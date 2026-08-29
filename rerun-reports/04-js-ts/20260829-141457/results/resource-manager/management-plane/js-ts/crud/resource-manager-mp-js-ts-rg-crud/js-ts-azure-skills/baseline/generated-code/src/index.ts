import { ResourceManagementClient } from "@azure/arm-resources";
import { DefaultAzureCredential } from "@azure/identity";

const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;
const resourceGroupName =
  process.env.AZURE_RESOURCE_GROUP_NAME ?? "sdk-resource-group-example";

if (!subscriptionId) {
  throw new Error("Set the AZURE_SUBSCRIPTION_ID environment variable.");
}

async function main(subscriptionId: string): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);

  console.log(`Creating resource group "${resourceGroupName}"...`);
  const created = await client.resourceGroups.createOrUpdate(
    resourceGroupName,
    { location: "eastus" },
  );
  console.log(`Created: ${created.id}`);

  console.log("\nResource groups in the subscription:");
  for await (const resourceGroup of client.resourceGroups.list()) {
    console.log(`- ${resourceGroup.name} (${resourceGroup.location})`);
  }

  const details = await client.resourceGroups.get(resourceGroupName);
  console.log("\nCreated resource group details:", details);

  const updated = await client.resourceGroups.update(resourceGroupName, {
    tags: {
      ...details.tags,
      managedBy: "azure-sdk-typescript-example",
    },
  });
  console.log("\nUpdated tags:", updated.tags);

  console.log(`\nDeleting resource group "${resourceGroupName}"...`);
  await client.resourceGroups.beginDeleteAndWait(resourceGroupName);
  console.log("Resource group deleted.");
}

main(subscriptionId).catch((error: unknown) => {
  console.error("Resource group operation failed:", error);
  process.exitCode = 1;
});
