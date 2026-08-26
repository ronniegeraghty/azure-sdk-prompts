import assert from "node:assert/strict";
import test from "node:test";
import { InMemorySecretClient } from "../demo/in-memory-secret-client.js";
import { SecretCache } from "../secret-cache.js";
import { SecretProvider } from "../secret-provider.js";
import { SecretRotationHelper } from "../secret-rotation.js";

const DAY_MS = 86_400_000;

test("provider supports defaults, versions, and expiry metadata", async () => {
  const client = new InMemorySecretClient();
  const expiry = new Date(Date.now() + DAY_MS);
  const first = await client.setSecret("setting", "one", { expiresOn: expiry });
  await client.setSecret("setting", "two");
  const provider = new SecretProvider(client);

  const missing = await provider.getSecret("missing", "fallback");
  assert.deepEqual(missing, {
    name: "missing",
    value: "fallback",
    found: false,
  });

  const versioned = await provider.getSecret(
    "setting",
    "fallback",
    first.properties.version,
  );
  assert.equal(versioned.value, "one");
  assert.equal(versioned.version, first.properties.version);
  assert.equal(versioned.expiresOn?.getTime(), expiry.getTime());
  assert.equal(provider.isNearExpiry(versioned, 2 * DAY_MS), true);
});

test("cache bulk-loads, refreshes, and re-fetches near-expiry entries", async () => {
  const client = new InMemorySecretClient();
  await client.setSecret("near", "one", {
    expiresOn: new Date(Date.now() + DAY_MS),
  });
  const cache = new SecretCache(new SecretProvider(client), 7 * DAY_MS);
  await cache.loadRequired([
    { name: "near", defaultValue: "fallback" },
    { name: "missing", defaultValue: "fallback" },
  ]);

  assert.equal(cache.getCached("missing").value, "fallback");
  const originalLoadedAt = cache.getCached("near").loadedAt;
  const refreshed = await cache.get("near");
  assert.ok(refreshed.loadedAt.getTime() >= originalLoadedAt.getTime());

  await client.setSecret("near", "two", {
    expiresOn: new Date(Date.now() + 30 * DAY_MS),
  });
  assert.equal((await cache.refresh("near")).value, "two");
  assert.equal(cache.findNearExpiry().length, 0);
});

test("rotation creates a version and cleanup waits before purge", async () => {
  const client = new InMemorySecretClient();
  const first = await client.setSecret("rotate-me", "one");
  const helper = new SecretRotationHelper(client);

  const result = await helper.rotate(
    "rotate-me",
    "two",
    new Date(Date.now() + 30 * DAY_MS),
  );
  assert.equal(result.previousVersion, first.properties.version);
  assert.equal((await client.getSecret("rotate-me")).value, "two");

  await assert.rejects(
    helper.deleteAndPurgeForNameReuse("rotate-me", "wrong-name"),
    /confirmation/,
  );
  await helper.deleteAndPurgeForNameReuse("rotate-me", "rotate-me");
  await assert.rejects(client.getSecret("rotate-me"), /not found/);

  const reused = await client.setSecret("rotate-me", "three");
  assert.equal(reused.value, "three");
});
