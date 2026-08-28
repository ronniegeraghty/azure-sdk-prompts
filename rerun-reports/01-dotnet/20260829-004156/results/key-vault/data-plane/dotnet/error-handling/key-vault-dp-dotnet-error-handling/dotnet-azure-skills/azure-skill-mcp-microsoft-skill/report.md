# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 162.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 4/7 |
| Duration | 162.1s |
| Timestamp | 2026-08-28T18:06:39Z |
| Files Generated | 0 |
| Event Count | 3278 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 134.7s |
| Review | 25.7s |
| **Total** | **162.1s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-security-keyvault-keys-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=329212 out=3805 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 650ms

**Input:**

```json
{"skill":"azure-security-keyvault-keys-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure.Security.KeyVault.Keys (.NET)

Client library for managing cryptographic keys in Azure Key Vault and Managed HSM.

## Installation

```bash
dotnet add package Azure.Security.KeyVault.Keys
dotnet add package Azure.Identity
```

**Current Version**: 4.7.0 (stable)

## Environment Variables

```bash
KEY_VAULT_NAME=<your-key-vault-name>  # Required: Key Vault name
AZURE_KEYVAULT_URL=https://<vault-name>.vault.azure.net  # Optional: full Key Vault URL
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
```

## Client Hierarchy

```
KeyClient (key management)
├── CreateKey / CreateRsaKey / CreateEcKey
├── GetKey / GetKeys
├── UpdateKeyProperties
├── DeleteKey / PurgeDeletedKey
├── BackupKey / RestoreKey
└── GetCryptographyClient() → CryptographyClient

CryptographyClient (cryptographic operations)
├── Encrypt / Decrypt
├── WrapKey / UnwrapKey
├── Sign / Verify
└── SignData / VerifyData

KeyResolver (key resolution)
└── Resolve(keyId) → CryptographyClient
```

## Authentication

### Microsoft Entra Token Credential

```csharp
using Azure.Identity;
using Azure.Security.KeyVault.Keys;

var keyVaultName = Environment.GetEnvironmentVariable("KEY_VAULT_NAME");
var kvUri = $"https://{keyVaultName}.vault.azure.net";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
var credential = new DefaultAzureCredential(
    DefaultAzureCredential.DefaultEnvironmentVariableName
);
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet#credential-classes
// var credential = new ManagedIdentityCredential();
var client = new KeyClient(new Uri(kvUri), credential);
```

### Service Principal

```csharp
var credential = new ClientSecretCredential(
    tenantId: "<tenant-id>
... (truncated)
```

### ✅ 2. glob — 1444ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 1074ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-security-keyvault-keys-dotnet","pattern":"**/examples.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 4638ms

**Input:**

```json
{"intent":"Provide .NET Azure.Security.KeyVault.Secrets error handling examples for RequestFailedException, HTTP 403, 404, 409 conflicts, and 429 throttling, including retry configuration.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 5779ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Security.KeyVault.Secrets .NET RequestFailedException Status ErrorCode and SecretClientOptions retry behavior for 403 404 409 429.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 5094ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET SDK error handling best practices for Key Vault Secrets code generation.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 6900ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find RequestFailedException and Key Vault Secrets retry documentation.","parameters":{"query":"Azure.Security.KeyVault.Secrets .NET RequestFailedException Status ErrorCode SecretClientOptions Retry 403 404 409 429"}}
```

**Output:**

```
{"results":[{"description":"description: Configures a SecretClient in C# with exponential backoff retry options to handle Azure Key Vault throttling, then retrieves a secret from the Key Vault.\npackage: Azure.Security.KeyVault.Secrets, Azure.Identity\nlanguage: csharp\n","codeSnippet":"SecretClientOptions options = new SecretClientOptions()\r\n    {\r\n        Retry =\r\n        {\r\n            Delay= TimeSpan.FromSeconds(2),\r\n            MaxDelay = TimeSpan.FromSeconds(16),\r\n            MaxRetries = 5,\r\n            Mode = RetryMode.Exponential\r\n         }\r\n    };\r\n    var client = new SecretClient(new Uri(\u0022https://keyVaultName.vault.azure.net\u0022), new DefaultAzureCredential(),options);\r\n                                 \r\n    //Retrieve Secret\r\n    secret = client.GetSecret(secretName);","link":"https://learn.microsoft.com/azure/key-vault/general/overview-throttling#how-to-throttle-your-app-in-response-to-service-limits","language":"csharp"},{"description":"description: Handles a RequestFailedException thrown when a secret is not found, checking for HTTP status 404 and printing the error code.\npackage: Azure.Core\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022NonexistentSecret\u0022);\r\n}\r\n// handle exception with status code 404\r\ncatch (RequestFailedException e) when (e.Status == 404)\r\n{\r\n    // handle not found error\r\n    Console.WriteLine(\u0022ErrorCode \u0022 \u002B e.ErrorCode);\r\n}","link":"https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Defines a console application that authenticates to Azure Key Vault using the VM\u0027s managed identity via DefaultAzureCredential, sets a secret with user input, retrieves it, and then deletes it, with retry policies configured for Key Vault operations.\npackage: Azure.Identity, Azure.Security.KeyVault.Secrets\nlanguage: csharp\n","codeSnippet":"cl
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 5966ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# code samples for Key Vault Secrets exception handling and retry settings.","parameters":{"language":"csharp","query":"Azure.Security.KeyVault.Secrets SecretClient RequestFailedException error handling retry SecretClientOptions"}}
```

**Output:**

```
{"results":[{"title":"Azure Key Vault secret client library for .NET - version 4.11.0","content":"# Azure Key Vault secret client library for .NET - version 4.11.0\n## Troubleshooting\nSee our [troubleshooting guide](https://github.com/Azure/azure-sdk-for-net/blob/Azure.Security.KeyVault.Secrets_4.11.0/sdk/keyvault/Azure.Security.KeyVault.Secrets/TROUBLESHOOTING.md) for details on how to diagnose various failure scenarios.\n### General\nWhen you interact with the Azure Key Vault secret client library using the .NET SDK, errors returned by the service correspond to the same HTTP status codes returned for [REST API](https://learn.microsoft.com/en-us/rest/api/keyvault/) requests.\nFor example, if you try to retrieve a secret that doesn\u0027t exist in your Azure Key Vault, a \u0060404\u0060 error is returned, indicating \u0060Not Found\u0060.\n\u0060\u0060\u0060C\ntry\n{\n    KeyVaultSecret secret = client.GetSecret(\u0022some_secret\u0022);\n}\ncatch (RequestFailedException ex)\n{\n    Console.WriteLine(ex.ToString());\n}\n\u0060\u0060\u0060\nYou will notice that additional information is logged, like the Client Request ID of the operation.","contentUrl":"https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting"},{"title":"Azure Key Vault throttling guidance","content":"# Azure Key Vault throttling guidance\n## How to throttle your app in response to service limits\nThe following are **best practices** you should implement when your service is throttled:\n1. Reduce the number of operations per request.\n2. Reduce the frequency of requests.\n3. Avoid immediate retries.\n3.1. All requests accrue against your usage limits.\nWhen you implement your app\u0027s error handling, use the HTTP error code 429 to detect the need for client-side throttling. If the request fails again with an HTTP 429 error code, you are still encountering an Azure service limit. Continue to use the recommended client-side throttling m
... (truncated)
```

### ✅ 9. rg — 4014ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt","pattern":"RequestFailedException|ErrorCode|Status|Retry-After|429|409|403|404"}
```

**Output:**

```
Output too large to read at once (32.1 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787940492544-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787940481508-copilot-tool-output-44a0c5b07d5945179b0de529e8b083c2.txt:{"results":[{"description":"description: Configures a SecretClient in C# with exponential backoff retry options to handle Azure Key Vault throttling, then retrieves a secret from the Key Vault.\npackage: Azure.Security.KeyVault.Secrets, Azure.Identity\nlanguage: csharp\n","codeSnippet":"SecretClientOptions options = new SecretClientOptions()\r\n    {\r\n        Retry =\r\n        {\r\n      
```

### ✅ 10. rg — 3111ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940481508-copilot-tool-output-44a0c5b07d5945179b0de529e8b083c2.txt","pattern":"SecretClientOptions|RetryMode|MaxRetries|MaxDelay|Delay|RequestFailedException|ErrorCode"}
```

**Output:**

```
Output too large to read at once (30.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787940492543-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt:{"results":[{"title":"Azure Key Vault secret client library for .NET - version 4.11.0","content":"# Azure Key Vault secret client library for .NET - version 4.11.0\n## Troubleshooting\nSee our [troubleshooting guide](https://github.com/Azure/azure-sdk-for-net/blob/Azure.Security.KeyVault.Secrets_4.11.0/sdk/keyvault/Azure.Security.KeyVault.Secrets/TROUBLESHOOTING.md) for details on how to dia
```

### ✅ 11. powershell — 2602ms

**Input:**

```json
{"command":"$doc = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt' | ConvertFrom-Json; $doc.results | Select-Object -First 10 title,contentUrl | Format-Table -AutoSize","description":"List Key Vault documentation search results"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt' | ConvertFrom-Json; $doc.r …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 12. view — 1376ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt","view_range":[1,5]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt b/C:/Users/weidxu/AppData/Local/Temp/1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787940481290-copilot-tool-output-9289f1ce9ea74a76ae74e2ad6a8c36f6.txt
@@ -1,5 +1,5 @@
 {"results":[{"title":"Azure Key Vault secret client library for .NET - version 4.11.0","content":"# Azure Key Vault secret client library for .NET - version 4.11.0\n## Troubleshooting\nSee our [troubleshooting guide](https://github.com/Azure/azure-sdk-for-net/blob/Azure.Security.KeyVault.Secrets_4.11.0/sdk/keyvault/Azure.Security.KeyVault.Secrets/TROUBLESHOOTING.md) for details on how to diagnose various failure scenarios.\n### General\nWhen you interact with the Azure Key Vault secret client library using the .NET SDK, errors returned by the service correspond to the same HTTP status codes returned for [REST API](https://learn.microsoft.com/en-us/rest/api/keyvault/) requests.\nFor example, if you try to retrieve a secret that doesn\u0027t exist in your Azure Key Vault, a \u0060404\u0060 error is returned, indicating \u0060Not Found\u0060.\n\u0060\u0060\u0060C\ntry\n{\n    KeyVaultSecret secret = client.GetSecret(\u0022some_secret\u0022);\n}\ncatch (RequestFailedException ex)\n{\n    Console.WriteLine(ex.ToString());\n}\n\u0060\u0060\u0060\nYou will notice that additional information is logged, like the Client Request ID of the operation.","contentUrl":"https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting"},{"title":"Azure Key Vault throttling guidance","content":"# Azure Key Vault throttling guidance\n## How to throttle your app in response to service limits\nThe following are **best practices** you should im
... (truncated)
```

## Copilot Response

I’m checking the SDK guidance for the exact `RequestFailedException` and retry patterns before showing the secret-specific code.

Use `RequestFailedException.Status` for reliable branching and `RequestFailedException.ErrorCode` for service-specific diagnostics. Configure retries on `SecretClient`; Azure SDK retries transient failures such as `429` before surfacing the exception.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

var options = new SecretClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(16),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

var client = new SecretClient(
    new Uri(Environment.GetEnvironmentVariable("AZURE_KEYVAULT_URL")
        ?? throw new InvalidOperationException("AZURE_KEYVAULT_URL is not configured.")),
    new DefaultAzureCredential(),
    options);
```

### Read a secret

Returning `null` for `404` is appropriate when “missing” is an expected domain outcome. Otherwise, rethrow it.

```csharp
static async Task<KeyVaultSecret?> GetSecretOrNullAsync(
    SecretClient client,
    string name,
    string? version = null,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.GetSecretAsync(name, version, cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        Console.Error.WriteLine(
            "Key Vault access denied. " +
            $"Status={ex.Status}, Code={ex.ErrorCode}, " +
            $"ClientRequestId={ex.ClientRequestId}");

        // Check the calling identity, Key Vault Secrets User/Officer RBAC role,
        // access policies, firewall rules, and private endpoint configuration.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 404)
    {
        Console.Error.WriteLine(
            $"Secret '{name}' or version '{version ?? "<latest>"}' was not found. " +
            $"Code={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        // This is reached only after SecretClient's configured retries are exhausted.
        Console.Error.WriteLine(
            "Key Vault throttling persisted after retries. " +
            $"Code={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        throw;
    }
    catch (RequestFailedException ex)
    {
        LogKeyVaultFailure(ex);
        throw;
    }
}
```

### Write a secret and handle conflicts

```csharp
static async Task<KeyVaultSecret> SetSecretAsync(
    SecretClient client,
    string name,
    string value,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.SetSecretAsync(name, value, cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        Console.Error.WriteLine(
            $"Not authorized to set secret '{name}'. " +
            $"Code={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 409)
    {
        Console.Error.WriteLine(
            $"Conflict while setting secret '{name}'. " +
            $"Code={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        // Common causes include a conflicting operation or a secret that is
        // deleted but still recoverable. Resolve the state before retrying.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        Console.Error.WriteLine(
            $"Key Vault throttled the write after retries. " +
            $"Code={ex.ErrorCode}, ClientRequestId={ex.ClientRequestId}");

        throw;
    }
    catch (RequestFailedException ex)
    {
        LogKeyVaultFailure(ex);
        throw;
    }
}
```

Secret values are immutable per version: `SetSecretAsync(name, value)` creates a **new version**. Key Vault Secrets does not provide a typical ETag-based compare-and-swap operation for secret values. A `409` generally indicates a vault/object state conflict rather than “that version already exists.” Retrieve a specific version by passing it to `GetSecretAsync`; a nonexistent version produces `404`.

### Extract status, error code, request ID, and response headers

```csharp
static void LogKeyVaultFailure(RequestFailedException ex)
{
    Response? response = ex.GetRawResponse();

    string? retryAfter = null;
    string? serviceRequestId = null;

    if (response is not null)
    {
        response.Headers.TryGetValue("Retry-After", out retryAfter);
        response.Headers.TryGetValue("x-ms-request-id", out serviceRequestId);
    }

    Console.Error.WriteLine(
        $"Key Vault request failed. " +
        $"HTTP={ex.Status}, " +
        $"ErrorCode={ex.ErrorCode ?? "<none>"}, " +
        $"ClientRequestId={ex.ClientRequestId ?? "<none>"}, " +
        $"ServiceRequestId={serviceRequestId ?? "<none>"}, " +
        $"RetryAfter={retryAfter ?? "<none>"}");
}
```

Typical error codes include `Forbidden`, `SecretNotFound`, `Conflict`, and `Throttled`, but codes can vary by operation and vault state. Branch primarily on `Status`; use `ErrorCode`, request IDs, and the exception message for diagnosis. Also note that credential acquisition can throw `AuthenticationFailedException` before any Key Vault HTTP response exists, so it has no Key Vault HTTP status.

References: [Key Vault Secrets .NET troubleshooting](https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting), [Key Vault throttling guidance](https://learn.microsoft.com/azure/key-vault/general/overview-throttling), and [Key Vault REST error codes](https://learn.microsoft.com/azure/key-vault/general/rest-error-codes).

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 7 criteria passed. The response provides strong RequestFailedException handling, diagnostics, throttling behavior, and retry configuration, but does not adequately distinguish authorization models, deleted-secret 404 scenarios, or soft-delete and purge-protection failures.

### Strengths

- Uses filtered RequestFailedException catches with a catch-all that rethrows.
- Clearly extracts and logs HTTP status, service error code, request IDs, and retry headers.
- Accurately explains that SetSecretAsync creates a new immutable version rather than producing a normal version collision.
- Provides a concrete exponential retry configuration and correctly describes exhausted 429 retries.

### Issues

- The 403 guidance does not distinguish RBAC authorization from legacy access-policy authorization or explain how to identify the active model.
- The 404 pattern does not check whether the secret exists in the deleted-secrets collection.
- Soft-delete recovery and purge-protection error handling are largely absent.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/7)
      - `RequestFailedException` for all Key Vault errors: Pass
      - Extracting `Status` and `ErrorCode` properties: Pass
      - 403 handling: diagnosing RBAC vs. access policy misconfiguration: Fail
      - 404 handling: secret not found vs. deleted secret: Fail
      - 429 throttling: Key Vault rate limits and retry behavior: Pass
      - Soft-delete and purge protection error scenarios: Fail
      - `SecretClientOptions` retry configuration: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 57% | 1.00 | 0.5714 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.5714** | **57.1%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
