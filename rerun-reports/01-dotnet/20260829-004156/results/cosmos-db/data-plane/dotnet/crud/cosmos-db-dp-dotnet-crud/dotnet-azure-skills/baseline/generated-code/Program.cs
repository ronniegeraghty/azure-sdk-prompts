using Microsoft.Azure.Cosmos;

const string databaseName = "TestDB";
const string containerName = "Items";
const string category = "electronics";

string? connectionString = Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING");
if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        "Set the COSMOS_CONNECTION_STRING environment variable before running the application.");
    return 1;
}

CosmosClientOptions clientOptions = new()
{
    ApplicationName = "CosmosCrudSample",
    SerializerOptions = new CosmosSerializationOptions
    {
        PropertyNamingPolicy = CosmosPropertyNamingPolicy.CamelCase
    }
};

try
{
    using CosmosClient client = new(connectionString, clientOptions);

    DatabaseResponse databaseResponse =
        await client.CreateDatabaseIfNotExistsAsync(databaseName);
    Database database = databaseResponse.Database;

    ContainerResponse containerResponse = await database.CreateContainerIfNotExistsAsync(
        new ContainerProperties(containerName, "/category"));
    Container container = containerResponse.Container;

    Item item = new()
    {
        Id = Guid.NewGuid().ToString(),
        Category = category,
        Name = "Wireless headphones",
        Quantity = 10
    };
    PartitionKey partitionKey = new(item.Category);

    ItemResponse<Item> createResponse =
        await container.CreateItemAsync(item, partitionKey);
    Console.WriteLine(
        $"Created item {createResponse.Resource.Id} " +
        $"(request charge: {createResponse.RequestCharge:F2} RU).");

    ItemResponse<Item> readResponse =
        await container.ReadItemAsync<Item>(item.Id, partitionKey);
    Console.WriteLine(
        $"Read item: {readResponse.Resource.Name}, " +
        $"quantity {readResponse.Resource.Quantity}.");

    QueryDefinition query = new(
        "SELECT * FROM c WHERE c.category = @category");
    query.WithParameter("@category", category);

    using FeedIterator<Item> iterator = container.GetItemQueryIterator<Item>(
        query,
        requestOptions: new QueryRequestOptions
        {
            PartitionKey = partitionKey
        });

    Console.WriteLine($"Items in category '{category}':");
    while (iterator.HasMoreResults)
    {
        FeedResponse<Item> page = await iterator.ReadNextAsync();
        foreach (Item result in page)
        {
            Console.WriteLine($"- {result.Id}: {result.Name} ({result.Quantity})");
        }
    }

    item.Quantity = 25;
    ItemResponse<Item> replaceResponse =
        await container.ReplaceItemAsync(item, item.Id, partitionKey);
    Console.WriteLine(
        $"Updated quantity to {replaceResponse.Resource.Quantity}.");

    ItemResponse<Item> deleteResponse =
        await container.DeleteItemAsync<Item>(item.Id, partitionKey);
    Console.WriteLine(
        $"Deleted item {item.Id} " +
        $"(request charge: {deleteResponse.RequestCharge:F2} RU).");

    return 0;
}
catch (CosmosException exception)
{
    Console.Error.WriteLine(
        $"Cosmos DB request failed with HTTP {(int)exception.StatusCode} " +
        $"({exception.StatusCode}).");
    Console.Error.WriteLine($"Activity ID: {exception.ActivityId}");
    Console.Error.WriteLine($"Request charge: {exception.RequestCharge:F2} RU");
    Console.Error.WriteLine(exception.Message);
    return 1;
}
catch (ArgumentException exception)
{
    Console.Error.WriteLine($"The Cosmos DB configuration is invalid: {exception.Message}");
    return 1;
}

internal sealed class Item
{
    public required string Id { get; init; }

    public required string Category { get; init; }

    public required string Name { get; init; }

    public int Quantity { get; set; }
}
