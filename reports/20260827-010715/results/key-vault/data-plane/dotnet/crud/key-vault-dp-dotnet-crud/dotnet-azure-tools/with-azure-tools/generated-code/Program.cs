using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

namespace KeyVaultSecretsCrud;

internal static class Program
{
    private const string SecretName = "my-secret";
    private const string InitialValue = "my-secret-value";
    private const string UpdatedValue = "updated-value";

    private static async Task<int> Main()
    {
        string? vaultUrl = Environment.GetEnvironmentVariable("KEY_VAULT_URL");
        if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri)
            || vaultUri.Scheme != Uri.UriSchemeHttps)
        {
            Console.Error.WriteLine(
                "Set KEY_VAULT_URL to a valid HTTPS URI, for example " +
                "https://my-vault.vault.azure.net/.");
            return 1;
        }

        using var cancellationSource = new CancellationTokenSource();
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
                MaxDelay = TimeSpan.FromSeconds(8),
                MaxRetries = 5
            }
        };

        var client = new SecretClient(
            vaultUri,
            new DefaultAzureCredential(),
            options);

        try
        {
            Console.WriteLine($"Creating secret '{SecretName}'...");
            await client.SetSecretAsync(
                SecretName,
                InitialValue,
                cancellationSource.Token);

            KeyVaultSecret secret = await client.GetSecretAsync(
                SecretName,
                cancellationToken: cancellationSource.Token);
            Console.WriteLine($"Read secret value: {secret.Value}");

            Console.WriteLine($"Updating secret '{SecretName}'...");
            await client.SetSecretAsync(
                SecretName,
                UpdatedValue,
                cancellationSource.Token);

            KeyVaultSecret updatedSecret = await client.GetSecretAsync(
                SecretName,
                cancellationToken: cancellationSource.Token);
            Console.WriteLine($"Updated secret value: {updatedSecret.Value}");

            Console.WriteLine($"Deleting secret '{SecretName}'...");
            DeleteSecretOperation deleteOperation =
                await client.StartDeleteSecretAsync(
                    SecretName,
                    cancellationSource.Token);

            // Purge is valid only after the soft-delete operation completes.
            await deleteOperation.WaitForCompletionAsync(cancellationSource.Token);

            Console.WriteLine($"Purging secret '{SecretName}'...");
            await client.PurgeDeletedSecretAsync(
                SecretName,
                cancellationSource.Token);

            Console.WriteLine("CRUD operations completed successfully.");
            return 0;
        }
        catch (CredentialUnavailableException exception)
        {
            Console.Error.WriteLine(
                $"No credential was available for DefaultAzureCredential: {exception.Message}");
            return 2;
        }
        catch (AuthenticationFailedException exception)
        {
            Console.Error.WriteLine($"Azure authentication failed: {exception.Message}");
            return 2;
        }
        catch (RequestFailedException exception)
        {
            Console.Error.WriteLine(
                $"Key Vault request failed ({exception.Status}, {exception.ErrorCode}): " +
                exception.Message);

            if (exception.Status == 403)
            {
                Console.Error.WriteLine(
                    "Verify that the identity has secret get, set, delete, and purge permissions. " +
                    "Purge also fails when purge protection is enabled.");
            }

            return 3;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Operation canceled.");
            return 4;
        }
    }
}
