# Evaluation Report: storage-dp-js-ts-encrypted-uploader

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 384.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-js-ts-encrypted-uploader` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 30/35 |
| Duration | 384.3s |
| Timestamp | 2026-08-29T09:45:33Z |
| Files Generated | 9 |
| Event Count | 11188 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 204.1s |
| Review | 178.6s |
| **Total** | **384.3s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** js-ts-azure-skills/azure-skill-mcp

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=369385 out=8963 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 3105ms

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

### ✅ 2. view — 14218ms

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

### ✅ 3. view — 12796ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk\\azure-storage-blob-ts.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-ts.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-ts.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-ts.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-ts.md
@@ -1,25 +1,25 @@
 # Blob Storage — TypeScript SDK Quick Reference
 
 > Condensed from **azure-storage-blob-ts**. Full patterns (SAS generation,
 > append/page blobs, streaming, browser uploads, error handling)
 > in the **azure-storage-blob-ts** plugin skill if installed.
 
 ## Install
 npm install @azure/storage-blob @azure/identity
 
 ## Quick Start
 ```typescript
 import { BlobServiceClient } from "@azure/storage-blob";
 import { DefaultAzureCredential } from "@azure/identity";
 const client = new BlobServiceClient(`https://${accountName}.blob.core.windows.net`, new DefaultAzureCredential());
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** — in production, use ManagedIdentityCredential. See [auth-best-practices.md](../auth-best-practices.md)
 - Use streaming for large files — `uploadStream`/`downloadToFile` for files > 256MB
 - Set appropriate content types — use `setHTTPHeaders` for correct MIME types
 - Use SAS tokens for client access — generate short-lived tokens for browser uploads
 - Handle errors gracefully — check `RestError.statusCode` for specific handling
 - Use `*IfNotExists` methods for idempotent container/blob creation
 - Close clients — good practice in long-running apps
 

```

### ✅ 4. view — 11399ms

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

### ✅ 5. glob — 8880ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 10210ms

**Input:**

```json
{"command":"get_ai_app_best_practices","intent":"Create a local TypeScript Node.js application using managed identity with one shared DefaultAzureCredential, Azure Blob Storage, and Azure Key Vault Keys CryptographyClient for AES-GCM envelope encryption. No live resources or deployment.","parameters":{}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 10183ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate secure TypeScript code for Azure Blob Storage and Azure Key Vault Keys envelope encryption using managed identity.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 13510ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official TypeScript JavaScript Azure Key Vault Keys CryptographyClient wrapKey unwrapKey RSA-OAEP-256 examples and result types, plus Blob Storage uploadData metadata and download metadata examples.","parameters":{"query":"JavaScript TypeScript Azure Key Vault Keys CryptographyClient wrapKey unwrapKey RSA-OAEP-256 Blob Storage uploadData metadata download metadata"}}
```

**Output:**

```
{"results":[{"title":"CryptographyClient class-Method Details","content":"# CryptographyClient class\r\n\r\n- Package:\r\n    - [@azure/keyvault-keys](https://learn.microsoft.com/en-us/javascript/api/preview-docs/@azure/keyvault-keys/)\r\n\r\nA client used to perform cryptographic operations on an Azure Key vault key or a local \u0026lt;xref:JsonWebKey\u0026gt;.\r\n\r\n## Method Details\r\n\r\n### unwrapKey(KeyWrapAlgorithm, Uint8Array, UnwrapKeyOptions)\r\n\r\nUnwraps the given wrapped key using the specified cryptography algorithm Example usage:\r\n\r\n\u0060\u0060\u0060ts\r\nlet client = new CryptographyClient(keyVaultKey, credentials);\r\nlet result = await client.unwrapKey(\u0022RSA1_5\u0022, keyToUnwrap);\r\n\u0060\u0060\u0060\r\n\r\n\u0060\u0060\u0060typescript\r\nfunction unwrapKey(algorithm: KeyWrapAlgorithm, encryptedKey: Uint8Array, options?: UnwrapKeyOptions)\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- algorithm\r\n    - [KeyWrapAlgorithm](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/keywrapalgorithm)\r\n\r\nThe decryption algorithm to use to unwrap the key.\r\n\r\n- encryptedKey\r\n    - Uint8Array\r\n\r\nThe encrypted key to unwrap.\r\n\r\n- options\r\n    - [UnwrapKeyOptions](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/unwrapkeyoptions)\r\n\r\nAdditional options.\r\n\r\n#### Returns\r\n\r\nPromise\u0026lt;[UnwrapResult](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/unwrapresult)\u0026gt;","contentUrl":"https://learn.microsoft.com/javascript/api/preview-docs/@azure/keyvault-keys/cryptographyclient?view=az-js-keyvault-v4"},{"title":"KeyVaultClient class-Method Details","content":"# KeyVaultClient class\r\n\r\n- Package:\r\n    - [azure-keyvault](https://learn.microsoft.com/en-us/javascript/api/preview-docs/azure-keyvault/)\r\n\r\n- Extends\r\n    - AzureServiceClient\r\n\r\n## Method Details\r\n\r\n### unwrapKey(string, string, string, string, Buffer, Object)\r\n\r\n\u0
... (truncated)
```

### ✅ 9. glob — 11694ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 10. rg — 6220ms

**Input:**

```json
{"-C":2,"-n":true,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787996793088-copilot-tool-output-1eeeeb13872048998737cfe475ce38fe.txt","pattern":"wrapKey\\(|unwrapKey\\(|encryptedKey|result:|metadata|uploadData"}
```

**Output:**

```
Output too large to read at once (39.4 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787996808139-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787996793088-copilot-tool-output-1eeeeb13872048998737cfe475ce38fe.txt:1:{"results":[{"title":"CryptographyClient class-Method Details","content":"# CryptographyClient class\r\n\r\n- Package:\r\n    - [@azure/keyvault-keys](https://learn.microsoft.com/en-us/javascript/api/preview-docs/@azure/keyvault-keys/)\r\n\r\nA client used to perform cryptographic operations on an Azure Key vault key or a local \u0026lt;xref:JsonWebKey\u0026gt;.\r\n\r\n## Method Details\r\
```

### ✅ 11. web_fetch — 6869ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/cryptographyclient?view=azure-node-latest"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blockblobclient?view=azure-node-latest:
---
layout: Reference
monikers:
- azure-node-latest
defaultMoniker: azure-node-latest
versioningType: Ranged
title: BlockBlobClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blockblobclient?view=azure-node-latest
config_moniker_range: azure-node-latest
uid: '@azure/storage-blob.BlockBlobClient'
package: '@azure/storage-blob'
inheritedProperties:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
uhfHeaderId: Azure
breadcrumb_path: /javascript/azure/azure_nodejs_bread/toc.json
ms.topic: generated-reference
feedback_system: None
feedback_product_url: ''
apiPlatform: javascript
author: rloutlaw
ms.manager: routlaw
ms.author: routlaw
ms.devlang: javascript
ms.date: 2018-08-09T00:00:00.0000000Z
locale: en-us
document_id: af44b145-8397-8e43-8df0-0e619dced82d
document_version_independent_id: 6226eb7d-af93-abd2-aea9-a7904edac962
updated_at: 2026-06-25T11:11:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-node/blob/live/docs-ref-autogen/@azure/storage-blob/BlockBlobClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-node/blob/3f9261f1114fe73a4a48a67ae7ea6e43d5e419ee/docs-ref-autogen/@azure/storage-blob/BlockBlobClient.yml
git_commit_id: 3f9261f1114fe73a4a48a67ae7ea6e43d5e419ee
default_moniker: azure-node-latest
site_name: Docs
depot_name: MSDN.node-api
page_type: typescript
page_kind: class
description: 'BlockBlobClient defines a set of operations applicable to block blobs. '
toc_rel: ../../azure_node_ref_toc/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
asset_id: api/@azure/storage-blob/blockblobclient
moniker_range_name: dc3f00e56d73e509b7ade314d9c7f100
monik
... (truncated)
```

### ✅ 12. web_fetch — 6551ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blockblobclient?view=azure-node-latest"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/cryptographyclient?view=azure-node-latest:
---
layout: Reference
monikers:
- azure-node-latest
defaultMoniker: azure-node-latest
versioningType: Ranged
title: CryptographyClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/cryptographyclient?view=azure-node-latest
config_moniker_range: azure-node-latest
uid: '@azure/keyvault-keys.CryptographyClient'
package: '@azure/keyvault-keys'
uhfHeaderId: Azure
breadcrumb_path: /javascript/azure/azure_nodejs_bread/toc.json
ms.topic: generated-reference
feedback_system: None
feedback_product_url: ''
apiPlatform: javascript
author: rloutlaw
ms.manager: routlaw
ms.author: routlaw
ms.devlang: javascript
ms.date: 2018-08-09T00:00:00.0000000Z
ms.service: azure-key-vault
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/f488294d-f483-456e-94e3-755f933b811b
locale: en-us
document_id: d19fa396-4bf1-14e3-1b1d-03c2f2dd1c29
document_version_independent_id: 77d13a9d-a75f-1f5f-cfe4-b1e0002a52fa
updated_at: 2025-06-11T11:13:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-node/blob/live/docs-ref-autogen/@azure/keyvault-keys/CryptographyClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-node/blob/bc0cbb1a80a941b53bc3161731f511846e5015b2/docs-ref-autogen/@azure/keyvault-keys/CryptographyClient.yml
git_commit_id: bc0cbb1a80a941b53bc3161731f511846e5015b2
default_moniker: azure-node-latest
site_name: Docs
depot_name: MSDN.node-api
page_type: typescript
page_kind: class
description: 'A client used to perform cryptographic operations on an Azure Key vault key or a local JsonWebKey. '
toc_rel: ../../azure_node_ref_toc/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
asset_id: api/@azure/keyvault-keys/cryptographyclient
moniker_range_name: dc3f00e56d73e509b7ade314d9c7f100
monikers:
- azure-node-latest
item_type: Content
sou
... (truncated)
```

### ✅ 13. powershell — 19157ms

**Input:**

```json
{"command":"$packages = '@azure/identity','@azure/keyvault-keys','@azure/storage-blob','typescript','@types/node','tsx'; foreach ($package in $packages) { $version = npm view $package version --silent; Write-Output \"$package=$version\" }","description":"Check current npm package versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
@azure/identity=4.13.2
@azure/keyvault-keys=4.10.2
@azure/storage-blob=12.33.0
typescript=7.0.2
@types/node=26.2.0
tsx=4.23.12
<shellId: 0 completed with exit code 0>
```

### ✅ 14. apply_patch — 688ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-client-side-encrypted-blob\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"description\": \"Client-side AES-GCM envelope encryption for Azure Blob Storage using an Azure Key Vault key.\",\n+  \"type\": \"module\",\n+  \"engines\": {\n+    \"node\": \"\u003e=20.0.0\"\n+  },\n+  \"scripts\": {\n+    \"build\": \"tsc -p tsconfig.json\",\n+    \"dev\": \"tsx src/main.ts\",\n+    \"start\": \"node dist/main.js\",\n+    \"typecheck\": \"tsc -p tsconfig.json --noEmit\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.13.2\",\n+    \"@azure/keyvault-keys\": \"^4.10.2\",\n+    \"@azure/storage-blob\": \"^12.33.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"tsx\": \"^4.23.12\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true,\n+    \"declaration\": true,\n+    \"sourceMap\": true\n+  },\n+  \"include\": [\n+    \"src/**/*.ts\"\n+  ]\n+}\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+.env\n+*.log\n*** Add File: .env.example\n+AZURE_STORAGE_BLOB_ENDPOINT=https://\u003cstorage-account\u003e.blob.core.windows.net\n+AZURE_STORAGE_CONTAINER=\u003cexisting-container\u003e\n+AZURE_KEY_VAULT_URL=https://\u003cvault-name\u003e.vault.azure.net\n+AZURE_KEY_NAME=\u003crsa-key-name\u003e\n+# Optional for a user-assigned managed identity:\n+# AZURE_CLIENT_ID=\u003cmanaged-identity-client-id\u003e\n+# Optional demo blob name:\n+# AZURE_BLOB_NAME=envelope-encryption-demo.txt\n*** Add File: src/errors.ts\n+export class ConfigurationError extends Error {\n+  public constructor(message: string) {\n+    super(message);\n+    this.name = \"ConfigurationError\";\n+  }\n+}\n+\n+export class KeyVaultOperationError extends Error {\n+  public constructor(\n+    operation: string,\n+    message: string,\n+    options?: ErrorOptions,\n+  ) {\n+    super(`Key Vault ${operation} failed: ${message}`, options);\n+    this.name = \"KeyVaultOperationError\";\n+  }\n+}\n+\n+export class BlobOperationError extends Error {\n+  public constructor(\n+    operation: string,\n+    message: string,\n+    options?: ErrorOptions,\n+  ) {\n+    super(`Blob ${operation} failed: ${message}`, options);\n+    this.name = \"BlobOperationError\";\n+  }\n+}\n+\n+export class EncryptionMetadataError extends Error {\n+  public constructor(message: string) {\n+    super(`Invalid encryption metadata: ${message}`);\n+    this.name = \"EncryptionMetadataError\";\n+  }\n+}\n+\n+export function getErrorMessage(error: unknown): string {\n+  return error instanceof Error ? error.message : String(error);\n+}\n+\n+export function getStatusCode(error: unknown): number | undefined {\n+  if (\n+    typeof error === \"object\" \u0026\u0026\n+    error !== null \u0026\u0026\n+    \"statusCode\" in error \u0026\u0026\n+    typeof error.statusCode === \"number\"\n+  ) {\n+    return error.statusCode;\n+  }\n+\n+  return undefined;\n+}\n*** Add File: src/config.ts\n+import { ManagedIdentityCredential } from \"@azure/identity\";\n+import { KeyClient } from \"@azure/keyvault-keys\";\n+import { BlobServiceClient, type ContainerClient } from \"@azure/storage-blob\";\n+\n+import { ConfigurationError } from \"./errors.js\";\n+\n+export interface AppConfig {\n+  blobName: string;\n+  containerClient: ContainerClient;\n+  credential: ManagedIdentityCredential;\n+  keyClient: KeyClient;\n+  keyName: string;\n+  keyVaultUrl: string;\n+}\n+\n+function requiredEnvironmentVariable(name: string): string {\n+  const value = process.env[name]?.trim();\n+  if (!value) {\n+    throw new ConfigurationError(`Environment variable ${name} is required.`);\n+  }\n+\n+  return value;\n+}\n+\n+function secureEndpoint(name: string): string {\n+  const value = requiredEnvironmentVariable(name).replace(/\\/+$/, \"\");\n+  let endpoint: URL;\n+\n+  try {\n+    endpoint = new URL(value);\n+  } catch (error) {\n+    throw new ConfigurationError(\n+      `${name} must be a valid absolute URL (${String(error)}).`,\n+    );\n+  }\n+\n+  if (endpoint.protocol !== \"https:\") {\n+    throw new ConfigurationError(`${name} must use HTTPS.`);\n+  }\n+\n+  return endpoint.toString().replace(/\\/+$/, \"\");\n+}\n+\n+export function createAppConfig(): AppConfig {\n+  const blobEndpoint = secureEndpoint(\"AZURE_STORAGE_BLOB_ENDPOINT\");\n+  const keyVaultUrl = secureEndpoint(\"AZURE_KEY_VAULT_URL\");\n+  const containerName = requiredEnvironmentVariable(\"AZURE_STORAGE_CONTAINER\");\n+  const keyName = requiredEnvironmentVariable(\"AZURE_KEY_NAME\");\n+  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();\n+\n+  const credential = managedIdentityClientId\n+    ? new ManagedIdentityCredential(managedIdentityClientId)\n+    : new ManagedIdentityCredential();\n+\n+  const blobServiceClient = new BlobServiceClient(blobEndpoint, credential);\n+  const keyClient = new KeyClient(keyVaultUrl, credential);\n+\n+  return {\n+    blobName:\n+      process.env.AZURE_BLOB_NAME?.trim() || \"envelope-encryption-demo.txt\",\n+    containerClient: blobServiceClient.getContainerClient(containerName),\n+    credential,\n+    keyClient,\n+    keyName,\n+    keyVaultUrl,\n+  };\n+}\n*** Add File: src/key-management.ts\n+import { randomBytes } from \"node:crypto\";\n+\n+import type { ManagedIdentityCredential } from \"@azure/identity\";\n+import { CryptographyClient, type KeyClient } from \"@azure/keyvault-keys\";\n+\n+import {\n+  EncryptionMetadataError,\n+  getErrorMessage,\n+  KeyVaultOperationError,\n+} from \"./errors.js\";\n+\n+const DATA_KEY_BYTES = 32;\n+export const KEY_WRAP_ALGORITHM = \"RSA-OAEP-256\" as const;\n+\n+export interface WrappedDataKey {\n+  keyId: string;\n+  plaintextKey: Buffer;\n+  wrappedKey: Buffer;\n+}\n+\n+export class KeyManagementClient {\n+  readonly #credential: ManagedIdentityCredential;\n+  readonly #keyClient: KeyClient;\n+  readonly #keyName: string;\n+  readonly #keyIdPrefix: string;\n+\n+  public constructor(\n+    keyClient: KeyClient,\n+    credential: ManagedIdentityCredential,\n+    keyVaultUrl: string,\n+    keyName: string,\n+  ) {\n+    this.#keyClient = keyClient;\n+    this.#credential = credential;\n+    this.#keyName = keyName;\n+    this.#keyIdPrefix = `${keyVaultUrl.replace(/\\/+$/, \"\")}/keys/`;\n+  }\n+\n+  public async generateAndWrapDataKey(): Promise\u003cWrappedDataKey\u003e {\n+    const plaintextKey = randomBytes(DATA_KEY_BYTES);\n+\n+    try {\n+      const vaultKey = await this.#keyClient.getKey(this.#keyName);\n+      if (!vaultKey.id) {\n+        throw new Error(`Key \"${this.#keyName}\" did not include a key ID.`);\n+      }\n+\n+      const cryptographyClient = new CryptographyClient(\n+        vaultKey.id,\n+        this.#credential,\n+      );\n+      const wrapResult = await cryptographyClient.wrapKey(\n+        KEY_WRAP_ALGORITHM,\n+        plaintextKey,\n+      );\n+\n+      return {\n+        keyId: vaultKey.id,\n+        plaintextKey,\n+        wrappedKey: Buffer.from(wrapResult.result),\n+      };\n+    } catch (error) {\n+      plaintextKey.fill(0);\n+      throw new KeyVaultOperationError(\n+        \"key wrapping\",\n+        getErrorMessage(error),\n+        { cause: error },\n+      );\n+    }\n+  }\n+\n+  public async unwrapDataKey(\n+    wrappedKey: Uint8Array,\n+    keyId: string,\n+  ): Promise\u003cBuffer\u003e {\n+    this.#assertKeyIdBelongsToConfiguredVault(keyId);\n+\n+    try {\n+      const cryptographyClient = new CryptographyClient(\n+        keyId,\n+        this.#credential,\n+      );\n+      const unwrapResult = await cryptographyClient.unwrapKey(\n+        KEY_WRAP_ALGORITHM,\n+        wrappedKey,\n+      );\n+      const plaintextKey = Buffer.from(unwrapResult.result);\n+\n+      if (plaintextKey.length !== DATA_KEY_BYTES) {\n+        plaintextKey.fill(0);\n+        throw new Error(\n+          `Unwrapped data key was ${plaintextKey.length} bytes; expected ${DATA_KEY_BYTES}.`,\n+        );\n+      }\n+\n+      return plaintextKey;\n+    } catch (error) {\n+      throw new KeyVaultOperationError(\n+        \"key unwrapping\",\n+        getErrorMessage(error),\n+        { cause: error },\n+      );\n+    }\n+  }\n+\n+  #assertKeyIdBelongsToConfiguredVault(keyId: string): void {\n+    if (!keyId.startsWith(this.#keyIdPrefix)) {\n+      throw new EncryptionMetadataError(\n+        \"the key ID does not belong to the configured Key Vault.\",\n+      );\n+    }\n+  }\n+}\n*** Add File: src/encrypted-blob-client.ts\n+import {\n+  createCipheriv,\n+  createDecipheriv,\n+  type CipherGCM,\n+  type DecipherGCM,\n+  randomBytes,\n+} from \"node:crypto\";\n+import { buffer as streamToBuffer } from \"node:stream/consumers\";\n+\n+import type { ContainerClient, Metadata } from \"@azure/storage-blob\";\n+\n+import {\n+  BlobOperationError,\n+  EncryptionMetadataError,\n+  getErrorMessage,\n+  getStatusCode,\n+} from \"./errors.js\";\n+import {\n+  KeyManagementClient,\n+  KEY_WRAP_ALGORITHM,\n+} from \"./key-management.js\";\n+\n+const CONTENT_ALGORITHM = \"AES-256-GCM\";\n+const ENCRYPTION_VERSION = \"1\";\n+const IV_BYTES = 12;\n+const AUTH_TAG_BYTES = 16;\n+\n+const METADATA = {\n+  aad: \"aad\",\n+  authenticationTag: \"authtag\",\n+  contentAlgorithm: \"contentalgorithm\",\n+  encryptionVersion: \"encryptionversion\",\n+  initializationVector: \"iv\",\n+  keyId: \"keyid\",\n+  keyWrapAlgorithm: \"keywrapalgorithm\",\n+  wrappedKey: \"wrappedkey\",\n+} as const;\n+\n+export interface EncryptedUploadResult {\n+  keyId: string;\n+  wrappedKeyBase64: string;\n+}\n+\n+interface EncryptionMetadata {\n+  authenticationTag: Buffer;\n+  initializationVector: Buffer;\n+  keyId: string;\n+  wrappedKey: Buffer;\n+}\n+\n+export class EncryptedBlobClient {\n+  readonly #containerClient: ContainerClient;\n+  readonly #keyManagementClient: KeyManagementClient;\n+\n+  public constructor(\n+    containerClient: ContainerClient,\n+    keyManagementClient: KeyManagementClient,\n+  ) {\n+    this.#containerClient = containerClient;\n+    this.#keyManagementClient = keyManagementClient;\n+  }\n+\n+  public async upload(\n+    blobName: string,\n+    plaintext: Buffer | string,\n+  ): Promise\u003cEncryptedUploadResult\u003e {\n+    const data = typeof plaintext === \"string\" ? Buffer.from(plaintext) : plaintext;\n+    const dataKey = await this.#keyManagementClient.generateAndWrapDataKey();\n+\n+    try {\n+      const initializationVector = randomBytes(IV_BYTES);\n+      const aad = this.#createAdditionalAuthenticatedData(blobName);\n+      const cipher = createCipheriv(\n+        \"aes-256-gcm\",\n+        dataKey.plaintextKey,\n+        initializationVector,\n+        { authTagLength: AUTH_TAG_BYTES },\n+      ) as CipherGCM;\n+      cipher.setAAD(aad);\n+\n+      const ciphertext = Buffer.concat([cipher.update(data), cipher.final()]);\n+      const authenticationTag = cipher.getAuthTag();\n+      const wrappedKeyBase64 = dataKey.wrappedKey.toString(\"base64\");\n+      const metadata: Metadata = {\n+        [METADATA.aad]: \"blob-path-v1\",\n+        [METADATA.authenticationTag]: authenticationTag.toString(\"base64\"),\n+        [METADATA.contentAlgorithm]: CONTENT_ALGORITHM,\n+        [METADATA.encryptionVersion]: ENCRYPTION_VERSION,\n+        [METADATA.initializationVector]:\n+          initializationVector.toString(\"base64\"),\n+        [METADATA.keyId]: dataKey.keyId,\n+        [METADATA.keyWrapAlgorithm]: KEY_WRAP_ALGORITHM,\n+        [METADATA.wrappedKey]: wrappedKeyBase64,\n+      };\n+\n+      try {\n+        await this.#containerClient\n+          .getBlockBlobClient(blobName)\n+          .uploadData(ciphertext, {\n+            blobHTTPHeaders: {\n+              blobContentType: \"application/octet-stream\",\n+            },\n+            metadata,\n+          });\n+      } catch (error) {\n+        throw new BlobOperationError(\n+          \"upload\",\n+          getErrorMessage(error),\n+          { cause: error },\n+        );\n+      }\n+\n+      return {\n+        keyId: dataKey.keyId,\n+        wrappedKeyBase64,\n+      };\n+    } finally {\n+      dataKey.plaintextKey.fill(0);\n+    }\n+  }\n+\n+  public async download(blobName: string): Promise\u003cBuffer\u003e {\n+    let ciphertext: Buffer;\n+    let metadata: Metadata | undefined;\n+\n+    try {\n+      const response = await this.#containerClient\n+        .getBlobClient(blobName)\n+        .download();\n+      metadata = response.metadata;\n+\n+      if (!response.readableStreamBody) {\n+        throw new Error(\"The Blob service returned no response body.\");\n+      }\n+\n+      ciphertext = await streamToBuffer(response.readableStreamBody);\n+    } catch (error) {\n+      const detail =\n+        getStatusCode(error) === 404\n+          ? `blob \"${blobName}\" does not exist.`\n+          : getErrorMessage(error);\n+      throw new BlobOperationError(\"download\", detail, { cause: error });\n+    }\n+\n+    const parameters = this.#parseMetadata(metadata);\n+    const dataKey = await this.#keyManagementClient.unwrapDataKey(\n+      parameters.wrappedKey,\n+      parameters.keyId,\n+    );\n+\n+    try {\n+      const decipher = createDecipheriv(\n+        \"aes-256-gcm\",\n+        dataKey,\n+        parameters.initializationVector,\n+        { authTagLength: AUTH_TAG_BYTES },\n+      ) as DecipherGCM;\n+      decipher.setAAD(this.#createAdditionalAuthenticatedData(blobName));\n+      decipher.setAuthTag(parameters.authenticationTag);\n+\n+      return Buffer.concat([\n+        decipher.update(ciphertext),\n+        decipher.final(),\n+      ]);\n+    } catch (error) {\n+      throw new BlobOperationError(\n+        \"decryption\",\n+        `authentication failed or the ciphertext is corrupt: ${getErrorMessage(error)}`,\n+        { cause: error },\n+      );\n+    } finally {\n+      dataKey.fill(0);\n+    }\n+  }\n+\n+  #createAdditionalAuthenticatedData(blobName: string): Buffer {\n+    return Buffer.from(\n+      `azure-blob:${this.#containerClient.containerName}/${blobName}:v1`,\n+      \"utf8\",\n+    );\n+  }\n+\n+  #parseMetadata(metadata: Metadata | undefined): EncryptionMetadata {\n+    if (!metadata) {\n+      throw new EncryptionMetadataError(\"metadata is missing.\");\n+    }\n+\n+    this.#expectMetadata(metadata, METADATA.encryptionVersion, ENCRYPTION_VERSION);\n+    this.#expectMetadata(metadata, METADATA.contentAlgorithm, CONTENT_ALGORITHM);\n+    this.#expectMetadata(metadata, METADATA.keyWrapAlgorithm, KEY_WRAP_ALGORITHM);\n+    this.#expectMetadata(metadata, METADATA.aad, \"blob-path-v1\");\n+\n+    const keyId = this.#requiredMetadata(metadata, METADATA.keyId);\n+    const wrappedKey = this.#decodeBase64(\n+      metadata,\n+      METADATA.wrappedKey,\n+    );\n+    const initializationVector = this.#decodeBase64(\n+      metadata,\n+      METADATA.initializationVector,\n+      IV_BYTES,\n+    );\n+    const authenticationTag = this.#decodeBase64(\n+      metadata,\n+      METADATA.authenticationTag,\n+      AUTH_TAG_BYTES,\n+    );\n+\n+    return {\n+      authenticationTag,\n+      initializationVector,\n+      keyId,\n+      wrappedKey,\n+    };\n+  }\n+\n+  #expectMetadata(\n+    metadata: Metadata,\n+    name: string,\n+    expected: string,\n+  ): void {\n+    const actual = this.#requiredMetadata(metadata, name);\n+    if (actual !== expected) {\n+      throw new EncryptionMetadataError(\n+        `${name} is \"${actual}\"; expected \"${expected}\".`,\n+      );\n+    }\n+  }\n+\n+  #requiredMetadata(metadata: Metadata, name: string): string {\n+    const value = metadata[name];\n+    if (!value) {\n+      throw new EncryptionMetadataError(`${name} is missing.`);\n+    }\n+\n+    return value;\n+  }\n+\n+  #decodeBase64(\n+    metadata: Metadata,\n+    name: string,\n+    expectedLength?: number,\n+  ): Buffer {\n+    const encoded = this.#requiredMetadata(metadata, name);\n+    if (\n+      !/^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$/.test(\n+        encoded,\n+      )\n+    ) {\n+      throw new EncryptionMetadataError(`${name} is not valid base64.`);\n+    }\n+\n+    const decoded = Buffer.from(encoded, \"base64\");\n+    if (decoded.length === 0) {\n+      throw new EncryptionMetadataError(`${name} decoded to an empty value.`);\n+    }\n+    if (expectedLength !== undefined \u0026\u0026 decoded.length !== expectedLength) {\n+      throw new EncryptionMetadataError(\n+        `${name} decoded to ${decoded.length} bytes; expected ${expectedLength}.`,\n+      );\n+    }\n+\n+    return decoded;\n+  }\n+}\n*** Add File: src/main.ts\n+import { createAppConfig } from \"./config.js\";\n+import { EncryptedBlobClient } from \"./encrypted-blob-client.js\";\n+import { getErrorMessage } from \"./errors.js\";\n+import { KeyManagementClient } from \"./key-management.js\";\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const config = createAppConfig();\n+  const keyManagementClient = new KeyManagementClient(\n+    config.keyClient,\n+    config.credential,\n+    config.keyVaultUrl,\n+    config.keyName,\n+  );\n+  const encryptedBlobClient = new EncryptedBlobClient(\n+    config.containerClient,\n+    keyManagementClient,\n+  );\n+\n+  const sample = \"Client-side envelope encryption with Azure Key Vault\";\n+  const uploadResult = await encryptedBlobClient.upload(config.blobName, sample);\n+  const decrypted = await encryptedBlobClient.download(config.blobName);\n+\n+  console.log(`Vault key ID: ${uploadResult.keyId}`);\n+  console.log(`Wrapped DEK (base64): ${uploadResult.wrappedKeyBase64}`);\n+  console.log(`Decrypted output: ${decrypted.toString(\"utf8\")}`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(`Round-trip failed: ${getErrorMessage(error)}`);\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure client-side encrypted blob uploader\n+\n+This TypeScript sample uses envelope encryption:\n+\n+1. Generate a random 256-bit data encryption key (DEK) in local memory.\n+2. Encrypt the payload locally with AES-256-GCM.\n+3. Wrap the DEK with a versioned RSA key in Azure Key Vault by using\n+   `RSA-OAEP-256`.\n+4. Store only ciphertext, the wrapped DEK, IV, authentication tag, algorithms,\n+   version, and Key Vault key ID in Blob metadata.\n+5. On download, ask Key Vault to unwrap the DEK and decrypt locally.\n+\n+The plaintext DEK is overwritten in memory after use and is never persisted.\n+The Key Vault private key never leaves Key Vault. The existing Key Vault key\n+must be an RSA key enabled for `wrapKey` and `unwrapKey` operations.\n+\n+## Configure and run\n+\n+Use an existing storage container, Key Vault, RSA key, and managed identity.\n+Do not put credentials in environment variables. Copy the endpoint and resource\n+settings from `.env.example` into the host's environment, then run:\n+\n+```powershell\n+npm install\n+npm run build\n+npm start\n+```\n+\n+The managed identity needs data-plane access equivalent to **Storage Blob Data\n+Contributor** on the container and **Key Vault Crypto User** on the key, scoped\n+as narrowly as possible. A user-assigned identity can be selected with\n+`AZURE_CLIENT_ID`; otherwise, the system-assigned identity is used.\n+\n+The demo intentionally does not create the container or key. It uploads a sample\n+string, downloads it, and prints the versioned key ID, wrapped DEK, and recovered\n+plaintext.\n+\n+## References\n+\n+- [Azure Key Vault `CryptographyClient`](https://learn.microsoft.com/javascript/api/@azure/keyvault-keys/cryptographyclient)\n+- [Azure Blob Storage JavaScript client library](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme)\n+- [Managed identity authentication](https://learn.microsoft.com/entra/identity/managed-identities-azure-resources/overview)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117/package.json
@@ -1,0 +1,27 @@
+{
+  "name": "azure-client-side-encrypted-blob",
+  "version": "1.0.0",
+  "private": true,
+  "description": "Client-side AES-GCM envelope encryption for Azure Blob Storage using an Azure Key Vault key.",
+  "type": "module",
+  "engines": {
+    "node": ">=20.0.0"
+  },
+  "scripts": {
+    "build": "tsc -p tsconfig.json",
+    "dev": "tsx src/main.ts",
+    "start": "node dist/main.js",
+    "typecheck": "tsc -p tsconfig.json --noEmit"
+  },
+  "dependencies": {
+    "@azure/identity": "^4.13.2",
+    "@azure/keyvault-keys": "^4.10.2",
+    "@azure/storage-blob": "^12.33.0"
+  },
+  "devDependencies": {
+    "@types/node": "^26.2.0",
+    "tsx": "^4.23.12",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-2812128117/tsconfig.json
@@ -1,0 +1,21 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "rootDir": "src",
+    "outDir": "dist",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "exactOptionalPropertyTypes": true,
+    "
... (truncated)
```

### ✅ 15. powershell — 28393ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund --silent \u0026\u0026 npm run build --silent","description":"Install dependencies and compile project","initial_wait":180,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\config.ts`
- `src\encrypted-blob-client.ts`
- `src\errors.ts`
- `src\key-management.ts`
- `src\main.ts`
- `tsconfig.json`

## Copilot Response

I’m grounding the implementation in the current Azure SDK authentication, Blob Storage, and Key Vault cryptography patterns before scaffolding the project.

The project will use a versioned RSA Key Vault key with `RSA-OAEP-256`, AES-256-GCM locally, strict metadata validation, and zeroization of in-memory DEKs after each operation. Service failures will be surfaced with operation-specific errors and preserved causes.

The implementation is scaffolded. I’m now resolving the declared SDK versions and compiling under strict TypeScript settings so any API or type mismatch is corrected before handoff.

Created the complete TypeScript Node.js project with:

- AES-256-GCM client-side encryption and authenticated blob-path binding
- RSA-OAEP-256 envelope encryption through Azure Key Vault **Keys**
- Versioned Key Vault key IDs and wrapped DEKs stored in blob metadata
- Shared `ManagedIdentityCredential` for Blob Storage and Key Vault
- Explicit Blob/Key Vault error handling and strict metadata validation
- Best-effort in-memory DEK zeroization
- Full upload/download round-trip demo
- `package.json`, lockfile, strict `tsconfig.json`, `.env.example`, and usage documentation

Configure the four required variables shown in `.env.example`, then run:

```powershell
npm start
```

SDK references are documented in `README.md`, including [Key Vault CryptographyClient](https://learn.microsoft.com/javascript/api/@azure/keyvault-keys/cryptographyclient) and [Azure Blob Storage for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses the correct Key Vault Keys package rather than Secrets.
- Includes complete TypeScript build configuration and Node.js type definitions.
- Provides a lockfile with resolved dependency versions.

### Issues

- The declared Node.js engine (>=20) is inconsistent with several locked Azure dependencies that require Node.js >=22.

## Grader Results

- encrypted-uploader.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (23/25)
      - Uses `@azure/keyvault-keys` (Keys, NOT Secrets) — critical distinction: Pass
      - Uses Node.js built-in `crypto` module for local AES-GCM encryption: Pass
      - Uses `KeyClient` for key management and `CryptographyClient` for wrap/unwrap operations (NOT `SecretClient`): Pass
      - Constructs `CryptographyClient` with the key ID or key name: Pass
      - Uses `CryptographyClient` for `wrapKey()` and `unwrapKey()` operations: Pass
      - Specifies RSA key wrap algorithm (e.g., `"RSA-OAEP"` or `"RSA-OAEP-256"`): Pass
      - Key material never leaves Key Vault (wrap/unwrap is server-side): Pass
      - Generates a random AES-256 DEK locally (32 bytes via `crypto.randomBytes`): Pass
      - Encrypts data with AES-GCM locally using the DEK (`crypto.createCipheriv("aes-256-gcm", ...)`): Pass
      - Wraps the DEK via Key Vault `wrapKey()`: Pass
      - Stores wrapped DEK as blob metadata (base64-encoded): Pass
      - Stores IV (initialization vector) in blob metadata (base64-encoded): Pass
      - Stores GCM auth tag in blob metadata (base64-encoded) — in Node.js the auth tag is separate from ciphertext: Pass
      - Stores vault key identifier in blob metadata: Pass
      - For decryption: retrieves wrapped DEK from metadata, unwraps via Key Vault, sets auth tag via `decipher.setAuthTag()`, decrypts locally: Pass
      - Uses AES-GCM (not AES-CBC, AES-ECB, or other modes): Pass
      - Generates random IV for each encryption (typically 12 bytes for GCM): Pass
      - Retrieves auth tag via `cipher.getAuthTag()` after encryption finalize: Pass
      - Sets auth tag via `decipher.setAuthTag()` before decryption finalize: Pass
      - Handles Key Vault errors (key disabled, key not found) via RestError: Fail
      - Handles Storage errors (blob not found) via RestError with statusCode: Fail
      - NOT using `SecretClient` instead of `KeyClient`/`CryptographyClient`: Pass
      - NOT encrypting data directly with the vault key (should be envelope encryption): Pass
      - NOT storing raw DEK in plaintext: Pass
      - NOT omitting the GCM auth tag from blob metadata (decryption will fail): Pass
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct @azure/ Scoped Packages**: All Azure SDK dependencies use @azure/ scoped packages (e.g., @azure/storage-blob, @azure/keyvault-secrets, @azure/identity). No deprecated unscoped packages (azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-*).: Pass
  - @azure/identity for Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**@azure/identity for Authentication**: Includes @azure/identity as a dependency. For local development, DefaultAzureCredential is acceptable. For production code, uses a specific Entra token credential (e.g., ManagedIdentityCredential, ClientSecretCredential, WorkloadIdentityCredential) or configures DefaultAzureCredential with the AZURE_TOKEN_CREDENTIALS environment variable set to "prod" or a specific token credential. No hardcoded connection strings, account keys, SAS tokens, or secrets in code.: Pass
  - Client Constructor with Endpoint and Credential (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Constructor with Endpoint and Credential**: SDK clients constructed with endpoint URL and credential object, e.g. new BlobServiceClient(url, credential) or new SecretClient(vaultUrl, credential). No legacy constructor patterns or connection string-based construction when identity-based auth is available.: Pass
  - Async/Await Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Fail
  - Pagination with for-await-of (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Pass
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Pass
  - RestError Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**RestError Exception Handling**: Catches RestError (from @azure/core-rest-pipeline or the service package) with statusCode inspection for error-specific handling. Not just generic Error or try/catch without type checking.: Fail
  - No Deprecated Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated Packages**: No deprecated or legacy Azure packages in package.json: azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-js, @azure/ms-rest-nodeauth, @azure/ms-rest-azure-js. All dependencies use the current @azure/* track.: Pass
  - Logging via @azure/logger (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Logging via @azure/logger**: Uses @azure/logger with setLogLevel() or the AZURE_LOG_LEVEL environment variable for SDK diagnostic logging. No custom logging wrappers around SDK internals.: Fail
  - package.json with Correct Dependencies (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**package.json with Correct Dependencies**: Includes a valid package.json with all required @azure/* dependencies listed. TypeScript projects include tsconfig.json and appropriate @types/* packages or use a recent TypeScript version with built-in types.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 92% | 1.00 | 0.9200 | 11.6% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 7.9200** | **72.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-js-ts-encrypted-uploader --config js-ts-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
