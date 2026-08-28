# Cosmos DB pagination with the .NET SDK v3

This console sample queries:

```sql
SELECT * FROM c WHERE c.category = 'electronics'
```

It requests at most 50 items per response, prints each page's continuation
token, persists that token as a checkpoint, and totals the request units (RUs)
reported by all pages processed during the current run.

## Configure and run

Set the connection settings with environment variables. For local development,
use a Cosmos DB emulator endpoint and key; do not commit credentials.

```powershell
$env:COSMOS_ENDPOINT = "https://localhost:8081"
$env:COSMOS_KEY = "<emulator-or-account-key>"
$env:COSMOS_DATABASE_ID = "<database-id>"
$env:COSMOS_CONTAINER_ID = "<container-id>"

dotnet run -- --continuation-token-file ".cosmos-continuation-token"
```

The checkpoint file is optional and defaults to
`.cosmos-continuation-token`. After every non-final page, the program replaces
the file with the continuation token for the next page. If the process stops,
run the same command again to resume. The file is removed after the query
finishes successfully.

A continuation token is tied to the query and SDK behavior that produced it.
Do not reuse it with a different query. Resuming provides cursor-based progress,
but it does not by itself provide exactly-once processing if data changes or the
process fails between handling an item and saving the next token.

## `MaxItemCount`

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB to return no more than 50
items in one response. It is an upper bound, not a promise that every page has
50 items. A page can be smaller because of response-size limits, available RUs,
query execution, or because the result set has ended.

## `FeedIterator` compared with LINQ

`FeedIterator<T>` is the asynchronous page reader used by the SDK. The sample
creates one directly with `GetItemQueryIterator<T>`, which is a good fit when
the SQL text is known and page metadata such as `ContinuationToken` and
`RequestCharge` must be handled explicitly.

LINQ is a query-construction option, not a different pagination mechanism. A
LINQ expression is translated into Cosmos SQL and remains deferred until it is
executed. For asynchronous, page-by-page processing, convert it with
`ToFeedIterator()` and then use the same `HasMoreResults` /
`ReadNextAsync()` loop:

```csharp
using FeedIterator<Item> iterator = container
    .GetItemLinqQueryable<Item>(
        continuationToken: savedToken,
        requestOptions: new QueryRequestOptions { MaxItemCount = 50 })
    .Where(item => item.Category == "electronics")
    .ToFeedIterator();
```

Direct SQL supports the full Cosmos SQL query language and makes the executed
query explicit. LINQ provides compile-time property access and composability,
but only supported expressions can be translated; inspect generated SQL with
`IQueryable.ToString()` when translation details matter. Avoid synchronous
LINQ enumeration for network queries; use `ToFeedIterator()` for asynchronous
execution and access to page-level metadata.
