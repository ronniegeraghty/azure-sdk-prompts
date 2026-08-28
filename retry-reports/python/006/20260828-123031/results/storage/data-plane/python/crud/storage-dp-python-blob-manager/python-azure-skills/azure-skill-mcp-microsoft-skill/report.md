# Evaluation Report: storage-dp-python-blob-manager

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 668.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-python-blob-manager` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 15/17 |
| Duration | 668.0s |
| Timestamp | 2026-08-28T04:30:31Z |
| Files Generated | 8 |
| Event Count | 17496 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 577.2s |
| Review | 89.9s |
| **Total** | **668.0s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage-blob-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=515423 out=13614 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 677ms

**Input:**

```json
{"skill":"azure-storage-blob-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Blob Storage SDK for Python

Client library for Azure Blob Storage — object storage for unstructured data.

## Installation

```bash
pip install azure-storage-blob azure-identity
```

## Environment Variables

```bash
AZURE_STORAGE_ACCOUNT_NAME=<your-storage-account>  # Required for all auth methods
# Or use full URL
AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net  # Alternative to account name
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication & Lifecycle

> **🔑 Two rules apply to every code sample below:**
>
> 1. **Prefer `DefaultAzureCredential`.** It works locally (Azure CLI / VS Code / Developer CLI) and in Azure (managed identity, workload identity) with no code change. Avoid connection strings, account/API keys — they bypass Entra audit and rotation.
>    - Local dev: `DefaultAzureCredential` works as-is.
>    - Production: set `AZURE_TOKEN_CREDENTIALS=prod` (or `AZURE_TOKEN_CREDENTIALS=<specific_credential>`) to constrain the credential chain to production-safe credentials.
> 2. **Wrap every client in a context manager** so HTTP transports, sockets, and token caches are released deterministically:
>    - Sync: `with <Client>(...) as client:`
>    - Async: `async with <Client>(...) as client:` **and** `async with DefaultAzureCredential() as credential:` (from `azure.identity.aio`)
>
> Snippets may abbreviate this setup, but production code should always follow both rules.

```python
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential
from azure.storage.blob import BlobServiceClient

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
credential = DefaultAzureCredential(require_envvar=True)
# Or use a specific credential directly in production:
# See https://learn.microsoft.com/python/api/overview/azure/identity-readme?view=azure-p
... (truncated)
```

### ✅ 2. view — 182399ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-storage-blob-py\\references\\capabilities.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/non-hero-scenarios.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/non-hero-scenarios.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/non-hero-scenarios.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/non-hero-scenarios.md
@@ -1,48 +1,48 @@
 # azure-storage-blob-py non-hero scenarios
 
 These scenarios are intentionally separate from hero flows in `SKILL.md`.
 They cover secondary/advanced patterns typically used after the primary end-to-end path is working.
 
 ## Blob Properties and Metadata
 
 ```python
 # Get properties
 properties = blob_client.get_blob_properties()
 print(f"Size: {properties.size}")
 print(f"Content-Type: {properties.content_settings.content_type}")
 print(f"Last modified: {properties.last_modified}")
 
 # Set metadata
 blob_client.set_blob_metadata(metadata={"category": "logs", "year": "2024"})
 
 # Set content type
 from azure.storage.blob import ContentSettings
 blob_client.set_http_headers(
     content_settings=ContentSettings(content_type="application/json")
 )
 ```
 
 ## Async Client
 
 ```python
 from azure.identity.aio import DefaultAzureCredential
 from azure.storage.blob.aio import BlobServiceClient
 
 async def upload_async():
     async with DefaultAzureCredential() as credential:
         async with BlobServiceClient(account_url, credential=credential) as client:
             blob_client = client.get_blob_client("mycontainer", "sample.txt")
             
             with open("./file.txt", "rb") as data:
                 await blob_client.upload_blob(data, overwrite=True)
 
 # Downlo
... (truncated)
```

### ✅ 3. view — 182076ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-storage-blob-py\\references\\non-hero-scenarios.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/capabilities.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/capabilities.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/capabilities.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-storage-blob-py/references/capabilities.md
@@ -1,33 +1,33 @@
 # azure-storage-blob-py capability coverage
 
 **SDK/package**: `azure-storage-blob`
 
 This index maps hero scenarios in `SKILL.md` and links non-hero scenarios documented in dedicated reference files.
 
 ## Hero scenarios covered in SKILL.md
 
 - `Client Hierarchy`
 - `Core Workflow`
 - `Performance Tuning`
 - `SAS Tokens (User Delegation)`
 
 ## Non-hero scenarios
 
 - `Blob Properties and Metadata`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#blob-properties-and-metadata`](non-hero-scenarios.md#blob-properties-and-metadata)
 - `Async Client`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#async-client`](non-hero-scenarios.md#async-client)
 
 ## Related deep-dive references
 
 - [`non-hero-scenarios.md`](non-hero-scenarios.md): Dedicated non-hero examples and implementation notes.
 
 ## API breadth checklist
 
 - Verify client/auth mode for the environment before coding.
 - Confirm operation-group/method names against current Microsoft Learn API reference.
 - For Python SDKs with both sync and async clients, document both forms without a blanket preference.
 - Include cleanup/delete paths for created resources in examples.
 - Prefer idempotent create/update operations where available.
 - Validate paging/LRO/error-handling patterns for prod
... (truncated)
```

### ✅ 4. view — 181773ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739 b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739
@@ -1,0 +1,0 @@


```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 182401ms

**Input:**

```json
{"intent":"Create a local-only Python project using DefaultAzureCredential and Azure Blob Storage SDK, with sync and async reusable services, large streaming uploads, leases/ETag concurrency, retries, logging, timeouts, and graceful errors.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 184932ms

**Input:**

```json
{"intent":"Find official Python Azure Blob Storage documentation for upload_blob ETag match_condition optimistic concurrency, blob leases, per-operation timeout, ExponentialRetry configuration, and sync/async streaming uploads and downloads.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. web_fetch — 184925ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.blobclient?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.aio.blobclient?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.storage.blob.aio.BlobClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.aio.blobclient?view=azure-python
config_moniker_range: azure-python
uid: azure.storage.blob.aio.BlobClient
module: azure.storage.blob.aio
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
document_id: 4d0efb74-cb1f-4e35-ce33-5c8769c8dff4
document_version_independent_id: 771489e0-27c4-0e90-6ec5-2b1985920fe7
updated_at: 2026-05-18T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-storage-blob/azure.storage.blob.aio.BlobClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/12f09d39a5412ab4ced175543b7b2cf9736c07fb/docs-ref-autogen/azure-storage-blob/azure.storage.blob.aio.BlobClient.yml
git_commit_id: 12f09d39a5412ab4ced175543b7b2cf9736c07fb
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'A client to interact with a specific blob, although that blob may not yet exist. '
toc_rel: ../_splitted/azure.storage.blob/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-storage-blob/azure.storage.blob.aio.blobclient
mo
... (truncated)
```

### ✅ 8. web_fetch — 184568ms

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.exponentialretry?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.blobclient?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.storage.blob.BlobClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.blobclient?view=azure-python
config_moniker_range: azure-python
uid: azure.storage.blob.BlobClient
module: azure.storage.blob
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
document_id: 2a557056-1da5-6c2d-fcee-e4e246a7a221
document_version_independent_id: 9ee6555a-aaca-243f-409e-1ac5881e3dbc
updated_at: 2026-06-12T11:10:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-storage-blob/azure.storage.blob.BlobClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/d4e56edda50d2e71ff1b2ef662123d57e188b608/docs-ref-autogen/azure-storage-blob/azure.storage.blob.BlobClient.yml
git_commit_id: d4e56edda50d2e71ff1b2ef662123d57e188b608
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'A client to interact with a specific blob, although that blob may not yet exist. For more optional configuration, please click here. '
toc_rel: ../_splitted/azure.storage.blob/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-storage-blob/azure.storag
... (truncated)
```

### ✅ 9. web_fetch — 184209ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.aio.blobclient?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.exponentialretry?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.storage.blob.ExponentialRetry class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-storage-blob/azure.storage.blob.exponentialretry?view=azure-python
config_moniker_range: azure-python
uid: azure.storage.blob.ExponentialRetry
module: azure.storage.blob
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
document_id: d9fceb71-8e8b-8722-0226-3322bcece09c
document_version_independent_id: 7901e21f-0cb8-b054-c853-2e112672ce4b
updated_at: 2025-07-17T11:06:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-storage-blob/azure.storage.blob.ExponentialRetry.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/70eb6e32c16de5c57052d4c4c8d1710cddd6ffb1/docs-ref-autogen/azure-storage-blob/azure.storage.blob.ExponentialRetry.yml
git_commit_id: 70eb6e32c16de5c57052d4c4c8d1710cddd6ffb1
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Exponential retry. Constructs an Exponential retry object. The initial_backoff is used for the first retry. Subsequent retries are retried after initial_backoff + increment_power^retry_count seconds. '
toc_rel: ../_splitted/azure.storage.blob/toc.json
feedback_help_link_type: ''

... (truncated)
```

### ✅ 10. apply_patch — 768ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-storage-blob\u003e=12.25.0,\u003c13.0.0\n*** Add File: blob_manager/__init__.py\n+\"\"\"Reusable synchronous and asynchronous Azure Blob Storage managers.\"\"\"\n+\n+from .async_service import AsyncBlobStorageManager\n+from .config import StorageSettings\n+from .models import BlobInfo, OperationResult\n+from .sync_service import BlobStorageManager\n+\n+__all__ = [\n+    \"AsyncBlobStorageManager\",\n+    \"BlobInfo\",\n+    \"BlobStorageManager\",\n+    \"OperationResult\",\n+    \"StorageSettings\",\n+]\n*** Add File: blob_manager/config.py\n+\"\"\"Environment-based Azure Blob Storage client configuration.\"\"\"\n+\n+from __future__ import annotations\n+\n+import logging\n+import os\n+from dataclasses import dataclass\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.storage.blob import BlobServiceClient, ExponentialRetry\n+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient\n+\n+_HTTP_LOGGER = \"azure.core.pipeline.policies.http_logging_policy\"\n+\n+\n+def _env_bool(name: str, default: bool) -\u003e bool:\n+    value = os.getenv(name)\n+    if value is None:\n+        return default\n+    normalized = value.strip().lower()\n+    if normalized in {\"1\", \"true\", \"yes\", \"on\"}:\n+        return True\n+    if normalized in {\"0\", \"false\", \"no\", \"off\"}:\n+        return False\n+    raise ValueError(f\"{name} must be true or false, not {value!r}.\")\n+\n+\n+def _env_int(name: str, default: int, minimum: int = 0) -\u003e int:\n+    value = os.getenv(name)\n+    parsed = default if value is None else int(value)\n+    if parsed \u003c minimum:\n+        raise ValueError(f\"{name} must be at least {minimum}.\")\n+    return parsed\n+\n+\n+@dataclass(frozen=True)\n+class StorageSettings:\n+    \"\"\"Settings used by both synchronous and asynchronous clients.\"\"\"\n+\n+    account_url: str\n+    max_retries: int = 5\n+    retry_delay: int = 2\n+    retry_increment: int = 2\n+    retry_jitter: int = 1\n+    http_logging_enabled: bool = False\n+    http_log_level: str = \"WARNING\"\n+    max_block_size: int = 4 * 1024 * 1024\n+    max_single_put_size: int = 64 * 1024 * 1024\n+    max_concurrency: int = 4\n+    connection_timeout: int = 20\n+    read_timeout: int = 120\n+\n+    @classmethod\n+    def from_env(cls) -\u003e \"StorageSettings\":\n+        account_url = os.getenv(\"AZURE_STORAGE_ACCOUNT_URL\", \"\").strip().rstrip(\"/\")\n+        if not account_url:\n+            raise ValueError(\n+                \"AZURE_STORAGE_ACCOUNT_URL is required \"\n+                \"(for example, https://\u003caccount\u003e.blob.core.windows.net).\"\n+            )\n+        if not account_url.lower().startswith(\"https://\"):\n+            raise ValueError(\"AZURE_STORAGE_ACCOUNT_URL must use HTTPS.\")\n+\n+        log_level = os.getenv(\"AZURE_STORAGE_HTTP_LOG_LEVEL\", \"WARNING\").upper()\n+        if log_level not in logging.getLevelNamesMapping():\n+            raise ValueError(\n+                f\"AZURE_STORAGE_HTTP_LOG_LEVEL is not a valid logging level: {log_level!r}.\"\n+            )\n+\n+        return cls(\n+            account_url=account_url,\n+            max_retries=_env_int(\"AZURE_STORAGE_MAX_RETRIES\", 5),\n+            retry_delay=_env_int(\"AZURE_STORAGE_RETRY_DELAY\", 2),\n+            retry_increment=_env_int(\"AZURE_STORAGE_RETRY_INCREMENT\", 2),\n+            retry_jitter=_env_int(\"AZURE_STORAGE_RETRY_JITTER\", 1),\n+            http_logging_enabled=_env_bool(\n+                \"AZURE_STORAGE_HTTP_LOGGING_ENABLED\", False\n+            ),\n+            http_log_level=log_level,\n+            max_block_size=_env_int(\n+                \"AZURE_STORAGE_MAX_BLOCK_SIZE\", 4 * 1024 * 1024, 1024\n+            ),\n+            max_single_put_size=_env_int(\n+                \"AZURE_STORAGE_MAX_SINGLE_PUT_SIZE\", 64 * 1024 * 1024, 1024\n+            ),\n+            max_concurrency=_env_int(\"AZURE_STORAGE_MAX_CONCURRENCY\", 4, 1),\n+            connection_timeout=_env_int(\n+                \"AZURE_STORAGE_CONNECTION_TIMEOUT\", 20, 1\n+            ),\n+            read_timeout=_env_int(\"AZURE_STORAGE_READ_TIMEOUT\", 120, 1),\n+        )\n+\n+    def retry_policy(self) -\u003e ExponentialRetry:\n+        return ExponentialRetry(\n+            initial_backoff=self.retry_delay,\n+            increment_base=self.retry_increment,\n+            retry_total=self.max_retries,\n+            random_jitter_range=self.retry_jitter,\n+        )\n+\n+    def configure_http_logging(self) -\u003e None:\n+        logging.getLogger(_HTTP_LOGGER).setLevel(self.http_log_level)\n+\n+    def client_options(self) -\u003e dict[str, object]:\n+        return {\n+            \"retry_policy\": self.retry_policy(),\n+            \"logging_enable\": self.http_logging_enabled,\n+            \"max_block_size\": self.max_block_size,\n+            \"max_single_put_size\": self.max_single_put_size,\n+            \"connection_timeout\": self.connection_timeout,\n+            \"read_timeout\": self.read_timeout,\n+        }\n+\n+\n+def create_sync_client(\n+    settings: StorageSettings, credential: DefaultAzureCredential\n+) -\u003e BlobServiceClient:\n+    settings.configure_http_logging()\n+    return BlobServiceClient(\n+        account_url=settings.account_url,\n+        credential=credential,\n+        **settings.client_options(),\n+    )\n+\n+\n+def create_async_client(\n+    settings: StorageSettings, credential: AsyncDefaultAzureCredential\n+) -\u003e AsyncBlobServiceClient:\n+    settings.configure_http_logging()\n+    return AsyncBlobServiceClient(\n+        account_url=settings.account_url,\n+        credential=credential,\n+        **settings.client_options(),\n+    )\n*** Add File: blob_manager/models.py\n+\"\"\"Public result models for blob operations.\"\"\"\n+\n+from __future__ import annotations\n+\n+from dataclasses import dataclass\n+from datetime import datetime\n+from typing import Generic, TypeVar\n+\n+T = TypeVar(\"T\")\n+\n+\n+@dataclass(frozen=True)\n+class OperationResult(Generic[T]):\n+    success: bool\n+    message: str\n+    value: T | None = None\n+\n+    @classmethod\n+    def ok(cls, message: str, value: T | None = None) -\u003e \"OperationResult[T]\":\n+        return cls(success=True, message=message, value=value)\n+\n+    @classmethod\n+    def fail(cls, message: str) -\u003e \"OperationResult[T]\":\n+        return cls(success=False, message=message)\n+\n+\n+@dataclass(frozen=True)\n+class BlobInfo:\n+    name: str\n+    size: int\n+    last_modified: datetime | None\n+    metadata: dict[str, str]\n+    tags: dict[str, str]\n*** Add File: blob_manager/errors.py\n+\"\"\"Azure Storage error translation shared by both service implementations.\"\"\"\n+\n+from __future__ import annotations\n+\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceExistsError,\n+    ResourceModifiedError,\n+    ResourceNotFoundError,\n+    ServiceRequestError,\n+    ServiceResponseError,\n+)\n+\n+HANDLED_AZURE_ERRORS = (\n+    ResourceNotFoundError,\n+    ResourceExistsError,\n+    ResourceModifiedError,\n+    ClientAuthenticationError,\n+    ServiceRequestError,\n+    ServiceResponseError,\n+    HttpResponseError,\n+)\n+\n+\n+def describe_storage_error(action: str, error: Exception) -\u003e str:\n+    if isinstance(error, ResourceNotFoundError):\n+        detail = \"the container or blob was not found\"\n+    elif isinstance(error, ResourceExistsError):\n+        detail = \"the blob already exists or a lease is already held\"\n+    elif isinstance(error, ResourceModifiedError):\n+        detail = \"the blob changed concurrently; retry with the latest version\"\n+    elif isinstance(error, ClientAuthenticationError):\n+        detail = \"Azure authentication failed\"\n+    elif isinstance(error, (ServiceRequestError, ServiceResponseError)):\n+        detail = \"Azure Storage could not be reached or returned an invalid response\"\n+    elif isinstance(error, HttpResponseError):\n+        status = getattr(error, \"status_code\", None)\n+        error_code = getattr(error, \"error_code\", None)\n+        if status == 403:\n+            detail = \"permission was denied\"\n+        elif status == 409 and error_code:\n+            detail = f\"the request conflicted with the blob state ({error_code})\"\n+        else:\n+            suffix = f\" ({error_code})\" if error_code else \"\"\n+            detail = f\"Azure Storage returned HTTP {status or 'error'}{suffix}\"\n+    else:\n+        detail = str(error)\n+    return f\"Could not {action}: {detail}.\"\n+\n+\n+def timeout_options(timeout: int | None) -\u003e dict[str, int]:\n+    if timeout is None:\n+        return {}\n+    if timeout \u003c= 0:\n+        raise ValueError(\"timeout must be greater than zero.\")\n+    return {\"timeout\": timeout}\n*** Add File: blob_manager/sync_service.py\n+\"\"\"Synchronous Azure Blob Storage management service.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from pathlib import Path\n+from types import TracebackType\n+from typing import Mapping\n+from uuid import uuid4\n+\n+from azure.core import MatchConditions\n+from azure.core.exceptions import ResourceNotFoundError\n+from azure.identity import DefaultAzureCredential\n+from azure.storage.blob import BlobLeaseClient, BlobServiceClient\n+\n+from .config import StorageSettings, create_sync_client\n+from .errors import HANDLED_AZURE_ERRORS, describe_storage_error, timeout_options\n+from .models import BlobInfo, OperationResult\n+\n+\n+class BlobStorageManager:\n+    \"\"\"Context-managed synchronous wrapper around common blob operations.\"\"\"\n+\n+    def __init__(self, settings: StorageSettings) -\u003e None:\n+        self.settings = settings\n+        self._credential: DefaultAzureCredential | None = None\n+        self._client: BlobServiceClient | None = None\n+\n+    def __enter__(self) -\u003e \"BlobStorageManager\":\n+        self._credential = DefaultAzureCredential()\n+        self._credential.__enter__()\n+        try:\n+            self._client = create_sync_client(self.settings, self._credential)\n+            self._client.__enter__()\n+        except Exception:\n+            self._credential.__exit__(None, None, None)\n+            self._credential = None\n+            raise\n+        return self\n+\n+    def __exit__(\n+        self,\n+        exc_type: type[BaseException] | None,\n+        exc_value: BaseException | None,\n+        traceback: TracebackType | None,\n+    ) -\u003e None:\n+        if self._client is not None:\n+            self._client.__exit__(exc_type, exc_value, traceback)\n+        if self._credential is not None:\n+            self._credential.__exit__(exc_type, exc_value, traceback)\n+        self._client = None\n+        self._credential = None\n+\n+    def _service_client(self) -\u003e BlobServiceClient:\n+        if self._client is None:\n+            raise RuntimeError(\"Use BlobStorageManager as a context manager.\")\n+        return self._client\n+\n+    def upload(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        source: str | Path,\n+        *,\n+        metadata: Mapping[str, str] | None = None,\n+        tags: Mapping[str, str] | None = None,\n+        lease: BlobLeaseClient | str | None = None,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[None]:\n+        source_path = Path(source)\n+        try:\n+            file_size = source_path.stat().st_size\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            request_options = timeout_options(timeout)\n+            try:\n+                properties = blob_client.get_blob_properties(**request_options)\n+            except ResourceNotFoundError:\n+                properties = None\n+\n+            conditions: dict[str, object]\n+            if properties is None:\n+                conditions = {\"overwrite\": False}\n+            else:\n+                conditions = {\n+                    \"overwrite\": True,\n+                    \"etag\": properties.etag,\n+                    \"match_condition\": MatchConditions.IfNotModified,\n+                }\n+\n+            with source_path.open(\"rb\") as stream:\n+                blob_client.upload_blob(\n+                    stream,\n+                    length=file_size,\n+                    metadata=dict(metadata) if metadata else None,\n+                    tags=dict(tags) if tags else None,\n+                    lease=lease,\n+                    max_concurrency=self.settings.max_concurrency,\n+                    **conditions,\n+                    **request_options,\n+                )\n+            return OperationResult.ok(\n+                f\"Uploaded {source_path} to {container}/{blob_name}.\"\n+            )\n+        except OSError as error:\n+            return OperationResult.fail(f\"Could not read {source_path}: {error}.\")\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"upload blob\", error))\n+\n+    def download(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        destination: str | Path,\n+        *,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[Path]:\n+        destination_path = Path(destination)\n+        temporary_path = destination_path.with_name(\n+            f\".{destination_path.name}.{uuid4().hex}.part\"\n+        )\n+        try:\n+            destination_path.parent.mkdir(parents=True, exist_ok=True)\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            downloader = blob_client.download_blob(\n+                max_concurrency=self.settings.max_concurrency,\n+                **timeout_options(timeout),\n+            )\n+            with temporary_path.open(\"wb\") as stream:\n+                downloader.readinto(stream)\n+            os.replace(temporary_path, destination_path)\n+            return OperationResult.ok(\n+                f\"Downloaded {container}/{blob_name} to {destination_path}.\",\n+                destination_path,\n+            )\n+        except OSError as error:\n+            return OperationResult.fail(\n+                f\"Could not write download to {destination_path}: {error}.\"\n+            )\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"download blob\", error))\n+        finally:\n+            temporary_path.unlink(missing_ok=True)\n+\n+    def list_blobs(\n+        self,\n+        container: str,\n+        *,\n+        prefix: str | None = None,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[list[BlobInfo]]:\n+        try:\n+            container_client = self._service_client().get_container_client(container)\n+            blobs = [\n+                BlobInfo(\n+                    name=item.name,\n+                    size=item.size or 0,\n+                    last_modified=item.last_modified,\n+                    metadata=dict(item.metadata or {}),\n+                    tags=dict(item.tags or {}),\n+                )\n+                for item in container_client.list_blobs(\n+                    name_starts_with=prefix,\n+                    include=[\"metadata\", \"tags\"],\n+                    **timeout_options(timeout),\n+                )\n+            ]\n+            return OperationResult.ok(\n+                f\"Listed {len(blobs)} blob(s) in {container}.\", blobs\n+            )\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"list blobs\", error))\n+\n+    def delete(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        *,\n+        lease: BlobLeaseClient | str | None = None,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[None]:\n+        try:\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            blob_client.delete_blob(\n+                delete_snapshots=\"include\",\n+                lease=lease,\n+                **timeout_options(timeout),\n+            )\n+            return OperationResult.ok(f\"Deleted {container}/{blob_name}.\")\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"delete blob\", error))\n+\n+    def acquire_lease(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        *,\n+        duration: int = 30,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[BlobLeaseClient]:\n+        try:\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            lease = blob_client.acquire_lease(\n+                lease_duration=duration, **timeout_options(timeout)\n+            )\n+            return OperationResult.ok(\n+                f\"Acquired a lease on {container}/{blob_name}.\", lease\n+            )\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(\n+                describe_storage_error(\"acquire blob lease\", error)\n+            )\n+\n+    def release_lease(\n+        self, lease: BlobLeaseClient, *, timeout: int | None = None\n+    ) -\u003e OperationResult[None]:\n+        try:\n+            lease.release(**timeout_options(timeout))\n+            return OperationResult.ok(\"Released the blob lease.\")\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(\n+                describe_storage_error(\"release blob lease\", error)\n+            )\n*** Add File: blob_manager/async_service.py\n+\"\"\"Asynchronous Azure Blob Storage management service.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import os\n+from collections.abc import AsyncIterator, Mapping\n+from pathlib import Path\n+from types import TracebackType\n+from uuid import uuid4\n+\n+from azure.core import MatchConditions\n+from azure.core.exceptions import ResourceNotFoundError\n+from azure.identity.aio import DefaultAzureCredential\n+from azure.storage.blob.aio import BlobLeaseClient, BlobServiceClient\n+\n+from .config import StorageSettings, create_async_client\n+from .errors import HANDLED_AZURE_ERRORS, describe_storage_error, timeout_options\n+from .models import BlobInfo, OperationResult\n+\n+\n+async def _file_chunks(path: Path, chunk_size: int) -\u003e AsyncIterator[bytes]:\n+    stream = await asyncio.to_thread(path.open, \"rb\")\n+    try:\n+        while chunk := await asyncio.to_thread(stream.read, chunk_size):\n+            yield chunk\n+    finally:\n+        await asyncio.to_thread(stream.close)\n+\n+\n+class AsyncBlobStorageManager:\n+    \"\"\"Context-managed asynchronous wrapper around common blob operations.\"\"\"\n+\n+    def __init__(self, settings: StorageSettings) -\u003e None:\n+        self.settings = settings\n+        self._credential: DefaultAzureCredential | None = None\n+        self._client: BlobServiceClient | None = None\n+\n+    async def __aenter__(self) -\u003e \"AsyncBlobStorageManager\":\n+        self._credential = DefaultAzureCredential()\n+        await self._credential.__aenter__()\n+        try:\n+            self._client = create_async_client(self.settings, self._credential)\n+            await self._client.__aenter__()\n+        except Exception:\n+            await self._credential.__aexit__(None, None, None)\n+            self._credential = None\n+            raise\n+        return self\n+\n+    async def __aexit__(\n+        self,\n+        exc_type: type[BaseException] | None,\n+        exc_value: BaseException | None,\n+        traceback: TracebackType | None,\n+    ) -\u003e None:\n+        if self._client is not None:\n+            await self._client.__aexit__(exc_type, exc_value, traceback)\n+        if self._credential is not None:\n+            await self._credential.__aexit__(exc_type, exc_value, traceback)\n+        self._client = None\n+        self._credential = None\n+\n+    def _service_client(self) -\u003e BlobServiceClient:\n+        if self._client is None:\n+            raise RuntimeError(\"Use AsyncBlobStorageManager as a context manager.\")\n+        return self._client\n+\n+    async def upload(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        source: str | Path,\n+        *,\n+        metadata: Mapping[str, str] | None = None,\n+        tags: Mapping[str, str] | None = None,\n+        lease: BlobLeaseClient | str | None = None,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[None]:\n+        source_path = Path(source)\n+        try:\n+            file_size = (await asyncio.to_thread(source_path.stat)).st_size\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            request_options = timeout_options(timeout)\n+            try:\n+                properties = await blob_client.get_blob_properties(**request_options)\n+            except ResourceNotFoundError:\n+                properties = None\n+\n+            conditions: dict[str, object]\n+            if properties is None:\n+                conditions = {\"overwrite\": False}\n+            else:\n+                conditions = {\n+                    \"overwrite\": True,\n+                    \"etag\": properties.etag,\n+                    \"match_condition\": MatchConditions.IfNotModified,\n+                }\n+\n+            await blob_client.upload_blob(\n+                _file_chunks(source_path, self.settings.max_block_size),\n+                length=file_size,\n+                metadata=dict(metadata) if metadata else None,\n+                tags=dict(tags) if tags else None,\n+                lease=lease,\n+                max_concurrency=self.settings.max_concurrency,\n+                **conditions,\n+                **request_options,\n+            )\n+            return OperationResult.ok(\n+                f\"Uploaded {source_path} to {container}/{blob_name}.\"\n+            )\n+        except OSError as error:\n+            return OperationResult.fail(f\"Could not read {source_path}: {error}.\")\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"upload blob\", error))\n+\n+    async def download(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        destination: str | Path,\n+        *,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[Path]:\n+        destination_path = Path(destination)\n+        temporary_path = destination_path.with_name(\n+            f\".{destination_path.name}.{uuid4().hex}.part\"\n+        )\n+        stream = None\n+        try:\n+            await asyncio.to_thread(\n+                destination_path.parent.mkdir, parents=True, exist_ok=True\n+            )\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            downloader = await blob_client.download_blob(\n+                max_concurrency=self.settings.max_concurrency,\n+                **timeout_options(timeout),\n+            )\n+            stream = await asyncio.to_thread(temporary_path.open, \"wb\")\n+            async for chunk in downloader.chunks():\n+                await asyncio.to_thread(stream.write, chunk)\n+            await asyncio.to_thread(stream.close)\n+            stream = None\n+            await asyncio.to_thread(os.replace, temporary_path, destination_path)\n+            return OperationResult.ok(\n+                f\"Downloaded {container}/{blob_name} to {destination_path}.\",\n+                destination_path,\n+            )\n+        except OSError as error:\n+            return OperationResult.fail(\n+                f\"Could not write download to {destination_path}: {error}.\"\n+            )\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"download blob\", error))\n+        finally:\n+            if stream is not None:\n+                await asyncio.to_thread(stream.close)\n+            await asyncio.to_thread(temporary_path.unlink, missing_ok=True)\n+\n+    async def list_blobs(\n+        self,\n+        container: str,\n+        *,\n+        prefix: str | None = None,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[list[BlobInfo]]:\n+        try:\n+            container_client = self._service_client().get_container_client(container)\n+            blobs = []\n+            async for item in container_client.list_blobs(\n+                name_starts_with=prefix,\n+                include=[\"metadata\", \"tags\"],\n+                **timeout_options(timeout),\n+            ):\n+                blobs.append(\n+                    BlobInfo(\n+                        name=item.name,\n+                        size=item.size or 0,\n+                        last_modified=item.last_modified,\n+                        metadata=dict(item.metadata or {}),\n+                        tags=dict(item.tags or {}),\n+                    )\n+                )\n+            return OperationResult.ok(\n+                f\"Listed {len(blobs)} blob(s) in {container}.\", blobs\n+            )\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"list blobs\", error))\n+\n+    async def delete(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        *,\n+        lease: BlobLeaseClient | str | None = None,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[None]:\n+        try:\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            await blob_client.delete_blob(\n+                delete_snapshots=\"include\",\n+                lease=lease,\n+                **timeout_options(timeout),\n+            )\n+            return OperationResult.ok(f\"Deleted {container}/{blob_name}.\")\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(describe_storage_error(\"delete blob\", error))\n+\n+    async def acquire_lease(\n+        self,\n+        container: str,\n+        blob_name: str,\n+        *,\n+        duration: int = 30,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[BlobLeaseClient]:\n+        try:\n+            blob_client = self._service_client().get_blob_client(\n+                container=container, blob=blob_name\n+            )\n+            lease = await blob_client.acquire_lease(\n+                lease_duration=duration, **timeout_options(timeout)\n+            )\n+            return OperationResult.ok(\n+                f\"Acquired a lease on {container}/{blob_name}.\", lease\n+            )\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(\n+                describe_storage_error(\"acquire blob lease\", error)\n+            )\n+\n+    async def release_lease(\n+        self, lease: BlobLeaseClient, *, timeout: int | None = None\n+    ) -\u003e OperationResult[None]:\n+        try:\n+            await lease.release(**timeout_options(timeout))\n+            return OperationResult.ok(\"Released the blob lease.\")\n+        except HANDLED_AZURE_ERRORS as error:\n+            return OperationResult.fail(\n+                describe_storage_error(\"release blob lease\", error)\n+            )\n*** Add File: main.py\n+\"\"\"Demonstrate synchronous and asynchronous blob management operations.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+from pathlib import Path\n+\n+from blob_manager import (\n+    AsyncBlobStorageManager,\n+    BlobStorageManager,\n+    OperationResult,\n+    StorageSettings,\n+)\n+\n+CONTAINER = os.getenv(\"AZURE_STORAGE_CONTAINER\", \"blob-manager-demo\")\n+SAMPLE_FILE = Path(\"sample.txt\")\n+SYNC_DOWNLOAD = Path(\"downloads/sync-sample.txt\")\n+ASYNC_DOWNLOAD = Path(\"downloads/async-sample.txt\")\n+\n+\n+def print_result(step: str, result: OperationResult[object]) -\u003e None:\n+    marker = \"OK\" if result.success else \"ERROR\"\n+    print(f\"[{marker}] {step}: {result.message}\")\n+\n+\n+def print_listing(result: OperationResult[object]) -\u003e None:\n+    print_result(\"List\", result)\n+    if result.success and isinstance(result.value, list):\n+        for blob in result.value:\n+            print(f\"      {blob.name} ({blob.size} bytes, tags={blob.tags})\")\n+\n+\n+def run_sync_demo(settings: StorageSettings) -\u003e None:\n+    blob_name = \"demo/sync-sample.txt\"\n+    print(\"\\n--- Synchronous demo ---\")\n+    SAMPLE_FILE.write_text(\"Initial synchronous content.\\n\", encoding=\"utf-8\")\n+\n+    with BlobStorageManager(settings) as manager:\n+        print_result(\n+            \"Upload\",\n+            manager.upload(\n+                CONTAINER,\n+                blob_name,\n+                SAMPLE_FILE,\n+                metadata={\"source\": \"sync-demo\"},\n+                tags={\"project\": \"blob-manager\", \"mode\": \"sync\"},\n+                timeout=120,\n+            ),\n+        )\n+        print_listing(manager.list_blobs(CONTAINER, prefix=\"demo/\", timeout=60))\n+        print_result(\n+            \"Download\",\n+            manager.download(CONTAINER, blob_name, SYNC_DOWNLOAD, timeout=120),\n+        )\n+\n+        lease_result = manager.acquire_lease(\n+            CONTAINER, blob_name, duration=30, timeout=30\n+        )\n+        print_result(\"Acquire lease\", lease_result)\n+        if lease_result.success and lease_result.value is not None:\n+            lease = lease_result.value\n+            try:\n+                SAMPLE_FILE.write_text(\n+                    \"Synchronous content overwritten while leased.\\n\",\n+                    encoding=\"utf-8\",\n+                )\n+                print_result(\n+                    \"Leased overwrite\",\n+                    manager.upload(\n+                        CONTAINER,\n+                        blob_name,\n+                        SAMPLE_FILE,\n+                        tags={\"project\": \"blob-manager\", \"mode\": \"sync\"},\n+                        lease=lease,\n+                        timeout=120,\n+                    ),\n+                )\n+            finally:\n+                print_result(\"Release lease\", manager.release_lease(lease, timeout=30))\n+\n+        print_result(\"Delete\", manager.delete(CONTAINER, blob_name, timeout=60))\n+\n+\n+async def run_async_demo(settings: StorageSettings) -\u003e None:\n+    blob_name = \"demo/async-sample.txt\"\n+    print(\"\\n--- Asynchronous demo ---\")\n+    await asyncio.to_thread(\n+        SAMPLE_FILE.write_text, \"Initial asynchronous content.\\n\", encoding=\"utf-8\"\n+    )\n+\n+    async with AsyncBlobStorageManager(settings) as manager:\n+        print_result(\n+            \"Upload\",\n+            await manager.upload(\n+                CONTAINER,\n+                blob_name,\n+                SAMPLE_FILE,\n+                metadata={\"source\": \"async-demo\"},\n+                tags={\"project\": \"blob-manager\", \"mode\": \"async\"},\n+                timeout=120,\n+            ),\n+        )\n+        print_listing(\n+            await manager.list_blobs(CONTAINER, prefix=\"demo/\", timeout=60)\n+        )\n+        print_result(\n+            \"Download\",\n+            await manager.download(\n+                CONTAINER, blob_name, ASYNC_DOWNLOAD, timeout=120\n+            ),\n+        )\n+\n+        lease_result = await manager.acquire_lease(\n+            CONTAINER, blob_name, duration=30, timeout=30\n+        )\n+        print_result(\"Acquire lease\", lease_result)\n+        if lease_result.success and lease_result.value is not None:\n+            lease = lease_result.value\n+            try:\n+                await asyncio.to_thread(\n+                    SAMPLE_FILE.write_text,\n+                    \"Asynchronous content overwritten while leased.\\n\",\n+                    encoding=\"utf-8\",\n+                )\n+                print_result(\n+                    \"Leased overwrite\",\n+                    await manager.upload(\n+                        CONTAINER,\n+                        blob_name,\n+                        SAMPLE_FILE,\n+                        tags={\"project\": \"blob-manager\", \"mode\": \"async\"},\n+                        lease=lease,\n+                        timeout=120,\n+                    ),\n+                )\n+            finally:\n+                print_result(\n+                    \"Release lease\", await manager.release_lease(lease, timeout=30)\n+                )\n+\n+        print_result(\n+            \"Delete\", await manager.delete(CONTAINER, blob_name, timeout=60)\n+        )\n+\n+\n+def main() -\u003e int:\n+    logging.basicConfig(\n+        level=os.getenv(\"APP_LOG_LEVEL\", \"INFO\"),\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    try:\n+        settings = StorageSettings.from_env()\n+    except (TypeError, ValueError) as error:\n+        print(f\"[ERROR] Configuration: {error}\")\n+        return 2\n+\n+    run_sync_demo(settings)\n+    asyncio.run(run_async_demo(settings))\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/requirements.txt
@@ -1,0 +1,3 @@
+azure-identity>=1.17.0,<2.0.0
+azure-storage-blob>=12.25.0,<13.0.0
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/__init__.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/__init__.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/__init__.py
@@ -1,0 +1,15 @@
+"""Reusable synchronous and asynchronous Azure Blob Storage managers."""
+
+from .async_service import AsyncBlobStorageManager
+from .config import StorageSettings
+from .models import BlobInfo, OperationResult
+from .sync_service import BlobStorageManager
+
+__all__ = [
+    "AsyncBlobStorageManager",
+    "BlobInfo",
+    "BlobStorageManager",
+    "OperationResult",
+    "StorageSettings",
+]
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/config.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/config.py
create file mode 100644
index 0000000..0000000
--- a/de
... (truncated)
```

### ✅ 11. powershell — 10151ms

**Input:**

```json
{"command":"python -m pip install -r requirements.txt --quiet \u0026\u0026 python -m compileall -q blob_manager main.py \u0026\u0026 python -c \"from blob_manager import BlobStorageManager, AsyncBlobStorageManager, StorageSettings; print('imports-ok')\" \u0026\u0026 python -m pip check","description":"Install dependencies and validate Python project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
imports-ok
onnx 1.21.0 requires ml-dtypes, which is not installed.
numba 0.62.1 has requirement numpy<2.4,>=1.22, but you have numpy 2.4.4.
streamlit 1.51.0 has requirement protobuf<7,>=3.20, but you have protobuf 7.34.1.
<shellId: 0 completed with exit code 1>
```

### ✅ 12. powershell — 12541ms

**Input:**

```json
{"command":"python -c \"import inspect; from azure.storage.blob import BlobClient; from azure.storage.blob.aio import BlobClient as A; print(inspect.signature(BlobClient.upload_blob)); print(inspect.signature(A.upload_blob)); print(inspect.signature(BlobClient.acquire_lease)); print(inspect.signature(A.acquire_lease))\"","description":"Inspect Azure Blob SDK method signatures","initial_wait":30,"mode":"sync"}
```

**Output:**

```
[ERROR] Configuration: AZURE_STORAGE_ACCOUNT_URL is required (for example, https://<account>.blob.core.windows.net).
<shellId: 3 completed with exit code 2>
```

### ✅ 13. powershell — 11908ms

**Input:**

```json
{"command":"python -c \"from blob_manager.config import StorageSettings; s=StorageSettings('https://example.blob.core.windows.net'); print(type(s.retry_policy()).__name__); print(s.client_options())\"","description":"Exercise storage configuration","initial_wait":30,"mode":"sync"}
```

**Output:**

```
ExponentialRetry
{'retry_policy': <azure.storage.blob._shared.policies.ExponentialRetry object at 0x000001DCA1A74B90>, 'logging_enable': False, 'max_block_size': 4194304, 'max_single_put_size': 67108864, 'connection_timeout': 20, 'read_timeout': 120}
<shellId: 2 completed with exit code 0>
```

### ✅ 14. powershell — 11343ms

**Input:**

```json
{"command":"$env:AZURE_STORAGE_ACCOUNT_URL=''; python main.py; if ($LASTEXITCODE -ne 2) { exit 1 }","description":"Check graceful configuration failure","initial_wait":30,"mode":"sync"}
```

**Output:**

```
(self, data: Union[bytes, str, Iterable[~AnyStr], IO[bytes]], blob_type: Union[str, azure.storage.blob._models.BlobType] = <BlobType.BLOCKBLOB: 'BlockBlob'>, length: Optional[int] = None, metadata: Optional[Dict[str, str]] = None, **kwargs: Any) -> Dict[str, Any]
(self, data: Union[bytes, str, Iterable[~AnyStr], AsyncIterable[~AnyStr], IO[bytes]], blob_type: Union[str, azure.storage.blob._models.BlobType] = <BlobType.BLOCKBLOB: 'BlockBlob'>, length: Optional[int] = None, metadata: Optional[Dict[str, str]] = None, **kwargs: Any) -> Dict[str, Any]
(self, lease_duration: int = -1, lease_id: Optional[str] = None, **kwargs: Any) -> azure.storage.blob._lease.BlobLeaseClient
(self, lease_duration: int = -1, lease_id: Optional[str] = None, **kwargs: Any) -> azure.storage.blob.aio._lease_async.BlobLeaseClient
<shellId: 1 completed with exit code 0>
```

### ✅ 15. powershell — 3839ms

**Input:**

```json
{"command":"@'\nimport asyncio\nimport tempfile\nfrom pathlib import Path\nfrom unittest.mock import AsyncMock, Mock\n\nfrom azure.core import MatchConditions\nfrom blob_manager.async_service import AsyncBlobStorageManager\nfrom blob_manager.config import StorageSettings\nfrom blob_manager.sync_service import BlobStorageManager\n\nsettings = StorageSettings('https://example.blob.core.windows.net', max_block_size=4, max_concurrency=2)\n\nwith tempfile.TemporaryDirectory() as directory:\n    source = Path(directory) / 'source.bin'\n    source.write_bytes(b'abcdefghij')\n\n    blob = Mock()\n    blob.get_blob_properties.return_value = Mock(etag='etag-1')\n    blob.upload_blob.return_value = {}\n    client = Mock()\n    client.get_blob_client.return_value = blob\n    manager = BlobStorageManager(settings)\n    manager._client = client\n    result = manager.upload('container', 'blob', source, tags={'kind': 'test'}, timeout=10)\n    assert result.success, result.message\n    kwargs = blob.upload_blob.call_args.kwargs\n    assert kwargs['etag'] == 'etag-1'\n    assert kwargs['match_condition'] is MatchConditions.IfNotModified\n    assert kwargs['length'] == 10\n    assert kwargs['timeout'] == 10\n    assert hasattr(blob.upload_blob.call_args.args[0], 'read')\n\n    async def async_check():\n        async_blob = Mock()\n        async_blob.get_blob_properties = AsyncMock(return_value=Mock(etag='etag-2'))\n        captured = []\n        async def consume(data, **kwargs):\n            async for chunk in data:\n                captured.append(chunk)\n            return {}\n        async_blob.upload_blob = AsyncMock(side_effect=consume)\n        async_client = Mock()\n        async_client.get_blob_client.return_value = async_blob\n        async_manager = AsyncBlobStorageManager(settings)\n        async_manager._client = async_client\n        async_result = await async_manager.upload('container', 'blob', source, timeout=11)\n        assert async_result.success, async_result.message\n        kwargs = async_blob.upload_blob.call_args.kwargs\n        assert kwargs['etag'] == 'etag-2'\n        assert kwargs['match_condition'] is MatchConditions.IfNotModified\n        assert kwargs['length'] == 10\n        assert kwargs['timeout'] == 11\n        assert captured == [b'abcd', b'efgh', b'ij']\n\n    asyncio.run(async_check())\n\nprint('behavior-ok')\n'@ | python -","description":"Test streaming and concurrency behavior locally","initial_wait":60,"mode":"sync"}
```

**Output:**

```
behavior-ok
<shellId: 4 completed with exit code 0>
```

### ✅ 16. apply_patch — 1599ms

**Input:**

```json
"*** Begin Patch\n*** Update File: blob_manager/config.py\n@@\n-        log_level = os.getenv(\"AZURE_STORAGE_HTTP_LOG_LEVEL\", \"WARNING\").upper()\n-        if log_level not in logging.getLevelNamesMapping():\n+        log_level = os.getenv(\"AZURE_STORAGE_HTTP_LOG_LEVEL\", \"WARNING\").upper()\n+        if not isinstance(getattr(logging, log_level, None), int):\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/config.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/config.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/config.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-microsoft-skill-1707574739/blob_manager/config.py
@@ -63,7 +63,7 @@
             raise ValueError("AZURE_STORAGE_ACCOUNT_URL must use HTTPS.")
 
         log_level = os.getenv("AZURE_STORAGE_HTTP_LOG_LEVEL", "WARNING").upper()
-        if log_level not in logging.getLevelNamesMapping():
+        if not isinstance(getattr(logging, log_level, None), int):
             raise ValueError(
                 f"AZURE_STORAGE_HTTP_LOG_LEVEL is not a valid logging level: {log_level!r}."
             )

```

### ✅ 17. powershell — 4839ms

**Input:**

```json
{"command":"python -m compileall -q blob_manager main.py \u0026\u0026 python -c \"from blob_manager.config import StorageSettings; assert StorageSettings('https://example.blob.core.windows.net').client_options()['logging_enable'] is False; print('validation-ok')\"; if (Test-Path 'blob_manager\\__pycache__') { Remove-Item -Recurse -Force 'blob_manager\\__pycache__' }; if (Test-Path '__pycache__') { Remove-Item -Recurse -Force '__pycache__' }; Get-ChildItem -Recurse -File | ForEach-Object { $_.FullName.Substring($PWD.Path.Length + 1) }","description":"Validate final files and remove bytecode","initial_wait":60,"mode":"sync"}
```

**Output:**

```
validation-ok
main.py
requirements.txt
blob_manager\__init__.py
blob_manager\async_service.py
blob_manager\config.py
blob_manager\errors.py
blob_manager\models.py
blob_manager\sync_service.py
<shellId: 5 completed with exit code 0>
```

## Generated Files

- `blob_manager\__init__.py`
- `blob_manager\async_service.py`
- `blob_manager\config.py`
- `blob_manager\errors.py`
- `blob_manager\models.py`
- `blob_manager\sync_service.py`
- `main.py`
- `requirements.txt`

## Copilot Response

I’m grounding the implementation in the Azure Blob SDK guidance and checking the workspace before creating the project.

The workspace is empty. I’m confirming the current SDK signatures for conditional uploads, leases, retries, and async streaming before writing the modules.

I’ve settled on optimistic ETag conditions for normal updates plus lease IDs for explicit locked overwrites, with chunked file I/O and result objects so expected Azure failures remain caller-visible without terminating the program.

The project files are in place. I’m now validating imports and SDK compatibility against the declared dependency range, then I’ll correct any signature or runtime issues found.

The SDK signatures confirm async iterables and lease-aware uploads are supported. I’m doing a focused behavioral check of transfer tuning, ETag conditions, bounded streaming, and cleanup paths rather than contacting Azure.

Created the reusable Azure Blob Storage project with:

- Sync and async service modules for upload, download, listing, deletion, and lease management.
- Memory-efficient chunked transfers suitable for multi-gigabyte files.
- ETag-based optimistic concurrency and lease-aware overwrites.
- `DefaultAzureCredential` authentication using `AZURE_STORAGE_ACCOUNT_URL`.
- Configurable exponential retries, network timeouts, concurrency, and HTTP logging.
- Graceful structured error results for missing blobs, permissions, lease conflicts, authentication, and network failures.
- A complete sync-then-async demo in `main.py`.
- Azure SDK dependencies in `requirements.txt`.

Set `AZURE_STORAGE_ACCOUNT_URL` and optionally `AZURE_STORAGE_CONTAINER`, then run:

`python main.py`

## Reviewer Notes (LLM-as-Judge)

### Summary

The exception-handling criterion is fully met.

### Strengths

- Consistent error handling across sync and async operations
- Clear translation of common Azure Storage failures
- Storage errors are returned as structured failure results rather than crashing

## Grader Results

- blob-storage-manager.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (10/10)
      - Configures a custom retry policy with exponential backoff: Pass
      - Enables HTTP request/response logging for debugging: Pass
      - Does NOT use connection strings or account keys: Pass
      - Handles large file upload efficiently without loading the entire file into memory: Pass
      - Supports blob index tags on upload (distinct from metadata): Pass
      - Implements blob lease acquisition to prevent concurrent overwrites: Pass
      - Provides both sync and async implementations: Pass
      - Catches and handles storage-specific errors from the Azure SDK: Pass
      - Handles lease conflicts when another client holds a lease: Pass
      - Includes per-operation timeout configuration: Pass
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
hyoka run --prompt-id storage-dp-python-blob-manager --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
