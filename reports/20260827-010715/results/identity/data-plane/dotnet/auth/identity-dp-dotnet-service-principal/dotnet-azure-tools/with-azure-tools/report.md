# Evaluation Report: identity-dp-dotnet-service-principal

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ✅ PASSED | **Duration:** 149.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-service-principal` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 149.1s |
| Timestamp | 2026-08-26T17:43:44Z |
| Files Generated | 0 |
| Event Count | 3232 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 120.9s |
| Review | 27.0s |
| **Total** | **149.1s** |

## Configuration

- **name:** dotnet-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=341867 out=3838 |
| Turn Count | 8 |

## Prompt Sent

```
Show me how to authenticate
to Azure using a Service Principal with client secret in C#. I need:
1. Required NuGet packages
2. How to create a ClientSecretCredential with tenant ID, client ID, and client secret
3. How to use it with an Azure SDK client (e.g., BlobServiceClient)
4. Best practices for storing the secret (environment variables vs configuration)
5. Error handling when credentials are invalid

Provide a complete example with proper error handling.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get","intent":"Get current best practices for C# Azure SDK authentication using a service principal client secret, secure secret configuration, BlobServiceClient usage, and authentication error handling."}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"command":"search","intent":"Find official Microsoft documentation for ClientSecretCredential in C#, Azure BlobServiceClient with TokenCredential, environment-variable secret storage, RBAC requirements, and handling AuthenticationFailedException and RequestFailedException."}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\acceptance-criteria.md"}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\examples.md"}
```

### 🔧 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get current best practices for C# Azure SDK authentication using a service principal client secret, secure secret configuration, BlobServiceClient usage, and authentication error handling.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Microsoft documentation for service principal client-secret authentication in Azure SDK for .NET.","parameters":{"query":"Azure SDK for .NET ClientSecretCredential tenant ID client ID client secret BlobServiceClient authentication error handling environment variables RBAC"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# samples for ClientSecretCredential and BlobServiceClient.","parameters":{"language":"csharp","query":"ClientSecretCredential BlobServiceClient service principal client secret AuthenticationFailedException"}}
```

### 🔧 9. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet"}
```

### 🔧 10. rg

**Input:**

```json
{"head_limit":30,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766281000-copilot-tool-output-8d6cec1732cf480f9876a88fe4f1cc78.txt"],"pattern":"\"(title|url)\":\"[^\"]+"}
```

### 🔧 11. rg

**Input:**

```json
{"head_limit":20,"output_mode":"content","paths":["C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766280675-copilot-tool-output-a7eab67b6ca14b939455203b9ff432e9.txt"],"pattern":"\"(title|url|description)\":\"[^\"]+"}
```

### 🔧 12. powershell

**Input:**

```json
{"command":"$paths = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766281000-copilot-tool-output-8d6cec1732cf480f9876a88fe4f1cc78.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766280675-copilot-tool-output-a7eab67b6ca14b939455203b9ff432e9.txt'); foreach ($path in $paths) { $json = Get-Content -Raw $path | ConvertFrom-Json; $json.results | Select-Object -First 10 | ForEach-Object { [pscustomobject]@{ Title = $_.title; Url = if ($_.contentUrl) {$_.contentUrl} elseif ($_.url) {$_.url} else {$_.source}; Description = $_.description } } | Format-List }","description":"Extract titles and URLs from docs results"}
```

### 🔧 13. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766281000-copilot-tool-output-8d6cec1732cf480f9876a88fe4f1cc78.txt","view_range":[1,8]}
```

### 🔧 14. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766280675-copilot-tool-output-a7eab67b6ca14b939455203b9ff432e9.txt","view_range":[1,8]}
```

### 🔧 15. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official guidance for ClientSecretCredential authentication in .NET apps.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/authentication/on-premises-apps"}}
```

### 🔧 16. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official Azure Identity error-handling and thread-safety guidance.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet"}}
```

## Copilot Response

I’m grounding the sample in the current Azure Identity guidance and the skill’s acceptance criteria before assembling the complete code.

## Required packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

- `Azure.Identity` provides `ClientSecretCredential`.
- `Azure.Storage.Blobs` provides `BlobServiceClient`.
- `Microsoft.Extensions.Azure` is optional if using ASP.NET Core dependency injection.

## Configuration

Set these values in the process environment:

```powershell
$env:AZURE_TENANT_ID = "<directory-tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_BLOB_ENDPOINT = "https://<storage-account>.blob.core.windows.net"
```

Use the **client secret value**, not the secret's ID. The service principal also needs an appropriate data-plane role, such as **Storage Blob Data Reader**, scoped as narrowly as possible.

## Complete example

```csharp
using Azure;
using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;

return await RunAsync();

static async Task<int> RunAsync()
{
    try
    {
        string tenantId = GetRequiredEnvironmentVariable("AZURE_TENANT_ID");
        string clientId = GetRequiredEnvironmentVariable("AZURE_CLIENT_ID");
        string clientSecret = GetRequiredEnvironmentVariable("AZURE_CLIENT_SECRET");
        Uri blobEndpoint = GetRequiredHttpsUri("AZURE_STORAGE_BLOB_ENDPOINT");

        var credential = new ClientSecretCredential(
            tenantId: tenantId,
            clientId: clientId,
            clientSecret: clientSecret);

        var clientOptions = new BlobClientOptions
        {
            Retry =
            {
                Mode = RetryMode.Exponential,
                MaxRetries = 3,
                Delay = TimeSpan.FromSeconds(1),
                MaxDelay = TimeSpan.FromSeconds(8),
                NetworkTimeout = TimeSpan.FromSeconds(30)
            }
        };

        var blobServiceClient = new BlobServiceClient(
            blobEndpoint,
            credential,
            clientOptions);

        // Authentication is lazy. The first service operation requests a token
        // and is therefore where invalid credentials are normally reported.
        Console.WriteLine("Containers:");

        await foreach (BlobContainerItem container in
            blobServiceClient.GetBlobContainersAsync())
        {
            Console.WriteLine($"- {container.Name}");
        }

        return 0;
    }
    catch (CredentialUnavailableException ex)
    {
        Console.Error.WriteLine(
            $"The configured credential could not be used: {ex.Message}");
        return 1;
    }
    catch (AuthenticationFailedException ex)
    {
        Console.Error.WriteLine(
            $"Microsoft Entra authentication failed: {ex.Message}");
        Console.Error.WriteLine(
            "Check the tenant ID, client ID, client secret value, and secret expiration.");
        return 2;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        Console.Error.WriteLine(
            $"Authentication succeeded, but access was denied: {ex.Message}");
        Console.Error.WriteLine(
            "Check the service principal's Azure Storage data-plane RBAC assignments.");
        return 3;
    }
    catch (RequestFailedException ex)
    {
        Console.Error.WriteLine(
            $"Azure Storage request failed. Status={ex.Status}, " +
            $"ErrorCode={ex.ErrorCode}, Message={ex.Message}");
        return 4;
    }
    catch (InvalidOperationException ex)
    {
        Console.Error.WriteLine($"Configuration error: {ex.Message}");
        return 5;
    }
    catch (UriFormatException ex)
    {
        Console.Error.WriteLine($"Invalid Blob Storage endpoint: {ex.Message}");
        return 6;
    }
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Required environment variable '{name}' is missing.");
    }

    return value;
}

static Uri GetRequiredHttpsUri(string name)
{
    string value = GetRequiredEnvironmentVariable(name);

    if (!Uri.TryCreate(value, UriKind.Absolute, out Uri? uri))
    {
        throw new UriFormatException(
            $"Environment variable '{name}' must contain an absolute URI.");
    }

    if (!string.Equals(uri.Scheme, Uri.UriSchemeHttps,
            StringComparison.OrdinalIgnoreCase))
    {
        throw new UriFormatException(
            $"Environment variable '{name}' must use HTTPS.");
    }

    return uri;
}
```

`ClientSecretCredential` itself is created directly as:

```csharp
var credential = new ClientSecretCredential(
    tenantId,
    clientId,
    clientSecret);
```

Creating the object does not immediately authenticate. Token acquisition normally happens during the first Azure SDK operation, so `AuthenticationFailedException` should surround service calls—not only credential construction.

## Secret-storage guidance

| Approach | Recommendation |
|---|---|
| Hardcoded in source or committed `appsettings.json` | Never use |
| Environment variables | Acceptable when securely injected by the host or CI/CD secret store |
| `.NET` User Secrets | Suitable for local development only; do not use in production |
| `IConfiguration` | Good configuration abstraction, but security depends on its underlying provider |
| Azure Key Vault or platform secret store | Prefer for production client secrets, with rotation and restricted access |
| Managed identity/workload identity | Preferred for Azure-hosted workloads because no client secret is stored |

Do not print the secret or include configuration values in logs. Rotate it before expiration, use a separate service principal per environment, and grant only the required RBAC role at the narrowest resource scope.

**References:** [Authenticate on-premises .NET apps with a service principal](https://learn.microsoft.com/dotnet/azure/sdk/authentication/on-premises-apps), [Azure Identity library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme), and [ClientSecretCredential API](https://learn.microsoft.com/dotnet/api/azure.identity.clientsecretcredential).

## Reviewer Notes (LLM-as-Judge)

### Summary

All 5 criteria passed (5/5). The response provides a complete C# service-principal authentication example with secure configuration guidance and specific error handling.

### Strengths

- Uses the required modern Azure SDK packages and credential type.
- Correctly accounts for lazy token acquisition by handling authentication failures around service operations.
- Distinguishes authentication failures, authorization failures, service errors, and configuration errors.
- Avoids hardcoding or logging the client secret.

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
hyoka run --prompt-id identity-dp-dotnet-service-principal --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
