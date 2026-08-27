---
id: key-vault-dp-dotnet-crud
properties:
  service: key-vault
  plane: data-plane
  language: dotnet
  category: crud
  difficulty: basic
  description: 'Can a developer create, read, update, and delete secrets in Azure Key Vault using the .NET SDK?

    '
  sdk_package: Azure.Security.KeyVault.Secrets
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/security.keyvault.secrets-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- secrets
- crud
- getting-started
---

# CRUD Secrets: Azure Key Vault (.NET)

## Prompt

I need to rotate a secret in Azure Key Vault — set a new value for an
existing secret and then purge the old deleted version. How do I do the full
lifecycle in .NET?
1. Create a secret with an initial value
2. Read it back and print its value
3. Update it to a new value
4. Delete the old secret and purge it (soft-delete enabled vault)

Authenticate securely using identity-based credentials. Include proper
error handling.

## Evaluation Criteria

The generated code should include:
- Creates a secret client with identity-based authentication
- Creates, reads, and updates secrets
- Handles soft-delete by waiting for deletion to complete before purging
- Handles errors appropriately

## Context

CRUD operations on secrets are the most fundamental Key Vault use case.
This tests whether the generated code provides a complete, runnable flow
covering the full lifecycle including soft-delete and async polling patterns.
