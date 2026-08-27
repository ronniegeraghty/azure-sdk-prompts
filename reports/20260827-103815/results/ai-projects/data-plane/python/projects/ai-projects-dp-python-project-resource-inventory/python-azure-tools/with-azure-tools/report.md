# Evaluation Report: ai-projects-dp-python-project-resource-inventory

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 440.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-projects-dp-python-project-resource-inventory` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 15/16 |
| Duration | 440.6s |
| Timestamp | 2026-08-27T02:38:15Z |
| Files Generated | 3 |
| Event Count | 6058 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 338.2s |
| Review | 101.1s |
| **Total** | **440.6s** |

## Configuration

- **name:** python-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-ai-projects-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=863793 out=6169 |
| Turn Count | 17 |

## Tool Calls

### ✅ 1. skill — 3874ms

**Input:**

```json
{"skill":"azure-ai-projects-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure AI Projects Python SDK (Foundry SDK)

Build AI applications on Microsoft Foundry using the `azure-ai-projects` SDK.

## Installation

```bash
pip install azure-ai-projects azure-identity
```

## Environment Variables

```bash
AZURE_AI_PROJECT_ENDPOINT="https://<resource>.services.ai.azure.com/api/projects/<project>"  # Required for all auth methods
AZURE_AI_MODEL_DEPLOYMENT_NAME="gpt-4o-mini"  # Required for all auth methods
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
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential
from azure.ai.projects import AIProjectClient

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
credential = DefaultAzureCredential()
# Or use a specific credential directly in production:
# See https://learn.microsoft.com/python/api/overview/azure/identity-readme?view=azure-py
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 3217ms

**Input:**

```json
{"intent":"Create a local-only synchronous Python console application using azure-ai-projects to inspect Microsoft Foundry project connections and model deployments without modifying Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-get_azure_bestpractices (via azure) — 2198ms

**Input:**

```json
{"command":"get_azure_bestpractices_ai_app","intent":"Get Microsoft Foundry code-generation guidance for a synchronous Python inventory console app.","parameters":{}}
```

**Output:**

```
{"status":200,"message":"","results":["# Best Practices for Building AI Apps with Azure AI Services\r\n\r\nThis guide covers best practices for:\r\n- Generating code for AI agent applications using Microsoft Foundry and any other Azure AI models\r\n- Building agents with Microsoft Foundry\r\n- Implementing chatbots, workflows, assistants, and multi-agent systems with Microsoft Agent Framework\r\n\r\n> Note: Understanding User Intent of Microsoft Foundry\r\n> \r\n> **Two Types of Operations:**\r\n> \r\n> **1. Resource Management** - Use `foundry` tool\r\n> - Keywords: \"create/deploy/provision agent/embedding IN Foundry\", \"set up agent/embedding resource\"\r\n> - Creates Microsoft Foundry agent resource via commands under `foundry`\r\n> \r\n> **2. Application Development** - Use this best practices tool\r\n> - Keywords: \"build/write/implement agent\", \"develop chatbot/assistant\", \"agent code\"\r\n> - Generates code using Microsoft Agent Framework SDK\r\n> - Returns code samples and implementation patterns\r\n> \r\n> **When unsure, ask:** \"Do you want to create an agent resource in Microsoft Foundry, or generate agent application code?\"\r\n\r\n## Key Principles\r\n**Defaults and Recommendations:**\r\n- **Default language**: If the user does not specify a programming language, use **.NET** as the default programming language.\r\n- **Recommended SDK**: Use **Microsoft Agent Framework** for building AI agents, chatbots, assistants, and multi-agent systems - it provides flexible orchestration, multi-agent patterns, and cross-platform support (.NET and Python)\r\n- **Select model first**: Model selection is REQUIRED for ANY AI implementation - always select the appropriate AI model before coding if the user does not specify a model\r\n\r\n**Process:**\r\n1. **Gather information first**: Always search / fetch / query for sufficient knowledge (quickstarts, tutorials, feature guides, code samples) BEFORE thinking or writing any code - this ensures high-quality code ge
... (truncated)
```

### ✅ 4. view — 9936ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\connections.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/deployments.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/deployments.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/deployments.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/deployments.md
@@ -1,135 +1,135 @@
 # Deployments Operations Reference
 
 ## Overview
 
 Deployments represent AI model deployments in your Azure AI Foundry project.
 
 ## List Deployments
 
 ### List All Deployments
 
 ```python
 deployments = project_client.deployments.list()
 for deployment in deployments:
     print(f"Name: {deployment.name}")
     print(f"Model: {deployment.model_name}")
     print(f"Publisher: {deployment.model_publisher}")
     print("---")
 ```
 
 ### Filter by Publisher
 
 ```python
 # List only OpenAI model deployments
 for deployment in project_client.deployments.list(model_publisher="OpenAI"):
     print(f"{deployment.name}: {deployment.model_name}")
 ```
 
 ### Filter by Model Name
 
 ```python
 # List deployments of a specific model
 for deployment in project_client.deployments.list(model_name="gpt-4o"):
     print(f"{deployment.name}: {deployment.model_version}")
 ```
 
 ## Get Deployment
 
 ```python
 from azure.ai.projects.models import ModelDeployment
 
 deployment = project_client.deployments.get("my-deployment-name")
 
 if isinstance(deployment, ModelDeployment):
     print(f"Type: {deployment.type}")
     print(f"Name: {deployment.name}")
     print(f"Model Name: {deployment.model_name}")
     print(f"Model Version: {deployment.model_version}")
     print(f"Model Publisher: {deployment.model_publisher}")
     
... (truncated)
```

### ✅ 5. view — 9560ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\deployments.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/connections.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/connections.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/connections.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/connections.md
@@ -1,204 +1,204 @@
 # Connections Operations Reference
 
 ## Overview
 
 Connections provide access to external Azure services like Azure OpenAI, Azure AI Search, Bing, and more.
 
 ## List Connections
 
 ### List All Connections
 
 ```python
 connections = project_client.connections.list()
 for conn in connections:
     print(f"Name: {conn.name}")
     print(f"Type: {conn.connection_type}")
     print(f"ID: {conn.id}")
     print("---")
 ```
 
 ### Filter by Connection Type
 
 ```python
 from azure.ai.projects.models import ConnectionType
 
 # List Azure OpenAI connections
 for conn in project_client.connections.list(
     connection_type=ConnectionType.AZURE_OPEN_AI
 ):
     print(f"Azure OpenAI: {conn.name}")
 
 # List Azure AI Search connections
 for conn in project_client.connections.list(
     connection_type=ConnectionType.AZURE_AI_SEARCH
 ):
     print(f"AI Search: {conn.name}")
 ```
 
 ## Connection Types
 
 ```python
 from azure.ai.projects.models import ConnectionType
 
 # Available connection types:
 # - ConnectionType.AZURE_OPEN_AI
 # - ConnectionType.AZURE_AI_SEARCH
 # - ConnectionType.AZURE_BLOB
 # - ConnectionType.AZURE_AI_SERVICES
 # - ConnectionType.API_KEY
 # - ConnectionType.COGNITIVE_SEARCH
 # - ConnectionType.COGNITIVE_SERVICE
 # - ConnectionType.CUSTOM
 ```
 
 ## Get Connection
 
 ### Get by Name
 
 ```python
... (truncated)
```

### ✅ 6. rg — 10305ms

**Input:**

```json
{"-n":true,"C":3,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\api-reference.md","pattern":"include_credentials|ModelDeployment|connection_type|is_default|model_publisher|model_name|model_version"}
```

**Output:**

```
Output too large to read at once (28.4 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787798348283-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787798332708-copilot-tool-output-aa5315b71f4e48c5968df19297e099c1.txt:5:    "# Best Practices for Building AI Apps with Azure AI Services\r\n\r\nThis guide covers best practices for:\r\n- Generating code for AI agent applications using Microsoft Foundry and any other Azure AI models\r\n- Building agents with Microsoft Foundry\r\n- Implementing chatbots, workflows, assistants, and multi-agent systems with Microsoft Agent Framework\r\n\r\n> Note: Understanding U
```

### ✅ 7. glob — 8803ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. glob — 8416ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-with-azure-tools-642855756","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. rg — 9177ms

**Input:**

```json
{"-i":true,"-n":true,"C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787798332708-copilot-tool-output-aa5315b71f4e48c5968df19297e099c1.txt","pattern":"DefaultAzureCredential|context manager|environment|exception|read-only|inventory"}
```

**Output:**

```
No matches found.
```

### ✅ 10. rg — 3874ms

**Input:**

```json
{"-i":true,"-n":true,"C":5,"head_limit":160,"output_mode":"content","paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\api-reference.md","pattern":"class .*Connection|ConnectionProperties|ConnectionDetails|target|is_default|ModelDeployment|Deployment"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-ai-projects-py\references\api-reference.md:48:client.deployments     # DeploymentsOperations
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-ai-projects-py\references\api-reference.md:446:    embedding_model_deployment="text-embedding-ada-002",
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-ai-projects-py\references\api-reference.md:1050:DeploymentType         # MODEL, EMBEDDING, etc.
```

### ✅ 11. azure-documentation (via azure) — 4638ms

**Input:**

```json
{"intent":"Find official Python azure-ai-projects API documentation for synchronous AIProjectClient connections list/get and deployments list/get, including model classes and fields for connection target/default and ModelDeployment.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 12. web_fetch — 4250ms

**Input:**

```json
{"max_length":3000,"url":"https://pypi.org/pypi/azure-ai-projects/json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-ai-projects/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com>","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.14"],"description":"# Azure AI Projects client library for Python\n\nThe AI Projects client library is part of the Microsoft Foundry SDK, and provides easy access to\nresources in your [Microsoft Foundry](https://ai.azure.com/) Project. Use it to:\n\n* **Create and run Agents** using methods on the `.agents` client property. This includes **Hosted Agents**, which let you run your own containerized agent runtime while using Microsoft Foundry for managed hosting and scaling.\n* **Enhance Agents with specialized tools and toolbox tools** such as:\n  * Agent-to-Agent (A2A)\n  * Azure AI Search\n  * Azure Functions\n  * Bing Custom Search (Preview)\n  * Bing Grounding\n  * Browser Automation (Preview)\n  * Code Interpreter\n  * Computer\n  * Computer Use (Preview)\n  * Fabric IQ (Preview)\n  * File Search\n  * Function Tool\n  * Image Generation\n  * Memory Search (Preview)\n  * Microsoft Fabric (Preview)\n  * Microsoft SharePoint (Preview)\n  * Model Context Protocol (MCP)\n  * OpenAPI\n  * Reminder Tool (Preview)\n  * Toolbox Search\n  * Web Search\n  * Work IQ (Preview)\n* **Get an OpenAI client** using `.get_openai_client()` method to run Responses, Conversations, Evaluations and Fine-Tuning operations with your Agent.\n* **Create and version toolboxes** that bundle collections of tools and skills for your agents, using `.toolboxes` operations
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 3725ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Locate the current official azure-ai-projects Python reference for connections and deployments.","parameters":{"query":"site:learn.microsoft.com python api azure-ai-projects AIProjectClient connections list get include_credentials deployments ModelDeployment"}}
```

**Output:**

```
{"results":[{"title":"Azure AI Projects client library for Python - version 2.5.0","content":"# Azure AI Projects client library for Python - version 2.5.0\n## Key concepts\n### Create and authenticate the client with Entra ID\nEntra ID is the only authentication method currently supported by the client.\nTo construct a synchronous client using a context manager:\n\u0060\u0060\u0060python\nimport os\nfrom azure.ai.projects import AIProjectClient\nfrom azure.identity import DefaultAzureCredential\n\nwith (\n    DefaultAzureCredential() as credential,\n    AIProjectClient(\n        endpoint=os.environ[\u0022FOUNDRY_PROJECT_ENDPOINT\u0022],\n        credential=credential\n    ) as project_client,\n):\n\u0060\u0060\u0060\nTo construct an asynchronous client, install the additional package [aiohttp](https://pypi.org/project/aiohttp/):\n\u0060\u0060\u0060bash\npip install aiohttp\n\u0060\u0060\u0060\nand run:\n\u0060\u0060\u0060python\nimport os\nimport asyncio\nfrom azure.ai.projects.aio import AIProjectClient\nfrom azure.identity.aio import DefaultAzureCredential\n\nasync with (\n    DefaultAzureCredential() as credential,\n    AIProjectClient(\n        endpoint=os.environ[\u0022FOUNDRY_PROJECT_ENDPOINT\u0022],\n        credential=credential\n    ) as project_client,\n):\n\u0060\u0060\u0060\n### Performing Responses operations using OpenAI client\nUse the \u0060.get_openai_client()\u0060 method to obtain an authenticated [OpenAI](https://github.com/openai/openai-python) client and run Responses, Conversations, Evaluations, Files, and Fine-Tuning operations. See the **responses**, **agents**, **evaluations**, **files**, and **finetuning** folders in the [samples](https://aka.ms/azsdk/azure-ai-projects-v2/python/samples/) for complete working examples.\nThe code below assumes the environment variable \u0060FOUNDRY_MODEL_NAME\u0060 is defined. It\u0027s the deployment name of an AI model in your Foundry Project. See first column titled \u0022Name\u0022 in the \u0022Deploye
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 3344ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official synchronous Python examples for listing/getting Foundry project connections and deployments.","parameters":{"language":"python","query":"azure-ai-projects AIProjectClient connections.list connections.get include_credentials deployments.list ModelDeployment"}}
```

**Output:**

```
{"results":[{"description":"description: Creates an AIProjectClient instance authenticated with DefaultAzureCredential, retrieves a Microsoft Fabric connection by name, initializes a FabricTool with the connection ID, and creates a new Azure AI agent with the Fabric tool enabled.\npackage: azure-ai-projects, azure-identity\nlanguage: python\n","codeSnippet":"project_client = AIProjectClient(\r\n    endpoint=os.environ[\u0022PROJECT_ENDPOINT\u0022],\r\n    credential=DefaultAzureCredential(),\r\n)\r\n\r\n# Look up the Fabric connection by name\r\nconn_id = project_client.connections.get(os.environ[\u0022FABRIC_CONNECTION_NAME\u0022]).id\r\n\r\n# Initialize the Fabric tool with the connection ID\r\nfabric = FabricTool(connection_id=conn_id)\r\n\r\nwith project_client:\r\n    agents_client = project_client.agents\r\n\r\n    agent = agents_client.create_agent(\r\n        model=os.environ[\u0022MODEL_DEPLOYMENT_NAME\u0022],\r\n        name=\u0022my-agent\u0022,\r\n        instructions=\u0022You are a helpful agent\u0022,\r\n        tools=fabric.definitions,\r\n    )\r\n    print(f\u0022Created agent, ID: {agent.id}\u0022)","link":"https://learn.microsoft.com/fabric/data-science/data-agent-foundry#add-fabric-data-agent-to-your-azure-ai-agent","language":"python"},{"description":"description: Imports necessary libraries and initializes the AIProjectClient with endpoint and default Azure credentials for authentication.\npackage: azure.ai.projects, azure.identity, azure.ai.agents.models\nlanguage: python\n","codeSnippet":"# Import necessary libraries\r\nimport os\r\nimport jsonref\r\nfrom azure.ai.projects import AIProjectClient\r\nfrom azure.identity import DefaultAzureCredential\r\n# import the following\r\nfrom azure.ai.agents.models import OpenApiTool, OpenApiAnonymousAuthDetails\r\n# use the following for connection auth\r\n# from azure.ai.agents.models import OpenApiTool, OpenApiConnectionAuthDetails, OpenApiConnectionSecurityScheme\r\n# use the following for managed i
... (truncated)
```

### ✅ 15. powershell — 4072ms

**Input:**

```json
{"command":"$metadata = Invoke-RestMethod 'https://pypi.org/pypi/azure-ai-projects/json'; [pscustomobject]@{ Version = $metadata.info.version; RequiresPython = $metadata.info.requires_python } | ConvertTo-Json -Compress","description":"Check latest azure-ai-projects version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
{"Version":"2.5.0","RequiresPython":">=3.10"}
<shellId: 0 completed with exit code 0>
```

### ✅ 16. view — 2975ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\api-reference.md","view_range":[1010,1070]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/api-reference.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/api-reference.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/api-reference.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/api-reference.md
@@ -1010,61 +1010,61 @@
 DatasetType            # FILE, FOLDER
 AzureAISearchQueryType # SIMPLE, FULL, SEMANTIC, VECTOR, VECTOR_SIMPLE_HYBRID, ...
 ```
 
 ### Evaluation Enums
 ```python
 EvaluatorType          # GROUNDEDNESS, RELEVANCE, COHERENCE, FLUENCY, etc.
 EvaluatorCategory      # QUALITY, SAFETY, CUSTOM
 EvaluatorDefinitionType # CODE_BASED, PROMPT_BASED
 EvaluatorMetricType    # NUMERIC, BOOLEAN, STRING
 EvaluatorMetricDirection # HIGHER_IS_BETTER, LOWER_IS_BETTER
 RiskCategory           # HATE_UNFAIRNESS, VIOLENCE, SELF_HARM, SEXUAL, etc.
 SampleType             # GOOD, BAD, NEUTRAL
 ```
 
 ### Memory Enums
 ```python
 MemoryItemKind         # CHAT_SUMMARY, USER_PROFILE
 MemoryStoreKind        # DEFAULT
 MemoryOperationKind    # ADD, REMOVE
 MemoryStoreUpdateStatus # COMPLETED, PENDING, FAILED
 ```
 
 ### Schedule & Trigger Enums
 ```python
 TriggerType            # CRON, RECURRENCE, ONE_TIME
 RecurrenceType         # HOURLY, DAILY, WEEKLY, MONTHLY
 DayOfWeek              # MONDAY through SUNDAY
 ScheduleTaskType       # EVALUATION, INSIGHT
 ScheduleProvisioningStatus # PROVISIONING, PROVISIONED, FAILED
 ```
 
 ### Credential Enums
 ```python
 CredentialType         # API_KEY, ENTRA_ID, SAS, NO_AUTH, AGENTIC_IDENTITY, CUSTOM
 ```
 
 ### Connection Enums
 ```python
 ConnectionType         # AZURE_OPEN_AI, AZURE_AI_S
... (truncated)
```

### ✅ 17. rg — 1599ms

**Input:**

```json
{"-i":true,"-n":true,"C":2,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt","pattern":"connections\\.list|connections\\.get|include_credentials|deployments\\.list|deployments\\.get|ModelDeployment|url|title"}
```

**Output:**

```
Output too large to read at once (20.9 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787798377208-3.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787798370026-copilot-tool-output-324003ad0a414ab7959722aaf63ee741.txt:1:{"results":[{"description":"description: Creates an AIProjectClient instance authenticated with DefaultAzureCredential, retrieves a Microsoft Fabric connection by name, initializes a FabricTool with the connection ID, and creates a new Azure AI agent with the Fabric tool enabled.\npackage: azure-ai-projects, azure-identity\nlanguage: python\n","codeSnippet":"project_client = AIProjectClien
```

### ✅ 18. rg — 1224ms

**Input:**

```json
{"-i":true,"-n":true,"C":2,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787798370026-copilot-tool-output-324003ad0a414ab7959722aaf63ee741.txt","pattern":"connections\\.list|connections\\.get|include_credentials|deployments\\.list|deployments\\.get|ModelDeployment|url|title"}
```

**Output:**

```
Output too large to read at once (40.4 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787798377168-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt:1:{"results":[{"title":"Azure AI Projects client library for Python - version 2.5.0","content":"# Azure AI Projects client library for Python - version 2.5.0\n## Key concepts\n### Create and authenticate the client with Entra ID\nEntra ID is the only authentication method currently supported by the client.\nTo construct a synchronous client using a context manager:\n\u0060\u0060\u0060python\
```

### ✅ 19. powershell — 1318ms

**Input:**

```json
{"command":"$docs = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt' -Raw | ConvertFrom-Json; $docs.results | Where-Object { $_.content -match 'connections|deployments|ModelDeployment' } | Select-Object title,content,contentUrl | ConvertTo-Json -Depth 4","description":"Extract relevant official documentation","initial_wait":30,"mode":"sync"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … f3de197e7cca44a0947f43655a83e494.txt' -Raw | ConvertFrom-Json; $docs. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 20. view — 820ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt","view_range":[1,10]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt b/C:/Users/weidxu/AppData/Local/Temp/1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787798370459-copilot-tool-output-f3de197e7cca44a0947f43655a83e494.txt
@@ -1,10 +1,10 @@
 {"results":[{"title":"Azure AI Projects client library for Python - version 2.5.0","content":"# Azure AI Projects client library for Python - version 2.5.0\n## Key concepts\n### Create and authenticate the client with Entra ID\nEntra ID is the only authentication method currently supported by the client.\nTo construct a synchronous client using a context manager:\n\u0060\u0060\u0060python\nimport os\nfrom azure.ai.projects import AIProjectClient\nfrom azure.identity import DefaultAzureCredential\n\nwith (\n    DefaultAzureCredential() as credential,\n    AIProjectClient(\n        endpoint=os.environ[\u0022FOUNDRY_PROJECT_ENDPOINT\u0022],\n        credential=credential\n    ) as project_client,\n):\n\u0060\u0060\u0060\nTo construct an asynchronous client, install the additional package [aiohttp](https://pypi.org/project/aiohttp/):\n\u0060\u0060\u0060bash\npip install aiohttp\n\u0060\u0060\u0060\nand run:\n\u0060\u0060\u0060python\nimport os\nimport asyncio\nfrom azure.ai.projects.aio import AIProjectClient\nfrom azure.identity.aio import DefaultAzureCredential\n\nasync with (\n    DefaultAzureCredential() as credential,\n    AIProjectClient(\n        endpoint=os.environ[\u0022FOUNDRY_PROJECT_ENDPOINT\u0022],\n        credential=credential\n    ) as project_client,\n):\n\u0060\u0060\u0060\n### Performing Responses operations using OpenAI client\nUse the \u0060.get_openai_client()\u0060 method to obtain an authenticated [OpenAI](https://github.com/openai/openai-python) client a
... (truncated)
```

### ✅ 21. azure-documentation (via azure) — 3827ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the current synchronous ConnectionsOperations API.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-ai-projects/azure.ai.projects.operations.connectionsoperations?view=azure-python"}}
```

**Output:**

```
# DeploymentsOperations Class

Warning

DO NOT instantiate this class directly.

Instead, you should access the following operations through

[AIProjectClient](azure.ai.projects.aiprojectclient?view=azure-python)'s

&lt;xref:deployments&gt; attribute.

## Constructor

```python
DeploymentsOperations(*args, **kwargs)
```

## Methods

| get | Get a deployment.<br><br>Retrieves a deployed model. |
| --- | --- |
| list | List deployments.<br><br>Returns the deployed models available in the current project, optionally filtered by publisher, model name, or deployment type. |

### get

Get a deployment.

Retrieves a deployed model.

```python
get(name: str, **kwargs: Any) -> Deployment
```

#### Parameters

| Name | Description |
| --- | --- |
| name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Name of the deployment. Required. |

#### Returns

| Type | Description |
| --- | --- |
| [Deployment](azure.ai.projects.models.deployment?view=azure-python) | Deployment. The Deployment is compatible with MutableMapping |

#### Exceptions

| Type | Description |
| --- | --- |
| [HttpResponseError](../azure-core/azure.core.exceptions.httpresponseerror?view=azure-python) |  |

### list

List deployments.

Returns the deployed models available in the current project, optionally filtered by publisher, model name, or deployment type.

```python
list(*, model_publisher: str | None = None, model_name: str | None = None, deployment_type: str | DeploymentType | None = None, **kwargs: Any) -> ItemPaged[Deployment]
```

#### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| model\_publisher | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Model publisher to filter models by. Default value is None.<br><br>Default value: None |
| model\_name | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Model name (the publisher specific name) to filter
... (truncated)
```

### ✅ 22. azure-documentation (via azure) — 3450ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the current synchronous DeploymentsOperations API.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-ai-projects/azure.ai.projects.operations.deploymentsoperations?view=azure-python"}}
```

**Output:**

```
# ModelDeployment Class

Model Deployment Definition.

## Constructor

```python
ModelDeployment(*args: Any, **kwargs: Any)
```

## Variables

| Name | Description |
| --- | --- |
| name | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Name of the deployment. Required. |
| type | [str](https://docs.python.org/3/library/stdtypes.html#str) or &lt;xref:azure.ai.projects.models.MODEL\_DEPLOYMENT&gt;<br><br>The type of the deployment. Required. Model deployment. |
| model\_name | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Publisher-specific name of the deployed model. Required. |
| model\_version | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Publisher-specific version of the deployed model. Required. |
| model\_publisher | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Name of the deployed model's publisher. Required. |
| capabilities | [dict](https://docs.python.org/3/library/stdtypes.html#dict)[[str](https://docs.python.org/3/library/stdtypes.html#str), [str](https://docs.python.org/3/library/stdtypes.html#str)] <br><br>Capabilities of deployed model. Required. |
| sku | [ModelDeploymentSku](azure.ai.projects.models.modeldeploymentsku?view=azure-python)<br><br>Sku of the model deployment. Required. |
| connection\_name | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Name of the connection the deployment comes from. |

## Methods

| as_dict | Return a dict that can be turned into json using json.dump. |
| --- | --- |
| clear | Remove all items from the dictionary. |
| copy |  |
| get | Get the value for key if key is in the dictionary, else default. :param str key: The key to look up. :param any default: The value to return if key is not in the dictionary. Defaults to None :returns: The value for key if key is in the dictionary, else default. :rtype: any |
| items |  |
| keys |  |
| pop | Removes specified key and return the correspondin
... (truncated)
```

### ✅ 23. azure-documentation (via azure) — 3106ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read current Connection model fields.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-ai-projects/azure.ai.projects.models.connection?view=azure-python"}}
```

**Output:**

```
# ConnectionsOperations Class

Warning

DO NOT instantiate this class directly.

Instead, you should access the following operations through

[AIProjectClient](azure.ai.projects.aiprojectclient?view=azure-python)'s

&lt;xref:connections&gt; attribute.

## Constructor

```python
ConnectionsOperations(*args, **kwargs)
```

## Methods

| get | Get a connection by name. |
| --- | --- |
| get_default | Get the default connection for a given connection type. |
| list | List connections.<br><br>Returns the connections available in the current project, optionally filtered by type or default status. |

### get

Get a connection by name.

```python
get(name: str, *, include_credentials: bool | None = False, **kwargs: Any) -> Connection
```

#### Parameters

| Name | Description |
| --- | --- |
| name<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The name of the connection. Required. |

#### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| include\_credentials | [bool](https://docs.python.org/3/library/functions.html#bool)<br><br>Whether to include credentials in the response. Default is False.<br><br>Default value: False |

#### Returns

| Type | Description |
| --- | --- |
| [Connection](azure.ai.projects.models.connection?view=azure-python) | Connection. The Connection is compatible with MutableMapping |

#### Exceptions

| Type | Description |
| --- | --- |
| [HttpResponseError](../azure-core/azure.core.exceptions.httpresponseerror?view=azure-python) |  |

### get\_default

Get the default connection for a given connection type.

```python
get_default(connection_type: str | ConnectionType, *, include_credentials: bool | None = False, **kwargs: Any) -> Connection
```

#### Parameters

| Name | Description |
| --- | --- |
| connection\_type<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str) or [ConnectionType](azure.ai.projects.m
... (truncated)
```

### ✅ 24. azure-documentation (via azure) — 2757ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read current ModelDeployment model fields.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-ai-projects/azure.ai.projects.models.modeldeployment?view=azure-python"}}
```

**Output:**

```
# Connection Class

Response from the list and get connections operations.

## Constructor

```python
Connection(*args: Any, **kwargs: Any)
```

## Variables

| Name | Description |
| --- | --- |
| name | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The friendly name of the connection, provided by the user. Required. |
| id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>A unique identifier for the connection, generated by the service. Required. |
| type | [str](https://docs.python.org/3/library/stdtypes.html#str) or [ConnectionType](azure.ai.projects.models.connectiontype?view=azure-python)<br><br>Category of the connection. Required. Known values are: "AzureOpenAI", "AzureBlob", "AzureStorageAccount", "CognitiveSearch", "CosmosDB", "ApiKey", "AppConfig", "AppInsights", "CustomKeys", and "RemoteTool\_Preview". |
| target | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The connection URL to be used for this service. Required. |
| is\_default | [bool](https://docs.python.org/3/library/functions.html#bool)<br><br>Whether the connection is tagged as the default connection of its type. Required. |
| credentials | [BaseCredentials](azure.ai.projects.models.basecredentials?view=azure-python)<br><br>The credentials used by the connection. Required. |
| metadata | [dict](https://docs.python.org/3/library/stdtypes.html#dict)[[str](https://docs.python.org/3/library/stdtypes.html#str), [str](https://docs.python.org/3/library/stdtypes.html#str)] <br><br>Metadata of the connection. Required. |

## Methods

| as_dict | Return a dict that can be turned into json using json.dump. |
| --- | --- |
| clear | Remove all items from the dictionary. |
| copy |  |
| get | Get the value for key if key is in the dictionary, else default. :param str key: The key to look up. :param any default: The value to return if key is not in the dictionary. Defaults to None :returns: The value for key if key is in th
... (truncated)
```

### ✅ 25. apply_patch — 750ms

**Input:**

```json
"*** Begin Patch\n*** Add File: app.py\n+import os\n+import sys\n+\n+from azure.ai.projects import AIProjectClient\n+from azure.ai.projects.models import Connection, ModelDeployment\n+from azure.core.exceptions import AzureError\n+from azure.identity import DefaultAzureCredential\n+\n+\n+def required_environment_variable(name: str) -\u003e str:\n+    value = os.environ.get(name, \"\").strip()\n+    if not value:\n+        raise ValueError(f\"Environment variable {name} must be set.\")\n+    return value\n+\n+\n+def display_value(value: object) -\u003e str:\n+    return str(getattr(value, \"value\", value))\n+\n+\n+def print_connection(connection: Connection) -\u003e None:\n+    print(f\"  Name: {connection.name}\")\n+    print(f\"  Type: {display_value(connection.type)}\")\n+    print(f\"  Target: {connection.target}\")\n+    print(f\"  Default: {connection.is_default}\")\n+\n+\n+def print_model_deployment(deployment: ModelDeployment) -\u003e None:\n+    print(f\"  Name: {deployment.name}\")\n+    print(f\"  Model publisher: {deployment.model_publisher}\")\n+    print(f\"  Model name: {deployment.model_name}\")\n+    print(f\"  Model version: {deployment.model_version}\")\n+\n+\n+def inspect_project(\n+    project_client: AIProjectClient,\n+    connection_name: str,\n+    deployment_name: str,\n+) -\u003e None:\n+    print(\"Project connections\")\n+    connection_count = 0\n+    for connection in project_client.connections.list():\n+        connection_count += 1\n+        print_connection(connection)\n+        print()\n+    if connection_count == 0:\n+        print(\"  No connections found.\")\n+        print()\n+\n+    print(f\"Requested connection: {connection_name}\")\n+    connection = project_client.connections.get(\n+        connection_name,\n+        include_credentials=False,\n+    )\n+    print_connection(connection)\n+    print()\n+\n+    print(\"Project model deployments\")\n+    model_deployment_count = 0\n+    for deployment in project_client.deployments.list():\n+        if isinstance(deployment, ModelDeployment):\n+            model_deployment_count += 1\n+            print_model_deployment(deployment)\n+            print()\n+    if model_deployment_count == 0:\n+        print(\"  No model deployments found.\")\n+        print()\n+\n+    print(f\"Requested deployment: {deployment_name}\")\n+    deployment = project_client.deployments.get(deployment_name)\n+    if not isinstance(deployment, ModelDeployment):\n+        raise TypeError(\n+            f\"Deployment {deployment_name!r} is not a model deployment \"\n+            f\"(received {type(deployment).__name__}).\"\n+        )\n+    print_model_deployment(deployment)\n+\n+\n+def main() -\u003e int:\n+    try:\n+        endpoint = required_environment_variable(\"FOUNDRY_PROJECT_ENDPOINT\")\n+        connection_name = required_environment_variable(\"CONNECTION_NAME\")\n+        deployment_name = required_environment_variable(\"DEPLOYMENT_NAME\")\n+\n+        with (\n+            DefaultAzureCredential() as credential,\n+            AIProjectClient(endpoint=endpoint, credential=credential) as project_client,\n+        ):\n+            inspect_project(project_client, connection_name, deployment_name)\n+    except (AzureError, TypeError, ValueError) as error:\n+        print(f\"Error: {error}\", file=sys.stderr)\n+        return 1\n+\n+    return 0\n+\n+\n+if __name__ == \"__main__\":\n+    raise SystemExit(main())\n*** Add File: requirements.txt\n+azure-ai-projects==2.5.0\n+azure-identity\u003e=1.25.0,\u003c2.0.0\n*** Add File: README.md\n+# Microsoft Foundry project inventory\n+\n+This synchronous Python console application lists a Microsoft Foundry project's\n+connections and model deployments, then retrieves one named connection and one\n+named model deployment. Connection credentials are never requested or printed.\n+\n+## Prerequisites\n+\n+- Python 3.10 or later\n+- A Microsoft Entra identity with permission to read the Foundry project\n+- Local authentication available to `DefaultAzureCredential`, such as an Azure\n+  CLI or Visual Studio Code sign-in\n+\n+## Install and run (PowerShell)\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\python -m pip install -r requirements.txt\n+\n+$env:FOUNDRY_PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:CONNECTION_NAME = \"\u003cconnection-name\u003e\"\n+$env:DEPLOYMENT_NAME = \"\u003cdeployment-name\u003e\"\n+.\\.venv\\Scripts\\python app.py\n+```\n+\n+For production-hosted execution, configure managed identity or workload identity\n+and set `AZURE_TOKEN_CREDENTIALS=prod` to constrain `DefaultAzureCredential` to\n+production-safe credential types.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-with-azure-tools-642855756/app.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-with-azure-tools-642855756/app.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-with-azure-tools-642855756/app.py
@@ -1,0 +1,99 @@
+import os
+import sys
+
+from azure.ai.projects import AIProjectClient
+from azure.ai.projects.models import Connection, ModelDeployment
+from azure.core.exceptions import AzureError
+from azure.identity import DefaultAzureCredential
+
+
+def required_environment_variable(name: str) -> str:
+    value = os.environ.get(name, "").strip()
+    if not value:
+        raise ValueError(f"Environment variable {name} must be set.")
+    return value
+
+
+def display_value(value: object) -> str:
+    return str(getattr(value, "value", value))
+
+
+def print_connection(connection: Connection) -> None:
+    print(f"  Name: {connection.name}")
+    print(f"  Type: {display_value(connection.type)}")
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
+def inspect_project(
+    project_client: AIProjectClient,
+    connection_name: str,
+    deployment_name: str,
+) -> None:
+    print("Project connections")
+    connection_count = 0
+    for connection in project_client.connections.list():
+        connection_count += 1
+        print_connection(connection)
+        print()
+    if connection_count == 0:
+        print("  No connec
... (truncated)
```

### ✅ 26. powershell — 148815ms

**Input:**

```json
{"command":"python -m venv .venv \u0026\u0026 .\\.venv\\Scripts\\python -m pip install --quiet -r requirements.txt \u0026\u0026 .\\.venv\\Scripts\\python -m compileall -q app.py \u0026\u0026 .\\.venv\\Scripts\\python -c \"import inspect; from azure.ai.projects.operations import ConnectionsOperations, DeploymentsOperations; from azure.ai.projects.models import Connection, ModelDeployment; import app; assert inspect.signature(ConnectionsOperations.get).parameters['include_credentials'].default is False; assert 'name' in inspect.signature(DeploymentsOperations.get).parameters; assert all(hasattr(Connection, field) for field in ('name', 'type', 'target', 'is_default')); assert all(hasattr(ModelDeployment, field) for field in ('name', 'model_publisher', 'model_name', 'model_version')); print('SDK contract and imports verified')\"; .\\.venv\\Scripts\\python app.py; if ($LASTEXITCODE -ne 1) { throw \"Expected missing-environment run to exit with code 1.\" }","description":"Install dependencies and validate application","initial_wait":180,"mode":"sync"}
```

**Output:**

```
[notice] A new release of pip is available: 25.2 -> 26.2.1
[notice] To update, run: C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-with-azure-tools-642855756\.venv\Scripts\python.exe -m pip install --upgrade pip
Traceback (most recent call last):
  File "<string>", line 1, in <module>
    import inspect; from azure.ai.projects.operations import ConnectionsOperations, DeploymentsOperations; from azure.ai.projects.models import Connection, ModelDeployment; import app; assert inspect.signature(ConnectionsOperations.get).parameters['include_credentials'].default is False; assert 'name' in inspect.signature(DeploymentsOperations.get).parameters; assert all(hasattr(Connection, field) for field in ('name', 'type', 'target', 'is_default')); assert all(hasattr(ModelDeployment, field) for field in ('name', 'model_publisher', 'model_name', 'model_version')); print('SDK contract and imports verified')
                                                                                                                                                                                                                                                                                                                                                                                  ~~~^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
AssertionError
Error: Environment variable FOUNDRY_PROJECT_ENDPOINT must be set.
<shellId: 2 completed with exit code 1>
```

### ✅ 27. powershell — 3285ms

**Input:**

```json
{"command":"@'\nimport inspect\nfrom azure.ai.projects.operations import ConnectionsOperations, DeploymentsOperations\nfrom azure.ai.projects.models import Connection, ModelDeployment, NoAuthenticationCredentials, ModelDeploymentSku\n\nprint(inspect.signature(ConnectionsOperations.get))\nprint(inspect.signature(DeploymentsOperations.get))\nprint(inspect.signature(Connection))\nprint(inspect.signature(ModelDeployment))\nprint(inspect.signature(NoAuthenticationCredentials))\nprint(inspect.signature(ModelDeploymentSku))\n'@ | .\\.venv\\Scripts\\python -","description":"Inspect installed SDK model contracts","initial_wait":30,"mode":"sync"}
```

**Output:**

```
(self, name: str, *, include_credentials: Optional[bool] = False, **kwargs: Any) -> azure.ai.projects.models._models.Connection
(self, name: str, **kwargs: Any) -> azure.ai.projects.models._models.Deployment
(*args: Any, **kwargs: Any) -> Self
(*args: Any, **kwargs: Any) -> None
(*args: Any, **kwargs: Any) -> None
(*args: Any, **kwargs: Any) -> None
<shellId: 3 completed with exit code 0>
```

### ✅ 28. powershell — 5126ms

**Input:**

```json
{"command":"@'\nfrom contextlib import redirect_stdout\nfrom io import StringIO\n\nfrom azure.ai.projects.models import (\n    Connection,\n    ModelDeployment,\n    ModelDeploymentSku,\n    NoAuthenticationCredentials,\n)\n\nfrom app import inspect_project\n\nconnection = Connection(\n    name=\"sample-connection\",\n    id=\"connection-id\",\n    type=\"AzureOpenAI\",\n    target=\"https://example.invalid\",\n    is_default=True,\n    credentials=NoAuthenticationCredentials(type=\"NoAuthentication\"),\n    metadata={},\n)\ndeployment = ModelDeployment(\n    name=\"sample-deployment\",\n    type=\"ModelDeployment\",\n    model_name=\"sample-model\",\n    model_version=\"1\",\n    model_publisher=\"sample-publisher\",\n    capabilities={},\n    sku=ModelDeploymentSku(name=\"GlobalStandard\", capacity=1),\n)\n\nclass Connections:\n    def __init__(self):\n        self.get_arguments = None\n\n    def list(self):\n        yield connection\n\n    def get(self, name, *, include_credentials=False):\n        self.get_arguments = (name, include_credentials)\n        return connection\n\nclass Deployments:\n    def list(self):\n        yield deployment\n\n    def get(self, name):\n        assert name == \"sample-deployment\"\n        return deployment\n\nclass ProjectClient:\n    def __init__(self):\n        self.connections = Connections()\n        self.deployments = Deployments()\n\nclient = ProjectClient()\noutput = StringIO()\nwith redirect_stdout(output):\n    inspect_project(client, \"sample-connection\", \"sample-deployment\")\n\nrendered = output.getvalue()\nassert client.connections.get_arguments == (\"sample-connection\", False)\nfor expected in (\n    \"Name: sample-connection\",\n    \"Type: AzureOpenAI\",\n    \"Target: https://example.invalid\",\n    \"Default: True\",\n    \"Name: sample-deployment\",\n    \"Model publisher: sample-publisher\",\n    \"Model name: sample-model\",\n    \"Model version: 1\",\n):\n    assert expected in rendered\nprint(\"Offline application behavior verified\")\n'@ | .\\.venv\\Scripts\\python -; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; .\\.venv\\Scripts\\python -m compileall -q app.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $missingOutput = .\\.venv\\Scripts\\python app.py 2\u003e\u00261; $missingExit = $LASTEXITCODE; if ($missingExit -ne 1 -or $missingOutput -notmatch 'FOUNDRY_PROJECT_ENDPOINT') { throw 'Missing configuration handling failed.' }; Write-Output 'Configuration error handling verified'","description":"Verify application offline behavior","initial_wait":120,"mode":"sync"}
```

**Output:**

```
Offline application behavior verified
Configuration error handling verified
<shellId: 4 completed with exit code 0>
```

### ❌ 29. glob — 20386ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-python-project-resource-inventory-python-azure-tools-with-azure-tools-642855756","pattern":"*"}
```

**Error:**

```
timeout
```

### ✅ 30. powershell — 4707ms

**Input:**

```json
{"command":"$targets = @('.venv', '__pycache__'); foreach ($target in $targets) { if (Test-Path -LiteralPath $target) { Remove-Item -LiteralPath $target -Recurse -Force } }; Get-ChildItem -File | Select-Object -ExpandProperty Name","description":"Clean validation artifacts and list deliverables","initial_wait":120,"mode":"sync"}
```

**Output:**

```
app.py
README.md
requirements.txt
<shellId: 5 completed with exit code 0>
```

## Generated Files

- `README.md`
- `app.py`
- `requirements.txt`

## Copilot Response

I’m grounding the implementation in the installed SDK reference, then I’ll create and locally validate the console app without contacting Azure.

The current SDK contract is confirmed (`azure-ai-projects` 2.5.0): connection retrieval supports `include_credentials=False`, and deployment results can be narrowed safely with `ModelDeployment`. I’m creating the app and its run instructions now.

The application imports and fails cleanly when configuration is absent. One SDK introspection assertion was too strict for generated model descriptors, so I’m replacing it with instance-level contract checks and an offline fake-client execution.

Created a complete synchronous Microsoft Foundry inventory application:

- `app.py` — lists and retrieves connections without credentials, then lists and validates model deployments.
- `requirements.txt` — pins `azure-ai-projects==2.5.0` with Azure Identity.
- `README.md` — concise PowerShell installation, configuration, and run commands.

The application uses `DefaultAzureCredential`, SDK context managers, pageable APIs, typed `Connection`/`ModelDeployment` metadata, required environment-variable validation, and clear error reporting.

## Reviewer Notes (LLM-as-Judge)

### Summary

The exception-handling criterion is fully met.

### Strengths

- Catches the Azure SDK exception hierarchy through AzureError.
- Reports failures clearly and returns exit code 1.

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
hyoka run --prompt-id ai-projects-dp-python-project-resource-inventory --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
