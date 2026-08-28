import {
  AzureCliCredential,
  ChainedTokenCredential,
  CredentialUnavailableError,
  ManagedIdentityCredential,
} from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const storageAccountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;
const userAssignedClientId =
  process.env.AZURE_USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID;

if (!storageAccountName) {
  throw new Error("AZURE_STORAGE_ACCOUNT_NAME must be set.");
}

if (!userAssignedClientId) {
  throw new Error(
    "AZURE_USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID must be set.",
  );
}

// An empty configuration selects the system-assigned identity.
const systemAssignedCredential = new ManagedIdentityCredential();

// A client ID selects a specific user-assigned managed identity.
const userAssignedCredential = new ManagedIdentityCredential({
  clientId: userAssignedClientId,
});

const managedIdentityCredential =
  process.env.MANAGED_IDENTITY_TYPE === "user-assigned"
    ? userAssignedCredential
    : systemAssignedCredential;

// Managed identity is attempted in Azure; `az login` is used as a local fallback.
const credential = new ChainedTokenCredential(
  managedIdentityCredential,
  new AzureCliCredential(),
);

async function reportManagedIdentityAvailability(): Promise<void> {
  try {
    await managedIdentityCredential.getToken(
      "https://storage.azure.com/.default",
    );
  } catch (error: unknown) {
    if (error instanceof CredentialUnavailableError) {
      console.info(
        "Managed identity is unavailable. Falling back to Azure CLI credentials.",
      );
      return;
    }

    throw error;
  }
}

async function main(): Promise<void> {
  await reportManagedIdentityAvailability();

  const serviceUrl = `https://${storageAccountName}.blob.core.windows.net`;
  const blobServiceClient = new BlobServiceClient(serviceUrl, credential);

  console.log(`Containers in ${storageAccountName}:`);
  for await (const container of blobServiceClient.listContainers()) {
    console.log(`- ${container.name}`);
  }
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Azure operation failed: ${message}`);
  process.exitCode = 1;
});
