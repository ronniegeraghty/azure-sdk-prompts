import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

const secretName = "my-secret";
const initialValue = "my-secret-value";
const updatedValue = "updated-value";

function getVaultUrl(): string {
  const vaultUrl = process.env.KEY_VAULT_URL;

  if (!vaultUrl) {
    throw new Error(
      "KEY_VAULT_URL is required (for example, https://your-vault.vault.azure.net).",
    );
  }

  let parsedUrl: URL;
  try {
    parsedUrl = new URL(vaultUrl);
  } catch {
    throw new Error("KEY_VAULT_URL must be a valid URL.");
  }

  if (parsedUrl.protocol !== "https:") {
    throw new Error("KEY_VAULT_URL must use HTTPS.");
  }

  return parsedUrl.toString().replace(/\/$/, "");
}

async function main(): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new SecretClient(getVaultUrl(), credential);

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
  if (error instanceof Error) {
    console.error(`Error: ${error.message}`);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
}
