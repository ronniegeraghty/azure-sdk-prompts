# Azure Cosmos DB Java CRUD example

This Maven project creates the `TestDB` database and an `Items` container partitioned
by `/category`, then creates, reads, queries, replaces, and deletes an item.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An Azure Cosmos DB for NoSQL endpoint and key

Set credentials as environment variables rather than placing them in source code:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
mvn compile exec:java
```

The required SDK dependency is:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-cosmos</artifactId>
    <version>4.81.0</version>
</dependency>
```

For production Azure-hosted applications, prefer Microsoft Entra ID and managed
identity over key authentication. Endpoint/key authentication is used here because
it is specifically required by the example.

References:

- [Azure Cosmos DB Java SDK overview](https://learn.microsoft.com/java/api/overview/azure/cosmos)
- [Azure Cosmos DB Java SDK samples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)
