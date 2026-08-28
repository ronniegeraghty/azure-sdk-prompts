# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 146.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 8/13 |
| Duration | 146.9s |
| Timestamp | 2026-08-28T22:57:46Z |
| Files Generated | 2 |
| Event Count | 1790 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 54.1s |
| Review | 91.7s |
| **Total** | **146.9s** |

## Configuration

- **name:** python-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=81719 out=1299 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. glob — 5609ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-1702304836","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 4375ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-1702304836","pattern":"**/.*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. apply_patch — 3083ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+import os\n+import sys\n+\n+from azure.cosmos import CosmosClient, PartitionKey, exceptions\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+\n+\n+def required_environment_variable(name: str) -\u003e str:\n+    value = os.environ.get(name)\n+    if not value:\n+        raise ValueError(f\"Required environment variable {name} is not set.\")\n+    return value\n+\n+\n+def main() -\u003e None:\n+    endpoint = required_environment_variable(\"COSMOS_ENDPOINT\")\n+    key = required_environment_variable(\"COSMOS_KEY\")\n+\n+    client = CosmosClient(endpoint, credential=key)\n+    database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+    container = database.create_container_if_not_exists(\n+        id=CONTAINER_NAME,\n+        partition_key=PartitionKey(path=\"/category\"),\n+    )\n+\n+    item = {\n+        \"id\": \"item-001\",\n+        \"category\": \"electronics\",\n+        \"name\": \"Wireless Mouse\",\n+        \"quantity\": 10,\n+    }\n+\n+    upserted_item = container.upsert_item(item)\n+    print(f\"Upserted item: {upserted_item}\")\n+\n+    read_item = container.read_item(\n+        item=item[\"id\"],\n+        partition_key=item[\"category\"],\n+    )\n+    print(f\"Read item: {read_item}\")\n+\n+    query = \"SELECT * FROM items i WHERE i.category = @category\"\n+    parameters = [{\"name\": \"@category\", \"value\": \"electronics\"}]\n+    queried_items = list(\n+        container.query_items(\n+            query=query,\n+            parameters=parameters,\n+            partition_key=\"electronics\",\n+        )\n+    )\n+    print(f\"Queried items: {queried_items}\")\n+\n+    updated_item = dict(read_item)\n+    updated_item[\"quantity\"] = 20\n+    replaced_item = container.replace_item(\n+        item=updated_item[\"id\"],\n+        body=updated_item,\n+    )\n+    print(f\"Replaced item: {replaced_item}\")\n+\n+    container.delete_item(\n+        item=updated_item[\"id\"],\n+        partition_key=updated_item[\"category\"],\n+    )\n+    print(f\"Deleted item: {updated_item['id']}\")\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        main()\n+    except exceptions.CosmosResourceNotFoundError as error:\n+        print(f\"Cosmos DB resource was not found: {error}\", file=sys.stderr)\n+        sys.exit(1)\n+    except exceptions.CosmosHttpResponseError as error:\n+        print(f\"Cosmos DB request failed: {error}\", file=sys.stderr)\n+        sys.exit(1)\n+    except ValueError as error:\n+        print(error, file=sys.stderr)\n+        sys.exit(1)\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.7.0,\u003c5.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-1702304836/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-1702304836/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-1702304836/cosmos_crud.py
@@ -1,0 +1,83 @@
+import os
+import sys
+
+from azure.cosmos import CosmosClient, PartitionKey, exceptions
+
+
+DATABASE_NAME = "TestDB"
+CONTAINER_NAME = "Items"
+
+
+def required_environment_variable(name: str) -> str:
+    value = os.environ.get(name)
+    if not value:
+        raise ValueError(f"Required environment variable {name} is not set.")
+    return value
+
+
+def main() -> None:
+    endpoint = required_environment_variable("COSMOS_ENDPOINT")
+    key = required_environment_variable("COSMOS_KEY")
+
+    client = CosmosClient(endpoint, credential=key)
+    database = client.create_database_if_not_exists(id=DATABASE_NAME)
+    container = database.create_container_if_not_exists(
+        id=CONTAINER_NAME,
+        partition_key=PartitionKey(path="/category"),
+    )
+
+    item = {
+        "id": "item-001",
+        "category": "electronics",
+        "name": "Wireless Mouse",
+        "quantity": 10,
+    }
+
+    upserted_item = container.upsert_item(item)
+    print(f"Upserted item: {upserted_item}")
+
+    read_item = container.read_item(
+        item=item["id"],
+        partition_key=item["category"],
+    )
+    print(f"Read item: {read_item}")
+
+    query = "SELECT * FROM items i WHERE i.category = @category"
+    parameters = [{"name": "@category", "value": "electronics"}]
+    queried_items = list(
+        container.query_items(
+            query=query,
+            parameters=parameters,
+            partition_key="electronics",
+        )
+    )
+    print(f"Queried items: {queried_items}")
+
+    
... (truncated)
```

### ✅ 4. powershell — 4378ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add the local-safe Cosmos DB CRUD script and dependency manifest.

The workspace is empty. I’ll create `cosmos_crud.py` using emulator/environment credentials and `requirements.txt` for the Azure SDK.

Created:

- `cosmos_crud.py` — Cosmos client setup and complete upsert, read, parameterized query, replace, and delete flow with Azure Cosmos exception handling.
- `requirements.txt` — requires `azure-cosmos>=4.7.0,<5.0.0`.

The script reads `COSMOS_ENDPOINT` and `COSMOS_KEY` environment variables, suitable for a local Cosmos DB emulator without provisioning Azure resources.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Azure Cosmos DB HTTP exceptions and relevant subclasses are handled appropriately without being silently swallowed.

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
      - tool_used: any tool (source=mcp, server=azure): Fail

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
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
