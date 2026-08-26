import { ManagedIdentityCredential } from "@azure/identity";
import { KeyClient } from "@azure/keyvault-keys";
import {
  BlobServiceClient,
  type ContainerClient,
} from "@azure/storage-blob";

export interface AppConfiguration {
  credential: ManagedIdentityCredential;
  containerClient: ContainerClient;
  keyClient: KeyClient;
  keyName: string;
}

export function createAppConfiguration(
  environment: NodeJS.ProcessEnv = process.env,
): AppConfiguration {
  const blobEndpoint = requireUrl(
    environment,
    "AZURE_STORAGE_BLOB_ENDPOINT",
  );
  const containerName = requireEnvironmentVariable(
    environment,
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const keyVaultUrl = requireUrl(environment, "AZURE_KEY_VAULT_URL");
  const keyName = requireEnvironmentVariable(
    environment,
    "AZURE_KEY_VAULT_KEY_NAME",
  );
  const managedIdentityClientId =
    environment["AZURE_MANAGED_IDENTITY_CLIENT_ID"]?.trim();

  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);

  return {
    credential,
    containerClient: blobServiceClient.getContainerClient(containerName),
    keyClient: new KeyClient(keyVaultUrl, credential),
    keyName,
  };
}

function requireEnvironmentVariable(
  environment: NodeJS.ProcessEnv,
  name: string,
): string {
  const value = environment[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function requireUrl(
  environment: NodeJS.ProcessEnv,
  name: string,
): string {
  const value = requireEnvironmentVariable(environment, name);

  let url: URL;
  try {
    url = new URL(value);
  } catch (error) {
    throw new Error(`${name} must be a valid HTTPS URL.`, { cause: error });
  }

  if (url.protocol !== "https:") {
    throw new Error(`${name} must use HTTPS.`);
  }

  return url.toString().replace(/\/$/, "");
}
