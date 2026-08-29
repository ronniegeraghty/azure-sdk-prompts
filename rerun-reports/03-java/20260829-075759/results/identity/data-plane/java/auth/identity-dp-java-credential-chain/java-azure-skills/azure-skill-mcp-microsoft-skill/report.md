# Evaluation Report: identity-dp-java-credential-chain

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 456.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-credential-chain` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 24/26 |
| Duration | 456.2s |
| Timestamp | 2026-08-29T00:59:37Z |
| Files Generated | 15 |
| Event Count | 16882 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 243.9s |
| Review | 211.1s |
| **Total** | **456.2s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=1162344 out=10993 |
| Turn Count | 20 |

## Tool Calls

### ✅ 1. skill — 2469ms

**Input:**

```json
{"skill":"azure-identity-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for Java

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.15.0</version>
</dependency>
```

## Key Concepts

| Credential | Use Case |
|------------|----------|
| `DefaultAzureCredential` | **Recommended** - Works in dev and production |
| `ManagedIdentityCredential` | Azure-hosted apps (App Service, Functions, VMs) |
| `EnvironmentCredential` | CI/CD pipelines with env vars |
| `ClientSecretCredential` | Service principals with secret |
| `ClientCertificateCredential` | Service principals with certificate |
| `AzureCliCredential` | Local dev using `az login` |
| `InteractiveBrowserCredential` | Interactive login flow |
| `DeviceCodeCredential` | Headless device authentication |

## DefaultAzureCredential (Recommended)

The `DefaultAzureCredential` tries multiple authentication methods in order. See [DefaultAzureCredential overview](https://aka.ms/azsdk/java/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

```java
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

// Simple usage
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();

// Use with any Azure client
BlobServiceClient blobClient = new BlobServiceClientBuilder()
    .endpoint("https://<storage-account>.blob.core.windows.net")
    .credential(credential)
    .buildClient();

KeyClient keyClient = new KeyClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();
```

### Configure DefaultAzureCredential

```java
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId("<user-assigned-identity-client-id>")  // For user-assigned MI
    .tenantId("<tenant-id
... (truncated)
```

### ✅ 2. view — 6549ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
@@ -1,495 +1,495 @@
 # Azure Identity library for Java - Examples
 
 Comprehensive code examples for the Azure Identity library for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [DefaultAzureCredential](#defaultazurecredential)
 - [ChainedTokenCredential](#chainedtokencredential)
 - [ClientSecretCredential](#clientsecretcredential)
 - [ClientCertificateCredential](#clientcertificatecredential)
 - [ManagedIdentityCredential](#managedidentitycredential)
 - [EnvironmentCredential](#environmentcredential)
 - [InteractiveBrowserCredential](#interactivebrowsercredential)
 - [DeviceCodeCredential](#devicecodecredential)
 - [AzureCliCredential](#azureclicredential)
 - [Using Credentials with Azure SDK Clients](#using-credentials-with-azure-sdk-clients)
 
 ## Maven Dependency
 
 ```xml
 <!-- Using Azure SDK BOM (recommended) -->
 <dependencyManagement>
     <dependencies>
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-sdk-bom</artifactId>
             <version>1.2.29</version>
             <type>pom</type>
             <scope>import</scope>
         </dependency>
     </dependencies>
 </dependencyManagement>
 
 <dependencies>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-identity</artifactId>
     </dependency>
 </dependencies>
 
 <!-- Or direct depend
... (truncated)
```

### ✅ 3. glob — 6171ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. glob — 5768ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 20543ms

**Input:**

```json
{"command":"get_code_generation_best_practices","intent":"Generate a local-only Java 17 Maven sample using Azure Identity custom credential chains for local development, Azure Pipelines and generic CI, managed identity and workload identity, with CAE-enabled token requests and precise authentication diagnostics."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 20855ms

**Input:**

```json
{"command":"search","intent":"Find current Java Azure Identity documentation for custom chained credentials, AzurePipelinesCredential builder configuration, and Continuous Access Evaluation enablement on TokenRequestContext or credential builders."}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 7. web_search (via github-mcp-server) — 21191ms

**Input:**

```json
{"query":"site:learn.microsoft.com java azure identity AzurePipelinesCredentialBuilder systemAccessToken serviceConnectionId tenantId clientId Continuous Access Evaluation TokenRequestContext setCaeEnabled"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s how the Azure Identity Java library supports the combination of the classes and settings you mentioned for use in Azure Pipelines, including systemAccessToken, serviceConnectionId, tenantId, clientId, Continuous Access Evaluation, and TokenRequestContext’s setCaeEnabled:\n\n## 1. AzurePipelinesCredentialBuilder Configuration\n\nThe `AzurePipelinesCredentialBuilder` is a builder for `AzurePipelinesCredential`. Use it to tie together your Azure service connection in a pipeline. You must provide:\n- `clientId`\n- `tenantId`\n- `serviceConnectionId`\n- `systemAccessToken` (e.g., from the environment variable `SYSTEM_ACCESSTOKEN` in Azure Pipelines)\n\n**Sample:**\n```java\nString systemAccessToken = System.getenv(\"SYSTEM_ACCESSTOKEN\");\n\nAzurePipelinesCredential credential = new AzurePipelinesCredentialBuilder()\n    .clientId(clientId)\n    .tenantId(tenantId)\n    .serviceConnectionId(serviceConnectionId)\n    .systemAccessToken(systemAccessToken)\n    .build();\n```\nThis credential can now be used any place that accepts a `TokenCredential` in the Azure SDK for Java【3:0†source】【3:1†source】.\n\n## 2. TokenRequestContext and Continuous Access Evaluation (CAE)\n\nThe `TokenRequestContext` class supports the `setCaeEnabled(boolean enableCae)` method to request a CAE-enabled access token.\n\n**Sample:**\n```java\nTokenRequestContext requestContext = new TokenRequestContext()\n    .addScopes(\"https://management.azure.com/.default\")\n    .setCaeEnabled(true); // Enables CAE for the token request\n```\n- `setCaeEnabled(true)` will ask for a CAE-enabled token if the resource and clients support it, improving real-time authorization enforcement (Continuous Access Evaluation)【3:3†source】.\n\n## 3. Acquiring a Token\n\nYou can acquire a token (async or sync) using the credential and token request context:\n\n**Async:**\n```java\nMono<AccessToken> tokenMono = credential.getToken(requestContext);\n```\n\n**Sy
... (truncated)
```

### ✅ 8. azure-get_azure_bestpractices (via azure) — 2602ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java SDK code-generation best practices for a Java 17 Azure Identity credential-chain sample that never provisions resources.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3350ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Search Azure Identity Java custom credential chain, Azure Pipelines credential, managed identity, workload identity, and CAE token request documentation.","parameters":{"query":"Azure Identity Java ChainedTokenCredential AzurePipelinesCredential ManagedIdentityCredential WorkloadIdentityCredential Continuous Access Evaluation TokenRequestContext setCaeEnabled"}}
```

**Output:**

```
{"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## DefaultAzureCredential overview\n7. \u0060VisualStudioCodeCredential\u0060\n8. \u0060WorkloadIdentityCredential\u0060\nImportant\nThe \u0060AZURE_TOKEN_CREDENTIALS\u0060 environment variable supports individual credential names in \u0060azure-identity\u0060 package versions 1.17.0 and later.\nTo ensure the environment variable is defined and set to a supported string, call method [requireEnvVars](https://learn.microsoft.com/en-us/java/api/com.azure.identity.defaultazurecredentialbuilder#com-azure-identity-defaultazurecredentialbuilder-requireenvvars%28com-azure-identity-azureidentityenvvars-%29) as follows:\n\u0060\u0060\u0060java\nDefaultAzureCredential credential = new DefaultAzureCredentialBuilder()\n    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)\n    .build();\n\u0060\u0060\u0060\n## ChainedTokenCredential overview\n[ChainedTokenCredential](https://learn.microsoft.com/en-us/java/api/com.azure.identity.chainedtokencredential) is an empty chain to which you add credentials to suit your app\u0027s needs. For example:\nThe preceding code example creates a tailored credential chain comprised of two development-time credentials. \u0060AzureCliCredential\u0060 is attempted first, followed by \u0060IntelliJCredential\u0060, if necessary. In graphical form, the chain looks like this:\n![Diagram that shows authentication flow for a ChainedTokenCredential instance that is composed of the Azure CLI and IntelliJ credentials.](https://learn.microsoft.com/azure/developer/java/sdk/media/mermaidjs/chained-token-credential-auth-flow.svg)\nTip\nFor improved performance, optimize credential ordering in \u0060ChainedTokenCredential\u0060 from most to least used credential.","contentUrl":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains#defaultazurecredential-overview"},{"title
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 2951ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve official Java samples for AzurePipelinesCredentialBuilder and CAE-enabled TokenRequestContext.","parameters":{"language":"java","query":"AzurePipelinesCredentialBuilder TokenRequestContext setCaeEnabled ChainedTokenCredentialBuilder Java"}}
```

**Output:**

```
{"results":[{"description":"description: Creates a ChainedTokenCredential instance by chaining together AzureCliCredential and IntelliJCredential, allowing sequential authentication attempts using these credentials.\npackage: com.azure.identity\nlanguage: java\n","codeSnippet":"import com.azure.identity.AzureCliCredential;\r\nimport com.azure.identity.AzureCliCredentialBuilder;\r\nimport com.azure.identity.ChainedTokenCredential;\r\nimport com.azure.identity.ChainedTokenCredentialBuilder;\r\nimport com.azure.identity.IntelliJCredential;\r\nimport com.azure.identity.IntelliJCredentialBuilder;\r\n\r\n// Code omitted for brevity\r\n\r\nAzureCliCredential cliCredential = new AzureCliCredentialBuilder()\r\n    .build();\r\nIntelliJCredential ijCredential = new IntelliJCredentialBuilder()\r\n    .build();\r\n\r\nChainedTokenCredential credential = new ChainedTokenCredentialBuilder()\r\n    .addLast(cliCredential)\r\n    .addLast(ijCredential)\r\n    .build();","link":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains#chainedtokencredential-overview","language":"java"},{"description":"description: Imports necessary classes for Azure authentication and communication management client initialization in Java.\npackage: com.azure.core.credential, com.azure.core.management, com.azure.identity, com.azure.resourcemanager.communication\nlanguage: java\n","codeSnippet":"import com.azure.core.credential.TokenCredential;\r\nimport com.azure.core.management.AzureEnvironment;\r\nimport com.azure.core.management.profile.AzureProfile;\r\nimport com.azure.identity.DefaultAzureCredentialBuilder;\r\nimport com.azure.resourcemanager.communication.CommunicationManager;","link":"https://learn.microsoft.com/azure/communication-services/quickstarts/email/manage-suppression-list-management-sdks?pivots=programming-language-java#initialize-the-management-client","language":"java"},{"description":"description: Creates a DefaultAzureCredential instance that require
... (truncated)
```

### ✅ 11. rg — 1976ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787965261489-copilot-tool-output-f59e9a54581748adacae6cf2c2f5a249.txt","pattern":"AzurePipelinesCredential|setCaeEnabled|WorkloadIdentityCredential|url|link"}
```

**Output:**

```
Output too large to read at once (54.1 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787965269318-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787965261489-copilot-tool-output-f59e9a54581748adacae6cf2c2f5a249.txt:{"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## DefaultAzureCredential overview\n7. \u0060VisualStudioCodeCredential\u0060\n8. \u0060WorkloadIdentityCredential\u0060\nImportant\nThe \u0060AZURE_TOKEN_CREDENTIALS\u0060 environment variable supports individual credential names in \u0060azu
```

### ✅ 12. powershell — 2321ms

**Input:**

```json
{"command":"java -version; mvn -version","description":"Check Java and Maven versions"}
```

**Output:**

```
openjdk version "25.0.1" 2025-10-21 LTS
OpenJDK Runtime Environment Microsoft-12574223 (build 25.0.1+8-LTS)
OpenJDK 64-Bit Server VM Microsoft-12574223 (build 25.0.1+8-LTS, mixed mode, sharing)
Apache Maven 3.9.11 (3e54c93a704957b63ee3494413a2b544fd3d825b)
Maven home: C:\apache-maven-3.9.11
Java version: 25.0.1, vendor: Microsoft, runtime: C:\Program Files\Microsoft\jdk-25.0.1.8-hotspot
Default locale: en_US, platform encoding: UTF-8
OS name: "windows 11", version: "10.0", arch: "amd64", family: "windows"
<shellId: 0 completed with exit code 0>
```

### ✅ 13. apply_patch — 817ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-credential-chains\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure credential chains\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+        \u003cjunit.version\u003e5.12.2\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.azureauth.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/azureauth/DeploymentEnvironment.java\n+package com.example.azureauth;\n+\n+public enum DeploymentEnvironment {\n+    DEVELOPMENT,\n+    CI,\n+    PRODUCTION\n+}\n*** Add File: src/main/java/com/example/azureauth/EnvironmentDetector.java\n+package com.example.azureauth;\n+\n+import java.util.Locale;\n+import java.util.Map;\n+import java.util.Set;\n+\n+public final class EnvironmentDetector {\n+    private static final Set\u003cString\u003e CI_MARKERS = Set.of(\n+        \"CI\",\n+        \"TF_BUILD\",\n+        \"GITHUB_ACTIONS\",\n+        \"GITLAB_CI\",\n+        \"JENKINS_URL\",\n+        \"BUILD_SOURCESDIRECTORY\",\n+        \"PIPELINE_WORKSPACE\"\n+    );\n+\n+    private static final Set\u003cString\u003e PRODUCTION_MARKERS = Set.of(\n+        \"IDENTITY_ENDPOINT\",\n+        \"MSI_ENDPOINT\",\n+        \"IMDS_ENDPOINT\",\n+        \"WEBSITE_INSTANCE_ID\",\n+        \"CONTAINER_APP_NAME\",\n+        \"AZURE_FEDERATED_TOKEN_FILE\"\n+    );\n+\n+    private final Map\u003cString, String\u003e environment;\n+\n+    public EnvironmentDetector() {\n+        this(System.getenv());\n+    }\n+\n+    EnvironmentDetector(Map\u003cString, String\u003e environment) {\n+        this.environment = Map.copyOf(environment);\n+    }\n+\n+    public DeploymentEnvironment detect() {\n+        String override = environment.get(\"APP_ENVIRONMENT\");\n+        if (hasText(override)) {\n+            return parseOverride(override);\n+        }\n+        if (hasAnyMarker(CI_MARKERS)) {\n+            return DeploymentEnvironment.CI;\n+        }\n+        if (hasAnyMarker(PRODUCTION_MARKERS) || hasWorkloadIdentityConfiguration()) {\n+            return DeploymentEnvironment.PRODUCTION;\n+        }\n+        return DeploymentEnvironment.DEVELOPMENT;\n+    }\n+\n+    private boolean hasAnyMarker(Set\u003cString\u003e names) {\n+        return names.stream().map(environment::get).anyMatch(EnvironmentDetector::hasText);\n+    }\n+\n+    private boolean hasWorkloadIdentityConfiguration() {\n+        return hasText(environment.get(\"AZURE_TENANT_ID\"))\n+            \u0026\u0026 hasText(environment.get(\"AZURE_CLIENT_ID\"))\n+            \u0026\u0026 hasText(environment.get(\"AZURE_FEDERATED_TOKEN_FILE\"));\n+    }\n+\n+    private static DeploymentEnvironment parseOverride(String value) {\n+        return switch (value.trim().toLowerCase(Locale.ROOT)) {\n+            case \"dev\", \"development\", \"local\" -\u003e DeploymentEnvironment.DEVELOPMENT;\n+            case \"ci\", \"pipeline\" -\u003e DeploymentEnvironment.CI;\n+            case \"prod\", \"production\" -\u003e DeploymentEnvironment.PRODUCTION;\n+            default -\u003e throw new IllegalArgumentException(\n+                \"APP_ENVIRONMENT must be one of: dev, ci, or production\"\n+            );\n+        };\n+    }\n+\n+    private static boolean hasText(String value) {\n+        return value != null \u0026\u0026 !value.isBlank();\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/CredentialSelection.java\n+package com.example.azureauth;\n+\n+import com.azure.core.credential.TokenCredential;\n+\n+import java.util.Objects;\n+\n+public record CredentialSelection(\n+    TokenCredential credential,\n+    String strategy,\n+    boolean caeEnabled\n+) {\n+    public CredentialSelection {\n+        Objects.requireNonNull(credential, \"credential\");\n+        Objects.requireNonNull(strategy, \"strategy\");\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/AzureCredentialFactory.java\n+package com.example.azureauth;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.AzureCliCredentialBuilder;\n+import com.azure.identity.AzureDeveloperCliCredentialBuilder;\n+import com.azure.identity.AzurePipelinesCredentialBuilder;\n+import com.azure.identity.ChainedTokenCredentialBuilder;\n+import com.azure.identity.EnvironmentCredentialBuilder;\n+import com.azure.identity.IntelliJCredentialBuilder;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.identity.WorkloadIdentityCredentialBuilder;\n+\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class AzureCredentialFactory {\n+    public static final String MANAGED_IDENTITY_CLIENT_ID = \"AZURE_MANAGED_IDENTITY_CLIENT_ID\";\n+    public static final String SERVICE_CONNECTION_ID = \"AZURE_SERVICE_CONNECTION_ID\";\n+\n+    private final Map\u003cString, String\u003e environment;\n+\n+    public AzureCredentialFactory() {\n+        this(System.getenv());\n+    }\n+\n+    AzureCredentialFactory(Map\u003cString, String\u003e environment) {\n+        this.environment = Map.copyOf(environment);\n+    }\n+\n+    public CredentialSelection create(DeploymentEnvironment deploymentEnvironment, boolean enableCae) {\n+        Objects.requireNonNull(deploymentEnvironment, \"deploymentEnvironment\");\n+        return switch (deploymentEnvironment) {\n+            case DEVELOPMENT -\u003e developmentCredential(enableCae);\n+            case CI -\u003e ciCredential(enableCae);\n+            case PRODUCTION -\u003e productionCredential(enableCae);\n+        };\n+    }\n+\n+    private CredentialSelection developmentCredential(boolean enableCae) {\n+        TokenCredential credential = new ChainedTokenCredentialBuilder()\n+            .addLast(new AzureCliCredentialBuilder().build())\n+            .addLast(new AzureDeveloperCliCredentialBuilder().build())\n+            .addLast(new IntelliJCredentialBuilder().build())\n+            .build();\n+\n+        return new CredentialSelection(\n+            credential,\n+            \"Azure CLI -\u003e Azure Developer CLI -\u003e IntelliJ\",\n+            enableCae\n+        );\n+    }\n+\n+    private CredentialSelection ciCredential(boolean enableCae) {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        String strategy;\n+\n+        if (hasAzurePipelinesFederationConfiguration()) {\n+            chain.addLast(new AzurePipelinesCredentialBuilder()\n+                .tenantId(environment.get(\"AZURE_TENANT_ID\"))\n+                .clientId(environment.get(\"AZURE_CLIENT_ID\"))\n+                .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))\n+                .systemAccessToken(environment.get(\"SYSTEM_ACCESSTOKEN\"))\n+                .build());\n+            strategy = \"Azure Pipelines workload federation -\u003e environment credential\";\n+        } else {\n+            strategy = \"Environment credential\";\n+        }\n+\n+        chain.addLast(new EnvironmentCredentialBuilder().build());\n+        return new CredentialSelection(chain.build(), strategy, enableCae);\n+    }\n+\n+    private CredentialSelection productionCredential(boolean enableCae) {\n+        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();\n+        String managedIdentityClientId = environment.get(MANAGED_IDENTITY_CLIENT_ID);\n+        String managedIdentityDescription = \"system-assigned managed identity\";\n+\n+        if (hasText(managedIdentityClientId)) {\n+            managedIdentity.clientId(managedIdentityClientId);\n+            managedIdentityDescription = \"user-assigned managed identity\";\n+        }\n+\n+        TokenCredential credential = new ChainedTokenCredentialBuilder()\n+            .addLast(managedIdentity.build())\n+            .addLast(new WorkloadIdentityCredentialBuilder().build())\n+            .build();\n+\n+        return new CredentialSelection(\n+            credential,\n+            managedIdentityDescription + \" -\u003e workload identity\",\n+            enableCae\n+        );\n+    }\n+\n+    private boolean hasAzurePipelinesFederationConfiguration() {\n+        return hasText(environment.get(\"AZURE_TENANT_ID\"))\n+            \u0026\u0026 hasText(environment.get(\"AZURE_CLIENT_ID\"))\n+            \u0026\u0026 hasText(environment.get(SERVICE_CONNECTION_ID))\n+            \u0026\u0026 hasText(environment.get(\"SYSTEM_ACCESSTOKEN\"));\n+    }\n+\n+    private static boolean hasText(String value) {\n+        return value != null \u0026\u0026 !value.isBlank();\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/AuthenticationFailureAnalyzer.java\n+package com.example.azureauth;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+\n+import java.util.Locale;\n+\n+public final class AuthenticationFailureAnalyzer {\n+    private AuthenticationFailureAnalyzer() {\n+    }\n+\n+    public static String explain(Throwable failure) {\n+        String details = collectMessages(failure);\n+        String normalized = details.toLowerCase(Locale.ROOT);\n+\n+        if (containsAny(normalized, \"aadsts7000222\", \"expired certificate\", \"certificate has expired\")) {\n+            return \"The client certificate or secret has expired. Rotate it and update the credential source.\";\n+        }\n+        if (containsAny(normalized, \"aadsts7000215\", \"invalid client secret\")) {\n+            return \"The client secret is invalid. Check that the secret value, not its identifier, was supplied.\";\n+        }\n+        if (containsAny(normalized, \"aadsts700027\", \"certificate was not found\", \"invalid certificate\")) {\n+            return \"The client certificate is invalid or is not registered on the application.\";\n+        }\n+        if (containsAny(normalized, \"aadsts90002\", \"tenant not found\")) {\n+            return \"The configured Microsoft Entra tenant does not exist or cannot be reached.\";\n+        }\n+        if (containsAny(normalized, \"aadsts700016\", \"application with identifier\")) {\n+            return \"The client application was not found in the configured tenant; verify client and tenant IDs.\";\n+        }\n+        if (containsAny(normalized, \"managed identity\", \"identity not found\", \"no identity has been assigned\")) {\n+            return \"No usable managed identity was found; assign one or verify its client ID and endpoint.\";\n+        }\n+        if (containsAny(normalized, \"federated\", \"subject claim\", \"token file\")) {\n+            return \"Workload identity federation is misconfigured; verify tenant, client, subject, and token file.\";\n+        }\n+        if (failure instanceof CredentialUnavailableException) {\n+            return \"No credential in the selected chain is configured or available.\";\n+        }\n+        if (failure instanceof ClientAuthenticationException) {\n+            return \"Microsoft Entra ID rejected the credential. SDK details: \" + firstLine(details);\n+        }\n+        return \"Token acquisition failed. SDK details: \" + firstLine(details);\n+    }\n+\n+    private static String collectMessages(Throwable failure) {\n+        StringBuilder messages = new StringBuilder();\n+        Throwable current = failure;\n+        while (current != null) {\n+            if (current.getMessage() != null \u0026\u0026 !current.getMessage().isBlank()) {\n+                if (!messages.isEmpty()) {\n+                    messages.append(\" | \");\n+                }\n+                messages.append(current.getMessage());\n+            }\n+            current = current.getCause();\n+        }\n+        return messages.isEmpty() ? failure.getClass().getSimpleName() : messages.toString();\n+    }\n+\n+    private static boolean containsAny(String value, String... fragments) {\n+        for (String fragment : fragments) {\n+            if (value.contains(fragment)) {\n+                return true;\n+            }\n+        }\n+        return false;\n+    }\n+\n+    private static String firstLine(String value) {\n+        int lineBreak = value.indexOf('\\n');\n+        return lineBreak \u003e= 0 ? value.substring(0, lineBreak) : value;\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/ConnectivityTestResult.java\n+package com.example.azureauth;\n+\n+import java.time.OffsetDateTime;\n+\n+public record ConnectivityTestResult(\n+    boolean successful,\n+    OffsetDateTime expiresAt,\n+    boolean caeRequested,\n+    String failureReason\n+) {\n+    public static ConnectivityTestResult success(OffsetDateTime expiresAt, boolean caeRequested) {\n+        return new ConnectivityTestResult(true, expiresAt, caeRequested, null);\n+    }\n+\n+    public static ConnectivityTestResult failure(boolean caeRequested, String reason) {\n+        return new ConnectivityTestResult(false, null, caeRequested, reason);\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java\n+package com.example.azureauth;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenRequestContext;\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+\n+import java.util.Objects;\n+\n+public final class SyncCredentialConnectivityTester {\n+    public ConnectivityTestResult test(CredentialSelection selection, String scope) {\n+        Objects.requireNonNull(selection, \"selection\");\n+        TokenRequestContext request = request(scope, selection.caeEnabled());\n+\n+        try {\n+            AccessToken token = selection.credential().getTokenSync(request);\n+            ConnectivityTestResult result =\n+                ConnectivityTestResult.success(token.getExpiresAt(), selection.caeEnabled());\n+            print(result);\n+            return result;\n+        } catch (CredentialUnavailableException | ClientAuthenticationException failure) {\n+            ConnectivityTestResult result = ConnectivityTestResult.failure(\n+                selection.caeEnabled(),\n+                AuthenticationFailureAnalyzer.explain(failure)\n+            );\n+            print(result);\n+            return result;\n+        }\n+    }\n+\n+    static TokenRequestContext request(String scope, boolean enableCae) {\n+        if (scope == null || scope.isBlank()) {\n+            throw new IllegalArgumentException(\"scope must not be blank\");\n+        }\n+        return new TokenRequestContext().addScopes(scope).setCaeEnabled(enableCae);\n+    }\n+\n+    static void print(ConnectivityTestResult result) {\n+        if (result.successful()) {\n+            System.out.printf(\n+                \"SUCCESS - token expires at %s; CAE requested: %s%n\",\n+                result.expiresAt(),\n+                result.caeRequested()\n+            );\n+        } else {\n+            System.out.printf(\n+                \"FAILURE - %s; CAE requested: %s%n\",\n+                result.failureReason(),\n+                result.caeRequested()\n+            );\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/AsyncCredentialConnectivityTester.java\n+package com.example.azureauth;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+public final class AsyncCredentialConnectivityTester {\n+    public Mono\u003cConnectivityTestResult\u003e test(CredentialSelection selection, String scope) {\n+        Objects.requireNonNull(selection, \"selection\");\n+\n+        return selection.credential()\n+            .getToken(SyncCredentialConnectivityTester.request(scope, selection.caeEnabled()))\n+            .map(token -\u003e ConnectivityTestResult.success(token.getExpiresAt(), selection.caeEnabled()))\n+            .onErrorResume(\n+                failure -\u003e failure instanceof CredentialUnavailableException\n+                    || failure instanceof ClientAuthenticationException,\n+                failure -\u003e Mono.just(ConnectivityTestResult.failure(\n+                    selection.caeEnabled(),\n+                    AuthenticationFailureAnalyzer.explain(failure)\n+                ))\n+            )\n+            .doOnNext(SyncCredentialConnectivityTester::print);\n+    }\n+}\n*** Add File: src/main/java/com/example/azureauth/Main.java\n+package com.example.azureauth;\n+\n+public final class Main {\n+    private static final String ARM_SCOPE = \"https://management.azure.com/.default\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        boolean enableCae = readCaeSetting();\n+        DeploymentEnvironment environment = new EnvironmentDetector().detect();\n+        CredentialSelection selection = new AzureCredentialFactory().create(environment, enableCae);\n+\n+        System.out.println(\"Detected environment: \" + environment);\n+        System.out.println(\"Credential strategy: \" + selection.strategy());\n+        System.out.println(\"Azure Resource Manager scope: \" + ARM_SCOPE);\n+\n+        System.out.println(\"\\nSynchronous connectivity test:\");\n+        new SyncCredentialConnectivityTester().test(selection, ARM_SCOPE);\n+\n+        System.out.println(\"\\nAsynchronous connectivity test:\");\n+        new AsyncCredentialConnectivityTester().test(selection, ARM_SCOPE).block();\n+    }\n+\n+    private static boolean readCaeSetting() {\n+        String configured = System.getenv().getOrDefault(\"AZURE_CAE_ENABLED\", \"true\");\n+        return Boolean.parseBoolean(configured);\n+    }\n+}\n*** Add File: src/test/java/com/example/azureauth/EnvironmentDetectorTest.java\n+package com.example.azureauth;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+class EnvironmentDetectorTest {\n+    @Test\n+    void defaultsToDevelopment() {\n+        assertEquals(DeploymentEnvironment.DEVELOPMENT, new EnvironmentDetector(Map.of()).detect());\n+    }\n+\n+    @Test\n+    void detectsCiBeforeAzureHostingSignals() {\n+        EnvironmentDetector detector = new EnvironmentDetector(Map.of(\n+            \"TF_BUILD\", \"True\",\n+            \"IDENTITY_ENDPOINT\", \"http://localhost/identity\"\n+        ));\n+\n+        assertEquals(DeploymentEnvironment.CI, detector.detect());\n+    }\n+\n+    @Test\n+    void detectsProductionFromManagedIdentityEndpoint() {\n+        EnvironmentDetector detector =\n+            new EnvironmentDetector(Map.of(\"IDENTITY_ENDPOINT\", \"http://localhost/identity\"));\n+\n+        assertEquals(DeploymentEnvironment.PRODUCTION, detector.detect());\n+    }\n+\n+    @Test\n+    void detectsProductionFromWorkloadIdentity() {\n+        EnvironmentDetector detector = new EnvironmentDetector(Map.of(\n+            \"AZURE_TENANT_ID\", \"tenant\",\n+            \"AZURE_CLIENT_ID\", \"client\",\n+            \"AZURE_FEDERATED_TOKEN_FILE\", \"token\"\n+        ));\n+\n+        assertEquals(DeploymentEnvironment.PRODUCTION, detector.detect());\n+    }\n+\n+    @Test\n+    void explicitOverrideWins() {\n+        EnvironmentDetector detector = new EnvironmentDetector(Map.of(\n+            \"APP_ENVIRONMENT\", \"dev\",\n+            \"TF_BUILD\", \"True\"\n+        ));\n+\n+        assertEquals(DeploymentEnvironment.DEVELOPMENT, detector.detect());\n+    }\n+\n+    @Test\n+    void rejectsUnknownOverride() {\n+        assertThrows(\n+            IllegalArgumentException.class,\n+            () -\u003e new EnvironmentDetector(Map.of(\"APP_ENVIRONMENT\", \"staging\")).detect()\n+        );\n+    }\n+}\n*** Add File: src/test/java/com/example/azureauth/ConnectivityTesterTest.java\n+package com.example.azureauth;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+import org.junit.jupiter.api.Test;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.OffsetDateTime;\n+import java.util.concurrent.atomic.AtomicReference;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class ConnectivityTesterTest {\n+    @Test\n+    void syncTesterRequestsScopeAndCae() {\n+        AtomicReference\u003cTokenRequestContext\u003e captured = new AtomicReference\u003c\u003e();\n+        OffsetDateTime expiry = OffsetDateTime.now().plusHours(1);\n+        CredentialSelection selection =\n+            new CredentialSelection(successfulCredential(captured, expiry), \"test\", true);\n+\n+        ConnectivityTestResult result =\n+            new SyncCredentialConnectivityTester().test(selection, \"https://management.azure.com/.default\");\n+\n+        assertTrue(result.successful());\n+        assertEquals(expiry, result.expiresAt());\n+        assertTrue(result.caeRequested());\n+        assertTrue(captured.get().isCaeEnabled());\n+        assertEquals(\n+            \"https://management.azure.com/.default\",\n+            captured.get().getScopes().get(0)\n+        );\n+    }\n+\n+    @Test\n+    void asyncTesterRequestsScopeAndCae() {\n+        AtomicReference\u003cTokenRequestContext\u003e captured = new AtomicReference\u003c\u003e();\n+        OffsetDateTime expiry = OffsetDateTime.now().plusHours(1);\n+        CredentialSelection selection =\n+            new CredentialSelection(successfulCredential(captured, expiry), \"test\", true);\n+\n+        ConnectivityTestResult result = new AsyncCredentialConnectivityTester()\n+            .test(selection, \"https://management.azure.com/.default\")\n+            .block();\n+\n+        assertTrue(result.successful());\n+        assertEquals(expiry, result.expiresAt());\n+        assertTrue(captured.get().isCaeEnabled());\n+    }\n+\n+    private static TokenCredential successfulCredential(\n+        AtomicReference\u003cTokenRequestContext\u003e captured,\n+        OffsetDateTime expiry\n+    ) {\n+        return request -\u003e {\n+            captured.set(request);\n+            return Mono.just(new AccessToken(\"fake-token\", expiry));\n+        };\n+    }\n+}\n*** Add File: src/test/java/com/example/azureauth/AuthenticationFailureAnalyzerTest.java\n+package com.example.azureauth;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+import org.junit.jupiter.api.Test;\n+\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class AuthenticationFailureAnalyzerTest {\n+    @Test\n+    void identifiesWrongTenantOrApplication() {\n+        ClientAuthenticationException failure =\n+            new ClientAuthenticationException(\"AADSTS700016: Application with identifier was not found\", null);\n+\n+        assertTrue(AuthenticationFailureAnalyzer.explain(failure).contains(\"client application\"));\n+    }\n+\n+    @Test\n+    void identifiesExpiredCredential() {\n+        ClientAuthenticationException failure =\n+            new ClientAuthenticationException(\"AADSTS7000222: The provided client secret has expired\", null);\n+\n+        assertTrue(AuthenticationFailureAnalyzer.explain(failure).contains(\"expired\"));\n+    }\n+\n+    @Test\n+    void identifiesUnavailableChain() {\n+        CredentialUnavailableException failure =\n+            new CredentialUnavailableException(\"No credential was able to authenticate\");\n+\n+        assertTrue(AuthenticationFailureAnalyzer.explain(failure).contains(\"No credential\"));\n+    }\n+}\n*** Add File: README.md\n+# Azure credential chains for Java\n+\n+Small Java 17 sample that selects a deliberately narrow Azure Identity credential chain for local\n+development, CI/CD, or production. It requests an Azure Resource Manager token synchronously and\n+asynchronously without creating or modifying Azure resources.\n+\n+## Credential strategies\n+\n+| Environment | Detection examples | Credential order |\n+|---|---|---|\n+| Development | No CI or Azure-hosting markers | Azure CLI, Azure Developer CLI, IntelliJ |\n+| CI/CD | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, `PIPELINE_WORKSPACE` | Azure Pipelines workload federation when configured, then `EnvironmentCredential` |\n+| Production | `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, `AZURE_FEDERATED_TOKEN_FILE` | Managed identity, then workload identity |\n+\n+Set `APP_ENVIRONMENT=dev`, `ci`, or `production` to override detection. For a user-assigned managed\n+identity, set `AZURE_MANAGED_IDENTITY_CLIENT_ID`; otherwise the production chain uses the\n+system-assigned identity.\n+\n+Generic CI uses the standard `EnvironmentCredential` variables such as `AZURE_TENANT_ID`,\n+`AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`. Azure Pipelines\n+workload federation uses:\n+\n+- `AZURE_TENANT_ID`\n+- `AZURE_CLIENT_ID`\n+- `AZURE_SERVICE_CONNECTION_ID`\n+- `SYSTEM_ACCESSTOKEN`\n+\n+CAE is requested by default through `TokenRequestContext.setCaeEnabled(true)`. Set\n+`AZURE_CAE_ENABLED=false` to disable the request. The output says **CAE requested** because\n+`AccessToken` does not expose a definitive CAE-capable flag; the target resource decides whether it\n+honors the request.\n+\n+## Build and run\n+\n+```text\n+mvn test\n+mvn exec:java\n+```\n+\n+The run command performs real token acquisition. Sign in with one of the configured developer\n+tools or provide the environment variables appropriate for the detected environment.\n+\n+References:\n+\n+- [Credential chains in Azure Identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)\n+- [AzurePipelinesCredentialBuilder](https://learn.microsoft.com/java/api/com.azure.identity.azurepipelinescredentialbuilder)\n+- [TokenRequestContext and CAE](https://learn.microsoft.com/java/api/com.azure.core.credential.tokenrequestcontext)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/pom.xml
@@ -1,0 +1,56 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>azure-credential-chains</artifactId>
+    <version>1.0.0</version>
+    <name>Azure credential chains</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.identity.version>1.18.2</azure.identity.version>
+        <junit.version>5.12.2</junit.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>org.junit.jupiter</groupId>
+            <artifactId>junit-jupiter</artifactId>
+            <version>${junit.version}</version>
+            <scope>test</scope>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-compiler-plugin</artifactId>
+                <version>3.14.0</version>
+            <
... (truncated)
```

### ✅ 14. powershell — 4838ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Compile project and run unit tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:azure-credential-chains >-----------------
[INFO] Building Azure credential chains 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.0:compile (default-compile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java:[21,51] Alternatives in a multi-catch statement cannot be related by subclassing
  Alternative com.azure.identity.CredentialUnavailableException is a subclass of alternative com.azure.core.exception.ClientAuthenticationException
[INFO] 1 error
[INFO] -------------------------------------------------------------
[INFO] -----------------------------------------------------------
... (truncated)
```

### ✅ 15. apply_patch — 770ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java\n@@\n-import com.azure.identity.CredentialUnavailableException;\n@@\n-        } catch (CredentialUnavailableException | ClientAuthenticationException failure) {\n+        } catch (ClientAuthenticationException failure) {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/SyncCredentialConnectivityTester.java
@@ -3,7 +3,6 @@
 import com.azure.core.credential.AccessToken;
 import com.azure.core.credential.TokenRequestContext;
 import com.azure.core.exception.ClientAuthenticationException;
-import com.azure.identity.CredentialUnavailableException;
 
 import java.util.Objects;
 
@@ -18,7 +17,7 @@
                 ConnectivityTestResult.success(token.getExpiresAt(), selection.caeEnabled());
             print(result);
             return result;
-        } catch (CredentialUnavailableException | ClientAuthenticationException failure) {
+        } catch (ClientAuthenticationException failure) {
             ConnectivityTestResult result = ConnectivityTestResult.failure(
                 selection.caeEnabled(),
                 AuthenticationFailureAnalyzer.explain(failure)

```

### ✅ 16. powershell — 6524ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Recompile project and run unit tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:azure-credential-chains >-----------------
[INFO] Building Azure credential chains 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.0:compile (default-compile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.0:testCompile (default-testCompile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 3 source files with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.3:test (default-test) @ azure-credential-chains ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider
[INFO] 
[INFO] ----------------------
... (truncated)
```

### ✅ 17. apply_patch — 758ms

**Input:**

```json
"*** Begin Patch\n*** Add File: src/test/java/com/example/azureauth/AzureCredentialFactoryTest.java\n+package com.example.azureauth;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertFalse;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class AzureCredentialFactoryTest {\n+    @Test\n+    void buildsDeveloperToolChain() {\n+        CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.DEVELOPMENT, true);\n+\n+        assertTrue(selection.strategy().contains(\"Azure CLI\"));\n+        assertTrue(selection.caeEnabled());\n+    }\n+\n+    @Test\n+    void buildsGenericCiChain() {\n+        CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.CI, false);\n+\n+        assertTrue(selection.strategy().contains(\"Environment credential\"));\n+        assertFalse(selection.caeEnabled());\n+    }\n+\n+    @Test\n+    void buildsAzurePipelinesFederatedChain() {\n+        CredentialSelection selection = new AzureCredentialFactory(Map.of(\n+            \"AZURE_TENANT_ID\", \"00000000-0000-0000-0000-000000000001\",\n+            \"AZURE_CLIENT_ID\", \"00000000-0000-0000-0000-000000000002\",\n+            AzureCredentialFactory.SERVICE_CONNECTION_ID,\n+            \"00000000-0000-0000-0000-000000000003\",\n+            \"SYSTEM_ACCESSTOKEN\", \"fake-system-access-token\"\n+        )).create(DeploymentEnvironment.CI, true);\n+\n+        assertTrue(selection.strategy().startsWith(\"Azure Pipelines\"));\n+    }\n+\n+    @Test\n+    void buildsSystemAssignedManagedIdentityChain() {\n+        CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.PRODUCTION, true);\n+\n+        assertTrue(selection.strategy().startsWith(\"system-assigned\"));\n+    }\n+\n+    @Test\n+    void buildsUserAssignedManagedIdentityChain() {\n+        CredentialSelection selection = new AzureCredentialFactory(Map.of(\n+            AzureCredentialFactory.MANAGED_IDENTITY_CLIENT_ID,\n+            \"00000000-0000-0000-0000-000000000004\"\n+        )).create(DeploymentEnvironment.PRODUCTION, true);\n+\n+        assertTrue(selection.strategy().startsWith(\"user-assigned\"));\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/test/java/com/example/azureauth/AzureCredentialFactoryTest.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/test/java/com/example/azureauth/AzureCredentialFactoryTest.java
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/test/java/com/example/azureauth/AzureCredentialFactoryTest.java
@@ -1,0 +1,60 @@
+package com.example.azureauth;
+
+import org.junit.jupiter.api.Test;
+
+import java.util.Map;
+
+import static org.junit.jupiter.api.Assertions.assertFalse;
+import static org.junit.jupiter.api.Assertions.assertTrue;
+
+class AzureCredentialFactoryTest {
+    @Test
+    void buildsDeveloperToolChain() {
+        CredentialSelection selection =
+            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.DEVELOPMENT, true);
+
+        assertTrue(selection.strategy().contains("Azure CLI"));
+        assertTrue(selection.caeEnabled());
+    }
+
+    @Test
+    void buildsGenericCiChain() {
+        CredentialSelection selection =
+            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.CI, false);
+
+        assertTrue(selection.strategy().contains("Environment credential"));
+        assertFalse(selection.caeEnabled());
+    }
+
+    @Test
+    void buildsAzurePipelinesFederatedChain() {
+        CredentialSelection selection = new AzureCredentialFactory(Map.of(
+            "AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000001",
+            "AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000002",
+            AzureCredentialFactory.SERVICE_CONNECTION_ID,
+            "00000000-0000-0000-0000-000000000003",
+            "S
... (truncated)
```

### ✅ 18. powershell — 6303ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Run complete unit test suite","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:azure-credential-chains >-----------------
[INFO] Building Azure credential chains 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.0:compile (default-compile) @ azure-credential-chains ---
[INFO] Nothing to compile - all classes are up to date.
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.0:testCompile (default-testCompile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 4 source files with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.3:test (default-test) @ azure-credential-chains ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider
[INFO] 
[INFO] -------------------------------------------------------
[INFO]  T E S T S
[INFO] --------------------------
... (truncated)
```

### ✅ 19. apply_patch — 800ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azureauth/AzureCredentialFactory.java\n@@\n                 .clientId(environment.get(\"AZURE_CLIENT_ID\"))\n                 .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))\n                 .systemAccessToken(environment.get(\"SYSTEM_ACCESSTOKEN\"))\n+                .oidcEndpoint(environment.get(\"SYSTEM_OIDCREQUESTURI\"))\n                 .build());\n@@\n-        TokenCredential credential = new ChainedTokenCredentialBuilder()\n-            .addLast(managedIdentity.build())\n-            .addLast(new WorkloadIdentityCredentialBuilder().build())\n-            .build();\n+        ChainedTokenCredentialBuilder chain =\n+            new ChainedTokenCredentialBuilder().addLast(managedIdentity.build());\n+        String strategy = managedIdentityDescription;\n+\n+        if (hasWorkloadIdentityConfiguration()) {\n+            chain.addLast(new WorkloadIdentityCredentialBuilder()\n+                .tenantId(environment.get(\"AZURE_TENANT_ID\"))\n+                .clientId(environment.get(\"AZURE_CLIENT_ID\"))\n+                .tokenFilePath(environment.get(\"AZURE_FEDERATED_TOKEN_FILE\"))\n+                .build());\n+            strategy += \" -\u003e workload identity\";\n+        } else {\n+            strategy += \" (workload identity fallback not configured)\";\n+        }\n \n         return new CredentialSelection(\n-            credential,\n-            managedIdentityDescription + \" -\u003e workload identity\",\n+            chain.build(),\n+            strategy,\n             enableCae\n         );\n@@\n             \u0026\u0026 hasText(environment.get(\"AZURE_CLIENT_ID\"))\n             \u0026\u0026 hasText(environment.get(SERVICE_CONNECTION_ID))\n-            \u0026\u0026 hasText(environment.get(\"SYSTEM_ACCESSTOKEN\"));\n+            \u0026\u0026 hasText(environment.get(\"SYSTEM_ACCESSTOKEN\"))\n+            \u0026\u0026 hasText(environment.get(\"SYSTEM_OIDCREQUESTURI\"));\n+    }\n+\n+    private boolean hasWorkloadIdentityConfiguration() {\n+        return hasText(environment.get(\"AZURE_TENANT_ID\"))\n+            \u0026\u0026 hasText(environment.get(\"AZURE_CLIENT_ID\"))\n+            \u0026\u0026 hasText(environment.get(\"AZURE_FEDERATED_TOKEN_FILE\"));\n     }\n*** Update File: src/test/java/com/example/azureauth/AzureCredentialFactoryTest.java\n@@\n             AzureCredentialFactory.SERVICE_CONNECTION_ID,\n             \"00000000-0000-0000-0000-000000000003\",\n-            \"SYSTEM_ACCESSTOKEN\", \"fake-system-access-token\"\n+            \"SYSTEM_ACCESSTOKEN\", \"fake-system-access-token\",\n+            \"SYSTEM_OIDCREQUESTURI\", \"https://example.invalid/oidc\"\n         )).create(DeploymentEnvironment.CI, true);\n@@\n         assertTrue(selection.strategy().startsWith(\"user-assigned\"));\n     }\n+\n+    @Test\n+    void addsWorkloadIdentityFallbackWhenConfigured() {\n+        CredentialSelection selection = new AzureCredentialFactory(Map.of(\n+            \"AZURE_TENANT_ID\", \"00000000-0000-0000-0000-000000000001\",\n+            \"AZURE_CLIENT_ID\", \"00000000-0000-0000-0000-000000000002\",\n+            \"AZURE_FEDERATED_TOKEN_FILE\", \"fake-token-path\"\n+        )).create(DeploymentEnvironment.PRODUCTION, true);\n+\n+        assertTrue(selection.strategy().endsWith(\"workload identity\"));\n+    }\n }\n*** Update File: README.md\n@@\n - `AZURE_SERVICE_CONNECTION_ID`\n - `SYSTEM_ACCESSTOKEN`\n+- `SYSTEM_OIDCREQUESTURI`\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java
@@ -60,6 +60,7 @@
                 .clientId(environment.get("AZURE_CLIENT_ID"))
                 .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))
                 .systemAccessToken(environment.get("SYSTEM_ACCESSTOKEN"))
+                .oidcEndpoint(environment.get("SYSTEM_OIDCREQUESTURI"))
                 .build());
             strategy = "Azure Pipelines workload federation -> environment credential";
         } else {
@@ -80,14 +81,24 @@
             managedIdentityDescription = "user-assigned managed identity";
         }
 
-        TokenCredential credential = new ChainedTokenCredentialBuilder()
-            .addLast(managedIdentity.build())
-            .addLast(new WorkloadIdentityCredentialBuilder().build())
-            .build();
+        ChainedTokenCredentialBuilder chain =
+            new ChainedTokenCredentialBuilder().addLast(managedIdentity.build());
+        String strategy = managedIdentityDescription;
+
+        if (hasWorkloadIdentityConfiguration()) {
+            chain.addLast(new WorkloadIdentityCredentialBuilder()
+                .tenantId(environment.get("AZURE_TENANT_ID"))
+       
... (truncated)
```

### ✅ 20. powershell — 4405ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Verify corrected credential chains","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:azure-credential-chains >-----------------
[INFO] Building Azure credential chains 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.0:compile (default-compile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java:[63,17] cannot find symbol
  symbol:   method oidcEndpoint(java.lang.String)
  location: class com.azure.identity.AzurePipelinesCredentialBuilder
[INFO] 1 error
[INFO] -------------------------------------------------------------
[INFO] ------------------------------------------------------------------------
[INFO] BUILD FAILURE
[INFO] ---------------------------------------------------
... (truncated)
```

### ✅ 21. powershell — 1721ms

**Input:**

```json
{"command":"$jar = Join-Path $env:USERPROFILE '.m2\\repository\\com\\azure\\azure-identity\\1.18.2\\azure-identity-1.18.2.jar'; javap -classpath $jar com.azure.identity.AzurePipelinesCredentialBuilder; javap -classpath $jar -c com.azure.identity.AzurePipelinesCredentialBuilder","description":"Inspect Azure Pipelines builder API"}
```

**Output:**

```
Compiled from "AzurePipelinesCredentialBuilder.java"
public class com.azure.identity.AzurePipelinesCredentialBuilder extends com.azure.identity.AadCredentialBuilderBase<com.azure.identity.AzurePipelinesCredentialBuilder> {
  public com.azure.identity.AzurePipelinesCredentialBuilder();
  public com.azure.identity.AzurePipelinesCredentialBuilder serviceConnectionId(java.lang.String);
  public com.azure.identity.AzurePipelinesCredentialBuilder systemAccessToken(java.lang.String);
  public com.azure.identity.AzurePipelinesCredentialBuilder tokenCachePersistenceOptions(com.azure.identity.TokenCachePersistenceOptions);
  public com.azure.identity.AzurePipelinesCredential build();
  static {};
}
Compiled from "AzurePipelinesCredentialBuilder.java"
public class com.azure.identity.AzurePipelinesCredentialBuilder extends com.azure.identity.AadCredentialBuilderBase<com.azure.identity.AzurePipelinesCredentialBuilder> {
  public com.azure.identity.AzurePipelinesCredentialBuilder();
    Code:
         0: aload_0
         1: invokespecial #1                  // Method com/azure/identity/AadCredentialBuilderBase."<init>":()V
         4: return

  public com.azure.identity.AzurePipelinesCredentialBuilder serviceConnectionId(java.lang.String);
    Code:
         0: aload_0
         1: aload_1
         2: putfield      #7                  // Field serviceConnectionId:Ljava/lang/String;
         5: aload_0
         6: areturn

  public com.azure.identity.AzurePipelinesCredentialBuilder systemAccessToken(java.lang.String);
    Code:
         0: aload_0
         1: aload_1
         2: putfield      #13                 // Field systemAccessToken:Ljava/lang/String;
         5: aload_0
         6: areturn

  public com.azure.identity.AzurePipelinesCredentialBuilder tokenCachePersistenceOptions(com.azure.identity.TokenCachePersistenceOptions);
    Code:
         0: aload_0
         1: getfield      #16                 // Field identityClientOptions:Lcom/azure/identity/implementation/Identity
... (truncated)
```

### ✅ 22. powershell — 1419ms

**Input:**

```json
{"command":"$jar = Join-Path $env:USERPROFILE '.m2\\repository\\com\\azure\\azure-identity\\1.18.2\\azure-identity-1.18.2.jar'; javap -classpath $jar com.azure.identity.AadCredentialBuilderBase","description":"Inspect shared Azure credential builder API"}
```

**Output:**

```
Compiled from "AadCredentialBuilderBase.java"
public abstract class com.azure.identity.AadCredentialBuilderBase<T extends com.azure.identity.AadCredentialBuilderBase<T>> extends com.azure.identity.CredentialBuilderBase<T> {
  java.lang.String clientId;
  java.lang.String tenantId;
  public com.azure.identity.AadCredentialBuilderBase();
  public T authorityHost(java.lang.String);
  public T clientId(java.lang.String);
  public T tenantId(java.lang.String);
  public T executorService(java.util.concurrent.ExecutorService);
  public T additionallyAllowedTenants(java.lang.String...);
  public T additionallyAllowedTenants(java.util.List<java.lang.String>);
  public T disableInstanceDiscovery();
  public T enableUnsafeSupportLogging();
  static {};
}
<shellId: 6 completed with exit code 0>
```

### ✅ 23. powershell — 1351ms

**Input:**

```json
{"command":"$jar = Join-Path $env:USERPROFILE '.m2\\repository\\com\\azure\\azure-identity\\1.18.2\\azure-identity-1.18.2.jar'; javap -classpath $jar com.azure.identity.CredentialBuilderBase","description":"Inspect credential builder configuration API"}
```

**Output:**

```
Compiled from "CredentialBuilderBase.java"
public abstract class com.azure.identity.CredentialBuilderBase<T extends com.azure.identity.CredentialBuilderBase<T>> implements com.azure.core.client.traits.HttpTrait<T> {
  com.azure.identity.implementation.IdentityClientOptions identityClientOptions;
  com.azure.identity.CredentialBuilderBase();
  public T maxRetry(int);
  public T retryTimeout(java.util.function.Function<java.time.Duration, java.time.Duration>);
  public T proxyOptions(com.azure.core.http.ProxyOptions);
  public T httpPipeline(com.azure.core.http.HttpPipeline);
  public T httpClient(com.azure.core.http.HttpClient);
  public T configuration(com.azure.core.util.Configuration);
  public T clientOptions(com.azure.core.util.ClientOptions);
  public T httpLogOptions(com.azure.core.http.policy.HttpLogOptions);
  public T retryPolicy(com.azure.core.http.policy.RetryPolicy);
  public T retryOptions(com.azure.core.http.policy.RetryOptions);
  public T addPolicy(com.azure.core.http.policy.HttpPipelinePolicy);
  public T pipeline(com.azure.core.http.HttpPipeline);
  public T enableAccountIdentifierLogging();
  public com.azure.core.client.traits.HttpTrait clientOptions(com.azure.core.util.ClientOptions);
  public com.azure.core.client.traits.HttpTrait httpLogOptions(com.azure.core.http.policy.HttpLogOptions);
  public com.azure.core.client.traits.HttpTrait retryOptions(com.azure.core.http.policy.RetryOptions);
  public com.azure.core.client.traits.HttpTrait addPolicy(com.azure.core.http.policy.HttpPipelinePolicy);
  public com.azure.core.client.traits.HttpTrait pipeline(com.azure.core.http.HttpPipeline);
  public com.azure.core.client.traits.HttpTrait httpClient(com.azure.core.http.HttpClient);
  static {};
}
<shellId: 7 completed with exit code 0>
```

### ✅ 24. powershell — 1607ms

**Input:**

```json
{"command":"$core = Get-ChildItem \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-core\\*\\azure-core-*.jar\" | Sort-Object FullName -Descending | Select-Object -First 1 -ExpandProperty FullName; javap -classpath $core com.azure.core.util.ConfigurationBuilder","description":"Inspect Azure configuration builder API"}
```

**Output:**

```
Compiled from "ConfigurationBuilder.java"
public final class com.azure.core.util.ConfigurationBuilder {
  public com.azure.core.util.ConfigurationBuilder();
  public com.azure.core.util.ConfigurationBuilder(com.azure.core.util.ConfigurationSource);
  public com.azure.core.util.ConfigurationBuilder(com.azure.core.util.ConfigurationSource, com.azure.core.util.ConfigurationSource, com.azure.core.util.ConfigurationSource);
  public com.azure.core.util.ConfigurationBuilder putProperty(java.lang.String, java.lang.String);
  public com.azure.core.util.ConfigurationBuilder root(java.lang.String);
  public com.azure.core.util.Configuration build();
  public com.azure.core.util.Configuration buildSection(java.lang.String);
  static com.azure.core.util.logging.ClientLogger access$200();
  static {};
}
<shellId: 8 completed with exit code 0>
```

### ✅ 25. apply_patch — 701ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azureauth/AzureCredentialFactory.java\n@@\n import com.azure.core.credential.TokenCredential;\n+import com.azure.core.util.ConfigurationBuilder;\n@@\n                 .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))\n                 .systemAccessToken(environment.get(\"SYSTEM_ACCESSTOKEN\"))\n-                .oidcEndpoint(environment.get(\"SYSTEM_OIDCREQUESTURI\"))\n+                .configuration(new ConfigurationBuilder()\n+                    .putProperty(\"SYSTEM_OIDCREQUESTURI\", environment.get(\"SYSTEM_OIDCREQUESTURI\"))\n+                    .build())\n                 .build());\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298/src/main/java/com/example/azureauth/AzureCredentialFactory.java
@@ -1,6 +1,7 @@
 package com.example.azureauth;
 
 import com.azure.core.credential.TokenCredential;
+import com.azure.core.util.ConfigurationBuilder;
 import com.azure.identity.AzureCliCredentialBuilder;
 import com.azure.identity.AzureDeveloperCliCredentialBuilder;
 import com.azure.identity.AzurePipelinesCredentialBuilder;
@@ -60,7 +61,9 @@
                 .clientId(environment.get("AZURE_CLIENT_ID"))
                 .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))
                 .systemAccessToken(environment.get("SYSTEM_ACCESSTOKEN"))
-                .oidcEndpoint(environment.get("SYSTEM_OIDCREQUESTURI"))
+                .configuration(new ConfigurationBuilder()
+                    .putProperty("SYSTEM_OIDCREQUESTURI", environment.get("SYSTEM_OIDCREQUESTURI"))
+                    .build())
                 .build());
             strategy = "Azure Pipelines workload federation -> environment credential";
         } else {

```

### ✅ 26. powershell — 6069ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Run final Maven test suite","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:azure-credential-chains >-----------------
[INFO] Building Azure credential chains 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.0:compile (default-compile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-microsoft-skill-2098168298\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.0:testCompile (default-testCompile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 4 source files with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.3:test (default-test) @ azure-credential-chains ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider
[INFO] 
[INFO] ----------------------
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\azureauth\AsyncCredentialConnectivityTester.java`
- `src\main\java\com\example\azureauth\AuthenticationFailureAnalyzer.java`
- `src\main\java\com\example\azureauth\AzureCredentialFactory.java`
- `src\main\java\com\example\azureauth\ConnectivityTestResult.java`
- `src\main\java\com\example\azureauth\CredentialSelection.java`
- `src\main\java\com\example\azureauth\DeploymentEnvironment.java`
- `src\main\java\com\example\azureauth\EnvironmentDetector.java`
- `src\main\java\com\example\azureauth\Main.java`
- `src\main\java\com\example\azureauth\SyncCredentialConnectivityTester.java`
- `src\test\java\com\example\azureauth\AuthenticationFailureAnalyzerTest.java`
- `src\test\java\com\example\azureauth\AzureCredentialFactoryTest.java`
- `src\test\java\com\example\azureauth\ConnectivityTesterTest.java`
- `src\test\java\com\example\azureauth\EnvironmentDetectorTest.java`

## Copilot Response

I’ll scaffold the Maven project from the Azure Identity patterns, then compile and run local-only tests that do not contact Azure.

The design will keep CAE explicit at token-request time (the Java SDK’s supported model), while returning the credential and strategy together so sync and async tests use identical settings.

Construction checks exposed two SDK-required inputs: Azure Pipelines federation needs `SYSTEM_OIDCREQUESTURI`, and workload identity can only be added when its token-file configuration is complete. I’m tightening both chains so partially configured environments fail over correctly instead of failing during factory creation.

Created the Java 17 Maven project with:

- Environment-specific credential chains for development, CI/CD, and production
- Azure Pipelines workload federation and generic environment credentials
- System/user-assigned managed identity with AKS workload identity fallback
- Configurable CAE token requests
- Sync and async Azure Resource Manager connectivity testers
- Detailed authentication failure classification
- Environment auto-detection with explicit override support
- Unit coverage for factory selection, detection, diagnostics, and token requests
- Complete usage and configuration guidance in `README.md`

Run with:

```text
mvn test
mvn exec:java
```

References: [Azure credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), [AzurePipelinesCredentialBuilder](https://learn.microsoft.com/java/api/com.azure.identity.azurepipelinescredentialbuilder), and [TokenRequestContext CAE](https://learn.microsoft.com/java/api/com.azure.core.credential.tokenrequestcontext).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK client is instantiated without appropriate resource management.

## Grader Results

- credential-chain.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (14/14)
      - Uses `ChainedTokenCredentialBuilder` to compose multiple credentials: Pass
      - Credentials added via `.addLast()` — order matters: Pass
      - **Dev chain**: includes `AzureCliCredential`; may include `IntelliJCredential`, `VisualStudioCodeCredential`, `AzurePowerShellCredential`: Pass
      - **CI chain**: uses `EnvironmentCredential` or `AzurePipelinesCredential` (not just `DefaultAzureCredential`): Pass
      - **Production chain**: `ManagedIdentityCredential` first (supports user-assigned via `clientId()`), `WorkloadIdentityCredential` as fallback: Pass
      - Enables CAE via `TokenRequestContext.setCaeEnabled(true)` or `enableCae()` on credential builders: Pass
      - Detects CI (checks `CI`, `TF_BUILD`, `AZURE_PIPELINE_WORKSPACE`, or similar): Pass
      - Detects production/managed identity (checks `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, or similar): Pass
      - Falls back to dev if neither detected: Pass
      - Creates `TokenRequestContext` with correct scope (`https://management.azure.com/.default`): Pass
      - Calls `getToken()` and prints token expiry from `AccessToken.getExpiresAt()`: Pass
      - Handles failure with specific exception info: Pass
      - Async tester uses reactive `getToken()` returning `Mono<AccessToken>`: Pass
      - NOT using `DefaultAzureCredential` as the CI credential (too broad): Pass
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
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.0000** | **84.6%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-credential-chain --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
