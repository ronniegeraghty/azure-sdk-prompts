import { ResourceManagementClient } from "@azure/arm-resources";
import { DefaultAzureCredential } from "@azure/identity";

const location = "eastus";
const resourceGroupName =
  process.env.AZURE_RESOURCE_GROUP_NAME ?? "typescript-sdk-example-rg";

async function manageResourceGroup(): Promise<void> {
  const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;
  if (!subscriptionId) {
    throw new Error("AZURE_SUBSCRIPTION_ID must be set.");
  }

  if (process.env.AZURE_EXECUTE !== "true") {
    console.log(
      `Dry run: would manage resource group "${resourceGroupName}" in ` +
        `"${location}" for subscription "${subscriptionId}".`,
    );
    console.log("Set AZURE_EXECUTE=true to perform the Azure operations.");
    return;
  }

  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);

  try {
    const created = await client.resourceGroups.createOrUpdate(
      resourceGroupName,
      { location },
    );
    console.log("Created resource group:", created);

    console.log("Resource groups in the subscription:");
    for await (const resourceGroup of client.resourceGroups.list()) {
      console.log(
        `- ${resourceGroup.name ?? "(unnamed)"} (${resourceGroup.location})`,
      );
    }

    const details = await client.resourceGroups.get(resourceGroupName);
    console.log("Created resource group details:", details);

    const updated = await client.resourceGroups.update(resourceGroupName, {
      tags: {
        ...details.tags,
        managedBy: "typescript-azure-sdk",
      },
    });
    console.log("Updated resource group:", updated);
  } finally {
    console.log(`Deleting resource group "${resourceGroupName}"...`);
    await client.resourceGroups.beginDeleteAndWait(resourceGroupName);
    console.log("Resource group deleted.");
  }
}

manageResourceGroup().catch((error: unknown) => {
  console.error("Resource group operation failed:", error);
  process.exitCode = 1;
});
