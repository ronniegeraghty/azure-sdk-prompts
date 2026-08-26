# Evaluation Report: identity-dp-java-default-credential

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 410.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-default-credential` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 410.6s |
| Timestamp | 2026-08-26T15:30:47Z |
| Files Generated | 0 |
| Event Count | 5028 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 160.9s |
| Review | 248.5s |
| **Total** | **410.6s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-java, azure-security-keyvault-secrets-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=550299 out=5075 |
| Turn Count | 10 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-security-keyvault-secrets-java"}
```

### 🔧 3. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java\\references\\examples.md"}
```

### 🔧 5. glob

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java","pattern":"**/acceptance-criteria.md"}
```

### 🔧 6. glob

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java","pattern":"**/acceptance-criteria.md"}
```

### 🔧 7. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide a complete, modern local-only Java example using DefaultAzureCredential with Azure Key Vault SecretClient, including dependencies, credential chain behavior, environment differences, and troubleshooting.","learn":true}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official Microsoft documentation for DefaultAzureCredential in the Azure Identity Java SDK: Maven dependency/BOM guidance, exact credential chain order and continuation policy, AZURE_TOKEN_CREDENTIALS behavior, local development versus managed identity/workload identity on Azure, troubleshooting/logging, and creating Key Vault SecretClient.","learn":true}
```

### 🔧 9. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for a Java application authenticating with DefaultAzureCredential and reading Azure Key Vault secrets.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official DefaultAzureCredential chain and troubleshooting details for Java.","parameters":{"query":"Azure Identity Java DefaultAzureCredential exact credential chain order continuation policy AZURE_TOKEN_CREDENTIALS troubleshooting logging managed identity local development"}}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Java Key Vault SecretClient dependency and authentication example.","parameters":{"query":"Azure Key Vault Secrets Java SecretClient DefaultAzureCredential Maven azure sdk bom dependency"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve the latest official Java code sample for DefaultAzureCredential with SecretClient.","parameters":{"language":"java","query":"DefaultAzureCredentialBuilder SecretClientBuilder Key Vault SecretClient complete example Maven"}}
```

### 🔧 13. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787758314411-copilot-tool-output-9d87437a1d2b4ef588ac134035f0e87a.txt' | ConvertFrom-Json; $d.results | Select-Object title, contentUrl | Format-Table -AutoSize -Wrap","description":"List identity documentation search results"}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787758314322-copilot-tool-output-64d4b14b6664435d913baff88159f198.txt' | ConvertFrom-Json; $d.results | Select-Object title, contentUrl | Format-Table -AutoSize -Wrap","description":"List Key Vault documentation results"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$d = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787758313506-copilot-tool-output-cf1b54496ecf4f77915a3dc72d781b33.txt' | ConvertFrom-Json; $d.results | Select-Object description, sourceUrl | Format-List","description":"List official Java sample sources"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787758314411-copilot-tool-output-9d87437a1d2b4ef588ac134035f0e87a.txt","view_range":[1,5]}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787758314322-copilot-tool-output-64d4b14b6664435d913baff88159f198.txt","view_range":[1,5]}
```

### 🔧 18. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787758313506-copilot-tool-output-cf1b54496ecf4f77915a3dc72d781b33.txt","view_range":[1,5]}
```

### 🔧 19. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the complete current Java DefaultAzureCredential chain documentation.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains"}}
```

### 🔧 20. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch current Azure Identity Java README for versions and continuation policy.","parameters":{"url":"https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable"}}
```

### 🔧 21. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch current troubleshooting guidance for Azure Identity Java.","parameters":{"url":"https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview"}}
```

### 🔧 22. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch current Key Vault Secrets Java package and SecretClient guidance.","parameters":{"url":"https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable"}}
```

### 🔧 23. web_search (via github-mcp-server)

**Input:**

```json
{"query":"What are the latest stable versions of Maven artifacts com.azure:azure-sdk-bom, com.azure:azure-identity, and com.azure:azure-security-keyvault-secrets as of August 2026? Prefer Maven Central and official Microsoft sources; exclude beta versions."}
```

### 🔧 24. web_fetch

**Input:**

```json
{"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-sdk-bom%22\u0026rows=1\u0026wt=json"}
```

### 🔧 25. web_fetch

**Input:**

```json
{"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-identity%22\u0026rows=1\u0026wt=json"}
```

### 🔧 26. web_fetch

**Input:**

```json
{"url":"https://search.maven.org/solrsearch/select?q=g:%22com.azure%22%20AND%20a:%22azure-security-keyvault-secrets%22\u0026rows=1\u0026wt=json"}
```

## Copilot Response

I’m grounding the example in the current Java SDK guidance, especially the credential-chain behavior, which can vary by `azure-identity` version.

`DefaultAzureCredential` lets the same Java code authenticate with developer credentials locally and workload/managed identity after deployment to Azure.

## 1. Maven dependencies

Use the Azure Identity and Key Vault Secrets libraries. These are the current GA versions shown in Microsoft’s Java SDK documentation:

```xml
<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-identity</artifactId>
        <version>1.18.5</version>
    </dependency>

    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
        <version>4.11.0</version>
    </dependency>
</dependencies>
```

For a larger Azure application, prefer the [Azure SDK BOM](https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/boms/azure-sdk-bom) so Azure dependency versions remain compatible:

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
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
</dependencies>
```

Use the latest GA BOM version listed in the [Azure SDK releases](https://azure.github.io/azure-sdk/releases/latest/java.html). If you want Visual Studio Code or OS-account broker authentication, also add `com.azure:azure-identity-broker`.

## 2. Complete `SecretClient` example

Configuration is supplied through environment variables—no vault URL, credentials, or secret values are hardcoded:

```java
package example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultExample {
    private KeyVaultExample() {
    }

    public static void main(String[] args) {
        String vaultUrl = requiredEnvironmentVariable("AZURE_KEY_VAULT_URL");
        String secretName = requiredEnvironmentVariable("SECRET_NAME");

        DefaultAzureCredential credential =
            new DefaultAzureCredentialBuilder().build();

        SecretClient secretClient = new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(credential)
            .buildClient();

        try {
            // The first service operation triggers token acquisition.
            KeyVaultSecret secret = secretClient.getSecret(secretName);

            // Use the value without writing it to logs.
            String secretValue = secret.getValue();
            useSecret(secretValue);

            System.out.printf(
                "Retrieved secret '%s', version '%s'.%n",
                secret.getName(),
                secret.getProperties().getVersion());
        } catch (CredentialUnavailableException exception) {
            System.err.println(
                "No credential in DefaultAzureCredential was available: "
                    + exception.getMessage());
        } catch (ClientAuthenticationException exception) {
            System.err.println(
                "Microsoft Entra authentication failed: "
                    + exception.getMessage());
        } catch (HttpResponseException exception) {
            int status = exception.getResponse().getStatusCode();
            System.err.printf(
                "Key Vault request failed with HTTP %d: %s%n",
                status,
                exception.getMessage());
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }

    private static void useSecret(String value) {
        // Pass the secret to the component that needs it. Do not log it.
        if (value.isEmpty()) {
            throw new IllegalStateException("The retrieved secret is empty.");
        }
    }
}
```

Example configuration:

```text
AZURE_KEY_VAULT_URL=https://my-vault.vault.azure.net
SECRET_NAME=database-password
```

Constructing the credential does not immediately authenticate. Authentication normally occurs when `getSecret()` makes the first request; acquired access tokens are then cached and refreshed by the credential.

## 3. Default credential chain order

With `azure-identity` 1.18.5, `DefaultAzureCredential` attempts these credentials in order and stops when one obtains a token:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal variables such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS workload identity |
| 3 | `ManagedIdentityCredential` | System- or user-assigned Azure managed identity |
| 4 | `IntelliJCredential` | Azure Toolkit for IntelliJ login |
| 5 | `VisualStudioCodeCredential` | VS Code Azure Resources login; requires `azure-identity-broker` |
| 6 | `AzureCliCredential` | Account authenticated through `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated through `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated through `azd auth login` |
| 9 | Broker credential | OS-brokered account; requires `azure-identity-broker` |

The chain stops at the first credential that **acquires a token**, even if that identity later lacks permission to access Key Vault. It does not retry another identity after Key Vault returns `403 Forbidden`.

Since `azure-identity` 1.10.0, failures from developer credentials do not stop the chain; all developer credentials are attempted until one succeeds. A deployed-service credential that is configured and able to attempt authentication, but receives an authentication failure, stops the chain rather than falling through to developer credentials.

## 4. Local development versus Azure

**Local development:** Managed identity and workload identity are normally unavailable, so the chain continues to IntelliJ, VS Code, Azure CLI, Azure PowerShell, or Azure Developer CLI. A common setup is:

```powershell
az login
```

The application then uses the signed-in Azure CLI account. That user must have Key Vault data-plane permission—for example, the **Key Vault Secrets User** role when using Azure RBAC.

**Azure-hosted application:** Enable a managed identity on App Service, Functions, Container Apps, a VM, or another supported host. The same code discovers `ManagedIdentityCredential`, without storing a client secret.

For a user-assigned managed identity, specify its client ID:

```java
DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder()
        .managedIdentityClientId(
            System.getenv("AZURE_CLIENT_ID"))
        .build();
```

For production, consider making credential selection deterministic:

```text
AZURE_TOKEN_CREDENTIALS=prod
```

This restricts `DefaultAzureCredential` to deployed-service credentials. Alternatively, set it to `ManagedIdentityCredential` with `azure-identity` 1.17.0 or later. For fail-fast configuration validation in 1.18.0 or later:

```java
import com.azure.identity.AzureIdentityEnvVars;

DefaultAzureCredential credential =
    new DefaultAzureCredentialBuilder()
        .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
        .build();
```

Microsoft also recommends using `ManagedIdentityCredential` directly in Azure production applications when managed identity is the only intended mechanism. That avoids unexpected credential selection and reduces chain-probing overhead.

## 5. Troubleshooting failures

1. **Inspect the full exception.** `ClientAuthenticationException` messages commonly include the failing credential, Microsoft Entra error code, correlation ID, and timestamp.

2. **Enable Azure Identity logging.** Configure your application’s SLF4J implementation to log `com.azure.identity` at `DEBUG` or `TRACE`. The output shows each attempted credential and identifies which one returned a token. Avoid unsafe MSAL logging unless working with support—it can contain tokens and personal data.

3. **Distinguish authentication from authorization.**
   - No credential available: local tool is not logged in, environment variables are incomplete, or managed/workload identity is unavailable.
   - `401`: token, tenant, audience, or authentication configuration problem.
   - `403`: authentication usually succeeded, but the selected identity lacks Key Vault data-plane access, the vault firewall blocks the request, or private networking/DNS is misconfigured.
   - `404`: vault URL, secret name, or secret version is incorrect.

4. **Check which identity won the chain.** Environment variables are checked before local developer credentials and managed identity. Stale `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, or `AZURE_CLIENT_SECRET` variables can cause an unintended service principal to be selected.

5. **Verify tenant and account.** Confirm the CLI/IDE is signed into the expected tenant and subscription. A token can be valid while belonging to an identity that has no access to the vault.

6. **Verify Key Vault authorization.** For retrieval with Azure RBAC, grant the authenticated identity a suitable data-plane role such as **Key Vault Secrets User**, scoped as narrowly as practical. Subscription `Owner` alone does not necessarily grant secret data access.

7. **Check user-assigned managed identity selection.** When more than one user-assigned identity exists, configure `managedIdentityClientId()` or `AZURE_CLIENT_ID` explicitly.

**References:** [DefaultAzureCredential chain](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains), [Azure Identity for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable), [authentication troubleshooting](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-authentication-overview), and [Key Vault Secrets Java SDK](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- The sample does not leave any AutoCloseable Azure SDK client unclosed.

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.0000** | **92.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-default-credential --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
