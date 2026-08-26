using Azure.Identity;
using Microsoft.Azure.Cosmos;
using Newtonsoft.Json.Linq;

const string QueryText =
    "SELECT * FROM c WHERE c.category = 'electronics'";
const string TokenFileName = "continuation-token.txt";

string endpoint = GetRequiredSetting("COSMOS_ENDPOINT");
string databaseName = GetRequiredSetting("COSMOS_DATABASE");
string containerName = GetRequiredSetting("COSMOS_CONTAINER");
int? maxPages = GetOptionalPositiveInteger("MAX_PAGES");

string? continuationToken = Environment.GetEnvironmentVariable(
    "COSMOS_CONTINUATION_TOKEN");

if (string.IsNullOrWhiteSpace(continuationToken)
    && File.Exists(TokenFileName))
{
    continuationToken = await File.ReadAllTextAsync(TokenFileName);
    Console.WriteLine($"Resuming from token saved in {TokenFileName}.");
}

using CosmosClient client = CreateCosmosClient(endpoint);
Container container = client.GetContainer(databaseName, containerName);

QueryRequestOptions requestOptions = new()
{
    // MaxItemCount is a page-size target, not a guaranteed page size.
    MaxItemCount = 50
};

QueryDefinition query = new(QueryText);
using FeedIterator<JObject> iterator =
    container.GetItemQueryIterator<JObject>(
        query,
        continuationToken,
        requestOptions);

double totalRequestUnits = 0;
int pageNumber = 0;

while (iterator.HasMoreResults
       && (maxPages is null || pageNumber < maxPages))
{
    FeedResponse<JObject> page = await iterator.ReadNextAsync();
    pageNumber++;
    totalRequestUnits += page.RequestCharge;

    Console.WriteLine(
        $"Page {pageNumber}: {page.Count} item(s), "
        + $"{page.RequestCharge:F2} RU");

    foreach (JObject item in page)
    {
        Console.WriteLine(item.ToString());
    }

    string? nextToken = page.ContinuationToken;
    Console.WriteLine(
        $"Continuation token: {nextToken ?? "<none>"}");

    if (string.IsNullOrEmpty(nextToken))
    {
        File.Delete(TokenFileName);
    }
    else
    {
        // Saving after processing the page makes the checkpoint represent
        // the first unprocessed page.
        await File.WriteAllTextAsync(TokenFileName, nextToken);
    }
}

Console.WriteLine($"Total RU consumed: {totalRequestUnits:F2}");

if (iterator.HasMoreResults)
{
    Console.WriteLine(
        $"Stopped after {pageNumber} page(s). Run again to resume "
        + $"from {TokenFileName}.");
}
else
{
    Console.WriteLine("Query completed.");
}

static CosmosClient CreateCosmosClient(string endpoint)
{
    string? key = Environment.GetEnvironmentVariable("COSMOS_KEY");

    if (!string.IsNullOrWhiteSpace(key))
    {
        // Use COSMOS_KEY only for local emulator/development scenarios.
        return new CosmosClient(endpoint, key);
    }

    return new CosmosClient(
        endpoint,
        new DefaultAzureCredential(),
        new CosmosClientOptions
        {
            ApplicationName = "CosmosPaginationSample"
        });
}

static string GetRequiredSetting(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new InvalidOperationException(
            $"Set the {name} environment variable.");
}

static int? GetOptionalPositiveInteger(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    if (string.IsNullOrWhiteSpace(value))
    {
        return null;
    }

    return int.TryParse(value, out int parsed) && parsed > 0
        ? parsed
        : throw new InvalidOperationException(
            $"{name} must be a positive integer.");
}
