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
created: 2026-04-09
author: JonathanCrd
---

# Dead-Letter Queue Handling: Azure Service Bus (.NET)

## Prompt

Some messages in my Service Bus queue are failing to process and I need to
handle them properly. How do I work with the dead-letter queue in .NET?
1. Dead-letter a message explicitly with a reason and description
2. Receive messages from the dead-letter sub-queue
3. Inspect why a message was dead-lettered
4. Re-submit a corrected message back to the main queue

Explain how automatic dead-lettering works when delivery attempts are
exhausted.

## Evaluation Criteria

- Explicitly dead-letters a message with a reason and error description
- Receives messages from the dead-letter sub-queue
- Inspects dead-letter reason and error description properties on messages
- Re-submits messages to the main queue after correction
- Explains automatic dead-lettering based on max delivery count

## Context

Dead-letter queue handling is a critical production pattern for Service Bus.
Messages that can't be processed (poison messages) must be triaged and either
fixed and re-submitted or discarded with an audit trail. This is a top question
from .NET developers building reliable messaging systems.
