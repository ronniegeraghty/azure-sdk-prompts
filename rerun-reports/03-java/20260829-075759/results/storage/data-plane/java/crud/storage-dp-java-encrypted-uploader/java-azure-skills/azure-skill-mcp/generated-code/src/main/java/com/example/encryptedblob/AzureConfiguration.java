package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import reactor.core.publisher.Mono;

import java.util.Map;
import java.util.Objects;

public final class AzureConfiguration {
    private final TokenCredential credential;
    private final String keyVaultEndpoint;
    private final String keyName;
    private final BlobContainerClient blobContainerClient;
    private final BlobContainerAsyncClient blobContainerAsyncClient;
    private final KeyClient keyClient;
    private final KeyAsyncClient keyAsyncClient;

    private AzureConfiguration(
        TokenCredential credential,
        String storageEndpoint,
        String containerName,
        String keyVaultEndpoint,
        String keyName
    ) {
        this.credential = Objects.requireNonNull(credential, "credential");
        this.keyVaultEndpoint = Objects.requireNonNull(keyVaultEndpoint, "keyVaultEndpoint");
        this.keyName = Objects.requireNonNull(keyName, "keyName");

        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()
            .endpoint(storageEndpoint)
            .credential(credential);
        BlobServiceClient blobServiceClient = blobBuilder.buildClient();
        BlobServiceAsyncClient blobServiceAsyncClient = blobBuilder.buildAsyncClient();
        blobContainerClient = blobServiceClient.getBlobContainerClient(containerName);
        blobContainerAsyncClient = blobServiceAsyncClient.getBlobContainerAsyncClient(containerName);

        KeyClientBuilder keyBuilder = new KeyClientBuilder()
            .vaultUrl(keyVaultEndpoint)
            .credential(credential);
        keyClient = keyBuilder.buildClient();
        keyAsyncClient = keyBuilder.buildAsyncClient();
    }

    public static AzureConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static AzureConfiguration fromEnvironment(Map<String, String> environment) {
        String managedIdentityClientId = environment.get("AZURE_CLIENT_ID");
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential sharedCredential = credentialBuilder.build();

        return new AzureConfiguration(
            sharedCredential,
            required(environment, "AZURE_STORAGE_BLOB_ENDPOINT"),
            required(environment, "AZURE_STORAGE_CONTAINER"),
            required(environment, "AZURE_KEY_VAULT_ENDPOINT"),
            required(environment, "AZURE_KEY_VAULT_KEY_NAME"));
    }

    public EncryptedBlobStore syncBlobStore(String blobName) {
        String versionedKeyId = keyClient.getKey(keyName).getId();
        KeyManagementService keyManagement = new KeyManagementService(
            new CryptographyClientBuilder()
                .keyIdentifier(versionedKeyId)
                .credential(credential)
                .buildClient(),
            versionedKeyId);
        return new EncryptedBlobStore(blobContainerClient.getBlobClient(blobName), keyManagement);
    }

    public Mono<AsyncEncryptedBlobStore> asyncBlobStore(String blobName) {
        return keyAsyncClient.getKey(keyName)
            .map(key -> {
                String versionedKeyId = key.getId();
                AsyncKeyManagementService keyManagement = new AsyncKeyManagementService(
                    new CryptographyClientBuilder()
                        .keyIdentifier(versionedKeyId)
                        .credential(credential)
                        .buildAsyncClient(),
                    versionedKeyId);
                return new AsyncEncryptedBlobStore(
                    blobContainerAsyncClient.getBlobAsyncClient(blobName),
                    keyManagement);
            });
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
