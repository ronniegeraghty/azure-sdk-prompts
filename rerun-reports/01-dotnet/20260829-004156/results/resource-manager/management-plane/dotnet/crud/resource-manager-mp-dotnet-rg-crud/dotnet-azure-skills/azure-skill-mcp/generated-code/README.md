# Azure Resource Group Manager

This .NET console application uses the current Azure management-plane SDK to:

1. Authenticate with `DefaultAzureCredential`.
2. Create a uniquely named resource group in `eastus`.
3. List every resource group in the default subscription.
4. Retrieve and display the new resource group's details.
5. Add a `ManagedBy=Azure.ResourceManager` tag.
6. Delete the resource group.

If a later operation fails after creation, the `finally` block attempts to delete the
resource group so the sample does not leave it behind.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager
```

The project currently pins:

```xml
<PackageReference Include="Azure.Identity" Version="1.21.0" />
<PackageReference Include="Azure.ResourceManager" Version="1.14.0" />
```

Resource-group types are included in `Azure.ResourceManager`; the older
`Microsoft.Azure.Management.*` packages are not used.

## Run

Authenticate using any credential supported by `DefaultAzureCredential`, such as
Azure CLI for local development or managed identity when hosted in Azure. The identity
needs permission to create, read, tag, and delete resource groups in the subscription;
the built-in `Resource Group Contributor` role is sufficient when scoped appropriately.

```powershell
dotnet run
```

The program uses the default subscription visible to the selected credential and creates
then deletes a real resource group. Do not interrupt it after creation; if cleanup fails,
delete the printed resource group manually.

## References

- [Resource management using the Azure SDK for .NET](https://learn.microsoft.com/dotnet/azure/sdk/resource-management)
- [Azure Resource Manager client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/resource-manager)
- [DefaultAzureCredential class](https://learn.microsoft.com/dotnet/api/azure.identity.defaultazurecredential)
