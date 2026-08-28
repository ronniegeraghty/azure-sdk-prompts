# Cosmos DB query pagination with .NET

This sample uses `Microsoft.Azure.Cosmos` SDK v3 to execute:

```sql
SELECT * FROM c WHERE c.category = 'electronics'
```

It reads results with `FeedIterator`, requests at most 50 items per page,
prints and persists each continuation token, and accumulates
`FeedResponse.RequestCharge` across all pages.

## Run

Set the connection details for a Cosmos DB account or local Cosmos DB
emulator. Do not store credentials in source control.

```powershell
$env:COSMOS_CONNECTION_STRING = "<connection-string>"
$env:COSMOS_DATABASE_ID = "<database-id>"
$env:COSMOS_CONTAINER_ID = "<container-id>"
dotnet run
```

By default, the current continuation token is stored in
`continuation-token.txt`. To choose another location:

```powershell
dotnet run -- --token-file .\state\electronics.token
```

Stop the program after any completed page and run the same command again to
resume from the saved token. The token file is removed after the final page.
Because a token is saved only after processing its entire page, a failure
during page processing can cause that page to be processed again. Consumers
should therefore make side effects idempotent.

## `MaxItemCount`

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB to return no more than
50 items in each response. It is not an exact page size: a response can
contain fewer items because of response-size limits, available request units,
query execution limits, or because the query has reached its end.

Continuation tokens are opaque SDK/service state. Store and pass them back
unchanged, and avoid exposing them unnecessarily.

## `FeedIterator` and LINQ queries

`GetItemQueryIterator<T>` accepts Cosmos SQL directly and returns a
`FeedIterator<T>`. It is the clearest option when the SQL text is known and
the application needs explicit control over continuation tokens, request
options, page boundaries, and per-page diagnostics such as RU charge.

The LINQ API starts with `GetItemLinqQueryable<T>()` and translates supported
C# expressions into Cosmos SQL. LINQ is useful for typed, composable query
construction, but unsupported expressions can fail translation and the
generated SQL may be less obvious. LINQ queries are not automatically
page-by-page: call `ToFeedIterator()` on the resulting query and then use the
same `HasMoreResults`/`ReadNextAsync` loop. For resumption, supply the saved
token and `QueryRequestOptions` to `GetItemLinqQueryable<T>` before calling
`ToFeedIterator()`.
