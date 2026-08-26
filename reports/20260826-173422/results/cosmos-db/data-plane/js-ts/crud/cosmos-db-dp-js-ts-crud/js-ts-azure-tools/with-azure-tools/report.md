# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 317.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 12/17 |
| Duration | 317.7s |
| Timestamp | 2026-08-26T11:19:30Z |
| Files Generated | 5 |
| Event Count | 5237 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 185.9s |
| Review | 130.6s |
| **Total** | **317.7s** |

## Configuration

- **name:** js-ts-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=469987 out=3976 |
| Turn Count | 11 |

## Prompt Sent

```
Write a TypeScript program
that performs CRUD operations on items in an Azure Cosmos DB NoSQL container:
1. Create a CosmosClient using endpoint and key
2. Create a database "TestDB" and container "Items" with partition key "/category"
3. Create an item with properties: id, category, name, quantity
4. Read the item back using item().read()
5. Query items where category equals "electronics" using parameterized query
6. Replace the item with updated quantity using item().replace()
7. Delete the item using item().delete()

Show required npm package and handle errors with appropriate status code checks.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-cosmos-ts"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-tools-with-azure-tools-4174975330","pattern":"**/*"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-cosmos-ts","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-cosmos-ts\\references\\query-patterns.md"}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local-only TypeScript example using @azure/cosmos with endpoint/key authentication, database and container create-if-not-exists, item CRUD, a parameterized category query, status-code-aware errors, and client cleanup. Do not deploy or modify Azure resources during development.","learn":true}
```

### 🔧 6. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript","pattern":"**/acceptance-criteria.md"}
```

### 🔧 7. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript","pattern":"**/examples.md"}
```

### 🔧 8. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a local TypeScript @azure/cosmos CRUD sample; no live Azure provisioning or execution.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"npm view @azure/cosmos version --silent","description":"Check latest Cosmos SDK version","initial_wait":30,"mode":"sync"}
```

### 🔧 10. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-crud-typescript\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"engines\": {\n+    \"node\": \"\u003e=20.0.0\"\n+  },\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.10.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+.env\n+*.log\n*** Add File: src/index.ts\n+import {\n+  CosmosClient,\n+  ErrorResponse,\n+  type SqlQuerySpec,\n+} from \"@azure/cosmos\";\n+import { randomUUID } from \"node:crypto\";\n+\n+const DATABASE_ID = \"TestDB\";\n+const CONTAINER_ID = \"Items\";\n+const PARTITION_KEY_PATH = \"/category\";\n+\n+interface InventoryItem {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+\n+  return value;\n+}\n+\n+function expectStatus(\n+  operation: string,\n+  actualStatus: number,\n+  expectedStatuses: readonly number[],\n+): void {\n+  if (!expectedStatuses.includes(actualStatus)) {\n+    throw new Error(\n+      `${operation} returned HTTP ${actualStatus}; expected ${expectedStatuses.join(\" or \")}`,\n+    );\n+  }\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+  const key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+  const client = new CosmosClient({ endpoint, key });\n+\n+  try {\n+    const databaseResponse = await client.databases.createIfNotExists({\n+      id: DATABASE_ID,\n+    });\n+    expectStatus(\n+      \"Create database\",\n+      databaseResponse.statusCode,\n+      [200, 201],\n+    );\n+    console.log(`Database ready: ${databaseResponse.database.id}`);\n+\n+    const containerResponse =\n+      await databaseResponse.database.containers.createIfNotExists({\n+        id: CONTAINER_ID,\n+        partitionKey: { paths: [PARTITION_KEY_PATH] },\n+      });\n+    expectStatus(\n+      \"Create container\",\n+      containerResponse.statusCode,\n+      [200, 201],\n+    );\n+    const { container } = containerResponse;\n+    console.log(`Container ready: ${container.id}`);\n+\n+    const item: InventoryItem = {\n+      id: randomUUID(),\n+      category: \"electronics\",\n+      name: \"Wireless keyboard\",\n+      quantity: 10,\n+    };\n+\n+    const createResponse = await container.items.create\u003cInventoryItem\u003e(item);\n+    expectStatus(\"Create item\", createResponse.statusCode, [201]);\n+    if (!createResponse.resource) {\n+      throw new Error(\"Create item succeeded without returning the item\");\n+    }\n+    console.log(\"Created:\", createResponse.resource);\n+\n+    const itemReference = container.item(item.id, item.category);\n+    const readResponse = await itemReference.read\u003cInventoryItem\u003e();\n+    expectStatus(\"Read item\", readResponse.statusCode, [200]);\n+    if (!readResponse.resource) {\n+      throw new Error(\"Read item succeeded without returning the item\");\n+    }\n+    console.log(\"Read:\", readResponse.resource);\n+\n+    const querySpec: SqlQuerySpec = {\n+      query: \"SELECT * FROM c WHERE c.category = @category\",\n+      parameters: [{ name: \"@category\", value: \"electronics\" }],\n+    };\n+    const queryResponse = await container.items\n+      .query\u003cInventoryItem\u003e(querySpec, { partitionKey: \"electronics\" })\n+      .fetchAll();\n+    expectStatus(\"Query items\", queryResponse.statusCode, [200]);\n+    console.log(\"Query results:\", queryResponse.resources);\n+\n+    const replacement: InventoryItem = {\n+      ...readResponse.resource,\n+      quantity: 25,\n+    };\n+    const replaceResponse =\n+      await itemReference.replace\u003cInventoryItem\u003e(replacement);\n+    expectStatus(\"Replace item\", replaceResponse.statusCode, [200]);\n+    if (!replaceResponse.resource) {\n+      throw new Error(\"Replace item succeeded without returning the item\");\n+    }\n+    console.log(\"Replaced:\", replaceResponse.resource);\n+\n+    const deleteResponse = await itemReference.delete();\n+    expectStatus(\"Delete item\", deleteResponse.statusCode, [204]);\n+    console.log(`Deleted item ${item.id}`);\n+  } finally {\n+    client.dispose();\n+  }\n+}\n+\n+function handleError(error: unknown): void {\n+  if (error instanceof ErrorResponse) {\n+    const statusCode = error.code;\n+\n+    switch (statusCode) {\n+      case 400:\n+        console.error(\"Cosmos DB rejected the request as invalid:\", error.message);\n+        break;\n+      case 401:\n+      case 403:\n+        console.error(\n+          `Cosmos DB authorization failed (HTTP ${statusCode}). Check the endpoint and key.`,\n+        );\n+        break;\n+      case 404:\n+        console.error(\"The requested Cosmos DB resource was not found.\");\n+        break;\n+      case 409:\n+        console.error(\"The item already exists.\");\n+        break;\n+      case 412:\n+        console.error(\"The item changed before it could be replaced.\");\n+        break;\n+      case 429:\n+        console.error(\n+          `Cosmos DB rate limit exceeded. Retry after ${error.retryAfterInMs ?? \"the server-specified delay\"} ms.`,\n+        );\n+        break;\n+      default:\n+        console.error(`Cosmos DB error ${statusCode}: ${error.message}`);\n+    }\n+  } else if (error instanceof Error) {\n+    console.error(error.message);\n+  } else {\n+    console.error(\"An unknown error occurred:\", error);\n+  }\n+\n+  process.exitCode = 1;\n+}\n+\n+void main().catch(handleError);\n*** Add File: README.md\n+# Azure Cosmos DB CRUD with TypeScript\n+\n+This sample creates a database and container, then creates, reads, queries,\n+replaces, and deletes an item with the Azure Cosmos DB for NoSQL SDK.\n+\n+## Requirements\n+\n+- Node.js 20 or later\n+- An Azure Cosmos DB for NoSQL endpoint and account key\n+\n+Install the required runtime package:\n+\n+```powershell\n+npm install @azure/cosmos\n+```\n+\n+Install all project dependencies and compile:\n+\n+```powershell\n+npm install\n+npm run build\n+```\n+\n+Set credentials without placing them in source control, then run the sample:\n+\n+```powershell\n+$env:COSMOS_ENDPOINT = \"https://\u003caccount\u003e.documents.azure.com:443/\"\n+$env:COSMOS_KEY = \"\u003caccount-key\u003e\"\n+npm start\n+```\n+\n+The program uses `createIfNotExists` for `TestDB` and `Items`, whose partition\n+key is `/category`. It creates a uniquely identified item and deletes it after\n+the CRUD sequence. For production workloads, prefer Microsoft Entra ID and a\n+managed identity over account-key authentication.\n+\n+SDK reference:\n+https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme\n*** End Patch\n"
```

### 🔧 11. powershell

**Input:**

```json
{"command":"npm install --silent \u0026\u0026 npm run build --silent","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n     const queryResponse = await container.items\n       .query\u003cInventoryItem\u003e(querySpec, { partitionKey: \"electronics\" })\n       .fetchAll();\n-    expectStatus(\"Query items\", queryResponse.statusCode, [200]);\n     console.log(\"Query results:\", queryResponse.resources);\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"npm run build --silent \u0026\u0026 $env:COSMOS_ENDPOINT=''; $env:COSMOS_KEY=''; node dist/index.js; if ($LASTEXITCODE -ne 1) { throw \"Expected missing configuration to exit with code 1, got $LASTEXITCODE\" }","description":"Compile and verify configuration error handling","initial_wait":120,"mode":"sync"}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"npm run build --silent; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; Remove-Item Env:COSMOS_ENDPOINT -ErrorAction SilentlyContinue; Remove-Item Env:COSMOS_KEY -ErrorAction SilentlyContinue; node dist\\index.js; if ($LASTEXITCODE -ne 1) { throw \"Expected missing configuration to exit with code 1, got $LASTEXITCODE\" }","description":"Compile and verify configuration error handling","initial_wait":120,"mode":"sync"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"npm run build --silent","description":"Compile TypeScript project","initial_wait":120,"mode":"sync"}
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the local project and SDK acceptance guidance, then add a runnable offline-safe TypeScript sample with status-aware error handling.

The workspace is empty. I’ll scaffold a minimal Node 20 TypeScript project, keep credentials in environment variables, and make every expected Cosmos response explicit.

The SDK’s query feed does not expose an HTTP status field; failed queries throw `ErrorResponse`. I’m adjusting that check while retaining explicit statuses for every CRUD/resource response.

Implemented a complete TypeScript Cosmos DB CRUD sample in `src/index.ts`.

- Uses `@azure/cosmos` 4.10.0 with endpoint/key environment variables.
- Creates `TestDB` and `Items` with partition key `/category`.
- Creates, reads, parameter-queries, replaces, and deletes an item.
- Checks expected HTTP status codes and handles common Cosmos errors including 400, 401/403, 404, 409, 412, and 429.
- Disposes `CosmosClient` reliably.
- Includes project configuration and usage instructions in `README.md`.

Install and run:

```powershell
npm install
npm run build
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
npm start
```

Reference: https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Includes the required Azure Cosmos DB SDK dependency.
- Provides TypeScript compiler configuration and Node.js type definitions.
- Defines build and start scripts for the TypeScript project.

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
  - Async/Await Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Pass
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
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 12.5% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Pagination with for-await-of` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 14.6% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 6.8571** | **62.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
