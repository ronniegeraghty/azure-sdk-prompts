using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const int PageSizeHint = 25;

string? vaultUrl = Environment.GetEnvironmentVariable("AZURE_KEY_VAULT_URL");
if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri))
{
    Console.Error.WriteLine(
        "Set AZURE_KEY_VAULT_URL to an absolute vault URL, for example " +
        "https://my-vault.vault.azure.net/.");
    return 1;
}

IterationMode mode;
try
{
    mode = ParseMode(args);
}
catch (ArgumentException exception)
{
    Console.Error.WriteLine(exception.Message);
    return 2;
}

TokenCredential credential = new DefaultAzureCredential();
SecretClient client = new(vaultUri, credential);

try
{
    if (mode is IterationMode.Sync or IterationMode.Both)
    {
        ListSecretsSynchronously(client);
    }

    if (mode is IterationMode.Async or IterationMode.Both)
    {
        await ListSecretsAsynchronously(client);
    }
}
catch (AuthenticationFailedException exception)
{
    Console.Error.WriteLine($"Authentication failed: {exception.Message}");
    return 3;
}
catch (RequestFailedException exception)
{
    Console.Error.WriteLine(
        $"Key Vault request failed ({exception.Status}, {exception.ErrorCode}): " +
        exception.Message);
    return 4;
}

return 0;

static void ListSecretsSynchronously(SecretClient client)
{
    Console.WriteLine("Synchronous page-by-page iteration");

    Pageable<SecretProperties> secrets = client.GetPropertiesOfSecrets();
    int pageNumber = 0;
    int secretCount = 0;

    foreach (Page<SecretProperties> page in secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"\nPage {pageNumber} ({page.Values.Count} secrets, " +
            $"has next page: {page.ContinuationToken is not null})");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
            secretCount++;
        }
    }

    Console.WriteLine($"\nSync total: {secretCount} secrets in {pageNumber} pages.");
}

static async Task ListSecretsAsynchronously(SecretClient client)
{
    Console.WriteLine("\nAsynchronous page-by-page iteration");

    AsyncPageable<SecretProperties> secrets =
        client.GetPropertiesOfSecretsAsync();
    int pageNumber = 0;
    int secretCount = 0;

    await foreach (Page<SecretProperties> page in
        secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"\nPage {pageNumber} ({page.Values.Count} secrets, " +
            $"has next page: {page.ContinuationToken is not null})");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
            secretCount++;
        }
    }

    Console.WriteLine($"\nAsync total: {secretCount} secrets in {pageNumber} pages.");
}

static void PrintSecret(SecretProperties secret)
{
    string contentType = string.IsNullOrWhiteSpace(secret.ContentType)
        ? "(not set)"
        : secret.ContentType;
    string enabledStatus = secret.Enabled switch
    {
        true => "enabled",
        false => "disabled",
        null => "not specified"
    };

    Console.WriteLine(
        $"  Name: {secret.Name,-32} " +
        $"Content type: {contentType,-24} " +
        $"Status: {enabledStatus}");
}

static IterationMode ParseMode(string[] args)
{
    if (args.Length == 0)
    {
        return IterationMode.Both;
    }

    if (args.Length != 2 || args[0] != "--mode")
    {
        throw new ArgumentException("Usage: dotnet run -- [--mode sync|async|both]");
    }

    return args[1].ToLowerInvariant() switch
    {
        "sync" => IterationMode.Sync,
        "async" => IterationMode.Async,
        "both" => IterationMode.Both,
        _ => throw new ArgumentException(
            "The --mode value must be sync, async, or both.")
    };
}

enum IterationMode
{
    Sync,
    Async,
    Both
}
