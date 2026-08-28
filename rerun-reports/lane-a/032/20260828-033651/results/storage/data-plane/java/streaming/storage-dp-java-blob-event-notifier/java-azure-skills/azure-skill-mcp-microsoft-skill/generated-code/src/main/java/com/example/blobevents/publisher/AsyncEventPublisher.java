package com.example.blobevents.publisher;

import com.azure.core.credential.TokenCredential;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.example.blobevents.model.CustomEvent;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Objects;
import java.util.function.Function;

public final class AsyncEventPublisher {
    private final Function<List<EventGridEvent>, Mono<Void>> sender;

    public AsyncEventPublisher(String topicEndpoint, TokenCredential credential) {
        EventGridPublisherAsyncClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
            .endpoint(Objects.requireNonNull(topicEndpoint, "topicEndpoint"))
            .credential(Objects.requireNonNull(credential, "credential"))
            .buildEventGridEventPublisherAsyncClient();
        this.sender = client::sendEvents;
    }

    public AsyncEventPublisher(Function<List<EventGridEvent>, Mono<Void>> sender) {
        this.sender = Objects.requireNonNull(sender, "sender");
    }

    public Mono<Void> publish(List<CustomEvent> events) {
        return sender.apply(events.stream().map(EventPublisher::toEventGridEvent).toList());
    }
}
