package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.util.Objects;

public final class SyncSecretProvider {
    private final SecretClient client;

    public SyncSecretProvider(SecretClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public SecretValue getSecret(String name, String defaultValue) {
        return getSecret(name, null, defaultValue);
    }

    public SecretValue getSecret(String name, String version, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        try {
            KeyVaultSecret secret = version == null
                    ? client.getSecret(name)
                    : client.getSecret(name, version);
            return toSecretValue(secret);
        } catch (ResourceNotFoundException exception) {
            return SecretValue.missing(name, defaultValue);
        }
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
