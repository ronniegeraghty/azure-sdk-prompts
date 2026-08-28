# Evaluation Report: event-hubs-dp-python-streaming

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 130.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-python-streaming` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 9/14 |
| Duration | 130.7s |
| Timestamp | 2026-08-27T20:04:43Z |
| Files Generated | 2 |
| Event Count | 1810 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 40.5s |
| Review | 89.4s |
| **Total** | **130.7s** |

## Configuration

- **name:** python-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=80142 out=1191 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. glob — 6502ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 6127ms

**Input:**

```json
{"pattern":"*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. apply_patch — 1159ms

**Input:**

```json
"*** Begin Patch\n*** Add File: event_hubs_async.py\n+\"\"\"Send a batch of events, then receive and checkpoint them asynchronously.\n+\n+Required environment variables:\n+    EVENT_HUB_CONNECTION_STRING\n+    EVENT_HUB_NAME\n+    AZURE_STORAGE_CONNECTION_STRING\n+    BLOB_CONTAINER_NAME\n+\n+Stop the receiver with Ctrl+C.\n+\"\"\"\n+\n+import asyncio\n+import os\n+\n+from azure.eventhub import EventData\n+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient\n+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore\n+\n+\n+EVENT_HUB_CONNECTION_STRING = os.environ[\"EVENT_HUB_CONNECTION_STRING\"]\n+EVENT_HUB_NAME = os.environ[\"EVENT_HUB_NAME\"]\n+AZURE_STORAGE_CONNECTION_STRING = os.environ[\n+    \"AZURE_STORAGE_CONNECTION_STRING\"\n+]\n+BLOB_CONTAINER_NAME = os.environ[\"BLOB_CONTAINER_NAME\"]\n+CONSUMER_GROUP = os.getenv(\"EVENT_HUB_CONSUMER_GROUP\", \"$Default\")\n+\n+\n+async def send_events() -\u003e None:\n+    producer = EventHubProducerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        eventhub_name=EVENT_HUB_NAME,\n+    )\n+\n+    async with producer:\n+        event_batch = await producer.create_batch()\n+        for index in range(1, 4):\n+            event_batch.add(EventData(f\"Async event {index}\"))\n+\n+        await producer.send_batch(event_batch)\n+        print(f\"Sent {len(event_batch)} events\")\n+\n+\n+async def on_event(partition_context, event) -\u003e None:\n+    print(\n+        f\"Partition {partition_context.partition_id}: \"\n+        f\"{event.body_as_str(encoding='UTF-8')}\"\n+    )\n+    await partition_context.update_checkpoint(event)\n+\n+\n+async def on_error(partition_context, error) -\u003e None:\n+    if partition_context is None:\n+        print(f\"Consumer error: {error!r}\")\n+    else:\n+        print(\n+            f\"Error on partition {partition_context.partition_id}: {error!r}\"\n+        )\n+\n+\n+async def receive_events() -\u003e None:\n+    checkpoint_store = BlobCheckpointStore.from_connection_string(\n+        conn_str=AZURE_STORAGE_CONNECTION_STRING,\n+        container_name=BLOB_CONTAINER_NAME,\n+    )\n+    consumer = EventHubConsumerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        consumer_group=CONSUMER_GROUP,\n+        eventhub_name=EVENT_HUB_NAME,\n+        checkpoint_store=checkpoint_store,\n+    )\n+\n+    async with consumer:\n+        await consumer.receive(\n+            on_event=on_event,\n+            on_error=on_error,\n+            starting_position=\"-1\",\n+        )\n+\n+\n+async def main() -\u003e None:\n+    await send_events()\n+    print(\"Receiving events; press Ctrl+C to stop\")\n+    await receive_events()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        asyncio.run(main())\n+    except KeyboardInterrupt:\n+        print(\"Receiver stopped\")\n*** Add File: requirements.txt\n+azure-eventhub\u003e=5.15,\u003c6\n+azure-eventhub-checkpointstoreblob-aio\u003e=1.2,\u003c2\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-4144548716/event_hubs_async.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-4144548716/event_hubs_async.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-4144548716/event_hubs_async.py
@@ -1,0 +1,92 @@
+"""Send a batch of events, then receive and checkpoint them asynchronously.
+
+Required environment variables:
+    EVENT_HUB_CONNECTION_STRING
+    EVENT_HUB_NAME
+    AZURE_STORAGE_CONNECTION_STRING
+    BLOB_CONTAINER_NAME
+
+Stop the receiver with Ctrl+C.
+"""
+
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
+AZURE_STORAGE_CONNECTION_STRING = os.environ[
+    "AZURE_STORAGE_CONNECTION_STRING"
+]
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
+        event_batch = await producer.create_batch()
+        for index in range(1, 4):
+            event_batch.add(EventData(f"Async event {index}"))
+
+        await producer.send_batch(event_batch)
+        print(f"Sent {len(event_batch)} events")
+
+
+async def on_event(partition_context, event) -> None:
+    print(
+        f"Partition {partition_context.partition_id}: "
+        f"{event.body_as_str(encoding='UTF-8')}"
+    )
+    await
... (truncated)
```

### ✅ 4. powershell — 4152ms

**Input:**

```json
{"command":"python -m py_compile .\\event_hubs_async.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `event_hubs_async.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add a local-only async Event Hubs example and its package requirements.

Created:

- `event_hubs_async.py` — async batch sending, `EventHubConsumerClient`, blob checkpointing, event/error callbacks, and continuous receiving.
- `requirements.txt` — includes `azure-eventhub` and `azure-eventhub-checkpointstoreblob-aio`.

Configuration uses environment variables for Event Hubs and Blob Storage connection details, without provisioning Azure resources.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Provides an asynchronous on_error callback that reports consumer and partition-specific receive errors.

### Issues

- Missing explicit handling for Azure SDK exceptions such as HttpResponseError and its subclasses.

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
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: any tool (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 33.3% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 33.3% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 33.3% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 3.0000** | **37.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-python-streaming --config python-azure-skills/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
