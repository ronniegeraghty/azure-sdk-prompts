# Azure Storage account management with TypeScript

This example uses the Azure management plane SDK to create, inspect, update,
and delete a StorageV2 account. It defines the operations but does not deploy
anything until you run it with Azure credentials and an existing resource
group.

## Install and build

```powershell
npm install
npm run build
```

The required runtime packages are:

```powershell
npm install @azure/arm-storage @azure/identity
```

## Configuration

Set the following environment variables:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-lowercase-name>"
```

The account name must be 3-24 characters and contain only lowercase letters
and numbers. The signed-in identity needs permission to manage Storage
Accounts and Blob Service properties in the target resource group.

Run the compiled example only when you intend to create and then delete the
account:

```powershell
npm start
```

`DefaultAzureCredential` is convenient for local development. For an
Azure-hosted production workload, prefer a specific managed identity
credential so authentication is deterministic.

## References

- [Azure Storage management SDK for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/storage-management-readme)
- [BlobServicePropertiesProperties.isVersioningEnabled](https://learn.microsoft.com/javascript/api/@azure/arm-storage/blobservicepropertiesproperties#@azure-arm-storage-blobservicepropertiesproperties-isversioningenabled)
- [DefaultAzureCredential](https://learn.microsoft.com/javascript/api/@azure/identity/defaultazurecredential)
