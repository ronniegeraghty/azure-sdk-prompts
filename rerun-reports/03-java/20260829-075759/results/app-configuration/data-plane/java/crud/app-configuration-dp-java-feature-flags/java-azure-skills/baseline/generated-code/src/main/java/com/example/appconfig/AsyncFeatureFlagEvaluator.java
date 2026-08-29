package com.example.appconfig;

import reactor.core.publisher.Mono;

public final class AsyncFeatureFlagEvaluator {
    private final AsyncConfigurationService configurationService;
    private final String label;

    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService, String label) {
        this.configurationService = java.util.Objects.requireNonNull(configurationService, "configurationService");
        this.label = label;
    }

    public Mono<Boolean> isEnabled(String flagName) {
        return isEnabled(flagName, null);
    }

    public Mono<Boolean> isEnabled(String flagName, String userId) {
        return configurationService
            .getSetting(FeatureFlagSupport.KEY_PREFIX + flagName, label)
            .map(payload -> payload
                .map(json -> FeatureFlagSupport.evaluate(flagName, json, userId))
                .orElse(false));
    }
}
