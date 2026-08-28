# Evaluation Report: service-bus-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 194.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/14 |
| Duration | 194.3s |
| Timestamp | 2026-08-28T21:33:41Z |
| Files Generated | 3 |
| Event Count | 4923 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 97.9s |
| Review | 95.3s |
| **Total** | **194.3s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=288971 out=3856 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 1914ms

**Input:**

```json
{"skill":"azure-messaging"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Messaging SDK Troubleshooting

## Quick Reference

| Property | Value |
|----------|-------|
| **Services** | Azure Event Hubs, Azure Service Bus |
| **MCP Tools** | `mcp_azure_mcp_eventhubs`, `mcp_azure_mcp_servicebus` |
| **Best For** | Diagnosing SDK connection, auth, and message processing issues |

## When to Use This Skill

- SDK connection failures, auth errors, or AMQP link errors
- Idle timeout, connection inactivity, or slow reconnection after disconnect
- AMQP link detach or detach-forced errors
- Message lock lost, message lock expired, lock renewal failures, or batch lock timeouts
- Session lock lost, session lock expired, or session receiver errors
- Event processor or message handler stops processing
- Duplicate events or checkpoint offset resets
- SDK configuration questions (retry, prefetch, batch size, receive batch behavior)

## MCP Tools

| Tool | Command | Use |
|------|---------|-----|
| `mcp_azure_mcp_eventhubs` | Namespace/hub ops | List namespaces, hubs, consumer groups |
| `mcp_azure_mcp_servicebus` | Queue/topic ops | List namespaces, queues, topics, subscriptions |
| `mcp_azure_mcp_monitor` | `logs_query` | Query diagnostic logs with KQL |
| `mcp_azure_mcp_resourcehealth` | `get` | Check service health status |
| `mcp_azure_mcp_documentation` | Doc search | Search Microsoft Learn for troubleshooting docs |

## Diagnosis Workflow

1. **Identify the SDK and version** — Check the prompt for SDK and version clues; if not stated, proceed with diagnosis and ask later if needed
2. **Check resource health** — Use `mcp_azure_mcp_resourcehealth` to verify the namespace is healthy
3. **Review the error message** — Match against language-specific troubleshooting guide
4. **Look up documentation** — Use `mcp_azure_mcp_documentation` to search Microsoft Learn for the error or topic
5. **Check configuration** — Verify connection string, entity name, consumer group
6. **Recommend fix** — Apply remediatio
... (truncated)
```

### ✅ 2. glob — 9423ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. azure-get_azure_bestpractices (via azure) — 10657ms

**Input:**

```json
{"intent":"Generate local-only Python examples for Azure Service Bus using synchronous and aio clients, queue batch send/receive/complete, and topic/subscription send/receive, with environment-variable connection strings and proper context managers.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 4. azure-documentation (via azure) — 12023ms

**Input:**

```json
{"intent":"Find current official Azure Service Bus Python SDK guidance for ServiceBusMessageBatch, receive_messages, complete_message, aio clients, and topic subscription messaging.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 5. glob — 8372ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 6071ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Python code-generation practices for a local-only Azure Service Bus sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 7222ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Azure Service Bus Python SDK guidance.","parameters":{"query":"Azure Service Bus Python azure-servicebus ServiceBusMessageBatch receive_messages complete_message aio topic subscription context manager"}}
```

**Output:**

```
{"results":[{"title":"Send messages to an Azure Service Bus topic and receive messages from subscriptions to the topic (Python)","content":"# Send messages to an Azure Service Bus topic and receive messages from subscriptions to the topic (Python)\n1. [C#](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-dotnet-how-to-use-topics-subscriptions)\n2. [Java](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-java-how-to-use-topics-subscriptions)\n3. [JavaScript](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-nodejs-how-to-use-topics-subscriptions)\n4. [Python](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-topics-subscriptions)\nIn this tutorial, you complete the following steps:\n1. Create a Service Bus namespace, using the Azure portal.\n2. Create a Service Bus topic, using the Azure portal.\n3. Create a Service Bus subscription to that topic, using the Azure portal.\n4. Write a Python application to use the [azure-servicebus](https://pypi.org/project/azure-servicebus/)package to:\n4.1. Send a set of messages to the topic.\n4.2. Receive those messages from the subscription.\nNote\nThis quickstart provides step-by-step instructions for a simple scenario of sending a batch of messages to a Service Bus topic and receiving those messages from a subscription of the topic. You can find pre-built Python samples for Azure Service Bus in the [Azure SDK for Python repository on GitHub](https://github.com/azure/azure-sdk-for-python/tree/main/sdk/servicebus/azure-servicebus/samples).\n## Prerequisites\n1. An [Azure subscription](https://azure.microsoft.com/pricing/purchase-options/azure-account?cid=msft_learn).\n2. [Python 3.8](https://www.python.org/downloads/) or higher\nNote\nThis tutorial works with samples that you can copy and run using Python. For instructions on how to create a Python application, see [Create and deploy a Python application to an Azure Website](https://learn.mi
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 6428ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Azure Service Bus Python code samples.","parameters":{"language":"python","query":"Azure Service Bus Python ServiceBusMessageBatch receive_messages complete_message aio topic subscription"}}
```

**Output:**

```
{"results":[{"description":"description: Defines the signature of the receive_messages method which receives a batch of messages from Azure Service Bus, with optional parameters for maximum message count and maximum wait time.\nlanguage: python\n","codeSnippet":"\r\n   with servicebus_receiver:\r\n       messages_sync = servicebus_receiver.receive_messages(max_wait_time=5)\r\n       for message in messages_sync:\r\n           print(str(message))\r\n           servicebus_receiver.complete_message(message)","link":"https://learn.microsoft.com/python/api/azure-servicebus/azure.servicebus.servicebusreceiver?view=azure-python","language":"python"},{"description":"description: Defines the complete_message method which completes a received message, removing it from the Azure Service Bus queue.\npackage: azure.servicebus\nlanguage: python\n","codeSnippet":"\r\n       messages_complete = servicebus_receiver.receive_messages(max_wait_time=5)\r\n       for message in messages_complete:\r\n           servicebus_receiver.complete_message(message)","link":"https://learn.microsoft.com/python/api/azure-servicebus/azure.servicebus.servicebusreceiver?view=azure-python","language":"python"},{"description":"description: Sends messages to a topic and receives messages from a subscription in Azure Service Bus using ServiceBusClient and authenticates with DefaultAzureCredential.\npackage: azure-servicebus, azure-identity\nlanguage: python\n","codeSnippet":"from azure.servicebus import ServiceBusClient, ServiceBusMessage\r\nfrom azure.identity import DefaultAzureCredential\r\n\r\nimport os\r\nfully_qualified_namespace = os.environ[\u0027SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\u0027]\r\ntopic_name = os.environ[\u0027SERVICE_BUS_TOPIC_NAME\u0027]\r\nsubscription_name = os.environ[\u0027SERVICE_BUS_SUBSCRIPTION_NAME\u0027]\r\n\r\ncredential = DefaultAzureCredential()\r\nwith ServiceBusClient(fully_qualified_namespace, credential) as client:\r\n    with client.get_topic_sender(topic_name) as send
... (truncated)
```

### ✅ 9. view — 6834ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
@@ -1,0 +1,1 @@
 SKILL.md

```

### ✅ 10. apply_patch — 865ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\u003e=1.25.1,\u003c2.0.0\n+azure-servicebus\u003e=7.14.2,\u003c8.0.0\n*** Add File: service_bus_demo.py\n+\"\"\"Azure Service Bus queue and topic/subscription messaging examples.\"\"\"\n+\n+import asyncio\n+import logging\n+import os\n+from dataclasses import dataclass\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+\n+MESSAGE_COUNT = 5\n+RECEIVE_WAIT_SECONDS = 10\n+\n+\n+@dataclass(frozen=True)\n+class Settings:\n+    fully_qualified_namespace: str\n+    queue_name: str\n+    topic_name: str\n+    subscription_name: str\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"Settings\":\n+        return cls(\n+            fully_qualified_namespace=os.environ[\n+                \"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\"\n+            ],\n+            queue_name=os.environ[\"SERVICEBUS_QUEUE_NAME\"],\n+            topic_name=os.environ[\"SERVICEBUS_TOPIC_NAME\"],\n+            subscription_name=os.environ[\"SERVICEBUS_SUBSCRIPTION_NAME\"],\n+        )\n+\n+\n+def process_message(source: str, message: object) -\u003e None:\n+    \"\"\"Replace this function with application-specific processing.\"\"\"\n+    logging.info(\"Processed message from %s: %s\", source, message)\n+\n+\n+def run_sync_queue_demo(settings: Settings) -\u003e None:\n+    with DefaultAzureCredential() as credential:\n+        with ServiceBusClient(\n+            fully_qualified_namespace=settings.fully_qualified_namespace,\n+            credential=credential,\n+        ) as client:\n+            with client.get_queue_sender(settings.queue_name) as sender:\n+                batch = sender.create_message_batch()\n+                for index in range(1, MESSAGE_COUNT + 1):\n+                    batch.add_message(\n+                        ServiceBusMessage(\n+                            f\"Synchronous queue message {index}\",\n+                            message_id=f\"sync-queue-{index}\",\n+                        )\n+                    )\n+                sender.send_messages(batch)\n+                logging.info(\"Sent a batch of %d queue messages\", MESSAGE_COUNT)\n+\n+            with client.get_queue_receiver(\n+                queue_name=settings.queue_name,\n+                max_wait_time=RECEIVE_WAIT_SECONDS,\n+            ) as receiver:\n+                messages = receiver.receive_messages(\n+                    max_message_count=MESSAGE_COUNT,\n+                    max_wait_time=RECEIVE_WAIT_SECONDS,\n+                )\n+                for message in messages:\n+                    process_message(settings.queue_name, message)\n+                    receiver.complete_message(message)\n+                logging.info(\"Completed %d queue messages\", len(messages))\n+\n+\n+def run_sync_topic_demo(settings: Settings) -\u003e None:\n+    with DefaultAzureCredential() as credential:\n+        with ServiceBusClient(\n+            fully_qualified_namespace=settings.fully_qualified_namespace,\n+            credential=credential,\n+        ) as client:\n+            with client.get_topic_sender(settings.topic_name) as sender:\n+                sender.send_messages(\n+                    ServiceBusMessage(\n+                        \"Synchronous topic message\",\n+                        message_id=\"sync-topic-1\",\n+                    )\n+                )\n+                logging.info(\"Sent a message to topic %s\", settings.topic_name)\n+\n+            with client.get_subscription_receiver(\n+                topic_name=settings.topic_name,\n+                subscription_name=settings.subscription_name,\n+                max_wait_time=RECEIVE_WAIT_SECONDS,\n+            ) as receiver:\n+                messages = receiver.receive_messages(\n+                    max_message_count=1,\n+                    max_wait_time=RECEIVE_WAIT_SECONDS,\n+                )\n+                for message in messages:\n+                    process_message(settings.subscription_name, message)\n+                    receiver.complete_message(message)\n+                logging.info(\n+                    \"Completed %d message(s) from subscription %s\",\n+                    len(messages),\n+                    settings.subscription_name,\n+                )\n+\n+\n+async def send_async_queue_batch(\n+    client: AsyncServiceBusClient, settings: Settings\n+) -\u003e None:\n+    async with client.get_queue_sender(settings.queue_name) as sender:\n+        batch = await sender.create_message_batch()\n+        for index in range(1, MESSAGE_COUNT + 1):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Asynchronous queue message {index}\",\n+                    message_id=f\"async-queue-{index}\",\n+                )\n+            )\n+        await sender.send_messages(batch)\n+        logging.info(\"Asynchronously sent a batch of %d messages\", MESSAGE_COUNT)\n+\n+\n+async def receive_async_queue(\n+    client: AsyncServiceBusClient, settings: Settings\n+) -\u003e None:\n+    async with client.get_queue_receiver(\n+        queue_name=settings.queue_name,\n+        max_wait_time=RECEIVE_WAIT_SECONDS,\n+    ) as receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=MESSAGE_COUNT,\n+            max_wait_time=RECEIVE_WAIT_SECONDS,\n+        )\n+        for message in messages:\n+            process_message(settings.queue_name, message)\n+            await receiver.complete_message(message)\n+        logging.info(\"Asynchronously completed %d queue messages\", len(messages))\n+\n+\n+async def send_async_topic_message(\n+    client: AsyncServiceBusClient, settings: Settings\n+) -\u003e None:\n+    async with client.get_topic_sender(settings.topic_name) as sender:\n+        await sender.send_messages(\n+            ServiceBusMessage(\n+                \"Asynchronous topic message\",\n+                message_id=\"async-topic-1\",\n+            )\n+        )\n+        logging.info(\"Asynchronously sent a topic message\")\n+\n+\n+async def receive_async_subscription(\n+    client: AsyncServiceBusClient, settings: Settings\n+) -\u003e None:\n+    async with client.get_subscription_receiver(\n+        topic_name=settings.topic_name,\n+        subscription_name=settings.subscription_name,\n+        max_wait_time=RECEIVE_WAIT_SECONDS,\n+    ) as receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=RECEIVE_WAIT_SECONDS,\n+        )\n+        for message in messages:\n+            process_message(settings.subscription_name, message)\n+            await receiver.complete_message(message)\n+        logging.info(\n+            \"Asynchronously completed %d subscription message(s)\", len(messages)\n+        )\n+\n+\n+async def run_async_demo(settings: Settings) -\u003e None:\n+    async with AsyncDefaultAzureCredential() as credential:\n+        async with AsyncServiceBusClient(\n+            fully_qualified_namespace=settings.fully_qualified_namespace,\n+            credential=credential,\n+        ) as client:\n+            # Independent sends and receives overlap instead of blocking one another.\n+            await asyncio.gather(\n+                send_async_queue_batch(client, settings),\n+                send_async_topic_message(client, settings),\n+            )\n+            await asyncio.gather(\n+                receive_async_queue(client, settings),\n+                receive_async_subscription(client, settings),\n+            )\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(level=logging.INFO, format=\"%(levelname)s %(message)s\")\n+    settings = Settings.from_environment()\n+\n+    run_sync_queue_demo(settings)\n+    run_sync_topic_demo(settings)\n+    asyncio.run(run_async_demo(settings))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: README.md\n+# Azure Service Bus Python messaging demo\n+\n+This sample demonstrates synchronous and asynchronous queue messaging, batch\n+sends, explicit message completion, and topic/subscription messaging. The queue,\n+topic, and subscription must already exist; the sample does not create or modify\n+Azure resources.\n+\n+## Install\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+Required packages:\n+\n+- `azure-servicebus`\n+- `azure-identity`\n+\n+## Configure\n+\n+Authenticate locally with an identity supported by `DefaultAzureCredential`,\n+then set the existing Service Bus namespace and entity names:\n+\n+```powershell\n+$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = \"your-namespace.servicebus.windows.net\"\n+$env:SERVICEBUS_QUEUE_NAME = \"your-queue\"\n+$env:SERVICEBUS_TOPIC_NAME = \"your-topic\"\n+$env:SERVICEBUS_SUBSCRIPTION_NAME = \"your-subscription\"\n+```\n+\n+The identity needs the **Azure Service Bus Data Sender** and **Azure Service Bus\n+Data Receiver** roles scoped as narrowly as practical.\n+\n+## Run\n+\n+```powershell\n+python service_bus_demo.py\n+```\n+\n+The synchronous examples use `with` for credentials, clients, senders, and\n+receivers. The higher-throughput example uses `azure.servicebus.aio`, `async\n+with`, and `asyncio.gather` to overlap independent queue and topic operations.\n+\n+## References\n+\n+- [Azure Service Bus client library for Python](https://learn.microsoft.com/python/api/overview/azure/servicebus-readme)\n+- [Send to topics and receive from subscriptions with Python](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-topics-subscriptions)\n+- [Passwordless authentication for Azure Service Bus](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-passwordless-messaging)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767/requirements.txt
@@ -1,0 +1,3 @@
+azure-identity>=1.25.1,<2.0.0
+azure-servicebus>=7.14.2,<8.0.0
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767/service_bus_demo.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767/service_bus_demo.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-715277767/service_bus_demo.py
@@ -1,0 +1,202 @@
+"""Azure Service Bus queue and topic/subscription messaging examples."""
+
+import asyncio
+import logging
+import os
+from dataclasses import dataclass
+
+from azure.identity import DefaultAzureCredential
+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
+from azure.servicebus import ServiceBusClient, ServiceBusMessage
+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
+
+
+MESSAGE_COUNT = 5
+RECEIVE_WAIT_SECONDS = 10
+
+
+@dataclass(frozen=True)
+class Settings:
+    fully_qualified_namespace: str
+    queue_name: str
+    topic_name: str
+    subscription_name: str
+
+    @classmethod
+    def from_environment(cls) -> "Settings":
+        return cls(
+            fully_qualified_namespace=os.environ[
+                "SERVICEBUS_FULLY_QUALIFIED_NAMESPACE"
+            ],
+            queue_name=os.environ["SERVICEBUS_QUEUE_NAME"],
+            topic_name=os.
... (truncated)
```

### ✅ 11. powershell — 1323ms

**Input:**

```json
{"command":"python -m py_compile .\\service_bus_demo.py","description":"Compile-check the Service Bus demo","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 12. powershell — 1321ms

**Input:**

```json
{"command":"$pyc = Get-ChildItem -Path .\\__pycache__ -Filter 'service_bus_demo.*.pyc' -File; $pyc | ForEach-Object { Remove-Item -LiteralPath $_.FullName }; if ((Get-ChildItem -Path .\\__pycache__ -Force | Measure-Object).Count -eq 0) { Remove-Item -LiteralPath .\\__pycache__ }","description":"Remove Python compile artifacts","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `requirements.txt`
- `service_bus_demo.py`

## Copilot Response

I’ll ground the sample in the current Service Bus SDK patterns, then create and locally validate the scripts without contacting Azure.

The current guidance favors passwordless `DefaultAzureCredential`; I’ll use environment-based entity names and namespace, with no embedded keys or live resource operations.

Created:

- `service_bus_demo.py` — synchronous and `aio` examples for queue batch messaging, explicit completion, and topic/subscription messaging.
- `requirements.txt` — `azure-servicebus` and `azure-identity`.
- `README.md` — installation, environment configuration, RBAC requirements, execution steps, and Microsoft references.

The script uses passwordless `DefaultAzureCredential`, environment variables, `with`/`async with` context managers, `ServiceBusMessageBatch` with five messages, and concurrent async operations via `asyncio.gather`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed. The messaging flows are implemented, but the required Azure SDK exception handling is absent.

### Strengths

- Uses proper synchronous and asynchronous context managers, batch messaging, explicit completion, and topic/subscription flows.

### Issues

- No handling for HttpResponseError or relevant subclasses around Azure Service Bus operations.

## Grader Results

- send-receive-messages.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/7)
      - `azure-servicebus` pip package: Pass
      - `ServiceBusSender` via `get_queue_sender()` or `get_topic_sender()`: Pass
      - `ServiceBusMessage` and `ServiceBusMessageBatch`: Fail
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
| `Criteria from prompt file` | prompt_review | 71% | 1.00 | 0.7143 | 12.5% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.7143** | **71.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-crud --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
