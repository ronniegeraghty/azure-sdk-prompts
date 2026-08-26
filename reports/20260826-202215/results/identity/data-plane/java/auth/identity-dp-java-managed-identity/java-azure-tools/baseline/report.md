# Evaluation Report: identity-dp-java-managed-identity

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 331.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-managed-identity` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 16/18 |
| Duration | 331.0s |
| Timestamp | 2026-08-26T15:37:37Z |
| Files Generated | 0 |
| Event Count | 5497 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 94.0s |
| Review | 236.1s |
| **Total** | **331.0s** |

## Configuration

- **name:** java-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=41487 out=3230 |
| Turn Count | 2 |

## Prompt Sent

```
Show me how to use
Managed Identity to authenticate Azure SDK clients in Java. Cover:
1. System-assigned vs user-assigned managed identity
2. How to create a ManagedIdentityCredential for each type
3. Using it with Azure SDK clients
4. Local development fallback strategies
5. Error handling and troubleshooting

Provide examples for both system-assigned and user-assigned identity.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Official Azure Java SDK documentation ManagedIdentityCredentialBuilder system-assigned user-assigned clientId resourceId DefaultAzureCredential local development troubleshooting 2026"}
```

## Copilot Response

## 1. System-assigned vs. user-assigned identity

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created and deleted with the Azure resource | Independent Azure resource |
| Sharing | Used by one hosting resource | Can be attached to multiple resources |
| Credential configuration | No identity ID required | Specify its **client ID** or **resource ID** |
| Typical use | One workload, one identity | Shared identity or multiple identities attached to a resource |

In either case, grant the identity the necessary Azure RBAC role or service-specific access policy. Authentication can succeed while the SDK operation still returns `403 Forbidden` if authorization is missing.

## 2. Dependency

Use the Azure SDK BOM so compatible library versions are selected together:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version><!-- current Azure SDK BOM version --></version>
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

    <!-- Example service clients -->
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-storage-blob</artifactId>
    </dependency>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies>
```

Replace the BOM placeholder with the current version from the [Azure SDK BOM documentation](https://learn.microsoft.com/azure/developer/java/sdk/get-started-maven).

## 3. Creating `ManagedIdentityCredential`

### System-assigned identity

Enable system-assigned managed identity on the Azure host, then build the credential without an identifier:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .build();
```

The Azure Identity library discovers the managed-identity endpoint supplied by App Service, Functions, Container Apps, Azure VMs, AKS workload environments, and other supported hosts.

### User-assigned identity using its client ID

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

String clientId = System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID");

if (clientId == null || clientId.isBlank()) {
    throw new IllegalStateException(
        "AZURE_MANAGED_IDENTITY_CLIENT_ID must contain the managed identity client ID");
}

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(clientId)
    .build();
```

The value must be the managed identity application's **client ID**, not its object/principal ID.

### User-assigned identity using its Azure resource ID

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

String resourceId = System.getenv("AZURE_MANAGED_IDENTITY_RESOURCE_ID");

if (resourceId == null || resourceId.isBlank()) {
    throw new IllegalStateException(
        "AZURE_MANAGED_IDENTITY_RESOURCE_ID must contain the managed identity resource ID");
}

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .resourceId(resourceId)
    .build();
```

A resource ID looks like:

```text
/subscriptions/<subscription-id>/resourceGroups/<resource-group>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity-name>
```

Using a resource ID can avoid ambiguity and makes the selected Azure resource explicit. Availability may depend on the managed-identity implementation used by the hosting service; client ID is the most broadly used selector.

## 4. Using the credential with Azure SDK clients

A `TokenCredential` can be reused safely across service clients. Do not create one for every request.

### Azure Blob Storage

```java
import com.azure.core.credential.TokenCredential;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

BlobServiceClient blobServiceClient = new BlobServiceClientBuilder()
    .endpoint("https://myaccount.blob.core.windows.net")
    .credential(credential)
    .buildClient();

blobServiceClient
    .getBlobContainerClient("documents")
    .listBlobs()
    .forEach(blob -> System.out.println(blob.getName()));
```

The identity needs an appropriate data-plane role, such as **Storage Blob Data Reader** or **Storage Blob Data Contributor**. A management-plane role such as Contributor does not automatically grant access to blob data.

### Azure Key Vault with a user-assigned identity

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

String clientId = System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID");

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(clientId)
    .build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://my-vault.vault.azure.net")
    .credential(credential)
    .buildClient();

String value = secretClient.getSecret("database-password").getValue();
```

Grant the identity an appropriate role, such as **Key Vault Secrets User**, when the vault uses Azure RBAC.

### Reusable client factory

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class AzureClients {
    private AzureClients() {}

    public static TokenCredential systemAssignedCredential() {
        return new ManagedIdentityCredentialBuilder().build();
    }

    public static TokenCredential userAssignedCredential(String clientId) {
        if (clientId == null || clientId.isBlank()) {
            throw new IllegalArgumentException(
                "A user-assigned managed identity client ID is required");
        }

        return new ManagedIdentityCredentialBuilder()
            .clientId(clientId)
            .build();
    }

    public static BlobServiceClient blobs(
            String accountName,
            TokenCredential credential) {

        return new BlobServiceClientBuilder()
            .endpoint("https://" + accountName + ".blob.core.windows.net")
            .credential(credential)
            .buildClient();
    }
}
```

System-assigned usage:

```java
TokenCredential credential = AzureClients.systemAssignedCredential();
BlobServiceClient blobs = AzureClients.blobs("myaccount", credential);
```

User-assigned usage:

```java
TokenCredential credential = AzureClients.userAssignedCredential(
    System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID"));

BlobServiceClient blobs = AzureClients.blobs("myaccount", credential);
```

## 5. Local development fallback

`ManagedIdentityCredential` normally cannot authenticate on a developer workstation because there is no Azure-managed-identity endpoint. Avoid replacing managed identity with a client secret unless your application specifically requires service-principal testing.

### Recommended: `DefaultAzureCredential`

`DefaultAzureCredential` can use developer credentials locally and managed identity after deployment:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

TokenCredential credential = new DefaultAzureCredentialBuilder()
    .build();
```

Locally, it can use supported developer credentials such as Azure CLI or IDE authentication. In Azure, it can use the system-assigned managed identity.

For a user-assigned managed identity:

```java
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId(
        System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID"))
    .build();
```

The client ID only selects the managed identity when that part of the credential chain is reached; local developer authentication still uses the signed-in developer identity.

### More deterministic environment-specific selection

If you do not want a broad credential chain in production, explicitly choose the credential:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

public final class CredentialFactory {
    private CredentialFactory() {}

    public static TokenCredential create() {
        String environment = System.getenv()
            .getOrDefault("APP_ENVIRONMENT", "development");

        if ("production".equalsIgnoreCase(environment)) {
            String clientId =
                System.getenv("AZURE_MANAGED_IDENTITY_CLIENT_ID");

            if (clientId != null && !clientId.isBlank()) {
                return new ManagedIdentityCredentialBuilder()
                    .clientId(clientId)
                    .build();
            }

            return new ManagedIdentityCredentialBuilder().build();
        }

        return new DefaultAzureCredentialBuilder()
            .excludeManagedIdentityCredential(true)
            .build();
    }
}
```

Excluding managed identity locally prevents unnecessary attempts to contact a managed-identity endpoint. Local authorization still must mirror production: grant the developer identity access to the development storage account, vault, or other service.

## 6. Error handling

Credential and service failures represent different problems:

```java
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.CredentialUnavailableException;
import com.azure.storage.blob.models.BlobStorageException;

try {
    blobServiceClient
        .getBlobContainerClient("documents")
        .listBlobs()
        .forEach(blob -> System.out.println(blob.getName()));
} catch (CredentialUnavailableException e) {
    System.err.println(
        "Managed identity is unavailable on this host: " + e.getMessage());
    throw e;
} catch (ClientAuthenticationException e) {
    System.err.println(
        "Azure could not issue an access token: " + e.getMessage());
    throw e;
} catch (BlobStorageException e) {
    System.err.printf(
        "Storage rejected the request: status=%d, errorCode=%s%n",
        e.getStatusCode(),
        e.getErrorCode());
    throw e;
} catch (HttpResponseException e) {
    System.err.println(
        "Azure service request failed with HTTP "
            + e.getResponse().getStatusCode());
    throw e;
}
```

Do not silently switch to another identity after an authentication failure in production. That can cause a workload to run under an unexpectedly privileged identity. Use an intentional credential chain or explicit environment configuration instead.

## 7. Troubleshooting checklist

| Symptom | Likely cause |
|---|---|
| Credential unavailable locally | Expected for `ManagedIdentityCredential`; use `DefaultAzureCredential` locally |
| No managed identity endpoint in Azure | Identity is not enabled, not attached, or the hosting environment does not support that managed-identity mechanism |
| User-assigned identity not found | Wrong client ID, resource ID, or identity not attached to the host |
| Multiple user-assigned identities | Explicitly call `.clientId(...)` or `.resourceId(...)` |
| `401 Unauthorized` | Token acquisition, tenant, audience, endpoint, or identity-selection problem |
| `403 Forbidden` | Identity authenticated but lacks the required data-plane role/access policy |
| Access fails immediately after role assignment | Azure RBAC and identity changes can take time to propagate |
| Works locally but not in Azure | Local developer identity has permission, but the managed identity does not |
| Works in Azure but not locally | Developer identity lacks access, CLI/IDE login expired, or the wrong tenant/subscription is active |
| Timeout contacting identity endpoint | Hosting configuration, proxy, firewall, or managed-identity endpoint connectivity issue |

Additional checks:

1. Verify the managed identity is enabled or attached to the **same Azure resource running the application**.
2. For user-assigned identity, verify the application uses the identity’s client ID—not its principal/object ID.
3. Check permissions on the exact target resource and scope.
4. Distinguish management-plane roles from data-plane roles.
5. Enable debug logging for `com.azure.identity` through your SLF4J logging implementation, but do not log access tokens.
6. If using `DefaultAzureCredential`, review its logs to identify which credential in the chain was selected.
7. Avoid manually calling or hard-coding managed-identity endpoint URLs; let `azure-identity` detect the hosting environment.

References: [ManagedIdentityCredentialBuilder](https://learn.microsoft.com/java/api/com.azure.identity.managedidentitycredentialbuilder), [Azure Identity for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme), [system-assigned identity tutorial](https://learn.microsoft.com/azure/developer/java/sdk/authentication/system-assigned-managed-identity), and [user-assigned identity tutorial](https://learn.microsoft.com/azure/developer/java/sdk/authentication/user-assigned-managed-identity).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK client is left unclosed.

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
hyoka run --prompt-id identity-dp-java-managed-identity --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
