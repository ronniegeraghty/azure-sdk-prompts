---
id: identity-dp-js-ts-service-principal
properties:
  service: identity
  plane: data-plane
  language: js-ts
  category: auth
  difficulty: intermediate
  description: 'Can a developer authenticate with a Service Principal (client secret) using the JavaScript/TypeScript SDK?

    '
  sdk_package: '@azure/identity'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- service-principal
- client-secret
---

# Service Principal Authentication: Azure Identity (JavaScript/TypeScript)

## Prompt

Write a TypeScript program that
authenticates to Azure using a Service Principal with client secret. The program should:
1. Install and import @azure/identity with ClientSecretCredential
2. Read tenant ID, client ID, and client secret from environment variables
3. Create a ClientSecretCredential instance
4. Use it to create an Azure SDK client (e.g., SecretClient from @azure/keyvault-secrets)
5. Perform an operation to verify the credential works
6. Handle AuthenticationError for invalid credentials

Include a package.json with all dependencies, use dotenv for environment variable
management, and use async/await throughout.

## Evaluation Criteria

The generated code should include:
- `@azure/identity` package with `ClientSecretCredential` class
- Constructor parameters: tenantId, clientId, clientSecret
- Passing credential to Azure SDK clients
- dotenv or environment variable patterns
- `AuthenticationError` handling

## Context

Service Principal authentication with client secrets is the most common pattern
for application-to-application auth in Azure. This tests whether the generated code
covers the full setup including credential creation, usage, and secret management best practices.
