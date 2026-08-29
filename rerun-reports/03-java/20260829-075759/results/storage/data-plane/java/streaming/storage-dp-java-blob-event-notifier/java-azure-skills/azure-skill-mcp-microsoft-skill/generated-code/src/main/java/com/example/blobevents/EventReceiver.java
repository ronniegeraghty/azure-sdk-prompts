package com.example.blobevents;

import com.azure.core.models.CloudEvent;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.example.blobevents.model.IncomingEvent;
import com.example.blobevents.model.IncomingEvent.Schema;
import java.util.List;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public final class EventReceiver {
    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);
    private final BlobEventHandler handler;

    public EventReceiver(BlobEventHandler handler) {
        this.handler = handler;
    }

    public List<IncomingEvent> receive(String jsonPayload) {
        List<IncomingEvent> events = deserialize(jsonPayload);
        events.forEach(this::route);
        return events;
    }

    public static List<IncomingEvent> deserialize(String jsonPayload) {
        if (jsonPayload == null || jsonPayload.isBlank()) {
            throw new IllegalArgumentException("Event Grid payload must not be empty");
        }

        if (looksLikeCloudEvents(jsonPayload)) {
            return CloudEvent.fromString(jsonPayload).stream()
                .map(event -> new IncomingEvent(
                    Schema.CLOUD_EVENT,
                    event.getId(),
                    event.getType(),
                    event.getSubject(),
                    event.getTime(),
                    event.getData()))
                .toList();
        }

        return EventGridEvent.fromString(jsonPayload).stream()
            .map(event -> new IncomingEvent(
                Schema.EVENT_GRID,
                event.getId(),
                event.getEventType(),
                event.getSubject(),
                event.getEventTime(),
                event.getData()))
            .toList();
    }

    private void route(IncomingEvent event) {
        switch (event.type()) {
            case BlobEventHandler.BLOB_CREATED -> handler.handleCreated(event);
            case BlobEventHandler.BLOB_DELETED -> handler.handleDeleted(event);
            default -> LOGGER.warn("Ignoring unrecognized Event Grid event type: {}", event.type());
        }
    }

    private static boolean looksLikeCloudEvents(String payload) {
        return payload.contains("\"specversion\"") || payload.contains("\"specVersion\"");
    }
}
