# Evaluation Report: cosmos-db-dp-java-crud

**Config:** java-cosmos-skill/without-azure-sdk-java | **Result:** ❌ FAILED | **Duration:** 345.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-java-crud` |
| Config | java-cosmos-skill/without-azure-sdk-java |
| Result | ❌ FAILED |
| Score | 14/19 |
| Duration | 345.4s |
| Timestamp | 2026-08-26T03:23:47Z |
| Files Generated | 3 |
| Event Count | 4699 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 117.8s |
| Review | 226.4s |
| **Total** | **345.4s** |

## Configuration

- **name:** java-cosmos-skill/without-azure-sdk-java
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=243999 out=3235 |
| Turn Count | 7 |

## Prompt Sent

```
Write a Java program that performs
CRUD operations on items in an Azure Cosmos DB NoSQL container:
1. Create a CosmosClient using endpoint and key with CosmosClientBuilder
2. Create a database "TestDB" and container "Items" with partition key "/category"
3. Insert a POJO item with properties: id, category, name, quantity
4. Read the item back by id and partition key value
5. Query items where category equals "electronics" using parameterized SQL
6. Replace the item with updated quantity
7. Delete the item by id and partition key

Show required Maven dependency and handle CosmosException appropriately.
```

## Tool Calls

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local-only Java Maven example using the Azure Cosmos DB NoSQL Java SDK that demonstrates CRUD, parameterized queries, and CosmosException handling without provisioning live resources.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only Java Maven example for Azure Cosmos DB NoSQL CRUD using the Java SDK, parameterized queries, and CosmosException handling.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Azure Cosmos DB Java SDK v4 examples for CosmosClientBuilder endpoint/key authentication, createDatabaseIfNotExists, createContainerIfNotExists with /category, item CRUD, parameterized SQL query, and CosmosException handling.","learn":true}
```

### 🔧 5. web_search (via github-mcp-server)

**Input:**

```json
{"query":"What is the latest stable Maven Central version of com.azure:azure-cosmos as of August 2026? Cite Maven Central or mvnrepository and official Azure documentation."}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve official current Java SDK examples for Azure Cosmos DB NoSQL CRUD and parameterized queries.","parameters":{"language":"java","query":"Azure Cosmos DB Java SDK v4 CosmosClientBuilder createDatabaseIfNotExists createContainerIfNotExists createItem readItem queryItems SqlQuerySpec SqlParameter replaceItem deleteItem CosmosException"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure Cosmos DB Java SDK v4 CRUD documentation and Java quickstart.","parameters":{"query":"Azure Cosmos DB NoSQL Java SDK v4 CRUD quickstart CosmosClientBuilder create item read query replace delete"}}
```

### 🔧 8. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ecosmos-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+            \u003cversion\u003e4.81.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.CosmosCrudApp\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/CosmosCrudApp.java\n+package com.example;\n+\n+import com.azure.cosmos.ConsistencyLevel;\n+import com.azure.cosmos.CosmosClient;\n+import com.azure.cosmos.CosmosClientBuilder;\n+import com.azure.cosmos.CosmosContainer;\n+import com.azure.cosmos.CosmosDatabase;\n+import com.azure.cosmos.CosmosException;\n+import com.azure.cosmos.models.CosmosContainerProperties;\n+import com.azure.cosmos.models.CosmosItemRequestOptions;\n+import com.azure.cosmos.models.CosmosItemResponse;\n+import com.azure.cosmos.models.CosmosQueryRequestOptions;\n+import com.azure.cosmos.models.PartitionKey;\n+import com.azure.cosmos.models.SqlParameter;\n+import com.azure.cosmos.models.SqlQuerySpec;\n+\n+import java.util.List;\n+import java.util.UUID;\n+\n+public final class CosmosCrudApp {\n+    private static final String DATABASE_NAME = \"TestDB\";\n+    private static final String CONTAINER_NAME = \"Items\";\n+    private static final String CATEGORY = \"electronics\";\n+\n+    private CosmosCrudApp() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String endpoint = requiredEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+        String key = requiredEnvironmentVariable(\"COSMOS_KEY\");\n+\n+        try (CosmosClient client = new CosmosClientBuilder()\n+                .endpoint(endpoint)\n+                .key(key)\n+                .consistencyLevel(ConsistencyLevel.SESSION)\n+                .buildClient()) {\n+\n+            client.createDatabaseIfNotExists(DATABASE_NAME);\n+            CosmosDatabase database = client.getDatabase(DATABASE_NAME);\n+\n+            CosmosContainerProperties properties =\n+                    new CosmosContainerProperties(CONTAINER_NAME, \"/category\");\n+            database.createContainerIfNotExists(properties);\n+            CosmosContainer container = database.getContainer(CONTAINER_NAME);\n+\n+            Item item = new Item(\n+                    UUID.randomUUID().toString(),\n+                    CATEGORY,\n+                    \"Wireless headphones\",\n+                    10);\n+\n+            CosmosItemResponse\u003cItem\u003e createResponse =\n+                    container.createItem(item, new PartitionKey(item.getCategory()),\n+                            new CosmosItemRequestOptions());\n+            System.out.printf(\"Created item %s (request charge: %.2f RU)%n\",\n+                    item.getId(), createResponse.getRequestCharge());\n+\n+            CosmosItemResponse\u003cItem\u003e readResponse =\n+                    container.readItem(item.getId(), new PartitionKey(item.getCategory()), Item.class);\n+            Item readItem = readResponse.getItem();\n+            System.out.printf(\"Read item: %s, quantity: %d%n\",\n+                    readItem.getName(), readItem.getQuantity());\n+\n+            SqlQuerySpec query = new SqlQuerySpec(\n+                    \"SELECT * FROM c WHERE c.category = @category\",\n+                    List.of(new SqlParameter(\"@category\", CATEGORY)));\n+\n+            System.out.println(\"Items in category '\" + CATEGORY + \"':\");\n+            container.queryItems(query, new CosmosQueryRequestOptions(), Item.class)\n+                    .forEach(result -\u003e System.out.printf(\"- %s: %s (quantity: %d)%n\",\n+                            result.getId(), result.getName(), result.getQuantity()));\n+\n+            readItem.setQuantity(25);\n+            CosmosItemResponse\u003cItem\u003e replaceResponse =\n+                    container.replaceItem(\n+                            readItem,\n+                            readItem.getId(),\n+                            new PartitionKey(readItem.getCategory()),\n+                            new CosmosItemRequestOptions());\n+            System.out.printf(\"Updated quantity to %d (request charge: %.2f RU)%n\",\n+                    replaceResponse.getItem().getQuantity(), replaceResponse.getRequestCharge());\n+\n+            CosmosItemResponse\u003cObject\u003e deleteResponse =\n+                    container.deleteItem(\n+                            readItem.getId(),\n+                            new PartitionKey(readItem.getCategory()),\n+                            new CosmosItemRequestOptions());\n+            System.out.printf(\"Deleted item %s (request charge: %.2f RU)%n\",\n+                    readItem.getId(), deleteResponse.getRequestCharge());\n+        } catch (CosmosException exception) {\n+            System.err.printf(\n+                    \"Cosmos DB request failed: status=%d, substatus=%d, activityId=%s, message=%s%n\",\n+                    exception.getStatusCode(),\n+                    exception.getSubStatusCode(),\n+                    exception.getActivityId(),\n+                    exception.getMessage());\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    public static final class Item {\n+        private String id;\n+        private String category;\n+        private String name;\n+        private int quantity;\n+\n+        public Item() {\n+        }\n+\n+        public Item(String id, String category, String name, int quantity) {\n+            this.id = id;\n+            this.category = category;\n+            this.name = name;\n+            this.quantity = quantity;\n+        }\n+\n+        public String getId() {\n+            return id;\n+        }\n+\n+        public void setId(String id) {\n+            this.id = id;\n+        }\n+\n+        public String getCategory() {\n+            return category;\n+        }\n+\n+        public void setCategory(String category) {\n+            this.category = category;\n+        }\n+\n+        public String getName() {\n+            return name;\n+        }\n+\n+        public void setName(String name) {\n+            this.name = name;\n+        }\n+\n+        public int getQuantity() {\n+            return quantity;\n+        }\n+\n+        public void setQuantity(int quantity) {\n+            this.quantity = quantity;\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Cosmos DB Java CRUD example\n+\n+This Maven project creates the `TestDB` database and an `Items` container partitioned\n+by `/category`, then creates, reads, queries, replaces, and deletes an item.\n+\n+## Prerequisites\n+\n+- Java 17 or later\n+- Maven 3.9 or later\n+- An Azure Cosmos DB for NoSQL endpoint and key\n+\n+Set credentials as environment variables rather than placing them in source code:\n+\n+```powershell\n+$env:COSMOS_ENDPOINT = \"https://\u003caccount\u003e.documents.azure.com:443/\"\n+$env:COSMOS_KEY = \"\u003caccount-key\u003e\"\n+mvn compile exec:java\n+```\n+\n+The required SDK dependency is:\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+    \u003cversion\u003e4.81.0\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+For production Azure-hosted applications, prefer Microsoft Entra ID and managed\n+identity over key authentication. Endpoint/key authentication is used here because\n+it is specifically required by the example.\n+\n+References:\n+\n+- [Azure Cosmos DB Java SDK overview](https://learn.microsoft.com/java/api/overview/azure/cosmos)\n+- [Azure Cosmos DB Java SDK samples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)\n*** End Patch\n"
```

### 🔧 9. powershell

**Input:**

```json
{"command":"mvn --quiet compile","description":"Compile the Maven Cosmos CRUD sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\CosmosCrudApp.java`

## Copilot Response

I’ll inspect the workspace and ground the implementation in current Azure Java guidance before creating the Maven example.

The workspace is empty. I’ll create a minimal Maven project that reads credentials from environment variables and performs the requested operations sequentially.

The implementation will use Azure Cosmos SDK `4.81.0`, environment-based endpoint/key configuration, a uniquely identified POJO, and a single `try`/`catch`/resource-cleanup flow with status-aware `CosmosException` reporting.

The project files are in place; I’m compiling the sample now to catch any SDK API or type mismatches.

Created a complete Maven Java 17 example:

- `pom.xml` — Azure Cosmos DB SDK `4.81.0`
- `src\main\java\com\example\CosmosCrudApp.java` — database/container creation and full CRUD flow
- `README.md` — configuration, execution instructions, security guidance, and official references

Credentials are read from `COSMOS_ENDPOINT` and `COSMOS_KEY`, SQL uses `SqlQuerySpec` with `SqlParameter`, and Cosmos failures report status, substatus, activity ID, and message.

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
mvn compile exec:java
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criterion passed.

### Strengths

- CosmosClient lifecycle is correctly managed with try-with-resources.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - `azure-cosmos` Maven dependency (com.azure:azure-cosmos): Pass
      - `CosmosClientBuilder` and `CosmosClient`: Pass
      - `CosmosDatabase` and `CosmosContainer` creation: Pass
      - `CosmosContainer.createItem()`, `readItem()`, `replaceItem()`, `deleteItem()`: Pass
      - `CosmosQueryRequestOptions` and `CosmosPagedIterable`: Fail
      - `SqlQuerySpec` with parameters: Pass
      - `CosmosException` error handling: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Fail
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Fail
  - Client Builder Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Fail
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Pass
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Pass
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 9.7% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 11.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 8.8571** | **68.1%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-java-crud --config java-cosmos-skill/without-azure-sdk-java
```

---

[← Back to Summary](../../../../../../summary.md)
