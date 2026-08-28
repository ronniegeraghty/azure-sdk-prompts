# Evaluation Report: service-bus-dp-python-crud

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 151.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-crud` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 8/14 |
| Duration | 151.5s |
| Timestamp | 2026-08-28T21:31:10Z |
| Files Generated | 2 |
| Event Count | 3971 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 74.1s |
| Review | 76.8s |
| **Total** | **151.5s** |

## Configuration

- **name:** python-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=129451 out=2927 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. glob — 8134ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 7746ms

**Input:**

```json
{"pattern":"requirements*.txt"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 7353ms

**Input:**

```json
{"pattern":"pyproject.toml"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. apply_patch — 693ms

**Input:**

```json
"*** Begin Patch\n*** Add File: service_bus_messaging.py\n+\"\"\"Azure Service Bus queue, batch, async, topic, and subscription examples.\"\"\"\n+\n+from __future__ import annotations\n+\n+import argparse\n+import asyncio\n+import os\n+from dataclasses import dataclass\n+\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+\n+@dataclass(frozen=True)\n+class Settings:\n+    connection_string: str\n+    queue_name: str\n+    topic_name: str\n+    subscription_name: str\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"Settings\":\n+        names = (\n+            \"SERVICE_BUS_CONNECTION_STRING\",\n+            \"SERVICE_BUS_QUEUE_NAME\",\n+            \"SERVICE_BUS_TOPIC_NAME\",\n+            \"SERVICE_BUS_SUBSCRIPTION_NAME\",\n+        )\n+        missing = [name for name in names if not os.getenv(name)]\n+        if missing:\n+            raise RuntimeError(\n+                \"Set the following environment variables before running: \"\n+                + \", \".join(missing)\n+            )\n+\n+        return cls(\n+            connection_string=os.environ[\"SERVICE_BUS_CONNECTION_STRING\"],\n+            queue_name=os.environ[\"SERVICE_BUS_QUEUE_NAME\"],\n+            topic_name=os.environ[\"SERVICE_BUS_TOPIC_NAME\"],\n+            subscription_name=os.environ[\"SERVICE_BUS_SUBSCRIPTION_NAME\"],\n+        )\n+\n+\n+def add_five_messages_to_batch(sender: object) -\u003e object:\n+    \"\"\"Create one ServiceBusMessageBatch containing exactly five messages.\"\"\"\n+    batch = sender.create_message_batch()\n+    for index in range(1, 6):\n+        message = ServiceBusMessage(\n+            f\"Synchronous queue message {index}\",\n+            application_properties={\"sequence\": index, \"example\": \"sync-queue\"},\n+        )\n+        if not batch.try_add_message(message):\n+            raise RuntimeError(\n+                \"The configured Service Bus message size limit cannot fit all 5 messages.\"\n+            )\n+    return batch\n+\n+\n+def run_sync_queue(settings: Settings) -\u003e None:\n+    \"\"\"Send a five-message batch, receive it, and complete each message.\"\"\"\n+    with ServiceBusClient.from_connection_string(\n+        settings.connection_string\n+    ) as client:\n+        with client.get_queue_sender(settings.queue_name) as sender:\n+            batch = add_five_messages_to_batch(sender)\n+            sender.send_messages(batch)\n+            print(f\"Sent {len(batch)} messages to queue {settings.queue_name!r}.\")\n+\n+        with client.get_queue_receiver(\n+            settings.queue_name, max_wait_time=10\n+        ) as receiver:\n+            messages = receiver.receive_messages(max_message_count=5, max_wait_time=10)\n+            for message in messages:\n+                print(f\"Processed queue message: {message}\")\n+                receiver.complete_message(message)\n+            print(f\"Completed {len(messages)} queue messages.\")\n+\n+\n+def run_sync_topic(settings: Settings) -\u003e None:\n+    \"\"\"Send to a topic, then receive and complete from its subscription.\"\"\"\n+    with ServiceBusClient.from_connection_string(\n+        settings.connection_string\n+    ) as client:\n+        with client.get_topic_sender(settings.topic_name) as sender:\n+            sender.send_messages(\n+                ServiceBusMessage(\n+                    \"Topic message\",\n+                    subject=\"service-bus-demo\",\n+                    application_properties={\"example\": \"topic\"},\n+                )\n+            )\n+            print(f\"Sent a message to topic {settings.topic_name!r}.\")\n+\n+        with client.get_subscription_receiver(\n+            topic_name=settings.topic_name,\n+            subscription_name=settings.subscription_name,\n+            max_wait_time=10,\n+        ) as receiver:\n+            messages = receiver.receive_messages(max_message_count=1, max_wait_time=10)\n+            for message in messages:\n+                print(f\"Processed subscription message: {message}\")\n+                receiver.complete_message(message)\n+            print(f\"Completed {len(messages)} subscription messages.\")\n+\n+\n+async def run_async_queue(settings: Settings) -\u003e None:\n+    \"\"\"Use azure.servicebus.aio to send and process messages asynchronously.\"\"\"\n+    client = AsyncServiceBusClient.from_connection_string(\n+        settings.connection_string\n+    )\n+    async with client:\n+        sender = client.get_queue_sender(settings.queue_name)\n+        async with sender:\n+            batch = await sender.create_message_batch()\n+            for index in range(1, 6):\n+                message = ServiceBusMessage(\n+                    f\"Asynchronous queue message {index}\",\n+                    application_properties={\n+                        \"sequence\": index,\n+                        \"example\": \"async-queue\",\n+                    },\n+                )\n+                if not batch.try_add_message(message):\n+                    raise RuntimeError(\n+                        \"The configured Service Bus message size limit cannot fit \"\n+                        \"all 5 messages.\"\n+                    )\n+            await sender.send_messages(batch)\n+            print(f\"Sent {len(batch)} messages asynchronously.\")\n+\n+        receiver = client.get_queue_receiver(\n+            settings.queue_name, max_wait_time=10\n+        )\n+        async with receiver:\n+            messages = await receiver.receive_messages(\n+                max_message_count=5, max_wait_time=10\n+            )\n+\n+            async def process_and_complete(message: ServiceBusMessage) -\u003e None:\n+                print(f\"Processed asynchronously: {message}\")\n+                await receiver.complete_message(message)\n+\n+            await asyncio.gather(\n+                *(process_and_complete(message) for message in messages)\n+            )\n+            print(f\"Completed {len(messages)} messages asynchronously.\")\n+\n+\n+def parse_args() -\u003e argparse.Namespace:\n+    parser = argparse.ArgumentParser(description=__doc__)\n+    parser.add_argument(\n+        \"mode\",\n+        choices=(\"sync-queue\", \"async-queue\", \"topic\", \"all\"),\n+        nargs=\"?\",\n+        default=\"all\",\n+        help=\"Example to run (default: all).\",\n+    )\n+    return parser.parse_args()\n+\n+\n+def main() -\u003e None:\n+    settings = Settings.from_environment()\n+    mode = parse_args().mode\n+\n+    if mode in (\"sync-queue\", \"all\"):\n+        run_sync_queue(settings)\n+    if mode in (\"async-queue\", \"all\"):\n+        asyncio.run(run_async_queue(settings))\n+    if mode in (\"topic\", \"all\"):\n+        run_sync_topic(settings)\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-servicebus\u003e=7.12.0,\u003c8.0.0\n*** Add File: .env.example\n+# Copy these placeholders into your shell environment. Do not commit real secrets.\n+SERVICE_BUS_CONNECTION_STRING=Endpoint=sb://your-namespace.servicebus.windows.net/;SharedAccessKeyName=your-policy;SharedAccessKey=your-key\n+SERVICE_BUS_QUEUE_NAME=your-queue\n+SERVICE_BUS_TOPIC_NAME=your-topic\n+SERVICE_BUS_SUBSCRIPTION_NAME=your-subscription\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py
@@ -1,0 +1,175 @@
+"""Azure Service Bus queue, batch, async, topic, and subscription examples."""
+
+from __future__ import annotations
+
+import argparse
+import asyncio
+import os
+from dataclasses import dataclass
+
+from azure.servicebus import ServiceBusClient, ServiceBusMessage
+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
+
+
+@dataclass(frozen=True)
+class Settings:
+    connection_string: str
+    queue_name: str
+    topic_name: str
+    subscription_name: str
+
+    @classmethod
+    def from_environment(cls) -> "Settings":
+        names = (
+            "SERVICE_BUS_CONNECTION_STRING",
+            "SERVICE_BUS_QUEUE_NAME",
+            "SERVICE_BUS_TOPIC_NAME",
+            "SERVICE_BUS_SUBSCRIPTION_NAME",
+        )
+        missing = [name for name in names if not os.getenv(name)]
+        if missing:
+            raise RuntimeError(
+                "Set the following environment variables before running: "
+                + ", ".join(missing)
+            )
+
+        return cls(
+            connection_string=os.environ["SERVICE_BUS_CONNECTION_STRING"],
+            queue_name=os.environ["SERVICE_BUS_QUEUE_NAME"],
+            topic_name=os.environ["SERVICE_BUS_TOPIC_NAME"],
+            subscription_name=os.environ["SERVICE_BUS_SUBSCRIPTION_NAME"],
+        )
+
+
+def add_five_messages_to_batch(sender: object) -> object:
+    """Create one ServiceBusMessageBatch containing exactly five messages."""
+    batch = sender.create_message_batch()
+   
... (truncated)
```

### ✅ 5. powershell — 2321ms

**Input:**

```json
{"command":"python -m py_compile .\\service_bus_messaging.py","description":"Compile Python sample","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

### ✅ 6. powershell — 1925ms

**Input:**

```json
{"command":"Get-Content .\\requirements.txt; Get-Content .\\.env.example","description":"Check dependency and configuration files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
azure-servicebus>=7.12.0,<8.0.0
# Copy these placeholders into your shell environment. Do not commit real secrets.
SERVICE_BUS_CONNECTION_STRING=Endpoint=sb://your-namespace.servicebus.windows.net/;SharedAccessKeyName=your-policy;SharedAccessKey=your-key
SERVICE_BUS_QUEUE_NAME=your-queue
SERVICE_BUS_TOPIC_NAME=your-topic
SERVICE_BUS_SUBSCRIPTION_NAME=your-subscription
<shellId: 0 completed with exit code 0>
```

### ✅ 7. apply_patch — 846ms

**Input:**

```json
"*** Begin Patch\n*** Update File: service_bus_messaging.py\n@@\n def main() -\u003e None:\n-    settings = Settings.from_environment()\n     mode = parse_args().mode\n+    settings = Settings.from_environment()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-crud-python-azure-skills-baseline-256288412/service_bus_messaging.py
@@ -159,8 +159,8 @@
 
 
 def main() -> None:
+    mode = parse_args().mode
     settings = Settings.from_environment()
-    mode = parse_args().mode
 
     if mode in ("sync-queue", "all"):
         run_sync_queue(settings)

```

### ✅ 8. powershell — 1946ms

**Input:**

```json
{"command":"python -m py_compile .\\service_bus_messaging.py; $exitCode = $LASTEXITCODE; if (Test-Path .\\__pycache__\\service_bus_messaging.cpython-*.pyc) { Remove-Item .\\__pycache__\\service_bus_messaging.cpython-*.pyc }; if ((Test-Path .\\__pycache__) -and -not (Get-ChildItem .\\__pycache__)) { Remove-Item .\\__pycache__ }; exit $exitCode","description":"Compile sample and clean cache","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 2 completed with exit code 0>
```

## Generated Files

- `requirements.txt`
- `service_bus_messaging.py`

## Copilot Response

I’ll inspect the workspace, then create a self-contained sample with dependency and configuration files.

The workspace is empty. I’ll add one runnable CLI sample covering synchronous queue batching, asynchronous queue throughput, and topic/subscription messaging, plus its pip requirements and environment template.

The sample is in place. I’m checking syntax and the command-line interface without connecting to or modifying Azure resources.

Created:

- `service_bus_messaging.py` — demonstrates:
  - Five-message `ServiceBusMessageBatch`
  - Queue receiving and `receiver.complete_message()`
  - Async messaging with `azure.servicebus.aio`
  - Topic sending and subscription receiving
  - Sync and async context managers
  - CLI modes: `sync-queue`, `async-queue`, `topic`, and `all`
- `requirements.txt` — required `azure-servicebus` package
- `.env.example` — required environment-variable placeholders without real credentials

## Reviewer Notes (LLM-as-Judge)

### Summary

0/1 criteria passed. The messaging examples are present, but the required Azure SDK exception handling is absent.

### Strengths

- Uses synchronous and asynchronous context managers and demonstrates queue batching, message completion, and topic/subscription messaging.

### Issues

- No handling for HttpResponseError or its subclasses around Service Bus operations.

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
      - tool_used: any tool (source=mcp, server=azure): Fail

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
hyoka run --prompt-id service-bus-dp-python-crud --config python-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
