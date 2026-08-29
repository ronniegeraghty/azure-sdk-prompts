package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;
import reactor.util.retry.Retry;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;
import java.util.concurrent.TimeoutException;

public final class AsyncSecretRotator {
    private final SecretAsyncClient client;
    private final Duration pollInterval;
    private final Duration timeout;

    public AsyncSecretRotator(SecretAsyncClient client, Duration pollInterval, Duration timeout) {
        this.client = Objects.requireNonNull(client, "client");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
        this.timeout = requirePositive(timeout, "timeout");
    }

    public Mono<KeyVaultSecret> rotate(
            String name,
            String newValue,
            OffsetDateTime expiresOn) {
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        return client.beginDeleteSecret(name)
                .last()
                .then(purgeWhenVisible(name))
                .then(waitUntilPurged(name))
                .then(client.setSecret(new KeyVaultSecret(name, newValue)
                        .setProperties(new SecretProperties().setExpiresOn(expiresOn))));
    }

    private Mono<Void> purgeWhenVisible(String name) {
        return client.purgeDeletedSecret(name)
                .retryWhen(Retry.fixedDelay(maxRetries(), pollInterval)
                        .filter(ResourceNotFoundException.class::isInstance)
                        .onRetryExhaustedThrow((spec, signal) -> new IllegalStateException(
                                "Deleted secret did not become purgeable before timeout: " + name,
                                signal.failure())));
    }

    private Mono<Void> waitUntilPurged(String name) {
        return Flux.interval(Duration.ZERO, pollInterval)
                .concatMap(ignored -> client.getDeletedSecret(name)
                        .map(secret -> false)
                        .onErrorResume(ResourceNotFoundException.class, exception -> Mono.just(true)))
                .filter(Boolean::booleanValue)
                .next()
                .timeout(timeout)
                .onErrorMap(TimeoutException.class, exception -> new IllegalStateException(
                        "Secret purge did not complete before timeout: " + name, exception))
                .then();
    }

    private long maxRetries() {
        return Math.max(1L, timeout.dividedBy(pollInterval));
    }

    private static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }
}
