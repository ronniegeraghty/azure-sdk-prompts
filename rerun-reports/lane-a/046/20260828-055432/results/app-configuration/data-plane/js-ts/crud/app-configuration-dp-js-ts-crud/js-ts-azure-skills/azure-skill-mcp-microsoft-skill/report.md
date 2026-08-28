# Evaluation Report: app-configuration-dp-js-ts-crud

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 672.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-js-ts-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 12/18 |
| Duration | 672.1s |
| Timestamp | 2026-08-27T21:54:32Z |
| Files Generated | 5 |
| Event Count | 4812 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 521.9s |
| Review | 149.0s |
| **Total** | **672.1s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-appconfiguration-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=568854 out=3850 |
| Turn Count | 14 |

## Tool Calls

### ✅ 1. skill — 4433ms

**Input:**

```json
{"skill":"azure-appconfiguration-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure App Configuration SDK for TypeScript

Centralized configuration management with feature flags and dynamic refresh.

## Installation

```bash
# Low-level CRUD SDK
npm install @azure/app-configuration @azure/identity

# High-level provider (recommended for apps)
npm install @azure/app-configuration-provider @azure/identity

# Feature flag management
npm install @microsoft/feature-management
```

## Environment Variables

```bash
AZURE_APPCONFIG_ENDPOINT=https://<your-resource>.azconfig.io
# OR
AZURE_APPCONFIG_CONNECTION_STRING=Endpoint=https://...;Id=...;Secret=...
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

```typescript
import { AppConfigurationClient } from "@azure/app-configuration";
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

const client = new AppConfigurationClient(
  process.env.AZURE_APPCONFIG_ENDPOINT!,
  credential
);

// Connection string
const client2 = new AppConfigurationClient(
  process.env.AZURE_APPCONFIG_CONNECTION_STRING!
);
```

## CRUD Operations

### Create/Update Settings

```typescript
// Add new (fails if exists)
await client.addConfigurationSetting({
  key: "app:settings:message",
  value: "Hello World",
  label: "production",
  contentType: "text/plain",
  tags: { environment: "prod" },
});

// Set (create or update)
await client.setConfigurationSetting({
  key: "app:settings:message",
  value: "Updated value",
  label: "production",
})
... (truncated)
```

### ❌ 2. view — 187567ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-appconfiguration-ts\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 187564ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-appconfiguration-ts\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. glob — 187560ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 187568ms

**Input:**

```json
{"intent":"Generate a local-only TypeScript Azure App Configuration CRUD sample using @azure/app-configuration, a connection string from an environment variable, feature flags, and RestError handling. Do not deploy or modify Azure resources during development.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. glob — 183806ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-appconfiguration-ts","pattern":"**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-typescript\skills\azure-appconfiguration-ts\SKILL.md
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 183809ms

**Input:**

```json
{"command":"get_best_practices","intent":"Get code-generation best practices for a TypeScript Azure App Configuration SDK CRUD sample that uses an environment-variable connection string and RestError handling, with no deployment.","parameters":{"language":"TypeScript","operation":"code-generation","resource":"Azure App Configuration"}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. powershell — 16837ms

**Input:**

```json
{"command":"npm view @azure/app-configuration version --json; npm view @azure/core-rest-pipeline version --json; npm view typescript version --json; npm view tsx version --json","description":"Check current npm package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
"1.12.1"
"1.25.0"
"7.0.2"
"4.23.12"
<shellId: 0 completed with exit code 0>
```

### ✅ 9. apply_patch — 2691ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-app-configuration-crud-sample\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"start\": \"tsx src/index.ts\",\n+    \"typecheck\": \"tsc --noEmit\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/app-configuration\": \"^1.12.1\",\n+    \"@azure/core-rest-pipeline\": \"^1.25.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"tsx\": \"^4.23.12\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"esModuleInterop\": true,\n+    \"skipLibCheck\": true,\n+    \"outDir\": \"dist\"\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: .env.example\n+AZURE_APPCONFIG_CONNECTION_STRING=Endpoint=https://your-store.azconfig.io;Id=your-id;Secret=your-secret\n*** Add File: src/index.ts\n+import {\n+  AppConfigurationClient,\n+  type ConfigurationSetting,\n+  type FeatureFlagValue,\n+  featureFlagContentType,\n+  featureFlagPrefix,\n+} from \"@azure/app-configuration\";\n+import { RestError } from \"@azure/core-rest-pipeline\";\n+\n+const key = \"app:Settings:FontSize\";\n+const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;\n+\n+if (!connectionString) {\n+  throw new Error(\n+    \"AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.\",\n+  );\n+}\n+\n+const client = new AppConfigurationClient(connectionString);\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  await client.setConfigurationSetting({ key, value: \"24\" });\n+\n+  await client.setConfigurationSetting({\n+    key,\n+    value: \"24\",\n+    label: \"Production\",\n+  });\n+\n+  const setting = await client.getConfigurationSetting({ key });\n+  console.log(`${setting.key} = ${setting.value}`);\n+\n+  console.log('Settings matching \"app:Settings:*\":');\n+  for await (const matchingSetting of client.listConfigurationSettings({\n+    keyFilter: \"app:Settings:*\",\n+  })) {\n+    const label = matchingSetting.label ?? \"(no label)\";\n+    console.log(\n+      `${matchingSetting.key} [${label}] = ${matchingSetting.value ?? \"\"}`,\n+    );\n+  }\n+\n+  const featureFlag: ConfigurationSetting\u003cFeatureFlagValue\u003e = {\n+    key: `${featureFlagPrefix}BetaFeature`,\n+    contentType: featureFlagContentType,\n+    value: {\n+      id: \"BetaFeature\",\n+      enabled: true,\n+      description: \"Enables the beta feature.\",\n+      conditions: {\n+        clientFilters: [],\n+      },\n+    },\n+  };\n+\n+  await client.setConfigurationSetting(featureFlag);\n+  await client.deleteConfigurationSetting({ key });\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  if (error instanceof RestError) {\n+    console.error(\n+      `Azure App Configuration request failed (${error.statusCode ?? \"unknown status\"}): ${error.message}`,\n+    );\n+    if (error.code) {\n+      console.error(`Error code: ${error.code}`);\n+    }\n+  } else if (error instanceof Error) {\n+    console.error(`Application error: ${error.message}`);\n+  } else {\n+    console.error(\"An unknown error occurred.\", error);\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure App Configuration CRUD sample\n+\n+Install the required Azure SDK package and the package that exports `RestError`:\n+\n+```powershell\n+npm install @azure/app-configuration @azure/core-rest-pipeline\n+```\n+\n+Install all project dependencies:\n+\n+```powershell\n+npm install\n+```\n+\n+Set the connection string without committing it:\n+\n+```powershell\n+$env:AZURE_APPCONFIG_CONNECTION_STRING = \"Endpoint=https://\u003cstore\u003e.azconfig.io;Id=\u003cid\u003e;Secret=\u003csecret\u003e\"\n+npm start\n+```\n+\n+The program creates an unlabeled setting and a `Production`-labeled setting,\n+reads and lists settings, creates the `BetaFeature` feature flag, and deletes\n+the unlabeled setting. The labeled setting and feature flag remain in the\n+store.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/package.json
@@ -1,0 +1,20 @@
+{
+  "name": "azure-app-configuration-crud-sample",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "start": "tsx src/index.ts",
+    "typecheck": "tsc --noEmit"
+  },
+  "dependencies": {
+    "@azure/app-configuration": "^1.12.1",
+    "@azure/core-rest-pipeline": "^1.25.0"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "tsx": "^4.23.12",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json
@@ -1,0 +1,15 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "exactOptionalPropertyTypes": true,
+    "esModuleInterop": true,
+    "skipLibCheck": true,
+    "outDir": "dist"
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-micros
... (truncated)
```

### ✅ 10. powershell — 21022ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run typecheck","description":"Install dependencies and type-check sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 26 packages in 17s

> azure-app-configuration-crud-sample@1.0.0 typecheck
> tsc --noEmit

src/index.ts(11,26): error TS2591: Cannot find name 'process'. Do you need to install type definitions for node? Try `npm i --save-dev @types/node` and then add 'node' to the types field in your tsconfig.
src/index.ts(43,9): error TS2322: Type '{ key: string; contentType: string; value: { id: string; enabled: true; description: string; conditions: { clientFilters: never[]; }; }; }' is not assignable to type 'ConfigurationSetting<FeatureFlagValue>'.
  Property 'isReadOnly' is missing in type '{ key: string; contentType: string; value: { id: string; enabled: true; description: string; conditions: { clientFilters: never[]; }; }; }' but required in type '{ isReadOnly: boolean; lastModified?: Date; }'.
src/index.ts(74,3): error TS2591: Cannot find name 'process'. Do you need to install type definitions for node? Try `npm i --save-dev @types/node` and then add 'node' to the types field in your tsconfig.
<shellId: 1 completed with exit code 1>
```

### ✅ 11. rg — 1954ms

**Input:**

```json
{"glob":"*.d.ts","head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\\node_modules\\@azure\\app-configuration","pattern":"ConfigurationSettingParam|FeatureFlagValue|ConfigurationSetting\u003c"}
```

**Output:**

```
Output too large to read at once (40.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787868154223-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\node_modules\@azure\app-configuration\dist\browser\appConfigurationClient.d.ts:import { type AddConfigurationSettingOptions, type AddConfigurationSettingParam, type AddConfigurationSettingResponse, type AppConfigurationClientOptions, type CheckConfigurationSettingsOptions, type ConfigurationSetting, type ConfigurationSettingId, type ConfigurationSnapshot, type Cr
```

### ✅ 12. view — 1614ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\\node_modules\\@azure\\app-configuration\\package.json"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/package.json
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/package.json
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/package.json
@@ -1,153 +1,153 @@
 {
   "name": "@azure/app-configuration",
   "author": "Microsoft Corporation",
   "description": "An isomorphic client library for the Azure App Configuration service.",
   "version": "1.12.1",
   "sdk-type": "client",
   "keywords": [
     "node",
     "azure",
     "typescript",
     "browser",
     "isomorphic",
     "cloud"
   ],
   "license": "MIT",
   "main": "./dist/commonjs/index.js",
   "types": "./dist/commonjs/index.d.ts",
   "homepage": "https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/appconfiguration/app-configuration/README.md",
   "repository": {
     "type": "git",
     "url": "git+https://github.com/Azure/azure-sdk-for-js",
     "directory": "sdk/appconfiguration/app-configuration"
   },
   "bugs": {
     "url": "https://github.com/Azure/azure-sdk-for-js/issues"
   },
   "files": [
     "dist/",
     "README.md",
     "LICENSE"
   ],
   "engines": {
     "node": ">=20.0.0"
   },
   "sideEffects": false,
   "autoPublish": false,
   "//metadata": {
     "constantPaths": [
       {
         "path": "src/internal/constants.ts",
         "prefix": "packageVersion"
       },
       {
         "path": "src/generated/src/appConfiguration.ts",
... (truncated)
```

### ✅ 13. rg — 920ms

**Input:**

```json
{"glob":"*.d.ts","head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\\node_modules\\@azure\\app-configuration\\dist\\esm","pattern":"export \\{[^\\n]*(SetConfigurationSettingParam|FeatureFlagValue)|type SetConfigurationSettingParam|interface SetConfigurationSettingParam|SetConfigurationSettingParam ="}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\node_modules\@azure\app-configuration\dist\esm\appConfigurationClient.d.ts:import { type AddConfigurationSettingOptions, type AddConfigurationSettingParam, type AddConfigurationSettingResponse, type AppConfigurationClientOptions, type CheckConfigurationSettingsOptions, type ConfigurationSetting, type ConfigurationSettingId, type ConfigurationSnapshot, type CreateSnapshotOptions, type CreateSnapshotResponse, type DeleteConfigurationSettingOptions, type DeleteConfigurationSettingResponse, type GetConfigurationSettingOptions, type GetConfigurationSettingResponse, type GetSnapshotOptions, type GetSnapshotResponse, type ListConfigurationSettingPage, type ListConfigurationSettingsForSnapshotOptions, type ListConfigurationSettingsOptions, type ListLabelsOptions, type ListLabelsPage, type ListRevisionsOptions, type ListRevisionsPage, type ListSnapshotsOptions, type ListSnapshotsPage, type PageSettings, type SetConfigurationSettingOptions, type SetConfigurationSettingParam, type SetConfigurationSettingResponse, type SetReadOnlyOptions, type SetReadOnlyResponse, type SettingLabel, type SnapshotInfo, type UpdateSnapshotOptions, type UpdateSnapshotResponse } from "./models.js";
C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\node_modules\@azure\app-configuration\dist\esm\models.d.ts:export type SetConfigurationSettingParam<T extends string | FeatureFlagValue | SecretReferenceValue | SnapshotReferenceValue = string> = ConfigurationSettingParam<T>;
C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\node_modules\@azure\app-configuration\dist\esm\index.d.ts:export { featureFlagContentType, featureFlagPrefix, type FeatureFlagValue, isFeatureFlag, parseFeatureFlag, } from "./featureF
... (truncated)
```

### ✅ 14. view — 653ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\\node_modules\\@azure\\app-configuration\\dist\\esm\\index.d.ts","view_range":[1,80]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/dist/esm/index.d.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/dist/esm/index.d.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/dist/esm/index.d.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/node_modules/@azure/app-configuration/dist/esm/index.d.ts
@@ -1,80 +1,80 @@
 export { AppConfigurationClient } from "./appConfigurationClient.js";
 export type { SimplePollerLike } from "./internal/lroShim.js";
 export { featureFlagContentType, featureFlagPrefix, type FeatureFlagValue, isFeatureFlag, parseFeatureFlag, } from "./featureFlag.js";
 export * from "./models.js";
 export { isSecretReference, parseSecretReference, secretReferenceContentType, type SecretReferenceValue, } from "./secretReference.js";
 export { isSnapshotReference, parseSnapshotReference, snapshotReferenceContentType, type SnapshotReferenceValue, } from "./snapshotReference.js";
 //# sourceMappingURL=index.d.ts.map

```

### ✅ 15. apply_patch — 814ms

**Input:**

```json
"*** Begin Patch\n*** Update File: tsconfig.json\n@@\n     \"module\": \"NodeNext\",\n     \"moduleResolution\": \"NodeNext\",\n+    \"types\": [\"node\"],\n     \"strict\": true,\n*** Update File: src/index.ts\n@@\n   AppConfigurationClient,\n-  type ConfigurationSetting,\n   type FeatureFlagValue,\n+  type SetConfigurationSettingParam,\n@@\n-  const featureFlag: ConfigurationSetting\u003cFeatureFlagValue\u003e = {\n+  const featureFlag: SetConfigurationSettingParam\u003cFeatureFlagValue\u003e = {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/tsconfig.json
@@ -3,6 +3,7 @@
     "target": "ES2022",
     "module": "NodeNext",
     "moduleResolution": "NodeNext",
+    "types": ["node"],
     "strict": true,
     "noUncheckedIndexedAccess": true,
     "exactOptionalPropertyTypes": true,


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts
@@ -1,7 +1,7 @@
 import {
   AppConfigurationClient,
-  type ConfigurationSetting,
   type FeatureFlagValue,
+  type SetConfigurationSettingParam,
   featureFlagContentType,
   featureFlagPrefix,
 } from "@azure/app-configuration";
@@ -40,7 +40,7 @@
     );
   }
 
-  const featureFlag: ConfigurationSetting<FeatureFlagValue> = {
+  const featureFlag: SetConfigurationSettingParam<FeatureFlagValue> = {
     key: `${featureFlagPrefix}BetaFeature`,
     contentType:
... (truncated)
```

### ✅ 16. powershell — 5568ms

**Input:**

```json
{"command":"npm run typecheck; if ($LASTEXITCODE -eq 0) { $env:AZURE_APPCONFIG_CONNECTION_STRING=''; npm start; if ($LASTEXITCODE -eq 0) { exit 1 } else { exit 0 } }","description":"Type-check and test missing-config error path","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-app-configuration-crud-sample@1.0.0 typecheck
> tsc --noEmit


> azure-app-configuration-crud-sample@1.0.0 start
> tsx src/index.ts

C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\src\index.ts:14
  throw new Error(
        ^

Error: AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.
    at <anonymous> (C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883\src\index.ts:14:9)
    at ModuleJob.run (node:internal/modules/esm/module_job:439:25)
    at async node:internal/modules/esm/loader:643:26
    at async asyncRunEntryPointWithESMLoader (node:internal/modules/run_main:101:5)

Node.js v24.18.0
<shellId: 2 completed with exit code 0>
```

### ✅ 17. apply_patch — 779ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n const key = \"app:Settings:FontSize\";\n-const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;\n-\n-if (!connectionString) {\n-  throw new Error(\n-    \"AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.\",\n-  );\n-}\n-\n-const client = new AppConfigurationClient(connectionString);\n \n async function main(): Promise\u003cvoid\u003e {\n+  const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;\n+  if (!connectionString) {\n+    throw new Error(\n+      \"AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.\",\n+    );\n+  }\n+\n+  const client = new AppConfigurationClient(connectionString);\n+\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-js-ts-crud-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-4282119883/src/index.ts
@@ -8,17 +8,17 @@
 import { RestError } from "@azure/core-rest-pipeline";
 
 const key = "app:Settings:FontSize";
-const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;
 
-if (!connectionString) {
-  throw new Error(
-    "AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.",
-  );
-}
+async function main(): Promise<void> {
+  const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;
+  if (!connectionString) {
+    throw new Error(
+      "AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.",
+    );
+  }
 
-const client = new AppConfigurationClient(connectionString);
+  const client = new AppConfigurationClient(connectionString);
 
-async function main(): Promise<void> {
   await client.setConfigurationSetting({ key, value: "24" });
 
   await client.setConfigurationSetting({

```

### ✅ 18. powershell — 2836ms

**Input:**

```json
{"command":"npm run typecheck; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; $env:AZURE_APPCONFIG_CONNECTION_STRING=''; $output = npm start 2\u003e\u00261; $exitCode = $LASTEXITCODE; $output | Write-Output; if ($exitCode -eq 0 -or ($output -join \"`n\") -notmatch 'Application error: AZURE_APPCONFIG_CONNECTION_STRING') { exit 1 }","description":"Verify types and handled error path","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-app-configuration-crud-sample@1.0.0 typecheck
> tsc --noEmit


> azure-app-configuration-crud-sample@1.0.0 start
> tsx src/index.ts

Application error: AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll align the sample with the SDK’s required feature-flag shape and error model, then create a local-only TypeScript project using an environment-variable connection string.

The skill package has no separate reference files beyond `SKILL.md`, and the Azure best-practices service is unavailable. I’ll use the loaded SDK guidance and verify the current package exports locally before finalizing imports.

Type checking exposed two SDK typing details: Node globals must be explicitly enabled under TypeScript 7, and writable feature flags use the SDK’s input setting type rather than the returned setting type. I’m correcting both without changing behavior.

Created a complete TypeScript sample in `src/index.ts` with:

- Connection-string authentication via `AZURE_APPCONFIG_CONNECTION_STRING`
- Unlabeled and `Production`-labeled settings
- Get, list with `app:Settings:*`, and delete operations
- A correctly typed `BetaFeature` feature flag
- Explicit `RestError` handling
- `package.json`, `tsconfig.json`, `.env.example`, and usage instructions

Install and run:

```powershell
npm install
$env:AZURE_APPCONFIG_CONNECTION_STRING = "Endpoint=https://<store>.azconfig.io;Id=<id>;Secret=<secret>"
npm start
```

Required runtime packages are `@azure/app-configuration` and `@azure/core-rest-pipeline`, with the latter providing `RestError`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete dependency manifest with matching lockfile
- Appropriate TypeScript configuration and Node.js type definitions

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
hyoka run --prompt-id app-configuration-dp-js-ts-crud --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
