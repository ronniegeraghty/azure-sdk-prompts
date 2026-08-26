using Microsoft.Azure.Cosmos;
using System.Net;

namespace hyoka_cosmos_db_dp_dotnet_crud_dotnet_azure_tools_with_azure_tools_3348390268;

internal static class Program
{
    private const string DatabaseName = "TestDB";
    private const string ContainerName = "Items";
    private const string Category = "electronics";

    private static async Task<int> Main()
    {
        try
        {
            string connectionString =
                Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING")
                ?? throw new InvalidOperationException(
                    "Set COSMOS_CONNECTION_STRING to a Cosmos DB connection string.");

            using CosmosClient client = new(
                connectionString,
                new CosmosClientOptions
                {
                    ApplicationName = "CosmosCrudSample"
                });

            DatabaseResponse databaseResponse =
                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);
            Database database = databaseResponse.Database;

            ContainerResponse containerResponse =
                await database.CreateContainerIfNotExistsAsync(
                    ContainerName,
                    partitionKeyPath: "/category");
            Container container = containerResponse.Container;

            Console.WriteLine(
                $"Using database '{DatabaseName}' and container '{ContainerName}'.");

            Item item = new()
            {
                id = Guid.NewGuid().ToString(),
                category = Category,
                name = "Wireless headphones",
                quantity = 5
            };
            PartitionKey partitionKey = new(item.category);

            ItemResponse<Item> createResponse =
                await container.CreateItemAsync(item, partitionKey);
            Console.WriteLine(
                $"Created item {createResponse.Resource.id} " +
                $"({createResponse.RequestCharge:F2} RU).");

            ItemResponse<Item> readResponse =
                await container.ReadItemAsync<Item>(item.id, partitionKey);
            Console.WriteLine(
                $"Read item: {readResponse.Resource.name}, " +
                $"quantity {readResponse.Resource.quantity}.");

            QueryDefinition query = new(
                "SELECT * FROM c WHERE c.category = @category");
            query.WithParameter("@category", Category);

            using FeedIterator<Item> results =
                container.GetItemQueryIterator<Item>(
                    query,
                    requestOptions: new QueryRequestOptions
                    {
                        PartitionKey = partitionKey
                    });

            while (results.HasMoreResults)
            {
                FeedResponse<Item> page = await results.ReadNextAsync();

                foreach (Item result in page)
                {
                    Console.WriteLine(
                        $"Query result: {result.id} - {result.name} " +
                        $"(quantity {result.quantity}).");
                }
            }

            item.quantity = 10;
            ItemResponse<Item> replaceResponse =
                await container.ReplaceItemAsync(item, item.id, partitionKey);
            Console.WriteLine(
                $"Updated item quantity to {replaceResponse.Resource.quantity} " +
                $"({replaceResponse.RequestCharge:F2} RU).");

            ItemResponse<Item> deleteResponse =
                await container.DeleteItemAsync<Item>(item.id, partitionKey);
            Console.WriteLine(
                $"Deleted item {item.id} " +
                $"({deleteResponse.RequestCharge:F2} RU).");

            return 0;
        }
        catch (CosmosException exception)
        {
            Console.Error.WriteLine(
                $"Cosmos DB request failed with HTTP {(int)exception.StatusCode} " +
                $"({exception.StatusCode}), substatus {exception.SubStatusCode}.");
            Console.Error.WriteLine($"Activity ID: {exception.ActivityId}");

            if (exception.StatusCode == HttpStatusCode.TooManyRequests
                && exception.RetryAfter is TimeSpan retryAfter)
            {
                Console.Error.WriteLine(
                    $"Retry after: {retryAfter.TotalMilliseconds:F0} ms.");
            }

            Console.Error.WriteLine(exception.Message);
            return 1;
        }
        catch (InvalidOperationException exception)
        {
            Console.Error.WriteLine($"Configuration error: {exception.Message}");
            return 2;
        }
        catch (ArgumentException exception)
        {
            Console.Error.WriteLine($"Invalid Cosmos DB configuration: {exception.Message}");
            return 2;
        }
    }
}

internal sealed class Item
{
    public required string id { get; init; }

    public required string category { get; init; }

    public required string name { get; init; }

    public int quantity { get; set; }
}
