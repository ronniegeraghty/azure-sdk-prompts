# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-cosmos-skill/without-azure-sdk-python | **Result:** ❌ FAILED | **Duration:** 230.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-cosmos-skill/without-azure-sdk-python |
| Result | ❌ FAILED |
| Score | 9/13 |
| Duration | 230.7s |
| Timestamp | 2026-08-26T02:24:27Z |
| Files Generated | 2 |
| Event Count | 1952 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 139.9s |
| Review | 89.6s |
| **Total** | **230.7s** |

## Configuration

- **name:** python-cosmos-skill/without-azure-sdk-python
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=200425 out=1753 |
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

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local Python script demonstrating Azure Cosmos DB for NoSQL CRUD operations, using environment-variable configuration and azure.cosmos exception handling. Do not provision or modify Azure resources from this environment; only write runnable sample code.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a Python Azure Cosmos DB for NoSQL CRUD sample that uses environment-variable credentials and is not executed against live Azure.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 4. glob

**Input:**

```json
{"pattern":"**/.*"}
```

### 🔧 5. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+import logging\n+import os\n+\n+from azure.cosmos import CosmosClient, PartitionKey\n+from azure.cosmos.exceptions import (\n+    CosmosHttpResponseError,\n+    CosmosResourceNotFoundError,\n+)\n+from azure.identity import DefaultAzureCredential\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+\n+\n+def main() -\u003e None:\n+    endpoint = os.environ.get(\"AZURE_COSMOS_ENDPOINT\")\n+    if not endpoint:\n+        raise RuntimeError(\"Set AZURE_COSMOS_ENDPOINT to the Cosmos DB account endpoint.\")\n+\n+    credential = DefaultAzureCredential()\n+    client = CosmosClient(url=endpoint, credential=credential)\n+\n+    item = {\n+        \"id\": \"item-001\",\n+        \"category\": \"electronics\",\n+        \"name\": \"Wireless keyboard\",\n+        \"quantity\": 10,\n+    }\n+\n+    try:\n+        database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+        container = database.create_container_if_not_exists(\n+            id=CONTAINER_NAME,\n+            partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+        )\n+\n+        container.upsert_item(body=item)\n+        logging.info(\"Upserted item %s.\", item[\"id\"])\n+\n+        stored_item = container.read_item(\n+            item=item[\"id\"],\n+            partition_key=item[\"category\"],\n+        )\n+        logging.info(\"Read item: %s\", stored_item)\n+\n+        query = \"SELECT * FROM items i WHERE i.category = @category\"\n+        parameters = [{\"name\": \"@category\", \"value\": \"electronics\"}]\n+        matching_items = list(\n+            container.query_items(\n+                query=query,\n+                parameters=parameters,\n+                partition_key=\"electronics\",\n+            )\n+        )\n+        logging.info(\"Query returned %d item(s): %s\", len(matching_items), matching_items)\n+\n+        stored_item[\"quantity\"] = 25\n+        replaced_item = container.replace_item(\n+            item=stored_item[\"id\"],\n+            body=stored_item,\n+        )\n+        logging.info(\"Replaced item with quantity %s.\", replaced_item[\"quantity\"])\n+\n+        container.delete_item(\n+            item=replaced_item[\"id\"],\n+            partition_key=replaced_item[\"category\"],\n+        )\n+        logging.info(\"Deleted item %s.\", replaced_item[\"id\"])\n+    except CosmosResourceNotFoundError as error:\n+        logging.error(\"A Cosmos DB resource was not found: %s\", error)\n+        raise\n+    except CosmosHttpResponseError as error:\n+        logging.error(\"Cosmos DB request failed: %s\", error)\n+        raise\n+    finally:\n+        client.close()\n+        credential.close()\n+\n+\n+if __name__ == \"__main__\":\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+    main()\n*** Add File: requirements.txt\n+azure-cosmos\n+azure-identity\n*** End Patch\n"
```

### 🔧 6. powershell

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add the smallest runnable sample and dependency file without connecting to Azure.

Created:

- `cosmos_crud.py` — CosmosClient setup, database/container creation, upsert, point read, parameterized query, replace, delete, logging, and `azure.cosmos.exceptions` handling.
- `requirements.txt` — includes `azure-cosmos` and `azure-identity`.

The script uses `DefaultAzureCredential`; set `AZURE_COSMOS_ENDPOINT` before running.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses specific Azure Cosmos DB exception types, orders the subclass handler before the broader handler, and preserves failure propagation.

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
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-cosmos-skill/without-azure-sdk-python
```

---

[← Back to Summary](../../../../../../summary.md)
