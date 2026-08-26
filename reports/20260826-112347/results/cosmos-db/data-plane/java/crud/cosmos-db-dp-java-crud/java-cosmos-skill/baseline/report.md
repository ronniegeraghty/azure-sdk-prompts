# Evaluation Report: cosmos-db-dp-java-crud

**Config:** java-cosmos-skill/baseline | **Result:** ❌ FAILED | **Duration:** 476.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-java-crud` |
| Config | java-cosmos-skill/baseline |
| Result | ❌ FAILED |
| Score | 16/19 |
| Duration | 476.6s |
| Timestamp | 2026-08-26T03:29:33Z |
| Files Generated | 3 |
| Event Count | 5308 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 137.6s |
| Review | 337.9s |
| **Total** | **476.6s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-cosmos-skill/baseline

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=565191 out=4753 |
| Turn Count | 12 |

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
{"intent":"Create a local-only Java Maven sample demonstrating synchronous Azure Cosmos DB NoSQL CRUD with CosmosClientBuilder, parameterized queries, partition keys, and CosmosException handling.","learn":true}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for a local Java Maven Azure Cosmos DB CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-cosmos-java\\references\\examples.md"}
```

### 🔧 5. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-java-crud-java-cosmos-skill-baseline-3139622555","pattern":"**/*"}
```

### 🔧 6. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-cosmos-java","pattern":"**/*acceptance*"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-cosmos/maven-metadata.xml"}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Java Azure Cosmos DB SDK v4 documentation for synchronous CRUD operations, parameterized SQL queries, and CosmosException handling.","learn":true}
```

### 🔧 9. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Ground a Java Cosmos DB CRUD sample in official documentation.","parameters":{"query":"Azure Cosmos DB for NoSQL Java SDK v4 synchronous create database container CRUD item parameterized query CosmosException"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find latest official Java examples for synchronous Cosmos DB CRUD and parameterized queries.","parameters":{"language":"java","query":"CosmosClientBuilder CosmosClient createDatabaseIfNotExists createContainerIfNotExists createItem readItem queryItems SqlQuerySpec replaceItem deleteItem CosmosException"}}
```

### 🔧 11. rg

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715040296-copilot-tool-output-e4c53c259cbc41f5a6a20f6f2bbf17cc.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

### 🔧 12. rg

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715039532-copilot-tool-output-ffd7abd8be2740cb8f43662497b68280.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

### 🔧 13. rg

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715040296-copilot-tool-output-e4c53c259cbc41f5a6a20f6f2bbf17cc.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715039532-copilot-tool-output-ffd7abd8be2740cb8f43662497b68280.txt"],"pattern":"parameter|SqlQuerySpec|createDatabaseIfNotExists|CosmosException"}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715040296-copilot-tool-output-e4c53c259cbc41f5a6a20f6f2bbf17cc.txt' | ConvertFrom-Json; $j.results | Select-Object title, contentUrl | Format-Table -AutoSize","description":"List documentation result URLs"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715039532-copilot-tool-output-ffd7abd8be2740cb8f43662497b68280.txt' | ConvertFrom-Json; $j.results | Select-Object description, sourceUrl | Format-List","description":"List code sample source URLs"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715040296-copilot-tool-output-e4c53c259cbc41f5a6a20f6f2bbf17cc.txt","view_range":[1,8]}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787715039532-copilot-tool-output-ffd7abd8be2740cb8f43662497b68280.txt","view_range":[1,8]}
```

### 🔧 18. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official Azure Cosmos DB Java SDK overview and examples before implementing the sample.","parameters":{"url":"https://learn.microsoft.com/java/api/overview/azure/cosmos-readme?view=azure-java-stable"}}
```

### 🔧 19. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ecosmos-crud\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+            \u003cversion\u003e4.81.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.cosmos.CosmosCrudExample\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/cosmos/CosmosCrudExample.java\n+package com.example.cosmos;\n+\n+import com.azure.cosmos.CosmosClient;\n+import com.azure.cosmos.CosmosClientBuilder;\n+import com.azure.cosmos.CosmosContainer;\n+import com.azure.cosmos.CosmosDatabase;\n+import com.azure.cosmos.CosmosException;\n+import com.azure.cosmos.models.CosmosItemResponse;\n+import com.azure.cosmos.models.CosmosQueryRequestOptions;\n+import com.azure.cosmos.models.PartitionKey;\n+import com.azure.cosmos.models.SqlParameter;\n+import com.azure.cosmos.models.SqlQuerySpec;\n+import com.azure.cosmos.util.CosmosPagedIterable;\n+\n+import java.util.List;\n+\n+public final class CosmosCrudExample {\n+    private static final String DATABASE_ID = \"TestDB\";\n+    private static final String CONTAINER_ID = \"Items\";\n+    private static final String PARTITION_KEY_PATH = \"/category\";\n+\n+    private CosmosCrudExample() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+        String key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+\n+        try (CosmosClient client = new CosmosClientBuilder()\n+            .endpoint(endpoint)\n+            .key(key)\n+            .buildClient()) {\n+\n+            runCrudOperations(client);\n+        } catch (CosmosException exception) {\n+            reportCosmosException(exception);\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static void runCrudOperations(CosmosClient client) {\n+        client.createDatabaseIfNotExists(DATABASE_ID);\n+        CosmosDatabase database = client.getDatabase(DATABASE_ID);\n+\n+        database.createContainerIfNotExists(CONTAINER_ID, PARTITION_KEY_PATH);\n+        CosmosContainer container = database.getContainer(CONTAINER_ID);\n+\n+        Item item = new Item(\"item-1\", \"electronics\", \"Wireless keyboard\", 10);\n+        PartitionKey partitionKey = new PartitionKey(item.getCategory());\n+\n+        CosmosItemResponse\u003cItem\u003e createResponse = container.createItem(item, partitionKey, null);\n+        System.out.printf(\n+            \"Created item %s (request charge: %.2f RU)%n\",\n+            createResponse.getItem().getId(),\n+            createResponse.getRequestCharge());\n+\n+        CosmosItemResponse\u003cItem\u003e readResponse =\n+            container.readItem(item.getId(), partitionKey, Item.class);\n+        Item readItem = readResponse.getItem();\n+        System.out.printf(\n+            \"Read item: id=%s, category=%s, name=%s, quantity=%d%n\",\n+            readItem.getId(),\n+            readItem.getCategory(),\n+            readItem.getName(),\n+            readItem.getQuantity());\n+\n+        String category = \"electronics\";\n+        SqlQuerySpec query = new SqlQuerySpec(\n+            \"SELECT * FROM c WHERE c.category = @category\",\n+            List.of(new SqlParameter(\"@category\", category)));\n+        CosmosQueryRequestOptions queryOptions = new CosmosQueryRequestOptions();\n+        queryOptions.setPartitionKey(new PartitionKey(category));\n+\n+        CosmosPagedIterable\u003cItem\u003e queryResults =\n+            container.queryItems(query, queryOptions, Item.class);\n+        queryResults.forEach(result -\u003e System.out.printf(\n+            \"Query result: id=%s, name=%s, quantity=%d%n\",\n+            result.getId(),\n+            result.getName(),\n+            result.getQuantity()));\n+\n+        readItem.setQuantity(25);\n+        CosmosItemResponse\u003cItem\u003e replaceResponse =\n+            container.replaceItem(readItem, readItem.getId(), partitionKey, null);\n+        System.out.printf(\n+            \"Replaced item %s with quantity %d (request charge: %.2f RU)%n\",\n+            replaceResponse.getItem().getId(),\n+            replaceResponse.getItem().getQuantity(),\n+            replaceResponse.getRequestCharge());\n+\n+        container.deleteItem(item.getId(), partitionKey, null);\n+        System.out.printf(\"Deleted item %s%n\", item.getId());\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\n+                \"Required environment variable \" + name + \" is not set.\");\n+        }\n+        return value;\n+    }\n+\n+    private static void reportCosmosException(CosmosException exception) {\n+        System.err.printf(\n+            \"Azure Cosmos DB request failed: status=%d, substatus=%d, activityId=%s, \"\n+                + \"requestCharge=%.2f RU, message=%s%n\",\n+            exception.getStatusCode(),\n+            exception.getSubStatusCode(),\n+            exception.getActivityId(),\n+            exception.getRequestCharge(),\n+            exception.getMessage());\n+\n+        if (exception.getStatusCode() == 409) {\n+            System.err.println(\"An item with the same id and partition key already exists.\");\n+        } else if (exception.getStatusCode() == 429) {\n+            System.err.printf(\n+                \"Request rate was limited. Retry after %s.%n\",\n+                exception.getRetryAfterDuration());\n+        }\n+    }\n+\n+    public static final class Item {\n+        private String id;\n+        private String category;\n+        private String name;\n+        private int quantity;\n+\n+        public Item() {\n+        }\n+\n+        public Item(String id, String category, String name, int quantity) {\n+            this.id = id;\n+            this.category = category;\n+            this.name = name;\n+            this.quantity = quantity;\n+        }\n+\n+        public String getId() {\n+            return id;\n+        }\n+\n+        public void setId(String id) {\n+            this.id = id;\n+        }\n+\n+        public String getCategory() {\n+            return category;\n+        }\n+\n+        public void setCategory(String category) {\n+            this.category = category;\n+        }\n+\n+        public String getName() {\n+            return name;\n+        }\n+\n+        public void setName(String name) {\n+            this.name = name;\n+        }\n+\n+        public int getQuantity() {\n+            return quantity;\n+        }\n+\n+        public void setQuantity(int quantity) {\n+            this.quantity = quantity;\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Cosmos DB Java CRUD example\n+\n+This Maven project creates a synchronous `CosmosClient` and demonstrates create,\n+read, parameterized query, replace, and delete operations in an Azure Cosmos DB\n+for NoSQL container.\n+\n+## Configuration\n+\n+Set the account endpoint and key as environment variables. Do not store account\n+keys in source control.\n+\n+```powershell\n+$env:COSMOS_ENDPOINT = \"https://\u003caccount\u003e.documents.azure.com:443/\"\n+$env:COSMOS_KEY = \"\u003caccount-key\u003e\"\n+```\n+\n+## Run\n+\n+```powershell\n+mvn compile exec:java\n+```\n+\n+The program creates `TestDB` and its `Items` container if they do not exist. The\n+container uses `/category` as its partition key.\n+\n+Key-based authentication is used because this example specifically demonstrates\n+the `CosmosClientBuilder.endpoint(...).key(...)` flow. For production workloads\n+hosted in Azure, prefer Microsoft Entra ID and managed identity when possible.\n+\n+## References\n+\n+- [Azure Cosmos DB client library for Java](https://learn.microsoft.com/java/api/overview/azure/cosmos-readme?view=azure-java-stable)\n+- [Azure Cosmos DB Java SDK v4 troubleshooting](https://learn.microsoft.com/azure/cosmos-db/troubleshoot-java-sdk-v4)\n*** End Patch\n"
```

### 🔧 20. powershell

**Input:**

```json
{"command":"mvn -q -DskipTests compile","description":"Compile the Maven project","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\cosmos\CosmosCrudExample.java`

## Copilot Response

I’ll build this as a small local Maven example using environment-variable credentials, then compile it without connecting to Azure.

The implementation will use a single reusable synchronous client, a partition-targeted parameterized query, and `try`-with-resources so the client always closes. Credentials remain in `COSMOS_ENDPOINT` and `COSMOS_KEY`.

The project files are in place. I’m compiling against Azure Cosmos SDK 4.81.0 now; this does not contact an Azure account.

Implemented a complete Maven project:

- `pom.xml` — Azure Cosmos DB SDK `4.81.0`
- `src\main\java\com\example\cosmos\CosmosCrudExample.java` — database/container creation and full CRUD workflow
- `README.md` — configuration and run instructions

The program uses `COSMOS_ENDPOINT` and `COSMOS_KEY`, a parameterized `SqlQuerySpec`, partition-targeted operations, `CosmosException` diagnostics, and automatic client cleanup.

Run with:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
mvn compile exec:java
```

Reference: [Azure Cosmos DB client library for Java](https://learn.microsoft.com/java/api/overview/azure/cosmos-readme?view=azure-java-stable).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses try-with-resources for the CosmosClient, providing automatic cleanup even when CRUD operations throw an exception.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `azure-cosmos` Maven dependency (com.azure:azure-cosmos): Pass
      - `CosmosClientBuilder` and `CosmosClient`: Pass
      - `CosmosDatabase` and `CosmosContainer` creation: Pass
      - `CosmosContainer.createItem()`, `readItem()`, `replaceItem()`, `deleteItem()`: Pass
      - `CosmosQueryRequestOptions` and `CosmosPagedIterable`: Pass
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
  - Client Builder Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.0000** | **76.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-java-crud --config java-cosmos-skill/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
