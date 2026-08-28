package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncKeyVaultSecretProvider {
    private final SecretAsyncClient client;

    public AsyncKeyVaultSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<ConfigSecret> getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    public Mono<ConfigSecret> getSecret(String name, String version, String defaultValue) {
        requireText(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");

        Mono<KeyVaultSecret> request = version == null || version.isBlank()
                ? client.getSecret(name)
                : client.getSecret(name, version);

        return request
                .map(secret -> new ConfigSecret(
                        secret.getName(),
                        secret.getValue(),
                        secret.getProperties().getVersion(),
                        secret.getProperties().getExpiresOn(),
                        false))
                .onErrorResume(
                        ResourceNotFoundException.class,
                        exception -> Mono.just(new ConfigSecret(name, defaultValue, version, null, true)));
    }

    private static void requireText(String value, String field) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(field + " must not be blank");
        }
    }
}
