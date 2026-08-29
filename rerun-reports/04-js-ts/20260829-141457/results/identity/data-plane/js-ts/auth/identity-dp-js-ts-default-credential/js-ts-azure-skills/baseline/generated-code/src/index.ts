import {
  AuthenticationError,
  DefaultAzureCredential,
} from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

async function main(): Promise<void> {
  const vaultUrl = process.env.AZURE_KEY_VAULT_URL;
  const secretName = process.env.AZURE_KEY_VAULT_SECRET_NAME;

  if (!vaultUrl || !secretName) {
    throw new Error(
      "Set AZURE_KEY_VAULT_URL and AZURE_KEY_VAULT_SECRET_NAME before running the program.",
    );
  }

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);

  try {
    const secret = await client.getSecret(secretName);
    console.log(secret.value);
  } catch (error: unknown) {
    if (error instanceof AuthenticationError) {
      console.error(`Azure authentication failed: ${error.message}`);
      process.exitCode = 1;
      return;
    }

    throw error;
  }
}

await main();
