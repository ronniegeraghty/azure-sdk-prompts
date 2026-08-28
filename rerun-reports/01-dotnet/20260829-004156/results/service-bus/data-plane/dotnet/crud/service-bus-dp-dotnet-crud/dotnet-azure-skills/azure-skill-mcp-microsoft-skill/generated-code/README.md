# Azure Service Bus messaging demo

This .NET 8 console application demonstrates:

- sending one queue message and a five-message `ServiceBusMessageBatch`;
- receiving queue messages with `ReceiveMessagesAsync` and settling them with
  `CompleteMessageAsync`;
- continuous processing with `ServiceBusProcessor` message and error handlers;
- publishing to a topic and receiving from a subscription; and
- asynchronous disposal of Service Bus clients, senders, receivers, and the
  processor with `await using`.

## NuGet package

The sample requires `Azure.Messaging.ServiceBus` 7.20.1:

```powershell
dotnet add package Azure.Messaging.ServiceBus --version 7.20.1
```

`Azure.Identity` is not required because this example specifically uses a
connection string. For production applications, prefer Microsoft Entra ID and
managed identity rather than connection strings.

## Configuration

Create the queue, topic, and subscription before running the sample. Then set
the configuration through environment variables; do not put credentials in
source code.

```powershell
$env:AZURE_SERVICEBUS_CONNECTION_STRING = "Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<policy>;SharedAccessKey=<key>"
$env:AZURE_SERVICEBUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICEBUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = "<subscription-name>"

dotnet run
```

The connection string policy needs data-plane permission to send and receive.
This application does not create or modify Service Bus entities.

## Reference

- [Azure Service Bus client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.servicebus-readme)
