import "dotenv/config";

import {
  AuthenticationError,
  ClientSecretCredential,
} from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();

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

  // Request one page of secret metadata to verify authentication and read access.
  const iterator = secretClient
    .listPropertiesOfSecrets()
    .byPage({ maxPageSize: 1 });
  await iterator.next();

  console.log("Authentication succeeded and Key Vault is accessible.");
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof AuthenticationError) {
    console.error(
      "Authentication failed. Check AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.",
    );
    console.error(error.message);
    process.exitCode = 1;
  } else {
    throw error;
  }
}
