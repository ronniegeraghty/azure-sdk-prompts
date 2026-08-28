using Azure.Identity;
using Microsoft.Azure.Cosmos;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

const string queryText = "SELECT * FROM c WHERE c.category = 'electronics'";
const string defaultTokenFile = "continuation-token.txt";

string endpoint = GetRequiredEnvironmentVariable("COSMOS_ENDPOINT");
string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE");
string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER");
string tokenFile = Environment.GetEnvironmentVariable("COSMOS_TOKEN_FILE") ?? defaultTokenFile;
string? continuationToken = await GetStartingTokenAsync(args, tokenFile);

using var cancellationSource = new CancellationTokenSource();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellationSource.Cancel();
};

try
{
    var credential = new DefaultAzureCredential();
    using var client = new CosmosClient(
        endpoint,
        credential,
        new CosmosClientOptions { ApplicationName = "CosmosPaginationSample" });

    Container container = client.GetContainer(databaseId, containerId);
    QueryDefinition query = new(queryText);
    QueryRequestOptions requestOptions = new()
    {
        // MaxItemCount is an upper bound; Cosmos DB can return fewer items.
        MaxItemCount = 50
    };

    using FeedIterator<CosmosItem> iterator =
        container.GetItemQueryIterator<CosmosItem>(
            queryDefinition: query,
            continuationToken: continuationToken,
            requestOptions: requestOptions);

    double totalRequestUnits = 0;
    int pageNumber = 0;

    while (iterator.HasMoreResults)
    {
        FeedResponse<CosmosItem> page =
            await iterator.ReadNextAsync(cancellationSource.Token);

        pageNumber++;
        totalRequestUnits += page.RequestCharge;

        Console.WriteLine(
            $"Page {pageNumber}: {page.Count} item(s), {page.RequestCharge:F2} RU");

        foreach (CosmosItem item in page)
        {
            Console.WriteLine(JsonConvert.SerializeObject(item));
        }

        continuationToken = page.ContinuationToken;
        Console.WriteLine(
            $"Continuation token: {continuationToken ?? "<none>"}");

        await SaveTokenAsync(
            tokenFile,
            continuationToken,
            cancellationSource.Token);
    }

    Console.WriteLine($"Total RU consumed: {totalRequestUnits:F2}");
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine(
        $"Query canceled. Resume later with: dotnet run -- --resume \"{tokenFile}\"");
    Environment.ExitCode = 2;
}
catch (CosmosException exception)
{
    Console.Error.WriteLine(
        $"Cosmos DB request failed ({(int)exception.StatusCode}): " +
        $"{exception.Message}; RU: {exception.RequestCharge:F2}");
    Environment.ExitCode = 1;
}
catch (Exception exception)
{
    Console.Error.WriteLine(exception.Message);
    Environment.ExitCode = 1;
}

static async Task<string?> GetStartingTokenAsync(
    string[] arguments,
    string defaultPath)
{
    if (arguments.Length == 0)
    {
        return null;
    }

    if (arguments[0] != "--resume")
    {
        throw new ArgumentException(
            "Usage: dotnet run -- [--resume [continuation-token-file]]");
    }

    string path = arguments.Length switch
    {
        1 => defaultPath,
        2 => arguments[1],
        _ => throw new ArgumentException(
            "Usage: dotnet run -- [--resume [continuation-token-file]]")
    };

    if (!File.Exists(path))
    {
        throw new FileNotFoundException(
            "The continuation token file does not exist.",
            path);
    }

    string token = await File.ReadAllTextAsync(path);
    if (string.IsNullOrWhiteSpace(token))
    {
        throw new InvalidOperationException(
            $"The continuation token file '{path}' is empty.");
    }

    Console.WriteLine($"Resuming from continuation token in '{path}'.");
    return token;
}

static async Task SaveTokenAsync(
    string path,
    string? token,
    CancellationToken cancellationToken)
{
    if (token is null)
    {
        if (File.Exists(path))
        {
            File.Delete(path);
        }

        return;
    }

    await File.WriteAllTextAsync(path, token, cancellationToken);
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException(
            $"Set the required environment variable '{name}'.")
        : value;
}

internal sealed class CosmosItem
{
    [JsonProperty("id")]
    public string? Id { get; init; }

    [JsonProperty("category")]
    public string? Category { get; init; }

    [JsonExtensionData]
    public IDictionary<string, JToken>? AdditionalProperties { get; init; }
}
