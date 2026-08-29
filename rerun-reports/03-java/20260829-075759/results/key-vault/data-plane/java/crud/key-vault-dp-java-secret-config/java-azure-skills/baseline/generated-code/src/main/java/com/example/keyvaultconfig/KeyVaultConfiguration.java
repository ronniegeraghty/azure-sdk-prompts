package com.example.keyvaultconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

import java.net.URI;
import java.net.URISyntaxException;
import java.util.Map;
import java.util.Objects;

public final class KeyVaultConfiguration {
    public static final String VAULT_URL_ENVIRONMENT_VARIABLE = "KEY_VAULT_URL";

    private final SecretClient secretClient;
    private final SecretAsyncClient secretAsyncClient;

    private KeyVaultConfiguration(SecretClient secretClient, SecretAsyncClient secretAsyncClient) {
        this.secretClient = secretClient;
        this.secretAsyncClient = secretAsyncClient;
    }

    public static KeyVaultConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static KeyVaultConfiguration fromEnvironment(Map<String, String> environment) {
        Objects.requireNonNull(environment, "environment");
        String vaultUrl = environment.get(VAULT_URL_ENVIRONMENT_VARIABLE);
        if (vaultUrl == null || vaultUrl.isBlank()) {
            throw new IllegalStateException(
                    VAULT_URL_ENVIRONMENT_VARIABLE + " must contain the Azure Key Vault URL");
        }
        validateVaultUrl(vaultUrl);

        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();
        SecretClientBuilder builder = new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential);
        return new KeyVaultConfiguration(builder.buildClient(), builder.buildAsyncClient());
    }

    private static void validateVaultUrl(String vaultUrl) {
        try {
            URI uri = new URI(vaultUrl);
            if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {
                throw new IllegalStateException("KEY_VAULT_URL must be an absolute HTTPS URL");
            }
        } catch (URISyntaxException exception) {
            throw new IllegalStateException("KEY_VAULT_URL is not a valid URL", exception);
        }
    }

    public SecretClient secretClient() {
        return secretClient;
    }

    public SecretAsyncClient secretAsyncClient() {
        return secretAsyncClient;
    }

    public SecretProvider secretProvider() {
        return new AzureKeyVaultSecretProvider(secretClient);
    }

    public AsyncSecretProvider asyncSecretProvider() {
        return new AzureKeyVaultAsyncSecretProvider(secretAsyncClient);
    }
}
