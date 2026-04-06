---
id: storage-dp-dotnet-batch
service: storage
plane: data-plane
language: dotnet
category: batch
difficulty: advanced
description: >
  Can a developer perform batch blob operations including bulk delete
  and bulk set-tier using the .NET SDK?
sdk_package: Azure.Storage.Blobs
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
tags:
  - batch
  - bulk-operations
  - delete
  - set-tier
created: 2025-07-27
author: ronniegeraghty
---

# Batch Operations: Azure Blob Storage (.NET)

## Prompt

I have a container with about 500 old blobs I need to delete in one shot
instead of calling DeleteBlobAsync in a loop. How do I use BlobBatchClient
to bulk-delete them in a single HTTP request? I also need to handle partial
failures where some deletes succeed and others fail.

Show me the setup with the Azure.Storage.Blobs.Batch package.

## Evaluation Criteria

- `BlobBatchClient` from `Azure.Storage.Blobs.Batch` package
- `BlobBatchClient.DeleteBlobsAsync()` for bulk delete
- Custom batch via `BlobBatchClient.CreateBatch()` and `SubmitBatchAsync()`
- Batch size limits (256 operations per batch)
- Partial failure handling: `AggregateException` with per-operation status
- `RequestFailedException` for individual operation failures within a batch

## Context

Batch operations are essential for cost optimization (changing storage tiers)
and cleanup (bulk deletion). Without batch support, developers resort to
sequential API calls that are slow and consume excessive transaction units.
