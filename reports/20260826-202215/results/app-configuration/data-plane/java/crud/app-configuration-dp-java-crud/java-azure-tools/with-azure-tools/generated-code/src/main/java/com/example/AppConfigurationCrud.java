package com.example;

import com.azure.core.exception.HttpResponseException;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

public final class AppConfigurationCrud {
    private static final String CONNECTION_STRING_ENV = "AZURE_APPCONFIG_CONNECTION_STRING";
    private static final String SETTING_KEY = "app:Settings:FontSize";
    private static final String PRODUCTION_LABEL = "Production";

    private AppConfigurationCrud() {
    }

    public static void main(String[] args) {
        String connectionString = System.getenv(CONNECTION_STRING_ENV);
        if (connectionString == null || connectionString.isBlank()) {
            System.err.printf("Set the %s environment variable before running.%n", CONNECTION_STRING_ENV);
            System.exit(1);
        }

        ConfigurationClient client = new ConfigurationClientBuilder()
            .connectionString(connectionString)
            .buildClient();

        try {
            client.setConfigurationSetting(SETTING_KEY, null, "24");
            client.setConfigurationSetting(SETTING_KEY, PRODUCTION_LABEL, "24");

            ConfigurationSetting setting = client.getConfigurationSetting(SETTING_KEY, null);
            System.out.printf("%s = %s%n", setting.getKey(), setting.getValue());

            SettingSelector selector = new SettingSelector()
                .setKeyFilter("app:Settings:*");
            client.listConfigurationSettings(selector)
                .forEach(item -> System.out.printf(
                    "%s [%s] = %s%n",
                    item.getKey(),
                    item.getLabel() == null ? "no label" : item.getLabel(),
                    item.getValue()));

            FeatureFlagConfigurationSetting betaFeature =
                new FeatureFlagConfigurationSetting("BetaFeature", true);
            client.setConfigurationSetting(betaFeature);

            client.deleteConfigurationSetting(SETTING_KEY, null);
            System.out.printf("Deleted setting: %s%n", SETTING_KEY);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf(
                "Azure App Configuration request failed (HTTP %s): %s%n",
                statusCode == -1 ? "unknown" : Integer.toString(statusCode),
                exception.getMessage());
            System.exit(1);
        }
    }
}
