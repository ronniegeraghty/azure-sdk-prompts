# Evaluation Report: ai-projects-dp-python-project-resource-inventory

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 222.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-projects-dp-python-project-resource-inventory` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 13/16 |
| Duration | 222.3s |
| Timestamp | 2026-08-26T12:06:16Z |
| Files Generated | 3 |
| Event Count | 3610 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 140.9s |
| Review | 80.5s |
| **Total** | **222.3s** |

## Configuration

- **name:** python-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=243968 out=2837 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob

**Input:**

```json
{"pattern":"**/.*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"azure-ai-projects Python synchronous AIProjectClient connections list get include_credentials false deployments list get model deployment current API 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's a summary of how to use the Python API for the Azure AI Projects SDK (`azure-ai-projects`) with a focus on the objects and operations you mentioned:\n\n## Main Client & Usage\n\n- **AIProjectClient**: This is the main entry point for most operations, wrapping agent creation, thread management, evaluations, connections, deployments, datasets, and indexes.\n\n**Typical Setup:**\n```python\nimport os\nfrom azure.identity import DefaultAzureCredential\nfrom azure.ai.projects import AIProjectClient\n\ncredential = DefaultAzureCredential()\nwith AIProjectClient(\n    endpoint=os.environ[\"AZURE_AI_PROJECT_ENDPOINT\"],\n    credential=credential,\n) as client:\n    deployments = list(client.deployments.list())\n    connections = list(client.connections.list())\n```\n- Use environment variable `AZURE_AI_PROJECT_ENDPOINT` for the endpoint (example: `\"https://<resource>.services.ai.azure.com/api/projects/<project>\"`).\n- Use `DefaultAzureCredential`—this is recommended for local development and production.\n\n## Key Operations\n\n- **client.deployments**: Use `.deployments.list()` to enumerate AI models deployed to your Foundry Project.\n- **client.connections**: Use `.connections.list()` to enumerate connected Azure resources.\n- Other key client properties: `.agents`, `.datasets`, `.indexes`, `.evaluations`, `.red_teams`, etc.\n\n## Models & Types\n\n- **ModelDeployment**: Represents a model deployment. Defined in `azure.ai.projects.models`.\n- **ConnectionType**: Enum describing types of connections available (also in `azure.ai.projects.models`).\n\nSee the models doc for all classes in the models package: [azure.ai.projects.models reference](https://learn.microsoft.com/python/api/azure-ai-projects/azure.ai.projects.models?view=azure-python)【3:2†source】.\n\n## Example References\n\n- [Official samples repo](https://github.com/Azure/azure-sdk-for-python/tree/main/sdk/ai/azure-ai-projects/samples) for real-world code sa
... (truncated)
```

### ✅ 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com python api azure-ai-projects azure.ai.projects.models ModelDeployment ConnectionType AIProjectClient deployments connections"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s how you can use the Azure AI Projects Python SDK (`azure-ai-projects`) as of 2026 to synchronously list connections (excluding credentials) and list deployments in your AI project.\n\n## Prerequisites\n\n- Install the library:  \n  ```bash\n  pip install azure-ai-projects azure-identity\n  ```\n\n- Set environment variables:\n  - `AZURE_AI_PROJECT_ENDPOINT`\n  - (Optionally, configure Azure credentials as per [Azure Identity guidelines](https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme?view=azure-python#credential-classes))\n\n## Example: Synchronous Usage with AIProjectClient\n\n```python\nimport os\nfrom azure.identity import DefaultAzureCredential\nfrom azure.ai.projects import AIProjectClient\n\n# Get credentials and endpoint\nendpoint = os.environ[\"AZURE_AI_PROJECT_ENDPOINT\"]\ncredential = DefaultAzureCredential()\n\n# Synchronously connect to AI Project\nwith AIProjectClient(endpoint=endpoint, credential=credential) as client:\n    # 1. List connections (excluding credentials)\n    connections = list(client.connections.list())\n    for connection in connections:\n        print(\"Connection name:\", connection.name)\n        print(\"Type:\", connection.type)\n        # Credentials will NOT be included by default\n\n    # 2. Get a specific connection, explicitly excluding credentials\n    conn = client.connections.get(\"my_connection_name\", include_credentials=False)\n    print(conn)\n\n    # 3. List all model deployments \n    deployments = list(client.deployments.list())\n    for deploy in deployments:\n        print(\"Deployment name:\", deploy.name)\n        print(\"Model:\", deploy.model)\n        print(\"Status:\", deploy.status)\n```\n\n### Notes\n\n- `include_credentials` defaults to `False` for both `get` and `get_default`.\n- You can filter or paginate connections and deployments if needed.\n- Deployment listing lets you track current model deployments in your project.\n\n### 
... (truncated)
```

### ✅ 5. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-python/main/sdk/ai/azure-ai-projects/samples/connections/sample_connections.py"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.models.modeldeployment?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.ai.projects.models.ModelDeployment class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.models.modeldeployment?view=azure-python
config_moniker_range: azure-python
uid: azure.ai.projects.models.ModelDeployment
module: azure.ai.projects.models
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
document_id: 4fc4c45d-68b1-38a9-d316-306c0df5e77d
document_version_independent_id: f5254f34-7409-c0f2-99c6-feebe416ada4
updated_at: 2026-08-25T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-ai-projects/azure.ai.projects.models.ModelDeployment.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/e74658802e80af785c26a422557b6189451805b9/docs-ref-autogen/azure-ai-projects/azure.ai.projects.models.ModelDeployment.yml
git_commit_id: e74658802e80af785c26a422557b6189451805b9
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Model Deployment Definition. '
toc_rel: ../_splitted/azure.ai.projects/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-ai-projects/azure.ai.projects.models.modeldeployment
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-autogen/azure-ai-pr
... (truncated)
```

### ✅ 6. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-python/main/sdk/ai/azure-ai-projects/samples/deployments/sample_deployments.py"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-python/main/sdk/ai/azure-ai-projects/samples/deployments/sample_deployments.py:
# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------

"""
DESCRIPTION:
    Given an AIProjectClient, this sample demonstrates how to use the synchronous
    `.deployments` methods to enumerate AI models deployed to your Microsoft Foundry Project.

USAGE:
    python sample_deployments.py

    Before running the sample:

    pip install "azure-ai-projects>=2.0.0" python-dotenv

    Set these environment variables with your own values:
    1) FOUNDRY_PROJECT_ENDPOINT - Required. The Azure AI Project endpoint, as found in the overview page of your
       Microsoft Foundry project.
    2) FOUNDRY_MODEL_NAME - Required. The name of the deployment to retrieve.
    3) MODEL_PUBLISHER - Optional. The publisher of the model to filter by.
    4) MODEL_NAME - Optional. The name of the model to filter by.
"""

import os
from dotenv import load_dotenv
from azure.identity import DefaultAzureCredential
from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import ModelDeployment

load_dotenv()

endpoint = os.environ["FOUNDRY_PROJECT_ENDPOINT"]
model_deployment_name = os.environ["FOUNDRY_MODEL_NAME"]
model_publisher = os.environ.get("MODEL_PUBLISHER", "Microsoft")
model_name = os.environ.get("MODEL_NAME", "Phi-4")

with (
    DefaultAzureCredential() as credential,
    AIProjectClient(endpoint=endpoint, credential=credential) as project_client,
):

    print("List all deployments:")
    for deployment in project_client.deployments.list():
        print(deployment)

    print(f"List all deployments by the model publisher `{model_publisher}`:")
    for deployment in project_client.deployments.list(model_publisher=model_publisher
... (truncated)
```

### ✅ 7. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.models.modeldeployment?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.models.connection?view=azure-python:
---
layout: Reference
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: azure.ai.projects.models.Connection class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.models.connection?view=azure-python
config_moniker_range: azure-python
uid: azure.ai.projects.models.Connection
module: azure.ai.projects.models
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
document_id: f1a78158-260f-1ad2-f5a4-2830c3253c8a
document_version_independent_id: 6a79cf40-54c5-22e9-a2eb-687e77db44b9
updated_at: 2026-08-25T11:09:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-autogen/azure-ai-projects/azure.ai.projects.models.Connection.yml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/e74658802e80af785c26a422557b6189451805b9/docs-ref-autogen/azure-ai-projects/azure.ai.projects.models.Connection.yml
git_commit_id: e74658802e80af785c26a422557b6189451805b9
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
in_right_rail: h2h3
page_type: python
page_kind: class
description: 'Response from the list and get connections operations. '
toc_rel: ../_splitted/azure.ai.projects/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
search.mshattr.devlang: python
asset_id: api/azure-ai-projects/azure.ai.projects.models.connection
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-autogen/azure-ai-projects/az
... (truncated)
```

### ✅ 8. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.models.connection?view=azure-python"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-python/main/sdk/ai/azure-ai-projects/samples/connections/sample_connections.py:
# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------

"""
DESCRIPTION:
    Given an AIProjectClient, this sample demonstrates how to use the synchronous
    `.connections` methods to enumerate the properties of all connections
    and get the properties of a connection by its name.

USAGE:
    python sample_connections.py

    Before running the sample:

    pip install "azure-ai-projects>=2.0.0" python-dotenv

    Set these environment variables with your own values:
    1) FOUNDRY_PROJECT_ENDPOINT - Required. The Azure AI Project endpoint, as found in the overview page of your
       Microsoft Foundry project.
    2) CONNECTION_NAME - The name of a connection, as found in the "Connected resources" tab
       in the Management Center of your Microsoft Foundry project.
"""

import os
from dotenv import load_dotenv
from azure.identity import DefaultAzureCredential
from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import ConnectionType

load_dotenv()

endpoint = os.environ["FOUNDRY_PROJECT_ENDPOINT"]
connection_name = os.environ["CONNECTION_NAME"]

with (
    DefaultAzureCredential() as credential,
    AIProjectClient(endpoint=endpoint, credential=credential) as project_client,
):
    print("List all connections:")
    for connection in project_client.connections.list():
        print(connection)

    print("List all connections of a particular type:")
    for connection in project_client.connections.list(
        connection_type=ConnectionType.AZURE_OPEN_AI,
    ):
        print(connection)

    print("Get the default connection of a particular type, without its credentials:")
    connection = project_clie
... (truncated)
```

### ✅ 9. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: app.py\n+import os\n+\n+from azure.ai.projects import AIProjectClient\n+from azure.ai.projects.models import Connection, ModelDeployment\n+from azure.identity import DefaultAzureCredential\n+\n+\n+def print_connection(connection: Connection) -\u003e None:\n+    print(f\"  Name: {connection.name}\")\n+    print(f\"  Type: {connection.type}\")\n+    print(f\"  Target: {connection.target}\")\n+    print(f\"  Default: {connection.is_default}\")\n+\n+\n+def print_model_deployment(deployment: ModelDeployment) -\u003e None:\n+    print(f\"  Name: {deployment.name}\")\n+    print(f\"  Model publisher: {deployment.model_publisher}\")\n+    print(f\"  Model name: {deployment.model_name}\")\n+    print(f\"  Model version: {deployment.model_version}\")\n+\n+\n+def required_environment_variable(name: str) -\u003e str:\n+    value = os.environ.get(name)\n+    if not value:\n+        raise RuntimeError(f\"Required environment variable {name} is not set.\")\n+    return value\n+\n+\n+def main() -\u003e None:\n+    endpoint = required_environment_variable(\"FOUNDRY_PROJECT_ENDPOINT\")\n+    connection_name = required_environment_variable(\"CONNECTION_NAME\")\n+    deployment_name = required_environment_variable(\"DEPLOYMENT_NAME\")\n+\n+    with (\n+        DefaultAzureCredential() as credential,\n+        AIProjectClient(endpoint=endpoint, credential=credential) as project_client,\n+    ):\n+        print(\"Project connections\")\n+        print(\"===================\")\n+        for connection in project_client.connections.list():\n+            print_connection(connection)\n+            print()\n+\n+        print(f\"Connection: {connection_name}\")\n+        print(\"=\" * (12 + len(connection_name)))\n+        connection = project_client.connections.get(\n+            connection_name,\n+            include_credentials=False,\n+        )\n+        print_connection(connection)\n+        print()\n+\n+        print(\"Model deployments\")\n+        print(\"=================\")\n+        for deployment in project_client.deployments.list():\n+            if isinstance(deployment, ModelDeployment):\n+                print_model_deployment(deployment)\n+                print()\n+\n+        print(f\"Deployment: {deployment_name}\")\n+        print(\"=\" * (12 + len(deployment_name)))\n+        deployment = project_client.deployments.get(deployment_name)\n+        if not isinstance(deployment, ModelDeployment):\n+            raise TypeError(\n+                f\"Deployment {deployment_name!r} is not a model deployment \"\n+                f\"(received {type(deployment).__name__}).\"\n+            )\n+        print_model_deployment(deployment)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-ai-projects\u003e=2.0.0,\u003c3.0.0\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n*** Add File: README.md\n+# Microsoft Foundry project inventory\n+\n+This console application uses the synchronous `azure-ai-projects` client to list\n+project connections and model deployments, then retrieve one named connection\n+and one named model deployment. Authentication uses `DefaultAzureCredential`.\n+\n+## Install\n+\n+Requires Python 3.10 or later.\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+python -m pip install -r requirements.txt\n+```\n+\n+## Configure and run\n+\n+Set the project endpoint shown on the Microsoft Foundry project's overview page,\n+plus an existing connection name and deployment name:\n+\n+```powershell\n+$env:FOUNDRY_PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:CONNECTION_NAME = \"\u003cconnection-name\u003e\"\n+$env:DEPLOYMENT_NAME = \"\u003cdeployment-name\u003e\"\n+python app.py\n+```\n+\n+`DefaultAzureCredential` supports local Azure CLI sign-in and other standard\n+Azure Identity credential sources. The application never requests connection\n+credentials.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-baseline-1423089613/app.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-baseline-1423089613/app.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-baseline-1423089613/app.py
@@ -1,0 +1,73 @@
+import os
+
+from azure.ai.projects import AIProjectClient
+from azure.ai.projects.models import Connection, ModelDeployment
+from azure.identity import DefaultAzureCredential
+
+
+def print_connection(connection: Connection) -> None:
+    print(f"  Name: {connection.name}")
+    print(f"  Type: {connection.type}")
+    print(f"  Target: {connection.target}")
+    print(f"  Default: {connection.is_default}")
+
+
+def print_model_deployment(deployment: ModelDeployment) -> None:
+    print(f"  Name: {deployment.name}")
+    print(f"  Model publisher: {deployment.model_publisher}")
+    print(f"  Model name: {deployment.model_name}")
+    print(f"  Model version: {deployment.model_version}")
+
+
+def required_environment_variable(name: str) -> str:
+    value = os.environ.get(name)
+    if not value:
+        raise RuntimeError(f"Required environment variable {name} is not set.")
+    return value
+
+
+def main() -> None:
+    endpoint = required_environment_variable("FOUNDRY_PROJECT_ENDPOINT")
+    connection_name = required_environment_variable("CONNECTION_NAME")
+    deployment_name = required_environment_variable("DEPLOYMENT_NAME")
+
+    with (
+        DefaultAzureCredential() as credential,
+        AIProjectClient(endpoint=endpoint, credential=credential) as project_client,
+    ):
+        print("Project connections")
+        print("===================")
+        for connection in project_client.connections.list():
+            print_connection(connection)
+ 
... (truncated)
```

### ✅ 10. powershell

**Input:**

```json
{"command":"python --version \u0026\u0026 python -m py_compile app.py \u0026\u0026 python -m pip show azure-ai-projects","description":"Validate Python app and SDK availability","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Python 3.13.9
Name: azure-ai-projects
Version: 2.5.0
Summary: Microsoft Corporation Azure AI Projects Client Library for Python
Home-page: 
Author: 
Author-email: Microsoft Corporation <azpysdkhelp@microsoft.com>
License-Expression: MIT
Location: C:\Users\weidxu\AppData\Roaming\Python\Python313\site-packages
Requires: azure-core, azure-identity, azure-storage-blob, isodate, openai, typing-extensions
Required-by:
<shellId: 0 completed with exit code 0>
```

### ✅ 11. powershell

**Input:**

```json
{"command":"python -c \"import inspect; from azure.ai.projects.operations import ConnectionsOperations, DeploymentsOperations; from azure.ai.projects.models import Connection, ModelDeployment; print(inspect.signature(ConnectionsOperations.list)); print(inspect.signature(ConnectionsOperations.get)); print(inspect.signature(DeploymentsOperations.list)); print(inspect.signature(DeploymentsOperations.get)); print(Connection.__name__, ModelDeployment.__name__)\"","description":"Verify SDK operations and model types","initial_wait":30,"mode":"sync"}
```

**Output:**

```
(self, *, connection_type: Union[str, azure.ai.projects.models._enums.ConnectionType, NoneType] = None, default_connection: Optional[bool] = None, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.Connection']
(self, name: str, *, include_credentials: Optional[bool] = False, **kwargs: Any) -> azure.ai.projects.models._models.Connection
(self, *, model_publisher: Optional[str] = None, model_name: Optional[str] = None, deployment_type: Union[str, azure.ai.projects.models._enums.DeploymentType, NoneType] = None, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.Deployment']
(self, name: str, **kwargs: Any) -> azure.ai.projects.models._models.Deployment
Connection ModelDeployment
<shellId: 1 completed with exit code 0>
```

### ✅ 12. powershell

**Input:**

```json
{"command":"@'\nimport os\nfrom unittest.mock import MagicMock, patch\nfrom azure.ai.projects.models import Connection, ModelDeployment\nimport app\n\nos.environ.update({\n    \"FOUNDRY_PROJECT_ENDPOINT\": \"https://example.invalid/api/projects/test\",\n    \"CONNECTION_NAME\": \"storage\",\n    \"DEPLOYMENT_NAME\": \"chat\",\n})\nconnection = Connection(name=\"storage\", id=\"1\", type=\"AzureBlob\", target=\"https://example.invalid\", is_default=True, credentials=None, metadata={})\ndeployment = ModelDeployment(name=\"chat\", type=\"ModelDeployment\", model_name=\"gpt-test\", model_version=\"1\", model_publisher=\"Test\", capabilities={}, sku={})\nclient = MagicMock()\nclient.__enter__.return_value = client\nclient.connections.list.return_value = iter([connection])\nclient.connections.get.return_value = connection\nclient.deployments.list.return_value = iter([deployment])\nclient.deployments.get.return_value = deployment\ncredential = MagicMock()\ncredential.__enter__.return_value = credential\nwith patch.object(app, \"DefaultAzureCredential\", return_value=credential), patch.object(app, \"AIProjectClient\", return_value=client):\n    app.main()\nclient.connections.get.assert_called_once_with(\"storage\", include_credentials=False)\nclient.deployments.get.assert_called_once_with(\"chat\")\nprint(\"OFFLINE_VALIDATION_OK\")\n'@ | python -","description":"Run offline mocked application validation","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Project connections
===================
  Name: storage
  Type: ConnectionType.AZURE_BLOB_STORAGE
  Target: https://example.invalid
  Default: True

Connection: storage
===================
  Name: storage
  Type: ConnectionType.AZURE_BLOB_STORAGE
  Target: https://example.invalid
  Default: True

Model deployments
=================
  Name: chat
  Model publisher: Test
  Model name: gpt-test
  Model version: 1

Deployment: chat
================
  Name: chat
  Model publisher: Test
  Model name: gpt-test
  Model version: 1
OFFLINE_VALIDATION_OK
<shellId: 2 completed with exit code 0>
```

### ✅ 13. powershell

**Input:**

```json
{"command":"if (Test-Path -LiteralPath '.\\__pycache__') { Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force }","description":"Remove Python cache artifacts","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `app.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then create the synchronous SDK app, dependency manifest, and run documentation.

The workspace is empty. I’m confirming the current `azure-ai-projects` connection and deployment APIs before writing the runnable files.

The client operations are confirmed; I’m checking the exact typed model attributes so the output and type checks match the stable SDK.

The stable SDK exposes `Connection.is_default` and `ModelDeployment.model_*` fields. I’m writing the app with explicit credential exclusion and runtime model-deployment checks.

The files are in place. I’m validating syntax and the installed SDK surface locally without contacting Azure.

The installed SDK signatures match the implementation. I’m running an offline mocked client pass to verify pageable iteration, typed output, and credential-free retrieval end to end.

Created a complete synchronous Python console application:

- `app.py` — lists and retrieves typed project connections and model deployments
- `requirements.txt` — includes `azure-ai-projects` and `azure-identity`
- `README.md` — concise PowerShell installation, configuration, and run commands

Connection retrieval explicitly uses `include_credentials=False`, and named deployment retrieval rejects non-model deployments.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed. The core SDK workflow is implemented, but Azure SDK exception handling is absent.

### Strengths

- Uses the synchronous client and pageable iteration, requests the named connection without credentials, and validates that the named deployment is a ModelDeployment.

### Issues

- No handling for HttpResponseError or subclasses around Azure SDK operations.

## Grader Results

- project-resource-inventory.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Creates `AIProjectClient` for `FOUNDRY_PROJECT_ENDPOINT`.: Pass
      - Iterates every result from `project_client.connections.list()` and reads typed: Pass
      - Calls `project_client.connections.get` for `CONNECTION_NAME` without credentials: Pass
      - Iterates every result from `project_client.deployments.list()`.: Pass
      - Uses `isinstance(..., ModelDeployment)` before printing publisher, model name, and: Pass
      - Calls `project_client.deployments.get` for `DEPLOYMENT_NAME` and rejects a result: Pass
      - Does not replace pageable iteration with a single assumed page.: Pass
      - Does not select resources only from the locally enumerated results instead of: Pass
      - Does not request or display connection credentials.: Pass
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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 5.0000** | **62.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id ai-projects-dp-python-project-resource-inventory --config python-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
