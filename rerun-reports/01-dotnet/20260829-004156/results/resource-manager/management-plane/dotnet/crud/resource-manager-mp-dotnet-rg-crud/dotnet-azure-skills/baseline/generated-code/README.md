# Azure Resource Group Manager

This console application uses the current Azure management-plane SDK to create,
list, retrieve, tag, and delete an Azure resource group.

## Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager
dotnet add package Azure.ResourceManager.Resources
```

The project pins the resolved package versions in
`AzureResourceGroupManager.csproj`.

## Authentication

`DefaultAzureCredential` tries supported credentials in its credential chain.
For example, configure these environment variables for a service principal:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
```

The identity must have permission to create, read, update, and delete resource
groups in the default subscription.

## Run

The optional argument is the resource group name. If omitted, the application
generates a timestamped name.

```powershell
dotnet run -- "rg-sdk-demo"
```

The resource group is deleted in a `finally` block, including when a later
operation fails. Authentication failures, Azure service errors, and
cancellation are reported separately with nonzero exit codes.
