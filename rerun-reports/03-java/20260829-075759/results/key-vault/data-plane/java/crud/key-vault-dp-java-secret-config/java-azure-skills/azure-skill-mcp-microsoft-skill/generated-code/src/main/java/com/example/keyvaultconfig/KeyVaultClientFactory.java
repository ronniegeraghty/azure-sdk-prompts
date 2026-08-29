package com.example.keyvaultconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

import java.util.Map;
import java.util.Objects;

public final class KeyVaultClientFactory {
    public static final String VAULT_URL_ENV = "AZURE_KEYVAULT_URL";
    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = "AZURE_CLIENT_ID";

    private KeyVaultClientFactory() {
    }

    public static SecretClient createSyncClient() {
        return createBuilder(System.getenv()).buildClient();
    }

    public static SecretAsyncClient createAsyncClient() {
        return createBuilder(System.getenv()).buildAsyncClient();
    }

    static SecretClientBuilder createBuilder(Map<String, String> environment) {
        Objects.requireNonNull(environment, "environment");
        String vaultUrl = requireEnvironmentVariable(environment, VAULT_URL_ENV);

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String clientId = environment.get(MANAGED_IDENTITY_CLIENT_ID_ENV);
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();

        return new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential);
    }

    private static String requireEnvironmentVariable(
            Map<String, String> environment,
            String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    "Required environment variable " + name + " is not set");
        }
        return value;
    }
}
