---
id: identity-dp-js-ts-managed-identity
properties:
  service: identity
  plane: data-plane
  language: js-ts
  category: auth
  difficulty: intermediate
  description: 'Can a developer use Managed Identity to authenticate Azure SDK clients using the JavaScript/TypeScript SDK?

    '
  sdk_package: '@azure/identity'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- managed-identity
- azure-hosted
---

# Managed Identity Authentication: Azure Identity (JavaScript/TypeScript)

## Prompt

Write a TypeScript program that
uses Managed Identity to authenticate Azure SDK clients. The program should:
1. Create a ManagedIdentityCredential for system-assigned identity
2. Create a ManagedIdentityCredential for user-assigned identity (with client ID)
3. Use ChainedTokenCredential to fall back to Azure CLI credential for local development
4. Pass the credential to an Azure SDK client and perform an operation
5. Handle CredentialUnavailableError when not running in Azure

Include a package.json with all dependencies and use async/await throughout.

## Evaluation Criteria

The generated code should include:
- `ManagedIdentityCredential` class from `@azure/identity`
- System-assigned: no parameters needed
- User-assigned: passing the client ID in options
- Integration with `DefaultAzureCredential` chain
- `CredentialUnavailableError` when not running in Azure
- `ChainedTokenCredential` for local fallback

## Context

Managed Identity is the recommended auth pattern for code running in Azure.
It eliminates the need for managing secrets entirely. This tests whether the
JavaScript/generated code demonstrates both system-assigned and user-assigned identity clearly,
including the critical local development fallback story.
