# Evaluation Report: key-vault-dp-java-secret-config

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 483.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-secret-config` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 19/22 |
| Duration | 483.3s |
| Timestamp | 2026-08-29T02:25:37Z |
| Files Generated | 12 |
| Event Count | 15705 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 217.4s |
| Review | 264.7s |
| **Total** | **483.3s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
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
| Token Usage | in=569293 out=12360 |
| Turn Count | 11 |

## Tool Calls

### ✅ 1. skill — 3008ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 7512ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven application using Azure Key Vault Secrets SDK, managed identity via DefaultAzureCredential, sync and async clients, caching, expiry inspection, and safe soft-delete rotation. Do not deploy or modify Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 9878ms

**Input:**

```json
{"intent":"Find official current Java SDK documentation for Azure Key Vault SecretClient and SecretAsyncClient: get latest and versioned secrets, expiry properties, handling 404 ResourceNotFoundException, beginDeleteSecret polling to completion, purge deleted secret and wait before recreating, and DefaultAzureCredential managed identity authentication.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 8738ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 7383ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code generation best practices for a local Java 17 Maven sample using Key Vault Secrets and managed identity, with no deployment.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 8068ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official current Java SDK docs for Key Vault secret operations and safe rotation.","parameters":{"query":"Azure Key Vault Secrets Java SecretClient SecretAsyncClient getSecret version getExpiresOn beginDeleteSecret poll completion purge soft delete recreate DefaultAzureCredential managed identity ResourceNotFoundException 404"}}
```

**Output:**

```
{"results":[{"title":"Azure Key Vault Secret client library for Java - version 4.11.0","content":"# Azure Key Vault Secret client library for Java - version 4.11.0\n## Troubleshooting\nSee our [troubleshooting guide](https://github.com/Azure/azure-sdk-for-java/blob/com.azure\u002Bazure-security-keyvault-secrets_4.11.0/sdk/keyvault/azure-security-keyvault-secrets/TROUBLESHOOTING.md) for details on how to diagnose various failure scenarios.\n### General\nAzure Key Vault Secret clients raise exceptions. For example, if you try to retrieve a secret after it is deleted a \u0060404\u0060 error is returned, indicating the resource was not found. In the following snippet, the error is handled gracefully by catching the exception and displaying additional information about the error.\n\u0060\u0060\u0060java\ntry {\n    secretClient.getSecret(\u0022\u003Cdeleted-secret-name\u003E\u0022);\n} catch (ResourceNotFoundException e) {\n    System.out.println(e.getMessage());\n}\n\u0060\u0060\u0060\n### Default HTTP Client\nAll client libraries by default use the Netty HTTP client. Adding the above dependency will automatically configure the client library to use the Netty HTTP client. Configuring or changing the HTTP client is detailed in the [HTTP clients wiki](https://learn.microsoft.com/en-us/azure/developer/java/sdk/http-client-pipeline#http-clients).\n### Default SSL library\nAll client libraries, by default, use the Tomcat-native Boring SSL library to enable native-level performance for SSL operations. The Boring SSL library is an Uber JAR containing native libraries for Linux / macOS / Windows, and provides better performance compared to the default SSL implementation within the JDK. For more information, including how to reduce the dependency size, refer to the [performance tuning](https://github.com/Azure/azure-sdk-for-java/blob/com.azure\u002Bazure-security-keyvault-secrets_4.11.0/docs/performance-tuning.md) section of the wiki.","contentUrl":"https://learn.microsoft.com/j
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 7725ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java code samples for Key Vault SecretClient and SecretAsyncClient.","parameters":{"language":"java","query":"Azure Key Vault Secrets Java SecretClient SecretAsyncClient getSecret version beginDeleteSecret waitForCompletion purgeDeletedSecret DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates how to use the Azure Key Vault Secret client library for Java to create, retrieve, and delete a secret in an Azure Key Vault instance, using DefaultAzureCredential for authentication.\npackage: com.azure.identity, com.azure.security.keyvault.secrets\nlanguage: java\n","codeSnippet":"package com.keyvault.secrets.quickstart;\r\n\r\nimport java.io.Console;\r\n\r\nimport com.azure.core.util.polling.SyncPoller;\r\nimport com.azure.identity.DefaultAzureCredentialBuilder;\r\n\r\nimport com.azure.security.keyvault.secrets.SecretClient;\r\nimport com.azure.security.keyvault.secrets.SecretClientBuilder;\r\nimport com.azure.security.keyvault.secrets.models.DeletedSecret;\r\nimport com.azure.security.keyvault.secrets.models.KeyVaultSecret;\r\n\r\npublic class App {\r\n    public static void main(String[] args) throws InterruptedException, IllegalArgumentException {\r\n        String keyVaultName = System.getenv(\u0022KEY_VAULT_NAME\u0022);\r\n        String keyVaultUri = \u0022https://\u0022 \u002B keyVaultName \u002B \u0022.vault.azure.net\u0022;\r\n\r\n        System.out.printf(\u0022key vault name = %s and key vault URI = %s \\n\u0022, keyVaultName, keyVaultUri);\r\n\r\n        SecretClient secretClient = new SecretClientBuilder()\r\n            .vaultUrl(keyVaultUri)\r\n            .credential(new DefaultAzureCredentialBuilder().build())\r\n            .buildClient();\r\n\r\n        Console con = System.console();\r\n\r\n        String secretName = \u0022mySecret\u0022;\r\n\r\n        System.out.println(\u0022Please provide the value of your secret \u003E \u0022);\r\n\r\n        String secretValue = con.readLine();\r\n\r\n        System.out.print(\u0022Creating a secret in \u0022 \u002B keyVaultName \u002B \u0022 called \u0027\u0022 \u002B secretName \u002B \u0022\u0027 with value \u0027\u0022 \u002B secretValue \u002B \u0022\u0027 ... \u0022);\r\n\r\n        secretClient.setSecret(new KeyVaultSecret(secretName, secretVa
... (truncated)
```

### ✅ 8. rg — 2491ms

**Input:**

```json
{"-n":true,"head_limit":80,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787970423010-copilot-tool-output-ae534e0528c845978a051eb5282e01c0.txt","C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787970423530-copilot-tool-output-2cf9caf8ae014dffa766cee5eeb43945.txt"],"pattern":"https://learn\\.microsoft\\.com[^\\\" ]+|beginDeleteSecret|purgeDeletedSecret|getSecret\\(|waitForCompletion|getExpiresOn"}
```

**Output:**

```
Output too large to read at once (117.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787970432657-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787970423010-copilot-tool-output-ae534e0528c845978a051eb5282e01c0.txt:1:{"results":[{"description":"description: Demonstrates how to use the Azure Key Vault Secret client library for Java to create, retrieve, and delete a secret in an Azure Key Vault instance, using DefaultAzureCredential for authentication.\npackage: com.azure.identity, com.azure.security.keyvault.secrets\nlanguage: java\n","codeSnippet":"package com.keyvault.secrets.quickstart;\r\n\r\nimport
```

### ✅ 9. web_fetch — 3469ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-sdk-bom</artifactId>
  <versioning>
    <latest>1.3.8</latest>
    <release>1.3.8</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.2.6</version>
      <version>1.2.7</version>
      <version>1.2.8</version>
      <version>1.2.9</version>
      <version>1.2.10</version>
      <version>1.2.11</version>
      <version>1.2.12</version>
      <version>1.2.13</version>
      <version>1.2.14</version>
      <version>1.2.15</version>
      <version>1.2.16</version>
      <version>1.2.17</version>
      <version>1.2.18</version>
      <version>1.2.19</version>
      <version>1.2.20</version>
      <version>1.2.21</version>
      <version>1.2.22</version>
      <version>1.2.23</version>
      <version>1.2.24</version>
      <version>1.2.25</version>
      <version>1.2.26</version>
      <version>1.2.27</version>
      <version>1.2.28</version>
      <version>1.2.29</version>
      <version>1.2.30</version>
      <version>1.2.31</version>
      <version>1.2.32</version>
      <version>1.2.33</version>
      <version>1.2.34</version>
      <version>1.2.35</version>
      <version>1.2.36</version>
      <version>1.2.37</version>
      <version>1.2.38</version>
      <version>1.3.0</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version
... (truncated)
```

### ✅ 10. web_fetch — 3134ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <versioning>
    <latest>1.19.0-beta.2</latest>
    <release>1.19.0-beta.2</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.0-preview.1</version>
      <version>1.0.0-preview.2</version>
      <version>1.0.0-preview.3</version>
      <version>1.0.0-preview.4</version>
      <version>1.0.0-preview.5</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.0.7</version>
      <version>1.0.8</version>
      <version>1.0.9</version>
      <version>1.1.0-beta.1</version>
      <version>1.1.0-beta.2</version>
      <version>1.1.0-beta.3</version>
      <version>1.1.0-beta.4</version>
      <version>1.1.0-beta.5</version>
      <version>1.1.0-beta.6</version>
      <version>1.1.0-beta.7</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.1.2</version>
      <version>1.1.3</version>
      <version>1.2.0-beta.1</version>
      <version>1.2.0-beta.2</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.3.0-beta.1</version>
      <version>1.3.0-beta.2</version>
      <version>1.3.0</version>
      <version>1.3.1</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version>1.3.4</version>
      <version>1.3.5</version>
      <version>1.3.6</version>
      <version>1.3.7</version>
      <version>1.4.0-beta.1</version>
      <version>1.4.0</version>
      <version>1.4.1</version>
      <version>
... (truncated)
```

### ✅ 11. web_fetch — 2812ms

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-security-keyvault-secrets</artifactId>
  <versioning>
    <latest>4.11.2</latest>
    <release>4.11.2</release>
    <versions>
      <version>4.0.0</version>
      <version>4.0.1</version>
      <version>4.1.0</version>
      <version>4.1.1</version>
      <version>4.1.2</version>
      <version>4.1.3</version>
      <version>4.1.4</version>
      <version>4.1.5</version>
      <version>4.2.0-beta.1</version>
      <version>4.2.0-beta.2</version>
      <version>4.2.0-beta.3</version>
      <version>4.2.0-beta.4</version>
      <version>4.2.0</version>
      <version>4.2.1</version>
      <version>4.2.2</version>
      <version>4.2.3</version>
      <version>4.2.4</version>
      <version>4.2.5</version>
      <version>4.2.6</version>
      <version>4.2.7</version>
      <version>4.2.8</version>
      <version>4.3.0-beta.1</version>
      <version>4.3.0-beta.2</version>
      <version>4.3.0-beta.3</version>
      <version>4.3.0-beta.4</version>
      <version>4.3.0-beta.5</version>
      <version>4.3.0-beta.6</version>
      <version>4.3.0</version>
      <version>4.3.1</version>
      <version>4.3.2</version>
      <version>4.3.3</version>
      <version>4.3.4</version>
      <version>4.3.5</version>
      <version>4.3.6</version>
      <version>4.3.7</version>
      <version>4.3.8</version>
      <version>4.4.0-beta.1</version>
      <version>4.4.0-beta.2</version>
      <version>4.4.0-beta.3</version>
      <version>4.4.0-beta.4</version>
      <version>4.4.0-beta.5</version>
      <version>4.4.0</version>
      <version>4.4.1</version>
      <version>4.4.2</version>
      <version>4.4.3</version>
      <version>4.4.4</version>
      <version>4.4.5</version>
      <version>4.4.6</ve
... (truncated)
```

### ✅ 12. apply_patch — 742ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-config-provider\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0-SNAPSHOT\u003c/version\u003e\n+    \u003cname\u003eKey Vault configuration provider\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+        \u003cazure.keyvault.secrets.version\u003e4.11.2\u003c/azure.keyvault.secrets.version\u003e\n+        \u003cslf4j.version\u003e2.0.17\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.keyvault.secrets.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.keyvaultconfig.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/keyvaultconfig/SecretSnapshot.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public record SecretSnapshot(\n+        String name,\n+        String value,\n+        String version,\n+        OffsetDateTime expiresOn,\n+        boolean found) {\n+\n+    public SecretSnapshot {\n+        Objects.requireNonNull(name, \"name\");\n+        Objects.requireNonNull(value, \"value\");\n+    }\n+\n+    public boolean expiresWithin(Duration warningWindow, Clock clock) {\n+        Objects.requireNonNull(warningWindow, \"warningWindow\");\n+        Objects.requireNonNull(clock, \"clock\");\n+        if (warningWindow.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return expiresOn != null\n+                \u0026\u0026 !expiresOn.isAfter(OffsetDateTime.now(clock).plus(warningWindow));\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+public interface SecretProvider {\n+    SecretSnapshot getSecret(String name, String defaultValue);\n+\n+    SecretSnapshot getSecret(String name, String version, String defaultValue);\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import reactor.core.publisher.Mono;\n+\n+public interface AsyncSecretProvider {\n+    Mono\u003cSecretSnapshot\u003e getSecret(String name, String defaultValue);\n+\n+    Mono\u003cSecretSnapshot\u003e getSecret(String name, String version, String defaultValue);\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/KeyVaultSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+import java.util.Objects;\n+\n+public final class KeyVaultSecretProvider implements SecretProvider {\n+    private final SecretClient client;\n+\n+    public KeyVaultSecretProvider(SecretClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    @Override\n+    public SecretSnapshot getSecret(String name, String defaultValue) {\n+        return getSecret(name, null, defaultValue);\n+    }\n+\n+    @Override\n+    public SecretSnapshot getSecret(String name, String version, String defaultValue) {\n+        validate(name, defaultValue);\n+        try {\n+            KeyVaultSecret secret = client.getSecret(name, version);\n+            return toSnapshot(secret);\n+        } catch (ResourceNotFoundException exception) {\n+            return missing(name, defaultValue);\n+        }\n+    }\n+\n+    private static SecretSnapshot toSnapshot(KeyVaultSecret secret) {\n+        return new SecretSnapshot(\n+                secret.getName(),\n+                secret.getValue(),\n+                secret.getProperties().getVersion(),\n+                secret.getProperties().getExpiresOn(),\n+                true);\n+    }\n+\n+    private static SecretSnapshot missing(String name, String defaultValue) {\n+        return new SecretSnapshot(name, defaultValue, null, null, false);\n+    }\n+\n+    private static void validate(String name, String defaultValue) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncKeyVaultSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+public final class AsyncKeyVaultSecretProvider implements AsyncSecretProvider {\n+    private final SecretAsyncClient client;\n+\n+    public AsyncKeyVaultSecretProvider(SecretAsyncClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    @Override\n+    public Mono\u003cSecretSnapshot\u003e getSecret(String name, String defaultValue) {\n+        return getSecret(name, null, defaultValue);\n+    }\n+\n+    @Override\n+    public Mono\u003cSecretSnapshot\u003e getSecret(String name, String version, String defaultValue) {\n+        validate(name, defaultValue);\n+        return client.getSecret(name, version)\n+                .map(AsyncKeyVaultSecretProvider::toSnapshot)\n+                .onErrorResume(ResourceNotFoundException.class,\n+                        exception -\u003e Mono.just(missing(name, defaultValue)));\n+    }\n+\n+    private static SecretSnapshot toSnapshot(KeyVaultSecret secret) {\n+        return new SecretSnapshot(\n+                secret.getName(),\n+                secret.getValue(),\n+                secret.getProperties().getVersion(),\n+                secret.getProperties().getExpiresOn(),\n+                true);\n+    }\n+\n+    private static SecretSnapshot missing(String name, String defaultValue) {\n+        return new SecretSnapshot(name, defaultValue, null, null, false);\n+    }\n+\n+    private static void validate(String name, String defaultValue) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/CachingSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+import java.util.concurrent.ConcurrentMap;\n+\n+public final class CachingSecretProvider {\n+    private final SecretProvider provider;\n+    private final Duration expiryWarningWindow;\n+    private final Clock clock;\n+    private final ConcurrentMap\u003cString, SecretSnapshot\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+    private final ConcurrentMap\u003cString, String\u003e defaults = new ConcurrentHashMap\u003c\u003e();\n+\n+    public CachingSecretProvider(SecretProvider provider, Duration expiryWarningWindow) {\n+        this(provider, expiryWarningWindow, Clock.systemUTC());\n+    }\n+\n+    CachingSecretProvider(SecretProvider provider, Duration expiryWarningWindow, Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.expiryWarningWindow = requireNonNegative(expiryWarningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public void loadRequired(Map\u003cString, String\u003e requiredSecrets) {\n+        Objects.requireNonNull(requiredSecrets, \"requiredSecrets\");\n+        requiredSecrets.forEach((name, defaultValue) -\u003e {\n+            validate(name, defaultValue);\n+            defaults.put(name, defaultValue);\n+            cache.put(name, provider.getSecret(name, defaultValue));\n+        });\n+    }\n+\n+    public SecretSnapshot get(String name) {\n+        SecretSnapshot cached = requireCached(name);\n+        return cached.expiresWithin(expiryWarningWindow, clock) ? refresh(name) : cached;\n+    }\n+\n+    public SecretSnapshot refresh(String name) {\n+        String defaultValue = requireDefault(name);\n+        SecretSnapshot refreshed = provider.getSecret(name, defaultValue);\n+        cache.put(name, refreshed);\n+        return refreshed;\n+    }\n+\n+    public List\u003cSecretSnapshot\u003e refreshExpiringSecrets() {\n+        List\u003cSecretSnapshot\u003e refreshed = new ArrayList\u003c\u003e();\n+        cache.forEach((name, secret) -\u003e {\n+            if (secret.expiresWithin(expiryWarningWindow, clock)) {\n+                refreshed.add(refresh(name));\n+            }\n+        });\n+        return List.copyOf(refreshed);\n+    }\n+\n+    public List\u003cSecretSnapshot\u003e expiringSecrets() {\n+        return cache.values().stream()\n+                .filter(secret -\u003e secret.expiresWithin(expiryWarningWindow, clock))\n+                .toList();\n+    }\n+\n+    private SecretSnapshot requireCached(String name) {\n+        SecretSnapshot secret = cache.get(name);\n+        if (secret == null) {\n+            throw new IllegalArgumentException(\"Secret is not loaded: \" + name);\n+        }\n+        return secret;\n+    }\n+\n+    private String requireDefault(String name) {\n+        String defaultValue = defaults.get(name);\n+        if (defaultValue == null) {\n+            throw new IllegalArgumentException(\"Secret is not loaded: \" + name);\n+        }\n+        return defaultValue;\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"expiryWarningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"expiryWarningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+\n+    private static void validate(String name, String defaultValue) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"Secret name must not be blank\");\n+        }\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncCachingSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+import java.util.concurrent.ConcurrentMap;\n+\n+public final class AsyncCachingSecretProvider {\n+    private final AsyncSecretProvider provider;\n+    private final Duration expiryWarningWindow;\n+    private final Clock clock;\n+    private final ConcurrentMap\u003cString, SecretSnapshot\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+    private final ConcurrentMap\u003cString, String\u003e defaults = new ConcurrentHashMap\u003c\u003e();\n+\n+    public AsyncCachingSecretProvider(AsyncSecretProvider provider, Duration expiryWarningWindow) {\n+        this(provider, expiryWarningWindow, Clock.systemUTC());\n+    }\n+\n+    AsyncCachingSecretProvider(\n+            AsyncSecretProvider provider,\n+            Duration expiryWarningWindow,\n+            Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.expiryWarningWindow = requireNonNegative(expiryWarningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public Mono\u003cVoid\u003e loadRequired(Map\u003cString, String\u003e requiredSecrets) {\n+        Objects.requireNonNull(requiredSecrets, \"requiredSecrets\");\n+        requiredSecrets.forEach(AsyncCachingSecretProvider::validate);\n+        defaults.putAll(requiredSecrets);\n+        return Flux.fromIterable(requiredSecrets.entrySet())\n+                .flatMap(entry -\u003e fetchAndCache(entry.getKey(), entry.getValue()))\n+                .then();\n+    }\n+\n+    public Mono\u003cSecretSnapshot\u003e get(String name) {\n+        return Mono.defer(() -\u003e {\n+            SecretSnapshot cached = cache.get(name);\n+            if (cached == null) {\n+                return Mono.error(new IllegalArgumentException(\"Secret is not loaded: \" + name));\n+            }\n+            return cached.expiresWithin(expiryWarningWindow, clock)\n+                    ? refresh(name)\n+                    : Mono.just(cached);\n+        });\n+    }\n+\n+    public Mono\u003cSecretSnapshot\u003e refresh(String name) {\n+        return Mono.defer(() -\u003e {\n+            String defaultValue = defaults.get(name);\n+            if (defaultValue == null) {\n+                return Mono.error(new IllegalArgumentException(\"Secret is not loaded: \" + name));\n+            }\n+            return fetchAndCache(name, defaultValue);\n+        });\n+    }\n+\n+    public Flux\u003cSecretSnapshot\u003e refreshExpiringSecrets() {\n+        return Flux.defer(() -\u003e Flux.fromIterable(cache.values())\n+                .filter(secret -\u003e secret.expiresWithin(expiryWarningWindow, clock))\n+                .flatMap(secret -\u003e refresh(secret.name())));\n+    }\n+\n+    public List\u003cSecretSnapshot\u003e expiringSecrets() {\n+        return cache.values().stream()\n+                .filter(secret -\u003e secret.expiresWithin(expiryWarningWindow, clock))\n+                .toList();\n+    }\n+\n+    private Mono\u003cSecretSnapshot\u003e fetchAndCache(String name, String defaultValue) {\n+        return provider.getSecret(name, defaultValue)\n+                .doOnNext(secret -\u003e cache.put(name, secret));\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"expiryWarningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"expiryWarningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+\n+    private static void validate(String name, String defaultValue) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"Secret name must not be blank\");\n+        }\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/KeyVaultClientFactory.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+\n+import java.net.URI;\n+import java.net.URISyntaxException;\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class KeyVaultClientFactory {\n+    public static final String VAULT_URL_ENV = \"KEY_VAULT_URL\";\n+    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = \"AZURE_CLIENT_ID\";\n+\n+    private KeyVaultClientFactory() {\n+    }\n+\n+    public static Clients fromEnvironment() {\n+        return fromEnvironment(System.getenv());\n+    }\n+\n+    static Clients fromEnvironment(Map\u003cString, String\u003e environment) {\n+        Objects.requireNonNull(environment, \"environment\");\n+        String vaultUrl = requireHttpsVaultUrl(environment.get(VAULT_URL_ENV));\n+        String clientId = environment.get(MANAGED_IDENTITY_CLIENT_ID_ENV);\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        if (clientId != null \u0026\u0026 !clientId.isBlank()) {\n+            credentialBuilder.clientId(clientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+        SecretClientBuilder clientBuilder = new SecretClientBuilder()\n+                .vaultUrl(vaultUrl)\n+                .credential(credential);\n+\n+        return new Clients(clientBuilder.buildClient(), clientBuilder.buildAsyncClient());\n+    }\n+\n+    private static String requireHttpsVaultUrl(String value) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(VAULT_URL_ENV + \" must be set\");\n+        }\n+        try {\n+            URI uri = new URI(value);\n+            if (!\"https\".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {\n+                throw new IllegalStateException(VAULT_URL_ENV + \" must be an HTTPS URL\");\n+            }\n+            return uri.toString();\n+        } catch (URISyntaxException exception) {\n+            throw new IllegalStateException(VAULT_URL_ENV + \" is not a valid URL\", exception);\n+        }\n+    }\n+\n+    public record Clients(SecretClient syncClient, SecretAsyncClient asyncClient) {\n+        public Clients {\n+            Objects.requireNonNull(syncClient, \"syncClient\");\n+            Objects.requireNonNull(asyncClient, \"asyncClient\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.core.util.polling.SyncPoller;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.DeletedSecret;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.time.Instant;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public final class SecretRotationHelper {\n+    private final SecretClient syncClient;\n+    private final SecretAsyncClient asyncClient;\n+    private final Duration purgeTimeout;\n+    private final Duration pollInterval;\n+\n+    public SecretRotationHelper(\n+            SecretClient syncClient,\n+            SecretAsyncClient asyncClient,\n+            Duration purgeTimeout,\n+            Duration pollInterval) {\n+        this.syncClient = Objects.requireNonNull(syncClient, \"syncClient\");\n+        this.asyncClient = Objects.requireNonNull(asyncClient, \"asyncClient\");\n+        this.purgeTimeout = requirePositive(purgeTimeout, \"purgeTimeout\");\n+        this.pollInterval = requirePositive(pollInterval, \"pollInterval\");\n+    }\n+\n+    public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {\n+        validateRotation(name, newValue, expiresOn);\n+        boolean deleted = deleteIfPresent(name);\n+        if (deleted) {\n+            Instant deadline = Instant.now().plus(purgeTimeout);\n+            purgeWithRetry(name, deadline);\n+            waitUntilPurged(name, deadline);\n+        }\n+        return syncClient.setSecret(new KeyVaultSecret(name, newValue).setExpiresOn(expiresOn));\n+    }\n+\n+    public Mono\u003cKeyVaultSecret\u003e rotateAsync(\n+            String name,\n+            String newValue,\n+            OffsetDateTime expiresOn) {\n+        validateRotation(name, newValue, expiresOn);\n+        return deleteIfPresentAsync(name)\n+                .flatMap(deleted -\u003e {\n+                    if (!deleted) {\n+                        return Mono.empty();\n+                    }\n+                    Instant deadline = Instant.now().plus(purgeTimeout);\n+                    return purgeWithRetryAsync(name, deadline)\n+                            .then(waitUntilPurgedAsync(name, deadline));\n+                })\n+                .then(asyncClient.setSecret(\n+                        new KeyVaultSecret(name, newValue).setExpiresOn(expiresOn)));\n+    }\n+\n+    private boolean deleteIfPresent(String name) {\n+        try {\n+            SyncPoller\u003cDeletedSecret, Void\u003e poller = syncClient.beginDeleteSecret(name);\n+            poller.waitForCompletion();\n+            return true;\n+        } catch (ResourceNotFoundException exception) {\n+            return false;\n+        }\n+    }\n+\n+    private Mono\u003cBoolean\u003e deleteIfPresentAsync(String name) {\n+        return Mono.defer(() -\u003e asyncClient.beginDeleteSecret(name).last().thenReturn(true))\n+                .onErrorResume(ResourceNotFoundException.class, exception -\u003e Mono.just(false));\n+    }\n+\n+    private void purgeWithRetry(String name, Instant deadline) {\n+        while (true) {\n+            try {\n+                syncClient.purgeDeletedSecret(name);\n+                return;\n+            } catch (ResourceNotFoundException exception) {\n+                waitOrThrow(name, deadline);\n+            }\n+        }\n+    }\n+\n+    private Mono\u003cVoid\u003e purgeWithRetryAsync(String name, Instant deadline) {\n+        return Mono.defer(() -\u003e asyncClient.purgeDeletedSecret(name))\n+                .onErrorResume(ResourceNotFoundException.class, exception -\u003e\n+                        delayOrError(name, deadline)\n+                                .then(purgeWithRetryAsync(name, deadline)));\n+    }\n+\n+    private void waitUntilPurged(String name, Instant deadline) {\n+        while (true) {\n+            try {\n+                syncClient.getDeletedSecret(name);\n+                waitOrThrow(name, deadline);\n+            } catch (ResourceNotFoundException exception) {\n+                return;\n+            }\n+        }\n+    }\n+\n+    private Mono\u003cVoid\u003e waitUntilPurgedAsync(String name, Instant deadline) {\n+        return Mono.defer(() -\u003e asyncClient.getDeletedSecret(name))\n+                .flatMap(secret -\u003e delayOrError(name, deadline)\n+                        .then(waitUntilPurgedAsync(name, deadline)))\n+                .onErrorResume(ResourceNotFoundException.class, exception -\u003e Mono.empty())\n+                .then();\n+    }\n+\n+    private void waitOrThrow(String name, Instant deadline) {\n+        if (!Instant.now().isBefore(deadline)) {\n+            throw purgeTimeout(name);\n+        }\n+        try {\n+            Thread.sleep(pollInterval.toMillis());\n+        } catch (InterruptedException exception) {\n+            Thread.currentThread().interrupt();\n+            throw new IllegalStateException(\"Interrupted while waiting for secret purge\", exception);\n+        }\n+    }\n+\n+    private Mono\u003cLong\u003e delayOrError(String name, Instant deadline) {\n+        if (!Instant.now().isBefore(deadline)) {\n+            return Mono.error(purgeTimeout(name));\n+        }\n+        return Mono.delay(pollInterval);\n+    }\n+\n+    private IllegalStateException purgeTimeout(String name) {\n+        return new IllegalStateException(\n+                \"Timed out waiting for secret '\" + name\n+                        + \"' to be purged. Purge protection may be enabled.\");\n+    }\n+\n+    private static void validateRotation(\n+            String name,\n+            String newValue,\n+            OffsetDateTime expiresOn) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        if (newValue == null || newValue.isBlank()) {\n+            throw new IllegalArgumentException(\"newValue must not be blank\");\n+        }\n+        Objects.requireNonNull(expiresOn, \"expiresOn\");\n+        if (!expiresOn.isAfter(OffsetDateTime.now())) {\n+            throw new IllegalArgumentException(\"expiresOn must be in the future\");\n+        }\n+    }\n+\n+    private static Duration requirePositive(Duration duration, String name) {\n+        Objects.requireNonNull(duration, name);\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(name + \" must be positive\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/Main.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Map;\n+import java.util.UUID;\n+\n+public final class Main {\n+    private static final Duration EXPIRY_WARNING_WINDOW = Duration.ofDays(7);\n+    private static final Map\u003cString, String\u003e REQUIRED_CONFIG = Map.of(\n+            \"database-url\", \"jdbc:postgresql://localhost/app\",\n+            \"external-api-key\", \"development-only-default\",\n+            \"feature-flag\", \"false\");\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        KeyVaultClientFactory.Clients clients = KeyVaultClientFactory.fromEnvironment();\n+        SecretRotationHelper rotationHelper = new SecretRotationHelper(\n+                clients.syncClient(),\n+                clients.asyncClient(),\n+                Duration.ofMinutes(5),\n+                Duration.ofSeconds(2));\n+\n+        runSyncDemo(clients, rotationHelper);\n+        runAsyncDemo(clients, rotationHelper);\n+    }\n+\n+    private static void runSyncDemo(\n+            KeyVaultClientFactory.Clients clients,\n+            SecretRotationHelper rotationHelper) {\n+        System.out.println(\"=== Synchronous provider ===\");\n+        KeyVaultSecretProvider provider = new KeyVaultSecretProvider(clients.syncClient());\n+        CachingSecretProvider cache =\n+                new CachingSecretProvider(provider, EXPIRY_WARNING_WINDOW);\n+\n+        cache.loadRequired(REQUIRED_CONFIG);\n+        REQUIRED_CONFIG.keySet().forEach(name -\u003e print(cache.get(name)));\n+\n+        System.out.println(\"Refreshing feature-flag\");\n+        print(cache.refresh(\"feature-flag\"));\n+\n+        cache.refreshExpiringSecrets();\n+        warnNearExpiry(cache.expiringSecrets());\n+\n+        rotationHelper.rotate(\n+                rotationSecretName(),\n+                \"sync-\" + UUID.randomUUID(),\n+                OffsetDateTime.now().plusDays(30));\n+        System.out.println(\"Synchronous rotation completed\");\n+    }\n+\n+    private static void runAsyncDemo(\n+            KeyVaultClientFactory.Clients clients,\n+            SecretRotationHelper rotationHelper) {\n+        System.out.println(\"=== Asynchronous provider ===\");\n+        AsyncKeyVaultSecretProvider provider =\n+                new AsyncKeyVaultSecretProvider(clients.asyncClient());\n+        AsyncCachingSecretProvider cache =\n+                new AsyncCachingSecretProvider(provider, EXPIRY_WARNING_WINDOW);\n+\n+        cache.loadRequired(REQUIRED_CONFIG)\n+                .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_CONFIG.keySet()))\n+                .concatMap(cache::get)\n+                .doOnNext(Main::print)\n+                .then(cache.refresh(\"feature-flag\"))\n+                .doOnNext(secret -\u003e System.out.println(\"Refreshed feature-flag\"))\n+                .doOnNext(Main::print)\n+                .thenMany(cache.refreshExpiringSecrets())\n+                .then(MonoSupport.fromRunnable(() -\u003e warnNearExpiry(cache.expiringSecrets())))\n+                .then(rotationHelper.rotateAsync(\n+                        rotationSecretName(),\n+                        \"async-\" + UUID.randomUUID(),\n+                        OffsetDateTime.now().plusDays(30)))\n+                .doOnSuccess(secret -\u003e System.out.println(\"Asynchronous rotation completed\"))\n+                .block();\n+    }\n+\n+    private static String rotationSecretName() {\n+        return System.getenv().getOrDefault(\"ROTATION_SECRET_NAME\", \"rotating-demo-secret\");\n+    }\n+\n+    private static void print(SecretSnapshot secret) {\n+        System.out.printf(\n+                \"%s = %s (version=%s, expires=%s, source=%s)%n\",\n+                secret.name(),\n+                secret.value(),\n+                secret.version(),\n+                secret.expiresOn(),\n+                secret.found() ? \"Key Vault\" : \"default\");\n+    }\n+\n+    private static void warnNearExpiry(Iterable\u003cSecretSnapshot\u003e secrets) {\n+        secrets.forEach(secret -\u003e System.out.printf(\n+                \"WARNING: secret '%s' expires on %s%n\",\n+                secret.name(),\n+                secret.expiresOn()));\n+    }\n+\n+    private static final class MonoSupport {\n+        private MonoSupport() {\n+        }\n+\n+        private static reactor.core.publisher.Mono\u003cVoid\u003e fromRunnable(Runnable runnable) {\n+            return reactor.core.publisher.Mono.fromRunnable(runnable);\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault configuration provider\n+\n+Java 17 sample with synchronous and asynchronous Key Vault secret providers,\n+in-memory caching, expiry-aware refresh, and soft-delete-safe rotation.\n+\n+## Configuration\n+\n+Set these environment variables on the Azure-hosted application:\n+\n+- `KEY_VAULT_URL`: required, for example `https://my-vault.vault.azure.net`\n+- `AZURE_CLIENT_ID`: optional client ID for a user-assigned managed identity;\n+  omit it to use the system-assigned managed identity\n+- `ROTATION_SECRET_NAME`: optional demo secret name; defaults to\n+  `rotating-demo-secret`\n+\n+Grant the managed identity only the Key Vault data-plane permissions it needs.\n+The rotation demo needs read, set, delete, and purge permissions. It cannot\n+recreate the same secret name when purge protection is enabled; in that case,\n+rotate by creating a new version with `setSecret` instead of delete-and-purge.\n+\n+## Build and run\n+\n+```powershell\n+mvn clean package\n+mvn exec:java\n+```\n+\n+The demo intentionally performs real secret deletion, purge, and creation\n+against the configured vault. Use a non-production secret and vault.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/pom.xml
@@ -1,0 +1,57 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>key-vault-config-provider</artifactId>
+    <version>1.0.0-SNAPSHOT</version>
+    <name>Key Vault configuration provider</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.identity.version>1.18.5</azure.identity.version>
+        <azure.keyvault.secrets.version>4.11.2</azure.keyvault.secrets.version>
+        <slf4j.version>2.0.17</slf4j.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-security-keyvault-secrets</artifactId>
+            <version>${azure.keyvault.secrets.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>org.slf4j</groupId>
+            <artifactId>slf4j-simple</artifactId>
+            <version>${slf4j.version}</version>
+            <scope>runtime</scope>
+        </depende
... (truncated)
```

### ✅ 13. powershell — 4466ms

**Input:**

```json
{"command":"mvn --no-transfer-progress clean package","description":"Compile and package the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ---------------< com.example:key-vault-config-provider >----------------
[INFO] Building Key Vault configuration provider 1.0.0-SNAPSHOT
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ key-vault-config-provider ---
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-config-provider ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ key-vault-config-provider ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java:[41,71] cannot find symbol
  symbol:   method setExpiresOn(java.time.OffsetDateTime)
  location: class com.azure.security.keyvault.secrets.models.KeyVaultSecret
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/
... (truncated)
```

### ✅ 14. apply_patch — 683ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java\n@@\n-        return syncClient.setSecret(new KeyVaultSecret(name, newValue).setExpiresOn(expiresOn));\n+        return syncClient.setSecret(secretWithExpiry(name, newValue, expiresOn));\n@@\n-                .then(asyncClient.setSecret(\n-                        new KeyVaultSecret(name, newValue).setExpiresOn(expiresOn)));\n+                .then(asyncClient.setSecret(secretWithExpiry(name, newValue, expiresOn)));\n@@\n     private static void validateRotation(\n             String name,\n             String newValue,\n             OffsetDateTime expiresOn) {\n@@\n         }\n     }\n \n+    private static KeyVaultSecret secretWithExpiry(\n+            String name,\n+            String value,\n+            OffsetDateTime expiresOn) {\n+        KeyVaultSecret secret = new KeyVaultSecret(name, value);\n+        secret.getProperties().setExpiresOn(expiresOn);\n+        return secret;\n+    }\n+\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java
@@ -38,7 +38,7 @@
             purgeWithRetry(name, deadline);
             waitUntilPurged(name, deadline);
         }
-        return syncClient.setSecret(new KeyVaultSecret(name, newValue).setExpiresOn(expiresOn));
+        return syncClient.setSecret(secretWithExpiry(name, newValue, expiresOn));
     }
 
     public Mono<KeyVaultSecret> rotateAsync(
@@ -55,8 +55,7 @@
                     return purgeWithRetryAsync(name, deadline)
                             .then(waitUntilPurgedAsync(name, deadline));
                 })
-                .then(asyncClient.setSecret(
-                        new KeyVaultSecret(name, newValue).setExpiresOn(expiresOn)));
+                .then(asyncClient.setSecret(secretWithExpiry(name, newValue, expiresOn)));
     }
 
     private boolean deleteIfPresent(String name) {
@@ -152,6 +151,15 @@
         }
     }
 
+    private static KeyVaultSecret secretWithExpiry(
+            String name,
+            String value,
+            OffsetDateTime expiresOn) {
+        KeyVaultSecret secret = new KeyVaultSecret(name, value);
+        secret.getProperties().setExpiresOn(expiresOn);
+        return secret;
+    }
+
     private static Duration requirePosi
... (truncated)
```

### ✅ 15. powershell — 5553ms

**Input:**

```json
{"command":"mvn --no-transfer-progress clean package","description":"Recompile and package the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ---------------< com.example:key-vault-config-provider >----------------
[INFO] Building Key Vault configuration provider 1.0.0-SNAPSHOT
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ key-vault-config-provider ---
[INFO] Deleting C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\target
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-config-provider ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ key-vault-config-provider ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ key-vault-config-provider ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ key-vault-config-provider ---
[INFO] No sources to compile
[INFO] 
[INFO] --- surefire:3.2.5:test (default-test) @ key-vault-config-provider ---
[INFO] No tests to run.
[INFO] 
[
... (truncated)
```

### ✅ 16. apply_patch — 664ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java\n@@\n         validateRotation(name, newValue, expiresOn);\n-        boolean deleted = deleteIfPresent(name);\n-        if (deleted) {\n+        boolean deletedSecretExists = deleteIfPresent(name) || deletedSecretExists(name);\n+        if (deletedSecretExists) {\n             Instant deadline = Instant.now().plus(purgeTimeout);\n             purgeWithRetry(name, deadline);\n             waitUntilPurged(name, deadline);\n@@\n         validateRotation(name, newValue, expiresOn);\n         return deleteIfPresentAsync(name)\n-                .flatMap(deleted -\u003e {\n-                    if (!deleted) {\n+                .flatMap(deleted -\u003e deleted\n+                        ? Mono.just(true)\n+                        : deletedSecretExistsAsync(name))\n+                .flatMap(deletedSecretExists -\u003e {\n+                    if (!deletedSecretExists) {\n                         return Mono.empty();\n                     }\n@@\n                 .onErrorResume(ResourceNotFoundException.class, exception -\u003e Mono.just(false));\n     }\n \n+    private boolean deletedSecretExists(String name) {\n+        try {\n+            syncClient.getDeletedSecret(name);\n+            return true;\n+        } catch (ResourceNotFoundException exception) {\n+            return false;\n+        }\n+    }\n+\n+    private Mono\u003cBoolean\u003e deletedSecretExistsAsync(String name) {\n+        return asyncClient.getDeletedSecret(name)\n+                .map(secret -\u003e true)\n+                .onErrorResume(ResourceNotFoundException.class, exception -\u003e Mono.just(false));\n+    }\n+\n*** Update File: src/main/java/com/example/keyvaultconfig/Main.java\n@@\n         System.out.printf(\n-                \"%s = %s (version=%s, expires=%s, source=%s)%n\",\n+                \"%s loaded (valueLength=%d, version=%s, expires=%s, source=%s)%n\",\n                 secret.name(),\n-                secret.value(),\n+                secret.value().length(),\n                 secret.version(),\n                 secret.expiresOn(),\n                 secret.found() ? \"Key Vault\" : \"default\");\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983/src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java
@@ -32,8 +32,8 @@
 
     public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {
         validateRotation(name, newValue, expiresOn);
-        boolean deleted = deleteIfPresent(name);
-        if (deleted) {
+        boolean deletedSecretExists = deleteIfPresent(name) || deletedSecretExists(name);
+        if (deletedSecretExists) {
             Instant deadline = Instant.now().plus(purgeTimeout);
             purgeWithRetry(name, deadline);
             waitUntilPurged(name, deadline);
@@ -47,8 +47,11 @@
             OffsetDateTime expiresOn) {
         validateRotation(name, newValue, expiresOn);
         return deleteIfPresentAsync(name)
-                .flatMap(deleted -> {
-                    if (!deleted) {
+                .flatMap(deleted -> deleted
+                        ? Mono.just(true)
+                        : deletedSecretExistsAsync(name))
+                .flatMap(deletedSecretExists -> {
+                    if (!deletedSecretExists) {
                         return Mono.empty();
                     }
                     Instant deadline = Instant.now().plus(purgeTimeout);
@@ -73,6 +76,21 @@
                 .onErrorResume(Reso
... (truncated)
```

### ✅ 17. powershell — 5496ms

**Input:**

```json
{"command":"mvn --no-transfer-progress clean package","description":"Build the finalized Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ---------------< com.example:key-vault-config-provider >----------------
[INFO] Building Key Vault configuration provider 1.0.0-SNAPSHOT
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ key-vault-config-provider ---
[INFO] Deleting C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\target
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-config-provider ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ key-vault-config-provider ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ key-vault-config-provider ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-1484660983\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ key-vault-config-provider ---
[INFO] No sources to compile
[INFO] 
[INFO] --- surefire:3.2.5:test (default-test) @ key-vault-config-provider ---
[INFO] No tests to run.
[INFO] 
[
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\keyvaultconfig\AsyncCachingSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\AsyncKeyVaultSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\AsyncSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\CachingSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\KeyVaultClientFactory.java`
- `src\main\java\com\example\keyvaultconfig\KeyVaultSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\Main.java`
- `src\main\java\com\example\keyvaultconfig\SecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\SecretRotationHelper.java`
- `src\main\java\com\example\keyvaultconfig\SecretSnapshot.java`

## Copilot Response

I’ll ground the implementation in current Azure Java SDK patterns, then scaffold and compile the project locally without contacting Azure.

The implementation will use managed identity directly, cache structured secret metadata (value/version/expiry), and purge-and-poll after soft deletion before recreating a name.

Compilation exposed the expected Azure model API shape; I’ve corrected that and am tightening rotation recovery so an interrupted prior delete/purge cannot block recreation.

Created the Java 17 Maven project with:

- Sync and async Key Vault secret providers, including version lookup, defaults for missing secrets, and expiry metadata.
- Expiry-aware sync/async in-memory caches with startup bulk loading and refresh operations.
- Managed identity client factory using `KEY_VAULT_URL` and optional `AZURE_CLIENT_ID`.
- Soft-delete-safe sync/async rotation that waits for deletion, purges, polls until fully removed, and then recreates the secret.
- A complete sync-first, async-second demo in `Main`.
- Azure SDK dependencies and run instructions in `pom.xml` and `README.md`.

Rotation requires purge permission and cannot recreate the same name when Key Vault purge protection is enabled; the README documents that limitation. References: [Key Vault Secrets Java SDK](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme), [Azure Identity Java SDK](https://learn.microsoft.com/java/api/overview/azure/identity-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

1/1 criteria passed.

### Strengths

- Client lifecycle handling is appropriate for the Azure Key Vault SecretClient and SecretAsyncClient APIs.

## Grader Results

- secret-config.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (10/10)
      - Secret versioning: retrieves specific version via `getSecret(name, version)`: Pass
      - Secret expiry: accesses `properties().getExpiresOn()` on `SecretProperties`: Pass
      - Configurable warning window for near-expiry detection: Pass
      - In-memory caching (e.g., `ConcurrentHashMap`) with bulk-load and single-key refresh: Pass
      - Secret rotation uses `beginDeleteSecret()` as a long-running operation: Pass
      - Sync uses `SyncPoller` to wait for delete completion: Pass
      - Async uses `PollerFlux` to wait for delete completion: Pass
      - Creates new secret only after delete completes (not concurrently): Pass
      - Returns a default value when secret is not found (does not crash): Pass
      - NOT using fire-and-forget `deleteSecret()` without waiting for completion: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Pass
  - Client Builder Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Pass
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Fail
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Fail
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.0000** | **76.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-java-secret-config --config java-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
