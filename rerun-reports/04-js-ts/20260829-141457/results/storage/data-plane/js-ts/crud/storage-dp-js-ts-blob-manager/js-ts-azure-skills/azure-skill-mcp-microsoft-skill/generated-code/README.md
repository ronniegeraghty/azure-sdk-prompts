# Azure Blob Manager

A reusable TypeScript service for streaming uploads and downloads, metadata,
blob index tags, listing, deletion, and lease-protected blob updates.

## Authentication

The project uses `ManagedIdentityCredential` and the HTTPS storage account
endpoint in `AZURE_STORAGE_ACCOUNT_ENDPOINT`. It does not accept connection
strings or account keys. Assign the workload's managed identity the least
privileged data-plane role needed by the application, typically **Storage Blob
Data Contributor**, scoped to the target container when possible.

For a user-assigned managed identity, also set `AZURE_CLIENT_ID`.

## Run the demo

1. Copy the values from `.env.example` into the Azure workload's environment.
2. Install and build:

   ```powershell
   npm install
   npm run build
   ```

3. Run:

   ```powershell
   npm run demo
   ```

The demo creates the configured container if it does not exist, uploads a
sample, lists and downloads it, performs a lease-protected overwrite, and
deletes the blob. Upload memory remains bounded by the configured buffer size
and concurrency rather than the source file size.

## Configuration

| Variable | Default | Purpose |
|---|---:|---|
| `AZURE_STORAGE_ACCOUNT_ENDPOINT` | required | HTTPS Blob service endpoint |
| `AZURE_STORAGE_CONTAINER_NAME` | `blob-manager-demo` | Demo container |
| `AZURE_STORAGE_MAX_RETRIES` | `5` | Maximum SDK request attempts |
| `AZURE_STORAGE_RETRY_DELAY_MS` | `1000` | Initial exponential retry delay |
| `AZURE_STORAGE_MAX_RETRY_DELAY_MS` | `30000` | Exponential retry delay cap |
| `AZURE_STORAGE_LEASE_WAIT_MS` | `30000` | Maximum time to wait for another writer |
| `AZURE_STORAGE_LEASE_POLL_MS` | `1000` | Lease acquisition polling interval |
| `AZURE_STORAGE_UPLOAD_BUFFER_SIZE` | `8388608` | Upload stream block buffer bytes |
| `AZURE_STORAGE_UPLOAD_CONCURRENCY` | `5` | Concurrent staged block uploads |
| `AZURE_SDK_LOG_LEVEL` | `warning` | `off`, `error`, `warning`, `info`, or `verbose` |

## SDK references

- [Azure Blob Storage JavaScript client library](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme)
- [Upload a block blob with JavaScript](https://learn.microsoft.com/azure/storage/blobs/storage-blob-upload-javascript)
- [Manage blob leases with JavaScript](https://learn.microsoft.com/azure/storage/blobs/storage-blob-lease-javascript)
