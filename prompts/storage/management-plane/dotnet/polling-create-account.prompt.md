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

I'm creating a storage account with CreateOrUpdateAsync and it returns
an ArmOperation. How do I properly wait for it to finish? I need to:
1. Start the create operation using CreateOrUpdateAsync
2. Wait for completion using WaitForCompletionAsync
3. Handle timeout scenarios where the operation takes too long
4. Understand the difference between WaitUntil.Completed and WaitUntil.Started

Use Azure.ResourceManager.Storage with DefaultAzureCredential. Show required
NuGet packages and explain the ArmOperation<T> pattern.

## Evaluation Criteria

- `StorageAccountCollection.CreateOrUpdateAsync()` returning `ArmOperation<StorageAccountResource>`
- `ArmOperation<T>.WaitForCompletionAsync()` for simple completion
- `ArmOperation<T>.HasCompleted` and `UpdateStatusAsync()` for manual polling
- `ArmOperation<T>.Value` to get the result after completion
- Timeout handling with `CancellationToken`
- `WaitUntil.Completed` vs `WaitUntil.Started` parameter
- Error handling when the LRO fails

## Context

Storage Account creation is an LRO that typically takes 10-30 seconds. The
ArmOperation pattern is used across all Azure management SDKs, so understanding
it here teaches a transferable skill for all resource provisioning.
