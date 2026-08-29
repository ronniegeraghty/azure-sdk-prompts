# Evaluation Report: resource-manager-mp-java-rg-crud

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 378.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-java-rg-crud` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/19 |
| Duration | 378.4s |
| Timestamp | 2026-08-29T02:52:55Z |
| Files Generated | 3 |
| Event Count | 5941 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 150.8s |
| Review | 226.4s |
| **Total** | **378.4s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-app-onboard-prereq |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=421807 out=5492 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 2947ms

**Input:**

```json
{"skill":"azure-app-onboard-prereq"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure App Onboard Prereq — Repository Evaluation

Evaluate a user's repository for build health, app completeness, and Azure deployment feasibility — before infrastructure planning. Produces per-component verdicts (PASS/WARN/FAIL) consumed by downstream phases.

> **Orchestrator relationship:** Called by `azure-app-onboard` at Step 3, or standalone for code readiness checks. When called by orchestrator, return control to `azure-app-onboard` after writing artifacts — do NOT invoke downstream phases directly.

Phase 1 of 4 in AppOnboard pipeline. Session: `.copilot-azure/sessions/{session-id}/`. Reads `context.json`. Writes `components[]`, `repo{}`, `detectedInfra[]`. Produces `prereq-output.json`. Schema: [`prereq-schemas.ts`](references/prereq-schemas.ts) — `PrereqOutput`, `BuildRequirements`. Direct entry supported.

## When NOT to Use

| Signal | Redirect |
|--------|----------|
| Validate infrastructure (Bicep/TF/azure.yaml) | **azure-validate** |
| Generate IaC | **azure-prepare** |
| End-to-end idea-to-production | **azure-app-onboard** |
| Run `azd up` or deploy | **azure-deploy** |

## Rules

> ⛔ **ABSOLUTE PROHIBITION — `npm install`, `npm test`, `npx jest`, `pytest`, and ALL install/build/test commands are NEVER allowed.**
> Under NO circumstances may you run `npm install`, `npm test`, `npx jest`, `pip install`, `pytest`, `dotnet build`, `dotnet restore`, `dotnet test`, `go mod download`, `cargo build`, or ANY package-manager install, build, or test command during the prereq phase. Do NOT run test suites to verify code — check for test config files statically instead. The prereq phase is read-only evaluation + static-only verification.
> **ONLY exception — two sanctioned contexts, both consent-gated:** (a) code the agent **modified** during migration/remediation (see [remediation-protocol.md](references/remediation-protocol.md) step 6), or (b) code the agent **wrote** from scratch on the zero-code path (see [
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 5925ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Generate a local-only Java Maven sample using the modern Azure Resource Manager management-plane SDK and DefaultAzureCredential to create, list, get, tag, and delete a resource group with robust exception handling."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. glob — 5579ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-3746155777","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. view — 6253ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-app-onboard-prereq\\references\\zero-code-path.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard-prereq/references/zero-code-path.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard-prereq/references/zero-code-path.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard-prereq/references/zero-code-path.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-app-onboard-prereq/references/zero-code-path.md
@@ -1,29 +1,29 @@
 # Zero-Code Path
 
 When the workspace is empty (no project files, no Dockerfile), prereq scaffolds a starter project before evaluating.
 
 ## Step 0: Check Existing Context
 
 ⛔ **Before asking anything, check `context.json.intent.userPrompt`.** The orchestrator (or the user's original message on direct entry) already captured what the user wants. If `userPrompt` contains an app description (e.g., "build me a todo app", "REST API with auth", "a notes app with three tabs"), treat it as the answer to "what are you building?" and skip directly to Step 2 (recommend stack). Only ask Step 1's question if `userPrompt` is vague ("help me get started", "deploy to Azure") or empty.
 
 ## Flow
 
 1. **Ask what they want to build** (skip if Step 0 resolved it): *"What kind of app are you building?"* (e.g., "a todo app", "REST API"). Users know what they want, not what stack to use.
 2. **Recommend a stack** based on the app description (from `userPrompt` or Step 1): *"For a REST API with a database, I'd suggest Node.js with a framework like Express or Fastify. Sound good?"* User can accept or override.
 3. ⛔ **Confirm before scaffolding** via `ask_user`.
 4. **Scaffold minimal starter project** — ⛔ **You MUST read [subagent-starter-scaffold.md](subagent-starter-scaffold.md).** Your
... (truncated)
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 3586ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java SDK code-generation guidance for a local Maven sample that manages resource groups with DefaultAzureCredential.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 4625ms

**Input:**

```json
{"command":"search","intent":"Find official Microsoft documentation for Java Azure Resource Manager authentication with DefaultAzureCredential and resource group create/list/get/update tags/delete operations using com.azure.resourcemanager.AzureResourceManager.","parameters":{"query":"Java AzureResourceManager DefaultAzureCredential resource groups create list get update tags delete azure-resourcemanager"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. web_fetch — 4281ms

**Input:**

```json
{"max_length":5000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure.resourcemanager%22%20AND%20a:%22azure-resourcemanager%22\u0026rows=1\u0026wt=json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-identity%22&rows=1&wt=json:
{"responseHeader":{"status":0,"QTime":0,"params":{"q":"g:\"com.azure\" AND a:\"azure-identity\"","core":"","indent":"off","spellcheck":"true","fl":"id,g,a,latestVersion,p,ec,repositoryId,text,timestamp,versionCount","start":"","spellcheck.count":"5","sort":"score desc,timestamp desc,g asc,a asc","rows":"1","wt":"json","version":"2.2"}},"response":{"numFound":1,"start":0,"docs":[{"id":"com.azure:azure-identity","g":"com.azure","a":"azure-identity","latestVersion":"1.16.2","repositoryId":"central","p":"jar","timestamp":1749530875060,"versionCount":117,"text":["com.azure","azure-identity","-sources.jar.sha256","-sources.jar","-readme.md",".pom","-javadoc.jar","-changelog.md","-javadoc.jar.sha256",".jar","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"],"ec":["-sources.jar.sha256","-sources.jar","-readme.md",".pom","-javadoc.jar","-changelog.md","-javadoc.jar.sha256",".jar","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"]}]},"spellcheck":{"suggestions":[]}}
```

### ✅ 8. web_fetch — 3923ms

**Input:**

```json
{"max_length":5000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-identity%22\u0026rows=1\u0026wt=json"}
```

**Output:**

```
Content type application/json cannot be simplified to markdown. Here is the raw content:
Contents of https://search.maven.org/solrsearch/select?q=g:%22com.azure.resourcemanager%22%20AND%20a:%22azure-resourcemanager%22&rows=1&wt=json:
{"responseHeader":{"status":0,"QTime":1,"params":{"q":"g:\"com.azure.resourcemanager\" AND a:\"azure-resourcemanager\"","core":"","indent":"off","spellcheck":"true","fl":"id,g,a,latestVersion,p,ec,repositoryId,text,timestamp,versionCount","start":"","spellcheck.count":"5","sort":"score desc,timestamp desc,g asc,a asc","rows":"1","wt":"json","version":"2.2"}},"response":{"numFound":1,"start":0,"docs":[{"id":"com.azure.resourcemanager:azure-resourcemanager","g":"com.azure.resourcemanager","a":"azure-resourcemanager","latestVersion":"2.51.0","repositoryId":"central","p":"jar","timestamp":1748264186466,"versionCount":58,"text":["com.azure.resourcemanager","azure-resourcemanager","-sources.jar","-sources.jar.sha256","-readme.md",".pom","-javadoc.jar","-changelog.md",".jar","-javadoc.jar.sha256","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"],"ec":["-sources.jar","-sources.jar.sha256","-readme.md",".pom","-javadoc.jar","-changelog.md",".jar","-javadoc.jar.sha256","-changelog.md.sha256",".pom.sha256","-readme.md.sha256",".jar.sha256"]}]},"spellcheck":{"suggestions":[]}}
```

### ✅ 9. azure-documentation (via azure) — 2420ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve official Java SDK examples for Azure Resource Manager resource group CRUD and authentication.","parameters":{"language":"java","query":"Azure Java AzureResourceManager DefaultAzureCredential resourceGroups define withRegion create list getByName update withTag deleteByName"}}
```

**Output:**

```
{"results":[{"title":"Azure Resource Manager support client library for Java - version 1.1.0","content":"# Azure Resource Manager support client library for Java - version 1.1.0\n## Getting started\n### Prerequisites\n1. [Java Development Kit (JDK)](https://learn.microsoft.com/en-us/azure/developer/java/fundamentals/) with version 8 or above\n2. [Azure Subscription](https://azure.microsoft.com/free/)\n### Adding the package to your product\n\u0060\u0060\u0060xml\n\u003Cdependency\u003E\n    \u003CgroupId\u003Ecom.azure.resourcemanager\u003C/groupId\u003E\n    \u003CartifactId\u003Eazure-resourcemanager-support\u003C/artifactId\u003E\n    \u003Cversion\u003E1.1.0\u003C/version\u003E\n\u003C/dependency\u003E\n\u0060\u0060\u0060\n### Include the recommended packages\nAzure Management Libraries require a \u0060TokenCredential\u0060 implementation for authentication and an \u0060HttpClient\u0060 implementation for HTTP client.\n[Azure Identity](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-support_1.1.0/sdk/identity/azure-identity) and [Azure Core Netty HTTP](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-support_1.1.0/sdk/core/azure-core-http-netty) packages provide the default implementation.\n### Authentication\nMicrosoft Entra ID token authentication relies on the [credential class](https://github.com/Azure/azure-sdk-for-java/tree/azure-resourcemanager-support_1.1.0/sdk/identity/azure-identity#credentials) from [Azure Identity](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-support_1.1.0/sdk/identity/azure-identity) package.\nAzure subscription ID can be configured via \u0060AZURE_SUBSCRIPTION_ID\u0060 environment variable.\nAssuming the use of the \u0060DefaultAzureCredential\u0060 credential class, the client can be authenticated using the following code:\n\u0060\u0060\u0060java\nAzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);\nTokenCredential credential = new DefaultAzureCred
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 2063ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Java Azure Resource Manager authentication and resource group management documentation.","parameters":{"query":"Java Azure Resource Manager SDK authenticate DefaultAzureCredential resource groups create delete"}}
```

**Output:**

```
{"results":[{"description":"description: Authenticates an AzureResourceManager instance using DefaultAzureCredential and AzureProfile, enabling management of Azure resources with the specified subscription.\npackage: azure-identity, azure-resourcemanager\nlanguage: java\n","codeSnippet":"AzureResourceManager azureResourceManager = AzureResourceManager.authenticate(\r\n        new DefaultAzureCredentialBuilder().build(),\r\n        new AzureProfile(AzureEnvironment.AZURE))\r\n    .withDefaultSubscription();","link":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview#examples","language":"java"},{"description":"description: Creates and authenticates an ArmClient instance using DefaultAzureCredential, retrieves the default subscription, and creates or updates a resource group with a specified name and location.\npackage: Azure.Identity, Azure.ResourceManager\nlanguage: c\n","codeSnippet":"ArmClient armClient = new ArmClient(new DefaultAzureCredential());\r\nSubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();\r\nResourceGroupCollection rgCollection = subscription.GetResourceGroups();\r\n// With the collection, we can create a new resource group with an specific name\r\nstring rgName = \u0022sample-rg\u0022;\r\nAzureLocation location = AzureLocation.WestUS2;\r\nArmOperation\u003CResourceGroupResource\u003E resourceGroupLro = await rgCollection.CreateOrUpdateAsync(WaitUntil.Completed, rgName, new ResourceGroupData(location));\r\nResourceGroupResource resourceGroup = lro.Value;","link":"https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.loadtesting-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Creates and authenticates an ArmClient instance using DefaultAzureCredential, retrieves the default subscription, and creates a new resource group with a specified name and location.\npackage: Azure.Identity, Azure.ResourceManager\nlanguage: c\n","codeSnippet":"ArmClient armC
... (truncated)
```

### ✅ 11. web_fetch — 2843ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/resourcemanager/azure-resourcemanager/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Resource Manager client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-07-08T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-resourcemanager, resourcemanager
ms.service: resourcemanager
locale: en-us
document_id: fdf479c2-6820-e586-078b-c36f0b321310
document_version_independent_id: 3da92e19-2d03-2ee2-06a2-45c73eac39e1
updated_at: 2026-07-08T11:06:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/resourcemanager-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/a9a3f5064eb0adb3919961b6ed30a5ab950c7634/docs-ref-services/latest/resourcemanager-readme.md
git_commit_id: a9a3f5064eb0adb3919961b6ed30a5ab950c7634
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 239
asset_id: api/overview/azure/resourcemanager-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/resourcemanager-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6
... (truncated)
```

### ✅ 12. web_fetch — 2430ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 13. web_fetch — 2029ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/resourcemanager/azure-resourcemanager/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure.resourcemanager</groupId>
  <artifactId>azure-resourcemanager</artifactId>
  <versioning>
    <latest>2.63.0</latest>
    <release>2.63.0</release>
    <versions>
      <version>1.0.0-hybrid</version>
      <version>2.0.0-beta.1</version>
      <version>2.0.0-beta.2</version>
      <version>2.0.0-beta.3</version>
      <version>2.0.0-beta.4</version>
      <version>2.0.0-beta.5</version>
      <version>2.0.0</version>
      <version>2.1.0</version>
      <version>2.2.0</version>
      <version>2.3.0</version>
      <version>2.4.0</version>
      <version>2.5.0</version>
      <version>2.6.0</version>
      <version>2.7.0</version>
      <version>2.8.0</version>
      <version>2.9.0</version>
      <version>2.10.0</version>
      <version>2.11.0</version>
      <version>2.12.0</version>
      <version>2.13.0</version>
      <version>2.14.0</version>
      <version>2.15.0</version>
      <version>2.16.0</version>
      <version>2.17.0</version>
      <version>2.18.0</version>
      <version>2.19.0</version>
      <version>2.20.0</version>
      <version>2.21.0</version>
      <version>2.22.0</version>
      <version>2.23.0</version>
      <version>2.24.0</version>
      <version>2.25.0</version>
      <version>2.26.0</version>
      <version>2.27.0</version>
      <version>2.28.0</version>
      <version>2.29.0</version>
      <version>2.30.0</version>
      <version>2.31.0</version>
      <version>2.32.0</version>
      <version>2.33.0</version>
      <version>2.34.0</version>
      <version>2.35.0</version>
      <version>2.36.0</version>
      <version>2.37.0</version>
      <version>2.38.0</version>
      <version>2.39.0</version>
      <version>2.40.0</version>
      <version>2.41.0</version>
      <versi
... (truncated)
```

### ✅ 14. web_fetch — 1088ms

**Input:**

```json
{"max_length":4000,"start_index":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
5.4</version>
      <version>1.16.0-beta.1</version>
      <version>1.16.0</version>
      <version>1.16.1</version>
      <version>1.16.2</version>
      <version>1.16.3</version>
      <version>1.17.0-beta.1</version>
      <version>1.17.0</version>
      <version>1.18.0</version>
      <version>1.18.1</version>
      <version>1.18.2</version>
      <version>1.18.3</version>
      <version>1.18.4</version>
      <version>1.18.5</version>
      <version>1.19.0-beta.1</version>
      <version>1.19.0-beta.2</version>
    </versions>
    <lastUpdated>20260824185316</lastUpdated>
  </versioning>
</metadata>

```

### ✅ 15. apply_patch — 727ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eresource-group-manager\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure.resourcemanager\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-resourcemanager\u003c/artifactId\u003e\n+            \u003cversion\u003e2.63.0\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e1.18.5\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.azure.ResourceGroupManagerApp\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/azure/ResourceGroupManagerApp.java\n+package com.example.azure;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.AzureException;\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.management.AzureEnvironment;\n+import com.azure.core.management.profile.AzureProfile;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.resourcemanager.AzureResourceManager;\n+import com.azure.resourcemanager.resources.models.ResourceGroup;\n+import com.azure.resourcemanager.resources.models.ResourceGroups;\n+\n+import java.util.Locale;\n+import java.util.Map;\n+import java.util.UUID;\n+import java.util.logging.Level;\n+import java.util.logging.Logger;\n+\n+public final class ResourceGroupManagerApp {\n+    private static final Logger LOGGER = Logger.getLogger(ResourceGroupManagerApp.class.getName());\n+    private static final String REGION = \"eastus\";\n+    private static final String TAG_NAME = \"managed-by\";\n+    private static final String TAG_VALUE = \"azure-java-sdk\";\n+\n+    private ResourceGroupManagerApp() {\n+    }\n+\n+    public static void main(String[] args) {\n+        try {\n+            String subscriptionId = requiredEnvironmentVariable(\"AZURE_SUBSCRIPTION_ID\");\n+            String resourceGroupName = resourceGroupName();\n+\n+            AzureResourceManager azure = createClient(subscriptionId);\n+            manageResourceGroup(azure.resourceGroups(), resourceGroupName);\n+        } catch (IllegalArgumentException exception) {\n+            LOGGER.log(Level.SEVERE, \"Invalid configuration: {0}\", exception.getMessage());\n+            System.exit(2);\n+        } catch (ClientAuthenticationException exception) {\n+            LOGGER.log(Level.SEVERE,\n+                \"Azure authentication failed. Check the DefaultAzureCredential configuration.\", exception);\n+            System.exit(3);\n+        } catch (HttpResponseException exception) {\n+            LOGGER.log(Level.SEVERE, () -\u003e String.format(\n+                \"Azure Resource Manager request failed with status %d: %s\",\n+                exception.getResponse().getStatusCode(),\n+                exception.getMessage()));\n+            System.exit(4);\n+        } catch (AzureException exception) {\n+            LOGGER.log(Level.SEVERE, \"Azure SDK operation failed.\", exception);\n+            System.exit(5);\n+        }\n+    }\n+\n+    private static AzureResourceManager createClient(String subscriptionId) {\n+        AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);\n+        TokenCredential credential = new DefaultAzureCredentialBuilder()\n+            .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())\n+            .build();\n+\n+        return AzureResourceManager.authenticate(credential, profile)\n+            .withSubscription(subscriptionId);\n+    }\n+\n+    private static void manageResourceGroup(ResourceGroups resourceGroups, String resourceGroupName) {\n+        boolean created = false;\n+        boolean deleted = false;\n+\n+        try {\n+            LOGGER.log(Level.INFO, \"Creating resource group {0} in {1}.\",\n+                new Object[] {resourceGroupName, REGION});\n+            ResourceGroup createdGroup = resourceGroups.define(resourceGroupName)\n+                .withRegion(REGION)\n+                .create();\n+            created = true;\n+            logDetails(\"Created\", createdGroup);\n+\n+            LOGGER.info(\"Resource groups in the subscription:\");\n+            resourceGroups.list().forEach(group -\u003e\n+                LOGGER.log(Level.INFO, \"  {0} ({1})\",\n+                    new Object[] {group.name(), group.regionName()}));\n+\n+            ResourceGroup fetchedGroup = resourceGroups.getByName(resourceGroupName);\n+            if (fetchedGroup == null) {\n+                throw new AzureException(\"Created resource group could not be retrieved: \" + resourceGroupName);\n+            }\n+            logDetails(\"Fetched\", fetchedGroup);\n+\n+            LOGGER.log(Level.INFO, \"Adding tag {0}={1}.\", new Object[] {TAG_NAME, TAG_VALUE});\n+            ResourceGroup taggedGroup = fetchedGroup.update()\n+                .withTag(TAG_NAME, TAG_VALUE)\n+                .apply();\n+            logDetails(\"Updated\", taggedGroup);\n+\n+            LOGGER.log(Level.INFO, \"Deleting resource group {0}.\", resourceGroupName);\n+            resourceGroups.deleteByName(resourceGroupName);\n+            deleted = true;\n+            LOGGER.log(Level.INFO, \"Deleted resource group {0}.\", resourceGroupName);\n+        } finally {\n+            if (created \u0026\u0026 !deleted) {\n+                cleanupResourceGroup(resourceGroups, resourceGroupName);\n+            }\n+        }\n+    }\n+\n+    private static void cleanupResourceGroup(ResourceGroups resourceGroups, String resourceGroupName) {\n+        try {\n+            LOGGER.log(Level.WARNING,\n+                \"An earlier operation failed; attempting to delete resource group {0}.\", resourceGroupName);\n+            resourceGroups.deleteByName(resourceGroupName);\n+        } catch (AzureException cleanupException) {\n+            LOGGER.log(Level.SEVERE,\n+                \"Cleanup failed. Delete the resource group manually: \" + resourceGroupName,\n+                cleanupException);\n+        }\n+    }\n+\n+    private static void logDetails(String operation, ResourceGroup resourceGroup) {\n+        Map\u003cString, String\u003e tags = resourceGroup.tags();\n+        LOGGER.log(Level.INFO,\n+            \"{0} resource group: name={1}, region={2}, provisioningState={3}, tags={4}\",\n+            new Object[] {\n+                operation,\n+                resourceGroup.name(),\n+                resourceGroup.regionName(),\n+                resourceGroup.provisioningState(),\n+                tags\n+            });\n+    }\n+\n+    private static String requiredEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(name + \" must be set.\");\n+        }\n+        return value.trim();\n+    }\n+\n+    private static String resourceGroupName() {\n+        String configuredName = System.getenv(\"RESOURCE_GROUP_NAME\");\n+        if (configuredName != null \u0026\u0026 !configuredName.isBlank()) {\n+            return configuredName.trim();\n+        }\n+\n+        String suffix = UUID.randomUUID().toString().substring(0, 8).toLowerCase(Locale.ROOT);\n+        return \"java-sdk-rg-\" + suffix;\n+    }\n+}\n*** Add File: README.md\n+# Azure Resource Group Manager (Java)\n+\n+This Maven console application uses the modern Azure management-plane SDK to:\n+\n+1. Authenticate with `DefaultAzureCredential`.\n+2. Create a resource group in `eastus`.\n+3. List the subscription's resource groups.\n+4. Retrieve the created resource group.\n+5. add the `managed-by=azure-java-sdk` tag.\n+6. Delete the resource group.\n+\n+If an operation fails after creation, the application attempts to delete the resource group in a\n+`finally` block and logs any cleanup failure.\n+\n+## Requirements\n+\n+- JDK 17+\n+- Maven 3.9+\n+- An Azure identity with permission to manage resource groups in the target subscription\n+- `AZURE_SUBSCRIPTION_ID` set to the target subscription ID\n+\n+`DefaultAzureCredential` supports local developer credentials and workload/managed identity\n+credentials. For service-principal authentication, set `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and\n+`AZURE_CLIENT_SECRET`. Do not store credentials in source control.\n+\n+Optionally set `RESOURCE_GROUP_NAME`. If omitted, the application generates a name such as\n+`java-sdk-rg-a1b2c3d4`.\n+\n+## Dependencies\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure.resourcemanager\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-resourcemanager\u003c/artifactId\u003e\n+    \u003cversion\u003e2.63.0\u003c/version\u003e\n+\u003c/dependency\u003e\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+    \u003cversion\u003e1.18.5\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+## Build and run\n+\n+PowerShell:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:RESOURCE_GROUP_NAME = \"my-java-sdk-rg\" # optional\n+mvn compile\n+mvn exec:java\n+```\n+\n+Running the application creates and deletes a real Azure resource group. Review the subscription\n+and resource-group name before executing it.\n+\n+## References\n+\n+- [Azure Resource Manager client library for Java](https://learn.microsoft.com/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable)\n+- [Authentication with Azure SDK for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview)\n*** Add File: .gitignore\n+target/\n+.idea/\n+*.iml\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-3746155777/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-3746155777/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-3746155777/pom.xml
@@ -1,0 +1,47 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example.azure</groupId>
+    <artifactId>resource-group-manager</artifactId>
+    <version>1.0.0</version>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure.resourcemanager</groupId>
+            <artifactId>azure-resourcemanager</artifactId>
+            <version>2.63.0</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>1.18.5</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-compiler-plugin</artifactId>
+                <version>3.14.1</version>
+            </plugin>
+            <plugin>
+                <groupId>org.codehaus.mojo</groupId>
+                <artifactId>exec-maven-plugin</artifactId>
+                <version>3.5.1</version>
+          
... (truncated)
```

### ✅ 16. powershell — 13067ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile the Maven application without running it","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example.azure:resource-group-manager >--------------
[INFO] Building resource-group-manager 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ resource-group-manager ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-3746155777\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ resource-group-manager ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  3.490 s
[INFO] Finished at: 2026-08-29T10:55:20+08:00
[INFO] ------------------------------------------------------------------------
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\azure\ResourceGroupManagerApp.java`

## Copilot Response

I’ll inspect the starter project and align the implementation with current Azure Java SDK guidance before making changes.

The workspace is empty, so I’m creating a focused Maven console application rather than modifying an existing project. I’m also pinning compatible SDK versions from current first-party/package metadata.

The implementation will use the current stable `azure-resourcemanager` 2.63.0 and `azure-identity` 1.18.5, with subscription and resource-group naming supplied through environment variables and cleanup attempted if any intermediate operation fails.

The project is now implemented. I’m compiling it only—without running the application—so no Azure resource is created or modified.

Implemented the Java Maven application using:

- `azure-resourcemanager` **2.63.0**
- `azure-identity` **1.18.5**
- `DefaultAzureCredential`
- Create, list, retrieve, tag, and delete operations
- Specific handling for configuration, authentication, HTTP, and Azure SDK failures
- Best-effort cleanup if an operation fails after resource-group creation

Key files:

- `pom.xml`
- `src\main\java\com\example\azure\ResourceGroupManagerApp.java`
- `README.md`

Run with:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:RESOURCE_GROUP_NAME = "my-java-sdk-rg" # optional
mvn compile
mvn exec:java
```

Running `mvn exec:java` performs real Azure create/delete operations.

References: [Azure Resource Manager SDK for Java](https://learn.microsoft.com/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable) and [Azure SDK authentication](https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The application does not leave any AutoCloseable Azure SDK client unclosed.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - Maven dependency for `azure-resourcemanager` and `azure-identity`: Pass
      - `AzureResourceManager.authenticate()` with credential and profile: Pass
      - `resourceGroups().define().withRegion().create()`: Pass
      - `resourceGroups().list()` iteration: Pass
      - `resourceGroups().getByName()` for details: Pass
      - Tag management via `update().withTag().apply()`: Pass
      - `resourceGroups().deleteByName()` for cleanup: Pass
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
  - Client Builder Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Fail
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
  - Service-Specific Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Pass
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
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.0000** | **76.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-java-rg-crud --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
