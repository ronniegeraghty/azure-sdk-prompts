package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class AzureConfiguration {
    private final TokenCredential credential;
    private final BlobContainerClient blobContainerClient;
    private final BlobContainerAsyncClient blobContainerAsyncClient;
    private final KeyClient keyClient;
    private final KeyAsyncClient keyAsyncClient;
    private final String keyName;

    private AzureConfiguration(
        TokenCredential credential,
        BlobContainerClient blobContainerClient,
        BlobContainerAsyncClient blobContainerAsyncClient,
        KeyClient keyClient,
        KeyAsyncClient keyAsyncClient,
        String keyName
    ) {
        this.credential = credential;
        this.blobContainerClient = blobContainerClient;
        this.blobContainerAsyncClient = blobContainerAsyncClient;
        this.keyClient = keyClient;
        this.keyAsyncClient = keyAsyncClient;
        this.keyName = keyName;
    }

    public static AzureConfiguration fromEnvironment() {
        String blobEndpoint = requiredEnvironment("AZURE_STORAGE_BLOB_ENDPOINT");
        String containerName = requiredEnvironment("AZURE_STORAGE_CONTAINER");
        String vaultEndpoint = requiredEnvironment("AZURE_KEY_VAULT_ENDPOINT");
        String keyName = requiredEnvironment("AZURE_KEY_VAULT_KEY_NAME");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();

        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()
            .endpoint(blobEndpoint)
            .credential(credential);
        BlobServiceClient blobServiceClient = blobBuilder.buildClient();
        BlobServiceAsyncClient blobServiceAsyncClient = blobBuilder.buildAsyncClient();

        KeyClientBuilder keyBuilder = new KeyClientBuilder()
            .vaultUrl(vaultEndpoint)
            .credential(credential);

        return new AzureConfiguration(
            credential,
            blobServiceClient.getBlobContainerClient(containerName),
            blobServiceAsyncClient.getBlobContainerAsyncClient(containerName),
            keyBuilder.buildClient(),
            keyBuilder.buildAsyncClient(),
            keyName);
    }

    public TokenCredential credential() {
        return credential;
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

    public String keyName() {
        return keyName;
    }

    private static String requiredEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
