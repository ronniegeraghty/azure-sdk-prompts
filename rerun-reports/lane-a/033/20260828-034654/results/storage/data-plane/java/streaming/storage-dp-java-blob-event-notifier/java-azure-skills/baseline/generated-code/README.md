# Azure Blob Event Notifier

Small Java 17 example for receiving Azure Storage lifecycle events from Event Grid, handling blobs, and publishing downstream custom events.

## Run the local demo

The demo uses in-memory blob metadata and a logging Event Grid sender, so it does not contact Azure:

```powershell
mvn clean test
mvn exec:java
```

## Use with Azure

Create production components with HTTPS endpoints:

```java
AzureConfiguration azure = new AzureConfiguration(
    "https://<account>.blob.core.windows.net");

EventReceiver receiver = new EventReceiver(azure.blobEventHandler());
AsyncEventReceiver asyncReceiver = new AsyncEventReceiver(azure.blobEventHandler());
EventPublisher publisher = azure.eventPublisher(
    "https://<topic>.<region>-1.eventgrid.azure.net/api/events");
AsyncEventPublisher asyncPublisher = azure.asyncEventPublisher(
    "https://<topic>.<region>-1.eventgrid.azure.net/api/events");
```

`DefaultAzureCredential` is used throughout. In Azure, assign the workload's managed identity the least-privilege roles it needs, such as **Storage Blob Data Reader** on the storage account and **EventGrid Data Sender** on the custom topic. No account keys or SAS tokens are accepted by this example.
