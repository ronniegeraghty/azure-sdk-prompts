import {
  CosmosClient,
  CosmosClientOptions,
  ErrorResponse,
  SqlQuerySpec,
} from "@azure/cosmos";

const databaseId = "TestDB";
const containerId = "Items";

interface Item {
  id: string;
  category: string;
  name: string;
  quantity: number;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function isCosmosError(error: unknown): error is ErrorResponse {
  return error instanceof ErrorResponse;
}

async function main(): Promise<void> {
  const clientOptions: CosmosClientOptions = {
    endpoint: requiredEnvironmentVariable("COSMOS_ENDPOINT"),
    key: requiredEnvironmentVariable("COSMOS_KEY"),
  };
  const client = new CosmosClient(clientOptions);

  const { database, statusCode: databaseStatus } =
    await client.databases.createIfNotExists({ id: databaseId });
  if (databaseStatus !== 200 && databaseStatus !== 201) {
    throw new Error(
      `Creating database returned unexpected status ${databaseStatus}.`,
    );
  }

  const { container, statusCode: containerStatus } =
    await database.containers.createIfNotExists({
      id: containerId,
      partitionKey: { paths: ["/category"] },
    });
  if (containerStatus !== 200 && containerStatus !== 201) {
    throw new Error(
      `Creating container returned unexpected status ${containerStatus}.`,
    );
  }

  const newItem: Item = {
    id: "item-1",
    category: "electronics",
    name: "Wireless keyboard",
    quantity: 10,
  };

  const { resource: created, statusCode: createStatus } =
    await container.items.create<Item>(newItem);
  if (createStatus !== 201 || !created) {
    throw new Error(`Creating item returned unexpected status ${createStatus}.`);
  }
  console.log("Created:", created);

  const item = container.item(newItem.id, newItem.category);
  const { resource: read, statusCode: readStatus } = await item.read<Item>();
  if (readStatus !== 200 || !read) {
    throw new Error(`Reading item returned unexpected status ${readStatus}.`);
  }
  console.log("Read:", read);

  const query: SqlQuerySpec = {
    query: "SELECT * FROM items i WHERE i.category = @category",
    parameters: [{ name: "@category", value: "electronics" }],
  };
  const { resources: queriedItems } = await container.items
    .query<Item>(query)
    .fetchAll();
  console.log("Query results:", queriedItems);

  const updatedItem: Item = { ...read, quantity: 25 };
  const { resource: replaced, statusCode: replaceStatus } =
    await item.replace<Item>(updatedItem);
  if (replaceStatus !== 200 || !replaced) {
    throw new Error(
      `Replacing item returned unexpected status ${replaceStatus}.`,
    );
  }
  console.log("Replaced:", replaced);

  const { statusCode: deleteStatus } = await item.delete();
  if (deleteStatus !== 204) {
    throw new Error(`Deleting item returned unexpected status ${deleteStatus}.`);
  }
  console.log(`Deleted item ${newItem.id}.`);
}

main().catch((error: unknown) => {
  if (isCosmosError(error)) {
    switch (error.statusCode) {
      case 401:
      case 403:
        console.error(
          `Cosmos DB authorization failed (${error.statusCode}): ${error.message}`,
        );
        break;
      case 404:
        console.error(`Cosmos DB resource not found (404): ${error.message}`);
        break;
      case 409:
        console.error(`Cosmos DB resource conflict (409): ${error.message}`);
        break;
      case 429:
        console.error(
          `Cosmos DB rate limit exceeded (429). Retry after ${error.retryAfterInMs ?? "the server-recommended delay"} ms.`,
        );
        break;
      default:
        console.error(
          `Cosmos DB request failed (${error.statusCode ?? "unknown status"}): ${error.message}`,
        );
    }
  } else {
    console.error(error instanceof Error ? error.message : error);
  }
  process.exitCode = 1;
});
