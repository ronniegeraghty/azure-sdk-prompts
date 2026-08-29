package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncSecretProvider {
    private static final System.Logger LOGGER = System.getLogger(AsyncSecretProvider.class.getName());

    private final SecretAsyncClient client;

    public AsyncSecretProvider(SecretAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<SecretValue> getSecret(String name, String defaultValue) {
        return getSecretVersion(name, null, defaultValue);
    }

    public Mono<SecretValue> getSecretVersion(String name, String version, String defaultValue) {
        requireText(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");

        Mono<KeyVaultSecret> request = version == null
                ? client.getSecret(name)
                : client.getSecret(name, version);

        return request
                .map(AsyncSecretProvider::fromKeyVault)
                .onErrorResume(ResourceNotFoundException.class, exception -> {
                    LOGGER.log(System.Logger.Level.WARNING,
                            "Secret \"{0}\" was not found; using its configured default.", name);
                    return Mono.just(new SecretValue(name, defaultValue, version, null, true));
                });
    }

    private static SecretValue fromKeyVault(KeyVaultSecret secret) {
        return new SecretValue(
                secret.getName(),
                secret.getValue(),
                secret.getProperties().getVersion(),
                secret.getProperties().getExpiresOn(),
                false);
    }

    private static void requireText(String value, String field) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(field + " must not be blank");
        }
    }
}
