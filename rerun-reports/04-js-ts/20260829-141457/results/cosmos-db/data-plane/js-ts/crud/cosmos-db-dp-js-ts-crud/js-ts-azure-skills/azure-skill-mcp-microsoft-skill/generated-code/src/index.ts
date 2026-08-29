import {
  CosmosClient,
  ErrorResponse,
  type SqlQuerySpec,
} from "@azure/cosmos";

interface InventoryItem {
  id: string;
  category: string;
  name: string;
  quantity: number;
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function assertStatus(
  operation: string,
  actualStatus: number,
  expectedStatuses: readonly number[],
): void {
  if (!expectedStatuses.includes(actualStatus)) {
    throw new Error(
      `${operation} returned HTTP ${actualStatus}; expected ${expectedStatuses.join(" or ")}`,
    );
  }
}

function requireResource<T>(
  operation: string,
  resource: T | undefined,
): T {
  if (resource === undefined) {
    throw new Error(`${operation} succeeded but returned no resource`);
  }
  return resource;
}

function reportError(error: unknown): void {
  if (!(error instanceof ErrorResponse)) {
    console.error(
      error instanceof Error ? error.message : "An unknown error occurred",
    );
    return;
  }

  switch (error.code) {
    case 400:
      console.error(`Bad request (400): ${error.message}`);
      break;
    case 401:
    case 403:
      console.error(`Authentication or authorization failed (${error.code}): ${error.message}`);
      break;
    case 404:
      console.error(`Resource not found (404): ${error.message}`);
      break;
    case 409:
      console.error(`Item already exists (409): ${error.message}`);
      break;
    case 412:
      console.error(`Precondition failed (412): ${error.message}`);
      break;
    case 429:
      console.error(
        `Rate limited (429); retry after ${error.retryAfterInMs ?? "an unspecified number of"} ms`,
      );
      break;
    default:
      console.error(`Cosmos DB request failed (${error.code}): ${error.message}`);
  }
}

async function main(): Promise<void> {
  const client = new CosmosClient({
    endpoint: requireEnvironmentVariable("COSMOS_ENDPOINT"),
    key: requireEnvironmentVariable("COSMOS_KEY"),
  });

  try {
    const databaseResponse = await client.databases.createIfNotExists({
      id: "TestDB",
    });
    assertStatus("Create database", databaseResponse.statusCode, [200, 201]);

    const containerResponse =
      await databaseResponse.database.containers.createIfNotExists({
        id: "Items",
        partitionKey: { paths: ["/category"] },
      });
    assertStatus("Create container", containerResponse.statusCode, [200, 201]);

    const container = containerResponse.container;
    const newItem: InventoryItem = {
      id: "item-1",
      category: "electronics",
      name: "Wireless Mouse",
      quantity: 10,
    };

    const createResponse =
      await container.items.create<InventoryItem>(newItem);
    assertStatus("Create item", createResponse.statusCode, [201]);
    console.log("Created:", requireResource("Create item", createResponse.resource));

    const itemReference = container.item(newItem.id, newItem.category);
    const readResponse = await itemReference.read<InventoryItem>();
    assertStatus("Read item", readResponse.statusCode, [200]);
    const storedItem = requireResource("Read item", readResponse.resource);
    console.log("Read:", storedItem);

    const querySpec: SqlQuerySpec = {
      query: "SELECT * FROM c WHERE c.category = @category",
      parameters: [{ name: "@category", value: "electronics" }],
    };
    const queryResponse = await container.items
      .query<InventoryItem>(querySpec, { partitionKey: "electronics" })
      .fetchAll();
    console.log("Query results:", queryResponse.resources);

    const updatedItem: InventoryItem = {
      ...storedItem,
      quantity: 25,
    };
    const replaceResponse =
      await itemReference.replace<InventoryItem>(updatedItem);
    assertStatus("Replace item", replaceResponse.statusCode, [200]);
    console.log(
      "Replaced:",
      requireResource("Replace item", replaceResponse.resource),
    );

    const deleteResponse = await itemReference.delete();
    assertStatus("Delete item", deleteResponse.statusCode, [204]);
    console.log(`Deleted item ${newItem.id}`);
  } finally {
    client.dispose();
  }
}

main().catch((error: unknown) => {
  reportError(error);
  process.exitCode = 1;
});
