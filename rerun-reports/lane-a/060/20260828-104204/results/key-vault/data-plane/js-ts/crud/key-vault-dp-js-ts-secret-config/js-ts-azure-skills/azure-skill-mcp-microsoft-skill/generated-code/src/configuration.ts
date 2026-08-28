import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";
import { CachedSecretProvider } from "./secret-cache.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";
import { SecretRotationHelper } from "./secret-rotation.js";

export interface KeyVaultConfiguration {
  vaultUrl: string;
  expiryWarningDays: number;
  managedIdentityClientId?: string;
}

export function loadKeyVaultConfiguration(
  environment: NodeJS.ProcessEnv = process.env,
): KeyVaultConfiguration {
  const vaultUrl = environment.KEY_VAULT_URL;
  if (vaultUrl === undefined || vaultUrl.trim().length === 0) {
    throw new Error("KEY_VAULT_URL must be set, for example https://my-vault.vault.azure.net.");
  }

  const parsedVaultUrl = new URL(vaultUrl);
  if (parsedVaultUrl.protocol !== "https:") {
    throw new Error("KEY_VAULT_URL must use HTTPS.");
  }

  const expiryWarningDays = Number(environment.SECRET_EXPIRY_WARNING_DAYS ?? "7");
  if (!Number.isFinite(expiryWarningDays) || expiryWarningDays < 0) {
    throw new Error("SECRET_EXPIRY_WARNING_DAYS must be a non-negative number.");
  }

  return {
    vaultUrl: parsedVaultUrl.toString(),
    expiryWarningDays,
    ...(environment.AZURE_CLIENT_ID === undefined
      ? {}
      : { managedIdentityClientId: environment.AZURE_CLIENT_ID }),
  };
}

export function createKeyVaultServices(configuration = loadKeyVaultConfiguration()) {
  const credential =
    configuration.managedIdentityClientId === undefined
      ? new ManagedIdentityCredential()
      : new ManagedIdentityCredential(configuration.managedIdentityClientId);
  const client = new SecretClient(configuration.vaultUrl, credential);
  const provider = new KeyVaultSecretProvider(client);

  return {
    client,
    provider,
    cache: new CachedSecretProvider(provider, configuration.expiryWarningDays),
    rotation: new SecretRotationHelper(client),
  };
}
