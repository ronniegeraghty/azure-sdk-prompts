package com.example;

import com.azure.core.exception.HttpResponseException;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

public final class AppConfigurationExample {
    private static final String CONNECTION_STRING_ENV = "AZURE_APP_CONFIGURATION_CONNECTION_STRING";
    private static final String FONT_SIZE_KEY = "app:Settings:FontSize";

    private AppConfigurationExample() {
    }

    public static void main(String[] args) {
        String connectionString = System.getenv(CONNECTION_STRING_ENV);
        if (connectionString == null || connectionString.isBlank()) {
            System.err.printf("Set the %s environment variable before running.%n", CONNECTION_STRING_ENV);
            return;
        }

        ConfigurationClient client = new ConfigurationClientBuilder()
            .connectionString(connectionString)
            .buildClient();

        try {
            // Create an unlabeled setting.
            client.setConfigurationSetting(FONT_SIZE_KEY, null, "24");

            // The same key can have a separate value for the Production label.
            client.setConfigurationSetting(FONT_SIZE_KEY, "Production", "24");

            ConfigurationSetting setting = client.getConfigurationSetting(FONT_SIZE_KEY, null);
            System.out.printf("%s = %s%n", setting.getKey(), setting.getValue());

            SettingSelector selector = new SettingSelector().setKeyFilter("app:Settings:*");
            client.listConfigurationSettings(selector).forEach(item ->
                System.out.printf("Found: key=%s, label=%s, value=%s%n",
                    item.getKey(), item.getLabel(), item.getValue()));

            FeatureFlagConfigurationSetting featureFlag =
                new FeatureFlagConfigurationSetting("BetaFeature", false);
            client.setConfigurationSetting(featureFlag);

            client.deleteConfigurationSetting(FONT_SIZE_KEY, null);
            System.out.printf("Deleted unlabeled setting: %s%n", FONT_SIZE_KEY);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf("Azure App Configuration request failed (HTTP %d): %s%n",
                statusCode, exception.getMessage());
        }
    }
}
