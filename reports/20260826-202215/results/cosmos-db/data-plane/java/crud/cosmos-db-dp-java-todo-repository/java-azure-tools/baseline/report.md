# Evaluation Report: cosmos-db-dp-java-todo-repository

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 321.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-java-todo-repository` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 17/26 |
| Duration | 321.6s |
| Timestamp | 2026-08-26T15:00:55Z |
| Files Generated | 9 |
| Event Count | 11579 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 135.1s |
| Review | 185.2s |
| **Total** | **321.6s** |

## Configuration

- **name:** java-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=94379 out=7419 |
| Turn Count | 4 |

## Prompt Sent

```
Create a small Java 17 Maven project that implements a ToDo item CRUD repository backed by Azure Cosmos DB (NoSQL API).

The project needs:

- A **model class** (shared by both implementations) for a ToDo item with fields for id, title, description, completed status, created timestamp, and category (where category is the partition key).

- A **synchronous repository class** that performs CRUD operations against Cosmos DB. It should support create, read, update, delete, and a query-by-category method. Each operation should log the request charge (RU cost consumed). The update operation should prevent lost updates — if another process modified the item since it was last read, the update should fail with a clear conflict error rather than silently overwriting the other process's changes. The query method should use safe, parameterized queries and must handle large result sets properly — paginate through results rather than loading everything into memory at once, and log progress as each page is retrieved.

- An **asynchronous repository class** that provides the same CRUD operations. The query method should return results as a stream of pages, and the caller should be able to process each page as it arrives.

- A **configuration/factory class** that connects to the Cosmos DB account using its endpoint from an environment variable. Authentication must use managed identity (no master keys). It should create the database and container if they don't already exist, setting a default TTL (time-to-live) of 90 days on the container and configuring the indexing policy to exclude the `description` field from indexing (since it's never queried on).

- A **Main class** that demos both implementations: runs the full CRUD cycle using the sync repository first (including paginated query output showing page-by-page results), then runs the same operations using the async repository. Print RU costs and results to the console.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
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
{"pattern":"*"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ecosmos-todo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eCosmos DB ToDo Repository\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.cosmos.version\u003e4.75.0\u003c/azure.cosmos.version\u003e\n+        \u003cazure.identity.version\u003e1.15.3\u003c/azure.identity.version\u003e\n+        \u003cjackson.annotations.version\u003e2.18.2\u003c/jackson.annotations.version\u003e\n+        \u003cslf4j.version\u003e2.0.16\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.cosmos.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-annotations\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.annotations.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.13.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.todo.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/todo/ToDoItem.java\n+package com.example.todo;\n+\n+import com.fasterxml.jackson.annotation.JsonIgnore;\n+import com.fasterxml.jackson.annotation.JsonProperty;\n+\n+import java.time.Instant;\n+import java.util.Objects;\n+\n+public class ToDoItem {\n+    private String id;\n+    private String title;\n+    private String description;\n+    private boolean completed;\n+    private Instant createdAt;\n+    private String category;\n+\n+    @JsonProperty(\"_etag\")\n+    private String eTag;\n+\n+    public ToDoItem() {\n+    }\n+\n+    public ToDoItem(\n+            String id,\n+            String title,\n+            String description,\n+            boolean completed,\n+            Instant createdAt,\n+            String category) {\n+        this.id = Objects.requireNonNull(id, \"id\");\n+        this.title = Objects.requireNonNull(title, \"title\");\n+        this.description = description;\n+        this.completed = completed;\n+        this.createdAt = Objects.requireNonNull(createdAt, \"createdAt\");\n+        this.category = Objects.requireNonNull(category, \"category\");\n+    }\n+\n+    public String getId() {\n+        return id;\n+    }\n+\n+    public void setId(String id) {\n+        this.id = id;\n+    }\n+\n+    public String getTitle() {\n+        return title;\n+    }\n+\n+    public void setTitle(String title) {\n+        this.title = title;\n+    }\n+\n+    public String getDescription() {\n+        return description;\n+    }\n+\n+    public void setDescription(String description) {\n+        this.description = description;\n+    }\n+\n+    public boolean isCompleted() {\n+        return completed;\n+    }\n+\n+    public void setCompleted(boolean completed) {\n+        this.completed = completed;\n+    }\n+\n+    public Instant getCreatedAt() {\n+        return createdAt;\n+    }\n+\n+    public void setCreatedAt(Instant createdAt) {\n+        this.createdAt = createdAt;\n+    }\n+\n+    public String getCategory() {\n+        return category;\n+    }\n+\n+    public void setCategory(String category) {\n+        this.category = category;\n+    }\n+\n+    @JsonIgnore\n+    public String getETag() {\n+        return eTag;\n+    }\n+\n+    public void setETag(String eTag) {\n+        this.eTag = eTag;\n+    }\n+\n+    @Override\n+    public String toString() {\n+        return \"ToDoItem{\"\n+                + \"id='\" + id + '\\''\n+                + \", title='\" + title + '\\''\n+                + \", completed=\" + completed\n+                + \", createdAt=\" + createdAt\n+                + \", category='\" + category + '\\''\n+                + '}';\n+    }\n+}\n*** Add File: src/main/java/com/example/todo/RepositoryResult.java\n+package com.example.todo;\n+\n+public record RepositoryResult\u003cT\u003e(T value, double requestCharge) {\n+}\n*** Add File: src/main/java/com/example/todo/RepositoryPage.java\n+package com.example.todo;\n+\n+import java.util.List;\n+\n+public record RepositoryPage\u003cT\u003e(\n+        List\u003cT\u003e items,\n+        double requestCharge,\n+        String continuationToken) {\n+\n+    public RepositoryPage {\n+        items = List.copyOf(items);\n+    }\n+}\n*** Add File: src/main/java/com/example/todo/ConcurrentUpdateException.java\n+package com.example.todo;\n+\n+public class ConcurrentUpdateException extends RuntimeException {\n+    public ConcurrentUpdateException(String itemId, Throwable cause) {\n+        super(\"ToDo item '\" + itemId\n+                + \"' was modified after it was read; reload it before updating.\", cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/todo/SyncToDoRepository.java\n+package com.example.todo;\n+\n+import com.azure.cosmos.CosmosContainer;\n+import com.azure.cosmos.CosmosException;\n+import com.azure.cosmos.models.CosmosItemRequestOptions;\n+import com.azure.cosmos.models.CosmosItemResponse;\n+import com.azure.cosmos.models.CosmosQueryRequestOptions;\n+import com.azure.cosmos.models.FeedResponse;\n+import com.azure.cosmos.models.PartitionKey;\n+import com.azure.cosmos.models.SqlParameter;\n+import com.azure.cosmos.models.SqlQuerySpec;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.util.Iterator;\n+import java.util.List;\n+import java.util.NoSuchElementException;\n+import java.util.Objects;\n+\n+public class SyncToDoRepository {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncToDoRepository.class);\n+\n+    private final CosmosContainer container;\n+\n+    public SyncToDoRepository(CosmosContainer container) {\n+        this.container = Objects.requireNonNull(container, \"container\");\n+    }\n+\n+    public RepositoryResult\u003cToDoItem\u003e create(ToDoItem item) {\n+        CosmosItemResponse\u003cToDoItem\u003e response = container.createItem(\n+                item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions());\n+        return itemResult(\"create\", response);\n+    }\n+\n+    public RepositoryResult\u003cToDoItem\u003e read(String id, String category) {\n+        CosmosItemResponse\u003cToDoItem\u003e response = container.readItem(\n+                id, new PartitionKey(category), ToDoItem.class);\n+        return itemResult(\"read\", response);\n+    }\n+\n+    public RepositoryResult\u003cToDoItem\u003e update(ToDoItem item) {\n+        if (item.getETag() == null || item.getETag().isBlank()) {\n+            throw new IllegalArgumentException(\n+                    \"An ETag from a previous read is required for a safe update.\");\n+        }\n+\n+        CosmosItemRequestOptions options = new CosmosItemRequestOptions()\n+                .setIfMatchETag(item.getETag());\n+        try {\n+            CosmosItemResponse\u003cToDoItem\u003e response = container.replaceItem(\n+                    item, item.getId(), new PartitionKey(item.getCategory()), options);\n+            return itemResult(\"update\", response);\n+        } catch (CosmosException exception) {\n+            if (exception.getStatusCode() == 412) {\n+                throw new ConcurrentUpdateException(item.getId(), exception);\n+            }\n+            throw exception;\n+        }\n+    }\n+\n+    public RepositoryResult\u003cVoid\u003e delete(String id, String category) {\n+        CosmosItemResponse\u003cObject\u003e response = container.deleteItem(\n+                id, new PartitionKey(category), new CosmosItemRequestOptions());\n+        LOGGER.info(\"delete id={} requestCharge={} RU\", id, response.getRequestCharge());\n+        return new RepositoryResult\u003c\u003e(null, response.getRequestCharge());\n+    }\n+\n+    public Iterable\u003cRepositoryPage\u003cToDoItem\u003e\u003e queryByCategory(String category, int pageSize) {\n+        if (pageSize \u003c= 0) {\n+            throw new IllegalArgumentException(\"pageSize must be greater than zero\");\n+        }\n+\n+        SqlQuerySpec query = new SqlQuerySpec(\n+                \"SELECT * FROM c WHERE c.category = @category\",\n+                List.of(new SqlParameter(\"@category\", category)));\n+        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()\n+                .setPartitionKey(new PartitionKey(category));\n+        Iterable\u003cFeedResponse\u003cToDoItem\u003e\u003e pages = container\n+                .queryItems(query, options, ToDoItem.class)\n+                .iterableByPage(pageSize);\n+\n+        return () -\u003e mapPages(pages.iterator(), category);\n+    }\n+\n+    private Iterator\u003cRepositoryPage\u003cToDoItem\u003e\u003e mapPages(\n+            Iterator\u003cFeedResponse\u003cToDoItem\u003e\u003e source,\n+            String category) {\n+        return new Iterator\u003c\u003e() {\n+            private int pageNumber;\n+\n+            @Override\n+            public boolean hasNext() {\n+                return source.hasNext();\n+            }\n+\n+            @Override\n+            public RepositoryPage\u003cToDoItem\u003e next() {\n+                if (!hasNext()) {\n+                    throw new NoSuchElementException();\n+                }\n+                FeedResponse\u003cToDoItem\u003e response = source.next();\n+                pageNumber++;\n+                LOGGER.info(\n+                        \"query category={} page={} items={} requestCharge={} RU\",\n+                        category,\n+                        pageNumber,\n+                        response.getResults().size(),\n+                        response.getRequestCharge());\n+                return new RepositoryPage\u003c\u003e(\n+                        response.getResults(),\n+                        response.getRequestCharge(),\n+                        response.getContinuationToken());\n+            }\n+        };\n+    }\n+\n+    private RepositoryResult\u003cToDoItem\u003e itemResult(\n+            String operation,\n+            CosmosItemResponse\u003cToDoItem\u003e response) {\n+        LOGGER.info(\n+                \"{} id={} requestCharge={} RU\",\n+                operation,\n+                response.getItem().getId(),\n+                response.getRequestCharge());\n+        return new RepositoryResult\u003c\u003e(response.getItem(), response.getRequestCharge());\n+    }\n+}\n*** Add File: src/main/java/com/example/todo/AsyncToDoRepository.java\n+package com.example.todo;\n+\n+import com.azure.cosmos.CosmosAsyncContainer;\n+import com.azure.cosmos.CosmosException;\n+import com.azure.cosmos.models.CosmosItemRequestOptions;\n+import com.azure.cosmos.models.CosmosItemResponse;\n+import com.azure.cosmos.models.CosmosQueryRequestOptions;\n+import com.azure.cosmos.models.FeedResponse;\n+import com.azure.cosmos.models.PartitionKey;\n+import com.azure.cosmos.models.SqlParameter;\n+import com.azure.cosmos.models.SqlQuerySpec;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+import java.util.Objects;\n+import java.util.concurrent.atomic.AtomicInteger;\n+\n+public class AsyncToDoRepository {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncToDoRepository.class);\n+\n+    private final CosmosAsyncContainer container;\n+\n+    public AsyncToDoRepository(CosmosAsyncContainer container) {\n+        this.container = Objects.requireNonNull(container, \"container\");\n+    }\n+\n+    public Mono\u003cRepositoryResult\u003cToDoItem\u003e\u003e create(ToDoItem item) {\n+        return container.createItem(\n+                        item,\n+                        new PartitionKey(item.getCategory()),\n+                        new CosmosItemRequestOptions())\n+                .map(response -\u003e itemResult(\"create\", response));\n+    }\n+\n+    public Mono\u003cRepositoryResult\u003cToDoItem\u003e\u003e read(String id, String category) {\n+        return container.readItem(id, new PartitionKey(category), ToDoItem.class)\n+                .map(response -\u003e itemResult(\"read\", response));\n+    }\n+\n+    public Mono\u003cRepositoryResult\u003cToDoItem\u003e\u003e update(ToDoItem item) {\n+        if (item.getETag() == null || item.getETag().isBlank()) {\n+            return Mono.error(new IllegalArgumentException(\n+                    \"An ETag from a previous read is required for a safe update.\"));\n+        }\n+\n+        CosmosItemRequestOptions options = new CosmosItemRequestOptions()\n+                .setIfMatchETag(item.getETag());\n+        return container.replaceItem(\n+                        item,\n+                        item.getId(),\n+                        new PartitionKey(item.getCategory()),\n+                        options)\n+                .map(response -\u003e itemResult(\"update\", response))\n+                .onErrorMap(\n+                        error -\u003e error instanceof CosmosException cosmosException\n+                                \u0026\u0026 cosmosException.getStatusCode() == 412,\n+                        error -\u003e new ConcurrentUpdateException(item.getId(), error));\n+    }\n+\n+    public Mono\u003cRepositoryResult\u003cVoid\u003e\u003e delete(String id, String category) {\n+        return container.deleteItem(\n+                        id,\n+                        new PartitionKey(category),\n+                        new CosmosItemRequestOptions())\n+                .map(response -\u003e {\n+                    LOGGER.info(\n+                            \"delete id={} requestCharge={} RU\",\n+                            id,\n+                            response.getRequestCharge());\n+                    return new RepositoryResult\u003cVoid\u003e(null, response.getRequestCharge());\n+                });\n+    }\n+\n+    public Flux\u003cRepositoryPage\u003cToDoItem\u003e\u003e queryByCategory(String category, int pageSize) {\n+        if (pageSize \u003c= 0) {\n+            return Flux.error(new IllegalArgumentException(\n+                    \"pageSize must be greater than zero\"));\n+        }\n+\n+        SqlQuerySpec query = new SqlQuerySpec(\n+                \"SELECT * FROM c WHERE c.category = @category\",\n+                List.of(new SqlParameter(\"@category\", category)));\n+        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()\n+                .setPartitionKey(new PartitionKey(category));\n+        AtomicInteger pageNumber = new AtomicInteger();\n+\n+        return container.queryItems(query, options, ToDoItem.class)\n+                .byPage(pageSize)\n+                .map(response -\u003e toPage(response, category, pageNumber.incrementAndGet()));\n+    }\n+\n+    private RepositoryPage\u003cToDoItem\u003e toPage(\n+            FeedResponse\u003cToDoItem\u003e response,\n+            String category,\n+            int pageNumber) {\n+        LOGGER.info(\n+                \"query category={} page={} items={} requestCharge={} RU\",\n+                category,\n+                pageNumber,\n+                response.getResults().size(),\n+                response.getRequestCharge());\n+        return new RepositoryPage\u003c\u003e(\n+                response.getResults(),\n+                response.getRequestCharge(),\n+                response.getContinuationToken());\n+    }\n+\n+    private RepositoryResult\u003cToDoItem\u003e itemResult(\n+            String operation,\n+            CosmosItemResponse\u003cToDoItem\u003e response) {\n+        LOGGER.info(\n+                \"{} id={} requestCharge={} RU\",\n+                operation,\n+                response.getItem().getId(),\n+                response.getRequestCharge());\n+        return new RepositoryResult\u003c\u003e(response.getItem(), response.getRequestCharge());\n+    }\n+}\n*** Add File: src/main/java/com/example/todo/CosmosToDoFactory.java\n+package com.example.todo;\n+\n+import com.azure.cosmos.ConsistencyLevel;\n+import com.azure.cosmos.CosmosAsyncClient;\n+import com.azure.cosmos.CosmosClient;\n+import com.azure.cosmos.CosmosClientBuilder;\n+import com.azure.cosmos.models.CosmosContainerProperties;\n+import com.azure.cosmos.models.ExcludedPath;\n+import com.azure.cosmos.models.IndexingMode;\n+import com.azure.cosmos.models.IndexingPolicy;\n+import com.azure.identity.DefaultAzureCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Objects;\n+\n+public final class CosmosToDoFactory implements AutoCloseable {\n+    public static final String ENDPOINT_ENVIRONMENT_VARIABLE = \"COSMOS_ENDPOINT\";\n+    public static final int DEFAULT_TTL_SECONDS = (int) Duration.ofDays(90).toSeconds();\n+\n+    private final CosmosClient syncClient;\n+    private final CosmosAsyncClient asyncClient;\n+    private final String databaseName;\n+    private final String containerName;\n+\n+    private CosmosToDoFactory(\n+            CosmosClient syncClient,\n+            CosmosAsyncClient asyncClient,\n+            String databaseName,\n+            String containerName) {\n+        this.syncClient = syncClient;\n+        this.asyncClient = asyncClient;\n+        this.databaseName = databaseName;\n+        this.containerName = containerName;\n+    }\n+\n+    public static CosmosToDoFactory create(String databaseName, String containerName) {\n+        String endpoint = requireEnvironmentVariable(ENDPOINT_ENVIRONMENT_VARIABLE);\n+        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();\n+\n+        CosmosClient syncClient = clientBuilder(endpoint, credential).buildClient();\n+        CosmosAsyncClient asyncClient = null;\n+        try {\n+            initializeSchema(syncClient, databaseName, containerName);\n+            asyncClient = clientBuilder(endpoint, credential).buildAsyncClient();\n+            return new CosmosToDoFactory(\n+                    syncClient, asyncClient, databaseName, containerName);\n+        } catch (RuntimeException exception) {\n+            if (asyncClient != null) {\n+                asyncClient.close();\n+            }\n+            syncClient.close();\n+            throw exception;\n+        }\n+    }\n+\n+    public SyncToDoRepository syncRepository() {\n+        return new SyncToDoRepository(\n+                syncClient.getDatabase(databaseName).getContainer(containerName));\n+    }\n+\n+    public AsyncToDoRepository asyncRepository() {\n+        return new AsyncToDoRepository(\n+                asyncClient.getDatabase(databaseName).getContainer(containerName));\n+    }\n+\n+    @Override\n+    public void close() {\n+        asyncClient.close();\n+        syncClient.close();\n+    }\n+\n+    private static CosmosClientBuilder clientBuilder(\n+            String endpoint,\n+            DefaultAzureCredential credential) {\n+        return new CosmosClientBuilder()\n+                .endpoint(endpoint)\n+                .credential(credential)\n+                .consistencyLevel(ConsistencyLevel.SESSION);\n+    }\n+\n+    private static void initializeSchema(\n+            CosmosClient client,\n+            String databaseName,\n+            String containerName) {\n+        Objects.requireNonNull(databaseName, \"databaseName\");\n+        Objects.requireNonNull(containerName, \"containerName\");\n+\n+        client.createDatabaseIfNotExists(databaseName);\n+\n+        IndexingPolicy indexingPolicy = new IndexingPolicy()\n+                .setAutomatic(true)\n+                .setIndexingMode(IndexingMode.CONSISTENT)\n+                .setExcludedPaths(List.of(new ExcludedPath(\"/description/?\")));\n+        CosmosContainerProperties properties = new CosmosContainerProperties(\n+                containerName, \"/category\");\n+        properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);\n+        properties.setIndexingPolicy(indexingPolicy);\n+\n+        client.getDatabase(databaseName).createContainerIfNotExists(properties);\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\n+                    \"Environment variable \" + name + \" must contain the Cosmos DB endpoint.\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/todo/Main.java\n+package com.example.todo;\n+\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Instant;\n+import java.util.UUID;\n+\n+public final class Main {\n+    private static final String DATABASE_NAME =\n+            System.getenv().getOrDefault(\"COSMOS_DATABASE\", \"todo-demo\");\n+    private static final String CONTAINER_NAME =\n+            System.getenv().getOrDefault(\"COSMOS_CONTAINER\", \"items\");\n+    private static final int PAGE_SIZE = 2;\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        try (CosmosToDoFactory factory =\n+                     CosmosToDoFactory.create(DATABASE_NAME, CONTAINER_NAME)) {\n+            runSyncDemo(factory.syncRepository());\n+            runAsyncDemo(factory.asyncRepository()).block();\n+        }\n+    }\n+\n+    private static void runSyncDemo(SyncToDoRepository repository) {\n+        String category = \"sync-demo\";\n+        ToDoItem item = newItem(\"Sync repository demo\", category);\n+\n+        System.out.println(\"\\n=== Synchronous CRUD ===\");\n+        print(\"create\", repository.create(item));\n+\n+        RepositoryResult\u003cToDoItem\u003e read = repository.read(item.getId(), category);\n+        print(\"read\", read);\n+        read.value().setCompleted(true);\n+        read.value().setTitle(\"Sync repository demo - completed\");\n+        print(\"update\", repository.update(read.value()));\n+\n+        int pageNumber = 0;\n+        for (RepositoryPage\u003cToDoItem\u003e page\n+                : repository.queryByCategory(category, PAGE_SIZE)) {\n+            printPage(++pageNumber, page);\n+        }\n+\n+        print(\"delete\", repository.delete(item.getId(), category));\n+    }\n+\n+    private static Mono\u003cVoid\u003e runAsyncDemo(AsyncToDoRepository repository) {\n+        String category = \"async-demo\";\n+        ToDoItem item = newItem(\"Async repository demo\", category);\n+\n+        System.out.println(\"\\n=== Asynchronous CRUD ===\");\n+        return repository.create(item)\n+                .doOnNext(result -\u003e print(\"create\", result))\n+                .then(repository.read(item.getId(), category))\n+                .doOnNext(result -\u003e print(\"read\", result))\n+                .map(RepositoryResult::value)\n+                .flatMap(readItem -\u003e {\n+                    readItem.setCompleted(true);\n+                    readItem.setTitle(\"Async repository demo - completed\");\n+                    return repository.update(readItem);\n+                })\n+                .doOnNext(result -\u003e print(\"update\", result))\n+                .thenMany(repository.queryByCategory(category, PAGE_SIZE))\n+                .index()\n+                .doOnNext(indexed -\u003e printPage(\n+                        Math.toIntExact(indexed.getT1() + 1), indexed.getT2()))\n+                .then(repository.delete(item.getId(), category))\n+                .doOnNext(result -\u003e print(\"delete\", result))\n+                .then();\n+    }\n+\n+    private static ToDoItem newItem(String title, String category) {\n+        return new ToDoItem(\n+                UUID.randomUUID().toString(),\n+                title,\n+                \"This field is stored but deliberately excluded from indexing.\",\n+                false,\n+                Instant.now(),\n+                category);\n+    }\n+\n+    private static void print(String operation, RepositoryResult\u003c?\u003e result) {\n+        System.out.printf(\n+                \"%-6s RU=%6.2f result=%s%n\",\n+                operation,\n+                result.requestCharge(),\n+                result.value());\n+    }\n+\n+    private static void printPage(int pageNumber, RepositoryPage\u003cToDoItem\u003e page) {\n+        System.out.printf(\n+                \"query page=%d RU=%.2f items=%d continuationToken=%s%n\",\n+                pageNumber,\n+                page.requestCharge(),\n+                page.items().size(),\n+                page.continuationToken() == null ? \"\u003cend\u003e\" : \"\u003cpresent\u003e\");\n+        page.items().forEach(item -\u003e System.out.println(\"  \" + item));\n+    }\n+}\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"mvn -q -DskipTests package","description":"Compile and package Maven project","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `pom.xml`
- `src\main\java\com\example\todo\AsyncToDoRepository.java`
- `src\main\java\com\example\todo\ConcurrentUpdateException.java`
- `src\main\java\com\example\todo\CosmosToDoFactory.java`
- `src\main\java\com\example\todo\Main.java`
- `src\main\java\com\example\todo\RepositoryPage.java`
- `src\main\java\com\example\todo\RepositoryResult.java`
- `src\main\java\com\example\todo\SyncToDoRepository.java`
- `src\main\java\com\example\todo\ToDoItem.java`

## Copilot Response

I’ll inspect the workspace, then create the Maven structure and wire the Cosmos SDK APIs consistently across sync and async paths.

The workspace is empty. I’m creating a compact project with shared result/page types so RU charges and continuation metadata remain visible without buffering query results.

The implementation is in place; I’m compiling it now to catch any Cosmos SDK signature or serialization issues before finalizing.

Created a Java 17 Maven project implementing Azure Cosmos DB ToDo CRUD repositories.

Key behavior includes:

- Shared `ToDoItem` model with category partition key and ETag concurrency metadata.
- Synchronous CRUD and lazy paginated category queries.
- Reactive asynchronous CRUD and `Flux<RepositoryPage<ToDoItem>>` page streaming.
- Parameterized Cosmos SQL queries.
- RU logging and returned request charges for every operation/page.
- Optimistic concurrency using `If-Match`; stale updates throw `ConcurrentUpdateException`.
- Managed identity authentication through `DefaultAzureCredential`.
- Automatic database/container creation with:
  - 90-day default TTL.
  - `/category` partition key.
  - `/description/?` excluded from indexing.
- `Main` demonstrations for both repository implementations.

Run with:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
mvn exec:java
```

Optional database and container overrides use `COSMOS_DATABASE` and `COSMOS_CONTAINER`.

## Grader Results

- todo-repository.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (11/14)
      - Correct partition key usage: `/category` path, `PartitionKey` in all point operations: Pass
      - ETag-based optimistic concurrency: captures ETag from read, passes `ifMatchETag` on update: Pass
      - Handles 412 Precondition Failed as a specific error case for conflicts: Pass
      - Parameterized queries using `SqlQuerySpec` with `SqlParameter` (no string concatenation): Pass
      - Page-by-page iteration using `iterableByPage()` or `CosmosPagedIterable`: Pass
      - Configurable page size via `QueryRequestOptions.setMaxItemCount`: Fail
      - Logs continuation token and item count per page: Fail
      - Async query uses `CosmosPagedFlux` returning pages as a stream: Pass
      - TTL configured at 90 days (7776000 seconds) via `ContainerProperties.setDefaultTimeToLiveInSeconds()`: Pass
      - Indexing policy excludes `/description` path: Pass
      - RU cost extracted from response via `getRequestCharge()` and logged per operation: Pass
      - Catches `CosmosException` with status code checks (404, 409, 412): Fail
      - Handles 412 separately for ETag conflicts: Pass
      - Does NOT flatten query results (`.stream()` / `.forEach()` without page iteration): Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Pass
  - Client Builder Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Pass
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Fail (0/1)
      - grader executed: Fail
  - Async Uses Project Reactor (Mono/Flux) (prompt): Fail (0/1)
      - grader executed: Fail
  - Service-Specific Exception Handling (prompt): Fail (0/1)
      - grader executed: Fail
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Fail (0/1)
      - grader executed: Fail
  - Try-With-Resources for Clients (prompt): Fail (0/1)
      - grader executed: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 79% | 1.00 | 0.7857 | 11.6% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 14.7% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 14.7% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 14.7% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 14.7% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 14.7% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 14.7% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Try-With-Resources for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 13.00** | **Σ 6.7857** | **52.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-java-todo-repository --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
