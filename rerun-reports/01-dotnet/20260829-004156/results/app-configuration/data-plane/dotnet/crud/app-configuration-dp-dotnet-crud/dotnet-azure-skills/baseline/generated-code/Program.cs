using Azure;
using Azure.Data.AppConfiguration;

const string key = "app:Settings:FontSize";
const string productionLabel = "Production";
const string featureFlagName = "BetaFeature";

string? connectionString =
    Environment.GetEnvironmentVariable("AZURE_APP_CONFIGURATION_CONNECTION_STRING");

if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        "Set AZURE_APP_CONFIGURATION_CONNECTION_STRING to an Azure App Configuration connection string.");
    return 1;
}

try
{
    var client = new ConfigurationClient(connectionString);

    // Create an unlabeled setting and a separate Production-labeled setting.
    client.SetConfigurationSetting(new ConfigurationSetting(key, "24"));
    client.SetConfigurationSetting(new ConfigurationSetting(key, "24", productionLabel));

    Response<ConfigurationSetting> response = client.GetConfigurationSetting(key);
    Console.WriteLine($"{response.Value.Key} = {response.Value.Value}");

    Console.WriteLine("Settings with prefix app:Settings:");
    var selector = new SettingSelector
    {
        KeyFilter = "app:Settings:*"
    };

    foreach (ConfigurationSetting setting in client.GetConfigurationSettings(selector))
    {
        string label = setting.Label ?? "(no label)";
        Console.WriteLine($"{setting.Key} = {setting.Value} [Label: {label}]");
    }

    var featureFlag = new FeatureFlagConfigurationSetting(featureFlagName, isEnabled: true);
    client.SetConfigurationSetting(featureFlag);
    Console.WriteLine($"Feature flag {featureFlagName} created and enabled.");

    client.DeleteConfigurationSetting(key);
    Console.WriteLine($"Deleted unlabeled setting {key}.");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. Status: {ex.Status}, " +
        $"ErrorCode: {ex.ErrorCode ?? "(none)"}, Message: {ex.Message}");
    return 1;
}
