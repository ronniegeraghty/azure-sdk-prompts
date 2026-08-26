# Azure Service Bus Python demo

This example uses an existing Azure Service Bus namespace, queue, topic, and
subscription. It does not create or modify Azure resources.

## Install

```powershell
python -m pip install -r requirements.txt
```

## Configure

Set the connection string and entity names in the current PowerShell session:

```powershell
$env:AZURE_SERVICEBUS_CONNECTION_STRING = "<service-bus-connection-string>"
$env:AZURE_SERVICEBUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICEBUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = "<subscription-name>"
```

The connection string must grant send and listen permissions. The queue, topic,
and subscription must already exist.

## Run

Run both synchronous and asynchronous examples:

```powershell
python .\service_bus_demo.py
```

Run only one pattern:

```powershell
python .\service_bus_demo.py --mode sync
python .\service_bus_demo.py --mode async
```

The script sends a five-message `ServiceBusMessageBatch` to the queue, receives
and completes those messages, then sends to a topic and receives and completes
the message through a subscription. The asynchronous example uses
`azure.servicebus.aio` and runs independent queue and topic operations
concurrently with `asyncio.gather`.
