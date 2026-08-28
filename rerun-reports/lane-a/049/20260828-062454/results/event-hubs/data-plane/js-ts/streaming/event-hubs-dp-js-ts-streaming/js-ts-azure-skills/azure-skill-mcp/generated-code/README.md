# Azure Event Hubs TypeScript sample

This sample sends a batch of 10 events, receives events with a blob-backed
checkpoint store, and closes all clients when `SIGINT` or `SIGTERM` is received.

## Install and build

```powershell
npm install
npm run build
```

The required Azure SDK packages are:

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

`@azure/storage-blob` supplies the `ContainerClient` required by
`BlobCheckpointStore`.

## Configure and run

Set the values from `.env.example` in the shell. The Event Hub and blob
container must already exist; the program does not provision Azure resources.

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<event-hubs-connection-string>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default'
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-connection-string>"
$env:BLOB_CONTAINER_NAME = "<existing-checkpoint-container>"
npm start
```

Use a dedicated consumer group and checkpoint container for repeatable runs.
If no checkpoint exists, the subscription starts at the earliest retained
event. Stop the program with `Ctrl+C` to close the subscription, consumer, and
producer cleanly.
