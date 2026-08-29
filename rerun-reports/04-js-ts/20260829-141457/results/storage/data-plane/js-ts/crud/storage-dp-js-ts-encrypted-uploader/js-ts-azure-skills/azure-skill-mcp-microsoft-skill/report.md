# Evaluation Report: storage-dp-js-ts-encrypted-uploader

**Config:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 365.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-js-ts-encrypted-uploader` |
| Config | js-ts-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 30/35 |
| Duration | 365.5s |
| Timestamp | 2026-08-29T09:56:51Z |
| Files Generated | 8 |
| Event Count | 11267 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 207.2s |
| Review | 157.0s |
| **Total** | **365.5s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, applicationinsights-web-ts, azure-ai-contentsafety-ts, azure-ai-document-intelligence-ts, azure-ai-projects-ts, azure-ai-translation-ts, azure-ai-voicelive-ts, azure-appconfiguration-ts, azure-cosmos-ts, azure-eventhub-ts, azure-identity-ts, azure-keyvault-keys-ts, azure-keyvault-secrets-ts, azure-microsoft-playwright-testing-ts, azure-monitor-opentelemetry-ts, azure-postgres-ts, azure-search-documents-ts, azure-servicebus-ts, azure-storage-blob-ts, azure-storage-file-share-ts, azure-storage-queue-ts, azure-web-pubsub-ts, frontend-ui-dark-ts, m365-agents-ts, react-flow-node-ts, zustand-store-ts, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage-blob-ts, azure-keyvault-keys-ts, azure-identity-ts |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=493905 out=8677 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 7357ms

**Input:**

```json
{"skill":"azure-storage-blob-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for TypeScript

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```bash
npm install @azure/identity

# For Visual Studio Code credential support
npm install @azure/identity-vscode
```

## Environment Variables

### Service Principal (Secret)

```bash
AZURE_TENANT_ID=<tenant-id>
AZURE_CLIENT_ID=<client-id>
AZURE_CLIENT_SECRET=<client-secret>
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

### Service Principal (Certificate)

```bash
AZURE_TENANT_ID=<tenant-id>
AZURE_CLIENT_ID=<client-id>
AZURE_CLIENT_CERTIFICATE_PATH=/path/to/cert.pem
AZURE_CLIENT_CERTIFICATE_PASSWORD=<optional-password>
```

### Workload Identity (Kubernetes)

```bash
AZURE_TENANT_ID=<tenant-id>
AZURE_CLIENT_ID=<client-id>
AZURE_FEDERATED_TOKEN_FILE=/var/run/secrets/tokens/azure-identity
```

## DefaultAzureCredential (Recommended for Local Development)

```typescript
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

// Use with any Azure SDK client
import { BlobServiceClient } from "@azure/storage-blob";
const blobClient = new BlobServiceClient(
  "https://<account>.blob.core.windows.net",
  credential
);
```

See [DefaultAzureCredential overview](https://aka.ms/azsdk/js/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

## Managed Identity

### System-Assigned

```typescript
import { ManagedIdentityCred
... (truncated)
```

### ✅ 2. skill — 6087ms

**Input:**

```json
{"skill":"azure-keyvault-keys-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Key Vault Keys SDK for TypeScript

Manage cryptographic keys with Azure Key Vault.

## Installation

```bash
# Keys SDK
npm install @azure/keyvault-keys @azure/identity
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
import { KeyClient, CryptographyClient } from "@azure/keyvault-keys";

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
for await (const secretProperties of secretCl
... (truncated)
```

### ✅ 3. skill — 4748ms

**Input:**

```json
{"skill":"azure-identity-ts"}
```

**Output:**

```
Skill loaded successfully ✅

# @azure/storage-blob (TypeScript/JavaScript)

SDK for Azure Blob Storage operations — upload, download, list, and manage blobs and containers.

## Installation

```bash
npm install @azure/storage-blob @azure/identity
```

**Current Version**: 12.x  
**Node.js**: >= 18.0.0

## Environment Variables

```bash
AZURE_STORAGE_ACCOUNT_NAME=<account-name>
AZURE_STORAGE_ACCOUNT_KEY=<account-key>
# OR connection string
AZURE_STORAGE_CONNECTION_STRING=DefaultEndpointsProtocol=https;AccountName=...
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication

### Microsoft Entra Token Credential (Recommended)

```typescript
import { BlobServiceClient } from "@azure/storage-blob";
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
const credential = new DefaultAzureCredential({requiredEnvVars: ["AZURE_TOKEN_CREDENTIALS"]});
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/javascript/api/overview/azure/identity-readme?view=azure-node-latest#credential-classes
// const credential = new ManagedIdentityCredential();

const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME!;
const client = new BlobServiceClient(
  `https://${accountName}.blob.core.windows.net`,
  credential
);
```

### Connection String

```typescript
import { BlobServiceClient } from "@azure/storage-blob";

const client = BlobServiceClient.fromConnectionString(
  process.env.AZURE_STORAGE_CONNECTION_STRING!
);
```

### StorageSharedKeyCredential (Node.js only)

```typescript
import { BlobServiceClient, StorageSharedKeyCredential } from "@azure/storage-blob";

const accountName = process.env.AZURE_STORAGE_ACCOUNT_NAME!;
const accountKey = process.env.AZURE_STORAGE_ACCOUNT_KEY!;

const sharedKeyCredential = new StorageShar
... (truncated)
```

### ✅ 4. glob — 7913ms

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-storage-blob-ts","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-keys-ts","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-identity-ts"],"pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 7128ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 7509ms

**Input:**

```json
{"command":"get_ai_application_best_practices","intent":"Create a local TypeScript Node.js application using managed identity, Azure Blob Storage, and Azure Key Vault Keys for envelope encryption without provisioning Azure resources."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 8299ms

**Input:**

```json
{"command":"search","intent":"Find official TypeScript guidance for Azure Key Vault CryptographyClient wrapKey/unwrapKey and Blob Storage metadata with ManagedIdentityCredential.","parameters":{"query":"TypeScript Key Vault CryptographyClient wrapKey unwrapKey Blob Storage metadata ManagedIdentityCredential"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 8. view — 5231ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-keyvault-keys-ts\\references\\keys.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-storage-blob-ts/references/streaming.md
@@ -1,426 +1,426 @@
 # @azure/storage-blob - Streaming Patterns
 
 Reference documentation for upload/download streaming in the Azure Blob Storage TypeScript SDK.
 
 **Source**: [Azure SDK for JS - storage-blob](https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/storage/storage-blob)
 
 ---
 
 ## Installation
 
 ```bash
 npm install @azure/storage-blob @azure/identity
 ```
 
 ---
 
 ## Client Setup
 
 ```typescript
 import {
   BlobServiceClient,
   ContainerClient,
   BlockBlobClient,
 } from "@azure/storage-blob";
 import { DefaultAzureCredential } from "@azure/identity";
 
 const credential = new DefaultAzureCredential();
 const accountUrl = `https://${process.env["STORAGE_ACCOUNT_NAME"]}.blob.core.windows.net`;
 
 const blobServiceClient = new BlobServiceClient(accountUrl, credential);
 const containerClient = blobServiceClient.getContainerClient("my-container");
 const blobClient = containerClient.getBlockBlobClient("my-blob.txt");
 ```
 
 ---
 
 ## Download Streaming
 
 ### Download to Buffer
 
 ```typescript
 const downloadResponse = await blobClient.download(0);
 const downloaded = await streamToBuffer(downloadResponse.readableStreamBody!);
 
 async function streamToBuffer(readableStream: NodeJS.ReadableStream): Promise<Buffer> {
   return new Promise((resolve, reject) => {
     const chunks: Buffer[] = [];

... (truncated)
```

### ✅ 9. view — 4833ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-storage-blob-ts\\references\\streaming.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-identity-ts/references/credential-types.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-identity-ts/references/credential-types.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-identity-ts/references/credential-types.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-identity-ts/references/credential-types.md
@@ -1,385 +1,385 @@
 # Credential Types Reference
 
 Azure Identity credential types for authenticating to Azure services using the @azure/identity TypeScript SDK.
 
 ## Overview
 
 The Azure Identity library provides various credential classes for different authentication scenarios. Choose the right credential based on your environment and security requirements.
 
 ## Credential Selection Guide
 
 | Scenario | Recommended Credential |
 |----------|------------------------|
 | Production (any environment) | `DefaultAzureCredential` |
 | Azure VM/App Service | `ManagedIdentityCredential` |
 | Service Principal (secret) | `ClientSecretCredential` |
 | Service Principal (cert) | `ClientCertificateCredential` |
 | Local development | `AzureCliCredential` or `AzureDeveloperCliCredential` |
 | Browser application | `InteractiveBrowserCredential` |
 | CI/CD pipeline | `ClientSecretCredential` or `WorkloadIdentityCredential` |
 | Kubernetes (AKS) | `WorkloadIdentityCredential` |
 
 ## DefaultAzureCredential (Recommended)
 
 The most versatile credential - automatically tries multiple authentication methods.
 
 ```typescript
 import { DefaultAzureCredential } from "@azure/identity";
 
 const credential = new DefaultAzureCredential();
 
 // Works in all environments - dev and production
 import { BlobServiceClient } fr
... (truncated)
```

### ✅ 10. view — 4449ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-typescript\\skills\\azure-identity-ts\\references\\credential-types.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-keys-ts/references/keys.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-keys-ts/references/keys.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-keys-ts/references/keys.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-typescript/skills/azure-keyvault-keys-ts/references/keys.md
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
 const rsaKey
... (truncated)
```

### ✅ 11. azure-get_azure_bestpractices (via azure) — 4064ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local TypeScript Node.js sample using Azure Blob Storage, Azure Key Vault Keys cryptography, and managed identity; do not provision resources.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 12. azure-documentation (via azure) — 6040ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Microsoft guidance for TypeScript Key Vault CryptographyClient wrapKey/unwrapKey and Blob Storage metadata.","parameters":{"query":"Azure SDK JavaScript TypeScript Key Vault CryptographyClient wrapKey unwrapKey Blob Storage metadata ManagedIdentityCredential"}}
```

**Output:**

```
{"results":[{"title":"CryptographyClient class-Method Details","content":"# CryptographyClient class\r\n\r\n- Package:\r\n    - [@azure/keyvault-keys](https://learn.microsoft.com/en-us/javascript/api/preview-docs/@azure/keyvault-keys/)\r\n\r\nA client used to perform cryptographic operations on an Azure Key vault key or a local \u0026lt;xref:JsonWebKey\u0026gt;.\r\n\r\n## Method Details\r\n\r\n### unwrapKey(KeyWrapAlgorithm, Uint8Array, UnwrapKeyOptions)\r\n\r\nUnwraps the given wrapped key using the specified cryptography algorithm Example usage:\r\n\r\n\u0060\u0060\u0060ts\r\nlet client = new CryptographyClient(keyVaultKey, credentials);\r\nlet result = await client.unwrapKey(\u0022RSA1_5\u0022, keyToUnwrap);\r\n\u0060\u0060\u0060\r\n\r\n\u0060\u0060\u0060typescript\r\nfunction unwrapKey(algorithm: KeyWrapAlgorithm, encryptedKey: Uint8Array, options?: UnwrapKeyOptions)\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n- algorithm\r\n    - [KeyWrapAlgorithm](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/keywrapalgorithm)\r\n\r\nThe decryption algorithm to use to unwrap the key.\r\n\r\n- encryptedKey\r\n    - Uint8Array\r\n\r\nThe encrypted key to unwrap.\r\n\r\n- options\r\n    - [UnwrapKeyOptions](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/unwrapkeyoptions)\r\n\r\nAdditional options.\r\n\r\n#### Returns\r\n\r\nPromise\u0026lt;[UnwrapResult](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-keys/unwrapresult)\u0026gt;","contentUrl":"https://learn.microsoft.com/javascript/api/preview-docs/@azure/keyvault-keys/cryptographyclient?view=az-js-keyvault-v4"},{"title":"Get started with Azure Key Vault keys in JavaScript","content":"# Get started with Azure Key Vault keys in JavaScript\n## Create a KeyClient object\nOnce your local environment and Key Vault authorization are set up, create a JavaScript file, which includes the [@azure/identity](https://www.npmjs.com/package/@azure/identity) and the [@azure/
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 5631ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official TypeScript samples for Key Vault key wrapping and Blob Storage uploads with metadata.","parameters":{"language":"typescript","query":"CryptographyClient wrapKey unwrapKey BlockBlobClient uploadData metadata ManagedIdentityCredential"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates unwrapping (decrypting) a wrapped key using CryptographyClient\u0027s unwrapKey method with RSA-OAEP algorithm.\npackage: @azure/identity, @azure/keyvault-keys\nlanguage: typescript\n","codeSnippet":"import { DefaultAzureCredential } from \u0022@azure/identity\u0022;\r\nimport { KeyClient, CryptographyClient } from \u0022@azure/keyvault-keys\u0022;\r\n\r\nconst credential = new DefaultAzureCredential();\r\n\r\nconst vaultName = \u0022\u003CYOUR KEYVAULT NAME\u003E\u0022;\r\nconst url = \u0060https://${vaultName}.vault.azure.net\u0060;\r\n\r\nconst client = new KeyClient(url, credential);\r\n\r\nconst myKey = await client.createKey(\u0022MyKey\u0022, \u0022RSA\u0022);\r\nconst cryptographyClient = new CryptographyClient(myKey, credential);\r\n\r\nconst wrapResult = await cryptographyClient.wrapKey(\u0022RSA-OAEP\u0022, Buffer.from(\u0022My Key\u0022));\r\nconsole.log(\u0022wrap result:\u0022, wrapResult.result);\r\n\r\nconst unwrapResult = await cryptographyClient.unwrapKey(\u0022RSA-OAEP\u0022, wrapResult.result);\r\nconsole.log(\u0022unwrap result: \u0022, unwrapResult.result);","link":"https://learn.microsoft.com/javascript/api/overview/azure/keyvault-keys-readme?view=azure-node-latest#cryptography","language":"typescript"},{"description":"description: Creates a ManagedIdentityCredential instance with customized retry options, specifying maximum retries and delay intervals for token acquisition attempts using the Azure Identity library for JavaScript in TypeScript.\npackage: @azure/identity\nlanguage: typescript\n","codeSnippet":"import { ManagedIdentityCredential } from \u0022@azure/identity\u0022;\r\n\r\nconst credential = new ManagedIdentityCredential(\r\n  process.env.AZURE_CLIENT_ID as string, // For user-assigned managed identity\r\n  {\r\n    retryOptions: {\r\n      maxRetries: 3,           // Maximum number of retry attempts\r\n      retryDelayInMs: 500,     // Initial delay between retries (in millise
... (truncated)
```

### ✅ 14. powershell — 19680ms

**Input:**

```json
{"command":"npm view @azure/identity version --json; npm view @azure/keyvault-keys version --json; npm view @azure/storage-blob version --json; npm view @azure/core-auth version --json; npm view typescript version --json; npm view @types/node version --json","description":"Check current package versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
"4.13.2"
"4.10.2"
"12.33.0"
"1.11.0"
"7.0.2"
"26.2.0"
<shellId: 0 completed with exit code 0>
```

### ✅ 15. apply_patch — 816ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-encrypted-blob-uploader\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"description\": \"Client-side AES-GCM envelope encryption for Azure Blob Storage using Azure Key Vault Keys.\",\n+  \"type\": \"module\",\n+  \"engines\": {\n+    \"node\": \"\u003e=20.0.0\"\n+  },\n+  \"scripts\": {\n+    \"build\": \"tsc -p tsconfig.json\",\n+    \"start\": \"node dist/main.js\",\n+    \"demo\": \"npm run build \u0026\u0026 npm start\",\n+    \"typecheck\": \"tsc -p tsconfig.json --noEmit\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/core-auth\": \"^1.11.0\",\n+    \"@azure/identity\": \"^4.13.2\",\n+    \"@azure/keyvault-keys\": \"^4.10.2\",\n+    \"@azure/storage-blob\": \"^12.33.0\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true,\n+    \"sourceMap\": true\n+  },\n+  \"include\": [\n+    \"src/**/*.ts\"\n+  ]\n+}\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+.env\n+*.log\n*** Add File: .env.example\n+AZURE_STORAGE_BLOB_ENDPOINT=https://\u003cstorage-account\u003e.blob.core.windows.net\n+AZURE_KEY_VAULT_ENDPOINT=https://\u003cvault-name\u003e.vault.azure.net\n+AZURE_KEY_VAULT_KEY_NAME=\u003crsa-key-name\u003e\n+AZURE_STORAGE_CONTAINER=\u003ccontainer-name\u003e\n+AZURE_STORAGE_BLOB_NAME=encrypted-demo.txt\n+\n+# Optional: the client ID of a user-assigned managed identity.\n+AZURE_CLIENT_ID=\n*** Add File: src/config.ts\n+import type { TokenCredential } from \"@azure/core-auth\";\n+import { ManagedIdentityCredential } from \"@azure/identity\";\n+import { KeyClient } from \"@azure/keyvault-keys\";\n+import {\n+  BlobServiceClient,\n+  type ContainerClient,\n+} from \"@azure/storage-blob\";\n+\n+export interface AppConfig {\n+  readonly credential: TokenCredential;\n+  readonly keyClient: KeyClient;\n+  readonly containerClient: ContainerClient;\n+  readonly keyName: string;\n+  readonly blobName: string;\n+}\n+\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name]?.trim();\n+  if (!value) {\n+    throw new Error(`Required environment variable ${name} is not set.`);\n+  }\n+\n+  return value;\n+}\n+\n+function requireHttpsEndpoint(name: string): string {\n+  const value = requireEnvironmentVariable(name);\n+  let endpoint: URL;\n+\n+  try {\n+    endpoint = new URL(value);\n+  } catch (error) {\n+    throw new Error(`${name} must be a valid URL.`, { cause: error });\n+  }\n+\n+  if (endpoint.protocol !== \"https:\") {\n+    throw new Error(`${name} must use HTTPS.`);\n+  }\n+\n+  return endpoint.toString().replace(/\\/$/, \"\");\n+}\n+\n+export function createAppConfig(): AppConfig {\n+  const storageEndpoint = requireHttpsEndpoint(\n+    \"AZURE_STORAGE_BLOB_ENDPOINT\",\n+  );\n+  const vaultEndpoint = requireHttpsEndpoint(\"AZURE_KEY_VAULT_ENDPOINT\");\n+  const keyName = requireEnvironmentVariable(\"AZURE_KEY_VAULT_KEY_NAME\");\n+  const containerName = requireEnvironmentVariable(\"AZURE_STORAGE_CONTAINER\");\n+  const blobName =\n+    process.env.AZURE_STORAGE_BLOB_NAME?.trim() || \"encrypted-demo.txt\";\n+  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();\n+\n+  const credential: TokenCredential = managedIdentityClientId\n+    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })\n+    : new ManagedIdentityCredential();\n+\n+  const blobServiceClient = new BlobServiceClient(\n+    storageEndpoint,\n+    credential,\n+  );\n+\n+  return {\n+    credential,\n+    keyClient: new KeyClient(vaultEndpoint, credential),\n+    containerClient: blobServiceClient.getContainerClient(containerName),\n+    keyName,\n+    blobName,\n+  };\n+}\n*** Add File: src/errors.ts\n+export type OperationCategory = \"storage\" | \"key-vault\" | \"cryptography\";\n+\n+export class EncryptedBlobError extends Error {\n+  public constructor(\n+    public readonly category: OperationCategory,\n+    public readonly operation: string,\n+    message: string,\n+    options?: ErrorOptions,\n+  ) {\n+    super(message, options);\n+    this.name = \"EncryptedBlobError\";\n+  }\n+}\n+\n+interface AzureErrorShape {\n+  readonly code?: string;\n+  readonly statusCode?: number;\n+}\n+\n+function readAzureErrorShape(error: unknown): AzureErrorShape {\n+  if (typeof error !== \"object\" || error === null) {\n+    return {};\n+  }\n+\n+  const candidate = error as Record\u003cstring, unknown\u003e;\n+  return {\n+    ...(typeof candidate.code === \"string\" ? { code: candidate.code } : {}),\n+    ...(typeof candidate.statusCode === \"number\"\n+      ? { statusCode: candidate.statusCode }\n+      : {}),\n+  };\n+}\n+\n+export function describeAzureFailure(error: unknown): string {\n+  const { code, statusCode } = readAzureErrorShape(error);\n+  const details = [\n+    code ? `code ${code}` : undefined,\n+    statusCode ? `HTTP ${statusCode}` : undefined,\n+  ].filter((item): item is string =\u003e item !== undefined);\n+\n+  if (error instanceof Error) {\n+    return details.length \u003e 0\n+      ? `${error.message} (${details.join(\", \")})`\n+      : error.message;\n+  }\n+\n+  return details.length \u003e 0 ? details.join(\", \") : \"Unknown Azure SDK error\";\n+}\n*** Add File: src/keyManagement.ts\n+import { randomBytes } from \"node:crypto\";\n+import type { TokenCredential } from \"@azure/core-auth\";\n+import {\n+  CryptographyClient,\n+  type KeyClient,\n+} from \"@azure/keyvault-keys\";\n+import {\n+  describeAzureFailure,\n+  EncryptedBlobError,\n+} from \"./errors.js\";\n+\n+const DATA_KEY_LENGTH_BYTES = 32;\n+export const KEY_WRAP_ALGORITHM = \"RSA-OAEP-256\" as const;\n+\n+export interface ProtectedDataKey {\n+  readonly keyId: string;\n+  readonly wrappedKey: Buffer;\n+  readonly wrapAlgorithm: typeof KEY_WRAP_ALGORITHM;\n+}\n+\n+export class KeyVaultKeyManager {\n+  public constructor(\n+    private readonly keyClient: KeyClient,\n+    private readonly credential: TokenCredential,\n+    private readonly keyName: string,\n+  ) {}\n+\n+  public generateDataKey(): Buffer {\n+    return randomBytes(DATA_KEY_LENGTH_BYTES);\n+  }\n+\n+  public async protectDataKey(dataKey: Buffer): Promise\u003cProtectedDataKey\u003e {\n+    if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"wrap data key\",\n+        `The data encryption key must be ${DATA_KEY_LENGTH_BYTES} bytes.`,\n+      );\n+    }\n+\n+    try {\n+      const key = await this.keyClient.getKey(this.keyName);\n+      if (!key.id) {\n+        throw new Error(\"Key Vault returned a key without a versioned key ID.\");\n+      }\n+\n+      const cryptographyClient = new CryptographyClient(\n+        key.id,\n+        this.credential,\n+      );\n+      const result = await cryptographyClient.wrapKey(\n+        KEY_WRAP_ALGORITHM,\n+        dataKey,\n+      );\n+\n+      return {\n+        keyId: key.id,\n+        wrappedKey: Buffer.from(result.result),\n+        wrapAlgorithm: KEY_WRAP_ALGORITHM,\n+      };\n+    } catch (error) {\n+      throw new EncryptedBlobError(\n+        \"key-vault\",\n+        \"wrap data key\",\n+        `Azure Key Vault could not protect the data key: ${describeAzureFailure(error)}`,\n+        { cause: error },\n+      );\n+    }\n+  }\n+\n+  public async recoverDataKey(\n+    keyId: string,\n+    wrappedKey: Buffer,\n+    wrapAlgorithm: string,\n+  ): Promise\u003cBuffer\u003e {\n+    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"unwrap data key\",\n+        `Unsupported key wrap algorithm: ${wrapAlgorithm}.`,\n+      );\n+    }\n+\n+    this.validateKeyId(keyId);\n+\n+    try {\n+      const cryptographyClient = new CryptographyClient(\n+        keyId,\n+        this.credential,\n+      );\n+      const result = await cryptographyClient.unwrapKey(\n+        KEY_WRAP_ALGORITHM,\n+        wrappedKey,\n+      );\n+      const dataKey = Buffer.from(result.result);\n+\n+      if (dataKey.length !== DATA_KEY_LENGTH_BYTES) {\n+        dataKey.fill(0);\n+        throw new Error(\n+          `Key Vault returned an invalid ${dataKey.length}-byte data key.`,\n+        );\n+      }\n+\n+      return dataKey;\n+    } catch (error) {\n+      throw new EncryptedBlobError(\n+        \"key-vault\",\n+        \"unwrap data key\",\n+        `Azure Key Vault could not recover the data key: ${describeAzureFailure(error)}`,\n+        { cause: error },\n+      );\n+    }\n+  }\n+\n+  private validateKeyId(keyId: string): void {\n+    let parsed: URL;\n+\n+    try {\n+      parsed = new URL(keyId);\n+    } catch (error) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"validate key ID\",\n+        \"Blob metadata contains an invalid Key Vault key ID.\",\n+        { cause: error },\n+      );\n+    }\n+\n+    if (\n+      parsed.protocol !== \"https:\" ||\n+      !/^\\/keys\\/[^/]+\\/[^/]+$/.test(parsed.pathname)\n+    ) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"validate key ID\",\n+        \"Blob metadata must reference a versioned HTTPS Key Vault key ID.\",\n+      );\n+    }\n+  }\n+}\n*** Add File: src/encryptedBlobStorage.ts\n+import {\n+  createCipheriv,\n+  createDecipheriv,\n+  randomBytes,\n+} from \"node:crypto\";\n+import { readFile, writeFile } from \"node:fs/promises\";\n+import type { ContainerClient } from \"@azure/storage-blob\";\n+import {\n+  describeAzureFailure,\n+  EncryptedBlobError,\n+} from \"./errors.js\";\n+import {\n+  KEY_WRAP_ALGORITHM,\n+  type KeyVaultKeyManager,\n+} from \"./keyManagement.js\";\n+\n+const ENCRYPTION_ALGORITHM = \"aes-256-gcm\";\n+const ENCRYPTION_VERSION = \"1\";\n+const IV_LENGTH_BYTES = 12;\n+const AUTH_TAG_LENGTH_BYTES = 16;\n+\n+const METADATA = {\n+  version: \"encryptionversion\",\n+  algorithm: \"encryptionalgorithm\",\n+  keyId: \"keyid\",\n+  wrapAlgorithm: \"keywrapalgorithm\",\n+  wrappedKey: \"wrappeddek\",\n+  iv: \"iv\",\n+  authenticationTag: \"authenticationtag\",\n+} as const;\n+\n+export interface UploadResult {\n+  readonly keyId: string;\n+  readonly wrappedDataKeyBase64: string;\n+  readonly blobUrl: string;\n+}\n+\n+interface EncryptionMetadata {\n+  readonly keyId: string;\n+  readonly wrappedKey: Buffer;\n+  readonly iv: Buffer;\n+  readonly authenticationTag: Buffer;\n+}\n+\n+export class EncryptedBlobStorage {\n+  public constructor(\n+    private readonly containerClient: ContainerClient,\n+    private readonly keyManager: KeyVaultKeyManager,\n+  ) {}\n+\n+  public async upload(\n+    blobName: string,\n+    plaintext: Buffer | string,\n+  ): Promise\u003cUploadResult\u003e {\n+    const data = typeof plaintext === \"string\"\n+      ? Buffer.from(plaintext, \"utf8\")\n+      : plaintext;\n+    const dataKey = this.keyManager.generateDataKey();\n+\n+    try {\n+      const iv = randomBytes(IV_LENGTH_BYTES);\n+      const cipher = createCipheriv(ENCRYPTION_ALGORITHM, dataKey, iv, {\n+        authTagLength: AUTH_TAG_LENGTH_BYTES,\n+      });\n+      const ciphertext = Buffer.concat([cipher.update(data), cipher.final()]);\n+      const authenticationTag = cipher.getAuthTag();\n+      const protectedKey = await this.keyManager.protectDataKey(dataKey);\n+      const wrappedDataKeyBase64 = protectedKey.wrappedKey.toString(\"base64\");\n+      const blockBlobClient =\n+        this.containerClient.getBlockBlobClient(blobName);\n+\n+      try {\n+        await blockBlobClient.uploadData(ciphertext, {\n+          blobHTTPHeaders: {\n+            blobContentType: \"application/octet-stream\",\n+          },\n+          metadata: {\n+            [METADATA.version]: ENCRYPTION_VERSION,\n+            [METADATA.algorithm]: ENCRYPTION_ALGORITHM,\n+            [METADATA.keyId]: protectedKey.keyId,\n+            [METADATA.wrapAlgorithm]: protectedKey.wrapAlgorithm,\n+            [METADATA.wrappedKey]: wrappedDataKeyBase64,\n+            [METADATA.iv]: iv.toString(\"base64\"),\n+            [METADATA.authenticationTag]:\n+              authenticationTag.toString(\"base64\"),\n+          },\n+        });\n+      } catch (error) {\n+        throw new EncryptedBlobError(\n+          \"storage\",\n+          \"upload encrypted blob\",\n+          `Azure Blob Storage upload failed: ${describeAzureFailure(error)}`,\n+          { cause: error },\n+        );\n+      }\n+\n+      return {\n+        keyId: protectedKey.keyId,\n+        wrappedDataKeyBase64,\n+        blobUrl: blockBlobClient.url,\n+      };\n+    } catch (error) {\n+      if (error instanceof EncryptedBlobError) {\n+        throw error;\n+      }\n+\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"encrypt blob\",\n+        error instanceof Error ? error.message : \"Local encryption failed.\",\n+        { cause: error },\n+      );\n+    } finally {\n+      dataKey.fill(0);\n+    }\n+  }\n+\n+  public async uploadFile(\n+    blobName: string,\n+    localPath: string,\n+  ): Promise\u003cUploadResult\u003e {\n+    return this.upload(blobName, await readFile(localPath));\n+  }\n+\n+  public async download(blobName: string): Promise\u003cBuffer\u003e {\n+    const blockBlobClient = this.containerClient.getBlockBlobClient(blobName);\n+    let ciphertext: Buffer;\n+    let metadata: Record\u003cstring, string\u003e;\n+\n+    try {\n+      const response = await blockBlobClient.download();\n+      if (!response.readableStreamBody) {\n+        throw new Error(\"Blob download returned no readable response body.\");\n+      }\n+\n+      ciphertext = await this.streamToBuffer(response.readableStreamBody);\n+      metadata = response.metadata;\n+    } catch (error) {\n+      throw new EncryptedBlobError(\n+        \"storage\",\n+        \"download encrypted blob\",\n+        `Azure Blob Storage download failed: ${describeAzureFailure(error)}`,\n+        { cause: error },\n+      );\n+    }\n+\n+    const encryptionMetadata = this.parseMetadata(metadata);\n+    const dataKey = await this.keyManager.recoverDataKey(\n+      encryptionMetadata.keyId,\n+      encryptionMetadata.wrappedKey,\n+      KEY_WRAP_ALGORITHM,\n+    );\n+\n+    try {\n+      const decipher = createDecipheriv(\n+        ENCRYPTION_ALGORITHM,\n+        dataKey,\n+        encryptionMetadata.iv,\n+        { authTagLength: AUTH_TAG_LENGTH_BYTES },\n+      );\n+      decipher.setAuthTag(encryptionMetadata.authenticationTag);\n+      return Buffer.concat([\n+        decipher.update(ciphertext),\n+        decipher.final(),\n+      ]);\n+    } catch (error) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"decrypt blob\",\n+        \"Local AES-GCM decryption failed. The ciphertext or encryption metadata may have been altered.\",\n+        { cause: error },\n+      );\n+    } finally {\n+      dataKey.fill(0);\n+    }\n+  }\n+\n+  public async downloadToFile(\n+    blobName: string,\n+    localPath: string,\n+  ): Promise\u003cvoid\u003e {\n+    await writeFile(localPath, await this.download(blobName));\n+  }\n+\n+  private parseMetadata(\n+    metadata: Record\u003cstring, string\u003e,\n+  ): EncryptionMetadata {\n+    const required = (name: string): string =\u003e {\n+      const value = metadata[name];\n+      if (!value) {\n+        throw new EncryptedBlobError(\n+          \"cryptography\",\n+          \"read encryption metadata\",\n+          `Blob metadata is missing ${name}.`,\n+        );\n+      }\n+\n+      return value;\n+    };\n+\n+    const version = required(METADATA.version);\n+    const algorithm = required(METADATA.algorithm);\n+    const wrapAlgorithm = required(METADATA.wrapAlgorithm);\n+\n+    if (version !== ENCRYPTION_VERSION) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        `Unsupported encryption metadata version: ${version}.`,\n+      );\n+    }\n+\n+    if (algorithm !== ENCRYPTION_ALGORITHM) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        `Unsupported content encryption algorithm: ${algorithm}.`,\n+      );\n+    }\n+\n+    if (wrapAlgorithm !== KEY_WRAP_ALGORITHM) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        `Unsupported key wrap algorithm: ${wrapAlgorithm}.`,\n+      );\n+    }\n+\n+    const wrappedKey = this.decodeBase64(\n+      required(METADATA.wrappedKey),\n+      METADATA.wrappedKey,\n+    );\n+    const iv = this.decodeBase64(required(METADATA.iv), METADATA.iv);\n+    const authenticationTag = this.decodeBase64(\n+      required(METADATA.authenticationTag),\n+      METADATA.authenticationTag,\n+    );\n+\n+    if (wrappedKey.length === 0) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        \"The wrapped data key is empty.\",\n+      );\n+    }\n+\n+    if (iv.length !== IV_LENGTH_BYTES) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        `The AES-GCM initialization vector must be ${IV_LENGTH_BYTES} bytes.`,\n+      );\n+    }\n+\n+    if (authenticationTag.length !== AUTH_TAG_LENGTH_BYTES) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        `The AES-GCM authentication tag must be ${AUTH_TAG_LENGTH_BYTES} bytes.`,\n+      );\n+    }\n+\n+    return {\n+      keyId: required(METADATA.keyId),\n+      wrappedKey,\n+      iv,\n+      authenticationTag,\n+    };\n+  }\n+\n+  private decodeBase64(value: string, fieldName: string): Buffer {\n+    if (\n+      value.length % 4 !== 0 ||\n+      !/^[A-Za-z0-9+/]*={0,2}$/.test(value)\n+    ) {\n+      throw new EncryptedBlobError(\n+        \"cryptography\",\n+        \"read encryption metadata\",\n+        `Blob metadata field ${fieldName} is not valid base64.`,\n+      );\n+    }\n+\n+    return Buffer.from(value, \"base64\");\n+  }\n+\n+  private async streamToBuffer(\n+    readable: NodeJS.ReadableStream,\n+  ): Promise\u003cBuffer\u003e {\n+    const chunks: Buffer[] = [];\n+    for await (const chunk of readable) {\n+      chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));\n+    }\n+\n+    return Buffer.concat(chunks);\n+  }\n+}\n*** Add File: src/main.ts\n+import { createAppConfig } from \"./config.js\";\n+import { EncryptedBlobStorage } from \"./encryptedBlobStorage.js\";\n+import { EncryptedBlobError } from \"./errors.js\";\n+import { KeyVaultKeyManager } from \"./keyManagement.js\";\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const config = createAppConfig();\n+  const keyManager = new KeyVaultKeyManager(\n+    config.keyClient,\n+    config.credential,\n+    config.keyName,\n+  );\n+  const encryptedStorage = new EncryptedBlobStorage(\n+    config.containerClient,\n+    keyManager,\n+  );\n+  const sample = \"Client-side encryption with Azure Key Vault envelope keys.\";\n+\n+  const upload = await encryptedStorage.upload(config.blobName, sample);\n+  const decrypted = await encryptedStorage.download(config.blobName);\n+\n+  console.log(`Vault key ID: ${upload.keyId}`);\n+  console.log(`Wrapped DEK (base64): ${upload.wrappedDataKeyBase64}`);\n+  console.log(`Decrypted output: ${decrypted.toString(\"utf8\")}`);\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  if (error instanceof EncryptedBlobError) {\n+    console.error(\n+      `[${error.category}] ${error.operation}: ${error.message}`,\n+    );\n+  } else {\n+    console.error(error instanceof Error ? error.message : error);\n+  }\n+\n+  process.exitCode = 1;\n+});\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/package.json
@@ -1,0 +1,27 @@
+{
+  "name": "azure-encrypted-blob-uploader",
+  "version": "1.0.0",
+  "private": true,
+  "description": "Client-side AES-GCM envelope encryption for Azure Blob Storage using Azure Key Vault Keys.",
+  "type": "module",
+  "engines": {
+    "node": ">=20.0.0"
+  },
+  "scripts": {
+    "build": "tsc -p tsconfig.json",
+    "start": "node dist/main.js",
+    "demo": "npm run build && npm start",
+    "typecheck": "tsc -p tsconfig.json --noEmit"
+  },
+  "dependencies": {
+    "@azure/core-auth": "^1.11.0",
+    "@azure/identity": "^4.13.2",
+    "@azure/keyvault-keys": "^4.10.2",
+    "@azure/storage-blob": "^12.33.0"
+  },
+  "devDependencies": {
+    "@types/node": "^26.2.0",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/tsconfig.json
@@ -1,0 +1,20 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "rootDir": "src",
+    "outDir": 
... (truncated)
```

### ✅ 16. powershell — 25995ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund --ignore-scripts \u0026\u0026 npm run typecheck \u0026\u0026 npm run build","description":"Install dependencies and compile project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 65 packages in 23s

> azure-encrypted-blob-uploader@1.0.0 typecheck
> tsc -p tsconfig.json --noEmit

src/encryptedBlobStorage.ts(137,7): error TS2322: Type '{ [propertyName: string]: string; } | undefined' is not assignable to type 'Record<string, string>'.
  Type 'undefined' is not assignable to type 'Record<string, string>'.
<shellId: 1 completed with exit code 1>
```

### ✅ 17. apply_patch — 753ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/encryptedBlobStorage.ts\n@@\n-      metadata = response.metadata;\n+      metadata = response.metadata ?? {};\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/src/encryptedBlobStorage.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/src/encryptedBlobStorage.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/src/encryptedBlobStorage.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-js-ts-encrypted-uploader-js-ts-azure-skills-azure-skill-mcp-microsoft-skill-1297731346/src/encryptedBlobStorage.ts
@@ -134,7 +134,7 @@
       }
 
       ciphertext = await this.streamToBuffer(response.readableStreamBody);
-      metadata = response.metadata;
+      metadata = response.metadata ?? {};
     } catch (error) {
       throw new EncryptedBlobError(
         "storage",

```

### ✅ 18. powershell — 2800ms

**Input:**

```json
{"command":"npm run typecheck \u0026\u0026 npm run build","description":"Compile TypeScript project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
> azure-encrypted-blob-uploader@1.0.0 typecheck
> tsc -p tsconfig.json --noEmit


> azure-encrypted-blob-uploader@1.0.0 build
> tsc -p tsconfig.json
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `package-lock.json`
- `package.json`
- `src\config.ts`
- `src\encryptedBlobStorage.ts`
- `src\errors.ts`
- `src\keyManagement.ts`
- `src\main.ts`
- `tsconfig.json`

## Copilot Response

I’m grounding the implementation in the SDK acceptance criteria, then I’ll scaffold the project and run its local TypeScript checks.

The design will pin the exact Key Vault key version in blob metadata, zeroize plaintext DEKs on every path, and distinguish storage, vault, and local-crypto failures.

The project is scaffolded; I’m now compiling against the current Azure SDK types and will correct any API/type mismatches before finishing.

Created the complete TypeScript project with:

- AES-256-GCM client-side encryption and separate IV/authentication-tag metadata.
- Azure Key Vault **Keys** envelope encryption using `RSA-OAEP-256` wrap/unwrap operations.
- Versioned Key Vault key IDs stored with each blob.
- Shared `ManagedIdentityCredential` for Blob Storage and Key Vault.
- Upload/download buffer and file APIs, structured service errors, metadata validation, and DEK zeroization.
- Demo round-trip in `src/main.ts`.
- Complete `package.json`, `package-lock.json`, `tsconfig.json`, and `.env.example`.

Configure the environment variables from `.env.example`, ensure the managed identity has Blob Data permissions plus Key Vault key `get`, `wrapKey`, and `unwrapKey` permissions, then run:

```powershell
npm run demo
```

References: [Azure Key Vault Keys SDK](https://learn.microsoft.com/javascript/api/@azure/keyvault-keys/) and [Azure Blob Storage JavaScript SDK](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-nodejs).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete package.json with appropriate Azure SDK dependencies
- Includes TypeScript and Node.js type definitions
- Provides a suitable tsconfig.json for an ESM Node.js project

## Grader Results

- encrypted-uploader.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (23/25)
      - Uses `@azure/keyvault-keys` (Keys, NOT Secrets) — critical distinction: Pass
      - Uses Node.js built-in `crypto` module for local AES-GCM encryption: Pass
      - Uses `KeyClient` for key management and `CryptographyClient` for wrap/unwrap operations (NOT `SecretClient`): Pass
      - Constructs `CryptographyClient` with the key ID or key name: Pass
      - Uses `CryptographyClient` for `wrapKey()` and `unwrapKey()` operations: Pass
      - Specifies RSA key wrap algorithm (e.g., `"RSA-OAEP"` or `"RSA-OAEP-256"`): Pass
      - Key material never leaves Key Vault (wrap/unwrap is server-side): Pass
      - Generates a random AES-256 DEK locally (32 bytes via `crypto.randomBytes`): Pass
      - Encrypts data with AES-GCM locally using the DEK (`crypto.createCipheriv("aes-256-gcm", ...)`): Pass
      - Wraps the DEK via Key Vault `wrapKey()`: Pass
      - Stores wrapped DEK as blob metadata (base64-encoded): Pass
      - Stores IV (initialization vector) in blob metadata (base64-encoded): Pass
      - Stores GCM auth tag in blob metadata (base64-encoded) — in Node.js the auth tag is separate from ciphertext: Pass
      - Stores vault key identifier in blob metadata: Pass
      - For decryption: retrieves wrapped DEK from metadata, unwraps via Key Vault, sets auth tag via `decipher.setAuthTag()`, decrypts locally: Pass
      - Uses AES-GCM (not AES-CBC, AES-ECB, or other modes): Pass
      - Generates random IV for each encryption (typically 12 bytes for GCM): Pass
      - Retrieves auth tag via `cipher.getAuthTag()` after encryption finalize: Pass
      - Sets auth tag via `decipher.setAuthTag()` before decryption finalize: Pass
      - Handles Key Vault errors (key disabled, key not found) via RestError: Fail
      - Handles Storage errors (blob not found) via RestError with statusCode: Fail
      - NOT using `SecretClient` instead of `KeyClient`/`CryptographyClient`: Pass
      - NOT encrypting data directly with the vault key (should be envelope encryption): Pass
      - NOT storing raw DEK in plaintext: Pass
      - NOT omitting the GCM auth tag from blob metadata (decryption will fail): Pass
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
| `Criteria from prompt file` | prompt_review | 92% | 1.00 | 0.9200 | 11.6% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `Async/Await Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 12.6% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 7.9200** | **72.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-js-ts-encrypted-uploader --config js-ts-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
