# Evaluation Report: storage-mp-python-account-mgmt

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 748.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-mp-python-account-mgmt` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 14/15 |
| Duration | 748.5s |
| Timestamp | 2026-08-28T04:41:42Z |
| Files Generated | 2 |
| Event Count | 5306 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 605.6s |
| Review | 139.5s |
| **Total** | **748.5s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage, azure-identity-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=689659 out=6894 |
| Turn Count | 13 |

## Error

```
evaluation failed: sending prompt: waiting for session.idle: context deadline exceeded
```

**Details:**

```
sending prompt: waiting for session.idle: context deadline exceeded
```

## Tool Calls

### ✅ 1. skill — 1414ms

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

### ✅ 2. skill — 1089ms

**Input:**

```json
{"skill":"azure-identity-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for Python

Authentication library for Azure SDK clients using Microsoft Entra ID.

Use this skill when:
- An app needs to authenticate to Azure services from Python
- You need `DefaultAzureCredential` for local dev + Azure deployment
- You need `ManagedIdentityCredential` for Azure-hosted workloads
- You need service principal auth with secret or certificate
- You need direct token acquisition with `get_token()`
- You need to troubleshoot credential chain failures

## Installation

```bash
pip install azure-identity
```

For VS Code or broker-based desktop auth:

```bash
pip install azure-identity-broker
```

## Python Version

`azure-identity` supports Python 3.9+.

## Environment Variables

```bash
# Service principal with client secret
AZURE_TENANT_ID=<your-tenant-id>
AZURE_CLIENT_ID=<your-client-id>
AZURE_CLIENT_SECRET=<your-client-secret>

# Service principal with certificate
AZURE_TENANT_ID=<your-tenant-id>
AZURE_CLIENT_ID=<your-client-id>
AZURE_CLIENT_CERTIFICATE_PATH=/path/to/cert.pem
AZURE_CLIENT_CERTIFICATE_PASSWORD=<optional-password>

# Authority (sovereign clouds)
AZURE_AUTHORITY_HOST=login.microsoftonline.com  # Default; or login.chinacloudapi.cn, login.microsoftonline.us

# User-assigned managed identity
AZURE_CLIENT_ID=<managed-identity-client-id>

# Credential selection (new)
AZURE_TOKEN_CREDENTIALS=dev|prod|<credential-name>  # Optional, restricts DAC chain
```

## Authentication & Lifecycle

> **🔑 Two rules apply to every code sample below:**
>
> 1. **Prefer `DefaultAzureCredential`.** It works locally (Azure CLI / VS Code / Developer CLI) and in Azure (managed identity, workload identity) with no code change. Avoid connection strings, account/API keys — they bypass Entra audit and rotation.
>    - Local dev: `DefaultAzureCredential` works as-is.
>    - Production: set `AZURE_TOKEN_CREDENTIALS=prod` (or `AZURE_TOKEN_CREDENTIALS=<specific_credential>`) to constrain the credential chain t
... (truncated)
```

### ❌ 3. azure-get_azure_bestpractices (via azure) — 186765ms

**Input:**

```json
{"intent":"Generate a local-only Python example that uses DefaultAzureCredential and the Azure Storage management-plane SDK to create, list, get, update blob versioning, and delete a storage account, with robust Azure SDK error handling. Do not execute Azure operations.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 4. azure-documentation (via azure) — 186768ms

**Input:**

```json
{"intent":"Find official current Python Azure SDK documentation for StorageManagementClient storage account create/list/get/delete operations and blob_services set_service_properties to enable blob versioning.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. view — 186758ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md
@@ -1,43 +1,43 @@
 # azure-identity-py capability coverage
 
 **SDK/package**: `azure-identity`
 
 This index maps hero scenarios in `SKILL.md` and links non-hero scenarios documented in dedicated reference files.
 
 ## Hero scenarios covered in SKILL.md
 
 - `get_bearer_token_provider`
 - `Credential Types`
 - `Specific Credential Examples`
 - `Getting Tokens Directly`
 
 ## Non-hero scenarios
 
 - `Async Credentials`: Async credentials are in `azure.identity.aio`. Always close them or use `async with`:  
   See: [`non-hero-scenarios.md#async-credentials`](non-hero-scenarios.md#async-credentials)
 - `Sovereign Clouds`: Use `AzureAuthorityHosts` or the `AZURE_AUTHORITY_HOST` env var:  
   See: [`non-hero-scenarios.md#sovereign-clouds`](non-hero-scenarios.md#sovereign-clouds)
 - `Persistent Token Caching`: Opt-in disk-based caching with `TokenCachePersistenceOptions`:  
   See: [`non-hero-scenarios.md#persistent-token-caching`](non-hero-scenarios.md#persistent-token-caching)
 - `Multi-Tenant Support`: Allow token acquisition for additional tenants beyond the configured one:  
   See: [`non-hero-scenarios.md#multi-tenant-support`](non-hero-scenarios.md#multi-tenant-support)
 - `Error Handling`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#error-handling`](non-hero-scenarios.md#error-handling)
 - `Logging`: Ena
... (truncated)
```

### ✅ 6. view — 186359ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\capabilities.md"}
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

### ✅ 7. glob — 185953ms

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py"],"pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. glob — 185453ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249","pattern":"*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. web_search (via github-mcp-server) — 182341ms

**Input:**

```json
{"query":"Official Microsoft Learn Python azure-mgmt-storage StorageManagementClient create list_by_resource_group get_properties delete BlobServices set_service_properties is_versioning_enabled example"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's a clear example using the latest Python SDK (`azure-mgmt-storage`) to demonstrate typical resource management client operations including:\n\n- Create\n- list_by_resource_group\n- get_properties\n- delete\n- BlobServices usage (set_service_properties, is_versioning_enabled)\n\nBelow is a full workflow using the `StorageManagementClient`. You will need the following prerequisites:\n- Install packages: `pip install azure-mgmt-storage azure-identity`\n- Authentication is shown using `DefaultAzureCredential` (you can use `az login` or Service Principal)\n\n```python\nimport os\nfrom azure.identity import DefaultAzureCredential\nfrom azure.mgmt.storage import StorageManagementClient\n\n# Set up credentials and client\nsubscription_id = os.getenv(\"AZURE_SUBSCRIPTION_ID\")  # Must set this in your environment!\ncredential = DefaultAzureCredential()\nstorage_client = StorageManagementClient(credential, subscription_id)\n\nresource_group = \"myResourceGroup\"\nlocation = \"eastus\"\naccount_name = \"mystorageacctuniquename1\"  # MUST BE globally unique!\n\n# 1. Create a storage account\nstorage_async_operation = storage_client.storage_accounts.begin_create(\n    resource_group,\n    account_name,\n    {\n        \"location\": location,\n        \"sku\": {\"name\": \"Standard_LRS\"},\n        \"kind\": \"StorageV2\",\n        \"enable_https_traffic_only\": True,\n    },\n)\nstorage_account = storage_async_operation.result()\n\n# 2. List storage accounts by resource group\nfor account in storage_client.storage_accounts.list_by_resource_group(resource_group):\n    print(f\"Storage Account: {account.name}\")\n\n# 3. Get storage account properties\naccount_properties = storage_client.storage_accounts.get_properties(resource_group, account_name)\nprint(f\"Primary endpoints: {account_properties.primary_endpoints}\")\n\n# 4. Delete storage account (be careful, this cannot be undone)\n# storage_client.storage_accounts.delete(resource_gro
... (truncated)
```

### ❌ 10. azure-get_azure_bestpractices (via azure) — 182357ms

**Input:**

```json
{"intent":"Python Azure SDK management-plane storage account example best practices, local code generation only; no Azure execution.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 11. azure-documentation (via azure) — 182348ms

**Input:**

```json
{"intent":"Python azure-mgmt-storage BlobServicesOperations set_service_properties is_versioning_enabled official API reference.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 12. web_fetch — 7438ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
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

### ✅ 13. web_fetch — 6562ms

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

### ✅ 14. web_fetch — 5693ms

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

### ✅ 15. web_fetch — 4860ms

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

### ✅ 16. web_fetch — 5700ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.blobservicepropertiesproperties?view=azure-python"}
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

### ✅ 17. web_fetch — 4805ms

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.models.storageaccountpropertiescreateparameters?view=azure-python"}
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

### ✅ 18. web_fetch — 4007ms

**Input:**

```json
{"max_length":12000,"start_index":30000,"url":"https://learn.microsoft.com/en-us/python/api/azure-mgmt-storage/azure.mgmt.storage.operations.storageaccountsoperations?view=azure-python"}
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

### ✅ 19. web_fetch — 1698ms

**Input:**

```json
{"max_length":3000,"url":"https://pypi.org/pypi/azure-mgmt-storage/json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-mgmt-storage/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com>","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13"],"description":"# Microsoft Azure SDK for Python\n\nThis is the Microsoft Azure Storage Management Client Library.\nThis package has been tested with Python 3.10+.\nFor a more complete view of Azure libraries, see the [azure sdk python release](https://aka.ms/azsdk/python/all).\n\n## _Disclaimer_\n\n_Azure SDK Python packages support for Python 2.7 has ended 01 January 2022. For more information and questions, please refer to https://github.com/Azure/azure-sdk-for-python/issues/20691_\n\n## Getting started\n\n### Prerequisites\n\n- Python 3.10+ is required to use this package.\n- [Azure subscription](https://azure.microsoft.com/free/)\n\n### Install the package\n\n```bash\npip install azure-mgmt-storage\npip install azure-identity\n```\n\n### Authentication\n\nBy default, [Microsoft Entra](https://learn.microsoft.com/entra/fundamentals/what-is-entra) token authentication depends on correct configuration of the following environment variables.\n\n- `AZURE_CLIENT_ID` for Azure client ID.\n- `AZURE_TENANT_ID` for Azure tenant ID.\n- `AZURE_CLIENT_SECRET` for Azure client secret.\n\nIn addition, Azure subscription ID can be configured via environment variable `AZURE_SUBSCRIPTION_ID`.\n\nWith above configuration, client can be authenticated by following code:\n\n```python\nfrom azure.identity import DefaultAzureCredential\nfrom azure.mgmt.storage import StorageManagementClient\nimpo
... (truncated)
```

### ✅ 20. web_fetch — 1704ms

**Input:**

```json
{"max_length":500,"url":"https://pypi.org/pypi/azure-identity/json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-identity/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com> License-Expression: MIT","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language ::

<note>Content truncated. Call the fetch tool with a start_index of 500 to get more content.</note>
```

### ✅ 21. apply_patch — 1268ms

**Input:**

```json
"*** Begin Patch\n*** Add File: storage_account_manager.py\n+\"\"\"Create, inspect, configure, and delete an Azure Storage account.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+import re\n+import sys\n+import uuid\n+\n+from azure.core.exceptions import (\n+    AzureError,\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+)\n+from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n+from azure.mgmt.storage import StorageManagementClient\n+from azure.mgmt.storage.models import (\n+    BlobServiceProperties,\n+    BlobServicePropertiesProperties,\n+    Sku,\n+    StorageAccountCreateParameters,\n+    StorageAccountPropertiesCreateParameters,\n+)\n+\n+LOCATION = \"eastus\"\n+STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r\"^[a-z0-9]{3,24}$\")\n+\n+\n+def required_environment_variable(name: str) -\u003e str:\n+    value = os.getenv(name)\n+    if not value:\n+        raise ValueError(f\"Required environment variable {name} is not set.\")\n+    return value\n+\n+\n+def get_storage_account_name() -\u003e str:\n+    name = os.getenv(\"AZURE_STORAGE_ACCOUNT_NAME\")\n+    if not name:\n+        name = f\"st{uuid.uuid4().hex[:22]}\"\n+\n+    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(name):\n+        raise ValueError(\n+            \"AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters \"\n+            \"and numbers.\"\n+        )\n+    return name\n+\n+\n+def describe_http_error(error: HttpResponseError) -\u003e str:\n+    status = f\"HTTP {error.status_code}\" if error.status_code else \"HTTP error\"\n+    code = getattr(getattr(error, \"error\", None), \"code\", None)\n+    return f\"{status} ({code}): {error.message}\" if code else f\"{status}: {error.message}\"\n+\n+\n+def manage_storage_account() -\u003e int:\n+    try:\n+        subscription_id = required_environment_variable(\"AZURE_SUBSCRIPTION_ID\")\n+        resource_group_name = required_environment_variable(\"AZURE_RESOURCE_GROUP\")\n+        account_name = get_storage_account_name()\n+    except ValueError as error:\n+        print(f\"Configuration error: {error}\", file=sys.stderr)\n+        return 2\n+\n+    account_created = False\n+\n+    try:\n+        with DefaultAzureCredential() as credential:\n+            client = StorageManagementClient(\n+                credential=credential,\n+                subscription_id=subscription_id,\n+            )\n+            try:\n+                print(f\"Creating storage account {account_name!r} in {LOCATION}...\")\n+                account = client.storage_accounts.begin_create(\n+                    resource_group_name=resource_group_name,\n+                    account_name=account_name,\n+                    parameters=StorageAccountCreateParameters(\n+                        sku=Sku(name=\"Standard_LRS\"),\n+                        kind=\"StorageV2\",\n+                        location=LOCATION,\n+                        properties=StorageAccountPropertiesCreateParameters(\n+                            enable_https_traffic_only=True,\n+                            minimum_tls_version=\"TLS1_2\",\n+                            allow_blob_public_access=False,\n+                            allow_cross_tenant_replication=False,\n+                        ),\n+                    ),\n+                ).result()\n+                account_created = True\n+                print(f\"Created: {account.id}\")\n+\n+                print(f\"\\nStorage accounts in resource group {resource_group_name!r}:\")\n+                accounts = client.storage_accounts.list_by_resource_group(\n+                    resource_group_name\n+                )\n+                for listed_account in accounts:\n+                    print(\n+                        f\"- {listed_account.name} \"\n+                        f\"({listed_account.location}, {listed_account.sku.name})\"\n+                    )\n+\n+                properties = client.storage_accounts.get_properties(\n+                    resource_group_name=resource_group_name,\n+                    account_name=account_name,\n+                )\n+                print(\n+                    \"\\nCreated account properties:\\n\"\n+                    f\"  name: {properties.name}\\n\"\n+                    f\"  location: {properties.location}\\n\"\n+                    f\"  sku: {properties.sku.name}\\n\"\n+                    f\"  kind: {properties.kind}\\n\"\n+                    f\"  provisioning state: {properties.provisioning_state}\"\n+                )\n+\n+                blob_properties = client.blob_services.set_service_properties(\n+                    resource_group_name=resource_group_name,\n+                    account_name=account_name,\n+                    parameters=BlobServiceProperties(\n+                        blob_service_properties=BlobServicePropertiesProperties(\n+                            is_versioning_enabled=True\n+                        )\n+                    ),\n+                )\n+                versioning_enabled = (\n+                    blob_properties.blob_service_properties.is_versioning_enabled\n+                )\n+                print(f\"\\nBlob versioning enabled: {versioning_enabled}\")\n+\n+                print(f\"\\nDeleting storage account {account_name!r}...\")\n+                client.storage_accounts.delete(\n+                    resource_group_name=resource_group_name,\n+                    account_name=account_name,\n+                )\n+                account_created = False\n+                print(\"Storage account deleted.\")\n+                return 0\n+            finally:\n+                if account_created:\n+                    print(\n+                        f\"Cleaning up storage account {account_name!r} after failure...\",\n+                        file=sys.stderr,\n+                    )\n+                    try:\n+                        client.storage_accounts.delete(\n+                            resource_group_name=resource_group_name,\n+                            account_name=account_name,\n+                        )\n+                    except HttpResponseError as cleanup_error:\n+                        print(\n+                            \"Cleanup failed: \"\n+                            f\"{describe_http_error(cleanup_error)}\",\n+                            file=sys.stderr,\n+                        )\n+                client.close()\n+    except CredentialUnavailableError as error:\n+        print(f\"No usable Azure credential was found: {error}\", file=sys.stderr)\n+    except ClientAuthenticationError as error:\n+        print(f\"Azure authentication failed: {error.message}\", file=sys.stderr)\n+    except HttpResponseError as error:\n+        print(f\"Azure request failed: {describe_http_error(error)}\", file=sys.stderr)\n+    except AzureError as error:\n+        print(f\"Azure SDK error: {error}\", file=sys.stderr)\n+\n+    return 1\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(manage_storage_account())\n*** Add File: requirements.txt\n+azure-identity\n+azure-mgmt-storage\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
@@ -1,0 +1,169 @@
+"""Create, inspect, configure, and delete an Azure Storage account."""
+
+from __future__ import annotations
+
+import os
+import re
+import sys
+import uuid
+
+from azure.core.exceptions import (
+    AzureError,
+    ClientAuthenticationError,
+    HttpResponseError,
+)
+from azure.identity import CredentialUnavailableError, DefaultAzureCredential
+from azure.mgmt.storage import StorageManagementClient
+from azure.mgmt.storage.models import (
+    BlobServiceProperties,
+    BlobServicePropertiesProperties,
+    Sku,
+    StorageAccountCreateParameters,
+    StorageAccountPropertiesCreateParameters,
+)
+
+LOCATION = "eastus"
+STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")
+
+
+def required_environment_variable(name: str) -> str:
+    value = os.getenv(name)
+    if not value:
+        raise ValueError(f"Required environment variable {name} is not set.")
+    return value
+
+
+def get_storage_account_name() -> str:
+    name = os.getenv("AZURE_STORAGE_ACCOUNT_NAME")
+    if not name:
+        name = f"st{uuid.uuid4().hex[:22]}"
+
+    if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(name):
+        raise ValueError(
+            "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters "
+            "and numbers."
+        )
+    return name
+
+
+def describe_http_error(error: HttpResponseError) -> str:
+    status = f"HTTP {error.status_code}" if error.status_code el
... (truncated)
```

### ✅ 22. powershell — 72061ms

**Input:**

```json
{"command":"python -m venv .venv \u0026\u0026 .\\.venv\\Scripts\\python.exe -m pip install --quiet -r requirements.txt \u0026\u0026 .\\.venv\\Scripts\\python.exe -m py_compile storage_account_manager.py \u0026\u0026 .\\.venv\\Scripts\\python.exe -c \"import inspect; from azure.mgmt.storage import StorageManagementClient; from azure.mgmt.storage.models import BlobServiceProperties, BlobServicePropertiesProperties, Sku, StorageAccountCreateParameters, StorageAccountPropertiesCreateParameters; create=StorageAccountCreateParameters(sku=Sku(name='Standard_LRS'), kind='StorageV2', location='eastus', properties=StorageAccountPropertiesCreateParameters(enable_https_traffic_only=True, minimum_tls_version='TLS1_2')); blob=BlobServiceProperties(blob_service_properties=BlobServicePropertiesProperties(is_versioning_enabled=True)); assert create.sku.name == 'Standard_LRS'; assert blob.blob_service_properties.is_versioning_enabled is True; print(inspect.signature(StorageManagementClient.storage_accounts.fget) if False else 'SDK model construction succeeded')\"","description":"Install dependencies and validate script","initial_wait":180,"mode":"sync"}
```

**Output:**

```
[notice] A new release of pip is available: 25.2 -> 26.2.1
[notice] To update, run: C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249\.venv\Scripts\python.exe -m pip install --upgrade pip
SDK model construction succeeded
<shellId: 0 completed with exit code 0>
```

### ✅ 23. view — 3009ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249\\storage_account_manager.py"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
@@ -1,169 +1,169 @@
 """Create, inspect, configure, and delete an Azure Storage account."""
 
 from __future__ import annotations
 
 import os
 import re
 import sys
 import uuid
 
 from azure.core.exceptions import (
     AzureError,
     ClientAuthenticationError,
     HttpResponseError,
 )
 from azure.identity import CredentialUnavailableError, DefaultAzureCredential
 from azure.mgmt.storage import StorageManagementClient
 from azure.mgmt.storage.models import (
     BlobServiceProperties,
     BlobServicePropertiesProperties,
     Sku,
     StorageAccountCreateParameters,
     StorageAccountPropertiesCreateParameters,
 )
 
 LOCATION = "eastus"
 STORAGE_ACCOUNT_NAME_PATTERN = re.compile(r"^[a-z0-9]{3,24}$")
 
 
 def required_environment_variable(name: str) -> str:
     value = os.getenv(name)
     if not value:
         raise ValueError(f"Required environment variable {name} is not set.")
     return value
 
 
 def get_storage_account_name() -> str:
     name = os.getenv("AZURE_STORAGE_ACCOUNT_NAME")
     if not name:
         name = f"st{uuid.uuid4().hex[:22]}"
 
     if not STORAGE_ACCOUNT_NAME_PATTERN.fullmatch(name):
         raise ValueError(
             "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters "
             "and numbers."
         )
     re
... (truncated)
```

### ✅ 24. powershell — 3537ms

**Input:**

```json
{"command":".\\.venv\\Scripts\\python.exe -c \"import inspect; from importlib.metadata import version; from azure.mgmt.storage.operations import BlobServicesOperations, StorageAccountsOperations; print('azure-identity', version('azure-identity')); print('azure-mgmt-storage', version('azure-mgmt-storage')); print('begin_create', inspect.signature(StorageAccountsOperations.begin_create)); print('list_by_resource_group', inspect.signature(StorageAccountsOperations.list_by_resource_group)); print('get_properties', inspect.signature(StorageAccountsOperations.get_properties)); print('set_service_properties', inspect.signature(BlobServicesOperations.set_service_properties)); print('delete', inspect.signature(StorageAccountsOperations.delete))\"","description":"Verify SDK operation signatures","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-identity 1.25.3
azure-mgmt-storage 25.1.0
begin_create (self, resource_group_name: str, account_name: str, parameters: Union[azure.mgmt.storage.models._models.StorageAccountCreateParameters, collections.abc.MutableMapping[str, Any], IO[bytes]], **kwargs: Any) -> azure.core.polling._poller.LROPoller[azure.mgmt.storage.models._models.StorageAccount]
list_by_resource_group (self, resource_group_name: str, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.StorageAccount']
get_properties (self, resource_group_name: str, account_name: str, *, expand: Union[str, azure.mgmt.storage.models._enums.StorageAccountExpand, NoneType] = None, **kwargs: Any) -> azure.mgmt.storage.models._models.StorageAccount
set_service_properties (self, resource_group_name: str, account_name: str, parameters: Union[azure.mgmt.storage.models._models.BlobServiceProperties, collections.abc.MutableMapping[str, Any], IO[bytes]], **kwargs: Any) -> azure.mgmt.storage.models._models.BlobServiceProperties
delete (self, resource_group_name: str, account_name: str, **kwargs: Any) -> None
<shellId: 1 completed with exit code 0>
```

### ✅ 25. apply_patch — 1270ms

**Input:**

```json
"*** Begin Patch\n*** Update File: storage_account_manager.py\n@@\n     BlobServicePropertiesProperties,\n     Sku,\n+    StorageAccountCheckNameAvailabilityParameters,\n     StorageAccountCreateParameters,\n     StorageAccountPropertiesCreateParameters,\n )\n@@\n def describe_http_error(error: HttpResponseError) -\u003e str:\n     status = f\"HTTP {error.status_code}\" if error.status_code else \"HTTP error\"\n     code = getattr(getattr(error, \"error\", None), \"code\", None)\n-    return f\"{status} ({code}): {error.message}\" if code else f\"{status}: {error.message}\"\n+    return (\n+        f\"{status} ({code}): {error.message}\"\n+        if code\n+        else f\"{status}: {error.message}\"\n+    )\n@@\n-    account_created = False\n+    cleanup_required = False\n \n     try:\n         with DefaultAzureCredential() as credential:\n@@\n             )\n             try:\n+                availability = client.storage_accounts.check_name_availability(\n+                    StorageAccountCheckNameAvailabilityParameters(\n+                        name=account_name,\n+                        type=\"Microsoft.Storage/storageAccounts\",\n+                    )\n+                )\n+                if not availability.name_available:\n+                    print(\n+                        f\"Storage account name {account_name!r} is unavailable: \"\n+                        f\"{availability.reason or availability.message}\",\n+                        file=sys.stderr,\n+                    )\n+                    return 2\n+\n                 print(f\"Creating storage account {account_name!r} in {LOCATION}...\")\n+                cleanup_required = True\n                 account = client.storage_accounts.begin_create(\n@@\n                     ),\n                 ).result()\n-                account_created = True\n                 print(f\"Created: {account.id}\")\n@@\n                     account_name=account_name,\n                 )\n-                account_created = False\n+                cleanup_required = False\n                 print(\"Storage account deleted.\")\n                 return 0\n             finally:\n-                if account_created:\n+                if cleanup_required:\n                     print(\n                         f\"Cleaning up storage account {account_name!r} after failure...\",\n@@\n                             f\"{describe_http_error(cleanup_error)}\",\n                             file=sys.stderr,\n                         )\n                 client.close()\n*** Update File: requirements.txt\n@@\n-azure-identity\n-azure-mgmt-storage\n+azure-identity\u003e=1.25.3,\u003c2.0.0\n+azure-mgmt-storage\u003e=25.1.0,\u003c26.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-mp-python-account-mgmt-python-azure-skills-azure-skill-mcp-microsoft-skill-3346692249/storage_account_manager.py
@@ -18,6 +18,7 @@
     BlobServiceProperties,
     BlobServicePropertiesProperties,
     Sku,
+    StorageAccountCheckNameAvailabilityParameters,
     StorageAccountCreateParameters,
     StorageAccountPropertiesCreateParameters,
 )
@@ -49,7 +50,11 @@
 def describe_http_error(error: HttpResponseError) -> str:
     status = f"HTTP {error.status_code}" if error.status_code else "HTTP error"
     code = getattr(getattr(error, "error", None), "code", None)
-    return f"{status} ({code}): {error.message}" if code else f"{status}: {error.message}"
+    return (
+        f"{status} ({code}): {error.message}"
+        if code
+        else f"{status}: {error.message}"
+    )
 
 
 def manage_storage_account() -> int:
@@ -61,7 +66,7 @@
         print(f"Configuration error: {error}", file=sys.stderr)
         return 2
 
-    account_created = False
+    cleanup_required = False
 
     try:
         with DefaultAzureCredential() as credential:
@@ -70,7 +75,22 @@
                 subscription_id=subscription_id,
             )
             try:
+                availability = client.storage_accounts.check_name_availability(
+                    StorageAccountCheckNameAvailabilityParameters(
+                        name=account_name,
+                        type="Microsoft.Storage/sto
... (truncated)
```

### 🔧 26. powershell

**Input:**

```json
{"command":".\\.venv\\Scripts\\python.exe -m pip install --quiet -r requirements.txt \u0026\u0026 .\\.venv\\Scripts\\python.exe -m py_compile storage_account_manager.py \u0026\u0026 .\\.venv\\Scripts\\python.exe -c \"from azure.mgmt.storage.models import StorageAccountCheckNameAvailabilityParameters; request=StorageAccountCheckNameAvailabilityParameters(name='stexample123', type='Microsoft.Storage/storageAccounts'); assert request.name == 'stexample123'; import storage_account_manager as app; assert app.get_storage_account_name(); print('Validation succeeded')\"","description":"Validate final script and dependencies","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `requirements.txt`
- `storage_account_manager.py`

## Copilot Response

I’m grounding the management-plane flow in the current Azure SDK patterns, then I’ll create the runnable script and package list.

The current management SDK uses the blob service properties operation—not the storage-account update operation—to enable versioning. I’m implementing that exact flow with cleanup on partial failure.

One safety edge case needs tightening: a user-supplied name could already belong to an account, and the create API can update it before the later delete. I’m adding a name-availability guard and making partial-create cleanup reliable.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. Azure SDK exceptions are handled appropriately in both the main workflow and cleanup path.

### Strengths

- Specific handling for HttpResponseError and authentication failures
- Cleanup failures are caught and reported without masking the original failure
- Azure error details are presented clearly

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
