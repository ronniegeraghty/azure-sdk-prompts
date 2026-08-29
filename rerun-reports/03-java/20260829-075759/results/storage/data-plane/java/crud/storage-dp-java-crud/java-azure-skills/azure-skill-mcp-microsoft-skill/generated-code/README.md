# Azure Blob Storage CRUD sample

This Java 17 application authenticates with `DefaultAzureCredential`, creates
`my-container` when needed, uploads `data.txt` as `uploads/data.txt`, lists
blobs and sizes, downloads the blob to `data-downloaded.txt`, and deletes the
blob and container.

## Prerequisites

- Java 17 and Maven
- An existing Azure Storage account
- An identity with the **Storage Blob Data Contributor** role on the account
- Local authentication supported by `DefaultAzureCredential`, such as Azure
  CLI sign-in, or workload/managed identity when hosted in Azure

Set the storage endpoint without embedding credentials:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
```

Build and run:

```powershell
mvn compile
mvn exec:java
```

The application intentionally performs destructive cleanup. Do not run it
against a container whose other contents must be preserved, because deleting a
non-empty container fails and the sample reports the resulting
`BlobStorageException`.

## References

- [Azure Blob Storage client library for Java](https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme)
- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
