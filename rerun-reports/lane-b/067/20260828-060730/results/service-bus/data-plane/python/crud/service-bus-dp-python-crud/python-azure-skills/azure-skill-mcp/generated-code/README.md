# Azure Service Bus Python messaging examples

These examples demonstrate:

- Sending exactly five queue messages in a `ServiceBusMessageBatch`
- Receiving and completing queue messages after processing
- Sending to a topic and receiving from a subscription
- Synchronous and `azure.servicebus.aio` asynchronous context-manager patterns

## Install

```powershell
python -m pip install -r requirements.txt
```

The only required package is `azure-servicebus`. Its AMQP transport dependencies
are installed automatically.

## Configure

Set these environment variables before connecting:

```powershell
$env:SERVICEBUS_CONNECTION_STR = "<namespace connection string>"
$env:SERVICEBUS_QUEUE_NAME = "<queue name>"
$env:SERVICEBUS_TOPIC_NAME = "<topic name>"
$env:SERVICEBUS_SUBSCRIPTION_NAME = "<subscription name>"
```

Use a connection string whose shared access policy has permission to send and
receive. Do not commit connection strings.

## Run

Both scripts default to a local dry run, which does not connect to Azure:

```powershell
python sync_service_bus.py
python async_service_bus.py
```

After configuring an Azure Service Bus namespace and entities, opt in to network
operations:

```powershell
python sync_service_bus.py --execute
python async_service_bus.py --execute
```

Each executable run sends a fresh batch before receiving, so it does not depend
on messages left by a previous run. The topic examples assume the subscription
already exists; subscriptions receive only messages sent after they are created.
