# Azure Resource Group management with TypeScript

This sample uses the Azure management plane SDK to create, list, retrieve,
tag, and delete a resource group. The resource group is deleted in a `finally`
block so cleanup is attempted even if a later operation fails.

## Install

```powershell
npm install
```

Required runtime packages:

```powershell
npm install @azure/identity @azure/arm-resources
```

The sample also uses TypeScript, `tsx`, and Node.js type declarations for local
development:

```powershell
npm install --save-dev typescript tsx @types/node
```

## Configure and run

Set the subscription and the name of the temporary resource group. Do not put
credentials in source code.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "<unique-resource-group-name>"
npm start
```

`DefaultAzureCredential` can use a supported local developer login or
environment-based service principal credentials. In Azure-hosted production
workloads, use a managed identity and grant only the management-plane RBAC
permissions needed to manage the target resource group.

> **Warning:** Running the sample creates and then deletes a real Azure resource
> group. Deleting a resource group also deletes every resource it contains. Use
> a new, disposable resource group name.
