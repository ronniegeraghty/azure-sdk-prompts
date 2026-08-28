# Azure Cosmos DB CRUD sample

This .NET 8 console application uses the local Azure Cosmos DB Emulator. It
refuses non-loopback endpoints so it cannot modify a live Azure account.

## Required NuGet package

```powershell
dotnet add package Microsoft.Azure.Cosmos --version 3.47.0
dotnet add package Newtonsoft.Json --version 13.0.3
```

Both packages are already declared in `CosmosCrudSample.csproj`.

## Run

Start the Azure Cosmos DB Emulator, copy its connection string, and run:

```powershell
$env:COSMOS_CONNECTION_STRING = "<local-emulator-connection-string>"
dotnet run
```

The program creates `TestDB` and the `Items` container in the emulator, creates
an item, reads and queries it, updates its quantity, and deletes it.
