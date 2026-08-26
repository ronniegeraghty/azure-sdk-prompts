import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

const secretName = "my-secret";

async function main(): Promise<void> {
  const vaultUrl = process.env.AZURE_KEY_VAULT_URL;

  if (!vaultUrl) {
    throw new Error(
      "AZURE_KEY_VAULT_URL is required (for example, https://my-vault.vault.azure.net).",
    );
  }

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);

  try {
    // Create
    await client.setSecret(secretName, "my-secret-value");
    console.log(`Created secret "${secretName}".`);

    // Read
    const secret = await client.getSecret(secretName);
    console.log(`Secret value: ${secret.value}`);

    // Update (creates a new secret version)
    await client.setSecret(secretName, "updated-value");
    console.log(`Updated secret "${secretName}".`);

    // Delete and wait until the soft-deleted secret is available.
    const deletePoller = await client.beginDeleteSecret(secretName);
    await deletePoller.pollUntilDone();
    console.log(`Deleted secret "${secretName}".`);

    // Permanently remove the soft-deleted secret.
    await client.purgeDeletedSecret(secretName);
    console.log(`Purged secret "${secretName}".`);
  } catch (error: unknown) {
    const message = error instanceof Error ? error.message : String(error);
    console.error(`Key Vault operation failed: ${message}`);
    throw error;
  }
}

main().catch(() => {
  process.exitCode = 1;
});
