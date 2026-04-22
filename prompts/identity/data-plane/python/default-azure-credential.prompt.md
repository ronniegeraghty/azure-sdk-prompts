---
id: identity-dp-python-default-credential
properties:
  service: identity
  plane: data-plane
  language: python
  category: auth
  difficulty: basic
  description: 'Can a developer set up DefaultAzureCredential for Azure SDK clients using the Python SDK?

    '
  sdk_package: azure-identity
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- authentication
- default-azure-credential
- getting-started
---

# DefaultAzureCredential: Azure Identity (Python)

## Prompt

Create a runnable Python project that demonstrates authenticating an Azure SDK
client using `DefaultAzureCredential`. You MUST write the following files to
the workspace:

1. `requirements.txt` — pinning `azure-identity` and `azure-keyvault-secrets`
2. `main.py` — a complete, executable script that:
   - constructs a `DefaultAzureCredential`
   - passes it to a `SecretClient` for an Azure Key Vault
   - reads a secret and prints its value
   - configures the `logging` module so credential-chain attempts are visible
   - handles `ClientAuthenticationError` with a useful error message
3. `README.md` — a short README that explains, in this order:
   1. What pip packages are needed and how to install them
   2. How `DefaultAzureCredential` is created and used
   3. The credential chain order (Environment → Workload Identity → Managed
      Identity → Azure CLI → etc.) and which credentials are tried
   4. How it works in local development (VS Code, Azure CLI) vs. Azure
      deployments
   5. How to troubleshoot authentication failures using the logging output
      from `main.py`

Do not just describe the solution in chat — write the files. The evaluation
inspects the workspace, not your assistant response.

## Evaluation Criteria

The generated code should include:
- `azure-identity` pip package installation
- `DefaultAzureCredential()` constructor and keyword arguments
- Credential chain: Environment → Workload Identity → Managed Identity → Azure CLI → etc.
- Passing credential to Azure SDK clients
- `ClientAuthenticationError` handling and `logging` module configuration

## Context

DefaultAzureCredential is the recommended starting point for Azure SDK authentication.
It abstracts away the complexity of credential selection and works across environments.
This tests whether the generated code demonstrates it clearly enough for first-time users.
