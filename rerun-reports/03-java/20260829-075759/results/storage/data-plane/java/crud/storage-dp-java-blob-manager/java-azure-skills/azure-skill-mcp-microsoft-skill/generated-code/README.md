# Azure Blob Manager

A Java 17 utility with synchronous and asynchronous Azure Blob Storage operations.
It authenticates only with Azure Managed Identity and never accepts account keys or
connection strings.

## Configuration

| Environment variable | Default | Purpose |
| --- | --- | --- |
| `AZURE_STORAGE_ACCOUNT_URL` | required | HTTPS blob endpoint, for example `https://account.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER` | `blob-manager-demo` | Demo container |
| `AZURE_CLIENT_ID` | system-assigned identity | Client ID for a user-assigned managed identity |
| `AZURE_STORAGE_MAX_RETRIES` | `5` | Retries after the initial request |
| `AZURE_STORAGE_RETRY_DELAY_MS` | `800` | Initial exponential retry delay |
| `AZURE_STORAGE_MAX_RETRY_DELAY_MS` | `10000` | Maximum retry delay |
| `AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS` | `120` | Timeout for each HTTP request attempt |
| `AZURE_STORAGE_HTTP_LOG_LEVEL` | `BASIC` | `NONE`, `BASIC`, `HEADERS`, or `BODY_AND_HEADERS` |

The managed identity needs an appropriate data-plane role, such as **Storage Blob
Data Contributor**, scoped as narrowly as practical. The demo also creates its
container if absent, so its role must permit that operation.

## Run

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "blob-manager-demo"
mvn compile exec:java
```

Uploads use bounded-memory parallel block transfer. Existing blobs are updated with
an `If-Match` ETag condition and new blobs use `If-None-Match: *`, so a concurrent
writer causes a precondition failure instead of silently losing data. Lease IDs can
be supplied for deliberate lease-protected updates.

Reference: [Azure Storage Blob SDK for Java examples](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java)
