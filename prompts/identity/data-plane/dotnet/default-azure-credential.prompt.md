---
id: identity-dp-dotnet-default-credential
properties:
  service: identity
  plane: data-plane
  language: dotnet
  category: auth
  difficulty: basic
  description: 'Can a developer set up DefaultAzureCredential for Azure SDK clients using the .NET SDK?

    '
  sdk_package: Azure.Identity
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- default-azure-credential
- getting-started
---

# DefaultAzureCredential: Azure Identity (.NET)

## Prompt

I keep getting authentication errors when using DefaultAzureCredential
and I don't understand where it's looking for credentials. Explain:
1. What packages are needed for Azure Identity
2. The credential chain order — which credentials are tried and in what sequence
3. How it behaves differently in local development vs deployed Azure environments
4. How to troubleshoot when authentication fails

Show a working example that creates an Azure SDK client with
DefaultAzureCredential, and explain what to check when it doesn't work.

## Evaluation Criteria

The generated code should include:
- Uses the Azure Identity package
- Creates a `DefaultAzureCredential` with appropriate options
- Explains the credential chain order
- Passes the credential to an Azure SDK client
- Handles authentication errors with diagnostics guidance

## Context

DefaultAzureCredential is the recommended starting point for Azure SDK authentication.
It abstracts away the complexity of credential selection and works across environments.
This tests whether the generated code demonstrates it clearly enough for first-time users.
