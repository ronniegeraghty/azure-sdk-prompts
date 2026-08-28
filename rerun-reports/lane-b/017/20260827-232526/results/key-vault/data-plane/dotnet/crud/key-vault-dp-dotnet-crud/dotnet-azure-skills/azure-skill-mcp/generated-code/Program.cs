using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

namespace KeyVaultSecretCrud;

internal static class Program
{
    private const string SecretName = "my-secret";

    private static async Task<int> Main()
    {
        string? vaultUrl = Environment.GetEnvironmentVariable("KEY_VAULT_URL");
        if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri))
        {
            Console.Error.WriteLine(
                "Set KEY_VAULT_URL to the vault URI, for example https://my-vault.vault.azure.net/.");
            return 1;
        }

        using var cancellationSource = new CancellationTokenSource();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        try
        {
            var client = new SecretClient(vaultUri, new DefaultAzureCredential());

            KeyVaultSecret createdSecret = await client.SetSecretAsync(
                SecretName,
                "my-secret-value",
                cancellationSource.Token);
            Console.WriteLine($"Created secret '{createdSecret.Name}'.");

            KeyVaultSecret readSecret = await client.GetSecretAsync(
                SecretName,
                cancellationToken: cancellationSource.Token);
            Console.WriteLine($"Secret value: {readSecret.Value}");

            KeyVaultSecret updatedSecret = await client.SetSecretAsync(
                SecretName,
                "updated-value",
                cancellationSource.Token);
            Console.WriteLine($"Updated secret '{updatedSecret.Name}' to '{updatedSecret.Value}'.");

            DeleteSecretOperation deleteOperation = await client.StartDeleteSecretAsync(
                SecretName,
                cancellationSource.Token);
            await deleteOperation.WaitForCompletionAsync(cancellationSource.Token);
            Console.WriteLine($"Deleted secret '{SecretName}'.");

            await client.PurgeDeletedSecretAsync(SecretName, cancellationSource.Token);
            Console.WriteLine($"Purged secret '{SecretName}'.");

            return 0;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            Console.Error.WriteLine(
                "Sign in with Azure CLI, Visual Studio, or another credential supported by DefaultAzureCredential.");
            return 2;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Key Vault request failed (HTTP {ex.Status}, code '{ex.ErrorCode ?? "unknown"}'): {ex.Message}");
            return 3;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Operation canceled.");
            return 4;
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Unexpected error: {ex.Message}");
            return 5;
        }
    }
}
