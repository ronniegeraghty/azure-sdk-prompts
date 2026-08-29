import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";
import { SecretCache } from "./secret-cache.js";
import { KeyVaultSecretProvider } from "./secret-provider.js";

const DAY_MS = 24 * 60 * 60 * 1000;

function requireHttpsUrl(value: string, variableName: string): string {
  let url: URL;
  try {
    url = new URL(value);
  } catch {
    throw new Error(`${variableName} must be a valid URL.`);
  }

  if (url.protocol !== "https:") {
    throw new Error(`${variableName} must use HTTPS.`);
  }

  if (url.username || url.password) {
    throw new Error(`${variableName} must not contain credentials.`);
  }

  return url.toString().replace(/\/$/, "");
}

function readNonNegativeNumber(value: string | undefined, fallback: number, name: string): number {
  if (value === undefined || value.trim() === "") {
    return fallback;
  }

  const number = Number(value);
  if (!Number.isFinite(number) || number < 0) {
    throw new Error(`${name} must be a non-negative number.`);
  }
  return number;
}

export interface ApplicationConfiguration {
  client: SecretClient;
  provider: KeyVaultSecretProvider;
  cache: SecretCache;
  expiryWarningWindowMs: number;
  autoRefreshIntervalMs: number;
}

export function createApplicationConfiguration(
  env: NodeJS.ProcessEnv = process.env
): ApplicationConfiguration {
  const rawVaultUrl = env.KEY_VAULT_URL;
  if (!rawVaultUrl) {
    throw new Error("KEY_VAULT_URL is required.");
  }

  const vaultUrl = requireHttpsUrl(rawVaultUrl, "KEY_VAULT_URL");
  const credential = env.AZURE_CLIENT_ID
    ? new ManagedIdentityCredential(env.AZURE_CLIENT_ID)
    : new ManagedIdentityCredential();
  const client = new SecretClient(vaultUrl, credential, {
    retryOptions: {
      maxRetries: 5,
      retryDelayInMs: 1_000,
      maxRetryDelayInMs: 10_000
    }
  });
  const provider = new KeyVaultSecretProvider(client);
  const expiryWarningWindowMs =
    readNonNegativeNumber(env.EXPIRY_WARNING_DAYS, 7, "EXPIRY_WARNING_DAYS") * DAY_MS;
  const autoRefreshIntervalMs =
    readNonNegativeNumber(
      env.AUTO_REFRESH_INTERVAL_MINUTES,
      60,
      "AUTO_REFRESH_INTERVAL_MINUTES"
    ) * 60 * 1000;

  if (autoRefreshIntervalMs === 0) {
    throw new Error("AUTO_REFRESH_INTERVAL_MINUTES must be greater than zero.");
  }

  return {
    client,
    provider,
    cache: new SecretCache(provider, expiryWarningWindowMs),
    expiryWarningWindowMs,
    autoRefreshIntervalMs
  };
}
