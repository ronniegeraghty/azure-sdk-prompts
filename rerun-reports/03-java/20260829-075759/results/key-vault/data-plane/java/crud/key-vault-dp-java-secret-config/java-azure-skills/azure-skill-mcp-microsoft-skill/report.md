# Evaluation Report: key-vault-dp-java-secret-config

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 386.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-secret-config` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 20/22 |
| Duration | 386.2s |
| Timestamp | 2026-08-29T02:33:40Z |
| Files Generated | 11 |
| Event Count | 14159 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 140.5s |
| Review | 244.5s |
| **Total** | **386.2s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-security-keyvault-secrets-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=437240 out=9118 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 3018ms

**Input:**

```json
{"skill":"azure-security-keyvault-secrets-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Key Vault Secrets (Java)

Securely store and manage secrets like passwords, API keys, and connection strings.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-secrets</artifactId>
    <version>4.9.0</version>
</dependency>
```

## Client Creation

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

// Sync client
SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();

// Async client
SecretAsyncClient secretAsyncClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildAsyncClient();
```

## Create/Set Secret

```java
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

// Simple secret
KeyVaultSecret secret = secretClient.setSecret("database-password", "P@ssw0rd123!");
System.out.println("Secret name: " + secret.getName());
System.out.println("Secret ID: " + secret.getId());

// Secret with options
KeyVaultSecret secretWithOptions = secretClient.setSecret(
    new KeyVaultSecre
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 7657ms

**Input:**

```json
{"command":"get_best_practices","intent":"Create a local-only Java 17 Maven application using Azure Key Vault Secrets sync and async clients with managed identity, caching, expiry checks, and safe soft-delete-aware rotation."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 8086ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-secrets-java/references/examples.md
@@ -1,396 +1,396 @@
 # Azure Key Vault Secrets SDK for Java - Examples
 
 Comprehensive code examples for the Azure Key Vault Secrets SDK for Java.
 
 ## Table of Contents
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Setting Secrets](#setting-secrets)
 - [Getting Secrets](#getting-secrets)
 - [Listing Secrets](#listing-secrets)
 - [Updating Secret Properties](#updating-secret-properties)
 - [Deleting and Recovering Secrets](#deleting-and-recovering-secrets)
 - [Purging Deleted Secrets](#purging-deleted-secrets)
 - [Backup and Restore](#backup-and-restore)
 - [Async Client Patterns](#async-client-patterns)
 - [Error Handling](#error-handling)
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-security-keyvault-secrets</artifactId>
     <version>4.11.0-beta.1</version>
 </dependency>
 
 <!-- Required for authentication -->
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.18.2</version>
 </dependency>
 ```
 
 ## Client Creation
 
 ### Sync SecretClient
 
 ```java
 import com.azure.identity.DefaultAzureCredentialBuilder;
 import com.azure.security.keyvault.secrets.SecretClient;
 import com.azure.security.keyvault.secrets.SecretClientBuilder;
 

... (truncated)
```

### ✅ 4. glob — 7268ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java","pattern":"**/*acceptance*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 6856ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. azure-get_azure_bestpractices (via azure) — 769ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code generation best practices for a Java 17 Maven app using managed identity and Azure Key Vault Secrets sync/async clients.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 7. web_fetch — 2878ms

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

### ✅ 8. web_fetch — 2462ms

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

### ✅ 9. web_fetch — 2105ms

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

### ✅ 10. apply_patch — 827ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-config\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eKey Vault configuration provider\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.17\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.keyvaultconfig.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/keyvaultconfig/SecretValue.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public record SecretValue(\n+        String name,\n+        String value,\n+        String version,\n+        OffsetDateTime expiresOn,\n+        boolean defaultValue) {\n+\n+    public SecretValue {\n+        Objects.requireNonNull(name, \"name\");\n+        Objects.requireNonNull(value, \"value\");\n+    }\n+\n+    public boolean expiresWithin(Duration warningWindow, OffsetDateTime now) {\n+        Objects.requireNonNull(warningWindow, \"warningWindow\");\n+        Objects.requireNonNull(now, \"now\");\n+        return expiresOn != null \u0026\u0026 !expiresOn.isAfter(now.plus(warningWindow));\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SyncSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+import java.util.Objects;\n+\n+public final class SyncSecretProvider {\n+    private static final System.Logger LOGGER = System.getLogger(SyncSecretProvider.class.getName());\n+\n+    private final SecretClient client;\n+\n+    public SyncSecretProvider(SecretClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public SecretValue getSecret(String name, String defaultValue) {\n+        return getSecretVersion(name, null, defaultValue);\n+    }\n+\n+    public SecretValue getSecretVersion(String name, String version, String defaultValue) {\n+        requireText(name, \"name\");\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+\n+        try {\n+            KeyVaultSecret secret = version == null\n+                    ? client.getSecret(name)\n+                    : client.getSecret(name, version);\n+            return fromKeyVault(secret);\n+        } catch (ResourceNotFoundException exception) {\n+            LOGGER.log(System.Logger.Level.WARNING,\n+                    \"Secret \\\"{0}\\\" was not found; using its configured default.\", name);\n+            return new SecretValue(name, defaultValue, version, null, true);\n+        }\n+    }\n+\n+    private static SecretValue fromKeyVault(KeyVaultSecret secret) {\n+        return new SecretValue(\n+                secret.getName(),\n+                secret.getValue(),\n+                secret.getProperties().getVersion(),\n+                secret.getProperties().getExpiresOn(),\n+                false);\n+    }\n+\n+    private static void requireText(String value, String field) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(field + \" must not be blank\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+public final class AsyncSecretProvider {\n+    private static final System.Logger LOGGER = System.getLogger(AsyncSecretProvider.class.getName());\n+\n+    private final SecretAsyncClient client;\n+\n+    public AsyncSecretProvider(SecretAsyncClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public Mono\u003cSecretValue\u003e getSecret(String name, String defaultValue) {\n+        return getSecretVersion(name, null, defaultValue);\n+    }\n+\n+    public Mono\u003cSecretValue\u003e getSecretVersion(String name, String version, String defaultValue) {\n+        requireText(name, \"name\");\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+\n+        Mono\u003cKeyVaultSecret\u003e request = version == null\n+                ? client.getSecret(name)\n+                : client.getSecret(name, version);\n+\n+        return request\n+                .map(AsyncSecretProvider::fromKeyVault)\n+                .onErrorResume(ResourceNotFoundException.class, exception -\u003e {\n+                    LOGGER.log(System.Logger.Level.WARNING,\n+                            \"Secret \\\"{0}\\\" was not found; using its configured default.\", name);\n+                    return Mono.just(new SecretValue(name, defaultValue, version, null, true));\n+                });\n+    }\n+\n+    private static SecretValue fromKeyVault(KeyVaultSecret secret) {\n+        return new SecretValue(\n+                secret.getName(),\n+                secret.getValue(),\n+                secret.getProperties().getVersion(),\n+                secret.getProperties().getExpiresOn(),\n+                false);\n+    }\n+\n+    private static void requireText(String value, String field) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(field + \" must not be blank\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SyncSecretCache.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Collection;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class SyncSecretCache {\n+    private final SyncSecretProvider provider;\n+    private final Map\u003cString, String\u003e defaults;\n+    private final Duration warningWindow;\n+    private final Clock clock;\n+    private final ConcurrentHashMap\u003cString, SecretValue\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public SyncSecretCache(\n+            SyncSecretProvider provider,\n+            Map\u003cString, String\u003e defaults,\n+            Duration warningWindow) {\n+        this(provider, defaults, warningWindow, Clock.systemUTC());\n+    }\n+\n+    SyncSecretCache(\n+            SyncSecretProvider provider,\n+            Map\u003cString, String\u003e defaults,\n+            Duration warningWindow,\n+            Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.defaults = Map.copyOf(Objects.requireNonNull(defaults, \"defaults\"));\n+        this.warningWindow = requireNonNegative(warningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public void loadRequired(Collection\u003cString\u003e names) {\n+        Objects.requireNonNull(names, \"names\").forEach(this::refresh);\n+    }\n+\n+    public String get(String name) {\n+        SecretValue cached = cache.computeIfAbsent(name, this::fetch);\n+        if (isNearExpiry(cached)) {\n+            cached = refresh(name);\n+        }\n+        return cached.value();\n+    }\n+\n+    public SecretValue refresh(String name) {\n+        SecretValue refreshed = fetch(name);\n+        cache.put(name, refreshed);\n+        return refreshed;\n+    }\n+\n+    public Map\u003cString, SecretValue\u003e refreshExpiring() {\n+        cache.forEach((name, secret) -\u003e {\n+            if (isNearExpiry(secret)) {\n+                refresh(name);\n+            }\n+        });\n+        return snapshot();\n+    }\n+\n+    public Map\u003cString, SecretValue\u003e snapshot() {\n+        return Map.copyOf(cache);\n+    }\n+\n+    public boolean isNearExpiry(SecretValue secret) {\n+        return secret.expiresWithin(warningWindow, OffsetDateTime.now(clock));\n+    }\n+\n+    private SecretValue fetch(String name) {\n+        String defaultValue = defaults.get(name);\n+        if (defaultValue == null) {\n+            throw new IllegalArgumentException(\"No default configured for secret: \" + name);\n+        }\n+        return provider.getSecret(name, defaultValue);\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"warningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncSecretCache.java\n+package com.example.keyvaultconfig;\n+\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Collection;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class AsyncSecretCache {\n+    private final AsyncSecretProvider provider;\n+    private final Map\u003cString, String\u003e defaults;\n+    private final Duration warningWindow;\n+    private final Clock clock;\n+    private final ConcurrentHashMap\u003cString, SecretValue\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public AsyncSecretCache(\n+            AsyncSecretProvider provider,\n+            Map\u003cString, String\u003e defaults,\n+            Duration warningWindow) {\n+        this(provider, defaults, warningWindow, Clock.systemUTC());\n+    }\n+\n+    AsyncSecretCache(\n+            AsyncSecretProvider provider,\n+            Map\u003cString, String\u003e defaults,\n+            Duration warningWindow,\n+            Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.defaults = Map.copyOf(Objects.requireNonNull(defaults, \"defaults\"));\n+        this.warningWindow = requireNonNegative(warningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public Mono\u003cVoid\u003e loadRequired(Collection\u003cString\u003e names) {\n+        return Flux.fromIterable(Objects.requireNonNull(names, \"names\"))\n+                .flatMap(this::refresh)\n+                .then();\n+    }\n+\n+    public Mono\u003cString\u003e get(String name) {\n+        return Mono.defer(() -\u003e {\n+            SecretValue cached = cache.get(name);\n+            if (cached == null || isNearExpiry(cached)) {\n+                return refresh(name).map(SecretValue::value);\n+            }\n+            return Mono.just(cached.value());\n+        });\n+    }\n+\n+    public Mono\u003cSecretValue\u003e refresh(String name) {\n+        String defaultValue = defaults.get(name);\n+        if (defaultValue == null) {\n+            return Mono.error(new IllegalArgumentException(\n+                    \"No default configured for secret: \" + name));\n+        }\n+        return provider.getSecret(name, defaultValue)\n+                .doOnNext(secret -\u003e cache.put(name, secret));\n+    }\n+\n+    public Mono\u003cMap\u003cString, SecretValue\u003e\u003e refreshExpiring() {\n+        return Flux.fromIterable(cache.entrySet())\n+                .filter(entry -\u003e isNearExpiry(entry.getValue()))\n+                .flatMap(entry -\u003e refresh(entry.getKey()))\n+                .then(Mono.fromSupplier(this::snapshot));\n+    }\n+\n+    public Map\u003cString, SecretValue\u003e snapshot() {\n+        return Map.copyOf(cache);\n+    }\n+\n+    public boolean isNearExpiry(SecretValue secret) {\n+        return secret.expiresWithin(warningWindow, OffsetDateTime.now(clock));\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"warningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/KeyVaultClientFactory.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class KeyVaultClientFactory {\n+    public static final String VAULT_URL_ENV = \"AZURE_KEYVAULT_URL\";\n+    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = \"AZURE_CLIENT_ID\";\n+\n+    private KeyVaultClientFactory() {\n+    }\n+\n+    public static SecretClient createSyncClient() {\n+        return createBuilder(System.getenv()).buildClient();\n+    }\n+\n+    public static SecretAsyncClient createAsyncClient() {\n+        return createBuilder(System.getenv()).buildAsyncClient();\n+    }\n+\n+    static SecretClientBuilder createBuilder(Map\u003cString, String\u003e environment) {\n+        Objects.requireNonNull(environment, \"environment\");\n+        String vaultUrl = requireEnvironmentVariable(environment, VAULT_URL_ENV);\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        String clientId = environment.get(MANAGED_IDENTITY_CLIENT_ID_ENV);\n+        if (clientId != null \u0026\u0026 !clientId.isBlank()) {\n+            credentialBuilder.clientId(clientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+\n+        return new SecretClientBuilder()\n+                .vaultUrl(vaultUrl)\n+                .credential(credential);\n+    }\n+\n+    private static String requireEnvironmentVariable(\n+            Map\u003cString, String\u003e environment,\n+            String name) {\n+        String value = environment.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\n+                    \"Required environment variable \" + name + \" is not set\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SyncSecretRotator.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import com.azure.security.keyvault.secrets.models.SecretProperties;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public final class SyncSecretRotator {\n+    private final SecretClient client;\n+    private final Duration pollInterval;\n+    private final Duration timeout;\n+\n+    public SyncSecretRotator(SecretClient client, Duration pollInterval, Duration timeout) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+        this.pollInterval = requirePositive(pollInterval, \"pollInterval\");\n+        this.timeout = requirePositive(timeout, \"timeout\");\n+    }\n+\n+    public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {\n+        Objects.requireNonNull(newValue, \"newValue\");\n+        Objects.requireNonNull(expiresOn, \"expiresOn\");\n+\n+        client.beginDeleteSecret(name).waitForCompletion();\n+        purgeWhenVisible(name);\n+        waitUntilPurged(name);\n+\n+        return client.setSecret(new KeyVaultSecret(name, newValue)\n+                .setProperties(new SecretProperties().setExpiresOn(expiresOn)));\n+    }\n+\n+    private void purgeWhenVisible(String name) {\n+        long deadline = System.nanoTime() + timeout.toNanos();\n+        while (true) {\n+            try {\n+                client.purgeDeletedSecret(name);\n+                return;\n+            } catch (ResourceNotFoundException exception) {\n+                if (System.nanoTime() \u003e= deadline) {\n+                    throw new IllegalStateException(\n+                            \"Deleted secret did not become purgeable before timeout: \" + name,\n+                            exception);\n+                }\n+                sleep();\n+            }\n+        }\n+    }\n+\n+    private void waitUntilPurged(String name) {\n+        long deadline = System.nanoTime() + timeout.toNanos();\n+        while (true) {\n+            try {\n+                client.getDeletedSecret(name);\n+                if (System.nanoTime() \u003e= deadline) {\n+                    throw new IllegalStateException(\n+                            \"Secret purge did not complete before timeout: \" + name);\n+                }\n+                sleep();\n+            } catch (ResourceNotFoundException exception) {\n+                return;\n+            }\n+        }\n+    }\n+\n+    private void sleep() {\n+        try {\n+            Thread.sleep(pollInterval.toMillis());\n+        } catch (InterruptedException exception) {\n+            Thread.currentThread().interrupt();\n+            throw new IllegalStateException(\"Interrupted while waiting for secret deletion\", exception);\n+        }\n+    }\n+\n+    private static Duration requirePositive(Duration duration, String name) {\n+        Objects.requireNonNull(duration, name);\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(name + \" must be positive\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import com.azure.security.keyvault.secrets.models.SecretProperties;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+import reactor.util.retry.Retry;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+import java.util.concurrent.TimeoutException;\n+\n+public final class AsyncSecretRotator {\n+    private final SecretAsyncClient client;\n+    private final Duration pollInterval;\n+    private final Duration timeout;\n+\n+    public AsyncSecretRotator(SecretAsyncClient client, Duration pollInterval, Duration timeout) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+        this.pollInterval = requirePositive(pollInterval, \"pollInterval\");\n+        this.timeout = requirePositive(timeout, \"timeout\");\n+    }\n+\n+    public Mono\u003cKeyVaultSecret\u003e rotate(\n+            String name,\n+            String newValue,\n+            OffsetDateTime expiresOn) {\n+        Objects.requireNonNull(newValue, \"newValue\");\n+        Objects.requireNonNull(expiresOn, \"expiresOn\");\n+\n+        return client.beginDeleteSecret(name)\n+                .waitForCompletion()\n+                .then(purgeWhenVisible(name))\n+                .then(waitUntilPurged(name))\n+                .then(client.setSecret(new KeyVaultSecret(name, newValue)\n+                        .setProperties(new SecretProperties().setExpiresOn(expiresOn))));\n+    }\n+\n+    private Mono\u003cVoid\u003e purgeWhenVisible(String name) {\n+        return client.purgeDeletedSecret(name)\n+                .retryWhen(Retry.fixedDelay(maxRetries(), pollInterval)\n+                        .filter(ResourceNotFoundException.class::isInstance)\n+                        .onRetryExhaustedThrow((spec, signal) -\u003e new IllegalStateException(\n+                                \"Deleted secret did not become purgeable before timeout: \" + name,\n+                                signal.failure())));\n+    }\n+\n+    private Mono\u003cVoid\u003e waitUntilPurged(String name) {\n+        return Flux.interval(Duration.ZERO, pollInterval)\n+                .concatMap(ignored -\u003e client.getDeletedSecret(name)\n+                        .map(secret -\u003e false)\n+                        .onErrorResume(ResourceNotFoundException.class, exception -\u003e Mono.just(true)))\n+                .filter(Boolean::booleanValue)\n+                .next()\n+                .timeout(timeout)\n+                .onErrorMap(TimeoutException.class, exception -\u003e new IllegalStateException(\n+                        \"Secret purge did not complete before timeout: \" + name, exception))\n+                .then();\n+    }\n+\n+    private long maxRetries() {\n+        return Math.max(1L, timeout.dividedBy(pollInterval));\n+    }\n+\n+    private static Duration requirePositive(Duration duration, String name) {\n+        Objects.requireNonNull(duration, name);\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(name + \" must be positive\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/Main.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class Main {\n+    private static final List\u003cString\u003e REQUIRED_KEYS =\n+            List.of(\"database-url\", \"service-api-key\", \"feature-mode\");\n+    private static final Map\u003cString, String\u003e DEFAULTS = Map.of(\n+            \"database-url\", \"jdbc:h2:mem:local\",\n+            \"service-api-key\", \"not-configured\",\n+            \"feature-mode\", \"safe\");\n+    private static final Duration WARNING_WINDOW = Duration.ofDays(7);\n+    private static final Duration ROTATION_POLL_INTERVAL = Duration.ofSeconds(2);\n+    private static final Duration ROTATION_TIMEOUT = Duration.ofMinutes(2);\n+    private static final String ROTATION_SECRET = \"rotating-demo-secret\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String rotatedValue = requireEnvironmentVariable(\"DEMO_ROTATED_SECRET_VALUE\");\n+\n+        runSyncDemo(rotatedValue);\n+        runAsyncDemo(rotatedValue + \"-async\");\n+    }\n+\n+    private static void runSyncDemo(String rotatedValue) {\n+        System.out.println(\"=== Synchronous implementation ===\");\n+        SecretClient client = KeyVaultClientFactory.createSyncClient();\n+        SyncSecretCache cache = new SyncSecretCache(\n+                new SyncSecretProvider(client), DEFAULTS, WARNING_WINDOW);\n+\n+        cache.loadRequired(REQUIRED_KEYS);\n+        REQUIRED_KEYS.forEach(name -\u003e\n+                System.out.printf(\"%s = %s%n\", name, cache.get(name)));\n+\n+        cache.refresh(\"service-api-key\");\n+        cache.refreshExpiring();\n+        printExpiryWarnings(cache.snapshot(), cache::isNearExpiry);\n+\n+        SyncSecretRotator rotator =\n+                new SyncSecretRotator(client, ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT);\n+        rotator.rotate(ROTATION_SECRET, rotatedValue, OffsetDateTime.now().plusDays(90));\n+        System.out.println(\"Synchronous rotation completed.\");\n+    }\n+\n+    private static void runAsyncDemo(String rotatedValue) {\n+        System.out.println(\"=== Asynchronous implementation ===\");\n+        SecretAsyncClient client = KeyVaultClientFactory.createAsyncClient();\n+        AsyncSecretCache cache = new AsyncSecretCache(\n+                new AsyncSecretProvider(client), DEFAULTS, WARNING_WINDOW);\n+\n+        cache.loadRequired(REQUIRED_KEYS)\n+                .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_KEYS)\n+                        .concatMap(name -\u003e cache.get(name)\n+                                .doOnNext(value -\u003e\n+                                        System.out.printf(\"%s = %s%n\", name, value))))\n+                .then(cache.refresh(\"service-api-key\"))\n+                .then(cache.refreshExpiring())\n+                .doOnNext(snapshot -\u003e\n+                        printExpiryWarnings(snapshot, cache::isNearExpiry))\n+                .then(new AsyncSecretRotator(\n+                        client, ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT)\n+                        .rotate(\n+                                ROTATION_SECRET,\n+                                rotatedValue,\n+                                OffsetDateTime.now().plusDays(90)))\n+                .doOnSuccess(secret -\u003e\n+                        System.out.println(\"Asynchronous rotation completed.\"))\n+                .block();\n+    }\n+\n+    private static void printExpiryWarnings(\n+            Map\u003cString, SecretValue\u003e secrets,\n+            java.util.function.Predicate\u003cSecretValue\u003e nearExpiry) {\n+        secrets.values().stream()\n+                .filter(nearExpiry)\n+                .forEach(secret -\u003e System.out.printf(\n+                        \"WARNING: %s expires at %s%n\",\n+                        secret.name(),\n+                        secret.expiresOn()));\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\n+                    \"Required environment variable \" + name + \" is not set\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault configuration provider\n+\n+A Java 17 Maven sample with synchronous and asynchronous Key Vault secret\n+providers, in-memory caches, expiry-aware refresh, and soft-delete-aware secret\n+rotation.\n+\n+## Authentication and permissions\n+\n+The client uses `ManagedIdentityCredential`; no application secret or certificate\n+is stored in code. Assign the workload's managed identity the least-privilege Key\n+Vault data-plane permissions needed by the demo:\n+\n+- secrets: get, set, delete, purge\n+\n+`purge` is required because a soft-deleted secret name cannot be recreated until\n+the deleted object has been purged. Rotation by delete/recreate is incompatible\n+with purge protection until the retention period ends; for purge-protected\n+production vaults, prefer creating a new secret version instead.\n+\n+Set these environment variables:\n+\n+```text\n+AZURE_KEYVAULT_URL=https://your-vault.vault.azure.net\n+AZURE_CLIENT_ID=\u003coptional user-assigned managed identity client ID\u003e\n+DEMO_ROTATED_SECRET_VALUE=\u003cnew demo value\u003e\n+```\n+\n+`AZURE_CLIENT_ID` is omitted for a system-assigned managed identity.\n+\n+## Build and run\n+\n+```text\n+mvn clean package\n+mvn exec:java\n+```\n+\n+The demo runs the synchronous flow first and then the asynchronous flow. Both\n+load required keys, read the cache, refresh one key, refresh near-expiry entries,\n+print expiry warnings, and rotate `rotating-demo-secret`.\n+\n+Missing required configuration secrets use the defaults declared in `Main`.\n+Authentication, authorization, throttling, and other service errors are not\n+treated as missing secrets and remain visible to the caller.\n+\n+SDK reference:\n+https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/pom.xml
@@ -1,0 +1,65 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>key-vault-config</artifactId>
+    <version>1.0.0</version>
+    <name>Key Vault configuration provider</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.sdk.bom.version>1.3.8</azure.sdk.bom.version>
+    </properties>
+
+    <dependencyManagement>
+        <dependencies>
+            <dependency>
+                <groupId>com.azure</groupId>
+                <artifactId>azure-sdk-bom</artifactId>
+                <version>${azure.sdk.bom.version}</version>
+                <type>pom</type>
+                <scope>import</scope>
+            </dependency>
+        </dependencies>
+    </dependencyManagement>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-security-keyvault-secrets</artifactId>
+        </dependency>
+        <dependency>
+            <groupId>or
... (truncated)
```

### ✅ 11. powershell — 4520ms

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
[INFO] --------------------< com.example:key-vault-config >--------------------
[INFO] Building Key Vault configuration provider 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ key-vault-config ---
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-config ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ key-vault-config ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java:[35,17] cannot find symbol
  symbol:   method waitForCompletion()
  location: class com.azure.core.util.polling.PollerFlux<com.azure.security.keyvault.secrets.models.DeletedSecret,java.lang.Void>
[INFO] 1 error
[INFO] -------------------------------------------------------------
[INFO] ----------------------------------
... (truncated)
```

### ✅ 12. apply_patch — 777ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java\n@@\n         Objects.requireNonNull(expiresOn, \"expiresOn\");\n \n         return client.beginDeleteSecret(name)\n-                .waitForCompletion()\n+                .last()\n                 .then(purgeWhenVisible(name))\n*** Update File: src/main/java/com/example/keyvaultconfig/Main.java\n@@\n         cache.loadRequired(REQUIRED_KEYS);\n         REQUIRED_KEYS.forEach(name -\u003e\n-                System.out.printf(\"%s = %s%n\", name, cache.get(name)));\n+                System.out.printf(\"%s loaded (%d characters)%n\", name, cache.get(name).length()));\n@@\n                         .concatMap(name -\u003e cache.get(name)\n                                 .doOnNext(value -\u003e\n-                                        System.out.printf(\"%s = %s%n\", name, value))))\n+                                        System.out.printf(\n+                                                \"%s loaded (%d characters)%n\",\n+                                                name,\n+                                                value.length()))))\n*** Update File: README.md\n@@\n-The demo runs the synchronous flow first and then the asynchronous flow. Both\n-load required keys, read the cache, refresh one key, refresh near-expiry entries,\n-print expiry warnings, and rotate `rotating-demo-secret`.\n+The demo runs the synchronous flow first and then the asynchronous flow. Both\n+load required keys, read the cache without logging secret values, refresh one\n+key, refresh near-expiry entries, print expiry warnings, and rotate\n+`rotating-demo-secret`.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/AsyncSecretRotator.java
@@ -32,7 +32,7 @@
         Objects.requireNonNull(expiresOn, "expiresOn");
 
         return client.beginDeleteSecret(name)
-                .waitForCompletion()
+                .last()
                 .then(purgeWhenVisible(name))
                 .then(waitUntilPurged(name))
                 .then(client.setSecret(new KeyVaultSecret(name, newValue)


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/Main.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/Main.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/Main.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507/src/main/java/com/example/keyvaultconfig/Ma
... (truncated)
```

### ✅ 13. powershell — 5217ms

**Input:**

```json
{"command":"mvn --no-transfer-progress clean package","description":"Rebuild the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------------< com.example:key-vault-config >--------------------
[INFO] Building Key Vault configuration provider 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ key-vault-config ---
[INFO] Deleting C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507\target
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ key-vault-config ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ key-vault-config ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ key-vault-config ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-microsoft-skill-3915169507\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ key-vault-config ---
[INFO] No sources to compile
[INFO] 
[INFO] --- surefire:3.2.5:test (default-test) @ key-vault-config ---
[INFO] No tests to run.
[INFO] 
[INFO] --- jar:3.
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\keyvaultconfig\AsyncSecretCache.java`
- `src\main\java\com\example\keyvaultconfig\AsyncSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\AsyncSecretRotator.java`
- `src\main\java\com\example\keyvaultconfig\KeyVaultClientFactory.java`
- `src\main\java\com\example\keyvaultconfig\Main.java`
- `src\main\java\com\example\keyvaultconfig\SecretValue.java`
- `src\main\java\com\example\keyvaultconfig\SyncSecretCache.java`
- `src\main\java\com\example\keyvaultconfig\SyncSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\SyncSecretRotator.java`

## Copilot Response

I’ll ground the implementation in the Key Vault Java SDK guidance, then scaffold and verify the Maven project locally without contacting Azure.

The rotation path must account for soft delete by waiting for deletion, purging the deleted object, and waiting until the name is reusable before creating the replacement; otherwise Key Vault returns a conflict. I’m using that behavior in both sync and async helpers.

The project is scaffolded. I’m compiling against the current stable Azure SDK BOM now; any SDK signature mismatches will be corrected before handoff.

Created the Java 17 Maven project with:

- Sync/async secret providers supporting defaults, versions, and expiry metadata
- Thread-safe in-memory caches with startup loading, refresh, and expiry-aware re-fetch
- Managed identity client factory using `AZURE_KEYVAULT_URL`
- Sync/async soft-delete-aware rotation helpers
- Full sync-then-async demo without logging secret values
- Azure SDK BOM `1.3.8` and Key Vault/Identity dependencies
- Build and runtime instructions in `README.md`

Run with:

```powershell
$env:AZURE_KEYVAULT_URL = "https://your-vault.vault.azure.net"
$env:DEMO_ROTATED_SECRET_VALUE = "replacement-value"
mvn exec:java
```

Rotation requires `get`, `set`, `delete`, and `purge` permissions. Delete/recreate cannot complete promptly when purge protection is enabled; in that environment, creating a new secret version is the appropriate rotation model.

Reference: https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Client lifecycle handling is appropriate for the Azure Key Vault client types used.

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.0000** | **84.6%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-java-secret-config --config java-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
