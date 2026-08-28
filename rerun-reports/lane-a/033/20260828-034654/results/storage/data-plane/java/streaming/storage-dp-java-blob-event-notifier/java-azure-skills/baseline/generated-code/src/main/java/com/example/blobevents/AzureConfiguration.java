package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class AzureConfiguration {
    private final TokenCredential credential;
    private final BlobServiceClient blobServiceClient;
    private final BlobServiceAsyncClient blobServiceAsyncClient;

    public AzureConfiguration(String storageAccountEndpoint) {
        this.credential = new DefaultAzureCredentialBuilder().build();
        BlobServiceClientBuilder builder = new BlobServiceClientBuilder()
                .endpoint(requireHttps(storageAccountEndpoint))
                .credential(credential);
        this.blobServiceClient = builder.buildClient();
        this.blobServiceAsyncClient = builder.buildAsyncClient();
    }

    public BlobEventHandler blobEventHandler() {
        AzureBlobOperations operations = new AzureBlobOperations(blobServiceClient, blobServiceAsyncClient);
        return new BlobEventHandler(operations, operations);
    }

    public EventPublisher eventPublisher(String topicEndpoint) {
        return new EventPublisher(topicEndpoint, credential);
    }

    public AsyncEventPublisher asyncEventPublisher(String topicEndpoint) {
        return new AsyncEventPublisher(topicEndpoint, credential);
    }

    private static String requireHttps(String endpoint) {
        if (endpoint == null || !endpoint.startsWith("https://")) {
            throw new IllegalArgumentException("Azure Storage endpoint must use HTTPS");
        }
        return endpoint;
    }
}
