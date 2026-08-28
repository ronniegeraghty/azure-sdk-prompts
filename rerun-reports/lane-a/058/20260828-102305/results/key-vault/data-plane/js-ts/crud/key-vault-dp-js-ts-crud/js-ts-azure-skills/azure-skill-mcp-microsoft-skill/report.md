# Evaluation Report: key-vault-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 390.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 12/15 |
| Duration | 390.7s |
| Timestamp | 2026-08-28T02:23:05Z |
| Files Generated | 5 |
| Event Count | 3091 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 260.8s |
| Review | 128.7s |
| **Total** | **390.7s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-keyvault-secrets-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=204993 out=2253 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 653ms

**Input:**

```json
{"skill":"azure-keyvault-secrets-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Key Vault Secrets SDK for TypeScript

Manage secrets with Azure Key Vault.

## Installation

```bash
# Secrets SDK
npm install @azure/keyvault-secrets @azure/identity
```

## Environment Variables

```bash
KEY_VAULT_URL=https://<vault-name>.vault.azure.net
# Or
AZURE_KEYVAULT_NAME=<vault-name>
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

```typescript
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";
import { SecretClient } from "@azure/keyvault-secrets";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();
const vaultUrl = `https://${process.env.AZURE_KEYVAULT_NAME}.vault.azure.net`;

const keyClient = new KeyClient(vaultUrl, credential);
const secretClient = new SecretClient(vaultUrl, credential);
```

## Secrets Operations

### Create/Set Secret

```typescript
const secret = await secretClient.setSecret("MySecret", "secret-value");

// With attributes
const secretWithAttrs = await secretClient.setSecret("MySecret", "value", {
  enabled: true,
  expiresOn: new Date("2025-12-31"),
  contentType: "application/json",
  tags: { environment: "production" }
});
```

### Get Secret

```typescript
// Get latest version
const secret = await secretClient.getSecret("MySecret");
console.log(secret.value);

// Get specific version
const specificSecret = await secretClient.getSecret("MySecret", {
  version: secret.properties.version
});
```

### List Secrets

```typescript
for await (const secretProperties of secretClient.listPropert
... (truncated)
```

### ✅ 2. view — 182526ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-secrets-ts\\references\\secrets.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/secrets.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/secrets.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/secrets.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/secrets.md
@@ -1,474 +1,474 @@
 # Secrets Reference
 
 Secret management operations using @azure/keyvault-secrets SDK.
 
 ## Overview
 
 The SecretClient provides operations for managing secrets in Azure Key Vault:
 - Create, update, and delete secrets
 - List secrets and versions
 - Soft-delete and purge operations
 - Backup and restore capabilities
 
 ## Core Types
 
 ```typescript
 import {
   SecretClient,
   KeyVaultSecret,
   SecretProperties,
   DeletedSecret,
   SetSecretOptions,
   GetSecretOptions,
   UpdateSecretPropertiesOptions,
   BeginDeleteSecretOptions,
   BeginRecoverDeletedSecretOptions,
   ListPropertiesOfSecretsOptions,
   ListPropertiesOfSecretVersionsOptions,
   ListDeletedSecretsOptions
 } from "@azure/keyvault-secrets";
 ```
 
 ## SecretClient Initialization
 
 ```typescript
 import { SecretClient } from "@azure/keyvault-secrets";
 import { DefaultAzureCredential } from "@azure/identity";
 
 const vaultUrl = `https://${process.env.AZURE_KEYVAULT_NAME}.vault.azure.net`;
 const credential = new DefaultAzureCredential();
 
 const secretClient = new SecretClient(vaultUrl, credential);
 ```
 
 ## Creating and Updating Secrets
 
 ### Set Secret (Create or Update)
 
 ```typescript
 // Basic secret
 const secret = await secretClient.setSecret("MySecret", "secret-value");
 console.log(`Secret: ${secret.name}
... (truncated)
```

### ✅ 3. view — 182174ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-secrets-ts\\references\\keys.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/keys.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/keys.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/keys.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-secrets-ts/references/keys.md
@@ -1,458 +1,458 @@
 # Keys Reference
 
 Cryptographic key management and operations using @azure/keyvault-keys SDK.
 
 ## Overview
 
 The Key Vault Keys SDK provides two main clients:
 - **KeyClient** - CRUD operations for keys (create, get, list, rotate, delete)
 - **CryptographyClient** - Cryptographic operations using keys (encrypt, decrypt, sign, verify, wrap, unwrap)
 
 ## Core Types
 
 ```typescript
 import {
   KeyClient,
   CryptographyClient,
   KeyVaultKey,
   KeyProperties,
   DeletedKey,
   KeyRotationPolicy,
   KeyRotationPolicyProperties,
   KeyRotationLifetimeAction,
   CreateKeyOptions,
   CreateRsaKeyOptions,
   CreateEcKeyOptions,
   EncryptParameters,
   DecryptParameters,
   SignResult,
   VerifyResult,
   WrapResult,
   UnwrapResult,
   KnownEncryptionAlgorithms,
   KnownSignatureAlgorithms,
   KnownKeyTypes,
   KnownKeyCurveNames
 } from "@azure/keyvault-keys";
 ```
 
 ## KeyClient Initialization
 
 ```typescript
 import { KeyClient } from "@azure/keyvault-keys";
 import { DefaultAzureCredential } from "@azure/identity";
 
 const vaultUrl = `https://${process.env.AZURE_KEYVAULT_NAME}.vault.azure.net`;
 const credential = new DefaultAzureCredential();
 
 const keyClient = new KeyClient(vaultUrl, credential);
 ```
 
 ## Creating Keys
 
 ### RSA Keys
 
 ```typescript
 // Basic RSA key (default 2048-bit)
 
... (truncated)
```

### ✅ 4. glob — 181815ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 182528ms

**Input:**

```json
{"intent":"Generate a local TypeScript Node.js script using DefaultAzureCredential and @azure/keyvault-secrets to create, read, update, soft-delete, and purge one Azure Key Vault secret, with proper error handling. Do not deploy or modify resources during generation.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. apply_patch — 755ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-key-vault-secret-crud\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.0.0\",\n+    \"@azure/keyvault-secrets\": \"^4.0.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.0.0\",\n+    \"tsx\": \"^4.0.0\",\n+    \"typescript\": \"^5.0.0\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true,\n+    \"outDir\": \"dist\"\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import { DefaultAzureCredential } from \"@azure/identity\";\n+import { SecretClient } from \"@azure/keyvault-secrets\";\n+\n+const secretName = \"my-secret\";\n+\n+function getVaultUrl(): string {\n+  const vaultUrl = process.env.KEY_VAULT_URL;\n+\n+  if (!vaultUrl) {\n+    throw new Error(\n+      \"KEY_VAULT_URL is required (for example, https://my-vault.vault.azure.net).\",\n+    );\n+  }\n+\n+  let parsedUrl: URL;\n+  try {\n+    parsedUrl = new URL(vaultUrl);\n+  } catch {\n+    throw new Error(\"KEY_VAULT_URL must be a valid URL.\");\n+  }\n+\n+  if (\n+    parsedUrl.protocol !== \"https:\" ||\n+    !parsedUrl.hostname.endsWith(\".vault.azure.net\")\n+  ) {\n+    throw new Error(\n+      \"KEY_VAULT_URL must be an HTTPS Azure Key Vault URL ending in .vault.azure.net.\",\n+    );\n+  }\n+\n+  return parsedUrl.toString().replace(/\\/$/, \"\");\n+}\n+\n+function describeError(error: unknown): string {\n+  if (!(error instanceof Error)) {\n+    return String(error);\n+  }\n+\n+  const azureError = error as Error \u0026 {\n+    code?: string;\n+    statusCode?: number;\n+  };\n+  const details = [\n+    azureError.statusCode \u0026\u0026 `HTTP ${azureError.statusCode}`,\n+    azureError.code,\n+  ].filter(Boolean);\n+\n+  return details.length \u003e 0\n+    ? `${azureError.message} (${details.join(\", \")})`\n+    : azureError.message;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const credential = new DefaultAzureCredential();\n+  const client = new SecretClient(getVaultUrl(), credential);\n+\n+  try {\n+    const created = await client.setSecret(secretName, \"my-secret-value\");\n+    console.log(\n+      `Created \"${created.name}\" (version ${created.properties.version}).`,\n+    );\n+\n+    const read = await client.getSecret(secretName);\n+    if (read.value === undefined) {\n+      throw new Error(`Secret \"${secretName}\" was returned without a value.`);\n+    }\n+    console.log(`Read \"${read.name}\": ${read.value}`);\n+\n+    const updated = await client.setSecret(secretName, \"updated-value\");\n+    console.log(\n+      `Updated \"${updated.name}\" (version ${updated.properties.version}).`,\n+    );\n+\n+    const deletePoller = await client.beginDeleteSecret(secretName);\n+    await deletePoller.pollUntilDone();\n+    console.log(`Soft-deleted \"${secretName}\".`);\n+\n+    await client.purgeDeletedSecret(secretName);\n+    console.log(`Purged \"${secretName}\".`);\n+  } catch (error: unknown) {\n+    console.error(`Key Vault operation failed: ${describeError(error)}`);\n+    throw error;\n+  }\n+}\n+\n+main().catch(() =\u003e {\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Key Vault secret CRUD (TypeScript)\n+\n+## Required packages\n+\n+```powershell\n+npm install @azure/identity @azure/keyvault-secrets\n+npm install --save-dev typescript tsx @types/node\n+```\n+\n+## Run\n+\n+Set the vault URL and authenticate with any identity supported by\n+`DefaultAzureCredential`, such as Azure CLI credentials for local development:\n+\n+```powershell\n+$env:KEY_VAULT_URL = \"https://\u003cvault-name\u003e.vault.azure.net\"\n+npm install\n+npm start\n+```\n+\n+The identity needs secret `get`, `set`, `delete`, and `purge` permissions. With\n+Azure RBAC, the **Key Vault Secrets Officer** role includes these operations.\n+Purging is irreversible and fails when purge protection is enabled; in that\n+case, the soft-deleted secret remains until the retention period expires.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/package.json
@@ -1,0 +1,20 @@
+{
+  "name": "azure-key-vault-secret-crud",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "tsx src/index.ts"
+  },
+  "dependencies": {
+    "@azure/identity": "^4.0.0",
+    "@azure/keyvault-secrets": "^4.0.0"
+  },
+  "devDependencies": {
+    "@types/node": "^24.0.0",
+    "tsx": "^4.0.0",
+    "typescript": "^5.0.0"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/tsconfig.json
@@ -1,0 +1,15 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "skipLibCheck": true,
+    "outDir": "dist"
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-110269406/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-
... (truncated)
```

### ✅ 7. powershell — 25640ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 54 packages in 22s

> azure-key-vault-secret-crud@1.0.0 build
> tsc
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the SDK guidance and workspace, then add the smallest runnable TypeScript example with package metadata.

The workspace is empty. I’m adding a focused TypeScript project that validates configuration, runs the CRUD sequence, and surfaces Azure HTTP errors without hiding failures.

The implementation is in place; I’m resolving the declared packages and compiling it locally to catch SDK or TypeScript incompatibilities without contacting Azure.

Implemented a runnable TypeScript example in `src/index.ts` covering:

1. Create `my-secret` with `my-secret-value`
2. Read and print its value
3. Update it to `updated-value`
4. Soft-delete and permanently purge it

It uses `DefaultAzureCredential`, validates `KEY_VAULT_URL`, waits for deletion completion, and reports Azure error details through `try/catch`.

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
npm install
npm start
```

Required runtime packages are `@azure/identity` and `@azure/keyvault-secrets`. Purging requires the appropriate RBAC permission and will fail if purge protection is enabled. No live Azure operations were executed.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete TypeScript project metadata with the required Azure SDK dependencies, Node.js type definitions, TypeScript tooling, and compiler configuration.

## Grader Results

- crud-secrets.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/5)
      - Installing `@azure/keyvault-secrets` and `@azure/identity` npm packages: Pass
      - Creating a `SecretClient` with vault URL and credential: Pass
      - `setSecret()`, `getSecret()`, `beginDeleteSecret()`, `purgeDeletedSecret()`: Pass
      - Awaiting the `DeleteSecretPoller` before purging: Pass
      - Error handling for `RestError`: Fail
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
| `Criteria from prompt file` | prompt_review | 80% | 1.00 | 0.8000 | 9.1% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 11.4% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 8.8000** | **80.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
