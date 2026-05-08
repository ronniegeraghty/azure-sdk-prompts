---
id: identity-dp-js-ts-default-credential
properties:
  service: identity
  plane: data-plane
  language: js-ts
  category: auth
  difficulty: basic
  description: 'Can a developer set up DefaultAzureCredential for Azure SDK clients using the JavaScript/TypeScript SDK?

    '
  sdk_package: '@azure/identity'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- default-azure-credential
- getting-started
---

# DefaultAzureCredential: Azure Identity (JavaScript/TypeScript)

## Prompt

Write a TypeScript program that
authenticates an Azure SDK client using DefaultAzureCredential. The program should:
1. Install and import the required npm packages
2. Create a DefaultAzureCredential instance
3. Use it to create a SecretClient from @azure/keyvault-secrets
4. Retrieve a secret from the vault and print its value
5. Handle AuthenticationError for credential failures

Enable SDK diagnostic logging using `@azure/logger` with a configurable log level.
Include a package.json with all dependencies and use async/await throughout.

## Evaluation Criteria

The generated code should include:
- `@azure/identity` npm package installation
- `DefaultAzureCredential` constructor and options
- Credential chain: Environment → Workload Identity → Managed Identity → Azure CLI → etc.
- Passing credential to Azure SDK clients
- `AuthenticationError` handling and logging

## Context

DefaultAzureCredential is the recommended starting point for Azure SDK authentication.
It abstracts away the complexity of credential selection and works across environments.
This tests whether the generated code demonstrates it clearly enough for first-time users.
