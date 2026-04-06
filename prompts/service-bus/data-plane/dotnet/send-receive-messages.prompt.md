---
id: service-bus-dp-dotnet-crud
service: service-bus
plane: data-plane
language: dotnet
category: crud
difficulty: intermediate
description: >
  Can a developer send and receive messages using Azure Service Bus
  queues and topics with the .NET SDK?
sdk_package: Azure.Messaging.ServiceBus
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/messaging.servicebus-readme
tags:
  - service-bus
  - messaging
  - queues
  - topics
created: 2025-07-28
author: ronniegeraghty
---

# Send and Receive Messages: Azure Service Bus (.NET)

## Prompt

How do I send and receive messages from an Azure Service Bus queue in C#?
I need to:
1. Create a ServiceBusClient using a connection string
2. Send a single message and a batch of messages to a queue
3. Receive and complete messages using ServiceBusReceiver
4. Set up a ServiceBusProcessor for continuous processing with handlers
5. Handle proper disposal with await using

Show required NuGet packages.

## Evaluation Criteria

The generated code should include:
- `Azure.Messaging.ServiceBus` NuGet package
- `ServiceBusClient` creation with connection string or `DefaultAzureCredential`
- `ServiceBusSender` and `ServiceBusMessage` for sending
- `ServiceBusMessageBatch` and `TryAddMessage()`
- `ServiceBusReceiver` and `ReceiveMessagesAsync()`
- `CompleteMessageAsync()`, `AbandonMessageAsync()`
- `ServiceBusProcessor` with `ProcessMessageAsync` and `ProcessErrorAsync`

## Context

Service Bus is Azure's enterprise messaging service supporting queues and pub/sub topics.
This tests both pull-based receiving and the processor pattern, plus the queue vs.
topic distinction that is fundamental to Service Bus architecture.
