# Azure Cosmos DB NoSQL CRUD sample

This .NET 8 console application creates `TestDB` and its `Items` container,
then creates, reads, queries, replaces, and deletes one item. The container
uses `/category` as its partition key.

## Required NuGet packages

```powershell
dotnet add package Microsoft.Azure.Cosmos --version 3.62.1
dotnet add package Newtonsoft.Json --version 13.0.4
```

`Microsoft.Azure.Cosmos` is the Azure Cosmos DB NoSQL SDK. `Newtonsoft.Json`
provides `JObject`, which lets the sample work directly with JSON documents.

## Run locally

Start the Azure Cosmos DB Emulator, copy its connection string, and set it in
an environment variable. Do not commit the connection string.

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-Cosmos-DB-Emulator-connection-string>"
dotnet run
```

The query uses a parameter rather than interpolating user input:

```sql
SELECT * FROM c WHERE c.category = @category
```

References:

- [Azure Cosmos DB .NET SDK](https://learn.microsoft.com/azure/cosmos-db/nosql/sdk-dotnet-v3)
- [Work with items using the .NET SDK](https://learn.microsoft.com/azure/cosmos-db/nosql/how-to-dotnet-create-item)
- [Azure Cosmos DB Emulator](https://learn.microsoft.com/azure/cosmos-db/emulator)
