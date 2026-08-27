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
created: 2026-04-09
author: JonathanCrd
---

# Queue Operations: Azure Storage Queues (.NET)

## Prompt

How do I send and receive messages with Azure Storage Queues in .NET?
I need a simple message queue and don't want the complexity of Service Bus.
1. Connect to a queue and ensure it exists
2. Send a message to the queue
3. Receive and process messages
4. Delete a message after processing
5. Handle message visibility timeout

Authenticate securely using identity-based credentials. Explain the
difference between receiving and peeking at messages.

## Evaluation Criteria

- Connects to Azure Storage Queues using identity-based authentication
- Creates the queue if it doesn't exist
- Sends messages to the queue
- Distinguishes between receiving (dequeue) and peeking at messages
- Deletes messages after processing using message ID and pop receipt
- Handles visibility timeout for reliable processing

## Context

Azure Storage Queues are the simplest queuing option in Azure, ideal for
lightweight workloads. The SDK patterns are different from Service Bus and
developers need to understand the visibility timeout model for reliable processing.
