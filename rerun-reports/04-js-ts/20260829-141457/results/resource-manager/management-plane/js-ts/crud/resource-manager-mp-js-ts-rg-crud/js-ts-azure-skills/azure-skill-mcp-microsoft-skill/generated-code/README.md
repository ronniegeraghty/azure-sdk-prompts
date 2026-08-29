# Azure Resource Group management with TypeScript

This sample uses the Azure management plane SDK to create, list, retrieve,
update, and delete a resource group. The resource group is deleted in a
`finally` block so cleanup is attempted if a later operation fails.

## Required packages

```powershell
npm install @azure/identity @azure/arm-resources
npm install --save-dev typescript @types/node
```

## Run

Set the subscription ID and optionally choose the resource group name:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "typescript-sdk-resource-group"
npm install
npm start
```

`DefaultAzureCredential` automatically tries supported credential sources,
including environment credentials, workload identity, managed identity, and
developer credentials. Do not store credentials in source code.

The signed-in identity needs permission to manage resource groups at the
subscription scope, such as the built-in **Resource Group Contributor** role.
Running the sample creates and then deletes a real Azure resource group.

## References

- [Azure Resource Manager SDK for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/resources)
- [`@azure/arm-resources` API reference](https://learn.microsoft.com/javascript/api/@azure/arm-resources/)
- [`DefaultAzureCredential` API reference](https://learn.microsoft.com/javascript/api/@azure/identity/defaultazurecredential)
