# Evaluation Report: service-bus-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 208.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 11/14 |
| Duration | 208.8s |
| Timestamp | 2026-08-28T21:36:56Z |
| Files Generated | 4 |
| Event Count | 6319 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 125.7s |
| Review | 81.9s |
| **Total** | **208.8s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-servicebus-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=440864 out=5031 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 3638ms

**Input:**

```json
{"skill":"azure-servicebus-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Service Bus SDK for Python

Enterprise messaging for reliable cloud communication with queues and pub/sub topics.

## Installation

```bash
pip install azure-servicebus azure-identity
```

## Environment Variables

```bash
SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net  # Required for all auth methods
SERVICEBUS_QUEUE_NAME=myqueue  # Required for queue operations
SERVICEBUS_TOPIC_NAME=mytopic  # Required for topic operations
SERVICEBUS_SUBSCRIPTION_NAME=mysubscription  # Required for subscription operations
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
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential
from azure.servicebus import ServiceBusClient

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
credential = DefaultAzureCredential(require_envvar=True)
# Or use a specific creden
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 8729ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Create local Python Azure Service Bus queue and topic/subscription examples using sync and aio clients, DefaultAzureCredential, batches, receive-and-complete settlement, and context managers."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. glob — 9964ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-servicebus-py\references\patterns.md
```

### ✅ 4. glob — 7513ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py","pattern":"**/{acceptance-criteria.md,examples.md,patterns.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 4587ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation guidance for Python Service Bus examples.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. view — 3230ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py\\references\\patterns.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md
@@ -1,406 +1,406 @@
 # Messaging Patterns Reference
 
 Advanced messaging patterns for Azure Service Bus.
 
 ## Competing Consumers
 
 Multiple receivers processing messages from the same queue in parallel:
 
 ```python
 import asyncio
 from azure.servicebus.aio import ServiceBusClient
 from azure.identity.aio import DefaultAzureCredential
 
 async def worker(worker_id: int, namespace: str, queue_name: str):
     """Worker that processes messages from a shared queue."""
     credential = DefaultAzureCredential()
     
     async with ServiceBusClient(
         fully_qualified_namespace=namespace,
         credential=credential
     ) as client:
         receiver = client.get_queue_receiver(queue_name=queue_name)
         
         async with receiver:
             while True:
                 messages = await receiver.receive_messages(
                     max_message_count=10,
                     max_wait_time=5
                 )
                 
                 if not messages:
                     continue
                 
                 for msg in messages:
                     try:
                         print(f"Worker {worker_id}: Processing {str(msg)}")
                         await process_message(msg)
                         await receiver.complete_message(msg)
                     except Exception as e:
                        
... (truncated)
```

### ✅ 7. web_fetch — 6181ms

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-servicebus/json"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/python/api/overview/azure/servicebus-readme?view=azure-python:
---
layout: Conceptual
monikers:
- azure-python
defaultMoniker: azure-python
versioningType: Ranged
title: Azure Service Bus client library for Python | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/python/api/overview/azure/servicebus-readme?view=azure-python
config_moniker_range: azure-python
uhfHeaderId: Azure
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-python/issues
breadcrumb_path: /python/azure/bread/toc.json
apiPlatform: python
author: azure-sdk
ms.author: azsdkdocs
ms.custom: tracking-python
keywords: Azure, python, SDK, API, azure-servicebus, azure-service-bus
ms.date: 2025-10-31T00:00:00.0000000Z
ms.topic: reference
ms.devlang: python
ms.service: azure-service-bus
locale: en-us
document_id: 8fa61aeb-289e-7b6a-bd70-ba4055d456b0
document_version_independent_id: ff2b2f4b-22df-7960-f3db-6d3e67d25f00
updated_at: 2025-10-31T11:08:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/live/docs-ref-services/latest/servicebus-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-python/blob/bca6bb664d34f522857145bdb8cbfadb505da89c/docs-ref-services/latest/servicebus-readme.md
git_commit_id: bca6bb664d34f522857145bdb8cbfadb505da89c
default_moniker: azure-python
site_name: Docs
depot_name: MSDN.python-sdk
page_type: conceptual
toc_rel: ../../toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 3932
asset_id: api/overview/azure/servicebus-readme
moniker_range_name: db4cc4146095cc059ef03e10f2246414
monikers:
- azure-python
item_type: Content
source_path: docs-ref-services/latest/servicebus-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/828c518e-a2df-46f1-8871-b760959c8fa8
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
- https://authoring-docs-m
... (truncated)
```

### ✅ 8. web_fetch — 5799ms

**Input:**

```json
{"max_length":1000,"url":"https://pypi.org/pypi/azure-identity/json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-servicebus/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com> License-Expression: MIT","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.9"],"description":"# Azure Service Bus client library for Python\n\nAzure Service Bus is a high performance cloud-managed messaging service for providing real-time and fault-tolerant communication between distributed senders and receivers.\n\nService Bus provides multiple mechanisms for asynchronous highly reliable communication, such as structured first-in-first-out messaging,\npublish/subscribe capabilities, and the ability to easily scale as your needs grow.\n\nUse the Service Bu

<note>Content truncated. Call the fetch tool with a start_index of 1000 to get more content.</note>
```

### ✅ 9. web_fetch — 5407ms

**Input:**

```json
{"max_length":2000,"url":"https://learn.microsoft.com/en-us/python/api/overview/azure/servicebus-readme?view=azure-python"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://pypi.org/pypi/azure-identity/json:
{"info":{"author":null,"author_email":"Microsoft Corporation <azpysdkhelp@microsoft.com> License-Expression: MIT","bugtrack_url":null,"classifiers":["Development Status :: 5 - Production/Stable","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.12","Programming Language :: Python :: 3.13","Programming Language :: Python :: 3.14","Programming Language :: Python :: 3.9"],"description":"# Azure Identity client library for Python\n\nThe Azure Identity library provides [Microsoft Entra ID](https://learn.microsoft.com/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [`TokenCredential`][token_cred_ref]/[`SupportsTokenInfo`][supports_token_info_ref] implementations, which can be used to construct Azure SDK clients that support Microsoft Entra token authent

<note>Content truncated. Call the fetch tool with a start_index of 1000 to get more content.</note>
```

### ✅ 10. powershell — 5668ms

**Input:**

```json
{"command":"$sb = Invoke-RestMethod -Uri 'https://pypi.org/pypi/azure-servicebus/json'; $id = Invoke-RestMethod -Uri 'https://pypi.org/pypi/azure-identity/json'; [pscustomobject]@{azure_servicebus=$sb.info.version; azure_identity=$id.info.version} | ConvertTo-Json -Compress","description":"Get latest Azure SDK package versions"}
```

**Output:**

```
{"azure_servicebus":"7.14.3","azure_identity":"1.25.3"}
<shellId: 0 completed with exit code 0>
```

### ✅ 11. apply_patch — 1874ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity==1.25.3\n+azure-servicebus==7.14.3\n*** Add File: sync_servicebus.py\n+import os\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+\n+\n+MESSAGE_COUNT = 5\n+RECEIVE_WAIT_SECONDS = 5\n+\n+\n+def required_setting(name: str) -\u003e str:\n+    value = os.environ.get(name)\n+    if not value:\n+        raise RuntimeError(f\"Set the {name} environment variable.\")\n+    return value\n+\n+\n+def process_message(source: str, message: ServiceBusMessage) -\u003e None:\n+    print(f\"{source}: {str(message)}\")\n+\n+\n+def send_queue_batch(\n+    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT\n+) -\u003e None:\n+    with client.get_queue_sender(queue_name=queue_name) as sender:\n+        batch = sender.create_message_batch()\n+        for index in range(1, message_count + 1):\n+            try:\n+                batch.add_message(\n+                    ServiceBusMessage(\n+                        f\"Queue message {index}\",\n+                        message_id=f\"sync-queue-{index}\",\n+                        content_type=\"text/plain\",\n+                    )\n+                )\n+            except ValueError as error:\n+                raise RuntimeError(\n+                    \"The five-message demonstration batch exceeds the entity size limit.\"\n+                ) from error\n+\n+        sender.send_messages(batch)\n+        print(f\"Sent {message_count} queue messages in one batch.\")\n+\n+\n+def receive_queue_messages(\n+    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT\n+) -\u003e None:\n+    with client.get_queue_receiver(queue_name=queue_name) as receiver:\n+        messages = receiver.receive_messages(\n+            max_message_count=message_count,\n+            max_wait_time=RECEIVE_WAIT_SECONDS,\n+        )\n+        for message in messages:\n+            process_message(\"queue\", message)\n+            receiver.complete_message(message)\n+\n+        print(f\"Completed {len(messages)} queue messages.\")\n+\n+\n+def send_topic_message(client: ServiceBusClient, topic_name: str) -\u003e None:\n+    with client.get_topic_sender(topic_name=topic_name) as sender:\n+        sender.send_messages(\n+            ServiceBusMessage(\n+                \"Topic message\",\n+                message_id=\"sync-topic-1\",\n+                content_type=\"text/plain\",\n+            )\n+        )\n+        print(\"Sent one topic message.\")\n+\n+\n+def receive_subscription_message(\n+    client: ServiceBusClient, topic_name: str, subscription_name: str\n+) -\u003e None:\n+    with client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+    ) as receiver:\n+        messages = receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=RECEIVE_WAIT_SECONDS,\n+        )\n+        for message in messages:\n+            process_message(\"subscription\", message)\n+            receiver.complete_message(message)\n+\n+        print(f\"Completed {len(messages)} subscription messages.\")\n+\n+\n+def main() -\u003e None:\n+    namespace = required_setting(\"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\")\n+    queue_name = required_setting(\"SERVICEBUS_QUEUE_NAME\")\n+    topic_name = required_setting(\"SERVICEBUS_TOPIC_NAME\")\n+    subscription_name = required_setting(\"SERVICEBUS_SUBSCRIPTION_NAME\")\n+\n+    with DefaultAzureCredential() as credential:\n+        with ServiceBusClient(\n+            fully_qualified_namespace=namespace,\n+            credential=credential,\n+        ) as client:\n+            send_queue_batch(client, queue_name)\n+            receive_queue_messages(client, queue_name)\n+            send_topic_message(client, topic_name)\n+            receive_subscription_message(client, topic_name, subscription_name)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: async_servicebus.py\n+import asyncio\n+import os\n+\n+from azure.identity.aio import DefaultAzureCredential\n+from azure.servicebus import ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient\n+from azure.servicebus.aio._base_handler_async import ServiceBusReceiver\n+\n+\n+MESSAGE_COUNT = 5\n+RECEIVE_WAIT_SECONDS = 5\n+\n+\n+def required_setting(name: str) -\u003e str:\n+    value = os.environ.get(name)\n+    if not value:\n+        raise RuntimeError(f\"Set the {name} environment variable.\")\n+    return value\n+\n+\n+async def process_message(source: str, message: ServiceBusMessage) -\u003e None:\n+    await asyncio.sleep(0)\n+    print(f\"{source}: {str(message)}\")\n+\n+\n+async def send_queue_batch(\n+    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT\n+) -\u003e None:\n+    sender = client.get_queue_sender(queue_name=queue_name)\n+    async with sender:\n+        batch = await sender.create_message_batch()\n+        for index in range(1, message_count + 1):\n+            try:\n+                batch.add_message(\n+                    ServiceBusMessage(\n+                        f\"Async queue message {index}\",\n+                        message_id=f\"async-queue-{index}\",\n+                        content_type=\"text/plain\",\n+                    )\n+                )\n+            except ValueError as error:\n+                raise RuntimeError(\n+                    \"The five-message demonstration batch exceeds the entity size limit.\"\n+                ) from error\n+\n+        await sender.send_messages(batch)\n+        print(f\"Sent {message_count} queue messages in one async batch.\")\n+\n+\n+async def process_and_complete(\n+    receiver: ServiceBusReceiver, source: str, message: ServiceBusMessage\n+) -\u003e None:\n+    await process_message(source, message)\n+    await receiver.complete_message(message)\n+\n+\n+async def receive_queue_messages(\n+    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT\n+) -\u003e None:\n+    receiver = client.get_queue_receiver(\n+        queue_name=queue_name,\n+        prefetch_count=message_count,\n+    )\n+    async with receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=message_count,\n+            max_wait_time=RECEIVE_WAIT_SECONDS,\n+        )\n+        await asyncio.gather(\n+            *(\n+                process_and_complete(receiver, \"queue\", message)\n+                for message in messages\n+            )\n+        )\n+        print(f\"Completed {len(messages)} queue messages concurrently.\")\n+\n+\n+async def send_topic_message(client: ServiceBusClient, topic_name: str) -\u003e None:\n+    sender = client.get_topic_sender(topic_name=topic_name)\n+    async with sender:\n+        await sender.send_messages(\n+            ServiceBusMessage(\n+                \"Async topic message\",\n+                message_id=\"async-topic-1\",\n+                content_type=\"text/plain\",\n+            )\n+        )\n+        print(\"Sent one async topic message.\")\n+\n+\n+async def receive_subscription_message(\n+    client: ServiceBusClient, topic_name: str, subscription_name: str\n+) -\u003e None:\n+    receiver = client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+    )\n+    async with receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=RECEIVE_WAIT_SECONDS,\n+        )\n+        await asyncio.gather(\n+            *(\n+                process_and_complete(receiver, \"subscription\", message)\n+                for message in messages\n+            )\n+        )\n+        print(f\"Completed {len(messages)} subscription messages.\")\n+\n+\n+async def main() -\u003e None:\n+    namespace = required_setting(\"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\")\n+    queue_name = required_setting(\"SERVICEBUS_QUEUE_NAME\")\n+    topic_name = required_setting(\"SERVICEBUS_TOPIC_NAME\")\n+    subscription_name = required_setting(\"SERVICEBUS_SUBSCRIPTION_NAME\")\n+\n+    async with DefaultAzureCredential() as credential:\n+        async with ServiceBusClient(\n+            fully_qualified_namespace=namespace,\n+            credential=credential,\n+        ) as client:\n+            await send_queue_batch(client, queue_name)\n+            await receive_queue_messages(client, queue_name)\n+            await send_topic_message(client, topic_name)\n+            await receive_subscription_message(\n+                client,\n+                topic_name,\n+                subscription_name,\n+            )\n+\n+\n+if __name__ == \"__main__\":\n+    asyncio.run(main())\n*** Add File: README.md\n+# Azure Service Bus Python messaging examples\n+\n+These scripts demonstrate queue batching, explicit message completion, asynchronous\n+processing, and topic/subscription messaging. They do not create or change Azure\n+resources; the queue, topic, and subscription must already exist.\n+\n+## Install\n+\n+Python 3.9 or newer is required.\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+Required packages:\n+\n+- `azure-servicebus==7.14.3`\n+- `azure-identity==1.25.3`\n+\n+## Configure\n+\n+Authenticate locally with a supported `DefaultAzureCredential` developer login.\n+In Azure, use managed identity. Grant the identity the Azure Service Bus Data Sender\n+and Azure Service Bus Data Receiver roles at the narrowest appropriate scope.\n+\n+```powershell\n+$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = \"\u003cnamespace\u003e.servicebus.windows.net\"\n+$env:SERVICEBUS_QUEUE_NAME = \"\u003cqueue-name\u003e\"\n+$env:SERVICEBUS_TOPIC_NAME = \"\u003ctopic-name\u003e\"\n+$env:SERVICEBUS_SUBSCRIPTION_NAME = \"\u003csubscription-name\u003e\"\n+```\n+\n+For production, also set `AZURE_TOKEN_CREDENTIALS=prod` to constrain\n+`DefaultAzureCredential` to production-safe credential types.\n+\n+## Run\n+\n+```powershell\n+python .\\sync_servicebus.py\n+python .\\async_servicebus.py\n+```\n+\n+Each script sends a five-message `ServiceBusMessageBatch` to the queue, receives up\n+to five messages, processes and completes each message, then publishes a topic\n+message and receives it from the configured subscription. Receives have bounded\n+wait times, so an empty entity does not block indefinitely. The async version uses\n+prefetching and concurrent processing/settlement for higher throughput.\n+\n+Reference: [Azure Service Bus client library for Python](https://learn.microsoft.com/en-us/python/api/overview/azure/servicebus-readme?view=azure-python)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/requirements.txt
@@ -1,0 +1,3 @@
+azure-identity==1.25.3
+azure-servicebus==7.14.3
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py
@@ -1,0 +1,109 @@
+import os
+
+from azure.identity import DefaultAzureCredential
+from azure.servicebus import ServiceBusClient, ServiceBusMessage
+
+
+MESSAGE_COUNT = 5
+RECEIVE_WAIT_SECONDS = 5
+
+
+def required_setting(name: str) -> str:
+    value = os.environ.get(name)
+    if not value:
+        raise RuntimeError(f"Set the {name} environment variable.")
+    return value
+
+
+def process_message(source: str, message: ServiceBusMessage) -> None:
+    print(f"{source}: {str(message)}")
+
+
+def send_queue_batch(
+    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT
+) -> None:
+    with client.get_queue_sender(queue_name=queue_name) as sender:
+        batch = sender.create_message_batch()
+        for index in range(1, message_count + 1):
+            try:
+                batch.add_message(
+                    ServiceB
... (truncated)
```

### ✅ 12. apply_patch — 760ms

**Input:**

```json
"*** Begin Patch\n*** Update File: sync_servicebus.py\n@@\n-from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus import (\n+    ServiceBusClient,\n+    ServiceBusMessage,\n+    ServiceBusReceivedMessage,\n+)\n@@\n-def process_message(source: str, message: ServiceBusMessage) -\u003e None:\n+def process_message(source: str, message: ServiceBusReceivedMessage) -\u003e None:\n*** Update File: async_servicebus.py\n@@\n from azure.identity.aio import DefaultAzureCredential\n-from azure.servicebus import ServiceBusMessage\n-from azure.servicebus.aio import ServiceBusClient\n-from azure.servicebus.aio._base_handler_async import ServiceBusReceiver\n+from azure.servicebus import ServiceBusMessage, ServiceBusReceivedMessage\n+from azure.servicebus.aio import ServiceBusClient, ServiceBusReceiver\n@@\n-async def process_message(source: str, message: ServiceBusMessage) -\u003e None:\n+async def process_message(source: str, message: ServiceBusReceivedMessage) -\u003e None:\n@@\n async def process_and_complete(\n-    receiver: ServiceBusReceiver, source: str, message: ServiceBusMessage\n+    receiver: ServiceBusReceiver, source: str, message: ServiceBusReceivedMessage\n ) -\u003e None:\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/sync_servicebus.py
@@ -1,7 +1,11 @@
 import os
 
 from azure.identity import DefaultAzureCredential
-from azure.servicebus import ServiceBusClient, ServiceBusMessage
+from azure.servicebus import (
+    ServiceBusClient,
+    ServiceBusMessage,
+    ServiceBusReceivedMessage,
+)
 
 
 MESSAGE_COUNT = 5
@@ -15,7 +19,7 @@
     return value
 
 
-def process_message(source: str, message: ServiceBusMessage) -> None:
+def process_message(source: str, message: ServiceBusReceivedMessage) -> None:
     print(f"{source}: {str(message)}")
 
 


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/async_servicebus.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/async_servicebus.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/async_servicebus.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-2471172224/async_servicebus.py
@@ -2,9 +2,8 @@
 import os
 
 from azure.identity.aio import DefaultAzureCredential
-from azure.servicebus import ServiceBusMessage
-from azure.servicebus.aio im
... (truncated)
```

### ✅ 13. powershell — 5070ms

**Input:**

```json
{"command":"python -m pip install --quiet -r .\\requirements.txt \u0026\u0026 python -m compileall -q .\\sync_servicebus.py .\\async_servicebus.py \u0026\u0026 python -c \"import sync_servicebus, async_servicebus; print('imports-ok')\"","description":"Install dependencies and validate Python examples","initial_wait":180,"mode":"sync"}
```

**Output:**

```
imports-ok
<shellId: 1 completed with exit code 0>
```

### ✅ 14. powershell — 1180ms

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -LiteralPath .\\__pycache__ -Recurse -Force }","description":"Remove Python cache files"}
```

**Output:**

```

<shellId: 2 completed with exit code 0>
```

## Generated Files

- `README.md`
- `async_servicebus.py`
- `requirements.txt`
- `sync_servicebus.py`

## Copilot Response

I’m checking the SDK’s required patterns and the current workspace layout before creating the runnable examples.

The examples will use Entra authentication with `DefaultAzureCredential`, separate sync/async modules, bounded receive waits, and deterministic cleanup for credentials, clients, senders, and receivers.

I’ll now add pinned dependencies, separate synchronous and `aio` demonstrations, and concise setup instructions without provisioning or modifying any Azure resources.

The implementation is complete; I’m validating the pinned SDK imports and both scripts locally without connecting to Service Bus.

Created:

- `sync_servicebus.py` — sends a five-message batch, receives queue messages, completes them, and demonstrates topic/subscription messaging.
- `async_servicebus.py` — equivalent `aio` workflow with prefetching and concurrent processing.
- `requirements.txt` — pinned `azure-servicebus` and `azure-identity` packages.
- `README.md` — environment variables, authentication/RBAC guidance, installation, and run commands.

All clients, credentials, senders, and receivers use appropriate sync or async context managers.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- The code handles batch-size ValueError and consistently uses proper synchronous and asynchronous context managers.

### Issues

- Missing handling for HttpResponseError and applicable Azure SDK exception subclasses in both synchronous and asynchronous workflows.

## Grader Results

- send-receive-messages.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - `azure-servicebus` pip package: Pass
      - `ServiceBusSender` via `get_queue_sender()` or `get_topic_sender()`: Pass
      - `ServiceBusMessage` and `ServiceBusMessageBatch`: Pass
      - `ServiceBusReceiver` via `get_queue_receiver()` or `get_subscription_receiver()`: Pass
      - `complete_message()`, `abandon_message()`, `dead_letter_message()`: Fail
      - Context manager pattern (`with` statements) for resource cleanup: Pass
      - Async variants in `azure.servicebus.aio`: Pass
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 14.6% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.8571** | **73.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-crud --config python-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
