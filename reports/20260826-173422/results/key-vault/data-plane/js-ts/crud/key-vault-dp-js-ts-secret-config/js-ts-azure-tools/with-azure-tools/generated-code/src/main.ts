import { InMemorySecretClient } from "./demo/in-memory-secret-client.js";
import { SecretCache, type RequiredSecret } from "./secret-cache.js";
import { SecretProvider } from "./secret-provider.js";
import { SecretRotationHelper } from "./secret-rotation.js";

const DAY_MS = 24 * 60 * 60 * 1_000;
const WARNING_WINDOW_MS = 7 * DAY_MS;

async function main(): Promise<void> {
  // The demo is intentionally offline. Production code should call
  // createKeyVaultSecretClient() from configuration.ts.
  const client = new InMemorySecretClient();
  await seedDemoVault(client);

  const provider = new SecretProvider(client);
  const cache = new SecretCache(provider, WARNING_WINDOW_MS);
  const required: readonly RequiredSecret[] = [
    { name: "api-endpoint", defaultValue: "https://localhost:3000" },
    { name: "database-password", defaultValue: "development-only" },
    { name: "optional-feature", defaultValue: "disabled" },
  ];

  console.log("1. Bulk-loading required configuration");
  const loaded = await cache.loadRequired(required);
  for (const entry of loaded) {
    printEntry(entry);
  }

  console.log("\n2. Reading api-endpoint from the in-memory cache");
  printEntry(cache.getCached("api-endpoint"));

  console.log("\n3. Creating a newer api-endpoint version and refreshing it");
  await client.setSecret("api-endpoint", "https://api-v2.example.test", {
    expiresOn: daysFromNow(60),
  });
  printEntry(await cache.refresh("api-endpoint"));

  console.log("\n4. Checking and automatically refreshing near-expiry secrets");
  const expiring = cache.findNearExpiry();
  if (expiring.length === 0) {
    console.log("No cached secrets are near expiry.");
  } else {
    for (const entry of expiring) {
      console.log(
        `WARNING: ${entry.name} expires on ${entry.expiresOn?.toISOString()}`,
      );
    }
    const refreshed = await cache.refreshNearExpiry();
    console.log(`Re-fetched ${refreshed.length} near-expiry secret(s).`);
  }

  console.log("\n5. Rotating database-password to a new version");
  const rotation = new SecretRotationHelper(client);
  const rotated = await rotation.rotate(
    "database-password",
    "rotated-demo-value",
    daysFromNow(90),
  );
  console.log(
    `Created ${rotated.secret.properties.version}; previous version was ${rotated.previousVersion}.`,
  );
  printEntry(await cache.refresh("database-password"));

  console.log("\n6. Demonstrating explicit delete-and-purge cleanup");
  console.log(
    "Cleanup is name-scoped: it deletes every database-password version.",
  );
  await rotation.deleteAndPurgeForNameReuse(
    "database-password",
    "database-password",
  );
  console.log(
    "Delete LRO completed and the soft-deleted secret was purged; the name can now be reused.",
  );
}

async function seedDemoVault(client: InMemorySecretClient): Promise<void> {
  await client.setSecret("api-endpoint", "https://api.example.test", {
    expiresOn: daysFromNow(30),
  });
  await client.setSecret("database-password", "demo-password", {
    expiresOn: daysFromNow(3),
  });
}

function daysFromNow(days: number): Date {
  return new Date(Date.now() + days * DAY_MS);
}

function printEntry(entry: {
  name: string;
  value: string;
  found: boolean;
  version?: string;
  expiresOn?: Date;
}): void {
  const displayValue =
    entry.name === "api-endpoint" || !entry.found
      ? entry.value
      : `<redacted:${entry.value.length} chars>`;
  console.log({
    name: entry.name,
    value: displayValue,
    source: entry.found ? "key-vault" : "default",
    version: entry.version ?? "none",
    expiresOn: entry.expiresOn?.toISOString() ?? "none",
  });
}

await main();
