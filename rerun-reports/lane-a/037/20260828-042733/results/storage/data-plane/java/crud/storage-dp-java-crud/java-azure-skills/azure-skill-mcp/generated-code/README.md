# Azure Blob Storage CRUD (Java)

This Maven application creates a container, uploads and lists a blob, downloads
it, and then deletes both the blob and container.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An Azure identity with Blob Storage data permissions
- A storage account Blob service endpoint

Set the endpoint in PowerShell:

```powershell
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<account>.blob.core.windows.net"
```

`DefaultAzureCredential` uses credentials from supported local developer tools
or environment-based service principal settings. No account key or connection
string is stored in the application.

## Build and run

```powershell
mvn compile
mvn exec:java
```

The application reads `data.txt`, writes `data-downloaded.txt`, and removes the
Azure resources it creates. If `my-container` already exists, the final step
still deletes it; use a dedicated development storage account.
