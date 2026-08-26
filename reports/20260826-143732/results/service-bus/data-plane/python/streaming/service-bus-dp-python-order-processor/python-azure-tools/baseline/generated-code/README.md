# Azure Service Bus Order Processor

This sample sends and processes orders with the synchronous and asynchronous
Azure Service Bus Python clients.

## Prerequisites

Use an existing Azure Service Bus queue with **sessions enabled**. The sender
sets each message's session ID to the customer name, and each processor drains
one session before accepting another. This preserves order for each customer
without interleaving customers.

Authenticate with any identity supported by `DefaultAzureCredential` and set:

```text
SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net
SERVICE_BUS_QUEUE_NAME=<session-enabled-queue>
```

Install dependencies with `python -m pip install -r requirements.txt`, then run
`python main.py`.

Orders above the sender's high-priority threshold (USD 1,000 by default) carry
a `priority=high` application property and are scheduled 30 seconds into the
future for a fraud-review window. Service Bus queues do not provide native
priority ordering, so consumers can use this property for application-specific
handling.
