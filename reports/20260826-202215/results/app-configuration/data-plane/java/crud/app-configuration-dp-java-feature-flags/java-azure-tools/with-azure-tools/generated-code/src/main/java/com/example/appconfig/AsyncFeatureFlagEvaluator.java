package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import reactor.core.publisher.Mono;

public final class AsyncFeatureFlagEvaluator {
    private final AsyncConfigurationService configurationService;

    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService) {
        this.configurationService = configurationService;
    }

    public Mono<Boolean> isEnabled(String flagId) {
        return isEnabled(flagId, null, null);
    }

    public Mono<Boolean> isEnabled(String flagId, String userId) {
        return isEnabled(flagId, userId, null);
    }

    public Mono<Boolean> isEnabled(String flagId, String userId, String label) {
        return configurationService
            .getSetting(FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + flagId, label)
            .map(payload -> FeatureFlagEvaluator.evaluatePayload(flagId, userId, payload))
            .onErrorReturn(ResourceNotFoundException.class, false);
    }
}
