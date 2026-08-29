package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.example.blobevents.storage.AzureBlobStore;
import java.util.Map;
import java.util.Objects;

public final class AzureConfiguration {
    private final String storageAccountUrl;
    private final String eventGridTopicEndpoint;
    private final TokenCredential credential;

    public AzureConfiguration(String storageAccountUrl, String eventGridTopicEndpoint, String managedIdentityClientId) {
        this.storageAccountUrl = requireHttpsUrl(storageAccountUrl, "storageAccountUrl");
        this.eventGridTopicEndpoint = requireHttpsUrl(eventGridTopicEndpoint, "eventGridTopicEndpoint");

        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            builder.clientId(managedIdentityClientId);
        }
        this.credential = builder.build();
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        return new AzureConfiguration(
            environment.get("AZURE_STORAGE_ACCOUNT_URL"),
            environment.get("EVENT_GRID_TOPIC_ENDPOINT"),
            environment.get("AZURE_CLIENT_ID"));
    }

    public AzureBlobStore blobStore() {
        BlobServiceClientBuilder builder = new BlobServiceClientBuilder()
            .endpoint(storageAccountUrl)
            .credential(credential);
        return new AzureBlobStore(builder.buildClient(), builder.buildAsyncClient());
    }

    public EventPublisher eventPublisher() {
        return new EventPublisher(new EventGridPublisherClientBuilder()
            .endpoint(eventGridTopicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherClient());
    }

    public AsyncEventPublisher asyncEventPublisher() {
        return new AsyncEventPublisher(new EventGridPublisherClientBuilder()
            .endpoint(eventGridTopicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherAsyncClient());
    }

    private static String requireHttpsUrl(String value, String name) {
        Objects.requireNonNull(value, name + " is required");
        if (!value.startsWith("https://")) {
            throw new IllegalArgumentException(name + " must use HTTPS");
        }
        return value;
    }
}
