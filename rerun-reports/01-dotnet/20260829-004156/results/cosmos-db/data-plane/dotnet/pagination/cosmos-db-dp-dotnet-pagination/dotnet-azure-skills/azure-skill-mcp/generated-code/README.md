# Cosmos DB SDK v3 pagination sample

This console app runs:

```sql
SELECT * FROM c WHERE c.category = 'electronics'
```

It reads results with `FeedIterator<T>`, prints each page's continuation
token, and accumulates `FeedResponse<T>.RequestCharge`. After each non-final
page it atomically saves the token, total RU charge, and page count in
`cosmos-query-state.json`.

## Configuration and usage

Set these variables to a local Cosmos DB emulator connection string and
database/container names:

```powershell
$env:COSMOS_CONNECTION_STRING = "<local-emulator-connection-string>"
$env:COSMOS_DATABASE_ID = "<database-id>"
$env:COSMOS_CONTAINER_ID = "<container-id>"
dotnet run
```

If execution stops after a state file has been written, restart from its saved
continuation token and RU total:

```powershell
dotnet run -- --resume
```

The state file is removed after the query completes. Continuation tokens are
opaque: do not modify them, and resume with the same query and SDK version.

## `MaxItemCount`

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB to return at most 50
items per query execution. It is an upper bound, not a guaranteed page size.
A page can contain fewer items, or occasionally no items, because of response
size, available RUs, execution time, and query-engine behavior. Keep reading
while `FeedIterator.HasMoreResults` is `true`.

## `FeedIterator` compared with LINQ

`GetItemQueryIterator<T>` accepts Cosmos SQL directly and exposes explicit,
asynchronous page boundaries. Each `ReadNextAsync` returns a
`FeedResponse<T>` containing the page's continuation token, request charge,
diagnostics, and other response metadata. This makes it the clearest option
when an application must checkpoint and resume a query.

`GetItemLinqQueryable<T>` provides deferred, strongly typed query composition.
The SDK translates supported LINQ operators into Cosmos SQL; unsupported .NET
expressions cannot run server-side. For asynchronous page-by-page execution,
convert the query with `.ToFeedIterator()` and then use the same
`HasMoreResults`/`ReadNextAsync` loop. LINQ is useful for type-safe query
construction, but it does not remove pagination and its translation should be
reviewed for correctness and efficiency.

## References

- [Pagination in Cosmos DB queries](https://learn.microsoft.com/azure/cosmos-db/nosql/query/pagination)
- [Tune query page size in the .NET SDK](https://learn.microsoft.com/azure/cosmos-db/performance-tips-query-sdk?pivots=programming-language-csharp#tune-the-page-size)
- [Get query request charge](https://learn.microsoft.com/azure/cosmos-db/query-metrics-performance#get-the-query-request-charge)
