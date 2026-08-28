# Azure Service Bus messaging demo

This .NET 8 console application demonstrates:

- Sending one queue message.
- Building and sending a five-message `ServiceBusMessageBatch`.
- Receiving queue messages with `ReceiveMessagesAsync` and settling them with
  `CompleteMessageAsync`.
- Continuous queue processing with `ServiceBusProcessor` message and error
  handlers.
- Sending to a topic and receiving from a subscription.
- Asynchronous disposal of Service Bus clients, senders, receivers, and
  processors with `await using`.

## Required NuGet package

The project references:

```xml
<PackageReference Include="Azure.Messaging.ServiceBus" Version="7.20.2" />
```

To add it to another project:

```powershell
dotnet add package Azure.Messaging.ServiceBus --version 7.20.2
```

No additional package is needed for connection-string authentication.

## Configuration

Create the queue, topic, and subscription before running the sample. The
connection string must grant data-plane send and receive permissions for those
entities. Store it outside source control and set these environment variables:

```powershell
$env:AZURE_SERVICE_BUS_CONNECTION_STRING = "<namespace-connection-string>"
$env:AZURE_SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
dotnet run
```

`ServiceBusMessageBatch` uses synchronous `using` because it implements
`IDisposable`; the network clients implement `IAsyncDisposable` and therefore
use `await using`.

## References

- [Azure Service Bus client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.servicebus-readme)
- [Service Bus topics and subscriptions quickstart for .NET](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-dotnet-how-to-use-topics-subscriptions)
