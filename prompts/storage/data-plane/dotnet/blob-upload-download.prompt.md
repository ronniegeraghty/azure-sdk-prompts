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

I need to upload a local file to Azure Blob Storage and download it back.
How do I do this with the Azure.Storage.Blobs SDK in C#?
1. Create a BlobContainerClient and ensure the container exists
2. Upload a file from disk using BlobClient.UploadAsync
3. Upload a stream (e.g., a MemoryStream) as a blob
4. Download a blob to a local file
5. Download a blob as a stream and read its contents

Use DefaultAzureCredential for auth. Show required NuGet packages and explain
the overwrite behavior on upload.

## Evaluation Criteria

- `Azure.Storage.Blobs` NuGet package
- `BlobServiceClient` or `BlobContainerClient` creation with `DefaultAzureCredential`
- `BlobContainerClient.CreateIfNotExistsAsync()`
- `BlobClient.UploadAsync()` with file path and stream overloads
- `BlobClient.DownloadToAsync()` or `DownloadContentAsync()`
- Overwrite parameter on upload (`overwrite: true`)
- Proper stream disposal with `using` or `await using`

## Context

Uploading and downloading blobs is the most fundamental Azure Storage operation.
This tests whether the generated code covers both file-based and stream-based
approaches, which are the two patterns every .NET developer needs.
