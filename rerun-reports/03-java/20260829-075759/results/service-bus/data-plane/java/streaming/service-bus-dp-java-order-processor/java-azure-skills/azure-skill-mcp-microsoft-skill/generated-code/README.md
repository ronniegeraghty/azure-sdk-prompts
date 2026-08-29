# Azure Service Bus order processor

A Java 17 Maven sample with synchronous and reactive senders/processors, JSON orders, size-aware batching,
scheduled high-value orders, explicit dead-lettering, and DLQ inspection/reprocessing.

## Prerequisites

- Java 17 and Maven 3.9+
- An existing Azure Service Bus queue with **sessions enabled**
- A managed identity (in Azure) or developer credential (locally) with the Azure Service Bus Data Sender and
  Azure Service Bus Data Receiver roles

No resources are provisioned by this project.

## Run

Set the fully qualified namespace, such as `contoso.servicebus.windows.net`. The queue defaults to `orders`.

```powershell
$env:AZURE_SERVICE_BUS_NAMESPACE = "contoso.servicebus.windows.net"
$env:AZURE_SERVICE_BUS_QUEUE = "orders"
mvn compile exec:java
```

`DefaultAzureCredential` uses managed identity when hosted in Azure and supported developer credentials locally.
Orders are assigned a session ID equal to the customer name. Batches are split by customer and by the
service-reported maximum batch size. Orders over `$1,000` get a `priority=high` application property and are
scheduled 30 seconds into the future for fraud review.

The processors use peek-lock mode and explicit settlement. Processing failures are dead-lettered with a reason and
description. DLQ readers log every inspected message and re-enqueue messages that still contain a valid `Order`.

## References

- [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme)
- [Service Bus message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions)
- [Passwordless authentication for Service Bus](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-java-how-to-use-queues)
