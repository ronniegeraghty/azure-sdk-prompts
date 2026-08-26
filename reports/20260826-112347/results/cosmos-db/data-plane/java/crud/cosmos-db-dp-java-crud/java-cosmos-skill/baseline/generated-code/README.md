# Azure Cosmos DB Java CRUD example

This Maven project creates a synchronous `CosmosClient` and demonstrates create,
read, parameterized query, replace, and delete operations in an Azure Cosmos DB
for NoSQL container.

## Configuration

Set the account endpoint and key as environment variables. Do not store account
keys in source control.

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
```

## Run

```powershell
mvn compile exec:java
```

The program creates `TestDB` and its `Items` container if they do not exist. The
container uses `/category` as its partition key.

Key-based authentication is used because this example specifically demonstrates
the `CosmosClientBuilder.endpoint(...).key(...)` flow. For production workloads
hosted in Azure, prefer Microsoft Entra ID and managed identity when possible.

## References

- [Azure Cosmos DB client library for Java](https://learn.microsoft.com/java/api/overview/azure/cosmos-readme?view=azure-java-stable)
- [Azure Cosmos DB Java SDK v4 troubleshooting](https://learn.microsoft.com/azure/cosmos-db/troubleshoot-java-sdk-v4)
