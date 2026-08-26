package com.example.keyvault;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

public final class KeyVaultClients {
    public static final String VAULT_URL_ENV = "KEY_VAULT_URL";

    private KeyVaultClients() {
    }

    public static SecretClient syncClient() {
        return builder().buildClient();
    }

    public static SecretAsyncClient asyncClient() {
        return builder().buildAsyncClient();
    }

    private static SecretClientBuilder builder() {
        String vaultUrl = requireEnvironment(VAULT_URL_ENV);
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();
        return new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(credential);
    }

    private static String requireEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
