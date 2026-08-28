# Evaluation Report: storage-dp-dotnet-auth

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 119.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-auth` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 3/5 |
| Duration | 119.1s |
| Timestamp | 2026-08-28T18:37:57Z |
| Files Generated | 0 |
| Event Count | 2085 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 103.3s |
| Review | 14.7s |
| **Total** | **119.1s** |

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
| Token Usage | in=203805 out=3348 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 641ms

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

### ✅ 2. view — 5103ms

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

### ✅ 3. view — 4786ms

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

### ✅ 4. glob — 4462ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 4152ms

**Input:**

```json
{"intent":"Provide current best practices for C# code using Azure.Storage.Blobs BlobServiceClient with DefaultAzureCredential, managed identity in production, and Azure CLI credentials in local development.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 5414ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for authenticating Azure Blob Storage in C# with DefaultAzureCredential, including managed identity in Azure, Azure CLI locally, required NuGet packages, user-assigned managed identity configuration, and required Blob data-plane RBAC roles.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2347ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure C# code-generation best practices for Blob Storage authentication with managed identity and local Azure CLI credentials.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3006ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official guidance for Azure Blob Storage C# authentication.","parameters":{"query":"C# Azure.Storage.Blobs BlobServiceClient DefaultAzureCredential managed identity Azure CLI local development NuGet packages RBAC Storage Blob Data Contributor"}}
```

**Output:**

```
{"results":[{"title":"Quickstart: Azure Blob Storage client library for .NET (blob-storage-quickstart-template)","content":"# Quickstart: Azure Blob Storage client library for .NET (blob-storage-quickstart-template)\n## Code examples\n### **Passwordless (Recommended)**\n\u0060DefaultAzureCredential\u0060 is a class provided by the Azure Identity client library for .NET, which you can learn more about on the [DefaultAzureCredential overview](https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication#defaultazurecredential). \u0060DefaultAzureCredential\u0060 supports multiple authentication methods and determines which method should be used at runtime. This approach enables your app to use different authentication methods in different environments (local vs. production) without implementing environment-specific code.\nThe order and locations in which \u0060DefaultAzureCredential\u0060 looks for credentials can be found in the [Azure Identity library overview](https://learn.microsoft.com/en-us/dotnet/api/overview/azure/Identity-readme#defaultazurecredential).\nFor example, your app can authenticate using your Visual Studio sign-in credentials with when developing locally. Your app can then use a [managed identity](https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview) once it has been deployed to Azure. No code changes are required for this transition.\n#### Assign roles to your Microsoft Entra user account\nWhen developing locally, make sure that the user account that is accessing blob data has the correct permissions. You\u0027ll need **Storage Blob Data Contributor** to read and write blob data. To assign yourself this role, you\u0027ll need to be assigned the **User Access Administrator** role, or another role that includes the **Microsoft.Authorization/roleAssignments/write** action. You can assign Azure RBAC roles to a user using the Azure portal, Azure CLI, or Azure PowerShell. For more information about the **Sto
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 2673ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# code samples for BlobServiceClient with Azure Identity credentials.","parameters":{"language":"csharp","query":"BlobServiceClient DefaultAzureCredential ManagedIdentityCredential Azure Storage Blob .NET"}}
```

**Output:**

```
{"results":[{"description":"description: Creates and authenticates BlobServiceClient and ServiceBusClient instances using DefaultAzureCredential for system-assigned managed identity in a .NET application.\npackage: Azure.Identity, Azure.Messaging.ServiceBus, Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"using Azure.Identity;\r\n    using Azure.Messaging.ServiceBus;\r\n    using Azure.Storage.Blobs;\r\n    \r\n    // Create DefaultAzureCredential instance that uses system-assigned managed identity\r\n    // in the underlying ManagedIdentityCredential.\r\n    DefaultAzureCredential credential = new();\r\n    \r\n    BlobServiceClient blobServiceClient = new(\r\n        new Uri(\u0022https://\u003Cyour-storage-account\u003E.blob.core.windows.net\u0022),\r\n        credential);\r\n    \r\n    ServiceBusClient serviceBusClient = new(\u0022\u003Cyour-namespace\u003E\u0022, credential);\r\n    ServiceBusSender sender = serviceBusClient.CreateSender(\u0022producttracking\u0022);","link":"https://learn.microsoft.com/azure/storage/common/multiple-identity-scenarios#connect-an-azure-hosted-app-to-multiple-azure-services","language":"csharp"},{"description":"description: Configures multiple Azure service clients in .NET using DefaultAzureCredential with user-assigned managed identities to connect to two Blob Storage accounts, Azure Cosmos DB, and Azure SQL Database.\npackage: Azure.Core, Azure.Identity, Azure.Storage.Blobs, Microsoft.Azure.Cosmos, Microsoft.Data.SqlClient\nlanguage: csharp\n","codeSnippet":"using Azure.Core;\r\n    using Azure.Identity;\r\n    using Azure.Storage.Blobs;\r\n    using Microsoft.Azure.Cosmos;\r\n    using Microsoft.Data.SqlClient;\r\n    \r\n    string clientIdStorage =\r\n        Environment.GetEnvironmentVariable(\u0022Managed_Identity_Client_ID_Storage\u0022)!;\r\n    \r\n    // Create a DefaultAzureCredential instance that configures the underlying\r\n    // ManagedIdentityCredential to use a user-assigned managed identity.\r\n    De
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 3118ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the official .NET Blob Storage quickstart for package and passwordless authentication details.","parameters":{"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-dotnet"}}
```

**Output:**

```
# Credential chains in the Azure Identity library for .NET

The Azure Identity library provides *credentials*—public classes derived from the Azure Core library's [TokenCredential](/en-us/dotnet/api/azure.core.tokencredential?view=azure-dotnet&amp;preserve-view=true) class. A credential represents a distinct authentication flow for acquiring an access token from Microsoft Entra ID. These credentials can be chained together to form an ordered sequence of authentication mechanisms to be attempted.

## How a chained credential works

At runtime, a credential chain attempts to authenticate using the sequence's first credential. If that credential fails to acquire an access token, the next credential in the sequence is attempted, and so on, until an access token is successfully obtained. The following sequence diagram illustrates this behavior:

![Credential chain sequence diagram](../media/mermaidjs/chain-sequence.svg)

## Why use credential chains

A chained credential can offer the following benefits:

- **Environment awareness**: Automatically selects the most appropriate credential based on the environment in which the app is running. Without it, you'd have to write code like this:

    ```csharp
    TokenCredential credential;
    
    if (app.Environment.IsProduction() || app.Environment.IsStaging())
    {
        credential = new ManagedIdentityCredential(
            ManagedIdentityId.FromUserAssignedClientId(userAssignedClientId));
    }
    else
    {
        // local development environment
        credential = new VisualStudioCredential();
    }
    ```
- **Seamless transitions**: Your app can move from local development to your staging or production environment without changing authentication code.
- **Improved resiliency**: Includes a fallback mechanism that moves to the next credential when the prior fails to acquire an access token.

## How to choose a chained credential

There are two disparate philosophies to creden
... (truncated)
```

### ✅ 11. azure-documentation (via azure) — 2800ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the official DefaultAzureCredential overview for credential chain behavior and managed identity configuration.","parameters":{"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains"}}
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

### ✅ 12. azure-documentation (via azure) — 2459ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official guidance for authorizing Blob Storage access with Microsoft Entra ID and Azure RBAC.","parameters":{"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/authorize-access-azure-active-directory"}}
```

**Output:**

```
# Quickstart: Azure Blob Storage client library for .NET

::: zone pivot="blob-storage-quickstart-scratch"

Note

The **Build from scratch** option walks you step by step through the process of creating a new project, installing packages, writing the code, and running a basic console app. This approach is recommended if you want to understand all the details involved in creating an app that connects to Azure Blob Storage. If you prefer to automate deployment tasks and start with a completed project, choose [Start with a template](storage-quickstart-blobs-dotnet?pivots=blob-storage-quickstart-template).

::: zone-end

::: zone pivot="blob-storage-quickstart-template"

Note

The **Start with a template** option uses the Azure Developer CLI to automate deployment tasks and starts you off with a completed project. This approach is recommended if you want to explore the code as quickly as possible without going through the setup tasks. If you prefer step by step instructions to build the app, choose [Build from scratch](storage-quickstart-blobs-dotnet?pivots=blob-storage-quickstart-scratch).

::: zone-end

This quickstart shows you how to use the Azure Blob Storage client library for .NET to create a container, upload and download blobs, and list blobs in a container.

::: zone pivot="blob-storage-quickstart-scratch"

In this article, you follow steps to install the package and try out example code for basic tasks.

::: zone-end

::: zone pivot="blob-storage-quickstart-template"

In this article, you use the [Azure Developer CLI](/en-us/azure/developer/azure-developer-cli/overview) to deploy Azure resources and run a completed console app with just a few commands.

::: zone-end

[API reference documentation](/en-us/dotnet/api/azure.storage.blobs) | [Library source code](https://github.com/Azure/azure-sdk-for-net/tree/master/sdk/storage/Azure.Storage.Blobs) | [Package (NuGet)](https://www.nuget.org/packages/Azure.Storage.Blobs) | [Samples](
... (truncated)
```

## Copilot Response

I’ll ground the setup in the current Azure SDK authentication guidance, including local and managed-identity behavior.

Use `DefaultAzureCredential`, but restrict it to **Azure CLI locally** and **managed identity in production**. This avoids accidentally selecting another credential from the default chain.

### NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

Use `Azure.Identity` **1.15.0 or later**, which supports selecting a specific credential through `DefaultAzureCredential`.

### Configuration

`appsettings.json`:

```json
{
  "AzureStorage": {
    "ServiceUri": "https://<storage-account-name>.blob.core.windows.net",
    "ManagedIdentityClientId": ""
  }
}
```

Leave `ManagedIdentityClientId` empty for a **system-assigned managed identity**. For a **user-assigned managed identity**, set it through an environment variable rather than hardcoding deployment-specific configuration:

```text
AzureStorage__ManagedIdentityClientId=<user-assigned-identity-client-id>
```

### ASP.NET Core setup

`Program.cs`:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

var builder = WebApplication.CreateBuilder(args);

var storageServiceUri =
    builder.Configuration["AzureStorage:ServiceUri"]
    ?? throw new InvalidOperationException(
        "AzureStorage:ServiceUri is required.");

var credentialOptions = new DefaultAzureCredentialOptions();

if (!builder.Environment.IsDevelopment())
{
    var managedIdentityClientId =
        builder.Configuration["AzureStorage:ManagedIdentityClientId"];

    if (!string.IsNullOrWhiteSpace(managedIdentityClientId))
    {
        credentialOptions.ManagedIdentityClientId = managedIdentityClientId;
    }
}

TokenCredential credential = builder.Environment.IsDevelopment()
    ? new DefaultAzureCredential(
        "AzureCliCredential",
        credentialOptions)
    : new DefaultAzureCredential(
        "ManagedIdentityCredential",
        credentialOptions);

builder.Services.AddSingleton(
    new BlobServiceClient(
        new Uri(storageServiceUri),
        credential));

var app = builder.Build();

app.MapGet("/blobs/{containerName}", async (
    string containerName,
    BlobServiceClient blobServiceClient,
    CancellationToken cancellationToken) =>
{
    var container = blobServiceClient.GetBlobContainerClient(containerName);
    var blobNames = new List<string>();

    await foreach (var blob in container.GetBlobsAsync(
        cancellationToken: cancellationToken))
    {
        blobNames.Add(blob.Name);
    }

    return Results.Ok(blobNames);
});

app.Run();
```

`BlobServiceClient` is thread-safe and should be registered once as a singleton.

### Local development

Authenticate the Azure CLI with a developer identity that has Blob data-plane access:

```powershell
az login
az account set --subscription "<subscription-name-or-id>"
```

Run under the Development environment:

```powershell
$env:ASPNETCORE_ENVIRONMENT = "Development"
dotnet run
```

The code will use only `AzureCliCredential`.

### Production

1. Enable a system-assigned or user-assigned managed identity on the Azure host.
2. Grant that identity an appropriate data-plane role:
   - **Storage Blob Data Reader** for downloads/listing.
   - **Storage Blob Data Contributor** for upload, update, and delete operations.
3. Scope the role as narrowly as possible—prefer the target container when appropriate.
4. Ensure `ASPNETCORE_ENVIRONMENT` is not set to `Development`.

The local developer account needs the same appropriate Blob data role. Azure subscription `Contributor` alone does **not** grant access to blob data.

If you do not require deterministic credential selection, the shorter form below naturally tries managed identity before Azure CLI:

```csharp
var client = new BlobServiceClient(
    new Uri("https://<storage-account-name>.blob.core.windows.net"),
    new DefaultAzureCredential());
```

However, the environment-specific version above is preferable for production because it prevents other credentials in the default chain from being selected unexpectedly.

References: [Credential chains in Azure Identity for .NET](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains), [Blob Storage quickstart for .NET](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-dotnet), and [Authorize Blob access with Microsoft Entra ID](https://learn.microsoft.com/azure/storage/blobs/authorize-access-azure-active-directory).

## Reviewer Notes (LLM-as-Judge)

### Summary

3 of 5 criteria passed. The setup correctly installs the packages and constructs BlobServiceClient with DefaultAzureCredential, but it omits the full requested chain explanation and credential-unavailable error handling.

### Strengths

- Includes both required NuGet package commands.
- Shows complete ASP.NET Core registration and managed-identity configuration.
- Uses explicit Azure CLI credentials locally and managed identity in production.
- Includes relevant RBAC and local-development guidance.

### Issues

- The credential-chain explanation does not cover environment-variable credentials.
- No authentication error handling is shown for unavailable credentials.

## Grader Results

- authentication.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (3/5)
      - Installing `Azure.Identity` and `Azure.Storage.Blobs` packages: Pass
      - Creating a `DefaultAzureCredential` instance: Pass
      - Passing the credential to `BlobServiceClient`: Pass
      - Explanation of the credential chain (managed identity → CLI → env vars): Fail
      - Error handling when no credential is available: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 60% | 1.00 | 0.6000 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6000** | **60.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-auth --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
