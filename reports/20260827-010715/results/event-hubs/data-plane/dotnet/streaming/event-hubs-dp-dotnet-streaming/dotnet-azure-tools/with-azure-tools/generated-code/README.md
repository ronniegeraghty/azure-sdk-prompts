# Azure Event Hubs send/receive sample

Required packages are declared in `EventHubsSample.csproj`:

```powershell
dotnet add package Azure.Messaging.EventHubs --version 5.12.2
dotnet add package Azure.Messaging.EventHubs.Processor --version 5.12.2
dotnet add package Azure.Storage.Blobs --version 12.25.1
```

`Azure.Storage.Blobs` supplies `BlobContainerClient`, which
`EventProcessorClient` uses for ownership and checkpoint storage.

Set configuration through environment variables; do not commit connection
strings:

```powershell
$env:EVENTHUB_CONNECTION_STRING = '<Event Hubs namespace connection string>'
$env:EVENTHUB_NAME = '<event hub name>'
$env:BLOB_STORAGE_CONNECTION_STRING = '<Storage account connection string>'
$env:BLOB_CONTAINER_NAME = '<checkpoint container name>'

dotnet run
```

The Event Hubs connection string needs permission to send and receive. The
program sends ten events, starts the processor, and checkpoints each event
after successful processing. Press Ctrl+C to stop processing gracefully.

Reference:
[Azure Event Hubs .NET client library](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.eventhubs-readme)
