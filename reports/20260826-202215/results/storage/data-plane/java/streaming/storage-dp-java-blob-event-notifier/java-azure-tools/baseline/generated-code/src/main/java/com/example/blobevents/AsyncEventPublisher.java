package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import reactor.core.publisher.Mono;

import java.util.List;

public final class AsyncEventPublisher {
    private final AsyncEventGridSender sender;

    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {
        EventGridPublisherAsyncClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
                .endpoint(topicEndpoint)
                .credential(credential)
                .buildEventGridEventPublisherAsyncClient();
        this.sender = client::sendEvents;
    }

    public AsyncEventPublisher(AsyncEventGridSender sender) {
        this.sender = sender;
    }

    public Mono<Void> publish(List<CustomEvent> events) {
        return sender.send(events.stream().map(EventPublisher::toEventGridEvent).toList());
    }
}
