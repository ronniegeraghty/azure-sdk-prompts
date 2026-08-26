package com.example.keyvault;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;
import java.time.Duration;
import java.time.Instant;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class SyncSecretRotator {
    private final SecretClient client;
    private final Duration pollInterval;
    private final Duration timeout;

    public SyncSecretRotator(SecretClient client, Duration pollInterval, Duration timeout) {
        this.client = Objects.requireNonNull(client, "client");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
        this.timeout = requirePositive(timeout, "timeout");
    }

    public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {
        requireRotationArguments(name, newValue, expiresOn);
        client.beginDeleteSecret(name).waitForCompletion();
        client.purgeDeletedSecret(name);
        waitUntilPurged(name);

        return client.setSecret(new KeyVaultSecret(name, newValue)
            .setProperties(new SecretProperties().setExpiresOn(expiresOn)));
    }

    private void waitUntilPurged(String name) {
        Instant deadline = Instant.now().plus(timeout);
        while (Instant.now().isBefore(deadline)) {
            try {
                client.getDeletedSecret(name);
                sleep();
            } catch (ResourceNotFoundException exception) {
                return;
            }
        }
        throw new IllegalStateException("Timed out waiting for purged secret name to become available: " + name);
    }

    private void sleep() {
        try {
            Thread.sleep(pollInterval.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for secret purge", exception);
        }
    }

    static void requireRotationArguments(String name, String newValue, OffsetDateTime expiresOn) {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("name must not be blank");
        }
        if (newValue == null || newValue.isEmpty()) {
            throw new IllegalArgumentException("newValue must not be empty");
        }
        Objects.requireNonNull(expiresOn, "expiresOn");
    }

    static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }
}
