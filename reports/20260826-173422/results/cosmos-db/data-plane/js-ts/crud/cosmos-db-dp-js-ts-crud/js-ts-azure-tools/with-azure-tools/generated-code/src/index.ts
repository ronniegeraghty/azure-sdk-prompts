import {
  CosmosClient,
  ErrorResponse,
  type SqlQuerySpec,
} from "@azure/cosmos";
import { randomUUID } from "node:crypto";

const DATABASE_ID = "TestDB";
const CONTAINER_ID = "Items";
const PARTITION_KEY_PATH = "/category";

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

function expectStatus(
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

async function main(): Promise<void> {
  const endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
  const key = requireEnvironmentVariable("COSMOS_KEY");
  const client = new CosmosClient({ endpoint, key });

  try {
    const databaseResponse = await client.databases.createIfNotExists({
      id: DATABASE_ID,
    });
    expectStatus(
      "Create database",
      databaseResponse.statusCode,
      [200, 201],
    );
    console.log(`Database ready: ${databaseResponse.database.id}`);

    const containerResponse =
      await databaseResponse.database.containers.createIfNotExists({
        id: CONTAINER_ID,
        partitionKey: { paths: [PARTITION_KEY_PATH] },
      });
    expectStatus(
      "Create container",
      containerResponse.statusCode,
      [200, 201],
    );
    const { container } = containerResponse;
    console.log(`Container ready: ${container.id}`);

    const item: InventoryItem = {
      id: randomUUID(),
      category: "electronics",
      name: "Wireless keyboard",
      quantity: 10,
    };

    const createResponse = await container.items.create<InventoryItem>(item);
    expectStatus("Create item", createResponse.statusCode, [201]);
    if (!createResponse.resource) {
      throw new Error("Create item succeeded without returning the item");
    }
    console.log("Created:", createResponse.resource);

    const itemReference = container.item(item.id, item.category);
    const readResponse = await itemReference.read<InventoryItem>();
    expectStatus("Read item", readResponse.statusCode, [200]);
    if (!readResponse.resource) {
      throw new Error("Read item succeeded without returning the item");
    }
    console.log("Read:", readResponse.resource);

    const querySpec: SqlQuerySpec = {
      query: "SELECT * FROM c WHERE c.category = @category",
      parameters: [{ name: "@category", value: "electronics" }],
    };
    const queryResponse = await container.items
      .query<InventoryItem>(querySpec, { partitionKey: "electronics" })
      .fetchAll();
    console.log("Query results:", queryResponse.resources);

    const replacement: InventoryItem = {
      ...readResponse.resource,
      quantity: 25,
    };
    const replaceResponse =
      await itemReference.replace<InventoryItem>(replacement);
    expectStatus("Replace item", replaceResponse.statusCode, [200]);
    if (!replaceResponse.resource) {
      throw new Error("Replace item succeeded without returning the item");
    }
    console.log("Replaced:", replaceResponse.resource);

    const deleteResponse = await itemReference.delete();
    expectStatus("Delete item", deleteResponse.statusCode, [204]);
    console.log(`Deleted item ${item.id}`);
  } finally {
    client.dispose();
  }
}

function handleError(error: unknown): void {
  if (error instanceof ErrorResponse) {
    const statusCode = error.code;

    switch (statusCode) {
      case 400:
        console.error("Cosmos DB rejected the request as invalid:", error.message);
        break;
      case 401:
      case 403:
        console.error(
          `Cosmos DB authorization failed (HTTP ${statusCode}). Check the endpoint and key.`,
        );
        break;
      case 404:
        console.error("The requested Cosmos DB resource was not found.");
        break;
      case 409:
        console.error("The item already exists.");
        break;
      case 412:
        console.error("The item changed before it could be replaced.");
        break;
      case 429:
        console.error(
          `Cosmos DB rate limit exceeded. Retry after ${error.retryAfterInMs ?? "the server-specified delay"} ms.`,
        );
        break;
      default:
        console.error(`Cosmos DB error ${statusCode}: ${error.message}`);
    }
  } else if (error instanceof Error) {
    console.error(error.message);
  } else {
    console.error("An unknown error occurred:", error);
  }

  process.exitCode = 1;
}

void main().catch(handleError);
