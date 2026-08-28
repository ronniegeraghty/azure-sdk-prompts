import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import { BlobServiceClient, type ContainerClient } from "@azure/storage-blob";

export interface AzureClients {
  credential: ManagedIdentityCredential;
  keyClient: KeyClient;
  containerClient: ContainerClient;
  keyVaultUrl: string;
  keyName: string;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }

  return value;
}

function endpointEnvironmentVariable(name: string): string {
  const value = requiredEnvironmentVariable(name);

  let endpoint: URL;
  try {
    endpoint = new URL(value);
  } catch (error) {
    throw new Error(`${name} must be a valid absolute URL.`, { cause: error });
  }

  if (endpoint.protocol !== "https:") {
    throw new Error(`${name} must use HTTPS.`);
  }

  return endpoint.toString().replace(/\/$/, "");
}

export function createAzureClients(): AzureClients {
  const blobEndpoint = endpointEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT");
  const keyVaultUrl = endpointEnvironmentVariable("AZURE_KEY_VAULT_URL");
  const containerName = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const keyName = requiredEnvironmentVariable("AZURE_KEY_VAULT_KEY_NAME");
  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);
  const keyClient = new KeyClient(keyVaultUrl, credential);

  return {
    credential,
    keyClient,
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyVaultUrl,
    keyName,
  };
}
