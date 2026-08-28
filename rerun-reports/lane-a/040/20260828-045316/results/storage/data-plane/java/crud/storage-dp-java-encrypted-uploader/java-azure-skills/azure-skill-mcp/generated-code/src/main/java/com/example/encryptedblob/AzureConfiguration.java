package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.util.Map;

public final class AzureConfiguration {
    private final TokenCredential credential;
    private final String keyName;
    private final BlobContainerClient blobContainerClient;
    private final BlobContainerAsyncClient blobContainerAsyncClient;
    private final KeyClient keyClient;
    private final KeyAsyncClient keyAsyncClient;

    private AzureConfiguration(
            TokenCredential credential,
            String keyName,
            BlobContainerClient blobContainerClient,
            BlobContainerAsyncClient blobContainerAsyncClient,
            KeyClient keyClient,
            KeyAsyncClient keyAsyncClient) {
        this.credential = credential;
        this.keyName = keyName;
        this.blobContainerClient = blobContainerClient;
        this.blobContainerAsyncClient = blobContainerAsyncClient;
        this.keyClient = keyClient;
        this.keyAsyncClient = keyAsyncClient;
    }

    public static AzureConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static AzureConfiguration fromEnvironment(Map<String, String> environment) {
        String storageEndpoint = required(environment, "AZURE_STORAGE_BLOB_ENDPOINT");
        String containerName = required(environment, "AZURE_STORAGE_CONTAINER_NAME");
        String vaultEndpoint = required(environment, "AZURE_KEY_VAULT_ENDPOINT");
        String keyName = required(environment, "AZURE_KEY_VAULT_KEY_NAME");
        String managedIdentityClientId = environment.get("AZURE_CLIENT_ID");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();

        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()
                .endpoint(storageEndpoint)
                .credential(credential);
        KeyClientBuilder keyBuilder = new KeyClientBuilder()
                .vaultUrl(vaultEndpoint)
                .credential(credential);

        return new AzureConfiguration(
                credential,
                keyName,
                blobBuilder.buildClient().getBlobContainerClient(containerName),
                blobBuilder.buildAsyncClient().getBlobContainerAsyncClient(containerName),
                keyBuilder.buildClient(),
                keyBuilder.buildAsyncClient());
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    public TokenCredential credential() {
        return credential;
    }

    public String keyName() {
        return keyName;
    }

    public BlobContainerClient blobContainerClient() {
        return blobContainerClient;
    }

    public BlobContainerAsyncClient blobContainerAsyncClient() {
        return blobContainerAsyncClient;
    }

    public KeyClient keyClient() {
        return keyClient;
    }

    public KeyAsyncClient keyAsyncClient() {
        return keyAsyncClient;
    }
}
