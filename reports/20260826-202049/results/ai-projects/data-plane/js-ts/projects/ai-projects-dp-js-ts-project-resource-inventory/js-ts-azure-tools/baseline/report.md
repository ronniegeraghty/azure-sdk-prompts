# Evaluation Report: ai-projects-dp-js-ts-project-resource-inventory

**Config:** js-ts-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 404.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-projects-dp-js-ts-project-resource-inventory` |
| Config | js-ts-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 16/19 |
| Duration | 404.2s |
| Timestamp | 2026-08-26T13:06:11Z |
| Files Generated | 5 |
| Event Count | 3461 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 218.9s |
| Review | 184.2s |
| **Total** | **404.2s** |

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
| Token Usage | in=224778 out=3358 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269","pattern":"**/package.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269","pattern":"**/tsconfig.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. powershell

**Input:**

```json
{"command":"npm view @azure/ai-projects version --json; npm view @azure/identity version --json; npm view typescript version --json","description":"Check current npm package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
"2.5.0"
"4.13.2"
"7.0.2"
<shellId: 0 completed with exit code 0>
```

### ✅ 5. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/src \"deployments\" \"connections\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"deploymentsBasics.js","path":"sdk/ai/ai-projects/samples/v2/javascript/deployments/deploymentsBasics.js","sha":"57e716fefaee218c0a136187475af2c3a21db1bf","text_matches":[{"fragment":"\n  for await (const deployment of project.deployments.list()) {\n    // Check if this is a ModelDeployment (has the required properties)","matches":[{"indices":[42,58],"text":"deployments.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/samples/v2/javascript/deployments/deploymentsBasics.js?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"  const filteredDeployments = [];\n  for await (const deployment of project.deployments.list({\n    modelPublisher,","matches":[{"indices":[75,91],"text":"deployments.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/samples/v2/javascript/deployments/deploymentsBasics.js?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"deploymentsBasics.ts","path":"sdk/ai/ai-projects/samples/v2/typescript/src/deployments/deploymentsBasics.ts","sha":"b0c6fd0afc92687e05cc0cf0ef36240b13ae824a","text_matches":[{"fragment":"\n  for await (const deployment of project.deployments.list()) {\n    // Check if this is a ModelDeployment (has the required properties)","matches":[{"indices":[42,58],"text":"deployments.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/samples/v2/typescript/src/deployments/deploymentsBasics.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"  const filteredDeployments: ModelDeployment[] = [];\n  for await (const deployment of project.deployments.list({\n    modelPublisher,","matches":[{"indices":[94,110],"text":"deployments.list"}],"object_type":"FileContent","object_url":"https://api.github.c
... (truncated)
```

### ✅ 6. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/src \"ModelDeployment\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"connectionsBasics.js","path":"sdk/ai/ai-projects/samples/v2/javascript/connections/connectionsBasics.js","sha":"9d2b32b5b46463cdddfb2fac7957bd1901bfbfa6","text_matches":[{"fragment":"  const connectionNames = [];\n  for await (const connection of project.connections.list()) {\n    connections.push(connection);","matches":[{"indices":[71,87],"text":"connections.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/samples/v2/javascript/connections/connectionsBasics.js?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"  const azureAIConnections = [];\n  for await (const azureOpenAIConnection of project.connections.list({\n    connectionType: \"AzureOpenAI\",","matches":[{"indices":[85,101],"text":"connections.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/samples/v2/javascript/connections/connectionsBasics.js?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"connectionsBasics.ts","path":"sdk/ai/ai-projects/samples/v2/typescript/src/connections/connectionsBasics.ts","sha":"fb63aed5a03f1504429c0c0fbe6c1d5a3e3e36f7","text_matches":[{"fragment":"  const connectionNames: string[] = [];\n  for await (const connection of project.connections.list()) {\n    connections.push(connection);","matches":[{"indices":[81,97],"text":"connections.list"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/samples/v2/typescript/src/connections/connectionsBasics.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"  const azureAIConnections: Connection[] = [];\n  for await (const azureOpenAIConnection of project.connections.list({\n    connectionType: \"AzureOpenAI\",","matches":[{"indices":[99,115],"text":"connections.list"}],"object_type":"FileC
... (truncated)
```

### ✅ 7. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/samples \"connections.list\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"index.ts","path":"sdk/ai/ai-projects/src/index.ts","sha":"fb724c13f03c38415fbb82f0d34cef70e47b4697","text_matches":[{"fragment":"  Deployment,\n  DeploymentUnion,\n  DeploymentType,\n  ModelDeployment,\n  ModelDeploymentSku,\n  Index,\n  IndexUnion,","matches":[{"indices":[53,68],"text":"ModelDeployment"},{"indices":[72,87],"text":"ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/src/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"index.ts","path":"sdk/ai/ai-projects/src/models/index.ts","sha":"1239536fc6c700e7c2246bcee0d1aa5160dbe56f","text_matches":[{"fragment":"  Deployment,\n  DeploymentUnion,\n  DeploymentType,\n  ModelDeployment,\n  ModelDeploymentSku,\n  Index,\n  IndexUnion,","matches":[{"indices":[53,68],"text":"ModelDeployment"},{"indices":[72,87],"text":"ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/src/models/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]}],"total_count":2}
```

### ✅ 8. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/samples \"deployments.list\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"index.ts","path":"sdk/ai/ai-projects/src/index.ts","sha":"fb724c13f03c38415fbb82f0d34cef70e47b4697","text_matches":[{"fragment":"  ConnectionsGetOptionalParams,\n} from \"./api/connections/index.js\";\nexport type {","matches":[{"indices":[2,13],"text":"Connections"},{"indices":[46,57],"text":"connections"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/src/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"},{"fragment":"  DeploymentsGetOptionalParams,\n} from \"./api/deployments/index.js\";\nexport type {","matches":[{"indices":[2,13],"text":"Deployments"},{"indices":[46,57],"text":"deployments"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/src/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"index.ts","path":"sdk/ai/ai-projects/src/classic/index.ts","sha":"7b67a9aa237ebb5d36fbdf2c3ddea0c1f9afc6a0","text_matches":[{"fragment":"\nexport type { AgentsOperations } from \"./agents/index.js\";\nexport type { BetaOperations } from \"./beta/index.js\";\nexport type { ConnectionsOperations } from \"./connections/index.js\";\nexport type { DatasetsOperations } from \"./datasets/index.js\";\nexport type { DeploymentsOperations } from \"./deployments/index.js\";\nexport type { EvaluationRulesOperations } from \"./evaluationRules/index.js\";","matches":[{"indices":[129,140],"text":"Connections"},{"indices":[161,172],"text":"connections"},{"indices":[261,272],"text":"Deployments"},{"indices":[293,304],"text":"deployments"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/ai/ai-projects/src/classic/index.ts?ref=9ac63ecc12bdfe8975ee911bac4d8dda6d456e33","property":"content"}]},{"name":"aiProjectClient.ts","path":"sdk/ai/ai-projects/src/aiProjectClient.ts","sha":"0871ede79ae4ef45af
... (truncated)
```

### ✅ 9. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-projects/samples/v2/typescript/src/connections/connectionsBasics.ts","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: c856f76a4d43fa3ad96816bf8f6eb99190865950)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

export { list, get } from "./operations.js";
export type { DeploymentsListOptionalParams, DeploymentsGetOptionalParams } from "./options.js";

```

### ✅ 10. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-projects/samples/v2/typescript/src/deployments/deploymentsBasics.ts","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: 3de7945b91f60be33a2001f4b6603f7f518b2b45)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

export { list, getWithCredentials, get } from "./operations.js";
export type {
  ConnectionsListOptionalParams,
  ConnectionsGetWithCredentialsOptionalParams,
  ConnectionsGetOptionalParams,
  ConnectionsGetDefaultOptionalParams,
} from "./options.js";

```

### ✅ 11. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-projects/src/models/models.ts","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: b0c6fd0afc92687e05cc0cf0ef36240b13ae824a)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/**
 * This sample demonstrates how to use the AIProjectClient to manage deployments.
 *
 * @summary Given an AIProjectClient, this sample demonstrates how to enumerate the properties of all deployments,
 * get the properties of a deployment by its name, and delete a deployment.
 */

import type { ModelDeployment } from "@azure/ai-projects";
import { AIProjectClient } from "@azure/ai-projects";
import { DefaultAzureCredential } from "@azure/identity";
import "dotenv/config";

const projectEndpoint = process.env["FOUNDRY_PROJECT_ENDPOINT"] || "<project endpoint string>";
const modelPublisher = process.env["MODEL_PUBLISHER"] || "<model publisher>";

export async function main(): Promise<void> {
  const project = new AIProjectClient(projectEndpoint, new DefaultAzureCredential());

  // List all deployments
  console.log("List all deployments:");
  const deployments: ModelDeployment[] = [];
  const properties: Array<Record<string, string>> = [];

  for await (const deployment of project.deployments.list()) {
    // Check if this is a ModelDeployment (has the required properties)
    if (
      deployment.type === "ModelDeployment" &&
      "modelName" in deployment &&
      "modelPublisher" in deployment &&
      "modelVersion" in deployment
    ) {
      deployments.push(deployment);
      properties.push({
        name: deployment.name,
        modelPublisher: deployment.modelPublisher,
        modelName: deployment.modelName,
      });
    }
  }
  console.log(`Retrieved deployments: ${JSON.stringify(properties, null, 2)}`);

  // List all deployments by a specific model publisher (assuming we have one from the list)
  console.log(`List all deployments by the model publisher '${modelPublisher}':`);
  const filteredDeployments: ModelDeployment[] = [];
  for await (const deployment of project.deployments.list({
    
... (truncated)
```

### ✅ 12. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-projects/src/api/connections/index.ts","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: 278cd9e25fde6c595a5e5df68f8cfefd38f1a7c2)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import type { NodeReadableStream } from "#platform/static-helpers/platform-types";
import type { FileContents } from "../static-helpers/multipartHelpers.js";
import { createFilePartDescriptor } from "../static-helpers/multipartHelpers.js";
import { serializeRecord } from "../static-helpers/serialization/serialize-record.js";

/**
 * This file contains only generated model types and their (de)serializers.
 * Disable the following rules for internal models with '_' prefix and deserializers which require 'any' for raw JSON input.
 */
/* eslint-disable @typescript-eslint/naming-convention */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/* eslint-disable tsdoc/syntax */

/** model interface Agent */
export interface Agent {
  /** The object type, which is always 'agent'. */
  object: "agent";
  /** The unique identifier of the agent. */
  id: string;
  /** The name of the agent. */
  name: string;
  /** The operational state of the agent. Controls whether the agent endpoint accepts or rejects requests. */
  readonly state: AgentState;
  /** The source of the agent's operational state. When the agent is disabled, indicates where the disabled state originates from. Empty when not derived from a specific source. */
  readonly state_source?: AgentStateSource;
  /** The latest version of the agent. */
  versions: {
    latest: AgentVersion;
  };
  /** The endpoint configuration for the agent */
  agent_endpoint?: AgentEndpointConfig;
  /** The instance identity of the agent */
  readonly instance_identity?: AgentIdentity;
  /** The blueprint for the agent */
  readonly blueprint?: AgentIdentity;
  /** The blueprint for the agent */
  readonly blueprint_reference?: AgentBlueprintReferenceUnion;
  agent_card?: AgentCard;
}

export function agentDeserializer(item: any): Agent {
  return {
    object
... (truncated)
```

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/ai-projects/src/api/deployments/index.ts","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: fb63aed5a03f1504429c0c0fbe6c1d5a3e3e36f7)// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

/**
 * This sample demonstrates how to use basic connections operations.
 *
 * @summary Given an AIProjectClient, this sample demonstrates how to enumerate the properties of all connections,
 * get the properties of a default connection, and get the properties of a connection by its name.
 */

import type { Connection } from "@azure/ai-projects";
import { AIProjectClient } from "@azure/ai-projects";
import { DefaultAzureCredential } from "@azure/identity";
import "dotenv/config";

const projectEndpoint = process.env["FOUNDRY_PROJECT_ENDPOINT"] || "<project endpoint string>";

export async function main(): Promise<void> {
  const project = new AIProjectClient(projectEndpoint, new DefaultAzureCredential());

  // List the details of all the connections
  const connections: Connection[] = [];
  const connectionNames: string[] = [];
  for await (const connection of project.connections.list()) {
    connections.push(connection);
    connectionNames.push(connection.name);
  }
  console.log(`Retrieved connections: ${connectionNames}`);

  // Get the details of a connection, without credentials
  const connectionName = connections[0].name;
  const connection = await project.connections.get(connectionName);
  console.log(
    "connection.type: ",
    connection.type,
    "connection.name: ",
    connection.name,
    "connection.target: ",
    connection.target,
  );

  const connectionWithCredentials = await project.connections.getWithCredentials(connectionName);
  const credentials = connectionWithCredentials.credentials;
  console.log("credentials.type: ", credentials.type, "credentials", credentials);

  // List all connections of a specific type
  const azureAIConnections: Connection[] = [];
  for await (const azureOpenAIConnection of project.connections.list({
    connectionType: "AzureOpenAI",
    defaultConnection
... (truncated)
```

### ✅ 14. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/src/models/models.ts \"export interface Connection\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 15. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/src/models/models.ts \"export interface ModelDeployment\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 16. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/src/models/models.ts \"isDefault\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 17. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js path:sdk/ai/ai-projects/src/models/models.ts \"defaultConnection\""}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 18. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"foundry-project-resource-inventory\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"description\": \"Inspect connections and model deployments in a Microsoft Foundry project.\",\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"engines\": {\n+    \"node\": \"\u003e=20\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/ai-projects\": \"^2.5.0\",\n+    \"@azure/identity\": \"^4.13.2\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import {\n+  AIProjectClient,\n+  type Connection,\n+  type DeploymentUnion,\n+  type ModelDeployment,\n+} from \"@azure/ai-projects\";\n+import { DefaultAzureCredential } from \"@azure/identity\";\n+\n+interface Configuration {\n+  projectEndpoint: string;\n+  connectionName: string;\n+  deploymentName: string;\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name]?.trim();\n+  if (!value) {\n+    throw new Error(`Required environment variable ${name} is not set.`);\n+  }\n+  return value;\n+}\n+\n+function loadConfiguration(): Configuration {\n+  return {\n+    projectEndpoint: requireEnvironmentVariable(\"FOUNDRY_PROJECT_ENDPOINT\"),\n+    connectionName: requireEnvironmentVariable(\"CONNECTION_NAME\"),\n+    deploymentName: requireEnvironmentVariable(\"DEPLOYMENT_NAME\"),\n+  };\n+}\n+\n+function printConnection(connection: Connection): void {\n+  console.log({\n+    name: connection.name,\n+    type: connection.type,\n+    target: connection.target,\n+    isDefault: connection.isDefault,\n+  });\n+}\n+\n+function isModelDeployment(deployment: DeploymentUnion): deployment is ModelDeployment {\n+  return deployment.type === \"ModelDeployment\";\n+}\n+\n+function printModelDeployment(deployment: ModelDeployment): void {\n+  console.log({\n+    name: deployment.name,\n+    modelPublisher: deployment.modelPublisher,\n+    modelName: deployment.modelName,\n+    modelVersion: deployment.modelVersion,\n+  });\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const config = loadConfiguration();\n+  const project = new AIProjectClient(\n+    config.projectEndpoint,\n+    new DefaultAzureCredential(),\n+  );\n+\n+  console.log(\"Project connections:\");\n+  for await (const connection of project.connections.list()) {\n+    printConnection(connection);\n+  }\n+\n+  console.log(`Connection \"${config.connectionName}\":`);\n+  const connection = await project.connections.get(config.connectionName);\n+  printConnection(connection);\n+\n+  console.log(\"Project model deployments:\");\n+  for await (const deployment of project.deployments.list()) {\n+    if (isModelDeployment(deployment)) {\n+      printModelDeployment(deployment);\n+    }\n+  }\n+\n+  console.log(`Deployment \"${config.deploymentName}\":`);\n+  const deployment = await project.deployments.get(config.deploymentName);\n+  if (!isModelDeployment(deployment)) {\n+    throw new Error(\n+      `Deployment \"${config.deploymentName}\" has type \"${deployment.type}\", not \"ModelDeployment\".`,\n+    );\n+  }\n+  printModelDeployment(deployment);\n+}\n+\n+await main().catch((error: unknown) =\u003e {\n+  console.error(\"Failed to inspect the Foundry project:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Foundry project resource inventory\n+\n+A TypeScript console application that lists a Microsoft Foundry project's\n+connections and model deployments, then retrieves one named connection and one\n+named model deployment.\n+\n+Authentication uses `DefaultAzureCredential`. Set these environment variables\n+before running:\n+\n+- `FOUNDRY_PROJECT_ENDPOINT`: Foundry project endpoint\n+- `CONNECTION_NAME`: connection to retrieve without credentials\n+- `DEPLOYMENT_NAME`: model deployment to retrieve\n+\n+## Install, build, and run\n+\n+```powershell\n+npm install\n+$env:FOUNDRY_PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:CONNECTION_NAME = \"\u003cconnection-name\u003e\"\n+$env:DEPLOYMENT_NAME = \"\u003cdeployment-name\u003e\"\n+npm run build\n+npm start\n+```\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/package.json
@@ -1,0 +1,23 @@
+{
+  "name": "foundry-project-resource-inventory",
+  "version": "1.0.0",
+  "private": true,
+  "description": "Inspect connections and model deployments in a Microsoft Foundry project.",
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "node dist/index.js"
+  },
+  "engines": {
+    "node": ">=20"
+  },
+  "dependencies": {
+    "@azure/ai-projects": "^2.5.0",
+    "@azure/identity": "^4.13.2"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/tsconfig.json
@@ -1,0 +1,17 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "rootDir": "src",
+    "outDir": "dist",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "exactOptionalPropertyTypes": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users
... (truncated)
```

### ✅ 19. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and build app","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 67 packages in 47s

> foundry-project-resource-inventory@1.0.0 build
> tsc
<shellId: 1 completed with exit code 0>
```

### ✅ 20. powershell

**Input:**

```json
{"command":"npm start; if ($LASTEXITCODE -ne 1) { throw \"Expected missing-configuration exit code 1, got $LASTEXITCODE\" }","description":"Check local configuration failure path","initial_wait":30,"mode":"sync"}
```

**Output:**

```
> foundry-project-resource-inventory@1.0.0 start
> node dist/index.js

Failed to inspect the Foundry project: Error: Required environment variable FOUNDRY_PROJECT_ENDPOINT is not set.
    at requireEnvironmentVariable (file:///C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/dist/index.js:6:15)
    at loadConfiguration (file:///C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/dist/index.js:12:26)
    at main (file:///C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/dist/index.js:37:20)
    at file:///C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-js-ts-project-resource-inventory-js-ts-azure-tools-baseline-3695192269/dist/index.js:59:7
    at ModuleJob.run (node:internal/modules/esm/module_job:439:25)
    at process.processTicksAndRejections (node:internal/process/task_queues:104:5)
    at async node:internal/modules/esm/loader:643:26
    at async asyncRunEntryPointWithESMLoader (node:internal/modules/run_main:101:5)
<shellId: 2 completed with exit code 1>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace structure and existing project conventions, then add the TypeScript app and verify it locally without contacting Azure.

The workspace is empty. I’m checking the installed SDK’s current type surface before writing the project so the pageable iteration and deployment narrowing compile against the declared package version.

The SDK uses `connections.list()` and `deployments.list()` as async pageable iterables, with `connections.get(name)` excluding credentials by default. I’m adding strict environment validation and a type guard so named non-model deployments fail explicitly.

The project compiles against `@azure/ai-projects` 2.5.0. I’m doing a local startup check with no Azure settings to confirm it exits through the intended configuration error before any network or Azure operation.

Created the complete runnable TypeScript application in the workspace.

- `src\index.ts` — pageable connection/deployment inspection with typed model narrowing
- `package.json` and `package-lock.json` — Azure SDK dependencies and build/start scripts
- `tsconfig.json` — strict Node.js ESM TypeScript configuration
- `README.md` — required environment variables and concise install, build, and run commands

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete TypeScript project manifest with the required Azure SDK dependencies, Node.js types, compiler configuration, and runnable scripts.

## Grader Results

- project-resource-inventory.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Creates `AIProjectClient` for `FOUNDRY_PROJECT_ENDPOINT`.: Pass
      - Iterates `client.connections.list()` with `for await...of` and reads typed: Pass
      - Calls `client.connections.get` for `CONNECTION_NAME`, not `getWithCredentials`,: Pass
      - Iterates `client.deployments.list()` with `for await...of`.: Pass
      - Uses the `ModelDeployment` discriminator and properties before printing publisher,: Pass
      - Calls `client.deployments.get` for `DEPLOYMENT_NAME` and rejects a result whose: Pass
      - Does not replace async pageable iteration with a single assumed page.: Pass
      - Does not select resources only from the locally enumerated results instead of: Pass
      - Does not request or display connection credentials.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 8.0000** | **72.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id ai-projects-dp-js-ts-project-resource-inventory --config js-ts-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
