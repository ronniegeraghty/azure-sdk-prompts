# Azure Event Hubs TypeScript example

This example sends a batch of 10 events, receives them with a
checkpoint-aware consumer, and closes all clients on completion or when
`Ctrl+C` is pressed.

## Install

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
npm install --save-dev typescript @types/node
```

`@azure/event-hubs` provides the producer and consumer clients.
`@azure/eventhubs-checkpointstore-blob` provides `BlobCheckpointStore`, and
`@azure/storage-blob` provides the container client used by that store.

## Configure and run

Set these values to your own resources. The Blob container must already
exist; the example does not create Azure resources.

```powershell
$env:EVENT_HUB_CONNECTION_STRING = "<event-hubs-namespace-connection-string>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-account-connection-string>"
$env:CHECKPOINT_CONTAINER_NAME = "<existing-blob-container-name>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default' # Optional

npm run build
npm start
```

The consumer starts at its stored checkpoint. If no checkpoint exists, it
starts at the earliest retained event so it can receive the batch that this
program sends before subscribing.
