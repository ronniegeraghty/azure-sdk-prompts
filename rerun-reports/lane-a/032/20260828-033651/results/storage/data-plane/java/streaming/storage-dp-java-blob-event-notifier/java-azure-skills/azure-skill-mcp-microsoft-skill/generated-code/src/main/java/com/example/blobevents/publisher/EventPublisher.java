package com.example.blobevents.publisher;

import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;
import com.example.blobevents.model.CustomEvent;
import com.azure.core.credential.TokenCredential;

import java.util.List;
import java.util.Objects;
import java.util.function.Consumer;

public final class EventPublisher {
    private final Consumer<List<EventGridEvent>> sender;

    public EventPublisher(String topicEndpoint, TokenCredential credential) {
        EventGridPublisherClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
            .endpoint(Objects.requireNonNull(topicEndpoint, "topicEndpoint"))
            .credential(Objects.requireNonNull(credential, "credential"))
            .buildEventGridEventPublisherClient();
        this.sender = client::sendEvents;
    }

    public EventPublisher(Consumer<List<EventGridEvent>> sender) {
        this.sender = Objects.requireNonNull(sender, "sender");
    }

    public void publish(List<CustomEvent> events) {
        sender.accept(events.stream().map(EventPublisher::toEventGridEvent).toList());
    }

    static EventGridEvent toEventGridEvent(CustomEvent event) {
        return new EventGridEvent(
            event.subject(), event.type(), BinaryData.fromObject(event.data()), event.dataVersion());
    }
}
