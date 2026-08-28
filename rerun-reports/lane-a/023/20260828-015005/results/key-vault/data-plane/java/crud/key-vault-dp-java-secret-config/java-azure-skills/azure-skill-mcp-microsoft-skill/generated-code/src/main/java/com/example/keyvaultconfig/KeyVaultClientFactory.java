package com.example.keyvaultconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

import java.net.URI;
import java.util.Map;
import java.util.Objects;

public final class KeyVaultClientFactory {
    public static final String VAULT_URL_ENV = "AZURE_KEYVAULT_URL";
    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = "AZURE_CLIENT_ID";

    private KeyVaultClientFactory() {
    }

    public static Clients fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static Clients fromEnvironment(Map<String, String> environment) {
        String vaultUrl = requireVaultUrl(environment.get(VAULT_URL_ENV));
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String clientId = environment.get(MANAGED_IDENTITY_CLIENT_ID_ENV);
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();
        SecretClientBuilder clientBuilder = new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential);
        return new Clients(clientBuilder.buildClient(), clientBuilder.buildAsyncClient());
    }

    private static String requireVaultUrl(String value) {
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(VAULT_URL_ENV + " must be set");
        }
        URI uri = URI.create(value);
        if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {
            throw new IllegalArgumentException(VAULT_URL_ENV + " must be an absolute HTTPS URL");
        }
        return value;
    }

    public record Clients(SecretClient sync, SecretAsyncClient async) {
        public Clients {
            Objects.requireNonNull(sync, "sync");
            Objects.requireNonNull(async, "async");
        }
    }
}
