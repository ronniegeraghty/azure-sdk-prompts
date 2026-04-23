---
id: service-bus-dp-dotnet-crud
properties:
  service: service-bus
  plane: data-plane
  language: dotnet
  category: crud
  difficulty: intermediate
  description: 'Can a developer send and receive messages using Azure Service Bus queues and topics with the .NET SDK?

    '
  sdk_package: Azure.Messaging.ServiceBus
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/messaging.servicebus-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- service-bus
- messaging
- queues
- topics
---

# Send and Receive Messages: Azure Service Bus (.NET)

## Prompt

How do I send and receive messages from an Azure Service Bus queue in .NET?
I need to:
1. Connect to Service Bus securely
2. Send a single message and a batch of messages to a queue
3. Receive and acknowledge messages
4. Set up continuous message processing with error handling
5. Ensure proper resource cleanup

## Evaluation Criteria

The generated code should include:
- Connects to Service Bus with identity-based authentication or connection string
- Sends individual and batched messages with size validation
- Receives messages and acknowledges them after processing
- Supports continuous processing via a processor or handler pattern
- Handles processing errors and properly disposes of clients

## Context

Service Bus is Azure's enterprise messaging service supporting queues and pub/sub topics.
This tests both pull-based receiving and the processor pattern, plus the queue vs.
topic distinction that is fundamental to Service Bus architecture.
