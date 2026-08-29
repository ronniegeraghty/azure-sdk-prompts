import { CosmosClient, type SqlQuerySpec } from "@azure/cosmos";

interface InventoryItem {
  id: string;
  category: string;
  name: string;
  quantity: number;
}

interface CosmosServiceError {
  code: number;
  message: string;
  activityId?: string;
}

const databaseId = "TestDB";
const containerId = "Items";
const partitionKeyPath = "/category";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function expectStatus(
  operation: string,
  statusCode: number | undefined,
  expectedStatusCodes: readonly number[],
): void {
  if (statusCode === undefined || !expectedStatusCodes.includes(statusCode)) {
    throw new Error(
      `${operation} returned HTTP ${statusCode ?? "unknown"}; expected ${expectedStatusCodes.join(" or ")}`,
    );
  }
}

function isCosmosServiceError(error: unknown): error is CosmosServiceError {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as Record<string, unknown>;
  return (
    typeof candidate.code === "number" && typeof candidate.message === "string"
  );
}

async function main(): Promise<void> {
  const endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
  const key = requireEnvironmentVariable("COSMOS_KEY");
  const client = new CosmosClient({ endpoint, key });

  const databaseResponse = await client.databases.createIfNotExists({
    id: databaseId,
  });
  expectStatus("Create database", databaseResponse.statusCode, [200, 201]);
  const { database } = databaseResponse;

  const containerResponse = await database.containers.createIfNotExists({
    id: containerId,
    partitionKey: partitionKeyPath,
  });
  expectStatus("Create container", containerResponse.statusCode, [200, 201]);
  const { container } = containerResponse;

  const newItem: InventoryItem = {
    id: "item-001",
    category: "electronics",
    name: "Wireless headphones",
    quantity: 10,
  };

  const createResponse = await container.items.create<InventoryItem>(newItem);
  expectStatus("Create item", createResponse.statusCode, [201]);
  console.log("Created:", createResponse.resource);

  const item = container.item(newItem.id, newItem.category);
  const readResponse = await item.read<InventoryItem>();
  expectStatus("Read item", readResponse.statusCode, [200]);
  if (!readResponse.resource) {
    throw new Error("Read item succeeded but returned no resource");
  }
  console.log("Read:", readResponse.resource);

  const query: SqlQuerySpec = {
    query: "SELECT * FROM items i WHERE i.category = @category",
    parameters: [{ name: "@category", value: "electronics" }],
  };
  const queryResponse = await container.items
    .query<InventoryItem>(query)
    .fetchAll();
  console.log("Query results:", queryResponse.resources);

  const updatedItem: InventoryItem = {
    ...readResponse.resource,
    quantity: 25,
  };
  const replaceResponse = await item.replace<InventoryItem>(updatedItem);
  expectStatus("Replace item", replaceResponse.statusCode, [200]);
  console.log("Replaced:", replaceResponse.resource);

  const deleteResponse = await item.delete();
  expectStatus("Delete item", deleteResponse.statusCode, [204]);
  console.log(`Deleted item "${newItem.id}"`);
}

main().catch((error: unknown) => {
  if (isCosmosServiceError(error)) {
    const activity = error.activityId
      ? `, activity ID ${error.activityId}`
      : "";
    console.error(
      `Azure Cosmos DB request failed with HTTP ${error.code}${activity}: ${error.message}`,
    );
  } else if (error instanceof Error) {
    console.error(error.message);
  } else {
    console.error("An unknown error occurred:", error);
  }

  process.exitCode = 1;
});
