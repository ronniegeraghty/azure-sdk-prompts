# Azure Cosmos DB ToDo Repository

A Java 17 Maven sample with synchronous and asynchronous CRUD repositories for
the Azure Cosmos DB NoSQL API.

## Authentication and configuration

The sample uses `ManagedIdentityCredential`; it does not accept account keys.
Assign the managed identity a Cosmos DB data-plane role that includes item CRUD
and the metadata permissions needed to create the database and container.

Set the account endpoint before running:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
```

For a user-assigned managed identity, also set `AZURE_CLIENT_ID`. A system-assigned
managed identity is used when `AZURE_CLIENT_ID` is absent.

The factory creates `todo-db` and `todos` if needed. The container uses
`/category` as its partition key, a 90-day default TTL, and excludes
`/description/?` from indexing.

## Build and run

```powershell
mvn clean package
mvn exec:java
```

Updates require the ETag returned by the last read or write. Cosmos DB rejects a
stale ETag with HTTP 412, which the repositories expose as
`OptimisticConcurrencyException`.
