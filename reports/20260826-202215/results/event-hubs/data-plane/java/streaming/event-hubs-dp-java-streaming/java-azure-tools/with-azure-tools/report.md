# Evaluation Report: event-hubs-dp-java-streaming

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 389.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-java-streaming` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 14/19 |
| Duration | 389.4s |
| Timestamp | 2026-08-26T14:48:12Z |
| Files Generated | 3 |
| Event Count | 4737 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 132.4s |
| Review | 255.7s |
| **Total** | **389.4s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-eventhub-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=289440 out=3789 |
| Turn Count | 7 |

## Prompt Sent

```
Write a Java program that demonstrates
sending and receiving events with Azure Event Hubs:
1. Create an EventHubProducerClient using EventHubClientBuilder with a connection string
2. Create an EventDataBatch and add 10 events with custom properties
3. Send the batch to the event hub
4. Create an EventProcessorClient with BlobCheckpointStore for checkpointing
5. Register processEvent and processError consumer functions
6. Start the processor, wait for events, and print received event bodies
7. Implement checkpointing with EventContext.updateCheckpoint()

Show required Maven dependencies (azure-messaging-eventhubs and
azure-messaging-eventhubs-checkpointstore-blob) and proper resource cleanup.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-eventhub-java"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local Java Maven sample that sends an EventDataBatch to Azure Event Hubs and receives with EventProcessorClient and BlobCheckpointStore, using environment-variable placeholders and no resource provisioning.","learn":true}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local Java Maven Event Hubs producer and EventProcessorClient sample with BlobCheckpointStore.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-eventhub-java\\references\\examples.md"}
```

### 🔧 5. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 6. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-eventhub-java","pattern":"**/acceptance-criteria.md"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-messaging-eventhubs/maven-metadata.xml"}
```

### 🔧 8. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-messaging-eventhubs-checkpointstore-blob/maven-metadata.xml"}
```

### 🔧 9. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eevent-hubs-producer-processor\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0-SNAPSHOT\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventhubs\u003c/artifactId\u003e\n+            \u003cversion\u003e5.21.6\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventhubs-checkpointstore-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e1.21.8\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.EventHubsSample\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/EventHubsSample.java\n+package com.example;\n+\n+import com.azure.messaging.eventhubs.EventData;\n+import com.azure.messaging.eventhubs.EventDataBatch;\n+import com.azure.messaging.eventhubs.EventHubClientBuilder;\n+import com.azure.messaging.eventhubs.EventHubProducerClient;\n+import com.azure.messaging.eventhubs.EventProcessorClient;\n+import com.azure.messaging.eventhubs.EventProcessorClientBuilder;\n+import com.azure.messaging.eventhubs.checkpointstore.blob.BlobCheckpointStore;\n+import com.azure.messaging.eventhubs.models.ErrorContext;\n+import com.azure.messaging.eventhubs.models.EventContext;\n+import com.azure.messaging.eventhubs.models.EventPosition;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.BlobContainerClientBuilder;\n+\n+import java.time.Duration;\n+import java.time.Instant;\n+import java.util.UUID;\n+import java.util.concurrent.CountDownLatch;\n+import java.util.concurrent.TimeUnit;\n+import java.util.function.Consumer;\n+\n+public final class EventHubsSample {\n+    private static final int EVENT_COUNT = 10;\n+    private static final Duration RECEIVE_TIMEOUT = Duration.ofMinutes(2);\n+\n+    private EventHubsSample() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String eventHubsConnectionString = requiredEnvironmentVariable(\n+            \"EVENT_HUBS_CONNECTION_STRING\");\n+        String eventHubName = requiredEnvironmentVariable(\"EVENT_HUB_NAME\");\n+        String storageConnectionString = requiredEnvironmentVariable(\n+            \"AZURE_STORAGE_CONNECTION_STRING\");\n+        String checkpointContainer = requiredEnvironmentVariable(\n+            \"BLOB_CHECKPOINT_CONTAINER\");\n+        String sampleRunId = UUID.randomUUID().toString();\n+\n+        sendEvents(eventHubsConnectionString, eventHubName, sampleRunId);\n+\n+        CountDownLatch receivedSampleEvents = new CountDownLatch(EVENT_COUNT);\n+        BlobContainerAsyncClient blobContainer = new BlobContainerClientBuilder()\n+            .connectionString(storageConnectionString)\n+            .containerName(checkpointContainer)\n+            .buildAsyncClient();\n+\n+        Consumer\u003cEventContext\u003e processEvent = eventContext -\u003e {\n+            EventData event = eventContext.getEventData();\n+            String partitionId = eventContext.getPartitionContext().getPartitionId();\n+\n+            System.out.printf(\n+                \"Received partition=%s, sequence=%d, body=%s, properties=%s%n\",\n+                partitionId,\n+                event.getSequenceNumber(),\n+                event.getBodyAsString(),\n+                event.getProperties());\n+\n+            // Checkpoint only after the event has been processed successfully.\n+            eventContext.updateCheckpoint();\n+\n+            if (sampleRunId.equals(event.getProperties().get(\"sampleRunId\"))) {\n+                receivedSampleEvents.countDown();\n+            }\n+        };\n+\n+        Consumer\u003cErrorContext\u003e processError = errorContext -\u003e {\n+            String partitionId = errorContext.getPartitionContext() == null\n+                ? \"N/A\"\n+                : errorContext.getPartitionContext().getPartitionId();\n+            System.err.printf(\n+                \"Processor error on partition %s: %s%n\",\n+                partitionId,\n+                errorContext.getThrowable());\n+        };\n+\n+        EventProcessorClient processor = new EventProcessorClientBuilder()\n+            .connectionString(eventHubsConnectionString, eventHubName)\n+            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)\n+            .checkpointStore(new BlobCheckpointStore(blobContainer))\n+            .initialPartitionEventPosition(partitionId -\u003e EventPosition.earliest())\n+            .processEvent(processEvent)\n+            .processError(processError)\n+            .buildEventProcessorClient();\n+\n+        try {\n+            processor.start();\n+            System.out.printf(\n+                \"Processor started; waiting up to %d seconds for this run's events.%n\",\n+                RECEIVE_TIMEOUT.toSeconds());\n+\n+            boolean receivedAll = receivedSampleEvents.await(\n+                RECEIVE_TIMEOUT.toMillis(), TimeUnit.MILLISECONDS);\n+            if (!receivedAll) {\n+                throw new IllegalStateException(\n+                    \"Timed out waiting for all events; remaining=\"\n+                        + receivedSampleEvents.getCount());\n+            }\n+        } finally {\n+            processor.stop();\n+            System.out.println(\"Processor stopped.\");\n+        }\n+    }\n+\n+    private static void sendEvents(\n+        String connectionString,\n+        String eventHubName,\n+        String sampleRunId) {\n+\n+        try (EventHubProducerClient producer = new EventHubClientBuilder()\n+            .connectionString(connectionString, eventHubName)\n+            .buildProducerClient()) {\n+\n+            EventDataBatch batch = producer.createBatch();\n+            for (int i = 1; i \u003c= EVENT_COUNT; i++) {\n+                EventData event = new EventData(\"Sample event \" + i);\n+                event.getProperties().put(\"eventNumber\", i);\n+                event.getProperties().put(\"source\", \"java-event-hubs-sample\");\n+                event.getProperties().put(\"sampleRunId\", sampleRunId);\n+                event.getProperties().put(\"createdAt\", Instant.now().toString());\n+\n+                if (!batch.tryAdd(event)) {\n+                    throw new IllegalStateException(\n+                        \"Event \" + i + \" does not fit in the EventDataBatch.\");\n+                }\n+            }\n+\n+            producer.send(batch);\n+            System.out.printf(\"Sent %d events in one batch.%n\", batch.getCount());\n+        }\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(\n+                \"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Azure Event Hubs Java producer and processor\n+\n+This sample sends 10 events in one `EventDataBatch`, then receives events with\n+`EventProcessorClient`. Each successfully processed event is checkpointed in\n+Azure Blob Storage with `EventContext.updateCheckpoint()`.\n+\n+## Prerequisites\n+\n+- Java 17 or later\n+- Maven 3.9 or later\n+- An existing Event Hub\n+- An existing Blob Storage container for checkpoints\n+\n+The Event Hubs connection string must grant send and listen permissions. The\n+Storage connection string must grant read and write access to the checkpoint\n+container. Keep connection strings outside source control.\n+\n+## Configure and run in PowerShell\n+\n+```powershell\n+$env:EVENT_HUBS_CONNECTION_STRING = \"Endpoint=sb://\u003cnamespace\u003e.servicebus.windows.net/;SharedAccessKeyName=\u003cpolicy\u003e;SharedAccessKey=\u003ckey\u003e\"\n+$env:EVENT_HUB_NAME = \"\u003cevent-hub-name\u003e\"\n+$env:AZURE_STORAGE_CONNECTION_STRING = \"\u003cstorage-connection-string\u003e\"\n+$env:BLOB_CHECKPOINT_CONTAINER = \"\u003cexisting-container-name\u003e\"\n+\n+mvn compile exec:java\n+```\n+\n+The processor uses the `$Default` consumer group. For an isolated sample, use a\n+dedicated consumer group or clear old sample checkpoints before running.\n+\n+## References\n+\n+- [Azure Event Hubs Java client library](https://learn.microsoft.com/java/api/overview/azure/messaging-eventhubs-readme)\n+- [Process events with `EventProcessorClient`](https://learn.microsoft.com/azure/event-hubs/event-hubs-java-get-started-send)\n*** End Patch\n"
```

### 🔧 10. powershell

**Input:**

```json
{"command":"mvn --quiet -DskipTests package","description":"Compile the Maven sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\EventHubsSample.java`

## Copilot Response

I’ll build this as a runnable Maven sample, using environment variables for both connection strings and a bounded wait so cleanup always occurs.

Created a complete Maven sample:

- `src/main/java/com/example/EventHubsSample.java` — sends 10 events with custom properties, processes them with `EventProcessorClient`, prints bodies, checkpoints via `EventContext.updateCheckpoint()`, handles errors, and cleans up clients.
- `pom.xml` — includes `azure-messaging-eventhubs` 5.21.6 and `azure-messaging-eventhubs-checkpointstore-blob` 1.21.8.
- `README.md` — documents required environment variables and PowerShell run commands.

The implementation follows the [Azure Event Hubs Java client guidance](https://learn.microsoft.com/java/api/overview/azure/messaging-eventhubs-readme) and requires existing Event Hub and Blob checkpoint-container resources.

## Reviewer Notes (LLM-as-Judge)

### Summary

The generated code fully satisfies the client resource-cleanup criterion.

### Strengths

- Uses try-with-resources for the producer client and guarantees processor shutdown with finally.

## Grader Results

- send-receive-events.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `azure-messaging-eventhubs` and `azure-messaging-eventhubs-checkpointstore-blob` Maven deps: Pass
      - `EventHubClientBuilder` and `EventHubProducerClient`: Pass
      - `createBatch()` and `EventDataBatch.tryAdd()`: Pass
      - `send()` for publishing events: Pass
      - `EventProcessorClientBuilder` with `BlobCheckpointStore`: Pass
      - `processEvent` and `processError` consumer functions: Pass
      - `EventContext.updateCheckpoint()` for reliable processing: Pass
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
hyoka run --prompt-id event-hubs-dp-java-streaming --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
