---
id: resource-manager-mp-dotnet-rg-crud
properties:
  service: resource-manager
  plane: management-plane
  language: dotnet
  category: crud
  difficulty: basic
  description: 'Can a developer create, list, update, and delete Azure Resource Groups using the .NET management SDK?

    '
  sdk_package: Azure.ResourceManager
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/resourcemanager-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- resource-groups
- management-plane
- provisioning
- getting-started
---

# Resource Group Management: Azure Resource Manager (.NET)

## Prompt

I'm setting up automation to create and manage resource groups using the
modern Azure management SDK in .NET. How do I:
1. Authenticate securely and get an ARM client
2. Create a new resource group in a specific region
3. List all resource groups in the subscription
4. Add tags to the resource group
5. Delete the resource group when done

Include proper error handling.

## Evaluation Criteria

The generated code should include:
- Uses the modern track 2 Azure Resource Manager SDK
- Authenticates with identity-based credentials
- Creates, lists, gets, and deletes resource groups
- Manages tags on resource groups
- Waits for deletion to complete

## Context

Resource group management is the foundation of Azure management plane operations.
This tests whether the generated code covers the modern Azure.ResourceManager SDK
(track 2) rather than the older Microsoft.Azure.Management packages.
