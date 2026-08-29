import { DefaultAzureCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

import { SecretCache } from "./secret-cache.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";

export interface KeyVaultConfiguration {
  client: SecretClient;
  provider: KeyVaultSecretProvider;
  cache: SecretCache;
}

export function createKeyVaultConfiguration(
  environment: NodeJS.ProcessEnv = process.env,
  warningWindowMs?: number,
): KeyVaultConfiguration {
  const vaultUrl = environment["AZURE_KEY_VAULT_URL"];
  if (vaultUrl === undefined || vaultUrl.trim() === "") {
    throw new Error("AZURE_KEY_VAULT_URL must be set");
  }

  const credential = new DefaultAzureCredential();
  const client = new SecretClient(vaultUrl, credential);
  const provider = new KeyVaultSecretProvider(client);
  const cache =
    warningWindowMs === undefined
      ? new SecretCache(provider)
      : new SecretCache(provider, warningWindowMs);

  return { client, provider, cache };
}
