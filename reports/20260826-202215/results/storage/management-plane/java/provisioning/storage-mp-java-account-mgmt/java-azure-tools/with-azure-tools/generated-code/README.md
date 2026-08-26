# Azure Storage Account management sample

This Java 17 sample uses `DefaultAzureCredential` and the Azure management-plane
SDK to create, list, inspect, update, and delete an Azure Storage account.

Set these environment variables before running:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-storage-account-name>"
```

For local development, `DefaultAzureCredential` can use a supported developer
credential. In Azure, prefer a managed identity with only the required
management-plane permissions. Never store credentials in source code.

Run the sample only when you intend to create and delete a real resource:

```powershell
mvn compile exec:java
```

The storage account name must be globally unique, contain 3-24 lowercase
letters and numbers, and the resource group must already exist.

References:

- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
- [Azure Resource Manager libraries for Java](https://learn.microsoft.com/azure/developer/java/sdk/management-sdk/overview)
- [Blob versioning](https://learn.microsoft.com/azure/storage/blobs/versioning-enable)
