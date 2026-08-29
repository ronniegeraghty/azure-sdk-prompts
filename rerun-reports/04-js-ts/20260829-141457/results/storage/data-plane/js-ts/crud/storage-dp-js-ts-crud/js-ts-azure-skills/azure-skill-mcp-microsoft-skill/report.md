# Evaluation Report: storage-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 362.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/18 |
| Duration | 362.1s |
| Timestamp | 2026-08-29T10:02:58Z |
| Files Generated | 5 |
| Event Count | 4194 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 181.5s |
| Review | 179.3s |
| **Total** | **362.1s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage-blob-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=713935 out=4228 |
| Turn Count | 14 |

## Tool Calls

### ✅ 1. skill — 3268ms

**Input:**

```json
{"skill":"azure-storage-blob-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# @azure/storage-blob (TypeScript/JavaScript)

SDK for Azure Blob Storage operations — upload, download, list, and manage blobs and containers.

## Installation

```bash
npm install @azure/storage-blob @azure/identity
```

**Current Version**: 12.x  
**Node.js**: >= 18.0.0

## Environment Variables

```bash
AZURE_STORAGE_ACCOUNT_NAME=<account-name>
AZURE_STORAGE_ACCOUNT_KEY=<account-key>
# OR connection string
AZURE_STORAGE_CONNECTION_STRING=DefaultEndpointsProtocol=https;AccountName=...
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

### Microsoft Entra Token Credential (Recommended)

```typescript
import { BlobServiceClient } from "@azure/storage-blob";
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME!;
const client = new BlobServiceClient(
  `https://${accountName}.blob.core.windows.net`,
  credential
);
```

### Connection String

```typescript
import { BlobServiceClient } from "@azure/storage-blob";

const client = BlobServiceClient.fromConnectionString(
  process.env.AZURE_STORAGE_CONNECTION_STRING!
);
```

### StorageSharedKeyCredential (Node.js only)

```typescript
import { BlobServiceClient, StorageSharedKeyCredential } from "@azure/storage-blob";

const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME!;
const accountKey = process.env.AZURE_STORAGE_ACCOUNT_KEY!;

const sharedKeyCredential = new StorageShar
... (truncated)
```

### ✅ 2. glob — 8580ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 8249ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-storage-blob-ts","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 8912ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-storage-blob-ts\\references\\streaming.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md
@@ -1,426 +1,426 @@
 # @azure/storage-blob - Streaming Patterns
 
 Reference documentation for upload/download streaming in the Azure Blob Storage TypeScript SDK.
 
 **Source**: [Azure SDK for JS - storage-blob](https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/storage/storage-blob)
 
 ---
 
 ## Installation
 
 ```bash
 npm install @azure/storage-blob @azure/identity
 ```
 
 ---
 
 ## Client Setup
 
 ```typescript
 import {
   BlobServiceClient,
   ContainerClient,
   BlockBlobClient,
 } from "@azure/storage-blob";
 import { DefaultAzureCredential } from "@azure/identity";
 
 const credential = new DefaultAzureCredential();
 const accountUrl = `https://${process.env["STORAGE_ACCOUNT_NAME"]}.blob.core.windows.net`;
 
 const blobServiceClient = new BlobServiceClient(accountUrl, credential);
 const containerClient = blobServiceClient.getContainerClient("my-container");
 const blobClient = containerClient.getBlockBlobClient("my-blob.txt");
 ```
 
 ---
 
 ## Download Streaming
 
 ### Download to Buffer
 
 ```typescript
 const downloadResponse = await blobClient.download(0);
 const downloaded = await streamToBuffer(downloadResponse.readableStreamBody!);
 
 async function streamToBuffer(readableStream: NodeJS.ReadableStream): Promise<Buffer> {
   return new Promise((resolve, reject) => {
     const chunks: Buffer[] = [];

... (truncated)
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 2795ms

**Input:**

```json
{"intent":"Generate a TypeScript Node.js CRUD sample for Azure Blob Storage using BlobServiceClient, DefaultAzureCredential, async/await, and RestError handling. Do not deploy or contact Azure.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 3141ms

**Input:**

```json
{"intent":"Find official TypeScript guidance for Azure Blob Storage BlobServiceClient with DefaultAzureCredential, createIfNotExists, block blob upload, listBlobsFlat, download as string, delete blob/container, and RestError handling.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2672ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure TypeScript code-generation best practices for a Blob Storage CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3350ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official TypeScript Blob Storage CRUD samples.","parameters":{"language":"typescript","query":"Azure Blob Storage TypeScript BlobServiceClient DefaultAzureCredential createIfNotExists upload listBlobsFlat download delete RestError"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates how to upload a blob to Azure Blob Storage using managed identities for authentication in a JavaScript web app.\npackage: @azure/identity, @azure/storage-blob\nlanguage: javascript\n","codeSnippet":"const { DefaultAzureCredential } = require(\u0022@azure/identity\u0022);\r\nconst { BlobServiceClient } = require(\u0022@azure/storage-blob\u0022);\r\nconst defaultAzureCredential = new DefaultAzureCredential();\r\n\r\n// Some code omitted for brevity.\r\n\r\nasync function uploadBlob(accountName, containerName, blobName, blobContents) {\r\n    const blobServiceClient = new BlobServiceClient(\r\n        \u0060https://${accountName}.blob.core.windows.net\u0060,\r\n        defaultAzureCredential\r\n    );\r\n\r\n    const containerClient = blobServiceClient.getContainerClient(containerName);\r\n\r\n    try {\r\n        await containerClient.createIfNotExists();\r\n        const blockBlobClient = containerClient.getBlockBlobClient(blobName);\r\n        const uploadBlobResponse = await blockBlobClient.upload(blobContents, blobContents.length);\r\n        console.log(\u0060Upload block blob ${blobName} successfully\u0060, uploadBlobResponse.requestId);\r\n    } catch (error) {\r\n        console.log(error);\r\n    }\r\n}","link":"https://learn.microsoft.com/azure/app-service/tutorial-connect-app-access-storage-javascript#javascript-example","language":"javascript"},{"description":"description: Creates a DefaultAzureCredential with Azure Government authority host and uses it to instantiate a BlobServiceClient for connecting to Azure Blob Storage in Azure Government, then creates a container if it does not exist.\npackage: @azure/storage-blob, @azure/identity\nlanguage: javascript\n","codeSnippet":"const { BlobServiceClient } = require(\u0027@azure/storage-blob\u0027);\r\nconst {\r\n DefaultAzureCredential,\r\n DefaultAzureCredentialOptions,\r\n AzureAuthorityHosts\r\n} = require(\u0027@azure/identity\u0027);\r\n\r\nconst cr
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3010ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure Blob Storage TypeScript CRUD and authorization guidance.","parameters":{"query":"Azure Blob Storage JavaScript TypeScript CRUD DefaultAzureCredential RestError create container upload list download delete"}}
```

**Output:**

```
{"results":[{"title":"Quickstart: Azure Blob Storage client library for Node.js with TypeScript","content":"# Quickstart: Azure Blob Storage client library for Node.js with TypeScript\n## Object model\nAzure Blob storage is optimized for storing massive amounts of unstructured data. Unstructured data is data that doesn\u0027t adhere to a particular data model or definition, such as text or binary data. Blob storage offers three types of resources:\n1. The storage account\n2. A container in the storage account\n3. A blob in the container\nThe following diagram shows the relationship between these resources.\n![Diagram of Blob storage architecture.](https://learn.microsoft.com/azure/storage/blobs/media/storage-blobs-introduction/blob1.png)\nUse the following JavaScript classes to interact with these resources:\n1. [BlobServiceClient](https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blobserviceclient): The \u0060BlobServiceClient\u0060 class allows you to manipulate Azure Storage resources and blob containers.\n2. [ContainerClient](https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/containerclient): The \u0060ContainerClient\u0060 class allows you to manipulate Azure Storage containers and their blobs.\n3. [BlobClient](https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blobclient): The \u0060BlobClient\u0060 class allows you to manipulate Azure Storage blobs.\n## Code examples\nThese example code snippets show you how to do the following tasks with the Azure Blob Storage client library for JavaScript:\n1. Authenticate to Azure and authorize access to blob data\n2. Create a container\n3. Upload blobs to a container\n4. List the blobs in a container\n5. Download blobs\n6. Delete a container\nSample code is also available on [GitHub](https://github.com/Azure-Samples/AzureStorageSnippets/tree/master/blobs/quickstarts/JavaScript/V12/nodejs).\n### Authenticate to Azure and authorize access to blob data\nApplication requ
... (truncated)
```

### ✅ 10. powershell — 2028ms

**Input:**

```json
{"command":"$data = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787997829703-copilot-tool-output-2c89bef1005b4006b2b842979089c2cb.txt' | ConvertFrom-Json; $data.results | Select-Object -First 8 title,url | Format-Table -AutoSize","description":"Extract official documentation links"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-2c89bef1005b4006b2b842979089c2cb.txt' | ConvertFrom-Json; $data. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 11. powershell — 1672ms

**Input:**

```json
{"command":"$data = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787997830093-copilot-tool-output-59d6a2b5c2434298a1a1fc5d78c80071.txt' | ConvertFrom-Json; $data.results | Select-Object -First 8 title,url | Format-Table -AutoSize","description":"Extract official sample links"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-59d6a2b5c2434298a1a1fc5d78c80071.txt' | ConvertFrom-Json; $data. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 12. view — 1432ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787997829703-copilot-tool-output-2c89bef1005b4006b2b842979089c2cb.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787997830093-copilot-tool-output-59d6a2b5c2434298a1a1fc5d78c80071.txt b/C:/Users/weidxu/AppData/Local/Temp/1787997830093-copilot-tool-output-59d6a2b5c2434298a1a1fc5d78c80071.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787997830093-copilot-tool-output-59d6a2b5c2434298a1a1fc5d78c80071.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787997830093-copilot-tool-output-59d6a2b5c2434298a1a1fc5d78c80071.txt
@@ -1,6 +1,6 @@
 {"results":[{"description":"description: Demonstrates how to upload a blob to Azure Blob Storage using managed identities for authentication in a JavaScript web app.\npackage: @azure/identity, @azure/storage-blob\nlanguage: javascript\n","codeSnippet":"const { DefaultAzureCredential } = require(\u0022@azure/identity\u0022);\r\nconst { BlobServiceClient } = require(\u0022@azure/storage-blob\u0022);\r\nconst defaultAzureCredential = new DefaultAzureCredential();\r\n\r\n// Some code omitted for brevity.\r\n\r\nasync function uploadBlob(accountName, containerName, blobName, blobContents) {\r\n    const blobServiceClient = new BlobServiceClient(\r\n        \u0060https://${accountName}.blob.core.windows.net\u0060,\r\n        defaultAzureCredential\r\n    );\r\n\r\n    const containerClient = blobServiceClient.getContainerClient(containerName);\r\n\r\n    try {\r\n        await containerClient.createIfNotExists();\r\n        const blockBlobClient = containerClient.getBlockBlobClient(blobName);\r\n        const uploadBlobResponse = await blockBlobClient.upload(blobContents, blobContents.length);\r\n        console.log(\u0060Upload block blob ${blobName} successfully\u0060, uploadBlobResponse.requestId);\r\n    } catch (error) {\r\n        console.log(error);\r\n    }\r\n}","link":"https://learn.microsoft.com/azure/app-service/tutorial-connect-app-access-storage-javascript#javascript-example","language":"javascript"},{"description":"description: Creates a DefaultAzureCredential with Az
... (truncated)
```

### ✅ 13. view — 1077ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787997830093-copilot-tool-output-59d6a2b5c2434298a1a1fc5d78c80071.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787997829703-copilot-tool-output-2c89bef1005b4006b2b842979089c2cb.txt b/C:/Users/weidxu/AppData/Local/Temp/1787997829703-copilot-tool-output-2c89bef1005b4006b2b842979089c2cb.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787997829703-copilot-tool-output-2c89bef1005b4006b2b842979089c2cb.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787997829703-copilot-tool-output-2c89bef1005b4006b2b842979089c2cb.txt
@@ -1,6 +1,6 @@
 {"results":[{"title":"Quickstart: Azure Blob Storage client library for Node.js with TypeScript","content":"# Quickstart: Azure Blob Storage client library for Node.js with TypeScript\n## Object model\nAzure Blob storage is optimized for storing massive amounts of unstructured data. Unstructured data is data that doesn\u0027t adhere to a particular data model or definition, such as text or binary data. Blob storage offers three types of resources:\n1. The storage account\n2. A container in the storage account\n3. A blob in the container\nThe following diagram shows the relationship between these resources.\n![Diagram of Blob storage architecture.](https://learn.microsoft.com/azure/storage/blobs/media/storage-blobs-introduction/blob1.png)\nUse the following JavaScript classes to interact with these resources:\n1. [BlobServiceClient](https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blobserviceclient): The \u0060BlobServiceClient\u0060 class allows you to manipulate Azure Storage resources and blob containers.\n2. [ContainerClient](https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/containerclient): The \u0060ContainerClient\u0060 class allows you to manipulate Azure Storage containers and their blobs.\n3. [BlobClient](https://learn.microsoft.com/en-us/javascript/api/@azure/storage-blob/blobclient): The \u0060BlobClient\u0060 class allows you to manipulate Azure Storage blobs.\n## Code examples\nThese example code snippets show you how to do 
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 1226ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official TypeScript Blob Storage CRUD quickstart before implementing the sample.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-nodejs-typescript"}}
```

**Output:**

```
# Quickstart: Azure Blob Storage client library for Node.js with TypeScript

Get started with the Azure Blob Storage client library for Node.js with TypeScript to manage blobs and containers.

In this article, you follow steps to install the package and try out example code for basic tasks.

[API reference](/en-us/javascript/api/@azure/storage-blob) | [Library source code](https://github.com/Azure/azure-sdk-for-js/tree/master/sdk/storage/storage-blob) | [Package (npm)](https://www.npmjs.com/package/@azure/storage-blob) | [Samples](../common/storage-samples-javascript?toc=/azure/storage/blobs/toc.json#blob-samples)

## Prerequisites

- Azure account with an active subscription - [create an account for free](https://azure.microsoft.com/pricing/purchase-options/azure-account?cid=msft_learn)
- Azure Storage account - [Create a storage account](../common/storage-account-create)
- [Node.js LTS](https://nodejs.org/en/download/package-manager/)
- [TypeScript](https://www.typescriptlang.org/download)

## Setting up

This section walks you through preparing a project to work with the Azure Blob Storage client library for Node.js.

### Create the Node.js project

Create a TypeScript application named *blob-quickstart*.

1. In a console window (such as cmd, PowerShell, or Bash), create a new directory for the project:

    ```console
    mkdir blob-quickstart
    ```
2. Switch to the newly created *blob-quickstart* directory:

    ```console
    cd blob-quickstart
    ```
3. Create a *package.json* file:

    ```console
    npm init -y
    ```
4. Open the project in Visual Studio Code:

    ```console
    code .
    ```
5. Edit the *package.json* file to add the following properties to support ESM with TypeScript:

    ```json
    "type": "module",
    ```

### Install the packages

From the project directory, install the following packages using the `npm install` command.

1. Install the Azure Storage npm package:

    ```
... (truncated)
```

### ✅ 15. powershell — 15178ms

**Input:**

```json
{"command":"npm view @azure/storage-blob version; npm view @azure/identity version; npm view typescript version; npm view @types/node version","description":"Check current npm package versions","initial_wait":120}
```

**Output:**

```
12.33.0
4.13.2
7.0.2
26.2.0
<shellId: 2 completed with exit code 0>
```

### ✅ 16. rg — 14860ms

**Input:**

```json
{"-n":true,"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787997847827-copilot-tool-output-3c373f23797f448db0798cfc26f37dc8.txt","pattern":"RestError|createIfNotExists|listBlobsFlat|readableStreamBody"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787997847827-copilot-tool-output-3c373f23797f448db0798cfc26f37dc8.txt:476:List the blobs in the container. The following code calls the [listBlobsFlat](/en-us/javascript/api/@azure/storage-blob/containerclient#listblobsflat-containerlistblobsoptions-) method. In this case, only one blob is in the container, so the listing operation returns just that one blob.
C:\Users\weidxu\AppData\Local\Temp\1787997847827-copilot-tool-output-3c373f23797f448db0798cfc26f37dc8.txt:483:  for await (const blob of containerClient.listBlobsFlat()) {
C:\Users\weidxu\AppData\Local\Temp\1787997847827-copilot-tool-output-3c373f23797f448db0798cfc26f37dc8.txt:508:    await streamToText(downloadBlockBlobResponse.readableStreamBody as NodeJS.ReadableStream)
```

### ✅ 17. apply_patch — 778ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-blob-crud-typescript\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.13.2\",\n+    \"@azure/storage-blob\": \"^12.33.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import { DefaultAzureCredential } from \"@azure/identity\";\n+import { BlobServiceClient, RestError } from \"@azure/storage-blob\";\n+import type { Readable } from \"node:stream\";\n+\n+const containerName = \"my-container\";\n+const blobName = \"greeting.txt\";\n+const blobContent = \"Hello Azure!\";\n+\n+async function streamToString(stream: Readable): Promise\u003cstring\u003e {\n+  const chunks: Buffer[] = [];\n+\n+  for await (const chunk of stream) {\n+    chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));\n+  }\n+\n+  return Buffer.concat(chunks).toString(\"utf8\");\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME;\n+  if (!accountName) {\n+    throw new Error(\n+      \"AZURE_STORAGE_ACCOUNT_NAME must contain the Azure Storage account name.\",\n+    );\n+  }\n+\n+  const credential = new DefaultAzureCredential();\n+  const blobServiceClient = new BlobServiceClient(\n+    `https://${accountName}.blob.core.windows.net`,\n+    credential,\n+  );\n+\n+  const containerClient =\n+    blobServiceClient.getContainerClient(containerName);\n+  await containerClient.createIfNotExists();\n+  console.log(`Container ready: ${containerName}`);\n+\n+  const blockBlobClient = containerClient.getBlockBlobClient(blobName);\n+  await blockBlobClient.upload(\n+    blobContent,\n+    Buffer.byteLength(blobContent, \"utf8\"),\n+    {\n+      blobHTTPHeaders: {\n+        blobContentType: \"text/plain; charset=utf-8\",\n+      },\n+    },\n+  );\n+  console.log(`Uploaded: ${blobName}`);\n+\n+  console.log(\"Blobs:\");\n+  for await (const blob of containerClient.listBlobsFlat()) {\n+    console.log(`- ${blob.name}`);\n+  }\n+\n+  const downloadResponse = await blockBlobClient.download();\n+  if (!downloadResponse.readableStreamBody) {\n+    throw new Error(`The download for \"${blobName}\" returned no content stream.`);\n+  }\n+\n+  const downloadedContent = await streamToString(\n+    downloadResponse.readableStreamBody,\n+  );\n+  console.log(`Downloaded content: ${downloadedContent}`);\n+\n+  await blockBlobClient.delete();\n+  console.log(`Deleted blob: ${blobName}`);\n+\n+  await containerClient.delete();\n+  console.log(`Deleted container: ${containerName}`);\n+}\n+\n+void main().catch((error: unknown) =\u003e {\n+  if (error instanceof RestError) {\n+    console.error(\"Azure Blob Storage request failed.\", {\n+      message: error.message,\n+      statusCode: error.statusCode,\n+      code: error.code,\n+      requestId: error.request?.headers.get(\"x-ms-request-id\"),\n+    });\n+  } else if (error instanceof Error) {\n+    console.error(`Application error: ${error.message}`);\n+  } else {\n+    console.error(\"An unknown error occurred.\", error);\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Blob Storage CRUD with TypeScript\n+\n+This sample creates a container, uploads and lists a block blob, downloads its\n+content, and deletes both the blob and container.\n+\n+## Install\n+\n+```powershell\n+npm install\n+```\n+\n+Runtime packages:\n+\n+- `@azure/storage-blob`\n+- `@azure/identity`\n+\n+Development packages:\n+\n+- `typescript`\n+- `@types/node`\n+\n+## Configure and run\n+\n+The authenticated identity needs the **Storage Blob Data Contributor** role (or\n+equivalent data-plane permissions) on the storage account.\n+\n+```powershell\n+$env:AZURE_STORAGE_ACCOUNT_NAME = \"\u003cstorage-account-name\u003e\"\n+$env:AZURE_TOKEN_CREDENTIALS = \"dev\"\n+npm run build\n+npm start\n+```\n+\n+`DefaultAzureCredential` can use a supported local developer credential. In an\n+Azure-hosted production environment, use managed identity and set\n+`AZURE_TOKEN_CREDENTIALS` to `prod`.\n+\n+Reference: [Azure Blob Storage client library quickstart for Node.js with\n+TypeScript](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-nodejs-typescript)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "azure-blob-crud-typescript",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "node dist/index.js"
+  },
+  "dependencies": {
+    "@azure/identity": "^4.13.2",
+    "@azure/storage-blob": "^12.33.0"
+  },
+  "devDependencies": {
+    "@types/node": "^26.2.0",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/tsconfig.json
@@ -1,0 +1,16 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "outDir": "dist",
+    "rootDir": "src",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-
... (truncated)
```

### ✅ 18. powershell — 23061ms

**Input:**

```json
{"command":"npm install --silent \u0026\u0026 npm run build --silent","description":"Install dependencies and type-check sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
src/index.ts(61,5): error TS2740: Type 'NodeJSReadableStream' is missing the following properties from type 'Readable': readableAborted, readableDidRead, readableEncoding, readableEnded, and 25 more.
<shellId: 3 completed with exit code 2>
```

### ✅ 19. apply_patch — 910ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n import { DefaultAzureCredential } from \"@azure/identity\";\n import { BlobServiceClient, RestError } from \"@azure/storage-blob\";\n-import type { Readable } from \"node:stream\";\n@@\n-async function streamToString(stream: Readable): Promise\u003cstring\u003e {\n+async function streamToString(\n+  stream: NodeJS.ReadableStream,\n+): Promise\u003cstring\u003e {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-3117592038/src/index.ts
@@ -1,12 +1,13 @@
 import { DefaultAzureCredential } from "@azure/identity";
 import { BlobServiceClient, RestError } from "@azure/storage-blob";
-import type { Readable } from "node:stream";
 
 const containerName = "my-container";
 const blobName = "greeting.txt";
 const blobContent = "Hello Azure!";
 
-async function streamToString(stream: Readable): Promise<string> {
+async function streamToString(
+  stream: NodeJS.ReadableStream,
+): Promise<string> {
   const chunks: Buffer[] = [];
 
   for await (const chunk of stream) {

```

### ✅ 20. powershell — 1908ms

**Input:**

```json
{"command":"npm run build --silent","description":"Type-check Blob Storage sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```

<shellId: 4 completed with exit code 0>
```

### ✅ 21. powershell — 21256ms

**Input:**

```json
{"command":"Remove-Item Env:AZURE_STORAGE_ACCOUNT_NAME -ErrorAction SilentlyContinue; $output = npm start --silent 2\u003e\u00261; if ($LASTEXITCODE -ne 1 -or ($output -notmatch 'AZURE_STORAGE_ACCOUNT_NAME')) { $output; throw 'Expected configuration error was not produced.' }; $output","description":"Verify offline configuration error handling","initial_wait":120,"mode":"sync"}
```

**Output:**

```
Application error: AZURE_STORAGE_ACCOUNT_NAME must contain the Azure Storage account name.
<shellId: 5 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll align the sample with the Azure SDK guidance, then add a locally type-checkable TypeScript project without contacting Azure.

The implementation will use `AZURE_STORAGE_ACCOUNT_NAME`, passwordless authentication, a byte-accurate upload length, streamed download conversion, and typed `unknown` error handling that distinguishes `RestError`.

Implemented a complete TypeScript Azure Blob Storage CRUD sample in `src/index.ts`.

- Uses `BlobServiceClient` with `DefaultAzureCredential`
- Creates `my-container` if needed
- Uploads `Hello Azure!` as `greeting.txt`
- Lists blobs and downloads content as UTF-8
- Deletes the blob and container
- Handles `RestError` with status code, error code, and request ID
- Includes strict TypeScript configuration and npm scripts
- Documents authentication, required packages, and execution in `README.md`

Run with:

```powershell
$env:AZURE_STORAGE_ACCOUNT_NAME = "<storage-account-name>"
$env:AZURE_TOKEN_CREDENTIALS = "dev"
npm install
npm run build
npm start
```

The identity must have **Storage Blob Data Contributor** permissions. Reference: [Azure Blob Storage TypeScript quickstart](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-nodejs-typescript).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete dependency manifest with required Azure SDK packages
- Includes TypeScript compiler and Node.js type definitions
- Provides a valid strict TypeScript configuration

## Grader Results

- crud-blobs.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - Installing `@azure/storage-blob` and `@azure/identity` packages: Pass
      - `BlobServiceClient` construction with `DefaultAzureCredential`: Pass
      - `ContainerClient.createIfNotExists()`: Pass
      - `BlockBlobClient.upload()` or `uploadData()` for string content: Pass
      - `ContainerClient.listBlobsFlat()` async iteration: Pass
      - `BlobClient.download()` and reading the response stream: Pass
      - `BlobClient.delete()` and `ContainerClient.delete()`: Pass
      - `RestError` handling with `statusCode`: Pass
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
  - Async/Await Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Fail
  - Pagination with for-await-of (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Pass
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Pass
  - RestError Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**RestError Exception Handling**: Catches RestError (from @azure/core-rest-pipeline or the service package) with statusCode inspection for error-specific handling. Not just generic Error or try/catch without type checking.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `RestError Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 11.1% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 9.0000** | **81.8%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
