package com.example.appconfig;

import java.util.Objects;

public class FeatureFlagEvaluator {
    public static final String FEATURE_FLAG_PREFIX = ".appconfig.featureflag/";

    private final ConfigurationService configurationService;
    private final String label;

    public FeatureFlagEvaluator(ConfigurationService configurationService) {
        this(configurationService, null);
    }

    public FeatureFlagEvaluator(ConfigurationService configurationService, String label) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
        this.label = label;
    }

    public boolean isEnabled(String flagName) {
        return isEnabled(flagName, null);
    }

    public boolean isEnabled(String flagName, String userId) {
        String key = flagName.startsWith(FEATURE_FLAG_PREFIX)
            ? flagName
            : FEATURE_FLAG_PREFIX + flagName;
        return configurationService.getSetting(key, label)
            .map(payload -> FeatureFlagLogic.evaluate(flagName, payload, userId))
            .orElse(false);
    }
}
