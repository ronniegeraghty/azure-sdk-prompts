import { ResourceManagementClient } from "@azure/arm-resources";
import { DefaultAzureCredential } from "@azure/identity";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Set the ${name} environment variable.`);
  }
  return value;
}

const subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
const resourceGroupName =
  process.env.AZURE_RESOURCE_GROUP_NAME ??
  `typescript-sdk-rg-${Date.now().toString(36)}`;
const location = "eastus";

async function main(): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  let resourceGroupCreated = false;

  try {
    console.log(`Creating resource group "${resourceGroupName}" in ${location}...`);
    const createdResourceGroup = await client.resourceGroups.createOrUpdate(
      resourceGroupName,
      { location },
    );
    resourceGroupCreated = true;
    console.log("Created:", createdResourceGroup);

    console.log("\nResource groups in the subscription:");
    for await (const resourceGroup of client.resourceGroups.list()) {
      console.log(`- ${resourceGroup.name} (${resourceGroup.location})`);
    }

    const resourceGroup = await client.resourceGroups.get(resourceGroupName);
    console.log("\nCreated resource group details:", resourceGroup);

    const updatedResourceGroup = await client.resourceGroups.update(
      resourceGroupName,
      {
        tags: {
          ...resourceGroup.tags,
          managedBy: "typescript-azure-sdk",
        },
      },
    );
    console.log("\nUpdated resource group:", updatedResourceGroup);
  } finally {
    if (resourceGroupCreated) {
      console.log(`\nDeleting resource group "${resourceGroupName}"...`);
      await client.resourceGroups.beginDeleteAndWait(resourceGroupName);
      console.log("Resource group deleted.");
    }
  }
}

main().catch((error: unknown) => {
  console.error("Resource group operation failed:", error);
  process.exitCode = 1;
});
