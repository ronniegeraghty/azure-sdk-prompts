import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

export const KEY_VAULT_URL_ENV = "KEY_VAULT_URL";
export const MANAGED_IDENTITY_CLIENT_ID_ENV =
  "AZURE_MANAGED_IDENTITY_CLIENT_ID";

export function createKeyVaultSecretClient(
  environment: NodeJS.ProcessEnv = process.env,
): SecretClient {
  const vaultUrl = requireVaultUrl(environment[KEY_VAULT_URL_ENV]);
  const clientId = environment[MANAGED_IDENTITY_CLIENT_ID_ENV];
  const credential =
    clientId === undefined || clientId.trim() === ""
      ? new ManagedIdentityCredential()
      : new ManagedIdentityCredential({ clientId });

  return new SecretClient(vaultUrl, credential, {
    retryOptions: {
      maxRetries: 4,
      retryDelayInMs: 800,
      maxRetryDelayInMs: 8_000,
    },
  });
}

function requireVaultUrl(value: string | undefined): string {
  if (value === undefined || value.trim() === "") {
    throw new Error(`${KEY_VAULT_URL_ENV} must be set`);
  }

  const url = new URL(value);
  if (
    url.protocol !== "https:" ||
    url.username !== "" ||
    url.password !== "" ||
    url.pathname !== "/" ||
    url.search !== "" ||
    url.hash !== ""
  ) {
    throw new Error(
      `${KEY_VAULT_URL_ENV} must be an HTTPS vault origin without credentials, path, query, or fragment`,
    );
  }

  return url.origin;
}
