---
id: storage-dp-python-crud
properties:
  service: storage
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: 'Can a developer upload, download, list, and delete blobs in Azure Blob Storage using the Python SDK?

    '
  sdk_package: azure-storage-blob
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/storage-blob-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- blob
- crud
- getting-started
---

# CRUD Blobs: Azure Blob Storage (Python)

## Prompt

Write a Python script that performs
CRUD operations on Azure Blob Storage:

**Write the code to files (use file-write tools, do not reply with code blocks).**

1. Create a BlobServiceClient using DefaultAzureCredential
2. Create a container named "my-container" if it doesn't exist
3. Upload a local file "report.csv" as a blob named "reports/report.csv"
4. List all blobs in the container and print each blob's name and content length
5. Download the blob and save it to "report-downloaded.csv"
6. Delete the blob and then delete the container

Include a `requirements.txt` with `azure-storage-blob` and `azure-identity`, and add proper error handling with `HttpResponseError` and `ResourceExistsError`.

## Evaluation Criteria

- Includes required Azure Storage and Identity SDK packages (via requirements.txt or install instructions)
- Authenticates the blob service client with DefaultAzureCredential
- Creates a container (handling the case where it already exists)
- Uploads a file as a blob with overwrite support
- Lists blobs in the container with their properties
- Downloads a blob to a local file
- Deletes the blob and the container
- Catches Azure-specific errors (both general HTTP errors and resource-exists errors)

## Context

Python is the most popular language for data engineering and ML workflows
that use Azure Blob Storage. This tests whether the generated code provides a
complete blob lifecycle tutorial that a data engineer can follow end-to-end.
