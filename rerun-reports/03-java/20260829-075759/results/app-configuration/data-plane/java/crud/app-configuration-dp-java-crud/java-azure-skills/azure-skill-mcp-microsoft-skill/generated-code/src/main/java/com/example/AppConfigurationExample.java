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
    private static final String PRODUCTION_LABEL = "Production";

    private AppConfigurationExample() {
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

            System.out.println("Matching settings:");
            client.listConfigurationSettings(selector).forEach(item ->
                System.out.printf("  %s [%s] = %s%n",
                    item.getKey(),
                    item.getLabel() == null ? "(no label)" : item.getLabel(),
                    item.getValue()));

            FeatureFlagConfigurationSetting featureFlag =
                new FeatureFlagConfigurationSetting("BetaFeature", true)
                    .setDescription("Enables the beta feature");
            client.setConfigurationSetting(featureFlag);

            client.deleteConfigurationSetting(SETTING_KEY, null);
            System.out.println("Deleted the unlabeled setting: " + SETTING_KEY);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();

            if (statusCode == 404) {
                System.err.println("The requested configuration setting was not found.");
            } else if (statusCode == 401 || statusCode == 403) {
                System.err.println("Authentication or authorization failed. Check the connection string.");
            } else if (statusCode == 429) {
                System.err.println("Azure App Configuration throttled the request. Retry later.");
            } else {
                System.err.printf("Azure App Configuration request failed (HTTP %d): %s%n",
                    statusCode, exception.getMessage());
            }
            System.exit(1);
        }
    }
}
