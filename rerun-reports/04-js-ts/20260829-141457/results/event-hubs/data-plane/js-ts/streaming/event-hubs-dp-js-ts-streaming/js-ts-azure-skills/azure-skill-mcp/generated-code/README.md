# Azure Event Hubs TypeScript sample

This sample sends 10 events in one batch, then receives events with a
Blob-backed checkpoint store until `Ctrl+C` is pressed.

## Install and build

Node.js 20 or later is required.

```powershell
npm install
npm run build
```

The required runtime packages are:

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

`@azure/storage-blob` supplies the `ContainerClient` required by
`BlobCheckpointStore`.

## Configure and run

Set the values shown in `.env.example` in your shell, then run:

```powershell
npm start
```

The Event Hubs connection string must allow sending and receiving. The storage
connection string must allow access to an existing blob container used for
ownership records and checkpoints. Credentials are read only from environment
variables and are not stored in source code.

## References

- [Azure Event Hubs JavaScript client examples](https://learn.microsoft.com/javascript/api/overview/azure/event-hubs-readme?view=azure-node-latest#examples)
- [Blob checkpoint store JavaScript examples](https://learn.microsoft.com/javascript/api/overview/azure/eventhubs-checkpointstore-blob-readme?view=azure-node-latest#examples)
