---
id: service-bus-dp-python-order-processor
properties:
  service: service-bus
  plane: data-plane
  language: python
  category: streaming
  difficulty: intermediate
  description: 'Can an agent generate a Service Bus order processing system with batch sending, scheduled delivery, dead-letter
    queue handling, session-aware processing, and proper error categorization?

    '
  sdk_package: azure-servicebus
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/servicebus-readme
  created: '2026-04-10'
  author: copilot
tags:
- service-bus
- batch-sending
- scheduled-delivery
- dead-letter-queue
- session-aware
- correlation
- async
---

# Order Processor: Azure Service Bus (Python)

## Prompt

Create a Python project that implements an order processing system using Azure Service Bus.

The project needs:

- A **model** for an Order with fields for order ID, customer name, product, quantity, total price, and status (pending/processing/completed/failed). It should be serializable to and from JSON. Use a dataclass or dictionary.

- A **sender module** (both sync and async versions) that publishes order messages to a Service Bus queue. It should support sending individual orders and sending a batch of orders efficiently (using `ServiceBusMessageBatch` and checking `add_message()` to respect the maximum batch size so messages aren't rejected). Each message should carry the order ID as a correlation property, and orders above a certain dollar threshold should be sent as high-priority with a scheduled delivery time of 30 seconds in the future (to allow for fraud review before processing).

- A **processor module** (both sync and async versions) that receives and processes orders from the queue. It should handle messages as they arrive, deserialize them, and log the results. If processing fails (e.g., a deserialization error), the message should be sent to the dead-letter queue with a reason string using `dead_letter_message()` rather than being silently abandoned. The processor should also be able to read from the dead-letter queue so failed messages can be inspected and reprocessed. It should guarantee that orders from the same customer are processed in sequence, not interleaved with other customers' orders.

- A **main script** that demos both implementations: connects to the Service Bus namespace (from an environment variable) with `DefaultAzureCredential`, runs the full send/receive/dead-letter cycle using the sync implementation first, then repeats with the async implementation.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Scenario-Specific Client Construction
- Sender uses `ServiceBusClient.get_queue_sender()` (or async equivalent)
- Processor uses `ServiceBusClient.get_queue_receiver()` or session-enabled receiver

### Scenario-Specific Patterns
- Batch sending: creates `ServiceBusMessageBatch`, checks `add_message()` return or catches `MessageSizeExceededError`
- Handles the case where a message doesn't fit in the current batch
- Scheduled delivery: uses `schedule_messages()` or `ServiceBusMessage(scheduled_enqueue_time_utc=...)` (~30s delay)
- Correlation: sets order ID via `correlation_id` property on `ServiceBusMessage`
- Dead-letter: explicitly dead-letters failed messages with `receiver.dead_letter_message()` and a reason string
- Dead-letter queue reading: uses `sub_queue=ServiceBusSubQueue.DEAD_LETTER` on receiver
- Session-aware processing: uses `session_id` on messages and session-enabled receiver
- Session ID keyed by customer name for ordered processing
- Context manager pattern (`with` statements) for all clients, senders, and receivers

### Scenario-Specific Error Handling
- Catches `ServiceBusError` and distinguishes transient vs non-transient errors
- Error handler logs entity name and error details

### Async Support
- Async versions use `azure.servicebus.aio` module

## Context

This goes beyond basic send/receive (covered by `send-receive-messages.prompt.md`) to test
production messaging patterns: batch sending with size-check guards, scheduled delivery for
fraud review workflows, dead-letter queue handling with reason strings for message forensics,
and session-aware processing to guarantee ordered delivery per customer. The session-aware
pattern is particularly important — without it, concurrent message processing can interleave
orders from the same customer, leading to race conditions in order fulfillment.
