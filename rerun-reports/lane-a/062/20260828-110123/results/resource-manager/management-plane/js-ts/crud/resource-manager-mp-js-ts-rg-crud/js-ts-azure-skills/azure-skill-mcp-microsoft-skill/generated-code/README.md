# Azure Resource Group Manager (TypeScript)

This sample uses the Azure management-plane SDK to create, list, read, tag, and
delete a resource group. The `finally` block ensures the created demo resource
group is deleted if a later operation fails.

## Required packages

```powershell
npm install @azure/identity @azure/arm-resources
npm install --save-dev typescript @types/node
```

## Run

Set the subscription ID and authenticate with any method supported by
`DefaultAzureCredential`, such as environment-based service principal
credentials:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"

npm install
npm run build
npm start
```

Optionally set `AZURE_RESOURCE_GROUP_NAME` to override the default demo name.
The identity must have permission to create, list, read, update, and delete
resource groups in the subscription.
