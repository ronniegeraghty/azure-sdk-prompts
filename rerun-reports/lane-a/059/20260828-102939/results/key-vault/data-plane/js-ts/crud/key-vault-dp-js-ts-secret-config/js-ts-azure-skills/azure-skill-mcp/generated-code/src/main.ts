import { CachingSecretProvider, type CachedSecret } from "./caching-secret-provider.js";
import { InMemorySecretClient } from "./in-memory-secret-client.js";
import { createKeyVaultSecretClient } from "./key-vault-config.js";
import type { SecretClientLike } from "./secret-client.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";
import { SecretRotationHelper } from "./secret-rotation.js";

const DAY_MS = 24 * 60 * 60 * 1_000;

function mask(value: string): string {
  if (!value) {
    return "(empty)";
  }
  return value.length <= 4 ? "****" : `${value.slice(0, 2)}***${value.slice(-2)}`;
}

function printSecret(label: string, secret: CachedSecret): void {
  console.log(
    `${label}: name=${secret.name}, value=${mask(secret.value)}, version=${secret.version ?? "n/a"}, ` +
      `source=${secret.found ? "vault" : "default"}, expires=${secret.expiresOn?.toISOString() ?? "none"}`,
  );
}

async function createDemoClient(mode: string): Promise<SecretClientLike> {
  if (mode === "azure") {
    return createKeyVaultSecretClient();
  }
  if (mode !== "mock") {
    throw new Error('KEY_VAULT_DEMO_MODE must be either "mock" or "azure".');
  }

  const client = new InMemorySecretClient();
  const now = Date.now();
  await client.setSecret("database-password", "local-db-password", {
    expiresOn: new Date(now + 30 * DAY_MS),
  });
  await client.setSecret("api-key", "local-api-key", {
    expiresOn: new Date(now + 3 * DAY_MS),
  });
  return client;
}

async function main(): Promise<void> {
  const mode = process.env.KEY_VAULT_DEMO_MODE ?? "mock";
  const warningDays = Number(process.env.SECRET_EXPIRY_WARNING_DAYS ?? "7");
  if (!Number.isFinite(warningDays) || warningDays < 0) {
    throw new Error("SECRET_EXPIRY_WARNING_DAYS must be a non-negative number.");
  }

  console.log(`1. Creating secret client in ${mode} mode.`);
  const client = await createDemoClient(mode);
  const provider = new KeyVaultSecretProvider(client);
  const cache = new CachingSecretProvider(provider, warningDays * DAY_MS);
  const rotation = new SecretRotationHelper(client);

  console.log("2. Bulk-loading required configuration.");
  await cache.loadRequired([
    { name: "database-password", defaultValue: "database-password-not-configured" },
    { name: "api-key", defaultValue: "api-key-not-configured" },
    { name: "optional-feature-token", defaultValue: "feature-disabled" },
  ]);
  for (const secret of cache.snapshot().values()) {
    printSecret("   loaded", secret);
  }

  console.log("3. Reading database-password from the in-memory cache.");
  printSecret("   cached", await cache.get("database-password"));

  console.log("4. Refreshing database-password on demand.");
  printSecret("   refreshed", await cache.refresh("database-password"));

  console.log(`5. Checking for secrets expiring within ${warningDays} day(s).`);
  const nearExpiry = cache.getNearExpiry();
  if (nearExpiry.length === 0) {
    console.log("   No cached secrets are near expiry.");
  } else {
    for (const secret of nearExpiry) {
      console.warn(`   WARNING: ${secret.name} expires on ${secret.expiresOn?.toISOString()}.`);
    }
    const refreshed = await cache.refreshNearExpiry();
    console.log(`   Automatically re-fetched ${refreshed.length} near-expiry secret(s).`);
  }

  console.log("6. Rotating api-key by creating a new Key Vault secret version.");
  const expiresOn = new Date(Date.now() + 90 * DAY_MS);
  const rotated = await rotation.rotateSecret("api-key", `rotated-${Date.now()}`, expiresOn);
  console.log(
    `   Created version=${rotated.version ?? "n/a"}, expires=${rotated.expiresOn.toISOString()}.`,
  );
  printSecret("   refreshed rotation", await cache.refresh("api-key"));

  console.log("7. Demonstrating long-running delete followed by purge.");
  const destructiveCleanupAllowed =
    mode === "mock" || process.env.RUN_DESTRUCTIVE_CLEANUP === "true";
  if (!destructiveCleanupAllowed) {
    console.log(
      "   Skipped for Azure. Set RUN_DESTRUCTIVE_CLEANUP=true only when permanently deleting all versions is intended.",
    );
  } else {
    await rotation.deleteAndPurgeSecret("api-key", "api-key");
    console.log("   Delete completed and the soft-deleted secret was purged.");
    const afterPurge = await provider.getSecret("api-key", "api-key-not-configured");
    console.log(`   Post-purge lookup source=${afterPurge.found ? "vault" : "default"}.`);
  }
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
