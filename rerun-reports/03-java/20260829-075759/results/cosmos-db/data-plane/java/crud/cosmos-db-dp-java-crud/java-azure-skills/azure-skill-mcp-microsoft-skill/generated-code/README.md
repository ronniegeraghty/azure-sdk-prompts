# Azure Cosmos DB Java CRUD sample

This Maven project demonstrates synchronous CRUD operations against a local Azure Cosmos DB
NoSQL emulator. It creates `TestDB`, creates the `Items` container with `/category` as its
partition key, and then creates, reads, queries, replaces, and deletes an item.

## Configuration

Set the emulator endpoint and key without placing credentials in source control:

```powershell
$env:COSMOS_ENDPOINT = "https://localhost:8081"
$env:COSMOS_KEY = "<local-emulator-key>"
```

The program rejects non-loopback endpoints so that it cannot modify a live Azure account.

## Build and run

```powershell
mvn compile
mvn exec:java
```

The required Azure SDK dependency is:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-cosmos</artifactId>
    <version>4.82.0</version>
</dependency>
```

References:

- [Azure Cosmos DB Java SDK examples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)
- [Azure Cosmos DB Java SDK API documentation](https://azuresdkdocs.z19.web.core.windows.net/java/azure-cosmos/latest/index.html)
