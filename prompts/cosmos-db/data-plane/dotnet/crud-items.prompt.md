---
id: cosmos-db-dp-dotnet-crud
service: cosmos-db
plane: data-plane
language: dotnet
category: crud
difficulty: basic
description: >
  Can a developer create, read, query, and delete items in an Azure Cosmos DB
  container using the .NET SDK?
sdk_package: Microsoft.Azure.Cosmos
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/microsoft.azure.cosmos-readme
tags:
  - cosmos-db
  - nosql
  - crud
  - getting-started
created: 2025-07-28
author: ronniegeraghty
---

# CRUD Items: Azure Cosmos DB (.NET)

## Prompt

How do I do basic CRUD operations against a Cosmos DB NoSQL container in C#?
I have an existing database and container and need to:
1. Create a CosmosClient and get a reference to my container
2. Insert a JSON item with properties: id, category, name, and quantity
3. Read the item back by id and partition key
4. Query items where category equals "electronics" using SQL-like syntax

Show required NuGet packages and proper error handling with CosmosException.

## Evaluation Criteria

The generated code should include:
- `Microsoft.Azure.Cosmos` NuGet package
- `CosmosClient` creation and configuration
- `Container.CreateItemAsync<T>()`, `ReadItemAsync<T>()`
- `Container.GetItemQueryIterator<T>()` with `QueryDefinition`
- `CosmosException` handling with status codes

## Context

Cosmos DB is one of the most popular Azure data services. CRUD operations test
whether the generated code covers the full item lifecycle including partitioning
and SQL-like query syntax in the NoSQL API.
