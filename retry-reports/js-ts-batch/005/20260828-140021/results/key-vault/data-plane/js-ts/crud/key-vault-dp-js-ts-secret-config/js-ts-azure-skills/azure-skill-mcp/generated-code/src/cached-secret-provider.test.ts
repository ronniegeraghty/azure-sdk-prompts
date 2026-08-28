import assert from "node:assert/strict";
import test from "node:test";
import {
  CachedSecretProvider,
  type SecretProvider,
} from "./cached-secret-provider.js";
import type { SecretValue } from "./secret-provider.js";

const fixedNow = new Date("2026-08-28T00:00:00.000Z");
const warningWindowMs = 7 * 24 * 60 * 60 * 1_000;

test("bulk loading populates the cache and later reads reuse it", async () => {
  const provider = new FakeSecretProvider([
    secret("api-key", "initial", "2026-10-01T00:00:00.000Z"),
  ]);
  const cache = createCache(provider);

  await cache.loadRequired([{ name: "api-key", defaultValue: "fallback" }]);

  assert.equal(await cache.get("api-key"), "initial");
  assert.equal(provider.getCalls, 1);
});

test("refresh replaces an individual cached value", async () => {
  const provider = new FakeSecretProvider([
    secret("api-key", "initial", "2026-10-01T00:00:00.000Z"),
    secret("api-key", "refreshed", "2026-11-01T00:00:00.000Z"),
  ]);
  const cache = createCache(provider);
  await cache.loadRequired([{ name: "api-key", defaultValue: "fallback" }]);

  await cache.refresh("api-key");

  assert.equal(await cache.get("api-key"), "refreshed");
  assert.equal(provider.getCalls, 2);
});

test("a near-expiry cached secret is automatically fetched again", async () => {
  const provider = new FakeSecretProvider([
    secret("api-key", "expiring", "2026-09-01T00:00:00.000Z"),
    secret("api-key", "rotated", "2026-12-01T00:00:00.000Z"),
  ]);
  const cache = createCache(provider);
  await cache.loadRequired([{ name: "api-key", defaultValue: "fallback" }]);

  assert.deepEqual(
    cache.getExpiringSecrets().map(({ name }) => name),
    ["api-key"],
  );
  assert.equal(await cache.get("api-key"), "rotated");
  assert.equal(provider.getCalls, 2);
  assert.deepEqual(cache.getExpiringSecrets(), []);
});

function createCache(provider: SecretProvider): CachedSecretProvider {
  return new CachedSecretProvider(provider, {
    warningWindowMs,
    now: () => fixedNow,
  });
}

function secret(name: string, value: string, expiresOn: string): SecretValue {
  return {
    name,
    value,
    version: `${value}-version`,
    expiresOn: new Date(expiresOn),
    usedDefault: false,
  };
}

class FakeSecretProvider implements SecretProvider {
  public getCalls = 0;

  public constructor(private readonly responses: readonly SecretValue[]) {}

  public async getSecret(): Promise<SecretValue> {
    const response = this.responses[this.getCalls];
    this.getCalls += 1;

    if (response === undefined) {
      throw new Error("No fake response configured.");
    }

    return response;
  }

  public isExpiringWithin(
    value: Pick<SecretValue, "expiresOn">,
    windowMs: number,
    now = new Date(),
  ): boolean {
    return (
      value.expiresOn !== undefined &&
      value.expiresOn.getTime() - now.getTime() <= windowMs
    );
  }
}
