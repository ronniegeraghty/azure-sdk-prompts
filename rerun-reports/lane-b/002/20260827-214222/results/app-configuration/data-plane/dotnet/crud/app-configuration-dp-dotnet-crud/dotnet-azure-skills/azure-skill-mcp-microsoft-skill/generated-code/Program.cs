using Azure;
using Azure.Data.AppConfiguration;

const string connectionStringVariable = "AZURE_APPCONFIG_CONNECTION_STRING";
const string key = "app:Settings:FontSize";
const string value = "24";
const string productionLabel = "Production";
const string keyPrefix = "app:Settings:";

string? connectionString = Environment.GetEnvironmentVariable(connectionStringVariable);

if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        $"Set the {connectionStringVariable} environment variable to an Azure App Configuration connection string.");
    return 1;
}

var client = new ConfigurationClient(connectionString);
var featureFlag = new FeatureFlagConfigurationSetting("BetaFeature", isEnabled: true);

try
{
    client.SetConfigurationSetting(new ConfigurationSetting(key, value));
    Console.WriteLine($"Set '{key}' without a label.");

    client.SetConfigurationSetting(
        new ConfigurationSetting(key, value, productionLabel));
    Console.WriteLine($"Set '{key}' with label '{productionLabel}'.");

    ConfigurationSetting setting = client.GetConfigurationSetting(key).Value;
    Console.WriteLine($"Value for '{setting.Key}': {setting.Value}");

    var selector = new SettingSelector
    {
        KeyFilter = $"{keyPrefix}*"
    };

    Console.WriteLine($"Settings with prefix '{keyPrefix}':");
    foreach (ConfigurationSetting matchingSetting in client.GetConfigurationSettings(selector))
    {
        string label = matchingSetting.Label ?? "(no label)";
        Console.WriteLine(
            $"  Key: {matchingSetting.Key}, Value: {matchingSetting.Value}, Label: {label}");
    }

    client.SetConfigurationSetting(featureFlag);
    Console.WriteLine(
        $"Created enabled feature flag '{featureFlag.FeatureId}'.");

    client.DeleteConfigurationSetting(key);
    Console.WriteLine($"Deleted '{key}' without a label.");

    client.DeleteConfigurationSetting(key, productionLabel);
    client.DeleteConfigurationSetting(featureFlag);
    Console.WriteLine("Deleted the labeled setting and feature flag.");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. Status: {ex.Status}, " +
        $"ErrorCode: {ex.ErrorCode ?? "(none)"}, Message: {ex.Message}");
    return 1;
}
