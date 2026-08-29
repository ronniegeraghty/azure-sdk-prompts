package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncKeyVaultSecretProvider implements AsyncSecretProvider {
    private final SecretAsyncClient client;

    public AsyncKeyVaultSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    @Override
    public Mono<SecretSnapshot> getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    @Override
    public Mono<SecretSnapshot> getSecret(String name, String version, String defaultValue) {
        validate(name, defaultValue);
        return client.getSecret(name, version)
                .map(AsyncKeyVaultSecretProvider::toSnapshot)
                .onErrorResume(ResourceNotFoundException.class,
                        exception -> Mono.just(missing(name, defaultValue)));
    }

    private static SecretSnapshot toSnapshot(KeyVaultSecret secret) {
        return new SecretSnapshot(
                secret.getName(),
                secret.getValue(),
                secret.getProperties().getVersion(),
                secret.getProperties().getExpiresOn(),
                true);
    }

    private static SecretSnapshot missing(String name, String defaultValue) {
        return new SecretSnapshot(name, defaultValue, null, null, false);
    }

    private static void validate(String name, String defaultValue) {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("name must not be blank");
        }
        Objects.requireNonNull(defaultValue, "defaultValue");
    }
}
