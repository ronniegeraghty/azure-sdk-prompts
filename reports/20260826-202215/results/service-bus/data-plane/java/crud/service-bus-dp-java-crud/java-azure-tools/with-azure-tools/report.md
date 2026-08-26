# Evaluation Report: service-bus-dp-java-crud

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 345.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-java-crud` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 16/19 |
| Duration | 345.5s |
| Timestamp | 2026-08-26T12:36:48Z |
| Files Generated | 3 |
| Event Count | 5108 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 116.7s |
| Review | 227.7s |
| **Total** | **345.5s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=291306 out=4393 |
| Turn Count | 7 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-servicebus-java"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local-only Java Maven sample demonstrating Azure Service Bus queue sender, batch sending, pull receiver settlement, processor handlers, and topic/subscription messaging with safe cleanup. Do not provision or modify Azure resources.","learn":true}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Java Azure Service Bus SDK guidance for ServiceBusClientBuilder sender receiver processor, ServiceBusMessageBatch, receiveMessages, complete, topic/subscription messaging, and close cleanup.","learn":true}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-java-crud-java-azure-tools-with-azure-tools-133677744","pattern":"**/*"}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java SDK code-generation best practices for a local-only Service Bus sample.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Azure Service Bus Java SDK guidance.","parameters":{"query":"Azure Service Bus Java ServiceBusClientBuilder sender receiver processor ServiceBusMessageBatch receiveMessages complete topic subscription close"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Azure Service Bus Java SDK samples.","parameters":{"language":"java","query":"ServiceBusClientBuilder ServiceBusSenderClient ServiceBusMessageBatch ServiceBusReceiverClient receiveMessages complete ServiceBusProcessorClient topic subscription"}}
```

### 🔧 8. web_search (via github-mcp-server)

**Input:**

```json
{"query":"What is the latest stable version of Maven artifact com.azure:azure-messaging-servicebus as of August 2026? Cite Maven Central or official Microsoft sources."}
```

### 🔧 9. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the complete current Azure Service Bus Java SDK README identified by documentation search.","parameters":{"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable"}}
```

### 🔧 10. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787747851288-copilot-tool-output-017f989b5dd442e98deae5a0ae6439aa.txt","view_range":[1,80]}
```

### 🔧 11. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787747850961-copilot-tool-output-64ba352ba9c64c1ea7fe2db3272b35c3.txt","view_range":[1,40]}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-service-bus-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+            \u003cversion\u003e7.17.17\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.16.3\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/servicebus/ServiceBusMessagingDemo.java\n+package com.example.servicebus;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.messaging.servicebus.ServiceBusClientBuilder;\n+import com.azure.messaging.servicebus.ServiceBusErrorContext;\n+import com.azure.messaging.servicebus.ServiceBusMessage;\n+import com.azure.messaging.servicebus.ServiceBusMessageBatch;\n+import com.azure.messaging.servicebus.ServiceBusProcessorClient;\n+import com.azure.messaging.servicebus.ServiceBusReceiverClient;\n+import com.azure.messaging.servicebus.ServiceBusReceivedMessage;\n+import com.azure.messaging.servicebus.ServiceBusSenderClient;\n+import com.azure.messaging.servicebus.models.ServiceBusReceiveMode;\n+\n+import java.time.Duration;\n+import java.util.concurrent.CountDownLatch;\n+import java.util.concurrent.TimeUnit;\n+\n+public final class ServiceBusMessagingDemo {\n+    private static final Duration RECEIVE_WAIT = Duration.ofSeconds(10);\n+    private static final Duration PROCESSOR_WAIT = Duration.ofSeconds(30);\n+\n+    private ServiceBusMessagingDemo() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String namespace = requiredEnvironmentVariable(\"SERVICE_BUS_NAMESPACE\");\n+        String queueName = requiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+        String topicName = requiredEnvironmentVariable(\"SERVICE_BUS_TOPIC_NAME\");\n+        String subscriptionName = requiredEnvironmentVariable(\"SERVICE_BUS_SUBSCRIPTION_NAME\");\n+\n+        TokenCredential credential = new DefaultAzureCredentialBuilder().build();\n+        ServiceBusClientBuilder clientBuilder = new ServiceBusClientBuilder()\n+            .credential(namespace, credential);\n+\n+        demonstrateQueueMessaging(clientBuilder, queueName);\n+        demonstrateProcessor(clientBuilder, queueName);\n+        demonstrateTopicAndSubscription(clientBuilder, topicName, subscriptionName);\n+    }\n+\n+    private static void demonstrateQueueMessaging(\n+        ServiceBusClientBuilder clientBuilder,\n+        String queueName\n+    ) {\n+        ServiceBusSenderClient sender = clientBuilder.sender()\n+            .queueName(queueName)\n+            .buildClient();\n+\n+        ServiceBusReceiverClient receiver = clientBuilder.receiver()\n+            .queueName(queueName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .buildClient();\n+\n+        try {\n+            sender.sendMessage(new ServiceBusMessage(\"Single queue message\"));\n+\n+            ServiceBusMessageBatch batch = sender.createMessageBatch();\n+            for (int i = 1; i \u003c= 5; i++) {\n+                ServiceBusMessage message = new ServiceBusMessage(\"Batch queue message \" + i);\n+                if (!batch.tryAddMessage(message)) {\n+                    throw new IllegalStateException(\"The batch cannot hold all 5 demo messages.\");\n+                }\n+            }\n+            sender.sendMessages(batch);\n+\n+            for (ServiceBusReceivedMessage message : receiver.receiveMessages(6, RECEIVE_WAIT)) {\n+                System.out.printf(\n+                    \"Queue message %s: %s%n\",\n+                    message.getMessageId(),\n+                    message.getBody().toString()\n+                );\n+                receiver.complete(message);\n+            }\n+        } finally {\n+            receiver.close();\n+            sender.close();\n+        }\n+    }\n+\n+    private static void demonstrateProcessor(\n+        ServiceBusClientBuilder clientBuilder,\n+        String queueName\n+    ) throws InterruptedException {\n+        CountDownLatch processed = new CountDownLatch(1);\n+\n+        ServiceBusProcessorClient processor = clientBuilder.processor()\n+            .queueName(queueName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .processMessage(context -\u003e {\n+                ServiceBusReceivedMessage message = context.getMessage();\n+                System.out.printf(\n+                    \"Processor received %s: %s%n\",\n+                    message.getMessageId(),\n+                    message.getBody().toString()\n+                );\n+                context.complete();\n+                processed.countDown();\n+            })\n+            .processError(ServiceBusMessagingDemo::processError)\n+            .buildProcessorClient();\n+\n+        ServiceBusSenderClient sender = clientBuilder.sender()\n+            .queueName(queueName)\n+            .buildClient();\n+\n+        try {\n+            processor.start();\n+            sender.sendMessage(new ServiceBusMessage(\"Message for continuous processor\"));\n+\n+            if (!processed.await(PROCESSOR_WAIT.toSeconds(), TimeUnit.SECONDS)) {\n+                throw new IllegalStateException(\"The processor did not receive the demo message in time.\");\n+            }\n+        } finally {\n+            processor.stop();\n+            processor.close();\n+            sender.close();\n+        }\n+    }\n+\n+    private static void demonstrateTopicAndSubscription(\n+        ServiceBusClientBuilder clientBuilder,\n+        String topicName,\n+        String subscriptionName\n+    ) {\n+        ServiceBusSenderClient topicSender = clientBuilder.sender()\n+            .topicName(topicName)\n+            .buildClient();\n+\n+        ServiceBusReceiverClient subscriptionReceiver = clientBuilder.receiver()\n+            .topicName(topicName)\n+            .subscriptionName(subscriptionName)\n+            .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)\n+            .disableAutoComplete()\n+            .buildClient();\n+\n+        try {\n+            topicSender.sendMessage(new ServiceBusMessage(\"Topic message\"));\n+\n+            for (ServiceBusReceivedMessage message\n+                : subscriptionReceiver.receiveMessages(1, RECEIVE_WAIT)) {\n+                System.out.printf(\n+                    \"Subscription message %s: %s%n\",\n+                    message.getMessageId(),\n+                    message.getBody().toString()\n+                );\n+                subscriptionReceiver.complete(message);\n+            }\n+        } finally {\n+            subscriptionReceiver.close();\n+            topicSender.close();\n+        }\n+    }\n+\n+    private static void processError(ServiceBusErrorContext context) {\n+        System.err.printf(\n+            \"Service Bus processor error. source=%s, namespace=%s, entity=%s, error=%s%n\",\n+            context.getErrorSource(),\n+            context.getFullyQualifiedNamespace(),\n+            context.getEntityPath(),\n+            context.getException()\n+        );\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Set the \" + name + \" environment variable.\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Azure Service Bus Java messaging demo\n+\n+This Maven sample demonstrates:\n+\n+- creating queue and topic senders with `ServiceBusClientBuilder`;\n+- sending one message and a five-message `ServiceBusMessageBatch`;\n+- receiving with `receiveMessages()` and explicitly completing messages;\n+- continuous queue processing with `ServiceBusProcessorClient` handlers; and\n+- sending to a topic and receiving from a subscription.\n+\n+The queue, topic, and subscription must already exist. The sample does not provision Azure resources.\n+\n+## Configuration\n+\n+The program uses `DefaultAzureCredential`, which supports local developer credentials and managed\n+identity without storing Service Bus keys in source code. Set these environment variables:\n+\n+```powershell\n+$env:SERVICE_BUS_NAMESPACE = \"\u003cnamespace\u003e.servicebus.windows.net\"\n+$env:SERVICE_BUS_QUEUE_NAME = \"\u003cqueue-name\u003e\"\n+$env:SERVICE_BUS_TOPIC_NAME = \"\u003ctopic-name\u003e\"\n+$env:SERVICE_BUS_SUBSCRIPTION_NAME = \"\u003csubscription-name\u003e\"\n+```\n+\n+The signed-in identity needs the **Azure Service Bus Data Sender** and **Azure Service Bus Data\n+Receiver** roles for the demonstrated entities.\n+\n+## Build and run\n+\n+```powershell\n+mvn compile\n+mvn exec:java \"-Dexec.mainClass=com.example.servicebus.ServiceBusMessagingDemo\"\n+```\n+\n+The required Service Bus Maven dependency is:\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-messaging-servicebus\u003c/artifactId\u003e\n+    \u003cversion\u003e7.17.17\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+`azure-identity` is also included because the sample uses passwordless `DefaultAzureCredential`.\n+\n+Reference: [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable)\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"mvn compile --quiet","description":"Compile the Maven Java sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\servicebus\ServiceBusMessagingDemo.java`

## Copilot Response

I’ll ground the sample in the current Azure Java SDK patterns, then inspect the workspace and implement it as a runnable local project without provisioning resources.

The workspace is empty, so I’ll create a minimal Maven project. The sample will use passwordless `DefaultAzureCredential`, explicit peek-lock settlement, bounded batch creation, processor error handlers, and deterministic client shutdown.

The implementation is in place. I’m compiling the exact SDK calls now; if Maven surfaces API or version issues, I’ll correct them before finishing.

Implemented a complete Maven Java sample:

- `ServiceBusSenderClient` for queue and topic messaging
- Single message and five-message `ServiceBusMessageBatch`
- `ServiceBusReceiverClient.receiveMessages()` with explicit `complete()`
- Continuous `ServiceBusProcessorClient` handlers with manual settlement
- Topic/subscription send and receive
- Passwordless `DefaultAzureCredential`
- Explicit `close()` cleanup for all clients

Files: `pom.xml`, `src\main\java\com\example\servicebus\ServiceBusMessagingDemo.java`, and `README.md`.

Reference: [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable)

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- All Service Bus clients have explicit cleanup in finally blocks, including the processor client.

## Grader Results

- send-receive-messages.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - `azure-messaging-servicebus` Maven dependency: Pass
      - `ServiceBusClientBuilder` with connection string: Fail
      - `ServiceBusSenderClient` and `ServiceBusMessage`: Pass
      - `createMessageBatch()` and `tryAddMessage()`: Pass
      - `ServiceBusReceiverClient.receiveMessages()` and `complete()`: Pass
      - `ServiceBusProcessorClient` with `processMessage` and `processError` handlers: Pass
      - Topic operations with `.topicName()` and `.subscriptionName()` on the builder: Pass
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 7.9% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.8571** | **83.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-java-crud --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
