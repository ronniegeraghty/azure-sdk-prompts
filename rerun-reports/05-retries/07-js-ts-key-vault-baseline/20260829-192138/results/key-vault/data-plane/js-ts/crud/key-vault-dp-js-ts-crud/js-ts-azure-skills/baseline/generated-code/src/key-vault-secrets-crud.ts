import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

const secretName = "my-secret";

function getVaultUrl(): string {
  const vaultUrl = process.env.AZURE_KEY_VAULT_URL;

  if (!vaultUrl) {
    throw new Error(
      "AZURE_KEY_VAULT_URL is required (for example, https://my-vault.vault.azure.net).",
    );
  }

  try {
    return new URL(vaultUrl).toString();
  } catch {
    throw new Error("AZURE_KEY_VAULT_URL must be a valid URL.");
  }
}

function formatError(error: unknown): string {
  if (error instanceof Error) {
    const statusCode =
      "statusCode" in error && typeof error.statusCode === "number"
        ? ` (HTTP ${error.statusCode})`
        : "";
    return `${error.name}${statusCode}: ${error.message}`;
  }

  return String(error);
}

async function main(): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new SecretClient(getVaultUrl(), credential);

  try {
    console.log(`Creating secret "${secretName}"...`);
    await client.setSecret(secretName, "my-secret-value");

    const createdSecret = await client.getSecret(secretName);
    if (createdSecret.value === undefined) {
      throw new Error(`Secret "${secretName}" was read without a value.`);
    }
    console.log(`Secret value: ${createdSecret.value}`);

    console.log(`Updating secret "${secretName}"...`);
    await client.setSecret(secretName, "updated-value");

    console.log(`Deleting secret "${secretName}"...`);
    const deletePoller = await client.beginDeleteSecret(secretName);
    await deletePoller.pollUntilDone();

    console.log(`Purging secret "${secretName}"...`);
    await client.purgeDeletedSecret(secretName);

    console.log("Secret CRUD operations completed successfully.");
  } catch (error: unknown) {
    console.error(`Key Vault operation failed: ${formatError(error)}`);
    process.exitCode = 1;
  }
}

void main();
