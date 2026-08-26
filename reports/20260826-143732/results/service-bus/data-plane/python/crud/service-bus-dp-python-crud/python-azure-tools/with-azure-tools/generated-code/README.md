# Azure Service Bus Python messaging demo

This sample demonstrates:

- A `ServiceBusMessageBatch` containing five queue messages.
- Queue receive and explicit `complete_message()` settlement after processing.
- Topic publishing and subscription receiving.
- Synchronous `with` and asynchronous `async with` lifecycle management.
- Concurrent `azure.servicebus.aio` queue and pub/sub flows for higher throughput.

The queue, topic, and subscription must already exist. Use dedicated demo entities
because the receivers settle and remove messages from them. No resources are
created by these scripts.

## Install

Create and activate a virtual environment, then install the required packages:

```powershell
py -m venv .venv
.venv\Scripts\Activate.ps1
py -m pip install -r requirements.txt
```

`requirements.txt` installs the latest compatible releases of:

- `azure-servicebus`
- `azure-identity`

## Configure authentication and entities

The scripts use `DefaultAzureCredential`; they do not accept connection strings
or keys. For local development, sign in with a supported developer credential
such as Azure CLI or Visual Studio Code. In Azure, use a managed identity. The
identity needs the Azure Service Bus Data Sender and Azure Service Bus Data
Receiver roles scoped to the demo entities or namespace.

Set these environment variables in the current PowerShell session:

```powershell
$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICEBUS_QUEUE_NAME = "demo-queue"
$env:SERVICEBUS_TOPIC_NAME = "demo-topic"
$env:SERVICEBUS_SUBSCRIPTION_NAME = "demo-subscription"
```

For production, constrain `DefaultAzureCredential` to production-safe
credentials by setting `AZURE_TOKEN_CREDENTIALS=prod`.

## Run

```powershell
py service_bus_sync_demo.py
py service_bus_async_demo.py
```

The default receive wait is 10 seconds, so each script exits rather than waiting
indefinitely when an entity has no available message.

## References

- [Azure Service Bus client library for Python](https://learn.microsoft.com/python/api/overview/azure/servicebus-readme)
- [Passwordless connections with the Azure Identity library](https://learn.microsoft.com/azure/developer/python/sdk/authentication-overview)
