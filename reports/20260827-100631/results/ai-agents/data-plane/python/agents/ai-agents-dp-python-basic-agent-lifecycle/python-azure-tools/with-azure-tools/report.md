# Evaluation Report: ai-agents-dp-python-basic-agent-lifecycle

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 342.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-agents-dp-python-basic-agent-lifecycle` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 15/17 |
| Duration | 342.9s |
| Timestamp | 2026-08-27T02:21:55Z |
| Files Generated | 9 |
| Event Count | 3573 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 241.4s |
| Review | 100.5s |
| **Total** | **342.9s** |

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
| Token Usage | in=929971 out=6271 |
| Turn Count | 17 |

## Tool Calls

### ✅ 1. skill — 3471ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 2878ms

**Input:**

```json
{"command":"get_ai_application_best_practices","intent":"Create a local-only synchronous Python console app using azure-ai-agents, DefaultAzureCredential, an agent/thread/run lifecycle, terminal-status polling, message output, and deterministic cleanup."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-get_azure_bestpractices (via azure) — 2428ms

**Input:**

```json
{"command":"get_azure_bestpractices_ai_app","intent":"Get guidance for a local-only synchronous Python azure-ai-agents console application.","parameters":{}}
```

**Output:**

```
{"status":200,"message":"","results":["# Best Practices for Building AI Apps with Azure AI Services\r\n\r\nThis guide covers best practices for:\r\n- Generating code for AI agent applications using Microsoft Foundry and any other Azure AI models\r\n- Building agents with Microsoft Foundry\r\n- Implementing chatbots, workflows, assistants, and multi-agent systems with Microsoft Agent Framework\r\n\r\n> Note: Understanding User Intent of Microsoft Foundry\r\n> \r\n> **Two Types of Operations:**\r\n> \r\n> **1. Resource Management** - Use `foundry` tool\r\n> - Keywords: \"create/deploy/provision agent/embedding IN Foundry\", \"set up agent/embedding resource\"\r\n> - Creates Microsoft Foundry agent resource via commands under `foundry`\r\n> \r\n> **2. Application Development** - Use this best practices tool\r\n> - Keywords: \"build/write/implement agent\", \"develop chatbot/assistant\", \"agent code\"\r\n> - Generates code using Microsoft Agent Framework SDK\r\n> - Returns code samples and implementation patterns\r\n> \r\n> **When unsure, ask:** \"Do you want to create an agent resource in Microsoft Foundry, or generate agent application code?\"\r\n\r\n## Key Principles\r\n**Defaults and Recommendations:**\r\n- **Default language**: If the user does not specify a programming language, use **.NET** as the default programming language.\r\n- **Recommended SDK**: Use **Microsoft Agent Framework** for building AI agents, chatbots, assistants, and multi-agent systems - it provides flexible orchestration, multi-agent patterns, and cross-platform support (.NET and Python)\r\n- **Select model first**: Model selection is REQUIRED for ANY AI implementation - always select the appropriate AI model before coding if the user does not specify a model\r\n\r\n**Process:**\r\n1. **Gather information first**: Always search / fetch / query for sufficient knowledge (quickstarts, tutorials, feature guides, code samples) BEFORE thinking or writing any code - this ensures high-quality code ge
... (truncated)
```

### ✅ 4. glob — 7757ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-ai-projects-py\references\agents.md
```

### ✅ 5. glob — 7416ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py","pattern":"**/{acceptance-criteria.md,examples.md,agents.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-documentation (via azure) — 8120ms

**Input:**

```json
{"command":"search","intent":"Find current official Python azure-ai-agents documentation for AgentsClient creation with endpoint and DefaultAzureCredential, creating an agent/thread/message/run, polling a run to terminal status, listing messages in chronological order, inspecting assistant text content, and deleting thread and agent.","parameters":{"query":"azure-ai-agents Python AgentsClient create_agent threads messages runs get_run RunStatus terminal list order assistant text delete_thread delete_agent"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. view — 2269ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\agents.md","view_range":[1,220]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md
@@ -1,220 +1,220 @@
 # Agent Operations Reference
 
 ## Agent Types and Kinds
 
 ```python
 from azure.ai.projects.models import AgentKind
 
 # Agent kinds
 # - "prompt": Standard prompt-based agents
 # - "hosted": Hosted agents
 # - "container_app": Container App agents
 # - "workflow": Workflow agents
 
 # Filter agents by kind
 agents = project_client.agents.list(kind=AgentKind.PROMPT)
 ```
 
 ## Basic Agent Creation
 
 ```python
 agent = project_client.agents.create_agent(
     model=os.environ["AZURE_AI_MODEL_DEPLOYMENT_NAME"],
     name="my-agent",
     instructions="You are a helpful assistant.",
 )
 print(f"Created agent, ID: {agent.id}")
 
 # Clean up when done
 project_client.agents.delete_agent(agent.id)
 ```
 
 ## Versioned Agents with PromptAgentDefinition
 
 For production deployments, use versioned agents:
 
 ```python
 from azure.ai.projects.models import PromptAgentDefinition
 
 agent = project_client.agents.create_version(
     agent_name="customer-support-agent",
     definition=PromptAgentDefinition(
         model=os.environ["AZURE_AI_MODEL_DEPLOYMENT_NAME"],
         instructions="You are a customer support specialist.",
         tools=[],  # Add tools as needed
     ),
     version_label="v1.0",
     description="Initial version",
 )
 print(f"Agent created (id: {agent.id}, name: {agent.name}, version: {agent.version})")
 ```
 
 #
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3024ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current low-level Azure AI Agents Python SDK lifecycle docs.","parameters":{"query":"azure-ai-agents Python AgentsClient create agent thread message run get_run RunStatus list messages order delete thread agent"}}
```

**Output:**

```
{"results":[{"title":"AzureAIAgentThread Class-Methods","content":"# AzureAIAgentThread Class\r\n\r\nAn Azure AI Agent Thread that manages creation, deletion, and retrieval of messages for an Azure-hosted agent conversation. This thread orchestrates calls to the Azure AI Project client API to perform operations like creating a new thread, adding messages, or retrieving conversation history.\r\n\r\n## Methods\r\n\r\n### get\\_messages\r\n\r\nRetrieve messages from the thread in either ascending or descending order.\r\n\r\n\u0060\u0060\u0060python\r\nasync get_messages(sort_order: Literal[\u0022asc\u0022,\u0022desc\u0022] = \u0022desc\u0022) -\u003E AsyncIterable[ChatMessageContent]\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r\n| sort\\_order | The sort order for the messages, defaults to descending.\u003Cbr\u003E\u003Cbr\u003EDefault value: desc |\r\n\r\n#### Returns\r\n\r\n| Type | Description |\r\n| --- | --- |\r\n|  | An async stream of ChatMessageContent for this thread\u0027s conversation. |","contentUrl":"https://learn.microsoft.com/python/api/semantic-kernel/semantic_kernel.agents.azure_ai.azure_ai_agent.azureaiagentthread?view=semantic-kernel-python"},{"title":"AzureAIAgentThread Class-Methods","content":"# AzureAIAgentThread Class\r\n\r\nAn Azure AI Agent Thread that manages creation, deletion, and retrieval of messages for an Azure-hosted agent conversation. This thread orchestrates calls to the Azure AI Project client API to perform operations like creating a new thread, adding messages, or retrieving conversation history.\r\n\r\n## Methods\r\n\r\n### \\_on\\_new\\_message\r\n\r\nInvoked on new user or system messages. Creates the message in the Azure AI thread if it doesn\u0027t match the existing thread ID.\r\n\r\n\u0060\u0060\u0060python\r\nasync _on_new_message(new_message: str | ChatMessageContent) -\u003E None\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 2646ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official Python code samples for a synchronous azure-ai-agents basic conversation lifecycle.","parameters":{"language":"python","query":"azure-ai-agents AgentsClient synchronous create_agent create_thread create_message create_run get_run list_messages delete_thread delete_agent"}}
```

**Output:**

```
{"results":[{"description":"description: Authenticates the user to their Azure subscription interactively using the Azure CLI.\nlanguage: python\n","codeSnippet":"import os\r\nfrom pathlib import Path\r\nfrom azure.ai.projects import AIProjectClient\r\nfrom azure.identity import DefaultAzureCredential\r\nfrom azure.ai.agents.models import CodeInterpreterTool\r\n\r\n# Create an AIProjectClient instance\r\nproject_client = AIProjectClient(\r\n    endpoint=os.getenv(\u0022PROJECT_ENDPOINT\u0022),\r\n    credential=DefaultAzureCredential(),  \r\n    # Use Azure Default Credential for authentication\r\n)\r\n\r\nwith project_client:\r\n\r\n    code_interpreter = CodeInterpreterTool()\r\n\r\n    agent = project_client.agents.create_agent(\r\n        model=os.getenv(\u0022MODEL_DEPLOYMENT_NAME\u0022),  # Model deployment name\r\n        name=\u0022my-agent\u0022,  # Name of the agent\r\n        instructions=\u0022\u0022\u0022You politely help with math questions. \r\n        Use the Code Interpreter tool when asked to visualize numbers.\u0022\u0022\u0022,  \r\n        # Instructions for the agent\r\n        tools=code_interpreter.definitions,  # Attach the tool\r\n        tool_resources=code_interpreter.resources,  # Attach tool resources\r\n    )\r\n    print(f\u0022Created agent, ID: {agent.id}\u0022)\r\n\r\n    # Create a thread for communication\r\n    thread = project_client.agents.threads.create()\r\n    print(f\u0022Created thread, ID: {thread.id}\u0022)\r\n\r\n    question = \u0022\u0022\u0022Draw a graph for a line with a slope of 4 \r\n    and y-intercept of 9 and provide the file to me?\u0022\u0022\u0022\r\n\r\n    # Add a message to the thread\r\n    message = project_client.agents.messages.create(\r\n        thread_id=thread.id,\r\n        role=\u0022user\u0022,  # Role of the message sender\r\n        content=question,  # Message content\r\n    )\r\n    print(f\u0022Created message, ID: {message[\u0027id\u0027]}\u0022)\r\n\r\n    # Create and process an agent 
... (truncated)
```

### ✅ 10. rg — 2895ms

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787797413191-copilot-tool-output-b94079b60da44f9ead788164430a0ad0.txt","pattern":"AgentsClient|create_run|get_run|RunStatus|ListSortOrder|delete_thread|delete_agent"}
```

**Output:**

```
Output too large to read at once (25.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787797421046-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787797413256-copilot-tool-output-26e09299af054ec6be2191e2018f928b.txt:{"results":[{"title":"AzureAIAgentThread Class-Methods","content":"# AzureAIAgentThread Class\r\n\r\nAn Azure AI Agent Thread that manages creation, deletion, and retrieval of messages for an Azure-hosted agent conversation. This thread orchestrates calls to the Azure AI Project client API to perform operations like creating a new thread, adding messages, or retrieving conversation history.\
```

### ✅ 11. rg — 2443ms

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787797413256-copilot-tool-output-26e09299af054ec6be2191e2018f928b.txt","pattern":"AgentsClient|create_run|get_run|RunStatus|ListSortOrder|delete_thread|delete_agent"}
```

**Output:**

```
Output too large to read at once (30.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787797421044-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787797413191-copilot-tool-output-b94079b60da44f9ead788164430a0ad0.txt:{"results":[{"description":"description: Authenticates the user to their Azure subscription interactively using the Azure CLI.\nlanguage: python\n","codeSnippet":"import os\r\nfrom pathlib import Path\r\nfrom azure.ai.projects import AIProjectClient\r\nfrom azure.identity import DefaultAzureCredential\r\nfrom azure.ai.agents.models import CodeInterpreterTool\r\n\r\n# Create an AIProjectClien
```

### ✅ 12. view — 2011ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-ai-projects-py\\references\\agents.md","view_range":[220,360]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-ai-projects-py/references/agents.md
@@ -220,141 +220,141 @@
 class MyHandler(AgentEventHandler):
     def on_message_delta(self, delta):
         if delta.text:
             print(delta.text.value, end="", flush=True)
 
     def on_error(self, data):
         print(f"Error: {data}")
 
 with project_client.agents.runs.stream(
     thread_id=thread.id,
     agent_id=agent.id,
     event_handler=MyHandler(),
 ) as stream:
     stream.until_done()
 ```
 
 ## File Operations
 
 ### Upload File
 
 ```python
 from azure.ai.agents.models import FilePurpose
 
 file = project_client.agents.files.upload_and_poll(
     file_path="./data/document.pdf",
     purpose=FilePurpose.AGENTS,
 )
 print(f"Uploaded file, ID: {file.id}")
 ```
 
 ### Create Vector Store
 
 ```python
 vector_store = project_client.agents.vector_stores.create_and_poll(
     file_ids=[file.id],
     name="my-vector-store",
 )
 print(f"Created vector store, ID: {vector_store.id}")
 ```
 
 ## Agent Lifecycle Best Practices
 
 ```python
 # 1. Use context managers
 with project_client:
     agent = project_client.agents.create_agent(...)
     thread = project_client.agents.threads.create()
     
     # ... use agent
     
     # Clean up
     project_client.agents.delete_agent(agent.id)
 
 # 2. For versioned agents, manage versions explicitly
 agent_v1 = project_client.agents.create_version(
     agent_name="my-agent",
     definition=
... (truncated)
```

### ✅ 13. web_fetch — 3536ms

**Input:**

```json
{"max_length":20000,"url":"https://pypi.org/pypi/azure-ai-agents/json"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python:
---
layout: Conceptual
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: Azure AI Agents client library for Python | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python
config_moniker_range: azure-python
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: azure-sdk
ms.author: azsdkdocs
ms.custom: tracking-python
keywords: Azure, python, SDK, API, azure-ai-agents, ai
ms.date: 2025-08-05T00:00:00.0000000Z
ms.topic: reference
ms.devlang: python
ms.service: ai
locale: en-us
document_id: c2011553-f14b-0d9f-d224-50ab688ed440
document_version_independent_id: dd35f162-44d6-64bd-0404-335e27975b4c
updated_at: 2025-08-06T11:06:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-services/latest/ai-agents-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/b4bbad572385f402118fd29f0f00e86ec3766d94/docs-ref-services/latest/ai-agents-readme.md
git_commit_id: b4bbad572385f402118fd29f0f00e86ec3766d94
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
page_type: conceptual
toc_rel: ../../toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 4864
asset_id: api/overview/azure/ai-agents-readme
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-services/latest/ai-agents-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/540ac133-a371-4dbb-8f94-28d6cc77a70b
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
- https://microsoft-devrel.poolparty.biz/DevRelOfferingOntology/1
... (truncated)
```

### ✅ 14. web_fetch — 2728ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-ai-agents/json:
{"info":{"author":"Microsoft Corporation","author_email":"azpysdkhelp@microsoft.com","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","License :: OSI Approved :: MIT License","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.9"],"description":"<!-- PIPY LONG DESCRIPTION BEGIN -->\n# Azure AI Agents client library for Python\n\nUse the AI Agents client library to:\n\n* **Develop Agents using the Azure AI Agents Service**, leveraging an extensive ecosystem of models, tools, and capabilities from OpenAI, Microsoft, and other LLM providers. The Azure AI Agents Service enables the building of Agents for a wide range of generative AI use cases.\n* **Note:** While this package can be used independently, we recommend using the Azure AI Projects client library (azure-ai-projects) for an enhanced experience. \nThe Projects library provides simplified access to advanced functionality, such as creating and managing agents, enumerating AI models, working with datasets and \nmanaging search indexes, evaluating generative AI performance, and enabling OpenTelemetry tracing.\n\n[Product documentation](https://aka.ms/azsdk/azure-ai-agents/product-doc)\n| [Samples][samples]\n| [API reference documentation](https://aka.ms/azsdk/azure-ai-agents/python/reference)\n| [Package (PyPI)](https://aka.ms/azsdk/azure-ai-agents/python/package)\n| [SDK source code](https://aka.ms/azsdk/azure-ai-agents/python/code)\n| [AI Starter Template](https://aka.ms/azsdk/azure-ai-agents/python/ai-starter-template)\n\n## Reporting issues\n\nTo report an issue with the client library, or request additional features
... (truncated)
```

### ✅ 15. powershell — 3146ms

**Input:**

```json
{"command":"$data = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787797413191-copilot-tool-output-b94079b60da44f9ead788164430a0ad0.txt' | ConvertFrom-Json; $data.results | Where-Object { $_.codeSnippet -match 'AgentsClient|create_run|get_run|RunStatus|ListSortOrder' -or $_.content -match 'AgentsClient|create_run|get_run|RunStatus|ListSortOrder' } | Select-Object -First 5 | ConvertTo-Json -Depth 8","description":"Extract relevant official SDK results"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-b94079b60da44f9ead788164430a0ad0.txt' | ConvertFrom-Json; $data. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 16. web_fetch — 1276ms

**Input:**

```json
{"max_length":15000,"start_index":19000,"url":"https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python:
e-ai-agents_1.1.0/sdk/ai/azure-ai-agents/samples/agents_streaming/sample_agents_stream_eventhandler_with_functions.py) or [`sample_agents_functions.py`](https://github.com/Azure/azure-sdk-for-python/blob/azure-ai-agents_1.1.0/sdk/ai/azure-ai-agents/samples/agents_tools/sample_agents_functions.py).

### Create Agent With Azure Function Call

The AI agent leverages Azure Functions triggered asynchronously via Azure Storage Queues. To enable the agent to perform Azure Function calls, you must set up the corresponding `AzureFunctionTool`, specifying input and output queues as well as parameter definitions.

Example Python snippet illustrating how you create an agent utilizing the Azure Function Tool:

```python
azure_function_tool = AzureFunctionTool(
    name="foo",
    description="Get answers from the foo bot.",
    parameters={
        "type": "object",
        "properties": {
            "query": {"type": "string", "description": "The question to ask."},
            "outputqueueuri": {"type": "string", "description": "The full output queue uri."},
        },
    },
    input_queue=AzureFunctionStorageQueue(
        queue_name="azure-function-foo-input",
        storage_service_endpoint=storage_service_endpoint,
    ),
    output_queue=AzureFunctionStorageQueue(
        queue_name="azure-function-tool-output",
        storage_service_endpoint=storage_service_endpoint,
    ),
)

agent = agents_client.create_agent(
    model=os.environ["MODEL_DEPLOYMENT_NAME"],
    name="azure-function-agent-foo",
    instructions=f"You are a helpful support agent. Use the provided function any time the prompt contains the string 'What would foo say?'. When you invoke the function, ALWAYS specify the output queue uri parameter as '{storage_service_endpoint}/azure-function-tool-output'. Always responds with \"Foo says\" and then the response from t
... (truncated)
```

### ✅ 17. web_fetch — 1349ms

**Input:**

```json
{"max_length":12000,"start_index":33000,"url":"https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/overview/azure/ai-agents-readme?view=azure-python:
ed file**
- **Using a public image accessible via URL**
- **Using a base64 encoded image string**

The following examples demonstrate each method:

#### Create message using uploaded image file

```python
# Upload the local image file
image_file = agents_client.files.upload_and_poll(file_path="image_file.png", purpose="assistants")

# Construct content using uploaded image
file_param = MessageImageFileParam(file_id=image_file.id, detail="high")
content_blocks = [
    MessageInputTextBlock(text="Hello, what is in the image?"),
    MessageInputImageFileBlock(image_file=file_param),
]

# Create the message
message = agents_client.messages.create(
    thread_id=thread.id,
    role="user",
    content=content_blocks
)
```

#### Create message with an image URL input

```python
# Specify the public image URL
image_url = "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"

# Create content directly referencing image URL
url_param = MessageImageUrlParam(url=image_url, detail="high")
content_blocks = [
    MessageInputTextBlock(text="Hello, what is in the image?"),
    MessageInputImageUrlBlock(image_url=url_param),
]

# Create the message
message = agents_client.messages.create(
    thread_id=thread.id,
    role="user",
    content=content_blocks
)
```

#### Create message with base64-encoded image input

```python
import base64

def image_file_to_base64(path: str) -> str:
    with open(path, "rb") as f:
        return base64.b64encode(f.read()).decode("utf-8")

# Convert your image file to base64 format
image_base64 = image_file_to_base64("image_file.png")

# Prepare the data URL
img_data_url = f"data:image/png;base64,{image_base64}"

# Use base64 encoded string as image URL parameter
url_param = 
... (truncated)
```

### ✅ 18. github-mcp-server-search_code (via github-mcp-server) — 5630ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/ai/azure-ai-agents/samples \"runs.create(thread_id=thread.id\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"sample_agents_multiple_connected_agents.py","path":"sdk/ai/azure-ai-agents/samples/agents_tools/sample_agents_multiple_connected_agents.py","sha":"9c042eae614763be0ef5baca741794fbf2ac49dc","text_matches":[{"fragment":"        if msg.text_messages:\n            last_text = msg.text_messages[-1]\n            text = last_text.text.value.replace(\"\\u3010\", \"[\").replace(\"\\u3011\", \"]\")\n            print(f\"{msg.role}: {text}\")\n    # [END list_messages]\n\n    agents_client.threads.delete(thread.id)","matches":[{"indices":[248,262],"text":"threads.delete"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-agents/samples/agents_tools/sample_agents_multiple_connected_agents.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]},{"name":"test_agents_client.py","path":"sdk/ai/azure-ai-agents/tests/test_agents_client.py","sha":"b8098a151f3cb70d8ea920bcdea6b2b9bd658ad3","text_matches":[{"fragment":"            # delete thread\n            client.threads.delete(thread.id)\n","matches":[{"indices":[47,61],"text":"threads.delete"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-agents/tests/test_agents_client.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"},{"fragment":"            client.delete_agent(agent.id)\n            client.threads.delete(thread.id)\n","matches":[{"indices":[61,75],"text":"threads.delete"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-agents/tests/test_agents_client.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]},{"name":"test_agents_client_async.py","path":"sdk/ai/azure-ai-agents/tests/test_agents_client_async.py","sha":"fb48e0d1dfe73ef717630c5a861e8761e8a9688b","text_matches":[{"fragment":"            # delete thread\n            await client.thre
... (truncated)
```

### ✅ 19. github-mcp-server-search_code (via github-mcp-server) — 5292ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/ai/azure-ai-agents \"class RunStatus\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"_enums.py","path":"sdk/ai/azure-ai-agents/azure/ai/agents/models/_enums.py","sha":"210bbe61530ac4924b67a432eeb1fa396c2c3ffe","text_matches":[{"fragment":"    \"\"\"File search result content.\"\"\"\n\n\nclass RunStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):\n    \"\"\"Possible values for the status of an agent thread run.\"\"\"\n\n    QUEUED = \"queued\"","matches":[{"indices":[40,55],"text":"class RunStatus"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-agents/azure/ai/agents/models/_enums.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]}],"total_count":1}
```

### ✅ 20. github-mcp-server-search_code (via github-mcp-server) — 4864ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"repo:Azure/azure-sdk-for-python path:sdk/ai/azure-ai-agents \"threads.delete\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"sample_agents_basics.py","path":"sdk/ai/azure-ai-agents/samples/sample_agents_basics.py","sha":"f5cdc8e56262a0f9ec6ac17d888c29c002b5d850","text_matches":[{"fragment":"    print(f\"Created message, message ID: {message.id}\")\n\n    # [START create_run]\n    run = agents_client.runs.create(thread_id=thread.id, agent_id=agent.id)\n\n    # Poll the run as long as run status is queued or in progress\n    while run.status in [\"queued\", \"in_progress\", \"requires_action\"]:","matches":[{"indices":[106,137],"text":"runs.create(thread_id=thread.id"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-agents/samples/sample_agents_basics.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]},{"name":"sample_agents_basics_async.py","path":"sdk/ai/azure-ai-agents/samples/agents_async/sample_agents_basics_async.py","sha":"d1ff46d873257a28c8e044ec246428fdd1f44628","text_matches":[{"fragment":"        message = await agents_client.messages.create(thread_id=thread.id, role=\"user\", content=\"Hello, tell me a joke\")\n        print(f\"Created message, message ID: {message.id}\")\n\n        run = await agents_client.runs.create(thread_id=thread.id, agent_id=agent.id)\n\n        # Poll the run as long as run status is queued or in progress\n        while run.status in [\"queued\", \"in_progress\", \"requires_action\"]:","matches":[{"indices":[216,247],"text":"runs.create(thread_id=thread.id"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/4127088/contents/sdk/ai/azure-ai-agents/samples/agents_async/sample_agents_basics_async.py?ref=1c7777f71f32868eb1ca36fcab099ac3d38c5014","property":"content"}]},{"name":"sample_agents_functions.py","path":"sdk/ai/azure-ai-agents/samples/agents_tools/sample_agents_functions.py","sha":"573fb800cd7c3fe3ea8d2c6e3d14eec83fa98813","text_matches":[{"fragment":"    )\n    print(f\"Created message, ID:
... (truncated)
```

### ✅ 21. powershell — 5997ms

**Input:**

```json
{"command":"python -m pip index versions azure-ai-agents --disable-pip-version-check","description":"Check available Agents SDK versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
azure-ai-agents (1.1.0)
Available versions: 1.1.0, 1.0.2, 1.0.1, 1.0.0
  INSTALLED: 1.1.0
  LATEST:    1.1.0
<shellId: 1 completed with exit code 0>
```

### ❌ 22. github-mcp-server-get_file_contents (via github-mcp-server) — 4294ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-agents/samples/sample_agents_basics.py","repo":"azure-sdk-for-python","sha":"f5cdc8e56262a0f9ec6ac17d888c29c002b5d850"}
```

**Error:**

```
MCP server 'github-mcp-server': failed to get git tree: Invalid object requested. SHA must identify a commit or a tree.
```

### ❌ 23. github-mcp-server-get_file_contents (via github-mcp-server) — 4292ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-agents/azure/ai/agents/models/_enums.py","repo":"azure-sdk-for-python","sha":"210bbe61530ac4924b67a432eeb1fa396c2c3ffe"}
```

**Error:**

```
MCP server 'github-mcp-server': failed to get git tree: Invalid object requested. SHA must identify a commit or a tree.
```

### ✅ 24. powershell — 4285ms

**Input:**

```json
{"command":"python -c \"import inspect; from azure.ai.agents import AgentsClient; from azure.ai.agents.models import RunStatus, ListSortOrder, MessageRole; import azure.ai.agents, azure.identity; print('agents', azure.ai.agents.__version__); print('identity', azure.identity.__version__); print('RunStatus', list(RunStatus)); print('MessageRole', list(MessageRole)); print('runs.get', inspect.signature(AgentsClient.runs.fget) if False else 'operations-bound-at-instance')\"","description":"Inspect installed SDK enums and versions"}
```

**Output:**

```
agents 1.1.0
identity 1.25.1
RunStatus [<RunStatus.QUEUED: 'queued'>, <RunStatus.IN_PROGRESS: 'in_progress'>, <RunStatus.REQUIRES_ACTION: 'requires_action'>, <RunStatus.CANCELLING: 'cancelling'>, <RunStatus.CANCELLED: 'cancelled'>, <RunStatus.FAILED: 'failed'>, <RunStatus.COMPLETED: 'completed'>, <RunStatus.EXPIRED: 'expired'>]
MessageRole [<MessageRole.USER: 'user'>, <MessageRole.AGENT: 'assistant'>]
runs.get operations-bound-at-instance
<shellId: 2 completed with exit code 0>
```

### ✅ 25. github-mcp-server-get_file_contents (via github-mcp-server) — 3630ms

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/azure-ai-agents/samples/sample_agents_basics.py","ref":"refs/heads/main","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: f5cdc8e56262a0f9ec6ac17d888c29c002b5d850)# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------

"""
DESCRIPTION:
    This sample demonstrates how to use basic agent operations from
    the Azure Agents service using a synchronous client.

USAGE:
    python sample_agents_basics.py

    Before running the sample:

    pip install azure-ai-projects azure-ai-agents azure-identity

    Set these environment variables with your own values:
    1) PROJECT_ENDPOINT - The Azure AI Project endpoint, as found in the Overview
                          page of your Azure AI Foundry portal.
    2) MODEL_DEPLOYMENT_NAME - The deployment name of the AI model, as found under the "Name" column in
       the "Models + endpoints" tab in your Azure AI Foundry project.
"""

import os, time
from azure.ai.projects import AIProjectClient
from azure.identity import DefaultAzureCredential
from azure.ai.agents.models import ListSortOrder

project_client = AIProjectClient(
    endpoint=os.environ["PROJECT_ENDPOINT"],
    credential=DefaultAzureCredential(),
)

with project_client:
    agents_client = project_client.agents

    # [START create_agent]
    agent = agents_client.create_agent(
        model=os.environ["MODEL_DEPLOYMENT_NAME"],
        name="my-agent",
        instructions="You are helpful agent",
    )
    # [END create_agent]
    print(f"Created agent, agent ID: {agent.id}")

    # [START create_thread]
    thread = agents_client.threads.create()
    # [END create_thread]
    print(f"Created thread, thread ID: {thread.id}")

    # List all threads for the agent
    # [START list_threads]
    threads = agents_client.threads.list()
    # [END list_threads]

    # [START create_message]
    message = agents_client.messages.create(thread_id=thread.id, role="user", content="Hello, tell me a joke")
    # [END create_message]
    print(f"Created message,
... (truncated)
```

### ✅ 26. powershell — 3220ms

**Input:**

```json
{"command":"python -c \"import inspect; from azure.ai.agents import AgentsClient; from azure.identity import DefaultAzureCredential; c=AgentsClient(endpoint='https://example.services.ai.azure.com/api/projects/example', credential=DefaultAzureCredential()); print('runs.create', inspect.signature(c.runs.create)); print('runs.get', inspect.signature(c.runs.get)); print('messages.list', inspect.signature(c.messages.list)); print('threads.delete', inspect.signature(c.threads.delete)); print('delete_agent', inspect.signature(c.delete_agent)); c.close()\"","description":"Inspect installed client method signatures"}
```

**Output:**

```
runs.create (thread_id: str, body: Union[collections.abc.MutableMapping[str, Any], IO[bytes]] = <object object at 0x0000027E1DC41310>, *, include: Optional[List[Union[str, azure.ai.agents.models._enums.RunAdditionalFieldList]]] = None, agent_id: str = <object object at 0x0000027E1DC41310>, model: Optional[str] = None, instructions: Optional[str] = None, additional_instructions: Optional[str] = None, additional_messages: Optional[List[azure.ai.agents.models._models.ThreadMessageOptions]] = None, tools: Optional[List[azure.ai.agents.models._models.ToolDefinition]] = None, tool_resources: Optional[azure.ai.agents.models._models.ToolResources] = None, temperature: Optional[float] = None, top_p: Optional[float] = None, max_prompt_tokens: Optional[int] = None, max_completion_tokens: Optional[int] = None, truncation_strategy: Optional[azure.ai.agents.models._models.TruncationObject] = None, tool_choice: Optional[ForwardRef('_types.AgentsToolChoiceOption')] = None, response_format: Optional[ForwardRef('_types.AgentsResponseFormatOption')] = None, parallel_tool_calls: Optional[bool] = None, metadata: Optional[Dict[str, str]] = None, **kwargs: Any) -> azure.ai.agents.models._models.ThreadRun
runs.get (thread_id: str, run_id: str, **kwargs: Any) -> azure.ai.agents.models._models.ThreadRun
messages.list (thread_id: str, *, run_id: Optional[str] = None, limit: Optional[int] = None, order: Union[str, azure.ai.agents.models._enums.ListSortOrder, NoneType] = None, before: Optional[str] = None, **kwargs: Any) -> azure.core.paging.ItemPaged['_models.ThreadMessage']
threads.delete (thread_id: str, **kwargs: Any) -> None
delete_agent (agent_id: str, **kwargs: Any) -> None
<shellId: 3 completed with exit code 0>
```

### ✅ 27. apply_patch — 886ms

**Input:**

```json
"*** Begin Patch\n*** Add File: main.py\n+import os\n+import time\n+\n+from azure.ai.agents import AgentsClient\n+from azure.ai.agents.models import ListSortOrder, MessageRole, RunStatus\n+from azure.identity import DefaultAzureCredential\n+\n+\n+POLL_INTERVAL_SECONDS = 1\n+TERMINAL_RUN_STATUSES = {\n+    RunStatus.CANCELLED,\n+    RunStatus.COMPLETED,\n+    RunStatus.EXPIRED,\n+    RunStatus.FAILED,\n+}\n+\n+\n+def main() -\u003e None:\n+    project_endpoint = os.environ[\"PROJECT_ENDPOINT\"]\n+    model_deployment_name = os.environ[\"MODEL_DEPLOYMENT_NAME\"]\n+\n+    agent_id: str | None = None\n+    thread_id: str | None = None\n+\n+    with DefaultAzureCredential() as credential:\n+        with AgentsClient(endpoint=project_endpoint, credential=credential) as client:\n+            try:\n+                agent = client.create_agent(\n+                    model=model_deployment_name,\n+                    name=\"hyoka-basic-agent\",\n+                    instructions=\"Answer the user's question clearly and concisely.\",\n+                )\n+                agent_id = agent.id\n+\n+                thread = client.threads.create()\n+                thread_id = thread.id\n+\n+                client.messages.create(\n+                    thread_id=thread_id,\n+                    role=MessageRole.USER,\n+                    content=\"What is the capital of France?\",\n+                )\n+\n+                run = client.runs.create(thread_id=thread_id, agent_id=agent_id)\n+                while run.status not in TERMINAL_RUN_STATUSES:\n+                    time.sleep(POLL_INTERVAL_SECONDS)\n+                    run = client.runs.get(thread_id=thread_id, run_id=run.id)\n+\n+                if run.status != RunStatus.COMPLETED:\n+                    error_detail = f\": {run.last_error}\" if run.last_error else \"\"\n+                    raise RuntimeError(f\"Agent run ended with status '{run.status}'{error_detail}\")\n+\n+                messages = client.messages.list(\n+                    thread_id=thread_id,\n+                    order=ListSortOrder.ASCENDING,\n+                )\n+                for message in messages:\n+                    if message.role == MessageRole.AGENT:\n+                        for text_message in message.text_messages:\n+                            print(text_message.text.value)\n+            finally:\n+                if thread_id is not None:\n+                    try:\n+                        client.threads.delete(thread_id)\n+                    finally:\n+                        if agent_id is not None:\n+                            client.delete_agent(agent_id)\n+                elif agent_id is not None:\n+                    client.delete_agent(agent_id)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: pyproject.toml\n+[build-system]\n+requires = [\"setuptools\u003e=77\"]\n+build-backend = \"setuptools.build_meta\"\n+\n+[project]\n+name = \"hyoka-basic-agent\"\n+version = \"0.1.0\"\n+description = \"A basic synchronous Azure AI Agent console conversation.\"\n+requires-python = \"\u003e=3.9\"\n+dependencies = [\n+    \"azure-ai-agents==1.1.0\",\n+    \"azure-identity==1.25.1\",\n+]\n+\n+[project.scripts]\n+hyoka-basic-agent = \"main:main\"\n+\n+[tool.setuptools]\n+py-modules = [\"main\"]\n*** Add File: README.md\n+# Hyoka Basic Azure AI Agent\n+\n+A synchronous Python console application that creates an Azure AI Agent, asks it\n+for the capital of France, prints every assistant text response, and deletes the\n+temporary thread and agent.\n+\n+## Prerequisites\n+\n+- Python 3.9 or later\n+- An Azure AI Foundry project and model deployment\n+- A Microsoft Entra identity with access to the project\n+- Local authentication available to `DefaultAzureCredential`\n+\n+## Restore\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\python -m pip install --upgrade pip\n+.\\.venv\\Scripts\\python -m pip install -e .\n+```\n+\n+## Build\n+\n+```powershell\n+.\\.venv\\Scripts\\python -m pip wheel . --no-deps --wheel-dir dist\n+```\n+\n+## Run\n+\n+```powershell\n+$env:PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:MODEL_DEPLOYMENT_NAME = \"\u003cmodel-deployment-name\u003e\"\n+.\\.venv\\Scripts\\hyoka-basic-agent.exe\n+```\n*** Add File: .gitignore\n+.venv/\n+__pycache__/\n+*.py[cod]\n+build/\n+dist/\n+*.egg-info/\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py
@@ -1,0 +1,74 @@
+import os
+import time
+
+from azure.ai.agents import AgentsClient
+from azure.ai.agents.models import ListSortOrder, MessageRole, RunStatus
+from azure.identity import DefaultAzureCredential
+
+
+POLL_INTERVAL_SECONDS = 1
+TERMINAL_RUN_STATUSES = {
+    RunStatus.CANCELLED,
+    RunStatus.COMPLETED,
+    RunStatus.EXPIRED,
+    RunStatus.FAILED,
+}
+
+
+def main() -> None:
+    project_endpoint = os.environ["PROJECT_ENDPOINT"]
+    model_deployment_name = os.environ["MODEL_DEPLOYMENT_NAME"]
+
+    agent_id: str | None = None
+    thread_id: str | None = None
+
+    with DefaultAzureCredential() as credential:
+        with AgentsClient(endpoint=project_endpoint, credential=credential) as client:
+            try:
+                agent = client.create_agent(
+                    model=model_deployment_name,
+                    name="hyoka-basic-agent",
+                    instructions="Answer the user's question clearly and concisely.",
+                )
+                agent_id = agent.id
+
+                thread = client.threads.create()
+                thread_id = thread.id
+
+                client.messages.create(
+                    thread_id=thread_id,
+                    role=MessageRole.USER,
+                    content="What is the capital of France?",
+                )
+
+                run = client.runs.create(thread_id=thread_id, agent_id=agent_id)
+                while run.status not in TERMINAL_RUN_STATUSES:
+                    tim
... (truncated)
```

### ✅ 28. apply_patch — 854ms

**Input:**

```json
"*** Begin Patch\n*** Update File: main.py\n@@\n import os\n import time\n+from typing import Optional\n@@\n-    agent_id: str | None = None\n-    thread_id: str | None = None\n+    agent_id: Optional[str] = None\n+    thread_id: Optional[str] = None\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991/main.py
@@ -1,5 +1,6 @@
 import os
 import time
+from typing import Optional
 
 from azure.ai.agents import AgentsClient
 from azure.ai.agents.models import ListSortOrder, MessageRole, RunStatus
@@ -19,8 +20,8 @@
     project_endpoint = os.environ["PROJECT_ENDPOINT"]
     model_deployment_name = os.environ["MODEL_DEPLOYMENT_NAME"]
 
-    agent_id: str | None = None
-    thread_id: str | None = None
+    agent_id: Optional[str] = None
+    thread_id: Optional[str] = None
 
     with DefaultAzureCredential() as credential:
         with AgentsClient(endpoint=project_endpoint, credential=credential) as client:

```

### ✅ 29. powershell — 17321ms

**Input:**

```json
{"command":"python -m compileall -q main.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -m pip wheel . --no-deps --wheel-dir dist --disable-pip-version-check --quiet","description":"Compile app and build wheel","initial_wait":120,"mode":"sync"}
```

**Output:**

```

<shellId: 4 completed with exit code 0>
```

### ✅ 30. powershell — 16915ms

**Input:**

```json
{"command":"@'\nimport contextlib\nimport io\nfrom types import SimpleNamespace\nfrom unittest.mock import MagicMock, patch\n\nimport main\n\ncredential = MagicMock()\ncredential.__enter__.return_value = credential\nclient = MagicMock()\nclient.__enter__.return_value = client\nclient.create_agent.return_value = SimpleNamespace(id=\"agent-1\")\nclient.threads.create.return_value = SimpleNamespace(id=\"thread-1\")\nclient.runs.create.return_value = SimpleNamespace(id=\"run-1\", status=main.RunStatus.QUEUED)\nclient.runs.get.return_value = SimpleNamespace(\n    id=\"run-1\", status=main.RunStatus.COMPLETED, last_error=None\n)\ntext = SimpleNamespace(text=SimpleNamespace(value=\"Paris is the capital of France.\"))\nclient.messages.list.return_value = [\n    SimpleNamespace(role=main.MessageRole.USER, text_messages=[]),\n    SimpleNamespace(role=main.MessageRole.AGENT, text_messages=[text]),\n]\n\noutput = io.StringIO()\nwith patch.dict(main.os.environ, {\n    \"PROJECT_ENDPOINT\": \"https://example.services.ai.azure.com/api/projects/example\",\n    \"MODEL_DEPLOYMENT_NAME\": \"example-model\",\n}, clear=True), patch.object(main, \"DefaultAzureCredential\", return_value=credential), patch.object(\n    main, \"AgentsClient\", return_value=client\n), patch.object(main.time, \"sleep\"), contextlib.redirect_stdout(output):\n    main.main()\n\nassert output.getvalue() == \"Paris is the capital of France.\\n\"\nclient.create_agent.assert_called_once_with(\n    model=\"example-model\",\n    name=\"hyoka-basic-agent\",\n    instructions=\"Answer the user's question clearly and concisely.\",\n)\nclient.messages.create.assert_called_once_with(\n    thread_id=\"thread-1\",\n    role=main.MessageRole.USER,\n    content=\"What is the capital of France?\",\n)\nclient.runs.get.assert_called_once_with(thread_id=\"thread-1\", run_id=\"run-1\")\nclient.threads.delete.assert_called_once_with(\"thread-1\")\nclient.delete_agent.assert_called_once_with(\"agent-1\")\nprint(\"mock lifecycle passed\")\n'@ | python -","description":"Test agent lifecycle with mocks","initial_wait":120,"mode":"sync"}
```

**Output:**

```
mock lifecycle passed
<shellId: 5 completed with exit code 0>
```

### ✅ 31. glob — 1121ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\__pycache__\main.cpython-313.pyc
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\README.md
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\pyproject.toml
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\main.py
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\hyoka_basic_agent.egg-info\top_level.txt
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\hyoka_basic_agent.egg-info\SOURCES.txt
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\hyoka_basic_agent.egg-info\requires.txt
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\hyoka_basic_agent.egg-info\PKG-INFO
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\hyoka_basic_agent.egg-info\entry_points.txt
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\hyoka_basic_agent.egg-info\dependency_links.txt
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\dist\hyoka_basic_agent-0.1.0-py3-none-any.whl
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-basic-agent-lifecycle-python-azure-tools-with-azure-tools-2966204991\build\lib\main.py
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-agents-dp-python-b
... (truncated)
```

### 🔧 32. powershell

**Input:**

```json
{"command":"Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force; Remove-Item -LiteralPath '.\\build' -Recurse -Force; Remove-Item -LiteralPath '.\\dist' -Recurse -Force; Remove-Item -LiteralPath '.\\hyoka_basic_agent.egg-info' -Recurse -Force","description":"Remove generated build artifacts"}
```

## Generated Files

- `README.md`
- `hyoka_basic_agent.egg-info\PKG-INFO`
- `hyoka_basic_agent.egg-info\SOURCES.txt`
- `hyoka_basic_agent.egg-info\dependency_links.txt`
- `hyoka_basic_agent.egg-info\entry_points.txt`
- `hyoka_basic_agent.egg-info\requires.txt`
- `hyoka_basic_agent.egg-info\top_level.txt`
- `main.py`
- `pyproject.toml`

## Copilot Response

I’m aligning the implementation with the current low-level Agents SDK lifecycle and cleanup semantics before writing the project.

The current SDK exposes explicit `RunStatus` terminal values and ascending message ordering, so I’m implementing manual synchronous polling with guaranteed thread-and-agent cleanup.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Resource cleanup is attempted with finally blocks.

### Issues

- Azure SDK exceptions are not caught or handled appropriately.

## Grader Results

- basic-agent-lifecycle.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (10/10)
      - Creates `AgentsClient` from `azure.ai.agents` with the project endpoint.: Pass
      - Calls `agents_client.create_agent` with the deployment from: Pass
      - Creates a thread with `agents_client.threads.create`.: Pass
      - Adds the exact user message with `agents_client.messages.create`, role `user`, and: Pass
      - Creates the run with the created thread ID and agent ID, then polls with: Pass
      - Retrieves messages only after successful completion, requests: Pass
      - Deletes both the created thread and the created agent through `AgentsClient`.: Pass
      - Does not print a hardcoded answer instead of retrieving agent messages.: Pass
      - Does not treat the initially created run as already completed.: Pass
      - Does not substitute an Azure OpenAI chat client or another non-agents API for the: Pass
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
hyoka run --prompt-id ai-agents-dp-python-basic-agent-lifecycle --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
