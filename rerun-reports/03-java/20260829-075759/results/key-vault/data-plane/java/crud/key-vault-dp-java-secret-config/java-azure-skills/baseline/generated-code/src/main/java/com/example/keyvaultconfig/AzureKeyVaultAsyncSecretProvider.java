package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AzureKeyVaultAsyncSecretProvider implements AsyncSecretProvider {
    private final SecretAsyncClient client;

    public AzureKeyVaultAsyncSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    @Override
    public Mono<SecretValue> getSecret(String name, String defaultValue) {
        return retrieve(name, null, defaultValue);
    }

    @Override
    public Mono<SecretValue> getSecret(String name, String version, String defaultValue) {
        Objects.requireNonNull(version, "version");
        return retrieve(name, version, defaultValue);
    }

    private Mono<SecretValue> retrieve(String name, String version, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        Mono<KeyVaultSecret> request = version == null
                ? client.getSecret(name)
                : client.getSecret(name, version);

        return request
                .map(secret -> new SecretValue(
                        secret.getName(),
                        secret.getValue(),
                        secret.getProperties().getExpiresOn()))
                .onErrorResume(
                        ResourceNotFoundException.class,
                        exception -> Mono.just(new SecretValue(name, defaultValue, null)));
    }
}
