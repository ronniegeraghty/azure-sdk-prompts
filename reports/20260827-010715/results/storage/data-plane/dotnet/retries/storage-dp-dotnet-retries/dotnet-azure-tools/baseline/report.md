# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 171.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 171.1s |
| Timestamp | 2026-08-26T18:41:03Z |
| Files Generated | 0 |
| Event Count | 3706 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 148.0s |
| Review | 22.2s |
| **Total** | **171.1s** |

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
| Token Usage | in=73224 out=3551 |
| Turn Count | 3 |

## Prompt Sent

```
How do I configure custom retry policies for Azure Blob Storage operations in .NET?
My application needs to:
1. Set a custom retry policy with 5 max retries and exponential backoff
2. Configure per-operation timeouts so a single upload doesn't hang forever
3. Handle network errors (transient) differently from auth errors (non-transient)
4. Use a custom retry policy for specific high-value operations
5. Implement circuit-breaker patterns for sustained failures

Show me how to configure BlobClientOptions with custom RetryOptions,
and explain which HTTP status codes the SDK considers retryable by default.
Use the Azure.Storage.Blobs SDK.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Azure.Storage.Blobs BlobClientOptions RetryOptions .NET retry status codes default NetworkTimeout MaxRetries exponential"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:github.com/Azure/azure-sdk-for-net Azure.Core ResponseClassifier retry status codes 408 429 500 502 503 504 RetryPolicy .NET"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Azure SDK for .NET custom HttpPipelinePolicy retry circuit breaker BlobClientOptions"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/ResponseClassifier.cs"}
```

### 🔧 6. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions?view=azure-dotnet"}
```

## Copilot Response

I’ll verify the current Azure SDK retry semantics and APIs against Microsoft’s documentation, then provide an end-to-end .NET example.

The built-in timeout is attempt/network-operation scoped, so I’m checking the exact extensibility signatures needed to distinguish failures without accidentally stacking two retry loops.

Configure retries once per client, and use a cancellation token for the **overall upload deadline**. `NetworkTimeout` only limits each individual HTTP/network attempt; it does not guarantee that a multipart upload—including retries and backoff—finishes within that time.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;

Uri accountUri =
    new("https://<storage-account>.blob.core.windows.net");

var options = new BlobClientOptions
{
    Retry =
    {
        MaxRetries = 5,                         // 5 retries + initial attempt
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),        // Exponential backoff base
        MaxDelay = TimeSpan.FromSeconds(15),
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

var serviceClient = new BlobServiceClient(
    accountUri,
    new DefaultAzureCredential(),
    options);

BlobClient blob = serviceClient
    .GetBlobContainerClient("documents")
    .GetBlobClient("report.pdf");
```

`MaxRetries = 5` means up to **six total attempts**: the original request plus five retries.

## Apply an overall operation timeout

Use a cancellation token to limit the entire upload, including all underlying requests, retry delays, and block-transfer work:

```csharp
await using FileStream content = File.OpenRead("report.pdf");

using var operationTimeout =
    new CancellationTokenSource(TimeSpan.FromMinutes(2));

try
{
    await blob.UploadAsync(
        content,
        new BlobUploadOptions
        {
            HttpHeaders = new BlobHttpHeaders
            {
                ContentType = "application/pdf"
            },
            TransferOptions = new StorageTransferOptions
            {
                MaximumConcurrency = 4,
                MaximumTransferSize = 4 * 1024 * 1024
            }
        },
        operationTimeout.Token);
}
catch (OperationCanceledException)
    when (operationTimeout.IsCancellationRequested)
{
    Console.Error.WriteLine("Upload exceeded its two-minute deadline.");
    throw;
}
```

These controls serve different purposes:

| Setting | Scope |
|---|---|
| `Retry.NetworkTimeout` | One individual network operation/attempt |
| `CancellationTokenSource.CancelAfter` | The complete SDK operation |
| `MaxRetries` | Number of retry attempts after the initial request |
| `MaxDelay` | Backoff cap unless the service supplies `Retry-After` |

## Transient versus authentication failures

Azure.Core’s default response classifier retries these HTTP statuses:

| Status | Meaning |
|---|---|
| `408` | Request Timeout |
| `429` | Too Many Requests |
| `500` | Internal Server Error |
| `502` | Bad Gateway |
| `503` | Service Unavailable |
| `504` | Gateway Timeout |

It also retries:

- `IOException`
- `RequestFailedException` with status `0`, usually indicating no HTTP response/network failure
- Non-user-initiated operation cancellation, such as an internal transport timeout

It does **not** retry `401 Unauthorized` or `403 Forbidden` by default. Those usually indicate invalid credentials, an expired/misconfigured token, RBAC propagation/configuration problems, or insufficient permissions. Retrying them with backoff generally cannot fix the problem.

```csharp
try
{
    await blob.UploadAsync(content, overwrite: true);
}
catch (RequestFailedException ex)
    when (ex.Status is 401 or 403)
{
    // Non-transient: alert, refresh/fix credentials, or correct RBAC.
    Console.Error.WriteLine(
        $"Blob authorization failed ({ex.Status}): {ex.ErrorCode}");

    throw;
}
catch (RequestFailedException ex)
    when (IsTransient(ex))
{
    // The SDK has already exhausted its configured retries.
    Console.Error.WriteLine(
        $"Transient storage failure after retries: {ex.Status}");

    throw;
}

static bool IsTransient(Exception exception)
{
    return exception switch
    {
        RequestFailedException ex =>
            ex.Status is 0 or 408 or 429 or 500 or 502 or 503 or 504,

        IOException => true,
        HttpRequestException => true,

        // Azure.Core can aggregate repeated transport exceptions.
        AggregateException ex =>
            ex.InnerExceptions.Count > 0 &&
            ex.InnerExceptions.All(IsTransient),

        // Caller cancellation/deadline is not a service transient fault.
        OperationCanceledException => false,

        _ => false
    };
}
```

When Storage returns `Retry-After`, `retry-after-ms`, or `x-ms-retry-after-ms`, the SDK honors the service-provided delay instead of the locally calculated backoff.

## A stronger policy for high-value operations

Client options apply to every operation made through that client. Create a separate client for high-value operations rather than adding another retry loop around a normally configured client:

```csharp
static BlobClientOptions CreateStandardOptions() => new()
{
    Retry =
    {
        MaxRetries = 5,
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(15),
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

static BlobClientOptions CreateHighValueOptions() => new()
{
    Retry =
    {
        MaxRetries = 8,
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(2),
        MaxDelay = TimeSpan.FromSeconds(45),
        NetworkTimeout = TimeSpan.FromSeconds(60)
    }
};

var credential = new DefaultAzureCredential();

var normalService = new BlobServiceClient(
    accountUri,
    credential,
    CreateStandardOptions());

var highValueService = new BlobServiceClient(
    accountUri,
    credential,
    CreateHighValueOptions());

BlobClient highValueBlob = highValueService
    .GetBlobContainerClient("critical-documents")
    .GetBlobClient("quarterly-ledger.pdf");
```

Azure SDK clients are thread-safe and intended to be reused, so register both client configurations as long-lived services rather than constructing them for every request.

For completely custom retry behavior, `BlobClientOptions.RetryPolicy` accepts an `HttpPipelinePolicy`. Setting `RetryPolicy` **replaces** the built-in retry policy; it should not be combined conceptually with an external retry loop unless you deliberately account for multiplied attempts.

## Circuit breaker for sustained failures

A circuit breaker is best placed around the high-level SDK operation. The SDK performs short-term retries internally; the breaker observes operations that still fail after those retries. This example uses Polly 8:

```csharp
using Polly;
using Polly.CircuitBreaker;

ResiliencePipeline circuitBreaker =
    new ResiliencePipelineBuilder()
        .AddCircuitBreaker(new CircuitBreakerStrategyOptions
        {
            ShouldHandle = new PredicateBuilder()
                .Handle<Exception>(IsTransient),

            // Open when at least 50% of 8+ operations fail in 30 seconds.
            FailureRatio = 0.50,
            MinimumThroughput = 8,
            SamplingDuration = TimeSpan.FromSeconds(30),
            BreakDuration = TimeSpan.FromSeconds(30)
        })
        .Build();

using var uploadDeadline =
    new CancellationTokenSource(TimeSpan.FromMinutes(5));

await using FileStream highValueContent =
    File.OpenRead("quarterly-ledger.pdf");

try
{
    await circuitBreaker.ExecuteAsync(
        async cancellationToken =>
        {
            await highValueBlob.UploadAsync(
                highValueContent,
                overwrite: true,
                cancellationToken);
        },
        uploadDeadline.Token);
}
catch (BrokenCircuitException)
{
    Console.Error.WriteLine(
        "Blob Storage circuit is open; the operation was not sent.");
    throw;
}
```

Because `IsTransient` excludes `401`, `403`, and caller cancellation, credential problems and operation deadlines do not contribute to opening the circuit. Avoid adding Polly retry to this pipeline unless you reduce or disable SDK retries; otherwise, retries multiply and can significantly increase traffic and latency.

References: [Blob Storage retry guidance](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), [RetryOptions](https://learn.microsoft.com/dotnet/api/azure.core.retryoptions), and the Azure.Core [default `ResponseClassifier`](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response thoroughly covers retry configuration, timeout scopes, transient handling, specialized clients, and circuit breaking, but omits geo-redundant retry configuration and an explicit complete list of non-retryable statuses.

### Strengths

- Correct BlobClientOptions retry configuration with exponential backoff.
- Clear distinction between NetworkTimeout and an overall CancellationToken deadline.
- Accurate default retryable status list and useful transient-network classification.
- Strong guidance for high-value-operation clients and Polly circuit-breaker integration.

### Issues

- Missing GeoRedundantSecondaryUri configuration and behavior.
- Non-retryable status discussion omits 400, 404, and 409.

## Grader Results

- retry-configuration.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `BlobClientOptions.Retry` configuration with `RetryOptions`: Pass
      - `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed): Pass
      - `NetworkTimeout` for per-request timeouts: Pass
      - Default retryable status codes (408, 429, 500, 502, 503, 504): Pass
      - Non-retryable errors (400, 401, 403, 404, 409): Fail
      - Per-operation `CancellationToken` for timeout control: Pass
      - Geo-redundant retry with `GeoRedundantSecondaryUri`: Fail
      - Interaction with Polly or other resilience libraries: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.7500** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
