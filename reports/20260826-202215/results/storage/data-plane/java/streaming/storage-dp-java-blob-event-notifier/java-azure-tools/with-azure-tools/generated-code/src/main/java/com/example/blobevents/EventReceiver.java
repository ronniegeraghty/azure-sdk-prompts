package com.example.blobevents;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public final class EventReceiver {
    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);
    private final BlobEventHandler handler;

    public EventReceiver(BlobEventHandler handler) {
        this.handler = handler;
    }

    public void receive(String jsonPayload) {
        for (BlobLifecycleEvent event : EventPayloadParser.parse(jsonPayload)) {
            if (BlobEventHandler.BLOB_CREATED.equals(event.type())
                || BlobEventHandler.BLOB_DELETED.equals(event.type())) {
                handler.handle(event);
            } else {
                LOGGER.warn("Ignoring unrecognized event type: type={}, eventId={}", event.type(), event.id());
            }
        }
    }
}
