# Evaluation Report: resource-manager-mp-python-rg-crud

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 222.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-python-rg-crud` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 11/14 |
| Duration | 222.5s |
| Timestamp | 2026-08-26T08:19:20Z |
| Files Generated | 3 |
| Event Count | 5881 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 144.3s |
| Review | 77.0s |
| **Total** | **222.5s** |

## Configuration

- **name:** python-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=562151 out=5634 |
| Turn Count | 11 |

## Prompt Sent

```
Write a Python script that manages Azure Resource Groups using the management plane SDK.

**Write the code to files (use file-write tools, do not reply with code blocks).**

3. List all resource groups in the subscription
4. Get details of the created resource group
5. Add a tag to the resource group
6. Delete the resource group

Show required pip packages and include proper error handling.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-py"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a Python script using Azure management-plane SDK to list, get, tag, and delete resource groups with DefaultAzureCredential and robust error handling. Do not execute Azure operations.","learn":true}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Python Azure SDK documentation for ResourceManagementClient resource_groups list/get/update/delete and DefaultAzureCredential usage.","learn":true}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-python-rg-crud-python-azure-tools-with-azure-tools-3984105282","pattern":"**/*"}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\capabilities.md"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\non-hero-scenarios.md"}
```

### 🔧 7. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate safe local Python management-plane SDK code for resource group list, get, tag update, and delete operations.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Confirm current ResourceManagementClient resource group operation signatures and LRO patterns.","parameters":{"query":"Python azure-mgmt-resource ResourceManagementClient resource_groups list get update delete begin_delete DefaultAzureCredential"}}
```

### 🔧 9. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python management SDK examples for Resource Groups.","parameters":{"language":"python","query":"ResourceManagementClient resource_groups list get update begin_delete DefaultAzureCredential"}}
```

### 🔧 10. web_fetch

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-identity/json"}
```

### 🔧 11. web_fetch

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-mgmt-resource/json"}
```

### 🔧 12. powershell

**Input:**

```json
{"command":"$paths = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787732408534-copilot-tool-output-b853017602fb4da08db1e965337e4659.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787732408913-copilot-tool-output-8faab0a8a7d74fb8993904e011e971c2.txt'); foreach ($path in $paths) { $data = Get-Content -Raw $path | ConvertFrom-Json; $data.results | Select-Object -First 5 | ForEach-Object { [pscustomobject]@{Title=$_.title; URL=$_.url; Description=$_.description} } | Format-List }","description":"Extract titles and URLs from Azure docs results"}
```

### 🔧 13. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787732408534-copilot-tool-output-b853017602fb4da08db1e965337e4659.txt","view_range":[1,8]}
```

### 🔧 14. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787732408913-copilot-tool-output-8faab0a8a7d74fb8993904e011e971c2.txt","view_range":[1,8]}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$identity = (Invoke-RestMethod 'https://pypi.org/pypi/azure-identity/json').info.version; $resource = (Invoke-RestMethod 'https://pypi.org/pypi/azure-mgmt-resource/json').info.version; \"azure-identity=$identity`nazure-mgmt-resource=$resource\"","description":"Get current Azure SDK package versions"}
```

### 🔧 16. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: resource_groups.py\n+\"\"\"Manage Azure Resource Groups with the Azure management-plane SDK.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import json\n+import logging\n+import os\n+import sys\n+from collections.abc import Sequence\n+\n+from azure.core.exceptions import (\n+    ClientAuthenticationError,\n+    HttpResponseError,\n+    ResourceNotFoundError,\n+    ServiceRequestError,\n+    ServiceResponseError,\n+)\n+from azure.identity import CredentialUnavailableError, DefaultAzureCredential\n+from azure.mgmt.resource import ResourceManagementClient\n+from azure.mgmt.resource.resources.models import ResourceGroupPatchable\n+\n+LOGGER = logging.getLogger(\"resource_groups\")\n+\n+\n+def resource_group_to_dict(resource_group: object) -\u003e dict[str, object]:\n+    \"\"\"Return the useful Resource Group fields as JSON-serializable data.\"\"\"\n+    return {\n+        \"id\": getattr(resource_group, \"id\", None),\n+        \"name\": getattr(resource_group, \"name\", None),\n+        \"location\": getattr(resource_group, \"location\", None),\n+        \"managed_by\": getattr(resource_group, \"managed_by\", None),\n+        \"provisioning_state\": getattr(\n+            getattr(resource_group, \"properties\", None),\n+            \"provisioning_state\",\n+            None,\n+        ),\n+        \"tags\": getattr(resource_group, \"tags\", None) or {},\n+    }\n+\n+\n+def list_resource_groups(client: ResourceManagementClient) -\u003e None:\n+    groups = [\n+        resource_group_to_dict(group)\n+        for group in client.resource_groups.list()\n+    ]\n+    print(json.dumps(groups, indent=2, sort_keys=True))\n+\n+\n+def get_resource_group(\n+    client: ResourceManagementClient, resource_group_name: str\n+) -\u003e None:\n+    group = client.resource_groups.get(resource_group_name)\n+    print(json.dumps(resource_group_to_dict(group), indent=2, sort_keys=True))\n+\n+\n+def add_tags(\n+    client: ResourceManagementClient,\n+    resource_group_name: str,\n+    tags_to_add: dict[str, str],\n+) -\u003e None:\n+    group = client.resource_groups.get(resource_group_name)\n+    merged_tags = dict(group.tags or {})\n+    merged_tags.update(tags_to_add)\n+\n+    updated_group = client.resource_groups.update(\n+        resource_group_name,\n+        ResourceGroupPatchable(tags=merged_tags),\n+    )\n+    print(json.dumps(resource_group_to_dict(updated_group), indent=2, sort_keys=True))\n+\n+\n+def delete_resource_group(\n+    client: ResourceManagementClient, resource_group_name: str\n+) -\u003e None:\n+    LOGGER.info(\"Deleting resource group %s\", resource_group_name)\n+    client.resource_groups.begin_delete(resource_group_name).result()\n+    print(f\"Deleted resource group '{resource_group_name}'.\")\n+\n+\n+def parse_tag(value: str) -\u003e tuple[str, str]:\n+    key, separator, tag_value = value.partition(\"=\")\n+    if not separator or not key.strip():\n+        raise argparse.ArgumentTypeError(\"tags must use the format KEY=VALUE\")\n+    return key.strip(), tag_value\n+\n+\n+def build_parser() -\u003e argparse.ArgumentParser:\n+    parser = argparse.ArgumentParser(\n+        description=\"Manage Azure Resource Groups with the Azure Python SDK.\"\n+    )\n+    parser.add_argument(\n+        \"--subscription-id\",\n+        default=os.getenv(\"AZURE_SUBSCRIPTION_ID\"),\n+        help=\"Azure subscription ID (defaults to AZURE_SUBSCRIPTION_ID).\",\n+    )\n+    parser.add_argument(\n+        \"--verbose\",\n+        action=\"store_true\",\n+        help=\"Enable informational logging.\",\n+    )\n+\n+    subparsers = parser.add_subparsers(dest=\"command\", required=True)\n+    subparsers.add_parser(\"list\", help=\"List every Resource Group.\")\n+\n+    get_parser = subparsers.add_parser(\"get\", help=\"Get one Resource Group.\")\n+    get_parser.add_argument(\"resource_group_name\")\n+\n+    tag_parser = subparsers.add_parser(\n+        \"tag\",\n+        help=\"Add or replace tags without removing other tags.\",\n+    )\n+    tag_parser.add_argument(\"resource_group_name\")\n+    tag_parser.add_argument(\n+        \"--tag\",\n+        action=\"append\",\n+        required=True,\n+        type=parse_tag,\n+        metavar=\"KEY=VALUE\",\n+        help=\"Tag to add. Repeat this option to add multiple tags.\",\n+    )\n+\n+    delete_parser = subparsers.add_parser(\n+        \"delete\",\n+        help=\"Delete a Resource Group and every resource it contains.\",\n+    )\n+    delete_parser.add_argument(\"resource_group_name\")\n+    delete_parser.add_argument(\n+        \"--yes\",\n+        action=\"store_true\",\n+        help=\"Confirm the destructive operation.\",\n+    )\n+    return parser\n+\n+\n+def execute(args: argparse.Namespace) -\u003e None:\n+    with DefaultAzureCredential() as credential:\n+        with ResourceManagementClient(\n+            credential=credential,\n+            subscription_id=args.subscription_id,\n+        ) as client:\n+            if args.command == \"list\":\n+                list_resource_groups(client)\n+            elif args.command == \"get\":\n+                get_resource_group(client, args.resource_group_name)\n+            elif args.command == \"tag\":\n+                add_tags(client, args.resource_group_name, dict(args.tag))\n+            elif args.command == \"delete\":\n+                delete_resource_group(client, args.resource_group_name)\n+\n+\n+def describe_http_error(error: HttpResponseError) -\u003e str:\n+    status = f\"HTTP {error.status_code}: \" if error.status_code else \"\"\n+    request_id = None\n+    if error.response is not None:\n+        request_id = error.response.headers.get(\"x-ms-request-id\")\n+    suffix = f\" (request ID: {request_id})\" if request_id else \"\"\n+    return f\"{status}{error.message or str(error)}{suffix}\"\n+\n+\n+def main(argv: Sequence[str] | None = None) -\u003e int:\n+    parser = build_parser()\n+    args = parser.parse_args(argv)\n+\n+    logging.basicConfig(\n+        level=logging.INFO if args.verbose else logging.WARNING,\n+        format=\"%(levelname)s: %(message)s\",\n+    )\n+\n+    if not args.subscription_id:\n+        parser.error(\n+            \"set AZURE_SUBSCRIPTION_ID or pass --subscription-id before the command\"\n+        )\n+    if args.command == \"delete\" and not args.yes:\n+        parser.error(\"delete requires --yes to confirm deletion\")\n+\n+    try:\n+        execute(args)\n+    except CredentialUnavailableError as error:\n+        LOGGER.error(\n+            \"No Azure credential is available. Configure managed identity, \"\n+            \"workload identity, or supported local developer credentials: %s\",\n+            error,\n+        )\n+        return 2\n+    except ClientAuthenticationError as error:\n+        LOGGER.error(\"Azure authentication failed: %s\", error.message or error)\n+        return 2\n+    except ResourceNotFoundError as error:\n+        LOGGER.error(\"Resource Group not found: %s\", describe_http_error(error))\n+        return 3\n+    except (ServiceRequestError, ServiceResponseError) as error:\n+        LOGGER.error(\"Could not communicate with Azure: %s\", error)\n+        return 4\n+    except HttpResponseError as error:\n+        LOGGER.error(\"Azure Resource Manager request failed: %s\", describe_http_error(error))\n+        return 5\n+    except KeyboardInterrupt:\n+        LOGGER.error(\"Operation cancelled.\")\n+        return 130\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    sys.exit(main())\n*** Add File: requirements.txt\n+azure-identity\u003e=1.25.3,\u003c2.0.0\n+azure-mgmt-resource\u003e=26.0.0,\u003c27.0.0\n*** Add File: README.md\n+# Azure Resource Group Manager\n+\n+This Python 3.10+ command-line script uses the Azure management-plane SDK and\n+`DefaultAzureCredential`. It does not contain credentials or a subscription ID.\n+\n+## Install\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+python -m pip install -r requirements.txt\n+```\n+\n+Required packages:\n+\n+- `azure-identity` for Microsoft Entra authentication\n+- `azure-mgmt-resource` for Azure Resource Manager operations\n+\n+Set the target subscription:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+```\n+\n+`DefaultAzureCredential` can use managed identity in Azure or supported local\n+developer credentials. Grant the identity only the Resource Group permissions\n+needed for the selected operations.\n+\n+## Use\n+\n+List all Resource Groups:\n+\n+```powershell\n+python .\\resource_groups.py list\n+```\n+\n+Get details of a Resource Group:\n+\n+```powershell\n+python .\\resource_groups.py get \"\u003cresource-group-name\u003e\"\n+```\n+\n+Add a tag while preserving existing tags:\n+\n+```powershell\n+python .\\resource_groups.py tag \"\u003cresource-group-name\u003e\" --tag Environment=Test\n+```\n+\n+Delete a Resource Group and all resources it contains:\n+\n+```powershell\n+python .\\resource_groups.py delete \"\u003cresource-group-name\u003e\" --yes\n+```\n+\n+The `--subscription-id` global option can be used instead of the environment\n+variable. Add `--verbose` before the command for informational logs.\n+\n+## References\n+\n+- [Manage Azure resource groups by using Python](https://learn.microsoft.com/azure/azure-resource-manager/management/manage-resource-groups-python)\n+- [Azure Identity client library for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme)\n*** End Patch\n"
```

### 🔧 17. powershell

**Input:**

```json
{"command":"python -m py_compile .\\resource_groups.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python .\\resource_groups.py --help","description":"Compile script and check CLI imports","initial_wait":120,"mode":"sync"}
```

### 🔧 18. powershell

**Input:**

```json
{"command":"@'\nimport contextlib\nimport io\nfrom types import SimpleNamespace\nfrom unittest.mock import Mock\n\nimport resource_groups\n\nrg = SimpleNamespace(\n    id=\"/subscriptions/sub/resourceGroups/demo\",\n    name=\"demo\",\n    location=\"eastus\",\n    managed_by=None,\n    properties=SimpleNamespace(provisioning_state=\"Succeeded\"),\n    tags={\"owner\": \"team\"},\n)\noperations = Mock()\noperations.list.return_value = [rg]\noperations.get.return_value = rg\noperations.update.return_value = rg\noperations.begin_delete.return_value.result.return_value = None\nclient = SimpleNamespace(resource_groups=operations)\n\nwith contextlib.redirect_stdout(io.StringIO()):\n    resource_groups.list_resource_groups(client)\n    resource_groups.get_resource_group(client, \"demo\")\n    resource_groups.add_tags(client, \"demo\", {\"environment\": \"test\"})\n    resource_groups.delete_resource_group(client, \"demo\")\n\npatch = operations.update.call_args.args[1]\nassert patch.tags == {\"owner\": \"team\", \"environment\": \"test\"}\noperations.begin_delete.return_value.result.assert_called_once_with()\nassert resource_groups.parse_tag(\"key=value=with=equals\") == (\"key\", \"value=with=equals\")\nprint(\"offline behavior checks passed\")\n'@ | python -","description":"Run offline resource group behavior checks","initial_wait":120,"mode":"sync"}
```

### 🔧 19. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-python-rg-crud-python-azure-tools-with-azure-tools-3984105282","pattern":"**/*"}
```

### 🔧 20. powershell

**Input:**

```json
{"command":"Remove-Item -LiteralPath '.\\__pycache__\\resource_groups.cpython-313.pyc'; Remove-Item -LiteralPath '.\\__pycache__'","description":"Remove generated Python cache"}
```

## Generated Files

- `README.md`
- `requirements.txt`
- `resource_groups.py`

## Copilot Response

I’ll build a local-only CLI script with safe defaults: read operations run directly, while tag/delete require explicit commands and confirmation.

The SDK contract is confirmed. I’m pinning compatible package ranges and adding a command-line interface with explicit validation and Azure-specific exception reporting.

Created:

- `resource_groups.py` — CLI for listing, retrieving, tagging, and deleting Resource Groups using `DefaultAzureCredential`
- `requirements.txt` — `azure-identity` and `azure-mgmt-resource`
- `README.md` — installation, authentication, command examples, and Microsoft documentation references

Deletion requires explicit `--yes`, tag updates preserve existing tags, and Azure authentication, networking, missing-resource, and HTTP failures return clear errors and distinct exit codes. No Azure operations were executed.

## Reviewer Notes (LLM-as-Judge)

### Summary

The generated script fully meets the exception-handling criterion.

### Strengths

- Correct exception hierarchy ordering
- Clear error messages with HTTP status and request ID when available
- Distinct exit codes for authentication, missing resources, networking, and general ARM failures

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 14.6% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 5.8571** | **73.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-python-rg-crud --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
