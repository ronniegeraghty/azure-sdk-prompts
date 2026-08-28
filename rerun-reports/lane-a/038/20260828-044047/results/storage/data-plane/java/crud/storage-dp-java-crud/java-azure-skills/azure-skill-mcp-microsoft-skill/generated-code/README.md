# Azure Blob Storage CRUD with Java

This sample creates a container, uploads and lists a blob, downloads it, and
then deletes the blob and container. It uses `DefaultAzureCredential`; no
credentials are stored in source code.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An existing Azure Storage account
- A local `data.txt` file in the project directory
- An authenticated identity with permission to manage blobs and containers,
  such as the **Storage Blob Data Contributor** role

Set the storage account Blob service URL:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
```

For local development, `DefaultAzureCredential` can use environment
credentials or an existing developer-tool login. For service principal
authentication, set `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and
`AZURE_CLIENT_SECRET`.

## Build and run

```powershell
mvn compile
mvn exec:java
```

The application overwrites `uploads/data.txt` and removes any existing local
`data-downloaded.txt`. It then deletes `my-container`, so use a disposable
container and do not store unrelated blobs in it.
