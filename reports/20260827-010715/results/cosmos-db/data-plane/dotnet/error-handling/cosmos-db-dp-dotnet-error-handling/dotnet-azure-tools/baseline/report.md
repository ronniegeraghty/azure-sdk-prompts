# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 147.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 147.0s |
| Timestamp | 2026-08-26T17:18:14Z |
| Files Generated | 0 |
| Event Count | 3240 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 122.9s |
| Review | 23.4s |
| **Total** | **147.0s** |

## Configuration

- **name:** dotnet-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=97042 out=2712 |
| Turn Count | 4 |

## Prompt Sent

```
How do I properly handle errors when working with Azure Cosmos DB in .NET?
I'm particularly concerned about throttling — when my application exceeds
the provisioned Request Units (RU/s), I get 429 errors. Show me how to
catch CosmosException, extract the status code and retry-after header,
and implement proper retry logic. Also cover handling 404 (item not found)
and 409 (conflict) responses. Use the Microsoft.Azure.Cosmos SDK v3.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Microsoft.Azure.Cosmos SDK v3 CosmosException StatusCode RetryAfter 429 MaxRetryAttemptsOnRateLimitedRequests documentation"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Microsoft Azure Cosmos DB .NET SDK v3 handle 404 NotFound 409 Conflict CosmosException official documentation"}
```

### 🔧 3. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests?view=azure-dotnet"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretrywaittimeonratelimitedrequests?view=azure-dotnet"}
```

### 🔧 6. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-cosmos-dotnet-v3 \"x-ms-retry-after-ms\" \"exception.Headers\" language:C#"}
```

## Copilot Response

I’ll verify the SDK v3 exception and retry option semantics against current Microsoft documentation, then provide a complete pattern.

The v3 SDK throws `CosmosException` for service responses. Its `StatusCode` exposes the HTTP status, while `RetryAfter` exposes the delay from the `x-ms-retry-after-ms` response header.

## Prefer the SDK’s built-in 429 retry policy

The SDK automatically retries `429 Too Many Requests` responses and honors the server-provided retry interval.

```csharp
using Microsoft.Azure.Cosmos;

CosmosClientOptions options = new()
{
    // Defaults are 9 retries and 30 seconds total wait time.
    MaxRetryAttemptsOnRateLimitedRequests = 9,
    MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30)
};

CosmosClient client = new(
    accountEndpoint: Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")!,
    authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")!,
    clientOptions: options);

Container container = client.GetContainer("my-database", "my-container");
```

If the retry count or cumulative wait limit is exhausted, the SDK surfaces the final 429 as a `CosmosException`.

```csharp
try
{
    ItemResponse<MyItem> response = await container.ReadItemAsync<MyItem>(
        id: itemId,
        partitionKey: new PartitionKey(partitionKeyValue),
        cancellationToken: cancellationToken);
}
catch (CosmosException ex)
{
    int statusCode = (int)ex.StatusCode;
    TimeSpan? retryAfter = ex.RetryAfter;

    Console.Error.WriteLine(
        $"Cosmos DB request failed. " +
        $"Status={statusCode}, " +
        $"SubStatus={ex.SubStatusCode}, " +
        $"RetryAfter={retryAfter}, " +
        $"ActivityId={ex.ActivityId}, " +
        $"RequestCharge={ex.RequestCharge}");

    throw;
}
```

`RetryAfter` is the supported way to read the retry-after header; it has already been parsed into a nullable `TimeSpan`.

## Custom retry policy

Use a custom policy only when you need an application-level retry budget or specialized logging. To avoid multiplying delays, disable the SDK’s automatic 429 retries when using this helper:

```csharp
CosmosClientOptions options = new()
{
    MaxRetryAttemptsOnRateLimitedRequests = 0
};
```

A bounded, cancellation-aware retry implementation:

```csharp
using System.Diagnostics;
using System.Net;
using Microsoft.Azure.Cosmos;

public static class CosmosRetry
{
    public static async Task<T> ExecuteAsync<T>(
        Func<CancellationToken, Task<T>> operation,
        int maxAttempts,
        TimeSpan maxTotalDelay,
        CancellationToken cancellationToken = default)
    {
        ArgumentNullException.ThrowIfNull(operation);

        if (maxAttempts < 1)
            throw new ArgumentOutOfRangeException(nameof(maxAttempts));

        if (maxTotalDelay < TimeSpan.Zero)
            throw new ArgumentOutOfRangeException(nameof(maxTotalDelay));

        Stopwatch stopwatch = Stopwatch.StartNew();

        for (int attempt = 1; ; attempt++)
        {
            try
            {
                return await operation(cancellationToken);
            }
            catch (CosmosException ex) when (
                ex.StatusCode == HttpStatusCode.TooManyRequests &&
                attempt < maxAttempts)
            {
                // RetryAfter comes from the x-ms-retry-after-ms response header.
                TimeSpan delay = ex.RetryAfter ??
                    GetFallbackDelay(attempt);

                TimeSpan remainingBudget = maxTotalDelay - stopwatch.Elapsed;
                if (delay > remainingBudget)
                {
                    // Preserve the original CosmosException instead of hiding it
                    // behind a retry-budget exception.
                    throw;
                }

                Console.Error.WriteLine(
                    $"Cosmos DB throttled request: status={(int)ex.StatusCode}, " +
                    $"attempt={attempt}/{maxAttempts}, " +
                    $"retryAfter={delay.TotalMilliseconds:N0}ms, " +
                    $"activityId={ex.ActivityId}");

                await Task.Delay(delay, cancellationToken);
            }
        }
    }

    private static TimeSpan GetFallbackDelay(int attempt)
    {
        // Used only if Cosmos DB did not provide RetryAfter.
        double exponentialMs = Math.Min(
            100 * Math.Pow(2, attempt - 1),
            5_000);

        return TimeSpan.FromMilliseconds(
            exponentialMs + Random.Shared.Next(0, 100));
    }
}
```

Usage:

```csharp
ItemResponse<MyItem> response = await CosmosRetry.ExecuteAsync(
    ct => container.ReadItemAsync<MyItem>(
        itemId,
        new PartitionKey(partitionKeyValue),
        cancellationToken: ct),
    maxAttempts: 5,
    maxTotalDelay: TimeSpan.FromSeconds(20),
    cancellationToken);
```

The retry delay must never be shorter than `RetryAfter`. Only the fallback delay needs exponential backoff and jitter.

## Handling 404 and 409

Neither response should normally be retried without changing the operation.

```csharp
public static async Task<MyItem?> ReadItemOrNullAsync(
    Container container,
    string id,
    string partitionKeyValue,
    CancellationToken cancellationToken)
{
    try
    {
        ItemResponse<MyItem> response =
            await CosmosRetry.ExecuteAsync(
                ct => container.ReadItemAsync<MyItem>(
                    id,
                    new PartitionKey(partitionKeyValue),
                    cancellationToken: ct),
                maxAttempts: 5,
                maxTotalDelay: TimeSpan.FromSeconds(20),
                cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.NotFound)
    {
        // Also verify that the supplied partition key is correct.
        return null;
    }
}
```

For create conflicts:

```csharp
public static async Task<bool> TryCreateItemAsync(
    Container container,
    MyItem item,
    string partitionKeyValue,
    CancellationToken cancellationToken)
{
    try
    {
        await CosmosRetry.ExecuteAsync(
            ct => container.CreateItemAsync(
                item,
                new PartitionKey(partitionKeyValue),
                cancellationToken: ct),
            maxAttempts: 5,
            maxTotalDelay: TimeSpan.FromSeconds(20),
            cancellationToken);

        return true;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.Conflict)
    {
        // An item with this ID and logical partition key already exists.
        // Retrying the same CreateItemAsync call will continue to fail.
        return false;
    }
}
```

Use `UpsertItemAsync` instead of `CreateItemAsync` only when “create or replace” is actually the desired behavior. For optimistic concurrency failures involving ETags, expect `412 Precondition Failed`, not `409 Conflict`.

References: [CosmosException.RetryAfter](https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter), [rate-limit retry attempts](https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests), and [resilient Cosmos DB SDK applications](https://learn.microsoft.com/en-us/azure/cosmos-db/conceptual-resilient-sdk-applications).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response thoroughly covers CosmosException handling, throttling retries, and 404/409 behavior, but omits successful-response RU tracking and Cosmos diagnostics output.

### Strengths

- Provides both SDK-managed and custom 429 retry approaches.
- Custom retry logic is bounded, cancellation-aware, and honors the server-provided RetryAfter delay.
- Correctly treats 404 and 409 as non-transient outcomes and distinguishes 409 from ETag-related 412 responses.

### Issues

- Does not track RequestCharge from successful ItemResponse values.
- Does not include the Diagnostics string for troubleshooting.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `CosmosException` as the primary exception type: Pass
      - Extracting `StatusCode`, `SubStatusCode`, and `RetryAfter` properties: Pass
      - Handling 429 (TooManyRequests) with retry-after backoff: Pass
      - Handling 404 (NotFound) for missing items or containers: Pass
      - Handling 409 (Conflict) for duplicate items: Pass
      - `CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests` configuration: Pass
      - RU consumption tracking via `RequestCharge` on responses: Fail
      - Diagnostics string for troubleshooting: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.7500** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
