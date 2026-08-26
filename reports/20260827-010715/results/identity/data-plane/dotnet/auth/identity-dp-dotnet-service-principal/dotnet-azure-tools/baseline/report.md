# Evaluation Report: identity-dp-dotnet-service-principal

**Config:** dotnet-azure-tools/baseline | **Result:** ✅ PASSED | **Duration:** 54.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-service-principal` |
| Config | dotnet-azure-tools/baseline |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 54.5s |
| Timestamp | 2026-08-26T17:32:56Z |
| Files Generated | 0 |
| Event Count | 2090 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 33.9s |
| Review | 19.9s |
| **Total** | **54.5s** |

## Configuration

- **name:** dotnet-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=19627 out=1408 |
| Turn Count | 1 |

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

## Copilot Response

Required packages:

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

Set the service principal and Storage account details as environment variables:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
$env:AZURE_STORAGE_ACCOUNT_NAME = "<storage-account-name>"
```

`ClientSecretCredential` acquires tokens lazily, so invalid credentials normally fail when the Azure SDK client performs its first operation—not when the credential object is constructed.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

internal static class Program
{
    private static async Task<int> Main()
    {
        try
        {
            string tenantId = GetRequiredEnvironmentVariable("AZURE_TENANT_ID");
            string clientId = GetRequiredEnvironmentVariable("AZURE_CLIENT_ID");
            string clientSecret = GetRequiredEnvironmentVariable("AZURE_CLIENT_SECRET");
            string storageAccountName =
                GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

            TokenCredential credential = new ClientSecretCredential(
                tenantId,
                clientId,
                clientSecret,
                new ClientSecretCredentialOptions
                {
                    // Exclude credential details and token contents from logs.
                    Diagnostics =
                    {
                        IsAccountIdentifierLoggingEnabled = false
                    }
                });

            var serviceUri = new Uri(
                $"https://{storageAccountName}.blob.core.windows.net");

            var blobServiceClient = new BlobServiceClient(
                serviceUri,
                credential);

            Console.WriteLine("Authenticated. Listing Blob containers:");

            await foreach (
                var container in blobServiceClient.GetBlobContainersAsync())
            {
                Console.WriteLine($"- {container.Name}");
            }

            return 0;
        }
        catch (InvalidOperationException exception)
        {
            // A required local configuration value is absent.
            Console.Error.WriteLine($"Configuration error: {exception.Message}");
            return 1;
        }
        catch (AuthenticationFailedException exception)
        {
            // Common causes include an invalid tenant ID, client ID, expired
            // secret, incorrect secret value, or disabled service principal.
            Console.Error.WriteLine(
                "Azure authentication failed. Verify the tenant ID, client ID, " +
                "client secret, and secret expiration.");

            // Use structured logging in production. Never log the secret or token.
            Console.Error.WriteLine($"Azure Identity error: {exception.Message}");
            return 2;
        }
        catch (RequestFailedException exception)
            when (exception.Status is 401 or 403)
        {
            // Authentication may have succeeded, but the principal may lack
            // permission on the Storage account.
            Console.Error.WriteLine(
                $"Azure Storage denied the request (HTTP {exception.Status}). " +
                "Verify the service principal's Azure RBAC role.");

            Console.Error.WriteLine(
                $"Error code: {exception.ErrorCode ?? "unknown"}");
            return 3;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Azure Storage request failed (HTTP {exception.Status}, " +
                $"code {exception.ErrorCode ?? "unknown"}): {exception.Message}");
            return 4;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The Azure operation was canceled.");
            return 5;
        }
        catch (Exception exception)
        {
            Console.Error.WriteLine(
                $"An unexpected error occurred: {exception.Message}");
            return 10;
        }
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);

        if (string.IsNullOrWhiteSpace(value))
        {
            throw new InvalidOperationException(
                $"Required environment variable '{name}' is not set.");
        }

        return value;
    }
}
```

The service principal also needs an Azure RBAC data-plane role such as **Storage Blob Data Reader** for listing and reading containers, or **Storage Blob Data Contributor** when writes are required. A successful token acquisition does not imply access to the Storage account; missing RBAC permissions generally produce HTTP `403`.

**Secret-storage practices:**

- Never commit a secret to source code or `appsettings.json`.
- Environment variables are suitable for local examples and CI/CD when populated by a protected secret store. Avoid placing the secret permanently in shell profiles or scripts.
- For local development, .NET User Secrets is preferable to plaintext configuration files.
- In production, use the platform’s secret manager or inject the value from Azure Key Vault. `IConfiguration` is only an abstraction—the security depends on its underlying provider.
- Restrict access to the secret, rotate it regularly, monitor its expiration, and avoid logging configuration values.
- When the application runs on Azure, prefer `ManagedIdentityCredential` or `DefaultAzureCredential` over a client secret whenever possible.

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response provides the required packages, credential construction, Azure SDK client integration, secure secret-storage guidance, and invalid-credential handling.

### Strengths

- Complete, practical C# example with lazy-authentication behavior explained.
- Distinguishes authentication failures from authorization and other Azure request failures.
- Avoids hard-coded secrets and includes production-oriented secret-management guidance.

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
hyoka run --prompt-id identity-dp-dotnet-service-principal --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
