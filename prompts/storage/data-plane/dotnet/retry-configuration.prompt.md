---
id: storage-dp-dotnet-retries
properties:
  service: storage
  plane: data-plane
  language: dotnet
  category: retries
  difficulty: advanced
  description: 'Can a developer configure custom retry policies for Azure Blob Storage including exponential backoff and per-operation
    timeouts in .NET?

    '
  sdk_package: Azure.Storage.Blobs
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- retries
- retry-policy
- resilience
- exponential-backoff
---

# Retry Configuration: Azure Blob Storage (.NET)

## Prompt

My blob uploads sometimes fail with transient errors under load.
How do I configure retry policies with exponential backoff for
Azure Blob Storage in .NET? I need to:
1. Set a custom retry policy with multiple retries and exponential backoff
2. Configure per-operation timeouts so a single upload doesn't hang forever
3. Understand which errors the SDK retries automatically vs. ones I need to handle myself

## Evaluation Criteria

- Configures custom retry options (max retries, delay, mode)
- Sets per-request network timeouts
- Explains which HTTP status codes are retried automatically vs. non-retryable
- Uses cancellation tokens for operation-level timeout control

## Context

Default retry policies work for simple scenarios, but production applications
need fine-tuned retry behavior. Developers building resilient storage pipelines
need to understand the full retry model and when to override defaults.
