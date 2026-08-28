# Evaluation Report: service-bus-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 555.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 11/14 |
| Duration | 555.5s |
| Timestamp | 2026-08-27T22:13:40Z |
| Files Generated | 4 |
| Event Count | 4654 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 462.1s |
| Review | 92.1s |
| **Total** | **555.5s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-servicebus-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=243345 out=3977 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3475ms

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

### ✅ 2. glob — 185989ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-servicebus-py\SKILL.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-servicebus-py\scripts\setup_servicebus.py
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-servicebus-py\references\patterns.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-python\skills\azure-servicebus-py\references\dead-letter.md
```

### ✅ 3. glob — 183566ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 184639ms

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

### ❌ 5. azure-get_azure_bestpractices (via azure) — 185988ms

**Input:**

```json
{"intent":"Generate local Python Azure Service Bus queue batch send/receive and topic/subscription examples using sync and aio clients, DefaultAzureCredential, and context managers.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 182247ms

**Input:**

```json
{"intent":"Get Python code-generation best practices for local Azure Service Bus examples using DefaultAzureCredential, synchronous and aio clients, message batches, queue settlement, and topic subscriptions.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. apply_patch — 3186ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\n+azure-servicebus\n*** Add File: .env.example\n+SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=your-namespace.servicebus.windows.net\n+SERVICEBUS_QUEUE_NAME=your-queue\n+SERVICEBUS_TOPIC_NAME=your-topic\n+SERVICEBUS_SUBSCRIPTION_NAME=your-subscription\n+# In production, constrain DefaultAzureCredential to production-safe credentials.\n+AZURE_TOKEN_CREDENTIALS=prod\n*** Add File: servicebus_sync.py\n+\"\"\"Synchronous Azure Service Bus queue and topic/subscription examples.\"\"\"\n+\n+import os\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.servicebus import (\n+    ServiceBusClient,\n+    ServiceBusMessage,\n+    ServiceBusMessageBatch,\n+)\n+\n+\n+def process_message(source: str, message: ServiceBusMessage) -\u003e None:\n+    \"\"\"Replace this function with application-specific processing.\"\"\"\n+    print(f\"{source}: {message}\")\n+\n+\n+def send_queue_batch(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    with client.get_queue_sender(queue_name=queue_name) as sender:\n+        batch: ServiceBusMessageBatch = sender.create_message_batch()\n+        for index in range(1, 6):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Synchronous queue message {index}\",\n+                    message_id=f\"sync-queue-{index}\",\n+                )\n+            )\n+        sender.send_messages(batch)\n+        print(f\"Sent {len(batch)} messages to queue {queue_name!r}.\")\n+\n+\n+def receive_queue_messages(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    with client.get_queue_receiver(\n+        queue_name=queue_name,\n+        max_wait_time=5,\n+    ) as receiver:\n+        messages = receiver.receive_messages(\n+            max_message_count=5,\n+            max_wait_time=5,\n+        )\n+        for message in messages:\n+            process_message(\"queue\", message)\n+            receiver.complete_message(message)\n+\n+\n+def topic_subscription_roundtrip(\n+    client: ServiceBusClient,\n+    topic_name: str,\n+    subscription_name: str,\n+) -\u003e None:\n+    with client.get_topic_sender(topic_name=topic_name) as sender:\n+        sender.send_messages(\n+            ServiceBusMessage(\n+                \"Synchronous topic message\",\n+                message_id=\"sync-topic-1\",\n+            )\n+        )\n+\n+    with client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+        max_wait_time=5,\n+    ) as receiver:\n+        messages = receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=5,\n+        )\n+        for message in messages:\n+            process_message(\"topic subscription\", message)\n+            receiver.complete_message(message)\n+\n+\n+def main() -\u003e None:\n+    namespace = os.environ[\"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\"]\n+    queue_name = os.environ[\"SERVICEBUS_QUEUE_NAME\"]\n+    topic_name = os.environ[\"SERVICEBUS_TOPIC_NAME\"]\n+    subscription_name = os.environ[\"SERVICEBUS_SUBSCRIPTION_NAME\"]\n+\n+    with DefaultAzureCredential() as credential:\n+        with ServiceBusClient(\n+            fully_qualified_namespace=namespace,\n+            credential=credential,\n+        ) as client:\n+            send_queue_batch(client, queue_name)\n+            receive_queue_messages(client, queue_name)\n+            topic_subscription_roundtrip(client, topic_name, subscription_name)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: servicebus_async.py\n+\"\"\"Asynchronous Azure Service Bus examples for higher-throughput processing.\"\"\"\n+\n+import asyncio\n+import os\n+\n+from azure.identity.aio import DefaultAzureCredential\n+from azure.servicebus import ServiceBusMessage, ServiceBusMessageBatch\n+from azure.servicebus.aio import ServiceBusClient, ServiceBusReceiver\n+\n+\n+async def process_message(source: str, message: ServiceBusMessage) -\u003e None:\n+    \"\"\"Replace this function with application-specific asynchronous work.\"\"\"\n+    await asyncio.sleep(0)\n+    print(f\"{source}: {message}\")\n+\n+\n+async def process_and_complete(\n+    receiver: ServiceBusReceiver,\n+    source: str,\n+    message: ServiceBusMessage,\n+) -\u003e None:\n+    await process_message(source, message)\n+    await receiver.complete_message(message)\n+\n+\n+async def send_queue_batch(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    async with client.get_queue_sender(queue_name=queue_name) as sender:\n+        batch: ServiceBusMessageBatch = await sender.create_message_batch()\n+        for index in range(1, 6):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Asynchronous queue message {index}\",\n+                    message_id=f\"async-queue-{index}\",\n+                )\n+            )\n+        await sender.send_messages(batch)\n+        print(f\"Sent {len(batch)} messages to queue {queue_name!r}.\")\n+\n+\n+async def receive_queue_messages(\n+    client: ServiceBusClient,\n+    queue_name: str,\n+) -\u003e None:\n+    async with client.get_queue_receiver(\n+        queue_name=queue_name,\n+        max_wait_time=5,\n+        prefetch_count=10,\n+    ) as receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=5,\n+            max_wait_time=5,\n+        )\n+        await asyncio.gather(\n+            *(\n+                process_and_complete(receiver, \"queue\", message)\n+                for message in messages\n+            )\n+        )\n+\n+\n+async def topic_subscription_roundtrip(\n+    client: ServiceBusClient,\n+    topic_name: str,\n+    subscription_name: str,\n+) -\u003e None:\n+    async with client.get_topic_sender(topic_name=topic_name) as sender:\n+        await sender.send_messages(\n+            ServiceBusMessage(\n+                \"Asynchronous topic message\",\n+                message_id=\"async-topic-1\",\n+            )\n+        )\n+\n+    async with client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+        max_wait_time=5,\n+    ) as receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=5,\n+        )\n+        await asyncio.gather(\n+            *(\n+                process_and_complete(\n+                    receiver,\n+                    \"topic subscription\",\n+                    message,\n+                )\n+                for message in messages\n+            )\n+        )\n+\n+\n+async def main() -\u003e None:\n+    namespace = os.environ[\"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\"]\n+    queue_name = os.environ[\"SERVICEBUS_QUEUE_NAME\"]\n+    topic_name = os.environ[\"SERVICEBUS_TOPIC_NAME\"]\n+    subscription_name = os.environ[\"SERVICEBUS_SUBSCRIPTION_NAME\"]\n+\n+    async with DefaultAzureCredential() as credential:\n+        async with ServiceBusClient(\n+            fully_qualified_namespace=namespace,\n+            credential=credential,\n+        ) as client:\n+            await send_queue_batch(client, queue_name)\n+            await receive_queue_messages(client, queue_name)\n+            await topic_subscription_roundtrip(\n+                client,\n+                topic_name,\n+                subscription_name,\n+            )\n+\n+\n+if __name__ == \"__main__\":\n+    asyncio.run(main())\n*** Add File: README.md\n+# Azure Service Bus Python examples\n+\n+These scripts use Microsoft Entra authentication through\n+`DefaultAzureCredential`; no connection string is stored in source.\n+\n+## Install\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+Required packages:\n+\n+- `azure-servicebus`\n+- `azure-identity`\n+\n+## Configure\n+\n+Create the queue, topic, and subscription before running the examples. Set the\n+values shown in `.env.example` in the current shell:\n+\n+```powershell\n+$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = \"your-namespace.servicebus.windows.net\"\n+$env:SERVICEBUS_QUEUE_NAME = \"your-queue\"\n+$env:SERVICEBUS_TOPIC_NAME = \"your-topic\"\n+$env:SERVICEBUS_SUBSCRIPTION_NAME = \"your-subscription\"\n+```\n+\n+Authenticate locally with a credential supported by `DefaultAzureCredential`,\n+such as Azure CLI or Visual Studio Code. In production, use managed identity\n+and set `AZURE_TOKEN_CREDENTIALS=prod` to constrain the credential chain.\n+\n+## Run\n+\n+Run the synchronous context-manager example:\n+\n+```powershell\n+python .\\servicebus_sync.py\n+```\n+\n+Run the asynchronous `aio` example:\n+\n+```powershell\n+python .\\servicebus_async.py\n+```\n+\n+Both examples send a five-message `ServiceBusMessageBatch` to a queue, receive\n+and complete the messages after processing, publish to a topic, and receive\n+and complete a message from its subscription. The async example also uses\n+prefetching and concurrent processing for higher throughput.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/requirements.txt
@@ -1,0 +1,3 @@
+azure-identity
+azure-servicebus
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/.env.example b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/.env.example
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/.env.example
@@ -1,0 +1,7 @@
+SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=your-namespace.servicebus.windows.net
+SERVICEBUS_QUEUE_NAME=your-queue
+SERVICEBUS_TOPIC_NAME=your-topic
+SERVICEBUS_SUBSCRIPTION_NAME=your-subscription
+# In production, constrain DefaultAzureCredential to production-safe credentials.
+AZURE_TOKEN_CREDENTIALS=prod
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/servicebus_sync.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/servicebus_sync.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053/servicebus_sync.py
@@ -1,0 +1,91 @@
+"""Synchronous Azure Serv
... (truncated)
```

### ✅ 8. powershell — 5040ms

**Input:**

```json
{"command":"python -m py_compile .\\servicebus_sync.py .\\servicebus_async.py","description":"Compile both Python examples","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 9. rg — 4551ms

**Input:**

```json
{"-n":true,"glob":"*.py","output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053","pattern":"ServiceBusMessageBatch|range\\(1, 6\\)|complete_message|get_queue_receiver|get_topic_sender|get_subscription_receiver|async with|with ServiceBusClient"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:7:from azure.servicebus import ServiceBusMessage, ServiceBusMessageBatch
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:23:    await receiver.complete_message(message)
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:27:    async with client.get_queue_sender(queue_name=queue_name) as sender:
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:28:        batch: ServiceBusMessageBatch = await sender.create_message_batch()
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:29:        for index in range(1, 6):
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:44:    async with client.get_queue_receiver(
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:66:    async with client.get_topic_sender(topic_name=topic_name) as sender:
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:74:    async with client.get_subscription_receiver(
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-microsoft-skill-1801022053\servicebus_async.py:101:    async with DefaultAzureCredential() as credential:
C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-crud-python-azure-skills
... (truncated)
```

## Generated Files

- `README.md`
- `requirements.txt`
- `servicebus_async.py`
- `servicebus_sync.py`

## Copilot Response

I’ll inspect the SDK guidance and workspace, then create separate synchronous and asynchronous examples with dependency metadata.

The workspace is empty. The SDK skill has no separate acceptance-criteria/examples files, so I’ll use its documented patterns and retry the Azure guidance lookup before writing the examples.

I’ll now add two executable examples plus dependency and environment setup files, keeping sync and `aio` call paths separate.

The examples are in place; I’ll run local syntax checks and inspect the generated files for the requested batch, settlement, async, and topic/subscription patterns.

Created the Azure Service Bus examples:

- `servicebus_sync.py` — batch send, queue receive/complete, topic/subscription messaging
- `servicebus_async.py` — `aio`, prefetching, concurrent processing, async context managers
- `requirements.txt` — `azure-servicebus` and `azure-identity`
- `.env.example` — required configuration
- `README.md` — installation, authentication, and run instructions

Both scripts use `DefaultAzureCredential` and proper context-manager cleanup.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed. The messaging examples implement the requested operations but do not meet the exception-handling criterion.

### Strengths

- Uses DefaultAzureCredential, ServiceBusMessageBatch, receiver.complete_message(), and proper synchronous and asynchronous context managers.

### Issues

- Missing handling for HttpResponseError and relevant subclasses in both synchronous and asynchronous workflows.

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
hyoka run --prompt-id service-bus-dp-python-crud --config python-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
