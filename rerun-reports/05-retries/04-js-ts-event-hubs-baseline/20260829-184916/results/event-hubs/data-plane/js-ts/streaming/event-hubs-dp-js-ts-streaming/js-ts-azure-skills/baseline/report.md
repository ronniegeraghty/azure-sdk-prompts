# Evaluation Report: event-hubs-dp-js-ts-streaming

**Config:** js-ts-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 265.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-js-ts-streaming` |
| Config | js-ts-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 14/18 |
| Duration | 265.7s |
| Timestamp | 2026-08-29T10:49:16Z |
| Files Generated | 4 |
| Event Count | 4131 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 103.4s |
| Review | 160.9s |
| **Total** | **265.7s** |

## Configuration

- **name:** js-ts-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=131661 out=3640 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. glob — 11657ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 10553ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"**/package.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 9545ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"**/tsconfig.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. apply_patch — 1896ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-event-hubs-typescript-example\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/event-hubs\": \"^6.0.0\",\n+    \"@azure/eventhubs-checkpointstore-blob\": \"^1.0.0\",\n+    \"@azure/storage-blob\": \"^12.0.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^22.0.0\",\n+    \"typescript\": \"^5.0.0\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"Node16\",\n+    \"moduleResolution\": \"Node16\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import {\n+  EventHubConsumerClient,\n+  EventHubProducerClient,\n+  earliestEventPosition,\n+  type Subscription,\n+} from \"@azure/event-hubs\";\n+import { BlobCheckpointStore } from \"@azure/eventhubs-checkpointstore-blob\";\n+import { BlobServiceClient } from \"@azure/storage-blob\";\n+import { randomUUID } from \"node:crypto\";\n+\n+const eventHubsConnectionString = requireEnvironmentVariable(\n+  \"EVENT_HUB_CONNECTION_STRING\",\n+);\n+const eventHubName = requireEnvironmentVariable(\"EVENT_HUB_NAME\");\n+const storageConnectionString = requireEnvironmentVariable(\n+  \"AZURE_STORAGE_CONNECTION_STRING\",\n+);\n+const checkpointContainerName = requireEnvironmentVariable(\n+  \"CHECKPOINT_CONTAINER_NAME\",\n+);\n+const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? \"$Default\";\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Required environment variable ${name} is not set.`);\n+  }\n+  return value;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const producer = new EventHubProducerClient(\n+    eventHubsConnectionString,\n+    eventHubName,\n+  );\n+  let consumer: EventHubConsumerClient | undefined;\n+  let subscription: Subscription | undefined;\n+  let resolveStop: (() =\u003e void) | undefined;\n+  const stopRequested = new Promise\u003cvoid\u003e((resolve) =\u003e {\n+    resolveStop = resolve;\n+  });\n+  const requestStop = (signal: NodeJS.Signals): void =\u003e {\n+    console.log(`Received ${signal}; shutting down.`);\n+    resolveStop?.();\n+  };\n+\n+  process.once(\"SIGINT\", requestStop);\n+  process.once(\"SIGTERM\", requestStop);\n+\n+  try {\n+    const runId = randomUUID();\n+    const batch = await producer.createBatch();\n+\n+    for (let index = 1; index \u003c= 10; index += 1) {\n+      const added = batch.tryAdd({\n+        body: {\n+          message: `Event ${index}`,\n+          createdAt: new Date().toISOString(),\n+        },\n+        properties: {\n+          runId,\n+          eventNumber: index,\n+          source: \"typescript-event-hubs-demo\",\n+        },\n+      });\n+\n+      if (!added) {\n+        throw new Error(`Event ${index} is too large for the Event Hubs batch.`);\n+      }\n+    }\n+\n+    await producer.sendBatch(batch);\n+    console.log(`Sent ${batch.count} events for run ${runId}.`);\n+\n+    const blobServiceClient = BlobServiceClient.fromConnectionString(\n+      storageConnectionString,\n+    );\n+    const containerClient =\n+      blobServiceClient.getContainerClient(checkpointContainerName);\n+    const checkpointStore = new BlobCheckpointStore(containerClient);\n+\n+    consumer = new EventHubConsumerClient(\n+      consumerGroup,\n+      eventHubsConnectionString,\n+      eventHubName,\n+      checkpointStore,\n+    );\n+\n+    const receivedEventNumbers = new Set\u003cnumber\u003e();\n+    subscription = consumer.subscribe(\n+      {\n+        processEvents: async (events, context) =\u003e {\n+          for (const event of events) {\n+            console.log(\n+              `Partition ${context.partitionId} received:`,\n+              event.body,\n+            );\n+\n+            if (\n+              event.properties?.runId === runId \u0026\u0026\n+              typeof event.properties.eventNumber === \"number\"\n+            ) {\n+              receivedEventNumbers.add(event.properties.eventNumber);\n+            }\n+\n+            await context.updateCheckpoint(event);\n+          }\n+\n+          if (receivedEventNumbers.size === 10) {\n+            console.log(\"Received and checkpointed all 10 demo events.\");\n+            resolveStop?.();\n+          }\n+        },\n+        processError: async (error, context) =\u003e {\n+          console.error(\n+            `Error on partition ${context.partitionId ?? \"unknown\"}:`,\n+            error,\n+          );\n+        },\n+      },\n+      {\n+        startPosition: earliestEventPosition,\n+      },\n+    );\n+\n+    console.log(\"Listening for events. Press Ctrl+C to stop.\");\n+    await stopRequested;\n+  } finally {\n+    process.removeListener(\"SIGINT\", requestStop);\n+    process.removeListener(\"SIGTERM\", requestStop);\n+    await subscription?.close();\n+    await consumer?.close();\n+    await producer.close();\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Event Hubs example failed:\", error);\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "azure-event-hubs-typescript-example",
+  "version": "1.0.0",
+  "private": true,
+  "scripts": {
+    "build": "tsc",
+    "start": "node dist/index.js"
+  },
+  "dependencies": {
+    "@azure/event-hubs": "^6.0.0",
+    "@azure/eventhubs-checkpointstore-blob": "^1.0.0",
+    "@azure/storage-blob": "^12.0.0"
+  },
+  "devDependencies": {
+    "@types/node": "^22.0.0",
+    "typescript": "^5.0.0"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/tsconfig.json
@@ -1,0 +1,14 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "Node16",
+    "moduleResolution": "Node16",
+    "outDir": "dist",
+    "rootDir": "src",
+    "strict": true,
+    "esModuleInterop": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/src/index.ts
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/T
... (truncated)
```

### ✅ 5. powershell — 26309ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund --quiet \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
npm warn cleanup Failed to remove some directories [
npm warn cleanup   [
npm warn cleanup     '\\\\?\\C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273\\node_modules\\@azure',
npm warn cleanup     [Error: EPERM: operation not permitted, rmdir 'C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273\node_modules\@azure\storage-blob\dist\commonjs'] {
npm warn cleanup       errno: -4048,
npm warn cleanup       code: 'EPERM',
npm warn cleanup       syscall: 'rmdir',
npm warn cleanup       path: 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273\\node_modules\\@azure\\storage-blob\\dist\\commonjs'
npm warn cleanup     }
npm warn cleanup   ]
npm warn cleanup ]
npm error code E401
npm error Unable to authenticate, your authentication token seems to be invalid.
npm error To correct this please try logging in again with:
npm error   npm login
npm error A complete log of this run can be found in: C:\Users\weidxu\AppData\Local\npm-cache\_logs\2026-08-29T10_50_07_499Z-debug-0.log
<shellId: 0 completed with exit code 1>
```

### ✅ 6. glob — 9758ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"node_modules/typescript/bin/tsc"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. glob — 8585ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"node_modules/@azure/event-hubs/package.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 8. glob — 7500ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"node_modules/@azure/eventhubs-checkpointstore-blob/package.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. glob — 6505ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273","pattern":"package-lock.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 10. apply_patch — 2342ms

**Input:**

```json
"*** Begin Patch\n*** Add File: README.md\n+# Azure Event Hubs TypeScript example\n+\n+This example sends a batch of 10 events, receives them with a\n+checkpoint-aware consumer, and closes all clients on completion or when\n+`Ctrl+C` is pressed.\n+\n+## Install\n+\n+```powershell\n+npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob\n+npm install --save-dev typescript @types/node\n+```\n+\n+`@azure/event-hubs` provides the producer and consumer clients.\n+`@azure/eventhubs-checkpointstore-blob` provides `BlobCheckpointStore`, and\n+`@azure/storage-blob` provides the container client used by that store.\n+\n+## Configure and run\n+\n+Set these values to your own resources. The Blob container must already\n+exist; the example does not create Azure resources.\n+\n+```powershell\n+$env:EVENT_HUB_CONNECTION_STRING = \"\u003cevent-hubs-namespace-connection-string\u003e\"\n+$env:EVENT_HUB_NAME = \"\u003cevent-hub-name\u003e\"\n+$env:AZURE_STORAGE_CONNECTION_STRING = \"\u003cstorage-account-connection-string\u003e\"\n+$env:CHECKPOINT_CONTAINER_NAME = \"\u003cexisting-blob-container-name\u003e\"\n+$env:EVENT_HUB_CONSUMER_GROUP = '$Default' # Optional\n+\n+npm run build\n+npm start\n+```\n+\n+The consumer starts at its stored checkpoint. If no checkpoint exists, it\n+starts at the earliest retained event so it can receive the batch that this\n+program sends before subscribing.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/README.md b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/README.md
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-baseline-430905273/README.md
@@ -1,0 +1,37 @@
+# Azure Event Hubs TypeScript example
+
+This example sends a batch of 10 events, receives them with a
+checkpoint-aware consumer, and closes all clients on completion or when
+`Ctrl+C` is pressed.
+
+## Install
+
+```powershell
+npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
+npm install --save-dev typescript @types/node
+```
+
+`@azure/event-hubs` provides the producer and consumer clients.
+`@azure/eventhubs-checkpointstore-blob` provides `BlobCheckpointStore`, and
+`@azure/storage-blob` provides the container client used by that store.
+
+## Configure and run
+
+Set these values to your own resources. The Blob container must already
+exist; the example does not create Azure resources.
+
+```powershell
+$env:EVENT_HUB_CONNECTION_STRING = "<event-hubs-namespace-connection-string>"
+$env:EVENT_HUB_NAME = "<event-hub-name>"
+$env:AZURE_STORAGE_CONNECTION_STRING = "<storage-account-connection-string>"
+$env:CHECKPOINT_CONTAINER_NAME = "<existing-blob-container-name>"
+$env:EVENT_HUB_CONSUMER_GROUP = '$Default' # Optional
+
+npm run build
+npm start
+```
+
+The consumer starts at its stored checkpoint. If no checkpoint exists, it
+starts at the earliest retained event so it can receive the batch that this
+program sends before subscribing.
+


```

## Generated Files

- `README.md`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace, then add a focused TypeScript example and the required package metadata without contacting Azure.

The workspace is empty. I’ll create a minimal project whose consumer starts from the earliest available event so a send-before-subscribe demo can still receive its newly sent batch.

The example is implemented with batch-size checks, per-event checkpoints, automatic completion after its 10 events, and signal-driven cleanup. I’ll install the declared dependencies and compile it now.

Dependency installation is blocked by the environment’s invalid npm authentication token. I’ll inspect the partial install and compile if the required packages and compiler are already available.

The source cannot be compiled in this environment because none of the dependencies were downloaded. I’ll add concise local setup instructions, including the required packages, placeholder environment variables, and the pre-existing checkpoint-container requirement.

Implemented the TypeScript Azure Event Hubs example:

- `src/index.ts` — sends 10 batched events with custom properties, consumes them using `BlobCheckpointStore`, updates checkpoints, and handles graceful shutdown.
- `package.json` / `tsconfig.json` — required Azure SDK and TypeScript dependencies.
- `README.md` — installation, environment variables, and execution instructions.

The local dependency installation was blocked by an invalid npm authentication token (`E401`), so compilation could not complete in this environment. No real Azure resources were accessed or modified.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criterion passed.

### Strengths

- Complete dependency manifest for the Event Hubs and BlobCheckpointStore implementation.
- Includes the TypeScript compiler, Node.js type definitions, build scripts, and tsconfig.json.

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
  - Async/Await Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 7.0000** | **63.6%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-js-ts-streaming --config js-ts-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
