# Azure Event Hubs TypeScript sample

This sample sends a batch of 10 events, receives events through a
checkpointing consumer, and closes all clients when `SIGINT` or `SIGTERM` is
received.

## Install

```powershell
npm install
```

The runtime packages are:

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

`@azure/storage-blob` is required to create the `ContainerClient` used by
`BlobCheckpointStore`.

## Configure and run

Create the Blob Storage container before running the sample. Copy
`.env.example` to `.env`, replace its placeholders, then load the variables
into the current PowerShell session:

```powershell
Get-Content .env | ForEach-Object {
  $name, $value = $_ -split "=", 2
  Set-Item -Path "Env:$name" -Value $value
}

npm start
```

Press `Ctrl+C` to close the subscription, consumer, and producer cleanly.

Connection strings are used because this sample explicitly demonstrates that
authentication method. Keep them out of source control and prefer passwordless
authentication with managed identity for production applications.

## References

- [Azure Event Hubs JavaScript client library](https://learn.microsoft.com/javascript/api/overview/azure/event-hubs-readme?view=azure-node-latest)
- [Azure Blob checkpoint store JavaScript client library](https://learn.microsoft.com/javascript/api/overview/azure/eventhubs-checkpointstore-blob-readme?view=azure-node-latest)
