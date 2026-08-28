# Azure Service Bus messaging demo

This .NET 8 console program demonstrates:

- Sending one queue message
- Sending five queue messages with `ServiceBusMessageBatch`
- Receiving queue messages with `ReceiveMessagesAsync`
- Settling messages with `CompleteMessageAsync`
- Continuous processing with `ServiceBusProcessor` handlers
- Publishing to a topic and receiving from a subscription
- Asynchronous disposal of clients, senders, receivers, and processors

## Required NuGet package

```powershell
dotnet add package Azure.Messaging.ServiceBus --version 7.20.1
```

`Azure.Identity` is not required because this sample explicitly uses a connection
string. For production applications, prefer Microsoft Entra authentication and add
`Azure.Identity`.

## Configuration

Create the queue, topic, and subscription before running the sample. Then set:

```powershell
$env:AZURE_SERVICEBUS_CONNECTION_STRING = "Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>"
$env:AZURE_SERVICEBUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICEBUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = "<subscription-name>"
```

The connection string policy needs send and listen permissions for the configured
entities.

## Run

```powershell
dotnet run
```
