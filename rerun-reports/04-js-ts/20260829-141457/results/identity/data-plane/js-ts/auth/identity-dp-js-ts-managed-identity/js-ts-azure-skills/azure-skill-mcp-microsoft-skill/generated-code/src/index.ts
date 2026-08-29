import {
  AzureCliCredential,
  ChainedTokenCredential,
  CredentialUnavailableError,
  ManagedIdentityCredential,
} from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const storageEndpoint = process.env.AZURE_STORAGE_BLOB_ENDPOINT;
const userAssignedClientId =
  process.env.AZURE_USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID;
const managedIdentityType = process.env.MANAGED_IDENTITY_TYPE ?? "system";

// No options selects the system-assigned managed identity of the Azure host.
const systemAssignedCredential = new ManagedIdentityCredential();

// A user-assigned identity is selected explicitly by its client ID.
const userAssignedCredential = userAssignedClientId
  ? new ManagedIdentityCredential({ clientId: userAssignedClientId })
  : undefined;

function selectManagedIdentityCredential(): ManagedIdentityCredential {
  if (managedIdentityType === "system") {
    return systemAssignedCredential;
  }

  if (managedIdentityType === "user") {
    if (!userAssignedCredential) {
      throw new Error(
        "Set AZURE_USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID when " +
          "MANAGED_IDENTITY_TYPE=user.",
      );
    }

    return userAssignedCredential;
  }

  throw new Error('MANAGED_IDENTITY_TYPE must be either "system" or "user".');
}

function containsOnlyUnavailableCredentials(error: unknown): boolean {
  if (error instanceof CredentialUnavailableError) {
    return true;
  }

  if (
    typeof error === "object" &&
    error !== null &&
    "errors" in error &&
    Array.isArray(error.errors)
  ) {
    return (
      error.errors.length > 0 &&
      error.errors.every(containsOnlyUnavailableCredentials)
    );
  }

  return false;
}

async function listContainers(): Promise<void> {
  if (!storageEndpoint) {
    throw new Error(
      "Set AZURE_STORAGE_BLOB_ENDPOINT to an HTTPS Blob service endpoint, " +
        "for example https://<account>.blob.core.windows.net.",
    );
  }

  const credential = new ChainedTokenCredential(
    selectManagedIdentityCredential(),
    new AzureCliCredential(),
  );
  const blobServiceClient = new BlobServiceClient(storageEndpoint, credential);

  console.log(`Listing containers from ${storageEndpoint}`);
  let count = 0;

  for await (const container of blobServiceClient.listContainers()) {
    console.log(container.name);
    count += 1;
  }

  console.log(`Found ${count} container(s).`);
}

async function main(): Promise<void> {
  try {
    await listContainers();
  } catch (error: unknown) {
    if (containsOnlyUnavailableCredentials(error)) {
      console.error(
        "Managed identity is unavailable because this process is not running " +
          "on an Azure host, and Azure CLI authentication is unavailable. " +
          "For local development, install Azure CLI and run `az login`.",
      );
      process.exitCode = 1;
      return;
    }

    throw error;
  }
}

await main();
