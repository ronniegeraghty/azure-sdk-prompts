# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 46.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 46.9s |
| Timestamp | 2026-08-28T18:49:26Z |
| Files Generated | 0 |
| Event Count | 2110 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 26.0s |
| Review | 19.8s |
| **Total** | **46.9s** |

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
| Token Usage | in=19612 out=1289 |
| Turn Count | 1 |

## Copilot Response

Azure Blob Storage operations typically throw `Azure.RequestFailedException` for service errors. Its `Status` property contains the HTTP status, and `ErrorCode` contains the Storage service error code.

```csharp
using Azure;
using Azure.Storage.Blobs;

BlobContainerClient container = new(
    Environment.GetEnvironmentVariable("AZURE_STORAGE_CONNECTION_STRING")
        ?? throw new InvalidOperationException(
            "AZURE_STORAGE_CONNECTION_STRING is not configured."),
    "documents");

try
{
    BlobClient blob = container.GetBlobClient("report.pdf");

    await using FileStream destination = File.Create("report.pdf");
    await blob.DownloadToAsync(destination);
}
catch (RequestFailedException ex) when (ex.Status == 404)
{
    // Common codes: ContainerNotFound, BlobNotFound
    Console.Error.WriteLine(
        $"Resource not found. Status={ex.Status}, Code={ex.ErrorCode}, " +
        $"RequestId={ex.GetRawResponse()?.ClientRequestId}");
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // Common codes include AuthenticationFailed,
    // AuthorizationFailure, and AuthorizationPermissionMismatch.
    Console.Error.WriteLine(
        $"Access denied. Status={ex.Status}, Code={ex.ErrorCode}, " +
        $"Message={ex.Message}");
}
catch (RequestFailedException ex) when (ex.Status == 429)
{
    // The SDK normally retries retryable responses before this is thrown.
    Console.Error.WriteLine(
        $"Storage request was throttled after retries. " +
        $"Status={ex.Status}, Code={ex.ErrorCode}");
}
catch (RequestFailedException ex)
{
    // Other service responses, such as 409 conflicts or 5xx failures.
    Console.Error.WriteLine(
        $"Blob Storage failed. Status={ex.Status}, " +
        $"Code={ex.ErrorCode}, Message={ex.Message}");

    throw;
}
```

A reusable pattern is to preserve the exception while translating it into application-specific errors:

```csharp
using Azure;
using Azure.Storage.Blobs;

public async Task DownloadAsync(
    BlobContainerClient container,
    string blobName,
    Stream destination,
    CancellationToken cancellationToken = default)
{
    try
    {
        await container
            .GetBlobClient(blobName)
            .DownloadToAsync(destination, cancellationToken);
    }
    catch (RequestFailedException ex) when (
        ex.Status == 404 &&
        ex.ErrorCode is "ContainerNotFound" or "BlobNotFound")
    {
        throw new FileNotFoundException(
            $"Blob '{blobName}' was not found.", blobName, ex);
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        throw new UnauthorizedAccessException(
            $"Blob access was denied ({ex.ErrorCode}).", ex);
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        // Usually let this propagate so an outer resilience policy can
        // reschedule the operation rather than immediately retrying again.
        throw;
    }
}
```

The important properties are:

```csharp
catch (RequestFailedException ex)
{
    int httpStatus = ex.Status;       // 403, 404, 429, etc.
    string? storageCode = ex.ErrorCode;
    Response? response = ex.GetRawResponse();

    Console.Error.WriteLine(
        "HTTP {0}; Storage code {1}; Reason {2}",
        httpStatus,
        storageCode ?? "<none>",
        response?.ReasonPhrase);
}
```

Configure the SDK’s built-in retry policy when constructing the client:

```csharp
using Azure.Core;
using Azure.Storage.Blobs;

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(30),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(100)
    }
};

var serviceClient = new BlobServiceClient(
    Environment.GetEnvironmentVariable("AZURE_STORAGE_CONNECTION_STRING"),
    options);
```

A few behavioral details matter:

| Scenario | Typical result |
|---|---|
| Container missing | `RequestFailedException`, status `404`, code `ContainerNotFound` |
| Blob missing | `RequestFailedException`, status `404`, code `BlobNotFound` |
| Invalid credentials | Status `403`, often `AuthenticationFailed` |
| Valid identity lacking permission | Status `403`, often `AuthorizationPermissionMismatch` or `AuthorizationFailure` |
| Throttling | Status `429`; some Azure Storage overload conditions instead use `503` with `ServerBusy` |
| `ExistsAsync()` on a missing resource | Usually returns `false` rather than throwing for the expected `404` |

Catch by HTTP status for broad handling, and inspect `ErrorCode` when the distinction affects application behavior. Avoid parsing `ex.Message`; it is intended for diagnostics and is not a stable programmatic contract.

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response handles the requested core exceptions and retry configuration well, but lacks explicit 409 handling and conditional-request examples for ETags and leases.

### Strengths

- Uses RequestFailedException idiomatically with exception filters.
- Shows direct extraction of Status and ErrorCode without parsing messages.
- Includes practical BlobClientOptions retry configuration.
- Provides useful troubleshooting fields in diagnostic output.

### Issues

- No dedicated 409 conflict catch or behavior-specific handling.
- No coverage of ETag or lease-related conditional failures.

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
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
