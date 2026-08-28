# Evaluation Report: identity-dp-python-credential-chain

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 642.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-python-credential-chain` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 19/21 |
| Duration | 642.3s |
| Timestamp | 2026-08-27T20:18:03Z |
| Files Generated | 6 |
| Event Count | 11041 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 530.7s |
| Review | 110.5s |
| **Total** | **642.3s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=499426 out=8269 |
| Turn Count | 13 |

## Tool Calls

### ✅ 1. skill — 4412ms

**Input:**

```json
{"skill":"azure-identity-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for Python

Authentication library for Azure SDK clients using Microsoft Entra ID.

Use this skill when:
- An app needs to authenticate to Azure services from Python
- You need `DefaultAzureCredential` for local dev + Azure deployment
- You need `ManagedIdentityCredential` for Azure-hosted workloads
- You need service principal auth with secret or certificate
- You need direct token acquisition with `get_token()`
- You need to troubleshoot credential chain failures

## Installation

```bash
pip install azure-identity
```

For VS Code or broker-based desktop auth:

```bash
pip install azure-identity-broker
```

## Python Version

`azure-identity` supports Python 3.9+.

## Environment Variables

```bash
# Service principal with client secret
AZURE_TENANT_ID=<your-tenant-id>
AZURE_CLIENT_ID=<your-client-id>
AZURE_CLIENT_SECRET=<your-client-secret>

# Service principal with certificate
AZURE_TENANT_ID=<your-tenant-id>
AZURE_CLIENT_ID=<your-client-id>
AZURE_CLIENT_CERTIFICATE_PATH=/path/to/cert.pem
AZURE_CLIENT_CERTIFICATE_PASSWORD=<optional-password>

# Authority (sovereign clouds)
AZURE_AUTHORITY_HOST=login.microsoftonline.com  # Default; or login.chinacloudapi.cn, login.microsoftonline.us

# User-assigned managed identity
AZURE_CLIENT_ID=<managed-identity-client-id>

# Credential selection (new)
AZURE_TOKEN_CREDENTIALS=dev|prod|<credential-name>  # Optional, restricts DAC chain
```

## Authentication & Lifecycle

> **🔑 Two rules apply to every code sample below:**
>
> 1. **Prefer `DefaultAzureCredential`.** It works locally (Azure CLI / VS Code / Developer CLI) and in Azure (managed identity, workload identity) with no code change. Avoid connection strings, account/API keys — they bypass Entra audit and rotation.
>    - Local dev: `DefaultAzureCredential` works as-is.
>    - Production: set `AZURE_TOKEN_CREDENTIALS=prod` (or `AZURE_TOKEN_CREDENTIALS=<specific_credential>`) to constrain the credential chain t
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 180802ms

**Input:**

```json
{"intent":"Generate a local-only Python sample project demonstrating secure Azure Identity credential chains for development, CI/CD, and production, including sync/async token acquisition and CAE.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. view — 188746ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\capabilities.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/capabilities.md
@@ -1,43 +1,43 @@
 # azure-identity-py capability coverage
 
 **SDK/package**: `azure-identity`
 
 This index maps hero scenarios in `SKILL.md` and links non-hero scenarios documented in dedicated reference files.
 
 ## Hero scenarios covered in SKILL.md
 
 - `get_bearer_token_provider`
 - `Credential Types`
 - `Specific Credential Examples`
 - `Getting Tokens Directly`
 
 ## Non-hero scenarios
 
 - `Async Credentials`: Async credentials are in `azure.identity.aio`. Always close them or use `async with`:  
   See: [`non-hero-scenarios.md#async-credentials`](non-hero-scenarios.md#async-credentials)
 - `Sovereign Clouds`: Use `AzureAuthorityHosts` or the `AZURE_AUTHORITY_HOST` env var:  
   See: [`non-hero-scenarios.md#sovereign-clouds`](non-hero-scenarios.md#sovereign-clouds)
 - `Persistent Token Caching`: Opt-in disk-based caching with `TokenCachePersistenceOptions`:  
   See: [`non-hero-scenarios.md#persistent-token-caching`](non-hero-scenarios.md#persistent-token-caching)
 - `Multi-Tenant Support`: Allow token acquisition for additional tenants beyond the configured one:  
   See: [`non-hero-scenarios.md#multi-tenant-support`](non-hero-scenarios.md#multi-tenant-support)
 - `Error Handling`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#error-handling`](non-hero-scenarios.md#error-handling)
 - `Logging`: Ena
... (truncated)
```

### ✅ 4. view — 187392ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\non-hero-scenarios.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/non-hero-scenarios.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/non-hero-scenarios.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/non-hero-scenarios.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-identity-py/references/non-hero-scenarios.md
@@ -1,136 +1,136 @@
 # azure-identity-py non-hero scenarios
 
 These scenarios are intentionally separate from hero flows in `SKILL.md`.
 They cover secondary/advanced patterns typically used after the primary end-to-end path is working.
 
 ## Async Credentials
 
 Async credentials are in `azure.identity.aio`. Always close them or use `async with`:
 
 ```python
 from azure.identity.aio import DefaultAzureCredential
 from azure.storage.blob.aio import BlobServiceClient
 
 async def main():
     # Preferred: use async context manager for both credential and client
     async with DefaultAzureCredential() as credential:
         async with BlobServiceClient(
             account_url="https://<account>.blob.core.windows.net",
             credential=credential,
         ) as client:
             # ... async operations
             pass
 ```
 
 > The async `get_bearer_token_provider` is at `azure.identity.aio.get_bearer_token_provider`.
 
 ## Sovereign Clouds
 
 Use `AzureAuthorityHosts` or the `AZURE_AUTHORITY_HOST` env var:
 
 ```python
 from azure.identity import DefaultAzureCredential, AzureAuthorityHosts
 
 # Azure Government
 credential = DefaultAzureCredential(authority=AzureAuthorityHosts.AZURE_GOVERNMENT)
 
 # Azure China
 credential = DefaultAzureCredential(authority=AzureAuthorityHosts.AZURE_CHINA)
 ```
 
 | Co
... (truncated)
```

### ✅ 5. glob — 186128ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 188747ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Get current Azure best practices for local-only Python code using azure-identity credential chains, managed identity, workload identity, Azure Pipelines federation, sync/async token requests, and CAE.","parameters":{}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. apply_patch — 3527ms

**Input:**

```json
"*** Begin Patch\n*** Add File: credential_factory.py\n+\"\"\"Build Azure credential chains tailored to the detected runtime environment.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from dataclasses import dataclass\n+from typing import Mapping\n+\n+from azure.core.credentials import TokenCredential\n+from azure.core.credentials_async import AsyncTokenCredential\n+from azure.identity import (\n+    AzureCliCredential,\n+    AzureDeveloperCliCredential,\n+    AzurePipelinesCredential,\n+    AzurePowerShellCredential,\n+    ChainedTokenCredential,\n+    EnvironmentCredential,\n+    ManagedIdentityCredential,\n+    VisualStudioCodeCredential,\n+    WorkloadIdentityCredential,\n+)\n+from azure.identity.aio import (\n+    AzureCliCredential as AsyncAzureCliCredential,\n+    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,\n+    AzurePipelinesCredential as AsyncAzurePipelinesCredential,\n+    AzurePowerShellCredential as AsyncAzurePowerShellCredential,\n+    ChainedTokenCredential as AsyncChainedTokenCredential,\n+    EnvironmentCredential as AsyncEnvironmentCredential,\n+    ManagedIdentityCredential as AsyncManagedIdentityCredential,\n+    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,\n+    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,\n+)\n+\n+from environment_detector import RuntimeEnvironment\n+\n+\n+@dataclass(frozen=True)\n+class SyncCredentialSelection:\n+    credential: TokenCredential\n+    strategy: str\n+    enable_cae: bool\n+\n+\n+@dataclass(frozen=True)\n+class AsyncCredentialSelection:\n+    credential: AsyncTokenCredential\n+    strategy: str\n+    enable_cae: bool\n+\n+\n+def build_sync_credential(\n+    environment: RuntimeEnvironment,\n+    *,\n+    enable_cae: bool = False,\n+    environ: Mapping[str, str] | None = None,\n+) -\u003e SyncCredentialSelection:\n+    \"\"\"Create a synchronous credential chain for one runtime environment.\"\"\"\n+    values = os.environ if environ is None else environ\n+\n+    if environment is RuntimeEnvironment.DEV:\n+        credential = ChainedTokenCredential(\n+            VisualStudioCodeCredential(),\n+            AzureCliCredential(),\n+            AzurePowerShellCredential(),\n+            AzureDeveloperCliCredential(),\n+        )\n+        strategy = \"developer tools: VS Code -\u003e Azure CLI -\u003e Azure PowerShell -\u003e Azure Developer CLI\"\n+    elif environment is RuntimeEnvironment.CI:\n+        credential, strategy = _build_sync_ci_credential(values)\n+    else:\n+        credential, strategy = _build_sync_production_credential(values)\n+\n+    return SyncCredentialSelection(credential, strategy, enable_cae)\n+\n+\n+def build_async_credential(\n+    environment: RuntimeEnvironment,\n+    *,\n+    enable_cae: bool = False,\n+    environ: Mapping[str, str] | None = None,\n+) -\u003e AsyncCredentialSelection:\n+    \"\"\"Create an asynchronous credential chain for one runtime environment.\"\"\"\n+    values = os.environ if environ is None else environ\n+\n+    if environment is RuntimeEnvironment.DEV:\n+        credential = AsyncChainedTokenCredential(\n+            AsyncVisualStudioCodeCredential(),\n+            AsyncAzureCliCredential(),\n+            AsyncAzurePowerShellCredential(),\n+            AsyncAzureDeveloperCliCredential(),\n+        )\n+        strategy = \"developer tools: VS Code -\u003e Azure CLI -\u003e Azure PowerShell -\u003e Azure Developer CLI\"\n+    elif environment is RuntimeEnvironment.CI:\n+        credential, strategy = _build_async_ci_credential(values)\n+    else:\n+        credential, strategy = _build_async_production_credential(values)\n+\n+    return AsyncCredentialSelection(credential, strategy, enable_cae)\n+\n+\n+def _pipeline_settings(values: Mapping[str, str]) -\u003e dict[str, str] | None:\n+    names = {\n+        \"tenant_id\": \"AZURE_TENANT_ID\",\n+        \"client_id\": \"AZURE_CLIENT_ID\",\n+        \"service_connection_id\": \"AZURE_PIPELINES_SERVICE_CONNECTION_ID\",\n+        \"system_access_token\": \"SYSTEM_ACCESSTOKEN\",\n+    }\n+    present = {key: values.get(name, \"\").strip() for key, name in names.items()}\n+    pipeline_specific = (\n+        present[\"service_connection_id\"],\n+        present[\"system_access_token\"],\n+    )\n+    if not any(pipeline_specific):\n+        return None\n+\n+    missing = [names[key] for key, value in present.items() if not value]\n+    if missing:\n+        raise ValueError(\n+            \"Azure Pipelines workload identity configuration is incomplete; \"\n+            f\"missing: {', '.join(missing)}\"\n+        )\n+    return present\n+\n+\n+def _build_sync_ci_credential(\n+    values: Mapping[str, str],\n+) -\u003e tuple[TokenCredential, str]:\n+    settings = _pipeline_settings(values)\n+    if settings:\n+        return (\n+            ChainedTokenCredential(\n+                AzurePipelinesCredential(**settings),\n+                EnvironmentCredential(),\n+            ),\n+            \"CI: Azure Pipelines workload identity -\u003e environment credential\",\n+        )\n+    return EnvironmentCredential(), \"CI: environment credential (secret or certificate)\"\n+\n+\n+def _build_async_ci_credential(\n+    values: Mapping[str, str],\n+) -\u003e tuple[AsyncTokenCredential, str]:\n+    settings = _pipeline_settings(values)\n+    if settings:\n+        return (\n+            AsyncChainedTokenCredential(\n+                AsyncAzurePipelinesCredential(**settings),\n+                AsyncEnvironmentCredential(),\n+            ),\n+            \"CI: Azure Pipelines workload identity -\u003e environment credential\",\n+        )\n+    return (\n+        AsyncEnvironmentCredential(),\n+        \"CI: environment credential (secret or certificate)\",\n+    )\n+\n+\n+def _workload_identity_is_configured(values: Mapping[str, str]) -\u003e bool:\n+    required = (\n+        \"AZURE_TENANT_ID\",\n+        \"AZURE_CLIENT_ID\",\n+        \"AZURE_FEDERATED_TOKEN_FILE\",\n+    )\n+    return all(values.get(name, \"\").strip() for name in required)\n+\n+\n+def _managed_identity_client_id(values: Mapping[str, str]) -\u003e str | None:\n+    return (\n+        values.get(\"AZURE_MANAGED_IDENTITY_CLIENT_ID\", \"\").strip()\n+        or values.get(\"AZURE_CLIENT_ID\", \"\").strip()\n+        or None\n+    )\n+\n+\n+def _build_sync_production_credential(\n+    values: Mapping[str, str],\n+) -\u003e tuple[TokenCredential, str]:\n+    client_id = _managed_identity_client_id(values)\n+    credentials: list[TokenCredential] = [\n+        ManagedIdentityCredential(client_id=client_id)\n+    ]\n+    strategy = (\n+        \"production: user-assigned managed identity\"\n+        if client_id\n+        else \"production: system-assigned managed identity\"\n+    )\n+    if _workload_identity_is_configured(values):\n+        credentials.append(\n+            WorkloadIdentityCredential(\n+                tenant_id=values[\"AZURE_TENANT_ID\"],\n+                client_id=values[\"AZURE_CLIENT_ID\"],\n+                token_file_path=values[\"AZURE_FEDERATED_TOKEN_FILE\"],\n+            )\n+        )\n+        strategy += \" -\u003e Kubernetes workload identity\"\n+    return ChainedTokenCredential(*credentials), strategy\n+\n+\n+def _build_async_production_credential(\n+    values: Mapping[str, str],\n+) -\u003e tuple[AsyncTokenCredential, str]:\n+    client_id = _managed_identity_client_id(values)\n+    credentials: list[AsyncTokenCredential] = [\n+        AsyncManagedIdentityCredential(client_id=client_id)\n+    ]\n+    strategy = (\n+        \"production: user-assigned managed identity\"\n+        if client_id\n+        else \"production: system-assigned managed identity\"\n+    )\n+    if _workload_identity_is_configured(values):\n+        credentials.append(\n+            AsyncWorkloadIdentityCredential(\n+                tenant_id=values[\"AZURE_TENANT_ID\"],\n+                client_id=values[\"AZURE_CLIENT_ID\"],\n+                token_file_path=values[\"AZURE_FEDERATED_TOKEN_FILE\"],\n+            )\n+        )\n+        strategy += \" -\u003e Kubernetes workload identity\"\n+    return AsyncChainedTokenCredential(*credentials), strategy\n*** Add File: environment_detector.py\n+\"\"\"Detect whether the process is running locally, in CI, or in production.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from enum import Enum\n+from typing import Mapping\n+\n+\n+class RuntimeEnvironment(str, Enum):\n+    DEV = \"dev\"\n+    CI = \"ci\"\n+    PRODUCTION = \"production\"\n+\n+\n+CI_MARKERS = (\n+    \"CI\",\n+    \"GITHUB_ACTIONS\",\n+    \"GITLAB_CI\",\n+    \"TF_BUILD\",\n+    \"BUILD_BUILDID\",\n+    \"PIPELINE_WORKSPACE\",\n+    \"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI\",\n+    \"JENKINS_URL\",\n+)\n+\n+PRODUCTION_MARKERS = (\n+    \"IDENTITY_ENDPOINT\",\n+    \"MSI_ENDPOINT\",\n+    \"IMDS_ENDPOINT\",\n+    \"IDENTITY_HEADER\",\n+    \"MSI_SECRET\",\n+    \"WEBSITE_INSTANCE_ID\",\n+    \"CONTAINER_APP_NAME\",\n+    \"KUBERNETES_SERVICE_HOST\",\n+    \"AZURE_FEDERATED_TOKEN_FILE\",\n+)\n+\n+\n+def detect_environment(\n+    environ: Mapping[str, str] | None = None,\n+) -\u003e RuntimeEnvironment:\n+    \"\"\"Classify the runtime using environment markers, with CI taking precedence.\"\"\"\n+    values = os.environ if environ is None else environ\n+\n+    if _has_marker(values, CI_MARKERS):\n+        return RuntimeEnvironment.CI\n+    if _has_marker(values, PRODUCTION_MARKERS):\n+        return RuntimeEnvironment.PRODUCTION\n+    return RuntimeEnvironment.DEV\n+\n+\n+def _has_marker(values: Mapping[str, str], names: tuple[str, ...]) -\u003e bool:\n+    return any(values.get(name, \"\").strip() for name in names)\n*** Add File: connectivity_tester.py\n+\"\"\"Synchronous and asynchronous Azure credential connectivity tests.\"\"\"\n+\n+from __future__ import annotations\n+\n+from dataclasses import dataclass\n+from datetime import datetime, timezone\n+\n+from azure.core.credentials import TokenCredential\n+from azure.core.credentials_async import AsyncTokenCredential\n+from azure.core.exceptions import ClientAuthenticationError\n+from azure.identity import CredentialUnavailableError\n+\n+\n+@dataclass(frozen=True)\n+class ConnectivityResult:\n+    success: bool\n+    cae_requested: bool\n+    expires_at: datetime | None = None\n+    failure_reason: str | None = None\n+\n+\n+def test_credential(\n+    credential: TokenCredential,\n+    scope: str,\n+    *,\n+    enable_cae: bool = False,\n+) -\u003e ConnectivityResult:\n+    \"\"\"Request a token and print a detailed synchronous test result.\"\"\"\n+    print(\"[sync] Requesting token...\")\n+    try:\n+        token = credential.get_token(scope, enable_cae=enable_cae)\n+    except (CredentialUnavailableError, ClientAuthenticationError) as error:\n+        result = _failure_result(error, enable_cae)\n+    except Exception as error:\n+        result = ConnectivityResult(\n+            success=False,\n+            cae_requested=enable_cae,\n+            failure_reason=(\n+                f\"Unexpected {type(error).__name__}: {str(error).strip() or 'no details'}\"\n+            ),\n+        )\n+    else:\n+        result = ConnectivityResult(\n+            success=True,\n+            cae_requested=enable_cae,\n+            expires_at=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),\n+        )\n+    _print_result(\"sync\", result)\n+    return result\n+\n+\n+async def test_credential_async(\n+    credential: AsyncTokenCredential,\n+    scope: str,\n+    *,\n+    enable_cae: bool = False,\n+) -\u003e ConnectivityResult:\n+    \"\"\"Request a token and print a detailed asynchronous test result.\"\"\"\n+    print(\"[async] Requesting token...\")\n+    try:\n+        token = await credential.get_token(scope, enable_cae=enable_cae)\n+    except (CredentialUnavailableError, ClientAuthenticationError) as error:\n+        result = _failure_result(error, enable_cae)\n+    except Exception as error:\n+        result = ConnectivityResult(\n+            success=False,\n+            cae_requested=enable_cae,\n+            failure_reason=(\n+                f\"Unexpected {type(error).__name__}: {str(error).strip() or 'no details'}\"\n+            ),\n+        )\n+    else:\n+        result = ConnectivityResult(\n+            success=True,\n+            cae_requested=enable_cae,\n+            expires_at=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),\n+        )\n+    _print_result(\"async\", result)\n+    return result\n+\n+\n+def _failure_result(\n+    error: CredentialUnavailableError | ClientAuthenticationError,\n+    enable_cae: bool,\n+) -\u003e ConnectivityResult:\n+    details = str(error).strip()\n+    normalized = details.lower()\n+\n+    if isinstance(error, CredentialUnavailableError):\n+        category = \"No identity is available for this credential\"\n+    elif any(\n+        phrase in normalized\n+        for phrase in (\n+            \"certificate has expired\",\n+            \"expired certificate\",\n+            \"aadsts700027\",\n+        )\n+    ):\n+        category = \"The client certificate is expired or invalid\"\n+    elif any(\n+        phrase in normalized\n+        for phrase in (\n+            \"secret is expired\",\n+            \"client secret\",\n+            \"aadsts7000222\",\n+        )\n+    ):\n+        category = \"The client secret is expired or invalid\"\n+    elif any(\n+        phrase in normalized\n+        for phrase in (\n+            \"tenant not found\",\n+            \"invalid tenant\",\n+            \"aadsts90002\",\n+        )\n+    ):\n+        category = \"The configured tenant is invalid or cannot be found\"\n+    elif any(\n+        phrase in normalized\n+        for phrase in (\n+            \"managed identity\",\n+            \"identity not found\",\n+            \"no response from the imds endpoint\",\n+        )\n+    ):\n+        category = \"The requested managed identity is unavailable\"\n+    elif any(\n+        phrase in normalized\n+        for phrase in (\n+            \"federated identity credential\",\n+            \"federated token\",\n+            \"aadsts70021\",\n+            \"aadsts700212\",\n+        )\n+    ):\n+        category = \"The workload identity federation settings do not match\"\n+    elif \"unauthorized_client\" in normalized or \"aadsts700016\" in normalized:\n+        category = \"The client ID is invalid or not registered in this tenant\"\n+    elif \"interaction_required\" in normalized:\n+        category = \"User interaction or a fresh developer login is required\"\n+    else:\n+        category = \"Microsoft Entra ID rejected the authentication request\"\n+\n+    reason = f\"{category}. SDK details: {details or type(error).__name__}\"\n+    return ConnectivityResult(\n+        success=False,\n+        cae_requested=enable_cae,\n+        failure_reason=reason,\n+    )\n+\n+\n+def _print_result(label: str, result: ConnectivityResult) -\u003e None:\n+    status = \"SUCCESS\" if result.success else \"FAILURE\"\n+    print(f\"[{label}] Result: {status}\")\n+    print(f\"[{label}] CAE requested: {'yes' if result.cae_requested else 'no'}\")\n+    if result.expires_at:\n+        print(f\"[{label}] Token expires (UTC): {result.expires_at.isoformat()}\")\n+    if result.failure_reason:\n+        print(f\"[{label}] Failure reason: {result.failure_reason}\")\n*** Add File: main.py\n+\"\"\"Run sync and async Azure Resource Manager authentication checks.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import asyncio\n+import os\n+\n+from connectivity_tester import test_credential, test_credential_async\n+from credential_factory import build_async_credential, build_sync_credential\n+from environment_detector import detect_environment\n+\n+\n+ARM_SCOPE = \"https://management.azure.com/.default\"\n+\n+\n+def _cae_default() -\u003e bool:\n+    value = os.getenv(\"AZURE_ENABLE_CAE\", \"\").strip().lower()\n+    return value in {\"1\", \"true\", \"yes\", \"on\"}\n+\n+\n+def _parse_args() -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(\n+        description=\"Test environment-specific Azure credential chains.\"\n+    )\n+    parser.add_argument(\n+        \"--cae\",\n+        action=argparse.BooleanOptionalAction,\n+        default=_cae_default(),\n+        help=\"request a CAE-enabled token (default: AZURE_ENABLE_CAE)\",\n+    )\n+    return parser.parse_args()\n+\n+\n+async def _run_async(environment, enable_cae: bool) -\u003e None:\n+    selection = build_async_credential(environment, enable_cae=enable_cae)\n+    print(f\"Async credential strategy: {selection.strategy}\")\n+    async with selection.credential:\n+        await test_credential_async(\n+            selection.credential,\n+            ARM_SCOPE,\n+            enable_cae=selection.enable_cae,\n+        )\n+\n+\n+def main() -\u003e None:\n+    args = _parse_args()\n+    environment = detect_environment()\n+    print(f\"Detected environment: {environment.value}\")\n+\n+    selection = build_sync_credential(environment, enable_cae=args.cae)\n+    print(f\"Sync credential strategy: {selection.strategy}\")\n+    with selection.credential:\n+        test_credential(\n+            selection.credential,\n+            ARM_SCOPE,\n+            enable_cae=selection.enable_cae,\n+        )\n+\n+    asyncio.run(_run_async(environment, args.cae))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-identity-broker\u003e=1.1.0,\u003c2.0.0\n+aiohttp\u003e=3.9.0,\u003c4.0.0\n*** Add File: test_project.py\n+\"\"\"Offline unit tests for environment selection and authentication diagnostics.\"\"\"\n+\n+from __future__ import annotations\n+\n+import unittest\n+\n+from azure.core.exceptions import ClientAuthenticationError\n+from azure.identity import CredentialUnavailableError\n+\n+from connectivity_tester import _failure_result\n+from credential_factory import build_sync_credential\n+from environment_detector import RuntimeEnvironment, detect_environment\n+\n+\n+class EnvironmentDetectorTests(unittest.TestCase):\n+    def test_defaults_to_development(self) -\u003e None:\n+        self.assertEqual(detect_environment({}), RuntimeEnvironment.DEV)\n+\n+    def test_ci_takes_precedence_over_production(self) -\u003e None:\n+        values = {\"PIPELINE_WORKSPACE\": \"work\", \"IDENTITY_ENDPOINT\": \"endpoint\"}\n+        self.assertEqual(detect_environment(values), RuntimeEnvironment.CI)\n+\n+    def test_detects_workload_identity_as_production(self) -\u003e None:\n+        values = {\"AZURE_FEDERATED_TOKEN_FILE\": \"token\"}\n+        self.assertEqual(detect_environment(values), RuntimeEnvironment.PRODUCTION)\n+\n+\n+class CredentialFactoryTests(unittest.TestCase):\n+    def test_production_selects_user_assigned_identity(self) -\u003e None:\n+        selection = build_sync_credential(\n+            RuntimeEnvironment.PRODUCTION,\n+            environ={\"AZURE_MANAGED_IDENTITY_CLIENT_ID\": \"client-id\"},\n+        )\n+        self.assertIn(\"user-assigned managed identity\", selection.strategy)\n+        selection.credential.close()\n+\n+    def test_rejects_incomplete_pipeline_federation_settings(self) -\u003e None:\n+        with self.assertRaisesRegex(ValueError, \"AZURE_TENANT_ID\"):\n+            build_sync_credential(\n+                RuntimeEnvironment.CI,\n+                environ={\"SYSTEM_ACCESSTOKEN\": \"token\"},\n+            )\n+\n+\n+class AuthenticationDiagnosticTests(unittest.TestCase):\n+    def test_reports_expired_certificate(self) -\u003e None:\n+        error = ClientAuthenticationError(\n+            message=\"AADSTS700027: The certificate has expired.\"\n+        )\n+        result = _failure_result(error, enable_cae=True)\n+        self.assertIn(\"certificate is expired\", result.failure_reason or \"\")\n+        self.assertTrue(result.cae_requested)\n+\n+    def test_reports_unavailable_identity(self) -\u003e None:\n+        result = _failure_result(\n+            CredentialUnavailableError(\"IMDS endpoint unavailable\"),\n+            enable_cae=False,\n+        )\n+        self.assertIn(\"No identity is available\", result.failure_reason or \"\")\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py
@@ -1,0 +1,221 @@
+"""Build Azure credential chains tailored to the detected runtime environment."""
+
+from __future__ import annotations
+
+import os
+from dataclasses import dataclass
+from typing import Mapping
+
+from azure.core.credentials import TokenCredential
+from azure.core.credentials_async import AsyncTokenCredential
+from azure.identity import (
+    AzureCliCredential,
+    AzureDeveloperCliCredential,
+    AzurePipelinesCredential,
+    AzurePowerShellCredential,
+    ChainedTokenCredential,
+    EnvironmentCredential,
+    ManagedIdentityCredential,
+    VisualStudioCodeCredential,
+    WorkloadIdentityCredential,
+)
+from azure.identity.aio import (
+    AzureCliCredential as AsyncAzureCliCredential,
+    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
+    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
+    AzurePowerShellCredential as AsyncAzurePowerShellCredential,
+    ChainedTokenCredential as AsyncChainedTokenCredential,
+    EnvironmentCredential as AsyncEnvironmentCredential,
+    ManagedIdentityCredential as AsyncManagedIdentityCredential,
+    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
+    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,
+)
+
+from environment_detector import RuntimeEnvironment
+
+
+@dataclass(frozen=True)
+class SyncCredentialSelection:
+    credential: TokenCredential
+    strategy: str
+    enable
... (truncated)
```

### ✅ 8. task — 25054ms

**Input:**

```json
{"agent_type":"task","description":"Run offline Python tests","mode":"sync","name":"python-validation","prompt":"In the current working directory, run the smallest local-only validation for the newly created Python Azure Identity project. First run `python -m py_compile credential_factory.py environment_detector.py connectivity_tester.py main.py test_project.py`, then run `python -m unittest -v test_project.py`. Do not contact Azure or run the main script. Return concise success status, or full relevant errors if validation fails."}
```

**Output:**

```
Validation succeeded: all 5 Python files compiled, and all 7 unit tests passed.
```

### ✅ 9. rg — 24680ms

**Input:**

```json
{"-i":true,"-n":true,"glob":"*.py","output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329","pattern":"(password|secret|token)\\s*=\\s*[\"'][^\"']+[\"']"}
```

**Output:**

```
No matches found.
```

### ✅ 10. powershell — 1553ms

**Input:**

```json
{"command":"python -m py_compile credential_factory.py environment_detector.py connectivity_tester.py main.py test_project.py","description":"Compile Python project files","initial_wait":60,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 11. powershell — 2871ms

**Input:**

```json
{"command":"python -m unittest -v test_project.py","description":"Run project unit tests","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_reports_expired_certificate (test_project.AuthenticationDiagnosticTests.test_reports_expired_certificate) ... ok
test_reports_unavailable_identity (test_project.AuthenticationDiagnosticTests.test_reports_unavailable_identity) ... ok
test_production_selects_user_assigned_identity (test_project.CredentialFactoryTests.test_production_selects_user_assigned_identity) ... ok
test_rejects_incomplete_pipeline_federation_settings (test_project.CredentialFactoryTests.test_rejects_incomplete_pipeline_federation_settings) ... ok
test_ci_takes_precedence_over_production (test_project.EnvironmentDetectorTests.test_ci_takes_precedence_over_production) ... ok
test_defaults_to_development (test_project.EnvironmentDetectorTests.test_defaults_to_development) ... ok
test_detects_workload_identity_as_production (test_project.EnvironmentDetectorTests.test_detects_workload_identity_as_production) ... ok

----------------------------------------------------------------------
Ran 7 tests in 0.234s

OK
<shellId: 1 completed with exit code 0>
```

### ✅ 12. apply_patch — 694ms

**Input:**

```json
"*** Begin Patch\n*** Update File: credential_factory.py\n@@\n def _pipeline_settings(values: Mapping[str, str]) -\u003e dict[str, str] | None:\n-    names = {\n-        \"tenant_id\": \"AZURE_TENANT_ID\",\n-        \"client_id\": \"AZURE_CLIENT_ID\",\n-        \"service_connection_id\": \"AZURE_PIPELINES_SERVICE_CONNECTION_ID\",\n-        \"system_access_token\": \"SYSTEM_ACCESSTOKEN\",\n+    names = {\n+        \"tenant_id\": (\"AZURESUBSCRIPTION_TENANT_ID\", \"AZURE_TENANT_ID\"),\n+        \"client_id\": (\"AZURESUBSCRIPTION_CLIENT_ID\", \"AZURE_CLIENT_ID\"),\n+        \"service_connection_id\": (\n+            \"AZURESUBSCRIPTION_SERVICE_CONNECTION_ID\",\n+            \"AZURE_PIPELINES_SERVICE_CONNECTION_ID\",\n+        ),\n+        \"system_access_token\": (\"SYSTEM_ACCESSTOKEN\",),\n     }\n-    present = {key: values.get(name, \"\").strip() for key, name in names.items()}\n+    present = {\n+        key: next(\n+            (\n+                values.get(name, \"\").strip()\n+                for name in aliases\n+                if values.get(name, \"\").strip()\n+            ),\n+            \"\",\n+        )\n+        for key, aliases in names.items()\n+    }\n@@\n-    missing = [names[key] for key, value in present.items() if not value]\n+    missing = [\"/\".join(names[key]) for key, value in present.items() if not value]\n*** Update File: test_project.py\n@@\n     def test_rejects_incomplete_pipeline_federation_settings(self) -\u003e None:\n-        with self.assertRaisesRegex(ValueError, \"AZURE_TENANT_ID\"):\n+        with self.assertRaisesRegex(ValueError, \"AZURESUBSCRIPTION_TENANT_ID\"):\n             build_sync_credential(\n                 RuntimeEnvironment.CI,\n                 environ={\"SYSTEM_ACCESSTOKEN\": \"token\"},\n             )\n+\n+    def test_accepts_standard_azure_pipelines_variables(self) -\u003e None:\n+        selection = build_sync_credential(\n+            RuntimeEnvironment.CI,\n+            environ={\n+                \"AZURESUBSCRIPTION_TENANT_ID\": \"tenant-id\",\n+                \"AZURESUBSCRIPTION_CLIENT_ID\": \"client-id\",\n+                \"AZURESUBSCRIPTION_SERVICE_CONNECTION_ID\": \"connection-id\",\n+                \"SYSTEM_ACCESSTOKEN\": \"token\",\n+            },\n+        )\n+        self.assertIn(\"Azure Pipelines workload identity\", selection.strategy)\n+        selection.credential.close()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-microsoft-skill-1127595329/credential_factory.py
@@ -100,12 +100,25 @@
 
 def _pipeline_settings(values: Mapping[str, str]) -> dict[str, str] | None:
     names = {
-        "tenant_id": "AZURE_TENANT_ID",
-        "client_id": "AZURE_CLIENT_ID",
-        "service_connection_id": "AZURE_PIPELINES_SERVICE_CONNECTION_ID",
-        "system_access_token": "SYSTEM_ACCESSTOKEN",
+        "tenant_id": ("AZURESUBSCRIPTION_TENANT_ID", "AZURE_TENANT_ID"),
+        "client_id": ("AZURESUBSCRIPTION_CLIENT_ID", "AZURE_CLIENT_ID"),
+        "service_connection_id": (
+            "AZURESUBSCRIPTION_SERVICE_CONNECTION_ID",
+            "AZURE_PIPELINES_SERVICE_CONNECTION_ID",
+        ),
+        "system_access_token": ("SYSTEM_ACCESSTOKEN",),
     }
-    present = {key: values.get(name, "").strip() for key, name in names.items()}
+    present = {
+        key: next(
+            (
+                values.get(name, "").strip()
+                for name in aliases
+                if values.get(name, "").strip()
+            ),
+            "",
+        )
+        for key, aliases in names.items()
+    }
     pipeline_specific = (
         present["service_connection_id"],
         present["system_access_token"],
@@ -113,7 +126,7 @@
     if not any(pipeline_specific):
         return None
 
-    missing = [names[key] for key, value in pr
... (truncated)
```

### ✅ 13. task — 16053ms

**Input:**

```json
{"agent_type":"task","description":"Re-run project validation","mode":"sync","name":"final-python-validation","prompt":"In the current working directory, validate the final local-only Python project without contacting Azure. Run `python -m py_compile credential_factory.py environment_detector.py connectivity_tester.py main.py test_project.py`, `python -m unittest -v test_project.py`, and `python main.py --help`. Return concise success status, or full relevant errors on failure."}
```

**Output:**

```
Validation succeeded: all 5 files compiled, all 8 unit tests passed, and `python main.py --help` completed successfully.
```

### ✅ 14. powershell — 3777ms

**Input:**

```json
{"command":"$failed = $false\nWrite-Output '=== py_compile ==='\npython -m py_compile credential_factory.py environment_detector.py connectivity_tester.py main.py test_project.py\nif ($LASTEXITCODE -ne 0) { $failed = $true }\nWrite-Output '=== unittest ==='\npython -m unittest -v test_project.py\nif ($LASTEXITCODE -ne 0) { $failed = $true }\nWrite-Output '=== main_help ==='\npython main.py --help\nif ($LASTEXITCODE -ne 0) { $failed = $true }\nif ($failed) { exit 1 } else { exit 0 }","description":"Compile, test, and check CLI help","initial_wait":300,"mode":"sync"}
```

**Output:**

```
=== py_compile ===
=== unittest ===
test_reports_expired_certificate (test_project.AuthenticationDiagnosticTests.test_reports_expired_certificate) ... ok
test_reports_unavailable_identity (test_project.AuthenticationDiagnosticTests.test_reports_unavailable_identity) ... ok
test_accepts_standard_azure_pipelines_variables (test_project.CredentialFactoryTests.test_accepts_standard_azure_pipelines_variables) ... Incomplete environment configuration for EnvironmentCredential. These variables are set: AZURE_TENANT_ID
ok
test_production_selects_user_assigned_identity (test_project.CredentialFactoryTests.test_production_selects_user_assigned_identity) ... ok
test_rejects_incomplete_pipeline_federation_settings (test_project.CredentialFactoryTests.test_rejects_incomplete_pipeline_federation_settings) ... ok
test_ci_takes_precedence_over_production (test_project.EnvironmentDetectorTests.test_ci_takes_precedence_over_production) ... ok
test_defaults_to_development (test_project.EnvironmentDetectorTests.test_defaults_to_development) ... ok
test_detects_workload_identity_as_production (test_project.EnvironmentDetectorTests.test_detects_workload_identity_as_production) ... ok

----------------------------------------------------------------------
Ran 8 tests in 0.230s

OK
=== main_help ===
usage: main.py [-h] [--cae | --no-cae]

Test environment-specific Azure credential chains.

options:
  -h, --help       show this help message and exit
  --cae, --no-cae  request a CAE-enabled token (default: AZURE_ENABLE_CAE)
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `connectivity_tester.py`
- `credential_factory.py`
- `environment_detector.py`
- `main.py`
- `requirements.txt`
- `test_project.py`

## Copilot Response

I’m implementing explicit sync/async chains with offline-testable detection and diagnostic classification; no Azure resources or live authentication will be invoked during validation.

Validation succeeded: all 5 Python files compiled, and all 7 unit tests passed.

The core project is working. I’m tightening Azure Pipelines compatibility to recognize the standard service-connection variable names in addition to generic Azure Identity variables.

Validation succeeded: all 5 files compiled, all 8 unit tests passed, and `python main.py --help` completed successfully.

Created the Python Azure credential-chain project with:

- Environment detection for local development, CI/CD, and production
- Sync and async credential factories
- Developer-tool credentials for local use
- Environment credentials and Azure Pipelines workload federation for CI
- System/user-assigned managed identity with Kubernetes workload identity fallback
- Optional CAE token requests via `--cae` or `AZURE_ENABLE_CAE=true`
- Detailed sync/async connectivity diagnostics against Azure Resource Manager
- Required dependencies in `requirements.txt`
- Offline unit tests in `test_project.py`

## Reviewer Notes (LLM-as-Judge)

### Summary

0/1 criteria passed. Authentication-specific exceptions are handled well, but Azure SDK HTTP exception coverage is incomplete.

### Strengths

- CredentialUnavailableError and ClientAuthenticationError receive targeted handling, including useful diagnostics for expired credentials, invalid tenants, unavailable identities, and federation failures.

### Issues

- Add an explicit HttpResponseError handler after the specialized authentication handlers and report its status/error details appropriately in both sync and async testers.

## Grader Results

- credential-chain.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (14/14)
      - Uses `ChainedTokenCredential` to compose multiple credentials: Pass
      - Credentials passed in order to the constructor — order matters: Pass
      - **Dev chain**: includes `AzureCliCredential`; may include `VisualStudioCodeCredential`, `AzurePowerShellCredential`: Pass
      - **CI chain**: uses `EnvironmentCredential` or `AzurePipelinesCredential` (not just `DefaultAzureCredential`): Pass
      - **Production chain**: `ManagedIdentityCredential` first (supports user-assigned via `client_id` kwarg), `WorkloadIdentityCredential` as fallback: Pass
      - Enables CAE via `enable_cae=True` in `get_token()` call or token request context: Pass
      - Detects CI (checks `CI`, `TF_BUILD`, `BUILD_SOURCESDIRECTORY`, or similar env vars): Pass
      - Detects production/managed identity (checks `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, or similar): Pass
      - Falls back to dev if neither detected: Pass
      - Calls `get_token()` with correct scope (`https://management.azure.com/.default`): Pass
      - Prints token expiry from `AccessToken.expires_on`: Pass
      - Handles failure with specific exception info from `ClientAuthenticationError`: Pass
      - Async tester uses `azure.identity.aio` async credentials and `await get_token()`: Pass
      - NOT using `DefaultAzureCredential` as the CI credential (too broad): Pass
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
  - Proper Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Fail
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 6.0000** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-python-credential-chain --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
