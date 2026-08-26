# Cosmos DB query pagination with .NET

This console app uses `Microsoft.Azure.Cosmos` v3 to run:

```sql
SELECT * FROM c WHERE c.category = 'electronics'
```

## Configuration

Set the following environment variables in PowerShell:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_DATABASE = "<database-name>"
$env:COSMOS_CONTAINER = "<container-name>"
dotnet run
```

By default, authentication uses `DefaultAzureCredential`. The signed-in
identity needs a Cosmos DB data-plane role that permits reading items and
executing queries. For a local Cosmos DB emulator only, set `COSMOS_KEY` from
your local emulator configuration.

The app saves each non-empty continuation token to
`continuation-token.txt`. To demonstrate stopping and resuming, process one
page, then run the app again:

```powershell
$env:MAX_PAGES = "1"
dotnet run

Remove-Item Env:MAX_PAGES
dotnet run
```

The second run automatically loads the saved token. You can instead provide
a token directly with `COSMOS_CONTINUATION_TOKEN`; that value takes precedence
over the token file. The file is removed when the query completes.

## Pagination details

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB for at most 50 items in
each query page. It is a maximum/target rather than a guarantee: a page can
contain fewer items because of response-size limits, throttling, query
execution, or an empty cross-partition continuation.

`FeedIterator<T>` is the direct pagination API. `HasMoreResults`,
`ReadNextAsync()`, `FeedResponse<T>.ContinuationToken`, and
`FeedResponse<T>.RequestCharge` expose page boundaries, checkpoints, and RU
charges explicitly.

LINQ starts with `container.GetItemLinqQueryable<T>()` and lets the SDK
translate supported C# expressions into Cosmos SQL. The query is not executed
until it is enumerated. For asynchronous page-by-page processing, convert the
LINQ query with `ToFeedIterator()` and use the same `ReadNextAsync()` loop.
LINQ is useful for type-safe query composition, while a SQL
`QueryDefinition` is preferable when the exact SQL text matters or the query
uses features that the LINQ provider cannot translate. Do not use synchronous
LINQ enumeration for this pagination pattern.

## References

- [Query performance tips: tune page size](https://learn.microsoft.com/azure/cosmos-db/performance-tips-query-sdk?pivots=programming-language-csharp#tune-the-page-size)
- [Get query request charge](https://learn.microsoft.com/azure/cosmos-db/query-metrics-performance#get-the-query-request-charge)
- [.NET SDK v3 item and query operations](https://learn.microsoft.com/azure/cosmos-db/migrate-dotnet-v3#item-and-query-operations)
