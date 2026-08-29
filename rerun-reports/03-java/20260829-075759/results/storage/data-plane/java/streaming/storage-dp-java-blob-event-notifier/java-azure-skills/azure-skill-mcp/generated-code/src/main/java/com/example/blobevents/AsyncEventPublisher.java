package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Objects;

public final class AsyncEventPublisher {
    private final AsyncEventSender sender;

    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {
        EventGridPublisherAsyncClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
            .endpoint(topicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherAsyncClient();
        this.sender = client::sendEvents;
    }

    AsyncEventPublisher(AsyncEventSender sender) {
        this.sender = Objects.requireNonNull(sender, "sender");
    }

    public Mono<Void> publish(List<CustomEvent> events) {
        return Mono.defer(() -> sender.send(EventPublisher.toEventGridEvents(events)));
    }

    @FunctionalInterface
    interface AsyncEventSender {
        Mono<Void> send(Iterable<EventGridEvent> events);
    }
}
