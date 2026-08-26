import {
  CosmosClient,
  type ErrorResponse,
  type OperationResponse,
} from "@azure/cosmos";

interface Item {
  id: string;
  category: string;
  name: string;
  quantity: number;
}

const endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
const key = requireEnvironmentVariable("COSMOS_KEY");

const client = new CosmosClient({ endpoint, key });

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function assertStatus(
  operation: string,
  response: OperationResponse,
  expectedStatusCodes: readonly number[],
): void {
  if (!expectedStatusCodes.includes(response.statusCode)) {
    throw new Error(
      `${operation} returned HTTP ${response.statusCode}; expected ${expectedStatusCodes.join(" or ")}`,
    );
  }
}

function isCosmosError(error: unknown): error is ErrorResponse {
  return (
    error instanceof Error &&
    "code" in error &&
    typeof error.code === "number"
  );
}

async function main(): Promise<void> {
  const databaseResponse = await client.databases.createIfNotExists({
    id: "TestDB",
  });
  assertStatus("Create database", databaseResponse, [200, 201]);
  console.log(`Database ready: ${databaseResponse.database.id}`);

  const containerResponse =
    await databaseResponse.database.containers.createIfNotExists({
      id: "Items",
      partitionKey: {
        paths: ["/category"],
      },
    });
  assertStatus("Create container", containerResponse, [200, 201]);
  console.log(`Container ready: ${containerResponse.container.id}`);

  const container = containerResponse.container;
  const newItem: Item = {
    id: "item-1",
    category: "electronics",
    name: "Wireless keyboard",
    quantity: 10,
  };

  const createResponse = await container.items.create<Item>(newItem);
  assertStatus("Create item", createResponse, [201]);
  console.log("Created item:", createResponse.resource);

  const item = container.item(newItem.id, newItem.category);
  const readResponse = await item.read<Item>();
  assertStatus("Read item", readResponse, [200]);
  if (!readResponse.resource) {
    throw new Error("Read item returned HTTP 200 without a resource");
  }
  console.log("Read item:", readResponse.resource);

  const query = {
    query: "SELECT * FROM item WHERE item.category = @category",
    parameters: [{ name: "@category", value: "electronics" }],
  };
  const queryResponse = await container.items.query<Item>(query).fetchAll();
  console.log(
    `Queried ${queryResponse.resources.length} electronics item(s):`,
    queryResponse.resources,
  );

  const updatedItem: Item = {
    ...readResponse.resource,
    quantity: 15,
  };
  const replaceResponse = await item.replace<Item>(updatedItem);
  assertStatus("Replace item", replaceResponse, [200]);
  console.log("Replaced item:", replaceResponse.resource);

  const deleteResponse = await item.delete();
  assertStatus("Delete item", deleteResponse, [204]);
  console.log(`Deleted item: ${newItem.id}`);
}

main().catch((error: unknown) => {
  if (isCosmosError(error)) {
    switch (error.code) {
      case 401:
      case 403:
        console.error(`Cosmos DB authorization failed (HTTP ${error.code}).`);
        break;
      case 404:
        console.error("Cosmos DB resource not found (HTTP 404).");
        break;
      case 409:
        console.error("Cosmos DB resource already exists (HTTP 409).");
        break;
      case 429:
        console.error(
          `Cosmos DB rate limit exceeded (HTTP 429). Retry after ${error.retryAfterInMs ?? "the server-specified delay"} ms.`,
        );
        break;
      default:
        console.error(
          `Cosmos DB request failed (HTTP ${error.code}): ${error.message}`,
        );
    }
  } else {
    console.error(
      "Unexpected error:",
      error instanceof Error ? error.message : error,
    );
  }

  process.exitCode = 1;
});
