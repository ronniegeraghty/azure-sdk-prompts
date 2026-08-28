# Evaluation Report: identity-dp-dotnet-managed-identity

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 497.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-managed-identity` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 3/6 |
| Duration | 497.2s |
| Timestamp | 2026-08-27T14:56:36Z |
| Files Generated | 0 |
| Event Count | 4485 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 466.3s |
| Review | 29.6s |
| **Total** | **497.2s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=202538 out=3670 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 4084ms

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

### ✅ 2. view — 189648ms

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

### ✅ 3. view — 188083ms

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

### ❌ 4. azure-get_azure_bestpractices (via azure) — 189651ms

**Input:**

```json
{"intent":"Provide modern C# Azure SDK examples using system-assigned and user-assigned managed identities, local development fallback, pitfalls, and error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 189653ms

**Input:**

```json
{"intent":"Find current official documentation for Azure Identity .NET ManagedIdentityCredential constructors/options, system-assigned and user-assigned managed identities, DefaultAzureCredential local development, and using credentials with BlobServiceClient and SecretClient.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_search (via github-mcp-server) — 185004ms

**Input:**

```json
{"query":"Official Microsoft Learn current Azure Identity .NET ManagedIdentityCredential system-assigned user-assigned client ID resource ID constructor examples DefaultAzureCredential local development BlobServiceClient SecretClient error handling 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's an up-to-date reference and summary for using Azure Identity credentials in .NET (as of 2026), covering `ManagedIdentityCredential` (system- and user-assigned, by client ID/resource ID), `DefaultAzureCredential` (including local development), using clients like `BlobServiceClient` and `SecretClient`, and error handling. All guidance is based on current Microsoft Learn documentation and relevant best practices.\n\n---\n\n### 1. Installing Azure Identity for .NET\n\n```sh\ndotnet add package Azure.Identity\n```\nFor clients:\n```sh\ndotnet add package Azure.Storage.Blobs\ndotnet add package Azure.Security.KeyVault.Secrets\n```\n\n---\n\n### 2. DefaultAzureCredential (Best for Local Dev & Production)\n\n#### How it works\n- **Local:** Uses developer login (Azure CLI, Visual Studio, VS Code, etc.).\n- **Azure:** Uses Managed Identity automatically (system or specified user-assigned).\n\n#### Example: BlobServiceClient & SecretClient\n\n```csharp\nusing Azure.Identity;\nusing Azure.Storage.Blobs;\nusing Azure.Security.KeyVault.Secrets;\n\n// Optionally set user-assigned managed identity in Azure\nstring userAssignedClientId = \"<USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID>\";\n\nvar credential = new DefaultAzureCredential(new DefaultAzureCredentialOptions\n{\n    ManagedIdentityClientId = userAssignedClientId // omit for system-assigned\n});\n\n// Blob Service\nvar blobServiceClient = new BlobServiceClient(new Uri(\"https://<account>.blob.core.windows.net/\"), credential);\n\n// Key Vault\nvar secretClient = new SecretClient(new Uri(\"https://<vault>.vault.azure.net/\"), credential);\n```\n\n**Local Development Setup:**\nAuthenticate locally with Azure CLI or Visual Studio:\n```sh\naz login\n```\nor login in Visual Studio for single sign-on.\n\n---\n\n### 3. ManagedIdentityCredential Explicit Use\n\n#### Constructors\n\n- **System-Assigned:**\n  ```csharp\n  var cred = new ManagedIdentityCredential();\n  ```\n\n- **User-Assigned
... (truncated)
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 185010ms

**Input:**

```json
{"intent":"Get Azure SDK authentication best practices for C# managed identity examples.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 8. azure-documentation (via azure) — 185007ms

**Input:**

```json
{"intent":"Search Microsoft Learn for .NET ManagedIdentityCredential system-assigned user-assigned and local development authentication.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 9. web_fetch — 8358ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.identity.managedidentitycredential.-ctor?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/user-assigned-managed-identity:
---
layout: Conceptual
title: Authenticate Azure-hosted .NET apps to Azure resources using a user-assigned managed identity - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/user-assigned-managed-identity
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: how-to
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn how to authenticate Azure-hosted .NET apps to other Azure services using a user-assigned managed identity.
ms.custom: devx-track-dotnet, engagement-fy23, devx-track-azurecli
ms.date: 2026-03-20T00:00:00.0000000Z
locale: en-us
document_id: 3e184050-15ea-6842-89e9-165e1bfbf7a4
document_version_independent_id: 66686b09-a1e1-7c43-2e78-b20f67e52350
updated_at: 2026-03-23T20:42:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/authentication/user-assigned-managed-identity.md
gitcommit: https://github.com/dotnet/docs/blob/870c28a07336caa46b0e4c36071881831a9b077e/docs/azure/sdk/authentication/user-assigned-managed-identity.md
git_commit_id: 870c28a07336caa46b0e4c36071881831a9b077e
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 2069
asset_id: azure/sdk/authentication/user-assigned-managed-identity
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/authentication/user-assigned-managed-identity.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-1
... (truncated)
```

### ✅ 10. web_fetch — 7932ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/local-development-dev-accounts"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/local-development-dev-accounts:
---
layout: Conceptual
title: Authenticate .NET apps to Azure using developer accounts - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/local-development-dev-accounts
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: how-to
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn how to authenticate your application to Azure services when using the Azure SDK for .NET during local development using developer accounts.
ms.date: 2025-11-25T00:00:00.0000000Z
ms.custom:
- devx-track-dotnet
- engagement-fy23
- devx-track-azurecli
- sfi-image-nochange
locale: en-us
document_id: bc2d6679-46d4-a78d-b6dc-773da45c4e71
document_version_independent_id: 1e2cf4f1-fe4c-ce68-591e-8271bd0bffc0
updated_at: 2026-03-20T17:26:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/authentication/local-development-dev-accounts.md
gitcommit: https://github.com/dotnet/docs/blob/863ac42cec71c43cff89653ed8951574358361b1/docs/azure/sdk/authentication/local-development-dev-accounts.md
git_commit_id: 863ac42cec71c43cff89653ed8951574358361b1
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1870
asset_id: azure/sdk/authentication/local-development-dev-accounts
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/authentication/local-development-dev-accounts.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec
... (truncated)
```

### ✅ 11. web_fetch — 7484ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity:
---
layout: Conceptual
title: Authenticate Azure-hosted .NET apps to Azure resources using a system-assigned managed identity - .NET | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/system-assigned-managed-identity
apiPlatform: dotnet
author: alexwolfmsft
breadcrumb_path: /dotnet/breadcrumb/toc.json
feedback_system: OpenSource
feedback_product_url: https://github.com/azure/azure-sdk-for-net
ms.author: alexwolf
ms.devlang: dotnet
ms.service: dotnet-azure
ms.topic: how-to
show_latex: true
uhfHeaderId: MSDocsHeader-DotNet
description: Learn how to authenticate Azure-hosted .NET apps to other Azure services using a system-assigned managed identity.
ms.custom: devx-track-dotnet, engagement-fy23, devx-track-azurecli
ms.date: 2025-11-10T00:00:00.0000000Z
locale: en-us
document_id: aafadbb3-d54c-03c6-864e-bdeb330ce49f
document_version_independent_id: 11530b71-136a-5f66-c1d7-0cd8f196e569
updated_at: 2026-03-20T17:26:00.0000000Z
original_content_git_url: https://github.com/dotnet/docs/blob/live/docs/azure/sdk/authentication/system-assigned-managed-identity.md
gitcommit: https://github.com/dotnet/docs/blob/863ac42cec71c43cff89653ed8951574358361b1/docs/azure/sdk/authentication/system-assigned-managed-identity.md
git_commit_id: 863ac42cec71c43cff89653ed8951574358361b1
site_name: Docs
depot_name: VS.core-docs
page_type: conceptual
toc_rel: ../../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/VS.core-docs/{branchName}{pdfName}
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: csharp
word_count: 1449
asset_id: azure/sdk/authentication/system-assigned-managed-identity
moniker_range_name: 
monikers: []
item_type: Content
source_path: docs/azure/sdk/authentication/system-assigned-managed-identity.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-
... (truncated)
```

### ✅ 12. web_fetch — 6988ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/user-assigned-managed-identity"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.identity.managedidentitycredential.-ctor?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: ManagedIdentityCredential Constructor (Azure.Identity) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.identity.managedidentitycredential.-ctor?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Identity.ManagedIdentityCredential.#ctor*
namespace: Azure.Identity
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-net/issues
author: azure-sdk
ms.author: azsdkdocs
ms.topic: generated-reference
ms.devlang: csharp
ms.date: 2021-04-13T00:00:00.0000000Z
uhfHeaderId: Azure
breadcrumb_path: ../azure_breadcrumb/toc.json
apiPlatform: dotnet
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/b1cfdec6-b0c3-4209-818c-736879856e0e
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Identity.ManagedIdentityCredential..ctor
api_location:
- Azure.Identity.dll
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: ab2acb6a-a0bf-144d-bb24-08b0ad8f0cab
document_version_independent_id: 59235b41-b247-800e-a953-3fbb0d0fef74
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Identity/ManagedIdentityCredential.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Identity/ManagedIdentityCredential.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: constructor
ms.assetid: Azure.Identity.ManagedIdentityCredential.#cto
... (truncated)
```

### ✅ 13. web_fetch — 6627ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.identity.managedidentityid?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.identity.managedidentityid?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
- azure-dotnet-preview
defaultMoniker: azure-dotnet
versioningType: Ranged
title: ManagedIdentityId Class (Azure.Identity) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.identity.managedidentityid?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Identity.ManagedIdentityId
namespace: Azure.Identity
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-net/issues
author: azure-sdk
ms.author: azsdkdocs
ms.topic: generated-reference
ms.devlang: csharp
ms.date: 2021-04-13T00:00:00.0000000Z
uhfHeaderId: Azure
breadcrumb_path: ../azure_breadcrumb/toc.json
apiPlatform: dotnet
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/b1cfdec6-b0c3-4209-818c-736879856e0e
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Identity.ManagedIdentityId
api_location:
- Azure.Identity.dll
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: d9c888fd-493f-32be-db8d-6b3225caa827
document_version_independent_id: cdb04905-25d0-fd0f-3832-ba5b0c12f6f2
updated_at: 2026-07-01T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Identity/ManagedIdentityId.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/5e97917524fb860c775e5ae372cb75ea7b8690aa/xml/Azure.Identity/ManagedIdentityId.xml
git_commit_id: 5e97917524fb860c775e5ae372cb75ea7b8690aa
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: class
ms.assetid: Azure.Identity.ManagedIdentityId
description: 'Defines the configuration for a managed identity enabled on a resource. '
toc_rel: azure_re
... (truncated)
```

## Copilot Response

Managed identities let Azure SDK clients obtain Microsoft Entra tokens without storing credentials. Use `ManagedIdentityCredential` explicitly in Azure-hosted production workloads and `DefaultAzureCredential` for local development.

Install the relevant packages:

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
dotnet add package Azure.Security.KeyVault.Secrets
```

> `KeyVaultClient` belongs to the legacy Key Vault SDK. For new applications, use `SecretClient`, `KeyClient`, or `CertificateClient`.

## 1. System-assigned vs. user-assigned

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created on and tied to one Azure resource | Independent Azure resource |
| Deletion | Deleted with the hosting resource | Remains until explicitly deleted |
| Sharing | Cannot be shared | Can be attached to multiple resources |
| Selection in code | Select `SystemAssigned` | Specify client ID, object ID, or resource ID |
| Best fit | One workload with its own permissions | Shared identity, stable permissions, deployment slots, multiple identities |
| Azure limit | One per hosting resource | Multiple can be attached where supported |

In both cases:

1. Enable or attach the identity to the Azure compute resource.
2. Grant the identity an Azure RBAC role on the target service.
3. Use its credential with an Azure SDK client.

## 2. Creating each credential

Current versions of `Azure.Identity` use `ManagedIdentityId`.

### System-assigned identity

```csharp
using Azure.Core;
using Azure.Identity;

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.SystemAssigned);
```

### User-assigned identity by client ID

The identity must first be attached to the resource running the application.

```csharp
using Azure.Core;
using Azure.Identity;

string clientId = Environment.GetEnvironmentVariable(
    "APP_MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "APP_MANAGED_IDENTITY_CLIENT_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(clientId));
```

The **client ID** is not the identity’s object/principal ID. If you want to identify it another way, recent SDK versions also provide:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;

var byObjectId = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedObjectId("<object-id>"));

var byResourceId = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedResourceId(
        new ResourceIdentifier(
            "/subscriptions/<subscription-id>" +
            "/resourceGroups/<resource-group>" +
            "/providers/Microsoft.ManagedIdentity" +
            "/userAssignedIdentities/<identity-name>")));
```

Prefer the client ID unless you have a reason to use another identifier.

## 3. Using managed identity with Azure SDK clients

### System-assigned identity with Blob Storage

The identity needs an appropriate **data-plane** role such as `Storage Blob Data Reader` or `Storage Blob Data Contributor`.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.SystemAssigned);

var serviceClient = new BlobServiceClient(
    new Uri("https://mystorageaccount.blob.core.windows.net"),
    credential);

BlobContainerClient container =
    serviceClient.GetBlobContainerClient("documents");

await foreach (BlobItem blob in container.GetBlobsAsync())
{
    Console.WriteLine(blob.Name);
}
```

### User-assigned identity with Key Vault

For an RBAC-enabled vault, grant a role such as `Key Vault Secrets User` at the narrowest practical scope.

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

string clientId = Environment.GetEnvironmentVariable(
    "APP_MANAGED_IDENTITY_CLIENT_ID")
    ?? throw new InvalidOperationException(
        "APP_MANAGED_IDENTITY_CLIENT_ID is not configured.");

TokenCredential credential = new ManagedIdentityCredential(
    ManagedIdentityId.FromUserAssignedClientId(clientId));

var secretClient = new SecretClient(
    new Uri("https://my-vault.vault.azure.net"),
    credential);

KeyVaultSecret secret =
    await secretClient.GetSecretAsync("DatabaseConnectionString");

Console.WriteLine($"Retrieved secret version {secret.Properties.Version}");
```

Azure SDK clients and credentials are thread-safe, so create and reuse them rather than constructing one per request. With dependency injection:

```csharp
builder.Services.AddSingleton<TokenCredential>(
    new ManagedIdentityCredential(ManagedIdentityId.SystemAssigned));

builder.Services.AddSingleton(provider =>
{
    var credential = provider.GetRequiredService<TokenCredential>();

    return new BlobServiceClient(
        new Uri(builder.Configuration["Storage:ServiceUri"]
            ?? throw new InvalidOperationException(
                "Storage:ServiceUri is missing.")),
        credential);
});
```

## 4. Testing locally

A managed identity endpoint normally exists only inside a supported Azure hosting environment. Your local computer cannot impersonate that managed identity directly.

Authenticate as your developer identity using Azure CLI, Visual Studio, VS Code, Azure PowerShell, or Azure Developer CLI, then use `DefaultAzureCredential` locally:

```csharp
using Azure.Core;
using Azure.Identity;

bool isDevelopment =
    string.Equals(
        Environment.GetEnvironmentVariable("DOTNET_ENVIRONMENT"),
        "Development",
        StringComparison.OrdinalIgnoreCase);

TokenCredential credential = isDevelopment
    ? new DefaultAzureCredential(new DefaultAzureCredentialOptions
    {
        // Avoid probing for an Azure managed-identity endpoint locally.
        ExcludeManagedIdentityCredential = true
    })
    : new ManagedIdentityCredential(ManagedIdentityId.SystemAssigned);
```

For a user-assigned production identity:

```csharp
using Azure.Core;
using Azure.Identity;

bool isDevelopment =
    string.Equals(
        Environment.GetEnvironmentVariable("DOTNET_ENVIRONMENT"),
        "Development",
        StringComparison.OrdinalIgnoreCase);

TokenCredential credential;

if (isDevelopment)
{
    credential = new DefaultAzureCredential(
        new DefaultAzureCredentialOptions
        {
            ExcludeManagedIdentityCredential = true
        });
}
else
{
    string clientId = Environment.GetEnvironmentVariable(
        "APP_MANAGED_IDENTITY_CLIENT_ID")
        ?? throw new InvalidOperationException(
            "APP_MANAGED_IDENTITY_CLIENT_ID is required in Azure.");

    credential = new ManagedIdentityCredential(
        ManagedIdentityId.FromUserAssignedClientId(clientId));
}
```

For Azure CLI-backed local authentication:

```powershell
az login
az account set --subscription "<subscription-id-or-name>"
```

Your developer account must have equivalent data-plane permissions. Giving permissions to the managed identity does not automatically give them to your developer account.

For unit tests, avoid requesting real tokens. Wrap Azure operations behind an application interface and mock that interface, or use Azure SDK client mocking facilities. For Blob integration tests that do not need Microsoft Entra authentication, use Azurite locally.

## 5. Error handling and common pitfalls

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;

try
{
    Response<KeyVaultSecret> response =
        await secretClient.GetSecretAsync("DatabaseConnectionString");

    KeyVaultSecret secret = response.Value;
}
catch (CredentialUnavailableException ex)
{
    // No managed-identity endpoint, or the selected identity isn't available.
    Console.Error.WriteLine(
        $"Managed identity is unavailable: {ex.Message}");
    throw;
}
catch (AuthenticationFailedException ex)
{
    // Azure Identity found a credential source but couldn't obtain a token.
    Console.Error.WriteLine(
        $"Microsoft Entra authentication failed: {ex.Message}");
    throw;
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // Authentication succeeded, but the identity lacks authorization.
    Console.Error.WriteLine(
        $"Access denied. Azure error code: {ex.ErrorCode}; " +
        $"request ID: {ex.ClientRequestId}");
    throw;
}
catch (RequestFailedException ex)
{
    // Service errors such as 404, throttling, or an invalid endpoint.
    Console.Error.WriteLine(
        $"Azure request failed ({ex.Status}/{ex.ErrorCode}); " +
        $"request ID: {ex.ClientRequestId}");
    throw;
}
```

Common problems include:

- **Identity enabled but not authorized:** Enabling an identity and assigning RBAC permissions are separate operations.
- **Wrong RBAC role:** Storage management roles such as `Contributor` do not necessarily grant blob data access. Use a `Storage Blob Data ...` role.
- **Wrong user-assigned ID:** Client ID, object/principal ID, and Azure resource ID are different values.
- **Identity not attached:** Creating a user-assigned identity is insufficient; it must also be associated with the hosting resource.
- **Ambiguous identity selection:** When several user-assigned identities are attached, select one explicitly by client ID.
- **RBAC propagation delay:** New assignments can take several minutes to become effective. Retry transient startup failures with bounded backoff rather than granting broader permissions.
- **Recreating system-assigned identities:** Disabling/re-enabling or recreating the host creates a new principal. Existing role assignments may reference the old principal.
- **Using `DefaultAzureCredential` indiscriminately in production:** An explicit `ManagedIdentityCredential` is more deterministic and easier to diagnose.
- **Creating clients per request:** Reuse credentials and service clients; token caching is handled internally.
- **Catching only authentication exceptions:** A valid token with insufficient permissions normally produces `RequestFailedException` with HTTP `403`, not `AuthenticationFailedException`.
- **Logging sensitive data:** Log status, error code, and request ID, but never tokens, secret values, or connection strings.

References: [system-assigned managed identity](https://learn.microsoft.com/dotnet/azure/sdk/authentication/system-assigned-managed-identity), [user-assigned managed identity](https://learn.microsoft.com/dotnet/azure/sdk/authentication/user-assigned-managed-identity), [local development authentication](https://learn.microsoft.com/dotnet/azure/sdk/authentication/local-development-dev-accounts), and [`ManagedIdentityId`](https://learn.microsoft.com/dotnet/api/azure.identity.managedidentityid).

## Reviewer Notes (LLM-as-Judge)

### Summary

3 of 6 criteria passed. The response is comprehensive about managed identity concepts, Azure client usage, local development, and errors, but it misses three specifically required credential-construction and fallback patterns.

### Strengths

- Clear system-assigned and user-assigned identity comparison.
- Correct user-assigned client-ID selection with ManagedIdentityId.FromUserAssignedClientId.
- Practical Blob Storage and Key Vault client integrations.
- Explicit CredentialUnavailableException, AuthenticationFailedException, and authorization error handling.

### Issues

- No parameterless ManagedIdentityCredential example for system-assigned identity.
- DefaultAzureCredential is configured to exclude managed identity rather than demonstrating managed identity in its chain.
- No ChainedTokenCredential local-fallback example.

## Grader Results

- managed-identity-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (3/6)
      - `ManagedIdentityCredential` class and constructors: Pass
      - System-assigned: no parameters needed: Fail
      - User-assigned: passing the client ID: Pass
      - Integration with `DefaultAzureCredential` (managed identity in the chain): Fail
      - `CredentialUnavailableException` when not running in Azure: Pass
      - Combining with `ChainedTokenCredential` for local fallback: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 50% | 1.00 | 0.5000 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.5000** | **50.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-managed-identity --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
