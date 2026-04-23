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
How do I consume the change feed in .NET?
1. Set up a processor that listens for item-level changes
2. Track progress so I don't reprocess events after a restart
3. Handle incoming changes as they arrive
4. Start and stop the processor gracefully

Explain how progress tracking works and what happens when my processor
restarts — does it replay from the beginning?

## Evaluation Criteria

- Sets up a change feed processor using the Cosmos DB SDK
- Configures a lease container for progress tracking
- Handles incoming changes via a delegate or handler
- Supports starting, stopping, and restarting the processor
- Explains lease-based progress tracking and multi-instance scaling

## Context

The Cosmos DB change feed is one of the most asked-about features for building
event-driven architectures. This tests whether the generated code demonstrates
the processor pattern correctly, including lease management for reliable delivery.
