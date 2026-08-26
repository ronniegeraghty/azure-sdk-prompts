# Evaluation Report: identity-dp-java-default-credential

**Config:** java-azure-tools/baseline | **Result:** ✅ PASSED | **Duration:** 362.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-default-credential` |
| Config | java-azure-tools/baseline |
| Result | ✅ PASSED |
| Score | 17/17 |
| Duration | 362.4s |
| Timestamp | 2026-08-26T15:24:44Z |
| Files Generated | 0 |
| Event Count | 4389 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 131.7s |
| Review | 229.6s |
| **Total** | **362.4s** |

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
| Token Usage | in=112162 out=3179 |
| Turn Count | 4 |

## Prompt Sent

```
Show me how to authenticate
an Azure SDK client using DefaultAzureCredential in Java. Explain:
1. What Maven dependencies are needed
2. How to create and use a DefaultAzureCredential instance
3. The credential chain order and which credentials are tried
4. How it behaves differently in local development vs Azure environments
5. How to troubleshoot authentication failures

Provide a complete example that creates a SecretClient using DefaultAzureCredential.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Official Microsoft Azure Identity Java DefaultAzureCredential credential chain order 2026 Maven azure-identity documentation"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Official Microsoft Azure Key Vault Secrets Java Maven dependencies azure-security-keyvault-secrets latest version DefaultAzureCredential example 2026"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Official Microsoft troubleshoot DefaultAzureCredential Java Azure Identity logging credential unavailable authentication failed"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml"}
```

### 🔧 6. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":16000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/credential-chains"}
```

### 🔧 8. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/1.3.8/azure-sdk-bom-1.3.8.pom"}
```

### 🔧 9. web_fetch

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/azure/developer/java/sdk/troubleshooting-authentication-overview"}
```

### 🔧 10. web_fetch

**Input:**

```json
{"max_length":10000,"url":"https://learn.microsoft.com/en-us/azure/key-vault/secrets/quick-create-java"}
```

## Copilot Response

`DefaultAzureCredential` provides one authentication API that works with developer sign-ins locally and managed identities or workload identities in Azure. It checks an ordered credential chain and uses the first credential that can obtain a Microsoft Entra access token.

## 1. Maven dependencies

Using the Azure SDK Bill of Materials keeps Azure library versions compatible. As of August 2026, the current BOM is `1.3.8`:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>1.3.8</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <!-- DefaultAzureCredential and other TokenCredential implementations -->
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-identity</artifactId>
    </dependency>

    <!-- SecretClient -->
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies>
```

BOM `1.3.8` selects `azure-identity` `1.18.4` and Key Vault Secrets `4.11.1`. If you do not use the BOM, current stable releases are:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.18.5</version>
</dependency>

<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-secrets</artifactId>
    <version>4.11.2</version>
</dependency>
```

Prefer the BOM over independently versioning Azure packages.

## 2. Creating and using `DefaultAzureCredential`

Create one credential and supply it to the SDK client's builder:

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder().build();

SecretClient client = new SecretClientBuilder()
    .vaultUrl("https://my-vault.vault.azure.net")
    .credential(credential)
    .buildClient();
```

Building the credential does not immediately authenticate. A token is normally requested when the first service operation runs, and the SDK handles token caching and refresh afterward. Reuse clients instead of building one for every request.

For a user-assigned managed identity, specify its client ID:

```java
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId(System.getenv("AZURE_CLIENT_ID"))
    .build();
```

## 3. Credential chain order

The default Java chain currently tries:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal configuration in environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS workload identity |
| 3 | `ManagedIdentityCredential` | System- or user-assigned managed identity on an Azure host |
| 4 | `IntelliJCredential` | Azure Toolkit for IntelliJ sign-in |
| 5 | `VisualStudioCodeCredential` | VS Code Azure sign-in; requires `azure-identity-broker` |
| 6 | `AzureCliCredential` | Current `az login` account |
| 7 | `AzurePowerShellCredential` | Current `Connect-AzAccount` account |
| 8 | `AzureDeveloperCliCredential` | Current `azd auth login` account |
| 9 | Brokered `InteractiveBrowserCredential` | OS account broker; requires `azure-identity-broker` |

The first credential that successfully obtains a token wins. That does **not** guarantee the identity has permission to access the requested resource.

`EnvironmentCredential` commonly uses:

```text
AZURE_TENANT_ID
AZURE_CLIENT_ID
AZURE_CLIENT_SECRET
```

For certificate authentication, it can use `AZURE_CLIENT_CERTIFICATE_PATH` instead of a client secret.

Starting with recent `azure-identity` versions, `AZURE_TOKEN_CREDENTIALS` can constrain the chain:

```text
AZURE_TOKEN_CREDENTIALS=dev
AZURE_TOKEN_CREDENTIALS=prod
AZURE_TOKEN_CREDENTIALS=AzureCliCredential
AZURE_TOKEN_CREDENTIALS=ManagedIdentityCredential
```

`dev` selects developer-tool credentials; `prod` selects deployed-service credentials. Restricting the production chain improves predictability and avoids unnecessary attempts.

## 4. Local development versus Azure

**Local development:** Environment, workload identity, and managed identity are usually unavailable, so the chain proceeds to developer credentials. A common setup is signing in with Azure CLI:

```powershell
az login
```

The application then uses the same identity through `AzureCliCredential`. That account must have data-plane permission on the vault, such as **Key Vault Secrets User** for reading secrets. Being an Azure subscription owner does not automatically grant Key Vault data access in every authorization configuration.

**Azure environments:** Configure a managed identity on App Service, Functions, Container Apps, or a VM, or use workload identity on AKS. No client secret is stored in the application. The same Java code resolves `ManagedIdentityCredential` or `WorkloadIdentityCredential` before reaching developer credentials.

For production, consider:

```text
AZURE_TOKEN_CREDENTIALS=prod
```

For a user-assigned managed identity, also provide its client ID through `managedIdentityClientId(...)`. The deployed identity must be granted the required Key Vault RBAC role or access-policy permissions.

## 5. Complete `SecretClient` example

```java
package example;

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
        String secretName = args.length > 0 ? args[0] : "example-secret";

        DefaultAzureCredential credential =
            new DefaultAzureCredentialBuilder().build();

        SecretClient secretClient = new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(credential)
            .buildClient();

        try {
            KeyVaultSecret secret = secretClient.getSecret(secretName);

            System.out.printf(
                "Retrieved secret '%s', version '%s'.%n",
                secret.getName(),
                secret.getProperties().getVersion());

            // Avoid printing secret.getValue() or writing it to logs.
        } catch (HttpResponseException exception) {
            System.err.printf(
                "Key Vault request failed with HTTP %d: %s%n",
                exception.getResponse().getStatusCode(),
                exception.getMessage());
            throw exception;
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is missing: " + name);
        }
        return value;
    }
}
```

Set the vault URL before running:

```powershell
$Env:KEY_VAULT_URL = "https://my-vault.vault.azure.net"
mvn compile exec:java -Dexec.mainClass="example.KeyVaultExample" `
    -Dexec.args="example-secret"
```

## Troubleshooting failures

1. **Inspect the complete exception.** `ClientAuthenticationException` means token acquisition failed. The message normally identifies each attempted credential and may include an `AADSTS` code, correlation ID, and timestamp.

2. **Distinguish authentication from authorization.**
   - Token acquisition failure: credential configuration or Microsoft Entra authentication problem.
   - HTTP `401`: invalid token, wrong tenant/audience, or authentication configuration.
   - HTTP `403`: authentication succeeded, but that identity lacks Key Vault permission.
   - HTTP `404`: the vault or secret name may be wrong; Key Vault configurations can also conceal unauthorized resources.

3. **Enable Azure Identity logging.** Azure SDK for Java uses SLF4J. Configure `com.azure.identity` at `DEBUG` or `INFO`. With `slf4j-simple`, for example:

   ```text
   -Dorg.slf4j.simpleLogger.log.com.azure.identity=debug
   ```

   Do not enable unsafe support logging in normal environments because it can include tokens and other sensitive identity data.

4. **Check which identity won.** A globally configured `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_SECRET` can cause `EnvironmentCredential` to win unexpectedly. Locally, confirm the active CLI account with `az account show`. In production, use `AZURE_TOKEN_CREDENTIALS=prod` or a specific credential name.

5. **Check environment-specific configuration.**
   - Service principal: verify tenant ID, client ID, and the **secret value**, not the secret's identifier.
   - Managed identity: verify it is enabled and assigned to the Azure host.
   - User-assigned identity: verify the correct client ID was supplied.
   - AKS workload identity: verify the service account annotation, federated identity, token file, tenant ID, and client ID.
   - Developer tools: verify the relevant CLI or IDE session is signed in to the intended tenant.

6. **Check network access.** The application must reach Microsoft Entra token endpoints and the vault. Private endpoints, DNS, proxies, firewalls, and blocked managed-identity endpoints can all produce authentication-like failures.

References: [Java credential chains](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), [Azure Identity troubleshooting](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview), and [Key Vault Secrets Java client](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

The response satisfies the client resource-management criterion.

### Strengths

- Uses the modern SecretClientBuilder pattern without introducing an unclosed AutoCloseable client.

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
hyoka run --prompt-id identity-dp-java-default-credential --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
