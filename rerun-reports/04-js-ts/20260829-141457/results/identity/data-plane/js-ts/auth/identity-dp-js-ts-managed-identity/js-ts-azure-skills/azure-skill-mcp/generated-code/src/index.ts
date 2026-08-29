import {
  AggregateAuthenticationError,
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

  return (
    error instanceof AggregateAuthenticationError &&
    error.errors.some(containsCredentialUnavailableError)
  );
}

async function main(): Promise<void> {
  const accountName = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
  const userAssignedClientId = requireEnvironmentVariable(
    "AZURE_USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID",
  );

  // No options selects the system-assigned identity of the Azure host.
  const systemAssignedCredential = new ManagedIdentityCredential();
  const userAssignedCredential = new ManagedIdentityCredential({
    clientId: userAssignedClientId,
  });

  const managedIdentityCredential =
    process.env.USE_USER_ASSIGNED_IDENTITY === "true"
      ? userAssignedCredential
      : systemAssignedCredential;

  // Managed identity is deterministic in Azure; Azure CLI supports `az login` locally.
  const credential = new ChainedTokenCredential(
    managedIdentityCredential,
    new AzureCliCredential(),
  );

  const blobServiceClient = new BlobServiceClient(
    `https://${accountName}.blob.core.windows.net`,
    credential,
  );

  console.log(`Containers in storage account "${accountName}":`);
  for await (const container of blobServiceClient.listContainers()) {
    console.log(`- ${container.name}`);
  }
}

try {
  await main();
} catch (error: unknown) {
  if (containsCredentialUnavailableError(error)) {
    console.error(
      "Managed identity is unavailable. When running locally, install Azure CLI and run `az login`.",
    );
    process.exitCode = 1;
  } else {
    throw error;
  }
}
