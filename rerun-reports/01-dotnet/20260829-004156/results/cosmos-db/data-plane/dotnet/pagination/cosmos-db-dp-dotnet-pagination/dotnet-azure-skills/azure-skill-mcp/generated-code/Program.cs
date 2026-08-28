using System.Globalization;
using System.Text.Json;
using Microsoft.Azure.Cosmos;

const string queryText = "SELECT * FROM c WHERE c.category = 'electronics'";
const string stateFileName = "cosmos-query-state.json";

try
{
    string connectionString = GetRequiredEnvironmentVariable("COSMOS_CONNECTION_STRING");
    string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE_ID");
    string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER_ID");
    bool resume = args.Contains("--resume", StringComparer.OrdinalIgnoreCase);

    QueryState state = resume
        ? await LoadStateAsync(stateFileName)
        : new QueryState();

    using CosmosClient client = new(connectionString);
    Container container = client.GetContainer(databaseId, containerId);

    QueryDefinition query = new(queryText);
    QueryRequestOptions options = new()
    {
        // This is an upper bound; Cosmos DB can return fewer items in a page.
        MaxItemCount = 50
    };

    using FeedIterator<CosmosItem> iterator = container.GetItemQueryIterator<CosmosItem>(
        queryDefinition: query,
        continuationToken: state.ContinuationToken,
        requestOptions: options);

    while (iterator.HasMoreResults)
    {
        FeedResponse<CosmosItem> page = await iterator.ReadNextAsync();

        state.PagesProcessed++;
        state.TotalRequestCharge += page.RequestCharge;
        state.ContinuationToken = page.ContinuationToken;

        Console.WriteLine(
            $"Page {state.PagesProcessed}: {page.Count} item(s), " +
            $"{page.RequestCharge.ToString("F2", CultureInfo.InvariantCulture)} RU");

        foreach (CosmosItem item in page)
        {
            Console.WriteLine($"  id={item.Id}, category={item.Category}");
        }

        Console.WriteLine($"Continuation token: {page.ContinuationToken ?? "<none>"}");
        Console.WriteLine(
            $"Total RU: {state.TotalRequestCharge.ToString("F2", CultureInfo.InvariantCulture)}");

        if (page.ContinuationToken is not null)
        {
            await SaveStateAsync(stateFileName, state);
        }
    }

    File.Delete(stateFileName);
    Console.WriteLine(
        $"Completed {state.PagesProcessed} page(s); total RU consumed: " +
        $"{state.TotalRequestCharge.ToString("F2", CultureInfo.InvariantCulture)}");
    return 0;
}
catch (CosmosException exception)
{
    Console.Error.WriteLine(
        $"Cosmos DB query failed ({(int)exception.StatusCode} {exception.StatusCode}). " +
        $"Request charge: {exception.RequestCharge.ToString("F2", CultureInfo.InvariantCulture)} RU. " +
        exception.Message);
    return 1;
}
catch (InvalidOperationException exception)
{
    Console.Error.WriteLine($"Configuration error: {exception.Message}");
    return 2;
}
catch (IOException exception)
{
    Console.Error.WriteLine($"Could not read or write query state: {exception.Message}");
    return 3;
}
catch (JsonException exception)
{
    Console.Error.WriteLine($"The saved query state is not valid JSON: {exception.Message}");
    return 3;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new InvalidOperationException($"Set the {name} environment variable.");
}

static async Task<QueryState> LoadStateAsync(string path)
{
    if (!File.Exists(path))
    {
        throw new FileNotFoundException(
            $"Cannot resume because the query state file does not exist: {path}",
            path);
    }

    await using FileStream stream = File.OpenRead(path);
    QueryState? state = await JsonSerializer.DeserializeAsync<QueryState>(stream);
    return state?.ContinuationToken is not null
        ? state
        : throw new InvalidDataException($"The query state file is invalid: {path}");
}

static async Task SaveStateAsync(string path, QueryState state)
{
    string temporaryPath = path + ".tmp";

    await using (FileStream stream = File.Create(temporaryPath))
    {
        await JsonSerializer.SerializeAsync(stream, state);
    }

    File.Move(temporaryPath, path, overwrite: true);
}

internal sealed class CosmosItem
{
    public string? Id { get; init; }

    public string? Category { get; init; }
}

internal sealed class QueryState
{
    public string? ContinuationToken { get; set; }

    public double TotalRequestCharge { get; set; }

    public int PagesProcessed { get; set; }
}
