# Azure Blob Manager

A reusable TypeScript service for streaming Azure Blob Storage file operations.
It authenticates with managed identity, never with account keys or connection
strings.

## Configuration

| Environment variable | Required | Default |
| --- | --- | --- |
| `AZURE_STORAGE_ACCOUNT_ENDPOINT` | Yes | Example: `https://myaccount.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER_NAME` | No | `blob-manager-demo` |
| `AZURE_CLIENT_ID` | No | System-assigned managed identity; set for a user-assigned identity |
| `AZURE_STORAGE_MAX_RETRIES` | No | `5` |
| `AZURE_STORAGE_RETRY_DELAY_MS` | No | `800` |
| `AZURE_STORAGE_MAX_RETRY_DELAY_MS` | No | `30000` |
| `AZURE_SDK_LOG_LEVEL` | No | `warning` |

The managed identity needs the **Storage Blob Data Contributor** role scoped as
narrowly as practical. The role assignment and container must normally be
provisioned ahead of time; the demo also calls `createIfNotExists` for
convenience.

## Run

```powershell
npm install
npm run build
$env:AZURE_STORAGE_ACCOUNT_ENDPOINT = "https://<account>.blob.core.windows.net"
npm start
```

`uploadFile` uses `uploadStream`, bounded buffering, and configurable
concurrency. Existing blobs are protected by a renewable 60-second lease while
being overwritten. A new blob uses an atomic create-only condition so another
writer cannot win the existence-check race and then be overwritten.

References:

- [Azure Blob Storage JavaScript client library](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme)
- [Authenticate JavaScript apps to Azure with managed identity](https://learn.microsoft.com/azure/developer/javascript/sdk/authentication/system-assigned-managed-identity)
- [Manage blob leases with JavaScript](https://learn.microsoft.com/azure/storage/blobs/storage-blob-lease-javascript)
