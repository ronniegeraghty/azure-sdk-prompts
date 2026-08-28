# Evaluation Report: resource-manager-mp-python-rg-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 418.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-python-rg-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 13/14 |
| Duration | 418.2s |
| Timestamp | 2026-08-28T21:01:53Z |
| Files Generated | 2 |
| Event Count | 6582 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 321.4s |
| Review | 95.6s |
| **Total** | **418.2s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-resource-lookup |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=520005 out=6524 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. skill — 3068ms

**Input:**

```json
{"skill":"azure-resource-lookup"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Resource Lookup

List, find, and discover Azure resources of any type across subscriptions and resource groups. Use Azure Resource Graph (ARG) for fast, cross-cutting queries when dedicated MCP tools don't cover the resource type.

## When to Use This Skill

Use this skill when the user wants to:
- **List resources** of any type (VMs, web apps, storage accounts, container apps, databases, etc.)
- **Show resources** in a specific subscription or resource group
- Query resources **across multiple subscriptions** or resource types
- Find **orphaned resources** (unattached disks, unused NICs, idle IPs)
- Discover resources **missing required tags** or configurations
- Get a **resource inventory** spanning multiple types
- Find resources in a **specific state** (unhealthy, failed provisioning, stopped)
- Answer "**what resources do I have?**" or "**show me my Azure resources**"
- **List web apps, websites, or App Services**

> ⚠️ **Warning:** App Service / Web Apps have no dedicated MCP `list` command. Prompts like "list websites", "list web apps", or "list app services" **must** route through this skill to use Azure Resource Graph.

> 💡 **Tip:** For single-resource-type queries, first check if a dedicated MCP tool can handle it (see routing table below). If none exists, use Azure Resource Graph.

## Quick Reference

| Property | Value |
|----------|-------|
| **Query Language** | KQL (Kusto Query Language subset) |
| **CLI Command** | `az graph query -q "<KQL>" -o table` |
| **Extension** | `az extension add --name resource-graph` |
| **MCP Tool** | `extension_cli_generate` with intent for `az graph query` |
| **Best For** | Cross-subscription queries, orphaned resources, tag audits |

## MCP Tools

| Tool | Purpose | When to Use |
|------|---------|-------------|
| `extension_cli_generate` | Generate `az graph query` commands | Primary tool — generate ARG queries from user intent |
| `mcp_azure_mcp_subscription_list` | Li
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 8505ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate a local Python script using Azure Identity and the Azure Resource Manager management-plane SDK to create, list, get, tag, and delete resource groups with robust Azure SDK error handling. Do not deploy or execute against Azure."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 9846ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-resource-lookup\\references\\azure-resource-graph.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md
@@ -1,185 +1,185 @@
 # Azure Resource Graph Query Patterns
 
 Azure Resource Graph (ARG) queries use a KQL subset against indexed Azure resource metadata. Results are near real-time across all subscriptions.
 
 ## Command Format
 
 ```bash
 az graph query -q "<KQL>" --query "data[].{col1:field1, col2:field2}" -o table
 ```
 
 | Flag | Purpose |
 |------|---------|
 | `-q` | KQL query string |
 | `--query` | JMESPath to shape output columns |
 | `--first N` | Limit to N results |
 | `--subscriptions` | Scope to specific subscription IDs |
 | `-o table` | Table output (also: json, tsv) |
 
 ## Key Tables
 
 | Table | Contents |
 |-------|----------|
 | `Resources` | All ARM resources — name, type, location, properties, tags, sku |
 | `ResourceContainers` | Subscriptions, resource groups, management groups |
 | `HealthResources` | Resource health availability status |
 | `ServiceHealthResources` | Azure service health events/incidents |
 | `AuthorizationResources` | Role assignments and definitions |
 | `AdvisorResources` | Azure Advisor recommendations |
 
 ## KQL Essentials
 
 - `=~` case-insensitive equals (use for `type` field — types are lowercase)
 - `properties.fieldName` navigates the properties JSON bag
 - `mv-expand` flattens arrays (subnets, IP configs)
 - `isempty()` / `is
... (truncated)
```

### ✅ 4. glob — 7247ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 3360ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Python management-plane SDK code-generation best practices for a script that authenticates with DefaultAzureCredential and creates, lists, gets, tags, and deletes resource groups without executing it against Azure.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. web_fetch — 9159ms

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-identity/json"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme?view=azure-python:
---
layout: Conceptual
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: Azure Identity client library for Python | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme?view=azure-python
config_moniker_range: azure-python
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: azure-sdk
ms.author: azsdkdocs
ms.custom: tracking-python
keywords: Azure, python, SDK, API, azure-identity, entra-id
ms.date: 2026-03-13T00:00:00.0000000Z
ms.topic: reference
ms.devlang: python
ms.service: entra-id
locale: en-us
document_id: 13bfe6a5-7920-bc4f-1604-403a406bc4dc
document_version_independent_id: 4ac2e58d-78a8-0675-284c-3df6c9cab28e
updated_at: 2026-03-13T11:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-services/latest/identity-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/7ef6598481fff5f5c02a183108a324df6f9dbac1/docs-ref-services/latest/identity-readme.md
git_commit_id: 7ef6598481fff5f5c02a183108a324df6f9dbac1
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
page_type: conceptual
toc_rel: ../../toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 1665
asset_id: api/overview/azure/identity-readme
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-services/latest/identity-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
- https://authoring-docs-microsoft.poolparty.biz/devrel/1ae5c491-970a-4062-8301-6336e69f9026
- https://microsoft-devrel.poolparty.biz/DevRelOfferingOntolo
... (truncated)
```

### ✅ 7. web_fetch — 8828ms

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-mgmt-resource/json"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/overview/azure/resources?view=azure-python:
---
layout: Conceptual
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: Azure Resources SDK for Python | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/overview/azure/resources?view=azure-python
config_moniker_range: azure-python
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: azure-sdk
ms.author: azsdkdocs
ms.custom: tracking-python
description: Reference for Azure Resources SDK for Python
ms.date: 2026-08-27T00:00:00.0000000Z
ms.topic: reference
ms.devlang: python
ms.service: resources
locale: en-us
document_id: 235574af-54a7-fffb-9a63-f65814aa9f9d
document_version_independent_id: ff8553a6-c819-6cbd-4201-f1ce7c382fc0
updated_at: 2026-08-28T11:12:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-services/latest/resources.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/40336509281dac7bd1be84dff1443c9d4a01ffdd/docs-ref-services/latest/resources.md
git_commit_id: 40336509281dac7bd1be84dff1443c9d4a01ffdd
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
page_type: conceptual
toc_rel: ../../toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 90
asset_id: api/overview/azure/resources
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-services/latest/resources.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
spProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/90370425-aca4-4a39-9533-d52e5e002a5d
platformId: f1db1493-011a-6d22-b5df-977b53177632
---

# Azure Resources SDK for Python | Microsoft
... (truncated)
```

### ✅ 8. web_fetch — 8501ms

**Input:**

```json
{"max_length":2000,"url":"https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme?view=azure-python"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-identity/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com> License-Expression: MIT","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.14","Programming Language :: Python :: 3.9"],"description":"# Azure Identity client library for Python\n\nThe Azure Identity library provides [Microsoft Entra ID](https://learn.microsoft.com/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [`TokenCredential`][token_cred_ref]/[`SupportsTokenInfo`][supports_token_info_ref] implementations, which can be used to construct Azure SDK clients that support Microsoft Entra token authent

<note>Content truncated. Call the fetch tool with a start_index of 1000 to get more content.</note>
```

### ✅ 9. web_fetch — 7214ms

**Input:**

```json
{"max_length":2000,"url":"https://learn.microsoft.com/en-us/python/api/overview/azure/resources?view=azure-python"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-mgmt-resource/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com>","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.14"],"description":"# Microsoft Azure SDK for Python\n\nThis is the Microsoft Azure Resource Management Client Library.\nThis package has been tested with Python 3.10+.\nFor a more complete view of Azure libraries, see the [azure sdk python release](https://aka.ms/azsdk/python/all).\n\n## _Disclaimer_\n\n_Azure SDK Python packages support for Python 2.7 has ended 01 January 2022. For more information and questions, please refer to https://github.com/Azure/azure-sdk-for-python/issues/20691_\n\n## Getting st

<note>Content truncated. Call the fetch tool with a start_index of 1000 to get more content.</note>
```

### ✅ 10. powershell — 1835ms

**Input:**

```json
{"command":"$identity = (Invoke-RestMethod -Uri 'https://pypi.org/pypi/azure-identity/json').info.version; $resource = (Invoke-RestMethod -Uri 'https://pypi.org/pypi/azure-mgmt-resource/json').info.version; Write-Output \"azure-identity==$identity\"; Write-Output \"azure-mgmt-resource==$resource\"","description":"Get current Azure SDK package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-identity==1.25.3
azure-mgmt-resource==26.0.0
<shellId: 0 completed with exit code 0>
```

### ✅ 11. apply_patch — 625ms

**Input:**

```json
"*** Begin Patch\n*** Add File: manage_resource_groups.py\n+\"\"\"Create, inspect, tag, and delete an Azure resource group.\n+\n+Authentication uses DefaultAzureCredential. Set AZURE_SUBSCRIPTION_ID,\n+AZURE_RESOURCE_GROUP, and AZURE_LOCATION, or pass their command-line options.\n+\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import logging\n+import os\n+import sys\n+from collections.abc import Sequence\n+\n+from azure.core.exceptions import (\n+    AzureError,\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceNotFoundError,\n+)\n+from azure.identity import DefaultAzureCredential\n+from azure.mgmt.resource import ResourceManagementClient\n+from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable\n+\n+LOGGER = logging.getLogger(\"resource-group-manager\")\n+\n+\n+def create_resource_group(\n+    client: ResourceManagementClient, name: str, location: str\n+) -\u003e ResourceGroup:\n+    \"\"\"Create a resource group, refusing to overwrite an existing group.\"\"\"\n+    try:\n+        existing = client.resource_groups.get(name)\n+    except ResourceNotFoundError:\n+        existing = None\n+\n+    if existing is not None:\n+        raise ValueError(\n+            f\"Resource group '{name}' already exists; refusing to modify or delete it.\"\n+        )\n+\n+    resource_group = client.resource_groups.create_or_update(\n+        name, ResourceGroup(location=location)\n+    )\n+    LOGGER.info(\"Created resource group '%s' in '%s'.\", name, location)\n+    return resource_group\n+\n+\n+def list_resource_groups(client: ResourceManagementClient) -\u003e None:\n+    \"\"\"List every resource group visible in the subscription.\"\"\"\n+    print(\"\\nResource groups in the subscription:\")\n+    found = False\n+    for resource_group in client.resource_groups.list():\n+        found = True\n+        print(\n+            f\"- {resource_group.name} \"\n+            f\"(location={resource_group.location}, tags={resource_group.tags or {}})\"\n+        )\n+    if not found:\n+        print(\"- None\")\n+\n+\n+def get_resource_group(\n+    client: ResourceManagementClient, name: str\n+) -\u003e ResourceGroup:\n+    \"\"\"Get and display the requested resource group.\"\"\"\n+    resource_group = client.resource_groups.get(name)\n+    print(\"\\nCreated resource group details:\")\n+    print(f\"  id: {resource_group.id}\")\n+    print(f\"  name: {resource_group.name}\")\n+    print(f\"  location: {resource_group.location}\")\n+    print(f\"  tags: {resource_group.tags or {}}\")\n+    return resource_group\n+\n+\n+def add_tag(\n+    client: ResourceManagementClient, name: str, tag_key: str, tag_value: str\n+) -\u003e ResourceGroup:\n+    \"\"\"Add or replace one tag while preserving all other tags.\"\"\"\n+    resource_group = client.resource_groups.get(name)\n+    tags = dict(resource_group.tags or {})\n+    tags[tag_key] = tag_value\n+\n+    updated = client.resource_groups.update(\n+        name, ResourceGroupPatchable(tags=tags)\n+    )\n+    LOGGER.info(\n+        \"Set tag '%s=%s' on resource group '%s'.\", tag_key, tag_value, name\n+    )\n+    print(f\"\\nUpdated tags: {updated.tags or {}}\")\n+    return updated\n+\n+\n+def delete_resource_group(client: ResourceManagementClient, name: str) -\u003e None:\n+    \"\"\"Delete a resource group and wait for the operation to finish.\"\"\"\n+    LOGGER.info(\"Deleting resource group '%s'...\", name)\n+    client.resource_groups.begin_delete(name).result()\n+    LOGGER.info(\"Deleted resource group '%s'.\", name)\n+\n+\n+def parse_args(argv: Sequence[str] | None = None) -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(description=__doc__)\n+    parser.add_argument(\n+        \"--subscription-id\",\n+        default=os.getenv(\"AZURE_SUBSCRIPTION_ID\"),\n+        help=\"Azure subscription ID (or set AZURE_SUBSCRIPTION_ID).\",\n+    )\n+    parser.add_argument(\n+        \"--resource-group\",\n+        default=os.getenv(\"AZURE_RESOURCE_GROUP\"),\n+        help=\"New resource group name (or set AZURE_RESOURCE_GROUP).\",\n+    )\n+    parser.add_argument(\n+        \"--location\",\n+        default=os.getenv(\"AZURE_LOCATION\"),\n+        help=\"Azure region, such as eastus (or set AZURE_LOCATION).\",\n+    )\n+    parser.add_argument(\"--tag-key\", default=\"Environment\")\n+    parser.add_argument(\"--tag-value\", default=\"Demo\")\n+    parser.add_argument(\n+        \"--keep-on-error\",\n+        action=\"store_true\",\n+        help=\"Do not delete the newly created group if a later operation fails.\",\n+    )\n+    args = parser.parse_args(argv)\n+\n+    missing = [\n+        option\n+        for option, value in (\n+            (\"--subscription-id\", args.subscription_id),\n+            (\"--resource-group\", args.resource_group),\n+            (\"--location\", args.location),\n+        )\n+        if not value\n+    ]\n+    if missing:\n+        parser.error(f\"missing required configuration: {', '.join(missing)}\")\n+    return args\n+\n+\n+def main(argv: Sequence[str] | None = None) -\u003e int:\n+    args = parse_args(argv)\n+    credential = DefaultAzureCredential()\n+    client = ResourceManagementClient(credential, args.subscription_id)\n+    created = False\n+\n+    try:\n+        create_resource_group(client, args.resource_group, args.location)\n+        created = True\n+        list_resource_groups(client)\n+        get_resource_group(client, args.resource_group)\n+        add_tag(\n+            client,\n+            args.resource_group,\n+            args.tag_key,\n+            args.tag_value,\n+        )\n+        delete_resource_group(client, args.resource_group)\n+        created = False\n+        return 0\n+    except ClientAuthenticationError as error:\n+        LOGGER.error(\"Azure authentication failed: %s\", error.message)\n+    except ResourceNotFoundError as error:\n+        LOGGER.error(\"The requested Azure resource was not found: %s\", error.message)\n+    except HttpResponseError as error:\n+        status = error.status_code if error.status_code is not None else \"unknown\"\n+        LOGGER.error(\"Azure request failed (HTTP %s): %s\", status, error.message)\n+    except ValueError as error:\n+        LOGGER.error(\"%s\", error)\n+    except AzureError as error:\n+        LOGGER.error(\"Azure SDK operation failed: %s\", error)\n+    finally:\n+        if created and not args.keep_on_error:\n+            try:\n+                LOGGER.warning(\"Cleaning up the resource group after an error.\")\n+                delete_resource_group(client, args.resource_group)\n+            except AzureError as cleanup_error:\n+                LOGGER.error(\"Cleanup failed: %s\", cleanup_error)\n+        client.close()\n+        credential.close()\n+\n+    return 1\n+\n+\n+if __name__ == \"__main__\":\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(message)s\",\n+    )\n+    sys.exit(main())\n*** Add File: requirements.txt\n+azure-identity==1.25.3\n+azure-mgmt-resource==26.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py
@@ -1,0 +1,191 @@
+"""Create, inspect, tag, and delete an Azure resource group.
+
+Authentication uses DefaultAzureCredential. Set AZURE_SUBSCRIPTION_ID,
+AZURE_RESOURCE_GROUP, and AZURE_LOCATION, or pass their command-line options.
+"""
+
+from __future__ import annotations
+
+import argparse
+import logging
+import os
+import sys
+from collections.abc import Sequence
+
+from azure.core.exceptions import (
+    AzureError,
+    ClientAuthenticationError,
+    HttpResponseError,
+    ResourceNotFoundError,
+)
+from azure.identity import DefaultAzureCredential
+from azure.mgmt.resource import ResourceManagementClient
+from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable
+
+LOGGER = logging.getLogger("resource-group-manager")
+
+
+def create_resource_group(
+    client: ResourceManagementClient, name: str, location: str
+) -> ResourceGroup:
+    """Create a resource group, refusing to overwrite an existing group."""
+    try:
+        existing = client.resource_groups.get(name)
+    except ResourceNotFoundError:
+        existing = None
+
+    if existing is not None:
+        raise ValueError(
+            f"Resource group '{name}' already exists; refusing to modify or delete it."
+        )
+
+    resource_group = client.resource_groups.create_or_update(
+        name, ResourceGroup(location=location)
+    )
+    LOGGER.info("Created resource group '%s' in '%s'.", name, location)
+    return resource_group
+
+
+
... (truncated)
```

### ✅ 12. powershell — 54666ms

**Input:**

```json
{"command":"$deps = Join-Path (Get-Location) '.validation-deps'; $cache = Join-Path (Get-Location) '__pycache__'; try { python -m pip install --quiet --disable-pip-version-check --target $deps -r requirements.txt; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $env:PYTHONPATH = $deps; python -m py_compile manage_resource_groups.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python manage_resource_groups.py --help | Select-Object -First 20; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -c \"import inspect; from azure.mgmt.resource import ResourceManagementClient; ops=ResourceManagementClient.resource_groups; print('SDK import succeeded')\" } finally { if (Test-Path -LiteralPath $deps) { Remove-Item -LiteralPath $deps -Recurse -Force }; if (Test-Path -LiteralPath $cache) { Remove-Item -LiteralPath $cache -Recurse -Force } }","description":"Validate script with pinned Azure SDKs","initial_wait":180,"mode":"sync"}
```

**Output:**

```
ERROR: pip's dependency resolver does not currently take into account all the packages that are installed. This behaviour is the source of the following dependency conflicts.
onnx 1.21.0 requires ml_dtypes>=0.5.0; platform_machine != "s390x", which is not installed.
pyopenssl 25.3.0 requires cryptography<47,>=45.0.7, but you have cryptography 50.0.0 which is incompatible.
streamlit 1.51.0 requires protobuf<7,>=3.20, but you have protobuf 7.34.1 which is incompatible.
Traceback (most recent call last):
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716\manage_resource_groups.py", line 22, in <module>
    from azure.mgmt.resource import ResourceManagementClient
ImportError: cannot import name 'ResourceManagementClient' from 'azure.mgmt.resource' (unknown location)
<shellId: 1 completed with exit code 1>
```

### ✅ 13. powershell — 54845ms

**Input:**

```json
{"command":"$venv = Join-Path (Get-Location) '.validation-venv'; $cache = Join-Path (Get-Location) '__pycache__'; try { python -m venv $venv; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" -m pip install --quiet --disable-pip-version-check -r requirements.txt; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" -c \"import azure.mgmt.resource as r; print(r.__file__); print([n for n in dir(r) if 'Client' in n])\"; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" -m py_compile manage_resource_groups.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" manage_resource_groups.py --help | Select-Object -First 20 } finally { if (Test-Path -LiteralPath $venv) { Remove-Item -LiteralPath $venv -Recurse -Force }; if (Test-Path -LiteralPath $cache) { Remove-Item -LiteralPath $cache -Recurse -Force } }","description":"Validate Azure imports in isolated environment","initial_wait":180,"mode":"sync"}
```

**Output:**

```
None
[]
Traceback (most recent call last):
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716\manage_resource_groups.py", line 22, in <module>
    from azure.mgmt.resource import ResourceManagementClient
ImportError: cannot import name 'ResourceManagementClient' from 'azure.mgmt.resource' (unknown location)
<shellId: 2 completed with exit code 0>
```

### ✅ 14. web_search (via github-mcp-server) — 15614ms

**Input:**

```json
{"query":"azure-mgmt-resource 26.0.0 Python ResourceManagementClient import package changes"}
```

**Output:**

```
{"type":"output_text","text":{"value":"With the release of `azure-mgmt-resource` version **26.0.0**, there is a breaking change regarding the import path for `ResourceManagementClient` in Python:\n\n**Old (pre-26.0.0) import pattern:**\n```python\nfrom azure.mgmt.resource import ResourceManagementClient\n```\n\n**New (26.0.0 and later) import pattern:**\n```python\nfrom azure.mgmt.resource.resources import ResourceManagementClient\n```\n\nIf you attempt the old import in 26.0.0 or later, you will get an `ImportError`. This change is not well-documented in the release notes, but the Azure SDK Python team has confirmed that this update is intentional. You only need to update your import statement to the new path above and your code should continue to work【3:4†source】【3:5†source】.\n\n**Additional notes on 26.0.0:**\n- Some operation groups (`Deployments` and `DeploymentOperations`) have moved to a new independent package: `azure-mgmt-resource-deployments`.\n- There are also model and method-level changes that may affect your code, such as changes to the `Resource` model instance variables; refer to [the migration guide](https://aka.ms/azsdk/python/migrate/hybrid-models) for more details.\n\n**Summary:**  \nReplace imports like  \n```python\nfrom azure.mgmt.resource import ResourceManagementClient\n```\nwith  \n```python\nfrom azure.mgmt.resource.resources import ResourceManagementClient\n```\nfor `azure-mgmt-resource>=26.0.0`【3:4†source】【3:5†source】.","annotations":[{"text":"","start_index":712,"end_index":724,"url_citation":{"title":"`azure-mgmt-resource==26.0.0` backwards-incompatible import - GitHub","url":"https://github.com/Azure/azure-sdk-for-python/issues/47651"}},{"text":"","start_index":724,"end_index":736,"url_citation":{"title":"Azure/azure-sdk-for-python azure-mgmt-resource_26.0.0b1 on GitHub","url":"https://newreleases.io/project/github/Azure/azure-sdk-for-python/release/azure-mgmt-resource_26.0.0b1"}},{"text":"","start_index":1380,
... (truncated)
```

### ✅ 15. apply_patch — 657ms

**Input:**

```json
"*** Begin Patch\n*** Update File: manage_resource_groups.py\n@@\n-from azure.mgmt.resource import ResourceManagementClient\n+from azure.mgmt.resource.resources import ResourceManagementClient\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-4212945716/manage_resource_groups.py
@@ -19,7 +19,7 @@
     ResourceNotFoundError,
 )
 from azure.identity import DefaultAzureCredential
-from azure.mgmt.resource import ResourceManagementClient
+from azure.mgmt.resource.resources import ResourceManagementClient
 from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable
 
 LOGGER = logging.getLogger("resource-group-manager")

```

### ✅ 16. powershell — 48856ms

**Input:**

```json
{"command":"$venv = Join-Path (Get-Location) '.validation-venv'; $cache = Join-Path (Get-Location) '__pycache__'; try { python -m venv $venv; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" -m pip install --quiet --disable-pip-version-check -r requirements.txt; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" -m py_compile manage_resource_groups.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; \u0026 \"$venv\\Scripts\\python.exe\" -c \"import inspect; from azure.mgmt.resource.resources import ResourceManagementClient; from azure.mgmt.resource.resources.operations import ResourceGroupsOperations; from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable; methods=('create_or_update','list','get','update','begin_delete'); [print(f'{name}{inspect.signature(getattr(ResourceGroupsOperations, name))}') for name in methods]\"; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $help = \u0026 \"$venv\\Scripts\\python.exe\" manage_resource_groups.py --help; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $help | Select-Object -First 20 } finally { if (Test-Path -LiteralPath $venv) { Remove-Item -LiteralPath $venv -Recurse -Force }; if (Test-Path -LiteralPath $cache) { Remove-Item -LiteralPath $cache -Recurse -Force } }","description":"Verify current SDK methods and CLI","initial_wait":180,"mode":"sync"}
```

**Output:**

```
create_or_update(self, resource_group_name: str, parameters: Union[azure.mgmt.resource.resources.models._models.ResourceGroup, collections.abc.MutableMapping[str, Any], IO[bytes]], **kwargs: Any) -> azure.mgmt.resource.resources.models._models.ResourceGroup
list(self, *, filter: Optional[str] = None, top: Optional[int] = None, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.ResourceGroup']
get(self, resource_group_name: str, **kwargs: Any) -> azure.mgmt.resource.resources.models._models.ResourceGroup
update(self, resource_group_name: str, parameters: Union[azure.mgmt.resource.resources.models._models.ResourceGroupPatchable, collections.abc.MutableMapping[str, Any], IO[bytes]], **kwargs: Any) -> azure.mgmt.resource.resources.models._models.ResourceGroup
begin_delete(self, resource_group_name: str, *, force_deletion_types: Optional[str] = None, **kwargs: Any) -> azure.core.polling._poller.LROPoller[NoneType]
usage: manage_resource_groups.py [-h] [--subscription-id SUBSCRIPTION_ID]
                                 [--resource-group RESOURCE_GROUP]
                                 [--location LOCATION] [--tag-key TAG_KEY]
                                 [--tag-value TAG_VALUE] [--keep-on-error]

Create, inspect, tag, and delete an Azure resource group. Authentication uses
DefaultAzureCredential. Set AZURE_SUBSCRIPTION_ID, AZURE_RESOURCE_GROUP, and
AZURE_LOCATION, or pass their command-line options.

options:
  -h, --help            show this help message and exit
  --subscription-id SUBSCRIPTION_ID
                        Azure subscription ID (or set AZURE_SUBSCRIPTION_ID).
  --resource-group RESOURCE_GROUP
                        New resource group name (or set AZURE_RESOURCE_GROUP).
  --location LOCATION   Azure region, such as eastus (or set AZURE_LOCATION).
  --tag-key TAG_KEY
  --tag-value TAG_VALUE
  --keep-on-error       Do not delete the newly created group if a later
                        operation fails.
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `manage_resource_groups.py`
- `requirements.txt`

## Copilot Response

I’m checking the workspace and grounding the implementation in the current Azure SDK guidance before writing the local-only script.

The workspace is empty, so I’ll add a self-contained CLI script plus a pinned dependency file. Configuration will come from arguments or environment variables, and no Azure operation will be executed here.

The implementation is in place. I’m validating the pinned SDK imports, operation signatures, and CLI startup locally without authenticating or sending any Azure requests.

The first import check was contaminated by globally installed namespace packages, not an Azure call. I’m switching to an isolated virtual environment to verify the current SDK’s actual import surface before adjusting anything.

Created:

- `manage_resource_groups.py` — creates, lists, retrieves, tags, and deletes a resource group using `DefaultAzureCredential`, with Azure-specific error handling and cleanup on failure.
- `requirements.txt` — pins `azure-identity==1.25.3` and `azure-mgmt-resource==26.0.0`.

Install with `pip install -r requirements.txt`, then run using `AZURE_SUBSCRIPTION_ID`, `AZURE_RESOURCE_GROUP`, and `AZURE_LOCATION`, or the equivalent command-line options. No Azure resources were accessed or modified.

References: [Azure Identity](https://learn.microsoft.com/python/api/overview/azure/identity-readme?view=azure-python) · [Azure Resources SDK](https://learn.microsoft.com/python/api/overview/azure/resources?view=azure-python)

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. Azure SDK exceptions are handled appropriately for the primary workflow and cleanup operation.

### Strengths

- Specific Azure exceptions are handled before broader exception types.
- HTTP failures include status and message details.
- Cleanup failures are caught and reported without masking the original operation failure.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `azure-mgmt-resource` and `azure-identity` pip packages: Pass
      - `ResourceManagementClient` creation with credential and subscription_id: Pass
      - `resource_groups.create_or_update()` with `ResourceGroup` model: Pass
      - `resource_groups.list()` iteration: Pass
      - `resource_groups.get()` for details: Pass
      - Tag updates via `resource_groups.update()`: Pass
      - `resource_groups.begin_delete()` with poller: Pass
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
hyoka run --prompt-id resource-manager-mp-python-rg-crud --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
