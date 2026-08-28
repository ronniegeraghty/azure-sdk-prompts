# Evaluation Report: identity-dp-java-credential-chain

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 865.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-credential-chain` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 24/26 |
| Duration | 865.4s |
| Timestamp | 2026-08-27T15:40:17Z |
| Files Generated | 14 |
| Event Count | 12743 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 604.6s |
| Review | 255.4s |
| **Total** | **865.4s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | entra-app-registration |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=436481 out=9834 |
| Turn Count | 9 |

## Error

```
evaluation failed: sending prompt: waiting for session.idle: context deadline exceeded
```

**Details:**

```
sending prompt: waiting for session.idle: context deadline exceeded
```

## Tool Calls

### ✅ 1. skill — 3678ms

**Input:**

```json
{"skill":"entra-app-registration"}
```

**Output:**

```
Skill loaded successfully ✅

## Overview

Microsoft Entra ID (formerly Azure Active Directory) is Microsoft's cloud-based identity and access management service. App registrations allow applications to authenticate users and access Azure resources securely.

### Key Concepts

| Concept | Description |
|---------|-------------|
| **App Registration** | Configuration that allows an app to use Microsoft identity platform |
| **Application (Client) ID** | Unique identifier for your application |
| **Tenant ID** | Unique identifier for your Azure AD tenant/directory |
| **Client Secret** | Password for the application (confidential clients only) |
| **Redirect URI** | URL where authentication responses are sent |
| **API Permissions** | Access scopes your app requests |
| **Service Principal** | Identity created in your tenant when you register an app |

### Application Types

| Type | Use Case |
|------|----------|
| **Web Application** | Server-side apps, APIs |
| **Single Page App (SPA)** | JavaScript/React/Angular apps |
| **Mobile/Native App** | Desktop, mobile apps |
| **Daemon/Service** | Background services, APIs |

## Core Workflow

### Step 1: Register the Application

Create an app registration in the Azure portal or using Azure CLI.

**Portal Method:**
1. Navigate to Azure Portal → Microsoft Entra ID → App registrations
2. Click "New registration"
3. Provide name, supported account types, and redirect URI
4. Click "Register"

**CLI Method:** See [references/cli-commands.md](references/cli-commands.md)
**IaC Method:** See [references/BICEP-EXAMPLE.bicep](references/BICEP-EXAMPLE.bicep)

It's highly recommended to use the IaC to manage Entra app registration if you already use IaC in your project, need a scalable solution for managing lots of app registrations or need fine-grained audit history of the configuration changes. 

### Step 2: Configure Authentication

Set up authentication settings based on your application type.

- **Web Apps**: Add redirect 
... (truncated)
```

### ✅ 2. view — 187999ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\sdk\\azure-identity-java.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md
@@ -1,31 +1,31 @@
 # Authentication — Java SDK Quick Reference
 
 > Condensed from **azure-identity-java**. Full patterns (workload identity,
 > certificate auth, device code, sovereign clouds)
 > in the **azure-identity-java** plugin skill if installed.
 
 ## Install
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.15.0</version>
 </dependency>
 ```
 
 ## Quick Start
 > **Auth:** `DefaultAzureCredential` is for local development. See [auth-best-practices.md](../auth-best-practices.md) for production patterns.
 
 ```java
 import com.azure.identity.DefaultAzureCredentialBuilder;
 var credential = new DefaultAzureCredentialBuilder().build();
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** (CLI, PowerShell, VS Code). In production, use ManagedIdentityCredential — see [auth-best-practices.md](../auth-best-practices.md)
 - Managed identity in production — no secrets to manage, automatic rotation
 - Azure CLI for local dev — run `az login` before running your app
 - Least privilege — grant only required permissions to service principals
 - Token caching — enabled by default, reduces auth round-trips
 - Environment variables — use for CI/CD, not hardcoded se
... (truncated)
```

### ✅ 3. view — 186460ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks enviro
... (truncated)
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 188001ms

**Input:**

```json
{"intent":"Create a local-only Java 17 Maven sample demonstrating environment-specific Azure Identity credential chains, CAE token requests, environment detection, and sync/async token connectivity testing. Do not deploy or modify Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. glob — 184765ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. web_search (via github-mcp-server) — 185156ms

**Input:**

```json
{"query":"Current Azure Identity Java SDK APIs ChainedTokenCredentialBuilder AzurePipelinesCredentialBuilder WorkloadIdentityCredentialBuilder ManagedIdentityCredentialBuilder CAE TokenRequestContext setCaeEnabled Maven version 2026 official Microsoft documentation"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here is a concise overview with references to the official Microsoft documentation for the Azure Identity Java SDK APIs you mentioned, including Maven usage, CAE (Continuous Access Evaluation), and feature support as of the 2026 releases:\n\n### Current Azure Identity Java SDK APIs\n\n- **ChainedTokenCredentialBuilder**\n  - Used to combine multiple credential types into a fallback chain. If one credential fails to authenticate, the next is tried until one succeeds.\n  - Example: You can combine ManagedIdentity, InteractiveBrowser, and other credentials for robust authentication across different environments (e.g., local dev, cloud automation)【3:1†source】【3:7†source】.\n  - [API Reference](https://learn.microsoft.com/en-us/java/api/com.azure.identity.chainedtokencredentialbuilder)【3:1†source】.\n\n- **AzurePipelinesCredentialBuilder**\n  - A credential optimized for Azure DevOps Pipelines, enabling seamless authentication (information on explicit builder available in later SDKs or preview; consult latest SDK docs for specifics as references may be evolving).\n\n- **WorkloadIdentityCredentialBuilder**\n  - Supports Azure AD workload identity federation, especially for cloud-native scenarios, enabling applications to use federated tokens for access instead of client secrets or certificates.\n\n- **ManagedIdentityCredentialBuilder**\n  - Used to obtain tokens from Azure managed identities—system-assigned or user-assigned—enabling keyless authentication within supported Azure services (VMs, App Services, AKS, etc.)【3:1†source】.\n  - [API Reference](https://learn.microsoft.com/en-us/java/api/com.azure.identity.managedidentitycredentialbuilder)\n\n### CAE (Continuous Access Evaluation)\n- Starting from SDK version 1.8.0 (and improvements in newer versions), CAE is supported. You can enable CAE for token requests:\n- **TokenRequestContext#setCaeEnabled**—this flag allows service clients to opt-in for CAE-en
... (truncated)
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 185160ms

**Input:**

```json
{"intent":"Provide Java Azure Identity best practices for a local-only Maven sample with environment-specific credential chains and CAE token requests. No resource deployment or Azure mutations.","learn":false}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. github-mcp-server-search_code (via github-mcp-server) — 4994ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":5,"query":"class AzurePipelinesCredentialBuilder repo:Azure/azure-sdk-for-java"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"README.md","path":"sdk/identity/azure-identity/README.md","sha":"1fe34673fee8854ba20262a24a671b9f38f9c1e2","text_matches":[{"fragment":"\n## Continuous Access Evaluation\n\nAs of v1.10.0, accessing resources protected by [Continuous Access Evaluation](https://learn.microsoft.com/entra/identity/conditional-access/concept-continuous-access-evaluation) (CAE) is possible on a per-request basis. This can be enabled using the [`TokenRequestContext.setCaeEnabled(boolean)` API](https://learn.microsoft.com/java/api/com.azure.core.credential.tokenrequestcontext?view=azure-java-stable#com-azure-core-credential-tokenrequestcontext-setcaeenabled(boolean)). CAE isn't supported for developer credentials.\n\n## Token caching\n","matches":[{"indices":[307,320],"text":"setCaeEnabled"},{"indices":[488,501],"text":"setcaeenabled"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/identity/azure-identity/README.md?ref=adb5c6c0d3ccc0c2ca17af296e814a24b404de43","property":"content"}]},{"name":"DeviceCodeCredentialTest.java","path":"sdk/identity/azure-identity/src/test/java/com/azure/identity/DeviceCodeCredentialTest.java","sha":"d9b75da1ccc257e32da3536306b1df57095352ed","text_matches":[{"fragment":"        TokenRequestContext request1\n            = new TokenRequestContext().addScopes(\"https://management.azure.com\").setCaeEnabled(true);\n        TokenRequestContext request2\n            = new TokenRequestContext().addScopes(\"https://vault.azure.net\").setCaeEnabled(true);\n        OffsetDateTime expiresAt = OffsetDateTime.now(ZoneOffset.UTC).plusHours(1);\n\n        // mock","matches":[{"indices":[119,132],"text":"setCaeEnabled"},{"indices":[254,267],"text":"setCaeEnabled"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/identity/azure-identity/src/test/java/com/azure/identity/DeviceCodeCredentialTest.java?ref=adb5c6c0d3ccc0c2ca17af296e814a
... (truncated)
```

### ✅ 9. github-mcp-server-search_code (via github-mcp-server) — 4654ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":5,"query":"class AzurePipelinesCredential repo:Azure/azure-sdk-for-java"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"AzurePipelinesCredential.java","path":"sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredential.java","sha":"026358ae1ccc39640b97b7844fe64f9441d44f8e","text_matches":[{"fragment":" *\n */\n@Immutable\npublic class AzurePipelinesCredential implements TokenCredential {\n    private static final ClientLogger LOGGER = new ClientLogger(AzurePipelinesCredential.class);\n    private final IdentityClient identityClient;\n    private final IdentitySyncClient identitySyncClient;","matches":[{"indices":[25,30],"text":"class"},{"indices":[31,55],"text":"AzurePipelinesCredential"},{"indices":[149,173],"text":"AzurePipelinesCredential"},{"indices":[174,179],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredential.java?ref=adb5c6c0d3ccc0c2ca17af296e814a24b404de43","property":"content"}]},{"name":"AzurePipelinesCredential.java","path":"sdk/identity-v2/azure-identity/src/main/java/com/azure/v2/identity/AzurePipelinesCredential.java","sha":"1bfe823a62bc2501422ada009172eec78e167777","text_matches":[{"fragment":" *     .build&#40;&#41;;\n * </pre>\n */\npublic class AzurePipelinesCredential implements TokenCredential {\n    private static final ClientLogger LOGGER = new ClientLogger(AzurePipelinesCredential.class);\n    private final ConfidentialClient confidentialClient;\n","matches":[{"indices":[46,51],"text":"class"},{"indices":[52,76],"text":"AzurePipelinesCredential"},{"indices":[170,194],"text":"AzurePipelinesCredential"},{"indices":[195,200],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/identity-v2/azure-identity/src/main/java/com/azure/v2/identity/AzurePipelinesCredential.java?ref=adb5c6c0d3ccc0c2ca17af296e814a24b404de43","property":"content"}]},{"name":"AzurePipelinesCredentialBuilder.java","path
... (truncated)
```

### ✅ 10. github-mcp-server-search_code (via github-mcp-server) — 4298ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"setCaeEnabled repo:Azure/azure-sdk-for-java path:sdk/identity/azure-identity"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"AzurePipelinesCredentialBuilder.java","path":"sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredentialBuilder.java","sha":"68387551ddec9e8a65f635117223df2c1c148f7e","text_matches":[{"fragment":" * </pre>\n * <!-- end com.azure.identity.credential.azurepipelinescredential.construct -->\n */\npublic class AzurePipelinesCredentialBuilder extends AadCredentialBuilderBase<AzurePipelinesCredentialBuilder> {\n    private static final ClientLogger LOGGER = new ClientLogger(AzurePipelinesCredentialBuilder.class);\n    private static final String OIDC_API_VERSION = \"7.1\";\n    private String serviceConnectionId;","matches":[{"indices":[102,107],"text":"class"},{"indices":[108,139],"text":"AzurePipelinesCredentialBuilder"},{"indices":[173,204],"text":"AzurePipelinesCredentialBuilder"},{"indices":[272,303],"text":"AzurePipelinesCredentialBuilder"},{"indices":[304,309],"text":"class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredentialBuilder.java?ref=adb5c6c0d3ccc0c2ca17af296e814a24b404de43","property":"content"}]},{"name":"AzurePipelinesCredentialBuilder.java","path":"sdk/identity-v2/azure-identity/src/main/java/com/azure/v2/identity/AzurePipelinesCredentialBuilder.java","sha":"6a6683c3ded9e1a4485b8275498c7cbcbc999fb7","text_matches":[{"fragment":" *     .build&#40;&#41;;\n * </pre>\n */\npublic class AzurePipelinesCredentialBuilder extends EntraIdCredentialBuilderBase<AzurePipelinesCredentialBuilder> {\n    private static final ClientLogger LOGGER = new ClientLogger(AzurePipelinesCredentialBuilder.class);\n    private static final String OIDC_API_VERSION = \"7.1\";\n    private String serviceConnectionId;","matches":[{"indices":[46,51],"text":"class"},{"indices":[52,83],"text":"AzurePipelinesCredentialBuilder"},{"indices":[121,152],"text":"AzurePipelinesCredentialBuilder"}
... (truncated)
```

### ✅ 11. github-mcp-server-get_file_contents (via github-mcp-server) — 10253ms

**Input:**

```json
{"owner":"Azure","path":"sdk/identity/azure-identity/src/main/java/com/azure/identity/AzurePipelinesCredentialBuilder.java","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 43032297e2c8c008d020e29f64b1094bddd1635f)<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <version>1.19.0-beta.3</version><!-- {x-version-update;com.azure:azure-identity;current} -->

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
      <artifactId>azure-core-http-netty</artifactId>
      <version>1.16.6</version> <!-- {x-version-update;com.azu
... (truncated)
```

### ✅ 12. github-mcp-server-get_file_contents (via github-mcp-server) — 9874ms

**Input:**

```json
{"owner":"Azure","path":"sdk/identity/azure-identity/src/main/java/com/azure/identity/ChainedTokenCredentialBuilder.java","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: d036657ee0bad52b4e6801775ab74e1e2483cad5)// Copyright (c) Microsoft Corporation. All rights reserved.
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
 * The WorkloadIdentityCredential supports Azure workload identity authentication on Azure Kubernetes and acquires
 * a token using the service account credentials available in the Azure Kubernetes environment.
 * Refer to <a href="https://learn.microsoft.com/azure/aks/work
... (truncated)
```

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server) — 9529ms

**Input:**

```json
{"owner":"Azure","path":"sdk/identity/azure-identity/src/main/java/com/azure/identity/WorkloadIdentityCredentialBuilder.java","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: a19c7effdd379b73da8e0d780c9d1b8ff9f7d2d3)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package com.azure.identity;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.logging.ClientLogger;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Deque;

/**
 * <p>Fluent credential builder for instantiating {@link ChainedTokenCredential}.</p>
 *
 * <p><strong>Sample: Construct a ChainedTokenCredential.</strong></p>
 *
 * <p>The following code sample demonstrates the creation of a {@link com.azure.identity.ChainedTokenCredential},
 * using the {@link com.azure.identity.ChainedTokenCredentialBuilder} to configure it. The sample below
 * tries managed identity first, then interactive browser login as needed
 * (e.g. when 2FA is turned on in the directory). Once this credential is created, it may be passed into the builder
 * of many of the Azure SDK for Java client builders as the 'credential' parameter.</p>
 *
 * <!-- src_embed com.azure.identity.credential.chainedtokencredential.construct -->
 * <pre>
 * TokenCredential managedIdentityCredential = new ManagedIdentityCredentialBuilder&#40;&#41;.build&#40;&#41;;
 * TokenCredential interactiveBrowserCredential = new InteractiveBrowserCredentialBuilder&#40;&#41;.clientId&#40;clientId&#41;
 *     .redirectUrl&#40;&quot;https:&#47;&#47;localhost:8765&quot;&#41;
 *     .build&#40;&#41;;
 * TokenCredential credential = new ChainedTokenCredentialBuilder&#40;&#41;.addLast&#40;managedIdentityCredential&#41;
 *     .addLast&#40;interactiveBrowserCredential&#41;
 *     .build&#40;&#41;;
 * </pre>
 * <!-- end com.azure.identity.credential.chainedtokencredential.construct -->
 *
 * @see ChainedTokenCredential
 */
public class ChainedTokenCredentialBuilder {
    private final Deque<TokenCredential> credentials;
    private static final ClientLogger LOGGER = new ClientLogge
... (truncated)
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server) — 9179ms

**Input:**

```json
{"owner":"Azure","path":"sdk/identity/azure-identity/pom.xml","repo":"azure-sdk-for-java"}
```

**Output:**

```
successfully downloaded text file (SHA: 68387551ddec9e8a65f635117223df2c1c148f7e)// Copyright (c) Microsoft Corporation. All rights reserved.
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
    public AzurePipelinesCredentialBuilder() {
        super();
    }

    /**
     * Sets the service connection id for the Azure Pipelines service connection. The service connection ID is
     * retrieved from the Service Con
... (truncated)
```

### ✅ 15. apply_patch — 937ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-credential-chains\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure credential chains by environment\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.5\u003c/azure.identity.version\u003e\n+        \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.azureidentity.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/azureidentity/DeploymentEnvironment.java\n+package com.example.azureidentity;\n+\n+public enum DeploymentEnvironment {\n+    DEV,\n+    CI,\n+    PRODUCTION\n+}\n*** Add File: src/main/java/com/example/azureidentity/EnvironmentDetector.java\n+package com.example.azureidentity;\n+\n+import java.util.Map;\n+import java.util.stream.Stream;\n+\n+public final class EnvironmentDetector {\n+    private static final String[] CI_MARKERS = {\n+        \"CI\",\n+        \"TF_BUILD\",\n+        \"BUILD_BUILDID\",\n+        \"BUILD_SOURCESDIRECTORY\",\n+        \"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI\",\n+        \"GITHUB_ACTIONS\",\n+        \"GITLAB_CI\",\n+        \"JENKINS_URL\"\n+    };\n+\n+    private static final String[] PRODUCTION_MARKERS = {\n+        \"IDENTITY_ENDPOINT\",\n+        \"MSI_ENDPOINT\",\n+        \"IMDS_ENDPOINT\",\n+        \"AZURE_FEDERATED_TOKEN_FILE\",\n+        \"WEBSITE_INSTANCE_ID\",\n+        \"CONTAINER_APP_NAME\",\n+        \"KUBERNETES_SERVICE_HOST\"\n+    };\n+\n+    public DeploymentEnvironment detect() {\n+        return detect(System.getenv());\n+    }\n+\n+    public DeploymentEnvironment detect(Map\u003cString, String\u003e environment) {\n+        if (hasAnyNonBlank(environment, CI_MARKERS)) {\n+            return DeploymentEnvironment.CI;\n+        }\n+        if (hasAnyNonBlank(environment, PRODUCTION_MARKERS)) {\n+            return DeploymentEnvironment.PRODUCTION;\n+        }\n+        return DeploymentEnvironment.DEV;\n+    }\n+\n+    private boolean hasAnyNonBlank(Map\u003cString, String\u003e environment, String[] names) {\n+        return Stream.of(names).map(environment::get).anyMatch(EnvironmentDetector::isNonBlank);\n+    }\n+\n+    private static boolean isNonBlank(String value) {\n+        return value != null \u0026\u0026 !value.isBlank();\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/CredentialSelection.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.TokenCredential;\n+\n+import java.util.Objects;\n+\n+public record CredentialSelection(\n+    TokenCredential credential,\n+    DeploymentEnvironment environment,\n+    String strategy,\n+    boolean caeEnabled\n+) {\n+    public CredentialSelection {\n+        Objects.requireNonNull(credential, \"credential\");\n+        Objects.requireNonNull(environment, \"environment\");\n+        Objects.requireNonNull(strategy, \"strategy\");\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/CaeAwareTokenCredential.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+final class CaeAwareTokenCredential implements TokenCredential {\n+    private final TokenCredential delegate;\n+    private final boolean caeEnabled;\n+\n+    CaeAwareTokenCredential(TokenCredential delegate, boolean caeEnabled) {\n+        this.delegate = Objects.requireNonNull(delegate, \"delegate\");\n+        this.caeEnabled = caeEnabled;\n+    }\n+\n+    @Override\n+    public Mono\u003cAccessToken\u003e getToken(TokenRequestContext request) {\n+        return delegate.getToken(applyCaeSetting(request));\n+    }\n+\n+    @Override\n+    public AccessToken getTokenSync(TokenRequestContext request) {\n+        return delegate.getTokenSync(applyCaeSetting(request));\n+    }\n+\n+    private TokenRequestContext applyCaeSetting(TokenRequestContext request) {\n+        return Objects.requireNonNull(request, \"request\").setCaeEnabled(caeEnabled);\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/AzureCredentialFactory.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.AzureCliCredentialBuilder;\n+import com.azure.identity.AzureDeveloperCliCredentialBuilder;\n+import com.azure.identity.AzurePipelinesCredentialBuilder;\n+import com.azure.identity.AzurePowerShellCredentialBuilder;\n+import com.azure.identity.ChainedTokenCredentialBuilder;\n+import com.azure.identity.EnvironmentCredentialBuilder;\n+import com.azure.identity.IntelliJCredentialBuilder;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.identity.WorkloadIdentityCredentialBuilder;\n+\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class AzureCredentialFactory {\n+    private static final String MANAGED_IDENTITY_CLIENT_ID = \"AZURE_MANAGED_IDENTITY_CLIENT_ID\";\n+    private static final String[] PIPELINES_CREDENTIAL_VARIABLES = {\n+        \"AZURE_CLIENT_ID\",\n+        \"AZURE_TENANT_ID\",\n+        \"AZURE_SERVICE_CONNECTION_ID\",\n+        \"SYSTEM_ACCESSTOKEN\",\n+        \"SYSTEM_OIDCREQUESTURI\"\n+    };\n+    private static final String[] WORKLOAD_IDENTITY_VARIABLES = {\n+        \"AZURE_CLIENT_ID\",\n+        \"AZURE_TENANT_ID\",\n+        \"AZURE_FEDERATED_TOKEN_FILE\"\n+    };\n+\n+    public CredentialSelection create(DeploymentEnvironment environment, boolean enableCae) {\n+        return create(environment, enableCae, System.getenv());\n+    }\n+\n+    CredentialSelection create(\n+        DeploymentEnvironment environment,\n+        boolean enableCae,\n+        Map\u003cString, String\u003e variables\n+    ) {\n+        BuiltCredential built = switch (environment) {\n+            case DEV -\u003e buildDeveloperCredential();\n+            case CI -\u003e buildCiCredential(variables);\n+            case PRODUCTION -\u003e buildProductionCredential(variables);\n+        };\n+\n+        TokenCredential credential = new CaeAwareTokenCredential(built.credential(), enableCae);\n+        return new CredentialSelection(credential, environment, built.strategy(), enableCae);\n+    }\n+\n+    private BuiltCredential buildDeveloperCredential() {\n+        TokenCredential credential = new ChainedTokenCredentialBuilder()\n+            .addLast(new AzureCliCredentialBuilder().build())\n+            .addLast(new AzureDeveloperCliCredentialBuilder().build())\n+            .addLast(new AzurePowerShellCredentialBuilder().build())\n+            .addLast(new IntelliJCredentialBuilder().build())\n+            .build();\n+\n+        return new BuiltCredential(\n+            credential,\n+            \"Developer tools: Azure CLI -\u003e Azure Developer CLI -\u003e Azure PowerShell -\u003e IntelliJ\"\n+        );\n+    }\n+\n+    private BuiltCredential buildCiCredential(Map\u003cString, String\u003e variables) {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        List\u003cString\u003e strategies = new ArrayList\u003c\u003e();\n+\n+        if (hasAny(variables, PIPELINES_CREDENTIAL_VARIABLES)) {\n+            requireAll(variables, \"Azure Pipelines workload federation\", PIPELINES_CREDENTIAL_VARIABLES);\n+            chain.addLast(new AzurePipelinesCredentialBuilder()\n+                .clientId(variables.get(\"AZURE_CLIENT_ID\"))\n+                .tenantId(variables.get(\"AZURE_TENANT_ID\"))\n+                .serviceConnectionId(variables.get(\"AZURE_SERVICE_CONNECTION_ID\"))\n+                .systemAccessToken(variables.get(\"SYSTEM_ACCESSTOKEN\"))\n+                .build());\n+            strategies.add(\"Azure Pipelines service connection\");\n+        }\n+\n+        chain.addLast(new EnvironmentCredentialBuilder().build());\n+        strategies.add(\"pipeline environment variables\");\n+\n+        return new BuiltCredential(chain.build(), String.join(\" -\u003e \", strategies));\n+    }\n+\n+    private BuiltCredential buildProductionCredential(Map\u003cString, String\u003e variables) {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        String managedIdentityClientId = variables.get(MANAGED_IDENTITY_CLIENT_ID);\n+\n+        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();\n+        String managedIdentityStrategy = \"system-assigned managed identity\";\n+        if (isNonBlank(managedIdentityClientId)) {\n+            managedIdentity.clientId(managedIdentityClientId);\n+            managedIdentityStrategy = \"user-assigned managed identity\";\n+        }\n+        chain.addLast(managedIdentity.build());\n+\n+        List\u003cString\u003e strategies = new ArrayList\u003c\u003e();\n+        strategies.add(managedIdentityStrategy);\n+        if (hasAny(variables, WORKLOAD_IDENTITY_VARIABLES)) {\n+            requireAll(variables, \"Kubernetes workload identity\", WORKLOAD_IDENTITY_VARIABLES);\n+            chain.addLast(new WorkloadIdentityCredentialBuilder()\n+                .clientId(variables.get(\"AZURE_CLIENT_ID\"))\n+                .tenantId(variables.get(\"AZURE_TENANT_ID\"))\n+                .tokenFilePath(variables.get(\"AZURE_FEDERATED_TOKEN_FILE\"))\n+                .build());\n+            strategies.add(\"Kubernetes workload identity\");\n+        }\n+\n+        return new BuiltCredential(chain.build(), String.join(\" -\u003e \", strategies));\n+    }\n+\n+    private static boolean hasAny(Map\u003cString, String\u003e variables, String[] names) {\n+        for (String name : names) {\n+            if (isNonBlank(variables.get(name))) {\n+                return true;\n+            }\n+        }\n+        return false;\n+    }\n+\n+    private static void requireAll(Map\u003cString, String\u003e variables, String label, String[] names) {\n+        List\u003cString\u003e missing = new ArrayList\u003c\u003e();\n+        for (String name : names) {\n+            if (!isNonBlank(variables.get(name))) {\n+                missing.add(name);\n+            }\n+        }\n+        if (!missing.isEmpty()) {\n+            throw new IllegalStateException(label + \" is partially configured; missing: \" + String.join(\", \", missing));\n+        }\n+    }\n+\n+    private static boolean isNonBlank(String value) {\n+        return value != null \u0026\u0026 !value.isBlank();\n+    }\n+\n+    private record BuiltCredential(TokenCredential credential, String strategy) {\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/ConnectivityResult.java\n+package com.example.azureidentity;\n+\n+import java.time.OffsetDateTime;\n+\n+public record ConnectivityResult(\n+    boolean successful,\n+    String scope,\n+    OffsetDateTime expiresAt,\n+    boolean caeEnabled,\n+    String failureReason\n+) {\n+    public static ConnectivityResult success(String scope, OffsetDateTime expiresAt, boolean caeEnabled) {\n+        return new ConnectivityResult(true, scope, expiresAt, caeEnabled, null);\n+    }\n+\n+    public static ConnectivityResult failure(String scope, boolean caeEnabled, String reason) {\n+        return new ConnectivityResult(false, scope, null, caeEnabled, reason);\n+    }\n+\n+    public void print(String testName) {\n+        System.out.println(testName + \" connectivity test:\");\n+        System.out.println(\"  Scope: \" + scope);\n+        System.out.println(\"  CAE-enabled request: \" + (caeEnabled ? \"yes\" : \"no\"));\n+        if (successful) {\n+            System.out.println(\"  Result: SUCCESS\");\n+            System.out.println(\"  Token expires at: \" + expiresAt);\n+        } else {\n+            System.out.println(\"  Result: FAILURE\");\n+            System.out.println(\"  Reason: \" + failureReason);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/AuthenticationFailureAnalyzer.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+\n+import java.util.Locale;\n+\n+final class AuthenticationFailureAnalyzer {\n+    private AuthenticationFailureAnalyzer() {\n+    }\n+\n+    static String describe(Throwable failure) {\n+        String details = collectMessages(failure);\n+        String normalized = details.toLowerCase(Locale.ROOT);\n+\n+        if (contains(normalized, \"aadsts7000222\", \"expired client secret\", \"certificate has expired\",\n+            \"expired certificate\")) {\n+            return \"The client secret or certificate has expired. \" + details;\n+        }\n+        if (contains(normalized, \"aadsts90002\", \"tenant not found\", \"invalid tenant\",\n+            \"aadsts700016\", \"application with identifier\")) {\n+            return \"The tenant ID is wrong, unavailable, or does not contain the configured application. \" + details;\n+        }\n+        if (contains(normalized, \"aadsts7000215\", \"invalid client secret\")) {\n+            return \"The client secret is invalid. \" + details;\n+        }\n+        if (contains(normalized, \"aadsts700027\", \"certificate\", \"client assertion is not within\")) {\n+            return \"The client certificate or assertion is invalid. \" + details;\n+        }\n+        if (contains(normalized, \"aadsts700024\", \"federated identity credential\", \"subject claim\",\n+            \"issuer claim\")) {\n+            return \"Workload identity federation is expired or does not match its federated credential. \" + details;\n+        }\n+        if (contains(normalized, \"aadsts500011\", \"invalid_resource\", \"resource principal\")) {\n+            return \"The requested Azure scope/resource is invalid for this tenant. \" + details;\n+        }\n+        if (failure instanceof CredentialUnavailableException\n+            || contains(normalized, \"credentialunavailableexception\", \"no managed identity\",\n+                \"managed identity endpoint\", \"identity not found\", \"no accounts were found\")) {\n+            return \"No configured identity is available in this environment. \" + details;\n+        }\n+        if (failure instanceof ClientAuthenticationException) {\n+            return \"Microsoft Entra ID rejected the credential. \" + details;\n+        }\n+        return failure.getClass().getSimpleName() + \": \" + details;\n+    }\n+\n+    private static boolean contains(String value, String... candidates) {\n+        for (String candidate : candidates) {\n+            if (value.contains(candidate)) {\n+                return true;\n+            }\n+        }\n+        return false;\n+    }\n+\n+    private static String collectMessages(Throwable failure) {\n+        StringBuilder messages = new StringBuilder();\n+        Throwable current = failure;\n+        while (current != null) {\n+            String message = current.getMessage();\n+            if (message != null \u0026\u0026 !message.isBlank() \u0026\u0026 messages.indexOf(message) \u003c 0) {\n+                if (!messages.isEmpty()) {\n+                    messages.append(\" Caused by: \");\n+                }\n+                messages.append(message.replaceAll(\"\\\\s+\", \" \").trim());\n+            }\n+            current = current.getCause();\n+        }\n+        return messages.isEmpty() ? \"No diagnostic message was provided.\" : messages.toString();\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/SyncCredentialConnectivityTester.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+\n+public final class SyncCredentialConnectivityTester {\n+    public ConnectivityResult test(\n+        TokenCredential credential,\n+        String scope,\n+        boolean caeEnabled\n+    ) {\n+        try {\n+            AccessToken token = credential.getTokenSync(new TokenRequestContext().addScopes(scope));\n+            ConnectivityResult result = ConnectivityResult.success(scope, token.getExpiresAt(), caeEnabled);\n+            result.print(\"Synchronous\");\n+            return result;\n+        } catch (RuntimeException failure) {\n+            ConnectivityResult result = ConnectivityResult.failure(\n+                scope,\n+                caeEnabled,\n+                AuthenticationFailureAnalyzer.describe(failure)\n+            );\n+            result.print(\"Synchronous\");\n+            return result;\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/AsyncCredentialConnectivityTester.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncCredentialConnectivityTester {\n+    public Mono\u003cConnectivityResult\u003e test(\n+        TokenCredential credential,\n+        String scope,\n+        boolean caeEnabled\n+    ) {\n+        return credential.getToken(new TokenRequestContext().addScopes(scope))\n+            .map(token -\u003e ConnectivityResult.success(scope, token.getExpiresAt(), caeEnabled))\n+            .onErrorResume(failure -\u003e Mono.just(ConnectivityResult.failure(\n+                scope,\n+                caeEnabled,\n+                AuthenticationFailureAnalyzer.describe(failure)\n+            )))\n+            .doOnNext(result -\u003e result.print(\"Asynchronous\"));\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/Main.java\n+package com.example.azureidentity;\n+\n+public final class Main {\n+    private static final String ARM_SCOPE = \"https://management.azure.com/.default\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        boolean enableCae = isCaeEnabled(args);\n+        DeploymentEnvironment environment = new EnvironmentDetector().detect();\n+        CredentialSelection selection = new AzureCredentialFactory().create(environment, enableCae);\n+\n+        System.out.println(\"Detected environment: \" + environment);\n+        System.out.println(\"Credential strategy: \" + selection.strategy());\n+        System.out.println(\"CAE requested: \" + selection.caeEnabled());\n+        System.out.println();\n+\n+        new SyncCredentialConnectivityTester()\n+            .test(selection.credential(), ARM_SCOPE, selection.caeEnabled());\n+        System.out.println();\n+        new AsyncCredentialConnectivityTester()\n+            .test(selection.credential(), ARM_SCOPE, selection.caeEnabled())\n+            .block();\n+    }\n+\n+    private static boolean isCaeEnabled(String[] args) {\n+        for (String arg : args) {\n+            if (\"--cae\".equalsIgnoreCase(arg)) {\n+                return true;\n+            }\n+        }\n+        return Boolean.parseBoolean(System.getenv().getOrDefault(\"AZURE_ENABLE_CAE\", \"false\"));\n+    }\n+}\n*** Add File: src/test/java/com/example/azureidentity/EnvironmentDetectorTest.java\n+package com.example.azureidentity;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class EnvironmentDetectorTest {\n+    private final EnvironmentDetector detector = new EnvironmentDetector();\n+\n+    @Test\n+    void defaultsToDeveloperEnvironment() {\n+        assertEquals(DeploymentEnvironment.DEV, detector.detect(Map.of()));\n+    }\n+\n+    @Test\n+    void detectsCiBeforeProductionMarkers() {\n+        assertEquals(\n+            DeploymentEnvironment.CI,\n+            detector.detect(Map.of(\"TF_BUILD\", \"True\", \"IDENTITY_ENDPOINT\", \"http://localhost/identity\"))\n+        );\n+    }\n+\n+    @Test\n+    void detectsManagedIdentityProductionEnvironment() {\n+        assertEquals(\n+            DeploymentEnvironment.PRODUCTION,\n+            detector.detect(Map.of(\"IDENTITY_ENDPOINT\", \"http://localhost/identity\"))\n+        );\n+    }\n+\n+    @Test\n+    void detectsKubernetesWorkloadIdentityEnvironment() {\n+        assertEquals(\n+            DeploymentEnvironment.PRODUCTION,\n+            detector.detect(Map.of(\"AZURE_FEDERATED_TOKEN_FILE\", \"/var/run/secrets/azure/tokens/azure-identity-token\"))\n+        );\n+    }\n+}\n*** Add File: README.md\n+# Azure credential chains for Java\n+\n+This Java 17 sample chooses an explicit Azure Identity credential chain for local development, CI, or production.\n+It only requests an Azure Resource Manager token; it does not create or modify Azure resources.\n+\n+| Environment | Detection examples | Credential strategy |\n+|---|---|---|\n+| Development | No CI or Azure hosting marker | Azure CLI, Azure Developer CLI, Azure PowerShell, then IntelliJ |\n+| CI | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, `BUILD_BUILDID` | Azure Pipelines service connection when configured, then `EnvironmentCredential` |\n+| Production | Managed identity endpoint, App Service, Container Apps, Kubernetes, or federated token marker | System/user-assigned managed identity, then configured Kubernetes workload identity |\n+\n+## Run\n+\n+```shell\n+mvn clean test\n+mvn exec:java\n+mvn exec:java -Dexec.args=\"--cae\"\n+```\n+\n+CAE can also be enabled with `AZURE_ENABLE_CAE=true`. Azure Identity applies CAE through\n+`TokenRequestContext.setCaeEnabled(true)`; the wrapper produced by the factory applies that setting to every request.\n+Developer credentials do not support CAE, so use the flag with CI or production credentials.\n+\n+## Configuration\n+\n+For a CI service principal using environment credentials, set `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and either\n+`AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`.\n+\n+For an Azure Pipelines workload-identity service connection, set:\n+\n+- `AZURE_TENANT_ID`\n+- `AZURE_CLIENT_ID`\n+- `AZURE_SERVICE_CONNECTION_ID` (the service connection resource ID)\n+- `SYSTEM_ACCESSTOKEN` (map `System.AccessToken` into this environment variable)\n+- `SYSTEM_OIDCREQUESTURI` (provided by Azure Pipelines)\n+\n+Production uses a system-assigned managed identity by default. Set `AZURE_MANAGED_IDENTITY_CLIENT_ID` to select a\n+user-assigned managed identity. Kubernetes workload identity is added as a fallback when `AZURE_TENANT_ID`,\n+`AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE` are all set.\n+\n+The identity needs an appropriate Azure RBAC role to use Azure Resource Manager. No secret or token is printed.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/pom.xml
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
+    <name>Azure credential chains by environment</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.identity.version>1.18.5</azure.identity.version>
+        <junit.version>5.13.4</junit.version>
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
+            </plugin>
+            <plugin>
+ 
... (truncated)
```

### ✅ 16. powershell — 7477ms

**Input:**

```json
{"command":"mvn --no-transfer-progress clean test","description":"Build project and run tests","initial_wait":180,"mode":"sync"}
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
[INFO] Building Azure credential chains by environment 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- clean:3.2.0:clean (default-clean) @ azure-credential-chains ---
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.0:compile (default-compile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 10 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-credential-chains ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.0:testCompile (default-testCompile) @ azure-credential-chains ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 1 source file with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.3:test (default-test) @ azure-credential-chains ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitpl
... (truncated)
```

### ✅ 17. apply_patch — 929ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azureidentity/AzureCredentialFactory.java\n@@\n     private static final String[] PIPELINES_CREDENTIAL_VARIABLES = {\n         \"AZURE_CLIENT_ID\",\n         \"AZURE_TENANT_ID\",\n         \"AZURE_SERVICE_CONNECTION_ID\",\n         \"SYSTEM_ACCESSTOKEN\",\n         \"SYSTEM_OIDCREQUESTURI\"\n     };\n+    private static final String[] PIPELINES_SERVICE_CONNECTION_MARKERS = {\n+        \"AZURE_SERVICE_CONNECTION_ID\",\n+        \"SYSTEM_ACCESSTOKEN\",\n+        \"SYSTEM_OIDCREQUESTURI\"\n+    };\n@@\n-        if (hasAny(variables, PIPELINES_CREDENTIAL_VARIABLES)) {\n+        if (hasAny(variables, PIPELINES_SERVICE_CONNECTION_MARKERS)) {\n@@\n-        if (hasAny(variables, WORKLOAD_IDENTITY_VARIABLES)) {\n+        if (isNonBlank(variables.get(\"AZURE_FEDERATED_TOKEN_FILE\"))) {\n*** Add File: src/test/java/com/example/azureidentity/AzureCredentialFactoryTest.java\n+package com.example.azureidentity;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class AzureCredentialFactoryTest {\n+    private final AzureCredentialFactory factory = new AzureCredentialFactory();\n+\n+    @Test\n+    void supportsCiServicePrincipalEnvironmentVariablesWithoutPipelineFederation() {\n+        CredentialSelection selection = factory.create(\n+            DeploymentEnvironment.CI,\n+            false,\n+            Map.of(\n+                \"AZURE_TENANT_ID\", \"tenant\",\n+                \"AZURE_CLIENT_ID\", \"client\",\n+                \"AZURE_CLIENT_SECRET\", \"secret\"\n+            )\n+        );\n+\n+        assertEquals(\"pipeline environment variables\", selection.strategy());\n+    }\n+\n+    @Test\n+    void rejectsPartialAzurePipelinesServiceConnectionConfiguration() {\n+        IllegalStateException failure = assertThrows(\n+            IllegalStateException.class,\n+            () -\u003e factory.create(\n+                DeploymentEnvironment.CI,\n+                false,\n+                Map.of(\"AZURE_SERVICE_CONNECTION_ID\", \"connection\")\n+            )\n+        );\n+\n+        assertTrue(failure.getMessage().contains(\"SYSTEM_ACCESSTOKEN\"));\n+        assertTrue(failure.getMessage().contains(\"SYSTEM_OIDCREQUESTURI\"));\n+    }\n+\n+    @Test\n+    void selectsUserAssignedManagedIdentityFromDedicatedVariable() {\n+        CredentialSelection selection = factory.create(\n+            DeploymentEnvironment.PRODUCTION,\n+            true,\n+            Map.of(\"AZURE_MANAGED_IDENTITY_CLIENT_ID\", \"managed-identity-client\")\n+        );\n+\n+        assertEquals(\"user-assigned managed identity\", selection.strategy());\n+        assertTrue(selection.caeEnabled());\n+    }\n+\n+    @Test\n+    void rejectsPartialWorkloadIdentityConfiguration() {\n+        IllegalStateException failure = assertThrows(\n+            IllegalStateException.class,\n+            () -\u003e factory.create(\n+                DeploymentEnvironment.PRODUCTION,\n+                false,\n+                Map.of(\"AZURE_FEDERATED_TOKEN_FILE\", \"token-file\")\n+            )\n+        );\n+\n+        assertTrue(failure.getMessage().contains(\"AZURE_CLIENT_ID\"));\n+        assertTrue(failure.getMessage().contains(\"AZURE_TENANT_ID\"));\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/src/main/java/com/example/azureidentity/AzureCredentialFactory.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/src/main/java/com/example/azureidentity/AzureCredentialFactory.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/src/main/java/com/example/azureidentity/AzureCredentialFactory.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-java-credential-chain-java-azure-skills-azure-skill-mcp-2342086755/src/main/java/com/example/azureidentity/AzureCredentialFactory.java
@@ -24,6 +24,11 @@
         "SYSTEM_ACCESSTOKEN",
         "SYSTEM_OIDCREQUESTURI"
     };
+    private static final String[] PIPELINES_SERVICE_CONNECTION_MARKERS = {
+        "AZURE_SERVICE_CONNECTION_ID",
+        "SYSTEM_ACCESSTOKEN",
+        "SYSTEM_OIDCREQUESTURI"
+    };
     private static final String[] WORKLOAD_IDENTITY_VARIABLES = {
         "AZURE_CLIENT_ID",
         "AZURE_TENANT_ID",
@@ -67,7 +72,7 @@
         ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
         List<String> strategies = new ArrayList<>();
 
-        if (hasAny(variables, PIPELINES_CREDENTIAL_VARIABLES)) {
+        if (hasAny(variables, PIPELINES_SERVICE_CONNECTION_MARKERS)) {
             requireAll(variables, "Azure Pipelines workload federation", PIPELINES_CREDENTIAL_VARIABLES);
             chain.addLast(new AzurePipelinesCredentialBuilder()
                 .clientId(variables.get("AZURE_CLIENT_ID"))
@@ -98,7 +103,7 @@
 
         List<String> strategies = new ArrayList<>();
         strategies.add(managedIdentityStrategy);
-        if (hasAny(variables, WORKLOAD_IDENTITY_VARIABLES)) {
+        if (isNonBlank(variables.get("AZURE_FEDERATED_TOKEN_FILE"))) {
... (truncated)
```

### 🔧 18. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress test package","description":"Run tests and package sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\azureidentity\AsyncCredentialConnectivityTester.java`
- `src\main\java\com\example\azureidentity\AuthenticationFailureAnalyzer.java`
- `src\main\java\com\example\azureidentity\AzureCredentialFactory.java`
- `src\main\java\com\example\azureidentity\CaeAwareTokenCredential.java`
- `src\main\java\com\example\azureidentity\ConnectivityResult.java`
- `src\main\java\com\example\azureidentity\CredentialSelection.java`
- `src\main\java\com\example\azureidentity\DeploymentEnvironment.java`
- `src\main\java\com\example\azureidentity\EnvironmentDetector.java`
- `src\main\java\com\example\azureidentity\Main.java`
- `src\main\java\com\example\azureidentity\SyncCredentialConnectivityTester.java`
- `src\test\java\com\example\azureidentity\AzureCredentialFactoryTest.java`
- `src\test\java\com\example\azureidentity\EnvironmentDetectorTest.java`

## Copilot Response

I’ll ground the implementation in the current Azure Identity Java patterns, then scaffold and verify the Maven project locally.

The initial best-practices endpoint timed out, so I’m using the skill’s Java guidance and checking the current SDK API surface before writing the sample.

The API confirms CAE is a per-token-request capability, so the factory will return a wrapped credential that consistently applies the CAE flag while preserving environment-specific chain behavior.

The build exposed no compiler issues, but reviewing the CI branch found an important configuration edge case: ordinary service-principal environment variables must not be mistaken for an Azure Pipelines service connection. I’m tightening that distinction and covering it with tests.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No Azure SDK client lifecycle violations are present.

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
hyoka run --prompt-id identity-dp-java-credential-chain --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
