# Evaluation Report: identity-dp-java-default-credential

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 640.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-default-credential` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 15/17 |
| Duration | 640.6s |
| Timestamp | 2026-08-27T16:06:26Z |
| Files Generated | 0 |
| Event Count | 4043 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 398.4s |
| Review | 240.9s |
| **Total** | **640.6s** |

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
| Token Usage | in=165262 out=3813 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 2825ms

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

### ✅ 2. view — 195655ms

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

### ✅ 3. view — 193806ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/troubleshooting.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/troubleshooting.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/troubleshooting.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/troubleshooting.md
@@ -1,270 +1,270 @@
 # Troubleshooting Microsoft Entra App Registration
 
 This guide helps you diagnose and fix common issues with app registrations and authentication.
 
 ## Table of Contents
 
 - [Authentication Errors](#authentication-errors)
 - [Token Issues](#token-issues)
 - [Permission Problems](#permission-problems)
 - [Redirect URI Issues](#redirect-uri-issues)
 - [Application Configuration](#application-configuration)
 - [Debugging Tools](#debugging-tools)
 
 ## Authentication Errors
 
 ### Redirect URI Mismatch
 
 **Error message:**
 ```
 AADSTS50011: The redirect URI 'http://localhost:3000' specified in the request 
 does not match the redirect URIs configured for the application.
 ```
 
 **Cause:** The redirect URI in your authentication request doesn't exactly match what's registered.
 
 **Solutions:**
 
 1. **Check exact match** (case-sensitive, trailing slash matters):
    ```
    Registered: https://myapp.com/callback
    Request:    https://myapp.com/callback/  ❌ (trailing slash)
    Request:    https://MyApp.com/callback   ❌ (case difference)
    Request:    https://myapp.com/callback   ✅
    ```
 
 2. **Add URI to app registration:**
    ```bash
    # Portal: Authentication → Add redirect URI
    # CLI:
    az ad app update --id $APP_ID \
      --web-redirect-uris "http://
... (truncated)
```

### ✅ 4. view — 191953ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\troubleshooting.md"}
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

### ❌ 5. azure-get_azure_bestpractices (via azure) — 195655ms

**Input:**

```json
{"intent":"Provide Java code showing Azure SDK authentication with DefaultAzureCredential and Azure Key Vault SecretClient, including troubleshooting guidance.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 195656ms

**Input:**

```json
{"intent":"Find current official documentation for DefaultAzureCredential in Azure Identity for Java, including Maven dependencies, credential chain order, local versus hosted behavior, logging and troubleshooting, and a Key Vault SecretClient example.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. web_search (via github-mcp-server) — 108368ms

**Input:**

```json
{"query":"site:learn.microsoft.com Java DefaultAzureCredential credential chain order EnvironmentCredential WorkloadIdentityCredential ManagedIdentityCredential IntelliJ Azure CLI Azure PowerShell Azure Developer CLI BrokerCredential current"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To troubleshoot authentication failures with `DefaultAzureCredential` in the Azure Identity Java SDK—especially \"credential unavailable\" errors or \"authentication failed\" messages—Microsoft provides comprehensive guidance. Here are the key steps and resources:\n\n---\n\n## 1. **Enable Logging**\nEnable logging in your Java app to see which credential types in the `DefaultAzureCredential` chain are being attempted and where the failure occurs. Logging will help you diagnose configuration issues or missing dependencies. See [how to enable and configure logging](https://learn.microsoft.com/en-us/azure/developer/java/sdk/identity/troubleshoot-identity-authentication?view=azure-java-stable#enable-and-configure-logging) in the Microsoft documentation.\n\n---\n\n## 2. **Common Exception Types**\n- **CredentialUnavailableException:** Indicates that a particular credential in the chain is not available (e.g., missing environment variables, no managed identity configured, missing Azure CLI, etc.).\n- **ClientAuthenticationException:** General authentication failure exception. It may wrap specific causes such as invalid credentials or permission denial【3:0†source】【3:1†source】.\n\n---\n\n## 3. **Typical Causes & Fixes**\n\n### **A. All Credentials Unavailable**\nThis usually means every credential in the chain (environment, managed identity, Azure CLI, etc.) failed. **Enable logging** to find out exactly which ones are tried and why each failed.\n\n- Make sure you set necessary environment variables if you want to use `EnvironmentCredential`.\n    - For service principal: `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_SECRET`\n    - For managed identity: ensure the VM, app service, or container has managed identity enabled and `AZURE_CLIENT_ID` (for user-assigned) is set if necessary.\n\n- Make sure Azure CLI or the required developer tool is installed and logged in if you depend on it for local authentication (`az
... (truncated)
```

### ✅ 8. web_search (via github-mcp-server) — 108023ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure SDK Java Key Vault SecretClient DefaultAzureCredential Maven azure-security-keyvault-secrets azure-identity dependency BOM current"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To use the Azure Key Vault SecretClient with DefaultAzureCredential in Java, you need to add the following Maven dependencies. The current recommended way to manage dependencies is using the Azure SDK BOM (Bill of Materials), which simplifies version management by aligning compatible versions of Azure libraries.\n\n### 1. Add the Azure SDK BOM to `dependencyManagement`:\nReplace `{bom_version_to_target}` with the latest version, which you can find on Maven Central or in the official Microsoft docs. As of the most recent documentation, the version examples are:\n\n```xml\n<dependencyManagement>\n  <dependencies>\n    <dependency>\n      <groupId>com.azure</groupId>\n      <artifactId>azure-sdk-bom</artifactId>\n      <version>{bom_version_to_target}</version>\n      <type>pom</type>\n      <scope>import</scope>\n    </dependency>\n  </dependencies>\n</dependencyManagement>\n```\n*For the very latest version, check [Azure SDK BOM README](https://learn.microsoft.com/en-us/java/api/overview/azure/sdk?view=azure-java-stable) or [Maven Central](https://search.maven.org/search?q=g:com.azure%20AND%20a:azure-sdk-bom). (As of August 2026, 1.2.19 is current)【3:2†source】.*\n\n### 2. Add individual dependencies WITHOUT specifying versions:\nThe versions will be controlled by the BOM for compatibility:\n\n```xml\n<dependencies>\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-security-keyvault-secrets</artifactId>\n  </dependency>\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-identity</artifactId>\n  </dependency>\n</dependencies>\n```\nReference for details: [Key Vault Secrets client library for Java \"Getting Started\"](https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable)【3:0†source】 and the [Azure Identity client library for Java docs](https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=az
... (truncated)
```

### ✅ 9. web_search (via github-mcp-server) — 107694ms

**Input:**

```json
{"query":"site:learn.microsoft.com Java Azure Identity troubleshooting DefaultAzureCredential logging credential unavailable authentication failed"}
```

**Output:**

```
{"type":"output_text","text":{"value":"For Java, the Azure SDK's `DefaultAzureCredential` class attempts to authenticate using a series of underlying credentials in a specific order (known as the credential chain). It checks each credential in sequence and uses the first one that successfully provides a token. The order for Java (as of recent official documentation) is:\n\n1. **EnvironmentCredential**\n2. **WorkloadIdentityCredential**\n3. **ManagedIdentityCredential**\n4. **IntelliJCredential** (Azure Toolkit for IntelliJ)\n5. **AzureCliCredential** (uses `az login`)\n6. **AzurePowerShellCredential** (uses `Connect-AzAccount`)\n7. **AzureDeveloperCliCredential** (uses Azure Developer CLI `azd login`)\n8. **A broker-enabled instance of InteractiveBrowserCredential** (requires additional broker configuration)\n\nThis means that if you have multiple possible sources of authentication on your machine or environment, `DefaultAzureCredential` will attempt each of these in this specific order until one provides a valid authentication token. This order allows for flexible authentication scenarios that work well both in local development and Azure-hosted (e.g., managed identity) environments.\n\nYou can find this documented in official Microsoft sources for Java:  \n> DefaultAzureCredential attempts to authenticate with each of these credentials, in the following order, stopping when one provides a token:\n> - EnvironmentCredential\n> - WorkloadIdentityCredential\n> - ManagedIdentityCredential\n> - IntelliJCredential\n> - AzureCliCredential\n> - AzurePowerShellCredential\n> - AzureDeveloperCliCredential\n> - A broker-enabled instance of InteractiveBrowserCredential【3:0†source】.\n\nIf you want more details or configuration options, you should check [DefaultAzureCredential Class documentation](https://learn.microsoft.com/en-us/java/api/com.azure.identity.defaultazurecredential?view=azure-java-stable).","annotations":[{"text":"","start_index":1624,"end_index":1636,"url_c
... (truncated)
```

### ✅ 10. web_fetch — 107280ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/identity/azure-identity/README.md"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/keyvault/azure-security-keyvault-secrets/README.md:
# Azure Key Vault Secret client library for Java
Azure Key Vault is a cloud service that provides secure storage for secrets, such as passwords and database connection strings.

The Azure Key Vault Secrets client library allows you to securely store and tightly control the access to tokens, passwords, API keys, and other secrets. This library offers operations to create, retrieve, update, delete, purge, backup, restore, and list the secrets and its versions.

Use the Azure Key Vault Secrets client library to create and manage secrets.

[Source code][source_code] | [API reference documentation][api_documentation] | [Product documentation][azkeyvault_docs] | [Samples][secrets_samples]

## Getting started
### Include the package
#### Include the BOM file
Please include the `azure-sdk-bom` to your project to take dependency on the General Availability (GA) version of the library. In the following snippet, replace the {bom_version_to_target} placeholder with the version number.
To learn more about the BOM, see the [AZURE SDK BOM README](https://github.com/Azure/azure-sdk-for-java/blob/main/sdk/boms/azure-sdk-bom/README.md).

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

and then include the direct dependency in the dependencies section without the version tag as shown below.

```xml
<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies
... (truncated)
```

### ✅ 11. web_fetch — 106874ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-java/main/sdk/keyvault/azure-security-keyvault-secrets/README.md"}
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

## Copilot Response

I’m checking the current Java credential chain and package guidance against first-party references, since both can change between Azure Identity releases.

`DefaultAzureCredential` lets the same Java code use developer credentials locally and passwordless workload credentials when deployed to Azure.

## 1. Maven dependencies

Add Azure Identity and the Key Vault Secrets client:

```xml
<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-identity</artifactId>
        <version>1.18.4</version>
    </dependency>

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
        <version>4.11.0</version>
    </dependency>
</dependencies>
```

These are the stable versions currently listed in the upstream Azure SDK readmes. For larger applications, Microsoft recommends importing the [`azure-sdk-bom`](https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/boms/azure-sdk-bom) and omitting versions from individual Azure dependencies.

If you want brokered authentication, including Visual Studio Code authentication through the Azure Resources extension, also add `com.azure:azure-identity-broker` using the BOM-managed version.

## 2. Complete `SecretClient` example

```java
package example;

import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultExample {
    private KeyVaultExample() {
    }

    public static void main(String[] args) {
        String vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");
        String secretName = requireEnvironmentVariable("SECRET_NAME");

        DefaultAzureCredential credential =
            new DefaultAzureCredentialBuilder().build();

        SecretClient secretClient =
            new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential)
                .buildClient();

        KeyVaultSecret secret = secretClient.getSecret(secretName);

        // Avoid printing secret.getValue() or otherwise exposing secret material.
        System.out.printf(
            "Retrieved secret '%s', version '%s'.%n",
            secret.getName(),
            secret.getProperties().getVersion());
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }
}
```

Configure non-secret application settings:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://my-vault.vault.azure.net"
$env:SECRET_NAME = "database-password"
mvn compile exec:java -Dexec.mainClass="example.KeyVaultExample"
```

The identity selected by `DefaultAzureCredential` must have data-plane permission to read the secret—for example, the **Key Vault Secrets User** RBAC role at the narrowest appropriate scope.

`DefaultAzureCredential` and Azure SDK clients are thread-safe. Create them once and reuse them rather than rebuilding them for every request.

## 3. Credential chain order

For current Azure Identity Java releases, credentials are tried in this order, stopping when one obtains a token:

1. **`EnvironmentCredential`**  
   Uses environment-based service-principal configuration, commonly:
   - `AZURE_CLIENT_ID`
   - `AZURE_TENANT_ID`
   - `AZURE_CLIENT_SECRET`

   Certificate-based environment configuration is also supported.

2. **`WorkloadIdentityCredential`**  
   Uses Microsoft Entra Workload ID federation, commonly on AKS or another Kubernetes environment.

3. **`ManagedIdentityCredential`**  
   Uses the managed identity assigned to an Azure resource such as App Service, Functions, a VM, or Container Apps.

4. **`IntelliJCredential`**  
   Uses the account signed into the Azure Toolkit for IntelliJ.

5. **`AzureCliCredential`**  
   Uses the account selected by `az login`.

6. **`AzurePowerShellCredential`**  
   Uses the account selected by `Connect-AzAccount`.

7. **`AzureDeveloperCliCredential`**  
   Uses the account selected by `azd auth login`.

8. **Broker-enabled `InteractiveBrowserCredential`**  
   Available when broker support is installed and configured. On Windows, this uses Web Account Manager and can support the account signed into the VS Code Azure Resources extension.

The chain is configurable through `DefaultAzureCredentialBuilder`, including methods such as `excludeAzureCliCredential()` and `excludeManagedIdentityCredential()`.

**Continuation behavior:** developer credentials are generally all attempted until one succeeds. For deployed credentials—environment, workload identity, and managed identity—an available credential that attempts token acquisition but fails causes the chain to stop. This prevents a broken production identity from silently falling back to an unrelated developer identity.

## 4. Local development versus Azure

### Local development

Managed identity and workload identity normally aren't available, so the chain proceeds to a signed-in development tool. Azure CLI is a common option:

```powershell
az login
az account set --subscription "<subscription-id-or-name>"
```

Your signed-in user needs the appropriate Key Vault data-plane role. Azure subscription **Owner** or **Contributor** alone does not grant permission to read Key Vault secret values.

Be aware that fully configured `AZURE_CLIENT_*` variables take precedence over Azure CLI and IDE credentials. Stale environment variables are a frequent reason the application uses an unexpected identity.

### Azure-hosted application

Enable a managed identity on the Azure resource and grant that identity the necessary Key Vault role. The application code does not contain a client secret; `DefaultAzureCredential` detects the managed identity endpoint automatically.

For a user-assigned managed identity, identify the intended identity explicitly:

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder()
        .managedIdentityClientId(
            requireEnvironmentVariable("AZURE_CLIENT_ID"))
        .build();
```

On AKS, prefer Microsoft Entra Workload ID. The injected workload-identity settings allow `WorkloadIdentityCredential` to succeed before managed identity is attempted.

For security-sensitive production applications, consider using the specific credential—such as `ManagedIdentityCredential` or `WorkloadIdentityCredential`—instead of the complete default chain. This makes identity selection deterministic and reduces troubleshooting ambiguity.

## 5. Troubleshooting failures

1. **Read the complete exception.** `DefaultAzureCredential` reports why each attempted credential was unavailable or failed.
   - `CredentialUnavailableException`: the credential cannot run in this environment.
   - `ClientAuthenticationException`: the credential attempted authentication but Microsoft Entra ID rejected it.

2. **Enable Azure Identity logging.** Azure SDK for Java uses SLF4J. Configure your logging implementation to enable `DEBUG` for:
   - `com.azure.identity`
   - `com.azure.core`

   Logs reveal which credential was selected and why earlier credentials failed. Do not enable sensitive HTTP-body logging in production.

3. **Check which identity is actually active.** Look for unexpected `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, or `AZURE_CLIENT_SECRET` variables. Locally, also verify the active CLI account:

   ```powershell
   az account show
   ```

4. **Separate authentication from authorization.**
   - Authentication errors mention token acquisition, tenant, client, certificate, or managed-identity failures.
   - Key Vault `403 Forbidden` usually means authentication succeeded but the selected identity lacks a Key Vault data-plane role.
   - `401 Unauthorized` can indicate an invalid token audience, tenant mismatch, or authentication configuration problem.

5. **Check environment-specific configuration.**
   - **Environment credential:** all required variables must be present and refer to the same app registration.
   - **Azure CLI:** ensure the CLI is installed, `az login` has completed, and the correct tenant/subscription is selected.
   - **Managed identity:** ensure it is enabled on the hosting resource. For multiple user-assigned identities, provide the desired client ID.
   - **Workload identity:** verify the service account annotation, federated identity credential, tenant/client IDs, and projected token file.
   - **Sovereign cloud:** configure the correct `authorityHost` and use the matching Key Vault DNS suffix.

6. **Check networking and time.** The application must reach Microsoft Entra token endpoints and the Key Vault endpoint. A substantially incorrect system clock can also invalidate tokens.

Official references:

- [DefaultAzureCredential for Java](https://learn.microsoft.com/java/api/com.azure.identity.defaultazurecredential?view=azure-java-stable)
- [Credential chains in Azure Identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)
- [Azure Identity troubleshooting for Java](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview)
- [Key Vault Secrets client library for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable)

## Reviewer Notes (LLM-as-Judge)

### Summary

The sole criterion passes because the example creates no Azure SDK client requiring explicit resource closure.

### Strengths

- Uses the synchronous SecretClient appropriately without unnecessary resource-management code.

## Grader Results

- default-azure-credential.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Maven dependency for `azure-identity`: Pass
      - `DefaultAzureCredentialBuilder` pattern: Pass
      - Credential chain order in Java SDK: Pass
      - Passing credential to client builders (e.g., `SecretClientBuilder`): Pass
      - Logging configuration for authentication troubleshooting: Pass
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
hyoka run --prompt-id identity-dp-java-default-credential --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
