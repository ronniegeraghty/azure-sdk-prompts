package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.util.Map;
import java.util.Objects;

public final class AzureConfiguration {
    private final String storageEndpoint;
    private final String eventGridTopicEndpoint;
    private final TokenCredential credential;

    public AzureConfiguration(String storageEndpoint, String eventGridTopicEndpoint) {
        this(storageEndpoint, eventGridTopicEndpoint, new ManagedIdentityCredentialBuilder().build());
    }

    AzureConfiguration(String storageEndpoint, String eventGridTopicEndpoint, TokenCredential credential) {
        this.storageEndpoint = requireHttpsEndpoint(storageEndpoint, "storageEndpoint");
        this.eventGridTopicEndpoint = requireHttpsEndpoint(eventGridTopicEndpoint, "eventGridTopicEndpoint");
        this.credential = Objects.requireNonNull(credential, "credential");
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        return new AzureConfiguration(
                requiredEnvironment(environment, "AZURE_STORAGE_ENDPOINT"),
                requiredEnvironment(environment, "EVENT_GRID_TOPIC_ENDPOINT"));
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

    public EventPublisher eventPublisher() {
        return new EventPublisher(eventGridTopicEndpoint, credential);
    }

    public AsyncEventPublisher asyncEventPublisher() {
        return new AsyncEventPublisher(eventGridTopicEndpoint, credential);
    }

    private static String requiredEnvironment(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static String requireHttpsEndpoint(String endpoint, String name) {
        if (endpoint == null || !endpoint.startsWith("https://")) {
            throw new IllegalArgumentException(name + " must be an HTTPS endpoint");
        }
        return endpoint;
    }
}
