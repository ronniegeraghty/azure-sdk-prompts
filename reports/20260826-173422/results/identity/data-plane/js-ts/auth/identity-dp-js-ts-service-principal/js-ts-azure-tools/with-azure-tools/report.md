# Evaluation Report: identity-dp-js-ts-service-principal

**Config:** js-ts-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 248.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-js-ts-service-principal` |
| Config | js-ts-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 12/15 |
| Duration | 248.8s |
| Timestamp | 2026-08-26T11:54:13Z |
| Files Generated | 4 |
| Event Count | 2415 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 121.5s |
| Review | 126.1s |
| **Total** | **248.8s** |

## Configuration

- **name:** js-ts-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-ts, azure-keyvault-secrets-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=257663 out=2482 |
| Turn Count | 6 |

## Prompt Sent

```
Write a TypeScript program that
authenticates to Azure using a Service Principal with client secret. The program should:
1. Install and import @azure/identity with ClientSecretCredential
2. Read tenant ID, client ID, and client secret from environment variables
3. Create a ClientSecretCredential instance
4. Use it to create an Azure SDK client (e.g., SecretClient from @azure/keyvault-secrets)
5. Perform an operation to verify the credential works
6. Handle AuthenticationError for invalid credentials

Include a package.json with all dependencies, use dotenv for environment variable
management, and use async/await throughout.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-ts"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-keyvault-secrets-ts"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local-only TypeScript example authenticating with an Azure service principal client secret and using Azure Key Vault SecretClient, with dotenv and AuthenticationError handling.","learn":true}
```

### 🔧 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure TypeScript code-generation best practices for a local-only ClientSecretCredential and Key Vault SecretClient sample.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-identity-ts\\acceptance-criteria.md"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-identity-ts\\examples.md"}
```

### 🔧 7. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-secrets-ts\\acceptance-criteria.md"}
```

### 🔧 8. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-secrets-ts\\examples.md"}
```

### 🔧 9. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-js-ts-service-principal-js-ts-azure-tools-with-azure-tools-4016237598","pattern":"**/*"}
```

### 🔧 10. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-identity-ts\\references\\credential-types.md"}
```

### 🔧 11. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-secrets-ts\\references\\secrets.md"}
```

### 🔧 12. powershell

**Input:**

```json
{"command":"npm view @azure/identity version; npm view @azure/keyvault-secrets version; npm view dotenv version; npm view typescript version; npm view @types/node version; npm view tsx version","description":"Check current npm package versions","initial_wait":120,"mode":"sync"}
```

### 🔧 13. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-service-principal-key-vault-example\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.13.2\",\n+    \"@azure/keyvault-secrets\": \"^4.11.2\",\n+    \"dotenv\": \"^17.4.2\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import \"dotenv/config\";\n+\n+import {\n+  AuthenticationError,\n+  ClientSecretCredential,\n+} from \"@azure/identity\";\n+import { SecretClient } from \"@azure/keyvault-secrets\";\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name]?.trim();\n+\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+\n+  return value;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const tenantId = requireEnvironmentVariable(\"AZURE_TENANT_ID\");\n+  const clientId = requireEnvironmentVariable(\"AZURE_CLIENT_ID\");\n+  const clientSecret = requireEnvironmentVariable(\"AZURE_CLIENT_SECRET\");\n+  const keyVaultUrl = requireEnvironmentVariable(\"AZURE_KEY_VAULT_URL\");\n+\n+  const credential = new ClientSecretCredential(\n+    tenantId,\n+    clientId,\n+    clientSecret,\n+  );\n+  const secretClient = new SecretClient(keyVaultUrl, credential);\n+\n+  // Request one page of secret metadata to verify authentication and read access.\n+  const iterator = secretClient\n+    .listPropertiesOfSecrets()\n+    .byPage({ maxPageSize: 1 });\n+  await iterator.next();\n+\n+  console.log(\"Authentication succeeded and Key Vault is accessible.\");\n+}\n+\n+try {\n+  await main();\n+} catch (error: unknown) {\n+  if (error instanceof AuthenticationError) {\n+    console.error(\n+      \"Authentication failed. Check AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.\",\n+    );\n+    console.error(error.message);\n+    process.exitCode = 1;\n+  } else {\n+    throw error;\n+  }\n+}\n*** Add File: .env.example\n+AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000\n+AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000\n+AZURE_CLIENT_SECRET=replace-with-client-secret\n+AZURE_KEY_VAULT_URL=https://replace-with-vault-name.vault.azure.net\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+.env\n*** End Patch\n"
```

### 🔧 14. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll align the sample with the Azure Identity and Key Vault SDK guidance, then create and locally type-check the project without contacting Azure.

The sample will use a read-only Key Vault metadata request, so it verifies both token acquisition and vault authorization without changing secrets. Missing configuration and invalid service-principal credentials will fail with distinct messages.

Created a complete TypeScript project with:

- `ClientSecretCredential` using `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`
- `SecretClient` using `AZURE_KEY_VAULT_URL`
- A read-only Key Vault metadata request to verify authentication and access
- Explicit `AuthenticationError` handling
- `dotenv`, strict TypeScript configuration, `.env.example`, and `.gitignore`

Run it with:

```powershell
Copy-Item .env.example .env
# Update .env with real values
npm run build
npm start
```

The service principal needs permission to list Key Vault secret metadata, such as the **Key Vault Secrets User** role. References: [Azure Identity authentication](https://learn.microsoft.com/azure/developer/javascript/sdk/authentication/) and [Key Vault Secrets JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/keyvault-secrets-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete dependency manifest with all required Azure SDK and environment-management packages.
- Includes TypeScript compiler configuration and Node.js type definitions.

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - `@azure/identity` package with `ClientSecretCredential` class: Pass
      - Constructor parameters: tenantId, clientId, clientSecret: Pass
      - Passing credential to Azure SDK clients: Pass
      - dotenv or environment variable patterns: Pass
      - `AuthenticationError` handling: Pass
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
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Pagination with for-await-of` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 8.0000** | **72.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-js-ts-service-principal --config js-ts-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
