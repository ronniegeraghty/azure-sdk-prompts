# Azure Blob Manager

A Java 17 Maven example with reusable synchronous and asynchronous Azure Blob Storage services.
It authenticates with Azure managed identity and never uses account keys or connection strings.

## Configuration

| Environment variable | Required | Default | Purpose |
|---|---:|---:|---|
| `AZURE_STORAGE_ENDPOINT` | Yes | - | HTTPS blob endpoint, such as `https://account.blob.core.windows.net` |
| `AZURE_CLIENT_ID` | No | System-assigned identity | Client ID of a user-assigned managed identity |
| `AZURE_STORAGE_MAX_RETRIES` | No | `5` | Retries after the initial request |
| `AZURE_STORAGE_RETRY_DELAY_SECONDS` | No | `2` | Initial exponential retry delay |
| `AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS` | No | `30` | Maximum exponential retry delay |
| `AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS` | No | `120` | Per-request and HTTP I/O timeout |
| `AZURE_STORAGE_HTTP_LOG_LEVEL` | No | `BASIC` | Azure SDK HTTP log level: `NONE`, `BASIC`, `HEADERS`, or `BODY_AND_HEADERS` |

Assign the managed identity the least-privileged data-plane role needed by the application, normally
`Storage Blob Data Contributor`, scoped to the target container where possible. The container must
already exist.

## Build and run

```powershell
mvn clean package
$env:AZURE_STORAGE_ENDPOINT = "https://<account>.blob.core.windows.net"
mvn exec:java -Dexec.args="<existing-container>"
```

Uploads use path-based parallel block transfers with bounded buffers. The request timeout is applied
to each HTTP attempt, not to the total multi-part transfer. Existing blobs are updated
with an ETag `If-Match` condition, while new blobs use `If-None-Match: *`; conflicting concurrent
writes therefore fail instead of silently overwriting data. The lease demo adds a finite lease to
the conditional update.

## References

- [Azure Storage Blob client library for Java](https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme)
- [Authenticate Java apps to Azure services](https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview)
- [Manage concurrency in Blob Storage](https://learn.microsoft.com/azure/storage/blobs/concurrency-manage)
- [Manage blob leases with Java](https://learn.microsoft.com/azure/storage/blobs/storage-blob-lease-java)
