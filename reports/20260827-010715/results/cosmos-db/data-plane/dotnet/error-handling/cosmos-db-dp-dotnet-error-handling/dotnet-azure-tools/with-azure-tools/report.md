# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 138.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 138.7s |
| Timestamp | 2026-08-26T17:20:41Z |
| Files Generated | 0 |
| Event Count | 4241 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 116.9s |
| Review | 21.0s |
| **Total** | **138.7s** |

## Configuration

- **name:** dotnet-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=406513 out=3826 |
| Turn Count | 10 |

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

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide correct .NET Microsoft.Azure.Cosmos SDK v3 error handling and retry guidance for 429 throttling, 404 not found, and 409 conflict responses, with local code examples only.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a Microsoft.Azure.Cosmos SDK v3 error-handling example covering 429, 404, and 409 responses.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos SDK v3 CosmosException handling, 429 RetryAfter behavior and retry configuration, plus 404 and 409 status codes.","learn":true}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find authoritative SDK v3 behavior for CosmosException status codes, RetryAfter, automatic retries, and retry options.","parameters":{"query":"Microsoft.Azure.Cosmos v3 CosmosException StatusCode RetryAfter 429 automatic retry MaxRetryAttemptsOnRateLimitedRequests MaxRetryWaitTimeOnRateLimitedRequests 404 409"}}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# examples for Microsoft.Azure.Cosmos v3 CosmosException handling and retry configuration.","parameters":{"language":"csharp","query":"Microsoft.Azure.Cosmos v3 catch CosmosException RetryAfter HttpStatusCode TooManyRequests NotFound Conflict CosmosClientOptions retry"}}
```

### 🔧 6. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787764892616-copilot-tool-output-41e6b68cd8d64a45bc2eb8e746d5548b.txt","pattern":"\"title\":\"[^\"]+\"|\"url\":\"[^\"]+\"|429|RetryAfter|MaxRetry"}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"$data = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787764892616-copilot-tool-output-41e6b68cd8d64a45bc2eb8e746d5548b.txt' | ConvertFrom-Json; $data.results | Select-Object title, url | Format-Table -AutoSize","description":"List Cosmos documentation search results"}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"$line = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787764892616-copilot-tool-output-41e6b68cd8d64a45bc2eb8e746d5548b.txt' | Select-Object -First 1; $data = $line | ConvertFrom-Json; $data.results | Select-Object title, url | Format-Table -Wrap","description":"List primary Cosmos documentation results"}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"$line = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787764892616-copilot-tool-output-41e6b68cd8d64a45bc2eb8e746d5548b.txt' | Select-Object -First 1; $data = $line | ConvertFrom-Json; $data.results | Select-Object title, contentUrl | Format-List","description":"Show Cosmos documentation URLs"}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official Cosmos DB resilient application retry guidance.","parameters":{"url":"https://learn.microsoft.com/azure/cosmos-db/conceptual-resilient-sdk-applications"}}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official Cosmos DB HTTP status code semantics for 429, 404, and 409.","parameters":{"url":"https://learn.microsoft.com/rest/api/cosmos-db/http-status-codes-for-cosmosdb"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read SDK v3 retry option semantics and defaults.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests?view=azure-dotnet"}}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read SDK v3 maximum retry wait option semantics and defaults.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretrywaittimeonratelimitedrequests?view=azure-dotnet"}}
```

## Copilot Response

I’ll ground the retry guidance in current Cosmos DB SDK v3 behavior, including what the SDK retries automatically and when application-level retries are appropriate.

The Cosmos DB .NET SDK v3 **automatically retries 429 responses by default**, honoring the server’s `x-ms-retry-after-ms` header. Configure that built-in behavior first; implement application-level retries only if you need a larger retry budget or custom logging—and avoid accidentally stacking two retry policies.

```csharp
using Azure.Identity;
using Microsoft.Azure.Cosmos;

CosmosClientOptions options = new()
{
    // Defaults are 9 retries and 30 seconds cumulative retry wait.
    MaxRetryAttemptsOnRateLimitedRequests = 9,
    MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30),
    ApplicationName = "OrdersApi"
};

// Use one long-lived CosmosClient per application/account.
CosmosClient client = new(
    Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")
        ?? throw new InvalidOperationException("COSMOS_ENDPOINT is required."),
    new DefaultAzureCredential(),
    options);

Container container = client.GetContainer("orders-db", "orders");
```

## Inspecting `CosmosException`

`RetryAfter` is the SDK’s typed representation of the `x-ms-retry-after-ms` response header.

```csharp
try
{
    ItemResponse<Order> response = await container.ReadItemAsync<Order>(
        id: orderId,
        partitionKey: new PartitionKey(customerId),
        cancellationToken: cancellationToken);
}
catch (CosmosException ex)
{
    HttpStatusCode statusCode = ex.StatusCode;
    TimeSpan retryAfter = ex.RetryAfter;

    logger.LogError(
        ex,
        "Cosmos request failed. Status={StatusCode}, SubStatus={SubStatusCode}, " +
        "RetryAfter={RetryAfter}, ActivityId={ActivityId}, RequestCharge={RequestCharge}",
        (int)statusCode,
        ex.SubStatusCode,
        retryAfter,
        ex.ActivityId,
        ex.RequestCharge);

    throw;
}
```

If a 429 reaches this catch block while automatic retries are enabled, the SDK has already exhausted the configured attempt or cumulative-wait limit.

## Handling 404 and 409

Neither response should normally be retried.

```csharp
public async Task<Order?> GetOrderAsync(
    Container container,
    string orderId,
    string customerId,
    CancellationToken cancellationToken)
{
    try
    {
        ItemResponse<Order> response = await container.ReadItemAsync<Order>(
            orderId,
            new PartitionKey(customerId),
            cancellationToken: cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.NotFound)
    {
        // Expected outcome when the item does not exist.
        return null;
    }
}

public async Task<bool> TryCreateOrderAsync(
    Container container,
    Order order,
    CancellationToken cancellationToken)
{
    try
    {
        await container.CreateItemAsync(
            order,
            new PartitionKey(order.CustomerId),
            cancellationToken: cancellationToken);

        return true;
    }
    catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.Conflict)
    {
        // The same ID/partition key or a unique-key value already exists.
        return false;
    }
}
```

A `409 Conflict` is a business/concurrency result, not a transient failure. Depending on the operation, handle it by rejecting the duplicate, reading the existing item, or using `UpsertItemAsync` when replacement is explicitly intended.

## Custom 429 retry policy

If you need full control, disable the SDK’s 429 retries so the custom policy is the **only** retry layer:

```csharp
CosmosClientOptions options = new()
{
    MaxRetryAttemptsOnRateLimitedRequests = 0
};
```

Then use the server-provided delay as the minimum, adding bounded exponential backoff and jitter:

```csharp
using System.Diagnostics;
using System.Net;
using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Logging;

public static class CosmosRetry
{
    public static async Task<T> ExecuteAsync<T>(
        Func<CancellationToken, Task<T>> operation,
        ILogger logger,
        int maxRetries,
        TimeSpan maxTotalWait,
        CancellationToken cancellationToken)
    {
        Stopwatch stopwatch = Stopwatch.StartNew();

        for (int attempt = 0; ; attempt++)
        {
            cancellationToken.ThrowIfCancellationRequested();

            try
            {
                return await operation(cancellationToken);
            }
            catch (CosmosException ex)
                when (ex.StatusCode == HttpStatusCode.TooManyRequests)
            {
                if (attempt >= maxRetries)
                {
                    throw;
                }

                // RetryAfter comes from the service's x-ms-retry-after-ms header.
                TimeSpan serverDelay = ex.RetryAfter;
                TimeSpan exponentialDelay = TimeSpan.FromMilliseconds(
                    Math.Min(5_000, 200 * Math.Pow(2, attempt)));

                TimeSpan delay = serverDelay > exponentialDelay
                    ? serverDelay
                    : exponentialDelay;

                delay += TimeSpan.FromMilliseconds(Random.Shared.Next(0, 250));

                if (stopwatch.Elapsed + delay > maxTotalWait)
                {
                    throw;
                }

                logger.LogWarning(
                    "Cosmos throttled request. Status={StatusCode}, RetryAfter={RetryAfter}, " +
                    "ActualDelay={Delay}, Retry={Retry}/{MaxRetries}, ActivityId={ActivityId}",
                    (int)ex.StatusCode,
                    ex.RetryAfter,
                    delay,
                    attempt + 1,
                    maxRetries,
                    ex.ActivityId);

                await Task.Delay(delay, cancellationToken);
            }
        }
    }
}
```

Usage:

```csharp
ItemResponse<Order> response = await CosmosRetry.ExecuteAsync(
    ct => container.ReadItemAsync<Order>(
        orderId,
        new PartitionKey(customerId),
        cancellationToken: ct),
    logger,
    maxRetries: 5,
    maxTotalWait: TimeSpan.FromSeconds(30),
    cancellationToken);
```

For sustained 429s, retries only absorb temporary bursts. Also inspect partition-key distribution, request concurrency, RU consumption, and provisioned/autoscale throughput.

**References:**

- [Design resilient applications with Azure Cosmos DB SDKs](https://learn.microsoft.com/azure/cosmos-db/conceptual-resilient-sdk-applications)
- [Cosmos DB HTTP status codes](https://learn.microsoft.com/rest/api/cosmos-db/http-status-codes-for-cosmosdb)
- [`MaxRetryAttemptsOnRateLimitedRequests`](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests?view=azure-dotnet)
- [`MaxRetryWaitTimeOnRateLimitedRequests`](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretrywaittimeonratelimitedrequests?view=azure-dotnet)

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response provides strong status-specific exception handling and a robust 429 retry implementation, but omits successful-response RU charge tracking and Cosmos diagnostics output.

### Strengths

- Uses CosmosException consistently with StatusCode, SubStatusCode, and RetryAfter.
- Explains built-in throttling retries and avoids stacking them with a custom retry policy.
- Implements bounded, cancellation-aware 429 backoff with server delay, exponential delay, and jitter.
- Handles 404 and 409 as expected non-transient outcomes.

### Issues

- Does not demonstrate tracking RU consumption through RequestCharge on ItemResponse values.
- Does not show the Diagnostics string from an exception or response.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
