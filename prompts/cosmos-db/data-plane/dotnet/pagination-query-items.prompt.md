---
id: cosmos-db-dp-dotnet-pagination
service: cosmos-db
plane: data-plane
language: dotnet
category: pagination
difficulty: intermediate
description: >
  Can a developer paginate through large Cosmos DB query results
  using continuation tokens in .NET?
sdk_package: Microsoft.Azure.Cosmos
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/microsoft.azure.cosmos-readme
tags:
  - pagination
  - query
  - continuation-token
  - feed-iterator
created: 2025-07-27
author: ronniegeraghty
---

# Pagination: Query Items in Azure Cosmos DB (.NET)

## Prompt

I'm querying a Cosmos DB container with thousands of items and my query
returns too much data at once. How do I paginate results using FeedIterator?
1. Execute a SQL query against a container with MaxItemCount set to 50
2. Process results page-by-page using FeedIterator
3. Save the continuation token so I can resume the query later
4. Track total RU consumption across all pages

Use the Microsoft.Azure.Cosmos SDK v3. Show how to configure MaxItemCount
and explain the difference between FeedIterator and LINQ-based queries.

## Evaluation Criteria

- `Container.GetItemQueryIterator<T>()` with `QueryDefinition`
- `QueryRequestOptions.MaxItemCount` for page size control
- `FeedIterator<T>.HasMoreResults` and `ReadNextAsync()` loop pattern
- `FeedResponse<T>.ContinuationToken` for resumable pagination
- Passing continuation token to resume a query
- `FeedResponse<T>.RequestCharge` for RU tracking
- Cross-partition query considerations
- LINQ alternative via `GetItemLinqQueryable<T>()`

## Context

Cosmos DB pagination via FeedIterator is fundamentally different from
traditional database cursors. Developers must understand continuation tokens
and RU consumption per page to build efficient data access patterns.
