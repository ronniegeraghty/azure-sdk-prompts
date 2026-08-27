# Evaluation Report: cosmos-db-dp-python-crud

**Config:** python-azure-skills/azure | **Result:** ❌ FAILED | **Duration:** 183.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-crud` |
| Config | python-azure-skills/azure |
| Result | ❌ FAILED |
| Score | 9/13 |
| Duration | 183.8s |
| Timestamp | 2026-08-27T02:53:52Z |
| Files Generated | 2 |
| Event Count | 2143 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 83.4s |
| Review | 99.1s |
| **Total** | **183.8s** |

## Configuration

- **name:** python-azure-skills/azure
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=239122 out=2377 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. azure-get_azure_bestpractices (via azure) — 4437ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate a local-only Python Azure Cosmos DB for NoSQL CRUD script using CosmosClient, parameterized queries, and azure.cosmos exception handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 8767ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Python Azure Cosmos DB SDK code-generation best practices for a local-only CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 10234ms

**Input:**

```json
{"command":"search","intent":"Find official Azure Cosmos DB Python SDK documentation for CosmosClient, create_database_if_not_exists, create_container_if_not_exists, upsert_item, read_item, parameterized query_items, replace_item, delete_item, and Cosmos exception handling.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 7442ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-documentation (via azure) — 5164ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve current official Python Azure Cosmos DB SDK CRUD examples.","parameters":{"language":"python","query":"Azure Cosmos DB for NoSQL Python SDK CosmosClient create_database_if_not_exists create_container_if_not_exists upsert_item read_item parameterized query_items replace_item delete_item exceptions"}}
```

**Output:**

```
{"results":[{"description":"description: Creates a database in Azure Cosmos DB using the create_database_if_not_exists method, which returns the database if it exists and handles request failures.\npackage: azure-cosmos\nlanguage: python\n","codeSnippet":"try:\r\n    database = client.create_database_if_not_exists(id=DATABASE_ID)\r\n    print(f\u0022Database created or returned: {database.id}\u0022)\r\n\r\nexcept CosmosHttpResponseError:\r\n    print(\u0022Request to the Azure Cosmos database service failed.\u0022)","link":"https://learn.microsoft.com/azure/cosmos-db/how-to-python-create-database#create-a-database","language":"python"},{"description":"description: Creates a database in Azure Cosmos DB using the create_database_if_not_exists method. This method returns the database if it exists and does not throw an exception, handling errors with a catch for CosmosHttpResponseError.\npackage: azure-cosmos\nlanguage: python\n","codeSnippet":"try:\r\n    database = client.create_database_if_not_exists(id=DATABASE_ID)\r\n    print(f\u0022Database created or returned: {database.id}\u0022)\r\n\r\nexcept CosmosHttpResponseError:\r\n    print(\u0022Request to the Azure Cosmos database service failed.\u0022)","link":"https://learn.microsoft.com/azure/cosmos-db/how-to-python-create-database#create-a-database","language":"python"},{"description":"description: Creates a container in Azure Cosmos DB using the create_container_if_not_exists method. It specifies a partition key and throughput, and handles exceptions related to HTTP response errors, avoiding errors if the container already exists.\npackage: azure-cosmos\nlanguage: python\n","codeSnippet":"try:\r\n    partition_key_path = PartitionKey(path=\u0022/categoryId\u0022)\r\n    container = database.create_container_if_not_exists(\r\n        id=CONTAINER_ID,\r\n        partition_key=partition_key_path,\r\n        offer_throughput=400,\r\n    )\r\n    print(f\u0022Container created or returned: {container.id}\u0022)\r\n\r\n
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 4780ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Retrieve current official Azure Cosmos DB Python SDK CRUD and authentication guidance.","parameters":{"query":"Azure Cosmos DB Python SDK CRUD operations CosmosClient parameterized query exceptions local emulator"}}
```

**Output:**

```
{"results":[{"title":"Azure Cosmos DB SQL API client library for Python - version 4.16.3","content":"# Azure Cosmos DB SQL API client library for Python - version 4.16.3\n## Examples\nThe preceding snippets also handle the [CosmosHttpResponseError](https://aka.ms/azsdk-python-cosmos-ref-http-failure) exception if the container creation failed. For more information on error handling and troubleshooting, see the Troubleshooting section.\n### Get an existing container\nRetrieve an existing container from the database:\n\u0060\u0060\u0060python\nfrom azure.cosmos import CosmosClient\nimport os\n\nURL = os.environ[\u0027ACCOUNT_URI\u0027]\nKEY = os.environ[\u0027ACCOUNT_KEY\u0027]\nclient = CosmosClient(URL, credential=KEY)\nDATABASE_NAME = \u0027testDatabase\u0027\ndatabase = client.get_database_client(DATABASE_NAME)\nCONTAINER_NAME = \u0027products\u0027\ncontainer = database.get_container_client(CONTAINER_NAME)\n\u0060\u0060\u0060\n### Insert data\nTo insert items into a container, pass a dictionary containing your data to [ContainerProxy.upsert_item](https://aka.ms/azsdk-python-cosmos-ref-upsert-item). Each item you add to a container must include an \u0060id\u0060 key with a value that uniquely identifies the item within the container.\nThis example inserts several items into the container, each with a unique \u0060id\u0060:\n### Delete data\nTo delete items from a container, use [ContainerProxy.delete_item](https://aka.ms/azsdk-python-cosmos-ref-delete-item). The SQL API in Cosmos DB does not support the SQL \u0060DELETE\u0060 statement.\n\nNOTE: If you are using partitioned collection, the value of the \u0060partitionKey\u0060 in the example code above, should be set to the value of the partition key for this particular item, not the name of the partition key column in your collection. This holds true for both point reads and deletes.\n\n### Query the database\nA Cosmos DB SQL API database supports querying the items in a container with [ContainerProxy.query_items
... (truncated)
```

### ✅ 7. apply_patch — 2462ms

**Input:**

```json
"*** Begin Patch\n*** Add File: cosmos_crud.py\n+import os\n+from urllib.parse import urlparse\n+\n+from azure.cosmos import CosmosClient, PartitionKey, exceptions\n+\n+\n+DATABASE_NAME = \"TestDB\"\n+CONTAINER_NAME = \"Items\"\n+PARTITION_KEY_PATH = \"/category\"\n+LOCAL_HOSTS = {\"localhost\", \"127.0.0.1\", \"::1\"}\n+\n+\n+def get_local_connection() -\u003e tuple[str, str]:\n+    endpoint = os.getenv(\"COSMOS_ENDPOINT\", \"https://localhost:8081\")\n+    key = os.getenv(\"COSMOS_KEY\")\n+\n+    if urlparse(endpoint).hostname not in LOCAL_HOSTS:\n+        raise ValueError(\"COSMOS_ENDPOINT must point to a local Cosmos DB Emulator.\")\n+    if not key:\n+        raise ValueError(\"Set COSMOS_KEY to the Cosmos DB Emulator primary key.\")\n+\n+    return endpoint, key\n+\n+\n+def main() -\u003e None:\n+    endpoint, key = get_local_connection()\n+    client = CosmosClient(endpoint, credential=key)\n+\n+    item = {\n+        \"id\": \"item-001\",\n+        \"category\": \"electronics\",\n+        \"name\": \"Wireless Mouse\",\n+        \"quantity\": 10,\n+    }\n+\n+    try:\n+        database = client.create_database_if_not_exists(id=DATABASE_NAME)\n+        container = database.create_container_if_not_exists(\n+            id=CONTAINER_NAME,\n+            partition_key=PartitionKey(path=PARTITION_KEY_PATH),\n+        )\n+\n+        upserted_item = container.upsert_item(body=item)\n+        print(f\"Upserted: {upserted_item}\")\n+\n+        read_item = container.read_item(\n+            item=item[\"id\"],\n+            partition_key=item[\"category\"],\n+        )\n+        print(f\"Read: {read_item}\")\n+\n+        query = \"SELECT * FROM items i WHERE i.category = @category\"\n+        parameters = [{\"name\": \"@category\", \"value\": \"electronics\"}]\n+        queried_items = list(\n+            container.query_items(\n+                query=query,\n+                parameters=parameters,\n+                partition_key=\"electronics\",\n+            )\n+        )\n+        print(f\"Query results: {queried_items}\")\n+\n+        updated_item = dict(read_item)\n+        updated_item[\"quantity\"] = 25\n+        replaced_item = container.replace_item(\n+            item=item[\"id\"],\n+            body=updated_item,\n+        )\n+        print(f\"Replaced: {replaced_item}\")\n+\n+        container.delete_item(\n+            item=item[\"id\"],\n+            partition_key=item[\"category\"],\n+        )\n+        print(f\"Deleted item: {item['id']}\")\n+    except exceptions.CosmosResourceNotFoundError as exc:\n+        print(f\"Cosmos DB resource was not found: {exc}\")\n+        raise\n+    except exceptions.CosmosHttpResponseError as exc:\n+        print(f\"Cosmos DB request failed (status {exc.status_code}): {exc}\")\n+        raise\n+    finally:\n+        client.close()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        main()\n+    except (ValueError, exceptions.CosmosHttpResponseError) as exc:\n+        raise SystemExit(1) from exc\n*** Add File: requirements.txt\n+azure-cosmos==4.16.3\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-2264254372/cosmos_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-2264254372/cosmos_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-python-crud-python-azure-skills-azure-2264254372/cosmos_crud.py
@@ -1,0 +1,91 @@
+import os
+from urllib.parse import urlparse
+
+from azure.cosmos import CosmosClient, PartitionKey, exceptions
+
+
+DATABASE_NAME = "TestDB"
+CONTAINER_NAME = "Items"
+PARTITION_KEY_PATH = "/category"
+LOCAL_HOSTS = {"localhost", "127.0.0.1", "::1"}
+
+
+def get_local_connection() -> tuple[str, str]:
+    endpoint = os.getenv("COSMOS_ENDPOINT", "https://localhost:8081")
+    key = os.getenv("COSMOS_KEY")
+
+    if urlparse(endpoint).hostname not in LOCAL_HOSTS:
+        raise ValueError("COSMOS_ENDPOINT must point to a local Cosmos DB Emulator.")
+    if not key:
+        raise ValueError("Set COSMOS_KEY to the Cosmos DB Emulator primary key.")
+
+    return endpoint, key
+
+
+def main() -> None:
+    endpoint, key = get_local_connection()
+    client = CosmosClient(endpoint, credential=key)
+
+    item = {
+        "id": "item-001",
+        "category": "electronics",
+        "name": "Wireless Mouse",
+        "quantity": 10,
+    }
+
+    try:
+        database = client.create_database_if_not_exists(id=DATABASE_NAME)
+        container = database.create_container_if_not_exists(
+            id=CONTAINER_NAME,
+            partition_key=PartitionKey(path=PARTITION_KEY_PATH),
+        )
+
+        upserted_item = container.upsert_item(body=item)
+        print(f"Upserted: {upserted_item}")
+
+        read_item = container.read_item(
+            item=item["id"],
+            partition_key=item["category"],
+        )
+        print(f"Read: {read_item}")
+
+        query = "SELECT * FROM items i WHERE i.category = @ca
... (truncated)
```

### ✅ 8. powershell — 3757ms

**Input:**

```json
{"command":"python -m py_compile .\\cosmos_crud.py","description":"Compile-check the Cosmos CRUD script","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `cosmos_crud.py`
- `requirements.txt`

## Copilot Response

I’ll create a local-only Cosmos DB CRUD script using environment-based emulator credentials, plus its pip dependency file.

The implementation will target the local Cosmos DB Emulator only, require its key via `COSMOS_KEY`, and keep all queries parameterized.

Created:

- `cosmos_crud.py` — complete CRUD flow with parameterized query, local Cosmos DB Emulator configuration, and `azure.cosmos.exceptions` handling.
- `requirements.txt` — pins `azure-cosmos==4.16.3`.

Set `COSMOS_KEY` to the emulator primary key; `COSMOS_ENDPOINT` defaults to `https://localhost:8081`. Based on the official [Azure Cosmos DB Python SDK guidance](https://learn.microsoft.com/azure/cosmos-db/how-to-python-create-database).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses specific and general Azure Cosmos DB exception handlers without silently swallowing failures.

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 17.2% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 20.7% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 4.8333** | **60.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-crud --config python-azure-skills/azure
```

---

[← Back to Summary](../../../../../../summary.md)
