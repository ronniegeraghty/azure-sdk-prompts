# Cosmos DB ToDo Repository

Small Java 17 sample with synchronous and asynchronous repositories for the
Azure Cosmos DB NoSQL API.

## Configuration

The application authenticates only with managed identity. Assign the identity
the Cosmos DB Built-in Data Contributor role for the account. For a
user-assigned managed identity, also set `AZURE_CLIENT_ID`.

```text
COSMOS_ENDPOINT=https://<account>.documents.azure.com:443/
COSMOS_DATABASE=todo-db       (optional)
COSMOS_CONTAINER=todos        (optional)
AZURE_CLIENT_ID=<identity-id> (optional)
```

The factory creates the database and container when absent. New containers use
`/category` as the partition key, expire items after 90 days, and exclude
`/description/?` from indexing.

## Build and run

```text
mvn clean package
mvn exec:java
```

Updates require the ETag returned by `read`, so Cosmos DB rejects stale writes
with HTTP 412 and the repository raises `ConcurrentUpdateException`.
