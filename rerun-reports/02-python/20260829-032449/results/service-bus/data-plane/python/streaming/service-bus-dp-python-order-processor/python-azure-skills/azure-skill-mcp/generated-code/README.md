# Azure Service Bus Order Processor

This sample sends and processes orders with the synchronous and asynchronous
Azure Service Bus Python clients. It uses `DefaultAzureCredential`; no secrets
or connection strings are stored in the project.

## Prerequisites

- Python 3.11 or newer
- An Azure Service Bus **Premium or Standard** namespace
- A queue with **sessions enabled**
- The authenticated identity assigned the **Azure Service Bus Data Sender** and
  **Azure Service Bus Data Receiver** roles for the queue or namespace

Sessions are required because each message uses the customer name as its
`session_id`. The processors accept and fully drain one session at a time,
which prevents orders for the same customer from being processed out of order
or interleaved with another customer's orders. Batch sends are grouped by
customer because every message in a session-aware batch must use the same
session ID.

## Run

Set the namespace hostname and, optionally, the queue name:

```powershell
$env:SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE = "example.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "orders"
python -m pip install -r requirements.txt
python main.py
```

The demo sends normal and high-priority orders. High-priority orders are
scheduled 30 seconds into the future for fraud review. It also sends malformed
JSON, dead-letters it, inspects the dead-letter queue, repairs and resubmits the
demo order, and processes it again. The synchronous cycle runs first, followed
by the asynchronous cycle.

## References

- [Azure Service Bus Python client library](https://learn.microsoft.com/python/api/overview/azure/servicebus-readme)
- [Get started with Azure Service Bus queues in Python](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-queues)
