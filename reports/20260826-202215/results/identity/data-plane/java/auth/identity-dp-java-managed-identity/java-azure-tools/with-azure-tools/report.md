# Evaluation Report: identity-dp-java-managed-identity

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 338.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-managed-identity` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 16/18 |
| Duration | 338.6s |
| Timestamp | 2026-08-26T15:43:08Z |
| Files Generated | 0 |
| Event Count | 4541 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 118.6s |
| Review | 218.8s |
| **Total** | **338.6s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=361841 out=4163 |
| Turn Count | 8 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 2. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide Java examples and guidance for authenticating Azure SDK clients with system-assigned and user-assigned managed identities, including local development fallback and troubleshooting.","learn":true}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Microsoft documentation for Azure Identity Java ManagedIdentityCredential and DefaultAzureCredential, covering system-assigned and user-assigned identities, local development, error handling, and troubleshooting.","learn":true}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for Azure Identity authentication in Java using managed identities and local developer credentials.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official guidance for managed identity authentication with Azure SDK for Java.","parameters":{"query":"Azure Identity Java ManagedIdentityCredential system-assigned user-assigned DefaultAzureCredential local development troubleshooting"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java code samples for system-assigned and user-assigned managed identities and Azure SDK clients.","parameters":{"language":"java","query":"ManagedIdentityCredentialBuilder system assigned user assigned clientId resourceId DefaultAzureCredentialBuilder Java"}}
```

### 🔧 8. rg

**Input:**

```json
{"head_limit":10,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787759040179-copilot-tool-output-7a5d182411a9465b9286c0e7736eef8e.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

### 🔧 9. rg

**Input:**

```json
{"head_limit":10,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787759039490-copilot-tool-output-628ca03e1b3c48fd8e46a20b084c3896.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\ ]+"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787759040179-copilot-tool-output-7a5d182411a9465b9286c0e7736eef8e.txt' | ConvertFrom-Json; $j.results | Select-Object -First 10 title,contentUrl | Format-Table -AutoSize","description":"List documentation result URLs"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"$j = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787759039490-copilot-tool-output-628ca03e1b3c48fd8e46a20b084c3896.txt' | ConvertFrom-Json; $j.results | Select-Object -First 10 description,url | Format-List","description":"List code sample result URLs"}
```

### 🔧 12. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787759040179-copilot-tool-output-7a5d182411a9465b9286c0e7736eef8e.txt","view_range":[1,8]}
```

### 🔧 13. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787759039490-copilot-tool-output-628ca03e1b3c48fd8e46a20b084c3896.txt","view_range":[1,8]}
```

### 🔧 14. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Retrieve complete official system-assigned managed identity guidance for Java.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/system-assigned-managed-identity"}}
```

### 🔧 15. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Retrieve complete official user-assigned managed identity guidance for Java.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/user-assigned-managed-identity"}}
```

### 🔧 16. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Retrieve complete official Java Azure Identity troubleshooting guidance for Azure-hosted applications.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-azure-hosted"}}
```

### 🔧 17. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Retrieve complete official local development authentication guidance for Azure SDK for Java.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-dev-accounts"}}
```

## Copilot Response

I’m grounding the examples in the current Java SDK credential patterns and production guidance.

Managed identities let Azure-hosted Java applications obtain Microsoft Entra tokens without storing secrets. Enable the identity on the hosting resource, assign it least-privilege Azure RBAC roles, then pass a `ManagedIdentityCredential` to any Azure SDK client that accepts `TokenCredential`.

## 1. System-assigned vs. user-assigned

| Type | Lifecycle | Typical use |
|---|---|---|
| **System-assigned** | Created on and deleted with one Azure resource | One workload needs its own identity |
| **User-assigned** | Independent Azure resource; can be attached to multiple workloads | Several resources share permissions, or identity lifecycle must remain independent |

If several identities are attached to a host, explicitly select the intended user-assigned identity.

## Dependencies

Using the Azure SDK BOM is recommended so dependency versions remain compatible:

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

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies>
```

Without the BOM, the current stable Azure Identity documentation references `azure-identity` **1.18.5**.

## 2. Create `ManagedIdentityCredential`

### System-assigned identity

No identifier is required because the identity belongs to the current Azure host:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;

TokenCredential credential =
    new ManagedIdentityCredentialBuilder().build();
```

### User-assigned identity by client ID

The client ID is generally the most portable selector. Keep it in configuration—not source code:

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

Do not confuse the identity's **client ID** with its principal/object ID.

You can alternatively select it by Azure resource ID:

```java
TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .resourceId(
        "/subscriptions/<subscription-id>"
        + "/resourceGroups/<resource-group>"
        + "/providers/Microsoft.ManagedIdentity"
        + "/userAssignedIdentities/<identity-name>")
    .build();
```

The builder also supports `.objectId(...)`, but client ID or resource ID is usually clearer.

## 3. Use the credential with Azure SDK clients

All supported data-plane clients accept the same `TokenCredential`.

### Blob Storage using a system-assigned identity

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

TokenCredential credential =
    new ManagedIdentityCredentialBuilder().build();

BlobServiceClient blobServiceClient = new BlobServiceClientBuilder()
    .endpoint("https://<storage-account>.blob.core.windows.net")
    .credential(credential)
    .buildClient();

blobServiceClient.listBlobContainers()
    .forEach(container -> System.out.println(container.getName()));
```

The managed identity needs an appropriate **data-plane** role, such as `Storage Blob Data Reader` or `Storage Blob Data Contributor`. Subscription `Reader` or `Owner` alone does not necessarily provide blob data access.

### Key Vault using a user-assigned identity

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

String managedIdentityClientId =
    System.getenv("MANAGED_IDENTITY_CLIENT_ID");

if (managedIdentityClientId == null || managedIdentityClientId.isBlank()) {
    throw new IllegalStateException(
        "MANAGED_IDENTITY_CLIENT_ID must be configured");
}

TokenCredential credential = new ManagedIdentityCredentialBuilder()
    .clientId(managedIdentityClientId)
    .build();

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();

String value = secretClient.getSecret("database-password").getValue();
```

The identity needs a role such as `Key Vault Secrets User`, scoped as narrowly as practical.

Build and reuse credentials and clients as long-lived objects. The SDK handles token acquisition, caching, and refresh.

## 4. Local development fallbacks

Managed identity endpoints exist only in supported Azure hosting environments. A `ManagedIdentityCredential` normally cannot authenticate on a developer workstation.

### Recommended: explicit production and development credentials

Use managed identity in Azure and a developer credential locally:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

final class CredentialFactory {
    private CredentialFactory() {
    }

    static TokenCredential create() {
        boolean azureHosted =
            Boolean.parseBoolean(System.getenv("AZURE_HOSTED"));

        if (!azureHosted) {
            return new DefaultAzureCredentialBuilder().build();
        }

        String clientId = System.getenv("MANAGED_IDENTITY_CLIENT_ID");

        if (clientId == null || clientId.isBlank()) {
            // System-assigned identity.
            return new ManagedIdentityCredentialBuilder().build();
        }

        // User-assigned identity.
        return new ManagedIdentityCredentialBuilder()
            .clientId(clientId)
            .build();
    }
}
```

Locally, `DefaultAzureCredential` can use credentials from Azure CLI, Azure Developer CLI, Azure PowerShell, IntelliJ, or Visual Studio Code. For example, authenticate your local tooling with `az login`; the application then runs as your developer identity.

For more deterministic local behavior, use the exact tool credential:

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureCliCredentialBuilder;

TokenCredential credential =
    new AzureCliCredentialBuilder().build();
```

Important considerations:

- Give the local developer identity equivalent **least-privilege** permissions.
- Avoid falling back to developer credentials in production. An explicit `ManagedIdentityCredential` prevents an unexpected environment or tooling credential from being selected.
- Avoid setting `AZURE_CLIENT_SECRET` for local convenience unless you intentionally want `DefaultAzureCredential` to use a service principal.
- If using `DefaultAzureCredential` with a user-assigned identity in both environments, configure it with `.managedIdentityClientId(clientId)`. This is convenient, but less deterministic in production than using `ManagedIdentityCredential` directly.

## 5. Error handling and troubleshooting

Authentication commonly happens on the first service request rather than when the client is constructed:

```java
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.CredentialUnavailableException;

try {
    secretClient.getSecret("database-password");
} catch (CredentialUnavailableException exception) {
    System.err.println(
        "No configured credential was available: "
        + exception.getMessage());
    throw exception;
} catch (ClientAuthenticationException exception) {
    System.err.println(
        "Microsoft Entra token acquisition failed: "
        + exception.getMessage());
    throw exception;
} catch (HttpResponseException exception) {
    int status = exception.getResponse().getStatusCode();

    if (status == 401) {
        System.err.println(
            "The service rejected the access token.");
    } else if (status == 403) {
        System.err.println(
            "Authentication succeeded, but the identity lacks permission.");
    }

    throw exception;
}
```

Common failure checks:

| Symptom | Likely cause |
|---|---|
| `ManagedIdentityCredential authentication unavailable` | Managed identity isn't enabled, or the app isn't running on a supported Azure host |
| `The requested identity hasn't been assigned to this resource` | Wrong client ID/resource ID, or the user-assigned identity isn't attached to the host |
| HTTP `401` | Invalid token audience, tenant/configuration problem, or service rejected authentication |
| HTTP `403` | Token acquisition succeeded, but the identity lacks the required service-specific RBAC role |
| Works locally but fails in Azure | Developer account has broader permissions than the managed identity |
| VM managed identity times out | IMDS endpoint `169.254.169.254` is blocked or routed through a proxy |
| Recently assigned role still fails | Azure RBAC changes may require time to propagate |

Enable SDK diagnostics with `AZURE_LOG_LEVEL=verbose`. Do not log access tokens, client secrets, or managed identity endpoint responses.

**References:**

- [System-assigned managed identity with Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/system-assigned-managed-identity)
- [User-assigned managed identity with Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/user-assigned-managed-identity)
- [Local development authentication](https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-dev-accounts)
- [Java credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)
- [Troubleshoot Azure-hosted Java authentication](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-azure-hosted)

## Reviewer Notes (LLM-as-Judge)

### Summary

The sole criterion passed.

### Strengths

- Correctly avoids unnecessary try-with-resources blocks for Azure SDK clients that do not implement AutoCloseable.

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
hyoka run --prompt-id identity-dp-java-managed-identity --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
