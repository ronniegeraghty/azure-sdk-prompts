import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

const secretName = "my-secret";

function getVaultUrl(): string {
  const vaultUrl = process.env.KEY_VAULT_URL;

  if (!vaultUrl) {
    throw new Error(
      "KEY_VAULT_URL is required (for example, https://<vault-name>.vault.azure.net).",
    );
  }

  const parsedUrl = new URL(vaultUrl);
  if (parsedUrl.protocol !== "https:") {
    throw new Error("KEY_VAULT_URL must use HTTPS.");
  }

  return parsedUrl.toString();
}

function describeError(error: unknown): string {
  if (error instanceof Error) {
    const details = error as Error & {
      code?: string;
      statusCode?: number;
    };
    const metadata = [
      details.statusCode ? `HTTP ${details.statusCode}` : undefined,
      details.code,
    ]
      .filter(Boolean)
      .join(", ");

    return metadata ? `${details.message} (${metadata})` : details.message;
  }

  return String(error);
}

async function main(): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new SecretClient(getVaultUrl(), credential);

  try {
    console.log(`Creating secret "${secretName}"...`);
    await client.setSecret(secretName, "my-secret-value");

    console.log(`Reading secret "${secretName}"...`);
    const secret = await client.getSecret(secretName);
    if (secret.value === undefined) {
      throw new Error(`Secret "${secretName}" was returned without a value.`);
    }
    console.log(`Secret value: ${secret.value}`);

    console.log(`Updating secret "${secretName}"...`);
    await client.setSecret(secretName, "updated-value");

    console.log(`Deleting secret "${secretName}"...`);
    const deletePoller = await client.beginDeleteSecret(secretName);
    await deletePoller.pollUntilDone();

    console.log(`Purging secret "${secretName}"...`);
    await client.purgeDeletedSecret(secretName);

    console.log("CRUD operations completed successfully.");
  } catch (error: unknown) {
    console.error(`Key Vault operation failed: ${describeError(error)}`);
    throw error;
  }
}

main().catch(() => {
  process.exitCode = 1;
});
