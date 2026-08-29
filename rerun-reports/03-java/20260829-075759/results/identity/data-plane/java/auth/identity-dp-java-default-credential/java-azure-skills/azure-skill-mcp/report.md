# Evaluation Report: identity-dp-java-default-credential

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 351.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-default-credential` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 15/17 |
| Duration | 351.7s |
| Timestamp | 2026-08-29T01:14:08Z |
| Files Generated | 0 |
| Event Count | 4492 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 131.6s |
| Review | 218.8s |
| **Total** | **351.7s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=330954 out=4091 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. azure-get_azure_bestpractices (via azure) — 2318ms

**Input:**

```json
{"intent":"Provide Java code and guidance for authenticating an Azure Key Vault SecretClient with DefaultAzureCredential, including dependencies, credential chain, environment differences, and troubleshooting.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 4675ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Java Azure SDK authentication code-generation best practices for DefaultAzureCredential and Key Vault SecretClient.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 5711ms

**Input:**

```json
{"intent":"Find current official documentation for Azure Identity DefaultAzureCredential in Java, its credential chain order, local versus Azure behavior, troubleshooting, and Key Vault SecretClient Maven dependencies and examples.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. azure-documentation (via azure) — 3881ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official documentation on Java DefaultAzureCredential chain order, continuation policy, local development, Azure managed identity, and troubleshooting.","parameters":{"query":"Java DefaultAzureCredential credential chain order EnvironmentCredential WorkloadIdentityCredential ManagedIdentityCredential Azure CLI IntelliJ Azure Developer CLI troubleshooting continuation policy"}}
```

**Output:**

```
{"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## Debug a chained credential\nTo diagnose an unexpected issue or to understand what a chained credential is doing, [enable logging](https://learn.microsoft.com/azure/developer/java/sdk/logging-overview) in your app.\nFor illustration purposes, assume the parameterless form of \u0060DefaultAzureCredential\u0060 is used to authenticate a request to a Blob Storage account. The app runs in the local development environment, and the developer authenticated to Azure using the Azure CLI. When the app is run, the following pertinent entries appear in the output:\nIn the preceding output, notice that:\n1. \u0060EnvironmentCredential\u0060, \u0060WorkloadIdentityCredential\u0060, \u0060ManagedIdentityCredential\u0060, \u0060IntelliJCredential\u0060, and \u0060VisualStudioCodeCredential\u0060 each failed to acquire a Microsoft Entra access token, in that order.\n2. The \u0060AzureCliCredential.getToken\u0060 call succeeds, as indicated by the \u0060returns a token\u0060-suffixed entry. Since \u0060AzureCliCredential\u0060 succeeded, no credentials beyond it were tried.","contentUrl":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains#debug-a-chained-credential"},{"title":"Azure Identity client library for Java - version 1.18.5","content":"# Azure Identity client library for Java - version 1.18.5\n## Getting started\n1. [Azure Toolkit for IntelliJ](https://learn.microsoft.com/en-us/azure/developer/java/sdk/identity-dev-env-auth#intellij-credential)\n2. [Azure CLI](https://learn.microsoft.com/en-us/azure/developer/java/sdk/identity-dev-env-auth#azure-cli-credential)\nSelect each item above to learn about how to configure them for Azure Identity authentication.\n## Key concepts\n### Credentials\nA credential is a class that contains or can obtain the data needed for a service client to authentica
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 3473ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find the latest official Java sample creating Azure Key Vault SecretClient with DefaultAzureCredential and current Maven dependencies.","parameters":{"language":"java","query":"Java SecretClientBuilder DefaultAzureCredentialBuilder Azure Key Vault Secrets Maven azure-identity azure-security-keyvault-secrets"}}
```

**Output:**

```
{"results":[{"description":"description: Imports necessary Azure SDK classes and defines a Java class that retrieves secret values from Azure Key Vault using the SecretClient with DefaultAzureCredential authentication.\npackage: com.azure.identity, com.azure.security.keyvault.secrets, com.azure.core.credential\nlanguage: java\n","codeSnippet":"import com.azure.identity.DefaultAzureCredentialBuilder;\r\nimport com.azure.security.keyvault.secrets.SecretClient;\r\nimport com.azure.security.keyvault.secrets.SecretClientBuilder;\r\nimport com.azure.core.credential.AzureKeyCredential;\r\n\r\npublic class Example {\r\n\r\n    public static void main(String[] args) {\r\n\r\n        String keyVaultName = System.getenv(\u0022KEY_VAULT_NAME\u0022);\r\n        String keyVaultUri = \u0022https://\u0022 \u002B keyVaultName \u002B \u0022.vault.azure.net\u0022;\r\n\r\n        //variables for retrieving the key and endpoint from your key vault.\r\n        //Set these variables to the names you created for your secrets\r\n        String keySecretName = \u0022Your-Key-Secret-Name\u0022;\r\n        String endpointSecretName = \u0022Your-Endpoint-Secret-Name\u0022;\r\n\r\n        //Create key vault secrets client\r\n        SecretClient secretClient = new SecretClientBuilder()\r\n                .vaultUrl(keyVaultUri)\r\n                .credential(new DefaultAzureCredentialBuilder().build())\r\n                .buildClient();\r\n\r\n        //retrieve key and endpoint from key vault\r\n        String keyValue = secretClient.getSecret(keySecretName).getValue();\r\n        String endpointValue = secretClient.getSecret(endpointSecretName).getValue();\r\n        System.out.printf(\u0022Your secret key value is: %s\u0022, keyValue)\r\n        System.out.printf(\u0022Your secret endpoint value is: %s\u0022, endpointValue)\r\n    }\r\n}","link":"https://learn.microsoft.com/azure/ai-services/use-key-vault?pivots=programming-language-java#import-the-example-code-2","language":"java"},{"descript
... (truncated)
```

### ✅ 6. rg — 4671ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt","pattern":"https://learn\\.microsoft\\.com[^\" ]+|DefaultAzureCredential|WorkloadIdentityCredential|VisualStudioCodeCredential|BrokerCredential|azure-identity|azure-security-keyvault-secrets"}
```

**Output:**

```
Output too large to read at once (47.7 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787966135860-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt:{"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## Debug a chained credential\nTo diagnose an unexpected issue or to understand what a chained credential is doing, [enable logging](https://learn.microsoft.com/azure/developer/java/sdk/logging-overview) in your app.\nFor illustration purposes
```

### ✅ 7. rg — 4312ms

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt","pattern":"https://learn\\.microsoft\\.com[^\" ]+|\u003cartifactId\u003e|\u003cversion\u003e|SecretClientBuilder|DefaultAzureCredentialBuilder"}
```

**Output:**

```
Output too large to read at once (27.5 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787966135860-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt:{"results":[{"description":"description: Imports necessary Azure SDK classes and defines a Java class that retrieves secret values from Azure Key Vault using the SecretClient with DefaultAzureCredential authentication.\npackage: com.azure.identity, com.azure.security.keyvault.secrets, com.azure.core.credential\nlanguage: java\n","codeSnippet":"import com.azure.identity.DefaultAzureCredential
```

### ✅ 8. powershell — 2268ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt' | ConvertFrom-Json; $d.results | Select-Object -First 10 title,contentUrl | Format-List","description":"List documentation result URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-4daac34efbd74963bab9ee035859271c.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 1 completed with exit code 1>
```

### ✅ 9. powershell — 1856ms

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt' | ConvertFrom-Json; $d.results | Select-Object -First 10 @{n='Description';e={$_.description}},contentUrl | Format-List","description":"List code sample result URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-1f586d02c01e490e8c52709c40690dd5.txt' | ConvertFrom-Json; $d.res …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 10. view — 1486ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt b/C:/Users/weidxu/AppData/Local/Temp/1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt
@@ -1,8 +1,8 @@
 {"results":[{"description":"description: Imports necessary Azure SDK classes and defines a Java class that retrieves secret values from Azure Key Vault using the SecretClient with DefaultAzureCredential authentication.\npackage: com.azure.identity, com.azure.security.keyvault.secrets, com.azure.core.credential\nlanguage: java\n","codeSnippet":"import com.azure.identity.DefaultAzureCredentialBuilder;\r\nimport com.azure.security.keyvault.secrets.SecretClient;\r\nimport com.azure.security.keyvault.secrets.SecretClientBuilder;\r\nimport com.azure.core.credential.AzureKeyCredential;\r\n\r\npublic class Example {\r\n\r\n    public static void main(String[] args) {\r\n\r\n        String keyVaultName = System.getenv(\u0022KEY_VAULT_NAME\u0022);\r\n        String keyVaultUri = \u0022https://\u0022 \u002B keyVaultName \u002B \u0022.vault.azure.net\u0022;\r\n\r\n        //variables for retrieving the key and endpoint from your key vault.\r\n        //Set these variables to the names you created for your secrets\r\n        String keySecretName = \u0022Your-Key-Secret-Name\u0022;\r\n        String endpointSecretName = \u0022Your-Endpoint-Secret-Name\u0022;\r\n\r\n        //Create key vault secrets client\r\n        SecretClient secretClient = new SecretClientBuilder()\r\n                .vaultUrl(keyVaultUri)\r\n                .credential(new DefaultAzureCredentialBuilder().build())\r\n                .buildClient();\r\n\r\n        //retrieve key and endpoint from key vault\r\n        St
... (truncated)
```

### ✅ 11. view — 1103ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787966126460-copilot-tool-output-4daac34efbd74963bab9ee035859271c.txt","view_range":[1,8]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt b/C:/Users/weidxu/AppData/Local/Temp/1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787966126687-copilot-tool-output-1f586d02c01e490e8c52709c40690dd5.txt
@@ -1,8 +1,8 @@
 {"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## Debug a chained credential\nTo diagnose an unexpected issue or to understand what a chained credential is doing, [enable logging](https://learn.microsoft.com/azure/developer/java/sdk/logging-overview) in your app.\nFor illustration purposes, assume the parameterless form of \u0060DefaultAzureCredential\u0060 is used to authenticate a request to a Blob Storage account. The app runs in the local development environment, and the developer authenticated to Azure using the Azure CLI. When the app is run, the following pertinent entries appear in the output:\nIn the preceding output, notice that:\n1. \u0060EnvironmentCredential\u0060, \u0060WorkloadIdentityCredential\u0060, \u0060ManagedIdentityCredential\u0060, \u0060IntelliJCredential\u0060, and \u0060VisualStudioCodeCredential\u0060 each failed to acquire a Microsoft Entra access token, in that order.\n2. The \u0060AzureCliCredential.getToken\u0060 call succeeds, as indicated by the \u0060returns a token\u0060-suffixed entry. Since \u0060AzureCliCredential\u0060 succeeded, no credentials beyond it were tried.","contentUrl":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains#debug-a-chained-credential"},{"title":"Azure Identity client library for Java - version 1.18.5","content":"# Azure Identity client library for Java - version 1.18.5\n## Gett
... (truncated)
```

### ✅ 12. azure-documentation (via azure) — 3314ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the complete official Java DefaultAzureCredential chain documentation, including all credential entries and order.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains"}}
```

**Output:**

```
# Credential chains in the Azure Identity library for Java

The Azure Identity library provides *credentials*—public classes that implement the Azure Core library's [TokenCredential](/en-us/java/api/com.azure.core.credential.tokencredential) interface. A credential represents a distinct authentication flow for acquiring an access token from Microsoft Entra ID. These credentials can be chained together to form an ordered sequence of authentication mechanisms to be attempted.

## How a chained credential works

At runtime, a credential chain attempts to authenticate using the sequence's first credential. If that credential fails to acquire an access token, the next credential in the sequence is attempted, and so on, until an access token is successfully obtained. The following sequence diagram illustrates this behavior:

![Diagram that shows credential chain sequence.](../media/mermaidjs/chain-sequence.svg)

## Why use credential chains

A chained credential can offer the following benefits:

- **Environment awareness**: Automatically selects the most appropriate credential based on the environment in which the app is running. Without it, you'd have to write code like this:

    ```java
    import com.azure.core.credential.TokenCredential;
    import com.azure.identity.AzureCliCredentialBuilder;
    import com.azure.identity.ManagedIdentityCredentialBuilder;
    
    // Code omitted for brevity
    
    TokenCredential credential = null;
    
    // Set up credential based on environment (Azure or local development)
    String environment = System.getenv("ENV");
    
    if (environment != null && environment.equals("production")) {
        credential = new ManagedIdentityCredentialBuilder()
            .clientId(userAssignedClientId)
            .build();
    } else {
        credential = new AzureCliCredentialBuilder()
            .build();
    }
    ```
- **Seamless transitions**: Your app can move from local development to yo
... (truncated)
```

### ✅ 13. azure-documentation (via azure) — 2944ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch complete official troubleshooting guidance for Azure Identity authentication failures in Java.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview"}}
```

**Output:**

```
# Azure Key Vault Secret client library for Java - version 4.11.0

Azure Key Vault is a cloud service that provides secure storage for secrets, such as passwords and database connection strings.

The Azure Key Vault Secrets client library allows you to securely store and tightly control the access to tokens, passwords, API keys, and other secrets. This library offers operations to create, retrieve, update, delete, purge, backup, restore, and list the secrets and its versions.

Use the Azure Key Vault Secrets client library to create and manage secrets.

[Source code](https://github.com/Azure/azure-sdk-for-java/blob/com.azure+azure-security-keyvault-secrets_4.11.0/sdk/keyvault/azure-security-keyvault-secrets/src) | [API reference documentation](https://azure.github.io/azure-sdk-for-java) | [Product documentation](/en-us/azure/key-vault/) | [Samples](https://github.com/Azure/azure-sdk-for-java/blob/com.azure+azure-security-keyvault-secrets_4.11.0/sdk/keyvault/azure-security-keyvault-secrets/src/samples/java/com/azure/security/keyvault/secrets)

## Getting started

### Include the package

#### Include the BOM file

Please include the `azure-sdk-bom` to your project to take dependency on the General Availability (GA) version of the library. In the following snippet, replace the {bom\_version\_to\_target} placeholder with the version number. To learn more about the BOM, see the [AZURE SDK BOM README](https://github.com/Azure/azure-sdk-for-java/blob/com.azure+azure-security-keyvault-secrets_4.11.0/sdk/boms/azure-sdk-bom/README.md).

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>{bom_version_to_target}</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>
```

and then include the direct dependency in the dependencies sect
... (truncated)
```

### ✅ 14. azure-documentation (via azure) — 2549ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the current Azure Key Vault Secrets Java SDK README for dependencies and SecretClient examples.","parameters":{"url":"https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable"}}
```

**Output:**

```
# Troubleshoot Azure Identity authentication problems

This article covers failure investigation techniques, common errors for the credential types in the Azure Identity Java client library, and mitigation steps to resolve these errors. Because many credential types are available in the Azure SDK for Java, this troubleshooting guide is split into sections based on usage scenario. The following sections are available:

- [Troubleshoot Azure-hosted application authentication](troubleshooting-authentication-azure-hosted)
- [Troubleshoot development environment authentication](troubleshooting-authentication-dev-env)
- [Troubleshoot service principal authentication](troubleshooting-authentication-service-principal)
- [Troubleshoot multi-tenant authentication](troubleshooting-authentication-multi-tenant)

The remainder of this article covers general troubleshooting techniques and guidance that apply to all credential types.

## Handle Azure Identity exceptions

As noted in the [Exception handling in the Azure SDK for Java](troubleshooting-overview#exception-handling-in-the-azure-sdk-for-java) section of [Troubleshooting overview](troubleshooting-overview), the Azure SDK for Java can throw a comprehensive set of exceptions and error codes. For Azure Identity specifically, a few key exception types are important to understand.

### ClientAuthenticationException

Any service client method that makes a request to the service can raise exceptions from authentication errors. These exceptions can occur because the token is requested from the credential on the first call to the service and on any subsequent requests to the service that need to refresh the token.

To distinguish these failures from failures in the service client, Azure Identity classes raise `ClientAuthenticationException` with details describing the source of the error in the exception message and possibly the error message. Depending on the application, these errors might be recoverable. The f
... (truncated)
```

## Copilot Response

I’ll ground the Java example in the current Azure Identity and Key Vault SDK guidance, including the exact credential-chain behavior.

The key version-sensitive detail is the Java chain order, so I’m narrowing the official results to that page and its troubleshooting guidance before presenting the example.

`DefaultAzureCredential` lets the same application use developer credentials locally and workload/managed identity after deployment, without storing credentials in code. Authentication is lazy: building the client does not request a token; the first service operation does.

## 1. Maven dependencies

The current stable documentation lists `azure-identity` **1.18.5** and Key Vault Secrets **4.11.0**:

```xml
<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-identity</artifactId>
        <version>1.18.5</version>
    </dependency>

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
        <version>4.11.0</version>
    </dependency>
</dependencies>
```

For larger applications, prefer the Azure SDK BOM to keep Azure dependencies compatible:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>${azure.sdk.bom.version}</version>
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
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies>
```

If you want Visual Studio Code or OS account broker authentication, also add the optional `com.azure:azure-identity-broker` dependency using the BOM-managed version.

## 2. Creating and using `DefaultAzureCredential`

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder().build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://my-vault.vault.azure.net")
    .credential(credential)
    .buildClient();
```

Create and reuse both the credential and client rather than rebuilding them for every request. The SDK caches and refreshes access tokens automatically.

For a user-assigned managed identity, identify it explicitly:

```java
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId(System.getenv("AZURE_CLIENT_ID"))
    .build();
```

## 3. Credential chain order

With Azure Identity 1.18.5, the default chain attempts:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal configuration in environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS |
| 3 | `ManagedIdentityCredential` | Azure system-assigned or user-assigned managed identity |
| 4 | `IntelliJCredential` | Azure Toolkit for IntelliJ account |
| 5 | `VisualStudioCodeCredential` | VS Code Azure account; requires broker support |
| 6 | `AzureCliCredential` | Account authenticated with `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated with `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated with `azd auth login` |
| 9 | Broker credential | Default OS account; requires `azure-identity-broker` |

The chain stops at the first credential that successfully obtains a token. It does **not** continue merely because that identity later receives a `403 Forbidden`.

Since Azure Identity 1.10.0, developer credentials continue to be tried after authentication errors so that another development tool can succeed. Deployed-service credentials are stricter: if one is configured and can attempt authentication but fails, the error is surfaced instead of silently falling through.

You can constrain the chain with `AZURE_TOKEN_CREDENTIALS`:

```text
AZURE_TOKEN_CREDENTIALS=dev
AZURE_TOKEN_CREDENTIALS=prod
AZURE_TOKEN_CREDENTIALS=AzureCliCredential
AZURE_TOKEN_CREDENTIALS=ManagedIdentityCredential
```

`dev` excludes deployed-service credentials; `prod` excludes developer-tool credentials.

## 4. Local development versus Azure

**Local development:** Managed identity normally isn't available, so the chain proceeds to IntelliJ, VS Code, Azure CLI, Azure PowerShell, Azure Developer CLI, or broker credentials. For example, after `az login`, `AzureCliCredential` can use that signed-in account. The developer identity must have access to the Key Vault, such as the **Key Vault Secrets User** RBAC role.

**Azure hosting:** Configure workload identity or managed identity on the host and grant that identity the required Key Vault data-plane role. `DefaultAzureCredential` normally succeeds near the beginning of the chain, without client secrets or certificates.

For a system-assigned identity, no identity-specific application configuration is usually needed. For a user-assigned identity, set `AZURE_CLIENT_ID` or call `managedIdentityClientId(...)`.

Be careful with `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_SECRET`: if present together, `EnvironmentCredential` is tried first and can override the identity you expected to use.

## 5. Complete `SecretClient` example

```java
package example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultExample {
    private KeyVaultExample() {
    }

    public static void main(String[] args) {
        String vaultUrl = requireEnvironmentVariable("KEY_VAULT_URL");
        String secretName = requireEnvironmentVariable("SECRET_NAME");

        DefaultAzureCredential credential =
            new DefaultAzureCredentialBuilder().build();

        SecretClient secretClient = new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(credential)
            .buildClient();

        try {
            KeyVaultSecret secret = secretClient.getSecret(secretName);

            // Avoid writing secret.getValue() to logs or console output.
            System.out.printf(
                "Retrieved secret '%s', version '%s'.%n",
                secret.getName(),
                secret.getProperties().getVersion()
            );
        } catch (ClientAuthenticationException exception) {
            System.err.println(
                "Microsoft Entra authentication failed: " + exception.getMessage()
            );
            System.exit(1);
        } catch (HttpResponseException exception) {
            System.err.printf(
                "Key Vault request failed with HTTP %d: %s%n",
                exception.getResponse().getStatusCode(),
                exception.getMessage()
            );
            System.exit(1);
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is missing: " + name
            );
        }
        return value;
    }
}
```

Example configuration:

```text
KEY_VAULT_URL=https://my-vault.vault.azure.net
SECRET_NAME=my-secret
```

## Troubleshooting

1. **Read the complete exception.** `DefaultAzureCredential` reports each attempted credential and why it was unavailable. Microsoft Entra errors also include an AADSTS code, correlation ID, and timestamp.
2. **Enable Azure SDK logging.** Configure an SLF4J provider and enable `com.azure.identity` logging at `INFO` or `DEBUG`. This shows which credentials were attempted and which one returned the token. Never enable unsafe MSAL support logging outside controlled diagnostics because it can contain tokens and personal data.
3. **Distinguish authentication from authorization.**
   - `ClientAuthenticationException`: no usable token was obtained or configured authentication failed.
   - HTTP `401`: token/audience/tenant authentication problem.
   - HTTP `403`: authentication probably succeeded, but the selected identity lacks Key Vault permissions.
4. **Confirm the actual identity.** A cached CLI, IDE, PowerShell, or broker account may differ from the account you expected. `DefaultAzureCredential` uses the first identity that obtains a token.
5. **Check Key Vault authorization.** For Azure RBAC, assign an appropriate data-plane role such as **Key Vault Secrets User** at the narrowest practical scope. Subscription `Owner` does not automatically grant access to secret contents.
6. **Check Azure-hosted configuration.** Ensure managed identity is enabled, the correct user-assigned client ID is selected, and workload identity has `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and `AZURE_FEDERATED_TOKEN_FILE`.
7. **Check connectivity.** Validate the vault URL, private endpoint DNS, firewall rules, proxy settings, and outbound access to both Microsoft Entra ID and Key Vault.

**References:** [Java credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), [Azure Identity troubleshooting](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview), [Key Vault Secrets Java SDK](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The sample does not omit cleanup for any AutoCloseable Azure SDK client.

## Grader Results

- default-azure-credential.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/5)
      - Maven dependency for `azure-identity`: Pass
      - `DefaultAzureCredentialBuilder` pattern: Pass
      - Credential chain order in Java SDK: Pass
      - Passing credential to client builders (e.g., `SecretClientBuilder`): Pass
      - Logging configuration for authentication troubleshooting: Fail
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
| `Criteria from prompt file` | prompt_review | 80% | 1.00 | 0.8000 | 6.8% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.8000** | **90.8%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-default-credential --config java-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
