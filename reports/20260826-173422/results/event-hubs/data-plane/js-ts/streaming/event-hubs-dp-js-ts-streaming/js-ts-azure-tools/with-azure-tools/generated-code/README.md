# Azure Event Hubs TypeScript sample

This sample sends 10 events in one Event Hubs batch, receives events with a
blob-backed checkpoint store, updates checkpoints after successful processing,
and closes all clients on completion, `SIGINT`, or `SIGTERM`.

## Install

```bash
npm install
```

The Azure runtime packages are:

```bash
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

`@azure/storage-blob` is required to create the `ContainerClient` used by
`BlobCheckpointStore`.

## Configure

Copy `.env.example` to your preferred local environment file or export its
variables in your shell. The Event Hubs and Storage connection strings are read
only from environment variables and must not be committed.

The checkpoint container must already exist. The Event Hubs connection string
needs send and listen permissions; the Storage connection string needs blob
read/write permissions.

## Run

```bash
npm start
```

The receiver exits after it observes the 10 events sent by the current run.
Press Ctrl+C to stop earlier.

## References

- [Send events to or receive events from Event Hubs using JavaScript](https://learn.microsoft.com/azure/event-hubs/event-hubs-node-get-started-send)
- [Azure Event Hubs JavaScript SDK samples](https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/eventhub/event-hubs/samples)
