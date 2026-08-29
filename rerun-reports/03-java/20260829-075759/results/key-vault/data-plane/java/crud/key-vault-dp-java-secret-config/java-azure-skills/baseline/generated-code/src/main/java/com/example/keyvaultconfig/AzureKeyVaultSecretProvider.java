package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.util.Objects;

public final class AzureKeyVaultSecretProvider implements SecretProvider {
    private final SecretClient client;

    public AzureKeyVaultSecretProvider(SecretClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    @Override
    public SecretValue getSecret(String name, String defaultValue) {
        return retrieve(name, null, defaultValue);
    }

    @Override
    public SecretValue getSecret(String name, String version, String defaultValue) {
        Objects.requireNonNull(version, "version");
        return retrieve(name, version, defaultValue);
    }

    private SecretValue retrieve(String name, String version, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        try {
            KeyVaultSecret secret = version == null
                    ? client.getSecret(name)
                    : client.getSecret(name, version);
            return new SecretValue(
                    secret.getName(),
                    secret.getValue(),
                    secret.getProperties().getExpiresOn());
        } catch (ResourceNotFoundException exception) {
            return new SecretValue(name, defaultValue, null);
        }
    }
}
