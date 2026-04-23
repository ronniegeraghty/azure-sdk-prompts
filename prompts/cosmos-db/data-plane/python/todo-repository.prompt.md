---
id: cosmos-db-dp-python-todo-repository
properties:
  service: cosmos-db
  plane: data-plane
  language: python
  category: crud
  difficulty: intermediate
  description: 'Can an agent generate a Cosmos DB CRUD repository with optimistic concurrency (ETags), parameterized queries,
    page-by-page pagination, TTL configuration, custom indexing policy, and RU cost logging?

    '
  sdk_package: azure-cosmos
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/cosmos-readme
  created: '2026-04-10'
  author: copilot
tags:
- cosmos-db
- etag
- optimistic-concurrency
- pagination
- parameterized-query
- ttl
- indexing-policy
- request-charge
- async
---

# ToDo Repository: Azure Cosmos DB (Python)

## Prompt

Create a Python project that implements a ToDo item CRUD repository backed by Azure Cosmos DB (NoSQL API).

**Write the code to files (use file-write tools, do not reply with code blocks).**

The project needs:

- A **model** (shared by both implementations) for a ToDo item with fields for id, title, description, completed status, created timestamp, and category (where category is the partition key). Use a dictionary or a dataclass.

- A **synchronous repository module** that performs CRUD operations against Cosmos DB. It should support create, read, update, delete, and a query-by-category method. Each operation should log the request charge (RU cost consumed) from the response headers. The update operation should prevent lost updates — if another process modified the item since it was last read, the update should fail with a clear conflict error rather than silently overwriting the other process's changes. The query method should use safe, parameterized queries and must handle large result sets properly — paginate through results page by page rather than loading everything into memory at once, and log progress as each page is retrieved.

- An **asynchronous repository module** that provides the same CRUD operations using the `azure.cosmos.aio` async client. The query method should iterate through pages asynchronously.

- A **configuration/factory module** that connects to the Cosmos DB account using its endpoint from an environment variable. Authentication must use `DefaultAzureCredential` (no master keys). It should create the database and container if they don't already exist, setting a default TTL (time-to-live) of 90 days on the container and configuring the indexing policy to exclude the `description` field from indexing (since it's never queried on).

- A **main script** that demos both implementations: runs the full CRUD cycle using the sync repository first (including paginated query output showing page-by-page results), then runs the same operations using the async repository. Print RU costs and results to the console.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Scenario-Specific Patterns
- Correct partition key usage: `/category` path, `partition_key` in all point operations
- ETag-based optimistic concurrency: captures `_etag` from read response, passes `if_match` on replace
- Handles 412 Precondition Failed as a specific error case for conflicts
- Parameterized queries using `parameters` list (no f-string or format-string concatenation)
- Page-by-page iteration using `query_items(...).by_page()` with `max_item_count`
- Logs continuation token and item count per page
- Async version uses `azure.cosmos.aio.CosmosClient`
- TTL configured at 90 days (7776000 seconds) via `default_ttl` in container properties
- Indexing policy excludes `/description` path using `excludedPaths`
- RU cost extracted from response headers via `x-ms-request-charge` or `response_headers`

### Scenario-Specific Error Handling
- Catches `CosmosHttpResponseError` with status code checks (404, 409, 412)
- Handles 412 separately for ETag conflicts

### Anti-Patterns (scenario-specific)
- Does NOT flatten query results by calling `list()` without page iteration

## Context

This goes beyond basic Cosmos DB CRUD (covered by `crud-items.prompt.md`) to test production
patterns: optimistic concurrency with ETags to prevent lost updates, parameterized queries
to avoid injection, page-by-page pagination with continuation tokens for large result sets,
TTL configuration for automatic document expiry, and custom indexing policy to optimize storage
and write cost. The RU cost logging tests whether the agent knows how to extract request charges
from the Cosmos DB response — critical for cost monitoring in production.
