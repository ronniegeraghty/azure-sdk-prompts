package com.example.appconfig;

import reactor.core.publisher.Mono;

import java.util.Objects;

public class AsyncFeatureFlagEvaluator {
    private final AsyncConfigurationService configurationService;
    private final String label;

    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService) {
        this(configurationService, null);
    }

    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService, String label) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
        this.label = label;
    }

    public Mono<Boolean> isEnabled(String flagName) {
        return isEnabled(flagName, null);
    }

    public Mono<Boolean> isEnabled(String flagName, String userId) {
        String key = flagName.startsWith(FeatureFlagEvaluator.FEATURE_FLAG_PREFIX)
            ? flagName
            : FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + flagName;
        return configurationService.getSetting(key, label)
            .map(payload -> payload
                .map(value -> FeatureFlagLogic.evaluate(flagName, value, userId))
                .orElse(false));
    }
}
