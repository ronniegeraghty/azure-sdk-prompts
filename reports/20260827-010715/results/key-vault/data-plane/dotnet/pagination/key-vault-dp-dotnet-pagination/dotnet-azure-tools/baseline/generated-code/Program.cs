using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const int PageSizeHint = 100;

if (args.Length != 2 ||
    !Uri.TryCreate(args[0], UriKind.Absolute, out Uri? vaultUri) ||
    (args[1] is not "sync" and not "async"))
{
    Console.Error.WriteLine(
        "Usage: dotnet run -- <https://vault-name.vault.azure.net/> <sync|async>");
    return 1;
}

SecretClient client = new(vaultUri, new DefaultAzureCredential());

try
{
    if (args[1] == "sync")
    {
        ListSecretsSynchronously(client);
    }
    else
    {
        await ListSecretsAsynchronously(client);
    }
}
catch (AuthenticationFailedException exception)
{
    Console.Error.WriteLine($"Authentication failed: {exception.Message}");
    return 2;
}
catch (RequestFailedException exception)
{
    Console.Error.WriteLine(
        $"Key Vault request failed ({exception.Status}, {exception.ErrorCode}): " +
        exception.Message);
    return 3;
}

return 0;

static void ListSecretsSynchronously(SecretClient client)
{
    Pageable<SecretProperties> secrets = client.GetPropertiesOfSecrets();
    int pageNumber = 0;

    foreach (Page<SecretProperties> page in
             secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"Page {pageNumber} ({page.Values.Count} secrets, " +
            $"more pages: {page.ContinuationToken is not null})");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
        }
    }
}

static async Task ListSecretsAsynchronously(SecretClient client)
{
    AsyncPageable<SecretProperties> secrets =
        client.GetPropertiesOfSecretsAsync();
    int pageNumber = 0;

    await foreach (Page<SecretProperties> page in
                   secrets.AsPages(pageSizeHint: PageSizeHint))
    {
        pageNumber++;
        Console.WriteLine(
            $"Page {pageNumber} ({page.Values.Count} secrets, " +
            $"more pages: {page.ContinuationToken is not null})");

        foreach (SecretProperties secret in page.Values)
        {
            PrintSecret(secret);
        }
    }
}

static void PrintSecret(SecretProperties secret)
{
    string contentType = secret.ContentType ?? "(not set)";
    string enabled = secret.Enabled switch
    {
        true => "enabled",
        false => "disabled",
        null => "not specified"
    };

    Console.WriteLine(
        $"  Name: {secret.Name}, Content type: {contentType}, Status: {enabled}");
}
