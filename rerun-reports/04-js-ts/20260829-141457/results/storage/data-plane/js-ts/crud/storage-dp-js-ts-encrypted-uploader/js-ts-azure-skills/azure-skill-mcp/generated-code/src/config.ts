import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import { BlobServiceClient, type ContainerClient } from "@azure/storage-blob";

import { ConfigurationError } from "./errors.js";

export interface AppConfig {
  blobName: string;
  containerClient: ContainerClient;
  credential: ManagedIdentityCredential;
  keyClient: KeyClient;
  keyName: string;
  keyVaultUrl: string;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new ConfigurationError(`Environment variable ${name} is required.`);
  }

  return value;
}

function secureEndpoint(name: string): string {
  const value = requiredEnvironmentVariable(name).replace(/\/+$/, "");
  let endpoint: URL;

  try {
    endpoint = new URL(value);
  } catch (error) {
    throw new ConfigurationError(
      `${name} must be a valid absolute URL (${String(error)}).`,
    );
  }

  if (endpoint.protocol !== "https:") {
    throw new ConfigurationError(`${name} must use HTTPS.`);
  }

  return endpoint.toString().replace(/\/+$/, "");
}

export function createAppConfig(): AppConfig {
  const blobEndpoint = secureEndpoint("AZURE_STORAGE_BLOB_ENDPOINT");
  const keyVaultUrl = secureEndpoint("AZURE_KEY_VAULT_URL");
  const containerName = requiredEnvironmentVariable("AZURE_STORAGE_CONTAINER");
  const keyName = requiredEnvironmentVariable("AZURE_KEY_NAME");
  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential(managedIdentityClientId)
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);
  const keyClient = new KeyClient(keyVaultUrl, credential);

  return {
    blobName:
      process.env.AZURE_BLOB_NAME?.trim() || "envelope-encryption-demo.txt",
    containerClient: blobServiceClient.getContainerClient(containerName),
    credential,
    keyClient,
    keyName,
    keyVaultUrl,
  };
}
