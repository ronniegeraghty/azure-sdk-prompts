# Azure Event Hubs Java send/receive sample

This sample sends 10 events in one batch, receives them with `EventProcessorClient`,
prints their bodies, and checkpoints each successfully processed event in Blob Storage.
The Event Hub and Blob container must already exist.

## Configuration

Set these environment variables. Keep secrets out of source control.

```powershell
$env:EVENT_HUBS_CONNECTION_STRING = "<event-hubs-namespace-connection-string>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-account-connection-string>"
$env:BLOB_CONTAINER_NAME = "<existing-checkpoint-container>"
```

The Event Hubs connection string needs permission to send and receive. The Blob Storage
connection string needs read/write access to the checkpoint container. For production,
use a dedicated consumer group rather than `$Default`.

## Run

```powershell
mvn compile exec:java
```

`EventContext.updateCheckpoint()` is called only after the event body is printed. Existing
checkpoints take precedence over the configured earliest starting position.
