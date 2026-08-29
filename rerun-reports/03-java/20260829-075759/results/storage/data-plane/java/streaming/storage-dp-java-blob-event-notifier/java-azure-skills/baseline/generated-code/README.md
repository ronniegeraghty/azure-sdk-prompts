# Blob Event Notifier

Java 17 sample for receiving Azure Storage lifecycle events from Event Grid and publishing downstream custom events.

The demo is local-only by default:

```powershell
mvn compile exec:java
```

Production clients use `DefaultAzureCredential`, which supports managed identity without storage keys or SAS tokens. Configure `EVENT_GRID_TOPIC_ENDPOINT` to enable publishing and construct the blob handlers with clients from `AzureConfiguration`:

```java
AzureConfiguration azure = new AzureConfiguration();
BlobEventHandler syncHandler =
    new BlobEventHandler(azure.blobServiceClient("https://<account>.blob.core.windows.net"));
AsyncBlobEventHandler asyncHandler =
    new AsyncBlobEventHandler(azure.blobServiceAsyncClient("https://<account>.blob.core.windows.net"));
```
