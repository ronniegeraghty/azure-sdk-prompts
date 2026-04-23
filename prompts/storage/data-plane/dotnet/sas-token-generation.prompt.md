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
without sharing my storage account key. How do I generate a SAS token in .NET?
1. Generate a SAS token without using the storage account key directly
2. Scope the token to read-only access with a limited expiry
3. Build a SAS URI for a specific blob
4. Verify that the SAS URI works by downloading the blob with it

Explain the different SAS types available and which is the most secure
approach.

## Evaluation Criteria

- Generates a user-delegation SAS using identity-based credentials (preferred)
- Scopes permissions to read-only with a time-limited expiry
- Builds a complete SAS URI for a specific blob
- Explains the security tradeoffs between user-delegation and account-key SAS
- Demonstrates verification of the generated SAS URI

## Context

SAS tokens are one of the most commonly used Azure Storage authentication
patterns for sharing access with external consumers. Getting the delegation
model right is a security-critical task that developers frequently ask about.
