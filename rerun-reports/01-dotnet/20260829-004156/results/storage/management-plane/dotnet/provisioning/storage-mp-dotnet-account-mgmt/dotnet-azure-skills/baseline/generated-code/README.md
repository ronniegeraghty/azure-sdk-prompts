# Azure Storage account management sample

Required NuGet packages:

```powershell
dotnet add package Azure.Identity --version 1.17.2
dotnet add package Azure.ResourceManager.Storage --version 1.6.0
```

The program is a dry run unless `--execute` is supplied. To execute the
management-plane workflow, authenticate with a method supported by
`DefaultAzureCredential`, then set:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP = "<existing-resource-group>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<globally-unique-name>" # optional
dotnet run -- --execute
```

The identity must be authorized to create, read, update, and delete storage
accounts in the resource group. If `AZURE_STORAGE_ACCOUNT_NAME` is omitted, the
program generates a valid lowercase name. The sample creates a StorageV2
account with Standard_LRS in eastus, lists the group's accounts, reads the
created account, enables blob versioning through its default blob service, and
deletes the account. A `finally` block attempts deletion if a later operation
fails after creation.
