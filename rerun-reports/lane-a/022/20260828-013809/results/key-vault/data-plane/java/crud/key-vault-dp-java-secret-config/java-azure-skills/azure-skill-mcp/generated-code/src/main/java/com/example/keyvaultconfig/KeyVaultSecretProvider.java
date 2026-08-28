package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.util.Objects;

public final class KeyVaultSecretProvider {
    private final SecretClient client;

    public KeyVaultSecretProvider(SecretClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public ConfigSecret getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    public ConfigSecret getSecret(String name, String version, String defaultValue) {
        requireText(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");

        try {
            KeyVaultSecret secret = version == null || version.isBlank()
                    ? client.getSecret(name)
                    : client.getSecret(name, version);
            return new ConfigSecret(
                    secret.getName(),
                    secret.getValue(),
                    secret.getProperties().getVersion(),
                    secret.getProperties().getExpiresOn(),
                    false);
        } catch (ResourceNotFoundException exception) {
            return new ConfigSecret(name, defaultValue, version, null, true);
        }
    }

    private static void requireText(String value, String field) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(field + " must not be blank");
        }
    }
}
