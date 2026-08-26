# Async Azure Event Hubs send and receive sample

This sample sends a batch of five events, receives events asynchronously, and
stores a checkpoint in Azure Blob Storage after each event is processed.
Authentication uses `DefaultAzureCredential`; no secrets or connection strings
are stored in the code.

## Install

Python 3.9 or later is required.

```powershell
python -m pip install -r requirements.txt
```

The required Event Hubs packages are `azure-eventhub` and
`azure-eventhub-checkpointstoreblob-aio`. The sample also installs
`azure-identity` for Microsoft Entra authentication.

## Configure and run

Set these variables to existing Azure resources:

```powershell
$env:EVENT_HUB_FULLY_QUALIFIED_NAMESPACE = "<namespace>.servicebus.windows.net"
$env:EVENT_HUB_NAME = "<event-hub-name>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default'
$env:STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:CHECKPOINT_CONTAINER = "<existing-container-name>"
python .\event_hubs_async.py
```

The identity selected by `DefaultAzureCredential` needs the Azure Event Hubs
Data Sender and Azure Event Hubs Data Receiver roles, plus Storage Blob Data
Contributor access to the checkpoint container. The receiver runs until
Ctrl+C. `starting_position="-1"` applies only when a partition has no stored
checkpoint; subsequent runs resume from the checkpoint.

## References

- [Azure Event Hubs Python client library](https://learn.microsoft.com/python/api/overview/azure/eventhub-readme)
- [Async Blob checkpoint store package](https://github.com/Azure/azure-sdk-for-python/tree/main/sdk/eventhub/azure-eventhub-checkpointstoreblob-aio)
