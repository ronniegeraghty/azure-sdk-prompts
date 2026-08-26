# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 131.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 5/7 |
| Duration | 131.8s |
| Timestamp | 2026-08-26T17:46:13Z |
| Files Generated | 0 |
| Event Count | 3282 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 107.8s |
| Review | 22.9s |
| **Total** | **131.8s** |

## Configuration

- **name:** dotnet-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=330120 out=3700 |
| Turn Count | 8 |

## Prompt Sent

```
How do I properly handle errors when working with Azure Key Vault secrets in .NET?
I need to handle common failures: access denied when RBAC or access policies
aren't configured correctly (403), secret not found (404), secret version
conflicts, and throttling when hitting Key Vault rate limits (429).
Show me try/catch patterns with Azure.Security.KeyVault.Secrets including
how to extract the error code and HTTP status from RequestFailedException.
```

## Tool Calls

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Provide modern .NET Azure Key Vault Secrets error-handling examples using Azure.Security.KeyVault.Secrets and Azure.RequestFailedException for 403, 404, version conflicts, and 429 throttling.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for .NET Azure Key Vault Secrets exception handling.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Azure SDK for .NET documentation for Azure.Security.KeyVault.Secrets RequestFailedException handling, including ErrorCode, Status, 403, 404, 409 conflicts, and 429 retry behavior.","learn":true}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find RequestFailedException Status and ErrorCode handling for Azure Key Vault Secrets .NET, including HTTP 403, 404, 409, and 429.","parameters":{"query":"Azure.Security.KeyVault.Secrets .NET RequestFailedException Status ErrorCode 403 404 409 429 retry policy secret version conflict"}}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# examples for Azure Key Vault SecretClient exception handling and retry configuration.","parameters":{"language":"csharp","query":"Azure.Security.KeyVault.Secrets SecretClient RequestFailedException ErrorCode Status RetryOptions"}}
```

### 🔧 6. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766419611-copilot-tool-output-2fe19e4d520946038e0d2f7111c1191e.txt","pattern":"https://[^\"\\\\ ]+|ErrorCode|Status ==|Retry|409|429|SecretAlreadyExists|Conflict"}
```

### 🔧 7. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766419863-copilot-tool-output-3c7cc576470c42c99c7d612127894098.txt","pattern":"https://[^\"\\\\ ]+|ErrorCode|Status ==|Retry|409|429|SecretAlreadyExists|Conflict"}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"$files = @('C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766419611-copilot-tool-output-2fe19e4d520946038e0d2f7111c1191e.txt','C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766419863-copilot-tool-output-3c7cc576470c42c99c7d612127894098.txt'); foreach ($file in $files) { $json = Get-Content -Raw $file | ConvertFrom-Json; $json.results | Select-Object -First 10 | ForEach-Object { [pscustomobject]@{ Title=$_.title; Url=$_.url; Description=$_.description; Content=if ($_.content) { $_.content.Substring(0,[Math]::Min(900,$_.content.Length)) } else { $null }; Code=$_.codeSnippet } } | Format-List }","description":"Extract relevant Azure documentation search results"}
```

### 🔧 9. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766419611-copilot-tool-output-2fe19e4d520946038e0d2f7111c1191e.txt","view_range":[1,8]}
```

### 🔧 10. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766419863-copilot-tool-output-3c7cc576470c42c99c7d612127894098.txt","view_range":[1,8]}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official SecretClient troubleshooting documentation.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet#troubleshooting"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read complete official Key Vault throttling guidance.","parameters":{"url":"https://learn.microsoft.com/azure/key-vault/general/overview-throttling#how-to-throttle-your-app-in-response-to-service-limits"}}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official RequestFailedException API documentation for Status and ErrorCode.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/azure.requestfailedexception?view=azure-dotnet"}}
```

### 🔧 14. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read official common Key Vault error code documentation, especially AccessDenied and ConflictError.","parameters":{"url":"https://learn.microsoft.com/azure/key-vault/general/common-error-codes"}}
```

## Copilot Response

`RequestFailedException` exposes both the HTTP status through `Status` and the Key Vault service error through `ErrorCode`. Use exception filters for expected outcomes, and configure the SDK retry policy for throttling.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;
using Microsoft.Extensions.Logging;

var options = new SecretClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(16),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

var client = new SecretClient(
    new Uri(Environment.GetEnvironmentVariable("KEY_VAULT_URI")
        ?? throw new InvalidOperationException("KEY_VAULT_URI is not configured.")),
    new DefaultAzureCredential(),
    options);
```

Reuse `SecretClient` as a singleton; its methods are thread-safe. In production, `DefaultAzureCredential` normally selects the application's managed identity.

```csharp
public static async Task<string?> GetSecretValueAsync(
    SecretClient client,
    string secretName,
    string? version,
    ILogger logger,
    CancellationToken cancellationToken)
{
    try
    {
        Response<KeyVaultSecret> response = await client.GetSecretAsync(
            secretName,
            version,
            cancellationToken);

        return response.Value.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        // Usually incorrect Key Vault RBAC/access policy, but firewall or
        // private-endpoint configuration can also produce a 403.
        logger.LogError(
            "Key Vault denied access. Status={Status}, ErrorCode={ErrorCode}, " +
            "RequestId={RequestId}",
            ex.Status,
            ex.ErrorCode,
            GetRequestId(ex));

        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 404)
    {
        // Common ErrorCode values include SecretNotFound.
        logger.LogInformation(
            "Secret {SecretName} version {Version} was not found. " +
            "Status={Status}, ErrorCode={ErrorCode}, RequestId={RequestId}",
            secretName,
            version ?? "<latest>",
            ex.Status,
            ex.ErrorCode,
            GetRequestId(ex));

        // Return null only if "not found" is an expected application outcome.
        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 409)
    {
        // Relevant mainly to write, delete, recover, and purge operations.
        // ConflictError can indicate concurrent operations on the same item.
        logger.LogWarning(
            "Key Vault operation conflicted with another operation. " +
            "Status={Status}, ErrorCode={ErrorCode}, RequestId={RequestId}",
            ex.Status,
            ex.ErrorCode,
            GetRequestId(ex));

        throw;
    }
    catch (RequestFailedException ex) when (ex.Status == 429)
    {
        // This is reached only after the configured SDK retries are exhausted.
        logger.LogWarning(
            "Key Vault throttling persisted after retries. " +
            "Status={Status}, ErrorCode={ErrorCode}, RequestId={RequestId}",
            ex.Status,
            ex.ErrorCode,
            GetRequestId(ex));

        throw;
    }
    catch (RequestFailedException ex)
    {
        logger.LogError(
            ex,
            "Unexpected Key Vault failure. Status={Status}, " +
            "ErrorCode={ErrorCode}, RequestId={RequestId}",
            ex.Status,
            ex.ErrorCode,
            GetRequestId(ex));

        throw;
    }
}

private static string? GetRequestId(RequestFailedException exception)
{
    Response? response = exception.GetRawResponse();

    return response is not null &&
           response.Headers.TryGetValue("x-ms-request-id", out string requestId)
        ? requestId
        : null;
}
```

The key fields are:

```csharp
catch (RequestFailedException ex)
{
    int httpStatus = ex.Status;       // 403, 404, 409, 429, etc.
    string? serviceCode = ex.ErrorCode; // AccessDenied, SecretNotFound,
                                         // ConflictError, etc.
    Response? response = ex.GetRawResponse();
}
```

Important behavior:

- **403:** Don’t repeatedly retry a persistent authorization failure. Verify the calling identity, its `Key Vault Secrets User` or other appropriate RBAC role, legacy access policy permissions, firewall rules, and private endpoint configuration. A short bounded retry may be reasonable immediately after changing RBAC because permissions can take time to propagate.
- **404:** Decide whether absence is expected. Return `null`/a result type when it is; otherwise rethrow or translate it to a domain exception.
- **Version conflicts:** `SetSecretAsync(name, value)` does **not** conflict merely because the name exists—it creates a new secret version. HTTP 409 generally indicates overlapping operations on the same item, such as concurrent delete/recover/purge/write workflows. Sequence those operations and retry only when the operation is idempotent.
- **429:** Azure Core automatically retries throttled requests using the configured policy and service-provided retry delay when available. A caught 429 means those retries were exhausted. Avoid layering an unbounded retry loop around the SDK; instead reduce request frequency, cache secret values in memory, and enforce a total retry/time budget.
- Never log `KeyVaultSecret.Value`.

**References:** [Key Vault Secrets .NET client 4.11.0](https://learn.microsoft.com/dotnet/api/overview/azure/security.keyvault.secrets-readme?view=azure-dotnet), [`RequestFailedException`](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception?view=azure-dotnet), [Key Vault throttling guidance](https://learn.microsoft.com/azure/key-vault/general/overview-throttling), and [common Key Vault error codes](https://learn.microsoft.com/azure/key-vault/general/common-error-codes).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 7 criteria passed. The response handles RequestFailedException, status and service codes, authorization failures, throttling, and SDK retry configuration well, but does not adequately cover deleted-secret detection or soft-delete and purge-protection errors.

### Strengths

- Uses exception filters while retaining a general RequestFailedException fallback.
- Clearly demonstrates Status and ErrorCode extraction and safe diagnostic logging.
- Accurately explains authorization diagnosis, exhausted throttling retries, and the semantics of SetSecretAsync creating a new version.
- Provides practical SecretClientOptions retry settings and warns against logging secret values.

### Issues

- The 404 discussion does not distinguish a never-existing or unavailable secret from one present in the deleted-secrets collection.
- Soft-delete and purge-protection scenarios lack concrete operations, exception handling, and error-code or status guidance.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/7)
      - `RequestFailedException` for all Key Vault errors: Pass
      - Extracting `Status` and `ErrorCode` properties: Pass
      - 403 handling: diagnosing RBAC vs. access policy misconfiguration: Pass
      - 404 handling: secret not found vs. deleted secret: Fail
      - 429 throttling: Key Vault rate limits and retry behavior: Pass
      - Soft-delete and purge protection error scenarios: Fail
      - `SecretClientOptions` retry configuration: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 71% | 1.00 | 0.7143 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.7143** | **71.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
