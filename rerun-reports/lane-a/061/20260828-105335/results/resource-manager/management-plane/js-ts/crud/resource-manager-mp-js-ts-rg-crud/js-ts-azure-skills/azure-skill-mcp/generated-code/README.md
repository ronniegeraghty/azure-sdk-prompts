# Azure Resource Group Manager

This TypeScript program uses the Azure management plane SDK to create, list,
read, tag, and delete an Azure resource group.

## Required packages

Runtime dependencies:

```powershell
npm install @azure/identity @azure/arm-resources
```

TypeScript development dependencies:

```powershell
npm install --save-dev typescript tsx @types/node
```

## Authentication and configuration

`DefaultAzureCredential` tries supported credential sources in order. For local
development, authenticate with a supported developer credential or configure a
service principal through environment variables.

Set the required program values:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "<resource-group-name>"
```

The authenticated identity needs permission to manage resource groups in the
subscription.

## Build and run

```powershell
npm install
npm run build
npm start
```

Running the program creates the resource group in `eastus`, lists the
subscription's resource groups, gets and tags the created group, and then waits
for its deletion to finish.
