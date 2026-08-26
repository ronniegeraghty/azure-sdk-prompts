# Azure Event Hubs producer and processor sample

This console application sends a batch of 10 events and then processes events
using Blob Storage for checkpointing.

## Required packages

```powershell
dotnet add package Azure.Messaging.EventHubs --version 5.12.2
dotnet add package Azure.Messaging.EventHubs.Processor --version 5.12.2
dotnet add package Azure.Storage.Blobs --version 12.25.0
```

`Azure.Storage.Blobs` supplies `BlobContainerClient`, which the processor uses
as its checkpoint store.

## Configuration

Create the Event Hub and Blob container separately. The sample intentionally
does not create Azure resources. Set these environment variables before
running:

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<event-hubs-namespace-connection-string>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:BLOB_STORAGE_CONNECTION_STRING = "<storage-account-connection-string>"
$env:BLOB_CONTAINER_NAME = "<existing-checkpoint-container-name>"

dotnet run
```

The Event Hubs connection string needs permission to send and receive events.
The Storage connection string needs permission to read and write blobs in the
checkpoint container. Press `Ctrl+C` to stop processing cleanly.
