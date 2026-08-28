# Evaluation Report: storage-dp-dotnet-batch

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 366.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-batch` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 5/8 |
| Duration | 366.8s |
| Timestamp | 2026-08-27T17:09:27Z |
| Files Generated | 0 |
| Event Count | 4322 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 342.6s |
| Review | 23.1s |
| **Total** | **366.8s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=258833 out=4085 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 2629ms

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

### ✅ 2. view — 189712ms

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

### ✅ 3. view — 188332ms

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

### ❌ 4. azure-get_azure_bestpractices (via azure) — 189715ms

**Input:**

```json
{"intent":"Provide current best practices for .NET code using Azure.Storage.Blobs.Batch BlobBatchClient for bulk delete, bulk set access tier, authentication, batching, and partial failure handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 189717ms

**Input:**

```json
{"intent":"Find official documentation for Azure.Storage.Blobs.Batch BlobBatchClient .NET, SubmitBatchAsync behavior, partial failures, maximum subrequests per batch, request size restrictions, and same-account requirements.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_search (via github-mcp-server) — 65944ms

**Input:**

```json
{"query":"Official Microsoft documentation Azure.Storage.Blobs.Batch BlobBatchClient .NET SubmitBatchAsync throwOnAnyFailure partial failures AggregateException Response[] maximum 256 4 MiB"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 7. web_search (via github-mcp-server) — 64429ms

**Input:**

```json
{"query":"site:learn.microsoft.com REST Blob Batch maximum 256 subrequests 4 MiB same storage account Delete Blob Set Blob Tier"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Based on Microsoft documentation for Azure Blob Storage REST API's Blob Batch operation:\n- **Blob Batch operation** allows you to execute up to **256 subrequests** in a single batch request.\n- The total size of the entire batch request body must not exceed **4 MiB (mebibytes)**.\n- All subrequests within the batch must target blobs in the **same storage account**.\n- The supported subrequest types are: **Delete Blob** and **Set Blob Tier** (for block blobs).\n\nThis lets you efficiently perform bulk deletes or tier changes with fewer HTTP roundtrips, but you must chunk your work accordingly if you have more than 256 blobs or if the request body would exceed 4 MiB【3:0†source】【3:8†source】.\n\nYou can find more details in the official [Blob Batch (REST API) documentation on Microsoft Learn](https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch).","annotations":[{"text":"","start_index":667,"end_index":679,"url_citation":{"title":"Blob Batch (REST API) - Azure Storage | Microsoft Learn","url":"https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch"}},{"text":"","start_index":679,"end_index":691,"url_citation":{"title":"Docs: batch limit of `azure.storage.blob.ContainerClient.delete_blobs ...","url":"https://github.com/Azure/azure-sdk-for-python/issues/22821"}}]},"bing_searches":[{"text":"learn.microsoft.com Blob Batch limits 256 subrequests 4 MiB Delete Blob Set Blob Tier same storage account","url":"https://www.bing.com/search?q=learn.microsoft.com+Blob+Batch+limits+256+subrequests+4+MiB+Delete+Blob+Set+Blob+Tier+same+storage+account"}],"annotations":null}
```

### ✅ 8. web_search (via github-mcp-server) — 62965ms

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api azure.storage.blobs.specialized.blobbatchclient DeleteBlobsAsync SetBlobsAccessTierAsync signature"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here are the official details for `BlobBatchClient.SubmitBatchAsync` in the Azure.Storage.Blobs.Batch .NET SDK:\n\n- **Method signature**:\n  ```csharp\n  public virtual Task<Response> SubmitBatchAsync(\n      BlobBatch batch,\n      bool throwOnAnyFailure = false,\n      CancellationToken cancellationToken = default\n  )\n  ```\n- **throwOnAnyFailure**:\n  - If `false`, sub-operation failures do _not_ throw, you need to inspect the returned response to check for partial failures.\n  - If `true`, if *any* sub-operation in the batch fails, an `AggregateException` is thrown, wrapping all exceptions from failed operations【3:0†source】【3:5†source】.\n\n- **Partial Failure Handling**:\n  - When `throwOnAnyFailure=false`, the response should be inspected for all sub-operation statuses. Failures and successes are mixed in the result array.\n  - When `throwOnAnyFailure=true`, _any_ sub-operation failure triggers an `AggregateException`, containing one or more exceptions for subprocesses.\n\n- **Response[]**:\n  - The result encapsulates the status for every sub-operation, so you must inspect each response to determine which ones succeeded/failed when not throwing on any failure.\n\n- **Limits**:\n  - **Maximum sub-operations per batch:** 256\n  - **Maximum batch request body size:** 4 MiB (4,194,304 bytes)\n  - These limits are enforced by the service and SDK; attempts to exceed them will result in errors【3:1†source】.\n\n- **Notable Behaviors**:\n  - Batch execution order is not guaranteed.\n  - Batch execution is _not atomic_. Some sub-operations may succeed and others may fail.\n  - All blobs operated on must be in the same storage account.\n\n**Docs and references:**\n- [BlobBatchClient.SubmitBatchAsync (Microsoft Docs)](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync)\n- [Azure Storage Batch client library docs](https://learn.microsoft.com/dotnet/api/azure.stor
... (truncated)
```

### ✅ 9. github-mcp-server-search_code (via github-mcp-server) — 67133ms

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":10,"query":"\"throwOnAnyFailure: false\" BlobBatchClient language:C#"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"AzureBlobStorage.cs","path":"src/Headless.Blobs.Azure/AzureBlobStorage.cs","repository":"xshaheen/headless-framework","text_matches":[{"fragment":"        {\n            var batchClient = blobServiceClient.GetBlobBatchClient();\n","matches":[{"indices":[61,76],"text":"BlobBatchClient"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/827351351/contents/src/Headless.Blobs.Azure/AzureBlobStorage.cs?ref=076479c9b120df5652fd1217c5d4d9ae2419190b","property":"content"},{"fragment":"            await batchClient\n                .SubmitBatchAsync(batch, throwOnAnyFailure: false, cancellationToken)\n                .ConfigureAwait(false);","matches":[{"indices":[71,95],"text":"throwOnAnyFailure: false"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/827351351/contents/src/Headless.Blobs.Azure/AzureBlobStorage.cs?ref=076479c9b120df5652fd1217c5d4d9ae2419190b","property":"content"}]},{"name":"BlobBatchClientTests.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"\n            BlobBatchClient client = scenario.GetBlobBatchClient();\n            AggregateException exes = Assert.ThrowsAsync<AggregateException>(","matches":[{"indices":[13,28],"text":"BlobBatchClient"},{"indices":[50,65],"text":"BlobBatchClient"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs?ref=80616424417f5335fcf8d3d3646afa25772dd86b","property":"content"},{"fragment":"            Response response3 = batch.DeleteBlob(bad[1]);\n            Response response = await client.SubmitBatchAsync(batch, throwOnAnyFailure: false);\n","matches":[{"indices":[128,152],"text":"throwOnAnyFailure: false"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/stora
... (truncated)
```

### ✅ 10. github-mcp-server-get_file_contents (via github-mcp-server) — 7754ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: ca94cdccdcbf1b3597a063d38ccbc10b20bea029)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Azure.Core.TestFramework;
using Azure.Core.TestFramework.Models;
using Azure.Storage.Blobs.Batch;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;
using Azure.Storage.Blobs.Tests;
using Azure.Storage.Sas;
using Azure.Storage.Test;
using Azure.Storage.Test.Shared;
using Moq;
using NUnit.Framework;

namespace Azure.Storage.Blobs.Test
{
    public class BlobBatchClientTests : BlobTestBase
    {
        private static Regex pattern = new Regex(@"sig=\S+\s", RegexOptions.Compiled);

        public BlobBatchClientTests(bool async, BlobClientOptions.ServiceVersion serviceVersion)
            : base(async, serviceVersion, null /* RecordedTestMode.Record /* to re-record */)
        {
            // Batch delimiters are random so disable body comparison
            CompareBodies = false;
            BodyRegexSanitizers.Add(new BodyRegexSanitizer(@"sig=(?<group>.*?)(?=\s+)")
            {
                GroupForReplace = "group",
                Value = SanitizeValue
            });
        }

        [SetUp]
        public void ResetDiagnostics()
        {
            // Some tests currently disable test diagnostics so we turn it back on
            TestDiagnostics = true;

            // Force multipart boundaries to be repeatably random
            Multipart.GetRandomGuid = () => Recording.Random.NewGuid();

            // Remove the x-ms-client-request-id headers from the batch
            // sub-operations because there's no convenient way to set them to
            // predictable values
            ((RemoveVersionHeaderPolicy)RemoveVersionHeaderPolicy.Shared).RemoveClientRequestIdHeaders = true;
        }

        #r
... (truncated)
```

### ✅ 11. github-mcp-server-get_file_contents (via github-mcp-server) — 7329ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/README.md","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: e85875c4adcc3b71ba671441d7d7b3e0d55427ea)# Azure Storage Blobs Batch client library for .NET

> Server Version: 2021-02-12, 2020-12-06, 2020-10-02, 2020-08-04, 2020-06-12, 2020-04-08, 2020-02-10, 2019-12-12, 2019-07-07, and 2019-02-02

Azure Blob storage is Microsoft's object storage solution for the cloud. Blob
storage is optimized for storing massive amounts of unstructured data.  This
library allows you to batch multiple Azure Blob Storage operations in a single request.

[Source code][source] | [Package (NuGet)][package] | [API reference documentation][docs] | [REST API documentation][rest_docs] | [Product documentation][product_docs]

## Getting started

### Install the package

Install the Azure Storage Blobs Batch client library for .NET with [NuGet][nuget]:

```dotnetcli
dotnet add package Azure.Storage.Blobs.Batch
```

### Prerequisites

You need an [Azure subscription][azure_sub] and a
[Storage Account][storage_account_docs] to use this package.

To create a new Storage Account, you can use the [Azure Portal][storage_account_create_portal],
[Azure PowerShell][storage_account_create_ps], or the [Azure CLI][storage_account_create_cli].
Here's an example using the Azure CLI:

```Powershell
az storage account create --name MyStorageAccount --resource-group MyResourceGroup --location westus --sku Standard_LRS
```

### Authenticate the client

In order to interact with the Azure Blobs Storage service for batch operations, you'll need to create an instance of the BlobServiceClient class.  The [Azure Identity library][identity] makes it easy to add Azure Active Directory support for authenticating Azure SDK clients with their corresponding Azure services.

```C# Snippet:SampleSnippetsBlobBatch_Auth
// Create a BlobServiceClient that will authenticate through Active Directory
Uri accountUri = new Uri("https://MYSTORAGEACCOUNT.blob.core.windows.net/");
BlobServiceClient client = new BlobServiceClient(accountUri, new DefaultAzu
... (truncated)
```

### ✅ 12. web_fetch — 6922ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch:
---
layout: Conceptual
title: Blob Batch (REST API) - Azure Storage | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch
enable_rest_try_it: true
rest_product: Azure
uhfHeaderId: azure
breadcrumb_path: ../../breadcrumb/toc.json
ms.author: seanmcc
manager: smmark
author: seanmcc-msft
ms.topic: reference
ms.devlang: rest-api
ms.date: 2023-11-28T00:00:00.0000000Z
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
ms.service: azure-storage
description: The Blob Batch operation allows multiple API calls to be embedded into a single HTTP request.
locale: en-us
moniker_definition_rel: ../../.monikers.Azure.AzureRestApi.json
document_id: 897ed44a-dcc2-f335-5b9c-4c21a9c3e0cb
document_version_independent_id: 9b3c52da-6533-8351-c14f-fe20daec613b
updated_at: 2024-07-16T22:05:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-rest-apis/blob/live/docs-ref-conceptual/storageservices/blob-batch.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-rest-apis/blob/ad1abc07a0db92c69443f571a14c54413100e27f/docs-ref-conceptual/storageservices/blob-batch.md
git_commit_id: ad1abc07a0db92c69443f571a14c54413100e27f
site_name: Docs
depot_name: Azure.AzureRestApi
page_type: conceptual
toc_rel: ../azure/toc.json
feedback_system: None
feedback_product_url: ''
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 2098
asset_id: api/storageservices/blob-batch
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs-ref-conceptual/storageservices/blob-batch.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/a06cf482-4ca9-4582-a142-bcf842258d42
platformId: 5c991847-08bf-68ef-e2bf-692985c1f515
---

# Blob Batch (REST API) - Azure Stora
... (truncated)
```

### ✅ 13. web_fetch — 6502ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch"}
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

### ✅ 14. rg — 1117ms

**Input:**

```json
{"-A":20,"-B":12,"-n":true,"head_limit":10,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787850873731-copilot-tool-output-0310f8a75d794b82895b4fe3844d66df.txt","pattern":"throwOnAnyFailure: false"}
```

**Output:**

```
[grep content: 132 matches across 1 file(s) under C:/Users/weidxu/AppData/Local/Temp]

1787850873731-copilot-tool-output-0310f8a75d794b82895b4fe3844d66df.txt (132 match(es)):
  589-         [RecordedTest]
  590-         public async Task Delete_OneFails_NoThrow()
  591-         {
  592-             await using TestScenario scenario = Scenario();
  593-             BlobClient[] good = await scenario.CreateBlobsAsync(2);
  594-             Uri[] bad = scenario.GetInvalidBlobUris(1);
  595- 
  596-             BlobBatchClient client = scenario.GetBlobBatchClient();
  597-             using BlobBatch batch = client.CreateBatch();
  598-             Response response1 = batch.DeleteBlob(good[0].Uri);
  599-             Response response2 = batch.DeleteBlob(good[1].Uri);
  600-             Response response3 = batch.DeleteBlob(bad[0]);
  601:             Response response = await client.SubmitBatchAsync(batch, throwOnAnyFailure: false);
  602- 
  603-             Assert.AreEqual(3, batch.RequestCount);
  604-             scenario.AssertStatus(202, response, response1, response2);
  605-             scenario.AssertStatus(404, response3);
  606-             await scenario.AssertDeleted(good);
  607-         }
  608- 
  609-         [RecordedTest]
  610-         public async Task Delete_MultipleFail()
  611-         {
  612-             await using TestScenario scenario = Scenario();
  613-             BlobClient[] good = await scenario.CreateBlobsAsync(1);
  614-             Uri[] bad = scenario.GetInvalidBlobUris(2);
  615- 
  616-             BlobBatchClient client = scenario.GetBlobBatchClient();
  617-             using BlobBatch batch = client.CreateBatch();
  618-             Response response1 = batch.DeleteBlob(good[0].Uri);
  619-             Response response2 = batch.DeleteBlob(bad[0]);
  620-             Response response3 = batch.DeleteBlob(bad[1]);
  621-             AggregateException exes = Assert.ThrowsAsync<AggregateException>(
  648-         [RecordedTest
... (truncated)
```

## Copilot Response

I’ll ground the sample in the current .NET batch API and its documented service limits.

The key constraint is already material: 500 deletes cannot fit in one Blob Batch request, so the implementation must split them into two requests while preserving per-blob results.

A single Blob Batch request supports **at most 256 operations**, so deleting 500 blobs requires **two HTTP batch requests**: 256 and 244. The 200 tier changes fit in one request.

Install the packages:

```powershell
dotnet add package Azure.Storage.Blobs.Batch
dotnet add package Azure.Identity
```

Use explicit `BlobBatch` objects rather than only the convenience methods when you need to associate every blob URI with its individual response:

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

const int MaxOperationsPerBatch = 256;

Uri accountUri = new(
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_URL")
    ?? throw new InvalidOperationException(
        "AZURE_STORAGE_ACCOUNT_URL is required. " +
        "Example: https://myaccount.blob.core.windows.net/"));

// DefaultAzureCredential is convenient for local development.
// In Azure-hosted production code, prefer ManagedIdentityCredential.
TokenCredential credential = new DefaultAzureCredential();

var serviceClient = new BlobServiceClient(accountUri, credential);
BlobBatchClient batchClient = serviceClient.GetBlobBatchClient();
```

Define a result type for per-blob outcomes:

```csharp
public sealed record BlobOperationResult(
    Uri BlobUri,
    bool Succeeded,
    int Status,
    string ReasonPhrase,
    string? ErrorCode);
```

### Bulk delete with partial-failure handling

```csharp
static async Task<IReadOnlyList<BlobOperationResult>> DeleteBlobsAsync(
    BlobBatchClient batchClient,
    IEnumerable<Uri> blobUris,
    CancellationToken cancellationToken = default)
{
    var results = new List<BlobOperationResult>();

    foreach (Uri[] chunk in blobUris.Chunk(MaxOperationsPerBatch))
    {
        using BlobBatch batch = batchClient.CreateBatch();

        // DeleteBlob returns a deferred Response. The SDK fills it in after
        // SubmitBatchAsync completes.
        var operations = chunk
            .Select(uri => (
                Uri: uri,
                Response: batch.DeleteBlob(
                    uri,
                    DeleteSnapshotsOption.IncludeSnapshots)))
            .ToArray();

        try
        {
            await batchClient.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);
        }
        catch (RequestFailedException ex)
        {
            // This indicates failure of the top-level batch request, such as
            // authentication, networking, malformed request, or service failure.
            Console.Error.WriteLine(
                $"Batch request failed: status={ex.Status}, " +
                $"errorCode={ex.ErrorCode}, message={ex.Message}");

            throw;
        }

        foreach (var operation in operations)
        {
            Response response = operation.Response;
            bool succeeded = response.Status is >= 200 and < 300;

            string? errorCode = null;
            if (!succeeded &&
                response.Headers.TryGetValue("x-ms-error-code", out string value))
            {
                errorCode = value;
            }

            results.Add(new BlobOperationResult(
                operation.Uri,
                succeeded,
                response.Status,
                response.ReasonPhrase,
                errorCode));
        }
    }

    return results;
}
```

For 500 blobs:

```csharp
BlobContainerClient container =
    serviceClient.GetBlobContainerClient("my-container");

IReadOnlyList<string> blobNames = GetThe500BlobNames();

Uri[] blobUris = blobNames
    .Select(name => container.GetBlobClient(name).Uri)
    .ToArray();

IReadOnlyList<BlobOperationResult> deleteResults =
    await DeleteBlobsAsync(batchClient, blobUris);

foreach (BlobOperationResult failure in deleteResults.Where(r => !r.Succeeded))
{
    Console.Error.WriteLine(
        $"Delete failed: blob={failure.BlobUri}, " +
        $"status={failure.Status}, error={failure.ErrorCode}, " +
        $"reason={failure.ReasonPhrase}");
}

Console.WriteLine(
    $"Deleted: {deleteResults.Count(r => r.Succeeded)}, " +
    $"failed: {deleteResults.Count(r => !r.Succeeded)}");
```

This sends two batch HTTP requests because 500 exceeds the 256-operation limit.

### Bulk change access tier

```csharp
static async Task<IReadOnlyList<BlobOperationResult>> SetAccessTierAsync(
    BlobBatchClient batchClient,
    IEnumerable<Uri> blobUris,
    AccessTier tier,
    CancellationToken cancellationToken = default)
{
    var results = new List<BlobOperationResult>();

    foreach (Uri[] chunk in blobUris.Chunk(MaxOperationsPerBatch))
    {
        using BlobBatch batch = batchClient.CreateBatch();

        var operations = chunk
            .Select(uri => (
                Uri: uri,
                Response: batch.SetBlobAccessTier(uri, tier)))
            .ToArray();

        try
        {
            await batchClient.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Tier batch request failed: status={ex.Status}, " +
                $"errorCode={ex.ErrorCode}, message={ex.Message}");

            throw;
        }

        foreach (var operation in operations)
        {
            Response response = operation.Response;
            bool succeeded = response.Status is >= 200 and < 300;

            string? errorCode = null;
            if (!succeeded &&
                response.Headers.TryGetValue("x-ms-error-code", out string value))
            {
                errorCode = value;
            }

            results.Add(new BlobOperationResult(
                operation.Uri,
                succeeded,
                response.Status,
                response.ReasonPhrase,
                errorCode));
        }
    }

    return results;
}
```

Change 200 block blobs from Hot to Cool:

```csharp
IReadOnlyList<string> namesToCool = GetThe200BlobNames();

Uri[] tierUris = namesToCool
    .Select(name => container.GetBlobClient(name).Uri)
    .ToArray();

IReadOnlyList<BlobOperationResult> tierResults =
    await SetAccessTierAsync(
        batchClient,
        tierUris,
        AccessTier.Cool);

foreach (BlobOperationResult failure in tierResults.Where(r => !r.Succeeded))
{
    Console.Error.WriteLine(
        $"Set tier failed: blob={failure.BlobUri}, " +
        $"status={failure.Status}, error={failure.ErrorCode}");
}
```

`Set Blob Tier` applies to **block blobs**. Setting an unsupported blob type or an invalid tier produces an individual failed subresponse.

### Failure modes

`throwOnAnyFailure` controls only how **subrequest** failures are reported:

| Setting | Behavior |
|---|---|
| `false` | Returns normally after a valid batch response; inspect each deferred `Response` for success or failure. Best for partial-success processing. |
| `true` | Throws `AggregateException` when one or more subrequests fail. Its inner exceptions are generally `RequestFailedException` instances. |
| Either setting | A failure of the top-level batch request can throw `RequestFailedException`. |

If you prefer exception-based handling:

```csharp
try
{
    await batchClient.SubmitBatchAsync(
        batch,
        throwOnAnyFailure: true,
        cancellationToken);
}
catch (AggregateException ex)
{
    foreach (RequestFailedException failure
             in ex.InnerExceptions.OfType<RequestFailedException>())
    {
        Console.Error.WriteLine(
            $"Subrequest failed: status={failure.Status}, " +
            $"errorCode={failure.ErrorCode}, message={failure.Message}");
    }
}
```

Using `throwOnAnyFailure: false` is generally better when you need exact success/failure mapping because the deferred responses remain associated with their blob URIs.

### Limits and semantics

| Constraint | Limit or behavior |
|---|---|
| Operations per request | Maximum **256** |
| Request body | Maximum **4 MB** |
| Supported operations | `Delete Blob` and `Set Blob Tier` |
| Operation types | All subrequests in one REST batch must be the same type |
| Scope | Every blob must belong to the same storage account |
| Atomicity | Not atomic; some operations can succeed while others fail |
| Ordering | Execution order is not guaranteed |
| SDK splitting | The SDK does not automatically split oversized input |
| Response | Top-level request normally returns `202`; each subrequest has its own status |

Chunking at 256 handles the operation-count limit, but extremely long blob names and headers can hit the 4 MB body limit before 256 operations. If your blob names are unusually large, use a smaller conservative chunk size such as 100–200 and retry only failed, retryable operations—not successful ones.

References: [Azure Storage Blobs Batch client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet) and [Blob Batch REST API](https://learn.microsoft.com/rest/api/storageservices/blob-batch).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 8 criteria passed. The custom batch implementation and partial-failure discussion are strong, but the required convenience APIs and authentication-scope guidance are missing.

### Strengths

- Correctly uses BlobBatchClient and custom BlobBatch submission.
- Correctly explains the 256-operation and 4 MB limits.
- Provides practical partial-success handling with deferred responses and AggregateException handling.

### Issues

- Does not demonstrate BlobBatchClient.DeleteBlobsAsync().
- Does not demonstrate BlobBatchClient.SetBlobsAccessTierAsync().
- Does not document the Azure Storage token scope or required data-plane RBAC permissions.

## Grader Results

- batch-blob-operations.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/8)
      - `BlobBatchClient` from `Azure.Storage.Blobs.Batch` package: Pass
      - `BlobBatchClient.DeleteBlobsAsync()` for bulk delete: Fail
      - `BlobBatchClient.SetBlobsAccessTierAsync()` for bulk tier changes: Fail
      - Custom batch via `BlobBatchClient.CreateBatch()` and `SubmitBatchAsync()`: Pass
      - Batch size limits (256 operations per batch): Pass
      - Partial failure handling: `AggregateException` with per-operation status: Pass
      - `RequestFailedException` for individual operation failures within a batch: Pass
      - Authentication scopes for batch operations: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 62% | 1.00 | 0.6250 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6250** | **62.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-batch --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
