using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

internal static class Program
{
    private const string SecretName = "my-secret";

    private static async Task<int> Main()
    {
        string? vaultUrl = Environment.GetEnvironmentVariable("AZURE_KEY_VAULT_URL");
        if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri)
            || vaultUri.Scheme != Uri.UriSchemeHttps)
        {
            Console.Error.WriteLine(
                "Set AZURE_KEY_VAULT_URL to a valid HTTPS vault URL, " +
                "for example https://<vault-name>.vault.azure.net/.");
            return 1;
        }

        try
        {
            var client = new SecretClient(vaultUri, new DefaultAzureCredential());

            KeyVaultSecret createdSecret =
                await client.SetSecretAsync(SecretName, "my-secret-value");
            Console.WriteLine($"Created secret '{createdSecret.Name}'.");

            KeyVaultSecret readSecret = await client.GetSecretAsync(SecretName);
            Console.WriteLine($"Read secret value: {readSecret.Value}");

            KeyVaultSecret updatedSecret =
                await client.SetSecretAsync(SecretName, "updated-value");
            Console.WriteLine($"Updated secret value: {updatedSecret.Value}");

            DeleteSecretOperation deleteOperation =
                await client.StartDeleteSecretAsync(SecretName);
            await deleteOperation.WaitForCompletionAsync();
            Console.WriteLine($"Deleted secret '{SecretName}'.");

            await client.PurgeDeletedSecretAsync(SecretName);
            Console.WriteLine($"Purged secret '{SecretName}'.");

            return 0;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Azure authentication failed: {ex.Message}");
            return 2;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Azure Key Vault request failed (HTTP {ex.Status}, " +
                $"error code '{ex.ErrorCode ?? "unknown"}'): {ex.Message}");
            return 3;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The Azure Key Vault operation was canceled.");
            return 4;
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Unexpected error: {ex.Message}");
            return 5;
        }
    }
}
