# Azure Blob Manager

A reusable Python utility with synchronous and asynchronous Azure Blob Storage
operations. Authentication uses `DefaultAzureCredential`; account keys and
connection strings are intentionally unsupported.

## Setup

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "<existing-container>"
python .\main.py
```

The signed-in identity or managed identity needs an appropriate data-plane role,
such as **Storage Blob Data Contributor**, on the target account or container.
The demo expects the container to exist; it does not provision Azure resources.

## Configuration

| Variable | Default | Purpose |
| --- | ---: | --- |
| `AZURE_STORAGE_ACCOUNT_URL` | required | HTTPS Blob service endpoint |
| `AZURE_STORAGE_CONTAINER` | required by demo | Existing container name |
| `AZURE_STORAGE_RETRY_TOTAL` | `5` | Maximum retry attempts |
| `AZURE_STORAGE_RETRY_BACKOFF_FACTOR` | `0.8` | Exponential backoff factor |
| `AZURE_STORAGE_RETRY_BACKOFF_MAX` | `30` | Maximum backoff in seconds |
| `AZURE_STORAGE_HTTP_LOGGING` | `false` | Enable Azure HTTP logging |
| `AZURE_STORAGE_LOG_LEVEL` | `WARNING` | Python/Azure log level |
| `AZURE_STORAGE_MAX_BLOCK_SIZE` | `8388608` | Streaming upload block size |
| `AZURE_STORAGE_MAX_SINGLE_PUT_SIZE` | `8388608` | Single-request upload threshold |
| `AZURE_STORAGE_MAX_CONCURRENCY` | `4` | Parallel transfer workers |
| `AZURE_STORAGE_CONNECTION_TIMEOUT` | `20` | Socket connection timeout |
| `AZURE_STORAGE_READ_TIMEOUT` | `120` | Socket read timeout |

Every service operation also accepts `timeout`, which is sent as Azure
Storage's per-operation server timeout. Uploads stream from disk in blocks, and
updates use an ETag condition so a concurrent change fails cleanly instead of
being overwritten.
