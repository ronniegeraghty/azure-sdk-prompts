import {
  AzureCliCredential,
  ChainedTokenCredential,
  CredentialUnavailableError,
  ManagedIdentityCredential,
} from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

const userAssignedClientId =
  process.env.AZURE_MANAGED_IDENTITY_CLIENT_ID ??
  "00000000-0000-0000-0000-000000000000";

const systemAssignedCredential = new ManagedIdentityCredential();
const userAssignedCredential = new ManagedIdentityCredential(
  userAssignedClientId,
);

const managedIdentityCredential =
  process.env.AZURE_USE_USER_ASSIGNED_IDENTITY === "true"
    ? userAssignedCredential
    : systemAssignedCredential;

const credential = new ChainedTokenCredential(
  managedIdentityCredential,
  new AzureCliCredential(),
);

function includesCredentialUnavailableError(error: unknown): boolean {
  if (error instanceof CredentialUnavailableError) {
    return true;
  }

  if (
    typeof error === "object" &&
    error !== null &&
    "errors" in error &&
    Array.isArray(error.errors)
  ) {
    return error.errors.some(includesCredentialUnavailableError);
  }

  return false;
}

async function main(): Promise<void> {
  const storageAccountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;
  if (!storageAccountName) {
    throw new Error("Set AZURE_STORAGE_ACCOUNT_NAME before running.");
  }

  const blobServiceClient = new BlobServiceClient(
    `https://${storageAccountName}.blob.core.windows.net`,
    credential,
  );

  try {
    const page = await blobServiceClient
      .listContainers()
      .byPage({ maxPageSize: 10 })
      .next();

    if (page.done) {
      console.log("No containers found.");
      return;
    }

    for (const container of page.value.containerItems) {
      console.log(container.name);
    }
  } catch (error: unknown) {
    if (includesCredentialUnavailableError(error)) {
      console.error(
        "Managed Identity is unavailable. When running locally, sign in with Azure CLI so AzureCliCredential can be used.",
      );
      return;
    }

    throw error;
  }
}

await main();
