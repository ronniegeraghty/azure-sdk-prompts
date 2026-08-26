package com.example.keyvault;

import com.azure.identity.ManagedIdentityCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import java.net.URI;
import java.net.URISyntaxException;

public final class KeyVaultClientFactory {
    public static final String VAULT_URL_ENV = "AZURE_KEY_VAULT_URL";
    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = "AZURE_MANAGED_IDENTITY_CLIENT_ID";

    private KeyVaultClientFactory() {
    }

    public static KeyVaultClients fromEnvironment() {
        String vaultUrl = requireEnvironment(VAULT_URL_ENV);
        validateHttpsUrl(vaultUrl);

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = System.getenv(MANAGED_IDENTITY_CLIENT_ID_ENV);
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        ManagedIdentityCredential credential = credentialBuilder.build();
        SecretClientBuilder clientBuilder = new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(credential);

        return new KeyVaultClients(clientBuilder.buildClient(), clientBuilder.buildAsyncClient());
    }

    private static String requireEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static void validateHttpsUrl(String value) {
        try {
            URI uri = new URI(value);
            if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {
                throw new IllegalStateException(VAULT_URL_ENV + " must be an absolute HTTPS URL");
            }
        } catch (URISyntaxException exception) {
            throw new IllegalStateException(VAULT_URL_ENV + " is not a valid URL", exception);
        }
    }

    public record KeyVaultClients(SecretClient syncClient, SecretAsyncClient asyncClient) {
    }
}
