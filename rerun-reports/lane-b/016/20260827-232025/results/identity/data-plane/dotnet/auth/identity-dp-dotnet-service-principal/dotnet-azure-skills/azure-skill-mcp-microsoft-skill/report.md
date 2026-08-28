# Evaluation Report: identity-dp-dotnet-service-principal

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 297.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-service-principal` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 297.0s |
| Timestamp | 2026-08-27T15:20:25Z |
| Files Generated | 0 |
| Event Count | 2914 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 273.0s |
| Review | 22.8s |
| **Total** | **297.0s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=184320 out=2400 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. skill — 4346ms

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for .NET

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```bash
dotnet add package Azure.Identity

# For ASP.NET Core integration
dotnet add package Microsoft.Extensions.Azure

# For brokered authentication and Visual Studio Code credential support
dotnet add package Azure.Identity.Broker
```

## Environment Variables

### Service Principal with Secret

```bash
AZURE_CLIENT_ID=<application-client-id>
AZURE_TENANT_ID=<directory-tenant-id>
AZURE_CLIENT_SECRET=<client-secret-value>
```

### Service Principal with Certificate

```bash
AZURE_CLIENT_ID=<application-client-id>
AZURE_TENANT_ID=<directory-tenant-id>
AZURE_CLIENT_CERTIFICATE_PATH=<path-to-pfx-or-pem>
AZURE_CLIENT_CERTIFICATE_PASSWORD=<certificate-password>  # Optional
```

### Managed Identity

```bash
AZURE_CLIENT_ID=<user-assigned-managed-identity-client-id>  # Only for user-assigned
```

## DefaultAzureCredential

The recommended credential for most scenarios. Tries multiple authentication methods in order. See [DefaultAzureCredential overview](https://aka.ms/azsdk/net/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

### Basic Usage

```csharp
using Azure.Identity;
using Azure.Storage.Blobs;

var credential = new DefaultAzureCredential();
var blobClient = new BlobServiceClient(
    new Uri("https://myaccount.blob.core.windows.net"),
    credential);
```

### ASP.NET Core with Dependency Injection

```csharp
using Azure.Identity;
using Microsoft.Extensions.Azure;

builder.Services.AddAzureClients(clientBuilder =>
{
    clientBuilder.AddBlobServiceClient(
        new Uri("https://myaccount.blob.core.windows.net"));
    clientBuilder.AddSecretClient(
        new Uri("https://myvault.vault.azure.net"));
    
    // Uses DefaultAzureCredential by default
    clientBuilder.UseCredential(new DefaultAzureCredential());
});
```

### Customizing DefaultAz
... (truncated)
```

### ❌ 2. view — 5520ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 5518ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. view — 3380ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-identity-dotnet b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-identity-dotnet
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-identity-dotnet
@@ -1,0 +1,1 @@
 SKILL.md

```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 185395ms

**Input:**

```json
{"intent":"Generate a local-only C# example using Azure.Identity ClientSecretCredential with Azure.Storage.Blobs, environment-based secret storage, and robust authentication/authorization error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. azure-documentation (via azure) — 185397ms

**Input:**

```json
{"intent":"Find official documentation for Azure.Identity ClientSecretCredential, BlobServiceClient token credentials, authentication exceptions, and secure secret storage best practices in .NET.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

## Copilot Response

I’ll ground the example in the Azure Identity skill’s acceptance criteria and storage-client patterns.

## 1. Required NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

- `Azure.Identity` provides `ClientSecretCredential`.
- `Azure.Storage.Blobs` provides `BlobServiceClient`.

## 2. Configuration

Set credentials outside source control:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
```

The service principal also needs an appropriate data-plane role on the storage account, such as **Storage Blob Data Reader** or **Storage Blob Data Contributor**.

## 3. Complete example

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

try
{
    string tenantId = GetRequiredEnvironmentVariable("AZURE_TENANT_ID");
    string clientId = GetRequiredEnvironmentVariable("AZURE_CLIENT_ID");
    string clientSecret = GetRequiredEnvironmentVariable("AZURE_CLIENT_SECRET");
    string storageAccountUrl =
        GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_URL");

    if (!Uri.TryCreate(storageAccountUrl, UriKind.Absolute, out Uri? serviceUri) ||
        serviceUri.Scheme != Uri.UriSchemeHttps)
    {
        throw new InvalidOperationException(
            "AZURE_STORAGE_ACCOUNT_URL must be a valid HTTPS URL.");
    }

    var credentialOptions = new ClientSecretCredentialOptions
    {
        Retry =
        {
            MaxRetries = 3,
            Delay = TimeSpan.FromSeconds(1),
            MaxDelay = TimeSpan.FromSeconds(5)
        }
    };

    TokenCredential credential = new ClientSecretCredential(
        tenantId: tenantId,
        clientId: clientId,
        clientSecret: clientSecret,
        options: credentialOptions);

    var blobServiceClient = new BlobServiceClient(serviceUri, credential);

    using var timeout = new CancellationTokenSource(TimeSpan.FromSeconds(30));

    Console.WriteLine("Authenticated. Listing blob containers:");

    await foreach (BlobContainerItem container in
        blobServiceClient.GetBlobContainersAsync(
            cancellationToken: timeout.Token))
    {
        Console.WriteLine($"- {container.Name}");
    }

    Console.WriteLine("Request completed successfully.");
}
catch (CredentialUnavailableException ex)
{
    // More common with chained credentials, but safe to handle explicitly.
    Console.Error.WriteLine(
        $"No usable Azure credential was available: {ex.Message}");
    Environment.ExitCode = 1;
}
catch (AuthenticationFailedException ex)
{
    // Examples: invalid tenant ID, client ID, expired secret, or wrong secret.
    Console.Error.WriteLine(
        "Azure authentication failed. Verify AZURE_TENANT_ID, " +
        "AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET.");

    // Do not log the secret or credential object.
    Console.Error.WriteLine($"Details: {ex.Message}");
    Environment.ExitCode = 2;
}
catch (RequestFailedException ex) when (ex.Status is 401 or 403)
{
    // Authentication may have succeeded, but Azure rejected the request.
    Console.Error.WriteLine(
        $"Azure rejected the storage request with HTTP {ex.Status}.");

    Console.Error.WriteLine(
        ex.Status == 403
            ? "Verify that the service principal has an appropriate Storage Blob Data role."
            : "Verify the credential and storage endpoint.");

    Console.Error.WriteLine(
        $"Azure error code: {ex.ErrorCode ?? "not provided"}");

    Environment.ExitCode = 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure Storage request failed with HTTP {ex.Status}. " +
        $"Error code: {ex.ErrorCode ?? "not provided"}. " +
        $"Message: {ex.Message}");

    Environment.ExitCode = 4;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine("The Azure Storage request timed out.");
    Environment.ExitCode = 5;
}
catch (InvalidOperationException ex)
{
    Console.Error.WriteLine($"Configuration error: {ex.Message}");
    Environment.ExitCode = 6;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Required environment variable '{name}' is not set.");
    }

    return value;
}
```

`ClientSecretCredential` acquires tokens lazily. Therefore, invalid credentials generally surface when `GetBlobContainersAsync` sends its first request, not when the credential object is constructed.

## 4. Secret-storage best practices

**Environment variables:** Better than hard-coding or committing secrets in `appsettings.json`. They work well in CI/CD systems and container platforms, but can still be exposed through process inspection, diagnostic dumps, or overly broad deployment access.

**.NET configuration:** `IConfiguration` is only an abstraction—it does not make a secret secure. This is acceptable:

```csharp
string secret = configuration["AzureAd:ClientSecret"]
    ?? throw new InvalidOperationException("Client secret is missing.");
```

But security depends on the underlying provider:

- Use environment variables or .NET User Secrets for local development.
- Use the hosting platform’s protected secret settings or a dedicated secret store in production.
- Never commit secrets to `appsettings.json`, source code, Dockerfiles, or repository history.
- Never log the secret.
- Rotate secrets regularly and keep their lifetime short.
- Prefer a certificate or workload identity over a client secret where supported.
- For applications hosted in Azure, **managed identity is preferable** because it eliminates stored credentials.
- Reuse a single credential and SDK client; Azure SDK clients are thread-safe.

For ASP.NET Core environment-variable configuration, hierarchical keys use double underscores:

```powershell
$env:AzureAd__TenantId = "<tenant-id>"
$env:AzureAd__ClientId = "<client-id>"
$env:AzureAd__ClientSecret = "<client-secret-value>"
```

Then access them as `AzureAd:TenantId`, `AzureAd:ClientId`, and `AzureAd:ClientSecret` through `IConfiguration`.

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response provides a complete C# example with the required packages, credential construction, Azure SDK client integration, secret-storage practices, and appropriate authentication error handling.

### Strengths

- Correctly notes that authentication is lazy and failures typically occur on the first service request.
- Distinguishes authentication failures from authorization and other Azure Storage request failures.
- Avoids hard-coded secrets and warns against logging or committing them.

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
hyoka run --prompt-id identity-dp-dotnet-service-principal --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
