# Evaluation Report: event-hubs-dp-java-streaming

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 432.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-java-streaming` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 14/19 |
| Duration | 432.3s |
| Timestamp | 2026-08-26T14:40:59Z |
| Files Generated | 2 |
| Event Count | 4256 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 114.2s |
| Review | 316.7s |
| **Total** | **432.3s** |

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
| Token Usage | in=136477 out=2971 |
| Turn Count | 6 |

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

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-java-streaming-java-azure-tools-baseline-1234380690","pattern":"**/*"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Latest stable Maven versions as of August 2026 for com.azure azure-messaging-eventhubs and azure-messaging-eventhubs-checkpointstore-blob, and current Java EventProcessorClient BlobCheckpointStore APIs"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eevent-hubs-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e11\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventhubs\u003c/artifactId\u003e\n+            \u003cversion\u003e5.21.6\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventhubs-checkpointstore-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e1.21.7\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.EventHubsDemo\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/EventHubsDemo.java\n+package com.example;\n+\n+import com.azure.messaging.eventhubs.EventData;\n+import com.azure.messaging.eventhubs.EventDataBatch;\n+import com.azure.messaging.eventhubs.EventHubClientBuilder;\n+import com.azure.messaging.eventhubs.EventHubProducerClient;\n+import com.azure.messaging.eventhubs.EventPosition;\n+import com.azure.messaging.eventhubs.EventProcessorClient;\n+import com.azure.messaging.eventhubs.EventProcessorClientBuilder;\n+import com.azure.messaging.eventhubs.checkpointstore.blob.BlobCheckpointStore;\n+import com.azure.messaging.eventhubs.models.ErrorContext;\n+import com.azure.messaging.eventhubs.models.EventContext;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.BlobContainerClientBuilder;\n+\n+import java.time.Duration;\n+import java.util.UUID;\n+import java.util.concurrent.CountDownLatch;\n+import java.util.concurrent.TimeUnit;\n+\n+public final class EventHubsDemo {\n+    private static final int EVENT_COUNT = 10;\n+    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(60);\n+\n+    private EventHubsDemo() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String eventHubsConnectionString = requireEnvironmentVariable(\"EVENT_HUB_CONNECTION_STRING\");\n+        String storageConnectionString = requireEnvironmentVariable(\"AZURE_STORAGE_CONNECTION_STRING\");\n+        String blobContainerName = requireEnvironmentVariable(\"BLOB_CONTAINER_NAME\");\n+        String runId = UUID.randomUUID().toString();\n+\n+        EventHubClientBuilder eventHubClientBuilder = new EventHubClientBuilder()\n+            .connectionString(eventHubsConnectionString);\n+\n+        sendEvents(eventHubClientBuilder, runId);\n+\n+        BlobContainerAsyncClient blobContainerClient = new BlobContainerClientBuilder()\n+            .connectionString(storageConnectionString)\n+            .containerName(blobContainerName)\n+            .buildAsyncClient();\n+        blobContainerClient.createIfNotExists().block();\n+\n+        CountDownLatch receivedEvents = new CountDownLatch(EVENT_COUNT);\n+        EventProcessorClient processor = new EventProcessorClientBuilder()\n+            .connectionString(eventHubsConnectionString)\n+            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)\n+            .checkpointStore(new BlobCheckpointStore(blobContainerClient))\n+            .initialPartitionEventPosition(partitionId -\u003e EventPosition.earliest())\n+            .processEvent(context -\u003e processEvent(context, runId, receivedEvents))\n+            .processError(EventHubsDemo::processError)\n+            .buildEventProcessorClient();\n+\n+        try {\n+            processor.start();\n+            boolean receivedAll = receivedEvents.await(RECEIVE_TIMEOUT.toSeconds(), TimeUnit.SECONDS);\n+            if (!receivedAll) {\n+                System.err.printf(\n+                    \"Timed out after %d seconds; received %d of %d events from this run.%n\",\n+                    RECEIVE_TIMEOUT.toSeconds(),\n+                    EVENT_COUNT - receivedEvents.getCount(),\n+                    EVENT_COUNT);\n+            }\n+        } finally {\n+            processor.stop();\n+        }\n+    }\n+\n+    private static void sendEvents(EventHubClientBuilder clientBuilder, String runId) {\n+        try (EventHubProducerClient producer = clientBuilder.buildProducerClient()) {\n+            EventDataBatch batch = producer.createBatch();\n+\n+            for (int i = 1; i \u003c= EVENT_COUNT; i++) {\n+                EventData event = new EventData(\"Event \" + i);\n+                event.getProperties().put(\"sequence-number\", i);\n+                event.getProperties().put(\"sample-run-id\", runId);\n+                event.getProperties().put(\"source\", \"java-event-hubs-demo\");\n+\n+                if (!batch.tryAdd(event)) {\n+                    throw new IllegalStateException(\"The 10-event sample batch exceeded the Event Hubs size limit.\");\n+                }\n+            }\n+\n+            producer.send(batch);\n+            System.out.printf(\"Sent %d events.%n\", EVENT_COUNT);\n+        }\n+    }\n+\n+    private static void processEvent(\n+        EventContext context,\n+        String runId,\n+        CountDownLatch receivedEvents) {\n+\n+        EventData event = context.getEventData();\n+        System.out.printf(\n+            \"Partition %s, sequence %d: %s%n\",\n+            context.getPartitionContext().getPartitionId(),\n+            event.getSequenceNumber(),\n+            event.getBodyAsString());\n+\n+        context.updateCheckpoint();\n+\n+        if (runId.equals(event.getProperties().get(\"sample-run-id\"))) {\n+            receivedEvents.countDown();\n+        }\n+    }\n+\n+    private static void processError(ErrorContext context) {\n+        String partitionId = context.getPartitionContext() == null\n+            ? \"unknown\"\n+            : context.getPartitionContext().getPartitionId();\n+        System.err.printf(\n+            \"Error while processing partition %s: %s%n\",\n+            partitionId,\n+            context.getThrowable().getMessage());\n+        context.getThrowable().printStackTrace(System.err);\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"mvn --quiet compile","description":"Compile the Maven Event Hubs example","initial_wait":180,"mode":"sync"}
```

### 🔧 5. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/EventHubsDemo.java\n@@\n-import com.azure.messaging.eventhubs.EventPosition;\n import com.azure.messaging.eventhubs.EventProcessorClient;\n import com.azure.messaging.eventhubs.EventProcessorClientBuilder;\n import com.azure.messaging.eventhubs.checkpointstore.blob.BlobCheckpointStore;\n import com.azure.messaging.eventhubs.models.ErrorContext;\n import com.azure.messaging.eventhubs.models.EventContext;\n+import com.azure.messaging.eventhubs.models.EventPosition;\n*** End Patch\n"
```

### 🔧 6. powershell

**Input:**

```json
{"command":"mvn --quiet compile","description":"Recompile the corrected Maven example","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `pom.xml`
- `src\main\java\com\example\EventHubsDemo.java`

## Copilot Response

I’ll create a minimal Maven example using environment-variable placeholders, then compile it locally if Maven is available.

The directory is empty. I’ll add a complete Maven project; the processor will start from the earliest position when no checkpoint exists so it can receive the batch sent before startup.

Compilation found the current SDK places `EventPosition` in the `models` package. I’ll correct the import and recompile.

Created a complete Java/Maven example:

- `pom.xml` — includes `azure-messaging-eventhubs` 5.21.6 and `azure-messaging-eventhubs-checkpointstore-blob` 1.21.7.
- `src\main\java\com\example\EventHubsDemo.java` — sends 10 events with custom properties, processes and prints them, handles errors, updates checkpoints, and cleans up producer/processor resources.

Run it after setting:

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<connection string including EntityPath>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage connection string>"
$env:BLOB_CONTAINER_NAME = "<checkpoint container>"
mvn exec:java
```

The checkpoint container is created if absent, and the processor starts from the earliest event when no checkpoint exists.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The producer is automatically closed, and the event processor is stopped even if waiting or processing fails.

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
hyoka run --prompt-id event-hubs-dp-java-streaming --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
