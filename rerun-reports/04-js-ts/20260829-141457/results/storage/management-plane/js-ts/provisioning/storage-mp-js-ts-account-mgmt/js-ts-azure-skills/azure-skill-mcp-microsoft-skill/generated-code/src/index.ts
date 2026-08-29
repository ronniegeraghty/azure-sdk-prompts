import { StorageManagementClient } from "@azure/arm-storage";
import { DefaultAzureCredential } from "@azure/identity";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

async function main(): Promise<void> {
  const subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
  const resourceGroupName = requireEnvironmentVariable("AZURE_RESOURCE_GROUP");
  const accountName = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  let accountCreated = false;

  try {
    console.log(`Creating storage account "${accountName}"...`);
    const createPoller = client.storageAccounts.create(
      resourceGroupName,
      accountName,
      {
        location: "eastus",
        kind: "StorageV2",
        sku: { name: "Standard_LRS" },
      },
    );
    const account = await createPoller.pollUntilDone();
    accountCreated = true;
    console.log(`Created: ${account.id}`);

    console.log(`Storage accounts in resource group "${resourceGroupName}":`);
    for await (const storageAccount of client.storageAccounts.listByResourceGroup(
      resourceGroupName,
    )) {
      console.log(`- ${storageAccount.name} (${storageAccount.location})`);
    }

    const properties = await client.storageAccounts.getProperties(
      resourceGroupName,
      accountName,
    );
    console.log("Created account properties:", {
      id: properties.id,
      name: properties.name,
      location: properties.location,
      provisioningState: properties.provisioningState,
      primaryEndpoints: properties.primaryEndpoints,
    });

    await client.blobServices.setServiceProperties(
      resourceGroupName,
      accountName,
      { isVersioningEnabled: true },
    );
    console.log("Blob versioning enabled.");
  } finally {
    if (accountCreated) {
      console.log(`Deleting storage account "${accountName}"...`);
      await client.storageAccounts.delete(resourceGroupName, accountName);
      console.log("Storage account deleted.");
    }
  }
}

main().catch((error: unknown) => {
  console.error("Storage account management failed:", error);
  process.exitCode = 1;
});
