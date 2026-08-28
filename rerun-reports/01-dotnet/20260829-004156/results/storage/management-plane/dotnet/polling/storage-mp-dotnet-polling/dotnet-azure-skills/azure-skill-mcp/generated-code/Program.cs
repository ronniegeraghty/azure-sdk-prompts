using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

try
{
    Settings settings = Settings.FromEnvironment(args);

    var credential = new DefaultAzureCredential();
    var armClient = new ArmClient(credential, settings.SubscriptionId);

    ResourceIdentifier resourceGroupId = ResourceGroupResource.CreateResourceIdentifier(
        settings.SubscriptionId,
        settings.ResourceGroupName);
    ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);
    StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

    var content = new StorageAccountCreateOrUpdateContent(
        new StorageSku(StorageSkuName.StandardLrs),
        StorageKind.StorageV2,
        new AzureLocation(settings.Location))
    {
        AccessTier = StorageAccountAccessTier.Hot,
        AllowBlobPublicAccess = false,
        MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2,
        EnableHttpsTrafficOnly = true
    };

    Console.WriteLine(
        "Starting creation of '{0}' in resource group '{1}'...",
        settings.AccountName,
        settings.ResourceGroupName);

    // Started returns after the service accepts the request instead of hiding the LRO.
    ArmOperation<StorageAccountResource> operation =
        await storageAccounts.CreateOrUpdateAsync(
            WaitUntil.Started,
            settings.AccountName,
            content);

    PrintStatus("Started", operation);

    StorageAccountResource account = settings.Mode switch
    {
        Settings.AutomaticMode => await WaitAutomaticallyAsync(operation, settings),
        Settings.ManualMode => await PollManuallyAsync(operation, settings),
        _ => throw new InvalidOperationException($"Unsupported mode '{settings.Mode}'.")
    };

    Console.WriteLine("Storage account ready: {0}", account.Id);
    return 0;
}
catch (OperationCanceledException)
{
    Console.Error.WriteLine(
        "Polling timed out. Stopping the client-side wait does not cancel the Azure operation; " +
        "the service may still finish creating the account.");
    return 2;
}
catch (AuthenticationFailedException ex)
{
    Console.Error.WriteLine("Authentication failed: {0}", ex.Message);
    return 3;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        "Azure request failed ({0}, {1}): {2}",
        ex.Status,
        ex.ErrorCode ?? "no error code",
        ex.Message);
    return 4;
}
catch (ArgumentException ex)
{
    Console.Error.WriteLine("Configuration error: {0}", ex.Message);
    return 5;
}

static async Task<StorageAccountResource> WaitAutomaticallyAsync(
    ArmOperation<StorageAccountResource> operation,
    Settings settings)
{
    Console.WriteLine(
        "Automatic mode: WaitForCompletionAsync polls internally every {0}.",
        settings.PollInterval);

    using var timeout = new CancellationTokenSource(settings.Timeout);
    Response<StorageAccountResource> completed =
        await operation.WaitForCompletionAsync(settings.PollInterval, timeout.Token);

    PrintStatus("Completed", operation);
    return completed.Value;
}

static async Task<StorageAccountResource> PollManuallyAsync(
    ArmOperation<StorageAccountResource> operation,
    Settings settings)
{
    Console.WriteLine(
        "Manual mode: the application calls UpdateStatusAsync and controls logging and delay.");

    using var timeout = new CancellationTokenSource(settings.Timeout);

    while (!operation.HasCompleted)
    {
        Response response = await operation.UpdateStatusAsync(timeout.Token);
        PrintStatus($"Polled (HTTP {response.Status})", operation);

        if (!operation.HasCompleted)
        {
            await Task.Delay(settings.PollInterval, timeout.Token);
        }
    }

    if (!operation.HasValue)
    {
        throw new InvalidOperationException(
            "The operation completed without producing a storage account.");
    }

    return operation.Value;
}

static void PrintStatus(
    string stage,
    ArmOperation<StorageAccountResource> operation)
{
    Response lastResponse = operation.GetRawResponse();
    Console.WriteLine(
        "[{0:O}] {1}: HasCompleted={2}, HasValue={3}, LastHttpStatus={4}, OperationId={5}",
        DateTimeOffset.UtcNow,
        stage,
        operation.HasCompleted,
        operation.HasValue,
        lastResponse.Status,
        operation.Id);
}

internal sealed record Settings(
    string SubscriptionId,
    string ResourceGroupName,
    string AccountName,
    string Location,
    string Mode,
    TimeSpan Timeout,
    TimeSpan PollInterval)
{
    public const string AutomaticMode = "automatic";
    public const string ManualMode = "manual";

    public static Settings FromEnvironment(string[] args)
    {
        string mode = ReadMode(args);

        return new Settings(
            Required("AZURE_SUBSCRIPTION_ID"),
            Required("AZURE_RESOURCE_GROUP"),
            Required("AZURE_STORAGE_ACCOUNT_NAME"),
            Environment.GetEnvironmentVariable("AZURE_LOCATION") ?? "eastus",
            mode,
            ReadPositiveSeconds("LRO_TIMEOUT_SECONDS", 600),
            ReadPositiveSeconds("LRO_POLL_SECONDS", 10));
    }

    private static string ReadMode(string[] args)
    {
        int modeIndex = Array.IndexOf(args, "--mode");
        string mode = modeIndex >= 0 && modeIndex + 1 < args.Length
            ? args[modeIndex + 1].ToLowerInvariant()
            : AutomaticMode;

        if (mode is not (AutomaticMode or ManualMode))
        {
            throw new ArgumentException(
                $"--mode must be '{AutomaticMode}' or '{ManualMode}'.");
        }

        return mode;
    }

    private static string Required(string name) =>
        Environment.GetEnvironmentVariable(name) is { Length: > 0 } value
            ? value
            : throw new ArgumentException($"Set the {name} environment variable.");

    private static TimeSpan ReadPositiveSeconds(string name, int defaultValue)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        if (value is null)
        {
            return TimeSpan.FromSeconds(defaultValue);
        }

        if (!int.TryParse(value, out int seconds) || seconds <= 0)
        {
            throw new ArgumentException($"{name} must be a positive integer.");
        }

        return TimeSpan.FromSeconds(seconds);
    }
}
