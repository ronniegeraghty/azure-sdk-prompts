# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 341.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 11/17 |
| Duration | 341.4s |
| Timestamp | 2026-08-29T07:50:59Z |
| Files Generated | 5 |
| Event Count | 4405 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 186.9s |
| Review | 153.3s |
| **Total** | **341.4s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=456557 out=4987 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. glob — 5907ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 7091ms

**Input:**

```json
{"command":"get","intent":"Generate a local-only TypeScript program using the Azure Cosmos DB JavaScript SDK for NoSQL CRUD operations, parameterized queries, and status-code-aware error handling.","parameters":{"language":"TypeScript","operation":"code-generation","resource":"Azure Cosmos DB"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 8458ms

**Input:**

```json
{"command":"search","intent":"Find current official Azure Cosmos DB JavaScript/TypeScript SDK guidance for CosmosClient, databases.containers createIfNotExists, item read replace delete, parameterized SQL queries, and error status codes.","parameters":{"query":"Azure Cosmos DB JavaScript SDK TypeScript CRUD item read replace delete parameterized query statusCode createIfNotExists"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 6671ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a local TypeScript Azure Cosmos DB CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 10862ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Azure Cosmos DB JavaScript SDK CRUD and error handling guidance.","parameters":{"query":"Azure Cosmos DB JavaScript SDK TypeScript CosmosClient createIfNotExists item read replace delete query parameters statusCode"}}
```

**Output:**

```
{"results":[{"title":"Create an item in Azure Cosmos DB for NoSQL using JavaScript","content":"# Create an item in Azure Cosmos DB for NoSQL using JavaScript\n## Access an item\nAccess an item through the [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) object. This can accessed from the [Container](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/container) object or changed from either the [Database](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/database) or [CosmosClient](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/cosmosclient) objects.\n\u0060\u0060\u0060javascript\n// Chained, then use a method of the Item object such as \u0060read\u0060\nconst { statusCode, item, resource, activityId, etag} = await client.database(databaseId).container(containerId).item(itemId).read();\n\u0060\u0060\u0060\nAccess by object:\n1. [Items](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/items) (plural): Create, batch, watch change feed, read all, upsert, or query items.\n2. [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) (singular): Read, patch, replace, or delete an item.\n## Replace an item\nReplace the data with the [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) object with the [replace](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item#@azure-cosmos-item-replace) method.\n\u0060\u0060\u0060javascript\nconst { statusCode, item, resource, activityId, etag} = await item.replace({ \n        id: \u00272\u0027, \n        category: \u0027gear-surf-surfboards-retro\u0027,\n        name: \u0027Sunnox Surfboard Retro\u0027,\n        quantity: 5,\n        sale: false \n    });\n\u0060\u0060\u0060\n## Read an item\nRead the most current data with the [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) object\u0027s [read](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item#@azure-cosm
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 8398ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official TypeScript examples for Azure Cosmos DB CRUD operations.","parameters":{"language":"typescript","query":"@azure/cosmos TypeScript CosmosClient database container item read replace delete parameterized query"}}
```

**Output:**

```
{"results":[{"description":"description: Defines server-side functions in a TypeScript file to interact with the Azure Cosmos DB emulator, including creating a client, database, container, and CRUD operations for a todo list.\npackage: @azure/cosmos\nlanguage: typescript\n","codeSnippet":"\u0022use server\u0022;\r\n    \r\n    import { CosmosClient } from \u0022@azure/cosmos\u0022;\r\n    import { revalidatePath } from \u0022next/cache\u0022;\r\n    \r\n    const connectionString = \u0022\u003Ccosmos-db-connection-string\u003E\u0022;\r\n    \r\n    const client = new CosmosClient(connectionString);\r\n    \r\n    const { database } = await client.databases.createIfNotExists({ id: \u0022work-management\u0022 });\r\n    const { container } = await database.containers.createIfNotExists({ id: \u0022tasks\u0022, partitionKey: \u0022/id\u0022 });\r\n    \r\n    export type Todo = {\r\n      id: string;\r\n      title: string;\r\n      completed: boolean;\r\n    };\r\n    \r\n    export async function getTodos(): Promise\u003CTodo[]\u003E {\r\n      const { resources } = await container.items\r\n        .query(\u0022SELECT * FROM c ORDER BY c._ts DESC\u0022)\r\n        .fetchAll();\r\n      return resources.map((doc) =\u003E ({\r\n        id: doc.id,\r\n        title: doc.title,\r\n        completed: doc.completed,\r\n      }));\r\n    }\r\n    \r\n    export async function addTodo(formData: FormData) {\r\n      const title = formData.get(\u0022title\u0022) as string;\r\n      if (!title?.trim()) return;\r\n      await container.items.create({\r\n        title: title.trim(),\r\n        completed: false,\r\n      });\r\n      revalidatePath(\u0022/\u0022);\r\n    }\r\n    \r\n    export async function toggleTodo(formData: FormData) {\r\n      const id = formData.get(\u0022id\u0022) as string;\r\n      const { resource: doc } = await container.item(id, id).read();\r\n      if (doc) {\r\n        await container.item(id, id).replace({\r\n          ...doc,\r\n          complete
... (truncated)
```

### ✅ 7. rg — 6662ms

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt"],"pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

**Output:**

```
Output too large to read at once (66.9 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787989919860-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt:{"results":[{"title":"Create an item in Azure Cosmos DB for NoSQL using JavaScript","content":"# Create an item in Azure Cosmos DB for NoSQL using JavaScript\n## Access an item\nAccess an item through the [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) object. This can accessed from the [Container](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/con
```

### ✅ 8. rg — 6337ms

**Input:**

```json
{"head_limit":50,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt"],"pattern":"statusCode|CosmosClient|createIfNotExists|item\\([^)]*\\)\\.read|replace\\(|delete\\(|parameters"}
```

**Output:**

```
Output too large to read at once (66.9 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787989919859-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt:{"results":[{"description":"description: Defines server-side functions in a TypeScript file to interact with the Azure Cosmos DB emulator, including creating a client, database, container, and CRUD operations for a todo list.\npackage: @azure/cosmos\nlanguage: typescript\n","codeSnippet":"\u0022use server\u0022;\r\n    \r\n    import { CosmosClient } from \u0022@azure/cosmos\u0022;\r\n    im
```

### ✅ 9. powershell — 13200ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt' | ConvertFrom-Json; $d.results | Select-Object -First 8 title,contentUrl | ConvertTo-Json -Depth 3","description":"Extract Cosmos documentation references"}
```

**Output:**

```
4.10.0
7.0.2
4.23.12
26.2.0
<shellId: 2 completed with exit code 0>
```

### ✅ 10. powershell — 12858ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt' | ConvertFrom-Json; $d.results | Select-Object -First 8 description,url | ConvertTo-Json -Depth 3","description":"Extract Cosmos code sample references"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-69d5a4cbd26647b4857b51e6f1e2f065.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 11. powershell — 12520ms

**Input:**

```json
{"command":"npm view @azure/cosmos version; npm view typescript version; npm view tsx version; npm view @types/node version","description":"Check current npm package versions","initial_wait":60}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-cce655d348854670b174d74861db2b58.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 12. view — 1331ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt b/C:/Users/weidxu/AppData/Local/Temp/1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt
@@ -1,6 +1,6 @@
 {"results":[{"description":"description: Defines server-side functions in a TypeScript file to interact with the Azure Cosmos DB emulator, including creating a client, database, container, and CRUD operations for a todo list.\npackage: @azure/cosmos\nlanguage: typescript\n","codeSnippet":"\u0022use server\u0022;\r\n    \r\n    import { CosmosClient } from \u0022@azure/cosmos\u0022;\r\n    import { revalidatePath } from \u0022next/cache\u0022;\r\n    \r\n    const connectionString = \u0022\u003Ccosmos-db-connection-string\u003E\u0022;\r\n    \r\n    const client = new CosmosClient(connectionString);\r\n    \r\n    const { database } = await client.databases.createIfNotExists({ id: \u0022work-management\u0022 });\r\n    const { container } = await database.containers.createIfNotExists({ id: \u0022tasks\u0022, partitionKey: \u0022/id\u0022 });\r\n    \r\n    export type Todo = {\r\n      id: string;\r\n      title: string;\r\n      completed: boolean;\r\n    };\r\n    \r\n    export async function getTodos(): Promise\u003CTodo[]\u003E {\r\n      const { resources } = await container.items\r\n        .query(\u0022SELECT * FROM c ORDER BY c._ts DESC\u0022)\r\n        .fetchAll();\r\n      return resources.map((doc) =\u003E ({\r\n        id: doc.id,\r\n        title: doc.title,\r\n        completed: doc.completed,\r\n      }));\r\n    }\r\n    \r\n    export async function addTodo(formData: FormData) {\r\n      const title = formData.get(\u0022title\u0022) as string;\r\n      if (!t
... (truncated)
```

### ✅ 13. view — 990ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787989902365-copilot-tool-output-cce655d348854670b174d74861db2b58.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt b/C:/Users/weidxu/AppData/Local/Temp/1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787989902905-copilot-tool-output-69d5a4cbd26647b4857b51e6f1e2f065.txt
@@ -1,6 +1,6 @@
 {"results":[{"title":"Create an item in Azure Cosmos DB for NoSQL using JavaScript","content":"# Create an item in Azure Cosmos DB for NoSQL using JavaScript\n## Access an item\nAccess an item through the [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) object. This can accessed from the [Container](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/container) object or changed from either the [Database](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/database) or [CosmosClient](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/cosmosclient) objects.\n\u0060\u0060\u0060javascript\n// Chained, then use a method of the Item object such as \u0060read\u0060\nconst { statusCode, item, resource, activityId, etag} = await client.database(databaseId).container(containerId).item(itemId).read();\n\u0060\u0060\u0060\nAccess by object:\n1. [Items](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/items) (plural): Create, batch, watch change feed, read all, upsert, or query items.\n2. [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) (singular): Read, patch, replace, or delete an item.\n## Replace an item\nReplace the data with the [Item](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item) object with the [replace](https://learn.microsoft.com/en-us/javascript/api/@azure/cosmos/item#@azure-cosmos-item-replace) method.\n\u0060\u0060\u0060javascript\nconst { statusCode, i
... (truncated)
```

### ✅ 14. apply_patch — 695ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-typescript-crud\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"start\": \"tsx src/index.ts\",\n+    \"typecheck\": \"tsc --noEmit\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.10.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"tsx\": \"^4.23.12\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"skipLibCheck\": true,\n+    \"types\": [\"node\"]\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .gitignore\n+node_modules/\n+\n*** Add File: src/index.ts\n+import { CosmosClient, type SqlQuerySpec } from \"@azure/cosmos\";\n+\n+interface InventoryItem {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+interface CosmosServiceError {\n+  code: number;\n+  message: string;\n+  activityId?: string;\n+}\n+\n+const databaseId = \"TestDB\";\n+const containerId = \"Items\";\n+const partitionKeyPath = \"/category\";\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+function expectStatus(\n+  operation: string,\n+  statusCode: number | undefined,\n+  expectedStatusCodes: readonly number[],\n+): void {\n+  if (statusCode === undefined || !expectedStatusCodes.includes(statusCode)) {\n+    throw new Error(\n+      `${operation} returned HTTP ${statusCode ?? \"unknown\"}; expected ${expectedStatusCodes.join(\" or \")}`,\n+    );\n+  }\n+}\n+\n+function isCosmosServiceError(error: unknown): error is CosmosServiceError {\n+  if (typeof error !== \"object\" || error === null) {\n+    return false;\n+  }\n+\n+  const candidate = error as Record\u003cstring, unknown\u003e;\n+  return (\n+    typeof candidate.code === \"number\" \u0026\u0026 typeof candidate.message === \"string\"\n+  );\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+  const key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+  const client = new CosmosClient({ endpoint, key });\n+\n+  const databaseResponse = await client.databases.createIfNotExists({\n+    id: databaseId,\n+  });\n+  expectStatus(\"Create database\", databaseResponse.statusCode, [200, 201]);\n+  const { database } = databaseResponse;\n+\n+  const containerResponse = await database.containers.createIfNotExists({\n+    id: containerId,\n+    partitionKey: partitionKeyPath,\n+  });\n+  expectStatus(\"Create container\", containerResponse.statusCode, [200, 201]);\n+  const { container } = containerResponse;\n+\n+  const newItem: InventoryItem = {\n+    id: \"item-001\",\n+    category: \"electronics\",\n+    name: \"Wireless headphones\",\n+    quantity: 10,\n+  };\n+\n+  const createResponse = await container.items.create\u003cInventoryItem\u003e(newItem);\n+  expectStatus(\"Create item\", createResponse.statusCode, [201]);\n+  console.log(\"Created:\", createResponse.resource);\n+\n+  const item = container.item(newItem.id, newItem.category);\n+  const readResponse = await item.read\u003cInventoryItem\u003e();\n+  expectStatus(\"Read item\", readResponse.statusCode, [200]);\n+  if (!readResponse.resource) {\n+    throw new Error(\"Read item succeeded but returned no resource\");\n+  }\n+  console.log(\"Read:\", readResponse.resource);\n+\n+  const query: SqlQuerySpec = {\n+    query: \"SELECT * FROM items i WHERE i.category = @category\",\n+    parameters: [{ name: \"@category\", value: \"electronics\" }],\n+  };\n+  const queryResponse = await container.items\n+    .query\u003cInventoryItem\u003e(query)\n+    .fetchAll();\n+  expectStatus(\"Query items\", queryResponse.statusCode, [200]);\n+  console.log(\"Query results:\", queryResponse.resources);\n+\n+  const updatedItem: InventoryItem = {\n+    ...readResponse.resource,\n+    quantity: 25,\n+  };\n+  const replaceResponse = await item.replace\u003cInventoryItem\u003e(updatedItem);\n+  expectStatus(\"Replace item\", replaceResponse.statusCode, [200]);\n+  console.log(\"Replaced:\", replaceResponse.resource);\n+\n+  const deleteResponse = await item.delete();\n+  expectStatus(\"Delete item\", deleteResponse.statusCode, [204]);\n+  console.log(`Deleted item \"${newItem.id}\"`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  if (isCosmosServiceError(error)) {\n+    const activity = error.activityId\n+      ? `, activity ID ${error.activityId}`\n+      : \"\";\n+    console.error(\n+      `Azure Cosmos DB request failed with HTTP ${error.code}${activity}: ${error.message}`,\n+    );\n+  } else if (error instanceof Error) {\n+    console.error(error.message);\n+  } else {\n+    console.error(\"An unknown error occurred:\", error);\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Cosmos DB TypeScript CRUD sample\n+\n+This sample uses `@azure/cosmos` to create `TestDB` and an `Items` container,\n+then creates, reads, queries, replaces, and deletes an item.\n+\n+Install dependencies:\n+\n+```powershell\n+npm install\n+```\n+\n+Set the Azure Cosmos DB for NoSQL endpoint and key without storing credentials\n+in source:\n+\n+```powershell\n+$env:COSMOS_ENDPOINT = \"https://\u003caccount-name\u003e.documents.azure.com:443/\"\n+$env:COSMOS_KEY = \"\u003caccount-key\u003e\"\n+npm start\n+```\n+\n+The required runtime package is:\n+\n+```powershell\n+npm install @azure/cosmos\n+```\n+\n+For production Azure-hosted applications, prefer Microsoft Entra ID and managed\n+identity over account keys.\n+\n+References:\n+\n+- [Azure Cosmos DB JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme?view=azure-node-latest)\n+- [Create and access items](https://learn.microsoft.com/azure/cosmos-db/how-to-javascript-create-item)\n+- [Official local development sample](https://learn.microsoft.com/azure/cosmos-db/development-loop#create-the-application)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "cosmos-db-typescript-crud",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "start": "tsx src/index.ts",
+    "typecheck": "tsc --noEmit"
+  },
+  "dependencies": {
+    "@azure/cosmos": "^4.10.0"
+  },
+  "devDependencies": {
+    "@types/node": "^26.2.0",
+    "tsx": "^4.23.12",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/tsconfig.json
@@ -1,0 +1,14 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "exactOptionalPropertyTypes": true,
+    "skipLibCheck": true,
+    "types": ["node"]
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/.gitignore b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/.gitignore
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos
... (truncated)
```

### ✅ 15. powershell — 27182ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run typecheck","description":"Install dependencies and type-check sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 29 packages in 25s

> cosmos-db-typescript-crud@1.0.0 typecheck
> tsc --noEmit

src/index.ts(95,45): error TS2339: Property 'statusCode' does not exist on type 'FeedResponse<InventoryItem>'.
<shellId: 3 completed with exit code 1>
```

### ✅ 16. apply_patch — 701ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n   const queryResponse = await container.items\n     .query\u003cInventoryItem\u003e(query)\n     .fetchAll();\n-  expectStatus(\"Query items\", queryResponse.statusCode, [200]);\n   console.log(\"Query results:\", queryResponse.resources);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-659219683/src/index.ts
@@ -92,7 +92,6 @@
   const queryResponse = await container.items
     .query<InventoryItem>(query)
     .fetchAll();
-  expectStatus("Query items", queryResponse.statusCode, [200]);
   console.log("Query results:", queryResponse.resources);
 
   const updatedItem: InventoryItem = {

```

### ✅ 17. powershell — 25337ms

**Input:**

```json
{"command":"npm run typecheck","description":"Type-check Cosmos CRUD sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> cosmos-db-typescript-crud@1.0.0 start
> tsx src/index.ts

Missing required environment variable: COSMOS_ENDPOINT
<shellId: 5 completed with exit code 1>
```

### ✅ 18. powershell — 24871ms

**Input:**

```json
{"command":"Remove-Item Env:COSMOS_ENDPOINT -ErrorAction SilentlyContinue; Remove-Item Env:COSMOS_KEY -ErrorAction SilentlyContinue; npm start; if ($LASTEXITCODE -ne 1) { throw \"Expected configuration failure exit code 1, got $LASTEXITCODE\" }","description":"Check missing credential handling","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> cosmos-db-typescript-crud@1.0.0 typecheck
> tsc --noEmit
<shellId: 4 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace and ground the implementation in current Azure SDK guidance before adding the local TypeScript sample.

The workspace is empty. I’ve confirmed the SDK patterns; next I’m extracting the exact API/error details and will create a minimal runnable project.

I’m implementing the sample with environment-based credentials, explicit expected HTTP statuses for every operation, and typed handling for thrown Cosmos service errors.

The query feed type does not expose an HTTP status in SDK v4, so query failures are handled through the SDK’s thrown service error code; point operations retain explicit response-status validation.

Created a complete runnable TypeScript project. `src/index.ts` implements database/container creation, create/read/parameterized-query/replace/delete operations, explicit HTTP status checks, environment-based credentials, and Cosmos service error handling.

```powershell
npm install
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
npm start
```

Required package: `@azure/cosmos` v4.10.0. References: [Cosmos DB JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme?view=azure-node-latest) and [item CRUD guidance](https://learn.microsoft.com/azure/cosmos-db/how-to-javascript-create-item).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Provides complete TypeScript project metadata, the required Azure Cosmos DB SDK dependency, Node type definitions, and an appropriate TypeScript configuration.

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
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
