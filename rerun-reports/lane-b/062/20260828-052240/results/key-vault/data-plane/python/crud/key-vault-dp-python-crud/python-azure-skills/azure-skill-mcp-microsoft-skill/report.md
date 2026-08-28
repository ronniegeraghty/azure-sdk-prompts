# Evaluation Report: key-vault-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 360.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 11/12 |
| Duration | 360.4s |
| Timestamp | 2026-08-27T21:22:40Z |
| Files Generated | 2 |
| Event Count | 3016 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 269.3s |
| Review | 89.9s |
| **Total** | **360.4s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-keyvault-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=231138 out=2622 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3396ms

**Input:**

```json
{"skill":"azure-keyvault-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Key Vault SDK for Python

Secure storage and management for secrets, cryptographic keys, and certificates.

## Installation

```bash
# Secrets
pip install azure-keyvault-secrets azure-identity

# Keys (cryptographic operations)
pip install azure-keyvault-keys azure-identity

# Certificates
pip install azure-keyvault-certificates azure-identity

# All
pip install azure-keyvault-secrets azure-keyvault-keys azure-keyvault-certificates azure-identity
```

## Environment Variables

```bash
AZURE_KEYVAULT_URL=https://<vault-name>.vault.azure.net/  # Required for all auth methods
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication & Lifecycle

> **🔑 Two rules apply to every code sample below:**
>
> 1. **Prefer `DefaultAzureCredential`.** It works locally (Azure CLI / VS Code / Developer CLI) and in Azure (managed identity, workload identity) with no code change. Avoid connection strings, account/API keys — they bypass Entra audit and rotation.
>    - Local dev: `DefaultAzureCredential` works as-is.
>    - Production: set `AZURE_TOKEN_CREDENTIALS=prod` (or `AZURE_TOKEN_CREDENTIALS=<specific_credential>`) to constrain the credential chain to production-safe credentials.
> 2. **Wrap every client in a context manager** so HTTP transports, sockets, and token caches are released deterministically:
>    - Sync: `with <Client>(...) as client:`
>    - Async: `async with <Client>(...) as client:` **and** `async with DefaultAzureCredential() as credential:` (from `azure.identity.aio`)
>
> Snippets may abbreviate this setup, but production code should always follow both rules.

## Secrets

### SecretClient Setup

```python
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential
from azure.keyvault.secrets import SecretClient

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
cre
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 193336ms

**Input:**

```json
{"intent":"Generate a local Python script using DefaultAzureCredential and Azure Key Vault SecretClient to create, read, update, delete, and purge a secret with robust error handling. Do not deploy or modify Azure infrastructure.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. view — 193333ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-keyvault-py\\references\\capabilities.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/capabilities.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/capabilities.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/capabilities.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/capabilities.md
@@ -1,33 +1,33 @@
 # azure-keyvault-py capability coverage
 
 **SDK/package**: `azure-keyvault-secrets, azure-keyvault-keys, azure-keyvault-certificates`
 
 This index maps hero scenarios in `SKILL.md` and links non-hero scenarios documented in dedicated reference files.
 
 ## Hero scenarios covered in SKILL.md
 
 - `Secrets`
 - `Keys`
 - `Certificates`
 - `Client Types Table`
 
 ## Non-hero scenarios
 
 - `Async Clients`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#async-clients`](non-hero-scenarios.md#async-clients)
 - `Error Handling`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#error-handling`](non-hero-scenarios.md#error-handling)
 
 ## Related deep-dive references
 
 - [`non-hero-scenarios.md`](non-hero-scenarios.md): Dedicated non-hero examples and implementation notes.
 
 ## API breadth checklist
 
 - Verify client/auth mode for the environment before coding.
 - Confirm operation-group/method names against current Microsoft Learn API reference.
 - For Python SDKs with both sync and async clients, document both forms without a blanket preference.
 - Include cleanup/delete paths for created resources in examples.
 - Prefer idempotent create/update operations where available.
 - Validate paging/LRO/error-handling patterns for production paths.
 

```

### ✅ 4. view — 192066ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-keyvault-py\\references\\non-hero-scenarios.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/non-hero-scenarios.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/non-hero-scenarios.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/non-hero-scenarios.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-keyvault-py/references/non-hero-scenarios.md
@@ -1,36 +1,36 @@
 # azure-keyvault-py non-hero scenarios
 
 These scenarios are intentionally separate from hero flows in `SKILL.md`.
 They cover secondary/advanced patterns typically used after the primary end-to-end path is working.
 
 ## Async Clients
 
 ```python
 from azure.identity.aio import DefaultAzureCredential
 from azure.keyvault.secrets.aio import SecretClient
 
 async def get_secret():
     async with DefaultAzureCredential() as credential:
         async with SecretClient(vault_url=vault_url, credential=credential) as client:
             secret = await client.get_secret("my-secret")
             print(f"Retrieved secret: {secret.name} (version: {secret.properties.version})")
 
 import asyncio
 asyncio.run(get_secret())
 ```
 
 ## Error Handling
 
 ```python
 from azure.core.exceptions import ResourceNotFoundError, HttpResponseError
 
 try:
     secret = client.get_secret("nonexistent")
 except ResourceNotFoundError:
     print("Secret not found")
 except HttpResponseError as e:
     if e.status_code == 403:
         print("Access denied - check RBAC permissions")
     raise
 ```
 

```

### ✅ 5. glob — 190748ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2465558328","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. glob — 189508ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-keyvault-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. apply_patch — 2619ms

**Input:**

```json
"*** Begin Patch\n*** Add File: key_vault_secret_crud.py\n+\"\"\"Create, read, update, delete, and purge an Azure Key Vault secret.\"\"\"\n+\n+import os\n+import sys\n+\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceNotFoundError,\n+)\n+from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n+from azure.keyvault.secrets import SecretClient\n+\n+\n+SECRET_NAME = os.getenv(\"AZURE_KEYVAULT_SECRET_NAME\", \"crud-demo-secret\")\n+INITIAL_VALUE = \"initial-value\"\n+UPDATED_VALUE = \"updated-value\"\n+\n+\n+def get_vault_url() -\u003e str:\n+    \"\"\"Return and validate the Key Vault URL from the environment.\"\"\"\n+    vault_url = os.getenv(\"AZURE_KEYVAULT_URL\")\n+    if not vault_url:\n+        raise ValueError(\n+            \"AZURE_KEYVAULT_URL is required, for example \"\n+            \"https://\u003cvault-name\u003e.vault.azure.net/\"\n+        )\n+\n+    if not vault_url.startswith(\"https://\") or \".vault.azure.net\" not in vault_url:\n+        raise ValueError(\"AZURE_KEYVAULT_URL must be a valid Azure Key Vault HTTPS URL\")\n+\n+    return vault_url\n+\n+\n+def run_crud_operations(vault_url: str) -\u003e None:\n+    \"\"\"Perform all CRUD operations on one secret.\"\"\"\n+    with DefaultAzureCredential() as credential:\n+        with SecretClient(vault_url=vault_url, credential=credential) as client:\n+            created = client.set_secret(SECRET_NAME, INITIAL_VALUE)\n+            print(\n+                f\"Created secret {created.name!r}, \"\n+                f\"version {created.properties.version!r}.\"\n+            )\n+\n+            retrieved = client.get_secret(SECRET_NAME)\n+            if retrieved.value != INITIAL_VALUE:\n+                raise RuntimeError(\"The retrieved secret does not match the created value\")\n+            print(f\"Read secret {retrieved.name!r}; its value matched the created value.\")\n+\n+            updated = client.set_secret(SECRET_NAME, UPDATED_VALUE)\n+            print(\n+                f\"Updated secret {updated.name!r}, \"\n+                f\"version {updated.properties.version!r}.\"\n+            )\n+\n+            retrieved_updated = client.get_secret(SECRET_NAME)\n+            if retrieved_updated.value != UPDATED_VALUE:\n+                raise RuntimeError(\"The retrieved secret does not match the updated value\")\n+            print(\"Read the updated secret; its value matched 'updated-value'.\")\n+\n+            client.begin_delete_secret(SECRET_NAME).result()\n+            print(f\"Deleted secret {SECRET_NAME!r} (soft delete).\")\n+\n+            client.purge_deleted_secret(SECRET_NAME)\n+            print(f\"Purged secret {SECRET_NAME!r}.\")\n+\n+\n+def main() -\u003e int:\n+    \"\"\"Run the example and translate expected failures into useful messages.\"\"\"\n+    try:\n+        run_crud_operations(get_vault_url())\n+    except ValueError as error:\n+        print(f\"Configuration error: {error}\", file=sys.stderr)\n+    except (CredentialUnavailableError, ClientAuthenticationError) as error:\n+        print(\n+            \"Authentication failed. Sign in with a supported developer credential \"\n+            \"(for example, Azure CLI) or configure a managed identity. \"\n+            f\"Details: {error}\",\n+            file=sys.stderr,\n+        )\n+    except ResourceNotFoundError as error:\n+        print(\n+            \"The vault or secret was not found. Check AZURE_KEYVAULT_URL and the \"\n+            f\"secret lifecycle state. Details: {error}\",\n+            file=sys.stderr,\n+        )\n+    except HttpResponseError as error:\n+        if error.status_code == 403:\n+            message = (\n+                \"Access denied. Grant secret get, set, delete, and purge permissions \"\n+                \"through Key Vault RBAC or an access policy.\"\n+            )\n+        elif error.status_code == 409:\n+            message = (\n+                \"The operation conflicted with the secret's current state. A \"\n+                \"previously deleted secret with the same name may still be retained.\"\n+            )\n+        else:\n+            message = \"Azure Key Vault request failed.\"\n+        print(f\"{message} Details: {error}\", file=sys.stderr)\n+    except RuntimeError as error:\n+        print(f\"Secret verification failed: {error}\", file=sys.stderr)\n+    else:\n+        return 0\n+\n+    return 1\n+\n+\n+if __name__ == \"__main__\":\n+    sys.exit(main())\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0\n+azure-keyvault-secrets\u003e=4.8.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2465558328/key_vault_secret_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2465558328/key_vault_secret_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2465558328/key_vault_secret_crud.py
@@ -1,0 +1,111 @@
+"""Create, read, update, delete, and purge an Azure Key Vault secret."""
+
+import os
+import sys
+
+from azure.core.exceptions import (
+    ClientAuthenticationError,
+    HttpResponseError,
+    ResourceNotFoundError,
+)
+from azure.identity import CredentialUnavailableError, DefaultAzureCredential
+from azure.keyvault.secrets import SecretClient
+
+
+SECRET_NAME = os.getenv("AZURE_KEYVAULT_SECRET_NAME", "crud-demo-secret")
+INITIAL_VALUE = "initial-value"
+UPDATED_VALUE = "updated-value"
+
+
+def get_vault_url() -> str:
+    """Return and validate the Key Vault URL from the environment."""
+    vault_url = os.getenv("AZURE_KEYVAULT_URL")
+    if not vault_url:
+        raise ValueError(
+            "AZURE_KEYVAULT_URL is required, for example "
+            "https://<vault-name>.vault.azure.net/"
+        )
+
+    if not vault_url.startswith("https://") or ".vault.azure.net" not in vault_url:
+        raise ValueError("AZURE_KEYVAULT_URL must be a valid Azure Key Vault HTTPS URL")
+
+    return vault_url
+
+
+def run_crud_operations(vault_url: str) -> None:
+    """Perform all CRUD operations on one secret."""
+    with DefaultAzureCredential() as credential:
+        with SecretClient(vault_url=vault_url, credential=credential) as client:
+            created = client.set_secret(SECRET_NAME, INITIAL_VALUE)
+            print(
+                f"Created secret {created.name!r}, "
+                f"version 
... (truncated)
```

### ✅ 8. powershell — 3465ms

**Input:**

```json
{"command":"python -m py_compile .\\key_vault_secret_crud.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $requirements = Get-Content .\\requirements.txt; if ($requirements -notcontains 'azure-identity\u003e=1.17.0' -or $requirements -notcontains 'azure-keyvault-secrets\u003e=4.8.0') { throw 'Required packages are missing from requirements.txt' }; Write-Output 'Local validation passed.'","description":"Validate Python syntax and package manifest","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Local validation passed.
<shellId: 0 completed with exit code 0>
```

### ✅ 9. powershell — 5907ms

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -LiteralPath .\\__pycache__ -Recurse -Force }","description":"Remove generated Python cache","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `key_vault_secret_crud.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the SDK guidance and workspace, then create a locally configurable script plus dependency file.

The workspace is empty. I’ll add a single executable CRUD script and `requirements.txt`, with secret values kept out of logs and explicit authentication/RBAC errors.

The implementation is in place. I’m doing a local-only syntax and dependency-manifest check; no Azure command or live operation will be run.

Created:

- `key_vault_secret_crud.py` — creates, reads, updates to `"updated-value"`, deletes, and purges a Key Vault secret using `DefaultAzureCredential`.
- `requirements.txt` — includes `azure-identity` and `azure-keyvault-secrets`.

The script reads `AZURE_KEYVAULT_URL` and optionally `AZURE_KEYVAULT_SECRET_NAME`, avoids logging secret values, and handles configuration, authentication, authorization, not-found, and conflict errors. Install dependencies with:

`pip install -r requirements.txt`

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Azure SDK exceptions are handled explicitly and ordered correctly, with specific messages for authentication, not-found, authorization, and conflict failures.

## Grader Results

- crud-secrets.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Installing `azure-keyvault-secrets` and `azure-identity` packages: Pass
      - Creating a `SecretClient` with vault URL and credential: Pass
      - `set_secret()`, `get_secret()`, `begin_delete_secret()`, `purge_deleted_secret()`: Pass
      - Handling soft-delete (waiting for delete to complete before purge): Pass
      - Exception handling for `ResourceNotFoundError`: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Pass
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Pass
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 7.0000** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-python-crud --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
