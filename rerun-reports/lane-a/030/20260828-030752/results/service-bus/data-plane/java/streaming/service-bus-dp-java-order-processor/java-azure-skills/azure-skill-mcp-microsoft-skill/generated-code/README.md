# Service Bus Order Processor

A Java 17 Maven sample with synchronous and asynchronous Azure Service Bus senders and session receivers.

## Prerequisites

- Java 17 and Maven 3.9+
- An Azure Service Bus queue created with **sessions enabled**
- A system-assigned or user-assigned managed identity with the **Azure Service Bus Data Sender** and
  **Azure Service Bus Data Receiver** roles

Set these environment variables:

```powershell
$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "orders"
# Only for a user-assigned managed identity:
$env:AZURE_CLIENT_ID = "00000000-0000-0000-0000-000000000000"
```

Build and run:

```powershell
mvn clean package
mvn exec:java
```

The demo intentionally sends malformed JSON to exercise dead-letter inspection and repair. Orders over
`$1,000` are marked high priority and scheduled 30 seconds into the future. The customer name is used as
the Service Bus session ID, and each processor consumes one session serially to preserve per-customer order.
