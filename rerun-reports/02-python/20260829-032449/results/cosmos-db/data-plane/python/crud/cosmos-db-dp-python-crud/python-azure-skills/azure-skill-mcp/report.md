# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 226.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 9/13 |
| Duration | 226.3s |
| Timestamp | 2026-08-28T23:00:13Z |
| Files Generated | 2 |
| Event Count | 2023 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 126.3s |
| Review | 98.9s |
| **Total** | **226.3s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=232195 out=2815 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. azure-get_azure_bestpractices (via azure) — 6932ms

**Input:**

```json
{"command":"get_best_practices","intent":"Create a local Python Azure Cosmos DB NoSQL CRUD sample using CosmosClient, parameterized queries, and azure.cosmos exception handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 2. azure-documentation (via azure) — 9370ms

**Input:**

```json
{"command":"search","intent":"Find current official Python Azure Cosmos DB NoSQL SDK examples for CosmosClient, create_database_if_not_exists, create_container_if_not_exists, upsert_item, read_item, parameterized query_items, replace_item, delete_item, and CosmosHttpResponseError handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 3. glob — 8108ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 4793ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a Python Azure Cosmos DB NoSQL CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 5522ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Search official Cosmos DB Python SDK CRUD documentation.","parameters":{"query":"Azure Cosmos DB NoSQL Python SDK CRUD create_database_if_not_exists create_container_if_not_exists upsert read query replace delete exceptions"}}
```

**Output:**

```
{"results":[{"title":"Create a database in Azure Cosmos DB for NoSQL using Python","content":"# Create a database in Azure Cosmos DB for NoSQL using Python\n## Parsing the response\nIn the examples above, the response from the requests is a [\u0060DatabaseProxy\u0060](https://learn.microsoft.com/en-us/python/api/azure-cosmos/azure.cosmos.databaseproxy), which is an interface to interact with a specific database. From the proxy, you can access methods to perform operations on the database.\nThe following example shows the **create\\_database\\_if\\_not\\_exists** method returning a **database** object.\n\u0060\u0060\u0060python\ndatabase = client.create_database_if_not_exists(id=DATABASE_ID)\nfor container in database.list_containers():\n    print(f\u0027Container name: {container[\u0022id\u0022]}\u0027)\n\u0060\u0060\u0060","contentUrl":"https://learn.microsoft.com/azure/cosmos-db/how-to-python-create-database#parsing-the-response"},{"title":"DatabaseProxy Class-Methods","content":"# DatabaseProxy Class\r\n\r\nAn interface to interact with a specific database.\r\n\r\nThis class should not be instantiated directly. Instead use the [get_database_client](https://learn.microsoft.com/python/api/azure-cosmos/azure.cosmos.cosmosclient?view=azure-python#azure-cosmos-cosmosclient-get-database-client) method.\r\n\r\nA database contains one or more containers, each of which can contain items, stored procedures, triggers, and user-defined functions.\r\n\r\nA database can also have associated users, each of which is configured with a set of permissions for accessing certain containers, stored procedures, triggers, user-defined functions, or items.\r\n\r\nAn Azure Cosmos DB SQL API database has the following system-generated properties. These properties are read-only:\r\n\r\n- *\\_rid*: The resource ID.\r\n- *\\_ts*: When the resource was last updated. The value is a timestamp.\r\n- *\\_self*: The unique addressable URI for the resource.\r\n- *\\_etag*: The resource etag required
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 5149ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official Python code samples for Cosmos DB CRUD and parameterized queries.","parameters":{"language":"python","query":"azure.cosmos CosmosClient create_database_if_not_exists create_container_if_not_exists upsert_item read_item parameterized query_items replace_item delete_item CosmosHttpResponseError"}}
```

**Output:**

```
{"results":[{"description":"description: Creates a container in Azure Cosmos DB using the create_container_if_not_exists method. It specifies a partition key and throughput, and handles exceptions related to HTTP response errors from the Cosmos DB service.\npackage: azure-cosmos\nlanguage: python\n","codeSnippet":"try:\r\n    partition_key_path = PartitionKey(path=\u0022/categoryId\u0022)\r\n    container = database.create_container_if_not_exists(\r\n        id=CONTAINER_ID,\r\n        partition_key=partition_key_path,\r\n        offer_throughput=400,\r\n    )\r\n    print(f\u0022Container created or returned: {container.id}\u0022)\r\n\r\nexcept CosmosHttpResponseError:\r\n    print(\u0022Request to the Azure Cosmos database service failed.\u0022)","link":"https://learn.microsoft.com/azure/cosmos-db/how-to-python-create-container#create-a-container","language":"python"},{"description":"description: Creates a container in Azure Cosmos DB using the create_container_if_not_exists method. It specifies a partition key and throughput, and handles exceptions related to HTTP response errors, avoiding errors if the container already exists.\npackage: azure-cosmos\nlanguage: python\n","codeSnippet":"try:\r\n    partition_key_path = PartitionKey(path=\u0022/categoryId\u0022)\r\n    container = database.create_container_if_not_exists(\r\n        id=CONTAINER_ID,\r\n        partition_key=partition_key_path,\r\n        offer_throughput=400,\r\n    )\r\n    print(f\u0022Container created or returned: {container.id}\u0022)\r\n\r\nexcept CosmosHttpResponseError:\r\n    print(\u0022Request to the Azure Cosmos database service failed.\u0022)","link":"https://learn.microsoft.com/azure/cosmos-db/how-to-python-create-container#create-a-container","language":"python"},{"description":"description: Creates a database in Azure Cosmos DB using the create_database_if_not_exists method. This method returns the database if it exists and does not throw an exception, handling errors with a catch f
... (truncated)
```

### ✅ 7. apply_patch — 3432ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+\"\"\"Perform a complete CRUD lifecycle in an Azure Cosmos DB for NoSQL container.\"\"\"\n+\n+import os\n+import sys\n+from typing import Any\n+\n+from azure.cosmos import CosmosClient, PartitionKey, exceptions\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+\n+\n+def run_crud_operations() -\u003e None:\n+    endpoint = os.environ[\"COSMOS_ENDPOINT\"]\n+    key = os.environ[\"COSMOS_KEY\"]\n+\n+    client = CosmosClient(endpoint, credential=key)\n+    database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+    container = database.create_container_if_not_exists(\n+        id=CONTAINER_NAME,\n+        partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+    )\n+\n+    item: dict[str, Any] = {\n+        \"id\": \"item-001\",\n+        \"category\": \"electronics\",\n+        \"name\": \"Wireless Mouse\",\n+        \"quantity\": 10,\n+    }\n+\n+    upserted_item = container.upsert_item(body=item)\n+    print(f\"Upserted item: {upserted_item}\")\n+\n+    read_item = container.read_item(\n+        item=item[\"id\"],\n+        partition_key=item[\"category\"],\n+    )\n+    print(f\"Read item: {read_item}\")\n+\n+    query = \"SELECT * FROM items i WHERE i.category = @category\"\n+    parameters = [{\"name\": \"@category\", \"value\": \"electronics\"}]\n+    queried_items = list(\n+        container.query_items(\n+            query=query,\n+            parameters=parameters,\n+            partition_key=\"electronics\",\n+        )\n+    )\n+    print(f\"Queried items: {queried_items}\")\n+\n+    item[\"quantity\"] = 15\n+    replaced_item = container.replace_item(item=item[\"id\"], body=item)\n+    print(f\"Replaced item: {replaced_item}\")\n+\n+    container.delete_item(\n+        item=item[\"id\"],\n+        partition_key=item[\"category\"],\n+    )\n+    print(f\"Deleted item: {item['id']}\")\n+\n+\n+def main() -\u003e int:\n+    try:\n+        run_crud_operations()\n+    except KeyError as error:\n+        print(\n+            f\"Missing required environment variable: {error.args[0]}\",\n+            file=sys.stderr,\n+        )\n+        return 2\n+    except exceptions.CosmosHttpResponseError as error:\n+        print(\n+            f\"Cosmos DB request failed ({error.status_code}): {error.message}\",\n+            file=sys.stderr,\n+        )\n+        return 1\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-cosmos\u003e=4.9.0,\u003c5.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
@@ -1,0 +1,85 @@
+"""Perform a complete CRUD lifecycle in an Azure Cosmos DB for NoSQL container."""
+
+import os
+import sys
+from typing import Any
+
+from azure.cosmos import CosmosClient, PartitionKey, exceptions
+
+
+DATABASE_NAME = "TestDB"
+CONTAINER_NAME = "Items"
+PARTITION_KEY_PATH = "/category"
+
+
+def run_crud_operations() -> None:
+    endpoint = os.environ["COSMOS_ENDPOINT"]
+    key = os.environ["COSMOS_KEY"]
+
+    client = CosmosClient(endpoint, credential=key)
+    database = client.create_database_if_not_exists(id=DATABASE_NAME)
+    container = database.create_container_if_not_exists(
+        id=CONTAINER_NAME,
+        partition_key=PartitionKey(path=PARTITION_KEY_PATH),
+    )
+
+    item: dict[str, Any] = {
+        "id": "item-001",
+        "category": "electronics",
+        "name": "Wireless Mouse",
+        "quantity": 10,
+    }
+
+    upserted_item = container.upsert_item(body=item)
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
+    item["quantity"] = 15
+    replaced_it
... (truncated)
```

### ✅ 8. apply_patch — 3437ms

**Input:**

```json
"*** Begin Patch\n*** Update File: cosmos_crud.py\n@@\n     container = database.create_container_if_not_exists(\n         id=CONTAINER_NAME,\n         partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+        offer_throughput=400,\n     )\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
@@ -21,6 +21,7 @@
     container = database.create_container_if_not_exists(
         id=CONTAINER_NAME,
         partition_key=PartitionKey(path=PARTITION_KEY_PATH),
+        offer_throughput=400,
     )
 
     item: dict[str, Any] = {

```

### ✅ 9. powershell — 9975ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 10. view — 8538ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089\\cosmos_crud.py"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/requirements.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/requirements.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/requirements.txt
@@ -1,2 +1,2 @@
 azure-cosmos>=4.9.0,<5.0.0
 

```

### ✅ 11. view — 7080ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089\\requirements.txt"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-skill-mcp-2409803089/cosmos_crud.py
@@ -1,86 +1,86 @@
 """Perform a complete CRUD lifecycle in an Azure Cosmos DB for NoSQL container."""
 
 import os
 import sys
 from typing import Any
 
 from azure.cosmos import CosmosClient, PartitionKey, exceptions
 
 
 DATABASE_NAME = "TestDB"
 CONTAINER_NAME = "Items"
 PARTITION_KEY_PATH = "/category"
 
 
 def run_crud_operations() -> None:
     endpoint = os.environ["COSMOS_ENDPOINT"]
     key = os.environ["COSMOS_KEY"]
 
     client = CosmosClient(endpoint, credential=key)
     database = client.create_database_if_not_exists(id=DATABASE_NAME)
     container = database.create_container_if_not_exists(
         id=CONTAINER_NAME,
         partition_key=PartitionKey(path=PARTITION_KEY_PATH),
         offer_throughput=400,
     )
 
     item: dict[str, Any] = {
         "id": "item-001",
         "category": "electronics",
         "name": "Wireless Mouse",
         "quantity": 10,
     }
 
     upserted_item = container.upsert_item(body=item)
     print(f"Upserted item: {upserted_item}")
 
     read_item = container.read_item(
         item=item["id"],
         partition_key=item["category"],
     )
     print(f"Read item: {read_item}")
 
     query = "SELECT * FROM items i WHERE i.category = @category"
     parameters = [{"name": "@category", "value": "electronics"}]
     queried_items = list(
         container.query_items(
             query=query,
             parameters=parameters,
             part
... (truncated)
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace and ground the implementation in current Azure SDK guidance before creating the local-only sample.

The workspace is empty. I’m retrieving the exact SDK method patterns, then I’ll add the executable script and dependency manifest.

I’ve settled on an emulator-friendly configuration: credentials come only from environment variables, and service failures are surfaced through `azure.cosmos.exceptions`.

The CRUD flow is in place. I’m making the container creation self-sufficient for provisioned-throughput accounts and checking the script’s syntax.

Created:

- `cosmos_crud.py` — complete Cosmos DB CRUD lifecycle with parameterized querying and `azure.cosmos.exceptions` handling.
- `requirements.txt` — requires `azure-cosmos>=4.9.0,<5.0.0`.

Set `COSMOS_ENDPOINT` and `COSMOS_KEY`, install with `pip install -r requirements.txt`, then run `python cosmos_crud.py`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Azure Cosmos DB request failures are explicitly caught and surfaced appropriately.

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
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 17.2% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 4.8333** | **60.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
