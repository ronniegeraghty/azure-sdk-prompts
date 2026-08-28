# Azure Blob Manager

A Java 17 Maven example that provides synchronous and asynchronous Azure Blob
Storage utilities. Authentication uses Azure managed identity; connection strings
and storage account keys are not supported.

## Configuration

The demo expects an existing container and these environment variables:

| Variable | Required | Default | Description |
|---|---:|---:|---|
| `AZURE_STORAGE_ACCOUNT_ENDPOINT` | Yes | - | HTTPS blob endpoint, such as `https://account.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER` | Yes | - | Existing container used by the demo |
| `AZURE_CLIENT_ID` | No | System-assigned identity | Client ID of a user-assigned managed identity |
| `AZURE_STORAGE_MAX_RETRIES` | No | `5` | Retries after the initial request |
| `AZURE_STORAGE_RETRY_DELAY_MS` | No | `800` | Initial exponential retry delay |
| `AZURE_STORAGE_MAX_RETRY_DELAY_MS` | No | `10000` | Maximum exponential retry delay |
| `AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS` | No | `120` | Timeout applied to each HTTP attempt |
| `AZURE_STORAGE_HTTP_LOG_LEVEL` | No | `BASIC` | `NONE`, `BASIC`, `HEADERS`, or `BODY_AND_HEADERS` |

The managed identity needs an appropriate data-plane role, such as **Storage Blob
Data Contributor**, scoped as narrowly as practical. Avoid logging headers or
bodies outside controlled debugging because they can contain sensitive data.

## Build and run

```powershell
mvn clean package
mvn exec:java
```

The demo deliberately does not create a container. It uploads separate sync and
async sample blobs, lists the container, downloads each blob, conditionally
overwrites it while holding a lease, and deletes it.

Uploads use the SDK's file-based parallel block upload. Each write is conditional:
new blobs require nonexistence, while updates require the caller's expected ETag.
For coordinated updates, pass both the expected ETag and a lease ID.
