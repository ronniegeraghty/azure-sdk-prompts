# Azure Service Bus Order Processor

Java 17 Maven sample with synchronous and asynchronous order senders and
session-aware processors. The Service Bus queue must be created with sessions
enabled because each message uses the customer name as its session ID.

Set these environment variables:

```text
SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net
SERVICE_BUS_QUEUE_NAME=orders
```

`DefaultAzureCredential` uses the workload's managed identity when the sample
runs in Azure. Assign that identity the **Azure Service Bus Data Sender** and
**Azure Service Bus Data Receiver** roles. For local development it can use
Azure CLI or another credential in its default chain. Then run:

```text
mvn compile exec:java
```

Orders above `$1,000` are marked high priority and scheduled 30 seconds into
the future. The demo includes an order with `FAILED` status to exercise
dead-letter inspection and reprocessing.
