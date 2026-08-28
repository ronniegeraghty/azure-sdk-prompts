# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 143.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 143.5s |
| Timestamp | 2026-08-28T17:09:47Z |
| Files Generated | 0 |
| Event Count | 4323 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 115.4s |
| Review | 26.9s |
| **Total** | **143.5s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-resource-manager-cosmosdb-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=351847 out=4368 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 633ms

**Input:**

```json
{"skill":"azure-resource-manager-cosmosdb-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure.ResourceManager.CosmosDB (.NET)

Management plane SDK for provisioning and managing Azure Cosmos DB resources via Azure Resource Manager.

> **⚠️ Management vs Data Plane**
> - **This SDK (Azure.ResourceManager.CosmosDB)**: Create accounts, databases, containers, configure throughput, manage RBAC
> - **Data Plane SDK (Microsoft.Azure.Cosmos)**: CRUD operations on documents, queries, stored procedures execution

## Installation

```bash
dotnet add package Azure.ResourceManager.CosmosDB
dotnet add package Azure.Identity
```

**Current Versions**: Stable v1.4.0, Preview v1.4.0-beta.13

## Environment Variables

```bash
AZURE_SUBSCRIPTION_ID=<your-subscription-id> # Required: Azure subscription ID
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
AZURE_TENANT_ID=<tenant-id> # For service principal auth (optional)
AZURE_CLIENT_ID=<client-id> # For service principal auth (optional)
AZURE_CLIENT_SECRET=<client-secret> # For service principal auth (optional)
```

## Authentication

```csharp
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
var credential = new DefaultAzureCredential(
    DefaultAzureCredential.DefaultEnvironmentVariableName
);
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet#credential-classes
// var credential = new ManagedIdentityCredential();
var armClient = new ArmClient(credential);

// Get subscription
var subscriptionId = Environment.GetEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
var subscription = armClient.GetSubscriptionResource(
    new ResourceIdentifier($"/subscriptions/{subscriptionId}"));
```

## Resource Hierarchy

```
ArmClient
└── SubscriptionResource
    └── ResourceGro
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 3114ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Provide current best practices for .NET code using Microsoft.Azure.Cosmos SDK v3 data-plane error handling, especially CosmosException 429 retry behavior, 404 not found, and 409 conflict."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 3754ms

**Input:**

```json
{"command":"search","intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos SDK v3 CosmosException properties StatusCode and RetryAfter, 429 throttling retry configuration, and handling 404/409 responses.","parameters":{"query":"Microsoft.Azure.Cosmos SDK v3 CosmosException RetryAfter StatusCode 429 MaxRetryAttemptsOnRateLimitedRequests 404 409"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 3432ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-resource-manager-cosmosdb-dotnet","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 2526ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices before providing Microsoft.Azure.Cosmos SDK v3 error-handling examples.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 3168ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Search official docs for Cosmos DB .NET SDK v3 exception and throttling behavior.","parameters":{"query":"Microsoft.Azure.Cosmos .NET SDK v3 CosmosException RetryAfter StatusCode 429 MaxRetryAttemptsOnRateLimitedRequests MaxRetryWaitTimeOnRateLimitedRequests 404 409"}}
```

**Output:**

```
{"results":[{"description":"description: Retrieves an item asynchronously from a container and accesses diagnostics information such as elapsed time and contacted regions, with exception handling for CosmosException in the Azure Cosmos DB .NET SDK v3.\npackage: Microsoft.Azure.Cosmos\nlanguage: csharp\n","codeSnippet":"catch (CosmosException ex)\r\n{\r\n    HttpStatusCode statusCode = ex.StatusCode;\r\n    CosmosDiagnostics diagnostics = ex.Diagnostics;\r\n    // store diagnostics optionally with diagnostics.ToString();\r\n    // or log the entire error details with ex.ToString();\r\n}","link":"https://learn.microsoft.com/azure/cosmos-db/migrate-dotnet-v3#major-name-changes-from-v2-sdk-to-v3-sdk","language":"csharp"},{"description":"description: Handles rate limiting by retrying a ReadItemAsync operation on a Cosmos container while excluding a specific region to route requests to alternate regions when a 429 Too Many Requests response is received.\npackage: Microsoft.Azure.Cosmos\nlanguage: csharp\n","codeSnippet":"ItemResponse\u003CCosmosItem\u003E item;\r\nitem = await container.ReadItemAsync\u003CCosmosItem\u003E(\u0022id\u0022, partitionKey);\r\n\r\nif (item.StatusCode == HttpStatusCode.TooManyRequests)\r\n{\r\n    ItemRequestOptions requestOptions = new ItemRequestOptions()\r\n    {\r\n        ExcludeRegions = new List\u003Cstring\u003E() { \u0022East US\u0022 }\r\n    };\r\n\r\n    item = await container.ReadItemAsync\u003CCosmosItem\u003E(\u0022id\u0022, partitionKey, requestOptions);\r\n}","link":"https://learn.microsoft.com/azure/cosmos-db/performance-tips-dotnet-sdk-v3#high-availability","language":"csharp"},{"description":"description: Demonstrates reading an item asynchronously from a container and accessing diagnostics information such as elapsed time and contacted regions, with error handling for CosmosException in the Azure Cosmos DB .NET SDK v3.\npackage: Microsoft.Azure.Cosmos\nlanguage: csharp\n","codeSnippet":"catch (CosmosException ex)\r\n{\r\n  
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 2845ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# samples for handling CosmosException throttling and status codes.","parameters":{"language":"csharp","query":"Microsoft.Azure.Cosmos CosmosException catch StatusCode RetryAfter TooManyRequests NotFound Conflict retry"}}
```

**Output:**

```
{"results":[{"title":"CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property-Definition","content":"# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Microsoft.Azure.Cosmos](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Microsoft.Azure.Cosmos.Client.dll\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.58.0\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.59.0-preview.0\r\n\r\n- Source:\r\n    - [CosmosClientOptions.cs](https://github.com/Azure/azure-cosmos-dotnet-v3/blob/5e6232d25f84074d5008e97670d3e0dc07891f0e/Microsoft.Azure.Cosmos/src/CosmosClientOptions.cs#L354C61-L354C65)\r\n\r\nGets or sets the maximum number of retries in the case where the request fails because the Azure Cosmos DB service has applied rate limiting on the client.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic int? MaxRetryAttemptsOnRateLimitedRequests { get; set; }\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n[Nullable](https://learn.microsoft.com/en-us/dotnet/api/system.nullable-1)\u0026lt;[Int32](https://learn.microsoft.com/en-us/dotnet/api/system.int32)\u0026gt; \r\n\r\nThe default value is 9. This means in the case where the request is rate limited, the same request will be issued for a maximum of 10 times to the server before an error is returned to the application.\r\n\r\nIf the value of this property is set to 0, there will be no automatic retry on rate limiting requests from the client and the exception needs to be handled at the application level.\r\n\r\n## Remarks\r\n\r\nWhen a client is sending requests faster than the allowed rate, the service will return HttpStatusCode 429 (Too Many Requests) to rate limit the client. The current implementation in the SDK will then wait for the amount of time the service tells it to wait and retry after the time has elapsed.\r\n\r\nFor more information, see [Handle rate limiting/request rat
... (truncated)
```

### ✅ 8. rg — 746ms

**Input:**

```json
{"head_limit":40,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt","pattern":"\"title\":\"[^\"]+\"|https://learn\\.microsoft\\.com[^\"\\\\]+"}
```

**Output:**

```
Output too large to read at once (52.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787937045228-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt:{"results":[{"title":"CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property-Definition","content":"# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Microsoft.Azure.Cosmos](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Microsoft.Azure.Cosmos.Client
```

### ✅ 9. powershell — 1195ms

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt' | ConvertFrom-Json; $j.results | Select-Object title,contentUrl | Format-Table -AutoSize","description":"List Cosmos documentation search results"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-13886096c26c4d36a7fc88d9d980b212.txt' | ConvertFrom-Json; $j.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 10. view — 656ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt b/C:/Users/weidxu/AppData/Local/Temp/1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787937039860-copilot-tool-output-13886096c26c4d36a7fc88d9d980b212.txt
@@ -1,6 +1,6 @@
 {"results":[{"title":"CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property-Definition","content":"# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Microsoft.Azure.Cosmos](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Microsoft.Azure.Cosmos.Client.dll\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.58.0\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.59.0-preview.0\r\n\r\n- Source:\r\n    - [CosmosClientOptions.cs](https://github.com/Azure/azure-cosmos-dotnet-v3/blob/5e6232d25f84074d5008e97670d3e0dc07891f0e/Microsoft.Azure.Cosmos/src/CosmosClientOptions.cs#L354C61-L354C65)\r\n\r\nGets or sets the maximum number of retries in the case where the request fails because the Azure Cosmos DB service has applied rate limiting on the client.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic int? MaxRetryAttemptsOnRateLimitedRequests { get; set; }\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n[Nullable](https://learn.microsoft.com/en-us/dotnet/api/system.nullable-1)\u0026lt;[Int32](https://learn.microsoft.com/en-us/dotnet/api/system.int32)\u0026gt; \r\n\r\nThe default value is 9. This means in the case where the request is rate limited, the same request will be issued for a maximum of 10 times to the server before an error is returned to the application.\r\n\r\nIf the value of this property is set to 0, there will be no automatic 
... (truncated)
```

### ✅ 11. azure-documentation (via azure) — 3389ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete official guidance on retryable Cosmos DB status codes.","parameters":{"url":"https://learn.microsoft.com/azure/cosmos-db/conceptual-resilient-sdk-applications"}}
```

**Output:**

```
# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property

## Definition

- Namespace:
    - [Microsoft.Azure.Cosmos](microsoft.azure.cosmos?view=azure-dotnet)

- Assembly:
    - Microsoft.Azure.Cosmos.Client.dll

- Package:
    - Microsoft.Azure.Cosmos v3.58.0

- Package:
    - Microsoft.Azure.Cosmos v3.59.0-preview.0

- Source:
    - [CosmosClientOptions.cs](https://github.com/Azure/azure-cosmos-dotnet-v3/blob/5e6232d25f84074d5008e97670d3e0dc07891f0e/Microsoft.Azure.Cosmos/src/CosmosClientOptions.cs#L354C61-L354C65)

::: moniker range=" azure-dotnet azure-dotnet-preview "

Gets or sets the maximum number of retries in the case where the request fails because the Azure Cosmos DB service has applied rate limiting on the client.

```csharp
public int? MaxRetryAttemptsOnRateLimitedRequests { get; set; }
```

```fsharp
member this.MaxRetryAttemptsOnRateLimitedRequests : Nullable<int> with get, set
```

```vb
Public Property MaxRetryAttemptsOnRateLimitedRequests As Nullable(Of Integer)
```

#### Property Value

[Nullable](/en-us/dotnet/api/system.nullable-1)&lt;[Int32](/en-us/dotnet/api/system.int32)&gt; 

The default value is 9. This means in the case where the request is rate limited, the same request will be issued for a maximum of 10 times to the server before an error is returned to the application.

If the value of this property is set to 0, there will be no automatic retry on rate limiting requests from the client and the exception needs to be handled at the application level.

## Remarks

When a client is sending requests faster than the allowed rate, the service will return HttpStatusCode 429 (Too Many Requests) to rate limit the client. The current implementation in the SDK will then wait for the amount of time the service tells it to wait and retry after the time has elapsed.

For more information, see [Handle rate limiting/request rate too large](/en-us/azure/cosmos-db/performance-tips#throughput).

... (truncated)
```

### ✅ 12. azure-documentation (via azure) — 3069ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Cosmos DB .NET SDK v3 throughput and retry behavior.","parameters":{"url":"https://learn.microsoft.com/azure/cosmos-db/performance-tips-dotnet-sdk-v3#throughput"}}
```

**Output:**

```
# CosmosException.RetryAfter Property

## Definition

- Namespace:
    - [Microsoft.Azure.Cosmos](microsoft.azure.cosmos?view=azure-dotnet)

- Assembly:
    - Microsoft.Azure.Cosmos.Client.dll

- Package:
    - Microsoft.Azure.Cosmos v3.58.0

- Package:
    - Microsoft.Azure.Cosmos v3.59.0-preview.0

- Source:
    - [CosmosException.cs](https://github.com/Azure/azure-cosmos-dotnet-v3/blob/5e6232d25f84074d5008e97670d3e0dc07891f0e/Microsoft.Azure.Cosmos/src/Resource/CosmosExceptions/CosmosException.cs#L121C48-L121C71)

::: moniker range=" azure-dotnet azure-dotnet-preview "

Gets the retry after time. This tells how long a request should wait before doing a retry.

```csharp
public virtual TimeSpan? RetryAfter { get; }
```

```fsharp
member this.RetryAfter : Nullable<TimeSpan>
```

```vb
Public Overridable ReadOnly Property RetryAfter As Nullable(Of TimeSpan)
```

#### Property Value

[Nullable](/en-us/dotnet/api/system.nullable-1)&lt;[TimeSpan](/en-us/dotnet/api/system.timespan)&gt; 

## Applies to

::: moniker-end
```

### ✅ 13. azure-documentation (via azure) — 2759ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official API reference for CosmosClientOptions throttling retry count.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests?view=azure-dotnet"}}
```

**Output:**

```
# Performance tips for Azure Cosmos DB and .NET

- [.NET SDK v3](performance-tips-dotnet-sdk-v3)
- [.NET SDK v2](performance-tips)
- [Java SDK v4](performance-tips-java-sdk-v4)
- [Async Java SDK v2](performance-tips-async-java)
- [Sync Java SDK v2](performance-tips-java)
- [Python SDK](performance-tips-python-sdk)

Azure Cosmos DB is a fast, flexible distributed database that scales seamlessly with guaranteed latency and throughput levels. You don't have to make major architecture changes or write complex code to scale your database with Azure Cosmos DB. Scaling up and down is as easy as making a single API call. To learn more, see [provision container throughput](how-to-provision-container-throughput) or [provision database throughput](how-to-provision-database-throughput).

Because Azure Cosmos DB is accessed via network calls, you can make client-side optimizations to achieve peak performance when you use the [SQL .NET SDK](sdk-dotnet-v3).

If you're trying to improve your database performance, consider the options presented in the following sections.

Tip

Agent Kit helps coding agents work with Azure Cosmos DB quickly and efficiently using recommended best practices. To get started, run:

```bash
npx skills add AzureCosmosDB/cosmosdb-agent-kit
```

To learn more, see [Azure Cosmos DB Agent Kit](gen-ai/agent-kit).

## Hosting recommendations

**Turn on server-side garbage collection**

Reducing the frequency of garbage collection can help in some cases. In .NET, set [gcServer](/en-us/dotnet/core/run-time-config/garbage-collector#flavors-of-garbage-collection) to `true`.

**Scale out your client workload**

If you're testing at high throughput levels, or at rates that are greater than 50,000 Request Units per second (RU/s), the client application could become a workload bottleneck because the machine might cap out on CPU or network utilization. If you reach this point, you can continue to push the Azure Cosmos DB account further b
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 2439ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official API reference for CosmosException RetryAfter.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet"}}
```

**Output:**

```
# Design resilient applications with Azure Cosmos DB SDKs

When authoring client applications that interact with Azure Cosmos DB through any of the SDKs, it's important to understand a few key fundamentals. This article is a design guide to help you understand these fundamentals and design resilient applications.

For a video overview of the concepts discussed in this article, see:

## Connectivity modes

Azure Cosmos DB SDKs can connect to the service in two [connectivity modes](sdk-connection-modes). The .NET and Java SDKs can connect to the service in either *Gateway* or *Direct* mode, while the others can only connect in Gateway mode. Gateway mode uses the HTTP protocol, and Direct mode typically uses the TCP protocol.

Gateway mode is always used to fetch metadata such as the account, container, and routing information regardless of which mode SDK is configured for use. This information is cached in memory and is used to connect to the [service replicas](partitioning#replica-sets).

In summary, for SDKs in Gateway mode, you can expect HTTP traffic, while for SDKs in Direct mode, you can expect a combination of HTTP and TCP traffic under different circumstances, such as initialization, fetching metadata, or routing information.

Note

Direct mode is available only in the .NET and Java SDKs. All other Azure Cosmos DB SDKs use Gateway mode. The error status codes and retry guidance in this article apply across all SDKs. For more information, see [connectivity modes](sdk-connection-modes).

## Client instances and connections

Regardless of the connectivity mode, it's critical to maintain a singleton instance of the SDK client per account per application. Connections, both HTTP and TCP, are scoped to the client instance. Most compute environments have limitations in terms of the number of connections that can be open at the same time. When these limits are reached, connectivity is affected.

## Distributed applications and networks

When you
... (truncated)
```

## Copilot Response

I’m grounding the example in the v3 data-plane SDK’s built-in throttling behavior, then I’ll show a manual retry variant that avoids stacked retries.

The Cosmos DB .NET SDK v3 already retries HTTP 429 responses by default and honors the service-provided `x-ms-retry-after-ms` header. Configure that behavior on a singleton `CosmosClient`; only add application-level retries if you need a retry budget beyond the SDK’s policy.

```bash
dotnet add package Microsoft.Azure.Cosmos --version 3.58.0
```

## Recommended: use the SDK retry policy

```csharp
using Microsoft.Azure.Cosmos;

CosmosClientOptions options = new()
{
    // Defaults shown explicitly:
    // 9 retries means up to 10 total attempts.
    MaxRetryAttemptsOnRateLimitedRequests = 9,

    // Stops retrying once cumulative throttling delays reach this limit.
    MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30)
};

// Keep one CosmosClient instance for the application's lifetime.
CosmosClient client = new(
    Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")
        ?? throw new InvalidOperationException("COSMOS_ENDPOINT is not configured."),
    Environment.GetEnvironmentVariable("COSMOS_KEY")
        ?? throw new InvalidOperationException("COSMOS_KEY is not configured."),
    options);

Container container = client.GetContainer("application-db", "items");
```

If a 429 reaches your code, the SDK has already exhausted the configured retry policy:

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;

try
{
    ItemResponse<MyItem> response = await container.ReadItemAsync<MyItem>(
        id: itemId,
        partitionKey: new PartitionKey(partitionKey),
        cancellationToken: cancellationToken);

    MyItem item = response.Resource;
}
catch (CosmosException ex)
{
    HttpStatusCode statusCode = ex.StatusCode;

    // This is the parsed value of the x-ms-retry-after-ms response header.
    TimeSpan? retryAfter = ex.RetryAfter;

    Console.Error.WriteLine(
        "Cosmos DB request failed. Status={0}, SubStatus={1}, " +
        "RetryAfter={2}, ActivityId={3}, RequestCharge={4}",
        (int)statusCode,
        ex.SubStatusCode,
        retryAfter,
        ex.ActivityId,
        ex.RequestCharge);

    throw;
}
```

Also log `ex.Diagnostics` for slow or failed operations, but avoid logging item content or credentials.

## Manual 429 retry policy

If you implement retries yourself, disable the SDK’s 429 retries so the two policies do not multiply each other:

```csharp
CosmosClientOptions options = new()
{
    MaxRetryAttemptsOnRateLimitedRequests = 0
};
```

A bounded retry helper that honors `RetryAfter`, supports cancellation, and uses exponential backoff with jitter when the header is unavailable:

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Logging;

public static class CosmosRetry
{
    public static async Task<T> ExecuteAsync<T>(
        Func<CancellationToken, Task<T>> operation,
        ILogger logger,
        int maxAttempts,
        CancellationToken cancellationToken = default)
    {
        ArgumentNullException.ThrowIfNull(operation);
        ArgumentNullException.ThrowIfNull(logger);

        if (maxAttempts < 1)
        {
            throw new ArgumentOutOfRangeException(nameof(maxAttempts));
        }

        for (int attempt = 1; ; attempt++)
        {
            cancellationToken.ThrowIfCancellationRequested();

            try
            {
                return await operation(cancellationToken);
            }
            catch (CosmosException ex) when (
                ex.StatusCode == HttpStatusCode.TooManyRequests &&
                attempt < maxAttempts)
            {
                TimeSpan retryAfter = ex.RetryAfter.GetValueOrDefault();

                TimeSpan backoff = retryAfter > TimeSpan.Zero
                    ? retryAfter
                    : TimeSpan.FromMilliseconds(
                        Math.Min(30_000, 200 * Math.Pow(2, attempt - 1)));

                TimeSpan jitter =
                    TimeSpan.FromMilliseconds(Random.Shared.Next(0, 250));

                TimeSpan delay = backoff + jitter;

                logger.LogWarning(
                    "Cosmos DB throttled request. Status={StatusCode}, " +
                    "Attempt={Attempt}/{MaxAttempts}, RetryAfter={RetryAfter}, " +
                    "Delay={Delay}, ActivityId={ActivityId}, RequestCharge={RequestCharge}",
                    (int)ex.StatusCode,
                    attempt,
                    maxAttempts,
                    ex.RetryAfter,
                    delay,
                    ex.ActivityId,
                    ex.RequestCharge);

                await Task.Delay(delay, cancellationToken);
            }
            catch (CosmosException ex)
            {
                logger.LogError(
                    ex,
                    "Cosmos DB request failed. Status={StatusCode}, " +
                    "SubStatus={SubStatusCode}, ActivityId={ActivityId}, " +
                    "RequestCharge={RequestCharge}, Diagnostics={Diagnostics}",
                    (int)ex.StatusCode,
                    ex.SubStatusCode,
                    ex.ActivityId,
                    ex.RequestCharge,
                    ex.Diagnostics);

                throw;
            }
        }
    }
}
```

Usage:

```csharp
ItemResponse<MyItem> response = await CosmosRetry.ExecuteAsync(
    ct => container.ReadItemAsync<MyItem>(
        itemId,
        new PartitionKey(partitionKey),
        cancellationToken: ct),
    logger,
    maxAttempts: 5,
    cancellationToken);
```

Once `maxAttempts` is reached, the final `CosmosException` is rethrown. A persistent stream of 429s usually indicates a hot partition, insufficient RU/s, or excessive concurrency; extending retries alone can substantially increase tail latency.

## Handle 404 without retrying

A 404 is normally a business outcome, not a transient failure:

```csharp
public static async Task<MyItem?> FindItemAsync(
    Container container,
    string itemId,
    string partitionKey,
    CancellationToken cancellationToken)
{
    try
    {
        ItemResponse<MyItem> response =
            await container.ReadItemAsync<MyItem>(
                itemId,
                new PartitionKey(partitionKey),
                cancellationToken: cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.NotFound)
    {
        return null;
    }
}
```

Confirm that both the item ID and partition-key value are correct before treating a 404 as an absent item.

## Handle 409 without retrying

A 409 commonly means an item with the same ID and partition key already exists, or a unique-key constraint was violated:

```csharp
public static async Task<bool> TryCreateItemAsync(
    Container container,
    MyItem item,
    string partitionKey,
    ILogger logger,
    CancellationToken cancellationToken)
{
    try
    {
        await container.CreateItemAsync(
            item,
            new PartitionKey(partitionKey),
            cancellationToken: cancellationToken);

        return true;
    }
    catch (CosmosException ex)
        when (ex.StatusCode == HttpStatusCode.Conflict)
    {
        logger.LogInformation(
            "Cosmos DB item already exists or violates a unique key. " +
            "ItemId={ItemId}, ActivityId={ActivityId}",
            item.Id,
            ex.ActivityId);

        return false;
    }
}
```

Do not blindly retry a 409. Resolve it according to application semantics:

- Use `UpsertItemAsync` when create-or-replace behavior is intentional.
- Read the existing item when a duplicate create should be treated as idempotent success.
- Return a conflict response when uniqueness is part of the business rules.
- After an ambiguous write timeout, read the item to determine whether the original create succeeded.

**References:** [Resilient Cosmos DB SDK applications](https://learn.microsoft.com/azure/cosmos-db/conceptual-resilient-sdk-applications), [Cosmos DB .NET SDK performance and throttling](https://learn.microsoft.com/azure/cosmos-db/performance-tips-dotnet-sdk-v3#throughput), [`CosmosException.RetryAfter`](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet), and [`MaxRetryAttemptsOnRateLimitedRequests`](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests?view=azure-dotnet).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response covers the relevant exception types, status handling, SDK retry configuration, and diagnostics, but its manual retry helper has a RetryAfter type error and it does not track RU charge from successful responses.

### Strengths

- Uses status-filtered CosmosException catches for 429, 404, and 409.
- Explains the SDK's built-in throttling policy and avoids multiplying SDK and application retry policies.
- Implements bounded cancellation-aware backoff with jitter conceptually and surfaces useful activity, substatus, and diagnostics data.

### Issues

- The manual retry helper does not compile because CosmosException.RetryAfter is a TimeSpan, not a nullable TimeSpan with GetValueOrDefault().
- RU consumption is shown only on exceptions; successful ItemResponse.RequestCharge is never captured or logged.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `CosmosException` as the primary exception type: Pass
      - Extracting `StatusCode`, `SubStatusCode`, and `RetryAfter` properties: Pass
      - Handling 429 (TooManyRequests) with retry-after backoff: Fail
      - Handling 404 (NotFound) for missing items or containers: Pass
      - Handling 409 (Conflict) for duplicate items: Pass
      - `CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests` configuration: Pass
      - RU consumption tracking via `RequestCharge` on responses: Fail
      - Diagnostics string for troubleshooting: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.7500** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
