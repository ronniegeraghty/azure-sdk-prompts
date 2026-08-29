# Azure Blob Event Notifier

Small Java 17 Maven sample that receives Azure Blob Storage lifecycle events in either Event Grid schema or
CloudEvents 1.0 schema, handles blob-created and blob-deleted events, and publishes downstream custom events.

The `Main` demo is fully local and uses injected in-memory adapters. Production constructors use
`ManagedIdentityCredential`; no keys, connection strings, or SAS tokens are used.

## Run the offline demo

```powershell
mvn test
mvn exec:java
```

## Production configuration

Set these environment variables before constructing `AzureConfiguration.fromEnvironment()`:

| Variable | Purpose |
|---|---|
| `AZURE_STORAGE_BLOB_ENDPOINT` | Storage endpoint, such as `https://account.blob.core.windows.net` |
| `AZURE_EVENTGRID_TOPIC_ENDPOINT` | Full custom topic endpoint |
| `AZURE_CLIENT_ID` | Optional client ID for a user-assigned managed identity |

The managed identity needs **Storage Blob Data Reader** on the required storage scope and an Event Grid data-plane
sender role, such as **EventGrid Data Sender**, on the custom topic scope.

```java
AzureConfiguration configuration = AzureConfiguration.fromEnvironment();
BlobEventHandler handler = new BlobEventHandler(
    configuration.blobServiceClient(),
    configuration.blobServiceAsyncClient());

new EventReceiver().receive(webhookJson, handler);
configuration.eventPublisher().publish(customEvents);
```

The receiver expects one delivery payload to use one schema, matching Event Grid subscription configuration.
