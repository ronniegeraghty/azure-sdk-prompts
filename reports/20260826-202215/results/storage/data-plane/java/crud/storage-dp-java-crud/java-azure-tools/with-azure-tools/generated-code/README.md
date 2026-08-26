# Azure Blob Storage CRUD (Java)

This Maven application uses `DefaultAzureCredential` to create a container,
upload and list a blob, download it, and then delete the blob and container.
It does not store credentials in source code.

## Prerequisites

- Java 17 or newer
- Maven 3.9 or newer
- An existing Azure Storage account
- An identity with the `Storage Blob Data Contributor` role scoped as narrowly
  as practical

Set the Blob service endpoint before running:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
```

For local development, `DefaultAzureCredential` can use supported developer
credentials. In an Azure-hosted environment, configure a managed identity.

## Build and run

```powershell
mvn package
mvn exec:java -Dexec.mainClass="com.example.BlobStorageCrudApp"
```

The application expects `data.txt` in its working directory. It overwrites the
remote `uploads/data.txt` blob, replaces an existing local
`data-downloaded.txt`, and deletes `my-container` at the end. Use a dedicated
container because deleting it also requires that no other blobs remain.

## References

- [Azure Blob Storage Java quickstart](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java)
- [Authenticate Java apps to Azure services](https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview)
