package com.example;

import com.azure.core.exception.HttpResponseException;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

public final class AppConfigurationDemo {
    private static final String CONNECTION_STRING_ENV = "AZURE_APPCONFIG_CONNECTION_STRING";
    private static final String FONT_SIZE_KEY = "app:Settings:FontSize";
    private static final String PRODUCTION_LABEL = "Production";

    private AppConfigurationDemo() {
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
            client.setConfigurationSetting(FONT_SIZE_KEY, null, "24");
            client.setConfigurationSetting(FONT_SIZE_KEY, PRODUCTION_LABEL, "24");

            ConfigurationSetting setting = client.getConfigurationSetting(FONT_SIZE_KEY, null);
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
            betaFeature.setDescription("Enables the beta feature");
            client.setConfigurationSetting(betaFeature);

            client.deleteConfigurationSetting(FONT_SIZE_KEY, null);
            System.out.printf("Deleted setting: %s%n", FONT_SIZE_KEY);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();

            switch (statusCode) {
                case 401:
                case 403:
                    System.err.println("Authentication or authorization failed. Check the connection string permissions.");
                    break;
                case 404:
                    System.err.println("The requested configuration setting was not found.");
                    break;
                case 429:
                    System.err.println("Azure App Configuration throttled the request. Retry after a delay.");
                    break;
                default:
                    System.err.printf(
                        "Azure App Configuration request failed (HTTP %s): %s%n",
                        statusCode == -1 ? "unknown" : Integer.toString(statusCode),
                        exception.getMessage());
            }
            System.exit(1);
        }
    }
}
