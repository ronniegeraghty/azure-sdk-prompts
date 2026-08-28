# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 356.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/13 |
| Duration | 356.9s |
| Timestamp | 2026-08-27T19:16:04Z |
| Files Generated | 2 |
| Event Count | 2041 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 265.9s |
| Review | 89.7s |
| **Total** | **356.9s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=225232 out=2592 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3242ms

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

### ❌ 2. azure-get_azure_bestpractices (via azure) — 189785ms

**Input:**

```json
{"intent":"Generate a local-only Python script using the Azure Cosmos DB for NoSQL SDK to create a client, database, container, and perform CRUD operations with exception handling. Do not deploy or modify Azure resources during validation.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 189788ms

**Input:**

```json
{"intent":"Find current official Python Azure Cosmos DB for NoSQL SDK guidance for CosmosClient, create_database_if_not_exists, create_container_if_not_exists with PartitionKey, upsert_item, read_item, parameterized query_items, replace_item, delete_item, and azure.cosmos.exceptions.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 189780ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-4103934252","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 188669ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. apply_patch — 2824ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+import os\n+import sys\n+\n+from azure.cosmos import CosmosClient, PartitionKey\n+from azure.cosmos.exceptions import CosmosHttpResponseError\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+ITEM_ID = \"item-001\"\n+CATEGORY = \"electronics\"\n+\n+\n+def main() -\u003e int:\n+    endpoint = os.environ.get(\"COSMOS_ENDPOINT\", \"https://localhost:8081/\")\n+    key = os.environ.get(\"COSMOS_KEY\")\n+    if not key:\n+        print(\n+            \"COSMOS_KEY is required. Set it to the local Cosmos DB emulator key.\",\n+            file=sys.stderr,\n+        )\n+        return 2\n+\n+    client = CosmosClient(endpoint, credential=key)\n+\n+    try:\n+        database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+        container = database.create_container_if_not_exists(\n+            id=CONTAINER_NAME,\n+            partition_key=PartitionKey(path=\"/category\"),\n+        )\n+\n+        item = {\n+            \"id\": ITEM_ID,\n+            \"category\": CATEGORY,\n+            \"name\": \"Wireless Headphones\",\n+            \"quantity\": 10,\n+        }\n+        upserted_item = container.upsert_item(body=item)\n+        print(f\"Upserted: {upserted_item}\")\n+\n+        read_item = container.read_item(\n+            item=ITEM_ID,\n+            partition_key=CATEGORY,\n+        )\n+        print(f\"Read: {read_item}\")\n+\n+        query = \"SELECT * FROM c WHERE c.category = @category\"\n+        parameters = [{\"name\": \"@category\", \"value\": CATEGORY}]\n+        queried_items = list(\n+            container.query_items(\n+                query=query,\n+                parameters=parameters,\n+                partition_key=CATEGORY,\n+            )\n+        )\n+        print(f\"Query results: {queried_items}\")\n+\n+        read_item[\"quantity\"] = 25\n+        replaced_item = container.replace_item(\n+            item=ITEM_ID,\n+            body=read_item,\n+        )\n+        print(f\"Replaced: {replaced_item}\")\n+\n+        container.delete_item(\n+            item=ITEM_ID,\n+            partition_key=CATEGORY,\n+        )\n+        print(f\"Deleted item: {ITEM_ID}\")\n+    except CosmosHttpResponseError as exc:\n+        print(\n+            f\"Cosmos DB request failed ({exc.status_code}): {exc.message}\",\n+            file=sys.stderr,\n+        )\n+        return 1\n+    finally:\n+        client.close()\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.7,\u003c5\n*** Add File: .env.example\n+COSMOS_ENDPOINT=https://localhost:8081/\n+COSMOS_KEY=replace-with-your-local-cosmos-db-emulator-key\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-4103934252/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-4103934252/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-4103934252/cosmos_crud.py
@@ -1,0 +1,85 @@
+import os
+import sys
+
+from azure.cosmos import CosmosClient, PartitionKey
+from azure.cosmos.exceptions import CosmosHttpResponseError
+
+
+DATABASE_NAME = "TestDB"
+CONTAINER_NAME = "Items"
+ITEM_ID = "item-001"
+CATEGORY = "electronics"
+
+
+def main() -> int:
+    endpoint = os.environ.get("COSMOS_ENDPOINT", "https://localhost:8081/")
+    key = os.environ.get("COSMOS_KEY")
+    if not key:
+        print(
+            "COSMOS_KEY is required. Set it to the local Cosmos DB emulator key.",
+            file=sys.stderr,
+        )
+        return 2
+
+    client = CosmosClient(endpoint, credential=key)
+
+    try:
+        database = client.create_database_if_not_exists(id=DATABASE_NAME)
+        container = database.create_container_if_not_exists(
+            id=CONTAINER_NAME,
+            partition_key=PartitionKey(path="/category"),
+        )
+
+        item = {
+            "id": ITEM_ID,
+            "category": CATEGORY,
+            "name": "Wireless Headphones",
+            "quantity": 10,
+        }
+        upserted_item = container.upsert_item(body=item)
+        print(f"Upserted: {upserted_item}")
+
+        read_item = container.read_item(
+            item=ITEM_ID,
+            partition_key=CATEGORY,
+        )
+        print(f"Read: {read_item}")
+
+        query = "SELECT * FROM c WHERE c.category = @category"
+        parameters = [{"name": "@category", "value": CATEGORY}]
+        queried_items = list(
+            container.query_items(
+                query=qu
... (truncated)
```

### ✅ 7. powershell — 3140ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile the Cosmos CRUD script","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 8. powershell — 1318ms

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -Recurse -Force .\\__pycache__ }","description":"Remove Python bytecode cache","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’m grounding the script in current Azure SDK guidance, then I’ll add the runnable files and local validation.

The SDK references do not include the requested acceptance/example files, and the documentation service timed out. I’ll use the established `azure-cosmos` v4 API surface with emulator-friendly environment configuration and no live execution.

Created:

- `cosmos_crud.py` — Cosmos client setup and complete upsert, read, parameterized query, replace, and delete flow with `CosmosHttpResponseError` handling.
- `requirements.txt` — requires `azure-cosmos>=4.7,<5`.
- `.env.example` — local Cosmos DB emulator endpoint and key placeholders.

The script defaults to `https://localhost:8081/` and requires `COSMOS_KEY`, avoiding embedded credentials or live Azure operations.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. Azure Cosmos DB request exceptions are handled appropriately.

### Strengths

- Catches the Azure Cosmos SDK-specific CosmosHttpResponseError around database, container, and item operations.
- Uses a finally block to close the CosmosClient even when an operation fails.

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
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
