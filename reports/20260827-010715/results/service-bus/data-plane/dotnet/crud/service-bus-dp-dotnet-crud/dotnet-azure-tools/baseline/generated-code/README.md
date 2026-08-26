# Azure Service Bus messaging demo

## Required package

The project references the Azure Service Bus SDK:

```powershell
dotnet add package Azure.Messaging.ServiceBus --version 7.18.2
```

No additional NuGet package is required for this console sample.

## Configuration

Create the queue, topic, and subscription separately. Then set configuration through
environment variables; do not put credentials in source control:

```powershell
$env:AZURE_SERVICE_BUS_CONNECTION_STRING = "<service-bus-connection-string>"
$env:AZURE_SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
dotnet run
```

The connection string must grant send and listen permissions for the configured
entities. The sample sends and receives queue messages, runs a queue processor until
Enter is pressed, and then sends a topic message and receives it from the subscription.
