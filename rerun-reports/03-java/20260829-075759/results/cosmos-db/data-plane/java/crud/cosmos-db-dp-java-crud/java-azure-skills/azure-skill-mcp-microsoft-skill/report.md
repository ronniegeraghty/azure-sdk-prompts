# Evaluation Report: cosmos-db-dp-java-crud

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 340.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-java-crud` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 15/19 |
| Duration | 340.3s |
| Timestamp | 2026-08-28T23:57:59Z |
| Files Generated | 3 |
| Event Count | 5621 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 108.3s |
| Review | 230.8s |
| **Total** | **340.3s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=289518 out=4567 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 880ms

**Input:**

```json
{"skill":"azure-cosmos-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Cosmos DB SDK for Java

Client library for Azure Cosmos DB NoSQL API with global distribution and reactive patterns.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-cosmos</artifactId>
    <version>LATEST</version>
</dependency>
```

Or use Azure SDK BOM:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>{bom_version}</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-cosmos</artifactId>
    </dependency>
</dependencies>
```

## Environment Variables

```bash
COSMOS_ENDPOINT=https://<account>.documents.azure.com:443/
COSMOS_KEY=<your-primary-key>
```

## Authentication

### Key-based Authentication

```java
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;

CosmosClient client = new CosmosClientBuilder()
    .endpoint(System.getenv("COSMOS_ENDPOINT"))
    .key(System.getenv("COSMOS_KEY"))
    .buildClient();
```

### Async Client

```java
import com.azure.cosmos.CosmosAsyncClient;

CosmosAsyncClient asyncClient = new CosmosClientBuilder()
    .endpoint(serviceEndpoint)
    .key(key)
    .buildAsyncClient();
```

### With Customizations

```java
import com.azure.cosmos.ConsistencyLevel;
import java.util.Arrays;

CosmosClient client = new CosmosClientBuilder()
    .endpoint(serviceEndpoint)
    .key(key)
    .directMode(directConnectionConfig, gatewayConnectionConfig)
    .consistencyLevel(ConsistencyLevel.SESSION)
    .connectionSharingAcrossClientsEnabled(true)
    .contentResponseOnWriteEnabled(true)
    .userAgentSuffix("my-application")
    .preferredRegions(Arrays.asList("West US", "East US"))
    .buildClient();
```

## Client Hierarc
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 2126ms

**Input:**

```json
{"intent":"Generate a local-only Java Maven program using the Azure Cosmos DB SDK for NoSQL CRUD, parameterized queries, partition keys, and CosmosException handling. Do not provision or connect to live Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. glob — 2491ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-cosmos-java","pattern":"references\\*.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. glob — 1743ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1203705856","pattern":"**\\*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 1701ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for a local Java Maven Azure Cosmos DB NoSQL CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. view — 2700ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-cosmos-java"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java/references/examples.md
@@ -1,410 +1,410 @@
 # Azure Cosmos DB Java SDK - Examples
 
 Comprehensive code examples for the Azure Cosmos DB SDK for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Database Operations](#database-operations)
 - [Container Operations](#container-operations)
 - [CRUD Operations (Sync)](#crud-operations-sync)
 - [CRUD Operations (Async)](#crud-operations-async)
 - [SQL Queries](#sql-queries)
 
 ---
 
 ## Maven Dependency
 
 ```xml
 <dependencyManagement>
     <dependencies>
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-sdk-bom</artifactId>
             <version>{bom_version}</version>
             <type>pom</type>
             <scope>import</scope>
         </dependency>
     </dependencies>
 </dependencyManagement>
 
 <dependencies>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-cosmos</artifactId>
     </dependency>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-identity</artifactId>
     </dependency>
 </dependencies>
 ```
 
 ---
 
 ## Client Creation
 
 ### Synchronous Client (CosmosClient)
 
 ```java
 import com.azure.cosmos.ConsistencyLevel;
 import com.azure.cosmos.CosmosClient;
 import com.azure.cosmos.CosmosClientBuilder;
 import java.util.Arrays;
 
 // Basic client w
... (truncated)
```

### ✅ 7. view — 2235ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-cosmos-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-cosmos-java
@@ -1,0 +1,2 @@
 SKILL.md
 references

```

### ✅ 8. web_search (via github-mcp-server) — 10517ms

**Input:**

```json
{"query":"What is the latest stable com.azure azure-cosmos Maven artifact version as of August 2026? Use authoritative Maven Central or Microsoft sources."}
```

**Output:**

```
{"type":"output_text","text":{"value":"The latest stable version of the `com.azure:azure-cosmos` Maven artifact as of August 2026 is **4.82.0**, published on August 26, 2026. This information is confirmed by data from Maven Central, which is an authoritative source for Java artifacts【3:1†source】.","annotations":[{"text":"","start_index":245,"end_index":257,"url_citation":{"title":"Maven Central: com.azure:azure-cosmos","url":"https://central.sonatype.com/artifact/com.azure/azure-cosmos/versions"}}]},"bing_searches":[{"text":"latest stable version of com.azure:azure-cosmos maven artifact as of August 2026","url":"https://www.bing.com/search?q=latest+stable+version+of+com.azure%3Aazure-cosmos+maven+artifact+as+of+August+2026"}],"annotations":null}
```

### ✅ 9. powershell — 10042ms

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace contents"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 10. apply_patch — 863ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ecosmos-crud-sample\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+            \u003cversion\u003e4.82.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.cosmos.CosmosCrudExample\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/cosmos/CosmosCrudExample.java\n+package com.example.cosmos;\n+\n+import com.azure.cosmos.CosmosClient;\n+import com.azure.cosmos.CosmosClientBuilder;\n+import com.azure.cosmos.CosmosContainer;\n+import com.azure.cosmos.CosmosDatabase;\n+import com.azure.cosmos.CosmosException;\n+import com.azure.cosmos.models.CosmosContainerProperties;\n+import com.azure.cosmos.models.CosmosItemRequestOptions;\n+import com.azure.cosmos.models.CosmosItemResponse;\n+import com.azure.cosmos.models.CosmosQueryRequestOptions;\n+import com.azure.cosmos.models.PartitionKey;\n+import com.azure.cosmos.models.SqlParameter;\n+import com.azure.cosmos.models.SqlQuerySpec;\n+import com.azure.cosmos.util.CosmosPagedIterable;\n+\n+import java.net.URI;\n+import java.util.Collections;\n+import java.util.Locale;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+\n+public final class CosmosCrudExample {\n+    private static final Logger LOGGER = Logger.getLogger(CosmosCrudExample.class.getName());\n+\n+    private static final String DATABASE_ID = \"TestDB\";\n+    private static final String CONTAINER_ID = \"Items\";\n+    private static final String PARTITION_KEY_PATH = \"/category\";\n+    private static final String CATEGORY = \"electronics\";\n+\n+    private CosmosCrudExample() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+        String key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+        requireLocalEmulatorEndpoint(endpoint);\n+\n+        try (CosmosClient client = new CosmosClientBuilder()\n+            .endpoint(endpoint)\n+            .key(key)\n+            .contentResponseOnWriteEnabled(true)\n+            .buildClient()) {\n+\n+            CosmosDatabase database = createDatabase(client);\n+            CosmosContainer container = createContainer(database);\n+            runCrudOperations(container);\n+        } catch (CosmosException exception) {\n+            logCosmosException(exception);\n+            System.exit(1);\n+        }\n+    }\n+\n+    private static CosmosDatabase createDatabase(CosmosClient client) {\n+        client.createDatabaseIfNotExists(DATABASE_ID);\n+        LOGGER.info(() -\u003e \"Database is ready: \" + DATABASE_ID);\n+        return client.getDatabase(DATABASE_ID);\n+    }\n+\n+    private static CosmosContainer createContainer(CosmosDatabase database) {\n+        CosmosContainerProperties properties =\n+            new CosmosContainerProperties(CONTAINER_ID, PARTITION_KEY_PATH);\n+        database.createContainerIfNotExists(properties);\n+        LOGGER.info(() -\u003e \"Container is ready: \" + CONTAINER_ID);\n+        return database.getContainer(CONTAINER_ID);\n+    }\n+\n+    private static void runCrudOperations(CosmosContainer container) {\n+        Item item = new Item(\"item-001\", CATEGORY, \"Wireless Headphones\", 10);\n+        PartitionKey partitionKey = new PartitionKey(item.getCategory());\n+\n+        CosmosItemResponse\u003cItem\u003e createResponse = container.createItem(\n+            item,\n+            partitionKey,\n+            new CosmosItemRequestOptions());\n+        logOperation(\"Created\", createResponse);\n+\n+        CosmosItemResponse\u003cItem\u003e readResponse =\n+            container.readItem(item.getId(), partitionKey, Item.class);\n+        Item storedItem = readResponse.getItem();\n+        logOperation(\"Read\", readResponse);\n+\n+        SqlQuerySpec query = new SqlQuerySpec(\n+            \"SELECT * FROM c WHERE c.category = @category\",\n+            Collections.singletonList(new SqlParameter(\"@category\", CATEGORY)));\n+        CosmosQueryRequestOptions queryOptions = new CosmosQueryRequestOptions()\n+            .setPartitionKey(partitionKey);\n+        CosmosPagedIterable\u003cItem\u003e queryResults =\n+            container.queryItems(query, queryOptions, Item.class);\n+\n+        queryResults.forEach(result -\u003e\n+            LOGGER.info(() -\u003e String.format(\n+                Locale.ROOT,\n+                \"Query result: id=%s, name=%s, quantity=%d\",\n+                result.getId(),\n+                result.getName(),\n+                result.getQuantity())));\n+\n+        storedItem.setQuantity(25);\n+        CosmosItemResponse\u003cItem\u003e replaceResponse = container.replaceItem(\n+            storedItem,\n+            storedItem.getId(),\n+            partitionKey,\n+            new CosmosItemRequestOptions());\n+        logOperation(\"Replaced\", replaceResponse);\n+\n+        container.deleteItem(\n+            storedItem.getId(),\n+            partitionKey,\n+            new CosmosItemRequestOptions());\n+        LOGGER.info(() -\u003e \"Deleted item: \" + storedItem.getId());\n+    }\n+\n+    private static void logOperation(String operation, CosmosItemResponse\u003cItem\u003e response) {\n+        LOGGER.info(() -\u003e String.format(\n+            Locale.ROOT,\n+            \"%s item %s (status=%d, requestCharge=%.2f RU)\",\n+            operation,\n+            response.getItem().getId(),\n+            response.getStatusCode(),\n+            response.getRequestCharge()));\n+    }\n+\n+    private static void logCosmosException(CosmosException exception) {\n+        LOGGER.log(\n+            Level.SEVERE,\n+            String.format(\n+                Locale.ROOT,\n+                \"Cosmos DB operation failed: status=%d, subStatus=%d, activityId=%s, \"\n+                    + \"requestCharge=%.2f RU, retryAfter=%s\",\n+                exception.getStatusCode(),\n+                exception.getSubStatusCode(),\n+                exception.getActivityId(),\n+                exception.getRequestCharge(),\n+                exception.getRetryAfterDuration()),\n+            exception);\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Environment variable \" + name + \" is required.\");\n+        }\n+        return value;\n+    }\n+\n+    private static void requireLocalEmulatorEndpoint(String endpoint) {\n+        URI endpointUri;\n+        try {\n+            endpointUri = URI.create(endpoint);\n+        } catch (IllegalArgumentException exception) {\n+            throw new IllegalStateException(\"COSMOS_ENDPOINT must be a valid URI.\", exception);\n+        }\n+\n+        String host = endpointUri.getHost();\n+        if (host == null\n+            || !(host.equalsIgnoreCase(\"localhost\")\n+            || host.equals(\"127.0.0.1\")\n+            || host.equals(\"::1\"))) {\n+            throw new IllegalStateException(\n+                \"This sample is restricted to a local Cosmos DB emulator endpoint.\");\n+        }\n+    }\n+\n+    public static final class Item {\n+        private String id;\n+        private String category;\n+        private String name;\n+        private int quantity;\n+\n+        public Item() {\n+        }\n+\n+        public Item(String id, String category, String name, int quantity) {\n+            this.id = id;\n+            this.category = category;\n+            this.name = name;\n+            this.quantity = quantity;\n+        }\n+\n+        public String getId() {\n+            return id;\n+        }\n+\n+        public void setId(String id) {\n+            this.id = id;\n+        }\n+\n+        public String getCategory() {\n+            return category;\n+        }\n+\n+        public void setCategory(String category) {\n+            this.category = category;\n+        }\n+\n+        public String getName() {\n+            return name;\n+        }\n+\n+        public void setName(String name) {\n+            this.name = name;\n+        }\n+\n+        public int getQuantity() {\n+            return quantity;\n+        }\n+\n+        public void setQuantity(int quantity) {\n+            this.quantity = quantity;\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Cosmos DB Java CRUD sample\n+\n+This Maven project demonstrates synchronous CRUD operations against a local Azure Cosmos DB\n+NoSQL emulator. It creates `TestDB`, creates the `Items` container with `/category` as its\n+partition key, and then creates, reads, queries, replaces, and deletes an item.\n+\n+## Configuration\n+\n+Set the emulator endpoint and key without placing credentials in source control:\n+\n+```powershell\n+$env:COSMOS_ENDPOINT = \"https://localhost:8081\"\n+$env:COSMOS_KEY = \"\u003clocal-emulator-key\u003e\"\n+```\n+\n+The program rejects non-loopback endpoints so that it cannot modify a live Azure account.\n+\n+## Build and run\n+\n+```powershell\n+mvn compile\n+mvn exec:java\n+```\n+\n+The required Azure SDK dependency is:\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-cosmos\u003c/artifactId\u003e\n+    \u003cversion\u003e4.82.0\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+References:\n+\n+- [Azure Cosmos DB Java SDK examples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)\n+- [Azure Cosmos DB Java SDK API documentation](https://azuresdkdocs.z19.web.core.windows.net/java/azure-cosmos/latest/index.html)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1203705856/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1203705856/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1203705856/pom.xml
@@ -1,0 +1,42 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>cosmos-crud-sample</artifactId>
+    <version>1.0.0</version>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-cosmos</artifactId>
+            <version>4.82.0</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-compiler-plugin</artifactId>
+                <version>3.14.1</version>
+            </plugin>
+            <plugin>
+                <groupId>org.codehaus.mojo</groupId>
+                <artifactId>exec-maven-plugin</artifactId>
+                <version>3.5.1</version>
+                <configuration>
+                    <mainClass>com.example.cosmos.CosmosCrudExample</mainClass>
+                </configuration>
+            </plugin>
+        </plugins>
+    </build>
+</project>
+



diff --git a/C:/Users/weidxu/AppD
... (truncated)
```

### ✅ 11. powershell — 13083ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile Cosmos CRUD sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] -------------------< com.example:cosmos-crud-sample >-------------------
[INFO] Building cosmos-crud-sample 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ cosmos-crud-sample ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-java-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-1203705856\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ cosmos-crud-sample ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  3.522 s
[INFO] Finished at: 2026-08-29T07:59:42+08:00
[INFO] ------------------------------------------------------------------------
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\cosmos\CosmosCrudExample.java`

## Copilot Response

I’ll scaffold a local Maven example using environment variables only, then compile it without contacting Azure.

The workspace is empty. I’m adding a complete Maven project with a synchronous client, point read, partition-scoped parameterized query, replace/delete operations, and emulator-only endpoint protection.

The implementation is in place; I’m compiling the Maven project now to catch SDK signature or Java type issues.

Implemented a complete Java 17 Maven sample:

- `src\main\java\com\example\cosmos\CosmosCrudExample.java` — database/container creation, POJO CRUD, partition-key point read, parameterized query, and detailed `CosmosException` logging.
- `pom.xml` — uses `com.azure:azure-cosmos:4.82.0`.
- `README.md` — emulator configuration and execution instructions.

The program reads `COSMOS_ENDPOINT` and `COSMOS_KEY` from environment variables and restricts execution to loopback Cosmos DB emulator endpoints, preventing accidental changes to live Azure resources. References: [Azure Cosmos DB Java samples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples) and [SDK documentation](https://azuresdkdocs.z19.web.core.windows.net/java/azure-cosmos/latest/index.html).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses try-with-resources correctly for the CosmosClient lifecycle.

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 9.0000** | **69.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-java-crud --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
