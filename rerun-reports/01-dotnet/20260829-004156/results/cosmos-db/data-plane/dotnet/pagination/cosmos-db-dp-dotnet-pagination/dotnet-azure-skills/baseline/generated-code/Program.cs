using Microsoft.Azure.Cosmos;

const string queryText =
    "SELECT * FROM c WHERE c.category = 'electronics'";

string endpoint = GetRequiredEnvironmentVariable("COSMOS_ENDPOINT");
string key = GetRequiredEnvironmentVariable("COSMOS_KEY");
string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE");
string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER");

CommandLineOptions options = CommandLineOptions.Parse(args);
string? continuationToken = options.ContinuationToken;

if (continuationToken is null &&
    options.Resume &&
    File.Exists(options.TokenFile))
{
    continuationToken = await File.ReadAllTextAsync(options.TokenFile);
    Console.WriteLine($"Resuming from token stored in '{options.TokenFile}'.");
}

using CosmosClient client = new(endpoint, key);
Container container = client.GetContainer(databaseId, containerId);

QueryRequestOptions requestOptions = new()
{
    // MaxItemCount is the requested maximum page size. Cosmos DB can return
    // fewer items because of throttling, response-size, or execution limits.
    MaxItemCount = 50
};

QueryDefinition query = new(queryText);
using FeedIterator<dynamic> iterator =
    container.GetItemQueryIterator<dynamic>(
        query,
        continuationToken,
        requestOptions);

double totalRequestCharge = 0;
int pageNumber = 0;
int totalItems = 0;

while (iterator.HasMoreResults)
{
    FeedResponse<dynamic> page = await iterator.ReadNextAsync();
    pageNumber++;
    totalRequestCharge += page.RequestCharge;

    foreach (dynamic item in page)
    {
        // Perform durable business processing before saving the page token.
        Console.WriteLine(item);
        totalItems++;
    }

    continuationToken = page.ContinuationToken;
    Console.WriteLine(
        $"Page {pageNumber}: {page.Count} items, " +
        $"{page.RequestCharge:F2} RU");
    Console.WriteLine(
        $"Continuation token: {continuationToken ?? "<none>"}");

    if (continuationToken is not null)
    {
        await SaveTokenAtomicallyAsync(
            options.TokenFile,
            continuationToken);
    }
    else if (File.Exists(options.TokenFile))
    {
        File.Delete(options.TokenFile);
    }
}

Console.WriteLine(
    $"Finished: {totalItems} items across {pageNumber} pages; " +
    $"total request charge: {totalRequestCharge:F2} RU.");

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new InvalidOperationException(
            $"Required environment variable '{name}' is not set.");
}

static async Task SaveTokenAtomicallyAsync(
    string tokenFile,
    string token)
{
    string fullPath = Path.GetFullPath(tokenFile);
    string? directory = Path.GetDirectoryName(fullPath);
    if (!string.IsNullOrEmpty(directory))
    {
        Directory.CreateDirectory(directory);
    }

    string temporaryFile = fullPath + ".tmp";
    await File.WriteAllTextAsync(temporaryFile, token);
    File.Move(temporaryFile, fullPath, overwrite: true);
}

internal sealed record CommandLineOptions(
    string TokenFile,
    bool Resume,
    string? ContinuationToken)
{
    public static CommandLineOptions Parse(string[] args)
    {
        string tokenFile = "continuation-token.txt";
        bool resume = false;
        string? continuationToken = null;

        for (int index = 0; index < args.Length; index++)
        {
            switch (args[index])
            {
                case "--resume":
                    resume = true;
                    break;
                case "--token-file" when index + 1 < args.Length:
                    tokenFile = args[++index];
                    break;
                case "--continuation-token" when index + 1 < args.Length:
                    continuationToken = args[++index];
                    break;
                default:
                    throw new ArgumentException(
                        $"Unknown or incomplete argument: {args[index]}");
            }
        }

        return new CommandLineOptions(
            tokenFile,
            resume,
            continuationToken);
    }
}
