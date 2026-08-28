# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 193.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 193.9s |
| Timestamp | 2026-08-28T19:00:18Z |
| Files Generated | 0 |
| Event Count | 4376 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 175.4s |
| Review | 17.3s |
| **Total** | **193.9s** |

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
| Token Usage | in=309812 out=5657 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 1534ms

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

### ✅ 2. glob — 5768ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md,*dotnet*.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 7251ms

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

### ✅ 4. azure-get_azure_bestpractices (via azure) — 6540ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Generate modern .NET Azure.Storage.Blobs code configuring BlobClientOptions retries, operation cancellation timeouts, custom per-operation retry behavior, transient versus authentication error handling, and circuit breaking.","parameters":{"language":"C#","operation":"code-generation","resource":"Azure Blob Storage"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 7945ms

**Input:**

```json
{"command":"search","intent":"Find official Microsoft documentation for Azure.Storage.Blobs .NET BlobClientOptions RetryOptions defaults, retryable HTTP status codes, NetworkTimeout, cancellation tokens, custom HttpPipelinePolicy, and circuit breaker guidance.","parameters":{"query":"Azure.Storage.Blobs .NET BlobClientOptions Retry RetryOptions retryable status codes NetworkTimeout cancellation token custom HttpPipelinePolicy circuit breaker"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 45739ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET Blob Storage code-generation best practices.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 46479ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official retry guidance.","parameters":{"query":"site:learn.microsoft.com Azure Storage .NET retry policy BlobClientOptions RetryOptions MaxRetries Delay MaxDelay Mode NetworkTimeout retry status codes"}}
```

**Output:**

```
{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\n## Configure retry options\n| Property | Type | Description | Default value| \n|  --- | --- | --- | ---  |\n| [Delay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.delay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The delay between retry attempts for a fixed approach or the delay on which to base calculations for a backoff-based approach. If the service provides a Retry-After response header, the next retry is delayed by the duration specified by the header value. | 0.8 second |\n| [MaxDelay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.maxdelay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The maximum permissible delay between retry attempts when the service doesn\u0027t provide a Retry-After response header. If the service provides a Retry-After response header, the next retry is delayed by the duration specified by the header value. | 1 minute |\n| [MaxRetries](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.maxretries) | int | The maximum number of retry attempts before giving up. | 5 (see note) |\n| [Mode](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.mode) | [RetryMode](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retrymode) | The approach to use for calculating retry delays. | Exponential |\n| [NetworkTimeout](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.networktimeout) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The timeout applied to an individual network operation. | 100 seconds |\n\nNote\n\u0060StorageClientOptions\u0060 increases the default value for \u0060MaxRetries\u0060 from 3 to 5. All other properties have the same default values as \u0060RetryOptions\u0060.\nIn this code example for Blob Storage, we configure the retry op
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 46100ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# Blob retry samples.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs BlobClientOptions RetryOptions custom retry policy cancellation token UploadAsync"}}
```

**Output:**

```
{"results":[{"description":"description: Configures retry handling for Azure Blob Storage client to use a primary then secondary approach, where read request retries from the primary region are redirected to the secondary region using BlobClientOptions and BlobServiceClient.\npackage: Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"string accountName = \u0022\u003CYOURSTORAGEACCOUNTNAME\u003E\u0022;\r\nUri primaryAccountUri = new Uri($\u0022https://{accountName}.blob.core.windows.net/\u0022);\r\nUri secondaryAccountUri = new Uri($\u0022https://{accountName}-secondary.blob.core.windows.net/\u0022);\r\n\r\n// Provide the client configuration options for connecting to Azure Blob storage\r\nBlobClientOptions blobClientOptions = new BlobClientOptions()\r\n{\r\n    Retry = {\r\n        // The delay between retry attempts for a fixed approach or the delay\r\n        // on which to base calculations for a backoff-based approach\r\n        Delay = TimeSpan.FromSeconds(2),\r\n\r\n        // The maximum number of retry attempts before giving up\r\n        MaxRetries = 5,\r\n\r\n        // The approach to use for calculating retry delays\r\n        Mode = RetryMode.Exponential,\r\n\r\n        // The maximum permissible delay between retry attempts\r\n        MaxDelay = TimeSpan.FromSeconds(10)\r\n    },\r\n\r\n    // If the GeoRedundantSecondaryUri property is set, the secondary Uri will be used for \r\n    // GET or HEAD requests during retries.\r\n    // If the status of the response from the secondary Uri is a 404, then subsequent retries\r\n    // for the request will not use the secondary Uri again, as this indicates that the resource \r\n    // may not have propagated there yet.\r\n    // Otherwise, subsequent retries will alternate back and forth between primary and secondary Uri.\r\n    GeoRedundantSecondaryUri = secondaryAccountUri\r\n};\r\n\r\n// Create a BlobServiceClient object using the configuration options above\r\nBlobServiceClient blobServiceClient = ne
... (truncated)
```

### ✅ 9. web_search (via github-mcp-server) — 46847ms

**Input:**

```json
{"query":"In the current Azure SDK for .NET Azure.Core and Azure.Storage.Blobs, what are the exact public APIs for ClientOptions.RetryPolicy and RetryPolicy customization, and which HTTP status codes are retried by default? Cite official Microsoft docs or azure-sdk-for-net source."}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here are the exact public APIs and behaviors for retry policy customization in Azure.Core and Azure.Storage.Blobs, with references to official Microsoft documentation and source.\n\n## 1. Azure.Core - `ClientOptions.RetryPolicy`\n\n- The `RetryPolicy` property is part of the abstract `ClientOptions` class in Azure.Core and allows you to set a custom retry policy on any client options type that inherits from it.\n- API signature:\n  ```csharp\n  public Azure.Core.Pipeline.HttpPipelinePolicy? RetryPolicy { get; set; }\n  ```\n  - If you set `RetryPolicy`, it will be used instead of the `Retry` property (which is of type `RetryOptions`) for all retries.\n  - You can subclass `RetryPolicy` to implement custom retry logic without having to write a full retry mechanism from scratch.\n- [Reference: Microsoft Learn - ClientOptions.RetryPolicy Property](https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy)【3:0†source】.\n- [Reference: Azure SDK source - ClientOptions.cs](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ClientOptions.cs)【3:5†source】.\n\n### Customizing a Retry Policy\n- Implement your own class that derives from `Azure.Core.Pipeline.RetryPolicy` and override methods such as:\n  - `ShouldRetry(HttpMessage, Exception)`\n  - `ShouldRetryAsync(HttpMessage, Exception)`\n  - `Process(HttpMessage, ...)`  \n  [Reference: RetryPolicy Class API](https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy)【3:8†source】.\n\n## 2. Azure.Storage.Blobs - Retry Policy Customization\n\n- For most use, you configure retries using the `Retry` property of `BlobClientOptions` (inherits from `StorageClientOptions`), which exposes a `RetryOptions` object.\n- This is *not* a settable property but has settable sub-properties:\n  - `Delay`\n  - `MaxDelay`\n  - `MaxRetries`\n  - `Mode` (exponential/fixed)\n  - `NetworkTimeout`\n\nExample:\n```csharp\nvar opti
... (truncated)
```

### ✅ 10. github-mcp-server-get_file_contents (via github-mcp-server) — 6803ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: c8102635d8888d976a9ec9269e44a09f8595bc44)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using Azure.Core.Pipeline;
using Microsoft.Extensions.Configuration;

namespace Azure.Core
{
    /// <summary>
    /// Base type for all client option types, exposes various common client options like <see cref="Diagnostics"/>, <see cref="Retry"/>, <see cref="Transport"/>.
    /// </summary>
    public abstract class ClientOptions
    {
        private HttpPipelineTransport _transport;
        internal bool IsCustomTransportSet { get; private set; }

        /// <summary>
        /// Gets the default set of <see cref="ClientOptions"/>. Changes to the <see cref="Default"/> options would be reflected
        /// in new instances of <see cref="ClientOptions"/> type created after changes to <see cref="Default"/> were made.
        /// </summary>
        public static ClientOptions Default { get; private set; } = new DefaultClientOptions();

        // For testing
        internal static void ResetDefaultOptions()
        {
            Default = new DefaultClientOptions();
        }

        /// <summary>
        /// Creates a new instance of <see cref="ClientOptions"/>.
        /// </summary>
        protected ClientOptions() : this(Default, null)
        {
        }

        /// <summary>
        /// Creates a new instance of <see cref="ClientOptions"/> with the specified <see cref="DiagnosticsOptions"/>.
        /// </summary>
        /// <param name="diagnostics"><see cref="DiagnosticsOptions"/> to be used for <see cref="Diagnostics"/>.</param>
        protected ClientOptions(DiagnosticsOptions? diagnostics)
            : this(Default, diagnostics)
        {
        }

        internal ClientOptions(ClientOptions? clientOptions, DiagnosticsOptions? diagnostics)
        {
       
... (truncated)
```

### ✅ 11. github-mcp-server-get_file_contents (via github-mcp-server) — 6437ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/ClientOptions.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
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

### ✅ 12. rg — 5644ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943673351-copilot-tool-output-d2183255e74146f7b97951b09290ea08.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943673195-copilot-tool-output-6bf319ae1d0d457292d60ca9c2d26df5.txt"],"pattern":"retry|Retry|408|429|500|502|503|504|NetworkTimeout|https://learn.microsoft.com"}
```

**Output:**

```
Output too large to read at once (66.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787943724925-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787943673351-copilot-tool-output-d2183255e74146f7b97951b09290ea08.txt:{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\n## Configure retry options\n| Property | Type | Description | Default value| \n|  --- | --- | --- | ---  |\n| [Delay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.delay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The delay betwee
```

### ✅ 13. azure-documentation (via azure) — 6051ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the official Blob Storage retry policy page in full.","parameters":{"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}}
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

### ✅ 14. github-mcp-server-search_code (via github-mcp-server) — 4490ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"IsRetriableResponse\" \"429\" path:sdk/core/Azure.Core/src"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"RetryPolicy.cs","path":"sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs","sha":"a729cb2a06b751b0f53499f0ca00d31ac7c919b1","text_matches":[{"fragment":"                }\n\n                // Response.IsError is true if we get here\n                return message.ResponseClassifier.IsRetriableResponse(message);\n            }\n\n            // out of retries","matches":[{"indices":[128,147],"text":"IsRetriableResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"SseHttpPipelineResponseHandler.cs","path":"sdk/core/Azure.Core/src/Pipeline/Internal/SseHttpPipelineResponseHandler.cs","sha":"d109bdecd43a5336ef5fad428527a136306b43b4","text_matches":[{"fragment":"            bool isRetriable = classifyResponse &&\n                _classifier.IsRetriableResponse(message);\n            message.ClearResponse();","matches":[{"indices":[79,98],"text":"IsRetriableResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Pipeline/Internal/SseHttpPipelineResponseHandler.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"},{"fragment":"            bool isRetriable = classifyResponse &&\n                _classifier.IsRetriableResponse(message);\n            message.ClearResponse();","matches":[{"indices":[79,98],"text":"IsRetriableResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/Pipeline/Internal/SseHttpPipelineResponseHandler.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]}],"total_count":2}
```

### ✅ 15. github-mcp-server-search_code (via github-mcp-server) — 4097ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-net \"IsRetriableResponse\" path:sdk/core/Azure.Core/src/Pipeline"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"ResponseClassifier.cs","path":"sdk/core/Azure.Core/src/ResponseClassifier.cs","sha":"483c5dcbc73e0819fd33fb35f6b91013d889eb7b","text_matches":[{"fragment":"        /// <summary>\n        /// Specifies if the request contained in the <paramref name=\"message\"/> should be retried.\n        /// </summary>\n        public virtual bool IsRetriableResponse(HttpMessage message)\n        {\n            switch (message.Response.Status)\n            {","matches":[{"indices":[173,192],"text":"IsRetriableResponse"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/core/Azure.Core/src/ResponseClassifier.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]}],"total_count":1}
```

### ✅ 16. azure-documentation (via azure) — 3707ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official .NET resilience and circuit breaker guidance compatible with async Azure Blob calls.","parameters":{"query":"site:learn.microsoft.com .NET resilience circuit breaker AddResiliencePipeline CircuitBreakerStrategyOptions ExecuteAsync cancellation token"}}
```

**Output:**

```
{"results":[{"title":"Introduction to resilient app development","content":"# Introduction to resilient app development\n## Build a resilience pipeline\nTo use resilience, you must first build a pipeline of resilience-based strategies. Each configured strategy executes in order of configuration. In other words, order is important. The entry point is an extension method on the [IServiceCollection](https://learn.microsoft.com/en-us/dotnet/api/microsoft.extensions.dependencyinjection.iservicecollection) type, named \u0060AddResiliencePipeline\u0060. This method takes an identifier of the pipeline and a delegate that configures the pipeline. The delegate is passed an instance of \u0060ResiliencePipelineBuilder\u0060, which is used to add resilience strategies to the pipeline.\nConsider the following string-based \u0060key\u0060 example:\nThe preceding code:\n1. Creates a new \u0060ServiceCollection\u0060 instance.\n2. Defines a \u0060key\u0060 to identify the pipeline.\n3. Adds a resilience pipeline to the \u0060ServiceCollection\u0060 instance.\n4. Configures the pipeline with a retry and timeout strategies.\nEach pipeline is configured for a given \u0060key\u0060, and each \u0060key\u0060 is used to identify its corresponding \u0060ResiliencePipeline\u0060 when getting the pipeline from the provider. The \u0060key\u0060 is a generic type parameter of the \u0060AddResiliencePipeline\u0060 method.\n### Resilience pipeline builder extensions\nTo add a strategy to the pipeline, call any of the available \u0060Add*\u0060 extension methods on the \u0060ResiliencePipelineBuilder\u0060 instance.\n1. \u0060AddRetry\u0060: Try again if something fails, which is useful when the problem is temporary and might go away.\n2. \u0060AddCircuitBreaker\u0060: Stop trying if something is broken or busy, which benefits you by avoiding wasted time and making things worse.\n3. \u0060AddTimeout\u0060: Give up if something takes too long, which can improve performance by freeing up resources.
... (truncated)
```

### ✅ 17. github-mcp-server-get_file_contents (via github-mcp-server) — 1922ms

**Input:**

```json
{"owner":"Azure","path":"sdk/core/Azure.Core/src/ResponseClassifier.cs","ref":"refs/heads/main","repo":"azure-sdk-for-net"}
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

## Copilot Response

I’m grounding the retry behavior in the current .NET SDK guidance, including the distinction between SDK retries, operation deadlines, and circuit breaking.

The storage skill has no .NET-specific acceptance file, so I’m using the official Azure.Core API references and samples for the exact retry hooks and defaults.

One point needs precision: Azure Storage’s documented retry set differs from some generic HTTP retry guidance around `429`, so I’m checking the current Azure.Core implementation rather than assuming throttling behavior.

Configure retries at the client level with `BlobClientOptions.Retry`, and enforce a separate end-to-end deadline with a cancellation token. `NetworkTimeout` only limits each individual network attempt; it does not cap the entire upload including retries and backoff.

```bash
dotnet add package Azure.Storage.Blobs
dotnet add package Azure.Identity
dotnet add package Polly.Core
```

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;
using Polly;
using Polly.CircuitBreaker;

public static class BlobClients
{
    public static BlobClient CreateStandardClient(
        Uri blobUri,
        TokenCredential credential)
    {
        var options = new BlobClientOptions
        {
            Retry =
            {
                Mode = RetryMode.Exponential,

                // Five retries after the initial request: up to six attempts total.
                MaxRetries = 5,

                // Base delay used by exponential backoff.
                Delay = TimeSpan.FromSeconds(1),

                // Caps the calculated delay between attempts.
                MaxDelay = TimeSpan.FromSeconds(30),

                // Timeout for each individual network operation.
                NetworkTimeout = TimeSpan.FromSeconds(30)
            }
        };

        return new BlobClient(blobUri, credential, options);
    }

    public static BlobClient CreateHighValueClient(
        Uri blobUri,
        TokenCredential credential)
    {
        // Keep this client as a singleton alongside the standard client.
        // Its policy applies only to operations sent through this instance.
        var options = new BlobClientOptions
        {
            Retry =
            {
                Mode = RetryMode.Exponential,
                MaxRetries = 8,
                Delay = TimeSpan.FromSeconds(2),
                MaxDelay = TimeSpan.FromMinutes(1),
                NetworkTimeout = TimeSpan.FromSeconds(45)
            }
        };

        return new BlobClient(blobUri, credential, options);
    }
}
```

In Azure-hosted applications, pass a `ManagedIdentityCredential`. For local development, `DefaultAzureCredential` is appropriate:

```csharp
TokenCredential credential = new DefaultAzureCredential();

var blobUri = new Uri(
    "https://<account>.blob.core.windows.net/<container>/<blob>");

BlobClient regularBlob = BlobClients.CreateStandardClient(blobUri, credential);
BlobClient highValueBlob = BlobClients.CreateHighValueClient(blobUri, credential);
```

## Per-operation timeout and error classification

Use `CancelAfter` to bound the complete upload—including SDK retries and backoff:

```csharp
public static async Task UploadWithDeadlineAsync(
    BlobClient blob,
    Stream content,
    TimeSpan operationTimeout,
    CancellationToken cancellationToken)
{
    using var timeoutCts =
        CancellationTokenSource.CreateLinkedTokenSource(cancellationToken);

    timeoutCts.CancelAfter(operationTimeout);

    try
    {
        await blob.UploadAsync(
            content,
            overwrite: true,
            cancellationToken: timeoutCts.Token);
    }
    catch (RequestFailedException ex) when (ex.Status is 401 or 403)
    {
        // Non-transient: bad/expired credentials, insufficient RBAC, SAS scope,
        // firewall authorization, or another access configuration problem.
        throw new InvalidOperationException(
            $"Blob authorization failed ({ex.Status}, {ex.ErrorCode}). " +
            "The operation should not be retried without correcting access.",
            ex);
    }
    catch (AuthenticationFailedException ex)
    {
        // Credential acquisition failed before a storage request was sent.
        throw new InvalidOperationException(
            "Azure credential acquisition failed. Retrying the blob request " +
            "with the same credential configuration is unlikely to help.",
            ex);
    }
    catch (OperationCanceledException ex)
        when (timeoutCts.IsCancellationRequested &&
              !cancellationToken.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"The upload exceeded its {operationTimeout} deadline.",
            ex);
    }
    catch (Exception ex) when (BlobFailureClassifier.IsTransient(ex))
    {
        // SDK retries are already exhausted at this point.
        // Record telemetry or propagate to a circuit breaker.
        throw new IOException(
            "The blob upload failed after transient retries were exhausted.",
            ex);
    }
}

public static class BlobFailureClassifier
{
    private static readonly HashSet<int> RetryableStatusCodes =
    [
        408, // Request Timeout
        429, // Too Many Requests
        500, // Internal Server Error
        502, // Bad Gateway
        503, // Service Unavailable
        504  // Gateway Timeout
    ];

    public static bool IsTransient(Exception exception)
    {
        return exception switch
        {
            RequestFailedException { Status: 0 } => true,
            RequestFailedException requestFailure =>
                RetryableStatusCodes.Contains(requestFailure.Status),

            IOException => true,
            HttpRequestException => true,
            TimeoutException => true,

            // Multiple transport failures can be aggregated after retries.
            AggregateException aggregate =>
                aggregate.InnerExceptions.Count > 0 &&
                aggregate.InnerExceptions.All(IsTransient),

            _ => false
        };
    }
}
```

Call it with an operation-wide limit:

```csharp
await using FileStream content = File.OpenRead("payload.bin");

await UploadWithDeadlineAsync(
    regularBlob,
    content,
    operationTimeout: TimeSpan.FromMinutes(5),
    cancellationToken);
```

Do not add another general-purpose retry loop around `UploadAsync`. That would multiply retries—for example, six SDK attempts multiplied by six outer attempts—and can produce unexpectedly long delays.

## Circuit breaker for sustained failures

`Azure.Storage.Blobs` provides retry behavior but not an application-level circuit breaker. Polly can wrap the complete SDK operation. The breaker below counts only failures classified as transient, so authentication and authorization errors do not trip it.

```csharp
ResiliencePipeline blobCircuitBreaker =
    new ResiliencePipelineBuilder()
        .AddCircuitBreaker(new CircuitBreakerStrategyOptions
        {
            ShouldHandle = new PredicateBuilder()
                .Handle<Exception>(BlobFailureClassifier.IsTransient),

            // Open when at least half of qualifying calls fail...
            FailureRatio = 0.50,

            // ...during this rolling interval...
            SamplingDuration = TimeSpan.FromSeconds(30),

            // ...provided enough calls occurred to make the sample meaningful.
            MinimumThroughput = 10,

            // Reject calls for this period before permitting a trial call.
            BreakDuration = TimeSpan.FromSeconds(30)
        })
        .Build();
```

Use the high-value client through that breaker:

```csharp
public static async Task UploadHighValueAsync(
    BlobClient highValueBlob,
    ResiliencePipeline circuitBreaker,
    Stream content,
    CancellationToken cancellationToken)
{
    try
    {
        await circuitBreaker.ExecuteAsync(
            async token =>
            {
                await UploadWithDeadlineAsync(
                    highValueBlob,
                    content,
                    operationTimeout: TimeSpan.FromMinutes(10),
                    cancellationToken: token);
            },
            cancellationToken);
    }
    catch (BrokenCircuitException ex)
    {
        throw new InvalidOperationException(
            "Blob Storage operations are temporarily disabled because " +
            "sustained transient failures opened the circuit.",
            ex);
    }
}
```

The circuit breaker operates on the result of the **entire SDK operation**. One failed operation counts once, after its internal retries are exhausted. Authentication failures are propagated immediately and excluded from the breaker.

## Default retry behavior

The current Azure.Core response classifier retries these HTTP statuses by default:

| Status | Meaning |
|---:|---|
| `408` | Request Timeout |
| `429` | Too Many Requests |
| `500` | Internal Server Error |
| `502` | Bad Gateway |
| `503` | Service Unavailable |
| `504` | Gateway Timeout |

It also retries:

- `IOException`
- `RequestFailedException` with status `0`, typically indicating no HTTP response due to a transport/network failure
- An `OperationCanceledException` caused by the SDK/transport timeout, but **not** cancellation requested through your own token

The default classifier does **not** retry `401` or `403`. Other client errors such as `400`, `404`, `409`, and `412` are also non-retryable by default because they generally require changing the request or application state. If the service returns `Retry-After`, Azure.Core uses that value rather than the locally calculated delay.

`BlobClientOptions.RetryPolicy` can replace `RetryOptions` with a custom `HttpPipelinePolicy`, but that is normally necessary only when changing response classification itself. For operation-specific retry counts and backoff, dedicated long-lived clients with different `RetryOptions` are simpler and safer.

**References:** [Azure Blob Storage retry policies for .NET](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), [ClientOptions.RetryPolicy](https://learn.microsoft.com/dotnet/api/azure.core.clientoptions.retrypolicy), and the current [Azure.Core `ResponseClassifier` source](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs).

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The response thoroughly covers SDK retry configuration, timeout scopes, error classification, operation-specific clients, and Polly circuit breaking, but omits geo-redundant secondary retry configuration.

### Strengths

- Clearly distinguishes per-request NetworkTimeout from an end-to-end CancellationToken deadline.
- Accurately lists default retryable and common non-retryable HTTP status codes.
- Provides separate retry settings for high-value operations and a transient-only Polly circuit breaker.
- Warns against nested general-purpose retry loops that multiply SDK attempts.

### Issues

- Missing BlobClientOptions.GeoRedundantSecondaryUri configuration and explanation.

## Grader Results

- retry-configuration.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `BlobClientOptions.Retry` configuration with `RetryOptions`: Pass
      - `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed): Pass
      - `NetworkTimeout` for per-request timeouts: Pass
      - Default retryable status codes (408, 429, 500, 502, 503, 504): Pass
      - Non-retryable errors (400, 401, 403, 404, 409): Pass
      - Per-operation `CancellationToken` for timeout control: Pass
      - Geo-redundant retry with `GeoRedundantSecondaryUri`: Fail
      - Interaction with Polly or other resilience libraries: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
