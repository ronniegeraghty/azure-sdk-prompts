# Azure Event Hubs Java producer and processor

This sample sends 10 events in one `EventDataBatch`, then receives events with
`EventProcessorClient`. Each successfully processed event is checkpointed in
Azure Blob Storage with `EventContext.updateCheckpoint()`.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An existing Event Hub
- An existing Blob Storage container for checkpoints

The Event Hubs connection string must grant send and listen permissions. The
Storage connection string must grant read and write access to the checkpoint
container. Keep connection strings outside source control.

## Configure and run in PowerShell

```powershell
$env:EVENT_HUBS_CONNECTION_STRING = "Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<policy>;SharedAccessKey=<key>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-connection-string>"
$env:BLOB_CHECKPOINT_CONTAINER = "<existing-container-name>"

mvn compile exec:java
```

The processor uses the `$Default` consumer group. For an isolated sample, use a
dedicated consumer group or clear old sample checkpoints before running.

## References

- [Azure Event Hubs Java client library](https://learn.microsoft.com/java/api/overview/azure/messaging-eventhubs-readme)
- [Process events with `EventProcessorClient`](https://learn.microsoft.com/azure/event-hubs/event-hubs-java-get-started-send)
