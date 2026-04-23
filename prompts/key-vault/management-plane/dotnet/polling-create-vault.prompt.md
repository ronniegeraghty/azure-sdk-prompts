---
id: key-vault-mp-dotnet-polling
properties:
  service: key-vault
  plane: management-plane
  language: dotnet
  category: polling
  difficulty: intermediate
  description: 'Can a developer use the LRO polling pattern to create a Key Vault and wait for completion using the .NET management
    SDK?

    '
  sdk_package: Azure.ResourceManager.KeyVault
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/resourcemanager.keyvault-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- polling
- lro
- long-running-operation
- management-plane
---

# Polling/LRO: Create Key Vault (.NET)

## Prompt

I need to create a Key Vault with RBAC authorization and purge protection
using the management SDK, and wait for the operation to complete. How do I
do this in .NET?
1. Authenticate securely using identity-based credentials
2. Create a Key Vault with RBAC authorization enabled
3. Configure soft-delete and purge protection
4. Wait for the vault creation to complete

Explain the difference between RBAC and access-policy authorization models.

## Evaluation Criteria

- Creates a Key Vault using the management SDK with RBAC authorization
- Configures soft-delete and purge protection at creation time
- Waits for the long-running operation to complete
- Explains RBAC vs access-policy authorization
- Handles errors for existing vaults and soft-deleted vaults

## Context

Key Vault creation requires configuring security-sensitive properties like
access policies and purge protection at creation time. The LRO pattern is
critical because vault DNS propagation adds latency to the operation.
