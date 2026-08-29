# Evaluation Report: service-bus-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 351.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 12/18 |
| Duration | 351.3s |
| Timestamp | 2026-08-29T09:16:22Z |
| Files Generated | 5 |
| Event Count | 5760 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 197.2s |
| Review | 152.8s |
| **Total** | **351.3s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-servicebus-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=398629 out=4495 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 2635ms

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

### ✅ 2. view — 11134ms

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

### ✅ 3. view — 9999ms

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

### ✅ 4. glob — 7650ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 8915ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate a local-only TypeScript Azure Service Bus SDK sample using a connection string, queue single and batch sends, receiveMessages with explicit completion, subscribe handlers, topic/subscription messaging, and proper cleanup."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 2859ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only TypeScript Azure Service Bus SDK sample using a connection string, queue single and batch sends, receiveMessages with explicit completion, subscribe handlers, topic/subscription messaging, and proper cleanup.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. powershell — 6827ms

**Input:**

```json
{"command":"npm view @azure/service-bus version --silent","description":"Check latest Service Bus package version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
7.10.0-alpha.20260312.1
<shellId: 0 completed with exit code 0>
```

### ✅ 8. powershell — 7138ms

**Input:**

```json
{"command":"npm view @azure/service-bus dist-tags --json --silent","description":"Check stable Service Bus package tag","initial_wait":30,"mode":"sync"}
```

**Output:**

```
{
  "dev": "7.10.0-alpha.20260828.2",
  "latest": "7.10.0-alpha.20260312.1",
  "beta": "7.10.0-beta.5",
  "next": "7.10.0-beta.1"
}
<shellId: 1 completed with exit code 0>
```

### ✅ 9. powershell — 5215ms

**Input:**

```json
{"command":"$versions = npm view @azure/service-bus versions --json --silent | ConvertFrom-Json; $versions | Where-Object { $_ -notmatch '-' } | Select-Object -Last 1","description":"Find latest stable Service Bus version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
7.9.5
<shellId: 2 completed with exit code 0>
```

### ✅ 10. apply_patch — 675ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-service-bus-typescript-demo\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/service-bus\": \"^7.9.5\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+SERVICEBUS_CONNECTION_STRING=Endpoint=sb://your-namespace.servicebus.windows.net/;SharedAccessKeyName=your-key-name;SharedAccessKey=your-key\n+SERVICEBUS_QUEUE_NAME=my-queue\n+SERVICEBUS_TOPIC_NAME=my-topic\n+SERVICEBUS_SUBSCRIPTION_NAME=my-subscription\n*** Add File: src/index.ts\n+import {\n+  ServiceBusClient,\n+  ServiceBusMessage,\n+  ServiceBusReceiver,\n+  ServiceBusSender,\n+  ServiceBusError,\n+  type ProcessErrorArgs,\n+  type ServiceBusReceivedMessage,\n+} from \"@azure/service-bus\";\n+\n+interface Closeable {\n+  close(): Promise\u003cvoid\u003e;\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+\n+  return value;\n+}\n+\n+function logMessage(prefix: string, message: ServiceBusReceivedMessage): void {\n+  console.log(`${prefix} ${String(message.messageId)}:`, message.body);\n+}\n+\n+function logSubscriptionError(args: ProcessErrorArgs): void {\n+  console.error(\n+    `Subscription error from ${args.errorSource} on ${args.entityPath}:`,\n+    args.error,\n+  );\n+\n+  if (args.error instanceof ServiceBusError) {\n+    console.error(\n+      `Service Bus error code=${args.error.code}, retryable=${args.error.retryable}`,\n+    );\n+  }\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const connectionString = requireEnvironmentVariable(\n+    \"SERVICEBUS_CONNECTION_STRING\",\n+  );\n+  const queueName = requireEnvironmentVariable(\"SERVICEBUS_QUEUE_NAME\");\n+  const topicName = requireEnvironmentVariable(\"SERVICEBUS_TOPIC_NAME\");\n+  const subscriptionName = requireEnvironmentVariable(\n+    \"SERVICEBUS_SUBSCRIPTION_NAME\",\n+  );\n+\n+  // The connection string is read from the environment; never embed credentials in code.\n+  const client = new ServiceBusClient(connectionString);\n+  const closeables: Closeable[] = [];\n+\n+  const track = \u003cT extends Closeable\u003e(resource: T): T =\u003e {\n+    closeables.push(resource);\n+    return resource;\n+  };\n+\n+  const closeNow = async (resource: Closeable): Promise\u003cvoid\u003e =\u003e {\n+    await resource.close();\n+    const index = closeables.lastIndexOf(resource);\n+    if (index !== -1) {\n+      closeables.splice(index, 1);\n+    }\n+  };\n+\n+  try {\n+    const queueSender: ServiceBusSender = track(\n+      client.createSender(queueName),\n+    );\n+\n+    const singleMessage: ServiceBusMessage = {\n+      body: { kind: \"single\", text: \"Hello from Azure Service Bus\" },\n+      contentType: \"application/json\",\n+      messageId: `single-${crypto.randomUUID()}`,\n+    };\n+    await queueSender.sendMessages(singleMessage);\n+    console.log(\"Sent one queue message\");\n+\n+    const batch = await queueSender.createMessageBatch();\n+    for (let index = 1; index \u003c= 5; index += 1) {\n+      const wasAdded = batch.tryAddMessage({\n+        body: { kind: \"batch\", sequence: index },\n+        contentType: \"application/json\",\n+        messageId: `batch-${index}-${crypto.randomUUID()}`,\n+      });\n+\n+      if (!wasAdded) {\n+        throw new Error(`Message ${index} did not fit in the Service Bus batch`);\n+      }\n+    }\n+\n+    await queueSender.sendMessages(batch);\n+    console.log(`Sent a batch of ${batch.count} queue messages`);\n+\n+    const queueReceiver: ServiceBusReceiver = track(\n+      client.createReceiver(queueName, { receiveMode: \"peekLock\" }),\n+    );\n+    const receivedMessages = await queueReceiver.receiveMessages(6, {\n+      maxWaitTimeInMs: 10_000,\n+    });\n+\n+    for (const message of receivedMessages) {\n+      logMessage(\"Pulled queue message\", message);\n+      await queueReceiver.completeMessage(message);\n+    }\n+    console.log(`Completed ${receivedMessages.length} pulled messages`);\n+    await closeNow(queueReceiver);\n+\n+    const subscribedMessageId = `subscribed-${crypto.randomUUID()}`;\n+    const subscribedReceiver: ServiceBusReceiver = track(\n+      client.createReceiver(queueName, { receiveMode: \"peekLock\" }),\n+    );\n+\n+    let notifySubscribedMessage: (() =\u003e void) | undefined;\n+    const subscribedMessageProcessed = new Promise\u003cvoid\u003e((resolve) =\u003e {\n+      notifySubscribedMessage = resolve;\n+    });\n+\n+    const streamingSubscription = track(\n+      subscribedReceiver.subscribe(\n+        {\n+          processMessage: async (message) =\u003e {\n+            logMessage(\"Subscribed queue message\", message);\n+            if (message.messageId === subscribedMessageId) {\n+              notifySubscribedMessage?.();\n+            }\n+            // subscribe() auto-completes when this handler succeeds.\n+          },\n+          processError: async (args) =\u003e {\n+            logSubscriptionError(args);\n+          },\n+        },\n+        {\n+          autoCompleteMessages: true,\n+          maxConcurrentCalls: 1,\n+        },\n+      ),\n+    );\n+\n+    await queueSender.sendMessages({\n+      body: { kind: \"subscription\", text: \"Process me with subscribe()\" },\n+      contentType: \"application/json\",\n+      messageId: subscribedMessageId,\n+    });\n+\n+    let subscriptionTimeout: NodeJS.Timeout | undefined;\n+    try {\n+      await Promise.race([\n+        subscribedMessageProcessed,\n+        new Promise\u003cnever\u003e((_, reject) =\u003e {\n+          subscriptionTimeout = setTimeout(\n+            () =\u003e reject(new Error(\"Timed out waiting for subscribed message\")),\n+            15_000,\n+          );\n+        }),\n+      ]);\n+    } finally {\n+      if (subscriptionTimeout) {\n+        clearTimeout(subscriptionTimeout);\n+      }\n+    }\n+\n+    await closeNow(streamingSubscription);\n+    await closeNow(subscribedReceiver);\n+    await closeNow(queueSender);\n+\n+    const topicSender: ServiceBusSender = track(\n+      client.createSender(topicName),\n+    );\n+    await topicSender.sendMessages({\n+      body: { event: \"demo.created\", createdAt: new Date().toISOString() },\n+      contentType: \"application/json\",\n+      subject: \"demo/created\",\n+      messageId: `topic-${crypto.randomUUID()}`,\n+    });\n+    console.log(\"Sent one topic message\");\n+\n+    const topicReceiver: ServiceBusReceiver = track(\n+      client.createReceiver(topicName, subscriptionName, {\n+        receiveMode: \"peekLock\",\n+      }),\n+    );\n+    const topicMessages = await topicReceiver.receiveMessages(1, {\n+      maxWaitTimeInMs: 10_000,\n+    });\n+\n+    if (topicMessages.length === 0) {\n+      throw new Error(\n+        `No message arrived on topic subscription \"${subscriptionName}\"`,\n+      );\n+    }\n+\n+    for (const message of topicMessages) {\n+      logMessage(\"Topic subscription message\", message);\n+      await topicReceiver.completeMessage(message);\n+    }\n+\n+    await closeNow(topicReceiver);\n+    await closeNow(topicSender);\n+  } finally {\n+    for (const resource of closeables.reverse()) {\n+      await resource.close();\n+    }\n+    await client.close();\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Service Bus demo failed:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Service Bus TypeScript demo\n+\n+This sample uses `@azure/service-bus` to send and receive queue and topic\n+messages. The queue, topic, and subscription must already exist.\n+\n+## Install\n+\n+```powershell\n+npm install\n+```\n+\n+The required runtime package is:\n+\n+```powershell\n+npm install @azure/service-bus\n+```\n+\n+## Configure and run\n+\n+Set the values from `.env.example` in your shell. Do not commit a real Service\n+Bus connection string.\n+\n+```powershell\n+$env:SERVICEBUS_CONNECTION_STRING = \"Endpoint=sb://...\"\n+$env:SERVICEBUS_QUEUE_NAME = \"my-queue\"\n+$env:SERVICEBUS_TOPIC_NAME = \"my-topic\"\n+$env:SERVICEBUS_SUBSCRIPTION_NAME = \"my-subscription\"\n+\n+npm run build\n+npm start\n+```\n+\n+The connection string needs send and receive permissions for all configured\n+entities. For production applications, prefer Microsoft Entra authentication\n+with a managed identity instead of a connection string.\n+\n+SDK references:\n+\n+- https://learn.microsoft.com/javascript/api/overview/azure/service-bus-readme\n+- https://learn.microsoft.com/azure/service-bus-messaging/service-bus-nodejs-how-to-use-queues\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/package.json
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
+    "@types/node": "^24.3.0",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/tsconfig.json
@@ -1,0 +1,15 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "rootDir": "src",
+    "outDir": "dist",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "esModuleInterop": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3658037132/.env.example b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-s
... (truncated)
```

### ✅ 11. powershell — 20662ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 76 packages in 18s

> azure-service-bus-typescript-demo@1.0.0 build
> tsc
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll scaffold a local TypeScript example with environment-based configuration, explicit message settlement, subscription handlers, and deterministic cleanup.

The workspace is empty, so I’m creating a minimal compile-and-run sample. Queue/topic entities must already exist; the code will only perform data-plane messaging and will never provision Azure resources.

Created a complete TypeScript Azure Service Bus example in `src\index.ts`.

It demonstrates:

- Connection-string `ServiceBusClient`
- Single and five-message batch queue sends
- `receiveMessages()` with `completeMessage()`
- `subscribe()` with `processMessage` and `processError`
- Topic publishing and subscription receiving
- Ordered `close()` cleanup for subscriptions, receivers, senders, and client
- Environment-based configuration with `.env.example`

Run with:

```powershell
npm install
npm run build
npm start
```

Required package:

```powershell
npm install @azure/service-bus
```

References: [Azure Service Bus SDK](https://learn.microsoft.com/javascript/api/overview/azure/service-bus-readme) and [Service Bus queue quickstart](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-nodejs-how-to-use-queues).

## Reviewer Notes (LLM-as-Judge)

### Summary

The sole dependency and TypeScript project-configuration criterion is fully met.

### Strengths

- Correctly declares @azure/service-bus, TypeScript, and Node.js type definitions.
- Includes a valid tsconfig.json and appropriate build/start scripts.

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
hyoka run --prompt-id service-bus-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
