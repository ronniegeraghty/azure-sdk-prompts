# Evaluation Report: service-bus-dp-python-crud

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 365.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-crud` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/14 |
| Duration | 365.8s |
| Timestamp | 2026-08-27T22:07:30Z |
| Files Generated | 4 |
| Event Count | 5991 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 283.1s |
| Review | 81.5s |
| **Total** | **365.8s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=259543 out=4023 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 3352ms

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

### ✅ 2. glob — 7940ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 6675ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/examples.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. glob — 5636ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 184504ms

**Input:**

```json
{"intent":"Generate local-only Python examples for Azure Service Bus queue batch send/receive/complete, async aio usage, and topic/subscription messaging using context managers.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 184507ms

**Input:**

```json
{"intent":"Find current official Python Azure Service Bus examples for ServiceBusMessageBatch, complete_message, aio clients, and topic/subscription context managers.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. apply_patch — 2676ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-servicebus\u003e=7.12,\u003c8\n*** Add File: README.md\n+# Azure Service Bus Python messaging examples\n+\n+These examples demonstrate:\n+\n+- Sending exactly five queue messages in a `ServiceBusMessageBatch`\n+- Receiving and completing queue messages after processing\n+- Sending to a topic and receiving from a subscription\n+- Synchronous and `azure.servicebus.aio` asynchronous context-manager patterns\n+\n+## Install\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+The only required package is `azure-servicebus`. Its AMQP transport dependencies\n+are installed automatically.\n+\n+## Configure\n+\n+Set these environment variables before connecting:\n+\n+```powershell\n+$env:SERVICEBUS_CONNECTION_STR = \"\u003cnamespace connection string\u003e\"\n+$env:SERVICEBUS_QUEUE_NAME = \"\u003cqueue name\u003e\"\n+$env:SERVICEBUS_TOPIC_NAME = \"\u003ctopic name\u003e\"\n+$env:SERVICEBUS_SUBSCRIPTION_NAME = \"\u003csubscription name\u003e\"\n+```\n+\n+Use a connection string whose shared access policy has permission to send and\n+receive. Do not commit connection strings.\n+\n+## Run\n+\n+Both scripts default to a local dry run, which does not connect to Azure:\n+\n+```powershell\n+python sync_service_bus.py\n+python async_service_bus.py\n+```\n+\n+After configuring an Azure Service Bus namespace and entities, opt in to network\n+operations:\n+\n+```powershell\n+python sync_service_bus.py --execute\n+python async_service_bus.py --execute\n+```\n+\n+Each executable run sends a fresh batch before receiving, so it does not depend\n+on messages left by a previous run. The topic examples assume the subscription\n+already exists; subscriptions receive only messages sent after they are created.\n*** Add File: sync_service_bus.py\n+\"\"\"Synchronous Azure Service Bus queue and topic/subscription example.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import os\n+from dataclasses import dataclass\n+\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+\n+BATCH_SIZE = 5\n+\n+\n+@dataclass(frozen=True)\n+class ServiceBusSettings:\n+    connection_string: str\n+    queue_name: str\n+    topic_name: str\n+    subscription_name: str\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"ServiceBusSettings\":\n+        names = (\n+            \"SERVICEBUS_CONNECTION_STR\",\n+            \"SERVICEBUS_QUEUE_NAME\",\n+            \"SERVICEBUS_TOPIC_NAME\",\n+            \"SERVICEBUS_SUBSCRIPTION_NAME\",\n+        )\n+        missing = [name for name in names if not os.environ.get(name)]\n+        if missing:\n+            raise RuntimeError(\n+                \"Missing required environment variables: \" + \", \".join(missing)\n+            )\n+\n+        return cls(\n+            connection_string=os.environ[\"SERVICEBUS_CONNECTION_STR\"],\n+            queue_name=os.environ[\"SERVICEBUS_QUEUE_NAME\"],\n+            topic_name=os.environ[\"SERVICEBUS_TOPIC_NAME\"],\n+            subscription_name=os.environ[\"SERVICEBUS_SUBSCRIPTION_NAME\"],\n+        )\n+\n+\n+def send_queue_batch(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    with client.get_queue_sender(queue_name=queue_name) as sender:\n+        batch = sender.create_message_batch()\n+        for index in range(1, BATCH_SIZE + 1):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Queue message {index}\",\n+                    message_id=f\"sync-queue-{index}\",\n+                    application_properties={\"sequence\": index},\n+                )\n+            )\n+        sender.send_messages(batch)\n+        print(f\"Sent a queue batch containing {BATCH_SIZE} messages.\")\n+\n+\n+def receive_and_complete_queue(\n+    client: ServiceBusClient, queue_name: str\n+) -\u003e None:\n+    with client.get_queue_receiver(\n+        queue_name=queue_name,\n+        max_wait_time=10,\n+    ) as receiver:\n+        messages = receiver.receive_messages(\n+            max_message_count=BATCH_SIZE,\n+            max_wait_time=10,\n+        )\n+        for message in messages:\n+            print(f\"Processed queue message: {message}\")\n+            receiver.complete_message(message)\n+        print(f\"Completed {len(messages)} queue messages.\")\n+\n+\n+def send_topic_message(client: ServiceBusClient, topic_name: str) -\u003e None:\n+    with client.get_topic_sender(topic_name=topic_name) as sender:\n+        sender.send_messages(\n+            ServiceBusMessage(\n+                \"Topic message\",\n+                message_id=\"sync-topic-1\",\n+                subject=\"sync-demo\",\n+            )\n+        )\n+        print(\"Sent one topic message.\")\n+\n+\n+def receive_and_complete_subscription(\n+    client: ServiceBusClient,\n+    topic_name: str,\n+    subscription_name: str,\n+) -\u003e None:\n+    with client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+        max_wait_time=10,\n+    ) as receiver:\n+        messages = receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=10,\n+        )\n+        for message in messages:\n+            print(f\"Processed subscription message: {message}\")\n+            receiver.complete_message(message)\n+        print(f\"Completed {len(messages)} subscription messages.\")\n+\n+\n+def run(settings: ServiceBusSettings) -\u003e None:\n+    with ServiceBusClient.from_connection_string(\n+        conn_str=settings.connection_string\n+    ) as client:\n+        send_queue_batch(client, settings.queue_name)\n+        receive_and_complete_queue(client, settings.queue_name)\n+        send_topic_message(client, settings.topic_name)\n+        receive_and_complete_subscription(\n+            client,\n+            settings.topic_name,\n+            settings.subscription_name,\n+        )\n+\n+\n+def main() -\u003e None:\n+    parser = argparse.ArgumentParser(description=__doc__)\n+    parser.add_argument(\n+        \"--execute\",\n+        action=\"store_true\",\n+        help=\"Connect to Azure Service Bus; without this flag, perform a dry run.\",\n+    )\n+    args = parser.parse_args()\n+\n+    if not args.execute:\n+        print(\n+            \"Dry run: would send a 5-message queue batch, receive and complete \"\n+            \"the messages, then send to a topic and receive from a subscription.\"\n+        )\n+        return\n+\n+    run(ServiceBusSettings.from_environment())\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: async_service_bus.py\n+\"\"\"Asynchronous Azure Service Bus queue and topic/subscription example.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import asyncio\n+import os\n+from dataclasses import dataclass\n+\n+from azure.servicebus import ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient\n+\n+BATCH_SIZE = 5\n+\n+\n+@dataclass(frozen=True)\n+class ServiceBusSettings:\n+    connection_string: str\n+    queue_name: str\n+    topic_name: str\n+    subscription_name: str\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"ServiceBusSettings\":\n+        names = (\n+            \"SERVICEBUS_CONNECTION_STR\",\n+            \"SERVICEBUS_QUEUE_NAME\",\n+            \"SERVICEBUS_TOPIC_NAME\",\n+            \"SERVICEBUS_SUBSCRIPTION_NAME\",\n+        )\n+        missing = [name for name in names if not os.environ.get(name)]\n+        if missing:\n+            raise RuntimeError(\n+                \"Missing required environment variables: \" + \", \".join(missing)\n+            )\n+\n+        return cls(\n+            connection_string=os.environ[\"SERVICEBUS_CONNECTION_STR\"],\n+            queue_name=os.environ[\"SERVICEBUS_QUEUE_NAME\"],\n+            topic_name=os.environ[\"SERVICEBUS_TOPIC_NAME\"],\n+            subscription_name=os.environ[\"SERVICEBUS_SUBSCRIPTION_NAME\"],\n+        )\n+\n+\n+async def send_queue_batch(\n+    client: ServiceBusClient, queue_name: str\n+) -\u003e None:\n+    sender = client.get_queue_sender(queue_name=queue_name)\n+    async with sender:\n+        batch = await sender.create_message_batch()\n+        for index in range(1, BATCH_SIZE + 1):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Async queue message {index}\",\n+                    message_id=f\"async-queue-{index}\",\n+                    application_properties={\"sequence\": index},\n+                )\n+            )\n+        await sender.send_messages(batch)\n+        print(f\"Sent an async queue batch containing {BATCH_SIZE} messages.\")\n+\n+\n+async def receive_and_complete_queue(\n+    client: ServiceBusClient, queue_name: str\n+) -\u003e None:\n+    receiver = client.get_queue_receiver(\n+        queue_name=queue_name,\n+        max_wait_time=10,\n+    )\n+    async with receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=BATCH_SIZE,\n+            max_wait_time=10,\n+        )\n+        for message in messages:\n+            print(f\"Processed async queue message: {message}\")\n+            await receiver.complete_message(message)\n+        print(f\"Completed {len(messages)} async queue messages.\")\n+\n+\n+async def send_topic_message(\n+    client: ServiceBusClient, topic_name: str\n+) -\u003e None:\n+    sender = client.get_topic_sender(topic_name=topic_name)\n+    async with sender:\n+        await sender.send_messages(\n+            ServiceBusMessage(\n+                \"Async topic message\",\n+                message_id=\"async-topic-1\",\n+                subject=\"async-demo\",\n+            )\n+        )\n+        print(\"Sent one async topic message.\")\n+\n+\n+async def receive_and_complete_subscription(\n+    client: ServiceBusClient,\n+    topic_name: str,\n+    subscription_name: str,\n+) -\u003e None:\n+    receiver = client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+        max_wait_time=10,\n+    )\n+    async with receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=10,\n+        )\n+        for message in messages:\n+            print(f\"Processed async subscription message: {message}\")\n+            await receiver.complete_message(message)\n+        print(f\"Completed {len(messages)} async subscription messages.\")\n+\n+\n+async def run(settings: ServiceBusSettings) -\u003e None:\n+    client = ServiceBusClient.from_connection_string(\n+        conn_str=settings.connection_string\n+    )\n+    async with client:\n+        await send_queue_batch(client, settings.queue_name)\n+        await receive_and_complete_queue(client, settings.queue_name)\n+        await send_topic_message(client, settings.topic_name)\n+        await receive_and_complete_subscription(\n+            client,\n+            settings.topic_name,\n+            settings.subscription_name,\n+        )\n+\n+\n+def main() -\u003e None:\n+    parser = argparse.ArgumentParser(description=__doc__)\n+    parser.add_argument(\n+        \"--execute\",\n+        action=\"store_true\",\n+        help=\"Connect to Azure Service Bus; without this flag, perform a dry run.\",\n+    )\n+    args = parser.parse_args()\n+\n+    if not args.execute:\n+        print(\n+            \"Dry run: would asynchronously send a 5-message queue batch, receive \"\n+            \"and complete the messages, then send to a topic and receive from a \"\n+            \"subscription.\"\n+        )\n+        return\n+\n+    asyncio.run(run(ServiceBusSettings.from_environment()))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540/requirements.txt
@@ -1,0 +1,2 @@
+azure-servicebus>=7.12,<8
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540/README.md b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540/README.md
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-azure-skill-mcp-190797540/README.md
@@ -1,0 +1,53 @@
+# Azure Service Bus Python messaging examples
+
+These examples demonstrate:
+
+- Sending exactly five queue messages in a `ServiceBusMessageBatch`
+- Receiving and completing queue messages after processing
+- Sending to a topic and receiving from a subscription
+- Synchronous and `azure.servicebus.aio` asynchronous context-manager patterns
+
+## Install
+
+```powershell
+python -m pip install -r requirements.txt
+```
+
+The only required package is `azure-servicebus`. Its AMQP transport dependencies
+are installed automatically.
+
+## Configure
+
+Set these environment variables before connecting:
+
+```powershell
+$env:SERVICEBUS_CONNECTION_STR = "<namespace connection string>"
+$env:SERVICEBUS_QUEUE_NAME = "<queue name>"
+$env:SERVICEBUS_TOPIC_NAME = "<topic name>"
+$env:SERVICEBUS_SUBSCRIPTION_NAME = "<subscription name>"
+```
+
+Use a connection string whose shared access policy has permission to send and
+receive. Do not commit connection strings.
+
+## Run
+
+Both scripts default
... (truncated)
```

### ✅ 8. powershell — 3633ms

**Input:**

```json
{"command":"python -m py_compile .\\sync_service_bus.py .\\async_service_bus.py; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -c \"import ast, pathlib; files=['sync_service_bus.py','async_service_bus.py']; [ast.parse(pathlib.Path(f).read_text(encoding='utf-8')) for f in files]; print('Syntax and AST validation passed')\"","description":"Validate Python examples","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Syntax and AST validation passed
<shellId: 0 completed with exit code 0>
```

### ✅ 9. powershell — 1373ms

**Input:**

```json
{"command":"if (Test-Path -LiteralPath '.\\__pycache__') { Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force }; Get-ChildItem -File | Select-Object -ExpandProperty Name","description":"Clean validation cache and list files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
async_service_bus.py
README.md
requirements.txt
sync_service_bus.py
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `async_service_bus.py`
- `requirements.txt`
- `sync_service_bus.py`

## Copilot Response

I’m checking the skill’s SDK-specific acceptance criteria and examples, then I’ll create a local-only sample with sync, async, queue, and topic/subscription flows.

The Azure guidance endpoints timed out, so I’m using the loaded SDK skill and current `azure-servicebus` 7.x API patterns. I’ll keep execution opt-in (`--execute`) so the examples are safe and runnable offline by default.

Created:

- `sync_service_bus.py` — synchronous batch send, queue receive/complete, and topic/subscription messaging using `with`.
- `async_service_bus.py` — equivalent higher-throughput `azure.servicebus.aio` pattern using `async with`.
- `requirements.txt` — requires `azure-servicebus>=7.12,<8`.
- `README.md` — installation, environment variables, and execution instructions.

Both scripts default to an offline dry run. Pass `--execute` after configuring the required Service Bus environment variables to perform network operations.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Uses appropriate synchronous and asynchronous context managers for Service Bus clients, senders, and receivers.

### Issues

- Missing handling for azure.core.exceptions.HttpResponseError and relevant subclasses in both synchronous and asynchronous operations.

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
  - DefaultAzureCredential Usage (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Fail
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 17.6% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 4.8571** | **60.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-crud --config python-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
