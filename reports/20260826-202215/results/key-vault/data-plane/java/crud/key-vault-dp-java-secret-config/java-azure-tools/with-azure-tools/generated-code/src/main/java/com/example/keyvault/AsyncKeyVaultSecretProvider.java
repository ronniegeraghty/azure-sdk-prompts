package com.example.keyvault;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import java.util.Objects;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public class AsyncKeyVaultSecretProvider {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncKeyVaultSecretProvider.class);

    private final SecretAsyncClient client;

    public AsyncKeyVaultSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<SecretSnapshot> getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    public Mono<SecretSnapshot> getSecret(String name, String version, String defaultValue) {
        requireText(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");

        Mono<KeyVaultSecret> request = version == null
            ? client.getSecret(name)
            : client.getSecret(name, version);
        return request
            .map(AsyncKeyVaultSecretProvider::toSnapshot)
            .onErrorResume(ResourceNotFoundException.class, exception -> {
                LOGGER.warn("Key Vault secret '{}' was not found; using its configured default", name);
                return Mono.just(new SecretSnapshot(name, defaultValue, version, null, true));
            });
    }

    private static SecretSnapshot toSnapshot(KeyVaultSecret secret) {
        return new SecretSnapshot(
            secret.getName(),
            secret.getValue(),
            secret.getProperties().getVersion(),
            secret.getProperties().getExpiresOn(),
            false
        );
    }

    private static void requireText(String value, String label) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(label + " must not be blank");
        }
    }
}
