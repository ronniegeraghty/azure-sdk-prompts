using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

const string ExecuteFlag = "--execute";

if (!args.Contains(ExecuteFlag, StringComparer.OrdinalIgnoreCase))
{
    Console.WriteLine(
        """
        Dry run only; no Azure request was sent.

        To run this sample against Azure, set:
          AZURE_SUBSCRIPTION_ID
          AZURE_RESOURCE_GROUP
          AZURE_STORAGE_ACCOUNT_NAME
          AZURE_LOCATION (optional; defaults to eastus)
          LRO_TIMEOUT_MINUTES (optional; defaults to 10)

        Then choose one polling strategy:
          dotnet run -- --execute wait
          dotnet run -- --execute manual
        """);
    return;
}

string subscriptionId = GetRequiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
string resourceGroupName = GetRequiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
string accountName = GetRequiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");
string location = Environment.GetEnvironmentVariable("AZURE_LOCATION") ?? "eastus";
TimeSpan timeout = TimeSpan.FromMinutes(GetPositiveDouble("LRO_TIMEOUT_MINUTES", 10));
TimeSpan pollingInterval = TimeSpan.FromSeconds(10);
string pollingMode = args.FirstOrDefault(
    argument => !argument.Equals(ExecuteFlag, StringComparison.OrdinalIgnoreCase)) ?? "wait";

using CancellationTokenSource timeoutSource = new(timeout);
ArmOperation<StorageAccountResource>? operation = null;

try
{
    ArmClient armClient = new(new DefaultAzureCredential());
    ResourceIdentifier resourceGroupId =
        ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
    ResourceGroupResource resourceGroup = armClient.GetResourceGroupResource(resourceGroupId);
    StorageAccountCollection storageAccounts = resourceGroup.GetStorageAccounts();

    StorageAccountCreateOrUpdateContent parameters = new(
        new StorageSku(StorageSkuName.StandardLrs),
        StorageKind.StorageV2,
        new AzureLocation(location));

    // WaitUntil.Started returns as soon as Azure accepts the request, exposing the LRO.
    operation = await storageAccounts.CreateOrUpdateAsync(
        WaitUntil.Started,
        accountName,
        parameters,
        timeoutSource.Token);

    Console.WriteLine(
        $"Started operation {operation.Id}; completed={operation.HasCompleted}; " +
        $"HTTP status={operation.GetRawResponse().Status}.");

    StorageAccountResource account = pollingMode.ToLowerInvariant() switch
    {
        "wait" => await CompleteWithSdkPollingAsync(
            operation,
            pollingInterval,
            timeoutSource.Token),
        "manual" => await CompleteWithManualPollingAsync(
            operation,
            pollingInterval,
            timeoutSource.Token),
        _ => throw new ArgumentException("Polling mode must be either 'wait' or 'manual'.")
    };

    Console.WriteLine($"Storage account created: {account.Data.Id}");
}
catch (OperationCanceledException) when (timeoutSource.IsCancellationRequested)
{
    Console.Error.WriteLine(
        $"Timed out after {timeout}. The client stopped polling, but Azure may still be " +
        $"processing operation {operation?.Id ?? "(not yet returned)"}.");
    Environment.ExitCode = 2;
}
catch (RequestFailedException exception)
{
    Console.Error.WriteLine(
        $"Azure request failed ({exception.Status}, {exception.ErrorCode}): {exception.Message}");
    Environment.ExitCode = 1;
}

static async Task<StorageAccountResource> CompleteWithSdkPollingAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan pollingInterval,
    CancellationToken cancellationToken)
{
    Console.WriteLine("Using WaitForCompletionAsync; the SDK owns the polling loop.");

    Response<StorageAccountResource> completed =
        await operation.WaitForCompletionAsync(pollingInterval, cancellationToken);

    Console.WriteLine(
        $"SDK polling finished; completed={operation.HasCompleted}; " +
        $"hasValue={operation.HasValue}; HTTP status={completed.GetRawResponse().Status}.");
    return completed.Value;
}

static async Task<StorageAccountResource> CompleteWithManualPollingAsync(
    ArmOperation<StorageAccountResource> operation,
    TimeSpan pollingInterval,
    CancellationToken cancellationToken)
{
    Console.WriteLine("Using manual polling; the application controls status checks and delays.");

    while (!operation.HasCompleted)
    {
        Response statusResponse = await operation.UpdateStatusAsync(cancellationToken);
        Console.WriteLine(
            $"{DateTimeOffset.UtcNow:O} completed={operation.HasCompleted}; " +
            $"HTTP status={statusResponse.Status}; reason={statusResponse.ReasonPhrase}.");

        if (!operation.HasCompleted)
        {
            await Task.Delay(pollingInterval, cancellationToken);
        }
    }

    if (!operation.HasValue)
    {
        throw new InvalidOperationException(
            $"Operation {operation.Id} completed without returning a storage account.");
    }

    return operation.Value;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException($"Environment variable {name} is required.")
        : value;
}

static double GetPositiveDouble(string name, double defaultValue)
{
    string? text = Environment.GetEnvironmentVariable(name);
    if (string.IsNullOrWhiteSpace(text))
    {
        return defaultValue;
    }

    return double.TryParse(text, out double value) && value > 0
        ? value
        : throw new InvalidOperationException(
            $"Environment variable {name} must be a positive number.");
}
