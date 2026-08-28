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
        EventGridPublisherClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
                .endpoint(requireHttps(topicEndpoint))
                .credential(credential)
                .buildEventGridEventPublisherClient();
        this.sender = client::sendEvents;
    }

    EventPublisher(Consumer<List<EventGridEvent>> sender) {
        this.sender = sender;
    }

    public void publish(List<CustomEvent> customEvents) {
        sender.accept(toEventGridEvents(customEvents));
    }

    static List<EventGridEvent> toEventGridEvents(List<CustomEvent> customEvents) {
        return customEvents.stream()
                .map(event -> new EventGridEvent(
                        event.subject(),
                        event.eventType(),
                        BinaryData.fromObject(event.data()),
                        "1.0"))
                .toList();
    }

    static String requireHttps(String endpoint) {
        if (endpoint == null || !endpoint.startsWith("https://")) {
            throw new IllegalArgumentException("Event Grid topic endpoint must use HTTPS");
        }
        return endpoint;
    }
}
