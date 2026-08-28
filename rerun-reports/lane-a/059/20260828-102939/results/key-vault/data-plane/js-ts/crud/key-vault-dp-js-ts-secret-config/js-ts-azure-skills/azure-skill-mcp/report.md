# Evaluation Report: key-vault-dp-js-ts-secret-config

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 742.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-js-ts-secret-config` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 20/23 |
| Duration | 742.0s |
| Timestamp | 2026-08-28T02:29:39Z |
| Files Generated | 12 |
| Event Count | 12463 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 611.1s |
| Review | 129.3s |
| **Total** | **742.0s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-app-onboard |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=506246 out=10906 |
| Turn Count | 12 |

## Error

```
evaluation failed: sending prompt: waiting for session.idle: context deadline exceeded
```

**Details:**

```
sending prompt: waiting for session.idle: context deadline exceeded
```

## Tool Calls

### ✅ 1. skill — 680ms

**Input:**

```json
{"skill":"azure-app-onboard"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure App Onboard

> ⛔ **Every repo goes through the full pipeline (Steps 1–10). No exceptions.** Do not skip steps, refuse, or short-circuit based on what you recognize. Follow the Workflow table below sequentially — read each step's references before acting.

## Quick Reference

| Property | Value |
|----------|-------|
| Best for | Developers who know what to build but not which Azure services to use |
| Inputs | Business idea or existing codebase, budget/scale preferences (optional) |
| Outputs | Architecture plan, cost estimate, IaC files, deployed Azure resources |
| Phases | Discover → Architect → Scaffold → Deploy (self-contained, no external skill calls) |

## When to Use This Skill

- Deploy existing code without knowing which Azure services to use
- Check if your existing code is ready to deploy to Azure
- Move an existing app to Azure without rewriting or with minimal changes
- Get cost estimates before committing to infrastructure
- Understand architecture decisions and rejected alternatives
- Get answers to Azure architecture or service selection questions (e.g., "What database should I use?")
- Get guided Azure onboarding without prior experience

## When NOT to Use

| Scenario | Use Instead |
|----------|-------------|
| Run `azd up` or execute an existing deployment | `azure-deploy` |
| Optimize existing Azure spend | `azure-cost` |
| Generate Bicep/Terraform for a known architecture | `azure-prepare` |
| Validate infrastructure or run preflight checks | `azure-validate` |
| Troubleshoot a running Azure deployment | `azure-diagnostics` |
| Deploy to or manage AKS/Kubernetes directly | `azure-kubernetes` |
| Look up or list existing Azure resources | `azure-resource-lookup` |

## Pipeline Rules

> ⛔ **You MUST read [`references/pipeline-rules.md`](references/pipeline-rules.md) at the start of every AppOnboard session.** It contains approval gates, phase lifecycle, session artifacts, deploy-as-is, and sec
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 183911ms

**Input:**

```json
{"intent":"Create a local TypeScript Node.js application that accesses Azure Key Vault secrets with DefaultAzureCredential/managed identity, caching, expiry checks, version retrieval, rotation, and safe delete/purge demonstration. No live resource operations or deployment.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 183913ms

**Input:**

```json
{"intent":"Find official current TypeScript @azure/keyvault-secrets and @azure/identity guidance for SecretClient getSecret with version, setSecret expiry, beginDeleteSecret waitUntilDone, purgeDeletedSecret, listPropertiesOfSecretVersions, DefaultAzureCredential managed identity, and RestError 404 handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 183903ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-app-onboard","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 183558ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 182120ms

**Input:**

```json
{"intent":"Get code-generation best practices for a TypeScript Node.js app using Azure Key Vault Secrets and managed identity.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 7. azure-documentation (via azure) — 182121ms

**Input:**

```json
{"intent":"Get available documentation search commands for Azure Key Vault Secrets JavaScript SDK.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. glob — 182115ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills","pattern":"**/*key*vault*/**/*"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-cost\cost-query\examples.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-cost\cost-forecast\examples.md
```

### ✅ 9. glob — 181801ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-prepare\references\services\key-vault\sdk.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-prepare\references\services\key-vault\README.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-prepare\references\services\key-vault\bicep.md
```

### ✅ 10. view — 33493ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-prepare\\references\\services\\key-vault\\sdk.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/README.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/README.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/README.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/README.md
@@ -1,56 +1,56 @@
 # Azure Key Vault
 
 Centralized secrets, keys, and certificate management.
 
 ## When to Use
 
 - Storing application secrets
 - Managing certificates
 - Storing encryption keys
 - Centralizing secret management
 - Enabling secret rotation
 
 ## Required Supporting Resources
 
 | Resource | Purpose |
 |----------|---------|
 | None required | Key Vault is self-contained |
 | Private Endpoint | Secure access (optional) |
 
 ## SKU Selection
 
 | SKU | Features |
 |-----|----------|
 | Standard | Software-protected keys |
 | Premium | HSM-protected keys |
 
 ## RBAC Roles
 
 | Role | Permissions |
 |------|-------------|
 | Key Vault Administrator | Full access |
 | Key Vault Secrets Officer | Manage secrets |
 | Key Vault Secrets User | Read secrets |
 | Key Vault Certificates Officer | Manage certificates |
 | Key Vault Crypto Officer | Manage keys |
 
 ## Environment Variables
 
 | Variable | Value |
 |----------|-------|
 | `KEY_VAULT_URL` | `https://{vault-name}.vault.azure.net/` |
 | `KEY_VAULT_NAME` | Vault name |
 
 ## Best Practices
 
 1. **Always use RBAC** over access policies
 2. **Enable soft delete and purge protection** for production
 3. **Use managed identities** instead of storing keys in apps
 4. **Set expiration dates** on secrets
 5. **Use separate vaults** f
... (truncated)
```

### ✅ 11. view — 33168ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-prepare\\references\\services\\key-vault\\README.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/sdk.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/sdk.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/sdk.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/services/key-vault/sdk.md
@@ -1,73 +1,73 @@
 # Key Vault - SDK Patterns
 
 ## Node.js
 
 > **Auth:** `DefaultAzureCredential` is for local development. See [auth-best-practices.md](../../auth-best-practices.md) for production patterns.
 
 ```javascript
 const { SecretClient } = require("@azure/keyvault-secrets");
 const { DefaultAzureCredential } = require("@azure/identity");
 
 const client = new SecretClient(
   process.env.KEY_VAULT_URL,
   new DefaultAzureCredential()
 );
 
 const secret = await client.getSecret("database-connection-string");
 console.log(secret.value);
 ```
 
 ## Python
 
 > **Auth:** `DefaultAzureCredential` is for local development. See [auth-best-practices.md](../../auth-best-practices.md) for production patterns.
 
 ```python
 from azure.keyvault.secrets import SecretClient
 from azure.identity import DefaultAzureCredential
 
 client = SecretClient(
     vault_url=os.environ["KEY_VAULT_URL"],
     credential=DefaultAzureCredential()
 )
 
 secret = client.get_secret("database-connection-string")
 print(secret.value)
 ```
 
 ## .NET
 
 > **Auth:** `DefaultAzureCredential` is for local development. See [auth-best-practices.md](../../auth-best-practices.md) for production patterns.
 
 ```csharp
 var client = new SecretClient(
     new Uri(Environment.GetEnvironmentVariable("KEY_VAULT_URL")),
     new DefaultAzure
... (truncated)
```

### ✅ 12. web_search (via github-mcp-server) — 34117ms

**Input:**

```json
{"query":"site:learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets SecretClient getSecret version beginDeleteSecret waitUntilDone purgeDeletedSecret setSecret expiresOn DefaultAzureCredential TypeScript"}
```

**Output:**

```
No relevant web results found.
```

### ✅ 13. web_search (via github-mcp-server) — 33797ms

**Input:**

```json
{"query":"site:learn.microsoft.com/en-us/javascript/api/@azure/identity DefaultAzureCredential managed identity TypeScript Azure hosted application"}
```

**Output:**

```
No relevant web results found.
```

### ❌ 14. view — 12975ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-prepare\\references\\services\\auth-best-practices.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 15. powershell — 13316ms

**Input:**

```json
{"command":"npm view @azure/keyvault-secrets version engines --json; npm view @azure/identity version engines --json; npm view typescript version --json; npm view tsx version --json","description":"Check current package versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
{
  "version": "4.11.2",
  "engines": {
    "node": ">=20.0.0"
  }
}
{
  "version": "4.13.2",
  "engines": {
    "node": ">=22.0.0"
  }
}
"7.0.2"
"4.23.12"
<shellId: 0 completed with exit code 0>
```

### ✅ 16. glob — 760ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-prepare","pattern":"**/auth-best-practices.md"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-prepare\references\auth-best-practices.md
```

### ✅ 17. view — 660ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-prepare\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-prepare/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks environment variables, CLI tokens, and oth
... (truncated)
```

### ✅ 18. github-mcp-server-search_code (via github-mcp-server) — 5965ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"getSecret(secretName\" path:sdk/keyvault/keyvault-secrets"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/keyvault/keyvault-secrets/src/index.ts","text_matches":[{"fragment":"   * @param value - The value of the secret.\n   * @param options - The optional parameters.\n   */\n  public setSecret(\n    secretName: string,\n    value: string,\n    options: SetSecretOptions = {},","matches":[{"indices":[107,116],"text":"setSecret"},{"indices":[174,183],"text":"SetSecret"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/src/index.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/review/keyvault-secrets-node.api.md","text_matches":[{"fragment":"    restoreSecretBackup(secretBundleBackup: Uint8Array, options?: RestoreSecretBackupOptions): Promise<SecretProperties>;\n    setSecret(secretName: string, value: string, options?: SetSecretOptions): Promise<KeyVaultSecret>;\n    updateSecretProperties(secretName: string, secretVersion: string, options?: UpdateSecretPropertiesOptions): Promise<SecretProperties>;","matches":[{"indices":[126,135],"text":"setSecret"},{"indices":[181,190],"text":"SetSecret"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/review/keyvault-secrets-node.api.md?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"},{"fragment":"    enabled?: boolean;\n    readonly expiresOn?: Date;\n    readonly notBefore?: Date;","matches":[{"indices":[36,45],"text":"expiresOn"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/review/keyvault-secrets-node.api.md?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/README.md","text_matches":[{"fragment":"- `notBefore`: A given date after which the secret value can be retrieved.\n- `expiresOn`: A given date after wh
... (truncated)
```

### ✅ 19. github-mcp-server-search_code (via github-mcp-server) — 5642ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"beginDeleteSecret\" \"waitUntilDone\" path:sdk/keyvault/keyvault-secrets"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/keyvault/keyvault-secrets/src/keyVaultClient.ts","text_matches":[{"fragment":"  }\n\n  /** The purge deleted secret operation removes the secret permanently, without the possibility of recovery. This operation can only be enabled on a soft-delete enabled vault. This operation requires the secrets/purge permission. */\n  purgeDeletedSecret(\n    secretName: string,\n    options: PurgeDeletedSecretOptionalParams = { requestOptions: {} },\n  ): Promise<void> {","matches":[{"indices":[241,259],"text":"purgeDeletedSecret"},{"indices":[298,316],"text":"PurgeDeletedSecret"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/src/keyVaultClient.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/generated/keyVaultClient.ts","text_matches":[{"fragment":"  }\n\n  /** The purge deleted secret operation removes the secret permanently, without the possibility of recovery. This operation can only be enabled on a soft-delete enabled vault. This operation requires the secrets/purge permission. */\n  purgeDeletedSecret(\n    secretName: string,\n    options: PurgeDeletedSecretOptionalParams = { requestOptions: {} },\n  ): Promise<void> {","matches":[{"indices":[241,259],"text":"purgeDeletedSecret"},{"indices":[298,316],"text":"PurgeDeletedSecret"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/generated/keyVaultClient.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/src/index.ts","text_matches":[{"fragment":"   * @param secretName - The name of the secret.\n   * @param options - The optional parameters.\n   */\n  public purgeDeletedSecret(\n    secretName: string,\n    options: PurgeDeletedSecretOptions = {},\n  ): Promise<void> {","matches":[{"indices":[111,
... (truncated)
```

### ✅ 20. github-mcp-server-search_code (via github-mcp-server) — 5305ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"purgeDeletedSecret\" path:sdk/keyvault/keyvault-secrets"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 21. github-mcp-server-search_code (via github-mcp-server) — 4974ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"expiresOn\" \"setSecret\" path:sdk/keyvault/keyvault-secrets"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/keyvault/keyvault-secrets/samples-dev/helloWorld.ts","text_matches":[{"fragment":"  console.log(\"result: \", result);\n\n  // Read the secret we created\n  const secret = await client.getSecret(secretName);\n  console.log(\"secret: \", secret);\n\n  // Update the secret with different attributes","matches":[{"indices":[98,118],"text":"getSecret(secretName"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/samples-dev/helloWorld.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/README.md","text_matches":[{"fragment":"\nconst latestSecret = await client.getSecret(secretName);\nconsole.log(`Latest version of the secret ${secretName}: `, latestSecret);","matches":[{"indices":[35,55],"text":"getSecret(secretName"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/README.md?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"},{"fragment":"\nconst specificSecret = await client.getSecret(secretName, {\n  version: latestSecret.properties.version!,","matches":[{"indices":[37,57],"text":"getSecret(secretName"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/README.md?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/samples/v4/javascript/helloWorld.js","text_matches":[{"fragment":"  console.log(\"result: \", result);\n\n  // Read the secret we created\n  const secret = await client.getSecret(secretName);\n  console.log(\"secret: \", secret);\n\n  // Update the secret with different attributes","matches":[{"indices":[98,118],"text":"getSecret(secretName"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvau
... (truncated)
```

### ✅ 22. github-mcp-server-search_code (via github-mcp-server) — 6299ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"new ManagedIdentityCredential(process.env\" path:sdk/identity/identity"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/keyvault/keyvault-secrets/src/index.ts","text_matches":[{"fragment":"   * @param secretName - The name of the secret.\n   * @param options - The optional parameters.\n   */\n  public async beginDeleteSecret(\n    name: string,\n    options: BeginDeleteSecretOptions = {},\n  ): Promise<PollerLike<PollOperationState<DeletedSecret>, DeletedSecret>> {","matches":[{"indices":[117,134],"text":"beginDeleteSecret"},{"indices":[167,184],"text":"BeginDeleteSecret"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/src/index.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/README.md","text_matches":[{"fragment":"\n// Deleted secrets can also be recovered or purged.\n\n// recoverDeletedSecret returns a poller, just like beginDeleteSecret.\nconst recoverPoller = await client.beginRecoverDeletedSecret(secretName);\nawait recoverPoller.pollUntilDone();\n","matches":[{"indices":[106,123],"text":"beginDeleteSecret"},{"indices":[219,232],"text":"pollUntilDone"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/README.md?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/keyvault/keyvault-secrets/samples-dev/backupAndRestore.ts","text_matches":[{"fragment":"\n  // Delete the secret\n  console.log(\"about to delete\");\n  const deletePoller = await client.beginDeleteSecret(secretName);\n  await deletePoller.pollUntilDone();\n\n  // Purge the deleted secret","matches":[{"indices":[94,111],"text":"beginDeleteSecret"},{"indices":[146,159],"text":"pollUntilDone"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/keyvault/keyvault-secrets/samples-dev/backupAndRestore.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"pat
... (truncated)
```

### ✅ 23. github-mcp-server-search_code (via github-mcp-server) — 5978ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"new ManagedIdentityCredential({\" path:sdk/identity/identity"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/identity/identity/integration/AzureKubernetes/index.js","text_matches":[{"fragment":"  try {\n    const blobServiceClient = new BlobServiceClient(\n      blobUrl,\n      new ManagedIdentityCredential({\n        clientId,\n      }),\n    );","matches":[{"indices":[82,113],"text":"new ManagedIdentityCredential({"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/identity/identity/integration/AzureKubernetes/index.js?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/identity/identity/integration/AzureWebApps/src/index.ts","text_matches":[{"fragment":"  }\n  try {\n    const accountUserAssigned = process.env.IDENTITY_STORAGE_NAME_USER_ASSIGNED;\n    const credentialUserAssigned = new ManagedIdentityCredential({\n      clientId: process.env.IDENTITY_USER_DEFINED_CLIENT_ID,\n    });\n    const clientUserAssigned = new BlobServiceClient(","matches":[{"indices":[128,159],"text":"new ManagedIdentityCredential({"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/identity/identity/integration/AzureWebApps/src/index.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},{"path":"sdk/identity/identity/integration/AzureKubernetes/src/index.ts","text_matches":[{"fragment":"\napp.get(\"/managed-identity/user-assigned\", async (req: express.Request, res: express.Response) => {\n  try {\n    const credential = new ManagedIdentityCredential({ clientId: userAssignedClientId });\n    await testStorageAccess(credential, storageAccountUserAssigned);\n\n    res.json({ success: true });","matches":[{"indices":[132,163],"text":"new ManagedIdentityCredential({"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/108311273/contents/sdk/identity/identity/integration/AzureKubernetes/src/index.ts?ref=528bf3b6a802c765f56fe0791d9d12f6968feec6","property":"content"}]},
... (truncated)
```

### ✅ 24. github-mcp-server-search_code (via github-mcp-server) — 5665ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-js \"pollUntilDone\" \"beginDeleteSecret\" path:sdk/keyvault/keyvault-secrets"}
```

**Output:**

```
{"incomplete_results":false,"items":[],"total_count":0}
```

### ✅ 25. github-mcp-server-get_file_contents (via github-mcp-server) — 6632ms

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/keyvault-secrets/README.md","ref":"refs/heads/main","repo":"azure-sdk-for-js"}
```

**Output:**

```
successfully downloaded text file (SHA: fd146ca5c10535bcaf7c961e9fb6789d0ffbd30e)# Azure Key Vault Secret client library for JavaScript

Azure Key Vault is a service that allows you to encrypt authentication keys, storage account keys, data encryption keys, .pfx files, and passwords by using secured keys.
If you would like to know more about Azure Key Vault, you may want to review: [What is Azure Key Vault?](https://learn.microsoft.com/azure/key-vault/key-vault-overview)

Azure Key Vault Secrets management allows you to securely store and
tightly control access to tokens, passwords, certificates, API keys,
and other secrets.

Use the client library for Azure Key Vault Secrets in your Node.js application to:

- Get, set and delete secrets.
- Update a secret and it's attributes.
- Backup and restore a secret.
- Get, purge or recover a deleted secret.
- Get all the versions of a secret.
- Get all secrets.
- Get all deleted secrets.

> Note: This package cannot be used in the browser due to Azure Key Vault service limitations, please refer to [this document](https://github.com/Azure/azure-sdk-for-js/blob/main/samples/cors/ts/README.md) for guidance.

Key links:

- [Source code](https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/keyvault/keyvault-secrets)
- [Package (npm)](https://www.npmjs.com/package/@azure/keyvault-secrets)
- [API Reference Documentation](https://learn.microsoft.com/javascript/api/@azure/keyvault-secrets)
- [Product documentation](https://azure.microsoft.com/services/key-vault/)
- [Samples](https://github.com/Azure/azure-sdk-for-js/tree/main/sdk/keyvault/keyvault-secrets/samples)

## Getting started

### Currently supported environments

- [LTS versions of Node.js](https://github.com/nodejs/release#release-schedule)

### Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- A [Key Vault resource](https://learn.microsoft.com/azure/key-vault/quick-create-portal)
- An existing [Azure Key Vault][azure_keyvault]. If you need to
... (truncated)
```

### ✅ 26. apply_patch — 651ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-key-vault-config-provider\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"engines\": {\n+    \"node\": \"\u003e=22.0.0\"\n+  },\n+  \"scripts\": {\n+    \"build\": \"tsc -p tsconfig.json\",\n+    \"check\": \"tsc -p tsconfig.json --noEmit\",\n+    \"demo\": \"tsx src/main.ts\",\n+    \"start\": \"node dist/main.js\",\n+    \"test\": \"tsx --test test/**/*.test.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.13.2\",\n+    \"@azure/keyvault-secrets\": \"^4.11.2\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^24.3.0\",\n+    \"tsx\": \"^4.23.12\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \".\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true,\n+    \"declaration\": true,\n+    \"sourceMap\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"],\n+  \"exclude\": [\"dist\", \"node_modules\", \"test\"]\n+}\n*** Add File: .env.example\n+# Use \"mock\" for the local, offline demo or \"azure\" for Azure Key Vault.\n+KEY_VAULT_DEMO_MODE=mock\n+\n+# Required when KEY_VAULT_DEMO_MODE=azure.\n+KEY_VAULT_URL=https://your-vault-name.vault.azure.net/\n+\n+# Optional: client ID of a user-assigned managed identity.\n+AZURE_CLIENT_ID=\n+\n+# Destructive Azure cleanup is skipped unless this is explicitly enabled.\n+RUN_DESTRUCTIVE_CLEANUP=false\n+\n+# Expiry warning window used by the cache.\n+SECRET_EXPIRY_WARNING_DAYS=7\n*** Add File: src/secret-client.ts\n+export interface SecretPropertiesLike {\n+  name: string;\n+  version?: string;\n+  expiresOn?: Date;\n+}\n+\n+export interface KeyVaultSecretLike {\n+  value?: string;\n+  properties: SecretPropertiesLike;\n+}\n+\n+export interface GetSecretOptionsLike {\n+  version?: string;\n+}\n+\n+export interface SetSecretOptionsLike {\n+  expiresOn?: Date;\n+}\n+\n+export interface DeleteSecretPollerLike {\n+  pollUntilDone(): Promise\u003cunknown\u003e;\n+}\n+\n+export interface SecretClientLike {\n+  getSecret(name: string, options?: GetSecretOptionsLike): Promise\u003cKeyVaultSecretLike\u003e;\n+  setSecret(\n+    name: string,\n+    value: string,\n+    options?: SetSecretOptionsLike,\n+  ): Promise\u003cKeyVaultSecretLike\u003e;\n+  beginDeleteSecret(name: string): Promise\u003cDeleteSecretPollerLike\u003e;\n+  purgeDeletedSecret(name: string): Promise\u003cvoid\u003e;\n+}\n*** Add File: src/secret-provider.ts\n+import type { SecretClientLike } from \"./secret-client.js\";\n+\n+export interface SecretValue {\n+  name: string;\n+  value: string;\n+  version?: string;\n+  expiresOn?: Date;\n+  found: boolean;\n+}\n+\n+interface ServiceErrorLike {\n+  code?: unknown;\n+  statusCode?: unknown;\n+}\n+\n+function isServiceErrorLike(error: unknown): error is ServiceErrorLike {\n+  return typeof error === \"object\" \u0026\u0026 error !== null;\n+}\n+\n+function isSecretNotFound(error: unknown): boolean {\n+  if (!isServiceErrorLike(error)) {\n+    return false;\n+  }\n+\n+  return error.statusCode === 404 || error.code === \"SecretNotFound\";\n+}\n+\n+export function expiresWithin(\n+  expiresOn: Date | undefined,\n+  warningWindowMs: number,\n+  now = new Date(),\n+): boolean {\n+  if (!expiresOn) {\n+    return false;\n+  }\n+\n+  return expiresOn.getTime() - now.getTime() \u003c= warningWindowMs;\n+}\n+\n+export class KeyVaultSecretProvider {\n+  public constructor(private readonly client: SecretClientLike) {}\n+\n+  public async getSecret(\n+    name: string,\n+    defaultValue: string,\n+    version?: string,\n+  ): Promise\u003cSecretValue\u003e {\n+    try {\n+      const secret = await this.client.getSecret(name, version ? { version } : undefined);\n+\n+      return {\n+        name,\n+        value: secret.value ?? defaultValue,\n+        version: secret.properties.version,\n+        expiresOn: secret.properties.expiresOn,\n+        found: true,\n+      };\n+    } catch (error: unknown) {\n+      if (!isSecretNotFound(error)) {\n+        throw error;\n+      }\n+\n+      return {\n+        name,\n+        value: defaultValue,\n+        found: false,\n+      };\n+    }\n+  }\n+\n+  public async getSecretVersion(\n+    name: string,\n+    version: string,\n+    defaultValue: string,\n+  ): Promise\u003cSecretValue\u003e {\n+    return this.getSecret(name, defaultValue, version);\n+  }\n+\n+  public async inspectExpiry(name: string, version?: string): Promise\u003cDate | undefined\u003e {\n+    const secret = await this.getSecret(name, \"\", version);\n+    return secret.expiresOn;\n+  }\n+\n+  public isNearExpiry(\n+    secret: Pick\u003cSecretValue, \"expiresOn\"\u003e,\n+    warningWindowMs: number,\n+    now = new Date(),\n+  ): boolean {\n+    return expiresWithin(secret.expiresOn, warningWindowMs, now);\n+  }\n+}\n*** Add File: src/caching-secret-provider.ts\n+import {\n+  expiresWithin,\n+  type KeyVaultSecretProvider,\n+  type SecretValue,\n+} from \"./secret-provider.js\";\n+\n+export interface CachedSecret extends SecretValue {\n+  fetchedAt: Date;\n+}\n+\n+export interface RequiredSecret {\n+  name: string;\n+  defaultValue: string;\n+}\n+\n+export class CachingSecretProvider {\n+  private readonly cache = new Map\u003cstring, CachedSecret\u003e();\n+  private readonly defaults = new Map\u003cstring, string\u003e();\n+\n+  public constructor(\n+    private readonly provider: KeyVaultSecretProvider,\n+    private readonly warningWindowMs: number,\n+  ) {\n+    if (!Number.isFinite(warningWindowMs) || warningWindowMs \u003c 0) {\n+      throw new RangeError(\"warningWindowMs must be a non-negative finite number.\");\n+    }\n+  }\n+\n+  public async loadRequired(secrets: readonly RequiredSecret[]): Promise\u003cReadonlyMap\u003cstring, CachedSecret\u003e\u003e {\n+    await Promise.all(\n+      secrets.map(async ({ name, defaultValue }) =\u003e {\n+        this.defaults.set(name, defaultValue);\n+        await this.refresh(name);\n+      }),\n+    );\n+\n+    return new Map(this.cache);\n+  }\n+\n+  public async get(name: string, defaultValue = \"\"): Promise\u003cCachedSecret\u003e {\n+    if (!this.defaults.has(name)) {\n+      this.defaults.set(name, defaultValue);\n+    }\n+\n+    const cached = this.cache.get(name);\n+    if (!cached || expiresWithin(cached.expiresOn, this.warningWindowMs)) {\n+      return this.refresh(name);\n+    }\n+\n+    return cached;\n+  }\n+\n+  public async refresh(name: string, defaultValue?: string): Promise\u003cCachedSecret\u003e {\n+    if (defaultValue !== undefined) {\n+      this.defaults.set(name, defaultValue);\n+    }\n+\n+    const resolvedDefault = this.defaults.get(name) ?? \"\";\n+    const secret = await this.provider.getSecret(name, resolvedDefault);\n+    const cached = { ...secret, fetchedAt: new Date() };\n+    this.cache.set(name, cached);\n+    return cached;\n+  }\n+\n+  public getNearExpiry(now = new Date()): readonly CachedSecret[] {\n+    return [...this.cache.values()].filter((secret) =\u003e\n+      expiresWithin(secret.expiresOn, this.warningWindowMs, now),\n+    );\n+  }\n+\n+  public async refreshNearExpiry(now = new Date()): Promise\u003creadonly CachedSecret[]\u003e {\n+    const names = this.getNearExpiry(now).map(({ name }) =\u003e name);\n+    return Promise.all(names.map((name) =\u003e this.refresh(name)));\n+  }\n+\n+  public snapshot(): ReadonlyMap\u003cstring, CachedSecret\u003e {\n+    return new Map(this.cache);\n+  }\n+}\n*** Add File: src/key-vault-config.ts\n+import { ManagedIdentityCredential } from \"@azure/identity\";\n+import { SecretClient } from \"@azure/keyvault-secrets\";\n+\n+export interface KeyVaultEnvironment {\n+  KEY_VAULT_URL?: string;\n+  AZURE_CLIENT_ID?: string;\n+}\n+\n+export function createKeyVaultSecretClient(\n+  environment: KeyVaultEnvironment = process.env,\n+): SecretClient {\n+  const vaultUrl = environment.KEY_VAULT_URL;\n+  if (!vaultUrl) {\n+    throw new Error(\"KEY_VAULT_URL is required when using Azure Key Vault.\");\n+  }\n+\n+  let parsedUrl: URL;\n+  try {\n+    parsedUrl = new URL(vaultUrl);\n+  } catch {\n+    throw new Error(\"KEY_VAULT_URL must be a valid URL.\");\n+  }\n+\n+  if (parsedUrl.protocol !== \"https:\") {\n+    throw new Error(\"KEY_VAULT_URL must use HTTPS.\");\n+  }\n+\n+  const credential = environment.AZURE_CLIENT_ID\n+    ? new ManagedIdentityCredential({ clientId: environment.AZURE_CLIENT_ID })\n+    : new ManagedIdentityCredential();\n+\n+  return new SecretClient(parsedUrl.toString(), credential);\n+}\n*** Add File: src/secret-rotation.ts\n+import type { SecretClientLike } from \"./secret-client.js\";\n+\n+export interface RotationResult {\n+  name: string;\n+  version?: string;\n+  expiresOn: Date;\n+}\n+\n+export class SecretRotationHelper {\n+  public constructor(private readonly client: SecretClientLike) {}\n+\n+  public async rotateSecret(\n+    name: string,\n+    newValue: string,\n+    expiresOn: Date,\n+    now = new Date(),\n+  ): Promise\u003cRotationResult\u003e {\n+    if (!name.trim()) {\n+      throw new Error(\"Secret name must not be empty.\");\n+    }\n+    if (!newValue) {\n+      throw new Error(\"Secret value must not be empty.\");\n+    }\n+    if (expiresOn.getTime() \u003c= now.getTime()) {\n+      throw new Error(\"The new secret expiry date must be in the future.\");\n+    }\n+\n+    const secret = await this.client.setSecret(name, newValue, { expiresOn });\n+    return {\n+      name,\n+      version: secret.properties.version,\n+      expiresOn,\n+    };\n+  }\n+\n+  /**\n+   * Permanently removes the secret name and every version under it.\n+   * Key Vault does not support deleting only one historical secret version.\n+   */\n+  public async deleteAndPurgeSecret(name: string, confirmation: string): Promise\u003cvoid\u003e {\n+    if (confirmation !== name) {\n+      throw new Error(`Permanent deletion requires confirmation equal to \"${name}\".`);\n+    }\n+\n+    const deletePoller = await this.client.beginDeleteSecret(name);\n+    await deletePoller.pollUntilDone();\n+    await this.client.purgeDeletedSecret(name);\n+  }\n+}\n*** Add File: src/in-memory-secret-client.ts\n+import type {\n+  GetSecretOptionsLike,\n+  KeyVaultSecretLike,\n+  SecretClientLike,\n+  SetSecretOptionsLike,\n+} from \"./secret-client.js\";\n+\n+interface StoredVersion {\n+  value: string;\n+  version: string;\n+  expiresOn?: Date;\n+}\n+\n+export class InMemorySecretClient implements SecretClientLike {\n+  private readonly active = new Map\u003cstring, StoredVersion[]\u003e();\n+  private readonly deleted = new Map\u003cstring, StoredVersion[]\u003e();\n+  private nextVersion = 1;\n+\n+  public async getSecret(\n+    name: string,\n+    options: GetSecretOptionsLike = {},\n+  ): Promise\u003cKeyVaultSecretLike\u003e {\n+    const versions = this.active.get(name);\n+    const secret = options.version\n+      ? versions?.find(({ version }) =\u003e version === options.version)\n+      : versions?.at(-1);\n+\n+    if (!secret) {\n+      throw Object.assign(new Error(`Secret \"${name}\" was not found.`), {\n+        code: \"SecretNotFound\",\n+        statusCode: 404,\n+      });\n+    }\n+\n+    return this.toSecret(name, secret);\n+  }\n+\n+  public async setSecret(\n+    name: string,\n+    value: string,\n+    options: SetSecretOptionsLike = {},\n+  ): Promise\u003cKeyVaultSecretLike\u003e {\n+    if (this.deleted.has(name)) {\n+      throw new Error(`Secret \"${name}\" is soft-deleted and must be recovered or purged first.`);\n+    }\n+\n+    const stored: StoredVersion = {\n+      value,\n+      version: `v${this.nextVersion++}`,\n+      expiresOn: options.expiresOn,\n+    };\n+    const versions = this.active.get(name) ?? [];\n+    versions.push(stored);\n+    this.active.set(name, versions);\n+    return this.toSecret(name, stored);\n+  }\n+\n+  public async beginDeleteSecret(name: string): Promise\u003c{ pollUntilDone(): Promise\u003cvoid\u003e }\u003e {\n+    if (!this.active.has(name)) {\n+      throw Object.assign(new Error(`Secret \"${name}\" was not found.`), {\n+        code: \"SecretNotFound\",\n+        statusCode: 404,\n+      });\n+    }\n+\n+    return {\n+      pollUntilDone: async () =\u003e {\n+        const versions = this.active.get(name);\n+        if (versions) {\n+          this.active.delete(name);\n+          this.deleted.set(name, versions);\n+        }\n+      },\n+    };\n+  }\n+\n+  public async purgeDeletedSecret(name: string): Promise\u003cvoid\u003e {\n+    if (!this.deleted.delete(name)) {\n+      throw new Error(`Deleted secret \"${name}\" was not found.`);\n+    }\n+  }\n+\n+  private toSecret(name: string, secret: StoredVersion): KeyVaultSecretLike {\n+    return {\n+      value: secret.value,\n+      properties: {\n+        name,\n+        version: secret.version,\n+        expiresOn: secret.expiresOn,\n+      },\n+    };\n+  }\n+}\n*** Add File: src/main.ts\n+import { CachingSecretProvider, type CachedSecret } from \"./caching-secret-provider.js\";\n+import { InMemorySecretClient } from \"./in-memory-secret-client.js\";\n+import { createKeyVaultSecretClient } from \"./key-vault-config.js\";\n+import type { SecretClientLike } from \"./secret-client.js\";\n+import { KeyVaultSecretProvider } from \"./secret-provider.js\";\n+import { SecretRotationHelper } from \"./secret-rotation.js\";\n+\n+const DAY_MS = 24 * 60 * 60 * 1_000;\n+\n+function mask(value: string): string {\n+  if (!value) {\n+    return \"(empty)\";\n+  }\n+  return value.length \u003c= 4 ? \"****\" : `${value.slice(0, 2)}***${value.slice(-2)}`;\n+}\n+\n+function printSecret(label: string, secret: CachedSecret): void {\n+  console.log(\n+    `${label}: name=${secret.name}, value=${mask(secret.value)}, version=${secret.version ?? \"n/a\"}, ` +\n+      `source=${secret.found ? \"vault\" : \"default\"}, expires=${secret.expiresOn?.toISOString() ?? \"none\"}`,\n+  );\n+}\n+\n+async function createDemoClient(mode: string): Promise\u003cSecretClientLike\u003e {\n+  if (mode === \"azure\") {\n+    return createKeyVaultSecretClient();\n+  }\n+  if (mode !== \"mock\") {\n+    throw new Error('KEY_VAULT_DEMO_MODE must be either \"mock\" or \"azure\".');\n+  }\n+\n+  const client = new InMemorySecretClient();\n+  const now = Date.now();\n+  await client.setSecret(\"database-password\", \"local-db-password\", {\n+    expiresOn: new Date(now + 30 * DAY_MS),\n+  });\n+  await client.setSecret(\"api-key\", \"local-api-key\", {\n+    expiresOn: new Date(now + 3 * DAY_MS),\n+  });\n+  return client;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const mode = process.env.KEY_VAULT_DEMO_MODE ?? \"mock\";\n+  const warningDays = Number(process.env.SECRET_EXPIRY_WARNING_DAYS ?? \"7\");\n+  if (!Number.isFinite(warningDays) || warningDays \u003c 0) {\n+    throw new Error(\"SECRET_EXPIRY_WARNING_DAYS must be a non-negative number.\");\n+  }\n+\n+  console.log(`1. Creating secret client in ${mode} mode.`);\n+  const client = await createDemoClient(mode);\n+  const provider = new KeyVaultSecretProvider(client);\n+  const cache = new CachingSecretProvider(provider, warningDays * DAY_MS);\n+  const rotation = new SecretRotationHelper(client);\n+\n+  console.log(\"2. Bulk-loading required configuration.\");\n+  await cache.loadRequired([\n+    { name: \"database-password\", defaultValue: \"database-password-not-configured\" },\n+    { name: \"api-key\", defaultValue: \"api-key-not-configured\" },\n+    { name: \"optional-feature-token\", defaultValue: \"feature-disabled\" },\n+  ]);\n+  for (const secret of cache.snapshot().values()) {\n+    printSecret(\"   loaded\", secret);\n+  }\n+\n+  console.log(\"3. Reading database-password from the in-memory cache.\");\n+  printSecret(\"   cached\", await cache.get(\"database-password\"));\n+\n+  console.log(\"4. Refreshing database-password on demand.\");\n+  printSecret(\"   refreshed\", await cache.refresh(\"database-password\"));\n+\n+  console.log(`5. Checking for secrets expiring within ${warningDays} day(s).`);\n+  const nearExpiry = cache.getNearExpiry();\n+  if (nearExpiry.length === 0) {\n+    console.log(\"   No cached secrets are near expiry.\");\n+  } else {\n+    for (const secret of nearExpiry) {\n+      console.warn(`   WARNING: ${secret.name} expires on ${secret.expiresOn?.toISOString()}.`);\n+    }\n+    const refreshed = await cache.refreshNearExpiry();\n+    console.log(`   Automatically re-fetched ${refreshed.length} near-expiry secret(s).`);\n+  }\n+\n+  console.log(\"6. Rotating api-key by creating a new Key Vault secret version.\");\n+  const expiresOn = new Date(Date.now() + 90 * DAY_MS);\n+  const rotated = await rotation.rotateSecret(\"api-key\", `rotated-${Date.now()}`, expiresOn);\n+  console.log(\n+    `   Created version=${rotated.version ?? \"n/a\"}, expires=${rotated.expiresOn.toISOString()}.`,\n+  );\n+  printSecret(\"   refreshed rotation\", await cache.refresh(\"api-key\"));\n+\n+  console.log(\"7. Demonstrating long-running delete followed by purge.\");\n+  const destructiveCleanupAllowed =\n+    mode === \"mock\" || process.env.RUN_DESTRUCTIVE_CLEANUP === \"true\";\n+  if (!destructiveCleanupAllowed) {\n+    console.log(\n+      \"   Skipped for Azure. Set RUN_DESTRUCTIVE_CLEANUP=true only when permanently deleting all versions is intended.\",\n+    );\n+  } else {\n+    await rotation.deleteAndPurgeSecret(\"api-key\", \"api-key\");\n+    console.log(\"   Delete completed and the soft-deleted secret was purged.\");\n+    const afterPurge = await provider.getSecret(\"api-key\", \"api-key-not-configured\");\n+    console.log(`   Post-purge lookup source=${afterPurge.found ? \"vault\" : \"default\"}.`);\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Demo failed:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: test/secret-provider.test.ts\n+import assert from \"node:assert/strict\";\n+import test from \"node:test\";\n+import { CachingSecretProvider } from \"../src/caching-secret-provider.js\";\n+import { InMemorySecretClient } from \"../src/in-memory-secret-client.js\";\n+import { KeyVaultSecretProvider } from \"../src/secret-provider.js\";\n+import { SecretRotationHelper } from \"../src/secret-rotation.js\";\n+\n+const DAY_MS = 24 * 60 * 60 * 1_000;\n+\n+test(\"returns a default only when the secret does not exist\", async () =\u003e {\n+  const provider = new KeyVaultSecretProvider(new InMemorySecretClient());\n+  const secret = await provider.getSecret(\"missing\", \"fallback\");\n+\n+  assert.equal(secret.value, \"fallback\");\n+  assert.equal(secret.found, false);\n+});\n+\n+test(\"retrieves a specific secret version\", async () =\u003e {\n+  const client = new InMemorySecretClient();\n+  const first = await client.setSecret(\"setting\", \"first\");\n+  await client.setSecret(\"setting\", \"second\");\n+  const provider = new KeyVaultSecretProvider(client);\n+\n+  const secret = await provider.getSecretVersion(\n+    \"setting\",\n+    first.properties.version ?? \"\",\n+    \"fallback\",\n+  );\n+\n+  assert.equal(secret.value, \"first\");\n+});\n+\n+test(\"bulk-loads, caches, refreshes, and detects near-expiry secrets\", async () =\u003e {\n+  const client = new InMemorySecretClient();\n+  await client.setSecret(\"soon\", \"old\", { expiresOn: new Date(Date.now() + DAY_MS) });\n+  const cache = new CachingSecretProvider(new KeyVaultSecretProvider(client), 7 * DAY_MS);\n+\n+  await cache.loadRequired([{ name: \"soon\", defaultValue: \"fallback\" }]);\n+  assert.equal(cache.getNearExpiry().length, 1);\n+\n+  await client.setSecret(\"soon\", \"new\", { expiresOn: new Date(Date.now() + 30 * DAY_MS) });\n+  const refreshed = await cache.get(\"soon\");\n+\n+  assert.equal(refreshed.value, \"new\");\n+  assert.equal(cache.getNearExpiry().length, 0);\n+});\n+\n+test(\"rotation creates a version and cleanup waits before purging all versions\", async () =\u003e {\n+  const events: string[] = [];\n+  const client = new InMemorySecretClient();\n+  await client.setSecret(\"rotating\", \"old\");\n+  const rotation = new SecretRotationHelper({\n+    getSecret: client.getSecret.bind(client),\n+    setSecret: client.setSecret.bind(client),\n+    beginDeleteSecret: async (name) =\u003e {\n+      const poller = await client.beginDeleteSecret(name);\n+      events.push(\"delete-started\");\n+      return {\n+        pollUntilDone: async () =\u003e {\n+          await poller.pollUntilDone();\n+          events.push(\"delete-completed\");\n+        },\n+      };\n+    },\n+    purgeDeletedSecret: async (name) =\u003e {\n+      events.push(\"purge-started\");\n+      await client.purgeDeletedSecret(name);\n+    },\n+  });\n+\n+  const result = await rotation.rotateSecret(\n+    \"rotating\",\n+    \"new\",\n+    new Date(Date.now() + 30 * DAY_MS),\n+  );\n+  assert.ok(result.version);\n+\n+  await rotation.deleteAndPurgeSecret(\"rotating\", \"rotating\");\n+  assert.deepEqual(events, [\"delete-started\", \"delete-completed\", \"purge-started\"]);\n+});\n*** Add File: README.md\n+# Azure Key Vault configuration provider\n+\n+A TypeScript Node.js example that reads versioned secrets from Azure Key Vault, caches configuration in memory, refreshes near-expiry values, and rotates secrets by creating new versions.\n+\n+## Run locally without Azure\n+\n+The default demo uses an in-memory Key Vault-compatible client, so it is safe to run offline:\n+\n+```powershell\n+npm install\n+npm run demo\n+```\n+\n+## Run against Azure Key Vault\n+\n+Set `KEY_VAULT_DEMO_MODE=azure` and `KEY_VAULT_URL` to the vault URL. The application uses `ManagedIdentityCredential`; it contains no client secret or certificate. `AZURE_CLIENT_ID` is optional and selects a user-assigned managed identity. Otherwise, the system-assigned managed identity is used.\n+\n+The identity needs **Key Vault Secrets User** to read configuration and **Key Vault Secrets Officer** (or a narrower custom role with equivalent secret permissions) for rotation and deletion.\n+\n+```powershell\n+$env:KEY_VAULT_DEMO_MODE = \"azure\"\n+$env:KEY_VAULT_URL = \"https://your-vault-name.vault.azure.net/\"\n+npm run demo\n+```\n+\n+The Azure demo does not perform permanent cleanup unless `RUN_DESTRUCTIVE_CLEANUP=true`. Purging requires additional permission and fails when purge protection is enabled.\n+\n+## Important deletion behavior\n+\n+Calling `setSecret` for an existing name creates a new version; this is the normal rotation flow. Azure Key Vault cannot delete one historical version. `deleteAndPurgeSecret` therefore deletes and permanently purges the secret name **and every version**, waits for the long-running delete operation with `pollUntilDone()`, and requires the secret name as an explicit confirmation value.\n+\n+Secret values are masked in demo output to avoid leaking credentials into logs.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705/package.json
@@ -1,0 +1,26 @@
+{
+  "name": "azure-key-vault-config-provider",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "engines": {
+    "node": ">=22.0.0"
+  },
+  "scripts": {
+    "build": "tsc -p tsconfig.json",
+    "check": "tsc -p tsconfig.json --noEmit",
+    "demo": "tsx src/main.ts",
+    "start": "node dist/main.js",
+    "test": "tsx --test test/**/*.test.ts"
+  },
+  "dependencies": {
+    "@azure/identity": "^4.13.2",
+    "@azure/keyvault-secrets": "^4.11.2"
+  },
+  "devDependencies": {
+    "@types/node": "^24.3.0",
+    "tsx": "^4.23.12",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3724879705/tsconfig.json
@@ -1,0 +1,20 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "rootDir": ".",
+    "outDir": "dist",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "exactOptionalPropertyTypes": true,
+    "esModuleInterop": true,
+    "forceConsistentCasingInFileNames": true,
+    "skipLibCheck": true,
+    "declaration": true,
+    "
... (truncated)
```

### ✅ 27. powershell — 27607ms

**Input:**

```json
{"command":"npm install --silent \u0026\u0026 npm run check --silent \u0026\u0026 npm test --silent \u0026\u0026 npm run build --silent \u0026\u0026 npm run demo --silent","description":"Install and validate project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
src/in-memory-secret-client.ts(47,11): error TS2375: Type '{ value: string; version: string; expiresOn: Date | undefined; }' is not assignable to type 'StoredVersion' with 'exactOptionalPropertyTypes: true'. Consider adding 'undefined' to the types of the target's properties.
  Types of property 'expiresOn' are incompatible.
    Type 'Date | undefined' is not assignable to type 'Date'.
      Type 'undefined' is not assignable to type 'Date'.
src/in-memory-secret-client.ts(86,7): error TS2375: Type '{ name: string; version: string; expiresOn: Date | undefined; }' is not assignable to type 'SecretPropertiesLike' with 'exactOptionalPropertyTypes: true'. Consider adding 'undefined' to the types of the target's properties.
  Types of property 'expiresOn' are incompatible.
    Type 'Date | undefined' is not assignable to type 'Date'.
      Type 'undefined' is not assignable to type 'Date'.
src/secret-provider.ts(51,7): error TS2375: Type '{ name: string; value: string; version: string | undefined; expiresOn: Date | undefined; found: true; }' is not assignable to type 'SecretValue' with 'exactOptionalPropertyTypes: true'. Consider adding 'undefined' to the types of the target's properties.
  Types of property 'version' are incompatible.
    Type 'string | undefined' is not assignable to type 'string'.
      Type 'undefined' is not assignable to type 'string'.
src/secret-rotation.ts(29,5): error TS2375: Type '{ name: string; version: string | undefined; expiresOn: Date; }' is not assignable to type 'RotationResult' with 'exactOptionalPropertyTypes: true'. Consider adding 'undefined' to the types of the target's properties.
  Types of property 'version' are incompatible.
    Type 'string | undefined' is not assignable to type 'string'.
      Type 'undefined' is not assignable to type 'string'.
<shellId: 1 completed with exit code 1>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\caching-secret-provider.ts`
- `src\in-memory-secret-client.ts`
- `src\key-vault-config.ts`
- `src\main.ts`
- `src\secret-client.ts`
- `src\secret-provider.ts`
- `src\secret-rotation.ts`
- `test\secret-provider.test.ts`
- `tsconfig.json`

## Copilot Response

I’ll build this as a local TypeScript sample only—no Azure resources or deployment commands—and ground the SDK usage before scaffolding.

The Azure reference endpoints timed out, so I’m narrowing the lookup and locating any Key Vault-specific skill guidance before writing files.

The design will use `ManagedIdentityCredential` in Azure, an in-memory client for the offline demo/tests, and a guarded purge API because Key Vault deletes a secret name and **all** versions—not one historical version.

The project is scaffolded; I’m now installing the declared dependencies and exercising type-check, tests, build, and the offline end-to-end demo.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete package manifest with appropriate Azure SDK and TypeScript dependencies.
- Includes a suitable tsconfig.json and reproducible package-lock.json.

## Grader Results

- secret-config.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (12/13)
      - Secret versioning: retrieves specific version via `client.getSecret(name, { version })`: Pass
      - Secret expiry: accesses `properties.expiresOn` on the secret response: Pass
      - Configurable warning window for near-expiry detection (compares expiresOn to current date): Pass
      - In-memory caching (e.g., `Map<string, KeyVaultSecret>`) with bulk-load and single-key refresh: Pass
      - Version-based rotation: calls `client.setSecret(name, newValue, { expiresOn })` to create a new version: Pass
      - Cleanup uses `client.beginDeleteSecret(name)` as a long-running operation: Pass
      - Awaits `poller.pollUntilDone()` before calling `client.purgeDeletedSecret(name)`: Pass
      - Creates new secret only after delete+purge completes (not concurrently): Fail
      - Returns a default value when secret is not found (404), does not crash: Pass
      - Handles RestError with statusCode check for not-found scenarios: Pass
      - NOT using fire-and-forget `deleteSecret()` without waiting for completion: Pass
      - NOT assuming deletion is instantaneous (must use poller): Pass
      - NOT hardcoding vault URL or credentials: Pass
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
| `Criteria from prompt file` | prompt_review | 92% | 1.00 | 0.9231 | 10.3% | ❌ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 11.2% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 8.9231** | **81.1%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-js-ts-secret-config --config js-ts-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
