# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 177.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 177.6s |
| Timestamp | 2026-08-26T18:16:07Z |
| Files Generated | 0 |
| Event Count | 3754 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 147.3s |
| Review | 29.2s |
| **Total** | **177.6s** |

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
| Token Usage | in=412719 out=3633 |
| Turn Count | 10 |

## Prompt Sent

```
How do I properly handle errors when working with Azure Blob Storage in .NET?
I need to understand what exceptions are thrown for common failure scenarios:
container not found (404), access denied (403), and throttling (429).
Show me idiomatic try/catch patterns with the Azure.Storage.Blobs SDK
including how to extract the error code and HTTP status from RequestFailedException.
```

## Tool Calls

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide idiomatic .NET error-handling examples for Azure.Storage.Blobs, covering HTTP 404 container not found, 403 access denied, 429 throttling, RequestFailedException.Status and ErrorCode extraction, and retry guidance. No Azure resources will be modified.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET SDK code-generation best practices for Azure Blob Storage exception handling.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Storage.Blobs .NET error handling, RequestFailedException Status and ErrorCode, Blob service error codes, and Azure SDK retry policies.","learn":true}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Locate official docs for RequestFailedException fields and Blob Storage error/retry behavior.","parameters":{"query":"Azure.Storage.Blobs .NET RequestFailedException Status ErrorCode 403 404 429 retry BlobErrorCode ContainerNotFound AuthorizationFailure ServerBusy"}}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# exception handling samples for Azure.Storage.Blobs.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs C# catch RequestFailedException Status ErrorCode BlobErrorCode ContainerNotFound AuthorizationFailure throttling 429 retry"}}
```

### 🔧 6. rg

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787768214143-copilot-tool-output-59fc5d8125e44fcd82183291664be9c7.txt","pattern":"\"title\"|\"url\"|\"content\""}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787768214143-copilot-tool-output-59fc5d8125e44fcd82183291664be9c7.txt' | ConvertFrom-Json; $j.results | Select-Object title,contentUrl | Format-Table -AutoSize","description":"List documentation result URLs"}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"$line = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787768214143-copilot-tool-output-59fc5d8125e44fcd82183291664be9c7.txt' -First 1; $j = $line | ConvertFrom-Json; $j.results | Select-Object title,contentUrl | Format-Table -AutoSize","description":"Parse documentation result URLs"}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"$line = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787768214143-copilot-tool-output-59fc5d8125e44fcd82183291664be9c7.txt' -First 1; $j = $line | ConvertFrom-Json; $j.results | ForEach-Object { \"TITLE=$($_.title)`nURL=$($_.contentUrl)`n\" }","description":"Print documentation result URLs"}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the current Azure Storage Blobs .NET troubleshooting guidance.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs-readme?view=azure-dotnet#troubleshooting"}}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Azure Storage common REST error codes.","parameters":{"url":"https://learn.microsoft.com/rest/api/storageservices/common-rest-api-error-codes"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official troubleshooting guidance for 403 Blob Storage errors and service error codes.","parameters":{"url":"https://learn.microsoft.com/troubleshoot/azure/azure-storage/blobs/authentication/storage-troubleshoot-403-errors#common-403-errors-related-to-authorization-failures"}}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Confirm the Azure SDK for .NET default retry status codes and Blob Storage throttling response behavior.","parameters":{"query":"Azure SDK .NET RetryOptions default HTTP status codes 408 429 500 502 503 504 Azure.Core Blob Storage throttling ServerBusy"}}
```

## Copilot Response

`Azure.Storage.Blobs` reports service failures as `Azure.RequestFailedException`. Its two most useful properties are:

- `Status`: HTTP status code, such as `404`, `403`, or `429`.
- `ErrorCode`: Azure Storage’s machine-readable service code, such as `ContainerNotFound` or `AuthorizationPermissionMismatch`.

```csharp
using Azure;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Microsoft.Extensions.Logging;

public static async Task<bool> DownloadAsync(
    BlobContainerClient container,
    string blobName,
    Stream destination,
    ILogger logger,
    CancellationToken cancellationToken = default)
{
    try
    {
        BlobClient blob = container.GetBlobClient(blobName);
        await blob.DownloadToAsync(destination, cancellationToken);
        return true;
    }
    catch (RequestFailedException ex)
        when (ex.Status == 404 &&
              ex.ErrorCode == BlobErrorCode.ContainerNotFound)
    {
        logger.LogInformation(
            "Container {Container} was not found. Status={Status}, ErrorCode={ErrorCode}",
            container.Name,
            ex.Status,
            ex.ErrorCode);

        return false;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        // Do not translate 403 into "not found": the resource might exist,
        // but the caller is not authorized to access it.
        logger.LogWarning(
            ex,
            "Blob access denied. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        throw;
    }
    catch (RequestFailedException ex)
        when (ex.Status is 429 or 503)
    {
        // Azure SDK retries transient responses before this catch executes.
        // Blob Storage frequently reports account throttling as 503/ServerBusy,
        // although 429 can also be returned by a service or intermediary.
        logger.LogWarning(
            ex,
            "Blob request was throttled after SDK retries. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        throw;
    }
    catch (RequestFailedException ex)
    {
        logger.LogError(
            ex,
            "Blob operation failed. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        throw;
    }
}
```

Typical values are:

| Scenario | `Status` | Common `ErrorCode` |
|---|---:|---|
| Container missing | 404 | `ContainerNotFound` |
| Blob missing | 404 | `BlobNotFound` |
| Permission missing | 403 | `AuthorizationPermissionMismatch` |
| Authentication/signature failure | 403 | `AuthenticationFailed` |
| Firewall/network restriction | 403 | `AuthorizationFailure`, `IpAuthorizationFailure` |
| Throttling | 429 | Varies by response |
| Blob account throughput exceeded | 503 | `ServerBusy` |

Because multiple conditions can produce the same HTTP status, use **both** properties when the distinction matters:

```csharp
catch (RequestFailedException ex) when (
    ex.Status == 404 &&
    ex.ErrorCode == BlobErrorCode.ContainerNotFound)
{
    // Specifically the container, rather than the blob, is missing.
}
```

For diagnostic data:

```csharp
catch (RequestFailedException ex)
{
    int status = ex.Status;
    string? storageErrorCode = ex.ErrorCode;
    string message = ex.Message;

    Response response = ex.GetRawResponse();

    response.Headers.TryGetValue("x-ms-request-id", out string? serviceRequestId);
    response.Headers.TryGetValue("x-ms-client-request-id", out string? clientRequestId);

    logger.LogError(
        ex,
        "Storage failure: Status={Status}, ErrorCode={ErrorCode}, " +
        "ServiceRequestId={ServiceRequestId}, ClientRequestId={ClientRequestId}",
        status,
        storageErrorCode,
        serviceRequestId,
        clientRequestId);

    throw;
}
```

Configure retries on the client rather than placing an unconditional retry loop around every operation:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(10),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(100)
    }
};

var serviceClient = new BlobServiceClient(
    new Uri("https://<storage-account>.blob.core.windows.net"),
    new DefaultAzureCredential(),
    options);
```

The SDK retry policy handles transient responses such as `429` and `503`, including server-provided retry delays. If a throttling exception reaches your catch block, the configured SDK retries have already been exhausted. Only add an application-level retry when the operation is safe to repeat, and avoid multiplying application retries by excessive SDK retries.

For simple existence checks, prefer APIs that represent absence without exceptions:

```csharp
Response<bool> exists = await container.ExistsAsync(cancellationToken);

if (!exists.Value)
{
    // Container does not exist.
}
```

`ExistsAsync` still throws for failures such as `403`; it only converts the expected not-found response into `false`.

References:

- [Azure Storage Blobs client library for .NET—troubleshooting](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs-readme?view=azure-dotnet#troubleshooting)
- [Azure Storage Blob service error codes](https://learn.microsoft.com/rest/api/storageservices/blob-service-error-codes)
- [Azure Storage common REST error codes](https://learn.microsoft.com/rest/api/storageservices/common-rest-api-error-codes)
- [Troubleshoot Azure Blob Storage 403 errors](https://learn.microsoft.com/troubleshoot/azure/azure-storage/blobs/authentication/storage-troubleshoot-403-errors)

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response thoroughly covers RequestFailedException, status/error extraction, retries, and diagnostics, but omits 409 conflict handling and conditional failures involving ETags and leases.

### Strengths

- Uses RequestFailedException consistently and preserves failures with throw.
- Shows idiomatic exception filters for 404, 403, and throttling responses.
- Provides a complete BlobClientOptions retry configuration.
- Includes useful structured diagnostic logging with x-ms request identifiers.

### Issues

- No handling example or explanation for HTTP 409 conflicts.
- No discussion of ETag precondition failures, lease-related failures, or their relevant status/error codes.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/6)
      - `RequestFailedException` as the primary exception type: Pass
      - Extracting `Status` and `ErrorCode` from the exception: Pass
      - Handling specific HTTP status codes (404, 403, 409, 429): Fail
      - Retry policy configuration via `BlobClientOptions`: Pass
      - Conditional request failures (ETags, leases): Fail
      - Logging and diagnostics for troubleshooting: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 67% | 1.00 | 0.6667 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6667** | **66.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
