# Evaluation Report: identity-dp-dotnet-default-credential

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ✅ PASSED | **Duration:** 152.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-dotnet-default-credential` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 152.3s |
| Timestamp | 2026-08-26T17:33:51Z |
| Files Generated | 0 |
| Event Count | 3724 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 120.6s |
| Review | 30.5s |
| **Total** | **152.3s** |

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
| Token Usage | in=205928 out=3962 |
| Turn Count | 5 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide a local-only C# example using Azure.Identity DefaultAzureCredential with Azure.Storage.Blobs BlobServiceClient, including authentication guidance and troubleshooting.","learn":true}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official documentation for DefaultAzureCredential in .NET, especially exact credential chain order, continuation policy, local versus Azure-hosted behavior, and troubleshooting.","learn":true}
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
{"command":"get_azure_bestpractices_get","intent":"Get Azure C# SDK authentication code-generation best practices for DefaultAzureCredential and BlobServiceClient.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official DefaultAzureCredential .NET chain, behavior, and troubleshooting documentation.","parameters":{"query":"DefaultAzureCredential .NET credential chain order continuation policy local development deployed Azure troubleshooting Azure.Identity 2026"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# code samples for BlobServiceClient with DefaultAzureCredential.","parameters":{"language":"csharp","query":"C# BlobServiceClient DefaultAzureCredential Azure.Storage.Blobs runnable console app"}}
```

### 🔧 9. glob

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet","pattern":"**/*"}
```

### 🔧 10. rg

**Input:**

```json
{"-n":true,"C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787765688207-copilot-tool-output-1996658f50e9400f8594bb9926b24b58.txt","pattern":"credential-chains|troubleshoot|continuation|EnvironmentCredential|BrokerCredential|VisualStudioCodeCredential"}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the complete current DefaultAzureCredential chain overview for .NET.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure Identity .NET authentication troubleshooting logging guidance.","parameters":{"query":"site:learn.microsoft.com/dotnet/azure/sdk authentication troubleshooting DefaultAzureCredential Azure Identity logging AzureEventSourceListener"}}
```

## Copilot Response

I’m grounding the example in the current Azure Identity chain and SDK guidance so the package list and credential order match today’s .NET behavior.

The current library separates production credentials from developer-tool credentials; I’m checking the exact order and continuation behavior before presenting the runnable sample.

`DefaultAzureCredential` provides one authentication API that works with developer sign-ins locally and Azure workload identities after deployment. It acquires tokens lazily—the first Azure service request triggers authentication.

## 1. NuGet packages

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

- `Azure.Identity`: provides `DefaultAzureCredential`.
- `Azure.Storage.Blobs`: provides `BlobServiceClient`.
- Optional: `Azure.Identity.Broker` enables Visual Studio Code and OS-broker authentication:

```powershell
dotnet add package Azure.Identity.Broker
```

`Azure.Core` is pulled in transitively; you normally should not reference it directly.

## 2. Complete runnable example

Create the project:

```powershell
dotnet new console --name DefaultCredentialBlobDemo
Set-Location DefaultCredentialBlobDemo
dotnet add package Azure.Identity
dotnet add package Azure.Storage.Blobs
```

Replace `Program.cs` with:

```csharp
using System.Diagnostics.Tracing;
using Azure;
using Azure.Core.Diagnostics;
using Azure.Identity;
using Azure.Storage.Blobs;

string accountName =
    Environment.GetEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME")
    ?? "examplestorage";

Uri serviceUri = new($"https://{accountName}.blob.core.windows.net");

// Optional diagnostic logging. Avoid enabling verbose logs routinely in production.
using AzureEventSourceListener? identityListener =
    Environment.GetEnvironmentVariable("AZURE_IDENTITY_DEBUG") == "1"
        ? new AzureEventSourceListener(
            (eventArgs, message) =>
            {
                if (eventArgs.EventSource.Name == "Azure-Identity")
                {
                    Console.Error.WriteLine(message);
                }
            },
            EventLevel.Verbose)
        : null;

// Credentials are thread-safe and should be reused.
DefaultAzureCredential credential = new();

BlobServiceClient blobServiceClient = new(serviceUri, credential);

Console.WriteLine($"Created BlobServiceClient for {blobServiceClient.Uri}");

// Client construction is offline: no token is requested until an operation runs.
// Pass --probe to authenticate and make a read-only request to Azure.
if (!args.Contains("--probe", StringComparer.OrdinalIgnoreCase))
{
    Console.WriteLine("Run with --probe to authenticate and list blob containers.");
    return;
}

try
{
    Console.WriteLine("Containers:");

    await foreach (BlobContainerItem container in
        blobServiceClient.GetBlobContainersAsync())
    {
        Console.WriteLine($"- {container.Name}");
    }
}
catch (CredentialUnavailableException ex)
{
    Console.Error.WriteLine(
        $"No credential in DefaultAzureCredential was available: {ex.Message}");
    Environment.ExitCode = 1;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Microsoft Entra authentication failed: {ex.Message}");
    Environment.ExitCode = 1;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Blob Storage request failed ({ex.Status}/{ex.ErrorCode}): {ex.Message}");
    Environment.ExitCode = 1;
}
```

Run without contacting Azure:

```powershell
dotnet run
```

To make the read-only request:

```powershell
$env:AZURE_STORAGE_ACCOUNT_NAME = "your-storage-account"
dotnet run -- --probe
```

The authenticated identity needs at least **Storage Blob Data Reader** on the storage account or appropriate container scope. For uploads or modifications, use **Storage Blob Data Contributor**.

## 3. Credential chain order

With the current `Azure.Identity` behavior, `DefaultAzureCredential` tries:

| Order | Credential | Typical source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity, commonly AKS |
| 3 | `ManagedIdentityCredential` | Azure managed identity |
| 4 | `VisualStudioCredential` | Account signed into Visual Studio |
| 5 | `VisualStudioCodeCredential` | VS Code account; requires `Azure.Identity.Broker` |
| 6 | `AzureCliCredential` | Account authenticated with `az login` |
| 7 | `AzurePowerShellCredential` | Account authenticated with `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account authenticated with `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser login; **disabled by default** |
| 10 | `BrokerCredential` | OS account broker; requires `Azure.Identity.Broker` |

The chain stops after one credential succeeds. Interactive browser authentication can be enabled explicitly:

```csharp
DefaultAzureCredential credential = new(
    new DefaultAzureCredentialOptions
    {
        ExcludeInteractiveBrowserCredential = false
    });
```

The continuation policy is important:

- An unavailable credential—for example, managed identity on a laptop—is skipped.
- Developer-tool credentials continue through the developer portion of the chain until one succeeds.
- If a configured deployed-service credential can attempt authentication but fails, the failure can stop the chain. This prevents a broken production identity from silently falling back to a developer identity.

## 4. Local versus deployed behavior

**Local development:** managed identity and workload identity are usually unavailable, so the chain reaches a developer credential such as Visual Studio, Azure CLI, Azure PowerShell, or Azure Developer CLI. The resulting access is that developer’s permissions.

**Azure deployment:** enable a managed identity or workload identity and assign it an Azure Storage data-plane role. The chain normally succeeds at `WorkloadIdentityCredential` or `ManagedIdentityCredential` before checking developer tools. No client secret needs to be stored.

For a user-assigned managed identity, identify it explicitly:

```csharp
DefaultAzureCredential credential = new(
    new DefaultAzureCredentialOptions
    {
        ManagedIdentityClientId = Environment.GetEnvironmentVariable(
            "AZURE_MANAGED_IDENTITY_CLIENT_ID")
    });
```

Although `DefaultAzureCredential` is convenient across environments, production applications often use a deterministic credential to avoid ambiguity and unnecessary probing:

```csharp
ManagedIdentityCredential credential = new(
    ManagedIdentityId.SystemAssigned);
```

## 5. Troubleshooting failures

1. **Enable Azure Identity logs:**

   ```powershell
   $env:AZURE_IDENTITY_DEBUG = "1"
   dotnet run -- --probe
   ```

   Look for `DefaultAzureCredential credential selected:` and the preceding failures.

2. **Separate authentication from authorization:**
   - Authentication errors mean no credential obtained a token.
   - HTTP `403` generally means authentication succeeded but the identity lacks a Blob Storage **data-plane** role.
   - Azure resource roles such as Contributor do not automatically grant access to blob data.

3. **Confirm local sign-in state:** sign into Visual Studio or authenticate one of the supported developer tools. Ensure the signed-in identity belongs to the correct tenant.

4. **Check environment configuration:** a service principal generally requires all of:

   ```text
   AZURE_TENANT_ID
   AZURE_CLIENT_ID
   AZURE_CLIENT_SECRET
   ```

   Partially configured or stale values can cause `EnvironmentCredential` failures. Never commit these values.

5. **Check deployed identity configuration:** verify that managed/workload identity is enabled, the intended user-assigned identity is selected, and its role assignment is scoped correctly. New role assignments can take several minutes to propagate.

6. **Narrow the chain while diagnosing:** with `Azure.Identity` 1.15.0 or later, set `AZURE_TOKEN_CREDENTIALS` to `dev`, `prod`, or a specific credential name such as `AzureCliCredential`. You can also use `DefaultAzureCredentialOptions` exclusion properties.

Official references:

- [Credential chains in Azure Identity for .NET](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains)
- [Azure SDK authentication best practices](https://learn.microsoft.com/dotnet/azure/sdk/authentication/best-practices)
- [BlobServiceClient authentication example](https://learn.microsoft.com/azure/storage/blobs/storage-blob-client-management#create-a-client-object)
- [Azure Identity troubleshooting](https://learn.microsoft.com/dotnet/azure/sdk/authentication/credential-chains#debug-a-chained-credential)

## Reviewer Notes (LLM-as-Judge)

### Summary

5/5 criteria passed. The response comprehensively covers package installation, credential configuration and ordering, client construction, environment-specific behavior, and authentication diagnostics.

### Strengths

- Includes practical DefaultAzureCredentialOptions examples for interactive browser and user-assigned managed identity scenarios.
- Clearly distinguishes authentication failures from authorization failures such as HTTP 403 responses.
- Explains credential-chain continuation behavior and local-versus-deployed identity selection.

### Issues

- The claimed runnable example is missing `using Azure.Storage.Blobs.Models;`, so `BlobContainerItem` will not resolve as written.

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
hyoka run --prompt-id identity-dp-dotnet-default-credential --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
