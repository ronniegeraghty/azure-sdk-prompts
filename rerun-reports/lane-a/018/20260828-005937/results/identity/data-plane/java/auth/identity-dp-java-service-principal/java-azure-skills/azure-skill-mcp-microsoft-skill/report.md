# Evaluation Report: identity-dp-java-service-principal

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 712.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-service-principal` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 17/17 |
| Duration | 712.3s |
| Timestamp | 2026-08-27T16:59:37Z |
| Files Generated | 0 |
| Event Count | 2739 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 475.9s |
| Review | 235.2s |
| **Total** | **712.3s** |

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
| Token Usage | in=154401 out=3405 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 4173ms

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

### ✅ 2. view — 185043ms

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

### ❌ 3. azure-get_azure_bestpractices (via azure) — 185045ms

**Input:**

```json
{"intent":"Get current best practices for Java code that authenticates to Azure with a service principal client secret and uses an Azure SDK client, including secure secret management and authentication error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 4. azure-documentation (via azure) — 185287ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Azure Identity Java ClientSecretCredentialBuilder, service principal environment variables, authentication errors, and using TokenCredential with Azure Storage Blob clients.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 185283ms

**Input:**

```json
{"intent":"Get best practices for Java Azure SDK service-principal client-secret authentication, secret management, least privilege, and exception handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

## Copilot Response

Use `azure-identity` to create the credential and pass it directly to an Azure SDK client builder. The example below uses Azure Blob Storage and reads credentials exclusively from environment variables.

### 1. Maven dependencies

Using the Azure SDK BOM keeps Azure library versions compatible:

```xml
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

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-storage-blob</artifactId>
    </dependency>
</dependencies>
```

### 2. Environment variables

`AZURE_CLIENT_SECRET` must contain the **secret value**, not the secret's ID.

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<storage-account>.blob.core.windows.net"
```

The service principal also needs an appropriate data-plane role, such as **Storage Blob Data Reader**, scoped as narrowly as possible.

### 3. Complete Java example

```java
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpRequestException;
import com.azure.identity.ClientSecretCredential;
import com.azure.identity.ClientSecretCredentialBuilder;
import com.azure.identity.CredentialUnavailableException;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobStorageException;

public final class ServicePrincipalExample {
    private ServicePrincipalExample() {
    }

    public static void main(String[] args) {
        try {
            String tenantId = requiredEnvironmentVariable("AZURE_TENANT_ID");
            String clientId = requiredEnvironmentVariable("AZURE_CLIENT_ID");
            String clientSecret = requiredEnvironmentVariable("AZURE_CLIENT_SECRET");
            String blobEndpoint =
                requiredEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT");

            ClientSecretCredential credential =
                new ClientSecretCredentialBuilder()
                    .tenantId(tenantId)
                    .clientId(clientId)
                    .clientSecret(clientSecret)
                    .build();

            BlobServiceClient blobServiceClient =
                new BlobServiceClientBuilder()
                    .endpoint(blobEndpoint)
                    .credential(credential)
                    .buildClient();

            // Credential creation is lazy. This operation triggers token acquisition
            // and verifies that the identity has access to Blob Storage.
            blobServiceClient.listBlobContainers().forEach(
                container -> System.out.println(container.getName())
            );
        } catch (IllegalStateException exception) {
            System.err.println(
                "Authentication configuration is incomplete: "
                    + exception.getMessage()
            );
            System.exit(2);
        } catch (CredentialUnavailableException exception) {
            System.err.println(
                "The configured credential could not acquire a token: "
                    + exception.getMessage()
            );
            System.exit(3);
        } catch (ClientAuthenticationException exception) {
            System.err.println(
                "Microsoft Entra authentication failed. Check the tenant ID, "
                    + "client ID, client secret value, and secret expiration."
            );
            // Log exception details only to a protected diagnostic destination.
            System.err.println("Authentication details: " + exception.getMessage());
            System.exit(4);
        } catch (BlobStorageException exception) {
            if (exception.getStatusCode() == 403) {
                System.err.println(
                    "Authentication succeeded, but the service principal is not "
                        + "authorized to access this storage account."
                );
            } else if (exception.getStatusCode() == 401) {
                System.err.println(
                    "Blob Storage rejected the access token. Verify the tenant, "
                        + "resource endpoint, and service-principal configuration."
                );
            } else {
                System.err.printf(
                    "Blob Storage request failed with status %d and error code %s.%n",
                    exception.getStatusCode(),
                    exception.getErrorCode()
                );
            }
            System.exit(5);
        } catch (HttpRequestException exception) {
            System.err.println(
                "Azure could not be reached. Check DNS, network connectivity, "
                    + "proxy configuration, and the storage endpoint."
            );
            System.err.println("Network details: " + exception.getMessage());
            System.exit(6);
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Environment variable " + name + " is not set."
            );
        }

        return value;
    }
}
```

`ClientSecretCredentialBuilder.build()` constructs the credential but does not immediately contact Microsoft Entra ID. Invalid or expired credentials are normally detected when the SDK client performs its first operation.

### 4. Secret-management best practices

- Never hardcode or commit client secrets. Keep `.env` files and local configuration containing secrets out of source control.
- Use a CI/CD secret store or protected environment-variable facility, and prevent secrets from appearing in command output.
- Prefer **managed identity** for Azure-hosted applications and workload identity/federated credentials for CI/CD or Kubernetes. These avoid stored secrets.
- If a secret is unavoidable, rotate it regularly, use a short expiration, and support overlapping secrets during rotation.
- Grant the service principal only the required roles at the narrowest resource scope.
- Never log `AZURE_CLIENT_SECRET`, authorization headers, or access tokens. Authentication error messages may also contain tenant or application identifiers, so send detailed diagnostics only to protected logs.
- Treat authentication and authorization separately: invalid credentials generally produce `ClientAuthenticationException` or HTTP `401`; valid credentials lacking the required role generally produce HTTP `403`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The sample does not leave any AutoCloseable Azure SDK client unclosed.

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
hyoka run --prompt-id identity-dp-java-service-principal --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
