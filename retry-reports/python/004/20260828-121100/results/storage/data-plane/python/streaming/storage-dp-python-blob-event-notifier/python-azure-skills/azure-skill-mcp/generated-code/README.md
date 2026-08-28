# Azure Blob Event Notifier

Processes Blob Storage created/deleted events delivered by Event Grid in either
Event Grid schema or CloudEvents 1.0 schema, then publishes custom downstream
events. Both synchronous and asynchronous Azure SDK clients are included.

## Setup

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

For real Azure clients, grant the workload identity `Storage Blob Data Reader`
on the storage scope and an Event Grid data-sender role on the topic, then set:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:AZURE_EVENTGRID_TOPIC_ENDPOINT = "https://<topic>.<region>-1.eventgrid.azure.net/api/events"
```

`blob_event_notifier.config` creates passwordless clients with
`DefaultAzureCredential`; no keys or SAS tokens are accepted. The demo is
local-only and uses in-memory clients:

```powershell
python main.py
python -m unittest discover -s tests -v
```

Each `receive` call accepts one structured JSON event body and uses the Azure
SDK's `CloudEvent.from_json` or `EventGridEvent.from_json` helper. An HTTP
adapter should invoke it once per event when an Event Grid delivery contains a
batch.
