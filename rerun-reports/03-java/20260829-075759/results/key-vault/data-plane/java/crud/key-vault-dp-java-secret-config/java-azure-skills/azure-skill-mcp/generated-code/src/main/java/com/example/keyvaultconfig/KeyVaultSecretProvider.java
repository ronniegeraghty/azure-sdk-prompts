package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.util.Objects;

public final class KeyVaultSecretProvider implements SecretProvider {
    private final SecretClient client;

    public KeyVaultSecretProvider(SecretClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    @Override
    public SecretSnapshot getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    @Override
    public SecretSnapshot getSecret(String name, String version, String defaultValue) {
        validate(name, defaultValue);
        try {
            KeyVaultSecret secret = client.getSecret(name, version);
            return toSnapshot(secret);
        } catch (ResourceNotFoundException exception) {
            return missing(name, defaultValue);
        }
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
