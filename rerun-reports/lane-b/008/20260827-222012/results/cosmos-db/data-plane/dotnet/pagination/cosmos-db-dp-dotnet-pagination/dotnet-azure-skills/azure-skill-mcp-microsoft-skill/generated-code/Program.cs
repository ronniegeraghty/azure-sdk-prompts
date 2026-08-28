using Microsoft.Azure.Cosmos;

const string QueryText =
    "SELECT * FROM c WHERE c.category = 'electronics'";
const int PageSize = 50;

string connectionString = GetRequiredEnvironmentVariable(
    "COSMOS_CONNECTION_STRING");
string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE_ID");
string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER_ID");
string tokenFile = GetOption(args, "--token-file")
    ?? "continuation-token.txt";

string? continuationToken = await ReadContinuationTokenAsync(tokenFile);

using CosmosClient client = new(connectionString);
Container container = client.GetContainer(databaseId, containerId);

QueryRequestOptions requestOptions = new()
{
    // This is a maximum, not a guarantee: Cosmos DB can return fewer items.
    MaxItemCount = PageSize
};

FeedIterator<dynamic> iterator = container.GetItemQueryIterator<dynamic>(
    new QueryDefinition(QueryText),
    continuationToken,
    requestOptions);

double totalRequestCharge = 0;
int pageNumber = 0;
int totalItemCount = 0;

if (continuationToken is not null)
{
    Console.WriteLine($"Resuming from token stored in '{tokenFile}'.");
}

while (iterator.HasMoreResults)
{
    FeedResponse<dynamic> page = await iterator.ReadNextAsync();

    pageNumber++;
    totalRequestCharge += page.RequestCharge;
    totalItemCount += page.Count;

    Console.WriteLine(
        $"Page {pageNumber}: {page.Count} items, " +
        $"{page.RequestCharge:F2} RU");

    foreach (dynamic item in page)
    {
        Console.WriteLine(item);
    }

    continuationToken = page.ContinuationToken;
    Console.WriteLine(
        $"Continuation token: {continuationToken ?? "<end>"}");

    // Persist only after the entire page has been processed. If the program
    // stops before this point, restarting safely reprocesses that page.
    await SaveContinuationTokenAsync(tokenFile, continuationToken);
}

Console.WriteLine($"Total items processed: {totalItemCount}");
Console.WriteLine($"Total request charge: {totalRequestCharge:F2} RU");

static string GetRequiredEnvironmentVariable(string name)
{
    return Environment.GetEnvironmentVariable(name)
        ?? throw new InvalidOperationException(
            $"Set the {name} environment variable before running.");
}

static string? GetOption(string[] arguments, string optionName)
{
    for (int index = 0; index < arguments.Length; index++)
    {
        if (arguments[index] != optionName)
        {
            continue;
        }

        if (index + 1 >= arguments.Length ||
            string.IsNullOrWhiteSpace(arguments[index + 1]))
        {
            throw new ArgumentException(
                $"Option {optionName} requires a value.");
        }

        return arguments[index + 1];
    }

    return null;
}

static async Task<string?> ReadContinuationTokenAsync(string path)
{
    if (!File.Exists(path))
    {
        return null;
    }

    string token = await File.ReadAllTextAsync(path);
    return string.IsNullOrWhiteSpace(token) ? null : token;
}

static async Task SaveContinuationTokenAsync(
    string path,
    string? continuationToken)
{
    if (continuationToken is null)
    {
        if (File.Exists(path))
        {
            File.Delete(path);
        }

        return;
    }

    string? directory = Path.GetDirectoryName(Path.GetFullPath(path));
    if (directory is not null)
    {
        Directory.CreateDirectory(directory);
    }

    await File.WriteAllTextAsync(path, continuationToken);
}
