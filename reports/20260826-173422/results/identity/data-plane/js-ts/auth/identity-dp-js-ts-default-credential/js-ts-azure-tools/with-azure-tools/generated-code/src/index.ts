import {
  AggregateAuthenticationError,
  AuthenticationError,
  DefaultAzureCredential,
} from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

function getRequiredEnvironmentVariable(name: string): string {
  const value = process.env[name];

  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

async function main(): Promise<void> {
  const vaultUrl = getRequiredEnvironmentVariable("KEY_VAULT_URL");
  const secretName = getRequiredEnvironmentVariable("SECRET_NAME");

  const credential = new DefaultAzureCredential();
  const secretClient = new SecretClient(vaultUrl, credential);

  const secret = await secretClient.getSecret(secretName);

  if (secret.value === undefined) {
    throw new Error(`Secret '${secretName}' does not contain a value.`);
  }

  console.log(secret.value);
}

try {
  await main();
} catch (error: unknown) {
  if (
    error instanceof AuthenticationError ||
    error instanceof AggregateAuthenticationError
  ) {
    console.error(`Azure authentication failed: ${error.message}`);
  } else {
    console.error("Failed to retrieve the secret:", error);
  }

  process.exitCode = 1;
}
