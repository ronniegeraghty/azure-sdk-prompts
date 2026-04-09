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

I need to set up pub/sub messaging with Azure Service Bus topics in C#.
How do I send messages to a topic and receive them from a subscription?
1. Create a ServiceBusSender for a topic and send a message
2. Create a ServiceBusReceiver for a specific subscription on that topic
3. Receive and complete messages from the subscription
4. Set up a ServiceBusProcessor for continuous subscription processing

Show the difference between queue-based and topic-based messaging with
the Azure.Messaging.ServiceBus SDK.

## Evaluation Criteria

- `ServiceBusClient.CreateSender(topicName)` for topic publishing
- `ServiceBusClient.CreateReceiver(topicName, subscriptionName)` for subscription
- `ServiceBusSender.SendMessageAsync()` to publish to topic
- `ServiceBusReceiver.ReceiveMessagesAsync()` from subscription
- `ServiceBusProcessor` with topic and subscription names
- `CompleteMessageAsync()` for message acknowledgment

## Context

Topics and subscriptions are the pub/sub model in Service Bus, enabling
fan-out to multiple consumers. This was split from the queue-based prompt
to test topic/subscription knowledge independently.
