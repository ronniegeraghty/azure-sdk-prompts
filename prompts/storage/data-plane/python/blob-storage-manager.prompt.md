---
id: storage-dp-python-blob-manager
properties:
  service: storage
  plane: data-plane
  language: python
  category: crud
  difficulty: advanced
  description: 'Can an agent generate a complete, production-ready Azure Blob Storage management utility with sync and async
    implementations, covering upload (large files, index tags), download, list, delete, concurrency prevention, retry configuration,
    and HTTP logging?

    '
  sdk_package: azure-storage-blob
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/storage-blob-readme
  created: '2026-04-10'
  author: copilot
tags:
- identity
- default-azure-credential
- blob-storage
- async
- retry
- lease
- parallel-upload
- index-tags
---

# Blob Storage Manager: Azure Blob Storage (Python)

## Prompt

Create a Python project that provides a reusable Azure Blob Storage management utility.

The project needs:

- A **service module** (both sync and async versions) that wraps blob operations: upload (with optional metadata and blob index tags for later querying), download, list blobs in a container, and delete. The upload method should handle large files efficiently so that uploading a multi-gigabyte file doesn't load the entire thing into memory or fail on slow connections. The service should also prevent concurrent writers from overwriting each other's changes when updating the same blob.

- A **configuration module** that connects to Azure securely using the storage account endpoint (from an environment variable). No connection strings or account keys should be used — the app will run in Azure with `DefaultAzureCredential`. The configuration should set up a custom retry policy (exponential backoff, configurable max retries and delay), so the app behaves predictably under transient failures. It should also enable HTTP request/response logging at a configurable level for debugging.

- The service should **handle errors gracefully** — if a storage operation fails (e.g., blob not found, permission denied, or a lease is already held by another client), the error should be caught and handled with a clear message rather than crashing. Each operation should also accept an **optional timeout**, so callers can control how long they're willing to wait for a response.

- A **main script** that wires everything together and demos each operation using the sync implementation first, then repeats the same operations using the async implementation: uploads a sample file with some index tags, lists blobs, downloads the file back, acquires a lease and overwrites it, and finally deletes it. Print status at each step.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Configuration & Resilience
- Configures a custom retry policy with exponential backoff
- Enables HTTP request/response logging for debugging
- Does NOT use connection strings or account keys

### Blob Operations
- Handles large file upload efficiently without loading the entire file into memory
- Supports blob index tags on upload (distinct from metadata)
- Implements blob lease acquisition to prevent concurrent overwrites
- Provides both sync and async implementations

### Error Handling
- Catches and handles storage-specific errors from the Azure SDK
- Handles lease conflicts when another client holds a lease
- Includes per-operation timeout configuration

## Context

This is the most common Azure Storage scenario: a reusable CRUD wrapper. It tests whether
the agent knows the modern Azure SDK for Python patterns (DefaultAzureCredential, async support,
streaming upload with concurrency) vs deprecated patterns that LLMs frequently generate.
The prompt is intentionally business-level — it says "handle large files efficiently" not
"use max_concurrency" — so skills must teach the agent the right SDK approach.

Cross-cutting concerns tested: authentication, retry/timeout configuration, HTTP pipeline
logging, async patterns, blob leasing for concurrency, and blob index tags.
