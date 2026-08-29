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
import reactor.core.publisher.Mono;

import java.util.Map;

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
            String keyName) {
        this.credential = credential;
        this.blobContainerClient = blobContainerClient;
        this.blobContainerAsyncClient = blobContainerAsyncClient;
        this.keyClient = keyClient;
        this.keyAsyncClient = keyAsyncClient;
        this.keyName = keyName;
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        String blobEndpoint = required(environment, "AZURE_STORAGE_BLOB_ENDPOINT");
        String containerName = required(environment, "AZURE_STORAGE_CONTAINER");
        String vaultEndpoint = required(environment, "AZURE_KEY_VAULT_ENDPOINT");
        String keyName = required(environment, "AZURE_KEY_VAULT_KEY_NAME");

        ManagedIdentityCredentialBuilder credentialBuilder =
                new ManagedIdentityCredentialBuilder();
        String clientId = environment.get("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();

        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()
                .endpoint(blobEndpoint)
                .credential(credential);
        BlobServiceClient syncBlobService = blobBuilder.buildClient();
        BlobServiceAsyncClient asyncBlobService = blobBuilder.buildAsyncClient();

        KeyClientBuilder keyBuilder = new KeyClientBuilder()
                .vaultUrl(vaultEndpoint)
                .credential(credential);

        return new AzureConfiguration(
                credential,
                syncBlobService.getBlobContainerClient(containerName),
                asyncBlobService.getBlobContainerAsyncClient(containerName),
                keyBuilder.buildClient(),
                keyBuilder.buildAsyncClient(),
                keyName);
    }

    public EncryptedBlobClient encryptedBlobClient() {
        KeyManagementService keyManagement =
                new KeyManagementService(keyClient, credential, keyName);
        return new EncryptedBlobClient(blobContainerClient, keyManagement);
    }

    public Mono<EncryptedBlobAsyncClient> encryptedBlobAsyncClient() {
        return KeyManagementAsyncService.create(keyAsyncClient, credential, keyName)
                .map(keyManagement ->
                        new EncryptedBlobAsyncClient(blobContainerAsyncClient, keyManagement));
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
