# Evaluation Report: event-hubs-dp-python-streaming

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 245.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-python-streaming` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 12/14 |
| Duration | 245.0s |
| Timestamp | 2026-08-26T07:20:45Z |
| Files Generated | 3 |
| Event Count | 3616 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 159.4s |
| Review | 84.3s |
| **Total** | **245.0s** |

## Configuration

- **name:** python-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-eventhub-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=510104 out=4552 |
| Turn Count | 11 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-eventhub-py"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get","intent":"Generate a local-only Python async Azure Event Hubs producer/consumer sample using DefaultAzureCredential and BlobCheckpointStore, with error handling and no resource provisioning.","parameters":{"language":"Python","operation":"code-generation","resource":"Azure Event Hubs"}}
```

### 🔧 3. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-eventhub-py\\references\\checkpointing.md"}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-with-azure-tools-937540895","pattern":"**/*"}
```

### 🔧 5. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-eventhub-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

### 🔧 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Python code-generation practices for an async Event Hubs producer/consumer sample with Blob checkpointing.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Python Azure Event Hubs async BlobCheckpointStore receive callback documentation and lifecycle guidance.","parameters":{"query":"Azure Event Hubs Python async EventHubConsumerClient BlobCheckpointStore aio on_event on_error receive"}}
```

### 🔧 8. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com/en-us/python/api azure.eventhub.extensions.checkpointstoreblob.aio.BlobCheckpointStore close async context manager"}
```

### 🔧 9. rg

**Input:**

```json
{"-n":true,"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787728916828-copilot-tool-output-28afec5693cb419387715364b1b427d1.txt","pattern":"BlobCheckpointStore|EventHubConsumerClient|on_error|receive\\("}
```

### 🔧 10. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"class BlobCheckpointStore repo:Azure/azure-sdk-for-python path:sdk/eventhub/azure-eventhub-checkpointstoreblob-aio"}
```

### 🔧 11. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/eventhub/azure-eventhub-checkpointstoreblob-aio/azure/eventhub/extensions/checkpointstoreblobaio/_blobstoragecsaio.py","repo":"azure-sdk-for-python"}
```

### 🔧 12. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/eventhub/azure-eventhub-checkpointstoreblob-aio/azure/eventhub/extensions/checkpointstoreblobaio/__init__.py","repo":"azure-sdk-for-python"}
```

### 🔧 13. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/eventhub/azure-eventhub-checkpointstoreblob-aio/samples/event_processor_sample.py","repo":"azure-sdk-for-python"}
```

### 🔧 14. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/eventhub/azure-eventhub-checkpointstoreblob-aio/README.md","repo":"azure-sdk-for-python"}
```

### 🔧 15. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: event_hubs_async.py\n+\"\"\"Send and receive Azure Event Hubs events with async checkpointing.\"\"\"\n+\n+import asyncio\n+import logging\n+import os\n+import sys\n+\n+from azure.core.exceptions import AzureError\n+from azure.eventhub import EventData\n+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient\n+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore\n+from azure.identity.aio import DefaultAzureCredential\n+\n+LOGGER = logging.getLogger(\"event_hubs_async\")\n+\n+\n+def required_setting(name: str) -\u003e str:\n+    value = os.getenv(name)\n+    if not value:\n+        raise ValueError(f\"Set the required environment variable {name}.\")\n+    return value\n+\n+\n+async def send_events(\n+    credential: DefaultAzureCredential,\n+    namespace: str,\n+    event_hub_name: str,\n+) -\u003e None:\n+    async with EventHubProducerClient(\n+        fully_qualified_namespace=namespace,\n+        eventhub_name=event_hub_name,\n+        credential=credential,\n+    ) as producer:\n+        batch = await producer.create_batch()\n+        for event_number in range(1, 6):\n+            batch.add(EventData(f\"Async event {event_number}\"))\n+\n+        await producer.send_batch(batch)\n+        LOGGER.info(\"Sent %d events.\", len(batch))\n+\n+\n+async def receive_events(\n+    credential: DefaultAzureCredential,\n+    namespace: str,\n+    event_hub_name: str,\n+    consumer_group: str,\n+    storage_account_url: str,\n+    checkpoint_container: str,\n+) -\u003e None:\n+    async def on_event(partition_context, event) -\u003e None:\n+        print(\n+            f\"Partition {partition_context.partition_id}: \"\n+            f\"{event.body_as_str(encoding='UTF-8')}\"\n+        )\n+        await partition_context.update_checkpoint(event)\n+\n+    async def on_error(partition_context, error: Exception) -\u003e None:\n+        partition_id = (\n+            partition_context.partition_id if partition_context else \"all partitions\"\n+        )\n+        LOGGER.error(\"Receive error for %s: %s\", partition_id, error)\n+\n+    async with BlobCheckpointStore(\n+        blob_account_url=storage_account_url,\n+        container_name=checkpoint_container,\n+        credential=credential,\n+    ) as checkpoint_store:\n+        async with EventHubConsumerClient(\n+            fully_qualified_namespace=namespace,\n+            eventhub_name=event_hub_name,\n+            consumer_group=consumer_group,\n+            credential=credential,\n+            checkpoint_store=checkpoint_store,\n+        ) as consumer:\n+            LOGGER.info(\"Receiving events. Press Ctrl+C to stop.\")\n+            await consumer.receive(\n+                on_event=on_event,\n+                on_error=on_error,\n+                starting_position=\"-1\",\n+            )\n+\n+\n+async def main() -\u003e None:\n+    namespace = required_setting(\"EVENT_HUB_FULLY_QUALIFIED_NAMESPACE\")\n+    event_hub_name = required_setting(\"EVENT_HUB_NAME\")\n+    storage_account_url = required_setting(\"STORAGE_ACCOUNT_URL\")\n+    checkpoint_container = required_setting(\"CHECKPOINT_CONTAINER\")\n+    consumer_group = os.getenv(\"EVENT_HUB_CONSUMER_GROUP\", \"$Default\")\n+\n+    async with DefaultAzureCredential() as credential:\n+        await send_events(credential, namespace, event_hub_name)\n+        await receive_events(\n+            credential,\n+            namespace,\n+            event_hub_name,\n+            consumer_group,\n+            storage_account_url,\n+            checkpoint_container,\n+        )\n+\n+\n+if __name__ == \"__main__\":\n+    logging.basicConfig(\n+        level=os.getenv(\"LOG_LEVEL\", \"INFO\").upper(),\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    try:\n+        asyncio.run(main())\n+    except KeyboardInterrupt:\n+        LOGGER.info(\"Receiver stopped.\")\n+    except (AzureError, ValueError) as error:\n+        LOGGER.error(\"%s\", error)\n+        sys.exit(1)\n*** Add File: requirements.txt\n+azure-eventhub\n+azure-eventhub-checkpointstoreblob-aio\n+azure-identity\n*** Add File: README.md\n+# Async Azure Event Hubs send and receive sample\n+\n+This sample sends a batch of five events, receives events asynchronously, and\n+stores a checkpoint in Azure Blob Storage after each event is processed.\n+Authentication uses `DefaultAzureCredential`; no secrets or connection strings\n+are stored in the code.\n+\n+## Install\n+\n+Python 3.9 or later is required.\n+\n+```powershell\n+python -m pip install -r requirements.txt\n+```\n+\n+The required Event Hubs packages are `azure-eventhub` and\n+`azure-eventhub-checkpointstoreblob-aio`. The sample also installs\n+`azure-identity` for Microsoft Entra authentication.\n+\n+## Configure and run\n+\n+Set these variables to existing Azure resources:\n+\n+```powershell\n+$env:EVENT_HUB_FULLY_QUALIFIED_NAMESPACE = \"\u003cnamespace\u003e.servicebus.windows.net\"\n+$env:EVENT_HUB_NAME = \"\u003cevent-hub-name\u003e\"\n+$env:EVENT_HUB_CONSUMER_GROUP = '$Default'\n+$env:STORAGE_ACCOUNT_URL = \"https://\u003caccount\u003e.blob.core.windows.net\"\n+$env:CHECKPOINT_CONTAINER = \"\u003cexisting-container-name\u003e\"\n+python .\\event_hubs_async.py\n+```\n+\n+The identity selected by `DefaultAzureCredential` needs the Azure Event Hubs\n+Data Sender and Azure Event Hubs Data Receiver roles, plus Storage Blob Data\n+Contributor access to the checkpoint container. The receiver runs until\n+Ctrl+C. `starting_position=\"-1\"` applies only when a partition has no stored\n+checkpoint; subsequent runs resume from the checkpoint.\n+\n+## References\n+\n+- [Azure Event Hubs Python client library](https://learn.microsoft.com/python/api/overview/azure/eventhub-readme)\n+- [Async Blob checkpoint store package](https://github.com/Azure/azure-sdk-for-python/tree/main/sdk/eventhub/azure-eventhub-checkpointstoreblob-aio)\n*** End Patch\n"
```

### 🔧 16. powershell

**Input:**

```json
{"command":"python -m py_compile .\\event_hubs_async.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-with-azure-tools-937540895\\event_hubs_async.py"}
```

### 🔧 18. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-with-azure-tools-937540895\\requirements.txt"}
```

### 🔧 19. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-tools-with-azure-tools-937540895\\README.md"}
```

### 🔧 20. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"class EventDataBatch def __len__ repo:Azure/azure-sdk-for-python path:sdk/eventhub/azure-eventhub"}
```

### 🔧 21. powershell

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -LiteralPath .\\__pycache__ -Recurse -Force }","description":"Remove generated Python cache","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `README.md`
- `event_hubs_async.py`
- `requirements.txt`

## Copilot Response

I’m grounding the async implementation in the Event Hubs checkpointing guidance, then I’ll write the runnable local sample and dependency manifest.

The implementation will use one async credential lifecycle, async producer/consumer context managers, per-event checkpoints, and callback-based receive errors; configuration stays entirely in environment variables.

Created the async Azure Event Hubs sample:

- `event_hubs_async.py` — sends a batch, receives via `EventHubConsumerClient`, checkpoints each event with async `BlobCheckpointStore`, and handles receive errors.
- `requirements.txt` — includes `azure-eventhub`, `azure-eventhub-checkpointstoreblob-aio`, and `azure-identity`.
- `README.md` — installation, environment configuration, RBAC requirements, and official references.

The sample uses `DefaultAzureCredential` and does not provision or modify Azure resources.

## Reviewer Notes (LLM-as-Judge)

### Summary

The exception-handling criterion is fully met.

### Strengths

- Handles Azure SDK exceptions without silently swallowing them.
- Logs receiver errors with relevant partition context.
- Separately handles configuration errors and graceful user interruption.

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
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Pass
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Pass
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 6.0000** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-python-streaming --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
