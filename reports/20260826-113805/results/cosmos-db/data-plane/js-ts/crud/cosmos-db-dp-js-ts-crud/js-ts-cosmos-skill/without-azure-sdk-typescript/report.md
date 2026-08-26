# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-cosmos-skill/without-azure-sdk-typescript | **Result:** ❌ FAILED | **Duration:** 588.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-cosmos-skill/without-azure-sdk-typescript |
| Result | ❌ FAILED |
| Score | 10/17 |
| Duration | 588.9s |
| Timestamp | 2026-08-26T03:47:51Z |
| Files Generated | 4 |
| Event Count | 4253 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 454.0s |
| Review | 133.7s |
| **Total** | **588.9s** |

## Configuration

- **name:** js-ts-cosmos-skill/without-azure-sdk-typescript
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=381231 out=3258 |
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

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local-only TypeScript example using the Azure Cosmos DB JavaScript SDK to perform CRUD against a NoSQL container, with endpoint/key environment variables and status-code-aware error handling.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices before creating a local TypeScript Cosmos DB CRUD example.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-cosmos-skill-without-azure-sdk-typescript-2858421137","pattern":"**/*"}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-cosmos-skill-without-azure-sdk-typescript-2858421137","pattern":"**/package.json"}
```

### 🔧 5. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-cosmos-skill-without-azure-sdk-typescript-2858421137","pattern":"**/tsconfig.json"}
```

### 🔧 6. powershell

**Input:**

```json
{"command":"npm view @azure/cosmos version","description":"Check latest Azure Cosmos SDK version","initial_wait":30}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"node --version; npm --version","description":"Check Node and npm versions","initial_wait":30}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"npm view typescript version","description":"Check latest TypeScript version","initial_wait":30}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"npm view @types/node version","description":"Check latest Node types version","initial_wait":30}
```

### 🔧 10. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-typescript-crud\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.10.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"noImplicitOverride\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import {\n+  CosmosClient,\n+  type ErrorResponse,\n+  type OperationResponse,\n+  type Resource,\n+} from \"@azure/cosmos\";\n+\n+interface Item {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+const endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+const key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+\n+const client = new CosmosClient({ endpoint, key });\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+  return value;\n+}\n+\n+function assertStatus\u003cT extends Resource\u003e(\n+  operation: string,\n+  response: OperationResponse\u003cT\u003e,\n+  expectedStatusCodes: readonly number[],\n+): void {\n+  if (!expectedStatusCodes.includes(response.statusCode)) {\n+    throw new Error(\n+      `${operation} returned HTTP ${response.statusCode}; expected ${expectedStatusCodes.join(\" or \")}`,\n+    );\n+  }\n+}\n+\n+function isCosmosError(error: unknown): error is ErrorResponse {\n+  return (\n+    error instanceof Error \u0026\u0026\n+    \"code\" in error \u0026\u0026\n+    typeof error.code === \"number\"\n+  );\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const databaseResponse = await client.databases.createIfNotExists({\n+    id: \"TestDB\",\n+  });\n+  assertStatus(\"Create database\", databaseResponse, [200, 201]);\n+  console.log(`Database ready: ${databaseResponse.database.id}`);\n+\n+  const containerResponse =\n+    await databaseResponse.database.containers.createIfNotExists({\n+      id: \"Items\",\n+      partitionKey: {\n+        paths: [\"/category\"],\n+      },\n+    });\n+  assertStatus(\"Create container\", containerResponse, [200, 201]);\n+  console.log(`Container ready: ${containerResponse.container.id}`);\n+\n+  const container = containerResponse.container;\n+  const newItem: Item = {\n+    id: \"item-1\",\n+    category: \"electronics\",\n+    name: \"Wireless keyboard\",\n+    quantity: 10,\n+  };\n+\n+  const createResponse = await container.items.create\u003cItem\u003e(newItem);\n+  assertStatus(\"Create item\", createResponse, [201]);\n+  console.log(\"Created item:\", createResponse.resource);\n+\n+  const item = container.item(newItem.id, newItem.category);\n+  const readResponse = await item.read\u003cItem\u003e();\n+  assertStatus(\"Read item\", readResponse, [200]);\n+  if (!readResponse.resource) {\n+    throw new Error(\"Read item returned HTTP 200 without a resource\");\n+  }\n+  console.log(\"Read item:\", readResponse.resource);\n+\n+  const query = {\n+    query: \"SELECT * FROM item WHERE item.category = @category\",\n+    parameters: [{ name: \"@category\", value: \"electronics\" }],\n+  };\n+  const queryResponse = await container.items.query\u003cItem\u003e(query).fetchAll();\n+  console.log(\n+    `Queried ${queryResponse.resources.length} electronics item(s):`,\n+    queryResponse.resources,\n+  );\n+\n+  const updatedItem: Item = {\n+    ...readResponse.resource,\n+    quantity: 15,\n+  };\n+  const replaceResponse = await item.replace\u003cItem\u003e(updatedItem);\n+  assertStatus(\"Replace item\", replaceResponse, [200]);\n+  console.log(\"Replaced item:\", replaceResponse.resource);\n+\n+  const deleteResponse = await item.delete();\n+  assertStatus(\"Delete item\", deleteResponse, [204]);\n+  console.log(`Deleted item: ${newItem.id}`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  if (isCosmosError(error)) {\n+    switch (error.code) {\n+      case 401:\n+      case 403:\n+        console.error(`Cosmos DB authorization failed (HTTP ${error.code}).`);\n+        break;\n+      case 404:\n+        console.error(\"Cosmos DB resource not found (HTTP 404).\");\n+        break;\n+      case 409:\n+        console.error(\"Cosmos DB resource already exists (HTTP 409).\");\n+        break;\n+      case 429:\n+        console.error(\n+          `Cosmos DB rate limit exceeded (HTTP 429). Retry after ${error.retryAfterInMs ?? \"the server-specified delay\"} ms.`,\n+        );\n+        break;\n+      default:\n+        console.error(\n+          `Cosmos DB request failed (HTTP ${error.code}): ${error.message}`,\n+        );\n+    }\n+  } else {\n+    console.error(\n+      \"Unexpected error:\",\n+      error instanceof Error ? error.message : error,\n+    );\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

### 🔧 11. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: tsconfig.json\n@@\n     \"rootDir\": \"src\",\n+    \"types\": [\"node\"],\n     \"strict\": true,\n*** Update File: src/index.ts\n@@\n import {\n   CosmosClient,\n   type ErrorResponse,\n   type OperationResponse,\n-  type Resource,\n } from \"@azure/cosmos\";\n@@\n-function assertStatus\u003cT extends Resource\u003e(\n+function assertStatus(\n   operation: string,\n-  response: OperationResponse\u003cT\u003e,\n+  response: OperationResponse,\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"npm run build","description":"Compile Cosmos CRUD example","initial_wait":120,"mode":"sync"}
```

### 🔧 14. rg

**Input:**

```json
{"-A":25,"-n":true,"glob":"*.d.ts","head_limit":10,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-cosmos-skill-without-azure-sdk-typescript-2858421137\\node_modules\\@azure\\cosmos","pattern":"class ErrorResponse|interface ErrorResponse|type ErrorResponse"}
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the existing project structure and package conventions, then add the smallest runnable Cosmos DB example and validate it locally without contacting Azure.

The workspace is empty. I’m creating a minimal Node/TypeScript project with environment-based credentials, explicit success-status assertions for each point operation, and typed Cosmos error reporting.

Implemented the TypeScript Cosmos DB CRUD program in `src\index.ts`.

- Required package: `@azure/cosmos` v4.10.0
- Creates `TestDB` and `Items` with partition key `/category`
- Performs create, point read, parameterized query, replace, and delete
- Checks expected HTTP statuses and handles common Cosmos DB errors
- Reads credentials from environment variables

```powershell
npm install
$env:COSMOS_ENDPOINT="https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY="<cosmos-key>"
npm run build
npm start
```

Project configuration is in `package.json` and `tsconfig.json`. Credentials are not hardcoded. Reference: [Azure Cosmos DB JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme?view=azure-node-latest).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Provides the required Azure Cosmos DB dependency and complete TypeScript project configuration.

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
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-cosmos-skill/without-azure-sdk-typescript
```

---

[← Back to Summary](../../../../../../summary.md)
