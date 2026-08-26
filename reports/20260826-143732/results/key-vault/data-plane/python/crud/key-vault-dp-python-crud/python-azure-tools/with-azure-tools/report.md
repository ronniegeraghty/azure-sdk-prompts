# Evaluation Report: key-vault-dp-python-crud

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 184.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-python-crud` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 10/12 |
| Duration | 184.0s |
| Timestamp | 2026-08-26T08:03:40Z |
| Files Generated | 2 |
| Event Count | 4105 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 104.1s |
| Review | 78.8s |
| **Total** | **184.0s** |

## Configuration

- **name:** python-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-keyvault-py, azure-identity-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=430261 out=3318 |
| Turn Count | 10 |

## Prompt Sent

```
Write a Python script that performs all four CRUD operations on Azure Key Vault secrets.

**Write the code to files (use file-write tools, do not reply with code blocks).**

3. Update the secret to a new value "updated-value"
4. Delete the secret and purge it (soft-delete enabled vault)

Use DefaultAzureCredential for authentication. Include proper error handling
and show required pip packages.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-keyvault-py"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-identity-py"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local Python script using DefaultAzureCredential and SecretClient to create, read, update, soft-delete, and purge an Azure Key Vault secret, with robust error handling and dependency documentation.","learn":true}
```

### 🔧 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation guidance for a Python Azure Key Vault secrets CRUD script.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 5. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-keyvault-py\\references\\capabilities.md"}
```

### 🔧 7. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\capabilities.md"}
```

### 🔧 8. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: key_vault_secret_crud.py\n+\"\"\"Create, read, update, delete, and purge an Azure Key Vault secret.\n+\n+Install dependencies with:\n+    python -m pip install -r requirements.txt\n+\n+Set AZURE_KEY_VAULT_URL to the vault URL, authenticate with any identity\n+supported by DefaultAzureCredential, and run:\n+    python key_vault_secret_crud.py\n+\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import logging\n+import os\n+import re\n+from collections.abc import Callable\n+from typing import TypeVar\n+from urllib.parse import urlparse\n+\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceNotFoundError,\n+)\n+from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n+from azure.keyvault.secrets import SecretClient\n+\n+\n+INITIAL_VALUE = \"initial-value\"\n+UPDATED_VALUE = \"updated-value\"\n+SECRET_NAME_PATTERN = re.compile(r\"^[0-9A-Za-z-]{1,127}$\")\n+T = TypeVar(\"T\")\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+class SecretOperationError(RuntimeError):\n+    \"\"\"An Azure Key Vault secret operation failed.\"\"\"\n+\n+\n+def execute(operation_name: str, operation: Callable[[], T]) -\u003e T:\n+    \"\"\"Run one Key Vault operation and add actionable error context.\"\"\"\n+    try:\n+        return operation()\n+    except (CredentialUnavailableError, ClientAuthenticationError) as error:\n+        raise SecretOperationError(\n+            f\"{operation_name} failed: DefaultAzureCredential could not authenticate. \"\n+            \"Configure a supported local credential or managed identity.\"\n+        ) from error\n+    except ResourceNotFoundError as error:\n+        raise SecretOperationError(\n+            f\"{operation_name} failed: the secret or vault was not found.\"\n+        ) from error\n+    except HttpResponseError as error:\n+        if error.status_code == 403:\n+            detail = (\n+                \"access denied; grant the identity secret get, set, delete, and purge \"\n+                \"permissions (for example, the Key Vault Secrets Officer RBAC role)\"\n+            )\n+        elif error.status_code == 409:\n+            detail = (\n+                \"the request conflicts with the vault state; a secret with this name \"\n+                \"may already be soft-deleted\"\n+            )\n+        else:\n+            detail = f\"Azure returned HTTP status {error.status_code or 'unknown'}\"\n+        raise SecretOperationError(f\"{operation_name} failed: {detail}.\") from error\n+\n+\n+def validate_vault_url(vault_url: str) -\u003e str:\n+    \"\"\"Validate and normalize the configured HTTPS vault URL.\"\"\"\n+    parsed = urlparse(vault_url)\n+    if parsed.scheme != \"https\" or not parsed.netloc or parsed.path not in (\"\", \"/\"):\n+        raise ValueError(\n+            \"AZURE_KEY_VAULT_URL must be an HTTPS vault URL such as \"\n+            \"https://\u003cvault-name\u003e.vault.azure.net/\"\n+        )\n+    return vault_url.rstrip(\"/\") + \"/\"\n+\n+\n+def perform_crud(vault_url: str, secret_name: str) -\u003e None:\n+    \"\"\"Perform the complete secret lifecycle, including permanent purge.\"\"\"\n+    with DefaultAzureCredential() as credential:\n+        with SecretClient(vault_url=vault_url, credential=credential) as client:\n+            created = execute(\n+                \"Create secret\",\n+                lambda: client.set_secret(secret_name, INITIAL_VALUE),\n+            )\n+            logger.info(\n+                \"Created secret %s (version %s).\",\n+                created.name,\n+                created.properties.version,\n+            )\n+\n+            retrieved = execute(\n+                \"Read secret\",\n+                lambda: client.get_secret(secret_name),\n+            )\n+            if retrieved.value != INITIAL_VALUE:\n+                raise SecretOperationError(\n+                    \"Read secret failed: the retrieved value did not match the created value.\"\n+                )\n+            logger.info(\"Read and verified secret %s.\", retrieved.name)\n+\n+            updated = execute(\n+                \"Update secret\",\n+                lambda: client.set_secret(secret_name, UPDATED_VALUE),\n+            )\n+            logger.info(\n+                \"Updated secret %s to a new version (%s).\",\n+                updated.name,\n+                updated.properties.version,\n+            )\n+\n+            updated_secret = execute(\n+                \"Verify updated secret\",\n+                lambda: client.get_secret(secret_name),\n+            )\n+            if updated_secret.value != UPDATED_VALUE:\n+                raise SecretOperationError(\n+                    \"Update secret failed: the retrieved value was not 'updated-value'.\"\n+                )\n+            logger.info(\"Verified the updated value for secret %s.\", secret_name)\n+\n+            execute(\n+                \"Delete secret\",\n+                lambda: client.begin_delete_secret(secret_name).result(),\n+            )\n+            logger.info(\"Soft-deleted secret %s.\", secret_name)\n+\n+            execute(\n+                \"Purge secret\",\n+                lambda: client.purge_deleted_secret(secret_name),\n+            )\n+            logger.info(\"Permanently purged secret %s.\", secret_name)\n+\n+\n+def parse_args() -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(\n+        description=\"Run all CRUD operations on an Azure Key Vault secret.\"\n+    )\n+    parser.add_argument(\n+        \"--vault-url\",\n+        default=os.getenv(\"AZURE_KEY_VAULT_URL\"),\n+        help=\"Key Vault URL; defaults to AZURE_KEY_VAULT_URL.\",\n+    )\n+    parser.add_argument(\n+        \"--secret-name\",\n+        default=\"crud-demo-secret\",\n+        help=\"Secret name to use (default: crud-demo-secret).\",\n+    )\n+    return parser.parse_args()\n+\n+\n+def main() -\u003e int:\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+    args = parse_args()\n+\n+    try:\n+        if not args.vault_url:\n+            raise ValueError(\n+                \"Set AZURE_KEY_VAULT_URL or provide the --vault-url argument.\"\n+            )\n+        vault_url = validate_vault_url(args.vault_url)\n+        if not SECRET_NAME_PATTERN.fullmatch(args.secret_name):\n+            raise ValueError(\n+                \"The secret name must contain 1-127 letters, numbers, or hyphens.\"\n+            )\n+        perform_crud(vault_url, args.secret_name)\n+    except (ValueError, SecretOperationError) as error:\n+        logger.error(\"%s\", error)\n+        return 1\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-identity\n+azure-keyvault-secrets\n*** End Patch\n"
```

### 🔧 9. powershell

**Input:**

```json
{"command":"python -m py_compile .\\key_vault_secret_crud.py","description":"Compile-check the Python script","initial_wait":30,"mode":"sync"}
```

### 🔧 10. glob

**Input:**

```json
{"pattern":"__pycache__/*"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"Remove-Item -LiteralPath '.\\__pycache__\\key_vault_secret_crud.cpython-313.pyc'; Remove-Item -LiteralPath '.\\__pycache__'","description":"Remove Python compile artifacts","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `key_vault_secret_crud.py`
- `requirements.txt`

## Copilot Response

I’m checking the workspace shape and SDK guidance, then I’ll add the script and dependency file with local-only validation.

The workspace is empty, so I’m creating a standalone CLI script plus `requirements.txt`; secret values will be verified but not printed to avoid leaking them into logs.

Created:

- `key_vault_secret_crud.py` — creates, reads, updates to `"updated-value"`, soft-deletes, and purges a secret using `DefaultAzureCredential`. Includes configuration validation, contextual Azure error handling, and avoids logging secret values.
- `requirements.txt` — lists `azure-identity` and `azure-keyvault-secrets`.

Run with:

```powershell
python -m pip install -r requirements.txt
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
python .\key_vault_secret_crud.py
```

The authenticated identity needs secret `get`, `set`, `delete`, and `purge` permissions, such as the **Key Vault Secrets Officer** role. References: [Key Vault Secrets Python client](https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme) and [Azure Identity for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. The script provides appropriate Azure SDK exception handling for every CRUD and purge operation.

### Strengths

- Catches specific authentication and not-found exceptions before the broader HttpResponseError handler.
- Provides actionable handling for common HTTP 403 and 409 responses.
- Preserves exception chains and converts failures into a controlled nonzero exit.

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 6.0000** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-python-crud --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
