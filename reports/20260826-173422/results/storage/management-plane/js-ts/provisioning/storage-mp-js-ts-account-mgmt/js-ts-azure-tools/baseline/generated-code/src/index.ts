import { StorageManagementClient } from "@azure/arm-storage";
import { DefaultAzureCredential } from "@azure/identity";

const subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
const resourceGroupName = requireEnvironmentVariable("AZURE_RESOURCE_GROUP");
const storageAccountName = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

validateStorageAccountName(storageAccountName);

const credential = new DefaultAzureCredential();
const client = new StorageManagementClient(credential, subscriptionId);

async function manageStorageAccount(): Promise<void> {
  let accountCreated = false;

  try {
    console.log(`Creating storage account "${storageAccountName}"...`);
    const createdAccount = await client.storageAccounts.beginCreateAndWait(
      resourceGroupName,
      storageAccountName,
      {
        location: "eastus",
        kind: "StorageV2",
        sku: {
          name: "Standard_LRS",
        },
      },
    );
    accountCreated = true;
    console.log(`Created: ${createdAccount.id}`);

    console.log(`Storage accounts in resource group "${resourceGroupName}":`);
    for await (const account of client.storageAccounts.listByResourceGroup(
      resourceGroupName,
    )) {
      console.log(`- ${account.name} (${account.location})`);
    }

    const account = await client.storageAccounts.getProperties(
      resourceGroupName,
      storageAccountName,
    );
    console.log("Created account properties:", {
      id: account.id,
      name: account.name,
      location: account.location,
      provisioningState: account.provisioningState,
      primaryEndpoints: account.primaryEndpoints,
    });

    // Blob versioning is configured on the account's default Blob Service.
    await client.blobServices.setServiceProperties(
      resourceGroupName,
      storageAccountName,
      {
        isVersioningEnabled: true,
      },
    );
    console.log("Blob versioning enabled.");
  } finally {
    if (accountCreated) {
      console.log(`Deleting storage account "${storageAccountName}"...`);
      await client.storageAccounts.delete(
        resourceGroupName,
        storageAccountName,
      );
      console.log("Storage account deleted.");
    }
  }
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function validateStorageAccountName(name: string): void {
  if (!/^[a-z0-9]{3,24}$/.test(name)) {
    throw new Error(
      "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters and numbers.",
    );
  }
}

await manageStorageAccount();
