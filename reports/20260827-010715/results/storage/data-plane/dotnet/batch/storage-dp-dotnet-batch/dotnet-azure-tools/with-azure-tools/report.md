# Evaluation Report: storage-dp-dotnet-batch

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 211.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-batch` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 5/8 |
| Duration | 211.0s |
| Timestamp | 2026-08-26T18:23:42Z |
| Files Generated | 0 |
| Event Count | 3530 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 166.2s |
| Review | 43.7s |
| **Total** | **211.0s** |

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
| Token Usage | in=347560 out=4900 |
| Turn Count | 8 |

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

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide modern .NET sample code for Azure Blob Storage batch delete and set-tier operations using Azure.Storage.Blobs.Batch, including partial-failure handling and service limits.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET code-generation best practices for Blob Storage batch operations.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Azure documentation for Azure.Storage.Blobs.Batch .NET BlobBatchClient, including SubmitBatchAsync partial failures, DeleteBlobsAsync, SetBlobsAccessTierAsync, 256-operation and 4-MiB limits.","learn":true}
```

### 🔧 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Azure.Storage.Blobs.Batch BlobBatchClient .NET SubmitBatchAsync partial failures 256 4 MiB"}
```

### 🔧 5. web_search (via github-mcp-server)

**Input:**

```json
{"query":"latest stable NuGet version Azure.Storage.Blobs.Batch August 2026"}
```

### 🔧 6. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha"],"page":1,"perPage":10,"query":"\"class BlobBatchClient\" repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch"}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Locate official BlobBatchClient .NET API docs and service limitations.","parameters":{"query":"Azure.Storage.Blobs.Batch BlobBatchClient SubmitBatchAsync throwOnAnyFailure partial failure 256 4 MiB .NET"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# examples for BlobBatchClient batch delete, set access tier, and partial failures.","parameters":{"language":"csharp","query":"BlobBatchClient CreateBatch SubmitBatchAsync throwOnAnyFailure DeleteBlob SetBlobAccessTier"}}
```

### 🔧 9. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs","repo":"azure-sdk-for-net"}
```

### 🔧 10. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/api/Azure.Storage.Blobs.Batch.net8.0.cs","repo":"azure-sdk-for-net"}
```

### 🔧 11. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/README.md","repo":"azure-sdk-for-net"}
```

### 🔧 12. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"page":1,"perPage":30,"query":"\"throwOnAnyFailure\" repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch"}
```

### 🔧 13. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs","repo":"azure-sdk-for-net"}
```

### 🔧 14. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"page":1,"perPage":10,"query":"\"AggregateException exes\" repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs"}
```

### 🔧 15. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"page":1,"perPage":20,"query":"\"DeleteBlobsAsync\" \"AggregateException\" repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch"}
```

### 🔧 16. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official BlobBatchClient SubmitBatchAsync API documentation.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet"}}
```

### 🔧 17. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"page":1,"perPage":20,"query":"\"same operation\" repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch/src"}
```

### 🔧 18. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","text_matches"],"page":1,"perPage":20,"query":"\"Cannot mix\" repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch"}
```

### 🔧 19. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Verify Azure Blob Batch REST restrictions such as whether operations can be mixed in one request, same-account rules, payload and operation limits.","parameters":{"query":"Azure Blob Batch request same operation type cannot mix Delete Blob Set Blob Tier 256 4 MiB"}}
```

### 🔧 20. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Blob Batch REST API restrictions and failure semantics.","parameters":{"url":"https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch"}}
```

## Copilot Response

I’m checking the current SDK signatures and service limits so the samples accurately show how oversized lists are chunked and how per-subrequest failures are surfaced.

You **cannot delete 500 blobs in one batch request**. Azure Blob Batch allows at most **256 subrequests**, so 500 deletes require two HTTP batch requests: **256 + 244**. The 200 tier changes fit in one request, provided the serialized request remains under 4 MB.

```powershell
dotnet add package Azure.Storage.Blobs.Batch --version 12.26.0
dotnet add package Azure.Identity
```

```csharp
using Azure;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

const int MaxBatchOperations = 256;

string accountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME")
    ?? throw new InvalidOperationException(
        "Set AZURE_STORAGE_ACCOUNT_NAME.");

var serviceClient = new BlobServiceClient(
    new Uri($"https://{accountName}.blob.core.windows.net"),
    new DefaultAzureCredential());

BlobBatchClient batchClient = serviceClient.GetBlobBatchClient();

// These can come from GetBlobsAsync(), a database, or another source.
IEnumerable<Uri> blobsToDelete = GetDeleteUris();
IEnumerable<Uri> blobsToCool = GetTierChangeUris();

// 500 blobs are automatically submitted as 256 + 244.
IReadOnlyList<BatchItemResult> deleteResults =
    await SubmitInChunksAsync(
        batchClient,
        blobsToDelete,
        static (batch, uri) =>
            batch.DeleteBlob(
                uri,
                DeleteSnapshotsOption.IncludeSnapshots));

// 200 blobs are submitted in one batch, assuming the request is under 4 MB.
IReadOnlyList<BatchItemResult> tierResults =
    await SubmitInChunksAsync(
        batchClient,
        blobsToCool,
        static (batch, uri) =>
            batch.SetBlobAccessTier(uri, AccessTier.Cool));

Report("delete", deleteResults);
Report("set-tier", tierResults);

static async Task<IReadOnlyList<BatchItemResult>> SubmitInChunksAsync(
    BlobBatchClient client,
    IEnumerable<Uri> blobUris,
    Func<BlobBatch, Uri, Response> addOperation,
    CancellationToken cancellationToken = default)
{
    var results = new List<BatchItemResult>();

    foreach (Uri[] chunk in blobUris.Chunk(MaxBatchOperations))
    {
        using BlobBatch batch = client.CreateBatch();

        // These Response objects are populated when SubmitBatchAsync completes.
        PendingOperation[] pending = chunk
            .Select(uri => new PendingOperation(
                uri,
                addOperation(batch, uri)))
            .ToArray();

        try
        {
            // False prevents a partial failure from hiding the successful
            // subrequests behind an AggregateException.
            await client.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);
        }
        catch (RequestFailedException ex)
        {
            // The parent batch request itself failed, for example because of
            // authentication, malformed content, throttling, or request size.
            Console.Error.WriteLine(
                $"Batch submission failed: HTTP {ex.Status}, " +
                $"code={ex.ErrorCode}, message={ex.Message}");

            throw;
        }

        foreach (PendingOperation operation in pending)
        {
            Response response = operation.Response;
            bool succeeded = response.Status is >= 200 and < 300;

            string? errorCode = null;
            if (!succeeded &&
                response.Headers.TryGetValue(
                    "x-ms-error-code",
                    out string returnedErrorCode))
            {
                errorCode = returnedErrorCode;
            }

            results.Add(new BatchItemResult(
                operation.BlobUri,
                succeeded,
                response.Status,
                errorCode,
                response.ReasonPhrase));
        }
    }

    return results;
}

static void Report(
    string operationName,
    IEnumerable<BatchItemResult> results)
{
    foreach (BatchItemResult failure in results.Where(r => !r.Succeeded))
    {
        Console.Error.WriteLine(
            $"{operationName} failed for {failure.BlobUri}: " +
            $"HTTP {failure.Status}, code={failure.ErrorCode ?? "unknown"}, " +
            $"reason={failure.ReasonPhrase}");
    }
}

static IEnumerable<Uri> GetDeleteUris()
{
    // Example only. Supply blob URIs belonging to the same storage account.
    return Array.Empty<Uri>();
}

static IEnumerable<Uri> GetTierChangeUris()
{
    return Array.Empty<Uri>();
}

internal sealed record PendingOperation(
    Uri BlobUri,
    Response Response);

internal sealed record BatchItemResult(
    Uri BlobUri,
    bool Succeeded,
    int Status,
    string? ErrorCode,
    string ReasonPhrase);
```

### Failure behavior

A batch is **not atomic**. For example, deleting 256 blobs may return 250 successful `202` responses and six `404 BlobNotFound` responses.

- `throwOnAnyFailure: false`: the parent request succeeds, and you inspect every captured `Response`. This is preferable when you need a URI-to-result mapping.
- `throwOnAnyFailure: true`: failed subrequests are collected into an `AggregateException`; its inner exceptions are generally `RequestFailedException` instances.
- `RequestFailedException` directly from `SubmitBatchAsync`: the parent batch request failed, so individual results should not be treated as completed.

The convenience methods `DeleteBlobsAsync` and `SetBlobsAccessTierAsync` return `Response[]`, but throw `AggregateException` when any subrequest fails. Explicit `CreateBatch`/`SubmitBatchAsync` provides better control over partial failures.

### Limits and restrictions

| Restriction | Value |
|---|---|
| Maximum subrequests | **256 per HTTP batch request** |
| Maximum request body | **4 MB**, including multipart framing, URLs, and headers |
| Supported operations | `Delete Blob` and `Set Blob Tier` |
| Mixing operation types | Not allowed; all subrequests in one batch must have the same type |
| Storage-account scope | Every subrequest must target the same storage account |
| Execution order | Not guaranteed |
| Atomicity | None; every subrequest executes independently |
| Tier support | `Set Blob Tier` applies to block blobs |
| Billing | The parent request and individual subrequests are billed as transactions |

If long blob URLs or request-condition headers push a 256-item request over 4 MB, use a smaller chunk size. Also reuse `BlobServiceClient` and `BlobBatchClient`; Azure SDK clients are thread-safe.

Authentication above uses `DefaultAzureCredential`, suitable for managed identity in Azure and developer credentials locally. The identity normally needs **Storage Blob Data Contributor** or equivalent least-privilege data-plane permissions.

**References:** [Azure.Storage.Blobs.Batch package documentation](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet), [SubmitBatchAsync API](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet), and [Blob Batch REST limits](https://learn.microsoft.com/rest/api/storageservices/blob-batch).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 8 criteria passed. The response gives a strong custom-batch implementation and accurate limits, but does not demonstrate either requested convenience method or AggregateException-based per-operation handling.

### Strengths

- Correctly rejects the impossible 500-operation single request and chunks it into 256 and 244.
- Provides robust CreateBatch()/SubmitBatchAsync() partial-result handling with URI-to-response mapping.
- Accurately documents the 256-operation and 4 MB limits, supported operations, non-atomic behavior, and account scope.
- Uses modern token-based authentication with DefaultAzureCredential and appropriate RBAC guidance.

### Issues

- No actual DeleteBlobsAsync() example.
- No actual SetBlobsAccessTierAsync() example.
- No AggregateException catch that inspects per-operation RequestFailedException status values.

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
hyoka run --prompt-id storage-dp-dotnet-batch --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
