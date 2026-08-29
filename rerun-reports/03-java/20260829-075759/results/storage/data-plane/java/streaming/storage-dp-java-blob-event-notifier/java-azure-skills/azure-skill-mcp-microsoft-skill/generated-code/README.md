# Blob Event Notifier

Small Java 17 sample for receiving Azure Blob Storage lifecycle events through Event Grid and publishing downstream events.

The demo is intentionally offline: it parses realistic Event Grid and CloudEvents payloads, uses an in-memory blob store, and logs mock publishes. Production adapters use only `ManagedIdentityCredential`; no access keys or SAS tokens are accepted.

## Run locally

```powershell
mvn test
mvn exec:java
```

## Use Azure-backed adapters

Set:

- `AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net`
- `EVENT_GRID_TOPIC_ENDPOINT=https://<topic>.<region>-1.eventgrid.azure.net/api/events`
- `AZURE_CLIENT_ID=<user-assigned-managed-identity-client-id>` only for a user-assigned identity

Construct `AzureConfiguration.fromEnvironment()`, then obtain `blobStore()`, `eventPublisher()`, and `asyncEventPublisher()`. Assign the managed identity **Storage Blob Data Reader** on the required storage scope and **EventGrid Data Sender** on the custom topic scope.

Event Grid webhook authentication and subscription-validation handling belong in the hosting HTTP framework. Pass the validated request body to `EventReceiver.receive` or `AsyncEventReceiver.receive`.
