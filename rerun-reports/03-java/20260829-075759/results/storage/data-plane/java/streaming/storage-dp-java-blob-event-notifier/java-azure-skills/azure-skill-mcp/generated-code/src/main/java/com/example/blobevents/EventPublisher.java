package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;

import java.util.List;
import java.util.Objects;

public final class EventPublisher {
    private final SyncEventSender sender;

    public EventPublisher(String topicEndpoint, TokenCredential credential) {
        EventGridPublisherClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
            .endpoint(topicEndpoint)
            .credential(credential)
            .buildEventGridEventPublisherClient();
        this.sender = client::sendEvents;
    }

    EventPublisher(SyncEventSender sender) {
        this.sender = Objects.requireNonNull(sender, "sender");
    }

    public void publish(List<CustomEvent> events) {
        sender.send(toEventGridEvents(events));
    }

    static List<EventGridEvent> toEventGridEvents(List<CustomEvent> events) {
        Objects.requireNonNull(events, "events");
        if (events.isEmpty()) {
            throw new IllegalArgumentException("At least one custom event is required");
        }
        return events.stream()
            .map(event -> new EventGridEvent(
                event.subject(),
                event.eventType(),
                BinaryData.fromObject(event.data()),
                event.dataVersion()))
            .toList();
    }

    @FunctionalInterface
    interface SyncEventSender {
        void send(Iterable<EventGridEvent> events);
    }
}
