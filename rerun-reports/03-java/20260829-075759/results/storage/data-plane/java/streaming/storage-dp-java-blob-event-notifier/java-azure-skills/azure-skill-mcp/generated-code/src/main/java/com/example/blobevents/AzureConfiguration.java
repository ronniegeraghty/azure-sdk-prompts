package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.net.URI;
import java.util.Map;
import java.util.Objects;

public final class AzureConfiguration {
    private final String storageEndpoint;
    private final String eventGridTopicEndpoint;
    private final TokenCredential credential;

    private AzureConfiguration(
        String storageEndpoint,
        String eventGridTopicEndpoint,
        TokenCredential credential
    ) {
        this.storageEndpoint = requireHttpsEndpoint(storageEndpoint, "storageEndpoint");
        this.eventGridTopicEndpoint = requireHttpsEndpoint(eventGridTopicEndpoint, "eventGridTopicEndpoint");
        this.credential = Objects.requireNonNull(credential, "credential");
    }

    public static AzureConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static AzureConfiguration fromEnvironment(Map<String, String> environment) {
        String clientId = environment.get("AZURE_CLIENT_ID");
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        return new AzureConfiguration(
            required(environment, "AZURE_STORAGE_BLOB_ENDPOINT"),
            required(environment, "AZURE_EVENTGRID_TOPIC_ENDPOINT"),
            credentialBuilder.build());
    }

    public BlobServiceClient blobServiceClient() {
        return blobClientBuilder().buildClient();
    }

    public BlobServiceAsyncClient blobServiceAsyncClient() {
        return blobClientBuilder().buildAsyncClient();
    }

    public EventPublisher eventPublisher() {
        return new EventPublisher(eventGridTopicEndpoint, credential);
    }

    public AsyncEventPublisher asyncEventPublisher() {
        return new AsyncEventPublisher(eventGridTopicEndpoint, credential);
    }

    private BlobServiceClientBuilder blobClientBuilder() {
        return new BlobServiceClientBuilder()
            .endpoint(storageEndpoint)
            .credential(credential);
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Missing required environment variable " + name);
        }
        return value;
    }

    private static String requireHttpsEndpoint(String value, String name) {
        URI uri;
        try {
            uri = URI.create(Objects.requireNonNull(value, name));
        } catch (IllegalArgumentException exception) {
            throw new IllegalArgumentException(name + " must be a valid URI", exception);
        }
        if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {
            throw new IllegalArgumentException(name + " must be an absolute HTTPS endpoint");
        }
        return value;
    }
}
