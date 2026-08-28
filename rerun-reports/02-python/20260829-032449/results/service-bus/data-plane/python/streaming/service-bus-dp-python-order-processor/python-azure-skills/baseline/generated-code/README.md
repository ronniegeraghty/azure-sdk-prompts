# Azure Service Bus order processor

This example sends and processes orders with the synchronous and asynchronous
Azure Service Bus Python clients. The target queue **must have sessions
enabled**. Each message uses the customer name as its session ID, so one
processor owns and drains a customer's session before accepting another
customer's session.

## Run

1. Create and activate a virtual environment.
2. Install dependencies with `python -m pip install -r requirements.txt`.
3. Set `SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE` to a namespace such as
   `example.servicebus.windows.net`.
4. Optionally set `SERVICE_BUS_QUEUE_NAME` (the default is `orders`).
5. Authenticate with a credential supported by `DefaultAzureCredential`.
6. Run `python main.py`.

The demo sends valid and malformed messages, processes available queue
sessions, and inspects the dead-letter queue. Orders over $1,000 are marked
high priority and scheduled 30 seconds into the future.
