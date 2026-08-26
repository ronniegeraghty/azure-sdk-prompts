package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;

import java.util.List;
import java.util.function.Consumer;

public final class EventPublisher {
    private final Consumer<List<EventGridEvent>> sender;

    public EventPublisher(String topicEndpoint, TokenCredential credential) {
        this(new EventGridPublisherClientBuilder()
            .endpoint(topicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherClient());
    }

    public EventPublisher(EventGridPublisherClient<EventGridEvent> client) {
        this(client::sendEvents);
    }

    public EventPublisher(Consumer<List<EventGridEvent>> sender) {
        this.sender = sender;
    }

    public void publish(List<CustomEvent> events) {
        if (events == null || events.isEmpty()) {
            throw new IllegalArgumentException("At least one event is required");
        }
        sender.accept(events.stream().map(EventPublisher::toEventGridEvent).toList());
    }

    static EventGridEvent toEventGridEvent(CustomEvent event) {
        return new EventGridEvent(
            event.subject(),
            event.type(),
            BinaryData.fromObject(event.data()),
            event.dataVersion());
    }
}
