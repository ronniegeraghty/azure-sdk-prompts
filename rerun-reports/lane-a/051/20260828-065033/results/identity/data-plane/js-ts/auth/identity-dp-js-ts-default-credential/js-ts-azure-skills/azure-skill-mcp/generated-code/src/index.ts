import {
  AuthenticationError,
  DefaultAzureCredential,
} from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];

  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

async function main(): Promise<void> {
  const vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");
  const secretName = requireEnvironmentVariable("AZURE_KEY_VAULT_SECRET_NAME");

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);
  const secret = await client.getSecret(secretName);

  console.log(secret.value);
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof AuthenticationError) {
    console.error(`Azure authentication failed: ${error.message}`);
    process.exitCode = 1;
  } else {
    throw error;
  }
}
