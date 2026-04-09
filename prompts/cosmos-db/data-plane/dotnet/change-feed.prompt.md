---
id: cosmos-db-dp-dotnet-change-feed
service: cosmos-db
plane: data-plane
language: dotnet
category: streaming
difficulty: intermediate
description: >
  Can a developer consume the Cosmos DB change feed to react to
  item-level changes using the .NET SDK?
sdk_package: Microsoft.Azure.Cosmos
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/microsoft.azure.cosmos-readme
tags:
  - change-feed
  - streaming
  - event-driven
  - real-time
created: 2026-04-09
author: JonathanCrd
---

# Change Feed: Azure Cosmos DB (.NET)

## Prompt

I want to react to changes in my Cosmos DB container in near-real-time.
How do I set up a change feed processor in C#?
1. Create a change feed processor using the builder pattern
2. Configure a lease container for tracking progress
3. Handle incoming changes in a delegate
4. Start and stop the processor gracefully

Use the Microsoft.Azure.Cosmos SDK v3. Show how the lease container works
and what happens when my processor restarts — does it replay from the beginning?

## Evaluation Criteria

- `Container.GetChangeFeedProcessorBuilder<T>()` builder pattern
- Lease container configuration with `WithLeaseContainer()`
- `WithInstanceName()` and `WithStartTime()` or `WithStartFromBeginning()`
- Change handler delegate: `ChangesHandler<T>`
- `ChangeFeedProcessor.StartAsync()` and `StopAsync()`
- Explanation of lease-based progress tracking and multi-instance scaling

## Context

The Cosmos DB change feed is one of the most asked-about features for building
event-driven architectures. This tests whether the generated code demonstrates
the processor pattern correctly, including lease management for reliable delivery.
