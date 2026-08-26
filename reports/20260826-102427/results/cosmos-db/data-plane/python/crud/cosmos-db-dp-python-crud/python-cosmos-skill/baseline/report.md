# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-cosmos-skill/baseline | **Result:** ❌ FAILED | **Duration:** 243.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-cosmos-skill/baseline |
| Result | ❌ FAILED |
| Score | 10/13 |
| Duration | 243.5s |
| Timestamp | 2026-08-26T02:28:18Z |
| Files Generated | 2 |
| Event Count | 2325 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 165.6s |
| Review | 76.2s |
| **Total** | **243.5s** |

## Configuration

- **name:** python-cosmos-skill/baseline
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
| Token Usage | in=339908 out=2562 |
| Turn Count | 8 |

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
{"intent":"Create a local-only Python script using Azure Cosmos DB SDK for NoSQL CRUD operations, DefaultAzureCredential, partition key /category, parameterized queries, and exception handling.","learn":true}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Cosmos DB Python code-generation best practices for a local-only CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-cosmos-skill-baseline-3205403014","pattern":"**/*"}
```

### 🔧 5. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-py","pattern":"**/{acceptance-criteria.md,examples.md,partitioning.md,query-patterns.md}"}
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

### 🔧 8. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+\"\"\"Run a complete CRUD workflow against an Azure Cosmos DB NoSQL container.\"\"\"\n+\n+import logging\n+import os\n+\n+from azure.cosmos import CosmosClient, PartitionKey\n+from azure.cosmos.exceptions import (\n+    CosmosHttpResponseError,\n+    CosmosResourceNotFoundError,\n+)\n+from azure.identity import DefaultAzureCredential\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+\n+\n+def run_crud_workflow() -\u003e None:\n+    endpoint = os.getenv(\"COSMOS_ENDPOINT\")\n+    if not endpoint:\n+        raise RuntimeError(\"Set COSMOS_ENDPOINT to the Azure Cosmos DB account endpoint.\")\n+\n+    credential = DefaultAzureCredential()\n+    try:\n+        with CosmosClient(url=endpoint, credential=credential) as client:\n+            database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+            container = database.create_container_if_not_exists(\n+                id=CONTAINER_NAME,\n+                partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+            )\n+\n+            item = {\n+                \"id\": \"item-001\",\n+                \"category\": \"electronics\",\n+                \"name\": \"Wireless Keyboard\",\n+                \"quantity\": 10,\n+            }\n+\n+            upserted_item = container.upsert_item(body=item)\n+            logging.info(\"Upserted item %s.\", upserted_item[\"id\"])\n+\n+            read_item = container.read_item(\n+                item=item[\"id\"],\n+                partition_key=item[\"category\"],\n+            )\n+            logging.info(\"Read item: %s\", read_item)\n+\n+            query = \"SELECT * FROM c WHERE c.category = @category\"\n+            queried_items = list(\n+                container.query_items(\n+                    query=query,\n+                    parameters=[\n+                        {\"name\": \"@category\", \"value\": \"electronics\"},\n+                    ],\n+                    partition_key=\"electronics\",\n+                )\n+            )\n+            logging.info(\"Query returned %d item(s): %s\", len(queried_items), queried_items)\n+\n+            read_item[\"quantity\"] = 25\n+            replaced_item = container.replace_item(\n+                item=read_item[\"id\"],\n+                body=read_item,\n+            )\n+            logging.info(\n+                \"Replaced item %s with quantity %d.\",\n+                replaced_item[\"id\"],\n+                replaced_item[\"quantity\"],\n+            )\n+\n+            container.delete_item(\n+                item=replaced_item[\"id\"],\n+                partition_key=replaced_item[\"category\"],\n+            )\n+            logging.info(\"Deleted item %s.\", replaced_item[\"id\"])\n+    finally:\n+        credential.close()\n+\n+\n+def main() -\u003e int:\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+\n+    try:\n+        run_crud_workflow()\n+    except CosmosResourceNotFoundError as error:\n+        logging.error(\"Cosmos DB resource was not found: %s\", error)\n+        return 1\n+    except CosmosHttpResponseError as error:\n+        logging.error(\n+            \"Cosmos DB request failed (status %s): %s\",\n+            error.status_code,\n+            error.message,\n+        )\n+        return 1\n+    except RuntimeError as error:\n+        logging.error(\"%s\", error)\n+        return 1\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-cosmos\n+azure-identity\n*** End Patch\n"
```

### 🔧 9. powershell

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile-check the Cosmos DB script","initial_wait":30,"mode":"sync"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"if (Test-Path -LiteralPath '.\\__pycache__') { Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force }","description":"Remove Python compile cache","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll create a runnable script plus a minimal dependency manifest, using Entra authentication and single-partition operations throughout.

Created:

- `cosmos_crud.py` — creates the client, database, and container; performs upsert, point read, parameterized partition query, replace, and delete; handles `azure.cosmos.exceptions`.
- `requirements.txt` — includes `azure-cosmos` and `azure-identity`.

Set `COSMOS_ENDPOINT`, authenticate through `DefaultAzureCredential`, then run:

`pip install -r requirements.txt`  
`python cosmos_crud.py`

Reference: [Azure Cosmos DB Python SDK documentation](https://learn.microsoft.com/azure/cosmos-db/nosql/sdk-python)

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Appropriately handles both a specific Cosmos DB not-found exception and the broader HTTP response exception hierarchy.

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
      - tool_used: azure (source=mcp, server=azure): Fail

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
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-cosmos-skill/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
