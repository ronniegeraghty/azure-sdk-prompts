using Azure;
using Azure.Data.AppConfiguration;

const string connectionStringVariable = "AZURE_APPCONFIG_CONNECTION_STRING";
const string settingKey = "app:Settings:FontSize";
const string settingValue = "24";
const string productionLabel = "Production";

string? connectionString = Environment.GetEnvironmentVariable(connectionStringVariable);
if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        $"Set the {connectionStringVariable} environment variable to an Azure App Configuration connection string.");
    return 1;
}

try
{
    // Connection strings are required by this sample; do not store them in source control.
    var client = new ConfigurationClient(connectionString);

    client.SetConfigurationSetting(settingKey, settingValue);
    Console.WriteLine($"Set {settingKey} with the default label.");

    client.SetConfigurationSetting(settingKey, settingValue, productionLabel);
    Console.WriteLine($"Set {settingKey} with label {productionLabel}.");

    ConfigurationSetting setting = client.GetConfigurationSetting(settingKey);
    Console.WriteLine($"{setting.Key} = {setting.Value}");

    Console.WriteLine($"Settings with prefix \"app:Settings:\":");
    var selector = new SettingSelector
    {
        KeyFilter = "app:Settings:*"
    };

    foreach (ConfigurationSetting matchingSetting in client.GetConfigurationSettings(selector))
    {
        string label = matchingSetting.Label ?? "(no label)";
        Console.WriteLine($"  {matchingSetting.Key} = {matchingSetting.Value} [{label}]");
    }

    var betaFeature = new FeatureFlagConfigurationSetting("BetaFeature", isEnabled: true);
    client.SetConfigurationSetting(betaFeature);
    Console.WriteLine("Created enabled feature flag BetaFeature.");

    client.DeleteConfigurationSetting(settingKey);
    Console.WriteLine($"Deleted {settingKey} with the default label.");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. Status: {ex.Status}; " +
        $"ErrorCode: {ex.ErrorCode ?? "(none)"}; Message: {ex.Message}");
    return 2;
}
catch (ArgumentException ex)
{
    Console.Error.WriteLine($"The App Configuration connection string is invalid: {ex.Message}");
    return 3;
}
