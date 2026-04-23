---
id: identity-dp-dotnet-managed-identity
properties:
  service: identity
  plane: data-plane
  language: dotnet
  category: auth
  difficulty: intermediate
  description: 'Can a developer use Managed Identity to authenticate Azure SDK clients using the .NET SDK?

    '
  sdk_package: Azure.Identity
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- managed-identity
- azure-hosted
---

# Managed Identity Authentication: Azure Identity (.NET)

## Prompt

I'm deploying my app to Azure and want to stop using connection strings
for authentication. How do I switch to managed identity for my Azure SDK
clients in .NET? I need to know:
1. System-assigned vs user-assigned managed identity — which should I pick?
2. How to create a credential for each type
3. Using the credential with Azure SDK clients
4. How to test locally when managed identity isn't available — what's the fallback?
5. Common pitfalls and error handling

Provide examples for both system-assigned and user-assigned identity.

## Evaluation Criteria

The generated code should include:
- Uses managed identity credentials (system-assigned and user-assigned)
- Passes the client ID for user-assigned identity
- Integrates with DefaultAzureCredential (managed identity in the chain)
- Handles the case where managed identity is unavailable (local dev fallback)
- Handles credential-unavailable errors appropriately

## Context

Managed Identity is the recommended auth pattern for code running in Azure.
It eliminates the need for managing secrets entirely. This tests whether the
.generated code demonstrates both system-assigned and user-assigned identity clearly,
including the critical local development fallback story.
