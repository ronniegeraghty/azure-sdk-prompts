import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

export interface KeyVaultEnvironment {
  KEY_VAULT_URL?: string;
  AZURE_CLIENT_ID?: string;
}

export function createKeyVaultSecretClient(
  environment: KeyVaultEnvironment = process.env,
): SecretClient {
  const vaultUrl = environment.KEY_VAULT_URL;
  if (!vaultUrl) {
    throw new Error("KEY_VAULT_URL is required when using Azure Key Vault.");
  }

  let parsedUrl: URL;
  try {
    parsedUrl = new URL(vaultUrl);
  } catch {
    throw new Error("KEY_VAULT_URL must be a valid URL.");
  }

  if (parsedUrl.protocol !== "https:") {
    throw new Error("KEY_VAULT_URL must use HTTPS.");
  }

  const credential = environment.AZURE_CLIENT_ID
    ? new ManagedIdentityCredential({ clientId: environment.AZURE_CLIENT_ID })
    : new ManagedIdentityCredential();

  return new SecretClient(parsedUrl.toString(), credential);
}
