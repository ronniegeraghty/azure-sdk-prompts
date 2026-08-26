# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 183.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 183.2s |
| Timestamp | 2026-08-26T18:37:59Z |
| Files Generated | 0 |
| Event Count | 5206 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 159.5s |
| Review | 22.6s |
| **Total** | **183.2s** |

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
| Token Usage | in=528739 out=6128 |
| Turn Count | 11 |

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

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide current best practices for configuring Azure.Storage.Blobs BlobClientOptions retries, per-operation timeouts, custom HttpPipelinePolicy behavior, and circuit breakers in .NET.","learn":true}
```

### 🔧 2. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Storage.Blobs BlobClientOptions RetryOptions, default retryable HTTP status codes, NetworkTimeout, cancellation tokens, custom per-retry pipeline policies, and circuit breaker guidance in .NET.","learn":true}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure SDK for .NET code-generation best practices for Blob Storage retry, timeout, custom pipeline policy, and circuit breaker examples.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find authoritative retry and timeout behavior for Azure.Storage.Blobs .NET.","parameters":{"query":"Azure.Storage.Blobs .NET BlobClientOptions RetryOptions MaxRetries Delay MaxDelay Mode NetworkTimeout default retryable HTTP status codes cancellation token AddPolicy HttpPipelinePosition PerRetry ResponseClassifier"}}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for BlobClientOptions RetryOptions and BlobClient upload cancellation.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs BlobClientOptions RetryOptions exponential NetworkTimeout UploadAsync CancellationToken AddPolicy"}}
```

### 🔧 6. rg

**Input:**

```json
{"-C":2,"-i":true,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769510609-copilot-tool-output-afad22e61f53447983678906edf6597a.txt","pattern":"408|429|500|502|503|504|retryable|ResponseClassifier|NetworkTimeout|CancellationToken|AddPolicy|PerRetry"}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Identify exact default Azure Core .NET retry classifier behavior and supported custom retry policy mechanisms.","parameters":{"query":"site:learn.microsoft.com .NET Azure.Core RetryPolicy default retry status codes 408 429 500 502 503 504 ResponseClassifier IsRetriableResponse HttpPipelinePolicy AddPolicy PerRetry custom retry policy"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read full official Blob Storage retry policy guidance.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy"}}
```

### 🔧 9. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Azure SDK for .NET retry guidance including exact default status codes.","parameters":{"url":"https://learn.microsoft.com/azure/developer/dotnet/sdk/connection-resiliency"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Confirm current Azure.Core API for replacing a client retry policy and custom response classification in .NET.","parameters":{"query":"site:learn.microsoft.com/en-us/dotnet/api/azure.core ClientOptions RetryPolicy property custom HttpPipelinePolicy ResponseClassifier Azure.Core .NET"}}
```

### 🔧 11. rg

**Input:**

```json
{"-C":3,"head_limit":50,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769532023-copilot-tool-output-15c603e679294e2e9fab41ddc260e456.txt","pattern":"ClientOptions.RetryPolicy Property|ResponseClassifier|IsRetriableResponse|StatusCodes|408|429|500|502|503|504|RetryPolicy ="}
```

### 🔧 12. rg

**Input:**

```json
{"-C":2,"-i":true,"head_limit":50,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769510609-copilot-tool-output-afad22e61f53447983678906edf6597a.txt","pattern":"408|429|500|502|503|504|retryable status|authentication|401|403|NetworkTimeout"}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official docs with exact default retriable HTTP status codes in Azure SDK for .NET.","parameters":{"query":"Azure SDK for .NET retries default HTTP status codes 408 429 500 502 503 504 Azure.Core RetryPolicy"}}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"$files = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769532023-copilot-tool-output-15c603e679294e2e9fab41ddc260e456.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769540322-copilot-tool-output-e9f0e0b3d45a486ea23cb8958c34e2c6.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769510609-copilot-tool-output-afad22e61f53447983678906edf6597a.txt'); foreach ($file in $files) { $json = Get-Content -Raw $file | ConvertFrom-Json; $json.results | Where-Object { $_.title -match 'RetryPolicy|ResponseClassifier|retry policy|RetryOptions' -or $_.content -match '408.*429|429.*500|IsRetriableResponse' } | Select-Object -First 8 | ForEach-Object { \"TITLE: $($_.title)`nURL: $($_.url)`nCONTENT:`n$($_.content)`n---\" } }","description":"Extract relevant Azure retry documentation","initial_wait":30,"mode":"sync"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$files = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769532023-copilot-tool-output-15c603e679294e2e9fab41ddc260e456.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769540322-copilot-tool-output-e9f0e0b3d45a486ea23cb8958c34e2c6.txt'); foreach ($file in $files) { $json = Get-Content -Raw $file | ConvertFrom-Json; $json.results | Where-Object { $_.title -match 'RetryPolicy|ResponseClassifier|retry policy|RetryOptions' -or $_.content -match '408.*429|429.*500|IsRetriableResponse' } | Select-Object -First 10 | ForEach-Object { \"TITLE: $($_.title)`nURL: $($_.url)`nCONTENT:`n$($_.content)`n---\" } }","description":"Extract retry API documentation details","initial_wait":30,"mode":"sync"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769532023-copilot-tool-output-15c603e679294e2e9fab41ddc260e456.txt","view_range":[1,5]}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769540322-copilot-tool-output-e9f0e0b3d45a486ea23cb8958c34e2c6.txt","view_range":[1,5]}
```

### 🔧 18. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find exact C# signatures and examples for deriving Azure.Core.Pipeline.RetryPolicy in Azure.Core 1.60.","parameters":{"query":"site:learn.microsoft.com/dotnet/api/azure.core.pipeline.retrypolicy ShouldRetryAsync CalculateDelay ProcessAsync Azure.Core 1.60 custom RetryPolicy C#"}}
```

### 🔧 19. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current Microsoft.Extensions.Resilience circuit breaker API examples suitable for wrapping Azure SDK operations.","parameters":{"query":"site:learn.microsoft.com .NET resilience pipeline AddCircuitBreaker ResiliencePipelineBuilder circuit breaker ShouldHandle RequestFailedException example 2026"}}
```

### 🔧 20. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769574140-copilot-tool-output-d55ac3702a474c8e9e68d668347af3b2.txt","view_range":[1,3]}
```

### 🔧 21. rg

**Input:**

```json
{"-C":2,"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787769574194-copilot-tool-output-6577f412120c457bb90ef2213f99637d.txt","pattern":"AddCircuitBreaker|CircuitBreakerStrategyOptions|ShouldHandle|FailureRatio|MinimumThroughput|BreakDuration|ExecuteAsync"}
```

## Copilot Response

I’ll ground the examples in the current Azure SDK retry pipeline and distinguish built-in retries, per-call cancellation, operation-specific policies, and circuit breaking.

The built-in `NetworkTimeout` is a per-network-operation idle timeout, not a deadline for the complete multipart upload. I’m checking the retry classifier details so the status-code guidance and custom-policy placement are exact.

For operation-specific behavior, the cleanest approach is usually a second long-lived `BlobClient` configured for that workload. If you truly need different classification or delay rules, `BlobClientOptions.RetryPolicy` replaces `RetryOptions`; I’m confirming the override signatures before showing that advanced option.

Use `BlobClientOptions.Retry` for SDK-level retries, and a cancellation token for the total upload deadline. `NetworkTimeout` only limits an individual network operation; it does **not** cap the duration of the complete upload.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

Uri serviceUri = new("https://<account-name>.blob.core.windows.net");
TokenCredential credential = new DefaultAzureCredential();

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 5,                       // 5 retries + the initial attempt
        Delay = TimeSpan.FromSeconds(1),      // Exponential backoff base
        MaxDelay = TimeSpan.FromSeconds(30),
        NetworkTimeout = TimeSpan.FromSeconds(60)
    }
};

var serviceClient = new BlobServiceClient(serviceUri, credential, options);
var blobClient = serviceClient
    .GetBlobContainerClient("documents")
    .GetBlobClient("report.pdf");
```

When Azure Storage returns `Retry-After` or `x-ms-retry-after-ms`, the SDK honors that server-provided delay instead of calculating its own delay.

## Total operation timeout

Use `CancellationTokenSource.CancelAfter` to bound the complete upload, including retries and multipart transfers:

```csharp
static async Task UploadWithDeadlineAsync(
    BlobClient blobClient,
    Stream content,
    TimeSpan timeout,
    CancellationToken cancellationToken)
{
    using var timeoutCts =
        CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);

    timeoutCts.CancelAfter(timeout);

    try
    {
        await blobClient.UploadAsync(
            content,
            overwrite: true,
            cancellationToken: timeoutCts.Token);
    }
    catch (OperationCanceledException)
        when (timeoutCts.IsCancellationRequested &&
              !cancellationToken.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"Blob upload exceeded its {timeout} operation deadline.");
    }
}
```

The distinction is:

| Setting | Scope |
|---|---|
| `Retry.NetworkTimeout` | One HTTP network operation, generally while waiting for data |
| `CancellationTokenSource.CancelAfter` | The entire logical upload |
| `MaxRetries = 5` | Up to six HTTP attempts: initial attempt plus five retries |

## Transient versus authentication errors

The built-in policy already retries network failures and recognized transient status codes. Authentication and authorization responses such as `401` and `403` aren't retried.

Handle final failures explicitly after SDK retries are exhausted:

```csharp
using Azure.Identity;
using System.Net.Http;

static bool IsTransient(Exception exception) =>
    exception switch
    {
        // Status 0 commonly represents a transport/network failure.
        RequestFailedException { Status: 0 } => true,

        RequestFailedException
        {
            Status: 408 or 429 or 500 or 502 or 503 or 504
        } => true,

        HttpRequestException => true,
        IOException => true,
        TimeoutException => true,

        _ => false
    };

try
{
    await UploadWithDeadlineAsync(
        blobClient,
        content,
        TimeSpan.FromMinutes(10),
        cancellationToken);
}
catch (AuthenticationFailedException ex)
{
    // Credential acquisition failed before a Storage request was sent.
    logger.LogError(ex, "Could not acquire an Azure credential.");
    throw;
}
catch (RequestFailedException ex) when (ex.Status is 401 or 403)
{
    // Non-transient: fix credentials, RBAC, SAS permissions, or token scope.
    logger.LogError(
        ex,
        "Blob authorization failed with HTTP {Status}. Error code: {Code}",
        ex.Status,
        ex.ErrorCode);

    throw;
}
catch (RequestFailedException ex) when (IsTransient(ex))
{
    // The SDK's configured retries have already been exhausted.
    logger.LogWarning(
        ex,
        "Blob operation failed after retries. HTTP status: {Status}",
        ex.Status);

    throw;
}
```

Do not add an unbounded outer retry around the SDK. Doing so multiplies attempts—for example, three outer attempts with six SDK attempts can produce up to 18 requests.

## Different policy for high-value operations

Azure clients are thread-safe and should be reused. Create a second long-lived client with more conservative options rather than changing options for every call:

```csharp
var standardOptions = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 5,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(30),
        NetworkTimeout = TimeSpan.FromSeconds(60)
    }
};

var highValueOptions = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 8,
        Delay = TimeSpan.FromSeconds(2),
        MaxDelay = TimeSpan.FromMinutes(1),
        NetworkTimeout = TimeSpan.FromMinutes(2)
    }
};

var standardServiceClient =
    new BlobServiceClient(serviceUri, credential, standardOptions);

var highValueServiceClient =
    new BlobServiceClient(serviceUri, credential, highValueOptions);

BlobClient highValueBlob = highValueServiceClient
    .GetBlobContainerClient("critical-documents")
    .GetBlobClient("contract.pdf");

await UploadWithDeadlineAsync(
    highValueBlob,
    content,
    timeout: TimeSpan.FromMinutes(20),
    cancellationToken);
```

This provides operation-specific retry behavior without rebuilding clients or mutating shared configuration. If you need to replace retry classification or delay calculation completely, `BlobClientOptions.RetryPolicy` accepts an `HttpPipelinePolicy` and replaces the `Retry` settings, but deriving from `Azure.Core.Pipeline.RetryPolicy` is an advanced option and requires preserving Azure pipeline processing context correctly.

## Circuit breaker

A circuit breaker should wrap the **logical SDK operation**, remain shared across calls, and count only transient failures. The SDK performs short-term retries; the breaker prevents continued traffic during sustained outages.

Using Polly v8:

```powershell
dotnet add package Polly.Core
```

```csharp
using Polly;
using Polly.CircuitBreaker;

ResiliencePipeline blobCircuitBreaker =
    new ResiliencePipelineBuilder()
        .AddCircuitBreaker(new CircuitBreakerStrategyOptions
        {
            ShouldHandle = new PredicateBuilder()
                .Handle<RequestFailedException>(IsTransient)
                .Handle<HttpRequestException>()
                .Handle<IOException>()
                .Handle<TimeoutException>(),

            // Open when at least 8 operations occur and 50% fail
            // during a 30-second sampling window.
            FailureRatio = 0.5,
            MinimumThroughput = 8,
            SamplingDuration = TimeSpan.FromSeconds(30),
            BreakDuration = TimeSpan.FromMinutes(1)
        })
        .Build();

try
{
    await blobCircuitBreaker.ExecuteAsync(
        async token =>
        {
            await UploadWithDeadlineAsync(
                highValueBlob,
                content,
                TimeSpan.FromMinutes(20),
                token);
        },
        cancellationToken);
}
catch (BrokenCircuitException ex)
{
    logger.LogWarning(
        ex,
        "Blob Storage circuit is open; the operation was rejected.");

    throw;
}
```

Keep the breaker as a singleton or otherwise long-lived. Creating one for every operation resets its failure history and defeats the pattern. Authentication failures are excluded from `ShouldHandle`, so bad credentials fail immediately without opening the circuit.

## Default retryable responses

The Azure SDK retry classifier treats these HTTP statuses as retryable by default:

| Status | Meaning |
|---:|---|
| `408` | Request Timeout |
| `429` | Too Many Requests |
| `500` | Internal Server Error |
| `502` | Bad Gateway |
| `503` | Service Unavailable |
| `504` | Gateway Timeout |

Transport failures, such as connection resets, socket failures, and network timeouts, are also retry candidates. `400`, `401`, `403`, `404`, and most other `4xx` responses aren't retried by default because they normally require changing the request, credentials, permissions, or resource name.

References: [Blob Storage retry policy](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), [`ClientOptions.RetryPolicy`](https://learn.microsoft.com/dotnet/api/azure.core.clientoptions.retrypolicy), and [`ClientOptions.AddPolicy`](https://learn.microsoft.com/dotnet/api/azure.core.clientoptions.addpolicy).

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The response thoroughly covers SDK retry configuration, timeout scopes, error classification, specialized clients, and circuit breaking, but omits geo-redundant secondary retries.

### Strengths

- Correct BlobClientOptions retry configuration with exponential backoff and five retries.
- Clear distinction between NetworkTimeout and an operation-wide CancellationToken deadline.
- Accurate retryable-status and transient-versus-authentication-error guidance.
- Strong Polly circuit-breaker example that avoids retry multiplication and excludes authentication failures.

### Issues

- No GeoRedundantSecondaryUri configuration or explanation is provided.

## Grader Results

- retry-configuration.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `BlobClientOptions.Retry` configuration with `RetryOptions`: Pass
      - `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed): Pass
      - `NetworkTimeout` for per-request timeouts: Pass
      - Default retryable status codes (408, 429, 500, 502, 503, 504): Pass
      - Non-retryable errors (400, 401, 403, 404, 409): Pass
      - Per-operation `CancellationToken` for timeout control: Pass
      - Geo-redundant retry with `GeoRedundantSecondaryUri`: Fail
      - Interaction with Polly or other resilience libraries: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
