import { ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

export function createKeyVaultSecretClient(
  environment: NodeJS.ProcessEnv = process.env,
): SecretClient {
  const vaultUrl = environment.KEY_VAULT_URL;
  if (vaultUrl === undefined || vaultUrl.trim() === "") {
    throw new Error("KEY_VAULT_URL is required");
  }

  const parsedUrl = new URL(vaultUrl);
  if (parsedUrl.protocol !== "https:") {
    throw new Error("KEY_VAULT_URL must use HTTPS");
  }

  const clientId = environment.AZURE_CLIENT_ID?.trim();
  const credential =
    clientId === undefined || clientId === ""
      ? new ManagedIdentityCredential()
      : new ManagedIdentityCredential({ clientId });

  return new SecretClient(parsedUrl.toString(), credential);
}
