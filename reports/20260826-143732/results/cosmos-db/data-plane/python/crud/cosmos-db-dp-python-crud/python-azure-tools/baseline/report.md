# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 127.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 8/13 |
| Duration | 127.0s |
| Timestamp | 2026-08-26T07:04:02Z |
| Files Generated | 2 |
| Event Count | 1915 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 40.4s |
| Review | 85.4s |
| **Total** | **127.0s** |

## Configuration

- **name:** python-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=81222 out=1195 |
| Turn Count | 4 |

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

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"pattern":"requirements*.txt"}
```

### 🔧 3. glob

**Input:**

```json
{"pattern":"*.py"}
```

### 🔧 4. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+\"\"\"Perform basic CRUD operations against an Azure Cosmos DB NoSQL container.\"\"\"\n+\n+import os\n+import sys\n+from typing import Any\n+\n+from azure.cosmos import CosmosClient, PartitionKey, exceptions\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+\n+\n+def require_environment_variable(name: str) -\u003e str:\n+    value = os.getenv(name)\n+    if not value:\n+        raise RuntimeError(f\"Required environment variable {name} is not set.\")\n+    return value\n+\n+\n+def main() -\u003e None:\n+    endpoint = require_environment_variable(\"COSMOS_ENDPOINT\")\n+    key = require_environment_variable(\"COSMOS_KEY\")\n+\n+    client = CosmosClient(endpoint, credential=key)\n+    database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+    container = database.create_container_if_not_exists(\n+        id=CONTAINER_NAME,\n+        partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+    )\n+\n+    item: dict[str, Any] = {\n+        \"id\": \"item-001\",\n+        \"category\": \"electronics\",\n+        \"name\": \"Wireless Keyboard\",\n+        \"quantity\": 10,\n+    }\n+\n+    upserted_item = container.upsert_item(item)\n+    print(f\"Upserted item: {upserted_item}\")\n+\n+    read_item = container.read_item(\n+        item=item[\"id\"],\n+        partition_key=item[\"category\"],\n+    )\n+    print(f\"Read item: {read_item}\")\n+\n+    query = \"SELECT * FROM items i WHERE i.category = @category\"\n+    parameters = [{\"name\": \"@category\", \"value\": \"electronics\"}]\n+    queried_items = list(\n+        container.query_items(\n+            query=query,\n+            parameters=parameters,\n+            partition_key=\"electronics\",\n+        )\n+    )\n+    print(f\"Queried items: {queried_items}\")\n+\n+    read_item[\"quantity\"] = 15\n+    replaced_item = container.replace_item(\n+        item=read_item[\"id\"],\n+        body=read_item,\n+    )\n+    print(f\"Replaced item: {replaced_item}\")\n+\n+    container.delete_item(\n+        item=read_item[\"id\"],\n+        partition_key=read_item[\"category\"],\n+    )\n+    print(f\"Deleted item: {read_item['id']}\")\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        main()\n+    except RuntimeError as error:\n+        print(f\"Configuration error: {error}\", file=sys.stderr)\n+        sys.exit(1)\n+    except exceptions.CosmosResourceNotFoundError as error:\n+        print(f\"Cosmos DB resource not found: {error}\", file=sys.stderr)\n+        sys.exit(1)\n+    except exceptions.CosmosHttpResponseError as error:\n+        print(\n+            f\"Cosmos DB request failed \"\n+            f\"(status={error.status_code}, message={error.message})\",\n+            file=sys.stderr,\n+        )\n+        sys.exit(1)\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.7.0,\u003c5.0.0\n*** End Patch\n"
```

### 🔧 5. powershell

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile the Cosmos CRUD script","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add a locally configurable Cosmos DB CRUD script and its required package manifest.

The workspace is empty. I’ll create `cosmos_crud.py` using environment-based credentials and `requirements.txt` with the Azure Cosmos SDK dependency.

The implementation is in place; I’ll run a syntax compilation check without connecting to Azure or creating resources.

Created:

- `cosmos_crud.py` — creates the client, database, and container; performs upsert, point read, parameterized query, replace, and delete; handles Cosmos SDK exceptions.
- `requirements.txt` — includes `azure-cosmos>=4.7.0,<5.0.0`.

The script reads credentials from `COSMOS_ENDPOINT` and `COSMOS_KEY` environment variables.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Azure Cosmos DB HTTP exceptions are handled with appropriate specificity and ordering.

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 21.7% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 26.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 26.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 26.1% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 3.8333** | **47.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
