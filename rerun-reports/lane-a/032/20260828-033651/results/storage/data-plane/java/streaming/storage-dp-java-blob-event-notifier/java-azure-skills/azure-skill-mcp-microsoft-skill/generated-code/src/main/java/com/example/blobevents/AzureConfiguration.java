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
    private final String storageEndpoint;
    private final String topicEndpoint;
    private final TokenCredential credential;

    public AzureConfiguration(String storageEndpoint, String topicEndpoint, String managedIdentityClientId) {
        this.storageEndpoint = requireHttps(storageEndpoint, "storageEndpoint");
        this.topicEndpoint = requireHttps(topicEndpoint, "topicEndpoint");
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        this.credential = credentialBuilder.build();
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        return new AzureConfiguration(
            required(environment, "AZURE_STORAGE_ACCOUNT_URL"),
            required(environment, "EVENT_GRID_TOPIC_ENDPOINT"),
            environment.get("AZURE_CLIENT_ID"));
    }

    public BlobServiceClient blobServiceClient() {
        return new BlobServiceClientBuilder().endpoint(storageEndpoint).credential(credential).buildClient();
    }

    public BlobServiceAsyncClient blobServiceAsyncClient() {
        return new BlobServiceClientBuilder().endpoint(storageEndpoint).credential(credential).buildAsyncClient();
    }

    public EventGridPublisherClient<EventGridEvent> eventPublisherClient() {
        return new EventGridPublisherClientBuilder().endpoint(topicEndpoint).credential(credential)
            .buildEventGridEventPublisherClient();
    }

    public EventGridPublisherAsyncClient<EventGridEvent> asyncEventPublisherClient() {
        return new EventGridPublisherClientBuilder().endpoint(topicEndpoint).credential(credential)
            .buildEventGridEventPublisherAsyncClient();
    }

    public TokenCredential credential() {
        return credential;
    }

    public String topicEndpoint() {
        return topicEndpoint;
    }

    private static String required(Map<String, String> values, String name) {
        String value = values.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is missing: " + name);
        }
        return value;
    }

    private static String requireHttps(String value, String name) {
        if (value == null || !value.startsWith("https://")) {
            throw new IllegalArgumentException(name + " must use HTTPS");
        }
        return value;
    }
}
