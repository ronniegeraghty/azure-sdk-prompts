# Azure Resource Group Management with TypeScript

This example uses the Azure management-plane SDK to create, list, read, tag,
and delete a resource group. The `finally` block deletes the resource group
even when a later operation fails.

## Packages

Runtime packages:

```powershell
npm install @azure/identity @azure/arm-resources
```

TypeScript development packages:

```powershell
npm install --save-dev typescript tsx @types/node
```

## Configuration and execution

Set the subscription ID and, optionally, a resource group name:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "sdk-rg-example"
npm install
npm run build
npm start
```

`DefaultAzureCredential` uses the available credential in its standard
credential chain. In Azure-hosted production environments, use a managed
identity with only the resource-group permissions the application requires.

> Running the program performs live create, update, and delete operations in
> the configured Azure subscription.

## References

- [Azure SDK for JavaScript: create a client and call methods](https://learn.microsoft.com/azure/developer/javascript/sdk/use-azure-sdk#create-an-sdk-client-and-call-methods)
- [`@azure/arm-resources` overview](https://learn.microsoft.com/javascript/api/overview/azure/arm-resources-readme?view=azure-node-latest)
