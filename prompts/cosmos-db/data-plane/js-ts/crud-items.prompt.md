---
id: cosmos-db-dp-js-ts-crud
properties:
  service: cosmos-db
  plane: data-plane
  language: js-ts
  category: crud
  difficulty: basic
  description: 'Can a developer create, read, query, and delete items in an Azure Cosmos DB container using the JavaScript/TypeScript
    SDK?

    '
  sdk_package: '@azure/cosmos'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/cosmos-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- cosmos-db
- nosql
- crud
- getting-started
---

# CRUD Items: Azure Cosmos DB (JavaScript/TypeScript)

## Prompt

Write a TypeScript program
that performs CRUD operations on items in an Azure Cosmos DB NoSQL container:
1. Create a CosmosClient using credential from `@azure/identity`
2. Create a database "TestDB" and container "Items" with partition key "/category"
3. Create an item with properties: id, category, name, quantity
4. Read the item back using item().read()
5. Query items where category equals "electronics" using parameterized query
6. Replace the item with updated quantity using item().replace()
7. Delete the item using item().delete()

Enable SDK diagnostic logging using `@azure/logger` with a configurable log level.
Handle errors using `RestError` from `@azure/core-rest-pipeline` with `statusCode` checks (e.g., 404 for not found).
Show required npm packages including `@azure/core-rest-pipeline`.

## Evaluation Criteria

The generated code should include:
- `@azure/cosmos` npm package
- `CosmosClient` constructor with credential from `@azure/identity`
- `client.databases.createIfNotExists()` and `database.containers.createIfNotExists()`
- `container.items.create()`, `container.item(id, partitionKey).read()`
- `container.items.query()` with `SqlQuerySpec`
- `container.item(id, partitionKey).replace()` and `.delete()`
- `FeedResponse` iteration and error status codes

## Context

The JavaScript Cosmos DB SDK uses a fluent chain pattern (container.item().read()).
This tests whether the generated code covers the chained resource model and the
FeedResponse pattern for query results.
