---
id: identity-dp-dotnet-managed-identity
service: identity
plane: data-plane
language: dotnet
category: auth
difficulty: intermediate
description: >
  Can a developer use Managed Identity to authenticate Azure SDK clients
  using the .NET SDK?
sdk_package: Azure.Identity
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/identity-readme
tags:
  - authentication
  - managed-identity
  - azure-hosted
created: 2025-07-28
author: ronniegeraghty
---

# Managed Identity Authentication: Azure Identity (.NET)

## Prompt

I'm deploying to Azure App Service and want to stop using connection strings.
How do I switch to managed identity for my Azure SDK clients in C#? I need to know:
1. System-assigned vs user-assigned managed identity — which should I pick?
2. How to create a ManagedIdentityCredential for each type
3. Using it with a BlobServiceClient or SecretClient
4. How to test locally when managed identity isn't available — what's the fallback?
5. Common pitfalls and error handling

Provide examples for both system-assigned and user-assigned identity.

## Evaluation Criteria

The generated code should include:
- `ManagedIdentityCredential` class and constructors
- System-assigned: no parameters needed
- User-assigned: passing the client ID
- Integration with `DefaultAzureCredential` (managed identity in the chain)
- `CredentialUnavailableException` when not running in Azure
- Combining with `ChainedTokenCredential` for local fallback

## Context

Managed Identity is the recommended auth pattern for code running in Azure.
It eliminates the need for managing secrets entirely. This tests whether the
.generated code demonstrates both system-assigned and user-assigned identity clearly,
including the critical local development fallback story.
