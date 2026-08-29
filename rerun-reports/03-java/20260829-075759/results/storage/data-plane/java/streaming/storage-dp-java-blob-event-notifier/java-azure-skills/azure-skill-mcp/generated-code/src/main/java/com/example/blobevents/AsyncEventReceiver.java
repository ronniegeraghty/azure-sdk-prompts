package com.example.blobevents;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.logging.Logger;

public final class AsyncEventReceiver {
    private static final Logger LOGGER = Logger.getLogger(AsyncEventReceiver.class.getName());
    private static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    private static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private final EventPayloadDeserializer deserializer = new EventPayloadDeserializer();

    public Flux<BlobLifecycleEvent> deserialize(String payload) {
        return Flux.fromIterable(deserializer.deserialize(payload));
    }

    public Mono<Void> receive(String payload, BlobEventHandler handler) {
        return deserialize(payload)
            .concatMap(event -> switch (event.type()) {
                case BLOB_CREATED -> handler.handleCreatedAsync(event);
                case BLOB_DELETED -> handler.handleDeletedAsync(event);
                default -> {
                    LOGGER.warning(() -> "Ignoring unrecognized event type " + event.type()
                        + " for event " + event.id());
                    yield Mono.empty();
                }
            })
            .then();
    }
}
