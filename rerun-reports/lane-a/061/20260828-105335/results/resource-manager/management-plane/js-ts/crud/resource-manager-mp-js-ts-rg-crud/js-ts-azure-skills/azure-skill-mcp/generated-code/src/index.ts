import { ResourceManagementClient } from "@azure/arm-resources";
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
  const resourceGroupName = requireEnvironmentVariable("AZURE_RESOURCE_GROUP_NAME");

  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);

  const createdResourceGroup = await client.resourceGroups.createOrUpdate(
    resourceGroupName,
    { location: "eastus" },
  );
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

  await client.resourceGroups.beginDeleteAndWait(resourceGroupName);
  console.log(`Deleted resource group: ${resourceGroupName}`);
}

main().catch((error: unknown) => {
  console.error("Resource group operation failed:", error);
  process.exitCode = 1;
});
