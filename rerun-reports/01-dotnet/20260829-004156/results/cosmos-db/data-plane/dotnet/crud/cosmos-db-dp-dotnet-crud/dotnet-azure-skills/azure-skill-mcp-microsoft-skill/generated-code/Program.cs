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
                "Set the COSMOS_CONNECTION_STRING environment variable before running.");
            return 1;
        }

        try
        {
            using CosmosClient client = new(connectionString);

            DatabaseResponse databaseResponse =
                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);
            Database database = databaseResponse.Database;

            ContainerResponse containerResponse =
                await database.CreateContainerIfNotExistsAsync(
                    new ContainerProperties(ContainerName, PartitionKeyPath),
                    throughput: 400);
            Container container = containerResponse.Container;

            Item item = new(
                id: Guid.NewGuid().ToString(),
                category: "electronics",
                name: "Wireless keyboard",
                quantity: 10);
            PartitionKey partitionKey = new(item.category);

            ItemResponse<Item> createResponse =
                await container.CreateItemAsync(item, partitionKey);
            Console.WriteLine(
                $"Created item {createResponse.Resource.id} " +
                $"(request charge: {createResponse.RequestCharge:F2} RUs).");

            ItemResponse<Item> readResponse =
                await container.ReadItemAsync<Item>(item.id, partitionKey);
            Console.WriteLine(
                $"Read item: {readResponse.Resource.name}, " +
                $"quantity {readResponse.Resource.quantity}.");

            QueryDefinition query = new(
                "SELECT * FROM c WHERE c.category = @category");
            query.WithParameter("@category", "electronics");

            using FeedIterator<Item> iterator = container.GetItemQueryIterator<Item>(
                query,
                requestOptions: new QueryRequestOptions
                {
                    PartitionKey = new PartitionKey("electronics")
                });

            Console.WriteLine("Electronics items:");
            while (iterator.HasMoreResults)
            {
                FeedResponse<Item> page = await iterator.ReadNextAsync();
                foreach (Item result in page)
                {
                    Console.WriteLine(
                        $"- {result.id}: {result.name}, quantity {result.quantity}");
                }
            }

            Item updatedItem = item with { quantity = 25 };
            ItemResponse<Item> replaceResponse =
                await container.ReplaceItemAsync(updatedItem, item.id, partitionKey);
            Console.WriteLine(
                $"Updated quantity to {replaceResponse.Resource.quantity}.");

            await container.DeleteItemAsync<Item>(item.id, partitionKey);
            Console.WriteLine($"Deleted item {item.id}.");

            return 0;
        }
        catch (CosmosException exception)
        {
            Console.Error.WriteLine(
                $"Cosmos DB request failed: {(int)exception.StatusCode} " +
                $"({exception.StatusCode}), substatus {exception.SubStatusCode}, " +
                $"activity ID {exception.ActivityId}.");
            Console.Error.WriteLine(exception.Message);

            if (exception.StatusCode == HttpStatusCode.TooManyRequests &&
                exception.RetryAfter is TimeSpan retryAfter)
            {
                Console.Error.WriteLine(
                    $"Retry after {retryAfter.TotalMilliseconds:F0} ms.");
            }

            return 1;
        }
        catch (ArgumentException exception)
        {
            Console.Error.WriteLine($"Invalid Cosmos DB configuration: {exception.Message}");
            return 1;
        }
    }

    private sealed record Item(
        string id,
        string category,
        string name,
        int quantity);
}
