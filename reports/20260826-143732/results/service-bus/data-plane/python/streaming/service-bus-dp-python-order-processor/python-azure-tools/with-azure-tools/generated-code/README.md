# Azure Service Bus Order Processor

This sample sends and processes orders with both synchronous and asynchronous
Azure SDK clients. It uses Microsoft Entra authentication through
`DefaultAzureCredential`; no connection strings or secrets are stored in the
project.

## Prerequisites

- Python 3.9 or newer
- An existing Azure Service Bus namespace and **session-enabled queue**
- Azure Service Bus Data Sender and Azure Service Bus Data Receiver access
- A local identity supported by `DefaultAzureCredential`

The queue must have sessions enabled because each message uses the customer name
as its session ID. A processor locks and drains one customer session at a time,
which preserves FIFO order for that customer and prevents customer sessions from
being interleaved by one processor.

## Run

Install dependencies and set:

```text
SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net
SERVICEBUS_QUEUE_NAME=<session-enabled-queue>
```

Then run `python main.py`. The demo sends normal and high-priority orders,
delays orders worth more than $1,000 for 30 seconds, explicitly dead-letters an
invalid JSON message, inspects and resubmits it, and finally inspects and removes
the message after it fails again.

References:

- https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-queues
- https://learn.microsoft.com/python/api/overview/azure/servicebus-readme
