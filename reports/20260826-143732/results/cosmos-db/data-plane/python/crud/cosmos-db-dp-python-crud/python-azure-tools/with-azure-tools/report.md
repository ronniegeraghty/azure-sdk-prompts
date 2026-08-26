# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 150.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 9/13 |
| Duration | 150.6s |
| Timestamp | 2026-08-26T07:06:09Z |
| Files Generated | 2 |
| Event Count | 2135 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 73.4s |
| Review | 76.0s |
| **Total** | **150.6s** |

## Configuration

- **name:** python-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=252682 out=2445 |
| Turn Count | 6 |

## Prompt Sent

```
Write a Python script that performs CRUD operations on items in an Azure Cosmos DB NoSQL container.

**Write the code to files (use file-write tools, do not reply with code blocks).**

1. Create a CosmosClient
2. Create a database "TestDB" and container "Items" with partition key "/category"
3. Upsert an item dict with keys: id, category, name, quantity
4. Read the item back using read_item() with id and partition key
5. Query items where category equals "electronics" using parameterized query
6. Replace the item with updated quantity
7. Delete the item

Show required pip packages and handle exceptions from azure.cosmos.exceptions.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-cosmos-py"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local Python script demonstrating Azure Cosmos DB NoSQL CRUD with DefaultAzureCredential, partition-scoped parameterized queries, and SDK exception handling. Do not deploy or modify Azure resources from tooling; only generate local code.","learn":true}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py","pattern":"**/{acceptance-criteria.md,examples.md,partitioning.md,query-patterns.md}"}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-tools-with-azure-tools-274352906","pattern":"**/*"}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation guidance for a local Python Azure Cosmos DB NoSQL CRUD sample using the Azure SDK. No live deployment.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py\\references\\partitioning.md"}
```

### 🔧 7. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py\\references\\query-patterns.md"}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace files"}
```

### 🔧 9. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+import logging\n+import os\n+import sys\n+from typing import Any\n+\n+from azure.cosmos import CosmosClient, PartitionKey\n+from azure.cosmos.exceptions import CosmosHttpResponseError\n+from azure.identity import DefaultAzureCredential\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+ITEM_ID = \"item-001\"\n+ITEM_CATEGORY = \"electronics\"\n+\n+logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+logger = logging.getLogger(__name__)\n+\n+\n+def run_crud_operations() -\u003e None:\n+    endpoint = os.environ.get(\"COSMOS_ENDPOINT\")\n+    if not endpoint:\n+        raise ValueError(\n+            \"COSMOS_ENDPOINT must be set, for example \"\n+            \"'https://\u003caccount\u003e.documents.azure.com:443/'.\"\n+        )\n+\n+    credential = DefaultAzureCredential()\n+    item: dict[str, Any] = {\n+        \"id\": ITEM_ID,\n+        \"category\": ITEM_CATEGORY,\n+        \"name\": \"Laptop\",\n+        \"quantity\": 10,\n+    }\n+\n+    try:\n+        with CosmosClient(url=endpoint, credential=credential) as client:\n+            database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+            container = database.create_container_if_not_exists(\n+                id=CONTAINER_NAME,\n+                partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+            )\n+\n+            upserted_item = container.upsert_item(body=item)\n+            logger.info(\"Upserted item %s\", upserted_item[\"id\"])\n+\n+            read_item = container.read_item(\n+                item=ITEM_ID,\n+                partition_key=ITEM_CATEGORY,\n+            )\n+            logger.info(\"Read item: %s\", read_item)\n+\n+            query = \"SELECT * FROM c WHERE c.category = @category\"\n+            query_results = list(\n+                container.query_items(\n+                    query=query,\n+                    parameters=[\n+                        {\"name\": \"@category\", \"value\": ITEM_CATEGORY},\n+                    ],\n+                    partition_key=ITEM_CATEGORY,\n+                )\n+            )\n+            logger.info(\n+                \"Found %d item(s) in category %s\",\n+                len(query_results),\n+                ITEM_CATEGORY,\n+            )\n+\n+            read_item[\"quantity\"] = 20\n+            replaced_item = container.replace_item(\n+                item=read_item[\"id\"],\n+                body=read_item,\n+            )\n+            logger.info(\n+                \"Replaced item %s; quantity is now %s\",\n+                replaced_item[\"id\"],\n+                replaced_item[\"quantity\"],\n+            )\n+\n+            container.delete_item(\n+                item=ITEM_ID,\n+                partition_key=ITEM_CATEGORY,\n+            )\n+            logger.info(\"Deleted item %s\", ITEM_ID)\n+    except CosmosHttpResponseError as error:\n+        logger.error(\n+            \"Cosmos DB request failed (status %s): %s\",\n+            error.status_code,\n+            error.message,\n+        )\n+        raise\n+    finally:\n+        credential.close()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        run_crud_operations()\n+    except (CosmosHttpResponseError, ValueError):\n+        sys.exit(1)\n*** Add File: requirements.txt\n+azure-cosmos\n+azure-identity\n*** End Patch\n"
```

### 🔧 10. powershell

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile Cosmos CRUD script"}
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’m grounding the script in the Cosmos SDK’s current authentication, partitioning, and exception-handling patterns before writing the files.

The workspace is empty, so I’m creating a self-contained script plus a minimal dependency manifest. The script will use Entra authentication, a partition-scoped parameterized query, and explicit Cosmos SDK error handling.

Created:

- `cosmos_crud.py` — creates `CosmosClient`, database/container, and performs upsert, point read, parameterized query, replace, and delete with `CosmosHttpResponseError` handling.
- `requirements.txt` — includes `azure-cosmos` and `azure-identity`.

Set `COSMOS_ENDPOINT`, authenticate through `DefaultAzureCredential`, then run `pip install -r requirements.txt` followed by `python cosmos_crud.py`.

Reference: [Azure Cosmos DB Python SDK samples](https://learn.microsoft.com/azure/cosmos-db/nosql/how-to-python).

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Uses the SDK-specific CosmosHttpResponseError and avoids silently swallowing request failures.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/6)
      - `azure-cosmos` pip package: Pass
      - `CosmosClient` creation: Pass
      - `database_client.create_database_if_not_exists()`: Fail
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
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 67% | 1.00 | 0.6667 | 11.8% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.6% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.6% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.6% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.6% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 17.6% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 5.6667** | **70.8%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
