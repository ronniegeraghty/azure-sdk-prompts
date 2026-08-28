import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";
import {
  CachedSecretProvider,
  type CachedSecretProviderOptions,
} from "./cached-secret-provider.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";

export interface KeyVaultConfiguration {
  client: SecretClient;
  provider: KeyVaultSecretProvider;
  cache: CachedSecretProvider;
}

export function createKeyVaultConfiguration(
  cacheOptions: CachedSecretProviderOptions = {},
): KeyVaultConfiguration {
  const vaultUrl = process.env.KEY_VAULT_URL;
  if (vaultUrl === undefined || vaultUrl.trim() === "") {
    throw new Error("KEY_VAULT_URL must contain the Azure Key Vault URL.");
  }

  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();
  const credential =
    managedIdentityClientId === undefined || managedIdentityClientId === ""
      ? new ManagedIdentityCredential()
      : new ManagedIdentityCredential(managedIdentityClientId);
  const client = new SecretClient(vaultUrl, credential);
  const provider = new KeyVaultSecretProvider(client);

  return {
    client,
    provider,
    cache: new CachedSecretProvider(provider, cacheOptions),
  };
}
