using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const string secretName = "my-secret";
const string initialValue = "my-secret-value";
const string updatedValue = "updated-value";

string? vaultUrl = Environment.GetEnvironmentVariable("AZURE_KEY_VAULT_URL");
if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri) ||
    !string.Equals(vaultUri.Scheme, Uri.UriSchemeHttps, StringComparison.OrdinalIgnoreCase))
{
    Console.Error.WriteLine(
        "Set AZURE_KEY_VAULT_URL to a valid HTTPS vault URL, " +
        "for example https://<vault-name>.vault.azure.net/.");
    return 2;
}

var credential = new DefaultAzureCredential();
var client = new SecretClient(vaultUri, credential);

try
{
    KeyVaultSecret createdSecret = await client.SetSecretAsync(secretName, initialValue);
    Console.WriteLine($"Created secret '{createdSecret.Name}'.");

    KeyVaultSecret readSecret = await client.GetSecretAsync(secretName);
    Console.WriteLine($"Read secret value: {readSecret.Value}");

    KeyVaultSecret updatedSecret = await client.SetSecretAsync(secretName, updatedValue);
    Console.WriteLine($"Updated secret value: {updatedSecret.Value}");

    DeleteSecretOperation deleteOperation = await client.StartDeleteSecretAsync(secretName);
    await deleteOperation.WaitForCompletionAsync();
    Console.WriteLine($"Deleted secret '{secretName}'.");

    await client.PurgeDeletedSecretAsync(secretName);
    Console.WriteLine($"Purged secret '{secretName}'.");

    return 0;
}
catch (CredentialUnavailableException ex)
{
    Console.Error.WriteLine($"No credential is available: {ex.Message}");
    return 3;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Authentication failed: {ex.Message}");
    return 4;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Key Vault request failed (HTTP {ex.Status}, {ex.ErrorCode ?? "no error code"}): " +
        ex.Message);
    return 5;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine("The Key Vault operation was canceled.");
    return 6;
}
