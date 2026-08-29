# Azure Event Hubs TypeScript sample

This sample sends a batch of 10 events and then receives events with blob-backed
partition ownership and checkpointing.

## Required packages

```bash
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
npm install --save-dev typescript @types/node
```

`@azure/storage-blob` is required to create the `ContainerClient` used by
`BlobCheckpointStore`.

## Run

Set the variables shown in `.env.example` in your shell, then run:

```bash
npm install
npm run build
npm start
```

Press `Ctrl+C` to close the subscription, consumer, and producer gracefully.
For production deployments, prefer Microsoft Entra ID with managed identity
over connection strings.

References:

- https://learn.microsoft.com/javascript/api/overview/azure/event-hubs-readme
- https://learn.microsoft.com/javascript/api/overview/azure/eventhubs-checkpointstore-blob-readme
