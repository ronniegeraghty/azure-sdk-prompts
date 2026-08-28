# Evaluation Report: service-bus-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 470.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 12/18 |
| Duration | 470.9s |
| Timestamp | 2026-08-28T03:08:28Z |
| Files Generated | 4 |
| Event Count | 4103 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 302.0s |
| Review | 167.6s |
| **Total** | **470.9s** |

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
| Token Usage | in=298066 out=3929 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 773ms

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

### ✅ 2. glob — 1718ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 1355ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 182793ms

**Input:**

```json
{"intent":"Generate a local-only TypeScript sample using @azure/service-bus that sends to queues and topics, receives, completes messages, subscribes with handlers, and closes all resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 182794ms

**Input:**

```json
{"intent":"Find current official TypeScript/JavaScript @azure/service-bus examples for ServiceBusClient, createSender, createMessageBatch, tryAddMessage, receiveMessages, completeMessage, subscribe, topic subscriptions, and cleanup.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. apply_patch — 1061ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-service-bus-typescript-demo\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/service-bus\": \"^7.9.5\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^22.10.2\",\n+    \"typescript\": \"^5.7.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+SERVICE_BUS_CONNECTION_STRING=Endpoint=sb://your-namespace.servicebus.windows.net/;SharedAccessKeyName=your-policy;SharedAccessKey=your-key\n+SERVICE_BUS_QUEUE_NAME=demo-queue\n+SERVICE_BUS_TOPIC_NAME=demo-topic\n+SERVICE_BUS_SUBSCRIPTION_NAME=demo-subscription\n*** Add File: src/index.ts\n+import {\n+  ServiceBusClient,\n+  type ProcessErrorArgs,\n+  type ServiceBusReceivedMessage,\n+} from \"@azure/service-bus\";\n+\n+const connectionString = requireEnvironmentVariable(\n+  \"SERVICE_BUS_CONNECTION_STRING\",\n+);\n+const queueName = requireEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+const topicName = requireEnvironmentVariable(\"SERVICE_BUS_TOPIC_NAME\");\n+const subscriptionName = requireEnvironmentVariable(\n+  \"SERVICE_BUS_SUBSCRIPTION_NAME\",\n+);\n+\n+const client = new ServiceBusClient(connectionString);\n+const queueSender = client.createSender(queueName);\n+const queueReceiver = client.createReceiver(queueName, {\n+  receiveMode: \"peekLock\",\n+});\n+const topicSender = client.createSender(topicName);\n+const topicSubscriptionReceiver = client.createReceiver(\n+  topicName,\n+  subscriptionName,\n+  { receiveMode: \"peekLock\" },\n+);\n+\n+let closeQueueSubscription: (() =\u003e Promise\u003cvoid\u003e) | undefined;\n+\n+try {\n+  await queueSender.sendMessages({\n+    body: \"Single queue message\",\n+    contentType: \"text/plain\",\n+  });\n+  console.log(\"Sent one message to the queue.\");\n+\n+  const batch = await queueSender.createMessageBatch();\n+  for (let sequence = 1; sequence \u003c= 5; sequence += 1) {\n+    const added = batch.tryAddMessage({\n+      body: { sequence, text: `Batch message ${sequence}` },\n+      contentType: \"application/json\",\n+    });\n+\n+    if (!added) {\n+      throw new Error(`Batch message ${sequence} exceeded the batch size limit.`);\n+    }\n+  }\n+  await queueSender.sendMessages(batch);\n+  console.log(\"Sent a batch of 5 messages to the queue.\");\n+\n+  const receivedMessages = await queueReceiver.receiveMessages(6, {\n+    maxWaitTimeInMs: 10_000,\n+  });\n+\n+  for (const message of receivedMessages) {\n+    await processMessage(message);\n+    await queueReceiver.completeMessage(message);\n+  }\n+  console.log(`Received and completed ${receivedMessages.length} queue messages.`);\n+\n+  const subscribedMessageProcessed = new Promise\u003cvoid\u003e((resolve) =\u003e {\n+    const subscription = queueReceiver.subscribe({\n+      processMessage: async (message: ServiceBusReceivedMessage) =\u003e {\n+        await processMessage(message);\n+        await queueReceiver.completeMessage(message);\n+        resolve();\n+      },\n+      processError: async (args: ProcessErrorArgs) =\u003e {\n+        console.error(\n+          `Queue subscription error (${args.errorSource}) in ${args.entityPath}:`,\n+          args.error,\n+        );\n+      },\n+    });\n+\n+    closeQueueSubscription = () =\u003e subscription.close();\n+  });\n+\n+  await queueSender.sendMessages({\n+    body: \"Message handled by subscribe()\",\n+    contentType: \"text/plain\",\n+  });\n+  await waitWithTimeout(\n+    subscribedMessageProcessed,\n+    15_000,\n+    \"Timed out waiting for the subscribe() handler.\",\n+  );\n+  console.log(\"Processed and completed one message using subscribe().\");\n+\n+  await topicSender.sendMessages({\n+    body: \"Topic message\",\n+    subject: \"topic-demo\",\n+    contentType: \"text/plain\",\n+  });\n+\n+  const topicMessages = await topicSubscriptionReceiver.receiveMessages(1, {\n+    maxWaitTimeInMs: 10_000,\n+  });\n+  if (topicMessages.length === 0) {\n+    throw new Error(\"No message was received from the topic subscription.\");\n+  }\n+\n+  for (const message of topicMessages) {\n+    await processMessage(message);\n+    await topicSubscriptionReceiver.completeMessage(message);\n+  }\n+  console.log(\"Sent to a topic and received from its subscription.\");\n+} finally {\n+  const closeOperations: Promise\u003cvoid\u003e[] = [\n+    queueReceiver.close(),\n+    topicSubscriptionReceiver.close(),\n+    queueSender.close(),\n+    topicSender.close(),\n+    client.close(),\n+  ];\n+\n+  if (closeQueueSubscription) {\n+    closeOperations.unshift(closeQueueSubscription());\n+  }\n+\n+  await Promise.all(closeOperations);\n+}\n+\n+async function processMessage(\n+  message: ServiceBusReceivedMessage,\n+): Promise\u003cvoid\u003e {\n+  console.log(`Processing message ${message.messageId}:`, message.body);\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+function waitWithTimeout(\n+  operation: Promise\u003cvoid\u003e,\n+  timeoutInMs: number,\n+  timeoutMessage: string,\n+): Promise\u003cvoid\u003e {\n+  return new Promise((resolve, reject) =\u003e {\n+    const timeout = setTimeout(() =\u003e reject(new Error(timeoutMessage)), timeoutInMs);\n+\n+    operation.then(\n+      () =\u003e {\n+        clearTimeout(timeout);\n+        resolve();\n+      },\n+      (error: unknown) =\u003e {\n+        clearTimeout(timeout);\n+        reject(error);\n+      },\n+    );\n+  });\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/package.json
@@ -1,0 +1,18 @@
+{
+  "name": "azure-service-bus-typescript-demo",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "node dist/index.js"
+  },
+  "dependencies": {
+    "@azure/service-bus": "^7.9.5"
+  },
+  "devDependencies": {
+    "@types/node": "^22.10.2",
+    "typescript": "^5.7.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/tsconfig.json
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



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/.env.example b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/.env.example
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js
... (truncated)
```

### ✅ 7. powershell — 29702ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund --quiet \u0026\u0026 npm run build --silent","description":"Install dependencies and build TypeScript sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 76 packages in 24s
<shellId: 0 completed with exit code 0>
```

### ✅ 8. apply_patch — 882ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n } finally {\n-  const closeOperations: Promise\u003cvoid\u003e[] = [\n-    queueReceiver.close(),\n-    topicSubscriptionReceiver.close(),\n-    queueSender.close(),\n-    topicSender.close(),\n-    client.close(),\n-  ];\n-\n-  if (closeQueueSubscription) {\n-    closeOperations.unshift(closeQueueSubscription());\n-  }\n-\n-  await Promise.all(closeOperations);\n+  try {\n+    if (closeQueueSubscription) {\n+      await closeQueueSubscription();\n+    }\n+  } finally {\n+    try {\n+      await Promise.all([\n+        queueReceiver.close(),\n+        topicSubscriptionReceiver.close(),\n+        queueSender.close(),\n+        topicSender.close(),\n+      ]);\n+    } finally {\n+      await client.close();\n+    }\n+  }\n }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-1995963970/src/index.ts
@@ -106,19 +106,22 @@
   }
   console.log("Sent to a topic and received from its subscription.");
 } finally {
-  const closeOperations: Promise<void>[] = [
-    queueReceiver.close(),
-    topicSubscriptionReceiver.close(),
-    queueSender.close(),
-    topicSender.close(),
-    client.close(),
-  ];
-
-  if (closeQueueSubscription) {
-    closeOperations.unshift(closeQueueSubscription());
+  try {
+    if (closeQueueSubscription) {
+      await closeQueueSubscription();
+    }
+  } finally {
+    try {
+      await Promise.all([
+        queueReceiver.close(),
+        topicSubscriptionReceiver.close(),
+        queueSender.close(),
+        topicSender.close(),
+      ]);
+    } finally {
+      await client.close();
+    }
   }
-
-  await Promise.all(closeOperations);
 }
 
 async function processMessage(

```

### ✅ 9. powershell — 11861ms

**Input:**

```json
{"command":"npm run build --silent","description":"Rebuild TypeScript sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the project shape and the Service Bus SDK guidance, then add a locally runnable sample without contacting Azure.

The workspace is empty, so I’ll create a minimal TypeScript project with environment-based placeholders and deterministic queue/topic cleanup.

Created a complete TypeScript Azure Service Bus sample in `src\index.ts` demonstrating:

- Queue single-message and five-message batch sending
- `receiveMessages()` with peek-lock completion
- `subscribe()` with `processMessage` and `processError`
- Topic sending and subscription receiving
- Environment-based connection configuration
- Ordered cleanup of subscriptions, receivers, senders, and `ServiceBusClient`

The required `@azure/service-bus` package, TypeScript configuration, scripts, and `.env.example` are included. Run with:

```powershell
$env:SERVICE_BUS_CONNECTION_STRING="..."
$env:SERVICE_BUS_QUEUE_NAME="demo-queue"
$env:SERVICE_BUS_TOPIC_NAME="demo-topic"
$env:SERVICE_BUS_SUBSCRIPTION_NAME="demo-subscription"
npm start
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Includes the required Azure Service Bus dependency and complete TypeScript tooling configuration.

## Grader Results

- send-receive-messages.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `@azure/service-bus` npm package: Pass
      - `ServiceBusClient` constructor with connection string: Pass
      - `createSender()` for queue or topic: Pass
      - `ServiceBusMessageBatch` with `tryAddMessage()`: Pass
      - `createReceiver()` for queue or subscription: Pass
      - `receiveMessages()` for batch receive and `subscribe()` for streaming: Pass
      - `completeMessage()`, `abandonMessage()`, `deadLetterMessage()`: Fail
      - `close()` on sender, receiver, and client for cleanup: Pass
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
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 14.9% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 17.0% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 17.0% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 17.0% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 17.0% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 17.0% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 5.8750** | **53.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
