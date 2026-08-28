package com.example.keyvaultconfig;

import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;
import reactor.util.retry.Retry;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class AsyncSecretRotationHelper {
    private final SecretAsyncClient client;
    private final Duration timeout;
    private final Duration pollInterval;

    public AsyncSecretRotationHelper(
            SecretAsyncClient client,
            Duration timeout,
            Duration pollInterval) {
        this.client = Objects.requireNonNull(client, "client");
        this.timeout = SecretRotationHelper.requirePositive(timeout, "timeout");
        this.pollInterval = SecretRotationHelper.requirePositive(pollInterval, "pollInterval");
    }

    public Mono<KeyVaultSecret> rotate(
            String name,
            String newValue,
            OffsetDateTime newExpiresOn) {
        SecretRotationHelper.requireRotationArguments(name, newValue, newExpiresOn);

        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue);
        replacement.getProperties().setExpiresOn(newExpiresOn);

        return client.beginDeleteSecret(name)
                .last()
                .then(client.purgeDeletedSecret(name))
                .then(waitUntilPurged(name))
                .then(createAfterPurge(replacement))
                .timeout(timeout);
    }

    private Mono<Void> waitUntilPurged(String name) {
        return Flux.interval(Duration.ZERO, pollInterval)
                .concatMap(ignored -> client.getDeletedSecret(name)
                        .map(secret -> false)
                        .onErrorResume(
                                ResourceNotFoundException.class,
                                exception -> Mono.just(true)))
                .filter(Boolean::booleanValue)
                .next()
                .then();
    }

    private Mono<KeyVaultSecret> createAfterPurge(KeyVaultSecret replacement) {
        long retries = Math.max(1, timeout.dividedBy(pollInterval));
        return client.setSecret(replacement)
                .retryWhen(Retry.fixedDelay(retries, pollInterval)
                        .filter(this::isNameConflict));
    }

    private boolean isNameConflict(Throwable throwable) {
        return throwable instanceof HttpResponseException exception
                && exception.getResponse() != null
                && exception.getResponse().getStatusCode() == 409;
    }
}
