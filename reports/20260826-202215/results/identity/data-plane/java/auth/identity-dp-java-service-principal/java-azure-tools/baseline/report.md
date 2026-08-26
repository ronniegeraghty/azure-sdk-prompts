# Evaluation Report: identity-dp-java-service-principal

**Config:** java-azure-tools/baseline | **Result:** ✅ PASSED | **Duration:** 377.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-service-principal` |
| Config | java-azure-tools/baseline |
| Result | ✅ PASSED |
| Score | 17/17 |
| Duration | 377.0s |
| Timestamp | 2026-08-26T15:48:47Z |
| Files Generated | 0 |
| Event Count | 3140 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 110.4s |
| Review | 265.7s |
| **Total** | **377.0s** |

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
| Token Usage | in=42502 out=1909 |
| Turn Count | 2 |

## Prompt Sent

```
Show me how to authenticate
to Azure using a Service Principal with client secret in Java. I need:
1. Required Maven dependencies
2. How to create a ClientSecretCredential with ClientSecretCredentialBuilder
3. Using it with an Azure SDK client
4. Best practices for secret management
5. Error handling for invalid credentials

Provide a complete example with proper exception handling.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Java Azure Identity ClientSecretCredentialBuilder Azure SDK service principal client secret Maven dependency 2026"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Azure SDK for Java dependency BOM azure-sdk-bom latest Maven 2026 azure-identity azure-storage-blob"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com Azure Java authentication exceptions ClientAuthenticationException invalid client secret error handling"}
```

## Copilot Response

I’ll verify the current Azure Java SDK dependency versions and authentication APIs against official documentation.

Use `ClientSecretCredential` from `azure-identity`, then pass it directly to an Azure SDK client builder. Authentication is lazy: invalid credentials normally fail on the first Azure service request, not when `build()` creates the credential.

### 1. Maven dependencies

Using the Azure SDK BOM keeps compatible library versions aligned:

```xml
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0
                             https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>azure-service-principal-example</artifactId>
    <version>1.0.0</version>

    <properties>
        <maven.compiler.release>17</maven.compiler.release>
        <azure-sdk-bom.version>1.3.8</azure-sdk-bom.version>
    </properties>

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
        <!-- ClientSecretCredential and other Azure credentials -->
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-identity</artifactId>
        </dependency>

        <!-- Example Azure SDK client -->
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-storage-blob</artifactId>
        </dependency>
    </dependencies>
</project>
```

### 2. Complete Java example

This example reads credentials from environment variables, creates a `ClientSecretCredential`, uses it with Blob Storage, and distinguishes authentication failures from service and configuration errors.

```java
package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.ClientSecretCredential;
import com.azure.identity.ClientSecretCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobStorageException;

public final class ServicePrincipalExample {

    private ServicePrincipalExample() {
    }

    public static void main(String[] args) {
        try {
            AzureConfig config = AzureConfig.fromEnvironment();

            ClientSecretCredential credential =
                    new ClientSecretCredentialBuilder()
                            .tenantId(config.tenantId())
                            .clientId(config.clientId())
                            .clientSecret(config.clientSecret())
                            .build();

            BlobServiceClient blobServiceClient =
                    new BlobServiceClientBuilder()
                            .endpoint(config.storageEndpoint())
                            .credential(credential)
                            .buildClient();

            // Authentication occurs when the SDK makes this request.
            blobServiceClient.listBlobContainers()
                    .stream()
                    .forEach(container ->
                            System.out.println(container.getName()));

        } catch (IllegalStateException e) {
            System.err.println("Configuration error: " + e.getMessage());
            System.exit(2);

        } catch (ClientAuthenticationException e) {
            System.err.println(
                    "Azure authentication failed. Verify the tenant ID, client ID, "
                            + "client-secret value, and secret expiration.");

            // Do not print the client secret or dump configuration values.
            if (e.getMessage() != null) {
                System.err.println("Azure response: " + sanitize(e.getMessage()));
            }
            System.exit(3);

        } catch (BlobStorageException e) {
            System.err.printf(
                    "Azure Storage request failed: status=%d, errorCode=%s%n",
                    e.getStatusCode(),
                    e.getErrorCode());

            /*
             * A 403 usually means authentication succeeded but the service
             * principal lacks an appropriate Azure RBAC role, such as
             * Storage Blob Data Reader.
             */
            System.exit(4);
        }
    }

    private static String sanitize(String message) {
        // Keep diagnostics single-line and avoid forwarding arbitrary line breaks.
        return message.replace('\r', ' ').replace('\n', ' ');
    }

    private record AzureConfig(
            String tenantId,
            String clientId,
            String clientSecret,
            String storageEndpoint) {

        static AzureConfig fromEnvironment() {
            return new AzureConfig(
                    requiredEnvironmentVariable("AZURE_TENANT_ID"),
                    requiredEnvironmentVariable("AZURE_CLIENT_ID"),
                    requiredEnvironmentVariable("AZURE_CLIENT_SECRET"),
                    requiredEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT"));
        }

        private static String requiredEnvironmentVariable(String name) {
            String value = System.getenv(name);
            if (value == null || value.isBlank()) {
                throw new IllegalStateException(
                        "Required environment variable " + name + " is not set.");
            }
            return value;
        }
    }
}
```

Set the variables before running it in PowerShell:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<account-name>.blob.core.windows.net"

mvn compile exec:java `
  "-Dexec.mainClass=com.example.ServicePrincipalExample"
```

Use the secret’s **value**, not its secret ID. Azure only displays the value when the secret is created.

### 3. Secret-management practices

- Never hard-code the secret or commit it to `application.properties`, source code, Dockerfiles, or scripts.
- For local development, use environment variables or a local secret store excluded from version control.
- In production, retrieve secrets from a protected deployment secret store and expose them only to the process that needs them.
- Prefer managed identity or workload identity when running on Azure. These avoid long-lived client secrets entirely.
- Grant the service principal only the required Azure RBAC roles. For this example, read-only listing generally requires an appropriate data-plane role such as **Storage Blob Data Reader**.
- Use short expiration periods, rotate secrets before expiration, and support overlapping old/new credentials during rotation.
- Never log the secret, environment dumps, authentication request bodies, or full credential objects.

Common credential errors include:

| Entra error | Meaning |
|---|---|
| `AADSTS7000215` | Invalid secret, frequently because the secret ID was supplied instead of its value |
| `AADSTS7000222` | Client secret expired |
| `AADSTS700016` | Application/client ID or tenant is incorrect |
| HTTP `403` from the service | Authentication may have succeeded, but the identity lacks authorization |

Official references: [ClientSecretCredentialBuilder](https://learn.microsoft.com/java/api/com.azure.identity.clientsecretcredentialbuilder), [Azure Identity for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme), and [service-principal authentication troubleshooting](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-service-principal).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The example does not leave any AutoCloseable Azure SDK client unclosed.

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
hyoka run --prompt-id identity-dp-java-service-principal --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
