package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import java.util.Objects;

public final class AzureConfiguration {
    private final String storageEndpoint;
    private final String eventGridTopicEndpoint;
    private final TokenCredential credential;

    public AzureConfiguration(
        String storageEndpoint,
        String eventGridTopicEndpoint,
        String managedIdentityClientId
    ) {
        this.storageEndpoint = requireHttpsEndpoint(storageEndpoint, "storageEndpoint");
        this.eventGridTopicEndpoint = requireHttpsEndpoint(eventGridTopicEndpoint, "eventGridTopicEndpoint");
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            builder.clientId(managedIdentityClientId);
        }
        this.credential = builder.build();
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

    public EventGridPublisherClient<EventGridEvent> eventGridPublisherClient() {
        return new EventGridPublisherClientBuilder()
            .endpoint(eventGridTopicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherClient();
    }

    public EventGridPublisherAsyncClient<EventGridEvent> eventGridPublisherAsyncClient() {
        return new EventGridPublisherClientBuilder()
            .endpoint(eventGridTopicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherAsyncClient();
    }

    private static String requireHttpsEndpoint(String endpoint, String name) {
        Objects.requireNonNull(endpoint, name);
        if (!endpoint.startsWith("https://")) {
            throw new IllegalArgumentException(name + " must use HTTPS");
        }
        return endpoint;
    }
}
