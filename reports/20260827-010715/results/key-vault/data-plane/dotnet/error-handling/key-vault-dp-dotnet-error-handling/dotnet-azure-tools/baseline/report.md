# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 166.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 5/7 |
| Duration | 166.9s |
| Timestamp | 2026-08-26T17:51:15Z |
| Files Generated | 0 |
| Event Count | 2736 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 121.6s |
| Review | 44.2s |
| **Total** | **166.9s** |

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
| Token Usage | in=67174 out=2322 |
| Turn Count | 3 |

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

### 🔧 1. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure.Security.KeyVault.Secrets .NET RequestFailedException Status ErrorCode 403 404 409 429 retry official documentation"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure Key Vault REST API secret error codes 409 conflict secret version .NET Azure.Security.KeyVault.Secrets official"}
```

### 🔧 3. web_fetch

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.requestfailedexception?view=azure-dotnet"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/azure/key-vault/general/common-error-codes"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.retryoptions?view=azure-dotnet"}
```

## Copilot Response

`Azure.Security.KeyVault.Secrets` surfaces service failures as `Azure.RequestFailedException`. Use `Status` for the HTTP status and `ErrorCode` for Key Vault’s service-specific code.

```csharp
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

var options = new SecretClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(30),
        MaxRetries = 5,
        NetworkTimeout = TimeSpan.FromSeconds(30)
    }
};

var client = new SecretClient(
    new Uri("https://my-vault.vault.azure.net/"),
    new DefaultAzureCredential(),
    options);

try
{
    KeyVaultSecret secret =
        await client.GetSecretAsync("database-password", cancellationToken: cancellationToken);

    // Use secret.Value, but never log it.
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // ErrorCode is commonly "Forbidden"; the message may identify
    // ForbiddenByRbac, ForbiddenByPolicy, or ForbiddenByFirewall.
    LogKeyVaultFailure("Access denied", ex);

    // Usually not transient. Correct RBAC/access policies/firewall settings.
    throw new UnauthorizedAccessException(
        $"Not authorized to read the Key Vault secret. Error code: {ex.ErrorCode}.",
        ex);
}
catch (RequestFailedException ex) when (ex.Status == 404)
{
    // Common ErrorCode: "SecretNotFound".
    LogKeyVaultFailure("Secret or secret version not found", ex);

    // Return a domain-specific result, use a fallback, or rethrow.
    throw new InvalidOperationException(
        $"The requested secret does not exist. Error code: {ex.ErrorCode}.",
        ex);
}
catch (RequestFailedException ex) when (ex.Status == 409)
{
    // Common ErrorCode: "Conflict" or "ConflictError".
    LogKeyVaultFailure("Conflicting Key Vault operation", ex);

    // Sequence concurrent operations on the same secret and retry only when
    // the operation is safe/idempotent.
    throw;
}
catch (RequestFailedException ex) when (ex.Status == 429)
{
    // This is reached only after the Azure SDK exhausts its configured retries.
    LogKeyVaultFailure("Key Vault throttling limit exceeded", ex);

    // Propagate so the caller can shed load, queue the operation, or fail cleanly.
    throw;
}
catch (RequestFailedException ex)
{
    LogKeyVaultFailure("Unexpected Key Vault service failure", ex);
    throw;
}
catch (AuthenticationFailedException ex)
{
    // Credential acquisition failed before a successful Key Vault request.
    // Examples: unavailable managed identity or invalid local credentials.
    Console.Error.WriteLine($"Azure authentication failed: {ex.Message}");
    throw;
}
```

A logging helper can safely extract the status, service error code, request ID, and retry information:

```csharp
static void LogKeyVaultFailure(string description, RequestFailedException ex)
{
    Response? response = ex.GetRawResponse();

    string? requestId = null;
    string? retryAfter = null;

    if (response is not null)
    {
        response.Headers.TryGetValue("x-ms-request-id", out requestId);
        response.Headers.TryGetValue("Retry-After", out retryAfter);
    }

    Console.Error.WriteLine(
        "{0}: HTTP {1}, ErrorCode={2}, RequestId={3}, RetryAfter={4}, Message={5}",
        description,
        ex.Status,                    // 403, 404, 409, 429, etc.; 0 if no response
        ex.ErrorCode ?? "<none>",     // Key Vault service-specific code
        requestId ?? "<none>",
        retryAfter ?? "<none>",
        ex.Message);
}
```

For an optional secret, handling `404` as an expected result is often cleaner:

```csharp
static async Task<string?> GetOptionalSecretAsync(
    SecretClient client,
    string name,
    string? version = null,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.GetSecretAsync(name, version, cancellationToken);

        return response.Value.Value;
    }
    catch (RequestFailedException ex) when (
        ex.Status == 404 &&
        string.Equals(ex.ErrorCode, "SecretNotFound",
            StringComparison.OrdinalIgnoreCase))
    {
        return null;
    }
}
```

For writes, note that secret values are versioned:

```csharp
try
{
    // This creates a new version when the secret name already exists.
    KeyVaultSecret created =
        await client.SetSecretAsync("api-key", newValue, cancellationToken);
}
catch (RequestFailedException ex) when (ex.Status == 409)
{
    // Possible overlapping delete/recover/purge/write operations.
    // Avoid concurrent operations on the same secret name.
    LogKeyVaultFailure("Secret operation conflict", ex);
    throw;
}
```

A secret version’s **value is immutable**—`SetSecretAsync` creates a new version rather than updating an existing version. Metadata can be changed through `UpdateSecretPropertiesAsync`. A `409` generally represents concurrent or incompatible operations on the same item, not a normal “version already exists” result.

For `429`, prefer the client’s built-in retry policy. It honors `Retry-After` when Key Vault supplies it and uses exponential backoff otherwise. If `429` still reaches your catch block, retries are exhausted; avoid immediately adding another unbounded retry loop. Reduce request volume, cache secrets in memory for an appropriate period, and use a bounded higher-level retry only where the operation is safe.

Also, treat `Status` as the primary classification. `ErrorCode` provides useful diagnostics but can vary—for example, a `403` may contain `Forbidden` at the top level while the message or inner service details indicate RBAC, policy, or firewall denial.

References: [RequestFailedException](https://learn.microsoft.com/dotnet/api/azure.requestfailedexception), [RetryOptions](https://learn.microsoft.com/dotnet/api/azure.core.retryoptions), and [Key Vault common error codes](https://learn.microsoft.com/azure/key-vault/general/common-error-codes).

## Reviewer Notes (LLM-as-Judge)

### Summary

5 of 7 criteria passed. The response thoroughly covers RequestFailedException classification, diagnostics, conflicts, throttling, and SDK retry configuration, but omits substantive deleted-secret and purge-protection handling.

### Strengths

- Clear status-filtered RequestFailedException catch patterns for 403, 404, 409, and 429.
- Correct extraction and logging of Status, ErrorCode, request ID, and retry metadata.
- Accurate explanation that SetSecretAsync creates a new version and that 429 normally reaches user code only after configured retries are exhausted.
- Concrete SecretClientOptions retry configuration.

### Issues

- No deleted-secret lookup or explanation for distinguishing a soft-deleted secret from an ordinary 404.
- No handling or diagnosis of soft-delete and purge-protection error scenarios.

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
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
