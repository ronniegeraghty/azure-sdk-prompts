using System.Net;
using Microsoft.Azure.Cosmos;

const string databaseName = "TestDB";
const string containerName = "Items";
const string category = "electronics";
const string itemId = "item-001";

string? connectionString = Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING");
if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        "Set COSMOS_CONNECTION_STRING to a Cosmos DB for NoSQL connection string. " +
        "For local development, use the Azure Cosmos DB Emulator connection string.");
    return 1;
}

try
{
    using CosmosClient client = new(
        connectionString,
        new CosmosClientOptions
        {
            ApplicationName = "CosmosCrudSample"
        });

    DatabaseResponse databaseResponse =
        await client.CreateDatabaseIfNotExistsAsync(databaseName);
    Database database = databaseResponse.Database;

    ContainerResponse containerResponse =
        await database.CreateContainerIfNotExistsAsync(containerName, "/category");
    Container container = containerResponse.Container;

    Item item = new(itemId, category, "Wireless Headphones", 10);

    ItemResponse<Item> createResponse = await container.CreateItemAsync(
        item,
        new PartitionKey(item.category));
    Console.WriteLine(
        $"Created item '{createResponse.Resource.id}' " +
        $"(request charge: {createResponse.RequestCharge:F2} RU).");

    ItemResponse<Item> readResponse = await container.ReadItemAsync<Item>(
        itemId,
        new PartitionKey(category));
    Console.WriteLine(
        $"Read item: {readResponse.Resource.name}, " +
        $"quantity {readResponse.Resource.quantity}.");

    QueryDefinition query = new QueryDefinition(
        "SELECT * FROM items i WHERE i.category = @category")
        .WithParameter("@category", category);

    using FeedIterator<Item> queryResults = container.GetItemQueryIterator<Item>(
        query,
        requestOptions: new QueryRequestOptions
        {
            PartitionKey = new PartitionKey(category)
        });

    while (queryResults.HasMoreResults)
    {
        FeedResponse<Item> page = await queryResults.ReadNextAsync();
        foreach (Item result in page)
        {
            Console.WriteLine(
                $"Query result: {result.id} - {result.name} " +
                $"(quantity {result.quantity}).");
        }
    }

    Item updatedItem = item with { quantity = 25 };
    ItemRequestOptions replaceOptions = new()
    {
        IfMatchEtag = readResponse.ETag
    };

    ItemResponse<Item> replaceResponse = await container.ReplaceItemAsync(
        updatedItem,
        updatedItem.id,
        new PartitionKey(updatedItem.category),
        replaceOptions);
    Console.WriteLine(
        $"Updated item quantity to {replaceResponse.Resource.quantity}.");

    await container.DeleteItemAsync<Item>(
        itemId,
        new PartitionKey(category));
    Console.WriteLine($"Deleted item '{itemId}'.");

    return 0;
}
catch (CosmosException exception) when (exception.StatusCode == HttpStatusCode.Conflict)
{
    Console.Error.WriteLine(
        $"The item already exists. Cosmos DB returned {(int)exception.StatusCode} " +
        $"{exception.StatusCode}. Activity ID: {exception.ActivityId}. " +
        $"Request charge: {exception.RequestCharge:F2} RU.");
    return 2;
}
catch (CosmosException exception) when (exception.StatusCode == HttpStatusCode.PreconditionFailed)
{
    Console.Error.WriteLine(
        "The item changed after it was read, so the replacement was rejected. " +
        $"Activity ID: {exception.ActivityId}.");
    return 3;
}
catch (CosmosException exception)
{
    Console.Error.WriteLine(
        $"Cosmos DB request failed with {(int)exception.StatusCode} " +
        $"{exception.StatusCode}: {exception.Message}{Environment.NewLine}" +
        $"Activity ID: {exception.ActivityId}; " +
        $"Request charge: {exception.RequestCharge:F2} RU; " +
        $"Retry after: {exception.RetryAfter}.");
    return 4;
}
catch (ArgumentException exception)
{
    Console.Error.WriteLine(
        $"The Cosmos DB connection string or request configuration is invalid: " +
        exception.Message);
    return 5;
}

internal sealed record Item(
    string id,
    string category,
    string name,
    int quantity);
