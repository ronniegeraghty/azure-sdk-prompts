# Service Bus Order Processor

Java 17 example with synchronous and asynchronous Azure Service Bus senders and
session-aware processors. The queue must have **sessions enabled when it is
created**. Each message uses the customer name as its session ID, which preserves
FIFO processing for that customer while allowing different customer sessions to
run concurrently.

## Configuration

The authenticated identity needs the **Azure Service Bus Data Sender** and
**Azure Service Bus Data Receiver** roles for the queue.

```powershell
$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "orders"
$env:HIGH_PRIORITY_THRESHOLD = "1000.00" # optional
mvn clean test
mvn exec:java
```

`DefaultAzureCredential` uses managed identity in Azure. For local development,
it can use a supported developer credential. The application never accepts a
connection string or access key.

Orders above the configured threshold receive a `priority=high` application
property and a 30-second scheduled enqueue time for fraud review. Batch senders
use `ServiceBusMessageBatch.tryAddMessage` and split full batches before sending.
Malformed demo messages are explicitly dead-lettered, inspected, corrected, and
requeued.

## References

- [Azure Service Bus Java client library](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme)
- [Service Bus message sessions](https://learn.microsoft.com/azure/service-bus-messaging/message-sessions)
- [Service Bus dead-letter queues](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-dead-letter-queues)
