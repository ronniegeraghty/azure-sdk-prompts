# Evaluation Report: identity-dp-dotnet-service-principal

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 152.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-service-principal` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 152.9s |
| Timestamp | 2026-08-28T17:53:51Z |
| Files Generated | 0 |
| Event Count | 3196 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 134.0s |
| Review | 17.6s |
| **Total** | **152.9s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-dotnet, azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=445075 out=4281 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 2330ms

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for .NET

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```bash
dotnet add package Azure.Identity

# For ASP.NET Core integration
dotnet add package Microsoft.Extensions.Azure

# For brokered authentication and Visual Studio Code credential support
dotnet add package Azure.Identity.Broker
```

## Environment Variables

### Service Principal with Secret

```bash
AZURE_CLIENT_ID=<application-client-id>
AZURE_TENANT_ID=<directory-tenant-id>
AZURE_CLIENT_SECRET=<client-secret-value>
```

### Service Principal with Certificate

```bash
AZURE_CLIENT_ID=<application-client-id>
AZURE_TENANT_ID=<directory-tenant-id>
AZURE_CLIENT_CERTIFICATE_PATH=<path-to-pfx-or-pem>
AZURE_CLIENT_CERTIFICATE_PASSWORD=<certificate-password>  # Optional
```

### Managed Identity

```bash
AZURE_CLIENT_ID=<user-assigned-managed-identity-client-id>  # Only for user-assigned
```

## DefaultAzureCredential

The recommended credential for most scenarios. Tries multiple authentication methods in order. See [DefaultAzureCredential overview](https://aka.ms/azsdk/net/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

### Basic Usage

```csharp
using Azure.Identity;
using Azure.Storage.Blobs;

var credential = new DefaultAzureCredential();
var blobClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);
```

### ASP.NET Core with Dependency Injection

```csharp
using Azure.Identity;
using Microsoft.Extensions.Azure;

builder.Services.AddAzureClients(clientBuilder =>
{
    clientBuilder.AddBlobServiceClient(
        new Uri("https://myaccount.blob.core.windows.net"));
    clientBuilder.AddSecretClient(
        new Uri("https://myvault.vault.azure.net"));
    
    // Uses DefaultAzureCredential by default
    clientBuilder.UseCredential(new DefaultAzureCredential());
});
```

### Customizing DefaultAz
... (truncated)
```

### ✅ 2. skill — 3462ms

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

### ✅ 3. glob — 10211ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 13204ms

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

### ✅ 5. view — 11555ms

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

### ✅ 6. glob — 8932ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 5723ms

**Input:**

```json
{"intent":"Generate a local-only C# example that authenticates to Azure Blob Storage with ClientSecretCredential, validates configuration, and handles invalid credentials and authorization errors.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 7160ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for C# ClientSecretCredential, BlobServiceClient token credentials, AuthenticationFailedException, RequestFailedException, and secure client-secret storage guidance.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 9. azure-get_azure_bestpractices (via azure) — 4886ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only C# example that authenticates to Azure Blob Storage with ClientSecretCredential, validates configuration, and handles invalid credentials and authorization errors.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 5552ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Microsoft documentation for C# ClientSecretCredential, BlobServiceClient token credentials, AuthenticationFailedException, RequestFailedException, and secure client-secret storage guidance.","parameters":{"query":"C# Azure.Identity ClientSecretCredential BlobServiceClient AuthenticationFailedException RequestFailedException client secret environment variables Key Vault managed identity"}}
```

**Output:**

```
{"results":[{"title":"Azure Identity client library for .NET - version 1.17.2","content":"# Azure Identity client library for .NET - version 1.17.2\n## Environment variables\n[\u0060DefaultAzureCredential\u0060](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet) and [\u0060EnvironmentCredential\u0060](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.environmentcredential?view=azure-dotnet) can be configured with environment variables. Each type of authentication requires values for specific variables. Configuration is attempted in the order in which these environment variables are listed. For example, if values for a client secret and certificate are both present, the client secret is used by \u0060EnvironmentCredential\u0060.\n### Service principal with secret\n| Variable name | Value| \n|  --- | ---  |\n| \u0060AZURE_CLIENT_ID\u0060 | ID of a Microsoft Entra application |\n| \u0060AZURE_TENANT_ID\u0060 | ID of the application\u0027s Microsoft Entra tenant |\n| \u0060AZURE_CLIENT_SECRET\u0060 | one of the application\u0027s client secrets |\n\n### Service principal with certificate\n| Variable name | Value| \n|  --- | ---  |\n| \u0060AZURE_CLIENT_ID\u0060 | ID of a Microsoft Entra application |\n| \u0060AZURE_TENANT_ID\u0060 | ID of the application\u0027s Microsoft Entra tenant |\n| \u0060AZURE_CLIENT_CERTIFICATE_PATH\u0060 | path to a PFX or PEM-encoded certificate file including private key |\n| \u0060AZURE_CLIENT_CERTIFICATE_PASSWORD\u0060 | (optional) the password protecting the certificate file (currently only supported for PFX (PKCS12) certificates) |\n| \u0060AZURE_CLIENT_SEND_CERTIFICATE_CHAIN\u0060 | (optional) send certificate chain in x5c header to support subject name / issuer based authentication |\n\n### Workload identity (\u0060DefaultAzureCredential\u0060)\n| Variable name | Value| \n|  --- | ---  |\n| \u0060AZURE_CLIENT_ID\u0060 | The client ID of the application the workload identity wi
... (truncated)
```

### ✅ 11. azure-documentation (via azure) — 5205ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for ClientSecretCredential with BlobServiceClient and authentication error handling.","parameters":{"language":"csharp","query":"Azure.Identity ClientSecretCredential Azure.Storage.Blobs BlobServiceClient AuthenticationFailedException RequestFailedException"}}
```

**Output:**

```
{"results":[{"description":"description: Connects to Azure Blob Storage using managed identity or service principal authentication with Azure SDK for .NET, retrieving the Blob endpoint from environment variables.\npackage: Azure.Identity, Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"using Azure.Identity;\r\nusing Azure.Storage.Blobs;\r\n\r\n// get Blob endpoint\r\nvar blobEndpoint = Environment.GetEnvironmentVariable(\u0022AZURE_STORAGEBLOB_RESOURCEENDPOINT\u0022);\r\n\r\n// Uncomment the following lines corresponding to the authentication type you want to use.\r\n// system-assigned managed identity\r\n// var credential = new DefaultAzureCredential();\r\n\r\n// user-assigned managed identity\r\n// var credential = new DefaultAzureCredential(\r\n//     new DefaultAzureCredentialOptions\r\n//     {\r\n//         ManagedIdentityClientId = Environment.GetEnvironmentVariable(\u0022AZURE_STORAGEBLOB_CLIENTID\u0022);\r\n//     });\r\n\r\n// service principal \r\n// var tenantId = Environment.GetEnvironmentVariable(\u0022AZURE_STORAGEBLOB_TENANTID\u0022);\r\n// var clientId = Environment.GetEnvironmentVariable(\u0022AZURE_STORAGEBLOB_CLIENTID\u0022);\r\n// var clientSecret = Environment.GetEnvironmentVariable(\u0022AZURE_STORAGEBLOB_CLIENTSECRET\u0022);\r\n// var credential = new ClientSecretCredential(tenantId, clientId, clientSecret);\r\n\r\nvar blobServiceClient = new BlobServiceClient(\r\n        new Uri(blobEndpoint),\r\n        credential);","link":"https://learn.microsoft.com/azure/service-connector/how-to-integrate-storage-blob#default-environment-variable-names-or-application-properties-and-sample-code","language":"csharp"},{"description":"description: Connects to Azure Blob Storage in .NET using connection string from environment variables with Azure Storage Blob client.\npackage: Azure.Storage.Blob\nlanguage: csharp\n","codeSnippet":"using Azure.Identity;\r\nusing Azure.Storage.Blobs;\r\n\r\n// get Blob endpoint\r\nvar blobEndpoint = Environment.GetEnv
... (truncated)
```

### ✅ 12. rg — 2494ms

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

**Output:**

```
Output too large to read at once (45.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787939711005-3.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt:{"results":[{"title":"Azure Identity client library for .NET - version 1.17.2","content":"# Azure Identity client library for .NET - version 1.17.2\n## Environment variables\n[\u0060DefaultAzureCredential\u0060](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet) and [\u0060EnvironmentCredential\u0060](https://learn.microsoft.com/en-us/dotnet
```

### ✅ 13. rg — 2078ms

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787939703244-copilot-tool-output-833bc58c039d40f48c5fdeed437c8fd6.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

**Output:**

```
Output too large to read at once (45.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787939711004-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt:{"results":[{"title":"Azure Identity client library for .NET - version 1.17.2","content":"# Azure Identity client library for .NET - version 1.17.2\n## Environment variables\n[\u0060DefaultAzureCredential\u0060](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet) and [\u0060EnvironmentCredential\u0060](https://learn.microsoft.com/en-us/dotnet
```

### ✅ 14. rg — 1615ms

**Input:**

```json
{"head_limit":15,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt","pattern":"\"title\":\"[^\"]+\""}
```

**Output:**

```
Output too large to read at once (21.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787939711005-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787939703244-copilot-tool-output-833bc58c039d40f48c5fdeed437c8fd6.txt:{"results":[{"description":"description: Connects to Azure Blob Storage using managed identity or service principal authentication with Azure SDK for .NET, retrieving the Blob endpoint from environment variables.\npackage: Azure.Identity, Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"using Azure.Identity;\r\nusing Azure.Storage.Blobs;\r\n\r\n// get Blob endpoint\r\nvar blobEndpoint
```

### ✅ 15. powershell — 1424ms

**Input:**

```json
{"command":"$p = 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt'; $j = Get-Content -Raw $p | ConvertFrom-Json; $j.results | Select-Object title, contentUrl | Format-Table -AutoSize","description":"Summarize Microsoft documentation search results"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … 3464583f7e3.txt'; $j = Get-Content -Raw $p | ConvertFrom-Json; $j.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 16. view — 748ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt","view_range":[1,10]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt b/C:/Users/weidxu/AppData/Local/Temp/1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787939703854-copilot-tool-output-9b626ba03685465d8f3783464583f7e3.txt
@@ -1,10 +1,10 @@
 {"results":[{"title":"Azure Identity client library for .NET - version 1.17.2","content":"# Azure Identity client library for .NET - version 1.17.2\n## Environment variables\n[\u0060DefaultAzureCredential\u0060](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.defaultazurecredential?view=azure-dotnet) and [\u0060EnvironmentCredential\u0060](https://learn.microsoft.com/en-us/dotnet/api/azure.identity.environmentcredential?view=azure-dotnet) can be configured with environment variables. Each type of authentication requires values for specific variables. Configuration is attempted in the order in which these environment variables are listed. For example, if values for a client secret and certificate are both present, the client secret is used by \u0060EnvironmentCredential\u0060.\n### Service principal with secret\n| Variable name | Value| \n|  --- | ---  |\n| \u0060AZURE_CLIENT_ID\u0060 | ID of a Microsoft Entra application |\n| \u0060AZURE_TENANT_ID\u0060 | ID of the application\u0027s Microsoft Entra tenant |\n| \u0060AZURE_CLIENT_SECRET\u0060 | one of the application\u0027s client secrets |\n\n### Service principal with certificate\n| Variable name | Value| \n|  --- | ---  |\n| \u0060AZURE_CLIENT_ID\u0060 | ID of a Microsoft Entra application |\n| \u0060AZURE_TENANT_ID\u0060 | ID of the application\u0027s Microsoft Entra tenant |\n| \u0060AZURE_CLIENT_CERTIFICATE_PATH\u0060 | path to a PFX or PEM-encoded certificate file including private key |\n| \u0060AZURE_CLI
... (truncated)
```

### ✅ 17. azure-documentation (via azure) — 3230ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official Azure Identity .NET documentation for ClientSecretCredential environment variable configuration and exceptions.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet"}}
```

**Output:**

```
# Authenticate to Azure resources from .NET apps hosted on-premises

Apps hosted outside of Azure, such as on-premises or in a third-party data center, should use an application service principal through [Microsoft Entra ID](/en-us/entra/fundamentals/whatis) to authenticate to Azure services. In the sections ahead, you learn:

- How to register an application with Microsoft Entra to create a service principal
- How to assign roles to scope permissions
- How to authenticate using a service principal from your app code

Using dedicated application service principals allows you to adhere to the principle of least privilege when accessing Azure resources. Permissions are limited to the specific requirements of the app during development, preventing accidental access to Azure resources intended for other apps or services. This approach also helps avoid issues when the app is moved to production by ensuring it isn't over-privileged in the development environment.

A different app registration should be created for each environment the app is hosted in. This allows environment specific resource permissions to be configured for each service principal and make sure an app deployed to one environment doesn't talk to Azure resources that are part of another environment.

## Register the app in Azure

Application service principal objects are created through an app registration in Azure using either the Azure portal or Azure CLI.

# [Azure portal](#tab/azure-portal)
1. In the Azure portal, use the search bar to navigate to the **App registrations** page.
2. On the **App registrations** page, select **+ New registration**.
3. On the **Register an application** page:

    - For the **Name** field, enter a descriptive value that includes the app name and the target environment.
    - For the **Supported account types**, select **Accounts in this organizational directory only (Microsoft Customer Led only - Single tenant)**, or whichever option best fits your r
... (truncated)
```

### ✅ 18. azure-documentation (via azure) — 2846ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official on-premises .NET service principal authentication guidance, especially secure storage of client secrets.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/authentication/on-premises-apps"}}
```

**Output:**

```
# Authorize access to blobs by using Microsoft Entra ID

Azure Storage supports using Microsoft Entra ID to authorize requests to blob data. By using Microsoft Entra ID, you can use Azure role-based access control (Azure RBAC) to grant permissions to a security principal, which can be a user, group, or application service principal. Microsoft Entra ID authenticates the security principal and returns an OAuth 2.0 token. Use the token to authorize a request against the Blob service.

You can use Microsoft Entra ID authorization with all general-purpose and Blob storage accounts in all public regions and national clouds. Only storage accounts created by using the Azure Resource Manager deployment model support Microsoft Entra authorization.

Important

For optimal security, Microsoft recommends using Microsoft Entra ID with managed identities to authorize requests against blob, queue, and table data, whenever possible. Authorization with Microsoft Entra ID and managed identities provides superior security and ease of use over Shared Key authorization. To learn more about managed identities, see [What are managed identities for Azure resources](/en-us/entra/identity/managed-identities-azure-resources/overview). For an example of how to enable and use a managed identity for a .NET application, see [Authenticating Azure-hosted apps to Azure resources with .NET](/en-us/dotnet/azure/sdk/authentication/azure-hosted-apps).

For resources hosted outside of Azure, such as on-premises applications, you can use managed identities through Azure Arc. For example, apps running on Azure Arc-enabled servers can use managed identities to connect to Azure services. To learn more, see [Authenticate against Azure resources with Azure Arc-enabled servers](/en-us/azure/azure-arc/servers/managed-identity-authentication).

For scenarios where shared access signatures (SAS) are used, Microsoft recommends using a user delegation SAS. A user delegation SAS is secured with Microsoft E
... (truncated)
```

### ✅ 19. azure-documentation (via azure) — 2490ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Azure Blob Storage .NET authorization guidance for TokenCredential and RBAC.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/authorize-access-azure-active-directory"}}
```

**Output:**

```
# Azure Identity client library for .NET - version 1.17.2

The Azure Identity library provides [Microsoft Entra ID](/en-us/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [`TokenCredential`](/en-us/dotnet/api/azure.core.tokencredential?view=azure-dotnet) implementations that can be used to construct Azure SDK clients that support Microsoft Entra token authentication.

[Source code](https://github.com/Azure/azure-sdk-for-net/tree/Azure.Identity_1.17.2/sdk/identity/Azure.Identity/src) | [Package (NuGet)](https://www.nuget.org/packages/Azure.Identity) | [API reference documentation](/en-us/dotnet/api/azure.identity?view=azure-dotnet) | [Microsoft Entra ID documentation](/en-us/entra/identity/)

## Getting started

### Install the package

Install the Azure Identity client library for .NET with NuGet:

```dotnetcli
dotnet add package Azure.Identity
```

### Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/dotnet/).
- The [Azure CLI](/en-us/cli/azure) can also be useful for authenticating in a development environment, creating accounts, and managing account roles.

### Authenticate the client

When debugging and executing code locally, it's typical for a developer to use their own account for authenticating calls to Azure services. There are several developer tools that can be used to perform this authentication in your development environment. For more information, see [Authentication during local development](/en-us/dotnet/azure/sdk/authentication/#authentication-during-local-development).

## Key concepts

### Credentials

A credential is a class that contains or can obtain the data needed for a service client to authenticate requests. Service clients across the Azure SDK accept credentials when they're constructed. Service clients use those credentials to authenticate requests to the service.

The Azure Identity library focuses on OAuth authentication with
... (truncated)
```

## Copilot Response

I’ll ground the example in the Azure Identity and Blob Storage SDK guidance, including authentication versus authorization failures.

Install the identity and Blob Storage SDK packages:

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

`ClientSecretCredential` receives the service principal’s **tenant ID**, **application/client ID**, and **client-secret value**. Creating the credential does not contact Microsoft Entra ID; authentication occurs when the SDK performs its first operation.

### Complete `Program.cs`

```csharp
using Azure;
using Azure.Identity;
using Azure.Storage.Blobs;

return await RunAsync();

static async Task<int> RunAsync()
{
    try
    {
        string tenantId = GetRequiredEnvironmentVariable("AZURE_TENANT_ID");
        string clientId = GetRequiredEnvironmentVariable("AZURE_CLIENT_ID");
        string clientSecret = GetRequiredEnvironmentVariable("AZURE_CLIENT_SECRET");
        Uri blobEndpoint = GetRequiredHttpsUri("AZURE_STORAGE_BLOB_ENDPOINT");

        var credential = new ClientSecretCredential(
            tenantId: tenantId,
            clientId: clientId,
            clientSecret: clientSecret);

        // Azure SDK clients and credentials are thread-safe and should be reused.
        var blobServiceClient = new BlobServiceClient(
            serviceUri: blobEndpoint,
            credential: credential);

        // The first service operation triggers token acquisition.
        Console.WriteLine("Containers:");

        await foreach (var container in blobServiceClient.GetBlobContainersAsync())
        {
            Console.WriteLine($"- {container.Name}");
        }

        return 0;
    }
    catch (CredentialUnavailableException ex)
    {
        Console.Error.WriteLine(
            $"The credential could not be used in this environment: {ex.Message}");
        return 2;
    }
    catch (AuthenticationFailedException ex)
    {
        // Typical causes: invalid tenant/client ID, expired secret, or using
        // the secret's ID instead of its value.
        Console.Error.WriteLine(
            $"Microsoft Entra authentication failed: {ex.Message}");
        return 3;
    }
    catch (RequestFailedException ex) when (ex.Status == 401)
    {
        Console.Error.WriteLine(
            $"Azure Storage rejected the access token. " +
            $"ErrorCode={ex.ErrorCode}; RequestId={ex.ClientRequestId}");
        return 4;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        // Authentication succeeded, but the service principal lacks Blob
        // data-plane permissions or a storage policy/firewall denied access.
        Console.Error.WriteLine(
            $"Authenticated, but not authorized for Blob Storage. " +
            $"ErrorCode={ex.ErrorCode}; RequestId={ex.ClientRequestId}");
        return 5;
    }
    catch (RequestFailedException ex)
    {
        Console.Error.WriteLine(
            $"Azure Storage request failed. Status={ex.Status}; " +
            $"ErrorCode={ex.ErrorCode}; RequestId={ex.ClientRequestId}; " +
            $"Message={ex.Message}");
        return 6;
    }
    catch (InvalidOperationException ex)
    {
        Console.Error.WriteLine($"Configuration error: {ex.Message}");
        return 7;
    }
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Required environment variable '{name}' is not set.");
    }

    return value;
}

static Uri GetRequiredHttpsUri(string name)
{
    string value = GetRequiredEnvironmentVariable(name);

    if (!Uri.TryCreate(value, UriKind.Absolute, out Uri? uri) ||
        uri.Scheme != Uri.UriSchemeHttps)
    {
        throw new InvalidOperationException(
            $"Environment variable '{name}' must contain an absolute HTTPS URI.");
    }

    return uri;
}
```

Set process-scoped variables before running the application:

```powershell
$env:AZURE_TENANT_ID = "<directory-tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<account-name>.blob.core.windows.net"

dotnet run
```

The service principal also needs a Blob **data-plane** RBAC role. For this read-only listing example, use **Storage Blob Data Reader**, scoped as narrowly as practical. Use **Storage Blob Data Contributor** only if the application must write or delete blobs. An Azure management-plane role such as Contributor does not automatically grant access to blob data.

### Secret-storage guidance

| Value | Recommended storage |
|---|---|
| Tenant ID, client ID, Blob endpoint | Environment variables or non-secret configuration such as `appsettings.json` |
| Client secret | A secret manager that injects it at runtime; environment variable when necessary |
| Azure-hosted production workload | Prefer managed identity instead of a client secret |

Do not commit the secret to `appsettings.json`, `launchSettings.json`, source code, container images, or CI configuration files. Environment variables are preferable to checked-in configuration, but they are a delivery mechanism rather than a full secret-management system and may be visible to privileged processes or diagnostics.

For production, use the hosting platform’s secret store or a managed secret manager, restrict access, avoid logging configuration values, and rotate the secret before expiration. When hosted in Azure, replace the service principal secret with `ManagedIdentityCredential` where possible.

References:

- [Azure Identity library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme)
- [Authenticate on-premises .NET apps with a service principal](https://learn.microsoft.com/dotnet/azure/sdk/authentication/on-premises-apps)
- [Authorize Blob Storage with Microsoft Entra ID](https://learn.microsoft.com/azure/storage/blobs/authorize-access-azure-active-directory)

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response provides a complete C# example covering packages, credential construction, Blob client integration, environment-based secret handling, and authentication error handling.

### Strengths

- Complete, directly usable example with input validation and distinct authentication, authorization, and service error handling.
- Strong secret-management guidance, including avoiding checked-in secrets, rotating credentials, and preferring managed identity for Azure-hosted workloads.

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - `Azure.Identity` package with `ClientSecretCredential` class: Pass
      - Constructor parameters: tenantId, clientId, clientSecret: Pass
      - Passing credential to Azure SDK clients: Pass
      - Environment variable patterns for secret storage: Pass
      - `AuthenticationFailedException` for invalid credentials: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-service-principal --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
