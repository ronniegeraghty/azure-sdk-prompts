# Azure Event Hubs TypeScript example

This sample sends a batch of 10 events, then receives and checkpoints them with
Azure Blob Storage. The checkpoint Blob container must already exist.

## Install

```powershell
npm install
```

The runtime packages are:

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

`@azure/storage-blob` supplies the `ContainerClient` required by
`BlobCheckpointStore`.

## Run

Set the values shown in `.env.example` as environment variables, then run:

```powershell
npm start
```

Press Ctrl+C to close the subscription, consumer, and producer gracefully.
