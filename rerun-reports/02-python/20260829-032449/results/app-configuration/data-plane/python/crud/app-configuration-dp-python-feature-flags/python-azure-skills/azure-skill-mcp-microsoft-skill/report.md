# Evaluation Report: app-configuration-dp-python-feature-flags

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 321.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-python-feature-flags` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 14/16 |
| Duration | 321.5s |
| Timestamp | 2026-08-28T20:54:03Z |
| Files Generated | 7 |
| Event Count | 15087 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 212.5s |
| Review | 107.8s |
| **Total** | **321.5s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-appconfiguration-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=710375 out=10653 |
| Turn Count | 14 |

## Tool Calls

### ✅ 1. skill — 2669ms

**Input:**

```json
{"skill":"azure-appconfiguration-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure App Configuration SDK for Python

Centralized configuration management with feature flags and dynamic settings.

## Installation

```bash
pip install azure-appconfiguration
```

## Environment Variables

```bash
AZURE_APPCONFIGURATION_ENDPOINT=https://<name>.azconfig.io  # Required for Entra ID auth
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

```python
import os
from azure.appconfiguration import AzureAppConfigurationClient
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
credential = DefaultAzureCredential(require_envvar=True)
# Or use a specific credential directly in production:
# See https://learn.microsoft.com/python/api/overview/azure/identity-readme?view=azure-python#credential-classes
# credential = ManagedIdentityCredential()

with AzureAppConfigurati
... (truncated)
```

### ✅ 2. view — 6325ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-appconfiguration-py\\references\\capabilities.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/capabilities.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/capabilities.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/capabilities.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/capabilities.md
@@ -1,35 +1,35 @@
 # azure-appconfiguration-py capability coverage
 
 **SDK/package**: `azure-appconfiguration`
 
 This index maps hero scenarios in `SKILL.md` and links non-hero scenarios documented in dedicated reference files.
 
 ## Hero scenarios covered in SKILL.md
 
 - `Configuration Settings`
 - `List Settings`
 - `Feature Flags`
 - `Read-Only Settings`
 
 ## Non-hero scenarios
 
 - `Snapshots`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#snapshots`](non-hero-scenarios.md#snapshots)
 - `Async Client`: Dedicated example and implementation notes.  
   See: [`non-hero-scenarios.md#async-client`](non-hero-scenarios.md#async-client)
 - `Client Operations`: | Operation | Description |  
   See: [`non-hero-scenarios.md#client-operations`](non-hero-scenarios.md#client-operations)
 
 ## Related deep-dive references
 
 - [`non-hero-scenarios.md`](non-hero-scenarios.md): Dedicated non-hero examples and implementation notes.
 
 ## API breadth checklist
 
 - Verify client/auth mode for the environment before coding.
 - Confirm operation-group/method names against current Microsoft Learn API reference.
 - For Python SDKs with both sync and async clients, document both forms without a blanket preference.
 - Include cleanup/delete paths for created resources in examples.
 - Prefer ide
... (truncated)
```

### ✅ 3. view — 5954ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-appconfiguration-py\\references\\non-hero-scenarios.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/non-hero-scenarios.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/non-hero-scenarios.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/non-hero-scenarios.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-appconfiguration-py/references/non-hero-scenarios.md
@@ -1,60 +1,60 @@
 # azure-appconfiguration-py non-hero scenarios
 
 These scenarios are intentionally separate from hero flows in `SKILL.md`.
 They cover secondary/advanced patterns typically used after the primary end-to-end path is working.
 
 ## Snapshots
 
 ### Create Snapshot
 
 ```python
 from azure.appconfiguration import ConfigurationSnapshot, ConfigurationSettingsFilter
 
 snapshot = ConfigurationSnapshot(
     filters=[
         ConfigurationSettingsFilter(key="app:*", label="production")
     ]
 )
 
 created = client.begin_create_snapshot(
     name="v1-snapshot",
     snapshot=snapshot
 ).result()
 ```
 
 ### List Snapshot Settings
 
 ```python
 settings = client.list_configuration_settings(
     snapshot_name="v1-snapshot"
 )
 ```
 
 ## Async Client
 
 ```python
 from azure.appconfiguration.aio import AzureAppConfigurationClient
 from azure.identity.aio import DefaultAzureCredential
 
 async def main():
     async with DefaultAzureCredential() as credential:
         async with AzureAppConfigurationClient(
             base_url=endpoint,
             credential=credential
         ) as client:
             setting = await client.get_configuration_setting(key="app:message")
             print(setting.value)
 ```
 
 ## Client Operations
 
 | Operation | Description |
 |----
... (truncated)
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 5597ms

**Input:**

```json
{"command":"get_best_practices","intent":"Create a local-only Python project using Azure App Configuration and DefaultAzureCredential, with sync and async configuration services, ETag-aware caching, feature flag evaluation, and sentinel-based polling.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 3504ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Python SDK code-generation best practices for App Configuration with sync/async clients and Entra ID authentication.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 3880ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Confirm the current Azure App Configuration Python SDK API for conditional GET using ETag and MatchConditions, including async behavior and response for not modified.","parameters":{"query":"Azure App Configuration Python get_configuration_setting etag match_condition IfNotModified ResourceNotModifiedError async"}}
```

**Output:**

```
{"results":[{"title":"AzureAppConfigurationClient Class-Methods","content":"# AzureAppConfigurationClient Class\r\n\r\nRepresents a client that calls restful API of Azure App Configuration service.\r\n\r\nThis is the async version of [AzureAppConfigurationClient](https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python)\r\n\r\n## Methods\r\n\r\n### get\\_configuration\\_setting\r\n\r\nGet the matched ConfigurationSetting from Azure App Configuration service\r\n\r\nExample\r\n\r\n\u0060\u0060\u0060python\r\n\r\n   # in async function\r\n   fetched_config_setting = await async_client.get_configuration_setting(\r\n       key=\u0022MyKey\u0022, label=\u0022MyLabel\u0022\r\n   )\r\n\u0060\u0060\u0060\r\n\r\n\u0060\u0060\u0060python\r\nasync get_configuration_setting(key: str, label: str | None = None, etag: str | None = \u0027*\u0027, match_condition: MatchConditions = MatchConditions.Unconditionally, *, accept_datetime: datetime | str | None = None, **kwargs: Any) -\u003E None | ConfigurationSetting\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r\n| key\u003Cbr\u003E\u003Cbr\u003ERequired | [str](https://docs.python.org/3/library/stdtypes.html#str)\u003Cbr\u003E\u003Cbr\u003EKey of the ConfigurationSetting |\r\n| label | [str](https://docs.python.org/3/library/stdtypes.html#str) or [None](https://docs.python.org/3/library/constants.html#None)\u003Cbr\u003E\u003Cbr\u003ELabel used to identify the ConfigurationSetting. Default is *None*.\u003Cbr\u003E\u003Cbr\u003EDefault value: None |\r\n| etag | [str](https://docs.python.org/3/library/stdtypes.html#str) or [None](https://docs.python.org/3/library/constants.html#None)\u003Cbr\u003E\u003Cbr\u003ECheck if the ConfigurationSetting is changed. Set None to skip checking etag\u003Cbr\u003E\u003Cbr\u003EDefault value: \\* |\r\n| match\\_condition | [MatchConditions](https://learn.microsoft.com/python/api/azure
... (truncated)
```

### ✅ 7. rg — 1546ms

**Input:**

```json
{"-i":true,"C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787950517923-copilot-tool-output-bca6c11220f648cb9728cc85bdf57a79.txt","pattern":"etag|match_condition|ResourceNotModifiedError|IfNotModified"}
```

**Output:**

```
Output too large to read at once (65.2 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787950523564-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787950517923-copilot-tool-output-bca6c11220f648cb9728cc85bdf57a79.txt:{"results":[{"title":"AzureAppConfigurationClient Class-Methods","content":"# AzureAppConfigurationClient Class\r\n\r\nRepresents a client that calls restful API of Azure App Configuration service.\r\n\r\nThis is the async version of [AzureAppConfigurationClient](https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python
```

### ✅ 8. glob — 1196ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. web_fetch — 1521ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.appconfiguration.AzureAppConfigurationClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python
config_moniker_range: azure-python
uid: azure.appconfiguration.AzureAppConfigurationClient
module: azure.appconfiguration
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
document_id: 8deb0ffe-7281-ca03-6b5b-8d51723d10b7
document_version_independent_id: ae594a34-2ec4-838d-1bb9-2bf403f4c7bf
updated_at: 2026-06-26T11:08:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-appconfiguration/azure.appconfiguration.AzureAppConfigurationClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/838959d3ef8b18f722b32401b06e24c4c1914c51/docs-ref-autogen/azure-appconfiguration/azure.appconfiguration.AzureAppConfigurationClient.yml
git_commit_id: 838959d3ef8b18f722b32401b06e24c4c1914c51
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Represents a client that calls restful API of Azure App Configuration service. '
toc_rel: ../_splitted/azure.appconfiguration/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-appconfiguration/azure.appconfiguration.azureappconfigurat
... (truncated)
```

### ✅ 10. web_fetch — 1225ms

**Input:**

```json
{"max_length":12000,"start_index":20000,"url":"https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python:
 App Configuration resource) used to access the Azure App Configuration. |

#### Returns

| Type | Description |
| --- | --- |
| [AzureAppConfigurationClient](azure.appconfiguration.azureappconfigurationclient) | An AzureAppConfigurationClient authenticated with the connection string |

### get\_configuration\_setting

Get the matched ConfigurationSetting from Azure App Configuration service

Example

```python

   fetched_config_setting = client.get_configuration_setting(
       key="MyKey", label="MyLabel"
   )
```

```python
get_configuration_setting(key: str, label: str | None = None, etag: str | None = '*', match_condition: MatchConditions = MatchConditions.Unconditionally, *, accept_datetime: datetime | str | None = None, **kwargs: Any) -> None | ConfigurationSetting
```

#### Parameters

| Name | Description |
| --- | --- |
| key<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Key of the ConfigurationSetting |
| label | [str](https://docs.python.org/3/library/stdtypes.html#str) or [None](https://docs.python.org/3/library/constants.html#None)<br><br>Label used to identify the ConfigurationSetting. Default is *None*.<br><br>Default value: None |
| etag | [str](https://docs.python.org/3/library/stdtypes.html#str) or [None](https://docs.python.org/3/library/constants.html#None)<br><br>Check if the ConfigurationSetting is changed. Set None to skip checking etag<br><br>Default value: \* |
| match\_condition | [MatchConditions](../azure-core/azure.core.matchconditions)<br><br>The match condition to use upon the etag<br><br>Default value: MatchConditions.Unconditionally |

#### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| accept\_datetime | [datetime](https://docs.python.org/3/library/datetime.html#datetime.datetime) or [str](ht
... (truncated)
```

### ✅ 11. powershell — 5683ms

**Input:**

```json
{"command":"python -m pip index versions azure-appconfiguration --disable-pip-version-check","description":"Check available App Configuration SDK versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-appconfiguration (1.9.0)
Available versions: 1.9.0, 1.8.1, 1.8.0, 1.7.2, 1.7.1, 1.7.0, 1.6.0, 1.5.0, 1.4.0, 1.3.0, 1.2.0, 1.1.1, 1.1.0, 1.0.1, 1.0.0
  INSTALLED: 1.9.0
  LATEST:    1.9.0
<shellId: 0 completed with exit code 0>
```

### ✅ 12. powershell — 5361ms

**Input:**

```json
{"command":"python -m pip index versions azure-identity --disable-pip-version-check","description":"Check available Azure Identity versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-identity (1.25.3)
Available versions: 1.25.3, 1.25.2, 1.25.1, 1.25.0, 1.24.0, 1.23.1, 1.23.0, 1.22.0, 1.21.0, 1.20.0, 1.19.0, 1.18.0, 1.17.1, 1.17.0, 1.16.1, 1.16.0, 1.15.0, 1.14.1, 1.14.0, 1.13.0, 1.12.0, 1.11.0, 1.10.0, 1.9.0, 1.8.0, 1.7.1, 1.7.0, 1.6.1, 1.6.0, 1.5.0, 1.4.1, 1.4.0, 1.3.1, 1.3.0, 1.2.0, 1.1.0, 1.0.1, 1.0.0
  INSTALLED: 1.25.3
  LATEST:    1.25.3
<shellId: 1 completed with exit code 0>
```

### ✅ 13. web_fetch — 4997ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.aio.azureappconfigurationclient?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.aio.azureappconfigurationclient?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.appconfiguration.aio.AzureAppConfigurationClient class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-appconfiguration/azure.appconfiguration.aio.azureappconfigurationclient?view=azure-python
config_moniker_range: azure-python
uid: azure.appconfiguration.aio.AzureAppConfigurationClient
module: azure.appconfiguration.aio
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
document_id: 046864ed-751b-a271-592e-f10b7b335fb7
document_version_independent_id: 35e920f7-2362-87a8-5d24-c6ccaaeba0ca
updated_at: 2026-06-26T11:08:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-appconfiguration/azure.appconfiguration.aio.AzureAppConfigurationClient.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/838959d3ef8b18f722b32401b06e24c4c1914c51/docs-ref-autogen/azure-appconfiguration/azure.appconfiguration.aio.AzureAppConfigurationClient.yml
git_commit_id: 838959d3ef8b18f722b32401b06e24c4c1914c51
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Represents a client that calls restful API of Azure App Configuration service. This is the async version of AzureAppConfigurationClient '
toc_rel: ../_splitted/azure.appconfiguration/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: 
... (truncated)
```

### ✅ 14. powershell — 1825ms

**Input:**

```json
{"command":"python -c \"import inspect; from azure.appconfiguration import AzureAppConfigurationClient; from azure.appconfiguration.aio import AzureAppConfigurationClient as A; print(inspect.getsource(AzureAppConfigurationClient.get_configuration_setting)); print('---ASYNC---'); print(inspect.getsource(A.get_configuration_setting))\"","description":"Inspect SDK conditional GET implementation","initial_wait":30,"mode":"sync"}
```

**Output:**

```
@distributed_trace
    def get_configuration_setting(
        self,
        key: str,
        label: Optional[str] = None,
        etag: Optional[str] = "*",
        match_condition: MatchConditions = MatchConditions.Unconditionally,
        *,
        accept_datetime: Optional[Union[datetime, str]] = None,
        **kwargs: Any,
    ) -> Union[None, ConfigurationSetting]:
        """Get the matched ConfigurationSetting from Azure App Configuration service

        :param key: Key of the ConfigurationSetting
        :type key: str
        :param label: Label used to identify the ConfigurationSetting. Default is `None`.
        :type label: str or None
        :param etag: Check if the ConfigurationSetting is changed. Set None to skip checking etag
        :type etag: str or None
        :param match_condition: The match condition to use upon the etag
        :type match_condition: ~azure.core.MatchConditions
        :keyword accept_datetime: Retrieve ConfigurationSetting that existed at this datetime
        :paramtype accept_datetime: ~datetime.datetime or str or None
        :return: The matched ConfigurationSetting object
        :rtype: ~azure.appconfiguration.ConfigurationSetting or None
        :raises: :class:`~azure.core.exceptions.HttpResponseError`, \
            :class:`~azure.core.exceptions.ClientAuthenticationError`, \
            :class:`~azure.core.exceptions.ResourceNotFoundError`, \
            :class:`~azure.core.exceptions.ResourceModifiedError`, \
            :class:`~azure.core.exceptions.ResourceExistsError`

        Example

        .. code-block:: python

            fetched_config_setting = client.get_configuration_setting(
                key="MyKey", label="MyLabel"
            )
        """
        if isinstance(accept_datetime, datetime):
            accept_datetime = str(accept_datetime)
        try:
            key_value = self._impl.get_key_value(
                key=key,
                label=label,
                accept_datetime=a
... (truncated)
```

### ✅ 15. apply_patch — 695ms

**Input:**

```json
"*** Begin Patch\n*** Add File: config_service.py\n+from __future__ import annotations\n+\n+import asyncio\n+from dataclasses import dataclass\n+from threading import RLock\n+from typing import TypeAlias\n+\n+from azure.appconfiguration import AzureAppConfigurationClient, ConfigurationSetting\n+from azure.appconfiguration.aio import (\n+    AzureAppConfigurationClient as AsyncAzureAppConfigurationClient,\n+)\n+from azure.core import MatchConditions\n+from azure.core.exceptions import ResourceNotFoundError\n+\n+\n+CacheKey: TypeAlias = tuple[str, str | None]\n+PrefixQuery: TypeAlias = tuple[str, str | None]\n+\n+\n+@dataclass(frozen=True)\n+class _CachedSetting:\n+    value: str | None\n+    etag: str\n+\n+\n+@dataclass(frozen=True)\n+class _CachedPrefix:\n+    values: dict[str, str | None]\n+    page_etags: tuple[str, ...]\n+\n+\n+def _label_filter(label: str | None) -\u003e str:\n+    return label if label is not None else \"\\0\"\n+\n+\n+def _cache_setting(setting: ConfigurationSetting) -\u003e _CachedSetting:\n+    return _CachedSetting(value=setting.value, etag=str(setting.etag))\n+\n+\n+class ConfigurationService:\n+    \"\"\"Cached, synchronous access to Azure App Configuration.\"\"\"\n+\n+    def __init__(self, client: AzureAppConfigurationClient) -\u003e None:\n+        self._client = client\n+        self._settings: dict[CacheKey, _CachedSetting] = {}\n+        self._prefixes: dict[PrefixQuery, _CachedPrefix] = {}\n+        self._lock = RLock()\n+\n+    def get_setting(self, key: str, label: str | None = None) -\u003e str | None:\n+        identity = (key, label)\n+        with self._lock:\n+            cached = self._settings.get(identity)\n+            try:\n+                if cached is None:\n+                    setting = self._client.get_configuration_setting(\n+                        key=key, label=label\n+                    )\n+                else:\n+                    setting = self._client.get_configuration_setting(\n+                        key=key,\n+                        label=label,\n+                        etag=cached.etag,\n+                        match_condition=MatchConditions.IfModified,\n+                    )\n+            except ResourceNotFoundError:\n+                self._settings.pop(identity, None)\n+                return None\n+\n+            if setting is None:\n+                return cached.value if cached is not None else None\n+\n+            self._settings[identity] = _cache_setting(setting)\n+            return setting.value\n+\n+    def list_settings(\n+        self, key_prefix: str, label: str | None = None\n+    ) -\u003e dict[str, str | None]:\n+        query = (key_prefix, label)\n+        key_filter = f\"{key_prefix}*\"\n+        label_filter = _label_filter(label)\n+\n+        with self._lock:\n+            cached = self._prefixes.get(query)\n+            if cached is not None:\n+                page_etags = tuple(\n+                    str(page.etag)\n+                    for page in self._client.check_configuration_settings(\n+                        key_filter=key_filter, label_filter=label_filter\n+                    ).by_page()\n+                )\n+                if page_etags == cached.page_etags:\n+                    return dict(cached.values)\n+\n+            values: dict[str, str | None] = {}\n+            for setting in self._client.list_configuration_settings(\n+                key_filter=key_filter, label_filter=label_filter\n+            ):\n+                values[setting.key] = setting.value\n+\n+            page_etags = tuple(\n+                str(page.etag)\n+                for page in self._client.check_configuration_settings(\n+                    key_filter=key_filter, label_filter=label_filter\n+                ).by_page()\n+            )\n+            self._prefixes[query] = _CachedPrefix(values, page_etags)\n+            return dict(values)\n+\n+    def refresh_all(self) -\u003e None:\n+        \"\"\"Invalidate and reload every key and prefix requested so far.\"\"\"\n+        with self._lock:\n+            setting_queries = tuple(self._settings)\n+            prefix_queries = tuple(self._prefixes)\n+            self._settings.clear()\n+            self._prefixes.clear()\n+\n+        for key, label in setting_queries:\n+            self.get_setting(key, label)\n+        for prefix, label in prefix_queries:\n+            self.list_settings(prefix, label)\n+\n+\n+class AsyncConfigurationService:\n+    \"\"\"Cached, asynchronous access to Azure App Configuration.\"\"\"\n+\n+    def __init__(self, client: AsyncAzureAppConfigurationClient) -\u003e None:\n+        self._client = client\n+        self._settings: dict[CacheKey, _CachedSetting] = {}\n+        self._prefixes: dict[PrefixQuery, _CachedPrefix] = {}\n+        self._lock = asyncio.Lock()\n+\n+    async def get_setting(self, key: str, label: str | None = None) -\u003e str | None:\n+        identity = (key, label)\n+        async with self._lock:\n+            cached = self._settings.get(identity)\n+            try:\n+                if cached is None:\n+                    setting = await self._client.get_configuration_setting(\n+                        key=key, label=label\n+                    )\n+                else:\n+                    setting = await self._client.get_configuration_setting(\n+                        key=key,\n+                        label=label,\n+                        etag=cached.etag,\n+                        match_condition=MatchConditions.IfModified,\n+                    )\n+            except ResourceNotFoundError:\n+                self._settings.pop(identity, None)\n+                return None\n+\n+            if setting is None:\n+                return cached.value if cached is not None else None\n+\n+            self._settings[identity] = _cache_setting(setting)\n+            return setting.value\n+\n+    async def list_settings(\n+        self, key_prefix: str, label: str | None = None\n+    ) -\u003e dict[str, str | None]:\n+        query = (key_prefix, label)\n+        key_filter = f\"{key_prefix}*\"\n+        label_filter = _label_filter(label)\n+\n+        async with self._lock:\n+            cached = self._prefixes.get(query)\n+            if cached is not None:\n+                page_etags = tuple(\n+                    [\n+                        str(page.etag)\n+                        async for page in self._client.check_configuration_settings(\n+                            key_filter=key_filter, label_filter=label_filter\n+                        ).by_page()\n+                    ]\n+                )\n+                if page_etags == cached.page_etags:\n+                    return dict(cached.values)\n+\n+            values: dict[str, str | None] = {}\n+            async for setting in self._client.list_configuration_settings(\n+                key_filter=key_filter, label_filter=label_filter\n+            ):\n+                values[setting.key] = setting.value\n+\n+            page_etags = tuple(\n+                [\n+                    str(page.etag)\n+                    async for page in self._client.check_configuration_settings(\n+                        key_filter=key_filter, label_filter=label_filter\n+                    ).by_page()\n+                ]\n+            )\n+            self._prefixes[query] = _CachedPrefix(values, page_etags)\n+            return dict(values)\n+\n+    async def refresh_all(self) -\u003e None:\n+        \"\"\"Invalidate and reload every key and prefix requested so far.\"\"\"\n+        async with self._lock:\n+            setting_queries = tuple(self._settings)\n+            prefix_queries = tuple(self._prefixes)\n+            self._settings.clear()\n+            self._prefixes.clear()\n+\n+        for key, label in setting_queries:\n+            await self.get_setting(key, label)\n+        for prefix, label in prefix_queries:\n+            await self.list_settings(prefix, label)\n*** Add File: feature_flags.py\n+from __future__ import annotations\n+\n+import hashlib\n+import json\n+from typing import Any\n+\n+from config_service import AsyncConfigurationService, ConfigurationService\n+\n+\n+FEATURE_FLAG_PREFIX = \".appconfig.featureflag/\"\n+PERCENTAGE_FILTER_NAMES = {\"Microsoft.Percentage\", \"Percentage\"}\n+\n+\n+class FeatureFlagError(ValueError):\n+    pass\n+\n+\n+def _parse_flag(payload: str | None, flag_id: str) -\u003e dict[str, Any] | None:\n+    if payload is None:\n+        return None\n+    try:\n+        flag = json.loads(payload)\n+    except json.JSONDecodeError as error:\n+        raise FeatureFlagError(f\"Feature flag {flag_id!r} contains invalid JSON\") from error\n+    if not isinstance(flag, dict):\n+        raise FeatureFlagError(f\"Feature flag {flag_id!r} must contain a JSON object\")\n+    return flag\n+\n+\n+def _rollout_percentage(flag: dict[str, Any], flag_id: str) -\u003e float | None:\n+    conditions = flag.get(\"conditions\", {})\n+    filters = conditions.get(\"client_filters\", []) if isinstance(conditions, dict) else []\n+    if not isinstance(filters, list):\n+        raise FeatureFlagError(f\"Feature flag {flag_id!r} has invalid client filters\")\n+\n+    for client_filter in filters:\n+        if not isinstance(client_filter, dict):\n+            raise FeatureFlagError(f\"Feature flag {flag_id!r} has an invalid filter\")\n+        if client_filter.get(\"name\") not in PERCENTAGE_FILTER_NAMES:\n+            continue\n+        parameters = client_filter.get(\"parameters\", {})\n+        try:\n+            percentage = float(parameters[\"Value\"])\n+        except (KeyError, TypeError, ValueError) as error:\n+            raise FeatureFlagError(\n+                f\"Feature flag {flag_id!r} has an invalid rollout percentage\"\n+            ) from error\n+        if not 0 \u003c= percentage \u003c= 100:\n+            raise FeatureFlagError(\n+                f\"Feature flag {flag_id!r} rollout percentage must be between 0 and 100\"\n+            )\n+        return percentage\n+    return None\n+\n+\n+def _is_enabled_for_user(\n+    flag: dict[str, Any] | None, flag_id: str, user_id: str | None\n+) -\u003e bool:\n+    if flag is None or flag.get(\"enabled\") is not True:\n+        return False\n+\n+    percentage = _rollout_percentage(flag, flag_id)\n+    if percentage is None:\n+        return True\n+    if user_id is None:\n+        return False\n+\n+    digest = hashlib.sha256(f\"{flag_id}:{user_id}\".encode(\"utf-8\")).digest()\n+    bucket = int.from_bytes(digest[:8], \"big\") * 100 / 2**64\n+    return bucket \u003c percentage\n+\n+\n+class FeatureFlagEvaluator:\n+    def __init__(self, configuration: ConfigurationService) -\u003e None:\n+        self._configuration = configuration\n+\n+    def is_enabled(\n+        self, flag_id: str, user_id: str | None = None, label: str | None = None\n+    ) -\u003e bool:\n+        payload = self._configuration.get_setting(\n+            f\"{FEATURE_FLAG_PREFIX}{flag_id}\", label\n+        )\n+        return _is_enabled_for_user(_parse_flag(payload, flag_id), flag_id, user_id)\n+\n+\n+class AsyncFeatureFlagEvaluator:\n+    def __init__(self, configuration: AsyncConfigurationService) -\u003e None:\n+        self._configuration = configuration\n+\n+    async def is_enabled(\n+        self, flag_id: str, user_id: str | None = None, label: str | None = None\n+    ) -\u003e bool:\n+        payload = await self._configuration.get_setting(\n+            f\"{FEATURE_FLAG_PREFIX}{flag_id}\", label\n+        )\n+        return _is_enabled_for_user(_parse_flag(payload, flag_id), flag_id, user_id)\n*** Add File: config_watcher.py\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import threading\n+from collections.abc import Sequence\n+\n+from azure.core.exceptions import AzureError\n+\n+from config_service import AsyncConfigurationService, ConfigurationService\n+\n+\n+logger = logging.getLogger(__name__)\n+Sentinel = tuple[str, str | None]\n+\n+\n+def _normalize_sentinels(\n+    sentinels: Sequence[str | Sentinel],\n+) -\u003e tuple[Sentinel, ...]:\n+    return tuple(\n+        sentinel if isinstance(sentinel, tuple) else (sentinel, None)\n+        for sentinel in sentinels\n+    )\n+\n+\n+class ConfigurationWatcher:\n+    def __init__(\n+        self,\n+        configuration: ConfigurationService,\n+        sentinels: Sequence[str | Sentinel],\n+        polling_interval: float = 30.0,\n+    ) -\u003e None:\n+        if polling_interval \u003c= 0:\n+            raise ValueError(\"polling_interval must be greater than zero\")\n+        if not sentinels:\n+            raise ValueError(\"at least one sentinel key is required\")\n+        self._configuration = configuration\n+        self._sentinels = _normalize_sentinels(sentinels)\n+        self._polling_interval = polling_interval\n+        self._values: dict[Sentinel, str | None] | None = None\n+\n+    def poll_once(self) -\u003e bool:\n+        current = {\n+            sentinel: self._configuration.get_setting(*sentinel)\n+            for sentinel in self._sentinels\n+        }\n+        changed = self._values is not None and current != self._values\n+        self._values = current\n+        if changed:\n+            logger.info(\"Sentinel changed; refreshing all cached configuration\")\n+            self._configuration.refresh_all()\n+        return changed\n+\n+    def run(self, stop_event: threading.Event) -\u003e None:\n+        while not stop_event.is_set():\n+            try:\n+                self.poll_once()\n+            except AzureError:\n+                logger.exception(\"Unable to poll Azure App Configuration\")\n+            stop_event.wait(self._polling_interval)\n+\n+\n+class AsyncConfigurationWatcher:\n+    def __init__(\n+        self,\n+        configuration: AsyncConfigurationService,\n+        sentinels: Sequence[str | Sentinel],\n+        polling_interval: float = 30.0,\n+    ) -\u003e None:\n+        if polling_interval \u003c= 0:\n+            raise ValueError(\"polling_interval must be greater than zero\")\n+        if not sentinels:\n+            raise ValueError(\"at least one sentinel key is required\")\n+        self._configuration = configuration\n+        self._sentinels = _normalize_sentinels(sentinels)\n+        self._polling_interval = polling_interval\n+        self._values: dict[Sentinel, str | None] | None = None\n+\n+    async def poll_once(self) -\u003e bool:\n+        current = {\n+            sentinel: await self._configuration.get_setting(*sentinel)\n+            for sentinel in self._sentinels\n+        }\n+        changed = self._values is not None and current != self._values\n+        self._values = current\n+        if changed:\n+            logger.info(\"Sentinel changed; refreshing all cached configuration\")\n+            await self._configuration.refresh_all()\n+        return changed\n+\n+    async def run(self, stop_event: asyncio.Event) -\u003e None:\n+        while not stop_event.is_set():\n+            try:\n+                await self.poll_once()\n+            except AzureError:\n+                logger.exception(\"Unable to poll Azure App Configuration\")\n+            try:\n+                await asyncio.wait_for(\n+                    stop_event.wait(), timeout=self._polling_interval\n+                )\n+            except TimeoutError:\n+                pass\n*** Add File: main.py\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+import threading\n+import time\n+\n+from azure.appconfiguration import AzureAppConfigurationClient\n+from azure.appconfiguration.aio import (\n+    AzureAppConfigurationClient as AsyncAzureAppConfigurationClient,\n+)\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+\n+from config_service import AsyncConfigurationService, ConfigurationService\n+from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher\n+from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator\n+\n+\n+LABEL = os.getenv(\"APPCONFIG_LABEL\", \"production\")\n+FLAG_ID = os.getenv(\"APPCONFIG_DEMO_FLAG\", \"beta-dashboard\")\n+SENTINEL_KEY = os.getenv(\"APPCONFIG_SENTINEL_KEY\", \"app:sentinel\")\n+POLL_INTERVAL = float(os.getenv(\"APPCONFIG_POLL_INTERVAL\", \"5\"))\n+WATCH_SECONDS = float(os.getenv(\"APPCONFIG_DEMO_WATCH_SECONDS\", \"6\"))\n+SAMPLE_USERS = (\"alice\", \"bob\", \"carol\", \"dave\")\n+\n+\n+def run_sync_demo(endpoint: str) -\u003e None:\n+    print(\"\\n--- Synchronous demo ---\")\n+    with DefaultAzureCredential() as credential:\n+        with AzureAppConfigurationClient(endpoint, credential) as client:\n+            configuration = ConfigurationService(client)\n+            flags = FeatureFlagEvaluator(configuration)\n+\n+            print(\"Unlabeled message:\", configuration.get_setting(\"app:message\"))\n+            print(\n+                f\"{LABEL} message:\",\n+                configuration.get_setting(\"app:message\", LABEL),\n+            )\n+            print(\n+                f\"{LABEL} settings:\",\n+                configuration.list_settings(\"app:\", LABEL),\n+            )\n+            for user_id in SAMPLE_USERS:\n+                enabled = flags.is_enabled(FLAG_ID, user_id, LABEL)\n+                print(f\"{FLAG_ID} for {user_id}: {enabled}\")\n+\n+            watcher = ConfigurationWatcher(\n+                configuration,\n+                [(SENTINEL_KEY, LABEL)],\n+                polling_interval=POLL_INTERVAL,\n+            )\n+            stop_event = threading.Event()\n+            thread = threading.Thread(\n+                target=watcher.run,\n+                args=(stop_event,),\n+                name=\"app-configuration-watcher\",\n+                daemon=True,\n+            )\n+            print(f\"Watching {SENTINEL_KEY!r} for {WATCH_SECONDS:g} seconds...\")\n+            thread.start()\n+            time.sleep(WATCH_SECONDS)\n+            stop_event.set()\n+            thread.join()\n+\n+\n+async def run_async_demo(endpoint: str) -\u003e None:\n+    print(\"\\n--- Asynchronous demo ---\")\n+    async with AsyncDefaultAzureCredential() as credential:\n+        async with AsyncAzureAppConfigurationClient(endpoint, credential) as client:\n+            configuration = AsyncConfigurationService(client)\n+            flags = AsyncFeatureFlagEvaluator(configuration)\n+\n+            print(\n+                \"Unlabeled message:\",\n+                await configuration.get_setting(\"app:message\"),\n+            )\n+            print(\n+                f\"{LABEL} message:\",\n+                await configuration.get_setting(\"app:message\", LABEL),\n+            )\n+            print(\n+                f\"{LABEL} settings:\",\n+                await configuration.list_settings(\"app:\", LABEL),\n+            )\n+            for user_id in SAMPLE_USERS:\n+                enabled = await flags.is_enabled(FLAG_ID, user_id, LABEL)\n+                print(f\"{FLAG_ID} for {user_id}: {enabled}\")\n+\n+            watcher = AsyncConfigurationWatcher(\n+                configuration,\n+                [(SENTINEL_KEY, LABEL)],\n+                polling_interval=POLL_INTERVAL,\n+            )\n+            stop_event = asyncio.Event()\n+            task = asyncio.create_task(watcher.run(stop_event))\n+            print(f\"Watching {SENTINEL_KEY!r} for {WATCH_SECONDS:g} seconds...\")\n+            await asyncio.sleep(WATCH_SECONDS)\n+            stop_event.set()\n+            await task\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    endpoint = os.environ[\"AZURE_APPCONFIGURATION_ENDPOINT\"]\n+    run_sync_demo(endpoint)\n+    asyncio.run(run_async_demo(endpoint))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-appconfiguration==1.9.0\n+azure-identity==1.25.3\n*** Add File: tests/test_feature_flags.py\n+import asyncio\n+import json\n+import unittest\n+\n+from feature_flags import (\n+    AsyncFeatureFlagEvaluator,\n+    FeatureFlagError,\n+    FeatureFlagEvaluator,\n+)\n+\n+\n+def flag_payload(enabled=True, percentage=None):\n+    filters = []\n+    if percentage is not None:\n+        filters.append(\n+            {\n+                \"name\": \"Microsoft.Percentage\",\n+                \"parameters\": {\"Value\": percentage},\n+            }\n+        )\n+    return json.dumps(\n+        {\n+            \"id\": \"test-flag\",\n+            \"enabled\": enabled,\n+            \"conditions\": {\"client_filters\": filters},\n+        }\n+    )\n+\n+\n+class FakeConfiguration:\n+    def __init__(self, payload):\n+        self.payload = payload\n+\n+    def get_setting(self, key, label=None):\n+        return self.payload\n+\n+\n+class AsyncFakeConfiguration:\n+    def __init__(self, payload):\n+        self.payload = payload\n+\n+    async def get_setting(self, key, label=None):\n+        return self.payload\n+\n+\n+class FeatureFlagEvaluatorTests(unittest.TestCase):\n+    def test_disabled_flag_is_disabled(self):\n+        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload(False)))\n+        self.assertFalse(evaluator.is_enabled(\"test-flag\", \"alice\"))\n+\n+    def test_enabled_flag_without_filters_is_enabled(self):\n+        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload()))\n+        self.assertTrue(evaluator.is_enabled(\"test-flag\", \"alice\"))\n+\n+    def test_percentage_rollout_is_deterministic(self):\n+        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=30)))\n+        first = evaluator.is_enabled(\"test-flag\", \"alice\")\n+        self.assertEqual(first, evaluator.is_enabled(\"test-flag\", \"alice\"))\n+\n+    def test_rollout_boundaries(self):\n+        zero = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=0)))\n+        full = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=100)))\n+        self.assertFalse(zero.is_enabled(\"test-flag\", \"alice\"))\n+        self.assertTrue(full.is_enabled(\"test-flag\", \"alice\"))\n+\n+    def test_invalid_percentage_raises(self):\n+        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=101)))\n+        with self.assertRaises(FeatureFlagError):\n+            evaluator.is_enabled(\"test-flag\", \"alice\")\n+\n+    def test_async_evaluator(self):\n+        evaluator = AsyncFeatureFlagEvaluator(\n+            AsyncFakeConfiguration(flag_payload(percentage=100))\n+        )\n+        self.assertTrue(asyncio.run(evaluator.is_enabled(\"test-flag\", \"alice\")))\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** Add File: tests/test_config_service.py\n+import asyncio\n+import unittest\n+from types import SimpleNamespace\n+from unittest.mock import MagicMock\n+\n+from azure.core import MatchConditions\n+\n+from config_service import AsyncConfigurationService, ConfigurationService\n+from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher\n+\n+\n+def setting(key, value, etag):\n+    return SimpleNamespace(key=key, value=value, etag=etag)\n+\n+\n+class AsyncItemPager:\n+    def __init__(self, items):\n+        self._items = items\n+\n+    def __aiter__(self):\n+        async def iterate():\n+            for item in self._items:\n+                yield item\n+\n+        return iterate()\n+\n+\n+class AsyncPagePager(AsyncItemPager):\n+    def by_page(self):\n+        return self\n+\n+\n+class AsyncClient:\n+    def __init__(self):\n+        self.responses = []\n+        self.get_calls = []\n+\n+    async def get_configuration_setting(self, **kwargs):\n+        self.get_calls.append(kwargs)\n+        return self.responses.pop(0)\n+\n+\n+class ConfigurationServiceTests(unittest.TestCase):\n+    def test_single_setting_uses_etag_after_first_read(self):\n+        client = MagicMock()\n+        client.get_configuration_setting.side_effect = [\n+            setting(\"app:key\", \"one\", \"etag-1\"),\n+            None,\n+        ]\n+        service = ConfigurationService(client)\n+\n+        self.assertEqual(\"one\", service.get_setting(\"app:key\"))\n+        self.assertEqual(\"one\", service.get_setting(\"app:key\"))\n+        second_call = client.get_configuration_setting.call_args_list[1].kwargs\n+        self.assertEqual(\"etag-1\", second_call[\"etag\"])\n+        self.assertEqual(MatchConditions.IfModified, second_call[\"match_condition\"])\n+\n+    def test_prefix_uses_head_etags_to_avoid_redownload(self):\n+        client = MagicMock()\n+        client.list_configuration_settings.return_value = [\n+            setting(\"app:a\", \"A\", \"a\")\n+        ]\n+        page = SimpleNamespace(etag=\"page-etag\")\n+        client.check_configuration_settings.return_value.by_page.return_value = [page]\n+        service = ConfigurationService(client)\n+\n+        expected = {\"app:a\": \"A\"}\n+        self.assertEqual(expected, service.list_settings(\"app:\"))\n+        self.assertEqual(expected, service.list_settings(\"app:\"))\n+        client.list_configuration_settings.assert_called_once()\n+\n+    def test_watcher_refreshes_after_sentinel_change(self):\n+        service = MagicMock()\n+        service.get_setting.side_effect = [\"v1\", \"v2\"]\n+        watcher = ConfigurationWatcher(service, [\"app:sentinel\"], 1)\n+\n+        self.assertFalse(watcher.poll_once())\n+        self.assertTrue(watcher.poll_once())\n+        service.refresh_all.assert_called_once()\n+\n+    def test_async_single_setting_uses_etag(self):\n+        async def run():\n+            client = AsyncClient()\n+            client.responses = [\n+                setting(\"app:key\", \"one\", \"etag-1\"),\n+                None,\n+            ]\n+            service = AsyncConfigurationService(client)\n+            self.assertEqual(\"one\", await service.get_setting(\"app:key\"))\n+            self.assertEqual(\"one\", await service.get_setting(\"app:key\"))\n+            self.assertEqual(\"etag-1\", client.get_calls[1][\"etag\"])\n+\n+        asyncio.run(run())\n+\n+    def test_async_watcher_refreshes_after_sentinel_change(self):\n+        class Service:\n+            def __init__(self):\n+                self.values = iter((\"v1\", \"v2\"))\n+                self.refresh_count = 0\n+\n+            async def get_setting(self, key, label=None):\n+                return next(self.values)\n+\n+            async def refresh_all(self):\n+                self.refresh_count += 1\n+\n+        async def run():\n+            service = Service()\n+            watcher = AsyncConfigurationWatcher(service, [\"app:sentinel\"], 1)\n+            self.assertFalse(await watcher.poll_once())\n+            self.assertTrue(await watcher.poll_once())\n+            self.assertEqual(1, service.refresh_count)\n+\n+        asyncio.run(run())\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** Add File: .gitignore\n+__pycache__/\n+*.py[cod]\n+.venv/\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py
@@ -1,0 +1,208 @@
+from __future__ import annotations
+
+import asyncio
+from dataclasses import dataclass
+from threading import RLock
+from typing import TypeAlias
+
+from azure.appconfiguration import AzureAppConfigurationClient, ConfigurationSetting
+from azure.appconfiguration.aio import (
+    AzureAppConfigurationClient as AsyncAzureAppConfigurationClient,
+)
+from azure.core import MatchConditions
+from azure.core.exceptions import ResourceNotFoundError
+
+
+CacheKey: TypeAlias = tuple[str, str | None]
+PrefixQuery: TypeAlias = tuple[str, str | None]
+
+
+@dataclass(frozen=True)
+class _CachedSetting:
+    value: str | None
+    etag: str
+
+
+@dataclass(frozen=True)
+class _CachedPrefix:
+    values: dict[str, str | None]
+    page_etags: tuple[str, ...]
+
+
+def _label_filter(label: str | None) -> str:
+    return label if label is not None else "\0"
+
+
+def _cache_setting(setting: ConfigurationSetting) -> _CachedSetting:
+    return _CachedSetting(value=setting.value, etag=str(setting.etag))
+
+
+class ConfigurationService:
+    """Cached, synchronous access to Azure App Configuration."""
+
+    def __init__(self, client: AzureAppConfigurationClient) -> None:
+        self._client = client
+        self._settings: dict[CacheKey, _CachedSetting] = {}
+        self._prefixes: dict[PrefixQuery, _CachedPrefix] = {}
+        self._lock = RLock()
+
+    def get_setting(self, key: str, label: str
... (truncated)
```

### ✅ 16. powershell — 1884ms

**Input:**

```json
{"command":"python -m unittest discover -s tests -v \u0026\u0026 python -m compileall -q .","description":"Run unit tests and compile Python files","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_async_single_setting_uses_etag (test_config_service.ConfigurationServiceTests.test_async_single_setting_uses_etag) ... ok
test_async_watcher_refreshes_after_sentinel_change (test_config_service.ConfigurationServiceTests.test_async_watcher_refreshes_after_sentinel_change) ... ok
test_prefix_uses_head_etags_to_avoid_redownload (test_config_service.ConfigurationServiceTests.test_prefix_uses_head_etags_to_avoid_redownload) ... ok
test_single_setting_uses_etag_after_first_read (test_config_service.ConfigurationServiceTests.test_single_setting_uses_etag_after_first_read) ... ok
test_watcher_refreshes_after_sentinel_change (test_config_service.ConfigurationServiceTests.test_watcher_refreshes_after_sentinel_change) ... ok
test_async_evaluator (test_feature_flags.FeatureFlagEvaluatorTests.test_async_evaluator) ... ok
test_disabled_flag_is_disabled (test_feature_flags.FeatureFlagEvaluatorTests.test_disabled_flag_is_disabled) ... ok
test_enabled_flag_without_filters_is_enabled (test_feature_flags.FeatureFlagEvaluatorTests.test_enabled_flag_without_filters_is_enabled) ... ok
test_invalid_percentage_raises (test_feature_flags.FeatureFlagEvaluatorTests.test_invalid_percentage_raises) ... ok
test_percentage_rollout_is_deterministic (test_feature_flags.FeatureFlagEvaluatorTests.test_percentage_rollout_is_deterministic) ... ok
test_rollout_boundaries (test_feature_flags.FeatureFlagEvaluatorTests.test_rollout_boundaries) ... ok

----------------------------------------------------------------------
Ran 11 tests in 0.004s

OK
<shellId: 3 completed with exit code 0>
```

### ✅ 17. powershell — 1617ms

**Input:**

```json
{"command":"python -c \"import inspect; from azure.appconfiguration._models import ConfigurationSettingPaged; print(inspect.getsource(ConfigurationSettingPaged)); from azure.appconfiguration.aio._models import ConfigurationSettingPaged as A; print('---ASYNC---'); print(inspect.getsource(A))\"","description":"Inspect SDK paging ETag behavior","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Traceback (most recent call last):
  File "<string>", line 1, in <module>
    import inspect; from azure.appconfiguration._models import ConfigurationSettingPaged; print(inspect.getsource(ConfigurationSettingPaged)); from azure.appconfiguration.aio._models import ConfigurationSettingPaged as A; print('---ASYNC---'); print(inspect.getsource(A))
                                                                                                                                               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
ModuleNotFoundError: No module named 'azure.appconfiguration.aio._models'
class ConfigurationSettingPaged(ItemPaged[ConfigurationSetting]):
    """
    An iterable of ConfigurationSettings that supports etag-based change detection.

    This class extends ItemPaged to provide efficient monitoring of configuration changes
    by using ETags. When used with the `match_conditions` parameter in `by_page()`,
    it only returns pages that have changed since the provided ETags were collected.

    Example:

    .. code-block:: python

        # Get initial page ETags
        items = client.list_configuration_settings(key_filter="sample_*")
        match_conditions = [page.etag for page in items.by_page()]

        # Later, check for changes - only changed pages are returned
        items = client.list_configuration_settings(key_filter="sample_*")
        for page in items.by_page(match_conditions=match_conditions):
            # Process only changed pages
            pass
    """

    def by_page(self, continuation_token: Optional[str] = None, *, match_conditions: Optional[List[str]] = None) -> Any:
        """Get an iterator of pages of objects, instead of an iterator of objects.

        :param str continuation_token:
            An opaque continuation token. This value can be retrieved from the
            continuation_token field of a previous generator object. If specified,
            this generator will begin
... (truncated)
```

### ✅ 18. apply_patch — 663ms

**Input:**

```json
"*** Begin Patch\n*** Update File: config_service.py\n@@\n             if cached is not None:\n                 page_etags = tuple(\n                     str(page.etag)\n@@\n                 if page_etags == cached.page_etags:\n                     return dict(cached.values)\n \n             values: dict[str, str | None] = {}\n-            for setting in self._client.list_configuration_settings(\n+            page_etags_list: list[str] = []\n+            pages = self._client.list_configuration_settings(\n                 key_filter=key_filter, label_filter=label_filter\n-            ):\n-                values[setting.key] = setting.value\n-\n-            page_etags = tuple(\n-                str(page.etag)\n-                for page in self._client.check_configuration_settings(\n-                    key_filter=key_filter, label_filter=label_filter\n-                ).by_page()\n-            )\n+            ).by_page()\n+            for page in pages:\n+                page_etags_list.append(str(page.etag))\n+                for setting in page:\n+                    values[setting.key] = setting.value\n+\n+            page_etags = tuple(page_etags_list)\n             self._prefixes[query] = _CachedPrefix(values, page_etags)\n             return dict(values)\n@@\n                 if page_etags == cached.page_etags:\n                     return dict(cached.values)\n \n             values: dict[str, str | None] = {}\n-            async for setting in self._client.list_configuration_settings(\n+            page_etags_list: list[str] = []\n+            pages = self._client.list_configuration_settings(\n                 key_filter=key_filter, label_filter=label_filter\n-            ):\n-                values[setting.key] = setting.value\n-\n-            page_etags = tuple(\n-                [\n-                    str(page.etag)\n-                    async for page in self._client.check_configuration_settings(\n-                        key_filter=key_filter, label_filter=label_filter\n-                    ).by_page()\n-                ]\n-            )\n+            ).by_page()\n+            async for page in pages:\n+                page_etags_list.append(str(page.etag))\n+                async for setting in page:\n+                    values[setting.key] = setting.value\n+\n+            page_etags = tuple(page_etags_list)\n             self._prefixes[query] = _CachedPrefix(values, page_etags)\n             return dict(values)\n*** Update File: tests/test_config_service.py\n@@\n     def test_prefix_uses_head_etags_to_avoid_redownload(self):\n         client = MagicMock()\n-        client.list_configuration_settings.return_value = [\n-            setting(\"app:a\", \"A\", \"a\")\n-        ]\n-        page = SimpleNamespace(etag=\"page-etag\")\n-        client.check_configuration_settings.return_value.by_page.return_value = [page]\n+        list_page = MagicMock()\n+        list_page.etag = \"page-etag\"\n+        list_page.__iter__.return_value = iter([setting(\"app:a\", \"A\", \"a\")])\n+        list_pager = MagicMock()\n+        list_pager.by_page.return_value = [list_page]\n+        client.list_configuration_settings.return_value = list_pager\n+        check_page = SimpleNamespace(etag=\"page-etag\")\n+        client.check_configuration_settings.return_value.by_page.return_value = [\n+            check_page\n+        ]\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-feature-flags-python-azure-skills-azure-skill-mcp-microsoft-skill-3486405777/config_service.py
@@ -92,17 +92,16 @@
                     return dict(cached.values)
 
             values: dict[str, str | None] = {}
-            for setting in self._client.list_configuration_settings(
+            page_etags_list: list[str] = []
+            pages = self._client.list_configuration_settings(
                 key_filter=key_filter, label_filter=label_filter
-            ):
-                values[setting.key] = setting.value
+            ).by_page()
+            for page in pages:
+                page_etags_list.append(str(page.etag))
+                for setting in page:
+                    values[setting.key] = setting.value
 
-            page_etags = tuple(
-                str(page.etag)
-                for page in self._client.check_configuration_settings(
-                    key_filter=key_filter, label_filter=label_filter
-                ).by_page()
-            )
+            page_etags = tuple(page_etags_list)
             self._prefixes[query] = _CachedPrefix(values, page_etags)
             return dict(values)
 
@@ -177,19 +176,16 @@
                     return dict(cached.values)
 
             values: dict[str, str | None] = {}
-            async for setting in self._client.list_configuration_settings(
+            page_etags_list: list[str] = [
... (truncated)
```

### ✅ 19. powershell — 2733ms

**Input:**

```json
{"command":"python -m unittest discover -s tests -v \u0026\u0026 python -m compileall -q . \u0026\u0026 python -c \"import config_service, config_watcher, feature_flags, main\"","description":"Re-run tests and import all modules","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_async_single_setting_uses_etag (test_config_service.ConfigurationServiceTests.test_async_single_setting_uses_etag) ... ok
test_async_watcher_refreshes_after_sentinel_change (test_config_service.ConfigurationServiceTests.test_async_watcher_refreshes_after_sentinel_change) ... ok
test_prefix_uses_head_etags_to_avoid_redownload (test_config_service.ConfigurationServiceTests.test_prefix_uses_head_etags_to_avoid_redownload) ... ok
test_single_setting_uses_etag_after_first_read (test_config_service.ConfigurationServiceTests.test_single_setting_uses_etag_after_first_read) ... ok
test_watcher_refreshes_after_sentinel_change (test_config_service.ConfigurationServiceTests.test_watcher_refreshes_after_sentinel_change) ... ok
test_async_evaluator (test_feature_flags.FeatureFlagEvaluatorTests.test_async_evaluator) ... ok
test_disabled_flag_is_disabled (test_feature_flags.FeatureFlagEvaluatorTests.test_disabled_flag_is_disabled) ... ok
test_enabled_flag_without_filters_is_enabled (test_feature_flags.FeatureFlagEvaluatorTests.test_enabled_flag_without_filters_is_enabled) ... ok
test_invalid_percentage_raises (test_feature_flags.FeatureFlagEvaluatorTests.test_invalid_percentage_raises) ... ok
test_percentage_rollout_is_deterministic (test_feature_flags.FeatureFlagEvaluatorTests.test_percentage_rollout_is_deterministic) ... ok
test_rollout_boundaries (test_feature_flags.FeatureFlagEvaluatorTests.test_rollout_boundaries) ... ok

----------------------------------------------------------------------
Ran 11 tests in 0.005s

OK
<shellId: 5 completed with exit code 0>
```

## Generated Files

- `config_service.py`
- `config_watcher.py`
- `feature_flags.py`
- `main.py`
- `requirements.txt`
- `tests\test_config_service.py`
- `tests\test_feature_flags.py`

## Copilot Response

I’m grounding the implementation in the Azure SDK’s client lifecycle, conditional-read, and async patterns before creating the project.

The SDK exposes ETag-aware conditional reads and HEAD-based collection checks. I’m using those to keep single-key and prefix polling efficient, with sentinel-triggered cache refresh layered above them.

I’ve finalized the project shape: shared cache semantics across sync/async services, deterministic SHA-256 rollout bucketing, and watcher loops that refresh every previously requested key/prefix when a sentinel changes.

The implementation is written. I’m exercising the cache, watcher, rollout boundaries, and async paths now; any SDK-shape or concurrency issue will be corrected before completion.

The first pass exposed a stronger SDK paging pattern: collection page ETags can be captured from the same response that supplies values. I’m switching to that atomic snapshot so a concurrent update cannot leave stale prefix data paired with a newer ETag.

Created the Python project with:

- Sync and async ETag-cached configuration services
- Labeled and prefix-based setting retrieval
- Deterministic SHA-256 percentage feature-flag rollout
- Sync and async sentinel watchers with full cache refresh
- Sequential sync/async demo using `DefaultAzureCredential`
- Pinned Azure SDK dependencies
- Unit coverage for caching, watchers, and feature evaluation

Run with `AZURE_APPCONFIGURATION_ENDPOINT` set, then execute `python main.py`.

Reference: [Azure App Configuration Python SDK](https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python)

## Reviewer Notes (LLM-as-Judge)

### Summary

The generated project does not fully satisfy the required Azure SDK exception-handling criterion.

### Strengths

- Missing individual settings are handled through ResourceNotFoundError.
- Both watcher implementations catch and log AzureError so transient polling failures do not stop their loops.

### Issues

- Conditional and listing requests do not appropriately handle HttpResponseError and its relevant subclasses.
- The main demo allows Azure SDK failures to terminate execution without a clear handled error path.

## Grader Results

- feature-flags.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Lists settings filtered by key prefix using `key_filter`: Pass
      - Implements conditional reads using `match_condition` with `MatchConditions.IfModified` and the setting's `etag`: Pass
      - Handles 304 Not Modified (setting unchanged since last read): Pass
      - Uses `.appconfig.featureflag/` prefix for feature flag keys: Pass
      - Parses the JSON payload in feature flag setting values: Pass
      - Implements deterministic percentage rollout (consistent hash via `hashlib`, not `random`): Pass
      - Implements sentinel key watching with configurable polling interval: Pass
      - Detects sentinel value change via ETag or value comparison and triggers full refresh: Pass
      - Async version uses `azure.appconfiguration.aio.AzureAppConfigurationClient`: Pass
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
hyoka run --prompt-id app-configuration-dp-python-feature-flags --config python-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
