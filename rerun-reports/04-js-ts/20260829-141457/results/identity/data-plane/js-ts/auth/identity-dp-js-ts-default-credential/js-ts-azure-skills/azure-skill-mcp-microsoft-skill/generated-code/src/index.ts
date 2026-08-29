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
  try {
    const vaultUrl = requireEnvironmentVariable("KEY_VAULT_URL");
    const secretName = requireEnvironmentVariable("SECRET_NAME");

    const credential = new DefaultAzureCredential();
    const client = new SecretClient(vaultUrl, credential);
    const secret = await client.getSecret(secretName);

    if (secret.value === undefined) {
      throw new Error(`Secret "${secretName}" has no value.`);
    }

    console.log(secret.value);
  } catch (error: unknown) {
    if (
      error instanceof AuthenticationError ||
      error instanceof AggregateAuthenticationError
    ) {
      console.error(`Azure authentication failed: ${error.message}`);
    } else if (error instanceof Error) {
      console.error(`Unable to retrieve the secret: ${error.message}`);
    } else {
      console.error("Unable to retrieve the secret due to an unknown error.");
    }

    process.exitCode = 1;
  }
}

await main();
