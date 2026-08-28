# Azure Cosmos DB CRUD sample

This console application creates the `TestDB` database and an `Items` container
partitioned by `/category`, then creates, reads, queries, replaces, and deletes
one item.

## Required NuGet package

```powershell
dotnet add package Microsoft.Azure.Cosmos
dotnet add package Newtonsoft.Json
```

## Run

Set the connection string through the environment rather than storing it in
source control. For local development, use the Azure Cosmos DB Emulator:

```powershell
$env:COSMOS_CONNECTION_STRING = "<local-cosmos-db-emulator-connection-string>"
dotnet run
```

The identity represented by the connection string must be allowed to create
databases and containers and perform item operations.
