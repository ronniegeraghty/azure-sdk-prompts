# Azure Blob Manager

A reusable Python utility for synchronous and asynchronous Azure Blob Storage
operations. It uses passwordless authentication, streams large transfers, uses
ETags to prevent lost updates, and supports exclusive blob leases.

## Setup

Install the dependencies:

```powershell
python -m pip install -r requirements.txt
```

Set the endpoint of an existing storage account and container:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "<existing-container>"
python main.py
```

No account key or connection string is accepted. `DefaultAzureCredential`
automatically uses the workload's managed identity in Azure. Assign that
identity the least-privileged data-plane role it needs, such as **Storage Blob
Data Contributor**, at the container scope.

## Configuration

| Environment variable | Default | Purpose |
|---|---:|---|
| `AZURE_STORAGE_ACCOUNT_URL` | required | HTTPS Blob service endpoint |
| `AZURE_STORAGE_CONTAINER` | `blob-manager-demo` | Existing container |
| `AZURE_STORAGE_MAX_RETRIES` | `5` | Maximum exponential retry attempts |
| `AZURE_STORAGE_RETRY_DELAY` | `1.0` | Initial retry backoff in seconds |
| `AZURE_STORAGE_RETRY_MAX_DELAY` | `30.0` | Maximum retry delay in seconds |
| `AZURE_STORAGE_CONNECTION_TIMEOUT` | `20` | Connection timeout in seconds |
| `AZURE_STORAGE_READ_TIMEOUT` | `120` | Socket read timeout in seconds |
| `AZURE_STORAGE_MAX_CONCURRENCY` | `4` | Parallel transfer workers |
| `AZURE_STORAGE_BLOCK_SIZE` | `8388608` | Upload block size in bytes |
| `AZURE_STORAGE_LOG_LEVEL` | `WARNING` | Azure HTTP logging level |

The optional per-operation `timeout` is passed to Azure Storage as its
server-side request timeout. Connection and socket read limits are configured
separately on the clients.

## Concurrency behavior

Uploads first read the current ETag and conditionally replace the blob only if
that ETag is still current. New blobs use create-only semantics. A competing
writer therefore receives a clear conflict result instead of silently losing
its changes. Pass a lease returned by `acquire_lease` when an exclusive update
window is required.

Large uploads are read in blocks rather than loaded into memory. The async
implementation moves local file reads and writes off the event loop.

## References

- [Azure Blob Storage Python client library](https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme)
- [Manage concurrency in Blob Storage](https://learn.microsoft.com/azure/storage/blobs/concurrency-manage)
- [Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview)
