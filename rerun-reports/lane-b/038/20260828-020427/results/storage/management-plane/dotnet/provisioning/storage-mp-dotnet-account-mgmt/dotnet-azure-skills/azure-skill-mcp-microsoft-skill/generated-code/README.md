# Azure Storage Account management sample

This .NET 8 console application uses the Azure management-plane SDK to:

1. Authenticate with `DefaultAzureCredential`.
2. Create a general-purpose v2 storage account with `Standard_LRS` redundancy in `eastus`.
3. List the storage accounts in the target resource group.
4. Read and display the created account's properties.
5. Enable blob versioning.
6. Delete the created account.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity --version 1.21.0
dotnet add package Azure.ResourceManager.Storage --version 1.7.0
```

The project file already contains these package references. Resource Manager base and resource-group
types are included transitively by `Azure.ResourceManager.Storage`.

## Configuration

Set the following environment variables. The resource group must already exist, and the storage
account name must be globally unique, 3-24 characters long, and contain only lowercase letters and
digits.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-account-name>"
```

For local development, sign in through a developer credential supported by
`DefaultAzureCredential`, then run:

```powershell
dotnet run
```

Running the application creates and then deletes a real Azure Storage account and may incur a small
charge. The authenticated identity needs permission to read the resource group and create, read,
update, and delete storage accounts. The program refuses to continue if an account with the same
name already exists in the resource group, preventing the cleanup step from deleting that account.

`DefaultAzureCredential` is convenient for local development. For an Azure-hosted production
application, prefer a deterministic `ManagedIdentityCredential` and grant it least-privilege RBAC.
