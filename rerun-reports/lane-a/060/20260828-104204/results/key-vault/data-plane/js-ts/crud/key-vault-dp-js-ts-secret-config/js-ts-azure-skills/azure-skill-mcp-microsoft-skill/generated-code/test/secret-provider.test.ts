import assert from "node:assert/strict";
import test from "node:test";
import type { GetSecretOptions, KeyVaultSecret } from "@azure/keyvault-secrets";
import { CachedSecretProvider } from "../src/secret-cache.js";
import { KeyVaultSecretProvider, type SecretClientReader } from "../src/secret-provider.js";

function secret(
  name: string,
  value: string,
  version: string,
  expiresOn?: Date,
): KeyVaultSecret {
  return {
    name,
    value,
    properties: {
      name,
      vaultUrl: "https://unit-test.vault.azure.net",
      version,
      ...(expiresOn === undefined ? {} : { expiresOn }),
    },
  };
}

test("returns a default value only when a secret is not found", async () => {
  const client: SecretClientReader = {
    async getSecret(): Promise<KeyVaultSecret> {
      throw { statusCode: 404, code: "SecretNotFound" };
    },
  };
  const provider = new KeyVaultSecretProvider(client);

  const result = await provider.getSecret("missing", "fallback");

  assert.deepEqual(result, {
    name: "missing",
    value: "fallback",
    found: false,
    usedDefault: true,
  });
});

test("passes a requested version and exposes expiry", async () => {
  let requestedVersion: string | undefined;
  const expiresOn = new Date("2030-01-01T00:00:00.000Z");
  const client: SecretClientReader = {
    async getSecret(_name: string, options?: GetSecretOptions): Promise<KeyVaultSecret> {
      requestedVersion = options?.version;
      return secret("setting", "value", "v1", expiresOn);
    },
  };
  const provider = new KeyVaultSecretProvider(client);

  const result = await provider.getSecret("setting", "", "v1");

  assert.equal(requestedVersion, "v1");
  assert.equal(result.version, "v1");
  assert.equal(result.expiresOn, expiresOn);
});

test("caches values and automatically refreshes near-expiry secrets", async () => {
  const now = new Date("2030-01-01T00:00:00.000Z");
  let calls = 0;
  const provider = {
    async getSecret(name: string) {
      calls += 1;
      return {
        name,
        value: `value-${calls}`,
        found: true,
        usedDefault: false,
        version: `v${calls}`,
        expiresOn: new Date("2030-01-03T00:00:00.000Z"),
      };
    },
  };
  const cache = new CachedSecretProvider(provider, 7, () => now);

  await cache.loadRequired([{ name: "setting" }]);
  const refreshed = await cache.get("setting");

  assert.equal(calls, 2);
  assert.equal(refreshed.value, "value-2");
});

test("keeps non-expiring values in cache", async () => {
  let calls = 0;
  const provider = {
    async getSecret(name: string) {
      calls += 1;
      return {
        name,
        value: "stable",
        found: true,
        usedDefault: false,
      };
    },
  };
  const cache = new CachedSecretProvider(provider);

  await cache.loadRequired([{ name: "setting" }]);
  const cached = await cache.get("setting");

  assert.equal(calls, 1);
  assert.equal(cached.value, "stable");
});
