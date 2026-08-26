package com.example.blobevents;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

public final class AsyncEventReceiver {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncEventReceiver.class);
    private final AsyncBlobEventHandler handler;

    public AsyncEventReceiver(AsyncBlobEventHandler handler) {
        this.handler = handler;
    }

    public Mono<Void> receive(String jsonPayload) {
        return Flux.fromIterable(EventPayloadParser.parse(jsonPayload))
            .concatMap(event -> {
                if (BlobEventHandler.BLOB_CREATED.equals(event.type())
                    || BlobEventHandler.BLOB_DELETED.equals(event.type())) {
                    return handler.handle(event);
                }
                LOGGER.warn("Ignoring unrecognized event type: type={}, eventId={}", event.type(), event.id());
                return Mono.empty();
            })
            .then();
    }
}
