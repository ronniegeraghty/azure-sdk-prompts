import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import {
  BlobServiceClient,
  ContainerClient,
} from "@azure/storage-blob";

export interface AzureConnections {
  readonly credential: ManagedIdentityCredential;
  readonly keyClient: KeyClient;
  readonly containerClient: ContainerClient;
  readonly keyName: string;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set`);
  }

  return value;
}

export function createAzureConnections(): AzureConnections {
  const blobEndpoint = requiredEnvironmentVariable(
    "AZURE_STORAGE_BLOB_ENDPOINT",
  );
  const vaultEndpoint = requiredEnvironmentVariable("AZURE_KEY_VAULT_ENDPOINT");
  const containerName = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const keyName = requiredEnvironmentVariable("AZURE_KEY_VAULT_KEY_NAME");

  // One credential instance is deliberately shared by every Azure SDK client.
  const credential = new ManagedIdentityCredential();
  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);

  return {
    credential,
    keyClient: new KeyClient(vaultEndpoint, credential),
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyName,
  };
}
