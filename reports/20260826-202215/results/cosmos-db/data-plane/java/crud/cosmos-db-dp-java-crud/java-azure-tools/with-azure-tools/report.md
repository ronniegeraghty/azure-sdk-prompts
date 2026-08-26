# Evaluation Report: cosmos-db-dp-java-crud

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 373.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-java-crud` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 14/19 |
| Duration | 373.5s |
| Timestamp | 2026-08-26T14:54:41Z |
| Files Generated | 3 |
| Event Count | 5279 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 109.6s |
| Review | 262.6s |
| **Total** | **373.5s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=307096 out=4203 |
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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-cosmos-java"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_code_gen_best_practices","intent":"Generate a local Java Maven example using the Azure Cosmos DB SDK for NoSQL CRUD operations with key authentication and CosmosException handling."}
```

### 🔧 3. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-cosmos-java\\references\\examples.md"}
```

### 🔧 4. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Java SDK v4 examples for Cosmos DB NoSQL synchronous CRUD, parameterized SQL query using SqlQuerySpec and SqlParameter, create database/container if not exists, replace and delete items.","parameters":{"query":"Azure Cosmos DB Java SDK v4 synchronous CRUD SqlQuerySpec SqlParameter createDatabaseIfNotExists createContainerIfNotExists"}}
```

### 🔧 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure code-generation best practices to a local Java Maven Cosmos DB NoSQL CRUD sample. Do not provision or deploy resources.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. web_fetch

**Input:**

```json
{"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-cosmos%22\u0026rows=1\u0026wt=json"}
```

### 🔧 8. rg

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787756121408-copilot-tool-output-ace296b2faac40f9bee740a7232b8351.txt","pattern":"SqlQuerySpec|createContainerIfNotExists|replaceItem|CosmosException"}
```

### 🔧 9. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-cosmos/maven-metadata.xml"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"mvn --version; java -version","description":"Check installed Java and Maven versions","initial_wait":30}
```

### 🔧 11. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\copilot-tool-output-1787756131439-1.txt","view_range":[1,40]}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ecosmos-crud-sample\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0-SNAPSHOT\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+            \u003cversion\u003e4.81.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/cosmos/CosmosCrudExample.java\n+package com.example.cosmos;\n+\n+import com.azure.cosmos.CosmosClient;\n+import com.azure.cosmos.CosmosClientBuilder;\n+import com.azure.cosmos.CosmosContainer;\n+import com.azure.cosmos.CosmosDatabase;\n+import com.azure.cosmos.CosmosException;\n+import com.azure.cosmos.models.CosmosItemRequestOptions;\n+import com.azure.cosmos.models.CosmosItemResponse;\n+import com.azure.cosmos.models.CosmosQueryRequestOptions;\n+import com.azure.cosmos.models.PartitionKey;\n+import com.azure.cosmos.models.SqlParameter;\n+import com.azure.cosmos.models.SqlQuerySpec;\n+\n+import java.util.List;\n+import java.util.UUID;\n+\n+public final class CosmosCrudExample {\n+    private static final String DATABASE_ID = \"TestDB\";\n+    private static final String CONTAINER_ID = \"Items\";\n+    private static final String PARTITION_KEY_PATH = \"/category\";\n+    private static final String CATEGORY = \"electronics\";\n+\n+    private CosmosCrudExample() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+        String key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+\n+        try (CosmosClient client = new CosmosClientBuilder()\n+            .endpoint(endpoint)\n+            .key(key)\n+            .buildClient()) {\n+\n+            CosmosDatabase database = createDatabase(client);\n+            CosmosContainer container = createContainer(database);\n+\n+            Item createdItem = createItem(container);\n+            Item readItem = readItem(container, createdItem.getId(), createdItem.getCategory());\n+            queryItems(container, CATEGORY);\n+            replaceItem(container, readItem, 25);\n+            deleteItem(container, createdItem.getId(), createdItem.getCategory());\n+        } catch (CosmosException exception) {\n+            reportCosmosException(exception);\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static CosmosDatabase createDatabase(CosmosClient client) {\n+        client.createDatabaseIfNotExists(DATABASE_ID);\n+        System.out.printf(\"Database ready: %s%n\", DATABASE_ID);\n+        return client.getDatabase(DATABASE_ID);\n+    }\n+\n+    private static CosmosContainer createContainer(CosmosDatabase database) {\n+        database.createContainerIfNotExists(CONTAINER_ID, PARTITION_KEY_PATH);\n+        System.out.printf(\"Container ready: %s (partition key: %s)%n\",\n+            CONTAINER_ID, PARTITION_KEY_PATH);\n+        return database.getContainer(CONTAINER_ID);\n+    }\n+\n+    private static Item createItem(CosmosContainer container) {\n+        Item item = new Item(\n+            UUID.randomUUID().toString(),\n+            CATEGORY,\n+            \"Wireless headphones\",\n+            10\n+        );\n+\n+        CosmosItemResponse\u003cItem\u003e response = container.createItem(\n+            item,\n+            new PartitionKey(item.getCategory()),\n+            new CosmosItemRequestOptions()\n+        );\n+        System.out.printf(\"Created item %s (request charge: %.2f RU)%n\",\n+            response.getItem().getId(), response.getRequestCharge());\n+        return response.getItem();\n+    }\n+\n+    private static Item readItem(CosmosContainer container, String id, String category) {\n+        CosmosItemResponse\u003cItem\u003e response = container.readItem(\n+            id,\n+            new PartitionKey(category),\n+            Item.class\n+        );\n+        Item item = response.getItem();\n+        System.out.printf(\"Read item: %s, quantity=%d (request charge: %.2f RU)%n\",\n+            item.getName(), item.getQuantity(), response.getRequestCharge());\n+        return item;\n+    }\n+\n+    private static void queryItems(CosmosContainer container, String category) {\n+        SqlQuerySpec query = new SqlQuerySpec(\n+            \"SELECT * FROM c WHERE c.category = @category\",\n+            List.of(new SqlParameter(\"@category\", category))\n+        );\n+        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()\n+            .setPartitionKey(new PartitionKey(category));\n+\n+        System.out.printf(\"Items in category '%s':%n\", category);\n+        container.queryItems(query, options, Item.class)\n+            .forEach(item -\u003e System.out.printf(\"  %s: %s (quantity=%d)%n\",\n+                item.getId(), item.getName(), item.getQuantity()));\n+    }\n+\n+    private static Item replaceItem(CosmosContainer container, Item item, int newQuantity) {\n+        item.setQuantity(newQuantity);\n+        CosmosItemResponse\u003cItem\u003e response = container.replaceItem(\n+            item,\n+            item.getId(),\n+            new PartitionKey(item.getCategory()),\n+            new CosmosItemRequestOptions()\n+        );\n+        System.out.printf(\"Replaced item %s with quantity=%d (request charge: %.2f RU)%n\",\n+            response.getItem().getId(),\n+            response.getItem().getQuantity(),\n+            response.getRequestCharge());\n+        return response.getItem();\n+    }\n+\n+    private static void deleteItem(CosmosContainer container, String id, String category) {\n+        container.deleteItem(\n+            id,\n+            new PartitionKey(category),\n+            new CosmosItemRequestOptions()\n+        );\n+        System.out.printf(\"Deleted item %s%n\", id);\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\n+                \"Required environment variable is not set: \" + name\n+            );\n+        }\n+        return value;\n+    }\n+\n+    private static void reportCosmosException(CosmosException exception) {\n+        System.err.printf(\n+            \"Cosmos DB request failed: status=%d, subStatus=%d, activityId=%s, message=%s%n\",\n+            exception.getStatusCode(),\n+            exception.getSubStatusCode(),\n+            exception.getActivityId(),\n+            exception.getMessage()\n+        );\n+\n+        if (exception.getStatusCode() == 409) {\n+            System.err.println(\"A resource with the same id and partition key already exists.\");\n+        } else if (exception.getStatusCode() == 429) {\n+            System.err.printf(\"Request was rate limited. Retry after %s.%n\",\n+                exception.getRetryAfterDuration());\n+        }\n+\n+        if (exception.getDiagnostics() != null) {\n+            System.err.println(\"Diagnostics: \" + exception.getDiagnostics());\n+        }\n+    }\n+\n+    public static final class Item {\n+        private String id;\n+        private String category;\n+        private String name;\n+        private int quantity;\n+\n+        public Item() {\n+        }\n+\n+        public Item(String id, String category, String name, int quantity) {\n+            this.id = id;\n+            this.category = category;\n+            this.name = name;\n+            this.quantity = quantity;\n+        }\n+\n+        public String getId() {\n+            return id;\n+        }\n+\n+        public void setId(String id) {\n+            this.id = id;\n+        }\n+\n+        public String getCategory() {\n+            return category;\n+        }\n+\n+        public void setCategory(String category) {\n+            this.category = category;\n+        }\n+\n+        public String getName() {\n+            return name;\n+        }\n+\n+        public void setName(String name) {\n+            this.name = name;\n+        }\n+\n+        public int getQuantity() {\n+            return quantity;\n+        }\n+\n+        public void setQuantity(int quantity) {\n+            this.quantity = quantity;\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Cosmos DB Java CRUD sample\n+\n+This Maven project uses the synchronous Azure Cosmos DB Java SDK to:\n+\n+1. Create the `TestDB` database and the `Items` container if they do not exist.\n+2. Create, point-read, query, replace, and delete an item.\n+3. Use `/category` as the partition key and a parameterized SQL query.\n+4. Report Cosmos DB status, retry, activity, and diagnostic information when a request fails.\n+\n+## Configuration\n+\n+Set credentials through environment variables. For local-only development, use values from\n+the Azure Cosmos DB emulator rather than credentials for a live account.\n+\n+```powershell\n+$env:COSMOS_ENDPOINT = \"https://localhost:8081/\"\n+$env:COSMOS_KEY = \"\u003cCOSMOS-DB-EMULATOR-KEY\u003e\"\n+```\n+\n+Do not commit keys to source control.\n+\n+## Build and run\n+\n+```powershell\n+mvn clean package\n+mvn exec:java -Dexec.mainClass=com.example.cosmos.CosmosCrudExample\n+```\n+\n+The required Maven dependency is:\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+    \u003cversion\u003e4.81.0\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+## References\n+\n+- [Azure Cosmos DB Java SDK overview](https://learn.microsoft.com/java/api/overview/azure/cosmos-readme)\n+- [Azure Cosmos DB Java SDK v4 samples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)\n+- [Troubleshoot the Azure Cosmos DB Java SDK v4](https://learn.microsoft.com/azure/cosmos-db/troubleshoot-java-sdk-v4)\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"mvn --quiet clean package","description":"Compile the Cosmos DB Java sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\cosmos\CosmosCrudExample.java`

## Copilot Response

I’ll build this as a small Maven project using the current `azure-cosmos` SDK, with credentials supplied only through environment variables and targeted partition-key querying.

Implemented a complete Maven-based Java CRUD sample.

- `pom.xml` — includes `com.azure:azure-cosmos:4.81.0`
- `src\main\java\com\example\cosmos\CosmosCrudExample.java` — creates `TestDB`/`Items`, inserts, reads, parameterized-queries, replaces, and deletes the POJO item
- `README.md` — documents emulator environment variables, build/run commands, dependency, and official references

Credentials are read from `COSMOS_ENDPOINT` and `COSMOS_KEY`; `CosmosException` handling reports status codes, rate limiting, activity IDs, and diagnostics.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses try-with-resources correctly for the CosmosClient lifecycle.

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
hyoka run --prompt-id cosmos-db-dp-java-crud --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
