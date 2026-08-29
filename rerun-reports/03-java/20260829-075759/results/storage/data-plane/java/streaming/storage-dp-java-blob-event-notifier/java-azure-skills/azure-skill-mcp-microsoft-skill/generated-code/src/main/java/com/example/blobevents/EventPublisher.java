package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.example.blobevents.model.CustomEvent;
import java.util.List;
import java.util.function.Consumer;

public final class EventPublisher {
    private final Consumer<List<EventGridEvent>> sender;

    public EventPublisher(EventGridPublisherClient<EventGridEvent> client) {
        this(client::sendEvents);
    }

    public EventPublisher(Consumer<List<EventGridEvent>> sender) {
        this.sender = sender;
    }

    public void publish(List<CustomEvent> customEvents) {
        if (customEvents == null || customEvents.isEmpty()) {
            throw new IllegalArgumentException("At least one custom event is required");
        }
        sender.accept(customEvents.stream().map(EventPublisher::toEventGridEvent).toList());
    }

    static EventGridEvent toEventGridEvent(CustomEvent event) {
        return new EventGridEvent(
            event.subject(),
            event.type(),
            BinaryData.fromObject(event.data()),
            event.dataVersion());
    }
}
