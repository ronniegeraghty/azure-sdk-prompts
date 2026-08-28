import { createKeyVaultConfiguration } from "./configuration.js";
import { SecretRotationHelper } from "./secret-rotation.js";
import type { SecretLookup } from "./secret-provider.js";

const requiredSecrets: readonly SecretLookup[] = [
  { name: "database-connection-string", defaultValue: "not-configured" },
  { name: "external-api-key", defaultValue: "not-configured" },
  { name: "feature-flags", defaultValue: "{}" },
];

async function main(): Promise<void> {
  const warningDays = parsePositiveNumber(
    process.env.SECRET_EXPIRY_WARNING_DAYS,
    7,
  );
  const { client, cache } = createKeyVaultConfiguration({
    warningWindowMs: warningDays * 24 * 60 * 60 * 1_000,
  });

  console.log("1. Loading required configuration secrets...");
  await cache.loadRequired(requiredSecrets);
  for (const lookup of requiredSecrets) {
    const cached = cache.inspect(lookup.name);
    console.log(
      `   ${lookup.name}: ${cached?.usedDefault === true ? "default value" : `version ${cached?.version ?? "unknown"}`}`,
    );
  }

  console.log("\n2. Reading configuration from the in-memory cache...");
  for (const lookup of requiredSecrets) {
    console.log(`   ${lookup.name}=${redact(await cache.get(lookup.name))}`);
  }

  const refreshName = requiredSecrets[0]?.name;
  if (refreshName === undefined) {
    throw new Error("At least one required secret must be configured.");
  }

  console.log(`\n3. Refreshing "${refreshName}" on demand...`);
  const refreshed = await cache.refresh(refreshName);
  console.log(
    `   Refreshed ${refreshed.name} at version ${refreshed.version ?? "unknown"}.`,
  );

  console.log(`\n4. Checking for secrets expiring within ${warningDays} days...`);
  const expiringSecrets = cache.getExpiringSecrets();
  if (expiringSecrets.length === 0) {
    console.log("   No cached secrets are near expiry.");
  } else {
    for (const secret of expiringSecrets) {
      console.warn(
        `   WARNING: ${secret.name} expires ${secret.expiresOn?.toISOString() ?? "without a future expiry date"}.`,
      );
    }
  }

  const rotationName = process.env.ROTATION_SECRET_NAME?.trim();
  const rotationValue = process.env.ROTATION_SECRET_VALUE;
  if (
    rotationName === undefined ||
    rotationName === "" ||
    rotationValue === undefined
  ) {
    console.log(
      "\n5. Rotation skipped. Set ROTATION_SECRET_NAME and ROTATION_SECRET_VALUE to run it.",
    );
    return;
  }

  const rotation = new SecretRotationHelper(client);
  const expiryDays = parsePositiveNumber(
    process.env.ROTATION_EXPIRY_DAYS,
    90,
  );
  const expiresOn = new Date(Date.now() + expiryDays * 24 * 60 * 60 * 1_000);

  console.log(`\n5. Rotating "${rotationName}" by creating a new version...`);
  const rotated = await rotation.rotateSecret(rotationName, rotationValue, {
    expiresOn,
    tags: { rotatedBy: "key-vault-config-provider-demo" },
  });
  console.log(
    `   Created version ${rotated.properties.version ?? "unknown"} expiring ${expiresOn.toISOString()}.`,
  );

  if (process.env.ENABLE_DELETE_AND_PURGE_DEMO !== "true") {
    console.log(
      "   Delete-and-purge skipped. Set ENABLE_DELETE_AND_PURGE_DEMO=true to run the destructive cleanup demo.",
    );
    return;
  }

  console.log(
    `\n6. Deleting all versions of "${rotationName}", waiting for soft-delete, then purging...`,
  );
  await rotation.deleteAndPurgeSecret(rotationName);
  console.log(`   Deleted and purged "${rotationName}".`);
}

function parsePositiveNumber(value: string | undefined, fallback: number): number {
  if (value === undefined) {
    return fallback;
  }

  const parsed = Number(value);
  if (!Number.isFinite(parsed) || parsed <= 0) {
    throw new Error(`Expected a positive number, received "${value}".`);
  }

  return parsed;
}

function redact(value: string): string {
  if (value.length <= 4) {
    return "****";
  }

  return `${value.slice(0, 2)}${"*".repeat(Math.min(value.length - 4, 12))}${value.slice(-2)}`;
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
