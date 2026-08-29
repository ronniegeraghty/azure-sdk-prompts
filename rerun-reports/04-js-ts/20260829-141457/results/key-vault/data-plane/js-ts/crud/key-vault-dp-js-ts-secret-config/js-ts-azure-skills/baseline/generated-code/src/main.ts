import { InMemorySecretStore } from "./in-memory-secret-store.js";
import { SecretCache, type RequiredSecret } from "./secret-cache.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";
import { SecretRotationHelper } from "./secret-rotation.js";

const DAY_MS = 24 * 60 * 60 * 1_000;

async function main(): Promise<void> {
  console.log("Using the local in-memory store; no Azure resources are changed.");

  const store = new InMemorySecretStore();
  await store.setSecret("database-url", "postgres://demo.local/app", {
    expiresOn: new Date(Date.now() + 90 * DAY_MS),
  });
  await store.setSecret("api-key", "api-key-v1", {
    expiresOn: new Date(Date.now() + 5 * DAY_MS),
  });

  const provider = new KeyVaultSecretProvider(store);
  const cache = new SecretCache(provider, 7 * DAY_MS);
  const rotation = new SecretRotationHelper(store);
  const required: RequiredSecret[] = [
    { name: "database-url" },
    { name: "api-key" },
    { name: "optional-feature-token", defaultValue: "disabled" },
  ];

  console.log("\n1. Bulk-loading required configuration...");
  await cache.loadRequired(required);
  for (const [name, secret] of cache.snapshot()) {
    console.log(`   ${name} = ${secret.value} (found: ${secret.found})`);
  }

  console.log("\n2. Reading database-url from cache...");
  console.log(`   database-url = ${await cache.get("database-url")}`);

  console.log("\n3. Refreshing database-url on demand...");
  const refreshed = await cache.refresh("database-url");
  console.log(`   refreshed version ${refreshed.version}: ${refreshed.value}`);

  console.log("\n4. Checking cached secrets for near expiry...");
  const expiring = cache.expiringSoon();
  if (expiring.length === 0) {
    console.log("   No secrets are near expiry.");
  }
  for (const secret of expiring) {
    console.warn(
      `   WARNING: ${secret.name} expires on ${secret.expiresOn?.toISOString()}`,
    );
  }

  console.log("\n5. Rotating api-key by creating a new version...");
  const rotated = await rotation.rotate("api-key", "api-key-v2", {
    expiresOn: new Date(Date.now() + 180 * DAY_MS),
    tags: { rotatedBy: "offline-demo" },
  });
  console.log(
    `   created ${rotated.secret.properties.version}; previous versions: ${rotated.previousVersions.join(", ")}`,
  );
  const oldVersion = rotated.previousVersions.at(-1);
  if (oldVersion !== undefined) {
    const oldSecret = await provider.getSecretVersion("api-key", oldVersion);
    console.log(`   previous version ${oldVersion} is still readable: ${oldSecret.value}`);
  }
  console.log(`   latest value = ${(await cache.refresh("api-key")).value}`);

  console.log("\n6. Demonstrating full-name delete-and-purge cleanup...");
  console.log("   This removes every version, including the newly rotated version.");
  await rotation.deleteAndPurge("api-key");
  const afterPurge = await provider.getSecret("api-key", "missing");
  console.log(`   after purge = ${afterPurge.value} (found: ${afterPurge.found})`);
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
