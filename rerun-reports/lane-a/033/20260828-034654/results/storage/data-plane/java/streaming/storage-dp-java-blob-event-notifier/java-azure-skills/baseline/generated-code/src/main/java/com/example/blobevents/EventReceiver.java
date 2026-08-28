package com.example.blobevents;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

public final class EventReceiver {
    public static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    public static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);
    private final BlobEventHandler handler;

    public EventReceiver(BlobEventHandler handler) {
        this.handler = handler;
    }

    public List<EventEnvelope> receive(String jsonPayload) {
        List<EventEnvelope> events = EventPayloadParser.parse(jsonPayload);
        events.forEach(this::route);
        return events;
    }

    private void route(EventEnvelope event) {
        switch (event.type()) {
            case BLOB_CREATED -> handler.handleCreated(event);
            case BLOB_DELETED -> handler.handleDeleted(event);
            default -> LOGGER.warn("Ignoring unrecognized event type {} for event {}", event.type(), event.id());
        }
    }
}
