# Evaluation Report: identity-dp-js-ts-service-principal

**Config:** js-ts-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 261.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-js-ts-service-principal` |
| Config | js-ts-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 12/15 |
| Duration | 261.1s |
| Timestamp | 2026-08-26T11:45:31Z |
| Files Generated | 4 |
| Event Count | 2303 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 101.5s |
| Review | 158.4s |
| **Total** | **261.1s** |

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
| Token Usage | in=82009 out=1514 |
| Turn Count | 4 |

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

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. powershell

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace contents"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-service-principal-example\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"description\": \"Authenticate to Azure with a service principal and access Azure Key Vault.\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\",\n+    \"dev\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.0.0\",\n+    \"@azure/keyvault-secrets\": \"^4.0.0\",\n+    \"dotenv\": \"^16.0.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^22.0.0\",\n+    \"tsx\": \"^4.0.0\",\n+    \"typescript\": \"^5.0.0\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+AZURE_TENANT_ID=your-tenant-id\n+AZURE_CLIENT_ID=your-client-id\n+AZURE_CLIENT_SECRET=your-client-secret\n+AZURE_KEY_VAULT_URL=https://your-key-vault-name.vault.azure.net\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+.env\n*** Add File: src/index.ts\n+import \"dotenv/config\";\n+\n+import {\n+  AuthenticationError,\n+  ClientSecretCredential,\n+} from \"@azure/identity\";\n+import { SecretClient } from \"@azure/keyvault-secrets\";\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+\n+  if (!value) {\n+    throw new Error(`Missing required environment variable: ${name}`);\n+  }\n+\n+  return value;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const tenantId = requireEnvironmentVariable(\"AZURE_TENANT_ID\");\n+  const clientId = requireEnvironmentVariable(\"AZURE_CLIENT_ID\");\n+  const clientSecret = requireEnvironmentVariable(\"AZURE_CLIENT_SECRET\");\n+  const keyVaultUrl = requireEnvironmentVariable(\"AZURE_KEY_VAULT_URL\");\n+\n+  const credential = new ClientSecretCredential(\n+    tenantId,\n+    clientId,\n+    clientSecret,\n+  );\n+\n+  // Requesting a token directly makes invalid service-principal credentials\n+  // available as AuthenticationError before making the Key Vault request.\n+  await credential.getToken(\"https://vault.azure.net/.default\");\n+  console.log(\"Service principal authentication succeeded.\");\n+\n+  const secretClient = new SecretClient(keyVaultUrl, credential);\n+\n+  for await (const secret of secretClient.listPropertiesOfSecrets()) {\n+    console.log(`Key Vault access succeeded. Found secret: ${secret.name}`);\n+    return;\n+  }\n+\n+  console.log(\"Key Vault access succeeded. The vault contains no secrets.\");\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  if (error instanceof AuthenticationError) {\n+    console.error(\n+      \"Azure authentication failed. Check AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.\",\n+    );\n+    console.error(error.message);\n+    process.exitCode = 1;\n+    return;\n+  }\n+\n+  console.error(\"The Azure Key Vault operation failed:\", error);\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile project","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace, then create the TypeScript project and run its local build without contacting Azure.

The workspace is empty. I’ll add a minimal project that explicitly requests an Azure token to surface `AuthenticationError`, then uses the same credential with Key Vault to list one secret property.

The project files are in place. I’ll install the declared packages and compile the TypeScript code to confirm the SDK imports and error handling match the installed versions.

Created a complete TypeScript project with:

- `ClientSecretCredential` and `AuthenticationError` handling
- Environment variables loaded via `dotenv`
- `SecretClient` for Azure Key Vault
- Token authentication and Key Vault access verification
- `package.json`, `tsconfig.json`, `.env.example`, and `.gitignore`

Dependencies are installed and the project compiles successfully. Copy `.env.example` to `.env`, provide valid Azure credentials and a Key Vault URL, then run:

```powershell
npm run dev
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Includes all required runtime and TypeScript development dependencies.
- Provides a valid strict TypeScript configuration.

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
hyoka run --prompt-id identity-dp-js-ts-service-principal --config js-ts-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
