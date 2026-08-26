import { createKeyVaultConfiguration } from "./config.js";
import { InMemoryKeyVaultClient } from "./in-memory-key-vault.js";
import { SecretCache, type RequiredSecrets } from "./secret-cache.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";
import { SecretRotator } from "./secret-rotator.js";

const DAY_MS = 24 * 60 * 60 * 1_000;
const WARNING_WINDOW_MS = 7 * DAY_MS;
const REQUIRED_SECRETS: RequiredSecrets = {
  "database-connection": "Server=localhost;Database=app",
  "service-api-key": "development-api-key",
  "missing-with-default": "safe-default",
};

async function createDemoConfiguration(): Promise<{
  provider: KeyVaultSecretProvider;
  cache: SecretCache;
  rotator: SecretRotator;
}> {
  if (process.env["DEMO_MODE"] === "azure") {
    const configuration = createKeyVaultConfiguration(
      REQUIRED_SECRETS,
      WARNING_WINDOW_MS,
    );
    return configuration;
  }

  const client = new InMemoryKeyVaultClient();
  await client.setSecret("database-connection", "Server=demo;Database=app", {
    expiresOn: new Date(Date.now() + 30 * DAY_MS),
  });
  await client.setSecret("service-api-key", "api-key-v1", {
    expiresOn: new Date(Date.now() + 3 * DAY_MS),
  });
  const provider = new KeyVaultSecretProvider(client);

  return {
    provider,
    cache: new SecretCache(provider, REQUIRED_SECRETS, WARNING_WINDOW_MS),
    rotator: new SecretRotator(client),
  };
}

async function main(): Promise<void> {
  const { provider, cache, rotator } = await createDemoConfiguration();

  console.log("1. Bulk-loading required configuration...");
  await cache.loadRequired();
  for (const entry of cache.inspectAll()) {
    console.log(
      `   ${entry.name}=${entry.value} (version=${entry.version ?? "default"})`,
    );
  }

  console.log("\n2. Reading database-connection from cache...");
  console.log(`   ${await cache.get("database-connection")}`);

  console.log("\n3. Refreshing database-connection on demand...");
  const refreshed = await cache.refresh("database-connection");
  console.log(`   Refreshed version ${refreshed.version ?? "default"}`);

  console.log("\n4. Checking for secrets near expiry...");
  for (const secret of cache.expiringSecrets()) {
    console.warn(
      `   WARNING: ${secret.name} expires on ${secret.expiresOn?.toISOString()}`,
    );
  }
  const automaticallyRefetched = await cache.refreshExpiring();
  console.log(
    `   Automatically re-fetched ${automaticallyRefetched.length} secret(s).`,
  );

  console.log("\n5. Rotating service-api-key to a new version...");
  const rotated = await rotator.rotate("service-api-key", "api-key-v2", {
    expiresOn: new Date(Date.now() + 90 * DAY_MS),
  });
  console.log(`   Created version ${rotated.secret.properties.version}.`);

  const previousVersion = cache.inspect("service-api-key")?.version;
  if (previousVersion !== undefined) {
    const previous = await provider.getSecret("service-api-key", {
      defaultValue: "not-found",
      version: previousVersion,
    });
    console.log(
      `   Previous version ${previousVersion} is still available as ${previous.value}.`,
    );
  }

  console.log("\n6. Demonstrating safe delete-and-purge name reuse...");
  if (process.env["DEMO_MODE"] === "azure") {
    console.log(
      "   Skipped in Azure mode. Set DEMO_ALLOW_PURGE=true only in an isolated test vault.",
    );
  } else {
    const cleaned = await rotator.rotate("service-api-key", "api-key-v3", {
      expiresOn: new Date(Date.now() + 120 * DAY_MS),
      cleanupForFullNameReuse: true,
    });
    console.log(
      `   Delete completed, purge completed, and version ${cleaned.secret.properties.version} was recreated.`,
    );
  }
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
