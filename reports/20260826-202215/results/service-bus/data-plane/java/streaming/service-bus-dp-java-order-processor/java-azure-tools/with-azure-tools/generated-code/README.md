# Azure Service Bus Order Processor

Small Java 17 Maven sample with synchronous and Reactor-based asynchronous senders and
processors for an Azure Service Bus queue.

## Prerequisites

- Java 17 and Maven 3.9+
- An Azure Service Bus queue with **sessions enabled**
- A managed identity with the `Azure Service Bus Data Sender` and
  `Azure Service Bus Data Receiver` roles scoped to the namespace or queue

The customer name is used as the Service Bus session ID. A receiver holds an exclusive
session lock and processes all messages in that session serially, so one customer's orders
cannot be interleaved with another customer's orders.

## Configuration

Set these environment variables:

```text
SERVICE_BUS_NAMESPACE=<namespace-name-or-fully-qualified-namespace>
SERVICE_BUS_QUEUE_NAME=<session-enabled-queue-name>
ORDER_HIGH_PRIORITY_THRESHOLD=1000.00
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

If `AZURE_CLIENT_ID` is omitted, the application uses the system-assigned managed identity.
No connection string or account key is used.

Orders above `ORDER_HIGH_PRIORITY_THRESHOLD` receive a `priority=high` application property
and a scheduled enqueue time 30 seconds in the future for fraud review. Other orders receive
`priority=normal`. All messages use the order ID as both the message ID and correlation ID.

## Build and run

```text
mvn clean verify
mvn exec:java
```

`Main` runs the synchronous cycle first and then the asynchronous cycle. Each cycle sends a
deliberately failed order, dead-letters it with a reason and description, reads it from the
dead-letter subqueue, resets it to `pending`, re-enqueues it, and processes it successfully.

Because the demo includes a scheduled high-value order, each cycle can wait up to 30 seconds
for that order to become available.

## References

- [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable)
- [Service Bus message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions)
- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
