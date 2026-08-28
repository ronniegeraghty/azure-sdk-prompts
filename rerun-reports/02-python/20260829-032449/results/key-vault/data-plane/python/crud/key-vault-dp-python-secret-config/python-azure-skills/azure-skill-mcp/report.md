# Evaluation Report: key-vault-dp-python-secret-config

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 325.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-python-secret-config` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 12/19 |
| Duration | 325.1s |
| Timestamp | 2026-08-28T20:43:22Z |
| Files Generated | 9 |
| Event Count | 14034 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 209.5s |
| Review | 114.2s |
| **Total** | **325.1s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
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
| Token Usage | in=924180 out=12388 |
| Turn Count | 15 |

## Tool Calls

### ✅ 1. skill — 2981ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 9243ms

**Input:**

```json
{"intent":"Create a local Python application configuration provider using Azure Key Vault Secrets sync and async clients, DefaultAzureCredential, in-memory caching, expiry checks, and safe secret rotation without deploying or modifying Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 9555ms

**Input:**

```json
{"intent":"Find official current Python SDK documentation and examples for azure-keyvault-secrets SecretClient and aio SecretClient: get_secret with versions, ResourceNotFoundError handling, expires_on properties, begin_delete_secret poller wait/result, and set_secret expires_on.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 8921ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 8599ms

**Input:**

```json
{"pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. glob — 8271ms

**Input:**

```json
{"pattern":"**/examples.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2310ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for a Python Azure Key Vault Secrets provider using DefaultAzureCredential, sync/async clients, caching, expiry handling, and safe deletion pollers.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 2931ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Azure SDK for Python docs for Key Vault Secrets sync/async secret retrieval, versions, expiry properties, deletion pollers, and setting expiration.","parameters":{"query":"Azure Key Vault Secrets Python SecretClient aio get_secret version expires_on begin_delete_secret poller wait set_secret expires_on DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"title":"Azure Key Vault Secrets client library for Python - version 4.11.2","content":"# Azure Key Vault Secrets client library for Python - version 4.11.2\n## Examples\n[begin_delete_secret](https://aka.ms/azsdk/python/keyvault-secrets/docs#azure.keyvault.secrets.SecretClient.begin_delete_secret) requests Key Vault delete a secret, returning a poller which allows you to wait for the deletion to finish. Waiting is helpful when the vault has [soft-delete](https://learn.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview) enabled, and you want to purge (permanently delete) the secret as soon as possible. When [soft-delete](https://learn.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview) is disabled, \u0060begin_delete_secret\u0060 itself is permanent.\n\u0060\u0060\u0060python\nfrom azure.identity import DefaultAzureCredential\nfrom azure.keyvault.secrets import SecretClient\n\ncredential = DefaultAzureCredential()\n\nsecret_client = SecretClient(vault_url=\u0022https://my-key-vault.vault.azure.net/\u0022, credential=credential)\ndeleted_secret = secret_client.begin_delete_secret(\u0022secret-name\u0022).result()\n\nprint(deleted_secret.name)\nprint(deleted_secret.deleted_date)\n\u0060\u0060\u0060\n### List secrets\n[list_properties_of_secrets](https://aka.ms/azsdk/python/keyvault-secrets/docs#azure.keyvault.secrets.SecretClient.list_properties_of_secrets) lists the properties of all of the secrets in the client\u0027s vault. This list doesn\u0027t include the secret\u0027s values.\n\u0060\u0060\u0060python\nfrom azure.identity import DefaultAzureCredential\nfrom azure.keyvault.secrets import SecretClient\n\ncredential = DefaultAzureCredential()\n\nsecret_client = SecretClient(vault_url=\u0022https://my-key-vault.vault.azure.net/\u0022, credential=credential)\nsecret_properties = secret_client.list_properties_of_secrets()\n\nfor secret_property in secret_properties:\n    # the list doesn\u0027t include values or versions of th
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 2622ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python examples for Azure Key Vault SecretClient and aio SecretClient including versioned retrieval, ResourceNotFoundError, deletion polling, set_secret expiry, and DefaultAzureCredential.","parameters":{"language":"python","query":"azure-keyvault-secrets Python SecretClient aio get_secret version ResourceNotFoundError begin_delete_secret wait set_secret expires_on DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"description":"description: Creates and authenticates a SecretClient instance using DefaultAzureCredential, then begins deletion of a secret in Azure Key Vault and waits for the deletion to complete, printing the deleted secret\u0027s name and deletion date.\npackage: azure.identity, azure.keyvault.secrets\nlanguage: python\n","codeSnippet":"from azure.identity import DefaultAzureCredential\r\nfrom azure.keyvault.secrets import SecretClient\r\n\r\ncredential = DefaultAzureCredential()\r\n\r\nsecret_client = SecretClient(vault_url=\u0022https://my-key-vault.vault.azure.net/\u0022, credential=credential)\r\ndeleted_secret = secret_client.begin_delete_secret(\u0022secret-name\u0022).result()\r\n\r\nprint(deleted_secret.name)\r\nprint(deleted_secret.deleted_date)","link":"https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme?view=azure-python#examples","language":"python"},{"description":"description: Demonstrates how to create an Azure Key Vault SecretClient using DefaultAzureCredential, set a secret, retrieve it, and delete it from the Key Vault.\npackage: azure-keyvault-secrets, azure-identity\nlanguage: python\n","codeSnippet":"import os\r\nfrom azure.keyvault.secrets import SecretClient\r\nfrom azure.identity import DefaultAzureCredential\r\n\r\nkeyVaultName = os.environ[\u0022KEY_VAULT_NAME\u0022]\r\nKVUri = f\u0022https://{keyVaultName}.vault.azure.net\u0022\r\n\r\ncredential = DefaultAzureCredential()\r\nclient = SecretClient(vault_url=KVUri, credential=credential)\r\n\r\nsecretName = input(\u0022Input a name for your secret \u003E \u0022)\r\nsecretValue = input(\u0022Input a value for your secret \u003E \u0022)\r\n\r\nprint(f\u0022Creating a secret in {keyVaultName} called \u0027{secretName}\u0027 with the value \u0027{secretValue}\u0027 ...\u0022)\r\n\r\nclient.set_secret(secretName, secretValue)\r\n\r\nprint(\u0022 done.\u0022)\r\n\r\nprint(f\u0022Retrieving your secret from {keyVaultName}.\u0022)\r\n\r\nretrieved_secret =
... (truncated)
```

### ✅ 10. rg — 2113ms

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt","pattern":"begin_delete_secret|poller|wait\\(|set_secret|expires_on"}
```

**Output:**

```
Output too large to read at once (61.1 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787949862182-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt:{"results":[{"description":"description: Creates and authenticates a SecretClient instance using DefaultAzureCredential, then begins deletion of a secret in Azure Key Vault and waits for the deletion to complete, printing the deleted secret\u0027s name and deletion date.\npackage: azure.identity, azure.keyvault.secrets\nlanguage: python\n","codeSnippet":"from azure.identity import DefaultAzu
```

### ✅ 11. rg — 1796ms

**Input:**

```json
{"-C":2,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt","pattern":"async|aio|begin_delete_secret|poller|expires_on|get_secret\\("}
```

**Output:**

```
Output too large to read at once (39.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787949862181-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt:{"results":[{"title":"Azure Key Vault Secrets client library for Python - version 4.11.2","content":"# Azure Key Vault Secrets client library for Python - version 4.11.2\n## Examples\n[begin_delete_secret](https://aka.ms/azsdk/python/keyvault-secrets/docs#azure.keyvault.secrets.SecretClient.begin_delete_secret) requests Key Vault delete a secret, returning a poller which allows you to wait f
```

### ✅ 12. rg — 1491ms

**Input:**

```json
{"head_limit":40,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt"],"pattern":"https://learn.microsoft.com[^\\\" ]+|https://aka.ms[^\\\" ]+"}
```

**Output:**

```
Output too large to read at once (21.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787949862183-3.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt:{"results":[{"description":"description: Creates and authenticates a SecretClient instance using DefaultAzureCredential, then begins deletion of a secret in Azure Key Vault and waits for the deletion to complete, printing the deleted secret\u0027s name and deletion date.\npackage: azure.identity, azure.keyvault.secrets\nlanguage: python\n","codeSnippet":"from azure.identity import DefaultAzu
```

### ✅ 13. powershell — 1152ms

**Input:**

```json
{"command":"$sample = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt' | ConvertFrom-Json; $sample.results | Select-Object -First 8 | ForEach-Object { \"---`n$($_.description)`n$($_.codeSnippet)`n$($_.link)\" }; $docs = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt' | ConvertFrom-Json; $docs.results | Select-Object -First 6 | ForEach-Object { \"=== $($_.title) ===`n$($_.content)`n$($_.contentUrl)\" }","description":"Extract relevant Azure SDK documentation","initial_wait":30,"mode":"sync"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-fbcc59c14d784891b79b6b19c72e9392.txt' | ConvertFrom-Json; $sampl …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 14. view — 1336ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt b/C:/Users/weidxu/AppData/Local/Temp/1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt
@@ -1,8 +1,8 @@
 {"results":[{"title":"Azure Key Vault Secrets client library for Python - version 4.11.2","content":"# Azure Key Vault Secrets client library for Python - version 4.11.2\n## Examples\n[begin_delete_secret](https://aka.ms/azsdk/python/keyvault-secrets/docs#azure.keyvault.secrets.SecretClient.begin_delete_secret) requests Key Vault delete a secret, returning a poller which allows you to wait for the deletion to finish. Waiting is helpful when the vault has [soft-delete](https://learn.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview) enabled, and you want to purge (permanently delete) the secret as soon as possible. When [soft-delete](https://learn.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview) is disabled, \u0060begin_delete_secret\u0060 itself is permanent.\n\u0060\u0060\u0060python\nfrom azure.identity import DefaultAzureCredential\nfrom azure.keyvault.secrets import SecretClient\n\ncredential = DefaultAzureCredential()\n\nsecret_client = SecretClient(vault_url=\u0022https://my-key-vault.vault.azure.net/\u0022, credential=credential)\ndeleted_secret = secret_client.begin_delete_secret(\u0022secret-name\u0022).result()\n\nprint(deleted_secret.name)\nprint(deleted_secret.deleted_date)\n\u0060\u0060\u0060\n### List secrets\n[list_properties_of_secrets](https://aka.ms/azsdk/python/keyvault-secrets/docs#azure.keyvault.secrets.SecretClient.list_properties_of_secrets) lists the properties of all of the secrets in the client\u0027s vault. This list d
... (truncated)
```

### ✅ 15. view — 998ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949852961-copilot-tool-output-8bf94b64e1fd4f65805e0ed2af439b01.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt b/C:/Users/weidxu/AppData/Local/Temp/1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787949852729-copilot-tool-output-fbcc59c14d784891b79b6b19c72e9392.txt
@@ -1,8 +1,8 @@
 {"results":[{"description":"description: Creates and authenticates a SecretClient instance using DefaultAzureCredential, then begins deletion of a secret in Azure Key Vault and waits for the deletion to complete, printing the deleted secret\u0027s name and deletion date.\npackage: azure.identity, azure.keyvault.secrets\nlanguage: python\n","codeSnippet":"from azure.identity import DefaultAzureCredential\r\nfrom azure.keyvault.secrets import SecretClient\r\n\r\ncredential = DefaultAzureCredential()\r\n\r\nsecret_client = SecretClient(vault_url=\u0022https://my-key-vault.vault.azure.net/\u0022, credential=credential)\r\ndeleted_secret = secret_client.begin_delete_secret(\u0022secret-name\u0022).result()\r\n\r\nprint(deleted_secret.name)\r\nprint(deleted_secret.deleted_date)","link":"https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme?view=azure-python#examples","language":"python"},{"description":"description: Demonstrates how to create an Azure Key Vault SecretClient using DefaultAzureCredential, set a secret, retrieve it, and delete it from the Key Vault.\npackage: azure-keyvault-secrets, azure-identity\nlanguage: python\n","codeSnippet":"import os\r\nfrom azure.keyvault.secrets import SecretClient\r\nfrom azure.identity import DefaultAzureCredential\r\n\r\nkeyVaultName = os.environ[\u0022KEY_VAULT_NAME\u0022]\r\nKVUri = f\u0022https://{keyVaultName}.vault.azure.net\u0022\r\n\r\ncredential = DefaultAzureCredential()\r\nclient = SecretClient(vault_url=
... (truncated)
```

### ✅ 16. azure-documentation (via azure) — 2935ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete current sync SecretClient method documentation for versioned get, expiry, delete poller, and set secret.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-keyvault-secrets/azure.keyvault.secrets.secretclient?view=azure-python"}}
```

**Output:**

```
# Azure Key Vault Secrets client library for Python - version 4.11.2

Azure Key Vault helps solve the following problems:

- Secrets management (this library) - securely store and control access to tokens, passwords, certificates, API keys, and other secrets
- Cryptographic key management ([azure-keyvault-keys](https://github.com/Azure/azure-sdk-for-python/tree/azure-keyvault-secrets_4.11.2/sdk/keyvault/azure-keyvault-keys)) - create, store, and control access to the keys used to encrypt your data
- Certificate management ([azure-keyvault-certificates](https://github.com/Azure/azure-sdk-for-python/tree/azure-keyvault-secrets_4.11.2/sdk/keyvault/azure-keyvault-certificates)) - create, manage, and deploy public and private SSL/TLS certificates
- Vault administration ([azure-keyvault-administration](https://github.com/Azure/azure-sdk-for-python/tree/azure-keyvault-secrets_4.11.2/sdk/keyvault/azure-keyvault-administration)) - role-based access control (RBAC), and vault-level backup and restore options

[Source code](https://github.com/Azure/azure-sdk-for-python/tree/azure-keyvault-secrets_4.11.2/sdk/keyvault/azure-keyvault-secrets/azure/keyvault/secrets) | [Package (PyPI)](https://pypi.org/project/azure-keyvault-secrets/) | [Package (Conda)](https://anaconda.org/microsoft/azure-keyvault/) | [API reference documentation](https://aka.ms/azsdk/python/keyvault-secrets/docs) | [Product documentation](/en-us/azure/key-vault/general/overview) | [Samples](https://github.com/Azure/azure-sdk-for-python/tree/azure-keyvault-secrets_4.11.2/sdk/keyvault/azure-keyvault-secrets/samples)

## *Disclaimer*

*Azure SDK Python packages support for Python 2.7 has ended 01 January 2022. For more information and questions, please refer to https://github.com/Azure/azure-sdk-for-python/issues/20691*. *Python 3.9 or later is required to use this package. For more details, please refer to [Azure SDK for Python version support policy](https://github.com/Azure/azure-sdk-for-python/blob/
... (truncated)
```

### ✅ 17. azure-documentation (via azure) — 2603ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete current async SecretClient method documentation for versioned get, expiry, delete poller, and set secret.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-keyvault-secrets/azure.keyvault.secrets.aio.secretclient?view=azure-python"}}
```

**Output:**

```
# SecretClient Class

A high-level interface for managing a vault's secrets.

## Constructor

```python
SecretClient(vault_url: str, credential: TokenCredential, **kwargs: Any)
```

### Parameters

| Name | Description |
| --- | --- |
| vault\_url<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>URL of the vault the client will access. This is also called the vault's "DNS Name". You should validate that this URL references a valid Key Vault resource. See https://aka.ms/azsdk/blog/vault-uri for details. |
| credential<br><br>Required | [TokenCredential](../azure-core/azure.core.credentials.tokencredential?view=azure-python)<br><br>An object which can provide an access token for the vault, such as a credential from [identity](../azure-identity/azure.identity?view=azure-python) |

### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| api\_version | [ApiVersion](azure.keyvault.secrets.apiversion?view=azure-python) or [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Version of the service API to use. Defaults to the most recent. |
| verify\_challenge\_resource | [bool](https://docs.python.org/3/library/functions.html#bool)<br><br>Whether to verify the authentication challenge resource matches the Key Vault domain. Defaults to True. |

## Methods

| backup_secret | Back up a secret in a protected form useable only by Azure Key Vault. Requires secrets/backup permission. |
| --- | --- |
| begin_delete_secret | Delete all versions of a secret. Requires secrets/delete permission.<br><br>When this method returns Key Vault has begun deleting the secret. Deletion may take several seconds in a vault with soft-delete enabled. This method therefore returns a poller enabling you to wait for deletion to complete. |
| begin_recover_deleted_secret | Recover a deleted secret to its latest version. Possible only in a vault with soft-delete enabled.<br><br>Requires the secrets/recover permission
... (truncated)
```

### ✅ 18. azure-documentation (via azure) — 2298ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch current Azure Key Vault Secrets Python readme examples including sync and async lifecycle guidance.","parameters":{"url":"https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme?view=azure-python"}}
```

**Output:**

```
# SecretClient Class

A high-level asynchronous interface for managing a vault's secrets.

## Constructor

```python
SecretClient(vault_url: str, credential: AsyncTokenCredential, **kwargs: Any)
```

### Parameters

| Name | Description |
| --- | --- |
| vault\_url<br><br>Required | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>URL of the vault the client will access. This is also called the vault's "DNS Name". You should validate that this URL references a valid Key Vault resource. See https://aka.ms/azsdk/blog/vault-uri for details. |
| credential<br><br>Required | [AsyncTokenCredential](../azure-core/azure.core.credentials_async.asynctokencredential?view=azure-python)<br><br>An object which can provide an access token for the vault, such as a credential from [aio](../azure-identity/azure.identity.aio?view=azure-python) |

### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| api\_version | [ApiVersion](azure.keyvault.secrets.apiversion?view=azure-python) or [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Version of the service API to use. Defaults to the most recent. |
| verify\_challenge\_resource | [bool](https://docs.python.org/3/library/functions.html#bool)<br><br>Whether to verify the authentication challenge resource matches the Key Vault domain. Defaults to True. |

## Methods

| backup_secret | Back up a secret in a protected form useable only by Azure Key Vault. Requires secrets/backup permission. |
| --- | --- |
| close | Close sockets opened by the client.<br><br>Calling this method is unnecessary when using the client as a context manager. |
| delete_secret | Delete all versions of a secret. Requires secrets/delete permission.<br><br>If the vault has soft-delete enabled, deletion may take several seconds to complete. |
| get_deleted_secret | Get a deleted secret. Possible only in vaults with soft-delete enabled. Requires secrets/get permission. |
| get_secret | Get a se
... (truncated)
```

### ✅ 19. rg — 2685ms

**Input:**

```json
{"-C":4,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt","pattern":"begin_delete_secret|async begin|AsyncLROPoller|poller|wait method|result method"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-Get a secret. Requires the secrets/get permission.
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-```python
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt:get_secret(name: str, version: str | None = None, *, out_content_type: str | ContentType | None = None, **kwargs: Any) -> KeyVaultSecret
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-```
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-#### Parameters
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
--
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-Requires secrets/set permission.
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-```python
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt:set_secret(name: str, value: str, *, enabled: bool | None = None, tags: Dict[str, str] | None = None, content_type: str | None = None, not_before: datetime | None = None, expires_on: datetime | None = None, **kwargs: Any) -> KeyVault
... (truncated)
```

### ✅ 20. rg — 2368ms

**Input:**

```json
{"-C":4,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt","pattern":"get_secret\\(|version.*Optional|expires_on|set_secret\\("}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-## Methods
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-| backup_secret | Back up a secret in a protected form useable only by Azure Key Vault. Requires secrets/backup permission. |
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt-| --- | --- |
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt:| begin_delete_secret | Delete all versions of a secret. Requires secrets/delete permission.<br><br>When this method returns Key Vault has begun deleting the secret. Deletion may take several seconds in a vault with soft-delete enabled. This method therefore returns a poller enabling you to wait for deletion to complete. |
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt:| begin_recover_deleted_secret | Recover a deleted secret to its latest version. Possible only in a vault with soft-delete enabled.<br><br>Requires the secrets/recover permission. If the vault does not have soft-delete enabled, [begin_delete_secret](azure.keyvault.secrets.secretclient?view=azure-python#azure-keyvault-secrets-secretclient-begin-delete-secret) is permanent, and this method will return an error. Attempting to recover a non-deleted secret will also return an error. When this method returns Key Vault has begun recovering the secret. Recovery may take several seconds. This method therefore returns a poller enabling you to wait for recovery to complete. Waiting is only necessary when you want to use the recovered secret in another operation immediately. |
C:\Users\weidxu\AppData\Local\Temp\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt
... (truncated)
```

### ✅ 21. rg — 2024ms

**Input:**

```json
{"-C":4,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt","pattern":"begin_delete_secret|LROPoller|poller|wait method|result method"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-Get a secret. Requires the secrets/get permission.
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-```python
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt:async get_secret(name: str, version: str | None = None, *, out_content_type: str | ContentType | None = None, **kwargs: Any) -> KeyVaultSecret
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-```
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-#### Parameters
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-
--
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-Requires secrets/set permission.
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-```python
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt:async set_secret(name: str, value: str, *, enabled: bool | None = None, tags: Dict[str, str] | None = None, content_type: str | None = None, not_before: datetime | None = None, expires_on: datetime | None = None, **kwargs: Any)
... (truncated)
```

### ✅ 22. rg — 1697ms

**Input:**

```json
{"-C":4,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949893529-copilot-tool-output-e6f3a0d326e542c2a56cdfb849e3cf0f.txt","pattern":"get_secret\\(|version.*Optional|expires_on|set_secret\\("}
```

**Output:**

```
No matches found.
```

### ✅ 23. rg — 710ms

**Input:**

```json
{"-C":3,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt","pattern":"\\| begin_delete_secret|### begin|delete_secret"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| backup_secret | Back up a secret in a protected form useable only by Azure Key Vault. Requires secrets/backup permission. |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| --- | --- |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| close | Close sockets opened by the client.<br><br>Calling this method is unnecessary when using the client as a context manager. |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt:| delete_secret | Delete all versions of a secret. Requires secrets/delete permission.<br><br>If the vault has soft-delete enabled, deletion may take several seconds to complete. |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| get_deleted_secret | Get a deleted secret. Possible only in vaults with soft-delete enabled. Requires secrets/get permission. |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| get_secret | Get a secret. Requires the secrets/get permission. |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| list_deleted_secrets | Lists all deleted secrets. Possible only in vaults with soft-delete enabled.<br><br>Requires secrets/list permission. |
C:\Users\weidxu\AppData\Local\Temp\1787949893258-copilot-tool-output-191d1c0e4a4b4685aeaa4ac82f57f11d.txt-| list_properties_of_secret_versions | List properties of all versions of a secret, excluding their values. Requires secrets/list permission.<br><br>List items don't include secret values. Use [get_secret](azure.keyvault.secrets.aio.secretclient?view=azure-python#azure-keyvault-secrets-aio-secretclient-get-secret) to get a secret's value. |
C:\
... (truncated)
```

### ✅ 24. github-mcp-server-get_file_contents (via github-mcp-server) — 2816ms

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/azure-keyvault-secrets/azure/keyvault/secrets/aio/_client.py","ref":"refs/tags/azure-keyvault-secrets_4.11.2","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: e9c048af41d91cac58a4c0d5b0d5d60fc9145a38)# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------
from datetime import datetime
from typing import Any, cast, Dict, Optional, Union
from functools import partial

from azure.core.tracing.decorator import distributed_trace
from azure.core.tracing.decorator_async import distributed_trace_async
from azure.core.async_paging import AsyncItemPaged

from .._generated.models import ContentType
from .._models import KeyVaultSecret, DeletedSecret, SecretProperties
from .._shared import AsyncKeyVaultClientBase
from .._shared._polling_async import AsyncDeleteRecoverPollingMethod


class SecretClient(AsyncKeyVaultClientBase):
    """A high-level asynchronous interface for managing a vault's secrets.

    :param str vault_url: URL of the vault the client will access. This is also called the vault's "DNS Name".
        You should validate that this URL references a valid Key Vault resource. See https://aka.ms/azsdk/blog/vault-uri
        for details.
    :param credential: An object which can provide an access token for the vault, such as a credential from
        :mod:`azure.identity.aio`
    :type credential: ~azure.core.credentials_async.AsyncTokenCredential

    :keyword api_version: Version of the service API to use. Defaults to the most recent.
    :paramtype api_version: ~azure.keyvault.secrets.ApiVersion or str
    :keyword bool verify_challenge_resource: Whether to verify the authentication challenge resource matches the Key
        Vault domain. Defaults to True.

    Example:
        .. literalinclude:: ../tests/test_samples_secrets_async.py
            :start-after: [START create_secret_client]
            :end-before: [END create_secret_client]
            :language: python
            :caption: Create a new ``SecretClient``
            :dedent: 4
    """

    # pylint:disable=protected-ac
... (truncated)
```

### ✅ 25. github-mcp-server-get_file_contents (via github-mcp-server) — 2566ms

**Input:**

```json
{"owner":"Azure","path":"sdk/keyvault/azure-keyvault-secrets/azure/keyvault/secrets/aio/_client.py","ref":"refs/tags/azure-keyvault-secrets_4.10.0","repo":"azure-sdk-for-python"}
```

**Output:**

```
successfully downloaded text file (SHA: 20904cf11646ef84fd1dfebd0a43ac705a7359b5)# ------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------
from datetime import datetime
from typing import Any, cast, Dict, Optional
from functools import partial

from azure.core.tracing.decorator import distributed_trace
from azure.core.tracing.decorator_async import distributed_trace_async
from azure.core.async_paging import AsyncItemPaged

from .._models import KeyVaultSecret, DeletedSecret, SecretProperties
from .._shared import AsyncKeyVaultClientBase
from .._shared._polling_async import AsyncDeleteRecoverPollingMethod


class SecretClient(AsyncKeyVaultClientBase):
    """A high-level asynchronous interface for managing a vault's secrets.

    :param str vault_url: URL of the vault the client will access. This is also called the vault's "DNS Name".
        You should validate that this URL references a valid Key Vault resource. See https://aka.ms/azsdk/blog/vault-uri
        for details.
    :param credential: An object which can provide an access token for the vault, such as a credential from
        :mod:`azure.identity.aio`
    :type credential: ~azure.core.credentials_async.AsyncTokenCredential

    :keyword api_version: Version of the service API to use. Defaults to the most recent.
    :paramtype api_version: ~azure.keyvault.secrets.ApiVersion or str
    :keyword bool verify_challenge_resource: Whether to verify the authentication challenge resource matches the Key
        Vault domain. Defaults to True.

    Example:
        .. literalinclude:: ../tests/test_samples_secrets_async.py
            :start-after: [START create_secret_client]
            :end-before: [END create_secret_client]
            :language: python
            :caption: Create a new ``SecretClient``
            :dedent: 4
    """

    # pylint:disable=protected-access

    @distributed_trace_async
    async def ge
... (truncated)
```

### ✅ 26. apply_patch — 1509ms

**Input:**

```json
"*** Begin Patch\n*** Add File: azure_keyvault_config/__init__.py\n+\"\"\"Azure Key Vault-backed application configuration.\"\"\"\n+\n+from .cache import AsyncSecretCache, SecretCache\n+from .factory import create_async_provider, create_sync_provider\n+from .provider import AsyncSecretProvider, SecretDetails, SecretProvider\n+from .rotation import AsyncSecretRotator, SecretRotator\n+\n+__all__ = [\n+    \"AsyncSecretCache\",\n+    \"AsyncSecretProvider\",\n+    \"AsyncSecretRotator\",\n+    \"SecretCache\",\n+    \"SecretDetails\",\n+    \"SecretProvider\",\n+    \"SecretRotator\",\n+    \"create_async_provider\",\n+    \"create_sync_provider\",\n+]\n*** Add File: azure_keyvault_config/provider.py\n+\"\"\"Synchronous and asynchronous Key Vault secret providers.\"\"\"\n+\n+from __future__ import annotations\n+\n+from dataclasses import dataclass\n+from datetime import datetime\n+from typing import Optional\n+\n+from azure.core.exceptions import ResourceNotFoundError\n+from azure.keyvault.secrets import SecretClient\n+from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient\n+\n+\n+@dataclass(frozen=True)\n+class SecretDetails:\n+    \"\"\"A secret value and the metadata needed by the configuration cache.\"\"\"\n+\n+    name: str\n+    value: Optional[str]\n+    version: Optional[str]\n+    expires_on: Optional[datetime]\n+    found: bool\n+\n+\n+class SecretProvider:\n+    \"\"\"Read secrets through the synchronous Azure Key Vault client.\"\"\"\n+\n+    def __init__(self, client: SecretClient, credential: object | None = None) -\u003e None:\n+        self.client = client\n+        self._credential = credential\n+\n+    def get_secret(\n+        self, name: str, default: Optional[str] = None, version: Optional[str] = None\n+    ) -\u003e Optional[str]:\n+        return self.get_secret_details(name, default=default, version=version).value\n+\n+    def get_secret_details(\n+        self, name: str, default: Optional[str] = None, version: Optional[str] = None\n+    ) -\u003e SecretDetails:\n+        try:\n+            secret = self.client.get_secret(name, version)\n+        except ResourceNotFoundError:\n+            return SecretDetails(name, default, version, None, False)\n+\n+        return SecretDetails(\n+            name=secret.name,\n+            value=secret.value,\n+            version=secret.properties.version,\n+            expires_on=secret.properties.expires_on,\n+            found=True,\n+        )\n+\n+    def get_expiry(\n+        self, name: str, version: Optional[str] = None\n+    ) -\u003e Optional[datetime]:\n+        return self.get_secret_details(name, version=version).expires_on\n+\n+    def close(self) -\u003e None:\n+        self.client.close()\n+        close = getattr(self._credential, \"close\", None)\n+        if close is not None:\n+            close()\n+\n+    def __enter__(self) -\u003e \"SecretProvider\":\n+        return self\n+\n+    def __exit__(self, exc_type: object, exc_value: object, traceback: object) -\u003e None:\n+        self.close()\n+\n+\n+class AsyncSecretProvider:\n+    \"\"\"Read secrets through the asynchronous Azure Key Vault client.\"\"\"\n+\n+    def __init__(\n+        self, client: AsyncSecretClient, credential: object | None = None\n+    ) -\u003e None:\n+        self.client = client\n+        self._credential = credential\n+\n+    async def get_secret(\n+        self, name: str, default: Optional[str] = None, version: Optional[str] = None\n+    ) -\u003e Optional[str]:\n+        details = await self.get_secret_details(name, default=default, version=version)\n+        return details.value\n+\n+    async def get_secret_details(\n+        self, name: str, default: Optional[str] = None, version: Optional[str] = None\n+    ) -\u003e SecretDetails:\n+        try:\n+            secret = await self.client.get_secret(name, version)\n+        except ResourceNotFoundError:\n+            return SecretDetails(name, default, version, None, False)\n+\n+        return SecretDetails(\n+            name=secret.name,\n+            value=secret.value,\n+            version=secret.properties.version,\n+            expires_on=secret.properties.expires_on,\n+            found=True,\n+        )\n+\n+    async def get_expiry(\n+        self, name: str, version: Optional[str] = None\n+    ) -\u003e Optional[datetime]:\n+        details = await self.get_secret_details(name, version=version)\n+        return details.expires_on\n+\n+    async def close(self) -\u003e None:\n+        await self.client.close()\n+        close = getattr(self._credential, \"close\", None)\n+        if close is not None:\n+            await close()\n+\n+    async def __aenter__(self) -\u003e \"AsyncSecretProvider\":\n+        return self\n+\n+    async def __aexit__(\n+        self, exc_type: object, exc_value: object, traceback: object\n+    ) -\u003e None:\n+        await self.close()\n*** Add File: azure_keyvault_config/cache.py\n+\"\"\"In-memory caching for Key Vault-backed configuration.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+from datetime import datetime, timedelta, timezone\n+from typing import Iterable, Mapping, Optional\n+\n+from .provider import AsyncSecretProvider, SecretDetails, SecretProvider\n+\n+\n+def _utc(value: datetime) -\u003e datetime:\n+    if value.tzinfo is None:\n+        return value.replace(tzinfo=timezone.utc)\n+    return value.astimezone(timezone.utc)\n+\n+\n+def _is_near_expiry(\n+    details: SecretDetails, warning_window: timedelta, now: datetime\n+) -\u003e bool:\n+    return (\n+        details.expires_on is not None\n+        and _utc(details.expires_on) \u003c= now + warning_window\n+    )\n+\n+\n+class SecretCache:\n+    \"\"\"Cache secret values and refresh entries approaching expiration.\"\"\"\n+\n+    def __init__(\n+        self,\n+        provider: SecretProvider,\n+        required_keys: Iterable[str] = (),\n+        defaults: Mapping[str, Optional[str]] | None = None,\n+        warning_window: timedelta = timedelta(days=7),\n+    ) -\u003e None:\n+        if warning_window \u003c timedelta(0):\n+            raise ValueError(\"warning_window must not be negative\")\n+        self._provider = provider\n+        self._required_keys = tuple(dict.fromkeys(required_keys))\n+        self._defaults = dict(defaults or {})\n+        self._warning_window = warning_window\n+        self._entries: dict[str, SecretDetails] = {}\n+\n+    def load_required(self) -\u003e dict[str, Optional[str]]:\n+        for name in self._required_keys:\n+            self.refresh(name)\n+        return self.values()\n+\n+    def get(self, name: str, default: Optional[str] = None) -\u003e Optional[str]:\n+        entry = self._entries.get(name)\n+        if entry is None or _is_near_expiry(\n+            entry, self._warning_window, datetime.now(timezone.utc)\n+        ):\n+            entry = self.refresh(name, default)\n+        return entry.value\n+\n+    def refresh(\n+        self, name: str, default: Optional[str] = None\n+    ) -\u003e SecretDetails:\n+        fallback = self._defaults.get(name, default)\n+        entry = self._provider.get_secret_details(name, default=fallback)\n+        self._entries[name] = entry\n+        return entry\n+\n+    def refresh_expiring(self) -\u003e tuple[str, ...]:\n+        now = datetime.now(timezone.utc)\n+        names = tuple(\n+            name\n+            for name, entry in self._entries.items()\n+            if _is_near_expiry(entry, self._warning_window, now)\n+        )\n+        for name in names:\n+            self.refresh(name)\n+        return names\n+\n+    def expiring_secrets(self) -\u003e dict[str, datetime]:\n+        now = datetime.now(timezone.utc)\n+        return {\n+            name: entry.expires_on\n+            for name, entry in self._entries.items()\n+            if entry.expires_on is not None\n+            and _is_near_expiry(entry, self._warning_window, now)\n+        }\n+\n+    def values(self) -\u003e dict[str, Optional[str]]:\n+        return {name: entry.value for name, entry in self._entries.items()}\n+\n+\n+class AsyncSecretCache:\n+    \"\"\"Asynchronous counterpart to :class:`SecretCache`.\"\"\"\n+\n+    def __init__(\n+        self,\n+        provider: AsyncSecretProvider,\n+        required_keys: Iterable[str] = (),\n+        defaults: Mapping[str, Optional[str]] | None = None,\n+        warning_window: timedelta = timedelta(days=7),\n+    ) -\u003e None:\n+        if warning_window \u003c timedelta(0):\n+            raise ValueError(\"warning_window must not be negative\")\n+        self._provider = provider\n+        self._required_keys = tuple(dict.fromkeys(required_keys))\n+        self._defaults = dict(defaults or {})\n+        self._warning_window = warning_window\n+        self._entries: dict[str, SecretDetails] = {}\n+\n+    async def load_required(self) -\u003e dict[str, Optional[str]]:\n+        await asyncio.gather(*(self.refresh(name) for name in self._required_keys))\n+        return self.values()\n+\n+    async def get(self, name: str, default: Optional[str] = None) -\u003e Optional[str]:\n+        entry = self._entries.get(name)\n+        if entry is None or _is_near_expiry(\n+            entry, self._warning_window, datetime.now(timezone.utc)\n+        ):\n+            entry = await self.refresh(name, default)\n+        return entry.value\n+\n+    async def refresh(\n+        self, name: str, default: Optional[str] = None\n+    ) -\u003e SecretDetails:\n+        fallback = self._defaults.get(name, default)\n+        entry = await self._provider.get_secret_details(name, default=fallback)\n+        self._entries[name] = entry\n+        return entry\n+\n+    async def refresh_expiring(self) -\u003e tuple[str, ...]:\n+        now = datetime.now(timezone.utc)\n+        names = tuple(\n+            name\n+            for name, entry in self._entries.items()\n+            if _is_near_expiry(entry, self._warning_window, now)\n+        )\n+        await asyncio.gather(*(self.refresh(name) for name in names))\n+        return names\n+\n+    def expiring_secrets(self) -\u003e dict[str, datetime]:\n+        now = datetime.now(timezone.utc)\n+        return {\n+            name: entry.expires_on\n+            for name, entry in self._entries.items()\n+            if entry.expires_on is not None\n+            and _is_near_expiry(entry, self._warning_window, now)\n+        }\n+\n+    def values(self) -\u003e dict[str, Optional[str]]:\n+        return {name: entry.value for name, entry in self._entries.items()}\n*** Add File: azure_keyvault_config/factory.py\n+\"\"\"Factories for securely authenticated Key Vault providers.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from urllib.parse import urlparse\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.keyvault.secrets import SecretClient\n+from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient\n+\n+from .provider import AsyncSecretProvider, SecretProvider\n+\n+VAULT_URL_ENV = \"AZURE_KEY_VAULT_URL\"\n+\n+\n+def _vault_url_from_environment() -\u003e str:\n+    vault_url = os.environ.get(VAULT_URL_ENV)\n+    if not vault_url:\n+        raise RuntimeError(f\"{VAULT_URL_ENV} must contain the Key Vault URL\")\n+\n+    parsed = urlparse(vault_url)\n+    if parsed.scheme != \"https\" or not parsed.netloc or parsed.username:\n+        raise ValueError(f\"{VAULT_URL_ENV} must be a valid HTTPS URL\")\n+    return vault_url\n+\n+\n+def create_sync_provider() -\u003e SecretProvider:\n+    credential = DefaultAzureCredential()\n+    client = SecretClient(\n+        vault_url=_vault_url_from_environment(), credential=credential\n+    )\n+    return SecretProvider(client, credential)\n+\n+\n+def create_async_provider() -\u003e AsyncSecretProvider:\n+    credential = AsyncDefaultAzureCredential()\n+    client = AsyncSecretClient(\n+        vault_url=_vault_url_from_environment(), credential=credential\n+    )\n+    return AsyncSecretProvider(client, credential)\n*** Add File: azure_keyvault_config/rotation.py\n+\"\"\"Safe secret rotation helpers.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import time\n+from datetime import datetime, timezone\n+\n+from azure.core.exceptions import ResourceNotFoundError\n+from azure.keyvault.secrets import KeyVaultSecret, SecretClient\n+from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient\n+\n+\n+def _validate_expiry(expires_on: datetime) -\u003e None:\n+    normalized = (\n+        expires_on.replace(tzinfo=timezone.utc)\n+        if expires_on.tzinfo is None\n+        else expires_on.astimezone(timezone.utc)\n+    )\n+    if normalized \u003c= datetime.now(timezone.utc):\n+        raise ValueError(\"expires_on must be in the future\")\n+\n+\n+class SecretRotator:\n+    \"\"\"Delete, fully remove, and recreate a secret with the same name.\"\"\"\n+\n+    def __init__(\n+        self,\n+        client: SecretClient,\n+        purge_timeout: float = 120.0,\n+        polling_interval: float = 1.0,\n+    ) -\u003e None:\n+        if purge_timeout \u003c= 0 or polling_interval \u003c= 0:\n+            raise ValueError(\"timeouts and polling intervals must be positive\")\n+        self._client = client\n+        self._purge_timeout = purge_timeout\n+        self._polling_interval = polling_interval\n+\n+    def rotate(\n+        self, name: str, value: str, expires_on: datetime\n+    ) -\u003e KeyVaultSecret:\n+        _validate_expiry(expires_on)\n+        try:\n+            poller = self._client.begin_delete_secret(name)\n+            poller.wait()\n+            deleted = poller.result()\n+        except ResourceNotFoundError:\n+            deleted = None\n+\n+        if deleted is not None and deleted.recovery_id is not None:\n+            self._client.purge_deleted_secret(name)\n+            self._wait_until_purged(name)\n+\n+        return self._client.set_secret(name, value, expires_on=expires_on)\n+\n+    def _wait_until_purged(self, name: str) -\u003e None:\n+        deadline = time.monotonic() + self._purge_timeout\n+        while True:\n+            try:\n+                self._client.get_deleted_secret(name)\n+            except ResourceNotFoundError:\n+                return\n+            if time.monotonic() \u003e= deadline:\n+                raise TimeoutError(\n+                    f\"Secret {name!r} was not purged within \"\n+                    f\"{self._purge_timeout:.1f} seconds\"\n+                )\n+            time.sleep(self._polling_interval)\n+\n+\n+class AsyncSecretRotator:\n+    \"\"\"Asynchronously delete, fully remove, and recreate a secret.\"\"\"\n+\n+    def __init__(\n+        self,\n+        client: AsyncSecretClient,\n+        purge_timeout: float = 120.0,\n+        polling_interval: float = 1.0,\n+    ) -\u003e None:\n+        if purge_timeout \u003c= 0 or polling_interval \u003c= 0:\n+            raise ValueError(\"timeouts and polling intervals must be positive\")\n+        self._client = client\n+        self._purge_timeout = purge_timeout\n+        self._polling_interval = polling_interval\n+\n+    async def rotate(\n+        self, name: str, value: str, expires_on: datetime\n+    ) -\u003e KeyVaultSecret:\n+        _validate_expiry(expires_on)\n+        try:\n+            # The aio client runs its deletion poller internally before returning.\n+            deleted = await self._client.delete_secret(name)\n+        except ResourceNotFoundError:\n+            deleted = None\n+\n+        if deleted is not None and deleted.recovery_id is not None:\n+            await self._client.purge_deleted_secret(name)\n+            await self._wait_until_purged(name)\n+\n+        return await self._client.set_secret(name, value, expires_on=expires_on)\n+\n+    async def _wait_until_purged(self, name: str) -\u003e None:\n+        deadline = time.monotonic() + self._purge_timeout\n+        while True:\n+            try:\n+                await self._client.get_deleted_secret(name)\n+            except ResourceNotFoundError:\n+                return\n+            if time.monotonic() \u003e= deadline:\n+                raise TimeoutError(\n+                    f\"Secret {name!r} was not purged within \"\n+                    f\"{self._purge_timeout:.1f} seconds\"\n+                )\n+            await asyncio.sleep(self._polling_interval)\n*** Add File: main.py\n+\"\"\"Demonstrate synchronous and asynchronous Key Vault configuration.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import os\n+from datetime import datetime, timedelta, timezone\n+\n+from azure_keyvault_config import (\n+    AsyncSecretCache,\n+    AsyncSecretRotator,\n+    SecretCache,\n+    SecretRotator,\n+    create_async_provider,\n+    create_sync_provider,\n+)\n+\n+REQUIRED_KEYS = (\"database-url\", \"api-key\", \"feature-flags\")\n+WARNING_WINDOW = timedelta(days=7)\n+\n+\n+def _print_cached_values(values: dict[str, str | None]) -\u003e None:\n+    for name, value in values.items():\n+        status = \"loaded\" if value is not None else \"missing\"\n+        print(f\"{name}: {status}\")\n+\n+\n+def _print_expiry_warnings(expiring: dict[str, datetime]) -\u003e None:\n+    for name, expires_on in expiring.items():\n+        print(f\"WARNING: {name} expires at {expires_on.isoformat()}\")\n+\n+\n+def run_sync_demo(rotation_value: str) -\u003e None:\n+    print(\"Synchronous implementation\")\n+    with create_sync_provider() as provider:\n+        cache = SecretCache(\n+            provider,\n+            required_keys=REQUIRED_KEYS,\n+            warning_window=WARNING_WINDOW,\n+        )\n+        cache.load_required()\n+        _print_cached_values({name: cache.get(name) for name in REQUIRED_KEYS})\n+\n+        cache.refresh(\"api-key\")\n+        cache.refresh_expiring()\n+        _print_expiry_warnings(cache.expiring_secrets())\n+\n+        rotator = SecretRotator(provider.client)\n+        rotator.rotate(\n+            \"api-key\",\n+            rotation_value,\n+            datetime.now(timezone.utc) + timedelta(days=90),\n+        )\n+        cache.refresh(\"api-key\")\n+\n+\n+async def run_async_demo(rotation_value: str) -\u003e None:\n+    print(\"Asynchronous implementation\")\n+    async with create_async_provider() as provider:\n+        cache = AsyncSecretCache(\n+            provider,\n+            required_keys=REQUIRED_KEYS,\n+            warning_window=WARNING_WINDOW,\n+        )\n+        await cache.load_required()\n+        values = {\n+            name: await cache.get(name)\n+            for name in REQUIRED_KEYS\n+        }\n+        _print_cached_values(values)\n+\n+        await cache.refresh(\"api-key\")\n+        await cache.refresh_expiring()\n+        _print_expiry_warnings(cache.expiring_secrets())\n+\n+        rotator = AsyncSecretRotator(provider.client)\n+        await rotator.rotate(\n+            \"api-key\",\n+            rotation_value,\n+            datetime.now(timezone.utc) + timedelta(days=90),\n+        )\n+        await cache.refresh(\"api-key\")\n+\n+\n+def main() -\u003e None:\n+    rotation_value = os.environ.get(\"ROTATED_SECRET_VALUE\")\n+    if not rotation_value:\n+        raise RuntimeError(\n+            \"ROTATED_SECRET_VALUE must contain the demo's replacement secret value\"\n+        )\n+\n+    run_sync_demo(rotation_value)\n+    asyncio.run(run_async_demo(rotation_value))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.25.1,\u003c2.0\n+azure-keyvault-secrets\u003e=4.11.2,\u003c5.0\n+aiohttp\u003e=3.13.2,\u003c4.0\n*** Add File: tests/__init__.py\n+\"\"\"Project tests.\"\"\"\n*** Add File: tests/test_keyvault_config.py\n+from __future__ import annotations\n+\n+import asyncio\n+import unittest\n+from datetime import datetime, timedelta, timezone\n+from types import SimpleNamespace\n+from unittest.mock import AsyncMock, Mock, patch\n+\n+from azure.core.exceptions import ResourceNotFoundError\n+\n+from azure_keyvault_config.cache import AsyncSecretCache, SecretCache\n+from azure_keyvault_config.provider import AsyncSecretProvider, SecretProvider\n+from azure_keyvault_config.rotation import AsyncSecretRotator, SecretRotator\n+\n+\n+def secret(\n+    name: str,\n+    value: str,\n+    expires_on: datetime | None = None,\n+    version: str = \"v1\",\n+) -\u003e SimpleNamespace:\n+    return SimpleNamespace(\n+        name=name,\n+        value=value,\n+        properties=SimpleNamespace(version=version, expires_on=expires_on),\n+    )\n+\n+\n+def not_found() -\u003e ResourceNotFoundError:\n+    return ResourceNotFoundError(\"not found\")\n+\n+\n+class ProviderTests(unittest.TestCase):\n+    def test_sync_provider_returns_default_and_passes_version(self) -\u003e None:\n+        client = Mock()\n+        client.get_secret.side_effect = not_found()\n+        provider = SecretProvider(client)\n+\n+        self.assertEqual(provider.get_secret(\"missing\", \"fallback\", \"v2\"), \"fallback\")\n+        client.get_secret.assert_called_once_with(\"missing\", \"v2\")\n+\n+    def test_sync_provider_exposes_expiry(self) -\u003e None:\n+        expiry = datetime.now(timezone.utc) + timedelta(days=30)\n+        client = Mock()\n+        client.get_secret.return_value = secret(\"item\", \"value\", expiry)\n+\n+        self.assertEqual(SecretProvider(client).get_expiry(\"item\"), expiry)\n+\n+    def test_async_provider_returns_default_and_passes_version(self) -\u003e None:\n+        async def run() -\u003e None:\n+            client = Mock()\n+            client.get_secret = AsyncMock(side_effect=not_found())\n+            provider = AsyncSecretProvider(client)\n+\n+            value = await provider.get_secret(\"missing\", \"fallback\", \"v2\")\n+\n+            self.assertEqual(value, \"fallback\")\n+            client.get_secret.assert_awaited_once_with(\"missing\", \"v2\")\n+\n+        asyncio.run(run())\n+\n+\n+class CacheTests(unittest.TestCase):\n+    def test_loads_required_keys_and_uses_cache(self) -\u003e None:\n+        provider = Mock()\n+        provider.get_secret_details.side_effect = [\n+            SimpleNamespace(\n+                name=\"one\",\n+                value=\"1\",\n+                version=\"v1\",\n+                expires_on=None,\n+                found=True,\n+            ),\n+            SimpleNamespace(\n+                name=\"two\",\n+                value=\"2\",\n+                version=\"v1\",\n+                expires_on=None,\n+                found=True,\n+            ),\n+        ]\n+        cache = SecretCache(provider, required_keys=(\"one\", \"two\"))\n+\n+        self.assertEqual(cache.load_required(), {\"one\": \"1\", \"two\": \"2\"})\n+        self.assertEqual(cache.get(\"one\"), \"1\")\n+        self.assertEqual(provider.get_secret_details.call_count, 2)\n+\n+    def test_get_automatically_refreshes_near_expiry(self) -\u003e None:\n+        soon = datetime.now(timezone.utc) + timedelta(days=1)\n+        later = datetime.now(timezone.utc) + timedelta(days=30)\n+        provider = Mock()\n+        provider.get_secret_details.side_effect = [\n+            SimpleNamespace(\n+                name=\"key\",\n+                value=\"old\",\n+                version=\"v1\",\n+                expires_on=soon,\n+                found=True,\n+            ),\n+            SimpleNamespace(\n+                name=\"key\",\n+                value=\"new\",\n+                version=\"v2\",\n+                expires_on=later,\n+                found=True,\n+            ),\n+        ]\n+        cache = SecretCache(provider, required_keys=(\"key\",))\n+        cache.load_required()\n+\n+        self.assertEqual(cache.get(\"key\"), \"new\")\n+        self.assertEqual(provider.get_secret_details.call_count, 2)\n+\n+    def test_async_cache_loads_and_refreshes(self) -\u003e None:\n+        async def run() -\u003e None:\n+            soon = datetime.now(timezone.utc) + timedelta(days=1)\n+            later = datetime.now(timezone.utc) + timedelta(days=30)\n+            provider = Mock()\n+            provider.get_secret_details = AsyncMock(\n+                side_effect=[\n+                    SimpleNamespace(\n+                        name=\"key\",\n+                        value=\"old\",\n+                        version=\"v1\",\n+                        expires_on=soon,\n+                        found=True,\n+                    ),\n+                    SimpleNamespace(\n+                        name=\"key\",\n+                        value=\"new\",\n+                        version=\"v2\",\n+                        expires_on=later,\n+                        found=True,\n+                    ),\n+                ]\n+            )\n+            cache = AsyncSecretCache(provider, required_keys=(\"key\",))\n+\n+            await cache.load_required()\n+            self.assertEqual(await cache.get(\"key\"), \"new\")\n+\n+        asyncio.run(run())\n+\n+\n+class RotationTests(unittest.TestCase):\n+    @patch(\"azure_keyvault_config.rotation.time.sleep\", return_value=None)\n+    def test_sync_rotation_waits_purges_then_sets(self, _sleep: Mock) -\u003e None:\n+        client = Mock()\n+        poller = Mock()\n+        deleted = SimpleNamespace(recovery_id=\"https://recovery\")\n+        poller.result.return_value = deleted\n+        client.begin_delete_secret.return_value = poller\n+        client.get_deleted_secret.side_effect = [deleted, not_found()]\n+        client.set_secret.return_value = secret(\"key\", \"new\")\n+        expiry = datetime.now(timezone.utc) + timedelta(days=30)\n+\n+        SecretRotator(client).rotate(\"key\", \"new\", expiry)\n+\n+        poller.wait.assert_called_once_with()\n+        client.purge_deleted_secret.assert_called_once_with(\"key\")\n+        client.set_secret.assert_called_once_with(\"key\", \"new\", expires_on=expiry)\n+\n+    def test_async_rotation_waits_purges_then_sets(self) -\u003e None:\n+        async def run() -\u003e None:\n+            client = Mock()\n+            deleted = SimpleNamespace(recovery_id=\"https://recovery\")\n+            client.delete_secret = AsyncMock(return_value=deleted)\n+            client.purge_deleted_secret = AsyncMock()\n+            client.get_deleted_secret = AsyncMock(side_effect=not_found())\n+            client.set_secret = AsyncMock(return_value=secret(\"key\", \"new\"))\n+            expiry = datetime.now(timezone.utc) + timedelta(days=30)\n+\n+            await AsyncSecretRotator(client).rotate(\"key\", \"new\", expiry)\n+\n+            client.delete_secret.assert_awaited_once_with(\"key\")\n+            client.purge_deleted_secret.assert_awaited_once_with(\"key\")\n+            client.set_secret.assert_awaited_once_with(\n+                \"key\", \"new\", expires_on=expiry\n+            )\n+\n+        asyncio.run(run())\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-secret-config-python-azure-skills-azure-skill-mcp-1513341870/azure_keyvault_config/__init__.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-secret-config-python-azure-skills-azure-skill-mcp-1513341870/azure_keyvault_config/__init__.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-secret-config-python-azure-skills-azure-skill-mcp-1513341870/azure_keyvault_config/__init__.py
@@ -1,0 +1,19 @@
+"""Azure Key Vault-backed application configuration."""
+
+from .cache import AsyncSecretCache, SecretCache
+from .factory import create_async_provider, create_sync_provider
+from .provider import AsyncSecretProvider, SecretDetails, SecretProvider
+from .rotation import AsyncSecretRotator, SecretRotator
+
+__all__ = [
+    "AsyncSecretCache",
+    "AsyncSecretProvider",
+    "AsyncSecretRotator",
+    "SecretCache",
+    "SecretDetails",
+    "SecretProvider",
+    "SecretRotator",
+    "create_async_provider",
+    "create_sync_provider",
+]
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-secret-config-python-azure-skills-azure-skill-mcp-1513341870/azure_keyvault_config/provider.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-secret-config-python-azure-skills-azure-skill-mcp-1513341870/azure_keyvault_config/provider.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-python-secret-config-python-azure-skills-azure-skill-mcp-1513341870/azure_keyvault_config/provider.py
@@ -1,0 +1,121 @@
+"""Synchronous and asynchronous Key Vault secret providers."""
+
+from __future__ import annotations
+
+from dataclasses import dataclass
+from datetime import datetime
+from typing import Optional
+
+from azure.core.exceptions import ResourceNotFoundError
+from azure.keyvault.secrets import SecretClient

... (truncated)
```

### ✅ 27. powershell — 5499ms

**Input:**

```json
{"command":"python -m unittest discover -s tests -v","description":"Run project unit tests","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_async_cache_loads_and_refreshes (test_keyvault_config.CacheTests.test_async_cache_loads_and_refreshes) ... ok
test_get_automatically_refreshes_near_expiry (test_keyvault_config.CacheTests.test_get_automatically_refreshes_near_expiry) ... ok
test_loads_required_keys_and_uses_cache (test_keyvault_config.CacheTests.test_loads_required_keys_and_uses_cache) ... ok
test_async_provider_returns_default_and_passes_version (test_keyvault_config.ProviderTests.test_async_provider_returns_default_and_passes_version) ... ok
test_sync_provider_exposes_expiry (test_keyvault_config.ProviderTests.test_sync_provider_exposes_expiry) ... ok
test_sync_provider_returns_default_and_passes_version (test_keyvault_config.ProviderTests.test_sync_provider_returns_default_and_passes_version) ... ok
test_async_rotation_waits_purges_then_sets (test_keyvault_config.RotationTests.test_async_rotation_waits_purges_then_sets) ... ok
test_sync_rotation_waits_purges_then_sets (test_keyvault_config.RotationTests.test_sync_rotation_waits_purges_then_sets) ... ok

----------------------------------------------------------------------
Ran 8 tests in 0.010s

OK
<shellId: 1 completed with exit code 0>
```

### ✅ 28. powershell — 4803ms

**Input:**

```json
{"command":"python -m compileall -q azure_keyvault_config main.py","description":"Compile Python project","initial_wait":120,"mode":"sync"}
```

**Output:**

```

<shellId: 2 completed with exit code 0>
```

### ✅ 29. glob — 1709ms

**Input:**

```json
{"pattern":"**/__pycache__/**"}
```

**Output:**

```
.\azure_keyvault_config\__pycache__\__init__.cpython-313.pyc
.\azure_keyvault_config\__pycache__\rotation.cpython-313.pyc
.\azure_keyvault_config\__pycache__\provider.cpython-313.pyc
.\azure_keyvault_config\__pycache__\factory.cpython-313.pyc
.\azure_keyvault_config\__pycache__\cache.cpython-313.pyc
.\__pycache__\main.cpython-313.pyc
.\tests\__pycache__\test_keyvault_config.cpython-313.pyc
```

### 🔧 30. powershell

**Input:**

```json
{"command":"$paths = @('.\\azure_keyvault_config\\__pycache__', '.\\tests\\__pycache__', '.\\__pycache__'); foreach ($path in $paths) { if (Test-Path -LiteralPath $path) { Remove-Item -LiteralPath $path -Recurse -Force } }","description":"Remove generated Python caches","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `azure_keyvault_config\__init__.py`
- `azure_keyvault_config\cache.py`
- `azure_keyvault_config\factory.py`
- `azure_keyvault_config\provider.py`
- `azure_keyvault_config\rotation.py`
- `main.py`
- `requirements.txt`
- `tests\__init__.py`
- `tests\test_keyvault_config.py`

## Copilot Response

I’ll build this as a local-only SDK project, using environment-based Key Vault configuration and no live Azure operations. I’m first grounding the implementation in current Azure Python SDK guidance and inspecting the workspace.

The SDK guidance confirms managed identity via `DefaultAzureCredential` and poller-based deletion. I’m narrowing the exact sync/async method contracts now, then I’ll create the modules and focused local tests with mocked clients.

The design now accounts for a critical Key Vault detail: after soft deletion, the name remains reserved. The rotation helper will wait for deletion, purge when soft-delete is active, wait until the deleted record disappears, and only then recreate the secret; the current async SDK performs deletion polling internally.

The project is implemented, including local mock-based coverage for versioned lookup, missing-secret defaults, expiry-driven refresh, and deletion/purge ordering. I’m running the smallest complete test pass and syntax compilation now.

## Reviewer Notes (LLM-as-Judge)

### Summary

The exception-handling criterion is fully met.

### Strengths

- Graceful default values for missing secrets
- Specific handling of ResourceNotFoundError
- Unexpected Azure SDK errors remain visible to callers

## Grader Results

- secret-config.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/12)
      - Secret versioning: retrieves specific version via `get_secret(name, version=version)`: Fail
      - Secret expiry: accesses `secret.properties.expires_on`: Pass
      - Configurable warning window for near-expiry detection: Pass
      - In-memory caching (e.g., `dict`) with bulk-load and single-key refresh: Pass
      - Secret rotation uses `begin_delete_secret()` as a long-running operation: Fail
      - Sync uses `LROPoller` — calls `.wait()` or `.result()` to wait for delete completion: Pass
      - Async uses the async poller — `await poller.wait()` for delete completion: Fail
      - Creates new secret only after delete completes (not concurrently): Fail
      - Async version uses `azure.keyvault.secrets.aio.SecretClient`: Pass
      - Returns a default value when secret is not found (`ResourceNotFoundError`): Pass
      - NOT using fire-and-forget `delete_secret()` without the long-running operation: Fail
      - NOT ignoring the poller and creating the new secret immediately: Fail
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Pass
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Pass
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 50% | 1.00 | 0.5000 | 7.7% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 15.4% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 15.4% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 15.4% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 15.4% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 15.4% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 15.4% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 6.5000** | **81.2%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-python-secret-config --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
