# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 57.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 57.1s |
| Timestamp | 2026-08-26T18:27:13Z |
| Files Generated | 0 |
| Event Count | 2126 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 36.1s |
| Review | 19.8s |
| **Total** | **57.1s** |

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
| Token Usage | in=19605 out=1162 |
| Turn Count | 1 |

## Prompt Sent

```
How do I properly handle errors when working with Azure Blob Storage in .NET?
I need to understand what exceptions are thrown for common failure scenarios:
container not found (404), access denied (403), and throttling (429).
Show me idiomatic try/catch patterns with the Azure.Storage.Blobs SDK
including how to extract the error code and HTTP status from RequestFailedException.
```

## Copilot Response

Azure Blob Storage service failures are surfaced as `Azure.RequestFailedException`. Its key properties are:

- `Status`: HTTP status code, such as `404`, `403`, or `429`
- `ErrorCode`: storage service error code, such as `ContainerNotFound` or `AuthorizationPermissionMismatch`
- `Message`: SDK-generated diagnostic message
- `GetRawResponse()`: response headers, request ID, and other HTTP details

```csharp
using Azure;
using Azure.Storage.Blobs;

BlobContainerClient containerClient =
    new(connectionString, containerName);

try
{
    // GetPropertiesAsync forces a service request and throws if the
    // container does not exist or cannot be accessed.
    await containerClient.GetPropertiesAsync(cancellationToken: cancellationToken);
}
catch (RequestFailedException ex) when (ex.Status == 404)
{
    Console.Error.WriteLine(
        $"Container not found. Status={ex.Status}, ErrorCode={ex.ErrorCode}");
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    Console.Error.WriteLine(
        $"Access denied. Status={ex.Status}, ErrorCode={ex.ErrorCode}");
}
catch (RequestFailedException ex) when (ex.Status == 429)
{
    Console.Error.WriteLine(
        $"Request throttled. Status={ex.Status}, ErrorCode={ex.ErrorCode}");

    // Usually let the SDK retry policy handle this rather than implementing
    // an immediate retry loop here.
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Storage request failed. Status={ex.Status}, " +
        $"ErrorCode={ex.ErrorCode}, Message={ex.Message}");

    throw;
}
```

Common service codes vary with the operation and authorization mechanism:

| Scenario | Status | Typical `ErrorCode` values |
|---|---:|---|
| Container missing | 404 | `ContainerNotFound` |
| Access denied | 403 | `AuthorizationFailure`, `AuthorizationPermissionMismatch`, `AuthenticationFailed` |
| Throttling/service busy | 429 | Often `ServerBusy` |

Prefer status plus service error code when handling a specific condition:

```csharp
try
{
    await containerClient.DeleteAsync(cancellationToken: cancellationToken);
}
catch (RequestFailedException ex)
    when (ex.Status == 404 && ex.ErrorCode == "ContainerNotFound")
{
    // Desired end state already exists: the container is absent.
}
```

For reusable handling, switch on the exception:

```csharp
try
{
    BlobClient blobClient = containerClient.GetBlobClient(blobName);

    await blobClient.UploadAsync(
        BinaryData.FromString("content"),
        overwrite: true,
        cancellationToken);
}
catch (RequestFailedException ex)
{
    switch (ex.Status)
    {
        case 404:
            throw new InvalidOperationException(
                $"Storage resource was not found ({ex.ErrorCode}).", ex);

        case 403:
            throw new UnauthorizedAccessException(
                $"Storage authorization failed ({ex.ErrorCode}).", ex);

        case 429:
            // Preserve the original exception so the caller can decide
            // whether the operation is safe to retry.
            throw;

        default:
            throw;
    }
}
```

You can also inspect response metadata:

```csharp
catch (RequestFailedException ex)
{
    Response? response = ex.GetRawResponse();

    string? requestId = response?.Headers.TryGetValue(
        "x-ms-request-id",
        out string value) == true
            ? value
            : null;

    Console.Error.WriteLine(
        $"Status={ex.Status}, ErrorCode={ex.ErrorCode}, RequestId={requestId}");

    throw;
}
```

Configure retries through `BlobClientOptions`; transient errors such as `429`, `500`, `502`, `503`, and `504` are retried automatically:

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

var serviceClient = new BlobServiceClient(connectionString, options);
```

Two important details:

1. `ExistsAsync()` returns `false` for a normal `404` rather than requiring exception handling:

   ```csharp
   Response<bool> result =
       await containerClient.ExistsAsync(cancellationToken);

   if (!result.Value)
   {
       Console.WriteLine("Container does not exist.");
   }
   ```

   Other failures, including `403`, still throw `RequestFailedException`.

2. Do not treat every `403` as merely “insufficient permissions.” It can also indicate an invalid/expired SAS, authentication failure, network/firewall restrictions, or a role assignment that has not propagated. Log `ErrorCode` and `x-ms-request-id`, but avoid logging credentials, SAS tokens, or connection strings.

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response thoroughly covers common RequestFailedException handling, retry configuration, and diagnostics, but omits 409 handling and conditional failures involving ETags or leases.

### Strengths

- Idiomatic filtered catches and switch-based RequestFailedException handling.
- Clear extraction of HTTP status, service error code, and request ID.
- Appropriate BlobClientOptions retry-policy example and security-conscious logging guidance.

### Issues

- No specific handling or explanation for HTTP 409 conflicts.
- No examples or guidance for ETag precondition failures or lease-related errors.

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
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
