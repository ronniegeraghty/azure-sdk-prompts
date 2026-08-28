# Azure Blob Event Notifier

A Java 17 Maven sample that receives Azure Blob Storage lifecycle events in
Event Grid or CloudEvents 1.0 schema, downloads created blobs, logs deleted
blobs, and publishes downstream Event Grid events.

The demo is local-only: it uses in-memory blob metadata and dry-run publishers.

```powershell
mvn compile exec:java
```

For Azure-hosted use, construct `AzureConfiguration` from environment-provided
HTTPS endpoints, then inject its sync or async clients into `BlobEventHandler`
and the matching publisher:

```java
var config = new AzureConfiguration(
    System.getenv("AZURE_STORAGE_BLOB_ENDPOINT"),
    System.getenv("AZURE_EVENT_GRID_TOPIC_ENDPOINT"),
    System.getenv("AZURE_CLIENT_ID")
);
```

`AZURE_CLIENT_ID` is optional for a system-assigned identity and identifies a
user-assigned managed identity when set. Grant the identity **Storage Blob Data
Reader** on the storage account/container and **EventGrid Data Sender** on the
custom topic. No account keys, topic keys, connection strings, or SAS tokens
are used.
