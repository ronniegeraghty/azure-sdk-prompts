package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.util.Map;

public final class AzureConfiguration {
    public static final String STORAGE_ACCOUNT_URL = "AZURE_STORAGE_ACCOUNT_URL";
    public static final String EVENT_GRID_TOPIC_ENDPOINT = "EVENT_GRID_TOPIC_ENDPOINT";
    public static final String MANAGED_IDENTITY_CLIENT_ID = "AZURE_CLIENT_ID";

    private final String storageAccountUrl;
    private final String eventGridTopicEndpoint;
    private final TokenCredential credential;

    public AzureConfiguration(String storageAccountUrl, String eventGridTopicEndpoint, String managedIdentityClientId) {
        this.storageAccountUrl = requireHttpsUrl(storageAccountUrl, "storageAccountUrl");
        this.eventGridTopicEndpoint = requireHttpsUrl(eventGridTopicEndpoint, "eventGridTopicEndpoint");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        this.credential = credentialBuilder.build();
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        return new AzureConfiguration(
            requireEnvironmentVariable(environment, STORAGE_ACCOUNT_URL),
            requireEnvironmentVariable(environment, EVENT_GRID_TOPIC_ENDPOINT),
            environment.get(MANAGED_IDENTITY_CLIENT_ID));
    }

    public BlobServiceClient blobServiceClient() {
        return new BlobServiceClientBuilder()
            .endpoint(storageAccountUrl)
            .credential(credential)
            .buildClient();
    }

    public BlobServiceAsyncClient blobServiceAsyncClient() {
        return new BlobServiceClientBuilder()
            .endpoint(storageAccountUrl)
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

    public EventPublisher eventPublisher() {
        return new EventPublisher(eventGridTopicEndpoint, credential);
    }

    public AsyncEventPublisher asyncEventPublisher() {
        return new AsyncEventPublisher(eventGridTopicEndpoint, credential);
    }

    private static String requireEnvironmentVariable(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static String requireHttpsUrl(String value, String name) {
        if (value == null || !value.startsWith("https://")) {
            throw new IllegalArgumentException(name + " must be an HTTPS URL");
        }
        return value;
    }
}
