import { createKeyVaultServices } from "./configuration.js";
import type { SecretResult } from "./secret-provider.js";

const requiredConfig = [
  { name: "database-connection", defaultValue: "not-configured" },
  { name: "external-api-key", defaultValue: "not-configured" },
  { name: "feature-flag", defaultValue: "false" },
] as const;

async function main(): Promise<void> {
  const { cache, provider, rotation } = createKeyVaultServices();

  console.log("1. Loading required configuration keys...");
  const loaded = await cache.loadRequired(requiredConfig);
  for (const secret of loaded.values()) {
    printSecretResult(secret);
  }

  console.log("\n2. Reading configuration from the in-memory cache...");
  for (const { name } of requiredConfig) {
    const cached = cache.peek(name);
    console.log(`${name}: ${cached === undefined ? "cache miss" : describeSecret(cached)}`);
  }

  const refreshName = process.env.REFRESH_SECRET_NAME ?? requiredConfig[0].name;
  console.log(`\n3. Refreshing '${refreshName}' on demand...`);
  printSecretResult(await cache.refresh(refreshName));

  console.log("\n4. Checking cached secret expiry dates...");
  const expiryStatuses = cache.getExpiryStatuses();
  const warnings = expiryStatuses.filter(({ isNearExpiry }) => isNearExpiry);
  if (warnings.length === 0) {
    console.log("No cached secrets are within the configured expiry warning window.");
  } else {
    for (const warning of warnings) {
      console.warn(
        `WARNING: '${warning.name}' expires on ${warning.expiresOn?.toISOString() ?? "unknown"}.`,
      );
    }
  }

  const automaticallyRefreshed = await cache.refreshExpiring();
  console.log(
    `Automatic expiry refresh: ${
      automaticallyRefreshed.length === 0 ? "none" : automaticallyRefreshed.join(", ")
    }.`,
  );

  const versionName = process.env.DEMO_SECRET_VERSION_NAME;
  const version = process.env.DEMO_SECRET_VERSION;
  if (versionName !== undefined && version !== undefined) {
    console.log(`\n5. Reading version '${version}' of '${versionName}'...`);
    printSecretResult(await provider.getSecret(versionName, "not-configured", version));
  } else {
    console.log(
      "\n5. Specific-version read skipped; set DEMO_SECRET_VERSION_NAME and DEMO_SECRET_VERSION.",
    );
  }

  const rotationName = process.env.ROTATION_SECRET_NAME;
  const rotationValue = process.env.ROTATION_SECRET_VALUE;
  if (rotationName === undefined || rotationValue === undefined) {
    console.log(
      "\n6. Rotation skipped; set ROTATION_SECRET_NAME and ROTATION_SECRET_VALUE to create a version.",
    );
    return;
  }

  const expiryDays = parsePositiveNumber(process.env.ROTATION_EXPIRY_DAYS ?? "90");
  const expiresOn = new Date(Date.now() + expiryDays * 24 * 60 * 60 * 1_000);
  console.log(`\n6. Rotating '${rotationName}' by creating a new version...`);
  const rotated = await rotation.rotateSecret(rotationName, rotationValue, expiresOn);
  console.log(
    `Created version ${rotated.newVersion ?? "unknown"}; previous latest version was ${
      rotated.previousVersion ?? "none"
    }; expires ${rotated.expiresOn.toISOString()}.`,
  );

  if (process.env.PURGE_ROTATED_SECRET !== "true") {
    console.log(
      "7. Delete-and-purge cleanup skipped. Set PURGE_ROTATED_SECRET=true and " +
        "PURGE_CONFIRM_SECRET_NAME to the exact name to permanently remove every version.",
    );
    return;
  }

  console.log(
    `7. Permanently cleaning up '${rotationName}' (all versions) after soft-delete completes...`,
  );
  await rotation.deleteAndPurgeSecret(rotationName, {
    confirmPermanentDeletion: process.env.PURGE_CONFIRM_SECRET_NAME ?? "",
  });
  console.log(`Deleted and purged '${rotationName}'. The name can now be reused.`);
}

function printSecretResult(secret: SecretResult): void {
  console.log(`${secret.name}: ${describeSecret(secret)}`);
}

function describeSecret(secret: SecretResult): string {
  const source = secret.found ? "Key Vault" : "default value";
  const expiry = secret.expiresOn?.toISOString() ?? "not set";
  return `loaded from ${source}; value=${redact(secret.value)}; version=${
    secret.version ?? "none"
  }; expires=${expiry}`;
}

function redact(value: string): string {
  return value.length === 0 ? "<empty>" : `<redacted:${value.length} chars>`;
}

function parsePositiveNumber(value: string): number {
  const parsed = Number(value);
  if (!Number.isFinite(parsed) || parsed <= 0) {
    throw new Error("ROTATION_EXPIRY_DAYS must be a positive number.");
  }
  return parsed;
}

main().catch((error: unknown) => {
  console.error("Demo failed:", error);
  process.exitCode = 1;
});
