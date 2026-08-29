package com.example.blobevents;

import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.example.blobevents.model.CustomEvent;
import java.util.List;
import java.util.function.Function;
import reactor.core.publisher.Mono;

public final class AsyncEventPublisher {
    private final Function<List<EventGridEvent>, Mono<Void>> sender;

    public AsyncEventPublisher(EventGridPublisherAsyncClient<EventGridEvent> client) {
        this(client::sendEvents);
    }

    public AsyncEventPublisher(Function<List<EventGridEvent>, Mono<Void>> sender) {
        this.sender = sender;
    }

    public Mono<Void> publish(List<CustomEvent> customEvents) {
        if (customEvents == null || customEvents.isEmpty()) {
            return Mono.error(new IllegalArgumentException("At least one custom event is required"));
        }
        return sender.apply(customEvents.stream().map(EventPublisher::toEventGridEvent).toList());
    }
}
