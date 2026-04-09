---
id: storage-dp-dotnet-sas-token
service: storage
plane: data-plane
language: dotnet
category: authentication
difficulty: intermediate
description: >
  Can a developer generate SAS tokens for Azure Blob Storage
  using both account-level and service-level delegation in .NET?
sdk_package: Azure.Storage.Blobs
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
tags:
  - sas
  - shared-access-signature
  - security
  - authentication
created: 2026-04-09
author: JonathanCrd
---

# SAS Token Generation: Azure Blob Storage (.NET)

## Prompt

I need to give a third-party temporary read-only access to a specific blob
without sharing my storage account key. How do I generate a SAS token in C#?
1. Create a user-delegation SAS using DefaultAzureCredential (no account key needed)
2. Set permissions to read-only with a 1-hour expiry
3. Generate the SAS URI for a specific blob
4. Show how to verify the SAS URI works by downloading the blob with it

Also explain when I'd use a user-delegation SAS vs. an account-key SAS
and why user-delegation is preferred.

## Evaluation Criteria

- `BlobSasBuilder` for defining SAS parameters
- `BlobServiceClient.GetUserDelegationKeyAsync()` for user-delegation SAS
- `BlobSasPermissions.Read` for read-only access
- Setting `ExpiresOn` with a `DateTimeOffset`
- `BlobUriBuilder` or `BlobClient.GenerateSasUri()` for URI generation
- Explanation of user-delegation vs. account-key SAS security tradeoffs

## Context

SAS tokens are one of the most commonly used Azure Storage authentication
patterns for sharing access with external consumers. Getting the delegation
model right is a security-critical task that developers frequently ask about.
