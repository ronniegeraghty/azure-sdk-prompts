# Evaluation Report: event-hubs-dp-js-ts-streaming

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 537.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-js-ts-streaming` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 13/18 |
| Duration | 537.9s |
| Timestamp | 2026-08-28T05:27:05Z |
| Files Generated | 3 |
| Event Count | 3449 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 372.9s |
| Review | 163.8s |
| **Total** | **537.9s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=294773 out=3074 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 1315ms

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

### ✅ 2. glob — 182839ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 182393ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724","pattern":"*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 182842ms

**Input:**

```json
{"intent":"Generate a local-only TypeScript example using Azure Event Hubs producer and consumer clients with BlobCheckpointStore, async/await, checkpointing, and graceful shutdown. Do not provision Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. apply_patch — 910ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-event-hubs-typescript-sample\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/event-hubs\": \"^5.12.2\",\n+    \"@azure/eventhubs-checkpointstore-blob\": \"^1.0.1\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"tsx\": \"^4.20.5\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+EVENT_HUB_CONNECTION_STRING=Endpoint=sb://\u003cnamespace\u003e.servicebus.windows.net/;SharedAccessKeyName=\u003ckey-name\u003e;SharedAccessKey=\u003ckey\u003e\n+EVENT_HUB_NAME=\u003cevent-hub-name\u003e\n+STORAGE_CONNECTION_STRING=DefaultEndpointsProtocol=https;AccountName=\u003cstorage-account\u003e;AccountKey=\u003ckey\u003e;EndpointSuffix=core.windows.net\n+BLOB_CONTAINER_NAME=\u003ccheckpoint-container-name\u003e\n+EVENT_HUB_CONSUMER_GROUP=$Default\n*** Add File: src/index.ts\n+import {\n+  EventDataBatch,\n+  EventHubConsumerClient,\n+  EventHubProducerClient,\n+  ReceivedEventData,\n+  Subscription,\n+} from \"@azure/event-hubs\";\n+import { BlobCheckpointStore } from \"@azure/eventhubs-checkpointstore-blob\";\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+async function addEvents(batch: EventDataBatch): Promise\u003cvoid\u003e {\n+  for (let index = 1; index \u003c= 10; index += 1) {\n+    const added = batch.tryAdd({\n+      body: {\n+        id: index,\n+        message: `Event ${index}`,\n+        createdAt: new Date().toISOString(),\n+      },\n+      properties: {\n+        eventType: \"sample.created\",\n+        sequenceNumber: index,\n+        source: \"typescript-sample\",\n+      },\n+    });\n+\n+    if (!added) {\n+      throw new Error(`Event ${index} did not fit in the Event Hubs batch.`);\n+    }\n+  }\n+}\n+\n+function waitForShutdownSignal(): Promise\u003cNodeJS.Signals\u003e {\n+  return new Promise((resolve) =\u003e {\n+    const signals: NodeJS.Signals[] = [\"SIGINT\", \"SIGTERM\"];\n+    for (const signal of signals) {\n+      process.once(signal, () =\u003e resolve(signal));\n+    }\n+  });\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const eventHubConnectionString = requireEnvironmentVariable(\n+    \"EVENT_HUB_CONNECTION_STRING\",\n+  );\n+  const eventHubName = requireEnvironmentVariable(\"EVENT_HUB_NAME\");\n+  const storageConnectionString = requireEnvironmentVariable(\n+    \"STORAGE_CONNECTION_STRING\",\n+  );\n+  const blobContainerName = requireEnvironmentVariable(\"BLOB_CONTAINER_NAME\");\n+  const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? \"$Default\";\n+\n+  const producer = new EventHubProducerClient(\n+    eventHubConnectionString,\n+    eventHubName,\n+  );\n+  const checkpointStore = new BlobCheckpointStore(\n+    storageConnectionString,\n+    blobContainerName,\n+  );\n+  const consumer = new EventHubConsumerClient(\n+    consumerGroup,\n+    eventHubConnectionString,\n+    eventHubName,\n+    checkpointStore,\n+  );\n+\n+  let subscription: Subscription | undefined;\n+\n+  try {\n+    const batch = await producer.createBatch();\n+    await addEvents(batch);\n+    await producer.sendBatch(batch);\n+    console.log(`Sent ${batch.count} events.`);\n+\n+    subscription = consumer.subscribe({\n+      processEvents: async (events, context) =\u003e {\n+        for (const event of events) {\n+          printEvent(event, context.partitionId);\n+        }\n+\n+        const lastEvent = events.at(-1);\n+        if (lastEvent) {\n+          await context.updateCheckpoint(lastEvent);\n+          console.log(\n+            `Updated checkpoint for partition ${context.partitionId} at sequence ${lastEvent.sequenceNumber}.`,\n+          );\n+        }\n+      },\n+      processError: async (error, context) =\u003e {\n+        console.error(\n+          `Error while processing partition ${context.partitionId}:`,\n+          error,\n+        );\n+      },\n+    });\n+\n+    console.log(\"Receiving events. Press Ctrl+C to stop.\");\n+    const signal = await waitForShutdownSignal();\n+    console.log(`Received ${signal}; shutting down.`);\n+  } finally {\n+    await subscription?.close();\n+    await consumer.close();\n+    await producer.close();\n+  }\n+}\n+\n+function printEvent(event: ReceivedEventData, partitionId: string): void {\n+  console.log(\n+    `Partition ${partitionId}, sequence ${event.sequenceNumber}:`,\n+    event.body,\n+  );\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Event Hubs sample failed:\", error);\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "azure-event-hubs-typescript-sample",
+  "version": "1.0.0",
+  "private": true,
+  "scripts": {
+    "build": "tsc",
+    "start": "tsx src/index.ts"
+  },
+  "dependencies": {
+    "@azure/event-hubs": "^5.12.2",
+    "@azure/eventhubs-checkpointstore-blob": "^1.0.1"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "tsx": "^4.20.5",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/tsconfig.json
@@ -1,0 +1,14 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "outDir": "dist",
+    "rootDir": "src",
+    "strict": true,
+    "esModuleInterop": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/.env.example b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724/.env.example
create file mode 100644
index 0000000..0000000
--
... (truncated)
```

### ✅ 6. powershell — 58188ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund --silent \u0026\u0026 npm run build --silent","description":"Install dependencies and compile TypeScript sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 1>
```

### ✅ 7. powershell — 3210ms

**Input:**

```json
{"command":"npm run build","description":"Show TypeScript compilation errors","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-event-hubs-typescript-sample@1.0.0 build
> tsc

'tsc' is not recognized as an internal or external command,
operable program or batch file.
<shellId: 1 completed with exit code 1>
```

### ✅ 8. powershell — 38830ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund","description":"Install sample npm dependencies","initial_wait":180,"mode":"sync"}
```

**Output:**

```
npm error code E401
npm error Unable to authenticate, your authentication token seems to be invalid.
npm error To correct this please try logging in again with:
npm error   npm login
npm error A complete log of this run can be found in: C:\Users\weidxu\AppData\Local\npm-cache\_logs\2026-08-28T05_32_25_837Z-debug-0.log
<shellId: 2 completed with exit code 1>
```

### ✅ 9. glob — 2520ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724","pattern":"package-lock.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 10. glob — 1963ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-1978235724","pattern":"node_modules\\**\\package.json"}
```

**Output:**

```
No files matched the pattern.
```

## Generated Files

- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll scaffold a local-only TypeScript sample with environment-variable placeholders, package metadata, and graceful signal handling.

The sample will fail fast on missing configuration or an oversized event, keep the subscription alive until `SIGINT`/`SIGTERM`, and close the subscription and both clients exactly once.

The implementation is complete; I’m compiling it against the declared Azure SDK versions to catch API or type mismatches.

Implemented the TypeScript Event Hubs sample in `src\index.ts`, including batching 10 events with custom properties, Blob checkpointing, `subscribe()` handlers, checkpoint updates, and graceful `SIGINT`/`SIGTERM` shutdown.

Required packages are declared in `package.json`:

```powershell
npm install
npm start
```

Configure values from `.env.example` as environment variables before running. TypeScript compilation could not be completed because the configured npm registry rejected the local authentication token with `E401`.

## Reviewer Notes (LLM-as-Judge)

### Summary

The dependency and TypeScript project configuration criterion is fully satisfied.

### Strengths

- Includes both required Azure SDK dependencies.
- Provides TypeScript build and execution tooling.
- Includes Node.js type definitions and a strict tsconfig.json.

## Grader Results

- send-receive-events.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - `@azure/event-hubs` and `@azure/eventhubs-checkpointstore-blob` npm packages: Pass
      - `EventHubProducerClient` constructor: Pass
      - `createBatch()` and `EventDataBatch.tryAdd()`: Pass
      - `sendBatch()` for publishing: Pass
      - `EventHubConsumerClient` with `BlobCheckpointStore`: Pass
      - `subscribe()` with `SubscriptionEventHandlers` (processEvents, processError): Pass
      - `updateCheckpoint()` in the processEvents handler: Pass
      - `close()` for cleanup: Pass
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct @azure/ Scoped Packages**: All Azure SDK dependencies use @azure/ scoped packages (e.g., @azure/storage-blob, @azure/keyvault-secrets, @azure/identity). No deprecated unscoped packages (azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-*).: Pass
  - @azure/identity for Authentication (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**@azure/identity for Authentication**: Includes @azure/identity as a dependency. For local development, DefaultAzureCredential is acceptable. For production code, uses a specific Entra token credential (e.g., ManagedIdentityCredential, ClientSecretCredential, WorkloadIdentityCredential) or configures DefaultAzureCredential with the AZURE_TOKEN_CREDENTIALS environment variable set to "prod" or a specific token credential. No hardcoded connection strings, account keys, SAS tokens, or secrets in code.: Fail
  - Client Constructor with Endpoint and Credential (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Constructor with Endpoint and Credential**: SDK clients constructed with endpoint URL and credential object, e.g. new BlobServiceClient(url, credential) or new SecretClient(vaultUrl, credential). No legacy constructor patterns or connection string-based construction when identity-based auth is available.: Fail
  - Async/Await Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Fail
  - Pagination with for-await-of (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Pass
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Pass
  - RestError Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**RestError Exception Handling**: Catches RestError (from @azure/core-rest-pipeline or the service package) with statusCode inspection for error-specific handling. Not just generic Error or try/catch without type checking.: Fail
  - No Deprecated Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated Packages**: No deprecated or legacy Azure packages in package.json: azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-js, @azure/ms-rest-nodeauth, @azure/ms-rest-azure-js. All dependencies use the current @azure/* track.: Pass
  - Logging via @azure/logger (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Logging via @azure/logger**: Uses @azure/logger with setLogLevel() or the AZURE_LOG_LEVEL environment variable for SDK diagnostic logging. No custom logging wrappers around SDK internals.: Fail
  - package.json with Correct Dependencies (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**package.json with Correct Dependencies**: Includes a valid package.json with all required @azure/* dependencies listed. TypeScript projects include tsconfig.json and appropriate @types/* packages or use a recent TypeScript version with built-in types.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 6.0000** | **54.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-js-ts-streaming --config js-ts-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
