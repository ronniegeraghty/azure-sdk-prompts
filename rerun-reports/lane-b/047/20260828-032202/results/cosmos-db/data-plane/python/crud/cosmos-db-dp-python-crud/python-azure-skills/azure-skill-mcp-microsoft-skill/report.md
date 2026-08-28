# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 352.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 10/13 |
| Duration | 352.0s |
| Timestamp | 2026-08-27T19:22:02Z |
| Files Generated | 2 |
| Event Count | 2440 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 261.6s |
| Review | 89.2s |
| **Total** | **352.0s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-cosmos-db-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=256251 out=2466 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 4460ms

**Input:**

```json
{"skill":"azure-cosmos-db-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Cosmos DB Service Implementation

Build production-grade Azure Cosmos DB NoSQL services following clean code, security best practices, and TDD principles.

## Installation

```bash
pip install azure-cosmos azure-identity
```

## Environment Variables

```bash
COSMOS_ENDPOINT=https://<account>.documents.azure.com:443/  # Required for all auth methods
COSMOS_DATABASE_NAME=<database-name>  # Required for all auth methods
COSMOS_CONTAINER_ID=<container-id>  # Required for all auth methods
# For emulator only (not production)
COSMOS_KEY=<emulator-key>  # Only required for key-based auth or emulator
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

**DefaultAzureCredential (preferred)**:
```python
import os
from azure.cosmos import CosmosClient
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<speci
... (truncated)
```

### ✅ 2. glob — 188435ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2445360638","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 192498ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\references\\client-setup.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md
@@ -1,215 +1,215 @@
 # Cosmos DB Client Setup
 
 ## Table of Contents
 
 1. [Dual Authentication Strategy](#dual-authentication-strategy)
 2. [Singleton Pattern](#singleton-pattern)
 3. [Async Wrapping](#async-wrapping)
 4. [Configuration Management](#configuration-management)
 5. [Connection Reset](#connection-reset)
 6. [Complete Implementation](#complete-implementation)
 
 ---
 
 ## Dual Authentication Strategy
 
 Use `DefaultAzureCredential` for Azure deployments and key-based auth only for the local emulator:
 
 ```python
 from azure.cosmos import CosmosClient
 from azure.identity import DefaultAzureCredential
 
 def _is_emulator_endpoint(endpoint: str) -> bool:
     """Detect Cosmos emulator by endpoint URL."""
     return "localhost" in endpoint.lower() or "127.0.0.1" in endpoint
 
 def _create_client(settings) -> CosmosClient:
     if _is_emulator_endpoint(settings.cosmos_endpoint):
         # Emulator: use well-known key, disable SSL verification
         return CosmosClient(
             url=settings.cosmos_endpoint,
             credential=settings.cosmos_key,
             connection_verify=False  # Emulator uses self-signed cert
         )
     else:
         # Azure: use RBAC via DefaultAzureCredential
         credential = DefaultAzureCredential()
         return CosmosClient(
             url=settings.cosmos_endpoint,
   
... (truncated)
```

### ✅ 4. view — 191280ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\references\\error-handling.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md
@@ -1,271 +1,271 @@
 # Partition Key Strategies
 
 ## Table of Contents
 
 1. [Partition Key Fundamentals](#partition-key-fundamentals)
 2. [Common Strategies](#common-strategies)
 3. [Cross-Partition Queries](#cross-partition-queries)
 4. [Move Operations](#move-operations)
 5. [Query Optimization](#query-optimization)
 
 ---
 
 ## Partition Key Fundamentals
 
 ### What is a Partition Key?
 
 The partition key determines data distribution and query efficiency:
 
 - **Same partition key** → Data co-located → Fast queries, transactional writes
 - **Different partition keys** → Data distributed → Cross-partition queries required
 
 ### Choosing a Partition Key
 
 Good partition keys have:
 
 1. **High cardinality** — Many distinct values to distribute load
 2. **Even distribution** — No single value dominates storage/throughput
 3. **Query alignment** — Most queries filter by partition key
 
 ---
 
 ## Common Strategies
 
 ### Strategy 1: Self-Partitioned Entities
 
 Use the entity's own ID when entities are accessed individually:
 
 ```python
 # Workspaces partition by their own ID
 class WorkspaceInDB(BaseModel):
     id: str                    # Also used as partition key
     name: str
     doc_type: str = "workspace"
 
 # Query always includes the workspace ID
 doc = await get_document(workspace_id, partition_key=worksp
... (truncated)
```

### ✅ 5. view — 189763ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\references\\partitioning.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md
@@ -1,383 +1,383 @@
 # Error Handling Patterns
 
 ## Table of Contents
 
 1. [Cosmos Exception Types](#cosmos-exception-types)
 2. [Client-Level Error Handling](#client-level-error-handling)
 3. [Service-Level Error Handling](#service-level-error-handling)
 4. [Router-Level Error Mapping](#router-level-error-mapping)
 5. [Logging Patterns](#logging-patterns)
 6. [Retry Strategies](#retry-strategies)
 
 ---
 
 ## Cosmos Exception Types
 
 ### Common Exceptions
 
 ```python
 from azure.cosmos.exceptions import (
     CosmosResourceNotFoundError,    # 404 - Document not found
     CosmosResourceExistsError,      # 409 - Conflict (duplicate ID)
     CosmosHttpResponseError,        # Base class for HTTP errors
 )
 ```
 
 | Exception | HTTP Status | Common Cause |
 |-----------|-------------|--------------|
 | `CosmosResourceNotFoundError` | 404 | Document/container doesn't exist |
 | `CosmosResourceExistsError` | 409 | Document with ID already exists |
 | `CosmosHttpResponseError` | 429 | Rate limiting (too many RU) |
 | `CosmosHttpResponseError` | 503 | Service unavailable |
 
 ---
 
 ## Client-Level Error Handling
 
 Handle exceptions in the Cosmos client module:
 
 ```python
 from azure.cosmos.exceptions import CosmosResourceNotFoundError, CosmosHttpResponseError
 import logging
 
 logger = logging.getLogger(__name__)
 
 
 async d
... (truncated)
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 192501ms

**Input:**

```json
{"intent":"Generate a local-only Python script demonstrating Azure Cosmos DB NoSQL CRUD with DefaultAzureCredential, parameterized queries, and azure.cosmos exception handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. glob — 6315ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. glob — 5960ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py","pattern":"**/examples.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. view — 6716ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\assets\\cosmos_client_template.py"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py
@@ -1,226 +1,226 @@
 """
 Cosmos DB Client Module Template
 
 Production-ready Azure Cosmos DB NoSQL client with:
 - Dual authentication (DefaultAzureCredential for Azure, key for emulator)
 - Singleton pattern for connection reuse
 - Async wrapping via run_in_threadpool
 - Graceful error handling
 
 Usage:
     from app.db.cosmos import get_container, upsert_document, get_document
 """
 from __future__ import annotations
 
 import logging
 from typing import Any, Optional
 
 from azure.cosmos import ContainerProxy, CosmosClient
 from azure.cosmos.exceptions import CosmosResourceNotFoundError
 from azure.identity import DefaultAzureCredential
 from starlette.concurrency import run_in_threadpool
 
 from app.config import settings
 
 logger = logging.getLogger(__name__)
 
 # Module-level singleton state
 _cosmos_container: Optional[ContainerProxy] = None
 _credential: Optional[DefaultAzureCredential] = None
 _init_attempted: bool = False
 
 
 def _is_emulator_endpoint(endpoint: str) -> bool:
     """Detect if endpoint is Cosmos emulator."""
     return "localhost" in endpoint.lower() or "127.0.0.1" in endpoint
 
 
 def _create_client() -> CosmosClient:
     """Create Cosmos client with appropriate authentication."""
     global _credential
 
     if _is_emulator_endpoint(settings.cosmos_endpoint):
         logger.i
... (truncated)
```

### ✅ 10. apply_patch — 679ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+from __future__ import annotations\n+\n+import os\n+import sys\n+from typing import Any\n+from urllib.parse import urlparse\n+\n+from azure.cosmos import CosmosClient, PartitionKey\n+from azure.cosmos.exceptions import (\n+    CosmosHttpResponseError,\n+    CosmosResourceNotFoundError,\n+)\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+ITEM_ID = \"item-001\"\n+ITEM_CATEGORY = \"electronics\"\n+LOCAL_HOSTS = {\"localhost\", \"127.0.0.1\", \"::1\"}\n+\n+\n+def get_local_configuration() -\u003e tuple[str, str]:\n+    endpoint = os.getenv(\"COSMOS_ENDPOINT\", \"https://localhost:8081\")\n+    key = os.getenv(\"COSMOS_KEY\")\n+\n+    if urlparse(endpoint).hostname not in LOCAL_HOSTS:\n+        raise ValueError(\n+            \"COSMOS_ENDPOINT must target a local Cosmos DB emulator.\"\n+        )\n+    if not key:\n+        raise ValueError(\n+            \"Set COSMOS_KEY to the key shown by your local Cosmos DB emulator.\"\n+        )\n+\n+    return endpoint, key\n+\n+\n+def run_crud_operations() -\u003e None:\n+    endpoint, key = get_local_configuration()\n+\n+    with CosmosClient(\n+        url=endpoint,\n+        credential=key,\n+        connection_verify=False,\n+    ) as client:\n+        database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+        container = database.create_container_if_not_exists(\n+            id=CONTAINER_NAME,\n+            partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+            offer_throughput=400,\n+        )\n+\n+        item: dict[str, Any] = {\n+            \"id\": ITEM_ID,\n+            \"category\": ITEM_CATEGORY,\n+            \"name\": \"Wireless Headphones\",\n+            \"quantity\": 10,\n+        }\n+\n+        upserted_item = container.upsert_item(body=item)\n+        print(f\"Upserted: {upserted_item}\")\n+\n+        read_item = container.read_item(\n+            item=ITEM_ID,\n+            partition_key=ITEM_CATEGORY,\n+        )\n+        print(f\"Read: {read_item}\")\n+\n+        query = \"SELECT * FROM items i WHERE i.category = @category\"\n+        parameters = [{\"name\": \"@category\", \"value\": ITEM_CATEGORY}]\n+        queried_items = list(\n+            container.query_items(\n+                query=query,\n+                parameters=parameters,\n+                partition_key=ITEM_CATEGORY,\n+            )\n+        )\n+        print(f\"Query results: {queried_items}\")\n+\n+        updated_item = dict(read_item)\n+        updated_item[\"quantity\"] = 25\n+        replaced_item = container.replace_item(\n+            item=ITEM_ID,\n+            body=updated_item,\n+        )\n+        print(f\"Replaced: {replaced_item}\")\n+\n+        container.delete_item(\n+            item=ITEM_ID,\n+            partition_key=ITEM_CATEGORY,\n+        )\n+        print(f\"Deleted item {ITEM_ID!r}.\")\n+\n+\n+def main() -\u003e int:\n+    try:\n+        run_crud_operations()\n+    except CosmosResourceNotFoundError as error:\n+        print(f\"Cosmos DB resource was not found: {error}\", file=sys.stderr)\n+        return 1\n+    except CosmosHttpResponseError as error:\n+        print(\n+            f\"Cosmos DB request failed \"\n+            f\"(status {error.status_code}): {error.message}\",\n+            file=sys.stderr,\n+        )\n+        return 1\n+    except ValueError as error:\n+        print(f\"Configuration error: {error}\", file=sys.stderr)\n+        return 2\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.7.0,\u003c5.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2445360638/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2445360638/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2445360638/cosmos_crud.py
@@ -1,0 +1,117 @@
+from __future__ import annotations
+
+import os
+import sys
+from typing import Any
+from urllib.parse import urlparse
+
+from azure.cosmos import CosmosClient, PartitionKey
+from azure.cosmos.exceptions import (
+    CosmosHttpResponseError,
+    CosmosResourceNotFoundError,
+)
+
+DATABASE_NAME = "TestDB"
+CONTAINER_NAME = "Items"
+PARTITION_KEY_PATH = "/category"
+ITEM_ID = "item-001"
+ITEM_CATEGORY = "electronics"
+LOCAL_HOSTS = {"localhost", "127.0.0.1", "::1"}
+
+
+def get_local_configuration() -> tuple[str, str]:
+    endpoint = os.getenv("COSMOS_ENDPOINT", "https://localhost:8081")
+    key = os.getenv("COSMOS_KEY")
+
+    if urlparse(endpoint).hostname not in LOCAL_HOSTS:
+        raise ValueError(
+            "COSMOS_ENDPOINT must target a local Cosmos DB emulator."
+        )
+    if not key:
+        raise ValueError(
+            "Set COSMOS_KEY to the key shown by your local Cosmos DB emulator."
+        )
+
+    return endpoint, key
+
+
+def run_crud_operations() -> None:
+    endpoint, key = get_local_configuration()
+
+    with CosmosClient(
+        url=endpoint,
+        credential=key,
+        connection_verify=False,
+    ) as client:
+        database = client.create_database_if_not_exists(id=DATABASE_NAME)
+        container = database.create_container_if_not_exists(
+            id=CONTAINER_NAME,
+            partition_key=PartitionKey(path=PARTITION_KEY_PATH),
+            offer_throughput=400,
+        )
+
+      
... (truncated)
```

### ✅ 11. powershell — 1503ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile the Cosmos DB CRUD script","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’m checking the SDK’s required CRUD and exception-handling patterns, then I’ll create a locally configurable script and dependency file.

The SDK guidance supports an emulator-key path and parameterized single-partition queries. I’m implementing that path only, so the sample cannot accidentally modify a live Azure account.

Created:

- `cosmos_crud.py` — creates `TestDB/Items`, then upserts, reads, queries, replaces, and deletes an item with Cosmos SDK exception handling.
- `requirements.txt` — includes `azure-cosmos>=4.7.0,<5.0.0`.

The script targets a **local Cosmos DB emulator only**. Set `COSMOS_KEY` and optionally `COSMOS_ENDPOINT` (defaults to `https://localhost:8081`), then run:

```powershell
pip install -r requirements.txt
python .\cosmos_crud.py
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Azure Cosmos DB HTTP exceptions and the resource-not-found subclass are handled appropriately.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/6)
      - `azure-cosmos` pip package: Pass
      - `CosmosClient` creation: Pass
      - `database_client.create_database_if_not_exists()`: Pass
      - `database.create_container_if_not_exists()` with `PartitionKey`: Pass
      - `container.query_items()` with `enable_cross_partition_query`: Fail
      - `CosmosHttpResponseError` exception handling: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Fail
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
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 14.3% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.8333** | **72.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
