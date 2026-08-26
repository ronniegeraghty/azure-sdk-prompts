# Azure Blob Storage Manager

Reusable synchronous and asynchronous Python services for streaming uploads and
downloads, metadata and index tags, optimistic ETag concurrency, blob leases,
listing, and deletion.

## Setup

1. Create a virtual environment and run `pip install -r requirements.txt`.
2. Copy `.env.example` values into your environment and set
   `AZURE_STORAGE_ACCOUNT_URL` to the HTTPS blob endpoint.
3. Grant the workload identity an appropriate data-plane role, such as
   **Storage Blob Data Contributor**, scoped as narrowly as possible.
4. Ensure `AZURE_STORAGE_CONTAINER` already exists, then run `python main.py`.

No connection string or account key is accepted. Authentication uses
`DefaultAzureCredential`; set `AZURE_TOKEN_CREDENTIALS=prod` in Azure-hosted
production environments to constrain its credential chain.

The operation `timeout` is sent to Azure Storage as the service-side timeout.
Connection/read retries are controlled by the retry environment variables.

## References

- [Upload performance tuning](https://learn.microsoft.com/azure/storage/blobs/storage-blobs-tune-upload-download-python)
- [Blob leases](https://learn.microsoft.com/azure/storage/blobs/storage-blob-lease-python)
- [Azure SDK Python usage patterns](https://learn.microsoft.com/azure/developer/python/sdk/azure-sdk-library-usage-patterns)
