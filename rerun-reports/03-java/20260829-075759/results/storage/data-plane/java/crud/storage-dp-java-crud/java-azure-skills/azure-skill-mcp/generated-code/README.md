# Azure Blob Storage CRUD (Java)

This Maven application uses passwordless authentication to:

1. Create `my-container` if needed.
2. Upload `data.txt` as `uploads/data.txt`.
3. List blob names and sizes.
4. Download the blob as `data-downloaded.txt`.
5. Delete the blob and container.

## Prerequisites

- Java 17+
- Maven 3.9+
- A local identity recognized by `DefaultAzureCredential`
- The **Storage Blob Data Contributor** role on the target storage account
- A `data.txt` file in the working directory

Set the Blob service endpoint without embedding credentials:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
```

Compile and run:

```powershell
mvn compile
mvn exec:java
```

`DefaultAzureCredential` is convenient for local development. For an Azure-hosted
production workload, prefer a specific `ManagedIdentityCredential` to make the
authentication behavior deterministic.

References:

- https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java
- https://learn.microsoft.com/java/api/overview/azure/identity-readme
