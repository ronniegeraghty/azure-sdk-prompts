# Evaluation Report: identity-dp-dotnet-default-credential

**Config:** dotnet-azure-tools/baseline | **Result:** ✅ PASSED | **Duration:** 197.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-default-credential` |
| Config | dotnet-azure-tools/baseline |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 197.8s |
| Timestamp | 2026-08-26T17:23:00Z |
| Files Generated | 0 |
| Event Count | 4107 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 174.1s |
| Review | 23.0s |
| **Total** | **197.8s** |

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
| Token Usage | in=84373 out=3103 |
| Turn Count | 3 |

## Prompt Sent

```
Show me how to authenticate
an Azure SDK client using DefaultAzureCredential in C#. Explain:
1. What NuGet packages are needed
2. How to create and use a DefaultAzureCredential instance
3. The credential chain order (which credentials are tried and in what sequence)
4. How it behaves differently in local development vs deployed Azure environments
5. How to troubleshoot when authentication fails

Provide a complete, runnable example that creates a BlobServiceClient using
DefaultAzureCredential.
```

## Tool Calls

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com/dotnet/api Azure.Identity DefaultAzureCredential credential chain order .NET 2026"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com azure dotnet DefaultAzureCredential troubleshoot authentication BlobServiceClient package Azure.Storage.Blobs Azure.Identity"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:learn.microsoft.com dotnet azure sdk authentication local development DefaultAzureCredential deployed managed identity"}
```

### 🔧 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"site:github.com/Azure/azure-sdk-for-net DefaultAzureCredential.cs credential order Azure.Identity latest"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/credential-chains"}
```

### 🔧 6. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/troubleshoot-azure-identity"}
```

### 🔧 7. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication/best-practices"}
```

### 🔧 8. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-dotnet-get-started"}
```

### 🔧 9. web_fetch

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/dotnet/api/overview/azure/identity-readme?view=azure-dotnet"}
```

### 🔧 10. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Identity/DefaultAzureCredentialFactory.cs"}
```

## Copilot Response

I’ll verify the current .NET credential chain and troubleshooting behavior against Microsoft’s Azure SDK documentation.

## 1. NuGet packages

Create a console project and install the Azure Blob Storage and Azure Identity libraries:

```powershell
dotnet new console -n AzureBlobDefaultCredential
Set-Location AzureBlobDefaultCredential

dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

- [`Azure.Identity`](https://www.nuget.org/packages/Azure.Identity) provides `DefaultAzureCredential`.
- [`Azure.Storage.Blobs`](https://www.nuget.org/packages/Azure.Storage.Blobs) provides `BlobServiceClient`.
- `Azure.Core`, which defines `TokenCredential`, is installed transitively.
- Optional: [`Azure.Identity.Broker`](https://www.nuget.org/packages/Azure.Identity.Broker) enables brokered authentication, including the signed-in OS or supported VS Code account.

## 2. Complete runnable example

Replace `Program.cs` with:

```csharp
using Azure;
using Azure.Identity;
using Azure.Storage.Blobs;

string? accountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

if (string.IsNullOrWhiteSpace(accountName))
{
    Console.Error.WriteLine(
        "Set AZURE_STORAGE_ACCOUNT_NAME to the name of an existing storage account.");
    return 1;
}

Uri serviceUri =
    new($"https://{accountName}.blob.core.windows.net");

// Reuse credential instances so their token caches can be reused.
DefaultAzureCredential credential = new();

BlobServiceClient blobServiceClient = new(serviceUri, credential);

try
{
    Console.WriteLine($"Containers in {accountName}:");

    await foreach (var container in blobServiceClient.GetBlobContainersAsync())
    {
        Console.WriteLine($"- {container.Name}");
    }

    return 0;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine("Microsoft Entra authentication failed.");
    Console.Error.WriteLine(ex.Message);
    return 2;
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    Console.Error.WriteLine(
        "Authentication succeeded, but this identity isn't authorized to list containers.");
    Console.Error.WriteLine(
        "Assign an appropriate data-plane role, such as Storage Blob Data Reader.");
    Console.Error.WriteLine(ex.Message);
    return 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Storage request failed: HTTP {ex.Status}, code {ex.ErrorCode}");
    Console.Error.WriteLine(ex.Message);
    return 4;
}
```

Set the storage account name and run it:

```powershell
$env:AZURE_STORAGE_ACCOUNT_NAME = "yourstorageaccount"
dotnet run
```

The authenticated identity needs an Azure data-plane role such as **Storage Blob Data Reader** at the container or storage-account scope. A management role such as **Reader** does not grant access to blob data.

The essential construction is:

```csharp
DefaultAzureCredential credential = new();

BlobServiceClient client = new(
    new Uri("https://yourstorageaccount.blob.core.windows.net"),
    credential);
```

Azure SDK clients request and refresh tokens as needed, so application code normally doesn't call `GetTokenAsync` directly. Reuse the credential and client rather than constructing them for every operation.

## 3. Credential chain order

For the current `Azure.Identity` documentation, the default chain is attempted in this order:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal configuration in environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS or Kubernetes |
| 3 | `ManagedIdentityCredential` | System- or user-assigned Azure managed identity |
| 4 | `VisualStudioCredential` | Account signed into Visual Studio |
| 5 | `VisualStudioCodeCredential` | Supported VS Code account; requires broker support |
| 6 | `AzureCliCredential` | Account authenticated through `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated through `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated through `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser login; **disabled by default** |
| 10 | `BrokerCredential` | Brokered OS account; requires `Azure.Identity.Broker` |

The first credential that successfully returns a token is used. A credential that is unavailable—such as managed identity on a developer laptop—is skipped.

Since `Azure.Identity` 1.10.1, the chain continues through developer-tool credentials even if an earlier developer credential fails. A configured deployed-service credential that attempts authentication but cannot obtain a token can terminate the chain with an authentication error.

You can narrow the chain with `DefaultAzureCredentialOptions`, or with `AZURE_TOKEN_CREDENTIALS` in Azure.Identity 1.15.0 and later:

```powershell
# Only developer credentials
$env:AZURE_TOKEN_CREDENTIALS = "dev"

# Only deployed-service credentials
$env:AZURE_TOKEN_CREDENTIALS = "prod"

# Exactly one credential
$env:AZURE_TOKEN_CREDENTIALS = "AzureCliCredential"
```

See Microsoft’s [credential-chain reference](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains).

## 4. Local development versus Azure deployment

**Local development:** Managed and workload identity generally aren't available, so the chain normally reaches a developer credential. Sign in using one supported tool—for example, Visual Studio, Azure CLI, Azure PowerShell, or Azure Developer CLI. The signed-in user must have the required Blob Storage data role.

Alternatively, `EnvironmentCredential` can use a service principal configured with:

```text
AZURE_TENANT_ID
AZURE_CLIENT_ID
AZURE_CLIENT_SECRET
```

Avoid storing these values in source code or committed configuration.

**Deployed in Azure:** Enable a managed identity or workload identity on the hosting service and assign that identity the required resource role. The same source code then authenticates through `ManagedIdentityCredential` or `WorkloadIdentityCredential`; no client secret is needed.

For a user-assigned managed identity, identify it explicitly:

```csharp
DefaultAzureCredential credential = new(
    new DefaultAzureCredentialOptions
    {
        ManagedIdentityClientId =
            Environment.GetEnvironmentVariable("AZURE_CLIENT_ID")
    });
```

Although `DefaultAzureCredential` is convenient, Microsoft recommends a deterministic credential such as `ManagedIdentityCredential` for production once the intended authentication mechanism is known. This prevents an unexpected environment variable or developer login from changing the selected identity. See the [Azure Identity authentication best practices](https://learn.microsoft.com/dotnet/azure/sdk/authentication/best-practices).

## 5. Troubleshooting authentication failures

1. **Enable Azure Identity logging** to see every credential attempted and why it was skipped or failed:

   ```csharp
   using System.Diagnostics.Tracing;
   using Azure.Core.Diagnostics;

   using AzureEventSourceListener listener = new(
       (args, message) =>
       {
           if (args.EventSource.Name == "Azure-Identity")
           {
               Console.Error.WriteLine(message);
           }
       },
       EventLevel.Verbose);
   ```

   Create the listener before the Azure SDK request.

2. **Distinguish authentication from authorization.**
   - `AuthenticationFailedException`: no credential successfully obtained a token.
   - HTTP `401`: token, tenant, audience, or authentication configuration problem.
   - HTTP `403`: authentication often succeeded, but the selected identity lacks permission, the role assignment hasn't propagated, or a storage networking/policy restriction applies.

3. **Determine which identity was selected.** A successful CLI, Visual Studio, environment, or managed-identity credential may represent a different principal than expected. Identity logging shows the successful credential type.

4. **Check local login state and tenant.** Reauthenticate the chosen developer tool and ensure it is using the tenant containing the storage account. Cached credentials can be valid but belong to the wrong tenant.

5. **Check environment variables.** Partially configured `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, or `AZURE_CLIENT_SECRET` values can cause `EnvironmentCredential` failures. Also inspect `AZURE_TOKEN_CREDENTIALS`, which can restrict the chain.

6. **Check managed identity configuration.** Confirm the identity is enabled on the Azure host. If multiple user-assigned identities exist, specify the intended client ID or resource ID. Managed identity isn't available when the program runs directly on a normal development workstation.

7. **Check Blob Storage data-plane RBAC.** Use a role such as **Storage Blob Data Reader** or **Storage Blob Data Contributor**, not merely a subscription or storage-account management-plane role. New role assignments can take several minutes to propagate.

8. **Reduce the chain while diagnosing.** For example, force Azure CLI authentication:

   ```powershell
   $env:AZURE_TOKEN_CREDENTIALS = "AzureCliCredential"
   dotnet run
   ```

   This requires `Azure.Identity` 1.15.0 or later and makes failures more deterministic.

Official references: [Azure Identity library](https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme), [Blob Storage .NET setup](https://learn.microsoft.com/azure/storage/blobs/storage-blob-dotnet-get-started), and [credential-chain debugging](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains#debug-a-chained-credential).

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response comprehensively covers package setup, credential construction and configuration, chain behavior, BlobServiceClient integration, and authentication troubleshooting.

### Strengths

- Complete runnable C# Blob Storage example
- Detailed and correctly ordered credential-chain explanation
- Clear local-development versus Azure-deployment guidance
- Strong exception handling and diagnostic logging guidance

## Grader Results

- default-azure-credential.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - `Azure.Identity` NuGet package installation: Pass
      - `DefaultAzureCredential` constructor and options: Pass
      - Credential chain: Environment → Workload Identity → Managed Identity → Azure CLI → etc.: Pass
      - Passing credential to any Azure SDK client constructor: Pass
      - `AuthenticationFailedException` handling and diagnostics: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-dotnet-default-credential --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
