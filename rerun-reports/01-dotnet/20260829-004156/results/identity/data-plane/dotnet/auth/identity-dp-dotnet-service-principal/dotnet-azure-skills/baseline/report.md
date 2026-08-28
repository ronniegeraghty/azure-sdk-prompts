# Evaluation Report: identity-dp-dotnet-service-principal

**Config:** dotnet-azure-skills/baseline | **Result:** ✅ PASSED | **Duration:** 58.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-service-principal` |
| Config | dotnet-azure-skills/baseline |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 58.1s |
| Timestamp | 2026-08-28T17:50:54Z |
| Files Generated | 0 |
| Event Count | 2544 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 36.5s |
| Review | 19.5s |
| **Total** | **58.1s** |

## Configuration

- **name:** dotnet-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=19631 out=1518 |
| Turn Count | 1 |

## Copilot Response

Install the Azure Identity and Blob Storage SDK packages:

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

`ClientSecretCredential` authenticates when the SDK makes its first service request—not when the credential object is constructed.

### Complete example

```csharp
using Azure;
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
            string blobServiceUrl =
                GetRequiredEnvironmentVariable("AZURE_STORAGE_BLOB_SERVICE_URL");

            var credential = new ClientSecretCredential(
                tenantId,
                clientId,
                clientSecret,
                new ClientSecretCredentialOptions
                {
                    AuthorityHost = AzureAuthorityHosts.AzurePublicCloud
                });

            var blobServiceClient = new BlobServiceClient(
                new Uri(blobServiceUrl),
                credential);

            // Authentication happens when this request is made.
            AccountInfo accountInfo =
                await blobServiceClient.GetAccountInfoAsync();

            Console.WriteLine(
                $"Authenticated successfully. Account kind: {accountInfo.AccountKind}");

            Console.WriteLine("Containers:");

            await foreach (BlobContainerItem container
                in blobServiceClient.GetBlobContainersAsync())
            {
                Console.WriteLine($"- {container.Name}");
            }

            return 0;
        }
        catch (InvalidOperationException ex)
        {
            // Missing or empty local configuration.
            Console.Error.WriteLine($"Configuration error: {ex.Message}");
            return 1;
        }
        catch (AuthenticationFailedException ex)
        {
            // Examples: invalid tenant ID, client ID, secret, expired secret,
            // or an application that does not exist in the specified tenant.
            Console.Error.WriteLine(
                "Azure authentication failed. Verify the tenant ID, client ID, " +
                "and client secret, and confirm the secret has not expired.");

            // Log detailed diagnostics only to an appropriately protected sink.
            Console.Error.WriteLine($"Details: {ex.Message}");
            return 2;
        }
        catch (RequestFailedException ex) when (ex.Status == 401)
        {
            Console.Error.WriteLine(
                "Azure rejected the access token. Verify the service principal " +
                "and token audience configuration.");
            Console.Error.WriteLine(
                $"Azure error: {ex.ErrorCode ?? "unknown"}; Request ID: {ex.ClientRequestId}");
            return 3;
        }
        catch (RequestFailedException ex) when (ex.Status == 403)
        {
            // Authentication may have succeeded, but the principal lacks authorization.
            Console.Error.WriteLine(
                "Authentication succeeded, but the service principal is not authorized " +
                "for this storage account. Verify its Azure RBAC role assignment.");
            Console.Error.WriteLine(
                $"Azure error: {ex.ErrorCode ?? "unknown"}; Request ID: {ex.ClientRequestId}");
            return 4;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure request failed with HTTP {ex.Status}. " +
                $"Error: {ex.ErrorCode ?? "unknown"}; Request ID: {ex.ClientRequestId}");
            return 5;
        }
        catch (UriFormatException)
        {
            Console.Error.WriteLine(
                "AZURE_STORAGE_BLOB_SERVICE_URL is not a valid URI.");
            return 6;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The Azure operation was canceled.");
            return 7;
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Unexpected error: {ex.Message}");
            return 8;
        }
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);

        if (string.IsNullOrWhiteSpace(value))
        {
            throw new InvalidOperationException(
                $"Required environment variable '{name}' is missing or empty.");
        }

        return value;
    }
}
```

`BlobContainerItem` requires this additional import:

```csharp
using Azure.Storage.Blobs.Models;
```

Alternatively, add it to the imports in the complete example above.

### Environment variables

PowerShell, for the current process:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_STORAGE_BLOB_SERVICE_URL = "https://<account-name>.blob.core.windows.net"

dotnet run
```

Use the **client secret value**, not the secret's identifier. The service principal also needs an appropriate data-plane RBAC role, such as **Storage Blob Data Reader** for listing containers. A management-plane role such as `Reader` alone does not grant access to blob data.

### Secret-storage best practices

| Method | Recommended use |
|---|---|
| Environment variables | Containers, CI/CD, and externally injected runtime configuration. Keep them out of scripts, logs, and source control. |
| .NET user secrets | Local development: `dotnet user-secrets`. They are not encrypted, but remain outside the project tree. |
| `appsettings.json` | Safe for tenant ID, client ID, and service URL; do **not** store a real client secret in a committed file. |
| Azure Key Vault | Prefer for centrally managed production secrets, ideally accessed with Managed Identity. |
| Managed Identity | Preferred for applications hosted in Azure because it eliminates client-secret storage and rotation. |

Environment variables are better than committed configuration files, but they are still secrets and can sometimes be inspected by privileged users or process-management tooling. Inject them through the hosting platform or CI/CD secret store, restrict access, rotate them regularly, and never include them in exception messages or telemetry.

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response covers the required packages, credential construction, Azure SDK integration, environment-variable secret storage, and authentication failure handling.

### Strengths

- Clearly explains that authentication occurs on the first service request rather than during credential construction.
- Distinguishes authentication failures from 401, 403, configuration, cancellation, and other request failures.
- Provides strong production secret-storage and Azure RBAC guidance.

### Issues

- The purported complete example omits using Azure.Storage.Blobs.Models, so BlobContainerItem will not resolve unless the separately mentioned import is added.

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
hyoka run --prompt-id identity-dp-dotnet-service-principal --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
