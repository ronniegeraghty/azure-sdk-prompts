# Azure Service Bus Order Processor

This sample sends and processes orders with both synchronous and asynchronous
Azure Service Bus clients. Customer names are used as Service Bus session IDs,
so the queue must be created with **sessions enabled**. A processor drains one
customer session before opening the next, preserving per-customer FIFO order.

High-value orders have an application property of `priority=high` and a
scheduled enqueue time 30 seconds in the future. Azure Service Bus queues do
not provide a native priority queue, so consumers can use this property for
telemetry or downstream routing.

## Run

Use Python 3.10 or newer, install the dependencies, and configure an identity
with the Azure Service Bus Data Sender and Azure Service Bus Data Receiver
roles on the namespace or queue.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = "example.servicebus.windows.net"
$env:SERVICEBUS_QUEUE_NAME = "orders"
python main.py
```

The demo first runs the synchronous flow and then the asynchronous flow. Each
flow sends normal and scheduled orders, intentionally sends malformed JSON,
dead-letters that message, inspects it, replaces it with a valid order, and
processes the replacement.
