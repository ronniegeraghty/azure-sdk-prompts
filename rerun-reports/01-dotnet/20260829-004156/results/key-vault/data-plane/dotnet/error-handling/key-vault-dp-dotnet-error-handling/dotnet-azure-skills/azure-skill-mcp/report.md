# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 134.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 4/7 |
| Duration | 134.4s |
| Timestamp | 2026-08-28T18:04:25Z |
| Files Generated | 0 |
| Event Count | 2966 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 108.2s |
| Review | 25.1s |
| **Total** | **134.4s** |

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
| Token Usage | in=405440 out=4987 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. azure-get_azure_bestpractices (via azure) — 3101ms

**Input:**

```json
{"intent":"Provide .NET Azure SDK code patterns for handling Azure Key Vault Secrets failures with Azure.Security.KeyVault.Secrets and Azure.RequestFailedException, including HTTP 403, 404, 409/version conflicts, and 429 throttling.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 2. azure-documentation (via azure) — 3445ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Security.KeyVault.Secrets SecretClient exception handling, RequestFailedException Status and ErrorCode properties, Key Vault 403/404/409/429 errors, and retry configuration in .NET.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 3. azure-get_azure_bestpractices (via azure) — 3146ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure .NET SDK code-generation best practices for robust Key Vault Secrets exception and retry handling.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 4. azure-documentation (via azure) — 3999ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official documentation for .NET Key Vault Secrets error handling.","parameters":{"query":"Azure.Security.KeyVault.Secrets .NET RequestFailedException Status ErrorCode 403 404 409 429 retry SecretClientOptions Retry"}}
```

**Output:**

```
{"results":[{"description":"description: Handles a RequestFailedException thrown when a secret is not found, checking for HTTP status 404 and printing the error code.\npackage: Azure.Core\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022NonexistentSecret\u0022);\r\n}\r\n// handle exception with status code 404\r\ncatch (RequestFailedException e) when (e.Status == 404)\r\n{\r\n    // handle not found error\r\n    Console.WriteLine(\u0022ErrorCode \u0022 \u002B e.ErrorCode);\r\n}","link":"https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Demonstrates how to handle exceptions when attempting to retrieve a secret from Azure Key Vault using the .NET client library, specifically catching a RequestFailedException when the secret is not found.\npackage: Azure.Security.KeyVault.Secrets\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022some_secret\u0022);\r\n}\r\ncatch (RequestFailedException ex)\r\n{\r\n    Console.WriteLine(ex.ToString());\r\n}","link":"https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting","language":"c"},{"description":"description: Initializes a ManagedIdentityCredential to authenticate a Service Fabric application and uses it to create a SecretClient for Azure Key Vault, then fetches a secret asynchronously while handling potential exceptions.\npackage: Azure.Identity, Azure.Security.KeyVault.Secrets\nlanguage: csharp\n","codeSnippet":"using Azure.Identity;\r\nusing Azure.Security.KeyVault.Secrets;\r\n\r\nnamespace MyMIService\r\n{\r\n    internal sealed class MyMIService : StatelessService\r\n    {\r\n        protected override async Task RunAsync(CancellationToken cancellationToken)\r\n        {\r\n            try\r\n            {\r\n                // Load the service fabric application managed identity assign
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 3579ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for handling Key Vault Secrets RequestFailedException.","parameters":{"language":"csharp","query":"Azure.Security.KeyVault.Secrets SecretClient RequestFailedException Status ErrorCode retry C#"}}
```

**Output:**

```
{"results":[{"title":"Azure Key Vault throttling guidance","content":"# Azure Key Vault throttling guidance\n## How to throttle your app in response to service limits\nThe following are **best practices** you should implement when your service is throttled:\n1. Reduce the number of operations per request.\n2. Reduce the frequency of requests.\n3. Avoid immediate retries.\n3.1. All requests accrue against your usage limits.\nWhen you implement your app\u0027s error handling, use the HTTP error code 429 to detect the need for client-side throttling. If the request fails again with an HTTP 429 error code, you are still encountering an Azure service limit. Continue to use the recommended client-side throttling method, retrying the request until it succeeds.\nHere is code that implements exponential backoff:\n\u0060\u0060\u0060csharp\nSecretClientOptions options = new SecretClientOptions()\n    {\n        Retry =\n        {\n            Delay= TimeSpan.FromSeconds(2),\n            MaxDelay = TimeSpan.FromSeconds(16),\n            MaxRetries = 5,\n            Mode = RetryMode.Exponential\n         }\n    };\n    var client = new SecretClient(new Uri(\u0022https://keyVaultName.vault.azure.net\u0022), new DefaultAzureCredential(),options);\n\n    //Retrieve Secret\n    secret = client.GetSecret(secretName);\n\u0060\u0060\u0060\nUsing this code in a client C# application is straightforward.\n### Recommended client-side throttling method\nOn HTTP error code 429, begin throttling your client using an exponential backoff approach:\n1. Wait 1 second, retry request\n2. If still throttled wait 2 seconds, retry request\n3. If still throttled wait 4 seconds, retry request\n4. If still throttled wait 8 seconds, retry request\n5. If still throttled wait 16 seconds, retry request\nAt this point, you should not be getting HTTP 429 response codes.","contentUrl":"https://learn.microsoft.com/azure/key-vault/general/overview-throttling#how-to-throttle-your-app-in-response-to-service-limits"
... (truncated)
```

### ✅ 6. rg — 1902ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt","pattern":"\"title\"|\"url\"|RequestFailedException|409|429|Retry-After|SecretClientOptions|ErrorCode"}
```

**Output:**

```
Output too large to read at once (34.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787940301450-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt:{"results":[{"title":"Azure Key Vault throttling guidance","content":"# Azure Key Vault throttling guidance\n## How to throttle your app in response to service limits\nThe following are **best practices** you should implement when your service is throttled:\n1. Reduce the number of operations per request.\n2. Reduce the frequency of requests.\n3. Avoid immediate retries.\n3.1. All requests a
```

### ✅ 7. rg — 1503ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt","pattern":"\"url\"|RequestFailedException|Status|ErrorCode|409|429|SecretClientOptions|Retry"}
```

**Output:**

```
Output too large to read at once (25.5 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787940301449-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt:{"results":[{"description":"description: Handles a RequestFailedException thrown when a secret is not found, checking for HTTP status 404 and printing the error code.\npackage: Azure.Core\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022NonexistentSecret\u0022);\r\n}\r\n// handle exception with status code 404\r\ncatch (RequestFailedException e) w
```

### ✅ 8. powershell — 2599ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt' | ConvertFrom-Json; $d.results | Select-Object title,url | Format-Table -AutoSize","description":"Extract documentation titles and URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-77421d4895cb4a88834ff8b1b8ec1a14.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 9. powershell — 2192ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt' | ConvertFrom-Json; $d.results | Select-Object -First 8 description,url | Format-List","description":"Extract sample descriptions and URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-8333573689d1454285bd49e7331e6676.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 10. view — 1417ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt b/C:/Users/weidxu/AppData/Local/Temp/1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt
@@ -1,6 +1,6 @@
 {"results":[{"description":"description: Handles a RequestFailedException thrown when a secret is not found, checking for HTTP status 404 and printing the error code.\npackage: Azure.Core\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022NonexistentSecret\u0022);\r\n}\r\n// handle exception with status code 404\r\ncatch (RequestFailedException e) when (e.Status == 404)\r\n{\r\n    // handle not found error\r\n    Console.WriteLine(\u0022ErrorCode \u0022 \u002B e.ErrorCode);\r\n}","link":"https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Demonstrates how to handle exceptions when attempting to retrieve a secret from Azure Key Vault using the .NET client library, specifically catching a RequestFailedException when the secret is not found.\npackage: Azure.Security.KeyVault.Secrets\nlanguage: c\n","codeSnippet":"try\r\n{\r\n    KeyVaultSecret secret = client.GetSecret(\u0022some_secret\u0022);\r\n}\r\ncatch (RequestFailedException ex)\r\n{\r\n    Console.WriteLine(ex.ToString());\r\n}","link":"https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting","language":"c"},{"description":"description: Initializes a ManagedIdentityCredential to authenticate a Service Fabric application and uses it to create a SecretClient for Azure Key Vault, then fetches a secret asynchronously while handling p
... (truncated)
```

### ✅ 11. view — 1071ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940295693-copilot-tool-output-8333573689d1454285bd49e7331e6676.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt b/C:/Users/weidxu/AppData/Local/Temp/1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787940294900-copilot-tool-output-77421d4895cb4a88834ff8b1b8ec1a14.txt
@@ -1,6 +1,6 @@
 {"results":[{"title":"Azure Key Vault throttling guidance","content":"# Azure Key Vault throttling guidance\n## How to throttle your app in response to service limits\nThe following are **best practices** you should implement when your service is throttled:\n1. Reduce the number of operations per request.\n2. Reduce the frequency of requests.\n3. Avoid immediate retries.\n3.1. All requests accrue against your usage limits.\nWhen you implement your app\u0027s error handling, use the HTTP error code 429 to detect the need for client-side throttling. If the request fails again with an HTTP 429 error code, you are still encountering an Azure service limit. Continue to use the recommended client-side throttling method, retrying the request until it succeeds.\nHere is code that implements exponential backoff:\n\u0060\u0060\u0060csharp\nSecretClientOptions options = new SecretClientOptions()\n    {\n        Retry =\n        {\n            Delay= TimeSpan.FromSeconds(2),\n            MaxDelay = TimeSpan.FromSeconds(16),\n            MaxRetries = 5,\n            Mode = RetryMode.Exponential\n         }\n    };\n    var client = new SecretClient(new Uri(\u0022https://keyVaultName.vault.azure.net\u0022), new DefaultAzureCredential(),options);\n\n    //Retrieve Secret\n    secret = client.GetSecret(secretName);\n\u0060\u0060\u0060\nUsing this code in a client C# application is straightforward.\n### Recommended client-side throttling method\nOn HTTP error code 429, begin throttling your cl
... (truncated)
```

### ✅ 12. azure-documentation (via azure) — 2549ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Clarify Azure Key Vault Secrets HTTP 409 conflict scenarios in the .NET SDK.","parameters":{"query":"Azure Key Vault secrets HTTP 409 Conflict SecretExists soft deleted secret same name recover purge SetSecret version conflict RequestFailedException"}}
```

**Output:**

```
{"results":[{"title":"RequestFailedException.ErrorCode Property-Definition","content":"# RequestFailedException.ErrorCode Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure](https://learn.microsoft.com/dotnet/api/azure?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.Core.dll\r\n\r\n- Package:\r\n    - Azure.Core v1.62.0\r\n\r\n- Source:\r\n    - [RequestFailedException.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/RequestFailedException.cs#L34C36-L34C40)\r\n\r\nGets the service specific error code if available. Please refer to the client documentation for the list of supported error codes.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic string? ErrorCode { get; }\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n[String](https://learn.microsoft.com/en-us/dotnet/api/system.string)","contentUrl":"https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.errorcode?view=azure-dotnet"},{"title":"Azure Core shared client library for .NET - version 1.62.0","content":"# Azure Core shared client library for .NET - version 1.62.0\n## Examples\n\u0060\u0060\u0060C\n// Setup a listener to monitor logged events.\nusing AzureEventSourceListener listener = AzureEventSourceListener.CreateConsoleLogger();\n\u0060\u0060\u0060\nMore on logging in [diagnostics samples](https://github.com/Azure/azure-sdk-for-net/blob/Azure.Core_1.62.0/sdk/core/Azure.Core/samples/Diagnostics.md).\n### Reporting Errors \u0060RequestFailedException\u0060\nWhen a service call fails \u0060Azure.RequestFailedException\u0060 would get thrown. The exception type provides a Status property with an HTTP status code and an ErrorCode property with a service-specific error code.\n\u0060\u0060\u0060C\ntry\n{\n    KeyVaultSecret secret = client.GetSecret(\u0022NonexistentSecret\u0022);\n}\n// handle exception with status code 404\ncatch (RequestFailedException e) when (e.Status == 404)\n{\n    // handle not found error
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 2141ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find RequestFailedException .NET API property details.","parameters":{"query":"Azure.Core RequestFailedException .NET Status ErrorCode GetRawResponse ClientRequestId API documentation"}}
```

**Output:**

```
{"results":[{"title":"Azure Key Vault recovery management with soft delete and purge protection","content":"# Azure Key Vault recovery management with soft delete and purge protection\n## List, recover, or purge a soft-deleted key vault\n1. Sign in to the Azure portal.\n2. Select the search bar at the top of the page.\n3. Search for the \u0022Key Vault\u0022 service. Do not select an individual key vault.\n4. At the top of the screen, select the option to \u0022Manage deleted vaults\u0022\n5. A context pane opens on the right side of your screen.\n6. Select your subscription.\n7. If your key vault has been soft deleted, it appears in the context pane on the right.\n8. If there are too many vaults, you can either select \u0022Load More\u0022 at the bottom of the context pane or use CLI or PowerShell to get the results.\n9. Once you find the vault you wish to recover or purge, select the checkbox next to it.\n10. Select the recover option at the bottom of the context pane if you would like to recover the key vault.\n11. Select the purge option if you would like to permanently delete the key vault.\n![On Key vaults, the Manage deleted vaults option is highlighted.](https://learn.microsoft.com/azure/key-vault/media/key-vault-recovery-3.png)\n![On Manage deleted key vaults, the only listed key vault is highlighted and selected, and the Recover button is highlighted.](https://learn.microsoft.com/azure/key-vault/media/key-vault-recovery-4.png)\n## List, recover or purge soft deleted secrets, keys, and certificates\n1. Sign in to the Azure portal.\n2. Select your key vault.\n3. Select the blade corresponding to the secret type you want to manage (keys, secrets, or certificates).\n4. At the top of the screen, select \u0022Manage deleted (keys, secrets, or certificates)\n5. A context pane appears on the right side of your screen.\n6. If your secret, key, or certificate does not appear in the list, it is not in the soft-deleted state.\n7. Select the secret, key, or certificate
... (truncated)
```

### ✅ 14. powershell — 1920ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940324165-copilot-tool-output-179735ef83764d78b182b37a3d121dbe.txt' -First 1) | ConvertFrom-Json; $d.results | ForEach-Object { \"TITLE: $($_.title)`nURL: $($_.contentUrl)`n$($_.content.Substring(0, [Math]::Min(700, $_.content.Length)))`n\" }","description":"Summarize 409 conflict documentation"}
```

**Output:**

```
TITLE: Azure Key Vault recovery management with soft delete and purge protection
URL: https://learn.microsoft.com/azure/key-vault/general/key-vault-recovery#list,-recover,-or-purge-a-soft-deleted-key-vault
# Azure Key Vault recovery management with soft delete and purge protection
## List, recover, or purge a soft-deleted key vault
1. Sign in to the Azure portal.
2. Select the search bar at the top of the page.
3. Search for the "Key Vault" service. Do not select an individual key vault.
4. At the top of the screen, select the option to "Manage deleted vaults"
5. A context pane opens on the right side of your screen.
6. Select your subscription.
7. If your key vault has been soft deleted, it appears in the context pane on the right.
8. If there are too many vaults, you can either select "Load More" at the bottom of the context pane or use CLI or PowerShell to get the results.
9. Once you find 

TITLE: Azure Key Vault: soft-delete overview
URL: https://learn.microsoft.com/azure/key-vault/general/soft-delete-overview
# Azure Key Vault: soft-delete overview
Important
If a key vault does not have soft-delete protection enabled, deleting a key deletes it permanently. Customers are strongly encouraged to turn on soft delete enforcement for their vaults via [Azure Policy](https://learn.microsoft.com/azure/key-vault/policy-reference).
Important
When a Key Vault is soft-deleted, services that are integrated with the Key Vault are deleted. For example: Azure RBAC roles assignments and Event Grid subscriptions. Recovering a soft-deleted Key Vault does not restore these services. They must be recreated.
Key Vault's soft-delete feature allows recovery of the deleted vaults and deleted key vault objects (for example

TITLE: Azure Key Vault recovery management with soft delete and purge protection
URL: https://learn.microsoft.com/azure/key-vault/general/key-vault-recovery#secrets-cli
# Azure Key Vault recovery management with soft delete and purge protection
## Secrets (CLI)
1. Gr
... (truncated)
```

### ✅ 15. powershell — 1568ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787940324191-copilot-tool-output-06e4e8afddb5471ca89a24456e7b76e4.txt' -First 1) | ConvertFrom-Json; $d.results | ForEach-Object { \"TITLE: $($_.title)`nURL: $($_.contentUrl)`n$($_.content.Substring(0, [Math]::Min(650, $_.content.Length)))`n\" }","description":"Summarize RequestFailedException documentation"}
```

**Output:**

```
TITLE: RequestFailedException.ErrorCode Property-Definition
URL: https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.errorcode?view=azure-dotnet
# RequestFailedException.ErrorCode Property

## Definition

- Namespace:
    - [Azure](https://learn.microsoft.com/dotnet/api/azure?view=azure-dotnet)

- Assembly:
    - Azure.Core.dll

- Package:
    - Azure.Core v1.62.0

- Source:
    - [RequestFailedException.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/RequestFailedException.cs#L34C36-L34C40)

Gets the service specific error code if available. Please refer to the client documentation for the list of supported error codes.

```csharp
public string? ErrorCode { get; }
```

#### Property Value

[S

TITLE: Azure Core shared client library for .NET - version 1.62.0
URL: https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#examples
# Azure Core shared client library for .NET - version 1.62.0
## Examples
```C
// Setup a listener to monitor logged events.
using AzureEventSourceListener listener = AzureEventSourceListener.CreateConsoleLogger();
```
More on logging in [diagnostics samples](https://github.com/Azure/azure-sdk-for-net/blob/Azure.Core_1.62.0/sdk/core/Azure.Core/samples/Diagnostics.md).
### Reporting Errors `RequestFailedException`
When a service call fails `Azure.RequestFailedException` would get thrown. The exception type provides a Status property with an HTTP status code and an ErrorCode property with a service-specific error code.
```C
try
{
    KeyVaultSec

TITLE: RequestFailedException.Status Property-Definition
URL: https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.status?view=azure-dotnet
# RequestFailedException.Status Property

## Definition

- Namespace:
    - [Azure](https://learn.microsoft.com/dotnet/api/azure?view=azure-dotnet)

- Assembly:
    - Azure.Core.dll

- Package:
    - Azure.Core v1.62.0

- Source:
    - [Req
... (truncated)
```

### ✅ 16. azure-documentation (via azure) — 3308ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Azure Core RequestFailedException handling guidance.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#examples"}}
```

**Output:**

```
# Azure Key Vault throttling guidance

Throttling is a process you initiate that limits the number of concurrent calls to the Azure service to prevent overuse of resources. Azure Key Vault (AKV) is designed to handle a high volume of requests. If an overwhelming number of requests occurs, throttling your client's requests helps maintain optimal performance and resiliency of the AKV service.

Throttling limits vary based on the scenario. For example, if you are performing a large volume of writes, the possibility for throttling is higher than if you are only performing reads.

## How does Key Vault handle its limits?

Service limits in Key Vault prevent misuse of resources and ensure quality of service for all of Key Vault's clients. When a service threshold is exceeded, Key Vault limits any further requests from that client, returns HTTP status code 429 (Too many requests), and the request fails. Failed requests that return a 429 do not count towards the throttle limits tracked by Key Vault.

Key Vault was originally designed to store and retrieve your secrets at deployment time. As technology evolved, Key Vault is now increasingly used at runtime to store and retrieve secrets. Many applications and services use Key Vault similar to a database. However, the current service limits are not designed to support such high throughput scenarios.

Key Vault was originally created with the limits specified in [Azure Key Vault service limits](service-limits). To maximize your Key Vault throughput rates, here are some recommended guidelines/best practices for maximizing your throughput:

1. Ensure that throttling is in place. Client must honor exponential back-off policies for 429s and ensure you are doing retries per the guidance.
2. Divide your Key Vault traffic among multiple vaults and different regions. Use a separate vault for each security/availability domain. If you have five apps, each in two regions, then we recommend 10 vaults each containing the secr
... (truncated)
```

### ✅ 17. azure-documentation (via azure) — 2942ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Azure Key Vault throttling guidance.","parameters":{"url":"https://learn.microsoft.com/azure/key-vault/general/overview-throttling#how-to-throttle-your-app-in-response-to-service-limits"}}
```

**Output:**

```
# Azure Core shared client library for .NET - version 1.62.0

Azure.Core provides shared primitives, abstractions, and helpers for modern .NET Azure SDK client libraries. These libraries follow the [Azure SDK Design Guidelines for .NET](https://azure.github.io/azure-sdk/dotnet_introduction.html) and can be easily identified by package and namespaces names starting with 'Azure', e.g. `Azure.Storage.Blobs`. A more complete list of client libraries using Azure.Core can be found [here](https://github.com/Azure/azure-sdk-for-net#core-services).

Azure.Core allows client libraries to expose common functionality in a consistent fashion, so that once you learn how to use these APIs in one client library, you will know how to use them in other client libraries.

[Source code](https://github.com/Azure/azure-sdk-for-net/tree/Azure.Core_1.62.0/sdk/core/Azure.Core/src) | [Package (NuGet)](https://www.nuget.org/packages/Azure.Core/) | [API reference documentation](/en-us/dotnet/api/azure.core)

## Getting started

Typically, you will not need to install Azure.Core; it will be installed for you when you install one of the client libraries using it. In case you want to install it explicitly (to implement your own client library, for example), you can find the NuGet package [here](https://www.nuget.org/packages/Azure.Core).

## Key concepts

The main shared concepts of Azure.Core (and so Azure SDK libraries using Azure.Core) include:

- Configuring service clients, e.g. configuring retries, logging (`ClientOptions`).
- Accessing HTTP response details (`Response`, `Response<T>`).
- Calling long-running operations (`Operation<T>`).
- Paging and asynchronous streams (`AsyncPageable<T>`).
- Exceptions for reporting errors from service requests in a consistent fashion. (`RequestFailedException`).
- Customizing requests (`RequestContext`).
- Abstractions for representing Azure SDK credentials. (`TokenCredentials`).
- Authenticating Azure SDK clients to Microsoft Ent
... (truncated)
```

### ✅ 18. azure-documentation (via azure) — 2577ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch official Azure Key Vault Secrets .NET client guidance.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting"}}
```

**Output:**

```
# Azure Key Vault secret client library for .NET - version 4.11.0

Azure Key Vault is a cloud service that provides a secure storage of secrets, such as passwords and database connection strings.

The Azure Key Vault secrets client library allows you to securely store and control the access to tokens, passwords, API keys, and other secrets. This library offers operations to create, retrieve, update, delete, purge, backup, restore, and list the secrets and its versions.

[Source code](https://github.com/Azure/azure-sdk-for-net/tree/Azure.Security.KeyVault.Secrets_4.11.0/sdk/keyvault/Azure.Security.KeyVault.Secrets/src) | [Package (NuGet)](https://www.nuget.org/packages/Azure.Security.KeyVault.Secrets/) | [API reference documentation](/en-us/dotnet/api/azure.security.keyvault.secrets) | [Product documentation](/en-us/azure/key-vault/) | [Samples](https://github.com/Azure/azure-sdk-for-net/tree/Azure.Security.KeyVault.Secrets_4.11.0/sdk/keyvault/Azure.Security.KeyVault.Secrets/samples) | [Migration guide](https://github.com/Azure/azure-sdk-for-net/blob/Azure.Security.KeyVault.Secrets_4.11.0/sdk/keyvault/Azure.Security.KeyVault.Secrets/MigrationGuide.md)

## Getting started

### Install the package

Install the Azure Key Vault secrets client library for .NET with [NuGet](https://www.nuget.org/):

```dotnetcli
dotnet add package Azure.Security.KeyVault.Secrets
```

### Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/dotnet/).
- An existing Azure Key Vault. If you need to create an Azure Key Vault, you can use the Azure Portal or [Azure CLI](/en-us/cli/azure).
- Authorization to an existing Azure Key Vault using either [RBAC](/en-us/azure/key-vault/general/rbac-guide) (recommended) or [access control](/en-us/azure/key-vault/general/assign-access-policy).

If you use the Azure CLI, replace `<your-resource-group-name>` and `<your-key-vault-name>` with your own, unique names:

```PowerShell
az keyvault create --resource-gr
... (truncated)
```

## Copilot Response

I’ll ground the exception patterns in the current Azure SDK behavior and official Key Vault guidance.

The important nuance is that `SetSecretAsync` normally creates a new version rather than conflicting; I’m checking which secret operations actually surface `409` so the catch pattern does not imply incorrect behavior.

Use `RequestFailedException.Status` for the HTTP status and `RequestFailedException.ErrorCode` for Key Vault’s service-specific code. Configure retries on the shared `SecretClient`; by the time a `429` reaches your catch block, the SDK has already exhausted those retries.

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
    new Uri(Environment.GetEnvironmentVariable("KEY_VAULT_URI")
        ?? throw new InvalidOperationException("KEY_VAULT_URI is not configured.")),
    new DefaultAzureCredential(),
    options);
```

### Reading a secret

Treat `404` as an expected result only if “missing” is valid for your application. Re-throw authorization, throttling, and unexpected failures so callers do not mistake them for success.

```csharp
static async Task<KeyVaultSecret?> GetSecretOrNullAsync(
    SecretClient client,
    string secretName,
    string? version,
    CancellationToken cancellationToken)
{
    try
    {
        Response<KeyVaultSecret> response = await client.GetSecretAsync(
            secretName,
            version,
            cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        LogKeyVaultFailure("Access denied", ex);

        // Check Key Vault data-plane RBAC/access policies and firewall rules.
        // Do not repeatedly retry a persistent permissions error.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 404)
    {
        LogKeyVaultFailure("Secret or requested version not found", ex);
        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        LogKeyVaultFailure("Key Vault throttling persisted after SDK retries", ex);
        throw;
    }
    catch (RequestFailedException ex)
    {
        LogKeyVaultFailure("Unexpected Key Vault failure", ex);
        throw;
    }
}
```

### Writing a secret and handling conflicts

`SetSecretAsync` does **not** conflict merely because the secret already exists—it creates a new version. A common `409` occurs when the name belongs to a soft-deleted, recoverable secret. That condition requires recovering or purging the deleted secret, or choosing another name; blindly retrying will not fix it.

```csharp
static async Task<KeyVaultSecret> SetSecretAsync(
    SecretClient client,
    string secretName,
    string secretValue,
    CancellationToken cancellationToken)
{
    try
    {
        Response<KeyVaultSecret> response = await client.SetSecretAsync(
            secretName,
            secretValue,
            cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        LogKeyVaultFailure("Access denied while setting secret", ex);
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 409)
    {
        LogKeyVaultFailure("Secret conflict", ex);

        // Often means the secret is deleted but recoverable.
        // Resolve explicitly through recovery/purge workflow.
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        LogKeyVaultFailure("Key Vault throttling persisted after SDK retries", ex);
        throw;
    }
    catch (RequestFailedException ex)
    {
        LogKeyVaultFailure("Unexpected Key Vault failure", ex);
        throw;
    }
}
```

### Extracting status, error code, and response headers

`ErrorCode` is service-specific and nullable, so use the HTTP status as the primary classification. The raw response can provide request IDs and retry information useful for diagnostics.

```csharp
static void LogKeyVaultFailure(string context, RequestFailedException ex)
{
    Response? response = ex.GetRawResponse();

    string? requestId = null;
    string? retryAfter = null;

    if (response is not null)
    {
        response.Headers.TryGetValue("x-ms-request-id", out requestId);
        response.Headers.TryGetValue("Retry-After", out retryAfter);
    }

    Console.Error.WriteLine(
        "{0}: HTTP {1}, ErrorCode={2}, RequestId={3}, RetryAfter={4}, Message={5}",
        context,
        ex.Status,                         // 403, 404, 409, 429, etc.; 0 if no response
        ex.ErrorCode ?? "<not supplied>", // e.g. SecretNotFound
        requestId ?? "<not supplied>",
        retryAfter ?? "<not supplied>",
        ex.Message);
}
```

Avoid logging `KeyVaultSecret.Value` or other secret material. Also distinguish service authorization failures from credential acquisition failures: invalid or unavailable credentials may throw `AuthenticationFailedException` or `CredentialUnavailableException` from `Azure.Identity`, rather than a `403 RequestFailedException`.

**References:** [Azure Core failure handling](https://learn.microsoft.com/dotnet/api/overview/azure/core-readme?view=azure-dotnet#reporting-errors-requestfailedexception), [Key Vault Secrets .NET client](https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet), [Key Vault throttling guidance](https://learn.microsoft.com/azure/key-vault/general/overview-throttling#how-to-throttle-your-app-in-response-to-service-limits).

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 7 criteria passed. The response gives strong RequestFailedException patterns, diagnostics, throttling guidance, and retry configuration, but incompletely addresses authorization-model diagnosis, deleted-secret detection, and purge-protection scenarios.

### Strengths

- Uses status-filtered and fallback RequestFailedException catches consistently.
- Clearly demonstrates extracting Status, ErrorCode, request ID, and retry headers.
- Correctly explains SetSecretAsync version behavior and the soft-deleted-name 409 conflict.
- Provides practical SecretClientOptions exponential retry configuration and avoids logging secret values.

### Issues

- The 403 guidance does not distinguish how to troubleshoot RBAC mode versus access-policy mode.
- The 404 example does not check whether the secret exists in the deleted-secrets collection.
- Purge protection and errors from recovery or purge workflows are not covered.

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
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
