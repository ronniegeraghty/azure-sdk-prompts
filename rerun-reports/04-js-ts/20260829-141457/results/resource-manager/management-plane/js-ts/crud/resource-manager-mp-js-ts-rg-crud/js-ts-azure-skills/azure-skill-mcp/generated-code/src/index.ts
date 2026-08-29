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
  process.env.AZURE_RESOURCE_GROUP_NAME ?? `sdk-rg-example-${Date.now()}`;
const location = "eastus";

async function main(): Promise<void> {
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  let resourceGroupCreated = false;

  try {
    console.log(`Creating resource group "${resourceGroupName}" in "${location}"...`);
    const created = await client.resourceGroups.createOrUpdate(resourceGroupName, {
      location,
    });
    resourceGroupCreated = true;
    console.log(`Created: ${created.id}`);

    console.log("\nResource groups in the subscription:");
    for await (const resourceGroup of client.resourceGroups.list()) {
      console.log(`- ${resourceGroup.name} (${resourceGroup.location})`);
    }

    const details = await client.resourceGroups.get(resourceGroupName);
    console.log("\nCreated resource group details:", {
      id: details.id,
      name: details.name,
      location: details.location,
      tags: details.tags,
    });

    const updated = await client.resourceGroups.update(resourceGroupName, {
      tags: {
        ...details.tags,
        environment: "demo",
      },
    });
    console.log("\nUpdated tags:", updated.tags);
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
