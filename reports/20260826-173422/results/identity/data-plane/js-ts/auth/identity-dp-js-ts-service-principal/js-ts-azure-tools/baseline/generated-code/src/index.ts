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
  const keyVaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");

  const credential = new ClientSecretCredential(
    tenantId,
    clientId,
    clientSecret,
  );

  // Requesting a token directly makes invalid service-principal credentials
  // available as AuthenticationError before making the Key Vault request.
  await credential.getToken("https://vault.azure.net/.default");
  console.log("Service principal authentication succeeded.");

  const secretClient = new SecretClient(keyVaultUrl, credential);

  for await (const secret of secretClient.listPropertiesOfSecrets()) {
    console.log(`Key Vault access succeeded. Found secret: ${secret.name}`);
    return;
  }

  console.log("Key Vault access succeeded. The vault contains no secrets.");
}

main().catch((error: unknown) => {
  if (error instanceof AuthenticationError) {
    console.error(
      "Azure authentication failed. Check AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.",
    );
    console.error(error.message);
    process.exitCode = 1;
    return;
  }

  console.error("The Azure Key Vault operation failed:", error);
  process.exitCode = 1;
});
