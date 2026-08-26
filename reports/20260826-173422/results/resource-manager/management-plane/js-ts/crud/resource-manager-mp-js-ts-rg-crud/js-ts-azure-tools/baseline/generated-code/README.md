# Azure Resource Group management with TypeScript

This example uses `DefaultAzureCredential` and the Azure Resource Manager
management-plane SDK to create, list, read, update, and delete a resource
group.

## Required packages

Runtime packages:

```powershell
npm install @azure/identity @azure/arm-resources
```

TypeScript tooling:

```powershell
npm install --save-dev typescript tsx @types/node
```

## Run

Install dependencies and set the subscription ID:

```powershell
npm install
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
npm start
```

The example is an offline dry run by default. To perform the operations after
authenticating through a method supported by `DefaultAzureCredential`, set:

```powershell
$env:AZURE_EXECUTE = "true"
$env:AZURE_RESOURCE_GROUP_NAME = "typescript-sdk-example-rg"
npm start
```

Use a unique resource group name. The program deletes the resource group in a
`finally` block with `beginDeleteAndWait`, including when listing, reading, or
updating fails after creation.
