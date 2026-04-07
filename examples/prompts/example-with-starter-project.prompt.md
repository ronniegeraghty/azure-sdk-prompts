---
id: example-starter-project
properties:
  service: key-vault
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: 'Example prompt demonstrating starter_project frontmatter field'
  sdk_package: azure-keyvault-secrets
  doc_url: https://learn.microsoft.com/python/api/azure-keyvault-secrets
  created: '2025-04-04'
  author: hyoka-examples
tags:
- example
- starter-project
starter_project: |
  # Starter project files
  src/
    __init__.py
    main.py  # Contains boilerplate KeyVault client setup
  requirements.txt  # Has azure-keyvault-secrets==4.8.0
  .env.example  # Shows AZURE_KEY_VAULT_URL format
---

# Example: CRUD Operations with Starter Project

## Prompt

I have a Python project already set up with the Azure Key Vault Secrets SDK.
Extend the existing `main.py` to add functions for:
- Creating a secret
- Reading a secret
- Updating a secret value
- Deleting a secret

Use the existing KeyVault client that's already initialized in the starter code.

## Evaluation Criteria

The generated code should:
- Work with the existing project structure
- Add CRUD functions without breaking existing initialization
- Include error handling for each operation
- Use async/await if the starter project uses async patterns

## Context

This example demonstrates how the `starter_project` field provides context
about existing project structure. The generator should extend rather than
create from scratch.
