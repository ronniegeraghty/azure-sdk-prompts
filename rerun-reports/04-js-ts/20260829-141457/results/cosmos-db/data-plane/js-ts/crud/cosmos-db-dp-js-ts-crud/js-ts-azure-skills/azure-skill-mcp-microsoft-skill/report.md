# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 302.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 11/17 |
| Duration | 302.3s |
| Timestamp | 2026-08-29T07:32:23Z |
| Files Generated | 4 |
| Event Count | 4223 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 149.8s |
| Review | 150.8s |
| **Total** | **302.3s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=343268 out=3335 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 2130ms

**Input:**

```json
{"skill":"azure-cosmos-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# @azure/cosmos (TypeScript/JavaScript)

Data plane SDK for Azure Cosmos DB NoSQL API operations — CRUD on documents, queries, bulk operations.

> **⚠️ Data vs Management Plane**
> - **This SDK (@azure/cosmos)**: CRUD operations on documents, queries, stored procedures
> - **Management SDK (@azure/arm-cosmosdb)**: Create accounts, databases, containers via ARM

## Installation

```bash
npm install @azure/cosmos @azure/identity
```

**Current Version**: 4.9.0  
**Node.js**: >= 20.0.0

## Environment Variables

```bash
COSMOS_ENDPOINT=https://<account>.documents.azure.com:443/
COSMOS_DATABASE=<database-name>
COSMOS_CONTAINER=<container-name>
# For key-based auth only (prefer AAD)
COSMOS_KEY=<account-key>
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

### Microsoft Entra Token Credential (Recommended)

```typescript
import { CosmosClient } from "@azure/cosmos";
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

const client = new CosmosClient({
  endpoint: process.env.COSMOS_ENDPOINT!,
  aadCredentials: credential,
});
```

### Key-Based Authentication

```typescript
import { CosmosClient } from "@azure/cosmos";

// Option 1: Endpoint + Key
const client = new CosmosClient({
  endpoint: process.env.COSMOS_ENDPOINT!,
  key: process.env.COSMOS_KEY!,
});

// Option 2: Connection String
const client = new CosmosClient(process.env.COSMOS_CONNECTION_STRING!);
```

## Resource Hierarc
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 10587ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Generate a local-only TypeScript program using @azure/cosmos for NoSQL CRUD, parameterized queries, endpoint/key authentication, and status-aware error handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 11916ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-cosmos-ts\\references\\query-patterns.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-cosmos-ts/references/query-patterns.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-cosmos-ts/references/query-patterns.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-cosmos-ts/references/query-patterns.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-cosmos-ts/references/query-patterns.md
@@ -1,397 +1,397 @@
 # Query Patterns Reference
 
 Advanced query patterns for Azure Cosmos DB using the @azure/cosmos TypeScript SDK.
 
 ## Overview
 
 Cosmos DB supports SQL-like queries with support for JSON documents. This reference covers parameterized queries, pagination, cross-partition queries, and advanced query patterns.
 
 ## SqlQuerySpec Interface
 
 ```typescript
 import { SqlQuerySpec, SqlParameter } from "@azure/cosmos";
 
 interface SqlQuerySpec {
   /** SQL query text */
   query: string;
   /** Array of parameters */
   parameters?: SqlParameter[];
 }
 
 interface SqlParameter {
   /** Parameter name (including @) */
   name: string;
   /** Parameter value */
   value: unknown;
 }
 ```
 
 ## Parameterized Queries (Recommended)
 
 Always use parameterized queries to prevent injection and improve plan caching.
 
 ```typescript
 import { SqlQuerySpec, Container } from "@azure/cosmos";
 
 interface Product {
   id: string;
   category: string;
   name: string;
   price: number;
   inStock: boolean;
 }
 
 // Single parameter
 const querySpec: SqlQuerySpec = {
   query: "SELECT * FROM c WHERE c.category = @category",
   parameters: [
     { name: "@category", value: "electronics" }
   ]
 };
 
 const { resources } = await container.items
   .query<Product>(querySpec)
   .fetchAll();
 
 // Multiple parameters
 cons
... (truncated)
```

### ✅ 4. glob — 9421ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-cosmos-ts","pattern":"**\\{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 8254ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753","pattern":"**\\*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 3027ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only TypeScript Azure Cosmos DB NoSQL CRUD sample using endpoint/key authentication and @azure/cosmos.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. apply_patch — 2207ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-crud-typescript\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node --env-file=.env dist/index.js\"\n+  },\n+  \"engines\": {\n+    \"node\": \"\u003e=20.6.0\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.9.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.0.0\",\n+    \"typescript\": \"^5.9.0\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+# Use a local Azure Cosmos DB Emulator endpoint and key.\n+COSMOS_ENDPOINT=https://localhost:8081/\n+COSMOS_KEY=\u003clocal-emulator-key\u003e\n*** Add File: src/index.ts\n+import {\n+  CosmosClient,\n+  ErrorResponse,\n+  type SqlQuerySpec,\n+} from \"@azure/cosmos\";\n+\n+interface InventoryItem {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+function assertStatus(\n+  operation: string,\n+  actualStatus: number,\n+  expectedStatuses: readonly number[],\n+): void {\n+  if (!expectedStatuses.includes(actualStatus)) {\n+    throw new Error(\n+      `${operation} returned HTTP ${actualStatus}; expected ${expectedStatuses.join(\" or \")}`,\n+    );\n+  }\n+}\n+\n+function requireResource\u003cT\u003e(\n+  operation: string,\n+  resource: T | undefined,\n+): T {\n+  if (resource === undefined) {\n+    throw new Error(`${operation} succeeded but returned no resource`);\n+  }\n+  return resource;\n+}\n+\n+function reportError(error: unknown): void {\n+  if (!(error instanceof ErrorResponse)) {\n+    console.error(\n+      error instanceof Error ? error.message : \"An unknown error occurred\",\n+    );\n+    return;\n+  }\n+\n+  switch (error.code) {\n+    case 400:\n+      console.error(`Bad request (400): ${error.message}`);\n+      break;\n+    case 401:\n+    case 403:\n+      console.error(`Authentication or authorization failed (${error.code}): ${error.message}`);\n+      break;\n+    case 404:\n+      console.error(`Resource not found (404): ${error.message}`);\n+      break;\n+    case 409:\n+      console.error(`Item already exists (409): ${error.message}`);\n+      break;\n+    case 412:\n+      console.error(`Precondition failed (412): ${error.message}`);\n+      break;\n+    case 429:\n+      console.error(\n+        `Rate limited (429); retry after ${error.retryAfterInMs ?? \"an unspecified number of\"} ms`,\n+      );\n+      break;\n+    default:\n+      console.error(`Cosmos DB request failed (${error.code}): ${error.message}`);\n+  }\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const client = new CosmosClient({\n+    endpoint: requireEnvironmentVariable(\"COSMOS_ENDPOINT\"),\n+    key: requireEnvironmentVariable(\"COSMOS_KEY\"),\n+  });\n+\n+  try {\n+    const databaseResponse = await client.databases.createIfNotExists({\n+      id: \"TestDB\",\n+    });\n+    assertStatus(\"Create database\", databaseResponse.statusCode, [200, 201]);\n+\n+    const containerResponse =\n+      await databaseResponse.database.containers.createIfNotExists({\n+        id: \"Items\",\n+        partitionKey: { paths: [\"/category\"] },\n+      });\n+    assertStatus(\"Create container\", containerResponse.statusCode, [200, 201]);\n+\n+    const container = containerResponse.container;\n+    const newItem: InventoryItem = {\n+      id: \"item-1\",\n+      category: \"electronics\",\n+      name: \"Wireless Mouse\",\n+      quantity: 10,\n+    };\n+\n+    const createResponse =\n+      await container.items.create\u003cInventoryItem\u003e(newItem);\n+    assertStatus(\"Create item\", createResponse.statusCode, [201]);\n+    console.log(\"Created:\", requireResource(\"Create item\", createResponse.resource));\n+\n+    const itemReference = container.item(newItem.id, newItem.category);\n+    const readResponse = await itemReference.read\u003cInventoryItem\u003e();\n+    assertStatus(\"Read item\", readResponse.statusCode, [200]);\n+    const storedItem = requireResource(\"Read item\", readResponse.resource);\n+    console.log(\"Read:\", storedItem);\n+\n+    const querySpec: SqlQuerySpec = {\n+      query: \"SELECT * FROM c WHERE c.category = @category\",\n+      parameters: [{ name: \"@category\", value: \"electronics\" }],\n+    };\n+    const queryResponse = await container.items\n+      .query\u003cInventoryItem\u003e(querySpec, { partitionKey: \"electronics\" })\n+      .fetchAll();\n+    assertStatus(\"Query items\", queryResponse.statusCode, [200]);\n+    console.log(\"Query results:\", queryResponse.resources);\n+\n+    const updatedItem: InventoryItem = {\n+      ...storedItem,\n+      quantity: 25,\n+    };\n+    const replaceResponse =\n+      await itemReference.replace\u003cInventoryItem\u003e(updatedItem);\n+    assertStatus(\"Replace item\", replaceResponse.statusCode, [200]);\n+    console.log(\n+      \"Replaced:\",\n+      requireResource(\"Replace item\", replaceResponse.resource),\n+    );\n+\n+    const deleteResponse = await itemReference.delete();\n+    assertStatus(\"Delete item\", deleteResponse.statusCode, [204]);\n+    console.log(`Deleted item ${newItem.id}`);\n+  } finally {\n+    client.dispose();\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  reportError(error);\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/package.json
@@ -1,0 +1,21 @@
+{
+  "name": "cosmos-db-crud-typescript",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "node --env-file=.env dist/index.js"
+  },
+  "engines": {
+    "node": ">=20.6.0"
+  },
+  "dependencies": {
+    "@azure/cosmos": "^4.9.0"
+  },
+  "devDependencies": {
+    "@types/node": "^24.0.0",
+    "typescript": "^5.9.0"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/tsconfig.json
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



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/.env.example b/C:/Users/weidxu/AppData/Local/T
... (truncated)
```

### ✅ 8. powershell — 26589ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 25 packages in 21s

> cosmos-db-crud-typescript@1.0.0 build
> tsc

src/index.ts(124,47): error TS2339: Property 'statusCode' does not exist on type 'FeedResponse<InventoryItem>'.
<shellId: 0 completed with exit code 2>
```

### ✅ 9. apply_patch — 1783ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n     const queryResponse = await container.items\n       .query\u003cInventoryItem\u003e(querySpec, { partitionKey: \"electronics\" })\n       .fetchAll();\n-    assertStatus(\"Query items\", queryResponse.statusCode, [200]);\n     console.log(\"Query results:\", queryResponse.resources);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4283398753/src/index.ts
@@ -121,7 +121,6 @@
     const queryResponse = await container.items
       .query<InventoryItem>(querySpec, { partitionKey: "electronics" })
       .fetchAll();
-    assertStatus("Query items", queryResponse.statusCode, [200]);
     console.log("Query results:", queryResponse.resources);
 
     const updatedItem: InventoryItem = {

```

### ✅ 10. powershell — 2413ms

**Input:**

```json
{"command":"npm run build","description":"Compile the TypeScript project","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> cosmos-db-crud-typescript@1.0.0 build
> tsc
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’m grounding the implementation in the Cosmos DB SDK’s current query and error-handling patterns, then I’ll create a locally configurable TypeScript example without contacting Azure.

The workspace is empty, so I’m creating a minimal Node 20 TypeScript project with environment-based emulator credentials, explicit success-status assertions, typed CRUD operations, and Cosmos-specific error reporting.

The SDK’s query feed type does not expose an HTTP status code; query failures are surfaced as `ErrorResponse`. I’m removing that invalid assertion while retaining explicit status checks for every point operation.

Implemented the TypeScript Cosmos DB CRUD example in `src/index.ts`.

- Required package: `npm install @azure/cosmos`
- Creates `TestDB` and `Items` with partition key `/category`
- Performs create, point read, parameterized partition-scoped query, replace, and delete
- Checks HTTP statuses for resource and item operations
- Handles common Cosmos DB errors including `404`, `409`, `412`, and `429`
- Disposes the `CosmosClient` during cleanup
- Reads endpoint/key from environment variables rather than hardcoding credentials

Run against a local Cosmos DB Emulator:

```powershell
Copy-Item .env.example .env
# Set the local emulator key in .env
npm run build
npm start
```

Reference: [Azure Cosmos DB Node.js SDK examples](https://learn.microsoft.com/azure/cosmos-db/nosql/sdk-nodejs)

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Includes the required Azure Cosmos DB SDK dependency and complete TypeScript development dependencies and configuration.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - `@azure/cosmos` npm package: Pass
      - `CosmosClient` constructor with endpoint and key: Pass
      - `client.databases.createIfNotExists()` and `database.containers.createIfNotExists()`: Pass
      - `container.items.create()`, `container.item(id, partitionKey).read()`: Pass
      - `container.items.query()` with `SqlQuerySpec`: Pass
      - `container.item(id, partitionKey).replace()` and `.delete()`: Pass
      - `FeedResponse` iteration and error status codes: Fail
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct @azure/ Scoped Packages**: All Azure SDK dependencies use @azure/ scoped packages (e.g., @azure/storage-blob, @azure/keyvault-secrets, @azure/identity). No deprecated unscoped packages (azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-*).: Pass
  - @azure/identity for Authentication (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**@azure/identity for Authentication**: Includes @azure/identity as a dependency. For local development, DefaultAzureCredential is acceptable. For production code, uses a specific Entra token credential (e.g., ManagedIdentityCredential, ClientSecretCredential, WorkloadIdentityCredential) or configures DefaultAzureCredential with the AZURE_TOKEN_CREDENTIALS environment variable set to "prod" or a specific token credential. No hardcoded connection strings, account keys, SAS tokens, or secrets in code.: Fail
  - Client Constructor with Endpoint and Credential (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Constructor with Endpoint and Credential**: SDK clients constructed with endpoint URL and credential object, e.g. new BlobServiceClient(url, credential) or new SecretClient(vaultUrl, credential). No legacy constructor patterns or connection string-based construction when identity-based auth is available.: Pass
  - Async/Await Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Fail
  - Pagination with for-await-of (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Fail
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 14.6% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 17.1% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 5.8571** | **53.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
