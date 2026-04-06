---
id: storage-mp-dotnet-account-mgmt
properties:
  service: storage
  plane: management-plane
  language: dotnet
  category: provisioning
  difficulty: intermediate
  description: 'Can a developer create, configure, and manage Azure Storage Accounts using the .NET management SDK?

    '
  sdk_package: Azure.ResourceManager.Storage
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/resourcemanager.storage-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- storage-account
- management-plane
- provisioning
---

# Storage Account Management: Azure Storage (.NET)

## Prompt

I need to create a Standard_LRS storage account in eastus using the
Azure.ResourceManager.Storage SDK, then enable blob versioning on it
afterward. How do I do this with the new Track 2 management SDK?

Show me authentication with DefaultAzureCredential, the create call,
and how to update properties. Include required NuGet packages.

## Evaluation Criteria

The generated code should include:
- `Azure.ResourceManager.Storage` NuGet package
- `ArmClient` and subscription/resource group navigation
- `StorageAccountCollection.CreateOrUpdate()` with `StorageAccountCreateOrUpdateContent`
- SKU and kind configuration (`StorageSku`, `StorageKind`)
- Listing and getting storage accounts
- Updating properties via `StorageAccountPatch`
- Delete operation

## Context

Storage Account management is one of the most common management plane tasks.
This tests whether the generated code covers the full lifecycle of a Storage Account
including the more complex configuration options like SKU, kind, and feature toggles.
