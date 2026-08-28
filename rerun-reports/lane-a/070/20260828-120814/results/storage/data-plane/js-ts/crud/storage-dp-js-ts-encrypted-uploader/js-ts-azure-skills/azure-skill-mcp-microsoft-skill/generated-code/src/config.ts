import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import {
  BlobServiceClient,
  type ContainerClient,
} from "@azure/storage-blob";

export interface AzureConnections {
  readonly credential: ManagedIdentityCredential;
  readonly containerClient: ContainerClient;
  readonly keyClient: KeyClient;
  readonly keyName: string;
  readonly vaultUrl: string;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }

  return value;
}

function validateHttpsEndpoint(name: string, value: string): string {
  let endpoint: URL;
  try {
    endpoint = new URL(value);
  } catch {
    throw new Error(`${name} must be a valid absolute URL.`);
  }

  if (endpoint.protocol !== "https:") {
    throw new Error(`${name} must use HTTPS.`);
  }

  endpoint.pathname = endpoint.pathname.replace(/\/+$/, "");
  endpoint.search = "";
  endpoint.hash = "";
  return endpoint.toString().replace(/\/$/, "");
}

export function buildAzureConnections(): AzureConnections {
  const blobEndpoint = validateHttpsEndpoint(
    "AZURE_STORAGE_BLOB_ENDPOINT",
    requiredEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT"),
  );
  const vaultUrl = validateHttpsEndpoint(
    "AZURE_KEY_VAULT_URL",
    requiredEnvironmentVariable("AZURE_KEY_VAULT_URL"),
  );
  const containerName = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const keyName = requiredEnvironmentVariable("AZURE_KEY_VAULT_KEY_NAME");
  const managedIdentityClientId =
    process.env.AZURE_MANAGED_IDENTITY_CLIENT_ID?.trim();

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);

  return {
    credential,
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyClient: new KeyClient(vaultUrl, credential),
    keyName,
    vaultUrl,
  };
}
