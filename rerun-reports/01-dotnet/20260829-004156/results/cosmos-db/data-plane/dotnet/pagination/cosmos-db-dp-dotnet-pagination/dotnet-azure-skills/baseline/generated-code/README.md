# Cosmos DB pagination with the .NET SDK v3

This sample executes:

```sql
SELECT * FROM c WHERE c.category = 'electronics'
```

It reads each response through `FeedIterator`, requests at most 50 items per
page, prints and checkpoints the continuation token after processing the page,
and totals the request charge reported by every response.

## Configure and run

Set placeholder configuration through environment variables rather than
putting credentials in source:

```powershell
$env:COSMOS_ENDPOINT = "https://your-account.documents.azure.com:443/"
$env:COSMOS_KEY = "<your-key>"
$env:COSMOS_DATABASE = "catalog"
$env:COSMOS_CONTAINER = "items"

dotnet run
```

The default checkpoint file is `continuation-token.txt`. Resume from it after
an interrupted run:

```powershell
dotnet run -- --resume
```

Choose another checkpoint file or supply a previously saved token directly:

```powershell
dotnet run -- --resume --token-file ".state\electronics.token"
dotnet run -- --continuation-token "<saved-token>"
```

The checkpoint is written only after all items in a page have been processed.
That ordering avoids skipping a page if processing fails. The checkpoint is
removed after the final page. Continuation tokens should be treated as opaque
SDK values and used with the same query and compatible request options.

## `MaxItemCount`

`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB to return no more than
50 items in each response. It is not a minimum or an exact page size: a page
can contain fewer items because of response-size limits, available RUs, or
query execution limits. Pagination must therefore follow `HasMoreResults` and
the returned continuation token rather than assuming that a short page is the
last page.

## `FeedIterator` compared with LINQ

`FeedIterator<T>` is the asynchronous page-oriented API. Each call to
`ReadNextAsync` exposes response metadata such as `ContinuationToken`,
`RequestCharge`, diagnostics, and activity ID, so it is the appropriate API
when an application needs explicit paging, resumability, or RU accounting.

The SDK's LINQ provider starts with `container.GetItemLinqQueryable<T>()` and
translates supported LINQ expressions into Cosmos SQL. LINQ is useful for
type-safe query composition, but the queryable itself does not execute
asynchronously or expose pages. For production iteration, convert it with
`ToFeedIterator()` and consume that iterator in the same page-by-page pattern.
Calling synchronous LINQ enumeration can drain pages implicitly, makes
continuation-token handling less direct, and should be avoided for this use
case.
