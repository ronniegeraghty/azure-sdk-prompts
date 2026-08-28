package com.example.blobevents;

import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import java.util.List;
import java.util.Objects;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public final class EventPublisherAsync {
    private static final Logger LOGGER = LoggerFactory.getLogger(EventPublisherAsync.class);

    private final EventGridPublisherAsyncClient<EventGridEvent> client;

    public EventPublisherAsync(EventGridPublisherAsyncClient<EventGridEvent> client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public static EventPublisherAsync dryRun() {
        return new EventPublisherAsync();
    }

    private EventPublisherAsync() {
        this.client = null;
    }

    public Mono<Void> publish(List<CustomEvent> events) {
        List<EventGridEvent> sdkEvents = EventPublisher.toSdkEvents(events);
        if (client == null) {
            return Mono.fromRunnable(() -> sdkEvents.forEach(event -> LOGGER.info(
                "Dry-run async publish: type='{}', subject='{}'",
                event.getEventType(),
                event.getSubject()
            )));
        }
        return client.sendEvents(sdkEvents);
    }
}
