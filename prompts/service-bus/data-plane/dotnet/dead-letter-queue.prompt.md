---
id: service-bus-dp-dotnet-dead-letter
service: service-bus
plane: data-plane
language: dotnet
category: error-handling
difficulty: intermediate
description: >
  Can a developer handle dead-letter queue messages and implement
  retry-then-dead-letter patterns using the .NET SDK?
sdk_package: Azure.Messaging.ServiceBus
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/messaging.servicebus-readme
tags:
  - dead-letter
  - error-handling
  - poison-messages
  - reliability
created: 2026-03-28
author: jonathangiles
---

# Dead-Letter Queue Handling: Azure Service Bus (.NET)

## Prompt

Some messages in my Service Bus queue are failing to process and I need to
handle them properly. How do I work with the dead-letter queue in C#?
1. Dead-letter a message explicitly with a reason and description
2. Create a receiver for the dead-letter sub-queue
3. Read dead-lettered messages and inspect the dead-letter reason
4. Re-submit a dead-lettered message back to the main queue after fixing the issue

Show me the dead-letter queue path format and how to use
DeadLetterMessageAsync with the Azure.Messaging.ServiceBus SDK.

## Evaluation Criteria

- `ServiceBusReceiver.DeadLetterMessageAsync()` with reason and description
- Dead-letter sub-queue path: `$"{queueName}/$deadletterqueue"`
- `ServiceBusClient.CreateReceiver()` with `SubQueue.DeadLetter` option
- Accessing `DeadLetterReason` and `DeadLetterErrorDescription` properties
- Re-submitting messages via `ServiceBusSender.SendMessageAsync()`
- `MaxDeliveryCount` and automatic dead-lettering behavior

## Context

Dead-letter queue handling is a critical production pattern for Service Bus.
Messages that can't be processed (poison messages) must be triaged and either
fixed and re-submitted or discarded with an audit trail. This is a top question
from .NET developers building reliable messaging systems.
