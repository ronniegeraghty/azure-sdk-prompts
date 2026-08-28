import { StorageManagementClient } from "@azure/arm-storage";
import { DefaultAzureCredential } from "@azure/identity";

function getRequiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

async function main(): Promise<void> {
  const subscriptionId = getRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
  const resourceGroupName = getRequiredEnvironmentVariable(
    "AZURE_RESOURCE_GROUP_NAME",
  );
  const storageAccountName = getRequiredEnvironmentVariable(
    "AZURE_STORAGE_ACCOUNT_NAME",
  );

  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  let accountCreated = false;

  try {
    console.log(`Creating storage account "${storageAccountName}"...`);
    const createdAccount = await client.storageAccounts.beginCreateAndWait(
      resourceGroupName,
      storageAccountName,
      {
        location: "eastus",
        kind: "StorageV2",
        sku: { name: "Standard_LRS" },
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
      kind: account.kind,
      sku: account.sku?.name,
      provisioningState: account.provisioningState,
      primaryEndpoints: account.primaryEndpoints,
    });

    await client.blobServices.setServiceProperties(
      resourceGroupName,
      storageAccountName,
      { isVersioningEnabled: true },
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
