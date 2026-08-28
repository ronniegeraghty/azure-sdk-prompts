# Azure Resource Group Manager

A .NET console sample using the modern `Azure.ResourceManager` management-plane
SDK to create, list, inspect, tag, and delete an Azure resource group.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager
```

The project currently resolves:

- `Azure.Identity` 1.21.0
- `Azure.ResourceManager` 1.14.0

`Azure.ResourceManager` contains the resource-group APIs in the
`Azure.ResourceManager.Resources` namespace. No
`Microsoft.Azure.Management.*` packages are used.

## Run

`DefaultAzureCredential` checks supported credential sources such as environment
credentials, workload identity, managed identity, Azure CLI, and developer-tool
sign-ins. The authenticated identity needs permission to manage resource groups
at subscription scope, such as the built-in **Resource Group Contributor** role.

```powershell
dotnet run -- my-resource-group-name
```

If the name is omitted, the program creates a timestamped name. The sample
deletes the resource group after completing the demonstration. If an operation
fails after creation, the `finally` block attempts cleanup and reports any
cleanup failure.

## References

- [Azure Resource Manager client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager-readme)
- [DefaultAzureCredential](https://learn.microsoft.com/dotnet/api/azure.identity.defaultazurecredential)
- [ResourceGroupCollection.CreateOrUpdateAsync](https://learn.microsoft.com/dotnet/api/azure.resourcemanager.resources.resourcegroupcollection.createorupdateasync)
