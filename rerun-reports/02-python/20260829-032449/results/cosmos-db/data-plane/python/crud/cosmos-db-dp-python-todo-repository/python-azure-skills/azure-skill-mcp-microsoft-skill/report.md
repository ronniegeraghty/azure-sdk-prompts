# Evaluation Report: cosmos-db-dp-python-todo-repository

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 343.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-python-todo-repository` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 17/20 |
| Duration | 343.3s |
| Timestamp | 2026-08-28T23:31:33Z |
| Files Generated | 0 |
| Event Count | 11044 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 221.8s |
| Review | 120.3s |
| **Total** | **343.3s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-db-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=668210 out=10338 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. skill — 1838ms

**Input:**

```json
{"skill":"azure-cosmos-db-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Cosmos DB Service Implementation

Build production-grade Azure Cosmos DB NoSQL services following clean code, security best practices, and TDD principles.

## Installation

```bash
pip install azure-cosmos azure-identity
```

## Environment Variables

```bash
COSMOS_ENDPOINT=https://<account>.documents.azure.com:443/  # Required for all auth methods
COSMOS_DATABASE_NAME=<database-name>  # Required for all auth methods
COSMOS_CONTAINER_ID=<container-id>  # Required for all auth methods
# For emulator only (not production)
COSMOS_KEY=<emulator-key>  # Only required for key-based auth or emulator
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

**DefaultAzureCredential (preferred)**:
```python
import os
from azure.cosmos import CosmosClient
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<speci
... (truncated)
```

### ✅ 2. glob — 12217ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. view — 17680ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\references\\client-setup.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/error-handling.md
@@ -1,383 +1,383 @@
 # Error Handling Patterns
 
 ## Table of Contents
 
 1. [Cosmos Exception Types](#cosmos-exception-types)
 2. [Client-Level Error Handling](#client-level-error-handling)
 3. [Service-Level Error Handling](#service-level-error-handling)
 4. [Router-Level Error Mapping](#router-level-error-mapping)
 5. [Logging Patterns](#logging-patterns)
 6. [Retry Strategies](#retry-strategies)
 
 ---
 
 ## Cosmos Exception Types
 
 ### Common Exceptions
 
 ```python
 from azure.cosmos.exceptions import (
     CosmosResourceNotFoundError,    # 404 - Document not found
     CosmosResourceExistsError,      # 409 - Conflict (duplicate ID)
     CosmosHttpResponseError,        # Base class for HTTP errors
 )
 ```
 
 | Exception | HTTP Status | Common Cause |
 |-----------|-------------|--------------|
 | `CosmosResourceNotFoundError` | 404 | Document/container doesn't exist |
 | `CosmosResourceExistsError` | 409 | Document with ID already exists |
 | `CosmosHttpResponseError` | 429 | Rate limiting (too many RU) |
 | `CosmosHttpResponseError` | 503 | Service unavailable |
 
 ---
 
 ## Client-Level Error Handling
 
 Handle exceptions in the Cosmos client module:
 
 ```python
 from azure.cosmos.exceptions import CosmosResourceNotFoundError, CosmosHttpResponseError
 import logging
 
 logger = logging.getLogger(__name__)
 
 
 async d
... (truncated)
```

### ✅ 4. view — 16101ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\references\\error-handling.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/partitioning.md
@@ -1,271 +1,271 @@
 # Partition Key Strategies
 
 ## Table of Contents
 
 1. [Partition Key Fundamentals](#partition-key-fundamentals)
 2. [Common Strategies](#common-strategies)
 3. [Cross-Partition Queries](#cross-partition-queries)
 4. [Move Operations](#move-operations)
 5. [Query Optimization](#query-optimization)
 
 ---
 
 ## Partition Key Fundamentals
 
 ### What is a Partition Key?
 
 The partition key determines data distribution and query efficiency:
 
 - **Same partition key** → Data co-located → Fast queries, transactional writes
 - **Different partition keys** → Data distributed → Cross-partition queries required
 
 ### Choosing a Partition Key
 
 Good partition keys have:
 
 1. **High cardinality** — Many distinct values to distribute load
 2. **Even distribution** — No single value dominates storage/throughput
 3. **Query alignment** — Most queries filter by partition key
 
 ---
 
 ## Common Strategies
 
 ### Strategy 1: Self-Partitioned Entities
 
 Use the entity's own ID when entities are accessed individually:
 
 ```python
 # Workspaces partition by their own ID
 class WorkspaceInDB(BaseModel):
     id: str                    # Also used as partition key
     name: str
     doc_type: str = "workspace"
 
 # Query always includes the workspace ID
 doc = await get_document(workspace_id, partition_key=worksp
... (truncated)
```

### ✅ 5. view — 14499ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\references\\partitioning.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/references/client-setup.md
@@ -1,215 +1,215 @@
 # Cosmos DB Client Setup
 
 ## Table of Contents
 
 1. [Dual Authentication Strategy](#dual-authentication-strategy)
 2. [Singleton Pattern](#singleton-pattern)
 3. [Async Wrapping](#async-wrapping)
 4. [Configuration Management](#configuration-management)
 5. [Connection Reset](#connection-reset)
 6. [Complete Implementation](#complete-implementation)
 
 ---
 
 ## Dual Authentication Strategy
 
 Use `DefaultAzureCredential` for Azure deployments and key-based auth only for the local emulator:
 
 ```python
 from azure.cosmos import CosmosClient
 from azure.identity import DefaultAzureCredential
 
 def _is_emulator_endpoint(endpoint: str) -> bool:
     """Detect Cosmos emulator by endpoint URL."""
     return "localhost" in endpoint.lower() or "127.0.0.1" in endpoint
 
 def _create_client(settings) -> CosmosClient:
     if _is_emulator_endpoint(settings.cosmos_endpoint):
         # Emulator: use well-known key, disable SSL verification
         return CosmosClient(
             url=settings.cosmos_endpoint,
             credential=settings.cosmos_key,
             connection_verify=False  # Emulator uses self-signed cert
         )
     else:
         # Azure: use RBAC via DefaultAzureCredential
         credential = DefaultAzureCredential()
         return CosmosClient(
             url=settings.cosmos_endpoint,
   
... (truncated)
```

### ✅ 6. glob — 11097ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-python-todo-repository-python-azure-skills-azure-skill-mcp-microsoft-skill-1401394887","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 13431ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate a local Python project using Azure Cosmos DB NoSQL SDK with sync and async clients, DefaultAzureCredential, CRUD, optimistic concurrency via ETags, paginated parameterized queries, RU charge logging, and database/container initialization without deploying Azure resources.","parameters":{"language":"Python","resource":"Azure Cosmos DB"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 8. azure-get_azure_bestpractices (via azure) — 8411ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a Python Azure Cosmos DB NoSQL repository using Entra authentication, optimistic concurrency, pagination, and RU logging.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 9621ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official current Azure Cosmos DB Python SDK documentation for optimistic concurrency using ETag and MatchConditions.IfNotModified, response_hook request charge headers, query_items pagination by_page, and azure.cosmos.aio async page iteration.","parameters":{"query":"Azure Cosmos DB Python SDK ETag MatchConditions IfNotModified response_hook request charge query_items by_page async azure.cosmos.aio"}}
```

**Output:**

```
{"results":[{"title":"Azure Cosmos DB SQL API client library for Python - version 4.16.3","content":"# Azure Cosmos DB SQL API client library for Python - version 4.16.3\n## Examples\nTransactional Batch operations look very similar to the singular operations apis, and are tuples containing (\u0060operation_type_string\u0060, \u0060args_tuple\u0060, \u0060batch_operation_kwargs_dictionary\u0060), with the kwargs dictionary being optional:\n\u0060\u0060\u0060python\nbatch_operations = [\n        (\u0022create\u0022, (item_body,), kwargs),\n        (\u0022replace\u0022, (item_id, item_body), kwargs),\n        (\u0022read\u0022, (item_id,), kwargs),\n        (\u0022upsert\u0022, (item_body,), kwargs),\n        (\u0022patch\u0022, (item_id, operations), kwargs),\n        (\u0022delete\u0022, (item_id,), kwargs),\n    ]\nbatch_results = container.execute_item_batch(batch_operations=batch_operations, partition_key=partition_key)\n\u0060\u0060\u0060\nThe batch operation kwargs dictionary is limited, and only takes a total of three different key values. In the case of wanting to use conditional patching within the batch, the use of \u0060filter_predicate\u0060 key is available for the patch operation, or in case of wanting to use etags with any of the operations, the use of the \u0060if_match_etag\u0060/\u0060if_none_match_etag\u0060 keys is available as well.\n\u0060\u0060\u0060python\nbatch_operations = [\n        (\u0022replace\u0022, (item_id, item_body), {\u0022if_match_etag\u0022: etag}),\n        (\u0022patch\u0022, (item_id, operations), {\u0022filter_predicate\u0022: filter_predicate, \u0022if_none_match_etag\u0022: etag}),\n    ]\n\u0060\u0060\u0060\nWe also have some samples showing these transactional batch operations in action with both the [sync](https://github.com/Azure/azure-sdk-for-python/tree/azure-cosmos_4.16.3/sdk/cosmos/azure-cosmos/samples/document_management.py) and [async](https://github.com/Azure/azure-sdk-for-python/tree/azure-cosmos_4.16.3/sdk/cos
... (truncated)
```

### ✅ 10. view — 9253ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\assets\\cosmos_client_template.py"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/service_template.py b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/service_template.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/service_template.py
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/service_template.py
@@ -1,315 +1,315 @@
 """
 Service Layer Template
 
 Production-ready service class pattern for Cosmos DB with:
 - Document ↔ Model conversion
 - CRUD operations
 - Graceful degradation
 - Unique slug generation
 
 Usage:
     Rename and customize for your entity type.
 """
 from __future__ import annotations
 
 import uuid
 from datetime import datetime, timezone
 from typing import Any, Optional
 
 from app.db.cosmos import (
     delete_document,
     get_container,
     get_document,
     query_documents,
     upsert_document,
 )
 
 # TODO: Import your Pydantic models
 # from app.models.entity import Entity, EntityCreate, EntityUpdate, EntityInDB
 
 
 def slugify(text: str) -> str:
     """Convert text to URL-friendly slug."""
     import re
 
     slug = text.lower().strip()
     slug = re.sub(r"[^\w\s-]", "", slug)
     slug = re.sub(r"[\s_]+", "-", slug)
     slug = re.sub(r"-+", "-", slug)
     return slug.strip("-")
 
 
 class EntityService:
     """
     Service for Entity CRUD operations.
 
     Replace 'Entity' with your actual entity name (Project, Workspace, etc.)
     """
 
     # -------------------------------------------------------------------------
     # Helper Methods
     # -------------------------------------------------------------------------
 
     def _use_cosmos(self) -> bool:
         """Check if Cosmos D
... (truncated)
```

### ✅ 11. view — 8866ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-cosmos-db-py\\assets\\service_template.py"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-cosmos-db-py/assets/cosmos_client_template.py
@@ -1,226 +1,226 @@
 """
 Cosmos DB Client Module Template
 
 Production-ready Azure Cosmos DB NoSQL client with:
 - Dual authentication (DefaultAzureCredential for Azure, key for emulator)
 - Singleton pattern for connection reuse
 - Async wrapping via run_in_threadpool
 - Graceful error handling
 
 Usage:
     from app.db.cosmos import get_container, upsert_document, get_document
 """
 from __future__ import annotations
 
 import logging
 from typing import Any, Optional
 
 from azure.cosmos import ContainerProxy, CosmosClient
 from azure.cosmos.exceptions import CosmosResourceNotFoundError
 from azure.identity import DefaultAzureCredential
 from starlette.concurrency import run_in_threadpool
 
 from app.config import settings
 
 logger = logging.getLogger(__name__)
 
 # Module-level singleton state
 _cosmos_container: Optional[ContainerProxy] = None
 _credential: Optional[DefaultAzureCredential] = None
 _init_attempted: bool = False
 
 
 def _is_emulator_endpoint(endpoint: str) -> bool:
     """Detect if endpoint is Cosmos emulator."""
     return "localhost" in endpoint.lower() or "127.0.0.1" in endpoint
 
 
 def _create_client() -> CosmosClient:
     """Create Cosmos client with appropriate authentication."""
     global _credential
 
     if _is_emulator_endpoint(settings.cosmos_endpoint):
         logger.i
... (truncated)
```

### ✅ 12. rg — 777ms

**Input:**

```json
{"-i":true,"C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt","pattern":"etag|match_condition|response_hook|by_page|async for|request-charge|x-ms-request-charge|priority"}
```

**Output:**

```
Output too large to read at once (45.2 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787959990611-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt:{"results":[{"title":"Azure Cosmos DB SQL API client library for Python - version 4.16.3","content":"# Azure Cosmos DB SQL API client library for Python - version 4.16.3\n## Examples\nTransactional Batch operations look very similar to the singular operations apis, and are tuples containing (\u0060operation_type_string\u0060, \u0060args_tuple\u0060, \u0060batch_operation_kwargs_dictionary\u0
```

### ✅ 13. powershell — 1370ms

**Input:**

```json
{"command":"$doc = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt' | ConvertFrom-Json; $doc.results | Where-Object { $_.content -match 'etag|match_condition|response_hook|by_page|request-charge' } | ForEach-Object { \"TITLE: $($_.title)`nURL: $($_.contentUrl)`n\" + (($_.content -split \"`n\") | Where-Object { $_ -match 'etag|match_condition|response_hook|by_page|request-charge|async for' } | Select-Object -First 20) + \"`n---\" }","description":"Extract relevant Cosmos SDK documentation snippets"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-92462757c5494046b98e944d2771f8da.txt' | ConvertFrom-Json; $doc.r …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 14. view — 836ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt b/C:/Users/weidxu/AppData/Local/Temp/1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787959986093-copilot-tool-output-92462757c5494046b98e944d2771f8da.txt
@@ -1,8 +1,8 @@
 {"results":[{"title":"Azure Cosmos DB SQL API client library for Python - version 4.16.3","content":"# Azure Cosmos DB SQL API client library for Python - version 4.16.3\n## Examples\nTransactional Batch operations look very similar to the singular operations apis, and are tuples containing (\u0060operation_type_string\u0060, \u0060args_tuple\u0060, \u0060batch_operation_kwargs_dictionary\u0060), with the kwargs dictionary being optional:\n\u0060\u0060\u0060python\nbatch_operations = [\n        (\u0022create\u0022, (item_body,), kwargs),\n        (\u0022replace\u0022, (item_id, item_body), kwargs),\n        (\u0022read\u0022, (item_id,), kwargs),\n        (\u0022upsert\u0022, (item_body,), kwargs),\n        (\u0022patch\u0022, (item_id, operations), kwargs),\n        (\u0022delete\u0022, (item_id,), kwargs),\n    ]\nbatch_results = container.execute_item_batch(batch_operations=batch_operations, partition_key=partition_key)\n\u0060\u0060\u0060\nThe batch operation kwargs dictionary is limited, and only takes a total of three different key values. In the case of wanting to use conditional patching within the batch, the use of \u0060filter_predicate\u0060 key is available for the patch operation, or in case of wanting to use etags with any of the operations, the use of the \u0060if_match_etag\u0060/\u0060if_none_match_etag\u0060 keys is available as well.\n\u0060\u0060\u0060python\nbatch_operations = [\n        (\u0022replace\u0022, (item_id, item_body), {\u0022if_match_etag\u002
... (truncated)
```

### ✅ 15. github-mcp-server-search_code (via github-mcp-server) — 4514ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":5,"query":"repo:Azure/azure-sdk-for-python path:sdk/cosmos/azure-cosmos/azure/cosmos/container.py \"def replace_item\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 16. github-mcp-server-search_code (via github-mcp-server) — 4193ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":5,"query":"repo:Azure/azure-sdk-for-python path:sdk/cosmos/azure-cosmos/azure/cosmos/aio/_container.py \"async def replace_item\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 17. github-mcp-server-search_code (via github-mcp-server) — 3866ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/cosmos/azure-cosmos/samples \"by_page\" \"response_hook\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 18. github-mcp-server-get_file_contents (via github-mcp-server) — 10033ms

**Input:**

```json
{"owner":"Azure","path":"sdk/cosmos/azure-cosmos/azure/cosmos/container.py","ref":"refs/tags/azure-cosmos_4.16.3","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: fb817a3c77472e366d81ba8cb13f332402b9f2b6)# -------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See LICENSE.txt in the project root for
# license information.
# -------------------------------------------------------------------------
from azure.cosmos.aio import CosmosClient
import azure.cosmos.exceptions as exceptions
from azure.cosmos.http_constants import StatusCodes
from azure.cosmos.partition_key import PartitionKey
import datetime

import asyncio
import config

# ----------------------------------------------------------------------------------------------------------
# Prerequisites -
#
# 1. An Azure Cosmos account -
#    https:#azure.microsoft.com/documentation/articles/documentdb-create-account/
#
# 2. Microsoft Azure Cosmos PyPi package -
#    https://pypi.python.org/pypi/azure-cosmos/
# ----------------------------------------------------------------------------------------------------------
# Sample - demonstrates the basic CRUD operations on a Item resource for Azure Cosmos
# ----------------------------------------------------------------------------------------------------------

HOST = config.settings['host']
MASTER_KEY = config.settings['master_key']
DATABASE_ID = config.settings['database_id']
CONTAINER_ID = config.settings['container_id']
CONTAINER_MH_ID = config.settings['container_mh_id']


async def create_items(container):
    print('Creating Items')
    print('\n1.1 Create Item\n')

    # Create a SalesOrder object. This object has nested properties and various types including numbers, DateTimes and strings.
    # This can be saved as JSON as is without converting into rows/columns.
    sales_order = get_sales_order("SalesOrder1")
    await container.create_item(body=sales_order)

    # As your app evolves, let's say your object has a new schema. You can insert SalesOrderV2 objects without
... (truncated)
```

### ✅ 19. github-mcp-server-get_file_contents (via github-mcp-server) — 9695ms

**Input:**

```json
{"owner":"Azure","path":"sdk/cosmos/azure-cosmos/azure/cosmos/aio/_container.py","ref":"refs/tags/azure-cosmos_4.16.3","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: b39acb5f80dbe489fc6721d053f7426c55328f67)# -------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See LICENSE.txt in the project root for
# license information.
# -------------------------------------------------------------------------
import azure.cosmos.cosmos_client as cosmos_client
import azure.cosmos.exceptions as exceptions
from azure.cosmos.http_constants import StatusCodes
from azure.cosmos.partition_key import PartitionKey
import datetime

import config

# ----------------------------------------------------------------------------------------------------------
# Prerequisites -
#
# 1. An Azure Cosmos account -
#    https:#azure.microsoft.com/documentation/articles/documentdb-create-account/
#
# 2. Microsoft Azure Cosmos PyPi package -
#    https://pypi.python.org/pypi/azure-cosmos/
# ----------------------------------------------------------------------------------------------------------
# Sample - demonstrates the basic CRUD operations on a Item resource for Azure Cosmos
# ----------------------------------------------------------------------------------------------------------

HOST = config.settings['host']
MASTER_KEY = config.settings['master_key']
DATABASE_ID = config.settings['database_id']
CONTAINER_ID = config.settings['container_id']
CONTAINER_MH_ID = config.settings['container_mh_id']


def create_items(container):
    print('Creating Items')
    print('\n1.1 Create Item\n')

    # Create a SalesOrder object. This object has nested properties and various types including numbers, DateTimes and strings.
    # This can be saved as JSON as is without converting into rows/columns.
    sales_order = get_sales_order("SalesOrder1")
    container.create_item(body=sales_order)

    # As your app evolves, let's say your object has a new schema. You can insert SalesOrderV2 objects without any
    # changes
... (truncated)
```

### ✅ 20. github-mcp-server-get_file_contents (via github-mcp-server) — 9356ms

**Input:**

```json
{"owner":"Azure","path":"sdk/cosmos/azure-cosmos/samples/document_management.py","ref":"refs/tags/azure-cosmos_4.16.3","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: cc5f5ad7210c6477077ef0269bd9c3b498024192)# The MIT License (MIT)
# Copyright (c) 2014 Microsoft Corporation

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
"""Create, read, update and delete items in the Azure Cosmos DB SQL API service.
"""
import threading
import warnings
from concurrent.futures.thread import ThreadPoolExecutor
from datetime import datetime
from typing import Any, Callable, cast, Iterable, Mapping, Optional, overload, Sequence, Tuple, Union, Dict
from typing_extensions import Literal

from azure.core import MatchConditions
from azure.core.paging import ItemPaged
from azure.core.tracing.decorator import distributed_trace
from azure.cosmos._change_feed.change_feed_utils import add_args_to_kwargs, validate_kwargs

from . import _utils as utils
from ._availability_strategy_config import _validate_request_hedging_strategy
from ._base import (_build_properties_cache, _deserialize_throughput, _replace_throughput, build_options,
 
... (truncated)
```

### ✅ 21. github-mcp-server-get_file_contents (via github-mcp-server) — 9024ms

**Input:**

```json
{"owner":"Azure","path":"sdk/cosmos/azure-cosmos/samples/document_management_async.py","ref":"refs/tags/azure-cosmos_4.16.3","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: aea8af11bcec8434296391bf396b4f86d0985906)# The MIT License (MIT)
# Copyright (c) 2021 Microsoft Corporation

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

"""Create, read, update and delete items in the Azure Cosmos DB SQL API service.
"""
import asyncio  # pylint: disable=do-not-import-asyncio
import warnings
from datetime import datetime
from typing import (Any, Mapping, Optional, Sequence, Union, Tuple, cast, overload, AsyncIterable,
                    Callable, Dict)
from typing_extensions import Literal

from azure.core import MatchConditions
from azure.core.async_paging import AsyncItemPaged, AsyncList
from azure.core.tracing.decorator import distributed_trace
from azure.core.tracing.decorator_async import distributed_trace_async  # type: ignore
from azure.cosmos._change_feed.change_feed_utils import validate_kwargs

from ._cosmos_client_connection_async import CosmosClientConnection
from ._scripts import ScriptsProxy
from .. impor
... (truncated)
```

### ✅ 22. rg — 2663ms

**Input:**

```json
{"-n":true,"A":20,"head_limit":150,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt","pattern":"def (create_item|read_item|replace_item|delete_item|query_items)\\("}
```

**Output:**

```
[grep content: 28 matches across 2 file(s) under C:/Users/weidxu/AppData/Local/Temp]

1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt (14 match(es)):
  114:     item_pages = query_iterable.by_page()
  122:     items_from_continuation = query_iterable.by_page(continuation_token)
  149:     item_pages = query_iterable.by_page()
  183:     item_pages = query_iterable.by_page()
  216:     item_pages = query_iterable.by_page()
  238: async def replace_item_using_etags(container, doc_id):
  240:     # The use of etags and if-match/if-none-match options allows users to run conditional replace operations
  241:     # based on the etag value passed. When using if-match, the request will only succeed if the item's latest etag
  246:     item_etag = read_item["_etag"]
  251:         if_match=item_etag)
  260:         if_none_match="some-etag")
  338:                                        {"if_match_etag": create_response.get_response_headers().get("etag")})
  342:                                             {"if_none_match_etag": create_response.get_response_headers().get("etag")})
  652:             await replace_item_using_etags(container, 'SalesOrder1')

1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt (14 match(es)):
  98:     item_pages = query_iterable.by_page()
  106:     items_from_continuation = query_iterable.by_page(continuation_token)
  132:     item_pages = query_iterable.by_page()
  167:     item_pages = query_iterable.by_page()
  200:     item_pages = query_iterable.by_page()
  228: def replace_item_using_etags(container, doc_id):
  230:     # The use of etags and if-match/if-none-match options allows users to run conditional replace operations
  231:     # based on the etag value passed. When using if-match, the request will only succeed if the item's latest etag
  236:     item_etag = read_item["_etag"]
  241:         if_match=item_etag)
  250:         if_none_match="some-etag")
  326:                              
... (truncated)
```

### ✅ 23. rg — 2189ms

**Input:**

```json
{"-n":true,"A":20,"head_limit":150,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt","pattern":"async def (create_item|read_item|replace_item|delete_item)|def query_items\\("}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt:746:    def query_items(
C:\Users\weidxu\AppData\Local\Temp\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt:848:    def query_items(
C:\Users\weidxu\AppData\Local\Temp\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt:947:    def query_items(  # pylint:disable=docstring-missing-param,too-many-statements
```

### ✅ 24. rg — 1750ms

**Input:**

```json
{"-n":true,"C":4,"head_limit":150,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt"],"pattern":"by_page|response_hook|IfNotModified|match_condition|etag"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:220:    async def create_item(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:320:    async def read_item(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:544:    def query_items(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:641:    def query_items(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:735:    def query_items(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:827:    def query_items(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:1388:    async def replace_item(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:1583:    async def delete_item(
```

### ✅ 25. view — 4691ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt","view_range":[80,115]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt
@@ -1360,91 +1360,91 @@
             )
         return result
 
     @distributed_trace
     def create_item(  # pylint:disable=docstring-missing-param
         self,
         body: dict[str, Any],
         populate_query_metrics: Optional[bool] = None,
         pre_trigger_include: Optional[str] = None,
         post_trigger_include: Optional[str] = None,
         indexing_directive: Optional[int] = None,
         *,
         enable_automatic_id_generation: bool = False,
         session_token: Optional[str] = None,
         initial_headers: Optional[dict[str, str]] = None,
         priority: Optional[Literal["High", "Low"]] = None,
         no_response: Optional[bool] = None,
         retry_write: Optional[int] = None,
         throughput_bucket: Optional[int] = None,
         availability_strategy: Optional[Union[bool, dict[str, Any]]] = None,
         response_hook: Optional[Callable[[Mapping[str, str], dict[str, Any]], None]] = None,
         **kwargs: Any
     ) -> CosmosDict:
         """Create an item in the container.
 
         To update or replace an existing item, use the
         :func:`ContainerProxy.upsert_item` method.
 
         :param body: A dict-like object representing the item to create.
         :type body: dict[str, Any]
         :param str pre_trigger_include: trigger id to be used as pre operation trigger.
         :param str post_trigger_include: trigger id to be used as post operation trigger.
         :param indexing_directive: Enumerates the possible values to ind
... (truncated)
```

### ✅ 26. view — 4286ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt","view_range":[220,255]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt
@@ -1360,96 +1360,96 @@
          * **return_documents** (bool): Whether to return the document text in the response. If False, only scores and indices are returned. Default is True.
          * **top_k** (int): Maximum number of documents to return in the reranked results. If not specified, all documents are returned.
          * **batch_size** (int): Number of documents to process in each batch. Used for optimizing performance with large document sets.
          * **sort** (bool): Whether to sort the results by relevance score in descending order. Default is True.
          * **document_type** (str): Type of documents being reranked. Supported values are "string" and "json".
          * **target_paths** (str): If document_type is "json", the list of JSON paths to extract text from for reranking. Comma-separated string.
 
         :returns: A CosmosDict containing the reranking results. The structure typically includes results list with reranked documents and their relevance scores. Each result contains index, relevance_score, and optionally document.
         :rtype: ~azure.cosmos.CosmosDict[str, Any]
         :raises ~azure.cosmos.exceptions.CosmosHttpResponseError: If the semantic reranking operation fails.
         """
 
         inference_service = self.client_connection._get_inference_service()
         if inference_service is None:
             raise CosmosHttpResponseError(
                 message="Semantic reranking requires AAD credentials (inference service not initialized).",
   
... (truncated)
```

### ✅ 27. view — 3879ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt","view_range":[95,130]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt
@@ -232,34 +232,34 @@
     read_item['subtotal'] = read_item['subtotal'] + 1
     response = await container.replace_item(item=read_item, body=read_item)
 
     print('Replaced Item\'s Id is {0}, new subtotal={1}'.format(response['id'], response['subtotal']))
 
 
 async def replace_item_using_etags(container, doc_id):
     print('\n1.7 Replace an Item using Etags and IfMatch\n')
     # The use of etags and if-match/if-none-match options allows users to run conditional replace operations
     # based on the etag value passed. When using if-match, the request will only succeed if the item's latest etag
     # matches the passed in value. For more on optimistic concurrency control, see the link below:
     # https://learn.microsoft.com/azure/cosmos-db/nosql/database-transactions-optimistic-concurrency
 
     read_item = await container.read_item(item=doc_id, partition_key=doc_id)
     item_etag = read_item["_etag"]
     read_item['subtotal'] = read_item['subtotal'] + 1
     response = await container.replace_item(
         read_item,
         read_item,
         if_match=item_etag)
 
     print('Replaced Item\'s Id is {0}, new subtotal={1}'.format(response['id'], response['subtotal']))
 
     read_item = await container.read_item(item=doc_id, partition_key=doc_id)
     read_item['subtotal'] = read_item['subtotal'] + 1
     response = await container.replace_item(
         read_item,
         read_item,
         if_none_match="some-etag")
 
     print('Replaced Item\'s Id is {0}, new subtotal={1}'
... (truncated)
```

### ✅ 28. view — 3473ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt","view_range":[232,265]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960037611-copilot-tool-output-13623f89f02542dd88ecdc42c87ff7f5.txt
@@ -95,36 +95,36 @@
         parameters=[
             {"name": "@id", "value": doc_id}
         ]
     )
 
     items = [item async for item in query_items_response]
 
     print('Item queried by Id {0}'.format(items[0].get("id")))
 
 
 async def query_items_with_continuation_token(container):
     print('\n1.5 Querying for Items using Pagination and Continuation Token\n')
 
     # When using Pagination, max_item_count will limit the number of items in each page
     query_iterable = container.query_items(
         query="SELECT * FROM r",
         max_item_count=1
     )
 
     item_pages = query_iterable.by_page()
     first_page = await anext(item_pages)  # type: ignore[name-defined]  # cspell:disable-line
     continuation_token = item_pages.continuation_token
 
     # Other code logic where you only need the first page of results would go here
 
     # Now we use the continuation token from the first page to pick up where we left off and
     # access the second page of items
     items_from_continuation = query_iterable.by_page(continuation_token)
     second_page_items_with_continuation = \
         [i async for i in await anext(items_from_continuation)]  # type: ignore[name-defined]  # cspell:disable-line
 
     print('The single items in the second page are {}'.format(second_page_items_with_continuation[0].get("id")))
 
 
 async def query_items_single_partition_with_pagination(container):
     print('\n1.5a Querying with Pagination - Demonstrating max_item_count and Counting Results\
... (truncated)
```

### ✅ 29. view — 2978ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt","view_range":[1360,1450]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt
@@ -220,36 +220,36 @@
 
     read_item = container.read_item(item=doc_id, partition_key=doc_id)
     read_item['subtotal'] = read_item['subtotal'] + 1
     response = container.replace_item(item=read_item, body=read_item)
 
     print('Replaced Item\'s Id is {0}, new subtotal={1}'.format(response['id'], response['subtotal']))
 
 
 def replace_item_using_etags(container, doc_id):
     print('\n1.7 Replace an Item using Etags and IfMatch\n')
     # The use of etags and if-match/if-none-match options allows users to run conditional replace operations
     # based on the etag value passed. When using if-match, the request will only succeed if the item's latest etag
     # matches the passed in value. For more on optimistic concurrency control, see the link below:
     # https://learn.microsoft.com/azure/cosmos-db/nosql/database-transactions-optimistic-concurrency
 
     read_item = container.read_item(item=doc_id, partition_key=doc_id)
     item_etag = read_item["_etag"]
     read_item['subtotal'] = read_item['subtotal'] + 1
     response = container.replace_item(
         read_item,
         read_item,
         if_match=item_etag)
 
     print('Replaced Item\'s Id is {0}, new subtotal={1}'.format(response['id'], response['subtotal']))
 
     read_item = container.read_item(item=doc_id, partition_key=doc_id)
     read_item['subtotal'] = read_item['subtotal'] + 1
     response = container.replace_item(
         read_item,
         read_item,
         if_none_match="some-etag")
 
     print('Replace
... (truncated)
```

### ✅ 30. view — 2556ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt","view_range":[1360,1455]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960037609-copilot-tool-output-f0ad41778e554bdd9e96b73f22784f35.txt
@@ -80,36 +80,36 @@
             {"name": "@id", "value": doc_id}
         ],
         enable_cross_partition_query=True
     ))
 
     print('Item queried by Id {0}'.format(items[0].get("id")))
 
 
 def query_items_with_continuation_token(container):
     print('\n1.5 Querying for Items using Pagination and Continuation Token\n')
 
     # When using Pagination, max_item_count will limit the number of items in each page
     query_iterable = container.query_items(
         query="SELECT * FROM r",
         enable_cross_partition_query=True,
         max_item_count=1
     )
 
     item_pages = query_iterable.by_page()
     first_page = item_pages.next()
     continuation_token = item_pages.continuation_token
 
     # Other code logic where you only need the first page of results would go here
 
     # Now we use the continuation token from the first page to pick up where we left off and
     # access the second page of items
     items_from_continuation = query_iterable.by_page(continuation_token)
     second_page_items_with_continuation = list(items_from_continuation.next())
 
     print('The single items in the second page are {}.'.format(second_page_items_with_continuation[0].get("id")))
 
 
 def query_items_single_partition_with_pagination(container):
     print('\n1.5a Querying with Pagination - Demonstrating max_item_count and Counting Results\n')
     
     # max_item_count controls how many items are returned per page, not the total number of results

```

### ✅ 31. view — 2267ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt","view_range":[730,820]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt
@@ -530,101 +530,101 @@
         query_options = _build_options(kwargs)
         await self._get_properties_with_options(query_options)
         query_options[Constants.ContainerRID] = self.__get_client_container_caches()[self.container_link]["_rid"]
         query_options["enableCrossPartitionQuery"] = True
         query_options[Constants.TimeoutScope] = TimeoutScope.OPERATION
 
         item_tuples = [(item_id, await self._set_partition_key(pk)) for item_id, pk in items]
         return await self.client_connection.read_items(
             collection_link=self.container_link,
             items=item_tuples,
             options= query_options,
             **kwargs)
 
     @overload
     def query_items(
             self,
             query: str,
             *,
             continuation_token_limit: Optional[int] = None,
             enable_scan_in_query: Optional[bool] = None,
             full_text_score_scope: Optional[Literal["Local", "Global"]] = None,
             initial_headers: Optional[dict[str, str]] = None,
             max_integrated_cache_staleness_in_ms: Optional[int] = None,
             max_item_count: Optional[int] = None,
             parameters: Optional[list[dict[str, object]]] = None,
             partition_key: PartitionKeyType,
             populate_index_metrics: Optional[bool] = None,
             populate_query_metrics: Optional[bool] = None,
             populate_query_advice: Optional[bool] = None,
             priority: Optional[Literal["High", "Low"]] = None
... (truncated)
```

### ✅ 32. view — 1917ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt","view_range":[530,630]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt b/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt
@@ -730,91 +730,91 @@
         # populate availability_strategy
         if (Constants.Kwargs.AVAILABILITY_STRATEGY in feed_options
                 and feed_options[Constants.Kwargs.AVAILABILITY_STRATEGY] is not None):
             feed_options[Constants.Kwargs.AVAILABILITY_STRATEGY] =\
                 _validate_request_hedging_strategy(feed_options.pop(Constants.Kwargs.AVAILABILITY_STRATEGY))
 
         response_hook = kwargs.pop("response_hook", None)
         if hasattr(response_hook, "clear"):
             response_hook.clear()
 
         result = self.client_connection.QueryItemsChangeFeed(
             self.container_link, options=feed_options, response_hook=response_hook, **kwargs
         )
         return result
 
     @overload
     def query_items(
             self,
             query: str,
             parameters: Optional[list[dict[str, object]]] = None,
             partition_key: Optional[PartitionKeyType] = None,
             enable_cross_partition_query: Optional[bool] = None,
             max_item_count: Optional[int] = None,
             enable_scan_in_query: Optional[bool] = None,
             populate_query_metrics: Optional[bool] = None,
             *,
             continuation_token_limit: Optional[int] = None,
             full_text_score_scope: Optional[Literal["Local", "Global"]] = None,
             initial_headers: Optional[dict[str, str]] = None,
             max_integrated_cache_staleness_in_ms: Optional[int] = None,
             populate_index_metrics: Option
... (truncated)
```

### ✅ 33. rg — 6299ms

**Input:**

```json
{"-n":true,"A":18,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt","pattern":"def (read_item|replace_item|delete_item)\\("}
```

**Output:**

```
No matches found.
```

### ✅ 34. rg — 2627ms

**Input:**

```json
{"-n":true,"A":18,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787960036407-copilot-tool-output-96ab1944158a44a9883117838d8dab95.txt","pattern":"async def (create_item|read_item|replace_item|delete_item)\\("}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:320:    async def read_item(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:1388:    async def replace_item(
C:\Users\weidxu\AppData\Local\Temp\1787960036410-copilot-tool-output-8ff812e1c23b4a39bca5bcafe8fe6689.txt:1583:    async def delete_item(
```

## Copilot Response

I’m grounding the implementation in the Cosmos SDK’s concurrency, paging, and client-lifecycle requirements before creating the project.

The skill has no separate `acceptance-criteria.md` or `examples.md`; I’m using its supplied references and confirming the current SDK APIs for ETags, response headers, and async page iteration.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Expected Cosmos DB failure modes are handled consistently in synchronous and asynchronous implementations.
- Unexpected SDK HTTP errors preserve their original failure semantics by propagating to the caller.

## Grader Results

- todo-repository.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (11/13)
      - Correct partition key usage: `/category` path, `partition_key` in all point operations: Pass
      - ETag-based optimistic concurrency: captures `_etag` from read response, passes `if_match` on replace: Pass
      - Handles 412 Precondition Failed as a specific error case for conflicts: Pass
      - Parameterized queries using `parameters` list (no f-string or format-string concatenation): Pass
      - Page-by-page iteration using `query_items(...).by_page()` with `max_item_count`: Pass
      - Logs continuation token and item count per page: Fail
      - Async version uses `azure.cosmos.aio.CosmosClient`: Pass
      - TTL configured at 90 days (7776000 seconds) via `default_ttl` in container properties: Pass
      - Indexing policy excludes `/description` path using `excludedPaths`: Pass
      - RU cost extracted from response headers via `x-ms-request-charge` or `response_headers`: Pass
      - Catches `CosmosHttpResponseError` with status code checks (404, 409, 412): Fail
      - Handles 412 separately for ETag conflicts: Pass
      - Does NOT flatten query results by calling `list()` without page iteration: Pass
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
| `Criteria from prompt file` | prompt_review | 85% | 1.00 | 0.8462 | 12.4% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 6.8462** | **85.6%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-python-todo-repository --config python-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
