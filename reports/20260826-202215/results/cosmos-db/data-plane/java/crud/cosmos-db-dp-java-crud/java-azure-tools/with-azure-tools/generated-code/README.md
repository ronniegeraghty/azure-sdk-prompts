# Azure Cosmos DB Java CRUD sample

This Maven project uses the synchronous Azure Cosmos DB Java SDK to:

1. Create the `TestDB` database and the `Items` container if they do not exist.
2. Create, point-read, query, replace, and delete an item.
3. Use `/category` as the partition key and a parameterized SQL query.
4. Report Cosmos DB status, retry, activity, and diagnostic information when a request fails.

## Configuration

Set credentials through environment variables. For local-only development, use values from
the Azure Cosmos DB emulator rather than credentials for a live account.

```powershell
$env:COSMOS_ENDPOINT = "https://localhost:8081/"
$env:COSMOS_KEY = "<COSMOS-DB-EMULATOR-KEY>"
```

Do not commit keys to source control.

## Build and run

```powershell
mvn clean package
mvn exec:java -Dexec.mainClass=com.example.cosmos.CosmosCrudExample
```

The required Maven dependency is:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-cosmos</artifactId>
    <version>4.81.0</version>
</dependency>
```

## References

- [Azure Cosmos DB Java SDK overview](https://learn.microsoft.com/java/api/overview/azure/cosmos-readme)
- [Azure Cosmos DB Java SDK v4 samples](https://github.com/Azure-Samples/azure-cosmos-java-sql-api-samples)
- [Troubleshoot the Azure Cosmos DB Java SDK v4](https://learn.microsoft.com/azure/cosmos-db/troubleshoot-java-sdk-v4)
