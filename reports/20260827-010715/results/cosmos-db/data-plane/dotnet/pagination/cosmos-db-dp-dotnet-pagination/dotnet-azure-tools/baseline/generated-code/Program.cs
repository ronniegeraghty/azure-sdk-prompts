using Microsoft.Azure.Cosmos;

const string queryText =
    "SELECT * FROM c WHERE c.category = 'electronics'";

string endpoint = GetRequiredEnvironmentVariable("COSMOS_ENDPOINT");
string key = GetRequiredEnvironmentVariable("COSMOS_KEY");
string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE");
string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER");

// A command-line token takes precedence over the environment variable.
string? savedContinuationToken = args.Length > 0
    ? args[0]
    : Environment.GetEnvironmentVariable("COSMOS_CONTINUATION_TOKEN");

using CosmosClient client = new(endpoint, key);
Container container = client.GetContainer(databaseId, containerId);

QueryRequestOptions requestOptions = new()
{
    // MaxItemCount is a page-size ceiling, not a guarantee that every page has 50 items.
    MaxItemCount = 50
};

using FeedIterator<dynamic> iterator = container.GetItemQueryIterator<dynamic>(
    queryDefinition: new QueryDefinition(queryText),
    continuationToken: savedContinuationToken,
    requestOptions: requestOptions);

double totalRequestCharge = 0;
int pageNumber = 0;
int totalItemCount = 0;

while (iterator.HasMoreResults)
{
    FeedResponse<dynamic> page = await iterator.ReadNextAsync();

    pageNumber++;
    totalItemCount += page.Count;
    totalRequestCharge += page.RequestCharge;

    Console.WriteLine(
        $"Page {pageNumber}: {page.Count} item(s), {page.RequestCharge:F2} RU");

    foreach (dynamic item in page)
    {
        Console.WriteLine(item);
    }

    // Persist this value after a successful page if the application may stop here.
    // A null token means the query is complete.
    Console.WriteLine(
        $"Continuation token: {page.ContinuationToken ?? "<none>"}");
}

Console.WriteLine(
    $"Complete: {totalItemCount} item(s) across {pageNumber} page(s), " +
    $"{totalRequestCharge:F2} total RU");

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Required environment variable '{name}' is not set.");
    }

    return value;
}
