# Azure Cosmos DB Java CRUD example

This example uses the Azure Cosmos DB Java SDK to create `TestDB` and its
`Items` container, then create, read, query, replace, and delete an item.
The container uses `/category` as its partition key.

The required SDK dependency is declared in `pom.xml`:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-cosmos</artifactId>
    <version>4.81.0</version>
</dependency>
```

Set credentials for an Azure Cosmos DB Emulator or another test endpoint.
Do not put keys in source control:

```powershell
$env:COSMOS_ENDPOINT = "https://localhost:8081"
$env:COSMOS_KEY = "<emulator-or-test-key>"
mvn compile exec:java
```

The sample catches `CosmosException`, reports its status code, substatus,
activity ID, message, and any service-provided retry delay, and exits with a
nonzero status.
