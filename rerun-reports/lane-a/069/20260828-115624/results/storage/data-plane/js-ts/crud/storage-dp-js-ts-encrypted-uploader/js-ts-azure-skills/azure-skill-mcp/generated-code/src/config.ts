import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import { BlobServiceClient, ContainerClient } from "@azure/storage-blob";

export interface AzureConnections {
  credential: ManagedIdentityCredential;
  containerClient: ContainerClient;
  keyClient: KeyClient;
  keyName: string;
  keyVersion?: string;
}

function requireEnvironment(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function requireHttpsEndpoint(name: string): string {
  const value = requireEnvironment(name);

  let url: URL;
  try {
    url = new URL(value);
  } catch (error) {
    throw new Error(`${name} must be a valid URL.`, { cause: error });
  }

  if (url.protocol !== "https:") {
    throw new Error(`${name} must use HTTPS.`);
  }

  return url.toString().replace(/\/$/, "");
}

export function createAzureConnections(): AzureConnections {
  const blobEndpoint = requireHttpsEndpoint("AZURE_STORAGE_BLOB_ENDPOINT");
  const containerName = requireEnvironment("AZURE_STORAGE_CONTAINER");
  const vaultEndpoint = requireHttpsEndpoint("AZURE_KEY_VAULT_ENDPOINT");
  const keyName = requireEnvironment("AZURE_KEY_VAULT_KEY_NAME");
  const keyVersion = process.env.AZURE_KEY_VAULT_KEY_VERSION?.trim() || undefined;
  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);
  const keyClient = new KeyClient(vaultEndpoint, credential);

  return {
    credential,
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyClient,
    keyName,
    ...(keyVersion ? { keyVersion } : {}),
  };
}
