---
id: key-vault-dp-dotnet-pagination
properties:
  service: key-vault
  plane: data-plane
  language: dotnet
  category: pagination
  difficulty: intermediate
  description: 'Can a developer paginate through a large list of Key Vault secrets using the .NET SDK?

    '
  sdk_package: Azure.Security.KeyVault.Secrets
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/security.keyvault.secrets-readme
  created: '2025-07-27'
  author: ronniegeraghty
tags:
- pagination
- list-secrets
- async
---

# Pagination: List Key Vault Secrets (.NET)

## Prompt

My Key Vault has hundreds of secrets and I need to enumerate them all
without loading everything into memory. How do I paginate through them
in .NET?
1. Connect to Key Vault with identity-based authentication
2. Iterate through secrets page by page
3. Print the name, content type, and enabled status of each secret
4. Handle the case where some secrets are disabled

I want to understand how async pagination works for large result sets.

## Evaluation Criteria

- Lists secret properties using async pagination
- Supports both simple async iteration and explicit page-by-page control
- Accesses secret metadata (name, content type, enabled status)
- Handles errors during pagination

## Context

Key Vaults in enterprise environments often contain hundreds of secrets.
Developers need to understand the AsyncPageable pattern to efficiently
enumerate them without loading all results into memory at once.
