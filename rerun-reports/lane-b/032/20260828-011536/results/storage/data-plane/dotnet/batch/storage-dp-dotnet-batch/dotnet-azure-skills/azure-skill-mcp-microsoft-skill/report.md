# Evaluation Report: storage-dp-dotnet-batch

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 549.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-batch` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 549.8s |
| Timestamp | 2026-08-27T17:15:36Z |
| Files Generated | 0 |
| Event Count | 3923 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 515.0s |
| Review | 33.6s |
| **Total** | **549.8s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=262014 out=4856 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3943ms

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

### ✅ 2. glob — 188535ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 189213ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
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

### ✅ 4. view — 188872ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
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

### ❌ 5. azure-get_azure_bestpractices (via azure) — 189215ms

**Input:**

```json
{"intent":"Generate modern .NET sample code using Azure.Storage.Blobs.Batch BlobBatchClient with secure authentication and partial-failure handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 189219ms

**Input:**

```json
{"intent":"Find official current Microsoft documentation for .NET Azure.Storage.Blobs.Batch BlobBatchClient, SubmitBatchAsync partial failure handling, convenience methods, and Blob Batch REST limits including maximum subrequests and request size.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. web_search (via github-mcp-server) — 183402ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure.Storage.Blobs.Batch .NET BlobBatchClient SubmitBatchAsync throwOnAnyFailure partial failures BlobBatchOperation Response 256 4 MiB"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 8. web_search (via github-mcp-server) — 183002ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure.Storage.Blobs.Specialized.BlobBatchClient DeleteBlobsAsync SetBlobsAccessTierAsync SubmitBatchAsync"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The Azure Blob Batch REST API allows you to combine multiple API calls into a single HTTP request. Here are the specific limits:\n\n- **Maximum Subrequests:** You can include up to 256 subrequests (for example, Delete Blob or Set Blob Tier operations) in a single batch request.\n- **Maximum Request Body Size:** The entire batch request body (including all subrequests) cannot exceed 4 MiB in size.\n- **Scope:** All subrequests must pertain to resources in the same storage account. If you use the REST API with the `restype=container` parameter, all subrequests must also be scoped to the same container.\n\nThese limits are important for optimizing operations that affect many blobs at once, but if you exceed either limit, your request will fail.\n\nFor further details, you can review the official documentation on [Blob Batch (REST API) - Azure Storage](https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch) or see this summary in the .NET client doc, which matches the REST API constraints:\n- \"Only supports up to 256 subrequests in a single batch.\"\n- \"The size of the body for a batch request cannot exceed 4MB.\"\n- \"Each subrequest must be for a resource within the same storage account.\"【3:5†source】.","annotations":[{"text":"","start_index":1206,"end_index":1218,"url_citation":{"title":"Azure Storage Blobs Batch client library for .NET - Azure for .NET ...","url":"https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet"}}]},"bing_searches":[{"text":"blob batch REST API maximum subrequests 4 MiB same storage account site:learn.microsoft.com","url":"https://www.bing.com/search?q=blob+batch+REST+API+maximum+subrequests+4+MiB+same+storage+account+site%3Alearn.microsoft.com"}],"annotations":null}
```

### ✅ 9. web_search (via github-mcp-server) — 182600ms

**Input:**

```json
{"query":"site:learn.microsoft.com/rest/api/storageservices/blob-batch maximum 256 subrequests 4 MiB same storage account"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s a detailed answer addressing your queries about `Azure.Storage.Blobs.Batch` in .NET and the `BlobBatchClient.SubmitBatchAsync` behavior for partial failures, response limits, and batch size considerations:\n\n### 1. **Partial Failures and `throwOnAnyFailure`**\n\nThe `BlobBatchClient.SubmitBatchAsync` method allows batching multiple sub-operations (like Deletes or Set Tier) into a single request. The relevant parameter here is `throwOnAnyFailure`:\n\n- If `throwOnAnyFailure` is **false** (default): The method completes and returns a response even if some sub-operations failed; you must inspect individual operation results to detect failures.\n- If **true**: The method throws an exception (wrapped in `AggregateException`) if any of the sub-operations fail, so you’ll need to handle this at your callsite【3:0†source】【3:2†source】.\n\n**Summary:**  \n- Set `throwOnAnyFailure` to `true` for “all or nothing” behavior (exception on partial failure).\n- Set to `false` for granular control: check success/failure per operation in the response.\n\n### 2. **Partial Failure Handling in Response**\n\nWhen `throwOnAnyFailure` is **false**, the returned `Response` object encapsulates the results for each sub-operation. You’ll need to parse/check individual status codes to determine which items failed.\n\n### 3. **Batch Operation Limits**\n\n- **Max requests per batch:** 256 sub-operations per batch are allowed.\n- **Max batch payload size:** The total request body (including all batched sub-operations) must not exceed 4 MiB (4 mebibytes)【3:5†source】.\n\nSee REST API reference:\n> \"A batch request can contain up to 256 sub-requests and the maximum payload size for the entire batch request is 4 MiB.\"【3:5†source】\n\n### 4. **Relevant Documentation/References**\n\n- [BlobBatchClient.SubmitBatchAsync Method - Microsoft Learn](https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobba
... (truncated)
```

### ❌ 10. azure-get_azure_bestpractices (via azure) — 183403ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate modern .NET sample code using Azure.Storage.Blobs.Batch BlobBatchClient with secure authentication and partial-failure handling.","parameters":{}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 11. azure-documentation (via azure) — 182154ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find BlobBatchClient .NET API and Blob Batch REST limits and partial failure behavior.","parameters":{"query":"Azure.Storage.Blobs.Batch BlobBatchClient SubmitBatchAsync partial failure 256 4 MiB"}}
```

**Output:**

```
{"results":[{"title":"Blob Batch","content":"# Blob Batch\n## Remarks\nOne of the main benefits of using a batch request is the reduction in the number of connections a client has to open. Note the following restrictions:\n1. Supported subrequests in the batch are \u0060Set Blob Tier\u0060 (for block blobs) and \u0060Delete Blob\u0060.\n2. Only supports up to 256 subrequests in a single batch. The size of the body for a batch request can\u0027t exceed 4 MB.\n3. An empty batch request fails with code 400 (Bad Request).\n4. There are no guarantees on the order of execution of the batch subrequests.\n5. Batch subrequest execution isn\u0027t atomic. Each subrequest runs independently.\n6. Each subrequest must be for a resource within the same storage account. A single batch request doesn\u0027t support running requests from different storage accounts.\n7. A nested request body isn\u0027t supported.\n8. If the server fails to parse the request body, the entire batch fails, and no request will be run.\n9. Note that [Account SAS](https://learn.microsoft.com/rest/api/storageservices/create-account-sas) is the only shared access signature type supported by \u0060Blob Batch\u0060, when the batch isn\u0027t using \u0060restype=container\u0060.\n### Scope all subrequests to a specific container\nBeginning with REST version 2020-04-08, the \u0060Blob Batch\u0060 API supports scoping subrequests to a specified container. When the request URI includes the container name and the \u0060restype=container\u0060 parameter, then each subrequest must apply to the same container. If the container name specified for a subrequest doesn\u0027t match the container name provided in URI, the service returns error code 400 (Bad Request).\nAll authorization mechanisms supported for a container are valid for a \u0060Blob Batch\u0060 operation that is scoped to the container. Each subrequest sends an authorization header to the service.","contentUrl":"https://learn.microsoft.com/rest/api/storageser
... (truncated)
```

### ✅ 12. rg — 2711ms

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787851179300-copilot-tool-output-2d928e7ed59247059ee8dbc917959d75.txt","pattern":"BlobBatchOperation|DeleteBlob|SetBlobAccessTier|SubmitBatchAsync|throwOnAnyFailure|AggregateException|256|4 MB|4 MiB"}
```

**Output:**

```
Output too large to read at once (46.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787851369560-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787851179300-copilot-tool-output-2d928e7ed59247059ee8dbc917959d75.txt:{"results":[{"title":"Blob Batch","content":"# Blob Batch\n## Remarks\nOne of the main benefits of using a batch request is the reduction in the number of connections a client has to open. Note the following restrictions:\n1. Supported subrequests in the batch are \u0060Set Blob Tier\u0060 (for block blobs) and \u0060Delete Blob\u0060.\n2. Only supports up to 256 subrequests in a single batc
```

### ✅ 13. web_fetch — 3814ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobBatch Class (Azure.Storage.Blobs.Specialized) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Specialized.BlobBatch
namespace: Azure.Storage.Blobs.Specialized
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
- Azure.Storage.Blobs.Specialized.BlobBatch
api_location:
- Azure.Storage.Blobs.Batch.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 1e37300c-6c42-0650-30c5-817d17f5a3b3
document_version_independent_id: 5f464ae6-4ae2-bfab-0335-eee7d7b81e77
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Specialized/BlobBatch.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Specialized/BlobBatch.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Storage.Blobs.Specialized.BlobBatch
description: 'A BlobBa
... (truncated)
```

### ✅ 14. web_fetch — 3463ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobBatchClient.SubmitBatchAsync Method (Azure.Storage.Blobs.Specialized) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Specialized.BlobBatchClient.SubmitBatchAsync*
namespace: Azure.Storage.Blobs.Specialized
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
- Azure.Storage.Blobs.Specialized.BlobBatchClient.SubmitBatchAsync
api_location:
- Azure.Storage.Blobs.Batch.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: f5862cb9-c479-f23a-a953-baf546269862
document_version_independent_id: 36335e37-1a85-58c2-2856-e4e2730d8978
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Specialized/BlobBatchClient.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Specialized/BlobBatchClient.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: 
... (truncated)
```

### ✅ 15. web_fetch — 3126ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet:
---
layout: Conceptual
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: Azure Storage Blobs Batch client library for .NET - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet
config_moniker_range: azure-dotnet
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-net/issues
author: azure-sdk
ms.author: azsdkdocs
ms.topic: reference
ms.devlang: dotnet
ms.date: 2026-06-04T00:00:00.0000000Z
uhfHeaderId: Azure
breadcrumb_path: ../../../azure_breadcrumb/toc.json
apiPlatform: dotnet
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-authored
- azure-dotnet-svc-overview
keywords: Azure, dotnet, SDK, API, Azure.Storage.Blobs.Batch, storage
ms.service: storage
locale: en-us
document_id: efa03f33-bcd9-3d7c-3615-022a8a19f4b1
document_version_independent_id: 9ed2e40e-26ff-96b2-ae30-5dee5eacba41
updated_at: 2026-06-05T11:43:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/api/overview/azure/latest/storage.blobs.batch-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/54218fa5efb5f8acf2d764893f4ca7bf068c921a/api/overview/azure/latest/storage.blobs.batch-readme.md
git_commit_id: 54218fa5efb5f8acf2d764893f4ca7bf068c921a
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: conceptual
toc_rel: ../../azure_ref_toc/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 847
asset_id: api/overview/azure/storage.blobs.batch-readme
moniker_range_name: 0e50d6ab13f0a9e30600d6d5dbd12dc4
monikers:
- azure-dotnet
item_type: Content
source_path: api/over
... (truncated)
```

### ✅ 16. web_fetch — 4272ms

**Input:**

```json
{"max_length":14000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.deleteblobsasync?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobBatchClient.DeleteBlobsAsync Method (Azure.Storage.Blobs.Specialized) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.deleteblobsasync?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Specialized.BlobBatchClient.DeleteBlobsAsync*
namespace: Azure.Storage.Blobs.Specialized
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
- Azure.Storage.Blobs.Specialized.BlobBatchClient.DeleteBlobsAsync
api_location:
- Azure.Storage.Blobs.Batch.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 9ac35066-3931-5168-148e-5636f95727ff
document_version_independent_id: 1adc4b05-0ddf-1411-0962-7adb38c18a49
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Specialized/BlobBatchClient.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Specialized/BlobBatchClient.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: 
... (truncated)
```

### ✅ 17. web_fetch — 3900ms

**Input:**

```json
{"max_length":14000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.setblobaccesstier?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.setblobsaccesstierasync?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobBatchClient.SetBlobsAccessTierAsync Method (Azure.Storage.Blobs.Specialized) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.setblobsaccesstierasync?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Specialized.BlobBatchClient.SetBlobsAccessTierAsync*
namespace: Azure.Storage.Blobs.Specialized
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
- Azure.Storage.Blobs.Specialized.BlobBatchClient.SetBlobsAccessTierAsync
api_location:
- Azure.Storage.Blobs.Batch.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 9a0817cd-3691-3e1d-b1e9-a9403aa77e7b
document_version_independent_id: 92778d92-755e-fb19-d471-181de585eef5
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Specialized/BlobBatchClient.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Specialized/BlobBatchClient.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-
... (truncated)
```

### ✅ 18. web_fetch — 3530ms

**Input:**

```json
{"max_length":14000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchoperation?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobBatch.DeleteBlob Method (Azure.Storage.Blobs.Specialized) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Specialized.BlobBatch.DeleteBlob*
namespace: Azure.Storage.Blobs.Specialized
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
- Azure.Storage.Blobs.Specialized.BlobBatch.DeleteBlob
api_location:
- Azure.Storage.Blobs.Batch.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 92a1ce81-8222-ede3-2afb-63fc6ac19fee
document_version_independent_id: 9af6b596-9e70-d49a-1f6d-2b87bbc89128
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Specialized/BlobBatch.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Specialized/BlobBatch.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: method
ms.assetid: Azure.
... (truncated)
```

### ❌ 19. web_fetch — 3135ms

**Input:**

```json
{"max_length":16000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.deleteblobsasync?view=azure-dotnet"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchoperation?view=azure-dotnet - status code 404
```

### ✅ 20. web_fetch — 3131ms

**Input:**

```json
{"max_length":16000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.setblobsaccesstierasync?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.setblobaccesstier?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: BlobBatch.SetBlobAccessTier Method (Azure.Storage.Blobs.Specialized) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.setblobaccesstier?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Storage.Blobs.Specialized.BlobBatch.SetBlobAccessTier*
namespace: Azure.Storage.Blobs.Specialized
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
- Azure.Storage.Blobs.Specialized.BlobBatch.SetBlobAccessTier
api_location:
- Azure.Storage.Blobs.Batch.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 4a79e702-13e0-fe76-5878-82a9883f7b93
document_version_independent_id: 481ff8dc-966f-58f1-3a45-03d07382a64e
updated_at: 2026-07-23T11:31:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Storage.Blobs.Specialized/BlobBatch.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/3f60666cf8fd596774ed882982439513c6fcbdd7/xml/Azure.Storage.Blobs.Specialized/BlobBatch.xml
git_commit_id: 3f60666cf8fd596774ed882982439513c6fcbdd7
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
p
... (truncated)
```

## Copilot Response

I’m checking the current .NET batch API and service limits because 500 operations may exceed a single batch request.

The service cap is the key constraint; I’m verifying the exact API surface and failure semantics before giving code that chunks 500 deletes correctly.

You **cannot delete 500 blobs in one batch HTTP request**. Azure Blob Batch allows at most **256 subrequests**, so 500 deletes require two requests: **256 + 244**. Changing the tier of 200 blobs fits in one request.

Install the packages:

```powershell
dotnet add package Azure.Storage.Blobs.Batch
dotnet add package Azure.Identity
```

Use a manually constructed `BlobBatch` with `throwOnAnyFailure: false`. This preserves each subrequest’s `Response`, allowing successful and failed operations to be correlated with blob names.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

const int MaxOperationsPerBatch = 256;

public sealed record BatchItemResult(
    string BlobName,
    int Status,
    string? ErrorCode,
    string ReasonPhrase)
{
    public bool Succeeded => Status is >= 200 and < 300;
}

public static async Task<IReadOnlyList<BatchItemResult>> SubmitInBatchesAsync(
    BlobBatchClient batchClient,
    IReadOnlyList<string> blobNames,
    Func<BlobBatch, string, Response> addOperation,
    CancellationToken cancellationToken = default)
{
    var results = new List<BatchItemResult>(blobNames.Count);

    for (int offset = 0; offset < blobNames.Count; offset += MaxOperationsPerBatch)
    {
        string[] chunk = blobNames
            .Skip(offset)
            .Take(MaxOperationsPerBatch)
            .ToArray();

        using BlobBatch batch = batchClient.CreateBatch();

        // These are deferred responses. They become readable after SubmitBatchAsync.
        var operations = new List<(string BlobName, Response Response)>(chunk.Length);

        foreach (string blobName in chunk)
        {
            Response response = addOperation(batch, blobName);
            operations.Add((blobName, response));
        }

        // false means individual 4xx/5xx responses do not throw AggregateException.
        // A failure to submit the overall HTTP request still throws
        // RequestFailedException.
        await batchClient.SubmitBatchAsync(
            batch,
            throwOnAnyFailure: false,
            cancellationToken);

        foreach ((string blobName, Response response) in operations)
        {
            string? errorCode = null;
            response.Headers.TryGetValue("x-ms-error-code", out errorCode);

            results.Add(new BatchItemResult(
                blobName,
                response.Status,
                errorCode,
                response.ReasonPhrase));
        }
    }

    return results;
}
```

### Create the clients

For local development, use `DefaultAzureCredential`. In an Azure-hosted production application, use `ManagedIdentityCredential`.

```csharp
string accountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME")
    ?? throw new InvalidOperationException(
        "AZURE_STORAGE_ACCOUNT_NAME is required.");

string containerName = "my-container";

TokenCredential credential =
    Environment.GetEnvironmentVariable("DOTNET_ENVIRONMENT") == "Development"
        ? new DefaultAzureCredential()
        : new ManagedIdentityCredential();

var serviceClient = new BlobServiceClient(
    new Uri($"https://{accountName}.blob.core.windows.net"),
    credential);

BlobContainerClient containerClient =
    serviceClient.GetBlobContainerClient(containerName);

BlobBatchClient batchClient = serviceClient.GetBlobBatchClient();
```

The identity needs suitable data-plane RBAC, such as **Storage Blob Data Contributor**, scoped as narrowly as practical.

## Delete 500 blobs

This sends two batch HTTP requests because of the 256-operation limit:

```csharp
IReadOnlyList<string> blobsToDelete = Get500BlobNames();

IReadOnlyList<BatchItemResult> deleteResults;

try
{
    deleteResults = await SubmitInBatchesAsync(
        batchClient,
        blobsToDelete,
        (batch, blobName) => batch.DeleteBlob(
            containerClient.GetBlobClient(blobName).Uri,
            DeleteSnapshotsOption.IncludeSnapshots),
        cancellationToken);
}
catch (RequestFailedException ex)
{
    // The overall batch request could not be submitted or processed.
    // Do not treat this as an individual blob failure.
    Console.Error.WriteLine(
        $"Batch submission failed: HTTP {ex.Status}, " +
        $"code={ex.ErrorCode}, message={ex.Message}");
    throw;
}

ReportResults("delete", deleteResults);
```

`DeleteSnapshotsOption.IncludeSnapshots` deletes each blob together with its snapshots. Use `None` only when you know the blobs have no snapshots.

## Set 200 blobs from Hot to Cool

All 200 tier changes fit in one request:

```csharp
IReadOnlyList<string> blobsToCool = Get200BlobNames();

IReadOnlyList<BatchItemResult> tierResults;

try
{
    tierResults = await SubmitInBatchesAsync(
        batchClient,
        blobsToCool,
        (batch, blobName) => batch.SetBlobAccessTier(
            containerClient.GetBlobClient(blobName).Uri,
            AccessTier.Cool),
        cancellationToken);
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Batch submission failed: HTTP {ex.Status}, " +
        $"code={ex.ErrorCode}, message={ex.Message}");
    throw;
}

ReportResults("set tier", tierResults);
```

Tier changes apply to **block blobs** in supported Blob Storage or general-purpose v2 accounts.

## Report and retry partial failures

Batch execution is not atomic. For example, one delete can return `202`, another `404`, and another `409`.

```csharp
static void ReportResults(
    string operation,
    IReadOnlyList<BatchItemResult> results)
{
    BatchItemResult[] failures = results
        .Where(result => !result.Succeeded)
        .ToArray();

    Console.WriteLine(
        $"{operation}: {results.Count - failures.Length} succeeded, " +
        $"{failures.Length} failed.");

    foreach (BatchItemResult failure in failures)
    {
        Console.Error.WriteLine(
            $"{failure.BlobName}: HTTP {failure.Status} " +
            $"{failure.ErrorCode ?? "(no error code)"} " +
            $"{failure.ReasonPhrase}");
    }
}
```

Only retry failures that are actually transient, such as `408`, `429`, or selected `5xx` responses. Do not blindly retry authorization failures, lease conflicts, invalid blob types, or malformed requests.

## Convenience methods

For simple cases, these methods also submit one batch request:

```csharp
Response[] deleteResponses = await batchClient.DeleteBlobsAsync(
    blobUris,
    DeleteSnapshotsOption.IncludeSnapshots,
    cancellationToken);

Response[] tierResponses = await batchClient.SetBlobsAccessTierAsync(
    blobUris,
    AccessTier.Cool,
    cancellationToken: cancellationToken);
```

However, their individual subrequest failures are wrapped in an `AggregateException`. For reliable blob-name-to-response correlation, the explicit `CreateBatch`/`SubmitBatchAsync(..., false)` approach above is preferable.

## Limits and semantics

| Restriction | Limit or behavior |
|---|---|
| Subrequests per HTTP batch | Maximum **256** |
| Entire multipart request body | Maximum **4 MB** |
| Supported operations | Delete Blob and Set Blob Tier |
| Resource scope | Every subrequest must target the same storage account |
| Atomicity | None; each subrequest succeeds or fails independently |
| Execution order | Not guaranteed |
| Mixed operations | Keep deletes and tier changes in separate batches |
| Overall request failure | `RequestFailedException` |
| Subrequest failures with `throwOnAnyFailure: false` | Inspect each deferred `Response` |
| Subrequest failures with `throwOnAnyFailure: true` | `AggregateException` containing individual failures |

If unusually long blob URIs or added conditional headers cause the multipart body to approach 4 MB, use fewer than 256 operations per chunk.

Sources: [Azure Storage Blobs Batch client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet), [`SubmitBatchAsync`](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet), and [Blob Batch REST API](https://learn.microsoft.com/rest/api/storageservices/blob-batch).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response thoroughly covers batching, limits, custom response inspection, convenience APIs, and authentication, but omits the required exception-based partial-failure example.

### Strengths

- Correctly rejects a single 500-operation request and chunks it at 256.
- Provides robust blob-to-response correlation using deferred responses and throwOnAnyFailure: false.
- Correctly documents the 4 MB request limit, non-atomic behavior, supported operations, and account scope.

### Issues

- No catch/inspection example for AggregateException.
- No explicit handling of inner RequestFailedException values for failed subrequests.

## Grader Results

- batch-blob-operations.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `BlobBatchClient` from `Azure.Storage.Blobs.Batch` package: Pass
      - `BlobBatchClient.DeleteBlobsAsync()` for bulk delete: Pass
      - `BlobBatchClient.SetBlobsAccessTierAsync()` for bulk tier changes: Pass
      - Custom batch via `BlobBatchClient.CreateBatch()` and `SubmitBatchAsync()`: Pass
      - Batch size limits (256 operations per batch): Pass
      - Partial failure handling: `AggregateException` with per-operation status: Fail
      - `RequestFailedException` for individual operation failures within a batch: Fail
      - Authentication scopes for batch operations: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.7500** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-batch --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
