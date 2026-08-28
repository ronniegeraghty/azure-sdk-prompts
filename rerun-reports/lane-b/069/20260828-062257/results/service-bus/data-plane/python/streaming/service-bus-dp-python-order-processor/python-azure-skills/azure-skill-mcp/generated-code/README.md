# Azure Service Bus Order Processor

This sample publishes and processes orders with both synchronous and
asynchronous Azure Service Bus clients. It uses Service Bus sessions keyed by
customer to preserve per-customer ordering, message batching, scheduled
high-value orders, and explicit dead-letter handling.

## Prerequisites

- Python 3.10 or later
- An Azure Service Bus namespace and a **session-enabled queue**
- A signed-in identity supported by `DefaultAzureCredential` with Azure Service
  Bus Data Sender and Data Receiver permissions

Set these environment variables:

```text
SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net
SERVICE_BUS_QUEUE_NAME=<session-enabled-queue>
```

Install dependencies with `python -m pip install -r requirements.txt`, then run
`python main.py`.

Orders over $1,000 carry a `priority=high` application property and are
scheduled 30 seconds into the future. Azure Service Bus queues do not provide a
native priority field, so consumers can use this property when applying their
own prioritization policy. The demo also sends malformed JSON, dead-letters it,
repairs and republishes it through the dead-letter inspection API, and then
processes the repaired order.
