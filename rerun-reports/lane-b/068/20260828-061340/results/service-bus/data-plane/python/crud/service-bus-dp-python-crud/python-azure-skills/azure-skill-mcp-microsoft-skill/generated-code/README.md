# Azure Service Bus Python examples

These scripts use Microsoft Entra authentication through
`DefaultAzureCredential`; no connection string is stored in source.

## Install

```powershell
python -m pip install -r requirements.txt
```

Required packages:

- `azure-servicebus`
- `azure-identity`

## Configure

Create the queue, topic, and subscription before running the examples. Set the
values shown in `.env.example` in the current shell:

```powershell
$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICEBUS_QUEUE_NAME = "your-queue"
$env:SERVICEBUS_TOPIC_NAME = "your-topic"
$env:SERVICEBUS_SUBSCRIPTION_NAME = "your-subscription"
```

Authenticate locally with a credential supported by `DefaultAzureCredential`,
such as Azure CLI or Visual Studio Code. In production, use managed identity
and set `AZURE_TOKEN_CREDENTIALS=prod` to constrain the credential chain.

## Run

Run the synchronous context-manager example:

```powershell
python .\servicebus_sync.py
```

Run the asynchronous `aio` example:

```powershell
python .\servicebus_async.py
```

Both examples send a five-message `ServiceBusMessageBatch` to a queue, receive
and complete the messages after processing, publish to a topic, and receive
and complete a message from its subscription. The async example also uses
prefetching and concurrent processing for higher throughput.
