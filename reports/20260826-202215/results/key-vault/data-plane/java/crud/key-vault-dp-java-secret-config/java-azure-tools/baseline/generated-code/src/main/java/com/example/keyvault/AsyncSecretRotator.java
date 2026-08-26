package com.example.keyvault;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;
import reactor.core.publisher.Mono;
import reactor.util.retry.Retry;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class AsyncSecretRotator {
    private final SecretAsyncClient client;
    private final Duration pollInterval;
    private final Duration timeout;

    public AsyncSecretRotator(SecretAsyncClient client, Duration pollInterval, Duration timeout) {
        this.client = Objects.requireNonNull(client, "client");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
        this.timeout = requirePositive(timeout, "timeout");
    }

    public Mono<KeyVaultSecret> rotate(String name, String newValue, OffsetDateTime expiresOn) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue)
                .setProperties(new SecretProperties().setExpiresOn(expiresOn));

        Mono<Void> deleteActive = client.beginDeleteSecret(name)
                .last()
                .then()
                .onErrorResume(ResourceNotFoundException.class, ignored -> Mono.empty());
        Mono<Void> purgeDeleted = client.purgeDeletedSecret(name)
                .then(waitUntilPurged(name))
                .onErrorResume(ResourceNotFoundException.class, ignored -> Mono.empty());

        return deleteActive
                .then(purgeDeleted)
                .then(client.setSecret(replacement));
    }

    private Mono<Void> waitUntilPurged(String name) {
        return Mono.defer(() -> client.getDeletedSecret(name)
                        .flatMap(ignored -> Mono.<Void>error(new StillExistsException())))
                .retryWhen(Retry.fixedDelay(Long.MAX_VALUE, pollInterval)
                        .filter(StillExistsException.class::isInstance))
                .onErrorResume(ResourceNotFoundException.class, ignored -> Mono.empty())
                .timeout(timeout);
    }

    private static Duration requirePositive(Duration value, String name) {
        Objects.requireNonNull(value, name);
        if (value.isZero() || value.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return value;
    }

    private static final class StillExistsException extends RuntimeException {
    }
}
