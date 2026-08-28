# Evaluation Report: storage-mp-python-account-mgmt

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 814.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-mp-python-account-mgmt` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 13/15 |
| Duration | 814.7s |
| Timestamp | 2026-08-28T19:28:41Z |
| Files Generated | 2 |
| Event Count | 121379 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 613.1s |
| Review | 196.5s |
| **Total** | **814.7s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=483183 out=6009 |
| Turn Count | 11 |

## Error

```
evaluation failed: sending prompt: waiting for session.idle: context deadline exceeded
```

**Details:**

```
sending prompt: waiting for session.idle: context deadline exceeded
```

## Tool Calls

### ✅ 1. skill — 747ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 5976ms

**Input:**

```json
{"command":"get_best_practices","intent":"Write a Python script using DefaultAzureCredential and azure-mgmt-storage to create, list, inspect, update blob versioning, and delete a Storage Account, with proper error handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. glob — 5555ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 6767ms

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

### ✅ 5. view — 6396ms

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

### ✅ 6. azure-documentation (via azure) — 7164ms

**Input:**

```json
{"command":"search","intent":"Find official Python azure-mgmt-storage examples and API signatures for creating a storage account, listing by resource group, getting account properties, enabling blob versioning through blob service properties, and deleting the account."}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. glob — 5142ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-1193536646","pattern":"*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. azure-get_azure_bestpractices (via azure) — 2731ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Python code-generation best practices for a Storage Account management-plane SDK script.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3578ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python management-plane SDK samples for Azure Storage Account lifecycle and blob versioning.","parameters":{"language":"python","query":"Python azure-mgmt-storage StorageManagementClient storage_accounts begin_create list_by_resource_group get_properties delete blob_services set_service_properties BlobServiceProperties is_versioning_enabled"}}
```

**Output:**

```
{"results":[{"title":"BlobServicesOperations Class-Methods","content":"# BlobServicesOperations Class\r\n\r\nWarning\r\n\r\nDO NOT instantiate this class directly.\r\n\r\nInstead, you should access the following operations through\r\n\r\n[StorageManagementClient](https://learn.microsoft.com/python/api/azure-mgmt-storage/azure.mgmt.storage.storagemanagementclient?view=azure-python)\u0027s\r\n\r\n\u0026lt;xref:blob\\_services\u0026gt; attribute.\r\n\r\n## Methods\r\n\r\n### set\\_service\\_properties\r\n\r\nSets the properties of a storage account\u0027s Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.\r\n\r\n\u0060\u0060\u0060python\r\nset_service_properties(resource_group_name: str, account_name: str, parameters: _models.BlobServiceProperties, *, content_type: str = \u0027application/json\u0027, **kwargs: Any) -\u003E _models.BlobServiceProperties\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r\n| resource\\_group\\_name\u003Cbr\u003E\u003Cbr\u003ERequired | [str](https://docs.python.org/3/library/stdtypes.html#str)\u003Cbr\u003E\u003Cbr\u003EThe name of the resource group. The name is case insensitive. Required. |\r\n| account\\_name\u003Cbr\u003E\u003Cbr\u003ERequired | [str](https://docs.python.org/3/library/stdtypes.html#str)\u003Cbr\u003E\u003Cbr\u003EThe name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Required. |\r\n| parameters\u003Cbr\u003E\u003Cbr\u003ERequired | [BlobServiceProperties](https://learn.microsoft.com/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobserviceproperties?view=azure-python) or \u0026lt;xref:JSON\u0026gt; or [IO](https://docs.python.org/3/library/typing.html#typing.IO)[[bytes](https://docs.python.org/3/library/stdtypes.html#bytes)] \u003Cbr\u003E\u003Cbr\u003EThe properties of a storage account\
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 3157ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official API reference for Python azure-mgmt-storage StorageManagementClient operations used to manage accounts and blob service versioning.","parameters":{"query":"site:learn.microsoft.com/python/api/azure-mgmt-storage StorageAccountsOperations BlobServicesOperations set_service_properties Python"}}
```

**Output:**

```
{"results":[{"description":"description: Imports Azure SDK modules and creates a StorageManagementClient instance authenticated with Azure CLI credentials to manage Azure storage resources in Python.\npackage: azure-mgmt-storage, azure-storage-blob\nlanguage: python\n","codeSnippet":"from azure.mgmt.storage import StorageManagementClient\r\nfrom azure.storage.blob import ContainerClient\r\nfrom azure.mgmt.storage.models import Sku, StorageAccountCreateParameters, BlobContainer\r\n\r\ncredential = AzureCliCredential()\r\nstorage_client = StorageManagementClient(\r\n    credential=credential, subscription_id=subscription_id\r\n)","link":"https://learn.microsoft.com/azure/machine-learning/how-to-access-resources-from-endpoints-managed-identities?view=azureml-api-2#create-storage-account-and-container","language":"python"},{"description":"description: Imports Azure SDK modules and creates a StorageManagementClient instance authenticated with Azure CLI credentials to manage Azure storage resources in Python (user-assigned identity scenario).\npackage: azure-mgmt-storage, azure-storage-blob\nlanguage: python\n","codeSnippet":"from azure.mgmt.storage import StorageManagementClient\r\nfrom azure.storage.blob import ContainerClient\r\nfrom azure.mgmt.storage.models import Sku, StorageAccountCreateParameters, BlobContainer\r\n\r\ncredential = AzureCliCredential()\r\nstorage_client = StorageManagementClient(\r\n    credential=credential, subscription_id=subscription_id\r\n)","link":"https://learn.microsoft.com/azure/machine-learning/how-to-access-resources-from-endpoints-managed-identities?view=azureml-api-2#create-storage-account-and-container","language":"python"},{"description":"description: This Python script uses the Azure SDK for Python management libraries to create a resource group, an Azure Storage account, and a Blob container. It authenticates using DefaultAzureCredential, checks for storage account name availability, provisions the storage account, retrieves the ac
... (truncated)
```

### ✅ 11. rg — 1502ms

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787945379006-copilot-tool-output-5f3594c17d0741beb232526c2674c6c9.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787945379463-copilot-tool-output-fa05f210c40e426e9b4602c47035426b.txt"],"pattern":"set_service_properties|is_versioning_enabled|begin_create|list_by_resource_group|get_properties|begin_delete|\\.delete\\("}
```

**Output:**

```
Output too large to read at once (69.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787945385827-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787945379463-copilot-tool-output-fa05f210c40e426e9b4602c47035426b.txt:{"results":[{"title":"BlobServicesOperations Class-Methods","content":"# BlobServicesOperations Class\r\n\r\nWarning\r\n\r\nDO NOT instantiate this class directly.\r\n\r\nInstead, you should access the following operations through\r\n\r\n[StorageManagementClient](https://learn.microsoft.com/python/api/azure-mgmt-storage/azure.mgmt.storage.storagemanagementclient?view=azure-python)\u0027s\r\n
```

### ✅ 12. rg — 1167ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787945379006-copilot-tool-output-5f3594c17d0741beb232526c2674c6c9.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787945379463-copilot-tool-output-fa05f210c40e426e9b4602c47035426b.txt"],"pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

**Output:**

```
Output too large to read at once (69.5 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787945385827-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787945379006-copilot-tool-output-5f3594c17d0741beb232526c2674c6c9.txt:{"results":[{"description":"description: Imports Azure SDK modules and creates a StorageManagementClient instance authenticated with Azure CLI credentials to manage Azure storage resources in Python.\npackage: azure-mgmt-storage, azure-storage-blob\nlanguage: python\n","codeSnippet":"from azure.mgmt.storage import StorageManagementClient\r\nfrom azure.storage.blob import ContainerClient\r\nf
```

### ✅ 13. powershell — 1255ms

**Input:**

```json
{"command":"$paths = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787945379006-copilot-tool-output-5f3594c17d0741beb232526c2674c6c9.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787945379463-copilot-tool-output-fa05f210c40e426e9b4602c47035426b.txt'); foreach ($path in $paths) { $json = Get-Content -Raw $path | ConvertFrom-Json; $json.results | Where-Object { ($_.title + $_.description + $_.content + $_.codeSnippet) -match 'set_service_properties|is_versioning_enabled|begin_create|list_by_resource_group|delete' } | Select-Object -First 8 | ForEach-Object { 'TITLE: ' + $_.title; 'URL: ' + ($_.contentUrl ?? $_.url); if ($_.codeSnippet) { 'CODE: ' + $_.codeSnippet }; if ($_.content) { 'CONTENT: ' + $_.content.Substring(0, [Math]::Min(1800, $_.content.Length)) }; '---' } }","description":"Extract relevant Azure SDK documentation","initial_wait":30,"mode":"sync"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … n $paths) { $json = Get-Content -Raw $path | ConvertFrom-Json; $json. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 14. web_fetch — 3018ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.blobservicesoperations?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.operations.StorageAccountsOperations class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.operations.StorageAccountsOperations
module: azure.mgmt.storage.operations
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-storage
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
locale: en-us
document_id: 69f78b6c-d0a2-14c5-fcc6-ae2e905de74c
document_version_independent_id: ee5682da-04bc-59a2-f87c-fb774903a132
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.operations.StorageAccountsOperations.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.operations.StorageAccountsOperations.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: " Warning DO NOT instantiate this class directly. Instead, you should access the following operations through StorageManagementClient's &lt;xref:stora
... (truncated)
```

### ✅ 15. web_fetch — 2619ms

**Input:**

```json
{"max_length":16000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobserviceproperties?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.BlobServiceProperties class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobserviceproperties?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.BlobServiceProperties
module: azure.mgmt.storage.models
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-storage
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
locale: en-us
document_id: 9070b2e5-9376-276e-9bbb-ebdb45c88921
document_version_independent_id: 3131be1e-da96-d915-f25e-8c26e12906cd
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.BlobServiceProperties.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.BlobServiceProperties.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: "The properties of a storage account's Blob service. "
toc_rel: ../_splitted/azure.mgmt.storage/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: a
... (truncated)
```

### ✅ 16. web_fetch — 2230ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobserviceproperties?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.blobservicesoperations?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.operations.BlobServicesOperations class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.blobservicesoperations?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.operations.BlobServicesOperations
module: azure.mgmt.storage.operations
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-storage
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
locale: en-us
document_id: 4de5d10e-1a60-0ec1-80bf-6213d87e1097
document_version_independent_id: eef1d0a4-539f-6fb0-a90c-da2b8d2b4a47
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.operations.BlobServicesOperations.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.operations.BlobServicesOperations.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: " Warning DO NOT instantiate this class directly. Instead, you should access the following operations through StorageManagementClient's &lt;xref:blob_services&gt; attri
... (truncated)
```

### ✅ 17. web_fetch — 2895ms

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobservicepropertiesproperties?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountcreateparameters?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.StorageAccountCreateParameters class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountcreateparameters?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.StorageAccountCreateParameters
module: azure.mgmt.storage.models
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-storage
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
locale: en-us
document_id: f5a450b6-5cc6-aaa3-101d-682acf8dff44
document_version_independent_id: d53944fd-b87c-e8e1-4295-30e7649574c5
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.StorageAccountCreateParameters.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.StorageAccountCreateParameters.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'The parameters used when creating a storage account. '
toc_rel: ../_splitted/azure.mgmt.storage/toc.json
feedback_help_link_type: ''
feedback_help_
... (truncated)
```

### ✅ 18. web_fetch — 2539ms

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountcreateparameters?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobservicepropertiesproperties?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.BlobServicePropertiesProperties class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobservicepropertiesproperties?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.BlobServicePropertiesProperties
module: azure.mgmt.storage.models
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-storage
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/de8ce683-cbe1-461b-bae7-77db0888ec6d
locale: en-us
document_id: 87b74065-0339-6072-c2fc-6cdf2900b8fb
document_version_independent_id: f6ae2810-554d-8d2a-9639-d26136bad15c
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.BlobServicePropertiesProperties.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.BlobServicePropertiesProperties.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: "The properties of a storage account's Blob service. "
toc_rel: ../_splitted/azure.mgmt.storage/toc.json
feedback_help_link_type: ''
feedback_
... (truncated)
```

### ✅ 19. web_fetch — 2184ms

**Input:**

```json
{"max_length":14000,"start_index":30000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python:
resource\_group\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the resource group. The name is case insensitive. Required. |
| account\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Required. |
| regenerate\_key<br><br>Required | [StorageAccountRegenerateKeyParameters](azure.mgmt.storage.models.storageaccountregeneratekeyparameters) or &lt;xref:JSON&gt; or [IO](https://docs.python.org/3/library/typing.html#typing.IO)[[bytes](https://docs.python.org/3/library/stdtypes.html#bytes)] <br><br>Specifies name of the key which should be regenerated – key1, key2, kerb1, kerb2. Is one of the following types: StorageAccountRegenerateKeyParameters, JSON, IO[bytes] Required. |

#### Returns

| Type | Description |
| --- | --- |
| [StorageAccountListKeysResult](azure.mgmt.storage.models.storageaccountlistkeysresult) | StorageAccountListKeysResult. The StorageAccountListKeysResult is compatible with MutableMapping |

#### Exceptions

| Type | Description |
| --- | --- |
| [HttpResponseError](../azure-core/azure.core.exceptions.httpresponseerror) |  |

### revoke\_user\_delegation\_keys

Revoke user delegation keys.

```python
revoke_user_delegation_keys(resource_group_name: str, account_name: str, **kwargs: Any) -> None
```

#### Parameters

| Name | Description |
| --- | --- |
| resource\_group\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the resource group. The name is case insensitive. Required. |
| account\_name<br><br>Required | [str](https://docs.python.org/3/librar
... (truncated)
```

### ✅ 20. web_fetch — 2774ms

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-identity/json"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python:
python.org/3/library/constants.html#None)] | An instance of LROPoller that returns None |

#### Exceptions

| Type | Description |
| --- | --- |
| [HttpResponseError](../azure-core/azure.core.exceptions.httpresponseerror) |  |

### begin\_restore\_blob\_ranges

Restore blobs in the specified blob ranges.

```python
begin_restore_blob_ranges(resource_group_name: str, account_name: str, parameters: _models.BlobRestoreParameters, *, content_type: str = 'application/json', **kwargs: Any) -> LROPoller[_models.BlobRestoreStatus]
```

#### Parameters

| Name | Description |
| --- | --- |
| resource\_group\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the resource group. The name is case insensitive. Required. |
| account\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Required. |
| parameters<br><br>Required | [BlobRestoreParameters](azure.mgmt.storage.models.blobrestoreparameters) or &lt;xref:JSON&gt; or [IO](https://docs.python.org/3/library/typing.html#typing.IO)[[bytes](https://docs.python.org/3/library/stdtypes.html#bytes)] <br><br>The parameters to provide for restore blob ranges. Is one of the following types: BlobRestoreParameters, JSON, IO[bytes] Required. |

#### Returns

| Type | Description |
| --- | --- |
| [LROPoller](../azure-core/azure.core.polling.lropoller)[[BlobRestoreStatus](azure.mgmt.storage.models.blobrestorestatus)] | An instance of LROPoller that returns BlobRestoreStatus. The BlobRestoreStatus is compatible with MutableMapping |

#### Exceptions

| Type | Description |
| --- | --- |
| [HttpR
... (truncated)
```

### ✅ 21. web_fetch — 2379ms

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-mgmt-storage/json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-identity/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com> License-Expression: MIT","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.14","Programming Language :: Python :: 3.9"],"description":"# Azure Identity client library for Python\n\nThe Azure Identity library provides [Microsoft Entra ID](https://learn.microsoft.com/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [`TokenCredential`][token_cred_ref]/[`SupportsTokenInfo`][supports_token_info_ref] implementations, which can be used to construct Azure SDK clients that support Microsoft Entra token authent

<note>Content truncated. Call the fetch tool with a start_index of 1000 to get more content.</note>
```

### ✅ 22. web_fetch — 1995ms

**Input:**

```json
{"max_length":10000,"start_index":16000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-mgmt-storage/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com>","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13"],"description":"# Microsoft Azure SDK for Python\n\nThis is the Microsoft Azure Storage Management Client Library.\nThis package has been tested with Python 3.10+.\nFor a more complete view of Azure libraries, see the [azure sdk python release](https://aka.ms/azsdk/python/all).\n\n## _Disclaimer_\n\n_Azure SDK Python packages support for Python 2.7 has ended 01 January 2022. For more information and questions, please refer to https://github.com/Azure/azure-sdk-for-python/issues/20691_\n\n## Getting started\n\n### Prerequisites\n\n- Python 3.1

<note>Content truncated. Call the fetch tool with a start_index of 1000 to get more content.</note>
```

### ✅ 23. powershell — 2020ms

**Input:**

```json
{"command":"$identity = Invoke-RestMethod 'https://pypi.org/pypi/azure-identity/json'; $storage = Invoke-RestMethod 'https://pypi.org/pypi/azure-mgmt-storage/json'; \"azure-identity=$($identity.info.version)\"; \"azure-mgmt-storage=$($storage.info.version)\"","description":"Check latest Azure SDK package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-identity=1.25.3
azure-mgmt-storage=25.1.0
<shellId: 1 completed with exit code 0>
```

### ✅ 24. apply_patch — 750ms

**Input:**

```json
"*** Begin Patch\n*** Add File: storage_account_manager.py\n+\"\"\"Create, inspect, configure, and delete an Azure Storage Account.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import logging\n+import os\n+import re\n+\n+from azure.core.exceptions import AzureError, ClientAuthenticationError, HttpResponseError\n+from azure.identity import DefaultAzureCredential\n+from azure.mgmt.storage import StorageManagementClient\n+from azure.mgmt.storage.models import (\n+    BlobServicePropertiesProperties,\n+    Kind,\n+    Sku,\n+    StorageAccountCreateParameters,\n+)\n+\n+LOCATION = \"eastus\"\n+STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r\"^[a-z0-9]{3,24}$\")\n+\n+logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+logger = logging.getLogger(__name__)\n+\n+\n+def parse_args() -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(\n+        description=\"Manage the lifecycle of an Azure Storage Account.\"\n+    )\n+    parser.add_argument(\n+        \"--subscription-id\",\n+        default=os.getenv(\"AZURE_SUBSCRIPTION_ID\"),\n+        help=\"Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).\",\n+    )\n+    parser.add_argument(\n+        \"--resource-group\",\n+        default=os.getenv(\"AZURE_RESOURCE_GROUP\"),\n+        help=\"Existing resource group (defaults to AZURE_RESOURCE_GROUP).\",\n+    )\n+    parser.add_argument(\n+        \"--account-name\",\n+        default=os.getenv(\"AZURE_STORAGE_ACCOUNT_NAME\"),\n+        help=(\n+            \"Globally unique account name (defaults to \"\n+            \"AZURE_STORAGE_ACCOUNT_NAME).\"\n+        ),\n+    )\n+    return parser.parse_args()\n+\n+\n+def validate_args(args: argparse.Namespace) -\u003e None:\n+    missing = [\n+        name\n+        for name, value in (\n+            (\"subscription ID\", args.subscription_id),\n+            (\"resource group\", args.resource_group),\n+            (\"account name\", args.account_name),\n+        )\n+        if not value\n+    ]\n+    if missing:\n+        raise ValueError(f\"Missing required configuration: {', '.join(missing)}\")\n+\n+    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(args.account_name):\n+        raise ValueError(\n+            \"The account name must contain 3-24 lowercase letters and numbers.\"\n+        )\n+\n+\n+def delete_account(\n+    client: StorageManagementClient, resource_group: str, account_name: str\n+) -\u003e None:\n+    client.storage_accounts.delete(resource_group, account_name)\n+    logger.info(\"Deleted Storage Account '%s'.\", account_name)\n+\n+\n+def manage_storage_account(args: argparse.Namespace) -\u003e int:\n+    validate_args(args)\n+    credential = DefaultAzureCredential()\n+    client = StorageManagementClient(credential, args.subscription_id)\n+    account_created = False\n+\n+    try:\n+        availability = client.storage_accounts.check_name_availability(\n+            {\"name\": args.account_name, \"type\": \"Microsoft.Storage/storageAccounts\"}\n+        )\n+        if not availability.name_available:\n+            raise ValueError(\n+                f\"Storage Account name '{args.account_name}' is unavailable: \"\n+                f\"{availability.reason or 'unknown reason'}\"\n+            )\n+\n+        logger.info(\n+            \"Creating Storage Account '%s' in %s...\", args.account_name, LOCATION\n+        )\n+        account = client.storage_accounts.begin_create(\n+            args.resource_group,\n+            args.account_name,\n+            StorageAccountCreateParameters(\n+                sku=Sku(name=\"Standard_LRS\"),\n+                kind=Kind.STORAGE_V2,\n+                location=LOCATION,\n+            ),\n+        ).result()\n+        account_created = True\n+        logger.info(\"Created %s with resource ID %s.\", account.name, account.id)\n+\n+        logger.info(\"Storage Accounts in resource group '%s':\", args.resource_group)\n+        for listed_account in client.storage_accounts.list_by_resource_group(\n+            args.resource_group\n+        ):\n+            logger.info(\n+                \"  %s (%s, %s)\",\n+                listed_account.name,\n+                listed_account.location,\n+                listed_account.sku.name,\n+            )\n+\n+        properties = client.storage_accounts.get_properties(\n+            args.resource_group, args.account_name\n+        )\n+        logger.info(\n+            \"Properties: name=%s, location=%s, sku=%s, kind=%s, id=%s\",\n+            properties.name,\n+            properties.location,\n+            properties.sku.name,\n+            properties.kind,\n+            properties.id,\n+        )\n+\n+        blob_service = client.blob_services.get_service_properties(\n+            args.resource_group, args.account_name\n+        )\n+        if blob_service.blob_service_properties is None:\n+            blob_service.blob_service_properties = BlobServicePropertiesProperties()\n+        blob_service.blob_service_properties.is_versioning_enabled = True\n+\n+        updated_blob_service = client.blob_services.set_service_properties(\n+            args.resource_group,\n+            args.account_name,\n+            parameters=blob_service,\n+        )\n+        versioning_enabled = (\n+            updated_blob_service.blob_service_properties.is_versioning_enabled\n+        )\n+        logger.info(\"Blob versioning enabled: %s\", versioning_enabled)\n+\n+        delete_account(client, args.resource_group, args.account_name)\n+        account_created = False\n+        return 0\n+    except ClientAuthenticationError as error:\n+        logger.error(\"Azure authentication failed: %s\", error)\n+        return 2\n+    except HttpResponseError as error:\n+        status_code = error.status_code or \"unknown\"\n+        logger.error(\"Azure request failed (HTTP %s): %s\", status_code, error.message)\n+        return 3\n+    except AzureError as error:\n+        logger.error(\"Azure SDK operation failed: %s\", error)\n+        return 4\n+    except ValueError as error:\n+        logger.error(\"%s\", error)\n+        return 5\n+    finally:\n+        if account_created:\n+            logger.warning(\"Cleaning up the Storage Account after a failed operation.\")\n+            try:\n+                delete_account(client, args.resource_group, args.account_name)\n+            except AzureError as cleanup_error:\n+                logger.error(\n+                    \"Cleanup failed; delete '%s' manually: %s\",\n+                    args.account_name,\n+                    cleanup_error,\n+                )\n+        credential.close()\n+\n+\n+def main() -\u003e int:\n+    try:\n+        args = parse_args()\n+        return manage_storage_account(args)\n+    except ValueError as error:\n+        logger.error(\"%s\", error)\n+        return 5\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-identity==1.25.3\n+azure-mgmt-storage==25.1.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-1193536646/storage_account_manager.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-1193536646/storage_account_manager.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-1193536646/storage_account_manager.py
@@ -1,0 +1,190 @@
+"""Create, inspect, configure, and delete an Azure Storage Account."""
+
+from __future__ import annotations
+
+import argparse
+import logging
+import os
+import re
+
+from azure.core.exceptions import AzureError, ClientAuthenticationError, HttpResponseError
+from azure.identity import DefaultAzureCredential
+from azure.mgmt.storage import StorageManagementClient
+from azure.mgmt.storage.models import (
+    BlobServicePropertiesProperties,
+    Kind,
+    Sku,
+    StorageAccountCreateParameters,
+)
+
+LOCATION = "eastus"
+STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")
+
+logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
+logger = logging.getLogger(__name__)
+
+
+def parse_args() -> argparse.Namespace:
+    parser = argparse.ArgumentParser(
+        description="Manage the lifecycle of an Azure Storage Account."
+    )
+    parser.add_argument(
+        "--subscription-id",
+        default=os.getenv("AZURE_SUBSCRIPTION_ID"),
+        help="Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).",
+    )
+    parser.add_argument(
+        "--resource-group",
+        default=os.getenv("AZURE_RESOURCE_GROUP"),
+        help="Existing resource group (defaults to AZURE_RESOURCE_GROUP).",
+    )
+    parser.add_argument(
+        "--account-name",
+        default=os.getenv("AZURE_STORAGE_ACCOUNT_NAME"),
+        help=(
+            "Globally un
... (truncated)
```

## Generated Files

- `requirements.txt`
- `storage_account_manager.py`

## Copilot Response

I’m grounding the management-plane calls and authentication pattern, then I’ll add the script and dependency file without executing any Azure operation.

The skill references cover data-plane SDKs, so I’m checking the current `azure-mgmt-storage` method signatures before writing the management-plane implementation.

I’ve confirmed the latest SDK shape: blob versioning is updated through the account’s Blob service properties, while preserving existing settings. I’m adding a failure-safe cleanup path so a partially completed run does not leave the newly created account behind.

## Reviewer Notes (LLM-as-Judge)

### Summary

The Azure SDK exception-handling requirement is fully met.

### Strengths

- Uses specific Azure exception types before the broader AzureError handler.
- Returns distinct nonzero exit codes for different failure categories.
- Handles and reports Azure SDK errors during failure-safe cleanup.

## Grader Results

- storage-account-mgmt.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - Includes the required Azure management and identity SDK packages: Pass
      - Creates a management client authenticated with credential and subscription ID: Pass
      - Creates a storage account as a long-running operation with the correct SKU and kind: Pass
      - Lists all storage accounts in the resource group: Pass
      - Retrieves detailed properties of a specific storage account: Pass
      - Enables blob versioning on the account (via account update or blob service properties): Pass
      - Deletes the storage account: Pass
      - Code builds and runs without import errors or API misuse: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Fail
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Pass
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 6.0000** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-mp-python-account-mgmt --config python-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
