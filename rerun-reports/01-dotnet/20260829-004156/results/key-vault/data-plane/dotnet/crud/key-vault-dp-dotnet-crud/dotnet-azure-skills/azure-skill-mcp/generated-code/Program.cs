using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

namespace hyoka_key_vault_dp_dotnet_crud_dotnet_azure_skills_azure_skill_mcp_889749088;

internal static class Program
{
    private const string SecretName = "my-secret";

    private static async Task<int> Main()
    {
        string? vaultUrl = Environment.GetEnvironmentVariable("KEY_VAULT_URL");

        if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri) ||
            vaultUri.Scheme != Uri.UriSchemeHttps)
        {
            Console.Error.WriteLine(
                "Set KEY_VAULT_URL to the vault URI, for example https://my-vault.vault.azure.net/.");
            return 1;
        }

        var clientOptions = new SecretClientOptions
        {
            Retry =
            {
                Mode = RetryMode.Exponential,
                Delay = TimeSpan.FromSeconds(1),
                MaxRetries = 5,
                MaxDelay = TimeSpan.FromSeconds(16)
            }
        };

        var client = new SecretClient(
            vaultUri,
            new DefaultAzureCredential(),
            clientOptions);

        try
        {
            Console.WriteLine($"Creating secret '{SecretName}'...");
            await client.SetSecretAsync(SecretName, "my-secret-value");

            Console.WriteLine($"Reading secret '{SecretName}'...");
            Response<KeyVaultSecret> readResponse = await client.GetSecretAsync(SecretName);
            Console.WriteLine($"Secret value: {readResponse.Value.Value}");

            Console.WriteLine($"Updating secret '{SecretName}'...");
            await client.SetSecretAsync(SecretName, "updated-value");

            Console.WriteLine($"Deleting secret '{SecretName}'...");
            DeleteSecretOperation deleteOperation = await client.StartDeleteSecretAsync(SecretName);
            await deleteOperation.WaitForCompletionAsync();

            Console.WriteLine($"Purging secret '{SecretName}'...");
            await client.PurgeDeletedSecretAsync(SecretName);

            Console.WriteLine("All secret CRUD operations completed successfully.");
            return 0;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            return 2;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Key Vault request failed (HTTP {ex.Status}, error code '{ex.ErrorCode ?? "unknown"}'): " +
                ex.Message);
            return 3;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("The Key Vault operation was canceled.");
            return 4;
        }
    }
}
