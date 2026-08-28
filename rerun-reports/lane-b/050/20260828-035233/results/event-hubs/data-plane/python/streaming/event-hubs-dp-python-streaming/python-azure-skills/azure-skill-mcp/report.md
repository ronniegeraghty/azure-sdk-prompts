# Evaluation Report: event-hubs-dp-python-streaming

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 338.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-python-streaming` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/14 |
| Duration | 338.4s |
| Timestamp | 2026-08-27T19:52:33Z |
| Files Generated | 2 |
| Event Count | 1829 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 253.0s |
| Review | 84.2s |
| **Total** | **338.4s** |

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
| Token Usage | in=215111 out=1536 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 3453ms

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

### ✅ 2. glob — 5569ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 4291ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2781182276","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 184669ms

**Input:**

```json
{"intent":"Create a local-only asynchronous Python sample that sends an Event Hubs batch and receives with BlobCheckpointStore checkpointing and error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 184670ms

**Input:**

```json
{"intent":"Find current official Python async Azure Event Hubs examples for EventHubProducerClient, EventHubConsumerClient, BlobCheckpointStore, on_event checkpointing, and on_error callbacks.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. apply_patch — 2305ms

**Input:**

```json
"*** Begin Patch\n*** Add File: event_hubs_async.py\n+import asyncio\n+import os\n+\n+from azure.eventhub import EventData\n+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient\n+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore\n+\n+\n+EVENT_HUB_CONNECTION_STRING = os.environ[\"EVENT_HUB_CONNECTION_STRING\"]\n+EVENT_HUB_NAME = os.environ[\"EVENT_HUB_NAME\"]\n+STORAGE_CONNECTION_STRING = os.environ[\"AZURE_STORAGE_CONNECTION_STRING\"]\n+BLOB_CONTAINER_NAME = os.environ[\"BLOB_CONTAINER_NAME\"]\n+CONSUMER_GROUP = os.getenv(\"EVENT_HUB_CONSUMER_GROUP\", \"$Default\")\n+\n+\n+async def send_events() -\u003e None:\n+    producer = EventHubProducerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        eventhub_name=EVENT_HUB_NAME,\n+    )\n+\n+    async with producer:\n+        batch = await producer.create_batch()\n+        for index in range(10):\n+            batch.add(EventData(f\"Event {index}\"))\n+        await producer.send_batch(batch)\n+        print(f\"Sent {len(batch)} events\")\n+\n+\n+async def on_event(partition_context, event) -\u003e None:\n+    if event is None:\n+        return\n+\n+    print(\n+        f\"Partition {partition_context.partition_id}: \"\n+        f\"{event.body_as_str(encoding='UTF-8')}\"\n+    )\n+    await partition_context.update_checkpoint(event)\n+\n+\n+async def on_error(partition_context, error) -\u003e None:\n+    partition_id = (\n+        partition_context.partition_id if partition_context is not None else \"client\"\n+    )\n+    print(f\"Receive error on {partition_id}: {error!r}\")\n+\n+\n+async def receive_events() -\u003e None:\n+    checkpoint_store = BlobCheckpointStore.from_connection_string(\n+        STORAGE_CONNECTION_STRING,\n+        BLOB_CONTAINER_NAME,\n+    )\n+    consumer = EventHubConsumerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        consumer_group=CONSUMER_GROUP,\n+        eventhub_name=EVENT_HUB_NAME,\n+        checkpoint_store=checkpoint_store,\n+    )\n+\n+    async with consumer:\n+        await consumer.receive(\n+            on_event=on_event,\n+            on_error=on_error,\n+            starting_position=\"-1\",\n+        )\n+\n+\n+async def main() -\u003e None:\n+    await send_events()\n+    print(\"Receiving events. Press Ctrl+C to stop.\")\n+    await receive_events()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        asyncio.run(main())\n+    except KeyboardInterrupt:\n+        print(\"Stopped receiving\")\n*** Add File: requirements.txt\n+azure-eventhub\n+azure-eventhub-checkpointstoreblob-aio\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2781182276/event_hubs_async.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2781182276/event_hubs_async.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2781182276/event_hubs_async.py
@@ -1,0 +1,79 @@
+import asyncio
+import os
+
+from azure.eventhub import EventData
+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
+
+
+EVENT_HUB_CONNECTION_STRING = os.environ["EVENT_HUB_CONNECTION_STRING"]
+EVENT_HUB_NAME = os.environ["EVENT_HUB_NAME"]
+STORAGE_CONNECTION_STRING = os.environ["AZURE_STORAGE_CONNECTION_STRING"]
+BLOB_CONTAINER_NAME = os.environ["BLOB_CONTAINER_NAME"]
+CONSUMER_GROUP = os.getenv("EVENT_HUB_CONSUMER_GROUP", "$Default")
+
+
+async def send_events() -> None:
+    producer = EventHubProducerClient.from_connection_string(
+        conn_str=EVENT_HUB_CONNECTION_STRING,
+        eventhub_name=EVENT_HUB_NAME,
+    )
+
+    async with producer:
+        batch = await producer.create_batch()
+        for index in range(10):
+            batch.add(EventData(f"Event {index}"))
+        await producer.send_batch(batch)
+        print(f"Sent {len(batch)} events")
+
+
+async def on_event(partition_context, event) -> None:
+    if event is None:
+        return
+
+    print(
+        f"Partition {partition_context.partition_id}: "
+        f"{event.body_as_str(encoding='UTF-8')}"
+    )
+    await partition_context.update_checkpoint(event)
+
+
+async def on_error(partition_context, error) -> None:
+    partition_id = (
+        partition_context.partition_id if partition_context is not None else "client"
+    )
+    print(f"Receive error on {part
... (truncated)
```

### ✅ 7. powershell — 3259ms

**Input:**

```json
{"command":"python -m py_compile .\\event_hubs_async.py","description":"Check Python sample syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `event_hubs_async.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the Event Hubs skill criteria and the empty workspace, then create the async sample and dependency file.

The workspace is empty and the reference files are not present in the installed skill. I’ll use environment-based credentials and keep the sample declarative—no Azure resources or deployments will be touched.

Created:

- `event_hubs_async.py` — async batch sending, `EventHubConsumerClient`, Blob checkpointing, event/error callbacks, and continuous receiving.
- `requirements.txt` — includes `azure-eventhub` and `azure-eventhub-checkpointstoreblob-aio`.

Configuration uses `EVENT_HUB_CONNECTION_STRING`, `EVENT_HUB_NAME`, `AZURE_STORAGE_CONNECTION_STRING`, `BLOB_CONTAINER_NAME`, and optional `EVENT_HUB_CONSUMER_GROUP` environment variables.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- The asynchronous on_error callback reports errors surfaced by EventHubConsumerClient.receive().

### Issues

- Azure SDK exceptions, including HttpResponseError and its subclasses, are not explicitly caught or handled.

## Grader Results

- send-receive-events.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `azure-eventhub` and `azure-eventhub-checkpointstoreblob-aio` pip packages: Pass
      - `create_batch()` and `EventDataBatch.add()`: Pass
      - `send_batch()` for publishing: Pass
      - `EventHubConsumerClient` with `BlobCheckpointStore`: Pass
      - `receive()` with `on_event` and `on_error` callbacks: Pass
      - Async variants with `aio` module: Pass
      - Context manager (async with) patterns: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Fail
  - Context Manager for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Fail
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 4.0000** | **50.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-python-streaming --config python-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
