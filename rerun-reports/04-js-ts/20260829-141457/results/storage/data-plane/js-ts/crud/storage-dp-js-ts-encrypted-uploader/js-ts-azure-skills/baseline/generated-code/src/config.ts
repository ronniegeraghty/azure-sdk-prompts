import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import {
  BlobServiceClient,
  ContainerClient,
} from "@azure/storage-blob";

export interface AzureConfiguration {
  credential: ManagedIdentityCredential;
  containerClient: ContainerClient;
  keyClient: KeyClient;
  keyName: string;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

export function createAzureConfiguration(): AzureConfiguration {
  const blobEndpoint = requiredEnvironmentVariable(
    "AZURE_STORAGE_BLOB_ENDPOINT",
  );
  const containerName = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const keyVaultUrl = requiredEnvironmentVariable("AZURE_KEY_VAULT_URL");
  const keyName = requiredEnvironmentVariable("AZURE_KEY_NAME");
  const managedIdentityClientId = process.env.AZURE_CLIENT_ID;

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential(managedIdentityClientId)
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);
  const containerClient = blobServiceClient.getContainerClient(containerName);
  const keyClient = new KeyClient(keyVaultUrl, credential);

  return {
    credential,
    containerClient,
    keyClient,
    keyName,
  };
}
