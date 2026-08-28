# Azure Service Bus order processor

Java 17 sample with synchronous and Reactor-based asynchronous send/receive
implementations. The queue **must have sessions enabled** because each message
uses the customer name as its session ID. Processors consume one session at a
time, preserving per-customer FIFO order without interleaving customers.

## Configuration

The managed identity needs `Azure Service Bus Data Sender` and
`Azure Service Bus Data Receiver` on the queue or namespace.

```powershell
$env:SERVICE_BUS_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE = "orders"
# Optional user-assigned managed identity:
$env:AZURE_CLIENT_ID = "00000000-0000-0000-0000-000000000000"
# Optional; defaults to 1000.00:
$env:HIGH_PRIORITY_THRESHOLD = "1000.00"

mvn test
mvn exec:java
```

The demo deliberately sends one malformed message in each cycle to demonstrate
explicit dead-lettering and inspection. High-value orders are marked with the
`priority=high` application property and scheduled 30 seconds in the future.
The demo waits for that delay before its second processing pass.
