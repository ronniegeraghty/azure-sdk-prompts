# Evaluation Report: cosmos-db-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 173.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 6/8 |
| Duration | 173.5s |
| Timestamp | 2026-08-28T17:06:53Z |
| Files Generated | 0 |
| Event Count | 4407 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 149.2s |
| Review | 23.1s |
| **Total** | **173.5s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=335547 out=4132 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. azure-get_azure_bestpractices (via azure) — 3341ms

**Input:**

```json
{"intent":"Provide modern Microsoft.Azure.Cosmos SDK v3 C# error handling examples for 429 throttling, RetryAfter-based retries, 404 not found, and 409 conflict without provisioning resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 5688ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Cosmos DB .NET SDK v3 error-handling and retry best practices.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 7148ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos SDK v3 CosmosException, RetryAfter, 429 retries, 404, and 409 handling.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. azure-documentation (via azure) — 4344ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Cosmos DB .NET SDK v3 guidance for 429 throttling retries and CosmosException status codes.","parameters":{"query":"Azure Cosmos DB .NET SDK v3 CosmosException StatusCode RetryAfter 429 404 409 retry configuration MaxRetryAttemptsOnRateLimitedRequests MaxRetryWaitTimeOnRateLimitedRequests"}}
```

**Output:**

```
{"results":[{"title":"CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property-Definition","content":"# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Microsoft.Azure.Cosmos](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Microsoft.Azure.Cosmos.Client.dll\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.58.0\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.59.0-preview.0\r\n\r\n- Source:\r\n    - [CosmosClientOptions.cs](https://github.com/Azure/azure-cosmos-dotnet-v3/blob/5e6232d25f84074d5008e97670d3e0dc07891f0e/Microsoft.Azure.Cosmos/src/CosmosClientOptions.cs#L354C61-L354C65)\r\n\r\nGets or sets the maximum number of retries in the case where the request fails because the Azure Cosmos DB service has applied rate limiting on the client.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic int? MaxRetryAttemptsOnRateLimitedRequests { get; set; }\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n[Nullable](https://learn.microsoft.com/en-us/dotnet/api/system.nullable-1)\u0026lt;[Int32](https://learn.microsoft.com/en-us/dotnet/api/system.int32)\u0026gt; \r\n\r\nThe default value is 9. This means in the case where the request is rate limited, the same request will be issued for a maximum of 10 times to the server before an error is returned to the application.\r\n\r\nIf the value of this property is set to 0, there will be no automatic retry on rate limiting requests from the client and the exception needs to be handled at the application level.\r\n\r\n## Remarks\r\n\r\nWhen a client is sending requests faster than the allowed rate, the service will return HttpStatusCode 429 (Too Many Requests) to rate limit the client. The current implementation in the SDK will then wait for the amount of time the service tells it to wait and retry after the time has elapsed.\r\n\r\nFor more information, see [Handle rate limiting/request rat
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 4013ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# samples for CosmosException handling and retry behavior in Microsoft.Azure.Cosmos v3.","parameters":{"language":"csharp","query":"Microsoft.Azure.Cosmos v3 C# CosmosException StatusCode RetryAfter TooManyRequests NotFound Conflict retry"}}
```

**Output:**

```
{"results":[{"description":"description: Retrieves an item asynchronously from a container and accesses diagnostics information such as elapsed time and contacted regions, with exception handling for CosmosException in the Azure Cosmos DB .NET SDK v3.\npackage: Microsoft.Azure.Cosmos\nlanguage: csharp\n","codeSnippet":"catch (CosmosException ex)\r\n{\r\n    HttpStatusCode statusCode = ex.StatusCode;\r\n    CosmosDiagnostics diagnostics = ex.Diagnostics;\r\n    // store diagnostics optionally with diagnostics.ToString();\r\n    // or log the entire error details with ex.ToString();\r\n}","link":"https://learn.microsoft.com/azure/cosmos-db/migrate-dotnet-v3#major-name-changes-from-v2-sdk-to-v3-sdk","language":"csharp"},{"description":"description: Demonstrates reading an item asynchronously from a container and accessing diagnostics information such as elapsed time and contacted regions, with error handling for CosmosException in the Azure Cosmos DB .NET SDK v3.\npackage: Microsoft.Azure.Cosmos\nlanguage: csharp\n","codeSnippet":"catch (CosmosException ex)\r\n{\r\n    HttpStatusCode statusCode = ex.StatusCode;\r\n    CosmosDiagnostics diagnostics = ex.Diagnostics;\r\n    // store diagnostics optionally with diagnostics.ToString();\r\n    // or log the entire error details with ex.ToString();\r\n}","link":"https://learn.microsoft.com/azure/cosmos-db/migrate-dotnet-v3#major-name-changes-from-v2-sdk-to-v3-sdk","language":"csharp"},{"description":"description: Catches and handles CosmosException to access status code and diagnostics information from Azure Cosmos DB operations using the .NET SDK v3.\npackage: Microsoft.Azure.Cosmos\nlanguage: csharp\n","codeSnippet":"CosmosClientBuilder cosmosClientBuilder = new CosmosClientBuilder(\r\n    accountEndpoint: \u0022https://testcosmos.documents.azure.com:443/\u0022,\r\n    authKeyOrResourceToken: \u0022SuperSecretKey\u0022)\r\n.WithConsistencyLevel(ConsistencyLevel.Strong)\r\n.WithApplicationRegion(Regions.EastUS);\r\nCosmosCl
... (truncated)
```

### ✅ 6. rg — 761ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt","pattern":"\"title\"|\"url\"|RetryAfter|MaxRetry|429|404|409"}
```

**Output:**

```
Output too large to read at once (49.8 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787936901216-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt:{"results":[{"title":"CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property-Definition","content":"# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Microsoft.Azure.Cosmos](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Microsoft.Azure.Cosmos.Client
```

### ✅ 7. powershell — 1199ms

**Input:**

```json
{"command":"$data = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt' | ConvertFrom-Json; $data.results | Select-Object title,url | Format-Table -AutoSize","description":"List Cosmos documentation search results"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-6b3b0beddbd04dd2aafd35b7a5854397.txt' | ConvertFrom-Json; $data. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 8. view — 3947ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt b/C:/Users/weidxu/AppData/Local/Temp/1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787936896996-copilot-tool-output-6b3b0beddbd04dd2aafd35b7a5854397.txt
@@ -1,8 +1,8 @@
 {"results":[{"title":"CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property-Definition","content":"# CosmosClientOptions.MaxRetryAttemptsOnRateLimitedRequests Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Microsoft.Azure.Cosmos](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Microsoft.Azure.Cosmos.Client.dll\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.58.0\r\n\r\n- Package:\r\n    - Microsoft.Azure.Cosmos v3.59.0-preview.0\r\n\r\n- Source:\r\n    - [CosmosClientOptions.cs](https://github.com/Azure/azure-cosmos-dotnet-v3/blob/5e6232d25f84074d5008e97670d3e0dc07891f0e/Microsoft.Azure.Cosmos/src/CosmosClientOptions.cs#L354C61-L354C65)\r\n\r\nGets or sets the maximum number of retries in the case where the request fails because the Azure Cosmos DB service has applied rate limiting on the client.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic int? MaxRetryAttemptsOnRateLimitedRequests { get; set; }\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n[Nullable](https://learn.microsoft.com/en-us/dotnet/api/system.nullable-1)\u0026lt;[Int32](https://learn.microsoft.com/en-us/dotnet/api/system.int32)\u0026gt; \r\n\r\nThe default value is 9. This means in the case where the request is rate limited, the same request will be issued for a maximum of 10 times to the server before an error is returned to the application.\r\n\r\nIf the value of this property is set to 0, there will be no automatic 
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 6520ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official retry guidance for Cosmos DB SDK applications.","parameters":{"url":"https://learn.microsoft.com/azure/cosmos-db/conceptual-resilient-sdk-applications"}}
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

### ✅ 10. azure-documentation (via azure) — 6197ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete CosmosException RetryAfter property documentation for .NET SDK v3.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet"}}
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

### ✅ 11. azure-documentation (via azure) — 5880ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete .NET SDK v3 throughput and 429 retry guidance.","parameters":{"url":"https://learn.microsoft.com/azure/cosmos-db/performance-tips-dotnet-sdk-v3#throughput"}}
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

I’ll ground the guidance in the current Cosmos DB .NET v3 retry behavior and recommended exception patterns.

The Cosmos DB .NET SDK v3 automatically retries HTTP 429 responses by default, honoring the service’s `x-ms-retry-after-ms` header. A `CosmosException` with status 429 normally reaches your code only after the SDK exhausts its configured retry count or cumulative wait time.

## Configure the built-in retry policy

Prefer the SDK policy unless you need an application-wide retry budget or custom telemetry:

```csharp
using Azure.Identity;
using Microsoft.Azure.Cosmos;

var options = new CosmosClientOptions
{
    MaxRetryAttemptsOnRateLimitedRequests = 9,
    MaxRetryWaitTimeOnRateLimitedRequests = TimeSpan.FromSeconds(30)
};

// Register CosmosClient as a singleton and reuse it for the application's lifetime.
CosmosClient client = new(
    accountEndpoint: Environment.GetEnvironmentVariable("COSMOS_ENDPOINT")
        ?? throw new InvalidOperationException("COSMOS_ENDPOINT is required."),
    tokenCredential: new DefaultAzureCredential(),
    clientOptions: options);

Container container = client.GetContainer("app-db", "items");
```

`MaxRetryAttemptsOnRateLimitedRequests` defaults to 9, meaning up to 10 total requests. The default cumulative retry wait is 30 seconds.

## Inspect `CosmosException`

The SDK exposes the parsed `x-ms-retry-after-ms` response header through `CosmosException.RetryAfter`:

```csharp
try
{
    ItemResponse<MyItem> response = await container.ReadItemAsync<MyItem>(
        id: itemId,
        partitionKey: new PartitionKey(partitionKey),
        cancellationToken: cancellationToken);
}
catch (CosmosException ex)
{
    HttpStatusCode statusCode = ex.StatusCode;
    TimeSpan? retryAfter = ex.RetryAfter; // Parsed x-ms-retry-after-ms header.

    logger.LogError(
        ex,
        "Cosmos request failed. Status={StatusCode}, RetryAfter={RetryAfter}, " +
        "ActivityId={ActivityId}, RequestCharge={RequestCharge}, Diagnostics={Diagnostics}",
        (int)statusCode,
        retryAfter,
        ex.ActivityId,
        ex.RequestCharge,
        ex.Diagnostics);

    throw;
}
```

## Custom bounded retry logic for 429

If you implement retries yourself, disable the SDK’s 429 retries so you do not accidentally stack two retry policies:

```csharp
var options = new CosmosClientOptions
{
    MaxRetryAttemptsOnRateLimitedRequests = 0
};
```

Then use a bounded policy that honors `RetryAfter`, adds exponential backoff and jitter, supports cancellation, and stops after a fixed delay budget:

```csharp
using System.Net;
using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Logging;

public static class CosmosRetry
{
    public static async Task<T> ExecuteAsync<T>(
        Func<CancellationToken, Task<T>> operation,
        ILogger logger,
        int maxAttempts = 5,
        TimeSpan? maxTotalDelay = null,
        CancellationToken cancellationToken = default)
    {
        ArgumentNullException.ThrowIfNull(operation);
        ArgumentNullException.ThrowIfNull(logger);

        if (maxAttempts < 1)
        {
            throw new ArgumentOutOfRangeException(nameof(maxAttempts));
        }

        TimeSpan delayBudget = maxTotalDelay ?? TimeSpan.FromSeconds(30);
        TimeSpan totalDelay = TimeSpan.Zero;

        for (int attempt = 1; ; attempt++)
        {
            try
            {
                return await operation(cancellationToken);
            }
            catch (CosmosException ex) when (
                ex.StatusCode == HttpStatusCode.TooManyRequests &&
                attempt < maxAttempts)
            {
                // RetryAfter is the SDK representation of x-ms-retry-after-ms.
                TimeSpan serverDelay =
                    ex.RetryAfter ?? TimeSpan.FromMilliseconds(100);

                double exponentialMs = Math.Min(
                    5_000,
                    100 * Math.Pow(2, attempt - 1));

                TimeSpan exponentialDelay =
                    TimeSpan.FromMilliseconds(exponentialMs);

                // Never retry earlier than the server requested.
                TimeSpan delay = serverDelay > exponentialDelay
                    ? serverDelay
                    : exponentialDelay;

                delay += TimeSpan.FromMilliseconds(Random.Shared.Next(0, 100));

                if (totalDelay + delay > delayBudget)
                {
                    logger.LogWarning(
                        "Cosmos 429 retry budget exhausted. Status={StatusCode}, " +
                        "Attempt={Attempt}, RetryAfter={RetryAfter}, ActivityId={ActivityId}",
                        (int)ex.StatusCode,
                        attempt,
                        ex.RetryAfter,
                        ex.ActivityId);

                    throw;
                }

                logger.LogWarning(
                    "Cosmos request throttled. Status={StatusCode}, " +
                    "Attempt={Attempt}/{MaxAttempts}, Delay={Delay}, " +
                    "ActivityId={ActivityId}, RequestCharge={RequestCharge}",
                    (int)ex.StatusCode,
                    attempt,
                    maxAttempts,
                    delay,
                    ex.ActivityId,
                    ex.RequestCharge);

                await Task.Delay(delay, cancellationToken);
                totalDelay += delay;
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
    maxTotalDelay: TimeSpan.FromSeconds(20),
    cancellationToken);
```

For sustained 429 rates, retries are only mitigation. Also investigate hot partitions, excessive concurrency, query RU cost, and whether the provisioned RU/s should be increased.

## Handle 404: item not found

A 404 is not transient, so retrying the same read is normally inappropriate. Return `null` when absence is expected:

```csharp
public static async Task<T?> ReadOrNullAsync<T>(
    Container container,
    string id,
    string partitionKey,
    CancellationToken cancellationToken = default)
{
    try
    {
        ItemResponse<T> response = await container.ReadItemAsync<T>(
            id,
            new PartitionKey(partitionKey),
            cancellationToken: cancellationToken);

        return response.Resource;
    }
    catch (CosmosException ex) when (
        ex.StatusCode == HttpStatusCode.NotFound)
    {
        return default;
    }
}
```

A 404 can also indicate an incorrect partition-key value, not just a missing ID. Log or propagate it when absence is unexpected.

## Handle 409: conflict

A 409 typically means the combination of ID and partition key already exists, or a unique-key constraint was violated. Retrying the unchanged create will continue to fail:

```csharp
public static async Task<bool> TryCreateAsync<T>(
    Container container,
    T item,
    string partitionKey,
    ILogger logger,
    CancellationToken cancellationToken = default)
{
    try
    {
        await container.CreateItemAsync(
            item,
            new PartitionKey(partitionKey),
            cancellationToken: cancellationToken);

        return true;
    }
    catch (CosmosException ex) when (
        ex.StatusCode == HttpStatusCode.Conflict)
    {
        logger.LogInformation(
            "Cosmos item was not created because it already exists or " +
            "violates a unique-key constraint. ActivityId={ActivityId}",
            ex.ActivityId);

        return false;
    }
}
```

Choose the operation according to the desired semantics:

- `CreateItemAsync`: fail with 409 if the item exists.
- `UpsertItemAsync`: create or replace; useful when that behavior is intentional.
- Read the existing item after 409 if an earlier create may have succeeded but its response was lost.
- Do not convert every create to an upsert merely to suppress conflicts, because upsert can overwrite existing data.

**References**

- [Design resilient applications with Azure Cosmos DB SDKs](https://learn.microsoft.com/azure/cosmos-db/conceptual-resilient-sdk-applications)
- [`CosmosException.RetryAfter`](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosexception.retryafter?view=azure-dotnet)
- [Performance tips for Azure Cosmos DB .NET SDK v3](https://learn.microsoft.com/azure/cosmos-db/performance-tips-dotnet-sdk-v3#throughput)
- [`MaxRetryAttemptsOnRateLimitedRequests`](https://learn.microsoft.com/dotnet/api/microsoft.azure.cosmos.cosmosclientoptions.maxretryattemptsonratelimitedrequests?view=azure-dotnet)

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 8 criteria passed. The response provides strong CosmosException handling and robust throttling, 404, and 409 examples, but omits SubStatusCode extraction and successful-response RU charge tracking.

### Strengths

- Uses CosmosException consistently with status-specific exception filters.
- Provides bounded 429 retry logic that honors server retry timing and avoids stacked SDK and application retries.
- Clearly handles expected 404 and 409 outcomes and explains their semantics.
- Configures the SDK v3 rate-limit retry policy and includes diagnostics logging.

### Issues

- CosmosException.SubStatusCode is not extracted or logged.
- RequestCharge is shown only on exceptions; successful ItemResponse.RequestCharge tracking is not demonstrated.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `CosmosException` as the primary exception type: Pass
      - Extracting `StatusCode`, `SubStatusCode`, and `RetryAfter` properties: Fail
      - Handling 429 (TooManyRequests) with retry-after backoff: Pass
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
hyoka run --prompt-id cosmos-db-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
