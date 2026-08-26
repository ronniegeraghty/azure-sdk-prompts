# Evaluation Report: storage-dp-dotnet-batch

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 124.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-batch` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 5/8 |
| Duration | 124.2s |
| Timestamp | 2026-08-26T18:21:38Z |
| Files Generated | 0 |
| Event Count | 3878 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 98.8s |
| Review | 24.3s |
| **Total** | **124.2s** |

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
| Token Usage | in=105247 out=2453 |
| Turn Count | 4 |

## Prompt Sent

```
How do I perform batch operations on Azure Blob Storage in .NET?
I have a container with thousands of blobs and need to:
1. Bulk delete 500 blobs in a single batched HTTP request
2. Bulk set the access tier of 200 blobs from Hot to Cool
3. Handle partial failures where some operations in the batch succeed and others fail
4. Understand the limits — max operations per batch and size restrictions

Show me how to use BlobBatchClient to submit batch requests with
the Azure.Storage.Blobs.Batch package. Include proper error handling
for partial batch failures.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure.Storage.Blobs.Batch .NET BlobBatchClient SubmitBatchAsync DeleteBlobs SetBlobsAccessTier partial failures limits maximum 256 operations 4 MB current documentation"}
```

### 🔧 2. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet"}
```

### 🔧 3. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.deleteblobsasync?view=azure-dotnet"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.setblobsaccesstierasync?view=azure-dotnet"}
```

### 🔧 6. web_fetch

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.setblobaccesstier?view=azure-dotnet"}
```

### 🔧 8. web_fetch

**Input:**

```json
{"max_length":16000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatch.cs"}
```

## Copilot Response

I’m checking the current .NET batch SDK API and service limits so the examples match the package’s present behavior.

The key constraint is 256 subrequests per batch, so 500 deletes require two requests; I’m confirming the response-inspection types for reliable partial-failure handling.

I’ll use the lower-level `CreateBatch`/`SubmitBatchAsync` API because it preserves one response object per blob, letting you map failures back to the original URI without relying on `AggregateException` ordering.

A Blob batch supports **at most 256 subrequests**, so 500 deletes cannot be sent in one HTTP request. Split them into **256 + 244**. The 200 tier changes fit in one batch.

Install the packages:

```powershell
dotnet add package Azure.Storage.Blobs.Batch
dotnet add package Azure.Identity
```

```csharp
using Azure;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

Uri serviceUri = new("https://<account-name>.blob.core.windows.net");
BlobServiceClient serviceClient =
    new(serviceUri, new DefaultAzureCredential());

BlobBatchClient batchClient = serviceClient.GetBlobBatchClient();
BlobContainerClient containerClient =
    serviceClient.GetBlobContainerClient("my-container");
```

Your identity needs data-plane permissions such as **Storage Blob Data Contributor**.

## Submit batches with per-blob failure reporting

The lower-level `CreateBatch`/`SubmitBatchAsync` API is useful because the `Response` returned when each operation is added can later be associated with that blob.

```csharp
public sealed record BatchFailure(
    Uri BlobUri,
    int? Status,
    string? ErrorCode,
    string Message);

public sealed record BatchResult(
    int Succeeded,
    IReadOnlyList<BatchFailure> Failures);

private const int MaximumBatchOperations = 256;

private static IEnumerable<IReadOnlyList<T>> Chunk<T>(
    IEnumerable<T> values,
    int size)
{
    List<T> chunk = new(size);

    foreach (T value in values)
    {
        chunk.Add(value);

        if (chunk.Count == size)
        {
            yield return chunk;
            chunk = new List<T>(size);
        }
    }

    if (chunk.Count > 0)
    {
        yield return chunk;
    }
}

private static BatchResult InspectResponses(
    IReadOnlyList<(Uri BlobUri, Response Response)> operations)
{
    int succeeded = 0;
    List<BatchFailure> failures = [];

    foreach ((Uri blobUri, Response response) in operations)
    {
        try
        {
            // Delayed responses become usable after SubmitBatchAsync completes.
            // Accessing the response throws RequestFailedException when that
            // individual subrequest failed.
            _ = response.Status;
            succeeded++;
        }
        catch (RequestFailedException ex)
        {
            failures.Add(new BatchFailure(
                blobUri,
                ex.Status,
                ex.ErrorCode,
                ex.Message));
        }
    }

    return new BatchResult(succeeded, failures);
}
```

### Delete 500 blobs

```csharp
public static async Task<BatchResult> DeleteBlobsInBatchesAsync(
    BlobBatchClient batchClient,
    IEnumerable<Uri> blobUris,
    CancellationToken cancellationToken = default)
{
    int succeeded = 0;
    List<BatchFailure> failures = [];

    foreach (IReadOnlyList<Uri> chunk in
             Chunk(blobUris, MaximumBatchOperations))
    {
        using BlobBatch batch = batchClient.CreateBatch();

        List<(Uri BlobUri, Response Response)> operations =
            new(chunk.Count);

        foreach (Uri blobUri in chunk)
        {
            Response response = batch.DeleteBlob(
                blobUri,
                DeleteSnapshotsOption.IncludeSnapshots);

            operations.Add((blobUri, response));
        }

        try
        {
            await batchClient.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);
        }
        catch (RequestFailedException ex)
        {
            // The outer batch request failed. Do not assume any operation's
            // outcome; reconcile or retry this chunk according to your policy.
            throw new InvalidOperationException(
                $"Batch submission failed for {chunk.Count} blobs. " +
                $"Status={ex.Status}, ErrorCode={ex.ErrorCode}",
                ex);
        }

        BatchResult result = InspectResponses(operations);
        succeeded += result.Succeeded;
        failures.AddRange(result.Failures);
    }

    return new BatchResult(succeeded, failures);
}
```

Usage:

```csharp
List<string> namesToDelete = GetBlobNamesToDelete();

List<Uri> deleteUris = namesToDelete
    .Select(name => containerClient.GetBlobClient(name).Uri)
    .ToList();

BatchResult deleteResult = await DeleteBlobsInBatchesAsync(
    batchClient,
    deleteUris);

Console.WriteLine(
    $"Deleted: {deleteResult.Succeeded}; failed: {deleteResult.Failures.Count}");

foreach (BatchFailure failure in deleteResult.Failures)
{
    Console.Error.WriteLine(
        $"{failure.BlobUri}: HTTP {failure.Status}, " +
        $"{failure.ErrorCode}: {failure.Message}");
}
```

This makes **two batch HTTP requests**, not one:

| Request | Deletes |
|---|---:|
| Batch 1 | 256 |
| Batch 2 | 244 |

If deleting a nonexistent blob should count as success, handle status `404`/error code `BlobNotFound` accordingly when processing failures.

### Set 200 blobs from Hot to Cool

A batch must contain only one operation type, so tier changes cannot be mixed with deletes in the same batch.

```csharp
public static async Task<BatchResult> SetAccessTierInBatchesAsync(
    BlobBatchClient batchClient,
    IEnumerable<Uri> blobUris,
    AccessTier accessTier,
    CancellationToken cancellationToken = default)
{
    int succeeded = 0;
    List<BatchFailure> failures = [];

    foreach (IReadOnlyList<Uri> chunk in
             Chunk(blobUris, MaximumBatchOperations))
    {
        using BlobBatch batch = batchClient.CreateBatch();

        List<(Uri BlobUri, Response Response)> operations =
            new(chunk.Count);

        foreach (Uri blobUri in chunk)
        {
            Response response =
                batch.SetBlobAccessTier(blobUri, accessTier);

            operations.Add((blobUri, response));
        }

        try
        {
            await batchClient.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);
        }
        catch (RequestFailedException ex)
        {
            throw new InvalidOperationException(
                $"Tier batch submission failed for {chunk.Count} blobs. " +
                $"Status={ex.Status}, ErrorCode={ex.ErrorCode}",
                ex);
        }

        BatchResult result = InspectResponses(operations);
        succeeded += result.Succeeded;
        failures.AddRange(result.Failures);
    }

    return new BatchResult(succeeded, failures);
}
```

Usage:

```csharp
List<string> namesToCool = GetBlobNamesToCool();

List<Uri> tierUris = namesToCool
    .Select(name => containerClient.GetBlobClient(name).Uri)
    .ToList();

BatchResult tierResult = await SetAccessTierInBatchesAsync(
    batchClient,
    tierUris,
    AccessTier.Cool);
```

## Failure behavior and limits

| Constraint | Behavior |
|---|---|
| Maximum subrequests | **256 per batch** |
| Maximum batch body | **4 MiB** |
| Supported operations | Delete blob and set block-blob access tier |
| Mixed operation types | Not allowed; one batch must be all deletes or all tier changes |
| Storage-account scope | Every subrequest must target the same storage account |
| Atomicity | None; each operation succeeds or fails independently |
| Ordering | Execution order is not guaranteed |
| Batch reuse | A submitted `BlobBatch` cannot be submitted again |
| Outer request failure | `SubmitBatchAsync` throws `RequestFailedException` |
| Individual failure with `throwOnAnyFailure: false` | Submission returns; inspect each saved response |
| Individual failure with `throwOnAnyFailure: true` | Failures are surfaced in an `AggregateException` |

Using `throwOnAnyFailure: false` is preferable when you need an exact blob-to-result mapping. Also make retries operation-aware: deleting a blob is generally retry-friendly, while an outer batch failure can leave outcomes unknown because the server may have processed some subrequests before the connection failed.

References: [Azure Storage Blobs Batch client library](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs.batch-readme) and [`SubmitBatchAsync`](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 8 criteria passed. The response provides a strong custom-batch implementation with accurate limits and detailed per-blob failure mapping, but omits both requested convenience APIs and does not demonstrate AggregateException-based partial-failure handling.

### Strengths

- Correctly rejects the impossible request to place 500 deletes in one batch and chunks them at the 256-operation limit.
- Provides complete CreateBatch()/SubmitBatchAsync() examples for both supported operation types.
- Maps individual RequestFailedException details back to each blob when throwOnAnyFailure is false.
- Explains relevant body-size, atomicity, operation-mixing, account-scope, and authentication constraints.

### Issues

- Missing BlobBatchClient.DeleteBlobsAsync() example.
- Missing BlobBatchClient.SetBlobsAccessTierAsync() example.
- AggregateException behavior is stated but not handled to report per-operation status.

## Grader Results

- batch-blob-operations.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/8)
      - `BlobBatchClient` from `Azure.Storage.Blobs.Batch` package: Pass
      - `BlobBatchClient.DeleteBlobsAsync()` for bulk delete: Fail
      - `BlobBatchClient.SetBlobsAccessTierAsync()` for bulk tier changes: Fail
      - Custom batch via `BlobBatchClient.CreateBatch()` and `SubmitBatchAsync()`: Pass
      - Batch size limits (256 operations per batch): Pass
      - Partial failure handling: `AggregateException` with per-operation status: Fail
      - `RequestFailedException` for individual operation failures within a batch: Pass
      - Authentication scopes for batch operations: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 62% | 1.00 | 0.6250 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6250** | **62.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-batch --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
