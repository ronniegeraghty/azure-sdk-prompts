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

I need to create a storage account programmatically and then modify
its configuration using the Azure management SDK in .NET. How do I:
1. Authenticate and navigate the resource hierarchy
2. Create a storage account with specific SKU and region
3. Enable features like blob versioning after creation
4. List and manage storage accounts

## Evaluation Criteria

The generated code should include:
- Uses the modern track 2 Azure Resource Manager Storage SDK
- Authenticates with identity-based credentials
- Creates a storage account with SKU and kind configuration
- Lists and retrieves storage accounts
- Updates properties after creation
- Handles deletion

## Context

Storage Account management is one of the most common management plane tasks.
This tests whether the generated code covers the full lifecycle of a Storage Account
including the more complex configuration options like SKU, kind, and feature toggles.
