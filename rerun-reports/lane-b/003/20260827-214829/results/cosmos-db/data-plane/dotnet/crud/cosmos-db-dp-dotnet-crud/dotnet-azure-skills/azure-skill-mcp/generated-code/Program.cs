using System.Net;
using Microsoft.Azure.Cosmos;
using Newtonsoft.Json;

const string databaseName = "TestDB";
const string containerName = "Items";
const string partitionKeyPath = "/category";

try
{
    string connectionString = Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING")
        ?? throw new InvalidOperationException(
            "Set COSMOS_CONNECTION_STRING to the Azure Cosmos DB Emulator connection string.");

    EnsureLocalEmulatorConnection(connectionString);

    using CosmosClient client = new(
        connectionString,
        new CosmosClientOptions { ApplicationName = "CosmosCrudSample" });

    DatabaseResponse databaseResponse =
        await client.CreateDatabaseIfNotExistsAsync(databaseName);
    Database database = databaseResponse.Database;

    ContainerResponse containerResponse =
        await database.CreateContainerIfNotExistsAsync(
            new ContainerProperties(containerName, partitionKeyPath));
    Container container = containerResponse.Container;

    Item item = new()
    {
        Id = Guid.NewGuid().ToString(),
        Category = "electronics",
        Name = "Wireless keyboard",
        Quantity = 10
    };

    ItemResponse<Item> createResponse = await container.CreateItemAsync(
        item,
        new PartitionKey(item.Category));
    Console.WriteLine(
        $"Created item {createResponse.Resource.Id}; request charge: {createResponse.RequestCharge} RUs.");

    ItemResponse<Item> readResponse = await container.ReadItemAsync<Item>(
        item.Id,
        new PartitionKey(item.Category));
    Console.WriteLine(
        $"Read item: {readResponse.Resource.Name}, quantity {readResponse.Resource.Quantity}.");

    QueryDefinition query = new(
        "SELECT * FROM items i WHERE i.category = @category");
    query.WithParameter("@category", "electronics");

    using FeedIterator<Item> results = container.GetItemQueryIterator<Item>(
        query,
        requestOptions: new QueryRequestOptions
        {
            PartitionKey = new PartitionKey("electronics")
        });

    Console.WriteLine("Electronics:");
    while (results.HasMoreResults)
    {
        FeedResponse<Item> page = await results.ReadNextAsync();
        foreach (Item result in page)
        {
            Console.WriteLine($"- {result.Id}: {result.Name} ({result.Quantity})");
        }
    }

    Item updatedItem = readResponse.Resource;
    updatedItem.Quantity = 25;

    ItemResponse<Item> replaceResponse = await container.ReplaceItemAsync(
        updatedItem,
        updatedItem.Id,
        new PartitionKey(updatedItem.Category));
    Console.WriteLine(
        $"Updated item {replaceResponse.Resource.Id} to quantity {replaceResponse.Resource.Quantity}.");

    await container.DeleteItemAsync<Item>(
        updatedItem.Id,
        new PartitionKey(updatedItem.Category));
    Console.WriteLine($"Deleted item {updatedItem.Id}.");
}
catch (CosmosException exception) when (exception.StatusCode == HttpStatusCode.Conflict)
{
    ReportCosmosError("An item with the same id and partition key already exists.", exception);
    Environment.ExitCode = 1;
}
catch (CosmosException exception) when (exception.StatusCode == HttpStatusCode.NotFound)
{
    ReportCosmosError("The requested database, container, or item was not found.", exception);
    Environment.ExitCode = 1;
}
catch (CosmosException exception) when (
    exception.StatusCode == HttpStatusCode.TooManyRequests)
{
    ReportCosmosError(
        $"Request rate was too high. Retry after " +
        $"{exception.RetryAfter?.TotalMilliseconds ?? 0:N0} ms.",
        exception);
    Environment.ExitCode = 1;
}
catch (CosmosException exception)
{
    ReportCosmosError("Azure Cosmos DB operation failed.", exception);
    Environment.ExitCode = 1;
}
catch (InvalidOperationException exception)
{
    Console.Error.WriteLine($"Configuration error: {exception.Message}");
    Environment.ExitCode = 1;
}
catch (UriFormatException exception)
{
    Console.Error.WriteLine($"Invalid Cosmos DB endpoint: {exception.Message}");
    Environment.ExitCode = 1;
}

static void EnsureLocalEmulatorConnection(string connectionString)
{
    string? endpoint = connectionString
        .Split(';', StringSplitOptions.RemoveEmptyEntries)
        .Select(part => part.Split('=', 2))
        .Where(parts => parts.Length == 2)
        .Where(parts => parts[0].Trim().Equals(
            "AccountEndpoint",
            StringComparison.OrdinalIgnoreCase))
        .Select(parts => parts[1].Trim())
        .SingleOrDefault();

    if (endpoint is null)
    {
        throw new InvalidOperationException(
            "COSMOS_CONNECTION_STRING must contain AccountEndpoint.");
    }

    Uri endpointUri = new(endpoint, UriKind.Absolute);
    if (!endpointUri.IsLoopback)
    {
        throw new InvalidOperationException(
            "This sample is restricted to a local Azure Cosmos DB Emulator endpoint.");
    }
}

static void ReportCosmosError(string message, CosmosException exception)
{
    Console.Error.WriteLine(message);
    Console.Error.WriteLine(
        $"Status: {(int)exception.StatusCode} ({exception.StatusCode}); " +
        $"substatus: {exception.SubStatusCode}; " +
        $"activity id: {exception.ActivityId}; " +
        $"request charge: {exception.RequestCharge} RUs.");
}

internal sealed class Item
{
    [JsonProperty("id")]
    public required string Id { get; init; }

    [JsonProperty("category")]
    public required string Category { get; init; }

    [JsonProperty("name")]
    public required string Name { get; init; }

    [JsonProperty("quantity")]
    public int Quantity { get; set; }
}
