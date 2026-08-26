# Azure Event Hubs TypeScript sample

Install the required packages:

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
npm install --save-dev typescript tsx @types/node
```

`@azure/storage-blob` provides the `ContainerClient` required by
`BlobCheckpointStore`.

Set the required environment variables. Use your own values; do not commit
connection strings.

```powershell
$env:EVENT_HUBS_CONNECTION_STRING = "<event-hubs-connection-string>"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-connection-string>"
$env:CHECKPOINT_CONTAINER_NAME = "<blob-container-name>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default' # Optional
```

Run the sample:

```powershell
npm install
npm start
```

The identity represented by the storage connection string must be able to
create or access the checkpoint container.
