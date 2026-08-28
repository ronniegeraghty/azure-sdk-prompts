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

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function requireLocalEndpoint(endpoint: string): void {
  const hostname = new URL(endpoint).hostname;
  if (hostname !== "localhost" && hostname !== "127.0.0.1") {
    throw new Error(
      "This example is local-only. COSMOS_ENDPOINT must target a local Cosmos DB emulator.",
    );
  }
}

function checkStatus(operation: string, statusCode: number | undefined): void {
  if (statusCode === undefined || statusCode < 200 || statusCode >= 300) {
    throw new Error(
      `${operation} returned unexpected status code ${statusCode ?? "unknown"}`,
    );
  }
}

function reportError(error: unknown): void {
  if (!(error instanceof ErrorResponse)) {
    console.error(error instanceof Error ? error.message : error);
    return;
  }

  switch (error.code) {
    case 400:
      console.error(`Bad request (400): ${error.message}`);
      break;
    case 401:
    case 403:
      console.error(`Authentication or authorization failed (${error.code}).`);
      break;
    case 404:
      console.error("Database, container, or item not found (404).");
      break;
    case 409:
      console.error("The resource already exists (409).");
      break;
    case 412:
      console.error("The resource changed since it was read (412).");
      break;
    case 429:
      console.error(
        `Request rate-limited (429). Retry after ${error.retryAfterInMs ?? "the server-provided delay"} ms.`,
      );
      break;
    default:
      console.error(`Cosmos DB error ${error.code}: ${error.message}`);
  }
}

async function main(): Promise<void> {
  const endpoint = requiredEnvironmentVariable("COSMOS_ENDPOINT");
  const key = requiredEnvironmentVariable("COSMOS_KEY");
  requireLocalEndpoint(endpoint);

  const client = new CosmosClient({ endpoint, key });

  try {
    const databaseResponse = await client.databases.createIfNotExists({
      id: "TestDB",
    });
    checkStatus("Create database", databaseResponse.statusCode);

    const containerResponse =
      await databaseResponse.database.containers.createIfNotExists({
        id: "Items",
        partitionKey: { paths: ["/category"] },
      });
    checkStatus("Create container", containerResponse.statusCode);

    const container = containerResponse.container;
    const item: InventoryItem = {
      id: "item-1",
      category: "electronics",
      name: "Wireless headphones",
      quantity: 10,
    };

    const createResponse =
      await container.items.create<InventoryItem>(item);
    checkStatus("Create item", createResponse.statusCode);
    console.log("Created:", createResponse.resource);

    const itemReference = container.item(item.id, item.category);
    const readResponse = await itemReference.read<InventoryItem>();
    checkStatus("Read item", readResponse.statusCode);
    if (!readResponse.resource) {
      throw new Error("Read succeeded but returned no item.");
    }
    console.log("Read:", readResponse.resource);

    const query: SqlQuerySpec = {
      query: "SELECT * FROM c WHERE c.category = @category",
      parameters: [{ name: "@category", value: "electronics" }],
    };
    const queryResponse = await container.items
      .query<InventoryItem>(query, { partitionKey: "electronics" })
      .fetchAll();
    checkStatus("Query items", queryResponse.statusCode);
    console.log("Query results:", queryResponse.resources);

    const updatedItem: InventoryItem = {
      ...readResponse.resource,
      quantity: 15,
    };
    const replaceResponse =
      await itemReference.replace<InventoryItem>(updatedItem);
    checkStatus("Replace item", replaceResponse.statusCode);
    console.log("Updated:", replaceResponse.resource);

    const deleteResponse = await itemReference.delete();
    checkStatus("Delete item", deleteResponse.statusCode);
    console.log(`Deleted ${item.id} (status ${deleteResponse.statusCode}).`);
  } catch (error: unknown) {
    reportError(error);
    process.exitCode = 1;
  } finally {
    client.dispose();
  }
}

void main();
