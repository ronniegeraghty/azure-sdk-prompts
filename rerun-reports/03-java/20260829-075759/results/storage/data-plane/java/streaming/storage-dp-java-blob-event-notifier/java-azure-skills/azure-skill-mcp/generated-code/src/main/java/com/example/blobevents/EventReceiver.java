package com.example.blobevents;

import java.util.List;
import java.util.logging.Logger;

public final class EventReceiver {
    private static final Logger LOGGER = Logger.getLogger(EventReceiver.class.getName());
    private static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    private static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private final EventPayloadDeserializer deserializer = new EventPayloadDeserializer();

    public List<BlobLifecycleEvent> deserialize(String payload) {
        return deserializer.deserialize(payload);
    }

    public void receive(String payload, BlobEventHandler handler) {
        for (BlobLifecycleEvent event : deserialize(payload)) {
            switch (event.type()) {
                case BLOB_CREATED -> handler.handleCreated(event);
                case BLOB_DELETED -> handler.handleDeleted(event);
                default -> LOGGER.warning(() -> "Ignoring unrecognized event type " + event.type()
                    + " for event " + event.id());
            }
        }
    }
}
