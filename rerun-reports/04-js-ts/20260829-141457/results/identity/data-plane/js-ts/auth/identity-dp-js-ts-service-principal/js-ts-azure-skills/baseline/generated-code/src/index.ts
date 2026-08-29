import "dotenv/config";

import { AuthenticationError, ClientSecretCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];

  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

async function main(): Promise<void> {
  const tenantId = requireEnvironmentVariable("AZURE_TENANT_ID");
  const clientId = requireEnvironmentVariable("AZURE_CLIENT_ID");
  const clientSecret = requireEnvironmentVariable("AZURE_CLIENT_SECRET");
  const keyVaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");

  const credential = new ClientSecretCredential(
    tenantId,
    clientId,
    clientSecret,
  );
  const secretClient = new SecretClient(keyVaultUrl, credential);

  // Requesting the first page forces authentication without reading secret values.
  const iterator = secretClient.listPropertiesOfSecrets()[Symbol.asyncIterator]();
  await iterator.next();

  console.log("Authentication succeeded and Key Vault is accessible.");
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof AuthenticationError) {
    console.error(
      `Azure authentication failed. Check the tenant ID, client ID, and client secret: ${error.message}`,
    );
    process.exitCode = 1;
  } else {
    console.error(
      "Unable to verify Azure access:",
      error instanceof Error ? error.message : error,
    );
    process.exitCode = 1;
  }
}
