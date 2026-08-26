import {
  AzureCliCredential,
  ChainedTokenCredential,
  CredentialUnavailableError,
  ManagedIdentityCredential,
} from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const storageAccountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;
const userAssignedClientId = process.env.AZURE_CLIENT_ID;

if (!storageAccountName) {
  throw new Error("Set AZURE_STORAGE_ACCOUNT_NAME to the target storage account name.");
}

// No options selects the system-assigned identity.
const systemAssignedCredential = new ManagedIdentityCredential();

// Supplying a client ID selects a user-assigned identity.
const userAssignedCredential = userAssignedClientId
  ? new ManagedIdentityCredential({ clientId: userAssignedClientId })
  : undefined;

const managedIdentityCredential =
  userAssignedCredential ?? systemAssignedCredential;

// In Azure, managed identity is used. Locally, the chain falls back to `az login`.
const credential = new ChainedTokenCredential(
  managedIdentityCredential,
  new AzureCliCredential(),
);

async function reportManagedIdentityAvailability(): Promise<void> {
  try {
    await managedIdentityCredential.getToken(
      "https://storage.azure.com/.default",
    );
    console.log(
      userAssignedCredential
        ? "Using the user-assigned managed identity."
        : "Using the system-assigned managed identity.",
    );
  } catch (error: unknown) {
    if (error instanceof CredentialUnavailableError) {
      console.log(
        "Managed identity is unavailable; falling back to Azure CLI credentials.",
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

await main();
