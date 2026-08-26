using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

namespace KeyVaultSecretPagination;

internal static class Program
{
    private const int PageSizeHint = 25;

    private static async Task<int> Main(string[] args)
    {
        string? vaultUri = Environment.GetEnvironmentVariable("AZURE_KEY_VAULT_URI");
        if (!Uri.TryCreate(vaultUri, UriKind.Absolute, out Uri? parsedVaultUri))
        {
            Console.Error.WriteLine(
                "Set AZURE_KEY_VAULT_URI to a vault URI such as " +
                "https://my-vault.vault.azure.net/.");
            return 1;
        }

        string mode = args.FirstOrDefault()?.ToLowerInvariant() ?? "--async";
        if (mode is not ("--async" or "--sync" or "--both"))
        {
            Console.Error.WriteLine("Usage: dotnet run -- [--async|--sync|--both]");
            return 1;
        }

        var credential = new DefaultAzureCredential();
        var client = new SecretClient(parsedVaultUri, credential);

        try
        {
            if (mode is "--async" or "--both")
            {
                await ListSecretsAsync(client);
            }

            if (mode is "--sync" or "--both")
            {
                ListSecrets(client);
            }

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
                $"Key Vault request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
            return 3;
        }
    }

    private static async Task ListSecretsAsync(SecretClient client)
    {
        Console.WriteLine("Asynchronous page-by-page iteration");

        AsyncPageable<SecretProperties> secrets =
            client.GetPropertiesOfSecretsAsync();

        int pageNumber = 0;
        int secretCount = 0;

        await foreach (Page<SecretProperties> page in
            secrets.AsPages(pageSizeHint: PageSizeHint))
        {
            pageNumber++;
            Console.WriteLine($"\nPage {pageNumber} ({page.Values.Count} secrets)");

            foreach (SecretProperties secret in page.Values)
            {
                PrintSecret(secret);
                secretCount++;
            }

            PrintContinuationToken(page.ContinuationToken);
        }

        Console.WriteLine($"\nAsync total: {secretCount} secrets in {pageNumber} pages.");
    }

    private static void ListSecrets(SecretClient client)
    {
        Console.WriteLine("Synchronous page-by-page iteration");

        Pageable<SecretProperties> secrets = client.GetPropertiesOfSecrets();

        int pageNumber = 0;
        int secretCount = 0;

        foreach (Page<SecretProperties> page in
            secrets.AsPages(pageSizeHint: PageSizeHint))
        {
            pageNumber++;
            Console.WriteLine($"\nPage {pageNumber} ({page.Values.Count} secrets)");

            foreach (SecretProperties secret in page.Values)
            {
                PrintSecret(secret);
                secretCount++;
            }

            PrintContinuationToken(page.ContinuationToken);
        }

        Console.WriteLine($"\nSync total: {secretCount} secrets in {pageNumber} pages.");
    }

    private static void PrintSecret(SecretProperties secret)
    {
        string contentType = secret.ContentType ?? "(none)";
        string enabledStatus = secret.Enabled switch
        {
            true => "Enabled",
            false => "Disabled",
            null => "Not specified"
        };

        Console.WriteLine(
            $"  Name: {secret.Name,-30} Content type: {contentType,-20} " +
            $"Status: {enabledStatus}");
    }

    private static void PrintContinuationToken(string? continuationToken)
    {
        Console.WriteLine(
            continuationToken is null
                ? "  Continuation token: <end>"
                : "  Continuation token: <available>");
    }
}
