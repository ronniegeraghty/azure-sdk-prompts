package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.time.Instant;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class SecretRotationHelper {
    private final SecretClient syncClient;
    private final SecretAsyncClient asyncClient;
    private final Duration purgeTimeout;
    private final Duration pollInterval;

    public SecretRotationHelper(
            SecretClient syncClient,
            SecretAsyncClient asyncClient,
            Duration purgeTimeout,
            Duration pollInterval) {
        this.syncClient = Objects.requireNonNull(syncClient, "syncClient");
        this.asyncClient = Objects.requireNonNull(asyncClient, "asyncClient");
        this.purgeTimeout = requirePositive(purgeTimeout, "purgeTimeout");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
    }

    public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {
        validateRotation(name, newValue, expiresOn);
        boolean deletedSecretExists = deleteIfPresent(name) || deletedSecretExists(name);
        if (deletedSecretExists) {
            Instant deadline = Instant.now().plus(purgeTimeout);
            purgeWithRetry(name, deadline);
            waitUntilPurged(name, deadline);
        }
        return syncClient.setSecret(secretWithExpiry(name, newValue, expiresOn));
    }

    public Mono<KeyVaultSecret> rotateAsync(
            String name,
            String newValue,
            OffsetDateTime expiresOn) {
        validateRotation(name, newValue, expiresOn);
        return deleteIfPresentAsync(name)
                .flatMap(deleted -> deleted
                        ? Mono.just(true)
                        : deletedSecretExistsAsync(name))
                .flatMap(deletedSecretExists -> {
                    if (!deletedSecretExists) {
                        return Mono.empty();
                    }
                    Instant deadline = Instant.now().plus(purgeTimeout);
                    return purgeWithRetryAsync(name, deadline)
                            .then(waitUntilPurgedAsync(name, deadline));
                })
                .then(asyncClient.setSecret(secretWithExpiry(name, newValue, expiresOn)));
    }

    private boolean deleteIfPresent(String name) {
        try {
            SyncPoller<DeletedSecret, Void> poller = syncClient.beginDeleteSecret(name);
            poller.waitForCompletion();
            return true;
        } catch (ResourceNotFoundException exception) {
            return false;
        }
    }

    private Mono<Boolean> deleteIfPresentAsync(String name) {
        return Mono.defer(() -> asyncClient.beginDeleteSecret(name).last().thenReturn(true))
                .onErrorResume(ResourceNotFoundException.class, exception -> Mono.just(false));
    }

    private boolean deletedSecretExists(String name) {
        try {
            syncClient.getDeletedSecret(name);
            return true;
        } catch (ResourceNotFoundException exception) {
            return false;
        }
    }

    private Mono<Boolean> deletedSecretExistsAsync(String name) {
        return asyncClient.getDeletedSecret(name)
                .map(secret -> true)
                .onErrorResume(ResourceNotFoundException.class, exception -> Mono.just(false));
    }

    private void purgeWithRetry(String name, Instant deadline) {
        while (true) {
            try {
                syncClient.purgeDeletedSecret(name);
                return;
            } catch (ResourceNotFoundException exception) {
                waitOrThrow(name, deadline);
            }
        }
    }

    private Mono<Void> purgeWithRetryAsync(String name, Instant deadline) {
        return Mono.defer(() -> asyncClient.purgeDeletedSecret(name))
                .onErrorResume(ResourceNotFoundException.class, exception ->
                        delayOrError(name, deadline)
                                .then(purgeWithRetryAsync(name, deadline)));
    }

    private void waitUntilPurged(String name, Instant deadline) {
        while (true) {
            try {
                syncClient.getDeletedSecret(name);
                waitOrThrow(name, deadline);
            } catch (ResourceNotFoundException exception) {
                return;
            }
        }
    }

    private Mono<Void> waitUntilPurgedAsync(String name, Instant deadline) {
        return Mono.defer(() -> asyncClient.getDeletedSecret(name))
                .flatMap(secret -> delayOrError(name, deadline)
                        .then(waitUntilPurgedAsync(name, deadline)))
                .onErrorResume(ResourceNotFoundException.class, exception -> Mono.empty())
                .then();
    }

    private void waitOrThrow(String name, Instant deadline) {
        if (!Instant.now().isBefore(deadline)) {
            throw purgeTimeout(name);
        }
        try {
            Thread.sleep(pollInterval.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for secret purge", exception);
        }
    }

    private Mono<Long> delayOrError(String name, Instant deadline) {
        if (!Instant.now().isBefore(deadline)) {
            return Mono.error(purgeTimeout(name));
        }
        return Mono.delay(pollInterval);
    }

    private IllegalStateException purgeTimeout(String name) {
        return new IllegalStateException(
                "Timed out waiting for secret '" + name
                        + "' to be purged. Purge protection may be enabled.");
    }

    private static void validateRotation(
            String name,
            String newValue,
            OffsetDateTime expiresOn) {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("name must not be blank");
        }
        if (newValue == null || newValue.isBlank()) {
            throw new IllegalArgumentException("newValue must not be blank");
        }
        Objects.requireNonNull(expiresOn, "expiresOn");
        if (!expiresOn.isAfter(OffsetDateTime.now())) {
            throw new IllegalArgumentException("expiresOn must be in the future");
        }
    }

    private static KeyVaultSecret secretWithExpiry(
            String name,
            String value,
            OffsetDateTime expiresOn) {
        KeyVaultSecret secret = new KeyVaultSecret(name, value);
        secret.getProperties().setExpiresOn(expiresOn);
        return secret;
    }

    private static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }
}
