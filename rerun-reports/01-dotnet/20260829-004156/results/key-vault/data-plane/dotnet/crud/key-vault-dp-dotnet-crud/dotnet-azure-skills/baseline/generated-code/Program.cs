using Azure;
using Azure.Identity;
using Azure.Security.KeyVault.Secrets;

const string SecretName = "my-secret";
const string InitialValue = "my-secret-value";
const string UpdatedValue = "updated-value";

string? vaultUriValue = Environment.GetEnvironmentVariable("KEY_VAULT_URI");

if (!Uri.TryCreate(vaultUriValue, UriKind.Absolute, out Uri? vaultUri) ||
    vaultUri.Scheme != Uri.UriSchemeHttps)
{
    Console.Error.WriteLine(
        "Set KEY_VAULT_URI to a valid HTTPS vault URI, for example " +
        "https://your-vault-name.vault.azure.net/.");
    return 1;
}

var client = new SecretClient(vaultUri, new DefaultAzureCredential());

try
{
    Console.WriteLine($"Creating secret '{SecretName}'...");
    await client.SetSecretAsync(SecretName, InitialValue);

    Console.WriteLine($"Reading secret '{SecretName}'...");
    KeyVaultSecret secret = await client.GetSecretAsync(SecretName);
    Console.WriteLine($"Secret value: {secret.Value}");

    Console.WriteLine($"Updating secret '{SecretName}'...");
    KeyVaultSecret updatedSecret =
        await client.SetSecretAsync(SecretName, UpdatedValue);
    Console.WriteLine($"Updated secret value: {updatedSecret.Value}");

    Console.WriteLine($"Deleting secret '{SecretName}'...");
    DeleteSecretOperation deleteOperation =
        await client.StartDeleteSecretAsync(SecretName);
    await deleteOperation.WaitForCompletionAsync();

    Console.WriteLine($"Purging secret '{SecretName}'...");
    await client.PurgeDeletedSecretAsync(SecretName);

    Console.WriteLine("All CRUD operations completed successfully.");
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
        $"Azure Key Vault request failed ({ex.Status}, {ex.ErrorCode}): " +
        ex.Message);
    return 3;
}
catch (Exception ex)
{
    Console.Error.WriteLine($"Unexpected error: {ex.Message}");
    return 4;
}
