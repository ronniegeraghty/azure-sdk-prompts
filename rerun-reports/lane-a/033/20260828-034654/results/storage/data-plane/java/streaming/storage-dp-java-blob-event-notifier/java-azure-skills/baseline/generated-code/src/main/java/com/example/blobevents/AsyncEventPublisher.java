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
        EventGridPublisherAsyncClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
                .endpoint(EventPublisher.requireHttps(topicEndpoint))
                .credential(credential)
                .buildEventGridEventPublisherAsyncClient();
        this.sender = client::sendEvents;
    }

    AsyncEventPublisher(Function<List<EventGridEvent>, Mono<Void>> sender) {
        this.sender = sender;
    }

    public Mono<Void> publish(List<CustomEvent> customEvents) {
        return sender.apply(EventPublisher.toEventGridEvents(customEvents));
    }
}
