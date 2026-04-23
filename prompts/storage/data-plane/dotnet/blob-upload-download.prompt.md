---
id: storage-dp-dotnet-blob-upload-download
service: storage
plane: data-plane
language: dotnet
category: crud
difficulty: basic
description: >
  Can a developer upload and download blobs using the Azure Blob Storage
  .NET SDK, including stream-based and file-based approaches?
sdk_package: Azure.Storage.Blobs
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
tags:
  - blob
  - upload
  - download
  - getting-started
created: 2026-04-09
author: JonathanCrd
---

# Blob Upload & Download: Azure Blob Storage (.NET)

## Prompt

I need to upload a local file to Azure Blob Storage and download it back
using .NET. How do I:
1. Connect to a blob container and ensure it exists
2. Upload a file from disk
3. Upload from a stream (e.g., in-memory data)
4. Download a blob to a local file
5. Download a blob as a stream and read its contents

Authenticate securely without hardcoding credentials. Explain the
overwrite behavior on upload.

## Evaluation Criteria

- Connects to Blob Storage using identity-based authentication
- Creates the container if it doesn't exist
- Uploads blobs from both file paths and streams
- Downloads blobs to files and as streams
- Handles overwrite behavior on upload
- Properly disposes of streams and clients

## Context

Uploading and downloading blobs is the most fundamental Azure Storage operation.
This tests whether the generated code covers both file-based and stream-based
approaches, which are the two patterns every .NET developer needs.
