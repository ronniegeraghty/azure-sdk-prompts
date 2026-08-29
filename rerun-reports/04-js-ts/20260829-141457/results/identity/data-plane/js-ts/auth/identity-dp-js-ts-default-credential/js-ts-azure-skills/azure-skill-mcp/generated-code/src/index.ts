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

function isAuthenticationFailure(
  error: unknown,
): error is AuthenticationError | AggregateAuthenticationError {
  return (
    error instanceof AuthenticationError ||
    (error instanceof AggregateAuthenticationError &&
      error.errors.some(
        (credentialError: unknown) =>
          credentialError instanceof AuthenticationError,
      ))
  );
}

async function main(): Promise<void> {
  const vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");
  const secretName = requireEnvironmentVariable("AZURE_KEY_VAULT_SECRET_NAME");

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);
  const secret = await client.getSecret(secretName);

  if (secret.value === undefined) {
    throw new Error(`Secret "${secretName}" does not contain a value.`);
  }

  console.log(secret.value);
}

try {
  await main();
} catch (error: unknown) {
  if (isAuthenticationFailure(error)) {
    console.error("Azure authentication failed:", error.message);
  } else {
    console.error("Failed to retrieve the secret:", error);
  }

  process.exitCode = 1;
}
