# Azure Event Hubs streaming sample

This console app sends one batch of 10 events and then receives events with
`EventProcessorClient`. Each successfully processed event is checkpointed in an
existing Azure Blob Storage container.

## Packages

```powershell
dotnet add package Azure.Messaging.EventHubs --version 5.12.2
dotnet add package Azure.Messaging.EventHubs.Processor --version 5.12.2
dotnet add package Azure.Storage.Blobs --version 12.25.1
```

`Azure.Storage.Blobs` is also required because `EventProcessorClient` uses a
`BlobContainerClient` as its checkpoint store.

## Configuration

Set these environment variables without placing credentials in source code:

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<Event Hubs namespace connection string>"
$env:EVENT_HUB_NAME = "<event hub name>"
$env:BLOB_STORAGE_CONNECTION_STRING = "<storage connection string>"
$env:BLOB_CONTAINER_NAME = "<existing checkpoint container>"
dotnet run
```

The Event Hubs connection string needs send and listen permissions. The storage
connection string needs blob read/write permissions for the checkpoint
container.

## References

- [Send or receive events from Azure Event Hubs using .NET](https://learn.microsoft.com/azure/event-hubs/event-hubs-dotnet-standard-getstarted-send)
- [EventProcessorClient class](https://learn.microsoft.com/dotnet/api/azure.messaging.eventhubs.eventprocessorclient)
