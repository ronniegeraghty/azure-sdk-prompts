# Azure Resource Group Manager

This console application uses the current Azure management-plane SDK to:

1. Authenticate with `DefaultAzureCredential`.
2. Create a resource group in `eastus`.
3. List the subscription's resource groups.
4. Read the created resource group's details.
5. Add a tag.
6. Delete the resource group.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager
dotnet add package Azure.ResourceManager.Resources
```

The project file pins tested package versions.

## Authentication

For local development, sign in using a credential supported by
`DefaultAzureCredential`, such as Visual Studio, Azure CLI, or environment
variables for a service principal:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
```

The identity needs permission to create, read, update, and delete resource
groups in the target subscription, such as the `Contributor` role.

## Run

```powershell
dotnet run -- "rg-sdk-demo"
```

If no name is supplied, the program generates a timestamped resource group
name. Running the program performs real create, update, and delete operations
in the default Azure subscription selected by `DefaultAzureCredential`.
