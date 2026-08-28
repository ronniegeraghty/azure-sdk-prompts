# Evaluation Report: storage-mp-python-account-mgmt

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 692.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-mp-python-account-mgmt` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 14/15 |
| Duration | 692.3s |
| Timestamp | 2026-08-28T02:42:38Z |
| Files Generated | 2 |
| Event Count | 5357 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 607.9s |
| Review | 79.3s |
| **Total** | **692.3s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=710621 out=6476 |
| Turn Count | 14 |

## Error

```
evaluation failed: sending prompt: waiting for session.idle: context deadline exceeded
```

**Details:**

```
sending prompt: waiting for session.idle: context deadline exceeded
```

## Tool Calls

### ✅ 1. skill — 681ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 184729ms

**Input:**

```json
{"intent":"Generate a Python script using Azure Identity and the Azure Storage management-plane SDK to create, list, get, update blob versioning, and delete a storage account, with proper error handling and safe local placeholders.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 184731ms

**Input:**

```json
{"intent":"Find official Python Azure SDK documentation and examples for StorageManagementClient storage account create, list_by_resource_group, get_properties, update blob versioning, delete, and DefaultAzureCredential.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 184723ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-2642617046","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-storage\references\sdk-usage.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-storage\references\auth-best-practices.md
```

### ✅ 5. glob — 184137ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md,auth-best-practices.md,sdk-usage.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. view — 182834ms

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

### ✅ 7. view — 182359ms

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

### ✅ 8. web_search (via github-mcp-server) — 183295ms

**Input:**

```json
{"query":"site:learn.microsoft.com python azure-mgmt-storage StorageManagementClient storage_accounts begin_create update blob versioning allow_blob_public_access DefaultAzureCredential official docs"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here is a summary extracted from the official Microsoft documentation, including links, for using Python's `azure-mgmt-storage` package (and `StorageManagementClient`) to work with Storage Accounts—including creating/updating blob versioning or public access—using `DefaultAzureCredential`:\n\n---\n\n### 1. **Authentication with DefaultAzureCredential**  \nUse `DefaultAzureCredential` for authentication:\n```python\nfrom azure.identity import DefaultAzureCredential\nfrom azure.mgmt.storage import StorageManagementClient\nimport os\n\nsubscription_id = os.getenv(\"AZURE_SUBSCRIPTION_ID\")  # Set this in your environment\nclient = StorageManagementClient(\n    credential=DefaultAzureCredential(),\n    subscription_id=subscription_id\n)\n```\n[Docs Example][Authenticator Example]【3:3†source】\n\n---\n\n### 2. **Create/Update Storage Account with Options**  \nYou can use `begin_create` or `update` via the `storage_accounts` property, and pass parameters to control features like blob versioning and `allow_blob_public_access`:\n\n```python\nstorage_account_params = {\n    \"location\": \"eastus\",\n    \"sku\": {\"name\": \"Standard_LRS\"},\n    \"kind\": \"StorageV2\",\n    \"allow_blob_public_access\": False,    # or True\n    \"enable_blob_versioning\": True        # enable blob versioning\n    # add other parameters as needed\n}\n\nresult = client.storage_accounts.begin_create(\n    \"resource-group-name\",\n    \"storageaccountname\",\n    storage_account_params\n)\naccount = result.result()\n```\n[Docs Example][Create Account]【3:5†source】【3:9†source】\n\n---\n\n### 3. **Update Existing Storage Account**  \nTo update an existing account:\n```python\nupdate_params = {\n    \"allow_blob_public_access\": False,\n    \"enable_blob_versioning\": True\n}\nresult = client.storage_accounts.update(\n    \"resource-group-name\",\n    \"storageaccountname\",\n    update_params\n)\nupdated_account = result.result()\n```\n
... (truncated)
```

### ❌ 9. azure-get_azure_bestpractices (via azure) — 183298ms

**Input:**

```json
{"command":"get_best_practices","intent":"Python Azure Storage management-plane SDK code generation best practices using DefaultAzureCredential.","parameters":{"language":"python","resource":"storage"}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 10. web_fetch — 5818ms

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

### ✅ 11. web_fetch — 5398ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
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

### ✅ 12. web_fetch — 4839ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobserviceproperties?view=azure-python"}
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

### ✅ 13. web_fetch — 4215ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountcreateparameters?view=azure-python"}
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

### ✅ 14. web_fetch — 6517ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobservicepropertiesproperties?view=azure-python"}
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

### ✅ 15. web_fetch — 5626ms

**Input:**

```json
{"max_length":10000,"start_index":24000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python:
f a storage account.

```python
list_account_sas(resource_group_name: str, account_name: str, parameters: _models.AccountSasParameters, *, content_type: str = 'application/json', **kwargs: Any) -> _models.ListAccountSasResponse
```

#### Parameters

| Name | Description |
| --- | --- |
| resource\_group\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the resource group. The name is case insensitive. Required. |
| account\_name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Required. |
| parameters<br><br>Required | [AccountSasParameters](azure.mgmt.storage.models.accountsasparameters) or &lt;xref:JSON&gt; or [IO](https://docs.python.org/3/library/typing.html#typing.IO)[[bytes](https://docs.python.org/3/library/stdtypes.html#bytes)] <br><br>The parameters to provide to list SAS credentials for the storage account. Is one of the following types: AccountSasParameters, JSON, IO[bytes] Required. |

#### Returns

| Type | Description |
| --- | --- |
| [ListAccountSasResponse](azure.mgmt.storage.models.listaccountsasresponse) | ListAccountSasResponse. The ListAccountSasResponse is compatible with MutableMapping |

#### Exceptions

| Type | Description |
| --- | --- |
| [HttpResponseError](../azure-core/azure.core.exceptions.httpresponseerror) |  |

### list\_by\_resource\_group

Lists all the storage accounts available under the given resource group. Note that storage keys are not returned; use the ListKeys operation for this.

```python
list_by_resource_group(resource_group_name: str, **kwargs: Any) -> ItemPaged[StorageAccoun
... (truncated)
```

### ✅ 16. github-mcp-server-search_code (via github-mcp-server) — 8491ms

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":10,"query":"BlobServicePropertiesProperties is_versioning_enabled language:python org:Azure-Samples"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"storage_testcase.py","path":"eng/tools/azure-sdk-tools/devtools_testutils/storage_testcase.py","repository":"Azure/azure-sdk-for-python","text_matches":[{"fragment":"        if self.is_versioning_enabled is True:\n            props.is_versioning_enabled = True\n        if self.container_retention_days:","matches":[{"indices":[16,37],"text":"is_versioning_enabled"},{"indices":[65,86],"text":"is_versioning_enabled"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/eng/tools/azure-sdk-tools/devtools_testutils/storage_testcase.py?ref=eacd3701f6b72b48daedf7f3e19ae76d7097b55f","property":"content"},{"fragment":"        if not all(prop is None for prop in props.as_dict().values()):\n            self.client.blob_services.set_service_properties(resource_group_name, account_name, props)\n","matches":[{"indices":[95,131],"text":"blob_services.set_service_properties"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/eng/tools/azure-sdk-tools/devtools_testutils/storage_testcase.py?ref=eacd3701f6b72b48daedf7f3e19ae76d7097b55f","property":"content"}]},{"name":"disable_test_cli_mgmt_storage.py","path":"sdk/storage/azure-mgmt-storage/tests/disable_test_cli_mgmt_storage.py","repository":"Azure/azure-sdk-for-python","text_matches":[{"fragment":"            \"delete_retention_policy\": {\"enabled\": True, \"days\": \"300\"},\n            # \"is_versioning_enabled\": True,\n            # TODO: unsupport","matches":[{"indices":[88,109],"text":"is_versioning_enabled"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/storage/azure-mgmt-storage/tests/disable_test_cli_mgmt_storage.py?ref=eacd3701f6b72b48daedf7f3e19ae76d7097b55f","property":"content"},{"fragment":"        }\n        result = self.mgmt_client.blob_services.set_service_properties(resource_group.name, STORAGE_ACCOUNT_NAME, BODY)\n","matches"
... (truncated)
```

### ✅ 17. github-mcp-server-search_code (via github-mcp-server) — 7398ms

**Input:**

```json
{"fields":["name","path","repository","text_matches"],"perPage":10,"query":"blob_services.set_service_properties is_versioning_enabled language:python org:Azure"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 18. github-mcp-server-get_file_contents (via github-mcp-server) — 2420ms

**Input:**

```json
{"owner":"Azure","path":"eng/tools/azure-sdk-tools/devtools_testutils/storage_testcase.py","ref":"eacd3701f6b72b48daedf7f3e19ae76d7097b55f","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: 8bfb2a6efbbf4c91c878670ecc78625aa6538ec5)# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
# --------------------------------------------------------------------------
import functools
import logging
from time import sleep

# let this import fail, this package is used intentionally without the presence of azure-core
try:
    from azure.core.exceptions import ResourceExistsError
except:
    pass

try:
    # Note: these models are only available from v17.0.0 and higher, if you need them you'll also need azure-core 1.4.0 and higher
    from azure.mgmt.storage import StorageManagementClient
    from azure.mgmt.storage.models import (
        StorageAccount,
        Endpoints,
        LastAccessTimeTrackingPolicy,
        BlobServiceProperties,
        DeleteRetentionPolicy,
    )
except ImportError:
    pass

from . import AzureMgmtPreparer, ResourceGroupPreparer, FakeResource
from .exceptions import AzureTestError
from .resource_testcase import RESOURCE_GROUP_PARAM
from .sanitizers import add_general_string_sanitizer


FakeStorageAccount = FakeResource


# Storage Account Preparer and its shorthand decorator
class StorageAccountPreparer(AzureMgmtPreparer):
    def __init__(
        self,
        name_prefix="",
        sku="Standard_LRS",
        location="westus",
        kind="StorageV2",
        parameter_name="storage_account",
        resource_group_parameter_name=RESOURCE_GROUP_PARAM,
        disable_recording=True,
        playback_fake_resource=None,
        client_kwargs=None,
        random_name_enabled=False,
        use_cache=False,
    ):
        super(StorageAccountPreparer, self).__init__(
            name_prefix,
            24,
            disable_recording=disable_recording,
            playback_fake_resource=playback_fake_resource
... (truncated)
```

### ✅ 19. github-mcp-server-search_code (via github-mcp-server) — 6156ms

**Input:**

```json
{"fields":["name","path","sha"],"perPage":10,"query":"class BlobServiceProperties repo:Azure/azure-sdk-for-python path:sdk/storage/azure-mgmt-storage/azure/mgmt/storage/models"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 20. github-mcp-server-search_code (via github-mcp-server) — 5527ms

**Input:**

```json
{"fields":["name","path","sha"],"perPage":10,"query":"class BlobServicePropertiesProperties repo:Azure/azure-sdk-for-python path:sdk/storage/azure-mgmt-storage/azure/mgmt/storage/models"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 21. github-mcp-server-search_code (via github-mcp-server) — 4823ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"\"is_versioning_enabled\" repo:Azure/azure-sdk-for-python path:sdk/storage/azure-mgmt-storage/azure/mgmt/storage/models"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 22. web_fetch — 4093ms

**Input:**

```json
{"max_length":4000,"url":"https://pypi.org/pypi/azure-mgmt-storage/json"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-python/main/sdk/storage/azure-mgmt-storage/CHANGELOG.md:
# Release History

## 25.1.0 (2026-06-01)

### Features Added

  - Client `StorageManagementClient` added operation group `advanced_platform_metrics`
  - Model `FileShareLimits` added property `guardrail_io_scalar`
  - Model `FileShareLimits` added property `guardrail_bandwidth_scalar`
  - Added enum `AdvancedPlatformMetricsFilterType`
  - Added model `AdvancedPlatformMetricsRule`
  - Added model `AdvancedPlatformMetricsRuleConfig`
  - Added model `AdvancedPlatformMetricsRuleProperties`
  - Added enum `AdvancedPlatformMetricsRuleType`
  - Added enum `MetricsEmitted`
  - Added operation group `AdvancedPlatformMetricsOperations`

## 25.0.0 (2026-05-19)

### Features Added

  - Client `StorageManagementClient` added method `send_request`
  - Client `StorageManagementClient` added operation group `connectors`
  - Client `StorageManagementClient` added operation group `data_shares`
  - Enum `AccessTier` added member `SMART`
  - Enum `AllowedCopyScope` added member `ALL`
  - Enum `TriggerType` added member `MOCK_RUN`
  - Model `AzureEntityResource` added property `system_data`
  - Model `BlobContainer` added property `system_data`
  - Model `BlobServiceProperties` added property `system_data`
  - Model `DeletedAccount` added property `system_data`
  - Model `EncryptionScope` added property `system_data`
  - Model `FileServiceProperties` added property `system_data`
  - Model `FileServiceUsage` added property `system_data`
  - Model `FileShare` added property `system_data`
  - Model `FileShareItem` added property `system_data`
  - Model `ImmutabilityPolicy` added property `system_data`
  - Model `ListContainerItem` added property `system_data`
  - Model `ListQueue` added property `system_data`
  - Model `ManagementPolicy` added property `system_data`
  - M
... (truncated)
```

### ✅ 23. web_fetch — 3520ms

**Input:**

```json
{"max_length":3000,"url":"https://pypi.org/pypi/azure-identity/json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-identity/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com> License-Expression: MIT","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.14","Programming Language :: Python :: 3.9"],"description":"# Azure Identity client library for Python\n\nThe Azure Identity library provides [Microsoft Entra ID](https://learn.microsoft.com/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [`TokenCredential`][token_cred_ref]/[`SupportsTokenInfo`][supports_token_info_ref] implementations, which can be used to construct Azure SDK clients that support Microsoft Entra token authentication.\n\n[Source code](https://github.com/Azure/azure-sdk-for-python/blob/main/sdk/identity/azure-identity)\n| [Package (PyPI)](https://pypi.org/project/azure-identity/)\n| [Package (Conda)](https://anaconda.org/microsoft/azure-identity/)\n| [API reference documentation][ref_docs]\n| [Microsoft Entra ID documentation](https://learn.microsoft.com/entra/identity/)\n\n## Getting started\n\n### Install the package\n\nInstall Azure Identity with pip:\n\n```sh\npip install azure-identity\n```\n\n### Prerequisites\n\n- An [Azure subscription](https://azure.microsoft.com/free/python)\n- Python 3.9 or a recent version of Python 3 (this library doesn't support end-of-life versions)\n\n### Authenticate the client\n\nWhen debugging and executing code locally, it's typical for a developer to use their own account for authenticating calls to Azure service
... (truncated)
```

### ✅ 24. web_fetch — 3001ms

**Input:**

```json
{"max_length":8000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-python/main/sdk/storage/azure-mgmt-storage/CHANGELOG.md"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-mgmt-storage/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com>","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13"],"description":"# Microsoft Azure SDK for Python\n\nThis is the Microsoft Azure Storage Management Client Library.\nThis package has been tested with Python 3.10+.\nFor a more complete view of Azure libraries, see the [azure sdk python release](https://aka.ms/azsdk/python/all).\n\n## _Disclaimer_\n\n_Azure SDK Python packages support for Python 2.7 has ended 01 January 2022. For more information and questions, please refer to https://github.com/Azure/azure-sdk-for-python/issues/20691_\n\n## Getting started\n\n### Prerequisites\n\n- Python 3.10+ is required to use this package.\n- [Azure subscription](https://azure.microsoft.com/free/)\n\n### Install the package\n\n```bash\npip install azure-mgmt-storage\npip install azure-identity\n```\n\n### Authentication\n\nBy default, [Microsoft Entra](https://learn.microsoft.com/entra/fundamentals/what-is-entra) token authentication depends on correct configuration of the following environment variables.\n\n- `AZURE_CLIENT_ID` for Azure client ID.\n- `AZURE_TENANT_ID` for Azure tenant ID.\n- `AZURE_CLIENT_SECRET` for Azure client secret.\n\nIn addition, Azure subscription ID can be configured via environment variable `AZURE_SUBSCRIPTION_ID`.\n\nWith above configuration, client can be authenticated by following code:\n\n```python\nfrom azure.identity import DefaultAzureCredential\nfrom azure.mgmt.storage import StorageManagementClient\nimpo
... (truncated)
```

### ✅ 25. web_fetch — 4833ms

**Input:**

```json
{"max_length":9000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountpropertiescreateparameters?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.sku?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.Sku class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.sku?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.Sku
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
document_id: ea4bf5f2-35ab-6461-5fa7-81ac67bf02e3
document_version_independent_id: 9f30dad2-5526-1c6e-6992-f6a09127cb77
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.Sku.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.Sku.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'The SKU of the storage account. '
toc_rel: ../_splitted/azure.mgmt.storage/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-mgmt-storage/azure.mgmt.storage.models.sku
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-pytho
... (truncated)
```

### ✅ 26. web_fetch — 4350ms

**Input:**

```json
{"max_length":4000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.sku?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.kind?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.Kind enum | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.kind?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.Kind
fullName: azure.mgmt.storage.models.Kind
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
document_id: 5ef8d87f-f590-9625-f3d3-f710261acb07
document_version_independent_id: fa5afd31-0abb-7371-c4ee-f19e08917b27
updated_at: 2025-05-27T11:06:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.Kind.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/6fb08324376b86f0630bc9b2c6a4876082a923db/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.Kind.yml
git_commit_id: 6fb08324376b86f0630bc9b2c6a4876082a923db
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
page_type: python
page_kind: class
description: 'Indicates the type of storage account. '
toc_rel: ../_splitted/azure.mgmt.storage/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-mgmt-storage/azure.mgmt.storage.models.kind
moniker_range_name: db4cc4146095cc059ef03e
... (truncated)
```

### ✅ 27. web_fetch — 3897ms

**Input:**

```json
{"max_length":4000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.skuname?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.skuname?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.SkuName enum | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.skuname?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.SkuName
fullName: azure.mgmt.storage.models.SkuName
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
document_id: 2c9d52f4-b20e-d956-94b8-088c97d6c50e
document_version_independent_id: dfd4acf3-d4e5-a1b3-5f3d-9dbf64bee5e4
updated_at: 2025-05-27T11:06:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.SkuName.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/6fb08324376b86f0630bc9b2c6a4876082a923db/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.SkuName.yml
git_commit_id: 6fb08324376b86f0630bc9b2c6a4876082a923db
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
page_type: python
page_kind: class
description: 'The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called accountType. '
toc_rel: ../_splitted/azure.mgmt.storage/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: pyt
... (truncated)
```

### ✅ 28. web_fetch — 3400ms

**Input:**

```json
{"max_length":4000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.kind?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountpropertiescreateparameters?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.mgmt.storage.models.StorageAccountPropertiesCreateParameters class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountpropertiescreateparameters?view=azure-python
config_moniker_range: azure-python
uid: azure.mgmt.storage.models.StorageAccountPropertiesCreateParameters
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
document_id: 37c7a820-e2bb-d850-47c0-ef6353b4ffc0
document_version_independent_id: 7a928ad8-4736-b559-1543-b57aefd7dc90
updated_at: 2026-08-04T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.StorageAccountPropertiesCreateParameters.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/fbcd2edd0b3f4e31817ef3f083517f1c457b3591/docs-ref-autogen/azure-mgmt-storage/azure.mgmt.storage.models.StorageAccountPropertiesCreateParameters.yml
git_commit_id: fbcd2edd0b3f4e31817ef3f083517f1c457b3591
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'The parameters used to create the storage account. '
toc_rel: ../_splitted/azure.mgmt.s
... (truncated)
```

### ✅ 29. apply_patch — 1669ms

**Input:**

```json
"*** Begin Patch\n*** Add File: manage_storage_account.py\n+\"\"\"Manage an Azure Storage account through the Azure management-plane SDK.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import logging\n+import os\n+import re\n+import sys\n+from collections.abc import Sequence\n+\n+from azure.core.exceptions import AzureError, ClientAuthenticationError\n+from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n+from azure.mgmt.storage import StorageManagementClient\n+from azure.mgmt.storage.models import (\n+    BlobServiceProperties,\n+    BlobServicePropertiesProperties,\n+    Kind,\n+    MinimumTlsVersion,\n+    Sku,\n+    SkuName,\n+    StorageAccountCreateParameters,\n+    StorageAccountPropertiesCreateParameters,\n+)\n+\n+LOCATION = \"eastus\"\n+STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r\"^[a-z0-9]{3,24}$\")\n+LOGGER = logging.getLogger(\"storage-account-manager\")\n+\n+\n+def parse_args(argv: Sequence[str] | None = None) -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(\n+        description=\"Create, inspect, update, and delete an Azure Storage account.\"\n+    )\n+    parser.add_argument(\n+        \"--subscription-id\",\n+        default=os.getenv(\"AZURE_SUBSCRIPTION_ID\"),\n+        help=\"Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).\",\n+    )\n+    parser.add_argument(\n+        \"--resource-group\",\n+        default=os.getenv(\"AZURE_RESOURCE_GROUP\"),\n+        help=\"Existing resource group (defaults to AZURE_RESOURCE_GROUP).\",\n+    )\n+    parser.add_argument(\n+        \"--account-name\",\n+        default=os.getenv(\"AZURE_STORAGE_ACCOUNT_NAME\"),\n+        help=(\n+            \"Globally unique, 3-24 character lowercase alphanumeric storage \"\n+            \"account name (defaults to AZURE_STORAGE_ACCOUNT_NAME).\"\n+        ),\n+    )\n+    parser.add_argument(\n+        \"--execute\",\n+        action=\"store_true\",\n+        help=\"Perform the Azure operations. Without this flag, only print the plan.\",\n+    )\n+    return parser.parse_args(argv)\n+\n+\n+def validate_args(args: argparse.Namespace) -\u003e None:\n+    missing = [\n+        name\n+        for name, value in (\n+            (\"subscription ID\", args.subscription_id),\n+            (\"resource group\", args.resource_group),\n+            (\"storage account name\", args.account_name),\n+        )\n+        if not value\n+    ]\n+    if missing:\n+        raise ValueError(f\"Missing required value(s): {', '.join(missing)}\")\n+\n+    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(args.account_name):\n+        raise ValueError(\n+            \"Storage account name must contain 3-24 lowercase letters or digits.\"\n+        )\n+\n+\n+def print_plan(args: argparse.Namespace) -\u003e None:\n+    print(\"Dry run; no Azure requests were made.\")\n+    print(f\"1. Create {args.account_name!r} in {LOCATION} with Standard_LRS.\")\n+    print(f\"2. List storage accounts in resource group {args.resource_group!r}.\")\n+    print(f\"3. Get properties for {args.account_name!r}.\")\n+    print(\"4. Enable blob versioning through the Blob service properties.\")\n+    print(f\"5. Delete {args.account_name!r}.\")\n+    print(\"Re-run with --execute to perform these operations.\")\n+\n+\n+def describe_account(prefix: str, account: object) -\u003e None:\n+    sku = getattr(getattr(account, \"sku\", None), \"name\", None)\n+    print(\n+        f\"{prefix}: name={getattr(account, 'name', None)}, \"\n+        f\"location={getattr(account, 'location', None)}, \"\n+        f\"kind={getattr(account, 'kind', None)}, sku={sku}\"\n+    )\n+\n+\n+def create_parameters() -\u003e StorageAccountCreateParameters:\n+    return StorageAccountCreateParameters(\n+        location=LOCATION,\n+        sku=Sku(name=SkuName.STANDARD_LRS),\n+        kind=Kind.STORAGE_V2,\n+        properties=StorageAccountPropertiesCreateParameters(\n+            enable_https_traffic_only=True,\n+            minimum_tls_version=MinimumTlsVersion.TLS1_2,\n+            allow_blob_public_access=False,\n+        ),\n+    )\n+\n+\n+def enable_blob_versioning(\n+    client: StorageManagementClient,\n+    resource_group: str,\n+    account_name: str,\n+) -\u003e BlobServiceProperties:\n+    current = client.blob_services.get_service_properties(\n+        resource_group, account_name\n+    )\n+    if current.blob_service_properties is None:\n+        current.blob_service_properties = BlobServicePropertiesProperties()\n+\n+    current.blob_service_properties.is_versioning_enabled = True\n+    updated = client.blob_services.set_service_properties(\n+        resource_group, account_name, current\n+    )\n+    if (\n+        updated.blob_service_properties is None\n+        or updated.blob_service_properties.is_versioning_enabled is not True\n+    ):\n+        raise RuntimeError(\"Azure did not report blob versioning as enabled.\")\n+    return updated\n+\n+\n+def execute(args: argparse.Namespace) -\u003e int:\n+    account_created = False\n+    exit_code = 0\n+\n+    try:\n+        with DefaultAzureCredential() as credential:\n+            with StorageManagementClient(\n+                credential=credential,\n+                subscription_id=args.subscription_id,\n+            ) as client:\n+                try:\n+                    LOGGER.info(\"Creating storage account %s\", args.account_name)\n+                    created = client.storage_accounts.begin_create(\n+                        args.resource_group,\n+                        args.account_name,\n+                        create_parameters(),\n+                    ).result()\n+                    account_created = True\n+                    describe_account(\"Created\", created)\n+\n+                    print(\n+                        f\"Storage accounts in resource group \"\n+                        f\"{args.resource_group!r}:\"\n+                    )\n+                    for account in client.storage_accounts.list_by_resource_group(\n+                        args.resource_group\n+                    ):\n+                        describe_account(\" - Account\", account)\n+\n+                    properties = client.storage_accounts.get_properties(\n+                        args.resource_group, args.account_name\n+                    )\n+                    describe_account(\"Properties\", properties)\n+\n+                    enable_blob_versioning(\n+                        client, args.resource_group, args.account_name\n+                    )\n+                    print(\"Blob versioning enabled: True\")\n+                except (CredentialUnavailableError, ClientAuthenticationError) as exc:\n+                    LOGGER.error(\"Azure authentication failed: %s\", exc)\n+                    exit_code = 1\n+                except AzureError as exc:\n+                    LOGGER.error(\"Azure Storage management operation failed: %s\", exc)\n+                    exit_code = 1\n+                except RuntimeError as exc:\n+                    LOGGER.error(\"Storage account verification failed: %s\", exc)\n+                    exit_code = 1\n+                finally:\n+                    if account_created:\n+                        try:\n+                            LOGGER.info(\n+                                \"Deleting storage account %s\", args.account_name\n+                            )\n+                            client.storage_accounts.delete(\n+                                args.resource_group, args.account_name\n+                            )\n+                            print(f\"Deleted storage account {args.account_name!r}.\")\n+                        except AzureError as exc:\n+                            LOGGER.error(\n+                                \"Failed to delete storage account %s: %s\",\n+                                args.account_name,\n+                                exc,\n+                            )\n+                            exit_code = 1\n+    except (CredentialUnavailableError, ClientAuthenticationError) as exc:\n+        LOGGER.error(\"Unable to initialize Azure authentication: %s\", exc)\n+        return 1\n+\n+    return exit_code\n+\n+\n+def main(argv: Sequence[str] | None = None) -\u003e int:\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+    args = parse_args(argv)\n+\n+    try:\n+        validate_args(args)\n+    except ValueError as exc:\n+        LOGGER.error(\"%s\", exc)\n+        return 2\n+\n+    if not args.execute:\n+        print_plan(args)\n+        return 0\n+\n+    return execute(args)\n+\n+\n+if __name__ == \"__main__\":\n+    sys.exit(main())\n*** Add File: requirements.txt\n+azure-identity\u003e=1.25.0,\u003c2.0.0\n+azure-mgmt-storage==25.1.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-2642617046/manage_storage_account.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-2642617046/manage_storage_account.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-2642617046/manage_storage_account.py
@@ -1,0 +1,225 @@
+"""Manage an Azure Storage account through the Azure management-plane SDK."""
+
+from __future__ import annotations
+
+import argparse
+import logging
+import os
+import re
+import sys
+from collections.abc import Sequence
+
+from azure.core.exceptions import AzureError, ClientAuthenticationError
+from azure.identity import CredentialUnavailableError, DefaultAzureCredential
+from azure.mgmt.storage import StorageManagementClient
+from azure.mgmt.storage.models import (
+    BlobServiceProperties,
+    BlobServicePropertiesProperties,
+    Kind,
+    MinimumTlsVersion,
+    Sku,
+    SkuName,
+    StorageAccountCreateParameters,
+    StorageAccountPropertiesCreateParameters,
+)
+
+LOCATION = "eastus"
+STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")
+LOGGER = logging.getLogger("storage-account-manager")
+
+
+def parse_args(argv: Sequence[str] | None = None) -> argparse.Namespace:
+    parser = argparse.ArgumentParser(
+        description="Create, inspect, update, and delete an Azure Storage account."
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
+
... (truncated)
```

### ✅ 30. powershell — 72691ms

**Input:**

```json
{"command":"python -m venv .venv \u0026\u0026 .\\.venv\\Scripts\\python.exe -m pip install --quiet --disable-pip-version-check -r requirements.txt \u0026\u0026 .\\.venv\\Scripts\\python.exe -m py_compile manage_storage_account.py \u0026\u0026 .\\.venv\\Scripts\\python.exe manage_storage_account.py --subscription-id 00000000-0000-0000-0000-000000000000 --resource-group example-rg --account-name examplestorageacct123","description":"Install SDKs and validate the script locally","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Dry run; no Azure requests were made.
1. Create 'examplestorageacct123' in eastus with Standard_LRS.
2. List storage accounts in resource group 'example-rg'.
3. Get properties for 'examplestorageacct123'.
4. Enable blob versioning through the Blob service properties.
5. Delete 'examplestorageacct123'.
Re-run with --execute to perform these operations.
<shellId: 0 completed with exit code 0>
```

### ✅ 31. powershell — 2818ms

**Input:**

```json
{"command":"@'\nfrom manage_storage_account import create_parameters, enable_blob_versioning\nfrom azure.mgmt.storage.models import BlobServiceProperties, BlobServicePropertiesProperties\n\nparams = create_parameters()\nassert params.location == 'eastus'\nassert params.sku.name == 'Standard_LRS'\nassert params.kind == 'StorageV2'\nassert params.properties.enable_https_traffic_only is True\nassert params.properties.minimum_tls_version == 'TLS1_2'\n\nclass BlobServices:\n    def __init__(self):\n        self.saved = None\n    def get_service_properties(self, resource_group, account_name):\n        return BlobServiceProperties(blob_service_properties=BlobServicePropertiesProperties())\n    def set_service_properties(self, resource_group, account_name, properties):\n        self.saved = properties\n        return properties\n\nclass Client:\n    blob_services = BlobServices()\n\nupdated = enable_blob_versioning(Client(), 'example-rg', 'examplestorageacct123')\nassert updated.blob_service_properties.is_versioning_enabled is True\nprint('SDK model and blob-versioning flow validated.')\n'@ | .\\.venv\\Scripts\\python.exe -","description":"Validate SDK models and versioning flow","initial_wait":30,"mode":"sync"}
```

**Output:**

```
SDK model and blob-versioning flow validated.
<shellId: 1 completed with exit code 0>
```

### ✅ 32. powershell — 3105ms

**Input:**

```json
{"command":"$targets = @('.venv', '__pycache__') | Where-Object { Test-Path $_ }; foreach ($target in $targets) { Remove-Item -LiteralPath $target -Recurse -Force }","description":"Remove local validation artifacts","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 2 completed with exit code 0>
```

## Generated Files

- `manage_storage_account.py`
- `requirements.txt`

## Copilot Response

I’ll ground the management-plane operations and authentication pattern, then create a local-only script with guarded destructive execution.

The blob-versioning change belongs to `blob_services.set_service_properties`, not the storage-account `update` operation; I’m pinning the script to a verified SDK shape and making execution explicit so it remains safe offline by default.

## Reviewer Notes (LLM-as-Judge)

### Summary

The exception-handling criterion is fully met.

### Strengths

- Handles authentication failures distinctly.
- Catches Azure SDK operation errors and reports failure via exit code.
- Attempts cleanup in a finally block and handles cleanup errors.

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
  - Context Manager for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 7.0000** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-mp-python-account-mgmt --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
