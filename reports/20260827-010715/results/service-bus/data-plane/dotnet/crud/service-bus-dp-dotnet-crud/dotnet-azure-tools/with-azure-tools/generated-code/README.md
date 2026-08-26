# Azure Service Bus messaging demo

This .NET console application demonstrates:

- Sending one message to a queue
- Safely sending a batch of five messages
- Receiving and completing queue messages
- Continuous processing with `ServiceBusProcessor`
- Sending to a topic and receiving from a subscription

## Required NuGet package

```powershell
dotnet add package Azure.Messaging.ServiceBus --version 7.20.1
```

The package is already declared in `ServiceBusMessagingDemo.csproj`.
`Azure.Identity` is not required because this sample intentionally uses a
connection string. For production applications, prefer Microsoft Entra ID and
managed identity instead of connection-string credentials.

## Configuration

Create the queue, topic, and subscription before running the sample. Then set
these environment variables in PowerShell:

```powershell
$env:AZURE_SERVICEBUS_CONNECTION_STRING = "<connection-string>"
$env:AZURE_SERVICEBUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICEBUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = "<subscription-name>"
dotnet run
```

Do not commit the connection string to source control.

## References

- [Azure Service Bus client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.servicebus-readme)
- [Azure.Messaging.ServiceBus NuGet package](https://www.nuget.org/packages/Azure.Messaging.ServiceBus/7.20.1)
