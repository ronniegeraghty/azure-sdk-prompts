# Cosmos DB paginated query

This .NET 8 console application queries Azure Cosmos DB for NoSQL with the
Microsoft.Azure.Cosmos v3 SDK. It reads results one page at a time, prints and
saves each continuation token, and totals the request units (RU) charged for
the pages read.

## Configure and run

Authenticate with Microsoft Entra ID. The identity selected by
`DefaultAzureCredential` needs an appropriate Cosmos DB data-plane role, such
as **Cosmos DB Built-in Data Reader**, scoped as narrowly as possible.

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_DATABASE = "<database-id>"
$env:COSMOS_CONTAINER = "<container-id>"

dotnet run
```

The application writes the latest non-null token to
`continuation-token.txt`. To restart the same query from that saved position:

```powershell
dotnet run -- --resume
```

To use another token file:

```powershell
$env:COSMOS_TOKEN_FILE = "query-state.txt"
dotnet run -- --resume "query-state.txt"
```

Continuation tokens are tied to the query and SDK behavior. Do not reuse one
with a different query. Store tokens securely if they leave the local
application, and treat processing as at-least-once: a crash after processing a
page but before durably saving its token can cause that page to be processed
again.

## Pagination details

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB for at most 50 items in
each response page. It is an upper bound, not a guarantee: a page may contain
fewer items because of response-size, execution-time, throttling, or
cross-partition query constraints. `ReadNextAsync` performs one request for
the next page. Each `FeedResponse.RequestCharge` is added to the running RU
total, and `FeedResponse.ContinuationToken` is the position used to resume.

## FeedIterator compared with LINQ

`GetItemQueryIterator<T>` accepts SQL text or a `QueryDefinition` and directly
returns a `FeedIterator<T>`. It gives explicit asynchronous page boundaries,
continuation tokens, per-page diagnostics and RU charges, so it is the natural
choice for resumable pagination.

The LINQ provider starts with `container.GetItemLinqQueryable<T>()`, translates
supported LINQ expressions into Cosmos DB SQL, and is useful for composing
strongly typed queries. It does not execute merely because an `IQueryable<T>`
was created. For asynchronous paged execution, convert it with
`queryable.ToFeedIterator()` and consume that iterator in the same
`HasMoreResults`/`ReadNextAsync` loop. Enumerating the queryable synchronously
does not provide the same explicit asynchronous page-control pattern.

## References

- [Query items using the .NET SDK](https://learn.microsoft.com/azure/cosmos-db/nosql/how-to-dotnet-query-items)
- [Tune query page size](https://learn.microsoft.com/azure/cosmos-db/performance-tips-query-sdk?pivots=programming-language-csharp#tune-the-page-size)
- [Get query request charge](https://learn.microsoft.com/azure/cosmos-db/query-metrics-performance#get-the-query-request-charge)
- [LINQ to SQL translation](https://learn.microsoft.com/azure/cosmos-db/nosql/query/linq-to-sql)
