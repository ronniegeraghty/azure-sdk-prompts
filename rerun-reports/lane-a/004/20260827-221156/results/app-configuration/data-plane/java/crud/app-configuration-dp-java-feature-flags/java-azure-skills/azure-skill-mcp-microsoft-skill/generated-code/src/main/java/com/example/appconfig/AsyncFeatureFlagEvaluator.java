package com.example.appconfig;

import reactor.core.publisher.Mono;

public final class AsyncFeatureFlagEvaluator {
    private final AsyncConfigurationReader configuration;

    public AsyncFeatureFlagEvaluator(AsyncConfigurationReader configuration) {
        this.configuration = configuration;
    }

    public Mono<Boolean> isEnabled(String flagId, String userId) {
        return isEnabled(flagId, userId, null);
    }

    public Mono<Boolean> isEnabled(String flagId, String userId, String label) {
        return configuration.getSetting(FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + flagId, label)
            .map(payload -> FeatureFlagEvaluator.evaluate(flagId, userId, payload));
    }
}
