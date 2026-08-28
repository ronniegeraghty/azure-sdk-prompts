package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncSecretProvider {
    private final SecretAsyncClient client;

    public AsyncSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<SecretValue> getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    public Mono<SecretValue> getSecret(String name, String version, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        Mono<KeyVaultSecret> request = version == null
                ? client.getSecret(name)
                : client.getSecret(name, version);
        return request
                .map(AsyncSecretProvider::toSecretValue)
                .onErrorResume(
                        ResourceNotFoundException.class,
                        exception -> Mono.just(SecretValue.missing(name, defaultValue)));
    }

    private static SecretValue toSecretValue(KeyVaultSecret secret) {
        return new SecretValue(
                secret.getName(),
                secret.getValue(),
                secret.getProperties().getVersion(),
                secret.getProperties().getExpiresOn(),
                true);
    }
}
