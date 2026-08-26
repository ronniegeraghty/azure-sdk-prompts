package com.example.keyvault;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class KeyVaultAsyncSecretProvider implements AsyncSecretProvider {
    private final SecretAsyncClient client;

    public KeyVaultAsyncSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    @Override
    public Mono<SecretSnapshot> get(String name, String defaultValue) {
        return retrieve(name, null, defaultValue);
    }

    @Override
    public Mono<SecretSnapshot> getVersion(String name, String version, String defaultValue) {
        Objects.requireNonNull(version, "version");
        return retrieve(name, version, defaultValue);
    }

    private Mono<SecretSnapshot> retrieve(String name, String version, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        Mono<KeyVaultSecret> request = version == null
                ? client.getSecret(name)
                : client.getSecret(name, version);
        return request
                .map(secret -> new SecretSnapshot(
                        secret.getName(),
                        secret.getValue(),
                        secret.getProperties().getVersion(),
                        secret.getProperties().getExpiresOn(),
                        false))
                .onErrorResume(
                        ResourceNotFoundException.class,
                        ignored -> Mono.just(SecretSnapshot.missing(name, defaultValue)));
    }
}
