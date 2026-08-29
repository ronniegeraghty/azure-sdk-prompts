# Evaluation Report: identity-dp-java-managed-identity

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 306.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-managed-identity` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 16/18 |
| Duration | 306.8s |
| Timestamp | 2026-08-29T01:30:58Z |
| Files Generated | 0 |
| Event Count | 5385 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 103.0s |
| Review | 202.6s |
| **Total** | **306.8s** |

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
| Token Usage | in=236886 out=4121 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. azure-get_azure_bestpractices (via azure) — 5455ms

**Input:**

```json
{"intent":"Provide current best practices for authenticating Azure SDK for Java clients with system-assigned and user-assigned managed identities, including local development fallback and error handling.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 2. azure-documentation (via azure) — 6540ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure Identity Java ManagedIdentityCredential, DefaultAzureCredential local development, user-assigned identity client/resource ID options, and troubleshooting.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 3. azure-get_azure_bestpractices (via azure) — 7213ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure Java SDK authentication code-generation best practices.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 4. azure-documentation (via azure) — 11286ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official guidance for Java managed identity authentication and local fallback.","parameters":{"query":"Azure Identity Java ManagedIdentityCredential system-assigned user-assigned client ID resource ID DefaultAzureCredential local development troubleshooting credential chain"}}
```

**Output:**

```
{"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## Usage guidance for DefaultAzureCredential\n\u0060DefaultAzureCredential\u0060 is undoubtedly the easiest way to get started with the Azure Identity library, but with that convenience comes tradeoffs. After you deploy your app to Azure, you should understand the app\u0027s authentication requirements and consider whether \u0060DefaultAzureCredential\u0060 is appropriate for your scenario.\n\u0060DefaultAzureCredential\u0060 offers a key benefit: it decouples your application code from specific authentication mechanisms, enabling you to change your authentication configuration without modifying code. For experienced developers who consciously configure their production authentication, this flexibility can be valuable. However, this flexibility comes with potential drawbacks:\n1. **Debugging challenges**: When authentication fails, it can be challenging to debug and identify the offending credential. You must enable logging to see the progression from one credential to the next and the success/failure status of each. For more information, see Debug a chained credential.\n2. **Performance overhead**: The process of sequentially trying multiple credentials can introduce performance overhead. For example, when running on a local development machine, managed identity is unavailable. Consequently, \u0060ManagedIdentityCredential\u0060 always fails in the local development environment.\n3. **Unpredictable behavior**: \u0060DefaultAzureCredential\u0060 checks for the presence of certain [environment variables](https://github.com/Azure/azure-sdk-for-java/blob/main/sdk/identity/azure-identity/README.md#environment-variables). It\u0027s possible that someone could add or modify these environment variables at the system level on the host machine. Those changes apply globally and therefore alter the behavior of \u0060DefaultAzure
... (truncated)
```

### ✅ 5. azure-documentation (via azure) — 8879ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official Java code samples for ManagedIdentityCredential and Azure SDK client builders.","parameters":{"language":"java","query":"Java ManagedIdentityCredentialBuilder clientId resourceId DefaultAzureCredentialBuilder Azure SDK client builder credential"}}
```

**Output:**

```
{"results":[{"description":"description: Shows how to authenticate an Azure-hosted Java app using a user-assigned managed identity by configuring ManagedIdentityCredential with the resource ID of the managed identity or using DefaultAzureCredential for local development, then creating a SecretClient to access Azure Key Vault.\npackage: com.azure.core.credential, com.azure.identity, com.azure.security.keyvault.secrets\nlanguage: java\n","codeSnippet":"import com.azure.core.credential.TokenCredential;\r\n    import com.azure.identity.DefaultAzureCredentialBuilder;\r\n    import com.azure.identity.ManagedIdentityCredentialBuilder;\r\n    import com.azure.security.keyvault.secrets.SecretClient;\r\n    import com.azure.security.keyvault.secrets.SecretClientBuilder;\r\n    \r\n    TokenCredential credential = null;\r\n    \r\n    // Set up credential based on environment (Azure or local development)\r\n    String environment = System.getenv(\u0022ENV\u0022);\r\n    \r\n    if (environment != null \u0026\u0026 environment.equals(\u0022production\u0022)) {\r\n        // Specify the client ID of the user-assigned managed identity\r\n        credential = new ManagedIdentityCredentialBuilder()\r\n            .clientId(\u0022\u003Cuser-assigned-managed-identity-client-id\u003E\u0022)\r\n            .build();\r\n    } else {\r\n        credential = new DefaultAzureCredentialBuilder()\r\n            .build();\r\n    }\r\n    \r\n    // Azure SDK client builders accept the credential as a parameter\r\n    SecretClient client = new SecretClientBuilder()\r\n        .vaultUrl(\u0022https://\u003Cyour-key-vault-name\u003E.vault.azure.net\u0022)\r\n        .credential(credential)\r\n        .buildClient();","link":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/user-assigned-managed-identity#authenticate-to-azure-services-from-your-app","language":"java"},{"description":"description: Retrieves the client ID of a user-assigned managed identity using Azure CLI.\nlang
... (truncated)
```

### ✅ 6. rg — 8603ms

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt","pattern":"resourceId|clientId|ManagedIdentityCredentialBuilder|DefaultAzureCredentialBuilder"}
```

**Output:**

```
Output too large to read at once (50.0 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787967114223-2.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt:{"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## Usage guidance for DefaultAzureCredential\n\u0060DefaultAzureCredential\u0060 is undoubtedly the easiest way to get started with the Azure Identity library, but with that convenience comes tradeoffs. After you deploy your app to Azure, you 
```

### ✅ 7. rg — 7248ms

**Input:**

```json
{"head_limit":50,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+|ManagedIdentityCredentialBuilder|resourceId|clientId|AZURE_CLIENT_ID|DefaultAzureCredential"}
```

**Output:**

```
Output too large to read at once (23.1 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787967114222-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt:{"results":[{"description":"description: Shows how to authenticate an Azure-hosted Java app using a user-assigned managed identity by configuring ManagedIdentityCredential with the resource ID of the managed identity or using DefaultAzureCredential for local development, then creating a SecretClient to access Azure Key Vault.\npackage: com.azure.core.credential, com.azure.identity, com.azure
```

### ✅ 8. powershell — 4906ms

**Input:**

```json
{"command":"$docs = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt' | ConvertFrom-Json; $samples = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt' | ConvertFrom-Json; 'DOCS'; $docs.results | Select-Object -First 8 title,contentUrl | Format-Table -AutoSize; 'SAMPLES'; $samples.results | Select-Object -First 6 description,codeSampleUrl | Format-List","description":"Extract relevant Azure documentation titles and URLs"}
```

**Output:**

```
ConvertFrom-Json: 
Line |
   2 |  … tput-d58fcd0a33a24733b720181bc0330a94.txt' | ConvertFrom-Json; $sampl …
     |                                                 ~~~~~~~~~~~~~~~~
     | Conversion from JSON failed with error: Additional text encountered after finished reading JSON content: {. Path '', line 3, position 0.
<shellId: 0 completed with exit code 1>
```

### ✅ 9. view — 6566ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt","view_range":[1,5]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt b/C:/Users/weidxu/AppData/Local/Temp/1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt
@@ -1,5 +1,5 @@
 {"results":[{"description":"description: Shows how to authenticate an Azure-hosted Java app using a user-assigned managed identity by configuring ManagedIdentityCredential with the resource ID of the managed identity or using DefaultAzureCredential for local development, then creating a SecretClient to access Azure Key Vault.\npackage: com.azure.core.credential, com.azure.identity, com.azure.security.keyvault.secrets\nlanguage: java\n","codeSnippet":"import com.azure.core.credential.TokenCredential;\r\n    import com.azure.identity.DefaultAzureCredentialBuilder;\r\n    import com.azure.identity.ManagedIdentityCredentialBuilder;\r\n    import com.azure.security.keyvault.secrets.SecretClient;\r\n    import com.azure.security.keyvault.secrets.SecretClientBuilder;\r\n    \r\n    TokenCredential credential = null;\r\n    \r\n    // Set up credential based on environment (Azure or local development)\r\n    String environment = System.getenv(\u0022ENV\u0022);\r\n    \r\n    if (environment != null \u0026\u0026 environment.equals(\u0022production\u0022)) {\r\n        // Specify the client ID of the user-assigned managed identity\r\n        credential = new ManagedIdentityCredentialBuilder()\r\n            .clientId(\u0022\u003Cuser-assigned-managed-identity-client-id\u003E\u0022)\r\n            .build();\r\n    } else {\r\n        credential = new DefaultAzureCredentialBuilder()\r\n            .build();\r\n    }\r\n    \r\n    // Azure SDK client builders accept the credential as a pa
... (truncated)
```

### ✅ 10. view — 4644ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787967097656-copilot-tool-output-c2fe7fc667e84454abfaaf6fda0a4f85.txt","view_range":[1,5]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt b/C:/Users/weidxu/AppData/Local/Temp/1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt
+++ b/C:/Users/weidxu/AppData/Local/Temp/1787967098544-copilot-tool-output-d58fcd0a33a24733b720181bc0330a94.txt
@@ -1,5 +1,5 @@
 {"results":[{"title":"Credential chains in the Azure Identity library for Java","content":"# Credential chains in the Azure Identity library for Java\n## Usage guidance for DefaultAzureCredential\n\u0060DefaultAzureCredential\u0060 is undoubtedly the easiest way to get started with the Azure Identity library, but with that convenience comes tradeoffs. After you deploy your app to Azure, you should understand the app\u0027s authentication requirements and consider whether \u0060DefaultAzureCredential\u0060 is appropriate for your scenario.\n\u0060DefaultAzureCredential\u0060 offers a key benefit: it decouples your application code from specific authentication mechanisms, enabling you to change your authentication configuration without modifying code. For experienced developers who consciously configure their production authentication, this flexibility can be valuable. However, this flexibility comes with potential drawbacks:\n1. **Debugging challenges**: When authentication fails, it can be challenging to debug and identify the offending credential. You must enable logging to see the progression from one credential to the next and the success/failure status of each. For more information, see Debug a chained credential.\n2. **Performance overhead**: The process of sequentially trying multiple credentials can introduce performance overhead. For example, when running on a local development machine, managed identity is unavailable. Consequently, \u0060ManagedIdentityCredential\u006
... (truncated)
```

## Copilot Response

I’ll ground the Java patterns in current Azure guidance, then provide copy-ready examples and a troubleshooting checklist.

The core production recommendation is clear: use a deterministic `ManagedIdentityCredential` in Azure and reserve credential chains for development. I’m checking the exact Java builder APIs and supported identity selectors before finalizing the snippets.

Managed identity lets Azure-hosted Java applications obtain Microsoft Entra tokens without storing secrets. Grant the identity the required **data-plane RBAC role**, then pass a `TokenCredential` to an Azure SDK client builder.

## 1. System-assigned vs. user-assigned

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created on and tied to one Azure resource | Independent Azure resource |
| Sharing | Used only by its owning resource | Can be attached to multiple resources |
| Deletion | Deleted with the owning resource | Remains until explicitly deleted |
| Credential configuration | No identity ID required | Specify client ID, resource ID, or object ID |
| Typical use | One workload with simple lifecycle | Shared identity or stable permissions across deployments |

For authorization, assign roles to the identity’s **principal/object ID**. When selecting a user-assigned identity in Java, the most portable option is usually its **client ID**.

## 2. Dependencies

Use the Azure SDK BOM to manage compatible versions, then include `azure-identity` and the client library you need:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>${azure-sdk-bom.version}</version>
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

    <!-- Example service client -->
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies>
```

Use the current stable BOM version from the [Azure SDK releases page](https://azure.github.io/azure-sdk/releases/latest/java.html).

## 3. Creating each credential

### System-assigned managed identity

No identifier is required:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .build();
```

If several identities are attached but no system-assigned identity exists, this form won't automatically choose a user-assigned identity—you must identify it explicitly.

### User-assigned managed identity by client ID

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

String managedIdentityClientId =
    System.getenv("MANAGED_IDENTITY_CLIENT_ID");

if (managedIdentityClientId == null || managedIdentityClientId.isBlank()) {
    throw new IllegalStateException(
        "MANAGED_IDENTITY_CLIENT_ID must contain the user-assigned identity client ID");
}

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(managedIdentityClientId)
    .build();
```

`MANAGED_IDENTITY_CLIENT_ID` must be the identity’s **client/application ID**, not its principal/object ID.

You can instead select it by Azure resource ID:

```java
TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .resourceId(System.getenv("MANAGED_IDENTITY_RESOURCE_ID"))
    .build();
```

A resource ID has this form:

```text
/subscriptions/<subscription-id>/resourceGroups/<resource-group>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity-name>
```

Recent `azure-identity` versions also support `.objectId(...)`, but client ID or resource ID generally makes intent clearer.

## 4. Using the credential with Azure SDK clients

All modern Azure SDK builders that use Microsoft Entra authentication accept `TokenCredential`.

### Key Vault with system-assigned identity

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl(System.getenv("KEY_VAULT_URL"))
    .credential(credential)
    .buildClient();

String value = secretClient
    .getSecret("database-password")
    .getValue();
```

The identity needs an appropriate role, such as **Key Vault Secrets User**, scoped to the vault or narrower.

### Key Vault with user-assigned identity

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

String clientId = System.getenv("MANAGED_IDENTITY_CLIENT_ID");
String vaultUrl = System.getenv("KEY_VAULT_URL");

if (clientId == null || clientId.isBlank()) {
    throw new IllegalStateException("MANAGED_IDENTITY_CLIENT_ID is required");
}
if (vaultUrl == null || vaultUrl.isBlank()) {
    throw new IllegalStateException("KEY_VAULT_URL is required");
}

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(clientId)
    .build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl(vaultUrl)
    .credential(credential)
    .buildClient();

String value = secretClient
    .getSecret("database-password")
    .getValue();
```

The same credential can be reused across clients:

```java
BlobServiceClient blobClient = new BlobServiceClientBuilder()
    .endpoint(System.getenv("STORAGE_BLOB_ENDPOINT"))
    .credential(credential)
    .buildClient();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl(System.getenv("KEY_VAULT_URL"))
    .credential(credential)
    .buildClient();
```

Credential and service-client instances are designed to be reused rather than recreated per request.

## 5. Local development fallback

Managed identity endpoints exist only in supported Azure hosting environments. For local development, use `DefaultAzureCredential`, which can authenticate through developer tools such as Azure CLI, Azure Developer CLI, IntelliJ, or environment-based credentials.

Use an explicit environment switch so production never silently falls back to a developer credential:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

public final class CredentialFactory {
    private CredentialFactory() {
    }

    public static TokenCredential create() {
        boolean azureHosted =
            Boolean.parseBoolean(System.getenv("AZURE_HOSTED"));

        if (!azureHosted) {
            return new DefaultAzureCredentialBuilder()
                .excludeManagedIdentityCredential()
                .build();
        }

        String clientId = System.getenv("MANAGED_IDENTITY_CLIENT_ID");

        if (clientId == null || clientId.isBlank()) {
            return new ManagedIdentityCredentialBuilder()
                .build(); // System-assigned
        }

        return new ManagedIdentityCredentialBuilder()
            .clientId(clientId)
            .build(); // User-assigned
    }
}
```

Example configuration:

```text
# Local
AZURE_HOSTED=false

# Azure, system-assigned
AZURE_HOSTED=true

# Azure, user-assigned
AZURE_HOSTED=true
MANAGED_IDENTITY_CLIENT_ID=<identity-client-id>
```

Important considerations:

- `DefaultAzureCredentialBuilder.managedIdentityClientId(...)` selects a user-assigned identity **when managed identity is available**. It does not make your local Azure CLI session impersonate that identity.
- Give developers equivalent but preferably narrower access using their own Entra accounts.
- For CI/CD, prefer workload identity federation. Use a service principal only when federation isn't available, and never commit its secret.
- In production, prefer `ManagedIdentityCredential` over an unrestricted `DefaultAzureCredential` chain for predictable identity selection.

## 6. Error handling

Token acquisition is generally lazy, so credential errors often occur on the first service operation rather than while constructing the client.

```java
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.AuthenticationFailedException;
import com.azure.identity.CredentialUnavailableException;
import com.azure.security.keyvault.secrets.SecretClient;

public final class SecretReader {
    private final SecretClient client;

    public SecretReader(SecretClient client) {
        this.client = client;
    }

    public String read(String name) {
        try {
            return client.getSecret(name).getValue();
        } catch (CredentialUnavailableException e) {
            throw new IllegalStateException(
                "Managed identity is unavailable in this hosting environment", e);
        } catch (AuthenticationFailedException e) {
            throw new IllegalStateException(
                "Microsoft Entra token acquisition failed", e);
        } catch (HttpResponseException e) {
            int status = e.getResponse().getStatusCode();

            if (status == 403) {
                throw new IllegalStateException(
                    "Authentication succeeded, but the identity lacks permission", e);
            }

            throw e;
        }
    }
}
```

Do not log access tokens, secrets, authorization headers, or full credential configuration.

## Troubleshooting checklist

1. **Credential unavailable or metadata endpoint timeout**
   - Confirm the app is running in a managed-identity-supported Azure service.
   - Confirm managed identity is enabled and, for user-assigned identity, attached to the host.
   - Don't expect `ManagedIdentityCredential` to work on a developer workstation.

2. **“Multiple user-assigned identities” or wrong identity**
   - Specify `.clientId(...)` or `.resourceId(...)`.
   - Verify that the value is the identity’s client ID—not its principal ID.

3. **HTTP 401**
   - Check the service endpoint and cloud/authority configuration.
   - Ensure the Azure service supports Microsoft Entra authentication.
   - Check system time and network access to the Azure identity endpoint.

4. **HTTP 403**
   - Token acquisition succeeded; this is normally authorization, not authentication.
   - Assign the correct **data-plane** role to the managed identity’s principal ID.
   - Check role scope and allow time for role assignments to propagate.
   - Management-plane roles such as Contributor don't necessarily grant data access.

5. **Works locally but fails in Azure**
   - Determine which local identity `DefaultAzureCredential` selected.
   - Compare that developer’s permissions with the managed identity’s permissions.
   - Ensure the deployed app uses `ManagedIdentityCredential`, not an unexpected environment credential.

6. **Enable diagnostics**
   - Enable Azure SDK logging temporarily and inspect the credential type and endpoint being used.
   - Avoid verbose identity logging in normal production operation because logs can contain sensitive metadata.

**References:** [System-assigned identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/system-assigned-managed-identity), [User-assigned identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/user-assigned-managed-identity), [Java credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), and [Azure Identity Java API](https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The response uses modern Azure SDK client builders without introducing any unclosed AutoCloseable client instances.

## Grader Results

- managed-identity-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/6)
      - `ManagedIdentityCredentialBuilder` class: Pass
      - System-assigned: default builder with no client ID: Pass
      - User-assigned: `.clientId()` on the builder: Pass
      - Integration with `DefaultAzureCredential` chain: Pass
      - `CredentialUnavailableException` when not in Azure: Pass
      - `ChainedTokenCredentialBuilder` for local fallback: Fail
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
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Fail
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 7.0% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.5% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.8333** | **91.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-managed-identity --config java-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
