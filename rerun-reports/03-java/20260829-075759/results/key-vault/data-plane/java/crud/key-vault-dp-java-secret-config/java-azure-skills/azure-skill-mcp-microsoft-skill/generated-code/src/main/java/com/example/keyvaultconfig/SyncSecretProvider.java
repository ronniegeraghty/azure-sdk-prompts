package com.example.keyvaultconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

import java.util.Objects;

public final class SyncSecretProvider {
    private static final System.Logger LOGGER = System.getLogger(SyncSecretProvider.class.getName());

    private final SecretClient client;

    public SyncSecretProvider(SecretClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public SecretValue getSecret(String name, String defaultValue) {
        return getSecretVersion(name, null, defaultValue);
    }

    public SecretValue getSecretVersion(String name, String version, String defaultValue) {
        requireText(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");

        try {
            KeyVaultSecret secret = version == null
                    ? client.getSecret(name)
                    : client.getSecret(name, version);
            return fromKeyVault(secret);
        } catch (ResourceNotFoundException exception) {
            LOGGER.log(System.Logger.Level.WARNING,
                    "Secret \"{0}\" was not found; using its configured default.", name);
            return new SecretValue(name, defaultValue, version, null, true);
        }
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
