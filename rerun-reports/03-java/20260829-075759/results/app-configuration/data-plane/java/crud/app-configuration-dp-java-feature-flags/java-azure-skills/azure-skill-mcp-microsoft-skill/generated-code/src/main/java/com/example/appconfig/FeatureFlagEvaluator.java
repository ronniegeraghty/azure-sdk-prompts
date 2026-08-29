package com.example.appconfig;

import java.util.Objects;

public final class FeatureFlagEvaluator {
    public static final String FEATURE_FLAG_PREFIX = ".appconfig.featureflag/";

    private final SettingProvider settings;

    public FeatureFlagEvaluator(SettingProvider settings) {
        this.settings = Objects.requireNonNull(settings, "settings");
    }

    public boolean isEnabled(String featureId) {
        return isEnabled(featureId, null);
    }

    public boolean isEnabled(String featureId, String label) {
        return evaluate(featureId, label, null);
    }

    public boolean isEnabledForUser(String featureId, String userId) {
        return isEnabledForUser(featureId, null, userId);
    }

    public boolean isEnabledForUser(String featureId, String label, String userId) {
        Objects.requireNonNull(userId, "userId");
        return evaluate(featureId, label, userId);
    }

    private boolean evaluate(String featureId, String label, String userId) {
        Objects.requireNonNull(featureId, "featureId");
        return settings.getSetting(FEATURE_FLAG_PREFIX + featureId, label)
            .map(payload -> FeatureFlagRules.evaluate(payload, featureId, userId))
            .orElse(false);
    }
}
