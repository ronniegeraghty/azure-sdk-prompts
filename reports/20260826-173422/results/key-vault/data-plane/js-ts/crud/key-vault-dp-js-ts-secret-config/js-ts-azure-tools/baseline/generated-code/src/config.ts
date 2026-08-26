import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

import { SecretCache, type RequiredSecrets } from "./secret-cache.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";
import { SecretRotator } from "./secret-rotator.js";

export interface KeyVaultConfiguration {
  provider: KeyVaultSecretProvider;
  cache: SecretCache;
  rotator: SecretRotator;
}

export function createKeyVaultConfiguration(
  requiredSecrets: RequiredSecrets,
  expiryWarningWindowMs: number,
): KeyVaultConfiguration {
  const vaultUrl = process.env["KEY_VAULT_URL"];
  if (vaultUrl === undefined || vaultUrl.trim() === "") {
    throw new Error("KEY_VAULT_URL must contain the Azure Key Vault URL.");
  }

  const credential = new ManagedIdentityCredential();
  const client = new SecretClient(vaultUrl, credential);
  const provider = new KeyVaultSecretProvider(client);

  return {
    provider,
    cache: new SecretCache(provider, requiredSecrets, expiryWarningWindowMs),
    rotator: new SecretRotator(client),
  };
}
