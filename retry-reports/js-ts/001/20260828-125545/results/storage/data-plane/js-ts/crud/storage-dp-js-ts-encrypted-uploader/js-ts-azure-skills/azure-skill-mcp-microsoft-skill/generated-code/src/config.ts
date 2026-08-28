import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import {
  BlobServiceClient,
  type ContainerClient,
} from "@azure/storage-blob";

export interface AzureConfiguration {
  credential: ManagedIdentityCredential;
  keyClient: KeyClient;
  containerClient: ContainerClient;
  keyVaultUrl: string;
  keyName: string;
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function requireHttpsEndpoint(name: string): string {
  const value = requireEnvironmentVariable(name);

  let endpoint: URL;
  try {
    endpoint = new URL(value);
  } catch (error) {
    throw new Error(`${name} must be a valid URL.`, { cause: error });
  }

  if (endpoint.protocol !== "https:") {
    throw new Error(`${name} must use HTTPS.`);
  }

  return endpoint.toString().replace(/\/$/, "");
}

export function createAzureConfiguration(): AzureConfiguration {
  const blobEndpoint = requireHttpsEndpoint("AZURE_STORAGE_BLOB_ENDPOINT");
  const keyVaultUrl = requireHttpsEndpoint("AZURE_KEY_VAULT_URL");
  const containerName = requireEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const keyName = requireEnvironmentVariable("AZURE_KEY_VAULT_KEY_NAME");
  const managedIdentityClientId =
    process.env.AZURE_MANAGED_IDENTITY_CLIENT_ID?.trim();

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);

  return {
    credential,
    keyClient: new KeyClient(keyVaultUrl, credential),
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyVaultUrl,
    keyName,
  };
}
