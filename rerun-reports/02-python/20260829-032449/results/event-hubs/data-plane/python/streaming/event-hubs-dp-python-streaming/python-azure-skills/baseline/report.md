# Evaluation Report: event-hubs-dp-python-streaming

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 140.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-python-streaming` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 10/14 |
| Duration | 140.5s |
| Timestamp | 2026-08-28T23:29:12Z |
| Files Generated | 2 |
| Event Count | 1976 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 47.9s |
| Review | 92.0s |
| **Total** | **140.5s** |

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
| Token Usage | in=82214 out=1657 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. glob — 10442ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-1876095667","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 9338ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-1876095667","pattern":"**/pyproject.toml"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 8299ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-1876095667","pattern":"**/requirements*.txt"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. apply_patch — 2583ms

**Input:**

```json
"*** Begin Patch\n*** Add File: event_hubs_async.py\n+import asyncio\n+import os\n+from typing import Optional\n+\n+from azure.eventhub import EventData\n+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient\n+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore\n+from azure.eventhub.aio import PartitionContext\n+\n+\n+EVENT_HUB_CONNECTION_STRING = os.environ[\"EVENT_HUB_CONNECTION_STRING\"]\n+EVENT_HUB_NAME = os.environ[\"EVENT_HUB_NAME\"]\n+EVENT_HUB_CONSUMER_GROUP = os.getenv(\"EVENT_HUB_CONSUMER_GROUP\", \"$Default\")\n+STORAGE_CONNECTION_STRING = os.environ[\"AZURE_STORAGE_CONNECTION_STRING\"]\n+BLOB_CONTAINER_NAME = os.environ[\"BLOB_CONTAINER_NAME\"]\n+\n+\n+async def send_events() -\u003e None:\n+    producer = EventHubProducerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        eventhub_name=EVENT_HUB_NAME,\n+    )\n+\n+    async with producer:\n+        batch = await producer.create_batch()\n+        for index in range(10):\n+            batch.add(EventData(f\"Event {index}\"))\n+\n+        await producer.send_batch(batch)\n+        print(f\"Sent {len(batch)} events\")\n+\n+\n+async def on_event(\n+    partition_context: PartitionContext,\n+    event: EventData,\n+) -\u003e None:\n+    print(\n+        f\"Partition {partition_context.partition_id}: \"\n+        f\"{event.body_as_str(encoding='UTF-8')}\"\n+    )\n+    await partition_context.update_checkpoint(event)\n+\n+\n+async def on_error(\n+    partition_context: Optional[PartitionContext],\n+    error: Exception,\n+) -\u003e None:\n+    if partition_context is None:\n+        print(f\"Consumer error: {error!r}\")\n+    else:\n+        print(\n+            f\"Partition {partition_context.partition_id} error: {error!r}\"\n+        )\n+\n+\n+async def receive_events() -\u003e None:\n+    checkpoint_store = BlobCheckpointStore.from_connection_string(\n+        STORAGE_CONNECTION_STRING,\n+        BLOB_CONTAINER_NAME,\n+    )\n+    consumer = EventHubConsumerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        consumer_group=EVENT_HUB_CONSUMER_GROUP,\n+        eventhub_name=EVENT_HUB_NAME,\n+        checkpoint_store=checkpoint_store,\n+    )\n+\n+    async with consumer:\n+        await consumer.receive(\n+            on_event=on_event,\n+            on_error=on_error,\n+            starting_position=\"-1\",\n+        )\n+\n+\n+async def main() -\u003e None:\n+    await send_events()\n+    print(\"Receiving events. Press Ctrl+C to stop.\")\n+    await receive_events()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        asyncio.run(main())\n+    except KeyboardInterrupt:\n+        print(\"Stopped receiving events\")\n*** Add File: requirements.txt\n+azure-eventhub\u003e=5.15.0,\u003c6.0.0\n+azure-eventhub-checkpointstoreblob-aio\u003e=1.2.0,\u003c2.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-1876095667/event_hubs_async.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-1876095667/event_hubs_async.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-baseline-1876095667/event_hubs_async.py
@@ -1,0 +1,87 @@
+import asyncio
+import os
+from typing import Optional
+
+from azure.eventhub import EventData
+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
+from azure.eventhub.aio import PartitionContext
+
+
+EVENT_HUB_CONNECTION_STRING = os.environ["EVENT_HUB_CONNECTION_STRING"]
+EVENT_HUB_NAME = os.environ["EVENT_HUB_NAME"]
+EVENT_HUB_CONSUMER_GROUP = os.getenv("EVENT_HUB_CONSUMER_GROUP", "$Default")
+STORAGE_CONNECTION_STRING = os.environ["AZURE_STORAGE_CONNECTION_STRING"]
+BLOB_CONTAINER_NAME = os.environ["BLOB_CONTAINER_NAME"]
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
+
+        await producer.send_batch(batch)
+        print(f"Sent {len(batch)} events")
+
+
+async def on_event(
+    partition_context: PartitionContext,
+    event: EventData,
+) -> None:
+    print(
+        f"Partition {partition_context.partition_id}: "
+        f"{event.body_as_str(encoding='UTF-8')}"
+    )
+    await partition_context.update_checkpoint(event)
+
+
+async def on_error(
+    partition_context: Optional[PartitionContext],
+    error: Exception,
+) -> None:
+    if partition_context i
... (truncated)
```

### ✅ 5. powershell — 4549ms

**Input:**

```json
{"command":"python -m py_compile .\\event_hubs_async.py; $exitCode = $LASTEXITCODE; if (Test-Path .\\__pycache__) { Remove-Item -Recurse -Force .\\__pycache__ }; exit $exitCode","description":"Validate sample Python syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

## Generated Files

- `event_hubs_async.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add an async Event Hubs sample and its dependency file using environment-based configuration.

The workspace is empty. I’m adding a self-contained async producer/consumer sample plus pinned-minimum package requirements.

The sample is in place; I’m checking Python syntax and removing the compiler cache afterward.

Created:

- `event_hubs_async.py` — async batch sending, receiving, blob checkpointing, event/error callbacks, and graceful shutdown.
- `requirements.txt` — includes `azure-eventhub` and `azure-eventhub-checkpointstoreblob-aio`.

Configuration uses `EVENT_HUB_CONNECTION_STRING`, `EVENT_HUB_NAME`, `EVENT_HUB_CONSUMER_GROUP`, `AZURE_STORAGE_CONNECTION_STRING`, and `BLOB_CONTAINER_NAME` environment variables.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Provides an asynchronous on_error callback that reports consumer and partition-specific receive errors.

### Issues

- Missing handling for HttpResponseError and relevant subclasses during Azure SDK operations.

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 25.0% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 4.0000** | **50.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-python-streaming --config python-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
