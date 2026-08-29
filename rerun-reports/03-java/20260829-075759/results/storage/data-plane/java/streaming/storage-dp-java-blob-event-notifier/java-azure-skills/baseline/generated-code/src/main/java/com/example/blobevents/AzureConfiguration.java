package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

public final class AzureConfiguration {
    private final TokenCredential credential;

    public AzureConfiguration() {
        this.credential = new DefaultAzureCredentialBuilder().build();
    }

    public BlobServiceClient blobServiceClient(String accountEndpoint) {
        return new BlobServiceClientBuilder()
                .endpoint(accountEndpoint)
                .credential(credential)
                .buildClient();
    }

    public BlobServiceAsyncClient blobServiceAsyncClient(String accountEndpoint) {
        return new BlobServiceClientBuilder()
                .endpoint(accountEndpoint)
                .credential(credential)
                .buildAsyncClient();
    }

    public EventGridPublisherClient<EventGridEvent> eventGridPublisherClient(String topicEndpoint) {
        return new EventGridPublisherClientBuilder()
                .endpoint(topicEndpoint)
                .credential(credential)
                .buildEventGridEventPublisherClient();
    }

    public EventGridPublisherAsyncClient<EventGridEvent> eventGridPublisherAsyncClient(String topicEndpoint) {
        return new EventGridPublisherClientBuilder()
                .endpoint(topicEndpoint)
                .credential(credential)
                .buildEventGridEventPublisherAsyncClient();
    }
}
