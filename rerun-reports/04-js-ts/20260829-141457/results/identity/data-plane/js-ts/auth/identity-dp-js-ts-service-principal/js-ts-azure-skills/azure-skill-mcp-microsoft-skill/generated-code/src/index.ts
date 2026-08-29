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

async function verifyCredential(): Promise<void> {
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

  // Reading one metadata page verifies authentication and Key Vault authorization
  // without retrieving or printing any secret values.
  const iterator = secretClient
    .listPropertiesOfSecrets()
    [Symbol.asyncIterator]();
  const firstResult = await iterator.next();

  const resultDescription = firstResult.done
    ? "the vault contains no secrets"
    : `found secret metadata for "${firstResult.value.name}"`;
  console.log(`Credential verified successfully: ${resultDescription}.`);
}

async function main(): Promise<void> {
  try {
    await verifyCredential();
  } catch (error: unknown) {
    if (error instanceof AuthenticationError) {
      console.error(
        "Authentication failed. Check AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.",
      );
      console.error(error.message);
      process.exitCode = 1;
      return;
    }

    throw error;
  }
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Unable to verify Key Vault access: ${message}`);
  process.exitCode = 1;
});
