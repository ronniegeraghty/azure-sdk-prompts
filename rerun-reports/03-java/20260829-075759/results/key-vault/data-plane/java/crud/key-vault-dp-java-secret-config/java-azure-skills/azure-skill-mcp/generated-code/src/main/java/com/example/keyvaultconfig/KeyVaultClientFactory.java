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

public final class KeyVaultClientFactory {
    public static final String VAULT_URL_ENV = "KEY_VAULT_URL";
    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = "AZURE_CLIENT_ID";

    private KeyVaultClientFactory() {
    }

    public static Clients fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static Clients fromEnvironment(Map<String, String> environment) {
        Objects.requireNonNull(environment, "environment");
        String vaultUrl = requireHttpsVaultUrl(environment.get(VAULT_URL_ENV));
        String clientId = environment.get(MANAGED_IDENTITY_CLIENT_ID_ENV);

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();
        SecretClientBuilder clientBuilder = new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential);

        return new Clients(clientBuilder.buildClient(), clientBuilder.buildAsyncClient());
    }

    private static String requireHttpsVaultUrl(String value) {
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(VAULT_URL_ENV + " must be set");
        }
        try {
            URI uri = new URI(value);
            if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {
                throw new IllegalStateException(VAULT_URL_ENV + " must be an HTTPS URL");
            }
            return uri.toString();
        } catch (URISyntaxException exception) {
            throw new IllegalStateException(VAULT_URL_ENV + " is not a valid URL", exception);
        }
    }

    public record Clients(SecretClient syncClient, SecretAsyncClient asyncClient) {
        public Clients {
            Objects.requireNonNull(syncClient, "syncClient");
            Objects.requireNonNull(asyncClient, "asyncClient");
        }
    }
}
