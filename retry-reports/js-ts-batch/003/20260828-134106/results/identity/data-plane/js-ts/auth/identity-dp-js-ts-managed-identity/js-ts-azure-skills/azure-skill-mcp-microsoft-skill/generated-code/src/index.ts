import {
  AzureCliCredential,
  ChainedTokenCredential,
  CredentialUnavailableError,
  ManagedIdentityCredential,
} from "@azure/identity";
import { BlobServiceClient } from "@azure/storage-blob";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Set the ${name} environment variable before running this program.`);
  }

  return value;
}

function containsCredentialUnavailableError(error: unknown): boolean {
  if (error instanceof CredentialUnavailableError) {
    return true;
  }

  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { cause?: unknown; errors?: unknown[] };
  return (
    (Array.isArray(candidate.errors) &&
      candidate.errors.some(containsCredentialUnavailableError)) ||
    containsCredentialUnavailableError(candidate.cause)
  );
}

async function main(): Promise<void> {
  const storageAccountUrl = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_URL");
  const userAssignedClientId = requireEnvironmentVariable(
    "AZURE_USER_ASSIGNED_CLIENT_ID",
  );

  const systemAssignedCredential = new ManagedIdentityCredential();
  const userAssignedCredential = new ManagedIdentityCredential({
    clientId: userAssignedClientId,
  });

  const managedIdentityCredential =
    process.env.USE_USER_ASSIGNED_IDENTITY === "true"
      ? userAssignedCredential
      : systemAssignedCredential;

  const credential = new ChainedTokenCredential(
    managedIdentityCredential,
    new AzureCliCredential(),
  );

  const blobServiceClient = new BlobServiceClient(
    storageAccountUrl,
    credential,
  );

  try {
    console.log("Containers:");
    for await (const container of blobServiceClient.listContainers()) {
      console.log(`- ${container.name}`);
    }
  } catch (error: unknown) {
    if (containsCredentialUnavailableError(error)) {
      console.error(
        "Managed Identity is unavailable. When running locally, install the Azure CLI and run 'az login'.",
      );
      process.exitCode = 1;
      return;
    }

    throw error;
  }
}

await main();
