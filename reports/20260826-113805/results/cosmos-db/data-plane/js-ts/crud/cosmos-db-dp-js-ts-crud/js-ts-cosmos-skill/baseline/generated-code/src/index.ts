import { randomUUID } from "node:crypto";
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

const endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
const key = requireEnvironmentVariable("COSMOS_KEY");

const client = new CosmosClient({ endpoint, key });

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
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
      `${operation} returned status ${actual}; expected ${expected.join(" or ")}.`,
    );
  }
}

function describeCosmosError(error: ErrorResponse): string {
  const statusCode = Number(error.code);

  switch (statusCode) {
    case 400:
      return "Bad request. Check item data and partition key configuration.";
    case 401:
      return "Unauthorized. Check COSMOS_ENDPOINT and COSMOS_KEY.";
    case 403:
      return "Forbidden. The supplied key lacks permission for this operation.";
    case 404:
      return "The requested database, container, or item was not found.";
    case 409:
      return "Conflict. A resource with the same ID already exists.";
    case 412:
      return "Precondition failed. The resource was modified by another operation.";
    case 429:
      return `Rate limited. Retry after ${error.retryAfterInMs ?? "the advised delay"} ms.`;
    default:
      return `Cosmos DB request failed with status ${statusCode}.`;
  }
}

async function main(): Promise<void> {
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
    id: randomUUID(),
    category: "electronics",
    name: "Wireless keyboard",
    quantity: 10,
  };

  const createResponse =
    await container.items.create<InventoryItem>(newItem);
  assertStatus("Create item", createResponse.statusCode, [201]);
  if (!createResponse.resource) {
    throw new Error("Create item succeeded without returning the item.");
  }
  console.log("Created:", createResponse.resource);

  const itemReference = container.item(newItem.id, newItem.category);
  const readResponse = await itemReference.read<InventoryItem>();
  assertStatus("Read item", readResponse.statusCode, [200]);
  if (!readResponse.resource) {
    throw new Error("Read item succeeded without returning the item.");
  }
  console.log("Read:", readResponse.resource);

  const query: SqlQuerySpec = {
    query: "SELECT * FROM c WHERE c.category = @category",
    parameters: [{ name: "@category", value: "electronics" }],
  };
  const queryResponse = await container.items
    .query<InventoryItem>(query, { partitionKey: "electronics" })
    .fetchAll();
  console.log("Query results:", queryResponse.resources);

  const updatedItem: InventoryItem = {
    ...readResponse.resource,
    quantity: 25,
  };
  const replaceResponse =
    await itemReference.replace<InventoryItem>(updatedItem);
  assertStatus("Replace item", replaceResponse.statusCode, [200]);
  if (!replaceResponse.resource) {
    throw new Error("Replace item succeeded without returning the updated item.");
  }
  console.log("Replaced:", replaceResponse.resource);

  const deleteResponse = await itemReference.delete();
  assertStatus("Delete item", deleteResponse.statusCode, [204]);
  console.log(`Deleted item ${newItem.id}.`);
}

try {
  await main();
} catch (error: unknown) {
  if (error instanceof ErrorResponse) {
    console.error(describeCosmosError(error));
    console.error(error.message);
    process.exitCode = 1;
  } else if (error instanceof Error) {
    console.error(error.message);
    process.exitCode = 1;
  } else {
    console.error("An unknown error occurred.", error);
    process.exitCode = 1;
  }
} finally {
  client.dispose();
}
