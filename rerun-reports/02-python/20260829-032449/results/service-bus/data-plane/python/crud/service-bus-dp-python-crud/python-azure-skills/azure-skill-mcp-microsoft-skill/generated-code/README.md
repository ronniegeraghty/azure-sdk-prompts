# Azure Service Bus Python messaging examples

These scripts demonstrate queue batching, explicit message completion, asynchronous
processing, and topic/subscription messaging. They do not create or change Azure
resources; the queue, topic, and subscription must already exist.

## Install

Python 3.9 or newer is required.

```powershell
python -m pip install -r requirements.txt
```

Required packages:

- `azure-servicebus==7.14.3`
- `azure-identity==1.25.3`

## Configure

Authenticate locally with a supported `DefaultAzureCredential` developer login.
In Azure, use managed identity. Grant the identity the Azure Service Bus Data Sender
and Azure Service Bus Data Receiver roles at the narrowest appropriate scope.

```powershell
$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = "<namespace>.servicebus.windows.net"
$env:SERVICEBUS_QUEUE_NAME = "<queue-name>"
$env:SERVICEBUS_TOPIC_NAME = "<topic-name>"
$env:SERVICEBUS_SUBSCRIPTION_NAME = "<subscription-name>"
```

For production, also set `AZURE_TOKEN_CREDENTIALS=prod` to constrain
`DefaultAzureCredential` to production-safe credential types.

## Run

```powershell
python .\sync_servicebus.py
python .\async_servicebus.py
```

Each script sends a five-message `ServiceBusMessageBatch` to the queue, receives up
to five messages, processes and completes each message, then publishes a topic
message and receives it from the configured subscription. Receives have bounded
wait times, so an empty entity does not block indefinitely. The async version uses
prefetching and concurrent processing/settlement for higher throughput.

Reference: [Azure Service Bus client library for Python](https://learn.microsoft.com/en-us/python/api/overview/azure/servicebus-readme?view=azure-python)
