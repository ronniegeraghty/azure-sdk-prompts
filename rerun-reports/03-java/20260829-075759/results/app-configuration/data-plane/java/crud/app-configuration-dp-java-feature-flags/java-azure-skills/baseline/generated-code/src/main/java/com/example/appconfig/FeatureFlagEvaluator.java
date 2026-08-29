package com.example.appconfig;

public final class FeatureFlagEvaluator {
    private final SyncConfigurationService configurationService;
    private final String label;

    public FeatureFlagEvaluator(SyncConfigurationService configurationService, String label) {
        this.configurationService = java.util.Objects.requireNonNull(configurationService, "configurationService");
        this.label = label;
    }

    public boolean isEnabled(String flagName) {
        return isEnabled(flagName, null);
    }

    public boolean isEnabled(String flagName, String userId) {
        String payload = configurationService
            .getSetting(FeatureFlagSupport.KEY_PREFIX + flagName, label)
            .orElse(null);
        return payload != null && FeatureFlagSupport.evaluate(flagName, payload, userId);
    }
}
