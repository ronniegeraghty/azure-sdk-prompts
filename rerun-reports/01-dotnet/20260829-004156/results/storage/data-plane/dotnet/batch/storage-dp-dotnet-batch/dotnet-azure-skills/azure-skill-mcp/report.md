# Evaluation Report: storage-dp-dotnet-batch

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 141.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-batch` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 5/8 |
| Duration | 141.6s |
| Timestamp | 2026-08-28T18:43:44Z |
| Files Generated | 0 |
| Event Count | 4020 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 107.4s |
| Review | 33.0s |
| **Total** | **141.6s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=438481 out=4566 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 694ms

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

### ✅ 2. view — 4827ms

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

### ✅ 3. view — 4484ms

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

### ✅ 4. azure-get_azure_bestpractices (via azure) — 4145ms

**Input:**

```json
{"command":"get_best_practices","intent":"Provide modern .NET Azure Blob Storage batch operations code using Azure.Storage.Blobs.Batch, including authentication and partial failure handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 5183ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official documentation for Azure.Storage.Blobs.Batch BlobBatchClient .NET SubmitBatchAsync DeleteBlobsAsync SetBlobsAccessTierAsync partial failures and batch limits.","parameters":{"query":"Azure.Storage.Blobs.Batch BlobBatchClient .NET SubmitBatchAsync partial failures limits 256 4 MiB"}}
```

**Output:**

```
{"results":[{"title":"Blob Batch","content":"# Blob Batch\n## Remarks\nOne of the main benefits of using a batch request is the reduction in the number of connections a client has to open. Note the following restrictions:\n1. Supported subrequests in the batch are \u0060Set Blob Tier\u0060 (for block blobs) and \u0060Delete Blob\u0060.\n2. Only supports up to 256 subrequests in a single batch. The size of the body for a batch request can\u0027t exceed 4 MB.\n3. An empty batch request fails with code 400 (Bad Request).\n4. There are no guarantees on the order of execution of the batch subrequests.\n5. Batch subrequest execution isn\u0027t atomic. Each subrequest runs independently.\n6. Each subrequest must be for a resource within the same storage account. A single batch request doesn\u0027t support running requests from different storage accounts.\n7. A nested request body isn\u0027t supported.\n8. If the server fails to parse the request body, the entire batch fails, and no request will be run.\n9. Note that [Account SAS](https://learn.microsoft.com/rest/api/storageservices/create-account-sas) is the only shared access signature type supported by \u0060Blob Batch\u0060, when the batch isn\u0027t using \u0060restype=container\u0060.\n### Scope all subrequests to a specific container\nBeginning with REST version 2020-04-08, the \u0060Blob Batch\u0060 API supports scoping subrequests to a specified container. When the request URI includes the container name and the \u0060restype=container\u0060 parameter, then each subrequest must apply to the same container. If the container name specified for a subrequest doesn\u0027t match the container name provided in URI, the service returns error code 400 (Bad Request).\nAll authorization mechanisms supported for a container are valid for a \u0060Blob Batch\u0060 operation that is scoped to the container. Each subrequest sends an authorization header to the service.","contentUrl":"https://learn.microsoft.com/rest/api/storageser
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 2346ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a .NET Blob Storage batch operations example.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. rg — 2663ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787942651629-copilot-tool-output-e2b5c5dd2f4c41f1b5ccb32837e7650a.txt","pattern":"256|4 MB|partial|throwOnAnyFailure|SubmitBatchAsync|DeleteBlobsAsync|SetBlobsAccessTierAsync|RequestFailedException"}
```

**Output:**

```
Output too large to read at once (49.5 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787942659364-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787942651629-copilot-tool-output-e2b5c5dd2f4c41f1b5ccb32837e7650a.txt:{"results":[{"title":"Blob Batch","content":"# Blob Batch\n## Remarks\nOne of the main benefits of using a batch request is the reduction in the number of connections a client has to open. Note the following restrictions:\n1. Supported subrequests in the batch are \u0060Set Blob Tier\u0060 (for block blobs) and \u0060Delete Blob\u0060.\n2. Only supports up to 256 subrequests in a single batc
```

### ✅ 8. azure-documentation (via azure) — 2978ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official .NET API behavior for BlobBatchClient.SubmitBatchAsync including throwOnAnyFailure false and how to inspect individual BlobBatch results.","parameters":{"query":"site:learn.microsoft.com/dotnet/api Azure.Storage.Blobs.Specialized.BlobBatchClient SubmitBatchAsync throwOnAnyFailure BlobBatch response partial failures"}}
```

**Output:**

```
{"results":[{"title":"BlobBatchClient.SubmitBatchAsync Method-Definition","content":"# BlobBatchClient.SubmitBatchAsync Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure.Storage.Blobs.Specialized](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.Storage.Blobs.Batch.dll\r\n\r\n- Package:\r\n    - Azure.Storage.Blobs.Batch v12.26.0\r\n\r\n- Package:\r\n    - Azure.Storage.Blobs.Batch v12.27.0-beta.1\r\n\r\n- Source:\r\n    - [BlobBatchClient.cs](https://github.com/Azure/azure-sdk-for-net/blob/8608dd90245dcc1e8d9a30f4fdb6494d8d268d82/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs)\r\n\r\n- Source:\r\n    - [BlobBatchClient.cs](https://github.com/Azure/azure-sdk-for-net/blob/62cf4bdea6158c44231a76da938a736d91eb11d6/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs)\r\n\r\nSubmit a [BlobBatch](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatch?view=azure-dotnet) of sub-operations.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic virtual System.Threading.Tasks.Task\u003CAzure.Response\u003E SubmitBatchAsync(Azure.Storage.Blobs.Specialized.BlobBatch batch, bool throwOnAnyFailure = false, System.Threading.CancellationToken cancellationToken = default);\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- batch\r\n    - [BlobBatch](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatch?view=azure-dotnet)\r\n\r\nA [BlobBatch](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatch?view=azure-dotnet) of sub-operations.\r\n\r\n- throwOnAnyFailure\r\n    - [Boolean](https://learn.microsoft.com/en-us/dotnet/api/system.boolean)\r\n\r\nA value indicating whether or not to throw exceptions for sub-operation failures.\r\n\r\n- cancellationToken\r\n    - [CancellationToken](https://learn.microsoft.com/en-us/dotnet/api/system.threading.cancellationtoken)\r\n\r\nOptional [CancellationToken](https://le
... (truncated)
```

### ✅ 9. powershell — 2641ms

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787942660310-copilot-tool-output-47c4a18804f94c17bb49f91c2a0dc490.txt' | ConvertFrom-Json; $j.results | Where-Object { $_.title -match 'SubmitBatchAsync|CreateBatch|DeleteBlob|SetBlobAccessTier' } | Select-Object title,content,url | Format-List","description":"Extract relevant Blob Batch documentation","initial_wait":30}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-47c4a18804f94c17bb49f91c2a0dc490.txt' | ConvertFrom-Json; $j.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 10. azure-documentation (via azure) — 3264ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find the official Azure Blob Storage .NET batch client sample showing BlobBatch CreateBatch DeleteBlob SetBlobAccessTier SubmitBatchAsync and per-operation response inspection.","parameters":{"query":"site:learn.microsoft.com azure blob batch client library .NET BlobBatch CreateBatch DeleteBlob SetBlobAccessTier Response[] sample"}}
```

**Output:**

```
{"results":[{"title":"Azure Storage samples using .NET client libraries","content":"# Azure Storage samples using .NET client libraries\n## Blob samples\n| Topic | Developer guide | Samples on GitHub| \n|  --- | --- | ---  |\n| Access tiers | [Set or change a block blob\u0027s access tier](https://learn.microsoft.com/azure/storage/blobs/storage-blob-use-access-tier-dotnet) |  |\n| Append blob | [Append data to an append blob](https://learn.microsoft.com/azure/storage/blobs/storage-blob-append) |  |\n| Batching |  | [Delete several blobs in one request](https://github.com/Azure/azure-sdk-for-net/tree/master/sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs#L22)[Set several blob access tiers in one request](https://github.com/Azure/azure-sdk-for-net/blob/master/sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs#L56)[Fine-grained control in a batch request](https://github.com/Azure/azure-sdk-for-net/blob/master/sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs#L90)[Catch errors from a failed sub-operation](https://github.com/Azure/azure-sdk-for-net/blob/master/sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs#L136) |\n| Troubleshooting |  | [Trigger a recoverable error using a container client](https://github.com/Azure/azure-sdk-for-net/tree/master/sdk/storage/Azure.Storage.Blobs/samples/Sample01b_HelloWorldAsync.cs#L166) |\n","contentUrl":"https://learn.microsoft.com/azure/storage/common/storage-samples-dotnet#blob-samples"},{"title":"BlobBatch class-Method Details","content":"# BlobBatch class\r\n\r\n- Package:\r\n    - [@azure/storage-blob](https://learn.microsoft.com/en-us/javascript/api/preview-docs/@azure/storage-blob/)\r\n\r\nA BlobBatch represents an aggregated set of operations on blobs. Currently, only \u0060delete\u0060 and \u0060setAccessTier\u0060 are supported.\r\n\r\n## Method Details\r\n\r\n### deleteBlob(string, StorageSharedKeyCredential | AnonymousCredential | 
... (truncated)
```

### ✅ 11. web_fetch — 2944ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/README.md"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/README.md:
# Azure Storage Blobs Batch client library for .NET

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
Uri account
... (truncated)
```

### ❌ 12. web_fetch — 4016ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ✅ 13. web_fetch — 4013ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatch.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatch.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Threading;
using Azure.Core;
using Azure.Core.Pipeline;
using Azure.Storage.Blobs.Batch;
using Azure.Storage.Blobs.Batch.Models;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Shared;

namespace Azure.Storage.Blobs.Specialized
{
    /// <summary>
    /// A <see cref="BlobBatch"/> allows you to batch multiple Azure Storage
    /// operations in a single request via <see cref="BlobBatchClient.SubmitBatch"/>.
    ///
    /// For more information, see
    /// <see href="https://docs.microsoft.com/en-us/rest/api/storageservices/blob-batch">
    /// Blob Batch</see>.
    /// </summary>
    public class BlobBatch : IDisposable
    {
        /// <summary>
        /// The number of pending requests in the batch.
        /// </summary>
        public int RequestCount => _messages.Count;

        /// <summary>
        /// If this BlobBatch is container scoped.
        /// </summary>
        private readonly bool _isContainerScoped;

        /// <summary>
        /// If this BlobBatch is container scoped.
        /// </summary>
        internal bool IsContainerScoped => _isContainerScoped;

        /// <summary>
        /// The <see cref="BlobBatchClient"/> associated with this batch.  It
        /// provides the Uri, BatchOperationPipeline, etc.
        /// </summary>
        private readonly BlobBatchClient _client;

        /// <summary>
        /// Storage requires each batch request to contain the same type of
        /// operation.
        /// </summary>
        private BlobBatchOperationType? _operationType;

        /// <summary>
        /// The list of m
... (truncated)
```

### ✅ 14. web_fetch — 3625ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.IO;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Pipeline;
using Azure.Storage.Blobs.Batch;
using Azure.Storage.Blobs.Models;
using Microsoft.TypeSpec.Generator.Customizations;

#pragma warning disable SA1402  // File may only contain a single type

namespace Azure.Storage.Blobs.Specialized
{
    /// <summary>
    /// The <see cref="BlobBatchClient"/> allows you to batch multiple Azure
    /// Storage operations in a single request.
    /// </summary>
    // CUSTOM:
    // - Suppress unused methods, ctors, and fields.
    [CodeGenType("BlobBatchClient")]
    [CodeGenSuppress("_endpoint", typeof(Uri))]
    [CodeGenSuppress("_cachedServiceRestClient", typeof(ServiceRestClient))]
    [CodeGenSuppress("_cachedContainerRestClient", typeof(ContainerRestClient))]
    [CodeGenSuppress("BlobBatchClient", typeof(Uri), typeof(TokenCredential))]
    [CodeGenSuppress("BlobBatchClient", typeof(HttpPipelinePolicy), typeof(Uri), typeof(BlobBatchClientOptions))]
    [CodeGenSuppress("BlobBatchClient", typeof(Uri), typeof(TokenCredential), typeof(BlobBatchClientOptions))]
#pragma warning disable SCME0002 // Type is for evaluation purposes only and is subject to change or removal in future updates. Suppress this diagnostic to proceed.
    [CodeGenSuppress("BlobBatchClient", typeof(BlobBatchClientSettings))]
#pragma warning restore SCME0002 // Type is for evaluation purposes only and is subject to change or removal in future updates. Suppress this diagnostic to proceed.
    [CodeGenSuppress("GetServiceRestClient")]
    [CodeGenSuppress("GetContainerRestClie
... (truncated)
```

### ✅ 15. web_fetch — 3262ms

**Input:**

```json
{"max_length":2000,"url":"https://api.nuget.org/v3-flatcontainer/azure.storage.blobs.batch/index.json"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Blobs.Batch/samples/Sample03b_BatchingAsync.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.IO;
using System.Text;
using System.Threading.Tasks;
using Azure.Storage;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;
using NUnit.Framework;

namespace Azure.Storage.Blobs.Samples
{
    /// <summary>
    /// Basic Azure Blob Storage batching samples
    /// </summary>
    public class Sample03b_BatchingAsync : SampleTest
    {
        /// <summary>
        /// Delete several blobs in one request.
        /// </summary>
        [Test]
        public async Task BatchDeleteAsync()
        {
            // Get a connection string to our Azure Storage account.
            string connectionString = ConnectionString;

            // Get a reference to a container named "sample-container" and then create it
            BlobServiceClient service = new BlobServiceClient(connectionString);
            BlobContainerClient container = service.GetBlobContainerClient(Randomize("sample-container"));
            await container.CreateAsync();
            try
            {
                // Create three blobs named "foo", "bar", and "baz"
                BlobClient foo = container.GetBlobClient("foo");
                BlobClient bar = container.GetBlobClient("bar");
                BlobClient baz = container.GetBlobClient("baz");
                await foo.UploadAsync(BinaryData.FromString("Foo!"));
                await bar.UploadAsync(BinaryData.FromString("Bar!"));
                await baz.UploadAsync(BinaryData.FromString("Baz!"));

                // Delete all three blobs at once
                BlobBatchClient batch = service.GetBlobBatchClient();
         
... (truncated)
```

### ✅ 16. github-mcp-server-search_code (via github-mcp-server) — 4221ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"class DelayedResponse repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobBatchClientTests.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs","sha":"ca94cdccdcbf1b3597a063d38ccbc10b20bea029","text_matches":[{"fragment":"            // Act\n            await blobBatchClient.SubmitBatchAsync(batch, throwOnAnyFailure: true);\n        }","matches":[{"indices":[77,94],"text":"throwOnAnyFailure"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"},{"fragment":"            AggregateException exes = Assert.ThrowsAsync<AggregateException>(\n                async () => await client.SubmitBatchAsync(batch, throwOnAnyFailure: true));\n","matches":[{"indices":[143,160],"text":"throwOnAnyFailure"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/tests/BlobBatchClientTests.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]}],"total_count":1}
```

### ✅ 17. github-mcp-server-search_code (via github-mcp-server) — 3798ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"throwOnAnyFailure repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch/tests"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"DelayedResponse.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/DelayedResponse.cs","sha":"b04f4e6c2ccf80bd9a370b7c036c93ca6d35c6fb","text_matches":[{"fragment":"    /// A Response that can be returned before a request is actually sent and\n    /// will throw until a live response is provided to wrap.\n    /// </summary>\n    internal class DelayedResponse : Response\n    {\n        /// <summary>\n        /// The live Response to wrap.","matches":[{"indices":[172,177],"text":"class"},{"indices":[178,193],"text":"DelayedResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/DelayedResponse.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"BatchConstants.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BatchConstants.cs","sha":"58335c28cdafc5bafaeb9e1041fe8b100e78d5b5","text_matches":[{"fragment":"    /// </summary>\n    internal static class BatchConstants\n    {","matches":[{"indices":[39,44],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/BatchConstants.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"},{"fragment":"#pragma warning disable CA1802 // Use literals where appropriate\n        public static readonly string DelayedResponsePropertyName = $\"{nameof(BlobBatchClient)}.{nameof(BlobBatchClient.SubmitBatch)}:DelayedResponse\";\n#pragma warning restore CA1802 // Use literals where appropriate","matches":[{"indices":[103,118],"text":"DelayedResponse"},{"indices":[199,214],"text":"DelayedResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/BatchConstants.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"BlobBatch.cs","path":"sdk/stora
... (truncated)
```

### ✅ 18. web_fetch — 3415ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/rest/api/storageservices/blob-batch"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch (redirected from https://learn.microsoft.com/rest/api/storageservices/blob-batch):
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
platformI
... (truncated)
```

### ❌ 19. github-mcp-server-get_file_contents (via github-mcp-server) — 1509ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/DelayedResponse.cs","repo":"azure-sdk-for-net","sha":"b04f4e6c2ccf80bd9a370b7c036c93ca6d35c6fb"}
```

**Error:**

```
MCP server 'github-mcp-server': failed to get git tree: Invalid object requested. SHA must identify a commit or a tree.
```

### ✅ 20. github-mcp-server-get_file_contents (via github-mcp-server) — 1929ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/DelayedResponse.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: b04f4e6c2ccf80bd9a370b7c036c93ca6d35c6fb)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Azure.Core;

namespace Azure.Storage.Blobs.Specialized
{
    /// <summary>
    /// A Response that can be returned before a request is actually sent and
    /// will throw until a live response is provided to wrap.
    /// </summary>
    internal class DelayedResponse : Response
    {
        /// <summary>
        /// The live Response to wrap.
        /// </summary>
        private Response _live;

        /// <summary>
        /// An optional function that can be used to process the response.  It
        /// is responsible for parsing response bodies and throwing exceptions.
        /// (This would be more meaningful if we add a DelayedResponse{T} down
        /// the road.)
        ///
        /// This is intended to be one of the
        /// BlobRestClient.Group.OperationName_CreateResponse methods which
        /// correctly throw when necessary.
        /// </summary>
        private readonly Func<Response, Response> _processResponse;

        /// <summary>
        /// Gets the live Response or throws an InvalidOperationException if
        /// you attempt to use the Response before the batch operation has been
        /// submitted.
        /// </summary>
        private Response LiveResponse
        {
            get => _live ?? throw BatchErrors.UseDelayedResponseEarly();
        }

        /// <summary>
        /// Creates a new instance of the <see cref="DelayedResponse"/> class.
        /// </summary>
        /// <param name="message">The message this is a response for.</param>
        /// <param name="processResponse">
        /// An optional function that can be used to process the response.  It
        /// is responsible for parsing response bodies and throwing exceptions.
    
... (truncated)
```

## Copilot Response

I’m checking the current .NET batch API and service limits so the sample handles per-subrequest results correctly.

The service only allows 256 subrequests, so 500 deletes must be split into two HTTP batch requests; I’m confirming the exact response-inspection pattern before showing the implementation.

You cannot delete 500 blobs in one HTTP batch: Azure Blob Batch allows **at most 256 subrequests**. Split the deletes into two requests—256 and 244. The 200 tier changes fit in one request.

### Packages

```powershell
dotnet add package Azure.Storage.Blobs.Batch --version 12.26.0
dotnet add package Azure.Identity
```

### Batch implementation with per-blob results

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

public sealed record BlobBatchResult(
    Uri BlobUri,
    bool Succeeded,
    int Status,
    string? ErrorCode,
    string ReasonPhrase);

public static class BlobBatchOperations
{
    private const int MaximumOperationsPerBatch = 256;

    public static async Task<IReadOnlyList<BlobBatchResult>> DeleteBlobsAsync(
        BlobBatchClient batchClient,
        IEnumerable<Uri> blobUris,
        CancellationToken cancellationToken = default)
    {
        var results = new List<BlobBatchResult>();

        // 500 blobs become two HTTP batch requests: 256 + 244.
        foreach (Uri[] chunk in blobUris.Chunk(MaximumOperationsPerBatch))
        {
            using BlobBatch batch = batchClient.CreateBatch();

            // The returned Response objects are populated when the batch is submitted.
            var operations = chunk
                .Select(uri => (
                    Uri: uri,
                    Response: batch.DeleteBlob(
                        uri,
                        DeleteSnapshotsOption.IncludeSnapshots)))
                .ToArray();

            await SubmitAndCollectAsync(
                batchClient,
                batch,
                operations,
                results,
                cancellationToken);
        }

        return results;
    }

    public static async Task<IReadOnlyList<BlobBatchResult>> SetCoolTierAsync(
        BlobBatchClient batchClient,
        IEnumerable<Uri> blobUris,
        CancellationToken cancellationToken = default)
    {
        var results = new List<BlobBatchResult>();

        foreach (Uri[] chunk in blobUris.Chunk(MaximumOperationsPerBatch))
        {
            using BlobBatch batch = batchClient.CreateBatch();

            var operations = chunk
                .Select(uri => (
                    Uri: uri,
                    Response: batch.SetBlobAccessTier(uri, AccessTier.Cool)))
                .ToArray();

            await SubmitAndCollectAsync(
                batchClient,
                batch,
                operations,
                results,
                cancellationToken);
        }

        return results;
    }

    private static async Task SubmitAndCollectAsync(
        BlobBatchClient batchClient,
        BlobBatch batch,
        IEnumerable<(Uri Uri, Response Response)> operations,
        ICollection<BlobBatchResult> results,
        CancellationToken cancellationToken)
    {
        try
        {
            // false means a failed subrequest does not abort result processing.
            // A failure of the outer HTTP batch request still throws
            // RequestFailedException.
            await batchClient.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);
        }
        catch (RequestFailedException ex)
        {
            // The outer request failed, so individual operation outcomes might
            // not be available. Do not report these items as successful.
            throw new InvalidOperationException(
                $"The batch HTTP request failed with status {ex.Status} " +
                $"and error code '{ex.ErrorCode}'.",
                ex);
        }

        foreach ((Uri uri, Response response) in operations)
        {
            int status = response.Status;
            bool succeeded = status is >= 200 and < 300;

            string? errorCode = null;
            if (!succeeded &&
                response.Headers.TryGetValue("x-ms-error-code", out string value))
            {
                errorCode = value;
            }

            results.Add(new BlobBatchResult(
                uri,
                succeeded,
                status,
                errorCode,
                response.ReasonPhrase));
        }
    }
}
```

### Usage

```csharp
string accountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME")
    ?? throw new InvalidOperationException(
        "AZURE_STORAGE_ACCOUNT_NAME is required.");

Uri serviceUri = new($"https://{accountName}.blob.core.windows.net/");

// Use ManagedIdentityCredential in Azure-hosted production applications.
// DefaultAzureCredential is convenient for local development.
TokenCredential credential =
    Environment.GetEnvironmentVariable("DOTNET_ENVIRONMENT") == "Development"
        ? new DefaultAzureCredential()
        : new ManagedIdentityCredential();

var serviceClient = new BlobServiceClient(serviceUri, credential);
BlobBatchClient batchClient = serviceClient.GetBlobBatchClient();
BlobContainerClient containerClient =
    serviceClient.GetBlobContainerClient("my-container");

// Example inputs. Listing can also be done asynchronously with
// containerClient.GetBlobsAsync().
Uri[] deleteUris = Enumerable.Range(1, 500)
    .Select(i => containerClient.GetBlobClient($"delete/blob-{i}.dat").Uri)
    .ToArray();

Uri[] tierUris = Enumerable.Range(1, 200)
    .Select(i => containerClient.GetBlobClient($"archive/blob-{i}.dat").Uri)
    .ToArray();

IReadOnlyList<BlobBatchResult> deleteResults =
    await BlobBatchOperations.DeleteBlobsAsync(batchClient, deleteUris);

IReadOnlyList<BlobBatchResult> tierResults =
    await BlobBatchOperations.SetCoolTierAsync(batchClient, tierUris);

foreach (BlobBatchResult failure in
         deleteResults.Concat(tierResults).Where(r => !r.Succeeded))
{
    Console.Error.WriteLine(
        $"{failure.BlobUri}: HTTP {failure.Status} " +
        $"{failure.ErrorCode} - {failure.ReasonPhrase}");
}

Console.WriteLine(
    $"Deletes: {deleteResults.Count(r => r.Succeeded)} succeeded, " +
    $"{deleteResults.Count(r => !r.Succeeded)} failed.");

Console.WriteLine(
    $"Tier changes: {tierResults.Count(r => r.Succeeded)} succeeded, " +
    $"{tierResults.Count(r => !r.Succeeded)} failed.");
```

### Partial-failure behavior

`SubmitBatchAsync(..., throwOnAnyFailure: false)` is the best option when you need to associate each result with its blob:

- The outer `Response` indicates whether the multipart batch request was accepted.
- Each `Response` returned by `DeleteBlob` or `SetBlobAccessTier` is populated after submission.
- A failed subrequest does not roll back successful subrequests.
- An outer-request failure still throws `RequestFailedException`.
- With `throwOnAnyFailure: true`, failed subrequests are combined into an `AggregateException`. This is convenient when any failure should throw, but less convenient for producing a complete per-blob report.
- The higher-level `DeleteBlobsAsync` and `SetBlobsAccessTierAsync` methods also throw aggregated subrequest failures. Fine-grained `BlobBatch` construction provides better result-to-blob correlation.

### Service limits

| Restriction | Limit |
|---|---:|
| Subrequests per batch | **256** |
| Multipart batch request body | **4 MiB** |
| Supported operations | Delete Blob, Set Blob Tier |
| Operation types per batch | One homogeneous type |
| Atomicity | None; every subrequest executes independently |
| Execution order | Not guaranteed |
| Resource scope | Same storage account |
| Container-scoped client | All blobs must be in that container |

The 4-MiB restriction applies to the serialized multipart request—URLs, boundaries, and headers—not the blob contents. If unusually long blob names or large conditional headers approach that limit, use chunks smaller than 256.

References: [Blob Batch REST API](https://learn.microsoft.com/rest/api/storageservices/blob-batch), [Azure.Storage.Blobs.Batch README and samples](https://github.com/Azure/azure-sdk-for-net/tree/main/sdk/storage/Azure.Storage.Blobs.Batch), [BlobBatchClient API](https://learn.microsoft.com/dotnet/api/azure.storage.blobs.specialized.blobbatchclient).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 8 criteria passed. The response provides a strong custom-batch implementation with correct limits and partial-result collection, but omits actual use of both high-level bulk methods and does not show per-subrequest RequestFailedException handling through AggregateException.

### Strengths

- Correctly splits 500 deletes around the 256-operation limit.
- Uses CreateBatch() and SubmitBatchAsync() to correlate each blob with its response.
- Accurately explains non-atomic partial success and the 4 MiB multipart request-body restriction.
- Uses token-based Azure authentication and observes the same-account batching scope.

### Issues

- No direct BlobBatchClient.DeleteBlobsAsync() example.
- No direct BlobBatchClient.SetBlobsAccessTierAsync() example.
- No AggregateException catch that inspects individual RequestFailedException status and error details.

## Grader Results

- batch-blob-operations.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/8)
      - `BlobBatchClient` from `Azure.Storage.Blobs.Batch` package: Pass
      - `BlobBatchClient.DeleteBlobsAsync()` for bulk delete: Fail
      - `BlobBatchClient.SetBlobsAccessTierAsync()` for bulk tier changes: Fail
      - Custom batch via `BlobBatchClient.CreateBatch()` and `SubmitBatchAsync()`: Pass
      - Batch size limits (256 operations per batch): Pass
      - Partial failure handling: `AggregateException` with per-operation status: Pass
      - `RequestFailedException` for individual operation failures within a batch: Fail
      - Authentication scopes for batch operations: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 62% | 1.00 | 0.6250 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6250** | **62.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-batch --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
