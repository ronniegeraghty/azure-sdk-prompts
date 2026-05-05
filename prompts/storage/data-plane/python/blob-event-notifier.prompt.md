---
id: storage-dp-python-blob-event-notifier
properties:
  service: storage
  plane: data-plane
  language: python
  category: streaming
  difficulty: intermediate
  description: 'Can an agent generate a Blob Storage event processor using Event Grid, supporting both EventGridEvent and
    CloudEvents 1.0 schemas, event routing by type, blob subject parsing, custom event publishing, and race condition handling?

    '
  sdk_package: azure-eventgrid
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/eventgrid-readme
  created: '2026-04-10'
  author: copilot
tags:
- event-grid
- blob-storage
- cloud-events
- event-routing
- async
- multi-service
---

# Blob Event Notifier: Azure Event Grid + Blob Storage (Python)

## Prompt

Create a Python project that processes Azure Blob Storage lifecycle events delivered via Event Grid.

**Write the code to files (use file-write tools, do not reply with code blocks).**

The project needs:

- An **event receiver module** (both sync and async versions) that accepts a JSON payload (as if received from an Event Grid webhook endpoint) and deserializes it into structured event objects using the Azure Event Grid SDK's built-in deserialization helpers — not manual JSON parsing. It should support both Event Grid native schema and CloudEvents 1.0 schema (since Event Grid supports both and the configured schema may vary). It should route events by type — blob-created events get processed one way, blob-deleted another, and unrecognized types are logged as warnings.

- A **blob event handler module** that processes individual blob events. For blob-created events, it should parse the blob's container and name from the event subject, download the blob, and print a summary (name, size, content type, and the blob's access tier). For blob-deleted events, it should just log the deletion. It should handle race conditions gracefully — the blob may have already been deleted or moved to a different tier by the time we try to read it.

- An **event publisher module** (both sync and async versions) that can publish custom events to an Event Grid topic. Given a topic endpoint and a list of custom event objects, it should send them to Event Grid using the SDK's publisher client. This would be used for downstream notifications (e.g., "document processed" events). It should support setting a subject hierarchy for event filtering (e.g., "/documents/invoices/processed"). Handle publishing errors gracefully with proper exception handling.

- A **configuration module** that connects to Azure Blob Storage and Event Grid securely. Authentication should use `DefaultAzureCredential` — no access keys or SAS tokens.

- A **main script** that demos both implementations: constructs a sample Event Grid JSON payload (with both CloudEvents and EventGrid-schema examples) containing mock blob-created and blob-deleted events with realistic structure, feeds them through the receiver and handler, and publishes a custom downstream event. Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Event Deserialization
- Deserializes Event Grid native schema events using the SDK's built-in deserialization (not manual JSON parsing)
- Deserializes CloudEvents 1.0 schema events using the SDK's built-in deserialization (not manual JSON parsing)

### Event Routing & Processing
- Routes events by event type (blob-created vs blob-deleted vs unrecognized)
- Logs a warning for unrecognized event types
- Parses container name and blob name from event subject
- Retrieves and prints blob access tier from blob properties

### Event Publishing
- Publishes custom events to an Event Grid topic using the SDK's publisher client
- Sets subject hierarchy on custom events for filtering

### Error Handling
- Handles race condition where the blob may no longer exist by the time the handler runs
- Handles publishing errors with proper exception handling

### Async Support
- Async versions use the async variants of the Event Grid and Blob Storage clients

## Context

This is a multi-service scenario testing Event Grid integration with Blob Storage. It exercises
the agent's knowledge of Event Grid's dual-schema support (EventGridEvent vs CloudEvents 1.0),
event deserialization using SDK helpers (not manual JSON parsing), event routing by type, blob
subject parsing, and the common race condition where a blob is deleted between the event firing
and the handler processing it. The event publishing side tests custom event creation with
subject hierarchies for downstream filtering.
