---
id: storage-dp-dotnet-queues
service: storage
plane: data-plane
language: dotnet
category: crud
difficulty: basic
description: >
  Can a developer send and receive messages using Azure Storage Queues
  with the .NET SDK?
sdk_package: Azure.Storage.Queues
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.queues-readme
tags:
  - queues
  - messaging
  - storage-queues
  - getting-started
created: 2026-03-28
author: jonathangiles
---

# Queue Operations: Azure Storage Queues (.NET)

## Prompt

How do I send and receive messages with Azure Storage Queues in C#?
I need a simple message queue and don't want the complexity of Service Bus.
1. Create a QueueClient and ensure the queue exists
2. Send a message to the queue
3. Receive and process messages (peek vs. dequeue)
4. Delete a message after processing
5. Handle message visibility timeout

Use DefaultAzureCredential for auth. Show required NuGet packages and
explain the difference between receiving and peeking at messages.

## Evaluation Criteria

- `Azure.Storage.Queues` NuGet package
- `QueueClient` creation with URI and `DefaultAzureCredential`
- `QueueClient.CreateIfNotExistsAsync()`
- `SendMessageAsync()` for enqueueing
- `ReceiveMessagesAsync()` vs `PeekMessagesAsync()`
- `DeleteMessageAsync()` with message ID and pop receipt
- `visibilityTimeout` parameter on receive
- Base64 encoding behavior (default in v12)

## Context

Azure Storage Queues are the simplest queuing option in Azure, ideal for
lightweight workloads. The SDK patterns are different from Service Bus and
developers need to understand the visibility timeout model for reliable processing.
