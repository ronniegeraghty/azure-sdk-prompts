using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

namespace KeyVaultSecretPagination;

internal static class Program
{
    private const int PageSizeHint = 25;

    public static async Task<int> Main(string[] args)
    {
        string? vaultUrl = Environment.GetEnvironmentVariable("AZURE_KEY_VAULT_URL");
        if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri) ||
            vaultUri.Scheme != Uri.UriSchemeHttps)
        {
            Console.Error.WriteLine(
                "Set AZURE_KEY_VAULT_URL to an HTTPS vault URL, for example " +
                "https://my-vault.vault.azure.net/.");
            return 2;
        }

        string mode = args.FirstOrDefault()?.ToLowerInvariant() ?? "async";
        if (mode is not ("sync" or "async" or "both"))
        {
            Console.Error.WriteLine("Usage: dotnet run -- [sync|async|both]");
            return 2;
        }

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

        var client = new SecretClient(vaultUri, new DefaultAzureCredential(), options);
        using var cancellationSource = new CancellationTokenSource();
        Console.CancelKeyPress += (_, eventArgs) =>
        {
            eventArgs.Cancel = true;
            cancellationSource.Cancel();
        };

        try
        {
            if (mode is "sync" or "both")
            {
                ListSecretsSynchronously(client, cancellationSource.Token);
            }

            if (mode is "async" or "both")
            {
                await ListSecretsAsynchronously(client, cancellationSource.Token);
            }

            return 0;
        }
        catch (AuthenticationFailedException ex)
        {
            Console.Error.WriteLine($"Authentication failed: {ex.Message}");
            return 1;
        }
        catch (RequestFailedException ex)
        {
            Console.Error.WriteLine(
                $"Key Vault request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
            return 1;
        }
        catch (OperationCanceledException)
        {
            Console.Error.WriteLine("Listing canceled.");
            return 130;
        }
    }

    private static void ListSecretsSynchronously(
        SecretClient client,
        CancellationToken cancellationToken)
    {
        Console.WriteLine("Synchronous Pageable<SecretProperties> iteration");

        Pageable<SecretProperties> secrets =
            client.GetPropertiesOfSecrets(cancellationToken);

        int pageNumber = 0;
        foreach (Page<SecretProperties> page in
                 secrets.AsPages(pageSizeHint: PageSizeHint))
        {
            PrintPage(++pageNumber, page);
        }
    }

    private static async Task ListSecretsAsynchronously(
        SecretClient client,
        CancellationToken cancellationToken)
    {
        Console.WriteLine("Asynchronous AsyncPageable<SecretProperties> iteration");

        AsyncPageable<SecretProperties> secrets =
            client.GetPropertiesOfSecretsAsync(cancellationToken);

        int pageNumber = 0;
        await foreach (Page<SecretProperties> page in
                       secrets.AsPages(pageSizeHint: PageSizeHint)
                           .WithCancellation(cancellationToken))
        {
            PrintPage(++pageNumber, page);
        }
    }

    private static void PrintPage(
        int pageNumber,
        Page<SecretProperties> page)
    {
        Console.WriteLine(
            $"\nPage {pageNumber}: {page.Values.Count} secret(s), " +
            $"continuation token: {page.ContinuationToken ?? "<end>"}");

        foreach (SecretProperties secret in page.Values)
        {
            string contentType = secret.ContentType ?? "<not set>";
            string enabledStatus = secret.Enabled switch
            {
                true => "enabled",
                false => "disabled",
                null => "not specified"
            };

            Console.WriteLine(
                $"  Name: {secret.Name}, Content type: {contentType}, " +
                $"Enabled: {enabledStatus}");
        }
    }
}
