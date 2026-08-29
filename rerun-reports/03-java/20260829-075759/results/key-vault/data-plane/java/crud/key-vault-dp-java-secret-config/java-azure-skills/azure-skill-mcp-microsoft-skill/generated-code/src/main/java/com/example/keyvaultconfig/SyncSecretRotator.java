package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;

import java.time.Duration;
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
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        client.beginDeleteSecret(name).waitForCompletion();
        purgeWhenVisible(name);
        waitUntilPurged(name);

        return client.setSecret(new KeyVaultSecret(name, newValue)
                .setProperties(new SecretProperties().setExpiresOn(expiresOn)));
    }

    private void purgeWhenVisible(String name) {
        long deadline = System.nanoTime() + timeout.toNanos();
        while (true) {
            try {
                client.purgeDeletedSecret(name);
                return;
            } catch (ResourceNotFoundException exception) {
                if (System.nanoTime() >= deadline) {
                    throw new IllegalStateException(
                            "Deleted secret did not become purgeable before timeout: " + name,
                            exception);
                }
                sleep();
            }
        }
    }

    private void waitUntilPurged(String name) {
        long deadline = System.nanoTime() + timeout.toNanos();
        while (true) {
            try {
                client.getDeletedSecret(name);
                if (System.nanoTime() >= deadline) {
                    throw new IllegalStateException(
                            "Secret purge did not complete before timeout: " + name);
                }
                sleep();
            } catch (ResourceNotFoundException exception) {
                return;
            }
        }
    }

    private void sleep() {
        try {
            Thread.sleep(pollInterval.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for secret deletion", exception);
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
