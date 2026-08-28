# Azure Cosmos DB CRUD sample

This .NET 8 console application uses the Azure Cosmos DB for NoSQL SDK to
create a database and container, then create, read, query, replace, and delete
an item.

## Required NuGet package

```powershell
dotnet add package Microsoft.Azure.Cosmos --version 3.47.0
dotnet add package Newtonsoft.Json --version 13.0.3
```

The packages are already declared in `CosmosCrudSample.csproj`.

## Run

Set a connection string for a local Cosmos DB emulator or another explicitly
approved environment. Do not commit the connection string.

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-cosmos-db-connection-string>"
dotnet run
```

The account represented by the connection string must allow database,
container, and item data-plane operations.
