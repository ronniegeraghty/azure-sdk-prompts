# Azure Event Hubs producer and processor sample

This .NET 8 console application sends a batch of 10 events and then receives
events with `EventProcessorClient`. Each successfully processed event is
checkpointed in Azure Blob Storage.

## Required packages

```powershell
dotnet add package Azure.Messaging.EventHubs --version 5.12.2
dotnet add package Azure.Messaging.EventHubs.Processor --version 5.12.2
dotnet add package Azure.Storage.Blobs --version 12.23.0
```

`Azure.Storage.Blobs` supplies the `BlobContainerClient` used by the processor
as its checkpoint store.

## Configuration

Set these environment variables to development/test resources. Do not commit
connection strings.

```powershell
$env:EVENTHUB_CONNECTION_STRING = 'Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<policy>;SharedAccessKey=<key>'
$env:EVENTHUB_NAME = '<event-hub-name>'
$env:BLOB_STORAGE_CONNECTION_STRING = '<storage-connection-string>'
$env:BLOB_CONTAINER_NAME = '<checkpoint-container-name>'
```

The Event Hubs shared access policy needs send and listen permissions. The
storage credentials need permission to create and update blobs in the
checkpoint container.

Run the sample:

```powershell
dotnet run
```

Use a dedicated consumer group for each independently scaled application in
production rather than the default consumer group used by this sample.
