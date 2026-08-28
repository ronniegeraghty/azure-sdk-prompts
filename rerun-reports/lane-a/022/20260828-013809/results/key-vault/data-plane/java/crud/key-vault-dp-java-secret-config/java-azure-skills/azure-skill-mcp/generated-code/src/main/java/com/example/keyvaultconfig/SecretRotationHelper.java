package com.example.keyvaultconfig;

import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class SecretRotationHelper {
    private final SecretClient client;
    private final Duration timeout;
    private final Duration pollInterval;

    public SecretRotationHelper(SecretClient client, Duration timeout, Duration pollInterval) {
        this.client = Objects.requireNonNull(client, "client");
        this.timeout = requirePositive(timeout, "timeout");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
    }

    public KeyVaultSecret rotate(
            String name,
            String newValue,
            OffsetDateTime newExpiresOn) {
        requireRotationArguments(name, newValue, newExpiresOn);
        long deadlineNanos = deadlineNanos();

        SyncPoller<DeletedSecret, Void> deletePoller = client.beginDeleteSecret(name);
        deletePoller.waitForCompletion(remaining(deadlineNanos));

        // Soft-deleted names cannot be recreated until the deleted record is purged.
        client.purgeDeletedSecret(name);
        waitUntilPurged(name, deadlineNanos);

        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue);
        replacement.getProperties().setExpiresOn(newExpiresOn);
        return createAfterPurge(replacement, deadlineNanos);
    }

    private void waitUntilPurged(String name, long deadlineNanos) {
        while (true) {
            try {
                client.getDeletedSecret(name);
            } catch (ResourceNotFoundException exception) {
                return;
            }
            sleepOrThrow(deadlineNanos);
        }
    }

    private KeyVaultSecret createAfterPurge(KeyVaultSecret replacement, long deadlineNanos) {
        while (true) {
            try {
                return client.setSecret(replacement);
            } catch (HttpResponseException exception) {
                if (exception.getResponse() == null
                        || exception.getResponse().getStatusCode() != 409) {
                    throw exception;
                }
                sleepOrThrow(deadlineNanos);
            }
        }
    }

    private void sleepOrThrow(long deadlineNanos) {
        Duration remaining = remaining(deadlineNanos);
        Duration sleep = remaining.compareTo(pollInterval) < 0 ? remaining : pollInterval;
        try {
            Thread.sleep(sleep.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while waiting for Key Vault rotation", exception);
        }
    }

    private long deadlineNanos() {
        return System.nanoTime() + timeout.toNanos();
    }

    private static Duration remaining(long deadlineNanos) {
        long nanos = deadlineNanos - System.nanoTime();
        if (nanos <= 0) {
            throw new IllegalStateException("Timed out waiting for Key Vault secret rotation");
        }
        return Duration.ofNanos(nanos);
    }

    static void requireRotationArguments(String name, String newValue, OffsetDateTime newExpiresOn) {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("name must not be blank");
        }
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(newExpiresOn, "newExpiresOn");
    }

    static Duration requirePositive(Duration duration, String field) {
        Objects.requireNonNull(duration, field);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(field + " must be positive");
        }
        return duration;
    }
}
