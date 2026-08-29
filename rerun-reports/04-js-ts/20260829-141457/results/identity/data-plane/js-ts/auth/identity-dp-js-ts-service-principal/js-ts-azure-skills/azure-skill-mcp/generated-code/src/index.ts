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

  // Fetch at most one secret's metadata to verify authentication and access.
  const firstSecret = await secretClient.listPropertiesOfSecrets().next();

  console.log("Authentication succeeded and Key Vault is accessible.");
  console.log(
    firstSecret.done
      ? "The vault contains no secrets."
      : `Found secret metadata for: ${firstSecret.value.name}`,
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
  } else if (error instanceof Error) {
    console.error(`Unable to verify the credential: ${error.message}`);
  } else {
    console.error("Unable to verify the credential due to an unknown error.");
  }

  process.exitCode = 1;
}
