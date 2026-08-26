using System.Net;
using Microsoft.Azure.Cosmos;

namespace CosmosCrudSample;

internal static class Program
{
    private const string DatabaseName = "TestDB";
    private const string ContainerName = "Items";
    private const string PartitionKeyPath = "/category";

    public static async Task<int> Main()
    {
        string? connectionString =
            Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING");

        if (string.IsNullOrWhiteSpace(connectionString))
        {
            Console.Error.WriteLine(
                "Set the COSMOS_CONNECTION_STRING environment variable before running the program.");
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
                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);
            Database database = databaseResponse.Database;

            ContainerResponse containerResponse =
                await database.CreateContainerIfNotExistsAsync(
                    id: ContainerName,
                    partitionKeyPath: PartitionKeyPath);
            Container container = containerResponse.Container;

            Console.WriteLine(
                $"Using database '{DatabaseName}' and container '{ContainerName}'.");

            CosmosItem item = new(
                id: Guid.NewGuid().ToString(),
                category: "electronics",
                name: "Wireless keyboard",
                quantity: 10);
            PartitionKey partitionKey = new(item.category);

            ItemResponse<CosmosItem> createResponse =
                await container.CreateItemAsync(item, partitionKey);
            Console.WriteLine(
                $"Created item '{createResponse.Resource.id}' " +
                $"({createResponse.RequestCharge:F2} RU).");

            ItemResponse<CosmosItem> readResponse =
                await container.ReadItemAsync<CosmosItem>(item.id, partitionKey);
            Console.WriteLine(
                $"Read item: {readResponse.Resource.name}, " +
                $"quantity {readResponse.Resource.quantity}.");

            QueryDefinition query = new(
                "SELECT * FROM c WHERE c.category = @category");
            query.WithParameter("@category", "electronics");

            using FeedIterator<CosmosItem> iterator =
                container.GetItemQueryIterator<CosmosItem>(query);

            Console.WriteLine("Electronics items:");
            while (iterator.HasMoreResults)
            {
                FeedResponse<CosmosItem> page = await iterator.ReadNextAsync();
                foreach (CosmosItem result in page)
                {
                    Console.WriteLine(
                        $"- {result.id}: {result.name}, quantity {result.quantity}");
                }
            }

            CosmosItem updatedItem = item with { quantity = 25 };
            ItemResponse<CosmosItem> replaceResponse =
                await container.ReplaceItemAsync(
                    updatedItem,
                    updatedItem.id,
                    partitionKey);
            Console.WriteLine(
                $"Updated quantity to {replaceResponse.Resource.quantity}.");

            await container.DeleteItemAsync<CosmosItem>(item.id, partitionKey);
            Console.WriteLine($"Deleted item '{item.id}'.");

            return 0;
        }
        catch (CosmosException exception)
        {
            Console.Error.WriteLine(
                $"Cosmos DB request failed with status {(int)exception.StatusCode} " +
                $"({exception.StatusCode}).");
            Console.Error.WriteLine($"Message: {exception.Message}");
            Console.Error.WriteLine($"Activity ID: {exception.ActivityId}");
            Console.Error.WriteLine($"Request charge: {exception.RequestCharge:F2} RU");

            if (exception.StatusCode == HttpStatusCode.TooManyRequests &&
                exception.RetryAfter is TimeSpan retryAfter)
            {
                Console.Error.WriteLine(
                    $"Retry after: {retryAfter.TotalMilliseconds:F0} ms");
            }

            return 1;
        }
    }
}

internal sealed record CosmosItem(
    string id,
    string category,
    string name,
    int quantity);
