package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.util.Map;

public final class AzureClientConfiguration {
    private final TokenCredential credential;
    private final String storageEndpoint;
    private final String vaultEndpoint;
    private final String keyName;
    private final String containerName;

    private AzureClientConfiguration(
            TokenCredential credential,
            String storageEndpoint,
            String vaultEndpoint,
            String keyName,
            String containerName) {
        this.credential = credential;
        this.storageEndpoint = storageEndpoint;
        this.vaultEndpoint = vaultEndpoint;
        this.keyName = keyName;
        this.containerName = containerName;
    }

    public static AzureClientConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String clientId = environment.get("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }

        return new AzureClientConfiguration(
                credentialBuilder.build(),
                requireEnvironmentVariable(environment, "AZURE_STORAGE_BLOB_ENDPOINT"),
                requireEnvironmentVariable(environment, "AZURE_KEY_VAULT_ENDPOINT"),
                requireEnvironmentVariable(environment, "AZURE_KEY_VAULT_KEY_NAME"),
                requireEnvironmentVariable(environment, "AZURE_STORAGE_CONTAINER"));
    }

    public BlobServiceClient blobServiceClient() {
        return new BlobServiceClientBuilder()
                .endpoint(storageEndpoint)
                .credential(credential)
                .buildClient();
    }

    public BlobServiceAsyncClient blobServiceAsyncClient() {
        return new BlobServiceClientBuilder()
                .endpoint(storageEndpoint)
                .credential(credential)
                .buildAsyncClient();
    }

    public KeyClient keyClient() {
        return new KeyClientBuilder()
                .vaultUrl(vaultEndpoint)
                .credential(credential)
                .buildClient();
    }

    public KeyAsyncClient keyAsyncClient() {
        return new KeyClientBuilder()
                .vaultUrl(vaultEndpoint)
                .credential(credential)
                .buildAsyncClient();
    }

    public TokenCredential credential() {
        return credential;
    }

    public String keyName() {
        return keyName;
    }

    public String containerName() {
        return containerName;
    }

    private static String requireEnvironmentVariable(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is missing: " + name);
        }
        return value;
    }
}
