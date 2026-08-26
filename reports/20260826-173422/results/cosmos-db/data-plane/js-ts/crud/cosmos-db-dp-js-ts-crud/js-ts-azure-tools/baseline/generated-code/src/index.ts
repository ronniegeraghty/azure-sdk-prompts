import { CosmosClient } from "@azure/cosmos";

interface Item {
  id: string;
  category: string;
  name: string;
  quantity: number;
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(
      `${name} is required. Set it to the corresponding Cosmos DB emulator value.`,
    );
  }

  return value;
}

function requireStatus(
  operation: string,
  actual: number | undefined,
  expected: readonly number[],
): void {
  if (actual === undefined || !expected.includes(actual)) {
    throw new Error(
      `${operation} returned status ${actual ?? "unknown"}; expected ${expected.join(" or ")}.`,
    );
  }
}

function getStatusCode(error: unknown): number | undefined {
  if (typeof error !== "object" || error === null) {
    return undefined;
  }

  const statusCode = Reflect.get(error, "code") ?? Reflect.get(error, "statusCode");
  return typeof statusCode === "number" ? statusCode : undefined;
}

function getErrorMessage(error: unknown): string {
  return error instanceof Error ? error.message : String(error);
}

async function main(): Promise<void> {
  const endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
  const key = requireEnvironmentVariable("COSMOS_KEY");
  const client = new CosmosClient({ endpoint, key });

  const databaseResponse = await client.databases.createIfNotExists({
    id: "TestDB",
  });
  requireStatus("Create database", databaseResponse.statusCode, [200, 201]);

  const containerResponse =
    await databaseResponse.database.containers.createIfNotExists({
      id: "Items",
      partitionKey: { paths: ["/category"] },
    });
  requireStatus("Create container", containerResponse.statusCode, [200, 201]);

  const container = containerResponse.container;
  const item: Item = {
    id: "item-1",
    category: "electronics",
    name: "Wireless headphones",
    quantity: 10,
  };

  const createResponse = await container.items.create<Item>(item);
  requireStatus("Create item", createResponse.statusCode, [201]);
  console.log("Created:", createResponse.resource);

  const itemReference = container.item(item.id, item.category);
  const readResponse = await itemReference.read<Item>();
  requireStatus("Read item", readResponse.statusCode, [200]);
  if (!readResponse.resource) {
    throw new Error("Read item succeeded but returned no resource.");
  }
  console.log("Read:", readResponse.resource);

  const query = {
    query: "SELECT * FROM c WHERE c.category = @category",
    parameters: [{ name: "@category", value: "electronics" }],
  };
  const queryResponse = await container.items
    .query<Item>(query)
    .fetchAll();
  console.log("Query results:", queryResponse.resources);

  const updatedItem: Item = {
    ...readResponse.resource,
    quantity: 15,
  };
  const replaceResponse = await itemReference.replace<Item>(updatedItem);
  requireStatus("Replace item", replaceResponse.statusCode, [200]);
  console.log("Updated:", replaceResponse.resource);

  const deleteResponse = await itemReference.delete();
  requireStatus("Delete item", deleteResponse.statusCode, [204]);
  console.log(`Deleted item ${item.id}.`);
}

main().catch((error: unknown) => {
  const statusCode = getStatusCode(error);

  if (statusCode === 404) {
    console.error("Cosmos DB resource not found:", getErrorMessage(error));
  } else if (statusCode === 409) {
    console.error("Cosmos DB resource already exists:", getErrorMessage(error));
  } else if (statusCode === 429) {
    console.error("Cosmos DB request was rate limited:", getErrorMessage(error));
  } else if (statusCode !== undefined) {
    console.error(
      `Cosmos DB request failed with status ${statusCode}:`,
      getErrorMessage(error),
    );
  } else {
    console.error("Unexpected error:", getErrorMessage(error));
  }

  process.exitCode = 1;
});
