# Azure Blob Event Notifier

Java 17 Maven sample for receiving Azure Blob Storage lifecycle notifications in either Event Grid
schema or CloudEvents 1.0 schema, downloading created blobs, logging deleted blobs, and publishing
custom downstream Event Grid events.

The default demo is local-only: it uses realistic JSON payloads and injected fake download/publish
operations, so it does not authenticate or contact Azure.

## Run

```powershell
mvn test
mvn exec:java
```

## Connect the classes to Azure

`AzureConfiguration.fromEnvironment()` creates synchronous and asynchronous Blob Storage and Event
Grid clients with `ManagedIdentityCredential`. No access keys, connection strings, or SAS tokens are
used.

```text
AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net
EVENT_GRID_TOPIC_ENDPOINT=https://<topic>.<region>-1.eventgrid.azure.net/api/events
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

Construct production handlers and publishers from the configuration:

```java
AzureConfiguration config = AzureConfiguration.fromEnvironment();

EventReceiver receiver = new EventReceiver(new BlobEventHandler(config.blobServiceClient()));
EventPublisher publisher = config.eventPublisher();

AsyncEventReceiver asyncReceiver =
    new AsyncEventReceiver(new AsyncBlobEventHandler(config.blobServiceAsyncClient()));
AsyncEventPublisher asyncPublisher = config.asyncEventPublisher();
```

Grant the managed identity only the required data-plane roles, typically **Storage Blob Data Reader**
on the relevant storage scope and **EventGrid Data Sender** on the custom topic.

SDK references:

- https://learn.microsoft.com/java/api/overview/azure/messaging-eventgrid-readme
- https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme
- https://learn.microsoft.com/java/api/overview/azure/identity-readme
