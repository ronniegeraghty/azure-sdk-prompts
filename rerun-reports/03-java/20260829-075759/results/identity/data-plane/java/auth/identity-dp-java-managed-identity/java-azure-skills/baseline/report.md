# Evaluation Report: identity-dp-java-managed-identity

**Config:** java-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 321.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-managed-identity` |
| Config | java-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 15/18 |
| Duration | 321.0s |
| Timestamp | 2026-08-29T01:25:37Z |
| Files Generated | 0 |
| Event Count | 5611 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 45.2s |
| Review | 273.9s |
| **Total** | **321.0s** |

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
| Token Usage | in=19653 out=2788 |
| Turn Count | 1 |

## Copilot Response

# Managed Identity authentication for Azure SDK clients in Java

Managed Identity lets Azure-hosted Java applications authenticate without storing client secrets. The Azure Identity library obtains tokens from the hosting environment and Azure SDK clients use those tokens automatically.

## 1. System-assigned vs. user-assigned identity

| Type | Lifecycle | Identity selection | Common use |
|---|---|---|---|
| **System-assigned** | Created on an Azure resource and deleted with it | No client ID is normally specified | One identity dedicated to one resource |
| **User-assigned** | Independent Azure resource; can be attached to multiple compute resources | Specify its **client ID**, resource ID, or object ID | Shared identity, stable lifecycle, explicit identity selection |

Managed Identity only proves the application's identity. You must separately grant that identity the required Azure RBAC role, such as **Storage Blob Data Reader** or **Key Vault Secrets User**.

## 2. Dependencies

Azure SDK client libraries commonly bring in `azure-core`, but add `azure-identity` explicitly.

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>${azure.identity.version}</version>
</dependency>
```

Add the service-specific library too, for example:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-storage-blob</artifactId>
    <version>${azure.storage.blob.version}</version>
</dependency>
```

Use the current Azure SDK BOM or current versions from Maven Central rather than hard-coding the illustrative placeholders above.

## 3. System-assigned managed identity

For a system-assigned identity, create `ManagedIdentityCredential` without specifying an identity:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .build();
```

Use it with an Azure SDK client:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class SystemAssignedExample {
    public static void main(String[] args) {
        String accountName = requireEnv("AZURE_STORAGE_ACCOUNT_NAME");

        TokenCredential credential = new ManagedIdentityCredentialBuilder()
            .build();

        BlobServiceClient blobServiceClient = new BlobServiceClientBuilder()
            .endpoint("https://" + accountName + ".blob.core.windows.net")
            .credential(credential)
            .buildClient();

        blobServiceClient.listBlobContainers()
            .forEach(container -> System.out.println(container.getName()));
    }

    private static String requireEnv(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }
}
```

On supported Azure compute, the credential discovers the system-assigned identity endpoint automatically.

## 4. User-assigned managed identity

When multiple identities may be available, select the user-assigned identity explicitly. Using its client ID is the most common approach:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

String managedIdentityClientId =
    requireEnv("AZURE_CLIENT_ID");

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(managedIdentityClientId)
    .build();
```

A complete Blob Storage example:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class UserAssignedExample {
    public static void main(String[] args) {
        String accountName = requireEnv("AZURE_STORAGE_ACCOUNT_NAME");
        String clientId = requireEnv("AZURE_CLIENT_ID");

        TokenCredential credential = new ManagedIdentityCredentialBuilder()
            .clientId(clientId)
            .build();

        BlobServiceClient blobServiceClient = new BlobServiceClientBuilder()
            .endpoint("https://" + accountName + ".blob.core.windows.net")
            .credential(credential)
            .buildClient();

        blobServiceClient.listBlobContainers()
            .forEach(container -> System.out.println(container.getName()));
    }

    private static String requireEnv(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }
}
```

Depending on your `azure-identity` version, user-assigned identities can also be selected by Azure resource ID:

```java
TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .resourceId(
        "/subscriptions/<subscription-id>"
            + "/resourceGroups/<resource-group>"
            + "/providers/Microsoft.ManagedIdentity"
            + "/userAssignedIdentities/<identity-name>")
    .build();
```

Prefer client ID when you want the most broadly familiar configuration. Do not confuse these identifiers:

- **Client ID**: application/client identifier; accepted by `.clientId(...)`.
- **Resource ID**: full Azure resource path; accepted by `.resourceId(...)`.
- **Object/principal ID**: service principal object identifier; not interchangeable with the client ID.

## 5. Using the credential with other Azure SDK clients

The same `TokenCredential` works across Azure SDK client builders.

### Key Vault Secrets

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();

String value = secretClient.getSecret("database-password").getValue();
```

For a user-assigned identity:

```java
TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(System.getenv("AZURE_CLIENT_ID"))
    .build();
```

### Azure Resource Manager

```java
import com.azure.core.credential.TokenCredential;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .build();

AzureResourceManager azure = AzureResourceManager
    .authenticate(credential, new AzureProfile())
    .withDefaultSubscription();
```

The identity needs control-plane roles for Resource Manager operations. Data-plane roles, such as Storage Blob Data Reader, do not grant Resource Manager access.

## 6. Local development fallback strategies

`ManagedIdentityCredential` is intentionally Azure-host-only. It normally cannot authenticate on a developer workstation.

### Recommended: `DefaultAzureCredential`

Use the same application code locally and in Azure:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

TokenCredential credential = new DefaultAzureCredentialBuilder()
    .build();
```

Typical behavior:

- Locally, it can authenticate through configured developer credentials, such as Azure CLI or IDE authentication.
- In Azure, it can use Managed Identity.
- If a user-assigned identity is attached, configure its client ID:

```java
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId(System.getenv("AZURE_CLIENT_ID"))
    .build();
```

Be careful with `AZURE_CLIENT_ID`: it is also used by environment-based service-principal authentication. For clearer configuration, use a dedicated application variable:

```java
String managedIdentityClientId =
    System.getenv("MANAGED_IDENTITY_CLIENT_ID");

DefaultAzureCredentialBuilder builder =
    new DefaultAzureCredentialBuilder();

if (managedIdentityClientId != null
        && !managedIdentityClientId.isBlank()) {
    builder.managedIdentityClientId(managedIdentityClientId);
}

TokenCredential credential = builder.build();
```

### Explicit local/Azure selection

For stricter production behavior, select the credential deliberately:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureCliCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

public final class CredentialFactory {
    public static TokenCredential create() {
        String environment = requireEnv("APP_ENVIRONMENT");

        if ("azure".equalsIgnoreCase(environment)) {
            String clientId =
                System.getenv("MANAGED_IDENTITY_CLIENT_ID");

            ManagedIdentityCredentialBuilder builder =
                new ManagedIdentityCredentialBuilder();

            if (clientId != null && !clientId.isBlank()) {
                builder.clientId(clientId);
            }

            return builder.build();
        }

        if ("local".equalsIgnoreCase(environment)) {
            return new AzureCliCredentialBuilder().build();
        }

        throw new IllegalStateException(
            "APP_ENVIRONMENT must be 'azure' or 'local'");
    }

    private static String requireEnv(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }
}
```

This avoids accidentally falling back to a developer credential in production.

Other local options include:

- `AzureCliCredential` after an interactive `az login`.
- IDE credentials supported by `DefaultAzureCredential`.
- `EnvironmentCredential` with service-principal environment variables for CI.
- Workload identity federation for CI/CD or Kubernetes, avoiding client secrets.

Do not create a fake managed identity endpoint or put production secrets in source code.

## 7. Error handling

Authentication failures and authorization failures are different:

```java
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.identity.CredentialUnavailableException;

try {
    blobServiceClient.listBlobContainers()
        .forEach(container -> System.out.println(container.getName()));
} catch (CredentialUnavailableException e) {
    System.err.println(
        "Managed Identity is unavailable in this environment: "
            + e.getMessage());
    throw e;
} catch (HttpResponseException e) {
    int status = e.getResponse().getStatusCode();

    if (status == 401) {
        System.err.println(
            "Token rejected; check tenant, audience, and identity configuration.");
    } else if (status == 403) {
        System.err.println(
            "Authenticated but unauthorized; check Azure RBAC assignments.");
    } else {
        System.err.println(
            "Azure request failed with HTTP " + status + ": "
                + e.getMessage());
    }

    throw e;
}
```

Also expect `ClientAuthenticationException` when token acquisition is attempted but fails:

```java
import com.azure.core.exception.ClientAuthenticationException;

try {
    // Make an SDK request. Most clients acquire tokens lazily here.
} catch (CredentialUnavailableException e) {
    // No supported Managed Identity endpoint was available.
    throw e;
} catch (ClientAuthenticationException e) {
    // An identity endpoint existed, but authentication failed.
    throw e;
}
```

Avoid swallowing these exceptions or silently switching identities. A fallback that changes identity can hide production configuration problems.

## 8. Troubleshooting checklist

| Symptom | Likely cause | Check |
|---|---|---|
| Credential unavailable locally | No Managed Identity endpoint exists | Use `DefaultAzureCredential` or an explicit developer credential |
| Credential unavailable in Azure | Managed Identity not enabled or unsupported hosting setup | Verify identity is enabled and attached to the workload |
| User-assigned identity not found | Wrong identity selector | Confirm client ID versus object ID versus resource ID |
| HTTP 401 | Invalid token context or service endpoint configuration | Verify endpoint, Azure cloud, tenant, and requested service |
| HTTP 403 | Identity authenticated but lacks permission | Assign the correct RBAC role at the correct scope |
| Works after a delay | Role assignment propagation | Allow several minutes for Azure RBAC propagation |
| Wrong identity used | Multiple identities attached or credential chain selected another credential | Set the user-assigned client ID explicitly; simplify the credential chain |
| Timeout reaching identity endpoint | Network, proxy, container, or hosting configuration | Ensure local proxy settings do not intercept the platform identity endpoint |
| Key Vault access denied | Access model mismatch | Check whether the vault uses Azure RBAC or legacy access policies |
| Storage access denied | Management role used instead of data role | Assign a data-plane role such as Storage Blob Data Reader |

Enable Azure SDK logging during diagnosis, but avoid verbose identity logs in normal production operation:

```java
import com.azure.core.util.logging.ClientLogger;

ClientLogger logger = new ClientLogger(MyApplication.class);
logger.info("Starting Azure client operation");
```

SDK log verbosity is normally configured through your logging framework and environment. Treat diagnostic logs as potentially sensitive because they can contain tenant IDs, client IDs, endpoints, and request metadata.

A strong production pattern is:

- Use `ManagedIdentityCredential` explicitly in Azure when no fallback should occur.
- Select user-assigned identities by client ID when more than one identity may exist.
- Use `DefaultAzureCredential` or an explicit developer credential locally.
- Grant only the minimum required Azure RBAC roles.
- Distinguish token acquisition errors from HTTP `401`, `403`, and network failures.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The response does not leave any AutoCloseable Azure SDK client unclosed.

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
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 7.7% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.8333** | **83.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-managed-identity --config java-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
