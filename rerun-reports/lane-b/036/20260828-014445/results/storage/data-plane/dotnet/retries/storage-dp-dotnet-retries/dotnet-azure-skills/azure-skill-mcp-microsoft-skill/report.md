# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 584.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 0/8 |
| Duration | 584.0s |
| Timestamp | 2026-08-27T17:44:45Z |
| Files Generated | 0 |
| Event Count | 4106 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 548.7s |
| Review | 34.1s |
| **Total** | **584.0s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=492954 out=6595 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 3419ms

**Input:**

```json
{"skill":"azure-storage"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Services

## Services

| Service | Use When | MCP Tools | CLI |
|---------|----------|-----------|-----|
| Blob Storage | Objects, files, backups, static content | `azure__storage` | `az storage blob` |
| File Shares | SMB file shares, lift-and-shift | - | `az storage file` |
| Queue Storage | Async messaging, task queues | - | `az storage queue` |
| Table Storage | NoSQL key-value (consider Cosmos DB) | - | `az storage table` |
| Data Lake | Big data analytics, hierarchical namespace | - | `az storage fs` |

## MCP Server (Preferred)

When Azure MCP is enabled:

- `azure__storage` with command `storage_account_list` - List storage accounts
- `azure__storage` with command `storage_container_list` - List containers in account
- `azure__storage` with command `storage_blob_list` - List blobs in container
- `azure__storage` with command `storage_blob_get` - Download blob content
- `azure__storage` with command `storage_blob_put` - Upload blob content

**If Azure MCP is not enabled:** Run `/azure:setup` or enable via `/mcp`.

## CLI Fallback

```bash
# List storage accounts
az storage account list --output table

# List containers
az storage container list --account-name ACCOUNT --output table

# List blobs
az storage blob list --account-name ACCOUNT --container-name CONTAINER --output table

# Download blob
az storage blob download --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH

# Upload blob
az storage blob upload --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH
```

## Storage Account Tiers

| Tier | Use Case | Performance |
|------|----------|-------------|
| Standard | General purpose, backup | Milliseconds |
| Premium | Databases, high IOPS | Sub-millisecond |

## Blob Access Tiers

| Tier | Access Frequency | Cost |
|------|-----------------|------|
| Hot | Frequent | Higher storage, lower access |
| Cool | Infrequent (30+ days) | Lower storage, higher ac
... (truncated)
```

### ✅ 2. glob — 189172ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 189893ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
@@ -1,115 +1,115 @@
 # Azure Storage SDK Usage
 
 SDK packages and quick start examples for Azure Storage services.
 
 ## Storage SDKs by Language
 
 | Language | Blob | Queue | File Share | Data Lake |
 |----------|------|-------|------------|----------|
 | .NET | `Azure.Storage.Blobs` | `Azure.Storage.Queues` | `Azure.Storage.Files.Shares` | `Azure.Storage.Files.DataLake` |
 | Java | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | JavaScript | `@azure/storage-blob` | `@azure/storage-queue` | `@azure/storage-file-share` | `@azure/storage-file-datalake` |
 | Python | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | Go | `azblob` | `azqueue` | `azfile` | `azdatalake` |
 | Rust | `azure_storage_blob` | `azure_storage_queue` | - | - |
 
 ## Installation Commands
 
 | Language | Install Blob SDK + Identity |
 |----------|-----------------------------|
 | .NET | `dotnet add package Azure.Storage.Blobs` `dotnet add package Azure.Identity` |
 | Java | Maven: `com.azure:azure-storage-blob` `com.azure:azure-identity` |
 | JavaScript | `npm install @azure/storage-blob @azure/identity` |
 | Python | `pip install azure-storage-blob azure-identity` |
 | Go | `go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob github.com/Azure/azure-sdk-for-g
... (truncated)
```

### ✅ 4. view — 189526ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks environment variables, CLI tokens, and oth
... (truncated)
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 189896ms

**Input:**

```json
{"intent":"Provide current .NET Azure.Storage.Blobs code patterns for retries, per-operation cancellation timeouts, error classification, custom per-operation retry policies, and circuit breakers. No deployment or live resource changes.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 189898ms

**Input:**

```json
{"intent":"Find official documentation for Azure.Storage.Blobs BlobClientOptions RetryOptions in .NET, including default retryable HTTP status codes, Retry-After handling, NetworkTimeout behavior, cancellation tokens, and custom HttpPipelinePolicy support.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. web_search (via github-mcp-server) — 149439ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure.Core RetryOptions MaxRetries Delay MaxDelay NetworkTimeout RetryMode documentation"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 8. web_search (via github-mcp-server) — 149084ms

**Input:**

```json
{"query":"site:github.com/Azure/azure-sdk-for-net Azure.Core ResponseClassifier IsRetriableResponse status 408 429 500 502 503 504"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To implement circuit breaker resilience for Azure Blob Storage in a .NET application, you typically use the Polly library, which is a popular choice for transient fault handling and resilience strategies in .NET. While Polly is often discussed in the context of HTTP clients, its patterns—including retry, circuit breaker, timeout, and fallback—can also be applied when interacting with services like Azure Blob Storage, though with some distinctions.\n\n### Key Points From Microsoft Docs:\n\n1. **Resilience with Polly**:\n   - Polly enables you to define policies like Retry and Circuit Breaker, both of which can help when accessing cloud resources like Azure Blob Storage that may experience transient faults【3:2†source】.\n   - Circuit Breaker specifically stops calls from flooding a failing service, preventing resource exhaustion and allowing the service time to recover【3:0†source】【3:7†source】.\n\n2. **Azure SDK & Polly Integration**:\n   - For HTTP clients, it’s straightforward to use Polly via the `IHttpClientFactory` integration in .NET Core.\n   - For Azure SDKs that don’t natively use `HttpClient` (like the Azure.Storage.Blobs library), you'll need to apply Polly at the logical operation level—i.e., wrap your Blob Storage calls with your Polly policies.\n   - The .NET Azure SDK provides its own retry logic (see `RetryOptions`), but does not natively support circuit breaker within the SDK. So, to add circuit breaker, you have to wrap calls to the SDK within your own circuit breaker policy code.\n\n3. **Blob Storage Retry Example**:\n   - The Azure Blob Storage .NET client (v12 and above) supports customizable retry policies via `BlobClientOptions`, but this is only for retries, not full circuit breaker logic【3:3†source】.\n   - Example of configuring retries:\n     ```csharp\n     var options = new BlobClientOptions()\n     {\n         Retry =\n         {\n             MaxRetries = 5,\n         
... (truncated)
```

### ✅ 9. web_search (via github-mcp-server) — 148726ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Storage .NET retry policy BlobClientOptions Retry NetworkTimeout cancellation token custom policy"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s how Azure Storage .NET SDK lets you configure retry policy, network timeouts, and custom retry logic with `BlobClientOptions`:\n\n### 1. **Retry Policy with BlobClientOptions**\n\n- The Azure.Storage.Blobs library uses `BlobClientOptions` to let you configure retry logic for all blob requests.\n- Retry configuration is set through the `.Retry` property (of type `RetryOptions`), which includes:\n  - **Delay**: Time delay between retries (fixed or exponential).\n  - **MaxDelay**: Maximum time between retries.\n  - **MaxRetries**: How many times to retry before failing.\n  - **Mode**: Exponential or fixed policy.\n  - **NetworkTimeout**: Timeout for each individual network operation—not the total time for all retries, but each attempt's timeout【3:0†source】【3:3†source】.\n\n**Example:**\n```csharp\nvar options = new BlobClientOptions()\n{\n    Retry =\n    {\n        Mode = RetryMode.Exponential,\n        Delay = TimeSpan.FromSeconds(1),\n        MaxDelay = TimeSpan.FromSeconds(10),\n        MaxRetries = 3,\n        NetworkTimeout = TimeSpan.FromSeconds(30)\n    }\n};\n\nvar client = new BlobServiceClient(connectionString, options);\n```\nMore info: [Configure retry options](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy-dotnet)【3:0†source】\n\n### 2. **CancellationToken Support**\n\n- Almost all async APIs (like `UploadAsync`, `DownloadAsync`) accept a `CancellationToken`.\n- This lets your application cancel operations (for example, if a user cancels a file upload or your app needs to abort due to a timeout).\n```csharp\nusing var cts = new CancellationTokenSource();\nawait blobClient.UploadAsync(stream, cancellationToken: cts.Token);\n```\n[See BlobClient docs](https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.blobclient)【3:8†source】\n\n### 3. **Custom Retry Policy**\n\n- If you want to override default retry behaviors, derive from `RetryPolicy` (in Azur
... (truncated)
```

### ✅ 10. web_search (via github-mcp-server) — 148365ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure.Core.Pipeline.HttpPipelinePolicy PerRetry custom policy BlobClientOptions AddPolicy"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Yes, in Azure.Core for the .NET SDK, the default retry policy (which relies on `ResponseClassifier.IsRetriableResponse`) treats the following HTTP status codes as retriable: **408 (Request Timeout), 429 (Too Many Requests), 500 (Internal Server Error), 502 (Bad Gateway), 503 (Service Unavailable), and 504 (Gateway Timeout)**. This means that if a response returns with one of these status codes, the SDK's retry logic will typically attempt the request again, subject to retry settings such as maximum retry count and backoff strategy【3:7†source】.\n\nYou can find the logic for this in the `ResponseClassifier.cs` file in the Azure SDK for .NET repository. The relevant method (`IsRetriableResponse`) is responsible for identifying if a given HTTP response should be retried based on its status code【3:2†source】【3:1†source】.\n\nThis default set of retriable status codes aligns with best practices for handling transient faults in distributed systems.","annotations":[{"text":"","start_index":536,"end_index":548,"url_citation":{"title":"HTTP pipelines and retries in the Azure SDK for C++","url":"https://learn.microsoft.com/en-us/azure/developer/cpp/sdk/fundamentals/http-pipelines-and-retries"}},{"text":"","start_index":799,"end_index":811,"url_citation":{"title":"azure-sdk-for-net/sdk/core/Azure.Core/src/ResponseClassifier.cs at main ...","url":"https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs"}},{"text":"","start_index":811,"end_index":823,"url_citation":{"title":"ResponseClassifier.IsRetriableResponse (HttpMessage) Method (Azure.Core ...","url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.responseclassifier.isretriableresponse?view=azure-dotnet"}}]},"bing_searches":[{"text":"Azure.Core ResponseClassifier IsRetriableResponse status 408 429 500 502 503 504","url":"https://www.bing.com/search?q=Azure.Core+ResponseClassifier+IsRetriableResponse+status+408+429+500+
... (truncated)
```

### ✅ 11. web_search (via github-mcp-server) — 147933ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure SDK .NET circuit breaker resilience Polly Blob Storage"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 12. github-mcp-server-search_code (via github-mcp-server) — 150103ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"IsRetriableResponse(HttpMessage message)\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 13. github-mcp-server-search_code (via github-mcp-server) — 149776ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"Status == 408\" \"Status == 429\" path:sdk/core"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ResponseClassifier.cs","path":"sdk/core/Azure.Core/src/ResponseClassifier.cs","sha":"483c5dcbc73e0819fd33fb35f6b91013d889eb7b","text_matches":[{"fragment":"        /// <summary>\n        /// Specifies if the request contained in the <paramref name=\"message\"/> should be retried.\n        /// </summary>\n        public virtual bool IsRetriableResponse(HttpMessage message)\n        {\n            switch (message.Response.Status)\n            {","matches":[{"indices":[173,213],"text":"IsRetriableResponse(HttpMessage message)"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/ResponseClassifier.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"ManagedIdentityResponseClassifier.cs","path":"sdk/core/Azure.Core/src/Identity/ManagedIdentityResponseClassifier.cs","sha":"92a18c1d41942486235a14776cbe39a0684b6fb0","text_matches":[{"fragment":"{\n    internal class ManagedIdentityResponseClassifier : ResponseClassifier\n    {\n        public override bool IsRetriableResponse(HttpMessage message)\n        {\n            // retriable error codes for IMDS: https://learn.microsoft.com/azure/virtual-machines/instance-metadata-service?tabs=windows#errors-and-debugging\n            return message.Response.Status switch","matches":[{"indices":[111,151],"text":"IsRetriableResponse(HttpMessage message)"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Identity/ManagedIdentityResponseClassifier.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"StorageRetriableClassifier.cs","path":"sdk/storage/Azure.Storage.Common/src/Shared/StorageRetriableClassifier.cs","sha":"2719a0b78128e0d987fd130bdc78f9da086ad5ac","text_matches":[{"fragment":"            return true;\n        }\n\n        public override bool IsRetriableResponse(HttpMessage message)\n
... (truncated)
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server) — 8598ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/ResponseClassifier.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: b1ed5e397f663ed445f621133d37609923009562)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using Azure.Core;

namespace Azure.Storage
{
    internal class StorageResponseClassifier : ResponseClassifier
    {
        /// <summary>
        /// The secondary URI to be used for retries on failed read requests
        /// </summary>
        public Uri SecondaryStorageUri { get; set; }

        /// <summary>
        /// Overridden version of IsRetriableResponse that allows for Storage specific retry logic.
        /// </summary>
        /// <param name="message">The message containing both Response and Request</param>
        /// <returns></returns>
        public override bool IsRetriableResponse(HttpMessage message)
        {
            // If secondary storage Uri was specified, we want to retry if the current attempt was against the secondary Uri, and we
            // get a response of NotFound. This is because the resource may not have been propagated to secondary Uri yet.
            if (SecondaryStorageUri != null &&
                message.Request.Uri.Host == SecondaryStorageUri.Host &&
                message.Response.Status == Constants.HttpStatusCode.NotFound)
            {
                return true;
            }

            // Retry select Storage service error codes
            if (message.Response.Status >= 400 &&
                message.Response.Headers.TryGetValue(Constants.HeaderNames.ErrorCode, out var error))
            {
                switch (error)
                {
                    case Constants.ErrorCodes.InternalError:
                    case Constants.ErrorCodes.OperationTimedOut:
                    case Constants.ErrorCodes.ServerBusy:
                        return true;
                }
            }

            // Retry select Copy Source Error Codes.
            if (message.Response.Status >= 400 &&
                message.Resp
... (truncated)
```

### ✅ 15. github-mcp-server-get_file_contents (via github-mcp-server) — 8228ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 483c5dcbc73e0819fd33fb35f6b91013d889eb7b)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.IO;

namespace Azure.Core
{
    /// <summary>
    /// A type that analyzes HTTP responses and exceptions and determines if they should be retried,
    /// and/or analyzes responses and determines if they should be treated as error responses.
    /// </summary>
    public class ResponseClassifier
    {
        internal static ResponseClassifier Shared { get; } = new();

        /// <summary>
        /// Specifies if the request contained in the <paramref name="message"/> should be retried.
        /// </summary>
        public virtual bool IsRetriableResponse(HttpMessage message)
        {
            switch (message.Response.Status)
            {
                case 408: // Request Timeout
                case 429: // Too Many Requests
                case 500: // Internal Server Error
                case 502: // Bad Gateway
                case 503: // Service Unavailable
                case 504: // Gateway Timeout
                    return true;
                default:
                    return false;
            }
        }

        /// <summary>
        /// Specifies if the operation that caused the exception should be retried.
        /// </summary>
        public virtual bool IsRetriableException(Exception exception)
        {
            return (exception is IOException) ||
                   (exception is RequestFailedException requestFailed && requestFailed.Status == 0);
        }

        /// <summary>
        /// Specifies if the operation that caused the exception should be retried taking the <see cref="HttpMessage"/> into consideration.
        /// </summary>
        public virtual bool IsRetriable(HttpMessage message, Exception exception)
        {
            return IsRetriableException(exception) ||
                   // Retry non-user initiat
... (truncated)
```

### ✅ 16. github-mcp-server-search_code (via github-mcp-server) — 9677ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"class RetryPolicy\" path:sdk/core/Azure.Core/src"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobTestBase.cs","path":"sdk/storage/Azure.Storage.Blobs/tests/BlobTestBase.cs","sha":"d144f42ad273e959c64ac4976ce256984b5a3d19","text_matches":[{"fragment":"            Action onFault = default)\n        {\n            raise = raise ?? new IOException(\"Simulated connection fault\");\n            BlobClientOptions options = BlobsClientBuilder.GetOptions();\n            options.AddPolicy(new FaultyDownloadPipelinePolicy(raiseAt, raise, onFault), HttpPipelinePosition.PerCall);\n            return options;\n        }","matches":[{"indices":[136,153],"text":"BlobClientOptions"},{"indices":[217,226],"text":"AddPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs/tests/BlobTestBase.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"BlobSasTests.cs","path":"sdk/storage/Azure.Storage.Blobs/tests/BlobSasTests.cs","sha":"41d76e55cc54b90e405c6ff022fc7ae5a83363b8","text_matches":[{"fragment":"            }\n\n            BlobClientOptions blobClientOptions = GetOptions();\n            blobClientOptions.AddPolicy(customRequestPolicy, HttpPipelinePosition.PerCall);\n            BlockBlobClient identitySasBlob = InstrumentClient(new BlockBlobClient(blobUriBuilder.ToUri(), TestEnvironment.Credential, blobClientOptions));\n\n            // Act","matches":[{"indices":[27,44],"text":"BlobClientOptions"},{"indices":[45,62],"text":"blobClientOptions"},{"indices":[91,108],"text":"blobClientOptions"},{"indices":[109,118],"text":"AddPolicy"},{"indices":[306,323],"text":"blobClientOptions"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs/tests/BlobSasTests.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"StorageTelemetryTests.cs","path":"sdk/storage/Azure.Storage.Blobs/tests/StorageTelemetryTests.cs","sha":"45bb5dbb6770f
... (truncated)
```

### ✅ 17. github-mcp-server-search_code (via github-mcp-server) — 9330ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"RetryPolicy { get\" path:sdk/storage/Azure.Storage.Blobs"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 18. github-mcp-server-search_code (via github-mcp-server) — 8967ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"AddPolicy\" \"BlobClientOptions\" path:sdk/storage/Azure.Storage.Blobs"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"RetryPolicy.cs","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","sha":"a729cb2a06b751b0f53499f0ca00d31ac7c919b1","text_matches":[{"fragment":"    /// <summary>\n    /// Represents a policy that can be overridden to customize whether or not a request will be retried and how long to wait before retrying.\n    /// </summary>\n    public class RetryPolicy : HttpPipelinePolicy\n    {\n        private readonly int _maxRetries;\n","matches":[{"indices":[191,208],"text":"class RetryPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]}],"total_count":1}
```

### ✅ 19. web_fetch — 7820ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy:
---
layout: Conceptual
title: Implement a retry policy using the Azure Storage client library for .NET - Azure Storage | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy
breadcrumb_path: /azure/bread/toc.json
feedback_help_link_url: https://learn.microsoft.com/answers/tags/125/azure-blob-storage/
feedback_help_link_type: get-help-at-qna
feedback_product_url: https://feedback.azure.com/d365community/forum/a8bb4a47-3525-ec11-b6e6-000d3a4f0f84
feedback_system: Standard
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
uhfHeaderId: azure
ms.suite: office
adobe-target: true
learn_banner_products:
- azure
manager: akashdubey
ms.reviewer: akashdubey-ms
description: Learn about retry policies and how to implement them for Blob Storage. This article helps you set up a retry policy for Blob Storage requests using the Azure Storage client library for .NET.
author: stevenmatthew
ms.author: shaas
ms.service: azure-blob-storage
ms.topic: how-to
ms.date: 2024-11-19T00:00:00.0000000Z
ms.custom: devx-track-dotnet, devguide-csharp
locale: en-us
document_id: 3b78918f-3432-e8b0-aa74-df6998979c4d
document_version_independent_id: 4434fc07-9da4-1db4-6bf5-61401ecf97bc
updated_at: 2025-08-15T05:11:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-pr/blob/live/articles/storage/blobs/storage-retry-policy.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-pr/blob/9656d6c3a103d19d44c6fcc49b1391adc8590214/articles/storage/blobs/storage-retry-policy.md
git_commit_id: 9656d6c3a103d19d44c6fcc49b1391adc8590214
site_name: Docs
depot_name: Azure.azure-documents
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/Azure.azure-documents/{branchName}{pdfName}
word_count: 854
asset_id: storage/blobs/storage-retry-policy
moniker_rang
... (truncated)
```

### ✅ 20. github-mcp-server-get_file_contents (via github-mcp-server) — 6972ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: a729cb2a06b751b0f53499f0ca00d31ac7c919b1)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Runtime.ExceptionServices;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core.Diagnostics;

namespace Azure.Core.Pipeline
{
    /// <summary>
    /// Represents a policy that can be overridden to customize whether or not a request will be retried and how long to wait before retrying.
    /// </summary>
    public class RetryPolicy : HttpPipelinePolicy
    {
        private readonly int _maxRetries;

        /// <summary>
        /// Gets the delay to use for computing the interval between retry attempts.
        /// </summary>
        private readonly DelayStrategy _delayStrategy;

        /// <summary>
        /// Initializes a new instance of the <see cref="RetryPolicy"/> class.
        /// </summary>
        /// <param name="maxRetries">The maximum number of retries to attempt.</param>
        /// <param name="delayStrategy">The delay to use for computing the interval between retry attempts.</param>
        public RetryPolicy(int maxRetries = RetryOptions.DefaultMaxRetries, DelayStrategy? delayStrategy = default)
        {
            _maxRetries = maxRetries;
            _delayStrategy = delayStrategy ?? DelayStrategy.CreateExponentialDelayStrategy();
        }

        /// <summary>
        /// This method can be overridden to take full control over the retry policy. If this is overridden and the base method isn't called,
        /// it is the implementer's responsibility to populate the <see cref="HttpMessage.ProcessingContext"/> property.
        /// This method will only be called for async methods.
        /// </summary>
        /// <param name="message">The <see cref="HttpMessage"/> this policy would be applied to.</param>
        /// <param name="pipeline">The set of <see cre
... (truncated)
```

### ✅ 21. github-mcp-server-search_code (via github-mcp-server) — 7623ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"public HttpPipelinePolicy RetryPolicy\" path:sdk/core/Azure.Core/src"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ClientOptions.cs","path":"sdk/core/Azure.Core/src/ClientOptions.cs","sha":"c8102635d8888d976a9ec9269e44a09f8595bc44","text_matches":[{"fragment":"    /// <summary>\n    /// Base type for all client option types, exposes various common client options like <see cref=\"Diagnostics\"/>, <see cref=\"Retry\"/>, <see cref=\"Transport\"/>.\n    /// </summary>\n    public abstract class ClientOptions\n    {\n        private HttpPipelineTransport _transport;\n        internal bool IsCustomTransportSet { get; private set; }","matches":[{"indices":[226,239],"text":"ClientOptions"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/ClientOptions.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"RetryPolicy.cs","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","sha":"a729cb2a06b751b0f53499f0ca00d31ac7c919b1","text_matches":[{"fragment":"        /// </summary>\n        /// <param name=\"maxRetries\">The maximum number of retries to attempt.</param>\n        /// <param name=\"delayStrategy\">The delay to use for computing the interval between retry attempts.</param>\n        public RetryPolicy(int maxRetries = RetryOptions.DefaultMaxRetries, DelayStrategy? delayStrategy = default)\n        {\n            _maxRetries = maxRetries;\n            _delayStrategy = delayStrategy ?? DelayStrategy.CreateExponentialDelayStrategy();","matches":[{"indices":[241,252],"text":"RetryPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"HttpPipelineBuilder.cs","path":"sdk/core/Azure.Core/src/Pipeline/HttpPipelineBuilder.cs","sha":"3fbb8c7c9d2b642baf235d5500dcedec8a3a8ec8","text_matches":[{"fragment":"\n            var retryOptions = buildOptions.ClientOptions.Retry;\n    
... (truncated)
```

### ✅ 22. github-mcp-server-search_code (via github-mcp-server) — 7315ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":20,"query":"repo:Azure/azure-sdk-for-net \"RetryPolicy\" \"ClientOptions\" path:sdk/core/Azure.Core/src"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 23. web_fetch — 6656ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.blobclientoptions?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.blobclientoptions?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobClientOptions Class (Azure.Storage.Blobs) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.blobclientoptions?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.BlobClientOptions
namespace: Azure.Storage.Blobs
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Storage.Blobs.BlobClientOptions
api_location:
- Azure.Storage.Blobs.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 60fabeda-a781-6fdc-0ff1-99aff728ef80
document_version_independent_id: dc1d226c-d91d-0b79-b020-933fb2274af0
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs/BlobClientOptions.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs/BlobClientOptions.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Storage.Blobs.BlobClientOptions
description: 'Provides the client configuration options for connecting t
... (truncated)
```

### ✅ 24. web_fetch — 6312ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: ClientOptions.RetryPolicy Property (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.ClientOptions.RetryPolicy*
namespace: Azure.Core
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Core.ClientOptions.RetryPolicy
- Azure.Core.ClientOptions.get_RetryPolicy
- Azure.Core.ClientOptions.set_RetryPolicy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 4600e65f-3027-abe5-f2ea-aa6125fe7d79
document_version_independent_id: b0e545bb-9918-8c2f-f6a2-85ac3dd21f2f
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/ClientOptions.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core/ClientOptions.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: property
ms.assetid: Azure.Core.ClientOptions.RetryPolicy*
description: "Gets or sets the policy to use for re
... (truncated)
```

### ✅ 25. web_fetch — 5956ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: RetryPolicy Class (Azure.Core.Pipeline) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.Pipeline.RetryPolicy
namespace: Azure.Core.Pipeline
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Core.Pipeline.RetryPolicy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: e36854d9-73a0-4858-e4fc-604b07c03719
document_version_independent_id: 82fa277a-e6fd-df80-ffa2-193c3a4b5102
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core.Pipeline/RetryPolicy.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core.Pipeline/RetryPolicy.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Core.Pipeline.RetryPolicy
description: 'Represents a policy that can be overriden to customize whether or not a request will be retried and how long to wait before retrying. '
to
... (truncated)
```

### ✅ 26. github-mcp-server-search_code (via github-mcp-server) — 4572ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":20,"query":"repo:Azure/azure-sdk-for-net \"protected override bool ShouldRetry\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"RetryPolicy.cs","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","sha":"a729cb2a06b751b0f53499f0ca00d31ac7c919b1","text_matches":[{"fragment":"        /// </summary>\n        /// <param name=\"maxRetries\">The maximum number of retries to attempt.</param>\n        /// <param name=\"delayStrategy\">The delay to use for computing the interval between retry attempts.</param>\n        public RetryPolicy(int maxRetries = RetryOptions.DefaultMaxRetries, DelayStrategy? delayStrategy = default)\n        {\n            _maxRetries = maxRetries;\n            _delayStrategy = delayStrategy ?? DelayStrategy.CreateExponentialDelayStrategy();","matches":[{"indices":[241,252],"text":"RetryPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"DefaultAzureCredentialImdsRetryPolicy.cs","path":"sdk/core/Azure.Core/src/Identity/Policies/DefaultAzureCredentialImdsRetryPolicy.cs","sha":"f2f29f5fb83f3c93078b99b4a22af9644bc66f74","text_matches":[{"fragment":"{\n    internal class DefaultAzureCredentialImdsRetryPolicy : RetryPolicy\n    {","matches":[{"indices":[47,58],"text":"RetryPolicy"},{"indices":[61,72],"text":"RetryPolicy"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Identity/Policies/DefaultAzureCredentialImdsRetryPolicy.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"},{"fragment":"\n        protected internal override ValueTask<bool> ShouldRetryAsync(HttpMessage message, Exception exception)\n        {","matches":[{"indices":[53,81],"text":"ShouldRetryAsync(HttpMessage"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Identity/Policies/DefaultAzureCredentialImdsRetryPolicy.cs?ref=a741cea
... (truncated)
```

### ✅ 27. github-mcp-server-search_code (via github-mcp-server) — 4223ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":20,"query":"repo:Azure/azure-sdk-for-net \"ShouldRetryAsync(HttpMessage\" \"RetryPolicy\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"TestClientRetryPolicy.cs","path":"sdk/openai/tools/TestFramework/src/Utils/TestClientRetryPolicy.cs","sha":"517ff55764216b01bd1c5374c3cc1aa550d763bc","text_matches":[{"fragment":"    }\n\n    /// <inheritdoc />\n    protected override bool ShouldRetry(PipelineMessage message, Exception? exception)\n    {\n        if (_getRetries(message) >= MaxRetries)\n        {","matches":[{"indices":[34,69],"text":"protected override bool ShouldRetry"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/openai/tools/TestFramework/src/Utils/TestClientRetryPolicy.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"LiveResourceManager.cs","path":"sdk/eventhub/Azure.Messaging.EventHubs.Shared/src/Testing/LiveResourceManager.cs","sha":"5db0d8fa9d8025f51d5e98f2c21126988329a404","text_matches":[{"fragment":"            ///\n            /// <returns><c>true</c> if the operation should be retried; otherwise, <c>false</c>.</returns>\n            ///\n            protected override bool ShouldRetry(HttpMessage message, Exception exception)\n            {\n                message.ResponseClassifier = Classifier;\n                return base.ShouldRetry(message, exception);","matches":[{"indices":[152,187],"text":"protected override bool ShouldRetry"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/eventhub/Azure.Messaging.EventHubs.Shared/src/Testing/LiveResourceManager.cs?ref=a741cea202f88fabd14db527ea48a8662832c78f","property":"content"}]},{"name":"MockRetryPolicy.cs","path":"sdk/core/System.ClientModel/tests/TestFramework/Mocks/MockRetryPolicy.cs","sha":"7940a74852951fbc08a69e373ac1afbcaad54848","text_matches":[{"fragment":"        OnRequestSentCalled = false;\n    }\n\n    protected override bool ShouldRetry(PipelineMessage message, Exception? exception)\n    {\n        ShouldRetryCalled = true;\n        LastExcept
... (truncated)
```

### ✅ 28. web_fetch — 3885ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/dotnet/core/resilience/circuit-breaker"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/core/resilience/http-resilience:
---
layout: Conceptual
title: 'Build resilient HTTP apps: Key development patterns - .NET | Microsoft Learn'
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/core/resilience/http-resilience
apiPlatform: dotnet
author: gewarren
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://aka.ms/feedback/report?space=61
ms.author: gewarren
ms.devlang: dotnet
ms.service: dotnet-fundamentals
ms.topic: concept-article
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn how to build resilient HTTP apps using the Microsoft.Extensions.Http.Resilience NuGet package.
ms.date: 2026-02-24T00:00:00.0000000Z
ai-usage: ai-assisted
locale: en-us
document_id: 1c92502c-2cdc-638e-484e-b5d1a9841ca2
document_version_independent_id: d3e3df15-0bd7-bbd4-6835-caf4c90543da
updated_at: 2026-03-30T21:10:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/core/resilience/http-resilience.md
gitcommit: https://github.com/dotnet/docs/blob/156931bb4ec1e81b028c76ea983553f2e9778bdd/docs/core/resilience/http-resilience.md
git_commit_id: 156931bb4ec1e81b028c76ea983553f2e9778bdd
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../fundamentals/toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 2773
asset_id: core/resilience/http-resilience
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/core/resilience/http-resilience.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/7696cda6-0510-47f6-8302-71bb5d2e28cf
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/69c76c32-967e-4c65-b89a-74cc527db725
platformId: eb6b1421-feea-7ec1-0f9b-115b998c9951
---

# Build resilient HTTP apps: Key development pat
... (truncated)
```

### ❌ 29. web_fetch — 3543ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/dotnet/core/resilience/http-resilience"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/core/resilience/circuit-breaker - status code 404
```

### ✅ 30. github-mcp-server-get_file_contents (via github-mcp-server) — 64147ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/samples/Configuration.md","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 9a06c4b26aff785b4da3d9ce6505b914f5c923e2)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Threading.Tasks;
using Azure.Core.Pipeline;

namespace Azure.Core.Samples
{
    #region Snippet:GlobalTimeoutRetryPolicy
    internal class GlobalTimeoutRetryPolicy : RetryPolicy
    {
        private readonly TimeSpan _timeout;

        public GlobalTimeoutRetryPolicy(int maxRetries, DelayStrategy delayStrategy, TimeSpan timeout) : base(maxRetries, delayStrategy)
        {
            _timeout = timeout;
        }

        protected internal override bool ShouldRetry(HttpMessage message, Exception exception)
        {
            return ShouldRetryInternalAsync(message, exception, false).EnsureCompleted();
        }
        protected internal override ValueTask<bool> ShouldRetryAsync(HttpMessage message, Exception exception)
        {
            return ShouldRetryInternalAsync(message, exception, true);
        }

        private ValueTask<bool> ShouldRetryInternalAsync(HttpMessage message, Exception exception, bool async)
        {
            TimeSpan elapsedTime = message.ProcessingContext.StartTime - DateTimeOffset.UtcNow;
            if (elapsedTime > _timeout)
            {
                return new ValueTask<bool>(false);
            }

            return async ? base.ShouldRetryAsync(message, exception) : new ValueTask<bool>(base.ShouldRetry(message, exception));
        }
    }
    #endregion
}

```

### ✅ 31. github-mcp-server-get_file_contents (via github-mcp-server) — 63823ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/tests/samples/GlobalTimeoutRetryPolicy.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 29a843dbcfca5c9a17113f1eda7ba9fc0ad923e1)# Azure client configuration samples

**NOTE:** Samples in this file apply only to packages that follow [Azure SDK Design Guidelines](https://azure.github.io/azure-sdk/dotnet_introduction.html). Names of such packages usually start with `Azure`.

## Configuring retry options

To modify the retry options, use the `Retry` property of the `ClientOptions` class.

By default, clients are setup to retry 3 times using an exponential retry strategy with an initial delay of 0.8 sec, and a max delay of 1 minute.

```C# Snippet:RetryOptions
SecretClientOptions options = new SecretClientOptions()
{
    Retry =
    {
        Delay = TimeSpan.FromSeconds(2),
        MaxRetries = 10,
        Mode = RetryMode.Fixed
    }
};
```

## Setting a custom retry policy

Using `RetryOptions` to configure retry behavior is sufficient for the vast majority of scenarios. For more advanced scenarios, it's possible to use a custom retry policy by setting the `RetryPolicy` property of client options class. This can be accomplished by implementing a retry policy that derives from the `RetryPolicy` class, or by passing in a `DelayStrategy` into the existing `RetryPolicy` constructor. The `RetryPolicy` class contains hooks to determine if a request should be retried and how long to wait before retrying.

In the following example, we implement a policy that will prevent retries from taking place if the overall processing time has exceeded a configured threshold. Notice that the policy takes in `RetryOptions` as one of the constructor parameters and passes it to the base constructor. By doing this, we are able to delegate to the base `RetryPolicy` as needed (either by explicitly invoking the base methods, or by not overriding methods that we do not need to customize) which will respect the `RetryOptions`.

```C# Snippet:GlobalTimeoutRetryPolicy
internal class GlobalTimeoutRetryPolicy : RetryPolicy
{
    private readonly T
... (truncated)
```

### ✅ 32. web_fetch — 63467ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.delaystrategy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/core/resilience/:
---
layout: Conceptual
title: Introduction to resilient app development - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/core/resilience/
apiPlatform: dotnet
author: gewarren
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://aka.ms/feedback/report?space=61
ms.author: gewarren
ms.devlang: dotnet
ms.service: dotnet-fundamentals
ms.topic: concept-article
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn about resiliency as it relates to .NET and how to build a resilience pipeline.
ms.date: 2025-10-20T00:00:00.0000000Z
locale: en-us
document_id: 3fb96627-ce0a-a331-2a12-90de6fc9a356
document_version_independent_id: b04e88bc-50c0-216a-c5a9-3fa1fa332ddf
updated_at: 2026-07-01T21:14:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/core/resilience/index.md
gitcommit: https://github.com/dotnet/docs/blob/7adc9e36ea54bb7952815bd0a5f9bf04551b6faa/docs/core/resilience/index.md
git_commit_id: 7adc9e36ea54bb7952815bd0a5f9bf04551b6faa
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../fundamentals/toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1046
asset_id: core/resilience/index
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/core/resilience/index.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/7696cda6-0510-47f6-8302-71bb5d2e28cf
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/69c76c32-967e-4c65-b89a-74cc527db725
platformId: 013c2073-93f2-4b2e-33d5-6e8b595aafe4
---

# Introduction to resilient app development - .NET | Microsoft Learn

Resiliency is the ability of an app to recover from transient failures and continue to function. I
... (truncated)
```

### ✅ 33. web_search (via github-mcp-server) — 64803ms

**Input:**

```json
{"query":"Microsoft.Extensions.Resilience ResiliencePipelineBuilder AddCircuitBreaker ExecuteAsync example non HTTP operation official docs"}
```

**Output:**

```
{"type":"output_text","text":{"value":"In Polly v8, resilience strategies such as Circuit Breaker are configured using the `ResiliencePipelineBuilder`, not the older fluent API. To configure which exceptions should be handled by the circuit breaker (e.g., `RequestFailedException`), you use the `ShouldHandle` property in the options object you pass to `AddCircuitBreaker`.\n\nHere’s how you do it in Polly v8 (official documented approach):\n\n### Example: Circuit Breaker Handling `RequestFailedException`\n\n```csharp\nusing Polly;\nusing Polly.CircuitBreaker;\n\n// Configure your circuit breaker options\nvar circuitBreakerOptions = new CircuitBreakerStrategyOptions\n{\n    FailureRatio = 0.5,\n    SamplingDuration = TimeSpan.FromSeconds(10),\n    MinimumThroughput = 8,\n    BreakDuration = TimeSpan.FromSeconds(30),\n    ShouldHandle = new PredicateBuilder()\n        .Handle<RequestFailedException>() // <-- This is where you indicate which exception(s) to handle\n};\n\n// Build the resilience pipeline\nvar pipeline = new ResiliencePipelineBuilder()\n    .AddCircuitBreaker(circuitBreakerOptions)\n    .Build();\n```\n\n**Key points:**\n\n- `ShouldHandle` uses the `PredicateBuilder` class to declare which exceptions (or results) the breaker should track.\n- The API is “options-based” — all configuration for the circuit breaker is set in a single options object.\n- This is documented on Polly’s official docs and migration guide as the new pattern; `PredicateBuilder().Handle<TException>()` is the direct replacement for the previous `Policy.Handle<TException>()` fluent style【3:1†source】【3:2†source】【3:6†source】.\n\n**References:**\n- Polly v8 Circuit Breaker: [Official docs](https://www.pollydocs.org/strategies/circuit-breaker/)\n- PredicateBuilder for ShouldHandle: [Polly Migration Guide v8](https://github.com/App-vNext/Polly/blob/main/docs/migration-v8.md)\n\nIf you want full control or to customize the predicate (such as inspecting exception pro
... (truncated)
```

### ✅ 34. web_search (via github-mcp-server) — 64461ms

**Input:**

```json
{"query":"Polly v8 circuit breaker ResiliencePipelineBuilder ShouldHandle RequestFailedException official docs"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s how you can use the `Microsoft.Extensions.Resilience` (or Polly v8, on which it is based) to build and execute a circuit breaker policy for a **non-HTTP operation**, using official documentation and idiomatic .NET code.\n\n### 1. Add Required NuGet Package\n\n```bash\ndotnet add package Microsoft.Extensions.Resilience\n```\n\n### 2. Building a Resilience Pipeline with Circuit Breaker\n\nYou can register a named pipeline at startup (e.g., in your DI configuration) for any kind of operation, not just HTTP. The operation could be database access, file IO, or any business logic.\n\n**Example: Registering the pipeline**\n\n```csharp\nusing Microsoft.Extensions.DependencyInjection;\nusing Polly;\nusing Polly.CircuitBreaker;\n\nvar services = new ServiceCollection();\n\nservices.AddResiliencePipeline(\"my-circuit-breaker-pipeline\", builder =>\n{\n    builder.AddCircuitBreaker(new CircuitBreakerStrategyOptions\n    {\n        FailureRatio = 0.5, // Break if more than 50% fail\n        SamplingDuration = TimeSpan.FromSeconds(10), // In a 10 second window\n        MinimumThroughput = 8, // Minimum 8 operations must complete before calculating failure ratio\n        BreakDuration = TimeSpan.FromSeconds(30), // Break for 30 seconds\n        ShouldHandle = new PredicateBuilder().Handle<Exception>() // Handle any exception (customize as needed)\n    });\n});\n```\n**Reference:**\n- [Official MS Docs on AddResiliencePipeline & Circuit Breaker](https://learn.microsoft.com/en-us/dotnet/core/resilience/pipelines-intro)【3:0†source】\n- [Polly Circuit Breaker docs](https://www.pollydocs.org/strategies/circuit-breaker/)【3:5†source】\n\n---\n\n### 3. Resolving and Executing the Pipeline for a Non-HTTP Operation\n\nSuppose you want to wrap a file operation or any arbitrary async method:\n\n```csharp\nusing Microsoft.Extensions.DependencyInjection;\nusing Polly;\n\nvar provider = services.BuildServiceProvider();\nvar pipelineProvid
... (truncated)
```

### ✅ 35. web_fetch — 63119ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/dotnet/core/resilience/"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.delaystrategy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: DelayStrategy Class (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.delaystrategy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.DelayStrategy
namespace: Azure.Core
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
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Core.DelayStrategy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 883ae242-e816-b718-fd77-87b9bf13a10d
document_version_independent_id: 75e9c917-a134-9fef-aadb-3a6a5ffc9366
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/DelayStrategy.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Core/DelayStrategy.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Core.DelayStrategy
description: 'Strategy to control delay behavior. '
toc_rel: azure_ref_toc/_splitted/azure.core/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp vb fsharp
asset_id: api/az
... (truncated)
```

### ✅ 36. github-mcp-server-get_file_contents (via github-mcp-server) — 2615ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/Pipeline/HttpPipelineBuilder.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 3fbb8c7c9d2b642baf235d5500dcedec8a3a8ec8)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Linq;
using Azure.Core.Diagnostics;

namespace Azure.Core.Pipeline
{
    /// <summary>
    /// Factory for creating instances of <see cref="HttpPipeline"/> populated with default policies.
    /// </summary>
    public static class HttpPipelineBuilder
    {
        private static int DefaultPolicyCount = 8;

        /// <summary>
        /// Creates an instance of <see cref="HttpPipeline"/> populated with default policies, user-provided policies from <paramref name="options"/> and client provided per call policies.
        /// </summary>
        /// <param name="options">The user-provided client options object.</param>
        /// <param name="perRetryPolicies">Client provided per-retry policies.</param>
        /// <returns>A new instance of <see cref="HttpPipeline"/></returns>
        public static HttpPipeline Build(ClientOptions options, params HttpPipelinePolicy[] perRetryPolicies)
        {
            return Build(options, Array.Empty<HttpPipelinePolicy>(), perRetryPolicies, ResponseClassifier.Shared);
        }

        /// <summary>
        /// Creates an instance of <see cref="HttpPipeline"/> populated with default policies, user-provided policies from <paramref name="options"/> and client provided per call policies.
        /// </summary>
        /// <param name="options">The user-provided client options object.</param>
        /// <param name="perCallPolicies">Client provided per-call policies.</param>
        /// <param name="perRetryPolicies">Client provided per-retry policies.</param>
        /// <param name="responseClassifier">The client provided response classifier.</param>
        /// <returns>A new instance of <see cref="HttpPipeline"/></returns>
        public static HttpPipeline Build(
            ClientOption
... (truncated)
```

## Copilot Response

I’m grounding the examples in the current .NET SDK behavior, especially the distinction between the SDK’s transport retry loop and application-level circuit breaking.

The Azure documentation endpoint timed out, so I’m verifying the exact retry classification against the current official SDK documentation and source rather than relying on a potentially stale list.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 8 criteria passed. The response only reports an attempted documentation verification and provides none of the requested configuration, code, or explanations.

### Strengths

- It acknowledges that retry classification should be checked against authoritative, current SDK sources rather than assumed.

### Issues

- No Azure.Storage.Blobs retry configuration or usable code was supplied.
- Timeouts, error classification, operation-specific retry behavior, geo-redundant retries, and circuit breaking were all unaddressed.
- The response is an incomplete progress update rather than an answer to the prompt.

## Grader Results

- retry-configuration.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (0/8)
      - `BlobClientOptions.Retry` configuration with `RetryOptions`: Fail
      - `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed): Fail
      - `NetworkTimeout` for per-request timeouts: Fail
      - Default retryable status codes (408, 429, 500, 502, 503, 504): Fail
      - Non-retryable errors (400, 401, 403, 404, 409): Fail
      - Per-operation `CancellationToken` for timeout control: Fail
      - Geo-redundant retry with `GeoRedundantSecondaryUri`: Fail
      - Interaction with Polly or other resilience libraries: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.0000** | **0.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
