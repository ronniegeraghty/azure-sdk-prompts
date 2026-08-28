# Evaluation Report: identity-dp-dotnet-service-principal

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ✅ PASSED | **Duration:** 312.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-service-principal` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 312.3s |
| Timestamp | 2026-08-27T15:15:11Z |
| Files Generated | 0 |
| Event Count | 2736 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 289.7s |
| Review | 21.4s |
| **Total** | **312.3s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=110300 out=2027 |
| Turn Count | 3 |

## Tool Calls

### ✅ 1. skill — 3719ms

**Input:**

```json
{"skill":"azure-storage"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Services

## Services

| Service | Use When | MCP Tools | CLI |
|---------|----------|-----------|-----|
| Blob Storage | Objects, files, backups, static content | `azure__storage` | `az storage blob` |
| File Shares | SMB file shares, lift-and-shift | - | `az storage file` |
| Queue Storage | Async messaging, task queues | - | `az storage queue` |
| Table Storage | NoSQL key-value (consider Cosmos DB) | - | `az storage table` |
| Data Lake | Big data analytics, hierarchical namespace | - | `az storage fs` |

## MCP Server (Preferred)

When Azure MCP is enabled:

- `azure__storage` with command `storage_account_list` - List storage accounts
- `azure__storage` with command `storage_container_list` - List containers in account
- `azure__storage` with command `storage_blob_list` - List blobs in container
- `azure__storage` with command `storage_blob_get` - Download blob content
- `azure__storage` with command `storage_blob_put` - Upload blob content

**If Azure MCP is not enabled:** Run `/azure:setup` or enable via `/mcp`.

## CLI Fallback

```bash
# List storage accounts
az storage account list --output table

# List containers
az storage container list --account-name ACCOUNT --output table

# List blobs
az storage blob list --account-name ACCOUNT --container-name CONTAINER --output table

# Download blob
az storage blob download --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH

# Upload blob
az storage blob upload --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH
```

## Storage Account Tiers

| Tier | Use Case | Performance |
|------|----------|-------------|
| Standard | General purpose, backup | Milliseconds |
| Premium | Databases, high IOPS | Sub-millisecond |

## Blob Access Tiers

| Tier | Access Frequency | Cost |
|------|-----------------|------|
| Hot | Frequent | Higher storage, lower access |
| Cool | Infrequent (30+ days) | Lower storage, higher ac
... (truncated)
```

### ✅ 2. view — 191269ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk-usage.md
@@ -1,115 +1,115 @@
 # Azure Storage SDK Usage
 
 SDK packages and quick start examples for Azure Storage services.
 
 ## Storage SDKs by Language
 
 | Language | Blob | Queue | File Share | Data Lake |
 |----------|------|-------|------------|----------|
 | .NET | `Azure.Storage.Blobs` | `Azure.Storage.Queues` | `Azure.Storage.Files.Shares` | `Azure.Storage.Files.DataLake` |
 | Java | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | JavaScript | `@azure/storage-blob` | `@azure/storage-queue` | `@azure/storage-file-share` | `@azure/storage-file-datalake` |
 | Python | `azure-storage-blob` | `azure-storage-queue` | `azure-storage-file-share` | `azure-storage-file-datalake` |
 | Go | `azblob` | `azqueue` | `azfile` | `azdatalake` |
 | Rust | `azure_storage_blob` | `azure_storage_queue` | - | - |
 
 ## Installation Commands
 
 | Language | Install Blob SDK + Identity |
 |----------|-----------------------------|
 | .NET | `dotnet add package Azure.Storage.Blobs` `dotnet add package Azure.Identity` |
 | Java | Maven: `com.azure:azure-storage-blob` `com.azure:azure-identity` |
 | JavaScript | `npm install @azure/storage-blob @azure/identity` |
 | Python | `pip install azure-storage-blob azure-identity` |
 | Go | `go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob github.com/Azure/azure-sdk-for-g
... (truncated)
```

### ✅ 3. view — 189387ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk-usage.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
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
 2. **Broad surface area** — checks environment variables, CLI tokens, and oth
... (truncated)
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 191274ms

**Input:**

```json
{"intent":"Provide a complete local-only C# example authenticating to Azure Blob Storage with ClientSecretCredential, including secure configuration and invalid-credential error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 191276ms

**Input:**

```json
{"intent":"Find current official Microsoft documentation for C# ClientSecretCredential, Azure.Storage.Blobs BlobServiceClient with TokenCredential, and Azure Identity error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

## Copilot Response

I’m grounding the example in current Azure Identity and Blob SDK guidance, including secret-handling and failure semantics.

Use `Azure.Identity` for `ClientSecretCredential` and `Azure.Storage.Blobs` for `BlobServiceClient`. Authentication occurs on the first token request or SDK operation—not when the credential object is constructed.

### 1. Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

### 2–3. Complete example

Set these environment variables before running:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<storage-account-name>"

dotnet run
```

> `AZURE_CLIENT_SECRET` must contain the **secret value**, not the secret's identifier.

`Program.cs`:

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Required environment variable '{name}' is not configured.");
    }

    return value;
}

try
{
    string tenantId = GetRequiredEnvironmentVariable("AZURE_TENANT_ID");
    string clientId = GetRequiredEnvironmentVariable("AZURE_CLIENT_ID");
    string clientSecret = GetRequiredEnvironmentVariable("AZURE_CLIENT_SECRET");
    string storageAccountName =
        GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

    var credential = new ClientSecretCredential(
        tenantId,
        clientId,
        clientSecret,
        new ClientSecretCredentialOptions
        {
            AuthorityHost = AzureAuthorityHosts.AzurePublicCloud
        });

    var serviceUri = new Uri(
        $"https://{storageAccountName}.blob.core.windows.net");

    var blobServiceClient = new BlobServiceClient(serviceUri, credential);

    using var cancellationSource = new CancellationTokenSource(
        TimeSpan.FromSeconds(30));

    // ClientSecretCredential authenticates lazily. Requesting a token here
    // separates credential failures from subsequent Storage failures.
    AccessToken token = await credential.GetTokenAsync(
        new TokenRequestContext(
            new[] { "https://storage.azure.com/.default" }),
        cancellationSource.Token);

    Console.WriteLine(
        $"Authentication succeeded. Token expires at {token.ExpiresOn:u}");

    // Example authenticated Blob Storage operation.
    await foreach (BlobContainerItem container in
        blobServiceClient.GetBlobContainersAsync(
            cancellationToken: cancellationSource.Token))
    {
        Console.WriteLine(container.Name);
    }
}
catch (AuthenticationFailedException exception)
{
    // Common causes: incorrect tenant/client ID, invalid or expired secret,
    // disabled service principal, or tenant policy restrictions.
    Console.Error.WriteLine(
        "Azure authentication failed. Verify the tenant ID, client ID, " +
        "and client-secret value.");

    // Avoid printing the client secret, access token, or full configuration.
    Console.Error.WriteLine(exception.Message);
    Environment.ExitCode = 1;
}
catch (RequestFailedException exception) when (exception.Status == 403)
{
    // Authentication succeeded, but the identity lacks data-plane access.
    Console.Error.WriteLine(
        "Authentication succeeded, but the service principal is not " +
        "authorized to access this storage account.");

    Console.Error.WriteLine(
        $"Azure error: {exception.ErrorCode}; request ID: " +
        $"{exception.ClientRequestId}");

    Environment.ExitCode = 2;
}
catch (RequestFailedException exception)
{
    Console.Error.WriteLine(
        $"Azure Storage request failed with HTTP {exception.Status}. " +
        $"Error: {exception.ErrorCode}; request ID: " +
        $"{exception.ClientRequestId}");

    Environment.ExitCode = 3;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine("The Azure operation timed out or was cancelled.");
    Environment.ExitCode = 4;
}
catch (InvalidOperationException exception)
{
    Console.Error.WriteLine($"Configuration error: {exception.Message}");
    Environment.ExitCode = 5;
}
catch (UriFormatException exception)
{
    Console.Error.WriteLine(
        $"The storage account produced an invalid endpoint: " +
        $"{exception.Message}");

    Environment.ExitCode = 6;
}
```

The service principal also needs an appropriate Azure RBAC **data-plane** role on the storage account or narrower scope. For this listing example, use **Storage Blob Data Reader**. For uploads and modifications, use **Storage Blob Data Contributor**.

### 4. Secret-storage guidance

| Environment | Recommended approach |
|---|---|
| Local development | .NET User Secrets or temporary environment variables |
| CI/CD | The pipeline's protected secret store, preferably workload identity federation |
| Azure-hosted production | Managed identity instead of a client secret |
| Non-Azure production requiring a secret | Secret manager such as Azure Key Vault, injected at runtime |

Environment variables are preferable to committing secrets in `appsettings.json`, but they are not a full secret-management system: process administrators and diagnostic tools may expose them. Never commit the secret, log it, include it in exception messages, or place it directly in command history.

For an ASP.NET Core or worker application, configuration can still read environment variables:

```csharp
string tenantId = configuration["AZURE_TENANT_ID"]
    ?? throw new InvalidOperationException("AZURE_TENANT_ID is missing");
```

Use `appsettings.json` only for non-sensitive settings such as the storage account name. If a client secret is unavoidable, rotate it regularly and assign the service principal only the minimum required RBAC role. For production Azure workloads, prefer `ManagedIdentityCredential` because it eliminates stored secrets entirely.

Official references: [ClientSecretCredential](https://learn.microsoft.com/dotnet/api/azure.identity.clientsecretcredential), [Azure Blob Storage client library for .NET](https://learn.microsoft.com/azure/storage/blobs/storage-blob-dotnet-get-started), and [Azure authentication best practices](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response provides a complete C# example covering packages, credential construction, SDK client integration, secure secret storage, and authentication error handling.

### Strengths

- Correctly notes that authentication is lazy and explicitly requests a token to distinguish credential failures from storage failures.
- Separates authentication, authorization, request, timeout, configuration, and URI errors.
- Provides practical secret-management and least-privilege RBAC guidance.

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - `Azure.Identity` package with `ClientSecretCredential` class: Pass
      - Constructor parameters: tenantId, clientId, clientSecret: Pass
      - Passing credential to Azure SDK clients: Pass
      - Environment variable patterns for secret storage: Pass
      - `AuthenticationFailedException` for invalid credentials: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-service-principal --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
