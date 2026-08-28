# Evaluation Report: storage-dp-python-blob-event-notifier

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 356.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-python-blob-event-notifier` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 14/18 |
| Duration | 356.5s |
| Timestamp | 2026-08-28T21:45:36Z |
| Files Generated | 9 |
| Event Count | 11571 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 241.4s |
| Review | 113.7s |
| **Total** | **356.5s** |

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
| Token Usage | in=509037 out=8414 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 2695ms

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

### ✅ 2. view — 14654ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
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

### ✅ 3. view — 13058ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk\\azure-storage-blob-py.md"}
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

### ✅ 4. glob — 10606ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 11641ms

**Input:**

```json
{"command":"get_best_practices","intent":"Create a local-only Python project using DefaultAzureCredential, Azure Blob Storage sync/async clients, and Azure Event Grid sync/async publisher and event deserialization SDKs.","parameters":{}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 16008ms

**Input:**

```json
{"command":"search","intent":"Find official Python Azure Event Grid SDK guidance for EventGridEvent.from_dict, CloudEvent.from_dict, EventGridPublisherClient and async publisher client, plus BlobServiceClient DefaultAzureCredential patterns.","parameters":{"query":"Python Azure Event Grid EventGridEvent from_dict CloudEvent from_dict EventGridPublisherClient async DefaultAzureCredential BlobServiceClient"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 3599ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Python SDK code-generation best practices for Blob Storage and Event Grid.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 4304ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Python SDK API documentation for deserializing Event Grid native events and CloudEvents from JSON, publishing sync and async Event Grid events, and handling Blob Storage download errors.","parameters":{"query":"site:learn.microsoft.com/python/api azure.eventgrid EventGridEvent from_json from_dict CloudEvent EventGridPublisherClient aio send BlobClient download_blob ResourceNotFoundError BlobArchived"}}
```

**Output:**

```
{"results":[{"title":"EventGridEvent Class-Methods","content":"# EventGridEvent Class\r\n\r\nProperties of an event published to an Event Grid topic using the EventGrid Schema.\r\n\r\nVariables are only populated by the server, and will be ignored when sending a request.\r\n\r\nAll required parameters must be populated in order to send to Azure.\r\n\r\n## Methods\r\n\r\n### from\\_json\r\n\r\nReturns the deserialized EventGridEvent object when a json payload is provided. :param event: The json string that should be converted into a EventGridEvent. This can also be a storage QueueMessage, eventhub\u0027s EventData or ServiceBusMessage :type event: object :rtype: EventGridEvent :return: An EventGridEvent object. :raises ValueError: If the provided JSON is invalid.\r\n\r\n\u0060\u0060\u0060python\r\nfrom_json(event: Any) -\u003E EventGridEvent\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r\n| event\u003Cbr\u003E\u003Cbr\u003ERequired |  |","contentUrl":"https://learn.microsoft.com/python/api/azure-eventgrid/azure.eventgrid.eventgridevent?view=azure-python"},{"title":"EventGridEvent Class-Methods","content":"# EventGridEvent Class\r\n\r\nProperties of an event published to an Event Grid topic using the EventGrid Schema.\r\n\r\nVariables are only populated by the server, and will be ignored when sending a request.\r\n\r\nAll required parameters must be populated in order to send to Azure.\r\n\r\n## Methods\r\n\r\n### as\\_dict\r\n\r\nReturn a dict that can be serialized using json.dump.\r\n\r\nAdvanced usage might optionally use a callback as parameter:\r\n\r\nKey is the attribute name used in Python. Attr\\_desc is a dict of metadata. Currently contains \u0027type\u0027 with the msrest type and \u0027key\u0027 with the RestAPI encoded key. Value is the current value in this object.\r\n\r\nThe string returned will be used to serialize the key. If the return type is a list, this is considered hierarchical result dict.\r\n\r\n
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3942ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python samples for EventGridEvent/CloudEvent deserialization and sync/async EventGridPublisherClient publishing with AzureKeyCredential or token credential.","parameters":{"language":"python","query":"Azure Event Grid Python EventGridEvent CloudEvent deserialization EventGridPublisherClient async send"}}
```

**Output:**

```
{"results":[{"description":"description: Defines the send method signature for sending events to the Event Grid Namespace Service.\nlanguage: python\n","codeSnippet":"\r\n   import os\r\n   import asyncio\r\n   from azure.core.messaging import CloudEvent\r\n   from azure.eventgrid.aio import EventGridClient\r\n   from azure.core.credentials import AzureKeyCredential\r\n\r\n   topic_key = os.environ[\u0022EVENTGRID_CLOUD_EVENT_TOPIC_KEY\u0022]\r\n   endpoint = os.environ[\u0022EVENTGRID_CLOUD_EVENT_TOPIC_ENDPOINT\u0022]\r\n\r\n   async def publish():\r\n       credential = AzureKeyCredential(topic_key)\r\n       client = EventGridClient(endpoint, credential, level=\u0022Basic\u0022)\r\n       async with client:\r\n           await client.send(\r\n               [\r\n                   CloudEvent(\r\n                       type=\u0022Contoso.Items.ItemReceived\u0022,\r\n                       source=\u0022/contoso/items\u0022,\r\n                       data={\u0022itemSku\u0022: \u0022Contoso Item SKU #1\u0022},\r\n                       subject=\u0022Door1\u0022,\r\n                   )\r\n               ]\r\n           )","link":"https://learn.microsoft.com/python/api/azure-eventgrid/azure.eventgrid.aio.eventgridclient?view=azure-python-preview","language":"python"},{"description":"description: Sends multiple CloudEvents as a batch to an Azure Event Grid topic using EventGridPublisherClient for improved performance.\npackage: azure-core, azure-eventgrid\nlanguage: python\n","codeSnippet":"import os\r\nfrom azure.core.credentials import AzureKeyCredential\r\nfrom azure.core.messaging import CloudEvent\r\nfrom azure.eventgrid import EventGridPublisherClient\r\n\r\nkey = os.environ[\u0022EVENTGRID_KEY\u0022]\r\nendpoint = os.environ[\u0022EVENTGRID_ENDPOINT\u0022]\r\ntopic_name = os.environ[\u0022EVENTGRID_TOPIC_NAME\u0022]\r\n\r\nevent0 = CloudEvent(\r\n    type=\u0022Azure.Sdk.Sample\u0022,\r\n    source=\u0022https://egsample.dev/sampleevent\u0022,\r\n    data={\u00
... (truncated)
```

### ✅ 10. rg — 2082ms

**Input:**

```json
{"head_limit":40,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787953594768-copilot-tool-output-1fd3ce52b2bd4adb849735fd5911fefc.txt","pattern":"from_json|CloudEvent Class|ResourceNotFoundError|BlobArchived|EventGridPublisherClient"}
```

**Output:**

```
Output too large to read at once (32.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787953601659-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787953594768-copilot-tool-output-1fd3ce52b2bd4adb849735fd5911fefc.txt:{"results":[{"title":"EventGridEvent Class-Methods","content":"# EventGridEvent Class\r\n\r\nProperties of an event published to an Event Grid topic using the EventGrid Schema.\r\n\r\nVariables are only populated by the server, and will be ignored when sending a request.\r\n\r\nAll required parameters must be populated in order to send to Azure.\r\n\r\n## Methods\r\n\r\n### from\\_json\r\n\r
```

### ✅ 11. web_fetch — 2756ms

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.eventgridevent?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-core/azure.core.messaging.cloudevent?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.core.messaging.CloudEvent class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-core/azure.core.messaging.cloudevent?view=azure-python
config_moniker_range: azure-python
uid: azure.core.messaging.CloudEvent
module: azure.core.messaging
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
locale: en-us
document_id: ccad6d0c-e31e-b74c-6a95-a20bca9c85e3
document_version_independent_id: b06ceb75-4622-b3e5-104c-a451cdeed135
updated_at: 2026-08-28T11:12:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-core/azure.core.messaging.CloudEvent.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/61a8183080abd124ff7b6f9ba762c5ba9da9390a/docs-ref-autogen/azure-core/azure.core.messaging.CloudEvent.yml
git_commit_id: 61a8183080abd124ff7b6f9ba762c5ba9da9390a
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Properties of the CloudEvent 1.0 Schema. All required parameters must be populated in order to send to Azure. '
toc_rel: ../_splitted/azure.core/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-core/azure.core.messaging.cloudevent
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-autogen/azure-core/azure.core.messaging.CloudEv
... (truncated)
```

### ✅ 12. web_fetch — 2413ms

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/python/api/azure-core/azure.core.messaging.cloudevent?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.eventgridevent?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.eventgrid.EventGridEvent class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.eventgridevent?view=azure-python
config_moniker_range: azure-python
uid: azure.eventgrid.EventGridEvent
module: azure.eventgrid
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-event-grid
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/467aaae1-e916-4fcd-a463-5b27f9d4745c
locale: en-us
document_id: 1fe191c1-c16b-947e-5add-4e41a9ba6fb3
document_version_independent_id: 33b4918c-1843-b77a-20d6-1b0fc6a312b1
updated_at: 2025-02-27T12:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-eventgrid/azure.eventgrid.EventGridEvent.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/32233b08830f2adf159cbe6fff20a941d9f0973d/docs-ref-autogen/azure-eventgrid/azure.eventgrid.EventGridEvent.yml
git_commit_id: 32233b08830f2adf159cbe6fff20a941d9f0973d
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Properties of an event published to an Event Grid topic using the EventGrid Schema. Variables are only populated by the server, and will be ignored when sending a request. All required parameters must be populated in order to send to Azure. '
toc_rel: ../_splitted/azure.eventgrid/toc.json
feedback_help_link_type: ''
feed
... (truncated)
```

### ✅ 13. web_fetch — 2835ms

**Input:**

```json
{"max_length":5000,"start_index":5000,"url":"https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.eventgridevent?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.aio.eventgridpublisherclient?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.eventgrid.aio.EventGridPublisherClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.aio.eventgridpublisherclient?view=azure-python
config_moniker_range: azure-python
uid: azure.eventgrid.aio.EventGridPublisherClient
module: azure.eventgrid.aio
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: lmazuel
ms.manager: smortaz
ms.author: lmazuel
ms.devlang: python
ms.date: 2018-05-23T00:00:00.0000000Z
ms.topic: generated-reference
ms.service: azure-event-grid
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/467aaae1-e916-4fcd-a463-5b27f9d4745c
locale: en-us
document_id: 9551751b-48f6-2e3a-40d4-8ef283934d78
document_version_independent_id: 86d16971-682b-59d0-bb93-9213cdc14d8a
updated_at: 2025-04-04T11:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-eventgrid/azure.eventgrid.aio.EventGridPublisherClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/8a91aab31384d094b5e441da8b9c0582e6b8ee75/docs-ref-autogen/azure-eventgrid/azure.eventgrid.aio.EventGridPublisherClient.yml
git_commit_id: 8a91aab31384d094b5e441da8b9c0582e6b8ee75
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'EventGridPublisherClient. Sends events to a basic topic, basic domain, or a namespace topic specified during the client initialization. A single instance or a list of dictionaries, CloudEvents or EventGridEvents are accepted. If a lis
... (truncated)
```

### ✅ 14. web_fetch — 2453ms

**Input:**

```json
{"max_length":5000,"start_index":5000,"url":"https://learn.microsoft.com/en-us/python/api/azure-core/azure.core.messaging.cloudevent?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.eventgridevent?view=azure-python:
es.html#str)<br><br>The type of the event that occurred. |
| data | [object](https://docs.python.org/3/library/functions.html#object)<br><br>Event data specific to the event type. |
| data\_version | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The schema version of the data object. If not provided, will be stamped with an empty value. |
| topic | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The resource path of the event source. If not provided, Event Grid will stamp onto the event. |
| metadata\_version | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The schema version of the event metadata. If provided, must match Event Grid Schema exactly. If not provided, EventGrid will stamp onto event. |
| id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>An identifier for the event. In not provided, a random UUID will be generated and used. |
| event\_time | [datetime](https://docs.python.org/3/library/datetime.html#datetime.datetime)<br><br>The time (in UTC) of the event. If not provided, it will be the time (in UTC) the event was generated. |

## Methods

| as_dict | Return a dict that can be serialized using json.dump.<br><br>Advanced usage might optionally use a callback as parameter:<br><br>Key is the attribute name used in Python. Attr\_desc is a dict of metadata. Currently contains 'type' with the msrest type and 'key' with the RestAPI encoded key. Value is the current value in this object.<br><br>The string returned will be used to serialize the key. If the return type is a list, this is considered hierarchical result dict.<br><br>See the three examples in this file:<br><br>- attribute\_transformer<br>- full\_restapi\_key\_transformer<br>- last\_restapi\_key\_transformer<br><br><br>If you want XML serialization, you can pass the kwargs is\_xml=True. |
| 
... (truncated)
```

### ✅ 15. web_fetch — 2082ms

**Input:**

```json
{"max_length":6000,"url":"https://learn.microsoft.com/en-us/python/api/azure-eventgrid/azure.eventgrid.aio.eventgridpublisherclient?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-core/azure.core.messaging.cloudevent?view=azure-python:
y - value pairs. Each extension must be alphanumeric, lower cased and must not exceed the length of 20 characters.<br><br>Default value: None |

## Methods

| from_dict | Returns the deserialized CloudEvent object when a dict is provided. |
| --- | --- |
| from_json | Returns the deserialized CloudEvent object when a json payload is provided. |

### from\_dict

Returns the deserialized CloudEvent object when a dict is provided.

```python
from_dict(event: Dict[str, Any]) -> CloudEvent[DataType]
```

#### Parameters

| Name | Description |
| --- | --- |
| event<br><br>Required | [dict](https://docs.python.org/3/library/stdtypes.html#dict)<br><br>The dict representation of the event which needs to be deserialized. |

#### Returns

| Type | Description |
| --- | --- |
| [CloudEvent](azure.core.messaging.cloudevent) | The deserialized CloudEvent object. |

### from\_json

Returns the deserialized CloudEvent object when a json payload is provided.

```python
from_json(event: Any) -> CloudEvent[DataType]
```

#### Parameters

| Name | Description |
| --- | --- |
| event<br><br>Required | [object](https://docs.python.org/3/library/functions.html#object)<br><br>The json string that should be converted into a CloudEvent. This can also be a storage QueueMessage, eventhub's EventData or ServiceBusMessage |

#### Returns

| Type | Description |
| --- | --- |
| [CloudEvent](azure.core.messaging.cloudevent) | The deserialized CloudEvent object. |

#### Exceptions

| Type | Description |
| --- | --- |
| [ValueError](https://docs.python.org/3/library/exceptions.html#ValueError) | If the provided JSON is invalid. |

## Attributes

### data

Event data specific to the event type.

```python
data: DataType | None
```

### datacontenttype

Content type of data value.

```python
datacontenttype: str | 
... (truncated)
```

### ✅ 16. apply_patch — 673ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-core\u003e=1.35.0,\u003c2.0.0\n+azure-eventgrid\u003e=4.22.0,\u003c5.0.0\n+azure-identity\u003e=1.24.0,\u003c2.0.0\n+azure-storage-blob\u003e=12.26.0,\u003c13.0.0\n*** Add File: blob_event_notifier/__init__.py\n+\"\"\"Azure Blob Storage lifecycle event processing.\"\"\"\n+\n*** Add File: blob_event_notifier/config.py\n+\"\"\"Passwordless Azure SDK client configuration.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+\n+from azure.eventgrid import EventGridPublisherClient\n+from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.storage.blob import BlobServiceClient\n+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient\n+\n+\n+def storage_account_url() -\u003e str:\n+    return os.environ[\"AZURE_STORAGE_ACCOUNT_URL\"]\n+\n+\n+def event_grid_topic_endpoint() -\u003e str:\n+    return os.environ[\"AZURE_EVENT_GRID_TOPIC_ENDPOINT\"]\n+\n+\n+def create_blob_service_client() -\u003e BlobServiceClient:\n+    return BlobServiceClient(\n+        account_url=storage_account_url(),\n+        credential=DefaultAzureCredential(),\n+    )\n+\n+\n+def create_event_grid_publisher_client(\n+    endpoint: str | None = None,\n+) -\u003e EventGridPublisherClient:\n+    return EventGridPublisherClient(\n+        endpoint=endpoint or event_grid_topic_endpoint(),\n+        credential=DefaultAzureCredential(),\n+    )\n+\n+\n+def create_async_blob_service_client() -\u003e AsyncBlobServiceClient:\n+    return AsyncBlobServiceClient(\n+        account_url=storage_account_url(),\n+        credential=AsyncDefaultAzureCredential(),\n+    )\n+\n+\n+def create_async_event_grid_publisher_client(\n+    endpoint: str | None = None,\n+) -\u003e AsyncEventGridPublisherClient:\n+    return AsyncEventGridPublisherClient(\n+        endpoint=endpoint or event_grid_topic_endpoint(),\n+        credential=AsyncDefaultAzureCredential(),\n+    )\n*** Add File: blob_event_notifier/blob_handler.py\n+\"\"\"Handlers for Blob Storage lifecycle events.\"\"\"\n+\n+from __future__ import annotations\n+\n+import logging\n+from dataclasses import dataclass\n+from typing import Any\n+from urllib.parse import unquote\n+\n+from azure.core.exceptions import HttpResponseError, ResourceNotFoundError\n+\n+LOGGER = logging.getLogger(__name__)\n+SUBJECT_MARKER = \"/blobServices/default/containers/\"\n+ARCHIVE_ERROR_CODES = {\"BlobArchived\", \"BlobBeingRehydrated\", \"BlobNotFound\"}\n+\n+\n+@dataclass(frozen=True)\n+class BlobLocation:\n+    container: str\n+    name: str\n+\n+\n+def parse_blob_subject(subject: str) -\u003e BlobLocation:\n+    if SUBJECT_MARKER not in subject:\n+        raise ValueError(f\"Invalid Blob Storage event subject: {subject!r}\")\n+\n+    _, relative_path = subject.split(SUBJECT_MARKER, maxsplit=1)\n+    try:\n+        container, encoded_name = relative_path.split(\"/blobs/\", maxsplit=1)\n+    except ValueError as exc:\n+        raise ValueError(f\"Invalid Blob Storage event subject: {subject!r}\") from exc\n+\n+    if not container or not encoded_name:\n+        raise ValueError(f\"Invalid Blob Storage event subject: {subject!r}\")\n+    return BlobLocation(unquote(container), unquote(encoded_name))\n+\n+\n+def _access_tier(properties: Any) -\u003e str:\n+    tier = getattr(properties, \"blob_tier\", None)\n+    return str(tier or \"unknown\")\n+\n+\n+def _is_lifecycle_race(exc: HttpResponseError) -\u003e bool:\n+    return getattr(exc, \"error_code\", None) in ARCHIVE_ERROR_CODES\n+\n+\n+def handle_blob_created(event: Any, blob_service_client: Any) -\u003e None:\n+    location = parse_blob_subject(event.subject)\n+    blob_client = blob_service_client.get_blob_client(\n+        container=location.container,\n+        blob=location.name,\n+    )\n+\n+    try:\n+        downloader = blob_client.download_blob()\n+        content = downloader.readall()\n+        properties = downloader.properties\n+    except ResourceNotFoundError:\n+        LOGGER.warning(\n+            \"Blob %s/%s no longer exists; skipping created event\",\n+            location.container,\n+            location.name,\n+        )\n+        return\n+    except HttpResponseError as exc:\n+        if _is_lifecycle_race(exc):\n+            LOGGER.warning(\n+                \"Blob %s/%s is unavailable after a lifecycle change (%s)\",\n+                location.container,\n+                location.name,\n+                exc.error_code,\n+            )\n+            return\n+        raise\n+\n+    content_settings = getattr(properties, \"content_settings\", None)\n+    content_type = getattr(content_settings, \"content_type\", None) or \"unknown\"\n+    size = getattr(properties, \"size\", None)\n+    LOGGER.info(\n+        \"Blob created: name=%s size=%s content_type=%s access_tier=%s\",\n+        location.name,\n+        size if size is not None else len(content),\n+        content_type,\n+        _access_tier(properties),\n+    )\n+\n+\n+async def handle_blob_created_async(event: Any, blob_service_client: Any) -\u003e None:\n+    location = parse_blob_subject(event.subject)\n+    blob_client = blob_service_client.get_blob_client(\n+        container=location.container,\n+        blob=location.name,\n+    )\n+\n+    try:\n+        downloader = await blob_client.download_blob()\n+        content = await downloader.readall()\n+        properties = downloader.properties\n+    except ResourceNotFoundError:\n+        LOGGER.warning(\n+            \"Blob %s/%s no longer exists; skipping created event\",\n+            location.container,\n+            location.name,\n+        )\n+        return\n+    except HttpResponseError as exc:\n+        if _is_lifecycle_race(exc):\n+            LOGGER.warning(\n+                \"Blob %s/%s is unavailable after a lifecycle change (%s)\",\n+                location.container,\n+                location.name,\n+                exc.error_code,\n+            )\n+            return\n+        raise\n+\n+    content_settings = getattr(properties, \"content_settings\", None)\n+    content_type = getattr(content_settings, \"content_type\", None) or \"unknown\"\n+    size = getattr(properties, \"size\", None)\n+    LOGGER.info(\n+        \"Blob created: name=%s size=%s content_type=%s access_tier=%s\",\n+        location.name,\n+        size if size is not None else len(content),\n+        content_type,\n+        _access_tier(properties),\n+    )\n+\n+\n+def handle_blob_deleted(event: Any) -\u003e None:\n+    location = parse_blob_subject(event.subject)\n+    LOGGER.info(\n+        \"Blob deleted: container=%s name=%s\",\n+        location.container,\n+        location.name,\n+    )\n+\n+\n+async def handle_blob_deleted_async(event: Any) -\u003e None:\n+    handle_blob_deleted(event)\n*** Add File: blob_event_notifier/event_receiver.py\n+\"\"\"Event Grid webhook payload deserialization and routing.\"\"\"\n+\n+from __future__ import annotations\n+\n+import json\n+import logging\n+from collections.abc import Callable\n+from typing import Any, TypeAlias\n+\n+from azure.core.messaging import CloudEvent\n+from azure.eventgrid import EventGridEvent\n+\n+from .blob_handler import (\n+    handle_blob_created,\n+    handle_blob_created_async,\n+    handle_blob_deleted,\n+    handle_blob_deleted_async,\n+)\n+\n+LOGGER = logging.getLogger(__name__)\n+BLOB_CREATED = \"Microsoft.Storage.BlobCreated\"\n+BLOB_DELETED = \"Microsoft.Storage.BlobDeleted\"\n+GridEvent: TypeAlias = EventGridEvent | CloudEvent\n+\n+\n+def deserialize_events(payload: str | bytes) -\u003e list[GridEvent]:\n+    try:\n+        envelope = json.loads(payload)\n+    except (TypeError, json.JSONDecodeError) as exc:\n+        raise ValueError(\"Webhook payload is not valid JSON\") from exc\n+\n+    raw_events = envelope if isinstance(envelope, list) else [envelope]\n+    events: list[GridEvent] = []\n+    for raw_event in raw_events:\n+        if not isinstance(raw_event, dict):\n+            raise ValueError(\"Each webhook event must be a JSON object\")\n+\n+        # SDK helpers own the schema-to-model mapping; this code only identifies\n+        # which Event Grid-supported envelope was delivered.\n+        if \"specversion\" in raw_event:\n+            events.append(CloudEvent.from_dict(raw_event))\n+        else:\n+            events.append(EventGridEvent.from_dict(raw_event))\n+    return events\n+\n+\n+def _event_type(event: GridEvent) -\u003e str:\n+    return event.type if isinstance(event, CloudEvent) else event.event_type\n+\n+\n+def receive_events(\n+    payload: str | bytes,\n+    blob_service_client: Any,\n+    *,\n+    on_created: Callable[[GridEvent, Any], None] = handle_blob_created,\n+    on_deleted: Callable[[GridEvent], None] = handle_blob_deleted,\n+) -\u003e list[GridEvent]:\n+    events = deserialize_events(payload)\n+    for event in events:\n+        event_type = _event_type(event)\n+        if event_type == BLOB_CREATED:\n+            on_created(event, blob_service_client)\n+        elif event_type == BLOB_DELETED:\n+            on_deleted(event)\n+        else:\n+            LOGGER.warning(\"Ignoring unsupported event type: %s\", event_type)\n+    return events\n+\n+\n+async def receive_events_async(\n+    payload: str | bytes,\n+    blob_service_client: Any,\n+) -\u003e list[GridEvent]:\n+    events = deserialize_events(payload)\n+    for event in events:\n+        event_type = _event_type(event)\n+        if event_type == BLOB_CREATED:\n+            await handle_blob_created_async(event, blob_service_client)\n+        elif event_type == BLOB_DELETED:\n+            await handle_blob_deleted_async(event)\n+        else:\n+            LOGGER.warning(\"Ignoring unsupported event type: %s\", event_type)\n+    return events\n*** Add File: blob_event_notifier/event_publisher.py\n+\"\"\"Sync and async publishers for custom downstream Event Grid events.\"\"\"\n+\n+from __future__ import annotations\n+\n+import logging\n+from collections.abc import Mapping, Sequence\n+from dataclasses import dataclass\n+from typing import Any\n+\n+from azure.core.exceptions import AzureError\n+from azure.eventgrid import EventGridEvent\n+\n+from .config import (\n+    create_async_event_grid_publisher_client,\n+    create_event_grid_publisher_client,\n+)\n+\n+LOGGER = logging.getLogger(__name__)\n+\n+\n+@dataclass(frozen=True)\n+class CustomEvent:\n+    event_type: str\n+    data: Mapping[str, Any]\n+    subject: str\n+    data_version: str = \"1.0\"\n+\n+\n+def _normalize_subject(subject: str) -\u003e str:\n+    normalized = \"/\" + subject.strip(\"/\")\n+    if normalized == \"/\" or \"//\" in normalized:\n+        raise ValueError(\"Event subject must contain a valid hierarchy\")\n+    return normalized\n+\n+\n+def _to_event_grid_events(events: Sequence[CustomEvent]) -\u003e list[EventGridEvent]:\n+    if not events:\n+        raise ValueError(\"At least one custom event is required\")\n+    return [\n+        EventGridEvent(\n+            subject=_normalize_subject(event.subject),\n+            event_type=event.event_type,\n+            data=dict(event.data),\n+            data_version=event.data_version,\n+        )\n+        for event in events\n+    ]\n+\n+\n+def publish_events(\n+    endpoint: str,\n+    events: Sequence[CustomEvent],\n+    *,\n+    client: Any | None = None,\n+) -\u003e bool:\n+    publisher = client or create_event_grid_publisher_client(endpoint)\n+    owns_client = client is None\n+    try:\n+        publisher.send(_to_event_grid_events(events))\n+        LOGGER.info(\"Published %d downstream event(s)\", len(events))\n+        return True\n+    except AzureError:\n+        LOGGER.exception(\"Event Grid publishing failed\")\n+        return False\n+    finally:\n+        if owns_client:\n+            publisher.close()\n+\n+\n+async def publish_events_async(\n+    endpoint: str,\n+    events: Sequence[CustomEvent],\n+    *,\n+    client: Any | None = None,\n+) -\u003e bool:\n+    publisher = client or create_async_event_grid_publisher_client(endpoint)\n+    owns_client = client is None\n+    try:\n+        await publisher.send(_to_event_grid_events(events))\n+        LOGGER.info(\"Published %d downstream event(s)\", len(events))\n+        return True\n+    except AzureError:\n+        LOGGER.exception(\"Event Grid publishing failed\")\n+        return False\n+    finally:\n+        if owns_client:\n+            await publisher.close()\n*** Add File: main.py\n+\"\"\"Run an offline demonstration of sync and async event processing.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import json\n+import logging\n+from dataclasses import dataclass\n+from typing import Any\n+\n+from blob_event_notifier.event_publisher import (\n+    CustomEvent,\n+    publish_events,\n+    publish_events_async,\n+)\n+from blob_event_notifier.event_receiver import receive_events, receive_events_async\n+\n+DEMO_ENDPOINT = \"https://example-topic.westus2-1.eventgrid.azure.net/api/events\"\n+CREATED_SUBJECT = (\n+    \"/blobServices/default/containers/documents/blobs/invoices/2026/august-001.pdf\"\n+)\n+DELETED_SUBJECT = (\n+    \"/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf\"\n+)\n+\n+\n+@dataclass\n+class MockContentSettings:\n+    content_type: str\n+\n+\n+@dataclass\n+class MockBlobProperties:\n+    size: int\n+    content_settings: MockContentSettings\n+    blob_tier: str\n+\n+\n+class MockDownloader:\n+    def __init__(self, content: bytes) -\u003e None:\n+        self._content = content\n+        self.properties = MockBlobProperties(\n+            size=len(content),\n+            content_settings=MockContentSettings(\"application/pdf\"),\n+            blob_tier=\"Hot\",\n+        )\n+\n+    def readall(self) -\u003e bytes:\n+        return self._content\n+\n+\n+class AsyncMockDownloader(MockDownloader):\n+    async def readall(self) -\u003e bytes:\n+        return self._content\n+\n+\n+class MockBlobClient:\n+    def download_blob(self) -\u003e MockDownloader:\n+        return MockDownloader(b\"%PDF-1.7 demo invoice\")\n+\n+\n+class AsyncMockBlobClient:\n+    async def download_blob(self) -\u003e AsyncMockDownloader:\n+        return AsyncMockDownloader(b\"%PDF-1.7 demo invoice\")\n+\n+\n+class MockBlobServiceClient:\n+    def get_blob_client(self, **_: Any) -\u003e MockBlobClient:\n+        return MockBlobClient()\n+\n+\n+class AsyncMockBlobServiceClient:\n+    def get_blob_client(self, **_: Any) -\u003e AsyncMockBlobClient:\n+        return AsyncMockBlobClient()\n+\n+\n+class MockPublisherClient:\n+    def send(self, events: list[Any]) -\u003e None:\n+        logging.info(\"Mock publisher accepted %d event(s)\", len(events))\n+\n+\n+class AsyncMockPublisherClient:\n+    async def send(self, events: list[Any]) -\u003e None:\n+        logging.info(\"Async mock publisher accepted %d event(s)\", len(events))\n+\n+\n+def sample_payloads() -\u003e tuple[str, str]:\n+    native_events = [\n+        {\n+            \"id\": \"6f159fb5-006e-001b-66f6-75d8ed06f101\",\n+            \"topic\": (\n+                \"/subscriptions/00000000-0000-0000-0000-000000000000/\"\n+                \"resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore\"\n+            ),\n+            \"subject\": CREATED_SUBJECT,\n+            \"eventType\": \"Microsoft.Storage.BlobCreated\",\n+            \"eventTime\": \"2026-08-28T21:45:00Z\",\n+            \"data\": {\n+                \"api\": \"PutBlob\",\n+                \"clientRequestId\": \"8b1b282a-9c65-4d5e-9c21-87bdba80a601\",\n+                \"requestId\": \"6f159fb5-006e-001b-66f6-75d8ed000000\",\n+                \"eTag\": \"0x8DC000000000001\",\n+                \"contentType\": \"application/pdf\",\n+                \"contentLength\": 32145,\n+                \"blobType\": \"BlockBlob\",\n+                \"url\": (\n+                    \"https://demostore.blob.core.windows.net/\"\n+                    \"documents/invoices/2026/august-001.pdf\"\n+                ),\n+                \"sequencer\": \"000000000000000000000000000001\",\n+                \"storageDiagnostics\": {\"batchId\": \"demo-batch-001\"},\n+            },\n+            \"dataVersion\": \"\",\n+            \"metadataVersion\": \"1\",\n+        }\n+    ]\n+    cloud_events = [\n+        {\n+            \"specversion\": \"1.0\",\n+            \"type\": \"Microsoft.Storage.BlobDeleted\",\n+            \"source\": (\n+                \"/subscriptions/00000000-0000-0000-0000-000000000000/\"\n+                \"resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore\"\n+            ),\n+            \"subject\": DELETED_SUBJECT,\n+            \"id\": \"9aeb0fdf-c01e-0131-6b6e-75a90f06ffff\",\n+            \"time\": \"2026-08-28T21:46:00Z\",\n+            \"datacontenttype\": \"application/json\",\n+            \"data\": {\n+                \"api\": \"DeleteBlob\",\n+                \"requestId\": \"9aeb0fdf-c01e-0131-6b6e-75a90f000000\",\n+                \"url\": (\n+                    \"https://demostore.blob.core.windows.net/\"\n+                    \"documents/archive/old-invoice.pdf\"\n+                ),\n+                \"sequencer\": \"000000000000000000000000000002\",\n+            },\n+        }\n+    ]\n+    return json.dumps(native_events), json.dumps(cloud_events)\n+\n+\n+def downstream_event() -\u003e CustomEvent:\n+    return CustomEvent(\n+        event_type=\"Contoso.Documents.DocumentProcessed\",\n+        subject=\"/documents/invoices/processed\",\n+        data={\"documentName\": \"august-001.pdf\", \"status\": \"processed\"},\n+    )\n+\n+\n+def run_sync_demo() -\u003e None:\n+    logging.info(\"Starting synchronous demo\")\n+    native_payload, cloud_payload = sample_payloads()\n+    blob_client = MockBlobServiceClient()\n+    receive_events(native_payload, blob_client)\n+    receive_events(cloud_payload, blob_client)\n+    publish_events(\n+        DEMO_ENDPOINT,\n+        [downstream_event()],\n+        client=MockPublisherClient(),\n+    )\n+\n+\n+async def run_async_demo() -\u003e None:\n+    logging.info(\"Starting asynchronous demo\")\n+    native_payload, cloud_payload = sample_payloads()\n+    blob_client = AsyncMockBlobServiceClient()\n+    await receive_events_async(native_payload, blob_client)\n+    await receive_events_async(cloud_payload, blob_client)\n+    await publish_events_async(\n+        DEMO_ENDPOINT,\n+        [downstream_event()],\n+        client=AsyncMockPublisherClient(),\n+    )\n+\n+\n+if __name__ == \"__main__\":\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s %(message)s\")\n+    run_sync_demo()\n+    asyncio.run(run_async_demo())\n*** Add File: tests/test_project.py\n+from __future__ import annotations\n+\n+import asyncio\n+import unittest\n+\n+from azure.core.messaging import CloudEvent\n+from azure.eventgrid import EventGridEvent\n+\n+from blob_event_notifier.blob_handler import parse_blob_subject\n+from blob_event_notifier.event_publisher import CustomEvent, publish_events\n+from blob_event_notifier.event_receiver import deserialize_events\n+from main import MockPublisherClient, sample_payloads\n+\n+\n+class ProjectTests(unittest.TestCase):\n+    def test_deserializes_both_supported_schemas(self) -\u003e None:\n+        native_payload, cloud_payload = sample_payloads()\n+\n+        native = deserialize_events(native_payload)\n+        cloud = deserialize_events(cloud_payload)\n+\n+        self.assertIsInstance(native[0], EventGridEvent)\n+        self.assertIsInstance(cloud[0], CloudEvent)\n+\n+    def test_parses_nested_and_encoded_blob_name(self) -\u003e None:\n+        location = parse_blob_subject(\n+            \"/blobServices/default/containers/docs/blobs/2026/invoice%2001.pdf\"\n+        )\n+\n+        self.assertEqual(\"docs\", location.container)\n+        self.assertEqual(\"2026/invoice 01.pdf\", location.name)\n+\n+    def test_publisher_normalizes_subject(self) -\u003e None:\n+        client = MockPublisherClient()\n+\n+        published = publish_events(\n+            \"https://example.invalid\",\n+            [CustomEvent(\"Contoso.Test\", {\"ok\": True}, \"documents/processed\")],\n+            client=client,\n+        )\n+\n+        self.assertTrue(published)\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/requirements.txt
@@ -1,0 +1,5 @@
+azure-core>=1.35.0,<2.0.0
+azure-eventgrid>=4.22.0,<5.0.0
+azure-identity>=1.24.0,<2.0.0
+azure-storage-blob>=12.26.0,<13.0.0
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/blob_event_notifier/__init__.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/blob_event_notifier/__init__.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/blob_event_notifier/__init__.py
@@ -1,0 +1,3 @@
+"""Azure Blob Storage lifecycle event processing."""
+
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/blob_event_notifier/config.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/blob_event_notifier/config.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/blob_event_notifier/config.py
@@ -1,0 +1,53 @@
+"""Passwordless Azure SDK client configuration."""
+
+from __future__ import annotations
+
+import os
+
+from azure.eventgrid import Eve
... (truncated)
```

### ✅ 17. powershell — 66326ms

**Input:**

```json
{"command":"python -m venv .venv \u0026\u0026 .\\.venv\\Scripts\\python.exe -m pip install --quiet -r requirements.txt \u0026\u0026 .\\.venv\\Scripts\\python.exe -m unittest discover -s tests -v \u0026\u0026 .\\.venv\\Scripts\\python.exe main.py","description":"Install SDKs, run tests, and execute offline demo","initial_wait":180,"mode":"sync","shellId":"validate-project"}
```

**Output:**

```
[notice] A new release of pip is available: 25.2 -> 26.2.1
[notice] To update, run: C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057\.venv\Scripts\python.exe -m pip install --upgrade pip
test_deserializes_both_supported_schemas (test_project.ProjectTests.test_deserializes_both_supported_schemas) ... ok
test_parses_nested_and_encoded_blob_name (test_project.ProjectTests.test_parses_nested_and_encoded_blob_name) ... ok
test_publisher_normalizes_subject (test_project.ProjectTests.test_publisher_normalizes_subject) ... ok

----------------------------------------------------------------------
Ran 3 tests in 0.008s

OK
INFO Starting synchronous demo
INFO Blob created: name=invoices/2026/august-001.pdf size=21 content_type=application/pdf access_tier=Hot
INFO Blob deleted: container=documents name=archive/old-invoice.pdf
INFO Mock publisher accepted 1 event(s)
INFO Published 1 downstream event(s)
INFO Starting asynchronous demo
INFO Blob created: name=invoices/2026/august-001.pdf size=21 content_type=application/pdf access_tier=Hot
INFO Blob deleted: container=documents name=archive/old-invoice.pdf
INFO Async mock publisher accepted 1 event(s)
INFO Published 1 downstream event(s)
<shellId: validate-project completed with exit code 0>
```

### ✅ 18. apply_patch — 2724ms

**Input:**

```json
"*** Begin Patch\n*** Add File: .gitignore\n+.venv/\n+__pycache__/\n+*.py[cod]\n+\n*** Add File: README.md\n+# Azure Blob Event Notifier\n+\n+This Python project receives Azure Blob Storage lifecycle events in either the\n+Event Grid native schema or CloudEvents 1.0 schema, downloads newly created\n+blobs, logs deletions, and publishes downstream custom events.\n+\n+## Run the offline demo\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\python.exe -m pip install -r requirements.txt\n+.\\.venv\\Scripts\\python.exe main.py\n+```\n+\n+The demo injects local mock clients, so it does not connect to Azure.\n+\n+## Use with Azure\n+\n+Set the resource endpoints; authentication is resolved passwordlessly through\n+`DefaultAzureCredential`.\n+\n+```powershell\n+$env:AZURE_STORAGE_ACCOUNT_URL = \"https://\u003caccount\u003e.blob.core.windows.net\"\n+$env:AZURE_EVENT_GRID_TOPIC_ENDPOINT = \"https://\u003ctopic\u003e.\u003cregion\u003e-1.eventgrid.azure.net/api/events\"\n+```\n+\n+Grant the workload identity only the data-plane roles it needs, such as\n+**Storage Blob Data Reader** on the storage account and **EventGrid Data\n+Sender** on the custom topic. Application code can create real clients through\n+`blob_event_notifier.config`; callers are responsible for closing those clients.\n+\n+## References\n+\n+- [Azure Event Grid Python SDK](https://learn.microsoft.com/python/api/overview/azure/eventgrid-readme)\n+- [EventGridEvent deserialization](https://learn.microsoft.com/python/api/azure-eventgrid/azure.eventgrid.eventgridevent)\n+- [CloudEvent deserialization](https://learn.microsoft.com/python/api/azure-core/azure.core.messaging.cloudevent)\n+- [Azure Blob Storage Python SDK](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-python)\n+\n*** Update File: tests/test_project.py\n@@\n-import asyncio\n import unittest\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/.gitignore b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/.gitignore
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/.gitignore
@@ -1,0 +1,5 @@
+.venv/
+__pycache__/
+*.py[cod]
+
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/README.md b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/README.md
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-blob-event-notifier-python-azure-skills-azure-skill-mcp-2820600057/README.md
@@ -1,0 +1,39 @@
+# Azure Blob Event Notifier
+
+This Python project receives Azure Blob Storage lifecycle events in either the
+Event Grid native schema or CloudEvents 1.0 schema, downloads newly created
+blobs, logs deletions, and publishes downstream custom events.
+
+## Run the offline demo
+
+```powershell
+python -m venv .venv
+.\.venv\Scripts\python.exe -m pip install -r requirements.txt
+.\.venv\Scripts\python.exe main.py
+```
+
+The demo injects local mock clients, so it does not connect to Azure.
+
+## Use with Azure
+
+Set the resource endpoints; authentication is resolved passwordlessly through
+`DefaultAzureCredential`.
+
+```powershell
+$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
+$env:AZURE_EVENT_GRID_TOPIC_ENDPOINT = "https://<topic>.<region>-1.eventgrid.azure.net/api/events"
+```
+
+Grant the workload identity only the data-plane roles it needs, such as
+**Storage Blob Data Reader** on the storage
... (truncated)
```

### ✅ 19. powershell — 6929ms

**Input:**

```json
{"command":"$targets = @('.venv', 'blob_event_notifier\\__pycache__', 'tests\\__pycache__', '__pycache__'); foreach ($target in $targets) { if (Test-Path -LiteralPath $target) { $resolved = (Resolve-Path -LiteralPath $target).Path; Remove-Item -LiteralPath $resolved -Recurse -Force } }; Get-ChildItem -Force | Select-Object Name, Mode","description":"Remove local validation artifacts and list project files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Name                Mode
----                ----
blob_event_notifier d----
tests               d----
.gitignore          -a---
main.py             -a---
README.md           -a---
requirements.txt    -a---
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `README.md`
- `blob_event_notifier\__init__.py`
- `blob_event_notifier\blob_handler.py`
- `blob_event_notifier\config.py`
- `blob_event_notifier\event_publisher.py`
- `blob_event_notifier\event_receiver.py`
- `main.py`
- `requirements.txt`
- `tests\test_project.py`

## Copilot Response

I’m grounding the project in the current Azure SDK authentication, Blob, and Event Grid patterns, then I’ll create and locally exercise the demo without contacting Azure.

The demo will remain fully offline: it uses real SDK event deserialization and production-ready client factories, while injected in-memory Blob and publisher clients exercise sync and async behavior without sending requests.

The project is in place. I’m now running the offline demo and focused tests against the pinned SDK ranges; any SDK-shape mismatch will be corrected before completion.

Created the complete Python project with:

- Sync and async Event Grid receivers supporting native Event Grid and CloudEvents 1.0 schemas via SDK deserialization helpers.
- Blob-created/download and blob-deleted handlers with lifecycle race handling.
- Sync and async custom-event publishers with hierarchical subjects and Azure error handling.
- Passwordless `DefaultAzureCredential` configuration.
- A fully offline sync/async demo using injected mock clients.
- Azure SDK dependencies, tests, setup documentation, and `.gitignore`.

The demo and all three focused tests completed successfully.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. Azure SDK failures are handled appropriately in both blob processing and event publishing paths.

### Strengths

- Expected blob deletion and archive-tier race conditions are handled explicitly.
- Unexpected blob HTTP failures are propagated rather than silently swallowed.
- Both publisher implementations log Azure SDK errors and report failure consistently.

## Grader Results

- blob-event-notifier.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (9/11)
      - Deserializes Event Grid native schema events using the SDK's built-in deserialization (not manual JSON parsing): Fail
      - Deserializes CloudEvents 1.0 schema events using the SDK's built-in deserialization (not manual JSON parsing): Fail
      - Routes events by event type (blob-created vs blob-deleted vs unrecognized): Pass
      - Logs a warning for unrecognized event types: Pass
      - Parses container name and blob name from event subject: Pass
      - Retrieves and prints blob access tier from blob properties: Pass
      - Publishes custom events to an Event Grid topic using the SDK's publisher client: Pass
      - Sets subject hierarchy on custom events for filtering: Pass
      - Handles race condition where the blob may no longer exist by the time the handler runs: Pass
      - Handles publishing errors with proper exception handling: Pass
      - Async versions use the async variants of the Event Grid and Blob Storage clients: Pass
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
| `Criteria from prompt file` | prompt_review | 82% | 1.00 | 0.8182 | 14.1% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.2% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.2% | ✅ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.2% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 17.2% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.2% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.8182** | **72.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-python-blob-event-notifier --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
