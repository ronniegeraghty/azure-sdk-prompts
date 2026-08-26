import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

const secretName = "my-secret";
const initialValue = "my-secret-value";
const updatedValue = "updated-value";

async function main(): Promise<void> {
  const vaultUrl = process.env.KEY_VAULT_URL;
  if (!vaultUrl) {
    throw new Error(
      "KEY_VAULT_URL is required (for example, https://<vault-name>.vault.azure.net).",
    );
  }

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);

  console.log(`Creating secret "${secretName}"...`);
  await client.setSecret(secretName, initialValue);

  const createdSecret = await client.getSecret(secretName);
  console.log(`Read secret value: ${createdSecret.value}`);

  console.log(`Updating secret "${secretName}"...`);
  await client.setSecret(secretName, updatedValue);

  const updatedSecret = await client.getSecret(secretName);
  console.log(`Updated secret value: ${updatedSecret.value}`);

  console.log(`Deleting secret "${secretName}"...`);
  const deletePoller = await client.beginDeleteSecret(secretName);
  await deletePoller.pollUntilDone();

  console.log(`Purging secret "${secretName}"...`);
  await client.purgeDeletedSecret(secretName);
  console.log("Secret deleted and purged.");
}

try {
  await main();
} catch (error: unknown) {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Key Vault secret CRUD failed: ${message}`);
  process.exitCode = 1;
}
