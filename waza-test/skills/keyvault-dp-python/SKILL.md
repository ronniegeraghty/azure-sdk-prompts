---
name: keyvault-dp-python
type: utility
description: |
  USE FOR: Azure Key Vault secrets operations in Python (CRUD, pagination, error handling).
  DO NOT USE FOR: Key Vault keys/certificates, non-Python languages, control plane operations.
---

# Azure Key Vault Data Plane — Python SDK

## Overview

Generates Python code for Azure Key Vault secrets operations using
`azure-keyvault-secrets` and `azure-identity`. Covers CRUD lifecycle,
pagination with `ItemPaged`, and error handling with `HttpResponseError`.

## Usage

- "Write a script that creates, reads, updates, and deletes Azure Key Vault secrets"
- "How do I list all secrets in a large Key Vault with pagination?"
- "How do I handle 403/404/429 errors with azure-keyvault-secrets?"

## References

- [azure-keyvault-secrets PyPI](https://pypi.org/project/azure-keyvault-secrets/)
- [Azure Key Vault Secrets client library for Python](https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme)
