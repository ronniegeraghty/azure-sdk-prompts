# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 125.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 125.7s |
| Timestamp | 2026-08-28T18:41:38Z |
| Files Generated | 0 |
| Event Count | 3852 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 108.1s |
| Review | 16.4s |
| **Total** | **125.7s** |

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
| Token Usage | in=376896 out=5918 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 688ms

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

### ✅ 2. view — 4635ms

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

### ✅ 3. view — 4296ms

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

### ✅ 4. azure-get_azure_bestpractices (via azure) — 3979ms

**Input:**

```json
{"intent":"Provide current best practices for .NET Azure.Storage.Blobs client configuration, RetryOptions, per-operation cancellation timeouts, custom HttpPipelinePolicy retry handling, and circuit breakers. No resource deployment.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 4953ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Core RetryOptions defaults and retryable status codes, BlobClientOptions retry configuration, custom HttpPipelinePolicy behavior, cancellation tokens for Blob uploads, and .NET circuit breaker integration.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 2793ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET SDK code-generation best practices for Blob Storage retries, timeouts, custom pipeline policies, and resilience.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 3477ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official docs for Azure.Core RetryOptions retry status codes and Blob Storage .NET retry/timeouts.","parameters":{"query":"Azure.Storage.Blobs .NET BlobClientOptions RetryOptions MaxRetries Delay MaxDelay NetworkTimeout retryable HTTP status codes 408 429 500 502 503 504 cancellation token"}}
```

**Output:**

```
{"results":[{"description":"description: Configures retry options for Azure Blob Storage client using BlobClientOptions, setting parameters like delay, max retries, retry mode, max delay, and network timeout, then creates a BlobServiceClient instance with these options and DefaultAzureCredential for authentication.\npackage: Azure.Storage.Blobs, Azure.Identity\nlanguage: csharp\n","codeSnippet":"// Provide the client configuration options for connecting to Azure Blob Storage\r\nBlobClientOptions blobOptions = new BlobClientOptions()\r\n{\r\n    Retry = {\r\n        Delay = TimeSpan.FromSeconds(2),\r\n        MaxRetries = 5,\r\n        Mode = RetryMode.Exponential,\r\n        MaxDelay = TimeSpan.FromSeconds(10),\r\n        NetworkTimeout = TimeSpan.FromSeconds(100)\r\n    },\r\n};\r\n\r\nBlobServiceClient blobServiceClient = new BlobServiceClient(\r\n    accountUri,\r\n    new DefaultAzureCredential(),\r\n    blobOptions);","link":"https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy#configure-retry-options","language":"csharp"},{"description":"description: Configures a BlobClientOptions object with a retry policy including delay, max retries, exponential backoff, max delay, and network timeout, and sets the GeoRedundantSecondaryUri property to enable geo-redundant secondary endpoint usage for Blob Storage requests.\npackage: Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"Uri secondaryAccountUri = new Uri($\u0022https://{accountName}-secondary.blob.core.windows.net/\u0022);\r\n\r\n// Provide the client configuration options for connecting to Azure Blob Storage\r\nBlobClientOptions blobOptionsGRS = new BlobClientOptions()\r\n{\r\n    Retry = {\r\n        Delay = TimeSpan.FromSeconds(2),\r\n        MaxRetries = 5,\r\n        Mode = RetryMode.Exponential,\r\n        MaxDelay = TimeSpan.FromSeconds(10),\r\n        NetworkTimeout = TimeSpan.FromSeconds(100)\r\n    },\r\n\r\n    // Set the secondary storage URI\r\n    GeoRedundantSecondaryUri = sec
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3130ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for configuring BlobClientOptions Retry and custom HttpPipelinePolicy.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs BlobClientOptions RetryOptions custom HttpPipelinePolicy per call cancellation token UploadAsync"}}
```

**Output:**

```
{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\n## Configure retry options\n| Property | Type | Description | Default value| \n|  --- | --- | --- | ---  |\n| [Delay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.delay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The delay between retry attempts for a fixed approach or the delay on which to base calculations for a backoff-based approach. If the service provides a Retry-After response header, the next retry is delayed by the duration specified by the header value. | 0.8 second |\n| [MaxDelay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.maxdelay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The maximum permissible delay between retry attempts when the service doesn\u0027t provide a Retry-After response header. If the service provides a Retry-After response header, the next retry is delayed by the duration specified by the header value. | 1 minute |\n| [MaxRetries](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.maxretries) | int | The maximum number of retry attempts before giving up. | 5 (see note) |\n| [Mode](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.mode) | [RetryMode](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retrymode) | The approach to use for calculating retry delays. | Exponential |\n| [NetworkTimeout](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.networktimeout) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The timeout applied to an individual network operation. | 100 seconds |\n\nNote\n\u0060StorageClientOptions\u0060 increases the default value for \u0060MaxRetries\u0060 from 3 to 5. All other properties have the same default values as \u0060RetryOptions\u0060.\nIn this code example for Blob Storage, we configure the retry op
... (truncated)
```

### ✅ 9. rg — 2484ms

**Input:**

```json
{"-n":true,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787942531407-copilot-tool-output-0f5e3fc2e7114a8889bac8bb5f11b69c.txt","pattern":"retry|Retry|408|429|500|502|503|504|NetworkTimeout|MaxRetries|HttpPipeline"}
```

**Output:**

```
Output too large to read at once (23.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787942537885-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787942531407-copilot-tool-output-0f5e3fc2e7114a8889bac8bb5f11b69c.txt:1:{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\n## Configure retry options\n| Property | Type | Description | Default value| \n|  --- | --- | --- | ---  |\n| [Delay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.delay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The delay betw
```

### ✅ 10. azure-documentation (via azure) — 3121ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the full official Blob Storage retry policy article for accurate configuration and defaults.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy"}}
```

**Output:**

```
{"results":[{"title":"HTTP pipelines and retries in the Azure SDK for C\u002B\u002B","content":"# HTTP pipelines and retries in the Azure SDK for C\u002B\u002B\n## Retry policy\n### How retries work\nThe retry policy is designed to handle transient failures that might occur when making HTTP requests to Azure services. When a request fails due to a transient error, the retry policy does the following:\n1. Determines if the failure is retryable.\n2. Calculates an appropriate delay.\n3. Waits for that delay.\n4. Retries the request.\nThe policy supports retrying on both transport-level failures (network issues) and certain HTTP status codes.\n### Default retry behavior\nBy default, the retry policy is configured with:\n1. Maximum of three retry attempts\n2. Initial retry delay of 800 milliseconds\n3. Maximum retry delay of 60 seconds\n4. Retryable status codes: 408, 429, 500, 502, 503, 504\nThe retry delay uses an exponential backoff strategy with jitter:\n1. First retry: ~800 ms\n2. Second retry: ~1,600 ms\n3. Third retry: ~3,200 ms\n4. The delay keeps doubling until it reaches the maximum retry delay.\n### When retries happen\nThe retry policy attempts to retry a request in the following scenarios:\n1. **Transport failures**:\n1.1. Network connectivity issues\n1.2. Connection timeouts\n1.3. DNS (Domain Name System) resolution failures\n2. **HTTP status codes**:\n2.1. 408 (Request timeout)\n2.2. 429 (Too Many Requests)\n2.3. 500 (Internal Server Error)\n2.4. 502 (Bad Gateway)\n2.5. 503 (Service Unavailable)\n2.6. 504 (Gateway timeout)\n3. **Service-specific retry logic**:\n3.1. Some services like Storage implement specialized retry logic for failover scenarios\n## Customizing retry behavior\nYou can customize the retry behavior when creating a client by modifying the \u0060RetryOptions\u0060 in the client options.\n### Example: customizing retry options","contentUrl":"https://learn.microsoft.com/azure/developer/cpp/sdk/fundamentals/http-pipelines-and-retries#retry-pol
... (truncated)
```

### ✅ 11. azure-documentation (via azure) — 2807ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure.Core API docs for RetryPolicy and RetryOptions retryable status code behavior and custom policies.","parameters":{"query":"site:learn.microsoft.com/dotnet/api/azure.core.pipeline RetryPolicy Azure.Core retry response classifier status codes 408 429 500 502 503 504 custom retry policy"}}
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

### ✅ 12. azure-documentation (via azure) — 4676ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find exact .NET Azure.Core RetryPolicy subclass override signatures and ClientOptions RetryPolicy property usage.","parameters":{"query":"Azure.Core RetryPolicy custom subclass ShouldRetry HttpMessage exception Response CalculateDelay .NET API ClientOptions RetryPolicy example"}}
```

**Output:**

```
{"results":[{"title":"RetryPolicy.ShouldRetry(HttpMessage, Exception) Method-Definition","content":"# RetryPolicy.ShouldRetry(HttpMessage, Exception) Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure.Core.Pipeline](https://learn.microsoft.com/dotnet/api/azure.core.pipeline?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.Core.dll\r\n\r\n- Package:\r\n    - Azure.Core v1.62.0\r\n\r\n- Source:\r\n    - [RetryPolicy.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs#L186C99-L186C138)\r\n\r\nThis method can be overridden to control whether a request should be retried. It will be called for any response where [IsError](https://learn.microsoft.com/dotnet/api/azure.response.iserror?view=azure-dotnet#azure-response-iserror) is true, or if an exception is thrown from any subsequent pipeline policies or the transport. This method will only be called for sync methods.\r\n\r\n\u0060\u0060\u0060csharp\r\nprotected internal virtual bool ShouldRetry(Azure.Core.HttpMessage message, Exception? exception);\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- message\r\n    - [HttpMessage](https://learn.microsoft.com/dotnet/api/azure.core.httpmessage?view=azure-dotnet)\r\n\r\nThe message containing the request and response.\r\n\r\n- exception\r\n    - [Exception](https://learn.microsoft.com/en-us/dotnet/api/system.exception)\r\n\r\nThe exception that occurred, if any, which can be used to determine if a retry should occur.\r\n\r\n#### Returns\r\n\r\n[Boolean](https://learn.microsoft.com/en-us/dotnet/api/system.boolean)\r\n\r\nWhether or not to retry.","contentUrl":"https://learn.microsoft.com/dotnet/api/azure.core.pipeline.retrypolicy.shouldretry?view=azure-dotnet"},{"title":"RetryPolicy.ShouldRetryAsync(HttpMessage, Exception) Method-Definition","content":"# RetryPolicy.ShouldRetryAsync(HttpMessage, Exception) Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure.Cor
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 4291ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Azure.Core RetryPolicy API documentation.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.retrypolicy"}}
```

**Output:**

```
# ResponseClassifier Class

## Definition

- Namespace:
    - [Azure.Core](azure.core?view=azure-dotnet)

- Assembly:
    - Azure.Core.dll

- Package:
    - Azure.Core v1.62.0

- Source:
    - [ResponseClassifier.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/ResponseClassifier.cs)

A type that analyzes HTTP responses and exceptions and determines if they should be retried, and/or analyzes responses and determines if they should be treated as error responses.

```csharp
public class ResponseClassifier
```

```fsharp
type ResponseClassifier = class
```

```vb
Public Class ResponseClassifier
```

- Inheritance
    - [Object](/en-us/dotnet/api/system.object)
ResponseClassifier

- Derived
    - [Azure.Core.StatusCodeClassifier](azure.core.statuscodeclassifier?view=azure-dotnet)

## Constructors

| Name | Description |
| --- | --- |
| [ResponseClassifier()](azure.core.responseclassifier.-ctor?view=azure-dotnet#azure-core-responseclassifier-ctor) |

## Methods

| Name | Description |
| --- | --- |
| [IsErrorResponse(HttpMessage)](azure.core.responseclassifier.iserrorresponse?view=azure-dotnet#azure-core-responseclassifier-iserrorresponse%28azure-core-httpmessage%29) | Specifies if the response contained in the `message` is not successful. |
| [IsRetriable(HttpMessage, Exception)](azure.core.responseclassifier.isretriable?view=azure-dotnet#azure-core-responseclassifier-isretriable%28azure-core-httpmessage-system-exception%29) | Specifies if the operation that caused the exception should be retried taking the [HttpMessage](azure.core.httpmessage?view=azure-dotnet) into consideration. |
| [IsRetriableException(Exception)](azure.core.responseclassifier.isretriableexception?view=azure-dotnet#azure-core-responseclassifier-isretriableexception%28system-exception%29) | Specifies if the operation that caused the exception should be retried. |
| [IsRetriableRespon
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 3095ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Azure.Core ResponseClassifier API documentation.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.responseclassifier"}}
```

**Output:**

```
# RetryPolicy Class

## Definition

- Namespace:
    - [Azure.Core.Pipeline](azure.core.pipeline?view=azure-dotnet)

- Assembly:
    - Azure.Core.dll

- Package:
    - Azure.Core v1.62.0

- Source:
    - [RetryPolicy.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs)

Represents a policy that can be overridden to customize whether or not a request will be retried and how long to wait before retrying.

```csharp
public class RetryPolicy : Azure.Core.Pipeline.HttpPipelinePolicy
```

```fsharp
type RetryPolicy = class
    inherit HttpPipelinePolicy
```

```vb
Public Class RetryPolicy
Inherits HttpPipelinePolicy
```

- Inheritance
    - [Object](/en-us/dotnet/api/system.object)

[HttpPipelinePolicy](azure.core.pipeline.httppipelinepolicy?view=azure-dotnet)
RetryPolicy

## Constructors

| Name | Description |
| --- | --- |
| [RetryPolicy(Int32, DelayStrategy)](azure.core.pipeline.retrypolicy.-ctor?view=azure-dotnet#azure-core-pipeline-retrypolicy-ctor%28system-int32-azure-core-delaystrategy%29) | Initializes a new instance of the [RetryPolicy](azure.core.pipeline.retrypolicy?view=azure-dotnet) class. |

## Methods

| Name | Description |
| --- | --- |
| [OnRequestSent(HttpMessage)](azure.core.pipeline.retrypolicy.onrequestsent?view=azure-dotnet#azure-core-pipeline-retrypolicy-onrequestsent%28azure-core-httpmessage%29) | This method can be overridden to introduce logic that runs after the request is sent through the pipeline and control is returned to the retry policy. This method will only be called for sync methods. |
| [OnRequestSentAsync(HttpMessage)](azure.core.pipeline.retrypolicy.onrequestsentasync?view=azure-dotnet#azure-core-pipeline-retrypolicy-onrequestsentasync%28azure-core-httpmessage%29) | This method can be overridden to introduce logic that runs after the request is sent through the pipeline and control is returned to 
... (truncated)
```

### ✅ 15. web_fetch — 3896ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/ResponseClassifier.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
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
        public virtual bool IsRetriable(HttpMessage
... (truncated)
```

### ✅ 16. web_fetch — 3493ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/ResponseClassifier.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
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
        /// <param
... (truncated)
```

### ✅ 17. web_fetch — 1185ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
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
                        return true
... (truncated)
```

### ✅ 18. web_fetch — 1227ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/DelayStrategy.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/DelayStrategy.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using Azure.Core.Pipeline;

namespace Azure.Core
{
    /// <summary>
    /// An abstraction to control delay behavior.
    /// </summary>
    public abstract class DelayStrategy
    {
        internal const double DefaultJitterFactor = 0.2;

        private readonly Random _random = new ThreadSafeRandom();
        private readonly double _minJitterFactor;
        private readonly double _maxJitterFactor;
        private readonly TimeSpan _maxDelay;

        /// <summary>
        /// Constructs a new instance of <see cref="DelayStrategy"/>. This constructor can be used by derived classes to customize the jitter factor and max delay.
        /// </summary>
        /// <param name="maxDelay">The max delay value to apply on an individual delay.</param>
        /// <param name="jitterFactor">The jitter factor to apply to each delay. For example, if the delay is 1 second with a jitterFactor of 0.2, the actual
        /// delay used will be a random double between 0.8 and 1.2. If set to 0, no jitter will be applied.</param>
        protected DelayStrategy(TimeSpan? maxDelay = default, double jitterFactor = DefaultJitterFactor)
        {
            _minJitterFactor = 1.0 - jitterFactor;
            _maxJitterFactor = 1.0 + jitterFactor;

            // use same defaults as RetryOptions
            _maxDelay = maxDelay ?? RetryOptions.DefaultMaxDelay;
        }

        /// <summary>
        /// Constructs an exponential delay with jitter.
        /// </summary>
        /// <param name="initialDelay">The initial delay to use.</param>
        /// <param name="maxDelay">The maximum delay to use.</param>
        /// <returns>The <see cref="DelayStrategy"/> 
... (truncated)
```

## Copilot Response

I’m grounding the examples in the current Azure SDK retry pipeline, including the distinction between built-in retries and per-operation custom policies.

The key design point is that `NetworkTimeout` limits each HTTP attempt, while a caller `CancellationToken` limits the whole logical upload—including SDK retries. I’m checking the custom `RetryPolicy` extension points so the high-value example compiles against current Azure.Core.

Configure retries once on reusable clients, then use a cancellation token to bound the **entire** upload. `NetworkTimeout` only limits an individual HTTP attempt; it does not cap the total operation.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

static BlobClientOptions CreateBlobOptions(
    int maxRetries,
    TimeSpan networkTimeout)
{
    return new BlobClientOptions
    {
        Retry =
        {
            Mode = RetryMode.Exponential,
            MaxRetries = maxRetries,
            Delay = TimeSpan.FromSeconds(1),
            MaxDelay = TimeSpan.FromSeconds(30),
            NetworkTimeout = networkTimeout
        }
    };
}

Uri serviceUri = new(
    Environment.GetEnvironmentVariable("BLOB_SERVICE_URI")
    ?? throw new InvalidOperationException("BLOB_SERVICE_URI is required."));

// Use DefaultAzureCredential for development. Prefer ManagedIdentityCredential
// for applications hosted in Azure.
TokenCredential credential = new DefaultAzureCredential();

BlobServiceClient blobServiceClient = new(
    serviceUri,
    credential,
    CreateBlobOptions(
        maxRetries: 5,
        networkTimeout: TimeSpan.FromSeconds(30)));
```

`MaxRetries = 5` means up to **five retries after the initial request**, for a maximum of six attempts. The exponential delay is based on `Delay`, capped by `MaxDelay`. A service-provided `Retry-After` header takes precedence.

## Bound the complete upload

Use a linked `CancellationTokenSource` so application shutdown and the operation deadline are both honored:

```csharp
static async Task UploadWithTimeoutAsync(
    BlobClient blobClient,
    Stream content,
    TimeSpan operationTimeout,
    CancellationToken applicationCancellation)
{
    using var timeoutCts =
        CancellationTokenSource.CreateLinkedTokenSource(applicationCancellation);

    timeoutCts.CancelAfter(operationTimeout);

    try
    {
        await blobClient.UploadAsync(
            content,
            overwrite: true,
            cancellationToken: timeoutCts.Token);
    }
    catch (OperationCanceledException)
        when (!applicationCancellation.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"Blob upload exceeded {operationTimeout}.");
    }
}
```

For example:

```csharp
BlobClient blob = blobServiceClient
    .GetBlobContainerClient("documents")
    .GetBlobClient("report.pdf");

await UploadWithTimeoutAsync(
    blob,
    File.OpenRead("report.pdf"),
    operationTimeout: TimeSpan.FromMinutes(2),
    applicationCancellation: stoppingToken);
```

For large uploads, `NetworkTimeout` applies independently to each block request. The cancellation token is therefore necessary to enforce an overall deadline.

## Transient versus authentication failures

Azure.Core retries these HTTP statuses by default:

| Status | Meaning |
|---:|---|
| `408` | Request Timeout |
| `429` | Too Many Requests |
| `500` | Internal Server Error |
| `502` | Bad Gateway |
| `503` | Service Unavailable |
| `504` | Gateway Timeout |

It also retries common transport failures, including `IOException`, `RequestFailedException` with status `0`, and cancellation caused internally by a network timeout rather than by the caller.

Blob Storage adds retries for Storage service error codes such as `InternalError`, `OperationTimedOut`, and `ServerBusy`. When `GeoRedundantSecondaryUri` is configured, a `404` from the secondary endpoint is also handled specially because replication might not have reached that endpoint.

Authentication and authorization failures—normally `401` and `403`—are **not retryable** by default. Retrying them generally cannot help until credentials or RBAC assignments change.

```csharp
using System.IO;
using System.Net.Http;

static bool IsTransient(Exception exception)
{
    return exception switch
    {
        RequestFailedException ex =>
            ex.Status is 0 or 408 or 429 or 500 or 502 or 503 or 504,

        IOException => true,
        HttpRequestException => true,

        AggregateException ex when ex.InnerExceptions.Count > 0 =>
            ex.InnerExceptions.All(IsTransient),

        _ => false
    };
}

try
{
    await UploadWithTimeoutAsync(
        blob,
        content,
        TimeSpan.FromMinutes(2),
        stoppingToken);
}
catch (RequestFailedException ex) when (ex.Status is 401 or 403)
{
    // Non-transient: alert on credential/RBAC configuration and do not retry.
    logger.LogError(ex, "Blob authentication or authorization failed.");
    throw;
}
catch (Exception ex) when (IsTransient(ex))
{
    // SDK retries are exhausted. Queue for later, degrade gracefully, or alert.
    logger.LogWarning(ex, "Transient Blob Storage failure after retries.");
    throw;
}
```

Do not treat every `RequestFailedException` as transient: statuses such as `400`, `401`, `403`, `404`, `409`, and `412` usually require an application or concurrency decision rather than another identical attempt.

## A stronger policy for high-value operations

Azure SDK clients are thread-safe and intended to be reused. Create a second client with separate `BlobClientOptions` rather than changing a shared client per request:

```csharp
BlobServiceClient highValueBlobServiceClient = new(
    serviceUri,
    credential,
    CreateBlobOptions(
        maxRetries: 8,
        networkTimeout: TimeSpan.FromSeconds(45)));

BlobClient highValueBlob = highValueBlobServiceClient
    .GetBlobContainerClient("critical-documents")
    .GetBlobClient("settlement.dat");
```

This client can use a larger retry budget and operation timeout without slowing ordinary requests. Keep total latency in mind: eight retries plus exponential delays can make a request take several minutes.

If you need to change the SDK’s actual retry classification rather than merely its limits, `BlobClientOptions.RetryPolicy` accepts a custom `Azure.Core.Pipeline.RetryPolicy`. Setting it replaces the normal retry policy; for most applications, separate clients with tailored `RetryOptions` are safer and easier to maintain.

## Circuit breaker for sustained failures

Use a circuit breaker **outside** the SDK call. The Blob SDK performs short-term retries; the breaker observes only failures that remain after those retries and temporarily rejects new calls. With Polly v8:

```powershell
dotnet add package Polly
```

```csharp
using Polly;
using Polly.CircuitBreaker;

ResiliencePipeline blobCircuitBreaker =
    new ResiliencePipelineBuilder()
        .AddCircuitBreaker(new CircuitBreakerStrategyOptions
        {
            ShouldHandle = new PredicateBuilder()
                .Handle<Exception>(IsTransient),

            FailureRatio = 0.5,
            SamplingDuration = TimeSpan.FromSeconds(30),
            MinimumThroughput = 10,
            BreakDuration = TimeSpan.FromSeconds(30)
        })
        .Build();

await blobCircuitBreaker.ExecuteAsync(
    async cancellationToken =>
    {
        await UploadWithTimeoutAsync(
            highValueBlob,
            content,
            operationTimeout: TimeSpan.FromMinutes(5),
            applicationCancellation: cancellationToken);
    },
    stoppingToken);
```

Because `IsTransient` excludes `401` and `403`, authentication failures do not open the circuit. Do not add another general-purpose retry strategy around this pipeline unless deliberately coordinated with the SDK retry count—stacked retry layers can multiply attempts and overload an already unhealthy service.

**References:** [Configure retry policies for Blob Storage](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), [Azure.Core `RetryPolicy`](https://learn.microsoft.com/dotnet/api/azure.core.pipeline.retrypolicy), and [Azure.Core `ResponseClassifier`](https://learn.microsoft.com/dotnet/api/azure.core.responseclassifier).

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The response thoroughly covers SDK retry configuration, timeout control, error classification, operation-specific clients, and circuit breaking; geo-redundant retry configuration is incomplete.

### Strengths

- Correctly distinguishes per-request NetworkTimeout from an overall CancellationToken deadline.
- Accurately lists the required retryable and non-retryable status codes.
- Demonstrates tailored clients for high-value operations and a coordinated Polly circuit breaker.

### Issues

- No concrete BlobClientOptions configuration for GeoRedundantSecondaryUri or adequate explanation of endpoint failover behavior.

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
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
