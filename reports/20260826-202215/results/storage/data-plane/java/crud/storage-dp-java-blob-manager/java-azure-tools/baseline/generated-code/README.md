# Azure Blob Manager

A small Java 17 library and demo for synchronous and asynchronous Azure Blob Storage operations.
Authentication uses `DefaultAzureCredential`; no account keys or connection strings are accepted.
In Azure, assign the workload's managed identity an appropriate Blob Storage data-plane role.

## Configuration

| Environment variable | Required | Default |
|---|---:|---:|
| `AZURE_STORAGE_ENDPOINT` | yes | - |
| `AZURE_STORAGE_CONTAINER` | yes | - |
| `AZURE_STORAGE_MAX_RETRIES` | no | `5` |
| `AZURE_STORAGE_RETRY_DELAY_SECONDS` | no | `2` |
| `AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS` | no | `30` |
| `AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS` | no | `120` |
| `AZURE_STORAGE_LOG_LEVEL` | no | `BASIC` |

Valid log levels are `NONE`, `BASIC`, `HEADERS`, `BODY`, and `BODY_AND_HEADERS`.
Request and response bodies can contain sensitive application data, so enable body logging only
temporarily.

## Run

```powershell
$env:AZURE_STORAGE_ENDPOINT = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "<existing-container>"
mvn compile exec:java
```

Uploads use staged blocks with bounded concurrency, so file contents are streamed from disk rather
than loaded into memory. Existing blobs are updated with an ETag condition, making competing writes
fail with HTTP 412 instead of silently overwriting one another. The lease-aware overload is provided
for callers that need an exclusive update window.
