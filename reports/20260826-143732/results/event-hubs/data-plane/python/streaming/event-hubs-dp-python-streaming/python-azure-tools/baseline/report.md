# Evaluation Report: event-hubs-dp-python-streaming

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 129.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-python-streaming` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 10/14 |
| Duration | 129.8s |
| Timestamp | 2026-08-26T07:18:35Z |
| Files Generated | 2 |
| Event Count | 1695 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 49.7s |
| Review | 79.0s |
| **Total** | **129.8s** |

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
| Token Usage | in=82428 out=1596 |
| Turn Count | 4 |

## Prompt Sent

```
Write a Python script that demonstrates sending and receiving events with Azure Event Hubs.

**Write the code to files (use file-write tools, do not reply with code blocks).**

3. Send the batch to the event hub
4. Create an EventHubConsumerClient for receiving
5. Create a BlobCheckpointStore for checkpointing
6. Define an on_event callback that prints event body and updates checkpoint
7. Start receiving with receive() or receive_batch() using the callback
8. Handle errors with an on_error callback

Show required pip packages (azure-eventhub and
azure-eventhub-checkpointstoreblob-aio) and async patterns.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-baseline-3268921148","pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-baseline-3268921148","pattern":"**/.*"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: event_hubs_async.py\n+import asyncio\n+import os\n+\n+from azure.eventhub import EventData\n+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient\n+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore\n+\n+\n+EVENT_HUB_CONNECTION_STRING = os.environ[\"EVENT_HUB_CONNECTION_STRING\"]\n+EVENT_HUB_NAME = os.environ[\"EVENT_HUB_NAME\"]\n+STORAGE_CONNECTION_STRING = os.environ[\"AZURE_STORAGE_CONNECTION_STRING\"]\n+BLOB_CONTAINER_NAME = os.environ[\"BLOB_CONTAINER_NAME\"]\n+CONSUMER_GROUP = os.getenv(\"EVENT_HUB_CONSUMER_GROUP\", \"$Default\")\n+\n+\n+async def send_events() -\u003e None:\n+    producer = EventHubProducerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        eventhub_name=EVENT_HUB_NAME,\n+    )\n+\n+    async with producer:\n+        batch = await producer.create_batch()\n+        for message in (\"First event\", \"Second event\", \"Third event\"):\n+            batch.add(EventData(message))\n+\n+        await producer.send_batch(batch)\n+        print(f\"Sent {len(batch)} events.\")\n+\n+\n+async def on_event(partition_context, event) -\u003e None:\n+    print(\n+        f\"Partition {partition_context.partition_id}: \"\n+        f\"{event.body_as_str(encoding='UTF-8')}\"\n+    )\n+    await partition_context.update_checkpoint(event)\n+\n+\n+async def on_error(partition_context, error) -\u003e None:\n+    partition = (\n+        partition_context.partition_id if partition_context is not None else \"client\"\n+    )\n+    print(f\"Receive error ({partition}): {error!r}\")\n+\n+\n+async def receive_events() -\u003e None:\n+    checkpoint_store = BlobCheckpointStore.from_connection_string(\n+        STORAGE_CONNECTION_STRING,\n+        BLOB_CONTAINER_NAME,\n+    )\n+    consumer = EventHubConsumerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        consumer_group=CONSUMER_GROUP,\n+        eventhub_name=EVENT_HUB_NAME,\n+        checkpoint_store=checkpoint_store,\n+    )\n+\n+    async with consumer:\n+        print(\"Receiving events. Press Ctrl+C to stop.\")\n+        await consumer.receive(\n+            on_event=on_event,\n+            on_error=on_error,\n+            starting_position=\"-1\",\n+        )\n+\n+\n+async def main() -\u003e None:\n+    await send_events()\n+    await receive_events()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        asyncio.run(main())\n+    except KeyboardInterrupt:\n+        print(\"Stopped receiving.\")\n*** Add File: requirements.txt\n+azure-eventhub\n+azure-eventhub-checkpointstoreblob-aio\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"python -m py_compile .\\event_hubs_async.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-baseline-3268921148\\event_hubs_async.py"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-baseline-3268921148\\requirements.txt"}
```

## Generated Files

- `event_hubs_async.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add a local-only async Event Hubs example and its required package manifest.

The workspace is empty. I’ll create one async script that sends a batch, receives from all partitions with blob-backed checkpoints, and surfaces receive errors, plus a minimal `requirements.txt`.

The files are in place; I’m checking the script’s syntax and final file contents without contacting Azure.

Created:

- `event_hubs_async.py` — async batch producer, consumer, blob checkpointing, `on_event`, and `on_error` callbacks.
- `requirements.txt` — includes `azure-eventhub` and `azure-eventhub-checkpointstoreblob-aio`.

The script reads `EVENT_HUB_CONNECTION_STRING`, `EVENT_HUB_NAME`, `AZURE_STORAGE_CONNECTION_STRING`, and `BLOB_CONTAINER_NAME`; `EVENT_HUB_CONSUMER_GROUP` is optional and defaults to `$Default`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed. The async Event Hubs workflow is demonstrated, but the required Azure SDK exception handling is incomplete.

### Strengths

- Provides an asynchronous producer and consumer, batch sending, blob checkpointing, event checkpoint updates, and an on_error callback.

### Issues

- HttpResponseError and related Azure SDK exceptions are not explicitly caught and handled.

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
      - tool_used: azure (source=mcp, server=azure): Fail

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
hyoka run --prompt-id event-hubs-dp-python-streaming --config python-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
