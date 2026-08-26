package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.function.Function;

public final class AsyncEventPublisher {
    private final Function<List<EventGridEvent>, Mono<Void>> sender;

    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {
        this(new EventGridPublisherClientBuilder()
            .endpoint(topicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherAsyncClient());
    }

    public AsyncEventPublisher(EventGridPublisherAsyncClient<EventGridEvent> client) {
        this(client::sendEvents);
    }

    public AsyncEventPublisher(Function<List<EventGridEvent>, Mono<Void>> sender) {
        this.sender = sender;
    }

    public Mono<Void> publish(List<CustomEvent> events) {
        if (events == null || events.isEmpty()) {
            return Mono.error(new IllegalArgumentException("At least one event is required"));
        }
        return sender.apply(events.stream().map(EventPublisher::toEventGridEvent).toList());
    }
}
