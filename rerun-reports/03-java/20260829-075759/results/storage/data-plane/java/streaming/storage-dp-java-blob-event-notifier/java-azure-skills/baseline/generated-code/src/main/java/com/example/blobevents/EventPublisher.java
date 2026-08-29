package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherAsyncClient;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import reactor.core.publisher.Mono;

import java.util.List;

public final class EventPublisher {
    private final EventGridPublisherClient<EventGridEvent> syncClient;
    private final EventGridPublisherAsyncClient<EventGridEvent> asyncClient;

    public EventPublisher(EventGridPublisherClient<EventGridEvent> syncClient) {
        this.syncClient = syncClient;
        this.asyncClient = null;
    }

    public EventPublisher(EventGridPublisherAsyncClient<EventGridEvent> asyncClient) {
        this.syncClient = null;
        this.asyncClient = asyncClient;
    }

    public void publish(List<CustomEvent> events) {
        if (syncClient == null) {
            throw new IllegalStateException("This publisher was configured for asynchronous use");
        }
        syncClient.sendEvents(toEventGridEvents(events));
    }

    public Mono<Void> publishAsync(List<CustomEvent> events) {
        if (asyncClient == null) {
            return Mono.error(new IllegalStateException("This publisher was configured for synchronous use"));
        }
        return asyncClient.sendEvents(toEventGridEvents(events));
    }

    private static List<EventGridEvent> toEventGridEvents(List<CustomEvent> events) {
        return events.stream()
                .map(event -> new EventGridEvent(
                        normalizeSubject(event.subject()),
                        event.type(),
                        BinaryData.fromObject(event.data()),
                        event.dataVersion()))
                .toList();
    }

    private static String normalizeSubject(String subject) {
        if (subject == null || subject.isBlank()) {
            throw new IllegalArgumentException("Event subject must not be blank");
        }
        return subject.startsWith("/") ? subject : "/" + subject;
    }

    public record CustomEvent(String subject, String type, String dataVersion, Object data) {
        public CustomEvent {
            if (type == null || type.isBlank()) {
                throw new IllegalArgumentException("Event type must not be blank");
            }
            if (dataVersion == null || dataVersion.isBlank()) {
                throw new IllegalArgumentException("Data version must not be blank");
            }
        }
    }
}
