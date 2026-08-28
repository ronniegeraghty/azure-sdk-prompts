using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const int PageSizeHint = 100;

if (args.Length is < 1 or > 2)
{
    PrintUsage();
    return 1;
}

if (!Uri.TryCreate(args[0], UriKind.Absolute, out Uri? vaultUri) ||
    vaultUri.Scheme != Uri.UriSchemeHttps)
{
    Console.Error.WriteLine("The vault URL must be an absolute HTTPS URL.");
    PrintUsage();
    return 1;
}

string mode = args.Length == 2 ? args[1].ToLowerInvariant() : "--both";
if (mode is not ("--sync" or "--async" or "--both"))
{
    Console.Error.WriteLine($"Unknown mode: {mode}");
    PrintUsage();
    return 1;
}

var client = new SecretClient(vaultUri, new DefaultAzureCredential());

try
{
    if (mode is "--sync" or "--both")
    {
        ListSecretsSynchronously(client);
    }

    if (mode == "--both")
    {
        Console.WriteLine();
    }

    if (mode is "--async" or "--both")
    {
        await ListSecretsAsynchronously(client);
    }
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

return 0;

static void ListSecretsSynchronously(SecretClient client)
{
    Console.WriteLine("Synchronous page iteration");

    // Pageable<T>.AsPages() exposes each service response as Page<T>.
    Pageable<SecretProperties> secrets = client.GetPropertiesOfSecrets();
    int pageNumber = 0;

    foreach (Page<SecretProperties> page in secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"Page {pageNumber}: {page.Values.Count} item(s), " +
            $"more pages: {page.ContinuationToken is not null}");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
        }
    }
}

static async Task ListSecretsAsynchronously(SecretClient client)
{
    Console.WriteLine("Asynchronous page iteration");

    // AsyncPageable<T> fetches the next page only as await foreach requests it.
    AsyncPageable<SecretProperties> secrets = client.GetPropertiesOfSecretsAsync();
    int pageNumber = 0;

    await foreach (Page<SecretProperties> page in
        secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"Page {pageNumber}: {page.Values.Count} item(s), " +
            $"more pages: {page.ContinuationToken is not null}");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
        }
    }
}

static void PrintSecret(SecretProperties secret)
{
    string contentType = string.IsNullOrWhiteSpace(secret.ContentType)
        ? "(not set)"
        : secret.ContentType;

    string enabledStatus = secret.Enabled switch
    {
        true => "Enabled",
        false => "Disabled",
        null => "Not specified"
    };

    Console.WriteLine(
        $"  Name: {secret.Name}, Content type: {contentType}, Status: {enabledStatus}");
}

static void PrintUsage()
{
    Console.Error.WriteLine(
        "Usage: dotnet run -- <https://vault-name.vault.azure.net/> " +
        "[--sync|--async|--both]");
}
