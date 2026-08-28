# Cosmos DB CRUD sample

This .NET 9 console application uses the Azure Cosmos DB for NoSQL data-plane
SDK to create a database and container, then create, read, query, replace, and
delete an item.

## Required NuGet package

```powershell
dotnet add package Microsoft.Azure.Cosmos --version 3.62.1
dotnet add package Newtonsoft.Json --version 13.0.4
```

No Azure resources are provisioned by this project. Run it against an existing
account or a local Azure Cosmos DB emulator.

## Run

Set the connection string without placing credentials in source control:

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-connection-string>"
dotnet run
```

The account or emulator must already exist. The program idempotently creates
the `TestDB` database and the `Items` container with partition key `/category`
and 400 RU/s if they do not exist.

SDK reference:
<https://learn.microsoft.com/azure/cosmos-db/nosql/quickstart-dotnet>
