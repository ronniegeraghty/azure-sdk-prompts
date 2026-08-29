# Azure Blob Storage CRUD with TypeScript

This sample creates a container, uploads and lists a block blob, downloads its
content, and deletes both the blob and container.

## Install

```powershell
npm install
```

Runtime packages:

- `@azure/storage-blob`
- `@azure/identity`

Development packages:

- `typescript`
- `@types/node`

## Configure and run

The authenticated identity needs the **Storage Blob Data Contributor** role (or
equivalent data-plane permissions) on the storage account.

```powershell
$env:AZURE_STORAGE_ACCOUNT_NAME = "<storage-account-name>"
$env:AZURE_TOKEN_CREDENTIALS = "dev"
npm run build
npm start
```

`DefaultAzureCredential` can use a supported local developer credential. In an
Azure-hosted production environment, use managed identity and set
`AZURE_TOKEN_CREDENTIALS` to `prod`.

Reference: [Azure Blob Storage client library quickstart for Node.js with
TypeScript](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-nodejs-typescript)
