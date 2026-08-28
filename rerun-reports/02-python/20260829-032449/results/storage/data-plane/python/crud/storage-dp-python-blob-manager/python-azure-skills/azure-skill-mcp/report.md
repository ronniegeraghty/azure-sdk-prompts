# Evaluation Report: storage-dp-python-blob-manager

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 298.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-python-blob-manager` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 298.7s |
| Timestamp | 2026-08-28T22:03:47Z |
| Files Generated | 6 |
| Event Count | 16273 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 190.8s |
| Review | 106.7s |
| **Total** | **298.7s** |

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
| Token Usage | in=424783 out=12421 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 2853ms

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

### ✅ 2. view — 15985ms

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

### ✅ 3. view — 14549ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk\\azure-storage-blob-py.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
@@ -1,115 +1,115 @@
 # Azure Storage SDK Usage
 
 SDK packages and quick start examples for Azure Storage services.
 
 ## Storage SDKs by Language
 
 | Language | Blob | Queue | File Share | Data Lake |
 |----------|------|-------|------------|----------|
 | .NET | `Azure.Storage.Blobs` | `Azure.Storage.Queues` | `Azure.Storage.Files.Shares` | `Azure.Storage.Files.DataLake` |
 | Java | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | JavaScript | `@azure/storage-blob` | `@azure/storage-queue` | `@azure/storage-file-share` | `@azure/storage-file-datalake` |
 | Python | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | Go | `azblob` | `azqueue` | `azfile` | `azdatalake` |
 | Rust | `azure_storage_blob` | `azure_storage_queue` | - | - |
 
 ## Installation Commands
 
 | Language | Install Blob SDK + Identity |
 |----------|-----------------------------|
 | .NET | `dotnet add package Azure.Storage.Blobs` `dotnet add package Azure.Identity` |
 | Java | Maven: `com.azure:azure-storage-blob` `com.azure:azure-identity` |
 | JavaScript | `npm install @azure/storage-blob @azure/identity` |
 | Python | `pip install azure-storage-blob azure-identity` |
 | Go | `go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob github.com/Azure/azure-sdk-for-g
... (truncated)
```

### ✅ 4. view — 13207ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
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

### ✅ 5. azure-get_azure_bestpractices (via azure) — 11819ms

**Input:**

```json
{"intent":"Create a local-only Python project using DefaultAzureCredential and Azure Blob Storage sync/async SDKs, with streaming large uploads, ETag concurrency, leases, retries, logging, and operation timeouts.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 17193ms

**Input:**

```json
{"intent":"Find current official Python Azure Blob Storage SDK guidance for streaming uploads, async clients, ETag conditional writes, blob leases, retry policies, logging, and timeout parameters.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. glob — 10572ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. azure-get_azure_bestpractices (via azure) — 5408ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure Python SDK code-generation best practices to a reusable Blob Storage utility.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 6043ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official details for Azure Blob Python API behavior.","parameters":{"query":"Azure Storage Blob Python upload_blob streaming max_concurrency ETag MatchConditions lease timeout retry policy logging async BlobServiceClient"}}
```

**Output:**

```
{"results":[{"title":"ContainerClient Class-Methods","content":"# ContainerClient Class\r\n\r\nA client to interact with a specific container, although that container may not yet exist.\r\n\r\nFor operations relating to a specific blob within this container, a blob client can be retrieved using the [get_blob_client](https://learn.microsoft.com/python/api/azure-storage-blob/azure.storage.blob.aio.containerclient?view=azure-python#azure-storage-blob-aio-containerclient-get-blob-client) function.\r\n\r\n## Methods\r\n\r\n### upload\\_blob\r\n\r\nCreates a new blob from a data source with automatic chunking.\r\n\r\n\u0060\u0060\u0060python\r\nasync upload_blob(name: str, data: bytes | str | Iterable | AsyncIterable | IO, blob_type: str | BlobType = BlobType.BLOCKBLOB, length: int | None = None, metadata: Dict[str, str] | None = None, **kwargs) -\u003E BlobClient\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r\n| name\u003Cbr\u003E\u003Cbr\u003ERequired | [str](https://docs.python.org/3/library/stdtypes.html#str)\u003Cbr\u003E\u003Cbr\u003EThe blob with which to interact. |\r\n| data\u003Cbr\u003E\u003Cbr\u003ERequired | [Union](https://docs.python.org/3/library/typing.html#typing.Union)[[bytes](https://docs.python.org/3/library/stdtypes.html#bytes), [str](https://docs.python.org/3/library/stdtypes.html#str), [Iterable](https://docs.python.org/3/library/typing.html#typing.Iterable)[[AnyStr](https://docs.python.org/3/library/typing.html#typing.AnyStr)], [AsyncIterable](https://docs.python.org/3/library/typing.html#typing.AsyncIterable)[[AnyStr](https://docs.python.org/3/library/typing.html#typing.AnyStr)], [IO](https://docs.python.org/3/library/typing.html#typing.IO)[[AnyStr](https://docs.python.org/3/library/typing.html#typing.AnyStr)]] \u003Cbr\u003E\u003Cbr\u003EThe blob data to upload. |\r\n| blob\\_type | [BlobType](https://learn.microsoft.com/python/api/azure-storage-blob/azure.storage.blob.blobtype?view=azure-python)\u0
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 5727ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python samples for blob upload/download/leases and concurrency.","parameters":{"language":"python","query":"Azure Storage Blob Python BlobClient upload_blob etag match_condition BlobLeaseClient async streaming download readinto timeout"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates how to create an asynchronous client for Azure Blob Storage using the async version of the Azure Blob Storage library, upload a local file to Blob Storage, and properly close the credential asynchronously.\npackage: azure.storage.blob.aio, azure.identity\nlanguage: python\n","codeSnippet":"credential = DefaultAzureCredential()\r\n\r\nasync def run():\r\n\r\n    async with BlobClient(\r\n        storage_url,\r\n        container_name=\u0022blob-container-01\u0022,\r\n        blob_name=f\u0022sample-blob-{str(uuid.uuid4())[0:5]}.txt\u0022,\r\n        credential=credential,\r\n    ) as blob_client:\r\n\r\n        # Open a local file and upload its contents to Blob Storage\r\n        with open(\u0022./sample-source.txt\u0022, \u0022rb\u0022) as data:\r\n            await blob_client.upload_blob(data)\r\n            print(f\u0022Uploaded sample-source.txt to {blob_client.url}\u0022)\r\n\r\n        # Close credential\r\n        await credential.close()\r\n\r\nasyncio.run(run())","link":"https://learn.microsoft.com/azure/developer/python/sdk/azure-sdk-library-usage-patterns#asynchronous-operations","language":"python"},{"description":"description: Creates a blob client asynchronously using a connection string, container name, and blob name, then downloads a blob to a local file asynchronously.\npackage: azure-storage-blob\nlanguage: python\n","codeSnippet":"from azure.storage.blob.aio import BlobClient\r\n\r\nblob = BlobClient.from_connection_string(conn_str=\u0022\u003Cconnection_string\u003E\u0022, container_name=\u0022mycontainer\u0022, blob_name=\u0022my_blob\u0022)\r\n\r\nwith open(\u0022./BlockDestination.txt\u0022, \u0022wb\u0022) as my_blob:\r\n    stream = await blob.download_blob()\r\n    data = await stream.readall()\r\n    my_blob.write(data)","link":"https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme?view=azure-python#examples","language":"python"},{"description":"description: Creates
... (truncated)
```

### ✅ 11. apply_patch — 693ms

**Input:**

```json
"*** Begin Patch\n*** Add File: blob_manager/__init__.py\n+\"\"\"Reusable synchronous and asynchronous Azure Blob Storage utilities.\"\"\"\n+\n+from .config import BlobStorageSettings, create_async_client, create_sync_client\n+from .service import AsyncBlobStorageService, BlobStorageService, OperationResult\n+\n+__all__ = [\n+    \"AsyncBlobStorageService\",\n+    \"BlobStorageService\",\n+    \"BlobStorageSettings\",\n+    \"OperationResult\",\n+    \"create_async_client\",\n+    \"create_sync_client\",\n+]\n*** Add File: blob_manager/config.py\n+\"\"\"Configuration and authenticated client factories for Azure Blob Storage.\"\"\"\n+\n+from __future__ import annotations\n+\n+import logging\n+import os\n+from dataclasses import dataclass\n+\n+from azure.core.pipeline.policies import RetryPolicy\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.storage.blob import BlobServiceClient\n+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient\n+\n+\n+def _positive_int(name: str, default: int) -\u003e int:\n+    raw_value = os.getenv(name)\n+    if raw_value is None:\n+        return default\n+    try:\n+        value = int(raw_value)\n+    except ValueError as exc:\n+        raise ValueError(f\"{name} must be an integer, got {raw_value!r}\") from exc\n+    if value \u003c 1:\n+        raise ValueError(f\"{name} must be at least 1\")\n+    return value\n+\n+\n+def _non_negative_float(name: str, default: float) -\u003e float:\n+    raw_value = os.getenv(name)\n+    if raw_value is None:\n+        return default\n+    try:\n+        value = float(raw_value)\n+    except ValueError as exc:\n+        raise ValueError(f\"{name} must be a number, got {raw_value!r}\") from exc\n+    if value \u003c 0:\n+        raise ValueError(f\"{name} cannot be negative\")\n+    return value\n+\n+\n+@dataclass(frozen=True, slots=True)\n+class BlobStorageSettings:\n+    account_url: str\n+    container_name: str = \"blob-manager-demo\"\n+    max_retries: int = 5\n+    retry_delay: float = 1.0\n+    retry_max_delay: float = 30.0\n+    connection_timeout: int = 20\n+    read_timeout: int = 120\n+    max_concurrency: int = 4\n+    block_size: int = 8 * 1024 * 1024\n+    logging_level: str = \"WARNING\"\n+\n+    @classmethod\n+    def from_env(cls) -\u003e \"BlobStorageSettings\":\n+        account_url = os.getenv(\"AZURE_STORAGE_ACCOUNT_URL\", \"\").strip().rstrip(\"/\")\n+        if not account_url:\n+            raise ValueError(\n+                \"AZURE_STORAGE_ACCOUNT_URL is required, for example \"\n+                \"https://\u003caccount\u003e.blob.core.windows.net\"\n+            )\n+        if not account_url.startswith(\"https://\"):\n+            raise ValueError(\"AZURE_STORAGE_ACCOUNT_URL must use HTTPS\")\n+\n+        logging_level = os.getenv(\"AZURE_STORAGE_LOG_LEVEL\", \"WARNING\").upper()\n+        if logging_level not in logging.getLevelNamesMapping():\n+            raise ValueError(f\"Invalid AZURE_STORAGE_LOG_LEVEL: {logging_level!r}\")\n+\n+        return cls(\n+            account_url=account_url,\n+            container_name=os.getenv(\n+                \"AZURE_STORAGE_CONTAINER\", \"blob-manager-demo\"\n+            ).strip(),\n+            max_retries=_positive_int(\"AZURE_STORAGE_MAX_RETRIES\", 5),\n+            retry_delay=_non_negative_float(\"AZURE_STORAGE_RETRY_DELAY\", 1.0),\n+            retry_max_delay=_non_negative_float(\n+                \"AZURE_STORAGE_RETRY_MAX_DELAY\", 30.0\n+            ),\n+            connection_timeout=_positive_int(\n+                \"AZURE_STORAGE_CONNECTION_TIMEOUT\", 20\n+            ),\n+            read_timeout=_positive_int(\"AZURE_STORAGE_READ_TIMEOUT\", 120),\n+            max_concurrency=_positive_int(\"AZURE_STORAGE_MAX_CONCURRENCY\", 4),\n+            block_size=_positive_int(\n+                \"AZURE_STORAGE_BLOCK_SIZE\", 8 * 1024 * 1024\n+            ),\n+            logging_level=logging_level,\n+        )\n+\n+    def configure_logging(self) -\u003e None:\n+        level = logging.getLevelNamesMapping()[self.logging_level]\n+        logging.basicConfig(\n+            level=level,\n+            format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+        )\n+        logging.getLogger(\"azure.core.pipeline.policies.http_logging_policy\").setLevel(\n+            level\n+        )\n+\n+    def retry_policy(self) -\u003e RetryPolicy:\n+        return RetryPolicy(\n+            retry_total=self.max_retries,\n+            retry_connect=self.max_retries,\n+            retry_read=self.max_retries,\n+            retry_status=self.max_retries,\n+            retry_backoff_factor=self.retry_delay,\n+            retry_backoff_max=self.retry_max_delay,\n+            retry_mode=\"exponential\",\n+        )\n+\n+\n+def create_sync_client(\n+    settings: BlobStorageSettings,\n+) -\u003e tuple[BlobServiceClient, DefaultAzureCredential]:\n+    credential = DefaultAzureCredential()\n+    client = BlobServiceClient(\n+        account_url=settings.account_url,\n+        credential=credential,\n+        retry_policy=settings.retry_policy(),\n+        connection_timeout=settings.connection_timeout,\n+        read_timeout=settings.read_timeout,\n+        logging_enable=True,\n+        max_block_size=settings.block_size,\n+        max_single_put_size=settings.block_size,\n+    )\n+    return client, credential\n+\n+\n+def create_async_client(\n+    settings: BlobStorageSettings,\n+) -\u003e tuple[AsyncBlobServiceClient, AsyncDefaultAzureCredential]:\n+    credential = AsyncDefaultAzureCredential()\n+    client = AsyncBlobServiceClient(\n+        account_url=settings.account_url,\n+        credential=credential,\n+        retry_policy=settings.retry_policy(),\n+        connection_timeout=settings.connection_timeout,\n+        read_timeout=settings.read_timeout,\n+        logging_enable=True,\n+        max_block_size=settings.block_size,\n+        max_single_put_size=settings.block_size,\n+    )\n+    return client, credential\n*** Add File: blob_manager/service.py\n+\"\"\"High-level synchronous and asynchronous Azure Blob Storage operations.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+from dataclasses import dataclass\n+from pathlib import Path\n+from typing import Any, AsyncIterator, Generic, TypeVar\n+\n+from azure.core import MatchConditions\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceExistsError,\n+    ResourceNotFoundError,\n+    ServiceRequestError,\n+    ServiceResponseError,\n+)\n+from azure.storage.blob import BlobLeaseClient, BlobServiceClient\n+from azure.storage.blob.aio import BlobLeaseClient as AsyncBlobLeaseClient\n+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient\n+\n+LOGGER = logging.getLogger(__name__)\n+T = TypeVar(\"T\")\n+\n+\n+@dataclass(frozen=True, slots=True)\n+class OperationResult(Generic[T]):\n+    succeeded: bool\n+    message: str\n+    value: T | None = None\n+\n+\n+@dataclass(frozen=True, slots=True)\n+class BlobSummary:\n+    name: str\n+    size: int\n+    etag: str | None\n+    metadata: dict[str, str]\n+    tags: dict[str, str] | None\n+\n+\n+def _failure(operation: str, blob_name: str | None, exc: Exception) -\u003e OperationResult[Any]:\n+    target = f\" for blob {blob_name!r}\" if blob_name else \"\"\n+    if isinstance(exc, ResourceNotFoundError):\n+        detail = \"The blob or container was not found.\"\n+    elif isinstance(exc, ClientAuthenticationError):\n+        detail = \"Authentication failed. Check the managed identity and Azure RBAC role.\"\n+    elif isinstance(exc, ResourceExistsError):\n+        detail = \"Another writer created the blob before this operation completed.\"\n+    elif isinstance(exc, HttpResponseError) and getattr(exc, \"error_code\", None) in {\n+        \"LeaseAlreadyPresent\",\n+        \"LeaseIdMissing\",\n+        \"LeaseIdMismatchWithBlobOperation\",\n+        \"LeaseNotPresentWithBlobOperation\",\n+    }:\n+        detail = \"The blob lease is held by another client or the lease ID is invalid.\"\n+    elif isinstance(exc, HttpResponseError) and exc.status_code in {409, 412}:\n+        detail = (\n+            \"The blob changed concurrently or is protected by another client's lease.\"\n+        )\n+    elif isinstance(exc, HttpResponseError) and exc.status_code == 403:\n+        detail = \"Permission denied. Check the identity's Blob Data RBAC role.\"\n+    elif isinstance(exc, (ServiceRequestError, ServiceResponseError)):\n+        detail = \"Azure Storage could not be reached or returned an invalid response.\"\n+    else:\n+        detail = str(exc) or type(exc).__name__\n+\n+    message = f\"{operation} failed{target}: {detail}\"\n+    LOGGER.warning(message)\n+    return OperationResult(False, message)\n+\n+\n+class BlobStorageService:\n+    def __init__(\n+        self,\n+        client: BlobServiceClient,\n+        container_name: str,\n+        *,\n+        max_concurrency: int = 4,\n+    ) -\u003e None:\n+        self._container = client.get_container_client(container_name)\n+        self._max_concurrency = max_concurrency\n+\n+    def upload(\n+        self,\n+        source: str | Path,\n+        blob_name: str,\n+        *,\n+        metadata: dict[str, str] | None = None,\n+        tags: dict[str, str] | None = None,\n+        timeout: int | None = None,\n+        lease: str | BlobLeaseClient | None = None,\n+    ) -\u003e OperationResult[str]:\n+        path = Path(source)\n+        if not path.is_file():\n+            return OperationResult(False, f\"Upload failed: file not found: {path}\")\n+\n+        blob = self._container.get_blob_client(blob_name)\n+        request_options: dict[str, Any] = {}\n+        if timeout is not None:\n+            request_options[\"timeout\"] = timeout\n+\n+        try:\n+            try:\n+                properties = blob.get_blob_properties(**request_options)\n+            except ResourceNotFoundError:\n+                properties = None\n+\n+            conditions: dict[str, Any]\n+            if properties is None:\n+                conditions = {\"overwrite\": False}\n+            else:\n+                conditions = {\n+                    \"overwrite\": True,\n+                    \"etag\": properties.etag,\n+                    \"match_condition\": MatchConditions.IfNotModified,\n+                }\n+\n+            with path.open(\"rb\") as stream:\n+                response = blob.upload_blob(\n+                    stream,\n+                    length=path.stat().st_size,\n+                    metadata=metadata,\n+                    tags=tags,\n+                    lease=lease,\n+                    max_concurrency=self._max_concurrency,\n+                    **conditions,\n+                    **request_options,\n+                )\n+            return OperationResult(\n+                True,\n+                f\"Uploaded {path} to {blob_name!r}.\",\n+                response.get(\"etag\"),\n+            )\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceExistsError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"Upload\", blob_name, exc)\n+\n+    def download(\n+        self,\n+        blob_name: str,\n+        destination: str | Path,\n+        *,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[Path]:\n+        path = Path(destination)\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            path.parent.mkdir(parents=True, exist_ok=True)\n+            with path.open(\"wb\") as stream:\n+                self._container.download_blob(\n+                    blob_name,\n+                    max_concurrency=self._max_concurrency,\n+                    **request_options,\n+                ).readinto(stream)\n+            return OperationResult(True, f\"Downloaded {blob_name!r} to {path}.\", path)\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            path.unlink(missing_ok=True)\n+            return _failure(\"Download\", blob_name, exc)\n+\n+    def list_blobs(\n+        self,\n+        *,\n+        include_tags: bool = True,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[list[BlobSummary]]:\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            include = [\"metadata\", \"tags\"] if include_tags else [\"metadata\"]\n+            blobs = [\n+                BlobSummary(\n+                    name=item.name,\n+                    size=item.size or 0,\n+                    etag=item.etag,\n+                    metadata=item.metadata or {},\n+                    tags=getattr(item, \"tags\", None),\n+                )\n+                for item in self._container.list_blobs(\n+                    include=include, **request_options\n+                )\n+            ]\n+            return OperationResult(True, f\"Listed {len(blobs)} blob(s).\", blobs)\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"List blobs\", None, exc)\n+\n+    def acquire_lease(\n+        self,\n+        blob_name: str,\n+        *,\n+        duration: int = 60,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[BlobLeaseClient]:\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            lease = BlobLeaseClient(self._container.get_blob_client(blob_name))\n+            lease.acquire(lease_duration=duration, **request_options)\n+            return OperationResult(True, f\"Acquired lease for {blob_name!r}.\", lease)\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"Acquire lease\", blob_name, exc)\n+\n+    def delete(\n+        self,\n+        blob_name: str,\n+        *,\n+        timeout: int | None = None,\n+        lease: str | BlobLeaseClient | None = None,\n+    ) -\u003e OperationResult[None]:\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            self._container.delete_blob(\n+                blob_name,\n+                lease=lease,\n+                delete_snapshots=\"include\",\n+                **request_options,\n+            )\n+            return OperationResult(True, f\"Deleted {blob_name!r}.\")\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"Delete\", blob_name, exc)\n+\n+\n+class AsyncBlobStorageService:\n+    def __init__(\n+        self,\n+        client: AsyncBlobServiceClient,\n+        container_name: str,\n+        *,\n+        max_concurrency: int = 4,\n+        block_size: int = 8 * 1024 * 1024,\n+    ) -\u003e None:\n+        self._container = client.get_container_client(container_name)\n+        self._max_concurrency = max_concurrency\n+        self._block_size = block_size\n+\n+    async def _file_chunks(self, path: Path) -\u003e AsyncIterator[bytes]:\n+        with path.open(\"rb\") as stream:\n+            while chunk := await asyncio.to_thread(stream.read, self._block_size):\n+                yield chunk\n+\n+    async def upload(\n+        self,\n+        source: str | Path,\n+        blob_name: str,\n+        *,\n+        metadata: dict[str, str] | None = None,\n+        tags: dict[str, str] | None = None,\n+        timeout: int | None = None,\n+        lease: str | AsyncBlobLeaseClient | None = None,\n+    ) -\u003e OperationResult[str]:\n+        path = Path(source)\n+        if not path.is_file():\n+            return OperationResult(False, f\"Upload failed: file not found: {path}\")\n+\n+        blob = self._container.get_blob_client(blob_name)\n+        request_options: dict[str, Any] = {}\n+        if timeout is not None:\n+            request_options[\"timeout\"] = timeout\n+\n+        try:\n+            try:\n+                properties = await blob.get_blob_properties(**request_options)\n+            except ResourceNotFoundError:\n+                properties = None\n+\n+            conditions: dict[str, Any]\n+            if properties is None:\n+                conditions = {\"overwrite\": False}\n+            else:\n+                conditions = {\n+                    \"overwrite\": True,\n+                    \"etag\": properties.etag,\n+                    \"match_condition\": MatchConditions.IfNotModified,\n+                }\n+\n+            response = await blob.upload_blob(\n+                self._file_chunks(path),\n+                length=path.stat().st_size,\n+                metadata=metadata,\n+                tags=tags,\n+                lease=lease,\n+                max_concurrency=self._max_concurrency,\n+                **conditions,\n+                **request_options,\n+            )\n+            return OperationResult(\n+                True,\n+                f\"Uploaded {path} to {blob_name!r}.\",\n+                response.get(\"etag\"),\n+            )\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceExistsError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"Upload\", blob_name, exc)\n+\n+    async def download(\n+        self,\n+        blob_name: str,\n+        destination: str | Path,\n+        *,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[Path]:\n+        path = Path(destination)\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            path.parent.mkdir(parents=True, exist_ok=True)\n+            downloader = await self._container.download_blob(\n+                blob_name,\n+                max_concurrency=self._max_concurrency,\n+                **request_options,\n+            )\n+            with path.open(\"wb\") as stream:\n+                async for chunk in downloader.chunks():\n+                    await asyncio.to_thread(stream.write, chunk)\n+            return OperationResult(True, f\"Downloaded {blob_name!r} to {path}.\", path)\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            path.unlink(missing_ok=True)\n+            return _failure(\"Download\", blob_name, exc)\n+\n+    async def list_blobs(\n+        self,\n+        *,\n+        include_tags: bool = True,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[list[BlobSummary]]:\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            include = [\"metadata\", \"tags\"] if include_tags else [\"metadata\"]\n+            blobs = [\n+                BlobSummary(\n+                    name=item.name,\n+                    size=item.size or 0,\n+                    etag=item.etag,\n+                    metadata=item.metadata or {},\n+                    tags=getattr(item, \"tags\", None),\n+                )\n+                async for item in self._container.list_blobs(\n+                    include=include, **request_options\n+                )\n+            ]\n+            return OperationResult(True, f\"Listed {len(blobs)} blob(s).\", blobs)\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"List blobs\", None, exc)\n+\n+    async def acquire_lease(\n+        self,\n+        blob_name: str,\n+        *,\n+        duration: int = 60,\n+        timeout: int | None = None,\n+    ) -\u003e OperationResult[AsyncBlobLeaseClient]:\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            lease = AsyncBlobLeaseClient(\n+                self._container.get_blob_client(blob_name)\n+            )\n+            await lease.acquire(lease_duration=duration, **request_options)\n+            return OperationResult(True, f\"Acquired lease for {blob_name!r}.\", lease)\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"Acquire lease\", blob_name, exc)\n+\n+    async def delete(\n+        self,\n+        blob_name: str,\n+        *,\n+        timeout: int | None = None,\n+        lease: str | AsyncBlobLeaseClient | None = None,\n+    ) -\u003e OperationResult[None]:\n+        request_options = {\"timeout\": timeout} if timeout is not None else {}\n+        try:\n+            await self._container.delete_blob(\n+                blob_name,\n+                lease=lease,\n+                delete_snapshots=\"include\",\n+                **request_options,\n+            )\n+            return OperationResult(True, f\"Deleted {blob_name!r}.\")\n+        except (\n+            ClientAuthenticationError,\n+            HttpResponseError,\n+            ResourceNotFoundError,\n+            ServiceRequestError,\n+            ServiceResponseError,\n+        ) as exc:\n+            return _failure(\"Delete\", blob_name, exc)\n*** Add File: main.py\n+\"\"\"Demonstrate the synchronous and asynchronous blob management services.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+from pathlib import Path\n+\n+from blob_manager import (\n+    AsyncBlobStorageService,\n+    BlobStorageService,\n+    BlobStorageSettings,\n+    create_async_client,\n+    create_sync_client,\n+)\n+\n+SAMPLE_PATH = Path(\"sample.txt\")\n+SYNC_DOWNLOAD_PATH = Path(\"downloads/sync-sample.txt\")\n+ASYNC_DOWNLOAD_PATH = Path(\"downloads/async-sample.txt\")\n+SYNC_BLOB_NAME = \"demo/sync-sample.txt\"\n+ASYNC_BLOB_NAME = \"demo/async-sample.txt\"\n+TIMEOUT_SECONDS = 120\n+TAGS = {\"project\": \"blob-manager\", \"environment\": \"demo\"}\n+\n+\n+def print_result(step: str, result: object) -\u003e None:\n+    message = getattr(result, \"message\", str(result))\n+    succeeded = getattr(result, \"succeeded\", False)\n+    print(f\"[{'OK' if succeeded else 'ERROR'}] {step}: {message}\")\n+\n+\n+def run_sync(settings: BlobStorageSettings) -\u003e None:\n+    print(\"\\n=== Synchronous demo ===\")\n+    client, credential = create_sync_client(settings)\n+    service = BlobStorageService(\n+        client,\n+        settings.container_name,\n+        max_concurrency=settings.max_concurrency,\n+    )\n+    try:\n+        upload = service.upload(\n+            SAMPLE_PATH,\n+            SYNC_BLOB_NAME,\n+            metadata={\"source\": \"sync-demo\"},\n+            tags=TAGS,\n+            timeout=TIMEOUT_SECONDS,\n+        )\n+        print_result(\"upload\", upload)\n+        if not upload.succeeded:\n+            return\n+\n+        listing = service.list_blobs(timeout=TIMEOUT_SECONDS)\n+        print_result(\"list\", listing)\n+        if listing.value:\n+            for blob in listing.value:\n+                print(f\"  - {blob.name} ({blob.size} bytes, tags={blob.tags})\")\n+\n+        download = service.download(\n+            SYNC_BLOB_NAME, SYNC_DOWNLOAD_PATH, timeout=TIMEOUT_SECONDS\n+        )\n+        print_result(\"download\", download)\n+\n+        lease_result = service.acquire_lease(\n+            SYNC_BLOB_NAME, timeout=TIMEOUT_SECONDS\n+        )\n+        print_result(\"acquire lease\", lease_result)\n+        if not lease_result.succeeded or lease_result.value is None:\n+            return\n+\n+        lease = lease_result.value\n+        SAMPLE_PATH.write_text(\"Updated by the synchronous lease holder.\\n\")\n+        overwrite = service.upload(\n+            SAMPLE_PATH,\n+            SYNC_BLOB_NAME,\n+            metadata={\"source\": \"sync-lease-demo\"},\n+            tags=TAGS,\n+            lease=lease,\n+            timeout=TIMEOUT_SECONDS,\n+        )\n+        print_result(\"leased overwrite\", overwrite)\n+\n+        delete = service.delete(\n+            SYNC_BLOB_NAME, lease=lease, timeout=TIMEOUT_SECONDS\n+        )\n+        print_result(\"delete\", delete)\n+    finally:\n+        client.close()\n+        credential.close()\n+\n+\n+async def run_async(settings: BlobStorageSettings) -\u003e None:\n+    print(\"\\n=== Asynchronous demo ===\")\n+    client, credential = create_async_client(settings)\n+    service = AsyncBlobStorageService(\n+        client,\n+        settings.container_name,\n+        max_concurrency=settings.max_concurrency,\n+        block_size=settings.block_size,\n+    )\n+    try:\n+        SAMPLE_PATH.write_text(\"Hello from the asynchronous blob manager.\\n\")\n+        upload = await service.upload(\n+            SAMPLE_PATH,\n+            ASYNC_BLOB_NAME,\n+            metadata={\"source\": \"async-demo\"},\n+            tags=TAGS,\n+            timeout=TIMEOUT_SECONDS,\n+        )\n+        print_result(\"upload\", upload)\n+        if not upload.succeeded:\n+            return\n+\n+        listing = await service.list_blobs(timeout=TIMEOUT_SECONDS)\n+        print_result(\"list\", listing)\n+        if listing.value:\n+            for blob in listing.value:\n+                print(f\"  - {blob.name} ({blob.size} bytes, tags={blob.tags})\")\n+\n+        download = await service.download(\n+            ASYNC_BLOB_NAME, ASYNC_DOWNLOAD_PATH, timeout=TIMEOUT_SECONDS\n+        )\n+        print_result(\"download\", download)\n+\n+        lease_result = await service.acquire_lease(\n+            ASYNC_BLOB_NAME, timeout=TIMEOUT_SECONDS\n+        )\n+        print_result(\"acquire lease\", lease_result)\n+        if not lease_result.succeeded or lease_result.value is None:\n+            return\n+\n+        lease = lease_result.value\n+        SAMPLE_PATH.write_text(\"Updated by the asynchronous lease holder.\\n\")\n+        overwrite = await service.upload(\n+            SAMPLE_PATH,\n+            ASYNC_BLOB_NAME,\n+            metadata={\"source\": \"async-lease-demo\"},\n+            tags=TAGS,\n+            lease=lease,\n+            timeout=TIMEOUT_SECONDS,\n+        )\n+        print_result(\"leased overwrite\", overwrite)\n+\n+        delete = await service.delete(\n+            ASYNC_BLOB_NAME, lease=lease, timeout=TIMEOUT_SECONDS\n+        )\n+        print_result(\"delete\", delete)\n+    finally:\n+        await client.close()\n+        await credential.close()\n+\n+\n+def main() -\u003e None:\n+    settings = BlobStorageSettings.from_env()\n+    settings.configure_logging()\n+    SAMPLE_PATH.write_text(\"Hello from the synchronous blob manager.\\n\")\n+    try:\n+        run_sync(settings)\n+        asyncio.run(run_async(settings))\n+    finally:\n+        SAMPLE_PATH.unlink(missing_ok=True)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.19.0,\u003c2.0.0\n+azure-storage-blob\u003e=12.24.0,\u003c13.0.0\n*** Add File: README.md\n+# Azure Blob Manager\n+\n+A reusable Python utility for synchronous and asynchronous Azure Blob Storage\n+operations. It uses passwordless authentication, streams large transfers, uses\n+ETags to prevent lost updates, and supports exclusive blob leases.\n+\n+## Setup\n+\n+Install the dependencies:\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+Set the endpoint of an existing storage account and container:\n+\n+```powershell\n+$env:AZURE_STORAGE_ACCOUNT_URL = \"https://\u003caccount\u003e.blob.core.windows.net\"\n+$env:AZURE_STORAGE_CONTAINER = \"\u003cexisting-container\u003e\"\n+python main.py\n+```\n+\n+No account key or connection string is accepted. `DefaultAzureCredential`\n+automatically uses the workload's managed identity in Azure. Assign that\n+identity the least-privileged data-plane role it needs, such as **Storage Blob\n+Data Contributor**, at the container scope.\n+\n+## Configuration\n+\n+| Environment variable | Default | Purpose |\n+|---|---:|---|\n+| `AZURE_STORAGE_ACCOUNT_URL` | required | HTTPS Blob service endpoint |\n+| `AZURE_STORAGE_CONTAINER` | `blob-manager-demo` | Existing container |\n+| `AZURE_STORAGE_MAX_RETRIES` | `5` | Maximum exponential retry attempts |\n+| `AZURE_STORAGE_RETRY_DELAY` | `1.0` | Initial retry backoff in seconds |\n+| `AZURE_STORAGE_RETRY_MAX_DELAY` | `30.0` | Maximum retry delay in seconds |\n+| `AZURE_STORAGE_CONNECTION_TIMEOUT` | `20` | Connection timeout in seconds |\n+| `AZURE_STORAGE_READ_TIMEOUT` | `120` | Socket read timeout in seconds |\n+| `AZURE_STORAGE_MAX_CONCURRENCY` | `4` | Parallel transfer workers |\n+| `AZURE_STORAGE_BLOCK_SIZE` | `8388608` | Upload block size in bytes |\n+| `AZURE_STORAGE_LOG_LEVEL` | `WARNING` | Azure HTTP logging level |\n+\n+The optional per-operation `timeout` is passed to Azure Storage as its\n+server-side request timeout. Connection and socket read limits are configured\n+separately on the clients.\n+\n+## Concurrency behavior\n+\n+Uploads first read the current ETag and conditionally replace the blob only if\n+that ETag is still current. New blobs use create-only semantics. A competing\n+writer therefore receives a clear conflict result instead of silently losing\n+its changes. Pass a lease returned by `acquire_lease` when an exclusive update\n+window is required.\n+\n+Large uploads are read in blocks rather than loaded into memory. The async\n+implementation moves local file reads and writes off the event loop.\n+\n+## References\n+\n+- [Azure Blob Storage Python client library](https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme)\n+- [Manage concurrency in Blob Storage](https://learn.microsoft.com/azure/storage/blobs/concurrency-manage)\n+- [Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/__init__.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/__init__.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/__init__.py
@@ -1,0 +1,14 @@
+"""Reusable synchronous and asynchronous Azure Blob Storage utilities."""
+
+from .config import BlobStorageSettings, create_async_client, create_sync_client
+from .service import AsyncBlobStorageService, BlobStorageService, OperationResult
+
+__all__ = [
+    "AsyncBlobStorageService",
+    "BlobStorageService",
+    "BlobStorageSettings",
+    "OperationResult",
+    "create_async_client",
+    "create_sync_client",
+]
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py
@@ -1,0 +1,145 @@
+"""Configuration and authenticated client factories for Azure Blob Storage."""
+
+from __future__ import annotations
+
+import logging
+import os
+from dataclasses import dataclass
+
+from azure.core.pipeline.policies import RetryPolicy
+from azure.identity import DefaultAzureCredential
+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
+from azure.storage.blob import BlobServiceClient
+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient
+
+
+def _
... (truncated)
```

### ✅ 12. powershell — 3003ms

**Input:**

```json
{"command":"python -m compileall -q .; if ($LASTEXITCODE -eq 0) { Write-Output 'compileall passed' }","description":"Compile all Python files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
12.30.1 1.25.3
<shellId: 0 completed with exit code 0>
```

### ✅ 13. powershell — 2614ms

**Input:**

```json
{"command":"python -c \"import azure.storage.blob, azure.identity; print(azure.storage.blob.__version__, azure.identity.__version__)\"","description":"Check installed Azure SDKs","initial_wait":30,"mode":"sync"}
```

**Output:**

```
compileall passed
<shellId: 1 completed with exit code 0>
```

### ✅ 14. powershell — 2315ms

**Input:**

```json
{"command":"@'\nimport inspect\nfrom azure.core.pipeline.policies import RetryPolicy\nfrom azure.storage.blob import BlobClient, BlobLeaseClient, ContainerClient\nfrom azure.storage.blob.aio import BlobClient as AsyncBlobClient, BlobLeaseClient as AsyncBlobLeaseClient, ContainerClient as AsyncContainerClient\nfrom blob_manager.config import BlobStorageSettings\n\nprint('RetryPolicy:', inspect.signature(RetryPolicy))\nprint('sync upload:', inspect.signature(BlobClient.upload_blob))\nprint('async upload:', inspect.signature(AsyncBlobClient.upload_blob))\nprint('sync lease acquire:', inspect.signature(BlobLeaseClient.acquire))\nprint('async lease acquire:', inspect.signature(AsyncBlobLeaseClient.acquire))\nprint('sync list:', inspect.signature(ContainerClient.list_blobs))\nprint('async list:', inspect.signature(AsyncContainerClient.list_blobs))\np = BlobStorageSettings('https://example.blob.core.windows.net').retry_policy()\nprint('retry settings:', p.total_retries, p.backoff_factor, p.backoff_max)\n'@ | python -","description":"Inspect Azure SDK signatures","initial_wait":30,"mode":"sync"}
```

**Output:**

```
RetryPolicy: (**kwargs: Any) -> None
sync upload: (self, data: Union[bytes, str, Iterable[~AnyStr], IO[bytes]], blob_type: Union[str, azure.storage.blob._models.BlobType] = <BlobType.BLOCKBLOB: 'BlockBlob'>, length: Optional[int] = None, metadata: Optional[Dict[str, str]] = None, **kwargs: Any) -> Dict[str, Any]
async upload: (self, data: Union[bytes, str, Iterable[~AnyStr], AsyncIterable[~AnyStr], IO[bytes]], blob_type: Union[str, azure.storage.blob._models.BlobType] = <BlobType.BLOCKBLOB: 'BlockBlob'>, length: Optional[int] = None, metadata: Optional[Dict[str, str]] = None, **kwargs: Any) -> Dict[str, Any]
sync lease acquire: (self, lease_duration: int = -1, **kwargs: Any) -> None
async lease acquire: (self, lease_duration: int = -1, **kwargs: Any) -> None
sync list: (self, name_starts_with: Optional[str] = None, include: Union[str, List[str], NoneType] = None, **kwargs: Any) -> azure.core.paging.ItemPaged[azure.storage.blob._models.BlobProperties]
async list: (self, name_starts_with: Optional[str] = None, include: Union[str, List[str], NoneType] = None, **kwargs: Any) -> azure.core.async_paging.AsyncItemPaged[azure.storage.blob._models.BlobProperties]
retry settings: 5 1.0 30.0
<shellId: 2 completed with exit code 0>
```

### ✅ 15. apply_patch — 639ms

**Input:**

```json
"*** Begin Patch\n*** Update File: blob_manager/config.py\n@@\n-        return cls(\n+        container_name = os.getenv(\n+            \"AZURE_STORAGE_CONTAINER\", \"blob-manager-demo\"\n+        ).strip()\n+        if not container_name:\n+            raise ValueError(\"AZURE_STORAGE_CONTAINER cannot be empty\")\n+\n+        return cls(\n             account_url=account_url,\n-            container_name=os.getenv(\n-                \"AZURE_STORAGE_CONTAINER\", \"blob-manager-demo\"\n-            ).strip(),\n+            container_name=container_name,\n*** Update File: blob_manager/service.py\n@@\n from azure.core.exceptions import (\n+    AzureError,\n     ClientAuthenticationError,\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceExistsError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except (AzureError, OSError) as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except (AzureError, OSError) as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except AzureError as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except AzureError as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except AzureError as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceExistsError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except (AzureError, OSError) as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except (AzureError, OSError) as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except AzureError as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except AzureError as exc:\n@@\n-        except (\n-            ClientAuthenticationError,\n-            HttpResponseError,\n-            ResourceNotFoundError,\n-            ServiceRequestError,\n-            ServiceResponseError,\n-        ) as exc:\n+        except AzureError as exc:\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/config.py
@@ -67,11 +67,15 @@
         if logging_level not in logging.getLevelNamesMapping():
             raise ValueError(f"Invalid AZURE_STORAGE_LOG_LEVEL: {logging_level!r}")
 
+        container_name = os.getenv(
+            "AZURE_STORAGE_CONTAINER", "blob-manager-demo"
+        ).strip()
+        if not container_name:
+            raise ValueError("AZURE_STORAGE_CONTAINER cannot be empty")
+
         return cls(
             account_url=account_url,
-            container_name=os.getenv(
-                "AZURE_STORAGE_CONTAINER", "blob-manager-demo"
-            ).strip(),
+            container_name=container_name,
             max_retries=_positive_int("AZURE_STORAGE_MAX_RETRIES", 5),
             retry_delay=_non_negative_float("AZURE_STORAGE_RETRY_DELAY", 1.0),
             retry_max_delay=_non_negative_float(


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/service.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/service.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-manager-python-azure-skills-azure-skill-mcp-118940352/blob_manager/service.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-ma
... (truncated)
```

### ✅ 16. powershell — 6250ms

**Input:**

```json
{"command":"python -m compileall -q blob_manager main.py; python -m pip check","description":"Compile project and check dependencies","initial_wait":120,"mode":"sync"}
```

**Output:**

```
onnx 1.21.0 requires ml-dtypes, which is not installed.
numba 0.62.1 has requirement numpy<2.4,>=1.22, but you have numpy 2.4.4.
streamlit 1.51.0 has requirement protobuf<7,>=3.20, but you have protobuf 7.34.1.
<shellId: 3 completed with exit code 1>
```

### ✅ 17. powershell — 5809ms

**Input:**

```json
{"command":"@'\nimport asyncio\nimport tempfile\nfrom pathlib import Path\nfrom types import SimpleNamespace\nfrom azure.core.exceptions import ResourceNotFoundError\nfrom blob_manager.service import BlobStorageService, AsyncBlobStorageService\n\nclass Downloader:\n    def readinto(self, stream): stream.write(b'downloaded')\nclass Blob:\n    def __init__(self): self.exists = False; self.kwargs = None\n    def get_blob_properties(self, **kwargs):\n        if not self.exists: raise ResourceNotFoundError('missing')\n        return SimpleNamespace(etag='etag-1')\n    def upload_blob(self, data, **kwargs):\n        assert data.read() == b'payload'; self.kwargs = kwargs; self.exists = True\n        return {'etag': 'etag-2'}\nclass Container:\n    def __init__(self): self.blob = Blob()\n    def get_blob_client(self, name): return self.blob\n    def download_blob(self, name, **kwargs): return Downloader()\n    def list_blobs(self, **kwargs): return [SimpleNamespace(name='x', size=7, etag='e', metadata={}, tags={'a':'b'})]\n    def delete_blob(self, name, **kwargs): self.blob.exists = False\nclass Client:\n    def __init__(self, c): self.c = c\n    def get_container_client(self, name): return self.c\n\nclass AsyncDownloader:\n    async def chunks(self):\n        yield b'down'; yield b'loaded'\nclass AsyncPages:\n    def __aiter__(self): return self\n    async def __anext__(self):\n        if hasattr(self, 'done'): raise StopAsyncIteration\n        self.done = True\n        return SimpleNamespace(name='x', size=7, etag='e', metadata={}, tags={'a':'b'})\nclass AsyncBlob:\n    def __init__(self): self.exists = False; self.kwargs = None\n    async def get_blob_properties(self, **kwargs):\n        if not self.exists: raise ResourceNotFoundError('missing')\n        return SimpleNamespace(etag='etag-1')\n    async def upload_blob(self, data, **kwargs):\n        content = b''\n        async for chunk in data: content += chunk\n        assert content == b'payload'; self.kwargs = kwargs; self.exists = True\n        return {'etag': 'etag-2'}\nclass AsyncContainer:\n    def __init__(self): self.blob = AsyncBlob()\n    def get_blob_client(self, name): return self.blob\n    async def download_blob(self, name, **kwargs): return AsyncDownloader()\n    def list_blobs(self, **kwargs): return AsyncPages()\n    async def delete_blob(self, name, **kwargs): self.blob.exists = False\n\nasync def async_checks(root):\n    c = AsyncContainer(); s = AsyncBlobStorageService(Client(c), 'c', block_size=3)\n    src = root / 'src-a'; dst = root / 'dst-a'; src.write_bytes(b'payload')\n    assert (await s.upload(src, 'x', timeout=9)).succeeded\n    assert c.blob.kwargs['overwrite'] is False and c.blob.kwargs['timeout'] == 9\n    assert (await s.upload(src, 'x')).succeeded\n    assert c.blob.kwargs['etag'] == 'etag-1'\n    assert (await s.download('x', dst)).succeeded and dst.read_bytes() == b'downloaded'\n    assert len((await s.list_blobs()).value) == 1\n    assert (await s.delete('x')).succeeded\n\nwith tempfile.TemporaryDirectory() as d:\n    root = Path(d); src = root / 'src'; dst = root / 'dst'; src.write_bytes(b'payload')\n    c = Container(); s = BlobStorageService(Client(c), 'c')\n    assert s.upload(src, 'x', timeout=9).succeeded\n    assert c.blob.kwargs['overwrite'] is False and c.blob.kwargs['timeout'] == 9\n    assert s.upload(src, 'x').succeeded and c.blob.kwargs['etag'] == 'etag-1'\n    assert s.download('x', dst).succeeded and dst.read_bytes() == b'downloaded'\n    assert len(s.list_blobs().value) == 1\n    assert s.delete('x').succeeded\n    asyncio.run(async_checks(root))\nprint('local sync/async service checks passed')\n'@ | python -","description":"Exercise services with local fakes","initial_wait":120,"mode":"sync"}
```

**Output:**

```
local sync/async service checks passed
<shellId: 4 completed with exit code 0>
```

## Generated Files

- `README.md`
- `blob_manager\__init__.py`
- `blob_manager\config.py`
- `blob_manager\service.py`
- `main.py`
- `requirements.txt`

## Copilot Response

I’m grounding the implementation in the current Azure SDK patterns, then I’ll create and exercise the local project without contacting Azure.

The project will use streamed/block uploads, ETag-based optimistic concurrency, explicit lease support, and structured results so expected Azure errors are reported instead of raised.

The files are in place. I’m now checking the installed SDK’s actual call signatures and exercising the service paths with local fakes so subtle sync/async API mismatches are caught without an Azure connection.

Created the reusable Azure Blob Storage project with:

- `blob_manager/config.py` — environment-based endpoint configuration, `DefaultAzureCredential`, exponential retry policy, transfer tuning, and configurable HTTP logging.
- `blob_manager/service.py` — sync/async streaming upload, download, listing, delete, ETag concurrency protection, lease support, timeouts, and structured error results.
- `main.py` — complete sync and async demonstrations, including tagged uploads and leased overwrites.
- `requirements.txt` and `README.md` — dependencies, configuration, security guidance, and usage instructions.

Set `AZURE_STORAGE_ACCOUNT_URL` and `AZURE_STORAGE_CONTAINER`, then run `python main.py`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. Azure SDK storage failures are consistently caught and converted into clear structured results for both synchronous and asynchronous operations.

### Strengths

- Centralized exception-to-message mapping handles common Azure Storage failure cases.
- Both synchronous and asynchronous service methods consistently return OperationResult failures.
- HttpResponseError status codes and lease-related error codes receive actionable messages.

## Grader Results

- blob-storage-manager.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (10/10)
      - Configures a custom retry policy with exponential backoff: Pass
      - Enables HTTP request/response logging for debugging: Pass
      - Does NOT use connection strings or account keys: Pass
      - Handles large file upload efficiently without loading the entire file into memory: Pass
      - Supports blob index tags on upload (distinct from metadata): Pass
      - Implements blob lease acquisition to prevent concurrent overwrites: Pass
      - Provides both sync and async implementations: Pass
      - Catches and handles storage-specific errors from the Azure SDK: Pass
      - Handles lease conflicts when another client holds a lease: Pass
      - Includes per-operation timeout configuration: Pass
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
hyoka run --prompt-id storage-dp-python-blob-manager --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
