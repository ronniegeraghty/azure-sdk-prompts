using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const string SecretName = "my-secret";
const string InitialValue = "my-secret-value";
const string UpdatedValue = "updated-value";

string? keyVaultUrl = Environment.GetEnvironmentVariable("KEY_VAULT_URL");
if (!Uri.TryCreate(keyVaultUrl, UriKind.Absolute, out Uri? keyVaultUri) ||
    keyVaultUri.Scheme != Uri.UriSchemeHttps)
{
    Console.Error.WriteLine(
        "Set KEY_VAULT_URL to an HTTPS Key Vault URL, for example " +
        "https://<vault-name>.vault.azure.net/.");
    return 2;
}

using var cancellationSource = new CancellationTokenSource(TimeSpan.FromMinutes(2));
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellationSource.Cancel();
};

var options = new SecretClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(16),
        MaxRetries = 5
    }
};

var client = new SecretClient(
    keyVaultUri,
    new DefaultAzureCredential(),
    options);

try
{
    Console.WriteLine($"Creating secret '{SecretName}'...");
    await client.SetSecretAsync(
        SecretName,
        InitialValue,
        cancellationSource.Token);

    Response<KeyVaultSecret> getResponse = await client.GetSecretAsync(
        SecretName,
        cancellationToken: cancellationSource.Token);
    Console.WriteLine($"Read secret value: {getResponse.Value.Value}");

    Console.WriteLine($"Updating secret '{SecretName}'...");
    // Setting an existing secret name creates a new version with the new value.
    await client.SetSecretAsync(
        SecretName,
        UpdatedValue,
        cancellationSource.Token);

    Console.WriteLine($"Deleting secret '{SecretName}'...");
    DeleteSecretOperation deleteOperation = await client.StartDeleteSecretAsync(
        SecretName,
        cancellationSource.Token);
    await deleteOperation.WaitForCompletionAsync(cancellationSource.Token);

    Console.WriteLine($"Purging secret '{SecretName}'...");
    await client.PurgeDeletedSecretAsync(
        SecretName,
        cancellationSource.Token);

    Console.WriteLine("CRUD operations completed successfully.");
    return 0;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine(
        $"Authentication failed. Configure a credential supported by " +
        $"DefaultAzureCredential. {ex.Message}");
    return 1;
}
catch (RequestFailedException ex) when (ex.Status == 403)
{
    Console.Error.WriteLine(
        "Access denied. Grant the identity secret get, set, delete, and purge " +
        "permissions (for RBAC vaults, use Key Vault Secrets Officer). " +
        "Purge also fails when purge protection is enabled.");
    Console.Error.WriteLine($"Azure error: {ex.ErrorCode ?? "unknown"}");
    return 1;
}
catch (RequestFailedException ex) when (ex.Status == 409)
{
    Console.Error.WriteLine(
        $"The operation conflicted with the current state of '{SecretName}'. " +
        "A previously deleted secret with this name may still be retained.");
    Console.Error.WriteLine($"Azure error: {ex.ErrorCode ?? "unknown"}");
    return 1;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Key Vault request failed (HTTP {ex.Status}, " +
        $"{ex.ErrorCode ?? "unknown"}): {ex.Message}");
    return 1;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine(
        "The operation was canceled or exceeded the two-minute timeout.");
    return 1;
}
