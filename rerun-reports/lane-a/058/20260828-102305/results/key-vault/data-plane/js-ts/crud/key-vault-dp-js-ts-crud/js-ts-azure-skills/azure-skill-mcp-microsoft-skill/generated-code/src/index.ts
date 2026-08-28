import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

const secretName = "my-secret";

function getVaultUrl(): string {
  const vaultUrl = process.env.KEY_VAULT_URL;

  if (!vaultUrl) {
    throw new Error(
      "KEY_VAULT_URL is required (for example, https://my-vault.vault.azure.net).",
    );
  }

  let parsedUrl: URL;
  try {
    parsedUrl = new URL(vaultUrl);
  } catch {
    throw new Error("KEY_VAULT_URL must be a valid URL.");
  }

  if (
    parsedUrl.protocol !== "https:" ||
    !parsedUrl.hostname.endsWith(".vault.azure.net")
  ) {
    throw new Error(
      "KEY_VAULT_URL must be an HTTPS Azure Key Vault URL ending in .vault.azure.net.",
    );
  }

  return parsedUrl.toString().replace(/\/$/, "");
}

function describeError(error: unknown): string {
  if (!(error instanceof Error)) {
    return String(error);
  }

  const azureError = error as Error & {
    code?: string;
    statusCode?: number;
  };
  const details = [
    azureError.statusCode && `HTTP ${azureError.statusCode}`,
    azureError.code,
  ].filter(Boolean);

  return details.length > 0
    ? `${azureError.message} (${details.join(", ")})`
    : azureError.message;
}

async function main(): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new SecretClient(getVaultUrl(), credential);

  try {
    const created = await client.setSecret(secretName, "my-secret-value");
    console.log(
      `Created "${created.name}" (version ${created.properties.version}).`,
    );

    const read = await client.getSecret(secretName);
    if (read.value === undefined) {
      throw new Error(`Secret "${secretName}" was returned without a value.`);
    }
    console.log(`Read "${read.name}": ${read.value}`);

    const updated = await client.setSecret(secretName, "updated-value");
    console.log(
      `Updated "${updated.name}" (version ${updated.properties.version}).`,
    );

    const deletePoller = await client.beginDeleteSecret(secretName);
    await deletePoller.pollUntilDone();
    console.log(`Soft-deleted "${secretName}".`);

    await client.purgeDeletedSecret(secretName);
    console.log(`Purged "${secretName}".`);
  } catch (error: unknown) {
    console.error(`Key Vault operation failed: ${describeError(error)}`);
    throw error;
  }
}

main().catch(() => {
  process.exitCode = 1;
});
