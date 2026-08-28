import assert from "node:assert/strict";
import test from "node:test";
import { CachingSecretProvider } from "../src/caching-secret-provider.js";
import { InMemorySecretClient } from "../src/in-memory-secret-client.js";
import { KeyVaultSecretProvider } from "../src/secret-provider.js";
import { SecretRotationHelper } from "../src/secret-rotation.js";

const DAY_MS = 24 * 60 * 60 * 1_000;

test("returns a default only when the secret does not exist", async () => {
  const provider = new KeyVaultSecretProvider(new InMemorySecretClient());
  const secret = await provider.getSecret("missing", "fallback");

  assert.equal(secret.value, "fallback");
  assert.equal(secret.found, false);
});

test("retrieves a specific secret version", async () => {
  const client = new InMemorySecretClient();
  const first = await client.setSecret("setting", "first");
  await client.setSecret("setting", "second");
  const provider = new KeyVaultSecretProvider(client);

  const secret = await provider.getSecretVersion(
    "setting",
    first.properties.version ?? "",
    "fallback",
  );

  assert.equal(secret.value, "first");
});

test("bulk-loads, caches, refreshes, and detects near-expiry secrets", async () => {
  const client = new InMemorySecretClient();
  await client.setSecret("soon", "old", { expiresOn: new Date(Date.now() + DAY_MS) });
  const cache = new CachingSecretProvider(new KeyVaultSecretProvider(client), 7 * DAY_MS);

  await cache.loadRequired([{ name: "soon", defaultValue: "fallback" }]);
  assert.equal(cache.getNearExpiry().length, 1);

  await client.setSecret("soon", "new", { expiresOn: new Date(Date.now() + 30 * DAY_MS) });
  const refreshed = await cache.get("soon");

  assert.equal(refreshed.value, "new");
  assert.equal(cache.getNearExpiry().length, 0);
});

test("rotation creates a version and cleanup waits before purging all versions", async () => {
  const events: string[] = [];
  const client = new InMemorySecretClient();
  await client.setSecret("rotating", "old");
  const rotation = new SecretRotationHelper({
    getSecret: client.getSecret.bind(client),
    setSecret: client.setSecret.bind(client),
    beginDeleteSecret: async (name) => {
      const poller = await client.beginDeleteSecret(name);
      events.push("delete-started");
      return {
        pollUntilDone: async () => {
          await poller.pollUntilDone();
          events.push("delete-completed");
        },
      };
    },
    purgeDeletedSecret: async (name) => {
      events.push("purge-started");
      await client.purgeDeletedSecret(name);
    },
  });

  const result = await rotation.rotateSecret(
    "rotating",
    "new",
    new Date(Date.now() + 30 * DAY_MS),
  );
  assert.ok(result.version);

  await rotation.deleteAndPurgeSecret("rotating", "rotating");
  assert.deepEqual(events, ["delete-started", "delete-completed", "purge-started"]);
});
