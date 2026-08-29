package com.example;

import com.azure.core.exception.HttpResponseException;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

public final class AppConfigurationExample {
    private static final String CONNECTION_STRING_ENV = "AZURE_APPCONFIG_CONNECTION_STRING";
    private static final String SETTING_KEY = "app:Settings:FontSize";
    private static final String SETTING_VALUE = "24";

    private AppConfigurationExample() {
    }

    public static void main(String[] args) {
        String connectionString = System.getenv(CONNECTION_STRING_ENV);
        if (connectionString == null || connectionString.isBlank()) {
            System.err.printf("Set the %s environment variable before running the program.%n",
                CONNECTION_STRING_ENV);
            System.exit(1);
        }

        ConfigurationClient client = new ConfigurationClientBuilder()
            .connectionString(connectionString)
            .buildClient();

        try {
            client.setConfigurationSetting(SETTING_KEY, null, SETTING_VALUE);
            client.setConfigurationSetting(SETTING_KEY, "Production", SETTING_VALUE);

            ConfigurationSetting setting = client.getConfigurationSetting(SETTING_KEY, null);
            System.out.printf("%s = %s%n", setting.getKey(), setting.getValue());

            SettingSelector selector = new SettingSelector().setKeyFilter("app:Settings:*");
            for (ConfigurationSetting matchingSetting : client.listConfigurationSettings(selector)) {
                System.out.printf("Key: %s, Label: %s, Value: %s%n",
                    matchingSetting.getKey(),
                    matchingSetting.getLabel(),
                    matchingSetting.getValue());
            }

            FeatureFlagConfigurationSetting betaFeature =
                new FeatureFlagConfigurationSetting("BetaFeature", true);
            client.setConfigurationSetting(betaFeature);

            client.deleteConfigurationSetting(SETTING_KEY, null);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf("Azure App Configuration request failed (HTTP %d): %s%n",
                statusCode, exception.getMessage());
            System.exit(1);
        }
    }
}
