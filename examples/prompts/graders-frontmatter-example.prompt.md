---
id: key-vault-dp-python-graders-example
properties:
  service: key-vault
  plane: data-plane
  language: python
  category: crud
  difficulty: basic
  description: >
    Example prompt demonstrating the optional `graders:` frontmatter field
    for specifying prompt-level evaluation criteria.
  sdk_package: azure-keyvault-secrets
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/keyvault-secrets-readme
  created: '2025-04-21'
  author: oracle

tags:
  - key-vault
  - crud
  - example
  - graders-frontmatter

# Example: Optional prompt-level graders
# When specified, these graders are evaluated alongside any file-level
# or config-level criteria. Graders defined here take precedence over
# config-level graders with the same name.
graders:
  - name: "SDK Import Structure"
    weight: 1.0
    kind: prompt_review
    prompt: >
      Verify that imports follow the azure-keyvault-secrets pattern:
      `from azure.keyvault.secrets import SecretClient` (not legacy
      `from azure.keyvault import SecretClient`).

  - name: "Credential Handling"
    weight: 1.0
    kind: prompt_review
    prompt: >
      Code uses DefaultAzureCredential for authentication, not
      connection strings or account keys. The credential is passed
      to the SecretClient constructor.

  - name: "Client Lifecycle"
    weight: 1.0
    kind: file
    config:
      path: "main.py"
      pattern: "with.*SecretClient|finally.*close"
      must_exist: true
    gate: true

---

# Key Vault Secrets CRUD: Azure Key Vault (Python)

## Prompt

Write a Python script that demonstrates CRUD operations on Azure Key Vault secrets:

1. **Create:** Set a new secret in the vault
2. **Read:** Retrieve a secret by name
3. **Update:** Modify an existing secret value
4. **Delete:** Remove a secret from the vault

Use the `azure-keyvault-secrets` SDK with:
- `SecretClient` for vault operations
- `DefaultAzureCredential` for authentication
- Proper error handling for `ResourceNotFoundError` and `HttpResponseError`
- Context manager for client cleanup

## Evaluation Criteria

The generated code should demonstrate:

- Correct import: `from azure.keyvault.secrets import SecretClient`
- Client initialization with vault URL and credential
- Set/get/delete operations using `set_secret()`, `get_secret()`, `delete_secret()`
- Proper exception handling for missing secrets
- Context manager or explicit client closure
- No hardcoded credentials (uses environment variables)

## Context

This example demonstrates the optional `graders:` frontmatter field in prompt YAML.
When specified, prompt-level graders allow fine-grained control over evaluation
criteria without requiring separate criteria files. This is useful for:

- Prompts with unique evaluation needs
- Tests or demonstrations requiring specific graders
- Rapid iteration on evaluation criteria before committing to file-level criteria

Prompt-level graders combine with file-level and config-level criteria.
