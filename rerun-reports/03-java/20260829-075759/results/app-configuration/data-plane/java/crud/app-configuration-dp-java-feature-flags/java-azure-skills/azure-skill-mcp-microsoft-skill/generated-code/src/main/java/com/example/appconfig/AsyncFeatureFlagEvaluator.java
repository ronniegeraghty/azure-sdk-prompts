package com.example.appconfig;

import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncFeatureFlagEvaluator {
    private final AsyncSettingProvider settings;

    public AsyncFeatureFlagEvaluator(AsyncSettingProvider settings) {
        this.settings = Objects.requireNonNull(settings, "settings");
    }

    public Mono<Boolean> isEnabled(String featureId) {
        return isEnabled(featureId, null);
    }

    public Mono<Boolean> isEnabled(String featureId, String label) {
        return evaluate(featureId, label, null);
    }

    public Mono<Boolean> isEnabledForUser(String featureId, String userId) {
        return isEnabledForUser(featureId, null, userId);
    }

    public Mono<Boolean> isEnabledForUser(String featureId, String label, String userId) {
        Objects.requireNonNull(userId, "userId");
        return evaluate(featureId, label, userId);
    }

    private Mono<Boolean> evaluate(String featureId, String label, String userId) {
        Objects.requireNonNull(featureId, "featureId");
        return settings.getSetting(FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + featureId, label)
            .map(payload -> FeatureFlagRules.evaluate(payload, featureId, userId))
            .defaultIfEmpty(false);
    }
}
