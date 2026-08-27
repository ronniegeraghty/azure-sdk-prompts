---
id: storage-dp-dotnet-batch
properties:
  service: storage
  plane: data-plane
  language: dotnet
  category: batch
  difficulty: advanced
  description: 'Can a developer perform batch blob operations including bulk delete and bulk set-tier using the .NET SDK?

    '
  sdk_package: Azure.Storage.Blobs
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- batch
- bulk-operations
- delete
- set-tier
---

# Batch Operations: Azure Blob Storage (.NET)

## Prompt

I have a container with about 500 old blobs I need to delete in one shot
instead of deleting them one at a time. How do I perform bulk blob operations
in a single HTTP request in .NET? I also need to handle partial failures
where some operations succeed and others fail.

## Evaluation Criteria

- Performs bulk blob operations (delete or tier change) in batched requests
- Handles batch size limits
- Handles partial failures where individual operations fail within a batch
- Uses the appropriate batch SDK package

## Context

Batch operations are essential for cost optimization (changing storage tiers)
and cleanup (bulk deletion). Without batch support, developers resort to
sequential API calls that are slow and consume excessive transaction units.
