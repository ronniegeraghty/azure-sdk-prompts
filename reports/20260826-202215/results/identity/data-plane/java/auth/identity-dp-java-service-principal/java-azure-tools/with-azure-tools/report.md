# Evaluation Report: identity-dp-java-service-principal

**Config:** java-azure-tools/with-azure-tools | **Result:** ✅ PASSED | **Duration:** 360.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-service-principal` |
| Config | java-azure-tools/with-azure-tools |
| Result | ✅ PASSED |
| Score | 17/17 |
| Duration | 360.8s |
| Timestamp | 2026-08-26T12:30:47Z |
| Files Generated | 0 |
| Event Count | 3576 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 109.1s |
| Review | 250.6s |
| **Total** | **360.8s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-tools/with-azure-tools

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=325900 out=4312 |
| Turn Count | 7 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide a local-only Java example that authenticates to Azure with a service principal client secret, builds ClientSecretCredential, uses it with Azure Blob Storage SDK, handles invalid credentials, and follows secure secret-management practices.","learn":true}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Microsoft documentation for Java ClientSecretCredentialBuilder, Azure BlobServiceClientBuilder token credential authentication, ClientAuthenticationException handling, and service principal client secret security best practices.","learn":true}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java"}
```

### 🔧 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Java Azure authentication code-generation guidance for a service principal client secret example using Blob Storage.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official guidance for service principal client-secret authentication in Java and secure secret management.","parameters":{"query":"Azure Java authenticate service principal client secret ClientSecretCredentialBuilder BlobServiceClientBuilder ClientAuthenticationException secret management"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java samples for ClientSecretCredentialBuilder used with Azure Blob Storage and authentication error handling.","parameters":{"language":"java","query":"ClientSecretCredentialBuilder BlobServiceClientBuilder ClientAuthenticationException service principal environment variables"}}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787747488007-copilot-tool-output-fdd8ebca7b3b4d2eae9c3ed2a4b41107.txt' | ConvertFrom-Json; $d.results | Select-Object -First 6 title,content,url | ConvertTo-Json -Depth 4","description":"Extract relevant Microsoft docs results"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787747488675-copilot-tool-output-3bd3e0c928714d5c88ece4a97170b0de.txt' | ConvertFrom-Json; $d.results | Select-Object -First 6 description,codeSnippet,link | ConvertTo-Json -Depth 4","description":"Extract official Java code samples"}
```

### 🔧 11. web_fetch

**Input:**

```json
{"max_length":4000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-sdk-bom%22\u0026rows=1\u0026wt=json"}
```

### 🔧 12. web_fetch

**Input:**

```json
{"max_length":4000,"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-storage-blob%22\u0026rows=1\u0026wt=json"}
```

### 🔧 13. web_fetch

**Input:**

```json
{"max_length":4000,"url":"https://search.maven.org/solrsearch/select?q=g:%22org.slf4j%22%20AND%20a:%22slf4j-simple%22\u0026rows=1\u0026wt=json"}
```

### 🔧 14. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787747488007-copilot-tool-output-fdd8ebca7b3b4d2eae9c3ed2a4b41107.txt","view_range":[1,12]}
```

### 🔧 15. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787747488675-copilot-tool-output-3bd3e0c928714d5c88ece4a97170b0de.txt","view_range":[1,12]}
```

### 🔧 16. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/service-principal"}
```

### 🔧 17. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-service-principal"}
```

### 🔧 18. web_fetch

**Input:**

```json
{"max_length":5000,"raw":true,"url":"https://repo.maven.apache.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

### 🔧 19. web_fetch

**Input:**

```json
{"max_length":5000,"raw":true,"url":"https://repo.maven.apache.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

### 🔧 20. web_fetch

**Input:**

```json
{"max_length":5000,"raw":true,"url":"https://repo.maven.apache.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
```

## Copilot Response

I’ll ground the example in the current Azure Identity patterns and use a concrete Blob Storage client so the credential wiring and failures are explicit.

The SDK pattern is straightforward; I’m checking the official troubleshooting guidance and current artifact versions so the dependency block and exception handling aren’t stale.

Use `ClientSecretCredential` for an explicit service-principal flow. The client secret should come from environment variables or a secret manager—never source code.

### 1. Maven dependencies

The Azure SDK BOM keeps Azure library versions compatible. This example uses the current stable BOM `1.3.8`.

```xml
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
    <!-- ClientSecretCredential and other Microsoft Entra credentials -->
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
```

`azure-core`, including `ClientAuthenticationException`, is included transitively.

### 2. Configuration

Supply these values through your process environment:

```text
AZURE_TENANT_ID=<Microsoft-Entra-directory-tenant-ID>
AZURE_CLIENT_ID=<application-client-ID>
AZURE_CLIENT_SECRET=<client-secret-value>
AZURE_STORAGE_BLOB_ENDPOINT=https://<account-name>.blob.core.windows.net
```

The service principal also needs an appropriate data-plane role on the storage account, such as **Storage Blob Data Reader** for this read-only example.

### 3. Complete Java example

```java
package example;

import com.azure.core.exception.AzureException;
import com.azure.core.exception.ClientAuthenticationException;
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
            BlobServiceClient client = createBlobServiceClient();

            // Building the client is lazy. This service request triggers token
            // acquisition and therefore verifies the credentials.
            client.listBlobContainers()
                .stream()
                .limit(10)
                .forEach(container ->
                    System.out.println(container.getName()));

        } catch (CredentialUnavailableException e) {
            System.err.println(
                "The service-principal credential is unavailable. "
                    + "Check the authentication configuration.");
            System.exit(2);

        } catch (ClientAuthenticationException e) {
            reportAuthenticationFailure(e);
            System.exit(3);

        } catch (BlobStorageException e) {
            if (e.getStatusCode() == 403) {
                System.err.println(
                    "Authentication succeeded, but the service principal is not "
                        + "authorized for this storage operation. Check its Azure RBAC role.");
            } else {
                System.err.printf(
                    "Azure Blob Storage rejected the request. Status=%d, errorCode=%s%n",
                    e.getStatusCode(),
                    e.getErrorCode());
            }
            System.exit(4);

        } catch (AzureException e) {
            // Covers other Azure SDK failures, such as transport or service errors.
            System.err.println(
                "The Azure SDK operation failed. Check network connectivity, "
                    + "the service endpoint, and Azure service health.");
            System.exit(5);

        } catch (IllegalStateException e) {
            // Raised by requireEnvironmentVariable before any secret is used.
            System.err.println("Configuration error: " + e.getMessage());
            System.exit(6);
        }
    }

    private static BlobServiceClient createBlobServiceClient() {
        String tenantId = requireEnvironmentVariable("AZURE_TENANT_ID");
        String clientId = requireEnvironmentVariable("AZURE_CLIENT_ID");
        String clientSecret = requireEnvironmentVariable("AZURE_CLIENT_SECRET");
        String storageEndpoint =
            requireEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT");

        ClientSecretCredential credential =
            new ClientSecretCredentialBuilder()
                .tenantId(tenantId)
                .clientId(clientId)
                .clientSecret(clientSecret)
                .build();

        return new BlobServiceClientBuilder()
            .endpoint(storageEndpoint)
            .credential(credential)
            .buildClient();
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable " + name + " is not set.");
        }

        return value;
    }

    private static void reportAuthenticationFailure(
        ClientAuthenticationException exception) {

        String message = exception.getMessage();

        if (message != null && message.contains("AADSTS7000215")) {
            System.err.println(
                "Authentication failed: the client secret is invalid. "
                    + "Verify that AZURE_CLIENT_SECRET contains the secret value, "
                    + "not the secret ID.");
        } else if (message != null && message.contains("AADSTS7000222")) {
            System.err.println(
                "Authentication failed: the client secret has expired. "
                    + "Rotate it and update the secret store.");
        } else if (message != null && message.contains("AADSTS700016")) {
            System.err.println(
                "Authentication failed: the application was not found in the tenant. "
                    + "Verify AZURE_CLIENT_ID and AZURE_TENANT_ID.");
        } else {
            // Avoid printing the full exception because logs may be broadly accessible.
            System.err.println(
                "Microsoft Entra authentication failed. Verify the tenant ID, "
                    + "client ID, secret, authority host, and application status.");
        }
    }
}
```

`ClientSecretCredentialBuilder.build()` does **not** contact Microsoft Entra ID. Authentication occurs when the first SDK operation requests a token, which is why the exception handling surrounds `listBlobContainers()`.

### 4. Secret-management practices

- **Prefer managed identity or workload identity in Azure-hosted production applications.** They eliminate stored client secrets.
- When a secret is necessary, load it from a protected CI/CD secret store, orchestrator secret, or another secure runtime injection mechanism.
- Never commit secrets to source control, `application.properties`, container images, command-line arguments, or `.env` files tracked by Git.
- Grant the service principal only the required roles and scope them to the specific resource whenever possible.
- Use short secret lifetimes, rotate before expiration, and support overlapping credentials during rotation.
- Use separate service principals for development, testing, and production.
- Do not log the client secret, access tokens, environment dumps, or full HTTP request headers.
- Prefer a certificate credential over a client secret when managed identity or workload identity cannot be used.

Official references:

- [Authenticate Java apps using service principals](https://learn.microsoft.com/azure/developer/java/sdk/authentication/local-development-service-principal)
- [Troubleshoot service-principal authentication](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-service-principal)
- [Azure Blob Storage Java SDK](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java)

## Reviewer Notes (LLM-as-Judge)

### Summary

The client resource-management criterion is fully met.

### Strengths

- Uses modern Azure SDK client and credential types that do not require explicit closure.

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
hyoka run --prompt-id identity-dp-java-service-principal --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
