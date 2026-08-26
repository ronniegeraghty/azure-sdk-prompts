# Evaluation Report: service-bus-dp-python-crud

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 162.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-crud` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 8/14 |
| Duration | 162.6s |
| Timestamp | 2026-08-26T08:33:55Z |
| Files Generated | 3 |
| Event Count | 4569 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 71.7s |
| Review | 89.9s |
| **Total** | **162.6s** |

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
| Token Usage | in=107190 out=2966 |
| Turn Count | 5 |

## Prompt Sent

```
Write a Python script that demonstrates messaging with Azure Service Bus.

**Write the code to files (use file-write tools, do not reply with code blocks).**

3. Send a batch of 5 messages using a ServiceBusMessageBatch
4. Get a receiver for the queue and receive messages
5. Complete a message with receiver.complete_message() after processing
6. Demonstrate the async pattern using aio module for higher throughput
7. Send to a topic and receive from a subscription

Show required pip packages and proper context manager patterns (with statements).
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"pattern":"*"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: service_bus_demo.py\n+\"\"\"Azure Service Bus queue and topic messaging examples.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import asyncio\n+import os\n+from dataclasses import dataclass\n+\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+\n+@dataclass(frozen=True)\n+class Settings:\n+    connection_string: str\n+    queue_name: str\n+    topic_name: str\n+    subscription_name: str\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"Settings\":\n+        variable_names = (\n+            \"AZURE_SERVICEBUS_CONNECTION_STRING\",\n+            \"AZURE_SERVICEBUS_QUEUE_NAME\",\n+            \"AZURE_SERVICEBUS_TOPIC_NAME\",\n+            \"AZURE_SERVICEBUS_SUBSCRIPTION_NAME\",\n+        )\n+        missing = [name for name in variable_names if not os.environ.get(name)]\n+        if missing:\n+            raise RuntimeError(\n+                \"Set the following environment variables: \" + \", \".join(missing)\n+            )\n+\n+        return cls(\n+            connection_string=os.environ[\"AZURE_SERVICEBUS_CONNECTION_STRING\"],\n+            queue_name=os.environ[\"AZURE_SERVICEBUS_QUEUE_NAME\"],\n+            topic_name=os.environ[\"AZURE_SERVICEBUS_TOPIC_NAME\"],\n+            subscription_name=os.environ[\"AZURE_SERVICEBUS_SUBSCRIPTION_NAME\"],\n+        )\n+\n+\n+def send_queue_batch(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    with client.get_queue_sender(queue_name=queue_name) as sender:\n+        batch = sender.create_message_batch()\n+        for message_number in range(1, 6):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Sync queue message {message_number}\",\n+                    application_properties={\"sequence\": message_number},\n+                )\n+            )\n+        sender.send_messages(batch)\n+    print(\"Sent a batch of 5 messages to the queue.\")\n+\n+\n+def receive_queue_messages(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    with client.get_queue_receiver(\n+        queue_name=queue_name,\n+        max_wait_time=10,\n+    ) as receiver:\n+        messages = receiver.receive_messages(max_message_count=5, max_wait_time=10)\n+        for message in messages:\n+            print(f\"Processing queue message: {message}\")\n+            receiver.complete_message(message)\n+    print(f\"Completed {len(messages)} queue message(s).\")\n+\n+\n+def send_topic_message(client: ServiceBusClient, topic_name: str) -\u003e None:\n+    with client.get_topic_sender(topic_name=topic_name) as sender:\n+        sender.send_messages(ServiceBusMessage(\"Sync topic message\"))\n+    print(\"Sent a message to the topic.\")\n+\n+\n+def receive_subscription_message(\n+    client: ServiceBusClient,\n+    topic_name: str,\n+    subscription_name: str,\n+) -\u003e None:\n+    with client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+        max_wait_time=10,\n+    ) as receiver:\n+        messages = receiver.receive_messages(max_message_count=1, max_wait_time=10)\n+        for message in messages:\n+            print(f\"Processing subscription message: {message}\")\n+            receiver.complete_message(message)\n+    print(f\"Completed {len(messages)} subscription message(s).\")\n+\n+\n+def run_sync(settings: Settings) -\u003e None:\n+    with ServiceBusClient.from_connection_string(\n+        conn_str=settings.connection_string\n+    ) as client:\n+        send_queue_batch(client, settings.queue_name)\n+        receive_queue_messages(client, settings.queue_name)\n+        send_topic_message(client, settings.topic_name)\n+        receive_subscription_message(\n+            client,\n+            settings.topic_name,\n+            settings.subscription_name,\n+        )\n+\n+\n+async def send_queue_batch_async(\n+    client: AsyncServiceBusClient,\n+    queue_name: str,\n+) -\u003e None:\n+    async with client.get_queue_sender(queue_name=queue_name) as sender:\n+        batch = await sender.create_message_batch()\n+        for message_number in range(1, 6):\n+            batch.add_message(\n+                ServiceBusMessage(\n+                    f\"Async queue message {message_number}\",\n+                    application_properties={\"sequence\": message_number},\n+                )\n+            )\n+        await sender.send_messages(batch)\n+    print(\"Asynchronously sent a batch of 5 messages to the queue.\")\n+\n+\n+async def send_topic_message_async(\n+    client: AsyncServiceBusClient,\n+    topic_name: str,\n+) -\u003e None:\n+    async with client.get_topic_sender(topic_name=topic_name) as sender:\n+        await sender.send_messages(ServiceBusMessage(\"Async topic message\"))\n+    print(\"Asynchronously sent a message to the topic.\")\n+\n+\n+async def receive_queue_messages_async(\n+    client: AsyncServiceBusClient,\n+    queue_name: str,\n+) -\u003e None:\n+    async with client.get_queue_receiver(\n+        queue_name=queue_name,\n+        max_wait_time=10,\n+    ) as receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=5,\n+            max_wait_time=10,\n+        )\n+        for message in messages:\n+            print(f\"Async processing queue message: {message}\")\n+            await receiver.complete_message(message)\n+    print(f\"Asynchronously completed {len(messages)} queue message(s).\")\n+\n+\n+async def receive_subscription_message_async(\n+    client: AsyncServiceBusClient,\n+    topic_name: str,\n+    subscription_name: str,\n+) -\u003e None:\n+    async with client.get_subscription_receiver(\n+        topic_name=topic_name,\n+        subscription_name=subscription_name,\n+        max_wait_time=10,\n+    ) as receiver:\n+        messages = await receiver.receive_messages(\n+            max_message_count=1,\n+            max_wait_time=10,\n+        )\n+        for message in messages:\n+            print(f\"Async processing subscription message: {message}\")\n+            await receiver.complete_message(message)\n+    print(f\"Asynchronously completed {len(messages)} subscription message(s).\")\n+\n+\n+async def run_async(settings: Settings) -\u003e None:\n+    async with AsyncServiceBusClient.from_connection_string(\n+        conn_str=settings.connection_string\n+    ) as client:\n+        # Independent queue and topic operations run concurrently for higher throughput.\n+        await asyncio.gather(\n+            send_queue_batch_async(client, settings.queue_name),\n+            send_topic_message_async(client, settings.topic_name),\n+        )\n+        await asyncio.gather(\n+            receive_queue_messages_async(client, settings.queue_name),\n+            receive_subscription_message_async(\n+                client,\n+                settings.topic_name,\n+                settings.subscription_name,\n+            ),\n+        )\n+\n+\n+def parse_args() -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(description=__doc__)\n+    parser.add_argument(\n+        \"--mode\",\n+        choices=(\"sync\", \"async\", \"both\"),\n+        default=\"both\",\n+        help=\"Messaging pattern to run (default: both).\",\n+    )\n+    return parser.parse_args()\n+\n+\n+def main() -\u003e None:\n+    args = parse_args()\n+    settings = Settings.from_environment()\n+\n+    if args.mode in (\"sync\", \"both\"):\n+        run_sync(settings)\n+    if args.mode in (\"async\", \"both\"):\n+        asyncio.run(run_async(settings))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-servicebus==7.14.3\n*** Add File: README.md\n+# Azure Service Bus Python demo\n+\n+This example uses an existing Azure Service Bus namespace, queue, topic, and\n+subscription. It does not create or modify Azure resources.\n+\n+## Install\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+## Configure\n+\n+Set the connection string and entity names in the current PowerShell session:\n+\n+```powershell\n+$env:AZURE_SERVICEBUS_CONNECTION_STRING = \"\u003cservice-bus-connection-string\u003e\"\n+$env:AZURE_SERVICEBUS_QUEUE_NAME = \"\u003cqueue-name\u003e\"\n+$env:AZURE_SERVICEBUS_TOPIC_NAME = \"\u003ctopic-name\u003e\"\n+$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = \"\u003csubscription-name\u003e\"\n+```\n+\n+The connection string must grant send and listen permissions. The queue, topic,\n+and subscription must already exist.\n+\n+## Run\n+\n+Run both synchronous and asynchronous examples:\n+\n+```powershell\n+python .\\service_bus_demo.py\n+```\n+\n+Run only one pattern:\n+\n+```powershell\n+python .\\service_bus_demo.py --mode sync\n+python .\\service_bus_demo.py --mode async\n+```\n+\n+The script sends a five-message `ServiceBusMessageBatch` to the queue, receives\n+and completes those messages, then sends to a topic and receives and completes\n+the message through a subscription. The asynchronous example uses\n+`azure.servicebus.aio` and runs independent queue and topic operations\n+concurrently with `asyncio.gather`.\n*** End Patch\n"
```

### 🔧 4. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: requirements.txt\n@@\n-azure-servicebus==7.14.3\n+azure-servicebus\u003e=7.12.0,\u003c8.0.0\n*** End Patch\n"
```

### 🔧 5. powershell

**Input:**

```json
{"command":"python -m pip install --quiet -r .\\requirements.txt \u0026\u0026 python -m py_compile .\\service_bus_demo.py \u0026\u0026 python .\\service_bus_demo.py --help","description":"Install dependency and validate demo","initial_wait":120,"mode":"sync"}
```

## Generated Files

- `README.md`
- `requirements.txt`
- `service_bus_demo.py`

## Copilot Response

I’ll inspect the workspace, then add runnable sync and async examples plus dependency and configuration files.

The workspace is empty. I’m adding one configurable demo that covers queue batching, completion, async concurrency, and topic/subscription messaging without provisioning resources.

The implementation is complete; I’m tightening the dependency constraint for SDK compatibility and checking the script’s syntax and CLI entry point.

Created:

- `service_bus_demo.py` — synchronous and `aio` asynchronous examples covering five-message batches, queue receiving/completion, and topic/subscription messaging.
- `requirements.txt` — Azure Service Bus SDK dependency.
- `README.md` — installation, environment configuration, and run instructions.

The demo uses proper `with`/`async with` context managers and environment variables for existing Azure Service Bus resources.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed. The messaging examples are implemented, but the required Azure SDK exception handling is absent.

### Strengths

- Uses synchronous and asynchronous context managers and demonstrates queue batches, message completion, and topic/subscription messaging.

### Issues

- No handling for HttpResponseError or relevant subclasses around Service Bus operations.

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 71% | 1.00 | 0.7143 | 19.2% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 26.9% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 26.9% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 26.9% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 3.7143** | **46.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-crud --config python-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
