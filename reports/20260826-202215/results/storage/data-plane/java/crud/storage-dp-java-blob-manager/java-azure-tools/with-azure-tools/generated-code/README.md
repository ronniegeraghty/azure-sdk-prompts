# Azure Blob Manager

A Java 17 Maven example that exposes reusable synchronous and asynchronous Azure Blob Storage
services. Authentication uses Azure managed identity only; no account keys or connection strings
are accepted.

## Configuration

| Environment variable | Required | Default | Purpose |
|---|---:|---:|---|
| `AZURE_STORAGE_ACCOUNT_ENDPOINT` | Yes | - | HTTPS endpoint, for example `https://account.blob.core.windows.net` |
| `AZURE_CLIENT_ID` | No | System-assigned identity | Client ID of a user-assigned managed identity |
| `AZURE_STORAGE_CONTAINER` | No | `blob-manager-demo` | Container used by `Main` |
| `BLOB_MAX_RETRIES` | No | `5` | Maximum exponential-backoff retries |
| `BLOB_RETRY_BASE_DELAY_SECONDS` | No | `1` | Initial retry delay |
| `BLOB_RETRY_MAX_DELAY_SECONDS` | No | `30` | Maximum retry delay |
| `BLOB_REQUEST_TIMEOUT_SECONDS` | No | `60` | HTTP connect/read/write/response timeout per request |
| `BLOB_HTTP_LOG_LEVEL` | No | `BASIC` | `NONE`, `BASIC`, `HEADERS`, or `BODY_AND_HEADERS` |
| `BLOB_BLOCK_SIZE_MIB` | No | `8` | Block size for parallel file transfers |
| `BLOB_MAX_CONCURRENCY` | No | `4` | Maximum parallel block requests per transfer |

Assign the identity the least-privileged data-plane role needed by the application, typically
**Storage Blob Data Contributor**, scoped to the target container when possible. Avoid
`BODY_AND_HEADERS` logging for sensitive payloads.

## Build and run

```powershell
mvn package
$env:AZURE_STORAGE_ACCOUNT_ENDPOINT = "https://account.blob.core.windows.net"
mvn exec:java
```

Uploads use the SDK's file-based parallel block transfer, so memory remains bounded for
multi-gigabyte files. Normal uploads use an ETag `If-Match` condition (or `If-None-Match: *` for a
new blob), which turns concurrent modifications into a precondition failure instead of silently
overwriting newer data. `uploadWithLease` is available when an exclusive write lease is preferred.

References:

- [Azure Blob Storage client library for Java](https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme)
- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
