# Evaluation Report: service-bus-dp-java-crud

**Config:** java-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 300.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-crud` |
| Config | java-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 14/19 |
| Duration | 300.6s |
| Timestamp | 2026-08-29T03:21:49Z |
| Files Generated | 2 |
| Event Count | 4820 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 56.2s |
| Review | 243.9s |
| **Total** | **300.6s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-skills/baseline

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=132165 out=2869 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. glob — 7127ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 6718ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248","pattern":"**/pom.xml"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. apply_patch — 853ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-service-bus-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0-SNAPSHOT\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e7.17.15\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.ServiceBusDemo\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/ServiceBusDemo.java\n+package com.example;\n+\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.azure.messaging.servicebus.models.ServiceBusErrorContext;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+import com.azure.messaging.servicebus.models.ServiceBusReceivedMessageContext;\n+\n+import java.time.Duration;\n+import java.util.concurrent.CountDownLatch;\n+import java.util.concurrent.TimeUnit;\n+\n+public final class ServiceBusDemo {\n+    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(10);\n+\n+    private ServiceBusDemo() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String connectionString = requiredEnvironmentVariable(\"SERVICE_BUS_CONNECTION_STRING\");\n+        String queueName = requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+        String topicName = requiredEnvironmentVariable(\"SERVICE_BUS_TOPIC_NAME\");\n+        String subscriptionName = requiredEnvironmentVariable(\"SERVICE_BUS_SUBSCRIPTION_NAME\");\n+\n+        sendToQueue(connectionString, queueName);\n+        receiveFromQueue(connectionString, queueName);\n+        processQueueContinuously(connectionString, queueName);\n+        sendToTopicAndReceiveFromSubscription(\n+            connectionString, topicName, subscriptionName);\n+    }\n+\n+    private static void sendToQueue(String connectionString, String queueName) {\n+        ServiceBusSenderClient sender = new ServiceBusClientBuilder()\n+            .connectionString(connectionString)\n+            .sender()\n+            .queueName(queueName)\n+            .buildClient();\n+\n+        try {\n+            sender.sendMessage(new ServiceBusMessage(\"Single queue message\"));\n+\n+            ServiceBusMessageBatch batch = sender.createMessageBatch();\n+            for (int i = 1; i \u003c= 5; i++) {\n+                ServiceBusMessage message = new ServiceBusMessage(\"Batch message \" + i);\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalStateException(\n+                        \"The five demo messages do not fit in one Service Bus batch.\");\n+                }\n+            }\n+            sender.sendMessages(batch);\n+            System.out.println(\"Sent one message and a batch of five messages to \" + queueName);\n+        } finally {\n+            sender.close();\n+        }\n+    }\n+\n+    private static void receiveFromQueue(String connectionString, String queueName) {\n+        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()\n+            .connectionString(connectionString)\n+            .receiver()\n+            .queueName(queueName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .buildClient();\n+\n+        try {\n+            for (ServiceBusReceivedMessage message\n+                : receiver.receiveMessages(6, RECEIVE_TIMEOUT)) {\n+                System.out.printf(\n+                    \"Queue receiver processed %s: %s%n\",\n+                    message.getMessageId(),\n+                    message.getBody().toString());\n+                receiver.complete(message);\n+            }\n+        } finally {\n+            receiver.close();\n+        }\n+    }\n+\n+    private static void processQueueContinuously(\n+        String connectionString, String queueName) throws InterruptedException {\n+\n+        CountDownLatch processed = new CountDownLatch(1);\n+        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()\n+            .connectionString(connectionString)\n+            .processor()\n+            .queueName(queueName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .processMessage(context -\u003e processMessage(context, processed))\n+            .processError(ServiceBusDemo::processError)\n+            .buildProcessorClient();\n+\n+        ServiceBusSenderClient sender = new ServiceBusClientBuilder()\n+            .connectionString(connectionString)\n+            .sender()\n+            .queueName(queueName)\n+            .buildClient();\n+\n+        try {\n+            processor.start();\n+            sender.sendMessage(new ServiceBusMessage(\"Message for the continuous processor\"));\n+\n+            if (!processed.await(30, TimeUnit.SECONDS)) {\n+                throw new IllegalStateException(\n+                    \"The processor did not receive its demo message within 30 seconds.\");\n+            }\n+        } finally {\n+            sender.close();\n+            processor.close();\n+        }\n+    }\n+\n+    private static void processMessage(\n+        ServiceBusReceivedMessageContext context, CountDownLatch processed) {\n+\n+        ServiceBusReceivedMessage message = context.getMessage();\n+        System.out.printf(\n+            \"Processor handled %s: %s%n\",\n+            message.getMessageId(),\n+            message.getBody().toString());\n+        context.complete();\n+        processed.countDown();\n+    }\n+\n+    private static void processError(ServiceBusErrorContext context) {\n+        System.err.printf(\n+            \"Processor error in %s for entity %s: %s%n\",\n+            context.getErrorSource(),\n+            context.getEntityPath(),\n+            context.getException().getMessage());\n+    }\n+\n+    private static void sendToTopicAndReceiveFromSubscription(\n+        String connectionString, String topicName, String subscriptionName) {\n+\n+        ServiceBusSenderClient topicSender = new ServiceBusClientBuilder()\n+            .connectionString(connectionString)\n+            .sender()\n+            .topicName(topicName)\n+            .buildClient();\n+\n+        ServiceBusReceiverClient subscriptionReceiver = new ServiceBusClientBuilder()\n+            .connectionString(connectionString)\n+            .receiver()\n+            .topicName(topicName)\n+            .subscriptionName(subscriptionName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .buildClient();\n+\n+        try {\n+            topicSender.sendMessage(new ServiceBusMessage(\"Topic message\"));\n+\n+            for (ServiceBusReceivedMessage message\n+                : subscriptionReceiver.receiveMessages(1, RECEIVE_TIMEOUT)) {\n+                System.out.printf(\n+                    \"Subscription receiver processed %s: %s%n\",\n+                    message.getMessageId(),\n+                    message.getBody().toString());\n+                subscriptionReceiver.complete(message);\n+            }\n+        } finally {\n+            subscriptionReceiver.close();\n+            topicSender.close();\n+        }\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Set the \" + name + \" environment variable.\");\n+        }\n+        return value;\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/pom.xml
@@ -1,0 +1,37 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>azure-service-bus-demo</artifactId>
+    <version>1.0-SNAPSHOT</version>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-messaging-servicebus</artifactId>
+            <version>7.17.15</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.codehaus.mojo</groupId>
+                <artifactId>exec-maven-plugin</artifactId>
+                <version>3.5.0</version>
+                <configuration>
+                    <mainClass>com.example.ServiceBusDemo</mainClass>
+                </configuration>
+            </plugin>
+        </plugins>
+    </build>
+</project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/exampl
... (truncated)
```

### ✅ 4. powershell — 5243ms

**Input:**

```json
{"command":"mvn --quiet -DskipTests compile","description":"Compile the Maven Service Bus example","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[ERROR] COMPILATION ERROR : 
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java:[10,45] cannot find symbol
  symbol:   class ServiceBusErrorContext
  location: package com.azure.messaging.servicebus.models
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java:[12,45] cannot find symbol
  symbol:   class ServiceBusReceivedMessageContext
  location: package com.azure.messaging.servicebus.models
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java:[118,9] cannot find symbol
  symbol:   class ServiceBusReceivedMessageContext
  location: class com.example.ServiceBusDemo
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java:[129,38] cannot find symbol
  symbol:   class ServiceBusErrorContext
  location: class com.example.ServiceBusDemo
[ERROR] Failed to execute goal org.apache.maven.plugins:maven-compiler-plugin:3.13.0:compile (default-compile) on project azure-service-bus-demo: Compilation failure: Compilation failure: 
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java:[10,45] cannot find symbol
[ERROR]   symbol:   
... (truncated)
```

### ✅ 5. apply_patch — 907ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/ServiceBusDemo.java\n@@\n import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusErrorContext;\n import com.azure.messaging.servicebus.ServiceBusMessage;\n import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;\n import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n import com.azure.messaging.servicebus.ServiceBusSenderClient;\n-import com.azure.messaging.servicebus.models.ServiceBusErrorContext;\n import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n-import com.azure.messaging.servicebus.models.ServiceBusReceivedMessageContext;\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-java-crud-java-azure-skills-baseline-3598486248/src/main/java/com/example/ServiceBusDemo.java
@@ -1,15 +1,15 @@
 package com.example;
 
 import com.azure.messaging.servicebus.ServiceBusClientBuilder;
+import com.azure.messaging.servicebus.ServiceBusErrorContext;
 import com.azure.messaging.servicebus.ServiceBusMessage;
 import com.azure.messaging.servicebus.ServiceBusMessageBatch;
 import com.azure.messaging.servicebus.ServiceBusProcessorClient;
 import com.azure.messaging.servicebus.ServiceBusReceivedMessage;
+import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;
 import com.azure.messaging.servicebus.ServiceBusReceiverClient;
 import com.azure.messaging.servicebus.ServiceBusSenderClient;
-import com.azure.messaging.servicebus.models.ServiceBusErrorContext;
 import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;
-import com.azure.messaging.servicebus.models.ServiceBusReceivedMessageContext;
 
 import java.time.Duration;
 import java.util.concurrent.CountDownLatch;

```

### ✅ 6. powershell — 5258ms

**Input:**

```json
{"command":"mvn --quiet -DskipTests compile","description":"Recompile the corrected Maven example","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `pom.xml`
- `src\main\java\com\example\ServiceBusDemo.java`

## Copilot Response

I’ll inspect the workspace, then create a runnable Maven example covering queue, processor, and topic/subscription flows without provisioning Azure resources.

The workspace is empty. I’ll add a minimal Maven project whose entities and connection string come from environment variables, so it demonstrates SDK usage without creating or modifying Azure resources.

The example is complete; I’ll compile it locally to catch SDK signature or Maven configuration issues, without connecting to Azure.

Created a complete Maven example:

- `pom.xml` includes `com.azure:azure-messaging-servicebus:7.17.15`
- `src\main\java\com\example\ServiceBusDemo.java` demonstrates:
  - Queue sender creation
  - Single-message and five-message batch sending
  - Synchronous receiving with `receiveMessages()`
  - Peek-lock settlement with `complete()`
  - Continuous processing with `ServiceBusProcessorClient` handlers
  - Topic sending and subscription receiving
  - Explicit `close()` cleanup for all clients
  - Environment-based configuration without provisioning Azure resources

Run after configuring existing Service Bus entities:

```powershell
$env:SERVICE_BUS_CONNECTION_STRING = "<connection-string>"
$env:SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"

mvn exec:java
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- All Azure Service Bus client instances have explicit close() cleanup in finally blocks.

## Grader Results

- send-receive-messages.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `azure-messaging-servicebus` Maven dependency: Pass
      - `ServiceBusClientBuilder` with connection string: Pass
      - `ServiceBusSenderClient` and `ServiceBusMessage`: Pass
      - `createMessageBatch()` and `tryAddMessage()`: Pass
      - `ServiceBusReceiverClient.receiveMessages()` and `complete()`: Pass
      - `ServiceBusProcessorClient` with `processMessage` and `processError` handlers: Pass
      - Topic operations with `.topicName()` and `.subscriptionName()` on the builder: Pass
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
  - Service-Specific Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Fail
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 8.0000** | **61.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-java-crud --config java-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
