# Evaluation Report: storage-dp-dotnet-auth

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 46.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-auth` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 3/5 |
| Duration | 46.5s |
| Timestamp | 2026-08-26T18:06:11Z |
| Files Generated | 0 |
| Event Count | 1156 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 27.4s |
| Review | 18.3s |
| **Total** | **46.5s** |

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
| Token Usage | in=19585 out=940 |
| Turn Count | 1 |

## Prompt Sent

```
How do I authenticate to Azure Blob Storage using DefaultAzureCredential in C#?
I need to create a BlobServiceClient that uses managed identity in production
but falls back to Azure CLI credentials during local development.
Show me the complete setup including required NuGet packages.
```

## Copilot Response

Install the Azure Blob Storage and identity libraries:

```powershell
dotnet add package Azure.Storage.Blobs
dotnet add package Azure.Identity
```

Configure the storage account endpoint, not an account key or connection string:

```json
// appsettings.json
{
  "AzureStorage": {
    "ServiceUri": "https://<storage-account-name>.blob.core.windows.net"
  }
}
```

For ASP.NET Core, register a singleton `DefaultAzureCredential` and `BlobServiceClient`:

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddSingleton<TokenCredential>(_ =>
{
    var managedIdentityClientId =
        builder.Configuration["AZURE_CLIENT_ID"];

    return new DefaultAzureCredential(
        new DefaultAzureCredentialOptions
        {
            // Leave unset for a system-assigned managed identity.
            // Set AZURE_CLIENT_ID for a user-assigned managed identity.
            ManagedIdentityClientId = managedIdentityClientId
        });
});

builder.Services.AddSingleton(serviceProvider =>
{
    var serviceUri =
        builder.Configuration["AzureStorage:ServiceUri"]
        ?? throw new InvalidOperationException(
            "AzureStorage:ServiceUri is not configured.");

    var credential = serviceProvider.GetRequiredService<TokenCredential>();

    return new BlobServiceClient(new Uri(serviceUri), credential);
});

var app = builder.Build();

app.MapGet("/containers", async (BlobServiceClient client) =>
{
    var containers = new List<string>();

    await foreach (var container in client.GetBlobContainersAsync())
    {
        containers.Add(container.Name);
    }

    return containers;
});

app.Run();
```

`DefaultAzureCredential` automatically selects an available credential:

- **Production on Azure:** managed identity is used.
- **Local development:** after authenticating with `az login`, Azure CLI credentials are used.
- **User-assigned identity:** set `AZURE_CLIENT_ID` to its client ID.
- **System-assigned identity:** leave `AZURE_CLIENT_ID` unset.

Example local configuration:

```powershell
$env:AzureStorage__ServiceUri = "https://mystorageaccount.blob.core.windows.net"
az login
dotnet run
```

The identity also needs a Blob **data-plane** role on the storage account or relevant container, typically:

- `Storage Blob Data Reader` for read-only access
- `Storage Blob Data Contributor` for reading and writing
- `Storage Blob Data Owner` when ownership/ACL operations are required

The general Azure `Contributor` role alone does not grant access to blob data. Assign the same appropriate data role to your local developer identity and the production managed identity. No storage account keys or connection strings are required.

## Reviewer Notes (LLM-as-Judge)

### Summary

3 of 5 criteria passed. The setup correctly installs the packages and wires DefaultAzureCredential into BlobServiceClient, but the credential-chain explanation and unavailable-credential handling are incomplete.

### Strengths

- Includes both required NuGet package commands.
- Shows complete ASP.NET Core dependency-injection setup for DefaultAzureCredential and BlobServiceClient.
- Clearly distinguishes system-assigned and user-assigned managed identity configuration and documents required Blob data-plane roles.

### Issues

- Environment-based credentials are omitted from the credential-chain explanation.
- No error handling is shown for CredentialUnavailableException or AuthenticationFailedException when authentication cannot succeed.

## Grader Results

- authentication.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (3/5)
      - Installing `Azure.Identity` and `Azure.Storage.Blobs` packages: Pass
      - Creating a `DefaultAzureCredential` instance: Pass
      - Passing the credential to `BlobServiceClient`: Pass
      - Explanation of the credential chain (managed identity → CLI → env vars): Fail
      - Error handling when no credential is available: Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 60% | 1.00 | 0.6000 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.6000** | **60.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-auth --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
