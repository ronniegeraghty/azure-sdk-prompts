---
id: event-hubs-dp-dotnet-streaming
properties:
  service: event-hubs
  plane: data-plane
  language: dotnet
  category: streaming
  difficulty: intermediate
  description: 'Can a developer send and receive events using Azure Event Hubs with the .NET SDK?

    '
  sdk_package: Azure.Messaging.EventHubs
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/messaging.eventhubs-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- event-hubs
- streaming
- producer
- consumer
---

# Send and Receive Events: Azure Event Hubs (.NET)

## Prompt

I need to send a batch of events to Azure Event Hubs and then process
them reliably with checkpointing. How do I set up both the producer
and consumer sides in .NET?
1. Send a batch of events to an Event Hub
2. Receive and process events reliably with checkpoint-based tracking
3. Handle processing errors gracefully
4. Ensure proper resource cleanup

Explain how checkpointing works and why it's needed for reliable
event processing.

## Evaluation Criteria

The generated code should include:
- Sends events in batches with size validation
- Receives events with checkpoint-based progress tracking
- Uses a durable checkpoint store (e.g., blob storage)
- Registers event and error handlers
- Checkpoints after processing to avoid reprocessing
- Properly disposes of clients and resources

## Context

Event Hubs is Azure's high-throughput event streaming service. The producer/consumer
pattern with checkpointing is the core usage model. This tests whether the generated code
covers both sides of the pipeline with proper checkpoint storage.
