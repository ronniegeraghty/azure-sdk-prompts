package com.example.encryptedblob;

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
import com.azure.storage.blob.BlobContainerClientBuilder;

import java.util.Map;

public final class AzureConfiguration {
    public static final String BLOB_ENDPOINT_ENV = "AZURE_STORAGE_BLOB_ENDPOINT";
    public static final String CONTAINER_NAME_ENV = "AZURE_STORAGE_CONTAINER_NAME";
    public static final String VAULT_ENDPOINT_ENV = "AZURE_KEY_VAULT_ENDPOINT";
    public static final String KEY_NAME_ENV = "AZURE_KEY_VAULT_KEY_NAME";

    private final TokenCredential credential;
    private final String blobEndpoint;
    private final String containerName;
    private final String vaultEndpoint;
    private final String keyName;

    private AzureConfiguration(
            TokenCredential credential,
            String blobEndpoint,
            String containerName,
            String vaultEndpoint,
            String keyName) {
        this.credential = credential;
        this.blobEndpoint = blobEndpoint;
        this.containerName = containerName;
        this.vaultEndpoint = vaultEndpoint;
        this.keyName = keyName;
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        String blobEndpoint = required(environment, BLOB_ENDPOINT_ENV);
        String containerName = required(environment, CONTAINER_NAME_ENV);
        String vaultEndpoint = required(environment, VAULT_ENDPOINT_ENV);
        String keyName = required(environment, KEY_NAME_ENV);

        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();
        return new AzureConfiguration(
                credential, blobEndpoint, containerName, vaultEndpoint, keyName);
    }

    public BlobContainerClient blobContainerClient() {
        return new BlobContainerClientBuilder()
                .endpoint(blobEndpoint)
                .containerName(containerName)
                .credential(credential)
                .buildClient();
    }

    public BlobContainerAsyncClient blobContainerAsyncClient() {
        return new BlobContainerClientBuilder()
                .endpoint(blobEndpoint)
                .containerName(containerName)
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

    public CryptographyClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildClient();
    }

    public CryptographyAsyncClient cryptographyAsyncClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildAsyncClient();
    }

    public String keyName() {
        return keyName;
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
