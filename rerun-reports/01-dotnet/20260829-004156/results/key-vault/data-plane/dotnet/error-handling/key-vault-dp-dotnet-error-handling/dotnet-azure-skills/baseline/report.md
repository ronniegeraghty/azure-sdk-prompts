# Evaluation Report: key-vault-dp-dotnet-error-handling

**Config:** dotnet-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 61.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-error-handling` |
| Config | dotnet-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 4/7 |
| Duration | 61.3s |
| Timestamp | 2026-08-28T18:03:23Z |
| Files Generated | 0 |
| Event Count | 2590 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 32.1s |
| Review | 28.6s |
| **Total** | **61.3s** |

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
| Token Usage | in=19623 out=1543 |
| Turn Count | 1 |

## Copilot Response

`Azure.Security.KeyVault.Secrets` reports service failures as `Azure.RequestFailedException`. Use `Status` for the HTTP status and `ErrorCode` for Key Vault’s machine-readable error code.

```csharp
using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

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
    new Uri("https://my-vault.vault.azure.net/"),
    new DefaultAzureCredential(),
    options);

try
{
    KeyVaultSecret secret = await client.GetSecretAsync(
        name: "database-password",
        version: null); // null means latest version

    Console.WriteLine(secret.Value);
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    // Authentication succeeded, but the identity lacks a required data-plane
    // permission through Key Vault RBAC or a legacy access policy.
    Console.Error.WriteLine(
        $"Key Vault access denied. Status={ex.Status}, " +
        $"Code={ex.ErrorCode}, RequestId={ex.ClientRequestId}");

    throw;
}
catch (RequestFailedException ex) when (
    ex.Status == 404 ||
    string.Equals(ex.ErrorCode, "SecretNotFound",
        StringComparison.OrdinalIgnoreCase))
{
    Console.Error.WriteLine(
        $"Secret was not found. Status={ex.Status}, Code={ex.ErrorCode}");

    // Return an application-specific "not found" result if absence is expected.
}
catch (RequestFailedException ex) when (
    ex.Status == 409 || ex.Status == 412)
{
    Console.Error.WriteLine(
        $"Secret operation conflicted with the current vault state. " +
        $"Status={ex.Status}, Code={ex.ErrorCode}");

    // Re-read the secret or its properties before deciding whether to retry.
    throw;
}
catch (RequestFailedException ex) when (ex.Status == 429)
{
    // This runs only after the Azure SDK exhausts its configured retries.
    Console.Error.WriteLine(
        $"Key Vault throttling persisted after retries. " +
        $"Status={ex.Status}, Code={ex.ErrorCode}, " +
        $"RequestId={ex.ClientRequestId}");

    throw;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Key Vault request failed. Status={ex.Status}, " +
        $"Code={ex.ErrorCode ?? "<none>"}, " +
        $"RequestId={ex.ClientRequestId}, Message={ex.Message}");

    throw;
}
```

A reusable lookup pattern can distinguish “missing” from actual failures:

```csharp
public static async Task<string?> TryGetSecretValueAsync(
    SecretClient client,
    string secretName,
    string? version = null,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response = await client.GetSecretAsync(
            secretName,
            version,
            cancellationToken);

        return response.Value.Value;
    }
    catch (RequestFailedException ex) when (
        ex.Status == 404 &&
        string.Equals(ex.ErrorCode, "SecretNotFound",
            StringComparison.OrdinalIgnoreCase))
    {
        return null;
    }
    catch (RequestFailedException ex) when (ex.Status == 403)
    {
        throw new InvalidOperationException(
            $"The application cannot read secret '{secretName}'. " +
            $"Check Key Vault RBAC roles or access policies. " +
            $"Key Vault error: {ex.ErrorCode}; request: {ex.ClientRequestId}.",
            ex);
    }
}
```

For writes and conflicts:

```csharp
public static async Task<KeyVaultSecret> SetSecretAsync(
    SecretClient client,
    string name,
    string value,
    CancellationToken cancellationToken = default)
{
    try
    {
        Response<KeyVaultSecret> response =
            await client.SetSecretAsync(name, value, cancellationToken);

        return response.Value;
    }
    catch (RequestFailedException ex) when (ex.Status == 409)
    {
        // Examples include a secret name that is still in the deleted/recovery
        // state or another operation conflicting with the vault's current state.
        throw new InvalidOperationException(
            $"Secret '{name}' could not be written because of a conflict. " +
            $"Key Vault error: {ex.ErrorCode}; request: {ex.ClientRequestId}.",
            ex);
    }
    catch (RequestFailedException ex) when (ex.Status == 412)
    {
        // Optimistic-concurrency condition failed, such as stale properties/ETag.
        // Fetch the latest properties, reconcile changes, then retry deliberately.
        throw new InvalidOperationException(
            $"Secret '{name}' changed before the update completed. " +
            $"Key Vault error: {ex.ErrorCode}; request: {ex.ClientRequestId}.",
            ex);
    }
}
```

Important behavior:

- **403:** Do not retry repeatedly; correct the managed identity/service principal’s Key Vault data-plane role or access policy. RBAC propagation may take a short time after assignment.
- **404:** `SecretNotFound` can mean the secret name or requested version does not exist. It can also intentionally conceal a resource from an unauthorized caller in some security contexts, so investigate unexpected 404s.
- **409/412:** Secret values are versioned—`SetSecretAsync` normally creates a new version rather than overwriting one. Conflicts more commonly involve deleted/recoverable secrets or conditional property updates. Re-read state before retrying.
- **429:** Let the SDK honor server retry guidance and exponential backoff. If retries are exhausted, reduce request concurrency, cache secrets, and avoid fetching a secret for every application request.
- **Logging:** Record `Status`, `ErrorCode`, and `ClientRequestId`, but never log `KeyVaultSecret.Value`. Avoid branching only on the exception message because it is not a stable API.

## Reviewer Notes (LLM-as-Judge)

### Summary

4 of 7 criteria passed. The response covers the core exception properties, general error handling, throttling, conflicts, and retry configuration, but lacks complete permission-model diagnosis and deleted-secret/purge-protection handling.

### Strengths

- Uses RequestFailedException filters and a general fallback catch consistently.
- Clearly extracts Status, ErrorCode, and ClientRequestId without logging secret values.
- Accurately explains that SetSecretAsync normally creates a new version and that 409/412 conflicts require state-aware handling.
- Provides practical SDK retry configuration and sound guidance for persistent 429 throttling.

### Issues

- No concrete procedure for distinguishing RBAC mode from access-policy mode when diagnosing 403 responses.
- No deleted-secret lookup pattern to distinguish an absent secret or version from a recoverable soft-deleted secret.
- No purge-protection or recover/purge failure examples.

## Grader Results

- error-handling.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (4/7)
      - `RequestFailedException` for all Key Vault errors: Pass
      - Extracting `Status` and `ErrorCode` properties: Pass
      - 403 handling: diagnosing RBAC vs. access policy misconfiguration: Fail
      - 404 handling: secret not found vs. deleted secret: Fail
      - 429 throttling: Key Vault rate limits and retry behavior: Pass
      - Soft-delete and purge protection error scenarios: Fail
      - `SecretClientOptions` retry configuration: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 57% | 1.00 | 0.5714 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.5714** | **57.1%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-dotnet-error-handling --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
