---
id: key-vault-dp-dotnet-crud
service: key-vault
plane: data-plane
language: dotnet
category: crud
difficulty: basic
description: >
  Can a developer create, read, update, and delete secrets in Azure Key Vault
  using the .NET SDK?
sdk_package: Azure.Security.KeyVault.Secrets
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/security.keyvault.secrets-readme
tags:
  - secrets
  - crud
  - getting-started
created: 2025-07-28
author: ronniegeraghty
---

# CRUD Secrets: Azure Key Vault (.NET)

## Prompt

I need to rotate a secret in Azure Key Vault — set a new value for an
existing secret and then purge the old deleted version. How do I do the full
lifecycle with SecretClient?
1. Create a secret called "my-secret" with an initial value
2. Read it back and print its value
3. Update it to a new value
4. Delete the old secret and purge it (soft-delete enabled vault)

Use DefaultAzureCredential for authentication. Include proper error handling
and show required NuGet packages.

## Evaluation Criteria

The generated code should include:
- Installing `Azure.Security.KeyVault.Secrets` and `Azure.Identity` NuGet packages
- Creating a `SecretClient` with vault URI and credential
- `SetSecret()`, `GetSecret()`, `StartDeleteSecret()`, `PurgeDeletedSecret()`
- Handling soft-delete (polling `DeleteSecretOperation` to completion before purge)
- Exception handling for `RequestFailedException`

## Context

CRUD operations on secrets are the most fundamental Key Vault use case.
This tests whether the generated code provides a complete, runnable flow
covering the full lifecycle including soft-delete and async polling patterns.
