# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 680.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 10/17 |
| Duration | 680.9s |
| Timestamp | 2026-08-27T22:05:46Z |
| Files Generated | 4 |
| Event Count | 4170 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 528.7s |
| Review | 150.9s |
| **Total** | **680.9s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=266279 out=3653 |
| Turn Count | 7 |

## Tool Calls

### ❌ 1. azure-get_azure_bestpractices (via azure) — 185369ms

**Input:**

```json
{"intent":"Generate a local-only TypeScript example using @azure/cosmos for CRUD operations with endpoint/key configuration, parameterized queries, and status-code-aware error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 2. glob — 185366ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 3. azure-documentation (via azure) — 188205ms

**Input:**

```json
{"intent":"Find current official TypeScript/JavaScript Azure Cosmos DB NoSQL SDK examples for CosmosClient endpoint/key authentication, createIfNotExists database and container with /category partition key, item read/replace/delete, parameterized SQL query, and status-code error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. web_search (via github-mcp-server) — 188202ms

**Input:**

```json
{"query":"site:learn.microsoft.com azure cosmos db node.js @azure/cosmos TypeScript item read replace delete parameterized query statusCode"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The latest version of the `@azure/cosmos` package on npm as of August 2026 is **4.10.0**, published about 13 days ago. This is the official Azure Cosmos DB client library for JavaScript and TypeScript, supporting interaction with Cosmos DB's SQL API databases【3:0†source】【3:5†source】【3:6†source】.\n\nTo install the latest version, use:\n```bash\nnpm install @azure/cosmos\n```\n\nYou can view more details or changelogs on the [npm package page](https://www.npmjs.com/package/@azure/cosmos).","annotations":[{"text":"","start_index":259,"end_index":271,"url_citation":{"title":"@azure/cosmos - npm","url":"https://www.npmjs.com/package/@azure/cosmos"}},{"text":"","start_index":271,"end_index":283,"url_citation":{"title":"@azure/cosmos - npm.io","url":"https://npm.io/package/@azure/cosmos"}},{"text":"","start_index":283,"end_index":295,"url_citation":{"title":"Azure Cosmos DB client library for JavaScript | Microsoft Learn","url":"https://learn.microsoft.com/en-us/javascript/api/overview/azure/cosmos-readme?view=azure-node-latest"}}]},"bing_searches":[{"text":"npm @azure/cosmos latest version August 2026","url":"https://www.bing.com/search?q=npm+%40azure%2Fcosmos+latest+version+August+2026"}],"annotations":null}
```

### ✅ 5. web_search (via github-mcp-server) — 186450ms

**Input:**

```json
{"query":"npm @azure/cosmos latest version August 2026 package"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here is a summary of how you can perform item operations (read, replace, delete), parameterized queries, and check status codes in TypeScript/Node.js using the @azure/cosmos SDK, with direct documentation links and code examples:\n\n---\n\n### 1. Reading an Item\n\nYou use the `container.item(id, partitionKey).read<T>()` method to read an item. The response will include an HTTP `statusCode` property.\n\n```typescript\nimport { CosmosClient } from \"@azure/cosmos\";\n\nconst client = new CosmosClient({ endpoint, key });\nconst { resource: item, statusCode } = await container.item(\"id\", partitionKey).read<MyType>();\nconsole.log(\"Status code:\", statusCode);\n```\nReference: [Item.read() Docs]【3:7†source】\n\n---\n\n### 2. Replacing an Item\n\nUse `replace<T>()` to update the item with new content.\n\n```typescript\nconst updatedItem = { ...item, property: \"newValue\" };\nconst { resource: replaced, statusCode } = await container.item(\"id\", partitionKey).replace<MyType>(updatedItem);\nconsole.log(\"Status code:\", statusCode);\n```\nReference: [Item.replace() Example]【3:7†source】\n\n---\n\n### 3. Deleting an Item\n\nDelete using `delete<T>()`. Again, you get the HTTP status code in the response.\n\n```typescript\nconst { statusCode } = await container.item(\"id\", partitionKey).delete<MyType>();\nconsole.log(\"Status code:\", statusCode);\n```\nReference: [Item.delete() Example]【3:7†source】\n\n---\n\n### 4. Parameterized Queries\n\nYou can run parameterized SQL queries with `.items.query(querySpec).fetchAll()`. Here’s an example:\n\n```typescript\nconst querySpec = {\n  query: \"SELECT * FROM c WHERE c.name = @name\",\n  parameters: [{ name: \"@name\", value: \"myName\" }],\n};\nconst { resources, statusCode } = await container.items.query(querySpec).fetchAll();\nconsole.log(\"Status code:\", statusCode);\n```\nReference: [Parameterized Query Example]【3:6†source】\n\n---\n\n### 5. Checking Status Cod
... (truncated)
```

### ✅ 6. apply_patch — 4067ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-typescript-crud\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.10.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"tsx\": \"^4.20.5\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"outDir\": \"dist\",\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import { CosmosClient } from \"@azure/cosmos\";\n+\n+interface InventoryItem {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+const DATABASE_ID = \"TestDB\";\n+const CONTAINER_ID = \"Items\";\n+const PARTITION_KEY_PATH = \"/category\";\n+\n+function requiredEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+function assertStatus(\n+  operation: string,\n+  actual: number,\n+  expected: readonly number[],\n+): void {\n+  if (!expected.includes(actual)) {\n+    throw new Error(\n+      `${operation} returned HTTP ${actual}; expected ${expected.join(\" or \")}`,\n+    );\n+  }\n+}\n+\n+function requireResource\u003cT\u003e(operation: string, resource: T | undefined): T {\n+  if (resource === undefined) {\n+    throw new Error(`${operation} succeeded without returning a resource`);\n+  }\n+  return resource;\n+}\n+\n+function getStatusCode(error: unknown): number | undefined {\n+  if (typeof error !== \"object\" || error === null) {\n+    return undefined;\n+  }\n+\n+  const candidate = error as { code?: unknown; statusCode?: unknown };\n+  if (typeof candidate.code === \"number\") {\n+    return candidate.code;\n+  }\n+  if (typeof candidate.statusCode === \"number\") {\n+    return candidate.statusCode;\n+  }\n+  return undefined;\n+}\n+\n+function getErrorMessage(error: unknown): string {\n+  return error instanceof Error ? error.message : String(error);\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const endpoint = requiredEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+  const key = requiredEnvironmentVariable(\"COSMOS_KEY\");\n+  const client = new CosmosClient({ endpoint, key });\n+\n+  const databaseResponse = await client.databases.createIfNotExists({\n+    id: DATABASE_ID,\n+  });\n+  assertStatus(\"Create database\", databaseResponse.statusCode, [200, 201]);\n+  const database = databaseResponse.database;\n+\n+  const containerResponse = await database.containers.createIfNotExists({\n+    id: CONTAINER_ID,\n+    partitionKey: {\n+      paths: [PARTITION_KEY_PATH],\n+    },\n+  });\n+  assertStatus(\"Create container\", containerResponse.statusCode, [200, 201]);\n+  const container = containerResponse.container;\n+\n+  const item: InventoryItem = {\n+    id: \"item-001\",\n+    category: \"electronics\",\n+    name: \"Wireless headphones\",\n+    quantity: 5,\n+  };\n+\n+  const createResponse = await container.items.create\u003cInventoryItem\u003e(item);\n+  assertStatus(\"Create item\", createResponse.statusCode, [201]);\n+  const createdItem = requireResource(\"Create item\", createResponse.resource);\n+  console.log(\"Created:\", createdItem);\n+\n+  const readResponse = await container\n+    .item(item.id, item.category)\n+    .read\u003cInventoryItem\u003e();\n+  assertStatus(\"Read item\", readResponse.statusCode, [200]);\n+  const readItem = requireResource(\"Read item\", readResponse.resource);\n+  console.log(\"Read:\", readItem);\n+\n+  const querySpec = {\n+    query: \"SELECT * FROM items i WHERE i.category = @category\",\n+    parameters: [{ name: \"@category\", value: \"electronics\" }],\n+  };\n+  const queryResponse = await container.items\n+    .query\u003cInventoryItem\u003e(querySpec, { partitionKey: \"electronics\" })\n+    .fetchAll();\n+  assertStatus(\"Query items\", queryResponse.statusCode, [200]);\n+  console.log(\"Query results:\", queryResponse.resources);\n+\n+  const updatedItem: InventoryItem = {\n+    ...readItem,\n+    quantity: 10,\n+  };\n+  const replaceResponse = await container\n+    .item(updatedItem.id, updatedItem.category)\n+    .replace\u003cInventoryItem\u003e(updatedItem);\n+  assertStatus(\"Replace item\", replaceResponse.statusCode, [200]);\n+  console.log(\n+    \"Replaced:\",\n+    requireResource(\"Replace item\", replaceResponse.resource),\n+  );\n+\n+  const deleteResponse = await container\n+    .item(updatedItem.id, updatedItem.category)\n+    .delete\u003cInventoryItem\u003e();\n+  assertStatus(\"Delete item\", deleteResponse.statusCode, [204]);\n+  console.log(`Deleted item ${updatedItem.id}`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  const statusCode = getStatusCode(error);\n+  const message = getErrorMessage(error);\n+\n+  switch (statusCode) {\n+    case 400:\n+      console.error(`Bad Cosmos DB request (HTTP 400): ${message}`);\n+      break;\n+    case 401:\n+    case 403:\n+      console.error(\n+        `Cosmos DB authentication/authorization failed (HTTP ${statusCode}): ${message}`,\n+      );\n+      break;\n+    case 404:\n+      console.error(`Cosmos DB resource not found (HTTP 404): ${message}`);\n+      break;\n+    case 409:\n+      console.error(`Cosmos DB resource conflict (HTTP 409): ${message}`);\n+      break;\n+    case 429:\n+      console.error(`Cosmos DB rate limit exceeded (HTTP 429): ${message}`);\n+      break;\n+    default:\n+      console.error(\n+        statusCode === undefined\n+          ? `Unexpected error: ${message}`\n+          : `Cosmos DB request failed (HTTP ${statusCode}): ${message}`,\n+      );\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "cosmos-db-typescript-crud",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "tsx src/index.ts"
+  },
+  "dependencies": {
+    "@azure/cosmos": "^4.10.0"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "tsx": "^4.20.5",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/tsconfig.json
@@ -1,0 +1,14 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "strict": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "outDir": "dist",
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/src/index.ts
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp
... (truncated)
```

### ✅ 7. powershell — 69747ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 28 packages in 59s

> cosmos-db-typescript-crud@1.0.0 build
> tsc

src/index.ts(106,45): error TS2339: Property 'statusCode' does not exist on type 'FeedResponse<InventoryItem>'.
<shellId: 0 completed with exit code 2>
```

### ✅ 8. apply_patch — 3615ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n   const queryResponse = await container.items\n     .query\u003cInventoryItem\u003e(querySpec, { partitionKey: \"electronics\" })\n     .fetchAll();\n-  assertStatus(\"Query items\", queryResponse.statusCode, [200]);\n   console.log(\"Query results:\", queryResponse.resources);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-2760355505/src/index.ts
@@ -103,7 +103,6 @@
   const queryResponse = await container.items
     .query<InventoryItem>(querySpec, { partitionKey: "electronics" })
     .fetchAll();
-  assertStatus("Query items", queryResponse.statusCode, [200]);
   console.log("Query results:", queryResponse.resources);
 
   const updatedItem: InventoryItem = {

```

### ✅ 9. powershell — 4841ms

**Input:**

```json
{"command":"npm run build","description":"Compile the TypeScript sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> cosmos-db-typescript-crud@1.0.0 build
> tsc
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace and ground the implementation in current Azure SDK guidance before adding the TypeScript program.

The guidance endpoint timed out, so I’m switching to official Cosmos DB SDK documentation and package metadata, then I’ll implement the local sample.

The workspace is empty. I’m adding a runnable TypeScript sample with explicit success-status assertions, categorized Cosmos error handling, and environment-based credentials.

The SDK’s feed response intentionally has no `statusCode`; query failures are thrown instead. I’m aligning the query path with that API while retaining explicit HTTP checks for every point operation.

Implemented the TypeScript Cosmos DB CRUD sample in `src/index.ts`.

- Required package: `npm install @azure/cosmos`
- Configure credentials:
  ```powershell
  $env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
  $env:COSMOS_KEY = "<cosmos-db-key>"
  npm start
  ```
- Creates `TestDB` and the `Items` container with `/category`.
- Performs create, point read, parameterized query, replace, and delete.
- Checks expected HTTP status codes and handles common Cosmos DB errors including `400`, `401`, `403`, `404`, `409`, and `429`.

Project configuration is included in `package.json`, `package-lock.json`, and `tsconfig.json`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Complete TypeScript project configuration with the correct Cosmos DB SDK dependency and Node.js type definitions.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/7)
      - `@azure/cosmos` npm package: Pass
      - `CosmosClient` constructor with endpoint and key: Pass
      - `client.databases.createIfNotExists()` and `database.containers.createIfNotExists()`: Pass
      - `container.items.create()`, `container.item(id, partitionKey).read()`: Pass
      - `container.items.query()` with `SqlQuerySpec`: Fail
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
| `Criteria from prompt file` | prompt_review | 71% | 1.00 | 0.7143 | 12.5% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 5.7143** | **51.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
