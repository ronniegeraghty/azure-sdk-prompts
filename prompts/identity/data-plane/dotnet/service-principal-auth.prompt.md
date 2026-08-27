---
id: identity-dp-dotnet-service-principal
properties:
  service: identity
  plane: data-plane
  language: dotnet
  category: auth
  difficulty: intermediate
  description: 'Can a developer authenticate with a Service Principal (client secret) using the .NET SDK?

    '
  sdk_package: Azure.Identity
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- service-principal
- client-secret
---

# Service Principal Authentication: Azure Identity (.NET)

## Prompt

My CI pipeline needs to authenticate to Azure using a service principal.
How do I set up client-secret-based authentication in .NET? I need:
1. What packages are required
2. How to create a credential with tenant ID, client ID, and client secret
3. How to use the credential with Azure SDK clients
4. Best practices for storing the secret securely
5. Error handling when credentials are invalid

Show a complete example with proper error handling.

## Evaluation Criteria

The generated code should include:
- Uses a client secret credential with tenant, client ID, and secret
- Passes the credential to Azure SDK clients
- Stores secrets via environment variables or secure configuration
- Handles authentication errors for invalid credentials

## Context

Service Principal authentication with client secrets is the most common pattern
for application-to-application auth in Azure. This tests whether the generated code
covers the full setup including credential creation, usage, and secret management best practices.
