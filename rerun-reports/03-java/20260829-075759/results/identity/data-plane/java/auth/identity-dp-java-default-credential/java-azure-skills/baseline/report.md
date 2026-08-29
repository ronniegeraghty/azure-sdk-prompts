# Evaluation Report: identity-dp-java-default-credential

**Config:** java-azure-skills/baseline | **Result:** ✅ PASSED | **Duration:** 414.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-default-credential` |
| Config | java-azure-skills/baseline |
| Result | ✅ PASSED |
| Score | 17/17 |
| Duration | 414.6s |
| Timestamp | 2026-08-29T01:07:13Z |
| Files Generated | 0 |
| Event Count | 4439 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 126.5s |
| Review | 286.4s |
| **Total** | **414.6s** |

## Configuration

- **name:** java-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=79696 out=2931 |
| Turn Count | 3 |

## Tool Calls

### ✅ 1. web_search (via github-mcp-server) — 64073ms

**Input:**

```json
{"query":"site:learn.microsoft.com azure identity Java DefaultAzureCredential credential chain order 2026 official documentation"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To troubleshoot authentication issues with `DefaultAzureCredential` in the Azure Identity library for Java, Microsoft recommends a series of steps, including exception handling and enabling diagnostic logging to get a detailed trace of what each credential in the chain is attempting. Here’s how to proceed, drawing directly from the official Microsoft documentation:\n\n### 1. Handling Exceptions\n\n- **ClientAuthenticationException**: Thrown when authentication fails. The message includes details on the failure source.\n- **CredentialUnavailableException**: Specific to when the credential can’t authenticate due to missing configuration. Useful for identifying if a particular credential (like environment, managed identity, CLI, etc.) is unavailable in the execution context.\n\nExample of catching an authentication exception:\n```java\ntry {\n    // code using DefaultAzureCredential\n} catch (ClientAuthenticationException e) {\n    e.printStackTrace();\n}\n```\n\n### 2. Enable and Configure Logging\n\nEnabling logging is critical for diagnosing which credentials are being tried, which one is failing, and why. The logs will show a step-by-step account of attempts by all credentials within the `DefaultAzureCredential` chain.\n\n**To enable logging in Java:**\n```java\n// Set system properties for Azure SDK logging (SLF4J example)\nSystem.setProperty(\"org.slf4j.simpleLogger.defaultLogLevel\", \"debug\");\nSystem.setProperty(\"org.slf4j.simpleLogger.log.com.azure.identity\", \"debug\");\n```\nOr, using log4j or another supported logging framework, configure the Azure identity logger (e.g., `com.azure.identity`) at DEBUG or TRACE level.\n\n**Official docs for more advanced logging setups:**\n[Azure Identity authentication troubleshooting - Logging section (Java)](https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable#logging)【3:0†source】\n\n### 3. Diagnose Problems with DefaultAzureC
... (truncated)
```

### ✅ 2. web_search (via github-mcp-server) — 62087ms

**Input:**

```json
{"query":"site:learn.microsoft.com Java Azure Key Vault SecretClient DefaultAzureCredential Maven azure-security-keyvault-secrets azure-identity official"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To use the Azure Key Vault SecretClient in Java with DefaultAzureCredential via Maven, referencing the official packages (azure-security-keyvault-secrets and azure-identity), here’s a precise, official Microsoft Learn-based guide:\n\n---\n\n## 1. Add Maven Dependencies\n\nIt is recommended to use the Azure SDK BOM for dependency management. Add this to your `pom.xml`:\n\n```xml\n<dependencyManagement>\n  <dependencies>\n    <dependency>\n      <groupId>com.azure</groupId>\n      <artifactId>azure-sdk-bom</artifactId>\n      <version>1.2.18</version> <!-- Check for latest version -->\n      <type>pom</type>\n      <scope>import</scope>\n    </dependency>\n  </dependencies>\n</dependencyManagement>\n```\n\nThen, under `<dependencies>`:\n```xml\n<dependency>\n  <groupId>com.azure</groupId>\n  <artifactId>azure-security-keyvault-secrets</artifactId>\n</dependency>\n<dependency>\n  <groupId>com.azure</groupId>\n  <artifactId>azure-identity</artifactId>\n</dependency>\n```\nOr, if not using the BOM, specify explicit versions:\n\n```xml\n<dependency>\n  <groupId>com.azure</groupId>\n  <artifactId>azure-security-keyvault-secrets</artifactId>\n  <version>4.11.0</version>\n</dependency>\n<dependency>\n  <groupId>com.azure</groupId>\n  <artifactId>azure-identity</artifactId>\n  <version>1.18.5</version>\n</dependency>\n```\n【3:1†source】【3:7†source】\n\n---\n\n## 2. Sample Java Code\n\n```java\nimport com.azure.identity.DefaultAzureCredentialBuilder;\nimport com.azure.security.keyvault.secrets.SecretClient;\nimport com.azure.security.keyvault.secrets.SecretClientBuilder;\n\npublic class KeyVaultSecretSample {\n    public static void main(String[] args) {\n        String keyVaultUrl = \"https://<your-key-vault-name>.vault.azure.net/\";\n\n        // Build a Key Vault Secret client with DefaultAzureCredential\n        SecretClient secretClient = new SecretClientBuilder()\n            .vaultUrl(keyVaultUrl)\n            .credenti
... (truncated)
```

### ✅ 3. web_search (via github-mcp-server) — 60309ms

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Identity Java troubleshoot DefaultAzureCredential authentication logging official"}
```

**Output:**

```
{"type":"output_text","text":{"value":"As of the latest official Microsoft documentation for 2026, the `DefaultAzureCredential` class in the Azure Identity SDK for Java attempts authentication using a chain of credential types in a specific order. The order is designed to streamline authentication for both local development and deployment within Azure-hosted environments.\n\n### DefaultAzureCredential Chain Order (Java)\n\nWhen you use `DefaultAzureCredential`, the credential types are attempted **in the following order** until one succeeds in acquiring a token:\n\n1. **EnvironmentCredential**  \n   Looks for credentials in environment variables, typically set for service principals, certificates, or managed identities.\n\n2. **WorkloadIdentityCredential**  \n   Supports Kubernetes and other cloud-native workload identity scenarios.\n\n3. **ManagedIdentityCredential**  \n   Attempts to authenticate using a managed identity available in an Azure service like VM, App Service, or Azure Functions.\n\n4. **IntelliJCredential**  \n   Uses credentials from Azure Toolkit for IntelliJ, for interactive sign-ins via the IntelliJ IDE.\n\n5. **VisualStudioCodeCredential**  \n   Authenticates using credentials from Visual Studio Code IDE.\n\n6. **AzureCliCredential**  \n   Uses credentials from the Azure CLI (via `az login`).\n\n7. **AzurePowerShellCredential**  \n   Uses credentials from an active login in Azure PowerShell.\n\n8. **AzureDeveloperCliCredential**  \n   Uses credentials from the Azure Developer CLI.\n\n9. **InteractiveBrowserCredential (broker-enabled instance)**  \n   Prompts for a user login via the browser if all previous credentials fail.\n\nThis means `DefaultAzureCredential` checks each in turn, from top to bottom, until it finds one that works for the current environment.\n\nYou can configure credentials by setting environment variables or using setters on the `DefaultAzureCredentialBuilder`. For more details, consult the docs:  \n[DefaultAzureCredential Cla
... (truncated)
```

### ✅ 4. web_fetch — 10753ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Identity client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-08-24T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-identity, entra-id
ms.service: entra-id
locale: en-us
document_id: a0c04cc9-4127-ee61-109d-2d672636a46e
document_version_independent_id: 0e5b5560-efe9-0c0a-9a63-38dc58e6af8b
updated_at: 2026-08-25T11:08:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/identity-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/75b1775676c4b42b2562bb1c23a9a8a867297328/docs-ref-services/latest/identity-readme.md
git_commit_id: 75b1775676c4b42b2562bb1c23a9a8a867297328
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 2370
asset_id: api/overview/azure/identity-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/identity-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
spProducts:
- https://authoring-docs-microsoft
... (truncated)
```

### ✅ 5. web_fetch — 9193ms

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Key Vault Secret client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-05-27T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-security-keyvault-secrets, keyvault
ms.service: keyvault
locale: en-us
document_id: 009f13e3-83f0-8f58-4be1-ebb785e6a344
document_version_independent_id: 6c845f98-c6a3-d3cd-c0b7-23f6775ef5ad
updated_at: 2026-05-28T11:07:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/security-keyvault-secrets-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/4f5bf09b69d6f35fcbd5a7a92ad75413fdbedaf8/docs-ref-services/latest/security-keyvault-secrets-readme.md
git_commit_id: 4f5bf09b69d6f35fcbd5a7a92ad75413fdbedaf8
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 1484
asset_id: api/overview/azure/security-keyvault-secrets-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/security-keyvault-secrets-readme.md
cmProducts:
- https://au
... (truncated)
```

### ✅ 6. web_fetch — 7249ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains:
---
layout: Conceptual
title: Credential chains in the Azure Identity library for Java - Java on Azure | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains
breadcrumb_path: /azure/developer/breadcrumb/toc.json
uhfHeaderId: azure
feedback_system: Standard
recommendations: true
feedback_help_link_type: get-help-at-qna
feedback_help_link_url: https://learn.microsoft.com/answers/questions/
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
author: bmitchell287
manager: nitinme
ms.author: brendm
ms.devlang: java
ms.service: azure-java
ms.subservice: java-sdk
description: This article describes the DefaultAzureCredential and ChainedTokenCredential classes in the Azure Identity library.
ms.date: 2026-02-24T00:00:00.0000000Z
ms.topic: concept-article
ms.reviewer: scaddie
ms.custom: devx-track-java
ai-usage: ai-assisted
locale: en-us
document_id: 54718866-4bea-1757-d2ff-144acdb1bb70
document_version_independent_id: 54718866-4bea-1757-d2ff-144acdb1bb70
updated_at: 2026-03-04T00:04:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/live/articles/java/sdk/authentication/credential-chains.md
gitcommit: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/ac67d79b6d52421f5635629e7be7d54c779d6c95/articles/java/sdk/authentication/credential-chains.md
git_commit_id: ac67d79b6d52421f5635629e7be7d54c779d6c95
site_name: Docs
depot_name: MSDN.azure-dev-docs
page_type: conceptual
toc_rel: ../toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/MSDN.azure-dev-docs/{branchName}{pdfName}
word_count: 1453
asset_id: java/sdk/authentication/credential-chains
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/java/sdk/authentication/credential-chains.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/d
... (truncated)
```

### ✅ 7. web_fetch — 5569ms

**Input:**

```json
{"max_length":15000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview:
---
layout: Conceptual
title: Troubleshoot Azure Identity authentication - Azure SDK for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview
breadcrumb_path: /azure/developer/breadcrumb/toc.json
uhfHeaderId: azure
feedback_system: Standard
recommendations: true
feedback_help_link_type: get-help-at-qna
feedback_help_link_url: https://learn.microsoft.com/answers/questions/
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
author: bmitchell287
manager: nitinme
ms.author: brendm
ms.devlang: java
ms.service: azure-java
ms.subservice: java-sdk
description: Learn how to troubleshoot Azure Identity authentication issues in the Azure SDK for Java. Resolve common errors and get mitigation steps.
ms.date: 2026-07-27T00:00:00.0000000Z
ms.topic: troubleshooting-general
ms.custom: devx-track-java, devx-track-extended-java
ms.reviewer: jogiles
locale: en-us
document_id: b76e7ea0-b79d-a913-1fdb-777cff4947af
document_version_independent_id: b76e7ea0-b79d-a913-1fdb-777cff4947af
updated_at: 2026-07-27T17:05:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/live/articles/java/sdk/troubleshooting-authentication-overview.md
gitcommit: https://github.com/MicrosoftDocs/azure-dev-docs-pr/blob/242d11c3bf5e37d84c7db9948dc37e357f9881f5/articles/java/sdk/troubleshooting-authentication-overview.md
git_commit_id: 242d11c3bf5e37d84c7db9948dc37e357f9881f5
site_name: Docs
depot_name: MSDN.azure-dev-docs
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/MSDN.azure-dev-docs/{branchName}{pdfName}
word_count: 811
asset_id: java/sdk/troubleshooting-authentication-overview
moniker_range_name: 
monikers: []
item_type: Content
source_path: articles/java/sdk/troubleshooting-authentication-overview.md
... (truncated)
```

## Copilot Response

I’ll verify the current Java package versions and credential-chain behavior against Microsoft’s official documentation.

`DefaultAzureCredential` provides one `TokenCredential` that works across developer machines and Azure-hosted applications by trying an ordered chain of authentication methods.

## 1. Maven dependencies

Use the Azure SDK BOM so compatible versions are selected together:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>1.2.38</version>
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

Check [Maven Central](https://central.sonatype.com/artifact/com.azure/azure-sdk-bom) for the latest BOM version. Without the BOM, current Microsoft documentation references `azure-identity` 1.18.5 and `azure-security-keyvault-secrets` 4.11.0.

For Visual Studio Code or OS account-broker authentication, also add:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity-broker</artifactId>
</dependency>
```

The broker package is optional if you use Azure CLI, IntelliJ, environment credentials, workload identity, or managed identity.

## 2. Creating and using `DefaultAzureCredential`

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder().build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://my-vault.vault.azure.net")
    .credential(credential)
    .buildClient();
```

Constructing the client does not immediately authenticate. A token is normally requested when the first service operation is performed, such as `getSecret()`.

For a user-assigned managed identity:

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder()
        .managedIdentityClientId(System.getenv("AZURE_CLIENT_ID"))
        .build();
```

For a system-assigned managed identity, do not specify a client ID.

## 3. Credential-chain order

The current Java chain attempts:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal configuration in environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS |
| 3 | `ManagedIdentityCredential` | An identity assigned to the Azure host |
| 4 | `IntelliJCredential` | Azure Toolkit for IntelliJ login |
| 5 | `VisualStudioCodeCredential` | VS Code Azure Resources login; broker package required |
| 6 | `AzureCliCredential` | Account authenticated through `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated through `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated through `azd auth login` |
| 9 | Broker credential | Default OS account; broker package required |

The chain stops when a credential successfully returns a token. It does **not** continue just because that identity lacks permission to access Key Vault.

There is also an important continuation-policy distinction:

- **Deployed-service credentials**—environment, workload identity, and managed identity—stop the chain if they are configured and attempt authentication but receive an authentication error.
- **Developer credentials** are all attempted until one succeeds, even when an earlier developer credential encounters an authentication error.

You can narrow the chain with `AZURE_TOKEN_CREDENTIALS`:

```text
AZURE_TOKEN_CREDENTIALS=dev
AZURE_TOKEN_CREDENTIALS=prod
AZURE_TOKEN_CREDENTIALS=AzureCliCredential
AZURE_TOKEN_CREDENTIALS=ManagedIdentityCredential
```

`dev` excludes deployed-service credentials; `prod` excludes developer credentials. Selecting an individual credential requires `azure-identity` 1.17.0 or later.

## 4. Local development versus Azure

**Local development**

`DefaultAzureCredential` normally finds no workload or managed identity. It falls through to a signed-in developer tool, commonly Azure CLI:

```powershell
az login
az account show
```

The signed-in user must have permission to read secrets—for example, the **Key Vault Secrets User** RBAC role on the vault. IntelliJ, VS Code, Azure PowerShell, and Azure Developer CLI logins are alternatives.

For local service-principal authentication, configure:

```text
AZURE_TENANT_ID=<tenant-id>
AZURE_CLIENT_ID=<application-client-id>
AZURE_CLIENT_SECRET=<client-secret>
```

Do not commit these values. Use environment variables or a local secret-management mechanism.

**Azure-hosted application**

Use managed identity whenever the hosting service supports it:

1. Enable a system-assigned or user-assigned identity on the host.
2. Grant that identity the required Key Vault role.
3. Deploy the same application code without client secrets.

On AKS, workload identity is generally preferred and is attempted before managed identity.

Be careful with environment variables in Azure: a complete but invalid service-principal configuration is tried before managed identity and can prevent fallback. Production deployments can set `AZURE_TOKEN_CREDENTIALS=prod` or specifically `ManagedIdentityCredential` to make the intended authentication mechanism explicit.

## 5. Complete `SecretClient` example

This example reads the vault URL and secret name from environment variables, creates a client, and retrieves a secret without printing its value:

```java
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

        SecretClient secretClient =
            new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential)
                .buildClient();

        try {
            KeyVaultSecret secret = secretClient.getSecret(secretName);

            System.out.printf(
                "Retrieved secret '%s', version '%s'.%n",
                secret.getName(),
                secret.getProperties().getVersion());
        } catch (ClientAuthenticationException exception) {
            System.err.println(
                "Microsoft Entra authentication failed: "
                    + exception.getMessage());
            System.exit(1);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();

            System.err.printf(
                "Key Vault request failed with HTTP %d: %s%n",
                statusCode,
                exception.getMessage());
            System.exit(1);
        }
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

Example configuration:

```powershell
$env:KEY_VAULT_URL = "https://my-vault.vault.azure.net"
$env:SECRET_NAME = "database-password"
mvn compile exec:java
```

Avoid logging or printing `secret.getValue()`.

## Troubleshooting failures

1. **Read the entire exception.** `DefaultAzureCredential` reports each credential it tried and why it was unavailable or failed. Preserve correlation IDs and timestamps when escalating an Entra ID error.

2. **Enable Azure Identity logging.** Configure your SLF4J implementation so `com.azure.identity` logs at `DEBUG` or `TRACE`. The useful messages look like:

   ```text
   Attempted credential EnvironmentCredential is unavailable.
   Attempted credential ManagedIdentityCredential is unavailable.
   Attempted credential AzureCliCredential returns a token.
   ```

   Do not enable `enableUnsafeSupportLogging()` outside controlled support diagnostics—it can expose tokens and personal data.

3. **Distinguish authentication from authorization.**
   - `ClientAuthenticationException`: no usable token was obtained.
   - HTTP `401`: invalid token, audience, tenant, or authentication configuration.
   - HTTP `403`: authentication usually succeeded, but the selected identity lacks permission or Key Vault network/access policies deny access.

4. **Confirm which identity won.** An unexpected environment credential can take precedence over Azure CLI locally or managed identity in Azure. Temporarily select one credential with `AZURE_TOKEN_CREDENTIALS` to isolate the problem.

5. **Check common configuration errors.**
   - Wrong tenant, client ID, or client-secret **ID** used instead of the secret **value**.
   - Expired service-principal secret or certificate.
   - Developer tool logged into the wrong tenant/subscription.
   - User-assigned managed identity client ID missing or incorrect.
   - Workload identity issuer, subject, audience, or federated credential mismatch.
   - Key Vault URL incorrect; it should resemble `https://<name>.vault.azure.net`.
   - Required Key Vault RBAC role not assigned, or the assignment has not propagated.
   - Vault firewall/private endpoint/DNS configuration blocking the request.

Official references: [credential-chain order](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), [Azure Identity for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme), [Key Vault Secrets for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme), and [authentication troubleshooting](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criterion passed.

### Strengths

- The example correctly avoids unnecessary resource-management code for SecretClient.

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
hyoka run --prompt-id identity-dp-java-default-credential --config java-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
