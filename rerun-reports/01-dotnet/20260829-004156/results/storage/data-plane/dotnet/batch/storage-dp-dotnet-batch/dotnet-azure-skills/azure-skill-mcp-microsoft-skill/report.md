# Evaluation Report: storage-dp-dotnet-batch

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 201.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-batch` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 5/8 |
| Duration | 201.0s |
| Timestamp | 2026-08-28T18:46:05Z |
| Files Generated | 0 |
| Event Count | 4289 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 183.7s |
| Review | 16.2s |
| **Total** | **201.0s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=612174 out=6874 |
| Turn Count | 11 |

## Tool Calls

### ✅ 1. skill — 738ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 5130ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Generate modern .NET code using Azure.Storage.Blobs.Batch BlobBatchClient for bulk delete and access-tier updates, including authentication and partial batch failure handling.","parameters":{"language":"C#","operation":"code-generation","resource":"Azure Blob Storage"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 6436ms

**Input:**

```json
{"command":"search","intent":"Find official Microsoft documentation for Azure.Storage.Blobs.Batch BlobBatchClient .NET, SubmitBatchAsync, DeleteBlobsAsync, SetBlobsAccessTierAsync, partial failures, and batch operation/size limits.","parameters":{"query":"Azure.Storage.Blobs.Batch BlobBatchClient .NET SubmitBatchAsync partial failures 256 4 MiB limits"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. view — 5968ms

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

### ✅ 5. view — 5525ms

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

### ✅ 6. glob — 4723ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2578ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a .NET Blob Storage batch operations example.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3396ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Microsoft documentation for .NET BlobBatchClient batch limits and partial failure behavior.","parameters":{"query":"Azure.Storage.Blobs.Batch BlobBatchClient .NET SubmitBatchAsync partial failures RequestFailedException 256 subrequests 4 MiB batch limits"}}
```

**Output:**

```
{"results":[{"title":"Blob Batch","content":"# Blob Batch\n## Remarks\nOne of the main benefits of using a batch request is the reduction in the number of connections a client has to open. Note the following restrictions:\n1. Supported subrequests in the batch are \u0060Set Blob Tier\u0060 (for block blobs) and \u0060Delete Blob\u0060.\n2. Only supports up to 256 subrequests in a single batch. The size of the body for a batch request can\u0027t exceed 4 MB.\n3. An empty batch request fails with code 400 (Bad Request).\n4. There are no guarantees on the order of execution of the batch subrequests.\n5. Batch subrequest execution isn\u0027t atomic. Each subrequest runs independently.\n6. Each subrequest must be for a resource within the same storage account. A single batch request doesn\u0027t support running requests from different storage accounts.\n7. A nested request body isn\u0027t supported.\n8. If the server fails to parse the request body, the entire batch fails, and no request will be run.\n9. Note that [Account SAS](https://learn.microsoft.com/rest/api/storageservices/create-account-sas) is the only shared access signature type supported by \u0060Blob Batch\u0060, when the batch isn\u0027t using \u0060restype=container\u0060.\n### Scope all subrequests to a specific container\nBeginning with REST version 2020-04-08, the \u0060Blob Batch\u0060 API supports scoping subrequests to a specified container. When the request URI includes the container name and the \u0060restype=container\u0060 parameter, then each subrequest must apply to the same container. If the container name specified for a subrequest doesn\u0027t match the container name provided in URI, the service returns error code 400 (Bad Request).\nAll authorization mechanisms supported for a container are valid for a \u0060Blob Batch\u0060 operation that is scoped to the container. Each subrequest sends an authorization header to the service.","contentUrl":"https://learn.microsoft.com/rest/api/storageser
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 2971ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# code samples for BlobBatchClient delete and set access tier batch operations.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs.Batch BlobBatchClient DeleteBlobsAsync SetBlobsAccessTierAsync SubmitBatchAsync partial failure"}}
```

**Output:**

```
{"results":[{"description":"description: Illustrates fine-grained control over batch operations by creating a batch with specific delete options for blobs and submitting it using the Azure Storage Blobs Batch client library for .NET.\npackage: Azure.Storage.Blobs.Batch\nlanguage: c\n","codeSnippet":"// Get a connection string to our Azure Storage account.\r\nstring connectionString = \u0022\u003Cconnection_string\u003E\u0022;\r\nstring containerName = \u0022sample-container\u0022;\r\n\r\n// Get a reference to a container named \u0022sample-container\u0022 and then create it\r\nBlobServiceClient service = new BlobServiceClient(connectionString);\r\nBlobContainerClient container = service.GetBlobContainerClient(containerName);\r\ncontainer.Create();\r\n\r\n// Create three blobs named \u0022foo\u0022, \u0022bar\u0022, and \u0022baz\u0022\r\nBlobClient foo = container.GetBlobClient(\u0022foo\u0022);\r\nBlobClient bar = container.GetBlobClient(\u0022bar\u0022);\r\nBlobClient baz = container.GetBlobClient(\u0022baz\u0022);\r\nfoo.Upload(BinaryData.FromString(\u0022Foo!\u0022));\r\nfoo.CreateSnapshot();\r\nbar.Upload(BinaryData.FromString(\u0022Bar!\u0022));\r\nbar.CreateSnapshot();\r\nbaz.Upload(BinaryData.FromString(\u0022Baz!\u0022));\r\n\r\n// Create a batch with three deletes\r\nBlobBatchClient batchClient = service.GetBlobBatchClient();\r\nBlobBatch batch = batchClient.CreateBatch();\r\nbatch.DeleteBlob(foo.Uri, DeleteSnapshotsOption.IncludeSnapshots);\r\nbatch.DeleteBlob(bar.Uri, DeleteSnapshotsOption.OnlySnapshots);\r\nbatch.DeleteBlob(baz.Uri);\r\n\r\n// Submit the batch\r\nbatchClient.SubmitBatch(batch);","link":"https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Imports namespaces for Azure Storage Blobs and related models to enable blob storage operations.\npackage: Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"using Azure.Storage.Blobs;\r\nusing Azu
... (truncated)
```

### ✅ 10. rg — 2328ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787942802235-copilot-tool-output-bee51b4dc952495ab5e90ec56a619fa3.txt","pattern":"Only supports up to 256|4 MB|RequestFailedException|partial|throwOnAnyFailure|SubmitBatchAsync|DeleteBlobsAsync|SetBlobsAccessTierAsync|https://learn.microsoft.com"}
```

**Output:**

```
Output too large to read at once (47.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787942808361-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787942802235-copilot-tool-output-bee51b4dc952495ab5e90ec56a619fa3.txt:{"results":[{"title":"Blob Batch","content":"# Blob Batch\n## Remarks\nOne of the main benefits of using a batch request is the reduction in the number of connections a client has to open. Note the following restrictions:\n1. Supported subrequests in the batch are \u0060Set Blob Tier\u0060 (for block blobs) and \u0060Delete Blob\u0060.\n2. Only supports up to 256 subrequests in a single batc
```

### ✅ 11. web_fetch — 3133ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/rest/api/storageservices/blob-batch"}
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

### ✅ 12. web_fetch — 2730ms

**Input:**

```json
{"max_length":16000,"url":"https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet"}
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

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server) — 8181ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs","ref":"refs/tags/Azure.Storage.Blobs.Batch_12.26.0","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 86b9f524dc1b8dfd11a2266c53cad0a7550f256f)// Copyright (c) Microsoft Corporation. All rights reserved.
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

#pragma warning disable SA1402  // File may only contain a single type

namespace Azure.Storage.Blobs.Specialized
{
    /// <summary>
    /// The <see cref="BlobBatchClient"/> allows you to batch multiple Azure
    /// Storage operations in a single request.
    /// </summary>
    public class BlobBatchClient
    {
        /// <summary>
        /// Gets the blob service's primary <see cref="Uri"/> endpoint.
        /// </summary>
        private readonly Uri _uri;

        /// <summary>
        /// Gets the blob service's primary <see cref="Uri"/> endpoint.
        /// </summary>
        public virtual Uri Uri => _uri;

        /// <summary>
        /// If this BlobBatchClient is scoped to a container.
        /// </summary>
        private readonly bool _isContainerScoped;

        /// <summary>
        /// If this BlobBatchClient is scoped to a container.
        /// </summary>
        internal virtual bool IsContainerScoped => _isContainerScoped;

        /// <summary>
        /// The <see cref="HttpPipeline"/> transport pipeline used to send
        /// every request.
        /// </summary>
        private readonly HttpPipeline _pipeline;

        /// <summary>
        /// The <see cref="HttpPipeline"/> transport pipeline used to send
        /// every request.
        /// </summary>
        internal virtual HttpPipeline Pipeline => _pipeline;

        /// <summary>
        /// The version of the service to use when sending requests.
        /// </summary>
        private readonly BlobClientOptions.ServiceVersion _version;

        /// <summary>
        /// Th
... (truncated)
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server) — 7069ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatch.cs","ref":"refs/tags/Azure.Storage.Blobs.Batch_12.26.0","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: fcc154a289140f2fc41499453ed9bd32a2716350)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.ComponentModel;
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
        /// The list of messages that will be sent as part of this batch.
        /// </summary>
        private readonly IList<HttpMessage> _messages = new List<HttpMessage>();

        /// <summa
... (truncated)
```

### ✅ 15. azure-documentation (via azure) — 5858ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the current official BlobBatchClient API reference, including async overload signatures and throwOnAnyFailure behavior.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient?view=azure-dotnet"}}
```

**Output:**

```
# BlobBatchClient Class

## Definition

- Namespace:
    - [Azure.Storage.Blobs.Specialized](azure.storage.blobs.specialized?view=azure-dotnet)

- Assembly:
    - Azure.Storage.Blobs.Batch.dll

- Package:
    - Azure.Storage.Blobs.Batch v12.26.0

- Package:
    - Azure.Storage.Blobs.Batch v12.27.0-beta.1

- Source:
    - [BlobBatchClient.cs](https://github.com/Azure/azure-sdk-for-net/blob/8608dd90245dcc1e8d9a30f4fdb6494d8d268d82/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs)

- Source:
    - [BlobBatchClient.cs](https://github.com/Azure/azure-sdk-for-net/blob/62cf4bdea6158c44231a76da938a736d91eb11d6/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs)

The [BlobBatchClient](azure.storage.blobs.specialized.blobbatchclient?view=azure-dotnet) allows you to batch multiple Azure Storage operations in a single request.

```csharp
public class BlobBatchClient
```

```fsharp
type BlobBatchClient = class
```

```vb
Public Class BlobBatchClient
```

- Inheritance
    - [Object](/en-us/dotnet/api/system.object)
BlobBatchClient

## Constructors

| Name | Description |
| --- | --- |
| [BlobBatchClient()](azure.storage.blobs.specialized.blobbatchclient.-ctor?view=azure-dotnet#azure-storage-blobs-specialized-blobbatchclient-ctor) | Initializes a new instance of the [BlobBatchClient](azure.storage.blobs.specialized.blobbatchclient?view=azure-dotnet) class for mocking. |
| [BlobBatchClient(BlobContainerClient)](azure.storage.blobs.specialized.blobbatchclient.-ctor?view=azure-dotnet#azure-storage-blobs-specialized-blobbatchclient-ctor%28azure-storage-blobs-blobcontainerclient%29) | Initializes a new instance of the [BlobBatchClient](azure.storage.blobs.specialized.blobbatchclient?view=azure-dotnet) class for container associated with the [BlobContainerClient](azure.storage.blobs.blobcontainerclient?view=azure-dotnet). The new [BlobBatchClient](azure.storage.blobs.specialized.blobbatchclient?view=azure-dotnet) uses t
... (truncated)
```

### ✅ 16. rg — 2068ms

**Input:**

```json
{"-A":18,"-B":5,"-n":true,"head_limit":160,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787942825235-copilot-tool-output-84c342643af048caa7f279407c1f39cd.txt","pattern":"public virtual async Task\u003cResponse|public virtual Response\\[\\]|DeleteBlobsAsync\\(|SetBlobsAccessTierAsync\\(|SubmitBatchAsync\\(|throwOnAnyFailure"}
```

**Output:**

```
[grep content: 338 matches across 1 file(s) under C:/Users/weidxu/AppData/Local/Temp]

1787942825235-copilot-tool-output-84c342643af048caa7f279407c1f39cd.txt (338 match(es)):
  294-         /// Submit a <see cref="BlobBatch"/> of sub-operations.
  295-         /// </summary>
  296-         /// <param name="batch">
  297-         /// A <see cref="BlobBatch"/> of sub-operations.
  298-         /// </param>
  299:         /// <param name="throwOnAnyFailure">
  300-         /// A value indicating whether or not to throw exceptions for
  301-         /// sub-operation failures.
  302-         /// </param>
  303-         /// <param name="cancellationToken">
  304-         /// Optional <see cref="CancellationToken"/> to propagate notifications
  305-         /// that the operation should be cancelled.
  306-         /// </param>
  307-         /// <returns>
  308-         /// A <see cref="Response"/> on successfully submitting.
  309-         /// </returns>
  310-         /// <remarks>
  311-         /// A <see cref="RequestFailedException"/> will be thrown if
  312-         /// a failure to submit the batch occurs.  Individual sub-operation
  313:         /// failures will only throw if <paramref name="throwOnAnyFailure"/> is
  314-         /// true and be wrapped in an <see cref="AggregateException"/>.
  315-         /// </remarks>
  316-         [ForwardsClientCalls]
  317-         public virtual Response SubmitBatch(
  318-             BlobBatch batch,
  319:             bool throwOnAnyFailure = false,
  320-             CancellationToken cancellationToken = default) =>
  321-             SubmitBatchInternal(
  322-                 batch,
  323:                 throwOnAnyFailure,
  324-                 async: false,
  325-                 cancellationToken)
  326-                 .EnsureCompleted();
  327- 
  328-         /// <summary>
  329-         /// Submit a <see cref="BlobBatch"/> of sub-operations.
  330-         /// </summary>
  331-         /// <param name="ba
... (truncated)
```

### ✅ 17. github-mcp-server-search_code (via github-mcp-server) — 5132ms

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":20,"query":"GetBlobBatchClient repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch/src"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"BlobBatchClient.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs","text_matches":[{"fragment":"        /// <returns>A new <see cref=\"BlobBatchClient\"/> instance.</returns>\n        public static BlobBatchClient GetBlobBatchClient(this BlobServiceClient client)\n            => new BlobBatchClient(client);","matches":[{"indices":[115,133],"text":"GetBlobBatchClient"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"},{"fragment":"        /// <returns>A new <see cref=\"BlobBatchClient\"/> instance.</returns>\n        public static BlobBatchClient GetBlobBatchClient(this BlobContainerClient client)\n            => new BlobBatchClient(client);","matches":[{"indices":[115,133],"text":"GetBlobBatchClient"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]}],"total_count":1}
```

### ✅ 18. azure-documentation (via azure) — 4544ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official BlobBatch SubmitBatchAsync parameter behavior, especially throwOnAnyFailure false and delayed operation Responses.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatchclient.submitbatchasync?view=azure-dotnet"}}
```

**Output:**

```
# BlobBatch.DeleteBlob Method

## Definition

- Namespace:
    - [Azure.Storage.Blobs.Specialized](azure.storage.blobs.specialized?view=azure-dotnet)

- Assembly:
    - Azure.Storage.Blobs.Batch.dll

- Package:
    - Azure.Storage.Blobs.Batch v12.26.0

- Package:
    - Azure.Storage.Blobs.Batch v12.27.0-beta.1

## Overloads

| Name | Description |
| --- | --- |
| [DeleteBlob(String, String, BatchDeleteBlobOptions)](azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet#azure-storage-blobs-specialized-blobbatch-deleteblob%28system-string-system-string-azure-storage-blobs-batch-batchdeletebloboptions%29) | The [DeleteBlob(String, String, BatchDeleteBlobOptions)](azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet#azure-storage-blobs-specialized-blobbatch-deleteblob%28system-string-system-string-azure-storage-blobs-batch-batchdeletebloboptions%29) operation marks the specified blob or snapshot for deletion.<br><br>Note that in order to delete a blob, you must delete all of its snapshots. You can delete both at the same time using [IncludeSnapshots](azure.storage.blobs.models.deletesnapshotsoption?view=azure-dotnet#azure-storage-blobs-models-deletesnapshotsoption-includesnapshots) in `options`.<br><br>For more information, see [Delete Blob](/en-us/rest/api/storageservices/delete-blob). |
| [DeleteBlob(Uri, DeleteSnapshotsOption, BlobRequestConditions)](azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet#azure-storage-blobs-specialized-blobbatch-deleteblob%28system-uri-azure-storage-blobs-models-deletesnapshotsoption-azure-storage-blobs-models-blobrequestconditions%29) | The [DeleteBlob(Uri, DeleteSnapshotsOption, BlobRequestConditions)](azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet#azure-storage-blobs-specialized-blobbatch-deleteblob%28system-uri-azure-storage-blobs-models-deletesnapshotsoption-azure-storage-blobs-models-blobrequestconditions%29) operation marks the speci
... (truncated)
```

### ✅ 19. azure-documentation (via azure) — 3916ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official BlobBatch DeleteBlob return response semantics for per-operation inspection after submitting.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.storage.blobs.specialized.blobbatch.deleteblob?view=azure-dotnet"}}
```

**Output:**

```
# BlobBatchClient.SubmitBatchAsync Method

## Definition

- Namespace:
    - [Azure.Storage.Blobs.Specialized](azure.storage.blobs.specialized?view=azure-dotnet)

- Assembly:
    - Azure.Storage.Blobs.Batch.dll

- Package:
    - Azure.Storage.Blobs.Batch v12.26.0

- Package:
    - Azure.Storage.Blobs.Batch v12.27.0-beta.1

- Source:
    - [BlobBatchClient.cs](https://github.com/Azure/azure-sdk-for-net/blob/8608dd90245dcc1e8d9a30f4fdb6494d8d268d82/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs)

- Source:
    - [BlobBatchClient.cs](https://github.com/Azure/azure-sdk-for-net/blob/62cf4bdea6158c44231a76da938a736d91eb11d6/sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatchClient.cs)

::: moniker range=" azure-dotnet azure-dotnet-preview "

Submit a [BlobBatch](azure.storage.blobs.specialized.blobbatch?view=azure-dotnet) of sub-operations.

```csharp
public virtual System.Threading.Tasks.Task<Azure.Response> SubmitBatchAsync(Azure.Storage.Blobs.Specialized.BlobBatch batch, bool throwOnAnyFailure = false, System.Threading.CancellationToken cancellationToken = default);
```

```fsharp
abstract member SubmitBatchAsync : Azure.Storage.Blobs.Specialized.BlobBatch * bool * System.Threading.CancellationToken -> System.Threading.Tasks.Task<Azure.Response>
override this.SubmitBatchAsync : Azure.Storage.Blobs.Specialized.BlobBatch * bool * System.Threading.CancellationToken -> System.Threading.Tasks.Task<Azure.Response>
```

```vb
Public Overridable Function SubmitBatchAsync (batch As BlobBatch, Optional throwOnAnyFailure As Boolean = false, Optional cancellationToken As CancellationToken = Nothing) As Task(Of Response)
```

#### Parameters

- batch
    - [BlobBatch](azure.storage.blobs.specialized.blobbatch?view=azure-dotnet)

A [BlobBatch](azure.storage.blobs.specialized.blobbatch?view=azure-dotnet) of sub-operations.

- throwOnAnyFailure
    - [Boolean](/en-us/dotnet/api/system.boolean)

A value indicating whether o
... (truncated)
```

### ✅ 20. rg — 907ms

**Input:**

```json
{"-A":35,"-B":5,"-n":true,"head_limit":120,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787942825235-copilot-tool-output-84c342643af048caa7f279407c1f39cd.txt","pattern":"class DelayedResponse|SetLiveResponse|public override int Status|throwOnFailure"}
```

**Output:**

```
[grep content: 41 matches across 1 file(s) under C:/Users/weidxu/AppData/Local/Temp]

1787942825235-copilot-tool-output-84c342643af048caa7f279407c1f39cd.txt (41 match(es)):
  598-                 {
  599-                     if (messages[i].TryGetProperty(BatchConstants.DelayedResponsePropertyName, out object value) &&
  600-                         value is DelayedResponse response)
  601-                     {
  602- #pragma warning disable AZC0110 // DO NOT use await keyword in possibly synchronous scope.
  603:                         response.SetLiveResponse(responses[i], throwOnAnyFailure);
  604- #pragma warning restore AZC0110 // DO NOT use await keyword in possibly synchronous scope.
  605-                     }
  606-                 }
  607-                 catch (Exception ex)
  608-                 {
  609-                     failures.Add(ex);
  610-                 }
  611-             }
  612- 
  613-             // Throw any failures
  614-             if (failures.Count > 0)
  615-             {
  616-                 throw BatchErrors.ResponseFailures(failures);
  617-             }
  618-         }
  619-         #endregion Create/SubmitBatch
  620- 
  621-         #region DeleteBlobs
  622-         /// <summary>
  623-         /// The DeleteBlobs operation marks the specified blobs for deletion.
  624-         /// All of the deletions are sent as a single batched request.
  625-         /// </summary>
  626-         /// <param name="cancellationToken">
  627-         /// Optional <see cref="CancellationToken"/> to propagate notifications
  628-         /// that the operation should be cancelled.
  629-         /// </param>
  630-         /// <param name="blobUris">URIs of the blobs to delete.</param>
  631-         /// <param name="snapshotsOption">
  632-         /// Specifies options for deleting blob snapshots.
  633-         /// </param>
  634-         /// <returns>
  635-         /// The <see cref="Response"/>s for the individual Delete ope
... (truncated)
```

### ✅ 21. github-mcp-server-search_code (via github-mcp-server) — 2179ms

**Input:**

```json
{"fields":["name","path","text_matches"],"perPage":10,"query":"class DelayedResponse repo:Azure/azure-sdk-for-net path:sdk/storage/Azure.Storage.Blobs.Batch/src"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"DelayedResponse.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/DelayedResponse.cs","text_matches":[{"fragment":"    /// A Response that can be returned before a request is actually sent and\n    /// will throw until a live response is provided to wrap.\n    /// </summary>\n    internal class DelayedResponse : Response\n    {\n        /// <summary>\n        /// The live Response to wrap.","matches":[{"indices":[172,177],"text":"class"},{"indices":[178,193],"text":"DelayedResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/DelayedResponse.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"BatchConstants.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BatchConstants.cs","text_matches":[{"fragment":"    /// </summary>\n    internal static class BatchConstants\n    {","matches":[{"indices":[39,44],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/BatchConstants.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"},{"fragment":"#pragma warning disable CA1802 // Use literals where appropriate\n        public static readonly string DelayedResponsePropertyName = $\"{nameof(BlobBatchClient)}.{nameof(BlobBatchClient.SubmitBatch)}:DelayedResponse\";\n#pragma warning restore CA1802 // Use literals where appropriate","matches":[{"indices":[103,118],"text":"DelayedResponse"},{"indices":[199,214],"text":"DelayedResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Blobs.Batch/src/BatchConstants.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"BlobBatch.cs","path":"sdk/storage/Azure.Storage.Blobs.Batch/src/BlobBatch.cs","text_matches":[{"fragment":"        /// <summary>\
... (truncated)
```

### ❌ 22. web_fetch — 1929ms

**Input:**

```json
{"max_length":4000,"url":"https://api.nuget.org/v3-flatcontainer/azure.storage.blobs.batch/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

## Copilot Response

I’m grounding the sample in the current .NET batch API and service limits so the partial-failure handling matches the SDK’s response model.

The key constraint changes the delete plan: 500 deletes cannot fit in one Blob Batch request; they must be split into at least two batches (256 + 244). I’m checking the exact current overloads so the sample compiles as written.

You **cannot delete 500 blobs in one Blob Batch HTTP request**. Azure Blob Storage allows at most **256 subrequests**, so 500 deletes require at least two requests: **256 + 244**. The 200 tier changes fit in one request.

```powershell
dotnet add package Azure.Storage.Blobs.Batch
dotnet add package Azure.Identity
```

The following helper splits operations into valid batches, submits them with `SubmitBatchAsync`, and correlates each subresponse with its blob URI. Setting `throwOnAnyFailure: false` is important: it prevents one failed subrequest from hiding the successful results.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

public sealed record BlobBatchFailure(
    Uri BlobUri,
    int Status,
    string? ErrorCode,
    string Reason);

public sealed record BlobBatchResult(
    int Succeeded,
    IReadOnlyList<BlobBatchFailure> Failures);

public static class BlobBatchOperations
{
    private const int MaxOperationsPerBatch = 256;

    public static Task<BlobBatchResult> DeleteAsync(
        BlobBatchClient batchClient,
        IReadOnlyList<Uri> blobUris,
        CancellationToken cancellationToken = default)
    {
        return SubmitInChunksAsync(
            batchClient,
            blobUris,
            static (batch, uri) =>
                batch.DeleteBlob(
                    uri,
                    snapshotsOption: DeleteSnapshotsOption.IncludeSnapshots),
            cancellationToken);
    }

    public static Task<BlobBatchResult> SetCoolTierAsync(
        BlobBatchClient batchClient,
        IReadOnlyList<Uri> blobUris,
        CancellationToken cancellationToken = default)
    {
        return SubmitInChunksAsync(
            batchClient,
            blobUris,
            static (batch, uri) =>
                batch.SetBlobAccessTier(uri, AccessTier.Cool),
            cancellationToken);
    }

    private static async Task<BlobBatchResult> SubmitInChunksAsync(
        BlobBatchClient batchClient,
        IReadOnlyList<Uri> blobUris,
        Func<BlobBatch, Uri, Response> addOperation,
        CancellationToken cancellationToken)
    {
        ArgumentNullException.ThrowIfNull(batchClient);
        ArgumentNullException.ThrowIfNull(blobUris);
        ArgumentNullException.ThrowIfNull(addOperation);

        int succeeded = 0;
        var failures = new List<BlobBatchFailure>();

        for (int offset = 0; offset < blobUris.Count; offset += MaxOperationsPerBatch)
        {
            int count = Math.Min(MaxOperationsPerBatch, blobUris.Count - offset);

            using BlobBatch batch = batchClient.CreateBatch();
            var operations = new List<(Uri BlobUri, Response Response)>(count);

            for (int index = offset; index < offset + count; index++)
            {
                Uri blobUri = blobUris[index];

                // The returned Response is populated after SubmitBatchAsync completes.
                Response response = addOperation(batch, blobUri);
                operations.Add((blobUri, response));
            }

            // A RequestFailedException here represents failure of the overall
            // batch request, such as authentication or malformed-request failure.
            await batchClient.SubmitBatchAsync(
                batch,
                throwOnAnyFailure: false,
                cancellationToken);

            foreach ((Uri blobUri, Response response) in operations)
            {
                if (response.Status is >= 200 and < 300)
                {
                    succeeded++;
                    continue;
                }

                string? errorCode = response.Headers.TryGetValue(
                    "x-ms-error-code",
                    out string value)
                    ? value
                    : null;

                failures.Add(new BlobBatchFailure(
                    BlobUri: blobUri,
                    Status: response.Status,
                    ErrorCode: errorCode,
                    Reason: response.ReasonPhrase));
            }
        }

        return new BlobBatchResult(succeeded, failures);
    }
}
```

### Using the helper

```csharp
bool isDevelopment =
    string.Equals(
        Environment.GetEnvironmentVariable("DOTNET_ENVIRONMENT"),
        "Development",
        StringComparison.OrdinalIgnoreCase);

TokenCredential credential = isDevelopment
    ? new DefaultAzureCredential()   // Local development
    : new ManagedIdentityCredential(); // Azure-hosted production

var serviceClient = new BlobServiceClient(
    new Uri("https://<storage-account>.blob.core.windows.net"),
    credential);

BlobContainerClient containerClient =
    serviceClient.GetBlobContainerClient("documents");

// A container-scoped client ensures every subrequest targets this container.
BlobBatchClient batchClient = containerClient.GetBlobBatchClient();

IReadOnlyList<Uri> deleteUris = deleteBlobNames
    .Select(name => containerClient.GetBlobClient(name).Uri)
    .ToArray();

IReadOnlyList<Uri> tierUris = tierBlobNames
    .Select(name => containerClient.GetBlobClient(name).Uri)
    .ToArray();

try
{
    // For 500 URIs, this sends two batch HTTP requests: 256 and 244.
    BlobBatchResult deleteResult =
        await BlobBatchOperations.DeleteAsync(batchClient, deleteUris);

    Console.WriteLine(
        $"Deleted: {deleteResult.Succeeded}; failed: {deleteResult.Failures.Count}");

    foreach (BlobBatchFailure failure in deleteResult.Failures)
    {
        Console.Error.WriteLine(
            $"Delete failed: {failure.BlobUri}; " +
            $"HTTP {failure.Status}; " +
            $"code={failure.ErrorCode}; reason={failure.Reason}");
    }

    // For 200 URIs, this sends one batch HTTP request.
    BlobBatchResult tierResult =
        await BlobBatchOperations.SetCoolTierAsync(batchClient, tierUris);

    Console.WriteLine(
        $"Tier changed: {tierResult.Succeeded}; failed: {tierResult.Failures.Count}");

    foreach (BlobBatchFailure failure in tierResult.Failures)
    {
        Console.Error.WriteLine(
            $"Tier change failed: {failure.BlobUri}; " +
            $"HTTP {failure.Status}; " +
            $"code={failure.ErrorCode}; reason={failure.Reason}");
    }
}
catch (RequestFailedException ex)
{
    // The parent batch request failed. Per-item responses might not be available.
    Console.Error.WriteLine(
        $"Batch submission failed: HTTP {ex.Status}; " +
        $"code={ex.ErrorCode}; message={ex.Message}");

    throw;
}
```

### Limits and behavior

| Constraint | Limit/behavior |
|---|---|
| Operations per batch | Maximum **256** |
| Batch request body | Maximum **4 MB**, including multipart headers and boundaries |
| Supported operations | `Delete Blob` and `Set Blob Tier` only |
| Operation types | Every subrequest in a batch must be the **same type** |
| Resource scope | Same storage account; a container-scoped client requires the same container |
| Atomicity | **Not atomic**—successful operations are not rolled back |
| Execution order | Not guaranteed |
| Tier operation | Supported for block blobs in Blob Storage or general-purpose v2 accounts |

If `throwOnAnyFailure: true` is used, failed subrequests are wrapped in an `AggregateException` containing `RequestFailedException` instances. Using `false` and inspecting each delayed `Response`, as above, provides better URI-to-failure correlation.

**References:** [Azure Storage Blobs Batch .NET client](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs.batch-readme?view=azure-dotnet) · [Blob Batch REST API and limits](https://learn.microsoft.com/rest/api/storageservices/blob-batch)

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 8 criteria passed. The response gives a strong custom-batch implementation with chunking, limits, and partial-result inspection, but omits both requested convenience methods and authentication-scope details.

### Strengths

- Correctly rejects the impossible 500-operation single request and chunks at 256.
- Demonstrates CreateBatch() and SubmitBatchAsync() with per-operation response correlation.
- Accurately describes non-atomic partial failures, AggregateException behavior, and the 4 MB request-body limit.

### Issues

- Missing DeleteBlobsAsync() example.
- Missing SetBlobsAccessTierAsync() example.
- Authentication scopes and required permissions are not explained.

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
hyoka run --prompt-id storage-dp-dotnet-batch --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
