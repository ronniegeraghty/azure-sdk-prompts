# Evaluation Report: identity-dp-java-managed-identity

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 511.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-managed-identity` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 15/18 |
| Duration | 511.5s |
| Timestamp | 2026-08-27T16:38:35Z |
| Files Generated | 0 |
| Event Count | 4992 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 296.8s |
| Review | 213.6s |
| **Total** | **511.5s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=153021 out=3074 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 4196ms

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

### ✅ 2. view — 8267ms

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

### ✅ 3. glob — 6668ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-documentation (via azure) — 183695ms

**Input:**

```json
{"intent":"Find official Azure Java guidance for ManagedIdentityCredential, DefaultAzureCredential local development, user-assigned identity selectors, and troubleshooting.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

## Copilot Response

Managed identities let Azure SDK clients obtain Microsoft Entra tokens without storing credentials. The identity must be attached to the Azure host and granted the required data-plane or management-plane role on the target resource.

## 1. System-assigned vs. user-assigned

| Type | Lifecycle | Azure resource attachment | Credential selection |
|---|---|---|---|
| **System-assigned** | Created and deleted with the hosting resource | Belongs to one resource | No client ID required |
| **User-assigned** | Independent Azure resource | Can be attached to multiple resources | Select using client ID, resource ID, or object ID |

Use system-assigned identity when the workload and identity share the same lifecycle. Use user-assigned identity when multiple workloads need the same identity or the identity must survive resource replacement.

> For user-assigned identity, `clientId(...)` expects the identity's **application/client ID**, not its principal/object ID.

## 2. Dependencies

Using the Azure SDK BOM avoids manually coordinating compatible library versions:

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

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-storage-blob</artifactId>
    </dependency>
</dependencies>
```

## 3. System-assigned managed identity

Enable the system-assigned identity on the VM, App Service, Function App, Container App, or other supported host. Then create the credential without an identity selector:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class SystemAssignedExample {
    public static BlobServiceClient createBlobClient() {
        TokenCredential credential =
            new ManagedIdentityCredentialBuilder().build();

        return new BlobServiceClientBuilder()
            .endpoint("https://<storage-account>.blob.core.windows.net")
            .credential(credential)
            .buildClient();
    }

    public static void main(String[] args) {
        BlobServiceClient blobClient = createBlobClient();

        blobClient.listBlobContainers()
            .forEach(container ->
                System.out.println(container.getName()));
    }
}
```

No tenant ID, secret, or client ID is needed. The SDK discovers the managed-identity endpoint exposed by the Azure host.

## 4. User-assigned managed identity

Attach the user-assigned identity to the Azure host and select it by client ID:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class UserAssignedExample {
    public static BlobServiceClient createBlobClient() {
        String managedIdentityClientId =
            requireEnvironmentVariable("AZURE_MANAGED_IDENTITY_CLIENT_ID");

        TokenCredential credential =
            new ManagedIdentityCredentialBuilder()
                .clientId(managedIdentityClientId)
                .build();

        return new BlobServiceClientBuilder()
            .endpoint("https://<storage-account>.blob.core.windows.net")
            .credential(credential)
            .buildClient();
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is missing: " + name);
        }
        return value;
    }

    public static void main(String[] args) {
        BlobServiceClient blobClient = createBlobClient();

        blobClient.listBlobContainers()
            .forEach(container ->
                System.out.println(container.getName()));
    }
}
```

You can also identify the user-assigned identity by Azure resource ID:

```java
TokenCredential credential =
    new ManagedIdentityCredentialBuilder()
        .resourceId(
            "/subscriptions/<subscription-id>"
                + "/resourceGroups/<resource-group>"
                + "/providers/Microsoft.ManagedIdentity"
                + "/userAssignedIdentities/<identity-name>")
        .build();
```

Or by principal/object ID if needed:

```java
TokenCredential credential =
    new ManagedIdentityCredentialBuilder()
        .objectId("<user-assigned-identity-object-id>")
        .build();
```

Client ID is generally the clearest and most portable selector.

## 5. Using the credential with other Azure SDK clients

Azure client builders generally accept the common `TokenCredential` interface.

### Key Vault

```java
import com.azure.core.credential.TokenCredential;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();

String value = secretClient.getSecret("database-password").getValue();
```

### Service Bus

```java
import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusSenderClient;

ServiceBusSenderClient sender = new ServiceBusClientBuilder()
    .fullyQualifiedNamespace("<namespace>.servicebus.windows.net")
    .credential(credential)
    .sender()
    .queueName("<queue-name>")
    .buildClient();
```

### Event Hubs

```java
import com.azure.messaging.eventhubs.EventHubClientBuilder;
import com.azure.messaging.eventhubs.EventHubProducerClient;

EventHubProducerClient producer = new EventHubClientBuilder()
    .fullyQualifiedNamespace("<namespace>.servicebus.windows.net")
    .eventHubName("<event-hub-name>")
    .credential(credential)
    .buildProducerClient();
```

The identity needs an appropriate role, such as:

- **Storage Blob Data Reader/Contributor** for Blob Storage
- **Key Vault Secrets User** for reading Key Vault secrets
- **Azure Service Bus Data Sender/Receiver**
- **Azure Event Hubs Data Sender/Receiver**

Control-plane roles such as `Contributor` do not necessarily grant access to resource data.

## 6. Local development fallback

`ManagedIdentityCredential` only works on a supported Azure host. A local computer does not expose Azure's managed-identity endpoint.

### Recommended: `DefaultAzureCredential`

Use one credential chain for local development and Azure hosting:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

TokenCredential credential =
    new DefaultAzureCredentialBuilder().build();
```

Locally, it can use developer credentials such as Azure CLI authentication. On Azure, it can use the system-assigned managed identity.

For local CLI authentication:

```powershell
az login
```

The signed-in developer must have the same resource permissions required by the application.

For a user-assigned managed identity in Azure:

```java
String managedIdentityClientId =
    System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID");

DefaultAzureCredentialBuilder builder =
    new DefaultAzureCredentialBuilder();

if (managedIdentityClientId != null &&
    !managedIdentityClientId.isBlank()) {
    builder.managedIdentityClientId(managedIdentityClientId);
}

TokenCredential credential = builder.build();
```

Set `AZURE_MANAGED_IDENTITY_CLIENT_ID` only in the deployed environment. Locally, `DefaultAzureCredential` can continue to use Azure CLI or another configured developer credential.

### Explicit environment-based selection

For stricter production behavior, choose credentials explicitly:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureCliCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

public final class CredentialFactory {
    public static TokenCredential create() {
        String environment =
            System.getenv().getOrDefault("APP_ENVIRONMENT", "local");

        if ("azure".equalsIgnoreCase(environment)) {
            String clientId =
                System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID");

            ManagedIdentityCredentialBuilder builder =
                new ManagedIdentityCredentialBuilder();

            if (clientId != null && !clientId.isBlank()) {
                builder.clientId(clientId);
            }

            return builder.build();
        }

        return new AzureCliCredentialBuilder().build();
    }
}
```

This prevents a deployed application from silently falling back to developer-oriented credentials if managed identity is misconfigured.

For CI/CD, prefer workload identity federation or an environment credential supplied by the pipeline. Avoid putting client secrets in source code or configuration files.

## 7. Error handling

Credential construction normally does not request a token. Authentication failures usually occur during the first SDK operation, so catch errors around the operation rather than only around the builder:

```java
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.CredentialUnavailableException;

try {
    blobClient.listBlobContainers()
        .forEach(container ->
            System.out.println(container.getName()));
} catch (CredentialUnavailableException exception) {
    System.err.println(
        "Managed identity is unavailable on this host: "
            + exception.getMessage());
    throw exception;
} catch (ClientAuthenticationException exception) {
    System.err.println(
        "Microsoft Entra authentication failed: "
            + exception.getMessage());
    throw exception;
} catch (HttpResponseException exception) {
    int status = exception.getResponse().getStatusCode();

    if (status == 403) {
        System.err.println(
            "Authentication succeeded, but the identity lacks permission.");
    } else {
        System.err.println(
            "Azure service request failed with HTTP " + status);
    }

    throw exception;
}
```

Do not treat `CredentialUnavailableException` as success or continue anonymously. In production, surface it as a configuration or deployment failure.

## Troubleshooting checklist

1. **Confirm the identity is enabled or attached.** A system-assigned identity must be enabled on the host; a user-assigned identity must be attached to it.
2. **Verify the identity selector.** For `.clientId(...)`, use the user-assigned identity's client ID—not its object ID or Azure resource ID.
3. **Check RBAC on the correct scope.** Assign the required data-plane role at the resource, resource-group, or subscription scope.
4. **Allow for RBAC propagation.** New role assignments can take several minutes to become effective.
5. **Distinguish authentication from authorization.** A token acquisition failure is authentication; HTTP `403` usually means the identity authenticated but lacks permission.
6. **Check the service endpoint.** Ensure the account, vault, namespace, and cloud suffix are correct.
7. **Avoid testing `ManagedIdentityCredential` locally.** Use `DefaultAzureCredential`, Azure CLI, or an explicit development credential instead.
8. **Enable temporary SDK diagnostics.** Set `AZURE_LOG_LEVEL=verbose`, reproduce the failure, then disable verbose logging after troubleshooting.
9. **Check networking.** Private endpoints, DNS, proxies, and outbound restrictions can block either token acquisition or the target service.
10. **On AKS, prefer workload identity.** AKS workload identity normally uses `WorkloadIdentityCredential` or `DefaultAzureCredential`, rather than relying on node-level managed identity.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- The non-AutoCloseable BlobServiceClient and SecretClient examples do not require try-with-resources.

### Issues

- ServiceBusSenderClient sender and EventHubProducerClient producer are not closed.

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
  - Try-With-Resources for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 7.7% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Try-With-Resources for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 13.00** | **Σ 10.8333** | **83.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-managed-identity --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
