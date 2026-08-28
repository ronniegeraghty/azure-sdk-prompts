# Evaluation Report: storage-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 122.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 4/6 |
| Duration | 122.0s |
| Timestamp | 2026-08-28T18:51:39Z |
| Files Generated | 0 |
| Event Count | 3091 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 101.7s |
| Review | 19.1s |
| **Total** | **122.0s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=481555 out=4125 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 761ms

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

### ✅ 2. glob — 3950ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 5252ms

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

### ✅ 4. view — 4771ms

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

### ✅ 5. azure-get_azure_bestpractices (via azure) — 4328ms

**Input:**

```json
{"intent":"Provide current best practices for idiomatic .NET code using Azure.Storage.Blobs, especially RequestFailedException handling and retry behavior.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 5612ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure.Storage.Blobs .NET RequestFailedException properties Status and ErrorCode, Blob Storage REST error codes for 403, 404, and 429, and Azure SDK retry configuration.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2970ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get current Azure .NET SDK code-generation practices for handling Azure Blob Storage failures and retries.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3644ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Verify Azure.Storage.Blobs .NET exception handling details.","parameters":{"query":"Azure.Storage.Blobs .NET RequestFailedException Status ErrorCode 403 404 429 retry Blob Storage"}}
```

**Output:**

```
{"results":[{"title":"Troubleshoot 403 errors in Azure Blob Storage","content":"# Troubleshoot 403 errors in Azure Blob Storage\n## Comprehensive 403 error diagnostic checklist\nIf you receive an error other than the errors described in the previous sections, or if the causes and solutions described in those sections don\u0027t help you resolve your error, use this checklist as a guide to identify the cause and find a solution.\n1. **Secure file transfer**: If secure transfer is required, make sure that requests are made over HTTPS. (**[Learn more](https://learn.microsoft.com/en-us/azure/storage/common/storage-require-secure-transfer)**).\n2. **Azure role-based access control**: Verify the security principal has the correct Azure RBAC role. (**[Learn more](https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access)**).\n3. **Shared access signature (SAS) tokens**: Check token expiration, permissions, and generation method. (**Read more**).\n4. **User delegation SAS**: Confirm required fields, permissions, and key status for user delegation SAS. (**Read more**).\n5. **Stored access policies**: Review stored access policy settings and the propagation delay. (**Read more**).\n6. **Access control lists (ACLs)**: Ensure ACL entries and permissions are correct for the client. (**Read more**).\n7. **Shared Key authorization**: If you\u0027re using an account key, confirm that Shared Key authorization is allowed. (**[Learn more](https://learn.microsoft.com/en-us/azure/storage/common/shared-key-authorization-prevent)**).\n8. **Public network endpoint**: Check public endpoint access, firewall rules, and network security perimeter. (**Read more**).\n9. **Private endpoints**: Validate private endpoint subnet and resource configuration. (**Read more**).\n10. **Encryption scopes**: Review encryption scope settings and customer-managed key status. (**Read more**).\n11. **Disabled accounts**: Confirm the storage account or subscription isn\u0027t disabled. 
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3318ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# examples for Blob Storage RequestFailedException handling.","parameters":{"language":"csharp","query":"Azure.Storage.Blobs RequestFailedException catch Status ErrorCode retry BlobClient DownloadContentAsync"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates handling Azure SDK exceptions by wrapping calls in try-catch blocks to manage common failure scenarios like 404 (not found) and 403 (forbidden) errors.\npackage: Azure\nlanguage: csharp\n","codeSnippet":"using Azure;\r\n\r\ntry\r\n{\r\n    var blob = await blobClient.DownloadAsync();\r\n}\r\ncatch (RequestFailedException ex) when (ex.Status == 404)\r\n{\r\n    // Blob not found\r\n}\r\ncatch (RequestFailedException ex) when (ex.Status == 403)\r\n{\r\n    // Insufficient permissions\r\n}\r\ncatch (RequestFailedException ex)\r\n{\r\n    _logger.LogError(ex, \u0022Azure SDK call failed with status {Status}\u0022, ex.Status);\r\n}","link":"https://learn.microsoft.com/entra/msidweb/call-downstream-apis/azure-sdks#follow-best-practices","language":"csharp"},{"description":"description: Retrieves and displays system properties of a blob such as ContentLanguage, ContentType, CreatedOn, and LastModified, handling any request failures gracefully.\npackage: Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"private static async Task GetBlobPropertiesAsync(BlobClient blob)\r\n{\r\n    try\r\n    {\r\n        // Get the blob properties\r\n        BlobProperties properties = await blob.GetPropertiesAsync();\r\n\r\n        // Display some of the blob\u0027s property values\r\n        Console.WriteLine($\u0022 ContentLanguage: {properties.ContentLanguage}\u0022);\r\n        Console.WriteLine($\u0022 ContentType: {properties.ContentType}\u0022);\r\n        Console.WriteLine($\u0022 CreatedOn: {properties.CreatedOn}\u0022);\r\n        Console.WriteLine($\u0022 LastModified: {properties.LastModified}\u0022);\r\n    }\r\n    catch (RequestFailedException e)\r\n    {\r\n        Console.WriteLine($\u0022HTTP error code {e.Status}: {e.ErrorCode}\u0022);\r\n        Console.WriteLine(e.Message);\r\n        Console.ReadLine();\r\n    }\r\n}","link":"https://learn.microsoft.com/azure/storage/blobs/storage-blob-properties-metadata#set-
... (truncated)
```

### ✅ 10. powershell — 1911ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943136341-copilot-tool-output-4e08f177d65c429d85609dfabadf2444.txt' | ConvertFrom-Json; $d.results | Where-Object { $_.title -match 'RequestFailedException|retry|error codes|403|404|429' } | Select-Object -First 10 title,contentUrl | Format-Table -AutoSize","description":"Extract relevant official documentation links"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-4e08f177d65c429d85609dfabadf2444.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 11. powershell — 1585ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943135417-copilot-tool-output-ed56178f02934717958532692c561b1f.txt' | ConvertFrom-Json; $d.results | Select-Object -First 6 description,codeSnippet,link | Format-List","description":"Extract official exception code samples"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-ed56178f02934717958532692c561b1f.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 12. view — 1327ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943136341-copilot-tool-output-4e08f177d65c429d85609dfabadf2444.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787943135417-copilot-tool-output-ed56178f02934717958532692c561b1f.txt b/C:/Users/weidxu/AppData/Local/Temp/1787943135417-copilot-tool-output-ed56178f02934717958532692c561b1f.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787943135417-copilot-tool-output-ed56178f02934717958532692c561b1f.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787943135417-copilot-tool-output-ed56178f02934717958532692c561b1f.txt
@@ -1,8 +1,8 @@
 {"results":[{"description":"description: Demonstrates handling Azure SDK exceptions by wrapping calls in try-catch blocks to manage common failure scenarios like 404 (not found) and 403 (forbidden) errors.\npackage: Azure\nlanguage: csharp\n","codeSnippet":"using Azure;\r\n\r\ntry\r\n{\r\n    var blob = await blobClient.DownloadAsync();\r\n}\r\ncatch (RequestFailedException ex) when (ex.Status == 404)\r\n{\r\n    // Blob not found\r\n}\r\ncatch (RequestFailedException ex) when (ex.Status == 403)\r\n{\r\n    // Insufficient permissions\r\n}\r\ncatch (RequestFailedException ex)\r\n{\r\n    _logger.LogError(ex, \u0022Azure SDK call failed with status {Status}\u0022, ex.Status);\r\n}","link":"https://learn.microsoft.com/entra/msidweb/call-downstream-apis/azure-sdks#follow-best-practices","language":"csharp"},{"description":"description: Retrieves and displays system properties of a blob such as ContentLanguage, ContentType, CreatedOn, and LastModified, handling any request failures gracefully.\npackage: Azure.Storage.Blobs\nlanguage: csharp\n","codeSnippet":"private static async Task GetBlobPropertiesAsync(BlobClient blob)\r\n{\r\n    try\r\n    {\r\n        // Get the blob properties\r\n        BlobProperties properties = await blob.GetPropertiesAsync();\r\n\r\n        // Display some of the blob\u0027s property values\r\n        Console.WriteLine($\u0022 ContentLanguage: {properties.ContentLanguage}\u0022);\r\n        Console.WriteLine($\u0022 ContentType: {properties.ContentTyp
... (truncated)
```

### ✅ 13. view — 994ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943135417-copilot-tool-output-ed56178f02934717958532692c561b1f.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787943136341-copilot-tool-output-4e08f177d65c429d85609dfabadf2444.txt b/C:/Users/weidxu/AppData/Local/Temp/1787943136341-copilot-tool-output-4e08f177d65c429d85609dfabadf2444.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787943136341-copilot-tool-output-4e08f177d65c429d85609dfabadf2444.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787943136341-copilot-tool-output-4e08f177d65c429d85609dfabadf2444.txt
@@ -1,8 +1,8 @@
 {"results":[{"title":"Troubleshoot 403 errors in Azure Blob Storage","content":"# Troubleshoot 403 errors in Azure Blob Storage\n## Comprehensive 403 error diagnostic checklist\nIf you receive an error other than the errors described in the previous sections, or if the causes and solutions described in those sections don\u0027t help you resolve your error, use this checklist as a guide to identify the cause and find a solution.\n1. **Secure file transfer**: If secure transfer is required, make sure that requests are made over HTTPS. (**[Learn more](https://learn.microsoft.com/en-us/azure/storage/common/storage-require-secure-transfer)**).\n2. **Azure role-based access control**: Verify the security principal has the correct Azure RBAC role. (**[Learn more](https://learn.microsoft.com/en-us/azure/storage/blobs/assign-azure-role-data-access)**).\n3. **Shared access signature (SAS) tokens**: Check token expiration, permissions, and generation method. (**Read more**).\n4. **User delegation SAS**: Confirm required fields, permissions, and key status for user delegation SAS. (**Read more**).\n5. **Stored access policies**: Review stored access policy settings and the propagation delay. (**Read more**).\n6. **Access control lists (ACLs)**: Ensure ACL entries and permissions are correct for the client. (**Read more**).\n7. **Shared Key authorization**: If you\u0027re using an account key, confirm that Shared Key authorization is allowed. (**[Learn more](https://learn.microsoft.com/en-
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 3290ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official RequestFailedException API and Azure SDK error handling documentation.","parameters":{"query":"site:learn.microsoft.com/dotnet/api Azure.RequestFailedException Status ErrorCode GetRawResponse"}}
```

**Output:**

```
{"results":[{"title":"Status and error codes","content":"# Status and error codes\nREST API operations for Azure Storage return standard HTTP status codes, as defined in the [HTTP/1.1 Status Code Definitions](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).\nAPI operations can also return detailed error information that the service defines. This error information is returned in the response body.\nThe body of the error response follows the basic format shown here. The \u0060code\u0060 element indicates the error code string, whereas the \u0060message\u0060 element indicates the error message string.\n\u0060\u0060\u0060\n\u003C?xml version=\u00221.0\u0022 encoding=\u0022utf-8\u0022?\u003E  \n\u003CError\u003E  \n  \u003CCode\u003Estring-value\u003C/Code\u003E  \n  \u003CMessage\u003Estring-value\u003C/Message\u003E  \n\u003C/Error\u003E\n\u0060\u0060\u0060\nFor version 2017-07-29 and later, failed API operations also return the Azure Storage error code string in a response header.\n| Response header | Description| \n|  --- | ---  |\n| \u0060x-ms-error-code\u0060 | A value that indicates the error code string. |\n\nA particular error response might also include detailed error information to provide the developer with more information. For example, the following error response indicates an invalid query parameter specified on the request URI, and provides additional information about the invalid parameter\u0027s name and value, and the reason for the error.\nFor version 2024-08-04 and later, error response will include detailed error information when Account SAS\u0027s signed resource type doesn\u0027t match the resource being accessed. For example, the following error response indicates the signed resource type \u0027o\u0027 is required to perform the operation.\nThe following list of articles provide collections of error codes that you might find useful:\n1. [Common REST API error codes](https://learn.microsoft.com/rest/api/storageservices/common-rest-api-er
... (truncated)
```

### ✅ 15. azure-documentation (via azure) — 2969ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure SDK for .NET retry configuration documentation.","parameters":{"query":"site:learn.microsoft.com/dotnet/azure/sdk retry MaxRetries Delay MaxDelay NetworkTimeout RetryMode Azure.Core RetryOptions"}}
```

**Output:**

```
{"results":[{"title":"RequestFailedException.GetRawResponse Method-Definition","content":"# RequestFailedException.GetRawResponse Method\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure](https://learn.microsoft.com/dotnet/api/azure?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.Core.dll\r\n\r\n- Package:\r\n    - Azure.Core v1.62.0\r\n\r\n- Source:\r\n    - [RequestFailedException.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/RequestFailedException.cs#L159C46-L159C55)\r\n\r\nGets the response, if any, that led to the exception.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic Azure.Response? GetRawResponse();\r\n\u0060\u0060\u0060\r\n\r\n#### Returns\r\n\r\n[Response](https://learn.microsoft.com/dotnet/api/azure.response?view=azure-dotnet)","contentUrl":"https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.getrawresponse?view=azure-dotnet"},{"title":"RequestFailedException.Status Property-Definition","content":"# RequestFailedException.Status Property\r\n\r\n## Definition\r\n\r\n- Namespace:\r\n    - [Azure](https://learn.microsoft.com/dotnet/api/azure?view=azure-dotnet)\r\n\r\n- Assembly:\r\n    - Azure.Core.dll\r\n\r\n- Package:\r\n    - Azure.Core v1.62.0\r\n\r\n- Source:\r\n    - [RequestFailedException.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/RequestFailedException.cs#L29C29-L29C33)\r\n\r\nGets the HTTP status code of the response. Returns. \u00600\u0060 if response was not received.\r\n\r\n\u0060\u0060\u0060csharp\r\npublic int Status { get; }\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n[Int32](https://learn.microsoft.com/en-us/dotnet/api/system.int32)","contentUrl":"https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.status?view=azure-dotnet"},{"title":"RequestFailedException.ErrorCode Property-Definition","content":"# RequestFailedException.ErrorCode Property\r\n\r
... (truncated)
```

### ✅ 16. azure-documentation (via azure) — 2643ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure Blob Storage error codes for missing container and throttling.","parameters":{"query":"site:learn.microsoft.com/rest/api/storageservices ContainerNotFound ServerBusy 429 Blob service error codes"}}
```

**Output:**

```
{"results":[{"title":"Implement a retry policy with .NET","content":"# Implement a retry policy with .NET\n## Configure retry options\n| Property | Type | Description | Default value| \n|  --- | --- | --- | ---  |\n| [Delay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.delay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The delay between retry attempts for a fixed approach or the delay on which to base calculations for a backoff-based approach. If the service provides a Retry-After response header, the next retry is delayed by the duration specified by the header value. | 0.8 second |\n| [MaxDelay](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.maxdelay) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The maximum permissible delay between retry attempts when the service doesn\u0027t provide a Retry-After response header. If the service provides a Retry-After response header, the next retry is delayed by the duration specified by the header value. | 1 minute |\n| [MaxRetries](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.maxretries) | int | The maximum number of retry attempts before giving up. | 5 (see note) |\n| [Mode](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.mode) | [RetryMode](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retrymode) | The approach to use for calculating retry delays. | Exponential |\n| [NetworkTimeout](https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions.networktimeout) | [TimeSpan](https://learn.microsoft.com/en-us/dotnet/api/system.timespan) | The timeout applied to an individual network operation. | 100 seconds |\n\nNote\n\u0060StorageClientOptions\u0060 increases the default value for \u0060MaxRetries\u0060 from 3 to 5. All other properties have the same default values as \u0060RetryOptions\u0060.\nIn this code example for Blob Storage, we configure the retry op
... (truncated)
```

### ✅ 17. powershell — 2499ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943157089-copilot-tool-output-9f9fac1de53e4b7fb0fefa19d8e2fbbd.txt' -First 1 | ConvertFrom-Json); $d.results | Select-Object -First 5 title,contentUrl | Format-Table -AutoSize","description":"Extract retry documentation links"}
```

**Output:**

```
title                                                      contentUrl
-----                                                      ----------
RequestFailedException.GetRawResponse Method-Definition    https://learn.microsoft.com/dotnet/api/azure.requestfailede…
RequestFailedException.Status Property-Definition          https://learn.microsoft.com/dotnet/api/azure.requestfailede…
RequestFailedException.ErrorCode Property-Definition       https://learn.microsoft.com/dotnet/api/azure.requestfailede…
RequestFailedException Constructors-Definition             https://learn.microsoft.com/dotnet/api/azure.requestfailede…
NullableResponse&lt;T&gt;.GetRawResponse Method-Definition https://learn.microsoft.com/dotnet/api/azure.nullablerespon…
RequestFailedException Constructors-Definition             https://learn.microsoft.com/dotnet/api/azure.requestfailede…
RequestFailedException Constructors-Definition             https://learn.microsoft.com/dotnet/api/azure.requestfailede…
RequestFailedException Constructors-Definition             https://learn.microsoft.com/dotnet/api/azure.requestfailede…
<shellId: 3 completed with exit code 0>
```

### ✅ 18. powershell — 2159ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943157216-copilot-tool-output-3a6c5cba27d04ace91d1de5684d4a11a.txt' -First 1 | ConvertFrom-Json); $d.results | Select-Object -First 8 title,contentUrl | Format-Table -AutoSize","description":"Extract exception API links"}
```

**Output:**

```
title                                                            contentUrl
-----                                                            ----------
Status and error codes                                           https://learn.microsoft.com/rest/api/storageservices/…
Troubleshoot client application errors in Azure storage accounts https://learn.microsoft.com/troubleshoot/azure/azure-…
Azure Storage Blobs client library for .NET - version 12.29.2    https://learn.microsoft.com/dotnet/api/overview/azure…
REST API Status Codes and Error Handling                         https://learn.microsoft.com/bingmaps/rest-services/st…
Common Error Codes for Storage Resource Provider                 https://learn.microsoft.com/rest/api/storagerp/srp_er…
Common REST API error codes                                      https://learn.microsoft.com/rest/api/storageservices/…
Troubleshoot client application errors in Azure storage accounts https://learn.microsoft.com/troubleshoot/azure/azure-…
Troubleshoot client application errors in Azure storage accounts https://learn.microsoft.com/troubleshoot/azure/azure-…
<shellId: 4 completed with exit code 0>
```

### ✅ 19. powershell — 1831ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943157647-copilot-tool-output-e6b500452d8d4f74ac3017bacf7a8e66.txt' -First 1 | ConvertFrom-Json); $d.results | Select-Object -First 8 title,contentUrl | Format-Table -AutoSize","description":"Extract storage error code links"}
```

**Output:**

```
title                                                      contentUrl
-----                                                      ----------
Implement a retry policy with .NET                         https://learn.microsoft.com/azure/storage/blobs/storage-ret…
Manage storage account resources with .NET                 https://learn.microsoft.com/azure/storage/common/storage-sr…
RetryOptions.Mode Property-Definition                      https://learn.microsoft.com/dotnet/api/azure.core.retryopti…
RetryOptions.MaxRetries Property-Definition                https://learn.microsoft.com/dotnet/api/azure.core.retryopti…
Handle timeouts and configure retries in Azure Service Bus https://learn.microsoft.com/azure/service-bus-messaging/ser…
<shellId: 2 completed with exit code 0>
```

### ✅ 20. powershell — 2541ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943157089-copilot-tool-output-9f9fac1de53e4b7fb0fefa19d8e2fbbd.txt' -First 1 | ConvertFrom-Json); $d.results | Select-Object -First 1 title,contentUrl | Format-List","description":"Show full retry documentation URL"}
```

**Output:**

```
title      : Status and error codes
contentUrl : https://learn.microsoft.com/rest/api/storageservices/status-and-error-codes2
<shellId: 7 completed with exit code 0>
```

### ✅ 21. powershell — 2215ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943157216-copilot-tool-output-3a6c5cba27d04ace91d1de5684d4a11a.txt' -First 1 | ConvertFrom-Json); $d.results | Select-Object -First 3 title,contentUrl | Format-List","description":"Show full exception API URLs"}
```

**Output:**

```
title      : Implement a retry policy with .NET
contentUrl : https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy#configure-retry-options
<shellId: 5 completed with exit code 0>
```

### ✅ 22. powershell — 1896ms

**Input:**

```json
{"command":"$d = (Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787943157647-copilot-tool-output-e6b500452d8d4f74ac3017bacf7a8e66.txt' -First 1 | ConvertFrom-Json); $d.results | Select-Object -First 1 title,contentUrl | Format-List","description":"Show full storage error URL"}
```

**Output:**

```
title      : RequestFailedException.GetRawResponse Method-Definition
contentUrl : https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.getrawresponse?view=azure-dotnet

title      : RequestFailedException.Status Property-Definition
contentUrl : https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.status?view=azure-dotnet

title      : RequestFailedException.ErrorCode Property-Definition
contentUrl : https://learn.microsoft.com/dotnet/api/azure.requestfailedexception.errorcode?view=azure-dotnet
<shellId: 6 completed with exit code 0>
```

### ✅ 23. azure-documentation (via azure) — 3170ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official retry policy guidance for Blob Storage .NET.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy"}}
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

### ✅ 24. azure-documentation (via azure) — 2839ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete RequestFailedException API documentation.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/azure.requestfailedexception?view=azure-dotnet"}}
```

**Output:**

```
# RequestFailedException Class

## Definition

- Namespace:
    - [Azure](azure?view=azure-dotnet)

- Assembly:
    - Azure.Core.dll

- Package:
    - Azure.Core v1.62.0

- Source:
    - [RequestFailedException.cs](https://github.com/Azure/azure-sdk-for-net/blob/86e9f4dd0218b07ec408b9452893f9efea974ff8/sdk/core/Azure.Core/src/RequestFailedException.cs)

An exception thrown when service request fails.

```csharp
[System.Serializable]
public class RequestFailedException : Exception
```

```fsharp
[<System.Serializable>]
type RequestFailedException = class
    inherit Exception
    interface ISerializable
```

```vb
Public Class RequestFailedException
Inherits Exception
```

- Inheritance
    - [Exception](/en-us/dotnet/api/system.exception)
RequestFailedException

- Derived
    - [Azure.AI.AgentServer.Core.Storage.FoundryStorageException](azure.ai.agentserver.core.storage.foundrystorageexception?view=azure-dotnet)

[Azure.Connectors.Sdk.ConnectorException](azure.connectors.sdk.connectorexception?view=azure-dotnet)

[Azure.Data.Tables.TableTransactionFailedException](azure.data.tables.tabletransactionfailedexception?view=azure-dotnet)

- Attributes
    - [SerializableAttribute](/en-us/dotnet/api/system.serializableattribute)

- Implements
    - [ISerializable](/en-us/dotnet/api/system.runtime.serialization.iserializable)

## Constructors

| Name | Description |
| --- | --- |
| [RequestFailedException(Response, Exception, RequestFailedDetailsParser)](azure.requestfailedexception.-ctor?view=azure-dotnet#azure-requestfailedexception-ctor%28azure-response-system-exception-azure-core-requestfaileddetailsparser%29) | Initializes a new instance of the [RequestFailedException](azure.requestfailedexception?view=azure-dotnet) class with an error message, HTTP status code, and error code obtained from the specified response. |
| [RequestFailedException(Response, Exception)](azure.requestfailedexception.-ctor?view=azure-dotn
... (truncated)
```

### ✅ 25. azure-documentation (via azure) — 2502ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Blob Storage service error codes.","parameters":{"url":"https://learn.microsoft.com/rest/api/storageservices/blob-service-error-codes"}}
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

## Copilot Response

I’ll ground the patterns in the current Azure SDK guidance and verify the precise exception fields and retry behavior.

The key distinction is between HTTP status (broad category) and Storage service error code (specific cause); I’m narrowing the examples around that distinction and built-in retry semantics.

`Azure.Storage.Blobs` reports failed service requests as `Azure.RequestFailedException`. Its most useful properties are:

- `Status`: HTTP status code, or `0` if no response was received.
- `ErrorCode`: service-specific code such as `ContainerNotFound`; it may be `null`.
- `GetRawResponse()`: response headers, including `x-ms-request-id`.

Creating a `BlobContainerClient` or `BlobClient` does not send a request, so exceptions occur only when an operation such as `DownloadContentAsync` executes.

```csharp
using Azure;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Microsoft.Extensions.Logging;

public static async Task<string?> DownloadAsync(
    BlobContainerClient containerClient,
    string blobName,
    ILogger logger,
    CancellationToken cancellationToken)
{
    BlobClient blobClient = containerClient.GetBlobClient(blobName);

    try
    {
        BlobDownloadResult result =
            await blobClient.DownloadContentAsync(cancellationToken);

        return result.Content.ToString();
    }
    catch (RequestFailedException ex)
        when (ex.Status == 404 &&
              ex.ErrorCode == BlobErrorCode.ContainerNotFound.ToString())
    {
        logger.LogWarning(
            "Container was not found. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 404)
    {
        // A download can also return BlobNotFound.
        logger.LogWarning(
            "Blob resource was not found. Status={Status}, ErrorCode={ErrorCode}",
            ex.Status,
            ex.ErrorCode);

        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        // ErrorCode identifies the specific cause, for example:
        // AuthorizationPermissionMismatch, AuthorizationFailure,
        // AuthenticationFailed, or UnauthorizedBlobOverwrite.
        LogStorageFailure(logger, ex, "Blob access denied");
        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        // The SDK has already applied its configured retry policy before
        // this exception reaches application code.
        LogStorageFailure(logger, ex, "Blob request throttled");
        throw;
    }
    catch (RequestFailedException ex)
    {
        LogStorageFailure(logger, ex, "Blob Storage request failed");
        throw;
    }
}

private static void LogStorageFailure(
    ILogger logger,
    RequestFailedException exception,
    string description)
{
    Response? response = exception.GetRawResponse();

    string? requestId = null;
    response?.Headers.TryGetValue("x-ms-request-id", out requestId);

    logger.LogError(
        exception,
        "{Description}. Status={Status}, ErrorCode={ErrorCode}, RequestId={RequestId}",
        description,
        exception.Status,
        exception.ErrorCode,
        requestId);
}
```

For status-only dispatch, a `switch` is also idiomatic:

```csharp
catch (RequestFailedException ex)
{
    switch (ex.Status)
    {
        case 404:
            // ex.ErrorCode distinguishes ContainerNotFound from BlobNotFound.
            break;

        case 403:
            // Permanent until credentials, RBAC, SAS, firewall, or ACLs change.
            throw;

        case 429:
            // Transient throttling after SDK retries are exhausted.
            throw;

        default:
            throw;
    }
}
```

### Expected failure information

| Scenario | `Status` | Typical `ErrorCode` | Treatment |
|---|---:|---|---|
| Container missing | `404` | `ContainerNotFound` | Often translate to a domain “not found” result |
| Blob missing | `404` | `BlobNotFound` | Handle separately if blob absence is expected |
| Access denied | `403` | Varies, such as `AuthorizationPermissionMismatch` or `AuthorizationFailure` | Log and propagate; retrying unchanged credentials normally cannot help |
| Throttling | `429` | Service-dependent and potentially `null` | Allow the SDK retry policy to handle it first |

Blob Storage data-plane throttling is also commonly returned as `503 ServerBusy` or occasionally `500 OperationTimedOut`, rather than only `429`. If application logic classifies transient failures, account for those statuses too.

### Configure SDK retries

The client retries transient responses automatically and honors server-provided retry delays. Configure that policy when constructing the service client rather than putting an unconditional retry loop around each operation:

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
    new DefaultAzureCredential(), // Local development
    options);
```

Use a managed identity credential directly in Azure-hosted production applications. Avoid logging credentials, SAS query strings, authorization headers, or full signed URLs.

**References:** [RequestFailedException](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception?view=azure-dotnet), [Blob Storage retry policy](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), and [Blob Storage error codes](https://learn.microsoft.com/rest/api/storageservices/blob-service-error-codes).

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 6 criteria passed. The response thoroughly covers the prompt's requested 404, 403, and 429 scenarios, exception details, retry configuration, and diagnostics, but omits 409 handling and conditional failures involving ETags or leases.

### Strengths

- Uses RequestFailedException idiomatically with exception filters and a status switch.
- Clearly extracts and logs Status, ErrorCode, and the server request ID.
- Provides practical BlobClientOptions retry configuration and sensible guidance about transient failures.
- Distinguishes ContainerNotFound from BlobNotFound and notes additional Blob Storage throttling responses.

### Issues

- No handling or explanation for HTTP 409 Conflict.
- No coverage of ETag or lease-related conditional request failures, including relevant 409/412 responses and service error codes.

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
hyoka run --prompt-id storage-dp-dotnet-error-handling --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
