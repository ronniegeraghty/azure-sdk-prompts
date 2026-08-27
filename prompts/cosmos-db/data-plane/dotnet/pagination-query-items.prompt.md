---
id: cosmos-db-dp-dotnet-pagination
properties:
  service: cosmos-db
  plane: data-plane
  language: dotnet
  category: pagination
  difficulty: intermediate
  description: 'Can a developer paginate through large Cosmos DB query results using continuation tokens in .NET?

    '
  sdk_package: Microsoft.Azure.Cosmos
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/microsoft.azure.cosmos-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- pagination
- query
- continuation-token
- feed-iterator
---

# Pagination: Query Items in Azure Cosmos DB (.NET)

## Prompt

I'm querying a Cosmos DB container with thousands of items and my query
returns too much data at once. How do I paginate results in .NET?
1. Execute a query against a container with a fixed page size
2. Process results page by page
3. Save a continuation token so I can resume the query later
4. Track request unit consumption across pages

Explain the difference between iterator-based and LINQ-based query
approaches.

## Evaluation Criteria

- Executes paginated queries with configurable page size
- Iterates through result pages using a feed iterator pattern
- Saves and resumes queries using continuation tokens
- Tracks request unit (RU) consumption per page
- Handles cross-partition query considerations
- Mentions LINQ-based query as an alternative

## Context

Cosmos DB pagination via FeedIterator is fundamentally different from
traditional database cursors. Developers must understand continuation tokens
and RU consumption per page to build efficient data access patterns.
