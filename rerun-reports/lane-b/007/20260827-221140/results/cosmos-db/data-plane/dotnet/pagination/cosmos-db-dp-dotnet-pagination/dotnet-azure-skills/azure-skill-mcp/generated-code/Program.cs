using System.Text.Json;
using Microsoft.Azure.Cosmos;

const string QueryText =
    "SELECT * FROM c WHERE c.category = 'electronics'";

string endpoint = GetRequiredEnvironmentVariable("COSMOS_ENDPOINT");
string key = GetRequiredEnvironmentVariable("COSMOS_KEY");
string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE_ID");
string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER_ID");
string tokenFile = GetArgumentValue(args, "--continuation-token-file")
    ?? ".cosmos-continuation-token";

using CancellationTokenSource cancellation = new();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellation.Cancel();
};

string? continuationToken = await ReadContinuationTokenAsync(
    tokenFile,
    cancellation.Token);

if (continuationToken is not null)
{
    Console.WriteLine($"Resuming from the token saved in {tokenFile}.");
}

using CosmosClient client = new(endpoint, key);
Container container = client.GetContainer(databaseId, containerId);

QueryDefinition query = new(QueryText);
QueryRequestOptions options = new()
{
    // This is the maximum number requested per response, not a guaranteed count.
    MaxItemCount = 50
};

using FeedIterator<JsonElement> iterator =
    container.GetItemQueryIterator<JsonElement>(
        queryDefinition: query,
        continuationToken: continuationToken,
        requestOptions: options);

double totalRequestUnits = 0;
int pageNumber = 0;

while (iterator.HasMoreResults)
{
    FeedResponse<JsonElement> page =
        await iterator.ReadNextAsync(cancellation.Token);

    pageNumber++;
    totalRequestUnits += page.RequestCharge;

    Console.WriteLine(
        $"Page {pageNumber}: {page.Count} item(s), " +
        $"{page.RequestCharge:F2} RU");

    foreach (JsonElement item in page)
    {
        Console.WriteLine(item.GetRawText());
    }

    continuationToken = page.ContinuationToken;
    Console.WriteLine(
        $"Continuation token: {continuationToken ?? "<none>"}");

    if (continuationToken is not null)
    {
        await File.WriteAllTextAsync(
            tokenFile,
            continuationToken,
            cancellation.Token);
    }
}

if (File.Exists(tokenFile))
{
    File.Delete(tokenFile);
}

Console.WriteLine(
    $"Completed {pageNumber} page(s). Total RU consumed: " +
    $"{totalRequestUnits:F2}");

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

static string? GetArgumentValue(string[] arguments, string option)
{
    for (int index = 0; index < arguments.Length; index++)
    {
        if (arguments[index] != option)
        {
            continue;
        }

        if (index + 1 >= arguments.Length ||
            string.IsNullOrWhiteSpace(arguments[index + 1]))
        {
            throw new ArgumentException(
                $"Option '{option}' requires a file path.");
        }

        return arguments[index + 1];
    }

    return null;
}

static async Task<string?> ReadContinuationTokenAsync(
    string path,
    CancellationToken cancellationToken)
{
    if (!File.Exists(path))
    {
        return null;
    }

    string token = await File.ReadAllTextAsync(path, cancellationToken);
    return string.IsNullOrWhiteSpace(token) ? null : token;
}
