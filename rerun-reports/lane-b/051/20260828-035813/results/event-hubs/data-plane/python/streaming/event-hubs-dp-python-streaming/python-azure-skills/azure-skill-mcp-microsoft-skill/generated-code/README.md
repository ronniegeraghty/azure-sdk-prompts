# Async Azure Event Hubs Python demo

Install the required packages:

```powershell
python -m pip install -r requirements.txt
```

Set configuration through environment variables:

```powershell
$env:EVENT_HUB_FULLY_QUALIFIED_NAMESPACE = "<namespace>.servicebus.windows.net"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:STORAGE_ACCOUNT_URL = "https://<storage-account>.blob.core.windows.net"
$env:CHECKPOINT_CONTAINER = "<existing-container-name>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default'
```

Authenticate locally with a developer credential supported by
`DefaultAzureCredential`, then run:

```powershell
python event_hubs_demo.py
```

The identity needs permission to send and receive Event Hubs data and to read
and write blobs in the checkpoint container. The script sends five events,
then receives events continuously. It checkpoints each event only after it is
printed. Stop it with Ctrl+C.
