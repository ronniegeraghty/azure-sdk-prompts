using Azure;
using Azure.Data.AppConfiguration;

const string connectionStringVariable = "AZURE_APPCONFIG_CONNECTION_STRING";
const string settingKey = "app:Settings:FontSize";
const string settingPrefix = "app:Settings:";
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
    var client = new ConfigurationClient(connectionString);

    client.SetConfigurationSetting(settingKey, "24");
    Console.WriteLine($"Set '{settingKey}' to '24'.");

    var productionSetting = new ConfigurationSetting(settingKey, "24", productionLabel);
    client.SetConfigurationSetting(productionSetting);
    Console.WriteLine($"Set '{settingKey}' with label '{productionLabel}'.");

    ConfigurationSetting setting = client.GetConfigurationSetting(settingKey).Value;
    Console.WriteLine($"Value for '{settingKey}': {setting.Value}");

    Console.WriteLine($"Settings with prefix '{settingPrefix}':");
    var selector = new SettingSelector
    {
        KeyFilter = $"{settingPrefix}*"
    };

    foreach (ConfigurationSetting matchingSetting in client.GetConfigurationSettings(selector))
    {
        string label = matchingSetting.Label ?? "(no label)";
        Console.WriteLine($"  {matchingSetting.Key} = {matchingSetting.Value} [{label}]");
    }

    var featureFlag = new FeatureFlagConfigurationSetting("BetaFeature", isEnabled: true);
    client.SetConfigurationSetting(featureFlag);
    Console.WriteLine("Enabled feature flag 'BetaFeature'.");

    client.DeleteConfigurationSetting(settingKey);
    Console.WriteLine($"Deleted the unlabeled setting '{settingKey}'.");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. Status: {ex.Status}; " +
        $"Error code: {ex.ErrorCode ?? "(none)"}; Message: {ex.Message}");
    return 1;
}
catch (ArgumentException ex)
{
    Console.Error.WriteLine($"The connection string is invalid: {ex.Message}");
    return 1;
}
