import { ResourceManagementClient } from "@azure/arm-resources";
import { DefaultAzureCredential } from "@azure/identity";

const resourceGroupName =
  process.env.AZURE_RESOURCE_GROUP_NAME ?? "typescript-sdk-resource-group-demo";

async function main(): Promise<void> {
  const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;
  if (!subscriptionId) {
    throw new Error("AZURE_SUBSCRIPTION_ID must be set.");
  }

  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  let resourceGroupCreated = false;

  try {
    const createdResourceGroup = await client.resourceGroups.createOrUpdate(
      resourceGroupName,
      { location: "eastus" },
    );
    resourceGroupCreated = true;
    console.log("Created resource group:", createdResourceGroup);

    console.log("Resource groups in the subscription:");
    for await (const resourceGroup of client.resourceGroups.list()) {
      console.log(`- ${resourceGroup.name} (${resourceGroup.location})`);
    }

    const resourceGroup = await client.resourceGroups.get(resourceGroupName);
    console.log("Resource group details:", resourceGroup);

    const updatedResourceGroup = await client.resourceGroups.update(
      resourceGroupName,
      {
        tags: {
          ...resourceGroup.tags,
          managedBy: "typescript-sdk",
        },
      },
    );
    console.log("Updated resource group:", updatedResourceGroup);
  } finally {
    if (resourceGroupCreated) {
      await client.resourceGroups.beginDeleteAndWait(resourceGroupName);
      console.log(`Deleted resource group: ${resourceGroupName}`);
    }
  }
}

main().catch((error: unknown) => {
  console.error("Resource group operation failed:", error);
  process.exitCode = 1;
});
