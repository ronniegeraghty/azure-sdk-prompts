# Azure Service Bus Python messaging demo

This sample demonstrates synchronous and asynchronous queue messaging, batch
sends, explicit message completion, and topic/subscription messaging. The queue,
topic, and subscription must already exist; the sample does not create or modify
Azure resources.

## Install

```powershell
python -m pip install -r requirements.txt
```

Required packages:

- `azure-servicebus`
- `azure-identity`

## Configure

Authenticate locally with an identity supported by `DefaultAzureCredential`,
then set the existing Service Bus namespace and entity names:

```powershell
$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICEBUS_QUEUE_NAME = "your-queue"
$env:SERVICEBUS_TOPIC_NAME = "your-topic"
$env:SERVICEBUS_SUBSCRIPTION_NAME = "your-subscription"
```

The identity needs the **Azure Service Bus Data Sender** and **Azure Service Bus
Data Receiver** roles scoped as narrowly as practical.

## Run

```powershell
python service_bus_demo.py
```

The synchronous examples use `with` for credentials, clients, senders, and
receivers. The higher-throughput example uses `azure.servicebus.aio`, `async
with`, and `asyncio.gather` to overlap independent queue and topic operations.

## References

- [Azure Service Bus client library for Python](https://learn.microsoft.com/python/api/overview/azure/servicebus-readme)
- [Send to topics and receive from subscriptions with Python](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-topics-subscriptions)
- [Passwordless authentication for Azure Service Bus](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-passwordless-messaging)
