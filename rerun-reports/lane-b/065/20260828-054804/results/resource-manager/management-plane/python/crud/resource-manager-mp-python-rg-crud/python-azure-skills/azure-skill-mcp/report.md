# Evaluation Report: resource-manager-mp-python-rg-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 553.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-python-rg-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 11/14 |
| Duration | 553.3s |
| Timestamp | 2026-08-27T21:48:04Z |
| Files Generated | 2 |
| Event Count | 3453 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 445.0s |
| Review | 107.0s |
| **Total** | **553.3s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-resource-lookup |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=226384 out=3013 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3576ms

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

### ✅ 2. view — 190061ms

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

### ✅ 3. glob — 188662ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-2583837052","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 190062ms

**Input:**

```json
{"intent":"Get current best practices for writing a local Python script that uses Azure Identity and the Azure Resource Manager management-plane SDK to create, list, get, tag, and delete resource groups with robust error handling. Do not deploy or modify any live resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 190063ms

**Input:**

```json
{"intent":"Find official current Python SDK documentation for ResourceManagementClient resource group create_or_update, list, get, update/tagging, begin_delete, DefaultAzureCredential, and Azure SDK exception handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 182795ms

**Input:**

```json
{"intent":"Python Azure SDK management-plane resource group CRUD code best practices","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 7. azure-documentation (via azure) — 182792ms

**Input:**

```json
{"intent":"Official Azure Python SDK ResourceManagementClient resource_groups methods and DefaultAzureCredential documentation","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. apply_patch — 1474ms

**Input:**

```json
"*** Begin Patch\n*** Add File: azure_resource_groups.py\n+#!/usr/bin/env python3\n+\"\"\"Create, inspect, tag, and delete an Azure resource group.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import json\n+import os\n+import sys\n+from typing import Any\n+\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceNotFoundError,\n+)\n+from azure.identity import DefaultAzureCredential\n+from azure.mgmt.resource import ResourceManagementClient\n+\n+\n+def parse_args() -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(\n+        description=\"Manage an Azure resource group with the management-plane SDK.\"\n+    )\n+    parser.add_argument(\n+        \"--subscription-id\",\n+        default=os.environ.get(\"AZURE_SUBSCRIPTION_ID\"),\n+        help=\"Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).\",\n+    )\n+    parser.add_argument(\"--name\", required=True, help=\"Resource group name.\")\n+    parser.add_argument(\n+        \"--location\",\n+        default=\"eastus\",\n+        help=\"Azure region used when creating the resource group (default: eastus).\",\n+    )\n+    parser.add_argument(\n+        \"--tag-key\",\n+        default=\"managed-by\",\n+        help=\"Tag key to add (default: managed-by).\",\n+    )\n+    parser.add_argument(\n+        \"--tag-value\",\n+        default=\"python-sdk\",\n+        help=\"Tag value to add (default: python-sdk).\",\n+    )\n+    parser.add_argument(\n+        \"--confirm-delete\",\n+        action=\"store_true\",\n+        help=\"Delete the resource group after the other operations complete.\",\n+    )\n+    return parser.parse_args()\n+\n+\n+def resource_group_to_dict(resource_group: Any) -\u003e dict[str, Any]:\n+    return {\n+        \"id\": resource_group.id,\n+        \"name\": resource_group.name,\n+        \"location\": resource_group.location,\n+        \"managed_by\": resource_group.managed_by,\n+        \"properties\": (\n+            resource_group.properties.as_dict()\n+            if resource_group.properties is not None\n+            else None\n+        ),\n+        \"tags\": resource_group.tags or {},\n+    }\n+\n+\n+def print_resource_group(resource_group: Any) -\u003e None:\n+    print(json.dumps(resource_group_to_dict(resource_group), indent=2, default=str))\n+\n+\n+def run_lifecycle(\n+    client: ResourceManagementClient,\n+    resource_group_name: str,\n+    location: str,\n+    tag_key: str,\n+    tag_value: str,\n+    confirm_delete: bool,\n+) -\u003e None:\n+    if client.resource_groups.check_existence(resource_group_name):\n+        raise ValueError(\n+            f\"Resource group '{resource_group_name}' already exists. \"\n+            \"Choose a new name to avoid modifying or deleting an existing group.\"\n+        )\n+\n+    print(f\"Creating resource group '{resource_group_name}' in '{location}'...\")\n+    created = client.resource_groups.create_or_update(\n+        resource_group_name,\n+        {\"location\": location},\n+    )\n+    print_resource_group(created)\n+\n+    print(\"\\nResource groups in the subscription:\")\n+    resource_groups = list(client.resource_groups.list())\n+    if not resource_groups:\n+        print(\"(none)\")\n+    for resource_group in resource_groups:\n+        print(f\"- {resource_group.name} ({resource_group.location})\")\n+\n+    print(f\"\\nDetails for '{resource_group_name}':\")\n+    current = client.resource_groups.get(resource_group_name)\n+    print_resource_group(current)\n+\n+    updated_tags = dict(current.tags or {})\n+    updated_tags[tag_key] = tag_value\n+    print(f\"\\nAdding tag '{tag_key}={tag_value}'...\")\n+    updated = client.resource_groups.update(\n+        resource_group_name,\n+        {\"tags\": updated_tags},\n+    )\n+    print_resource_group(updated)\n+\n+    if not confirm_delete:\n+        print(\n+            \"\\nDeletion skipped. Re-run with --confirm-delete to delete \"\n+            f\"'{resource_group_name}'.\"\n+        )\n+        return\n+\n+    print(f\"\\nDeleting resource group '{resource_group_name}'...\")\n+    delete_operation = client.resource_groups.begin_delete(resource_group_name)\n+    delete_operation.result()\n+    print(f\"Resource group '{resource_group_name}' was deleted.\")\n+\n+\n+def main() -\u003e int:\n+    args = parse_args()\n+    if not args.subscription_id:\n+        print(\n+            \"Error: provide --subscription-id or set AZURE_SUBSCRIPTION_ID.\",\n+            file=sys.stderr,\n+        )\n+        return 2\n+\n+    credential = DefaultAzureCredential()\n+    client = ResourceManagementClient(credential, args.subscription_id)\n+\n+    try:\n+        run_lifecycle(\n+            client=client,\n+            resource_group_name=args.name,\n+            location=args.location,\n+            tag_key=args.tag_key,\n+            tag_value=args.tag_value,\n+            confirm_delete=args.confirm_delete,\n+        )\n+    except ValueError as error:\n+        print(f\"Error: {error}\", file=sys.stderr)\n+        return 2\n+    except ClientAuthenticationError as error:\n+        print(\n+            \"Authentication failed. Sign in with a supported credential or configure \"\n+            f\"service-principal environment variables. Details: {error}\",\n+            file=sys.stderr,\n+        )\n+        return 3\n+    except ResourceNotFoundError as error:\n+        print(f\"Resource group was not found: {error}\", file=sys.stderr)\n+        return 4\n+    except HttpResponseError as error:\n+        status_code = error.status_code if error.status_code is not None else \"unknown\"\n+        print(\n+            f\"Azure request failed (HTTP {status_code}): {error.message}\",\n+            file=sys.stderr,\n+        )\n+        return 5\n+    finally:\n+        credential.close()\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.1,\u003c2.0.0\n+azure-mgmt-resource\u003e=23.1.1,\u003c24.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-2583837052/azure_resource_groups.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-2583837052/azure_resource_groups.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-2583837052/azure_resource_groups.py
@@ -1,0 +1,176 @@
+#!/usr/bin/env python3
+"""Create, inspect, tag, and delete an Azure resource group."""
+
+from __future__ import annotations
+
+import argparse
+import json
+import os
+import sys
+from typing import Any
+
+from azure.core.exceptions import (
+    ClientAuthenticationError,
+    HttpResponseError,
+    ResourceNotFoundError,
+)
+from azure.identity import DefaultAzureCredential
+from azure.mgmt.resource import ResourceManagementClient
+
+
+def parse_args() -> argparse.Namespace:
+    parser = argparse.ArgumentParser(
+        description="Manage an Azure resource group with the management-plane SDK."
+    )
+    parser.add_argument(
+        "--subscription-id",
+        default=os.environ.get("AZURE_SUBSCRIPTION_ID"),
+        help="Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).",
+    )
+    parser.add_argument("--name", required=True, help="Resource group name.")
+    parser.add_argument(
+        "--location",
+        default="eastus",
+        help="Azure region used when creating the resource group (default: eastus).",
+    )
+    parser.add_argument(
+        "--tag-key",
+        default="managed-by",
+        help="Tag key to add (default: managed-by).",
+    )
+    parser.add_argument(
+        "--tag-value",
+        default="python-sdk",
+        help="Tag value to add (default: python-sdk).",
+    )
+    parser.add_argument(
+        "--confirm-delete",
+        action="store_true",
+        help="Delet
... (truncated)
```

### ✅ 9. powershell — 3337ms

**Input:**

```json
{"command":"python -m py_compile .\\azure_resource_groups.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -c \"import ast, pathlib; ast.parse(pathlib.Path('azure_resource_groups.py').read_text(encoding='utf-8')); print('Syntax validation passed')\"","description":"Validate Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Syntax validation passed
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `azure_resource_groups.py`
- `requirements.txt`

## Copilot Response

I’ll align the script with the current Azure Python management SDK patterns, then add a locally runnable CLI and dependency manifest.

The workspace is empty, so I’m creating a self-contained lifecycle script plus a pinned-minimum dependency manifest. Destructive deletion will require an explicit confirmation flag.

Created:

- `azure_resource_groups.py` — creates, lists, retrieves, tags, and optionally deletes a resource group using `ResourceManagementClient`, `DefaultAzureCredential`, and Azure-specific exception handling.
- `requirements.txt` — includes `azure-identity` and `azure-mgmt-resource`.

Run with `python azure_resource_groups.py --name my-resource-group --location eastus --confirm-delete`. Set `AZURE_SUBSCRIPTION_ID` or pass `--subscription-id`. The deletion step requires `--confirm-delete` to prevent accidental resource removal.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses specific Azure exception subclasses before the general HttpResponseError handler and ensures the credential is closed in a finally block.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - `azure-mgmt-resource` and `azure-identity` pip packages: Pass
      - `ResourceManagementClient` creation with credential and subscription_id: Pass
      - `resource_groups.create_or_update()` with `ResourceGroup` model: Fail
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 14.6% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.8571** | **73.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-python-rg-crud --config python-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
