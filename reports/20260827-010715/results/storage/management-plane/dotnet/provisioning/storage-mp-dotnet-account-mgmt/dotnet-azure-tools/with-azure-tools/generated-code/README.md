# Azure Storage Account management sample

This .NET console program uses the Azure management plane to:

1. Authenticate with `DefaultAzureCredential`.
2. Create a `StorageV2` account in `eastus` with the `Standard_LRS` SKU.
3. List the storage accounts in the target resource group.
4. Read and display the new account's properties.
5. Enable blob versioning through its blob-service resource.
6. Delete the account in a `finally` block.

The program refuses to continue if the requested account name already exists, so cleanup cannot
delete a pre-existing account.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager.Storage
```

`Azure.ResourceManager.Storage` brings in the core Azure Resource Manager dependencies
transitively.

## Configuration

The signed-in identity needs permission to read the resource group and manage storage accounts,
such as the built-in **Storage Account Contributor** role scoped to that resource group. Configure
the target with environment variables; do not store credentials in source code.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-guid>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-lowercase-name>"
```

`DefaultAzureCredential` can use local developer credentials or environment-based service
principal credentials. For example, a service principal can be configured with
`AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`. In Azure-hosted production code,
prefer a managed identity.

Run the sample only when you intend to create and immediately delete the named resource:

```powershell
dotnet run
```

Pressing Ctrl+C requests cancellation. Azure authentication, HTTP failures, invalid configuration,
and cancellation return distinct nonzero exit codes. If any step after account creation fails, the
program still attempts deletion; a deletion failure is surfaced rather than ignored.

## References

- [Azure Storage management library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.storage-readme)
- [Manage storage accounts with .NET](https://learn.microsoft.com/azure/storage/common/storage-srp-manage-account-dotnet)
- [`DefaultAzureCredential` overview](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains)
- [`BlobServiceResource.CreateOrUpdateAsync`](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.storage.blobserviceresource.createorupdateasync)
