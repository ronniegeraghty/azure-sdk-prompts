# Evaluation Report: event-hubs-dp-js-ts-streaming

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 396.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-js-ts-streaming` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 13/18 |
| Duration | 396.1s |
| Timestamp | 2026-08-29T11:13:13Z |
| Files Generated | 4 |
| Event Count | 4533 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 175.3s |
| Review | 219.5s |
| **Total** | **396.1s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-eventhub-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=356580 out=3966 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 2732ms

**Input:**

```json
{"skill":"azure-eventhub-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Event Hubs SDK for TypeScript

High-throughput event streaming and real-time data ingestion.

## Installation

```bash
npm install @azure/event-hubs @azure/identity
```

For checkpointing with consumer groups:
```bash
npm install @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

## Environment Variables

```bash
EVENTHUB_NAMESPACE=<namespace>.servicebus.windows.net
EVENTHUB_NAME=my-eventhub
STORAGE_ACCOUNT_NAME=<storage-account>
STORAGE_CONTAINER_NAME=checkpoints
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

```typescript
import { EventHubProducerClient, EventHubConsumerClient } from "@azure/event-hubs";
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

const fullyQualifiedNamespace = process.env.EVENTHUB_NAMESPACE!;
const eventHubName = process.env.EVENTHUB_NAME!;
// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

// Producer
const producer = new EventHubProducerClient(fullyQualifiedNamespace, eventHubName, credential);

// Consumer
const consumer = new EventHubConsumerClient(
  "$Default", // Consumer group
  fullyQualifiedNamespace,
  eventHubName,
  credential
);
```

## Core Workflow

### Send Events

```typescript
const producer = new EventHubProducerClient(namespace, eventHubName, credential);

// Create batch and add events
const batch = await producer.createBatch();
batch.tryAdd({ body: { temperature: 72.5, deviceId: "sensor-1" } });
batch.tryAdd({ body: { temperature: 68.2, deviceId: "sensor-2" } 
... (truncated)
```

### ✅ 2. view — 12430ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-eventhub-ts\\references\\checkpointing.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/event-processing.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/event-processing.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/event-processing.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/event-processing.md
@@ -1,434 +1,434 @@
 # Event Processing Reference
 
 Receiving and processing events with Azure Event Hubs using the @azure/event-hubs TypeScript SDK.
 
 ## Overview
 
 Event Hubs provides high-throughput event streaming. This reference covers the EventHubConsumerClient, subscription patterns, event handlers, and processing strategies.
 
 ## Key Interfaces
 
 ```typescript
 import {
   EventHubConsumerClient,
   ReceivedEventData,
   PartitionContext,
   SubscriptionEventHandlers,
   SubscribeOptions,
   Subscription
 } from "@azure/event-hubs";
 
 // ReceivedEventData - structure of received events
 interface ReceivedEventData {
   body: any;
   contentType?: string;
   correlationId?: string | number | Buffer;
   enqueuedTimeUtc: Date;
   messageId?: string | number | Buffer;
   offset: string;
   partitionKey: string | null;
   properties?: Record<string, any>;
   sequenceNumber: number;
   systemProperties?: Record<string, any>;
 }
 
 // PartitionContext - context for event handlers
 interface PartitionContext {
   readonly consumerGroup: string;
   readonly eventHubName: string;
   readonly fullyQualifiedNamespace: string;
   readonly partitionId: string;
   readonly lastEnqueuedEventProperties?: LastEnqueuedEventProperties;
   updateCheckpoint(eventData: ReceivedEventData): Promise<void>;
 }
 
 // Handl
... (truncated)
```

### ✅ 3. view — 11003ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-eventhub-ts\\references\\event-processing.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/checkpointing.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/checkpointing.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/checkpointing.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-eventhub-ts/references/checkpointing.md
@@ -1,426 +1,426 @@
 # Checkpointing Reference
 
 Persistent checkpointing with BlobCheckpointStore for Azure Event Hubs consumer applications.
 
 ## Overview
 
 Checkpointing tracks the last successfully processed event position per partition. This enables:
 - **Resumption** — Continue from where you left off after restart
 - **Load balancing** — Distribute partitions across multiple consumers
 - **Exactly-once processing** — When combined with idempotent downstream operations
 
 ## Installation
 
 ```bash
 npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob @azure/identity
 ```
 
 ## Key Interfaces
 
 ```typescript
 import { CheckpointStore, Checkpoint, PartitionOwnership } from "@azure/event-hubs";
 
 // CheckpointStore interface (implemented by BlobCheckpointStore)
 interface CheckpointStore {
   listCheckpoints(
     fullyQualifiedNamespace: string,
     eventHubName: string,
     consumerGroup: string
   ): Promise<Checkpoint[]>;
 
   updateCheckpoint(checkpoint: Checkpoint): Promise<void>;
 
   listOwnership(
     fullyQualifiedNamespace: string,
     eventHubName: string,
     consumerGroup: string
   ): Promise<PartitionOwnership[]>;
 
   claimOwnership(
     partitionOwnership: PartitionOwnership[]
   ): Promise<PartitionOwnership[]>;
 }
 
 // Checkpoint structure
 inte
... (truncated)
```

### ✅ 4. glob — 8814ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-eventhub-ts","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 7773ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 9926ms

**Input:**

```json
{"intent":"Get Azure SDK best practices for a local-only TypeScript Event Hubs producer/consumer example using BlobCheckpointStore and connection strings.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2411ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only TypeScript sample for Azure Event Hubs producer and consumer clients using connection strings, batched events, BlobCheckpointStore, async handlers, checkpointing, and graceful shutdown.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. powershell — 20674ms

**Input:**

```json
{"command":"$packages = '@azure/event-hubs','@azure/eventhubs-checkpointstore-blob','@azure/storage-blob','typescript','@types/node','tsx'; foreach ($package in $packages) { $version = npm view $package version --silent; Write-Output \"$package=$version\" }","description":"Get latest npm package versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
@azure/event-hubs=6.0.4
@azure/eventhubs-checkpointstore-blob=2.0.2-alpha.20260312.1
@azure/storage-blob=12.33.0
typescript=7.0.2
@types/node=26.2.0
tsx=4.23.12
<shellId: 0 completed with exit code 0>
```

### ✅ 9. powershell — 10354ms

**Input:**

```json
{"command":"npm view @azure/event-hubs dist-tags --json; npm view @azure/eventhubs-checkpointstore-blob dist-tags --json; npm view @azure/storage-blob dist-tags --json","description":"Check stable Azure SDK package tags","initial_wait":120,"mode":"sync"}
```

**Output:**

```
{
  "dev": "6.0.5-alpha.20260828.2",
  "latest": "6.0.4"
}
{
  "dev": "2.0.2-alpha.20260828.2",
  "latest": "2.0.2-alpha.20260312.1"
}
{
  "dev": "12.34.0-alpha.20260828.2",
  "latest": "12.33.0",
  "beta": "12.34.0-beta.1",
  "next": "12.29.0-beta.1"
}
<shellId: 1 completed with exit code 0>
```

### ✅ 10. apply_patch — 683ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-event-hubs-typescript-demo\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc --project tsconfig.json\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/event-hubs\": \"6.0.4\",\n+    \"@azure/eventhubs-checkpointstore-blob\": \"2.0.2-alpha.20260312.1\",\n+    \"@azure/storage-blob\": \"12.33.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"26.2.0\",\n+    \"tsx\": \"4.23.12\",\n+    \"typescript\": \"7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+EVENTHUB_CONNECTION_STRING=Endpoint=sb://\u003cnamespace\u003e.servicebus.windows.net/;SharedAccessKeyName=\u003ckey-name\u003e;SharedAccessKey=\u003ckey\u003e\n+EVENTHUB_NAME=\u003cevent-hub-name\u003e\n+EVENTHUB_CONSUMER_GROUP=$Default\n+AZURE_STORAGE_CONNECTION_STRING=DefaultEndpointsProtocol=https;AccountName=\u003caccount-name\u003e;AccountKey=\u003caccount-key\u003e;EndpointSuffix=core.windows.net\n+BLOB_CONTAINER_NAME=event-hub-checkpoints\n*** Add File: src/index.ts\n+import {\n+  EventHubConsumerClient,\n+  EventHubProducerClient,\n+  earliestEventPosition,\n+  type Subscription,\n+} from \"@azure/event-hubs\";\n+import { BlobCheckpointStore } from \"@azure/eventhubs-checkpointstore-blob\";\n+import { ContainerClient } from \"@azure/storage-blob\";\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+\n+  return value;\n+}\n+\n+function waitForShutdownSignal(): {\n+  promise: Promise\u003cvoid\u003e;\n+  dispose: () =\u003e void;\n+} {\n+  let resolveShutdown: () =\u003e void;\n+  const promise = new Promise\u003cvoid\u003e((resolve) =\u003e {\n+    resolveShutdown = resolve;\n+  });\n+\n+  const handleSignal = (signal: NodeJS.Signals): void =\u003e {\n+    console.log(`\\nReceived ${signal}; shutting down...`);\n+    resolveShutdown();\n+  };\n+\n+  process.once(\"SIGINT\", handleSignal);\n+  process.once(\"SIGTERM\", handleSignal);\n+\n+  return {\n+    promise,\n+    dispose: () =\u003e {\n+      process.off(\"SIGINT\", handleSignal);\n+      process.off(\"SIGTERM\", handleSignal);\n+    },\n+  };\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const eventHubConnectionString = requireEnvironmentVariable(\n+    \"EVENTHUB_CONNECTION_STRING\",\n+  );\n+  const eventHubName = requireEnvironmentVariable(\"EVENTHUB_NAME\");\n+  const storageConnectionString = requireEnvironmentVariable(\n+    \"AZURE_STORAGE_CONNECTION_STRING\",\n+  );\n+  const consumerGroup = process.env.EVENTHUB_CONSUMER_GROUP ?? \"$Default\";\n+  const containerName =\n+    process.env.BLOB_CONTAINER_NAME ?? \"event-hub-checkpoints\";\n+\n+  const producer = new EventHubProducerClient(\n+    eventHubConnectionString,\n+    eventHubName,\n+  );\n+  let consumer: EventHubConsumerClient | undefined;\n+  let subscription: Subscription | undefined;\n+  const shutdown = waitForShutdownSignal();\n+\n+  try {\n+    const batch = await producer.createBatch();\n+\n+    for (let index = 1; index \u003c= 10; index += 1) {\n+      const wasAdded = batch.tryAdd({\n+        body: {\n+          message: `Event ${index}`,\n+          createdAt: new Date().toISOString(),\n+        },\n+        contentType: \"application/json\",\n+        properties: {\n+          eventType: \"demo\",\n+          eventNumber: index,\n+          source: \"typescript-sample\",\n+        },\n+      });\n+\n+      if (!wasAdded) {\n+        throw new Error(`Event ${index} did not fit in the Event Hubs batch.`);\n+      }\n+    }\n+\n+    await producer.sendBatch(batch);\n+    console.log(`Sent ${batch.count} events.`);\n+\n+    const containerClient = new ContainerClient(\n+      storageConnectionString,\n+      containerName,\n+    );\n+    await containerClient.createIfNotExists();\n+\n+    const checkpointStore = new BlobCheckpointStore(containerClient);\n+    consumer = new EventHubConsumerClient(\n+      consumerGroup,\n+      eventHubConnectionString,\n+      eventHubName,\n+      checkpointStore,\n+    );\n+\n+    subscription = consumer.subscribe(\n+      {\n+        processEvents: async (events, context) =\u003e {\n+          if (events.length === 0) {\n+            return;\n+          }\n+\n+          for (const event of events) {\n+            console.log(\n+              `Partition ${context.partitionId}:`,\n+              JSON.stringify(event.body),\n+              \"properties:\",\n+              event.properties,\n+            );\n+          }\n+\n+          const lastEvent = events.at(-1);\n+          if (lastEvent) {\n+            await context.updateCheckpoint(lastEvent);\n+            console.log(\n+              `Checkpoint updated for partition ${context.partitionId} at sequence ${lastEvent.sequenceNumber}.`,\n+            );\n+          }\n+        },\n+        processError: async (error, context) =\u003e {\n+          console.error(\n+            `Error while processing partition ${context.partitionId}:`,\n+            error,\n+          );\n+        },\n+      },\n+      {\n+        startPosition: earliestEventPosition,\n+        maxBatchSize: 100,\n+        maxWaitTimeInSeconds: 10,\n+      },\n+    );\n+\n+    console.log(\"Receiving events. Press Ctrl+C to stop.\");\n+    await shutdown.promise;\n+  } finally {\n+    shutdown.dispose();\n+    await subscription?.close();\n+    await consumer?.close();\n+    await producer.close();\n+    console.log(\"Event Hubs clients closed.\");\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Application failed:\", error);\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869/package.json
@@ -1,0 +1,21 @@
+{
+  "name": "azure-event-hubs-typescript-demo",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc --project tsconfig.json",
+    "start": "tsx src/index.ts"
+  },
+  "dependencies": {
+    "@azure/event-hubs": "6.0.4",
+    "@azure/eventhubs-checkpointstore-blob": "2.0.2-alpha.20260312.1",
+    "@azure/storage-blob": "12.33.0"
+  },
+  "devDependencies": {
+    "@types/node": "26.2.0",
+    "tsx": "4.23.12",
+    "typescript": "7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1718118869/tsconfig.json
@@ -1,0 +1,15 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "outDir": "dist",
+    "rootDir": "src",
+    "strict": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-js-ts-streaming-js-
... (truncated)
```

### ✅ 11. powershell — 28606ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 86 packages in 26s

> azure-event-hubs-typescript-demo@1.0.0 build
> tsc --project tsconfig.json
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’m grounding the sample in the SDK’s checkpointing and event-processing guidance, then I’ll create a locally buildable TypeScript project.

The sample will use environment-provided connection strings, fail fast on missing configuration or an oversized batch, checkpoint only after successful processing, and close the subscription and both clients on shutdown.

Implemented a complete TypeScript sample in `src/index.ts` with:

- `EventHubProducerClient`, `createBatch()`, 10 events with custom properties, and `sendBatch()`
- `EventHubConsumerClient` with `BlobCheckpointStore`
- `processEvents` and `processError` handlers
- Event-body logging and per-partition checkpoint updates
- Graceful `SIGINT`/`SIGTERM` shutdown using `close()`
- Environment-variable validation and batch-size error handling

Install and run:

```powershell
npm install
Copy-Item .env.example .env
# Set the environment variables from .env in your shell
npm run start
```

Required runtime packages are declared in `package.json`:

```powershell
npm install @azure/event-hubs @azure/eventhubs-checkpointstore-blob @azure/storage-blob
```

`@azure/storage-blob` is required to construct the `ContainerClient` used by `BlobCheckpointStore`.

References: [Event Hubs JavaScript quickstart](https://learn.microsoft.com/azure/event-hubs/event-hubs-node-get-started-send), [Azure Event Hubs JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/event-hubs-readme), and [Blob checkpoint store package](https://www.npmjs.com/package/@azure/eventhubs-checkpointstore-blob).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete dependency manifest with matching lockfile
- Includes tsconfig.json and appropriate Node.js type definitions

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
hyoka run --prompt-id event-hubs-dp-js-ts-streaming --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
