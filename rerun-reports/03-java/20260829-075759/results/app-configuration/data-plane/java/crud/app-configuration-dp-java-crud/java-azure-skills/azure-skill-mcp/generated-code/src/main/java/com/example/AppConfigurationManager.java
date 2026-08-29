package com.example;

import com.azure.core.exception.HttpResponseException;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

public final class AppConfigurationManager {
    private static final String CONNECTION_STRING_ENVIRONMENT_VARIABLE =
        "AZURE_APPCONFIGURATION_CONNECTION_STRING";
    private static final String FONT_SIZE_KEY = "app:Settings:FontSize";
    private static final String PRODUCTION_LABEL = "Production";

    private AppConfigurationManager() {
    }

    public static void main(String[] args) {
        String connectionString = System.getenv(CONNECTION_STRING_ENVIRONMENT_VARIABLE);
        if (connectionString == null || connectionString.isBlank()) {
            System.err.printf(
                "Set the %s environment variable before running the program.%n",
                CONNECTION_STRING_ENVIRONMENT_VARIABLE);
            System.exit(1);
        }

        ConfigurationClient client = new ConfigurationClientBuilder()
            .connectionString(connectionString)
            .buildClient();

        try {
            client.setConfigurationSetting(FONT_SIZE_KEY, null, "24");

            client.setConfigurationSetting(
                new ConfigurationSetting()
                    .setKey(FONT_SIZE_KEY)
                    .setLabel(PRODUCTION_LABEL)
                    .setValue("24"));

            ConfigurationSetting setting =
                client.getConfigurationSetting(FONT_SIZE_KEY, null);
            System.out.printf("%s = %s%n", setting.getKey(), setting.getValue());

            System.out.println("Matching settings:");
            client.listConfigurationSettings(
                    new SettingSelector().setKeyFilter("app:Settings:*"))
                .forEach(item -> System.out.printf(
                    "  key=%s, label=%s, value=%s%n",
                    item.getKey(),
                    item.getLabel(),
                    item.getValue()));

            FeatureFlagConfigurationSetting betaFeature =
                new FeatureFlagConfigurationSetting("BetaFeature", true);
            client.setConfigurationSetting(betaFeature);

            client.deleteConfigurationSetting(FONT_SIZE_KEY, null);
            System.out.printf("Deleted the unlabeled setting %s.%n", FONT_SIZE_KEY);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf(
                "Azure App Configuration request failed (HTTP %d): %s%n",
                statusCode,
                exception.getMessage());
            System.exit(1);
        }
    }
}
