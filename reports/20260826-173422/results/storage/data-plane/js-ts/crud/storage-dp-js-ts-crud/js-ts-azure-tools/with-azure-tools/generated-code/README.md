# Azure Blob Storage CRUD with TypeScript

This sample uses Microsoft Entra authentication through `DefaultAzureCredential`.
The identity must have a blob data-plane role such as **Storage Blob Data Contributor**
on the target storage account.

## Required packages

```bash
npm install @azure/storage-blob @azure/identity
npm install --save-dev typescript tsx @types/node
```

## Run

Use Node.js 18 or later, authenticate with one of the credential sources supported by
`DefaultAzureCredential`, and set the storage account name:

```powershell
$env:AZURE_STORAGE_ACCOUNT_NAME = "<storage-account-name>"
npm install
npm start
```

The program creates `my-container`, uploads and reads `greeting.txt`, then deletes
the blob and container. Do not run it against a container whose contents must be
retained.

References:

- [Azure Blob Storage client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme)
- [DefaultAzureCredential overview](https://aka.ms/azsdk/js/identity/credential-chains#defaultazurecredential-overview)
