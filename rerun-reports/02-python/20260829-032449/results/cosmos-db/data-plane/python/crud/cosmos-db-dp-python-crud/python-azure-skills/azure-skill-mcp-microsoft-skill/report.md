# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 152.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 10/13 |
| Duration | 152.5s |
| Timestamp | 2026-08-28T23:04:00Z |
| Files Generated | 2 |
| Event Count | 1974 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 66.0s |
| Review | 85.4s |
| **Total** | **152.5s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=208918 out=1673 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 680ms

**Input:**

```json
{"skill":"azure-cosmos-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Cosmos DB SDK for Python

Client library for Azure Cosmos DB NoSQL API — globally distributed, multi-model database.

## Installation

```bash
pip install azure-cosmos azure-identity
```

## Environment Variables

```bash
COSMOS_ENDPOINT=https://<account>.documents.azure.com:443/  # Required for all auth methods
COSMOS_DATABASE=mydb  # Required for all auth methods
COSMOS_CONTAINER=mycontainer  # Required for all auth methods
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
import os
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential
from azure.cosmos import CosmosClient

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
credential = DefaultAzureCredential(require_envvar=True)
# Or use a specific credential directly in production:
# See https://learn.microsoft.com/python/api/overview/azure/identity-readme
... (truncated)
```

### ✅ 2. glob — 2129ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1987094751","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 1821ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 2782ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py\\references\\partitioning.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/partitioning.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/partitioning.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/partitioning.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/partitioning.md
@@ -1,347 +1,347 @@
 # Partition Key Strategies
 
 Comprehensive guide to partition key design and optimization for Azure Cosmos DB.
 
 ## Partition Key Fundamentals
 
 ### What is a Partition Key?
 
 A partition key determines how data is distributed across physical partitions:
 
 - **Logical partition**: All items with the same partition key value
 - **Physical partition**: Storage unit that holds one or more logical partitions
 - **Maximum logical partition size**: 20 GB
 
 ```python
 from azure.cosmos import PartitionKey
 
 # Define partition key at container creation
 container = database.create_container_if_not_exists(
     id="orders",
     partition_key=PartitionKey(path="/customer_id")
 )
 ```
 
 ## Choosing a Partition Key
 
 ### Good Partition Key Characteristics
 
 | Characteristic | Why It Matters |
 |---------------|----------------|
 | High cardinality | Many unique values = even distribution |
 | Even distribution | Prevents hot partitions |
 | Frequently used in queries | Enables efficient single-partition queries |
 | Immutable | Partition key cannot be changed after item creation |
 
 ### Common Partition Key Patterns
 
 ```python
 # E-commerce orders - partition by customer
 container = database.create_container_if_not_exists(
     id="orders",
     partition_key=PartitionKey(path="/customer_id")
 )
 # Query: "Get all orders for
... (truncated)
```

### ✅ 5. view — 2459ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py\\references\\query-patterns.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/query-patterns.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/query-patterns.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/query-patterns.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-py/references/query-patterns.md
@@ -1,396 +1,396 @@
 # Query Patterns Reference
 
 Advanced query patterns for Azure Cosmos DB NoSQL API.
 
 ## Parameterized Queries
 
 Always use parameters to prevent injection and enable query plan caching:
 
 ```python
 # GOOD: Parameterized query
 query = "SELECT * FROM c WHERE c.category = @category AND c.price < @max_price"
 items = container.query_items(
     query=query,
     parameters=[
         {"name": "@category", "value": "electronics"},
         {"name": "@max_price", "value": 500}
     ],
     partition_key="electronics"
 )
 
 # BAD: String interpolation (injection risk, no caching)
 category = "electronics"
 query = f"SELECT * FROM c WHERE c.category = '{category}'"  # Never do this!
 ```
 
 ## Query Optimization
 
 ### Single Partition vs Cross-Partition
 
 ```python
 # EFFICIENT: Single partition query (always preferred)
 items = container.query_items(
     query="SELECT * FROM c WHERE c.status = 'active'",
     partition_key="tenant-123"
 )
 
 # EXPENSIVE: Cross-partition query (use sparingly)
 items = container.query_items(
     query="SELECT * FROM c WHERE c.status = 'active'",
     enable_cross_partition_query=True  # Fans out to all partitions
 )
 ```
 
 ### Point Reads vs Queries
 
 ```python
 # MOST EFFICIENT: Point read (1 RU for 1KB document)
 item = container.read_item(
     item="doc-123",
     partition_key=
... (truncated)
```

### ✅ 6. apply_patch — 714ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+\"\"\"Perform CRUD operations against an Azure Cosmos DB for NoSQL container.\"\"\"\n+\n+import os\n+import sys\n+\n+from azure.cosmos import CosmosClient, PartitionKey\n+from azure.cosmos.exceptions import (\n+    CosmosHttpResponseError,\n+    CosmosResourceNotFoundError,\n+)\n+from azure.identity import DefaultAzureCredential\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY = \"electronics\"\n+ITEM_ID = \"item-001\"\n+\n+\n+def run_crud_operations() -\u003e None:\n+    endpoint = os.environ.get(\"COSMOS_ENDPOINT\")\n+    if not endpoint:\n+        raise ValueError(\"Set the COSMOS_ENDPOINT environment variable.\")\n+\n+    credential = DefaultAzureCredential()\n+\n+    with CosmosClient(url=endpoint, credential=credential) as client:\n+        database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+        container = database.create_container_if_not_exists(\n+            id=CONTAINER_NAME,\n+            partition_key=PartitionKey(path=\"/category\"),\n+        )\n+\n+        item = {\n+            \"id\": ITEM_ID,\n+            \"category\": PARTITION_KEY,\n+            \"name\": \"Laptop\",\n+            \"quantity\": 5,\n+        }\n+        upserted_item = container.upsert_item(body=item)\n+        print(f\"Upserted item: {upserted_item}\")\n+\n+        read_item = container.read_item(\n+            item=ITEM_ID,\n+            partition_key=PARTITION_KEY,\n+        )\n+        print(f\"Read item: {read_item}\")\n+\n+        query = \"SELECT * FROM c WHERE c.category = @category\"\n+        query_results = list(\n+            container.query_items(\n+                query=query,\n+                parameters=[{\"name\": \"@category\", \"value\": PARTITION_KEY}],\n+                partition_key=PARTITION_KEY,\n+            )\n+        )\n+        print(f\"Query results: {query_results}\")\n+\n+        read_item[\"quantity\"] = 10\n+        replaced_item = container.replace_item(\n+            item=read_item[\"id\"],\n+            body=read_item,\n+        )\n+        print(f\"Replaced item: {replaced_item}\")\n+\n+        container.delete_item(\n+            item=ITEM_ID,\n+            partition_key=PARTITION_KEY,\n+        )\n+        print(f\"Deleted item: {ITEM_ID}\")\n+\n+\n+def main() -\u003e int:\n+    try:\n+        run_crud_operations()\n+    except ValueError as error:\n+        print(f\"Configuration error: {error}\", file=sys.stderr)\n+        return 2\n+    except CosmosResourceNotFoundError as error:\n+        print(f\"Cosmos DB resource was not found: {error}\", file=sys.stderr)\n+        return 1\n+    except CosmosHttpResponseError as error:\n+        print(\n+            f\"Cosmos DB request failed with status {error.status_code}: \"\n+            f\"{error.message}\",\n+            file=sys.stderr,\n+        )\n+        return 1\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.7,\u003c5\n+azure-identity\u003e=1.17,\u003c2\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1987094751/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1987094751/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1987094751/cosmos_crud.py
@@ -1,0 +1,94 @@
+"""Perform CRUD operations against an Azure Cosmos DB for NoSQL container."""
+
+import os
+import sys
+
+from azure.cosmos import CosmosClient, PartitionKey
+from azure.cosmos.exceptions import (
+    CosmosHttpResponseError,
+    CosmosResourceNotFoundError,
+)
+from azure.identity import DefaultAzureCredential
+
+DATABASE_NAME = "TestDB"
+CONTAINER_NAME = "Items"
+PARTITION_KEY = "electronics"
+ITEM_ID = "item-001"
+
+
+def run_crud_operations() -> None:
+    endpoint = os.environ.get("COSMOS_ENDPOINT")
+    if not endpoint:
+        raise ValueError("Set the COSMOS_ENDPOINT environment variable.")
+
+    credential = DefaultAzureCredential()
+
+    with CosmosClient(url=endpoint, credential=credential) as client:
+        database = client.create_database_if_not_exists(id=DATABASE_NAME)
+        container = database.create_container_if_not_exists(
+            id=CONTAINER_NAME,
+            partition_key=PartitionKey(path="/category"),
+        )
+
+        item = {
+            "id": ITEM_ID,
+            "category": PARTITION_KEY,
+            "name": "Laptop",
+            "quantity": 5,
+        }
+        upserted_item = container.upsert_item(body=item)
+        print(f"Upserted item: {upserted_item}")
+
+        read_item = container.read_item(
+            item=ITEM_ID,
+            partition_key=PARTITION_KEY,
+        )
+        print(f"Read item: {read_item}")
+
+        query = "SELECT * FROM c WHERE c.category = @category"
+ 
... (truncated)
```

### ✅ 7. powershell — 1340ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile the Cosmos CRUD script","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace and the Cosmos SDK guidance, then create the runnable script and dependency file.

The workspace is empty, so I’ll add a focused `cosmos_crud.py` plus `requirements.txt`, using environment-based Entra authentication and single-partition operations.

Created:

- `cosmos_crud.py` — creates the client, database/container, and performs upsert, point read, parameterized query, replace, and delete operations with Cosmos SDK exception handling.
- `requirements.txt` — includes `azure-cosmos` and `azure-identity`.

The script authenticates with `DefaultAzureCredential` and reads the account endpoint from `COSMOS_ENDPOINT`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Azure Cosmos DB HTTP exceptions and a relevant resource-not-found subclass are handled explicitly and appropriately.

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: any tool (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 14.3% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 5.8333** | **72.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
