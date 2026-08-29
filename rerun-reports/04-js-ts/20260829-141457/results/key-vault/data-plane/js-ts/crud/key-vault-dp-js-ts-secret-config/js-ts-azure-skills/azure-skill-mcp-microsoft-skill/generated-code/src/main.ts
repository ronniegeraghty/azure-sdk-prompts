import { CachedSecretProvider } from "./secret-cache.js";
import { createKeyVaultSecretClient } from "./configuration.js";
import { InMemorySecretClient } from "./in-memory-secret-client.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";
import { SecretRotationHelper } from "./secret-rotation.js";
import type { SecretClientLike, SecretSnapshot } from "./types.js";

const REQUIRED_CONFIG = {
  "database-connection": "Server=localhost;Database=app",
  "external-api-key": "development-api-key",
  "feature-flags": '{"newCheckout":false}',
} as const;

function summarize(secret: SecretSnapshot): string {
  return `${secret.name}: ${secret.found ? "Key Vault value" : "default value"} ` +
    `(version=${secret.version ?? "none"}, length=${secret.value.length})`;
}

async function createDemoClient(mode: string): Promise<SecretClientLike> {
  if (mode === "azure") {
    return createKeyVaultSecretClient();
  }
  if (mode !== "mock") {
    throw new Error("DEMO_MODE must be either 'mock' or 'azure'");
  }

  const client = new InMemorySecretClient();
  await client.setSecret("database-connection", "Server=demo;Database=app", {
    expiresOn: new Date(Date.now() + 90 * 24 * 60 * 60 * 1000),
  });
  await client.setSecret("external-api-key", "offline-demo-key", {
    expiresOn: new Date(Date.now() + 3 * 24 * 60 * 60 * 1000),
  });
  await client.setSecret("demo-rotating-secret", "version-one", {
    expiresOn: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000),
  });
  return client;
}

async function main(): Promise<void> {
  const mode = process.env.DEMO_MODE ?? "mock";
  const client = await createDemoClient(mode);
  const provider = new KeyVaultSecretProvider(client);
  const cache = new CachedSecretProvider(provider, 7);
  const rotation = new SecretRotationHelper(client);

  console.log(`1. Loading required configuration (mode=${mode})`);
  const loaded = await cache.loadRequired(REQUIRED_CONFIG);
  for (const secret of loaded.values()) {
    console.log(`   ${summarize(secret)}`);
  }

  console.log("\n2. Reading configuration from the in-memory cache");
  console.log(`   ${summarize(await cache.get("database-connection"))}`);
  console.log(`   ${summarize(await cache.get("feature-flags"))}`);

  console.log("\n3. Refreshing one key on demand");
  console.log(`   ${summarize(await cache.refresh("external-api-key"))}`);

  console.log("\n4. Inspecting expiry and automatically re-fetching near-expiry keys");
  const warnings = cache.getExpiryWarnings();
  if (warnings.length === 0) {
    console.log("   No cached secrets are within the 7-day warning window.");
  }
  for (const warning of warnings) {
    console.warn(
      `   WARNING: ${warning.name} expires on ${warning.expiresOn?.toISOString()}`,
    );
  }
  const refreshed = await cache.refreshExpiringSecrets();
  console.log(`   Automatically re-fetched ${refreshed.length} secret(s).`);

  console.log("\n5. Rotating a secret by creating a new version");
  const rotationName = "demo-rotating-secret";
  const rotated = await rotation.rotate(
    rotationName,
    "version-two",
    new Date(Date.now() + 90 * 24 * 60 * 60 * 1000),
  );
  console.log(
    `   ${rotationName}: ${rotated.previousVersion ?? "none"} -> ` +
      `${rotated.newVersion ?? "unknown"}`,
  );
  console.log(
    `   Previous version still available: ${
      rotated.previousVersion === undefined
        ? "none"
        : summarize(
            await provider.getSecret(
              rotationName,
              "missing",
              rotated.previousVersion,
            ),
          )
    }`,
  );

  console.log("\n6. Demonstrating delete-and-purge cleanup for full name reuse");
  const cleanupEnabled =
    mode === "mock" || process.env.RUN_DESTRUCTIVE_CLEANUP === "true";
  if (!cleanupEnabled) {
    console.log(
      "   Skipped in Azure mode. Set RUN_DESTRUCTIVE_CLEANUP=true and " +
        "PURGE_CONFIRM_SECRET_NAME to the exact name to enable it.",
    );
    return;
  }

  const confirmation =
    mode === "mock"
      ? rotationName
      : (process.env.PURGE_CONFIRM_SECRET_NAME ?? "");
  await rotation.deleteAndPurgeForNameReuse(rotationName, confirmation);
  const afterPurge = await provider.getSecret(rotationName, "deleted-default");
  console.log(
    `   Delete completed before purge; subsequent read returned ` +
      `${afterPurge.found ? "a secret" : "the configured default"}.`,
  );
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
