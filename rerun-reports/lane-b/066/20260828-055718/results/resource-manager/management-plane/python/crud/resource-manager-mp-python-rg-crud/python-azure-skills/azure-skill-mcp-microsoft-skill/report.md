# Evaluation Report: resource-manager-mp-python-rg-crud

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 610.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-python-rg-crud` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 13/14 |
| Duration | 610.4s |
| Timestamp | 2026-08-27T21:57:18Z |
| Files Generated | 3 |
| Event Count | 5966 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 506.1s |
| Review | 103.1s |
| **Total** | **610.4s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-py, azure-resource-lookup |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=539906 out=5293 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. skill — 4995ms

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

### ✅ 2. skill — 4224ms

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

### ✅ 3. glob — 191804ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 195878ms

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

### ✅ 5. view — 194446ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\non-hero-scenarios.md"}
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

### ✅ 6. view — 192979ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-resource-lookup\\references\\azure-resource-graph.md"}
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

### ✅ 7. glob — 190640ms

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-resource-lookup"],"pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 8. azure-get_azure_bestpractices (via azure) — 195879ms

**Input:**

```json
{"intent":"Generate a local-only Python CLI using azure-identity and azure-mgmt-resource to create, list, get, tag, and delete Azure Resource Groups with robust error handling. Do not execute operations against Azure.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 9. azure-get_azure_bestpractices (via azure) — 182549ms

**Input:**

```json
{"intent":"Provide Azure SDK Python code-generation best practices for a local CLI that uses azure-identity and azure-mgmt-resource to manage resource groups without executing any Azure operations.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 10. azure-documentation (via azure) — 182551ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Python azure-mgmt-resource ResourceManagementClient resource_groups create_or_update, list, get, update tags, and begin_delete methods, including error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 11. apply_patch — 974ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\u003e=1.25.0,\u003c2.0.0\n+azure-mgmt-resource\u003e=24.0.0,\u003c25.0.0\n*** Add File: resource_group_manager.py\n+\"\"\"Manage Azure Resource Groups with the Azure management plane SDK.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import json\n+import os\n+import sys\n+from collections.abc import Sequence\n+\n+from azure.core.exceptions import (\n+    AzureError,\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceNotFoundError,\n+)\n+from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n+from azure.mgmt.resource import ResourceManagementClient\n+from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable\n+\n+\n+def print_resource_group(resource_group: ResourceGroup) -\u003e None:\n+    \"\"\"Print a resource group as formatted JSON.\"\"\"\n+    print(json.dumps(resource_group.as_dict(), indent=2, sort_keys=True, default=str))\n+\n+\n+def create_resource_group(\n+    client: ResourceManagementClient,\n+    name: str,\n+    location: str,\n+) -\u003e ResourceGroup:\n+    resource_group = client.resource_groups.create_or_update(\n+        name,\n+        ResourceGroup(location=location),\n+    )\n+    print(f\"Created or updated resource group '{name}'.\")\n+    return resource_group\n+\n+\n+def list_resource_groups(client: ResourceManagementClient) -\u003e None:\n+    groups = list(client.resource_groups.list())\n+    print(json.dumps([group.as_dict() for group in groups], indent=2, default=str))\n+    print(f\"Found {len(groups)} resource group(s).\", file=sys.stderr)\n+\n+\n+def get_resource_group(\n+    client: ResourceManagementClient,\n+    name: str,\n+) -\u003e ResourceGroup:\n+    resource_group = client.resource_groups.get(name)\n+    print_resource_group(resource_group)\n+    return resource_group\n+\n+\n+def add_resource_group_tag(\n+    client: ResourceManagementClient,\n+    name: str,\n+    key: str,\n+    value: str,\n+) -\u003e ResourceGroup:\n+    current = client.resource_groups.get(name)\n+    tags = dict(current.tags or {})\n+    tags[key] = value\n+\n+    updated = client.resource_groups.update(\n+        name,\n+        ResourceGroupPatchable(tags=tags),\n+    )\n+    print(f\"Set tag '{key}={value}' on resource group '{name}'.\")\n+    print_resource_group(updated)\n+    return updated\n+\n+\n+def delete_resource_group(client: ResourceManagementClient, name: str) -\u003e None:\n+    poller = client.resource_groups.begin_delete(name)\n+    poller.result()\n+    print(f\"Deleted resource group '{name}'.\")\n+\n+\n+def run_workflow(\n+    client: ResourceManagementClient,\n+    name: str,\n+    location: str,\n+    tag_key: str,\n+    tag_value: str,\n+) -\u003e None:\n+    create_resource_group(client, name, location)\n+    list_resource_groups(client)\n+    get_resource_group(client, name)\n+    add_resource_group_tag(client, name, tag_key, tag_value)\n+    delete_resource_group(client, name)\n+\n+\n+def build_parser() -\u003e argparse.ArgumentParser:\n+    parser = argparse.ArgumentParser(\n+        description=\"Create, list, inspect, tag, and delete Azure Resource Groups.\"\n+    )\n+    parser.add_argument(\n+        \"--subscription-id\",\n+        default=os.getenv(\"AZURE_SUBSCRIPTION_ID\"),\n+        help=\"Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).\",\n+    )\n+    subparsers = parser.add_subparsers(dest=\"command\", required=True)\n+\n+    create_parser = subparsers.add_parser(\"create\", help=\"Create a resource group.\")\n+    create_parser.add_argument(\"--name\", required=True)\n+    create_parser.add_argument(\"--location\", required=True)\n+\n+    subparsers.add_parser(\"list\", help=\"List all resource groups.\")\n+\n+    get_parser = subparsers.add_parser(\"get\", help=\"Get a resource group.\")\n+    get_parser.add_argument(\"--name\", required=True)\n+\n+    tag_parser = subparsers.add_parser(\"tag\", help=\"Add or replace a tag.\")\n+    tag_parser.add_argument(\"--name\", required=True)\n+    tag_parser.add_argument(\"--key\", required=True)\n+    tag_parser.add_argument(\"--value\", required=True)\n+\n+    delete_parser = subparsers.add_parser(\"delete\", help=\"Delete a resource group.\")\n+    delete_parser.add_argument(\"--name\", required=True)\n+    delete_parser.add_argument(\n+        \"--yes\",\n+        action=\"store_true\",\n+        help=\"Confirm deletion of the resource group and all resources it contains.\",\n+    )\n+\n+    workflow_parser = subparsers.add_parser(\n+        \"workflow\",\n+        help=\"Create, list, get, tag, and delete a resource group.\",\n+    )\n+    workflow_parser.add_argument(\"--name\", required=True)\n+    workflow_parser.add_argument(\"--location\", required=True)\n+    workflow_parser.add_argument(\"--tag-key\", default=\"managed-by\")\n+    workflow_parser.add_argument(\"--tag-value\", default=\"python-sdk\")\n+    workflow_parser.add_argument(\n+        \"--yes\",\n+        action=\"store_true\",\n+        help=\"Confirm final deletion of the resource group.\",\n+    )\n+    return parser\n+\n+\n+def execute(args: argparse.Namespace, client: ResourceManagementClient) -\u003e None:\n+    if args.command == \"create\":\n+        print_resource_group(create_resource_group(client, args.name, args.location))\n+    elif args.command == \"list\":\n+        list_resource_groups(client)\n+    elif args.command == \"get\":\n+        get_resource_group(client, args.name)\n+    elif args.command == \"tag\":\n+        add_resource_group_tag(client, args.name, args.key, args.value)\n+    elif args.command == \"delete\":\n+        if not args.yes:\n+            raise ValueError(\"Deletion requires --yes.\")\n+        delete_resource_group(client, args.name)\n+    elif args.command == \"workflow\":\n+        if not args.yes:\n+            raise ValueError(\"The workflow deletes the resource group; pass --yes.\")\n+        run_workflow(\n+            client,\n+            args.name,\n+            args.location,\n+            args.tag_key,\n+            args.tag_value,\n+        )\n+\n+\n+def main(argv: Sequence[str] | None = None) -\u003e int:\n+    parser = build_parser()\n+    args = parser.parse_args(argv)\n+\n+    if not args.subscription_id:\n+        parser.error(\n+            \"Provide --subscription-id or set the AZURE_SUBSCRIPTION_ID environment variable.\"\n+        )\n+\n+    try:\n+        with DefaultAzureCredential() as credential:\n+            with ResourceManagementClient(\n+                credential,\n+                args.subscription_id,\n+            ) as client:\n+                execute(args, client)\n+        return 0\n+    except CredentialUnavailableError as error:\n+        print(\n+            \"No Azure credential is available. Configure managed identity, service \"\n+            \"principal environment variables, or a supported developer login.\",\n+            file=sys.stderr,\n+        )\n+        print(f\"Details: {error}\", file=sys.stderr)\n+    except ClientAuthenticationError as error:\n+        print(f\"Azure authentication failed: {error}\", file=sys.stderr)\n+    except ResourceNotFoundError as error:\n+        print(f\"Resource group was not found: {error}\", file=sys.stderr)\n+    except HttpResponseError as error:\n+        status = error.status_code if error.status_code is not None else \"unknown\"\n+        print(f\"Azure request failed (HTTP {status}): {error}\", file=sys.stderr)\n+    except ValueError as error:\n+        print(f\"Invalid operation: {error}\", file=sys.stderr)\n+    except AzureError as error:\n+        print(f\"Azure SDK error: {error}\", file=sys.stderr)\n+    except KeyboardInterrupt:\n+        print(\"Operation cancelled.\", file=sys.stderr)\n+\n+    return 1\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: README.md\n+# Azure Resource Group Manager\n+\n+This local Python CLI uses the Azure management plane SDK to create, list, get,\n+tag, and delete resource groups. It does not contain credentials or execute any\n+operation until you run it.\n+\n+## Install\n+\n+Python 3.9 or newer is required.\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+Required packages:\n+\n+- `azure-identity` for Microsoft Entra authentication\n+- `azure-mgmt-resource` for Resource Group management\n+\n+Set the subscription ID and configure any authentication method supported by\n+`DefaultAzureCredential`, such as service principal environment variables or a\n+managed identity:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:AZURE_TENANT_ID = \"\u003ctenant-id\u003e\"\n+$env:AZURE_CLIENT_ID = \"\u003cclient-id\u003e\"\n+$env:AZURE_CLIENT_SECRET = \"\u003cclient-secret\u003e\"\n+```\n+\n+Do not store those values in source control.\n+\n+## Usage\n+\n+Run individual operations:\n+\n+```powershell\n+python resource_group_manager.py create --name example-rg --location eastus\n+python resource_group_manager.py list\n+python resource_group_manager.py get --name example-rg\n+python resource_group_manager.py tag --name example-rg --key environment --value dev\n+python resource_group_manager.py delete --name example-rg --yes\n+```\n+\n+Run the complete create, list, get, tag, and delete lifecycle:\n+\n+```powershell\n+python resource_group_manager.py workflow --name example-rg --location eastus --yes\n+```\n+\n+Deletion removes the resource group and every resource it contains. Both\n+deletion paths require `--yes`.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt
@@ -1,0 +1,3 @@
+azure-identity>=1.25.0,<2.0.0
+azure-mgmt-resource>=24.0.0,<25.0.0
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py
@@ -1,0 +1,211 @@
+"""Manage Azure Resource Groups with the Azure management plane SDK."""
+
+from __future__ import annotations
+
+import argparse
+import json
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
+from azure.identity import CredentialUnavailableError, DefaultAzureCredential
+from azure.mgmt.resource import ResourceManagementClient
+from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable
+
+
+def print_resource_group(resource_group: ResourceGroup) -> None:
+    """Print a resource group as formatted JSON."""
+    print(json.dumps(resource_group.as_dict(), indent=2, sort_keys=True, 
... (truncated)
```

### ✅ 12. powershell — 2550ms

**Input:**

```json
{"command":"python -m py_compile .\\resource_group_manager.py \u0026\u0026 python .\\resource_group_manager.py --help","description":"Validate script syntax and CLI","initial_wait":120,"mode":"sync"}
```

**Output:**

```
Traceback (most recent call last):
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853\resource_group_manager.py", line 18, in <module>
    from azure.mgmt.resource import ResourceManagementClient
ImportError: cannot import name 'ResourceManagementClient' from 'azure.mgmt.resource' (unknown location)
<shellId: 0 completed with exit code 1>
```

### ✅ 13. powershell — 7864ms

**Input:**

```json
{"command":"python -m pip show azure-mgmt-resource azure-identity","description":"Inspect installed Azure packages","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Name: azure-mgmt-resource
Version: 26.0.0
Summary: Microsoft Azure Resource Management Client Library for Python
Home-page: 
Author: 
Author-email: Microsoft Corporation <azpysdkhelp@microsoft.com>
License-Expression: MIT
Location: C:\Users\weidxu\AppData\Roaming\Python\Python313\site-packages
Requires: azure-mgmt-core, isodate, typing-extensions
Required-by: 
---
Name: azure-identity
Version: 1.25.3
Summary: Microsoft Azure Identity Library for Python
Home-page: 
Author: 
Author-email: Microsoft Corporation <azpysdkhelp@microsoft.com>
License-Expression: MIT
Location: C:\Users\weidxu\AppData\Roaming\Python\Python313\site-packages
Requires: azure-core, cryptography, msal, msal-extensions, typing-extensions
Required-by: azure-ai-projects, azure-identity-broker, azure-managed-identity-demo, azure-managed-identity-example, managed-identity-demo
<shellId: 1 completed with exit code 0>
```

### ✅ 14. powershell — 4665ms

**Input:**

```json
{"command":"python -c \"import azure.mgmt.resource as r; print(r); print(list(getattr(r, '__path__', []))); print(sorted(x for x in dir(r) if not x.startswith('_')))\"","description":"Inspect resource SDK exports","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure\mgmt\resource\resources\__init__.py
<shellId: 4 completed with exit code 0>
```

### ✅ 15. powershell — 4307ms

**Input:**

```json
{"command":"python -c \"from azure.mgmt.resource.resources import ResourceManagementClient; from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable; print(ResourceManagementClient, ResourceGroup, ResourceGroupPatchable)\"","description":"Verify current SDK imports","initial_wait":30,"mode":"sync"}
```

**Output:**

```
<class 'azure.mgmt.resource.resources._client.ResourceManagementClient'> <class 'azure.mgmt.resource.resources.models._models.ResourceGroup'> <class 'azure.mgmt.resource.resources.models._models.ResourceGroupPatchable'>
<shellId: 3 completed with exit code 0>
```

### ✅ 16. powershell — 3958ms

**Input:**

```json
{"command":"python -m pip show -f azure-mgmt-resource | Select-String -Pattern 'ResourceManagementClient|resource_management_client|resources\\\\__init__|resources/__init__'","description":"Find management client module","initial_wait":30,"mode":"sync"}
```

**Output:**

```
<module 'azure.mgmt.resource' (namespace) from ['C:\\Users\\weidxu\\AppData\\Roaming\\Python\\Python313\\site-packages\\azure\\mgmt\\resource']>
['C:\\Users\\weidxu\\AppData\\Roaming\\Python\\Python313\\site-packages\\azure\\mgmt\\resource']
[]
<shellId: 2 completed with exit code 0>
```

### ✅ 17. apply_patch — 662ms

**Input:**

```json
"*** Begin Patch\n*** Update File: requirements.txt\n@@\n azure-identity\u003e=1.25.0,\u003c2.0.0\n-azure-mgmt-resource\u003e=24.0.0,\u003c25.0.0\n+azure-mgmt-resource\u003e=26.0.0,\u003c27.0.0\n*** Update File: resource_group_manager.py\n@@\n from azure.core.exceptions import (\n@@\n )\n from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n-from azure.mgmt.resource import ResourceManagementClient\n+from azure.mgmt.resource.resources import ResourceManagementClient\n from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/requirements.txt
@@ -1,3 +1,3 @@
 azure-identity>=1.25.0,<2.0.0
-azure-mgmt-resource>=24.0.0,<25.0.0
+azure-mgmt-resource>=26.0.0,<27.0.0
 


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-python-rg-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-693730853/resource_group_manager.py
@@ -15,7 +15,7 @@
     ResourceNotFoundError,
 )
 from azure.identity import CredentialUnavailableError, DefaultAzureCredential
-from azure.mgmt.resource import ResourceManagementClient
+from azure.mgmt.resource.resources import ResourceManagementClient
 from azure.mgmt.resource.resources.models import ResourceGroup, ResourceGroupPatchable
 
 

```

### ✅ 18. powershell — 3028ms

**Input:**

```json
{"command":"python -m py_compile .\\resource_group_manager.py \u0026\u0026 python .\\resource_group_manager.py --help \u003e $null \u0026\u0026 python -c \"from inspect import signature; from azure.mgmt.resource.resources.operations import ResourceGroupsOperations; print(signature(ResourceGroupsOperations.create_or_update)); print(signature(ResourceGroupsOperations.update)); print(signature(ResourceGroupsOperations.begin_delete))\"","description":"Validate CLI and SDK method signatures","initial_wait":120,"mode":"sync"}
```

**Output:**

```
(self, resource_group_name: str, parameters: Union[azure.mgmt.resource.resources.models._models.ResourceGroup, collections.abc.MutableMapping[str, Any], IO[bytes]], **kwargs: Any) -> azure.mgmt.resource.resources.models._models.ResourceGroup
(self, resource_group_name: str, parameters: Union[azure.mgmt.resource.resources.models._models.ResourceGroupPatchable, collections.abc.MutableMapping[str, Any], IO[bytes]], **kwargs: Any) -> azure.mgmt.resource.resources.models._models.ResourceGroup
(self, resource_group_name: str, *, force_deletion_types: Optional[str] = None, **kwargs: Any) -> azure.core.polling._poller.LROPoller[NoneType]
<shellId: 5 completed with exit code 0>
```

### ✅ 19. powershell — 1241ms

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -Recurse -Force .\\__pycache__ }","description":"Remove generated bytecode cache","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 6 completed with exit code 0>
```

## Generated Files

- `README.md`
- `requirements.txt`
- `resource_group_manager.py`

## Copilot Response

I’ll implement this as a local CLI script using `DefaultAzureCredential`, with resource creation included because the later operations refer to the created group.

The SDK references don’t include the requested acceptance/example files, and the official guidance endpoint is unavailable, so I’m proceeding with the documented synchronous management-client patterns and explicit lifecycle cleanup.

The installed current SDK is version 26 and no longer exposes the client at the legacy package root, so I’m adjusting the imports and package constraint to the current module layout.

Implemented the Azure Resource Group management CLI:

- `resource_group_manager.py` — create, list, get, tag, delete, and complete lifecycle commands with authentication and Azure SDK error handling.
- `requirements.txt` — `azure-identity` and `azure-mgmt-resource`.
- `README.md` — installation, environment configuration, and usage examples.

Deletion requires explicit `--yes`; credentials and subscription IDs are supplied through arguments or environment variables. No Azure resources were created or modified.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed. The script provides appropriate, ordered exception handling for Azure authentication, missing resources, HTTP failures, and other SDK errors.

### Strengths

- Specific Azure SDK exceptions are handled before the broader AzureError fallback.
- HttpResponseError handling includes the HTTP status code when available.
- Failures return exit code 1 and emit actionable messages to stderr.

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
hyoka run --prompt-id resource-manager-mp-python-rg-crud --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
