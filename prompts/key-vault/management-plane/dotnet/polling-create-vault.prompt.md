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
using the management SDK, and wait for the LRO to complete. How do I do
this with Azure.ResourceManager.KeyVault?
1. Authenticate using DefaultAzureCredential
2. Create a Key Vault with RBAC authorization enabled in "eastus"
3. Configure soft-delete and purge protection
4. Wait for the vault creation to complete using the ArmOperation pattern

Show required NuGet packages and how to set RBAC roles during creation.

## Evaluation Criteria

- `KeyVaultCollection.CreateOrUpdateAsync()` returning `ArmOperation<KeyVaultResource>`
- `KeyVaultCreateOrUpdateContent` with `KeyVaultProperties`
- Configuring `EnableRbacAuthorization`, `EnableSoftDelete`, `EnablePurgeProtection`
- `VaultAccessPolicy` vs RBAC authorization model
- `ArmOperation<T>.WaitForCompletionAsync()` for completion
- `WaitUntil.Completed` vs `WaitUntil.Started`
- Tenant ID and object ID configuration
- Error handling for existing vaults and soft-deleted vaults

## Context

Key Vault creation requires configuring security-sensitive properties like
access policies and purge protection at creation time. The LRO pattern is
critical because vault DNS propagation adds latency to the operation.
