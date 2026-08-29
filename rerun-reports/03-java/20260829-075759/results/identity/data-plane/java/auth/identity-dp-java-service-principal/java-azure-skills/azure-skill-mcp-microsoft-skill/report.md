# Evaluation Report: identity-dp-java-service-principal

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 386.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-service-principal` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 17/17 |
| Duration | 386.9s |
| Timestamp | 2026-08-29T01:55:32Z |
| Files Generated | 0 |
| Event Count | 3855 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 145.5s |
| Review | 240.7s |
| **Total** | **386.9s** |

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
| Token Usage | in=435293 out=5136 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 3530ms

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

### ✅ 2. view — 11726ms

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

### ❌ 3. view — 8641ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 8636ms

**Input:**

```json
{"command":"get_bestpractices","intent":"Provide a Java code example using Azure Identity ClientSecretCredential with an Azure Blob Storage SDK client, secure secret configuration, and invalid credential exception handling.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 9988ms

**Input:**

```json
{"command":"search","intent":"Find official Microsoft documentation for Java ClientSecretCredentialBuilder, Azure Blob Storage Java client authentication, and ClientAuthenticationException handling.","parameters":{"query":"Java ClientSecretCredentialBuilder BlobServiceClient ClientAuthenticationException service principal client secret"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 5459ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Java Azure SDK authentication and Blob client code-generation best practices.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 6151ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Search official documentation for service principal client secret authentication in Java and Azure Blob SDK integration.","parameters":{"query":"Java Azure Identity ClientSecretCredentialBuilder BlobServiceClient ClientAuthenticationException service principal secret"}}
```

**Output:**

```
{"results":[{"title":"Troubleshoot Azure Identity authentication problems","content":"# Troubleshoot Azure Identity authentication problems\n## Handle Azure Identity exceptions\nAs noted in the [Exception handling in the Azure SDK for Java](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-overview#exception-handling-in-the-azure-sdk-for-java) section of [Troubleshooting overview](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-overview), the Azure SDK for Java can throw a comprehensive set of exceptions and error codes. For Azure Identity specifically, a few key exception types are important to understand.\n### ClientAuthenticationException\nAny service client method that makes a request to the service can raise exceptions from authentication errors. These exceptions can occur because the token is requested from the credential on the first call to the service and on any subsequent requests to the service that need to refresh the token.\nTo distinguish these failures from failures in the service client, Azure Identity classes raise \u0060ClientAuthenticationException\u0060 with details describing the source of the error in the exception message and possibly the error message. Depending on the application, these errors might be recoverable. The following code shows an example of catching \u0060ClientAuthenticationException\u0060:\n\u0060\u0060\u0060java\n// Create a secret client using the DefaultAzureCredential\nSecretClient client = new SecretClientBuilder()\n    .vaultUrl(\u0022https://myvault.vault.azure.net/\u0022)\n    .credential(new DefaultAzureCredentialBuilder().build())\n    .buildClient();\n\ntry {\n    KeyVaultSecret secret = client.getSecret(\u0022secret1\u0022);\n} catch (ClientAuthenticationException e) {\n    //Handle Exception\n    e.printStackTrace();\n}\n\u0060\u0060\u0060\n### CredentialUnavailableException\n\u0060CredentialUnavailableException\u0060 is a special exception type derived from \u0060ClientAuth
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 5797ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java code samples for ClientSecretCredentialBuilder and Azure Blob Storage client credential authentication.","parameters":{"language":"java","query":"ClientSecretCredentialBuilder BlobServiceClientBuilder credential Java"}}
```

**Output:**

```
{"results":[{"description":"description: Demonstrates how to authenticate a Java application to Azure services using a service principal by creating a ClientSecretCredential with tenant ID, client ID, and client secret, and then using this credential to build an Azure BlobServiceClient.\npackage: com.azure.identity, com.azure.storage.blob\nlanguage: java\n","codeSnippet":"import com.azure.identity.ClientSecretCredential;\r\nimport com.azure.identity.ClientSecretCredentialBuilder;\r\nimport com.azure.storage.blob.BlobServiceClient;\r\nimport com.azure.storage.blob.BlobServiceClientBuilder;\r\n\r\nString tenantId = System.getenv(\u0022AZURE_TENANT_ID\u0022);\r\nString clientId = System.getenv(\u0022AZURE_CLIENT_ID\u0022);\r\nString clientSecret = System.getenv(\u0022AZURE_CLIENT_SECRET\u0022);\r\n\r\nClientSecretCredential credential = new ClientSecretCredentialBuilder()\r\n    .tenantId(tenantId)\r\n    .clientId(clientId)\r\n    .clientSecret(clientSecret)\r\n    .build();\r\n\r\nBlobServiceClient blobServiceClient = new BlobServiceClientBuilder()\r\n    .endpoint(\u0022https://\u003Caccount-name\u003E.blob.core.windows.net\u0022)\r\n    .credential(credential)\r\n    .buildClient();","link":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-service-principal#authenticate-to-azure-services-from-your-app","language":"java"},{"description":"description: Imports necessary classes and demonstrates how to authenticate to Azure Blob Storage using a service principal with ClientSecretCredential from the Azure Identity library, then creates a BlobServiceClient instance with the authenticated credential.\npackage: com.azure.identity, com.azure.storage.blob\nlanguage: java\n","codeSnippet":"import com.azure.identity.ClientSecretCredential;\r\nimport com.azure.identity.ClientSecretCredentialBuilder;\r\nimport com.azure.storage.blob.BlobServiceClient;\r\nimport com.azure.storage.blob.BlobServiceClientBuilder;\r\n\r\nString tenantId = System.g
... (truncated)
```

### ✅ 9. rg — 1517ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt","pattern":"\"(title|url|description)\""}
```

**Output:**

```
Output too large to read at once (41.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787968620187-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt:{"results":[{"title":"Troubleshoot Azure Identity authentication problems","content":"# Troubleshoot Azure Identity authentication problems\n## Handle Azure Identity exceptions\nAs noted in the [Exception handling in the Azure SDK for Java](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-overview#exception-handling-in-the-azure-sdk-for-java) section of [Troubleshooting o
```

### ✅ 10. rg — 1171ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt","pattern":"\"(title|url|description)\""}
```

**Output:**

```
Output too large to read at once (24.2 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787968620188-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt:{"results":[{"description":"description: Demonstrates how to authenticate a Java application to Azure services using a service principal by creating a ClientSecretCredential with tenant ID, client ID, and client secret, and then using this credential to build an Azure BlobServiceClient.\npackage: com.azure.identity, com.azure.storage.blob\nlanguage: java\n","codeSnippet":"import com.azure.id
```

### ✅ 11. powershell — 2216ms

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt' | ConvertFrom-Json; $j.results | Select-Object -First 6 title,contentUrl | Format-List","description":"Extract Microsoft documentation titles and URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-4c76b3aaea4445ef8f8f5236068ff888.txt' | ConvertFrom-Json; $j.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 12. powershell — 1837ms

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt' | ConvertFrom-Json; $j.results | Select-Object -First 6 description,url | Format-List","description":"Extract Java sample descriptions and URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-5b538ad0f2e540c6a66f5d7606ed90eb.txt' | ConvertFrom-Json; $j.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 13. view — 1437ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt b/C:/Users/weidxu/AppData/Local/Temp/1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt
@@ -1,6 +1,6 @@
 {"results":[{"description":"description: Demonstrates how to authenticate a Java application to Azure services using a service principal by creating a ClientSecretCredential with tenant ID, client ID, and client secret, and then using this credential to build an Azure BlobServiceClient.\npackage: com.azure.identity, com.azure.storage.blob\nlanguage: java\n","codeSnippet":"import com.azure.identity.ClientSecretCredential;\r\nimport com.azure.identity.ClientSecretCredentialBuilder;\r\nimport com.azure.storage.blob.BlobServiceClient;\r\nimport com.azure.storage.blob.BlobServiceClientBuilder;\r\n\r\nString tenantId = System.getenv(\u0022AZURE_TENANT_ID\u0022);\r\nString clientId = System.getenv(\u0022AZURE_CLIENT_ID\u0022);\r\nString clientSecret = System.getenv(\u0022AZURE_CLIENT_SECRET\u0022);\r\n\r\nClientSecretCredential credential = new ClientSecretCredentialBuilder()\r\n    .tenantId(tenantId)\r\n    .clientId(clientId)\r\n    .clientSecret(clientSecret)\r\n    .build();\r\n\r\nBlobServiceClient blobServiceClient = new BlobServiceClientBuilder()\r\n    .endpoint(\u0022https://\u003Caccount-name\u003E.blob.core.windows.net\u0022)\r\n    .credential(credential)\r\n    .buildClient();","link":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-service-principal#authenticate-to-azure-services-from-your-app","language":"java"},{"description":"description: Imports necessary classes and demonstrates how to authenticate to Azure Blob Storage using 
... (truncated)
```

### ✅ 14. view — 1090ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787968614056-copilot-tool-output-5b538ad0f2e540c6a66f5d7606ed90eb.txt","view_range":[1,6]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt b/C:/Users/weidxu/AppData/Local/Temp/1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787968614481-copilot-tool-output-4c76b3aaea4445ef8f8f5236068ff888.txt
@@ -1,6 +1,6 @@
 {"results":[{"title":"Troubleshoot Azure Identity authentication problems","content":"# Troubleshoot Azure Identity authentication problems\n## Handle Azure Identity exceptions\nAs noted in the [Exception handling in the Azure SDK for Java](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-overview#exception-handling-in-the-azure-sdk-for-java) section of [Troubleshooting overview](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-overview), the Azure SDK for Java can throw a comprehensive set of exceptions and error codes. For Azure Identity specifically, a few key exception types are important to understand.\n### ClientAuthenticationException\nAny service client method that makes a request to the service can raise exceptions from authentication errors. These exceptions can occur because the token is requested from the credential on the first call to the service and on any subsequent requests to the service that need to refresh the token.\nTo distinguish these failures from failures in the service client, Azure Identity classes raise \u0060ClientAuthenticationException\u0060 with details describing the source of the error in the exception message and possibly the error message. Depending on the application, these errors might be recoverable. The following code shows an example of catching \u0060ClientAuthenticationException\u0060:\n\u0060\u0060\u0060java\n// Create a secret client using the DefaultAzureCredential\nSecretClient client = new
... (truncated)
```

### ✅ 15. web_fetch — 2722ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-storage-blob</artifactId>
  <versioning>
    <latest>12.36.0-beta.1</latest>
    <release>12.36.0-beta.1</release>
    <versions>
      <version>12.7.0</version>
      <version>12.8.0-beta.1</version>
      <version>12.8.0</version>
      <version>12.9.0-beta.1</version>
      <version>12.9.0-beta.2</version>
      <version>12.9.0</version>
      <version>12.10.0-beta.1</version>
      <version>12.10.0</version>
      <version>12.10.1</version>
      <version>12.10.2</version>
      <version>12.11.0-beta.1</version>
      <version>12.11.0-beta.2</version>
      <version>12.11.0-beta.3</version>
      <version>12.11.0</version>
      <version>12.11.1</version>
      <version>12.12.0-beta.1</version>
      <version>12.12.0</version>
      <version>12.13.0</version>
      <version>12.14.0-beta.1</version>
      <version>12.14.0</version>
      <version>12.14.1</version>
      <version>12.14.2</version>
      <version>12.14.3</version>
      <version>12.14.4</version>
      <version>12.15.0-beta.1</version>
      <version>12.15.0-beta.2</version>
      <version>12.15.0-beta.3</version>
      <version>12.15.0</version>
      <version>12.16.0</version>
      <version>12.16.1</version>
      <version>12.17.0-beta.1</version>
      <version>12.17.0</version>
      <version>12.17.1</version>
      <version>12.18.0-beta.1</version>
      <version>12.18.0</version>
      <version>12.19.0</version>
      <version>12.19.1</version>
      <version>12.20.0-beta.1</version>
      <version>12.20.0</version>
      <version>12.20.1</version>
      <version>12.20.2</version>
      <version>12.20.3</version>
      <version>12.21.0-beta.1</version>
      <version>12.21.0</version>
      <version>12.21.1</version>
     
... (truncated)
```

### ✅ 16. web_fetch — 2380ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 17. web_fetch — 2043ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
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

### ✅ 18. azure-documentation (via azure) — 2418ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete official Java service principal authentication guidance for the final example.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-service-principal"}}
```

**Output:**

```
# Authenticate Java apps to Azure services during local development by using service principals

During local development, applications need to authenticate to Azure to access various Azure services. You can authenticate locally by using one of the following approaches:

- Use a developer account with one of the developer tools supported by the Azure Identity library. For more information, see [Authenticate Java apps to Azure services during local development by using developer accounts](local-development-dev-accounts).
- Use a service principal.

This article explains how to use an application service principal. For more information about service principals, see [Application and service principal objects in Microsoft Entra ID](/en-us/entra/identity-platform/app-objects-and-service-principals). In this article, you learn:

- How to register an application with Microsoft Entra to create a service principal.
- How to use Microsoft Entra groups to efficiently manage permissions.
- How to assign roles to scope permissions.
- How to authenticate by using a service principal from your app code.

Using dedicated application service principals enables you to follow the principle of least privilege when accessing Azure resources. You can limit permissions to the specific requirements of the app during development to prevent accidental access to Azure resources intended for other apps or services. This approach also helps you avoid problems when you move the app to production by ensuring it isn't over-privileged in the development environment.

![A diagram that shows how a local Java app uses a service principal to connect to Azure resources.](../../../includes/authentication/media/mermaidjs/local-service-principal-authentication.svg)

When you register the app in Azure, an application service principal is created. For local development, you should:

- Create a separate app registration for each developer working on the app so each developer has their own 
... (truncated)
```

### ✅ 19. azure-documentation (via azure) — 2094ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete official troubleshooting guidance for invalid Java ClientSecretCredential credentials.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-service-principal"}}
```

**Output:**

```
# Troubleshoot service principal authentication

This article provides guidance on dealing with issues encountered when authenticating Azure SDK for Java applications via service principal, through various `TokenCredential` implementations. For more information, see [Azure authentication with service principal](authentication/service-principal).

## Troubleshoot ClientSecretCredential

When you use `ClientSecretCredential`, you can optionally try/catch for `ClientAuthenticationException`. The following table shows the errors that this exception indicates, and methods of mitigation:

| Error code | Issue | Mitigation |
| --- | --- | --- |
| `AADSTS7000215` | An invalid client secret was provided. | Ensure that the `clientSecret` provided when constructing the credential is valid. If unsure, create a new client secret using the Azure portal. For more information, see the [Create a new application secret](/en-us/azure/active-directory/develop/howto-create-service-principal-portal#option-3-create-a-new-application-secret) section of [Create a Microsoft Entra application and service principal that can access resources](/en-us/azure/active-directory/develop/howto-create-service-principal-portal). |
| `AADSTS7000222` | An expired client secret was provided. | Create a new client secret using the Azure portal. For more information, see the [Create a new application secret](/en-us/azure/active-directory/develop/howto-create-service-principal-portal#option-3-create-a-new-application-secret) section of [Create a Microsoft Entra application and service principal that can access resources](/en-us/azure/active-directory/develop/howto-create-service-principal-portal). |
| `AADSTS700016` | The specified application wasn't found in the specified tenant. | Ensure the specified `clientId` and `tenantId` are correct for your application registration. For multi-tenant apps, ensure that a tenant admin has added the application to the desired tenant. For more information, see 
... (truncated)
```

## Copilot Response

I’ll ground the example in the current Azure Identity Java patterns and include a concrete Blob Storage client plus authentication-specific error handling.

Use environment variables for the service-principal values, construct a `ClientSecretCredential`, and pass it to the Azure SDK client builder. Authentication is lazy—the first service operation, not `build()`, validates the credentials.

### Maven dependencies

Use the Azure SDK BOM to keep Azure library versions compatible. Version `1.3.8` is the current stable BOM as of August 2026.

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>1.3.8</version>
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

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-storage-blob</artifactId>
    </dependency>
</dependencies>
```

`azure-core`, including `ClientAuthenticationException`, is included transitively.

### Configuration

For PowerShell:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<storage-account-name>"
```

The service principal needs an appropriate data-plane role, such as **Storage Blob Data Reader**, scoped as narrowly as possible.

### Complete Java example

```java
package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.ClientSecretCredential;
import com.azure.identity.ClientSecretCredentialBuilder;
import com.azure.identity.CredentialUnavailableException;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobContainerItem;
import com.azure.storage.blob.models.BlobStorageException;

import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.regex.Pattern;

public final class ServicePrincipalExample {
    private static final Logger LOGGER =
        Logger.getLogger(ServicePrincipalExample.class.getName());

    private static final Pattern STORAGE_ACCOUNT_NAME =
        Pattern.compile("[a-z0-9]{3,24}");

    private ServicePrincipalExample() {
    }

    public static void main(String[] args) {
        try {
            String tenantId = requireEnvironmentVariable("AZURE_TENANT_ID");
            String clientId = requireEnvironmentVariable("AZURE_CLIENT_ID");
            String clientSecret = requireEnvironmentVariable("AZURE_CLIENT_SECRET");
            String accountName =
                requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

            if (!STORAGE_ACCOUNT_NAME.matcher(accountName).matches()) {
                throw new IllegalArgumentException(
                    "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase "
                        + "letters or numbers.");
            }

            ClientSecretCredential credential =
                new ClientSecretCredentialBuilder()
                    .tenantId(tenantId)
                    .clientId(clientId)
                    .clientSecret(clientSecret)
                    .build();

            BlobServiceClient blobServiceClient =
                new BlobServiceClientBuilder()
                    .endpoint(
                        "https://" + accountName + ".blob.core.windows.net")
                    .credential(credential)
                    .buildClient();

            // Client construction is lazy. Iterating this result requests a token
            // and sends an authenticated request to Blob Storage.
            LOGGER.info("Authentication succeeded. Blob containers:");

            int count = 0;
            for (BlobContainerItem container
                    : blobServiceClient.listBlobContainers()) {
                LOGGER.info(() -> " - " + container.getName());
                count++;
            }

            if (count == 0) {
                LOGGER.info("No blob containers were found.");
            }
        } catch (CredentialUnavailableException exception) {
            LOGGER.log(
                Level.SEVERE,
                "The credential could not run in this environment. "
                    + "Verify the required configuration is available.",
                exception);
            System.exit(2);
        } catch (ClientAuthenticationException exception) {
            reportAuthenticationFailure(exception);
            System.exit(3);
        } catch (BlobStorageException exception) {
            reportStorageFailure(exception);
            System.exit(4);
        } catch (IllegalArgumentException exception) {
            LOGGER.log(
                Level.SEVERE,
                "Invalid application configuration: {0}",
                exception.getMessage());
            System.exit(2);
        } catch (RuntimeException exception) {
            LOGGER.log(
                Level.SEVERE,
                "Unexpected failure while accessing Azure Blob Storage.",
                exception);
            System.exit(1);
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Required environment variable " + name + " is not set.");
        }

        return value;
    }

    private static void reportAuthenticationFailure(
            ClientAuthenticationException exception) {
        String message =
            exception.getMessage() == null ? "" : exception.getMessage();

        if (message.contains("AADSTS7000215")) {
            LOGGER.severe(
                "Authentication failed: the client secret is invalid. "
                    + "Verify that you supplied the secret value, not its ID.");
        } else if (message.contains("AADSTS7000222")) {
            LOGGER.severe(
                "Authentication failed: the client secret has expired. "
                    + "Rotate it and update the application's secret store.");
        } else if (message.contains("AADSTS700016")) {
            LOGGER.severe(
                "Authentication failed: the application was not found in "
                    + "the specified tenant. Verify AZURE_CLIENT_ID and "
                    + "AZURE_TENANT_ID.");
        } else {
            LOGGER.severe(
                "Microsoft Entra authentication failed. Verify the tenant ID, "
                    + "client ID, client secret, network connectivity, and "
                    + "service-principal status.");
        }

        // Avoid printing credentials or detailed identity data to normal logs.
        LOGGER.log(Level.FINE, "Authentication failure details", exception);
    }

    private static void reportStorageFailure(BlobStorageException exception) {
        int statusCode = exception.getStatusCode();

        if (statusCode == 401) {
            LOGGER.severe(
                "Blob Storage rejected the access token. Verify the identity "
                    + "and storage endpoint.");
        } else if (statusCode == 403) {
            LOGGER.severe(
                "Authentication succeeded, but the service principal is not "
                    + "authorized. Check its Azure Storage data-plane RBAC role "
                    + "and assignment scope.");
        } else {
            LOGGER.log(
                Level.SEVERE,
                "Blob Storage request failed with HTTP status {0}.",
                statusCode);
        }

        LOGGER.log(Level.FINE, "Storage failure details", exception);
    }
}
```

### Secret-management practices

- Never hardcode or commit the client secret, including in `.properties`, `.env`, source code, or container images.
- Inject secrets at runtime from a CI/CD secret store or platform configuration.
- For Azure-hosted applications, prefer **managed identity** or workload identity so no client secret is required.
- If a secret is unavoidable, store it in Azure Key Vault, rotate it before expiration, use a short lifetime, and allow overlapping secrets during rotation.
- Assign only the required RBAC roles at the narrowest resource scope.
- Never log the secret. Restrict verbose Azure Identity logs to protected diagnostic environments.
- Prefer certificate credentials or workload-identity federation over long-lived client secrets where supported.

**References:** [Service principal authentication for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-service-principal#authenticate-to-azure-services-from-your-app) · [Troubleshoot `ClientSecretCredential`](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-service-principal) · [Azure Identity Java reference](https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable)

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The example does not leave any AutoCloseable Azure SDK client unclosed.

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Maven dependency for `azure-identity`: Pass
      - `ClientSecretCredentialBuilder` with tenantId, clientId, clientSecret: Pass
      - Passing credential to Azure SDK client builders: Pass
      - Environment variable patterns for secret storage: Pass
      - `AuthenticationException` handling: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 13.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-service-principal --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
