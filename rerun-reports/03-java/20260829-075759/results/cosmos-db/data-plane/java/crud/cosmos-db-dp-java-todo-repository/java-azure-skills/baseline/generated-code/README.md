# Cosmos DB ToDo repository

Java 17 sample with synchronous and asynchronous Azure Cosmos DB repositories.
Updates use Cosmos DB ETags to prevent lost updates, and category queries emit
one page at a time.

## Run

Run the application on an Azure host with a managed identity, then set the
account endpoint:

```powershell
$env:AZURE_COSMOS_ENDPOINT = "https://your-account.documents.azure.com:443/"
mvn compile exec:java
```

The identity needs permission to create the `todo-db` database and its `items`
container. The container uses `/category` as its partition key, a 90-day default
TTL, and an indexing policy that excludes `/description/?`.
