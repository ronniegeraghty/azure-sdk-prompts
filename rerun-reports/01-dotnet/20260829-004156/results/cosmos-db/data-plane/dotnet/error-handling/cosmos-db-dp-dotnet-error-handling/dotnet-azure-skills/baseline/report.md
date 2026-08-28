# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 146.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 146.9s |
| Timestamp | 2026-08-28T17:04:26Z |
| Files Generated | 0 |
| Event Count | 3342 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 116.0s |
| Review | 29.3s |
| **Total** | **146.9s** |

## Configuration

- **name:** dotnet-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=42295 out=2320 |
| Turn Count | 2 |

## Tool Calls

### ✅ 1. web_search (via github-mcp-server) — 67693ms

**Input:**

```json
{"query":"Official Microsoft Azure Cosmos DB .NET SDK v3 CosmosException StatusCode RetryAfter 429 throttling retry configuration MaxRetryAttemptsOnRateLimitedRequests MaxRetryWaitTimeOnRateLimitedRequests"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here are example snippets for handling 404 NotFound and 409 Conflict exceptions with the official Microsoft Azure Cosmos DB .NET SDK v3 (`Microsoft.Azure.Cosmos`):\n\n### Handling 404 NotFound (`CosmosException.StatusCode == HttpStatusCode.NotFound`)\n\nA `404 NotFound` is thrown (as `CosmosException`) typically when reading or replacing an item that does not exist or is not in the expected partition. Your code should catch this and take appropriate action, such as reporting “not found” to the user:\n\n```csharp\nusing Microsoft.Azure.Cosmos;\nusing System.Net;\n\ntry\n{\n    ItemResponse<MyType> response = await container.ReadItemAsync<MyType>(id, new PartitionKey(partitionKey));\n    // process response.Resource\n}\ncatch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.NotFound)\n{\n    // Item not found, handle accordingly\n    Console.WriteLine($\"Item not found. Diagnostics: {ex.Diagnostics}\");\n}\n```\n**Note:** The `Diagnostics` property is useful for logging details about the request for troubleshooting【3:1†source】【3:9†source】.\n\n---\n\n### Handling 409 Conflict (`CosmosException.StatusCode == HttpStatusCode.Conflict`)\n\nA `409 Conflict` usually happens if you try to create an item with an `id` that already exists in the same partition (resource already exists):\n\n```csharp\nusing Microsoft.Azure.Cosmos;\nusing System.Net;\n\ntry\n{\n    ItemResponse<MyType> response = await container.CreateItemAsync<MyType>(item, new PartitionKey(item.PartitionKey));\n}\ncatch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.Conflict)\n{\n    // Conflict error (resource with specified id already exists)\n    Console.WriteLine($\"Conflict: Item with id '{item.id}' already exists. Diagnostics: {ex.Diagnostics}\");\n}\n```\n\n---\n\n### Catch-All with Status Code Handling\n\nYou can handle both in a single catch block (good for logging or generic workflows):\n\n```csharp\ntry\n{\n    // Your cosmos
... (truncated)
```

### ✅ 2. web_search (via github-mcp-server) — 65756ms

**Input:**

```json
{"query":"Official Microsoft Azure Cosmos DB .NET SDK v3 handle 404 NotFound 409 Conflict CosmosException examples"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's a comprehensive explanation of handling 429 errors (throttling) and the official retry configuration settings in the Microsoft Azure Cosmos DB .NET SDK v3:\n\n### What is a 429 in Cosmos DB?\nA 429 \"Too Many Requests\" status code in Azure Cosmos DB means you have exceeded the provisioned Request Units (RUs) per second for your container or database. The Cosmos DB service tells the SDK how long to wait before retrying via the `x-ms-retry-after-ms` header.\n\n### Does the .NET SDK v3 Retry Automatically?\n**Yes.** By default, the SDK will automatically retry write, read, and query requests that receive a 429 response, honoring the wait time indicated by the `x-ms-retry-after-ms` header sent from Cosmos DB. This helps to smooth out short bursts of over-capacity operations【3:2†source】【3:3†source】【3:6†source】.\n\n### How Do You Configure the Retry Policy?\nWhen you create your `CosmosClient` via `CosmosClientOptions`, you can control two important parameters:\n\n#### 1. `MaxRetryAttemptsOnRateLimitedRequests`\n- **Definition**: The maximum number of retry attempts for a single request that receives a 429 status.\n- **Default**: 9 attempts.\n- **Usage Example**:\n    ```csharp\n    var options = new CosmosClientOptions\n    {\n        MaxRetryAttemptsOnRateLimitedRequests = 20, // e.g., set to 20\n    };\n    ```\n\n#### 2. `MaxRetryWaitTimeOnRateLimitedRequests`\n- **Definition**: The maximum cumulative time that will be spent retrying a request due to 429 errors.\n- **Default**: 30 seconds.\n- **Usage Example**:\n    ```csharp\n    var options = new CosmosClientOptions\n    {\n        MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromMinutes(2), // e.g., 2 minutes\n    };\n    ```\n\nCombine both in your client configuration:\n```csharp\nvar options = new CosmosClientOptions\n{\n    MaxRetryAttemptsOnRateLimitedRequests = 20,\n    MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromMinutes(2)\n};\n\nv
... (truncated)
```

## Copilot Response

I’ll anchor the retry guidance in the current Cosmos DB .NET SDK v3 behavior and official recommendations.

`CosmosException` exposes the HTTP status through `StatusCode` and the server-requested throttling delay through `RetryAfter`, which corresponds to the `x-ms-retry-after-ms` response header.

The SDK v3 already retries 429 responses by default. Prefer configuring that built-in policy rather than layering another retry policy on top:

```csharp
using Microsoft.Azure.Cosmos;

CosmosClientOptions options = new()
{
    // Maximum retries for a single rate-limited request.
    MaxRetryAttemptsOnRateLimitedRequests = 9,

    // Maximum cumulative time spent retrying that request.
    MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30)
};

CosmosClient client = new(
    accountEndpoint: Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")!,
    authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")!,
    clientOptions: options);

Container container = client.GetContainer("my-database", "my-container");
```

If the retry count or cumulative wait limit is exhausted, the operation throws a `CosmosException` with status 429:

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;

try
{
    ItemResponse<Customer> response =
        await container.ReadItemAsync<Customer>(
            id: customerId,
            partitionKey: new PartitionKey(tenantId),
            cancellationToken: cancellationToken);

    Customer customer = response.Resource;
}
catch (CosmosException ex)
    when (ex.StatusCode == HttpStatusCode.TooManyRequests)
{
    Console.Error.WriteLine(
        $"Cosmos DB throttling persisted after SDK retries. " +
        $"Status={(int)ex.StatusCode}, " +
        $"RetryAfter={ex.RetryAfter.TotalMilliseconds}ms, " +
        $"RequestCharge={ex.RequestCharge}, " +
        $"ActivityId={ex.ActivityId}, " +
        $"Diagnostics={ex.Diagnostics}");

    throw;
}
```

## Explicit retry logic

If you need application-specific limits, logging, or metrics, you can implement the retry loop yourself. To avoid multiplying retries, disable the SDK’s built-in 429 retries for the client used with this helper:

```csharp
CosmosClientOptions options = new()
{
    MaxRetryAttemptsOnRateLimitedRequests = 0
};

CosmosClient client = new(endpoint, key, options);
```

A bounded retry helper should honor `RetryAfter`, support cancellation, add modest jitter, and impose both retry-count and total-wait limits:

```csharp
using System.Diagnostics;
using System.Net;
using Microsoft.Azure.Cosmos;

public static class CosmosRetry
{
    public static async Task<T> ExecuteAsync<T>(
        Func<CancellationToken, Task<T>> operation,
        int maxRetries,
        TimeSpan maxTotalWait,
        CancellationToken cancellationToken = default)
    {
        ArgumentNullException.ThrowIfNull(operation);

        if (maxRetries < 0)
            throw new ArgumentOutOfRangeException(nameof(maxRetries));

        if (maxTotalWait < TimeSpan.Zero)
            throw new ArgumentOutOfRangeException(nameof(maxTotalWait));

        Stopwatch stopwatch = Stopwatch.StartNew();

        for (int retry = 0; ; retry++)
        {
            try
            {
                return await operation(cancellationToken);
            }
            catch (CosmosException ex)
                when (ex.StatusCode == HttpStatusCode.TooManyRequests)
            {
                Console.Error.WriteLine(
                    $"Cosmos DB returned {(int)ex.StatusCode}. " +
                    $"RetryAfter={ex.RetryAfter.TotalMilliseconds}ms, " +
                    $"Attempt={retry + 1}, " +
                    $"RequestCharge={ex.RequestCharge}, " +
                    $"ActivityId={ex.ActivityId}");

                if (retry >= maxRetries)
                    throw;

                // Cosmos DB normally supplies RetryAfter for 429 responses.
                // The fallback protects against an absent or zero value.
                TimeSpan serverDelay = ex.RetryAfter > TimeSpan.Zero
                    ? ex.RetryAfter
                    : TimeSpan.FromMilliseconds(
                        Math.Min(100 * Math.Pow(2, retry), 5_000));

                TimeSpan delay =
                    serverDelay +
                    TimeSpan.FromMilliseconds(Random.Shared.Next(0, 250));

                TimeSpan remaining = maxTotalWait - stopwatch.Elapsed;
                if (remaining <= TimeSpan.Zero || delay > remaining)
                    throw;

                await Task.Delay(delay, cancellationToken);
            }
        }
    }
}
```

Usage:

```csharp
ItemResponse<Customer> response =
    await CosmosRetry.ExecuteAsync(
        ct => container.ReadItemAsync<Customer>(
            customerId,
            new PartitionKey(tenantId),
            cancellationToken: ct),
        maxRetries: 5,
        maxTotalWait: TimeSpan.FromSeconds(20),
        cancellationToken);
```

Do not combine this loop with Polly or the SDK’s default retry behavior unless you deliberately account for the compounded retry count and latency.

## Handling 404 and 409

Use exception filters so each expected response has explicit semantics, and allow unexpected failures to propagate:

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;

public async Task<Customer?> GetCustomerAsync(
    Container container,
    string customerId,
    string tenantId,
    CancellationToken cancellationToken)
{
    try
    {
        ItemResponse<Customer> response =
            await container.ReadItemAsync<Customer>(
                customerId,
                new PartitionKey(tenantId),
                cancellationToken: cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.NotFound)
    {
        // A wrong partition-key value can also produce a 404.
        Console.WriteLine(
            $"Customer '{customerId}' was not found. " +
            $"ActivityId={ex.ActivityId}");

        return null;
    }
}
```

For create conflicts:

```csharp
public async Task<bool> TryCreateCustomerAsync(
    Container container,
    Customer customer,
    CancellationToken cancellationToken)
{
    try
    {
        await container.CreateItemAsync(
            customer,
            new PartitionKey(customer.TenantId),
            cancellationToken: cancellationToken);

        return true;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.Conflict)
    {
        // Usually means an item with the same id and partition key exists.
        Console.WriteLine(
            $"Customer '{customer.Id}' already exists. " +
            $"ActivityId={ex.ActivityId}");

        return false;
    }
}
```

If “create or replace” is the intended behavior, use `UpsertItemAsync` instead of treating every 409 as a retryable error. Neither 404 nor 409 should normally be retried unchanged: a retry will produce the same result unless application state or request data changes.

Useful exception fields to log include:

```csharp
catch (CosmosException ex)
{
    Console.Error.WriteLine(
        $"Status={(int)ex.StatusCode} ({ex.StatusCode}), " +
        $"SubStatus={ex.SubStatusCode}, " +
        $"RetryAfter={ex.RetryAfter}, " +
        $"RequestCharge={ex.RequestCharge}, " +
        $"ActivityId={ex.ActivityId}, " +
        $"Diagnostics={ex.Diagnostics}");

    throw;
}
```

Frequent 429s are not necessarily failures—the SDK can absorb transient throttling—but sustained throttling generally calls for workload concurrency control, better partition-key distribution, query optimization, or increased/autoscale throughput. See Microsoft’s [resilient SDK application guidance](https://learn.microsoft.com/en-us/azure/cosmos-db/conceptual-resilient-sdk-applications) and [`CosmosException` API documentation](https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception).

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The response thoroughly covers CosmosException handling, throttling configuration and retries, expected 404 and 409 outcomes, and diagnostics, but omits successful-response RU tracking.

### Strengths

- Correctly prioritizes the SDK's built-in 429 retry policy and warns against compounded retries.
- Provides a bounded, cancellation-aware custom retry implementation that honors RetryAfter.
- Uses explicit exception filters for 429, 404, and 409 while allowing unexpected failures to propagate.
- Includes StatusCode, SubStatusCode, RetryAfter, ActivityId, and Diagnostics in troubleshooting output.

### Issues

- It does not demonstrate reading ItemResponse<T>.RequestCharge after a successful operation, so RU consumption is only reported for exceptions.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `CosmosException` as the primary exception type: Pass
      - Extracting `StatusCode`, `SubStatusCode`, and `RetryAfter` properties: Pass
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
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
