# Evaluation Report: service-bus-dp-java-crud

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 374.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-crud` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 14/19 |
| Duration | 374.8s |
| Timestamp | 2026-08-26T13:48:34Z |
| Files Generated | 3 |
| Event Count | 4362 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 97.0s |
| Review | 276.3s |
| **Total** | **374.8s** |

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
| Token Usage | in=106284 out=3013 |
| Turn Count | 5 |

## Prompt Sent

```
Write a Java program that demonstrates
messaging with Azure Service Bus:
1. Create a ServiceBusSenderClient using ServiceBusClientBuilder for a queue
2. Send a single message with ServiceBusMessage
3. Send a batch of 5 messages using ServiceBusMessageBatch
4. Create a ServiceBusReceiverClient and receive messages using receiveMessages()
5. Complete a message with complete() after processing
6. Create a ServiceBusProcessorClient for continuous processing with handlers
7. Demonstrate sending to a topic and receiving from a subscription

Show required Maven dependency (com.azure:azure-messaging-servicebus) and
proper resource cleanup with close().
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-crud-java-azure-tools-baseline-3271477279","pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-crud-java-azure-tools-baseline-3271477279","pattern":"**/pom.xml"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-crud-java-azure-tools-baseline-3271477279","pattern":"**/*.java"}
```

### 🔧 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Maven Central latest stable version com.azure azure-messaging-servicebus August 2026"}
```

### 🔧 5. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-service-bus-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0-SNAPSHOT\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e11\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e7.17.17\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.ServiceBusDemo\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/ServiceBusDemo.java\n+package com.example;\n+\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+\n+import java.time.Duration;\n+import java.util.concurrent.CountDownLatch;\n+import java.util.concurrent.TimeUnit;\n+\n+public final class ServiceBusDemo {\n+    private static final String CONNECTION_STRING =\n+            requiredEnvironmentVariable(\"SERVICE_BUS_CONNECTION_STRING\");\n+    private static final String QUEUE_NAME =\n+            requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+    private static final String TOPIC_NAME =\n+            requiredEnvironmentVariable(\"SERVICE_BUS_TOPIC_NAME\");\n+    private static final String SUBSCRIPTION_NAME =\n+            requiredEnvironmentVariable(\"SERVICE_BUS_SUBSCRIPTION_NAME\");\n+\n+    private ServiceBusDemo() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        sendQueueMessages();\n+        receiveAndCompleteQueueMessages();\n+        processQueueContinuously();\n+        sendToTopicAndReceiveFromSubscription();\n+    }\n+\n+    private static void sendQueueMessages() {\n+        ServiceBusSenderClient sender = new ServiceBusClientBuilder()\n+                .connectionString(CONNECTION_STRING)\n+                .sender()\n+                .queueName(QUEUE_NAME)\n+                .buildClient();\n+\n+        try {\n+            sender.sendMessage(new ServiceBusMessage(\"Single queue message\"));\n+            System.out.println(\"Sent one message to queue \" + QUEUE_NAME);\n+\n+            ServiceBusMessageBatch batch = sender.createMessageBatch();\n+            for (int i = 1; i \u003c= 5; i++) {\n+                ServiceBusMessage message = new ServiceBusMessage(\"Batch message \" + i);\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalStateException(\n+                            \"The batch cannot fit all five messages. Send multiple batches instead.\");\n+                }\n+            }\n+\n+            sender.sendMessages(batch);\n+            System.out.println(\"Sent a batch of five messages to queue \" + QUEUE_NAME);\n+        } finally {\n+            sender.close();\n+        }\n+    }\n+\n+    private static void receiveAndCompleteQueueMessages() {\n+        ServiceBusReceiverClient receiver = new ServiceBusClientBuilder()\n+                .connectionString(CONNECTION_STRING)\n+                .receiver()\n+                .queueName(QUEUE_NAME)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .buildClient();\n+\n+        try {\n+            receiver.receiveMessages(6, Duration.ofSeconds(20)).forEach(message -\u003e {\n+                System.out.printf(\"Received queue message: %s%n\", message.getBody());\n+\n+                // PEEK_LOCK messages remain on the queue until explicitly settled.\n+                receiver.complete(message);\n+                System.out.printf(\"Completed message %s%n\", message.getMessageId());\n+            });\n+        } finally {\n+            receiver.close();\n+        }\n+    }\n+\n+    private static void processQueueContinuously() throws InterruptedException {\n+        CountDownLatch processed = new CountDownLatch(1);\n+\n+        ServiceBusProcessorClient processor = new ServiceBusClientBuilder()\n+                .connectionString(CONNECTION_STRING)\n+                .processor()\n+                .queueName(QUEUE_NAME)\n+                .disableAutoComplete()\n+                .processMessage(context -\u003e {\n+                    System.out.printf(\"Processor received: %s%n\",\n+                            context.getMessage().getBody());\n+                    context.complete();\n+                    processed.countDown();\n+                })\n+                .processError(context -\u003e {\n+                    System.err.printf(\"Processor error from %s: %s%n\",\n+                            context.getErrorSource(), context.getException().getMessage());\n+                })\n+                .buildProcessorClient();\n+\n+        ServiceBusSenderClient sender = new ServiceBusClientBuilder()\n+                .connectionString(CONNECTION_STRING)\n+                .sender()\n+                .queueName(QUEUE_NAME)\n+                .buildClient();\n+\n+        try {\n+            processor.start();\n+            sender.sendMessage(new ServiceBusMessage(\"Message for continuous processor\"));\n+\n+            if (!processed.await(30, TimeUnit.SECONDS)) {\n+                throw new IllegalStateException(\n+                        \"The processor did not receive its demonstration message within 30 seconds.\");\n+            }\n+        } finally {\n+            sender.close();\n+            processor.close();\n+        }\n+    }\n+\n+    private static void sendToTopicAndReceiveFromSubscription() {\n+        ServiceBusSenderClient topicSender = new ServiceBusClientBuilder()\n+                .connectionString(CONNECTION_STRING)\n+                .sender()\n+                .topicName(TOPIC_NAME)\n+                .buildClient();\n+\n+        ServiceBusReceiverClient subscriptionReceiver = new ServiceBusClientBuilder()\n+                .connectionString(CONNECTION_STRING)\n+                .receiver()\n+                .topicName(TOPIC_NAME)\n+                .subscriptionName(SUBSCRIPTION_NAME)\n+                .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+                .buildClient();\n+\n+        try {\n+            topicSender.sendMessage(new ServiceBusMessage(\"Topic message\"));\n+            System.out.printf(\"Sent one message to topic %s%n\", TOPIC_NAME);\n+\n+            subscriptionReceiver.receiveMessages(1, Duration.ofSeconds(20))\n+                    .forEach(message -\u003e {\n+                        System.out.printf(\"Subscription %s received: %s%n\",\n+                                SUBSCRIPTION_NAME, message.getBody());\n+                        subscriptionReceiver.complete(message);\n+                    });\n+        } finally {\n+            subscriptionReceiver.close();\n+            topicSender.close();\n+        }\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Set the \" + name + \" environment variable.\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Azure Service Bus Java demo\n+\n+This example uses existing Azure Service Bus entities; it does not create or\n+deploy any Azure resources.\n+\n+## Configuration\n+\n+Set these environment variables:\n+\n+```text\n+SERVICE_BUS_CONNECTION_STRING=\u003cservice-bus-connection-string\u003e\n+SERVICE_BUS_QUEUE_NAME=\u003cexisting-queue-name\u003e\n+SERVICE_BUS_TOPIC_NAME=\u003cexisting-topic-name\u003e\n+SERVICE_BUS_SUBSCRIPTION_NAME=\u003cexisting-subscription-name\u003e\n+```\n+\n+The connection string must permit sending and receiving for the named entities.\n+The subscription must already belong to the configured topic.\n+\n+## Run\n+\n+```text\n+mvn compile exec:java\n+```\n+\n+The program sends one queue message, sends a five-message batch, receives and\n+completes queue messages, runs a continuous processor, and demonstrates topic\n+and subscription messaging. Every sender, receiver, and processor client is\n+closed in a `finally` block.\n*** End Patch\n"
```

### 🔧 6. powershell

**Input:**

```json
{"command":"mvn --quiet compile","description":"Compile the Maven Service Bus example","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\ServiceBusDemo.java`

## Copilot Response

I’ll inspect the workspace, then create a runnable Maven example with queue, processor, and topic/subscription flows.

The workspace is empty. I’ll add a minimal Maven project using environment-provided entity names and connection string, without provisioning Azure resources.

The example is implemented. I’m compiling it now to catch any SDK API or Maven configuration issues.

Implemented a complete Maven-based Azure Service Bus example.

- `pom.xml` includes `com.azure:azure-messaging-servicebus:7.17.17`
- `ServiceBusDemo.java` demonstrates:
  - Single queue message
  - Five-message `ServiceBusMessageBatch`
  - `receiveMessages()` with `complete()`
  - Continuous `ServiceBusProcessorClient` handlers
  - Topic sending and subscription receiving
  - Explicit `close()` cleanup in `finally` blocks
- `README.md` documents required environment variables and execution with `mvn compile exec:java`

The example compiles successfully.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- All sender, receiver, and processor clients have deterministic cleanup, including methods that manage multiple clients.

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
hyoka run --prompt-id service-bus-dp-java-crud --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
