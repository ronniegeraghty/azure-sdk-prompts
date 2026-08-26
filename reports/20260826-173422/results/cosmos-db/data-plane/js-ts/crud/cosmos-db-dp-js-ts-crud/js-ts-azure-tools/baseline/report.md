# Evaluation Report: cosmos-db-dp-js-ts-crud

**Config:** js-ts-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 251.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-js-ts-crud` |
| Config | js-ts-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 10/17 |
| Duration | 251.1s |
| Timestamp | 2026-08-26T11:15:17Z |
| Files Generated | 4 |
| Event Count | 3516 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 108.7s |
| Review | 141.6s |
| **Total** | **251.1s** |

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
| Token Usage | in=127597 out=2435 |
| Turn Count | 6 |

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

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-tools-baseline-3078958800","pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-tools-baseline-3078958800","pattern":"**/package.json"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-js-ts-crud-js-ts-azure-tools-baseline-3078958800","pattern":"**/tsconfig.json"}
```

### 🔧 4. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"cosmos-db-typescript-crud\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/cosmos\": \"^4.3.1\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^22.10.2\",\n+    \"typescript\": \"^5.7.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import { CosmosClient } from \"@azure/cosmos\";\n+\n+interface Item {\n+  id: string;\n+  category: string;\n+  name: string;\n+  quantity: number;\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(\n+      `${name} is required. Set it to the corresponding Cosmos DB emulator value.`,\n+    );\n+  }\n+\n+  return value;\n+}\n+\n+function requireStatus(\n+  operation: string,\n+  actual: number | undefined,\n+  expected: readonly number[],\n+): void {\n+  if (actual === undefined || !expected.includes(actual)) {\n+    throw new Error(\n+      `${operation} returned status ${actual ?? \"unknown\"}; expected ${expected.join(\" or \")}.`,\n+    );\n+  }\n+}\n+\n+function getStatusCode(error: unknown): number | undefined {\n+  if (typeof error !== \"object\" || error === null) {\n+    return undefined;\n+  }\n+\n+  const statusCode = Reflect.get(error, \"code\") ?? Reflect.get(error, \"statusCode\");\n+  return typeof statusCode === \"number\" ? statusCode : undefined;\n+}\n+\n+function getErrorMessage(error: unknown): string {\n+  return error instanceof Error ? error.message : String(error);\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const endpoint = requireEnvironmentVariable(\"COSMOS_ENDPOINT\");\n+  const key = requireEnvironmentVariable(\"COSMOS_KEY\");\n+  const client = new CosmosClient({ endpoint, key });\n+\n+  const databaseResponse = await client.databases.createIfNotExists({\n+    id: \"TestDB\",\n+  });\n+  requireStatus(\"Create database\", databaseResponse.statusCode, [200, 201]);\n+\n+  const containerResponse =\n+    await databaseResponse.database.containers.createIfNotExists({\n+      id: \"Items\",\n+      partitionKey: { paths: [\"/category\"] },\n+    });\n+  requireStatus(\"Create container\", containerResponse.statusCode, [200, 201]);\n+\n+  const container = containerResponse.container;\n+  const item: Item = {\n+    id: \"item-1\",\n+    category: \"electronics\",\n+    name: \"Wireless headphones\",\n+    quantity: 10,\n+  };\n+\n+  const createResponse = await container.items.create\u003cItem\u003e(item);\n+  requireStatus(\"Create item\", createResponse.statusCode, [201]);\n+  console.log(\"Created:\", createResponse.resource);\n+\n+  const itemReference = container.item(item.id, item.category);\n+  const readResponse = await itemReference.read\u003cItem\u003e();\n+  requireStatus(\"Read item\", readResponse.statusCode, [200]);\n+  if (!readResponse.resource) {\n+    throw new Error(\"Read item succeeded but returned no resource.\");\n+  }\n+  console.log(\"Read:\", readResponse.resource);\n+\n+  const query = {\n+    query: \"SELECT * FROM c WHERE c.category = @category\",\n+    parameters: [{ name: \"@category\", value: \"electronics\" }],\n+  };\n+  const queryResponse = await container.items\n+    .query\u003cItem\u003e(query)\n+    .fetchAll();\n+  requireStatus(\"Query items\", queryResponse.statusCode, [200]);\n+  console.log(\"Query results:\", queryResponse.resources);\n+\n+  const updatedItem: Item = {\n+    ...readResponse.resource,\n+    quantity: 15,\n+  };\n+  const replaceResponse = await itemReference.replace\u003cItem\u003e(updatedItem);\n+  requireStatus(\"Replace item\", replaceResponse.statusCode, [200]);\n+  console.log(\"Updated:\", replaceResponse.resource);\n+\n+  const deleteResponse = await itemReference.delete();\n+  requireStatus(\"Delete item\", deleteResponse.statusCode, [204]);\n+  console.log(`Deleted item ${item.id}.`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  const statusCode = getStatusCode(error);\n+\n+  if (statusCode === 404) {\n+    console.error(\"Cosmos DB resource not found:\", getErrorMessage(error));\n+  } else if (statusCode === 409) {\n+    console.error(\"Cosmos DB resource already exists:\", getErrorMessage(error));\n+  } else if (statusCode === 429) {\n+    console.error(\"Cosmos DB request was rate limited:\", getErrorMessage(error));\n+  } else if (statusCode !== undefined) {\n+    console.error(\n+      `Cosmos DB request failed with status ${statusCode}:`,\n+      getErrorMessage(error),\n+    );\n+  } else {\n+    console.error(\"Unexpected error:\", getErrorMessage(error));\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

### 🔧 5. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund --quiet \u0026\u0026 npm run build --silent","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

### 🔧 6. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n   const queryResponse = await container.items\n     .query\u003cItem\u003e(query)\n     .fetchAll();\n-  requireStatus(\"Query items\", queryResponse.statusCode, [200]);\n   console.log(\"Query results:\", queryResponse.resources);\n*** End Patch\n"
```

### 🔧 7. powershell

**Input:**

```json
{"command":"npm run build --silent","description":"Compile TypeScript project","initial_wait":120,"mode":"sync"}
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace, then add a runnable local TypeScript example with package metadata and Cosmos-specific error handling.

The workspace is empty. I’ll create a minimal project whose credentials come from environment variables, so it can target the local Cosmos DB emulator without embedding secrets or provisioning Azure resources.

The implementation is in place. I’ll restore the declared dependencies and compile it to catch SDK API or type mismatches.

The SDK’s query result type does not expose an HTTP status code; query failures are thrown instead. I’ll keep explicit status checks for point operations and rely on the SDK error path for the parameterized query.

Implemented a complete TypeScript Cosmos DB CRUD example in `src\index.ts`.

- Uses `@azure/cosmos`
- Creates `TestDB` and `Items` with partition key `/category`
- Creates, reads, parameterized-queries, replaces, and deletes an item
- Checks expected HTTP status codes
- Handles `404`, `409`, `429`, and other Cosmos errors
- Reads credentials from `COSMOS_ENDPOINT` and `COSMOS_KEY`

Run against a local Cosmos DB emulator:

```powershell
$env:COSMOS_ENDPOINT = "https://localhost:8081"
$env:COSMOS_KEY = "<emulator-key>"
npm install
npm run build
npm start
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Includes the required Azure Cosmos DB package and complete TypeScript project configuration.
- Provides appropriate Node.js type definitions and a recent TypeScript version.

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
hyoka run --prompt-id cosmos-db-dp-js-ts-crud --config js-ts-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
