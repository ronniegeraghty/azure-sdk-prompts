# Evaluation Report: identity-dp-java-credential-chain

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 432.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-credential-chain` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 24/26 |
| Duration | 432.0s |
| Timestamp | 2026-08-29T00:52:25Z |
| Files Generated | 10 |
| Event Count | 11436 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 200.0s |
| Review | 230.8s |
| **Total** | **432.0s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | entra-agent-id |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=755882 out=9689 |
| Turn Count | 13 |

## Tool Calls

### ✅ 1. skill — 2434ms

**Input:**

```json
{"skill":"entra-agent-id"}
```

**Output:**

```
Skill loaded successfully ✅

# Microsoft Entra Agent ID

Create and manage OAuth 2.0-capable identities for AI agents using Microsoft Graph. Every agent instance gets a distinct identity, audit trail, and independently-scoped permission grants.

## Quick Reference

| Property | Value |
|----------|-------|
| Service | Microsoft Entra Agent ID |
| API | Microsoft Graph (`https://graph.microsoft.com/v1.0`) |
| Required role | Agent Identity Developer, Agent Identity Administrator, or Application Administrator |
| Object model | Blueprint (application) → BlueprintPrincipal (SP) → Agent Identity (SP) |
| Runtime exchange | Two-step `fmi_path` exchange (autonomous and OBO) |
| .NET helper | `Microsoft.Identity.Web.AgentIdentities` |
| Polyglot helper | Microsoft Entra SDK for AgentID (sidecar container) |

## When to Use This Skill

- Provisioning a new Agent Identity Blueprint and BlueprintPrincipal
- Creating per-instance Agent Identities under a Blueprint
- Configuring credentials (FIC, Managed Identity, or client secret) on the Blueprint
- Implementing the two-step `fmi_path` runtime token exchange (autonomous or OBO)
- Cross-tenant agent token flows
- Deploying the Microsoft Entra SDK for AgentID sidecar for polyglot agents (Python, Node, Go, Java)
- Granting per-Agent-Identity application (`appRoleAssignments`) or delegated (`oauth2PermissionGrants`) permissions
- Diagnosing Agent ID errors such as `AADSTS82001`, `AADSTS700211`, or `PropertyNotCompatibleWithAgentIdentity`

## MCP Tools

| Tool | Use |
|------|-----|
| `mcp_azure_mcp_documentation` | Search Microsoft Learn for current Agent ID setup, Graph API shapes, and SDK configuration |

There is no dedicated Agent Identity MCP server today. This skill guides direct Microsoft Graph API calls (PowerShell or Python `requests`). Use `mcp_azure_mcp_documentation` to verify request bodies and endpoints against current docs before running.

## Before You Start

Use the `mcp_azure_mcp_documentation` tool to sear
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 5136ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Identity credentials, chained credentials, managed identity, workload identity, Azure Pipelines credentials, and CAE. Need secure modern SDK patterns without deploying Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 7088ms

**Input:**

```json
{"intent":"Find current official Java Azure Identity documentation for ChainedTokenCredential, AzureCliCredential, IntelliJCredential, AzureDeveloperCliCredential, EnvironmentCredential, AzurePipelinesCredential, ManagedIdentityCredential, WorkloadIdentityCredential, TokenRequestContext CAE, and AuthenticationFailedException diagnostics.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 6062ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 6591ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure SDK code-generation best practices to a Java 17 Maven credential-chain sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 10457ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Verify current Java Azure Identity APIs and recommended credential patterns.","parameters":{"query":"Azure Identity Java custom credential chain ChainedTokenCredential AzurePipelinesCredential ManagedIdentityCredential WorkloadIdentityCredential CAE TokenRequestContext enableCae authentication diagnostics"}}
```

**Output:**

```
{"results":[{"title":"Azure Identity client library for Java - version 1.18.5","content":"# Azure Identity client library for Java - version 1.18.5\n## Managed Identity support\n\u0060\u0060\u0060java\n/**\n * Authenticate with a system-assigned managed identity.\n */\npublic void createManagedIdentityCredential() {\n    ManagedIdentityCredential managedIdentityCredential = new ManagedIdentityCredentialBuilder()\n        .build();\n\n    // Azure SDK client builders accept the credential as a parameter\n    SecretClient client = new SecretClientBuilder()\n        .vaultUrl(\u0022https://{YOUR_VAULT_NAME}.vault.azure.net\u0022)\n        .credential(managedIdentityCredential)\n        .buildClient();\n}\n\u0060\u0060\u0060\n### Define a custom authentication flow with \u0060ChainedTokenCredential\u0060\nWhile \u0060DefaultAzureCredential\u0060 is generally the quickest way to authenticate apps for Azure, you can create a customized chain of credentials to be considered. \u0060ChainedTokenCredential\u0060 enables users to combine multiple credential instances to define a customized chain of credentials. For more information, see [ChainedTokenCredential overview](https://aka.ms/azsdk/java/identity/credential-chains#chainedtokencredential-overview).\n## Cloud / Sovereign configuration\nBy default, credentials authenticate to the Microsoft Entra endpoint for Azure Public Cloud. To access resources in other clouds, such as Azure US Government or a private cloud, use one of the following solutions:\n1. Configure credentials with the \u0060authorityHost\u0060 method. For example:\n\u0060\u0060\u0060java\nDefaultAzureCredential defaultAzureCredential = new DefaultAzureCredentialBuilder()\n    .authorityHost(AzureAuthorityHosts.AZURE_GOVERNMENT)\n    .build();\n\u0060\u0060\u0060\n[AzureAuthorityHosts](https://learn.microsoft.com/en-us/java/api/com.azure.identity.azureauthorityhosts?view=azure-java-stable) defines authorities for well-known clouds.\n1. Set the \u0060AZURE_AUTH
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 8650ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java code samples for Azure Identity credential chains and CAE.","parameters":{"language":"java","query":"Java azure-identity ChainedTokenCredential AzurePipelinesCredential ManagedIdentityCredential WorkloadIdentityCredential TokenRequestContext setCaeEnabled"}}
```

**Output:**

```
{"results":[{"description":"description: Creates a ChainedTokenCredential instance by chaining together AzureCliCredential and IntelliJCredential, allowing sequential authentication attempts using these credentials.\npackage: com.azure.identity\nlanguage: java\n","codeSnippet":"import com.azure.identity.AzureCliCredential;\r\nimport com.azure.identity.AzureCliCredentialBuilder;\r\nimport com.azure.identity.ChainedTokenCredential;\r\nimport com.azure.identity.ChainedTokenCredentialBuilder;\r\nimport com.azure.identity.IntelliJCredential;\r\nimport com.azure.identity.IntelliJCredentialBuilder;\r\n\r\n// Code omitted for brevity\r\n\r\nAzureCliCredential cliCredential = new AzureCliCredentialBuilder()\r\n    .build();\r\nIntelliJCredential ijCredential = new IntelliJCredentialBuilder()\r\n    .build();\r\n\r\nChainedTokenCredential credential = new ChainedTokenCredentialBuilder()\r\n    .addLast(cliCredential)\r\n    .addLast(ijCredential)\r\n    .build();","link":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains#chainedtokencredential-overview","language":"java"},{"description":"description: Imports necessary classes for Azure authentication and communication management client initialization in Java.\npackage: com.azure.core.credential, com.azure.core.management, com.azure.identity, com.azure.resourcemanager.communication\nlanguage: java\n","codeSnippet":"import com.azure.core.credential.TokenCredential;\r\nimport com.azure.core.management.AzureEnvironment;\r\nimport com.azure.core.management.profile.AzureProfile;\r\nimport com.azure.identity.DefaultAzureCredentialBuilder;\r\nimport com.azure.resourcemanager.communication.CommunicationManager;","link":"https://learn.microsoft.com/azure/communication-services/quickstarts/email/manage-suppression-list-management-sdks?pivots=programming-language-java#initialize-the-management-client","language":"java"},{"description":"description: Creates a DefaultAzureCredential instance that require
... (truncated)
```

### ✅ 8. glob — 7444ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-agent-id","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 9. rg — 4749ms

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787964796519-copilot-tool-output-774650ab44584a38a0ce67b38056d7f1.txt","pattern":"AzurePipelinesCredential|setCaeEnabled|enableCae|WorkloadIdentityCredential|AuthenticationFailedException"}
```

**Output:**

```
Output too large to read at once (44.2 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787964813715-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787964796519-copilot-tool-output-774650ab44584a38a0ce67b38056d7f1.txt:{"results":[{"title":"Azure Identity client library for Java - version 1.18.5","content":"# Azure Identity client library for Java - version 1.18.5\n## Managed Identity support\n\u0060\u0060\u0060java\n/**\n * Authenticate with a system-assigned managed identity.\n */\npublic void createManagedIdentityCredential() {\n    ManagedIdentityCredential managedIdentityCredential = new ManagedIdenti
```

### ✅ 10. azure-documentation (via azure) — 5497ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the full official Java credential chain guide before implementing the sample.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains"}}
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

### ✅ 11. web_fetch — 5116ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 12. web_fetch — 3748ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredentialBuilder.java"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/core/azure-core/src/main/java/com/azure/core/credential/TokenRequestContext.java:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.core.credential;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Objects;

/**
 * <p>
 * The {@link TokenRequestContext} is a class used to provide additional information and context when requesting an
 * access token from an authentication source. It allows you to customize the token request and specify additional
 * parameters, such as scopes, claims, or authentication options.
 * </p>
 *
 * <p>
 * The {@link TokenRequestContext} is typically used with authentication mechanisms that require more advanced
 * configurations or options, such as
 * <a href="https://learn.microsoft.com/azure/active-directory/fundamentals/">Azure Active Directory (Azure AD)</a>
 * authentication.
 * </p>
 *
 * <p>
 * Here's a high-level overview of how you can use the {@link TokenRequestContext}:
 * </p>
 *
 * <ol>
 * <li>Create an instance of the {@link TokenRequestContext} class and configure the required properties.
 * The {@link TokenRequestContext} class allows you to specify the scopes or resources for which you want to request
 * an access token, as well as any additional claims or options.</li>
 *
 * <li>Pass the TokenRequestContext instance to the appropriate authentication client or mechanism when
 * requesting an access token. The specific method or API to do this will depend on the authentication mechanism
 * you are using. For example, if you are using Azure Identity for AAD authentication, you would pass the
 * TokenRequestContext instance to the getToken method of the {@link TokenCredential} implementation.</li>
 *
 * <li>The authentication client or mechanis
... (truncated)
```

### ✅ 13. web_fetch — 3374ms

**Input:**

```json
{"max_length":8000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/src/main/java/com/azure/identity/WorkloadIdentityCredentialBuilder.java"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/src/main/java/com/azure/identity/WorkloadIdentityCredentialBuilder.java:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.identity;

import com.azure.core.util.Configuration;
import com.azure.core.util.CoreUtils;
import com.azure.core.util.logging.ClientLogger;
import com.azure.identity.implementation.customtokenproxy.CustomTokenProxyConfiguration;
import com.azure.identity.implementation.customtokenproxy.CustomTokenProxyHttpClient;
import com.azure.identity.implementation.customtokenproxy.ProxyConfig;
import com.azure.identity.implementation.util.ValidationUtil;

import static com.azure.identity.ManagedIdentityCredential.AZURE_FEDERATED_TOKEN_FILE;

/**
 * Fluent credential builder for instantiating a {@link WorkloadIdentityCredential}.
 *
 * <p>Workload Identity authentication is a feature in Azure that allows applications running on virtual machines (VMs)
 * to access other Azure resources without the need for a service principal or managed identity. With Workload Identity
 * authentication, applications authenticate themselves using their own identity, rather than using a shared service
 * principal or managed identity. Under the hood, Workload Identity authentication uses the concept of Service Account
 * Credentials (SACs), which are automatically created by Azure and stored securely in the VM. By using Workload
 * Identity authentication, you can avoid the need to manage and rotate service principals or managed identities for
 * each application on each VM. Additionally, because SACs are created automatically and managed by Azure, you don't
 * need to worry about storing and securing sensitive credentials themselves.
 * The WorkloadIdentityCredential supports Azure workload identity authenticatio
... (truncated)
```

### ✅ 14. web_fetch — 2997ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/README.md"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/README.md:
# Azure Identity client library for Java

The Azure Identity library provides [Microsoft Entra ID](https://learn.microsoft.com/entra/fundamentals/whatis) token-based authentication support across the Azure SDK. It provides a set of [TokenCredential](https://learn.microsoft.com/java/api/com.azure.core.credential.tokencredential?view=azure-java-stable) implementations that can be used to construct Azure SDK clients that support Microsoft Entra token authentication.

[Source code][source] | [API reference documentation][javadoc] | [Microsoft Entra ID documentation][entraid_doc]

## Getting started

### Include the package

#### Include the BOM file

Include the `azure-sdk-bom` in your project to take a dependency on the stable version of the library. In the following snippet, replace the `{bom_version_to_target}` placeholder with the version number. To learn more about the BOM, see the [Azure SDK BOM README](https://github.com/Azure/azure-sdk-for-java/blob/main/sdk/boms/azure-sdk-bom/README.md).

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

Then include the direct dependency in the `dependencies` section without the version tag:

```xml
<dependencies>
  <dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
  </dependency>
</dependencies>
```

#### Include direct dependency

To take dependency on a particular version of the library that isn't present in the BOM, add the direct dependency to your project as follows:

[//]: # (
... (truncated)
```

### ✅ 15. web_fetch — 2614ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/core/azure-core/src/main/java/com/azure/core/credential/TokenRequestContext.java"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredentialBuilder.java:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.identity;

import com.azure.core.http.policy.HttpLogOptions;
import com.azure.core.util.Configuration;
import com.azure.core.util.logging.ClientLogger;
import com.azure.identity.implementation.util.ValidationUtil;

import java.util.Arrays;

/**
 * The {@link AzurePipelinesCredentialBuilder} provides a fluent builder for {@link AzurePipelinesCredential}.
 *
 * <!-- src_embed com.azure.identity.credential.azurepipelinescredential.construct -->
 * <pre>
 * &#47;&#47; serviceConnectionId is retrieved from the portal.
 * &#47;&#47; systemAccessToken is retrieved from the pipeline environment as shown.
 * &#47;&#47; You may choose another name for this variable.
 *
 * String systemAccessToken = System.getenv&#40;&quot;SYSTEM_ACCESSTOKEN&quot;&#41;;
 * AzurePipelinesCredential credential = new AzurePipelinesCredentialBuilder&#40;&#41;.clientId&#40;clientId&#41;
 *     .tenantId&#40;tenantId&#41;
 *     .serviceConnectionId&#40;serviceConnectionId&#41;
 *     .systemAccessToken&#40;systemAccessToken&#41;
 *     .build&#40;&#41;;
 * </pre>
 * <!-- end com.azure.identity.credential.azurepipelinescredential.construct -->
 */
public class AzurePipelinesCredentialBuilder extends AadCredentialBuilderBase<AzurePipelinesCredentialBuilder> {
    private static final ClientLogger LOGGER = new ClientLogger(AzurePipelinesCredentialBuilder.class);
    private static final String OIDC_API_VERSION = "7.1";
    private String serviceConnectionId;
    private String systemAccessToken;

    /**
     * Creates an instance of the {@link AzurePipelinesCredentialBuilder}.
     */
    public AzurePipelinesCredential
... (truncated)
```

### ✅ 16. apply_patch — 741ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003ecredential-chain-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure credential chain demo\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity-broker\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.azure.identity.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/azure/identity/DeploymentEnvironment.java\n+package com.example.azure.identity;\n+\n+public enum DeploymentEnvironment {\n+    DEV,\n+    CI,\n+    PRODUCTION\n+}\n*** Add File: src/main/java/com/example/azure/identity/EnvironmentDetector.java\n+package com.example.azure.identity;\n+\n+import java.util.Locale;\n+import java.util.Map;\n+import java.util.Set;\n+\n+public final class EnvironmentDetector {\n+    private static final Set\u003cString\u003e CI_MARKERS = Set.of(\n+        \"CI\",\n+        \"TF_BUILD\",\n+        \"BUILD_BUILDID\",\n+        \"PIPELINE_WORKSPACE\",\n+        \"GITHUB_ACTIONS\",\n+        \"GITLAB_CI\",\n+        \"JENKINS_URL\"\n+    );\n+    private static final Set\u003cString\u003e MANAGED_IDENTITY_MARKERS = Set.of(\n+        \"IDENTITY_ENDPOINT\",\n+        \"MSI_ENDPOINT\",\n+        \"IMDS_ENDPOINT\"\n+    );\n+\n+    private final Map\u003cString, String\u003e environment;\n+\n+    public EnvironmentDetector() {\n+        this(System.getenv());\n+    }\n+\n+    EnvironmentDetector(Map\u003cString, String\u003e environment) {\n+        this.environment = Map.copyOf(environment);\n+    }\n+\n+    public DeploymentEnvironment detect() {\n+        String override = environment.get(\"APP_ENVIRONMENT\");\n+        if (isPresent(override)) {\n+            return switch (override.trim().toLowerCase(Locale.ROOT)) {\n+                case \"dev\", \"development\", \"local\" -\u003e DeploymentEnvironment.DEV;\n+                case \"ci\", \"pipeline\" -\u003e DeploymentEnvironment.CI;\n+                case \"prod\", \"production\" -\u003e DeploymentEnvironment.PRODUCTION;\n+                default -\u003e throw new IllegalArgumentException(\n+                    \"APP_ENVIRONMENT must be one of: dev, ci, production\");\n+            };\n+        }\n+\n+        if (CI_MARKERS.stream().anyMatch(this::hasVariable)) {\n+            return DeploymentEnvironment.CI;\n+        }\n+        if (MANAGED_IDENTITY_MARKERS.stream().anyMatch(this::hasVariable) || hasWorkloadIdentityConfiguration()) {\n+            return DeploymentEnvironment.PRODUCTION;\n+        }\n+        return DeploymentEnvironment.DEV;\n+    }\n+\n+    private boolean hasWorkloadIdentityConfiguration() {\n+        return hasVariable(\"AZURE_FEDERATED_TOKEN_FILE\")\n+            \u0026\u0026 hasVariable(\"AZURE_TENANT_ID\")\n+            \u0026\u0026 hasVariable(\"AZURE_CLIENT_ID\");\n+    }\n+\n+    private boolean hasVariable(String name) {\n+        return isPresent(environment.get(name));\n+    }\n+\n+    private static boolean isPresent(String value) {\n+        return value != null \u0026\u0026 !value.isBlank();\n+    }\n+}\n*** Add File: src/main/java/com/example/azure/identity/CaeTokenCredential.java\n+package com.example.azure.identity;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+final class CaeTokenCredential implements TokenCredential {\n+    private final TokenCredential delegate;\n+\n+    CaeTokenCredential(TokenCredential delegate) {\n+        this.delegate = Objects.requireNonNull(delegate, \"delegate\");\n+    }\n+\n+    @Override\n+    public Mono\u003cAccessToken\u003e getToken(TokenRequestContext request) {\n+        TokenRequestContext caeRequest = new TokenRequestContext()\n+            .setScopes(request.getScopes())\n+            .setClaims(request.getClaims())\n+            .setTenantId(request.getTenantId())\n+            .setProofOfPossessionOptions(request.getProofOfPossessionOptions())\n+            .setCaeEnabled(true);\n+        return delegate.getToken(caeRequest);\n+    }\n+}\n*** Add File: src/main/java/com/example/azure/identity/CredentialFactory.java\n+package com.example.azure.identity;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.AzureCliCredentialBuilder;\n+import com.azure.identity.AzureDeveloperCliCredentialBuilder;\n+import com.azure.identity.AzurePipelinesCredentialBuilder;\n+import com.azure.identity.AzurePowerShellCredentialBuilder;\n+import com.azure.identity.ChainedTokenCredentialBuilder;\n+import com.azure.identity.EnvironmentCredentialBuilder;\n+import com.azure.identity.IntelliJCredentialBuilder;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.identity.VisualStudioCodeCredentialBuilder;\n+import com.azure.identity.WorkloadIdentityCredentialBuilder;\n+\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class CredentialFactory {\n+    public static final String USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID\n+        = \"AZURE_MANAGED_IDENTITY_CLIENT_ID\";\n+\n+    private final Map\u003cString, String\u003e environment;\n+\n+    public CredentialFactory() {\n+        this(System.getenv());\n+    }\n+\n+    CredentialFactory(Map\u003cString, String\u003e environment) {\n+        this.environment = Map.copyOf(environment);\n+    }\n+\n+    public BuiltCredential create(DeploymentEnvironment deploymentEnvironment, boolean caeEnabled) {\n+        BuiltCredential built = switch (deploymentEnvironment) {\n+            case DEV -\u003e createDevelopmentCredential();\n+            case CI -\u003e createCiCredential();\n+            case PRODUCTION -\u003e createProductionCredential();\n+        };\n+\n+        TokenCredential credential = caeEnabled ? new CaeTokenCredential(built.credential()) : built.credential();\n+        return new BuiltCredential(credential, built.strategy(), caeEnabled);\n+    }\n+\n+    private BuiltCredential createDevelopmentCredential() {\n+        TokenCredential credential = new ChainedTokenCredentialBuilder()\n+            .addLast(new AzureCliCredentialBuilder().build())\n+            .addLast(new IntelliJCredentialBuilder().build())\n+            .addLast(new VisualStudioCodeCredentialBuilder().build())\n+            .addLast(new AzureDeveloperCliCredentialBuilder().build())\n+            .addLast(new AzurePowerShellCredentialBuilder().build())\n+            .build();\n+\n+        return new BuiltCredential(\n+            credential,\n+            \"Azure CLI -\u003e IntelliJ -\u003e Visual Studio Code -\u003e Azure Developer CLI -\u003e Azure PowerShell\",\n+            false\n+        );\n+    }\n+\n+    private BuiltCredential createCiCredential() {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        List\u003cString\u003e strategies = new ArrayList\u003c\u003e();\n+\n+        if (hasAzurePipelinesServiceConnectionConfiguration()) {\n+            chain.addLast(new AzurePipelinesCredentialBuilder()\n+                .tenantId(required(\"AZURE_TENANT_ID\"))\n+                .clientId(required(\"AZURE_CLIENT_ID\"))\n+                .serviceConnectionId(required(\"AZURE_SERVICE_CONNECTION_ID\"))\n+                .systemAccessToken(required(\"SYSTEM_ACCESSTOKEN\"))\n+                .build());\n+            strategies.add(\"Azure Pipelines workload identity service connection\");\n+        }\n+\n+        chain.addLast(new EnvironmentCredentialBuilder().build());\n+        strategies.add(\"environment-configured service principal (secret or certificate)\");\n+\n+        return new BuiltCredential(chain.build(), String.join(\" -\u003e \", strategies), false);\n+    }\n+\n+    private BuiltCredential createProductionCredential() {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        String managedIdentityClientId = environment.get(USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID);\n+        String managedIdentityStrategy;\n+\n+        if (isPresent(managedIdentityClientId)) {\n+            chain.addLast(new ManagedIdentityCredentialBuilder().clientId(managedIdentityClientId).build());\n+            managedIdentityStrategy = \"user-assigned managed identity\";\n+        } else {\n+            chain.addLast(new ManagedIdentityCredentialBuilder().build());\n+            managedIdentityStrategy = \"system-assigned managed identity\";\n+        }\n+\n+        String strategy = managedIdentityStrategy;\n+        if (hasWorkloadIdentityConfiguration()) {\n+            chain.addLast(new WorkloadIdentityCredentialBuilder()\n+                .tenantId(required(\"AZURE_TENANT_ID\"))\n+                .clientId(required(\"AZURE_CLIENT_ID\"))\n+                .tokenFilePath(required(\"AZURE_FEDERATED_TOKEN_FILE\"))\n+                .build());\n+            strategy += \" -\u003e Kubernetes workload identity\";\n+        }\n+\n+        return new BuiltCredential(chain.build(), strategy, false);\n+    }\n+\n+    private boolean hasAzurePipelinesServiceConnectionConfiguration() {\n+        return has(\"TF_BUILD\")\n+            \u0026\u0026 has(\"SYSTEM_OIDCREQUESTURI\")\n+            \u0026\u0026 has(\"SYSTEM_ACCESSTOKEN\")\n+            \u0026\u0026 has(\"AZURE_SERVICE_CONNECTION_ID\")\n+            \u0026\u0026 has(\"AZURE_TENANT_ID\")\n+            \u0026\u0026 has(\"AZURE_CLIENT_ID\");\n+    }\n+\n+    private boolean hasWorkloadIdentityConfiguration() {\n+        return has(\"AZURE_FEDERATED_TOKEN_FILE\")\n+            \u0026\u0026 has(\"AZURE_TENANT_ID\")\n+            \u0026\u0026 has(\"AZURE_CLIENT_ID\");\n+    }\n+\n+    private boolean has(String name) {\n+        return isPresent(environment.get(name));\n+    }\n+\n+    private String required(String name) {\n+        String value = environment.get(name);\n+        if (!isPresent(value)) {\n+            throw new IllegalStateException(\"Required environment variable is missing: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static boolean isPresent(String value) {\n+        return value != null \u0026\u0026 !value.isBlank();\n+    }\n+\n+    public record BuiltCredential(TokenCredential credential, String strategy, boolean caeEnabled) {\n+    }\n+}\n*** Add File: src/main/java/com/example/azure/identity/AuthenticationFailureAnalyzer.java\n+package com.example.azure.identity;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Locale;\n+\n+final class AuthenticationFailureAnalyzer {\n+    private AuthenticationFailureAnalyzer() {\n+    }\n+\n+    static String describe(Throwable failure) {\n+        Throwable root = rootCause(failure);\n+        String details = collectMessages(failure);\n+        String normalized = details.toLowerCase(Locale.ROOT);\n+\n+        String category;\n+        if (containsAny(normalized, \"aadsts7000222\", \"expired client secret\", \"certificate has expired\",\n+            \"expired certificate\", \"certificate is not within its validity period\")) {\n+            category = \"The client secret or certificate has expired\";\n+        } else if (containsAny(normalized, \"aadsts500011\", \"aadsts90002\", \"tenant not found\",\n+            \"invalid tenant\", \"wrong tenant\")) {\n+            category = \"The configured tenant is invalid or does not contain the requested application/resource\";\n+        } else if (containsAny(normalized, \"aadsts7000215\", \"invalid client secret\", \"invalid_client\")) {\n+            category = \"The client credential is invalid\";\n+        } else if (containsAny(normalized, \"aadsts700016\", \"application with identifier\")) {\n+            category = \"The configured client ID was not found in the tenant\";\n+        } else if (containsAny(normalized, \"aadsts700211\", \"federated identity record\")) {\n+            category = \"No matching workload identity federation record was found\";\n+        } else if (containsAny(normalized, \"aadsts700024\", \"client assertion is not within\")) {\n+            category = \"The workload identity assertion is expired or not yet valid\";\n+        } else if (containsAny(normalized, \"unauthorized_client\", \"consent\", \"permission\", \"forbidden\")) {\n+            category = \"The identity lacks consent or permission for this token request\";\n+        } else if (failure instanceof CredentialUnavailableException\n+            || containsAny(normalized, \"credentialunavailableexception\", \"authentication unavailable\",\n+                \"no accounts were found\", \"cannot be established\", \"not installed\", \"not logged in\")) {\n+            category = \"No configured credential source or Azure identity is available\";\n+        } else if (failure instanceof ClientAuthenticationException) {\n+            category = \"Microsoft Entra ID rejected the authentication request\";\n+        } else if (containsAny(normalized, \"timeout\", \"timed out\", \"connection refused\", \"unknownhost\")) {\n+            category = \"The identity endpoint could not be reached\";\n+        } else {\n+            category = root.getClass().getSimpleName();\n+        }\n+\n+        return category + \": \" + sanitize(root.getMessage());\n+    }\n+\n+    private static Throwable rootCause(Throwable failure) {\n+        Throwable current = failure;\n+        while (current.getCause() != null \u0026\u0026 current.getCause() != current) {\n+            current = current.getCause();\n+        }\n+        return current;\n+    }\n+\n+    private static String collectMessages(Throwable failure) {\n+        List\u003cString\u003e messages = new ArrayList\u003c\u003e();\n+        Throwable current = failure;\n+        while (current != null) {\n+            if (current.getMessage() != null) {\n+                messages.add(current.getMessage());\n+            }\n+            current = current.getCause();\n+        }\n+        return String.join(\" | \", messages);\n+    }\n+\n+    private static boolean containsAny(String value, String... candidates) {\n+        for (String candidate : candidates) {\n+            if (value.contains(candidate)) {\n+                return true;\n+            }\n+        }\n+        return false;\n+    }\n+\n+    private static String sanitize(String message) {\n+        if (message == null || message.isBlank()) {\n+            return \"No additional details were supplied.\";\n+        }\n+        return message.replaceAll(\"[\\\\r\\\\n]+\", \" \").trim();\n+    }\n+}\n*** Add File: src/main/java/com/example/azure/identity/CredentialConnectivityTester.java\n+package com.example.azure.identity;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenRequestContext;\n+\n+import java.time.Duration;\n+\n+public final class CredentialConnectivityTester {\n+    private static final Duration TOKEN_REQUEST_TIMEOUT = Duration.ofSeconds(30);\n+\n+    public boolean test(CredentialFactory.BuiltCredential builtCredential, String scope) {\n+        TokenRequestContext request = new TokenRequestContext()\n+            .addScopes(scope)\n+            .setCaeEnabled(builtCredential.caeEnabled());\n+\n+        System.out.println(\"[sync] Requesting a token...\");\n+        try {\n+            AccessToken token = builtCredential.credential().getToken(request).block(TOKEN_REQUEST_TIMEOUT);\n+            if (token == null) {\n+                throw new IllegalStateException(\"The credential completed without returning a token.\");\n+            }\n+            printSuccess(\"sync\", token, builtCredential.caeEnabled());\n+            return true;\n+        } catch (RuntimeException failure) {\n+            printFailure(\"sync\", failure, builtCredential.caeEnabled());\n+            return false;\n+        }\n+    }\n+\n+    static void printSuccess(String mode, AccessToken token, boolean caeEnabled) {\n+        System.out.printf(\"[%s] SUCCESS%n\", mode);\n+        System.out.printf(\"[%s] Token expires: %s%n\", mode, token.getExpiresAt());\n+        printCaeStatus(mode, caeEnabled);\n+    }\n+\n+    static void printFailure(String mode, Throwable failure, boolean caeEnabled) {\n+        System.out.printf(\"[%s] FAILURE%n\", mode);\n+        System.out.printf(\"[%s] Reason: %s%n\", mode, AuthenticationFailureAnalyzer.describe(failure));\n+        printCaeStatus(mode, caeEnabled);\n+    }\n+\n+    private static void printCaeStatus(String mode, boolean caeEnabled) {\n+        System.out.printf(\n+            \"[%s] CAE-enabled token request: %s\"\n+                + \" (CAE issuance is controlled by Microsoft Entra ID and the target resource)%n\",\n+            mode,\n+            caeEnabled ? \"yes\" : \"no\"\n+        );\n+    }\n+}\n*** Add File: src/main/java/com/example/azure/identity/AsyncCredentialConnectivityTester.java\n+package com.example.azure.identity;\n+\n+import com.azure.core.credential.TokenRequestContext;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncCredentialConnectivityTester {\n+    public Mono\u003cBoolean\u003e test(CredentialFactory.BuiltCredential builtCredential, String scope) {\n+        TokenRequestContext request = new TokenRequestContext()\n+            .addScopes(scope)\n+            .setCaeEnabled(builtCredential.caeEnabled());\n+\n+        System.out.println(\"[async] Requesting a token...\");\n+        return builtCredential.credential()\n+            .getToken(request)\n+            .doOnNext(token -\u003e CredentialConnectivityTester.printSuccess(\n+                \"async\", token, builtCredential.caeEnabled()))\n+            .map(token -\u003e true)\n+            .onErrorResume(failure -\u003e {\n+                CredentialConnectivityTester.printFailure(\"async\", failure, builtCredential.caeEnabled());\n+                return Mono.just(false);\n+            });\n+    }\n+}\n*** Add File: src/main/java/com/example/azure/identity/Main.java\n+package com.example.azure.identity;\n+\n+public final class Main {\n+    private static final String AZURE_RESOURCE_MANAGER_SCOPE\n+        = \"https://management.azure.com/.default\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        DeploymentEnvironment environment = new EnvironmentDetector().detect();\n+        boolean caeEnabled = Boolean.parseBoolean(\n+            System.getenv().getOrDefault(\"AZURE_ENABLE_CAE\", \"false\"));\n+        CredentialFactory.BuiltCredential credential\n+            = new CredentialFactory().create(environment, caeEnabled);\n+\n+        System.out.println(\"Detected environment: \" + environment);\n+        System.out.println(\"Credential strategy: \" + credential.strategy());\n+        System.out.println(\"Azure scope: \" + AZURE_RESOURCE_MANAGER_SCOPE);\n+        System.out.println();\n+\n+        new CredentialConnectivityTester().test(credential, AZURE_RESOURCE_MANAGER_SCOPE);\n+        System.out.println();\n+        new AsyncCredentialConnectivityTester()\n+            .test(credential, AZURE_RESOURCE_MANAGER_SCOPE)\n+            .block();\n+    }\n+}\n*** Add File: README.md\n+# Azure credential chain demo\n+\n+Small Java 17 sample that selects an explicit Azure Identity credential chain for local development,\n+CI/CD, or production and tests it against the Azure Resource Manager scope. It only requests tokens;\n+it does not create or modify Azure resources.\n+\n+## Credential strategies\n+\n+| Environment | Detection | Credential order |\n+| --- | --- | --- |\n+| Development | Default when no hosted-environment marker is present | Azure CLI, IntelliJ, Visual Studio Code, Azure Developer CLI, Azure PowerShell |\n+| CI/CD | `CI`, `TF_BUILD`, `BUILD_BUILDID`, `PIPELINE_WORKSPACE`, or another supported CI marker | Azure Pipelines workload identity service connection when fully configured, then `EnvironmentCredential` |\n+| Production | A managed identity endpoint variable or complete workload identity configuration | System- or user-assigned managed identity, then Kubernetes workload identity when configured |\n+\n+Set `APP_ENVIRONMENT` to `dev`, `ci`, or `production` to override auto-detection.\n+\n+## Configuration\n+\n+- Pipeline service principal: `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and either\n+  `AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`.\n+- Azure Pipelines workload identity service connection: `TF_BUILD`, `SYSTEM_OIDCREQUESTURI`,\n+  `SYSTEM_ACCESSTOKEN`, `AZURE_SERVICE_CONNECTION_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_ID`.\n+  Map the pipeline's `System.AccessToken` secret into `SYSTEM_ACCESSTOKEN`.\n+- User-assigned managed identity: `AZURE_MANAGED_IDENTITY_CLIENT_ID`.\n+  Leave it unset to use the system-assigned identity.\n+- Kubernetes workload identity fallback: `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and\n+  `AZURE_FEDERATED_TOKEN_FILE`.\n+- CAE: set `AZURE_ENABLE_CAE=true`. CAE is requested through `TokenRequestContext`; Microsoft Entra ID\n+  and the target resource decide whether a CAE token is issued. Azure Identity developer credentials\n+  don't support CAE, so normally leave CAE disabled for local development.\n+\n+Never commit secret values or pipeline access tokens.\n+\n+## Build and run\n+\n+```text\n+mvn clean package\n+mvn exec:java\n+```\n+\n+The program performs a synchronous token request and then an asynchronous request for\n+`https://management.azure.com/.default`. It prints no token contents.\n+\n+## References\n+\n+- [Credential chains in Azure Identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)\n+- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)\n+- [Continuous Access Evaluation](https://learn.microsoft.com/entra/identity/conditional-access/concept-continuous-access-evaluation)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml
@@ -1,0 +1,49 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example.azure</groupId>
+    <artifactId>credential-chain-demo</artifactId>
+    <version>1.0.0</version>
+    <name>Azure credential chain demo</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.identity.version>1.18.5</azure.identity.version>
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
+            <artifactId>azure-identity-broker</artifactId>
+            <version>${azure.identity.version}</version>
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
+                <artifact
... (truncated)
```

### ✅ 17. powershell — 4759ms

**Input:**

```json
{"command":"java -version; mvn -version; mvn -q -DskipTests package","description":"Compile the Java Maven project","initial_wait":180,"mode":"sync"}
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
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[ERROR] Failed to execute goal on project credential-chain-demo: Could not resolve dependencies for project com.example.azure:credential-chain-demo:jar:1.0.0
[ERROR] dependency: com.azure:azure-identity-broker:jar:1.18.5 (compile)
[ERROR] 	Could not find artifact com.azure:azure-identity-broker:jar:1.18.5 in central (https://repo.maven.apache.org/maven2)
[ERROR] 
[ERROR] -> [Help 1]
[ERROR] 
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR] 
[ERROR] For more information about the errors and possible solutions, please read the following articles:
[ERROR] [Help 1] http://cwiki.apache.org/confluence/display/MAVEN/DependencyResolutionException
<shellId: 0 completed with exit code 1>
```

### ✅ 18. web_fetch — 1895ms

**Input:**

```json
{"max_length":6000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity-broker/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity-broker/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-identity-broker</artifactId>
  <versioning>
    <latest>1.1.22</latest>
    <release>1.1.22</release>
    <versions>
      <version>1.0.0-beta.1</version>
      <version>1.0.0</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.1.2</version>
      <version>1.1.3</version>
      <version>1.1.4</version>
      <version>1.1.5</version>
      <version>1.1.6</version>
      <version>1.1.7</version>
      <version>1.1.8</version>
      <version>1.1.9</version>
      <version>1.1.10</version>
      <version>1.1.11</version>
      <version>1.1.12</version>
      <version>1.1.13</version>
      <version>1.1.14</version>
      <version>1.1.15</version>
      <version>1.1.16</version>
      <version>1.1.17</version>
      <version>1.1.18</version>
      <version>1.1.19</version>
      <version>1.1.20</version>
      <version>1.1.21</version>
      <version>1.1.22</version>
    </versions>
    <lastUpdated>20260824185322</lastUpdated>
  </versioning>
</metadata>

```

### ✅ 19. web_fetch — 1559ms

**Input:**

```json
{"max_length":8000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/1.18.5/azure-identity-1.18.5.pom"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/1.18.5/azure-identity-1.18.5.pom:
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <version>1.18.5</version><!-- {x-version-update;com.azure:azure-identity;current} -->

  <name>Microsoft Azure client library for Identity</name>
  <description>This module contains client library for Microsoft Azure Identity.</description>
  <url>https://github.com/Azure/azure-sdk-for-java</url>

  <properties>
    <jacoco.min.branchcoverage>0.35</jacoco.min.branchcoverage>
    <!-- Configures the Java 9+ run to perform the required module exports, opens, and reads that are necessary for testing but shouldn't be part of the module-info. -->
    <javaModulesSurefireArgLine>
      --add-opens java.xml/jdk.xml.internal=ALL-UNNAMED
      --add-exports com.azure.core/com.azure.core.implementation.util=ALL-UNNAMED
    </javaModulesSurefireArgLine>

    <spotbugs.skip>false</spotbugs.skip>
    <spotbugs.includeTests>true</spotbugs.includeTests>
  </properties>

  <parent>
    <groupId>com.azure</groupId>
    <artifactId>azure-client-sdk-parent</artifactId>
    <version>1.7.0</version> <!-- {x-version-update;com.azure:azure-client-sdk-parent;current} -->
    <relativePath>../../parents/azure-client-sdk-parent</relativePath>
  </parent>

  <dependencies>
    <dependency>
      <groupId>com.azure</groupId>
      <artifactId>azure-core</artifactId>
      <version>1.59.0</version> <!-- {x-version-update;com.azure:azure-core;dependency} -->
    </dependency>

    <dependency>
      <groupId>com.azure</groupId>
      <artifactId>az
... (truncated)
```

### ✅ 20. apply_patch — 688ms

**Input:**

```json
"*** Begin Patch\n*** Update File: pom.xml\n@@\n         \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n         \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+        \u003cazure.identity.broker.version\u003e1.1.22\u003c/azure.identity.broker.version\u003e\n@@\n             \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n             \u003cartifactId\u003eazure-identity-broker\u003c/artifactId\u003e\n-            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+            \u003cversion\u003e${azure.identity.broker.version}\u003c/version\u003e\n         \u003c/dependency\u003e\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594/pom.xml
@@ -13,6 +13,7 @@
         <maven.compiler.release>17</maven.compiler.release>
         <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
         <azure.identity.version>1.18.5</azure.identity.version>
+        <azure.identity.broker.version>1.1.22</azure.identity.broker.version>
     </properties>
 
     <dependencies>
@@ -24,7 +25,7 @@
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-identity-broker</artifactId>
-            <version>${azure.identity.version}</version>
+            <version>${azure.identity.broker.version}</version>
         </dependency>
     </dependencies>
 

```

### ✅ 21. powershell — 6703ms

**Input:**

```json
{"command":"mvn -q -DskipTests package","description":"Compile the corrected Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 1 completed with exit code 0>
```

### ✅ 22. powershell — 3597ms

**Input:**

```json
{"command":"mvn -q dependency:tree -Dincludes=com.azure:azure-identity,com.azure:azure-identity-broker; Get-ChildItem -Recurse -File target | Select-Object -ExpandProperty FullName","description":"Inspect resolved Azure dependencies and build output","initial_wait":120,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[ERROR] Error resolving version for plugin '.azure:azure-identity,com.azure' from the repositories [local (C:\Users\weidxu\.m2\repository), central (https://repo.maven.apache.org/maven2)]: Plugin not found in any plugin repository -> [Help 1]
[ERROR] 
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR] 
[ERROR] For more information about the errors and possible solutions, please read the following articles:
[ERROR] [Help 1] http://cwiki.apache.org/confluence/display/MAVEN/PluginVersionResolutionException
C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594\target\credential-chain-demo-1.0.0.jar
C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594\target\classes\com\example\azure\identity\AsyncCredentialConnectivityTester.class
C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594\target\classes\com\example\azure\identity\AuthenticationFailureAnalyzer.class
C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594\target\classes\com\example\azure\identity\CaeTokenCredential.class
C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2281420594\target\classes\com\example\azure\identity\CredentialConnectivityTester.c
... (truncated)
```

### ✅ 23. powershell — 4233ms

**Input:**

```json
{"command":"mvn dependency:tree \"-Dincludes=com.azure:azure-identity,com.azure:azure-identity-broker\" -DoutputType=text","description":"Verify resolved Azure Identity versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example.azure:credential-chain-demo >---------------
[INFO] Building Azure credential chain demo 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- dependency:3.7.0:tree (default-cli) @ credential-chain-demo ---
[INFO] com.example.azure:credential-chain-demo:jar:1.0.0
[INFO] +- com.azure:azure-identity:jar:1.18.5:compile
[INFO] \- com.azure:azure-identity-broker:jar:1.1.22:compile
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  1.751 s
[INFO] Finished at: 2026-08-29T08:55:39+08:00
[INFO] ------------------------------------------------------------------------
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\azure\identity\AsyncCredentialConnectivityTester.java`
- `src\main\java\com\example\azure\identity\AuthenticationFailureAnalyzer.java`
- `src\main\java\com\example\azure\identity\CaeTokenCredential.java`
- `src\main\java\com\example\azure\identity\CredentialConnectivityTester.java`
- `src\main\java\com\example\azure\identity\CredentialFactory.java`
- `src\main\java\com\example\azure\identity\DeploymentEnvironment.java`
- `src\main\java\com\example\azure\identity\EnvironmentDetector.java`
- `src\main\java\com\example\azure\identity\Main.java`

## Copilot Response

I’ve confirmed the current Azure Identity Java APIs and CAE’s per-request behavior. I’m creating the tailored chains with explicit environment-specific ordering and diagnostic failure classification.

Created the Java 17 Maven project with:

- Environment detection for development, CI/CD, and production
- Tailored credential chains for developer tools, pipeline credentials, managed identity, and Kubernetes workload identity
- Optional CAE via `AZURE_ENABLE_CAE`
- Synchronous and asynchronous ARM token connectivity tests
- Detailed authentication failure classification without exposing tokens
- Complete configuration and usage guidance in `README.md`

Run with:

```powershell
mvn clean package
mvn exec:java
```

References: [Azure Identity credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains) and [Continuous Access Evaluation](https://learn.microsoft.com/entra/identity/conditional-access/concept-continuous-access-evaluation).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK clients are created or left unclosed.

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
hyoka run --prompt-id identity-dp-java-credential-chain --config java-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
