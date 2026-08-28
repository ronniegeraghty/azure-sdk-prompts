using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const int PageSizeHint = 50;

string? vaultUrl = Environment.GetEnvironmentVariable("AZURE_KEYVAULT_URL");
if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri))
{
    Console.Error.WriteLine(
        "Set AZURE_KEYVAULT_URL to a valid vault URL, for example " +
        "https://my-vault.vault.azure.net/.");
    return 1;
}

var client = new SecretClient(vaultUri, new DefaultAzureCredential());

try
{
    Console.WriteLine("Synchronous page-by-page iteration");
    ListSecretsSynchronously(client);

    Console.WriteLine();
    Console.WriteLine("Asynchronous page-by-page iteration");
    await ListSecretsAsynchronously(client);
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Key Vault request failed ({ex.Status}, {ex.ErrorCode}): {ex.Message}");
    return 1;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine($"Azure authentication failed: {ex.Message}");
    return 1;
}

return 0;

static void ListSecretsSynchronously(SecretClient client)
{
    Pageable<SecretProperties> secrets = client.GetPropertiesOfSecrets();
    int pageNumber = 0;

    foreach (Page<SecretProperties> page in secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"Page {pageNumber} ({page.Values.Count} items, " +
            $"more pages: {page.ContinuationToken is not null})");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
        }
    }
}

static async Task ListSecretsAsynchronously(SecretClient client)
{
    AsyncPageable<SecretProperties> secrets = client.GetPropertiesOfSecretsAsync();
    int pageNumber = 0;

    await foreach (Page<SecretProperties> page in
        secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"Page {pageNumber} ({page.Values.Count} items, " +
            $"more pages: {page.ContinuationToken is not null})");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
        }
    }
}

static void PrintSecret(SecretProperties secret)
{
    string enabledStatus = secret.Enabled switch
    {
        true => "enabled",
        false => "disabled",
        null => "not set"
    };

    Console.WriteLine(
        $"  Name: {secret.Name}, " +
        $"Content type: {secret.ContentType ?? "(none)"}, " +
        $"Enabled: {enabledStatus}");
}
