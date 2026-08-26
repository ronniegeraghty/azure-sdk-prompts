# Evaluation Report: ai-agents-dp-js-ts-file-search

**Config:** js-ts-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 431.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-agents-dp-js-ts-file-search` |
| Config | js-ts-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 15/19 |
| Duration | 431.3s |
| Timestamp | 2026-08-26T12:45:26Z |
| Files Generated | 5 |
| Event Count | 4536 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 187.0s |
| Review | 243.0s |
| **Total** | **431.3s** |

## Configuration

- **name:** js-ts-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=255403 out=3401 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"@azure/ai-agents JavaScript TypeScript file search vector store upload createAndPoll current API 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's a summary and code sample for using the current (2026) `@azure/ai-agents` API in JavaScript/TypeScript to create a vector store, upload files, enable file search, and use create-and-poll patterns. This workflow supports document search and retrieval for generative AI or RAG use cases:\n\n---\n\n### Prerequisites\n- Node.js LTS version\n- Azure account and an AI Foundry project\n- `@azure/ai-agents` and `@azure/identity` NPM packages\n\n```bash\nnpm install @azure/ai-agents @azure/identity\n```\n\n---\n\n### Imports and Authentication\n\n```typescript\nimport { AgentsClient } from \"@azure/ai-agents\";\nimport { DefaultAzureCredential } from \"@azure/identity\";\n\nconst endpoint = process.env[\"PROJECT_ENDPOINT\"] || \"<your-ai-project-endpoint>\";\nconst client = new AgentsClient(endpoint, new DefaultAzureCredential());\n```\n\n---\n\n### Step 1: Upload File(s)\n\n```typescript\nconst uploadedFile = await client.uploadFile({\n  filePath: \"./path/to/myDocument.pdf\",\n  fileName: \"myDocument.pdf\"\n});\n// uploadedFile.id will be used in next operations\n```\n\n---\n\n### Step 2: Create Vector Store and Upload Documents\n\n```typescript\nconst vectorStore = await client.createVectorStoreWithFiles({  // check method name in actual SDK docs\n  files: [uploadedFile.id],\n  vectorStoreName: \"Example Knowledge Base\"\n});\n\n// Wait for completion or poll, depending on SDK version:\nconst operation = await client.createAndPollVectorStore({\n  files: [uploadedFile.id],\n  vectorStoreName: \"Example Vector Store\"\n});\n// operation.result gives you the completed vector store\n```\n\n---\n\n### Step 3: Enable File Search Tool for the Agent\n\n```typescript\nconst fileSearchTool = {\n  type: \"fileSearch\",\n  vectorStoreIds: [vectorStore.id]\n};\n\nconst agent = await client.createAgent({\n  name: \"Smart File Search Agent\",\n  tools: [fileSearchTool]\n});\n```\n\n---\n\n### Step 4: Attach Resources/Use File Search in Threa
... (truncated)
```

### ✅ 3. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-agents fileSearchTool vectorStoreIds language:TypeScript"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"index.ts","path":"sdk/ai/ai-agents/src/index.ts","sha":"98b84cf1b12542bc9838713991ada16d54ef92b0","text_matches":[{"fragment":"export { OpenApiTool } from \"./utils/OpenApiTool.js\";\nexport { AgentsClient } from \"./agentsClient.js\";\nexport {","matches":[{"indices":[63,75],"text":"AgentsClient"},{"indices":[86,98],"text":"agentsClient"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"  FilesDeleteFileOptionalParams,\n  FilesUploadFileOptionalParams,\n  FilesListFilesOptionalParams,","matches":[{"indices":[40,50],"text":"UploadFile"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]}],"total_count":1}
```

### ✅ 4. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-agents createVectorStoreAndPoll language:TypeScript"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"operations.ts","path":"sdk/ai/ai-agents/src/api/vectorStores/operations.ts","sha":"388a7b40c0d2d69e6fa4599c540c6ea469527765","text_matches":[{"fragment":"/**\n * Creates a vector store and poll.\n */\nexport function createVectorStoreAndPoll(\n  context: Client,\n  options: VectorStoresCreateVectorStoreOptionalParams = { requestOptions: {} },\n): PollerLike<OperationState<VectorStore>, VectorStore> {","matches":[{"indices":[60,84],"text":"createVectorStoreAndPoll"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/api/vectorStores/operations.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"index.ts","path":"sdk/ai/ai-agents/src/classic/vectorStores/index.ts","sha":"ddca1d4e03209c6dc4add7f138c1fbab45ad69f2","text_matches":[{"fragment":"  listVectorStores,\n  createVectorStoreAndPoll,\n} from \"../../api/vectorStores/operations.js\";","matches":[{"indices":[22,46],"text":"createVectorStoreAndPoll"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/classic/vectorStores/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"    createAndPoll: (options?: VectorStoresCreateVectorStoreOptionalParams) =>\n      createVectorStoreAndPoll(context, options),\n    list: (options?: VectorStoresListVectorStoresOptionalParams) =>","matches":[{"indices":[84,108],"text":"createVectorStoreAndPoll"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/classic/vectorStores/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]}],"total_count":2}
```

### ✅ 5. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-agents uploadFile language:TypeScript agentsClient"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"utils.ts","path":"sdk/ai/ai-agents/src/utils/utils.ts","sha":"3c214847c3d92ec94fd0cba21bb4129d3951ba9f","text_matches":[{"fragment":"    vectorStores?: Array<VectorStoreConfigurations>,\n    definitionDetails?: FileSearchToolDefinitionDetails,\n  ): { definition: FileSearchToolDefinition; resources: ToolResources } {\n    const tool = ToolUtility.createFileSearchTool(vectorStoreIds, vectorStores, definitionDetails);\n    this.toolDefinitions.push(tool.definition);\n    this.toolResources = { ...this.toolResources, ...tool.resources };\n    return tool;","matches":[{"indices":[77,91],"text":"FileSearchTool"},{"indices":[129,143],"text":"FileSearchTool"},{"indices":[219,233],"text":"FileSearchTool"},{"indices":[234,248],"text":"vectorStoreIds"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/utils/utils.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"models.ts","path":"sdk/ai/ai-agents/src/models/models.ts","sha":"b725ec78ad212ee75a1fc2a6db48abf983f16052","text_matches":[{"fragment":"\nexport function fileSearchToolResourceDeserializer(item: any): FileSearchToolResource {\n  return {\n    vectorStoreIds: !item[\"vector_store_ids\"]\n      ? item[\"vector_store_ids\"]\n      : item[\"vector_store_ids\"].map((p: any) => {\n          return p;","matches":[{"indices":[17,31],"text":"fileSearchTool"},{"indices":[64,78],"text":"FileSearchTool"},{"indices":[104,118],"text":"vectorStoreIds"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/src/models/models.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]}],"total_count":2}
```

### ✅ 6. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://azuresdkdocs.z19.web.core.windows.net/javascript/azure-ai-agents/1.1.0/index.html"}
```

**Output:**

```
Contents of https://azuresdkdocs.z19.web.core.windows.net/javascript/azure-ai-agents/1.1.0/index.html:
# Azure AI Agents client library for JavaScript[](about:blank#md:azure-ai-agents-client-library-for-javascript)

Use the AI Agents client library to:

*   **Develop Agents using the Azure AI Agent Service**, leveraging an extensive ecosystem of models, tools, and capabilities from OpenAI, Microsoft, and other LLM providers. The Azure AI Agent Service enables the building of Agents for a wide range of generative AI use cases.

*   **Note:** While this package can be used independently, we recommend using the [Azure AI Projects client library](https://www.npmjs.com/package/@azure/ai-projects) for an enhanced experience. The Projects library provides simplified access to advanced functionality, such as creating and managing agents, enumerating AI models, working with datasets and managing search indexes, evaluating generative AI performance, and enabling OpenTelemetry tracing.

[Product documentation](https://aka.ms/azsdk/azure-ai-projects/product-doc) | [Samples](https://github.com/Azure/azure-sdk-for-js/tree/@azure/ai-agents_1.1.0/sdk/ai/ai-agents/samples/) | [Package (npm)](https://www.npmjs.com/package/@azure/ai-agents) | [API reference documentation](https://learn.microsoft.com/javascript/api/overview/azure/ai-agents-readme?view=azure-node-latest)

## Table of contents[](about:blank#md:table-of-contents)

## Getting started[](about:blank#md:getting-started)

### Prerequisite[](about:blank#md:prerequisite)

### Authorization[](about:blank#md:authorization)

*   [Entra ID](https://learn.microsoft.com/azure/ai-services/authentication?tabs=powershell#authenticate-with-microsoft-entra-id) is needed to authenticate the client. Your application needs an object that implements the [TokenCredential](https://learn.microsoft.com/javascript/api/@azure/core-auth/tokencredential) interface. Code samples here use [DefaultAzureCredential](https://learn.microsoft.com/javascript/a
... (truncated)
```

### ✅ 7. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","type"],"owner":"Azure","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src","ref":"refs/tags/@azure/ai-agents_1.1.0","repo":"azure-sdk-for-js"}
```

**Output:**

```
[{"name":"agentTokenUsage.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentTokenUsage.ts","type":"file"},{"name":"agentsAzureAiSearch.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsAzureAiSearch.ts","type":"file"},{"name":"agentsBasics.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsBasics.ts","type":"file"},{"name":"agentsBingGrounding.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsBingGrounding.ts","type":"file"},{"name":"agentsBingGroundingWithStreaming.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsBingGroundingWithStreaming.ts","type":"file"},{"name":"agentsConnectedAgents.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsConnectedAgents.ts","type":"file"},{"name":"agentsImageInputWithBase64.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsImageInputWithBase64.ts","type":"file"},{"name":"agentsImageInputWithFile.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsImageInputWithFile.ts","type":"file"},{"name":"agentsImageInputWithUrl.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsImageInputWithUrl.ts","type":"file"},{"name":"agentsLocalTelemetry.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsLocalTelemetry.ts","type":"file"},{"name":"agentsRemoteTelemetry.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsRemoteTelemetry.ts","type":"file"},{"name":"agentsWithFunctionTool.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsWithFunctionTool.ts","type":"file"},{"name":"agentsWithOpenApi.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsWithOpenApi.ts","type":"file"},{"name":"agentsWithToolset.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsWithToolset.ts","type":"file"},{"name":"batchVectorStoreWithFiles.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/batchVectorStoreWithFiles.ts","type":"file"},{"name":"batchVectorSto
... (truncated)
```

### ✅ 8. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":20,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-agents/samples \"createAndPoll\" \"messages.list\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"fileSearch.js","path":"sdk/ai/ai-agents/samples/v1-beta/javascript/fileSearch.js","sha":"58c260aca31a36888176a53a344c1a382baefe1e","text_matches":[{"fragment":"  console.log(\"Creating run...\");\n  const run = await client.runs.createAndPoll(thread.id, agent.id, {\n    pollingOptions: {","matches":[{"indices":[66,79],"text":"createAndPoll"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/samples/v1-beta/javascript/fileSearch.js?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"\n  const messages = await client.messages.list(thread.id);\n  for await (const threadMessage of messages) {","matches":[{"indices":[33,46],"text":"messages.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/samples/v1-beta/javascript/fileSearch.js?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"fileSearch.ts","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/fileSearch.ts","sha":"956afdca0f47da45f7dcb3654ea272f84962dee7","text_matches":[{"fragment":"  console.log(\"Creating run...\");\n  const run = await client.runs.createAndPoll(thread.id, agent.id, {\n    pollingOptions: {","matches":[{"indices":[66,79],"text":"createAndPoll"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/samples/v1-beta/typescript/src/fileSearch.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"\n  const messages = await client.messages.list(thread.id);\n  for await (const threadMessage of messages) {","matches":[{"indices":[33,46],"text":"messages.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-agents/samples/v1-beta/typescript/src/fileSearch.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"con
... (truncated)
```

### ✅ 9. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":20,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-agents/samples \"deleteVectorStore\" OR \"vectorStores.delete\" language:TypeScript"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 10. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/fileSearch.ts","ref":"refs/tags/@azure/ai-agents_1.1.0","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: 400db87cf648ca5af9a0232fa965c9ce093f75ea){
  "compilerOptions": {
    "target": "ES2023",
    "module": "commonjs",
    "lib": [],
    "importHelpers": true,
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "moduleResolution": "node10",
    "esModuleInterop": true,
    "outDir": "./dist",
    "resolveJsonModule": true
  },
  "include": [
    "./src"
  ]
}

```

### ✅ 11. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/vectorStoreWithFilesAndPolling.ts","ref":"refs/tags/@azure/ai-agents_1.1.0","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: 41852394937608b2697554b80aac4110885c20ff){
  "name": "@azure-samples/ai-agents-ts-beta",
  "private": true,
  "version": "1.0.0",
  "description": "Azure AI Agents client library samples for TypeScript (Beta)",
  "engines": {
    "node": ">=20.0.0"
  },
  "scripts": {
    "build": "tsc",
    "prebuild": "rimraf dist/"
  },
  "repository": {
    "type": "git",
    "url": "git+https://github.com/Azure/azure-sdk-for-js.git",
    "directory": "sdk/ai/ai-agents"
  },
  "keywords": [
    "node",
    "azure",
    "cloud",
    "typescript",
    "browser",
    "isomorphic"
  ],
  "author": "Microsoft Corporation",
  "license": "MIT",
  "bugs": {
    "url": "https://github.com/Azure/azure-sdk-for-js/issues"
  },
  "homepage": "https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/ai/ai-agents",
  "dependencies": {
    "@azure/ai-agents": "next",
    "dotenv": "latest",
    "@azure/identity": "^4.3.0",
    "@opentelemetry/api": "^1.9.0",
    "@opentelemetry/instrumentation": "0.57.0",
    "@azure/opentelemetry-instrumentation-azure-sdk": "^1.0.0-beta.7",
    "@opentelemetry/sdk-trace-node": "^1.30.0",
    "@azure/monitor-opentelemetry": "^1.11.1"
  },
  "devDependencies": {
    "@types/node": "^20.0.0",
    "typescript": "~5.8.2",
    "rimraf": "latest"
  }
}

```

### ✅ 12. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/src/agentsBasics.ts","ref":"refs/tags/@azure/ai-agents_1.1.0","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: 554d300e4c69db7e438717d7adb29b1c4b727cf9)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/**
 * This sample demonstrates how to use basic agent operations from the Azure Agents service.
 *
 * @summary demonstrates how to use basic agent operations.
 */

import { AgentsClient } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

import "dotenv/config";

const projectEndpoint = process.env["PROJECT_ENDPOINT"] || "<project endpoint>";
const modelDeploymentName = process.env["MODEL_DEPLOYMENT_NAME"] || "gpt-4o";

export async function main(): Promise<void> {
  // Create an Azure AI Client
  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());

  // Create an agent
  const agent = await client.createAgent(modelDeploymentName, {
    name: "my-agent",
    instructions: "You are helpful agent",
  });
  console.log(`Created agent, agent ID : ${agent.id}`);

  // Create a thread
  const thread = await client.threads.create();
  console.log(`Created thread, thread ID : ${thread.id}`);

  // List all threads for the agent
  const threads = client.threads.list();
  console.log(`Threads for agent ${agent.id}:`);
  for await (const t of threads) {
    console.log(`Thread ID: ${t.id}`);
    console.log(`Created at: ${t.createdAt}`);
    console.log(`Metadata: ${t.metadata}`);
    console.log(`---- `);
  }

  // Create a message
  const message = await client.messages.create(thread.id, "user", "Hello, tell me a joke");
  console.log(`Created message, message ID : ${message.id}`);

  // Create and poll a run
  console.log("Creating run...");
  const run = await client.runs.createAndPoll(thread.id, agent.id, {
    pollingOptions: {
      intervalInMs: 2000,
    },
    onResponse: (response): void => {
      const parsedBody =
        typeof response.parsedBody === "object" && response.parsedBody !== null
          ? response.parsedBody
          : null;
      
... (truncated)
```

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/package.json","ref":"refs/tags/@azure/ai-agents_1.1.0","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: ff156f2e4ea1b6041d453d3c034447b1d345854a)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/**
 * This sample demonstrates how to create the vector store with the list of files using polling operation.
 *
 * @summary demonstrates how to create the vector store with the list of files using polling operation.
 */

import { AgentsClient } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";
import { Readable } from "stream";
import "dotenv/config";

const projectEndpoint = process.env["PROJECT_ENDPOINT"] || "<project endpoint>";

export async function main(): Promise<void> {
  // Create an Azure AI Client
  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());

  // Create vector store
  const vectorStore = await client.vectorStores.create();
  console.log(`Created vector store, vector store ID: ${vectorStore.id}`);

  // Create and upload file
  const fileContent = "Hello, Vector Store!";
  const readable = new Readable();
  await readable.push(fileContent);
  await readable.push(null); // end the stream
  const file = await client.files.upload(readable, "assistants", {
    fileName: "vectorFile.txt",
  });
  console.log(`Uploaded file, file ID: ${file.id}`);

  // Create vector store file, which will automatically poll until the operation is complete
  const vectorStoreFile1 = await client.vectorStoreFiles.create(vectorStore.id, {
    fileId: file.id,
    pollingOptions: {
      intervalInMs: 2000,
    },
    // (Optional) Define an onResponse callback to monitor the progress of polling
    onResponse: (response): void => {
      console.log(`Received response with status: ${response.parsedBody.status}`);
    },
  });
  console.log(
    `Created vector store file with status ${vectorStoreFile1.status}, vector store file ID: ${vectorStoreFile1.id}`,
  );

  // Alternatively, polling can be done using .poll() and .pollUntilDone() methods.
  // This 
... (truncated)
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-agents/samples/v1-beta/typescript/tsconfig.json","ref":"refs/tags/@azure/ai-agents_1.1.0","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: cdb19a34bbbae84a843b15cbc31f8e32c62d8202)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/**
 * This sample demonstrates how to use agent operations with file searching from the Azure Agents service.
 *
 * @summary This sample demonstrates how to use agent operations with file searching.
 */

import type { MessageContent, MessageImageFileContent, MessageTextContent } from "@azure/ai-agents";
import { AgentsClient, isOutputOfType, ToolUtility } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

import * as fs from "fs";
import "dotenv/config";

const projectEndpoint = process.env["PROJECT_ENDPOINT"] || "<project endpoint>";
const modelDeploymentName = process.env["MODEL_DEPLOYMENT_NAME"] || "gpt-4o";

export async function main(): Promise<void> {
  // Create an Azure AI Client
  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());

  // Upload file
  const filePath = "./data/sampleFileForUpload.txt";
  const localFileStream = fs.createReadStream(filePath);
  const file = await client.files.upload(localFileStream, "assistants", {
    fileName: "sampleFileForUpload.txt",
  });
  console.log(`Uploaded file, file ID: ${file.id}`);

  // Create vector store
  const vectorStore = await client.vectorStores.create({
    fileIds: [file.id],
    name: "myVectorStore",
  });
  console.log(`Created vector store, vector store ID: ${vectorStore.id}`);

  // Initialize file search tool
  const fileSearchTool = ToolUtility.createFileSearchTool([vectorStore.id]);

  // Create agent with files
  const agent = await client.createAgent(modelDeploymentName, {
    name: "File Search Agent",
    instructions: "You are helpful agent that can help fetch data from files you know about.",
    tools: [fileSearchTool.definition],
    toolResources: fileSearchTool.resources,
  });
  console.log(`Created agent, agent ID : ${agent.id}`);

  // Create thread
  const thread 
... (truncated)
```

### ✅ 15. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"hyoka-trail-guide-agent\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"description\": \"A TypeScript console application using Azure AI Agents file search.\",\n+  \"type\": \"module\",\n+  \"engines\": {\n+    \"node\": \"\u003e=20.0.0\"\n+  },\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/ai-agents\": \"^1.1.0\",\n+    \"@azure/identity\": \"^4.11.1\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"noUnusedLocals\": true,\n+    \"noUnusedParameters\": true,\n+    \"noImplicitReturns\": true,\n+    \"noFallthroughCasesInSwitch\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+trail-guide.txt\n+.env\n*** Add File: README.md\n+# Azure AI Agents document-grounded console app\n+\n+## Install, build, and run\n+\n+```powershell\n+npm install\n+$env:PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:MODEL_DEPLOYMENT_NAME = \"\u003cmodel-deployment-name\u003e\"\n+npm run build\n+npm start\n+```\n+\n+Authentication uses `DefaultAzureCredential`. Configure any supported local credential, such as Azure CLI login, before running.\n*** Add File: src/index.ts\n+import { AgentsClient, isOutputOfType, ToolUtility } from \"@azure/ai-agents\";\n+import type { MessageTextContent } from \"@azure/ai-agents\";\n+import { DefaultAzureCredential } from \"@azure/identity\";\n+import { createReadStream } from \"node:fs\";\n+import { writeFile } from \"node:fs/promises\";\n+import { resolve } from \"node:path\";\n+\n+const documentText =\n+  \"The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.\";\n+const question =\n+  \"According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?\";\n+const documentPath = resolve(\"trail-guide.txt\");\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const projectEndpoint = requireEnvironmentVariable(\"PROJECT_ENDPOINT\");\n+  const modelDeploymentName = requireEnvironmentVariable(\"MODEL_DEPLOYMENT_NAME\");\n+  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());\n+\n+  let uploadedFileId: string | undefined;\n+  let vectorStoreId: string | undefined;\n+  let agentId: string | undefined;\n+  let threadId: string | undefined;\n+\n+  await writeFile(documentPath, documentText, \"utf8\");\n+\n+  try {\n+    const uploadedFile = await client.files.upload(\n+      createReadStream(documentPath),\n+      \"assistants\",\n+      { fileName: \"trail-guide.txt\" },\n+    );\n+    uploadedFileId = uploadedFile.id;\n+\n+    const vectorStore = await client.vectorStores.create({\n+      name: \"hyoka-trail-guide-vector-store\",\n+    });\n+    vectorStoreId = vectorStore.id;\n+\n+    const indexedFile = await client.vectorStoreFiles\n+      .createAndPoll(vectorStore.id, {\n+        fileId: uploadedFile.id,\n+        pollingOptions: { intervalInMs: 2_000 },\n+      })\n+      .pollUntilDone();\n+\n+    if (indexedFile.status !== \"completed\") {\n+      throw new Error(\n+        `Document indexing failed with status \"${indexedFile.status}\"` +\n+          (indexedFile.lastError ? `: ${indexedFile.lastError.message}` : \"\"),\n+      );\n+    }\n+\n+    const fileSearchTool = ToolUtility.createFileSearchTool([vectorStore.id]);\n+    const agent = await client.createAgent(modelDeploymentName, {\n+      name: \"hyoka-trail-guide-agent\",\n+      instructions:\n+        \"Answer questions using the uploaded guide. Base factual claims on file search results.\",\n+      tools: [fileSearchTool.definition],\n+      toolResources: fileSearchTool.resources,\n+    });\n+    agentId = agent.id;\n+\n+    const thread = await client.threads.create();\n+    threadId = thread.id;\n+    await client.messages.create(thread.id, \"user\", question);\n+\n+    const run = await client.runs.createAndPoll(thread.id, agent.id, {\n+      pollingOptions: { intervalInMs: 2_000 },\n+    });\n+\n+    if (run.status !== \"completed\") {\n+      throw new Error(\n+        `Agent run ended with status \"${run.status}\"` +\n+          (run.lastError ? `: ${run.lastError.message}` : \"\"),\n+      );\n+    }\n+\n+    const messages = [];\n+    for await (const message of client.messages.list(thread.id, {\n+      order: \"asc\",\n+    })) {\n+      messages.push(message);\n+    }\n+\n+    for (const message of messages) {\n+      if (message.role !== \"assistant\") {\n+        continue;\n+      }\n+      for (const content of message.content) {\n+        if (isOutputOfType\u003cMessageTextContent\u003e(content, \"text\")) {\n+          console.log(content.text.value);\n+        }\n+      }\n+    }\n+  } finally {\n+    const cleanupErrors: unknown[] = [];\n+\n+    const cleanUp = async (operation: () =\u003e Promise\u003cunknown\u003e): Promise\u003cvoid\u003e =\u003e {\n+      try {\n+        await operation();\n+      } catch (error: unknown) {\n+        cleanupErrors.push(error);\n+      }\n+    };\n+\n+    if (threadId) {\n+      await cleanUp(() =\u003e client.threads.delete(threadId));\n+    }\n+    if (agentId) {\n+      await cleanUp(() =\u003e client.deleteAgent(agentId));\n+    }\n+    if (vectorStoreId) {\n+      await cleanUp(() =\u003e client.vectorStores.delete(vectorStoreId));\n+    }\n+    if (uploadedFileId) {\n+      await cleanUp(() =\u003e client.files.delete(uploadedFileId));\n+    }\n+\n+    if (cleanupErrors.length \u003e 0) {\n+      throw new AggregateError(cleanupErrors, \"One or more Azure resources could not be deleted\");\n+    }\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(error);\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/package.json
@@ -1,0 +1,23 @@
+{
+  "name": "hyoka-trail-guide-agent",
+  "version": "1.0.0",
+  "private": true,
+  "description": "A TypeScript console application using Azure AI Agents file search.",
+  "type": "module",
+  "engines": {
+    "node": ">=20.0.0"
+  },
+  "scripts": {
+    "build": "tsc",
+    "start": "node dist/index.js"
+  },
+  "dependencies": {
+    "@azure/ai-agents": "^1.1.0",
+    "@azure/identity": "^4.11.1"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/tsconfig.json
@@ -1,0 +1,18 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "outDir": "dist",
+    "rootDir": "src",
+    "strict": true,
+    "noUnusedLocals": true,
+    "noUnusedParameters": true,
+    "noImplicitReturns": true,
+    "noFallthroughCasesInSwitch": true,
+    "esModuleInterop": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/.gitign
... (truncated)
```

### ✅ 16. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 51 packages in 53s

> hyoka-trail-guide-agent@1.0.0 build
> tsc

src/index.ts(115,49): error TS2345: Argument of type 'string | undefined' is not assignable to parameter of type 'string'.
  Type 'undefined' is not assignable to type 'string'.
src/index.ts(118,46): error TS2345: Argument of type 'string | undefined' is not assignable to parameter of type 'string'.
  Type 'undefined' is not assignable to type 'string'.
src/index.ts(121,54): error TS2345: Argument of type 'string | undefined' is not assignable to parameter of type 'string'.
  Type 'undefined' is not assignable to type 'string'.
src/index.ts(124,47): error TS2345: Argument of type 'string | undefined' is not assignable to parameter of type 'string'.
  Type 'undefined' is not assignable to type 'string'.
<shellId: 0 completed with exit code 2>
```

### ✅ 17. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n     };\n \n     if (threadId) {\n-      await cleanUp(() =\u003e client.threads.delete(threadId));\n+      const id = threadId;\n+      await cleanUp(() =\u003e client.threads.delete(id));\n     }\n     if (agentId) {\n-      await cleanUp(() =\u003e client.deleteAgent(agentId));\n+      const id = agentId;\n+      await cleanUp(() =\u003e client.deleteAgent(id));\n     }\n     if (vectorStoreId) {\n-      await cleanUp(() =\u003e client.vectorStores.delete(vectorStoreId));\n+      const id = vectorStoreId;\n+      await cleanUp(() =\u003e client.vectorStores.delete(id));\n     }\n     if (uploadedFileId) {\n-      await cleanUp(() =\u003e client.files.delete(uploadedFileId));\n+      const id = uploadedFileId;\n+      await cleanUp(() =\u003e client.files.delete(id));\n     }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-agents-dp-js-ts-file-search-js-ts-azure-tools-baseline-2603841768/src/index.ts
@@ -112,16 +112,20 @@
     };
 
     if (threadId) {
-      await cleanUp(() => client.threads.delete(threadId));
+      const id = threadId;
+      await cleanUp(() => client.threads.delete(id));
     }
     if (agentId) {
-      await cleanUp(() => client.deleteAgent(agentId));
+      const id = agentId;
+      await cleanUp(() => client.deleteAgent(id));
     }
     if (vectorStoreId) {
-      await cleanUp(() => client.vectorStores.delete(vectorStoreId));
+      const id = vectorStoreId;
+      await cleanUp(() => client.vectorStores.delete(id));
     }
     if (uploadedFileId) {
-      await cleanUp(() => client.files.delete(uploadedFileId));
+      const id = uploadedFileId;
+      await cleanUp(() => client.files.delete(id));
     }
 
     if (cleanupErrors.length > 0) {

```

### ✅ 18. powershell

**Input:**

```json
{"command":"npm run build","description":"Compile the TypeScript application","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> hyoka-trail-guide-agent@1.0.0 build
> tsc
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace and confirm the current `@azure/ai-agents` API before creating the runnable project.

The workspace is empty. I’m narrowing the implementation to the SDK’s published TypeScript samples so the method names, polling, message ordering, and cleanup match the released API.

The released API supports polling vector-store ingestion directly and exposes the required delete operations. I’ll now create the project with failure-safe cleanup and explicit terminal-status checks, then compile it against the installed SDK types.

Compilation found TypeScript’s closure-narrowing limitation in cleanup callbacks. I’m making the captured IDs explicitly non-optional and rebuilding.

Created a complete TypeScript console application in the workspace.

- Uses `@azure/ai-agents` with document upload, vector-store indexing, file search, agent/thread/run polling, chronological assistant output, and dependency-ordered cleanup.
- Includes `package.json`, lockfile, `tsconfig.json`, `.gitignore`, and concise setup instructions in `README.md`.
- Build succeeds with `npm run build`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete TypeScript project manifest with the necessary Azure SDK, authentication, compiler, and Node.js type dependencies.

## Grader Results

- file-search.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Uploads a readable document stream with the `assistants` purpose and a filename.: Pass
      - Creates a vector store, adds the uploaded file with the SDK poller, and requires a: Pass
      - Uses `ToolUtility.createFileSearchTool` with the vector-store ID and passes both: Pass
      - Creates the thread, exact grounded question, and run with the created IDs.: Pass
      - Polls the run with `createAndPoll`, requires completed status, and retrieves: Pass
      - Deletes the thread before the agent, then deletes the vector store before the: Pass
      - Does not put a local file path or uploaded file ID directly in the user message as: Pass
      - Does not start the run before vector-store file indexing completes.: Pass
      - Does not print the known document fact directly as though it were the agent's: Pass
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct @azure/ Scoped Packages**: All Azure SDK dependencies use @azure/ scoped packages (e.g., @azure/storage-blob, @azure/keyvault-secrets, @azure/identity). No deprecated unscoped packages (azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-*).: Pass
  - @azure/identity for Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**@azure/identity for Authentication**: Includes @azure/identity as a dependency. For local development, DefaultAzureCredential is acceptable. For production code, uses a specific Entra token credential (e.g., ManagedIdentityCredential, ClientSecretCredential, WorkloadIdentityCredential) or configures DefaultAzureCredential with the AZURE_TOKEN_CREDENTIALS environment variable set to "prod" or a specific token credential. No hardcoded connection strings, account keys, SAS tokens, or secrets in code.: Pass
  - Client Constructor with Endpoint and Credential (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Constructor with Endpoint and Credential**: SDK clients constructed with endpoint URL and credential object, e.g. new BlobServiceClient(url, credential) or new SecretClient(vaultUrl, credential). No legacy constructor patterns or connection string-based construction when identity-based auth is available.: Pass
  - Async/Await Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Pass
  - Pagination with for-await-of (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Fail
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Fail
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
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Pagination with for-await-of` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 14.3% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 7.0000** | **63.6%** | |

## Re-run Command

```bash
hyoka run --prompt-id ai-agents-dp-js-ts-file-search --config js-ts-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
