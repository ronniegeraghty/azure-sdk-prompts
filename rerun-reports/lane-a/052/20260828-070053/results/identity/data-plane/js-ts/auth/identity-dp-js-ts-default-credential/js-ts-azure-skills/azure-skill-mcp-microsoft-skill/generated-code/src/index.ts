import {
  AggregateAuthenticationError,
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
  const vaultUrl = requireEnvironmentVariable("KEY_VAULT_URL");
  const secretName = requireEnvironmentVariable("SECRET_NAME");

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);

  try {
    const secret = await client.getSecret(secretName);
    if (secret.value === undefined) {
      throw new Error(`Secret '${secretName}' has no value.`);
    }

    console.log(secret.value);
  } catch (error: unknown) {
    if (
      error instanceof AuthenticationError ||
      error instanceof AggregateAuthenticationError
    ) {
      console.error("Azure authentication failed:", error.message);
      process.exitCode = 1;
      return;
    }

    throw error;
  }
}

await main();
