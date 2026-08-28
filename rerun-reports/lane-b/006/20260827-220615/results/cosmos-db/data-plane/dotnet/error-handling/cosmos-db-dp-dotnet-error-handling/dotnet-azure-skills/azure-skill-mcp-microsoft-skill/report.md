# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 320.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 320.9s |
| Timestamp | 2026-08-27T14:06:15Z |
| Files Generated | 0 |
| Event Count | 3280 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 284.4s |
| Review | 35.3s |
| **Total** | **320.9s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=144019 out=2643 |
| Turn Count | 4 |

## Tool Calls

### ❌ 1. skill — 2364ms

**Input:**

```json
{"skill":"azure-cosmos-db-dotnet"}
```

**Error:**

```
Skill not found: azure-cosmos-db-dotnet
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 184437ms

**Input:**

```json
{"intent":"Provide .NET code and guidance for handling Microsoft.Azure.Cosmos SDK v3 errors, especially 429 throttling retries, 404 not found, and 409 conflict.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 184440ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos SDK v3 CosmosException properties, RetryAfter, 429 automatic retry configuration, 404, and 409 handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. web_search (via github-mcp-server) — 34066ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Cosmos DB .NET SDK v3 CosmosException RetryAfter 429 MaxRetryAttemptsOnRateLimitedRequests MaxRetryWaitTimeOnRateLimitedRequests 404 409"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s a detailed explanation of the relevant Azure Cosmos DB .NET SDK v3 concepts you mentioned, with official documentation references:\n\n---\n\n### CosmosException and RetryAfter (429)\n- **CosmosException** is thrown by the Azure Cosmos DB SDK for documented error cases, including throttling (HTTP status code 429).\n- The **RetryAfter** property on `CosmosException` tells you how long the client should wait before retrying the operation. Its type is `TimeSpan?`【3:5†source】【3:0†source】.\n    ```csharp\n    catch (CosmosException ex) {\n        TimeSpan? retryAfter = ex.RetryAfter;\n        // Handle accordingly\n    }\n    ```\n\n- **429 status code** (\"Request rate is large\") means your client is being rate limited because RU/s (Request Units per second) are exceeded. The SDK respects the server's suggested retry interval, retrieved from the exception's `RetryAfter` property【3:6†source】【3:7†source】.\n\n---\n\n### Setting Retry Behavior:  \n#### 1. MaxRetryAttemptsOnRateLimitedRequests\n- Property in `CosmosClientOptions` that controls how many times the SDK will automatically retry requests that receive a 429 error.\n- Default: 9 retries (which means 10 total attempts including the initial one).\n- If set to 0, there will be no automatic retry for 429, and you must handle the exception yourself.\n    ```csharp\n    var options = new CosmosClientOptions\n    {\n        MaxRetryAttemptsOnRateLimitedRequests = 5 // example\n    };\n    ```\n- [See official docs here.](https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests)【3:2†source】\n\n#### 2. MaxRetryWaitTimeOnRateLimitedRequests\n- Also a property of `CosmosClientOptions`.\n- Sets the maximum cumulative wait time for all retries due to rate limiting.\n- Default: 30 seconds. If cumulative wait time for retries exceeds this, the SDK stops retrying and returns the error.\
... (truncated)
```

## Copilot Response

I’ll ground the example in the SDK v3 retry behavior so it doesn’t accidentally layer unsafe retries over the client’s built-in 429 handling.

The Azure documentation routers timed out, so I’m switching to direct first-party documentation lookup before finalizing the sample.

`Microsoft.Azure.Cosmos` v3 throws `CosmosException` for service responses. Its `StatusCode` contains the HTTP status, and `RetryAfter` exposes the server-provided retry delay.

### Prefer the SDK’s built-in 429 retry policy

The SDK automatically retries 429 responses and honors the service’s retry-after value:

```csharp
using Microsoft.Azure.Cosmos;

CosmosClient client = new(
    connectionString: Environment.GetEnvironmentVariable(
        "COSMOS_CONNECTION_STRING")
        ?? throw new InvalidOperationException(
            "COSMOS_CONNECTION_STRING is not configured."),
    clientOptions: new CosmosClientOptions
    {
        // Defaults are 9 retries and 30 seconds cumulative wait.
        MaxRetryAttemptsOnRateLimitedRequests = 9,
        MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30)
    });

Container container = client.GetContainer("application", "items");
```

A 429 reaches your `catch` block only after this built-in retry policy has been exhausted:

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;

try
{
    ItemResponse<MyItem> response =
        await container.ReadItemAsync<MyItem>(
            id: itemId,
            partitionKey: new PartitionKey(partitionKey),
            cancellationToken: cancellationToken);

    MyItem item = response.Resource;
}
catch (CosmosException ex)
    when (ex.StatusCode == HttpStatusCode.TooManyRequests)
{
    int statusCode = (int)ex.StatusCode;       // 429
    TimeSpan? retryAfter = ex.RetryAfter;      // Retry-After from Cosmos DB
    string activityId = ex.ActivityId;
    double requestCharge = ex.RequestCharge;

    logger.LogWarning(
        "Cosmos DB throttled the request. Status={StatusCode}, " +
        "RetryAfter={RetryAfter}, ActivityId={ActivityId}, RU={RequestCharge}, " +
        "Diagnostics={Diagnostics}",
        statusCode,
        retryAfter,
        activityId,
        requestCharge,
        ex.Diagnostics);

    throw;
}
```

Do not add another general retry layer around this configuration without considering the combined maximum attempts and wait time.

### Manual bounded 429 retry policy

If application-level retry control is required, disable the SDK’s automatic 429 retries to avoid multiplying retry attempts:

```csharp
CosmosClient client = new(
    connectionString,
    new CosmosClientOptions
    {
        MaxRetryAttemptsOnRateLimitedRequests = 0
    });
```

Then use a bounded policy that honors `RetryAfter`, supports cancellation, and limits cumulative waiting:

```csharp
using System.Diagnostics;
using System.Net;
using Microsoft.Azure.Cosmos;

static async Task<T> ExecuteWithCosmosRetryAsync<T>(
    Func<CancellationToken, Task<T>> operation,
    int maxRetries,
    TimeSpan maxTotalWait,
    ILogger logger,
    CancellationToken cancellationToken)
{
    Stopwatch stopwatch = Stopwatch.StartNew();

    for (int retry = 0; ; retry++)
    {
        try
        {
            return await operation(cancellationToken);
        }
        catch (CosmosException ex)
            when (ex.StatusCode == HttpStatusCode.TooManyRequests
                  && retry < maxRetries)
        {
            TimeSpan remaining = maxTotalWait - stopwatch.Elapsed;
            if (remaining <= TimeSpan.Zero)
            {
                throw;
            }

            // Normally supplied for 429 responses. The fallback protects
            // against a missing or invalid server delay.
            TimeSpan delay = ex.RetryAfter is { } retryAfter
                             && retryAfter > TimeSpan.Zero
                ? retryAfter
                : TimeSpan.FromSeconds(
                    Math.Min(Math.Pow(2, retry), 30));

            if (delay > remaining)
            {
                throw;
            }

            logger.LogWarning(
                "Cosmos DB returned {StatusCode}. Retry {Retry}/{MaxRetries} " +
                "after {Delay}. ActivityId={ActivityId}, RU={RequestCharge}",
                (int)ex.StatusCode,
                retry + 1,
                maxRetries,
                delay,
                ex.ActivityId,
                ex.RequestCharge);

            await Task.Delay(delay, cancellationToken);
        }
    }
}
```

Usage:

```csharp
ItemResponse<MyItem> response =
    await ExecuteWithCosmosRetryAsync(
        ct => container.ReadItemAsync<MyItem>(
            itemId,
            new PartitionKey(partitionKey),
            cancellationToken: ct),
        maxRetries: 5,
        maxTotalWait: TimeSpan.FromSeconds(30),
        logger,
        cancellationToken);
```

Only transient responses such as 429 should be retried. A persistent high 429 rate generally requires reducing request concurrency, optimizing RU usage, or increasing/autoscaling throughput.

### Handle 404 as an expected missing item

A point read for a nonexistent item throws a 404. Convert it to `null` when absence is expected:

```csharp
static async Task<T?> ReadOrNullAsync<T>(
    Container container,
    string id,
    string partitionKey,
    CancellationToken cancellationToken)
    where T : class
{
    try
    {
        ItemResponse<T> response = await container.ReadItemAsync<T>(
            id,
            new PartitionKey(partitionKey),
            cancellationToken: cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.NotFound)
    {
        return null;
    }
}
```

Make sure the partition key is correct: an incorrect partition-key value can also produce a 404 even when an item with that ID exists elsewhere.

### Handle 409 conflicts explicitly

A create can return 409 when the item ID or a unique-key value already exists:

```csharp
try
{
    ItemResponse<MyItem> response = await container.CreateItemAsync(
        item,
        new PartitionKey(item.PartitionKey),
        cancellationToken: cancellationToken);
}
catch (CosmosException ex)
    when (ex.StatusCode == HttpStatusCode.Conflict)
{
    logger.LogInformation(
        "Item creation conflicted with an existing item. " +
        "Status={StatusCode}, ActivityId={ActivityId}, Diagnostics={Diagnostics}",
        (int)ex.StatusCode,
        ex.ActivityId,
        ex.Diagnostics);

    // Translate to an application-specific result or exception.
    throw new ItemAlreadyExistsException(item.Id, ex);
}
```

Do not automatically retry a 409: retrying the same create normally produces the same conflict. Instead, return an “already exists” result, read the existing item if appropriate, use `UpsertItemAsync` when create-or-replace semantics are intended, or resolve the unique-key collision.

Official references: [CosmosException](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosexception), [RetryAfter](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter), and [troubleshooting 429 responses](https://learn.microsoft.com/azure/cosmos-db/troubleshoot-request-rate-too-large).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response provides strong CosmosException handling and robust throttling guidance, but omits SubStatusCode and successful-response RU tracking.

### Strengths

- Uses CosmosException consistently for service errors.
- Provides both SDK-managed and bounded manual 429 retry approaches.
- Handles 404 and 409 with appropriate non-retry semantics.
- Includes diagnostics and activity IDs for troubleshooting.

### Issues

- SubStatusCode is not extracted or logged.
- RequestCharge is only shown on exceptions, not on successful response objects.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `CosmosException` as the primary exception type: Pass
      - Extracting `StatusCode`, `SubStatusCode`, and `RetryAfter` properties: Fail
      - Handling 429 (TooManyRequests) with retry-after backoff: Pass
      - Handling 404 (NotFound) for missing items or containers: Pass
      - Handling 409 (Conflict) for duplicate items: Pass
      - `CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests` configuration: Pass
      - RU consumption tracking via `RequestCharge` on responses: Fail
      - Diagnostics string for troubleshooting: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.7500** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
