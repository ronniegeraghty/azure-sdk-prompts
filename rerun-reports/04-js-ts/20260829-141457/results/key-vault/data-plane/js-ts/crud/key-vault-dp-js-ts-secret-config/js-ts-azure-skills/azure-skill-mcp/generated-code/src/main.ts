import { createApplicationConfiguration } from "./configuration.js";
import { type RequiredConfigKey } from "./secret-cache.js";
import { SecretRotationHelper } from "./secret-rotation.js";

const DAY_MS = 24 * 60 * 60 * 1000;

function parseSecretNames(value: string | undefined): readonly string[] {
  const configured = value ?? "DatabaseConnectionString,ExternalApiKey,FeatureFlag";
  const names = configured
    .split(",")
    .map((name) => name.trim())
    .filter(Boolean);

  if (names.length === 0) {
    throw new Error("CONFIG_SECRET_NAMES must contain at least one secret name.");
  }
  return names;
}

function requirePositiveNumber(value: string | undefined, fallback: number, name: string): number {
  const number = value === undefined ? fallback : Number(value);
  if (!Number.isFinite(number) || number <= 0) {
    throw new Error(`${name} must be a positive number.`);
  }
  return number;
}

function isEnabled(value: string | undefined): boolean {
  return value?.toLowerCase() === "true";
}

function redact(value: string): string {
  return value.length === 0 ? "<empty>" : `<redacted:${value.length} chars>`;
}

async function main(): Promise<void> {
  const configuration = createApplicationConfiguration();
  const { cache, provider, client } = configuration;
  const requiredKeys: readonly RequiredConfigKey[] = parseSecretNames(
    process.env.CONFIG_SECRET_NAMES
  ).map((name) => ({ name, defaultValue: `<default:${name}>` }));

  console.log("1. Bulk-loading required configuration keys...");
  const loaded = await cache.bulkLoad(requiredKeys);
  for (const [name, entry] of loaded) {
    console.log(
      `   ${name}: ${redact(entry.value)}, source=${entry.found ? "Key Vault" : "default"}, `
      + `version=${entry.version ?? "none"}, expires=${entry.expiresOn?.toISOString() ?? "never"}`
    );
  }

  console.log("\n2. Reading values from the in-memory cache...");
  for (const key of requiredKeys) {
    console.log(`   ${key.name}: ${redact(await cache.get(key.name))}`);
  }

  const refreshTarget = requiredKeys[0];
  if (!refreshTarget) {
    throw new Error("No required configuration keys were configured.");
  }
  console.log(`\n3. Refreshing "${refreshTarget.name}" on demand...`);
  const refreshed = await cache.refresh(refreshTarget.name);
  console.log(
    `   Refreshed version=${refreshed.version ?? "none"}, cachedAt=${refreshed.cachedAt.toISOString()}`
  );

  console.log("\n4. Checking for secrets near expiry...");
  const expiring = cache.getExpiringSecrets();
  if (expiring.length === 0) {
    console.log("   No cached secrets are inside the expiry warning window.");
  } else {
    for (const entry of expiring) {
      console.warn(`   WARNING: ${entry.name} expires at ${entry.expiresOn?.toISOString()}.`);
    }
    const autoRefreshed = await cache.refreshExpiring();
    console.log(`   Automatically re-fetched ${autoRefreshed.length} expiring secret(s).`);
  }

  const stopAutoRefresh = cache.startAutoRefresh(
    configuration.autoRefreshIntervalMs,
    (error) => console.error("Automatic secret refresh failed:", error)
  );
  console.log(
    `   Periodic expiry refresh enabled every ${configuration.autoRefreshIntervalMs / 60_000} minute(s).`
  );

  const versionedName = process.env.VERSIONED_SECRET_NAME;
  const version = process.env.SECRET_VERSION;
  if (versionedName && version) {
    console.log(`\n5. Retrieving version "${version}" of "${versionedName}"...`);
    const versioned = await provider.getSecret(versionedName, {
      version,
      defaultValue: `<default:${versionedName}>`
    });
    console.log(
      `   ${versioned.name}: ${redact(versioned.value)}, found=${versioned.found}, `
      + `expires=${versioned.expiresOn?.toISOString() ?? "never"}`
    );
  } else {
    console.log("\n5. Version retrieval skipped; set VERSIONED_SECRET_NAME and SECRET_VERSION.");
  }

  const rotation = new SecretRotationHelper(client);
  if (isEnabled(process.env.ENABLE_ROTATION_DEMO)) {
    const name = process.env.ROTATE_SECRET_NAME;
    const value = process.env.ROTATE_SECRET_VALUE;
    if (!name || !value) {
      throw new Error(
        "ROTATE_SECRET_NAME and ROTATE_SECRET_VALUE are required when rotation is enabled."
      );
    }

    const expiryDays = requirePositiveNumber(
      process.env.ROTATE_SECRET_EXPIRY_DAYS,
      90,
      "ROTATE_SECRET_EXPIRY_DAYS"
    );
    const expiresOn = new Date(Date.now() + expiryDays * DAY_MS);
    console.log(`\n6. Rotating "${name}" by creating a new version...`);
    const result = await rotation.rotateSecret(name, value, expiresOn);
    console.log(
      `   Created version=${result.version ?? "unknown"}, expires=${result.expiresOn.toISOString()}.`
    );
  } else {
    console.log("\n6. Rotation skipped; set ENABLE_ROTATION_DEMO=true to run it.");
  }

  if (isEnabled(process.env.ENABLE_PURGE_DEMO)) {
    const name = process.env.PURGE_SECRET_NAME;
    if (!name) {
      throw new Error("PURGE_SECRET_NAME is required when purge is enabled.");
    }

    console.log(
      `\n7. Deleting and permanently purging "${name}" (all versions) after the delete LRO completes...`
    );
    await rotation.deleteAndPurgeSecret(name, true);
    console.log("   Delete completed and the deleted secret was purged.");
  } else {
    console.log(
      "\n7. Delete/purge skipped; set ENABLE_PURGE_DEMO=true and PURGE_SECRET_NAME to run it."
    );
  }

  stopAutoRefresh();
}

main().catch((error: unknown) => {
  console.error("Application failed:", error);
  process.exitCode = 1;
});
