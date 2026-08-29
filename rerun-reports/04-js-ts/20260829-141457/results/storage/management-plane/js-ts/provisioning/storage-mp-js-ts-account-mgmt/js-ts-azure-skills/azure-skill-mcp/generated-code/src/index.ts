import { DefaultAzureCredential } from "@azure/identity";
import { StorageManagementClient } from "@azure/arm-storage";

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
  const storageAccountName = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  let accountCreated = false;

  try {
    console.log(`Creating storage account "${storageAccountName}"...`);
    const createPoller = client.storageAccounts.create(
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
    const createdAccount = await createPoller.pollUntilDone();
    accountCreated = true;
    console.log(`Created: ${createdAccount.id}`);

    console.log(`Storage accounts in resource group "${resourceGroupName}":`);
    for await (const account of client.storageAccounts.listByResourceGroup(
      resourceGroupName,
    )) {
      console.log(`- ${account.name} (${account.location})`);
    }

    const accountProperties = await client.storageAccounts.getProperties(
      resourceGroupName,
      storageAccountName,
    );
    console.log("Created account properties:", {
      id: accountProperties.id,
      name: accountProperties.name,
      location: accountProperties.location,
      provisioningState: accountProperties.provisioningState,
      primaryEndpoints: accountProperties.primaryEndpoints,
    });

    // Blob versioning is configured on the account's Blob Service resource.
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

main().catch((error: unknown) => {
  console.error("Storage account management failed:", error);
  process.exitCode = 1;
});
