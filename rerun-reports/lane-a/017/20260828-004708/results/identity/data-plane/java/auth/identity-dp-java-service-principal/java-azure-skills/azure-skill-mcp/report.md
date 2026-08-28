# Evaluation Report: identity-dp-java-service-principal

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 745.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-service-principal` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 745.2s |
| Timestamp | 2026-08-27T16:47:08Z |
| Files Generated | 0 |
| Event Count | 3090 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 438.3s |
| Review | 245.7s |
| **Total** | **745.2s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | entra-app-registration |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=152025 out=2979 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. skill — 2808ms

**Input:**

```json
{"skill":"entra-app-registration"}
```

**Output:**

```
Skill loaded successfully ✅

## Overview

Microsoft Entra ID (formerly Azure Active Directory) is Microsoft's cloud-based identity and access management service. App registrations allow applications to authenticate users and access Azure resources securely.

### Key Concepts

| Concept | Description |
|---------|-------------|
| **App Registration** | Configuration that allows an app to use Microsoft identity platform |
| **Application (Client) ID** | Unique identifier for your application |
| **Tenant ID** | Unique identifier for your Azure AD tenant/directory |
| **Client Secret** | Password for the application (confidential clients only) |
| **Redirect URI** | URL where authentication responses are sent |
| **API Permissions** | Access scopes your app requests |
| **Service Principal** | Identity created in your tenant when you register an app |

### Application Types

| Type | Use Case |
|------|----------|
| **Web Application** | Server-side apps, APIs |
| **Single Page App (SPA)** | JavaScript/React/Angular apps |
| **Mobile/Native App** | Desktop, mobile apps |
| **Daemon/Service** | Background services, APIs |

## Core Workflow

### Step 1: Register the Application

Create an app registration in the Azure portal or using Azure CLI.

**Portal Method:**
1. Navigate to Azure Portal → Microsoft Entra ID → App registrations
2. Click "New registration"
3. Provide name, supported account types, and redirect URI
4. Click "Register"

**CLI Method:** See [references/cli-commands.md](references/cli-commands.md)
**IaC Method:** See [references/BICEP-EXAMPLE.bicep](references/BICEP-EXAMPLE.bicep)

It's highly recommended to use the IaC to manage Entra app registration if you already use IaC in your project, need a scalable solution for managing lots of app registrations or need fine-grained audit history of the configuration changes. 

### Step 2: Configure Authentication

Set up authentication settings based on your application type.

- **Web Apps**: Add redirect 
... (truncated)
```

### ✅ 2. view — 190557ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\sdk\\azure-identity-java.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/sdk/azure-identity-java.md
@@ -1,31 +1,31 @@
 # Authentication — Java SDK Quick Reference
 
 > Condensed from **azure-identity-java**. Full patterns (workload identity,
 > certificate auth, device code, sovereign clouds)
 > in the **azure-identity-java** plugin skill if installed.
 
 ## Install
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.15.0</version>
 </dependency>
 ```
 
 ## Quick Start
 > **Auth:** `DefaultAzureCredential` is for local development. See [auth-best-practices.md](../auth-best-practices.md) for production patterns.
 
 ```java
 import com.azure.identity.DefaultAzureCredentialBuilder;
 var credential = new DefaultAzureCredentialBuilder().build();
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** (CLI, PowerShell, VS Code). In production, use ManagedIdentityCredential — see [auth-best-practices.md](../auth-best-practices.md)
 - Managed identity in production — no secrets to manage, automatic rotation
 - Azure CLI for local dev — run `az login` before running your app
 - Least privilege — grant only required permissions to service principals
 - Token caching — enabled by default, reduces auth round-trips
 - Environment variables — use for CI/CD, not hardcoded se
... (truncated)
```

### ✅ 3. view — 189282ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\entra-app-registration\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/entra-app-registration/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks enviro
... (truncated)
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 190558ms

**Input:**

```json
{"intent":"Generate a Java example authenticating to Azure with a service principal client secret via ClientSecretCredentialBuilder, using an Azure SDK client, secure secret management, and robust invalid-credential error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 190561ms

**Input:**

```json
{"intent":"Find current official Java Azure Identity documentation for ClientSecretCredentialBuilder, Maven dependencies, use with an Azure SDK client, and handling authentication failures.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. web_search (via github-mcp-server) — 183407ms

**Input:**

```json
{"query":"What are the latest stable Maven Central versions as of August 2026 for com.azure azure-sdk-bom, azure-identity, and azure-storage-blob? Cite Maven Central or official Microsoft sources."}
```

**Output:**

```
{"type":"output_text","text":{"value":"Based on the latest official and Maven Central sources as of August 2026, here are the most recent stable (non-beta) versions:\n\n1. **com.azure:azure-sdk-bom**  \n   - **Latest stable version:** 1.3.8  \n   - Release date: July 23, 2026  \n   - Source: Maven Central directory for this artifact shows `azure-sdk-bom-1.3.8.pom` with a 2026-07-23 timestamp【3:6†source】.\n\n2. **com.azure:azure-identity**  \n   - **Latest stable version:** 1.18.5  \n   - Released: August 24, 2026  \n   - Source: Official Microsoft GitHub releases for azure-sdk-for-java shows version 1.18.5 released on 2026-08-24【3:7†source】.\n\n3. **com.azure:azure-storage-blob**  \n   - **Latest stable version:** 12.35.0  \n   - Note: There is a `12.36.0-beta.1`, but 12.35.0 is the most recent stable (i.e., non-beta) release【3:3†source】.\n\n**Official References:**  \n- [Maven Central: azure-sdk-bom](https://search.maven.org/artifact/com.azure/azure-sdk-bom)【3:6†source】\n- [Maven Central: azure-identity](https://search.maven.org/artifact/com.azure/azure-identity)【3:7†source】\n- [Maven Central: azure-storage-blob](https://search.maven.org/artifact/com.azure/azure-storage-blob)【3:3†source】\n\nThese versions are confirmed as current and stable as of late August 2026.","annotations":[{"text":"","start_index":350,"end_index":362,"url_citation":{"title":"Central Repository: com/azure/azure-sdk-bom/1.3.8","url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/1.3.8/"}},{"text":"","start_index":587,"end_index":599,"url_citation":{"title":"Releases: Azure/azure-sdk-for-java - GitHub","url":"https://github.com/Azure/azure-sdk-for-java/releases"}},{"text":"","start_index":784,"end_index":796,"url_citation":{"title":"Maven Central: com.azure:azure-storage-blob","url":"https://central.sonatype.com/artifact/com.azure/azure-storage-blob"}},{"text":"","start_index":917,"end_index":929,"url_citation":{"title":"Central Repository: com/azure
... (truncated)
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 183410ms

**Input:**

```json
{"intent":"Retrieve Azure Java SDK authentication and secret-management best practices for service principal client-secret authentication examples.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 8. azure-documentation (via azure) — 183408ms

**Input:**

```json
{"intent":"Retrieve official documentation for Azure Identity Java ClientSecretCredentialBuilder and ClientAuthenticationException, including lazy authentication behavior and service-client use.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

## Copilot Response

Use `ClientSecretCredential` for the OAuth 2.0 client-credentials flow. The example below authenticates a service principal and lists Azure Blob Storage containers.

### 1. Maven dependencies

Use the Azure SDK BOM to keep Azure library versions compatible:

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

The versions shown are current stable releases as of August 2026. The BOM manages the individual dependency versions.

### 2. Environment configuration

Do not put credentials in source code, `pom.xml`, command-line arguments, or committed configuration files.

PowerShell:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<account-name>.blob.core.windows.net"
```

The service principal needs an appropriate data-plane role, such as **Storage Blob Data Reader**, scoped as narrowly as possible.

### 3. Complete Java example

```java
package com.example;

import com.azure.core.exception.AzureException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.credential.TokenCredential;
import com.azure.identity.ClientAuthenticationException;
import com.azure.identity.ClientSecretCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobContainerItem;

public final class ServicePrincipalExample {
    private ServicePrincipalExample() {
    }

    public static void main(String[] args) {
        int exitCode = run();

        if (exitCode != 0) {
            System.exit(exitCode);
        }
    }

    private static int run() {
        final String tenantId;
        final String clientId;
        final String clientSecret;
        final String blobEndpoint;

        try {
            tenantId = requireEnvironmentVariable("AZURE_TENANT_ID");
            clientId = requireEnvironmentVariable("AZURE_CLIENT_ID");
            clientSecret = requireEnvironmentVariable("AZURE_CLIENT_SECRET");
            blobEndpoint = requireEnvironmentVariable(
                "AZURE_STORAGE_BLOB_ENDPOINT"
            );
        } catch (IllegalStateException exception) {
            System.err.println("Configuration error: " + exception.getMessage());
            return 2;
        }

        TokenCredential credential =
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

        try {
            // Authentication is lazy. The first SDK operation requests a token.
            System.out.println("Containers:");

            for (BlobContainerItem container
                    : blobServiceClient.listBlobContainers()) {
                System.out.println("- " + container.getName());
            }

            return 0;
        } catch (ClientAuthenticationException exception) {
            // Typical causes: invalid/expired secret, incorrect tenant ID,
            // incorrect client ID, or a disabled/deleted service principal.
            System.err.println(
                "Microsoft Entra authentication failed. Verify the tenant ID, "
                    + "client ID, and client secret, and confirm that the "
                    + "service principal is enabled."
            );
            System.err.println(
                "Authentication error: " + exception.getMessage()
            );

            // Never print clientSecret or the complete credential configuration.
            return 3;
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();

            if (statusCode == 403) {
                System.err.println(
                    "Authentication succeeded, but the service principal is "
                        + "not authorized. Check its Azure RBAC role and scope."
                );
            } else {
                System.err.printf(
                    "Azure Storage returned HTTP %d: %s%n",
                    statusCode,
                    exception.getMessage()
                );
            }

            return 4;
        } catch (AzureException exception) {
            // Covers other Azure SDK failures such as transport problems.
            System.err.println(
                "Azure SDK operation failed: " + exception.getMessage()
            );
            return 5;
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable " + name + " is not set"
            );
        }

        return value;
    }
}
```

`ClientSecretCredentialBuilder.build()` does not immediately contact Microsoft Entra ID. Authentication occurs when the Azure SDK client first requests a token, so the SDK operation—and any lazy iteration it returns—must remain inside the `try` block.

### 4. Error behavior

| Failure | Expected exception/result |
|---|---|
| Wrong or expired client secret | `ClientAuthenticationException` |
| Wrong tenant or client ID | `ClientAuthenticationException` |
| Valid identity but missing RBAC role | `HttpResponseException`, usually HTTP 403 |
| Incorrect endpoint or service failure | `HttpResponseException` or another `AzureException` |
| Missing environment configuration | Local `IllegalStateException` before authentication |

Avoid logging the exception request payload, access tokens, or credential values. Authentication messages are useful for diagnostics, but production logs should be access-controlled.

### 5. Secret-management best practices

- Prefer **managed identity** for applications hosted in Azure; it eliminates client secrets.
- Prefer **workload identity federation** for Kubernetes and CI/CD environments.
- If a secret is unavoidable, inject it at runtime from a protected secret store rather than committing it to source control.
- Store secrets in Azure Key Vault or the platform’s native secret manager, but bootstrap Key Vault access with managed identity or federation—not the same secret being retrieved.
- Grant the service principal only the required RBAC role at the narrowest resource scope.
- Set an expiration date, rotate before expiration, and support overlapping credentials during rotation.
- Never include secrets in command-line arguments, exception messages, telemetry, test fixtures, container images, or build artifacts.
- Treat client secrets as transitional credentials; certificates or federated credentials are preferable for non-Azure production workloads.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses the Azure SDK client without incorrectly applying resource-management requirements to non-AutoCloseable types.

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.0000** | **92.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-service-principal --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
