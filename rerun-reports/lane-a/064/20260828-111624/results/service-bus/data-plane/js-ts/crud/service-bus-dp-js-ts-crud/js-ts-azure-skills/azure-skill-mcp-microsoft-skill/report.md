# Evaluation Report: service-bus-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 603.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 13/18 |
| Duration | 603.3s |
| Timestamp | 2026-08-28T03:16:24Z |
| Files Generated | 5 |
| Event Count | 4848 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 475.3s |
| Review | 126.8s |
| **Total** | **603.3s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-servicebus-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=333008 out=4053 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 799ms

**Input:**

```json
{"skill":"azure-servicebus-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Service Bus SDK for TypeScript

Enterprise messaging with queues, topics, and subscriptions.

## Installation

```bash
npm install @azure/service-bus @azure/identity
```

## Environment Variables

```bash
SERVICEBUS_NAMESPACE=<namespace>.servicebus.windows.net
SERVICEBUS_QUEUE_NAME=my-queue
SERVICEBUS_TOPIC_NAME=my-topic
SERVICEBUS_SUBSCRIPTION_NAME=my-subscription
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

```typescript
import { ServiceBusClient } from "@azure/service-bus";
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

const fullyQualifiedNamespace = process.env.SERVICEBUS_NAMESPACE!;
const client = new ServiceBusClient(fullyQualifiedNamespace, credential);
```

## Core Workflow

### Send Messages to Queue

```typescript
const sender = client.createSender("my-queue");

// Single message
await sender.sendMessages({
  body: { orderId: "12345", amount: 99.99 },
  contentType: "application/json",
});

// Batch messages
const batch = await sender.createMessageBatch();
batch.tryAddMessage({ body: "Message 1" });
batch.tryAddMessage({ body: "Message 2" });
await sender.sendMessages(batch);

await sender.close();
```

### Receive Messages from Queue

```typescript
const receiver = client.createReceiver("my-queue");

// Receive batch
const messages = await receiver.receiveMessages(10, { maxWaitTimeInMs: 5000 });
for (const message of messages) {
  console.log(`Received: 
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 181884ms

**Input:**

```json
{"intent":"Generate a local-only TypeScript example using @azure/service-bus with a connection string, queue single and batch sends, receiveMessages settlement, subscribe handlers, topics/subscriptions, and reliable close cleanup.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. glob — 181880ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-typescript\skills\azure-servicebus-ts\references\queues-topics.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-typescript\skills\azure-servicebus-ts\references\error-handling.md
```

### ✅ 4. glob — 181543ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-servicebus-ts","pattern":"**/{acceptance-criteria.md,examples.md,error-handling.md,queues-topics.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. view — 181632ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-servicebus-ts\\references\\error-handling.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/error-handling.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/error-handling.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/error-handling.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/error-handling.md
@@ -1,400 +1,400 @@
 # Error Handling and Reliability
 
 Comprehensive error handling patterns for @azure/service-bus.
 
 ## ServiceBusError
 
 All Service Bus errors extend `ServiceBusError` with a `code` property:
 
 ```typescript
 import { ServiceBusError } from "@azure/service-bus";
 
 try {
   await sender.sendMessages(message);
 } catch (error) {
   if (error instanceof ServiceBusError) {
     console.log(`Code: ${error.code}`);
     console.log(`Message: ${error.message}`);
     console.log(`Retryable: ${error.retryable}`);
     
     handleServiceBusError(error);
   }
 }
 ```
 
 ## Error Codes Reference
 
 | Code | Description | Retryable | Action |
 |------|-------------|-----------|--------|
 | `GeneralError` | Unspecified error | Maybe | Log and investigate |
 | `MessagingEntityNotFound` | Queue/topic/subscription doesn't exist | No | Check entity name, create entity |
 | `MessageLockLost` | Lock expired before settlement | No | Message will be redelivered |
 | `MessageNotFound` | Message no longer available | No | Already processed or expired |
 | `MessageSizeExceeded` | Message too large (>256KB standard, >100MB premium) | No | Reduce message size or use claim check |
 | `MessagingEntityAlreadyExists` | Entity already exists | No | Use existing entity |
 | `MessagingEntityDisabled` | Entity is di
... (truncated)
```

### ✅ 6. view — 181312ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-servicebus-ts\\references\\queues-topics.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/queues-topics.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/queues-topics.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/queues-topics.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-servicebus-ts/references/queues-topics.md
@@ -1,371 +1,371 @@
 # Queues vs Topics Patterns
 
 Detailed patterns for Azure Service Bus queues and topics with @azure/service-bus.
 
 ## Queue Patterns (Point-to-Point)
 
 Queues deliver each message to exactly one consumer. Use for work distribution.
 
 ### Basic Queue Sender
 
 ```typescript
 import { ServiceBusClient, ServiceBusMessage } from "@azure/service-bus";
 import { DefaultAzureCredential } from "@azure/identity";
 
 const client = new ServiceBusClient(
   process.env.SERVICEBUS_NAMESPACE!,
   new DefaultAzureCredential()
 );
 
 const sender = client.createSender("order-queue");
 
 // Send single message
 const message: ServiceBusMessage = {
   body: { orderId: "12345", amount: 99.99 },
   contentType: "application/json",
   messageId: "unique-id-12345",
   correlationId: "request-abc",
   applicationProperties: {
     priority: "high",
     source: "web-app",
   },
 };
 
 await sender.sendMessages(message);
 await sender.close();
 ```
 
 ### Batch Sending (Recommended for Multiple Messages)
 
 ```typescript
 const sender = client.createSender("order-queue");
 
 // Create batch with size limits
 const batch = await sender.createMessageBatch();
 
 const orders = [
   { orderId: "001", amount: 50 },
   { orderId: "002", amount: 75 },
   { orderId: "003", amount: 100 },
 ];
 
 for (const order of orde
... (truncated)
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 181635ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Get TypeScript Azure Service Bus SDK code-generation best practices for queue/topic send, batch, receive, subscribe, settlement, and cleanup.","parameters":{"language":"TypeScript","operation":"code-generation","resource":"Azure Service Bus"}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. apply_patch — 729ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-service-bus-typescript-demo\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc --noEmit\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/service-bus\": \"^7.9.5\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"tsx\": \"^4.20.5\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true,\n+    \"types\": [\"node\"]\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+SERVICEBUS_CONNECTION_STRING=Endpoint=sb://your-namespace.servicebus.windows.net/;SharedAccessKeyName=your-policy;SharedAccessKey=your-key\n+SERVICEBUS_QUEUE_NAME=demo-queue\n+SERVICEBUS_TOPIC_NAME=demo-topic\n+SERVICEBUS_SUBSCRIPTION_NAME=demo-subscription\n*** Add File: src/index.ts\n+import {\n+  type ProcessErrorArgs,\n+  ServiceBusClient,\n+  type ServiceBusReceiver,\n+  type ServiceBusSender,\n+  type ServiceBusSubscriber,\n+} from \"@azure/service-bus\";\n+\n+function requiredEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+function logSubscriptionError(args: ProcessErrorArgs): void {\n+  console.error(\"Service Bus subscription error\", {\n+    source: args.errorSource,\n+    entityPath: args.entityPath,\n+    namespace: args.fullyQualifiedNamespace,\n+    error: args.error,\n+  });\n+}\n+\n+async function run(): Promise\u003cvoid\u003e {\n+  const connectionString = requiredEnvironmentVariable(\n+    \"SERVICEBUS_CONNECTION_STRING\",\n+  );\n+  const queueName = requiredEnvironmentVariable(\"SERVICEBUS_QUEUE_NAME\");\n+  const topicName = requiredEnvironmentVariable(\"SERVICEBUS_TOPIC_NAME\");\n+  const topicSubscriptionName = requiredEnvironmentVariable(\n+    \"SERVICEBUS_SUBSCRIPTION_NAME\",\n+  );\n+\n+  const client = new ServiceBusClient(connectionString);\n+  let queueSender: ServiceBusSender | undefined;\n+  let pullReceiver: ServiceBusReceiver | undefined;\n+  let subscriptionReceiver: ServiceBusReceiver | undefined;\n+  let subscription: ServiceBusSubscriber | undefined;\n+  let topicSender: ServiceBusSender | undefined;\n+  let topicReceiver: ServiceBusReceiver | undefined;\n+\n+  try {\n+    queueSender = client.createSender(queueName);\n+\n+    await queueSender.sendMessages({\n+      body: { kind: \"single\", text: \"Hello from a single message\" },\n+      contentType: \"application/json\",\n+      messageId: `single-${Date.now()}`,\n+    });\n+    console.log(\"Sent one queue message\");\n+\n+    const batch = await queueSender.createMessageBatch();\n+    for (let number = 1; number \u003c= 5; number += 1) {\n+      const wasAdded = batch.tryAddMessage({\n+        body: { kind: \"batch\", number },\n+        contentType: \"application/json\",\n+        messageId: `batch-${Date.now()}-${number}`,\n+      });\n+\n+      if (!wasAdded) {\n+        throw new Error(`Batch message ${number} did not fit in the batch`);\n+      }\n+    }\n+    await queueSender.sendMessages(batch);\n+    console.log(`Sent a batch of ${batch.count} queue messages`);\n+\n+    pullReceiver = client.createReceiver(queueName, {\n+      receiveMode: \"peekLock\",\n+    });\n+    const pulledMessages = await pullReceiver.receiveMessages(3, {\n+      maxWaitTimeInMs: 10_000,\n+    });\n+\n+    for (const message of pulledMessages) {\n+      console.log(\"Pulled queue message:\", message.body);\n+      await pullReceiver.completeMessage(message);\n+    }\n+\n+    subscriptionReceiver = client.createReceiver(queueName, {\n+      receiveMode: \"peekLock\",\n+    });\n+\n+    let subscribedMessageCount = 0;\n+    let resolveSubscription: (() =\u003e void) | undefined;\n+    let rejectSubscription: ((error: unknown) =\u003e void) | undefined;\n+    const subscriptionFinished = new Promise\u003cvoid\u003e((resolve, reject) =\u003e {\n+      resolveSubscription = resolve;\n+      rejectSubscription = reject;\n+    });\n+\n+    subscription = subscriptionReceiver.subscribe(\n+      {\n+        processMessage: async (message) =\u003e {\n+          console.log(\"Subscribed queue message:\", message.body);\n+          await subscriptionReceiver?.completeMessage(message);\n+          subscribedMessageCount += 1;\n+\n+          if (subscribedMessageCount \u003e= 3) {\n+            resolveSubscription?.();\n+          }\n+        },\n+        processError: async (args) =\u003e {\n+          logSubscriptionError(args);\n+          rejectSubscription?.(args.error);\n+        },\n+      },\n+      {\n+        autoCompleteMessages: false,\n+        maxConcurrentCalls: 1,\n+      },\n+    );\n+\n+    const timeout = new Promise\u003cnever\u003e((_, reject) =\u003e {\n+      setTimeout(\n+        () =\u003e reject(new Error(\"Timed out waiting for subscribed messages\")),\n+        15_000,\n+      );\n+    });\n+    await Promise.race([subscriptionFinished, timeout]);\n+    await subscription.close();\n+    subscription = undefined;\n+\n+    topicSender = client.createSender(topicName);\n+    await topicSender.sendMessages({\n+      body: { event: \"order.created\", orderId: \"order-123\" },\n+      contentType: \"application/json\",\n+      applicationProperties: { eventType: \"order.created\" },\n+      messageId: `topic-${Date.now()}`,\n+    });\n+    console.log(\"Sent one topic message\");\n+\n+    topicReceiver = client.createReceiver(topicName, topicSubscriptionName, {\n+      receiveMode: \"peekLock\",\n+    });\n+    const topicMessages = await topicReceiver.receiveMessages(1, {\n+      maxWaitTimeInMs: 10_000,\n+    });\n+\n+    for (const message of topicMessages) {\n+      console.log(\"Received topic subscription message:\", message.body);\n+      await topicReceiver.completeMessage(message);\n+    }\n+  } finally {\n+    try {\n+      await subscription?.close();\n+    } finally {\n+      try {\n+        await Promise.all([\n+          topicReceiver?.close(),\n+          topicSender?.close(),\n+          subscriptionReceiver?.close(),\n+          pullReceiver?.close(),\n+          queueSender?.close(),\n+        ]);\n+      } finally {\n+        await client.close();\n+      }\n+    }\n+  }\n+}\n+\n+run().catch((error: unknown) =\u003e {\n+  console.error(\"Service Bus demo failed:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Service Bus TypeScript demo\n+\n+Install the required Service Bus SDK and development dependencies:\n+\n+```powershell\n+npm install\n+```\n+\n+The Azure messaging package used by the program is:\n+\n+```powershell\n+npm install @azure/service-bus\n+```\n+\n+Create the queue, topic, and subscription in an Azure Service Bus namespace\n+before running the example. Copy `.env.example` to `.env`, replace the\n+placeholders, load those variables into the shell, and run:\n+\n+```powershell\n+npm start\n+```\n+\n+The connection string must allow sending to the queue and topic and receiving\n+from the queue and topic subscription. Do not commit a real connection string.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "azure-service-bus-typescript-demo",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc --noEmit",
+    "start": "tsx src/index.ts"
+  },
+  "dependencies": {
+    "@azure/service-bus": "^7.9.5"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "tsx": "^4.20.5",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/tsconfig.json
@@ -1,0 +1,14 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "esModuleInterop": true,
+    "skipLibCheck": true,
+    "types": ["node"]
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/.env.example b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-mic
... (truncated)
```

### ✅ 9. powershell — 26887ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund --quiet \u0026\u0026 npm run build","description":"Install dependencies and type-check the demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 79 packages in 23s

> azure-service-bus-typescript-demo@1.0.0 build
> tsc --noEmit

src/index.ts(6,8): error TS2305: Module '"@azure/service-bus"' has no exported member 'ServiceBusSubscriber'.
<shellId: 0 completed with exit code 2>
```

### ✅ 10. apply_patch — 656ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n   type ProcessErrorArgs,\n   ServiceBusClient,\n   type ServiceBusReceiver,\n   type ServiceBusSender,\n-  type ServiceBusSubscriber,\n } from \"@azure/service-bus\";\n \n+type ServiceBusSubscription = ReturnType\u003cServiceBusReceiver[\"subscribe\"]\u003e;\n+\n@@\n-  let subscription: ServiceBusSubscriber | undefined;\n+  let subscription: ServiceBusSubscription | undefined;\n@@\n-    const timeout = new Promise\u003cnever\u003e((_, reject) =\u003e {\n-      setTimeout(\n+    let timeoutId: NodeJS.Timeout | undefined;\n+    const timeout = new Promise\u003cnever\u003e((_, reject) =\u003e {\n+      timeoutId = setTimeout(\n         () =\u003e reject(new Error(\"Timed out waiting for subscribed messages\")),\n         15_000,\n       );\n     });\n-    await Promise.race([subscriptionFinished, timeout]);\n+    try {\n+      await Promise.race([subscriptionFinished, timeout]);\n+    } finally {\n+      clearTimeout(timeoutId);\n+    }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3376418739/src/index.ts
@@ -3,9 +3,10 @@
   ServiceBusClient,
   type ServiceBusReceiver,
   type ServiceBusSender,
-  type ServiceBusSubscriber,
 } from "@azure/service-bus";
 
+type ServiceBusSubscription = ReturnType<ServiceBusReceiver["subscribe"]>;
+
 function requiredEnvironmentVariable(name: string): string {
   const value = process.env[name];
   if (!value) {
@@ -37,7 +38,7 @@
   let queueSender: ServiceBusSender | undefined;
   let pullReceiver: ServiceBusReceiver | undefined;
   let subscriptionReceiver: ServiceBusReceiver | undefined;
-  let subscription: ServiceBusSubscriber | undefined;
+  let subscription: ServiceBusSubscription | undefined;
   let topicSender: ServiceBusSender | undefined;
   let topicReceiver: ServiceBusReceiver | undefined;
 
@@ -112,13 +113,18 @@
       },
     );
 
+    let timeoutId: NodeJS.Timeout | undefined;
     const timeout = new Promise<never>((_, reject) => {
-      setTimeout(
+      timeoutId = setTimeout(
         () => reject(new Error("Timed out waiting for subscribed messages")),
         15_000,
       );
     });
-    await Promise.race([subscriptionFinished, timeout]);
+    try {
+      await Promise.race([subscriptionFinished, timeout]);
+    } finally {
+      clearTimeout(timeoutId);
+    }
     await subscription.close();
     subscription = undefined;
 

```

### ✅ 11. powershell — 3141ms

**Input:**

```json
{"command":"npm run build","description":"Type-check the corrected demo","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-service-bus-typescript-demo@1.0.0 build
> tsc --noEmit
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace and the Service Bus reference patterns, then add a locally runnable TypeScript example without contacting Azure.

The example will use one reusable client, manually settle pulled and subscribed messages, and close the subscription before receivers, senders, and client.

The SDK’s subscription handle is not exported as `ServiceBusSubscriber`; I’m switching to the receiver method’s inferred return type and clearing the wait timer deterministically.

Implemented the Azure Service Bus TypeScript demo in `src/index.ts`.

It includes:

- `ServiceBusClient` connection-string authentication
- Single queue message and a five-message batch
- `receiveMessages()` with `completeMessage()`
- `subscribe()` with `processMessage` and `processError`
- Topic publishing and subscription receiving
- Ordered `close()` cleanup for subscriptions, receivers, senders, and client
- Required `@azure/service-bus` dependency in `package.json`
- Environment placeholders in `.env.example`
- Setup and execution instructions in `README.md`

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Includes the required Azure Service Bus dependency and complete TypeScript project configuration.

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
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 12.7% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.5% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 14.5% | ✅ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 14.5% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 14.5% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.5% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 14.5% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 6.8750** | **62.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
