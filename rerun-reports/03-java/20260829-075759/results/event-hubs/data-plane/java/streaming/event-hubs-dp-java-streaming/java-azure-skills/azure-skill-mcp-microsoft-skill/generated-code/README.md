# Azure Event Hubs Java producer and processor

This sample sends ten events in one `EventDataBatch`, then receives and
checkpoints them with `EventProcessorClient` and `BlobCheckpointStore`.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An existing Event Hub
- An existing Blob container dedicated to checkpoints

Set these PowerShell environment variables without committing their values:

```powershell
$env:EVENT_HUBS_CONNECTION_STRING = "Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<name>;SharedAccessKey=<key>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-connection-string>"
$env:BLOB_CHECKPOINT_CONTAINER = "<existing-container-name>"
```

The Event Hubs connection string must be namespace-scoped because the Event
Hub name is supplied separately. The checkpoint container must already exist.

Compile and run:

```powershell
mvn compile
mvn exec:java
```

On a new checkpoint store, the processor starts from the earliest retained
event in each partition. On later runs, it resumes from stored checkpoints.
Use a dedicated Event Hub or consumer group if exactly ten received events are
expected, because earlier uncheckpointed events may also be delivered.

## Maven dependencies

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-eventhubs</artifactId>
    <version>5.21.0</version>
</dependency>
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-eventhubs-checkpointstore-blob</artifactId>
    <version>1.21.0</version>
</dependency>
```

Reference: [Azure Event Hubs samples for Java](https://learn.microsoft.com/azure/event-hubs/event-hubs-java-get-started-send)
