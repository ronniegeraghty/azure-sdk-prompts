# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-cosmos-skill/baseline | **Result:** ❌ FAILED | **Duration:** 584.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-cosmos-skill/baseline |
| Result | ❌ FAILED |
| Score | 12/17 |
| Duration | 584.2s |
| Timestamp | 2026-08-26T03:38:05Z |
| Files Generated | 4 |
| Event Count | 3931 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 342.8s |
| Review | 240.4s |
| **Total** | **584.2s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** js-ts-cosmos-skill/baseline

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-cosmos-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=329118 out=3166 |
| Turn Count | 8 |

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
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-cosmos-skill-baseline-3127876017","pattern":"**/*"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-cosmos-ts","pattern":"**/{acceptance-criteria.md,examples.md,query-patterns.md}"}
```

### 🔧 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local-only TypeScript sample using @azure/cosmos with key authentication, database/container setup, CRUD, parameterized queries, and status-code-aware error handling.","learn":true}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only TypeScript Azure Cosmos DB CRUD sample with parameterized queries and robust error handling.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-cosmos-ts\\references\\query-patterns.md"}
```

### 🔧 7. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-crud-sample\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc --noEmit\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.9.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^22.0.0\",\n+    \"tsx\": \"^4.20.0\",\n+    \"typescript\": \"^5.8.0\"\n+  },\n+  \"engines\": {\n+    \"node\": \"\u003e=20\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true,\n+    \"outDir\": \"dist\"\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+COSMOS_ENDPOINT=https://your-account.documents.azure.com:443/\n+COSMOS_KEY=your-account-key\n*** Add File: src/index.ts\n+import { randomUUID } from \"node:crypto\";\n+import {\n+  CosmosClient,\n+  ErrorResponse,\n+  type SqlQuerySpec,\n+} from \"@azure/cosmos\";\n+\n+interface InventoryItem {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+const endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+const key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+\n+const client = new CosmosClient({ endpoint, key });\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Required environment variable ${name} is not set.`);\n+  }\n+  return value;\n+}\n+\n+function assertStatus(\n+  operation: string,\n+  actual: number,\n+  expected: readonly number[],\n+): void {\n+  if (!expected.includes(actual)) {\n+    throw new Error(\n+      `${operation} returned status ${actual}; expected ${expected.join(\" or \")}.`,\n+    );\n+  }\n+}\n+\n+function describeCosmosError(error: ErrorResponse): string {\n+  const statusCode = Number(error.code);\n+\n+  switch (statusCode) {\n+    case 400:\n+      return \"Bad request. Check item data and partition key configuration.\";\n+    case 401:\n+      return \"Unauthorized. Check COSMOS_ENDPOINT and COSMOS_KEY.\";\n+    case 403:\n+      return \"Forbidden. The supplied key lacks permission for this operation.\";\n+    case 404:\n+      return \"The requested database, container, or item was not found.\";\n+    case 409:\n+      return \"Conflict. A resource with the same ID already exists.\";\n+    case 412:\n+      return \"Precondition failed. The resource was modified by another operation.\";\n+    case 429:\n+      return `Rate limited. Retry after ${error.retryAfterInMs ?? \"the advised delay\"} ms.`;\n+    default:\n+      return `Cosmos DB request failed with status ${statusCode}.`;\n+  }\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const databaseResponse = await client.databases.createIfNotExists({\n+    id: \"TestDB\",\n+  });\n+  assertStatus(\"Create database\", databaseResponse.statusCode, [200, 201]);\n+\n+  const containerResponse =\n+    await databaseResponse.database.containers.createIfNotExists({\n+      id: \"Items\",\n+      partitionKey: { paths: [\"/category\"] },\n+    });\n+  assertStatus(\"Create container\", containerResponse.statusCode, [200, 201]);\n+\n+  const container = containerResponse.container;\n+  const newItem: InventoryItem = {\n+    id: randomUUID(),\n+    category: \"electronics\",\n+    name: \"Wireless keyboard\",\n+    quantity: 10,\n+  };\n+\n+  const createResponse =\n+    await container.items.create\u003cInventoryItem\u003e(newItem);\n+  assertStatus(\"Create item\", createResponse.statusCode, [201]);\n+  if (!createResponse.resource) {\n+    throw new Error(\"Create item succeeded without returning the item.\");\n+  }\n+  console.log(\"Created:\", createResponse.resource);\n+\n+  const itemReference = container.item(newItem.id, newItem.category);\n+  const readResponse = await itemReference.read\u003cInventoryItem\u003e();\n+  assertStatus(\"Read item\", readResponse.statusCode, [200]);\n+  if (!readResponse.resource) {\n+    throw new Error(\"Read item succeeded without returning the item.\");\n+  }\n+  console.log(\"Read:\", readResponse.resource);\n+\n+  const query: SqlQuerySpec = {\n+    query: \"SELECT * FROM c WHERE c.category = @category\",\n+    parameters: [{ name: \"@category\", value: \"electronics\" }],\n+  };\n+  const queryResponse = await container.items\n+    .query\u003cInventoryItem\u003e(query, { partitionKey: \"electronics\" })\n+    .fetchAll();\n+  assertStatus(\"Query items\", queryResponse.statusCode, [200]);\n+  console.log(\"Query results:\", queryResponse.resources);\n+\n+  const updatedItem: InventoryItem = {\n+    ...readResponse.resource,\n+    quantity: 25,\n+  };\n+  const replaceResponse =\n+    await itemReference.replace\u003cInventoryItem\u003e(updatedItem);\n+  assertStatus(\"Replace item\", replaceResponse.statusCode, [200]);\n+  if (!replaceResponse.resource) {\n+    throw new Error(\"Replace item succeeded without returning the updated item.\");\n+  }\n+  console.log(\"Replaced:\", replaceResponse.resource);\n+\n+  const deleteResponse = await itemReference.delete();\n+  assertStatus(\"Delete item\", deleteResponse.statusCode, [204]);\n+  console.log(`Deleted item ${newItem.id}.`);\n+}\n+\n+try {\n+  await main();\n+} catch (error: unknown) {\n+  if (error instanceof ErrorResponse) {\n+    console.error(describeCosmosError(error));\n+    console.error(error.message);\n+    process.exitCode = 1;\n+  } else if (error instanceof Error) {\n+    console.error(error.message);\n+    process.exitCode = 1;\n+  } else {\n+    console.error(\"An unknown error occurred.\", error);\n+    process.exitCode = 1;\n+  }\n+} finally {\n+  client.dispose();\n+}\n*** End Patch\n"
```

### 🔧 8. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund --silent \u0026\u0026 npm run build --silent","description":"Install dependencies and type-check sample","initial_wait":180,"mode":"sync"}
```

### 🔧 9. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n   const queryResponse = await container.items\n     .query\u003cInventoryItem\u003e(query, { partitionKey: \"electronics\" })\n     .fetchAll();\n-  assertStatus(\"Query items\", queryResponse.statusCode, [200]);\n   console.log(\"Query results:\", queryResponse.resources);\n*** End Patch\n"
```

### 🔧 10. powershell

**Input:**

```json
{"command":"npm run build --silent","description":"Type-check Cosmos DB sample","initial_wait":120,"mode":"sync"}
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll scaffold a runnable local TypeScript sample with environment-based credentials, typed CRUD operations, and explicit success/error status handling.

Implemented a complete TypeScript Cosmos DB CRUD sample in `src\index.ts`.

- Uses `@azure/cosmos` with endpoint/key from `COSMOS_ENDPOINT` and `COSMOS_KEY`
- Creates `TestDB` and `Items` with partition key `/category`
- Creates, reads, queries, replaces, and deletes an item
- Uses a parameterized, partition-scoped query
- Checks operation status codes and handles Cosmos errors including 400, 401, 403, 404, 409, 412, and 429
- Disposes the client reliably

Run it with:

```powershell
$env:COSMOS_ENDPOINT = "https://your-account.documents.azure.com:443/"
$env:COSMOS_KEY = "your-account-key"
npm start
```

The required package is declared as `@azure/cosmos` in `package.json`. Reference: [Azure Cosmos DB Node.js quickstart](https://learn.microsoft.com/azure/cosmos-db/nosql/quickstart-nodejs).

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Complete TypeScript project metadata with the required Azure Cosmos DB SDK, Node.js typings, TypeScript compiler, runtime tooling, and tsconfig.json.

### Issues

- The package declares Node.js >=20, while the lockfile resolves @azure/cosmos 4.10.0 and several transitive dependencies that require Node.js >=22.

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
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-cosmos-skill/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
