# Azure Blob Event Notifier

This Python project receives Azure Blob Storage lifecycle events in either the
Event Grid native schema or CloudEvents 1.0 schema, downloads newly created
blobs, logs deletions, and publishes downstream custom events.

## Run the offline demo

```powershell
python -m venv .venv
.\.venv\Scripts\python.exe -m pip install -r requirements.txt
.\.venv\Scripts\python.exe main.py
```

The demo injects local mock clients, so it does not connect to Azure.

## Use with Azure

Set the resource endpoints; authentication is resolved passwordlessly through
`DefaultAzureCredential`.

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:AZURE_EVENT_GRID_TOPIC_ENDPOINT = "https://<topic>.<region>-1.eventgrid.azure.net/api/events"
```

Grant the workload identity only the data-plane roles it needs, such as
**Storage Blob Data Reader** on the storage account and **EventGrid Data
Sender** on the custom topic. Application code can create real clients through
`blob_event_notifier.config`; callers are responsible for closing those clients.

## References

- [Azure Event Grid Python SDK](https://learn.microsoft.com/python/api/overview/azure/eventgrid-readme)
- [EventGridEvent deserialization](https://learn.microsoft.com/python/api/azure-eventgrid/azure.eventgrid.eventgridevent)
- [CloudEvent deserialization](https://learn.microsoft.com/python/api/azure-core/azure.core.messaging.cloudevent)
- [Azure Blob Storage Python SDK](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-python)

