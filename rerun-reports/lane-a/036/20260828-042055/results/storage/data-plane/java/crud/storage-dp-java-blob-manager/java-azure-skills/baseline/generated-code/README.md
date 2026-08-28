# Azure Blob Manager

A Java 17 utility with synchronous and asynchronous Azure Blob Storage operations.
Authentication uses `DefaultAzureCredential`, so no account keys or connection strings are needed.

## Configuration

| Environment variable | Required | Default |
|---|---:|---:|
| `AZURE_STORAGE_BLOB_ENDPOINT` | Yes | - |
| `AZURE_STORAGE_CONTAINER` | No | `blob-manager-demo` |
| `AZURE_STORAGE_MAX_RETRIES` | No | `5` |
| `AZURE_STORAGE_RETRY_DELAY_MS` | No | `800` |
| `AZURE_STORAGE_MAX_RETRY_DELAY_MS` | No | `10000` |
| `AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS` | No | `120` |
| `AZURE_STORAGE_HTTP_LOG_LEVEL` | No | `BASIC` |

The endpoint has the form `https://<account>.blob.core.windows.net`. In Azure, assign the
managed identity an appropriate data-plane role such as **Storage Blob Data Contributor**.
`DefaultAzureCredential` also supports developer credentials for local testing.

The demo expects the configured container to exist. Run it with:

```powershell
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "<existing-container>"
mvn compile exec:java
```

Uploads use staged block transfer from a file with bounded concurrency, avoiding whole-file
buffering. Every write requires an explicit concurrency condition: create-only, an expected ETag,
or an active lease ID.
