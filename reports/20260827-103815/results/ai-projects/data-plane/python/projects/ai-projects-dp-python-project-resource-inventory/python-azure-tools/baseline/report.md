# Evaluation Report: ai-projects-dp-python-project-resource-inventory

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 309.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-projects-dp-python-project-resource-inventory` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 13/16 |
| Duration | 309.4s |
| Timestamp | 2026-08-27T02:45:35Z |
| Files Generated | 3 |
| Event Count | 4739 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 202.0s |
| Review | 106.2s |
| **Total** | **309.4s** |

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
| Token Usage | in=281181 out=3887 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. glob — 7153ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 6068ms

**Input:**

```json
{"pattern":"**/.*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. rg — 8219ms

**Input:**

```json
{"-n":true,"glob":"*.py","head_limit":100,"output_mode":"content","pattern":"azure-ai-projects|AIProjectClient|FOUNDRY_PROJECT_ENDPOINT"}
```

**Output:**

```
No matches found.
```

### ✅ 4. web_search (via github-mcp-server) — 62301ms

**Input:**

```json
{"query":"Current Python azure-ai-projects synchronous AIProjectClient connections list get include_credentials deployments list get ModelDeployment API 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To list model deployments and connections using the `azure-ai-projects` package in the Azure SDK for Python, you can use the sample scripts in the official Azure SDK repository.\n\nKey pointers from the repo:\n- **Model Deployments**: The `samples/deployments/sample_deployments.py` script demonstrates how to interact with model deployments in your Azure AI Foundry project.\n- **Connections**: You can enumerate connected Azure resources using operations exposed in the `.connections` resource group.\n\n### Listing Model Deployments Example\n\nIn [sample_deployments.py](https://github.com/Azure/azure-sdk-for-python/blob/main/sdk/ai/azure-ai-projects/samples/deployments/sample_deployments.py), the following environment variables are required:\n- `FOUNDRY_PROJECT_ENDPOINT`: Your Azure AI Project endpoint.\n- `FOUNDRY_MODEL_NAME`: The name of the model deployment you wish to retrieve.\n\nThe sample includes code to list deployments and filter them by model name or publisher:\n\n```python\n# Required: FOUNDRY_PROJECT_ENDPOINT, FOUNDRY_MODEL_NAME\n# Optional: MODEL_PUBLISHER, MODEL_NAME\n\nimport os\nfrom azure.ai.projects import AIProjectsClient\n\nendpoint = os.getenv(\"FOUNDRY_PROJECT_ENDPOINT\")\nmodel_name = os.getenv(\"FOUNDRY_MODEL_NAME\")\nclient = AIProjectsClient(endpoint)\n\ndeployments = client.deployments.list()\nfor deployment in deployments:\n    print(deployment)\n```\nYou can adjust how you filter and display them according to your needs【3:3†source】.\n\n### Listing Connections\n\nThe repo exposes `.connections` operations to enumerate Azure resources connected to your Foundry project. You can check the implementation in the [operations/_patch_connections.py](https://github.com/Azure/azure-sdk-for-python/blob/main/sdk/ai/azure-ai-projects/azure/ai/projects/operations/_patch_connections.py) file.\n\n### More Samples\nYou can find further usage examples and explore other scenarios in the [samples/deployments folder]
... (truncated)
```

### ✅ 5. web_search (via github-mcp-server) — 59990ms

**Input:**

```json
{"query":"site:learn.microsoft.com python api azure-ai-projects AIProjectClient deployments connections ModelDeployment ConnectionType is_default get_connection 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here are the relevant Python API details for working with Azure AI Projects and the `AIProjectClient`, focusing on deployments, connections, and related classes, based on the official Microsoft documentation up to 2026:\n\n---\n\n### 1. **AIProjectClient Instantiation**\nYou use `AIProjectClient` to work with your Azure AI Foundry (Projects) resources:\n\n```python\nfrom azure.ai.projects.aio import AIProjectClient  # For async\nfrom azure.identity.aio import DefaultAzureCredential\n\nclient = AIProjectClient(\n    endpoint=\"https://<your-project-name>.<region>.inference.ai.azure.com\",\n    credential=DefaultAzureCredential(),\n    allow_preview=True  # if preview features are needed\n)\n```\nDo **not** use `from_connection_string` as it is deprecated. Use the project endpoint as described above【3:3†source】【3:9†source】.\n\n---\n\n### 2. **ModelDeployment and Deployment Classes**\n- Both `ModelDeployment` and `Deployment` represent deployed AI models and their metadata:\n\n```python\nfrom azure.ai.projects.models import ModelDeployment\n\n# Attributes include: name, type, model_name, model_version, model_publisher, capabilities, sku, connection_name\ndeployment = ModelDeployment(...)  # Usually received from listing or getting resources\nprint(deployment.name)\n```\nSee all attributes and methods described, such as `.as_dict()`, `.get()`, etc.【3:1†source】【3:4†source】.\n\n---\n\n### 3. **Connection and ConnectionType**\n- The `Connection` class represents configured Azure or external resources your project can integrate with; each has a `type` (as `str` or `ConnectionType`) and an `is_default` boolean.\n\n```python\nfrom azure.ai.projects.models import Connection\n\n# Example access\nconnection = ...  # From listing or get operation\nprint(connection.is_default)\nprint(connection.type)   # e.g., \"AzureOpenAI\", \"AzureBlob\"\n```\n`is_default` tells you if this is the \"default\" connection of its kind 
... (truncated)
```

### ✅ 6. web_search (via github-mcp-server) — 57831ms

**Input:**

```json
{"query":"site:github.com/Azure/azure-sdk-for-python azure-ai-projects samples connections deployments list model deployment Python"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s a concise guide for using the current (2026) **Python Azure AI Projects SDK** (azure-ai-projects, part of the Microsoft Foundry SDK) for your requested workflows:\n\n---\n\n### 1. **AIProjectClient Synchronous Usage**\n\n**Initialize the client (synchronous):**\n```python\nfrom azure.ai.projects import AIProjectClient\n\nclient = AIProjectClient(endpoint=\"https://YOUR_PROJECT_ENDPOINT\", credential=\"YOUR_CREDENTIAL\")\n```\nIf you need to use preview features:\n```python\nclient = AIProjectClient(endpoint=\"...\", credential=\"...\", allow_preview=True)\n```\n\n---\n\n### 2. **Enumerate Connections**  \n(inc. listing, getting details, and including credentials)\n\n#### **List All Connections**\n```python\nconnections = client.connections.list()  # Returns an iterator of Connection objects\nfor conn in connections:\n    print(conn)\n```\nYou may filter by type (e.g., \"AzureOpenAI\"):\n```python\nconnections = client.connections.list(connection_type=\"AzureOpenAI\")\n```\n\n#### **Get Specific Connection (with credentials)**\n```python\nconn_name = \"YOUR_CONNECTION_NAME\"\nconnection = client.connections.get(conn_name, include_credentials=True)\nprint(connection.credentials)  # Access credentials if included\n```\nYou can also get the default connection for a given type:\n```python\nconnection = client.connections.get_default(\"AzureOpenAI\", include_credentials=True)\n```\n**Documentation for these methods:**  \n- [`ConnectionsOperations`](https://learn.microsoft.com/en-us/python/api/azure-ai-projects/azure.ai.projects.operations.connectionsoperations?view=azure-python)【3:3†source】\n- [Sample code (`sample_connections.py`, GitHub)](https://github.com/Azure/azure-sdk-for-python/blob/main/sdk/ai/azure-ai-projects/samples/connections/sample_connections.py)【3:4†source】\n\n---\n\n### 3. **Enumerate Deployments (Model Deployment API)**\n\n#### **List All Deployments**\n```python\ndeployments = client.deploymen
... (truncated)
```

### ✅ 7. github-mcp-server-get_file_contents (via github-mcp-server) — 10657ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/samples/connections/sample_connections.py","ref":"refs/heads/main","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: dea352a197634de0a137fade32566a8178c5d69e)# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) Python Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

[build-system]
requires = ["setuptools>=77.0.3", "wheel"]
build-backend = "setuptools.build_meta"

[project]
name = "azure-ai-projects"
authors = [
  { name = "Microsoft Corporation", email = "azpysdkhelp@microsoft.com" },
]
description = "Microsoft Corporation Azure AI Projects Client Library for Python"
license = "MIT"
classifiers = [
    "Development Status :: 5 - Production/Stable",
    "Programming Language :: Python",
    "Programming Language :: Python :: 3 :: Only",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.10",
    "Programming Language :: Python :: 3.11",
    "Programming Language :: Python :: 3.12",
    "Programming Language :: Python :: 3.13",
    "Programming Language :: Python :: 3.14",
]
requires-python = ">=3.10"
keywords = ["azure", "azure sdk"]

dependencies = [
    "isodate>=0.6.1",
    "azure-core>=1.37.0",
    "typing-extensions>=4.11",
    "azure-identity>=1.15.0",
    "openai>=3.0.0",
    "azure-storage-blob>=12.15.0",
]
dynamic = [
"version", "readme"
]

[project.urls]
repository = "https://aka.ms/azsdk/azure-ai-projects-v2/python/code"

[tool.setuptools.dynamic]
version = {attr = "azure.ai.projects._version.VERSION"}
readme = {file = ["README.md", "CHANGELOG.md"], content-type = "text/markdown"}

[tool.setuptools.packages.find]
exclude = [
    "azure.ai",
    "azure",
    "doc*",
    "generated_samples*",
    "generated_tests*",
    "samples*",
    "tests*",
]

[tool.setuptools.
... (truncated)
```

### ✅ 8. github-mcp-server-get_file_contents (via github-mcp-server) — 10267ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/samples/deployments/sample_deployments.py","ref":"refs/heads/main","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: 3970566daddf279b8a5072f6ef31e57fcf356f0d)# pylint: disable=line-too-long,useless-suppression
# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------
"""Customize generated code here.

Follow our quickstart for examples: https://aka.ms/azsdk/python/dpcodegen/python/customize
"""

from functools import wraps
import inspect
from typing import Any, Callable, List
from ..models._patch import _FOUNDRY_FEATURES_HEADER_NAME, _BETA_OPERATION_FEATURE_HEADERS, _has_header_case_insensitive
from ._patch_agents import AgentsOperations, BetaAgentsOperations
from ._patch_datasets import BetaDatasetsOperations, DatasetsOperations
from ._patch_evaluators import BetaEvaluatorsOperations
from ._patch_evaluation_rules import EvaluationRulesOperations
from ._patch_telemetry import TelemetryOperations
from ._patch_connections import ConnectionsOperations
from ._patch_memories import BetaMemoryStoresOperations
from ._patch_models import BetaModelsOperations
from ._operations import (
    BetaEvaluationTaxonomiesOperations,
    BetaInsightsOperations,
    BetaOperations as GeneratedBetaOperations,
    BetaRedTeamsOperations,
    BetaRoutinesOperations,
    BetaSchedulesOperations,
    BetaSkillsOperations,
)


def _method_accepts_keyword_headers(method: Callable[..., Any]) -> bool:
    try:
        signature = inspect.signature(method)
    except (TypeError, ValueError):
        return False

    for parameter in signature.parameters.values():
        if parameter.name == "headers":
            return True
        if parameter.kind == inspect.Parameter.VAR_KEYWORD:
            return True

    return False


class _OperationMethodHeaderProxy:
    """Proxy that injects the Foundry-Features header into public operation method calls."""

    def __init__(self, operation: Any, foundry_features_value: str):
        object.__setattr__(self, "_operation", operation
... (truncated)
```

### ✅ 9. github-mcp-server-get_file_contents (via github-mcp-server) — 9868ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/azure/ai/projects/operations/_patch.py","ref":"refs/heads/main","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: 2d46858fc7c439cdbef8267d927a4c60618fe9c6)# ------------------------------------
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
    for deployment in project_client.deployments.list(model_publisher=model_publisher):
        print(deployment)

    print(f"List all deployments of model `{model_name}`:")
    for deployment in project_client.deployments.list(model_name=model_na
... (truncated)
```

### ✅ 10. github-mcp-server-get_file_contents (via github-mcp-server) — 9479ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/pyproject.toml","ref":"refs/heads/main","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: c9fcbf78bc8cd2430dc767ac03fde6d693fa7e3b)# ------------------------------------
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
    connection = project_client.connections.get_default(connection_type=ConnectionType.AZURE_OPEN_AI)
    print(connection)

    print("Get the default connection of a particular type, with it
... (truncated)
```

### ✅ 11. github-mcp-server-search_code (via github-mcp-server) — 5882ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/ai/azure-ai-projects/azure/ai/projects/models class Connection"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 12. github-mcp-server-search_code (via github-mcp-server) — 5504ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/ai/azure-ai-projects/azure/ai/projects/models class ModelDeployment"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/azure-ai-projects/azure/ai/projects/models/_enums.py","text_matches":[{"fragment":"\n    MODEL_DEPLOYMENT = \"ModelDeployment\"\n    \"\"\"Model deployment.\"\"\"","matches":[{"indices":[25,40],"text":"ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-projects/azure/ai/projects/models/_enums.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"},{"fragment":"\nclass EvaluationLevel(str, Enum, metaclass=CaseInsensitiveEnumMeta):\n    \"\"\"The level at which evaluation is performed.\"\"\"","matches":[{"indices":[1,6],"text":"class"},{"indices":[38,43],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-projects/azure/ai/projects/models/_enums.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]}],"total_count":1}
```

### ✅ 13. github-mcp-server-search_code (via github-mcp-server) — 5145ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/ai/azure-ai-projects/azure/ai/projects/operations class DeploymentsOperations"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/azure-ai-projects/azure/ai/projects/models/_enums.py","text_matches":[{"fragment":"    \"\"\"BROWSER.\"\"\"\n\n\nclass ConnectionType(str, Enum, metaclass=CaseInsensitiveEnumMeta):\n    \"\"\"The Type (or category) of the connection.\"\"\"\n\n    AZURE_OPEN_AI = \"AzureOpenAI\"","matches":[{"indices":[21,26],"text":"class"},{"indices":[27,37],"text":"Connection"},{"indices":[57,62],"text":"class"},{"indices":[126,136],"text":"connection"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-projects/azure/ai/projects/models/_enums.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]},{"path":"sdk/ai/azure-ai-projects/azure/ai/projects/models/_patch.py","text_matches":[{"fragment":"\nclass CustomCredential(CustomCredentialGenerated, discriminator=\"CustomKeys\"):\n    \"\"\"Custom credential definition.","matches":[{"indices":[1,6],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-projects/azure/ai/projects/models/_patch.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"},{"fragment":"        # Although the issue was filed on C# Projects SDK, the same problem exists in Python SDK.\n        # Assume your Foundry project has a connection of type `Custom`, named \"test_custom_connection\",\n        # and you defined two public and two secrete (private) keys. When you get the connection, the response","matches":[{"indices":[142,152],"text":"connection"},{"indices":[190,200],"text":"connection"},{"indices":[289,299],"text":"connection"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-projects/azure/ai/projects/models/_patch.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]}],"total_count":2}
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server) — 6225ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-projects/azure/ai/projects/operations/_patch_connections.py","ref":"refs/heads/main","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: 581d1367151681dbc5462603efe147f5987a5530)# pylint: disable=line-too-long,useless-suppression
# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------
"""Customize generated code here.

Follow our quickstart for examples: https://aka.ms/azsdk/python/dpcodegen/python/customize
"""

from typing import Optional, Any, Union
from azure.core.tracing.decorator import distributed_trace
from ._operations import ConnectionsOperations as ConnectionsOperationsGenerated
from ..models._models import Connection
from ..models._enums import ConnectionType


class ConnectionsOperations(ConnectionsOperationsGenerated):
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.ai.projects.AIProjectClient`'s
        :attr:`connections` attribute.
    """

    @distributed_trace
    def get(self, name: str, *, include_credentials: Optional[bool] = False, **kwargs: Any) -> Connection:
        """Get a connection by name.

        :param name: The name of the connection. Required.
        :type name: str
        :keyword include_credentials: Whether to include credentials in the response. Default is False.
        :paramtype include_credentials: bool
        :return: Connection. The Connection is compatible with MutableMapping
        :rtype: ~azure.ai.projects.models.Connection
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        if include_credentials:
            return super()._get_with_credentials(name, **kwargs)

        return super()._get(name, **kwargs)

    @distributed_trace
    def get_default(
        self, connection_type: Union[str, ConnectionType], *, include_credentials: Optional[bool] = False, **kwargs: Any
    ) -> Connection:
        """Get the default connection for a given connection type.

        :para
... (truncated)
```

### ✅ 15. apply_patch — 1074ms

**Input:**

```json
"*** Begin Patch\n*** Add File: app.py\n+\"\"\"Inspect connections and model deployments in a Microsoft Foundry project.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from typing import NoReturn\n+\n+from azure.ai.projects import AIProjectClient\n+from azure.ai.projects.models import Connection, ModelDeployment\n+from azure.identity import DefaultAzureCredential\n+\n+\n+def required_environment_variable(name: str) -\u003e str:\n+    \"\"\"Return a required, non-empty environment variable.\"\"\"\n+    value = os.environ.get(name, \"\").strip()\n+    if not value:\n+        raise ValueError(f\"Environment variable {name} must be set to a non-empty value.\")\n+    return value\n+\n+\n+def display_value(value: object) -\u003e object:\n+    \"\"\"Use an enum's wire value when the SDK returns an enum instance.\"\"\"\n+    return getattr(value, \"value\", value)\n+\n+\n+def print_connection(connection: Connection) -\u003e None:\n+    \"\"\"Print the requested connection metadata.\"\"\"\n+    print(f\"  Name: {connection.name}\")\n+    print(f\"  Type: {display_value(connection.type)}\")\n+    print(f\"  Target: {connection.target}\")\n+    print(f\"  Default: {connection.is_default}\")\n+\n+\n+def print_model_deployment(deployment: ModelDeployment) -\u003e None:\n+    \"\"\"Print the requested model deployment metadata.\"\"\"\n+    print(f\"  Name: {deployment.name}\")\n+    print(f\"  Model publisher: {deployment.model_publisher}\")\n+    print(f\"  Model name: {deployment.model_name}\")\n+    print(f\"  Model version: {deployment.model_version}\")\n+\n+\n+def require_model_deployment(deployment: object, name: str) -\u003e NoReturn:\n+    \"\"\"Raise a useful error when a named deployment is not a model deployment.\"\"\"\n+    deployment_type = display_value(getattr(deployment, \"type\", type(deployment).__name__))\n+    raise TypeError(\n+        f\"Deployment {name!r} has type {deployment_type!r}; \"\n+        \"a model deployment is required.\"\n+    )\n+\n+\n+def main() -\u003e None:\n+    endpoint = required_environment_variable(\"FOUNDRY_PROJECT_ENDPOINT\")\n+    connection_name = required_environment_variable(\"CONNECTION_NAME\")\n+    deployment_name = required_environment_variable(\"DEPLOYMENT_NAME\")\n+\n+    with (\n+        DefaultAzureCredential() as credential,\n+        AIProjectClient(endpoint=endpoint, credential=credential) as project_client,\n+    ):\n+        print(\"Project connections\")\n+        for connection in project_client.connections.list():\n+            print_connection(connection)\n+\n+        print(f\"\\nConnection {connection_name!r}\")\n+        connection = project_client.connections.get(\n+            connection_name,\n+            include_credentials=False,\n+        )\n+        print_connection(connection)\n+\n+        print(\"\\nModel deployments\")\n+        for deployment in project_client.deployments.list():\n+            if isinstance(deployment, ModelDeployment):\n+                print_model_deployment(deployment)\n+\n+        print(f\"\\nModel deployment {deployment_name!r}\")\n+        deployment = project_client.deployments.get(deployment_name)\n+        if not isinstance(deployment, ModelDeployment):\n+            require_model_deployment(deployment, deployment_name)\n+        print_model_deployment(deployment)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-ai-projects\u003e=2.0.0,\u003c3.0.0\n+azure-identity\u003e=1.15.0,\u003c2.0.0\n*** Add File: .env.example\n+FOUNDRY_PROJECT_ENDPOINT=https://your-resource.services.ai.azure.com/api/projects/your-project\n+CONNECTION_NAME=your-connection-name\n+DEPLOYMENT_NAME=your-deployment-name\n*** Add File: README.md\n+# Microsoft Foundry project inventory\n+\n+This console application uses the synchronous `azure-ai-projects` client and\n+`DefaultAzureCredential` to inspect a Microsoft Foundry project's connections\n+and model deployments. It reads configuration from the process environment and\n+never requests connection credentials.\n+\n+## Install\n+\n+Python 3.10 or later is required.\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+python -m pip install -r requirements.txt\n+```\n+\n+Authenticate locally with a credential supported by `DefaultAzureCredential`,\n+then set the project values:\n+\n+```powershell\n+$env:FOUNDRY_PROJECT_ENDPOINT = \"https://your-resource.services.ai.azure.com/api/projects/your-project\"\n+$env:CONNECTION_NAME = \"your-connection-name\"\n+$env:DEPLOYMENT_NAME = \"your-deployment-name\"\n+```\n+\n+## Run\n+\n+```powershell\n+python app.py\n+```\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-baseline-4196165523/app.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-baseline-4196165523/app.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-baseline-4196165523/app.py
@@ -1,0 +1,85 @@
+"""Inspect connections and model deployments in a Microsoft Foundry project."""
+
+from __future__ import annotations
+
+import os
+from typing import NoReturn
+
+from azure.ai.projects import AIProjectClient
+from azure.ai.projects.models import Connection, ModelDeployment
+from azure.identity import DefaultAzureCredential
+
+
+def required_environment_variable(name: str) -> str:
+    """Return a required, non-empty environment variable."""
+    value = os.environ.get(name, "").strip()
+    if not value:
+        raise ValueError(f"Environment variable {name} must be set to a non-empty value.")
+    return value
+
+
+def display_value(value: object) -> object:
+    """Use an enum's wire value when the SDK returns an enum instance."""
+    return getattr(value, "value", value)
+
+
+def print_connection(connection: Connection) -> None:
+    """Print the requested connection metadata."""
+    print(f"  Name: {connection.name}")
+    print(f"  Type: {display_value(connection.type)}")
+    print(f"  Target: {connection.target}")
+    print(f"  Default: {connection.is_default}")
+
+
+def print_model_deployment(deployment: ModelDeployment) -> None:
+    """Print the requested model deployment metadata."""
+    print(f"  Name: {deployment.name}")
+    print(f"  Model publisher: {deployment.model_publisher}")
+    print(f"  Model name: {deployment.model_name}")
+    print(f"  Model version: {deployment.model_version}")
+
+
+def require_model_deployment(deployment: object, n
... (truncated)
```

### ✅ 16. powershell — 13742ms

**Input:**

```json
{"command":"python -m pip install --quiet -r requirements.txt \u0026\u0026 python -m compileall -q app.py \u0026\u0026 python -c \"from inspect import signature; from azure.ai.projects import AIProjectClient; from azure.ai.projects.models import Connection, ModelDeployment; print('connections.get', signature(AIProjectClient.connections.fget) if isinstance(AIProjectClient.connections, property) else 'available'); print('imports and compilation succeeded')\"","description":"Install dependencies and validate application imports","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Traceback (most recent call last):
  File "<string>", line 1, in <module>
    from inspect import signature; from azure.ai.projects import AIProjectClient; from azure.ai.projects.models import Connection, ModelDeployment; print('connections.get', signature(AIProjectClient.connections.fget) if isinstance(AIProjectClient.connections, property) else 'available'); print('imports and compilation succeeded')
                                                                                                                                                                                                                                       ^^^^^^^^^^^^^^^^^^^^^^^^^^^
AttributeError: type object 'AIProjectClient' has no attribute 'connections'
<shellId: 0 completed with exit code 1>
```

### ✅ 17. powershell — 4281ms

**Input:**

```json
{"command":"@'\nfrom inspect import signature\nfrom azure.ai.projects import AIProjectClient\nfrom azure.ai.projects.models import Connection, ModelDeployment\n\nclass Credential:\n    def get_token(self, *scopes, **kwargs):\n        raise AssertionError(\"No network authentication expected\")\n    def close(self):\n        pass\n\nclient = AIProjectClient(\n    endpoint=\"https://example.services.ai.azure.com/api/projects/example\",\n    credential=Credential(),\n)\nprint(\"connections.list\", signature(client.connections.list))\nprint(\"connections.get\", signature(client.connections.get))\nprint(\"deployments.list\", signature(client.deployments.list))\nprint(\"deployments.get\", signature(client.deployments.get))\nprint(\"Connection fields\", list(Connection._attr_to_rest_field))\nprint(\"ModelDeployment fields\", list(ModelDeployment._attr_to_rest_field))\nclient.close()\n'@ | python -","description":"Verify installed SDK operation and model interfaces","initial_wait":60,"mode":"sync"}
```

**Output:**

```
connections.list (*, connection_type: Union[str, azure.ai.projects.models._enums.ConnectionType, NoneType] = None, default_connection: Optional[bool] = None, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.Connection']
connections.get (name: str, *, include_credentials: Optional[bool] = False, **kwargs: Any) -> azure.ai.projects.models._models.Connection
deployments.list (*, model_publisher: Optional[str] = None, model_name: Optional[str] = None, deployment_type: Union[str, azure.ai.projects.models._enums.DeploymentType, NoneType] = None, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.Deployment']
deployments.get (name: str, **kwargs: Any) -> azure.ai.projects.models._models.Deployment
Traceback (most recent call last):
  File "<stdin>", line 19, in <module>
AttributeError: type object 'Connection' has no attribute '_attr_to_rest_field'
<shellId: 1 completed with exit code 1>
```

### ✅ 18. powershell — 3788ms

**Input:**

```json
{"command":"python -c \"from inspect import signature; from azure.ai.projects.models import Connection, ModelDeployment; print(signature(Connection)); print(signature(ModelDeployment)); print(Connection.__annotations__); print(ModelDeployment.__annotations__)\"","description":"Inspect SDK model constructors and fields","initial_wait":30,"mode":"sync"}
```

**Output:**

```
(*args: Any, **kwargs: Any) -> Self
(*args: Any, **kwargs: Any) -> None
{'name': <class 'str'>, 'id': <class 'str'>, 'type': typing.Union[str, ForwardRef('_models.ConnectionType')], 'target': <class 'str'>, 'is_default': <class 'bool'>, 'credentials': '_models.BaseCredentials', 'metadata': dict[str, str]}
{'type': typing.Literal[<DeploymentType.MODEL_DEPLOYMENT: 'ModelDeployment'>], 'model_name': <class 'str'>, 'model_version': <class 'str'>, 'model_publisher': <class 'str'>, 'capabilities': dict[str, str], 'sku': '_models.ModelDeploymentSku', 'connection_name': typing.Optional[str]}
<shellId: 2 completed with exit code 0>
```

### ✅ 19. powershell — 4014ms

**Input:**

```json
{"command":"@'\nimport contextlib\nimport io\nimport os\nfrom unittest.mock import patch\n\nimport app\nfrom azure.ai.projects.models import Connection, ConnectionType, ModelDeployment\n\nconnection = Connection(name=\"storage\", type=ConnectionType.AZURE_STORAGE_ACCOUNT, target=\"https://example.invalid\", is_default=True)\ndeployment = ModelDeployment(name=\"chat\", model_name=\"gpt-test\", model_version=\"1\", model_publisher=\"Test Publisher\")\n\nclass Operations:\n    def __init__(self, item):\n        self.item = item\n        self.get_calls = []\n    def list(self):\n        yield self.item\n    def get(self, name, **kwargs):\n        self.get_calls.append((name, kwargs))\n        return self.item\n\nclass Client:\n    def __init__(self, **kwargs):\n        self.connections = Operations(connection)\n        self.deployments = Operations(deployment)\n    def __enter__(self):\n        return self\n    def __exit__(self, *args):\n        pass\n\nclass Credential:\n    def __enter__(self):\n        return self\n    def __exit__(self, *args):\n        pass\n\nvalues = {\n    \"FOUNDRY_PROJECT_ENDPOINT\": \"https://example.invalid/api/projects/test\",\n    \"CONNECTION_NAME\": \"storage\",\n    \"DEPLOYMENT_NAME\": \"chat\",\n}\noutput = io.StringIO()\nwith patch.dict(os.environ, values, clear=True), patch.object(app, \"DefaultAzureCredential\", Credential), patch.object(app, \"AIProjectClient\", Client), contextlib.redirect_stdout(output):\n    app.main()\ntext = output.getvalue()\nfor expected in (\"Name: storage\", \"Type: AzureStorageAccount\", \"Target: https://example.invalid\", \"Default: True\", \"Name: chat\", \"Model publisher: Test Publisher\", \"Model name: gpt-test\", \"Model version: 1\"):\n    assert expected in text, expected\nprint(\"offline application flow succeeded\")\n'@ | python -","description":"Exercise application with local fake SDK operations","initial_wait":30,"mode":"sync"}
```

**Output:**

```
offline application flow succeeded
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `app.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace and existing project conventions, then implement the synchronous SDK application and validate it locally without contacting Azure.

The workspace is empty, so I’m confirming the current `azure-ai-projects` connection and deployment APIs before creating the minimal project structure.

The official SDK confirms `connections.list()/get()` and `deployments.list()/get()` are synchronous pageable operations, with `include_credentials=False` supported explicitly. I’m creating the app with model-type guards, environment validation, and concise setup documentation.

The dependencies installed and the application compiles; the first API introspection probe targeted an instance-created operation as a class property. I’m rerunning the check against a local client instance and exercising the full output path with fake pageable operations only.

Created a complete runnable synchronous Python application:

- `app.py` — enumerates and retrieves connections and model deployments with typed metadata.
- `requirements.txt` — includes `azure-ai-projects` and `azure-identity`.
- `.env.example` — documents required environment variables.
- `README.md` — provides concise PowerShell install, authentication, configuration, and run commands.

The named connection explicitly uses `include_credentials=False`, and the named deployment is rejected unless it is a `ModelDeployment`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- The application uses the synchronous SDK client and correctly avoids requesting connection credentials.

### Issues

- Azure SDK service errors are not caught or presented with controlled error handling.

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
      - tool_used: any tool (source=mcp, server=azure): Fail

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
