using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const int PageSizeHint = 50;

if (!TryGetVaultUri(out Uri? vaultUri))
{
    return 1;
}

string mode = args.FirstOrDefault()?.ToLowerInvariant() ?? "async";
if (mode is not ("sync" or "async" or "both"))
{
    Console.Error.WriteLine("Usage: dotnet run -- [sync|async|both]");
    return 1;
}

using var cancellationSource = new CancellationTokenSource();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellationSource.Cancel();
};

var clientOptions = new SecretClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 5,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(16)
    }
};

var client = new SecretClient(
    vaultUri,
    new DefaultAzureCredential(),
    clientOptions);

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
catch (OperationCanceledException)
{
    Console.Error.WriteLine("Secret listing was canceled.");
    return 2;
}
catch (AuthenticationFailedException exception)
{
    Console.Error.WriteLine($"Authentication failed: {exception.Message}");
    return 3;
}
catch (RequestFailedException exception)
{
    Console.Error.WriteLine(
        $"Key Vault request failed ({exception.Status}, {exception.ErrorCode}): {exception.Message}");
    return 4;
}

static void ListSecretsSynchronously(SecretClient client, CancellationToken cancellationToken)
{
    Console.WriteLine("Synchronous page iteration");

    Pageable<SecretProperties> secrets =
        client.GetPropertiesOfSecrets(cancellationToken);

    var pageNumber = 0;
    var secretCount = 0;

    foreach (Page<SecretProperties> page in secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine($"\nPage {pageNumber} ({page.Values.Count} secrets)");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
            secretCount++;
        }

        Console.WriteLine(
            $"Continuation token: {FormatContinuationToken(page.ContinuationToken)}");
    }

    Console.WriteLine($"\nSync total: {secretCount} secrets in {pageNumber} pages.");
}

static async Task ListSecretsAsynchronously(
    SecretClient client,
    CancellationToken cancellationToken)
{
    Console.WriteLine("Asynchronous page iteration");

    AsyncPageable<SecretProperties> secrets =
        client.GetPropertiesOfSecretsAsync(cancellationToken);

    var pageNumber = 0;
    var secretCount = 0;

    await foreach (Page<SecretProperties> page in secrets
        .AsPages(pageSizeHint: PageSizeHint)
        .WithCancellation(cancellationToken))
    {
        pageNumber++;
        Console.WriteLine($"\nPage {pageNumber} ({page.Values.Count} secrets)");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
            secretCount++;
        }

        Console.WriteLine(
            $"Continuation token: {FormatContinuationToken(page.ContinuationToken)}");
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
        $"Name: {secret.Name}, Content type: {contentType}, Status: {enabledStatus}");
}

static string FormatContinuationToken(string? continuationToken) =>
    continuationToken is null ? "(end)" : "(present; another page follows)";

static bool TryGetVaultUri(out Uri? vaultUri)
{
    string? value = Environment.GetEnvironmentVariable("KEY_VAULT_URL");

    if (Uri.TryCreate(value, UriKind.Absolute, out vaultUri) &&
        vaultUri.Scheme == Uri.UriSchemeHttps)
    {
        return true;
    }

    Console.Error.WriteLine(
        "Set KEY_VAULT_URL to an HTTPS vault URL, for example " +
        "https://my-vault.vault.azure.net/.");
    return false;
}
