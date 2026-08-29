# Azure Resource Group management with TypeScript

This example uses the Azure management plane SDK to create, list, read, update,
and delete a resource group.

## Install and build

```powershell
npm install
npm run build
```

Required runtime packages:

```powershell
npm install @azure/identity @azure/arm-resources
```

Authenticate with any method supported by `DefaultAzureCredential`, then set
the subscription ID. For example, in PowerShell:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "sdk-resource-group-example" # Optional
npm start
```

Running the program creates and then permanently deletes the named resource
group. Ensure the selected identity has permission to manage resource groups in
the subscription.
