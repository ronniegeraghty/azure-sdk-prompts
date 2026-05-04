---
id: storage-dp-js-ts-blob-manager
properties:
  service: storage
  plane: data-plane
  language: js-ts
  category: crud
  difficulty: advanced
  description: >
    Can an agent generate a complete, production-ready Azure Blob Storage
    management utility in TypeScript with streaming upload for large files,
    download with Node.js stream handling, blob leasing for concurrency
    control, index tags, retry configuration, and SDK logging?
  sdk_package: '@azure/storage-blob'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/storage-blob-readme
  created: '2026-04-30'
  author: kaghiya
tags:
  - identity
  - default-azure-credential
  - blob-storage
  - streaming
  - retry
  - lease
  - index-tags
  - logging
---

# Blob Storage Manager: Azure Blob Storage (TypeScript)

## Prompt

Create a TypeScript Node.js project that provides a reusable Azure Blob Storage management utility.

The project needs:

- A **service class** that wraps blob operations: upload (with optional metadata and blob index tags for later querying), download, list blobs in a container, and delete. The upload method should handle large files efficiently using streaming — uploading a multi-gigabyte file should not load the entire thing into memory. The service should also prevent concurrent writers from overwriting each other's changes when updating the same blob by acquiring a lease before writing.

- A **configuration module** that connects to Azure securely using the storage account endpoint (from an environment variable). No connection strings or account keys should be used — the app will run in Azure with managed identity. The configuration should set up a custom retry policy (exponential backoff, configurable max retries and delay) and enable SDK logging at a configurable level for debugging.

- A **main script** that wires everything together and demos each operation: uploads a sample file with some index tags, lists all blobs in the container, downloads the file back and prints its content, acquires a lease and overwrites the blob, then finally deletes it. Print status at each step.

Include a complete `package.json` with the necessary Azure SDK dependencies and a `tsconfig.json`.

## Evaluation Criteria

### Scenario-Specific Patterns
- Configures custom retry policy via `StorageRetryOptions` (exponential backoff, max retries, retry delay)
- Enables SDK logging via `@azure/logger` `setLogLevel()` or `AZURE_LOG_LEVEL`
- Implements blob lease acquisition before overwrite using `BlobLeaseClient`
- Uses `BlockBlobClient.uploadStream()` for large file streaming upload (not `uploadData()` or `upload()` which buffer in memory)
- Sets blob index tags on upload via `tags` property in `BlockBlobUploadStreamOptions`
- Downloads blob and reads response via `readableStreamBody` (Node.js Readable stream)
- Lists blobs using `for await...of` async iteration over `ContainerClient.listBlobsFlat()`

### Scenario-Specific Error Handling
- Handles lease conflict errors (409 status code) when blob is already leased
- Handles blob not found errors (404 status code) on download/delete

### Anti-Patterns (scenario-specific)
- NOT using `uploadData()` or `upload()` with full buffer for large files
- NOT using connection strings or account keys for authentication
- NOT collecting all listed blobs into an array before processing

## Context

This is the most common Azure Storage scenario: a reusable CRUD wrapper. It tests whether
the agent knows the modern @azure/storage-blob v12 patterns (DefaultAzureCredential,
streaming upload, BlobLeaseClient, index tags) versus older deprecated patterns. The prompt
is intentionally business-level — it says "handle large files efficiently using streaming"
not "use uploadStream()" — so skills must teach the agent the right SDK approach.

Cross-cutting concerns tested: authentication, retry configuration, SDK logging, streaming
I/O, blob leasing for concurrency, and blob index tags.
