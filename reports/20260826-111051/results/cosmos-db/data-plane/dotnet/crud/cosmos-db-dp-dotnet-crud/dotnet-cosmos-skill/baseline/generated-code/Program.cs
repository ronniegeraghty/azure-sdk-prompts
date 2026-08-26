using System.Net;
using Microsoft.Azure.Cosmos;
using Newtonsoft.Json.Linq;

internal static class Program
{
    private const string DatabaseName = "TestDB";
    private const string ContainerName = "Items";
    private const string PartitionKeyPath = "/category";

    public static async Task<int> Main()
    {
        try
        {
            string connectionString =
                Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING")
                ?? throw new InvalidOperationException(
                    "Set COSMOS_CONNECTION_STRING to the Azure Cosmos DB Emulator connection string.");

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
                    new ContainerProperties(ContainerName, PartitionKeyPath));
            Container container = containerResponse.Container;

            const string itemId = "item-001";
            const string category = "electronics";
            PartitionKey partitionKey = new(category);

            JObject item = JObject.FromObject(new
            {
                id = itemId,
                category,
                name = "Wireless Mouse",
                quantity = 10
            });

            ItemResponse<JObject> createResponse =
                await container.CreateItemAsync(item, partitionKey);
            Console.WriteLine(
                $"Created {createResponse.Resource["id"]} " +
                $"(RU: {createResponse.RequestCharge:F2})");

            ItemResponse<JObject> readResponse =
                await container.ReadItemAsync<JObject>(itemId, partitionKey);
            Console.WriteLine($"Read: {readResponse.Resource}");

            QueryDefinition queryDefinition = new(
                "SELECT * FROM c WHERE c.category = @category");
            queryDefinition.WithParameter("@category", category);

            using FeedIterator<JObject> query =
                container.GetItemQueryIterator<JObject>(
                    queryDefinition,
                    requestOptions: new QueryRequestOptions
                    {
                        PartitionKey = partitionKey
                    });

            while (query.HasMoreResults)
            {
                FeedResponse<JObject> page = await query.ReadNextAsync();
                foreach (JObject queryItem in page)
                {
                    Console.WriteLine($"Query result: {queryItem}");
                }
            }

            item["quantity"] = 25;
            ItemResponse<JObject> replaceResponse =
                await container.ReplaceItemAsync(item, itemId, partitionKey);
            Console.WriteLine(
                $"Updated quantity to {replaceResponse.Resource["quantity"]} " +
                $"(RU: {replaceResponse.RequestCharge:F2})");

            ItemResponse<JObject> deleteResponse =
                await container.DeleteItemAsync<JObject>(itemId, partitionKey);
            Console.WriteLine(
                $"Deleted {itemId} (status: {deleteResponse.StatusCode}, " +
                $"RU: {deleteResponse.RequestCharge:F2})");

            return 0;
        }
        catch (CosmosException exception)
        {
            Console.Error.WriteLine(
                $"Cosmos DB request failed: {exception.StatusCode} " +
                $"({(int)exception.StatusCode}), substatus {exception.SubStatusCode}.");
            Console.Error.WriteLine($"Activity ID: {exception.ActivityId}");
            Console.Error.WriteLine($"Request charge: {exception.RequestCharge:F2} RU");
            Console.Error.WriteLine(exception.Message);

            if (exception.StatusCode == HttpStatusCode.Conflict)
            {
                Console.Error.WriteLine(
                    "The item already exists. Delete item-001 or use a different id.");
            }

            return 1;
        }
        catch (InvalidOperationException exception)
        {
            Console.Error.WriteLine($"Configuration error: {exception.Message}");
            return 2;
        }
    }
}
