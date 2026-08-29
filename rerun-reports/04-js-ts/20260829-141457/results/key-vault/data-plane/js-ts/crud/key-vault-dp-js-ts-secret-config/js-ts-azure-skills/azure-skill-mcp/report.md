# Evaluation Report: key-vault-dp-js-ts-secret-config

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 418.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-js-ts-secret-config` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 19/23 |
| Duration | 418.8s |
| Timestamp | 2026-08-29T08:33:47Z |
| Files Generated | 10 |
| Event Count | 15629 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 264.0s |
| Review | 153.5s |
| **Total** | **418.8s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-app-onboard |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=717675 out=13027 |
| Turn Count | 13 |

## Tool Calls

### ✅ 1. skill — 4187ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 10253ms

**Input:**

```json
{"intent":"Create a local TypeScript Node.js application using @azure/keyvault-secrets and @azure/identity with DefaultAzureCredential/managed identity, in-memory caching, secret expiry checks, version retrieval, rotation, and safe delete-and-purge behavior. No Azure resources will be created or modified during development.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 12824ms

**Input:**

```json
{"intent":"Find official Azure SDK for JavaScript/TypeScript documentation for Key Vault SecretClient: getSecret latest and specific version, setSecret with expiresOn, beginDeleteSecret waitUntilDone, purgeDeletedSecret, and DefaultAzureCredential managed identity.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 8971ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. view — 11570ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-app-onboard\\references\\pipeline-rules.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard/references/pipeline-rules.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard/references/pipeline-rules.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard/references/pipeline-rules.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard/references/pipeline-rules.md
@@ -1,89 +1,89 @@
 # Pipeline Rules — Reference
 
 Cross-cutting rules enforced across all workflow steps. Referenced from [SKILL.md](../SKILL.md) `## Pipeline Rules`.
 
 ## Approval gates
 
 ⛔ **Two separate approval gates are required — never merge them.**
 
 1. **Scaffold gate (Step 6):** "✅ Ready to proceed with scaffolding? (Yes / Edit plan / Cancel)" — approves IaC generation only. ⛔ When the plan includes PostgreSQL/MySQL, add `Private access` as a selectable choice — see approval-gates.md for the exact variant.
 2. **Deploy gate (Step 8):** "🚀 Ready to deploy? (Yes / Run manually / Edit plan / Cancel)" — approves resource provisioning.
 
 The scaffold gate does NOT grant deploy permission. After scaffold completes, you MUST present the deploy gate as a SEPARATE response. Never go from scaffold approval directly to `az group create` or `az deployment`.
 
 ⛔ **BOTH gates MUST show Subscription (name + ID), Resource Group, and Region** as standalone lines above the service table — users must see WHERE resources will be created before approving.
 
 ⛔ **NEVER create, write, or modify infrastructure files before the user explicitly says "Yes" to the scaffold gate.** No exceptions — not for simple apps, trivial plans, free-tier deployments, or single-component repos.
 
 ⛔ **Modifying existing Ia
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 5217ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure TypeScript code-generation guidance for a Key Vault-backed configuration provider using managed identity, caching, expiry inspection, secret versioning, rotation, and safe delete/purge.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 5869ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official API guidance for implementing Azure Key Vault secret configuration in TypeScript.","parameters":{"query":"Azure Key Vault JavaScript TypeScript SecretClient getSecret version properties expiresOn setSecret beginDeleteSecret pollUntilDone purgeDeletedSecret DefaultAzureCredential managed identity"}}
```

**Output:**

```
{"results":[{"title":"Delete, restore, or purge a secret in Azure Key Vault with JavaScript","content":"# Delete, restore, or purge a secret in Azure Key Vault with JavaScript\nCreate the [SecretClient](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient) with the appropriate [programmatic authentication credentials](https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-get-started#authorize-access-and-connect-to-key-vault), then use the client to delete an existing secret from Azure Key Vault.\n## Delete a secret\nTo delete a secret in Azure Key Vault, use the [beginDeleteSecret](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient#@azure-keyvault-secrets-secretclient-begindeletesecret) long running operation (LRO) method of the [SecretClient](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient) class, chained with the [pollUntilDone](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/pollerlike#@azure-keyvault-secrets-pollerlike-polluntildone) to wait until the deletion is complete.\nWhen a secret is deleted, it uses the configured [delete strategy](https://learn.microsoft.com/azure/key-vault/general/soft-delete-overview) for the key vault.\nThis \u0060deleteResult\u0060 is a [DeletedSecret](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/deletedsecret) object.","contentUrl":"https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-delete-secret"},{"title":"Delete, restore, or purge a secret in Azure Key Vault with JavaScript","content":"# Delete, restore, or purge a secret in Azure Key Vault with JavaScript\n## Recover a deleted secret\nTo recover a deleted secret in Azure Key Vault, use the [beginRecoverDeletedSecret](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient#@azure-keyvault-secrets-secretclient-beginrecoverdeletedsecret) long 
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 5536ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official TypeScript samples for Azure Key Vault secret retrieval, versioning, expiry, rotation, and safe delete/purge.","parameters":{"language":"typescript","query":"@azure/keyvault-secrets SecretClient getSecret version setSecret expiresOn beginDeleteSecret pollUntilDone purgeDeletedSecret DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"description":"description: Imports necessary classes from Azure Key Vault Secrets and Azure Identity packages, sets up environment variables, and defines helper functions to print secret details.\npackage: @azure/keyvault-secrets, @azure/identity, dotenv/config\nlanguage: typescript\n","codeSnippet":"import {\r\n      SecretClient,\r\n      KeyVaultSecret,\r\n      SecretProperties,\r\n    } from \u0022@azure/keyvault-secrets\u0022;\r\n    import { DefaultAzureCredential } from \u0022@azure/identity\u0022;\r\n    import \u0022dotenv/config\u0022;\r\n    \r\n    // Passwordless credential\r\n    const credential = new DefaultAzureCredential();\r\n    \r\n    // Get Key Vault name from environment variables\r\n    // such as \u0060https://${keyVaultName}.vault.azure.net\u0060\r\n    const keyVaultUrl = process.env.KEY_VAULT_URL;\r\n    if (!keyVaultUrl) throw new Error(\u0022KEY_VAULT_URL is empty\u0022);\r\n    \r\n    function printSecret(secret: KeyVaultSecret): void {\r\n      const { name, value, properties } = secret;\r\n      const { enabled, expiresOn, createdOn } = properties;\r\n      console.log(\u0022Secret: \u0022, { name, value, enabled, expiresOn, createdOn });\r\n    }\r\n    function printSecretProperties(secret: SecretProperties): void {\r\n      const { name, enabled, expiresOn, createdOn } = secret;\r\n      console.log(\u0022Secret: \u0022, { name, enabled, expiresOn, createdOn });\r\n    }\r\n    \r\n    async function main(): Promise\u003Cvoid\u003E {\r\n      // Create a new SecretClient\r\n      const client = new SecretClient(keyVaultUrl, credential);\r\n    \r\n      // Create a unique secret name\r\n      const uniqueString = new Date().getTime().toString();\r\n      const secretName = \u0060secret${uniqueString}\u0060;\r\n    \r\n      // Create a secret\r\n      const createSecretResult = await client.setSecret(\r\n        secretName,\r\n        \u0022MySecretValue\u0022\r\n      );\r\n      printSecret(createSecretResult);\
... (truncated)
```

### ✅ 9. rg — 16974ms

**Input:**

```json
{"head_limit":40,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787992488504-copilot-tool-output-48d733900e1e424eb33bb6d2dbae70a3.txt"],"pattern":"pollUntilDone|beginDeleteSecret|purgeDeletedSecret|getSecret\\(|expiresOn|setSecret\\("}
```

**Output:**

```
Output too large to read at once (38.4 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787992512415-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt:{"results":[{"title":"Delete, restore, or purge a secret in Azure Key Vault with JavaScript","content":"# Delete, restore, or purge a secret in Azure Key Vault with JavaScript\nCreate the [SecretClient](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient) with the appropriate [programmatic authentication credentials](https://learn.microsoft.com/azure/key-vau
```

### ✅ 10. rg — 16653ms

**Input:**

```json
{"head_limit":40,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt"],"pattern":"pollUntilDone|beginDeleteSecret|purgeDeletedSecret|getSecret\\(|expiresOn|setSecret\\("}
```

**Output:**

```
Output too large to read at once (29.8 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787992512414-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787992488504-copilot-tool-output-48d733900e1e424eb33bb6d2dbae70a3.txt:{"results":[{"description":"description: Imports necessary classes from Azure Key Vault Secrets and Azure Identity packages, sets up environment variables, and defines helper functions to print secret details.\npackage: @azure/keyvault-secrets, @azure/identity, dotenv/config\nlanguage: typescript\n","codeSnippet":"import {\r\n      SecretClient,\r\n      KeyVaultSecret,\r\n      SecretProper
```

### ✅ 11. powershell — 17336ms

**Input:**

```json
{"command":"npm view @azure/keyvault-secrets version; npm view @azure/identity version; npm view typescript version; npm view @types/node version","description":"Check current npm package versions","initial_wait":30,"mode":"sync"}
```

**Output:**

```
4.11.2
4.13.2
7.0.2
26.2.0
<shellId: 0 completed with exit code 0>
```

### ✅ 12. powershell — 1396ms

**Input:**

```json
{"command":"$docs = Get-Content 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt' -Raw | ConvertFrom-Json; $docs.results | Select-Object -First 10 title,contentUrl | Format-Table -AutoSize","description":"List relevant Microsoft documentation URLs","initial_wait":30,"mode":"sync"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … d8ba892319654906b0ff4172b65431a8.txt' -Raw | ConvertFrom-Json; $docs. …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 13. view — 630ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt b/C:/Users/weidxu/AppData/Local/Temp/1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787992488806-copilot-tool-output-d8ba892319654906b0ff4172b65431a8.txt
@@ -1,8 +1,8 @@
 {"results":[{"title":"Delete, restore, or purge a secret in Azure Key Vault with JavaScript","content":"# Delete, restore, or purge a secret in Azure Key Vault with JavaScript\nCreate the [SecretClient](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient) with the appropriate [programmatic authentication credentials](https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-get-started#authorize-access-and-connect-to-key-vault), then use the client to delete an existing secret from Azure Key Vault.\n## Delete a secret\nTo delete a secret in Azure Key Vault, use the [beginDeleteSecret](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient#@azure-keyvault-secrets-secretclient-begindeletesecret) long running operation (LRO) method of the [SecretClient](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/secretclient) class, chained with the [pollUntilDone](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/pollerlike#@azure-keyvault-secrets-pollerlike-polluntildone) to wait until the deletion is complete.\nWhen a secret is deleted, it uses the configured [delete strategy](https://learn.microsoft.com/azure/key-vault/general/soft-delete-overview) for the key vault.\nThis \u0060deleteResult\u0060 is a [DeletedSecret](https://learn.microsoft.com/en-us/javascript/api/@azure/keyvault-secrets/deletedsecret) object.","contentUrl":"https://learn.microsoft.com/a
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 1923ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete official guidance for safe Key Vault secret deletion and purge in JavaScript.","parameters":{"url":"https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-delete-secret"}}
```

**Output:**

```
# Quickstart: Azure Key Vault secret client library for JavaScript

Get started with the Azure Key Vault secret client library for JavaScript. [Azure Key Vault](../general/overview) is a cloud service that provides a secure store for secrets. You can securely store keys, passwords, certificates, and other secrets. Azure key vaults may be created and managed through the Azure portal. In this quickstart, you learn how to create, retrieve, and delete secrets from an Azure key vault using the JavaScript client library.

Key Vault client library resources:

[API reference documentation](/en-us/javascript/api/overview/azure/key-vault) | [Library source code](https://github.com/Azure/azure-sdk-for-js/tree/master/sdk/keyvault) | [Package (npm)](https://www.npmjs.com/package/@azure/keyvault-secrets)

For more information about Key Vault and secrets, see:

- [Key Vault Overview](../general/overview)
- [Secrets Overview](about-secrets)

::: zone pivot="programming-language-javascript"

## Prerequisites

- An Azure subscription - [create one for free](https://azure.microsoft.com/pricing/purchase-options/azure-account?cid=msft_learn).
- Current [Node.js LTS](https://nodejs.org).
- [Azure CLI](/en-us/cli/azure/install-azure-cli)

::: zone-end

::: zone pivot="programming-language-typescript"

## Prerequisites

- An Azure subscription - [create one for free](https://azure.microsoft.com/pricing/purchase-options/azure-account?cid=msft_learn).
- Current [Node.js LTS](https://nodejs.org).
- [TypeScript 5+](https://www.typescriptlang.org/download/)
- [Azure CLI](/en-us/cli/azure/install-azure-cli).

::: zone-end

This quickstart assumes you are running [Azure CLI](/en-us/cli/azure/install-azure-cli).

## Sign in to Azure

1. Run the `login` command.

    ```azurecli
    az login
    ```

    If the CLI can open your default browser, it will do so and load an Azure sign-in page.

    Otherwise, open a browser page at https://aka.ms/devicelog
... (truncated)
```

### ✅ 15. azure-documentation (via azure) — 1600ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete official guidance for Key Vault SecretClient passwordless authentication and secret operations in JavaScript.","parameters":{"url":"https://learn.microsoft.com/azure/key-vault/secrets/quick-create-node"}}
```

**Output:**

```
# Delete, restore, or purge a secret in Azure Key Vault with JavaScript

Create the [SecretClient](/en-us/javascript/api/@azure/keyvault-secrets/secretclient) with the appropriate [programmatic authentication credentials](javascript-developer-guide-get-started#authorize-access-and-connect-to-key-vault), then use the client to delete an existing secret from Azure Key Vault.

## Delete a secret

To delete a secret in Azure Key Vault, use the [beginDeleteSecret](/en-us/javascript/api/@azure/keyvault-secrets/secretclient#@azure-keyvault-secrets-secretclient-begindeletesecret) long running operation (LRO) method of the [SecretClient](/en-us/javascript/api/@azure/keyvault-secrets/secretclient) class, chained with the [pollUntilDone](/en-us/javascript/api/@azure/keyvault-secrets/pollerlike#@azure-keyvault-secrets-pollerlike-polluntildone) to wait until the deletion is complete.

When a secret is deleted, it uses the configured [delete strategy](../general/soft-delete-overview) for the key vault.

```javascript
const existingSecretName = 'myExistingSecret';

// Begin LRO
const deletePoller = await client.beginDeleteSecret(secretName);

// Wait for LRO to complete
const deleteResult = await deletePoller.pollUntilDone();

console.log(`SecretName: ${deleteResult.name}`);
console.log(`DeletedDate: ${deleteResult.deletedOn}`);
console.log(`Version: ${deleteResult.properties.deletedOn}`);
console.log(`PurgeDate: ${deleteResult.scheduledPurgeDate}`);
```

This `deleteResult` is a [DeletedSecret](/en-us/javascript/api/@azure/keyvault-secrets/deletedsecret) object.

## Recover a deleted secret

To recover a deleted secret in Azure Key Vault, use the [beginRecoverDeletedSecret](/en-us/javascript/api/@azure/keyvault-secrets/secretclient#@azure-keyvault-secrets-secretclient-beginrecoverdeletedsecret) long running operation (LRO) method of the SecretClient class, chained with the [pollUntilDone](/en-us/javascript/api/@azure/keyvault-secrets/pollerlike#@azure
... (truncated)
```

### ✅ 16. apply_patch — 655ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-key-vault-config-provider\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"description\": \"A managed-identity Azure Key Vault configuration provider with caching and rotation support.\",\n+  \"type\": \"module\",\n+  \"engines\": {\n+    \"node\": \"\u003e=20\"\n+  },\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"clean\": \"node -e \\\"require('node:fs').rmSync('dist', { recursive: true, force: true })\\\"\",\n+    \"start\": \"node dist\\\\src\\\\main.js\",\n+    \"dev\": \"tsx src\\\\main.ts\",\n+    \"test\": \"npm run build \u0026\u0026 node --test dist\\\\test\\\\secret-config.test.js\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/identity\": \"^4.13.2\",\n+    \"@azure/keyvault-secrets\": \"^4.11.2\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"^26.2.0\",\n+    \"tsx\": \"^4.20.6\",\n+    \"typescript\": \"^7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \".\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"noImplicitOverride\": true,\n+    \"esModuleInterop\": true,\n+    \"forceConsistentCasingInFileNames\": true,\n+    \"skipLibCheck\": true,\n+    \"sourceMap\": true,\n+    \"declaration\": true\n+  },\n+  \"include\": [\n+    \"src/**/*.ts\",\n+    \"test/**/*.ts\"\n+  ]\n+}\n*** Add File: .gitignore\n+node_modules/\n+dist/\n+.env\n+*.tsbuildinfo\n*** Add File: .env.example\n+KEY_VAULT_URL=https://your-vault-name.vault.azure.net\n+AZURE_CLIENT_ID=\n+CONFIG_SECRET_NAMES=DatabaseConnectionString,ExternalApiKey,FeatureFlag\n+EXPIRY_WARNING_DAYS=7\n+AUTO_REFRESH_INTERVAL_MINUTES=60\n+\n+# Optional version retrieval demo\n+VERSIONED_SECRET_NAME=\n+SECRET_VERSION=\n+\n+# Rotation is disabled unless explicitly enabled.\n+ENABLE_ROTATION_DEMO=false\n+ROTATE_SECRET_NAME=\n+ROTATE_SECRET_VALUE=\n+ROTATE_SECRET_EXPIRY_DAYS=90\n+\n+# Purge is irreversible and deletes every version under the secret name.\n+ENABLE_PURGE_DEMO=false\n+PURGE_SECRET_NAME=\n*** Add File: src/secret-provider.ts\n+import type { GetSecretOptions, KeyVaultSecret } from \"@azure/keyvault-secrets\";\n+\n+export interface SecretReader {\n+  getSecret(name: string, options?: GetSecretOptions): Promise\u003cKeyVaultSecret\u003e;\n+}\n+\n+export interface SecretRequest {\n+  defaultValue: string;\n+  version?: string;\n+}\n+\n+export interface ResolvedSecret {\n+  name: string;\n+  value: string;\n+  version?: string;\n+  expiresOn?: Date;\n+  found: boolean;\n+}\n+\n+function isSecretNotFound(error: unknown): boolean {\n+  if (typeof error !== \"object\" || error === null) {\n+    return false;\n+  }\n+\n+  const candidate = error as { statusCode?: unknown; code?: unknown };\n+  return candidate.statusCode === 404 || candidate.code === \"SecretNotFound\";\n+}\n+\n+export class KeyVaultSecretProvider {\n+  public constructor(private readonly client: SecretReader) {}\n+\n+  public async getSecret(name: string, request: SecretRequest): Promise\u003cResolvedSecret\u003e {\n+    if (!name.trim()) {\n+      throw new Error(\"Secret name must not be empty.\");\n+    }\n+\n+    try {\n+      const options = request.version ? { version: request.version } : undefined;\n+      const secret = await this.client.getSecret(name, options);\n+\n+      if (secret.value === undefined) {\n+        throw new Error(`Key Vault returned secret \"${name}\" without a value.`);\n+      }\n+\n+      return {\n+        name: secret.name,\n+        value: secret.value,\n+        version: secret.properties.version,\n+        expiresOn: secret.properties.expiresOn,\n+        found: true\n+      };\n+    } catch (error: unknown) {\n+      if (!isSecretNotFound(error)) {\n+        throw error;\n+      }\n+\n+      return {\n+        name,\n+        value: request.defaultValue,\n+        version: request.version,\n+        found: false\n+      };\n+    }\n+  }\n+\n+  public isExpiringSoon(\n+    secret: Pick\u003cResolvedSecret, \"expiresOn\"\u003e,\n+    warningWindowMs: number,\n+    now = new Date()\n+  ): boolean {\n+    if (warningWindowMs \u003c 0) {\n+      throw new Error(\"Expiry warning window must not be negative.\");\n+    }\n+\n+    return secret.expiresOn !== undefined\n+      \u0026\u0026 secret.expiresOn.getTime() - now.getTime() \u003c= warningWindowMs;\n+  }\n+}\n*** Add File: src/secret-cache.ts\n+import {\n+  KeyVaultSecretProvider,\n+  type ResolvedSecret\n+} from \"./secret-provider.js\";\n+\n+export interface RequiredConfigKey {\n+  name: string;\n+  defaultValue: string;\n+}\n+\n+export interface CachedSecret extends ResolvedSecret {\n+  cachedAt: Date;\n+  defaultValue: string;\n+}\n+\n+export type AutoRefreshErrorHandler = (error: unknown) =\u003e void;\n+\n+export class SecretCache {\n+  private readonly entries = new Map\u003cstring, CachedSecret\u003e();\n+  private readonly inFlight = new Map\u003cstring, Promise\u003cCachedSecret\u003e\u003e();\n+\n+  public constructor(\n+    private readonly provider: KeyVaultSecretProvider,\n+    private readonly expiryWarningWindowMs: number\n+  ) {\n+    if (expiryWarningWindowMs \u003c 0) {\n+      throw new Error(\"Expiry warning window must not be negative.\");\n+    }\n+  }\n+\n+  public async bulkLoad(keys: readonly RequiredConfigKey[]): Promise\u003cReadonlyMap\u003cstring, CachedSecret\u003e\u003e {\n+    const uniqueKeys = new Map(keys.map((key) =\u003e [key.name, key]));\n+    await Promise.all(\n+      [...uniqueKeys.values()].map((key) =\u003e this.fetchAndCache(key.name, key.defaultValue))\n+    );\n+    return this.snapshot();\n+  }\n+\n+  public async get(name: string, defaultValue = \"\"): Promise\u003cstring\u003e {\n+    const cached = this.entries.get(name);\n+    if (!cached) {\n+      return (await this.fetchAndCache(name, defaultValue)).value;\n+    }\n+\n+    if (this.provider.isExpiringSoon(cached, this.expiryWarningWindowMs)) {\n+      return (await this.fetchAndCache(name, cached.defaultValue)).value;\n+    }\n+\n+    return cached.value;\n+  }\n+\n+  public async refresh(name: string): Promise\u003cCachedSecret\u003e {\n+    const cached = this.entries.get(name);\n+    if (!cached) {\n+      throw new Error(`Cannot refresh \"${name}\" before it has been loaded into the cache.`);\n+    }\n+\n+    return this.fetchAndCache(name, cached.defaultValue);\n+  }\n+\n+  public async refreshExpiring(): Promise\u003creadonly CachedSecret[]\u003e {\n+    const expiring = this.getExpiringSecrets();\n+    return Promise.all(\n+      expiring.map((entry) =\u003e this.fetchAndCache(entry.name, entry.defaultValue))\n+    );\n+  }\n+\n+  public getExpiringSecrets(now = new Date()): readonly CachedSecret[] {\n+    return [...this.entries.values()].filter((entry) =\u003e\n+      this.provider.isExpiringSoon(entry, this.expiryWarningWindowMs, now)\n+    );\n+  }\n+\n+  public getEntry(name: string): CachedSecret | undefined {\n+    return this.entries.get(name);\n+  }\n+\n+  public snapshot(): ReadonlyMap\u003cstring, CachedSecret\u003e {\n+    return new Map(this.entries);\n+  }\n+\n+  public startAutoRefresh(\n+    intervalMs: number,\n+    onError: AutoRefreshErrorHandler\n+  ): () =\u003e void {\n+    if (intervalMs \u003c= 0) {\n+      throw new Error(\"Automatic refresh interval must be greater than zero.\");\n+    }\n+\n+    const timer = setInterval(() =\u003e {\n+      void this.refreshExpiring().catch(onError);\n+    }, intervalMs);\n+    timer.unref();\n+\n+    return () =\u003e clearInterval(timer);\n+  }\n+\n+  private fetchAndCache(name: string, defaultValue: string): Promise\u003cCachedSecret\u003e {\n+    const existingRequest = this.inFlight.get(name);\n+    if (existingRequest) {\n+      return existingRequest;\n+    }\n+\n+    const request = this.provider\n+      .getSecret(name, { defaultValue })\n+      .then((secret): CachedSecret =\u003e {\n+        const entry = {\n+          ...secret,\n+          defaultValue,\n+          cachedAt: new Date()\n+        };\n+        this.entries.set(name, entry);\n+        return entry;\n+      })\n+      .finally(() =\u003e {\n+        this.inFlight.delete(name);\n+      });\n+\n+    this.inFlight.set(name, request);\n+    return request;\n+  }\n+}\n*** Add File: src/configuration.ts\n+import { ManagedIdentityCredential } from \"@azure/identity\";\n+import { SecretClient } from \"@azure/keyvault-secrets\";\n+import { SecretCache } from \"./secret-cache.js\";\n+import { KeyVaultSecretProvider } from \"./secret-provider.js\";\n+\n+const DAY_MS = 24 * 60 * 60 * 1000;\n+\n+function requireHttpsUrl(value: string, variableName: string): string {\n+  let url: URL;\n+  try {\n+    url = new URL(value);\n+  } catch {\n+    throw new Error(`${variableName} must be a valid URL.`);\n+  }\n+\n+  if (url.protocol !== \"https:\") {\n+    throw new Error(`${variableName} must use HTTPS.`);\n+  }\n+\n+  if (url.username || url.password) {\n+    throw new Error(`${variableName} must not contain credentials.`);\n+  }\n+\n+  return url.toString().replace(/\\/$/, \"\");\n+}\n+\n+function readNonNegativeNumber(value: string | undefined, fallback: number, name: string): number {\n+  if (value === undefined || value.trim() === \"\") {\n+    return fallback;\n+  }\n+\n+  const number = Number(value);\n+  if (!Number.isFinite(number) || number \u003c 0) {\n+    throw new Error(`${name} must be a non-negative number.`);\n+  }\n+  return number;\n+}\n+\n+export interface ApplicationConfiguration {\n+  client: SecretClient;\n+  provider: KeyVaultSecretProvider;\n+  cache: SecretCache;\n+  expiryWarningWindowMs: number;\n+  autoRefreshIntervalMs: number;\n+}\n+\n+export function createApplicationConfiguration(\n+  env: NodeJS.ProcessEnv = process.env\n+): ApplicationConfiguration {\n+  const rawVaultUrl = env.KEY_VAULT_URL;\n+  if (!rawVaultUrl) {\n+    throw new Error(\"KEY_VAULT_URL is required.\");\n+  }\n+\n+  const vaultUrl = requireHttpsUrl(rawVaultUrl, \"KEY_VAULT_URL\");\n+  const credential = env.AZURE_CLIENT_ID\n+    ? new ManagedIdentityCredential(env.AZURE_CLIENT_ID)\n+    : new ManagedIdentityCredential();\n+  const client = new SecretClient(vaultUrl, credential, {\n+    retryOptions: {\n+      maxRetries: 5,\n+      retryDelayInMs: 1_000,\n+      maxRetryDelayInMs: 10_000\n+    }\n+  });\n+  const provider = new KeyVaultSecretProvider(client);\n+  const expiryWarningWindowMs =\n+    readNonNegativeNumber(env.EXPIRY_WARNING_DAYS, 7, \"EXPIRY_WARNING_DAYS\") * DAY_MS;\n+  const autoRefreshIntervalMs =\n+    readNonNegativeNumber(\n+      env.AUTO_REFRESH_INTERVAL_MINUTES,\n+      60,\n+      \"AUTO_REFRESH_INTERVAL_MINUTES\"\n+    ) * 60 * 1000;\n+\n+  if (autoRefreshIntervalMs === 0) {\n+    throw new Error(\"AUTO_REFRESH_INTERVAL_MINUTES must be greater than zero.\");\n+  }\n+\n+  return {\n+    client,\n+    provider,\n+    cache: new SecretCache(provider, expiryWarningWindowMs),\n+    expiryWarningWindowMs,\n+    autoRefreshIntervalMs\n+  };\n+}\n*** Add File: src/secret-rotation.ts\n+import type {\n+  DeleteSecretPoller,\n+  KeyVaultSecret,\n+  SetSecretOptions\n+} from \"@azure/keyvault-secrets\";\n+\n+export interface SecretWriter {\n+  setSecret(name: string, value: string, options?: SetSecretOptions): Promise\u003cKeyVaultSecret\u003e;\n+  beginDeleteSecret(name: string): Promise\u003cDeleteSecretPoller\u003e;\n+  purgeDeletedSecret(name: string): Promise\u003cvoid\u003e;\n+}\n+\n+export interface RotationResult {\n+  name: string;\n+  version?: string;\n+  expiresOn: Date;\n+}\n+\n+export class SecretRotationHelper {\n+  public constructor(private readonly client: SecretWriter) {}\n+\n+  public async rotateSecret(\n+    name: string,\n+    value: string,\n+    expiresOn: Date\n+  ): Promise\u003cRotationResult\u003e {\n+    if (!name.trim()) {\n+      throw new Error(\"Secret name must not be empty.\");\n+    }\n+    if (!value) {\n+      throw new Error(\"Rotated secret value must not be empty.\");\n+    }\n+    if (expiresOn.getTime() \u003c= Date.now()) {\n+      throw new Error(\"Rotated secret expiry must be in the future.\");\n+    }\n+\n+    const created = await this.client.setSecret(name, value, { expiresOn });\n+    return {\n+      name: created.name,\n+      version: created.properties.version,\n+      expiresOn\n+    };\n+  }\n+\n+  public async deleteAndPurgeSecret(\n+    name: string,\n+    confirmPermanentPurge: boolean\n+  ): Promise\u003cvoid\u003e {\n+    if (!confirmPermanentPurge) {\n+      throw new Error(\n+        \"Permanent purge was not confirmed. Deletion would remove every version of this secret.\"\n+      );\n+    }\n+\n+    const deletePoller = await this.client.beginDeleteSecret(name);\n+    await deletePoller.pollUntilDone();\n+    await this.client.purgeDeletedSecret(name);\n+  }\n+}\n*** Add File: src/main.ts\n+import { createApplicationConfiguration } from \"./configuration.js\";\n+import { type RequiredConfigKey } from \"./secret-cache.js\";\n+import { SecretRotationHelper } from \"./secret-rotation.js\";\n+\n+const DAY_MS = 24 * 60 * 60 * 1000;\n+\n+function parseSecretNames(value: string | undefined): readonly string[] {\n+  const configured = value ?? \"DatabaseConnectionString,ExternalApiKey,FeatureFlag\";\n+  const names = configured\n+    .split(\",\")\n+    .map((name) =\u003e name.trim())\n+    .filter(Boolean);\n+\n+  if (names.length === 0) {\n+    throw new Error(\"CONFIG_SECRET_NAMES must contain at least one secret name.\");\n+  }\n+  return names;\n+}\n+\n+function requirePositiveNumber(value: string | undefined, fallback: number, name: string): number {\n+  const number = value === undefined ? fallback : Number(value);\n+  if (!Number.isFinite(number) || number \u003c= 0) {\n+    throw new Error(`${name} must be a positive number.`);\n+  }\n+  return number;\n+}\n+\n+function isEnabled(value: string | undefined): boolean {\n+  return value?.toLowerCase() === \"true\";\n+}\n+\n+function redact(value: string): string {\n+  return value.length === 0 ? \"\u003cempty\u003e\" : `\u003credacted:${value.length} chars\u003e`;\n+}\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const configuration = createApplicationConfiguration();\n+  const { cache, provider, client } = configuration;\n+  const requiredKeys: readonly RequiredConfigKey[] = parseSecretNames(\n+    process.env.CONFIG_SECRET_NAMES\n+  ).map((name) =\u003e ({ name, defaultValue: `\u003cdefault:${name}\u003e` }));\n+\n+  console.log(\"1. Bulk-loading required configuration keys...\");\n+  const loaded = await cache.bulkLoad(requiredKeys);\n+  for (const [name, entry] of loaded) {\n+    console.log(\n+      `   ${name}: ${redact(entry.value)}, source=${entry.found ? \"Key Vault\" : \"default\"}, `\n+      + `version=${entry.version ?? \"none\"}, expires=${entry.expiresOn?.toISOString() ?? \"never\"}`\n+    );\n+  }\n+\n+  console.log(\"\\n2. Reading values from the in-memory cache...\");\n+  for (const key of requiredKeys) {\n+    console.log(`   ${key.name}: ${redact(await cache.get(key.name))}`);\n+  }\n+\n+  const refreshTarget = requiredKeys[0];\n+  if (!refreshTarget) {\n+    throw new Error(\"No required configuration keys were configured.\");\n+  }\n+  console.log(`\\n3. Refreshing \"${refreshTarget.name}\" on demand...`);\n+  const refreshed = await cache.refresh(refreshTarget.name);\n+  console.log(\n+    `   Refreshed version=${refreshed.version ?? \"none\"}, cachedAt=${refreshed.cachedAt.toISOString()}`\n+  );\n+\n+  console.log(\"\\n4. Checking for secrets near expiry...\");\n+  const expiring = cache.getExpiringSecrets();\n+  if (expiring.length === 0) {\n+    console.log(\"   No cached secrets are inside the expiry warning window.\");\n+  } else {\n+    for (const entry of expiring) {\n+      console.warn(`   WARNING: ${entry.name} expires at ${entry.expiresOn?.toISOString()}.`);\n+    }\n+    const autoRefreshed = await cache.refreshExpiring();\n+    console.log(`   Automatically re-fetched ${autoRefreshed.length} expiring secret(s).`);\n+  }\n+\n+  const stopAutoRefresh = cache.startAutoRefresh(\n+    configuration.autoRefreshIntervalMs,\n+    (error) =\u003e console.error(\"Automatic secret refresh failed:\", error)\n+  );\n+  console.log(\n+    `   Periodic expiry refresh enabled every ${configuration.autoRefreshIntervalMs / 60_000} minute(s).`\n+  );\n+\n+  const versionedName = process.env.VERSIONED_SECRET_NAME;\n+  const version = process.env.SECRET_VERSION;\n+  if (versionedName \u0026\u0026 version) {\n+    console.log(`\\n5. Retrieving version \"${version}\" of \"${versionedName}\"...`);\n+    const versioned = await provider.getSecret(versionedName, {\n+      version,\n+      defaultValue: `\u003cdefault:${versionedName}\u003e`\n+    });\n+    console.log(\n+      `   ${versioned.name}: ${redact(versioned.value)}, found=${versioned.found}, `\n+      + `expires=${versioned.expiresOn?.toISOString() ?? \"never\"}`\n+    );\n+  } else {\n+    console.log(\"\\n5. Version retrieval skipped; set VERSIONED_SECRET_NAME and SECRET_VERSION.\");\n+  }\n+\n+  const rotation = new SecretRotationHelper(client);\n+  if (isEnabled(process.env.ENABLE_ROTATION_DEMO)) {\n+    const name = process.env.ROTATE_SECRET_NAME;\n+    const value = process.env.ROTATE_SECRET_VALUE;\n+    if (!name || !value) {\n+      throw new Error(\n+        \"ROTATE_SECRET_NAME and ROTATE_SECRET_VALUE are required when rotation is enabled.\"\n+      );\n+    }\n+\n+    const expiryDays = requirePositiveNumber(\n+      process.env.ROTATE_SECRET_EXPIRY_DAYS,\n+      90,\n+      \"ROTATE_SECRET_EXPIRY_DAYS\"\n+    );\n+    const expiresOn = new Date(Date.now() + expiryDays * DAY_MS);\n+    console.log(`\\n6. Rotating \"${name}\" by creating a new version...`);\n+    const result = await rotation.rotateSecret(name, value, expiresOn);\n+    console.log(\n+      `   Created version=${result.version ?? \"unknown\"}, expires=${result.expiresOn.toISOString()}.`\n+    );\n+  } else {\n+    console.log(\"\\n6. Rotation skipped; set ENABLE_ROTATION_DEMO=true to run it.\");\n+  }\n+\n+  if (isEnabled(process.env.ENABLE_PURGE_DEMO)) {\n+    const name = process.env.PURGE_SECRET_NAME;\n+    if (!name) {\n+      throw new Error(\"PURGE_SECRET_NAME is required when purge is enabled.\");\n+    }\n+\n+    console.log(\n+      `\\n7. Deleting and permanently purging \"${name}\" (all versions) after the delete LRO completes...`\n+    );\n+    await rotation.deleteAndPurgeSecret(name, true);\n+    console.log(\"   Delete completed and the deleted secret was purged.\");\n+  } else {\n+    console.log(\n+      \"\\n7. Delete/purge skipped; set ENABLE_PURGE_DEMO=true and PURGE_SECRET_NAME to run it.\"\n+    );\n+  }\n+\n+  stopAutoRefresh();\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Application failed:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: test/secret-config.test.ts\n+import assert from \"node:assert/strict\";\n+import { describe, it } from \"node:test\";\n+import type {\n+  DeleteSecretPoller,\n+  GetSecretOptions,\n+  KeyVaultSecret,\n+  SetSecretOptions\n+} from \"@azure/keyvault-secrets\";\n+import { SecretCache } from \"../src/secret-cache.js\";\n+import {\n+  KeyVaultSecretProvider,\n+  type SecretReader\n+} from \"../src/secret-provider.js\";\n+import {\n+  SecretRotationHelper,\n+  type SecretWriter\n+} from \"../src/secret-rotation.js\";\n+\n+function secret(name: string, value: string, expiresOn?: Date, version = \"v1\"): KeyVaultSecret {\n+  return {\n+    name,\n+    value,\n+    properties: {\n+      name,\n+      vaultUrl: \"https://example.vault.azure.net\",\n+      id: `https://example.vault.azure.net/secrets/${name}/${version}`,\n+      version,\n+      expiresOn\n+    }\n+  };\n+}\n+\n+describe(\"KeyVaultSecretProvider\", () =\u003e {\n+  it(\"returns the default only for a missing secret\", async () =\u003e {\n+    const reader: SecretReader = {\n+      getSecret: async () =\u003e {\n+        throw { statusCode: 404 };\n+      }\n+    };\n+    const provider = new KeyVaultSecretProvider(reader);\n+\n+    const result = await provider.getSecret(\"missing\", { defaultValue: \"fallback\" });\n+\n+    assert.deepEqual(result, {\n+      name: \"missing\",\n+      value: \"fallback\",\n+      version: undefined,\n+      found: false\n+    });\n+  });\n+\n+  it(\"retrieves a requested secret version\", async () =\u003e {\n+    let receivedOptions: GetSecretOptions | undefined;\n+    const reader: SecretReader = {\n+      getSecret: async (name, options) =\u003e {\n+        receivedOptions = options;\n+        return secret(name, \"old-value\", undefined, \"v2\");\n+      }\n+    };\n+    const provider = new KeyVaultSecretProvider(reader);\n+\n+    const result = await provider.getSecret(\"api-key\", {\n+      defaultValue: \"fallback\",\n+      version: \"v2\"\n+    });\n+\n+    assert.equal(receivedOptions?.version, \"v2\");\n+    assert.equal(result.value, \"old-value\");\n+    assert.equal(result.version, \"v2\");\n+  });\n+\n+  it(\"does not hide non-404 service failures\", async () =\u003e {\n+    const failure = { statusCode: 403 };\n+    const provider = new KeyVaultSecretProvider({\n+      getSecret: async () =\u003e {\n+        throw failure;\n+      }\n+    });\n+\n+    await assert.rejects(\n+      provider.getSecret(\"forbidden\", { defaultValue: \"fallback\" }),\n+      (error) =\u003e error === failure\n+    );\n+  });\n+});\n+\n+describe(\"SecretCache\", () =\u003e {\n+  it(\"bulk-loads, caches, and refreshes individual values\", async () =\u003e {\n+    let calls = 0;\n+    const provider = new KeyVaultSecretProvider({\n+      getSecret: async (name) =\u003e secret(name, `value-${++calls}`)\n+    });\n+    const cache = new SecretCache(provider, 7 * 24 * 60 * 60 * 1000);\n+\n+    await cache.bulkLoad([{ name: \"api-key\", defaultValue: \"fallback\" }]);\n+    assert.equal(await cache.get(\"api-key\"), \"value-1\");\n+    assert.equal(calls, 1);\n+\n+    assert.equal((await cache.refresh(\"api-key\")).value, \"value-2\");\n+    assert.equal(calls, 2);\n+  });\n+\n+  it(\"automatically re-fetches a cached secret near expiry\", async () =\u003e {\n+    let calls = 0;\n+    const provider = new KeyVaultSecretProvider({\n+      getSecret: async (name) =\u003e {\n+        calls += 1;\n+        return calls === 1\n+          ? secret(name, \"expiring\", new Date(Date.now() + 60_000))\n+          : secret(name, \"rotated\", new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), \"v2\");\n+      }\n+    });\n+    const cache = new SecretCache(provider, 7 * 24 * 60 * 60 * 1000);\n+\n+    await cache.bulkLoad([{ name: \"api-key\", defaultValue: \"fallback\" }]);\n+\n+    assert.equal(await cache.get(\"api-key\"), \"rotated\");\n+    assert.equal(calls, 2);\n+  });\n+});\n+\n+describe(\"SecretRotationHelper\", () =\u003e {\n+  it(\"creates a new version with expiry\", async () =\u003e {\n+    let options: SetSecretOptions | undefined;\n+    const writer = {\n+      setSecret: async (name: string, value: string, received?: SetSecretOptions) =\u003e {\n+        options = received;\n+        return secret(name, value, received?.expiresOn, \"new-version\");\n+      },\n+      beginDeleteSecret: async () =\u003e {\n+        throw new Error(\"not used\");\n+      },\n+      purgeDeletedSecret: async () =\u003e undefined\n+    } satisfies SecretWriter;\n+    const helper = new SecretRotationHelper(writer);\n+    const expiresOn = new Date(Date.now() + 86_400_000);\n+\n+    const result = await helper.rotateSecret(\"api-key\", \"new-value\", expiresOn);\n+\n+    assert.equal(options?.expiresOn, expiresOn);\n+    assert.equal(result.version, \"new-version\");\n+  });\n+\n+  it(\"waits for deletion before purge\", async () =\u003e {\n+    const events: string[] = [];\n+    const poller = {\n+      pollUntilDone: async () =\u003e {\n+        events.push(\"delete-complete\");\n+        return {};\n+      }\n+    } as unknown as DeleteSecretPoller;\n+    const writer: SecretWriter = {\n+      setSecret: async (name, value) =\u003e secret(name, value),\n+      beginDeleteSecret: async () =\u003e {\n+        events.push(\"delete-start\");\n+        return poller;\n+      },\n+      purgeDeletedSecret: async () =\u003e {\n+        events.push(\"purge\");\n+      }\n+    };\n+    const helper = new SecretRotationHelper(writer);\n+\n+    await helper.deleteAndPurgeSecret(\"disposable-secret\", true);\n+\n+    assert.deepEqual(events, [\"delete-start\", \"delete-complete\", \"purge\"]);\n+  });\n+\n+  it(\"requires explicit confirmation before permanent purge\", async () =\u003e {\n+    const writer = {\n+      setSecret: async (name: string, value: string) =\u003e secret(name, value),\n+      beginDeleteSecret: async () =\u003e {\n+        throw new Error(\"must not be called\");\n+      },\n+      purgeDeletedSecret: async () =\u003e {\n+        throw new Error(\"must not be called\");\n+      }\n+    } satisfies SecretWriter;\n+\n+    await assert.rejects(\n+      new SecretRotationHelper(writer).deleteAndPurgeSecret(\"api-key\", false),\n+      /not confirmed/\n+    );\n+  });\n+});\n*** Add File: README.md\n+# Azure Key Vault configuration provider\n+\n+TypeScript/Node.js configuration provider backed by Azure Key Vault. It uses managed identity only, caches values in memory, refreshes near-expiry secrets, supports version-specific reads, and provides explicit rotation and delete/purge operations.\n+\n+## Run\n+\n+```powershell\n+npm install\n+npm test\n+npm run build\n+$env:KEY_VAULT_URL = \"https://your-vault.vault.azure.net\"\n+npm start\n+```\n+\n+The Azure host's managed identity needs data-plane permissions to read secrets. Rotation additionally needs `set`; cleanup needs `delete` and `purge`. For a user-assigned managed identity, set `AZURE_CLIENT_ID` to its client ID.\n+\n+Copy `.env.example` values into the host environment. The program does not load `.env` files or accept client secrets. Secret values are redacted in logs.\n+\n+## Safety\n+\n+`setSecret` creates a new version when a secret name already exists. Azure Key Vault cannot delete just one secret version: deleting a secret name soft-deletes every version under that name. For that reason, delete/purge is separate from rotation, waits for `beginDeleteSecret(...).pollUntilDone()`, and requires both `ENABLE_PURGE_DEMO=true` and a `PURGE_SECRET_NAME`. Purge is irreversible and can be blocked by vault purge protection.\n+\n+## References\n+\n+- [Azure Key Vault JavaScript quickstart](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-node)\n+- [Delete, recover, and purge a secret](https://learn.microsoft.com/azure/key-vault/secrets/javascript-developer-guide-delete-secret)\n+- [Managed identities for Azure resources](https://learn.microsoft.com/entra/identity/managed-identities-azure-resources/overview)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/package.json
@@ -1,0 +1,27 @@
+{
+  "name": "azure-key-vault-config-provider",
+  "version": "1.0.0",
+  "private": true,
+  "description": "A managed-identity Azure Key Vault configuration provider with caching and rotation support.",
+  "type": "module",
+  "engines": {
+    "node": ">=20"
+  },
+  "scripts": {
+    "build": "tsc",
+    "clean": "node -e \"require('node:fs').rmSync('dist', { recursive: true, force: true })\"",
+    "start": "node dist\\src\\main.js",
+    "dev": "tsx src\\main.ts",
+    "test": "npm run build && node --test dist\\test\\secret-config.test.js"
+  },
+  "dependencies": {
+    "@azure/identity": "^4.13.2",
+    "@azure/keyvault-secrets": "^4.11.2"
+  },
+  "devDependencies": {
+    "@types/node": "^26.2.0",
+    "tsx": "^4.20.6",
+    "typescript": "^7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/tsconfig.json
@@ -1,0 +1,22 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "rootDir": ".",
+    "outDir": "dist",
+    "strict": true,
+    "noUncheckedIndexedAccess": t
... (truncated)
```

### ✅ 17. powershell — 23783ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm test","description":"Install dependencies and run project tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 55 packages in 21s

> azure-key-vault-config-provider@1.0.0 test
> npm run build && node --test dist\test\secret-config.test.js


> azure-key-vault-config-provider@1.0.0 build
> tsc

src/secret-rotation.ts(2,3): error TS2724: '"@azure/keyvault-secrets"' has no exported member named 'DeleteSecretPoller'. Did you mean 'DeletedSecret'?
test/secret-config.test.ts(4,3): error TS2724: '"@azure/keyvault-secrets"' has no exported member named 'DeleteSecretPoller'. Did you mean 'DeletedSecret'?
<shellId: 2 completed with exit code 2>
```

### ✅ 18. rg — 1472ms

**Input:**

```json
{"-C":2,"-n":true,"glob":"*.d.ts","head_limit":40,"output_mode":"content","paths":"node_modules\\@azure\\keyvault-secrets","pattern":"beginDeleteSecret|type.*Poller|interface.*Poller"}
```

**Output:**

```
[grep content: 320 matches across 28 file(s) under node_modules/@azure/keyvault-secrets/dist]

browser/index.d.ts (31 match(es)):
  3- import type { PageSettings, PagedAsyncIterableIterator } from "@azure/core-paging";
  4- import type { PollOperationState } from "@azure/core-lro";
  5: import type { PollerLike } from "@azure/core-lro";
  6: import type { BackupSecretOptions, BeginDeleteSecretOptions, BeginRecoverDeletedSecretOptions, DeletedSecret, GetDeletedSecretOptions, GetSecretOptions, KeyVaultSecret, ListDeletedSecretsOptions, ListPropertiesOfSecretVersionsOptions, ListPropertiesOfSecretsOptions, PurgeDeletedSecretOptions, RestoreSecretBackupOptions, SecretClientOptions, SecretPollerOptions, SecretProperties, SetSecretOptions, UpdateSecretPropertiesOptions } from "./secretsModels.js";
  7- import type { DeletionRecoveryLevel } from "./models/models.js";
  8- import { KnownDeletionRecoveryLevel, KnownContentType } from "./models/models.js";
  10- import type { KeyVaultSecretIdentifier } from "./identifier.js";
  11- import { parseKeyVaultSecretIdentifier } from "./identifier.js";
  12: export { type SecretClientOptions, type DeletedSecret, type DeletionRecoveryLevel, KnownDeletionRecoveryLevel, type ContentType, KnownContentType, type GetSecretOptions, type GetDeletedSecretOptions, type PurgeDeletedSecretOptions, type BackupSecretOptions, type RestoreSecretBackupOptions, type ListPropertiesOfSecretVersionsOptions, type ListPropertiesOfSecretsOptions, type ListDeletedSecretsOptions, type PagedAsyncIterableIterator, type PageSettings, type KeyVaultSecretIdentifier, parseKeyVaultSecretIdentifier, type PollerLike, type PollOperationState, type KeyVaultSecret, type SecretProperties, type SecretPollerOptions, type BeginDeleteSecretOptions, type BeginRecoverDeletedSecretOptions, type SetSecretOptions, type UpdateSecretPropertiesOptions, logger, };
  13- /**
  14-  * The SecretClient provides methods to manage {@link KeyVaultSecret} in
  98-      * const secretName = 
... (truncated)
```

### ✅ 19. apply_patch — 775ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/secret-rotation.ts\n@@\n import type {\n-  DeleteSecretPoller,\n   KeyVaultSecret,\n   SetSecretOptions\n } from \"@azure/keyvault-secrets\";\n \n+export interface DeleteSecretPollerLike {\n+  pollUntilDone(): Promise\u003cunknown\u003e;\n+}\n+\n export interface SecretWriter {\n   setSecret(name: string, value: string, options?: SetSecretOptions): Promise\u003cKeyVaultSecret\u003e;\n-  beginDeleteSecret(name: string): Promise\u003cDeleteSecretPoller\u003e;\n+  beginDeleteSecret(name: string): Promise\u003cDeleteSecretPollerLike\u003e;\n   purgeDeletedSecret(name: string): Promise\u003cvoid\u003e;\n }\n*** Update File: test/secret-config.test.ts\n@@\n import { describe, it } from \"node:test\";\n import type {\n-  DeleteSecretPoller,\n   GetSecretOptions,\n   KeyVaultSecret,\n   SetSecretOptions\n@@\n import {\n+  type DeleteSecretPollerLike,\n   SecretRotationHelper,\n   type SecretWriter\n } from \"../src/secret-rotation.js\";\n@@\n     const poller = {\n       pollUntilDone: async () =\u003e {\n         events.push(\"delete-complete\");\n         return {};\n       }\n-    } as unknown as DeleteSecretPoller;\n+    } satisfies DeleteSecretPollerLike;\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/src/secret-rotation.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/src/secret-rotation.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/src/secret-rotation.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/src/secret-rotation.ts
@@ -1,12 +1,15 @@
 import type {
-  DeleteSecretPoller,
   KeyVaultSecret,
   SetSecretOptions
 } from "@azure/keyvault-secrets";
 
+export interface DeleteSecretPollerLike {
+  pollUntilDone(): Promise<unknown>;
+}
+
 export interface SecretWriter {
   setSecret(name: string, value: string, options?: SetSecretOptions): Promise<KeyVaultSecret>;
-  beginDeleteSecret(name: string): Promise<DeleteSecretPoller>;
+  beginDeleteSecret(name: string): Promise<DeleteSecretPollerLike>;
   purgeDeletedSecret(name: string): Promise<void>;
 }
 


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/test/secret-config.test.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/test/secret-config.test.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/test/secret-config.test.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-js-ts-secret-config-js-ts-azure-skills-azure-skill-mcp-3631340615/test/secret-config.test.ts
@@ -1,7 +1,6 @@
 import assert from "node:assert/strict";
 import { describe, it } from "node:test";
 import type {
-  DeleteSecretPoller,
   GetSecretOptions,
   KeyVaultSecret,
   Set
... (truncated)
```

### ✅ 20. powershell — 2743ms

**Input:**

```json
{"command":"npm test","description":"Build and run project tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
> azure-key-vault-config-provider@1.0.0 test
> npm run build && node --test dist\test\secret-config.test.js


> azure-key-vault-config-provider@1.0.0 build
> tsc

▶ KeyVaultSecretProvider
  ✔ returns the default only for a missing secret (1.0081ms)
  ✔ retrieves a requested secret version (0.1934ms)
  ✔ does not hide non-404 service failures (0.3667ms)
✔ KeyVaultSecretProvider (2.2922ms)
▶ SecretCache
  ✔ bulk-loads, caches, and refreshes individual values (0.3739ms)
  ✔ automatically re-fetches a cached secret near expiry (0.2354ms)
✔ SecretCache (0.7618ms)
▶ SecretRotationHelper
  ✔ creates a new version with expiry (0.2502ms)
  ✔ waits for deletion before purge (0.2339ms)
  ✔ requires explicit confirmation before permanent purge (0.1485ms)
✔ SecretRotationHelper (0.8139ms)
ℹ tests 8
ℹ suites 3
ℹ pass 8
ℹ fail 0
ℹ cancelled 0
ℹ skipped 0
ℹ todo 0
ℹ duration_ms 94.3432
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\configuration.ts`
- `src\main.ts`
- `src\secret-cache.ts`
- `src\secret-provider.ts`
- `src\secret-rotation.ts`
- `test\secret-config.test.ts`
- `tsconfig.json`

## Copilot Response

I’m grounding the implementation in current Azure SDK guidance, then I’ll scaffold the local-only TypeScript project without touching any Azure subscription.

The design will keep destructive cleanup separate from ordinary rotation because Key Vault deletion applies to the secret name and all versions, including the newly rotated version. The demo will make purge explicitly opt-in.

I’ve finalized the project shape: injectable SDK-facing classes for local tests, managed-identity-only client creation, access-triggered plus periodic expiry refresh, and an irreversible purge API that requires explicit confirmation.

Created the complete TypeScript Node.js project with:

- Managed-identity authentication via `ManagedIdentityCredential`
- Version-aware secret retrieval and 404-only default fallback
- Expiry inspection and configurable warning window
- In-memory bulk loading, on-demand refresh, access-triggered refresh, and periodic refresh
- Secret rotation using new Key Vault versions
- Explicit, safe delete-and-purge flow that waits for the deletion LRO
- Opt-in destructive demo controls and redacted logging
- Unit tests, `package.json`, `tsconfig.json`, `.env.example`, and usage documentation

Run with:

```powershell
npm test
$env:KEY_VAULT_URL = "https://your-vault.vault.azure.net"
npm start
```

Rotation and irreversible purge remain disabled unless explicitly enabled through the environment variables documented in `.env.example`. No Azure resources were accessed or modified.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Complete dependency manifest with appropriate Azure SDK packages
- Includes TypeScript and Node.js type definitions
- Provides a suitable tsconfig.json for an ESM Node.js project

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
| `Criteria from prompt file` | prompt_review | 92% | 1.00 | 0.9231 | 11.7% | ❌ |
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
| **Final** | | | **Σ 11.00** | **Σ 7.9231** | **72.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-js-ts-secret-config --config js-ts-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
