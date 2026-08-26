package com.example.keyvault;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.util.Objects;

public final class KeyVaultSyncSecretProvider implements SyncSecretProvider {
    private final SecretClient client;

    public KeyVaultSyncSecretProvider(SecretClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    @Override
    public SecretSnapshot get(String name, String defaultValue) {
        return retrieve(name, null, defaultValue);
    }

    @Override
    public SecretSnapshot getVersion(String name, String version, String defaultValue) {
        Objects.requireNonNull(version, "version");
        return retrieve(name, version, defaultValue);
    }

    private SecretSnapshot retrieve(String name, String version, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        try {
            KeyVaultSecret secret = version == null
                    ? client.getSecret(name)
                    : client.getSecret(name, version);
            return new SecretSnapshot(
                    secret.getName(),
                    secret.getValue(),
                    secret.getProperties().getVersion(),
                    secret.getProperties().getExpiresOn(),
                    false);
        } catch (ResourceNotFoundException ignored) {
            return SecretSnapshot.missing(name, defaultValue);
        }
    }
}
