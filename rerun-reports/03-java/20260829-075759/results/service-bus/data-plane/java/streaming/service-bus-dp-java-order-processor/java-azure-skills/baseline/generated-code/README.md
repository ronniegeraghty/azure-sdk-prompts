# Service Bus order processor

Java 17 example using Azure Service Bus queue sessions, managed identity, synchronous
clients, and asynchronous clients.

## Prerequisites

- Java 17 and Maven 3.9+
- An existing Azure Service Bus queue with **sessions enabled**
- A managed identity with the Azure Service Bus Data Sender and Data Receiver roles

No deployment commands are included; the project expects an existing queue.

## Run

Set the fully qualified namespace and queue name, then run the demo:

```powershell
$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "orders"
# Optional for a user-assigned managed identity:
$env:AZURE_CLIENT_ID = "managed-identity-client-id"
mvn compile exec:java
```

The demo sends valid orders and deliberately malformed JSON through each implementation.
Malformed messages are explicitly dead-lettered and then inspected without being removed
from the dead-letter queue. The dead-letter processors also expose reprocessing support,
which completes a dead-letter message only after it has been successfully re-enqueued.
Orders over `1000.00` receive a `priority=high` application property and are scheduled
30 seconds into the future. The demo's processing window includes this fraud-review delay.

Every message uses the order ID as its correlation ID and the customer name as its session
ID. Both processors limit concurrency to one session and one call per session, which
preserves order within a customer and prevents different customer sessions from being
processed concurrently.
