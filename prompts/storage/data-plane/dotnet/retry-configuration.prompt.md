---
id: storage-dp-dotnet-retries
service: storage
plane: data-plane
language: dotnet
category: retries
difficulty: advanced
description: >
  Can a developer configure custom retry policies for Azure Blob Storage
  including exponential backoff and per-operation timeouts in .NET?
sdk_package: Azure.Storage.Blobs
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
tags:
  - retries
  - retry-policy
  - resilience
  - exponential-backoff
created: 2025-07-27
author: ronniegeraghty
---

# Retry Configuration: Azure Blob Storage (.NET)

## Prompt

My blob uploads sometimes fail with 503 Service Unavailable under load.
How do I configure retry policies with exponential backoff for
Azure Blob Storage in .NET? I need to:
1. Set a custom retry policy with 5 max retries and exponential backoff
2. Configure per-operation timeouts so a single upload doesn't hang forever
3. Understand which HTTP errors the SDK retries automatically vs. ones I need to handle myself

Show me how to configure BlobClientOptions with custom RetryOptions.
Use the Azure.Storage.Blobs SDK.

## Evaluation Criteria

- `BlobClientOptions.Retry` configuration with `RetryOptions`
- `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed)
- `NetworkTimeout` for per-request timeouts
- Default retryable status codes (408, 429, 500, 502, 503, 504)
- Non-retryable errors (400, 401, 403, 404, 409)
- Per-operation `CancellationToken` for timeout control

## Context

Default retry policies work for simple scenarios, but production applications
need fine-tuned retry behavior. Developers building resilient storage pipelines
need to understand the full retry model and when to override defaults.
