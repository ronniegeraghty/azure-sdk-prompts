package com.example.blobevents.receiver;

import com.example.blobevents.blob.BlobEventHandler;
import com.example.blobevents.model.BlobLifecycleEvent;

import java.util.List;
import java.util.logging.Logger;

public final class EventReceiver {
    private static final Logger LOGGER = Logger.getLogger(EventReceiver.class.getName());
    private static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    private static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private final BlobEventHandler handler;

    public EventReceiver(BlobEventHandler handler) {
        this.handler = handler;
    }

    public List<BlobLifecycleEvent> receive(String jsonPayload) {
        List<BlobLifecycleEvent> events = EventPayloadParser.parse(jsonPayload);
        events.forEach(this::route);
        return events;
    }

    private void route(BlobLifecycleEvent event) {
        switch (event.type()) {
            case BLOB_CREATED -> handler.handleCreated(event);
            case BLOB_DELETED -> handler.handleDeleted(event);
            default -> LOGGER.warning(() -> "Ignoring unrecognized event type: " + event.type());
        }
    }
}
