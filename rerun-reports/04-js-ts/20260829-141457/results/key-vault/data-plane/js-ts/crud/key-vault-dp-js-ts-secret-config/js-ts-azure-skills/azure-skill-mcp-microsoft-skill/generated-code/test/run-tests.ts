import assert from "node:assert/strict";
import test from "node:test";
import { CachedSecretProvider } from "../src/secret-cache.js";
import { InMemorySecretClient } from "../src/in-memory-secret-client.js";
import { KeyVaultSecretProvider } from "../src/secret-provider.js";
import { SecretRotationHelper } from "../src/secret-rotation.js";

const NOW = new Date("2026-08-29T08:00:00.000Z");

test("provider returns defaults and retrieves a specific version", async () => {
  const client = new InMemorySecretClient();
  const provider = new KeyVaultSecretProvider(client, () => NOW);

  const missing = await provider.getSecret("missing", "fallback");
  assert.equal(missing.value, "fallback");
  assert.equal(missing.found, false);

  const first = await client.setSecret("versioned", "one");
  await client.setSecret("versioned", "two");
  const historical = await provider.getSecret(
    "versioned",
    "fallback",
    first.properties.version,
  );
  assert.equal(historical.value, "one");
});

test("cache bulk-loads, caches, refreshes, and re-fetches near expiry", async () => {
  const client = new InMemorySecretClient();
  await client.setSecret("stable", "value", {
    expiresOn: new Date("2026-10-01T00:00:00.000Z"),
  });
  await client.setSecret("expiring", "value", {
    expiresOn: new Date("2026-09-01T00:00:00.000Z"),
  });

  const provider = new KeyVaultSecretProvider(client, () => NOW);
  const cache = new CachedSecretProvider(provider, 7);
  await cache.loadRequired({ stable: "default", expiring: "default" });
  await cache.get("stable");
  assert.equal(client.getRequestCount("stable"), 1);

  await cache.get("expiring");
  assert.equal(client.getRequestCount("expiring"), 2);
  assert.deepEqual(
    cache.getExpiryWarnings().map(({ name }) => name),
    ["expiring"],
  );

  await cache.refresh("stable");
  assert.equal(client.getRequestCount("stable"), 2);
});

test("rotation creates a version and cleanup waits before purge", async () => {
  const client = new InMemorySecretClient();
  const first = await client.setSecret("rotate-me", "one");
  const helper = new SecretRotationHelper(client, () => NOW);

  const result = await helper.rotate(
    "rotate-me",
    "two",
    new Date("2026-12-01T00:00:00.000Z"),
  );
  assert.equal(result.previousVersion, first.properties.version);
  assert.notEqual(result.newVersion, result.previousVersion);

  await helper.deleteAndPurgeForNameReuse("rotate-me", "rotate-me");
  assert.deepEqual(client.operations.slice(-3), [
    "begin-delete:rotate-me",
    "delete-complete:rotate-me",
    "purge:rotate-me",
  ]);
});

test("cleanup requires exact-name confirmation", async () => {
  const client = new InMemorySecretClient();
  await client.setSecret("protected", "value");
  const helper = new SecretRotationHelper(client, () => NOW);
  await assert.rejects(
    helper.deleteAndPurgeForNameReuse("protected", "wrong-name"),
    /confirmation must exactly match/,
  );
});
