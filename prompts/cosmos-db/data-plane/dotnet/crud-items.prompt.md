---
id: cosmos-db-dp-dotnet-crud
properties:
  service: cosmos-db
  plane: data-plane
  language: dotnet
  category: crud
  difficulty: basic
  description: 'Can a developer create, read, query, and delete items in an Azure Cosmos DB container using the .NET SDK?

    '
  sdk_package: Microsoft.Azure.Cosmos
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/microsoft.azure.cosmos-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- cosmos-db
- nosql
- crud
- getting-started
---

# CRUD Items: Azure Cosmos DB (.NET)

## Prompt

How do I do basic CRUD operations against a Cosmos DB NoSQL container in .NET?
I have an existing database and container and need to:
1. Connect to Cosmos DB and get a reference to my container
2. Insert a JSON item with properties: id, category, name, and quantity
3. Read the item back by id and partition key
4. Query items where category equals a specific value

Include proper error handling for common failure scenarios.

## Evaluation Criteria

The generated code should include:
- Connects to Cosmos DB using the .NET SDK
- Creates, reads, and queries items in a container
- Uses parameterized queries for filtering
- Handles partition keys correctly
- Handles Cosmos DB-specific errors with status codes

## Context

Cosmos DB is one of the most popular Azure data services. CRUD operations test
whether the generated code covers the full item lifecycle including partitioning
and SQL-like query syntax in the NoSQL API.
