# Azure Event Hubs streaming demo

This console app sends a batch of 10 events, then receives events with
`EventProcessorClient`. The processor stores partition ownership and checkpoints
in an existing Azure Blob Storage container.

## Required packages

```xml
<PackageReference Include="Azure.Messaging.EventHubs" Version="5.12.2" />
<PackageReference Include="Azure.Messaging.EventHubs.Processor" Version="5.12.2" />
<PackageReference Include="Azure.Storage.Blobs" Version="12.25.1" />
```

`Azure.Storage.Blobs` is also required because `BlobContainerClient` supplies the
processor's checkpoint store.

## Configuration

Set these environment variables. Use a namespace-level Event Hubs connection
string with permission to send and listen; do not commit connection strings.
The blob container must already exist.

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<event-hubs-connection-string>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:BLOB_STORAGE_CONNECTION_STRING = "<blob-storage-connection-string>"
$env:BLOB_CONTAINER_NAME = "<existing-container-name>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default' # Optional

dotnet run
```

Press Ctrl+C to stop the processor cleanly.

`EventHubProducerClient` and `EventDataBatch` are disposed with `await using` and
`using`, respectively. `EventProcessorClient` is not disposable; its lifecycle
is closed by awaiting `StopProcessingAsync` in a `finally` block.

For production, prefer passwordless authentication with managed identity and
checkpoint periodically rather than after every event when throughput is high.

## References

- [Publish events using `EventHubProducerClient`](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.eventhubs-readme)
- [Process events using `EventProcessorClient`](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.eventhubs.processor-readme)
