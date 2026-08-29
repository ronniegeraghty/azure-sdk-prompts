import type { TokenCredential } from "@azure/core-auth";
import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import {
  BlobServiceClient,
  type ContainerClient,
} from "@azure/storage-blob";

export interface AppConfig {
  readonly credential: TokenCredential;
  readonly keyClient: KeyClient;
  readonly containerClient: ContainerClient;
  readonly keyName: string;
  readonly blobName: string;
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

export function createAppConfig(): AppConfig {
  const storageEndpoint = requireHttpsEndpoint(
    "AZURE_STORAGE_BLOB_ENDPOINT",
  );
  const vaultEndpoint = requireHttpsEndpoint("AZURE_KEY_VAULT_ENDPOINT");
  const keyName = requireEnvironmentVariable("AZURE_KEY_VAULT_KEY_NAME");
  const containerName = requireEnvironmentVariable("AZURE_STORAGE_CONTAINER");
  const blobName =
    process.env.AZURE_STORAGE_BLOB_NAME?.trim() || "encrypted-demo.txt";
  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();

  const credential: TokenCredential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(
    storageEndpoint,
    credential,
  );

  return {
    credential,
    keyClient: new KeyClient(vaultEndpoint, credential),
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyName,
    blobName,
  };
}
