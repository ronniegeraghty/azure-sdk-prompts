# Azure Event Hubs TypeScript example

This example sends a batch of 10 events with custom properties, receives events
with a blob-backed checkpoint store, and checkpoints each successfully processed
batch.

## Install

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
npm install --save-dev typescript tsx @types/node
```

`@azure/storage-blob` is required to create the `ContainerClient` used by
`BlobCheckpointStore`.

## Configure and run

Set the values shown in `.env.example` in your shell. The Event Hubs connection
string must have send and listen permissions, and the Storage connection string
must be able to create and update blobs in the checkpoint container.

```powershell
$env:EVENTHUB_CONNECTION_STRING = "<event-hubs-connection-string>"
$env:EVENTHUB_NAME = "<event-hub-name>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-connection-string>"
$env:BLOB_CONTAINER_NAME = "event-hub-checkpoints"

npm install
npm start
```

The consumer uses `$Default` unless `EVENTHUB_CONSUMER_GROUP` is set. Press
Ctrl+C to close the subscription, consumer, and producer gracefully.
