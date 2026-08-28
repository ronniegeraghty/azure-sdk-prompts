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
  const vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");

  const credential = new ClientSecretCredential(
    tenantId,
    clientId,
    clientSecret,
  );
  const secretClient = new SecretClient(vaultUrl, credential);

  // Fetching one page is a read-only request that verifies authentication and
  // that the service principal can list secret metadata.
  const firstPage = await secretClient
    .listPropertiesOfSecrets()
    .byPage({ maxPageSize: 1 })
    .next();

  const secretCount = firstPage.done ? 0 : firstPage.value.length;
  console.log(
    `Authentication succeeded. Key Vault returned ${secretCount} secret entr${secretCount === 1 ? "y" : "ies"} in the first page.`,
  );
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof AuthenticationError) {
    console.error(
      "Authentication failed. Verify AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.",
    );
    console.error(error.message);
    process.exitCode = 1;
  } else {
    throw error;
  }
}
