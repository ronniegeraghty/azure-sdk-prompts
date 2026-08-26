# Evaluation Report: storage-dp-java-blob-event-notifier

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 501.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-blob-event-notifier` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 16/22 |
| Duration | 501.9s |
| Timestamp | 2026-08-26T16:00:17Z |
| Files Generated | 25 |
| Event Count | 16925 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 184.4s |
| Review | 315.7s |
| **Total** | **501.9s** |

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
| Token Usage | in=235974 out=10146 |
| Turn Count | 8 |

## Prompt Sent

```
Create a small Java 17 Maven project that processes Azure Blob Storage lifecycle events delivered via Event Grid.

The project needs:

- An **event receiver class** (both sync and async versions) that accepts a JSON payload (as if received from an Event Grid webhook endpoint) and deserializes it into structured event objects. It should support both Event Grid schema and CloudEvents 1.0 schema (since Event Grid supports both and the configured schema may vary). It should route events by type — blob-created events get processed one way, blob-deleted another, and unrecognized types are logged as warnings.

- A **blob event handler class** that processes individual blob events. For blob-created events, it should parse the blob's container and name from the event subject, download the blob, and print a summary (name, size, content type, and the blob's access tier). For blob-deleted events, it should just log the deletion. It should handle race conditions gracefully — the blob may have already been deleted or moved to a different tier by the time we try to read it.

- An **event publisher class** (both sync and async versions) that can publish custom events to an Event Grid topic. Given a topic endpoint and a list of custom event objects, it should send them to Event Grid. This would be used for downstream notifications (e.g., "document processed" events). It should support setting a subject hierarchy for event filtering (e.g., "/documents/invoices/processed").

- A **configuration class** that connects to Azure Blob Storage and Event Grid securely. Authentication should use managed identity — no access keys or SAS tokens.

- A **Main class** that demos both implementations: constructs a sample Event Grid JSON payload (with both CloudEvents and EventGrid-schema examples) containing mock blob-created and blob-deleted events with realistic structure, feeds them through the receiver and handler, and publishes a custom downstream event. Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. powershell

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace contents"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eblob-event-notifier\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.15.4\u003c/azure.identity.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.30.0\u003c/azure.storage.blob.version\u003e\n+        \u003cazure.eventgrid.version\u003e4.30.0\u003c/azure.eventgrid.version\u003e\n+        \u003cjackson.version\u003e2.18.3\u003c/jackson.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+        \u003cjunit.version\u003e5.12.1\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-messaging-eventgrid\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.eventgrid.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.13.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.2\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.blobevents.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/blobevents/EventSchema.java\n+package com.example.blobevents;\n+\n+public enum EventSchema {\n+    EVENT_GRID,\n+    CLOUD_EVENTS\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobLifecycleEvent.java\n+package com.example.blobevents;\n+\n+import com.fasterxml.jackson.databind.JsonNode;\n+\n+import java.time.OffsetDateTime;\n+\n+public record BlobLifecycleEvent(\n+        EventSchema schema,\n+        String id,\n+        String type,\n+        String subject,\n+        OffsetDateTime time,\n+        JsonNode data) {\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPayloadParser.java\n+package com.example.blobevents;\n+\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.JsonNode;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+import java.time.OffsetDateTime;\n+import java.util.ArrayList;\n+import java.util.List;\n+\n+final class EventPayloadParser {\n+    private final ObjectMapper objectMapper;\n+\n+    EventPayloadParser(ObjectMapper objectMapper) {\n+        this.objectMapper = objectMapper;\n+    }\n+\n+    List\u003cBlobLifecycleEvent\u003e parse(String payload) {\n+        try {\n+            JsonNode root = objectMapper.readTree(payload);\n+            if (!root.isArray()) {\n+                throw new IllegalArgumentException(\"Event Grid webhook payload must be a JSON array\");\n+            }\n+\n+            List\u003cBlobLifecycleEvent\u003e events = new ArrayList\u003c\u003e();\n+            for (JsonNode node : root) {\n+                boolean cloudEvent = node.hasNonNull(\"specversion\");\n+                EventSchema schema = cloudEvent ? EventSchema.CLOUD_EVENTS : EventSchema.EVENT_GRID;\n+                String typeField = cloudEvent ? \"type\" : \"eventType\";\n+                String timeField = cloudEvent ? \"time\" : \"eventTime\";\n+                events.add(new BlobLifecycleEvent(\n+                        schema,\n+                        requiredText(node, \"id\"),\n+                        requiredText(node, typeField),\n+                        requiredText(node, \"subject\"),\n+                        OffsetDateTime.parse(requiredText(node, timeField)),\n+                        required(node, \"data\")));\n+            }\n+            return List.copyOf(events);\n+        } catch (JsonProcessingException e) {\n+            throw new IllegalArgumentException(\"Invalid event JSON payload\", e);\n+        }\n+    }\n+\n+    private static String requiredText(JsonNode node, String field) {\n+        JsonNode value = required(node, field);\n+        if (!value.isTextual() || value.textValue().isBlank()) {\n+            throw new IllegalArgumentException(\"Event field '\" + field + \"' must be a non-empty string\");\n+        }\n+        return value.textValue();\n+    }\n+\n+    private static JsonNode required(JsonNode node, String field) {\n+        JsonNode value = node.get(field);\n+        if (value == null || value.isNull()) {\n+            throw new IllegalArgumentException(\"Event is missing required field '\" + field + \"'\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobEventProcessor.java\n+package com.example.blobevents;\n+\n+public interface BlobEventProcessor {\n+    void onBlobCreated(BlobLifecycleEvent event);\n+\n+    void onBlobDeleted(BlobLifecycleEvent event);\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncBlobEventProcessor.java\n+package com.example.blobevents;\n+\n+import reactor.core.publisher.Mono;\n+\n+public interface AsyncBlobEventProcessor {\n+    Mono\u003cVoid\u003e onBlobCreated(BlobLifecycleEvent event);\n+\n+    Mono\u003cVoid\u003e onBlobDeleted(BlobLifecycleEvent event);\n+}\n*** Add File: src/main/java/com/example/blobevents/EventReceiver.java\n+package com.example.blobevents;\n+\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+import java.util.List;\n+\n+public final class EventReceiver {\n+    public static final String BLOB_CREATED = \"Microsoft.Storage.BlobCreated\";\n+    public static final String BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\";\n+\n+    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);\n+\n+    private final EventPayloadParser parser;\n+    private final BlobEventProcessor processor;\n+\n+    public EventReceiver(BlobEventProcessor processor) {\n+        this(processor, new ObjectMapper());\n+    }\n+\n+    EventReceiver(BlobEventProcessor processor, ObjectMapper objectMapper) {\n+        this.processor = processor;\n+        this.parser = new EventPayloadParser(objectMapper);\n+    }\n+\n+    public List\u003cBlobLifecycleEvent\u003e receive(String jsonPayload) {\n+        List\u003cBlobLifecycleEvent\u003e events = parser.parse(jsonPayload);\n+        events.forEach(this::route);\n+        return events;\n+    }\n+\n+    private void route(BlobLifecycleEvent event) {\n+        switch (event.type()) {\n+            case BLOB_CREATED -\u003e processor.onBlobCreated(event);\n+            case BLOB_DELETED -\u003e processor.onBlobDeleted(event);\n+            default -\u003e LOGGER.warn(\"Ignoring unsupported event type '{}' (event id {})\", event.type(), event.id());\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventReceiver.java\n+package com.example.blobevents;\n+\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncEventReceiver {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncEventReceiver.class);\n+\n+    private final EventPayloadParser parser;\n+    private final AsyncBlobEventProcessor processor;\n+\n+    public AsyncEventReceiver(AsyncBlobEventProcessor processor) {\n+        this(processor, new ObjectMapper());\n+    }\n+\n+    AsyncEventReceiver(AsyncBlobEventProcessor processor, ObjectMapper objectMapper) {\n+        this.processor = processor;\n+        this.parser = new EventPayloadParser(objectMapper);\n+    }\n+\n+    public Mono\u003cVoid\u003e receive(String jsonPayload) {\n+        return Mono.fromCallable(() -\u003e parser.parse(jsonPayload))\n+                .flatMapMany(Flux::fromIterable)\n+                .concatMap(this::route)\n+                .then();\n+    }\n+\n+    private Mono\u003cVoid\u003e route(BlobLifecycleEvent event) {\n+        return switch (event.type()) {\n+            case EventReceiver.BLOB_CREATED -\u003e processor.onBlobCreated(event);\n+            case EventReceiver.BLOB_DELETED -\u003e processor.onBlobDeleted(event);\n+            default -\u003e Mono.fromRunnable(() -\u003e\n+                    LOGGER.warn(\"Ignoring unsupported event type '{}' (event id {})\", event.type(), event.id()));\n+        };\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobLocation.java\n+package com.example.blobevents;\n+\n+import java.net.URLDecoder;\n+import java.nio.charset.StandardCharsets;\n+\n+public record BlobLocation(String container, String blobName) {\n+    private static final String CONTAINERS_MARKER = \"/containers/\";\n+    private static final String BLOBS_MARKER = \"/blobs/\";\n+\n+    public static BlobLocation fromSubject(String subject) {\n+        int containerStart = subject.indexOf(CONTAINERS_MARKER);\n+        int blobMarker = subject.indexOf(BLOBS_MARKER, containerStart + CONTAINERS_MARKER.length());\n+        if (containerStart \u003c 0 || blobMarker \u003c 0) {\n+            throw new IllegalArgumentException(\"Invalid blob event subject: \" + subject);\n+        }\n+\n+        String container = subject.substring(containerStart + CONTAINERS_MARKER.length(), blobMarker);\n+        String blobName = subject.substring(blobMarker + BLOBS_MARKER.length());\n+        if (container.isBlank() || blobName.isBlank()) {\n+            throw new IllegalArgumentException(\"Blob event subject has an empty container or blob name: \" + subject);\n+        }\n+        return new BlobLocation(decode(container), decode(blobName));\n+    }\n+\n+    private static String decode(String value) {\n+        return URLDecoder.decode(value.replace(\"+\", \"%2B\"), StandardCharsets.UTF_8);\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobSummary.java\n+package com.example.blobevents;\n+\n+public record BlobSummary(String name, long size, String contentType, String accessTier) {\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobAccess.java\n+package com.example.blobevents;\n+\n+@FunctionalInterface\n+public interface BlobAccess {\n+    BlobSummary download(String container, String blobName);\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncBlobAccess.java\n+package com.example.blobevents;\n+\n+import reactor.core.publisher.Mono;\n+\n+@FunctionalInterface\n+public interface AsyncBlobAccess {\n+    Mono\u003cBlobSummary\u003e download(String container, String blobName);\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureBlobAccess.java\n+package com.example.blobevents;\n+\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.models.BlobProperties;\n+\n+import java.io.OutputStream;\n+\n+public final class AzureBlobAccess implements BlobAccess {\n+    private final BlobServiceClient blobServiceClient;\n+\n+    public AzureBlobAccess(BlobServiceClient blobServiceClient) {\n+        this.blobServiceClient = blobServiceClient;\n+    }\n+\n+    @Override\n+    public BlobSummary download(String container, String blobName) {\n+        BlobClient client = blobServiceClient.getBlobContainerClient(container).getBlobClient(blobName);\n+        client.downloadStream(OutputStream.nullOutputStream());\n+        BlobProperties properties = client.getProperties();\n+        return new BlobSummary(\n+                blobName,\n+                properties.getBlobSize(),\n+                properties.getContentType(),\n+                properties.getAccessTier() == null ? \"unknown\" : properties.getAccessTier().toString());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureAsyncBlobAccess.java\n+package com.example.blobevents;\n+\n+import com.azure.storage.blob.BlobAsyncClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.models.BlobProperties;\n+import reactor.core.publisher.Mono;\n+\n+public final class AzureAsyncBlobAccess implements AsyncBlobAccess {\n+    private final BlobServiceAsyncClient blobServiceClient;\n+\n+    public AzureAsyncBlobAccess(BlobServiceAsyncClient blobServiceClient) {\n+        this.blobServiceClient = blobServiceClient;\n+    }\n+\n+    @Override\n+    public Mono\u003cBlobSummary\u003e download(String container, String blobName) {\n+        BlobAsyncClient client = blobServiceClient.getBlobContainerAsyncClient(container).getBlobAsyncClient(blobName);\n+        return client.downloadContent()\n+                .then(client.getProperties())\n+                .map(properties -\u003e toSummary(blobName, properties));\n+    }\n+\n+    private static BlobSummary toSummary(String blobName, BlobProperties properties) {\n+        return new BlobSummary(\n+                blobName,\n+                properties.getBlobSize(),\n+                properties.getContentType(),\n+                properties.getAccessTier() == null ? \"unknown\" : properties.getAccessTier().toString());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/BlobEventHandler.java\n+package com.example.blobevents;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+public final class BlobEventHandler implements BlobEventProcessor {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(BlobEventHandler.class);\n+\n+    private final BlobAccess blobAccess;\n+\n+    public BlobEventHandler(BlobAccess blobAccess) {\n+        this.blobAccess = blobAccess;\n+    }\n+\n+    @Override\n+    public void onBlobCreated(BlobLifecycleEvent event) {\n+        BlobLocation location = BlobLocation.fromSubject(event.subject());\n+        try {\n+            BlobSummary summary = blobAccess.download(location.container(), location.blobName());\n+            LOGGER.info(\"Blob downloaded: name={}, size={} bytes, contentType={}, accessTier={}\",\n+                    summary.name(), summary.size(), summary.contentType(), summary.accessTier());\n+        } catch (BlobStorageException e) {\n+            if (isExpectedRace(e)) {\n+                LOGGER.warn(\"Blob '{}' is no longer readable (status {}, error {}); it may have been deleted, \"\n+                                + \"renamed, or moved to an offline tier\",\n+                        location.blobName(), e.getStatusCode(), e.getErrorCode());\n+                return;\n+            }\n+            throw e;\n+        }\n+    }\n+\n+    @Override\n+    public void onBlobDeleted(BlobLifecycleEvent event) {\n+        BlobLocation location = BlobLocation.fromSubject(event.subject());\n+        LOGGER.info(\"Blob deleted: container={}, name={}\", location.container(), location.blobName());\n+    }\n+\n+    static boolean isExpectedRace(BlobStorageException error) {\n+        return error.getStatusCode() == 404\n+                || error.getStatusCode() == 409\n+                || error.getStatusCode() == 410;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncBlobEventHandler.java\n+package com.example.blobevents;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncBlobEventHandler implements AsyncBlobEventProcessor {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncBlobEventHandler.class);\n+\n+    private final AsyncBlobAccess blobAccess;\n+\n+    public AsyncBlobEventHandler(AsyncBlobAccess blobAccess) {\n+        this.blobAccess = blobAccess;\n+    }\n+\n+    @Override\n+    public Mono\u003cVoid\u003e onBlobCreated(BlobLifecycleEvent event) {\n+        return Mono.fromCallable(() -\u003e BlobLocation.fromSubject(event.subject()))\n+                .flatMap(location -\u003e blobAccess.download(location.container(), location.blobName())\n+                        .doOnNext(summary -\u003e LOGGER.info(\n+                                \"Blob downloaded: name={}, size={} bytes, contentType={}, accessTier={}\",\n+                                summary.name(), summary.size(), summary.contentType(), summary.accessTier()))\n+                        .then()\n+                        .onErrorResume(BlobStorageException.class, error -\u003e {\n+                            if (!BlobEventHandler.isExpectedRace(error)) {\n+                                return Mono.error(error);\n+                            }\n+                            LOGGER.warn(\"Blob '{}' is no longer readable (status {}, error {}); it may have been \"\n+                                            + \"deleted, renamed, or moved to an offline tier\",\n+                                    location.blobName(), error.getStatusCode(), error.getErrorCode());\n+                            return Mono.empty();\n+                        }));\n+    }\n+\n+    @Override\n+    public Mono\u003cVoid\u003e onBlobDeleted(BlobLifecycleEvent event) {\n+        return Mono.fromRunnable(() -\u003e {\n+            BlobLocation location = BlobLocation.fromSubject(event.subject());\n+            LOGGER.info(\"Blob deleted: container={}, name={}\", location.container(), location.blobName());\n+        });\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/CustomEvent.java\n+package com.example.blobevents;\n+\n+import java.util.Objects;\n+\n+public record CustomEvent(String subject, String eventType, String dataVersion, Object data) {\n+    public CustomEvent {\n+        requireHierarchy(subject);\n+        Objects.requireNonNull(eventType, \"eventType\");\n+        Objects.requireNonNull(dataVersion, \"dataVersion\");\n+        Objects.requireNonNull(data, \"data\");\n+    }\n+\n+    private static void requireHierarchy(String subject) {\n+        if (subject == null || !subject.startsWith(\"/\") || subject.length() == 1) {\n+            throw new IllegalArgumentException(\"subject must be a hierarchy beginning with '/', for example \"\n+                    + \"'/documents/invoices/processed'\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/EventGridSender.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+\n+import java.util.List;\n+\n+@FunctionalInterface\n+public interface EventGridSender {\n+    void send(List\u003cEventGridEvent\u003e events);\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventGridSender.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+\n+@FunctionalInterface\n+public interface AsyncEventGridSender {\n+    Mono\u003cVoid\u003e send(List\u003cEventGridEvent\u003e events);\n+}\n*** Add File: src/main/java/com/example/blobevents/EventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.BinaryData;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+\n+import java.util.List;\n+\n+public final class EventPublisher {\n+    private final EventGridSender sender;\n+\n+    public EventPublisher(String topicEndpoint, TokenCredential credential) {\n+        EventGridPublisherClient\u003cEventGridEvent\u003e client = new EventGridPublisherClientBuilder()\n+                .endpoint(topicEndpoint)\n+                .credential(credential)\n+                .buildEventGridEventPublisherClient();\n+        this.sender = client::sendEvents;\n+    }\n+\n+    public EventPublisher(EventGridSender sender) {\n+        this.sender = sender;\n+    }\n+\n+    public void publish(List\u003cCustomEvent\u003e events) {\n+        sender.send(events.stream().map(EventPublisher::toEventGridEvent).toList());\n+    }\n+\n+    static EventGridEvent toEventGridEvent(CustomEvent event) {\n+        return new EventGridEvent(\n+                event.subject(),\n+                event.eventType(),\n+                BinaryData.fromObject(event.data()),\n+                event.dataVersion());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AsyncEventPublisher.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;\n+import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+\n+public final class AsyncEventPublisher {\n+    private final AsyncEventGridSender sender;\n+\n+    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {\n+        EventGridPublisherAsyncClient\u003cEventGridEvent\u003e client = new EventGridPublisherClientBuilder()\n+                .endpoint(topicEndpoint)\n+                .credential(credential)\n+                .buildEventGridEventPublisherAsyncClient();\n+        this.sender = client::sendEvents;\n+    }\n+\n+    public AsyncEventPublisher(AsyncEventGridSender sender) {\n+        this.sender = sender;\n+    }\n+\n+    public Mono\u003cVoid\u003e publish(List\u003cCustomEvent\u003e events) {\n+        return sender.send(events.stream().map(EventPublisher::toEventGridEvent).toList());\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/AzureConfiguration.java\n+package com.example.blobevents;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class AzureConfiguration {\n+    private final String storageEndpoint;\n+    private final String eventGridTopicEndpoint;\n+    private final TokenCredential credential;\n+\n+    public AzureConfiguration(String storageEndpoint, String eventGridTopicEndpoint) {\n+        this(storageEndpoint, eventGridTopicEndpoint, new ManagedIdentityCredentialBuilder().build());\n+    }\n+\n+    AzureConfiguration(String storageEndpoint, String eventGridTopicEndpoint, TokenCredential credential) {\n+        this.storageEndpoint = requireHttpsEndpoint(storageEndpoint, \"storageEndpoint\");\n+        this.eventGridTopicEndpoint = requireHttpsEndpoint(eventGridTopicEndpoint, \"eventGridTopicEndpoint\");\n+        this.credential = Objects.requireNonNull(credential, \"credential\");\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        Map\u003cString, String\u003e environment = System.getenv();\n+        return new AzureConfiguration(\n+                requiredEnvironment(environment, \"AZURE_STORAGE_ENDPOINT\"),\n+                requiredEnvironment(environment, \"EVENT_GRID_TOPIC_ENDPOINT\"));\n+    }\n+\n+    public BlobServiceClient blobServiceClient() {\n+        return new BlobServiceClientBuilder()\n+                .endpoint(storageEndpoint)\n+                .credential(credential)\n+                .buildClient();\n+    }\n+\n+    public BlobServiceAsyncClient blobServiceAsyncClient() {\n+        return new BlobServiceClientBuilder()\n+                .endpoint(storageEndpoint)\n+                .credential(credential)\n+                .buildAsyncClient();\n+    }\n+\n+    public EventPublisher eventPublisher() {\n+        return new EventPublisher(eventGridTopicEndpoint, credential);\n+    }\n+\n+    public AsyncEventPublisher asyncEventPublisher() {\n+        return new AsyncEventPublisher(eventGridTopicEndpoint, credential);\n+    }\n+\n+    private static String requiredEnvironment(Map\u003cString, String\u003e environment, String name) {\n+        String value = environment.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static String requireHttpsEndpoint(String endpoint, String name) {\n+        if (endpoint == null || !endpoint.startsWith(\"https://\")) {\n+            throw new IllegalArgumentException(name + \" must be an HTTPS endpoint\");\n+        }\n+        return endpoint;\n+    }\n+}\n*** Add File: src/main/java/com/example/blobevents/Main.java\n+package com.example.blobevents;\n+\n+import com.azure.messaging.eventgrid.EventGridEvent;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class Main {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);\n+\n+    private static final String EVENT_GRID_PAYLOAD = \"\"\"\n+            [\n+              {\n+                \"id\": \"eg-created-001\",\n+                \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/invoices/2026/invoice-1042.pdf\",\n+                \"eventTime\": \"2026-08-26T15:10:00Z\",\n+                \"dataVersion\": \"3\",\n+                \"metadataVersion\": \"1\",\n+                \"data\": {\n+                  \"api\": \"PutBlob\",\n+                  \"clientRequestId\": \"f4f7f56c-f5d6-4b92-86be-021789745aed\",\n+                  \"requestId\": \"3d2e6d94-901e-0023-37a6-a72c68000000\",\n+                  \"eTag\": \"0x8DC000000000001\",\n+                  \"contentType\": \"application/pdf\",\n+                  \"contentLength\": 48291,\n+                  \"blobType\": \"BlockBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/invoices/2026/invoice-1042.pdf\",\n+                  \"sequencer\": \"000000000000000000000000000012340000000000000001\"\n+                }\n+              },\n+              {\n+                \"id\": \"eg-deleted-001\",\n+                \"eventType\": \"Microsoft.Storage.BlobDeleted\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf\",\n+                \"eventTime\": \"2026-08-26T15:11:00Z\",\n+                \"dataVersion\": \"3\",\n+                \"metadataVersion\": \"1\",\n+                \"data\": {\n+                  \"api\": \"DeleteBlob\",\n+                  \"requestId\": \"8c4b7552-401e-0025-08a6-a77a70000000\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf\",\n+                  \"sequencer\": \"000000000000000000000000000012350000000000000001\"\n+                }\n+              }\n+            ]\n+            \"\"\";\n+\n+    private static final String CLOUD_EVENTS_PAYLOAD = \"\"\"\n+            [\n+              {\n+                \"specversion\": \"1.0\",\n+                \"id\": \"ce-created-001\",\n+                \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage\",\n+                \"type\": \"Microsoft.Storage.BlobCreated\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/reports/quarterly-report.csv\",\n+                \"time\": \"2026-08-26T15:12:00Z\",\n+                \"datacontenttype\": \"application/json\",\n+                \"data\": {\n+                  \"api\": \"PutBlockList\",\n+                  \"contentType\": \"text/csv\",\n+                  \"contentLength\": 17320,\n+                  \"blobType\": \"BlockBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/reports/quarterly-report.csv\"\n+                }\n+              },\n+              {\n+                \"specversion\": \"1.0\",\n+                \"id\": \"ce-deleted-001\",\n+                \"source\": \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage\",\n+                \"type\": \"Microsoft.Storage.BlobDeleted\",\n+                \"subject\": \"/blobServices/default/containers/documents/blobs/reports/draft.csv\",\n+                \"time\": \"2026-08-26T15:13:00Z\",\n+                \"datacontenttype\": \"application/json\",\n+                \"data\": {\n+                  \"api\": \"DeleteBlob\",\n+                  \"url\": \"https://examplestorage.blob.core.windows.net/documents/reports/draft.csv\"\n+                }\n+              }\n+            ]\n+            \"\"\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        runSyncDemo();\n+        runAsyncDemo();\n+    }\n+\n+    private static void runSyncDemo() {\n+        LOGGER.info(\"--- Synchronous demo ---\");\n+        BlobEventHandler handler = new BlobEventHandler(Main::mockDownload);\n+        EventReceiver receiver = new EventReceiver(handler);\n+        receiver.receive(EVENT_GRID_PAYLOAD);\n+        receiver.receive(CLOUD_EVENTS_PAYLOAD);\n+\n+        EventPublisher publisher = new EventPublisher(Main::logPublishedEvents);\n+        publisher.publish(List.of(processedEvent(\"invoice-1042.pdf\")));\n+    }\n+\n+    private static void runAsyncDemo() {\n+        LOGGER.info(\"--- Asynchronous demo ---\");\n+        AsyncBlobEventHandler handler = new AsyncBlobEventHandler(\n+                (container, blobName) -\u003e Mono.fromSupplier(() -\u003e mockDownload(container, blobName)));\n+        AsyncEventReceiver receiver = new AsyncEventReceiver(handler);\n+        AsyncEventPublisher publisher = new AsyncEventPublisher(events -\u003e {\n+            logPublishedEvents(events);\n+            return Mono.empty();\n+        });\n+\n+        receiver.receive(EVENT_GRID_PAYLOAD)\n+                .then(receiver.receive(CLOUD_EVENTS_PAYLOAD))\n+                .then(publisher.publish(List.of(processedEvent(\"quarterly-report.csv\"))))\n+                .block();\n+    }\n+\n+    private static BlobSummary mockDownload(String container, String blobName) {\n+        LOGGER.info(\"Offline demo download: container={}, name={}\", container, blobName);\n+        String contentType = blobName.endsWith(\".pdf\") ? \"application/pdf\" : \"text/csv\";\n+        return new BlobSummary(blobName, blobName.endsWith(\".pdf\") ? 48_291 : 17_320, contentType, \"Hot\");\n+    }\n+\n+    private static CustomEvent processedEvent(String name) {\n+        return new CustomEvent(\n+                \"/documents/invoices/processed\",\n+                \"Contoso.Documents.Processed\",\n+                \"1.0\",\n+                Map.of(\"blobName\", name, \"status\", \"processed\"));\n+    }\n+\n+    private static void logPublishedEvents(List\u003cEventGridEvent\u003e events) {\n+        events.forEach(event -\u003e LOGGER.info(\n+                \"Offline demo publish: type={}, subject={}, data={}\",\n+                event.getEventType(), event.getSubject(), event.getData().toString()));\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/EventReceiverTest.java\n+package com.example.blobevents;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.ArrayList;\n+import java.util.List;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+class EventReceiverTest {\n+    @Test\n+    void parsesAndRoutesBothSchemas() {\n+        CapturingProcessor processor = new CapturingProcessor();\n+        EventReceiver receiver = new EventReceiver(processor);\n+\n+        List\u003cBlobLifecycleEvent\u003e gridEvents = receiver.receive(payload(\"eventType\", \"eventTime\", null));\n+        List\u003cBlobLifecycleEvent\u003e cloudEvents = receiver.receive(payload(\"type\", \"time\", \"\\\"specversion\\\":\\\"1.0\\\",\"));\n+\n+        assertEquals(EventSchema.EVENT_GRID, gridEvents.getFirst().schema());\n+        assertEquals(EventSchema.CLOUD_EVENTS, cloudEvents.getFirst().schema());\n+        assertEquals(List.of(\"created\", \"deleted\", \"created\", \"deleted\"), processor.calls);\n+    }\n+\n+    @Test\n+    void rejectsNonArrayPayload() {\n+        EventReceiver receiver = new EventReceiver(new CapturingProcessor());\n+        assertThrows(IllegalArgumentException.class, () -\u003e receiver.receive(\"{}\"));\n+    }\n+\n+    private static String payload(String typeField, String timeField, String extra) {\n+        String prefix = extra == null ? \"\" : extra;\n+        return \"\"\"\n+                [\n+                  {%s\"id\":\"1\",\"%s\":\"Microsoft.Storage.BlobCreated\",\"subject\":\"/blobServices/default/containers/c/blobs/a.txt\",\"%s\":\"2026-08-26T15:00:00Z\",\"data\":{}},\n+                  {%s\"id\":\"2\",\"%s\":\"Microsoft.Storage.BlobDeleted\",\"subject\":\"/blobServices/default/containers/c/blobs/b.txt\",\"%s\":\"2026-08-26T15:01:00Z\",\"data\":{}}\n+                ]\n+                \"\"\".formatted(prefix, typeField, timeField, prefix, typeField, timeField);\n+    }\n+\n+    private static final class CapturingProcessor implements BlobEventProcessor {\n+        private final List\u003cString\u003e calls = new ArrayList\u003c\u003e();\n+\n+        @Override\n+        public void onBlobCreated(BlobLifecycleEvent event) {\n+            calls.add(\"created\");\n+        }\n+\n+        @Override\n+        public void onBlobDeleted(BlobLifecycleEvent event) {\n+            calls.add(\"deleted\");\n+        }\n+    }\n+}\n*** Add File: src/test/java/com/example/blobevents/BlobLocationTest.java\n+package com.example.blobevents;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+class BlobLocationTest {\n+    @Test\n+    void parsesNestedAndEncodedBlobName() {\n+        BlobLocation location = BlobLocation.fromSubject(\n+                \"/blobServices/default/containers/documents/blobs/invoices/2026/Invoice%20A%2BB.pdf\");\n+\n+        assertEquals(\"documents\", location.container());\n+        assertEquals(\"invoices/2026/Invoice A+B.pdf\", location.blobName());\n+    }\n+\n+    @Test\n+    void rejectsInvalidSubject() {\n+        assertThrows(IllegalArgumentException.class, () -\u003e BlobLocation.fromSubject(\"/not/a/blob\"));\n+    }\n+}\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress test exec:java","description":"Build, test, and run demo","initial_wait":180,"mode":"sync"}
```

### 🔧 5. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/test/java/com/example/blobevents/EventReceiverTest.java\n@@\n-        assertEquals(EventSchema.EVENT_GRID, gridEvents.getFirst().schema());\n-        assertEquals(EventSchema.CLOUD_EVENTS, cloudEvents.getFirst().schema());\n+        assertEquals(EventSchema.EVENT_GRID, gridEvents.get(0).schema());\n+        assertEquals(EventSchema.CLOUD_EVENTS, cloudEvents.get(0).schema());\n*** End Patch\n"
```

### 🔧 6. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress test exec:java","description":"Retest and run complete demo","initial_wait":180,"mode":"sync"}
```

### 🔧 7. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/blobevents/Main.java\n@@\n-        AsyncEventPublisher publisher = new AsyncEventPublisher(events -\u003e {\n-            logPublishedEvents(events);\n-            return Mono.empty();\n-        });\n+        AsyncEventPublisher publisher = new AsyncEventPublisher(\n+                events -\u003e Mono.fromRunnable(() -\u003e logPublishedEvents(events)));\n*** End Patch\n"
```

### 🔧 8. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress --quiet test exec:java","description":"Verify final tests and demo order","initial_wait":180,"mode":"sync"}
```

## Generated Files

Generated 25 files (24 .java, 1 .xml)

- `pom.xml`
- `src\main\java\com\example\blobevents\AsyncBlobAccess.java`
- `src\main\java\com\example\blobevents\AsyncBlobEventHandler.java`
- `src\main\java\com\example\blobevents\AsyncBlobEventProcessor.java`
- `src\main\java\com\example\blobevents\AsyncEventGridSender.java`
- `src\main\java\com\example\blobevents\AsyncEventPublisher.java`
- `src\main\java\com\example\blobevents\AsyncEventReceiver.java`
- `src\main\java\com\example\blobevents\AzureAsyncBlobAccess.java`
- `src\main\java\com\example\blobevents\AzureBlobAccess.java`
- `src\main\java\com\example\blobevents\AzureConfiguration.java`
- `src\main\java\com\example\blobevents\BlobAccess.java`
- `src\main\java\com\example\blobevents\BlobEventHandler.java`
- `src\main\java\com\example\blobevents\BlobEventProcessor.java`
- `src\main\java\com\example\blobevents\BlobLifecycleEvent.java`
- `src\main\java\com\example\blobevents\BlobLocation.java`
- `src\main\java\com\example\blobevents\BlobSummary.java`
- `src\main\java\com\example\blobevents\CustomEvent.java`
- `src\main\java\com\example\blobevents\EventGridSender.java`
- `src\main\java\com\example\blobevents\EventPayloadParser.java`
- `src\main\java\com\example\blobevents\EventPublisher.java`
- ... and 5 more

## Copilot Response

I’ll inspect the empty workspace, then scaffold the Maven sources and focused tests around schema detection, routing, subject parsing, and managed-identity clients.

The workspace is empty. I’m implementing production Azure adapters behind small interfaces so the demo remains fully offline while exercising the same receivers, handlers, and publishers.

The production and offline demo paths are in place. I’m compiling and running the tests plus the full sync/async demo now, which will catch any Azure SDK signature mismatches.

Created a complete Java 17 Maven project with:

- Sync and async receivers supporting Event Grid and CloudEvents 1.0 schemas.
- Event-type routing for blob-created, blob-deleted, and unsupported events.
- Sync and async blob handlers with subject parsing, download summaries, access tiers, and race-condition handling.
- Sync and async custom Event Grid publishers with hierarchical subjects.
- Managed-identity configuration for Blob Storage and Event Grid.
- A fully offline `Main` demo running sync first, then async, without accessing real Azure resources.
- Unit tests for schema routing and blob-subject parsing.
- Complete Azure SDK dependency configuration in `pom.xml`.

Run it with:

```powershell
mvn test exec:java
```

For production clients, set `AZURE_STORAGE_ENDPOINT` and `EVENT_GRID_TOPIC_ENDPOINT`, then use `AzureConfiguration.fromEnvironment()`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- BlobServiceClient and BlobServiceAsyncClient do not implement AutoCloseable, so they do not require explicit closure.

### Issues

- Event Grid publisher clients are never closed, potentially leaking their underlying HTTP resources.

## Grader Results

- blob-event-notifier.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/10)
      - Handles Event Grid native schema via `EventGridEvent.fromString()` deserialization: Fail
      - Handles CloudEvents 1.0 schema via `CloudEvent.fromString()` deserialization: Fail
      - Does NOT manually parse JSON without the SDK's deserialization helpers: Fail
      - Routes events based on event type string (`Microsoft.Storage.BlobCreated`, `Microsoft.Storage.BlobDeleted`): Pass
      - Logs a warning for unrecognized event types: Pass
      - Parses container name and blob name from event subject (`/blobServices/default/containers/{container}/blobs/{blob}`): Pass
      - Publishes custom events with subject hierarchy for filtering: Pass
      - Retrieves and prints blob access tier from blob properties: Pass
      - Handles race condition: blob may no longer exist (catches `BlobStorageException` with 404 status): Pass
      - Catches Event Grid-specific exceptions for publishing errors: Fail
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
  - Service-Specific Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Pass
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 60% | 1.00 | 0.6000 | 5.7% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.4% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 13.00** | **Σ 10.6000** | **81.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-blob-event-notifier --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
