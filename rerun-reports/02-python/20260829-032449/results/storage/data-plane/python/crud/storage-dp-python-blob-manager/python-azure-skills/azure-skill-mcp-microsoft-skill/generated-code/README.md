# Azure Blob Manager

A reusable Python utility that provides synchronous and asynchronous Azure Blob
Storage uploads, downloads, listing, deletion, and lease acquisition.

The utility authenticates only with `DefaultAzureCredential`. It never accepts a
connection string or account key. Uploads stream from disk in blocks and use
ETag preconditions so a concurrent writer cannot silently overwrite a newer
blob version. Downloads also stream to a temporary file before atomically
replacing the destination.

## Setup

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
```

Set the endpoint and the name of an existing container:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "<container>"
$env:AZURE_TOKEN_CREDENTIALS = "prod"
python .\main.py
```

The identity needs an appropriate data-plane role, such as **Storage Blob Data
Contributor**, scoped as narrowly as possible. For local development,
`DefaultAzureCredential` can use a supported developer credential.

## Optional configuration

| Environment variable | Default | Purpose |
|---|---:|---|
| `AZURE_STORAGE_RETRY_TOTAL` | `5` | Maximum retry attempts |
| `AZURE_STORAGE_RETRY_INITIAL_DELAY` | `2` | Initial backoff in seconds |
| `AZURE_STORAGE_RETRY_INCREMENT` | `2` | Exponential retry increment |
| `AZURE_STORAGE_HTTP_LOG_LEVEL` | `WARNING` | Set to `DEBUG` for HTTP request/response logging |
| `AZURE_STORAGE_MAX_CONCURRENCY` | `4` | Parallel transfer workers |
| `AZURE_STORAGE_MAX_BLOCK_SIZE` | `8388608` | Upload block size in bytes |
| `AZURE_STORAGE_MAX_SINGLE_PUT_SIZE` | `67108864` | Single-request upload threshold in bytes |

HTTP logging can contain request details. Enable `DEBUG` only while diagnosing
an issue and protect the resulting logs.

## References

- [Azure Blob Storage client library for Python](https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme)
- [Authenticate Python apps to Azure services](https://learn.microsoft.com/azure/developer/python/sdk/authentication-overview)
- [Manage concurrency in Blob Storage](https://learn.microsoft.com/azure/storage/blobs/concurrency-manage)
