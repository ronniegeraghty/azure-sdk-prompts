package com.example.blobevents;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

public final class EventReceiver {
    public static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    public static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);

    private final EventPayloadParser parser;
    private final BlobEventProcessor processor;

    public EventReceiver(BlobEventProcessor processor) {
        this(processor, new ObjectMapper());
    }

    EventReceiver(BlobEventProcessor processor, ObjectMapper objectMapper) {
        this.processor = processor;
        this.parser = new EventPayloadParser(objectMapper);
    }

    public List<BlobLifecycleEvent> receive(String jsonPayload) {
        List<BlobLifecycleEvent> events = parser.parse(jsonPayload);
        events.forEach(this::route);
        return events;
    }

    private void route(BlobLifecycleEvent event) {
        switch (event.type()) {
            case BLOB_CREATED -> processor.onBlobCreated(event);
            case BLOB_DELETED -> processor.onBlobDeleted(event);
            default -> LOGGER.warn("Ignoring unsupported event type '{}' (event id {})", event.type(), event.id());
        }
    }
}
