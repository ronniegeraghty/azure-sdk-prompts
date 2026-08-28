# Cosmos DB ToDo Repository

Small Java 17 sample with synchronous and asynchronous Azure Cosmos DB for
NoSQL repositories.

## Configuration

The sample authenticates only with an Azure managed identity. Set:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:AZURE_CLIENT_ID = "<user-assigned-managed-identity-client-id>" # optional
$env:COSMOS_DATABASE = "todo-db"                                    # optional
$env:COSMOS_CONTAINER = "todos"                                     # optional
```

The managed identity needs a Cosmos DB data-plane role that permits item,
metadata, database, and container operations. Run the demo from a managed
Azure host:

```powershell
mvn compile exec:java
```

The container uses `/category` as its partition key, a 90-day default TTL, and
an indexing policy that excludes `/description/?`.
