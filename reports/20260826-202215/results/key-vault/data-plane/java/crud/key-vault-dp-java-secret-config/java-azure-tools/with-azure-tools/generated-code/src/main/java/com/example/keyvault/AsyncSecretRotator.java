package com.example.keyvault;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;
import reactor.core.publisher.Mono;

public final class AsyncSecretRotator {
    private final SecretAsyncClient client;
    private final Duration pollInterval;
    private final Duration timeout;

    public AsyncSecretRotator(SecretAsyncClient client, Duration pollInterval, Duration timeout) {
        this.client = Objects.requireNonNull(client, "client");
        this.pollInterval = SyncSecretRotator.requirePositive(pollInterval, "pollInterval");
        this.timeout = SyncSecretRotator.requirePositive(timeout, "timeout");
    }

    public Mono<KeyVaultSecret> rotate(String name, String newValue, OffsetDateTime expiresOn) {
        SyncSecretRotator.requireRotationArguments(name, newValue, expiresOn);
        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue)
            .setProperties(new SecretProperties().setExpiresOn(expiresOn));

        return client.beginDeleteSecret(name)
            .last()
            .then(client.purgeDeletedSecret(name))
            .then(waitUntilPurged(name))
            .then(client.setSecret(replacement));
    }

    private Mono<Void> waitUntilPurged(String name) {
        return Mono.defer(() -> client.getDeletedSecret(name)
                .then(Mono.delay(pollInterval))
                .then(waitUntilPurged(name)))
            .onErrorResume(ResourceNotFoundException.class, exception -> Mono.empty())
            .timeout(timeout, Mono.error(new IllegalStateException(
                "Timed out waiting for purged secret name to become available: " + name)));
    }
}
