using Azure;
using Azure.Data.AppConfiguration;

const string connectionStringVariable = "AZURE_APP_CONFIGURATION_CONNECTION_STRING";
const string key = "app:Settings:FontSize";
const string value = "24";
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

    client.SetConfigurationSetting(key, value);
    Console.WriteLine($"Set {key} = {value}");

    client.SetConfigurationSetting(key, value, productionLabel);
    Console.WriteLine($"Set {key} = {value} with label {productionLabel}");

    ConfigurationSetting setting = client.GetConfigurationSetting(key).Value;
    Console.WriteLine($"Retrieved {setting.Key} = {setting.Value}");

    Console.WriteLine("Settings with prefix \"app:Settings:\":");
    var selector = new SettingSelector
    {
        KeyFilter = "app:Settings:*"
    };

    foreach (ConfigurationSetting matchingSetting in client.GetConfigurationSettings(selector))
    {
        string label = matchingSetting.Label is null ? "(no label)" : matchingSetting.Label;
        Console.WriteLine($"  {matchingSetting.Key} = {matchingSetting.Value}, label = {label}");
    }

    var featureFlag = new FeatureFlagConfigurationSetting("BetaFeature", isEnabled: true);
    client.SetConfigurationSetting(featureFlag);
    Console.WriteLine("Enabled feature flag BetaFeature");

    client.DeleteConfigurationSetting(key);
    Console.WriteLine($"Deleted the unlabeled setting {key}");

    return 0;
}
catch (RequestFailedException ex)
{
    Console.Error.WriteLine(
        $"Azure App Configuration request failed. Status: {ex.Status}, ErrorCode: {ex.ErrorCode ?? "unknown"}");
    Console.Error.WriteLine(ex.Message);
    return 2;
}
