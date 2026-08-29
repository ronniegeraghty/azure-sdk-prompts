package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.time.Instant;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class AsyncSecretRotationHelper {
    private final SecretAsyncClient client;
    private final Duration purgeTimeout;
    private final Duration pollInterval;

    public AsyncSecretRotationHelper(
            SecretAsyncClient client,
            Duration purgeTimeout,
            Duration pollInterval) {
        this.client = Objects.requireNonNull(client, "client");
        this.purgeTimeout = requirePositive(purgeTimeout, "purgeTimeout");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
    }

    public Mono<KeyVaultSecret> rotate(String name, String newValue, OffsetDateTime expiresOn) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        return client.beginDeleteSecret(name)
                .last()
                .then(client.purgeDeletedSecret(name))
                .then(waitUntilPurged(name, Instant.now().plus(purgeTimeout)))
                .then(Mono.defer(() -> {
                    KeyVaultSecret replacement = new KeyVaultSecret(name, newValue);
                    replacement.getProperties().setExpiresOn(expiresOn);
                    return client.setSecret(replacement);
                }));
    }

    private Mono<Void> waitUntilPurged(String name, Instant deadline) {
        return Mono.defer(() -> client.getDeletedSecret(name)
                .flatMap(deletedSecret -> {
                    if (!Instant.now().isBefore(deadline)) {
                        return Mono.<Void>error(new IllegalStateException(
                                "Timed out waiting for deleted secret to be purged: " + name));
                    }
                    return Mono.delay(pollInterval)
                            .then(waitUntilPurged(name, deadline));
                })
                .onErrorResume(ResourceNotFoundException.class, exception -> Mono.empty()));
    }

    private static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }
}
