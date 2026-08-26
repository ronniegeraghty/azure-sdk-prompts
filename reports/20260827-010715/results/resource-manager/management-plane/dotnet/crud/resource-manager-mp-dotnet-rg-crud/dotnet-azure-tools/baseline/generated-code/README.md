# Azure Resource Group management sample

This .NET 8 console application uses the current management-plane SDK:

- `Azure.Identity` for `DefaultAzureCredential`
- `Azure.ResourceManager.Resources` for subscription and resource-group operations

The sample creates a resource group in `eastus`, lists the subscription's resource
groups, reads the new group, adds a tag, and deletes the group. If an error occurs
after creation, it attempts to delete the group in `finally`.

## Configure and run

Authenticate with any identity supported by `DefaultAzureCredential`, such as
Azure CLI login, Visual Studio, workload identity, managed identity, or these
service-principal environment variables:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
```

Optionally set a resource-group name. Otherwise, the program generates a unique
name:

```powershell
$env:AZURE_RESOURCE_GROUP_NAME = "rg-sdk-sample"
```

Restore and run:

```powershell
dotnet restore
dotnet run
```

The authenticated identity needs permission to list, create, update, and delete
resource groups in the selected subscription, such as the `Contributor` role.

The package references are:

```xml
<PackageReference Include="Azure.Identity" Version="1.17.0" />
<PackageReference Include="Azure.ResourceManager.Resources" Version="1.10.0" />
```
