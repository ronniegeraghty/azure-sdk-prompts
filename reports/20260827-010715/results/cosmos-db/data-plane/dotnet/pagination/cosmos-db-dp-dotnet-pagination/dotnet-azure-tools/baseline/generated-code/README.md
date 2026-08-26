# Cosmos DB paginated query sample

This .NET 8 console program uses `Microsoft.Azure.Cosmos` SDK v3 to execute:

```sql
SELECT * FROM c WHERE c.category = 'electronics'
```

Configure the container through environment variables:

```powershell
$env:COSMOS_ENDPOINT = "https://localhost:8081"
$env:COSMOS_KEY = "<emulator-key>"
$env:COSMOS_DATABASE = "catalog"
$env:COSMOS_CONTAINER = "items"
dotnet run
```

To resume, save the continuation token printed after a successfully processed page
and pass it to the next run:

```powershell
dotnet run -- "<saved-continuation-token>"
```

Alternatively, set `COSMOS_CONTINUATION_TOKEN`. The command-line argument takes
precedence. Continuation tokens should be treated as opaque values and reused with
the same query and compatible query options.

## Pagination behavior

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB to return at most 50 items
per page. A page can contain fewer items because of response-size limits, RU
availability, partition behavior, or query execution limits. `FeedIterator`
explicitly exposes each service response through `ReadNextAsync()`, including its
items, continuation token, and `RequestCharge`. This makes it suitable for
checkpoints and page-by-page RU accounting.

Cosmos LINQ queries begin with `container.GetItemLinqQueryable<T>()` and translate
C# expressions into Cosmos SQL. They are convenient for type-safe query
construction, but enumeration alone hides service-page boundaries. For explicit
pagination, convert the LINQ query with `ToFeedIterator()` and use the same
`HasMoreResults`/`ReadNextAsync()` loop. A SQL `QueryDefinition` is preferable when
the query text is already known or when exact SQL behavior is important.
