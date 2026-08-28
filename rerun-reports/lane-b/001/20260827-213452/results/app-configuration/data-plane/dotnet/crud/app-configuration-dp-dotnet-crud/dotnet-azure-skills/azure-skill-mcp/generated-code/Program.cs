using Azure;
using Azure.Data.AppConfiguration;

const string connectionStringEnvironmentVariable = "AZURE_APP_CONFIGURATION_CONNECTION_STRING";
const string settingsPrefix = "app:Settings:";
const string fontSizeKey = $"{settingsPrefix}FontSize";
const string productionLabel = "Production";

string? connectionString = Environment.GetEnvironmentVariable(connectionStringEnvironmentVariable);
if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        $"Set the {connectionStringEnvironmentVariable} environment variable to an Azure App Configuration connection string.");
    return 1;
}

try
{
    var client = new ConfigurationClient(connectionString);

    ConfigurationSetting fontSizeSetting = client.SetConfigurationSetting(
        new ConfigurationSetting(fontSizeKey, "24"));
    Console.WriteLine($"Set {fontSizeSetting.Key} = {fontSizeSetting.Value}");

    ConfigurationSetting productionSetting = client.SetConfigurationSetting(
        new ConfigurationSetting(fontSizeKey, "24", productionLabel));
    Console.WriteLine(
        $"Set {productionSetting.Key} = {productionSetting.Value} (label: {productionSetting.Label})");

    ConfigurationSetting retrievedSetting = client.GetConfigurationSetting(fontSizeKey);
    Console.WriteLine($"Retrieved value: {retrievedSetting.Value}");

    Console.WriteLine($"Settings with prefix \"{settingsPrefix}\":");
    var selector = new SettingSelector
    {
        KeyFilter = $"{settingsPrefix}*"
    };

    foreach (ConfigurationSetting setting in client.GetConfigurationSettings(selector))
    {
        string label = setting.Label is null ? "<no label>" : setting.Label;
        Console.WriteLine($"  {setting.Key} = {setting.Value} (label: {label})");
    }

    var betaFeature = new FeatureFlagConfigurationSetting("BetaFeature", isEnabled: true);
    ConfigurationSetting featureFlag = client.SetConfigurationSetting(betaFeature);
    Console.WriteLine($"Created enabled feature flag: {featureFlag.Key}");

    client.DeleteConfigurationSetting(fontSizeKey);
    Console.WriteLine($"Deleted unlabeled setting: {fontSizeKey}");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. Status: {ex.Status}; Error code: {ex.ErrorCode ?? "<none>"}; Message: {ex.Message}");
    return 1;
}
