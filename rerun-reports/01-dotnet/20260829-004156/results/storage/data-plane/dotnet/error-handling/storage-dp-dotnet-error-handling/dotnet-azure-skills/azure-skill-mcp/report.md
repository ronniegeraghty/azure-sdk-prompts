# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 85.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 85.5s |
| Timestamp | 2026-08-28T18:50:13Z |
| Files Generated | 0 |
| Event Count | 3002 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 68.4s |
| Review | 16.0s |
| **Total** | **85.5s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=354585 out=2813 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 677ms

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

### ✅ 2. glob — 3667ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 3991ms

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

### ✅ 4. azure-get_azure_bestpractices (via azure) — 3333ms

**Input:**

```json
{"intent":"Provide idiomatic .NET Azure.Storage.Blobs error handling patterns for RequestFailedException, including 403, 404, 429, error codes, HTTP status, retries, and cancellation.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 4321ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for .NET Azure.Storage.Blobs RequestFailedException handling, BlobErrorCode values, HTTP 403/404/429 behavior, and retry configuration.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 2265ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET SDK code-generation best practices for Blob Storage exception handling.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 2909ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find RequestFailedException and Blob Storage retry/error documentation.","parameters":{"query":"Azure.Storage.Blobs .NET RequestFailedException Status ErrorCode 403 404 429 BlobErrorCode retry options"}}
```

**Output:**

```
{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\nAny application that runs in the cloud or communicates with remote services and resources must be able to handle transient faults. It\u0027s common for these applications to experience faults due to a momentary loss of network connectivity, a request timeout when a service or resource is busy, or other factors. Developers should build applications to handle transient faults transparently to improve stability and resiliency.\nIn this article, you learn how to use the Azure Storage client library for .NET to set up a retry policy for an application that connects to Azure Blob Storage. Retry policies define how the application handles failed requests, and should always be tuned to match the business requirements of the application and the nature of the failure.\n## Configure retry options\nRetry policies for Blob Storage are configured programmatically, offering control over how retry options are applied to various service requests and scenarios. For example, a web app issuing requests based on user interaction might implement a policy with fewer retries and shorter delays to increase responsiveness and notify the user when an error occurs. Alternatively, an app or component running batch requests in the background might increase the number of retries and use an exponential backoff strategy to allow the request time to complete successfully.\nThe following table lists the properties of the [RetryOptions](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions) class, along with the type, a brief description, and the default value if you make no changes. You should be proactive in tuning the values of these properties to meet the needs of your app.","contentUrl":"https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy"},{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\n## Use geo-redundancy to i
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 2583ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# Blob Storage exception handling examples.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs RequestFailedException catch Status ErrorCode BlobErrorCode C# retry"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates handling Azure SDK exceptions by wrapping calls in try-catch blocks to manage common failure scenarios like 404 (not found) and 403 (forbidden) errors.\npackage: Azure\nlanguage: csharp\n","codeSnippet":"using Azure;\r\n\r\ntry\r\n{\r\n    var blob = await blobClient.DownloadAsync();\r\n}\r\ncatch (RequestFailedException ex) when (ex.Status == 404)\r\n{\r\n    // Blob not found\r\n}\r\ncatch (RequestFailedException ex) when (ex.Status == 403)\r\n{\r\n    // Insufficient permissions\r\n}\r\ncatch (RequestFailedException ex)\r\n{\r\n    _logger.LogError(ex, \u0022Azure SDK call failed with status {Status}\u0022, ex.Status);\r\n}","link":"https://learn.microsoft.com/entra/msidweb/call-downstream-apis/azure-sdks#follow-best-practices","language":"csharp"},{"description":"description: Imports namespaces for Azure Storage Blobs and related models to enable blob storage operations.\npackage: Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"using Azure.Storage.Blobs;\r\nusing Azure.Storage.Blobs.Models;","link":"https://learn.microsoft.com/azure/storage/common/storage-use-data-movement-library#use-extension-methods-for-blobcontainerclient","language":"csharp"},{"description":"description: Handles a RequestFailedException thrown when a secret is not found, checking for HTTP status 404 and printing the error code.\npackage: Azure.Core\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022NonexistentSecret\u0022);\r\n}\r\n// handle exception with status code 404\r\ncatch (RequestFailedException e) when (e.Status == 404)\r\n{\r\n    // handle not found error\r\n    Console.WriteLine(\u0022ErrorCode \u0022 \u002B e.ErrorCode);\r\n}","link":"https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Demonstrates handling of RequestFailedException when calling the Complete method on the Azure AI Inferen
... (truncated)
```

### ✅ 9. rg — 780ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt","pattern":"\"title\"|\"content\"|\"url\""}
```

**Output:**

```
Output too large to read at once (35.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787943049745-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt:{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\nAny application that runs in the cloud or communicates with remote services and resources must be able to handle transient faults. It\u0027s common for these applications to experience faults due to a momentary loss of network connectivity, a request timeout when a service or resource i
```

### ✅ 10. powershell — 1255ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt' | ConvertFrom-Json; $d.results | Select-Object title,url | Format-Table -AutoSize","description":"List relevant Microsoft documentation results"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-5b784eba08f5424196287ca61fca50d6.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 11. view — 729ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt b/C:/Users/weidxu/AppData/Local/Temp/1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787943045839-copilot-tool-output-5b784eba08f5424196287ca61fca50d6.txt
@@ -1,6 +1,6 @@
 {"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\nAny application that runs in the cloud or communicates with remote services and resources must be able to handle transient faults. It\u0027s common for these applications to experience faults due to a momentary loss of network connectivity, a request timeout when a service or resource is busy, or other factors. Developers should build applications to handle transient faults transparently to improve stability and resiliency.\nIn this article, you learn how to use the Azure Storage client library for .NET to set up a retry policy for an application that connects to Azure Blob Storage. Retry policies define how the application handles failed requests, and should always be tuned to match the business requirements of the application and the nature of the failure.\n## Configure retry options\nRetry policies for Blob Storage are configured programmatically, offering control over how retry options are applied to various service requests and scenarios. For example, a web app issuing requests based on user interaction might implement a policy with fewer retries and shorter delays to increase responsiveness and notify the user when an error occurs. Alternatively, an app or component running batch requests in the background might increase the number of retries and use an exponential backoff strategy to allow the request time to complete successfully.\nThe following table lists the pr
... (truncated)
```

### ✅ 12. azure-documentation (via azure) — 3061ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read full official retry guidance for Blob Storage .NET.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy"}}
```

**Output:**

```
# Azure Blob Storage error codes

You might see the following error codes returned by an operation against Azure Blob Storage.

| Error code | HTTP status code | User message |
| --- | --- | --- |
| AppendPositionConditionNotMet | Precondition Failed (412) | The append position condition specified was not met. |
| BlobAlreadyExists | Conflict (409) | The specified blob already exists. |
| BlobImmutableDueToLegalHold | Conflict (409) | This operation is not permitted as the blob is immutable due to one or more legal holds. |
| BlobImmutableDueToPolicy | Conflict (409) | This operation is not permitted as the blob is immutable due to a policy. |
| BlobNotFound | Not Found (404) | The specified blob does not exist. |
| BlobOperationNotSupported | Conflict (409) | The operation is not supported in this scenario. This error can appear if you use the [PutBlock](put-block) or [PutBlockList](put-block-list) operation against a blob that was created by using [Path - Create](/en-us/rest/api/storageservices/datalakestoragegen2/path/create) operation of the [Azure Data Lake Storage Gen2 REST API](data-lake-storage-gen2). |
| BlobOverwritten | Conflict (409) | The blob has been recreated since the previous snapshot was taken. |
| BlobTierInadequateForContentLength | Conflict (409) | The specified blob tier size limit cannot be less than content length. |
| BlockCountExceedsLimit | Conflict (409) | The committed block count cannot exceed the maximum limit of 50,000 blocks. |
| BlockCountExceedsLimit | Conflict (409) | The uncommitted block count cannot exceed the maximum limit of 100,000 blocks. |
| BlockListTooLong | Bad Request (400) | The blocklist may not contain more than 50,000 blocks. |
| CannotChangeToLowerTier | Conflict (409) | A higher blob tier has already been explicitly set. |
| CannotVerifyCopySource | Internal Server Error (500) | Could not verify the copy source within the specified time. Examine the HTTP status code and message for more informa
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 2701ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read current Azure Storage Blobs .NET troubleshooting details.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs-readme?view=azure-dotnet#troubleshooting"}}
```

**Output:**

```
# Azure Storage Blobs client library for .NET - version 12.29.2

> 
> Server Version: 2021-02-12, 2020-12-06, 2020-10-02, 2020-08-04, 2020-06-12, 2020-04-08, 2020-02-10, 2019-12-12, 2019-07-07, and 2019-02-02

Azure Blob storage is Microsoft's object storage solution for the cloud. Blob storage is optimized for storing massive amounts of unstructured data. Unstructured data is data that does not adhere to a particular data model or definition, such as text or binary data.

[Source code](https://github.com/Azure/azure-sdk-for-net/tree/Azure.Storage.Blobs_12.29.2/sdk/storage/Azure.Storage.Blobs/src) | [Package (NuGet)](https://www.nuget.org/packages/Azure.Storage.Blobs/) | [API reference documentation](/en-us/dotnet/api/azure.storage.blobs) | [REST API documentation](/en-us/rest/api/storageservices/blob-service-rest-api) | [Product documentation](/en-us/azure/storage/blobs/storage-blobs-overview)

## Getting started

### Install the package

Install the Azure Storage Blobs client library for .NET with [NuGet](https://www.nuget.org/):

```dotnetcli
dotnet add package Azure.Storage.Blobs
```

### Prerequisites

You need an [Azure subscription](https://azure.microsoft.com/free/dotnet/) and a [Storage Account](/en-us/azure/storage/common/storage-account-overview) to use this package.

To create a new Storage Account, you can use the [Azure Portal](/en-us/azure/storage/common/storage-quickstart-create-account?tabs=azure-portal), [Azure PowerShell](/en-us/azure/storage/common/storage-quickstart-create-account?tabs=azure-powershell), or the [Azure CLI](/en-us/azure/storage/common/storage-quickstart-create-account?tabs=azure-cli). Here's an example using the Azure CLI:

```Powershell
az storage account create --name MyStorageAccount --resource-group MyResourceGroup --location westus --sku Standard_LRS
```

### Authenticate the client

In order to interact with the Azure Blobs Storage service, you'll need to create an instance of the BlobServiceC
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 2372ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Blob Storage service error code reference.","parameters":{"url":"https://learn.microsoft.com/rest/api/storageservices/blob-service-error-codes"}}
```

**Output:**

```
# Implement a retry policy with .NET

Any application that runs in the cloud or communicates with remote services and resources must be able to handle transient faults. It's common for these applications to experience faults due to a momentary loss of network connectivity, a request timeout when a service or resource is busy, or other factors. Developers should build applications to handle transient faults transparently to improve stability and resiliency.

In this article, you learn how to use the Azure Storage client library for .NET to set up a retry policy for an application that connects to Azure Blob Storage. Retry policies define how the application handles failed requests, and should always be tuned to match the business requirements of the application and the nature of the failure.

## Configure retry options

Retry policies for Blob Storage are configured programmatically, offering control over how retry options are applied to various service requests and scenarios. For example, a web app issuing requests based on user interaction might implement a policy with fewer retries and shorter delays to increase responsiveness and notify the user when an error occurs. Alternatively, an app or component running batch requests in the background might increase the number of retries and use an exponential backoff strategy to allow the request time to complete successfully.

The following table lists the properties of the [RetryOptions](/en-us/dotnet/api/azure.core.retryoptions) class, along with the type, a brief description, and the default value if you make no changes. You should be proactive in tuning the values of these properties to meet the needs of your app.

| Property | Type | Description | Default value |
| --- | --- | --- | --- |
| [Delay](/en-us/dotnet/api/azure.core.retryoptions.delay) | [TimeSpan](/en-us/dotnet/api/system.timespan) | The delay between retry attempts for a fixed approach or the delay on which to base calculations for a backo
... (truncated)
```

## Copilot Response

I’m grounding the patterns in the current Azure SDK guidance, including retry behavior and service error codes.

`Azure.Storage.Blobs` reports service failures as `Azure.RequestFailedException`. Its key properties are:

- `Status`: HTTP status code, such as `403`, `404`, or `429`
- `ErrorCode`: storage service code, such as `ContainerNotFound` or `AuthorizationPermissionMismatch`
- `Message`: diagnostic message
- `GetRawResponse()`: response headers, including `x-ms-request-id`

```csharp
using Azure;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Microsoft.Extensions.Logging;

public static async Task DownloadBlobAsync(
    BlobClient blobClient,
    Stream destination,
    ILogger logger,
    CancellationToken cancellationToken)
{
    try
    {
        await blobClient.DownloadToAsync(destination, cancellationToken);
    }
    catch (RequestFailedException ex)
        when (ex.Status == 404 &&
              ex.ErrorCode == BlobErrorCode.ContainerNotFound)
    {
        logger.LogWarning(
            "Container was not found. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        throw new InvalidOperationException(
            "The configured blob container does not exist.", ex);
    }
    catch (RequestFailedException ex)
        when (ex.Status == 404 &&
              ex.ErrorCode == BlobErrorCode.BlobNotFound)
    {
        logger.LogInformation("Blob was not found.");
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        logger.LogError(
            ex,
            "Blob access denied. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        // Usually a credentials, RBAC, SAS, firewall, or ACL problem.
        // Retrying unchanged credentials normally will not help.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        logger.LogWarning(
            ex,
            "Blob request remained throttled after SDK retries. " +
            "Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        throw;
    }
    catch (RequestFailedException ex)
    {
        string? requestId = null;
        ex.GetRawResponse()?.Headers.TryGetValue(
            "x-ms-request-id",
            out requestId);

        logger.LogError(
            ex,
            "Blob operation failed. Status={Status}, ErrorCode={ErrorCode}, " +
            "RequestId={RequestId}",
            ex.Status,
            ex.ErrorCode,
            requestId);

        throw;
    }
}
```

Use the error code when you need to distinguish failures sharing the same status. For example, both a missing blob and a missing container produce `404`:

```csharp
catch (RequestFailedException ex)
    when (ex.ErrorCode == BlobErrorCode.ContainerNotFound)
{
    // Specifically the container is missing.
}
```

For authorization failures, inspect `ErrorCode`; common values include `AuthenticationFailed`, `AuthorizationFailure`, and `AuthorizationPermissionMismatch`. Prefer status-based handling for the general `403` case because the exact code depends on whether the problem is authentication, RBAC, SAS permissions, networking, or ACLs.

### Throttling and retries

The SDK automatically retries transient responses, including `429`, using the server’s `Retry-After` header when present. Therefore, a caught `429` generally means configured retries were exhausted. Configure retry behavior on the client rather than writing a manual retry loop around each operation:

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
    new Uri("https://<account-name>.blob.core.windows.net"),
    new DefaultAzureCredential(),
    options);
```

Blob Storage can also indicate overload with `503 Service Unavailable` and an error code such as `ServerBusy`, so transient-failure handling often includes both:

```csharp
catch (RequestFailedException ex)
    when (ex.Status is 429 or 503)
{
    // The SDK has already exhausted its configured retries.
    logger.LogWarning(
        ex,
        "Storage service is busy. Status={Status}, ErrorCode={ErrorCode}",
        ex.Status,
        ex.ErrorCode);

    throw;
}
```

Avoid check-then-act patterns such as calling `ExistsAsync()` before downloading; the resource can disappear between calls. Catch the operation’s `RequestFailedException` instead. For operations where absence is expected, APIs such as `DeleteIfExistsAsync` and `ExistsAsync` provide non-exceptional results.

References:

- [Azure Storage Blobs client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/storage.blobs-readme?view=azure-dotnet#troubleshooting)
- [Configure Blob Storage retry policies in .NET](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy)
- [Azure Blob Storage error codes](https://learn.microsoft.com/rest/api/storageservices/blob-service-error-codes)

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response thoroughly covers RequestFailedException, error metadata, core requested statuses, retries, and diagnostics, but omits 409 handling and conditional ETag/lease failures.

### Strengths

- Uses idiomatic exception filters with RequestFailedException.
- Clearly extracts Status, ErrorCode, and the storage request ID.
- Configures retries correctly through BlobClientOptions instead of a manual retry loop.
- Provides practical structured logging and troubleshooting guidance.

### Issues

- No handling or explanation of HTTP 409 conflict responses.
- No examples of ETag or lease conditional failures, including relevant request conditions and service error codes.

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
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
