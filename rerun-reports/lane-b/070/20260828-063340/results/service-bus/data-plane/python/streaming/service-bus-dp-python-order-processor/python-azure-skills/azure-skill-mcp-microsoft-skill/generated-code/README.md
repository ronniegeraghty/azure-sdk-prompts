# Azure Service Bus Order Processor

This sample sends and processes orders with both synchronous and asynchronous
Azure Service Bus clients. It uses message sessions keyed by customer, explicit
dead-letter settlement, size-aware message batches, and delayed high-priority
orders.

## Prerequisites

- Python 3.10 or newer
- An existing Azure Service Bus queue with **sessions enabled**
- `Azure Service Bus Data Sender` and `Azure Service Bus Data Receiver` access
  for the identity used by `DefaultAzureCredential`

The project does not create or modify Azure resources.

## Run

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = "your-namespace.servicebus.windows.net"
$env:SERVICEBUS_QUEUE_NAME = "orders"
python main.py
```

The demo first runs the synchronous sender and processor, intentionally
dead-letters malformed JSON, inspects and repairs it, and processes the repaired
order. It then repeats the cycle with the asynchronous implementation.

Orders whose `total_price` exceeds the configured threshold are marked as high
priority and scheduled 30 seconds in the future. Later orders for that customer
are held behind the same scheduling barrier so they cannot overtake the delayed
order.
