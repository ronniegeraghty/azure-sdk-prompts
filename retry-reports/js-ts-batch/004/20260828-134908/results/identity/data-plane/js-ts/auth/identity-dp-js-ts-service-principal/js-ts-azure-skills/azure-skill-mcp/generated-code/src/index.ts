import "dotenv/config";

import {
  AuthenticationError,
  ClientSecretCredential,
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
  const tenantId = requireEnvironmentVariable("AZURE_TENANT_ID");
  const clientId = requireEnvironmentVariable("AZURE_CLIENT_ID");
  const clientSecret = requireEnvironmentVariable("AZURE_CLIENT_SECRET");
  const vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");

  const credential = new ClientSecretCredential(
    tenantId,
    clientId,
    clientSecret,
  );
  const secretClient = new SecretClient(vaultUrl, credential);

  // Fetching one page forces token acquisition and makes an authenticated request.
  await secretClient
    .listPropertiesOfSecrets()
    .byPage({ maxPageSize: 1 })
    .next();

  console.log("Authentication succeeded and Azure Key Vault is accessible.");
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof AuthenticationError) {
    console.error(
      "Azure authentication failed. Check AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.",
    );
  } else if (error instanceof Error) {
    console.error(`The Key Vault verification failed: ${error.message}`);
  } else {
    console.error("The Key Vault verification failed with an unknown error.");
  }

  process.exitCode = 1;
}
