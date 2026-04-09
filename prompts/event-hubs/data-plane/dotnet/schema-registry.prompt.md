---
id: event-hubs-dp-dotnet-schema-registry
service: event-hubs
plane: data-plane
language: dotnet
category: streaming
difficulty: advanced
description: >
  Can a developer use Azure Schema Registry with Event Hubs to serialize
  and deserialize events using Avro schemas in .NET?
sdk_package: Azure.Messaging.EventHubs
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/messaging.eventhubs-readme
tags:
  - schema-registry
  - avro
  - serialization
  - enterprise
created: 2026-04-09
author: JonathanCrd
---

# Schema Registry with Event Hubs: Azure Event Hubs (.NET)

## Prompt

I'm using Event Hubs in an enterprise setting and need to enforce schemas
on my events. How do I use Azure Schema Registry with Avro serialization in C#?
1. Register an Avro schema with the Schema Registry client
2. Serialize an event using SchemaRegistryAvroSerializer before sending
3. Send the serialized event to Event Hubs
4. Deserialize the received event back using the schema from the registry

Show the required NuGet packages (Microsoft.Azure.Data.SchemaRegistry.ApacheAvro
and Azure.Data.SchemaRegistry) and explain the schema ID embedding model.

## Evaluation Criteria

- `Azure.Data.SchemaRegistry` and `Microsoft.Azure.Data.SchemaRegistry.ApacheAvro` NuGet packages
- `SchemaRegistryClient` creation with endpoint and credential
- `SchemaRegistryAvroSerializer` for Avro serialization/deserialization
- `SerializeAsync()` to create `EventData` with embedded schema ID
- `DeserializeAsync<T>()` to reconstruct the typed object
- Schema group and schema name concepts

## Context

Schema Registry is critical for enterprise Event Hubs deployments where
producers and consumers need a shared contract. The Avro serializer
handles schema evolution and ID embedding transparently, but the setup
is non-obvious and frequently asked about.
