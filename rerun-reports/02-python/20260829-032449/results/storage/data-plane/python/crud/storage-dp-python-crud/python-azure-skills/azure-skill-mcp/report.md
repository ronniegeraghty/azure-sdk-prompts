# Evaluation Report: storage-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 190.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 14/15 |
| Duration | 190.0s |
| Timestamp | 2026-08-28T22:17:00Z |
| Files Generated | 2 |
| Event Count | 2347 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 94.4s |
| Review | 94.5s |
| **Total** | **190.0s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=321551 out=2398 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 2330ms

**Input:**

```json
{"skill":"azure-storage"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Services

## Services

| Service | Use When | MCP Tools | CLI |
|---------|----------|-----------|-----|
| Blob Storage | Objects, files, backups, static content | `azure__storage` | `az storage blob` |
| File Shares | SMB file shares, lift-and-shift | - | `az storage file` |
| Queue Storage | Async messaging, task queues | - | `az storage queue` |
| Table Storage | NoSQL key-value (consider Cosmos DB) | - | `az storage table` |
| Data Lake | Big data analytics, hierarchical namespace | - | `az storage fs` |

## MCP Server (Preferred)

When Azure MCP is enabled:

- `azure__storage` with command `storage_account_list` - List storage accounts
- `azure__storage` with command `storage_container_list` - List containers in account
- `azure__storage` with command `storage_blob_list` - List blobs in container
- `azure__storage` with command `storage_blob_get` - Download blob content
- `azure__storage` with command `storage_blob_put` - Upload blob content

**If Azure MCP is not enabled:** Run `/azure:setup` or enable via `/mcp`.

## CLI Fallback

```bash
# List storage accounts
az storage account list --output table

# List containers
az storage container list --account-name ACCOUNT --output table

# List blobs
az storage blob list --account-name ACCOUNT --container-name CONTAINER --output table

# Download blob
az storage blob download --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH

# Upload blob
az storage blob upload --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH
```

## Storage Account Tiers

| Tier | Use Case | Performance |
|------|----------|-------------|
| Standard | General purpose, backup | Milliseconds |
| Premium | Databases, high IOPS | Sub-millisecond |

## Blob Access Tiers

| Tier | Access Frequency | Cost |
|------|-----------------|------|
| Hot | Frequent | Higher storage, lower access |
| Cool | Infrequent (30+ days) | Lower storage, higher ac
... (truncated)
```

### ✅ 2. view — 15628ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk\\azure-storage-blob-py.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-py.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-py.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-py.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-py.md
@@ -1,25 +1,25 @@
 # Blob Storage — Python SDK Quick Reference
 
 > Condensed from **azure-storage-blob-py**. Full patterns (SAS tokens,
 > async client, performance tuning, blob properties/metadata)
 > in the **azure-storage-blob-py** plugin skill if installed.
 
 ## Install
 pip install azure-storage-blob azure-identity
 
 ## Quick Start
 ```python
 from azure.storage.blob import BlobServiceClient
 from azure.identity import DefaultAzureCredential
 blob_service_client = BlobServiceClient("https://<account>.blob.core.windows.net", DefaultAzureCredential())
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** — in production, use ManagedIdentityCredential. See [auth-best-practices.md](../auth-best-practices.md)
 - Use context managers for async clients
 - Set `overwrite=True` explicitly when re-uploading
 - Use `max_concurrency` for large file transfers
 - Prefer `readinto()` over `readall()` for memory efficiency
 - Use `walk_blobs()` for hierarchical listing
 - Set appropriate content types for web-served blobs
 

```

### ✅ 3. view — 14109ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
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
 2. **Broad surface area** — checks environment variables, CLI tokens, and oth
... (truncated)
```

### ✅ 4. glob — 11498ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 10251ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 12824ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate a local-only Python script using DefaultAzureCredential and azure-storage-blob for Blob Storage container and blob CRUD operations with proper Azure SDK exception handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 3258ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only Python script using DefaultAzureCredential and azure-storage-blob for Blob Storage CRUD operations.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. apply_patch — 2699ms

**Input:**

```json
"*** Begin Patch\n*** Add File: blob_crud.py\n+import logging\n+import os\n+from pathlib import Path\n+\n+from azure.core.exceptions import HttpResponseError, ResourceExistsError\n+from azure.identity import DefaultAzureCredential\n+from azure.storage.blob import BlobServiceClient\n+\n+\n+CONTAINER_NAME = \"my-container\"\n+BLOB_NAME = \"reports/report.csv\"\n+SOURCE_FILE = Path(\"report.csv\")\n+DOWNLOAD_FILE = Path(\"report-downloaded.csv\")\n+\n+\n+def run_blob_crud() -\u003e None:\n+    account_url = os.environ.get(\"AZURE_STORAGE_ACCOUNT_URL\")\n+    if not account_url:\n+        raise ValueError(\n+            \"AZURE_STORAGE_ACCOUNT_URL must be set, for example \"\n+            \"https://\u003cstorage-account\u003e.blob.core.windows.net\"\n+        )\n+\n+    credential = DefaultAzureCredential()\n+    blob_service_client = BlobServiceClient(\n+        account_url=account_url,\n+        credential=credential,\n+    )\n+\n+    try:\n+        container_client = blob_service_client.get_container_client(CONTAINER_NAME)\n+        try:\n+            container_client.create_container()\n+            logging.info(\"Created container %s.\", CONTAINER_NAME)\n+        except ResourceExistsError:\n+            logging.info(\"Container %s already exists.\", CONTAINER_NAME)\n+\n+        blob_client = container_client.get_blob_client(BLOB_NAME)\n+        with SOURCE_FILE.open(\"rb\") as source:\n+            blob_client.upload_blob(source, overwrite=True)\n+        logging.info(\"Uploaded %s as %s.\", SOURCE_FILE, BLOB_NAME)\n+\n+        for blob in container_client.list_blobs():\n+            print(f\"{blob.name}: {blob.size} bytes\")\n+\n+        with DOWNLOAD_FILE.open(\"wb\") as destination:\n+            blob_client.download_blob().readinto(destination)\n+        logging.info(\"Downloaded %s to %s.\", BLOB_NAME, DOWNLOAD_FILE)\n+\n+        blob_client.delete_blob()\n+        logging.info(\"Deleted blob %s.\", BLOB_NAME)\n+\n+        container_client.delete_container()\n+        logging.info(\"Deleted container %s.\", CONTAINER_NAME)\n+    except HttpResponseError:\n+        logging.exception(\"An Azure Blob Storage operation failed.\")\n+        raise\n+    finally:\n+        blob_service_client.close()\n+        credential.close()\n+\n+\n+if __name__ == \"__main__\":\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s: %(message)s\")\n+    try:\n+        run_blob_crud()\n+    except (OSError, ValueError) as error:\n+        logging.error(\"%s\", error)\n+        raise SystemExit(1) from error\n+    except HttpResponseError as error:\n+        logging.error(\n+            \"Azure request failed (status %s): %s\",\n+            error.status_code,\n+            error.message,\n+        )\n+        raise SystemExit(1) from error\n*** Add File: requirements.txt\n+azure-identity\n+azure-storage-blob\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py
@@ -1,0 +1,77 @@
+import logging
+import os
+from pathlib import Path
+
+from azure.core.exceptions import HttpResponseError, ResourceExistsError
+from azure.identity import DefaultAzureCredential
+from azure.storage.blob import BlobServiceClient
+
+
+CONTAINER_NAME = "my-container"
+BLOB_NAME = "reports/report.csv"
+SOURCE_FILE = Path("report.csv")
+DOWNLOAD_FILE = Path("report-downloaded.csv")
+
+
+def run_blob_crud() -> None:
+    account_url = os.environ.get("AZURE_STORAGE_ACCOUNT_URL")
+    if not account_url:
+        raise ValueError(
+            "AZURE_STORAGE_ACCOUNT_URL must be set, for example "
+            "https://<storage-account>.blob.core.windows.net"
+        )
+
+    credential = DefaultAzureCredential()
+    blob_service_client = BlobServiceClient(
+        account_url=account_url,
+        credential=credential,
+    )
+
+    try:
+        container_client = blob_service_client.get_container_client(CONTAINER_NAME)
+        try:
+            container_client.create_container()
+            logging.info("Created container %s.", CONTAINER_NAME)
+        except ResourceExistsError:
+            logging.info("Container %s already exists.", CONTAINER_NAME)
+
+        blob_client = container_client.get_blob_client(BLOB_NAME)
+        with SOURCE_FILE.open("rb") as source:
+            blob_client.upload_blob(source, overwrite=True)
+        logging.info("Uploaded %s as %s.", SOURCE_FILE, BLOB_NAME)
+
+        for blob in container_client.list_blobs():
+            print(f"{blob.name}: {blob.size} bytes")
+

... (truncated)
```

### ✅ 9. apply_patch — 1631ms

**Input:**

```json
"*** Begin Patch\n*** Update File: blob_crud.py\n@@\n-    except HttpResponseError:\n-        logging.exception(\"An Azure Blob Storage operation failed.\")\n-        raise\n     finally:\n         blob_service_client.close()\n         credential.close()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-crud-python-azure-skills-azure-skill-mcp-100811347/blob_crud.py
@@ -52,9 +52,6 @@
 
         container_client.delete_container()
         logging.info("Deleted container %s.", CONTAINER_NAME)
-    except HttpResponseError:
-        logging.exception("An Azure Blob Storage operation failed.")
-        raise
     finally:
         blob_service_client.close()
         credential.close()

```

### ✅ 10. powershell — 4491ms

**Input:**

```json
{"command":"python -m py_compile .\\blob_crud.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; Get-Content .\\requirements.txt","description":"Validate script syntax and dependencies","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-identity
azure-storage-blob
<shellId: 0 completed with exit code 0>
```

### ✅ 11. powershell — 1145ms

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -LiteralPath .\\__pycache__ -Recurse -Force }; Get-ChildItem -File | Select-Object -ExpandProperty Name","description":"Clean compile artifacts and list files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
blob_crud.py
requirements.txt
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `blob_crud.py`
- `requirements.txt`

## Copilot Response

I’m grounding the implementation in the current Azure SDK authentication and Blob Storage patterns, then I’ll create the script and dependency file.

The script will use an environment-provided account URL, passwordless `DefaultAzureCredential`, streaming file transfers, and operation-specific Azure SDK error handling without embedding credentials.

Created:

- `blob_crud.py` — performs container creation, upload, listing, download, blob deletion, and container deletion using `DefaultAzureCredential`, with `ResourceExistsError` and `HttpResponseError` handling.
- `requirements.txt` — includes `azure-identity` and `azure-storage-blob`.

Set `AZURE_STORAGE_ACCOUNT_URL` before running. Reference: [Azure Blob Storage Python quickstart](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-python).

## Reviewer Notes (LLM-as-Judge)

### Summary

The generated script fully satisfies the exception-handling criterion.

### Strengths

- Uses targeted handling for ResourceExistsError while providing broader HttpResponseError handling for Azure SDK operations.
- Ensures clients and credentials are closed with a finally block.

## Grader Results

- crud-blobs.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - Includes required Azure Storage and Identity SDK packages (via requirements.txt or install instructions): Pass
      - Authenticates the blob service client with DefaultAzureCredential: Pass
      - Creates a container (handling the case where it already exists): Pass
      - Uploads a file as a blob with overwrite support: Pass
      - Lists blobs in the container with their properties: Pass
      - Downloads a blob to a local file: Pass
      - Deletes the blob and the container: Pass
      - Catches Azure-specific errors (both general HTTP errors and resource-exists errors): Pass
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
hyoka run --prompt-id storage-dp-python-crud --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
