package com.example.keyvault;

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
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        try {
            client.beginDeleteSecret(name).waitForCompletion();
        } catch (ResourceNotFoundException ignored) {
            // The active secret may already have been deleted by an earlier rotation attempt.
        }

        try {
            client.purgeDeletedSecret(name);
            waitUntilPurged(name);
        } catch (ResourceNotFoundException ignored) {
            // No soft-deleted secret remains, so the name is ready to reuse.
        }

        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue)
                .setProperties(new SecretProperties().setExpiresOn(expiresOn));
        return client.setSecret(replacement);
    }

    private void waitUntilPurged(String name) {
        long deadline = System.nanoTime() + timeout.toNanos();
        while (System.nanoTime() < deadline) {
            try {
                client.getDeletedSecret(name);
                sleep();
            } catch (ResourceNotFoundException ignored) {
                return;
            }
        }
        throw new IllegalStateException("Timed out waiting for secret purge: " + name);
    }

    private void sleep() {
        try {
            Thread.sleep(pollInterval.toMillis());
        } catch (InterruptedException interrupted) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for secret purge", interrupted);
        }
    }

    private static Duration requirePositive(Duration value, String name) {
        Objects.requireNonNull(value, name);
        if (value.isZero() || value.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return value;
    }
}
