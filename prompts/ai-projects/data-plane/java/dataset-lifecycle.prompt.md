---
id: ai-projects-dp-java-dataset-lifecycle
properties:
  service: ai-projects
  plane: data-plane
  language: java
  category: projects
  difficulty: advanced
  description: "Can an application upload, retrieve, download, verify, and delete one exact Azure AI project dataset version?"
  sdk_package: com.azure:azure-ai-projects
  doc_url: https://learn.microsoft.com/en-us/java/api/overview/azure/ai-projects-readme
  created: "2026-08-26"
  author: weidongxu-microsoft
tags:
  - foundry
  - ai-projects
  - datasets
  - lifecycle
  - storage-blob
---

# Azure AI project dataset lifecycle (Java)

## Prompt

Create a complete, runnable Java console application using
`com.azure:azure-ai-projects` 2.4.0 and `com.azure:azure-storage-blob` 12.35.1.

**Write the application to files in the workspace. Do not reply with code blocks.**

The application must:

- Read `FOUNDRY_PROJECT_ENDPOINT`, `DATASET_NAME`, `DATASET_VERSION`,
  `DATA_FILE_PATH`, and `DOWNLOAD_FILE_PATH`.
- Upload the local file at `DATA_FILE_PATH` as the exact project dataset name and
  version supplied by the user, using the project's default Azure Storage
  connection.
- Retrieve that exact dataset version and print its SDK-defined name, version, ID,
  type, and data URL.
- Request the dataset's SDK credential and blob reference. Use them with the Azure
  Storage Blob SDK to download the referenced blob to `DOWNLOAD_FILE_PATH` without
  displaying the SAS URL.
- Compare the downloaded bytes with the original file and fail if they differ.
  Print the verified byte count only after the comparison succeeds.
- Delete the exact dataset name and version even when retrieval or download fails
  after upload.
- Include the project manifest and concise restore, build, and run commands.

Use the synchronous SDK clients throughout.

## Evaluation Criteria

### Dataset lifecycle workflow

- Builds `DatasetsClient` with `AIProjectClientBuilder` and calls
  `createDatasetWithFile` with the supplied name, version, and `Path`.
- Calls `getDatasetVersion` with the same name and version and prints typed
  `DatasetVersion` metadata, including its type and data URL.
- Calls `getCredentials` for that exact version, combines
  `DatasetCredential.getBlobReference().getBlobUrl()` with its SAS credential, and
  downloads the correct blob through `BlobClient`.
- Persists the blob at `DOWNLOAD_FILE_PATH`, compares its bytes with the original
  file, and prints a byte count derived from the verified download.
- Calls `deleteDatasetVersion` with the exact name and version in cleanup after a
  successful upload.

### Scenario-specific anti-patterns

- Does not invent a `downloadDataset` API.
- Does not print the SAS URL, hardcode downloaded content, or skip the byte
  comparison.
- Does not replace the project dataset upload or Storage Blob download with direct
  HTTP calls.
- Does not delete only by dataset name or omit cleanup on a post-upload failure.

## Context

The reference application is in
`reference-apps/ai-projects/dataset-lifecycle/java`.
