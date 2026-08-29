# Azure Event Hubs Java batch and processor sample

This sample sends a batch of 10 events, receives them with `EventProcessorClient`,
and stores ownership and checkpoints in Azure Blob Storage. The Event Hubs
connection string must include `EntityPath`.

The Blob container must already exist. Use a dedicated consumer group and Blob
container for each independently running processor application.

## Configuration

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<Event Hubs connection string with EntityPath>"
$env:BLOB_STORAGE_CONNECTION_STRING = "<Blob Storage connection string>"
$env:BLOB_CONTAINER_NAME = "<existing checkpoint container>"
$env:EVENT_HUB_CONSUMER_GROUP = "<consumer group>" # Optional; defaults to `$Default`
```

Credentials are read from environment variables and are not stored in source
control. For production workloads, prefer passwordless authentication with
managed identity instead of connection strings.

## Run

```powershell
mvn compile exec:java
```

The processor checkpoints after every event by calling
`EventContext.updateCheckpoint()`. For high-throughput production workloads,
checkpoint less frequently to reduce Blob Storage operations.

References:

- [Send and receive events with Java](https://learn.microsoft.com/azure/event-hubs/event-hubs-java-get-started-send)
- [Event processor with Blob checkpoint store sample](https://github.com/Azure/azure-sdk-for-java/blob/main/sdk/eventhubs/azure-messaging-eventhubs-checkpointstore-blob/src/samples/java/com/azure/messaging/eventhubs/checkpointstore/blob/EventProcessorBlobCheckpointStoreSample.java)
