# Evaluation Report: app-configuration-dp-js-ts-crud

**Config:** js-ts-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 282.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-js-ts-crud` |
| Config | js-ts-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 12/18 |
| Duration | 282.8s |
| Timestamp | 2026-08-26T11:05:54Z |
| Files Generated | 4 |
| Event Count | 3376 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 149.2s |
| Review | 132.7s |
| **Total** | **282.8s** |

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
| Token Usage | in=271206 out=2970 |
| Turn Count | 11 |

## Prompt Sent

```
Write a TypeScript program
that manages configuration settings in Azure App Configuration:
1. Create an AppConfigurationClient using a connection string
2. Set a configuration setting with key "app:Settings:FontSize" and value "24"
3. Set a setting with label "Production"
4. Get the setting by key and print its value
5. List all settings matching the key filter "app:Settings:*" using listConfigurationSettings
6. Create a feature flag configuration setting for "BetaFeature"
7. Delete the setting by key

Show required npm package (@azure/app-configuration) and
proper error handling with RestError.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-tools-baseline-970527954","pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-tools-baseline-970527954","pattern":"**/package.json"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-tools-baseline-970527954","pattern":"**/tsconfig.json"}
```

### 🔧 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Current @azure/app-configuration TypeScript API featureFlagContentType featureFlagPrefix setConfigurationSetting example RestError 2026"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://www.npmjs.com/package/@azure/app-configuration"}
```

### 🔧 6. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-js/main/sdk/appconfiguration/app-configuration/samples/v1/typescript/src/featureFlag.ts"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-js/main/sdk/appconfiguration/app-configuration/src/index.ts"}
```

### 🔧 8. web_fetch

**Input:**

```json
{"max_length":8000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-js/main/sdk/appconfiguration/app-configuration/package.json"}
```

### 🔧 9. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-app-configuration-example\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/app-configuration\": \"^1.12.2\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^22.18.0\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import {\n+  AppConfigurationClient,\n+  type ConfigurationSetting,\n+  type FeatureFlagValue,\n+  RestError,\n+  featureFlagContentType,\n+  featureFlagPrefix,\n+} from \"@azure/app-configuration\";\n+\n+const settingKey = \"app:Settings:FontSize\";\n+const settingValue = \"24\";\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const connectionString = process.env.AZURE_APP_CONFIGURATION_CONNECTION_STRING;\n+  if (!connectionString) {\n+    throw new Error(\n+      \"Set AZURE_APP_CONFIGURATION_CONNECTION_STRING before running this program.\",\n+    );\n+  }\n+\n+  const client = new AppConfigurationClient(connectionString);\n+\n+  await client.setConfigurationSetting({\n+    key: settingKey,\n+    value: settingValue,\n+  });\n+\n+  await client.setConfigurationSetting({\n+    key: settingKey,\n+    value: settingValue,\n+    label: \"Production\",\n+  });\n+\n+  const setting = await client.getConfigurationSetting({ key: settingKey });\n+  console.log(`${setting.key} = ${setting.value}`);\n+\n+  console.log(\"Matching configuration settings:\");\n+  for await (const matchingSetting of client.listConfigurationSettings({\n+    keyFilter: \"app:Settings:*\",\n+  })) {\n+    console.log(\n+      `${matchingSetting.key} [${matchingSetting.label ?? \"no label\"}] = ${matchingSetting.value}`,\n+    );\n+  }\n+\n+  const betaFeature: ConfigurationSetting\u003cFeatureFlagValue\u003e = {\n+    key: `${featureFlagPrefix}BetaFeature`,\n+    contentType: featureFlagContentType,\n+    value: {\n+      id: \"BetaFeature\",\n+      enabled: true,\n+      description: \"Enables the beta feature.\",\n+      conditions: {\n+        clientFilters: [],\n+      },\n+    },\n+  };\n+  await client.setConfigurationSetting(betaFeature);\n+\n+  await client.deleteConfigurationSetting({ key: settingKey });\n+  console.log(`Deleted the unlabeled setting \"${settingKey}\".`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  if (error instanceof RestError) {\n+    console.error(\n+      `Azure App Configuration request failed (${error.statusCode ?? \"no status\"}): ${error.message}`,\n+    );\n+  } else if (error instanceof Error) {\n+    console.error(`Program failed: ${error.message}`);\n+  } else {\n+    console.error(\"Program failed with an unknown error.\");\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

### 🔧 10. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"npm view @azure/app-configuration version --json; npm view @azure/app-configuration versions --json","description":"Check published SDK versions","initial_wait":30,"mode":"sync"}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: package.json\n@@\n-    \"@azure/app-configuration\": \"^1.12.2\"\n+    \"@azure/app-configuration\": \"^1.12.1\"\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install published SDK and compile","initial_wait":180,"mode":"sync"}
```

### 🔧 14. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: package.json\n@@\n   \"dependencies\": {\n-    \"@azure/app-configuration\": \"^1.12.1\"\n+    \"@azure/app-configuration\": \"^1.12.1\",\n+    \"@azure/core-rest-pipeline\": \"^1.22.1\"\n*** Update File: src/index.ts\n@@\n   type ConfigurationSetting,\n   type FeatureFlagValue,\n-  RestError,\n   featureFlagContentType,\n   featureFlagPrefix,\n } from \"@azure/app-configuration\";\n+import { RestError } from \"@azure/core-rest-pipeline\";\n@@\n   const betaFeature: ConfigurationSetting\u003cFeatureFlagValue\u003e = {\n     key: `${featureFlagPrefix}BetaFeature`,\n+    isReadOnly: false,\n     contentType: featureFlagContentType,\n*** End Patch\n"
```

### 🔧 15. powershell

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Update dependencies and compile example","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace, then add a runnable TypeScript example and dependency metadata without contacting Azure.

The workspace is empty. I’m confirming the current SDK’s feature-flag and `RestError` APIs before creating the example.

The current SDK re-exports `RestError` and supports typed feature-flag values. I’m adding an environment-driven example with the unlabeled and `Production` variants kept explicit.

The implementation is complete; I’m installing the declared packages and compiling it locally to catch SDK type mismatches.

`1.12.2` is the repository’s unreleased version, not an npm release. I’m switching to the latest published SDK version and will recompile against that API.

Published `1.12.1` does not yet re-export `RestError`, so the correct stable import is its underlying `@azure/core-rest-pipeline` package. I’m applying that compatibility adjustment and the published feature-flag type requirement.

Implemented the TypeScript Azure App Configuration example in `src/index.ts`. It:

- Creates `AppConfigurationClient` from a connection string
- Creates unlabeled and `Production` settings
- Retrieves, lists, and deletes the unlabeled setting
- Creates the `BetaFeature` feature flag
- Handles Azure failures using `RestError`

Dependencies are declared in `package.json`:

```bash
npm install @azure/app-configuration @azure/core-rest-pipeline
```

Run it with:

```powershell
$env:AZURE_APP_CONFIGURATION_CONNECTION_STRING="<connection-string>"
npm run build
npm start
```

`@azure/core-rest-pipeline` is included because the published `@azure/app-configuration` version does not directly export `RestError`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- All required Azure and TypeScript dependencies are declared.
- A suitable tsconfig.json is included.
- The lockfile is consistent with the declared dependencies.

## Grader Results

- config-values.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/8)
      - `@azure/app-configuration` npm package: Pass
      - `AppConfigurationClient` constructor with connection string: Pass
      - `setConfigurationSetting()` with key, value, label: Pass
      - `getConfigurationSetting()` by key and label: Fail
      - `listConfigurationSettings()` with `ListConfigurationSettingOptions`: Fail
      - Feature flag settings with `featureFlagContentType`: Pass
      - `deleteConfigurationSetting()` and `RestError` handling: Pass
      - Async iteration with `for await...of` pattern: Pass
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct @azure/ Scoped Packages**: All Azure SDK dependencies use @azure/ scoped packages (e.g., @azure/storage-blob, @azure/keyvault-secrets, @azure/identity). No deprecated unscoped packages (azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-*).: Pass
  - @azure/identity for Authentication (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**@azure/identity for Authentication**: Includes @azure/identity as a dependency. For local development, DefaultAzureCredential is acceptable. For production code, uses a specific Entra token credential (e.g., ManagedIdentityCredential, ClientSecretCredential, WorkloadIdentityCredential) or configures DefaultAzureCredential with the AZURE_TOKEN_CREDENTIALS environment variable set to "prod" or a specific token credential. No hardcoded connection strings, account keys, SAS tokens, or secrets in code.: Fail
  - Client Constructor with Endpoint and Credential (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Constructor with Endpoint and Credential**: SDK clients constructed with endpoint URL and credential object, e.g. new BlobServiceClient(url, credential) or new SecretClient(vaultUrl, credential). No legacy constructor patterns or connection string-based construction when identity-based auth is available.: Fail
  - Async/Await Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Fail
  - Pagination with for-await-of (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Pass
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Pass
  - RestError Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**RestError Exception Handling**: Catches RestError (from @azure/core-rest-pipeline or the service package) with statusCode inspection for error-specific handling. Not just generic Error or try/catch without type checking.: Pass
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
| `Criteria from prompt file` | prompt_review | 75% | 1.00 | 0.7500 | 11.1% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.8% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 14.8% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 14.8% | ✅ |
| `RestError Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 14.8% | ✅ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 14.8% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 14.8% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 6.7500** | **61.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id app-configuration-dp-js-ts-crud --config js-ts-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
