using System.Net;
using Microsoft.Azure.Cosmos;

const string databaseName = "TestDB";
const string containerName = "Items";
const string partitionKeyPath = "/category";
const string category = "electronics";

try
{
    string connectionString =
        Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING")
        ?? throw new InvalidOperationException(
            "Set COSMOS_CONNECTION_STRING to an Azure Cosmos DB connection string.");

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
        await database.CreateContainerIfNotExistsAsync(
            id: containerName,
            partitionKeyPath: partitionKeyPath);
    Container container = containerResponse.Container;

    Item item = new(
        id: Guid.NewGuid().ToString(),
        category: category,
        name: "Bluetooth Speaker",
        quantity: 10);
    PartitionKey partitionKey = new(item.category);

    ItemResponse<Item> createResponse =
        await container.CreateItemAsync(item, partitionKey);
    Console.WriteLine(
        $"Created '{createResponse.Resource.name}' " +
        $"(RU: {createResponse.RequestCharge:F2}).");

    ItemResponse<Item> readResponse =
        await container.ReadItemAsync<Item>(item.id, partitionKey);
    Console.WriteLine(
        $"Read '{readResponse.Resource.name}' with quantity " +
        $"{readResponse.Resource.quantity} " +
        $"(RU: {readResponse.RequestCharge:F2}).");

    QueryDefinition query = new(
        "SELECT * FROM c WHERE c.category = @category");
    query.WithParameter("@category", category);

    using FeedIterator<Item> iterator =
        container.GetItemQueryIterator<Item>(
            query,
            requestOptions: new QueryRequestOptions
            {
                PartitionKey = partitionKey
            });

    while (iterator.HasMoreResults)
    {
        FeedResponse<Item> page = await iterator.ReadNextAsync();
        foreach (Item result in page)
        {
            Console.WriteLine(
                $"Query result: {result.id}, {result.name}, " +
                $"quantity {result.quantity}");
        }
    }

    Item updatedItem = item with { quantity = 25 };
    ItemResponse<Item> replaceResponse =
        await container.ReplaceItemAsync(
            updatedItem,
            updatedItem.id,
            partitionKey);
    Console.WriteLine(
        $"Updated quantity to {replaceResponse.Resource.quantity} " +
        $"(RU: {replaceResponse.RequestCharge:F2}).");

    ItemResponse<Item> deleteResponse =
        await container.DeleteItemAsync<Item>(item.id, partitionKey);
    Console.WriteLine(
        $"Deleted item '{item.id}' " +
        $"(RU: {deleteResponse.RequestCharge:F2}).");
}
catch (CosmosException exception)
{
    Console.Error.WriteLine(
        $"Cosmos DB request failed: {(int)exception.StatusCode} " +
        $"{exception.StatusCode}");
    Console.Error.WriteLine($"Message: {exception.Message}");
    Console.Error.WriteLine($"Activity ID: {exception.ActivityId}");
    Console.Error.WriteLine($"Substatus code: {exception.SubStatusCode}");
    Console.Error.WriteLine($"Request charge: {exception.RequestCharge:F2} RU");

    if (exception.StatusCode == HttpStatusCode.TooManyRequests &&
        exception.RetryAfter is TimeSpan retryAfter)
    {
        Console.Error.WriteLine(
            $"Retry after: {retryAfter.TotalMilliseconds:F0} ms");
    }

    Environment.ExitCode = 1;
}
catch (InvalidOperationException exception)
{
    Console.Error.WriteLine($"Configuration error: {exception.Message}");
    Environment.ExitCode = 2;
}
catch (Exception exception)
{
    Console.Error.WriteLine(
        $"Unexpected error ({exception.GetType().Name}): {exception.Message}");
    Environment.ExitCode = 3;
}

public sealed record Item(
    string id,
    string category,
    string name,
    int quantity);
