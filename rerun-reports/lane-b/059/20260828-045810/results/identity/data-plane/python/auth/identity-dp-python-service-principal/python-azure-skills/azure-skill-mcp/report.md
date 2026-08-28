# Evaluation Report: identity-dp-python-service-principal

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 520.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-python-service-principal` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/12 |
| Duration | 520.6s |
| Timestamp | 2026-08-27T20:58:10Z |
| Files Generated | 4 |
| Event Count | 5167 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 414.5s |
| Review | 104.9s |
| **Total** | **520.6s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | entra-app-registration |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=239632 out=3778 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3994ms

**Input:**

```json
{"skill":"entra-app-registration"}
```

**Output:**

```
Skill loaded successfully ✅

## Overview

Microsoft Entra ID (formerly Azure Active Directory) is Microsoft's cloud-based identity and access management service. App registrations allow applications to authenticate users and access Azure resources securely.

### Key Concepts

| Concept | Description |
|---------|-------------|
| **App Registration** | Configuration that allows an app to use Microsoft identity platform |
| **Application (Client) ID** | Unique identifier for your application |
| **Tenant ID** | Unique identifier for your Azure AD tenant/directory |
| **Client Secret** | Password for the application (confidential clients only) |
| **Redirect URI** | URL where authentication responses are sent |
| **API Permissions** | Access scopes your app requests |
| **Service Principal** | Identity created in your tenant when you register an app |

### Application Types

| Type | Use Case |
|------|----------|
| **Web Application** | Server-side apps, APIs |
| **Single Page App (SPA)** | JavaScript/React/Angular apps |
| **Mobile/Native App** | Desktop, mobile apps |
| **Daemon/Service** | Background services, APIs |

## Core Workflow

### Step 1: Register the Application

Create an app registration in the Azure portal or using Azure CLI.

**Portal Method:**
1. Navigate to Azure Portal → Microsoft Entra ID → App registrations
2. Click "New registration"
3. Provide name, supported account types, and redirect URI
4. Click "Register"

**CLI Method:** See [references/cli-commands.md](references/cli-commands.md)
**IaC Method:** See [references/BICEP-EXAMPLE.bicep](references/BICEP-EXAMPLE.bicep)

It's highly recommended to use the IaC to manage Entra app registration if you already use IaC in your project, need a scalable solution for managing lots of app registrations or need fine-grained audit history of the configuration changes. 

### Step 2: Configure Authentication

Set up authentication settings based on your application type.

- **Web Apps**: Add redirect 
... (truncated)
```

### ✅ 2. view — 191195ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\sdk\\azure-identity-py.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-py.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-py.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-py.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-py.md
@@ -1,28 +1,28 @@
 # Authentication — Python SDK Quick Reference
 
 > Condensed from **azure-identity-py**. Full patterns (async,
 > ChainedTokenCredential, token caching, all credential types)
 > in the **azure-identity-py** plugin skill if installed.
 
 ## Install
 ```bash
 pip install azure-identity
 ```
 
 ## Quick Start
 > **Auth:** `DefaultAzureCredential` is for local development. See [auth-best-practices.md](../auth-best-practices.md) for production patterns.
 
 ```python
 from azure.identity import DefaultAzureCredential
 credential = DefaultAzureCredential()
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** (CLI, PowerShell, VS Code). In production, use ManagedIdentityCredential — see [auth-best-practices.md](../auth-best-practices.md)
 - Never hardcode credentials — use environment variables or managed identity
 - Prefer managed identity in production Azure deployments
 - Use ChainedTokenCredential when you need a custom credential order
 - Close async credentials explicitly or use context managers
 - Set AZURE_CLIENT_ID env var for user-assigned managed identities
 - Exclude unused credentials to speed up authentication
 

```

### ✅ 3. view — 189650ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks enviro
... (truncated)
```

### ✅ 4. glob — 188476ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 191197ms

**Input:**

```json
{"intent":"Create a local-only runnable Python example that uses azure-identity ClientSecretCredential and azure-mgmt-resource ResourceManagementClient, loads service principal settings from environment variables or a gitignored .env file, and handles authentication and Azure SDK HTTP failures. Do not provision or modify Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. apply_patch — 3416ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-mgmt-resource\u003e=23.1.0,\u003c24.0.0\n+python-dotenv\u003e=1.0.1,\u003c2.0.0\n*** Add File: .env.example\n+# Copy this file to .env for local development only.\n+# Never commit .env or put real credentials in this example file.\n+AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000\n+AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000\n+AZURE_CLIENT_SECRET=replace-with-a-local-development-secret\n+AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\n*** Add File: .gitignore\n+.env\n+.venv/\n+__pycache__/\n+*.py[cod]\n+.pytest_cache/\n*** Add File: app.py\n+\"\"\"Authenticate with an Azure service principal and list resource groups.\"\"\"\n+\n+from __future__ import annotations\n+\n+import logging\n+import os\n+import sys\n+from dataclasses import dataclass\n+from typing import Sequence\n+\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ServiceRequestError,\n+)\n+from azure.identity import ClientSecretCredential\n+from azure.mgmt.resource import ResourceManagementClient\n+from dotenv import load_dotenv\n+\n+LOGGER = logging.getLogger(__name__)\n+ARM_SCOPE = \"https://management.azure.com/.default\"\n+\n+\n+class ConfigurationError(ValueError):\n+    \"\"\"Raised when required application configuration is missing.\"\"\"\n+\n+\n+@dataclass(frozen=True)\n+class AzureSettings:\n+    tenant_id: str\n+    client_id: str\n+    client_secret: str\n+    subscription_id: str\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"AzureSettings\":\n+        values = {\n+            name: os.getenv(name, \"\").strip()\n+            for name in (\n+                \"AZURE_TENANT_ID\",\n+                \"AZURE_CLIENT_ID\",\n+                \"AZURE_CLIENT_SECRET\",\n+                \"AZURE_SUBSCRIPTION_ID\",\n+            )\n+        }\n+        missing = [name for name, value in values.items() if not value]\n+        if missing:\n+            raise ConfigurationError(\n+                \"Missing required environment variables: \" + \", \".join(missing)\n+            )\n+\n+        return cls(\n+            tenant_id=values[\"AZURE_TENANT_ID\"],\n+            client_id=values[\"AZURE_CLIENT_ID\"],\n+            client_secret=values[\"AZURE_CLIENT_SECRET\"],\n+            subscription_id=values[\"AZURE_SUBSCRIPTION_ID\"],\n+        )\n+\n+\n+def create_credential(settings: AzureSettings) -\u003e ClientSecretCredential:\n+    \"\"\"Create the deterministic credential used by this service.\"\"\"\n+    return ClientSecretCredential(\n+        tenant_id=settings.tenant_id,\n+        client_id=settings.client_id,\n+        client_secret=settings.client_secret,\n+    )\n+\n+\n+def list_resource_group_names(\n+    settings: AzureSettings,\n+    credential: ClientSecretCredential,\n+) -\u003e list[str]:\n+    \"\"\"Verify authentication and return resource group names.\"\"\"\n+    credential.get_token(ARM_SCOPE)\n+    client = ResourceManagementClient(credential, settings.subscription_id)\n+    return [resource_group.name for resource_group in client.resource_groups.list()]\n+\n+\n+def run() -\u003e int:\n+    load_dotenv(override=False)\n+\n+    try:\n+        settings = AzureSettings.from_environment()\n+        credential = create_credential(settings)\n+        names = list_resource_group_names(settings, credential)\n+    except ConfigurationError as error:\n+        LOGGER.error(\"Configuration error: %s\", error)\n+        return 2\n+    except ClientAuthenticationError:\n+        LOGGER.error(\n+            \"Azure authentication failed. Verify the tenant ID, client ID, \"\n+            \"client secret, secret expiration, and service principal status.\"\n+        )\n+        return 3\n+    except ServiceRequestError as error:\n+        LOGGER.error(\"Could not reach Azure: %s\", error)\n+        return 4\n+    except HttpResponseError as error:\n+        status_code = getattr(error, \"status_code\", None)\n+        LOGGER.error(\n+            \"Azure rejected the Resource Manager request (status %s). Verify \"\n+            \"the subscription ID and least-privilege RBAC assignment.\",\n+            status_code if status_code is not None else \"unknown\",\n+        )\n+        return 5\n+\n+    if names:\n+        print(\"Resource groups:\")\n+        for name in names:\n+            print(f\"- {name}\")\n+    else:\n+        print(\"No resource groups found.\")\n+    return 0\n+\n+\n+def main(argv: Sequence[str] | None = None) -\u003e int:\n+    del argv\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+    return run()\n+\n+\n+if __name__ == \"__main__\":\n+    sys.exit(main())\n*** Add File: test_app.py\n+\"\"\"Local-only tests for configuration and Azure error handling.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+import unittest\n+from unittest.mock import MagicMock, patch\n+\n+from azure.core.exceptions import ClientAuthenticationError\n+\n+import app\n+\n+\n+class AzureSettingsTests(unittest.TestCase):\n+    @patch.dict(os.environ, {}, clear=True)\n+    def test_missing_environment_variables_are_reported(self) -\u003e None:\n+        with self.assertRaisesRegex(\n+            app.ConfigurationError, \"AZURE_CLIENT_SECRET\"\n+        ):\n+            app.AzureSettings.from_environment()\n+\n+\n+class RunTests(unittest.TestCase):\n+    SETTINGS = {\n+        \"AZURE_TENANT_ID\": \"tenant\",\n+        \"AZURE_CLIENT_ID\": \"client\",\n+        \"AZURE_CLIENT_SECRET\": \"secret\",\n+        \"AZURE_SUBSCRIPTION_ID\": \"subscription\",\n+    }\n+\n+    @patch.dict(os.environ, SETTINGS, clear=True)\n+    @patch(\"app.list_resource_group_names\", return_value=[\"example-rg\"])\n+    def test_success(self, list_names: MagicMock) -\u003e None:\n+        with patch(\"app.create_credential\", return_value=MagicMock()):\n+            self.assertEqual(app.run(), 0)\n+        list_names.assert_called_once()\n+\n+    @patch.dict(os.environ, SETTINGS, clear=True)\n+    @patch(\n+        \"app.list_resource_group_names\",\n+        side_effect=ClientAuthenticationError(\"invalid credential\"),\n+    )\n+    def test_authentication_failure_has_distinct_exit_code(\n+        self, list_names: MagicMock\n+    ) -\u003e None:\n+        with patch(\"app.create_credential\", return_value=MagicMock()):\n+            self.assertEqual(app.run(), 3)\n+        list_names.assert_called_once()\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** Add File: README.md\n+# Azure service principal authentication with Python\n+\n+This runnable example uses the OAuth 2.0 client credentials flow through\n+`ClientSecretCredential`, then passes that credential to the Azure Resource\n+Manager SDK to list resource groups. The Azure operation is read-only.\n+\n+## Requirements\n+\n+- Python 3.10 or later\n+- A Microsoft Entra service principal with a client secret\n+- The service principal assigned the least-privilege `Reader` role at the\n+  subscription or narrower scope that it must inspect\n+\n+Install the required pip packages in a virtual environment:\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+python -m pip install -r requirements.txt\n+```\n+\n+The packages are:\n+\n+- `azure-identity` for `ClientSecretCredential`\n+- `azure-mgmt-resource` for `ResourceManagementClient`\n+- `python-dotenv` for optional local `.env` loading\n+\n+## Configure and run\n+\n+For local development, copy `.env.example` to `.env` and replace each\n+placeholder. The application calls `load_dotenv(override=False)`, so existing\n+environment variables take precedence over `.env` values.\n+\n+Alternatively, set the values directly in the shell:\n+\n+```powershell\n+$env:AZURE_TENANT_ID = \"\u003ctenant-id\u003e\"\n+$env:AZURE_CLIENT_ID = \"\u003capplication-client-id\u003e\"\n+$env:AZURE_CLIENT_SECRET = \"\u003cclient-secret-value\u003e\"\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+python app.py\n+```\n+\n+`create_credential` constructs the credential explicitly from `tenant_id`,\n+`client_id`, and `client_secret`. Before creating\n+`ResourceManagementClient`, the application requests an Azure Resource\n+Manager token so authentication failures are reported separately from\n+authorization, subscription, and network failures.\n+\n+Run the local-only tests without making Azure requests:\n+\n+```powershell\n+python -m unittest -v\n+```\n+\n+## Secret-management practices\n+\n+- Never hardcode or commit a client secret. `.env` is gitignored; commit only\n+  `.env.example` with placeholders.\n+- Treat `.env` as a local-development convenience, not a production secret\n+  store. Inject production secrets through the hosting platform or retrieve\n+  them from Azure Key Vault.\n+- Restrict access with least-privilege Azure RBAC at the narrowest practical\n+  scope, set secret expirations, and rotate secrets regularly.\n+- Do not log credentials or exception details that may contain sensitive\n+  request data. This example emits actionable messages without printing the\n+  secret.\n+- Prefer managed identity for Azure-hosted production workloads, or workload\n+  identity/certificate credentials where managed identity is unavailable, to\n+  avoid long-lived client secrets.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/requirements.txt
@@ -1,0 +1,4 @@
+azure-identity>=1.17.0,<2.0.0
+azure-mgmt-resource>=23.1.0,<24.0.0
+python-dotenv>=1.0.1,<2.0.0
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/.env.example b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/.env.example
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/.env.example
@@ -1,0 +1,7 @@
+# Copy this file to .env for local development only.
+# Never commit .env or put real credentials in this example file.
+AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000
+AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000
+AZURE_CLIENT_SECRET=replace-with-a-local-development-secret
+AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/.gitignore b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/.gitignore
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919/.gitignore
@@ -1,0 +1,6 
... (truncated)
```

### ✅ 7. powershell — 6455ms

**Input:**

```json
{"command":"python -m compileall -q . \u0026\u0026 python -m unittest -v","description":"Compile project and run unit tests","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_app (unittest.loader._FailedTest.test_app) ... ERROR

======================================================================
ERROR: test_app (unittest.loader._FailedTest.test_app)
----------------------------------------------------------------------
ImportError: Failed to import test module: test_app
Traceback (most recent call last):
  File "C:\ProgramData\anaconda3\Lib\unittest\loader.py", line 396, in _find_test_path
    module = self._get_module_from_name(name)
  File "C:\ProgramData\anaconda3\Lib\unittest\loader.py", line 339, in _get_module_from_name
    __import__(name)
    ~~~~~~~~~~^^^^^^
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919\test_app.py", line 11, in <module>
    import app
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919\app.py", line 17, in <module>
    from azure.mgmt.resource import ResourceManagementClient
ImportError: cannot import name 'ResourceManagementClient' from 'azure.mgmt.resource' (unknown location)


----------------------------------------------------------------------
Ran 1 test in 0.000s

FAILED (errors=1)
<shellId: 0 completed with exit code 1>
```

### ✅ 8. powershell — 101280ms

**Input:**

```json
{"command":"python -m venv .venv \u0026\u0026 .\\.venv\\Scripts\\python.exe -m pip install --quiet -r requirements.txt \u0026\u0026 .\\.venv\\Scripts\\python.exe -m unittest -v","description":"Install dependencies and run tests in venv","initial_wait":180,"mode":"sync"}
```

**Output:**

```
[notice] A new release of pip is available: 25.2 -> 26.2.1
[notice] To update, run: C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-python-service-principal-python-azure-skills-azure-skill-mcp-3768024919\.venv\Scripts\python.exe -m pip install --upgrade pip
test_missing_environment_variables_are_reported (test_app.AzureSettingsTests.test_missing_environment_variables_are_reported) ... ok
test_authentication_failure_has_distinct_exit_code (test_app.RunTests.test_authentication_failure_has_distinct_exit_code) ... Azure authentication failed. Verify the tenant ID, client ID, client secret, secret expiration, and service principal status.
ok
test_success (test_app.RunTests.test_success) ... ok
Resource groups:
- example-rg

----------------------------------------------------------------------
Ran 3 tests in 0.006s

OK
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `app.py`
- `requirements.txt`
- `test_app.py`

## Copilot Response

I’ll ground the example in the Azure Identity guidance, then create the local project and validate its failure paths without contacting Azure.

The workspace is empty. I’m creating a small read-only Resource Manager example with explicit configuration validation, distinct authentication/authorization/network errors, and tests that use mocks only.

The system Python has an incomplete `azure.mgmt.resource` installation, so I’m validating in the project’s isolated virtual environment using the declared dependencies.

Created a complete runnable Python project with:

- `app.py` — `ClientSecretCredential`, `ResourceManagementClient`, and explicit authentication/network/HTTP error handling
- `requirements.txt` — required Azure SDK and dotenv packages
- `.env.example` and `.gitignore` — safe local secret configuration
- `README.md` — setup, execution, RBAC, rotation, Key Vault, and managed identity guidance
- `test_app.py` — local-only tests with mocked Azure access

Run with `.\.venv\Scripts\python.exe app.py` after copying `.env.example` to `.env` and supplying valid service-principal values.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Correct exception ordering preserves specialized authentication handling while covering other HttpResponseError subclasses.
- Authentication, network, and general Azure HTTP failures receive distinct messages and exit codes.

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - `azure-identity` package with `ClientSecretCredential` class: Pass
      - Constructor keyword arguments: tenant_id, client_id, client_secret: Pass
      - Passing credential to Azure SDK clients: Pass
      - `os.environ` or python-dotenv patterns: Pass
      - `ClientAuthenticationError` exception handling: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Fail
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 6.0000** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-python-service-principal --config python-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
