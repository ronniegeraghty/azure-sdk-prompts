package com.example.azureencrypted;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.util.Map;

public final class AzureConfiguration {
    private final TokenCredential credential;
    private final BlobContainerClient blobContainerClient;
    private final BlobContainerAsyncClient blobContainerAsyncClient;
    private final KeyClient keyClient;
    private final KeyAsyncClient keyAsyncClient;
    private final CryptographyClient cryptographyClient;
    private final CryptographyAsyncClient cryptographyAsyncClient;
    private final String keyId;

    private AzureConfiguration(
            TokenCredential credential,
            BlobContainerClient blobContainerClient,
            BlobContainerAsyncClient blobContainerAsyncClient,
            KeyClient keyClient,
            KeyAsyncClient keyAsyncClient,
            CryptographyClient cryptographyClient,
            CryptographyAsyncClient cryptographyAsyncClient,
            String keyId) {
        this.credential = credential;
        this.blobContainerClient = blobContainerClient;
        this.blobContainerAsyncClient = blobContainerAsyncClient;
        this.keyClient = keyClient;
        this.keyAsyncClient = keyAsyncClient;
        this.cryptographyClient = cryptographyClient;
        this.cryptographyAsyncClient = cryptographyAsyncClient;
        this.keyId = keyId;
    }

    public static AzureConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static AzureConfiguration fromEnvironment(Map<String, String> environment) {
        String storageEndpoint = required(environment, "AZURE_STORAGE_BLOB_ENDPOINT");
        String containerName = required(environment, "AZURE_STORAGE_CONTAINER");
        String vaultEndpoint = required(environment, "AZURE_KEY_VAULT_ENDPOINT");
        String keyName = required(environment, "AZURE_KEY_NAME");

        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()
                .endpoint(storageEndpoint)
                .credential(credential);
        BlobServiceClient blobServiceClient = blobBuilder.buildClient();
        BlobServiceAsyncClient blobServiceAsyncClient = blobBuilder.buildAsyncClient();

        KeyClientBuilder keyBuilder = new KeyClientBuilder()
                .vaultUrl(vaultEndpoint)
                .credential(credential);
        KeyClient keyClient = keyBuilder.buildClient();
        KeyAsyncClient keyAsyncClient = keyBuilder.buildAsyncClient();

        String keyId = keyClient.getKey(keyName).getId();
        CryptographyClientBuilder cryptographyBuilder = new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential);

        return new AzureConfiguration(
                credential,
                blobServiceClient.getBlobContainerClient(containerName),
                blobServiceAsyncClient.getBlobContainerAsyncClient(containerName),
                keyClient,
                keyAsyncClient,
                cryptographyBuilder.buildClient(),
                cryptographyBuilder.buildAsyncClient(),
                keyId);
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

    public CryptographyClient cryptographyClient() {
        return cryptographyClient;
    }

    public CryptographyAsyncClient cryptographyAsyncClient() {
        return cryptographyAsyncClient;
    }

    public String keyId() {
        return keyId;
    }
}
