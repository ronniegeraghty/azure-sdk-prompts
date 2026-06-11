---
name: azure-sdk-for-python-bestpractices
description: "Best practices for generating Python code that uses the Azure SDK for Python. Use when generating Python code targeting Azure services to ensure correct package selection, credential usage, and idiomatic SDK patterns."
applyTo: "**/*.py,**/requirements.txt,**/pyproject.toml"
---

# Azure SDK for Python Best Practices

You are generating Python code that uses the **Azure SDK for Python**. Follow these best practices to produce correct, idiomatic code.

## References

- [Azure SDK for Python repository](https://github.com/Azure/azure-sdk-for-python)
- [Azure SDK Design Guidelines for Python](https://azure.github.io/azure-sdk/python_introduction.html)

## Package Selection

Use only the **official** `azure-*` packages published by the Azure SDK team. Prefer data-plane client libraries (e.g., `azure-keyvault-secrets`) over the legacy `azure-mgmt-*` libraries unless the task is explicitly a management-plane operation.

### Common Packages

| Service            | Package                    |
| ------------------ | -------------------------- |
| Identity           | `azure-identity`           |
| Key Vault Secrets  | `azure-keyvault-secrets`   |
| Blob Storage       | `azure-storage-blob`       |
| Cosmos DB          | `azure-cosmos`             |
| Service Bus        | `azure-servicebus`         |
| Event Hubs         | `azure-eventhub`           |
| App Configuration  | `azure-appconfiguration`   |

Do **not** use connection strings, shared access signatures, or account keys for authentication in sample code.

## Authentication

Authenticate with `DefaultAzureCredential` from `azure.identity`:

```python
from azure.identity import DefaultAzureCredential

credential = DefaultAzureCredential()
```

Pass the credential directly to client constructors. Never hardcode secrets. Read endpoints and resource names from environment variables.

## Client Construction

Construct service clients once and reuse them:

```python
from azure.keyvault.secrets import SecretClient

client = SecretClient(vault_url=vault_url, credential=credential)
```

Clients are thread-safe. Close clients via `with` blocks or call `client.close()` when done. For async clients, use `async with`.

## Error Handling

Catch `azure.core.exceptions.HttpResponseError` (and subclasses like `ResourceNotFoundError`, `ResourceExistsError`) rather than broad `Exception`. Let transport-level retries in the SDK handle transient failures — do not wrap every call in retry loops.

```python
from azure.core.exceptions import ResourceNotFoundError

try:
    secret = client.get_secret("my-secret")
except ResourceNotFoundError:
    ...
```

## Pagination

List operations return `ItemPaged` iterables. Iterate items directly rather than pages:

```python
for secret_properties in client.list_properties_of_secrets():
    print(secret_properties.name)
```

Use `.by_page()` only when page-level access is explicitly required.

## Long-Running Operations

LRO methods return a poller. Call `.result()` to wait, or `.wait()` / `.status()` for manual control. LRO methods are prefixed with `begin_` (e.g., `begin_create_or_update`).

## Style

- Follow PEP 8 naming (snake_case functions/variables, PascalCase classes).
- Use type hints for function signatures.
- Prefer f-strings for string formatting.
- Keep imports at the top of the module.
