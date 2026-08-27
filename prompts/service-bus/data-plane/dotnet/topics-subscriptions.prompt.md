---
id: service-bus-dp-dotnet-topics
service: service-bus
plane: data-plane
language: dotnet
category: streaming
difficulty: intermediate
description: >
  Can a developer use Azure Service Bus topics and subscriptions for
  pub/sub messaging using the .NET SDK?
sdk_package: Azure.Messaging.ServiceBus
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/messaging.servicebus-readme
tags:
  - topics
  - subscriptions
  - pub-sub
  - messaging
created: 2026-04-09
author: JonathanCrd
---

# Topics & Subscriptions: Azure Service Bus (.NET)

## Prompt

I need to set up pub/sub messaging with Azure Service Bus topics in .NET.
How do I:
1. Publish a message to a topic
2. Receive messages from a specific subscription on that topic
3. Acknowledge messages after processing
4. Set up continuous processing for a subscription

Explain how topic-based messaging differs from queue-based messaging.

## Evaluation Criteria

- Publishes messages to a Service Bus topic
- Receives messages from a named subscription
- Acknowledges (completes) messages after successful processing
- Supports continuous processing via a processor or handler pattern
- Distinguishes between queue-based and topic/subscription-based messaging

## Context

Topics and subscriptions are the pub/sub model in Service Bus, enabling
fan-out to multiple consumers. This was split from the queue-based prompt
to test topic/subscription knowledge independently.
