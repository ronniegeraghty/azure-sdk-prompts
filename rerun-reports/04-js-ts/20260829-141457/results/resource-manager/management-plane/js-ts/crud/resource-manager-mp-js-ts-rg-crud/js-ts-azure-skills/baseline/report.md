# Evaluation Report: resource-manager-mp-js-ts-rg-crud

**Config:** js-ts-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 351.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-js-ts-rg-crud` |
| Config | js-ts-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 15/18 |
| Duration | 351.8s |
| Timestamp | 2026-08-29T08:47:25Z |
| Files Generated | 5 |
| Event Count | 3397 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 161.0s |
| Review | 190.1s |
| **Total** | **351.8s** |

## Configuration

- **name:** js-ts-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=242199 out=2612 |
| Turn Count | 11 |

## Tool Calls

### ✅ 1. glob — 7637ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 7232ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028","pattern":"**/package.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 6794ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028","pattern":"**/tsconfig.json"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. web_search (via github-mcp-server) — 14838ms

**Input:**

```json
{"query":"npm current versions @azure/identity @azure/arm-resources August 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"As of August 2026, here are the current npm versions for the Azure SDK packages you asked about:\n\n- **@azure/identity**: The latest version is **15.3.7**, last published on August 25, 2026【5:4†source】.\n- **@azure/arm-resources**: The latest version is **8.0.0**, last published about a month ago【5:0†source】【5:1†source】.\n\nYou can install or update them using the following commands:\n```sh\nnpm install @azure/identity@15.3.7\nnpm install @azure/arm-resources@8.0.0\n```\n\nLet me know if you need details for other Azure packages!","annotations":[{"text":"","start_index":188,"end_index":200,"url_citation":{"title":"versions on NPM | Release Alert","url":"https://releasealert.dev/npmjs/_/versions"}},{"text":"","start_index":295,"end_index":307,"url_citation":{"title":"@azure/arm-resources - npm","url":"https://www.npmjs.com/package/@azure/arm-resources"}},{"text":"","start_index":307,"end_index":319,"url_citation":{"title":"Azure ResourceManagement client library for JavaScript","url":"https://learn.microsoft.com/en-us/javascript/api/overview/azure/arm-resources-readme?view=azure-node-latest"}}]},"bing_searches":[{"text":"npm current versions @azure/identity and @azure/arm-resources August 2026","url":"https://www.bing.com/search?q=npm+current+versions+%40azure%2Fidentity+and+%40azure%2Farm-resources+August+2026"},{"text":"npm current version @azure/arm-resources August 2026","url":"https://www.bing.com/search?q=npm+current+version+%40azure%2Farm-resources+August+2026"}],"annotations":null}
```

### ✅ 5. apply_patch — 2108ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-resource-group-management-example\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"node dist/index.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/arm-resources\": \"^8.0.0\",\n+    \"@azure/identity\": \"^15.3.7\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"typescript\": \"^5.9.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"rootDir\": \"src\",\n+    \"strict\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import { ResourceManagementClient } from \"@azure/arm-resources\";\n+import { DefaultAzureCredential } from \"@azure/identity\";\n+\n+const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;\n+const resourceGroupName =\n+  process.env.AZURE_RESOURCE_GROUP_NAME ?? \"sdk-resource-group-example\";\n+\n+if (!subscriptionId) {\n+  throw new Error(\"Set the AZURE_SUBSCRIPTION_ID environment variable.\");\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const credential = new DefaultAzureCredential();\n+  const client = new ResourceManagementClient(credential, subscriptionId);\n+\n+  console.log(`Creating resource group \"${resourceGroupName}\"...`);\n+  const created = await client.resourceGroups.createOrUpdate(\n+    resourceGroupName,\n+    { location: \"eastus\" },\n+  );\n+  console.log(`Created: ${created.id}`);\n+\n+  console.log(\"\\nResource groups in the subscription:\");\n+  for await (const resourceGroup of client.resourceGroups.list()) {\n+    console.log(`- ${resourceGroup.name} (${resourceGroup.location})`);\n+  }\n+\n+  const details = await client.resourceGroups.get(resourceGroupName);\n+  console.log(\"\\nCreated resource group details:\", details);\n+\n+  const updated = await client.resourceGroups.update(resourceGroupName, {\n+    tags: {\n+      ...details.tags,\n+      managedBy: \"azure-sdk-typescript-example\",\n+    },\n+  });\n+  console.log(\"\\nUpdated tags:\", updated.tags);\n+\n+  console.log(`\\nDeleting resource group \"${resourceGroupName}\"...`);\n+  await client.resourceGroups.beginDeleteAndWait(resourceGroupName);\n+  console.log(\"Resource group deleted.\");\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Resource group operation failed:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Resource Group management with TypeScript\n+\n+This example uses the Azure management plane SDK to create, list, read, update,\n+and delete a resource group.\n+\n+## Install and build\n+\n+```powershell\n+npm install\n+npm run build\n+```\n+\n+Required runtime packages:\n+\n+```powershell\n+npm install @azure/identity @azure/arm-resources\n+```\n+\n+Authenticate with any method supported by `DefaultAzureCredential`, then set\n+the subscription ID. For example, in PowerShell:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:AZURE_RESOURCE_GROUP_NAME = \"sdk-resource-group-example\" # Optional\n+npm start\n+```\n+\n+Running the program creates and then permanently deletes the named resource\n+group. Ensure the selected identity has permission to manage resource groups in\n+the subscription.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json
@@ -1,0 +1,19 @@
+{
+  "name": "azure-resource-group-management-example",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "node dist/index.js"
+  },
+  "dependencies": {
+    "@azure/arm-resources": "^8.0.0",
+    "@azure/identity": "^15.3.7"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "typescript": "^5.9.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/tsconfig.json
@@ -1,0 +1,15 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "outDir": "dist",
+    "rootDir": "src",
+    "strict": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/src/index.ts
create file mode 100644
index 0000000..0
... (truncated)
```

### ✅ 6. powershell — 14712ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
npm error code ETARGET
npm error notarget No matching version found for @azure/identity@^15.3.7.
npm error notarget In most cases you or one of your dependencies are requesting a package version that doesn't exist.
npm error A complete log of this run can be found in: C:\Users\weidxu\AppData\Local\npm-cache\_logs\2026-08-29T08_48_29_728Z-debug-0.log
<shellId: 0 completed with exit code 1>
```

### ✅ 7. powershell — 17029ms

**Input:**

```json
{"command":"npm view @azure/identity version","description":"Check available Azure Identity version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
7.0.2
26.2.0
<shellId: 3 completed with exit code 0>
```

### ✅ 8. powershell — 15827ms

**Input:**

```json
{"command":"npm view @azure/arm-resources version","description":"Check available ARM Resources version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
4.13.2
<shellId: 1 completed with exit code 0>
```

### ✅ 9. powershell — 14675ms

**Input:**

```json
{"command":"npm view typescript version; npm view @types/node version","description":"Check available TypeScript package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
8.0.0
<shellId: 2 completed with exit code 0>
```

### ✅ 10. powershell — 16166ms

**Input:**

```json
{"command":"Write-Output \"identity=$(npm view @azure/identity version)\"; Write-Output \"arm-resources=$(npm view @azure/arm-resources version)\"; Write-Output \"typescript=$(npm view typescript version)\"; Write-Output \"node-types=$(npm view @types/node version)\"","description":"Confirm exact npm package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
identity=4.13.2
arm-resources=8.0.0
typescript=7.0.2
node-types=26.2.0
<shellId: 4 completed with exit code 0>
```

### ✅ 11. apply_patch — 2798ms

**Input:**

```json
"*** Begin Patch\n*** Update File: package.json\n@@\n-    \"@azure/identity\": \"^15.3.7\"\n+    \"@azure/identity\": \"^4.13.2\"\n@@\n-    \"@types/node\": \"^24.3.0\",\n-    \"typescript\": \"^5.9.2\"\n+    \"@types/node\": \"^26.2.0\",\n+    \"typescript\": \"^7.0.2\"\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/package.json
@@ -9,11 +9,11 @@
   },
   "dependencies": {
     "@azure/arm-resources": "^8.0.0",
-    "@azure/identity": "^15.3.7"
+    "@azure/identity": "^4.13.2"
   },
   "devDependencies": {
-    "@types/node": "^24.3.0",
-    "typescript": "^5.9.2"
+    "@types/node": "^26.2.0",
+    "typescript": "^7.0.2"
   }
 }
 

```

### ✅ 12. powershell — 20672ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install corrected dependencies and compile","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 50 packages in 17s

> azure-resource-group-management-example@1.0.0 build
> tsc

src/index.ts(14,59): error TS2769: No overload matches this call.
  The last overload gave the following error.
    Argument of type 'string | undefined' is not assignable to parameter of type 'string'.
      Type 'undefined' is not assignable to type 'string'.
<shellId: 5 completed with exit code 2>
```

### ✅ 13. apply_patch — 812ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n-async function main(): Promise\u003cvoid\u003e {\n+async function main(subscriptionId: string): Promise\u003cvoid\u003e {\n@@\n-main().catch((error: unknown) =\u003e {\n+main(subscriptionId).catch((error: unknown) =\u003e {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-baseline-1028316028/src/index.ts
@@ -9,7 +9,7 @@
   throw new Error("Set the AZURE_SUBSCRIPTION_ID environment variable.");
 }
 
-async function main(): Promise<void> {
+async function main(subscriptionId: string): Promise<void> {
   const credential = new DefaultAzureCredential();
   const client = new ResourceManagementClient(credential, subscriptionId);
 
@@ -41,7 +41,7 @@
   console.log("Resource group deleted.");
 }
 
-main().catch((error: unknown) => {
+main(subscriptionId).catch((error: unknown) => {
   console.error("Resource group operation failed:", error);
   process.exitCode = 1;
 });

```

### ✅ 14. powershell — 1793ms

**Input:**

```json
{"command":"npm run build","description":"Compile corrected TypeScript example","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-resource-group-management-example@1.0.0 build
> tsc
<shellId: 6 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll inspect the workspace, then add a focused TypeScript example and package setup without executing any Azure operations.

The example is in place. I’m compiling it now to catch SDK signature or TypeScript configuration issues; no Azure command or program execution will occur.

Implemented the complete TypeScript Azure Resource Group management example.

- `src/index.ts` uses `DefaultAzureCredential` and `ResourceManagementClient`
- Creates a resource group in `eastus`
- Lists resource groups with `for await...of`
- Gets and updates the created group with a tag
- Deletes it using `beginDeleteAndWait`
- `package.json` includes `@azure/identity` and `@azure/arm-resources`
- `README.md` documents installation, authentication, and execution

Run after configuring Azure credentials:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
npm start
```

This command performs real create, update, and delete operations against the configured subscription.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed. The TypeScript project supplies the required Azure SDK dependencies and compiler/type configuration.

### Strengths

- Valid package.json with both required Azure SDK packages
- Includes TypeScript and Node.js type definitions
- Includes a suitable strict tsconfig.json

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - `@azure/arm-resources` and `@azure/identity` npm packages: Pass
      - `DefaultAzureCredential` for authentication: Pass
      - `ResourceManagementClient` constructor with credential and subscriptionId: Pass
      - `resourceGroups.createOrUpdate()` with resource group name and parameters: Pass
      - `resourceGroups.list()` with async iteration (`for await...of`): Pass
      - `resourceGroups.get()` for fetching details: Pass
      - `resourceGroups.beginDeleteAndWait()` for long-running delete: Pass
      - Tag updates via `resourceGroups.update()` with tags parameter: Pass
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
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Fail
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
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 8.0000** | **72.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-js-ts-rg-crud --config js-ts-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
