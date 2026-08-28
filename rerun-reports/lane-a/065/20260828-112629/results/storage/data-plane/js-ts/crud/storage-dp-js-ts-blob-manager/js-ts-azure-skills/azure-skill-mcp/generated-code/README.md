# Azure Blob Storage Manager

A reusable TypeScript utility for uploading, downloading, listing, and deleting
Azure block blobs. Uploads use bounded-memory streams and an exclusive blob
lease to prevent concurrent writers from overwriting the same blob.

## Requirements

- Node.js 22 or later
- An existing Azure Storage container
- A system-assigned or user-assigned managed identity with the least-privilege
  `Storage Blob Data Contributor` role on the target container

The project uses Microsoft Entra authentication only. It does not accept storage
account keys or connection strings.

## Configuration

| Environment variable | Required | Default | Description |
| --- | --- | --- | --- |
| `AZURE_STORAGE_BLOB_ENDPOINT` | Yes | - | Account endpoint, such as `https://account.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER` | Demo only | - | Existing container used by the demo |
| `AZURE_CLIENT_ID` | No | System identity | Client ID of a user-assigned managed identity |
| `AZURE_STORAGE_MAX_RETRIES` | No | `4` | Maximum retries after the initial request |
| `AZURE_STORAGE_RETRY_DELAY_MS` | No | `800` | Initial exponential retry delay |
| `AZURE_STORAGE_MAX_RETRY_DELAY_MS` | No | `30000` | Maximum exponential retry delay |
| `AZURE_SDK_LOG_LEVEL` | No | `off` | `verbose`, `info`, `warning`, `error`, or `off` |
| `AZURE_STORAGE_DEMO_BLOB` | No | `blob-manager-demo.txt` | Blob name used by the demo |

## Build and run

```powershell
npm install
npm run build
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<account>.blob.core.windows.net"
$env:AZURE_STORAGE_CONTAINER = "<container>"
npm run demo
```

The demo creates only temporary local files. It uploads a sample, lists the
container, downloads and prints the sample, performs a lease-protected
overwrite, and deletes the blob.

## Library usage

```typescript
import {
  BlobStorageManager,
  createBlobServiceClient,
} from "azure-blob-storage-manager";

const manager = new BlobStorageManager(
  createBlobServiceClient(),
  "my-container",
);

await manager.upload("./large-file.bin", "large-file.bin", {
  metadata: { source: "batch-import" },
  tags: { project: "archive", status: "ready" },
});
```

By default, uploads use five concurrent 8 MiB buffers. Override `bufferSize` and
`maxConcurrency` per upload to tune throughput and memory use.
