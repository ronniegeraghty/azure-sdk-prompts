# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 334.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 334.0s |
| Timestamp | 2026-08-27T14:00:39Z |
| Files Generated | 0 |
| Event Count | 3270 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 300.6s |
| Review | 32.2s |
| **Total** | **334.0s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-diagnostics |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=194793 out=3517 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 3809ms

**Input:**

```json
{"skill":"azure-diagnostics"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Diagnostics

> **AUTHORITATIVE GUIDANCE — MANDATORY COMPLIANCE**
>
> This document is the **official source** for debugging and troubleshooting Azure production issues. Follow these instructions to diagnose and resolve common Azure service problems systematically.

## Triggers

Activate this skill when user wants to:
- Debug or troubleshoot production issues
- Diagnose errors in Azure services
- Analyze application logs or metrics
- Fix image pull, cold start, or health probe issues
- Investigate why Azure resources are failing
- Find root cause of application errors
- Troubleshoot App Service issues (high CPU, deployment failures, crashes, slow responses, TLS/custom domains)
- Respond to prompts like "troubleshoot app service", "app service high CPU", or "app service deployment failure"
- Troubleshoot Azure Function Apps (invocation failures, timeouts, binding errors)
- Find the App Insights or Log Analytics workspace linked to a Function App
- Troubleshoot AKS clusters, nodes, pods, ingress, or Kubernetes networking issues
- Troubleshoot Azure VM connectivity issues (RDP/SSH failures, port 3389/22 timeouts, NSG or firewall blocking, credential resets)
- Troubleshoot Azure Messaging SDK issues (Event Hubs, Service Bus connection failures, AMQP errors, message lock issues)

## Rules

1. Start with systematic diagnosis flow
2. Use AppLens (MCP) for AI-powered diagnostics when available
3. Check resource health before deep-diving into logs
4. Select appropriate troubleshooting guide based on service type
5. Document findings and attempted remediation steps
6. Route AKS incidents to the dedicated AKS troubleshooting document

---

## Quick Diagnosis Flow

1. **Identify symptoms** - What's failing?
2. **Check resource health** - Is Azure healthy?
3. **Review logs** - What do logs show?
4. **Analyze metrics** - Performance patterns?
5. **Investigate recent changes** - What changed?

---

## Troubleshooting Guides by Service

| Ser
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 184757ms

**Input:**

```json
{"intent":"Generate local-only .NET code using Microsoft.Azure.Cosmos SDK v3 to handle CosmosException responses, especially 429 throttling with RetryAfter, plus 404 and 409.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 184760ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos SDK v3 CosmosException StatusCode RetryAfter, built-in 429 retry configuration, 404 item not found, and 409 conflict handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. web_search (via github-mcp-server) — 30399ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Cosmos DB .NET SDK v3 CosmosException RetryAfter 429 MaxRetryAttemptsOnRateLimitedRequests 404 409"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's an overview addressing your queries regarding Azure Cosmos DB .NET SDK v3—specifically handling `CosmosException`, `RetryAfter`, 429 errors, `MaxRetryAttemptsOnRateLimitedRequests`, and common status codes like 404 and 409:\n\n### 1. `CosmosException` and `RetryAfter`\n- The `.NET SDK` throws a `CosmosException` for Cosmos DB errors, which includes HTTP status codes and diagnostic information.\n- The property `CosmosException.RetryAfter` tells you how long (in `TimeSpan`) your application should wait before retrying after hitting a throttling (HTTP 429) error. Use this value to control delay between retry attempts【3:5†source】.\n\n### 2. Handling 429 (\"Too Many Requests\") and Retry Logic\n- Cosmos DB uses HTTP 429 to signal rate limiting. This happens when client requests exceed the provisioned RU/s (Request Units/sec).\n- The SDK honors the \"x-ms-retry-after-ms\" header from Cosmos DB, using this time before retrying automatically【3:4†source】【3:1†source】.\n- Configuration: \n  - `CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests` property sets the maximum number of retry attempts for 429 errors. Default is 9, so SDK will attempt a total of 10 tries before surfacing the error.\n  - Setting this property to 0 disables automatic retry, so your application must handle the logic【3:3†source】.\n\n### 3. Common Status Codes\n- **404 Not Found:** Usually means the resource (document, container, database, etc.) no longer exists or never did.\n- **409 Conflict:** Happens when you try to create a resource (like a document) with an ID that already exists in the partition【3:9†source】.\n\n### 4. Additional Guidance and Best Practices\n- Always log and inspect `CosmosException.Diagnostics` for troubleshooting complex issues.\n- Too many 404s or 409s could indicate logic errors—check existence before creating/deleting, and ensure idempotency or proper conflict handling.\n\n### References and F
... (truncated)
```

### ✅ 5. web_fetch — 3522ms

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: CosmosException.RetryAfter Property (Microsoft.Azure.Cosmos) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Microsoft.Azure.Cosmos.CosmosException.RetryAfter*
namespace: Microsoft.Azure.Cosmos
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-net/issues
author: azure-sdk
ms.author: azsdkdocs
ms.topic: generated-reference
ms.devlang: csharp
ms.date: 2021-04-13T00:00:00.0000000Z
uhfHeaderId: Azure
breadcrumb_path: ../azure_breadcrumb/toc.json
apiPlatform: dotnet
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/cd668c2f-f5b3-4573-8ad1-019570e3e2db
ms.service: azure-cosmos-db
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Microsoft.Azure.Cosmos.CosmosException.RetryAfter
- Microsoft.Azure.Cosmos.CosmosException.get_RetryAfter
api_location:
- Microsoft.Azure.Cosmos.Client.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 10fb499b-bbf1-7f4f-8728-aa3ffe4e9ec0
document_version_independent_id: 94b3a53b-8a90-8b0e-a982-c54f9f094e38
updated_at: 2026-04-08T11:32:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Microsoft.Azure.Cosmos/CosmosException.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/c5c49becd84b5571cceb0df00619ac7dd9f7ebbc/xml/Microsoft.Azure.Cosmos/CosmosException.xml
git_commit_id: c5c49becd84b5571cceb0df00619ac7dd9f7ebbc
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: d
... (truncated)
```

## Copilot Response

I’ll ground the retry pattern in the current SDK v3 guidance, including where built-in retries end and application retries begin.

The Azure documentation endpoint timed out, so I’m switching to the public Microsoft Learn material while keeping the example SDK-v3-specific.

The Cosmos DB .NET SDK v3 already retries HTTP 429 responses automatically. Prefer configuring that behavior rather than adding another retry loop. If you need custom retry behavior, disable SDK retries to avoid nested retries and excessive latency.

### Recommended: SDK-managed 429 retries

```csharp
using Microsoft.Azure.Cosmos;

CosmosClient client = new(
    accountEndpoint: Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")
        ?? throw new InvalidOperationException("COSMOS_ENDPOINT is missing."),
    authKeyOrResourceToken: Environment.GetEnvironmentVariable("COSMOS_KEY")
        ?? throw new InvalidOperationException("COSMOS_KEY is missing."),
    clientOptions: new CosmosClientOptions
    {
        // Retries after the initial request.
        MaxRetryAttemptsOnRateLimitedRequests = 9,

        // Maximum cumulative time spent waiting for 429 retries.
        MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30)
    });
```

The SDK reads the `x-ms-retry-after-ms` response header, waits for that duration, and retries. A `CosmosException` with status 429 is normally surfaced only after these limits are exhausted.

### Catching status codes and retry-after

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;

try
{
    ItemResponse<Order> response = await container.ReadItemAsync<Order>(
        id: orderId,
        partitionKey: new PartitionKey(customerId),
        cancellationToken: cancellationToken);

    Order order = response.Resource;
}
catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.TooManyRequests)
{
    int statusCode = (int)ex.StatusCode;           // 429
    TimeSpan? retryAfter = ex.RetryAfter;          // x-ms-retry-after-ms
    string activityId = ex.ActivityId;
    double requestCharge = ex.RequestCharge;

    Console.Error.WriteLine(
        $"Cosmos DB throttled the request. " +
        $"Status={statusCode}, " +
        $"RetryAfterMs={retryAfter?.TotalMilliseconds}, " +
        $"RequestCharge={requestCharge}, " +
        $"ActivityId={activityId}");

    throw;
}
catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.NotFound)
{
    // The item does not exist for this ID and partition key.
}
catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.Conflict)
{
    // An item with this ID already exists in the logical partition.
}
```

`RetryAfter` is nullable, but a throttling response normally supplies it. Log `ActivityId` and `Diagnostics` when troubleshooting:

```csharp
Console.Error.WriteLine(ex.Diagnostics);
```

### Custom bounded 429 retry

Use this only when you need application-specific behavior. Configure the client with `MaxRetryAttemptsOnRateLimitedRequests = 0` first so the SDK and application do not both retry.

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;

static async Task<T> ExecuteWith429RetryAsync<T>(
    Func<CancellationToken, Task<T>> operation,
    int maxRetries,
    CancellationToken cancellationToken)
{
    int retryCount = 0;

    while (true)
    {
        try
        {
            return await operation(cancellationToken);
        }
        catch (CosmosException ex)
            when (ex.StatusCode == HttpStatusCode.TooManyRequests &&
                  retryCount < maxRetries)
        {
            retryCount++;

            // Honor the service-provided x-ms-retry-after-ms value.
            TimeSpan serviceDelay = ex.RetryAfter
                ?? TimeSpan.FromMilliseconds(
                    Math.Min(1_000 * Math.Pow(2, retryCount - 1), 30_000));

            // Small jitter prevents many callers from retrying simultaneously.
            TimeSpan delay =
                serviceDelay + TimeSpan.FromMilliseconds(Random.Shared.Next(0, 250));

            Console.Error.WriteLine(
                $"Cosmos DB returned {(int)ex.StatusCode}. " +
                $"Retry {retryCount}/{maxRetries} after {delay.TotalMilliseconds:N0} ms. " +
                $"ActivityId={ex.ActivityId}");

            await Task.Delay(delay, cancellationToken);
        }
    }
}
```

Example:

```csharp
ItemResponse<Order> response = await ExecuteWith429RetryAsync(
    ct => container.ReadItemAsync<Order>(
        orderId,
        new PartitionKey(customerId),
        cancellationToken: ct),
    maxRetries: 5,
    cancellationToken);
```

The final 429 propagates after `maxRetries`, allowing the caller to fail the request, enqueue it for later processing, or return an appropriate response.

### Handling 404 without a preflight read

Catch 404 directly instead of first checking whether the item exists; a check-then-read sequence costs extra RUs and introduces a race.

```csharp
static async Task<T?> ReadItemOrNullAsync<T>(
    Container container,
    string id,
    PartitionKey partitionKey,
    CancellationToken cancellationToken)
    where T : class
{
    try
    {
        ItemResponse<T> response = await container.ReadItemAsync<T>(
            id,
            partitionKey,
            cancellationToken: cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.NotFound)
    {
        return null;
    }
}
```

A 404 can also indicate that the supplied partition-key value is wrong, even if an item with that ID exists elsewhere.

### Handling 409 conflicts

A create operation returns 409 when the same item ID already exists in the same logical partition:

```csharp
try
{
    await container.CreateItemAsync(
        order,
        new PartitionKey(order.CustomerId),
        cancellationToken: cancellationToken);
}
catch (CosmosException ex) when (ex.StatusCode == HttpStatusCode.Conflict)
{
    Console.Error.WriteLine(
        $"Order '{order.Id}' already exists. ActivityId={ex.ActivityId}");

    // Apply domain-specific behavior: report duplicate, read the existing
    // item, or treat the create as idempotently successful.
}
```

Do not retry 409 blindly—the same create will generally conflict again. If the intended behavior is “create or replace,” use `UpsertItemAsync`. For optimistic-concurrency ETag failures, handle HTTP **412 Precondition Failed**, not 409.

References: [CosmosException.RetryAfter](https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter), [rate-limit retry options](https://learn.microsoft.com/en-us/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests), and [Cosmos DB status codes](https://learn.microsoft.com/en-us/rest/api/cosmos-db/http-status-codes-for-cosmosdb).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response gives strong SDK-managed and custom throttling guidance plus appropriate 404 and 409 handling, but omits SubStatusCode and successful-response RU tracking.

### Strengths

- Correctly treats CosmosException as the central SDK exception type.
- Provides bounded 429 retry logic that honors RetryAfter and avoids nested retries.
- Clearly handles missing items and duplicate-item conflicts.
- Includes client retry configuration and diagnostics logging.

### Issues

- SubStatusCode is not extracted or logged.
- RequestCharge is shown only on CosmosException, not on successful ItemResponse or other response objects.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
