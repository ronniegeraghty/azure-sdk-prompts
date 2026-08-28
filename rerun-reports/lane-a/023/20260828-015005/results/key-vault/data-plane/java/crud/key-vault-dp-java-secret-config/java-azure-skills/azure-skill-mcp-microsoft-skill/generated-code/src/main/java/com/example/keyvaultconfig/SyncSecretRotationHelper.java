package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;

import java.time.Duration;
import java.time.Instant;
import java.time.OffsetDateTime;
import java.util.Objects;

public final class SyncSecretRotationHelper {
    private final SecretClient client;
    private final Duration purgeTimeout;
    private final Duration pollInterval;

    public SyncSecretRotationHelper(SecretClient client) {
        this(client, Duration.ofMinutes(2), Duration.ofSeconds(2));
    }

    public SyncSecretRotationHelper(
            SecretClient client,
            Duration purgeTimeout,
            Duration pollInterval) {
        this.client = Objects.requireNonNull(client, "client");
        this.purgeTimeout = requirePositive(purgeTimeout, "purgeTimeout");
        this.pollInterval = requirePositive(pollInterval, "pollInterval");
    }

    public KeyVaultSecret rotate(
            String name,
            String newValue,
            OffsetDateTime expiresOn) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(newValue, "newValue");
        Objects.requireNonNull(expiresOn, "expiresOn");

        SyncPoller<DeletedSecret, Void> deletePoller = client.beginDeleteSecret(name);
        deletePoller.waitForCompletion();
        client.purgeDeletedSecret(name);
        waitUntilPurged(name);

        return client.setSecret(new KeyVaultSecret(name, newValue)
                .setProperties(new SecretProperties().setExpiresOn(expiresOn)));
    }

    private void waitUntilPurged(String name) {
        Instant deadline = Instant.now().plus(purgeTimeout);
        while (Instant.now().isBefore(deadline)) {
            try {
                client.getDeletedSecret(name);
                sleep();
            } catch (ResourceNotFoundException exception) {
                return;
            }
        }
        throw new IllegalStateException(
                "Timed out waiting for secret '" + name + "' to be fully purged");
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
