---
id: ai-projects-dp-js-ts-dataset-lifecycle
properties:
  service: ai-projects
  plane: data-plane
  language: js-ts
  category: projects
  difficulty: advanced
  description: "Can an application upload, retrieve, download, verify, and delete one exact Azure AI project dataset version?"
  sdk_package: "@azure/ai-projects"
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/ai-projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - datasets
  - lifecycle
  - storage-blob
---

# Azure AI project dataset lifecycle (JavaScript/TypeScript)

## Prompt

Create a complete, runnable TypeScript console application using
`@azure/ai-projects` 2.5.0 and `@azure/storage-blob` 12.33.0. Target Node.js 22.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `FOUNDRY_PROJECT_ENDPOINT`, `DATASET_NAME`, `DATASET_VERSION`,
  `DATA_FILE_PATH`, and `DOWNLOAD_FILE_PATH`.
- Upload the local file at `DATA_FILE_PATH` as the exact project dataset name and
  version supplied by the user, using the project's default Azure Storage
  connection.
- Retrieve that exact dataset version and print its SDK-defined name, version, ID,
  type, and data URI.
- Request the dataset's SDK credential and blob reference. Use them with the Azure
  Storage Blob SDK to download the referenced blob to `DOWNLOAD_FILE_PATH` without
  displaying the SAS URI.
- Compare the downloaded bytes with the original file and fail if they differ.
  Print the verified byte count only after the comparison succeeds.
- Delete the exact dataset name and version even when retrieval or download fails
  after upload.
- Include the project manifest and concise install, build, and run commands.

Use async/await throughout.

## Evaluation Criteria

### Dataset lifecycle workflow

- Creates `AIProjectClient` and calls `client.datasets.uploadFile` with the supplied
  name, version, and file path.
- Calls `client.datasets.get` with the same name and version and prints typed
  dataset-version metadata, including `type` and `dataUri`.
- Calls `client.datasets.getCredentials` for that exact version, uses both
  `credential.blobReference.blobUri` and its `credential.sasUri`, and downloads the
  correct blob through `BlobClient` from `@azure/storage-blob`.
- Persists the blob at `DOWNLOAD_FILE_PATH`, compares its bytes with the original
  file, and prints a byte count derived from the verified download.
- Calls `client.datasets.delete` with the exact name and version in cleanup after a
  successful upload.

### Scenario-specific anti-patterns

- Does not invent a `client.datasets.download` API.
- Does not print the SAS URI, hardcode downloaded content, or skip the byte
  comparison.
- Does not replace the project dataset upload or Storage Blob download with `fetch`
  or another direct HTTP call.
- Does not delete only by dataset name or omit cleanup on a post-upload failure.

## Context

The reference application is in
`reference-apps/ai-projects/dataset-lifecycle/js-ts`.
