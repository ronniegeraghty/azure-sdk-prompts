package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.time.Duration;
import java.time.Instant;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class SecretRotationHelper {
    private final SecretClient client;
    private final Duration purgeTimeout;
    private final Duration pollInterval;

    public SecretRotationHelper(SecretClient client, Duration purgeTimeout, Duration pollInterval) {
        this.client = Objects.requireNonNull(client, "client");
        this.purgeTimeout = requirePositive(purgeTimeout, "purgeTimeout");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
    }

    public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        client.beginDeleteSecret(name).waitForCompletion();
        client.purgeDeletedSecret(name);
        waitUntilPurged(name);

        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue);
        replacement.getProperties().setExpiresOn(expiresOn);
        return client.setSecret(replacement);
    }

    private void waitUntilPurged(String name) {
        Instant deadline = Instant.now().plus(purgeTimeout);
        while (Instant.now().isBefore(deadline)) {
            try {
                client.getDeletedSecret(name);
            } catch (ResourceNotFoundException exception) {
                return;
            }
            sleep();
        }
        throw new IllegalStateException(
                "Timed out waiting for deleted secret to be purged: " + name);
    }

    private void sleep() {
        try {
            Thread.sleep(pollInterval.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for secret purge", exception);
        }
    }

    private static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }
}
