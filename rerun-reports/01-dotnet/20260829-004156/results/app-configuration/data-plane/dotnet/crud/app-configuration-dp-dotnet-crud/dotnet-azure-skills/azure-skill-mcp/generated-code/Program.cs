using Azure;
using Azure.Data.AppConfiguration;

const string connectionStringEnvironmentVariable = "AZURE_APPCONFIG_CONNECTION_STRING";
const string settingKey = "app:Settings:FontSize";
const string productionLabel = "Production";

string? connectionString =
    Environment.GetEnvironmentVariable(connectionStringEnvironmentVariable);

if (string.IsNullOrWhiteSpace(connectionString))
{
    Console.Error.WriteLine(
        $"Set the {connectionStringEnvironmentVariable} environment variable to an " +
        "Azure App Configuration connection string.");
    return 1;
}

var client = new ConfigurationClient(connectionString);
var featureFlag = new FeatureFlagConfigurationSetting("BetaFeature", isEnabled: true);

try
{
    // Create or replace the unlabeled setting.
    client.SetConfigurationSetting(settingKey, "24");

    // The same key can have a separate value for each label.
    client.SetConfigurationSetting(
        new ConfigurationSetting(settingKey, "24", productionLabel));

    ConfigurationSetting setting = client.GetConfigurationSetting(settingKey);
    Console.WriteLine($"{setting.Key} = {setting.Value}");

    Console.WriteLine("Settings with prefix app:Settings:");
    var selector = new SettingSelector
    {
        KeyFilter = "app:Settings:*"
    };

    foreach (ConfigurationSetting matchingSetting
             in client.GetConfigurationSettings(selector))
    {
        string label = matchingSetting.Label ?? "<no label>";
        Console.WriteLine(
            $"  {matchingSetting.Key} = {matchingSetting.Value} (label: {label})");
    }

    client.SetConfigurationSetting(featureFlag);
    Console.WriteLine(
        $"Feature flag {featureFlag.FeatureId} created and enabled.");

    client.DeleteConfigurationSetting(settingKey);
    Console.WriteLine($"Deleted unlabeled setting {settingKey}.");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. " +
        $"Status: {ex.Status}; ErrorCode: {ex.ErrorCode ?? "<none>"}; " +
        $"Message: {ex.Message}");
    return 1;
}
finally
{
    // Remove the additional labeled setting and feature flag created by this sample.
    TryDelete(client, settingKey, productionLabel);
    TryDelete(client, featureFlag.Key, featureFlag.Label);
}

static void TryDelete(
    ConfigurationClient client,
    string key,
    string? label = null)
{
    try
    {
        client.DeleteConfigurationSetting(key, label);
    }
    catch (RequestFailedException ex) when (ex.Status == 404)
    {
        // The setting was not created or was already removed.
    }
    catch (RequestFailedException ex)
    {
        Console.Error.WriteLine(
            $"Cleanup failed for key '{key}' and label '{label ?? "<no label>"}'. " +
            $"Status: {ex.Status}; ErrorCode: {ex.ErrorCode ?? "<none>"}; " +
            $"Message: {ex.Message}");
    }
}
