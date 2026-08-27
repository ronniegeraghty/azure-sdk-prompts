# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 162.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 8/13 |
| Duration | 162.8s |
| Timestamp | 2026-08-27T02:51:09Z |
| Files Generated | 2 |
| Event Count | 2539 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 69.0s |
| Review | 92.7s |
| **Total** | **162.8s** |

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
| Token Usage | in=126080 out=2124 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. glob — 5258ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 4084ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109","pattern":"**/.*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. apply_patch — 2185ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+import os\n+import sys\n+\n+from azure.cosmos import CosmosClient, PartitionKey, exceptions\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+\n+\n+def main() -\u003e None:\n+    endpoint = os.environ[\"COSMOS_ENDPOINT\"]\n+    key = os.environ[\"COSMOS_KEY\"]\n+\n+    client = CosmosClient(endpoint, credential=key)\n+    database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+    container = database.create_container_if_not_exists(\n+        id=CONTAINER_NAME,\n+        partition_key=PartitionKey(path=\"/category\"),\n+    )\n+\n+    item = {\n+        \"id\": \"item-001\",\n+        \"category\": \"electronics\",\n+        \"name\": \"Wireless keyboard\",\n+        \"quantity\": 10,\n+    }\n+\n+    try:\n+        upserted_item = container.upsert_item(item)\n+        print(f\"Upserted item: {upserted_item}\")\n+\n+        read_item = container.read_item(\n+            item=item[\"id\"],\n+            partition_key=item[\"category\"],\n+        )\n+        print(f\"Read item: {read_item}\")\n+\n+        query = \"SELECT * FROM items i WHERE i.category = @category\"\n+        parameters = [\n+            {\"name\": \"@category\", \"value\": \"electronics\"},\n+        ]\n+        queried_items = list(\n+            container.query_items(\n+                query=query,\n+                parameters=parameters,\n+                partition_key=\"electronics\",\n+            )\n+        )\n+        print(f\"Queried items: {queried_items}\")\n+\n+        read_item[\"quantity\"] = 25\n+        replaced_item = container.replace_item(\n+            item=read_item[\"id\"],\n+            body=read_item,\n+        )\n+        print(f\"Replaced item: {replaced_item}\")\n+\n+        container.delete_item(\n+            item=read_item[\"id\"],\n+            partition_key=read_item[\"category\"],\n+        )\n+        print(f\"Deleted item: {read_item['id']}\")\n+    except exceptions.CosmosResourceNotFoundError as error:\n+        print(f\"Cosmos DB resource was not found: {error}\", file=sys.stderr)\n+        raise\n+    except exceptions.CosmosAccessConditionFailedError as error:\n+        print(f\"Cosmos DB access condition failed: {error}\", file=sys.stderr)\n+        raise\n+    except exceptions.CosmosHttpResponseError as error:\n+        print(\n+            f\"Cosmos DB request failed with status {error.status_code}: {error}\",\n+            file=sys.stderr,\n+        )\n+        raise\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        main()\n+    except KeyError as error:\n+        variable = error.args[0]\n+        print(\n+            f\"Missing required environment variable: {variable}\",\n+            file=sys.stderr,\n+        )\n+        sys.exit(2)\n+    except exceptions.CosmosHttpResponseError:\n+        sys.exit(1)\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.7,\u003c5\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
@@ -1,0 +1,90 @@
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
+def main() -> None:
+    endpoint = os.environ["COSMOS_ENDPOINT"]
+    key = os.environ["COSMOS_KEY"]
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
+        "name": "Wireless keyboard",
+        "quantity": 10,
+    }
+
+    try:
+        upserted_item = container.upsert_item(item)
+        print(f"Upserted item: {upserted_item}")
+
+        read_item = container.read_item(
+            item=item["id"],
+            partition_key=item["category"],
+        )
+        print(f"Read item: {read_item}")
+
+        query = "SELECT * FROM items i WHERE i.category = @category"
+        parameters = [
+            {"name": "@category", "value": "electronics"},
+        ]
+        queried_items = list(
+            container.query_items(
+                query=query,
+                parameters=parameters,
+                partition_key="electronics",
+            )
+        )
+        print(f"Queried items: {queried_items}")
+
+        read_item["quantity"] = 25
+        replaced_item = container.replace_item(
+            item=read_item["id"],
+            body=read_item,
... (truncated)
```

### ✅ 4. powershell — 2637ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 5. apply_patch — 2740ms

**Input:**

```json
"*** Begin Patch\n*** Update File: cosmos_crud.py\n@@\n def main() -\u003e None:\n     endpoint = os.environ[\"COSMOS_ENDPOINT\"]\n     key = os.environ[\"COSMOS_KEY\"]\n \n-    client = CosmosClient(endpoint, credential=key)\n-    database = client.create_database_if_not_exists(id=DATABASE_NAME)\n-    container = database.create_container_if_not_exists(\n-        id=CONTAINER_NAME,\n-        partition_key=PartitionKey(path=\"/category\"),\n-    )\n-\n-    item = {\n-        \"id\": \"item-001\",\n-        \"category\": \"electronics\",\n-        \"name\": \"Wireless keyboard\",\n-        \"quantity\": 10,\n-    }\n-\n     try:\n+        client = CosmosClient(endpoint, credential=key)\n+        database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+        container = database.create_container_if_not_exists(\n+            id=CONTAINER_NAME,\n+            partition_key=PartitionKey(path=\"/category\"),\n+        )\n+\n+        item = {\n+            \"id\": \"item-001\",\n+            \"category\": \"electronics\",\n+            \"name\": \"Wireless keyboard\",\n+            \"quantity\": 10,\n+        }\n+\n         upserted_item = container.upsert_item(item)\n         print(f\"Upserted item: {upserted_item}\")\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
@@ -12,21 +12,21 @@
     endpoint = os.environ["COSMOS_ENDPOINT"]
     key = os.environ["COSMOS_KEY"]
 
-    client = CosmosClient(endpoint, credential=key)
-    database = client.create_database_if_not_exists(id=DATABASE_NAME)
-    container = database.create_container_if_not_exists(
-        id=CONTAINER_NAME,
-        partition_key=PartitionKey(path="/category"),
-    )
+    try:
+        client = CosmosClient(endpoint, credential=key)
+        database = client.create_database_if_not_exists(id=DATABASE_NAME)
+        container = database.create_container_if_not_exists(
+            id=CONTAINER_NAME,
+            partition_key=PartitionKey(path="/category"),
+        )
 
-    item = {
-        "id": "item-001",
-        "category": "electronics",
-        "name": "Wireless keyboard",
-        "quantity": 10,
-    }
+        item = {
+            "id": "item-001",
+            "category": "electronics",
+            "name": "Wireless keyboard",
+            "quantity": 10,
+        }
 
-    try:
         upserted_item = container.upsert_item(item)
         print(f"Upserted item: {upserted_item}")
 

```

### ✅ 6. view — 6114ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109\\cosmos_crud.py"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/requirements.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/requirements.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/requirements.txt
@@ -1,2 +1,2 @@
 azure-cosmos>=4.7,<5
 

```

### ✅ 7. view — 4916ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109\\requirements.txt"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-baseline-690713109/cosmos_crud.py
@@ -1,90 +1,90 @@
 import os
 import sys
 
 from azure.cosmos import CosmosClient, PartitionKey, exceptions
 
 
 DATABASE_NAME = "TestDB"
 CONTAINER_NAME = "Items"
 
 
 def main() -> None:
     endpoint = os.environ["COSMOS_ENDPOINT"]
     key = os.environ["COSMOS_KEY"]
 
     try:
         client = CosmosClient(endpoint, credential=key)
         database = client.create_database_if_not_exists(id=DATABASE_NAME)
         container = database.create_container_if_not_exists(
             id=CONTAINER_NAME,
             partition_key=PartitionKey(path="/category"),
         )
 
         item = {
             "id": "item-001",
             "category": "electronics",
             "name": "Wireless keyboard",
             "quantity": 10,
         }
 
         upserted_item = container.upsert_item(item)
         print(f"Upserted item: {upserted_item}")
 
         read_item = container.read_item(
             item=item["id"],
             partition_key=item["category"],
         )
         print(f"Read item: {read_item}")
 
         query = "SELECT * FROM items i WHERE i.category = @category"
         parameters = [
             {"name": "@category", "value": "electronics"},
         ]
         queried_items = list(
             container.query_items(
                 query=query,
                 parameters=parameters,
                 partition_key="electronics",
             )
         )
         print(f"Queried items: {queried_items}")
 
     
... (truncated)
```

### ✅ 8. powershell — 7265ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Recheck Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add the local-ready Cosmos DB CRUD script and dependency manifest.

The workspace is empty. I’ll create a self-contained script using environment variables plus `requirements.txt`, with parameterized querying and Cosmos-specific error handling.

I’m tightening exception coverage so client initialization and database/container creation receive the same explicit Cosmos error reporting as item operations.

Created:

- `cosmos_crud.py` — creates the client, database, and partitioned container; performs upsert, point read, parameterized query, replace, and delete operations; handles `azure.cosmos.exceptions`.
- `requirements.txt` — specifies `azure-cosmos>=4.7,<5`.

The script reads credentials from `COSMOS_ENDPOINT` and `COSMOS_KEY` environment variables.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses exception classes from azure.cosmos.exceptions, handles specific subclasses before the base CosmosHttpResponseError, and exits with an appropriate failure status.

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
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
