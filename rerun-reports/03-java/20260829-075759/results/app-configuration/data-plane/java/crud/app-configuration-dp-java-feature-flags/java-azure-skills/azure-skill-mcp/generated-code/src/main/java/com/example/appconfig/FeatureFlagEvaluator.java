package com.example.appconfig;

import java.util.Objects;

public final class FeatureFlagEvaluator {
    private final ConfigurationService configurationService;

    public FeatureFlagEvaluator(ConfigurationService configurationService) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
    }

    public boolean isEnabled(String flagId, String userId) {
        return isEnabled(flagId, null, userId);
    }

    public boolean isEnabled(String flagId, String label, String userId) {
        String id = requireFlagId(flagId);
        return configurationService.getSetting(FeatureFlag.KEY_PREFIX + id, label)
            .map(json -> FeatureFlag.evaluate(id, json, userId))
            .orElse(false);
    }

    private static String requireFlagId(String flagId) {
        if (flagId == null || flagId.isBlank()) {
            throw new IllegalArgumentException("Feature flag ID must not be blank");
        }
        return flagId;
    }
}
