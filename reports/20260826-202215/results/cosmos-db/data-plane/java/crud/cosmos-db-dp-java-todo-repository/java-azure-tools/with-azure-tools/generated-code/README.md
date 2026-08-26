# Cosmos DB ToDo Repository

Java 17 Maven sample with synchronous and asynchronous repositories for the
Azure Cosmos DB for NoSQL API.

## Configuration

The account identity needs Cosmos DB data-plane permissions for item CRUD and
queries, plus permission to create the database and container. The sample uses
managed identity only; it does not accept account keys.

Set the account endpoint:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
```

For a user-assigned managed identity, also set its client ID. Leave this unset
to use the system-assigned identity:

```powershell
$env:AZURE_CLIENT_ID = "<managed-identity-client-id>"
```

Run the demo:

```powershell
mvn compile exec:java
```

On first use, the factory creates `todo-db/items` with `/category` as its
partition key, a 90-day default TTL, and `/description/?` excluded from the
index. Existing containers are left unchanged by Cosmos DB's
`createContainerIfNotExists` operation.

The model carries the Cosmos DB `_etag` returned by reads and writes. Repository
updates send that ETag as an `If-Match` condition and report HTTP 412 responses
as `OptimisticConcurrencyException`, preventing lost updates.

## References

- [Azure Cosmos DB Java SDK samples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)
- [Azure Cosmos DB Java SDK v4 performance tips](https://learn.microsoft.com/azure/cosmos-db/nosql/performance-tips-java-sdk-v4)
- [Azure Identity credential chains for Java](https://aka.ms/azsdk/java/identity/credential-chains)
