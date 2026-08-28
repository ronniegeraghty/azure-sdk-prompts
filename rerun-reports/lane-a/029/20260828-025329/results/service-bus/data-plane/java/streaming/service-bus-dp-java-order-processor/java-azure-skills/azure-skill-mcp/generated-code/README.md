# Azure Service Bus Order Processor

Small Java 17 sample showing synchronous and asynchronous Azure Service Bus order senders and processors.

## Prerequisites

- Java 17 and Maven 3.9+
- An existing Azure Service Bus namespace
- An existing **session-enabled** queue
- A managed identity, Azure CLI login, or another credential supported by `DefaultAzureCredential`
- The identity must have the Azure Service Bus Data Sender and Azure Service Bus Data Receiver roles

No Azure resources are created by this project.

## Configuration

Set:

```powershell
$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "orders"
```

The queue must have sessions enabled. The sender uses the customer name as the Service Bus session ID, and each processor accepts and drains one customer session at a time. This preserves order within each customer and avoids interleaving customers in a processing run.

Orders above `$1,000` receive the `priority=high` application property and are scheduled 30 seconds into the future for fraud review.

## Build and run

```powershell
mvn test
mvn exec:java
```

`DefaultAzureCredential` uses managed identity when deployed to Azure and can use developer credentials locally. The demo deliberately publishes malformed JSON so both the synchronous and asynchronous dead-letter inspection paths are exercised.
