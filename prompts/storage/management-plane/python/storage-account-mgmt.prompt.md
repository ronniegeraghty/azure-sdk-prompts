---
id: storage-mp-python-account-mgmt
properties:
  service: storage
  plane: management-plane
  language: python
  category: provisioning
  difficulty: intermediate
  description: 'Can a developer create, configure, and manage Azure Storage Accounts using the Python management SDK?

    '
  sdk_package: azure-mgmt-storage
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/mgmt-storage-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- storage-account
- management-plane
- provisioning
---

# Storage Account Management: Azure Storage (Python)

## Prompt

Write a Python script that manages
Azure Storage Accounts using the management plane SDK:
1. Authenticate using DefaultAzureCredential
2. Create a new Storage Account with Standard_LRS SKU in "eastus"
3. List all Storage Accounts in a resource group
4. Get the properties of the created Storage Account
5. Update the account to enable blob versioning
6. Delete the Storage Account

Show required pip packages and include proper error handling.

## Evaluation Criteria

- Includes the required Azure management and identity SDK packages
- Creates a management client authenticated with credential and subscription ID
- Creates a storage account as a long-running operation with the correct SKU and kind
- Lists all storage accounts in the resource group
- Retrieves detailed properties of a specific storage account
- Enables blob versioning on the account (via account update or blob service properties)
- Deletes the storage account
- Code builds and runs without import errors or API misuse

## Context

Storage Account management is one of the most common management plane tasks.
This tests whether the generated code covers the full lifecycle including
the long-running create operation and model configuration for SKUs and features.
