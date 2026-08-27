---
id: storage-mp-dotnet-polling
properties:
  service: storage
  plane: management-plane
  language: dotnet
  category: polling
  difficulty: intermediate
  description: 'Can a developer use the LRO polling pattern to create a Storage Account and wait for completion using the
    .NET management SDK?

    '
  sdk_package: Azure.ResourceManager.Storage
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/resourcemanager.storage-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- polling
- lro
- long-running-operation
- management-plane
---

# Polling/LRO: Create Storage Account (.NET)

## Prompt

I'm creating a storage account programmatically and the operation takes
a while to complete. How do I properly wait for long-running operations
in the Azure management SDK for .NET? I need to:
1. Start the create operation
2. Wait for it to complete
3. Handle timeout scenarios where the operation takes too long
4. Understand the different completion modes available

Authenticate securely using identity-based credentials.

## Evaluation Criteria

- Creates a storage account using the management SDK
- Waits for the long-running operation to complete
- Supports both blocking and polling completion modes
- Handles timeouts with cancellation tokens
- Handles errors when the LRO fails

## Context

Storage Account creation is an LRO that typically takes 10-30 seconds. The
ArmOperation pattern is used across all Azure management SDKs, so understanding
it here teaches a transferable skill for all resource provisioning.
