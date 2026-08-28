package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import java.util.List;
import java.util.Objects;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public final class EventPublisher {
    private static final Logger LOGGER = LoggerFactory.getLogger(EventPublisher.class);

    private final EventGridPublisherClient<EventGridEvent> client;

    public EventPublisher(EventGridPublisherClient<EventGridEvent> client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public static EventPublisher dryRun() {
        return new EventPublisher();
    }

    private EventPublisher() {
        this.client = null;
    }

    public void publish(List<CustomEvent> events) {
        List<EventGridEvent> sdkEvents = toSdkEvents(events);
        if (client == null) {
            sdkEvents.forEach(event -> LOGGER.info(
                "Dry-run publish: type='{}', subject='{}'",
                event.getEventType(),
                event.getSubject()
            ));
            return;
        }
        client.sendEvents(sdkEvents);
    }

    static List<EventGridEvent> toSdkEvents(List<CustomEvent> events) {
        return events.stream()
            .map(event -> new EventGridEvent(
                event.subject(),
                event.eventType(),
                BinaryData.fromObject(event.data()),
                event.dataVersion()
            ))
            .toList();
    }
}
