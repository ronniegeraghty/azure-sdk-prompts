import assert from "node:assert/strict";
import test from "node:test";
import type {
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions,
} from "@azure/keyvault-secrets";
import {
  SecretRotationHelper,
  type SecretRotationClient,
} from "../src/secret-rotation.js";

function keyVaultSecret(name: string, value: string, version: string): KeyVaultSecret {
  return {
    name,
    value,
    properties: {
      name,
      vaultUrl: "https://unit-test.vault.azure.net",
      version,
      contentType: "text/plain",
      tags: { owner: "test" },
    },
  };
}

test("rotation creates a new version and preserves metadata", async () => {
  let setOptions: SetSecretOptions | undefined;
  const client: SecretRotationClient = {
    async getSecret(_name: string, _options?: GetSecretOptions) {
      return keyVaultSecret("setting", "old", "v1");
    },
    async setSecret(name: string, value: string, options?: SetSecretOptions) {
      assert.equal(name, "setting");
      assert.equal(value, "new");
      setOptions = options;
      return keyVaultSecret(name, value, "v2");
    },
    async beginDeleteSecret() {
      return { async pollUntilDone() {} };
    },
    async purgeDeletedSecret() {},
  };
  const expiry = new Date(Date.now() + 86_400_000);

  const result = await new SecretRotationHelper(client).rotateSecret("setting", "new", expiry);

  assert.equal(result.previousVersion, "v1");
  assert.equal(result.newVersion, "v2");
  assert.equal(setOptions?.contentType, "text/plain");
  assert.equal(setOptions?.tags?.owner, "test");
});

test("cleanup waits for delete completion before purging", async () => {
  const events: string[] = [];
  const client: SecretRotationClient = {
    async getSecret() {
      return keyVaultSecret("setting", "value", "v1");
    },
    async setSecret(name: string, value: string) {
      return keyVaultSecret(name, value, "v2");
    },
    async beginDeleteSecret() {
      events.push("delete-started");
      return {
        async pollUntilDone() {
          events.push("delete-completed");
        },
      };
    },
    async purgeDeletedSecret() {
      events.push("purged");
    },
  };

  await new SecretRotationHelper(client).deleteAndPurgeSecret("setting", {
    confirmPermanentDeletion: "setting",
  });

  assert.deepEqual(events, ["delete-started", "delete-completed", "purged"]);
});

test("cleanup requires an exact secret-name confirmation", async () => {
  const client: SecretRotationClient = {
    async getSecret() {
      return keyVaultSecret("setting", "value", "v1");
    },
    async setSecret(name: string, value: string) {
      return keyVaultSecret(name, value, "v2");
    },
    async beginDeleteSecret() {
      return { async pollUntilDone() {} };
    },
    async purgeDeletedSecret() {},
  };

  await assert.rejects(
    new SecretRotationHelper(client).deleteAndPurgeSecret("setting", {
      confirmPermanentDeletion: "wrong-name",
    }),
    /Permanent deletion was not confirmed/,
  );
});
