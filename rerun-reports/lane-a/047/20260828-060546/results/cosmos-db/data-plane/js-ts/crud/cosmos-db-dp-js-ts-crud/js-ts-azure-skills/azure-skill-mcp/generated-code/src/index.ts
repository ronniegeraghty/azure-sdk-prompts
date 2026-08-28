import { CosmosClient } from "@azure/cosmos";

interface InventoryItem {
  id: string;
  category: string;
  name: string;
  quantity: number;
}

const DATABASE_ID = "TestDB";
const CONTAINER_ID = "Items";
const PARTITION_KEY_PATH = "/category";

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function assertStatus(
  operation: string,
  actual: number,
  expected: readonly number[],
): void {
  if (!expected.includes(actual)) {
    throw new Error(
      `${operation} returned HTTP ${actual}; expected ${expected.join(" or ")}`,
    );
  }
}

function requireResource<T>(operation: string, resource: T | undefined): T {
  if (resource === undefined) {
    throw new Error(`${operation} succeeded without returning a resource`);
  }
  return resource;
}

function getStatusCode(error: unknown): number | undefined {
  if (typeof error !== "object" || error === null) {
    return undefined;
  }

  const candidate = error as { code?: unknown; statusCode?: unknown };
  if (typeof candidate.code === "number") {
    return candidate.code;
  }
  if (typeof candidate.statusCode === "number") {
    return candidate.statusCode;
  }
  return undefined;
}

function getErrorMessage(error: unknown): string {
  return error instanceof Error ? error.message : String(error);
}

async function main(): Promise<void> {
  const endpoint = requiredEnvironmentVariable("COSMOS_ENDPOINT");
  const key = requiredEnvironmentVariable("COSMOS_KEY");
  const client = new CosmosClient({ endpoint, key });

  const databaseResponse = await client.databases.createIfNotExists({
    id: DATABASE_ID,
  });
  assertStatus("Create database", databaseResponse.statusCode, [200, 201]);
  const database = databaseResponse.database;

  const containerResponse = await database.containers.createIfNotExists({
    id: CONTAINER_ID,
    partitionKey: {
      paths: [PARTITION_KEY_PATH],
    },
  });
  assertStatus("Create container", containerResponse.statusCode, [200, 201]);
  const container = containerResponse.container;

  const item: InventoryItem = {
    id: "item-001",
    category: "electronics",
    name: "Wireless headphones",
    quantity: 5,
  };

  const createResponse = await container.items.create<InventoryItem>(item);
  assertStatus("Create item", createResponse.statusCode, [201]);
  const createdItem = requireResource("Create item", createResponse.resource);
  console.log("Created:", createdItem);

  const readResponse = await container
    .item(item.id, item.category)
    .read<InventoryItem>();
  assertStatus("Read item", readResponse.statusCode, [200]);
  const readItem = requireResource("Read item", readResponse.resource);
  console.log("Read:", readItem);

  const querySpec = {
    query: "SELECT * FROM items i WHERE i.category = @category",
    parameters: [{ name: "@category", value: "electronics" }],
  };
  const queryResponse = await container.items
    .query<InventoryItem>(querySpec, { partitionKey: "electronics" })
    .fetchAll();
  console.log("Query results:", queryResponse.resources);

  const updatedItem: InventoryItem = {
    ...readItem,
    quantity: 10,
  };
  const replaceResponse = await container
    .item(updatedItem.id, updatedItem.category)
    .replace<InventoryItem>(updatedItem);
  assertStatus("Replace item", replaceResponse.statusCode, [200]);
  console.log(
    "Replaced:",
    requireResource("Replace item", replaceResponse.resource),
  );

  const deleteResponse = await container
    .item(updatedItem.id, updatedItem.category)
    .delete<InventoryItem>();
  assertStatus("Delete item", deleteResponse.statusCode, [204]);
  console.log(`Deleted item ${updatedItem.id}`);
}

main().catch((error: unknown) => {
  const statusCode = getStatusCode(error);
  const message = getErrorMessage(error);

  switch (statusCode) {
    case 400:
      console.error(`Bad Cosmos DB request (HTTP 400): ${message}`);
      break;
    case 401:
    case 403:
      console.error(
        `Cosmos DB authentication/authorization failed (HTTP ${statusCode}): ${message}`,
      );
      break;
    case 404:
      console.error(`Cosmos DB resource not found (HTTP 404): ${message}`);
      break;
    case 409:
      console.error(`Cosmos DB resource conflict (HTTP 409): ${message}`);
      break;
    case 429:
      console.error(`Cosmos DB rate limit exceeded (HTTP 429): ${message}`);
      break;
    default:
      console.error(
        statusCode === undefined
          ? `Unexpected error: ${message}`
          : `Cosmos DB request failed (HTTP ${statusCode}): ${message}`,
      );
  }

  process.exitCode = 1;
});
