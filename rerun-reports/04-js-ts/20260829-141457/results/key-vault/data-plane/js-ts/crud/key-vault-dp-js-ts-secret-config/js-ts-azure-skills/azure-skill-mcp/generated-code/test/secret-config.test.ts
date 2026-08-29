import assert from "node:assert/strict";
import { describe, it } from "node:test";
import type {
  GetSecretOptions,
  KeyVaultSecret,
  SetSecretOptions
} from "@azure/keyvault-secrets";
import { SecretCache } from "../src/secret-cache.js";
import {
  KeyVaultSecretProvider,
  type SecretReader
} from "../src/secret-provider.js";
import {
  type DeleteSecretPollerLike,
  SecretRotationHelper,
  type SecretWriter
} from "../src/secret-rotation.js";

function secret(name: string, value: string, expiresOn?: Date, version = "v1"): KeyVaultSecret {
  return {
    name,
    value,
    properties: {
      name,
      vaultUrl: "https://example.vault.azure.net",
      id: `https://example.vault.azure.net/secrets/${name}/${version}`,
      version,
      expiresOn
    }
  };
}

describe("KeyVaultSecretProvider", () => {
  it("returns the default only for a missing secret", async () => {
    const reader: SecretReader = {
      getSecret: async () => {
        throw { statusCode: 404 };
      }
    };
    const provider = new KeyVaultSecretProvider(reader);

    const result = await provider.getSecret("missing", { defaultValue: "fallback" });

    assert.deepEqual(result, {
      name: "missing",
      value: "fallback",
      version: undefined,
      found: false
    });
  });

  it("retrieves a requested secret version", async () => {
    let receivedOptions: GetSecretOptions | undefined;
    const reader: SecretReader = {
      getSecret: async (name, options) => {
        receivedOptions = options;
        return secret(name, "old-value", undefined, "v2");
      }
    };
    const provider = new KeyVaultSecretProvider(reader);

    const result = await provider.getSecret("api-key", {
      defaultValue: "fallback",
      version: "v2"
    });

    assert.equal(receivedOptions?.version, "v2");
    assert.equal(result.value, "old-value");
    assert.equal(result.version, "v2");
  });

  it("does not hide non-404 service failures", async () => {
    const failure = { statusCode: 403 };
    const provider = new KeyVaultSecretProvider({
      getSecret: async () => {
        throw failure;
      }
    });

    await assert.rejects(
      provider.getSecret("forbidden", { defaultValue: "fallback" }),
      (error) => error === failure
    );
  });
});

describe("SecretCache", () => {
  it("bulk-loads, caches, and refreshes individual values", async () => {
    let calls = 0;
    const provider = new KeyVaultSecretProvider({
      getSecret: async (name) => secret(name, `value-${++calls}`)
    });
    const cache = new SecretCache(provider, 7 * 24 * 60 * 60 * 1000);

    await cache.bulkLoad([{ name: "api-key", defaultValue: "fallback" }]);
    assert.equal(await cache.get("api-key"), "value-1");
    assert.equal(calls, 1);

    assert.equal((await cache.refresh("api-key")).value, "value-2");
    assert.equal(calls, 2);
  });

  it("automatically re-fetches a cached secret near expiry", async () => {
    let calls = 0;
    const provider = new KeyVaultSecretProvider({
      getSecret: async (name) => {
        calls += 1;
        return calls === 1
          ? secret(name, "expiring", new Date(Date.now() + 60_000))
          : secret(name, "rotated", new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), "v2");
      }
    });
    const cache = new SecretCache(provider, 7 * 24 * 60 * 60 * 1000);

    await cache.bulkLoad([{ name: "api-key", defaultValue: "fallback" }]);

    assert.equal(await cache.get("api-key"), "rotated");
    assert.equal(calls, 2);
  });
});

describe("SecretRotationHelper", () => {
  it("creates a new version with expiry", async () => {
    let options: SetSecretOptions | undefined;
    const writer = {
      setSecret: async (name: string, value: string, received?: SetSecretOptions) => {
        options = received;
        return secret(name, value, received?.expiresOn, "new-version");
      },
      beginDeleteSecret: async () => {
        throw new Error("not used");
      },
      purgeDeletedSecret: async () => undefined
    } satisfies SecretWriter;
    const helper = new SecretRotationHelper(writer);
    const expiresOn = new Date(Date.now() + 86_400_000);

    const result = await helper.rotateSecret("api-key", "new-value", expiresOn);

    assert.equal(options?.expiresOn, expiresOn);
    assert.equal(result.version, "new-version");
  });

  it("waits for deletion before purge", async () => {
    const events: string[] = [];
    const poller = {
      pollUntilDone: async () => {
        events.push("delete-complete");
        return {};
      }
    } satisfies DeleteSecretPollerLike;
    const writer: SecretWriter = {
      setSecret: async (name, value) => secret(name, value),
      beginDeleteSecret: async () => {
        events.push("delete-start");
        return poller;
      },
      purgeDeletedSecret: async () => {
        events.push("purge");
      }
    };
    const helper = new SecretRotationHelper(writer);

    await helper.deleteAndPurgeSecret("disposable-secret", true);

    assert.deepEqual(events, ["delete-start", "delete-complete", "purge"]);
  });

  it("requires explicit confirmation before permanent purge", async () => {
    const writer = {
      setSecret: async (name: string, value: string) => secret(name, value),
      beginDeleteSecret: async () => {
        throw new Error("must not be called");
      },
      purgeDeletedSecret: async () => {
        throw new Error("must not be called");
      }
    } satisfies SecretWriter;

    await assert.rejects(
      new SecretRotationHelper(writer).deleteAndPurgeSecret("api-key", false),
      /not confirmed/
    );
  });
});
