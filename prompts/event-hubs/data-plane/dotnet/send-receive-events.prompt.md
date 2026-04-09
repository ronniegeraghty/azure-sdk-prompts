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
and consumer sides in C#?
1. Create an EventHubProducerClient and send a batch of events
2. Use EventProcessorClient with Blob Storage checkpointing to receive events
3. Register ProcessEventAsync and ProcessErrorAsync handlers
4. Implement checkpointing with ProcessEventArgs.UpdateCheckpointAsync()
5. Handle proper disposal of clients

Show required NuGet packages (Azure.Messaging.EventHubs and
Azure.Messaging.EventHubs.Processor).

## Evaluation Criteria

The generated code should include:
- `Azure.Messaging.EventHubs` and `Azure.Messaging.EventHubs.Processor` NuGet packages
- `EventHubProducerClient` with `CreateBatchAsync()` and `SendAsync()`
- `EventDataBatch.TryAdd()` for adding events to a batch
- `EventProcessorClient` with `BlobContainerClient` for checkpointing
- Event handler delegates and `ProcessEventArgs`
- `UpdateCheckpointAsync()` for reliable processing

## Context

Event Hubs is Azure's high-throughput event streaming service. The producer/consumer
pattern with checkpointing is the core usage model. This tests whether the generated code
covers both sides of the pipeline with proper checkpoint storage.
