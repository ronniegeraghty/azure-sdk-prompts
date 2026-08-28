# Evaluation Report: event-hubs-dp-python-streaming

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 233.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-python-streaming` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/14 |
| Duration | 233.2s |
| Timestamp | 2026-08-28T23:37:16Z |
| Files Generated | 2 |
| Event Count | 1860 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 120.9s |
| Review | 111.1s |
| **Total** | **233.2s** |

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
| Token Usage | in=222414 out=1735 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 2500ms

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

### ✅ 2. glob — 8513ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2120679032","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 6196ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 7216ms

**Input:**

```json
{"command":"get_ai_app_best_practices","intent":"Create a local-only Python async Azure Event Hubs send/receive sample using BlobCheckpointStore and environment-variable configuration."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. glob — 7542ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/*.md"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-messaging\SKILL.md
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 6163ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation guidance for an asynchronous Python Azure Event Hubs sample with Blob checkpointing.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. apply_patch — 3279ms

**Input:**

```json
"*** Begin Patch\n*** Add File: event_hubs_async.py\n+import asyncio\n+import os\n+from typing import Optional\n+\n+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient\n+from azure.eventhub import EventData\n+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore\n+\n+\n+EVENT_HUB_CONNECTION_STRING = os.environ[\"EVENT_HUB_CONNECTION_STRING\"]\n+EVENT_HUB_NAME = os.environ[\"EVENT_HUB_NAME\"]\n+STORAGE_CONNECTION_STRING = os.environ[\"AZURE_STORAGE_CONNECTION_STRING\"]\n+BLOB_CONTAINER_NAME = os.environ[\"BLOB_CHECKPOINT_CONTAINER\"]\n+CONSUMER_GROUP = os.getenv(\"EVENT_HUB_CONSUMER_GROUP\", \"$Default\")\n+\n+\n+async def send_events() -\u003e None:\n+    producer = EventHubProducerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        eventhub_name=EVENT_HUB_NAME,\n+    )\n+\n+    async with producer:\n+        batch = await producer.create_batch()\n+        for index in range(10):\n+            batch.add(EventData(f\"Event {index}\"))\n+        await producer.send_batch(batch)\n+        print(f\"Sent {len(batch)} events\")\n+\n+\n+async def on_event(partition_context, event) -\u003e None:\n+    print(\n+        f\"Partition {partition_context.partition_id}: \"\n+        f\"{event.body_as_str(encoding='UTF-8')}\"\n+    )\n+    await partition_context.update_checkpoint(event)\n+\n+\n+async def on_error(partition_context, error: Exception) -\u003e None:\n+    partition_id: Optional[str] = (\n+        partition_context.partition_id if partition_context else None\n+    )\n+    if partition_id is None:\n+        print(f\"Consumer error: {error!r}\")\n+    else:\n+        print(f\"Consumer error on partition {partition_id}: {error!r}\")\n+\n+\n+async def receive_events() -\u003e None:\n+    checkpoint_store = BlobCheckpointStore.from_connection_string(\n+        STORAGE_CONNECTION_STRING,\n+        BLOB_CONTAINER_NAME,\n+    )\n+    consumer = EventHubConsumerClient.from_connection_string(\n+        conn_str=EVENT_HUB_CONNECTION_STRING,\n+        consumer_group=CONSUMER_GROUP,\n+        eventhub_name=EVENT_HUB_NAME,\n+        checkpoint_store=checkpoint_store,\n+    )\n+\n+    async with consumer:\n+        print(\"Receiving events. Press Ctrl+C to stop.\")\n+        await consumer.receive(\n+            on_event=on_event,\n+            on_error=on_error,\n+            starting_position=\"-1\",\n+        )\n+\n+\n+async def main() -\u003e None:\n+    await send_events()\n+    await receive_events()\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        asyncio.run(main())\n+    except KeyboardInterrupt:\n+        print(\"Receiver stopped.\")\n*** Add File: requirements.txt\n+azure-eventhub\n+azure-eventhub-checkpointstoreblob-aio\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2120679032/event_hubs_async.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2120679032/event_hubs_async.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-python-streaming-python-azure-skills-azure-skill-mcp-2120679032/event_hubs_async.py
@@ -1,0 +1,80 @@
+import asyncio
+import os
+from typing import Optional
+
+from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
+from azure.eventhub import EventData
+from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
+
+
+EVENT_HUB_CONNECTION_STRING = os.environ["EVENT_HUB_CONNECTION_STRING"]
+EVENT_HUB_NAME = os.environ["EVENT_HUB_NAME"]
+STORAGE_CONNECTION_STRING = os.environ["AZURE_STORAGE_CONNECTION_STRING"]
+BLOB_CONTAINER_NAME = os.environ["BLOB_CHECKPOINT_CONTAINER"]
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
+    print(
+        f"Partition {partition_context.partition_id}: "
+        f"{event.body_as_str(encoding='UTF-8')}"
+    )
+    await partition_context.update_checkpoint(event)
+
+
+async def on_error(partition_context, error: Exception) -> None:
+    partition_id: Optional[str] = (
+        partition_context.partition_id if partition_context else None
+    )
+    if partition_id is None:
+
... (truncated)
```

### ✅ 8. powershell — 5663ms

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

I’m checking the local project shape and the SDK-specific acceptance guidance before creating the async sample.

The sample will use async clients end-to-end, environment-only secrets, batched publishing, per-event checkpoint updates, and a receiver error callback.

Created:

- `event_hubs_async.py` — async batch sending, `EventHubConsumerClient`, `BlobCheckpointStore`, event checkpointing, and error handling.
- `requirements.txt` — includes `azure-eventhub` and `azure-eventhub-checkpointstoreblob-aio`.

Configuration uses environment variables: `EVENT_HUB_CONNECTION_STRING`, `EVENT_HUB_NAME`, `AZURE_STORAGE_CONNECTION_STRING`, `BLOB_CHECKPOINT_CONTAINER`, and optional `EVENT_HUB_CONSUMER_GROUP`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- The asynchronous on_error callback reports consumer errors with partition context when available.

### Issues

- Azure SDK HttpResponseError exceptions are not explicitly caught and handled.

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
hyoka run --prompt-id event-hubs-dp-python-streaming --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
