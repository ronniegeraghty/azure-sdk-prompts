# Azure Storage account management with TypeScript

This example uses the Azure management-plane SDK to create, list, inspect,
update, and delete a StorageV2 account. The resource group must already exist.
The account name must be globally unique, 3-24 characters long, and contain
only lowercase letters and numbers.

## Install

```powershell
npm install
```

The required runtime packages are:

```powershell
npm install @azure/arm-storage @azure/identity
```

## Configure and run

Set the variables shown in `.env.example` in your shell, authenticate with a
developer credential supported by `DefaultAzureCredential`, and run:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-name>"
npm start
```

The program deletes the account in a `finally` block after the other operations
finish. `DefaultAzureCredential` is convenient for local development; prefer a
specific managed identity credential for Azure-hosted production applications.

Reference:
[Azure Storage management client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/arm-storage-readme)
