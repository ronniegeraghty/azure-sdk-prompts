using System.Net;
using Microsoft.Azure.Cosmos;

const string databaseName = "TestDB";
const string containerName = "Items";
const string category = "electronics";
const string itemId = "item-001";

string? connectionString =
    Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING");

if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        "Set the COSMOS_CONNECTION_STRING environment variable before running the sample.");
    return 1;
}

try
{
    using CosmosClient client = new(connectionString);

    DatabaseResponse databaseResponse =
        await client.CreateDatabaseIfNotExistsAsync(databaseName);
    Database database = databaseResponse.Database;

    ContainerResponse containerResponse =
        await database.CreateContainerIfNotExistsAsync(
            new ContainerProperties(containerName, "/category"));
    Container container = containerResponse.Container;

    Item item = new()
    {
        id = itemId,
        category = category,
        name = "Wireless headphones",
        quantity = 10
    };

    ItemResponse<Item> createResponse = await container.CreateItemAsync(
        item,
        new PartitionKey(item.category));
    Console.WriteLine(
        $"Created {createResponse.Resource.id}; RU charge: {createResponse.RequestCharge:F2}");

    ItemResponse<Item> readResponse = await container.ReadItemAsync<Item>(
        itemId,
        new PartitionKey(category));
    Console.WriteLine(
        $"Read {readResponse.Resource.name}, quantity {readResponse.Resource.quantity}");

    QueryDefinition query = new(
        "SELECT * FROM items i WHERE i.category = @category");
    query.WithParameter("@category", category);

    using FeedIterator<Item> queryIterator = container.GetItemQueryIterator<Item>(
        query,
        requestOptions: new QueryRequestOptions
        {
            PartitionKey = new PartitionKey(category)
        });

    while (queryIterator.HasMoreResults)
    {
        FeedResponse<Item> page = await queryIterator.ReadNextAsync();
        foreach (Item result in page)
        {
            Console.WriteLine(
                $"Query result: {result.id}, {result.name}, quantity {result.quantity}");
        }
    }

    item.quantity = 25;
    ItemResponse<Item> replaceResponse = await container.ReplaceItemAsync(
        item,
        item.id,
        new PartitionKey(item.category));
    Console.WriteLine(
        $"Updated {replaceResponse.Resource.id} to quantity {replaceResponse.Resource.quantity}");

    ItemResponse<Item> deleteResponse = await container.DeleteItemAsync<Item>(
        item.id,
        new PartitionKey(item.category));
    Console.WriteLine(
        $"Deleted {item.id}; status: {deleteResponse.StatusCode}");

    return 0;
}
catch (CosmosException exception)
{
    Console.Error.WriteLine(
        $"Cosmos DB request failed ({(int)exception.StatusCode} {exception.StatusCode}).");
    Console.Error.WriteLine($"Message: {exception.Message}");
    Console.Error.WriteLine($"Activity ID: {exception.ActivityId}");
    Console.Error.WriteLine($"Request charge: {exception.RequestCharge:F2} RU");

    if (exception.StatusCode == HttpStatusCode.TooManyRequests)
    {
        Console.Error.WriteLine($"Retry after: {exception.RetryAfter}");
    }

    return 1;
}
catch (ArgumentException exception)
{
    Console.Error.WriteLine($"Invalid Cosmos DB configuration: {exception.Message}");
    return 1;
}

internal sealed class Item
{
    public required string id { get; init; }

    public required string category { get; init; }

    public required string name { get; init; }

    public int quantity { get; set; }
}
