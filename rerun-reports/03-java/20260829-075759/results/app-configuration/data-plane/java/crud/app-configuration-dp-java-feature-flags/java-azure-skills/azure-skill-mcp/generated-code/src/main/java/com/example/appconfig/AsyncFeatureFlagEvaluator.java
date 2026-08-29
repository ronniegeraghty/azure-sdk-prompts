package com.example.appconfig;

import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncFeatureFlagEvaluator {
    private final AsyncConfigurationService configurationService;

    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
    }

    public Mono<Boolean> isEnabled(String flagId, String userId) {
        return isEnabled(flagId, null, userId);
    }

    public Mono<Boolean> isEnabled(String flagId, String label, String userId) {
        String id = requireFlagId(flagId);
        return configurationService.getSetting(FeatureFlag.KEY_PREFIX + id, label)
            .map(value -> value.map(json -> FeatureFlag.evaluate(id, json, userId)).orElse(false));
    }

    private static String requireFlagId(String flagId) {
        if (flagId == null || flagId.isBlank()) {
            throw new IllegalArgumentException("Feature flag ID must not be blank");
        }
        return flagId;
    }
}
